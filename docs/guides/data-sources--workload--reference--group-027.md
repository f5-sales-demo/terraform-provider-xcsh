---
page_title: "xcsh_workload reference"
subcategory: "Container"
description: "Complete grouped canonical reference for xcsh_workload reference."
---

# xcsh_workload reference

<a id="canonical-1102032003123222-0300302031311133-2112000011300220-3022000312022212-1133120031222301-2301022113101013-3023002323033022-0103223012031213"></a>

## `stateful_service.containers.image.container_registry.namespace` property

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

<a id="canonical-1322211111023111-2120010120210012-2233331210211312-0132113321001011-1233211030210103-2011232102002211-1300201030320031-0220230130332300"></a>

<a id="canonical-1013030310013130-1222322023023200-3222210200313002-0301222203101021-1311201303332132-3002213011003210-0130101013303223-3332212100222113"></a>

## `stateful_service.containers.image.container_registry.tenant` property

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

<a id="canonical-3103202230232111-3001333120213010-1133130321013002-2223022030303023-1133301012110330-0133320232312002-0022120311113300-1323331231023030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.containers.image.public` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.containers](data-sources--workload--reference--group-026.md#canonical-3101311130033212-3010211101012322-2332300122313021-2300111123200012-3222222201300313-2322010011032230-2110031201133132-3103131301302003)
- [stateful_service.containers.image](data-sources--workload--reference--group-026.md#canonical-0021330333100321-2210032021313101-0103302100233222-1221130213013113-2331021231110011-0122333011022231-0333012303023310-3330131303211312)
- stateful_service.containers.image.public

<a id="canonical-0031232113313331-1133312323102101-3301132002110310-0111011100301013-2322303312002330-2120101210303203-1231023230120322-2321112131221031"></a>

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

<a id="canonical-1000133100210330-0101332000310233-1022300301200130-1202221333012211-3002300330030301-2323010320303223-3322100021113321-2121321202120322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.containers.liveness_check` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.containers](data-sources--workload--reference--group-026.md#canonical-3101311130033212-3010211101012322-2332300122313021-2300111123200012-3222222201300313-2322010011032230-2110031201133132-3103131301302003)
- stateful_service.containers.liveness_check

<a id="canonical-3320031302030110-3222201310013100-2321033132201101-3320000223000303-1302201200211221-1032002233132333-2231322213211123-0000023302130130"></a>

Type: `"single"`. Computed.

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

<a id="canonical-3101203330021303-2012300202303131-3032123100110133-3113113303013202-0222023331002221-0333020301103021-0330100001003123-0203232211003032"></a>

### Direct properties for `stateful_service.containers.liveness_check`

- [exec_health_check](data-sources--workload--reference--group-027.md#canonical-1132220311013222-1220302112001030-2122320100000223-0201323003102123-3132111333220202-0222233122023211-3030003230120131-2102013312221301): complete subsection reference.

<a id="canonical-3010002322323130-0023023103100303-3031300310311020-3032123023300111-0311120232012133-0310211231012232-3311100200123302-3313212303012211"></a>

<a id="canonical-2122010303321100-3032332232221223-2210012222230022-0210222332031331-3320101330220120-0033310011011310-3033303320000022-0001321311100102"></a>

#### `stateful_service.containers.liveness_check.healthy_threshold` property

Type: `"number"`. Computed.

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

- [http_health_check](data-sources--workload--reference--group-027.md#canonical-0100221312120310-1003030310020211-3332103202011122-2000031211231222-2133023310103003-1300320030132231-2122022123333022-3201320212013312): complete subsection reference.

<a id="canonical-1131220033223321-3033102020111021-1122330311200002-2323231333020210-3002232100333321-2003302131022221-3220221102321003-1332003001313021"></a>

<a id="canonical-3220222302212102-2131313200323211-1000121032233121-2310220323230013-1310201122332311-0103220113131202-3002113210102011-2312133032032231"></a>

#### `stateful_service.containers.liveness_check.initial_delay` property

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

<a id="canonical-0003323112313222-3302213000233133-2321133132202311-1232003133332132-0320230130211313-3132223220110230-2232013010322113-0033011103103131"></a>

<a id="canonical-0323113230002230-2122010330323310-0231221303322030-1221333011210212-2210301131033112-3102133223002102-0033222322313333-3310230002232002"></a>

#### `stateful_service.containers.liveness_check.interval` property

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

- [tcp_health_check](data-sources--workload--reference--group-027.md#canonical-3010213213100032-1012001102113101-2302023302130330-2010220100111201-1033033130200033-2202000112003332-1301203212300333-2110221011101301): complete subsection reference.

<a id="canonical-1310310313102130-2211132030331201-1113333331010012-3311111220122111-1322201012000123-3132011112311311-2300001122022033-0332013022233211"></a>

<a id="canonical-2000120231113110-3121110221121103-3333032020200303-3231102320032200-3331330210211030-1221030232022110-2333313212103033-0032212302230301"></a>

#### `stateful_service.containers.liveness_check.timeout` property

Type: `"number"`. Computed.

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

<a id="canonical-2122231002103123-2302031011013321-2123303302122200-0133320113130122-0110121132321023-3111112322103021-2132020200031201-3133311310313202"></a>

<a id="canonical-1031022333103032-3020021202001201-3101201333021101-0302103223132322-1223301001332023-1021102312031103-2333000101223233-3002201311012312"></a>

#### `stateful_service.containers.liveness_check.unhealthy_threshold` property

Type: `"number"`. Computed.

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

<a id="canonical-1132220311013222-1220302112001030-2122320100000223-0201323003102123-3132111333220202-0222233122023211-3030003230120131-2102013312221301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.containers.liveness_check.exec_health_check` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.containers](data-sources--workload--reference--group-026.md#canonical-3101311130033212-3010211101012322-2332300122313021-2300111123200012-3222222201300313-2322010011032230-2110031201133132-3103131301302003)
- [stateful_service.containers.liveness_check](data-sources--workload--reference--group-027.md#canonical-1000133100210330-0101332000310233-1022300301200130-1202221333012211-3002300330030301-2323010320303223-3322100021113321-2121321202120322)
- stateful_service.containers.liveness_check.exec_health_check

<a id="canonical-3023222320333210-1113203301221020-0002110020023133-2122100021021213-0120331123013133-3221333030012320-2330100210110301-0133011011313320"></a>

Type: `"single"`. Computed.

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

<a id="canonical-1110312321131011-1030213311133231-1112321322220322-3122332022300132-1330133202300000-1121212300232323-2132002102321030-1132322232313321"></a>

### Direct properties for `stateful_service.containers.liveness_check.exec_health_check`

<a id="canonical-1213300302030003-1202303323320022-3030222132020232-0002021311011300-1300113133202221-0020230003113221-0300322330021010-2322110110301232"></a>

#### `stateful_service.containers.liveness_check.exec_health_check.command` property

Type: `["list", "string"]`. Computed.

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

<a id="canonical-0100221312120310-1003030310020211-3332103202011122-2000031211231222-2133023310103003-1300320030132231-2122022123333022-3201320212013312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.containers.liveness_check.http_health_check` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.containers](data-sources--workload--reference--group-026.md#canonical-3101311130033212-3010211101012322-2332300122313021-2300111123200012-3222222201300313-2322010011032230-2110031201133132-3103131301302003)
- [stateful_service.containers.liveness_check](data-sources--workload--reference--group-027.md#canonical-1000133100210330-0101332000310233-1022300301200130-1202221333012211-3002300330030301-2323010320303223-3322100021113321-2121321202120322)
- stateful_service.containers.liveness_check.http_health_check

<a id="canonical-0310323321033032-0222132201130201-0303232100321111-3102110032232232-3121220212023313-3103123100323200-3023022301013200-0211222021010232"></a>

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

<a id="canonical-0132011111330001-2333112300233030-1001221333011202-1030231210103313-3203102330031123-3210203113221031-3200001133100031-2121120103000132"></a>

### Direct properties for `stateful_service.containers.liveness_check.http_health_check`

<a id="canonical-3101012220132122-1033230312011123-0103003131203230-2130110210222012-1202330101311330-1120101203113130-2000020203031213-2121220312111031"></a>

#### `stateful_service.containers.liveness_check.http_health_check.headers` property

Type: `["map", "string"]`. Computed.

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

<a id="canonical-2221132210123303-1023321111111031-1320311310101312-0113303211131301-3300231212113100-3303312003212313-1221133221220023-1021012030202323"></a>

<a id="canonical-0210031311202300-1131000010313320-0231113203233210-3302030311202220-1012210201302012-3312200032332221-1312020130202312-1221103220213311"></a>

#### `stateful_service.containers.liveness_check.http_health_check.host_header` property

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

<a id="canonical-1233323103011212-1230202003101222-0323013111033033-3000121031313111-2332130130022133-1011013102032230-3100310210011302-3221011012212321"></a>

<a id="canonical-2000130213112210-1210211000101222-1020022201001000-3202131111310012-2133002202313331-2320002112112112-3302303111122132-3322032120302002"></a>

#### `stateful_service.containers.liveness_check.http_health_check.path` property

Type: `"string"`. Computed.

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

- [port](data-sources--workload--reference--group-027.md#canonical-1310032210323030-3033100031312022-2131302223001113-1332031022200130-3333002311103212-1122010200122032-1101020331132220-3102131320020331): complete subsection reference.

<a id="canonical-1310032210323030-3033100031312022-2131302223001113-1332031022200130-3333002311103212-1122010200122032-1101020331132220-3102131320020331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.containers.liveness_check.http_health_check.port` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.containers](data-sources--workload--reference--group-026.md#canonical-3101311130033212-3010211101012322-2332300122313021-2300111123200012-3222222201300313-2322010011032230-2110031201133132-3103131301302003)
- [stateful_service.containers.liveness_check](data-sources--workload--reference--group-027.md#canonical-1000133100210330-0101332000310233-1022300301200130-1202221333012211-3002300330030301-2323010320303223-3322100021113321-2121321202120322)
- [stateful_service.containers.liveness_check.http_health_check](data-sources--workload--reference--group-027.md#canonical-0100221312120310-1003030310020211-3332103202011122-2000031211231222-2133023310103003-1300320030132231-2122022123333022-3201320212013312)
- stateful_service.containers.liveness_check.http_health_check.port

<a id="canonical-1302130201103200-3221220023101303-0333212201130100-0011021233202120-1010001003112203-1111201313030122-2131201311133323-1201131313131133"></a>

Type: `"single"`. Computed.

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

<a id="canonical-0213100321103221-3230020010201112-1111001311330010-0323131003102331-0221122222000111-3311320110132020-0021132302131201-2200201223123220"></a>

### Direct properties for `stateful_service.containers.liveness_check.http_health_check.port`

<a id="canonical-1121120231211213-0110003313003010-0212333223012332-2203221231200212-3310311330110032-1100333030323101-2200103112001212-1113323310320211"></a>

#### `stateful_service.containers.liveness_check.http_health_check.port.name` property

Type: `"string"`. Computed.

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

<a id="canonical-2133113223323120-1200000111100111-0021113130022332-3112001123332210-3213102210313131-3331220220232210-1323023212232011-2113131302121221"></a>

<a id="canonical-0230232320312312-3313013300103122-1101332223330321-0301311200230010-1312133101123103-3012222130101312-3212112113020231-0123120010113001"></a>

#### `stateful_service.containers.liveness_check.http_health_check.port.num` property

Type: `"number"`. Computed.

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

<a id="canonical-3010213213100032-1012001102113101-2302023302130330-2010220100111201-1033033130200033-2202000112003332-1301203212300333-2110221011101301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.containers.liveness_check.tcp_health_check` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.containers](data-sources--workload--reference--group-026.md#canonical-3101311130033212-3010211101012322-2332300122313021-2300111123200012-3222222201300313-2322010011032230-2110031201133132-3103131301302003)
- [stateful_service.containers.liveness_check](data-sources--workload--reference--group-027.md#canonical-1000133100210330-0101332000310233-1022300301200130-1202221333012211-3002300330030301-2323010320303223-3322100021113321-2121321202120322)
- stateful_service.containers.liveness_check.tcp_health_check

<a id="canonical-2233233232012023-3231333103323302-0113323223100222-2032200222310011-3010213132012020-2101232100121111-3023131200231320-1201011103130121"></a>

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

<a id="canonical-2122211032012202-2103112002222031-0233223202101102-2313303321213333-3032101322221313-3311020201210212-2000232013103202-3122322100310121"></a>

### Direct properties for `stateful_service.containers.liveness_check.tcp_health_check`

- [port](data-sources--workload--reference--group-027.md#canonical-3300323201020302-1210100032132100-3222120112220012-0012030222310100-3132201311032203-3322023120200100-2311331120120301-1123311202023003): complete subsection reference.

<a id="canonical-3300323201020302-1210100032132100-3222120112220012-0012030222310100-3132201311032203-3322023120200100-2311331120120301-1123311202023003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.containers.liveness_check.tcp_health_check.port` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.containers](data-sources--workload--reference--group-026.md#canonical-3101311130033212-3010211101012322-2332300122313021-2300111123200012-3222222201300313-2322010011032230-2110031201133132-3103131301302003)
- [stateful_service.containers.liveness_check](data-sources--workload--reference--group-027.md#canonical-1000133100210330-0101332000310233-1022300301200130-1202221333012211-3002300330030301-2323010320303223-3322100021113321-2121321202120322)
- [stateful_service.containers.liveness_check.tcp_health_check](data-sources--workload--reference--group-027.md#canonical-3010213213100032-1012001102113101-2302023302130330-2010220100111201-1033033130200033-2202000112003332-1301203212300333-2110221011101301)
- stateful_service.containers.liveness_check.tcp_health_check.port

<a id="canonical-3101020302101321-2221031033323323-1000111011033331-1220232311123003-1322323032100223-2311013312332031-2231322332220010-3203111312232000"></a>

Type: `"single"`. Computed.

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

<a id="canonical-2231111110022132-3121011031231133-2110312233323101-2012202332100130-3002132301112233-2330011323103010-3303022011102000-1111310212300332"></a>

### Direct properties for `stateful_service.containers.liveness_check.tcp_health_check.port`

<a id="canonical-0133120321332133-1223100033331201-1312112010332133-3233131103100112-3210020213321312-2131212113320102-0111231220301003-1323030211133103"></a>

#### `stateful_service.containers.liveness_check.tcp_health_check.port.name` property

Type: `"string"`. Computed.

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

<a id="canonical-3012301110300231-0311032213223203-1131031303220231-3232222201031203-0211310130311202-2332012101312103-3020330211100033-2203110131222302"></a>

<a id="canonical-3000201030213332-1113103123210023-2312232333300210-3103100320211331-1100110320202133-0220030003212332-2000330303223212-1303221223113133"></a>

#### `stateful_service.containers.liveness_check.tcp_health_check.port.num` property

Type: `"number"`. Computed.

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

<a id="canonical-1203230022130102-2123013130220100-0133020010102133-0210200010302332-2332232103111100-2033223132011311-0213121011301300-3330312202013233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.containers.readiness_check` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.containers](data-sources--workload--reference--group-026.md#canonical-3101311130033212-3010211101012322-2332300122313021-2300111123200012-3222222201300313-2322010011032230-2110031201133132-3103131301302003)
- stateful_service.containers.readiness_check

<a id="canonical-2021101222001232-2303012303121310-2301002100231010-0033101332210321-0002112121113211-0033101020032223-0010333130012201-2211012020112113"></a>

Type: `"single"`. Computed.

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

<a id="canonical-3112102032033133-3233023120233122-2302322203203213-0123100330302211-2023002230121212-0231321122001300-0311301312333213-2320312313032002"></a>

### Direct properties for `stateful_service.containers.readiness_check`

- [exec_health_check](data-sources--workload--reference--group-027.md#canonical-0320120322013123-2223012322302220-3220023131010202-1012331131121101-2231211232233011-1011030232002311-1200003031322003-3122213231302232): complete subsection reference.

<a id="canonical-2232013002011030-3122002201203001-2212030132000100-0130232221331001-3002001312313203-2100003102002222-1013103132120013-3020001011130003"></a>

<a id="canonical-3203033212101303-2212332300102023-0203333321003232-3030232331201011-0131113002331202-3031203231332312-3133123323132323-2030102332102202"></a>

#### `stateful_service.containers.readiness_check.healthy_threshold` property

Type: `"number"`. Computed.

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

- [http_health_check](data-sources--workload--reference--group-027.md#canonical-2221120320110222-0203030303123120-1022210310203022-2133303311100130-2023121320332201-2122111211230120-0232202330231011-2210300332021223): complete subsection reference.

<a id="canonical-3322323000302302-3032133122131231-0302323333113213-1303220121023123-2021323123030020-1030122232000233-3302131121330330-1222220203331333"></a>

<a id="canonical-0232113022033002-3300222310311033-0101002123001030-2201323230001101-1232233032213303-2330312210210121-0122301330101032-2313023213212203"></a>

#### `stateful_service.containers.readiness_check.initial_delay` property

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

<a id="canonical-0313313020321111-0122320233101202-3313312011223232-1202330322222311-0110002102201201-2121001113313302-0310131023201231-2131012321213233"></a>

<a id="canonical-0333101220332112-2111322030302300-2323203213103012-2322331002212120-3030101212210102-1102212231332220-3122331110120011-0001330110301332"></a>

#### `stateful_service.containers.readiness_check.interval` property

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

- [tcp_health_check](data-sources--workload--reference--group-027.md#canonical-0021333302201130-1003113203330013-3311321231101013-3101213222001310-0032300102300031-2223113122000233-1130230313011130-1302303120023300): complete subsection reference.

<a id="canonical-1331112333201203-1232230320323232-2001012303103023-3233122101113231-2323322221301332-0102323012023220-0002201031330003-3133203012012121"></a>

<a id="canonical-1333103012220121-3110011322300302-0230303011312111-3202111031010002-0032132300222321-1122202320303032-3221121022301320-3033220230203310"></a>

#### `stateful_service.containers.readiness_check.timeout` property

Type: `"number"`. Computed.

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

<a id="canonical-3223101310301032-1110013233213112-0020112023103200-2100301211212100-3321301231122121-0023220200221113-2332312330030102-0022020320003302"></a>

<a id="canonical-3131021301232211-2101113110312230-1112033132012321-3112101210302320-0310221132323230-1230120123122020-0210010221111112-1232010032200113"></a>

#### `stateful_service.containers.readiness_check.unhealthy_threshold` property

Type: `"number"`. Computed.

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

<a id="canonical-0320120322013123-2223012322302220-3220023131010202-1012331131121101-2231211232233011-1011030232002311-1200003031322003-3122213231302232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.containers.readiness_check.exec_health_check` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.containers](data-sources--workload--reference--group-026.md#canonical-3101311130033212-3010211101012322-2332300122313021-2300111123200012-3222222201300313-2322010011032230-2110031201133132-3103131301302003)
- [stateful_service.containers.readiness_check](data-sources--workload--reference--group-027.md#canonical-1203230022130102-2123013130220100-0133020010102133-0210200010302332-2332232103111100-2033223132011311-0213121011301300-3330312202013233)
- stateful_service.containers.readiness_check.exec_health_check

<a id="canonical-0001121130111011-2103001233321213-0102033121103130-2203123203022233-1300301111232231-2211022030102200-0220201221312001-2203123132013321"></a>

Type: `"single"`. Computed.

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

<a id="canonical-0321212203101020-2333110103321200-1331130231111130-1110230122020300-1120200322223211-3000320212303033-1132133122313133-1302133303221313"></a>

### Direct properties for `stateful_service.containers.readiness_check.exec_health_check`

<a id="canonical-0102323121203001-1302331023323321-1300122111133011-1220013333111102-2000312310102023-3213322330022312-1302221123311213-1032222103110111"></a>

#### `stateful_service.containers.readiness_check.exec_health_check.command` property

Type: `["list", "string"]`. Computed.

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

<a id="canonical-2221120320110222-0203030303123120-1022210310203022-2133303311100130-2023121320332201-2122111211230120-0232202330231011-2210300332021223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.containers.readiness_check.http_health_check` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.containers](data-sources--workload--reference--group-026.md#canonical-3101311130033212-3010211101012322-2332300122313021-2300111123200012-3222222201300313-2322010011032230-2110031201133132-3103131301302003)
- [stateful_service.containers.readiness_check](data-sources--workload--reference--group-027.md#canonical-1203230022130102-2123013130220100-0133020010102133-0210200010302332-2332232103111100-2033223132011311-0213121011301300-3330312202013233)
- stateful_service.containers.readiness_check.http_health_check

<a id="canonical-0230301310111302-1221223012203032-1320022121010231-0012021033023230-2222213210200233-0220130111233301-3203232020032223-0231213133021320"></a>

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

<a id="canonical-1001320322022032-0313011303202011-3102000001131020-0301021113010121-0123210020321022-2012032323011023-2332323322333210-3201022001223031"></a>

### Direct properties for `stateful_service.containers.readiness_check.http_health_check`

<a id="canonical-2323120331123100-2022131320213300-1021312112200130-2021302301321002-3021021313032132-1102020030101221-1123112130013010-3303110323301032"></a>

#### `stateful_service.containers.readiness_check.http_health_check.headers` property

Type: `["map", "string"]`. Computed.

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

<a id="canonical-0020130210100311-3020030332130231-2031001221303220-3231333200121031-1220232211012131-2132313322001201-1233113333210003-1321211200122210"></a>

<a id="canonical-0023130213222321-3012233322111311-1221211200031322-1211201231330102-0333123323101131-1000323120200023-2132231302213333-1031223131220010"></a>

#### `stateful_service.containers.readiness_check.http_health_check.host_header` property

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

<a id="canonical-2321301101313130-0302103211232021-2023230202101012-2120132110100133-0230113301310210-2301001013011323-1223121110220012-2121022123100111"></a>

<a id="canonical-1132221131333313-1133313111223320-2332132032111031-0331113020211003-2333110300212112-0302320013002221-1033320133303032-0120112212331001"></a>

#### `stateful_service.containers.readiness_check.http_health_check.path` property

Type: `"string"`. Computed.

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

- [port](data-sources--workload--reference--group-027.md#canonical-1013230231132202-2322213000211130-2033303203210002-3311323120302002-1302300133301311-0331212223230202-1101121122122221-3113323113210212): complete subsection reference.

<a id="canonical-1013230231132202-2322213000211130-2033303203210002-3311323120302002-1302300133301311-0331212223230202-1101121122122221-3113323113210212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.containers.readiness_check.http_health_check.port` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.containers](data-sources--workload--reference--group-026.md#canonical-3101311130033212-3010211101012322-2332300122313021-2300111123200012-3222222201300313-2322010011032230-2110031201133132-3103131301302003)
- [stateful_service.containers.readiness_check](data-sources--workload--reference--group-027.md#canonical-1203230022130102-2123013130220100-0133020010102133-0210200010302332-2332232103111100-2033223132011311-0213121011301300-3330312202013233)
- [stateful_service.containers.readiness_check.http_health_check](data-sources--workload--reference--group-027.md#canonical-2221120320110222-0203030303123120-1022210310203022-2133303311100130-2023121320332201-2122111211230120-0232202330231011-2210300332021223)
- stateful_service.containers.readiness_check.http_health_check.port

<a id="canonical-3201220220303100-0220120303202302-1203300030101212-1220220101031310-1310022220011103-0131130230012330-0121033312311033-0312322132001120"></a>

Type: `"single"`. Computed.

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

<a id="canonical-0333201130012313-2001032312331110-0033112200311033-2001301100321013-2121013020003221-1023023131120033-3311112303012321-2033300000130331"></a>

### Direct properties for `stateful_service.containers.readiness_check.http_health_check.port`

<a id="canonical-3321013000112003-3230321332302030-1200120123301003-0023310220002212-1213031332302310-3130311332213001-0301012003200233-1230210221233212"></a>

#### `stateful_service.containers.readiness_check.http_health_check.port.name` property

Type: `"string"`. Computed.

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

<a id="canonical-3011331232110210-1120310113322331-0133333111000002-0032302000020300-0212021120302302-3112323223233212-2112131133303212-1103210221311212"></a>

<a id="canonical-3032301020221333-0211001331320320-1003121002220122-0130031231331131-0233022331031031-3222021130001322-0303330312201112-1010123302302033"></a>

#### `stateful_service.containers.readiness_check.http_health_check.port.num` property

Type: `"number"`. Computed.

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

<a id="canonical-0021333302201130-1003113203330013-3311321231101013-3101213222001310-0032300102300031-2223113122000233-1130230313011130-1302303120023300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.containers.readiness_check.tcp_health_check` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.containers](data-sources--workload--reference--group-026.md#canonical-3101311130033212-3010211101012322-2332300122313021-2300111123200012-3222222201300313-2322010011032230-2110031201133132-3103131301302003)
- [stateful_service.containers.readiness_check](data-sources--workload--reference--group-027.md#canonical-1203230022130102-2123013130220100-0133020010102133-0210200010302332-2332232103111100-2033223132011311-0213121011301300-3330312202013233)
- stateful_service.containers.readiness_check.tcp_health_check

<a id="canonical-2121000300313301-2133030200031122-1230200121223310-1101232013302010-0021010220223010-0233130112101211-2322001012112201-1223001120311233"></a>

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

<a id="canonical-3323111101130313-2122002003131101-1320022301013031-3002122003013023-2210122312003123-0130130132132000-3300221113131303-2333213223230032"></a>

### Direct properties for `stateful_service.containers.readiness_check.tcp_health_check`

- [port](data-sources--workload--reference--group-027.md#canonical-2003202201121221-3022203000202323-0130021031331303-3120110332300121-2130303333130212-2232332320233101-2203121023001203-0132032223321001): complete subsection reference.

<a id="canonical-2003202201121221-3022203000202323-0130021031331303-3120110332300121-2130303333130212-2232332320233101-2203121023001203-0132032223321001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.containers.readiness_check.tcp_health_check.port` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.containers](data-sources--workload--reference--group-026.md#canonical-3101311130033212-3010211101012322-2332300122313021-2300111123200012-3222222201300313-2322010011032230-2110031201133132-3103131301302003)
- [stateful_service.containers.readiness_check](data-sources--workload--reference--group-027.md#canonical-1203230022130102-2123013130220100-0133020010102133-0210200010302332-2332232103111100-2033223132011311-0213121011301300-3330312202013233)
- [stateful_service.containers.readiness_check.tcp_health_check](data-sources--workload--reference--group-027.md#canonical-0021333302201130-1003113203330013-3311321231101013-3101213222001310-0032300102300031-2223113122000233-1130230313011130-1302303120023300)
- stateful_service.containers.readiness_check.tcp_health_check.port

<a id="canonical-2001302222212012-1312113302211032-0311313131321130-0131003332311200-0221200030030203-3133123112010011-0123003011200213-2232233020012003"></a>

Type: `"single"`. Computed.

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

<a id="canonical-1011101211131033-2202100122111112-0102223233131010-0122222320102221-2111221002302212-2313332200221033-2022021223020321-3223100303013010"></a>

### Direct properties for `stateful_service.containers.readiness_check.tcp_health_check.port`

<a id="canonical-3010000110000233-0222233023211003-1032100030221311-3301223221001211-3212113100020302-3211033220023313-2211222130200012-2302312130032320"></a>

#### `stateful_service.containers.readiness_check.tcp_health_check.port.name` property

Type: `"string"`. Computed.

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

<a id="canonical-2300013002203121-0322012332321133-0321332101332120-3112130302111202-2001323222032020-1310202002003222-3300301003321001-0112130103122213"></a>

<a id="canonical-1032120321103101-3022320213100233-3223032331110031-2233312023323102-3322210221221121-3303030320003220-0032131322133012-2021201031111033"></a>

#### `stateful_service.containers.readiness_check.tcp_health_check.port.num` property

Type: `"number"`. Computed.

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

<a id="canonical-1013001210313000-3001012211030201-1210322313302221-3132213121000320-0012111130102110-3332131123011111-1230211023121203-1131013202130133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.deploy_options` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- stateful_service.deploy_options

<a id="canonical-1132023023322302-2321300032302301-3100022033300201-1130122110331131-2231203223120122-3301121111201310-3333031021133101-3232100331201203"></a>

Type: `"single"`. Computed.

Deploy OPTIONS are used to configure the workload deployment OPTIONS.

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

<a id="canonical-2003333331333220-3333112222332313-0030231011332231-2212110330022313-3011331130203312-0100031010301112-0132323020021020-3223220233220322"></a>

### Direct properties for `stateful_service.deploy_options`

- [all_res](data-sources--workload--reference--group-027.md#canonical-0112010112103121-2110132231312002-0122021302103123-1023223203121322-1202123021310332-0330111030120303-0311000302211232-2221130332320223): complete subsection reference.

- [default_virtual_sites](data-sources--workload--reference--group-027.md#canonical-1032122122323231-2201023322013232-3021033301033031-3301233103200013-1322000303031300-0031120101130010-2311002212001111-3221230231001131): complete subsection reference.

- [deploy_ce_sites](data-sources--workload--reference--group-027.md#canonical-0331120123230233-2122000011230021-0222222332013221-3231012212320100-1230021300310130-0202332212333100-1010010101020110-0000323233332233): complete subsection reference.

- [deploy_ce_virtual_sites](data-sources--workload--reference--group-027.md#canonical-3220302333132100-3232010003023122-3011323020231132-3103313223111220-1201222012133121-0120311120331102-1102320132202323-3230001113113103): complete subsection reference.

- [deploy_re_sites](data-sources--workload--reference--group-027.md#canonical-3232220331300320-1231311312321222-3020011203121022-0230310300013202-0323220333320003-0333201100202213-1220302002200133-1303031030233131): complete subsection reference.

- [deploy_re_virtual_sites](data-sources--workload--reference--group-027.md#canonical-3121001123303100-3221011200111132-3002021330303313-1333120230130323-1031101200021213-1323132121332323-3112123112321231-2032233011011030): complete subsection reference.

<a id="canonical-0112010112103121-2110132231312002-0122021302103123-1023223203121322-1202123021310332-0330111030120303-0311000302211232-2221130332320223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.deploy_options.all_res` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.deploy_options](data-sources--workload--reference--group-027.md#canonical-1013001210313000-3001012211030201-1210322313302221-3132213121000320-0012111130102110-3332131123011111-1230211023121203-1131013202130133)
- stateful_service.deploy_options.all_res

<a id="canonical-0311332022210022-3310201111021112-3212101220320111-1032212011032103-2000200011331323-2023132330030031-3223011000100322-2332110322113123"></a>

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

<a id="canonical-1032122122323231-2201023322013232-3021033301033031-3301233103200013-1322000303031300-0031120101130010-2311002212001111-3221230231001131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.deploy_options.default_virtual_sites` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.deploy_options](data-sources--workload--reference--group-027.md#canonical-1013001210313000-3001012211030201-1210322313302221-3132213121000320-0012111130102110-3332131123011111-1230211023121203-1131013202130133)
- stateful_service.deploy_options.default_virtual_sites

<a id="canonical-0001000033021020-3201123313313221-2303032313321200-3130332111322331-3111302220020013-0112111000331300-2132233101303123-2122021202220012"></a>

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

<a id="canonical-0331120123230233-2122000011230021-0222222332013221-3231012212320100-1230021300310130-0202332212333100-1010010101020110-0000323233332233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.deploy_options.deploy_ce_sites` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.deploy_options](data-sources--workload--reference--group-027.md#canonical-1013001210313000-3001012211030201-1210322313302221-3132213121000320-0012111130102110-3332131123011111-1230211023121203-1131013202130133)
- stateful_service.deploy_options.deploy_ce_sites

<a id="canonical-0131003020130212-1122321300220312-3220312033101311-3133321011110220-0223110201102313-3313002102130322-1100021123121203-0323011113012110"></a>

Type: `"single"`. Computed.

This defines a way to deploy a workload on specific Customer sites.

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

<a id="canonical-0302132230213331-1321303021132013-0001013302023112-2302202131030322-3032311122012333-3231021303100022-2232013122233123-1333330332301012"></a>

### Direct properties for `stateful_service.deploy_options.deploy_ce_sites`

- [site](data-sources--workload--reference--group-027.md#canonical-1212012033330321-1032132031121313-2032001032022021-3003111211132101-3111301333331020-3002222321301023-3103031321033120-1100033132131123): complete subsection reference.

<a id="canonical-1212012033330321-1032132031121313-2032001032022021-3003111211132101-3111301333331020-3002222321301023-3103031321033120-1100033132131123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.deploy_options.deploy_ce_sites.site` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.deploy_options](data-sources--workload--reference--group-027.md#canonical-1013001210313000-3001012211030201-1210322313302221-3132213121000320-0012111130102110-3332131123011111-1230211023121203-1131013202130133)
- [stateful_service.deploy_options.deploy_ce_sites](data-sources--workload--reference--group-027.md#canonical-0331120123230233-2122000011230021-0222222332013221-3231012212320100-1230021300310130-0202332212333100-1010010101020110-0000323233332233)
- stateful_service.deploy_options.deploy_ce_sites.site

<a id="canonical-2323010032301221-0101332203033322-2203230002310123-2020113310120101-2322033200333133-1212303011021230-2030310101031103-1213030122032131"></a>

Type: `"list"`. Computed.

Which customer sites should this workload be deployed.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-1322310223310300-0130001312120221-2300232201231022-2103032002332320-2120021021202002-2003102321100123-1032022303110022-1233023131312131"></a>

### Direct properties for `stateful_service.deploy_options.deploy_ce_sites.site`

<a id="canonical-0123000230020100-0213221122102320-2330222122023110-2311033000113113-0020000130010202-0012333111011223-0313100213221322-3223213301222200"></a>

#### `stateful_service.deploy_options.deploy_ce_sites.site.name` property

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

<a id="canonical-1202001300113031-0313033112213212-0311203033100313-2032323302200330-0302200002033130-2232012033201331-1103132331133012-2130200010223312"></a>

<a id="canonical-3003322011322310-0211203113032222-0133230111220301-0222311100230332-2013333322301032-1200100312221303-1112332110022333-2202201212020332"></a>

#### `stateful_service.deploy_options.deploy_ce_sites.site.namespace` property

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

<a id="canonical-0030000000011201-1313100123001333-1221033112213031-2300030313221203-2130201013130223-1231213032120321-2220033221310312-2321230113310313"></a>

<a id="canonical-3331100121210103-2301312001230221-3201001203332201-1233222023213033-1010103113323131-3112223220021010-2021101032300120-2221010300322101"></a>

#### `stateful_service.deploy_options.deploy_ce_sites.site.tenant` property

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

<a id="canonical-3220302333132100-3232010003023122-3011323020231132-3103313223111220-1201222012133121-0120311120331102-1102320132202323-3230001113113103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.deploy_options.deploy_ce_virtual_sites` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.deploy_options](data-sources--workload--reference--group-027.md#canonical-1013001210313000-3001012211030201-1210322313302221-3132213121000320-0012111130102110-3332131123011111-1230211023121203-1131013202130133)
- stateful_service.deploy_options.deploy_ce_virtual_sites

<a id="canonical-1300233010212030-3230120011333201-3221101011213020-3230031100101213-3302323231120202-1311212212223301-1200200001300321-1110120232120221"></a>

Type: `"single"`. Computed.

This defines a way to deploy a workload on specific Customer virtual sites.

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

<a id="canonical-0031020022202213-1013202230003120-2222132013012232-1130300132322133-2202200122100201-3322002213210003-2330121022112001-2210232320233321"></a>

### Direct properties for `stateful_service.deploy_options.deploy_ce_virtual_sites`

- [virtual_site](data-sources--workload--reference--group-027.md#canonical-2030022011332212-3003311100031102-2002203322332302-1221311001033011-3322030200323230-3011202332200200-1011203212121213-1333212222323102): complete subsection reference.

<a id="canonical-2030022011332212-3003311100031102-2002203322332302-1221311001033011-3322030200323230-3011202332200200-1011203212121213-1333212222323102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.deploy_options.deploy_ce_virtual_sites.virtual_site` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.deploy_options](data-sources--workload--reference--group-027.md#canonical-1013001210313000-3001012211030201-1210322313302221-3132213121000320-0012111130102110-3332131123011111-1230211023121203-1131013202130133)
- [stateful_service.deploy_options.deploy_ce_virtual_sites](data-sources--workload--reference--group-027.md#canonical-3220302333132100-3232010003023122-3011323020231132-3103313223111220-1201222012133121-0120311120331102-1102320132202323-3230001113113103)
- stateful_service.deploy_options.deploy_ce_virtual_sites.virtual_site

<a id="canonical-0013310003001312-3201200333012200-1322123322022302-2101201133013032-0211302312102032-3020000323131331-0330113003013222-0203213130312222"></a>

Type: `"list"`. Computed.

Which customer virtual sites should this workload be deployed.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-1203221113000223-3120333200231322-2103002302333311-1131022331313320-2212201310223313-1233100321300303-2101032113022011-2200132002030302"></a>

### Direct properties for `stateful_service.deploy_options.deploy_ce_virtual_sites.virtual_site`

<a id="canonical-2011111111120120-3023312300123013-1001333302133013-3012223230132010-2201221120030102-1001022133133032-0302331033021233-0322211122220030"></a>

#### `stateful_service.deploy_options.deploy_ce_virtual_sites.virtual_site.name` property

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

<a id="canonical-2012110100331330-1202023200311020-0130113003001120-0122222230103032-1203213033032320-3320002010333302-1031012333231311-3023231330322033"></a>

<a id="canonical-1111013030001102-2001221012130013-2313233123100200-2001200202230203-0222211100323231-3133213212113032-2311011321113233-3330310313000132"></a>

#### `stateful_service.deploy_options.deploy_ce_virtual_sites.virtual_site.namespace` property

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

<a id="canonical-0231110021323220-1201211302221001-2330033011010001-0320123033201121-1131213112012303-1203200311030200-0003121001020312-3023103320130032"></a>

<a id="canonical-3011032120310020-3012030132120201-0001120202201110-0322022212120102-1331221011213022-2302010131300130-0302023030200033-0121122222103010"></a>

#### `stateful_service.deploy_options.deploy_ce_virtual_sites.virtual_site.tenant` property

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

<a id="canonical-3232220331300320-1231311312321222-3020011203121022-0230310300013202-0323220333320003-0333201100202213-1220302002200133-1303031030233131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.deploy_options.deploy_re_sites` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.deploy_options](data-sources--workload--reference--group-027.md#canonical-1013001210313000-3001012211030201-1210322313302221-3132213121000320-0012111130102110-3332131123011111-1230211023121203-1131013202130133)
- stateful_service.deploy_options.deploy_re_sites

<a id="canonical-0113010311010031-2020211002011202-3201003002033200-1322132333002022-2102312210212030-1003321223310320-0131332321100012-3013311112332222"></a>

Type: `"single"`. Computed.

This defines a way to deploy a workload on specific Regional Edge sites.

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

<a id="canonical-0032221120003320-0313222331103221-1122103213232113-1122230313331013-3020203032222211-2213223322023000-0202012123202213-1101120003123121"></a>

### Direct properties for `stateful_service.deploy_options.deploy_re_sites`

- [site](data-sources--workload--reference--group-027.md#canonical-0320122303120312-1201023133021320-3032301020101120-1112222111202012-3213300221102020-3020203302010030-1213103332331233-2003233000300130): complete subsection reference.

<a id="canonical-0320122303120312-1201023133021320-3032301020101120-1112222111202012-3213300221102020-3020203302010030-1213103332331233-2003233000300130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.deploy_options.deploy_re_sites.site` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.deploy_options](data-sources--workload--reference--group-027.md#canonical-1013001210313000-3001012211030201-1210322313302221-3132213121000320-0012111130102110-3332131123011111-1230211023121203-1131013202130133)
- [stateful_service.deploy_options.deploy_re_sites](data-sources--workload--reference--group-027.md#canonical-3232220331300320-1231311312321222-3020011203121022-0230310300013202-0323220333320003-0333201100202213-1220302002200133-1303031030233131)
- stateful_service.deploy_options.deploy_re_sites.site

<a id="canonical-1021330312123212-1100111122023310-3122230113210230-0011012330133331-0103002331323032-0110120332000232-3123001203010000-1210333110321103"></a>

Type: `"list"`. Computed.

Which regional edge sites should this workload be deployed.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-2203330300301231-3330233131131203-1101120221003031-0103133122001100-3001123121212212-1010211101331023-3123333130231013-3132223201221203"></a>

### Direct properties for `stateful_service.deploy_options.deploy_re_sites.site`

<a id="canonical-2113321113111011-1312223113303000-2331330032200200-3233031022031120-3133303012102111-1303000111001133-3212111033133230-3132322122112022"></a>

#### `stateful_service.deploy_options.deploy_re_sites.site.name` property

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

<a id="canonical-1033022211132333-2311332232131001-3130002322230022-1011100202102323-1110211311332032-3211123212321130-2311222231300212-3221002302312211"></a>

<a id="canonical-0232221222101320-2310000102000030-2202122023000210-1100302001232010-0123310230200321-3323023332210012-1230332210103231-1331222123020313"></a>

#### `stateful_service.deploy_options.deploy_re_sites.site.namespace` property

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

<a id="canonical-1103033203301203-1232313310022210-3323030112302233-3301010131001232-2301131213033012-2013031131301300-3213031001021033-1230322232213330"></a>

<a id="canonical-0030311311310202-3222312011113301-2103120333200003-1010100303112012-2101312303300212-2001020311232120-1331200111311212-0120132213130000"></a>

#### `stateful_service.deploy_options.deploy_re_sites.site.tenant` property

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

<a id="canonical-3121001123303100-3221011200111132-3002021330303313-1333120230130323-1031101200021213-1323132121332323-3112123112321231-2032233011011030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.deploy_options.deploy_re_virtual_sites` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.deploy_options](data-sources--workload--reference--group-027.md#canonical-1013001210313000-3001012211030201-1210322313302221-3132213121000320-0012111130102110-3332131123011111-1230211023121203-1131013202130133)
- stateful_service.deploy_options.deploy_re_virtual_sites

<a id="canonical-2303312212301020-0031011133322103-3120223332100110-0032130010230122-1333103311002000-0123011200033210-1311000331230131-0030100222230230"></a>

Type: `"single"`. Computed.

This defines a way to deploy a workload on specific Regional Edge virtual sites.

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

<a id="canonical-3021000032021331-3323321312211312-0010100112200230-2002130112122333-3130102102321320-2313100303110001-3233331303032311-2132201001021020"></a>

### Direct properties for `stateful_service.deploy_options.deploy_re_virtual_sites`

- [virtual_site](data-sources--workload--reference--group-027.md#canonical-3302332221220223-1321200010232020-0130010001322311-2021113021211012-1311213201113321-1220103021230332-0303323133120110-2211101031220032): complete subsection reference.

<a id="canonical-3302332221220223-1321200010232020-0130010001322311-2021113021211012-1311213201113321-1220103021230332-0303323133120110-2211101031220032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.deploy_options.deploy_re_virtual_sites.virtual_site` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.deploy_options](data-sources--workload--reference--group-027.md#canonical-1013001210313000-3001012211030201-1210322313302221-3132213121000320-0012111130102110-3332131123011111-1230211023121203-1131013202130133)
- [stateful_service.deploy_options.deploy_re_virtual_sites](data-sources--workload--reference--group-027.md#canonical-3121001123303100-3221011200111132-3002021330303313-1333120230130323-1031101200021213-1323132121332323-3112123112321231-2032233011011030)
- stateful_service.deploy_options.deploy_re_virtual_sites.virtual_site

<a id="canonical-0112302002113330-2032112300102122-0111122210030003-2302322131323120-3301313201010321-3101113313022311-3331331101123113-1212300132210300"></a>

Type: `"list"`. Computed.

Which regional edge virtual sites should this workload be deployed.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-0110202233313000-0232123023311231-0320130212120001-0222330111122021-1322111112212000-0220122120321311-2300122021012310-0313310303221300"></a>

### Direct properties for `stateful_service.deploy_options.deploy_re_virtual_sites.virtual_site`

<a id="canonical-3231212030310021-3230221133231202-2302213132233221-0230220032321020-3202331303310233-1011020111323222-3222321011033303-0032323313212010"></a>

#### `stateful_service.deploy_options.deploy_re_virtual_sites.virtual_site.name` property

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

<a id="canonical-1111103010220230-2311120321323010-1220123202010202-0330122221102332-0230111213310221-1323210233201020-0023310121113202-3102313302302301"></a>

<a id="canonical-0030303003031202-2031133100030311-0233220103120101-2303330222113210-0031202020012033-3102032021033001-1313312033231112-3222132002121032"></a>

#### `stateful_service.deploy_options.deploy_re_virtual_sites.virtual_site.namespace` property

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

<a id="canonical-3000020112022310-2300231111330021-2331322130210211-3213312333303033-2302121320022320-0100031121121222-1001231203122103-1130222331110013"></a>

<a id="canonical-0231313003221013-2032021023320303-3100222201331200-3331322320103110-3103102030221120-3331102133221133-3001330101330102-0030313301331122"></a>

#### `stateful_service.deploy_options.deploy_re_virtual_sites.virtual_site.tenant` property

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

<a id="canonical-1132102002130121-0202013313231301-3123121303133333-3111103021121020-1113123203313322-3332020232010320-2113332300313132-2131322221012303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.persistent_volumes` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- stateful_service.persistent_volumes

<a id="canonical-3100030133312132-2120231301313222-0131230132101330-2330001331120121-0121100113122010-0010311333132011-3011020032300011-2023200012131002"></a>

Type: `"list"`. Computed.

Persistent storage configuration for the service.

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

<a id="canonical-0213333121032033-3203110220110031-2331200021321230-0300113020213223-3123112322320132-2100230321033223-3233232303232132-0302210202103103"></a>

### Direct properties for `stateful_service.persistent_volumes`

<a id="canonical-1303212202100113-2021300000313030-1302012001330320-2321201120013221-1123001230202022-0313323110322132-3122121203231222-0221010030211223"></a>

#### `stateful_service.persistent_volumes.name` property

Type: `"string"`. Computed.

Name. Name of the volume.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

- [persistent_volume](data-sources--workload--reference--group-027.md#canonical-0313133013332231-3300200130030103-2002302022200030-1003021003231002-1121111231212110-2232311322111203-3202032103100110-0130333321220032): complete subsection reference.

<a id="canonical-0313133013332231-3300200130030103-2002302022200030-1003021003231002-1121111231212110-2232311322111203-3202032103100110-0130333321220032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.persistent_volumes.persistent_volume` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.persistent_volumes](data-sources--workload--reference--group-027.md#canonical-1132102002130121-0202013313231301-3123121303133333-3111103021121020-1113123203313322-3332020232010320-2113332300313132-2131322221012303)
- stateful_service.persistent_volumes.persistent_volume

<a id="canonical-1223322300101233-1201112011300032-2100333233210131-2232211131303123-1201303021330133-3323202130203130-1000333032113010-3311201323333131"></a>

Type: `"single"`. Computed.

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

<a id="canonical-3313111122100101-0232000100100202-3222203010100111-2312200023302330-1330303332011102-0201102233110120-3320121120201211-1220231100122033"></a>

### Direct properties for `stateful_service.persistent_volumes.persistent_volume`

- [mount](data-sources--workload--reference--group-027.md#canonical-0221103203320223-3323303223222203-1302323101021332-3203033330232003-0030302111200300-0113302132200212-1203133210122121-3210213202200000): complete subsection reference.

- [storage](data-sources--workload--reference--group-028.md#canonical-3201222030031023-0013133222021331-0121210112232103-3003022313223010-3023331320110200-1003221211311002-0220112010311110-0111303311113230): complete subsection reference.

<a id="canonical-0221103203320223-3323303223222203-1302323101021332-3203033330232003-0030302111200300-0113302132200212-1203133210122121-3210213202200000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.persistent_volumes.persistent_volume.mount` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.persistent_volumes](data-sources--workload--reference--group-027.md#canonical-1132102002130121-0202013313231301-3123121303133333-3111103021121020-1113123203313322-3332020232010320-2113332300313132-2131322221012303)
- [stateful_service.persistent_volumes.persistent_volume](data-sources--workload--reference--group-027.md#canonical-0313133013332231-3300200130030103-2002302022200030-1003021003231002-1121111231212110-2232311322111203-3202032103100110-0130333321220032)
- stateful_service.persistent_volumes.persistent_volume.mount

<a id="canonical-0113212211331202-1100330230032233-0033023121221312-2303123322000101-3100321110230203-0302202220310123-1333202322102233-3101020311102002"></a>

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

<a id="canonical-2131133230022013-2031200020033031-0011302321003221-0231031333121101-0222213130003330-2211323112313223-1323010030330112-3011122333031300"></a>

### Direct properties for `stateful_service.persistent_volumes.persistent_volume.mount`

<a id="canonical-3302120302320010-0221323303032320-3230303101032022-1232300233000001-1121110310323201-0212301321223123-1203001130323203-2313222210013333"></a>

#### `stateful_service.persistent_volumes.persistent_volume.mount.mode` property

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

<a id="canonical-0130031122210211-3301301121211232-0101033002231233-0221110022213023-3033323000120122-1210002130121323-1213101133110021-1312033100312023"></a>

<a id="canonical-3311121222011112-2113122011330301-1102320232213231-2021302031210013-1232001020033020-2132100323213300-1232110031233120-0322322221032013"></a>

#### `stateful_service.persistent_volumes.persistent_volume.mount.mount_path` property

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

<a id="canonical-3310012330120030-1011210130231031-0220101012322121-1221232302323123-0023120232103232-1123002010131113-3120331302312023-2300320200231222"></a>

<a id="canonical-0002100013223130-1300221210001012-1311010232213120-1133221233201310-1012233332110230-2302130333333330-0130001022323100-3102301121302322"></a>

#### `stateful_service.persistent_volumes.persistent_volume.mount.sub_path` property

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
