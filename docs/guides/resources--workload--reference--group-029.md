---
page_title: "xcsh_workload reference"
subcategory: "Container"
description: "Complete grouped canonical reference for xcsh_workload reference."
---

# xcsh_workload reference

<a id="canonical-1012003011110000-3030322122122011-3003300123011100-0123202120132222-2022021210202023-3101123002032122-0110320133133212-1031031201201302"></a>

## stateful_service.containers.readiness_check — readiness_check / 323122100101 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-017.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.containers](resources--workload--reference--group-028.md#canonical-3221212333032323-1011122002030020-0332100102213333-2230123311320001-3101132000200220-3323302320310311-1200332113313321-1231002100113301)
- stateful_service.containers.readiness_check

<a id="canonical-0323220310301030-3133212133211033-3332322000133113-2120302312100102-2312200300213201-2120220022300011-3202202202331300-1320200230223221"></a>

Type: `"object"`. single nested block, Optional.

HealthCheckType describes a health check to be performed against a container to determine whether it
has started up or is alive or ready to receive traffic.

Upstream description:

HealthCheckType describes a health check to be performed against a container to determine whether it
has started up or is alive or ready to receive traffic.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("healthy_threshold",
    "interval",
    "timeout",
    "unhealthy_threshold"),
  validators.ConflictingObjectAttributes("exec_health_check",
    "http_health_check"),
  validators.ConflictingObjectAttributes("exec_health_check",
    "tcp_health_check"),
  validators.ConflictingObjectAttributes("http_health_check",
    "tcp_health_check")}
```

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

<a id="canonical-3231122303200333-1002213010231023-1010120213013130-3010133323000211-3020032202020120-1221211113230200-0030012132230022-2312220312132013"></a>

## Direct properties — readiness_check / 323122100101 / 3

- [exec_health_check](resources--workload--reference--group-029.md#canonical-2120310202200113-3301220212111331-0022001202232223-3112320131020100-3232021031113102-2200212230203103-2200201020131302-1323310100003320): complete subsection reference.

<a id="canonical-2200202110212210-0010130300032112-1300012231113002-2013220302232202-3203311311103032-0123033202012022-3330011321023321-0200002200020001"></a>

<a id="canonical-2310211223211031-0201210302222320-1100303101002232-2100023222311133-2302322211222033-2201023031301231-0010312310133303-0310310021021113"></a>

## healthy_threshold property — readiness_check / 323122100101 / 4

Type: `"number"`. Optional.

Number of consecutive successful responses after having failed before declaring healthy. In other
words, this is the number of healthy health checks required before marking healthy. Note that during
startup and liveliness, only a single successful health check is required to mark a container..

Upstream description:

Number of consecutive successful responses after having failed before declaring healthy. In other
words, this is the number of healthy health checks required before marking healthy. Note that during
startup and liveliness, only a single successful health check is required to mark a container
healthy.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(1, 16),
}
```

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

- [http_health_check](resources--workload--reference--group-029.md#canonical-3103313332112223-2130033020321112-0322301232002211-2223322202123210-2331330121022122-1311321200133031-0301021310022213-3322311200230001): complete subsection reference.

<a id="canonical-2032323233022000-2111301312233212-3220202111202203-2321033120131000-3313310033332030-3311230232200223-0111123032232123-1111211103230331"></a>

<a id="canonical-1013010000101333-1103120312232122-2210311011120003-0230130113331310-0102321213032312-2330233032121112-2130113023230233-0301101123030313"></a>

## initial_delay property — readiness_check / 323122100101 / 5

Type: `"number"`. Optional.

Number of seconds after the container has started before health checks are initiated.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.AtMost(600),
}
```

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

<a id="canonical-3111111033122312-3203113301120013-3031100222000030-1311010023233033-3111023131321220-2021023130023003-0211113012332210-0232102320332321"></a>

<a id="canonical-0211321201021130-3203023122013313-0033111111202330-1002203231323230-0330322112200132-1201210313120310-3102022110231121-1023130023120321"></a>

## interval property — readiness_check / 323122100101 / 6

Type: `"number"`. Optional.

Time interval in seconds between two health check requests.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(1, 600),
}
```

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

- [tcp_health_check](resources--workload--reference--group-029.md#canonical-2300221210013310-2003113111031132-1030311021321222-3322010001000232-1230010311010331-2111102232333312-1002231210303301-0322323223330032): complete subsection reference.

<a id="canonical-1233010033201322-1122030301233232-0202011331311310-0211001302311300-0013301100330102-3230022213333220-0312333021121123-0322233333303313"></a>

<a id="canonical-3121303122302223-0002002213231003-0012030000132323-1312310322230330-2212101112212312-3013301020113032-3120122110311220-3021002121312002"></a>

## timeout property — readiness_check / 323122100101 / 7

Type: `"number"`. Optional.

Timeout in seconds to wait for successful response. In other words, it is the time to wait for a
health check response. If the timeout is reached the health check attempt will be considered a
failure.

Upstream description:

Timeout in seconds to wait for successful response. In other words, it is the time to wait for a
health check response. If the timeout is reached the health check attempt will be considered a
failure.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(1, 600),
}
```

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

<a id="canonical-3222202101002313-0221320333112330-2020022233313320-0200323033322023-0203233000103112-2313133133221033-0031220312013033-3232013200230132"></a>

<a id="canonical-3311223023302220-1111101210003300-2101101202322113-1303021232301001-1000220130202032-0221222010231132-2321200221303302-1322212310001013"></a>

## unhealthy_threshold property — readiness_check / 323122100101 / 8

Type: `"number"`. Optional.

Number of consecutive failed responses before declaring unhealthy. In other words, this is the
number of unhealthy health checks required before a container is marked unhealthy.

Upstream description:

Number of consecutive failed responses before declaring unhealthy. In other words, this is the
number of unhealthy health checks required before a container is marked unhealthy.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(1, 16),
}
```

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

<a id="canonical-2323111313003221-1321123213212213-1201120003211221-2310321232112330-3010302313330033-3133121332321303-1320223110311132-0213133123032103"></a>

## Next pages — readiness_check / 323122100101 / 9

- [stateful_service.containers.readiness_check.exec_health_check](resources--workload--reference--group-029.md#canonical-2120310202200113-3301220212111331-0022001202232223-3112320131020100-3232021031113102-2200212230203103-2200201020131302-1323310100003320)
- [stateful_service.containers.readiness_check.http_health_check](resources--workload--reference--group-029.md#canonical-3103313332112223-2130033020321112-0322301232002211-2223322202123210-2331330121022122-1311321200133031-0301021310022213-3322311200230001)
- [stateful_service.containers.readiness_check.tcp_health_check](resources--workload--reference--group-029.md#canonical-2300221210013310-2003113111031132-1030311021321222-3322010001000232-1230010311010331-2111102232333312-1002231210303301-0322323223330032)
- [stateful_service.containers](resources--workload--reference--group-028.md#canonical-3221212333032323-1011122002030020-0332100102213333-2230123311320001-3101132000200220-3323302320310311-1200332113313321-1231002100113301)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-2120310202200113-3301220212111331-0022001202232223-3112320131020100-3232021031113102-2200212230203103-2200201020131302-1323310100003320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1102102202022032-2121230101013321-1112333110132013-2303120110123332-0320003211000302-0012222033312301-0020002303020011-3210121323031300"></a>

## stateful_service.containers.readiness_check.exec_health_check — exec_health_check / 312121022231 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-017.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.containers](resources--workload--reference--group-028.md#canonical-3221212333032323-1011122002030020-0332100102213333-2230123311320001-3101132000200220-3323302320310311-1200332113313321-1231002100113301)
- [stateful_service.containers.readiness_check](resources--workload--reference--group-028.md#canonical-2101233010110302-0023221233331300-0020012001202110-3111101033233313-3131302323332021-0313320323031313-1031100100130333-1131132213032321)
- stateful_service.containers.readiness_check.exec_health_check

<a id="canonical-1021330120223030-2122321022222031-2220202113013110-2132123213112332-1222301211310121-3211102000203223-3100122121320031-2002020302130300"></a>

Type: `"object"`. single nested block, Optional.

ExecHealthCheckType describes a health check based on 'run in container' action. Exit status of 0 is
treated as live/healthy and non-zero is unhealthy.

Upstream description:

ExecHealthCheckType describes a health check based on "run in container" action. Exit status of 0 is
treated as live/healthy and non-zero is unhealthy.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("command")}
```

Receipt-pinned upstream constraints:

```json
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

<a id="canonical-0312013201210300-2002222213210213-0313122203311203-3032302202110033-2230213020103012-1231210302130232-2112212302130201-1203321033121022"></a>

## Direct properties — exec_health_check / 312121022231 / 3

<a id="canonical-1311101110021022-0030212232231232-1300320032323013-0011030312230013-1123031320001100-0022102101330330-1010222302212102-3132121021312211"></a>

<a id="canonical-3210221301002001-1123322101000023-3100100230321220-1103200333301103-3210121211221013-0300110203113103-1020131300131301-1310023312020011"></a>

## command property — exec_health_check / 312121022231 / 4

Type: `["list", "string"]`. Optional.

Command is the command line to execute inside the container, the working directory for the command
is root ('/') in the container's filesystem. The command is simply exec'd, it is not run inside a
shell, so traditional shell instructions ('|', etc) won't work. To use a shell, you need to..

Upstream description:

Command is the command line to execute inside the container, the working directory for the command
is root ('/') in the container's filesystem. The command is simply exec'd, it is not run inside a
shell, so traditional shell instructions ('|', etc) won't work. To use a shell, you need to
explicitly call out to that shell.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeAtMost(16),
}
```

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

<a id="canonical-3112203231030313-1011010323112223-3312201223123023-1233323323222003-3020210222322301-1103130302220333-1320223100213023-0203033332312311"></a>

## Next pages — exec_health_check / 312121022231 / 5

- [stateful_service.containers.readiness_check](resources--workload--reference--group-028.md#canonical-2101233010110302-0023221233331300-0020012001202110-3111101033233313-3131302323332021-0313320323031313-1031100100130333-1131132213032321)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-3103313332112223-2130033020321112-0322301232002211-2223322202123210-2331330121022122-1311321200133031-0301021310022213-3322311200230001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1331121022331023-2300233320011202-3330133013310130-3333300003203312-1232132230112230-0310020003311110-3230301233030032-2213102111232200"></a>

## stateful_service.containers.readiness_check.http_health_check — http_health_check / 011111022021 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-017.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.containers](resources--workload--reference--group-028.md#canonical-3221212333032323-1011122002030020-0332100102213333-2230123311320001-3101132000200220-3323302320310311-1200332113313321-1231002100113301)
- [stateful_service.containers.readiness_check](resources--workload--reference--group-028.md#canonical-2101233010110302-0023221233331300-0020012001202110-3111101033233313-3131302323332021-0313320323031313-1031100100130333-1131132213032321)
- stateful_service.containers.readiness_check.http_health_check

<a id="canonical-2003032130003110-0133333131320013-3310311320222313-3333131330211120-0023231130312103-0302021303013003-2132102130203133-0213303213023131"></a>

Type: `"object"`. single nested block, Optional.

HTTPHealthCheckType describes a health check based on HTTP GET requests.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("path")}
```

Receipt-pinned upstream constraints:

```json
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

<a id="canonical-2112330010010332-0222133121120232-1303211112312010-0021302332130333-1210221323313311-2320333001102112-0311001020301003-1312032132230223"></a>

## Direct properties — http_health_check / 011111022021 / 3

<a id="canonical-0331301032010221-1213233122200002-0203012323313201-1020312223233211-0002331320312201-2233212131030220-2013300313033212-1003212200220200"></a>

<a id="canonical-2330332233221032-1022113122303110-2231123211321223-3213033200321110-0202202102013231-3111302121133303-1232022133210311-1213330201322131"></a>

## headers property — http_health_check / 011111022021 / 4

Type: `["map", "string"]`. Optional.

Specifies a list of HTTP headers that should be added to each request that is sent to the health
checked container. This is a list of key-value pairs.

Upstream description:

Specifies a list of HTTP headers that should be added to each request that is sent to the health
checked container. This is a list of key-value pairs.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Map{validators.MapConstraintsValidator("{\"cardinality\":{\"maxProperties\":16},\"category\":\"discovery\",\"constraintType\":\"map\",\"deterministic\":true,\"keys\":{\"maxLength\":256,\"minLength\":1,\"type\":\"string\"},\"originalRules\":{\"ves.io.schema.rules.map.keys.string.max_len\":\"256\",\"ves.io.schema.rules.map.keys.string.min_len\":\"1\",\"ves.io.schema.rules.map.max_pairs\":\"16\",\"ves.io.schema.rules.map.values.string.max_len\":\"2048\",\"ves.io.schema.rules.map.values.string.min_len\":\"1\"},\"values\":{\"maxLength\":2048,\"minLength\":1,\"type\":\"string\"}}")}
```

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

<a id="canonical-0300100120103133-0311203130330020-3112003301012300-1203110331121101-0320003301131123-3222310032120033-0210112233232323-2233012103110320"></a>

<a id="canonical-2121131203232013-1323003211011202-2311131001031330-0123200132110232-2232023222301113-3302032111213023-3001320121303313-1010232010333202"></a>

## host_header property — http_health_check / 011111022021 / 5

Type: `"string"`. Optional.

The value of the host header in the HTTP health check request.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(262),
}
```

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

<a id="canonical-2112331031223133-3123002102002110-3321001131130010-1131002002110101-1030230113302111-3313023011112300-0230300312131330-0032131023013112"></a>

<a id="canonical-1230103300210001-3030030210300330-2303203233100332-0311011331020011-2121003232113033-2212312000100120-2111330223332332-3000000210201302"></a>

## path property — http_health_check / 011111022021 / 6

Type: `"string"`. Optional.

Path. Path to access on the HTTP server.

Upstream description:

Path to access on the HTTP server.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 2048),
}
```

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

- [port](resources--workload--reference--group-029.md#canonical-3231130010320201-0320033112203300-3112302023120132-2000032130130302-0101323200333230-1121220101133011-2332020330112110-0222103200100130): complete subsection reference.

<a id="canonical-0201300322320030-0310002312320020-1213320212223002-0220020221032331-3232212133333303-1021321122112131-2320210103021111-3211223203031220"></a>

## Next pages — http_health_check / 011111022021 / 7

- [stateful_service.containers.readiness_check.http_health_check.port](resources--workload--reference--group-029.md#canonical-3231130010320201-0320033112203300-3112302023120132-2000032130130302-0101323200333230-1121220101133011-2332020330112110-0222103200100130)
- [stateful_service.containers.readiness_check](resources--workload--reference--group-028.md#canonical-2101233010110302-0023221233331300-0020012001202110-3111101033233313-3131302323332021-0313320323031313-1031100100130333-1131132213032321)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-3231130010320201-0320033112203300-3112302023120132-2000032130130302-0101323200333230-1121220101133011-2332020330112110-0222103200100130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2030322033221210-2101012132010232-0031112031102020-0003013231312123-2231221300113331-2002031130133121-3030333301210130-0111202301303303"></a>

## stateful_service.containers.readiness_check.http_health_check.port — port / 211021210023 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-017.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.containers](resources--workload--reference--group-028.md#canonical-3221212333032323-1011122002030020-0332100102213333-2230123311320001-3101132000200220-3323302320310311-1200332113313321-1231002100113301)
- [stateful_service.containers.readiness_check](resources--workload--reference--group-028.md#canonical-2101233010110302-0023221233331300-0020012001202110-3111101033233313-3131302323332021-0313320323031313-1031100100130333-1131132213032321)
- [stateful_service.containers.readiness_check.http_health_check](resources--workload--reference--group-029.md#canonical-3103313332112223-2130033020321112-0322301232002211-2223322202123210-2331330121022122-1311321200133031-0301021310022213-3322311200230001)
- stateful_service.containers.readiness_check.http_health_check.port

<a id="canonical-0221112331321211-2220201111002102-1330010110320311-0010313111110001-3223332331100130-0211230333023300-3101311222031101-2221333132022123"></a>

Type: `"object"`. single nested block, Optional.

Port. Port

Upstream description:

Port

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("name",
    "num")}
```

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

<a id="canonical-3220102101313010-3030330031322133-1012310311020330-2233111120133130-2101022033032323-1321001110220221-0120210202030000-2032232011303113"></a>

## Direct properties — port / 211021210023 / 3

<a id="canonical-1302120101221101-1322011210031210-0102220213022222-3220002031033310-2313233331233121-2023000011011320-0202320133213211-0222333222202010"></a>

<a id="canonical-2023030230111322-3220311112110311-1001000001300003-0101332111332012-2320133310202021-2131023321133302-2112320323330000-3130332132303331"></a>

## name property — port / 211021210023 / 4

Type: `"string"`. Optional.

Port Name. Exclusive with \[num\] Port Name.

Upstream description:

Exclusive with \[num\] Port Name.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
  stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z]([-a-z0-9]*[a-z0-9])?$`),
    ""),
}
```

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

<a id="canonical-1212332103010031-0230002021032223-0203012010230032-3010212301210233-3300022012321113-3203212011030031-3320323030312100-0330322022232100"></a>

<a id="canonical-3302022201020131-2023030021032303-1130312222102232-1103301311130030-1132333033031123-1120133230022321-0003312300201322-0330113132212301"></a>

## num property — port / 211021210023 / 5

Type: `"number"`. Optional.

Port Number. Exclusive with \[name\] Port number.

Upstream description:

Exclusive with \[name\] Port number.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(1, 65535),
}
```

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

<a id="canonical-3301223030313213-0231332003113202-0322113113301012-1222212323021020-0211200011232202-0213111010312033-2130033001032300-1120232300312103"></a>

## Next pages — port / 211021210023 / 6

- [stateful_service.containers.readiness_check.http_health_check](resources--workload--reference--group-029.md#canonical-3103313332112223-2130033020321112-0322301232002211-2223322202123210-2331330121022122-1311321200133031-0301021310022213-3322311200230001)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-2300221210013310-2003113111031132-1030311021321222-3322010001000232-1230010311010331-2111102232333312-1002231210303301-0322323223330032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1312033201002130-1121320122301313-1032202103202033-3001113230201212-3020031123103002-0003110010030203-1012320033002002-1021213010233303"></a>

## stateful_service.containers.readiness_check.tcp_health_check — tcp_health_check / 031101323003 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-017.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.containers](resources--workload--reference--group-028.md#canonical-3221212333032323-1011122002030020-0332100102213333-2230123311320001-3101132000200220-3323302320310311-1200332113313321-1231002100113301)
- [stateful_service.containers.readiness_check](resources--workload--reference--group-028.md#canonical-2101233010110302-0023221233331300-0020012001202110-3111101033233313-3131302323332021-0313320323031313-1031100100130333-1131132213032321)
- stateful_service.containers.readiness_check.tcp_health_check

<a id="canonical-0010011112021013-1232333101202322-1102223121222300-2211300231023333-1330212330220333-0102100222211223-0033003212002022-0222333110332222"></a>

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

<a id="canonical-2332313311211233-2231210213223033-1321321130203201-2013123333220113-1130302230111112-3132302332031312-1322122121112320-3132302222133322"></a>

## Direct properties — tcp_health_check / 031101323003 / 3

- [port](resources--workload--reference--group-029.md#canonical-3120000133220120-0330123030002233-3310320201010121-0132113123322322-1011101202103210-1030201022132110-2020301001333320-1130031101121132): complete subsection reference.

<a id="canonical-0110031030230320-1010111211133322-1310203131130211-3101310003213330-3030103302011331-0130232122030123-1133200320302031-0311101102123201"></a>

## Next pages — tcp_health_check / 031101323003 / 4

- [stateful_service.containers.readiness_check.tcp_health_check.port](resources--workload--reference--group-029.md#canonical-3120000133220120-0330123030002233-3310320201010121-0132113123322322-1011101202103210-1030201022132110-2020301001333320-1130031101121132)
- [stateful_service.containers.readiness_check](resources--workload--reference--group-028.md#canonical-2101233010110302-0023221233331300-0020012001202110-3111101033233313-3131302323332021-0313320323031313-1031100100130333-1131132213032321)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-3120000133220120-0330123030002233-3310320201010121-0132113123322322-1011101202103210-1030201022132110-2020301001333320-1130031101121132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0200223221311022-2130320212113021-3323022331202130-3203001330103233-0301203133313331-0303323003033033-2301322231330122-0023132020032330"></a>

## stateful_service.containers.readiness_check.tcp_health_check.port — port / 303100302020 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-017.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.containers](resources--workload--reference--group-028.md#canonical-3221212333032323-1011122002030020-0332100102213333-2230123311320001-3101132000200220-3323302320310311-1200332113313321-1231002100113301)
- [stateful_service.containers.readiness_check](resources--workload--reference--group-028.md#canonical-2101233010110302-0023221233331300-0020012001202110-3111101033233313-3131302323332021-0313320323031313-1031100100130333-1131132213032321)
- [stateful_service.containers.readiness_check.tcp_health_check](resources--workload--reference--group-029.md#canonical-2300221210013310-2003113111031132-1030311021321222-3322010001000232-1230010311010331-2111102232333312-1002231210303301-0322323223330032)
- stateful_service.containers.readiness_check.tcp_health_check.port

<a id="canonical-1133020313013130-0110220310000331-3130231002233111-0301012321212022-1101033332220131-2001300031113133-2331200033033123-2013030310133211"></a>

Type: `"object"`. single nested block, Optional.

Port. Port

Upstream description:

Port

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("name",
    "num")}
```

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

<a id="canonical-2010202231322032-2103303331123330-1211301200223102-2213111201023301-3132131031102113-3220133331020102-0001022211120312-0310100313311000"></a>

## Direct properties — port / 303100302020 / 3

<a id="canonical-0333212302133010-2023133230000213-1002011122330101-0021211213030013-0123012100112101-1333002200231222-1033222322230330-1031210010332120"></a>

<a id="canonical-3020323123011221-1102123003113202-0111332221022030-3303222002232130-0113032321202030-3211003010013123-3032003323011020-0330302001301303"></a>

## name property — port / 303100302020 / 4

Type: `"string"`. Optional.

Port Name. Exclusive with \[num\] Port Name.

Upstream description:

Exclusive with \[num\] Port Name.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
  stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z]([-a-z0-9]*[a-z0-9])?$`),
    ""),
}
```

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

<a id="canonical-2120132120321302-2022223003203210-2302331320200122-3301203113222030-2013130010200020-3013003000332031-0301013201200202-1201132131310202"></a>

<a id="canonical-1013020103232303-3103213322223121-3321110233103023-3231033132100102-1000200210000223-1032330132013123-3230310000001310-2202230220122213"></a>

## num property — port / 303100302020 / 5

Type: `"number"`. Optional.

Port Number. Exclusive with \[name\] Port number.

Upstream description:

Exclusive with \[name\] Port number.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(1, 65535),
}
```

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

<a id="canonical-2321003013221303-2211002203111112-1200331113011211-1210200120101033-3132100002003211-1233110002330003-1010313232021013-2333322322320233"></a>

## Next pages — port / 303100302020 / 6

- [stateful_service.containers.readiness_check.tcp_health_check](resources--workload--reference--group-029.md#canonical-2300221210013310-2003113111031132-1030311021321222-3322010001000232-1230010311010331-2111102232333312-1002231210303301-0322323223330032)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-0211010322210330-0002130101323013-1320121223020011-3332231210010332-1331310201011323-2210300012032000-2333113200230221-2013202000203003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3120133122013320-2232132002333232-0112303133310100-1323210222003212-1130122001233313-1213110300301230-3003300130002021-2320331110102202"></a>

## stateful_service.deploy_options — deploy_options / 230100332333 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-017.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- stateful_service.deploy_options

<a id="canonical-3122221203230131-0223311201302130-3132322123022311-3201123113111001-2131133230231230-1200023310310112-1200331023201300-2121220133203102"></a>

Type: `"object"`. single nested block, Optional.

Deploy OPTIONS are used to configure the workload deployment OPTIONS.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("all_res",
    "default_virtual_sites"),
  validators.ConflictingObjectAttributes("all_res",
    "deploy_ce_sites"),
  validators.ConflictingObjectAttributes("all_res",
    "deploy_ce_virtual_sites"),
  validators.ConflictingObjectAttributes("all_res",
    "deploy_re_sites"),
  validators.ConflictingObjectAttributes("all_res",
    "deploy_re_virtual_sites"),
  validators.ConflictingObjectAttributes("default_virtual_sites",
    "deploy_ce_sites"),
  validators.ConflictingObjectAttributes("default_virtual_sites",
    "deploy_ce_virtual_sites"),
  validators.ConflictingObjectAttributes("default_virtual_sites",
    "deploy_re_sites"),
  validators.ConflictingObjectAttributes("default_virtual_sites",
    "deploy_re_virtual_sites"),
  validators.ConflictingObjectAttributes("deploy_ce_sites",
    "deploy_ce_virtual_sites"),
  validators.ConflictingObjectAttributes("deploy_ce_sites",
    "deploy_re_sites"),
  validators.ConflictingObjectAttributes("deploy_ce_sites",
    "deploy_re_virtual_sites"),
  validators.ConflictingObjectAttributes("deploy_ce_virtual_sites",
    "deploy_re_sites"),
  validators.ConflictingObjectAttributes("deploy_ce_virtual_sites",
    "deploy_re_virtual_sites"),
  validators.ConflictingObjectAttributes("deploy_re_sites",
    "deploy_re_virtual_sites")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-deploy_choice": "[\"all_res\",\"default_virtual_sites\",\"deploy_ce_sites\",\"deploy_ce_virtual_sites\",\"deploy_re_sites\",\"deploy_re_virtual_sites\"]"
}
```

Terraform syntax:

```terraform
deploy_options {
  # Configure direct properties listed below.
}
```

<a id="canonical-1311223132201030-1302213020032102-3230131102101330-2221302311100220-2030021222221333-1210333201020221-2212113210010133-3021231331220211"></a>

## Direct properties — deploy_options / 230100332333 / 3

- [all_res](resources--workload--reference--group-029.md#canonical-0333020122030013-2232102120330100-0001232112111203-1122121033013213-1130210113022032-2131213330302211-1310213210222213-0123001003100333): complete subsection reference.

- [default_virtual_sites](resources--workload--reference--group-029.md#canonical-2032123320222002-1201323103121311-3200302232301122-1220200031310213-2033010301031101-2300321332333230-1002301103313220-3123233111201303): complete subsection reference.

- [deploy_ce_sites](resources--workload--reference--group-029.md#canonical-3303011322113203-2110011010212132-3231211223311110-0233203331210301-0312101210021013-2233130301322323-0331112200310331-2211313013233200): complete subsection reference.

- [deploy_ce_virtual_sites](resources--workload--reference--group-029.md#canonical-3023302212111110-1101130302032313-1320112013332333-0123300303110312-0223110332333123-3100332131220102-3003203012220032-0023213120323130): complete subsection reference.

- [deploy_re_sites](resources--workload--reference--group-029.md#canonical-3022120320012232-3033131302323222-1323221000120213-0313233111030310-1332302100333111-0231322223320330-3021210112112203-3200003302312222): complete subsection reference.

- [deploy_re_virtual_sites](resources--workload--reference--group-029.md#canonical-0133000313021212-3311021101103013-2200323130111022-2212320032022221-2201211303220333-3332200210223031-0003300231210302-1000202011023033): complete subsection reference.

<a id="canonical-3231333233023201-2331132022332021-0000233102313232-1000022023002103-1111013101100333-0332032211320002-1333111302122100-1302323111031213"></a>

## Next pages — deploy_options / 230100332333 / 4

- [stateful_service.deploy_options.all_res](resources--workload--reference--group-029.md#canonical-0333020122030013-2232102120330100-0001232112111203-1122121033013213-1130210113022032-2131213330302211-1310213210222213-0123001003100333)
- [stateful_service.deploy_options.default_virtual_sites](resources--workload--reference--group-029.md#canonical-2032123320222002-1201323103121311-3200302232301122-1220200031310213-2033010301031101-2300321332333230-1002301103313220-3123233111201303)
- [stateful_service.deploy_options.deploy_ce_sites](resources--workload--reference--group-029.md#canonical-3303011322113203-2110011010212132-3231211223311110-0233203331210301-0312101210021013-2233130301322323-0331112200310331-2211313013233200)
- [stateful_service.deploy_options.deploy_ce_virtual_sites](resources--workload--reference--group-029.md#canonical-3023302212111110-1101130302032313-1320112013332333-0123300303110312-0223110332333123-3100332131220102-3003203012220032-0023213120323130)
- [stateful_service.deploy_options.deploy_re_sites](resources--workload--reference--group-029.md#canonical-3022120320012232-3033131302323222-1323221000120213-0313233111030310-1332302100333111-0231322223320330-3021210112112203-3200003302312222)
- [stateful_service.deploy_options.deploy_re_virtual_sites](resources--workload--reference--group-029.md#canonical-0133000313021212-3311021101103013-2200323130111022-2212320032022221-2201211303220333-3332200210223031-0003300231210302-1000202011023033)
- [stateful_service](resources--workload--reference--group-017.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-0333020122030013-2232102120330100-0001232112111203-1122121033013213-1130210113022032-2131213330302211-1310213210222213-0123001003100333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1111122030322310-2211200020111231-3131013103122113-0322010230320112-1122100230112313-2230222313110133-3313310221310213-3023332222301230"></a>

## stateful_service.deploy_options.all_res — all_res / 133033132232 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-017.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.deploy_options](resources--workload--reference--group-029.md#canonical-0211010322210330-0002130101323013-1320121223020011-3332231210010332-1331310201011323-2210300012032000-2333113200230221-2013202000203003)
- stateful_service.deploy_options.all_res

<a id="canonical-1231030111210213-1132300233032311-1232210323102313-1300202033131010-1013121310132011-0220302013101121-2102020321231111-3220011333220222"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
all_res = {}
```

<a id="canonical-1001121110101312-0033130223321311-3301030210022111-1110000100202232-0323023133223302-3131031000220302-1003001323331100-2120010221302211"></a>

## Direct properties — all_res / 133033132232 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1310323011003130-0210031130212020-2222122230123003-1001212333331231-1100210001101312-3232310333011033-3210212122202121-0113230213200010"></a>

## Next pages — all_res / 133033132232 / 4

- [stateful_service.deploy_options](resources--workload--reference--group-029.md#canonical-0211010322210330-0002130101323013-1320121223020011-3332231210010332-1331310201011323-2210300012032000-2333113200230221-2013202000203003)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-2032123320222002-1201323103121311-3200302232301122-1220200031310213-2033010301031101-2300321332333230-1002301103313220-3123233111201303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0023320130202222-0003111112301201-3231121011030121-2221232003311330-0031001023221230-3122231011032103-0112313320211331-2211301130012301"></a>

## stateful_service.deploy_options.default_virtual_sites — default_virtual_sites / 030113123310 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-017.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.deploy_options](resources--workload--reference--group-029.md#canonical-0211010322210330-0002130101323013-1320121223020011-3332231210010332-1331310201011323-2210300012032000-2333113200230221-2013202000203003)
- stateful_service.deploy_options.default_virtual_sites

<a id="canonical-0331222002031300-3030001312022302-3322030020121313-3323311101330321-0211031023110203-1313323032111000-0231221322102023-2213223301302003"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
default_virtual_sites = {}
```

<a id="canonical-3122221131031022-3220330110231133-1210323210002203-1231112313201303-2130233221120020-0023222111031211-2102220022300121-2130331010020133"></a>

## Direct properties — default_virtual_sites / 030113123310 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3020033331011312-1322102110110230-1210222313231232-0131132100120121-0030130003103303-3302200313210103-0213300221022121-3303210121321012"></a>

## Next pages — default_virtual_sites / 030113123310 / 4

- [stateful_service.deploy_options](resources--workload--reference--group-029.md#canonical-0211010322210330-0002130101323013-1320121223020011-3332231210010332-1331310201011323-2210300012032000-2333113200230221-2013202000203003)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-3303011322113203-2110011010212132-3231211223311110-0233203331210301-0312101210021013-2233130301322323-0331112200310331-2211313013233200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2300221322312301-0100132231233133-0300211210302321-3321113300330313-3112202102130222-1112232002103021-0202111132210332-3200130020311112"></a>

## stateful_service.deploy_options.deploy_ce_sites — deploy_ce_sites / 322213000112 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-017.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.deploy_options](resources--workload--reference--group-029.md#canonical-0211010322210330-0002130101323013-1320121223020011-3332231210010332-1331310201011323-2210300012032000-2333113200230221-2013202000203003)
- stateful_service.deploy_options.deploy_ce_sites

<a id="canonical-0013312010302011-3223033130232103-1311300012311300-3031120012113203-3020310120322111-3003101332332201-0012101023011031-3100030303002331"></a>

Type: `"object"`. single nested block, Optional.

Defines a way to deploy a workload on specific Customer sites.

Upstream description:

This defines a way to deploy a workload on specific Customer sites.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("site")}
```

Receipt-pinned upstream constraints:

```json
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
deploy_ce_sites {
  # Configure direct properties listed below.
}
```

<a id="canonical-0002022130203301-3211023320123320-3133012122130033-2323013230312212-1013313231210113-0021011032112211-2000322022031233-0113102223022322"></a>

## Direct properties — deploy_ce_sites / 322213000112 / 3

- [site](resources--workload--reference--group-029.md#canonical-2021320012100113-1310113320012122-2012222230031133-0223310211311010-0101213130000222-1322222111113133-3011220021320220-1212320322302320): complete subsection reference.

<a id="canonical-2100303120031101-1320122110332031-0303122113220103-0331120232101212-1130322011131123-0231202221313222-3310322022313122-2100101020100202"></a>

## Next pages — deploy_ce_sites / 322213000112 / 4

- [stateful_service.deploy_options.deploy_ce_sites.site](resources--workload--reference--group-029.md#canonical-2021320012100113-1310113320012122-2012222230031133-0223310211311010-0101213130000222-1322222111113133-3011220021320220-1212320322302320)
- [stateful_service.deploy_options](resources--workload--reference--group-029.md#canonical-0211010322210330-0002130101323013-1320121223020011-3332231210010332-1331310201011323-2210300012032000-2333113200230221-2013202000203003)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-2021320012100113-1310113320012122-2012222230031133-0223310211311010-0101213130000222-1322222111113133-3011220021320220-1212320322302320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1323101331331103-3122122110301331-0110312010303101-2331323031103301-3202220133010302-1111223032330203-1312001213230213-3231120103212332"></a>

## stateful_service.deploy_options.deploy_ce_sites.site — site / 103211332023 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-017.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.deploy_options](resources--workload--reference--group-029.md#canonical-0211010322210330-0002130101323013-1320121223020011-3332231210010332-1331310201011323-2210300012032000-2333113200230221-2013202000203003)
- [stateful_service.deploy_options.deploy_ce_sites](resources--workload--reference--group-029.md#canonical-3303011322113203-2110011010212132-3231211223311110-0233203331210301-0312101210021013-2233130301322323-0331112200310331-2211313013233200)
- stateful_service.deploy_options.deploy_ce_sites.site

<a id="canonical-2312211010330001-3102212221312220-3203132030321111-3332201230123331-2112130303200030-2003010031200230-1102101321322012-0312311012312202"></a>

Type: `"object"`. list nested block, Optional.

Which customer sites should this workload be deployed.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.RequiredListObjectAttributes("name")}
```

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
site {
  # Configure direct properties listed below.
}
```

<a id="canonical-0213213323123022-3232011212122300-0112323111033033-0202303023331201-1010213011123313-3113013321001011-2203023100301321-1013020001231130"></a>

## Direct properties — site / 103211332023 / 3

<a id="canonical-1002030321023330-0223031011221111-1313122332003203-2113312332121223-3312001231013303-0313210230233221-1131020020310320-3002202330123223"></a>

<a id="canonical-0213212301323100-0132012131032232-2021212323213013-2220003332013013-3203313311001310-2201232000011203-1333010232111230-2233232023332333"></a>

## name property — site / 103211332023 / 4

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 128),
}
```

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

<a id="canonical-3323321132330102-2001132212111200-2202013330331323-3211230231022301-2132020110201232-0132320023001202-1212110210011332-2002200302001230"></a>

<a id="canonical-1130012023310110-1323123221133032-1200200132102201-1111322003332213-3011322110303321-1332112202111031-0230111230233232-3232031303021102"></a>

## namespace property — site / 103211332023 / 5

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

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

<a id="canonical-3312233012023201-2001332031022302-0002121110021203-3213120033012110-2110223130311111-2301213020133331-3133320121330330-3112310012032131"></a>

<a id="canonical-3210321320022303-1233032313223102-2220021121110202-1300033002020233-1230310302010002-1012211002033233-0132133130213322-0330130312331210"></a>

## tenant property — site / 103211332023 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(64),
}
```

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

<a id="canonical-2111332021023330-1222002320122132-3300213332100002-2321111113233221-3003321123220002-0012021333133232-1300201002012033-3013101211030102"></a>

## Next pages — site / 103211332023 / 7

- [stateful_service.deploy_options.deploy_ce_sites](resources--workload--reference--group-029.md#canonical-3303011322113203-2110011010212132-3231211223311110-0233203331210301-0312101210021013-2233130301322323-0331112200310331-2211313013233200)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-3023302212111110-1101130302032313-1320112013332333-0123300303110312-0223110332333123-3100332131220102-3003203012220032-0023213120323130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0220132201202213-3000013301222322-0021330121022221-3103132000311001-0220300311000311-0331231102031033-0032231101112200-2033230301130200"></a>

## stateful_service.deploy_options.deploy_ce_virtual_sites — deploy_ce_virtual_sites / 030312032302 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-017.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.deploy_options](resources--workload--reference--group-029.md#canonical-0211010322210330-0002130101323013-1320121223020011-3332231210010332-1331310201011323-2210300012032000-2333113200230221-2013202000203003)
- stateful_service.deploy_options.deploy_ce_virtual_sites

<a id="canonical-2011331213312121-0023313020300321-3223213011100031-2122033123310200-2002313213303303-3122133210212332-0022131310030302-1221101231102201"></a>

Type: `"object"`. single nested block, Optional.

Defines a way to deploy a workload on specific Customer virtual sites.

Upstream description:

This defines a way to deploy a workload on specific Customer virtual sites.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("virtual_site")}
```

Receipt-pinned upstream constraints:

```json
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
deploy_ce_virtual_sites {
  # Configure direct properties listed below.
}
```

<a id="canonical-1201200330203202-0310012310020002-1200232330323030-2213112012220201-1131022213131123-2032112203112010-3212331312312101-1333200120203002"></a>

## Direct properties — deploy_ce_virtual_sites / 030312032302 / 3

- [virtual_site](resources--workload--reference--group-029.md#canonical-1301213132310112-0121131113111312-0020123310311103-3300101030002210-1013300131122001-0110012323112020-0201201302333123-3323302231203023): complete subsection reference.

<a id="canonical-2213200233322331-2110010123223220-0031202213310131-2331010320120313-1101003321121321-2101132211301320-0311020023321001-0312010010232303"></a>

## Next pages — deploy_ce_virtual_sites / 030312032302 / 4

- [stateful_service.deploy_options.deploy_ce_virtual_sites.virtual_site](resources--workload--reference--group-029.md#canonical-1301213132310112-0121131113111312-0020123310311103-3300101030002210-1013300131122001-0110012323112020-0201201302333123-3323302231203023)
- [stateful_service.deploy_options](resources--workload--reference--group-029.md#canonical-0211010322210330-0002130101323013-1320121223020011-3332231210010332-1331310201011323-2210300012032000-2333113200230221-2013202000203003)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-1301213132310112-0121131113111312-0020123310311103-3300101030002210-1013300131122001-0110012323112020-0201201302333123-3323302231203023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1221120031231022-3232001321131331-0332111310220111-1102210201012232-3133102130323221-2111302312233103-3132212302100320-2132120022331022"></a>

## stateful_service.deploy_options.deploy_ce_virtual_sites.virtual_site — virtual_site / 123120230313 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-017.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.deploy_options](resources--workload--reference--group-029.md#canonical-0211010322210330-0002130101323013-1320121223020011-3332231210010332-1331310201011323-2210300012032000-2333113200230221-2013202000203003)
- [stateful_service.deploy_options.deploy_ce_virtual_sites](resources--workload--reference--group-029.md#canonical-3023302212111110-1101130302032313-1320112013332333-0123300303110312-0223110332333123-3100332131220102-3003203012220032-0023213120323130)
- stateful_service.deploy_options.deploy_ce_virtual_sites.virtual_site

<a id="canonical-0022010321203302-2313222201312100-0210313103211312-2333331231322120-0120223021311221-0333201232032233-3211322101023012-3130023112333232"></a>

Type: `"object"`. list nested block, Optional.

Which customer virtual sites should this workload be deployed.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.RequiredListObjectAttributes("name")}
```

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
virtual_site {
  # Configure direct properties listed below.
}
```

<a id="canonical-1203230332230223-0312020100230021-0033311122310031-0220210100220100-1303332301230231-3102132203201202-2320331333011130-3202213030301322"></a>

## Direct properties — virtual_site / 123120230313 / 3

<a id="canonical-3212223233002203-1013223103020030-0131120132322021-2100200311102300-2202332330302032-2031001230330000-1000121202200310-1101332200130003"></a>

<a id="canonical-0120031132220202-0310231322111120-1133002123011103-0130323033103031-0212013121012132-2202103013311323-1120021222301331-1303121313332112"></a>

## name property — virtual_site / 123120230313 / 4

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 128),
}
```

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

<a id="canonical-3022211001130101-1023203030323330-0032100002011300-0321120133132323-3323031230001311-3321122033203002-0020021101110020-0022011211012023"></a>

<a id="canonical-1101321221022200-2020312110120121-3030120212000020-2023311202333013-0201222221001310-2032010033331311-2112020301330211-3332322120202331"></a>

## namespace property — virtual_site / 123120230313 / 5

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

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

<a id="canonical-3030330032232301-3333033032100131-1112302121233000-3302101023031132-0123122123000013-3302003001001330-2301300200023023-0121221300001121"></a>

<a id="canonical-2301300311022302-2310030312011123-3311223120331211-1003031203301130-2323221211322302-0313012203302311-1033122000311121-2000001300321112"></a>

## tenant property — virtual_site / 123120230313 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(64),
}
```

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

<a id="canonical-3303223112032030-2311021311303331-3221123111131222-3031311013132022-0113210231321321-1231223022321011-1120210013230222-2130100233012102"></a>

## Next pages — virtual_site / 123120230313 / 7

- [stateful_service.deploy_options.deploy_ce_virtual_sites](resources--workload--reference--group-029.md#canonical-3023302212111110-1101130302032313-1320112013332333-0123300303110312-0223110332333123-3100332131220102-3003203012220032-0023213120323130)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-3022120320012232-3033131302323222-1323221000120213-0313233111030310-1332302100333111-0231322223320330-3021210112112203-3200003302312222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3121233230103022-0202322031201323-1213000123131232-1032200322321113-3213032203111230-3022101230122131-2133113330132203-0303003202212321"></a>

## stateful_service.deploy_options.deploy_re_sites — deploy_re_sites / 202023223021 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-017.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.deploy_options](resources--workload--reference--group-029.md#canonical-0211010322210330-0002130101323013-1320121223020011-3332231210010332-1331310201011323-2210300012032000-2333113200230221-2013202000203003)
- stateful_service.deploy_options.deploy_re_sites

<a id="canonical-1121323023112301-1313303331133222-1102012203031332-1003231000130330-3331321113232132-2213102103333003-3201332323133301-3130023110130320"></a>

Type: `"object"`. single nested block, Optional.

Defines a way to deploy a workload on specific Regional Edge sites.

Upstream description:

This defines a way to deploy a workload on specific Regional Edge sites.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("site")}
```

Receipt-pinned upstream constraints:

```json
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
deploy_re_sites {
  # Configure direct properties listed below.
}
```

<a id="canonical-2002321111213223-3010203330030113-2220012130013001-3320123110213011-3301313020332030-0210223021331211-0330200333331000-1220110032023322"></a>

## Direct properties — deploy_re_sites / 202023223021 / 3

- [site](resources--workload--reference--group-029.md#canonical-2201021001332122-0003122023021320-1210003201132313-1300321012303103-1133301332013202-3232300301133221-0001011031130012-2002132321032321): complete subsection reference.

<a id="canonical-3302030102101333-1310012133230031-0303022211021203-3212303212111111-1110012110221232-3302120210311011-0230001231012113-1311113013222003"></a>

## Next pages — deploy_re_sites / 202023223021 / 4

- [stateful_service.deploy_options.deploy_re_sites.site](resources--workload--reference--group-029.md#canonical-2201021001332122-0003122023021320-1210003201132313-1300321012303103-1133301332013202-3232300301133221-0001011031130012-2002132321032321)
- [stateful_service.deploy_options](resources--workload--reference--group-029.md#canonical-0211010322210330-0002130101323013-1320121223020011-3332231210010332-1331310201011323-2210300012032000-2333113200230221-2013202000203003)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-2201021001332122-0003122023021320-1210003201132313-1300321012303103-1133301332013202-3232300301133221-0001011031130012-2002132321032321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2210031203023102-2202302131122030-1300112302021320-0201012001321012-2031001122322301-2113210212331232-2203022130101202-1103030131002110"></a>

## stateful_service.deploy_options.deploy_re_sites.site — site / 233331012211 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-017.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.deploy_options](resources--workload--reference--group-029.md#canonical-0211010322210330-0002130101323013-1320121223020011-3332231210010332-1331310201011323-2210300012032000-2333113200230221-2013202000203003)
- [stateful_service.deploy_options.deploy_re_sites](resources--workload--reference--group-029.md#canonical-3022120320012232-3033131302323222-1323221000120213-0313233111030310-1332302100333111-0231322223320330-3021210112112203-3200003302312222)
- stateful_service.deploy_options.deploy_re_sites.site

<a id="canonical-1330230122133230-2103220011023001-3330311311031011-1333103013020313-0332313021222103-3320031300303310-3000010223001010-2303322112322100"></a>

Type: `"object"`. list nested block, Optional.

Which regional edge sites should this workload be deployed.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.RequiredListObjectAttributes("name")}
```

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
site {
  # Configure direct properties listed below.
}
```

<a id="canonical-2113103031321213-1001013221200333-0002032321021223-1021313223330310-1231010021011232-1032313203300223-1202312103023310-1202203012023131"></a>

## Direct properties — site / 233331012211 / 3

<a id="canonical-3133110330322313-3003321300323123-3020303301010032-3313220302111200-1211133231323121-3333100030333032-3033033012122111-2023333003032301"></a>

<a id="canonical-2133312300130222-1112312122231001-0220302032022300-0123113230202220-0303313100113223-3102313013132330-1101103321021323-3023022221321220"></a>

## name property — site / 233331012211 / 4

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 128),
}
```

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

<a id="canonical-1131022213100210-3321232002311321-1202131000132110-1302321122120301-3032331001000101-0030200210033323-0332001302112302-0312213000213020"></a>

<a id="canonical-0303213203313130-1130123301232311-0312322033101110-3233222302231322-3132223302111300-1213130313233323-0231221213022322-3202100301201003"></a>

## namespace property — site / 233331012211 / 5

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

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

<a id="canonical-1330212230113131-2221311320123230-2312313301220111-2001302201121130-1122233202120330-3003202232001330-0111301331020313-0302132313132122"></a>

<a id="canonical-1131032102123013-1110223221211330-0230122132202011-2033222321110022-3312232120332313-0211330330033331-1310130033121123-0300113003033120"></a>

## tenant property — site / 233331012211 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(64),
}
```

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

<a id="canonical-1310021102303330-1120131002103301-0011030122211232-1111233110123231-1122120222321213-3132033013101203-2120333333203012-1120100030222031"></a>

## Next pages — site / 233331012211 / 7

- [stateful_service.deploy_options.deploy_re_sites](resources--workload--reference--group-029.md#canonical-3022120320012232-3033131302323222-1323221000120213-0313233111030310-1332302100333111-0231322223320330-3021210112112203-3200003302312222)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-0133000313021212-3311021101103013-2200323130111022-2212320032022221-2201211303220333-3332200210223031-0003300231210302-1000202011023033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2301230011202323-0322010202122202-3200223021233031-3232312333203102-0112031230013302-1201233303323021-0313002211001231-3101230232202100"></a>

## stateful_service.deploy_options.deploy_re_virtual_sites — deploy_re_virtual_sites / 111322321103 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-017.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.deploy_options](resources--workload--reference--group-029.md#canonical-0211010322210330-0002130101323013-1320121223020011-3332231210010332-1331310201011323-2210300012032000-2333113200230221-2013202000203003)
- stateful_service.deploy_options.deploy_re_virtual_sites

<a id="canonical-2332002033202301-1112011020220102-0303311002123322-3331122321203302-3230310000023332-1121130221311101-1320032113012112-3233113020311103"></a>

Type: `"object"`. single nested block, Optional.

Defines a way to deploy a workload on specific Regional Edge virtual sites.

Upstream description:

This defines a way to deploy a workload on specific Regional Edge virtual sites.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("virtual_site")}
```

Receipt-pinned upstream constraints:

```json
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
deploy_re_virtual_sites {
  # Configure direct properties listed below.
}
```

<a id="canonical-3222031023203010-0202233023223311-1133233233012133-2132003133200323-1220303323213012-2213232020100020-3231310022132312-3320330010023011"></a>

## Direct properties — deploy_re_virtual_sites / 111322321103 / 3

- [virtual_site](resources--workload--reference--group-029.md#canonical-3230002312003101-1031112020110220-0011330120231212-0023121323131203-1220121321302310-3032220301113331-2300022030320213-0023131202201322): complete subsection reference.

<a id="canonical-3020311331303003-1123002200120112-0212020321011321-1330130031320330-1023321320032301-0013311011230011-2320322230003101-3230023200021120"></a>

## Next pages — deploy_re_virtual_sites / 111322321103 / 4

- [stateful_service.deploy_options.deploy_re_virtual_sites.virtual_site](resources--workload--reference--group-029.md#canonical-3230002312003101-1031112020110220-0011330120231212-0023121323131203-1220121321302310-3032220301113331-2300022030320213-0023131202201322)
- [stateful_service.deploy_options](resources--workload--reference--group-029.md#canonical-0211010322210330-0002130101323013-1320121223020011-3332231210010332-1331310201011323-2210300012032000-2333113200230221-2013202000203003)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-3230002312003101-1031112020110220-0011330120231212-0023121323131203-1220121321302310-3032220301113331-2300022030320213-0023131202201322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1000202003302210-0012003311101213-3123123110200331-2103203311301330-1030211211030010-2111211112000013-0200302030213021-3223122331220010"></a>

## stateful_service.deploy_options.deploy_re_virtual_sites.virtual_site — virtual_site / 211130120210 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-017.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.deploy_options](resources--workload--reference--group-029.md#canonical-0211010322210330-0002130101323013-1320121223020011-3332231210010332-1331310201011323-2210300012032000-2333113200230221-2013202000203003)
- [stateful_service.deploy_options.deploy_re_virtual_sites](resources--workload--reference--group-029.md#canonical-0133000313021212-3311021101103013-2200323130111022-2212320032022221-2201211303220333-3332200210223031-0003300231210302-1000202011023033)
- stateful_service.deploy_options.deploy_re_virtual_sites.virtual_site

<a id="canonical-2312133300332320-2322312031331102-0030231000211233-3303113332030213-3230331313322103-2232130321331213-3132133302102332-3310133120000220"></a>

Type: `"object"`. list nested block, Optional.

Which regional edge virtual sites should this workload be deployed.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.RequiredListObjectAttributes("name")}
```

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
virtual_site {
  # Configure direct properties listed below.
}
```

<a id="canonical-2103021023313221-0000223202110132-0331032300333022-0112110103010231-3123012223223101-1233112201201320-3101202012122012-2231110201231031"></a>

## Direct properties — virtual_site / 211130120210 / 3

<a id="canonical-2033132330103303-0322200330201100-2202013301302303-1201202002222021-1021220302003023-2233312301223033-0102333120121202-2220110211231020"></a>

<a id="canonical-2110202323013031-2113010001223021-0130231032001233-0030202332000212-0101333313000002-0001333333202011-1220030333302310-2112010211222132"></a>

## name property — virtual_site / 211130120210 / 4

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 128),
}
```

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

<a id="canonical-0131312013123113-0300213120120213-2000302013320310-3221103232231002-0023202232101230-3333003033020102-0021220232231201-3102212002122031"></a>

<a id="canonical-3131212211210101-2212200130312331-1102033201233312-1201211220100330-1221022033033302-0231131311132323-2023223013133201-1132001211300110"></a>

## namespace property — virtual_site / 211130120210 / 5

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

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

<a id="canonical-3003000323322230-2121113133311100-0202113121133003-3022021113122223-3313122220201333-0230330233021022-2331202231013222-0301233203301023"></a>

<a id="canonical-3222013203311102-2131321202201100-0210033111130123-1123003001032320-2333332323233013-1310300023111223-2112312033110120-2111322132312000"></a>

## tenant property — virtual_site / 211130120210 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(64),
}
```

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

<a id="canonical-2203221201310230-0111313103103020-1223111013011231-3302331323310100-0332131001020020-1021223331310211-1322200103130130-2233113030232232"></a>

## Next pages — virtual_site / 211130120210 / 7

- [stateful_service.deploy_options.deploy_re_virtual_sites](resources--workload--reference--group-029.md#canonical-0133000313021212-3311021101103013-2200323130111022-2212320032022221-2201211303220333-3332200210223031-0003300231210302-1000202011023033)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-2202000332001022-3031133313102112-0003233201110202-2303003230303311-3002213310230220-0131313123131032-0010302220210011-2300030100331332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1323322331122200-3021101012032202-0121133011111101-0133213322000213-1210322020303311-1333310212231201-1221320302103023-3132222301220012"></a>

## stateful_service.persistent_volumes — persistent_volumes / 132002201310 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-017.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- stateful_service.persistent_volumes

<a id="canonical-2300133301120232-1323322232111300-3202011021331210-0121120231233002-2203333102130012-2033331313310100-3112223332031121-2220210100313233"></a>

Type: `"object"`. list nested block, Optional.

Persistent storage configuration for the service.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.RequiredListObjectAttributes("name")}
```

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

Terraform syntax:

```terraform
persistent_volumes {
  # Configure direct properties listed below.
}
```

<a id="canonical-3300203200313030-3223300303221020-3303100111301101-3300020113100303-2231001112103131-3121220201003321-1201232030121001-2120330110300311"></a>

## Direct properties — persistent_volumes / 132002201310 / 3

<a id="canonical-1333023200113321-0222312122011212-0311212102130103-1213203311101131-0213203330113023-3223103332102223-2231101313233212-2320001213323212"></a>

<a id="canonical-3201001310301213-0220110030030313-1130012321033333-2213101222323301-0231132113331211-0111101023200211-3102221012202211-0230232122312302"></a>

## name property — persistent_volumes / 132002201310 / 4

Type: `"string"`. Optional.

Name. Name of the volume.

Upstream description:

Name of the volume.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

Receipt-pinned upstream constraints:

```json
{
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
    "maxLength": 63,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1,
    "pattern": "^[a-z0-9]([-a-z0-9]*[a-z0-9])?$",
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
    "ves.io.schema.rules.string.dns_1123_label": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.dns_1123_label": "true"
  }
}
```

- [persistent_volume](resources--workload--reference--group-029.md#canonical-2112312000232122-2312223222320010-2033200313031001-1131110333101220-3321201023230022-2131001030110101-1110201233102113-0013122123223132): complete subsection reference.

<a id="canonical-0033222332132330-1120022322133200-1213332302210300-1300233230021120-0203012002021021-3300213311323221-3232023322220011-0212323202110332"></a>

## Next pages — persistent_volumes / 132002201310 / 5

- [stateful_service.persistent_volumes.persistent_volume](resources--workload--reference--group-029.md#canonical-2112312000232122-2312223222320010-2033200313031001-1131110333101220-3321201023230022-2131001030110101-1110201233102113-0013122123223132)
- [stateful_service](resources--workload--reference--group-017.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-2112312000232122-2312223222320010-2033200313031001-1131110333101220-3321201023230022-2131001030110101-1110201233102113-0013122123223132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1120022211010322-1000320312303133-3120213220122100-0112112003333131-2031103032000322-3022233011121221-2102021013011223-0332033003132021"></a>

## stateful_service.persistent_volumes.persistent_volume — persistent_volume / 221121100031 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-017.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.persistent_volumes](resources--workload--reference--group-029.md#canonical-2202000332001022-3031133313102112-0003233201110202-2303003230303311-3002213310230220-0131313123131032-0010302220210011-2300030100331332)
- stateful_service.persistent_volumes.persistent_volume

<a id="canonical-0212231030210300-2333322302200302-1301233113333001-0111233211213112-0300333322310203-1013133303022202-3301022001322103-1010132000310213"></a>

Type: `"object"`. single nested block, Optional.

Volume containing the Persistent Storage for the workload.

Receipt-pinned upstream constraints:

```json
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
persistent_volume {
  # Configure direct properties listed below.
}
```

<a id="canonical-2000000132210121-1331123221010132-0313221213232023-0220310202203310-2130103222003102-3210222222220313-2103233300133212-2223013032113000"></a>

## Direct properties — persistent_volume / 221121100031 / 3

- [mount](resources--workload--reference--group-029.md#canonical-3301113202313211-2012203311231321-3201220012310011-1100230203103300-0033001210103031-2233020202130120-3113231222213203-3131002202213222): complete subsection reference.

- [storage](resources--workload--reference--group-029.md#canonical-2303112111002111-0301211130322101-1303203032331132-1303120113203310-1303311030303232-2300030233330020-3201133032112123-2222323020330130): complete subsection reference.

<a id="canonical-3103133011230200-2031201301220013-2322222330132101-0310010200121301-3001123020322201-3121130000032221-3203221003221132-0033302032002201"></a>

## Next pages — persistent_volume / 221121100031 / 4

- [stateful_service.persistent_volumes.persistent_volume.mount](resources--workload--reference--group-029.md#canonical-3301113202313211-2012203311231321-3201220012310011-1100230203103300-0033001210103031-2233020202130120-3113231222213203-3131002202213222)
- [stateful_service.persistent_volumes.persistent_volume.storage](resources--workload--reference--group-029.md#canonical-2303112111002111-0301211130322101-1303203032331132-1303120113203310-1303311030303232-2300030233330020-3201133032112123-2222323020330130)
- [stateful_service.persistent_volumes](resources--workload--reference--group-029.md#canonical-2202000332001022-3031133313102112-0003233201110202-2303003230303311-3002213310230220-0131313123131032-0010302220210011-2300030100331332)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-3301113202313211-2012203311231321-3201220012310011-1100230203103300-0033001210103031-2233020202130120-3113231222213203-3131002202213222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0230122312230001-0310013223322332-3030330112030112-3033211323033332-1201230101003322-0002021321301320-1232321020213301-3022102102321311"></a>

## stateful_service.persistent_volumes.persistent_volume.mount — mount / 233021013131 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-017.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.persistent_volumes](resources--workload--reference--group-029.md#canonical-2202000332001022-3031133313102112-0003233201110202-2303003230303311-3002213310230220-0131313123131032-0010302220210011-2300030100331332)
- [stateful_service.persistent_volumes.persistent_volume](resources--workload--reference--group-029.md#canonical-2112312000232122-2312223222320010-2033200313031001-1131110333101220-3321201023230022-2131001030110101-1110201233102113-0013122123223132)
- stateful_service.persistent_volumes.persistent_volume.mount

<a id="canonical-1113321313013030-0230200030211203-2212102030310110-0213032333020021-0100233233312332-2121222230010122-1232022003211201-0300002202301222"></a>

Type: `"object"`. single nested block, Optional.

Volume mount describes how volume is mounted inside a workload.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("mount_path")}
```

Receipt-pinned upstream constraints:

```json
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

<a id="canonical-0102221001003113-2323002221000022-1120013210032213-1312022002333312-0031323333121301-2213011310031103-3313332131010111-0300331033202120"></a>

## Direct properties — mount / 233021013131 / 3

<a id="canonical-0213013230311212-3213221122221110-0321321023300013-0133103001020210-1012212111000212-3313000331131213-1311301222011121-3102013313000220"></a>

<a id="canonical-2030333313322012-2232221211221303-3322321230001021-2100002133002323-2120120101302012-1013222331303002-3031212121321311-1310011012332210"></a>

## mode property — mount / 233021013131 / 4

Type: `"string"`. Optional.

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

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["VOLUME_MOUNT_READ_ONLY","VOLUME_MOUNT_READ_WRITE"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
Validators: []validator.String{
  stringvalidator.OneOf("VOLUME_MOUNT_READ_ONLY",
    "VOLUME_MOUNT_READ_WRITE"),
}
```

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

<a id="canonical-1210223203312321-3112202210321131-3032322011012020-0113232010113121-1202012313323313-2113011130013030-3331202001132123-3330030030223200"></a>

<a id="canonical-3130031001001313-3313002012323131-0113213012322101-2223230310201113-0013331203312212-2312233332023022-2110322003313032-1011323111031310"></a>

## mount_path property — mount / 233021013131 / 5

Type: `"string"`. Optional.

Path within the workload container at which the volume should be mounted. Must not contain ':'.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

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

<a id="canonical-1313313120332123-0222121032203213-1020131102231010-3303133312302103-3222313302111220-1003312120011320-0113203130000231-0323002332132220"></a>

<a id="canonical-0323332003202232-0121033232201302-0213003112100231-3211222120033232-2211113201010010-3111002320112130-2323203203311230-3112230301032112"></a>

## sub_path property — mount / 233021013131 / 6

Type: `"string"`. Optional.

Path within the volume from which the workload's volume should be mounted. Defaults to '' (volume's
root).

Upstream description:

Path within the volume from which the workload's volume should be mounted. Defaults to "" (volume's
root).

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

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

<a id="canonical-1302303112333200-2210221200311013-1222122122010312-1302020232022101-1113010032031123-1110200010233320-1023223332302232-3312012100113120"></a>

## Next pages — mount / 233021013131 / 7

- [stateful_service.persistent_volumes.persistent_volume](resources--workload--reference--group-029.md#canonical-2112312000232122-2312223222320010-2033200313031001-1131110333101220-3321201023230022-2131001030110101-1110201233102113-0013122123223132)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-2303112111002111-0301211130322101-1303203032331132-1303120113203310-1303311030303232-2300030233330020-3201133032112123-2222323020330130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0331323300003122-3220131313213100-2112203023133023-2003201102201230-2212333310331331-1330201100202210-2233103021022202-0022023313330311"></a>

## stateful_service.persistent_volumes.persistent_volume.storage — storage / 030021220323 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-017.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.persistent_volumes](resources--workload--reference--group-029.md#canonical-2202000332001022-3031133313102112-0003233201110202-2303003230303311-3002213310230220-0131313123131032-0010302220210011-2300030100331332)
- [stateful_service.persistent_volumes.persistent_volume](resources--workload--reference--group-029.md#canonical-2112312000232122-2312223222320010-2033200313031001-1131110333101220-3321201023230022-2131001030110101-1110201233102113-0013122123223132)
- stateful_service.persistent_volumes.persistent_volume.storage

<a id="canonical-2013002300003120-1131022000322102-2320112100302323-3301320011331121-1000003202313332-3200212013231120-3103033130020230-3222213221212121"></a>

Type: `"object"`. single nested block, Optional.

Persistent storage configuration is used to configure Persistent Volume Claim (PVC).

Upstream description:

Persistent storage configuration is used to configure Persistent Volume Claim (PVC)

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("storage_size"),
  validators.ConflictingObjectAttributes("class_name",
    "default")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-class_name_choice": "[\"class_name\",\"default\"]"
}
```

Terraform syntax:

```terraform
storage {
  # Configure direct properties listed below.
}
```

<a id="canonical-0112210020013203-1233022033233002-3320212220331103-2323023233231021-3102002130210200-3213332300032323-2003310112111132-2120003200103010"></a>

## Direct properties — storage / 030021220323 / 3

<a id="canonical-2112201320020300-0202102331201021-0320032121021030-2101120301320230-3332023211223131-1223103333220100-3102022300322302-0030221330003113"></a>

<a id="canonical-2011110113221223-0200021300231012-2201320220332231-2231131332230320-2332321220311102-0001133223003200-1032213102310200-1233331013100201"></a>

## access_mode property — storage / 030021220323 / 4

Type: `"string"`. Optional.

\[Enum:
ACCESS\_MODE\_READ\_WRITE\_ONCE|ACCESS\_MODE\_READ\_WRITE\_MANY|ACCESS\_MODE\_READ\_ONLY\_MANY\]
Persistence storage access mode is used to configure access mode for persistent storage -
ACCESS\_MODE\_READ\_WRITE\_ONCE: Read Write Once Read Write Once is used to mount persistent storage
in read/write mode to exactly 1 host - ACCESS\_MODE\_READ\_WRITE\_MANY: Read Write Many Read Write
Many is used.. Possible values are \`ACCESS\_MODE\_READ\_WRITE\_ONCE\`,
\`ACCESS\_MODE\_READ\_WRITE\_MANY\`, \`ACCESS\_MODE\_READ\_ONLY\_MANY\`. Defaults to
\`ACCESS\_MODE\_READ\_WRITE\_ONCE\`.

Upstream description:

Persistence storage access mode is used to configure access mode for persistent storage

&#8203;- ACCESS\_MODE\_READ\_WRITE\_ONCE: Read Write Once

Read Write Once is used to mount persistent storage in read/write mode to exactly 1 host &#8203;-
ACCESS\_MODE\_READ\_WRITE\_MANY: Read Write Many

Read Write Many is used to mount persistent storage in read/write mode to many hosts &#8203;-
ACCESS\_MODE\_READ\_ONLY\_MANY: Read Only Many

Read Only Many is used to mount persistent storage in read-only mode to many hosts.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["ACCESS_MODE_READ_ONLY_MANY","ACCESS_MODE_READ_WRITE_MANY","ACCESS_MODE_READ_WRITE_ONCE"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
Validators: []validator.String{
  stringvalidator.OneOf("ACCESS_MODE_READ_WRITE_ONCE",
    "ACCESS_MODE_READ_WRITE_MANY",
    "ACCESS_MODE_READ_ONLY_MANY"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "ACCESS_MODE_READ_WRITE_ONCE",
  "enum": [
    "ACCESS_MODE_READ_WRITE_ONCE",
    "ACCESS_MODE_READ_WRITE_MANY",
    "ACCESS_MODE_READ_ONLY_MANY"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0030110210122000-3302320010133330-1130220120111300-2110001230320222-3033121130333322-2301203221302111-3212131332231231-2100013201303000"></a>

<a id="canonical-2332103113030221-0333211011031220-3033121323112001-2303122101301230-1021012022201101-1111221211110311-3232210113013000-2120302000101111"></a>

## class_name property — storage / 030021220323 / 5

Type: `"string"`. Optional.

Exclusive with \[default\] Use the specified class name.

Upstream description:

Exclusive with \[default\] Use the specified class name.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

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

- [default](resources--workload--reference--group-029.md#canonical-3020112000023313-0201001203000200-0022013031202121-3202000232331110-0132102331232122-1130233003202332-3123001201102003-0230030220210222): complete subsection reference.

<a id="canonical-0122300312303202-0033232213100333-1032123001213323-3032102132010322-1023333323200221-2002032032230123-1021021200021123-2103021023003102"></a>

<a id="canonical-3032001233333100-0213033313000133-0102313120120133-2320221232002122-1300113201001332-0003133300232031-2032012311322233-1000203311303110"></a>

## storage_size property — storage / 030021220323 / 6

Type: `"number"`. Optional.

Size (in GiB). Size in GiB of the persistent storage.

Upstream description:

Size in GiB of the persistent storage.

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
    "ves.io.schema.rules.double.gte": "0.004",
    "ves.io.schema.rules.double.lte": "256",
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.double.gte": "0.004",
    "ves.io.schema.rules.double.lte": "256",
    "ves.io.schema.rules.message.required": "true"
  }
}
```

<a id="canonical-2212012030301003-3023121332220003-1322233131202033-3033322220222332-3220201312203220-2131110323000130-0330103033231222-0232130033232202"></a>

## Next pages — storage / 030021220323 / 7

- [stateful_service.persistent_volumes.persistent_volume.storage.default](resources--workload--reference--group-029.md#canonical-3020112000023313-0201001203000200-0022013031202121-3202000232331110-0132102331232122-1130233003202332-3123001201102003-0230030220210222)
- [stateful_service.persistent_volumes.persistent_volume](resources--workload--reference--group-029.md#canonical-2112312000232122-2312223222320010-2033200313031001-1131110333101220-3321201023230022-2131001030110101-1110201233102113-0013122123223132)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-3020112000023313-0201001203000200-0022013031202121-3202000232331110-0132102331232122-1130233003202332-3123001201102003-0230030220210222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0220130022130103-1013323312311132-0121113222023232-2320312230322312-0131222230322320-3131211232132020-3303322113311231-2012231130201211"></a>

## stateful_service.persistent_volumes.persistent_volume.storage.default — default / 213020002010 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-017.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.persistent_volumes](resources--workload--reference--group-029.md#canonical-2202000332001022-3031133313102112-0003233201110202-2303003230303311-3002213310230220-0131313123131032-0010302220210011-2300030100331332)
- [stateful_service.persistent_volumes.persistent_volume](resources--workload--reference--group-029.md#canonical-2112312000232122-2312223222320010-2033200313031001-1131110333101220-3321201023230022-2131001030110101-1110201233102113-0013122123223132)
- [stateful_service.persistent_volumes.persistent_volume.storage](resources--workload--reference--group-029.md#canonical-2303112111002111-0301211130322101-1303203032331132-1303120113203310-1303311030303232-2300030233330020-3201133032112123-2222323020330130)
- stateful_service.persistent_volumes.persistent_volume.storage.default

<a id="canonical-1223223201110332-2030212031030312-0000211231302011-1212110030313000-1102212111130321-2300002030111220-3220311133001031-0203010330203032"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
default = {}
```

<a id="canonical-3020213222220233-3011123101122312-0302010313213030-2313130313002123-3012202333201331-0323312112010131-3303310123022333-1330312330312112"></a>

## Direct properties — default / 213020002010 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2302320221012123-2103203003130231-1311333033023130-3313231231100101-2030330230032132-3103123110320013-0011232001313020-3111312112303022"></a>

## Next pages — default / 213020002010 / 4

- [stateful_service.persistent_volumes.persistent_volume.storage](resources--workload--reference--group-029.md#canonical-2303112111002111-0301211130322101-1303203032331132-1303120113203310-1303311030303232-2300030233330020-3201133032112123-2222323020330130)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-0221221200321221-1133220231313302-3122113332002302-1303223333210013-3310313331033111-2320113332020010-2301111102010030-3110331212031230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1010000010013131-0021001120323202-3122202103103010-1101233001102100-1202233320011001-3233110210033002-0030232201333210-0302222132031111"></a>

## stateful_service.scale_to_zero — scale_to_zero / 303303322010 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-017.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- stateful_service.scale_to_zero

<a id="canonical-2112321232032003-2313003201313020-1130203001113131-3132312011310220-0222033330010103-3313103032203223-3012311321121300-2231003312001030"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for scale to zero.

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

Terraform syntax:

```terraform
scale_to_zero = {}
```

<a id="canonical-0001313320112122-3022122330321003-3132102131331021-3003103302022003-1320320123121032-3311111121311112-1312220303001101-3210200113231102"></a>

## Direct properties — scale_to_zero / 303303322010 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2002023231202000-0001121221123021-2020133200200223-2302033232002022-3322213212220300-3313310322022232-3100031303322320-0012102011312013"></a>

## Next pages — scale_to_zero / 303303322010 / 4

- [stateful_service](resources--workload--reference--group-017.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-0320321332330022-3200203220020212-3113213300013223-3331000302332100-2331130020201320-0110232101322132-3201120100233223-3122231121020213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1201011013013223-0131112230123203-1320231003021233-0110202001112210-1131000300111032-2201200130122121-0020322033020021-3230102223112010"></a>

## stateful_service.volumes — volumes / 220003202313 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-017.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- stateful_service.volumes

<a id="canonical-0002232221033022-0212210210213012-3302303030230011-0203222213210020-1332001131132322-0300102331113133-0010131212030111-0121322311213013"></a>

Type: `"object"`. list nested block, Optional.

Ephemeral Volumes. Ephemeral volumes for the service.

Upstream description:

Ephemeral volumes for the service.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.ConflictingListObjectAttributes("empty_dir",
    "host_path")}
```

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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16"
  }
}
```

Terraform syntax:

```terraform
volumes {
  # Configure direct properties listed below.
}
```

<a id="canonical-0313330210330220-1221020223002103-1031333132000010-0332213233000023-0320101311122021-0001023200212232-3311301130310311-1201020101112203"></a>

## Direct properties — volumes / 220003202313 / 3

- [empty_dir](resources--workload--reference--group-029.md#canonical-1233111023320131-2313133001020012-3013001223020200-0301013301200313-0310102301020311-1212003232313031-1122230103210021-0203223003200300): complete subsection reference.

- [host_path](resources--workload--reference--group-029.md#canonical-0200113013001232-2100121231233113-3010302211030210-3130201033310001-2300331110310102-0333102232201120-1110200121203133-2221202001022121): complete subsection reference.

<a id="canonical-3213100232213233-3330112323320111-0112013223123002-2020230030103211-2331310101310233-0310002023220113-3020221231011231-3211132232330232"></a>

<a id="canonical-3230300130332020-3023010102033210-3223023110210223-1300100312200212-2113031302303320-2231223010210210-1010303322331232-0311312300311122"></a>

## name property — volumes / 220003202313 / 4

Type: `"string"`. Optional.

Name. Name of the volume.

Upstream description:

Name of the volume.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

Receipt-pinned upstream constraints:

```json
{
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
    "maxLength": 63,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1,
    "pattern": "^[a-z0-9]([-a-z0-9]*[a-z0-9])?$",
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
    "ves.io.schema.rules.string.dns_1123_label": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.dns_1123_label": "true"
  }
}
```

<a id="canonical-0312120133033030-3021121123200320-0113222012100031-2112323031213000-2311011111103130-2002232031132010-0203120332030011-2302211123332301"></a>

## Next pages — volumes / 220003202313 / 5

- [stateful_service.volumes.empty_dir](resources--workload--reference--group-029.md#canonical-1233111023320131-2313133001020012-3013001223020200-0301013301200313-0310102301020311-1212003232313031-1122230103210021-0203223003200300)
- [stateful_service.volumes.host_path](resources--workload--reference--group-029.md#canonical-0200113013001232-2100121231233113-3010302211030210-3130201033310001-2300331110310102-0333102232201120-1110200121203133-2221202001022121)
- [stateful_service](resources--workload--reference--group-017.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-1233111023320131-2313133001020012-3013001223020200-0301013301200313-0310102301020311-1212003232313031-1122230103210021-0203223003200300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2213210122020111-3301013230322020-3323033213120220-0002130012313202-0331203113023321-1201110332022031-3301330023333102-0231311103303100"></a>

## stateful_service.volumes.empty_dir — empty_dir / 322003011323 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-017.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.volumes](resources--workload--reference--group-029.md#canonical-0320321332330022-3200203220020212-3113213300013223-3331000302332100-2331130020201320-0110232101322132-3201120100233223-3122231121020213)
- stateful_service.volumes.empty_dir

<a id="canonical-2031330010200212-1123130110322133-0201323121030122-1300223121030212-0231320213021320-2222323110112023-2002330100313021-2133221121300003"></a>

Type: `"object"`. single nested block, Optional.

Volume containing a temporary directory whose lifetime is the same as a replica of a workload.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("size_limit")}
```

Receipt-pinned upstream constraints:

```json
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
empty_dir {
  # Configure direct properties listed below.
}
```

<a id="canonical-3002230132330213-3021230113202001-0231100103331101-3300220032233333-1213221332013132-3032003012232233-1202110100013002-2002211012100130"></a>

## Direct properties — empty_dir / 322003011323 / 3

- [mount](resources--workload--reference--group-029.md#canonical-1111230202033002-3211231323102023-2300210222020100-3122322030122000-1323103000020023-2112301330112331-0111222230222030-2131122133111121): complete subsection reference.

<a id="canonical-2331133302003133-0230333122222000-2232101011000001-0033220231103133-1003122032010013-1301130130221312-2133103131310211-3010131202122231"></a>

<a id="canonical-0312310012010221-2321300332130110-2031022320021010-2201123112021200-0011332212201130-3031002003200310-2203031103033303-1032330313230222"></a>

## size_limit property — empty_dir / 322003011323 / 4

Type: `"number"`. Optional.

Size Limit (in GiB). Configuration parameter for size limit

Upstream description:

Configuration parameter for size limit

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
    "ves.io.schema.rules.double.lte": "10",
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.double.lte": "10",
    "ves.io.schema.rules.message.required": "true"
  }
}
```

<a id="canonical-3022132012030021-1133002210321303-3103102333103030-0120122113123201-1222022233133222-2320322120011313-0022001003232110-2030130202322113"></a>

## Next pages — empty_dir / 322003011323 / 5

- [stateful_service.volumes.empty_dir.mount](resources--workload--reference--group-029.md#canonical-1111230202033002-3211231323102023-2300210222020100-3122322030122000-1323103000020023-2112301330112331-0111222230222030-2131122133111121)
- [stateful_service.volumes](resources--workload--reference--group-029.md#canonical-0320321332330022-3200203220020212-3113213300013223-3331000302332100-2331130020201320-0110232101322132-3201120100233223-3122231121020213)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-1111230202033002-3211231323102023-2300210222020100-3122322030122000-1323103000020023-2112301330112331-0111222230222030-2131122133111121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0013231000013333-2201321330233222-3210320311103032-0321230231132110-0302112221302302-3110032300021320-1211231311001221-2231310221100223"></a>

## stateful_service.volumes.empty_dir.mount — mount / 032121031331 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-017.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.volumes](resources--workload--reference--group-029.md#canonical-0320321332330022-3200203220020212-3113213300013223-3331000302332100-2331130020201320-0110232101322132-3201120100233223-3122231121020213)
- [stateful_service.volumes.empty_dir](resources--workload--reference--group-029.md#canonical-1233111023320131-2313133001020012-3013001223020200-0301013301200313-0310102301020311-1212003232313031-1122230103210021-0203223003200300)
- stateful_service.volumes.empty_dir.mount

<a id="canonical-2331130330102033-3110301310010302-0300300110133032-0332123302101320-1331312132202123-1233102132312012-3103313101323100-1020010033212233"></a>

Type: `"object"`. single nested block, Optional.

Volume mount describes how volume is mounted inside a workload.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("mount_path")}
```

Receipt-pinned upstream constraints:

```json
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

<a id="canonical-0301212321230230-3101112333231121-0321101231321222-3132323013323010-3311021203013002-2010300123201203-3210033200330003-1312320130130131"></a>

## Direct properties — mount / 032121031331 / 3

<a id="canonical-2213131202232003-2201330321333001-2212330020303022-2332130033202313-2322020222322112-0102100121110033-1010122200330320-0021231231332013"></a>

<a id="canonical-3100211222311230-3022330031221023-3212010100133232-1310000010031330-3123122223200313-2211103211020230-0011111211330100-3232213000133130"></a>

## mode property — mount / 032121031331 / 4

Type: `"string"`. Optional.

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

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["VOLUME_MOUNT_READ_ONLY","VOLUME_MOUNT_READ_WRITE"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
Validators: []validator.String{
  stringvalidator.OneOf("VOLUME_MOUNT_READ_ONLY",
    "VOLUME_MOUNT_READ_WRITE"),
}
```

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

<a id="canonical-3110033313120110-3113200223010133-3233233003121011-3132131323303310-3013131132331000-3222300032223032-3122311221233220-1021230200303113"></a>

<a id="canonical-1112132023123112-3301122113013333-1300202123020032-3203201332120110-0212010112100110-2203001331102213-0310303312213213-3011003021300102"></a>

## mount_path property — mount / 032121031331 / 5

Type: `"string"`. Optional.

Path within the workload container at which the volume should be mounted. Must not contain ':'.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

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

<a id="canonical-3132123100030132-1133222330320112-3022122010212231-1313102310023231-0111322213010102-3031311132231122-1121130131233332-2332102021033220"></a>

<a id="canonical-2112202132230302-2002312010123330-1023310212123213-0120211102023321-2311130312010320-3121313022333010-1113302330022023-1322321222023212"></a>

## sub_path property — mount / 032121031331 / 6

Type: `"string"`. Optional.

Path within the volume from which the workload's volume should be mounted. Defaults to '' (volume's
root).

Upstream description:

Path within the volume from which the workload's volume should be mounted. Defaults to "" (volume's
root).

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

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

<a id="canonical-0321222223111231-1131231122101120-1311212003000313-0310023231302000-3021102223000133-3303123211220002-0210200332132300-3011012220211213"></a>

## Next pages — mount / 032121031331 / 7

- [stateful_service.volumes.empty_dir](resources--workload--reference--group-029.md#canonical-1233111023320131-2313133001020012-3013001223020200-0301013301200313-0310102301020311-1212003232313031-1122230103210021-0203223003200300)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-0200113013001232-2100121231233113-3010302211030210-3130201033310001-2300331110310102-0333102232201120-1110200121203133-2221202001022121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3331202001302120-3132022202222012-2333000100311320-2030220122212113-2232131223331123-1020122133332210-3312210011320011-0212211211020203"></a>

## stateful_service.volumes.host_path — host_path / 233120320321 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-017.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.volumes](resources--workload--reference--group-029.md#canonical-0320321332330022-3200203220020212-3113213300013223-3331000302332100-2331130020201320-0110232101322132-3201120100233223-3122231121020213)
- stateful_service.volumes.host_path

<a id="canonical-3123321323021111-2201130322030323-0333300012120123-2122331030002013-1212301023132122-1222230332030032-3332312023233302-1210212003332113"></a>

Type: `"object"`. single nested block, Optional.

Volume containing a host mapped path into the workload.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("path")}
```

Receipt-pinned upstream constraints:

```json
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
host_path {
  # Configure direct properties listed below.
}
```

<a id="canonical-2300130012222313-3203301002111100-1031130331301003-3103122311333331-1320033310310121-3011311322330000-3003120202001231-1322001303000111"></a>

## Direct properties — host_path / 233120320321 / 3

- [mount](resources--workload--reference--group-029.md#canonical-3010313323033132-1302103320223112-0223313212330113-3121033022012103-1111001301030330-1100013113123203-1120311130233031-3130310203013032): complete subsection reference.

<a id="canonical-1231330100211020-2031301230120331-2032202320111303-2001123311223022-2101133232232113-2110022200010112-1322222011333300-2203220001002121"></a>

<a id="canonical-2023223130103021-0210231221122020-2223030230032003-2302012200233321-3221103213123100-0320332010220123-2221323120200230-2133202011210020"></a>

## path property — host_path / 233120320321 / 4

Type: `"string"`. Optional.

Path. Path of the directory on the host.

Upstream description:

Path of the directory on the host.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

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
    "pattern": "[^\\\\0]+"
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
    "ves.io.schema.rules.string.pattern": "[^\\\\0]+"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.pattern": "[^\\\\0]+"
  }
}
```

<a id="canonical-0221230131133333-3101130100332111-3031220311223211-0313212131303233-2231210311110112-0022232331213330-3003132222033020-2013121232201123"></a>

## Next pages — host_path / 233120320321 / 5

- [stateful_service.volumes.host_path.mount](resources--workload--reference--group-029.md#canonical-3010313323033132-1302103320223112-0223313212330113-3121033022012103-1111001301030330-1100013113123203-1120311130233031-3130310203013032)
- [stateful_service.volumes](resources--workload--reference--group-029.md#canonical-0320321332330022-3200203220020212-3113213300013223-3331000302332100-2331130020201320-0110232101322132-3201120100233223-3122231121020213)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-3010313323033132-1302103320223112-0223313212330113-3121033022012103-1111001301030330-1100013113123203-1120311130233031-3130310203013032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0312222203102132-0111122212211120-2011131113220013-3213311221012301-3132000003101023-0012200233313300-0030130200222323-2233010133000101"></a>

## stateful_service.volumes.host_path.mount — mount / 112122001322 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-017.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.volumes](resources--workload--reference--group-029.md#canonical-0320321332330022-3200203220020212-3113213300013223-3331000302332100-2331130020201320-0110232101322132-3201120100233223-3122231121020213)
- [stateful_service.volumes.host_path](resources--workload--reference--group-029.md#canonical-0200113013001232-2100121231233113-3010302211030210-3130201033310001-2300331110310102-0333102232201120-1110200121203133-2221202001022121)
- stateful_service.volumes.host_path.mount

<a id="canonical-1020120030201301-3123331213303303-1011301102201002-2220003020213303-1230301333322132-1320033302321102-2210333103122213-0230100120111332"></a>

Type: `"object"`. single nested block, Optional.

Volume mount describes how volume is mounted inside a workload.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("mount_path")}
```

Receipt-pinned upstream constraints:

```json
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

<a id="canonical-3100221023113110-3010220012113121-1300310231323133-0210230302312210-1233121012232322-0331031103200232-3233302311211233-3200330133300203"></a>

## Direct properties — mount / 112122001322 / 3

<a id="canonical-3012222112101000-0210322332003200-0300102210301331-2110021000110302-3113111210102103-0202320033211003-3023133023200201-3301113333230022"></a>

<a id="canonical-0300112333311331-2123030321123012-0112203232002301-3122120211033013-1211130123303211-1221001301002031-0100210310021223-1023322220022330"></a>

## mode property — mount / 112122001322 / 4

Type: `"string"`. Optional.

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

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["VOLUME_MOUNT_READ_ONLY","VOLUME_MOUNT_READ_WRITE"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
Validators: []validator.String{
  stringvalidator.OneOf("VOLUME_MOUNT_READ_ONLY",
    "VOLUME_MOUNT_READ_WRITE"),
}
```

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

<a id="canonical-0133131010230000-0100020011011002-2320103231002002-2033022203233030-1002131123012300-1311102330201210-0001032012020310-2210013010013320"></a>

<a id="canonical-3023110302000320-2220010302102133-3122222103233233-3120302330120123-3200203020201132-0312332110131223-2030312123103223-3320103012230311"></a>

## mount_path property — mount / 112122001322 / 5

Type: `"string"`. Optional.

Path within the workload container at which the volume should be mounted. Must not contain ':'.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

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

<a id="canonical-3210123120331222-0320333030200102-2211120300303202-2300231320131220-1131210303201221-2013213002023210-2113110233033132-1223210333033301"></a>

<a id="canonical-3103011313220033-1021203223202120-1302032100223321-3021210120033110-1200120200320120-1100200022323112-1211211110001210-1012330100321030"></a>

## sub_path property — mount / 112122001322 / 6

Type: `"string"`. Optional.

Path within the volume from which the workload's volume should be mounted. Defaults to '' (volume's
root).

Upstream description:

Path within the volume from which the workload's volume should be mounted. Defaults to "" (volume's
root).

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

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

<a id="canonical-0101130112031323-1131110122301011-1223313211102211-3211101201321000-3231121010222000-0012113311112302-2313311330222111-3230331101102110"></a>

## Next pages — mount / 112122001322 / 7

- [stateful_service.volumes.host_path](resources--workload--reference--group-029.md#canonical-0200113013001232-2100121231233113-3010302211030210-3130201033310001-2300331110310102-0333102232201120-1110200121203133-2221202001022121)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-2102032211323123-0023120212303332-1230101130231233-1131123111320033-3302303022232001-0303333332310230-0230313031222332-2031033213122120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0030322213032031-2321233202200113-1220203303111220-3012100211100203-3022231120113200-1201000212022013-3131310210022302-3223332112000333"></a>

## timeouts — timeouts / 030302230013 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- timeouts

<a id="canonical-3231210002232200-0311023001013021-3300012201120001-0031223311331310-3010101110311202-3103130211030202-0130311220031102-1313031132033322"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-3312122103333000-2022031030111310-1210231231131130-1133300100220131-1121213332110100-0032012310021301-1113202202120103-2331312033210322"></a>

## Direct properties — timeouts / 030302230013 / 3

<a id="canonical-3303312302203131-0202332022333013-2110232232313220-2022313333213033-0020213101102022-0300030302130030-2023301103212002-2113212132123221"></a>

<a id="canonical-1132013200130102-2133032333121121-1203300203000032-1332203301222311-0211321221012121-3130100020300200-2020011323302302-0001013301313231"></a>

## create property — timeouts / 030302230013 / 4

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-0001120312210313-2002331131200112-0212121202311020-0113131223112311-0231101033113330-2211312100320102-2220002300213232-0030030221221122"></a>

<a id="canonical-0330230012123020-3110200023013113-0113122210121200-2213331212103212-0002010312020011-3002333203202123-2000133223310121-3211122332201113"></a>

## delete property — timeouts / 030302230013 / 5

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-2312023012301213-1120033230002133-3030331122001101-2221101330332011-0321000121001323-1331110123113313-1133001221030213-3302222131231232"></a>

<a id="canonical-2210031101001331-0031303301333100-2102210122303330-3121213201330330-2300011012230331-3312303230103102-0011131130332331-2311022222312302"></a>

## read property — timeouts / 030302230013 / 6

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-0003201223320203-3023131033333200-1131313330200203-0210120331201113-0211131020013000-0012001203313012-3211110022021302-1112121220311331"></a>

<a id="canonical-3121310223023202-2112001010132122-1102313032130120-0113220103323032-3113330131020010-3232121321111010-2100110222011103-1120030013100233"></a>

## update property — timeouts / 030302230013 / 7

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-2212221332021200-3113122103033312-0221022003002203-3233123103331132-1010130021331302-3312300022112003-2231013021320331-2002002303201332"></a>

## Next pages — timeouts / 030302230013 / 8

- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
