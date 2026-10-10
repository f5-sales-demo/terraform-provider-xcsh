---
page_title: "xcsh_workload reference"
subcategory: "Container"
description: "Complete grouped canonical reference for xcsh_workload reference."
---

# xcsh_workload reference

<a id="canonical-1200333303302210-1101011312333201-2221003330332123-3010232020133102-1232001233112201-1313103133310301-1220211130202110-2203200203302013"></a>

## `simple_service.container.liveness_check.http_health_check.host_header` property

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

<a id="canonical-3202222300330111-2233031313323300-1030203300110011-3123130123121323-1121312100122110-2021323222231003-2233231133300011-1322111213100231"></a>

<a id="canonical-3111113121132333-1021311033133323-1110001101122130-0230312312303010-1310302212101021-0331001131212213-3310003011021233-0003332100123330"></a>

## `simple_service.container.liveness_check.http_health_check.path` property

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

- [port](resources--workload--reference--group-018.md#canonical-0332012012023132-1213232202130103-1201221221101131-2300330101111132-1002201203020001-3203321221133102-3123333133031123-2120302130201231): complete subsection reference.

<a id="canonical-0332012012023132-1213232202130103-1201221221101131-2300330101111132-1002201203020001-3203321221133102-3123333133031123-2120302130201231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `simple_service.container.liveness_check.http_health_check.port` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [simple_service](resources--workload--reference--group-017.md#canonical-2103120123000022-0231212311020021-3011030332210231-1230230313313103-1213231032231100-0321113222111303-0220111021013230-2333203212100121)
- [simple_service.container](resources--workload--reference--group-017.md#canonical-1011303000022103-1122231110321130-3202121331100121-1301320113212120-3220331111111112-1120022220012313-3313013203321210-0332330312030213)
- [simple_service.container.liveness_check](resources--workload--reference--group-017.md#canonical-1123003201223020-3031001003003030-2323130113230021-2110112120231030-0331101121132200-2010210232213121-3023312203232202-1112033223321112)
- [simple_service.container.liveness_check.http_health_check](resources--workload--reference--group-017.md#canonical-0123112211202011-1011003322102323-3231201102302033-2302101301120031-3120122012122032-3212311301223311-2220133131010031-0003200300112301)
- simple_service.container.liveness_check.http_health_check.port

<a id="canonical-0331301120110322-0223002023311331-0022013120103100-2033123121220211-2212100220332220-0311231223321321-2320121101111123-0003031230311132"></a>

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

<a id="canonical-2021033011303120-3110201001200211-3223322031230321-3321101113213031-0332231233332002-3300230311011001-0211132123303012-2232311132201033"></a>

### Direct properties for `simple_service.container.liveness_check.http_health_check.port`

<a id="canonical-0022330232302330-1231202320101031-1312322212222031-2230311111211331-0022102221210300-1223023233102112-2231000113211233-0110302123133132"></a>

#### `simple_service.container.liveness_check.http_health_check.port.name` property

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

<a id="canonical-3101220022330013-2002020302213230-2312112000130112-2330332102003102-3230122201212122-3011200012002033-3222122120112132-0013211133010220"></a>

<a id="canonical-3312221121223112-0202220202332303-3130003112132333-1310102102110023-1112321021113013-3100220103211230-3011333000021223-0231322223322312"></a>

#### `simple_service.container.liveness_check.http_health_check.port.num` property

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

<a id="canonical-1321011121310301-2221231110312101-2203132301320030-3030201110120231-3002002032122030-1130100002311331-3130212221120312-1230122223333111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `simple_service.container.liveness_check.tcp_health_check` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [simple_service](resources--workload--reference--group-017.md#canonical-2103120123000022-0231212311020021-3011030332210231-1230230313313103-1213231032231100-0321113222111303-0220111021013230-2333203212100121)
- [simple_service.container](resources--workload--reference--group-017.md#canonical-1011303000022103-1122231110321130-3202121331100121-1301320113212120-3220331111111112-1120022220012313-3313013203321210-0332330312030213)
- [simple_service.container.liveness_check](resources--workload--reference--group-017.md#canonical-1123003201223020-3031001003003030-2323130113230021-2110112120231030-0331101121132200-2010210232213121-3023312203232202-1112033223321112)
- simple_service.container.liveness_check.tcp_health_check

<a id="canonical-1313220130101003-0202110231310202-0120011102033103-1221331221210011-1320000232303112-3120100002113112-3320002201131230-2102120213301000"></a>

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

<a id="canonical-3223010033200012-1201001000122031-3022030210320022-3312221313303111-1100213020132231-2302301322000033-3312213321131012-2301322203300121"></a>

### Direct properties for `simple_service.container.liveness_check.tcp_health_check`

- [port](resources--workload--reference--group-018.md#canonical-0032202331212122-1213122130100333-0123002331311201-0021322022122110-0103132013120120-0120333032213021-0231331110003102-3111031301213030): complete subsection reference.

<a id="canonical-0032202331212122-1213122130100333-0123002331311201-0021322022122110-0103132013120120-0120333032213021-0231331110003102-3111031301213030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `simple_service.container.liveness_check.tcp_health_check.port` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [simple_service](resources--workload--reference--group-017.md#canonical-2103120123000022-0231212311020021-3011030332210231-1230230313313103-1213231032231100-0321113222111303-0220111021013230-2333203212100121)
- [simple_service.container](resources--workload--reference--group-017.md#canonical-1011303000022103-1122231110321130-3202121331100121-1301320113212120-3220331111111112-1120022220012313-3313013203321210-0332330312030213)
- [simple_service.container.liveness_check](resources--workload--reference--group-017.md#canonical-1123003201223020-3031001003003030-2323130113230021-2110112120231030-0331101121132200-2010210232213121-3023312203232202-1112033223321112)
- [simple_service.container.liveness_check.tcp_health_check](resources--workload--reference--group-018.md#canonical-1321011121310301-2221231110312101-2203132301320030-3030201110120231-3002002032122030-1130100002311331-3130212221120312-1230122223333111)
- simple_service.container.liveness_check.tcp_health_check.port

<a id="canonical-1121112132311022-2301302113132321-3120111103101022-1022111331202020-2130201013220133-1003131321121211-3010303011003101-0131321120233222"></a>

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

<a id="canonical-3203020230221100-1110232000212000-3032023003011233-3031133131112023-3030113330100320-0023102131100312-1200002022020222-1322022233202312"></a>

### Direct properties for `simple_service.container.liveness_check.tcp_health_check.port`

<a id="canonical-2233010232302311-1013122110312131-3133021330222021-2123223213231002-3323011102132230-0102333110211103-0223231210020222-2023023103302011"></a>

#### `simple_service.container.liveness_check.tcp_health_check.port.name` property

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

<a id="canonical-3013102001023303-3332031023033213-0213203133103320-1130233300221333-1021333102203032-1302011211200110-2321222221331013-0103333221110021"></a>

<a id="canonical-3021222330331003-1030111220223320-2231333110011002-3220121302000202-3201101102131322-1103133113101120-1113022031322231-2013013121312121"></a>

#### `simple_service.container.liveness_check.tcp_health_check.port.num` property

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

<a id="canonical-3212121030101233-3131221033310023-1233033330130102-3013330332203213-3033003011032322-1323201011111220-2202323123332313-3220300202001323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `simple_service.container.readiness_check` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [simple_service](resources--workload--reference--group-017.md#canonical-2103120123000022-0231212311020021-3011030332210231-1230230313313103-1213231032231100-0321113222111303-0220111021013230-2333203212100121)
- [simple_service.container](resources--workload--reference--group-017.md#canonical-1011303000022103-1122231110321130-3202121331100121-1301320113212120-3220331111111112-1120022220012313-3313013203321210-0332330312030213)
- simple_service.container.readiness_check

<a id="canonical-1312221123021031-3300121033033321-2130230221003102-0330131012201000-1020220321300000-0003030102322323-1300231233133020-2030223313132331"></a>

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

<a id="canonical-2202222122101112-3110113113210121-3331020111103310-3223130101322223-1333011221120033-3110132103202331-3333011131211100-3200210121210221"></a>

### Direct properties for `simple_service.container.readiness_check`

- [exec_health_check](resources--workload--reference--group-018.md#canonical-2232303312023003-3302232112300100-0133310100100211-3110312312010310-1032110113031121-1212102123203212-0231111010310002-2113113310202331): complete subsection reference.

<a id="canonical-1301011330112211-3010100133113033-1111100020233010-3311233123332000-2001100012312132-1122322221300020-0102030211113011-2320303223132132"></a>

<a id="canonical-3313213030033211-0311213332232332-0112303313100322-3220332120300013-2203030231231211-0320311323213001-1202022220332033-0013211331232200"></a>

#### `simple_service.container.readiness_check.healthy_threshold` property

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

- [http_health_check](resources--workload--reference--group-018.md#canonical-0013303313233031-1110230113303003-0321332113313013-2131211101210102-2303120323310310-3011231232223303-3031011103321102-1111122103321322): complete subsection reference.

<a id="canonical-3013011102213032-1111323122120123-2303231221302000-1032303101201233-1212003302332322-1310130203122223-0302333131123011-1230021330201020"></a>

<a id="canonical-1200303003012322-3221300132131303-0110211001022102-0011211113330231-3333333200001312-1303000212222010-1103031010132223-1002022300021102"></a>

#### `simple_service.container.readiness_check.initial_delay` property

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

<a id="canonical-3013313331213202-3003212003301200-2222103320012222-1203303122322132-3321311100210221-0221202030231103-1011230132030332-2333131020312320"></a>

<a id="canonical-2330112032203300-1130323333003223-0130203302100331-1302201203222213-1322210233131001-0002100012311011-1221300031332201-3030320210301031"></a>

#### `simple_service.container.readiness_check.interval` property

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

- [tcp_health_check](resources--workload--reference--group-018.md#canonical-2120132203003331-2020222300211130-1230021011312132-1001012000221222-1131332001122301-3000222103312120-1010001332230132-0210313102002011): complete subsection reference.

<a id="canonical-3101320013012203-2113023323323131-0001201201101031-3332330323013021-2101011121320302-3033203030333100-1033112012312010-2031303321022013"></a>

<a id="canonical-2122202331211013-1023311133033113-3012330020312120-0012031131323301-3013200212111323-1310200210323200-2221033023020203-2020101120000012"></a>

#### `simple_service.container.readiness_check.timeout` property

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

<a id="canonical-1221311032010002-2221310312112211-3233233120330112-1310232020230021-0223212300332022-1211331332323222-0201130033111131-3033202022010301"></a>

<a id="canonical-0302312100212213-0031130323232010-0101023121202003-0022123023130220-2113023130020212-3000310111331100-1330231312101013-0131101303231320"></a>

#### `simple_service.container.readiness_check.unhealthy_threshold` property

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

<a id="canonical-2232303312023003-3302232112300100-0133310100100211-3110312312010310-1032110113031121-1212102123203212-0231111010310002-2113113310202331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `simple_service.container.readiness_check.exec_health_check` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [simple_service](resources--workload--reference--group-017.md#canonical-2103120123000022-0231212311020021-3011030332210231-1230230313313103-1213231032231100-0321113222111303-0220111021013230-2333203212100121)
- [simple_service.container](resources--workload--reference--group-017.md#canonical-1011303000022103-1122231110321130-3202121331100121-1301320113212120-3220331111111112-1120022220012313-3313013203321210-0332330312030213)
- [simple_service.container.readiness_check](resources--workload--reference--group-018.md#canonical-3212121030101233-3131221033310023-1233033330130102-3013330332203213-3033003011032322-1323201011111220-2202323123332313-3220300202001323)
- simple_service.container.readiness_check.exec_health_check

<a id="canonical-2111111330333303-1233103131022332-1121010023303033-1001010112123202-0222103033232113-0133000113212202-0023102003123203-2010222111201330"></a>

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

<a id="canonical-2233021120213121-1110213230322123-3013110212101031-3121230230121030-0030003231312203-1010002210123333-3120312212003001-3133320023332110"></a>

### Direct properties for `simple_service.container.readiness_check.exec_health_check`

<a id="canonical-2031221000000332-3113203013023220-0031312313301203-1300100013132111-1120231322012113-1231020132220122-2303031022313110-1100221332001230"></a>

#### `simple_service.container.readiness_check.exec_health_check.command` property

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

<a id="canonical-0013303313233031-1110230113303003-0321332113313013-2131211101210102-2303120323310310-3011231232223303-3031011103321102-1111122103321322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `simple_service.container.readiness_check.http_health_check` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [simple_service](resources--workload--reference--group-017.md#canonical-2103120123000022-0231212311020021-3011030332210231-1230230313313103-1213231032231100-0321113222111303-0220111021013230-2333203212100121)
- [simple_service.container](resources--workload--reference--group-017.md#canonical-1011303000022103-1122231110321130-3202121331100121-1301320113212120-3220331111111112-1120022220012313-3313013203321210-0332330312030213)
- [simple_service.container.readiness_check](resources--workload--reference--group-018.md#canonical-3212121030101233-3131221033310023-1233033330130102-3013330332203213-3033003011032322-1323201011111220-2202323123332313-3220300202001323)
- simple_service.container.readiness_check.http_health_check

<a id="canonical-3100220130333110-2110122001101133-3010203133332033-1323013203320123-1330222130203300-0123230111230311-1232210010333131-1020233323101023"></a>

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

<a id="canonical-3110203203020011-3101211111330110-2102330003112230-1230011120213132-1030213230313333-1310301211223022-0021012130111332-2330302333121030"></a>

### Direct properties for `simple_service.container.readiness_check.http_health_check`

<a id="canonical-1112111131010302-1111330020230020-2201113233221332-0131310302212231-2132001010023030-1233000202212220-2003333330212121-3210000210300200"></a>

#### `simple_service.container.readiness_check.http_health_check.headers` property

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

<a id="canonical-0220010322303023-3001213011212022-1222122122010300-2232301311113201-0323321111231101-0122332001323230-2333131313013020-0133020333000030"></a>

<a id="canonical-1130011003020030-3022020002001323-2101102213112333-1020321300300122-3130213200233022-3213011330120032-1113301330222023-1203111103021213"></a>

#### `simple_service.container.readiness_check.http_health_check.host_header` property

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

<a id="canonical-1212032033330102-3023033303000331-2010132133201222-3033322231300202-0111330200213210-0210011011120312-3013212020103112-2300303131033020"></a>

<a id="canonical-0332213133120312-3220022103133222-3300222130132222-2020222232312030-2212332232220222-1313022303322132-2033111230312330-3130010013330210"></a>

#### `simple_service.container.readiness_check.http_health_check.path` property

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

- [port](resources--workload--reference--group-018.md#canonical-2100330323023130-0012320000123300-1222011220120312-0102120120210020-1210331021123232-3230110113032102-3310002103203322-3121033003300200): complete subsection reference.

<a id="canonical-2100330323023130-0012320000123300-1222011220120312-0102120120210020-1210331021123232-3230110113032102-3310002103203322-3121033003300200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `simple_service.container.readiness_check.http_health_check.port` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [simple_service](resources--workload--reference--group-017.md#canonical-2103120123000022-0231212311020021-3011030332210231-1230230313313103-1213231032231100-0321113222111303-0220111021013230-2333203212100121)
- [simple_service.container](resources--workload--reference--group-017.md#canonical-1011303000022103-1122231110321130-3202121331100121-1301320113212120-3220331111111112-1120022220012313-3313013203321210-0332330312030213)
- [simple_service.container.readiness_check](resources--workload--reference--group-018.md#canonical-3212121030101233-3131221033310023-1233033330130102-3013330332203213-3033003011032322-1323201011111220-2202323123332313-3220300202001323)
- [simple_service.container.readiness_check.http_health_check](resources--workload--reference--group-018.md#canonical-0013303313233031-1110230113303003-0321332113313013-2131211101210102-2303120323310310-3011231232223303-3031011103321102-1111122103321322)
- simple_service.container.readiness_check.http_health_check.port

<a id="canonical-2213021231023100-1232023013022121-1221103122023122-2123022300122310-1000132301101021-1102030110113122-2332103303021200-0330013332312022"></a>

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

<a id="canonical-1112222130002313-2320112112320233-2212232101013302-2121310012011121-2110202110100113-0131030110222123-3330332120203131-3033200000223033"></a>

### Direct properties for `simple_service.container.readiness_check.http_health_check.port`

<a id="canonical-1131002211102322-2020132311122320-0030021121330223-1332311230322232-1231031330132221-3021012323133011-1312333322122320-0023202122321012"></a>

#### `simple_service.container.readiness_check.http_health_check.port.name` property

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

<a id="canonical-2230013213132122-3330232212211133-1132000322030203-0321320322103200-3333133000312022-3201022201200121-1313023030000210-2332201322102202"></a>

<a id="canonical-1313221320303111-0011121302221120-1133211102013120-0000332123002112-1331303010111321-1131321301310122-1332031310011000-2010223230113011"></a>

#### `simple_service.container.readiness_check.http_health_check.port.num` property

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

<a id="canonical-2120132203003331-2020222300211130-1230021011312132-1001012000221222-1131332001122301-3000222103312120-1010001332230132-0210313102002011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `simple_service.container.readiness_check.tcp_health_check` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [simple_service](resources--workload--reference--group-017.md#canonical-2103120123000022-0231212311020021-3011030332210231-1230230313313103-1213231032231100-0321113222111303-0220111021013230-2333203212100121)
- [simple_service.container](resources--workload--reference--group-017.md#canonical-1011303000022103-1122231110321130-3202121331100121-1301320113212120-3220331111111112-1120022220012313-3313013203321210-0332330312030213)
- [simple_service.container.readiness_check](resources--workload--reference--group-018.md#canonical-3212121030101233-3131221033310023-1233033330130102-3013330332203213-3033003011032322-1323201011111220-2202323123332313-3220300202001323)
- simple_service.container.readiness_check.tcp_health_check

<a id="canonical-1113003001321000-1133203331120113-1213011131220030-2020202002203111-3133022323311202-3112202030211021-3300312030020020-3023120323321302"></a>

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

<a id="canonical-3232330222222012-1320313230103110-1013220220311232-2210333100301331-1330020202230010-0030130103112130-1321221002333202-2301003331211111"></a>

### Direct properties for `simple_service.container.readiness_check.tcp_health_check`

- [port](resources--workload--reference--group-018.md#canonical-2223220333130313-3102120110210000-3321133301022331-0213020030220300-3000100212221322-0231113001300020-1001031301313020-1113301333232231): complete subsection reference.

<a id="canonical-2223220333130313-3102120110210000-3321133301022331-0213020030220300-3000100212221322-0231113001300020-1001031301313020-1113301333232231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `simple_service.container.readiness_check.tcp_health_check.port` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [simple_service](resources--workload--reference--group-017.md#canonical-2103120123000022-0231212311020021-3011030332210231-1230230313313103-1213231032231100-0321113222111303-0220111021013230-2333203212100121)
- [simple_service.container](resources--workload--reference--group-017.md#canonical-1011303000022103-1122231110321130-3202121331100121-1301320113212120-3220331111111112-1120022220012313-3313013203321210-0332330312030213)
- [simple_service.container.readiness_check](resources--workload--reference--group-018.md#canonical-3212121030101233-3131221033310023-1233033330130102-3013330332203213-3033003011032322-1323201011111220-2202323123332313-3220300202001323)
- [simple_service.container.readiness_check.tcp_health_check](resources--workload--reference--group-018.md#canonical-2120132203003331-2020222300211130-1230021011312132-1001012000221222-1131332001122301-3000222103312120-1010001332230132-0210313102002011)
- simple_service.container.readiness_check.tcp_health_check.port

<a id="canonical-2032320211203132-2012020332331213-0310203122312311-1121103110322223-1001312222322112-0230112112312213-1212012300132321-1113110002013113"></a>

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

<a id="canonical-0111102103110230-3001222320311123-0211201131131301-3213201320102101-2223210020133010-0103210233322332-0300132103313023-3230010020102111"></a>

### Direct properties for `simple_service.container.readiness_check.tcp_health_check.port`

<a id="canonical-1311210310003221-2032323310021102-1010232022121010-3001311313330312-3121023331230112-1321031202110123-2231210200022331-1233003313021131"></a>

#### `simple_service.container.readiness_check.tcp_health_check.port.name` property

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

<a id="canonical-0103322122023311-3222010022113230-2201332331011213-2232110312223020-1012320203222033-0231303301331323-0332330321120101-2102123122110320"></a>

<a id="canonical-3130103231113222-0030220130203012-2121021221001333-3101030002120312-1223121303122133-3132321301322012-2201131001122220-1201311230002303"></a>

#### `simple_service.container.readiness_check.tcp_health_check.port.num` property

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

<a id="canonical-1110023311313011-0202000212230320-0210021130232023-0323322102330130-3223121303212321-2033302130023021-3120001020331002-1031122330013020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `simple_service.disabled` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [simple_service](resources--workload--reference--group-017.md#canonical-2103120123000022-0231212311020021-3011030332210231-1230230313313103-1213231032231100-0321113222111303-0220111021013230-2333203212100121)
- simple_service.disabled

<a id="canonical-1331010000111000-1030202022131033-2012020223002100-0003132030122233-3022132001330130-3220121200323313-3101011022232320-0213200112210222"></a>

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
disabled = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0003300212123321-0030022111222023-1030022232221021-1132031212101132-3230032012113311-1303232022202331-3111103333310112-3310021021132230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `simple_service.do_not_advertise` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [simple_service](resources--workload--reference--group-017.md#canonical-2103120123000022-0231212311020021-3011030332210231-1230230313313103-1213231032231100-0321113222111303-0220111021013230-2333203212100121)
- simple_service.do_not_advertise

<a id="canonical-3201133201002001-2102111030231333-3131002012103102-0312101212121323-0001330223220013-0030332222032310-1212110332012133-3210100130302133"></a>

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

<a id="canonical-0022120332023020-2312100113020133-0123031032111130-3302333320122221-1232003333212033-1103032302323333-2231001310003211-0310320121030030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `simple_service.enabled` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [simple_service](resources--workload--reference--group-017.md#canonical-2103120123000022-0231212311020021-3011030332210231-1230230313313103-1213231032231100-0321113222111303-0220111021013230-2333203212100121)
- simple_service.enabled

<a id="canonical-1013123123330323-1013010313012311-2000301120012101-3202132002222132-1320332112320030-2323023103313302-3012202221213302-3011202333223233"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
enabled {
  # Configure direct properties listed below.
}
```

<a id="canonical-1331323031320021-0110122023013313-0312123012032100-0313020100000011-2322120102323032-3331231123223202-3013202122213233-0002012220032203"></a>

### Direct properties for `simple_service.enabled`

<a id="canonical-2122211312113320-3320233131201231-0213002312332013-2121231321302322-3122030133100323-1102333211001202-0223330211322232-3102110110110121"></a>

#### `simple_service.enabled.name` property

Type: `"string"`. Optional.

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

- [persistent_volume](resources--workload--reference--group-018.md#canonical-3113200332101222-1011331333203002-1210020131130330-0320221103232321-3010131320311330-1112331013232233-2101133032201232-1113223131002320): complete subsection reference.

<a id="canonical-3113200332101222-1011331333203002-1210020131130330-0320221103232321-3010131320311330-1112331013232233-2101133032201232-1113223131002320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `simple_service.enabled.persistent_volume` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [simple_service](resources--workload--reference--group-017.md#canonical-2103120123000022-0231212311020021-3011030332210231-1230230313313103-1213231032231100-0321113222111303-0220111021013230-2333203212100121)
- [simple_service.enabled](resources--workload--reference--group-018.md#canonical-0022120332023020-2312100113020133-0123031032111130-3302333320122221-1232003333212033-1103032302323333-2231001310003211-0310320121030030)
- simple_service.enabled.persistent_volume

<a id="canonical-1121031211201101-0230221011311012-1320210121011002-3010012020313132-3211031101333331-3220101211012313-3113211201110013-3003023001003131"></a>

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

<a id="canonical-2033011223320002-3031303111120031-3132023120233131-0022321330323011-0333203320333332-0130022333110012-0223312333021112-2023031300200303"></a>

### Direct properties for `simple_service.enabled.persistent_volume`

- [mount](resources--workload--reference--group-018.md#canonical-2111120313310323-2221023331321012-2030221311033320-3003121230321330-3323312203123022-2322013321110122-0100031032103110-0022321200201303): complete subsection reference.

- [storage](resources--workload--reference--group-018.md#canonical-0212031230210030-0110133221302101-1211130030200033-0222110222222113-1311020123323113-1302101310023300-1101100122133002-0221000223123310): complete subsection reference.

<a id="canonical-2111120313310323-2221023331321012-2030221311033320-3003121230321330-3323312203123022-2322013321110122-0100031032103110-0022321200201303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `simple_service.enabled.persistent_volume.mount` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [simple_service](resources--workload--reference--group-017.md#canonical-2103120123000022-0231212311020021-3011030332210231-1230230313313103-1213231032231100-0321113222111303-0220111021013230-2333203212100121)
- [simple_service.enabled](resources--workload--reference--group-018.md#canonical-0022120332023020-2312100113020133-0123031032111130-3302333320122221-1232003333212033-1103032302323333-2231001310003211-0310320121030030)
- [simple_service.enabled.persistent_volume](resources--workload--reference--group-018.md#canonical-3113200332101222-1011331333203002-1210020131130330-0320221103232321-3010131320311330-1112331013232233-2101133032201232-1113223131002320)
- simple_service.enabled.persistent_volume.mount

<a id="canonical-0220032131123200-0013033203120011-3331122310210130-2221212231023312-3200220300103100-0323312333100232-2330001202230120-3333102212130211"></a>

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

<a id="canonical-3323222300202021-3212100311320323-1211130303132100-1330133113300202-2102121123303030-2110021010123122-1131311121130020-0100201312003330"></a>

### Direct properties for `simple_service.enabled.persistent_volume.mount`

<a id="canonical-0132010212030021-3223033023223330-0312131223203131-2122022231032021-0001123311012013-1301011231131303-3020203101330133-3012122111132313"></a>

#### `simple_service.enabled.persistent_volume.mount.mode` property

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

<a id="canonical-3010210323030012-1320021010100201-0221312310322320-2321001122012233-2321213122220211-2113223230010230-0112103132221132-3330222210331232"></a>

<a id="canonical-2232213212132313-1330221032021011-3002323300311312-0202112323201012-2303110312320123-3111211332120231-2332222132301122-0313020013320000"></a>

#### `simple_service.enabled.persistent_volume.mount.mount_path` property

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

<a id="canonical-0211301020323111-3320220232130113-0110333333231123-2023100310031221-1013330302012201-3213332221331311-3101332203301330-1300120210122111"></a>

<a id="canonical-0332300312133220-1003301300012000-3303310210312111-0302221301111002-1313210232321103-2001031101220133-0131000123133031-1010202111132100"></a>

#### `simple_service.enabled.persistent_volume.mount.sub_path` property

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

<a id="canonical-0212031230210030-0110133221302101-1211130030200033-0222110222222113-1311020123323113-1302101310023300-1101100122133002-0221000223123310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `simple_service.enabled.persistent_volume.storage` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [simple_service](resources--workload--reference--group-017.md#canonical-2103120123000022-0231212311020021-3011030332210231-1230230313313103-1213231032231100-0321113222111303-0220111021013230-2333203212100121)
- [simple_service.enabled](resources--workload--reference--group-018.md#canonical-0022120332023020-2312100113020133-0123031032111130-3302333320122221-1232003333212033-1103032302323333-2231001310003211-0310320121030030)
- [simple_service.enabled.persistent_volume](resources--workload--reference--group-018.md#canonical-3113200332101222-1011331333203002-1210020131130330-0320221103232321-3010131320311330-1112331013232233-2101133032201232-1113223131002320)
- simple_service.enabled.persistent_volume.storage

<a id="canonical-1000100132031322-3002211033330331-1123302221120222-1312132220011223-3200313220103133-1203203103230310-2111231230102213-0232313112303032"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
storage {
  # Configure direct properties listed below.
}
```

<a id="canonical-3002113023121033-1123313000023001-0330032130132130-3330330231002202-0012221231313212-3310333230301201-1013222322013200-3133220020333222"></a>

### Direct properties for `simple_service.enabled.persistent_volume.storage`

<a id="canonical-1212230222002111-3312131323303032-1331132123020330-1122321213323003-2021231213033221-0123222200311110-2031233002110000-0222022122301301"></a>

#### `simple_service.enabled.persistent_volume.storage.access_mode` property

Type: `"string"`. Optional.

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

<a id="canonical-3323031332311133-0033013133220223-3303222202031320-0012332200011011-0332221200310020-2313023030201222-2012300030021132-3303212122313003"></a>

<a id="canonical-1201123023210313-2120313303123303-0020312333233123-0013021110302033-0333213130212232-2313033132000330-0300213121230201-3201331011310000"></a>

#### `simple_service.enabled.persistent_volume.storage.class_name` property

Type: `"string"`. Optional.

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

- [default](resources--workload--reference--group-018.md#canonical-3012102102201331-1221031112211213-1222123121211333-0030232213203203-0312010203111212-2102322332230021-0201012112121333-3231320112022100): complete subsection reference.

<a id="canonical-3132133011031120-1223310211331330-3333332101012110-0023333121120310-1131301122122131-0133000103100123-2312321331330113-2221322032130001"></a>

<a id="canonical-2301333013123012-1311301220123123-3102312210022220-1032120221312133-2113001112032213-3032330023210122-2331001321020310-0220020032310301"></a>

#### `simple_service.enabled.persistent_volume.storage.storage_size` property

Type: `"number"`. Optional.

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

<a id="canonical-3012102102201331-1221031112211213-1222123121211333-0030232213203203-0312010203111212-2102322332230021-0201012112121333-3231320112022100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `simple_service.enabled.persistent_volume.storage.default` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [simple_service](resources--workload--reference--group-017.md#canonical-2103120123000022-0231212311020021-3011030332210231-1230230313313103-1213231032231100-0321113222111303-0220111021013230-2333203212100121)
- [simple_service.enabled](resources--workload--reference--group-018.md#canonical-0022120332023020-2312100113020133-0123031032111130-3302333320122221-1232003333212033-1103032302323333-2231001310003211-0310320121030030)
- [simple_service.enabled.persistent_volume](resources--workload--reference--group-018.md#canonical-3113200332101222-1011331333203002-1210020131130330-0320221103232321-3010131320311330-1112331013232233-2101133032201232-1113223131002320)
- [simple_service.enabled.persistent_volume.storage](resources--workload--reference--group-018.md#canonical-0212031230210030-0110133221302101-1211130030200033-0222110222222113-1311020123323113-1302101310023300-1101100122133002-0221000223123310)
- simple_service.enabled.persistent_volume.storage.default

<a id="canonical-2222232222023233-3312203132202011-0303300232132200-3300331123133322-2031303031121030-3312031001313302-1231303221321120-3011301303203320"></a>

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
default = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3310302312301012-0213323302330302-2012000032301123-3011230301130001-2021102001021321-2321111011331112-1000322300330133-3212013211230202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `simple_service.simple_advertise` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [simple_service](resources--workload--reference--group-017.md#canonical-2103120123000022-0231212311020021-3011030332210231-1230230313313103-1213231032231100-0321113222111303-0220111021013230-2333203212100121)
- simple_service.simple_advertise

<a id="canonical-3130001232212321-0332110202020133-2132201100212101-0000322300221133-1210203232221223-2303320002123131-2312022013122221-1230131020132102"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
simple_advertise {
  # Configure direct properties listed below.
}
```

<a id="canonical-3221333133112213-0332220003031033-1200010210303310-3333120333132212-2030333221322231-2322211200002001-3002133220011022-0323011132110022"></a>

### Direct properties for `simple_service.simple_advertise`

<a id="canonical-0122012001223110-0120001223303221-3102231311113331-1121221211010232-3331100020203300-1112030012302303-2120231002101320-1102032302013231"></a>

#### `simple_service.simple_advertise.domains` property

Type: `["list", "string"]`. Optional.

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

<a id="canonical-3000000110013201-2013100122101320-0201322210302332-1031200303002001-0220303321303021-0231110213210333-0011302301233032-0300100221311013"></a>

<a id="canonical-2000100231231103-2022113212320033-0113030001030300-0310231011223320-3130111201120300-2203323013331300-3323121000122301-2331111213211333"></a>

#### `simple_service.simple_advertise.service_port` property

Type: `"number"`. Optional.

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

<a id="canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- stateful_service

<a id="canonical-3131320200112003-1302100202032202-0133102112232312-1213202222213131-2120333030221010-2233200311233003-3313123001231110-2320312322302201"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
stateful_service {
  # Configure direct properties listed below.
}
```

<a id="canonical-1022323221232223-0223203203302322-1033001223102223-2122131202221323-0013301100022211-0300130230002033-0101220101323013-1232212002123001"></a>

### Direct properties for `stateful_service`

- [advertise_options](resources--workload--reference--group-018.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323): complete subsection reference.

- [configuration](resources--workload--reference--group-028.md#canonical-3203012310220003-2210330300331110-2103223003030222-2222201323310103-2300110002113011-0130333021003321-0201020033321220-1102301103203230): complete subsection reference.

- [containers](resources--workload--reference--group-028.md#canonical-3221212333032323-1011122002030020-0332100102213333-2230123311320001-3101132000200220-3323302320310311-1200332113313321-1231002100113301): complete subsection reference.

- [deploy_options](resources--workload--reference--group-029.md#canonical-0211010322210330-0002130101323013-1320121223020011-3332231210010332-1331310201011323-2210300012032000-2333113200230221-2013202000203003): complete subsection reference.

<a id="canonical-1303111221221331-1323333220111212-3302320313231221-0200210101220130-2031031300320011-2000123031220021-2233201111020111-3123113032312120"></a>

<a id="canonical-1133203010131332-0300112021200321-3223133101100223-0303002212130201-0120011301102202-2121310303303210-3110201132031020-0321323132012322"></a>

#### `stateful_service.num_replicas` property

Type: `"number"`. Optional.

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

- [persistent_volumes](resources--workload--reference--group-029.md#canonical-2202000332001022-3031133313102112-0003233201110202-2303003230303311-3002213310230220-0131313123131032-0010302220210011-2300030100331332): complete subsection reference.

- [scale_to_zero](resources--workload--reference--group-029.md#canonical-0221221200321221-1133220231313302-3122113332002302-1303223333210013-3310313331033111-2320113332020010-2301111102010030-3110331212031230): complete subsection reference.

- [volumes](resources--workload--reference--group-029.md#canonical-0320321332330022-3200203220020212-3113213300013223-3331000302332100-2331130020201320-0110232101322132-3201120100233223-3122231121020213): complete subsection reference.

<a id="canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-018.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- stateful_service.advertise_options

<a id="canonical-3223330011020031-0330212002303111-2103131212302003-0310313201230333-1322000211100323-0301031101002100-1200021112333203-3200030221111101"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
advertise_options {
  # Configure direct properties listed below.
}
```

<a id="canonical-2010331033010110-2011311110010103-3021302230332311-2000030002111011-0031320200113003-3321212202030031-3303020112222222-1212013203302131"></a>

### Direct properties for `stateful_service.advertise_options`

- [advertise_custom](resources--workload--reference--group-018.md#canonical-0203102001320123-1120130012130212-0111101220231333-3101331123203121-1110000311003232-1321231201202312-3200200021132112-0200303231330002): complete subsection reference.

- [advertise_in_cluster](resources--workload--reference--group-022.md#canonical-2203200033111100-1131311021203132-1232001231010211-3000112112233321-3002013120002033-2111032232010121-0222102100212312-0012310120331313): complete subsection reference.

- [advertise_on_public](resources--workload--reference--group-022.md#canonical-3300101333202310-3223231311312310-2030111033033203-1011103231310023-1313110120311121-1221031312313020-1210302113302033-2110231331000202): complete subsection reference.

- [do_not_advertise](resources--workload--reference--group-028.md#canonical-3202321232103323-1000231221221030-0001122222031132-1003133321010023-1002201331103010-1230130022220311-0032100201113003-0333121223233010): complete subsection reference.

<a id="canonical-0203102001320123-1120130012130212-0111101220231333-3101331123203121-1110000311003232-1321231201202312-3200200021132112-0200303231330002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_custom` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-018.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-018.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- stateful_service.advertise_options.advertise_custom

<a id="canonical-1211201231012312-1032003212201201-3213201111203000-0232213003331200-3201332312303101-1200023022122231-3323232333232221-2010023221333112"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
advertise_custom {
  # Configure direct properties listed below.
}
```

<a id="canonical-3302231213303220-1200222203031002-1301202003113103-3213030121100331-1120020301311332-1230211103333322-0303100201210333-2132022023233033"></a>

### Direct properties for `stateful_service.advertise_options.advertise_custom`

- [advertise_where](resources--workload--reference--group-018.md#canonical-1202202011031000-3220322121323302-0033022130203330-3112100110300123-1310210111123123-3122221103132313-0322021321232201-3122311330133310): complete subsection reference.

- [ports](resources--workload--reference--group-018.md#canonical-3111311322332303-1201030020021201-3133021233000313-2022333331131202-3221312032323033-3133120120113232-3330021120310331-2331303302032231): complete subsection reference.

<a id="canonical-1202202011031000-3220322121323302-0033022130203330-3112100110300123-1310210111123123-3122221103132313-0322021321232201-3122311330133310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_custom.advertise_where` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-018.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-018.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_custom](resources--workload--reference--group-018.md#canonical-0203102001320123-1120130012130212-0111101220231333-3101331123203121-1110000311003232-1321231201202312-3200200021132112-0200303231330002)
- stateful_service.advertise_options.advertise_custom.advertise_where

<a id="canonical-3002101332033022-1110203201303323-3111222110113121-2013301001130212-3122120231011221-0103031132322131-0202303321122110-3220303302033333"></a>

Type: `"object"`. list nested block, Optional.

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

Terraform syntax:

```terraform
advertise_where {
  # Configure direct properties listed below.
}
```

<a id="canonical-0113022102202132-1101001210202203-0312021310110020-2022031210311222-2121220033220310-1230221002013301-3030331331101112-3331113213131120"></a>

### Direct properties for `stateful_service.advertise_options.advertise_custom.advertise_where`

- [site](resources--workload--reference--group-018.md#canonical-2223211010300233-0033203221303312-1010220021220302-2021010313033010-0231212132123031-3022301303002012-3030332322213232-3221210002300132): complete subsection reference.

- [virtual_site](resources--workload--reference--group-018.md#canonical-0031013123213022-1332033033011221-1030121103331012-1020201303132332-0101213311110012-2300122113201200-2102222023002321-1301121133120003): complete subsection reference.

- [vk8s_service](resources--workload--reference--group-018.md#canonical-3113201100321032-2110322211311302-0001210313012211-3231233213313313-3102200133233230-2000333233123013-1333110321022331-3013320123303013): complete subsection reference.

<a id="canonical-2223211010300233-0033203221303312-1010220021220302-2021010313033010-0231212132123031-3022301303002012-3030332322213232-3221210002300132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_custom.advertise_where.site` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-018.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-018.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_custom](resources--workload--reference--group-018.md#canonical-0203102001320123-1120130012130212-0111101220231333-3101331123203121-1110000311003232-1321231201202312-3200200021132112-0200303231330002)
- [stateful_service.advertise_options.advertise_custom.advertise_where](resources--workload--reference--group-018.md#canonical-1202202011031000-3220322121323302-0033022130203330-3112100110300123-1310210111123123-3122221103132313-0322021321232201-3122311330133310)
- stateful_service.advertise_options.advertise_custom.advertise_where.site

<a id="canonical-3030100122133203-0100200333101100-3012210101031123-2121200020330231-0031123232233300-2131032303301313-3101033323000321-2111021033022101"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
site {
  # Configure direct properties listed below.
}
```

<a id="canonical-3331321203112213-2230213110030012-2033313311322330-3030033131121130-2332112003201031-0013300002100312-3213131002302322-2012121011321112"></a>

### Direct properties for `stateful_service.advertise_options.advertise_custom.advertise_where.site`

<a id="canonical-2131231302322001-1000330012021011-2003031320011331-1102021013131200-1333130300222303-1220301200313130-3203332012120323-3021030301221010"></a>

#### `stateful_service.advertise_options.advertise_custom.advertise_where.site.ip` property

Type: `"string"`. Optional.

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

<a id="canonical-0323010001001310-2100333320222223-2132021223301130-3302332001030002-1312100023001330-1020123300220113-1132011312033333-1011111302330332"></a>

<a id="canonical-1213132200202121-3232011111220210-2221222321212120-1231100003121323-0210112120200023-2323001110012321-2332013222130200-1013113331033011"></a>

#### `stateful_service.advertise_options.advertise_custom.advertise_where.site.network` property

Type: `"string"`. Optional.

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

- [site](resources--workload--reference--group-018.md#canonical-2323021101211302-3220332313132300-3003021010013310-2112330231220120-3012002303222131-3001333101022320-3333013030233110-1023103020210111): complete subsection reference.

<a id="canonical-2323021101211302-3220332313132300-3003021010013310-2112330231220120-3012002303222131-3001333101022320-3333013030233110-1023103020210111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_custom.advertise_where.site.site` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-018.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-018.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_custom](resources--workload--reference--group-018.md#canonical-0203102001320123-1120130012130212-0111101220231333-3101331123203121-1110000311003232-1321231201202312-3200200021132112-0200303231330002)
- [stateful_service.advertise_options.advertise_custom.advertise_where](resources--workload--reference--group-018.md#canonical-1202202011031000-3220322121323302-0033022130203330-3112100110300123-1310210111123123-3122221103132313-0322021321232201-3122311330133310)
- [stateful_service.advertise_options.advertise_custom.advertise_where.site](resources--workload--reference--group-018.md#canonical-2223211010300233-0033203221303312-1010220021220302-2021010313033010-0231212132123031-3022301303002012-3030332322213232-3221210002300132)
- stateful_service.advertise_options.advertise_custom.advertise_where.site.site

<a id="canonical-2330000103132010-1101121120211222-2100133012313021-3220203232302231-2212100033212123-3003033033003203-1233032121131111-3011311322022312"></a>

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
site {
  # Configure direct properties listed below.
}
```

<a id="canonical-1212033032301112-2213232320101331-3031132233223201-0030221220032110-1221321212022331-1202001232213103-3120301202113300-1013022011332230"></a>

### Direct properties for `stateful_service.advertise_options.advertise_custom.advertise_where.site.site`

<a id="canonical-0300000300201101-2031002131022313-0301331331333211-2121233122210013-1310003120300123-2002123031222112-1112221323023032-0031220130130301"></a>

#### `stateful_service.advertise_options.advertise_custom.advertise_where.site.site.name` property

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

<a id="canonical-0030032102230132-3330031201221033-2023130231300311-0023103301221233-3232012313222113-2211111330310203-3001003103100333-2030211302211020"></a>

<a id="canonical-1200103011201030-0322102210213013-2130213032002123-0212102232121200-1312020121002000-0121111132232331-0323030323121122-2311110333200023"></a>

#### `stateful_service.advertise_options.advertise_custom.advertise_where.site.site.namespace` property

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

<a id="canonical-3000010211310313-2100123322311212-2223011023120002-0333030223312110-3332220301013003-0000102202102301-2210001331110200-3000321200102203"></a>

<a id="canonical-2032221333033003-2320323120320033-3331321111301021-3133312102110011-1312130031130301-3011131133333002-2311110330202321-0031133220003313"></a>

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

<a id="canonical-0031013123213022-1332033033011221-1030121103331012-1020201303132332-0101213311110012-2300122113201200-2102222023002321-1301121133120003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_custom.advertise_where.virtual_site` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-018.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-018.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_custom](resources--workload--reference--group-018.md#canonical-0203102001320123-1120130012130212-0111101220231333-3101331123203121-1110000311003232-1321231201202312-3200200021132112-0200303231330002)
- [stateful_service.advertise_options.advertise_custom.advertise_where](resources--workload--reference--group-018.md#canonical-1202202011031000-3220322121323302-0033022130203330-3112100110300123-1310210111123123-3122221103132313-0322021321232201-3122311330133310)
- stateful_service.advertise_options.advertise_custom.advertise_where.virtual_site

<a id="canonical-0220333133222202-1213221113130000-3210212200321203-3230111233102330-3212122111330023-1211102032122223-0123310322232321-0023033112112100"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
virtual_site {
  # Configure direct properties listed below.
}
```

<a id="canonical-3330213332303021-0311130132002031-2030310010132322-2221132330210123-0111131121001133-2230133133023332-3231033333203211-0322010313133021"></a>

### Direct properties for `stateful_service.advertise_options.advertise_custom.advertise_where.virtual_site`

<a id="canonical-0001022133021103-0021013300301032-3120332331130301-2123233012201020-1331121012033303-3210212000311220-0023130030001333-1220031211012002"></a>

#### `stateful_service.advertise_options.advertise_custom.advertise_where.virtual_site.network` property

Type: `"string"`. Optional.

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

- [virtual_site](resources--workload--reference--group-018.md#canonical-1333212202330222-1102223021022033-2001133332013310-3013122023221222-1100021323030200-3331213011031030-0333012123031301-3033223122131333): complete subsection reference.

<a id="canonical-1333212202330222-1102223021022033-2001133332013310-3013122023221222-1100021323030200-3331213011031030-0333012123031301-3033223122131333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_custom.advertise_where.virtual_site.virtual_site` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-018.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-018.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_custom](resources--workload--reference--group-018.md#canonical-0203102001320123-1120130012130212-0111101220231333-3101331123203121-1110000311003232-1321231201202312-3200200021132112-0200303231330002)
- [stateful_service.advertise_options.advertise_custom.advertise_where](resources--workload--reference--group-018.md#canonical-1202202011031000-3220322121323302-0033022130203330-3112100110300123-1310210111123123-3122221103132313-0322021321232201-3122311330133310)
- [stateful_service.advertise_options.advertise_custom.advertise_where.virtual_site](resources--workload--reference--group-018.md#canonical-0031013123213022-1332033033011221-1030121103331012-1020201303132332-0101213311110012-2300122113201200-2102222023002321-1301121133120003)
- stateful_service.advertise_options.advertise_custom.advertise_where.virtual_site.virtual_site

<a id="canonical-3112231022331201-0122320233201312-1113011321130010-1330333023030022-0001303222133023-1330200301130133-2321213131030000-2330220130100022"></a>

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
virtual_site {
  # Configure direct properties listed below.
}
```

<a id="canonical-1331130111223210-1123103223012311-3332321121211032-1220011013020222-0112312112220113-1232202210202121-1312332301001003-0132112022100133"></a>

### Direct properties for `stateful_service.advertise_options.advertise_custom.advertise_where.virtual_site.virtual_site`

<a id="canonical-1031222202322033-2232223033002112-2110030221232210-3202010133323211-0212031221000031-3310210320213002-0120310122030010-3032011323013000"></a>

#### `stateful_service.advertise_options.advertise_custom.advertise_where.virtual_site.virtual_site.name` property

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

<a id="canonical-3013210031012013-2001020300131230-2001110100330032-0322211202003223-2330321311222002-3321303103231230-1010112013113233-3333302101332113"></a>

<a id="canonical-1020303012330200-2030331021221003-3101023332012103-0232011331301323-3231130123113312-3000020122020231-1213100122331330-1033010203302101"></a>

#### `stateful_service.advertise_options.advertise_custom.advertise_where.virtual_site.virtual_site.namespace` property

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

<a id="canonical-1210320201031133-0112333113202130-3203310003130332-3300303110301213-2000011303110032-1202332013223023-3023212012000133-3113313330112312"></a>

<a id="canonical-1321310202113222-2232323100220102-0232213002332001-1103111111302131-2220220321320100-2310133303023222-1121213201000310-0102100002220022"></a>

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

<a id="canonical-3113201100321032-2110322211311302-0001210313012211-3231233213313313-3102200133233230-2000333233123013-1333110321022331-3013320123303013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_custom.advertise_where.vk8s_service` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-018.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-018.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_custom](resources--workload--reference--group-018.md#canonical-0203102001320123-1120130012130212-0111101220231333-3101331123203121-1110000311003232-1321231201202312-3200200021132112-0200303231330002)
- [stateful_service.advertise_options.advertise_custom.advertise_where](resources--workload--reference--group-018.md#canonical-1202202011031000-3220322121323302-0033022130203330-3112100110300123-1310210111123123-3122221103132313-0322021321232201-3122311330133310)
- stateful_service.advertise_options.advertise_custom.advertise_where.vk8s_service

<a id="canonical-2103022233101201-0210213120123221-0010323320303022-0121300023000103-3231021132010100-0321001232303102-3202003000111012-0312021331330302"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
vk8s_service {
  # Configure direct properties listed below.
}
```

<a id="canonical-1222321311220002-2020222221131120-0213301213121012-2120101113221320-0311022030112202-3033110032233302-1103221103203311-1311113330323300"></a>

### Direct properties for `stateful_service.advertise_options.advertise_custom.advertise_where.vk8s_service`

- [site](resources--workload--reference--group-018.md#canonical-2020100103012113-0010110101130222-1022200321203230-2222130032123120-1013300002223112-3121322123010013-0110323230231133-1310002322321131): complete subsection reference.

- [virtual_site](resources--workload--reference--group-018.md#canonical-3003101031302223-0033311311220102-1003320120221221-2231021023123210-3331032233001201-1033020022011030-3201113002230220-3232213303301221): complete subsection reference.

<a id="canonical-2020100103012113-0010110101130222-1022200321203230-2222130032123120-1013300002223112-3121322123010013-0110323230231133-1310002322321131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_custom.advertise_where.vk8s_service.site` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-018.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-018.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_custom](resources--workload--reference--group-018.md#canonical-0203102001320123-1120130012130212-0111101220231333-3101331123203121-1110000311003232-1321231201202312-3200200021132112-0200303231330002)
- [stateful_service.advertise_options.advertise_custom.advertise_where](resources--workload--reference--group-018.md#canonical-1202202011031000-3220322121323302-0033022130203330-3112100110300123-1310210111123123-3122221103132313-0322021321232201-3122311330133310)
- [stateful_service.advertise_options.advertise_custom.advertise_where.vk8s_service](resources--workload--reference--group-018.md#canonical-3113201100321032-2110322211311302-0001210313012211-3231233213313313-3102200133233230-2000333233123013-1333110321022331-3013320123303013)
- stateful_service.advertise_options.advertise_custom.advertise_where.vk8s_service.site

<a id="canonical-2021321331120132-0203203332000011-2302001223300012-2222332130313321-2300012203213100-2301302313233022-3333033332123302-3213132212232131"></a>

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
site {
  # Configure direct properties listed below.
}
```

<a id="canonical-3331303133002203-3011301323313132-3102311231020123-0202002300311220-0021211123132200-2011202030322020-2210220313300131-3002300002000013"></a>

### Direct properties for `stateful_service.advertise_options.advertise_custom.advertise_where.vk8s_service.site`

<a id="canonical-1002303202210120-2231121023110201-0322121002131002-1031301321312122-0033033310111333-3322230311310121-1103102310100121-2021331121312201"></a>

#### `stateful_service.advertise_options.advertise_custom.advertise_where.vk8s_service.site.name` property

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

<a id="canonical-1032232102111301-0122131032203001-3130323121030202-3311202211001112-2000233232111112-3131232201102202-3132210132232222-0103310012100212"></a>

<a id="canonical-2012220231031011-2303203100131110-2032311101001131-3102112312131223-2010122131332213-0233213113320310-0123023132202310-2202323001230231"></a>

#### `stateful_service.advertise_options.advertise_custom.advertise_where.vk8s_service.site.namespace` property

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

<a id="canonical-2103033322011011-3330320231012311-2111103013312122-0001023032012001-0331303200320113-3131213232010230-0300121011303001-1123311030320203"></a>

<a id="canonical-0321211030001132-1322302132122323-3300323310230201-3031103210030313-1021310022211123-1312303211203012-3312020103030100-0012000111300020"></a>

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

<a id="canonical-3003101031302223-0033311311220102-1003320120221221-2231021023123210-3331032233001201-1033020022011030-3201113002230220-3232213303301221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_custom.advertise_where.vk8s_service.virtual_site` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-018.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-018.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_custom](resources--workload--reference--group-018.md#canonical-0203102001320123-1120130012130212-0111101220231333-3101331123203121-1110000311003232-1321231201202312-3200200021132112-0200303231330002)
- [stateful_service.advertise_options.advertise_custom.advertise_where](resources--workload--reference--group-018.md#canonical-1202202011031000-3220322121323302-0033022130203330-3112100110300123-1310210111123123-3122221103132313-0322021321232201-3122311330133310)
- [stateful_service.advertise_options.advertise_custom.advertise_where.vk8s_service](resources--workload--reference--group-018.md#canonical-3113201100321032-2110322211311302-0001210313012211-3231233213313313-3102200133233230-2000333233123013-1333110321022331-3013320123303013)
- stateful_service.advertise_options.advertise_custom.advertise_where.vk8s_service.virtual_site

<a id="canonical-0123302000200031-3003123302123202-3000303113233003-1212302221222031-1302311201103300-1310330222002303-1120022012031313-0112232303300002"></a>

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
virtual_site {
  # Configure direct properties listed below.
}
```

<a id="canonical-1333222303302131-0123001210230212-1312022303131323-2201222023030102-2101333121030021-1033300023223132-2322123222301110-3232332201301003"></a>

### Direct properties for `stateful_service.advertise_options.advertise_custom.advertise_where.vk8s_service.virtual_site`

<a id="canonical-1321330033133012-1233221220030003-0112211103011012-2233332112303100-3011003112131033-2122230302132320-2032030333311032-0322230321002110"></a>

#### `stateful_service.advertise_options.advertise_custom.advertise_where.vk8s_service.virtual_site.name` property

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

<a id="canonical-1103330123133031-3023132311220333-2102100213003230-3322123031210213-3023332230220231-3322113300321002-0113333331000020-0113320303223223"></a>

<a id="canonical-2130211221202310-2102322223223003-1002101030212303-0231013030133302-0311113302120312-0230333100032321-2012032002310003-0111030121221323"></a>

#### `stateful_service.advertise_options.advertise_custom.advertise_where.vk8s_service.virtual_site.namespace` property

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

<a id="canonical-3103000031321332-0231130232010110-2200221101223112-3021012101200031-2313303333110012-0301130023332231-3203211030020322-3023313002020122"></a>

<a id="canonical-2302003220202301-1301221123212301-3303313022232133-3311103031102101-0301010200320300-3323002113013012-0023023122223022-2330102112010113"></a>

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

<a id="canonical-3111311322332303-1201030020021201-3133021233000313-2022333331131202-3221312032323033-3133120120113232-3330021120310331-2331303302032231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_custom.ports` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-018.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-018.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_custom](resources--workload--reference--group-018.md#canonical-0203102001320123-1120130012130212-0111101220231333-3101331123203121-1110000311003232-1321231201202312-3200200021132112-0200303231330002)
- stateful_service.advertise_options.advertise_custom.ports

<a id="canonical-2110313010033020-0020221230302220-2320320022122010-0231201122002231-3313000303103002-2213120300202302-3333311300012020-1312232001022110"></a>

Type: `"object"`. list nested block, Optional.

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

Terraform syntax:

```terraform
ports {
  # Configure direct properties listed below.
}
```

<a id="canonical-2103220210020311-3012023013130001-2000121021330011-0010202323200100-0333113313213212-1313012031310011-0000001200311310-0122103020010202"></a>

### Direct properties for `stateful_service.advertise_options.advertise_custom.ports`

- [http_loadbalancer](resources--workload--reference--group-018.md#canonical-0313222022003221-2301103230203023-3120322210033100-2121220010332201-3133231021021220-1313001311100021-0313102021211212-0331322021202002): complete subsection reference.

- [port](resources--workload--reference--group-021.md#canonical-2103101222313132-0030333210220102-2031201232312001-1103110230303213-2123210111312322-2333302111312301-1031201133012330-3303120000023223): complete subsection reference.

- [tcp_loadbalancer](resources--workload--reference--group-022.md#canonical-0020101203031230-1223300211330033-2123113022210223-3310333011133100-3330221301313310-3131021233011311-3221200100200012-0300221100330331): complete subsection reference.

<a id="canonical-0313222022003221-2301103230203023-3120322210033100-2121220010332201-3133231021021220-1313001311100021-0313102021211212-0331322021202002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-018.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-018.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_custom](resources--workload--reference--group-018.md#canonical-0203102001320123-1120130012130212-0111101220231333-3101331123203121-1110000311003232-1321231201202312-3200200021132112-0200303231330002)
- [stateful_service.advertise_options.advertise_custom.ports](resources--workload--reference--group-018.md#canonical-3111311322332303-1201030020021201-3133021233000313-2022333331131202-3221312032323033-3133120120113232-3330021120310331-2331303302032231)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer

<a id="canonical-3233013221221131-2012133322101333-2333103311222030-1300332102121212-3213323301000303-2100212310212031-3102012233011233-3232213013123002"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
http_loadbalancer {
  # Configure direct properties listed below.
}
```

<a id="canonical-1322200313001203-0021221323310002-3013233010031033-3331132203113302-2321110310023333-0203312012321110-1201203211233333-3032323132210023"></a>

### Direct properties for `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer`

- [default_route](resources--workload--reference--group-018.md#canonical-2031112003330101-0310120302332023-3120021113320130-3110000310223201-1012000003321130-2312213320212011-0022300121030033-3100123121103201): complete subsection reference.

<a id="canonical-3122002032322020-1000311301022012-2113012331300210-3101110120322332-2103231112330231-2022313201323200-0133313033130311-1212231000203011"></a>

<a id="canonical-2131000232000212-0232120112133222-3122221022102011-3001012201200331-3122022333011332-3010331121001131-3223030121133211-1220031133213213"></a>

#### `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.domains` property

Type: `["list", "string"]`. Optional.

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

- [http](resources--workload--reference--group-019.md#canonical-0001001230203210-1010213221120331-1102200030010210-1101322011000303-0020221022232213-1231230310223023-2110212231011313-1321232000313021): complete subsection reference.

- [https](resources--workload--reference--group-019.md#canonical-0112020200111022-0013103013103331-1110301001233101-1300202101312213-3301031210112102-1221232301300211-0231000232130131-3132113303133333): complete subsection reference.

- [https_auto_cert](resources--workload--reference--group-020.md#canonical-3032121232310131-1323102031301330-0021130303300211-3312120223313333-2332212203311131-0113312021100231-3031113230222323-1032033323320101): complete subsection reference.

- [specific_routes](resources--workload--reference--group-020.md#canonical-0102002011120221-0321102211211200-1233000013222030-0112300012022023-1023113003101102-0220010103202030-2311210030210102-2200103323322232): complete subsection reference.

<a id="canonical-2031112003330101-0310120302332023-3120021113320130-3110000310223201-1012000003321130-2312213320212011-0022300121030033-3100123121103201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.default_route` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-018.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-018.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_custom](resources--workload--reference--group-018.md#canonical-0203102001320123-1120130012130212-0111101220231333-3101331123203121-1110000311003232-1321231201202312-3200200021132112-0200303231330002)
- [stateful_service.advertise_options.advertise_custom.ports](resources--workload--reference--group-018.md#canonical-3111311322332303-1201030020021201-3133021233000313-2022333331131202-3221312032323033-3133120120113232-3330021120310331-2331303302032231)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-018.md#canonical-0313222022003221-2301103230203023-3120322210033100-2121220010332201-3133231021021220-1313001311100021-0313102021211212-0331322021202002)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.default_route

<a id="canonical-2012112020211212-2002133331032331-0301122010221321-1233000233300003-0002212101002131-0210220123322310-1033033101022002-2112000101333121"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
default_route {
  # Configure direct properties listed below.
}
```

<a id="canonical-3202011303331332-2303132010013121-0013322312232221-2322122030123212-2332300133130110-2002123002223213-2211330102232200-1023222131200021"></a>

### Direct properties for `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.default_route`

- [auto_host_rewrite](resources--workload--reference--group-019.md#canonical-3102103313032210-0102301102011232-1303110012101212-3032312111101121-2212320022233233-2110332122303130-0322022331103012-2002002322210003): complete subsection reference.

- [disable_host_rewrite](resources--workload--reference--group-019.md#canonical-2233022211333231-1323213121201103-1312330322030123-0003203301103101-0332011031223232-2322233001313302-0312301223030330-1003213211032213): complete subsection reference.

<a id="canonical-3123011213032330-0310102133232303-2213001223032221-2132211311231011-2011231300311232-1120310313313203-2300202213032233-3011121100110321"></a>

<a id="canonical-2212333013311332-1112132230113102-0201211300333021-1030330033100111-3012022301232110-2010322310002131-0330130102030302-0200223333330022"></a>

#### `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.default_route.host_rewrite` property

Type: `"string"`. Optional.

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
