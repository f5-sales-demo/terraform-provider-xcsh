---
page_title: "xcsh_workload reference"
subcategory: "Container"
description: "Complete grouped canonical reference for xcsh_workload reference."
---

# xcsh_workload reference

<a id="canonical-2122211231322113-0321010120020121-2320333323010130-3302111003330012-3233300100130133-3212112333301203-1100133201311012-2011321220012100"></a>

## Direct properties for `job.containers.liveness_check.http_health_check`

<a id="canonical-1311231333220320-0221313220111131-3312220201133303-2232132331320201-0321321110201333-2102100231000121-3330300333231320-0002133221012011"></a>

### `job.containers.liveness_check.http_health_check.headers` property

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

<a id="canonical-0133123333123220-0310123121300020-3230231122320300-0133021032320210-3133032001110332-2213331330113031-3320132320102110-2301212111132122"></a>

<a id="canonical-3103310113020133-3323011110012230-2111002310311012-0213012001023102-0331030223212100-1321320300331311-3123233300301223-0233010131210112"></a>

### `job.containers.liveness_check.http_health_check.host_header` property

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

<a id="canonical-1020100021022302-0121330220320112-3131100113030032-0213310230223010-3322013211303002-1233220320130333-2000302121032120-0031110230323011"></a>

<a id="canonical-1031231012302210-0302013133231301-1132120013202211-1111110010111301-2001320222200332-2222112020011213-3311202311301122-3032110110323313"></a>

### `job.containers.liveness_check.http_health_check.path` property

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

- [port](resources--workload--reference--group-005.md#canonical-1321311011101302-0312010032213201-0233111222302031-2302013031332330-0202212332333210-3311202311100333-0322211021330003-0022312120322202): complete subsection reference.

<a id="canonical-1321311011101302-0312010032213201-0233111222302031-2302013031332330-0202212332333210-3311202311100333-0322211021330003-0022312120322202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `job.containers.liveness_check.http_health_check.port` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [job](resources--workload--reference--group-004.md#canonical-2101133223131122-0231303130011203-2130203200233301-2332212033103313-3210300031001001-1312100313213300-1013120312121230-0320112032220311)
- [job.containers](resources--workload--reference--group-004.md#canonical-2002301323122233-0221102100003021-3013132233211233-3012203030111110-1222020112111222-3222330102102021-2002120323232230-2011302203232120)
- [job.containers.liveness_check](resources--workload--reference--group-004.md#canonical-1212102312001320-2330002220103201-0331323102112032-0231200202213223-0111102010022023-3221030302200030-0103233001300201-2012131231021122)
- [job.containers.liveness_check.http_health_check](resources--workload--reference--group-004.md#canonical-3330020102131300-1013303302110121-2033120022321212-0020233021130030-2220201230022113-0331011203200210-0201231102023020-0302333022122313)
- job.containers.liveness_check.http_health_check.port

<a id="canonical-1232322213032310-1000132233131303-2232010301303212-2121211221101021-2032031300001310-3311101210003133-1131111230001120-1112030102330020"></a>

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

<a id="canonical-3211232133322130-1030031221220230-3331233103213332-2202221011202302-1220220000320112-3033002202200121-1001220113210220-0120311323221131"></a>

### Direct properties for `job.containers.liveness_check.http_health_check.port`

<a id="canonical-1013100032113021-2011233013122103-2220101220202033-3222332322111303-2012030030032313-1000020210332301-2302111311113112-0200213330101210"></a>

#### `job.containers.liveness_check.http_health_check.port.name` property

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

<a id="canonical-1121103321112030-3022333020020112-3333230110313022-2003233330030311-1303330003322200-3112122231031112-2211332032020300-1023203120321031"></a>

<a id="canonical-2323023003201223-1101332202222330-0101201011203023-0133232000020222-0112313203322133-3131033313111030-0303301133032212-3221332002303322"></a>

#### `job.containers.liveness_check.http_health_check.port.num` property

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

<a id="canonical-1223231110101332-2322020122131330-1123010210221202-3020222322221232-1201100102333333-3011310033110012-1312221031213301-1100110320100300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `job.containers.liveness_check.tcp_health_check` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [job](resources--workload--reference--group-004.md#canonical-2101133223131122-0231303130011203-2130203200233301-2332212033103313-3210300031001001-1312100313213300-1013120312121230-0320112032220311)
- [job.containers](resources--workload--reference--group-004.md#canonical-2002301323122233-0221102100003021-3013132233211233-3012203030111110-1222020112111222-3222330102102021-2002120323232230-2011302203232120)
- [job.containers.liveness_check](resources--workload--reference--group-004.md#canonical-1212102312001320-2330002220103201-0331323102112032-0231200202213223-0111102010022023-3221030302200030-0103233001300201-2012131231021122)
- job.containers.liveness_check.tcp_health_check

<a id="canonical-0133120230323122-2130020203320120-0201130231210100-1333303232312103-2313101323312133-1131322212013332-1233330221213311-3201102230313213"></a>

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

<a id="canonical-2231230301210300-2022020330231320-1221100022122101-2202122023213021-0311132120232221-0231010223322210-0321111220212332-2302000001020311"></a>

### Direct properties for `job.containers.liveness_check.tcp_health_check`

- [port](resources--workload--reference--group-005.md#canonical-1213111300230323-2301130032101021-1112110311230100-1211212322321301-2201230303022033-3011312311212121-0310311103333021-3103013013330332): complete subsection reference.

<a id="canonical-1213111300230323-2301130032101021-1112110311230100-1211212322321301-2201230303022033-3011312311212121-0310311103333021-3103013013330332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `job.containers.liveness_check.tcp_health_check.port` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [job](resources--workload--reference--group-004.md#canonical-2101133223131122-0231303130011203-2130203200233301-2332212033103313-3210300031001001-1312100313213300-1013120312121230-0320112032220311)
- [job.containers](resources--workload--reference--group-004.md#canonical-2002301323122233-0221102100003021-3013132233211233-3012203030111110-1222020112111222-3222330102102021-2002120323232230-2011302203232120)
- [job.containers.liveness_check](resources--workload--reference--group-004.md#canonical-1212102312001320-2330002220103201-0331323102112032-0231200202213223-0111102010022023-3221030302200030-0103233001300201-2012131231021122)
- [job.containers.liveness_check.tcp_health_check](resources--workload--reference--group-005.md#canonical-1223231110101332-2322020122131330-1123010210221202-3020222322221232-1201100102333333-3011310033110012-1312221031213301-1100110320100300)
- job.containers.liveness_check.tcp_health_check.port

<a id="canonical-3213211100313031-2000000202321133-3302121013212112-3212200330012000-0301230211212333-2203113021221032-2321012003001213-0221303202312302"></a>

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

<a id="canonical-3311200100200120-3330001213331230-3100031132323131-2333332300223123-0023021001021010-1233032020013202-0030013213000033-2230020332321011"></a>

### Direct properties for `job.containers.liveness_check.tcp_health_check.port`

<a id="canonical-0210122011002100-0030112021320333-2120300113302302-2302320030233113-2111321313233312-1002021332332020-2310033330012323-2213011332332011"></a>

#### `job.containers.liveness_check.tcp_health_check.port.name` property

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

<a id="canonical-3130202213022101-3320013333031312-0021023122320121-2133203222010133-1103203113230331-0101213100020301-2301020033033303-1323302300331021"></a>

<a id="canonical-3122133022213113-1301030113002110-1112303313332303-1211011023020203-0211220003300320-1033222103102101-1213300310113310-3302211310230212"></a>

#### `job.containers.liveness_check.tcp_health_check.port.num` property

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

<a id="canonical-0222311231110223-3301213322101001-0122100312221013-1333033012100312-0133231033110100-3131310123021323-2322322213211230-3110310102123121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `job.containers.readiness_check` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [job](resources--workload--reference--group-004.md#canonical-2101133223131122-0231303130011203-2130203200233301-2332212033103313-3210300031001001-1312100313213300-1013120312121230-0320112032220311)
- [job.containers](resources--workload--reference--group-004.md#canonical-2002301323122233-0221102100003021-3013132233211233-3012203030111110-1222020112111222-3222330102102021-2002120323232230-2011302203232120)
- job.containers.readiness_check

<a id="canonical-0111121221312133-1101103100033312-1032311132301211-3232300233033120-0111013111302123-3312122130232323-1101102121202210-3310303012221133"></a>

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

<a id="canonical-2222030330222333-0302002020303312-2303010113313200-0232031213330013-2220120123301303-1312132331313301-2002311121011101-1110233132120100"></a>

### Direct properties for `job.containers.readiness_check`

- [exec_health_check](resources--workload--reference--group-005.md#canonical-2330302322132120-0133031312222322-1020001300011202-2222223031212223-0200031223212121-0130302130222100-1231320103231212-3000101133002221): complete subsection reference.

<a id="canonical-2001311320000332-0222112120212332-0322122003202312-2302112003310031-2220102110032210-3132102030112031-2121100311010212-1332112233301201"></a>

<a id="canonical-1320211132232311-3311022322322333-0011331322300130-3331101301331320-0233011003331111-3313121322232133-1022122311321130-1033113223030232"></a>

#### `job.containers.readiness_check.healthy_threshold` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

- [http_health_check](resources--workload--reference--group-005.md#canonical-0233132103213321-3313220020312110-2101103210301101-0310230020133310-0222131330330222-3011213321020020-3122230111300132-0313022002201010): complete subsection reference.

<a id="canonical-2022212302023111-1300111320120111-3102203321212102-2013102213201323-1213200302212223-1022012030212103-0222223113121113-2133131201121133"></a>

<a id="canonical-1221113303311302-3320122011330200-3100202022030203-3102210020133301-1030311202302213-1110010011122311-0333201232332310-2300231310230323"></a>

#### `job.containers.readiness_check.initial_delay` property

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
    "ves.io.schema.rules.uint32.lte": "600"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "600"
  }
}
```

<a id="canonical-2020221333122321-2112310111111322-2112123332102132-0233312311210213-2231012003320230-2313331232220322-0011231321232312-1311013313310112"></a>

<a id="canonical-2232232222332230-1100312221213021-2331332303132301-1211222120032133-3030110212012311-3300210312121122-0303202321200102-2320213311200202"></a>

#### `job.containers.readiness_check.interval` property

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

- [tcp_health_check](resources--workload--reference--group-005.md#canonical-1331303311323011-1231220102032100-0033122102322121-0323121321110233-3230223220133313-1001001130132012-1012102132021112-3000233323110222): complete subsection reference.

<a id="canonical-2302021121203313-1202202102132123-0303000010101220-0233000020103013-1021130100333230-0322313123303033-2331303222110001-3330320133121131"></a>

<a id="canonical-2232332110322221-3303122313202310-3301321221301202-0230120021333231-2212230133310103-3103001300013002-1211121221023013-1202333023131333"></a>

#### `job.containers.readiness_check.timeout` property

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

<a id="canonical-0230232232300313-2331220323011211-2221200222112320-2333000022323211-1202000331323110-2033030102012321-1102023322013201-2300012023122120"></a>

<a id="canonical-3321132330131020-3311211120110033-3003201330121211-2033232013222022-2022000231101202-1011201221123121-2201111121030313-0032320332200311"></a>

#### `job.containers.readiness_check.unhealthy_threshold` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-2330302322132120-0133031312222322-1020001300011202-2222223031212223-0200031223212121-0130302130222100-1231320103231212-3000101133002221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `job.containers.readiness_check.exec_health_check` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [job](resources--workload--reference--group-004.md#canonical-2101133223131122-0231303130011203-2130203200233301-2332212033103313-3210300031001001-1312100313213300-1013120312121230-0320112032220311)
- [job.containers](resources--workload--reference--group-004.md#canonical-2002301323122233-0221102100003021-3013132233211233-3012203030111110-1222020112111222-3222330102102021-2002120323232230-2011302203232120)
- [job.containers.readiness_check](resources--workload--reference--group-005.md#canonical-0222311231110223-3301213322101001-0122100312221013-1333033012100312-0133231033110100-3131310123021323-2322322213211230-3110310102123121)
- job.containers.readiness_check.exec_health_check

<a id="canonical-1110233002123001-0121201302022113-1012110121122212-2222002201210120-0020302113220232-2313201322201203-3211210112210300-3321133322321010"></a>

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

<a id="canonical-0202021310312202-2323333301233230-0313300110020110-0013201212332120-1230223333220002-1001303032321031-3210002322113230-1323032223123101"></a>

### Direct properties for `job.containers.readiness_check.exec_health_check`

<a id="canonical-2000020120330313-3230012012232012-3301103012302223-1230020321210033-0331012133101102-1232230021332112-3201203221030212-0012211010212010"></a>

#### `job.containers.readiness_check.exec_health_check.command` property

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

<a id="canonical-0233132103213321-3313220020312110-2101103210301101-0310230020133310-0222131330330222-3011213321020020-3122230111300132-0313022002201010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `job.containers.readiness_check.http_health_check` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [job](resources--workload--reference--group-004.md#canonical-2101133223131122-0231303130011203-2130203200233301-2332212033103313-3210300031001001-1312100313213300-1013120312121230-0320112032220311)
- [job.containers](resources--workload--reference--group-004.md#canonical-2002301323122233-0221102100003021-3013132233211233-3012203030111110-1222020112111222-3222330102102021-2002120323232230-2011302203232120)
- [job.containers.readiness_check](resources--workload--reference--group-005.md#canonical-0222311231110223-3301213322101001-0122100312221013-1333033012100312-0133231033110100-3131310123021323-2322322213211230-3110310102123121)
- job.containers.readiness_check.http_health_check

<a id="canonical-1012100031113031-1223210113102200-0303200130321200-1330133321333123-1312131312301110-3102130232101213-1102010002202103-1313311300112311"></a>

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

<a id="canonical-3103221023300232-2100311013102132-3122020131111032-2031100123201232-0100232200031323-1131332330203321-3212233013012023-0012023333113113"></a>

### Direct properties for `job.containers.readiness_check.http_health_check`

<a id="canonical-2302131102021030-0332012003300123-1032132023133000-2130201013212100-0001231310110032-3313332013302022-2101222033302010-1312223030012032"></a>

#### `job.containers.readiness_check.http_health_check.headers` property

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

<a id="canonical-0213010301003333-0022333331023013-1002230323310123-3031233301221021-2321333110331232-0332221201222122-1310231323211321-3210231000320111"></a>

<a id="canonical-2132000222221310-1132002232000112-3032210310013101-1002300321030311-0200211302310122-1031310311122032-3133101202211310-0302120131103030"></a>

#### `job.containers.readiness_check.http_health_check.host_header` property

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

<a id="canonical-3112313121032103-1322200103300313-1121301131002033-2231001330232333-0303110011220311-1113331230323023-2221121001331232-2032200012002310"></a>

<a id="canonical-2002210201102222-0303130113301223-0332111230011112-2321132323130302-0001310130132022-2203311320223221-2322110303300311-1131311232020221"></a>

#### `job.containers.readiness_check.http_health_check.path` property

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

- [port](resources--workload--reference--group-005.md#canonical-3200122021033233-2130333101111311-0022021202122100-1203331110110301-1022122010002222-1212233220021022-1011113020100322-2013330123313030): complete subsection reference.

<a id="canonical-3200122021033233-2130333101111311-0022021202122100-1203331110110301-1022122010002222-1212233220021022-1011113020100322-2013330123313030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `job.containers.readiness_check.http_health_check.port` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [job](resources--workload--reference--group-004.md#canonical-2101133223131122-0231303130011203-2130203200233301-2332212033103313-3210300031001001-1312100313213300-1013120312121230-0320112032220311)
- [job.containers](resources--workload--reference--group-004.md#canonical-2002301323122233-0221102100003021-3013132233211233-3012203030111110-1222020112111222-3222330102102021-2002120323232230-2011302203232120)
- [job.containers.readiness_check](resources--workload--reference--group-005.md#canonical-0222311231110223-3301213322101001-0122100312221013-1333033012100312-0133231033110100-3131310123021323-2322322213211230-3110310102123121)
- [job.containers.readiness_check.http_health_check](resources--workload--reference--group-005.md#canonical-0233132103213321-3313220020312110-2101103210301101-0310230020133310-0222131330330222-3011213321020020-3122230111300132-0313022002201010)
- job.containers.readiness_check.http_health_check.port

<a id="canonical-1032311302021100-1022330120130133-0203223101010322-0013333213012301-1231123101331010-1303021121031001-3113333311111303-3201222102222020"></a>

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

<a id="canonical-2200001011123130-0021331103012303-0130103112221320-1303210131033232-0332033202020300-0031301231131012-1032012330000213-2300330012310131"></a>

### Direct properties for `job.containers.readiness_check.http_health_check.port`

<a id="canonical-3010011210103101-0021230213330223-2331300103331112-0000203120313012-0323301123021003-3002222223330220-2321233003101121-0211122020003322"></a>

#### `job.containers.readiness_check.http_health_check.port.name` property

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

<a id="canonical-2221122111122201-2021012220212130-1313110102120312-0223220111121101-1300221100100212-1210212300201212-2212101223201300-3103132032111201"></a>

<a id="canonical-2311122120202220-1321023200321133-2322102112321100-2032103220001330-0233201100221323-1110312211020212-0321102200201113-1333100000102311"></a>

#### `job.containers.readiness_check.http_health_check.port.num` property

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

<a id="canonical-1331303311323011-1231220102032100-0033122102322121-0323121321110233-3230223220133313-1001001130132012-1012102132021112-3000233323110222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `job.containers.readiness_check.tcp_health_check` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [job](resources--workload--reference--group-004.md#canonical-2101133223131122-0231303130011203-2130203200233301-2332212033103313-3210300031001001-1312100313213300-1013120312121230-0320112032220311)
- [job.containers](resources--workload--reference--group-004.md#canonical-2002301323122233-0221102100003021-3013132233211233-3012203030111110-1222020112111222-3222330102102021-2002120323232230-2011302203232120)
- [job.containers.readiness_check](resources--workload--reference--group-005.md#canonical-0222311231110223-3301213322101001-0122100312221013-1333033012100312-0133231033110100-3131310123021323-2322322213211230-3110310102123121)
- job.containers.readiness_check.tcp_health_check

<a id="canonical-3331002011201311-2100021031110020-1003332232231323-3010201200022321-0303302332012333-2213101203133120-2330033313331010-2121000230103212"></a>

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

<a id="canonical-2030113010210020-2221130201123200-1232002202122230-3331133213132110-1111121121133023-0132223312120303-2030021001223131-3322023303030111"></a>

### Direct properties for `job.containers.readiness_check.tcp_health_check`

- [port](resources--workload--reference--group-005.md#canonical-1313330000022102-0322002022102033-1323002200211000-2122000111002120-0121301120211313-2122102013222132-3330213210000233-3132220012133221): complete subsection reference.

<a id="canonical-1313330000022102-0322002022102033-1323002200211000-2122000111002120-0121301120211313-2122102013222132-3330213210000233-3132220012133221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `job.containers.readiness_check.tcp_health_check.port` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [job](resources--workload--reference--group-004.md#canonical-2101133223131122-0231303130011203-2130203200233301-2332212033103313-3210300031001001-1312100313213300-1013120312121230-0320112032220311)
- [job.containers](resources--workload--reference--group-004.md#canonical-2002301323122233-0221102100003021-3013132233211233-3012203030111110-1222020112111222-3222330102102021-2002120323232230-2011302203232120)
- [job.containers.readiness_check](resources--workload--reference--group-005.md#canonical-0222311231110223-3301213322101001-0122100312221013-1333033012100312-0133231033110100-3131310123021323-2322322213211230-3110310102123121)
- [job.containers.readiness_check.tcp_health_check](resources--workload--reference--group-005.md#canonical-1331303311323011-1231220102032100-0033122102322121-0323121321110233-3230223220133313-1001001130132012-1012102132021112-3000233323110222)
- job.containers.readiness_check.tcp_health_check.port

<a id="canonical-3012211012302320-0200232330230133-2211312000011131-2033211032112202-0301323211101121-0010133311313003-3002121122021022-1203123220000223"></a>

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

<a id="canonical-3133220103013132-1322222200323320-3031113330130330-2123012223003012-2203012123131122-2301320031033331-3031003123330001-2130022121302010"></a>

### Direct properties for `job.containers.readiness_check.tcp_health_check.port`

<a id="canonical-1012110333010302-1030333122022303-1121020001100222-3200110012202023-0320102113131313-2312111313132013-2110222122311032-2302221211322320"></a>

#### `job.containers.readiness_check.tcp_health_check.port.name` property

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

<a id="canonical-2230021030012103-3031232220303203-1213003133023221-2011013001011111-3012212130033003-3112022123332020-3333001202332300-0131220220022202"></a>

<a id="canonical-3032212221311000-0003103223202011-2001333233000300-2131023121102220-3313222221121221-3312212100220200-1302301113310022-2003103132322131"></a>

#### `job.containers.readiness_check.tcp_health_check.port.num` property

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

<a id="canonical-0302031202020321-2011011103202223-2203332002201122-3213130121212012-1321322223212202-2012010100321200-2301013113222301-2013123123001030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `job.deploy_options` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [job](resources--workload--reference--group-004.md#canonical-2101133223131122-0231303130011203-2130203200233301-2332212033103313-3210300031001001-1312100313213300-1013120312121230-0320112032220311)
- job.deploy_options

<a id="canonical-1300233003122113-3130222002303302-3303301233020212-2303212212131023-2113011102301213-3322113112030033-2333222031301222-0231210030020302"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
deploy_options {
  # Configure direct properties listed below.
}
```

<a id="canonical-0310001221123210-0312232003120101-0300203223322231-3003321202313212-2102232112232012-1030001113101333-1133322303112012-0301031133002000"></a>

### Direct properties for `job.deploy_options`

- [all_res](resources--workload--reference--group-005.md#canonical-3131221000303320-1002331010001331-0332121303112322-3121123332120010-1003200023102022-2210023012213201-3232021120032200-2300331010112300): complete subsection reference.

- [default_virtual_sites](resources--workload--reference--group-005.md#canonical-2123101231213202-0333113001323311-1102130211000000-0120320311011312-1121031322123201-1132210012200231-1033021121101112-2312133113213303): complete subsection reference.

- [deploy_ce_sites](resources--workload--reference--group-005.md#canonical-3011331012111131-1200211221323333-2222201313111211-2132322200111103-1210332013232031-2303233120011013-2233031312000221-0310210001111123): complete subsection reference.

- [deploy_ce_virtual_sites](resources--workload--reference--group-005.md#canonical-3123031132321222-1211312320321323-2100330023231222-1011202220031232-2332313222331100-1333220330323313-0212113310011030-0000102112201113): complete subsection reference.

- [deploy_re_sites](resources--workload--reference--group-005.md#canonical-3202012002300303-1220012332222101-0310120133000321-0300120210310013-2101111121220301-0321202013303020-1302001321033011-1133320020002000): complete subsection reference.

- [deploy_re_virtual_sites](resources--workload--reference--group-006.md#canonical-2303323000212010-1331320302002323-2131023001233320-2232113103000220-1031323101300203-3020202312102100-3332323210033031-2112133212011023): complete subsection reference.

<a id="canonical-3131221000303320-1002331010001331-0332121303112322-3121123332120010-1003200023102022-2210023012213201-3232021120032200-2300331010112300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `job.deploy_options.all_res` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [job](resources--workload--reference--group-004.md#canonical-2101133223131122-0231303130011203-2130203200233301-2332212033103313-3210300031001001-1312100313213300-1013120312121230-0320112032220311)
- [job.deploy_options](resources--workload--reference--group-005.md#canonical-0302031202020321-2011011103202223-2203332002201122-3213130121212012-1321322223212202-2012010100321200-2301013113222301-2013123123001030)
- job.deploy_options.all_res

<a id="canonical-3001220121222132-3213003131230210-1132301203312301-1000002322333313-1311232212303113-2322013100331200-3213213013222012-0033203013010320"></a>

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
all_res = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2123101231213202-0333113001323311-1102130211000000-0120320311011312-1121031322123201-1132210012200231-1033021121101112-2312133113213303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `job.deploy_options.default_virtual_sites` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [job](resources--workload--reference--group-004.md#canonical-2101133223131122-0231303130011203-2130203200233301-2332212033103313-3210300031001001-1312100313213300-1013120312121230-0320112032220311)
- [job.deploy_options](resources--workload--reference--group-005.md#canonical-0302031202020321-2011011103202223-2203332002201122-3213130121212012-1321322223212202-2012010100321200-2301013113222301-2013123123001030)
- job.deploy_options.default_virtual_sites

<a id="canonical-0121221313230031-3321123323301111-2312113210103210-1023113210010211-2003033011233212-2211212011231111-3201121120230123-0110103113332030"></a>

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
default_virtual_sites = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3011331012111131-1200211221323333-2222201313111211-2132322200111103-1210332013232031-2303233120011013-2233031312000221-0310210001111123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `job.deploy_options.deploy_ce_sites` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [job](resources--workload--reference--group-004.md#canonical-2101133223131122-0231303130011203-2130203200233301-2332212033103313-3210300031001001-1312100313213300-1013120312121230-0320112032220311)
- [job.deploy_options](resources--workload--reference--group-005.md#canonical-0302031202020321-2011011103202223-2203332002201122-3213130121212012-1321322223212202-2012010100321200-2301013113222301-2013123123001030)
- job.deploy_options.deploy_ce_sites

<a id="canonical-0130221002120312-2120013012100022-1311011112211003-2033300232110013-1103331323202330-3110223032321120-1101321202203022-2330121321121213"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
deploy_ce_sites {
  # Configure direct properties listed below.
}
```

<a id="canonical-3030202312013020-2331300020103200-3222012331202123-0211312322223101-3002201310201210-2223100333330112-0313030330111012-0021233221022203"></a>

### Direct properties for `job.deploy_options.deploy_ce_sites`

- [site](resources--workload--reference--group-005.md#canonical-1230231020332032-3102101212021311-1032313121321003-2210233300123223-3310211000231201-1002201202130100-1032212021100123-1322303021300203): complete subsection reference.

<a id="canonical-1230231020332032-3102101212021311-1032313121321003-2210233300123223-3310211000231201-1002201202130100-1032212021100123-1322303021300203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `job.deploy_options.deploy_ce_sites.site` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [job](resources--workload--reference--group-004.md#canonical-2101133223131122-0231303130011203-2130203200233301-2332212033103313-3210300031001001-1312100313213300-1013120312121230-0320112032220311)
- [job.deploy_options](resources--workload--reference--group-005.md#canonical-0302031202020321-2011011103202223-2203332002201122-3213130121212012-1321322223212202-2012010100321200-2301013113222301-2013123123001030)
- [job.deploy_options.deploy_ce_sites](resources--workload--reference--group-005.md#canonical-3011331012111131-1200211221323333-2222201313111211-2132322200111103-1210332013232031-2303233120011013-2233031312000221-0310210001111123)
- job.deploy_options.deploy_ce_sites.site

<a id="canonical-0312312101122000-0020310003122231-2013100033132310-0003220322310301-0220012202103112-1203030021000131-0223323121202132-1110000300221303"></a>

Type: `"object"`. list nested block, Optional.

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

Terraform syntax:

```terraform
site {
  # Configure direct properties listed below.
}
```

<a id="canonical-2102133001133310-2303302121232022-0021132323221231-1300020010311222-3003220310323102-2303231103100212-0202331010000323-2320103333011200"></a>

### Direct properties for `job.deploy_options.deploy_ce_sites.site`

<a id="canonical-3001230221103003-1131113103130113-2300031313011222-3200012121112230-1302132231311101-3130232221122100-3222111232332112-3201122222000030"></a>

#### `job.deploy_options.deploy_ce_sites.site.name` property

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

<a id="canonical-3011032112012302-0230221012313010-2213100003203220-2310122302323332-1212212220110231-1223212013301030-3013300310332230-2133332120210301"></a>

<a id="canonical-0310112230202002-0111022230002230-3222310110302032-2311101013221123-1120210331213213-2202302203110213-1002202011121120-2033200231333122"></a>

#### `job.deploy_options.deploy_ce_sites.site.namespace` property

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

<a id="canonical-3113310102231213-3203013123211331-3032300203223012-2221020321323013-2321303320310320-2022203033313233-2133330010223231-2013233222200032"></a>

<a id="canonical-3202112110313102-2012113010131013-0112010102231033-1030233223111311-0111133323032311-2101313130222230-1111102203321132-0001220202222213"></a>

#### `job.deploy_options.deploy_ce_sites.site.tenant` property

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

<a id="canonical-3123031132321222-1211312320321323-2100330023231222-1011202220031232-2332313222331100-1333220330323313-0212113310011030-0000102112201113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `job.deploy_options.deploy_ce_virtual_sites` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [job](resources--workload--reference--group-004.md#canonical-2101133223131122-0231303130011203-2130203200233301-2332212033103313-3210300031001001-1312100313213300-1013120312121230-0320112032220311)
- [job.deploy_options](resources--workload--reference--group-005.md#canonical-0302031202020321-2011011103202223-2203332002201122-3213130121212012-1321322223212202-2012010100321200-2301013113222301-2013123123001030)
- job.deploy_options.deploy_ce_virtual_sites

<a id="canonical-3220112202232302-0211321011332220-2310203023010222-2300311222201313-3110021032003212-2023001123122032-2111032002331110-1333300331130202"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
deploy_ce_virtual_sites {
  # Configure direct properties listed below.
}
```

<a id="canonical-2133022102122202-0001310311123112-2003032312321010-1010230221003223-1201033211320121-1002131021231333-3320130133101032-2123222212223331"></a>

### Direct properties for `job.deploy_options.deploy_ce_virtual_sites`

- [virtual_site](resources--workload--reference--group-005.md#canonical-2123110231002301-0321331133020211-1102300222122111-3213213233323302-3132321020001221-2203011331020000-0200113020321133-1020330210002001): complete subsection reference.

<a id="canonical-2123110231002301-0321331133020211-1102300222122111-3213213233323302-3132321020001221-2203011331020000-0200113020321133-1020330210002001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `job.deploy_options.deploy_ce_virtual_sites.virtual_site` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [job](resources--workload--reference--group-004.md#canonical-2101133223131122-0231303130011203-2130203200233301-2332212033103313-3210300031001001-1312100313213300-1013120312121230-0320112032220311)
- [job.deploy_options](resources--workload--reference--group-005.md#canonical-0302031202020321-2011011103202223-2203332002201122-3213130121212012-1321322223212202-2012010100321200-2301013113222301-2013123123001030)
- [job.deploy_options.deploy_ce_virtual_sites](resources--workload--reference--group-005.md#canonical-3123031132321222-1211312320321323-2100330023231222-1011202220031232-2332313222331100-1333220330323313-0212113310011030-0000102112201113)
- job.deploy_options.deploy_ce_virtual_sites.virtual_site

<a id="canonical-0220210230000110-0033131113020322-0210010022203220-0311000211033232-2032311120323220-3322003113201322-1222122133110303-0221202323131322"></a>

Type: `"object"`. list nested block, Optional.

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

Terraform syntax:

```terraform
virtual_site {
  # Configure direct properties listed below.
}
```

<a id="canonical-0311231111313231-2332033303002000-3310133232220301-3221133112322223-3100132102211310-2030212201231122-1133021302002020-1033320112112021"></a>

### Direct properties for `job.deploy_options.deploy_ce_virtual_sites.virtual_site`

<a id="canonical-1013302323312320-3012330213330123-3220230231023331-3310110222023222-1002222310300032-2332112032202331-2302021201120131-1023330213111200"></a>

#### `job.deploy_options.deploy_ce_virtual_sites.virtual_site.name` property

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

<a id="canonical-3322113202131002-2331121322210100-3230132002302123-3323002322333212-2323022112021122-2202330020312220-0001032230112010-3033032230110110"></a>

<a id="canonical-0330113120012100-0333023010120023-3330031000320221-0331221233212102-3313222222300031-0000201020021100-0023221220211020-0100200000013313"></a>

#### `job.deploy_options.deploy_ce_virtual_sites.virtual_site.namespace` property

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

<a id="canonical-3020130112320032-3320022102002220-3331100123123130-0022031333032131-1111333133032110-0330313201130300-2100220322203330-1120002132021022"></a>

<a id="canonical-1031311233302310-2202011213132121-0300212102220110-0332130103210233-2200211013331333-0321221112203101-1303332312332112-1333322122230102"></a>

#### `job.deploy_options.deploy_ce_virtual_sites.virtual_site.tenant` property

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

<a id="canonical-3202012002300303-1220012332222101-0310120133000321-0300120210310013-2101111121220301-0321202013303020-1302001321033011-1133320020002000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `job.deploy_options.deploy_re_sites` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [job](resources--workload--reference--group-004.md#canonical-2101133223131122-0231303130011203-2130203200233301-2332212033103313-3210300031001001-1312100313213300-1013120312121230-0320112032220311)
- [job.deploy_options](resources--workload--reference--group-005.md#canonical-0302031202020321-2011011103202223-2203332002201122-3213130121212012-1321322223212202-2012010100321200-2301013113222301-2013123123001030)
- job.deploy_options.deploy_re_sites

<a id="canonical-0220020110011203-1332010013330210-2032002112221130-0030012022101301-1011211101200120-1013303212013010-0012023103333031-0121112133221202"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
deploy_re_sites {
  # Configure direct properties listed below.
}
```

<a id="canonical-2210120100120200-0022132123312011-0120012233333022-3101121023332230-0320313101003101-2122203132233100-3222033001331321-0302232333332232"></a>

### Direct properties for `job.deploy_options.deploy_re_sites`

- [site](resources--workload--reference--group-006.md#canonical-1013311331031231-3112002103021332-3330003313111330-3023232231020030-2033001221022131-3102223310202001-0323311203112313-2022123331132012): complete subsection reference.
