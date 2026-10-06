---
page_title: "xcsh_workload reference"
subcategory: "Container"
description: "Complete grouped canonical reference for xcsh_workload reference."
---

# xcsh_workload reference

<a id="canonical-0112230132313323-2122103132021131-1012002001203012-1000103302310330-0032032123212331-0123013323310310-2230212223000111-0220101331331003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.simple_route.auto_host_rewrite` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-2121230312030330-3010113231323001-1233013032320322-1032203002210311-2022021002113232-1123312032112001-0012103213230102-1112311001110020)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-006.md#canonical-0133221101333210-0222103021001232-3212010032013103-3331120032012221-1012203003232110-1233102323001233-2233321121201101-0122313122210200)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-006.md#canonical-2311001323203002-2023333020201320-1330320322101233-1033123102030111-2013210002333120-2130122003123103-1133301323123220-2302130221012102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes](resources--workload--reference--group-007.md#canonical-3011211302110102-1110003330130012-2330003231002123-1333212203033211-1212101000211203-3333313212202213-1122232332203332-0220110332323232)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes](resources--workload--reference--group-007.md#canonical-3101223212120012-3332300223001100-0333220310012333-1031332113331120-2210213133122022-1333032223122231-2030231210333100-1103221131133311)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.simple_route](resources--workload--reference--group-008.md#canonical-2332230303301330-0120223130133121-2131103033210110-0000331031322110-2102032123012022-0000121011033231-1222320000321202-3323210133011002)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.simple_route.auto_host_rewrite

<a id="canonical-3020001302123220-3032021000301010-0312301322220200-3332312312310232-2212301110110320-0333302321001121-2030223301130312-3223212123122120"></a>

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

<a id="canonical-1111112032212222-1331322310310203-3123130031332003-3000123221101021-0102220030220211-3130102131232020-3010130002220003-2212212321323022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.simple_route.disable_host_rewrite` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-2121230312030330-3010113231323001-1233013032320322-1032203002210311-2022021002113232-1123312032112001-0012103213230102-1112311001110020)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-006.md#canonical-0133221101333210-0222103021001232-3212010032013103-3331120032012221-1012203003232110-1233102323001233-2233321121201101-0122313122210200)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-006.md#canonical-2311001323203002-2023333020201320-1330320322101233-1033123102030111-2013210002333120-2130122003123103-1133301323123220-2302130221012102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes](resources--workload--reference--group-007.md#canonical-3011211302110102-1110003330130012-2330003231002123-1333212203033211-1212101000211203-3333313212202213-1122232332203332-0220110332323232)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes](resources--workload--reference--group-007.md#canonical-3101223212120012-3332300223001100-0333220310012333-1031332113331120-2210213133122022-1333032223122231-2030231210333100-1103221131133311)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.simple_route](resources--workload--reference--group-008.md#canonical-2332230303301330-0120223130133121-2131103033210110-0000331031322110-2102032123012022-0000121011033231-1222320000321202-3323210133011002)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.simple_route.disable_host_rewrite

<a id="canonical-3013131201302111-3200103322030120-3211031323212022-2023230122322203-0132210130102222-3122311100232130-0100223111333101-3002100012031012"></a>

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

<a id="canonical-3221021331221300-2122202121211303-0333120310113302-2003110323312131-1320212313100121-1101212002232103-1001323321022113-3210123212220031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.simple_route.path` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-2121230312030330-3010113231323001-1233013032320322-1032203002210311-2022021002113232-1123312032112001-0012103213230102-1112311001110020)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-006.md#canonical-0133221101333210-0222103021001232-3212010032013103-3331120032012221-1012203003232110-1233102323001233-2233321121201101-0122313122210200)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-006.md#canonical-2311001323203002-2023333020201320-1330320322101233-1033123102030111-2013210002333120-2130122003123103-1133301323123220-2302130221012102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes](resources--workload--reference--group-007.md#canonical-3011211302110102-1110003330130012-2330003231002123-1333212203033211-1212101000211203-3333313212202213-1122232332203332-0220110332323232)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes](resources--workload--reference--group-007.md#canonical-3101223212120012-3332300223001100-0333220310012333-1031332113331120-2210213133122022-1333032223122231-2030231210333100-1103221131133311)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.simple_route](resources--workload--reference--group-008.md#canonical-2332230303301330-0120223130133121-2131103033210110-0000331031322110-2102032123012022-0000121011033231-1222320000321202-3323210133011002)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.simple_route.path

<a id="canonical-2233001333001332-2303111302302332-0202302322123220-1130113312221013-1332110210121221-3220120321013001-2321331130311303-3113023221202222"></a>

Type: `"object"`. single nested block, Optional.

Path match of the URI can be either be, Prefix match or exact match or regular expression match.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("path",
    "prefix"),
  validators.ConflictingObjectAttributes("path",
    "regex"),
  validators.ConflictingObjectAttributes("prefix",
    "regex")}
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
  "x-ves-oneof-field-path_match": "[\"path\",\"prefix\",\"regex\"]"
}
```

Terraform syntax:

```terraform
path {
  # Configure direct properties listed below.
}
```

<a id="canonical-2232330312223222-2213113302122220-2103111133032330-3331102022332003-3023203331021123-1133000103323030-2123230211331013-3113022313033223"></a>

### Direct properties for `service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.simple_route.path`

<a id="canonical-0322133032200232-2132032232022301-3002330223220223-2230033003112322-3232112213331330-0321002121010012-3010130310101133-3330032230110120"></a>

#### `service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.simple_route.path.path` property

Type: `"string"`. Optional.

Exclusive with \[prefix regular expression\] Exact path value to match.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-3033133211102031-0210002131003012-1001020331002313-1122212103103133-1230012021113022-0330122332123010-3113331230000132-0232000002301211"></a>

<a id="canonical-1321313232101020-0031333013231121-1333102202122232-3101221012131022-3112311223023310-3221323032200110-0200000310122233-2312203020111102"></a>

#### `service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.simple_route.path.prefix` property

Type: `"string"`. Optional.

Exclusive with \[path regular expression\] Path prefix to match (e.g. The value / will match on all paths)

Provider validators and defaults (from schema source):

```go
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
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-0330323132022321-3203032310233020-0003231130130101-3002120102210110-1203123113113213-3220212033011012-0113210033132000-3121231311030012"></a>

<a id="canonical-2210030130330012-3112231310121032-2123030121102123-2311201333020032-0301003130332123-2101103113010123-1002200300220302-2102032001002103"></a>

#### `service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.simple_route.path.regex` property

Type: `"string"`. Optional.

Exclusive with \[path prefix\] Regular expression of path match (e.g. The value .\* will match on
all paths).

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

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

<a id="canonical-2011231233103202-1130133301231221-3112112222130011-0133100320022222-3002023210301231-3323320331313221-3230300303311111-0302223303200123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_custom.ports.port` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-2121230312030330-3010113231323001-1233013032320322-1032203002210311-2022021002113232-1123312032112001-0012103213230102-1112311001110020)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-006.md#canonical-0133221101333210-0222103021001232-3212010032013103-3331120032012221-1012203003232110-1233102323001233-2233321121201101-0122313122210200)
- service.advertise_options.advertise_custom.ports.port

<a id="canonical-0303220310031220-2231211003010121-2130101032232221-3001101213022122-0330311011320112-2211322120222030-2133311323320300-1103211130031312"></a>

Type: `"object"`. single nested block, Optional.

Port. Port of the workload.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("name")}
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
port {
  # Configure direct properties listed below.
}
```

<a id="canonical-2000312131021111-1322002100000322-3001222320013321-0010203332221311-1033201021322303-0201313210131122-2133003110330332-1313131300231110"></a>

### Direct properties for `service.advertise_options.advertise_custom.ports.port`

- [info](resources--workload--reference--group-009.md#canonical-1122232230031120-0022203213323110-0233011123013101-3033030322231311-0002300310301212-3213332001002102-0103000132313210-3031331333211321): complete subsection reference.

<a id="canonical-0023000021233322-3210122310230003-3133020310012111-0330331320122320-3221113220030013-3021322010001231-2202303311022010-0322231210130020"></a>

<a id="canonical-3033303032032110-2002330332133132-2020011122101331-1320002101313011-1301030222100212-2301011031221323-3310011230022130-3313211233330230"></a>

#### `service.advertise_options.advertise_custom.ports.port.name` property

Type: `"string"`. Optional.

Name. Name of the Port.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-1122232230031120-0022203213323110-0233011123013101-3033030322231311-0002300310301212-3213332001002102-0103000132313210-3031331333211321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_custom.ports.port.info` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-2121230312030330-3010113231323001-1233013032320322-1032203002210311-2022021002113232-1123312032112001-0012103213230102-1112311001110020)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-006.md#canonical-0133221101333210-0222103021001232-3212010032013103-3331120032012221-1012203003232110-1233102323001233-2233321121201101-0122313122210200)
- [service.advertise_options.advertise_custom.ports.port](resources--workload--reference--group-009.md#canonical-2011231233103202-1130133301231221-3112112222130011-0133100320022222-3002023210301231-3323320331313221-3230300303311111-0302223303200123)
- service.advertise_options.advertise_custom.ports.port.info

<a id="canonical-0213020202001320-1303321131021231-0321003301313021-1203023131303220-2300321033130002-0332110330301121-0201213033203212-1100001020102031"></a>

Type: `"object"`. single nested block, Optional.

Port Information. Port information.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("port"),
  validators.ConflictingObjectAttributes("same_as_port",
    "target_port")}
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
  "x-ves-oneof-field-target_port_choice": "[\"same_as_port\",\"target_port\"]"
}
```

Terraform syntax:

```terraform
info {
  # Configure direct properties listed below.
}
```

<a id="canonical-3113332212100112-2202131001232123-1111100030312232-0030012321033322-0031010202110332-2033123322130221-0012023323232032-3030322332021120"></a>

### Direct properties for `service.advertise_options.advertise_custom.ports.port.info`

<a id="canonical-3121333321222000-2003212003020111-1030002011031030-3102203233310312-2211330113133222-1230013222032000-2201011323220120-0031331310323333"></a>

#### `service.advertise_options.advertise_custom.ports.port.info.port` property

Type: `"number"`. Optional.

Port. Port the workload can be reached on.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-2130113332133012-3321112320130113-2221132013233301-1223132011211110-3311213231113003-2022331022001012-3213320010101102-3110111032013302"></a>

<a id="canonical-2313033323023301-0001202030113222-3133111022302303-2011003220013202-0011202322001022-2123330303223032-2332022302311012-0322303333033020"></a>

#### `service.advertise_options.advertise_custom.ports.port.info.protocol` property

Type: `"string"`. Optional.

\[Enum: PROTOCOL\_TCP|PROTOCOL\_HTTP|PROTOCOL\_HTTP2|PROTOCOL\_TLS\_WITH\_SNI|PROTOCOL\_UDP\] Type
of protocol - PROTOCOL\_TCP: TCP TCP - PROTOCOL\_HTTP: HTTP HTTP - PROTOCOL\_HTTP2: HTTP2 HTTP2 -
PROTOCOL\_TLS\_WITH\_SNI: TLS with SNI TLS with SNI - PROTOCOL\_UDP: UDP UDP. Possible values are
\`PROTOCOL\_TCP\`, \`PROTOCOL\_HTTP\`, \`PROTOCOL\_HTTP2\`, \`PROTOCOL\_TLS\_WITH\_SNI\`,
\`PROTOCOL\_UDP\`. Defaults to \`PROTOCOL\_TCP\`.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("PROTOCOL_TCP",
    "PROTOCOL_HTTP",
    "PROTOCOL_HTTP2",
    "PROTOCOL_TLS_WITH_SNI",
    "PROTOCOL_UDP"),
}
```

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

- [same_as_port](resources--workload--reference--group-009.md#canonical-1110030120221301-2033120210021001-0101031112323320-0213132221023103-0223022301323322-1022213103310103-2011031302113000-1011103331323132): complete subsection reference.

<a id="canonical-1213323310203333-3101130010212232-1133311103130032-3022003211030311-1013000230233330-0313121111120121-0021302223010123-0131032002020310"></a>

<a id="canonical-0311210300001112-0221100103100020-1220232032013100-2221203333321000-3312220011121033-0030232022031123-3021133203031021-0130011133231100"></a>

#### `service.advertise_options.advertise_custom.ports.port.info.target_port` property

Type: `"number"`. Optional.

Exclusive with \[same\_as\_port\] Port the workload is listening on.

Provider validators and defaults (from schema source):

```go
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
    "ves.io.schema.rules.uint32.gt": "0",
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gt": "0",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-1110030120221301-2033120210021001-0101031112323320-0213132221023103-0223022301323322-1022213103310103-2011031302113000-1011103331323132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_custom.ports.port.info.same_as_port` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-2121230312030330-3010113231323001-1233013032320322-1032203002210311-2022021002113232-1123312032112001-0012103213230102-1112311001110020)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-006.md#canonical-0133221101333210-0222103021001232-3212010032013103-3331120032012221-1012203003232110-1233102323001233-2233321121201101-0122313122210200)
- [service.advertise_options.advertise_custom.ports.port](resources--workload--reference--group-009.md#canonical-2011231233103202-1130133301231221-3112112222130011-0133100320022222-3002023210301231-3323320331313221-3230300303311111-0302223303200123)
- [service.advertise_options.advertise_custom.ports.port.info](resources--workload--reference--group-009.md#canonical-1122232230031120-0022203213323110-0233011123013101-3033030322231311-0002300310301212-3213332001002102-0103000132313210-3031331333211321)
- service.advertise_options.advertise_custom.ports.port.info.same_as_port

<a id="canonical-1330133210030312-0100112101001102-0012030120131323-2213310121331210-3332013323130330-1232301301130103-1000010033011211-1131010012013133"></a>

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

<a id="canonical-2333000323211321-3320220213111111-3000113231121121-2211220223011110-3303030010002012-3322030301222113-2111223031202010-0320112332022000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_custom.ports.tcp_loadbalancer` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-2121230312030330-3010113231323001-1233013032320322-1032203002210311-2022021002113232-1123312032112001-0012103213230102-1112311001110020)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-006.md#canonical-0133221101333210-0222103021001232-3212010032013103-3331120032012221-1012203003232110-1233102323001233-2233321121201101-0122313122210200)
- service.advertise_options.advertise_custom.ports.tcp_loadbalancer

<a id="canonical-2100023101202221-1121112011012321-1010332030111311-1232012122332331-2011223133312332-3311201231213313-0000332123100013-3301321001033101"></a>

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

<a id="canonical-0120222130300023-3233330321032112-0333212200202303-0002310313302132-3223312022302230-0022010310231303-1003222211320013-2311002003030022"></a>

### Direct properties for `service.advertise_options.advertise_custom.ports.tcp_loadbalancer`

<a id="canonical-3300103310101212-3023222200221303-0010222232310312-1322201021232112-2110012302301311-3113012120331232-1231122021233310-2333330130101200"></a>

#### `service.advertise_options.advertise_custom.ports.tcp_loadbalancer.domains` property

Type: `["list", "string"]`. Optional.

List of additional domains (host/authority header) that will be matched to this loadbalancer.
Domains are also used for SNI matching if the is true Domains also indicate the list of names for
which DNS resolution will be done by VER.

Additional upstream details:

A list of additional domains (host/authority header) that will be matched to this loadbalancer.
Domains are also used for SNI matching if the \`with\_sni\` is true Domains also indicate the list
of names for which DNS resolution will be done by VER.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeBetween(1, 32),
}
```

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

<a id="canonical-1300131122302020-0330022100313311-2012212200121003-1001101012200320-1333123111120300-0020003323023132-0102203001103331-1113211132013111"></a>

<a id="canonical-3021032321001211-2001333313320310-2122202330333031-1203121002002122-2000100102012303-2011000232010101-2332313021003231-2312320313203102"></a>

#### `service.advertise_options.advertise_custom.ports.tcp_loadbalancer.with_sni` property

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

<a id="canonical-0330023301131021-0130211002331300-3002122032320301-2112231133101233-3303101211030122-3300030030023300-1213020320030020-0213001312211213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_in_cluster` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- service.advertise_options.advertise_in_cluster

<a id="canonical-3303010210113203-0000123133102021-2220101212333303-2022210232202312-1121210030320001-2122031211022321-0021132013231232-2231133202303313"></a>

Type: `"object"`. single nested block, Optional.

Advertise the workload locally in-cluster.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("multi_ports",
    "port")}
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
  "x-ves-oneof-field-port_choice": "[\"multi_ports\",\"port\"]"
}
```

Terraform syntax:

```terraform
advertise_in_cluster {
  # Configure direct properties listed below.
}
```

<a id="canonical-0330220031113300-0003030303312010-3230010030031303-2330233220133033-2012210020230323-3020122113211001-3033031213012202-0002313032203301"></a>

### Direct properties for `service.advertise_options.advertise_in_cluster`

- [multi_ports](resources--workload--reference--group-009.md#canonical-3233102030323303-0010001111332222-2221330300220032-2322233002300322-1030033321333231-1200300001103302-1331010302322302-2231113103111232): complete subsection reference.

- [port](resources--workload--reference--group-009.md#canonical-2032101121303011-1103222233000330-0203232311001010-2322120323102000-3202033332221123-2310113223210032-0131222201221033-2112013321213130): complete subsection reference.

<a id="canonical-3233102030323303-0010001111332222-2221330300220032-2322233002300322-1030033321333231-1200300001103302-1331010302322302-2231113103111232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_in_cluster.multi_ports` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_in_cluster](resources--workload--reference--group-009.md#canonical-0330023301131021-0130211002331300-3002122032320301-2112231133101233-3303101211030122-3300030030023300-1213020320030020-0213001312211213)
- service.advertise_options.advertise_in_cluster.multi_ports

<a id="canonical-0010032311313302-2100112001110300-0332012333311013-1301111313330310-3033232010320211-1312310313022101-1123110330123221-0022210010330200"></a>

Type: `"object"`. single nested block, Optional.

Multiple Ports. Multiple ports.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("ports")}
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
multi_ports {
  # Configure direct properties listed below.
}
```

<a id="canonical-2121202131001323-2330111213331230-0212203132030221-2132211323022111-3310123332033313-0010331001230213-1122020300200303-2032130221210202"></a>

### Direct properties for `service.advertise_options.advertise_in_cluster.multi_ports`

- [ports](resources--workload--reference--group-009.md#canonical-0302230202101333-3033110013223110-0210311023223222-0203030332102133-2201333301103120-1000302302100130-2222303013302022-1022120223133312): complete subsection reference.

<a id="canonical-0302230202101333-3033110013223110-0210311023223222-0203030332102133-2201333301103120-1000302302100130-2222303013302022-1022120223133312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_in_cluster.multi_ports.ports` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_in_cluster](resources--workload--reference--group-009.md#canonical-0330023301131021-0130211002331300-3002122032320301-2112231133101233-3303101211030122-3300030030023300-1213020320030020-0213001312211213)
- [service.advertise_options.advertise_in_cluster.multi_ports](resources--workload--reference--group-009.md#canonical-3233102030323303-0010001111332222-2221330300220032-2322233002300322-1030033321333231-1200300001103302-1331010302322302-2231113103111232)
- service.advertise_options.advertise_in_cluster.multi_ports.ports

<a id="canonical-3310220231101131-2122202103203132-2032032312323112-0111111123010110-0210101110331033-3311220102322010-2300011333310100-3021201133033222"></a>

Type: `"object"`. list nested block, Optional.

Ports. Ports to advertise.

Provider validators and defaults (from schema source):

```go
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
ports {
  # Configure direct properties listed below.
}
```

<a id="canonical-3101111113211202-3212031000032220-2230230213230302-1312223130030012-2231222033032331-1101133300332102-2330100330301210-2223123203000103"></a>

### Direct properties for `service.advertise_options.advertise_in_cluster.multi_ports.ports`

- [info](resources--workload--reference--group-009.md#canonical-1032330310201200-3101102022121101-0202311121001130-1310203222000011-1113321122130331-0113223320013030-1012012320323211-2111201022201331): complete subsection reference.

<a id="canonical-0211032121020212-0320101100100331-2030002032011211-2133120033210201-1201202203121000-2303322233330300-3011300100221030-1302003300331112"></a>

<a id="canonical-2033313000001030-3130111022233102-3310333033111030-1130001320101323-3201202122330332-0312022313103002-1302023022033322-0321000210330333"></a>

#### `service.advertise_options.advertise_in_cluster.multi_ports.ports.name` property

Type: `"string"`. Optional.

Name. Name of the Port.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-1032330310201200-3101102022121101-0202311121001130-1310203222000011-1113321122130331-0113223320013030-1012012320323211-2111201022201331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_in_cluster.multi_ports.ports.info` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_in_cluster](resources--workload--reference--group-009.md#canonical-0330023301131021-0130211002331300-3002122032320301-2112231133101233-3303101211030122-3300030030023300-1213020320030020-0213001312211213)
- [service.advertise_options.advertise_in_cluster.multi_ports](resources--workload--reference--group-009.md#canonical-3233102030323303-0010001111332222-2221330300220032-2322233002300322-1030033321333231-1200300001103302-1331010302322302-2231113103111232)
- [service.advertise_options.advertise_in_cluster.multi_ports.ports](resources--workload--reference--group-009.md#canonical-0302230202101333-3033110013223110-0210311023223222-0203030332102133-2201333301103120-1000302302100130-2222303013302022-1022120223133312)
- service.advertise_options.advertise_in_cluster.multi_ports.ports.info

<a id="canonical-2212212020221220-2223322222013330-3113113200233311-3132033221200000-1010220031230330-2100032331322220-0102331032210211-0032331101303110"></a>

Type: `"object"`. single nested block, Optional.

Port Information. Port information.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("port"),
  validators.ConflictingObjectAttributes("same_as_port",
    "target_port")}
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
  "x-ves-oneof-field-target_port_choice": "[\"same_as_port\",\"target_port\"]"
}
```

Terraform syntax:

```terraform
info {
  # Configure direct properties listed below.
}
```

<a id="canonical-2310030110231120-1232203023211031-3310322122130032-3112203112130110-0001111211203312-2122011230031300-3013313101000023-2102213112030202"></a>

### Direct properties for `service.advertise_options.advertise_in_cluster.multi_ports.ports.info`

<a id="canonical-3121101201110211-1101213022120113-0211330330113021-0100031313112122-1323300222332221-1331102122010211-3001012123210231-3303321312013331"></a>

#### `service.advertise_options.advertise_in_cluster.multi_ports.ports.info.port` property

Type: `"number"`. Optional.

Port. Port the workload can be reached on.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-3203030333002131-3013112331220123-1033102222310020-2110213032002311-2312312312003330-1222331111031320-2010021123002110-1000203203323021"></a>

<a id="canonical-0110200323111113-2113210211213103-2331313020320230-0222302132313120-0011030112131011-2201232203220200-0302222230003311-1323210213013110"></a>

#### `service.advertise_options.advertise_in_cluster.multi_ports.ports.info.protocol` property

Type: `"string"`. Optional.

\[Enum: PROTOCOL\_TCP|PROTOCOL\_HTTP|PROTOCOL\_HTTP2|PROTOCOL\_TLS\_WITH\_SNI|PROTOCOL\_UDP\] Type
of protocol - PROTOCOL\_TCP: TCP TCP - PROTOCOL\_HTTP: HTTP HTTP - PROTOCOL\_HTTP2: HTTP2 HTTP2 -
PROTOCOL\_TLS\_WITH\_SNI: TLS with SNI TLS with SNI - PROTOCOL\_UDP: UDP UDP. Possible values are
\`PROTOCOL\_TCP\`, \`PROTOCOL\_HTTP\`, \`PROTOCOL\_HTTP2\`, \`PROTOCOL\_TLS\_WITH\_SNI\`,
\`PROTOCOL\_UDP\`. Defaults to \`PROTOCOL\_TCP\`.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("PROTOCOL_TCP",
    "PROTOCOL_HTTP",
    "PROTOCOL_HTTP2",
    "PROTOCOL_TLS_WITH_SNI",
    "PROTOCOL_UDP"),
}
```

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

- [same_as_port](resources--workload--reference--group-009.md#canonical-3203022201002002-3230201120302112-2201031221333112-2131003320203232-0313201303030011-2101023111002031-2322011130001010-0231013131320131): complete subsection reference.

<a id="canonical-1032202030013211-0202010310302001-1323333032331310-0323120012210031-2012210221013002-3011132221003323-0110111333003113-1202131203000022"></a>

<a id="canonical-1220030212122020-0211033113221113-1100211031233331-0332230221210223-1201132023330321-2020330022233032-3313321321010203-1211200122300001"></a>

#### `service.advertise_options.advertise_in_cluster.multi_ports.ports.info.target_port` property

Type: `"number"`. Optional.

Exclusive with \[same\_as\_port\] Port the workload is listening on.

Provider validators and defaults (from schema source):

```go
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
    "ves.io.schema.rules.uint32.gt": "0",
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gt": "0",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-3203022201002002-3230201120302112-2201031221333112-2131003320203232-0313201303030011-2101023111002031-2322011130001010-0231013131320131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_in_cluster.multi_ports.ports.info.same_as_port` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_in_cluster](resources--workload--reference--group-009.md#canonical-0330023301131021-0130211002331300-3002122032320301-2112231133101233-3303101211030122-3300030030023300-1213020320030020-0213001312211213)
- [service.advertise_options.advertise_in_cluster.multi_ports](resources--workload--reference--group-009.md#canonical-3233102030323303-0010001111332222-2221330300220032-2322233002300322-1030033321333231-1200300001103302-1331010302322302-2231113103111232)
- [service.advertise_options.advertise_in_cluster.multi_ports.ports](resources--workload--reference--group-009.md#canonical-0302230202101333-3033110013223110-0210311023223222-0203030332102133-2201333301103120-1000302302100130-2222303013302022-1022120223133312)
- [service.advertise_options.advertise_in_cluster.multi_ports.ports.info](resources--workload--reference--group-009.md#canonical-1032330310201200-3101102022121101-0202311121001130-1310203222000011-1113321122130331-0113223320013030-1012012320323211-2111201022201331)
- service.advertise_options.advertise_in_cluster.multi_ports.ports.info.same_as_port

<a id="canonical-1100100120332232-2031310231312130-3033131012331023-2331230011302021-3233033232202301-0010313031313132-0333012102302210-3033111123331223"></a>

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

<a id="canonical-2032101121303011-1103222233000330-0203232311001010-2322120323102000-3202033332221123-2310113223210032-0131222201221033-2112013321213130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_in_cluster.port` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_in_cluster](resources--workload--reference--group-009.md#canonical-0330023301131021-0130211002331300-3002122032320301-2112231133101233-3303101211030122-3300030030023300-1213020320030020-0213001312211213)
- service.advertise_options.advertise_in_cluster.port

<a id="canonical-2101101323122211-1123110212023233-2013223033212303-1110022300332000-3202333301102030-1131030212001331-2311013113311312-3003012311022321"></a>

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

<a id="canonical-3203221012003020-3303032232023210-1103311231233011-2201302023210110-3320032110122312-0110020010011320-2203303121312013-2012002300030103"></a>

### Direct properties for `service.advertise_options.advertise_in_cluster.port`

- [info](resources--workload--reference--group-009.md#canonical-3000021121133300-0112013033122013-0332121130331320-2320022103131303-3211132020321221-0102021013132121-3121210012321313-1032112131200133): complete subsection reference.

<a id="canonical-3000021121133300-0112013033122013-0332121130331320-2320022103131303-3211132020321221-0102021013132121-3121210012321313-1032112131200133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_in_cluster.port.info` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_in_cluster](resources--workload--reference--group-009.md#canonical-0330023301131021-0130211002331300-3002122032320301-2112231133101233-3303101211030122-3300030030023300-1213020320030020-0213001312211213)
- [service.advertise_options.advertise_in_cluster.port](resources--workload--reference--group-009.md#canonical-2032101121303011-1103222233000330-0203232311001010-2322120323102000-3202033332221123-2310113223210032-0131222201221033-2112013321213130)
- service.advertise_options.advertise_in_cluster.port.info

<a id="canonical-3102321302002202-1313330301020233-3220111220002100-1000101110111200-2210122331010211-2330233022003102-3221333221311132-3122313223223010"></a>

Type: `"object"`. single nested block, Optional.

Port Information. Port information.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("port"),
  validators.ConflictingObjectAttributes("same_as_port",
    "target_port")}
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
  "x-ves-oneof-field-target_port_choice": "[\"same_as_port\",\"target_port\"]"
}
```

Terraform syntax:

```terraform
info {
  # Configure direct properties listed below.
}
```

<a id="canonical-3232012010303303-1212111321313200-1131303132331232-3121122131202110-3101232121231211-1301222130220132-0322102111311210-2321100101031010"></a>

### Direct properties for `service.advertise_options.advertise_in_cluster.port.info`

<a id="canonical-0203013130130121-2211213131321210-0110230102033000-3202003111200310-3231033231311301-2201220120101032-2211032220132101-2133230030021321"></a>

#### `service.advertise_options.advertise_in_cluster.port.info.port` property

Type: `"number"`. Optional.

Port. Port the workload can be reached on.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-2303100213300302-0020323211102102-0022222100230221-3213010211333131-0122201230303120-2003100210023103-1000330010313301-3303333003122111"></a>

<a id="canonical-3211101122322133-1301203032331312-0201203011220322-0331300221020323-2033302021023123-1102203102201031-0311302013023203-1032001022031212"></a>

#### `service.advertise_options.advertise_in_cluster.port.info.protocol` property

Type: `"string"`. Optional.

\[Enum: PROTOCOL\_TCP|PROTOCOL\_HTTP|PROTOCOL\_HTTP2|PROTOCOL\_TLS\_WITH\_SNI|PROTOCOL\_UDP\] Type
of protocol - PROTOCOL\_TCP: TCP TCP - PROTOCOL\_HTTP: HTTP HTTP - PROTOCOL\_HTTP2: HTTP2 HTTP2 -
PROTOCOL\_TLS\_WITH\_SNI: TLS with SNI TLS with SNI - PROTOCOL\_UDP: UDP UDP. Possible values are
\`PROTOCOL\_TCP\`, \`PROTOCOL\_HTTP\`, \`PROTOCOL\_HTTP2\`, \`PROTOCOL\_TLS\_WITH\_SNI\`,
\`PROTOCOL\_UDP\`. Defaults to \`PROTOCOL\_TCP\`.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("PROTOCOL_TCP",
    "PROTOCOL_HTTP",
    "PROTOCOL_HTTP2",
    "PROTOCOL_TLS_WITH_SNI",
    "PROTOCOL_UDP"),
}
```

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

- [same_as_port](resources--workload--reference--group-009.md#canonical-0331030002203301-2000202112302103-1121310003132201-1133302201230133-3333102123022321-3023221303113110-3202332010212000-1012220001110232): complete subsection reference.

<a id="canonical-2313132133230101-0113333300123012-3201210202101322-2230231221112011-0202100311201003-0233323033131220-2010111222221211-1120311323001122"></a>

<a id="canonical-1010300213103323-0233102322203322-0211313221110120-3021211003003301-0102233000313232-1213033220201302-2000103132121122-2102211021321231"></a>

#### `service.advertise_options.advertise_in_cluster.port.info.target_port` property

Type: `"number"`. Optional.

Exclusive with \[same\_as\_port\] Port the workload is listening on.

Provider validators and defaults (from schema source):

```go
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
    "ves.io.schema.rules.uint32.gt": "0",
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gt": "0",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-0331030002203301-2000202112302103-1121310003132201-1133302201230133-3333102123022321-3023221303113110-3202332010212000-1012220001110232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_in_cluster.port.info.same_as_port` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_in_cluster](resources--workload--reference--group-009.md#canonical-0330023301131021-0130211002331300-3002122032320301-2112231133101233-3303101211030122-3300030030023300-1213020320030020-0213001312211213)
- [service.advertise_options.advertise_in_cluster.port](resources--workload--reference--group-009.md#canonical-2032101121303011-1103222233000330-0203232311001010-2322120323102000-3202033332221123-2310113223210032-0131222201221033-2112013321213130)
- [service.advertise_options.advertise_in_cluster.port.info](resources--workload--reference--group-009.md#canonical-3000021121133300-0112013033122013-0332121130331320-2320022103131303-3211132020321221-0102021013132121-3121210012321313-1032112131200133)
- service.advertise_options.advertise_in_cluster.port.info.same_as_port

<a id="canonical-2310130320320001-0321130312020100-1102312221011200-0331132113122323-0220233323202133-2231321002231203-2232231322102110-1233331301201022"></a>

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

<a id="canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- service.advertise_options.advertise_on_public

<a id="canonical-3101011112220000-1112230302331111-0310010332131220-0300200300310022-0030222321031212-1120302233113133-0133032102333001-0032321033213020"></a>

Type: `"object"`. single nested block, Optional.

Advertise this workload via loadbalancer on internet with default VIP.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("multi_ports",
    "port")}
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
  "x-ves-oneof-field-advertise_choice": "[\"multi_ports\",\"port\"]"
}
```

Terraform syntax:

```terraform
advertise_on_public {
  # Configure direct properties listed below.
}
```

<a id="canonical-3232101100321100-3112220210213311-2300313113013112-3230031232321011-3312023000200031-2312310003001312-1030213202230022-1203200031222110"></a>

### Direct properties for `service.advertise_options.advertise_on_public`

- [multi_ports](resources--workload--reference--group-009.md#canonical-3210102113031103-2332201101220100-1103122221002212-3002101012001000-1332031130002011-0021231133303122-2121222212212213-1123320220010203): complete subsection reference.

- [port](resources--workload--reference--group-012.md#canonical-3013031222022030-2021222233122213-2220022232301003-3102220000332302-0311213311310132-3133321213021123-0133300222212330-1122131003023301): complete subsection reference.

<a id="canonical-3210102113031103-2332201101220100-1103122221002212-3002101012001000-1332031130002011-0021231133303122-2121222212212213-1123320220010203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.multi_ports` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-009.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- service.advertise_options.advertise_on_public.multi_ports

<a id="canonical-1030002101120112-1320101011031223-2332221312121111-3230202222201310-1320001001201120-0212221333123301-0013210013013102-2130112130310310"></a>

Type: `"object"`. single nested block, Optional.

Advertise Multiple Ports. Advertise multiple ports.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("ports")}
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
multi_ports {
  # Configure direct properties listed below.
}
```

<a id="canonical-2223010133032312-2020301302212331-2031232212223010-3000200230011323-3303120020320330-3003300023223302-2011013032101022-2310020121221012"></a>

### Direct properties for `service.advertise_options.advertise_on_public.multi_ports`

- [ports](resources--workload--reference--group-009.md#canonical-2311132233010120-0223230230122210-0332200203312210-3000132111100333-1133101301233010-3211310033122333-2323002110210000-1300012002110122): complete subsection reference.

<a id="canonical-2311132233010120-0223230230122210-0332200203312210-3000132111100333-1133101301233010-3211310033122333-2323002110210000-1300012002110122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.multi_ports.ports` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-009.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-009.md#canonical-3210102113031103-2332201101220100-1103122221002212-3002101012001000-1332031130002011-0021231133303122-2121222212212213-1123320220010203)
- service.advertise_options.advertise_on_public.multi_ports.ports

<a id="canonical-1212022013113012-1020030022221212-3120222021020231-2303332232323312-0301331020110111-2000302130200213-3322222032321023-1302033110100212"></a>

Type: `"object"`. list nested block, Optional.

Ports. Ports to advertise.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("http_loadbalancer",
    "tcp_loadbalancer")}
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
ports {
  # Configure direct properties listed below.
}
```

<a id="canonical-0031032302032013-0102131002121322-0101100201033203-3312221312032212-2010120011031222-2022111120302331-3213013002033322-1130001331322301"></a>

### Direct properties for `service.advertise_options.advertise_on_public.multi_ports.ports`

- [http_loadbalancer](resources--workload--reference--group-009.md#canonical-1202020123020333-2332300320113330-1321112113331202-2112220001121101-3022001013310321-0020100001212100-0032232320012323-1311122212310031): complete subsection reference.

- [port](resources--workload--reference--group-012.md#canonical-0331331222102030-2203010231311222-1103320331132010-0113201101003101-3301201000333230-1221211320313131-1102020121310021-1030210213331030): complete subsection reference.

- [tcp_loadbalancer](resources--workload--reference--group-012.md#canonical-1201000230031032-2312033000023331-1111332222123100-2001203111302231-0133102032000130-2121122320101311-2223311020113123-1220222002313220): complete subsection reference.

<a id="canonical-1202020123020333-2332300320113330-1321112113331202-2112220001121101-3022001013310321-0020100001212100-0032232320012323-1311122212310031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-009.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-009.md#canonical-3210102113031103-2332201101220100-1103122221002212-3002101012001000-1332031130002011-0021231133303122-2121222212212213-1123320220010203)
- [service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-009.md#canonical-2311132233010120-0223230230122210-0332200203312210-3000132111100333-1133101301233010-3211310033122333-2323002110210000-1300012002110122)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer

<a id="canonical-0212010320303113-1320002303012030-3100331023002103-2332010133112122-0022210313302310-1220022300232120-2010322322321312-3321002321021320"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for http loadbalancer.

Additional upstream details:

HTTP/HTTPS Load balancer.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("domains"),
  validators.ConflictingObjectAttributes("default_route",
    "specific_routes"),
  validators.ConflictingObjectAttributes("http",
    "https"),
  validators.ConflictingObjectAttributes("http",
    "https_auto_cert"),
  validators.ConflictingObjectAttributes("https",
    "https_auto_cert")}
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

<a id="canonical-2200331202212111-0230313203121013-2313211220233030-1303113332213013-1033311013133232-3100033303311300-1003020213330332-0121332213010000"></a>

### Direct properties for `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer`

- [default_route](resources--workload--reference--group-009.md#canonical-0232013311111022-0003320023211122-1321321011010330-1223120121220333-0213130300023300-2001123101323323-3002132100221212-2311012111330101): complete subsection reference.

<a id="canonical-2032303100131001-3230001322010003-3003202000122332-3011000120123020-2101013013103112-0132003102012100-3331201111300130-1220333313331030"></a>

<a id="canonical-2333303023130320-3320030321000323-1202033112210310-2201320303010331-1330301300212321-3120010101013112-1011032130131322-0233101013331211"></a>

#### `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.domains` property

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

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeBetween(1, 32),
}
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

- [http](resources--workload--reference--group-009.md#canonical-2032123312300020-0313233113231322-1312210332031202-2133210330020310-0110031022301023-1212321320020231-2033231203003230-2310210201131220): complete subsection reference.

- [https](resources--workload--reference--group-009.md#canonical-3212310132100213-2000301220120030-0112323230310220-3200032133332210-3223310210320320-2331013310222010-2010102133130322-2111123310321223): complete subsection reference.

- [https_auto_cert](resources--workload--reference--group-011.md#canonical-0123203011200310-2323132333110331-0320211012002122-2323013220112302-2000030132131023-1312320130303213-2223130103233300-3130200202200302): complete subsection reference.

- [specific_routes](resources--workload--reference--group-011.md#canonical-3333233100013331-2322011200101232-2320001330000031-2000212200300310-3210113330031012-2012103222021320-0322112233332312-3331120230102200): complete subsection reference.

<a id="canonical-0232013311111022-0003320023211122-1321321011010330-1223120121220333-0213130300023300-2001123101323323-3002132100221212-2311012111330101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.default_route` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-009.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-009.md#canonical-3210102113031103-2332201101220100-1103122221002212-3002101012001000-1332031130002011-0021231133303122-2121222212212213-1123320220010203)
- [service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-009.md#canonical-2311132233010120-0223230230122210-0332200203312210-3000132111100333-1133101301233010-3211310033122333-2323002110210000-1300012002110122)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](resources--workload--reference--group-009.md#canonical-1202020123020333-2332300320113330-1321112113331202-2112220001121101-3022001013310321-0020100001212100-0032232320012323-1311122212310031)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.default_route

<a id="canonical-0032213203000102-1303101223331031-3301022322131211-0111132120030201-2201330132133323-3013023020210223-3102103202021013-1302312200212130"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for default route.

Additional upstream details:

Default route matching all APIs.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("auto_host_rewrite",
    "disable_host_rewrite"),
  validators.ConflictingObjectAttributes("auto_host_rewrite",
    "host_rewrite"),
  validators.ConflictingObjectAttributes("disable_host_rewrite",
    "host_rewrite")}
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
  "x-ves-oneof-field-host_rewrite_params": "[\"auto_host_rewrite\",\"disable_host_rewrite\",\"host_rewrite\"]"
}
```

Terraform syntax:

```terraform
default_route {
  # Configure direct properties listed below.
}
```

<a id="canonical-1013322212100213-1012022103231212-3012110111323301-2002032133133031-2011310302320220-3102202010323302-3022231120101111-1002002132331203"></a>

### Direct properties for `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.default_route`

- [auto_host_rewrite](resources--workload--reference--group-009.md#canonical-2322120332220030-0300010303222132-0231233230310120-3312130212030203-3323022301121303-1300013301332323-1310232112322310-3321122202110221): complete subsection reference.

- [disable_host_rewrite](resources--workload--reference--group-009.md#canonical-1211002103313333-1231121300012003-1312223130210212-0002230330122201-0012213033030211-0020320230121233-2113131113313320-3110331123031333): complete subsection reference.

<a id="canonical-1333232102233301-2200021232321101-1212202010312032-0001211311320112-3010213230112123-0010130011022311-1303013101323111-3230333320313332"></a>

<a id="canonical-3201323120202231-2333133312312301-0030113110002131-0000011323313230-3100210131221011-2023000122233320-1331212320113022-3130132031111121"></a>

#### `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.default_route.host_rewrite` property

Type: `"string"`. Optional.

Exclusive with \[auto\_host\_rewrite disable\_host\_rewrite\] Host header will be swapped with this
value.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
}
```

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

<a id="canonical-2322120332220030-0300010303222132-0231233230310120-3312130212030203-3323022301121303-1300013301332323-1310232112322310-3321122202110221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.default_route.auto_host_rewrite` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-009.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-009.md#canonical-3210102113031103-2332201101220100-1103122221002212-3002101012001000-1332031130002011-0021231133303122-2121222212212213-1123320220010203)
- [service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-009.md#canonical-2311132233010120-0223230230122210-0332200203312210-3000132111100333-1133101301233010-3211310033122333-2323002110210000-1300012002110122)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](resources--workload--reference--group-009.md#canonical-1202020123020333-2332300320113330-1321112113331202-2112220001121101-3022001013310321-0020100001212100-0032232320012323-1311122212310031)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.default_route](resources--workload--reference--group-009.md#canonical-0232013311111022-0003320023211122-1321321011010330-1223120121220333-0213130300023300-2001123101323323-3002132100221212-2311012111330101)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.default_route.auto_host_rewrite

<a id="canonical-2021331302113012-2002312120020201-1201102220322311-3332112321330033-2032311123000131-3103221031211122-2111223213330233-0310230203113012"></a>

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

<a id="canonical-1211002103313333-1231121300012003-1312223130210212-0002230330122201-0012213033030211-0020320230121233-2113131113313320-3110331123031333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.default_route.disable_host_rewrite` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-009.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-009.md#canonical-3210102113031103-2332201101220100-1103122221002212-3002101012001000-1332031130002011-0021231133303122-2121222212212213-1123320220010203)
- [service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-009.md#canonical-2311132233010120-0223230230122210-0332200203312210-3000132111100333-1133101301233010-3211310033122333-2323002110210000-1300012002110122)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](resources--workload--reference--group-009.md#canonical-1202020123020333-2332300320113330-1321112113331202-2112220001121101-3022001013310321-0020100001212100-0032232320012323-1311122212310031)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.default_route](resources--workload--reference--group-009.md#canonical-0232013311111022-0003320023211122-1321321011010330-1223120121220333-0213130300023300-2001123101323323-3002132100221212-2311012111330101)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.default_route.disable_host_rewrite

<a id="canonical-2212013233013002-1312320020100322-3302231222322102-0110111123312230-3100011031213210-0312300200200023-2001211011112031-2322002010130110"></a>

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

<a id="canonical-2032123312300020-0313233113231322-1312210332031202-2133210330020310-0110031022301023-1212321320020231-2033231203003230-2310210201131220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.http` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-009.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-009.md#canonical-3210102113031103-2332201101220100-1103122221002212-3002101012001000-1332031130002011-0021231133303122-2121222212212213-1123320220010203)
- [service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-009.md#canonical-2311132233010120-0223230230122210-0332200203312210-3000132111100333-1133101301233010-3211310033122333-2323002110210000-1300012002110122)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](resources--workload--reference--group-009.md#canonical-1202020123020333-2332300320113330-1321112113331202-2112220001121101-3022001013310321-0020100001212100-0032232320012323-1311122212310031)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.http

<a id="canonical-2332011233212301-1201303130312203-2102323333231203-0321211302131113-0023013221221131-0220013031113301-0332003232032130-2102031322312301"></a>

Type: `"object"`. single nested block, Optional.

HTTP Choice. Choice for selecting HTTP proxy.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("port",
    "port_ranges")}
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
  "x-ves-oneof-field-port_choice": "[\"port\",\"port_ranges\"]"
}
```

Terraform syntax:

```terraform
http {
  # Configure direct properties listed below.
}
```

<a id="canonical-3312112202231101-2220333012310323-3011311002330012-1030000321200021-3230021103122212-2133320311303231-2103113333321120-2013102332120022"></a>

### Direct properties for `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.http`

<a id="canonical-1223312120322332-2111100302101223-3230113112302323-3203010232202120-2030321301012203-1133302133122223-0202313321003212-1200021000203000"></a>

#### `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.http.dns_volterra_managed` property

Type: `"bool"`. Optional.

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

<a id="canonical-1131223103331013-3031300201323330-2330032102302033-3322012101133320-2321021330220033-0100013222300132-0210313232132210-1012102033112323"></a>

<a id="canonical-3322212020322231-1032020312313301-3032201323220012-3302223323123233-2310022110110332-1110020013222312-1230021222323331-0201102013233323"></a>

#### `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.http.port` property

Type: `"number"`. Optional.

Exclusive with \[port\_ranges\] HTTP port to Listen.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-2003312200311022-1010100231102332-0320022020321302-0020023001130110-0322023012202103-2211132033321131-2130310322303022-1110332121232311"></a>

<a id="canonical-0331110211203033-0210332113310321-3021010133003201-3300330033322210-1211302000032133-0002321000023033-3210212130211321-2032120012320320"></a>

#### `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.http.port_ranges` property

Type: `"string"`. Optional.

Exclusive with \[port\] A string containing a comma separated list of port ranges. Each port range
consists of a single port or two ports separated by '-'.

Additional upstream details:

Each port range consists of a single port or two ports separated by "-".

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 512),
}
```

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

<a id="canonical-3212310132100213-2000301220120030-0112323230310220-3200032133332210-3223310210320320-2331013310222010-2010102133130322-2111123310321223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-009.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-009.md#canonical-3210102113031103-2332201101220100-1103122221002212-3002101012001000-1332031130002011-0021231133303122-2121222212212213-1123320220010203)
- [service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-009.md#canonical-2311132233010120-0223230230122210-0332200203312210-3000132111100333-1133101301233010-3211310033122333-2323002110210000-1300012002110122)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](resources--workload--reference--group-009.md#canonical-1202020123020333-2332300320113330-1321112113331202-2112220001121101-3022001013310321-0020100001212100-0032232320012323-1311122212310031)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https

<a id="canonical-3130003112033033-3220213301212302-1322310013120233-2201132331312332-0223301310331120-0233313230022200-3203201002323001-2131102210122102"></a>

Type: `"object"`. single nested block, Optional.

Choice for selecting HTTP proxy with bring your own certificates.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("append_server_name",
    "default_header"),
  validators.ConflictingObjectAttributes("append_server_name",
    "pass_through"),
  validators.ConflictingObjectAttributes("append_server_name",
    "server_name"),
  validators.ConflictingObjectAttributes("default_header",
    "pass_through"),
  validators.ConflictingObjectAttributes("default_header",
    "server_name"),
  validators.ConflictingObjectAttributes("default_loadbalancer",
    "non_default_loadbalancer"),
  validators.ConflictingObjectAttributes("disable_path_normalize",
    "enable_path_normalize"),
  validators.ConflictingObjectAttributes("pass_through",
    "server_name"),
  validators.ConflictingObjectAttributes("port",
    "port_ranges"),
  validators.ConflictingObjectAttributes("tls_cert_params",
    "tls_parameters")}
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
  "x-ves-oneof-field-default_lb_choice": "[\"default_loadbalancer\",\"non_default_loadbalancer\"]",
  "x-ves-oneof-field-path_normalize_choice": "[\"disable_path_normalize\",\"enable_path_normalize\"]",
  "x-ves-oneof-field-port_choice": "[\"port\",\"port_ranges\"]",
  "x-ves-oneof-field-server_header_choice": "[\"append_server_name\",\"default_header\",\"pass_through\",\"server_name\"]",
  "x-ves-oneof-field-tls_certificates_choice": "[\"tls_cert_params\",\"tls_parameters\"]"
}
```

Terraform syntax:

```terraform
https {
  # Configure direct properties listed below.
}
```

<a id="canonical-1200021102301023-1200220023333111-3011331033323131-1322220311203021-0331111212111132-1001130212323221-3211023132322032-2200002030230031"></a>

### Direct properties for `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https`

<a id="canonical-2100221011330111-1032323112010021-3000201322301230-1222103123300302-1030301203020232-1303021310232132-3330130113033132-3232223310012012"></a>

#### `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.add_hsts` property

Type: `"bool"`. Optional.

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

<a id="canonical-3122313213122211-2012200202101300-3200121133112200-2311201312030301-0322012303131331-0003233322210012-2121103010012322-2000312111211123"></a>

<a id="canonical-3023031033023102-3313000131032031-2013111111011033-3131002133321002-2031030100333131-3133001300112130-2311033220011110-2030323103331200"></a>

#### `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.append_server_name` property

Type: `"string"`. Optional.

Exclusive with \[default\_header pass\_through server\_name\] Define the header value for the header
name “server”. If header value is already present, it is not overwritten and passed as-is.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(8096),
}
```

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
    "ves.io.schema.rules.string.max_len": "8096"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "8096"
  }
}
```

- [coalescing_options](resources--workload--reference--group-010.md#canonical-2300311201121333-1311112110203331-3310212300212011-2313230012130122-1332203113120320-3020302213011022-2331223222320311-0303223313003001): complete subsection reference.

<a id="canonical-2133123010310131-3103011110210332-3231202110332310-3210123230300111-3013212133112123-2000031202301230-2002213013121130-0331023133021323"></a>

<a id="canonical-1020133232212231-1001132210210312-2302013312321033-0203222132120302-0313223133332121-2233013101120030-0001220313212223-0003221122123133"></a>

#### `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.connection_idle_timeout` property

Type: `"number"`. Optional.

The idle timeout for downstream connections. The idle timeout is defined as the period in which
there are no active requests. When the idle timeout is reached the connection will be closed. Note
that request based timeouts mean that HTTP/2 PINGs will not keep the connection alive. This is
specified in milliseconds. The default value is 2 minutes.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.AtMost(600000),
}
```

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
    "ves.io.schema.rules.uint32.lte": "600000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "600000"
  }
}
```

- [default_header](resources--workload--reference--group-010.md#canonical-0001220200312001-1113113220200311-2121232002032023-3311312103002000-0121010101202330-2111213223210002-2321330322111220-2331333113001213): complete subsection reference.

- [default_loadbalancer](resources--workload--reference--group-010.md#canonical-0030131323133032-3230131331021321-0120331323111323-2311202313312123-2110223231212033-3320331321232223-0102313012203303-2311102313220320): complete subsection reference.

- [disable_path_normalize](resources--workload--reference--group-010.md#canonical-0100323020130132-1101221000003031-1021130013322030-0031130202003212-2332220321111301-0033022220301031-1033200212220030-3302131102131121): complete subsection reference.

- [enable_path_normalize](resources--workload--reference--group-010.md#canonical-1233033122010133-0313233003202220-3123310102130302-0030010200322301-1302221332333123-3300232123111230-0220102032030330-3101210201002321): complete subsection reference.

- [http_protocol_options](resources--workload--reference--group-010.md#canonical-2233110201213030-3333133212100020-2011121111212120-0113211322030320-0002310120033130-3322203130310011-3012120201131022-3120322301300011): complete subsection reference.

<a id="canonical-1312222321012332-0222033323033123-1213012320332303-0120100122120323-2302231202232322-1201102202332332-1013010130022332-2000111332110303"></a>

<a id="canonical-3321022203023313-2321022033311012-0101311110200322-3031130312332333-1130203132131000-1033102132102131-3103212211123321-1212103223012120"></a>

#### `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_redirect` property

Type: `"bool"`. Optional.

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

- [non_default_loadbalancer](resources--workload--reference--group-010.md#canonical-3033302113113321-1000032130332122-0133011213100323-3323120132112130-1222100231201310-3230202123223130-2103032130300101-1013121302101213): complete subsection reference.

- [pass_through](resources--workload--reference--group-010.md#canonical-0203331121231312-0110013320313123-1233303021130123-3331222120302133-1332231120213020-3203030031311331-2322322120023130-0002030203333232): complete subsection reference.

<a id="canonical-1022031321301103-3030302301310003-2010203122211331-0031323101330311-1212030030232202-1313220321200231-3032231202031301-0121200210022301"></a>

<a id="canonical-3011103101113012-2321130301110121-1201002010202112-1330312232201330-2002333222203331-2301321203020311-1010211322003233-0221121000022300"></a>

#### `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.port` property

Type: `"number"`. Optional.

Exclusive with \[port\_ranges\] HTTPS port to Listen.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-0320321213203201-0302112230213223-2110302210231121-3301023130033122-3013331121231203-2213112331210023-1310300113010311-2210210222310101"></a>

<a id="canonical-3313102331121203-1332131230230030-2021303110122031-1301331021233113-1300221310010120-3212300230033311-1120010222023020-1000110011133011"></a>

#### `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.port_ranges` property

Type: `"string"`. Optional.

Exclusive with \[port\] A string containing a comma separated list of port ranges. Each port range
consists of a single port or two ports separated by '-'.

Additional upstream details:

Each port range consists of a single port or two ports separated by "-".

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 512),
}
```

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

<a id="canonical-0001021113322132-3212310031122323-3111002200002300-2020310311331013-0303300031233100-2203132300101303-1031121301012221-1120132200020023"></a>

<a id="canonical-3033030112332011-1311022031211100-0333011322201332-3321202230002302-2122001232220120-3132232020233032-3213101313300003-0112222320133103"></a>

#### `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.server_name` property

Type: `"string"`. Optional.

Exclusive with \[append\_server\_name default\_header pass\_through\] Define the header value for
the header name “server”. This will overwrite existing values, if any, for the server header.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(8096),
}
```

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
    "ves.io.schema.rules.string.max_len": "8096"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "8096"
  }
}
```

- [tls_cert_params](resources--workload--reference--group-010.md#canonical-0212122113133333-1000102320023221-2003032220233320-0021231111201312-3001031012012332-3313221303003312-0112011332100001-1120030302022032): complete subsection reference.

- [tls_parameters](resources--workload--reference--group-010.md#canonical-3230132323111312-0132020110123203-0222302101120231-0202012321012303-1213310020220000-0101203211110131-1303320330302033-3001313221002112): complete subsection reference.
