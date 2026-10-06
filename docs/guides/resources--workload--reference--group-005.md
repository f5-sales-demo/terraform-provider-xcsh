---
page_title: "xcsh_workload reference"
subcategory: "Container"
description: "Complete grouped canonical reference for xcsh_workload reference."
---

# xcsh_workload reference

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
- [job.containers.readiness_check](resources--workload--reference--group-004.md#canonical-0222311231110223-3301213322101001-0122100312221013-1333033012100312-0133231033110100-3131310123021323-2322322213211230-3110310102123121)
- [job.containers.readiness_check.tcp_health_check](resources--workload--reference--group-004.md#canonical-1331303311323011-1231220102032100-0033122102322121-0323121321110233-3230223220133313-1001001130132012-1012102132021112-3000233323110222)
- job.containers.readiness_check.tcp_health_check.port

<a id="canonical-3012211012302320-0200232330230133-2211312000011131-2033211032112202-0301323211101121-0010133311313003-3002121122021022-1203123220000223"></a>

Type: `"object"`. single nested block, Optional.

Port. Port

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-3133220103013132-1322222200323320-3031113330130330-2123012223003012-2203012123131122-2301320031033331-3031003123330001-2130022121302010"></a>

### Direct properties for `job.containers.readiness_check.tcp_health_check.port`

<a id="canonical-1012110333010302-1030333122022303-1121020001100222-3200110012202023-0320102113131313-2312111313132013-2110222122311032-2302221211322320"></a>

#### `job.containers.readiness_check.tcp_health_check.port.name` property

Type: `"string"`. Optional.

Port Name. Exclusive with \[num\] Port Name.

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

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-0310001221123210-0312232003120101-0300203223322231-3003321202313212-2102232112232012-1030001113101333-1133322303112012-0301031133002000"></a>

### Direct properties for `job.deploy_options`

- [all_res](resources--workload--reference--group-005.md#canonical-3131221000303320-1002331010001331-0332121303112322-3121123332120010-1003200023102022-2210023012213201-3232021120032200-2300331010112300): complete subsection reference.

- [default_virtual_sites](resources--workload--reference--group-005.md#canonical-2123101231213202-0333113001323311-1102130211000000-0120320311011312-1121031322123201-1132210012200231-1033021121101112-2312133113213303): complete subsection reference.

- [deploy_ce_sites](resources--workload--reference--group-005.md#canonical-3011331012111131-1200211221323333-2222201313111211-2132322200111103-1210332013232031-2303233120011013-2233031312000221-0310210001111123): complete subsection reference.

- [deploy_ce_virtual_sites](resources--workload--reference--group-005.md#canonical-3123031132321222-1211312320321323-2100330023231222-1011202220031232-2332313222331100-1333220330323313-0212113310011030-0000102112201113): complete subsection reference.

- [deploy_re_sites](resources--workload--reference--group-005.md#canonical-3202012002300303-1220012332222101-0310120133000321-0300120210310013-2101111121220301-0321202013303020-1302001321033011-1133320020002000): complete subsection reference.

- [deploy_re_virtual_sites](resources--workload--reference--group-005.md#canonical-2303323000212010-1331320302002323-2131023001233320-2232113103000220-1031323101300203-3020202312102100-3332323210033031-2112133212011023): complete subsection reference.

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

Provider validators and defaults (from schema source):

```go
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

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-2102133001133310-2303302121232022-0021132323221231-1300020010311222-3003220310323102-2303231103100212-0202331010000323-2320103333011200"></a>

### Direct properties for `job.deploy_options.deploy_ce_sites.site`

<a id="canonical-3001230221103003-1131113103130113-2300031313011222-3200012121112230-1302132231311101-3130232221122100-3222111232332112-3201122222000030"></a>

#### `job.deploy_options.deploy_ce_sites.site.name` property

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-3011032112012302-0230221012313010-2213100003203220-2310122302323332-1212212220110231-1223212013301030-3013300310332230-2133332120210301"></a>

<a id="canonical-0310112230202002-0111022230002230-3222310110302032-2311101013221123-1120210331213213-2202302203110213-1002202011121120-2033200231333122"></a>

#### `job.deploy_options.deploy_ce_sites.site.namespace` property

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-3113310102231213-3203013123211331-3032300203223012-2221020321323013-2321303320310320-2022203033313233-2133330010223231-2013233222200032"></a>

<a id="canonical-3202112110313102-2012113010131013-0112010102231033-1030233223111311-0111133323032311-2101313130222230-1111102203321132-0001220202222213"></a>

#### `job.deploy_options.deploy_ce_sites.site.tenant` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Provider validators and defaults (from schema source):

```go
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

Provider validators and defaults (from schema source):

```go
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

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-0311231111313231-2332033303002000-3310133232220301-3221133112322223-3100132102211310-2030212201231122-1133021302002020-1033320112112021"></a>

### Direct properties for `job.deploy_options.deploy_ce_virtual_sites.virtual_site`

<a id="canonical-1013302323312320-3012330213330123-3220230231023331-3310110222023222-1002222310300032-2332112032202331-2302021201120131-1023330213111200"></a>

#### `job.deploy_options.deploy_ce_virtual_sites.virtual_site.name` property

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-3322113202131002-2331121322210100-3230132002302123-3323002322333212-2323022112021122-2202330020312220-0001032230112010-3033032230110110"></a>

<a id="canonical-0330113120012100-0333023010120023-3330031000320221-0331221233212102-3313222222300031-0000201020021100-0023221220211020-0100200000013313"></a>

#### `job.deploy_options.deploy_ce_virtual_sites.virtual_site.namespace` property

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-3020130112320032-3320022102002220-3331100123123130-0022031333032131-1111333133032110-0330313201130300-2100220322203330-1120002132021022"></a>

<a id="canonical-1031311233302310-2202011213132121-0300212102220110-0332130103210233-2200211013331333-0321221112203101-1303332312332112-1333322122230102"></a>

#### `job.deploy_options.deploy_ce_virtual_sites.virtual_site.tenant` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Provider validators and defaults (from schema source):

```go
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

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-2210120100120200-0022132123312011-0120012233333022-3101121023332230-0320313101003101-2122203132233100-3222033001331321-0302232333332232"></a>

### Direct properties for `job.deploy_options.deploy_re_sites`

- [site](resources--workload--reference--group-005.md#canonical-1013311331031231-3112002103021332-3330003313111330-3023232231020030-2033001221022131-3102223310202001-0323311203112313-2022123331132012): complete subsection reference.

<a id="canonical-1013311331031231-3112002103021332-3330003313111330-3023232231020030-2033001221022131-3102223310202001-0323311203112313-2022123331132012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `job.deploy_options.deploy_re_sites.site` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [job](resources--workload--reference--group-004.md#canonical-2101133223131122-0231303130011203-2130203200233301-2332212033103313-3210300031001001-1312100313213300-1013120312121230-0320112032220311)
- [job.deploy_options](resources--workload--reference--group-005.md#canonical-0302031202020321-2011011103202223-2203332002201122-3213130121212012-1321322223212202-2012010100321200-2301013113222301-2013123123001030)
- [job.deploy_options.deploy_re_sites](resources--workload--reference--group-005.md#canonical-3202012002300303-1220012332222101-0310120133000321-0300120210310013-2101111121220301-0321202013303020-1302001321033011-1133320020002000)
- job.deploy_options.deploy_re_sites.site

<a id="canonical-1211222110103303-2210010103230330-1220111300303033-0133323032221201-3030133123211331-1103001200012011-1010302133000200-2331132000132231"></a>

Type: `"object"`. list nested block, Optional.

Which regional edge sites should this workload be deployed.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-1120133332333002-3013013230132203-0010323233022203-3231231020122201-2303130202112322-1011012200001333-1020110002103003-1201002031330331"></a>

### Direct properties for `job.deploy_options.deploy_re_sites.site`

<a id="canonical-3331320310131200-0003321123212332-2132311322000002-1213122131322222-2201333110123200-1201220120323031-3313121321202310-3031211300221302"></a>

#### `job.deploy_options.deploy_re_sites.site.name` property

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-2331102000001121-0223312131313320-1211223200221132-1023322232032033-1133130131133203-2230211232123301-0220220001321300-1312110311030233"></a>

<a id="canonical-2101322221120111-2102210320202012-2313333133202133-2303132013301232-1020230002013333-2230221122130332-2300133002202021-1331202121200322"></a>

#### `job.deploy_options.deploy_re_sites.site.namespace` property

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-0202021202202100-2123103111203132-1022023013322001-2003100020021101-2113313231122032-3200012312212130-3223030120201021-2303300023222200"></a>

<a id="canonical-1231002103210031-3130102113213301-0203320002121313-1201013020221001-0200030023301112-2211131121200233-0213121111133300-1301013212013112"></a>

#### `job.deploy_options.deploy_re_sites.site.tenant` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-2303323000212010-1331320302002323-2131023001233320-2232113103000220-1031323101300203-3020202312102100-3332323210033031-2112133212011023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `job.deploy_options.deploy_re_virtual_sites` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [job](resources--workload--reference--group-004.md#canonical-2101133223131122-0231303130011203-2130203200233301-2332212033103313-3210300031001001-1312100313213300-1013120312121230-0320112032220311)
- [job.deploy_options](resources--workload--reference--group-005.md#canonical-0302031202020321-2011011103202223-2203332002201122-3213130121212012-1321322223212202-2012010100321200-2301013113222301-2013123123001030)
- job.deploy_options.deploy_re_virtual_sites

<a id="canonical-2002320031213220-1100110103132113-1023311212100122-0003032033230302-0102112102133003-0002301322230021-0211110312121120-2110110331331111"></a>

Type: `"object"`. single nested block, Optional.

This defines a way to deploy a workload on specific Regional Edge virtual sites.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-1000323323022131-0331321002213323-3032223031123233-2112101131223011-1330300133011233-2312022002221233-0301031132012112-0222112320123211"></a>

### Direct properties for `job.deploy_options.deploy_re_virtual_sites`

- [virtual_site](resources--workload--reference--group-005.md#canonical-2301021131031223-0310023300031121-3331011202300221-3000201100101203-0112032312211332-1110131112231110-3133330101302303-3321203122222130): complete subsection reference.

<a id="canonical-2301021131031223-0310023300031121-3331011202300221-3000201100101203-0112032312211332-1110131112231110-3133330101302303-3321203122222130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `job.deploy_options.deploy_re_virtual_sites.virtual_site` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [job](resources--workload--reference--group-004.md#canonical-2101133223131122-0231303130011203-2130203200233301-2332212033103313-3210300031001001-1312100313213300-1013120312121230-0320112032220311)
- [job.deploy_options](resources--workload--reference--group-005.md#canonical-0302031202020321-2011011103202223-2203332002201122-3213130121212012-1321322223212202-2012010100321200-2301013113222301-2013123123001030)
- [job.deploy_options.deploy_re_virtual_sites](resources--workload--reference--group-005.md#canonical-2303323000212010-1331320302002323-2131023001233320-2232113103000220-1031323101300203-3020202312102100-3332323210033031-2112133212011023)
- job.deploy_options.deploy_re_virtual_sites.virtual_site

<a id="canonical-3023132123133311-2330123110201220-2233020031021000-2111103100031122-3220233102333021-1103231000221313-3212321032112020-2020110201213103"></a>

Type: `"object"`. list nested block, Optional.

Which regional edge virtual sites should this workload be deployed.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-2302323022102010-2130033321220032-3223130221020103-0112300312120000-0123133220301320-1222013100211301-2102102301100132-1032123013312022"></a>

### Direct properties for `job.deploy_options.deploy_re_virtual_sites.virtual_site`

<a id="canonical-1012223100013321-0001223002020312-0333331131211030-2131332202322200-3023031022020201-3132301020130033-3000322103033330-2123213232221220"></a>

#### `job.deploy_options.deploy_re_virtual_sites.virtual_site.name` property

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-3211002011230232-0100030003131022-3313330333131033-3323213131110221-3300000121110113-0132131230120231-3130003000102031-0013012213203201"></a>

<a id="canonical-2221031113212323-1213032303130320-3031312311212221-0203121330220032-0120030321222211-0100130113030010-2212123133110321-0130111213223230"></a>

#### `job.deploy_options.deploy_re_virtual_sites.virtual_site.namespace` property

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-0300010100213102-2301222002231202-3302032233032023-1032303122102102-0111102220332223-2103323212302321-3012312200220212-2301021221200320"></a>

<a id="canonical-2003030113013302-3111321211130223-3002121201113233-3222231303330022-0203132232223032-3022132120201221-0002303311203303-3130210303133003"></a>

#### `job.deploy_options.deploy_re_virtual_sites.virtual_site.tenant` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-3030211032210202-3200123100133322-2330332311200201-3221211012330020-2120013021021113-1301223221032102-0030211333033310-1111121220320131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `job.volumes` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [job](resources--workload--reference--group-004.md#canonical-2101133223131122-0231303130011203-2130203200233301-2332212033103313-3210300031001001-1312100313213300-1013120312121230-0320112032220311)
- job.volumes

<a id="canonical-0231212110000201-1120003012103001-0303010011303320-3332020301232122-0110210333012032-3331213303300212-3122101033222331-3202202211330322"></a>

Type: `"object"`. list nested block, Optional.

Volumes. Volumes for the job.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("empty_dir",
    "host_path"),
  validators.ConflictingListObjectAttributes("empty_dir",
    "persistent_volume"),
  validators.ConflictingListObjectAttributes("host_path",
    "persistent_volume")}
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

<a id="canonical-0030010103301010-1012223302331233-1303133012203113-0302332201221031-0213233322320021-0213232332101132-3103020310320113-1222233132202210"></a>

### Direct properties for `job.volumes`

- [empty_dir](resources--workload--reference--group-005.md#canonical-2103002013003112-0200112230012330-1120110101321210-2130121011113010-1002233201310100-2121020032313303-2220023232010202-1102001023221221): complete subsection reference.

- [host_path](resources--workload--reference--group-005.md#canonical-1223231211123302-0302203312202000-2111232130221103-1230023203102122-1133130211023032-2231232202003100-0213320303000202-3033231213122010): complete subsection reference.

<a id="canonical-2123333023031333-2001132033133301-0003300112201003-3113232322220202-1320120313120031-1333032210230232-3023322210322110-3120300123222003"></a>

<a id="canonical-2030003313111333-2301300303000311-0332100012131302-3012133202332021-3013020131123030-2011123331131301-1033320310010302-1121223031113103"></a>

#### `job.volumes.name` property

Type: `"string"`. Optional.

Name. Name of the volume.

Provider validators and defaults (from schema source):

```go
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

- [persistent_volume](resources--workload--reference--group-005.md#canonical-3203200012213001-3111300330011222-2132220323122220-1123013113230202-1310102202333231-2032301213111230-0312131213202310-3223220232002311): complete subsection reference.

<a id="canonical-2103002013003112-0200112230012330-1120110101321210-2130121011113010-1002233201310100-2121020032313303-2220023232010202-1102001023221221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `job.volumes.empty_dir` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [job](resources--workload--reference--group-004.md#canonical-2101133223131122-0231303130011203-2130203200233301-2332212033103313-3210300031001001-1312100313213300-1013120312121230-0320112032220311)
- [job.volumes](resources--workload--reference--group-005.md#canonical-3030211032210202-3200123100133322-2330332311200201-3221211012330020-2120013021021113-1301223221032102-0030211333033310-1111121220320131)
- job.volumes.empty_dir

<a id="canonical-3121213232002230-0311310301111200-1231303321030010-3212103233313023-0302110310013000-3100333331330003-2001221130333313-3000130111120232"></a>

Type: `"object"`. single nested block, Optional.

Volume containing a temporary directory whose lifetime is the same as a replica of a workload.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-0032030212021223-1200121331132003-1112222321323301-3201212213131220-3123101100131222-2131313233110210-3110122200313100-0231023100321202"></a>

### Direct properties for `job.volumes.empty_dir`

- [mount](resources--workload--reference--group-005.md#canonical-3212003132033033-0010122212331320-2323223221123322-1110010301223210-0310102122003131-2213322012011102-1213022220300032-3232212312231111): complete subsection reference.

<a id="canonical-1132032233323012-3231233013031030-3212002010010111-0232131000310332-1213221231000131-2232102233030030-0101320002303211-0130111313311012"></a>

<a id="canonical-3101111331002111-2020030002211222-3001202111203301-3030302131311232-2231102200303130-1312032202131230-0100102202110200-3120103232302230"></a>

#### `job.volumes.empty_dir.size_limit` property

Type: `"number"`. Optional.

Size Limit (in GiB). Configuration parameter for size limit

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

<a id="canonical-3212003132033033-0010122212331320-2323223221123322-1110010301223210-0310102122003131-2213322012011102-1213022220300032-3232212312231111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `job.volumes.empty_dir.mount` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [job](resources--workload--reference--group-004.md#canonical-2101133223131122-0231303130011203-2130203200233301-2332212033103313-3210300031001001-1312100313213300-1013120312121230-0320112032220311)
- [job.volumes](resources--workload--reference--group-005.md#canonical-3030211032210202-3200123100133322-2330332311200201-3221211012330020-2120013021021113-1301223221032102-0030211333033310-1111121220320131)
- [job.volumes.empty_dir](resources--workload--reference--group-005.md#canonical-2103002013003112-0200112230012330-1120110101321210-2130121011113010-1002233201310100-2121020032313303-2220023232010202-1102001023221221)
- job.volumes.empty_dir.mount

<a id="canonical-1311032111003231-3031212031201332-1021133201032121-3213101332031321-0331123002002032-1221113122133100-3121122012022132-0022002200120032"></a>

Type: `"object"`. single nested block, Optional.

Volume mount describes how volume is mounted inside a workload.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-0013002112133221-3001200231210232-1100213101221230-3113312011213223-0130313013010233-2301201002030300-3110033012003103-0000111032102202"></a>

### Direct properties for `job.volumes.empty_dir.mount`

<a id="canonical-0230131103021003-1212322033311020-3110001212020102-3120202020022203-2313221110232321-1222203213212223-3013110231110010-2122031001032033"></a>

#### `job.volumes.empty_dir.mount.mode` property

Type: `"string"`. Optional.

\[Enum: VOLUME\_MOUNT\_READ\_ONLY|VOLUME\_MOUNT\_READ\_WRITE\] Mode in which the volume should be
mounted to the workload - VOLUME\_MOUNT\_READ\_ONLY: ReadOnly Mount the volume in read-only mode -
VOLUME\_MOUNT\_READ\_WRITE: Read Write Mount the volume in read-write mode. Possible values are
\`VOLUME\_MOUNT\_READ\_ONLY\`, \`VOLUME\_MOUNT\_READ\_WRITE\`. Defaults to
\`VOLUME\_MOUNT\_READ\_ONLY\`.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-1331222123022331-0000332213331121-0312223103122132-0021120330123020-2200012212122123-3121101012022213-2012232123222132-2232201322332223"></a>

<a id="canonical-1033321223212012-3023031110210311-0321032230312112-1010110003101122-2332020031223210-1200002303321223-3213310311330213-0032130110201032"></a>

#### `job.volumes.empty_dir.mount.mount_path` property

Type: `"string"`. Optional.

Path within the workload container at which the volume should be mounted. Must not contain ':'.

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

<a id="canonical-3011133121122231-3332231032300023-2210102102032122-1301130233300002-3030033132220201-2131002020120311-2021110311102001-1323201132233322"></a>

<a id="canonical-0311122332313301-3223303120100113-3002212233111231-2210120132113031-1322301122210000-1001310312021122-2020300033233221-3322010222000202"></a>

#### `job.volumes.empty_dir.mount.sub_path` property

Type: `"string"`. Optional.

Path within the volume from which the workload's volume should be mounted. Defaults to '' (volume's
root).

Additional upstream details:

Defaults to "" (volume's root).

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
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-1223231211123302-0302203312202000-2111232130221103-1230023203102122-1133130211023032-2231232202003100-0213320303000202-3033231213122010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `job.volumes.host_path` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [job](resources--workload--reference--group-004.md#canonical-2101133223131122-0231303130011203-2130203200233301-2332212033103313-3210300031001001-1312100313213300-1013120312121230-0320112032220311)
- [job.volumes](resources--workload--reference--group-005.md#canonical-3030211032210202-3200123100133322-2330332311200201-3221211012330020-2120013021021113-1301223221032102-0030211333033310-1111121220320131)
- job.volumes.host_path

<a id="canonical-2310111133311202-2212321203311312-3031110223112231-1010131011222123-1131122331022323-2021333011110123-2030031230203023-1320021331032131"></a>

Type: `"object"`. single nested block, Optional.

Volume containing a host mapped path into the workload.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-1003011111131030-0203233201110122-2010322111122022-3231312011300322-1303321321323103-2032211200113303-1332031300211101-2011033020113133"></a>

### Direct properties for `job.volumes.host_path`

- [mount](resources--workload--reference--group-005.md#canonical-0120233102003331-3110220213311222-0201010013211013-0313133003021232-2201212223230113-0022201212002011-0130031223322301-2100303001311123): complete subsection reference.

<a id="canonical-3210322300203001-3332300110020211-3202302001301331-2020110321033103-1131203330110223-2232203233012010-0113101000232310-2101233203313032"></a>

<a id="canonical-1232331301222130-1022210323101222-1003023233212021-3131113110311032-1121123300022331-1111203323200032-2132102330202322-3231111321331100"></a>

#### `job.volumes.host_path.path` property

Type: `"string"`. Optional.

Path. Path of the directory on the host.

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

<a id="canonical-0120233102003331-3110220213311222-0201010013211013-0313133003021232-2201212223230113-0022201212002011-0130031223322301-2100303001311123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `job.volumes.host_path.mount` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [job](resources--workload--reference--group-004.md#canonical-2101133223131122-0231303130011203-2130203200233301-2332212033103313-3210300031001001-1312100313213300-1013120312121230-0320112032220311)
- [job.volumes](resources--workload--reference--group-005.md#canonical-3030211032210202-3200123100133322-2330332311200201-3221211012330020-2120013021021113-1301223221032102-0030211333033310-1111121220320131)
- [job.volumes.host_path](resources--workload--reference--group-005.md#canonical-1223231211123302-0302203312202000-2111232130221103-1230023203102122-1133130211023032-2231232202003100-0213320303000202-3033231213122010)
- job.volumes.host_path.mount

<a id="canonical-2102110233303021-0321200123311013-0212332300123002-1133113130302101-0010102211232311-2102210033201213-0201230130233111-3130333021100132"></a>

Type: `"object"`. single nested block, Optional.

Volume mount describes how volume is mounted inside a workload.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-0132100222023233-2203132311320210-0011333311021030-1010223103320003-2301223332333133-3313122001030320-1213332000233103-2022033021310000"></a>

### Direct properties for `job.volumes.host_path.mount`

<a id="canonical-3111210131333202-3001001333102111-1131002302030123-0112301030022030-3103233020133302-3310102010233331-1323002100302332-1031233212220300"></a>

#### `job.volumes.host_path.mount.mode` property

Type: `"string"`. Optional.

\[Enum: VOLUME\_MOUNT\_READ\_ONLY|VOLUME\_MOUNT\_READ\_WRITE\] Mode in which the volume should be
mounted to the workload - VOLUME\_MOUNT\_READ\_ONLY: ReadOnly Mount the volume in read-only mode -
VOLUME\_MOUNT\_READ\_WRITE: Read Write Mount the volume in read-write mode. Possible values are
\`VOLUME\_MOUNT\_READ\_ONLY\`, \`VOLUME\_MOUNT\_READ\_WRITE\`. Defaults to
\`VOLUME\_MOUNT\_READ\_ONLY\`.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-1303121021321231-3320011231302200-2303330211330032-2112122111100111-1122301331332032-1020121301220320-2323332321331110-3112331111230032"></a>

<a id="canonical-0211301321002112-2313203322213312-3132113300322201-1321131021222201-0000223031020233-1131003133000112-3011000322012000-0211013010232212"></a>

#### `job.volumes.host_path.mount.mount_path` property

Type: `"string"`. Optional.

Path within the workload container at which the volume should be mounted. Must not contain ':'.

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

<a id="canonical-1133332103022332-0310032012330101-1200331223110321-1203201232202023-1023302332111202-0331232120033332-0131221213002321-3111011120231200"></a>

<a id="canonical-3100203023032112-2201112130313122-0103332313202102-3303310223313020-0212221311122001-1102100123212211-1300201232033322-3133212020210220"></a>

#### `job.volumes.host_path.mount.sub_path` property

Type: `"string"`. Optional.

Path within the volume from which the workload's volume should be mounted. Defaults to '' (volume's
root).

Additional upstream details:

Defaults to "" (volume's root).

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
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-3203200012213001-3111300330011222-2132220323122220-1123013113230202-1310102202333231-2032301213111230-0312131213202310-3223220232002311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `job.volumes.persistent_volume` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [job](resources--workload--reference--group-004.md#canonical-2101133223131122-0231303130011203-2130203200233301-2332212033103313-3210300031001001-1312100313213300-1013120312121230-0320112032220311)
- [job.volumes](resources--workload--reference--group-005.md#canonical-3030211032210202-3200123100133322-2330332311200201-3221211012330020-2120013021021113-1301223221032102-0030211333033310-1111121220320131)
- job.volumes.persistent_volume

<a id="canonical-3111011201203203-3300310131112033-0232012332322131-0311203232211302-1202021322001233-2302130021231001-0021320232200200-2310113120300103"></a>

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

<a id="canonical-2330023211002221-1030313112033122-3320301301230002-3211320320333220-1000131222131002-3121212013321230-1133203131003012-0102330303331222"></a>

### Direct properties for `job.volumes.persistent_volume`

- [mount](resources--workload--reference--group-005.md#canonical-2023331022012002-2120231221033010-3011232212230232-3001013023321101-2102102031311231-0231022101220013-3212313133022021-0231112012320031): complete subsection reference.

- [storage](resources--workload--reference--group-005.md#canonical-0031100121220303-2023112232312211-2312022203121230-2022312101303030-2310210212123232-0230131333201103-3220011000221100-1223333313213203): complete subsection reference.

<a id="canonical-2023331022012002-2120231221033010-3011232212230232-3001013023321101-2102102031311231-0231022101220013-3212313133022021-0231112012320031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `job.volumes.persistent_volume.mount` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [job](resources--workload--reference--group-004.md#canonical-2101133223131122-0231303130011203-2130203200233301-2332212033103313-3210300031001001-1312100313213300-1013120312121230-0320112032220311)
- [job.volumes](resources--workload--reference--group-005.md#canonical-3030211032210202-3200123100133322-2330332311200201-3221211012330020-2120013021021113-1301223221032102-0030211333033310-1111121220320131)
- [job.volumes.persistent_volume](resources--workload--reference--group-005.md#canonical-3203200012213001-3111300330011222-2132220323122220-1123013113230202-1310102202333231-2032301213111230-0312131213202310-3223220232002311)
- job.volumes.persistent_volume.mount

<a id="canonical-3212121222330030-2330010331030301-2203101221223201-1013030101111321-2202022232230313-3132010231012332-2002110032302013-1122101123122002"></a>

Type: `"object"`. single nested block, Optional.

Volume mount describes how volume is mounted inside a workload.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-2023312233213203-3232202133312212-0021023012211001-1231212133310113-2210033323212102-2221301001001331-3322032031333330-1221031010223000"></a>

### Direct properties for `job.volumes.persistent_volume.mount`

<a id="canonical-0203321331200100-3211121023213301-3232222133210202-3212203322130333-1123232323101020-0012132300312101-1223023213002300-1203122211200311"></a>

#### `job.volumes.persistent_volume.mount.mode` property

Type: `"string"`. Optional.

\[Enum: VOLUME\_MOUNT\_READ\_ONLY|VOLUME\_MOUNT\_READ\_WRITE\] Mode in which the volume should be
mounted to the workload - VOLUME\_MOUNT\_READ\_ONLY: ReadOnly Mount the volume in read-only mode -
VOLUME\_MOUNT\_READ\_WRITE: Read Write Mount the volume in read-write mode. Possible values are
\`VOLUME\_MOUNT\_READ\_ONLY\`, \`VOLUME\_MOUNT\_READ\_WRITE\`. Defaults to
\`VOLUME\_MOUNT\_READ\_ONLY\`.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-1111220221131312-0120220311032313-1130132123320231-3001303320332211-0021201223220123-1123332202123221-2001111102033312-2323230131020221"></a>

<a id="canonical-3213132323200213-3203201002130102-0230023103303223-3320020121230320-3032201113011232-0222333131213303-0102210231012301-0132203113323123"></a>

#### `job.volumes.persistent_volume.mount.mount_path` property

Type: `"string"`. Optional.

Path within the workload container at which the volume should be mounted. Must not contain ':'.

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

<a id="canonical-0113300023002022-3031300333321011-3222133011113111-3003230313130323-1321200220330101-1132103333323123-2100131312323201-0002100100203101"></a>

<a id="canonical-1210232111013232-3023032201220120-2312321113011323-1323003102222213-3130102323112021-0123301311022310-3312101020030100-1301213022001131"></a>

#### `job.volumes.persistent_volume.mount.sub_path` property

Type: `"string"`. Optional.

Path within the volume from which the workload's volume should be mounted. Defaults to '' (volume's
root).

Additional upstream details:

Defaults to "" (volume's root).

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
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-0031100121220303-2023112232312211-2312022203121230-2022312101303030-2310210212123232-0230131333201103-3220011000221100-1223333313213203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `job.volumes.persistent_volume.storage` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [job](resources--workload--reference--group-004.md#canonical-2101133223131122-0231303130011203-2130203200233301-2332212033103313-3210300031001001-1312100313213300-1013120312121230-0320112032220311)
- [job.volumes](resources--workload--reference--group-005.md#canonical-3030211032210202-3200123100133322-2330332311200201-3221211012330020-2120013021021113-1301223221032102-0030211333033310-1111121220320131)
- [job.volumes.persistent_volume](resources--workload--reference--group-005.md#canonical-3203200012213001-3111300330011222-2132220323122220-1123013113230202-1310102202333231-2032301213111230-0312131213202310-3223220232002311)
- job.volumes.persistent_volume.storage

<a id="canonical-0233303313001032-2212132132321200-3313233231012323-3232002222110131-2000131033121112-1001002232133220-0302301111310012-1001023312122000"></a>

Type: `"object"`. single nested block, Optional.

Persistent storage configuration is used to configure Persistent Volume Claim (PVC).

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-2022111113301002-2100002232131223-0203000123232002-0331333211201203-3000121113331010-2222320131121232-0030230101323003-2302332203222323"></a>

### Direct properties for `job.volumes.persistent_volume.storage`

<a id="canonical-0310011032001330-3312010133030210-2301222333301122-0313132313210233-1211000223203231-1231011110312010-1000333030033032-1023221313212231"></a>

#### `job.volumes.persistent_volume.storage.access_mode` property

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

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-1010222220321123-2200101223203123-0223200210321020-2332303210122130-0020203323320011-0300331301201031-2102020121223323-2212032101011211"></a>

<a id="canonical-0212322110232332-3002121212223103-0111030001130030-1210320133312301-2223121320233133-3230111010321312-0322203012202022-0033210021303132"></a>

#### `job.volumes.persistent_volume.storage.class_name` property

Type: `"string"`. Optional.

Exclusive with \[default\] Use the specified class name.

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
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

- [default](resources--workload--reference--group-005.md#canonical-3212012101330020-1220300220102232-2230312110120333-2202331201031221-1033110101001300-1223031221320100-3203213122300311-3320003322322112): complete subsection reference.

<a id="canonical-2311003321313230-2322101013131003-0321211121010211-1212132100220103-2202030012312002-3313313002013123-1232321331212030-0131010011233313"></a>

<a id="canonical-0300233231302230-2212011131311002-1323100221012200-0213013032003221-2131203333001203-1103112133030122-2222200121223233-0300120213110132"></a>

#### `job.volumes.persistent_volume.storage.storage_size` property

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

<a id="canonical-3212012101330020-1220300220102232-2230312110120333-2202331201031221-1033110101001300-1223031221320100-3203213122300311-3320003322322112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `job.volumes.persistent_volume.storage.default` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [job](resources--workload--reference--group-004.md#canonical-2101133223131122-0231303130011203-2130203200233301-2332212033103313-3210300031001001-1312100313213300-1013120312121230-0320112032220311)
- [job.volumes](resources--workload--reference--group-005.md#canonical-3030211032210202-3200123100133322-2330332311200201-3221211012330020-2120013021021113-1301223221032102-0030211333033310-1111121220320131)
- [job.volumes.persistent_volume](resources--workload--reference--group-005.md#canonical-3203200012213001-3111300330011222-2132220323122220-1123013113230202-1310102202333231-2032301213111230-0312131213202310-3223220232002311)
- [job.volumes.persistent_volume.storage](resources--workload--reference--group-005.md#canonical-0031100121220303-2023112232312211-2312022203121230-2022312101303030-2310210212123232-0230131333201103-3220011000221100-1223333313213203)
- job.volumes.persistent_volume.storage.default

<a id="canonical-2303131013301203-3201113313311111-0321303231023110-3323102013201003-2012203013311222-0010002203202223-3033231203111103-0030030112131202"></a>

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

<a id="canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- service

<a id="canonical-1332310211103303-1212000223320202-2221213223003333-1121213113131100-0301313231102131-2310301013222100-0012222131003032-0011220233300032"></a>

Type: `"object"`. single nested block, Optional.

Service does not maintain per replica state, however it can be configured to use persistent storage
that is shared amongst all the replicas. Replicas of a service are fungible and do not have a stable
network identity or storage. Common examples of services are web servers, application servers,
traditional SQL databases, etc.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("containers"),
  validators.ConflictingObjectAttributes("num_replicas",
    "scale_to_zero")}
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
  "x-ves-oneof-field-scaling_choice": "[\"num_replicas\",\"scale_to_zero\"]"
}
```

Terraform syntax:

```terraform
service {
  # Configure direct properties listed below.
}
```

<a id="canonical-3330132120323330-3112121100333103-1110013030322003-2032131330100331-1000102112111231-3313012221121221-1101032021013233-3122320011022321"></a>

### Direct properties for `service`

- [advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130): complete subsection reference.

- [configuration](resources--workload--reference--group-015.md#canonical-0120300131213121-2030323030023013-0301331320133010-0123120231213121-1220222031210011-1001323001213020-0213022223133211-2110030001312333): complete subsection reference.

- [containers](resources--workload--reference--group-015.md#canonical-0333202010020223-0311032212110312-0322013032012300-3111121223203012-2222131221222013-3020133002033023-0203232021132003-0203222130000122): complete subsection reference.

- [deploy_options](resources--workload--reference--group-016.md#canonical-0011011211333310-1311031311001132-3121223032103322-0223110022300031-0201021032310133-3010100110301231-0220203002321211-2031113202323003): complete subsection reference.

<a id="canonical-0333212012011110-2130112003031213-0311010210012302-1221223333233322-0201111111323232-0030112320000202-1212020101231213-2123201101310111"></a>

<a id="canonical-2102122230223311-3232333012320330-2013133232301132-2221112313101222-2312320213301230-0112300120021212-1103221121120131-3301230101233202"></a>

#### `service.num_replicas` property

Type: `"number"`. Optional.

Exclusive with \[scale\_to\_zero\] Number of replicas of service to spawn per site.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 5),
}
```

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
    "ves.io.schema.rules.int32.gt": "0",
    "ves.io.schema.rules.int32.lte": "5"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.int32.gt": "0",
    "ves.io.schema.rules.int32.lte": "5"
  }
}
```

- [scale_to_zero](resources--workload--reference--group-016.md#canonical-2112203323032023-1022033211220013-2331111010311212-0200221020302323-2123121101110311-2211200012322311-2122132200232301-3022222221010211): complete subsection reference.

- [volumes](resources--workload--reference--group-016.md#canonical-3322223031121220-1122333131231013-0332331210233133-1311312333300020-2001331131131300-0232320011132231-2013021033301233-0003032133003003): complete subsection reference.

<a id="canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- service.advertise_options

<a id="canonical-1330113111333321-0202110223311103-1033021331320313-1330313231121010-1101130003002030-0202122233113330-2031032023320312-1202202212321031"></a>

Type: `"object"`. single nested block, Optional.

Advertise OPTIONS are used to configure how and where to advertise the workload using load
balancers.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("advertise_custom",
    "advertise_in_cluster"),
  validators.ConflictingObjectAttributes("advertise_custom",
    "advertise_on_public"),
  validators.ConflictingObjectAttributes("advertise_custom",
    "do_not_advertise"),
  validators.ConflictingObjectAttributes("advertise_in_cluster",
    "advertise_on_public"),
  validators.ConflictingObjectAttributes("advertise_in_cluster",
    "do_not_advertise"),
  validators.ConflictingObjectAttributes("advertise_on_public",
    "do_not_advertise")}
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
  "x-ves-oneof-field-advertise_choice": "[\"advertise_custom\",\"advertise_in_cluster\",\"advertise_on_public\",\"do_not_advertise\"]"
}
```

Terraform syntax:

```terraform
advertise_options {
  # Configure direct properties listed below.
}
```

<a id="canonical-2100013133221331-2301110112300101-2323131002220001-3330201102301312-2323313032133333-0032003103332332-3321220323033133-3311310100311021"></a>

### Direct properties for `service.advertise_options`

- [advertise_custom](resources--workload--reference--group-005.md#canonical-2121230312030330-3010113231323001-1233013032320322-1032203002210311-2022021002113232-1123312032112001-0012103213230102-1112311001110020): complete subsection reference.

- [advertise_in_cluster](resources--workload--reference--group-009.md#canonical-0330023301131021-0130211002331300-3002122032320301-2112231133101233-3303101211030122-3300030030023300-1213020320030020-0213001312211213): complete subsection reference.

- [advertise_on_public](resources--workload--reference--group-009.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101): complete subsection reference.

- [do_not_advertise](resources--workload--reference--group-015.md#canonical-0310302023323212-1013021312131132-1222231111312001-2121122023132203-1332202130321213-0133012220300033-3111032100113201-3031211330022333): complete subsection reference.

<a id="canonical-2121230312030330-3010113231323001-1233013032320322-1032203002210311-2022021002113232-1123312032112001-0012103213230102-1112311001110020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_custom` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- service.advertise_options.advertise_custom

<a id="canonical-2211032332323103-1333121332123121-3032103321322003-3331122020211113-1201023212110002-0323233232022301-2102101111213233-3113013033132003"></a>

Type: `"object"`. single nested block, Optional.

Advertise this workload via loadbalancer on specific sites.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("advertise_where",
    "ports")}
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
advertise_custom {
  # Configure direct properties listed below.
}
```

<a id="canonical-0132212123130021-1102321130220100-1231230002133302-3333201112300303-2110121332222111-0023032220210231-0121300103033130-2231203120223312"></a>

### Direct properties for `service.advertise_options.advertise_custom`

- [advertise_where](resources--workload--reference--group-005.md#canonical-0000302211310030-1312212333220303-2330202200321201-0311101022030310-1322222011232201-0120032203031033-3002313020202001-2002031210022033): complete subsection reference.

- [ports](resources--workload--reference--group-006.md#canonical-0133221101333210-0222103021001232-3212010032013103-3331120032012221-1012203003232110-1233102323001233-2233321121201101-0122313122210200): complete subsection reference.

<a id="canonical-0000302211310030-1312212333220303-2330202200321201-0311101022030310-1322222011232201-0120032203031033-3002313020202001-2002031210022033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_custom.advertise_where` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-2121230312030330-3010113231323001-1233013032320322-1032203002210311-2022021002113232-1123312032112001-0012103213230102-1112311001110020)
- service.advertise_options.advertise_custom.advertise_where

<a id="canonical-0002001030223131-0201101132010012-0320131300310320-2223311102012032-3111223311000333-1111223012201302-0211112201230303-2010200302312220"></a>

Type: `"object"`. list nested block, Optional.

Where should this load balancer be available.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("site",
    "virtual_site"),
  validators.ConflictingListObjectAttributes("site",
    "vk8s_service"),
  validators.ConflictingListObjectAttributes("virtual_site",
    "vk8s_service")}
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
advertise_where {
  # Configure direct properties listed below.
}
```

<a id="canonical-0121131022221230-2132033031121131-3221202123102010-0120133302330131-1030231122132033-2232310031023232-3133111003221220-1132123032311301"></a>

### Direct properties for `service.advertise_options.advertise_custom.advertise_where`

- [site](resources--workload--reference--group-005.md#canonical-0030333100303013-1333122033000310-2322131031010012-3320311021122031-0010002121033331-2002122020332203-1232011312312311-0000322030003001): complete subsection reference.

- [virtual_site](resources--workload--reference--group-005.md#canonical-0233210232101101-1111132120230132-3012030232123130-2031231011320101-3311201113011302-2113330313323032-0302331031323203-3101303230022131): complete subsection reference.

- [vk8s_service](resources--workload--reference--group-006.md#canonical-3203023030311011-1123233030130010-2121223300322103-0331031333102030-1003323232232011-2122300232331102-3312332311122002-1200323203122001): complete subsection reference.

<a id="canonical-0030333100303013-1333122033000310-2322131031010012-3320311021122031-0010002121033331-2002122020332203-1232011312312311-0000322030003001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_custom.advertise_where.site` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-2121230312030330-3010113231323001-1233013032320322-1032203002210311-2022021002113232-1123312032112001-0012103213230102-1112311001110020)
- [service.advertise_options.advertise_custom.advertise_where](resources--workload--reference--group-005.md#canonical-0000302211310030-1312212333220303-2330202200321201-0311101022030310-1322222011232201-0120032203031033-3002313020202001-2002031210022033)
- service.advertise_options.advertise_custom.advertise_where.site

<a id="canonical-0001312130223102-0011232213223010-1103222321330121-1331102022003330-2112310003313313-3333212302310121-1022021211311223-3012002311030212"></a>

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

<a id="canonical-1102022010222230-2012033031111020-0012021002333210-1112331232231022-0132000321310313-2023131032330102-0231111103113201-0131122313331332"></a>

### Direct properties for `service.advertise_options.advertise_custom.advertise_where.site`

<a id="canonical-1322022022303123-2223023033133120-0000011031321213-0332211112012000-0213010322021133-2132223201102132-3222132111113200-1202210120120010"></a>

#### `service.advertise_options.advertise_custom.advertise_where.site.ip` property

Type: `"string"`. Optional.

Use given IP address as VIP on the site.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
  validators.IPv4Validator(),
}
```

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
    "ves.io.schema.rules.string.ipv4": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv4": "true"
  }
}
```

<a id="canonical-0330200022012011-1210312131230032-0203103211113130-2101222330100002-1123102132011211-2321103102123232-0000223000101231-0002323012030232"></a>

<a id="canonical-3131330022221233-3223230111021121-2022222101022020-2311221220131031-3330311320100021-0311311201332200-3110003213322303-2203303010123223"></a>

#### `service.advertise_options.advertise_custom.advertise_where.site.network` property

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

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("SITE_NETWORK_INSIDE_AND_OUTSIDE",
    "SITE_NETWORK_INSIDE",
    "SITE_NETWORK_OUTSIDE",
    "SITE_NETWORK_SERVICE",
    "SITE_NETWORK_OUTSIDE_WITH_INTERNET_VIP",
    "SITE_NETWORK_INSIDE_AND_OUTSIDE_WITH_INTERNET_VIP",
    "SITE_NETWORK_IP_FABRIC"),
}
```

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

- [site](resources--workload--reference--group-005.md#canonical-3320022311110022-3103030003202322-1131012232020133-2320321110203332-0111013210321122-0322313101132331-2131233111312213-1120001332232122): complete subsection reference.

<a id="canonical-3320022311110022-3103030003202322-1131012232020133-2320321110203332-0111013210321122-0322313101132331-2131233111312213-1120001332232122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_custom.advertise_where.site.site` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-2121230312030330-3010113231323001-1233013032320322-1032203002210311-2022021002113232-1123312032112001-0012103213230102-1112311001110020)
- [service.advertise_options.advertise_custom.advertise_where](resources--workload--reference--group-005.md#canonical-0000302211310030-1312212333220303-2330202200321201-0311101022030310-1322222011232201-0120032203031033-3002313020202001-2002031210022033)
- [service.advertise_options.advertise_custom.advertise_where.site](resources--workload--reference--group-005.md#canonical-0030333100303013-1333122033000310-2322131031010012-3320311021122031-0010002121033331-2002122020332203-1232011312312311-0000322030003001)
- service.advertise_options.advertise_custom.advertise_where.site.site

<a id="canonical-2132121311133003-2132020123320132-3332203131112003-1131100121303200-0130032222331033-0100222032331213-0223032110201120-3321312123022323"></a>

Type: `"object"`. single nested block, Optional.

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

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
site {
  # Configure direct properties listed below.
}
```

<a id="canonical-2120213110202332-1030032021020233-1120130300001031-0120011322313000-2321231220333321-2021323301023323-2310311003100332-2331231023112110"></a>

### Direct properties for `service.advertise_options.advertise_custom.advertise_where.site.site`

<a id="canonical-3200323120100221-1013303103220103-1212213103010333-3211133232330113-3311223230210302-0102102230123132-3320023200001002-3103223120133023"></a>

#### `service.advertise_options.advertise_custom.advertise_where.site.site.name` property

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-0022220303010221-3313003230212012-2322222211012013-2300330211123220-0033303222322331-2131032113030022-0101102333023103-1313312020300213"></a>

<a id="canonical-1213022002010013-2332223020301110-0023010010322011-0020102130322123-3212000031300132-0310100131211213-3011231313321300-0221211012030332"></a>

#### `service.advertise_options.advertise_custom.advertise_where.site.site.namespace` property

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-3023233233311000-2101130031123223-1032311333203101-2112111033231330-3101303002200213-3103103120330010-3103310310032212-0003222122013030"></a>

<a id="canonical-3320102030232001-2000110230301030-3122200210331020-3003113110311023-1020202202310023-2321122000003012-3112233212211033-3231003023001232"></a>

#### `service.advertise_options.advertise_custom.advertise_where.site.site.tenant` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-0233210232101101-1111132120230132-3012030232123130-2031231011320101-3311201113011302-2113330313323032-0302331031323203-3101303230022131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_custom.advertise_where.virtual_site` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-2121230312030330-3010113231323001-1233013032320322-1032203002210311-2022021002113232-1123312032112001-0012103213230102-1112311001110020)
- [service.advertise_options.advertise_custom.advertise_where](resources--workload--reference--group-005.md#canonical-0000302211310030-1312212333220303-2330202200321201-0311101022030310-1322222011232201-0120032203031033-3002313020202001-2002031210022033)
- service.advertise_options.advertise_custom.advertise_where.virtual_site

<a id="canonical-3333313023011111-1100321202102100-2131232110322201-3201021102210023-2301210230301202-1023033022001233-2322030032002232-0333113131003301"></a>

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

<a id="canonical-2310223220102202-1110133102200010-1312001021210201-0210132012013033-3333012101222333-1113132312100022-3333103001030132-3113120003230231"></a>

### Direct properties for `service.advertise_options.advertise_custom.advertise_where.virtual_site`

<a id="canonical-3030021213122133-2211120333111013-1233323112022120-3301203032310133-1202333003200213-1231000002001330-2021001233210123-0120333013230003"></a>

#### `service.advertise_options.advertise_custom.advertise_where.virtual_site.network` property

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

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("SITE_NETWORK_INSIDE_AND_OUTSIDE",
    "SITE_NETWORK_INSIDE",
    "SITE_NETWORK_OUTSIDE",
    "SITE_NETWORK_SERVICE",
    "SITE_NETWORK_OUTSIDE_WITH_INTERNET_VIP",
    "SITE_NETWORK_INSIDE_AND_OUTSIDE_WITH_INTERNET_VIP",
    "SITE_NETWORK_IP_FABRIC"),
}
```

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

- [virtual_site](resources--workload--reference--group-006.md#canonical-3111211030302213-0110213100301232-1020233303203221-3032002013303100-2020003131020020-2123313002320022-0102003213012311-0112003323132321): complete subsection reference.
