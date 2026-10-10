---
page_title: "xcsh_workload reference"
subcategory: "Container"
description: "Complete grouped canonical reference for xcsh_workload reference."
---

# xcsh_workload reference

<a id="canonical-2020220101001102-2133303210311300-0310110110222130-0100110211223233-2030002322201020-2310231223210200-2213111230102200-3011023323231031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `simple_service.container.readiness_check.http_health_check` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [simple_service](data-sources--workload--reference--group-015.md#canonical-1133330001222133-0023013333013330-3330013322032021-3302213310300033-2312032031201020-3233031203301231-0303033100232200-0223111033220010)
- [simple_service.container](data-sources--workload--reference--group-015.md#canonical-3120213033223022-3001202202000120-1331313132130330-1102012032232032-1332103223123321-2212010010003232-3123223220322033-1032222002231223)
- [simple_service.container.readiness_check](data-sources--workload--reference--group-015.md#canonical-0312322121032303-0120231311113002-3232323333022313-0313301223202010-0113120313112233-3202000230211322-1223203103333003-2223013222202203)
- simple_service.container.readiness_check.http_health_check

<a id="canonical-0011101110202110-3311200001001210-2012222201302123-3020231320233133-3123321101210003-2113032202101210-2312332101203112-0100212223213020"></a>

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

<a id="canonical-2313200120221200-3023312210201322-3103101011023112-2210210300012121-1021232222011321-1201112303310332-0011102213130321-3302103103110320"></a>

### Direct properties for `simple_service.container.readiness_check.http_health_check`

<a id="canonical-1322002212023233-1332301120223020-1023130331032133-1123222113113211-3001330021300220-2213033203232013-2012130202312023-0003211223213020"></a>

#### `simple_service.container.readiness_check.http_health_check.headers` property

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

<a id="canonical-3011132032230131-3220031320200133-0031022111302130-1022230102330031-3320000111223203-3113200322012232-3211302201213130-1231130000001020"></a>

<a id="canonical-0123300202221213-3200010101212333-3330023222232020-1032220332130103-1332112130200330-2322132021210122-3232103131103333-1023330322103223"></a>

#### `simple_service.container.readiness_check.http_health_check.host_header` property

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
    "ves.io.schema.rules.string.hostport": "true",
    "ves.io.schema.rules.string.max_len": "262"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostport": "true",
    "ves.io.schema.rules.string.max_len": "262"
  }
}
```

<a id="canonical-3031220300221223-3002230201122030-2033300110120033-3120123311020101-0012321212103132-0310100333331021-1322101321122000-3030322230213132"></a>

<a id="canonical-2130201210122121-1300133323233303-1202132311311320-3301032013213330-0021331312220032-3102031213132333-2313113220312202-3203211220101020"></a>

#### `simple_service.container.readiness_check.http_health_check.path` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

- [port](data-sources--workload--reference--group-016.md#canonical-3012203012020033-3212012332003203-1011303232113213-1311031123301011-1333223031113102-0010112133103022-1232333111202001-3023211311112032): complete subsection reference.

<a id="canonical-3012203012020033-3212012332003203-1011303232113213-1311031123301011-1333223031113102-0010112133103022-1232333111202001-3023211311112032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `simple_service.container.readiness_check.http_health_check.port` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [simple_service](data-sources--workload--reference--group-015.md#canonical-1133330001222133-0023013333013330-3330013322032021-3302213310300033-2312032031201020-3233031203301231-0303033100232200-0223111033220010)
- [simple_service.container](data-sources--workload--reference--group-015.md#canonical-3120213033223022-3001202202000120-1331313132130330-1102012032232032-1332103223123321-2212010010003232-3123223220322033-1032222002231223)
- [simple_service.container.readiness_check](data-sources--workload--reference--group-015.md#canonical-0312322121032303-0120231311113002-3232323333022313-0313301223202010-0113120313112233-3202000230211322-1223203103333003-2223013222202203)
- [simple_service.container.readiness_check.http_health_check](data-sources--workload--reference--group-016.md#canonical-2020220101001102-2133303210311300-0310110110222130-0100110211223233-2030002322201020-2310231223210200-2213111230102200-3011023323231031)
- simple_service.container.readiness_check.http_health_check.port

<a id="canonical-1030202312311103-0333023222331030-1221213133032330-2022021002220012-1020031111310032-1110210102212133-0130202211313323-0112322100332300"></a>

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

<a id="canonical-1302332310310032-3232231223130033-1302230100122133-2030231002202221-0021013003321321-1321113230002103-0221113303321332-0012133220103120"></a>

### Direct properties for `simple_service.container.readiness_check.http_health_check.port`

<a id="canonical-1012010003221231-2210123001300300-0001003223130030-3312013033331011-3212013223323301-2113232011033213-0202320322113211-0021332313132211"></a>

#### `simple_service.container.readiness_check.http_health_check.port.name` property

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
    "ves.io.schema.rules.string.iana_svc_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.iana_svc_name": "true"
  }
}
```

<a id="canonical-1130322221310121-3101232000202220-1130101100131111-2322311021112023-2333010131103313-0302203100321231-0320303222131112-2201100010002023"></a>

<a id="canonical-3120100221201300-1003232320003101-1011023211312223-1233323033332010-1101301110033221-1301011210003303-0113211111222300-0213331103301311"></a>

#### `simple_service.container.readiness_check.http_health_check.port.num` property

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
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-2003212223310021-3331301230033133-2110300232002230-1202301130320132-3231030132313333-2123332103111333-2332101310010223-2302122111302123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `simple_service.container.readiness_check.tcp_health_check` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [simple_service](data-sources--workload--reference--group-015.md#canonical-1133330001222133-0023013333013330-3330013322032021-3302213310300033-2312032031201020-3233031203301231-0303033100232200-0223111033220010)
- [simple_service.container](data-sources--workload--reference--group-015.md#canonical-3120213033223022-3001202202000120-1331313132130330-1102012032232032-1332103223123321-2212010010003232-3123223220322033-1032222002231223)
- [simple_service.container.readiness_check](data-sources--workload--reference--group-015.md#canonical-0312322121032303-0120231311113002-3232323333022313-0313301223202010-0113120313112233-3202000230211322-1223203103333003-2223013222202203)
- simple_service.container.readiness_check.tcp_health_check

<a id="canonical-3213231301331133-0033121123031230-3022331211211311-1232032020103233-2013022102032211-3023330222331010-0313023211221023-2311222333221313"></a>

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

<a id="canonical-2013000201010302-0321003330123230-2302330313213033-1330111113310201-0031311313222100-3123103321300221-2233001022201112-2002112120201322"></a>

### Direct properties for `simple_service.container.readiness_check.tcp_health_check`

- [port](data-sources--workload--reference--group-016.md#canonical-3203032101212220-2130333301311320-1231003231210211-3011130133302331-1110300300030003-1233220323132332-2002112230133110-1233100000303333): complete subsection reference.

<a id="canonical-3203032101212220-2130333301311320-1231003231210211-3011130133302331-1110300300030003-1233220323132332-2002112230133110-1233100000303333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `simple_service.container.readiness_check.tcp_health_check.port` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [simple_service](data-sources--workload--reference--group-015.md#canonical-1133330001222133-0023013333013330-3330013322032021-3302213310300033-2312032031201020-3233031203301231-0303033100232200-0223111033220010)
- [simple_service.container](data-sources--workload--reference--group-015.md#canonical-3120213033223022-3001202202000120-1331313132130330-1102012032232032-1332103223123321-2212010010003232-3123223220322033-1032222002231223)
- [simple_service.container.readiness_check](data-sources--workload--reference--group-015.md#canonical-0312322121032303-0120231311113002-3232323333022313-0313301223202010-0113120313112233-3202000230211322-1223203103333003-2223013222202203)
- [simple_service.container.readiness_check.tcp_health_check](data-sources--workload--reference--group-016.md#canonical-2003212223310021-3331301230033133-2110300232002230-1202301130320132-3231030132313333-2123332103111333-2332101310010223-2302122111302123)
- simple_service.container.readiness_check.tcp_health_check.port

<a id="canonical-0012332202330310-0002023210211033-2031330322333001-1300102322120131-1121123002210120-2021231132002033-2332210333123202-1323011122123020"></a>

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

<a id="canonical-3021302021012012-1313032111213313-0012000330200310-2012132323022003-1201132300101233-2123321033133120-2223131231323230-2130321211212303"></a>

### Direct properties for `simple_service.container.readiness_check.tcp_health_check.port`

<a id="canonical-1223022312202202-0313203022322002-1103023303333212-2333303133323230-2312310210020330-2200012101201310-1311131231032200-0213201220210120"></a>

#### `simple_service.container.readiness_check.tcp_health_check.port.name` property

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
    "ves.io.schema.rules.string.iana_svc_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.iana_svc_name": "true"
  }
}
```

<a id="canonical-3232030210322330-3110323013222130-2030331212323020-3031113023310201-2211131233113032-0122123302123113-0220013211120100-3012321212331120"></a>

<a id="canonical-1001330200013110-1011210032033032-2000232221223010-0233100302332313-0013211013101120-0020010011112001-1233323333020113-3231000230023133"></a>

#### `simple_service.container.readiness_check.tcp_health_check.port.num` property

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
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-2112103133232330-2212313300022031-2132030012302003-0033223033113333-2311301323022003-0110122323101031-0101130200003103-2203110002233113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `simple_service.disabled` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [simple_service](data-sources--workload--reference--group-015.md#canonical-1133330001222133-0023013333013330-3330013322032021-3302213310300033-2312032031201020-3233031203301231-0303033100232200-0223111033220010)
- simple_service.disabled

<a id="canonical-1230223131002211-3031103301233131-0310231020102201-2131313210102210-3303321202330111-1220223210213313-2232003201202212-0310301130123300"></a>

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

<a id="canonical-1210220313300213-3013331300310122-2232001332132003-2302230111202103-0230012033213001-3213221121020312-3303110313013123-1101012020030220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `simple_service.do_not_advertise` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [simple_service](data-sources--workload--reference--group-015.md#canonical-1133330001222133-0023013333013330-3330013322032021-3302213310300033-2312032031201020-3233031203301231-0303033100232200-0223111033220010)
- simple_service.do_not_advertise

<a id="canonical-1330003310220120-0230312011113200-3223131133322232-3033310030032111-1032030010133312-3101021220122300-2130020310301320-0021132302232233"></a>

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

<a id="canonical-3001213300210321-2011120130303023-2211001300322013-1332022123332020-0232131121303302-0131000112110201-1133003122121203-3212202000010030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `simple_service.enabled` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [simple_service](data-sources--workload--reference--group-015.md#canonical-1133330001222133-0023013333013330-3330013322032021-3302213310300033-2312032031201020-3233031203301231-0303033100232200-0223111033220010)
- simple_service.enabled

<a id="canonical-1322110013113212-0213001333002132-1330211012032023-0322333022321312-2100330122102130-1202011111103203-1103122331123212-1033302000002210"></a>

Type: `"single"`. Computed.

Persistent storage volume configuration for the workload.

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

<a id="canonical-3112301033220330-3221220203203031-3110213102031333-2101331310132120-1230332121000123-2130332201113331-0231232223002123-0323032301233322"></a>

### Direct properties for `simple_service.enabled`

<a id="canonical-2202200001010330-0332313010010121-2101122010231133-0211332202130231-3303322102002331-3210012032033211-0031222013021033-0131232022202111"></a>

#### `simple_service.enabled.name` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

- [persistent_volume](data-sources--workload--reference--group-016.md#canonical-2123111321221000-0222301112113233-3033212121211311-2222113101333011-3111222331002110-1021332110233023-1113010330312332-2232102112321331): complete subsection reference.

<a id="canonical-2123111321221000-0222301112113233-3033212121211311-2222113101333011-3111222331002110-1021332110233023-1113010330312332-2232102112321331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `simple_service.enabled.persistent_volume` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [simple_service](data-sources--workload--reference--group-015.md#canonical-1133330001222133-0023013333013330-3330013322032021-3302213310300033-2312032031201020-3233031203301231-0303033100232200-0223111033220010)
- [simple_service.enabled](data-sources--workload--reference--group-016.md#canonical-3001213300210321-2011120130303023-2211001300322013-1332022123332020-0232131121303302-0131000112110201-1133003122121203-3212202000010030)
- simple_service.enabled.persistent_volume

<a id="canonical-0102001303002020-2223302120021122-1223332231230032-0203202222000101-2320311120322222-1322320131021030-3221131012103133-3210131012122000"></a>

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

<a id="canonical-0223311232123113-1212022322023222-0201223020323331-0003001311002132-2321320222313332-1022000200132201-0322232332133031-3103212212112322"></a>

### Direct properties for `simple_service.enabled.persistent_volume`

- [mount](data-sources--workload--reference--group-016.md#canonical-1121310022032211-3322100203123030-2130032220221232-1003303032222002-2233101013131002-2032133200010313-0033221033000303-1032322332101333): complete subsection reference.

- [storage](data-sources--workload--reference--group-016.md#canonical-3222112122312030-0100131331212110-3311100330320223-2221122302332032-2221333323323122-0210301032111313-0001233130123103-3111000232322233): complete subsection reference.

<a id="canonical-1121310022032211-3322100203123030-2130032220221232-1003303032222002-2233101013131002-2032133200010313-0033221033000303-1032322332101333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `simple_service.enabled.persistent_volume.mount` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [simple_service](data-sources--workload--reference--group-015.md#canonical-1133330001222133-0023013333013330-3330013322032021-3302213310300033-2312032031201020-3233031203301231-0303033100232200-0223111033220010)
- [simple_service.enabled](data-sources--workload--reference--group-016.md#canonical-3001213300210321-2011120130303023-2211001300322013-1332022123332020-0232131121303302-0131000112110201-1133003122121203-3212202000010030)
- [simple_service.enabled.persistent_volume](data-sources--workload--reference--group-016.md#canonical-2123111321221000-0222301112113233-3033212121211311-2222113101333011-3111222331002110-1021332110233023-1113010330312332-2232102112321331)
- simple_service.enabled.persistent_volume.mount

<a id="canonical-1130201020220121-2313130002312300-3331011100003021-1222030220213300-2103103313000023-2123011320110130-1031002112303001-2101033330203211"></a>

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

<a id="canonical-1330012313132101-2302222220111211-1020110211210002-1231303122323021-0113010003021123-0111232231111022-2022201310021020-0330201212202102"></a>

### Direct properties for `simple_service.enabled.persistent_volume.mount`

<a id="canonical-3001213221110001-1203121211303001-1001120230322211-0013210311031211-0210031223233210-2310022211021300-2331202120113300-3320131200320122"></a>

#### `simple_service.enabled.persistent_volume.mount.mode` property

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

<a id="canonical-2313311310320331-1211303011313033-2222300202321312-3120200020132031-0311331331332312-3120313001211212-0000313000101233-0313221132210002"></a>

<a id="canonical-1030212101021333-1223210110033122-3303111231011210-2200230031323112-2313032122333003-2100011302103111-2031030313203121-2213221010122031"></a>

#### `simple_service.enabled.persistent_volume.mount.mount_path` property

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

<a id="canonical-2102012230331121-2010303122100322-0330032303100120-3031330021111202-2112232100232320-3002210232121120-2222333311321322-3022113223112013"></a>

<a id="canonical-2011103202010211-3321013323130220-3011320321112033-0302220230011101-1320023132101123-0110031103030100-0313213113233300-1103302231102033"></a>

#### `simple_service.enabled.persistent_volume.mount.sub_path` property

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

<a id="canonical-3222112122312030-0100131331212110-3311100330320223-2221122302332032-2221333323323122-0210301032111313-0001233130123103-3111000232322233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `simple_service.enabled.persistent_volume.storage` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [simple_service](data-sources--workload--reference--group-015.md#canonical-1133330001222133-0023013333013330-3330013322032021-3302213310300033-2312032031201020-3233031203301231-0303033100232200-0223111033220010)
- [simple_service.enabled](data-sources--workload--reference--group-016.md#canonical-3001213300210321-2011120130303023-2211001300322013-1332022123332020-0232131121303302-0131000112110201-1133003122121203-3212202000010030)
- [simple_service.enabled.persistent_volume](data-sources--workload--reference--group-016.md#canonical-2123111321221000-0222301112113233-3033212121211311-2222113101333011-3111222331002110-1021332110233023-1113010330312332-2232102112321331)
- simple_service.enabled.persistent_volume.storage

<a id="canonical-3200132300201303-1130203302131102-0102231320112310-2000201302120011-0223123231103233-3023133120302022-0332311330210231-3013111112001021"></a>

Type: `"single"`. Computed.

Persistent storage configuration is used to configure Persistent Volume Claim (PVC).

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

<a id="canonical-1233302122330230-1103102302003133-0321330233202110-1121221112133023-0031010320032201-0013222103321021-1103233331303023-3131302111331120"></a>

### Direct properties for `simple_service.enabled.persistent_volume.storage`

<a id="canonical-0230303230311022-3302003220120332-3133303130133020-3223310213101301-3231232031010320-3233000122122203-1321311032321221-2010130230322100"></a>

#### `simple_service.enabled.persistent_volume.storage.access_mode` property

Type: `"string"`. Computed.

\[Enum:
ACCESS\_MODE\_READ\_WRITE\_ONCE|ACCESS\_MODE\_READ\_WRITE\_MANY|ACCESS\_MODE\_READ\_ONLY\_MANY\]
Persistence storage access mode is used to configure access mode for persistent storage -
ACCESS\_MODE\_READ\_WRITE\_ONCE: Read Write Once Read Write Once is used to mount persistent storage
in read/write mode to exactly 1 host - ACCESS\_MODE\_READ\_WRITE\_MANY: Read Write Many Read Write
Many is used.. Possible values are \`ACCESS\_MODE\_READ\_WRITE\_ONCE\`,
\`ACCESS\_MODE\_READ\_WRITE\_MANY\`, \`ACCESS\_MODE\_READ\_ONLY\_MANY\`. Defaults to
\`ACCESS\_MODE\_READ\_WRITE\_ONCE\`.

Additional upstream details:

Persistence storage access mode is used to configure access mode for persistent storage

&#8203;- ACCESS\_MODE\_READ\_WRITE\_ONCE: Read Write Once

Read Write Once is used to mount persistent storage in read/write mode to exactly 1 host &#8203;-
ACCESS\_MODE\_READ\_WRITE\_MANY: Read Write Many

Read Write Many is used to mount persistent storage in read/write mode to many hosts &#8203;-
ACCESS\_MODE\_READ\_ONLY\_MANY: Read Only Many

Read Only Many is used to mount persistent storage in read-only mode to many hosts.

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

<a id="canonical-1332222100302022-2320303122300130-0310020001323133-3303332212210222-3100221323300101-1001330011112220-0002111321121312-1210213111221110"></a>

<a id="canonical-1101301030013010-0300210121013223-3121010212010330-0033222230211213-0321131131131031-1022001222331101-0213123101100033-1310001131320210"></a>

#### `simple_service.enabled.persistent_volume.storage.class_name` property

Type: `"string"`. Computed.

Exclusive with \[default\] Use the specified class name.

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

- [default](data-sources--workload--reference--group-016.md#canonical-2010000222012032-0002313022130011-3123231302300212-0221121100202132-2233032213121330-1130033321022310-3200020030203002-2320103331230221): complete subsection reference.

<a id="canonical-0323213130210212-1101312012012013-1200331133230021-2003120011310311-1333202321201121-3021022022022131-1233331301002132-0121131020301223"></a>

<a id="canonical-1101010301323201-3312130110113102-1123222012103333-1201333130123001-1122203232122133-1021210321333110-2010020002330301-1311003002001230"></a>

#### `simple_service.enabled.persistent_volume.storage.storage_size` property

Type: `"number"`. Computed.

Size (in GiB). Size in GiB of the persistent storage.

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

<a id="canonical-2010000222012032-0002313022130011-3123231302300212-0221121100202132-2233032213121330-1130033321022310-3200020030203002-2320103331230221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `simple_service.enabled.persistent_volume.storage.default` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [simple_service](data-sources--workload--reference--group-015.md#canonical-1133330001222133-0023013333013330-3330013322032021-3302213310300033-2312032031201020-3233031203301231-0303033100232200-0223111033220010)
- [simple_service.enabled](data-sources--workload--reference--group-016.md#canonical-3001213300210321-2011120130303023-2211001300322013-1332022123332020-0232131121303302-0131000112110201-1133003122121203-3212202000010030)
- [simple_service.enabled.persistent_volume](data-sources--workload--reference--group-016.md#canonical-2123111321221000-0222301112113233-3033212121211311-2222113101333011-3111222331002110-1021332110233023-1113010330312332-2232102112321331)
- [simple_service.enabled.persistent_volume.storage](data-sources--workload--reference--group-016.md#canonical-3222112122312030-0100131331212110-3311100330320223-2221122302332032-2221333323323122-0210301032111313-0001233130123103-3111000232322233)
- simple_service.enabled.persistent_volume.storage.default

<a id="canonical-0201223002010212-3110102220213203-0222230320310113-2102012102232230-1130133322201020-1130133212230300-0310202332201112-1313201011022323"></a>

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

<a id="canonical-1332023020220230-0231312012221313-1132321323222002-2223303033022101-2220302230211001-0323332011222312-2312021312021221-1032021121111230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `simple_service.simple_advertise` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [simple_service](data-sources--workload--reference--group-015.md#canonical-1133330001222133-0023013333013330-3330013322032021-3302213310300033-2312032031201020-3233031203301231-0303033100232200-0223111033220010)
- simple_service.simple_advertise

<a id="canonical-2322232111322303-2221013322111133-2202332330023003-2321130113230202-0111120120132233-1203332010213001-2102213320012002-3133012213113121"></a>

Type: `"single"`. Computed.

Configuration parameter for simple advertise.

Additional upstream details:

Advertise OPTIONS for Simple Service.

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

<a id="canonical-0200213031222231-2301323213221003-3313301231320233-1333312101003313-0321200112302113-1313232211302131-3212320221100122-3203231213211131"></a>

### Direct properties for `simple_service.simple_advertise`

<a id="canonical-1222302213332212-0303231030022002-2120322232333111-1320002331212032-1021132200203002-2200212200321320-2000030303033232-3201312220020033"></a>

#### `simple_service.simple_advertise.domains` property

Type: `["list", "string"]`. Computed.

A list of Domains (host/authority header) that will be matched to Load Balancer. Wildcard hosts are
supported in the suffix or prefix form

Supported Domains and search order: &#8203;1. Exact Domain names: www&#46;example.com. &#8203;2.
Domains starting with a Wildcard: \*.example.com.

Not supported Domains: &#8203;- Just a Wildcard: \* &#8203;- A Wildcard and TLD with no root Domain:
\*.com. &#8203;- A Wildcard not matching a whole DNS label. E.g. \*.example.com and
\*.bar.example.com are valid Wildcards however \*bar.example.com, \*-bar.example.com, and
bar\*.example.com are all invalid.

Additional notes: A Wildcard will not match empty string. E.g. \*.example.com will match
bar.example.com and baz-bar.example.com but not .example.com. The longest Wildcards match first.
Only a single virtual host in the entire route configuration can match on \*. Also a Domain must be
unique across all virtual hosts within an advertise policy.

Domains are also used for SNI matching if the Load Balancer type is HTTPS. Domains also indicate the
list of names for which DNS resolution will be automatically resolved to IP addresses by the system.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-3012203233110003-0322302302023231-3203001032331001-3210023223023133-1200213002110103-1301222023331320-2103210322100200-0232033321321321"></a>

<a id="canonical-0233033021123012-3003032020012120-0312312300111203-1021200021223123-2121123213321313-3120310003233313-3311013021212113-0230201313002211"></a>

#### `simple_service.simple_advertise.service_port` property

Type: `"number"`. Computed.

Service port to advertise on internet via HTTP loadbalancer using port 80.

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
    "minimum": 1024
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1024",
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1024",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- stateful_service

<a id="canonical-3132102321212312-1030011200002201-0312221231111333-3000303301000002-2111003212203322-3013233131032031-0201123032301323-2201123002233101"></a>

Type: `"single"`. Computed.

StatefulService maintains per replica state and each replica has its own persistent storage. Each
replica has a unique network identity and stable storage. Stateful service are used for distributed
stateful applications like Cassandra, MongoDB, redis, etc.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-scaling_choice": "[\"num_replicas\",\"scale_to_zero\"]"
}
```

<a id="canonical-2212123223033322-0222131012302323-1222020132333011-3331221000110112-0030333113311022-0133011333311202-3122103310033210-1033100110312032"></a>

### Direct properties for `stateful_service`

- [advertise_options](data-sources--workload--reference--group-016.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220): complete subsection reference.

- [configuration](data-sources--workload--reference--group-026.md#canonical-2032211221320222-3233032312121000-0321111221221122-3331201113022312-0020000223303230-0020033200320022-2222023002102320-2200212222310100): complete subsection reference.

- [containers](data-sources--workload--reference--group-026.md#canonical-3101311130033212-3010211101012322-2332300122313021-2300111123200012-3222222201300313-2322010011032230-2110031201133132-3103131301302003): complete subsection reference.

- [deploy_options](data-sources--workload--reference--group-027.md#canonical-1013001210313000-3001012211030201-1210322313302221-3132213121000320-0012111130102110-3332131123011111-1230211023121203-1131013202130133): complete subsection reference.

<a id="canonical-0211202302123311-0002302301110323-2121323221310032-1102131221120231-0021022131232020-0033031300212031-3230322112212110-1322232212032232"></a>

<a id="canonical-0002131321331033-3200320233012111-2131131130012131-2213211201232110-1112211123300013-2303231300012000-1021001213023030-2303331022200001"></a>

#### `stateful_service.num_replicas` property

Type: `"number"`. Computed.

Exclusive with \[scale\_to\_zero\] Number of replicas of service to spawn per site.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 5,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
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
    "ves.io.schema.rules.int32.gt": "0",
    "ves.io.schema.rules.int32.lte": "5"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.int32.gt": "0",
    "ves.io.schema.rules.int32.lte": "5"
  }
}
```

- [persistent_volumes](data-sources--workload--reference--group-027.md#canonical-1132102002130121-0202013313231301-3123121303133333-3111103021121020-1113123203313322-3332020232010320-2113332300313132-2131322221012303): complete subsection reference.

- [scale_to_zero](data-sources--workload--reference--group-028.md#canonical-2130310200302331-3301113001002211-0231300130002032-0320200012113302-2010020220101022-0221323130323300-0010123130122332-1232023022230312): complete subsection reference.

- [volumes](data-sources--workload--reference--group-028.md#canonical-2011330200003232-3232210131033220-1331201211101011-0113112011320103-3212220101322221-0100131112031303-3331333222200110-3212321201321212): complete subsection reference.

<a id="canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- stateful_service.advertise_options

<a id="canonical-0111221301333012-1330211132212122-2102332121012301-1222122213133322-2001112100000333-1200220223303322-3302212113001302-2302223221302012"></a>

Type: `"single"`. Computed.

Advertise OPTIONS are used to configure how and where to advertise the workload using load
balancers.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-advertise_choice": "[\"advertise_custom\",\"advertise_in_cluster\",\"advertise_on_public\",\"do_not_advertise\"]"
}
```

<a id="canonical-0203231022303232-3121213122031011-2010030223123110-2202321013202113-0120131321301001-2002003110110213-3121210001311022-0230330000331020"></a>

### Direct properties for `stateful_service.advertise_options`

- [advertise_custom](data-sources--workload--reference--group-016.md#canonical-1111211210302023-1130301301031112-0133023021111122-2200322020203101-0223203003301101-2022121313331313-1323300011032311-0202020122011323): complete subsection reference.

- [advertise_in_cluster](data-sources--workload--reference--group-019.md#canonical-2000012203330031-3013213320003321-1130203203021012-0301212320010201-2013312012132313-2201113011320310-2123010110120120-3030200031122122): complete subsection reference.

- [advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220): complete subsection reference.

- [do_not_advertise](data-sources--workload--reference--group-026.md#canonical-2130301222010233-1223132131303222-0322121021211321-2133011133112020-3312223303312212-3213023022123103-1032021220021202-2003132313121010): complete subsection reference.

<a id="canonical-1111211210302023-1130301301031112-0133023021111122-2200322020203101-0223203003301101-2022121313331313-1323300011032311-0202020122011323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_custom` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- stateful_service.advertise_options.advertise_custom

<a id="canonical-2202303311011230-1323020101303130-2233010130123230-0213033131123000-1113233022133222-3303020021030133-3121112110020212-0122312130311121"></a>

Type: `"single"`. Computed.

Advertise this workload via loadbalancer on specific sites.

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

<a id="canonical-2022302203223122-0203003322313102-2222303323311332-3110001102032212-0111200011310131-1233222023301211-1002211102123221-2212333013300101"></a>

### Direct properties for `stateful_service.advertise_options.advertise_custom`

- [advertise_where](data-sources--workload--reference--group-016.md#canonical-2023022020012320-2131122313011201-2300000120322300-1332100012103313-0221311213321232-0302220100100100-0302230032132330-0021310212030123): complete subsection reference.

- [ports](data-sources--workload--reference--group-016.md#canonical-0013330003233312-1103212221031032-2233120012231121-3101311130002333-2222211213103202-3221232132020020-1233011032133130-3130112100210221): complete subsection reference.

<a id="canonical-2023022020012320-2131122313011201-2300000120322300-1332100012103313-0221311213321232-0302220100100100-0302230032132330-0021310212030123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_custom.advertise_where` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_custom](data-sources--workload--reference--group-016.md#canonical-1111211210302023-1130301301031112-0133023021111122-2200322020203101-0223203003301101-2022121313331313-1323300011032311-0202020122011323)
- stateful_service.advertise_options.advertise_custom.advertise_where

<a id="canonical-2213103203011233-2321231203232002-2003221030021302-1100133330020120-3211300331102113-3303031310020302-2320103112131132-0122001210010333"></a>

Type: `"list"`. Computed.

Where should this load balancer be available.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-1110122211300022-2231000201330321-2330023113131211-3033233313031333-0100300000322310-3213300200102103-1300100200301131-0202322122011233"></a>

### Direct properties for `stateful_service.advertise_options.advertise_custom.advertise_where`

- [site](data-sources--workload--reference--group-016.md#canonical-3320202201130032-0203220333223120-2030320330020021-0221300332330021-1030120320110000-2222122233221333-0132001012231210-0133323000302013): complete subsection reference.

- [virtual_site](data-sources--workload--reference--group-016.md#canonical-2032022022213011-2330313321021313-1332303311331122-3303001132120112-1231213133110121-1001301121011121-0022023232231103-1202232213310331): complete subsection reference.

- [vk8s_service](data-sources--workload--reference--group-016.md#canonical-3011201020212332-2011332133320103-1300120313022330-3010230001020133-2322002100330210-1233003022123011-1021320030320131-2111111310111303): complete subsection reference.

<a id="canonical-3320202201130032-0203220333223120-2030320330020021-0221300332330021-1030120320110000-2222122233221333-0132001012231210-0133323000302013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_custom.advertise_where.site` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_custom](data-sources--workload--reference--group-016.md#canonical-1111211210302023-1130301301031112-0133023021111122-2200322020203101-0223203003301101-2022121313331313-1323300011032311-0202020122011323)
- [stateful_service.advertise_options.advertise_custom.advertise_where](data-sources--workload--reference--group-016.md#canonical-2023022020012320-2131122313011201-2300000120322300-1332100012103313-0221311213321232-0302220100100100-0302230032132330-0021310212030123)
- stateful_service.advertise_options.advertise_custom.advertise_where.site

<a id="canonical-2321120210020110-0010102111230011-3330311021122332-3211222221130230-0123110211223021-1011330300212123-2101031323023333-1103020213000210"></a>

Type: `"single"`. Computed.

This defines a reference to a CE site along with network type and an optional IP address where a
load balancer could be advertised.

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

<a id="canonical-3212011000003000-3310331233311211-3033320132130213-0003232321101002-3203230031101102-3221302021013122-3023113312203302-0131011212110012"></a>

### Direct properties for `stateful_service.advertise_options.advertise_custom.advertise_where.site`

<a id="canonical-3310311013212020-0300120022032111-3220122212032021-1323212311001020-3320121212330131-1132220311130020-2022312132301312-3020001013231120"></a>

#### `stateful_service.advertise_options.advertise_custom.advertise_where.site.ip` property

Type: `"string"`. Computed.

Use given IP address as VIP on the site.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ipv4",
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
    "ves.io.schema.rules.string.ipv4": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv4": "true"
  }
}
```

<a id="canonical-3120102131120100-0020333300301213-3220201321203303-0022111311310320-1300310202222322-3021103121202232-1301132033112322-0233331223323011"></a>

<a id="canonical-0131301002333212-3122020333031233-2223300032200130-3201201313030111-1321332002110031-0211331000022012-2310323110020122-0202321313301313"></a>

#### `stateful_service.advertise_options.advertise_custom.advertise_where.site.network` property

Type: `"string"`. Computed.

\[Enum:
SITE\_NETWORK\_INSIDE\_AND\_OUTSIDE|SITE\_NETWORK\_INSIDE|SITE\_NETWORK\_OUTSIDE|SITE\_NETWORK\_SERVICE|SITE\_NETWORK\_OUTSIDE\_WITH\_INTERNET\_VIP|SITE\_NETWORK\_INSIDE\_AND\_OUTSIDE\_WITH\_INTERNET\_VIP|SITE\_NETWORK\_IP\_FABRIC\]
Defines network types to be used on site All inside and outside networks. All inside and outside
networks with internet VIP support. All inside networks. Possible values are
\`SITE\_NETWORK\_INSIDE\_AND\_OUTSIDE\`, \`SITE\_NETWORK\_INSIDE\`, \`SITE\_NETWORK\_OUTSIDE\`,
\`SITE\_NETWORK\_SERVICE\`, \`SITE\_NETWORK\_OUTSIDE\_WITH\_INTERNET\_VIP\`,
\`SITE\_NETWORK\_INSIDE\_AND\_OUTSIDE\_WITH\_INTERNET\_VIP\`, \`SITE\_NETWORK\_IP\_FABRIC\`.
Defaults to \`SITE\_NETWORK\_INSIDE\_AND\_OUTSIDE\`.

Additional upstream details:

This defines network types to be used on site

All inside and outside networks. All outside networks. All outside networks with internet VIP
support. VK8s service network. &#8203;- SITE\_NETWORK\_IP\_FABRIC: VER IP Fabric network for the
site

This Virtual network type is used for exposing virtual host on IP Fabric network on the VER site or
for endpoint in IP Fabric network.

Receipt-pinned upstream constraints:

```json
{
  "default": "SITE_NETWORK_INSIDE_AND_OUTSIDE",
  "enum": [
    "SITE_NETWORK_INSIDE_AND_OUTSIDE",
    "SITE_NETWORK_INSIDE",
    "SITE_NETWORK_OUTSIDE",
    "SITE_NETWORK_SERVICE",
    "SITE_NETWORK_OUTSIDE_WITH_INTERNET_VIP",
    "SITE_NETWORK_INSIDE_AND_OUTSIDE_WITH_INTERNET_VIP",
    "SITE_NETWORK_IP_FABRIC"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [site](data-sources--workload--reference--group-016.md#canonical-0311323121320121-1121012301100112-1233122132103231-2202232321201330-1300031303100112-3332030020233302-1120023312101132-3022221031130011): complete subsection reference.

<a id="canonical-0311323121320121-1121012301100112-1233122132103231-2202232321201330-1300031303100112-3332030020233302-1120023312101132-3022221031130011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_custom.advertise_where.site.site` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_custom](data-sources--workload--reference--group-016.md#canonical-1111211210302023-1130301301031112-0133023021111122-2200322020203101-0223203003301101-2022121313331313-1323300011032311-0202020122011323)
- [stateful_service.advertise_options.advertise_custom.advertise_where](data-sources--workload--reference--group-016.md#canonical-2023022020012320-2131122313011201-2300000120322300-1332100012103313-0221311213321232-0302220100100100-0302230032132330-0021310212030123)
- [stateful_service.advertise_options.advertise_custom.advertise_where.site](data-sources--workload--reference--group-016.md#canonical-3320202201130032-0203220333223120-2030320330020021-0221300332330021-1030120320110000-2222122233221333-0132001012231210-0133323000302013)
- stateful_service.advertise_options.advertise_custom.advertise_where.site.site

<a id="canonical-2110023130200132-1212200321022322-2230133122200310-1033030131303031-3023231012101030-2233331302211100-3100010111303302-2102020010333103"></a>

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

<a id="canonical-2312133311213330-3310131232302332-1210001120211110-1302311300022323-0313331030131202-0131312132201233-2230213013202302-3023212331303131"></a>

### Direct properties for `stateful_service.advertise_options.advertise_custom.advertise_where.site.site`

<a id="canonical-1311130300100313-0221100332013233-3103331331131133-3212101131303000-3111101023101112-0031213012130202-0203100203122120-2212213122123302"></a>

#### `stateful_service.advertise_options.advertise_custom.advertise_where.site.site.name` property

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

<a id="canonical-0110131230120301-1221321110212120-2203011212300110-3033323221132232-1020222103312030-1202020031121103-1122210333311220-2230200021233100"></a>

<a id="canonical-3200230221300233-0220011103310021-0012310333122111-1220111331111000-3201110320321332-2120232223300031-2232031023013303-0202320102230200"></a>

#### `stateful_service.advertise_options.advertise_custom.advertise_where.site.site.namespace` property

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

<a id="canonical-2323221213120100-2122230131100212-0132000331201023-3120331212323203-2033230120112111-0321030100122301-3001020210333010-1021301112002002"></a>

<a id="canonical-1022030100300321-1222002220001220-3330032300221211-2310110032121020-1012001320210002-3130333323120202-2111102331110321-2330033011311013"></a>

#### `stateful_service.advertise_options.advertise_custom.advertise_where.site.site.tenant` property

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

<a id="canonical-2032022022213011-2330313321021313-1332303311331122-3303001132120112-1231213133110121-1001301121011121-0022023232231103-1202232213310331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_custom.advertise_where.virtual_site` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_custom](data-sources--workload--reference--group-016.md#canonical-1111211210302023-1130301301031112-0133023021111122-2200322020203101-0223203003301101-2022121313331313-1323300011032311-0202020122011323)
- [stateful_service.advertise_options.advertise_custom.advertise_where](data-sources--workload--reference--group-016.md#canonical-2023022020012320-2131122313011201-2300000120322300-1332100012103313-0221311213321232-0302220100100100-0302230032132330-0021310212030123)
- stateful_service.advertise_options.advertise_custom.advertise_where.virtual_site

<a id="canonical-0223123031021033-0000033123311213-3320332121312303-0100103312303031-0003313301012120-0322102212121330-1133131102002332-2200000130333033"></a>

Type: `"single"`. Computed.

This defines a reference to a customer site virtual site along with network type where a load
balancer could be advertised.

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

<a id="canonical-1123212310133333-1022020223031232-0033033133313132-2100211102322233-2030311330222032-2210030311131000-0232223032222333-2321021233133022"></a>

### Direct properties for `stateful_service.advertise_options.advertise_custom.advertise_where.virtual_site`

<a id="canonical-1023032331321300-3313033000233133-2222230220302301-3011122212311010-1212202000311321-1330103312310130-1303020010031030-1210130022210220"></a>

#### `stateful_service.advertise_options.advertise_custom.advertise_where.virtual_site.network` property

Type: `"string"`. Computed.

\[Enum:
SITE\_NETWORK\_INSIDE\_AND\_OUTSIDE|SITE\_NETWORK\_INSIDE|SITE\_NETWORK\_OUTSIDE|SITE\_NETWORK\_SERVICE|SITE\_NETWORK\_OUTSIDE\_WITH\_INTERNET\_VIP|SITE\_NETWORK\_INSIDE\_AND\_OUTSIDE\_WITH\_INTERNET\_VIP|SITE\_NETWORK\_IP\_FABRIC\]
Defines network types to be used on site All inside and outside networks. All inside and outside
networks with internet VIP support. All inside networks. Possible values are
\`SITE\_NETWORK\_INSIDE\_AND\_OUTSIDE\`, \`SITE\_NETWORK\_INSIDE\`, \`SITE\_NETWORK\_OUTSIDE\`,
\`SITE\_NETWORK\_SERVICE\`, \`SITE\_NETWORK\_OUTSIDE\_WITH\_INTERNET\_VIP\`,
\`SITE\_NETWORK\_INSIDE\_AND\_OUTSIDE\_WITH\_INTERNET\_VIP\`, \`SITE\_NETWORK\_IP\_FABRIC\`.
Defaults to \`SITE\_NETWORK\_INSIDE\_AND\_OUTSIDE\`.

Additional upstream details:

This defines network types to be used on site

All inside and outside networks. All outside networks. All outside networks with internet VIP
support. VK8s service network. &#8203;- SITE\_NETWORK\_IP\_FABRIC: VER IP Fabric network for the
site

This Virtual network type is used for exposing virtual host on IP Fabric network on the VER site or
for endpoint in IP Fabric network.

Receipt-pinned upstream constraints:

```json
{
  "default": "SITE_NETWORK_INSIDE_AND_OUTSIDE",
  "enum": [
    "SITE_NETWORK_INSIDE_AND_OUTSIDE",
    "SITE_NETWORK_INSIDE",
    "SITE_NETWORK_OUTSIDE",
    "SITE_NETWORK_SERVICE",
    "SITE_NETWORK_OUTSIDE_WITH_INTERNET_VIP",
    "SITE_NETWORK_INSIDE_AND_OUTSIDE_WITH_INTERNET_VIP",
    "SITE_NETWORK_IP_FABRIC"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [virtual_site](data-sources--workload--reference--group-016.md#canonical-3110123121231321-3121211321023220-2233013112201330-2121100330320032-0012123012203020-3130100111122332-0123223000133331-3031333313133110): complete subsection reference.

<a id="canonical-3110123121231321-3121211321023220-2233013112201330-2121100330320032-0012123012203020-3130100111122332-0123223000133331-3031333313133110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_custom.advertise_where.virtual_site.virtual_site` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_custom](data-sources--workload--reference--group-016.md#canonical-1111211210302023-1130301301031112-0133023021111122-2200322020203101-0223203003301101-2022121313331313-1323300011032311-0202020122011323)
- [stateful_service.advertise_options.advertise_custom.advertise_where](data-sources--workload--reference--group-016.md#canonical-2023022020012320-2131122313011201-2300000120322300-1332100012103313-0221311213321232-0302220100100100-0302230032132330-0021310212030123)
- [stateful_service.advertise_options.advertise_custom.advertise_where.virtual_site](data-sources--workload--reference--group-016.md#canonical-2032022022213011-2330313321021313-1332303311331122-3303001132120112-1231213133110121-1001301121011121-0022023232231103-1202232213310331)
- stateful_service.advertise_options.advertise_custom.advertise_where.virtual_site.virtual_site

<a id="canonical-3232023033111213-1110232213210203-1101333233012212-1223022222131101-0103020310211021-0221313203130313-0000231020020012-0223111222333020"></a>

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

<a id="canonical-3331111312200213-3130331032303301-3120311123013023-0201011023110010-2131012120210023-3303232231130100-0101301302320202-0031222221103010"></a>

### Direct properties for `stateful_service.advertise_options.advertise_custom.advertise_where.virtual_site.virtual_site`

<a id="canonical-1310100322133320-1132100103333233-2001200223113222-3022232303323212-3100221312221331-2323310232220011-1110030102311101-3030123121010331"></a>

#### `stateful_service.advertise_options.advertise_custom.advertise_where.virtual_site.virtual_site.name` property

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

<a id="canonical-2100203120322223-1033320202100131-2100201111000113-0222203113132100-0120323010231231-3103002002222213-3321111233031210-1312002312313022"></a>

<a id="canonical-0032010232313113-3221001100220303-0010122122100132-3332311123032113-3233033331331332-1001301132100322-1030122120121323-2132011021331112"></a>

#### `stateful_service.advertise_options.advertise_custom.advertise_where.virtual_site.virtual_site.namespace` property

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

<a id="canonical-2210210232010130-3221132320123313-0231130122121311-1332122302132311-3121312100011312-0331031230313120-3002311022300110-1212122032132203"></a>

<a id="canonical-1331230231220233-3111131220333012-0322200020212023-1200011001002003-1001222111001023-0330320032100201-1332232133100121-1313320110213130"></a>

#### `stateful_service.advertise_options.advertise_custom.advertise_where.virtual_site.virtual_site.tenant` property

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

<a id="canonical-3011201020212332-2011332133320103-1300120313022330-3010230001020133-2322002100330210-1233003022123011-1021320030320131-2111111310111303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_custom.advertise_where.vk8s_service` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_custom](data-sources--workload--reference--group-016.md#canonical-1111211210302023-1130301301031112-0133023021111122-2200322020203101-0223203003301101-2022121313331313-1323300011032311-0202020122011323)
- [stateful_service.advertise_options.advertise_custom.advertise_where](data-sources--workload--reference--group-016.md#canonical-2023022020012320-2131122313011201-2300000120322300-1332100012103313-0221311213321232-0302220100100100-0302230032132330-0021310212030123)
- stateful_service.advertise_options.advertise_custom.advertise_where.vk8s_service

<a id="canonical-1201310232312011-2310012112333322-3130102301103210-1212322301213032-0031210332130333-2020022000112031-2122330310133020-2111332012223311"></a>

Type: `"single"`. Computed.

This defines a reference to a RE site or virtual site where a load balancer could be advertised in
the vK8s service network.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-choice": "[\"site\",\"virtual_site\"]"
}
```

<a id="canonical-1301103212001123-3033203323311101-3132002122232103-1213123230200200-0000001033220303-2330233211320213-0020311010333022-0101233021233112"></a>

### Direct properties for `stateful_service.advertise_options.advertise_custom.advertise_where.vk8s_service`

- [site](data-sources--workload--reference--group-016.md#canonical-0110130013013011-1110302030103030-0321201230331211-3232303310333121-3310112222313131-1022011002122101-1132220313123223-1032121300032111): complete subsection reference.

- [virtual_site](data-sources--workload--reference--group-016.md#canonical-2222303112221113-0331322312133021-3033323021222002-2232332300210110-3121100200032000-2333032111222110-1120002312333303-3322122122201213): complete subsection reference.

<a id="canonical-0110130013013011-1110302030103030-0321201230331211-3232303310333121-3310112222313131-1022011002122101-1132220313123223-1032121300032111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_custom.advertise_where.vk8s_service.site` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_custom](data-sources--workload--reference--group-016.md#canonical-1111211210302023-1130301301031112-0133023021111122-2200322020203101-0223203003301101-2022121313331313-1323300011032311-0202020122011323)
- [stateful_service.advertise_options.advertise_custom.advertise_where](data-sources--workload--reference--group-016.md#canonical-2023022020012320-2131122313011201-2300000120322300-1332100012103313-0221311213321232-0302220100100100-0302230032132330-0021310212030123)
- [stateful_service.advertise_options.advertise_custom.advertise_where.vk8s_service](data-sources--workload--reference--group-016.md#canonical-3011201020212332-2011332133320103-1300120313022330-3010230001020133-2322002100330210-1233003022123011-1021320030320131-2111111310111303)
- stateful_service.advertise_options.advertise_custom.advertise_where.vk8s_service.site

<a id="canonical-2111110233131321-1320220132223301-0212120323231211-2211202000032331-3321213321100010-3200230220220030-1233113310202110-1010010133210020"></a>

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

<a id="canonical-2310232131322020-3201123300303310-3032310000300100-0103321210212030-2021012231210322-3132113202221210-3332003233122022-3121222230000312"></a>

### Direct properties for `stateful_service.advertise_options.advertise_custom.advertise_where.vk8s_service.site`

<a id="canonical-0322302020001021-3231232032301011-0230112331220113-1310021103020221-0013111030131323-2123200202303302-2120130323202120-0021332113100323"></a>

#### `stateful_service.advertise_options.advertise_custom.advertise_where.vk8s_service.site.name` property

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

<a id="canonical-3103100301000023-3332313002203001-1213120113322222-1302003222201032-2103331322313322-0102012221223203-3020313022032032-3012010330133001"></a>

<a id="canonical-0100321031030203-0033001302131211-0212232002023111-2302320232020303-2102302301303000-1131202003221102-1333133221132213-1333111300211001"></a>

#### `stateful_service.advertise_options.advertise_custom.advertise_where.vk8s_service.site.namespace` property

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

<a id="canonical-2011031030301012-3133102100130030-0111132310230300-1131111200023220-3333031112230031-2300022213002332-3221003221120031-3012121133301313"></a>

<a id="canonical-3003013212112322-2300230323201222-1103330132212123-0112012100123203-3112311230111032-2013321211312122-3131030201030000-0222233112200311"></a>

#### `stateful_service.advertise_options.advertise_custom.advertise_where.vk8s_service.site.tenant` property

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

<a id="canonical-2222303112221113-0331322312133021-3033323021222002-2232332300210110-3121100200032000-2333032111222110-1120002312333303-3322122122201213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_custom.advertise_where.vk8s_service.virtual_site` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_custom](data-sources--workload--reference--group-016.md#canonical-1111211210302023-1130301301031112-0133023021111122-2200322020203101-0223203003301101-2022121313331313-1323300011032311-0202020122011323)
- [stateful_service.advertise_options.advertise_custom.advertise_where](data-sources--workload--reference--group-016.md#canonical-2023022020012320-2131122313011201-2300000120322300-1332100012103313-0221311213321232-0302220100100100-0302230032132330-0021310212030123)
- [stateful_service.advertise_options.advertise_custom.advertise_where.vk8s_service](data-sources--workload--reference--group-016.md#canonical-3011201020212332-2011332133320103-1300120313022330-3010230001020133-2322002100330210-1233003022123011-1021320030320131-2111111310111303)
- stateful_service.advertise_options.advertise_custom.advertise_where.vk8s_service.virtual_site

<a id="canonical-1001331030330331-1223200200201333-1321331330020020-3220331320211002-2301223300031003-1201011200330103-1233322010230223-0010310231132321"></a>

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

<a id="canonical-1021302100121222-2100011023120213-1332123003332020-3202112102311113-3112301133332301-3221131302012122-3232212013203102-0222203331011211"></a>

### Direct properties for `stateful_service.advertise_options.advertise_custom.advertise_where.vk8s_service.virtual_site`

<a id="canonical-2002221300202301-3023313302220213-0101200303011100-1312023020203132-2300013333301100-1212233312121030-0100200331232130-0101023211031021"></a>

#### `stateful_service.advertise_options.advertise_custom.advertise_where.vk8s_service.virtual_site.name` property

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

<a id="canonical-0300303201102131-2331023330012011-0210320222213002-1123023010120033-0120302113131320-0121222222101110-0310210101100010-3003022323211222"></a>

<a id="canonical-1203130300000320-3113002310102012-1121323133233120-3111223310130001-1230011001033323-2332321021000223-0310111132320300-1312111322330201"></a>

#### `stateful_service.advertise_options.advertise_custom.advertise_where.vk8s_service.virtual_site.namespace` property

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

<a id="canonical-0300200302203012-2321023133233203-0013131131213300-1102301030213121-2102201213110200-2131220131221201-3111312012222103-2223110233020012"></a>

<a id="canonical-0033013300111002-2030100133232222-2021321001012230-1301023303012203-0003230031111012-2001133212232210-2200332303311311-0012201231021032"></a>

#### `stateful_service.advertise_options.advertise_custom.advertise_where.vk8s_service.virtual_site.tenant` property

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

<a id="canonical-0013330003233312-1103212221031032-2233120012231121-3101311130002333-2222211213103202-3221232132020020-1233011032133130-3130112100210221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_custom.ports` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_custom](data-sources--workload--reference--group-016.md#canonical-1111211210302023-1130301301031112-0133023021111122-2200322020203101-0223203003301101-2022121313331313-1323300011032311-0202020122011323)
- stateful_service.advertise_options.advertise_custom.ports

<a id="canonical-2332031023012310-1211012112123322-0230032031121211-3231301231203332-3022010032300331-0323330323120321-1023131332101033-1123302030303131"></a>

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

<a id="canonical-3003022000111201-3011030013031311-0110230020003321-2201032101132032-0230200103331132-1320220331300310-1121302101122221-2132033023103010"></a>

### Direct properties for `stateful_service.advertise_options.advertise_custom.ports`

- [http_loadbalancer](data-sources--workload--reference--group-016.md#canonical-0323333020010100-1132122030202301-0220233131023222-0121213003321102-3000102211200300-2122203321110300-0131101323132000-2220030303220201): complete subsection reference.

- [port](data-sources--workload--reference--group-019.md#canonical-3020203222030212-3033113130221110-2320320300332011-1001030302320102-2022321230310000-0112020113321222-3002102322302000-3221222122031101): complete subsection reference.

- [tcp_loadbalancer](data-sources--workload--reference--group-019.md#canonical-3021012100101331-3130231312320112-3123200110222222-2102132131213021-0312223301023023-0100202023203122-0211120233110023-0332201103002102): complete subsection reference.

<a id="canonical-0323333020010100-1132122030202301-0220233131023222-0121213003321102-3000102211200300-2122203321110300-0131101323132000-2220030303220201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_custom](data-sources--workload--reference--group-016.md#canonical-1111211210302023-1130301301031112-0133023021111122-2200322020203101-0223203003301101-2022121313331313-1323300011032311-0202020122011323)
- [stateful_service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-016.md#canonical-0013330003233312-1103212221031032-2233120012231121-3101311130002333-2222211213103202-3221232132020020-1233011032133130-3130112100210221)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer

<a id="canonical-2022122301122103-3000320010020112-0122200002001233-2031231101310101-1223030013300123-2120321210212123-3112121031201101-3110312113333023"></a>

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

<a id="canonical-2032300123313101-2233003220033103-1033223200232212-2102110302013001-1110021031000003-1021002013023111-3033031320013212-2312001000103103"></a>

### Direct properties for `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer`

- [default_route](data-sources--workload--reference--group-016.md#canonical-2020222323220111-1331211203202000-3201020021101201-0111321110223013-1230323223112113-0223023100222323-0302123123231102-2203231302311220): complete subsection reference.

<a id="canonical-0330302213321211-2111203033222023-2002303332300311-2000303132023011-1102231133113203-3001023021322100-3003133021320101-3222132310103302"></a>

<a id="canonical-0102332101103010-1023013331203230-3203220200202021-1231102022312233-3103212210031321-2232200103203300-0030320202023012-2112031232000330"></a>

#### `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.domains` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

- [http](data-sources--workload--reference--group-016.md#canonical-0210000123110012-0203033330020312-0012332023112212-1310121212222023-3210023231230233-1213201111200133-1221001123112101-1011211000210303): complete subsection reference.

- [https](data-sources--workload--reference--group-016.md#canonical-1033211310210331-3122131003202312-2011303010012203-1210330231200011-0103302332022003-0002020201221022-0330311011330033-0101322031121003): complete subsection reference.

- [https_auto_cert](data-sources--workload--reference--group-017.md#canonical-2222310212003321-3311011031323313-1033123102310023-1131322023022203-0310203302121223-0212302033021012-1111313132200003-2303112232301100): complete subsection reference.

- [specific_routes](data-sources--workload--reference--group-018.md#canonical-1013211100312123-2303001111231323-0313131132122000-0333201330302230-1031323200021000-1112233102332030-3102023323212313-2010030013121010): complete subsection reference.

<a id="canonical-2020222323220111-1331211203202000-3201020021101201-0111321110223013-1230323223112113-0223023100222323-0302123123231102-2203231302311220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.default_route` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_custom](data-sources--workload--reference--group-016.md#canonical-1111211210302023-1130301301031112-0133023021111122-2200322020203101-0223203003301101-2022121313331313-1323300011032311-0202020122011323)
- [stateful_service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-016.md#canonical-0013330003233312-1103212221031032-2233120012231121-3101311130002333-2222211213103202-3221232132020020-1233011032133130-3130112100210221)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-016.md#canonical-0323333020010100-1132122030202301-0220233131023222-0121213003321102-3000102211200300-2122203321110300-0131101323132000-2220030303220201)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.default_route

<a id="canonical-1331001032123101-1232131002331201-3012212220101231-3303000331300212-0023321111032120-0030233301312201-3023230301032110-1202203331102313"></a>

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

<a id="canonical-2303030230012203-0300003213313012-2302021332010201-1222032303023322-2123203333331301-1032011003003200-0131110233123201-0222300111300020"></a>

### Direct properties for `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.default_route`

- [auto_host_rewrite](data-sources--workload--reference--group-016.md#canonical-3020030333232332-0113230332320033-0203210331002213-1200033131111332-0111322130123223-1323233013002303-0232323031301101-0322221310113002): complete subsection reference.

- [disable_host_rewrite](data-sources--workload--reference--group-016.md#canonical-3133001210012120-1000122332022200-3132023300022300-2211300033302131-3131012322223130-3131032012210310-2121301122033202-2222323213321303): complete subsection reference.

<a id="canonical-1121330220210123-2302103130113233-1011012021110210-0011330031133232-1210123102010323-1131333312001012-3132312003201301-0202202030002212"></a>

<a id="canonical-2001012001022320-3212213103230302-0002031022203303-1211101221201323-1220031210303312-3210130233311210-0201322011321331-2131022201112101"></a>

#### `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.default_route.host_rewrite` property

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

<a id="canonical-3020030333232332-0113230332320033-0203210331002213-1200033131111332-0111322130123223-1323233013002303-0232323031301101-0322221310113002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.default_route.auto_host_rewrite` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_custom](data-sources--workload--reference--group-016.md#canonical-1111211210302023-1130301301031112-0133023021111122-2200322020203101-0223203003301101-2022121313331313-1323300011032311-0202020122011323)
- [stateful_service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-016.md#canonical-0013330003233312-1103212221031032-2233120012231121-3101311130002333-2222211213103202-3221232132020020-1233011032133130-3130112100210221)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-016.md#canonical-0323333020010100-1132122030202301-0220233131023222-0121213003321102-3000102211200300-2122203321110300-0131101323132000-2220030303220201)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.default_route](data-sources--workload--reference--group-016.md#canonical-2020222323220111-1331211203202000-3201020021101201-0111321110223013-1230323223112113-0223023100222323-0302123123231102-2203231302311220)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.default_route.auto_host_rewrite

<a id="canonical-1031322330221100-3010131221113312-0312212133222203-0312131302031310-1030030222130131-3112012100300300-2331130302222221-2323002211102300"></a>

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

<a id="canonical-3133001210012120-1000122332022200-3132023300022300-2211300033302131-3131012322223130-3131032012210310-2121301122033202-2222323213321303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.default_route.disable_host_rewrite` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_custom](data-sources--workload--reference--group-016.md#canonical-1111211210302023-1130301301031112-0133023021111122-2200322020203101-0223203003301101-2022121313331313-1323300011032311-0202020122011323)
- [stateful_service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-016.md#canonical-0013330003233312-1103212221031032-2233120012231121-3101311130002333-2222211213103202-3221232132020020-1233011032133130-3130112100210221)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-016.md#canonical-0323333020010100-1132122030202301-0220233131023222-0121213003321102-3000102211200300-2122203321110300-0131101323132000-2220030303220201)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.default_route](data-sources--workload--reference--group-016.md#canonical-2020222323220111-1331211203202000-3201020021101201-0111321110223013-1230323223112113-0223023100222323-0302123123231102-2203231302311220)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.default_route.disable_host_rewrite

<a id="canonical-0011213122113022-1113132022020021-1211320213133213-1003133113011012-0233320311303030-2213203112020302-0023313232330213-1013333300011210"></a>

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

<a id="canonical-0210000123110012-0203033330020312-0012332023112212-1310121212222023-3210023231230233-1213201111200133-1221001123112101-1011211000210303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_custom](data-sources--workload--reference--group-016.md#canonical-1111211210302023-1130301301031112-0133023021111122-2200322020203101-0223203003301101-2022121313331313-1323300011032311-0202020122011323)
- [stateful_service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-016.md#canonical-0013330003233312-1103212221031032-2233120012231121-3101311130002333-2222211213103202-3221232132020020-1233011032133130-3130112100210221)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-016.md#canonical-0323333020010100-1132122030202301-0220233131023222-0121213003321102-3000102211200300-2122203321110300-0131101323132000-2220030303220201)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http

<a id="canonical-2131221122103110-2132101202331321-0113230103212231-0312010022332110-3302120211232230-0022102012022231-1022000221133313-3023210113333111"></a>

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

<a id="canonical-0313333012103000-2113000210011310-0331233221200123-2203310322233110-2212330301202020-1210003031311312-2031233121031323-2001323321302112"></a>

### Direct properties for `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http`

<a id="canonical-0100113010111322-3312032011311323-1232313301113332-1111201103222311-0232123331200231-0112012030222113-0011310031111201-2031333101010120"></a>

#### `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http.dns_volterra_managed` property

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

<a id="canonical-0112231201101032-0113212012320101-0031020330200200-2202001000121002-2022100013210011-3213320231100303-0320003201020321-0013213200023133"></a>

<a id="canonical-2021322032113220-2000312301311030-2230100333032232-0323331131000001-2322203113023033-0022010221230120-3013113232000100-3202222312233021"></a>

#### `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http.port` property

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

<a id="canonical-0300020030233032-1010223203221332-2302220121102332-1221202313332222-1203110020033003-1311032001110102-3023102100030220-1131320132300310"></a>

<a id="canonical-0211022330002032-0130322322020132-3102321103211122-1032210132021302-3320303220133131-0130031010312212-1112110131311121-0000033103113013"></a>

#### `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http.port_ranges` property

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

<a id="canonical-1033211310210331-3122131003202312-2011303010012203-1210330231200011-0103302332022003-0002020201221022-0330311011330033-0101322031121003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_custom](data-sources--workload--reference--group-016.md#canonical-1111211210302023-1130301301031112-0133023021111122-2200322020203101-0223203003301101-2022121313331313-1323300011032311-0202020122011323)
- [stateful_service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-016.md#canonical-0013330003233312-1103212221031032-2233120012231121-3101311130002333-2222211213103202-3221232132020020-1233011032133130-3130112100210221)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-016.md#canonical-0323333020010100-1132122030202301-0220233131023222-0121213003321102-3000102211200300-2122203321110300-0131101323132000-2220030303220201)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https

<a id="canonical-0102303302013230-1330320233233310-2113122031113012-3130002023203133-2013130231201011-2332120303302122-1322302211021320-2001313330201211"></a>

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

<a id="canonical-3312222220301231-1002313220320120-0201211001213111-3203002002233220-1103322033310212-1321113331023120-3101323010102121-0231122113211330"></a>

### Direct properties for `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https`

<a id="canonical-2023211030203023-0332211033031000-3113001110013310-3233332332303200-0322212201120211-3303001121031001-3031130200121303-2222110302003113"></a>

#### `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.add_hsts` property

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

<a id="canonical-0333131301310121-1113202230332201-0102101202110101-0000230011233302-2101323222031312-2300331321121321-2200132031303122-0203032023303131"></a>

<a id="canonical-3030120210313213-1012220321032003-2230330322013213-1320001310030120-2023312011301310-1013011213013300-3301113033003011-0111201133100012"></a>

#### `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.append_server_name` property

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
    "ves.io.schema.rules.string.max_len": "8096"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "8096"
  }
}
```

- [coalescing_options](data-sources--workload--reference--group-016.md#canonical-3222230311111111-3020113330012230-0002302032113022-3221131030313033-2131321231133022-0132211032010123-0101033132301003-2200000120321203): complete subsection reference.

<a id="canonical-1202230333231220-2320221010310103-0010310113130021-1323113233231003-0010210312321210-3102221330311201-0320221210101102-2220130230223013"></a>

<a id="canonical-2002010222121212-0210320213302122-2103322020133303-3023210123233321-3213023100202202-2133031311013331-0323212210021032-0323003031213103"></a>

#### `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.connection_idle_timeout` property

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
    "ves.io.schema.rules.uint32.lte": "600000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "600000"
  }
}
```

- [default_header](data-sources--workload--reference--group-016.md#canonical-0023232310300200-3223031333313121-2212120213302110-1333203102110320-3003312012200002-2202003121222200-1031312300201132-0031122320322313): complete subsection reference.

- [default_loadbalancer](data-sources--workload--reference--group-016.md#canonical-0012001103030201-3203200132211322-3232213203333223-3312110000013103-3010201310230301-0300213202122012-0103012122313002-1023213201333300): complete subsection reference.

- [disable_path_normalize](data-sources--workload--reference--group-016.md#canonical-3233221102122003-1130211310313313-2323010000022332-1312131231102103-3101321021113300-1323200323023012-1230101113333022-0202333112310002): complete subsection reference.

- [enable_path_normalize](data-sources--workload--reference--group-016.md#canonical-2013223221202330-3132232032332332-1132202330202010-3210121032132010-2033202003222023-1321220120101313-2201212323132112-0331002302300020): complete subsection reference.

- [http_protocol_options](data-sources--workload--reference--group-016.md#canonical-1120222011000333-0211100013303031-0300010010332233-1031300033323310-3030100021331110-1303033230310102-2010030333330102-1101231301332331): complete subsection reference.

<a id="canonical-2323233222130220-3222032011200211-3223310000021123-3211020332212011-1012331300011200-2232321230210001-1332032012221223-0133001333023032"></a>

<a id="canonical-2013312231110302-0030011103003332-2221331010100210-1123123213111123-1313310010222231-0131022233023130-1323222322102302-2311220110331020"></a>

#### `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_redirect` property

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

- [non_default_loadbalancer](data-sources--workload--reference--group-016.md#canonical-2233212220331030-3333132301233302-1131111310300010-1020030131001213-2022003333122002-1023230213320033-0102023112233011-0323230323302323): complete subsection reference.

- [pass_through](data-sources--workload--reference--group-016.md#canonical-3210111122032323-1101022110110331-1023201230033121-2011003332013100-1002021032212313-0320000121303232-3101312000213323-3232133323310302): complete subsection reference.

<a id="canonical-0312003211101322-3200003333232333-0021102011221020-3203313002313021-0221103131023013-1101113301101100-1223220333232112-0331203131231020"></a>

<a id="canonical-3222120111111001-3330312303010323-0210033202233023-3033112032211113-2220332310002303-3022230121131303-0331222300131020-2200103103130222"></a>

#### `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.port` property

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

<a id="canonical-1203111001210101-3230123313322322-3132030210020303-0101320020220232-1231002112311233-3233233330022101-1123033023222021-1103032120323213"></a>

<a id="canonical-3233002132112201-1002310111111113-3212230231101312-1223323320322233-3322100210311212-2011132302310231-1233212211101223-3132120002212100"></a>

#### `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.port_ranges` property

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

<a id="canonical-1303010020132212-1202002030101021-2213112100111012-0022123103233033-2233010000031101-3200322032230202-0320211113111201-3231123033031232"></a>

<a id="canonical-1033013110301300-1121221123112231-1322221120032100-1111120011312231-1331313133202100-2233120233112310-0100112133023310-3103220321010031"></a>

#### `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.server_name` property

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
    "ves.io.schema.rules.string.max_len": "8096"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "8096"
  }
}
```

- [tls_cert_params](data-sources--workload--reference--group-016.md#canonical-2010001011230201-3020111320131112-2333021133221211-0200112033223200-3003002233220012-1123230332003001-0103023230030232-1202122133301320): complete subsection reference.

- [tls_parameters](data-sources--workload--reference--group-017.md#canonical-3230312021230322-3310001132303313-1110012112202213-3330322321030323-1121221221133120-2321102112023333-2011212100221003-0220120111022302): complete subsection reference.

<a id="canonical-3222230311111111-3020113330012230-0002302032113022-3221131030313033-2131321231133022-0132211032010123-0101033132301003-2200000120321203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.coalescing_options` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_custom](data-sources--workload--reference--group-016.md#canonical-1111211210302023-1130301301031112-0133023021111122-2200322020203101-0223203003301101-2022121313331313-1323300011032311-0202020122011323)
- [stateful_service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-016.md#canonical-0013330003233312-1103212221031032-2233120012231121-3101311130002333-2222211213103202-3221232132020020-1233011032133130-3130112100210221)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-016.md#canonical-0323333020010100-1132122030202301-0220233131023222-0121213003321102-3000102211200300-2122203321110300-0131101323132000-2220030303220201)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https](data-sources--workload--reference--group-016.md#canonical-1033211310210331-3122131003202312-2011303010012203-1210330231200011-0103302332022003-0002020201221022-0330311011330033-0101322031121003)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.coalescing_options

<a id="canonical-0011120122220210-3300331323112202-3111121200130022-2201011323330020-1022111012122223-2331123132332331-3313300313231231-0310110103110001"></a>

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

<a id="canonical-3203121220012233-0231300121011232-0221200211133032-3022032330032300-3030101000131312-0102312202201120-0320010330121210-0010010331031232"></a>

### Direct properties for `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.coalescing_options`

- [default_coalescing](data-sources--workload--reference--group-016.md#canonical-3302211123200123-3303023033113022-0310023022321220-3301022002322232-2121001320100231-2111032202120100-0132133100231101-1001313330303202): complete subsection reference.

- [strict_coalescing](data-sources--workload--reference--group-016.md#canonical-1021033211222331-3220122130223113-1230202132011320-3032113310232333-0302213321012102-0232220222320001-3310023320113101-0333002002031212): complete subsection reference.

<a id="canonical-3302211123200123-3303023033113022-0310023022321220-3301022002322232-2121001320100231-2111032202120100-0132133100231101-1001313330303202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.coalescing_options.default_coalescing` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_custom](data-sources--workload--reference--group-016.md#canonical-1111211210302023-1130301301031112-0133023021111122-2200322020203101-0223203003301101-2022121313331313-1323300011032311-0202020122011323)
- [stateful_service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-016.md#canonical-0013330003233312-1103212221031032-2233120012231121-3101311130002333-2222211213103202-3221232132020020-1233011032133130-3130112100210221)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-016.md#canonical-0323333020010100-1132122030202301-0220233131023222-0121213003321102-3000102211200300-2122203321110300-0131101323132000-2220030303220201)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https](data-sources--workload--reference--group-016.md#canonical-1033211310210331-3122131003202312-2011303010012203-1210330231200011-0103302332022003-0002020201221022-0330311011330033-0101322031121003)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.coalescing_options](data-sources--workload--reference--group-016.md#canonical-3222230311111111-3020113330012230-0002302032113022-3221131030313033-2131321231133022-0132211032010123-0101033132301003-2200000120321203)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.coalescing_options.default_coalescing

<a id="canonical-3000023113021233-3313310221312132-1100013300012132-2103221133011202-1200000102321000-3332323223221022-0303122031111312-0231003003000120"></a>

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

<a id="canonical-1021033211222331-3220122130223113-1230202132011320-3032113310232333-0302213321012102-0232220222320001-3310023320113101-0333002002031212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.coalescing_options.strict_coalescing` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_custom](data-sources--workload--reference--group-016.md#canonical-1111211210302023-1130301301031112-0133023021111122-2200322020203101-0223203003301101-2022121313331313-1323300011032311-0202020122011323)
- [stateful_service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-016.md#canonical-0013330003233312-1103212221031032-2233120012231121-3101311130002333-2222211213103202-3221232132020020-1233011032133130-3130112100210221)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-016.md#canonical-0323333020010100-1132122030202301-0220233131023222-0121213003321102-3000102211200300-2122203321110300-0131101323132000-2220030303220201)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https](data-sources--workload--reference--group-016.md#canonical-1033211310210331-3122131003202312-2011303010012203-1210330231200011-0103302332022003-0002020201221022-0330311011330033-0101322031121003)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.coalescing_options](data-sources--workload--reference--group-016.md#canonical-3222230311111111-3020113330012230-0002302032113022-3221131030313033-2131321231133022-0132211032010123-0101033132301003-2200000120321203)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.coalescing_options.strict_coalescing

<a id="canonical-1103030131202011-1332212233021202-0023200200232010-2133303321332211-3110133033223130-0333320100032322-0333233333323111-2032332111212220"></a>

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

<a id="canonical-0023232310300200-3223031333313121-2212120213302110-1333203102110320-3003312012200002-2202003121222200-1031312300201132-0031122320322313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.default_header` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_custom](data-sources--workload--reference--group-016.md#canonical-1111211210302023-1130301301031112-0133023021111122-2200322020203101-0223203003301101-2022121313331313-1323300011032311-0202020122011323)
- [stateful_service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-016.md#canonical-0013330003233312-1103212221031032-2233120012231121-3101311130002333-2222211213103202-3221232132020020-1233011032133130-3130112100210221)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-016.md#canonical-0323333020010100-1132122030202301-0220233131023222-0121213003321102-3000102211200300-2122203321110300-0131101323132000-2220030303220201)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https](data-sources--workload--reference--group-016.md#canonical-1033211310210331-3122131003202312-2011303010012203-1210330231200011-0103302332022003-0002020201221022-0330311011330033-0101322031121003)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.default_header

<a id="canonical-0202120112230021-2330211320002031-2020211310301333-1110202220210121-2012310322221200-2012021320102013-2311200031131121-2231203031023301"></a>

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

<a id="canonical-0012001103030201-3203200132211322-3232213203333223-3312110000013103-3010201310230301-0300213202122012-0103012122313002-1023213201333300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.default_loadbalancer` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_custom](data-sources--workload--reference--group-016.md#canonical-1111211210302023-1130301301031112-0133023021111122-2200322020203101-0223203003301101-2022121313331313-1323300011032311-0202020122011323)
- [stateful_service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-016.md#canonical-0013330003233312-1103212221031032-2233120012231121-3101311130002333-2222211213103202-3221232132020020-1233011032133130-3130112100210221)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-016.md#canonical-0323333020010100-1132122030202301-0220233131023222-0121213003321102-3000102211200300-2122203321110300-0131101323132000-2220030303220201)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https](data-sources--workload--reference--group-016.md#canonical-1033211310210331-3122131003202312-2011303010012203-1210330231200011-0103302332022003-0002020201221022-0330311011330033-0101322031121003)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.default_loadbalancer

<a id="canonical-2013333302102123-2201103223101032-3132320311301202-2312232233031031-3312300123303332-2310003023110301-1021002201231332-0333001312302301"></a>

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

<a id="canonical-3233221102122003-1130211310313313-2323010000022332-1312131231102103-3101321021113300-1323200323023012-1230101113333022-0202333112310002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.disable_path_normalize` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_custom](data-sources--workload--reference--group-016.md#canonical-1111211210302023-1130301301031112-0133023021111122-2200322020203101-0223203003301101-2022121313331313-1323300011032311-0202020122011323)
- [stateful_service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-016.md#canonical-0013330003233312-1103212221031032-2233120012231121-3101311130002333-2222211213103202-3221232132020020-1233011032133130-3130112100210221)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-016.md#canonical-0323333020010100-1132122030202301-0220233131023222-0121213003321102-3000102211200300-2122203321110300-0131101323132000-2220030303220201)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https](data-sources--workload--reference--group-016.md#canonical-1033211310210331-3122131003202312-2011303010012203-1210330231200011-0103302332022003-0002020201221022-0330311011330033-0101322031121003)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.disable_path_normalize

<a id="canonical-3130333013230322-3321013313020312-1111133211031213-0130213030223330-3022303322110200-2113012212113210-2230000123112232-3113222101031010"></a>

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

<a id="canonical-2013223221202330-3132232032332332-1132202330202010-3210121032132010-2033202003222023-1321220120101313-2201212323132112-0331002302300020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.enable_path_normalize` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_custom](data-sources--workload--reference--group-016.md#canonical-1111211210302023-1130301301031112-0133023021111122-2200322020203101-0223203003301101-2022121313331313-1323300011032311-0202020122011323)
- [stateful_service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-016.md#canonical-0013330003233312-1103212221031032-2233120012231121-3101311130002333-2222211213103202-3221232132020020-1233011032133130-3130112100210221)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-016.md#canonical-0323333020010100-1132122030202301-0220233131023222-0121213003321102-3000102211200300-2122203321110300-0131101323132000-2220030303220201)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https](data-sources--workload--reference--group-016.md#canonical-1033211310210331-3122131003202312-2011303010012203-1210330231200011-0103302332022003-0002020201221022-0330311011330033-0101322031121003)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.enable_path_normalize

<a id="canonical-3203102120213132-3220323212200121-3002102321212130-2331013210130220-3003122323301013-1311123222311121-3310310201023330-0103031112103101"></a>

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

<a id="canonical-1120222011000333-0211100013303031-0300010010332233-1031300033323310-3030100021331110-1303033230310102-2010030333330102-1101231301332331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_custom](data-sources--workload--reference--group-016.md#canonical-1111211210302023-1130301301031112-0133023021111122-2200322020203101-0223203003301101-2022121313331313-1323300011032311-0202020122011323)
- [stateful_service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-016.md#canonical-0013330003233312-1103212221031032-2233120012231121-3101311130002333-2222211213103202-3221232132020020-1233011032133130-3130112100210221)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-016.md#canonical-0323333020010100-1132122030202301-0220233131023222-0121213003321102-3000102211200300-2122203321110300-0131101323132000-2220030303220201)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https](data-sources--workload--reference--group-016.md#canonical-1033211310210331-3122131003202312-2011303010012203-1210330231200011-0103302332022003-0002020201221022-0330311011330033-0101322031121003)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options

<a id="canonical-1100300322212023-1103310203022101-2313321131110032-2222221311201103-2112212101123111-3210202301113001-0200300103232121-3203033103021131"></a>

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

<a id="canonical-1112122211012233-3131010221010231-1323020210131131-3203333320011112-2213331322122032-0020200030211021-1313233102131300-3010201113103000"></a>

### Direct properties for `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options`

- [http_protocol_enable_v1_only](data-sources--workload--reference--group-016.md#canonical-2030232200121110-2100031231111330-2023320230132100-0300321031312321-0222130321013122-2202030033211322-1330310030303102-2312021102321131): complete subsection reference.

- [http_protocol_enable_v1_v2](data-sources--workload--reference--group-016.md#canonical-1120022312332203-3300020131303230-3333113203222313-3332321311320301-1103311333100132-3320203113100031-0220132331312223-1013030020310030): complete subsection reference.

- [http_protocol_enable_v2_only](data-sources--workload--reference--group-016.md#canonical-2002032323000231-2311320210121132-0203220023130321-0313122310230003-2230102333222001-1330130233002122-1002200201123013-1232113113210320): complete subsection reference.

<a id="canonical-2030232200121110-2100031231111330-2023320230132100-0300321031312321-0222130321013122-2202030033211322-1330310030303102-2312021102321131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_custom](data-sources--workload--reference--group-016.md#canonical-1111211210302023-1130301301031112-0133023021111122-2200322020203101-0223203003301101-2022121313331313-1323300011032311-0202020122011323)
- [stateful_service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-016.md#canonical-0013330003233312-1103212221031032-2233120012231121-3101311130002333-2222211213103202-3221232132020020-1233011032133130-3130112100210221)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-016.md#canonical-0323333020010100-1132122030202301-0220233131023222-0121213003321102-3000102211200300-2122203321110300-0131101323132000-2220030303220201)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https](data-sources--workload--reference--group-016.md#canonical-1033211310210331-3122131003202312-2011303010012203-1210330231200011-0103302332022003-0002020201221022-0330311011330033-0101322031121003)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options](data-sources--workload--reference--group-016.md#canonical-1120222011000333-0211100013303031-0300010010332233-1031300033323310-3030100021331110-1303033230310102-2010030333330102-1101231301332331)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only

<a id="canonical-3323323110312321-2123331132323201-3331213102132330-2013313031213323-2311233133032300-1200022312132311-3132202032030231-2212332301110003"></a>

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

<a id="canonical-1131021010220003-3331013100303122-2121202221223023-3322002303021231-2220202001121302-2133012212111120-1311030121032331-1121101131302011"></a>

### Direct properties for `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only`

- [header_transformation](data-sources--workload--reference--group-016.md#canonical-2200203320303131-2023111232202232-1211300022202210-1000312310010200-2031131200010101-0320121333322101-0022110332101033-1213011011100300): complete subsection reference.

<a id="canonical-2200203320303131-2023111232202232-1211300022202210-1000312310010200-2031131200010101-0320121333322101-0022110332101033-1213011011100300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_custom](data-sources--workload--reference--group-016.md#canonical-1111211210302023-1130301301031112-0133023021111122-2200322020203101-0223203003301101-2022121313331313-1323300011032311-0202020122011323)
- [stateful_service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-016.md#canonical-0013330003233312-1103212221031032-2233120012231121-3101311130002333-2222211213103202-3221232132020020-1233011032133130-3130112100210221)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-016.md#canonical-0323333020010100-1132122030202301-0220233131023222-0121213003321102-3000102211200300-2122203321110300-0131101323132000-2220030303220201)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https](data-sources--workload--reference--group-016.md#canonical-1033211310210331-3122131003202312-2011303010012203-1210330231200011-0103302332022003-0002020201221022-0330311011330033-0101322031121003)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options](data-sources--workload--reference--group-016.md#canonical-1120222011000333-0211100013303031-0300010010332233-1031300033323310-3030100021331110-1303033230310102-2010030333330102-1101231301332331)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only](data-sources--workload--reference--group-016.md#canonical-2030232200121110-2100031231111330-2023320230132100-0300321031312321-0222130321013122-2202030033211322-1330310030303102-2312021102321131)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation

<a id="canonical-0012320031220201-0103233013103012-3131300313101000-0323101131102232-1031012112012111-1211213012320221-0020330123300121-2231223122200223"></a>

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

<a id="canonical-3212003331213311-2020021303323013-0121020322101313-1022311230220130-3310011310300002-1132111311013103-1300122311033230-0112331213330202"></a>

### Direct properties for `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation`

- [default_header_transformation](data-sources--workload--reference--group-016.md#canonical-1320331202000003-1323303202003022-3131303223233211-3111011131200010-0133122002030230-3201201100212023-2222333220223011-1020300112200221): complete subsection reference.

- [preserve_case_header_transformation](data-sources--workload--reference--group-016.md#canonical-0103212202233220-2332113321232210-1122020103202310-3220032001031210-0012332323221332-1321133202231220-0211003210100001-1233202320113332): complete subsection reference.

- [proper_case_header_transformation](data-sources--workload--reference--group-016.md#canonical-2133201310333311-3030302202212131-0003112221112301-0223321231302320-2230011022230212-0021221031133000-2303220232003220-2212222003230223): complete subsection reference.

<a id="canonical-1320331202000003-1323303202003022-3131303223233211-3111011131200010-0133122002030230-3201201100212023-2222333220223011-1020300112200221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.default_header_transformation` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_custom](data-sources--workload--reference--group-016.md#canonical-1111211210302023-1130301301031112-0133023021111122-2200322020203101-0223203003301101-2022121313331313-1323300011032311-0202020122011323)
- [stateful_service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-016.md#canonical-0013330003233312-1103212221031032-2233120012231121-3101311130002333-2222211213103202-3221232132020020-1233011032133130-3130112100210221)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-016.md#canonical-0323333020010100-1132122030202301-0220233131023222-0121213003321102-3000102211200300-2122203321110300-0131101323132000-2220030303220201)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https](data-sources--workload--reference--group-016.md#canonical-1033211310210331-3122131003202312-2011303010012203-1210330231200011-0103302332022003-0002020201221022-0330311011330033-0101322031121003)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options](data-sources--workload--reference--group-016.md#canonical-1120222011000333-0211100013303031-0300010010332233-1031300033323310-3030100021331110-1303033230310102-2010030333330102-1101231301332331)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only](data-sources--workload--reference--group-016.md#canonical-2030232200121110-2100031231111330-2023320230132100-0300321031312321-0222130321013122-2202030033211322-1330310030303102-2312021102321131)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation](data-sources--workload--reference--group-016.md#canonical-2200203320303131-2023111232202232-1211300022202210-1000312310010200-2031131200010101-0320121333322101-0022110332101033-1213011011100300)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.default_header_transformation

<a id="canonical-1113313133310002-1223100012302300-0213032030332132-0031221100333112-2013123010303031-1033013030323303-1022230213221010-2000202021232112"></a>

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

<a id="canonical-0103212202233220-2332113321232210-1122020103202310-3220032001031210-0012332323221332-1321133202231220-0211003210100001-1233202320113332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.preserve_case_header_transformation` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_custom](data-sources--workload--reference--group-016.md#canonical-1111211210302023-1130301301031112-0133023021111122-2200322020203101-0223203003301101-2022121313331313-1323300011032311-0202020122011323)
- [stateful_service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-016.md#canonical-0013330003233312-1103212221031032-2233120012231121-3101311130002333-2222211213103202-3221232132020020-1233011032133130-3130112100210221)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-016.md#canonical-0323333020010100-1132122030202301-0220233131023222-0121213003321102-3000102211200300-2122203321110300-0131101323132000-2220030303220201)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https](data-sources--workload--reference--group-016.md#canonical-1033211310210331-3122131003202312-2011303010012203-1210330231200011-0103302332022003-0002020201221022-0330311011330033-0101322031121003)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options](data-sources--workload--reference--group-016.md#canonical-1120222011000333-0211100013303031-0300010010332233-1031300033323310-3030100021331110-1303033230310102-2010030333330102-1101231301332331)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only](data-sources--workload--reference--group-016.md#canonical-2030232200121110-2100031231111330-2023320230132100-0300321031312321-0222130321013122-2202030033211322-1330310030303102-2312021102321131)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation](data-sources--workload--reference--group-016.md#canonical-2200203320303131-2023111232202232-1211300022202210-1000312310010200-2031131200010101-0320121333322101-0022110332101033-1213011011100300)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.preserve_case_header_transformation

<a id="canonical-1210302133332333-3013010011032123-0102022120011010-1300303021111001-2013321232010311-0333000313121100-1330100332213000-3212000112110211"></a>

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

<a id="canonical-2133201310333311-3030302202212131-0003112221112301-0223321231302320-2230011022230212-0021221031133000-2303220232003220-2212222003230223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.proper_case_header_transformation` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_custom](data-sources--workload--reference--group-016.md#canonical-1111211210302023-1130301301031112-0133023021111122-2200322020203101-0223203003301101-2022121313331313-1323300011032311-0202020122011323)
- [stateful_service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-016.md#canonical-0013330003233312-1103212221031032-2233120012231121-3101311130002333-2222211213103202-3221232132020020-1233011032133130-3130112100210221)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-016.md#canonical-0323333020010100-1132122030202301-0220233131023222-0121213003321102-3000102211200300-2122203321110300-0131101323132000-2220030303220201)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https](data-sources--workload--reference--group-016.md#canonical-1033211310210331-3122131003202312-2011303010012203-1210330231200011-0103302332022003-0002020201221022-0330311011330033-0101322031121003)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options](data-sources--workload--reference--group-016.md#canonical-1120222011000333-0211100013303031-0300010010332233-1031300033323310-3030100021331110-1303033230310102-2010030333330102-1101231301332331)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only](data-sources--workload--reference--group-016.md#canonical-2030232200121110-2100031231111330-2023320230132100-0300321031312321-0222130321013122-2202030033211322-1330310030303102-2312021102321131)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation](data-sources--workload--reference--group-016.md#canonical-2200203320303131-2023111232202232-1211300022202210-1000312310010200-2031131200010101-0320121333322101-0022110332101033-1213011011100300)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.proper_case_header_transformation

<a id="canonical-3311112121000322-0002013333033031-2203123223101302-3313011003021031-3301121033212320-2030122030303120-1332130113020201-3121001120200223"></a>

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

<a id="canonical-1120022312332203-3300020131303230-3333113203222313-3332321311320301-1103311333100132-3320203113100031-0220132331312223-1013030020310030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_v2` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_custom](data-sources--workload--reference--group-016.md#canonical-1111211210302023-1130301301031112-0133023021111122-2200322020203101-0223203003301101-2022121313331313-1323300011032311-0202020122011323)
- [stateful_service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-016.md#canonical-0013330003233312-1103212221031032-2233120012231121-3101311130002333-2222211213103202-3221232132020020-1233011032133130-3130112100210221)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-016.md#canonical-0323333020010100-1132122030202301-0220233131023222-0121213003321102-3000102211200300-2122203321110300-0131101323132000-2220030303220201)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https](data-sources--workload--reference--group-016.md#canonical-1033211310210331-3122131003202312-2011303010012203-1210330231200011-0103302332022003-0002020201221022-0330311011330033-0101322031121003)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options](data-sources--workload--reference--group-016.md#canonical-1120222011000333-0211100013303031-0300010010332233-1031300033323310-3030100021331110-1303033230310102-2010030333330102-1101231301332331)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_v2

<a id="canonical-0313312211121013-0303301122102023-0333310001103331-1100222201011302-0031322212302031-2031011100123002-0020112133132030-2113332322211100"></a>

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

<a id="canonical-2002032323000231-2311320210121132-0203220023130321-0313122310230003-2230102333222001-1330130233002122-1002200201123013-1232113113210320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v2_only` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_custom](data-sources--workload--reference--group-016.md#canonical-1111211210302023-1130301301031112-0133023021111122-2200322020203101-0223203003301101-2022121313331313-1323300011032311-0202020122011323)
- [stateful_service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-016.md#canonical-0013330003233312-1103212221031032-2233120012231121-3101311130002333-2222211213103202-3221232132020020-1233011032133130-3130112100210221)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-016.md#canonical-0323333020010100-1132122030202301-0220233131023222-0121213003321102-3000102211200300-2122203321110300-0131101323132000-2220030303220201)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https](data-sources--workload--reference--group-016.md#canonical-1033211310210331-3122131003202312-2011303010012203-1210330231200011-0103302332022003-0002020201221022-0330311011330033-0101322031121003)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options](data-sources--workload--reference--group-016.md#canonical-1120222011000333-0211100013303031-0300010010332233-1031300033323310-3030100021331110-1303033230310102-2010030333330102-1101231301332331)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v2_only

<a id="canonical-3321201002311000-2312033122300100-0222021202331100-3300133203023233-2200130031310300-0330221103323000-0133013221123030-3001231231301233"></a>

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

<a id="canonical-2233212220331030-3333132301233302-1131111310300010-1020030131001213-2022003333122002-1023230213320033-0102023112233011-0323230323302323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.non_default_loadbalancer` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_custom](data-sources--workload--reference--group-016.md#canonical-1111211210302023-1130301301031112-0133023021111122-2200322020203101-0223203003301101-2022121313331313-1323300011032311-0202020122011323)
- [stateful_service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-016.md#canonical-0013330003233312-1103212221031032-2233120012231121-3101311130002333-2222211213103202-3221232132020020-1233011032133130-3130112100210221)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-016.md#canonical-0323333020010100-1132122030202301-0220233131023222-0121213003321102-3000102211200300-2122203321110300-0131101323132000-2220030303220201)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https](data-sources--workload--reference--group-016.md#canonical-1033211310210331-3122131003202312-2011303010012203-1210330231200011-0103302332022003-0002020201221022-0330311011330033-0101322031121003)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.non_default_loadbalancer

<a id="canonical-3330003133011223-3112020013033023-2002002102013021-1213101101132002-3101213003111322-0113031032331103-3203331103313303-0011203102231202"></a>

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

<a id="canonical-3210111122032323-1101022110110331-1023201230033121-2011003332013100-1002021032212313-0320000121303232-3101312000213323-3232133323310302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.pass_through` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_custom](data-sources--workload--reference--group-016.md#canonical-1111211210302023-1130301301031112-0133023021111122-2200322020203101-0223203003301101-2022121313331313-1323300011032311-0202020122011323)
- [stateful_service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-016.md#canonical-0013330003233312-1103212221031032-2233120012231121-3101311130002333-2222211213103202-3221232132020020-1233011032133130-3130112100210221)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-016.md#canonical-0323333020010100-1132122030202301-0220233131023222-0121213003321102-3000102211200300-2122203321110300-0131101323132000-2220030303220201)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https](data-sources--workload--reference--group-016.md#canonical-1033211310210331-3122131003202312-2011303010012203-1210330231200011-0103302332022003-0002020201221022-0330311011330033-0101322031121003)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.pass_through

<a id="canonical-2312010032331130-1210222230031201-3032232120231302-0321133213310232-0230322132300010-1012313122302032-1223001333310013-0123033312020002"></a>

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

<a id="canonical-2010001011230201-3020111320131112-2333021133221211-0200112033223200-3003002233220012-1123230332003001-0103023230030232-1202122133301320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_custom](data-sources--workload--reference--group-016.md#canonical-1111211210302023-1130301301031112-0133023021111122-2200322020203101-0223203003301101-2022121313331313-1323300011032311-0202020122011323)
- [stateful_service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-016.md#canonical-0013330003233312-1103212221031032-2233120012231121-3101311130002333-2222211213103202-3221232132020020-1233011032133130-3130112100210221)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-016.md#canonical-0323333020010100-1132122030202301-0220233131023222-0121213003321102-3000102211200300-2122203321110300-0131101323132000-2220030303220201)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https](data-sources--workload--reference--group-016.md#canonical-1033211310210331-3122131003202312-2011303010012203-1210330231200011-0103302332022003-0002020201221022-0330311011330033-0101322031121003)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params

<a id="canonical-3230201021031222-0310032013203213-2203000010021020-1220000320102013-3212132122201031-2132122013231101-1011212010121331-3311321231323020"></a>

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

<a id="canonical-1000221210012322-2313211113211311-1320011221022300-0031301022133302-2212212312232320-2320012031220231-3310131112020123-3301300220003111"></a>

### Direct properties for `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params`

- [certificates](data-sources--workload--reference--group-016.md#canonical-0200320033213133-0020001321201233-0202213032020213-0102012010121312-3222330131202032-1103130023001201-2333113032201212-2121120212121202): complete subsection reference.

- [no_mtls](data-sources--workload--reference--group-016.md#canonical-0132330310010120-1311211000130301-2012002111302230-0322000022133313-1213120133132200-0102110202300300-1102212333033030-3323221220033122): complete subsection reference.

- [tls_config](data-sources--workload--reference--group-016.md#canonical-2232320011300222-0303022031002312-1022321103212313-0133003330203133-1221203232301103-2221122022022121-0222033321201310-0221330023302121): complete subsection reference.

- [use_mtls](data-sources--workload--reference--group-017.md#canonical-2313300102011220-2013300230111223-1221213101333022-3233013120131303-1012130321112222-0033102323232122-3133200002310023-0033322000223103): complete subsection reference.

<a id="canonical-0200320033213133-0020001321201233-0202213032020213-0102012010121312-3222330131202032-1103130023001201-2333113032201212-2121120212121202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.certificates` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_custom](data-sources--workload--reference--group-016.md#canonical-1111211210302023-1130301301031112-0133023021111122-2200322020203101-0223203003301101-2022121313331313-1323300011032311-0202020122011323)
- [stateful_service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-016.md#canonical-0013330003233312-1103212221031032-2233120012231121-3101311130002333-2222211213103202-3221232132020020-1233011032133130-3130112100210221)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-016.md#canonical-0323333020010100-1132122030202301-0220233131023222-0121213003321102-3000102211200300-2122203321110300-0131101323132000-2220030303220201)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https](data-sources--workload--reference--group-016.md#canonical-1033211310210331-3122131003202312-2011303010012203-1210330231200011-0103302332022003-0002020201221022-0330311011330033-0101322031121003)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params](data-sources--workload--reference--group-016.md#canonical-2010001011230201-3020111320131112-2333021133221211-0200112033223200-3003002233220012-1123230332003001-0103023230030232-1202122133301320)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.certificates

<a id="canonical-1121203023303030-3000002122320000-0323111101322111-3121230220111223-1112003110223133-3202222330310203-0012231130101011-2332112110332030"></a>

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-1032231220323320-0222130220223211-0331122023231212-3213223122221001-0332130202223111-1232312300302302-3202020011330321-2120122003031003"></a>

### Direct properties for `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.certificates`

<a id="canonical-1110211203300212-2020212313222001-0313021312021132-2132103203023210-0110322201013322-3311123111031123-2032130022333130-0210002131310200"></a>

#### `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.certificates.name` property

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

<a id="canonical-1032211033000332-0030300013122013-1303303023013312-1033330220032322-0120223303211023-1313130203032101-1321203221131301-0133332331210101"></a>

<a id="canonical-1123033101212320-0032232212213322-1301123112103023-0200230103332220-0010213333221212-1112223203131231-2110111222322000-1313212130232111"></a>

#### `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.certificates.namespace` property

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

<a id="canonical-2122100213301030-0130203320231030-1221222103330023-1210002103101121-0011333313012221-3231203313030322-1220031321212312-1321221203022330"></a>

<a id="canonical-2111200320031100-1303131133103321-3203121120003123-3323203333323233-0323103111222123-0013301212033022-1302023033012012-0332103132221113"></a>

#### `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.certificates.tenant` property

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

<a id="canonical-0132330310010120-1311211000130301-2012002111302230-0322000022133313-1213120133132200-0102110202300300-1102212333033030-3323221220033122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.no_mtls` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_custom](data-sources--workload--reference--group-016.md#canonical-1111211210302023-1130301301031112-0133023021111122-2200322020203101-0223203003301101-2022121313331313-1323300011032311-0202020122011323)
- [stateful_service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-016.md#canonical-0013330003233312-1103212221031032-2233120012231121-3101311130002333-2222211213103202-3221232132020020-1233011032133130-3130112100210221)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-016.md#canonical-0323333020010100-1132122030202301-0220233131023222-0121213003321102-3000102211200300-2122203321110300-0131101323132000-2220030303220201)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https](data-sources--workload--reference--group-016.md#canonical-1033211310210331-3122131003202312-2011303010012203-1210330231200011-0103302332022003-0002020201221022-0330311011330033-0101322031121003)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params](data-sources--workload--reference--group-016.md#canonical-2010001011230201-3020111320131112-2333021133221211-0200112033223200-3003002233220012-1123230332003001-0103023230030232-1202122133301320)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.no_mtls

<a id="canonical-0221313203310003-2012230130030320-0332121002201103-2102221330101210-2312203133111310-1221120320001303-2020101131212122-3333103112200320"></a>

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

<a id="canonical-2232320011300222-0303022031002312-1022321103212313-0133003330203133-1221203232301103-2221122022022121-0222033321201310-0221330023302121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.tls_config` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_custom](data-sources--workload--reference--group-016.md#canonical-1111211210302023-1130301301031112-0133023021111122-2200322020203101-0223203003301101-2022121313331313-1323300011032311-0202020122011323)
- [stateful_service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-016.md#canonical-0013330003233312-1103212221031032-2233120012231121-3101311130002333-2222211213103202-3221232132020020-1233011032133130-3130112100210221)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-016.md#canonical-0323333020010100-1132122030202301-0220233131023222-0121213003321102-3000102211200300-2122203321110300-0131101323132000-2220030303220201)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https](data-sources--workload--reference--group-016.md#canonical-1033211310210331-3122131003202312-2011303010012203-1210330231200011-0103302332022003-0002020201221022-0330311011330033-0101322031121003)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params](data-sources--workload--reference--group-016.md#canonical-2010001011230201-3020111320131112-2333021133221211-0200112033223200-3003002233220012-1123230332003001-0103023230030232-1202122133301320)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.tls_config

<a id="canonical-3302202201230100-3212233023133201-3031003201133200-3101000332002103-1322230012021020-1201303012120120-1222321223300120-3000100121201320"></a>

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

<a id="canonical-1231102301101000-1303003003222101-2101123220013303-2022010323211101-1022111330101131-3233323203321332-0111131323003013-1113033111230001"></a>

### Direct properties for `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.tls_config`

- [custom_security](data-sources--workload--reference--group-016.md#canonical-1121230202210110-3233312123130221-2002030021101033-0023113330320230-0002203130003210-1302322102321210-0103110100331103-3003333101303023): complete subsection reference.

- [default_security](data-sources--workload--reference--group-016.md#canonical-2123321331121002-1301211000221111-2000231021222011-2322230320030013-0212032130032103-2311023331033331-0000112303201220-3230332221010310): complete subsection reference.

- [low_security](data-sources--workload--reference--group-016.md#canonical-0032303100131112-3111010033002211-1233020131302032-2113303232330212-2131322032310132-1220032300100033-3323013003011111-3021000320211213): complete subsection reference.

- [medium_security](data-sources--workload--reference--group-017.md#canonical-1113202323300132-1211133022030123-1210023101301101-2222203120320211-0322323333232303-2012223302312311-0323301033311113-2020223020210020): complete subsection reference.

<a id="canonical-1121230202210110-3233312123130221-2002030021101033-0023113330320230-0002203130003210-1302322102321210-0103110100331103-3003333101303023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.tls_config.custom_security` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_custom](data-sources--workload--reference--group-016.md#canonical-1111211210302023-1130301301031112-0133023021111122-2200322020203101-0223203003301101-2022121313331313-1323300011032311-0202020122011323)
- [stateful_service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-016.md#canonical-0013330003233312-1103212221031032-2233120012231121-3101311130002333-2222211213103202-3221232132020020-1233011032133130-3130112100210221)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-016.md#canonical-0323333020010100-1132122030202301-0220233131023222-0121213003321102-3000102211200300-2122203321110300-0131101323132000-2220030303220201)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https](data-sources--workload--reference--group-016.md#canonical-1033211310210331-3122131003202312-2011303010012203-1210330231200011-0103302332022003-0002020201221022-0330311011330033-0101322031121003)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params](data-sources--workload--reference--group-016.md#canonical-2010001011230201-3020111320131112-2333021133221211-0200112033223200-3003002233220012-1123230332003001-0103023230030232-1202122133301320)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.tls_config](data-sources--workload--reference--group-016.md#canonical-2232320011300222-0303022031002312-1022321103212313-0133003330203133-1221203232301103-2221122022022121-0222033321201310-0221330023302121)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.tls_config.custom_security

<a id="canonical-0233201201120203-3212131210130331-2310123233201221-0321033311220330-1113201000203130-1322112111011230-0022303201321123-2222321312313032"></a>

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

<a id="canonical-1313013121302001-0133322222303213-2002001310210311-2210130111331022-1202123122323321-2032012102211002-3033023133200130-1330223331213031"></a>

### Direct properties for `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.tls_config.custom_security`

<a id="canonical-2122122110301330-1132231133000333-3331123332113113-3333121000012021-2130320111121121-2320002221132202-3132313310113110-2212201122020120"></a>

#### `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.tls_config.custom_security.cipher_suites` property

Type: `["list", "string"]`. Computed.

The TLS listener will only support the specified cipher list.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-10T03:47:03+00:00"
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
    "ves.io.schema.rules.repeated.items.string.in": "[\\\"TLS_AES_128_GCM_SHA256\\\",\\\"TLS_AES_256_GCM_SHA384\\\",\\\"TLS_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384\\\",\\\"TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384\\\",\\\"TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_RSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_256_GCM_SHA384\\\"]",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.in": "[\\\"TLS_AES_128_GCM_SHA256\\\",\\\"TLS_AES_256_GCM_SHA384\\\",\\\"TLS_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384\\\",\\\"TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384\\\",\\\"TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_RSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_256_GCM_SHA384\\\"]",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3000302210201033-3231022231001212-3131332311302322-3221112200323001-1133131030022120-2022012210233221-3211330323211313-2322133001023320"></a>

<a id="canonical-2111111223030330-2020321321101103-1220121333320202-0233213302002303-0203203032001330-0101220211201311-0213230011130011-3110300221011012"></a>

#### `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.tls_config.custom_security.max_version` property

Type: `"string"`. Computed.

\[Enum: TLS\_AUTO|TLSv1\_0|TLSv1\_1|TLSv1\_2|TLSv1\_3\] TlsProtocol is enumeration of supported TLS
versions F5 Distributed Cloud will choose the optimal TLS version. Possible values are
\`TLS\_AUTO\`, \`TLSv1\_0\`, \`TLSv1\_1\`, \`TLSv1\_2\`, \`TLSv1\_3\`. Defaults to \`TLS\_AUTO\`.

Receipt-pinned upstream constraints:

```json
{
  "default": "TLS_AUTO",
  "enum": [
    "TLS_AUTO",
    "TLSv1_0",
    "TLSv1_1",
    "TLSv1_2",
    "TLSv1_3"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3001302112302211-0221003020331200-3001101202101330-2230203120231302-1012330103303133-3310301311301002-0230322101103212-1221330301322220"></a>

<a id="canonical-0302003012301323-3103000212321303-3312333113321232-0332130032211022-0333231233321310-1212123211222010-3303110302023011-1020113120023103"></a>

#### `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.tls_config.custom_security.min_version` property

Type: `"string"`. Computed.

\[Enum: TLS\_AUTO|TLSv1\_0|TLSv1\_1|TLSv1\_2|TLSv1\_3\] TlsProtocol is enumeration of supported TLS
versions F5 Distributed Cloud will choose the optimal TLS version. Possible values are
\`TLS\_AUTO\`, \`TLSv1\_0\`, \`TLSv1\_1\`, \`TLSv1\_2\`, \`TLSv1\_3\`. Defaults to \`TLS\_AUTO\`.

Receipt-pinned upstream constraints:

```json
{
  "default": "TLS_AUTO",
  "enum": [
    "TLS_AUTO",
    "TLSv1_0",
    "TLSv1_1",
    "TLSv1_2",
    "TLSv1_3"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2123321331121002-1301211000221111-2000231021222011-2322230320030013-0212032130032103-2311023331033331-0000112303201220-3230332221010310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.tls_config.default_security` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_custom](data-sources--workload--reference--group-016.md#canonical-1111211210302023-1130301301031112-0133023021111122-2200322020203101-0223203003301101-2022121313331313-1323300011032311-0202020122011323)
- [stateful_service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-016.md#canonical-0013330003233312-1103212221031032-2233120012231121-3101311130002333-2222211213103202-3221232132020020-1233011032133130-3130112100210221)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-016.md#canonical-0323333020010100-1132122030202301-0220233131023222-0121213003321102-3000102211200300-2122203321110300-0131101323132000-2220030303220201)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https](data-sources--workload--reference--group-016.md#canonical-1033211310210331-3122131003202312-2011303010012203-1210330231200011-0103302332022003-0002020201221022-0330311011330033-0101322031121003)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params](data-sources--workload--reference--group-016.md#canonical-2010001011230201-3020111320131112-2333021133221211-0200112033223200-3003002233220012-1123230332003001-0103023230030232-1202122133301320)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.tls_config](data-sources--workload--reference--group-016.md#canonical-2232320011300222-0303022031002312-1022321103212313-0133003330203133-1221203232301103-2221122022022121-0222033321201310-0221330023302121)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.tls_config.default_security

<a id="canonical-2330133231322101-2332323013202021-3300322000012012-2030012201232201-2130002202002222-0233010211320202-3131110313112211-3111332212030310"></a>

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

<a id="canonical-0032303100131112-3111010033002211-1233020131302032-2113303232330212-2131322032310132-1220032300100033-3323013003011111-3021000320211213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.tls_config.low_security` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_custom](data-sources--workload--reference--group-016.md#canonical-1111211210302023-1130301301031112-0133023021111122-2200322020203101-0223203003301101-2022121313331313-1323300011032311-0202020122011323)
- [stateful_service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-016.md#canonical-0013330003233312-1103212221031032-2233120012231121-3101311130002333-2222211213103202-3221232132020020-1233011032133130-3130112100210221)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-016.md#canonical-0323333020010100-1132122030202301-0220233131023222-0121213003321102-3000102211200300-2122203321110300-0131101323132000-2220030303220201)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https](data-sources--workload--reference--group-016.md#canonical-1033211310210331-3122131003202312-2011303010012203-1210330231200011-0103302332022003-0002020201221022-0330311011330033-0101322031121003)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params](data-sources--workload--reference--group-016.md#canonical-2010001011230201-3020111320131112-2333021133221211-0200112033223200-3003002233220012-1123230332003001-0103023230030232-1202122133301320)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.tls_config](data-sources--workload--reference--group-016.md#canonical-2232320011300222-0303022031002312-1022321103212313-0133003330203133-1221203232301103-2221122022022121-0222033321201310-0221330023302121)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.tls_config.low_security

<a id="canonical-0032130023022330-2200230200201321-2003212020300000-2331221130121231-2333033333313001-3101112033100120-2233231032110133-2101320002223011"></a>

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
