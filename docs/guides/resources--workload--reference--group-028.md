---
page_title: "xcsh_workload reference"
subcategory: "Container"
description: "Complete grouped canonical reference for xcsh_workload reference."
---

# xcsh_workload reference

<a id="canonical-1310133202102013-0033223222322201-2132133221033113-0103303032013320-3201132300120302-0112231333011111-1013201200320203-2110133102322321"></a>

## `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.simple_route.http_method` property

Type: `"string"`. Optional.

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

- [path](resources--workload--reference--group-028.md#canonical-3021333212112313-2101331311213202-0232212101132113-3023302321322000-1220322010120102-3322031001132322-3000132030301332-0020122313213010): complete subsection reference.

<a id="canonical-0212220221111023-1000023213211100-3010222331203110-2120131200321123-0312032321201232-1111033102332123-1111121221230001-1321212321130012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.simple_route.auto_host_rewrite` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-018.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-018.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-022.md#canonical-3300101333202310-3223231311312310-2030111033033203-1011103231310023-1313110120311121-1221031312313020-1210302113302033-2110231331000202)
- [stateful_service.advertise_options.advertise_on_public.port](resources--workload--reference--group-025.md#canonical-0120012201110120-2320030110233132-3233102311321321-2202232022223030-2133131131013111-3231210131302330-1022020011332122-3111230313032011)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-025.md#canonical-2203002110120031-3203013111221013-1321101003110222-2103121303322022-0021002301311230-0203100203111101-0321302220100020-1320220233023212)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes](resources--workload--reference--group-027.md#canonical-2300220120202032-2120133330213102-2332100100003133-0101032330232302-2302101320100002-1302333333023222-3232223323232230-0320110220223201)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes](resources--workload--reference--group-027.md#canonical-2302003233200103-3312120222212200-1000311013000201-0201011231133032-2032003100312302-3001032323113031-1031222100120212-0102213202320111)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.simple_route](resources--workload--reference--group-027.md#canonical-3223313232022132-1300023030303322-3120001113102132-2101020221132002-0011333231033012-1032011333210012-0303212231230122-1132022333212200)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.simple_route.auto_host_rewrite

<a id="canonical-0221103303322032-1323133333203112-0203312122100112-1331121213202101-0131231203232113-0320021322101130-3203213311112311-1323103302322001"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
auto_host_rewrite = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3201103303230233-1311003111323331-3201331333211200-2033111332103310-0101023203110110-3332121123222232-1103002012213120-0010321001223021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.simple_route.disable_host_rewrite` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-018.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-018.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-022.md#canonical-3300101333202310-3223231311312310-2030111033033203-1011103231310023-1313110120311121-1221031312313020-1210302113302033-2110231331000202)
- [stateful_service.advertise_options.advertise_on_public.port](resources--workload--reference--group-025.md#canonical-0120012201110120-2320030110233132-3233102311321321-2202232022223030-2133131131013111-3231210131302330-1022020011332122-3111230313032011)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-025.md#canonical-2203002110120031-3203013111221013-1321101003110222-2103121303322022-0021002301311230-0203100203111101-0321302220100020-1320220233023212)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes](resources--workload--reference--group-027.md#canonical-2300220120202032-2120133330213102-2332100100003133-0101032330232302-2302101320100002-1302333333023222-3232223323232230-0320110220223201)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes](resources--workload--reference--group-027.md#canonical-2302003233200103-3312120222212200-1000311013000201-0201011231133032-2032003100312302-3001032323113031-1031222100120212-0102213202320111)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.simple_route](resources--workload--reference--group-027.md#canonical-3223313232022132-1300023030303322-3120001113102132-2101020221132002-0011333231033012-1032011333210012-0303212231230122-1132022333212200)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.simple_route.disable_host_rewrite

<a id="canonical-1302022122012203-0131322123323220-2330210323202011-3202011210100032-3022131310303133-2131123120003130-2010113112121212-0003111123213230"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
disable_host_rewrite = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3021333212112313-2101331311213202-0232212101132113-3023302321322000-1220322010120102-3322031001132322-3000132030301332-0020122313213010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.simple_route.path` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-018.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-018.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-022.md#canonical-3300101333202310-3223231311312310-2030111033033203-1011103231310023-1313110120311121-1221031312313020-1210302113302033-2110231331000202)
- [stateful_service.advertise_options.advertise_on_public.port](resources--workload--reference--group-025.md#canonical-0120012201110120-2320030110233132-3233102311321321-2202232022223030-2133131131013111-3231210131302330-1022020011332122-3111230313032011)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-025.md#canonical-2203002110120031-3203013111221013-1321101003110222-2103121303322022-0021002301311230-0203100203111101-0321302220100020-1320220233023212)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes](resources--workload--reference--group-027.md#canonical-2300220120202032-2120133330213102-2332100100003133-0101032330232302-2302101320100002-1302333333023222-3232223323232230-0320110220223201)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes](resources--workload--reference--group-027.md#canonical-2302003233200103-3312120222212200-1000311013000201-0201011231133032-2032003100312302-3001032323113031-1031222100120212-0102213202320111)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.simple_route](resources--workload--reference--group-027.md#canonical-3223313232022132-1300023030303322-3120001113102132-2101020221132002-0011333231033012-1032011333210012-0303212231230122-1132022333212200)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.simple_route.path

<a id="canonical-0310013323322201-1202302121113021-1300200201013120-0201210320313022-0101012213312010-1122312131132302-1122103100233223-3022001010113202"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
path {
  # Configure direct properties listed below.
}
```

<a id="canonical-3100201222022320-2213330032311103-2201321303112002-2023132222302210-3332201233221210-0332103231301101-1132330232022310-2333220312211131"></a>

### Direct properties for `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.simple_route.path`

<a id="canonical-0102203013100130-1223100310023132-1022212223310223-3020020301221022-0331331102333010-0331023000220221-0103202012031303-2301301222102020"></a>

#### `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.simple_route.path.path` property

Type: `"string"`. Optional.

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

<a id="canonical-3130020022112022-1001330102122313-0220332113232103-1031123101021300-2010000000101110-3123101010013010-0220332332113232-2102011303121323"></a>

<a id="canonical-2001322312032001-1223013230323030-3203232332323221-1131330332223013-0311100303300321-0310331201221110-0101132212121022-0111100312103020"></a>

#### `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.simple_route.path.prefix` property

Type: `"string"`. Optional.

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

<a id="canonical-0321332322203211-0122323313332021-0210000311013321-1322322213232203-0312032212201321-1212322211303333-3321311122102030-3103312102321220"></a>

<a id="canonical-0213212200003202-1023311131101330-0302103132233013-3303132130120210-1003113232321003-0113001103222013-0022000002221022-1322030202232331"></a>

#### `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.simple_route.path.regex` property

Type: `"string"`. Optional.

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

<a id="canonical-2323101323000230-3130303013113131-1312013231212223-1212222320213220-2013000033112020-2320301032133202-3111123300030002-2013131210032122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_on_public.port.port` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-018.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-018.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-022.md#canonical-3300101333202310-3223231311312310-2030111033033203-1011103231310023-1313110120311121-1221031312313020-1210302113302033-2110231331000202)
- [stateful_service.advertise_options.advertise_on_public.port](resources--workload--reference--group-025.md#canonical-0120012201110120-2320030110233132-3233102311321321-2202232022223030-2133131131013111-3231210131302330-1022020011332122-3111230313032011)
- stateful_service.advertise_options.advertise_on_public.port.port

<a id="canonical-1131031200011330-2003320130233000-2310101133021033-2302211303123330-2002100332333100-0222332012221222-3213002133233322-2130331123333212"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
port {
  # Configure direct properties listed below.
}
```

<a id="canonical-3223212202130331-3323003120002300-0221133230321230-0300100020333323-3300330012011022-3000103231100001-1211103212310023-3223122202011000"></a>

### Direct properties for `stateful_service.advertise_options.advertise_on_public.port.port`

- [info](resources--workload--reference--group-028.md#canonical-2322021320231203-3320131000003231-1103323320202131-3000213021202011-0332120303031122-2330331300030221-3032131102120330-2202233022300200): complete subsection reference.

<a id="canonical-2322021320231203-3320131000003231-1103323320202131-3000213021202011-0332120303031122-2330331300030221-3032131102120330-2202233022300200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_on_public.port.port.info` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-018.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-018.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-022.md#canonical-3300101333202310-3223231311312310-2030111033033203-1011103231310023-1313110120311121-1221031312313020-1210302113302033-2110231331000202)
- [stateful_service.advertise_options.advertise_on_public.port](resources--workload--reference--group-025.md#canonical-0120012201110120-2320030110233132-3233102311321321-2202232022223030-2133131131013111-3231210131302330-1022020011332122-3111230313032011)
- [stateful_service.advertise_options.advertise_on_public.port.port](resources--workload--reference--group-028.md#canonical-2323101323000230-3130303013113131-1312013231212223-1212222320213220-2013000033112020-2320301032133202-3111123300030002-2013131210032122)
- stateful_service.advertise_options.advertise_on_public.port.port.info

<a id="canonical-1003200323012333-3123333100232332-0221302030111010-1122030020112123-2111323020001212-2332020222010003-0320231121331300-0202201002111130"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
info {
  # Configure direct properties listed below.
}
```

<a id="canonical-2221011213200023-1102333303300313-0221020212300200-0231331132000233-2103033221031221-0021121112232031-3023011220010313-3220030122023131"></a>

### Direct properties for `stateful_service.advertise_options.advertise_on_public.port.port.info`

<a id="canonical-2130233332120323-2122131012111103-1210300211000022-3332230100113013-2002312130330033-1221120303203303-1022130221321032-0033332212131113"></a>

#### `stateful_service.advertise_options.advertise_on_public.port.port.info.port` property

Type: `"number"`. Optional.

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

<a id="canonical-2032230310010231-0012200323011202-0220320313023101-0003102032213212-3321003002333110-3111231212211022-3112121103121001-3313223020131202"></a>

<a id="canonical-0013100022101233-0031003110301321-2021333303030210-2312131300102301-0331231130323312-3312122010133130-1221333301030233-1330200132101332"></a>

#### `stateful_service.advertise_options.advertise_on_public.port.port.info.protocol` property

Type: `"string"`. Optional.

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

- [same_as_port](resources--workload--reference--group-028.md#canonical-2321001300010003-1102133001310130-2102300131332133-0012202022311302-1023033220032013-1331030230213323-3132030033000222-0302013110321223): complete subsection reference.

<a id="canonical-1302320102002022-0131012232103000-1203033133022133-2113012321323021-2321310223313232-3121333222033301-1201232130313301-1313303013230301"></a>

<a id="canonical-3022320030102103-0020031210033332-3111322333231101-0003323213313330-0213001110213030-0133013020333032-1000233002103002-1013000202303033"></a>

#### `stateful_service.advertise_options.advertise_on_public.port.port.info.target_port` property

Type: `"number"`. Optional.

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

<a id="canonical-2321001300010003-1102133001310130-2102300131332133-0012202022311302-1023033220032013-1331030230213323-3132030033000222-0302013110321223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_on_public.port.port.info.same_as_port` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-018.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-018.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-022.md#canonical-3300101333202310-3223231311312310-2030111033033203-1011103231310023-1313110120311121-1221031312313020-1210302113302033-2110231331000202)
- [stateful_service.advertise_options.advertise_on_public.port](resources--workload--reference--group-025.md#canonical-0120012201110120-2320030110233132-3233102311321321-2202232022223030-2133131131013111-3231210131302330-1022020011332122-3111230313032011)
- [stateful_service.advertise_options.advertise_on_public.port.port](resources--workload--reference--group-028.md#canonical-2323101323000230-3130303013113131-1312013231212223-1212222320213220-2013000033112020-2320301032133202-3111123300030002-2013131210032122)
- [stateful_service.advertise_options.advertise_on_public.port.port.info](resources--workload--reference--group-028.md#canonical-2322021320231203-3320131000003231-1103323320202131-3000213021202011-0332120303031122-2330331300030221-3032131102120330-2202233022300200)
- stateful_service.advertise_options.advertise_on_public.port.port.info.same_as_port

<a id="canonical-2313130303311000-3211112201320300-2210211200333123-0000002212022323-0002031002133011-0023021323032210-3003321203121000-0213203130030321"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
same_as_port = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0201031010010211-0110122211331300-2100213123230032-2210131211202000-0123012121332012-2132032030312023-0111231211021031-2320223020033130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_on_public.port.tcp_loadbalancer` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-018.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-018.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-022.md#canonical-3300101333202310-3223231311312310-2030111033033203-1011103231310023-1313110120311121-1221031312313020-1210302113302033-2110231331000202)
- [stateful_service.advertise_options.advertise_on_public.port](resources--workload--reference--group-025.md#canonical-0120012201110120-2320030110233132-3233102311321321-2202232022223030-2133131131013111-3231210131302330-1022020011332122-3111230313032011)
- stateful_service.advertise_options.advertise_on_public.port.tcp_loadbalancer

<a id="canonical-2002323122321000-2112022202131013-0020210332321210-2110030111322210-0230333322111022-2303230321311131-3000022230202211-0112030103210210"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
tcp_loadbalancer {
  # Configure direct properties listed below.
}
```

<a id="canonical-0032133002213222-3122123222023331-3203211322122220-0103212233021002-0230102222123301-1022323033030113-3022011032123201-3033021121021230"></a>

### Direct properties for `stateful_service.advertise_options.advertise_on_public.port.tcp_loadbalancer`

<a id="canonical-1301230131031211-1002122223323222-1020200132030311-1012113121300110-3031230310310131-3331030332311000-0331332220001120-3030322330112030"></a>

#### `stateful_service.advertise_options.advertise_on_public.port.tcp_loadbalancer.domains` property

Type: `["list", "string"]`. Optional.

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

<a id="canonical-2320313100122030-1230022111301233-0102031311332303-3002013303111302-2313303230010022-1323112332003230-3012230300101132-0020201133031222"></a>

<a id="canonical-0331112030003310-0220123303200133-2131001123103321-1132332133223030-2100333233000013-1231312331023111-3213220033131132-3300033332021300"></a>

#### `stateful_service.advertise_options.advertise_on_public.port.tcp_loadbalancer.with_sni` property

Type: `"bool"`. Optional.

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

<a id="canonical-3202321232103323-1000231221221030-0001122222031132-1003133321010023-1002201331103010-1230130022220311-0032100201113003-0333121223233010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.do_not_advertise` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-018.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-018.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- stateful_service.advertise_options.do_not_advertise

<a id="canonical-3002301133331320-1322203131133310-1112020032333111-0133302112022120-0233211032202220-0122333032100300-2110130011122200-1221220222220303"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
do_not_advertise = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3203012310220003-2210330300331110-2103223003030222-2222201323310103-2300110002113011-0130333021003321-0201020033321220-1102301103203230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.configuration` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-018.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- stateful_service.configuration

<a id="canonical-1201322303303202-0221320010021012-1303000331110131-0101010002332012-2322111220213321-1101133003320203-1013112320231110-0212330121232013"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
configuration {
  # Configure direct properties listed below.
}
```

<a id="canonical-3101202312333220-2211200333103323-3221321230223013-3003313332022001-1200012223110010-0030111101023122-3322333232202020-0322003100320102"></a>

### Direct properties for `stateful_service.configuration`

- [parameters](resources--workload--reference--group-028.md#canonical-1302213033220212-3202132112103233-0311330233023033-0321030031223222-2102010112112002-2320312322333220-0032233021200100-1022102313033123): complete subsection reference.

<a id="canonical-1302213033220212-3202132112103233-0311330233023033-0321030031223222-2102010112112002-2320312322333220-0032233021200100-1022102313033123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.configuration.parameters` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-018.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.configuration](resources--workload--reference--group-028.md#canonical-3203012310220003-2210330300331110-2103223003030222-2222201323310103-2300110002113011-0130333021003321-0201020033321220-1102301103203230)
- stateful_service.configuration.parameters

<a id="canonical-0333303303203013-2021330333311330-0211211331210301-1220230230230102-1013220331210112-1110310111112203-1213212300210113-0003113111231111"></a>

Type: `"object"`. list nested block, Optional.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

Terraform syntax:

```terraform
parameters {
  # Configure direct properties listed below.
}
```

<a id="canonical-3332110010001010-2013313311220303-3202133223200100-2202303301013330-0102032213323300-2000021002321300-1023233113132200-2213210011233132"></a>

### Direct properties for `stateful_service.configuration.parameters`

- [env_var](resources--workload--reference--group-028.md#canonical-1323013220320120-3113201120222231-3323230231223312-3132031213122030-0230203302303023-0002322101130002-2012310022211131-2033101310312030): complete subsection reference.

- [file](resources--workload--reference--group-028.md#canonical-3111133103101303-0323302220303200-2031333233121200-0222313132030022-3332112322312012-1320033010330232-0302211321120013-3033112100233210): complete subsection reference.

<a id="canonical-1323013220320120-3113201120222231-3323230231223312-3132031213122030-0230203302303023-0002322101130002-2012310022211131-2033101310312030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.configuration.parameters.env_var` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-018.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.configuration](resources--workload--reference--group-028.md#canonical-3203012310220003-2210330300331110-2103223003030222-2222201323310103-2300110002113011-0130333021003321-0201020033321220-1102301103203230)
- [stateful_service.configuration.parameters](resources--workload--reference--group-028.md#canonical-1302213033220212-3202132112103233-0311330233023033-0321030031223222-2102010112112002-2320312322333220-0032233021200100-1022102313033123)
- stateful_service.configuration.parameters.env_var

<a id="canonical-0222323200322103-1001330330310020-2223213220122033-0312213222000102-2220311212203301-1313212103302002-0322202321313222-0202323113132010"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
env_var {
  # Configure direct properties listed below.
}
```

<a id="canonical-0020312123011131-2300331300210112-0112330203310230-1222203312201120-1301313110223211-3013223220201110-2332000031220123-2123232231000001"></a>

### Direct properties for `stateful_service.configuration.parameters.env_var`

<a id="canonical-3222100201300110-1130032302211021-2102200003012220-0013002332221001-1231100211012313-3131321100320010-1221203201011112-0022301132202233"></a>

#### `stateful_service.configuration.parameters.env_var.name` property

Type: `"string"`. Optional.

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

<a id="canonical-0332132212021231-2033021132011311-0000202101331333-3303302033331310-2013233012021300-1011230132310033-0202212131213213-2110112201302011"></a>

<a id="canonical-0310112321302022-0322131203330230-3303032032233223-3313020000000201-0031020131301020-1101003112103321-3203012310201033-2200000231211302"></a>

#### `stateful_service.configuration.parameters.env_var.value` property

Type: `"string"`. Optional.

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
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-3111133103101303-0323302220303200-2031333233121200-0222313132030022-3332112322312012-1320033010330232-0302211321120013-3033112100233210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.configuration.parameters.file` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-018.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.configuration](resources--workload--reference--group-028.md#canonical-3203012310220003-2210330300331110-2103223003030222-2222201323310103-2300110002113011-0130333021003321-0201020033321220-1102301103203230)
- [stateful_service.configuration.parameters](resources--workload--reference--group-028.md#canonical-1302213033220212-3202132112103233-0311330233023033-0321030031223222-2102010112112002-2320312322333220-0032233021200100-1022102313033123)
- stateful_service.configuration.parameters.file

<a id="canonical-0020013323122130-3121213332131131-1201313310110320-3100321010323213-0131121232132012-3131122213103311-0321012310213221-2011021002111231"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
file {
  # Configure direct properties listed below.
}
```

<a id="canonical-1121301021132111-3203011123333011-3332331121322311-3131210110221130-1311210122122222-0310111212232122-2002011302332002-0230322022322002"></a>

### Direct properties for `stateful_service.configuration.parameters.file`

<a id="canonical-0203223030221121-2121202222200002-1233201032223132-2033103132103213-1203001310112123-0211231132221030-2131003000033103-2032110113000032"></a>

#### `stateful_service.configuration.parameters.file.data` property

Type: `"string"`. Optional.

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
    "ves.io.schema.rules.string.max_len": "16384",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "16384",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

- [mount](resources--workload--reference--group-028.md#canonical-0233000201230203-0013120130001113-1020123313011012-3331231213030323-1031331300313332-1022000203000000-1200013002010223-1333333023011110): complete subsection reference.

<a id="canonical-0112000310120033-3010033122211331-2221033233201121-2000303021222011-0003121003123032-3230030011331330-3020223203233221-0200322123023311"></a>

<a id="canonical-2000303021211022-3301133000112123-2231230102132211-3112101010011111-3200220123132203-1012312333201122-1012303330012313-0133031333010202"></a>

#### `stateful_service.configuration.parameters.file.name` property

Type: `"string"`. Optional.

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
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-1110122130212220-3323121330233113-0323300323301321-2231021321210230-0001320213231033-3220033211102002-3001221000023113-3033222101211100"></a>

<a id="canonical-0333130101010120-2011033110033132-1302022222131223-3101202122101221-2203123132300301-1032021023300221-0023131320010120-1223100001022321"></a>

#### `stateful_service.configuration.parameters.file.volume_name` property

Type: `"string"`. Optional.

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
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-0233000201230203-0013120130001113-1020123313011012-3331231213030323-1031331300313332-1022000203000000-1200013002010223-1333333023011110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.configuration.parameters.file.mount` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-018.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.configuration](resources--workload--reference--group-028.md#canonical-3203012310220003-2210330300331110-2103223003030222-2222201323310103-2300110002113011-0130333021003321-0201020033321220-1102301103203230)
- [stateful_service.configuration.parameters](resources--workload--reference--group-028.md#canonical-1302213033220212-3202132112103233-0311330233023033-0321030031223222-2102010112112002-2320312322333220-0032233021200100-1022102313033123)
- [stateful_service.configuration.parameters.file](resources--workload--reference--group-028.md#canonical-3111133103101303-0323302220303200-2031333233121200-0222313132030022-3332112322312012-1320033010330232-0302211321120013-3033112100233210)
- stateful_service.configuration.parameters.file.mount

<a id="canonical-3222300012011013-3222112022030031-0333003331130000-2003211133012212-1331102322213200-0221012030311100-3120102101322322-1213110130012231"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
mount {
  # Configure direct properties listed below.
}
```

<a id="canonical-1133111230301231-1112301120323003-1302001200203301-1332001011201012-0210232323021222-3201131123130113-1233130200102232-1202010312222211"></a>

### Direct properties for `stateful_service.configuration.parameters.file.mount`

<a id="canonical-2322311313103131-0101221031322302-3232233102110110-1001010021312101-0212133310022003-3301312102022112-2002122131021111-1210230132133110"></a>

#### `stateful_service.configuration.parameters.file.mount.mode` property

Type: `"string"`. Optional.

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

<a id="canonical-3031321122023211-2211202131311023-0223203332202131-3032112021230313-1212110323111112-1012200221120121-2312323132212233-1230310101222201"></a>

<a id="canonical-0010210003102222-0102222213323010-2100322233300121-1311323232231303-2330012330303101-0230133222321222-2333011232113322-1001212231011322"></a>

#### `stateful_service.configuration.parameters.file.mount.mount_path` property

Type: `"string"`. Optional.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-3312322121030111-2300020310330113-1113101332223203-2112310233010112-1112301303032300-2110222221021323-2230233323030312-3300300103123303"></a>

<a id="canonical-2022200001002102-3032002203003311-2010220331111232-0113112021331110-1121333111203120-2010300020032301-0210130003220110-2212023023311203"></a>

#### `stateful_service.configuration.parameters.file.mount.sub_path` property

Type: `"string"`. Optional.

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
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-3221212333032323-1011122002030020-0332100102213333-2230123311320001-3101132000200220-3323302320310311-1200332113313321-1231002100113301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.containers` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-018.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- stateful_service.containers

<a id="canonical-1113030003011000-0010130220333121-0021121013302333-2232320020302313-0033322320320313-0222330100312121-0130123102002030-2300120133302233"></a>

Type: `"object"`. list nested block, Optional.

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

Terraform syntax:

```terraform
containers {
  # Configure direct properties listed below.
}
```

<a id="canonical-1020003202220121-3212302322020012-1333223330210231-2001323200032233-0332202313310032-0011330300332133-3110131212031121-3100123012133313"></a>

### Direct properties for `stateful_service.containers`

<a id="canonical-0302200203121231-3303000323310310-1102102212320233-2010201021110301-1221000222221312-2120310110302112-2233223023110002-1131122322321320"></a>

#### `stateful_service.containers.args` property

Type: `["list", "string"]`. Optional.

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
    "ves.io.schema.rules.repeated.max_items": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "128"
  }
}
```

<a id="canonical-0210003003322332-0203311023330301-0220013100311120-1223131303213312-2100000110232023-2003100103013110-3311020230120200-0201001010103321"></a>

<a id="canonical-3011123332120023-0103333322013300-1331201323113113-3022332310003321-2300023203123100-1301212313101203-0331112111001102-3202311201112102"></a>

#### `stateful_service.containers.command` property

Type: `["list", "string"]`. Optional.

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
    "ves.io.schema.rules.repeated.max_items": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "128"
  }
}
```

- [custom_flavor](resources--workload--reference--group-028.md#canonical-3213302013021002-0200100013303102-2133231113010230-0012320031302120-0111021223212331-2223102120311311-1102330223321001-0011001011331122): complete subsection reference.

- [default_flavor](resources--workload--reference--group-028.md#canonical-2220323323131313-3213223223212102-1023010101002123-1010210210130022-3213010130010023-0000113022111132-1222132313331300-2003202303001101): complete subsection reference.

<a id="canonical-3021233321033321-3101113132213002-2310100123313111-0020132221001310-0130100310023230-2100001201220301-2121322320201111-1133020322102031"></a>

<a id="canonical-2013330122001002-2112313231131333-2311112213011022-1000103021110213-1233221332121103-2030302232200302-1113110333300033-1131130310120323"></a>

#### `stateful_service.containers.flavor` property

Type: `"string"`. Optional.

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

- [image](resources--workload--reference--group-028.md#canonical-0210130211131331-2131110001320022-1333130131033002-1332001300303301-1003303222231001-1031320230211023-3113003031002110-0122201210203302): complete subsection reference.

<a id="canonical-0330321120110120-2302311101022233-3302032333232112-0302123232213202-2200202311200113-3220232030131123-3132130110230333-3331011120111102"></a>

<a id="canonical-0131322230132320-3121221122212301-1011003032022033-1213100001233030-2331121331122001-0110000130120312-1022111103210013-2101112133102000"></a>

#### `stateful_service.containers.init_container` property

Type: `"bool"`. Optional.

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

- [liveness_check](resources--workload--reference--group-028.md#canonical-3020201133320111-3333012222010310-1020001033321333-2031333323312030-0322021121300122-0110330133313201-1000022031310321-0321102330322111): complete subsection reference.

<a id="canonical-3233232112033033-3233003233311113-3110313330213300-2232123200001302-1330311112103130-3030221002203301-2002303301333022-2321333121132212"></a>

<a id="canonical-0333013021102313-0100201232132321-0301113132032013-2303120333003101-0222023033222222-2202320233030022-3211220011123312-1031030003001312"></a>

#### `stateful_service.containers.name` property

Type: `"string"`. Optional.

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
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

- [readiness_check](resources--workload--reference--group-028.md#canonical-2101233010110302-0023221233331300-0020012001202110-3111101033233313-3131302323332021-0313320323031313-1031100100130333-1131132213032321): complete subsection reference.

<a id="canonical-3213302013021002-0200100013303102-2133231113010230-0012320031302120-0111021223212331-2223102120311311-1102330223321001-0011001011331122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.containers.custom_flavor` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-018.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.containers](resources--workload--reference--group-028.md#canonical-3221212333032323-1011122002030020-0332100102213333-2230123311320001-3101132000200220-3323302320310311-1200332113313321-1231002100113301)
- stateful_service.containers.custom_flavor

<a id="canonical-3033231233130120-1132332211121203-1332322321023230-0321020313200330-1332100013131120-0000223100131203-1133313320211013-3331212112000001"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
custom_flavor {
  # Configure direct properties listed below.
}
```

<a id="canonical-2111231221210110-0131131321002333-2130321201202030-2031102121122221-3300113330033210-3123311100320203-3202312131032200-3132110112313011"></a>

### Direct properties for `stateful_service.containers.custom_flavor`

<a id="canonical-3130313112031001-3332132002101332-1001112020200220-1012232221301123-1111031200120222-1001111121330233-1302131223022300-0200002333101232"></a>

#### `stateful_service.containers.custom_flavor.name` property

Type: `"string"`. Optional.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-3110232020213111-1022113322121023-2121012112310223-3113201233001103-3323011033222003-1010213333010033-3133133320313110-2121300131323010"></a>

<a id="canonical-2300100101122211-2311020311223323-1213011012223322-3212003203020033-0323312333100112-1330102103312320-2310122200033122-0121131320013212"></a>

#### `stateful_service.containers.custom_flavor.namespace` property

Type: `"string"`. Optional, Computed.

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

<a id="canonical-2322213101110213-3223320233320031-2320333221121011-2100311012133311-2213202022300311-1312010303311022-3103033323210023-2202133022031122"></a>

<a id="canonical-2121221000211101-1201220202003120-1130303113333210-0210000232102132-1033032233110132-1222122313201233-3023003223303130-0031100111222300"></a>

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
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-2220323323131313-3213223223212102-1023010101002123-1010210210130022-3213010130010023-0000113022111132-1222132313331300-2003202303001101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.containers.default_flavor` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-018.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.containers](resources--workload--reference--group-028.md#canonical-3221212333032323-1011122002030020-0332100102213333-2230123311320001-3101132000200220-3323302320310311-1200332113313321-1231002100113301)
- stateful_service.containers.default_flavor

<a id="canonical-3233211301023200-1303330031010210-2303303020321200-3120032331013331-3330302300023332-1230201322232003-3012120300021030-0130110333103310"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
default_flavor = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0210130211131331-2131110001320022-1333130131033002-1332001300303301-1003303222231001-1031320230211023-3113003031002110-0122201210203302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.containers.image` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-018.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.containers](resources--workload--reference--group-028.md#canonical-3221212333032323-1011122002030020-0332100102213333-2230123311320001-3101132000200220-3323302320310311-1200332113313321-1231002100113301)
- stateful_service.containers.image

<a id="canonical-3001023131330212-3030133103033233-3102010310130023-1233001101222100-2211212033323210-1223220120212231-3320001030021301-0330311222111212"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
image {
  # Configure direct properties listed below.
}
```

<a id="canonical-3030212333300221-2210113203303132-0031223331120313-3322131001300122-0132132331022210-2120320232310031-1332033103132111-3330200001032212"></a>

### Direct properties for `stateful_service.containers.image`

- [container_registry](resources--workload--reference--group-028.md#canonical-1310113331123012-2113223012123210-3020033210132330-2130201100002301-3023312013303313-3102233310233203-1331211112031221-1100202031131313): complete subsection reference.

<a id="canonical-1313132001301020-3131313223113330-1201013132303011-0113123120230310-3212200222323303-2330120222123310-1202010132330130-0212111203232232"></a>

<a id="canonical-3222032330310231-3021012313231310-2310330313000111-0331130121312131-3330022122120001-3201031121120031-1021222000013012-0103132211121322"></a>

#### `stateful_service.containers.image.name` property

Type: `"string"`. Optional.

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
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

- [public](resources--workload--reference--group-028.md#canonical-1010202101123303-0233303332120030-3031220123331031-2201301010020310-2120002112132330-0230230121311201-3221211223232113-0231313011210133): complete subsection reference.

<a id="canonical-3210102203303223-1212113132033300-3203323112012213-2200320232331210-1300023312202232-0220132122221030-2020030331230222-1302333101212122"></a>

<a id="canonical-3033120102133220-0313221121130231-2313321322001311-3012331312223201-1130201031103002-1222223021003311-1331021032213232-3011002031331023"></a>

#### `stateful_service.containers.image.pull_policy` property

Type: `"string"`. Optional.

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

<a id="canonical-1310113331123012-2113223012123210-3020033210132330-2130201100002301-3023312013303313-3102233310233203-1331211112031221-1100202031131313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.containers.image.container_registry` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-018.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.containers](resources--workload--reference--group-028.md#canonical-3221212333032323-1011122002030020-0332100102213333-2230123311320001-3101132000200220-3323302320310311-1200332113313321-1231002100113301)
- [stateful_service.containers.image](resources--workload--reference--group-028.md#canonical-0210130211131331-2131110001320022-1333130131033002-1332001300303301-1003303222231001-1031320230211023-3113003031002110-0122201210203302)
- stateful_service.containers.image.container_registry

<a id="canonical-3101103113313320-3033213030112022-0311203332333223-0210201103123211-1231213013131010-3332331011020231-0230213102133110-3113232101000123"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
container_registry {
  # Configure direct properties listed below.
}
```

<a id="canonical-0122101001221030-1321001302203021-1331031131031021-3232130211203213-0212033132313200-1023010333302121-3020023111313013-1200203210220312"></a>

### Direct properties for `stateful_service.containers.image.container_registry`

<a id="canonical-1221113233112033-0133110213203212-1313232030311232-1213330021332302-0021113133200011-3320011330003003-0311322033202333-2310221201123002"></a>

#### `stateful_service.containers.image.container_registry.name` property

Type: `"string"`. Optional.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-0312131233211133-1330101321013023-0223010130023313-3301212002100233-0233210102011231-1232200321132320-2132010213013321-0313011220013220"></a>

<a id="canonical-2313022230111132-2331232100021210-1021202020112022-3303211131220023-2103311332333021-0330313232203132-0322012231311002-1322001303212002"></a>

#### `stateful_service.containers.image.container_registry.namespace` property

Type: `"string"`. Optional, Computed.

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

<a id="canonical-0221012130201021-3232113212001123-3303332333000322-1130321120232221-1203323332222121-1320023012110322-1320220323200201-0023021302313012"></a>

<a id="canonical-1020021310032220-1023321121000300-0232123022201003-0113103113202012-1103223311303100-0303023100310320-2132301010122113-1103031323310302"></a>

#### `stateful_service.containers.image.container_registry.tenant` property

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
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-1010202101123303-0233303332120030-3031220123331031-2201301010020310-2120002112132330-0230230121311201-3221211223232113-0231313011210133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.containers.image.public` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-018.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.containers](resources--workload--reference--group-028.md#canonical-3221212333032323-1011122002030020-0332100102213333-2230123311320001-3101132000200220-3323302320310311-1200332113313321-1231002100113301)
- [stateful_service.containers.image](resources--workload--reference--group-028.md#canonical-0210130211131331-2131110001320022-1333130131033002-1332001300303301-1003303222231001-1031320230211023-3113003031002110-0122201210203302)
- stateful_service.containers.image.public

<a id="canonical-0210003223122312-0333203231302013-3100003211123231-0112320200321100-2112331000110030-0131313112220211-2332101212113130-2322332322112131"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
public = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3020201133320111-3333012222010310-1020001033321333-2031333323312030-0322021121300122-0110330133313201-1000022031310321-0321102330322111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.containers.liveness_check` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-018.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.containers](resources--workload--reference--group-028.md#canonical-3221212333032323-1011122002030020-0332100102213333-2230123311320001-3101132000200220-3323302320310311-1200332113313321-1231002100113301)
- stateful_service.containers.liveness_check

<a id="canonical-0320023333012002-2311211021333230-3011110221210203-1110030132210302-1101021123023103-1312200232210212-1332232213330211-2201331123001213"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
liveness_check {
  # Configure direct properties listed below.
}
```

<a id="canonical-0033300213223020-2212111221110023-2230313301120320-1123213333310111-1012213301303210-0002221232000232-3322210331220011-1130011303120012"></a>

### Direct properties for `stateful_service.containers.liveness_check`

- [exec_health_check](resources--workload--reference--group-028.md#canonical-3202211021213013-0010223010220301-3100011320332220-1133312022220330-3102212031212130-0300013331012011-3320111013133111-2110023110323311): complete subsection reference.

<a id="canonical-3322303113312020-0223121313300111-0010233030022022-3002011111202131-1302230030023122-1303001322310031-0311001303032323-1030331220023000"></a>

<a id="canonical-3211001002311011-0123222122113231-2020110011233212-3302021112101310-0200131002311233-0130232330312212-3333101332210300-3010123333313200"></a>

#### `stateful_service.containers.liveness_check.healthy_threshold` property

Type: `"number"`. Optional.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

- [http_health_check](resources--workload--reference--group-028.md#canonical-0110313303111000-0201320232102000-3002132013033323-0000210321132302-0130313122020010-1120231201101213-2331033302000122-2030120120202000): complete subsection reference.

<a id="canonical-0123111122003333-1010233333030220-1210120211023102-2033320031232111-3231102313222201-0121133000021320-2210010112023331-2122031333130032"></a>

<a id="canonical-2101001222001031-1111310020200210-0313322031321022-2210012320211001-2001300332233312-1232233221003133-3222023111200133-3213213202330021"></a>

#### `stateful_service.containers.liveness_check.initial_delay` property

Type: `"number"`. Optional.

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
    "ves.io.schema.rules.uint32.lte": "600"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "600"
  }
}
```

<a id="canonical-1030313032132003-0111222133222120-1011000201323130-2323232013303030-2311310033201302-1310003333222332-3021233202231130-0320103223200132"></a>

<a id="canonical-1331102303011101-1213133122132312-3311320132032330-0030221102123003-0220031121001030-0222203311132002-1321211311022020-1100012022122133"></a>

#### `stateful_service.containers.liveness_check.interval` property

Type: `"number"`. Optional.

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

- [tcp_health_check](resources--workload--reference--group-028.md#canonical-2211202333323310-1203311303331120-2033130120013202-0200033312120121-3033212311030220-3331130220320002-0311323221310111-0032011220101332): complete subsection reference.

<a id="canonical-1102232021212222-1032032020013311-2031203321102101-1130033330102223-3110323200210002-0131303302032030-3013113213211011-2321221322322311"></a>

<a id="canonical-1122103011132331-1122320221231330-2120102102331311-3030203023300321-0002310133130111-0130320021211012-3223102212123020-2132010230011103"></a>

#### `stateful_service.containers.liveness_check.timeout` property

Type: `"number"`. Optional.

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

<a id="canonical-2130032331112332-0303112021320233-2121303012233223-2220213313020121-0100323102213231-1222211201232112-3011233203331223-0000133200233232"></a>

<a id="canonical-2103333202221110-1110123302011031-0102210110101023-0331021322013311-1202120002332002-3232323333320301-3200213020130131-3120102311111223"></a>

#### `stateful_service.containers.liveness_check.unhealthy_threshold` property

Type: `"number"`. Optional.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-3202211021213013-0010223010220301-3100011320332220-1133312022220330-3102212031212130-0300013331012011-3320111013133111-2110023110323311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.containers.liveness_check.exec_health_check` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-018.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.containers](resources--workload--reference--group-028.md#canonical-3221212333032323-1011122002030020-0332100102213333-2230123311320001-3101132000200220-3323302320310311-1200332113313321-1231002100113301)
- [stateful_service.containers.liveness_check](resources--workload--reference--group-028.md#canonical-3020201133320111-3333012222010310-1020001033321333-2031333323312030-0322021121300122-0110330133313201-1000022031310321-0321102330322111)
- stateful_service.containers.liveness_check.exec_health_check

<a id="canonical-0223120101113202-3032331300013333-0223123233121021-2310311233213303-0322011202031033-0311102221130112-0021013130312300-3232313322213123"></a>

Type: `"object"`. single nested block, Optional.

ExecHealthCheckType describes a health check based on 'run in container' action. Exit status of 0 is
treated as live/healthy and non-zero is unhealthy.

Additional upstream details:

ExecHealthCheckType describes a health check based on "run in container" action.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
exec_health_check {
  # Configure direct properties listed below.
}
```

<a id="canonical-3312222310321021-2232022320021122-2001012220022332-3002101102000230-3311233312220201-1030212032232310-2002032232023130-3323021003033212"></a>

### Direct properties for `stateful_service.containers.liveness_check.exec_health_check`

<a id="canonical-0232202132022102-0301223332331101-0311222232322313-2213301310101101-1230013021110003-1121121131330011-3231132121031233-0302220220210030"></a>

#### `stateful_service.containers.liveness_check.exec_health_check.command` property

Type: `["list", "string"]`. Optional.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-0110313303111000-0201320232102000-3002132013033323-0000210321132302-0130313122020010-1120231201101213-2331033302000122-2030120120202000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.containers.liveness_check.http_health_check` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-018.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.containers](resources--workload--reference--group-028.md#canonical-3221212333032323-1011122002030020-0332100102213333-2230123311320001-3101132000200220-3323302320310311-1200332113313321-1231002100113301)
- [stateful_service.containers.liveness_check](resources--workload--reference--group-028.md#canonical-3020201133320111-3333012222010310-1020001033321333-2031333323312030-0322021121300122-0110330133313201-1000022031310321-0321102330322111)
- stateful_service.containers.liveness_check.http_health_check

<a id="canonical-0211101021013333-2301021331102013-0301213132000120-2111100322220323-2323121311200200-2312211223202023-3011202132210013-0311003301333103"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
http_health_check {
  # Configure direct properties listed below.
}
```

<a id="canonical-3233322133120130-0030233311301321-1230320102111100-2123230022002203-2211233121111331-2132031203121303-2320110130020300-0021312122333132"></a>

### Direct properties for `stateful_service.containers.liveness_check.http_health_check`

<a id="canonical-0220003000022122-3212233013312302-0020311323032101-3212102012023232-1211001113032230-1001223213002212-2232233202322113-0312333200322332"></a>

#### `stateful_service.containers.liveness_check.http_health_check.headers` property

Type: `["map", "string"]`. Optional.

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

<a id="canonical-3201003023133333-0203021131021100-2113010312212101-0201231132033221-3331030122010303-3101012121213010-1133122320103103-0120223110001232"></a>

<a id="canonical-1110020212233201-2002112202210023-1302113112021113-2323311011002023-1211121003030022-3230332202003231-0213010301320020-2302221003303101"></a>

#### `stateful_service.containers.liveness_check.http_health_check.host_header` property

Type: `"string"`. Optional.

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
    "ves.io.schema.rules.string.hostport": "true",
    "ves.io.schema.rules.string.max_len": "262"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostport": "true",
    "ves.io.schema.rules.string.max_len": "262"
  }
}
```

<a id="canonical-1113001212213102-0203301231003132-1320003333310121-3322223013310200-0020121000032030-0210312011330312-3332303230201302-1110132222230102"></a>

<a id="canonical-2013233223122111-0133311320202210-0200001210210202-0031131323201133-1023213001123323-2311101223311310-3130023301020020-1321311110121112"></a>

#### `stateful_service.containers.liveness_check.http_health_check.path` property

Type: `"string"`. Optional.

Path. Path to access on the HTTP server.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

- [port](resources--workload--reference--group-028.md#canonical-0303322310300133-0233313213223100-1012003101331013-0010311203033303-2102103210121101-1012100013320123-3311232311313232-2011322020221233): complete subsection reference.

<a id="canonical-0303322310300133-0233313213223100-1012003101331013-0010311203033303-2102103210121101-1012100013320123-3311232311313232-2011322020221233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.containers.liveness_check.http_health_check.port` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-018.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.containers](resources--workload--reference--group-028.md#canonical-3221212333032323-1011122002030020-0332100102213333-2230123311320001-3101132000200220-3323302320310311-1200332113313321-1231002100113301)
- [stateful_service.containers.liveness_check](resources--workload--reference--group-028.md#canonical-3020201133320111-3333012222010310-1020001033321333-2031333323312030-0322021121300122-0110330133313201-1000022031310321-0321102330322111)
- [stateful_service.containers.liveness_check.http_health_check](resources--workload--reference--group-028.md#canonical-0110313303111000-0201320232102000-3002132013033323-0000210321132302-0130313122020010-1120231201101213-2331033302000122-2030120120202000)
- stateful_service.containers.liveness_check.http_health_check.port

<a id="canonical-1001013021330012-0222020302120323-2322330332331013-1223133222120110-0221201111100031-2201221013332322-3112020313021130-1011210022031002"></a>

Type: `"object"`. single nested block, Optional.

Port. Port

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

Terraform syntax:

```terraform
port {
  # Configure direct properties listed below.
}
```

<a id="canonical-1110121123321303-3112101132111121-2200113332110321-1233033122020223-0222022102003120-2212132111121230-0012321031113101-3233223322123001"></a>

### Direct properties for `stateful_service.containers.liveness_check.http_health_check.port`

<a id="canonical-0320200322102100-2010132323022323-0330131102013213-1110221020301112-1313301313033023-0103201020023302-2013100133123101-2300100011303322"></a>

#### `stateful_service.containers.liveness_check.http_health_check.port.name` property

Type: `"string"`. Optional.

Port Name. Exclusive with \[num\] Port Name.

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

<a id="canonical-2321322102312113-0022130022201311-0023302001123111-0331022313320113-2310232033231203-3320301221033302-3100310230221113-3330213233302030"></a>

<a id="canonical-3111131110130202-1303203030220221-1111311232332030-2133313332313110-0112322333210312-1120303021310011-2313000322310332-0210212133121221"></a>

#### `stateful_service.containers.liveness_check.http_health_check.port.num` property

Type: `"number"`. Optional.

Port Number. Exclusive with \[name\] Port number.

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
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-2211202333323310-1203311303331120-2033130120013202-0200033312120121-3033212311030220-3331130220320002-0311323221310111-0032011220101332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.containers.liveness_check.tcp_health_check` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-018.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.containers](resources--workload--reference--group-028.md#canonical-3221212333032323-1011122002030020-0332100102213333-2230123311320001-3101132000200220-3323302320310311-1200332113313321-1231002100113301)
- [stateful_service.containers.liveness_check](resources--workload--reference--group-028.md#canonical-3020201133320111-3333012222010310-1020001033321333-2031333323312030-0322021121300122-0110330133313201-1000022031310321-0321102330322111)
- stateful_service.containers.liveness_check.tcp_health_check

<a id="canonical-0031212212331002-3110330332021301-2021021233110110-2120230113313213-0300020223200232-2033033032333333-2122223312221301-3020002222202330"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
tcp_health_check {
  # Configure direct properties listed below.
}
```

<a id="canonical-3110122010310333-0101302301002201-3032010002223210-0213123202022103-2303131101111031-2021033211001132-1130033311122130-0302033230103303"></a>

### Direct properties for `stateful_service.containers.liveness_check.tcp_health_check`

- [port](resources--workload--reference--group-028.md#canonical-1023303303332203-1232211220000333-2212302213222012-0123311331133103-3030230132332133-1332333111132332-0111333123031020-0130101220000030): complete subsection reference.

<a id="canonical-1023303303332203-1232211220000333-2212302213222012-0123311331133103-3030230132332133-1332333111132332-0111333123031020-0130101220000030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.containers.liveness_check.tcp_health_check.port` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-018.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.containers](resources--workload--reference--group-028.md#canonical-3221212333032323-1011122002030020-0332100102213333-2230123311320001-3101132000200220-3323302320310311-1200332113313321-1231002100113301)
- [stateful_service.containers.liveness_check](resources--workload--reference--group-028.md#canonical-3020201133320111-3333012222010310-1020001033321333-2031333323312030-0322021121300122-0110330133313201-1000022031310321-0321102330322111)
- [stateful_service.containers.liveness_check.tcp_health_check](resources--workload--reference--group-028.md#canonical-2211202333323310-1203311303331120-2033130120013202-0200033312120121-3033212311030220-3331130220320002-0311323221310111-0032011220101332)
- stateful_service.containers.liveness_check.tcp_health_check.port

<a id="canonical-2202201102011123-2023313313010221-0102033000003123-1033300322012313-0131103000332200-0301332023303121-0330233231112122-1221023021021301"></a>

Type: `"object"`. single nested block, Optional.

Port. Port

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

Terraform syntax:

```terraform
port {
  # Configure direct properties listed below.
}
```

<a id="canonical-3322011322212301-3131033232321323-0120131233201333-3100331012110212-2101312331303232-0100213313323031-3302002300221213-1320213130123213"></a>

### Direct properties for `stateful_service.containers.liveness_check.tcp_health_check.port`

<a id="canonical-3303120011233211-1023012031123231-0001203322220000-1102230303123110-1131133332301212-2222233012001232-1032233210312033-3322133110111133"></a>

#### `stateful_service.containers.liveness_check.tcp_health_check.port.name` property

Type: `"string"`. Optional.

Port Name. Exclusive with \[num\] Port Name.

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

<a id="canonical-2021201200212211-3312221121211321-3113110330110001-0031203212111223-3220002213230233-1120303110213221-1131021003130131-3023212311002323"></a>

<a id="canonical-0312020311322330-2002322230201002-1200121223233021-3032202210313230-0030332231021232-3313130310203300-0221210333120003-2231010122323001"></a>

#### `stateful_service.containers.liveness_check.tcp_health_check.port.num` property

Type: `"number"`. Optional.

Port Number. Exclusive with \[name\] Port number.

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
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-2101233010110302-0023221233331300-0020012001202110-3111101033233313-3131302323332021-0313320323031313-1031100100130333-1131132213032321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.containers.readiness_check` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-018.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.containers](resources--workload--reference--group-028.md#canonical-3221212333032323-1011122002030020-0332100102213333-2230123311320001-3101132000200220-3323302320310311-1200332113313321-1231002100113301)
- stateful_service.containers.readiness_check

<a id="canonical-0323220310301030-3133212133211033-3332322000133113-2120302312100102-2312200300213201-2120220022300011-3202202202331300-1320200230223221"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
readiness_check {
  # Configure direct properties listed below.
}
```

<a id="canonical-1012003011110000-3030322122122011-3003300123011100-0123202120132222-2022021210202023-3101123002032122-0110320133133212-1031031201201302"></a>

### Direct properties for `stateful_service.containers.readiness_check`

- [exec_health_check](resources--workload--reference--group-029.md#canonical-2120310202200113-3301220212111331-0022001202232223-3112320131020100-3232021031113102-2200212230203103-2200201020131302-1323310100003320): complete subsection reference.

<a id="canonical-2200202110212210-0010130300032112-1300012231113002-2013220302232202-3203311311103032-0123033202012022-3330011321023321-0200002200020001"></a>

<a id="canonical-3231122303200333-1002213010231023-1010120213013130-3010133323000211-3020032202020120-1221211113230200-0030012132230022-2312220312132013"></a>

#### `stateful_service.containers.readiness_check.healthy_threshold` property

Type: `"number"`. Optional.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

- [http_health_check](resources--workload--reference--group-029.md#canonical-3103313332112223-2130033020321112-0322301232002211-2223322202123210-2331330121022122-1311321200133031-0301021310022213-3322311200230001): complete subsection reference.

<a id="canonical-2032323233022000-2111301312233212-3220202111202203-2321033120131000-3313310033332030-3311230232200223-0111123032232123-1111211103230331"></a>

<a id="canonical-2310211223211031-0201210302222320-1100303101002232-2100023222311133-2302322211222033-2201023031301231-0010312310133303-0310310021021113"></a>

#### `stateful_service.containers.readiness_check.initial_delay` property

Type: `"number"`. Optional.

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
    "ves.io.schema.rules.uint32.lte": "600"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "600"
  }
}
```

<a id="canonical-3111111033122312-3203113301120013-3031100222000030-1311010023233033-3111023131321220-2021023130023003-0211113012332210-0232102320332321"></a>

<a id="canonical-1013010000101333-1103120312232122-2210311011120003-0230130113331310-0102321213032312-2330233032121112-2130113023230233-0301101123030313"></a>

#### `stateful_service.containers.readiness_check.interval` property

Type: `"number"`. Optional.

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

- [tcp_health_check](resources--workload--reference--group-029.md#canonical-2300221210013310-2003113111031132-1030311021321222-3322010001000232-1230010311010331-2111102232333312-1002231210303301-0322323223330032): complete subsection reference.

<a id="canonical-1233010033201322-1122030301233232-0202011331311310-0211001302311300-0013301100330102-3230022213333220-0312333021121123-0322233333303313"></a>
