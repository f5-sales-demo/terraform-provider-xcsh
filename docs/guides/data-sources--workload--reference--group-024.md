---
page_title: "xcsh_workload reference"
subcategory: "Container"
description: "Complete grouped canonical reference for xcsh_workload reference."
---

# xcsh_workload reference

<a id="canonical-3032333121223003-2021222320021023-0310030000101032-3030002111323031-0323203132331231-1212113333002102-1120310330001323-2313121221133221"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.redirect_route.headers — headers / 021320231210 / 2

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
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.redirect_route](data-sources--workload--reference--group-023.md#canonical-1132000213131301-3210120212301310-3301022312202023-2233211123100020-0010000313311330-3133001202320010-3022300313000021-0233300001331101)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.redirect_route.headers

<a id="canonical-3303130110332110-1232232103323103-2133321212310012-1233312132232213-0022022330120130-0302030231322100-3132300031230320-1131010133331112"></a>

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

<a id="canonical-2033221330020201-2302030102031112-1312300000333322-3021131323322332-3100312101030330-0231133231100311-1131333222013311-2200113303121113"></a>

## Direct properties — headers / 021320231210 / 3

<a id="canonical-2232110002112300-1321210012231013-3120323230032133-2131232103210203-1110311313033120-2100021331201012-3332002231001033-0201333333133221"></a>

<a id="canonical-0120023012111023-1003111132100133-3020100221303022-3202100002032103-3030103200213301-1030102301003130-0130232330021122-2221032200102013"></a>

## exact property — headers / 021320231210 / 4

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

<a id="canonical-3211012311233020-1011321303201312-1012332313131111-1222313121212003-2000030132210331-3130033303123200-3032030113221331-0222130022121201"></a>

<a id="canonical-1333012012311123-1010232032120333-0322031133201021-0200033202033021-1002102111131200-0101331233322130-0210313220130012-0132210211210221"></a>

## invert_match property — headers / 021320231210 / 5

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

<a id="canonical-2130203131111011-3101023321002200-3210231301233021-1012021203321030-2212331121121122-3230313330222231-3122002313011131-3323000130321003"></a>

<a id="canonical-3121300011303203-0003232111013112-3023200223112323-0312022201200310-3031331120303233-0332201333002121-1011332320111133-1013213213031012"></a>

## name property — headers / 021320231210 / 6

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

<a id="canonical-1212030201103133-3022312132210012-1113032302023003-2301012133000301-0300331230320223-2313132301031300-0330303103111311-3220311123230113"></a>

<a id="canonical-3232313102300232-2133022030231331-0321203133033113-1201021320010123-3222000121032010-0322300221213201-0103203000310333-2121323323323130"></a>

## presence property — headers / 021320231210 / 7

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

<a id="canonical-0312031231300321-1021233332031200-0232221202322312-3202113012031321-2002131032330330-3123101232231310-0000023020103200-2331221010110002"></a>

<a id="canonical-3332002212021213-1003201320120120-0002120311221321-3021230333330211-1003102031000323-3102321203212221-3020121220130221-2331220320211100"></a>

## regex property — headers / 021320231210 / 8

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

<a id="canonical-1001021102110323-3030113221200333-0313310211102201-0122301110111021-2010201131003320-3300231320201003-1013013002011312-1331003011110131"></a>

## Next pages — headers / 021320231210 / 9

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.redirect_route](data-sources--workload--reference--group-023.md#canonical-1132000213131301-3210120212301310-3301022312202023-2233211123100020-0010000313311330-3133001202320010-3022300313000021-0233300001331101)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-3111310222120011-1312331330223212-3133120333322312-2233123210220020-0032013011232223-1333032020012112-1132100033233330-2222202321201012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1233210101231033-2112002220302120-1311212003201110-3032113002310013-1200221231012013-2232301100021222-2330212322123133-1320022313001123"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.redirect_route.incoming_port — incoming_port / 200033321002 / 2

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
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.redirect_route](data-sources--workload--reference--group-023.md#canonical-1132000213131301-3210120212301310-3301022312202023-2233211123100020-0010000313311330-3133001202320010-3022300313000021-0233300001331101)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.redirect_route.incoming_port

<a id="canonical-2011321023030202-0111211020021030-0031301101111210-2322333020112301-3232321121111310-3301011112322001-2321110303123111-2103203033213031"></a>

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

<a id="canonical-0321330002032100-2012123320331100-0220222121032212-0021301230000020-2122333321301021-0230211210322310-2132321220320010-3221022112122031"></a>

## Direct properties — incoming_port / 200033321002 / 3

- [no_port_match](data-sources--workload--reference--group-024.md#canonical-2101233030121101-2030032211220002-1323301120203100-0013030300320013-3321030021303313-1330302333000132-2021000001120330-1021320102031111): complete subsection reference.

<a id="canonical-0201323202220203-3032123112331210-3323031031111221-1013130230111012-2013312320113203-2213321321211022-1303021321131001-0100120113100332"></a>

<a id="canonical-3113111221011232-2130112211131220-1013112131232213-0011211201131032-0023102112212200-2110123100013203-1022131322003331-3222233001210101"></a>

## port property — incoming_port / 200033321002 / 4

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

<a id="canonical-2123002203011331-0113123101320031-1230332033313123-3200132001002023-3110303212310001-0012221000002011-3113201210202322-2303010121031021"></a>

<a id="canonical-0013213212121102-0320212303322302-2333002221200330-3003202221220122-1302021123221132-3231123300301210-0301131121313301-1123211211023203"></a>

## port_ranges property — incoming_port / 200033321002 / 5

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

<a id="canonical-0023030202313113-2213101320100101-1310101210113200-3122133322301010-2213001311023210-1113311223333231-3132302313011202-1212301132231133"></a>

## Next pages — incoming_port / 200033321002 / 6

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.redirect_route.incoming_port.no_port_match](data-sources--workload--reference--group-024.md#canonical-2101233030121101-2030032211220002-1323301120203100-0013030300320013-3321030021303313-1330302333000132-2021000001120330-1021320102031111)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.redirect_route](data-sources--workload--reference--group-023.md#canonical-1132000213131301-3210120212301310-3301022312202023-2233211123100020-0010000313311330-3133001202320010-3022300313000021-0233300001331101)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-2101233030121101-2030032211220002-1323301120203100-0013030300320013-3321030021303313-1330302333000132-2021000001120330-1021320102031111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2302203030013103-0113222110223321-2012222032333231-1202223000132330-0301012112131330-3332212330310113-1120010213130022-2112331210113023"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.redirect_route.incoming_port.no_port_match — no_port_match / 123030312130 / 2

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
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.redirect_route](data-sources--workload--reference--group-023.md#canonical-1132000213131301-3210120212301310-3301022312202023-2233211123100020-0010000313311330-3133001202320010-3022300313000021-0233300001331101)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.redirect_route.incoming_port](data-sources--workload--reference--group-024.md#canonical-3111310222120011-1312331330223212-3133120333322312-2233123210220020-0032013011232223-1333032020012112-1132100033233330-2222202321201012)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.redirect_route.incoming_port.no_port_match

<a id="canonical-0300133110031120-1322031111122032-0211021022213220-2131113313230332-1012333312101123-1200131300212002-2103013010103133-3213133320231000"></a>

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

<a id="canonical-3330231112211112-2300330210133102-0032321030303131-1101301312002221-2023122320222020-3233122012031130-1010003120233013-2313212033201321"></a>

## Direct properties — no_port_match / 123030312130 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0132322001310131-0332103123302213-0123113023123002-3200222322122321-1201200112101111-3331213312201123-1002001121130212-3301012212013001"></a>

## Next pages — no_port_match / 123030312130 / 4

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.redirect_route.incoming_port](data-sources--workload--reference--group-024.md#canonical-3111310222120011-1312331330223212-3133120333322312-2233123210220020-0032013011232223-1333032020012112-1132100033233330-2222202321201012)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-2030021223211200-0220022322131110-3000001103231221-1102213111213203-0301100030020313-0022322101021100-3031031012212133-0023103012332120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0112010221223233-3101130201111102-0131022302222331-2011302333210232-0103023211131233-2333112313211011-0332310322030110-2011030302210112"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.redirect_route.path — path / 223302332203 / 2

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
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.redirect_route](data-sources--workload--reference--group-023.md#canonical-1132000213131301-3210120212301310-3301022312202023-2233211123100020-0010000313311330-3133001202320010-3022300313000021-0233300001331101)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.redirect_route.path

<a id="canonical-1320033220030030-1031322113311020-1203030202133301-1001301320313021-2012013101030011-0313231111223123-3110332330032123-2211300211300132"></a>

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

<a id="canonical-3000203110120000-2120333210131313-3300332300302023-1233123010230010-3012032300220211-2310031330000132-1231322003331200-3020100003033313"></a>

## Direct properties — path / 223302332203 / 3

<a id="canonical-2121213103013121-1320021212000010-2220302001303210-3231332000303311-0030221313211123-3011332230320233-3003320121320330-0200221013230303"></a>

<a id="canonical-0010000231312101-2102130121211213-1231303110333133-0122322213023023-2110320032233220-0310311112103123-0222113231231332-0321310003103233"></a>

## path property — path / 223302332203 / 4

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

<a id="canonical-3311100230232102-2321111230302210-3102001322012312-2133130033201233-1021310301233110-2111121312133112-3202222132122222-0012121331130110"></a>

<a id="canonical-2002022000120131-3022311202131123-1313031000333013-0132101200300120-2110023212101333-3131023331232120-3100021010011010-0121102310012331"></a>

## prefix property — path / 223302332203 / 5

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

<a id="canonical-3013321213332021-0213220211020220-0102223331202012-3123033022330021-2332222133033030-2002103310132220-0033120233233122-2321032013210002"></a>

<a id="canonical-2001110233032301-2231212221211011-2322302121210011-3133212131022212-3131222003000201-0133212013302223-3000113313132213-3010321031100203"></a>

## regex property — path / 223302332203 / 6

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

<a id="canonical-3223101301112222-2222220100021011-3321322322223032-0110021113331220-3230203222113033-0100103032313131-2030130213012123-0321211303321132"></a>

## Next pages — path / 223302332203 / 7

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.redirect_route](data-sources--workload--reference--group-023.md#canonical-1132000213131301-3210120212301310-3301022312202023-2233211123100020-0010000313311330-3133001202320010-3022300313000021-0233300001331101)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-2323211032212033-2302322032020303-1333300100001201-2220011220303100-3333113102013320-3013212233000332-3122202001112233-2233313303121232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1123122200003133-1203011212311230-0231102123121101-1310010210330001-3311232103022231-0310301321223000-0013022020112132-1211002121231002"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect — route_redirect / 132302012202 / 2

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
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.redirect_route](data-sources--workload--reference--group-023.md#canonical-1132000213131301-3210120212301310-3301022312202023-2233211123100020-0010000313311330-3133001202320010-3022300313000021-0233300001331101)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect

<a id="canonical-0021111231333032-2212223002323331-2003100233230302-2013201000300133-0303201113300101-2120121223300012-1130102230323320-1000101323310213"></a>

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

<a id="canonical-1102032132323333-0301000211300233-1123302033100320-2112302310111310-1303200111310220-0012121212102230-2303030310222122-1003233023010000"></a>

## Direct properties — route_redirect / 132302012202 / 3

<a id="canonical-2200102302031021-0123321320301211-1001001133301223-0110121032232100-3213002323021222-2123003200113011-0133301000230221-3032100101330133"></a>

<a id="canonical-2000303131102013-2132002201112112-1230123333223330-0310212021100021-3000003310022220-3023112331111103-3222300121211122-0302212032203302"></a>

## host_redirect property — route_redirect / 132302012202 / 4

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

<a id="canonical-1000102133001300-0011033302222131-1222200102000321-2320233301210113-0333233020103221-0222301323032012-0021213103312113-1320032110013301"></a>

<a id="canonical-3321120103302323-2012121203322011-3020000103102030-1012213200001212-2120123112230001-2123133203100122-1230020123320121-2020331321101110"></a>

## path_redirect property — route_redirect / 132302012202 / 5

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

<a id="canonical-0232120031233313-0132221331010013-0331211003101011-0102032013232322-2001232012300010-0221330210333233-2311020322233003-1321301323213220"></a>

<a id="canonical-1201201130331221-3033203130232120-3312333322033031-3012111031333310-0002301132221003-1233230013311231-3311311220122121-3331202332003131"></a>

## prefix_rewrite property — route_redirect / 132302012202 / 6

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

<a id="canonical-0321213330213201-1021110103011020-0333322012001112-0203232230022021-2311311221102310-1133300113313003-1030302003212001-3011310132113313"></a>

<a id="canonical-1023023033230302-0333112010131102-0233101132210000-1330321130131232-3313112203133323-1002320310113110-0023113011102220-3322023013203022"></a>

## proto_redirect property — route_redirect / 132302012202 / 7

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

- [remove_all_params](data-sources--workload--reference--group-024.md#canonical-0320031201103001-3330012021333233-3131201123213222-3033123213233131-1132232110300023-0113110203330031-2230132210323113-2321001222332311): complete subsection reference.

<a id="canonical-2120010030230122-0121313030001333-2303330303230122-1133321002122110-3302132000011310-3112113033133001-1332203030003000-3010232012330131"></a>

<a id="canonical-0121131321301000-1120231230302112-1222232013110001-0111203333021121-0221001311011202-3230110322031301-2120221232220122-3003113200221313"></a>

## replace_params property — route_redirect / 132302012202 / 8

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

<a id="canonical-3121033312333231-3130212220031132-1002120232213203-0021122221010221-1220311020003231-2003231112110100-1323030103323030-0113321023300001"></a>

<a id="canonical-2312032103003010-1012312000101003-2230101002131211-2133133330232233-1233233230110232-3130230321000220-3013102322211013-0003031220021330"></a>

## response_code property — route_redirect / 132302012202 / 9

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

- [retain_all_params](data-sources--workload--reference--group-024.md#canonical-0231020032203210-1300103120320111-2313300213322233-1120123332320003-2200212232211310-0222031020030132-1010012101211200-2001032302013111): complete subsection reference.

<a id="canonical-1313302102322011-0030320110311210-1010313010303211-0221300211102202-1200201302303112-1033013111222133-3302030213112011-0110002132320113"></a>

## Next pages — route_redirect / 132302012202 / 10

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect.remove_all_params](data-sources--workload--reference--group-024.md#canonical-0320031201103001-3330012021333233-3131201123213222-3033123213233131-1132232110300023-0113110203330031-2230132210323113-2321001222332311)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect.retain_all_params](data-sources--workload--reference--group-024.md#canonical-0231020032203210-1300103120320111-2313300213322233-1120123332320003-2200212232211310-0222031020030132-1010012101211200-2001032302013111)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.redirect_route](data-sources--workload--reference--group-023.md#canonical-1132000213131301-3210120212301310-3301022312202023-2233211123100020-0010000313311330-3133001202320010-3022300313000021-0233300001331101)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-0320031201103001-3330012021333233-3131201123213222-3033123213233131-1132232110300023-0113110203330031-2230132210323113-2321001222332311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1201101122030313-0202230011200003-2312031113101312-3221130101222220-0212113323121100-1220222202120021-3123321000323313-2332021133000222"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect.remove_all_params — remove_all_params / 332320330120 / 2

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
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.redirect_route](data-sources--workload--reference--group-023.md#canonical-1132000213131301-3210120212301310-3301022312202023-2233211123100020-0010000313311330-3133001202320010-3022300313000021-0233300001331101)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect](data-sources--workload--reference--group-024.md#canonical-2323211032212033-2302322032020303-1333300100001201-2220011220303100-3333113102013320-3013212233000332-3122202001112233-2233313303121232)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect.remove_all_params

<a id="canonical-2212320232213201-3310230320102222-3202221000031133-0113313223001130-0101102132323113-2032023111321113-1231130301223033-1233333311102131"></a>

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

<a id="canonical-0120330222111231-0301132301013133-3332111120302033-0021032033110213-2032312101001231-3112223011000113-0033333111203120-3200203020122303"></a>

## Direct properties — remove_all_params / 332320330120 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3221220132313211-1033333300333201-3233313300010333-1000012201010121-0111000203311213-0002312022202233-3101100230203301-2022332311123301"></a>

## Next pages — remove_all_params / 332320330120 / 4

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect](data-sources--workload--reference--group-024.md#canonical-2323211032212033-2302322032020303-1333300100001201-2220011220303100-3333113102013320-3013212233000332-3122202001112233-2233313303121232)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-0231020032203210-1300103120320111-2313300213322233-1120123332320003-2200212232211310-0222031020030132-1010012101211200-2001032302013111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2020231123302100-0030022323031002-1231202220222300-1010323230232220-2011122331202111-2133031301110320-1003313231020203-0120222111003233"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect.retain_all_params — retain_all_params / 023203122222 / 2

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
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.redirect_route](data-sources--workload--reference--group-023.md#canonical-1132000213131301-3210120212301310-3301022312202023-2233211123100020-0010000313311330-3133001202320010-3022300313000021-0233300001331101)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect](data-sources--workload--reference--group-024.md#canonical-2323211032212033-2302322032020303-1333300100001201-2220011220303100-3333113102013320-3013212233000332-3122202001112233-2233313303121232)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect.retain_all_params

<a id="canonical-0322331323202122-0322230121002033-2120110030133131-3222033310131002-0121320200122013-3020220013221302-0232323123002330-0000013110332113"></a>

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

<a id="canonical-3303133101130001-2301203002020103-1122212130031123-1123221122003120-1232101302303030-1320123001301101-3223010311100322-2202000313232232"></a>

## Direct properties — retain_all_params / 023203122222 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1202222320333033-3233221101031210-2013210323033332-2311333123200121-0102032311210233-2202032100301001-1321212210221133-0333313130233123"></a>

## Next pages — retain_all_params / 023203122222 / 4

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect](data-sources--workload--reference--group-024.md#canonical-2323211032212033-2302322032020303-1333300100001201-2220011220303100-3333113102013320-3013212233000332-3122202001112233-2233313303121232)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-2222012210320200-3020020103010020-1310330110032302-3031321112003031-2302313332132112-0211011323231201-0210100133220232-0030303122321111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3312221131231011-3030131031312212-0200122021303131-3002200110330010-1332103110010022-2001121120211011-0131111001033133-1221121021223101"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.simple_route — simple_route / 003132003222 / 2

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
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.simple_route

<a id="canonical-0202323223310300-0333221103100201-0112100320310232-2331101012232123-3333312203220332-0230012033002321-2103101230021311-2210300200311112"></a>

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

<a id="canonical-1123321111301110-0210213022010313-3303130312032300-2300010312313111-3022301120232321-0220323122320010-0323032023233200-1323111210223002"></a>

## Direct properties — simple_route / 003132003222 / 3

- [auto_host_rewrite](data-sources--workload--reference--group-024.md#canonical-0333032102213133-0320310030322113-1001202003232233-0032210220022303-1120302022020201-0013303032223112-1110312330323301-0002323020222231): complete subsection reference.

- [disable_host_rewrite](data-sources--workload--reference--group-024.md#canonical-2231332030030201-2021100032132023-0130100013302102-3320133231001123-2313102133330122-2231222311130320-3120203331321213-0202230201113333): complete subsection reference.

<a id="canonical-0111221133323203-1311233101110302-2001121233120003-2231000300133131-0022301202032003-1202110332312301-0123200203202221-3230210220233113"></a>

<a id="canonical-2302111303110112-0230332032201110-3322121030310003-1002102100310300-2101320220123202-3203310222120032-3120112131003203-3202022132112212"></a>

## host_rewrite property — simple_route / 003132003222 / 4

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

<a id="canonical-3122203321002111-0212211022011123-3331330122222133-1021300201231023-1000331021210233-0032332303001332-1110203133213323-3023211033220200"></a>

<a id="canonical-2012002230123212-3210012002202033-3303200312323220-3212200001203231-3303110330312021-3230023102013321-3123022303300111-0133000300032020"></a>

## http_method property — simple_route / 003132003222 / 5

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

- [path](data-sources--workload--reference--group-024.md#canonical-1001123320322230-1302200132021233-1223020223023000-0111233303130210-1322003020001223-1312131031021303-3021330131031033-1231121202011131): complete subsection reference.

<a id="canonical-0000230120022221-1320202312023012-0333211223311101-1333203200322120-0123332200211330-2030100300302110-3010311122122133-2302033312301001"></a>

## Next pages — simple_route / 003132003222 / 6

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.simple_route.auto_host_rewrite](data-sources--workload--reference--group-024.md#canonical-0333032102213133-0320310030322113-1001202003232233-0032210220022303-1120302022020201-0013303032223112-1110312330323301-0002323020222231)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.simple_route.disable_host_rewrite](data-sources--workload--reference--group-024.md#canonical-2231332030030201-2021100032132023-0130100013302102-3320133231001123-2313102133330122-2231222311130320-3120203331321213-0202230201113333)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.simple_route.path](data-sources--workload--reference--group-024.md#canonical-1001123320322230-1302200132021233-1223020223023000-0111233303130210-1322003020001223-1312131031021303-3021330131031033-1231121202011131)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes](data-sources--workload--reference--group-023.md#canonical-3000130013133302-0123022001123220-3123112310130213-0322132302122231-0223033323021001-2013101011230000-2302131330020100-2000131230331230)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-0333032102213133-0320310030322113-1001202003232233-0032210220022303-1120302022020201-0013303032223112-1110312330323301-0002323020222231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1312321311332231-2131012312202230-2202113302132312-3212031202112323-3032222231003230-0202022310002122-2313131300000023-0123300033032312"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.simple_route.auto_host_rewrite — auto_host_rewrite / 003310321011 / 2

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
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.simple_route](data-sources--workload--reference--group-024.md#canonical-2222012210320200-3020020103010020-1310330110032302-3031321112003031-2302313332132112-0211011323231201-0210100133220232-0030303122321111)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.simple_route.auto_host_rewrite

<a id="canonical-3032203032102101-1012221000220332-2212000322101120-2312230320232131-2202313111200032-3310020233100332-3301121202323123-0230021012300112"></a>

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

<a id="canonical-1211103331131102-0302002123323101-3230232021031331-2303112033033121-2203310132031333-0131020120000022-0013232311332201-0103313022230232"></a>

## Direct properties — auto_host_rewrite / 003310321011 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3031301022223230-2222120232223302-3133030111310121-1003202311222122-3332122223011301-3323122020213020-3033222311011101-3232220322111310"></a>

## Next pages — auto_host_rewrite / 003310321011 / 4

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.simple_route](data-sources--workload--reference--group-024.md#canonical-2222012210320200-3020020103010020-1310330110032302-3031321112003031-2302313332132112-0211011323231201-0210100133220232-0030303122321111)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-2231332030030201-2021100032132023-0130100013302102-3320133231001123-2313102133330122-2231222311130320-3120203331321213-0202230201113333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2223023321200222-2011200230112121-1132332302300211-3212012311213113-3320312002111020-2100213120123013-3102012023222312-2023301100032223"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.simple_route.disable_host_rewrite — disable_host_rewrite / 123332120211 / 2

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
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.simple_route](data-sources--workload--reference--group-024.md#canonical-2222012210320200-3020020103010020-1310330110032302-3031321112003031-2302313332132112-0211011323231201-0210100133220232-0030303122321111)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.simple_route.disable_host_rewrite

<a id="canonical-3132122102202012-3121233132310312-0003030233301301-0001103203130210-2221221312033223-2013130000223112-2311222302300322-3033212003301023"></a>

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

<a id="canonical-0232220032323021-1302313031323320-0302300012011211-1031230012022023-1102103001303030-1002311113123021-1032223322331031-1221122330020330"></a>

## Direct properties — disable_host_rewrite / 123332120211 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1300113032312023-3112231302322023-2002323230313322-0112331311323100-1233230230021200-3200132130312013-3320001002123331-1003330011113323"></a>

## Next pages — disable_host_rewrite / 123332120211 / 4

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.simple_route](data-sources--workload--reference--group-024.md#canonical-2222012210320200-3020020103010020-1310330110032302-3031321112003031-2302313332132112-0211011323231201-0210100133220232-0030303122321111)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-1001123320322230-1302200132021233-1223020223023000-0111233303130210-1322003020001223-1312131031021303-3021330131031033-1231121202011131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0131312011130121-1130211322211231-3321212233231031-1230113210120300-1331022032121313-0230011331200101-2110213200031132-0032020122320031"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.simple_route.path — path / 131222212201 / 2

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
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.simple_route](data-sources--workload--reference--group-024.md#canonical-2222012210320200-3020020103010020-1310330110032302-3031321112003031-2302313332132112-0211011323231201-0210100133220232-0030303122321111)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.simple_route.path

<a id="canonical-2222010101101211-0003111230333232-2222010320223131-0020322013211112-1232220230312211-2210110102001030-3320113122132322-1032333233230033"></a>

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

<a id="canonical-0323330020332111-0033231010132131-2332212020210030-2003022232231301-1101132123002131-1333201313101112-0331202200233130-0103220311010102"></a>

## Direct properties — path / 131222212201 / 3

<a id="canonical-3223113313232301-0011121220220232-2100011002113232-1312122210112200-1033103002230332-3212301020311233-0013321213201202-2313111300302012"></a>

<a id="canonical-2123031230210231-1203022233101003-3230133113021032-0123223330303103-1121220100303212-0331132223222302-0233321000213110-3110313012012131"></a>

## path property — path / 131222212201 / 4

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

<a id="canonical-1223111232000213-2130100132221212-3320330223033102-1201123133310223-3002330203132002-2231133210123013-3123013310203210-1221311132123003"></a>

<a id="canonical-2212222121102331-1320320112112331-1113101003330332-3111130013322232-0310120003223322-0312123232033110-3111200232001201-1110230331211313"></a>

## prefix property — path / 131222212201 / 5

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

<a id="canonical-3001103130012121-0321222130103033-2332200021121102-3101331031130223-1321002310011112-1120321312000012-2013000010120302-3300330300112232"></a>

<a id="canonical-0232233330203132-1100230233022220-3312211010111103-2101110310121303-1210023300330213-3320301211203230-1023100120202112-0130223221213113"></a>

## regex property — path / 131222212201 / 6

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

<a id="canonical-1323022201201221-3121232112002130-0312321320203010-0233012220303123-3003310030210112-3001201211303231-3033031033322003-2103020230230013"></a>

## Next pages — path / 131222212201 / 7

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.simple_route](data-sources--workload--reference--group-024.md#canonical-2222012210320200-3020020103010020-1310330110032302-3031321112003031-2302313332132112-0211011323231201-0210100133220232-0030303122321111)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-2032122332230132-1232301210033133-2310220303021120-1333300211130330-0130020012323310-0231032130121032-0330110103233210-2320011201211330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0012220113130110-3021201213222221-2012302130331332-0020103011322101-3313112220020230-2012303201300031-1331011213132222-2232012031211010"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.port — port / 031222300012 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-017.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-017.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-020.md#canonical-1220310313301222-1022021330323111-3000020103020220-0001323023100101-1212003022023123-2103220331130211-1221230310000302-1003233331103333)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-020.md#canonical-0102110102031213-1121223312110232-2131020030001132-1300230131002230-2121033300310012-3210001131010120-0310232322332121-2033221232113110)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.port

<a id="canonical-1031321120313212-1122121100012303-1331302201120110-2000330113200131-1330111113000021-3211323010332333-3302112310301202-3033312323010222"></a>

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

<a id="canonical-1022322223302112-0021331332110000-0311332121010301-2020331033230210-2232021030112231-1210310000100301-3203211111321031-3103120101321033"></a>

## Direct properties — port / 031222300012 / 3

- [info](data-sources--workload--reference--group-024.md#canonical-2130002111313332-1030033002002100-2133210222323201-0020222000120320-1000001232131311-3123130320320312-1012011302032212-1212032030212202): complete subsection reference.

<a id="canonical-1021210021013123-3133000201020221-3020201333111110-0021023121103331-1131202320222101-3131010321130212-0220302323220302-2023121210032101"></a>

<a id="canonical-1101021001311323-3231120202133100-0221132001103111-2031023123301022-1031012212200122-2221221033312310-3210232133023301-0123023001010231"></a>

## name property — port / 031222300012 / 4

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

<a id="canonical-2321011112133203-1103021100312013-0003233313312202-2231301112333231-0033212113133102-3132110322132300-3203330021102111-0023301200031212"></a>

## Next pages — port / 031222300012 / 5

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.port.info](data-sources--workload--reference--group-024.md#canonical-2130002111313332-1030033002002100-2133210222323201-0020222000120320-1000001232131311-3123130320320312-1012011302032212-1212032030212202)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-020.md#canonical-0102110102031213-1121223312110232-2131020030001132-1300230131002230-2121033300310012-3210001131010120-0310232322332121-2033221232113110)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-2130002111313332-1030033002002100-2133210222323201-0020222000120320-1000001232131311-3123130320320312-1012011302032212-1212032030212202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0222120103121213-0111100112201012-1201113120311002-2311310302232302-2120311111233210-3310000120211312-3310010111122011-3211103303202102"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.port.info — info / 111103112302 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-017.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-017.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-020.md#canonical-1220310313301222-1022021330323111-3000020103020220-0001323023100101-1212003022023123-2103220331130211-1221230310000302-1003233331103333)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-020.md#canonical-0102110102031213-1121223312110232-2131020030001132-1300230131002230-2121033300310012-3210001131010120-0310232322332121-2033221232113110)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.port](data-sources--workload--reference--group-024.md#canonical-2032122332230132-1232301210033133-2310220303021120-1333300211130330-0130020012323310-0231032130121032-0330110103233210-2320011201211330)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.port.info

<a id="canonical-3223230223222010-1321011111013300-1113200302232113-0022303131332310-3231033332110201-0121121301230200-2130311111221131-2200130033312003"></a>

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

<a id="canonical-2311013011001222-0223300301002103-2210203031211020-3232201303310012-1022131020302010-3311100300123203-0103331030312101-2010003011111022"></a>

## Direct properties — info / 111103112302 / 3

<a id="canonical-0301100100210323-1212101113312122-3100103232010320-3211111222203220-0132211110332013-0102120021100033-3031223003221021-0310113031301210"></a>

<a id="canonical-1223210323201302-1010101021133100-2321313121300002-0013322100322013-1220322103311030-3002002121222322-2100000121330013-0230130210323232"></a>

## port property — info / 111103112302 / 4

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

<a id="canonical-0330003020313220-1121303201011021-3311211213133330-0033022230130130-1010233133303201-1300103132031033-1313210210003231-0201112001023001"></a>

<a id="canonical-2122211130212233-0310200122230133-2222210300020233-3100221013030101-0031331101123131-2103213302222022-3233122210232121-0112023021223302"></a>

## protocol property — info / 111103112302 / 5

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

- [same_as_port](data-sources--workload--reference--group-024.md#canonical-3032023012310010-2103303330221033-3231120103211013-3330101010032002-2023101123010320-1031230300100322-2303301201002232-3130313102333302): complete subsection reference.

<a id="canonical-1103323021302101-0023300330030133-2130030203311322-3112010103233120-0331033300300000-1312111123233323-0203230301133310-1303112010303211"></a>

<a id="canonical-0133230321212030-0022203210021133-2333000030102311-0112311312320123-2132112321202303-3031130030020230-3001123012313301-3313200323222301"></a>

## target_port property — info / 111103112302 / 6

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

<a id="canonical-0113111002013231-3133010101010000-0220230330101113-2202131102313122-2203100022121322-2022013100303123-3113310130002132-1222321020332020"></a>

## Next pages — info / 111103112302 / 7

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.port.info.same_as_port](data-sources--workload--reference--group-024.md#canonical-3032023012310010-2103303330221033-3231120103211013-3330101010032002-2023101123010320-1031230300100322-2303301201002232-3130313102333302)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.port](data-sources--workload--reference--group-024.md#canonical-2032122332230132-1232301210033133-2310220303021120-1333300211130330-0130020012323310-0231032130121032-0330110103233210-2320011201211330)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-3032023012310010-2103303330221033-3231120103211013-3330101010032002-2023101123010320-1031230300100322-2303301201002232-3130313102333302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0100302201323220-0012132222312031-1130020332222320-1221011311330212-2223213233331020-2232033211132231-2302303332222330-0212210010331123"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.port.info.same_as_port — same_as_port / 313000330001 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-017.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-017.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-020.md#canonical-1220310313301222-1022021330323111-3000020103020220-0001323023100101-1212003022023123-2103220331130211-1221230310000302-1003233331103333)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-020.md#canonical-0102110102031213-1121223312110232-2131020030001132-1300230131002230-2121033300310012-3210001131010120-0310232322332121-2033221232113110)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.port](data-sources--workload--reference--group-024.md#canonical-2032122332230132-1232301210033133-2310220303021120-1333300211130330-0130020012323310-0231032130121032-0330110103233210-2320011201211330)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.port.info](data-sources--workload--reference--group-024.md#canonical-2130002111313332-1030033002002100-2133210222323201-0020222000120320-1000001232131311-3123130320320312-1012011302032212-1212032030212202)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.port.info.same_as_port

<a id="canonical-0303313133301031-3302233112202202-2223232203330323-2313112112113010-1321012303211120-3220112303310121-3333101203310033-1013013233223001"></a>

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

<a id="canonical-2132321023302231-3232100032330333-2301132020020313-3033201002212010-2102331332223213-3223013212310321-1232100033002220-0031303102321323"></a>

## Direct properties — same_as_port / 313000330001 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3032200112212310-2112013000131210-0321103201233201-1333210322300233-3031232220012203-0232033022013011-3112023112310203-1210003021030213"></a>

## Next pages — same_as_port / 313000330001 / 4

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.port.info](data-sources--workload--reference--group-024.md#canonical-2130002111313332-1030033002002100-2133210222323201-0020222000120320-1000001232131311-3123130320320312-1012011302032212-1212032030212202)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-0023311210022323-0202111021131210-2023111033133223-1012103322100303-2311011122322021-0003122222031302-1322312130010112-3220023131111022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2103021220300221-1223101133213230-0102023010132012-1311301033313221-1201200213223113-0011013223301300-1123131331132330-3131331311103002"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.tcp_loadbalancer — tcp_loadbalancer / 121111232013 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-017.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-017.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-020.md#canonical-1220310313301222-1022021330323111-3000020103020220-0001323023100101-1212003022023123-2103220331130211-1221230310000302-1003233331103333)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-020.md#canonical-0102110102031213-1121223312110232-2131020030001132-1300230131002230-2121033300310012-3210001131010120-0310232322332121-2033221232113110)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.tcp_loadbalancer

<a id="canonical-2002010330122000-2112131313201311-0031103013103103-0213011302133133-1203232121030221-3023132120131113-3200331303201311-1133100211020113"></a>

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

<a id="canonical-0233312230122202-3022113311122120-0311003333331211-2102220320012321-0003312202300023-3023132133100002-2312013200011012-2130023012122220"></a>

## Direct properties — tcp_loadbalancer / 121111232013 / 3

<a id="canonical-3010011203133022-3223003223112032-1320200223303132-0002132323311012-3013102321112212-2102021233123013-0300022001131221-2131021011023301"></a>

<a id="canonical-2101023023011230-2300332331220200-1113212202211033-2323221220010102-2033113311222223-0203130212021103-1233033002101013-1113331223103213"></a>

## domains property — tcp_loadbalancer / 121111232013 / 4

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

<a id="canonical-1311132023002333-1033230231101222-2103130330130302-3212323103121222-0222210020321231-0332103033020312-3330322301310033-0002103113000131"></a>

<a id="canonical-0022313122103311-2012133322003301-1210231201320002-3031313110312033-1031121112210030-0013210020112103-2312210303221021-0322332200033223"></a>

## with_sni property — tcp_loadbalancer / 121111232013 / 5

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

<a id="canonical-0031031002000330-2201010030122301-2032130112332123-1021331330132113-0003110210322001-1311031213030220-1031133133113322-2303121000103301"></a>

## Next pages — tcp_loadbalancer / 121111232013 / 6

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-020.md#canonical-0102110102031213-1121223312110232-2131020030001132-1300230131002230-2121033300310012-3210001131010120-0310232322332121-2033221232113110)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-2011132303102222-3023301203121330-2233013320233300-0112111313030012-1103313203030222-1002221110321002-3110022320101011-2333330131330303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2121311013202331-3301001231003322-0123312200013321-1022303032212011-1100002110332322-0302302133032101-3222112300133032-1132232003320120"></a>

## stateful_service.advertise_options.advertise_on_public.port — port / 310011233011 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-017.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-017.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- stateful_service.advertise_options.advertise_on_public.port

<a id="canonical-1131003110210222-1231011122102023-1112123210202120-1130000222301032-0320102330131210-0230102223103113-3100022213310331-2303221212331302"></a>

Type: `"single"`. Computed.

Advertise Port. Advertise single port.

Upstream description:

Advertise single port.

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

<a id="canonical-0102113202100133-1211010333232300-1231101222311101-2303012320000022-2220222032222330-0033203230213302-2310202322220013-0130213302112013"></a>

## Direct properties — port / 310011233011 / 3

- [http_loadbalancer](data-sources--workload--reference--group-024.md#canonical-0122221112123120-0023010121003330-3033210010113330-0100100333002100-0211222302303110-3202021130000322-3023123330231032-2033332332032221): complete subsection reference.

- [port](data-sources--workload--reference--group-027.md#canonical-0131322030120213-1100100231233311-1130301232021101-2213213330300032-0021030111212103-0111121122012212-3320321121032303-0300331333012133): complete subsection reference.

- [tcp_loadbalancer](data-sources--workload--reference--group-027.md#canonical-1033132331311233-3020310332033313-0130113122310122-2013000233110030-2111310033030212-0200233001232302-0310022301231111-1011312300303021): complete subsection reference.

<a id="canonical-0100222032110211-0133001102020011-3310211023001113-0220211033013301-3203130330102101-1023031101332221-3313303103030021-0133130230321203"></a>

## Next pages — port / 310011233011 / 4

- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-024.md#canonical-0122221112123120-0023010121003330-3033210010113330-0100100333002100-0211222302303110-3202021130000322-3023123330231032-2033332332032221)
- [stateful_service.advertise_options.advertise_on_public.port.port](data-sources--workload--reference--group-027.md#canonical-0131322030120213-1100100231233311-1130301232021101-2213213330300032-0021030111212103-0111121122012212-3320321121032303-0300331333012133)
- [stateful_service.advertise_options.advertise_on_public.port.tcp_loadbalancer](data-sources--workload--reference--group-027.md#canonical-1033132331311233-3020310332033313-0130113122310122-2013000233110030-2111310033030212-0200233001232302-0310022301231111-1011312300303021)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-0122221112123120-0023010121003330-3033210010113330-0100100333002100-0211222302303110-3202021130000322-3023123330231032-2033332332032221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3301232200131202-3032313103112322-0010020331100011-3303133313210302-1020312220210310-3213030113010221-0332210100321322-2020213013032131"></a>

## stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer — http_loadbalancer / 210123101210 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-017.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-017.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-024.md#canonical-2011132303102222-3023301203121330-2233013320233300-0112111313030012-1103313203030222-1002221110321002-3110022320101011-2333330131330303)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer

<a id="canonical-1203112322110322-1113032100110031-2121010031113332-1020103232200122-2011133121123000-2323301211111210-2010011031220120-2002021011120020"></a>

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

<a id="canonical-2213011001023213-3111220300210230-0013033311333321-3311330102001330-0201200021323123-1203201013210210-0133232223312230-1201022211221331"></a>

## Direct properties — http_loadbalancer / 210123101210 / 3

- [default_route](data-sources--workload--reference--group-024.md#canonical-3122313211003000-3321131000121321-3032211122321213-3200023233000030-1231010213120220-3030300231023213-1000212301201002-3222210011311221): complete subsection reference.

<a id="canonical-1022322001123101-1123012121101022-1012011321120102-2022003022120312-2302031001110323-2223321302110312-0321323203302120-3330022110013332"></a>

<a id="canonical-1013130003222121-2033312100113020-1211122111213233-0313022123321032-3032210311233120-0023310030322323-0011301232230222-1321003212000331"></a>

## domains property — http_loadbalancer / 210123101210 / 4

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

- [http](data-sources--workload--reference--group-024.md#canonical-2021301101032032-3223012002101000-2200130332220120-1102231211102321-0122001310231321-1233300103101013-2002121133222203-2112020233120102): complete subsection reference.

- [https](data-sources--workload--reference--group-024.md#canonical-3130202111132300-0301212030033211-3312303220203330-1002132121130131-3322330013202010-3013320012301202-1031201103211333-3201303313311311): complete subsection reference.

- [https_auto_cert](data-sources--workload--reference--group-026.md#canonical-2202113120330223-0103121120023301-3332222222333111-2112113203120113-0331031000100122-2110333300021322-3333230230101301-0211032212120200): complete subsection reference.

- [specific_routes](data-sources--workload--reference--group-026.md#canonical-1121003100212032-3332312231032301-0223112200303302-0330012100210020-0203121122320321-2230332022321022-0312211301020331-1321232022223123): complete subsection reference.

<a id="canonical-0030003213322302-3113133310032300-2013023032002102-1233120320230313-3031021113012202-0120130003102112-0123133123010012-1333002202113212"></a>

## Next pages — http_loadbalancer / 210123101210 / 5

- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.default_route](data-sources--workload--reference--group-024.md#canonical-3122313211003000-3321131000121321-3032211122321213-3200023233000030-1231010213120220-3030300231023213-1000212301201002-3222210011311221)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.http](data-sources--workload--reference--group-024.md#canonical-2021301101032032-3223012002101000-2200130332220120-1102231211102321-0122001310231321-1233300103101013-2002121133222203-2112020233120102)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https](data-sources--workload--reference--group-024.md#canonical-3130202111132300-0301212030033211-3312303220203330-1002132121130131-3322330013202010-3013320012301202-1031201103211333-3201303313311311)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-026.md#canonical-2202113120330223-0103121120023301-3332222222333111-2112113203120113-0331031000100122-2110333300021322-3333230230101301-0211032212120200)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes](data-sources--workload--reference--group-026.md#canonical-1121003100212032-3332312231032301-0223112200303302-0330012100210020-0203121122320321-2230332022321022-0312211301020331-1321232022223123)
- [stateful_service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-024.md#canonical-2011132303102222-3023301203121330-2233013320233300-0112111313030012-1103313203030222-1002221110321002-3110022320101011-2333330131330303)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-3122313211003000-3321131000121321-3032211122321213-3200023233000030-1231010213120220-3030300231023213-1000212301201002-3222210011311221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2311221110133013-2001311013122321-1123033131032312-1013023002020320-0132121213030233-1113222221001030-2131112320133010-2233030312003301"></a>

## stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.default_route — default_route / 003201111102 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-017.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-017.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-024.md#canonical-2011132303102222-3023301203121330-2233013320233300-0112111313030012-1103313203030222-1002221110321002-3110022320101011-2333330131330303)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-024.md#canonical-0122221112123120-0023010121003330-3033210010113330-0100100333002100-0211222302303110-3202021130000322-3023123330231032-2033332332032221)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.default_route

<a id="canonical-3333222302132233-3330101121120103-1010301010301222-2131333122132022-3303210233012230-0101300203212100-1021101211220212-3130113000022320"></a>

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

<a id="canonical-3331310103030313-0331001331330020-3301103032122132-1332112220101030-0103022011130322-0331223213231220-0211130233310000-3102211031213101"></a>

## Direct properties — default_route / 003201111102 / 3

- [auto_host_rewrite](data-sources--workload--reference--group-024.md#canonical-2121111230330230-3303312200332111-2202100221013132-2300103121030221-3121121102233203-1200300303000100-2302312320113102-0032003023022200): complete subsection reference.

- [disable_host_rewrite](data-sources--workload--reference--group-024.md#canonical-2000103220303020-0220120130003100-1301002030332333-2133013330302132-2001230111002203-3320200311123302-3322232110030212-1332012300220122): complete subsection reference.

<a id="canonical-3222012202012202-3011030032201321-2130030211022002-2123002310010221-3210131221032232-1012132232011200-1203032022202003-3333102203123322"></a>

<a id="canonical-1320023221312333-1020202002102202-3032003122022103-3122120123221011-0022000112221021-1011101120003330-1003020001200233-1323231321021323"></a>

## host_rewrite property — default_route / 003201111102 / 4

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

<a id="canonical-2111023100312223-0212112211222322-0220130222212100-1033203232110132-0033022203230133-0033002301123223-3300021232112031-1322231303302032"></a>

## Next pages — default_route / 003201111102 / 5

- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.default_route.auto_host_rewrite](data-sources--workload--reference--group-024.md#canonical-2121111230330230-3303312200332111-2202100221013132-2300103121030221-3121121102233203-1200300303000100-2302312320113102-0032003023022200)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.default_route.disable_host_rewrite](data-sources--workload--reference--group-024.md#canonical-2000103220303020-0220120130003100-1301002030332333-2133013330302132-2001230111002203-3320200311123302-3322232110030212-1332012300220122)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-024.md#canonical-0122221112123120-0023010121003330-3033210010113330-0100100333002100-0211222302303110-3202021130000322-3023123330231032-2033332332032221)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-2121111230330230-3303312200332111-2202100221013132-2300103121030221-3121121102233203-1200300303000100-2302312320113102-0032003023022200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0001203002012023-3312200022123223-1122121101300322-1313213301000021-2201131010301022-1203000320030003-1121302300123330-3210221012022202"></a>

## stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.default_route.auto_host_rewrite — auto_host_rewrite / 120123132032 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-017.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-017.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-024.md#canonical-2011132303102222-3023301203121330-2233013320233300-0112111313030012-1103313203030222-1002221110321002-3110022320101011-2333330131330303)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-024.md#canonical-0122221112123120-0023010121003330-3033210010113330-0100100333002100-0211222302303110-3202021130000322-3023123330231032-2033332332032221)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.default_route](data-sources--workload--reference--group-024.md#canonical-3122313211003000-3321131000121321-3032211122321213-3200023233000030-1231010213120220-3030300231023213-1000212301201002-3222210011311221)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.default_route.auto_host_rewrite

<a id="canonical-1320322000113110-0000200221230302-3112033213132310-1232101303002102-0320233010220123-1223010113200100-0321033111302101-0031301321100302"></a>

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

<a id="canonical-1001313213221330-1033313132220313-0133330201313212-1330302130203212-0312300310200300-0332331320203022-1111023302021331-0012223032010210"></a>

## Direct properties — auto_host_rewrite / 120123132032 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3013201210112231-0103202232111220-3031200333200021-3030210233103022-3311021031300032-2032003333023112-1231100120103231-0233300101011320"></a>

## Next pages — auto_host_rewrite / 120123132032 / 4

- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.default_route](data-sources--workload--reference--group-024.md#canonical-3122313211003000-3321131000121321-3032211122321213-3200023233000030-1231010213120220-3030300231023213-1000212301201002-3222210011311221)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-2000103220303020-0220120130003100-1301002030332333-2133013330302132-2001230111002203-3320200311123302-3322232110030212-1332012300220122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0201020133001121-2221323012010322-0100122002021022-3102303221100311-2310022021301123-1210131120222301-0303210320021112-1332013101001101"></a>

## stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.default_route.disable_host_rewrite — disable_host_rewrite / 103310302130 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-017.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-017.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-024.md#canonical-2011132303102222-3023301203121330-2233013320233300-0112111313030012-1103313203030222-1002221110321002-3110022320101011-2333330131330303)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-024.md#canonical-0122221112123120-0023010121003330-3033210010113330-0100100333002100-0211222302303110-3202021130000322-3023123330231032-2033332332032221)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.default_route](data-sources--workload--reference--group-024.md#canonical-3122313211003000-3321131000121321-3032211122321213-3200023233000030-1231010213120220-3030300231023213-1000212301201002-3222210011311221)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.default_route.disable_host_rewrite

<a id="canonical-0231031313000113-0003120310220323-0032330301031303-2032313221231023-2321330123202013-3333031202112103-1200300302113020-2111112112002010"></a>

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

<a id="canonical-1003312231111220-0001231231012323-1303012032131021-3121211221000312-3120101021232303-2121330113313202-0132022111131232-1102033100021201"></a>

## Direct properties — disable_host_rewrite / 103310302130 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1302013312010021-1000030130333211-0113320110201200-3122103210203032-0002210113312010-0011001131022310-1201020310110210-1103023232033103"></a>

## Next pages — disable_host_rewrite / 103310302130 / 4

- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.default_route](data-sources--workload--reference--group-024.md#canonical-3122313211003000-3321131000121321-3032211122321213-3200023233000030-1231010213120220-3030300231023213-1000212301201002-3222210011311221)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-2021301101032032-3223012002101000-2200130332220120-1102231211102321-0122001310231321-1233300103101013-2002121133222203-2112020233120102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2211301130333212-0211200031320300-2120031111102230-2300212001111330-2332212033321000-0010322111011023-1031021333110331-0033311301213232"></a>

## stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.http — http / 133112212200 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-017.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-017.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-024.md#canonical-2011132303102222-3023301203121330-2233013320233300-0112111313030012-1103313203030222-1002221110321002-3110022320101011-2333330131330303)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-024.md#canonical-0122221112123120-0023010121003330-3033210010113330-0100100333002100-0211222302303110-3202021130000322-3023123330231032-2033332332032221)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.http

<a id="canonical-3120331311332000-0312121223231213-1030232011200311-2122032322321022-1121313100020001-3202130202200211-0001233203002100-1332212122130120"></a>

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

<a id="canonical-1110213131100131-1321113230021112-0303022031100013-1103111022133331-3003222013223103-1133120211000132-3003001322132110-0232111331201102"></a>

## Direct properties — http / 133112212200 / 3

<a id="canonical-1323100231220232-1212201330313110-2013231222201133-1213122100101000-3013321100000310-1103231013001322-2023203311021330-2302231333012310"></a>

<a id="canonical-1213201130212011-0322231120323011-0131330202202310-0011313133311320-3023101323301230-1203232223013110-2020132031022231-3323300332021202"></a>

## dns_volterra_managed property — http / 133112212200 / 4

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

<a id="canonical-2000302011132313-0113102201120030-3132000000212123-1120221220200101-1302231222000112-0222101103102100-3120322320031232-1022203320010130"></a>

<a id="canonical-3221130232323011-3103202123101132-3311302232120321-0210203133202313-2202231300301223-1022000013003102-3023321321223233-2310031333220130"></a>

## port property — http / 133112212200 / 5

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

<a id="canonical-0112113301112323-1320332103300232-2101211011122333-3312023013121322-0003231211320321-0201312232222311-2311310302331321-0013130220203100"></a>

<a id="canonical-0001302021203031-2111110213230000-1131033220330100-2222031001330102-2300131110321302-2131102312222322-1332231102300100-3002123110330112"></a>

## port_ranges property — http / 133112212200 / 6

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

<a id="canonical-1221310202110230-1102320201121211-0102023331100231-0210021032032313-0100020003222322-3102212323022220-2203303211210002-1221132012020313"></a>

## Next pages — http / 133112212200 / 7

- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-024.md#canonical-0122221112123120-0023010121003330-3033210010113330-0100100333002100-0211222302303110-3202021130000322-3023123330231032-2033332332032221)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-3130202111132300-0301212030033211-3312303220203330-1002132121130131-3322330013202010-3013320012301202-1031201103211333-3201303313311311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2313211213012313-1202313331322231-1123331123301112-0101102221222322-3102203211133223-0221103131213110-2331023113133201-2302100032312223"></a>

## stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https — https / 110300132113 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-017.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-017.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-024.md#canonical-2011132303102222-3023301203121330-2233013320233300-0112111313030012-1103313203030222-1002221110321002-3110022320101011-2333330131330303)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-024.md#canonical-0122221112123120-0023010121003330-3033210010113330-0100100333002100-0211222302303110-3202021130000322-3023123330231032-2033332332032221)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https

<a id="canonical-2010312133310032-2331131323333120-2211313013332301-3020322223221132-2013010202023322-2001020110311313-0322323101331013-0330111000023110"></a>

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

<a id="canonical-1210003223333032-0220130222013103-3321110131112032-2220313031221210-3102320123333300-3001220103001011-1212202013123101-0112121112212211"></a>

## Direct properties — https / 110300132113 / 3

<a id="canonical-0330303312112230-0313123332131022-0022203130122113-0313122131323111-1013312212020111-0311210210130000-1323232130213213-0212023123130122"></a>

<a id="canonical-2130202203223333-2131031010003303-3331230123310012-2322003010333113-2330132001232321-1013003232000101-0013211301221321-3200101010002130"></a>

## add_hsts property — https / 110300132113 / 4

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

<a id="canonical-2033023131021203-0022300232032010-2332032203330100-1131222113122130-2030010232210000-3033003230322131-2303233200133211-0132313131203230"></a>

<a id="canonical-1320103103302020-0012033210121121-3233203112011133-1112132131111230-0303102013320122-1102202333031002-1033110212122313-2032222230013132"></a>

## append_server_name property — https / 110300132113 / 5

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

- [coalescing_options](data-sources--workload--reference--group-024.md#canonical-1023120122302200-2011233320230320-1102233120232223-2013001132122323-1223332131001032-3021210021131100-2203100132021113-3312331033302103): complete subsection reference.

<a id="canonical-1111310220013220-1303301322311130-3021223223302121-0222001102133212-3013223023302102-3300010320120223-2122312310013301-2130111333202121"></a>

<a id="canonical-2132300000322302-3010220312000101-3211212130333201-1332021332123321-3123013323331222-3320302001202030-1312332213122112-1131131321133302"></a>

## connection_idle_timeout property — https / 110300132113 / 6

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
    "ves.io.schema.rules.uint32.lte": "600000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "600000"
  }
}
```

- [default_header](data-sources--workload--reference--group-024.md#canonical-3212030333331111-1323213222123123-3113032102033012-1200212310332211-0113033032320311-1220000132023221-2113121320322131-2232332232333031): complete subsection reference.

- [default_loadbalancer](data-sources--workload--reference--group-024.md#canonical-0223223101131313-1300310221112102-0321210132001022-3112212300000023-0132301001002132-3302030003012101-3320110120100223-1110101022012123): complete subsection reference.

- [disable_path_normalize](data-sources--workload--reference--group-024.md#canonical-2333223311220130-2111013331020103-1221120102101213-3220002012231222-2021122321032332-1333013010022023-0131031320013100-0120221203102300): complete subsection reference.

- [enable_path_normalize](data-sources--workload--reference--group-024.md#canonical-0223132321010103-3022011311232223-3212232213012101-1320003202112021-3211202013132331-3121121312002301-1330203100230032-2221331332130333): complete subsection reference.

- [http_protocol_options](data-sources--workload--reference--group-024.md#canonical-3333101033030012-1032120202302303-2102001301123323-2131030113312320-3213223211120231-2113300031233322-0220322000113332-0333121113331310): complete subsection reference.

<a id="canonical-3203120123030000-1111011332333113-2103122200310022-1032121012001012-2233201203001023-2323112102332302-2200033331121033-0033312303310000"></a>

<a id="canonical-0330030302310032-0330023110211221-3020212230211200-2323210221211122-0201330332312101-0013003020203203-2010003323032022-2200100133212002"></a>

## http_redirect property — https / 110300132113 / 7

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

- [non_default_loadbalancer](data-sources--workload--reference--group-025.md#canonical-0301011213122133-1300033113213030-3120002233210213-1112230010011111-0020221320132203-0321110303123233-2332103123233312-1221221101223010): complete subsection reference.

- [pass_through](data-sources--workload--reference--group-025.md#canonical-0122012011231231-1200131133311312-1030012131023030-2131222022303011-3123012333133233-0323201300031012-2231023102333330-1101231222212323): complete subsection reference.

<a id="canonical-3320320211321320-0032032320320130-3102122222122121-1212213132032223-2302212031302323-0021311131030010-2212233221310003-3300301121110021"></a>

<a id="canonical-3321223303203113-1132001131331113-3130021210120331-1001103220201330-1200300212122032-3033122101021221-1213222011122320-2110301203010201"></a>

## port property — https / 110300132113 / 8

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

<a id="canonical-1111021121311212-0223333312220000-2101030103212022-3133000132200023-1131011123233010-2100120233113222-1112323012001212-2330230130103011"></a>

<a id="canonical-2023000201002200-0211121120330212-3120122312322222-3323322321111012-2111022031013312-0310033212012330-2301133132130133-0031332223101310"></a>

## port_ranges property — https / 110300132113 / 9

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

<a id="canonical-3001220033131230-3332300321110321-1303022313133212-0123333000301030-1332112112132103-0030120021321310-0221231132133100-1012232133111012"></a>

<a id="canonical-2231321313212021-2010222303201311-0030230302111021-2321232113330200-0331230212103312-2032123222213212-1131021322033020-2220131102110311"></a>

## server_name property — https / 110300132113 / 10

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

- [tls_cert_params](data-sources--workload--reference--group-025.md#canonical-3322122302333010-1112232022110101-3332210121013303-0203013103132310-3121123232001030-0330332323311120-3221321223111332-3210202001303013): complete subsection reference.

- [tls_parameters](data-sources--workload--reference--group-025.md#canonical-1213321303121301-1102133032030121-0221111123231231-2130000023010203-0303210000223333-1100020323111330-0120221132032311-1020313231211133): complete subsection reference.

<a id="canonical-0320203320111322-3223131131123033-3230101100112323-3000012311010130-1311213203133211-0010013111002221-2321212310123303-0331221111033013"></a>

## Next pages — https / 110300132113 / 11

- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.coalescing_options](data-sources--workload--reference--group-024.md#canonical-1023120122302200-2011233320230320-1102233120232223-2013001132122323-1223332131001032-3021210021131100-2203100132021113-3312331033302103)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.default_header](data-sources--workload--reference--group-024.md#canonical-3212030333331111-1323213222123123-3113032102033012-1200212310332211-0113033032320311-1220000132023221-2113121320322131-2232332232333031)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.default_loadbalancer](data-sources--workload--reference--group-024.md#canonical-0223223101131313-1300310221112102-0321210132001022-3112212300000023-0132301001002132-3302030003012101-3320110120100223-1110101022012123)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.disable_path_normalize](data-sources--workload--reference--group-024.md#canonical-2333223311220130-2111013331020103-1221120102101213-3220002012231222-2021122321032332-1333013010022023-0131031320013100-0120221203102300)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.enable_path_normalize](data-sources--workload--reference--group-024.md#canonical-0223132321010103-3022011311232223-3212232213012101-1320003202112021-3211202013132331-3121121312002301-1330203100230032-2221331332130333)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options](data-sources--workload--reference--group-024.md#canonical-3333101033030012-1032120202302303-2102001301123323-2131030113312320-3213223211120231-2113300031233322-0220322000113332-0333121113331310)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.non_default_loadbalancer](data-sources--workload--reference--group-025.md#canonical-0301011213122133-1300033113213030-3120002233210213-1112230010011111-0020221320132203-0321110303123233-2332103123233312-1221221101223010)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.pass_through](data-sources--workload--reference--group-025.md#canonical-0122012011231231-1200131133311312-1030012131023030-2131222022303011-3123012333133233-0323201300031012-2231023102333330-1101231222212323)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params](data-sources--workload--reference--group-025.md#canonical-3322122302333010-1112232022110101-3332210121013303-0203013103132310-3121123232001030-0330332323311120-3221321223111332-3210202001303013)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters](data-sources--workload--reference--group-025.md#canonical-1213321303121301-1102133032030121-0221111123231231-2130000023010203-0303210000223333-1100020323111330-0120221132032311-1020313231211133)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-024.md#canonical-0122221112123120-0023010121003330-3033210010113330-0100100333002100-0211222302303110-3202021130000322-3023123330231032-2033332332032221)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-1023120122302200-2011233320230320-1102233120232223-2013001132122323-1223332131001032-3021210021131100-2203100132021113-3312331033302103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2310220020013310-3230302303221223-2233121331033021-0201001102201213-3223011312300033-0222010012230121-0203022023031110-3031010213122131"></a>

## stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.coalescing_options — coalescing_options / 313102222302 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-017.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-017.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-024.md#canonical-2011132303102222-3023301203121330-2233013320233300-0112111313030012-1103313203030222-1002221110321002-3110022320101011-2333330131330303)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-024.md#canonical-0122221112123120-0023010121003330-3033210010113330-0100100333002100-0211222302303110-3202021130000322-3023123330231032-2033332332032221)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https](data-sources--workload--reference--group-024.md#canonical-3130202111132300-0301212030033211-3312303220203330-1002132121130131-3322330013202010-3013320012301202-1031201103211333-3201303313311311)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.coalescing_options

<a id="canonical-2203323301213223-3231023200122010-0203013223011010-1203101230213020-1332312123201321-3022301311200330-0110002301123311-2130222212110020"></a>

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

<a id="canonical-3032032303033221-3110012220012331-0302121201332232-0321010000220102-2333321112211312-3311132030222013-1223322210132123-0311312330323002"></a>

## Direct properties — coalescing_options / 313102222302 / 3

- [default_coalescing](data-sources--workload--reference--group-024.md#canonical-2103211223112000-1203311230111230-1221320100113101-0132330220101102-0013033232113021-3310121201133210-2031320011000311-3010132331303100): complete subsection reference.

- [strict_coalescing](data-sources--workload--reference--group-024.md#canonical-3011320110330222-2030112311201003-2112133010201030-0222301212130320-0103330022012101-3332200131230123-0322122330323231-0200323013130132): complete subsection reference.

<a id="canonical-3123331322320310-0210201321133021-2130110101111033-0033230132231133-3132012122121020-1322330003011020-3222213011303133-2021013222033310"></a>

## Next pages — coalescing_options / 313102222302 / 4

- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.coalescing_options.default_coalescing](data-sources--workload--reference--group-024.md#canonical-2103211223112000-1203311230111230-1221320100113101-0132330220101102-0013033232113021-3310121201133210-2031320011000311-3010132331303100)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.coalescing_options.strict_coalescing](data-sources--workload--reference--group-024.md#canonical-3011320110330222-2030112311201003-2112133010201030-0222301212130320-0103330022012101-3332200131230123-0322122330323231-0200323013130132)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https](data-sources--workload--reference--group-024.md#canonical-3130202111132300-0301212030033211-3312303220203330-1002132121130131-3322330013202010-3013320012301202-1031201103211333-3201303313311311)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-2103211223112000-1203311230111230-1221320100113101-0132330220101102-0013033232113021-3310121201133210-2031320011000311-3010132331303100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2233102103003201-2021100133130130-3231013033113021-1021001003223111-3323302323201323-2120320130233332-3232123222033201-1302321133131023"></a>

## stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.coalescing_options.default_coalescing — default_coalescing / 222201112103 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-017.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-017.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-024.md#canonical-2011132303102222-3023301203121330-2233013320233300-0112111313030012-1103313203030222-1002221110321002-3110022320101011-2333330131330303)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-024.md#canonical-0122221112123120-0023010121003330-3033210010113330-0100100333002100-0211222302303110-3202021130000322-3023123330231032-2033332332032221)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https](data-sources--workload--reference--group-024.md#canonical-3130202111132300-0301212030033211-3312303220203330-1002132121130131-3322330013202010-3013320012301202-1031201103211333-3201303313311311)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.coalescing_options](data-sources--workload--reference--group-024.md#canonical-1023120122302200-2011233320230320-1102233120232223-2013001132122323-1223332131001032-3021210021131100-2203100132021113-3312331033302103)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.coalescing_options.default_coalescing

<a id="canonical-0332300103311012-3001223203123001-3003031013333003-2303031100123233-2031323200302030-2211123300320102-1213123310222031-3321022323003211"></a>

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

<a id="canonical-3030200330002013-2312120230322313-0022202123220332-2333211012132133-2112133032003303-0033010210320201-2332300323230202-3213233313112232"></a>

## Direct properties — default_coalescing / 222201112103 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3210023212202132-3020300210122032-0020131331331300-1302112210312030-3200032003302303-1031303310030333-2232132033003110-1123130211213132"></a>

## Next pages — default_coalescing / 222201112103 / 4

- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.coalescing_options](data-sources--workload--reference--group-024.md#canonical-1023120122302200-2011233320230320-1102233120232223-2013001132122323-1223332131001032-3021210021131100-2203100132021113-3312331033302103)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-3011320110330222-2030112311201003-2112133010201030-0222301212130320-0103330022012101-3332200131230123-0322122330323231-0200323013130132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2212121321030030-1233112313002122-0233123011112302-0213300322221220-2233002011310022-2100233201133201-0321001221203010-2101032003030102"></a>

## stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.coalescing_options.strict_coalescing — strict_coalescing / 031133011011 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-017.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-017.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-024.md#canonical-2011132303102222-3023301203121330-2233013320233300-0112111313030012-1103313203030222-1002221110321002-3110022320101011-2333330131330303)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-024.md#canonical-0122221112123120-0023010121003330-3033210010113330-0100100333002100-0211222302303110-3202021130000322-3023123330231032-2033332332032221)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https](data-sources--workload--reference--group-024.md#canonical-3130202111132300-0301212030033211-3312303220203330-1002132121130131-3322330013202010-3013320012301202-1031201103211333-3201303313311311)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.coalescing_options](data-sources--workload--reference--group-024.md#canonical-1023120122302200-2011233320230320-1102233120232223-2013001132122323-1223332131001032-3021210021131100-2203100132021113-3312331033302103)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.coalescing_options.strict_coalescing

<a id="canonical-1303221103313030-2330002231312200-0200021333101231-0220111122021313-3202231212110113-0230323233133103-1110133320232000-1031003221313211"></a>

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

<a id="canonical-2132303101021103-0301310023302221-0210031210013232-3102130000310010-2002103002323213-1002213021330220-1230201011110010-1311013203302012"></a>

## Direct properties — strict_coalescing / 031133011011 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3333203202110320-2300213011132120-3112111313110230-1231221133300212-0311203211130312-1023111031121213-3133302131003003-3221013010031003"></a>

## Next pages — strict_coalescing / 031133011011 / 4

- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.coalescing_options](data-sources--workload--reference--group-024.md#canonical-1023120122302200-2011233320230320-1102233120232223-2013001132122323-1223332131001032-3021210021131100-2203100132021113-3312331033302103)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-3212030333331111-1323213222123123-3113032102033012-1200212310332211-0113033032320311-1220000132023221-2113121320322131-2232332232333031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0003222010222013-1030213221020201-3013003321113132-0212000322300121-0023321130220211-0110121010131331-2103312121220231-2033111002011113"></a>

## stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.default_header — default_header / 222120013103 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-017.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-017.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-024.md#canonical-2011132303102222-3023301203121330-2233013320233300-0112111313030012-1103313203030222-1002221110321002-3110022320101011-2333330131330303)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-024.md#canonical-0122221112123120-0023010121003330-3033210010113330-0100100333002100-0211222302303110-3202021130000322-3023123330231032-2033332332032221)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https](data-sources--workload--reference--group-024.md#canonical-3130202111132300-0301212030033211-3312303220203330-1002132121130131-3322330013202010-3013320012301202-1031201103211333-3201303313311311)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.default_header

<a id="canonical-2120032322102101-0131120102310322-0021032210033303-2111303001201312-1113303100023120-3112212213123231-2100022011112333-0001320000222101"></a>

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

<a id="canonical-0233210201121212-2110020231122101-1220023100331313-2003233200230301-2231001121221331-0311301310021101-2133111120123222-3221232330232322"></a>

## Direct properties — default_header / 222120013103 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0023303323122011-0220211302032213-1332231013033231-1131203222311321-0221022023302321-1100322201011113-2130113213223333-0223322133021130"></a>

## Next pages — default_header / 222120013103 / 4

- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https](data-sources--workload--reference--group-024.md#canonical-3130202111132300-0301212030033211-3312303220203330-1002132121130131-3322330013202010-3013320012301202-1031201103211333-3201303313311311)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-0223223101131313-1300310221112102-0321210132001022-3112212300000023-0132301001002132-3302030003012101-3320110120100223-1110101022012123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1030331200132121-2020131103021100-3001022113002131-3300133032021321-3023323000030210-1322232012013031-3001103020031112-2032312123203123"></a>

## stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.default_loadbalancer — default_loadbalancer / 000321102010 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-017.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-017.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-024.md#canonical-2011132303102222-3023301203121330-2233013320233300-0112111313030012-1103313203030222-1002221110321002-3110022320101011-2333330131330303)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-024.md#canonical-0122221112123120-0023010121003330-3033210010113330-0100100333002100-0211222302303110-3202021130000322-3023123330231032-2033332332032221)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https](data-sources--workload--reference--group-024.md#canonical-3130202111132300-0301212030033211-3312303220203330-1002132121130131-3322330013202010-3013320012301202-1031201103211333-3201303313311311)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.default_loadbalancer

<a id="canonical-3011333321320210-0020320200302030-3100221212332001-1101133220001133-1220132321212222-1321121022331123-1101333022032300-2130233233022221"></a>

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

<a id="canonical-1210330320302121-3322013220333101-0011322022322031-3030112312331000-1310323322101322-0301010312011113-3130020200200132-2230022012132011"></a>

## Direct properties — default_loadbalancer / 000321102010 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0221221332332022-1332330201310211-3333211023312203-1031220223110223-3202030001322210-1030112133001113-1331310133023231-2323211313303001"></a>

## Next pages — default_loadbalancer / 000321102010 / 4

- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https](data-sources--workload--reference--group-024.md#canonical-3130202111132300-0301212030033211-3312303220203330-1002132121130131-3322330013202010-3013320012301202-1031201103211333-3201303313311311)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-2333223311220130-2111013331020103-1221120102101213-3220002012231222-2021122321032332-1333013010022023-0131031320013100-0120221203102300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3010110302003302-0002330030310000-1311023113120003-1123001302013322-1032323030210113-2000301013100103-0100121113211010-2123110303303000"></a>

## stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.disable_path_normalize — disable_path_normalize / 122213010203 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-017.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-017.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-024.md#canonical-2011132303102222-3023301203121330-2233013320233300-0112111313030012-1103313203030222-1002221110321002-3110022320101011-2333330131330303)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-024.md#canonical-0122221112123120-0023010121003330-3033210010113330-0100100333002100-0211222302303110-3202021130000322-3023123330231032-2033332332032221)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https](data-sources--workload--reference--group-024.md#canonical-3130202111132300-0301212030033211-3312303220203330-1002132121130131-3322330013202010-3013320012301202-1031201103211333-3201303313311311)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.disable_path_normalize

<a id="canonical-0012133022203313-0002301230330313-2130112010230333-3230132023203113-3223333232030302-3223100123322230-2012321301201030-0331310110231213"></a>

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

<a id="canonical-2112111223223302-1212110000022303-2133223013220203-3003131013233300-1200001030031231-0300110111113111-2002121210023023-0122311312201221"></a>

## Direct properties — disable_path_normalize / 122213010203 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1030020222121311-2133210001210313-0001231302022101-1332300333113321-0312312312322032-2310232333110013-3312101232330221-2223111000133332"></a>

## Next pages — disable_path_normalize / 122213010203 / 4

- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https](data-sources--workload--reference--group-024.md#canonical-3130202111132300-0301212030033211-3312303220203330-1002132121130131-3322330013202010-3013320012301202-1031201103211333-3201303313311311)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-0223132321010103-3022011311232223-3212232213012101-1320003202112021-3211202013132331-3121121312002301-1330203100230032-2221331332130333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3223331222223100-1203322010231312-0122300330330303-3020000001302001-1002203111230010-3113332313200011-2233001122331020-0131113301102301"></a>

## stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.enable_path_normalize — enable_path_normalize / 000110210320 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-017.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-017.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-024.md#canonical-2011132303102222-3023301203121330-2233013320233300-0112111313030012-1103313203030222-1002221110321002-3110022320101011-2333330131330303)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-024.md#canonical-0122221112123120-0023010121003330-3033210010113330-0100100333002100-0211222302303110-3202021130000322-3023123330231032-2033332332032221)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https](data-sources--workload--reference--group-024.md#canonical-3130202111132300-0301212030033211-3312303220203330-1002132121130131-3322330013202010-3013320012301202-1031201103211333-3201303313311311)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.enable_path_normalize

<a id="canonical-2130223123320021-0033132311231022-3220133232103221-1312301000020130-2321103113221103-0201101131132003-2311202313003133-1031213122032010"></a>

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

<a id="canonical-0231020312332030-3000101311231121-1220011311221220-1203332110311233-1010011121331311-3121310330320103-3111010112001323-1232100230311200"></a>

## Direct properties — enable_path_normalize / 000110210320 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0322011001103111-1202122232201220-3001003330001122-2213221122112233-3121311033231303-1313120330222213-3321102220203302-2120321223100133"></a>

## Next pages — enable_path_normalize / 000110210320 / 4

- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https](data-sources--workload--reference--group-024.md#canonical-3130202111132300-0301212030033211-3312303220203330-1002132121130131-3322330013202010-3013320012301202-1031201103211333-3201303313311311)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-3333101033030012-1032120202302303-2102001301123323-2131030113312320-3213223211120231-2113300031233322-0220322000113332-0333121113331310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2211200222203012-2320201210100332-3101231112022310-2113002301301222-2323201333223130-0200311033123101-2301120313322301-2032212021310133"></a>

## stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options — http_protocol_options / 313233000011 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-017.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-017.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-024.md#canonical-2011132303102222-3023301203121330-2233013320233300-0112111313030012-1103313203030222-1002221110321002-3110022320101011-2333330131330303)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-024.md#canonical-0122221112123120-0023010121003330-3033210010113330-0100100333002100-0211222302303110-3202021130000322-3023123330231032-2033332332032221)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https](data-sources--workload--reference--group-024.md#canonical-3130202111132300-0301212030033211-3312303220203330-1002132121130131-3322330013202010-3013320012301202-1031201103211333-3201303313311311)
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

<a id="canonical-0123222323122231-2212023103003032-0223233233321330-3132212030203033-2203011031303331-2213311233132123-1203321001131113-0023321111200300"></a>

## Direct properties — http_protocol_options / 313233000011 / 3

- [http_protocol_enable_v1_only](data-sources--workload--reference--group-024.md#canonical-3300320011203011-1210133323120233-2131113302030100-1230323010213120-0203303020111003-2013130012011210-2102210122232110-2112230323021100): complete subsection reference.

- [http_protocol_enable_v1_v2](data-sources--workload--reference--group-025.md#canonical-0220131310311122-3120003032210212-2001333311320302-3123100130023201-1020032013313313-1232310322030222-3110103122332311-0310201020323021): complete subsection reference.

- [http_protocol_enable_v2_only](data-sources--workload--reference--group-025.md#canonical-1230003231101131-0102300131032130-2022231311330100-1021022132012311-1121102112112022-0002330130333303-3011301112221000-1002232020213000): complete subsection reference.

<a id="canonical-3013331022031031-3302213023322031-3101313010302003-3010331301013102-3200120312203223-0103220230313210-2213100222200102-0200220012200112"></a>

## Next pages — http_protocol_options / 313233000011 / 4

- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only](data-sources--workload--reference--group-024.md#canonical-3300320011203011-1210133323120233-2131113302030100-1230323010213120-0203303020111003-2013130012011210-2102210122232110-2112230323021100)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_v2](data-sources--workload--reference--group-025.md#canonical-0220131310311122-3120003032210212-2001333311320302-3123100130023201-1020032013313313-1232310322030222-3110103122332311-0310201020323021)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v2_only](data-sources--workload--reference--group-025.md#canonical-1230003231101131-0102300131032130-2022231311330100-1021022132012311-1121102112112022-0002330130333303-3011301112221000-1002232020213000)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https](data-sources--workload--reference--group-024.md#canonical-3130202111132300-0301212030033211-3312303220203330-1002132121130131-3322330013202010-3013320012301202-1031201103211333-3201303313311311)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-3300320011203011-1210133323120233-2131113302030100-1230323010213120-0203303020111003-2013130012011210-2102210122232110-2112230323021100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0131212012022320-0003021003001112-3211121030232223-1300120112111323-1033330030011121-1323010222300310-2122320330001211-3313312310233033"></a>

## stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only — http_protocol_enable_v1_only / 101022211210 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-017.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-017.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-024.md#canonical-2011132303102222-3023301203121330-2233013320233300-0112111313030012-1103313203030222-1002221110321002-3110022320101011-2333330131330303)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-024.md#canonical-0122221112123120-0023010121003330-3033210010113330-0100100333002100-0211222302303110-3202021130000322-3023123330231032-2033332332032221)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https](data-sources--workload--reference--group-024.md#canonical-3130202111132300-0301212030033211-3312303220203330-1002132121130131-3322330013202010-3013320012301202-1031201103211333-3201303313311311)
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

<a id="canonical-2313032102202113-1301313323313003-2002112111310113-3233131100122000-0133130012112321-1011223323033103-0313332031033020-3002123332023031"></a>

## Direct properties — http_protocol_enable_v1_only / 101022211210 / 3

- [header_transformation](data-sources--workload--reference--group-024.md#canonical-3122222212030101-1203221120031210-3333023032021123-0223230011222133-3233112301020012-2011220311220011-1122123312122111-0133320312023110): complete subsection reference.

<a id="canonical-2133213313130112-1100233022301232-1123033200302023-0213213333221231-3132322233223303-0233120023123020-2203002302033013-3301001311320330"></a>

## Next pages — http_protocol_enable_v1_only / 101022211210 / 4

- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation](data-sources--workload--reference--group-024.md#canonical-3122222212030101-1203221120031210-3333023032021123-0223230011222133-3233112301020012-2011220311220011-1122123312122111-0133320312023110)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options](data-sources--workload--reference--group-024.md#canonical-3333101033030012-1032120202302303-2102001301123323-2131030113312320-3213223211120231-2113300031233322-0220322000113332-0333121113331310)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-3122222212030101-1203221120031210-3333023032021123-0223230011222133-3233112301020012-2011220311220011-1122123312122111-0133320312023110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0030121333133203-0132132110022221-0333100313232013-2101021221031110-3001011112102222-2121322001013300-2010100003302101-1122233220310313"></a>

## stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation — header_transformation / 101320012120 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-017.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-017.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-024.md#canonical-2011132303102222-3023301203121330-2233013320233300-0112111313030012-1103313203030222-1002221110321002-3110022320101011-2333330131330303)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-024.md#canonical-0122221112123120-0023010121003330-3033210010113330-0100100333002100-0211222302303110-3202021130000322-3023123330231032-2033332332032221)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https](data-sources--workload--reference--group-024.md#canonical-3130202111132300-0301212030033211-3312303220203330-1002132121130131-3322330013202010-3013320012301202-1031201103211333-3201303313311311)
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

<a id="canonical-1001110202312112-2212211322003211-1331310232222221-1313311203113303-2132222110013023-3123312102211332-2031301323333111-1100322210323131"></a>

## Direct properties — header_transformation / 101320012120 / 3

- [default_header_transformation](data-sources--workload--reference--group-024.md#canonical-0031323121102300-0333321012113223-0011020330023110-0023320100323302-2233001223221013-1112130133013220-3120222213021313-0213203101321111): complete subsection reference.

- [preserve_case_header_transformation](data-sources--workload--reference--group-024.md#canonical-2101222223223301-2010302221021123-3103213310223003-3221232220011220-0012133332002131-0121301010130302-0323032003213322-0231033023001123): complete subsection reference.

- [proper_case_header_transformation](data-sources--workload--reference--group-024.md#canonical-1022122133111101-1332003303032202-1301322031220110-2332232111311102-3200000213102003-3031002202333202-2133022321322313-2003323102002120): complete subsection reference.

<a id="canonical-1131100002330321-1230220310311012-2111001322032113-1313230110011310-0231113310211200-0302222330320023-3310230212022102-2323110201231013"></a>

## Next pages — header_transformation / 101320012120 / 4

- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.default_header_transformation](data-sources--workload--reference--group-024.md#canonical-0031323121102300-0333321012113223-0011020330023110-0023320100323302-2233001223221013-1112130133013220-3120222213021313-0213203101321111)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.preserve_case_header_transformation](data-sources--workload--reference--group-024.md#canonical-2101222223223301-2010302221021123-3103213310223003-3221232220011220-0012133332002131-0121301010130302-0323032003213322-0231033023001123)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.proper_case_header_transformation](data-sources--workload--reference--group-024.md#canonical-1022122133111101-1332003303032202-1301322031220110-2332232111311102-3200000213102003-3031002202333202-2133022321322313-2003323102002120)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only](data-sources--workload--reference--group-024.md#canonical-3300320011203011-1210133323120233-2131113302030100-1230323010213120-0203303020111003-2013130012011210-2102210122232110-2112230323021100)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-0031323121102300-0333321012113223-0011020330023110-0023320100323302-2233001223221013-1112130133013220-3120222213021313-0213203101321111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2212130112311132-1211221331332003-1100323332132310-3223000003303100-3223132211033332-1220133233212122-3212021022323333-0302331101010210"></a>

## stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.default_header_transformation — default_header_transformation / 022101001322 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-017.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-017.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-024.md#canonical-2011132303102222-3023301203121330-2233013320233300-0112111313030012-1103313203030222-1002221110321002-3110022320101011-2333330131330303)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-024.md#canonical-0122221112123120-0023010121003330-3033210010113330-0100100333002100-0211222302303110-3202021130000322-3023123330231032-2033332332032221)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https](data-sources--workload--reference--group-024.md#canonical-3130202111132300-0301212030033211-3312303220203330-1002132121130131-3322330013202010-3013320012301202-1031201103211333-3201303313311311)
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

<a id="canonical-3130321130202103-1120211213103102-2013232201232323-1232022302233333-0210302111003233-2333031323230302-0223231101230011-2101002310330322"></a>

## Direct properties — default_header_transformation / 022101001322 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0203321001110231-2001123122001220-2230313230131103-3131222013210010-0230200121301313-1113023111122312-0133103010103222-0302320133020230"></a>

## Next pages — default_header_transformation / 022101001322 / 4

- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation](data-sources--workload--reference--group-024.md#canonical-3122222212030101-1203221120031210-3333023032021123-0223230011222133-3233112301020012-2011220311220011-1122123312122111-0133320312023110)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-2101222223223301-2010302221021123-3103213310223003-3221232220011220-0012133332002131-0121301010130302-0323032003213322-0231033023001123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3312213211102220-1003101013000032-3000201010023030-3333222121301233-1013201100111232-2223233103133012-1011233330312011-3311231332100121"></a>

## stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.preserve_case_header_transformation — preserve_case_header_transformation / 023120303003 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-017.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-017.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-024.md#canonical-2011132303102222-3023301203121330-2233013320233300-0112111313030012-1103313203030222-1002221110321002-3110022320101011-2333330131330303)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-024.md#canonical-0122221112123120-0023010121003330-3033210010113330-0100100333002100-0211222302303110-3202021130000322-3023123330231032-2033332332032221)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https](data-sources--workload--reference--group-024.md#canonical-3130202111132300-0301212030033211-3312303220203330-1002132121130131-3322330013202010-3013320012301202-1031201103211333-3201303313311311)
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

<a id="canonical-2211212130110202-0323123200110022-3022320123001222-2323011033333223-3230213211230002-2031312333221322-3301122131232120-3223222100103200"></a>

## Direct properties — preserve_case_header_transformation / 023120303003 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3123001230230323-3032031232313330-2201111000010221-1233101322133320-1221132332102312-0203213303202012-1022211333121022-0100332201301210"></a>

## Next pages — preserve_case_header_transformation / 023120303003 / 4

- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation](data-sources--workload--reference--group-024.md#canonical-3122222212030101-1203221120031210-3333023032021123-0223230011222133-3233112301020012-2011220311220011-1122123312122111-0133320312023110)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-1022122133111101-1332003303032202-1301322031220110-2332232111311102-3200000213102003-3031002202333202-2133022321322313-2003323102002120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->
