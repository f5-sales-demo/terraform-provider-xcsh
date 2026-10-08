---
page_title: "xcsh_workload reference"
subcategory: "Container"
description: "Complete grouped canonical reference for xcsh_workload reference."
---

# xcsh_workload reference

<a id="canonical-3230310223020302-1210221202333333-1323022020110122-3211121301021031-0011123011310121-3120203212311313-2203133000210032-0121212320231310"></a>

## `stateful_service.advertise_options.advertise_custom.ports.port.info.protocol` property

Type: `"string"`. Optional.

\[Enum: PROTOCOL\_TCP|PROTOCOL\_HTTP|PROTOCOL\_HTTP2|PROTOCOL\_TLS\_WITH\_SNI|PROTOCOL\_UDP\] Type
of protocol - PROTOCOL\_TCP: TCP TCP - PROTOCOL\_HTTP: HTTP HTTP - PROTOCOL\_HTTP2: HTTP2 HTTP2 -
PROTOCOL\_TLS\_WITH\_SNI: TLS with SNI TLS with SNI - PROTOCOL\_UDP: UDP UDP. Possible values are
\`PROTOCOL\_TCP\`, \`PROTOCOL\_HTTP\`, \`PROTOCOL\_HTTP2\`, \`PROTOCOL\_TLS\_WITH\_SNI\`,
\`PROTOCOL\_UDP\`. Defaults to \`PROTOCOL\_TCP\`.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["PROTOCOL_HTTP","PROTOCOL_HTTP2","PROTOCOL_TCP","PROTOCOL_TLS_WITH_SNI","PROTOCOL_UDP"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
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

- [same_as_port](resources--workload--reference--group-021.md#canonical-0031000132002020-1121213003003220-2120303333130311-2230321130311233-2333101122223220-0231301013200300-3000223303103332-3020010003231110): complete subsection reference.

<a id="canonical-0310321302123300-2201130330332110-3220220220202101-2033313302222131-3021233113130033-3021211102020000-1202321300301312-0321232120201212"></a>

<a id="canonical-0100010020010313-1122132110201222-0222310203110212-0223122220223132-1200232312323032-3032122033212020-0203002323233002-2331201123221223"></a>

## `stateful_service.advertise_options.advertise_custom.ports.port.info.target_port` property

Type: `"number"`. Optional.

Exclusive with \[same\_as\_port\] Port the workload is listening on.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-0031000132002020-1121213003003220-2120303333130311-2230321130311233-2333101122223220-0231301013200300-3000223303103332-3020010003231110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_custom.ports.port.info.same_as_port` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-017.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_custom](resources--workload--reference--group-017.md#canonical-0203102001320123-1120130012130212-0111101220231333-3101331123203121-1110000311003232-1321231201202312-3200200021132112-0200303231330002)
- [stateful_service.advertise_options.advertise_custom.ports](resources--workload--reference--group-017.md#canonical-3111311322332303-1201030020021201-3133021233000313-2022333331131202-3221312032323033-3133120120113232-3330021120310331-2331303302032231)
- [stateful_service.advertise_options.advertise_custom.ports.port](resources--workload--reference--group-020.md#canonical-2103101222313132-0030333210220102-2031201232312001-1103110230303213-2123210111312322-2333302111312301-1031201133012330-3303120000023223)
- [stateful_service.advertise_options.advertise_custom.ports.port.info](resources--workload--reference--group-020.md#canonical-3313132111210032-0022013311111223-1003113033223223-1030113031032131-1313012113300322-2200110331232231-1313321330202023-1012032021102011)
- stateful_service.advertise_options.advertise_custom.ports.port.info.same_as_port

<a id="canonical-2121320023122122-2201001321212211-3021201020312112-2113203012230113-3213110311200113-1032322320210103-3112110220000122-2103011212020200"></a>

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

<a id="canonical-0020101203031230-1223300211330033-2123113022210223-3310333011133100-3330221301313310-3131021233011311-3221200100200012-0300221100330331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_custom.ports.tcp_loadbalancer` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-017.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_custom](resources--workload--reference--group-017.md#canonical-0203102001320123-1120130012130212-0111101220231333-3101331123203121-1110000311003232-1321231201202312-3200200021132112-0200303231330002)
- [stateful_service.advertise_options.advertise_custom.ports](resources--workload--reference--group-017.md#canonical-3111311322332303-1201030020021201-3133021233000313-2022333331131202-3221312032323033-3133120120113232-3330021120310331-2331303302032231)
- stateful_service.advertise_options.advertise_custom.ports.tcp_loadbalancer

<a id="canonical-2012302312002110-2112313121332331-3230203320003100-1003300002322100-0120233133322333-0322322031010300-2321102211011110-0020322202321021"></a>

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

<a id="canonical-2131222223111010-2111120233203103-2320030020233212-1100313122102001-0220103033100103-2321130110123033-3121232330030010-1331303011210121"></a>

### Direct properties for `stateful_service.advertise_options.advertise_custom.ports.tcp_loadbalancer`

<a id="canonical-0220201320032020-2203230112002331-0221100223202321-0231110232312303-0200330330020033-0010303313232001-1230303212232200-1113110330323200"></a>

#### `stateful_service.advertise_options.advertise_custom.ports.tcp_loadbalancer.domains` property

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
EnumExtractionComplete: false
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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-3031013323100010-2303022332011301-2111223000321312-3030010320032033-0332211210232313-1020012211133232-2000211212001022-2131012202130332"></a>

<a id="canonical-0011110323121132-1312330213311331-2132102031020031-2211100313030310-2133013220122111-2312032121112123-1001103101123031-3321230212310023"></a>

#### `stateful_service.advertise_options.advertise_custom.ports.tcp_loadbalancer.with_sni` property

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

<a id="canonical-2203200033111100-1131311021203132-1232001231010211-3000112112233321-3002013120002033-2111032232010121-0222102100212312-0012310120331313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_in_cluster` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-017.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- stateful_service.advertise_options.advertise_in_cluster

<a id="canonical-2210103033003032-2200113122333020-3110203010210121-2320210330220212-1020233221231313-0320120232122112-1211200220020300-1233222202012012"></a>

Type: `"object"`. single nested block, Optional.

Advertise the workload locally in-cluster.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

<a id="canonical-0222032312012133-0003130002102322-1230030230321201-2122132111321221-3310003323011211-0031221302133112-2323120221211232-0210122023320222"></a>

### Direct properties for `stateful_service.advertise_options.advertise_in_cluster`

- [multi_ports](resources--workload--reference--group-021.md#canonical-1011312021130103-1132201221323011-3121333020103321-3032102100312103-3112201131010023-2123121010313000-2113022332010030-0031032321111332): complete subsection reference.

- [port](resources--workload--reference--group-021.md#canonical-0321122020003212-1202311332110210-0323031000133111-3021322002001202-3302313330222033-1233211302112320-3101312321102200-3230320123113312): complete subsection reference.

<a id="canonical-1011312021130103-1132201221323011-3121333020103321-3032102100312103-3112201131010023-2123121010313000-2113022332010030-0031032321111332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_in_cluster.multi_ports` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-017.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_in_cluster](resources--workload--reference--group-021.md#canonical-2203200033111100-1131311021203132-1232001231010211-3000112112233321-3002013120002033-2111032232010121-0222102100212312-0012310120331313)
- stateful_service.advertise_options.advertise_in_cluster.multi_ports

<a id="canonical-2101030231011132-0020001303322101-0313202013032230-1311002111110211-0202322333122311-2002210131213013-1201333220313300-2233201221100202"></a>

Type: `"object"`. single nested block, Optional.

Multiple Ports. Multiple ports.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

<a id="canonical-3021211120003311-1030133321013233-2223112010232002-1012220011021203-0120021033022133-1101132301002012-2113232121012002-1302331012233330"></a>

### Direct properties for `stateful_service.advertise_options.advertise_in_cluster.multi_ports`

- [ports](resources--workload--reference--group-021.md#canonical-0022302331010133-0021121031232202-3131012213033020-0112302012122330-0231112202323133-0200332330311323-0100100012032123-0020333100103302): complete subsection reference.

<a id="canonical-0022302331010133-0021121031232202-3131012213033020-0112302012122330-0231112202323133-0200332330311323-0100100012032123-0020333100103302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_in_cluster.multi_ports.ports` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-017.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_in_cluster](resources--workload--reference--group-021.md#canonical-2203200033111100-1131311021203132-1232001231010211-3000112112233321-3002013120002033-2111032232010121-0222102100212312-0012310120331313)
- [stateful_service.advertise_options.advertise_in_cluster.multi_ports](resources--workload--reference--group-021.md#canonical-1011312021130103-1132201221323011-3121333020103321-3032102100312103-3112201131010023-2123121010313000-2113022332010030-0031032321111332)
- stateful_service.advertise_options.advertise_in_cluster.multi_ports.ports

<a id="canonical-1202311322020121-0201233210022022-1121011003213102-0100103322302013-0232012313111101-1013003221012010-2200033322321012-2200100111031321"></a>

Type: `"object"`. list nested block, Optional.

Ports. Ports to advertise.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-1011233131333332-0332033313302300-3202221130113101-2032003122312130-2102232213231100-0110212033331312-3001030311333232-3122133322311120"></a>

### Direct properties for `stateful_service.advertise_options.advertise_in_cluster.multi_ports.ports`

- [info](resources--workload--reference--group-021.md#canonical-2321303312100200-2202211220202233-3112100312131322-0210100102330323-3111022100200021-1121332221111210-1312102132112232-2322011231321113): complete subsection reference.

<a id="canonical-0200032100120303-0333230233001131-1220013310033320-2100131223202313-2033102111001233-2113131103011032-1333212301123001-3122331213131011"></a>

<a id="canonical-0221330321112233-2100303231301212-2332312322013220-2203232220121212-0001321100200022-0300223203321223-1310120100320021-0010121110221111"></a>

#### `stateful_service.advertise_options.advertise_in_cluster.multi_ports.ports.name` property

Type: `"string"`. Optional.

Name. Name of the Port.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-2321303312100200-2202211220202233-3112100312131322-0210100102330323-3111022100200021-1121332221111210-1312102132112232-2322011231321113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_in_cluster.multi_ports.ports.info` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-017.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_in_cluster](resources--workload--reference--group-021.md#canonical-2203200033111100-1131311021203132-1232001231010211-3000112112233321-3002013120002033-2111032232010121-0222102100212312-0012310120331313)
- [stateful_service.advertise_options.advertise_in_cluster.multi_ports](resources--workload--reference--group-021.md#canonical-1011312021130103-1132201221323011-3121333020103321-3032102100312103-3112201131010023-2123121010313000-2113022332010030-0031032321111332)
- [stateful_service.advertise_options.advertise_in_cluster.multi_ports.ports](resources--workload--reference--group-021.md#canonical-0022302331010133-0021121031232202-3131012213033020-0112302012122330-0231112202323133-0200332330311323-0100100012032123-0020333100103302)
- stateful_service.advertise_options.advertise_in_cluster.multi_ports.ports.info

<a id="canonical-0200130332302332-2222301012201120-2310220032330333-3002222121101020-1230302321011322-3100012032332001-2023311302211123-2230013203021300"></a>

Type: `"object"`. single nested block, Optional.

Port Information. Port information.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

<a id="canonical-0202113313301212-3301010003002101-1021111030310331-0102110332322302-3130100330111101-1312322203112130-3210012333123211-3021022312232333"></a>

### Direct properties for `stateful_service.advertise_options.advertise_in_cluster.multi_ports.ports.info`

<a id="canonical-3130221203321203-3021000323231122-1323113320020220-1000130300121301-1211302123220230-2022211003300303-2312112232311211-2113320003020102"></a>

#### `stateful_service.advertise_options.advertise_in_cluster.multi_ports.ports.info.port` property

Type: `"number"`. Optional.

Port. Port the workload can be reached on.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-2320011012320311-1333330231132102-1103130001203301-3020032333313230-2001233123220303-0220300131203122-2323100010220012-1220113301032121"></a>

<a id="canonical-3332333020232031-0123213303322231-0222012322020033-1230303121133132-0203322110212132-1211310313313212-3323003032013213-2023201333113311"></a>

#### `stateful_service.advertise_options.advertise_in_cluster.multi_ports.ports.info.protocol` property

Type: `"string"`. Optional.

\[Enum: PROTOCOL\_TCP|PROTOCOL\_HTTP|PROTOCOL\_HTTP2|PROTOCOL\_TLS\_WITH\_SNI|PROTOCOL\_UDP\] Type
of protocol - PROTOCOL\_TCP: TCP TCP - PROTOCOL\_HTTP: HTTP HTTP - PROTOCOL\_HTTP2: HTTP2 HTTP2 -
PROTOCOL\_TLS\_WITH\_SNI: TLS with SNI TLS with SNI - PROTOCOL\_UDP: UDP UDP. Possible values are
\`PROTOCOL\_TCP\`, \`PROTOCOL\_HTTP\`, \`PROTOCOL\_HTTP2\`, \`PROTOCOL\_TLS\_WITH\_SNI\`,
\`PROTOCOL\_UDP\`. Defaults to \`PROTOCOL\_TCP\`.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["PROTOCOL_HTTP","PROTOCOL_HTTP2","PROTOCOL_TCP","PROTOCOL_TLS_WITH_SNI","PROTOCOL_UDP"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
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

- [same_as_port](resources--workload--reference--group-021.md#canonical-1103111031300130-1221311311212132-1132201022332110-3122122133312111-0102213010213001-3223333213110103-2332023033211020-2231231003002102): complete subsection reference.

<a id="canonical-1030320133223011-0120110103133220-1020311301002022-1331100323200020-0310300223001323-2100122313222130-2030301020310122-3331322313220300"></a>

<a id="canonical-1231223301303111-1312132301300202-3321310033122101-0132030311323010-1010002013131102-2212021000223102-2220102010131003-1211232010023301"></a>

#### `stateful_service.advertise_options.advertise_in_cluster.multi_ports.ports.info.target_port` property

Type: `"number"`. Optional.

Exclusive with \[same\_as\_port\] Port the workload is listening on.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-1103111031300130-1221311311212132-1132201022332110-3122122133312111-0102213010213001-3223333213110103-2332023033211020-2231231003002102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_in_cluster.multi_ports.ports.info.same_as_port` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-017.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_in_cluster](resources--workload--reference--group-021.md#canonical-2203200033111100-1131311021203132-1232001231010211-3000112112233321-3002013120002033-2111032232010121-0222102100212312-0012310120331313)
- [stateful_service.advertise_options.advertise_in_cluster.multi_ports](resources--workload--reference--group-021.md#canonical-1011312021130103-1132201221323011-3121333020103321-3032102100312103-3112201131010023-2123121010313000-2113022332010030-0031032321111332)
- [stateful_service.advertise_options.advertise_in_cluster.multi_ports.ports](resources--workload--reference--group-021.md#canonical-0022302331010133-0021121031232202-3131012213033020-0112302012122330-0231112202323133-0200332330311323-0100100012032123-0020333100103302)
- [stateful_service.advertise_options.advertise_in_cluster.multi_ports.ports.info](resources--workload--reference--group-021.md#canonical-2321303312100200-2202211220202233-3112100312131322-0210100102330323-3111022100200021-1121332221111210-1312102132112232-2322011231321113)
- stateful_service.advertise_options.advertise_in_cluster.multi_ports.ports.info.same_as_port

<a id="canonical-2101013222013121-2110221302212303-0201201031210020-2122321112201123-2333210012301301-1210111120002012-1010200031030031-2322023103110020"></a>

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

<a id="canonical-0321122020003212-1202311332110210-0323031000133111-3021322002001202-3302313330222033-1233211302112320-3101312321102200-3230320123113312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_in_cluster.port` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-017.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_in_cluster](resources--workload--reference--group-021.md#canonical-2203200033111100-1131311021203132-1232001231010211-3000112112233321-3002013120002033-2111032232010121-0222102100212312-0012310120331313)
- stateful_service.advertise_options.advertise_in_cluster.port

<a id="canonical-3121323322001330-3110201130122013-3023210123200121-3121103020031013-2111210133032231-3321120121011013-1321332230212222-3012220130010000"></a>

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

<a id="canonical-1333233103311323-3203210132023222-3230301021303130-3100322130032001-0123011110023312-2202222231312303-3200333013221323-3110231323313201"></a>

### Direct properties for `stateful_service.advertise_options.advertise_in_cluster.port`

- [info](resources--workload--reference--group-021.md#canonical-3312122023300110-2123020210133120-1021222230013202-0030331331100202-1120301003301303-3220211111200102-3011033220302123-1201201000223032): complete subsection reference.

<a id="canonical-3312122023300110-2123020210133120-1021222230013202-0030331331100202-1120301003301303-3220211111200102-3011033220302123-1201201000223032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_in_cluster.port.info` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-017.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_in_cluster](resources--workload--reference--group-021.md#canonical-2203200033111100-1131311021203132-1232001231010211-3000112112233321-3002013120002033-2111032232010121-0222102100212312-0012310120331313)
- [stateful_service.advertise_options.advertise_in_cluster.port](resources--workload--reference--group-021.md#canonical-0321122020003212-1202311332110210-0323031000133111-3021322002001202-3302313330222033-1233211302112320-3101312321102200-3230320123113312)
- stateful_service.advertise_options.advertise_in_cluster.port.info

<a id="canonical-2233221110213101-3333013232020021-1221212111133220-0021201212331010-3120122003322101-2213103133131010-3310302331031012-2021311310302300"></a>

Type: `"object"`. single nested block, Optional.

Port Information. Port information.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

<a id="canonical-1131022133012201-0110220230232131-0033323223010302-2232320032101210-0200102232030003-2313013031323333-1213030210020320-2333132222123323"></a>

### Direct properties for `stateful_service.advertise_options.advertise_in_cluster.port.info`

<a id="canonical-0220323213122321-1210233011002012-2103102101311113-3112030102231023-3033300102031003-0120022320223033-3323032301333200-2223023111133320"></a>

#### `stateful_service.advertise_options.advertise_in_cluster.port.info.port` property

Type: `"number"`. Optional.

Port. Port the workload can be reached on.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-2312113123233130-2120211233302212-1130031033033102-1232002022011020-1102001012323103-1011031321301023-3032120013233232-3012330202313021"></a>

<a id="canonical-3212021211102221-3312123030121200-0330232000332102-1211012333302231-3031110301010323-3013120300231012-2013110331201033-3111030132112022"></a>

#### `stateful_service.advertise_options.advertise_in_cluster.port.info.protocol` property

Type: `"string"`. Optional.

\[Enum: PROTOCOL\_TCP|PROTOCOL\_HTTP|PROTOCOL\_HTTP2|PROTOCOL\_TLS\_WITH\_SNI|PROTOCOL\_UDP\] Type
of protocol - PROTOCOL\_TCP: TCP TCP - PROTOCOL\_HTTP: HTTP HTTP - PROTOCOL\_HTTP2: HTTP2 HTTP2 -
PROTOCOL\_TLS\_WITH\_SNI: TLS with SNI TLS with SNI - PROTOCOL\_UDP: UDP UDP. Possible values are
\`PROTOCOL\_TCP\`, \`PROTOCOL\_HTTP\`, \`PROTOCOL\_HTTP2\`, \`PROTOCOL\_TLS\_WITH\_SNI\`,
\`PROTOCOL\_UDP\`. Defaults to \`PROTOCOL\_TCP\`.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["PROTOCOL_HTTP","PROTOCOL_HTTP2","PROTOCOL_TCP","PROTOCOL_TLS_WITH_SNI","PROTOCOL_UDP"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
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

- [same_as_port](resources--workload--reference--group-021.md#canonical-3210020200110330-3333021331320222-0230001130001233-2303222221011122-2131231021120313-1203132023111112-0320302000031110-2020300310221113): complete subsection reference.

<a id="canonical-2322110110120111-1201032021221020-1223200112301020-1201200310301230-0022003130223303-0013023131303003-0002213120132121-1301010323220022"></a>

<a id="canonical-3003201232120113-0012003230120110-0110232233211331-2130132023010310-2033310020010133-3213213123302201-0322230023210331-1030020132122301"></a>

#### `stateful_service.advertise_options.advertise_in_cluster.port.info.target_port` property

Type: `"number"`. Optional.

Exclusive with \[same\_as\_port\] Port the workload is listening on.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-3210020200110330-3333021331320222-0230001130001233-2303222221011122-2131231021120313-1203132023111112-0320302000031110-2020300310221113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_in_cluster.port.info.same_as_port` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-017.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_in_cluster](resources--workload--reference--group-021.md#canonical-2203200033111100-1131311021203132-1232001231010211-3000112112233321-3002013120002033-2111032232010121-0222102100212312-0012310120331313)
- [stateful_service.advertise_options.advertise_in_cluster.port](resources--workload--reference--group-021.md#canonical-0321122020003212-1202311332110210-0323031000133111-3021322002001202-3302313330222033-1233211302112320-3101312321102200-3230320123113312)
- [stateful_service.advertise_options.advertise_in_cluster.port.info](resources--workload--reference--group-021.md#canonical-3312122023300110-2123020210133120-1021222230013202-0030331331100202-1120301003301303-3220211111200102-3011033220302123-1201201000223032)
- stateful_service.advertise_options.advertise_in_cluster.port.info.same_as_port

<a id="canonical-1320010313101000-2330101303322012-1330033120212002-2101332210133030-1123133023210230-0310100223220202-3231133202002320-0323313113120103"></a>

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

<a id="canonical-3300101333202310-3223231311312310-2030111033033203-1011103231310023-1313110120311121-1221031312313020-1210302113302033-2110231331000202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_on_public` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-017.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- stateful_service.advertise_options.advertise_on_public

<a id="canonical-3230130300023120-0320313000010332-2211103230120121-1301300330213210-2020331002232102-2203313321122123-1032130102320212-3030001221320330"></a>

Type: `"object"`. single nested block, Optional.

Advertise this workload via loadbalancer on internet with default VIP.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

<a id="canonical-1222230021032022-2032212321320222-2311110323010130-3132033020100130-0332003110103121-3231322120020310-3011231232010303-1222210001332012"></a>

### Direct properties for `stateful_service.advertise_options.advertise_on_public`

- [multi_ports](resources--workload--reference--group-021.md#canonical-1100112212100001-0021111102323333-1323322022200132-0012103312111033-0030020021321023-2123212213221211-2130231010230301-0201331202212302): complete subsection reference.

- [port](resources--workload--reference--group-024.md#canonical-0120012201110120-2320030110233132-3233102311321321-2202232022223030-2133131131013111-3231210131302330-1022020011332122-3111230313032011): complete subsection reference.

<a id="canonical-1100112212100001-0021111102323333-1323322022200132-0012103312111033-0030020021321023-2123212213221211-2130231010230301-0201331202212302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_on_public.multi_ports` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-017.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-021.md#canonical-3300101333202310-3223231311312310-2030111033033203-1011103231310023-1313110120311121-1221031312313020-1210302113302033-2110231331000202)
- stateful_service.advertise_options.advertise_on_public.multi_ports

<a id="canonical-2123313023223102-3221002033020311-2233013012321200-3032031302212011-0230320220220123-1121032121130033-3122110121303222-2023210302200121"></a>

Type: `"object"`. single nested block, Optional.

Advertise Multiple Ports. Advertise multiple ports.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

<a id="canonical-2330332031133210-0113101111121230-2120030333210201-2001003300023001-1310212120013321-1132122122220000-2222212331013110-0321030120103210"></a>

### Direct properties for `stateful_service.advertise_options.advertise_on_public.multi_ports`

- [ports](resources--workload--reference--group-021.md#canonical-3213200112122133-3131300232313020-2123002301331230-0311302300110122-1031102122031002-1200202213022002-0022210022203323-0120011222030133): complete subsection reference.

<a id="canonical-3213200112122133-3131300232313020-2123002301331230-0311302300110122-1031102122031002-1200202213022002-0022210022203323-0120011222030133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_on_public.multi_ports.ports` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-017.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-021.md#canonical-3300101333202310-3223231311312310-2030111033033203-1011103231310023-1313110120311121-1221031312313020-1210302113302033-2110231331000202)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-021.md#canonical-1100112212100001-0021111102323333-1323322022200132-0012103312111033-0030020021321023-2123212213221211-2130231010230301-0201331202212302)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports

<a id="canonical-3122301330113311-1211330120102320-0321223300100313-0002310201020203-2312111330220321-1201233203021001-1332213323111320-2333022102310010"></a>

Type: `"object"`. list nested block, Optional.

Ports. Ports to advertise.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-2333203220320311-1302021131103223-2301133000113321-0330130231212131-1133221013022211-0102213303123332-3232323312220303-2120123233022301"></a>

### Direct properties for `stateful_service.advertise_options.advertise_on_public.multi_ports.ports`

- [http_loadbalancer](resources--workload--reference--group-021.md#canonical-3321321030220132-0123111201021120-1223031102031120-2000323012112212-1003100231313121-1032130002221132-3301112102311102-0022133230122013): complete subsection reference.

- [port](resources--workload--reference--group-024.md#canonical-1033111330232111-3032000001323323-1021313331201023-3300200110033300-2331233133210221-0213013303022233-3320021122001110-3222333133221131): complete subsection reference.

- [tcp_loadbalancer](resources--workload--reference--group-024.md#canonical-0311120130302300-0212300231130012-0030210231023313-2321020221010333-0011113212201320-1011033002112333-1130110102303303-1322201032210213): complete subsection reference.

<a id="canonical-3321321030220132-0123111201021120-1223031102031120-2000323012112212-1003100231313121-1032130002221132-3301112102311102-0022133230122013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-017.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-021.md#canonical-3300101333202310-3223231311312310-2030111033033203-1011103231310023-1313110120311121-1221031312313020-1210302113302033-2110231331000202)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-021.md#canonical-1100112212100001-0021111102323333-1323322022200132-0012103312111033-0030020021321023-2123212213221211-2130231010230301-0201331202212302)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-021.md#canonical-3213200112122133-3131300232313020-2123002301331230-0311302300110122-1031102122031002-1200202213022002-0022210022203323-0120011222030133)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer

<a id="canonical-2301311020111022-0001121120013200-3133231120021131-0302322210101110-1313212003331131-0103331132030330-2101002113332111-0101322132023302"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for http loadbalancer.

Additional upstream details:

HTTP/HTTPS Load balancer.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

<a id="canonical-2322103022020100-2120033221213212-0003230300103201-1233133010322131-1011333020231120-1311231322211222-3023222233033011-1202000003110332"></a>

### Direct properties for `stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer`

- [default_route](resources--workload--reference--group-021.md#canonical-2103121200313112-2212331322311120-3212220213031300-2000233130310232-1310210011012230-3212322000313311-3302122130101033-3010210130203121): complete subsection reference.

<a id="canonical-1101022311123010-2101133001232032-2332001210313230-0131133232232001-0331023012132122-2002300333313002-3210122110302223-1211220001200122"></a>

<a id="canonical-0011013211201221-1120313033101030-2331122011300311-2202032210013320-2112313112001010-1230332112321322-3221033213332000-3023333102001231"></a>

#### `stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.domains` property

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
EnumExtractionComplete: false
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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

- [http](resources--workload--reference--group-021.md#canonical-1112300021023200-3003230202133313-1122112123303221-2313230112331002-2020220332001032-3122132320101231-0023122110201233-2201313013302101): complete subsection reference.

- [https](resources--workload--reference--group-021.md#canonical-0321203222112202-1111303122230123-2002231022310211-1312303013230233-2031132232132311-1111220203200233-0202310101333222-2011201032133231): complete subsection reference.

- [https_auto_cert](resources--workload--reference--group-023.md#canonical-2032211123003303-0001312001113121-2111212330333101-0133320113201220-2301311001003001-0233320211201111-0011021020303132-3310011112331000): complete subsection reference.

- [specific_routes](resources--workload--reference--group-023.md#canonical-3131312222321330-1311103031033113-2013033101200323-1130033022133322-2103120313300211-0311222332323111-0103003001003203-0330210232300303): complete subsection reference.

<a id="canonical-2103121200313112-2212331322311120-3212220213031300-2000233130310232-1310210011012230-3212322000313311-3302122130101033-3010210130203121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.default_route` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-017.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-021.md#canonical-3300101333202310-3223231311312310-2030111033033203-1011103231310023-1313110120311121-1221031312313020-1210302113302033-2110231331000202)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-021.md#canonical-1100112212100001-0021111102323333-1323322022200132-0012103312111033-0030020021321023-2123212213221211-2130231010230301-0201331202212302)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-021.md#canonical-3213200112122133-3131300232313020-2123002301331230-0311302300110122-1031102122031002-1200202213022002-0022210022203323-0120011222030133)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](resources--workload--reference--group-021.md#canonical-3321321030220132-0123111201021120-1223031102031120-2000323012112212-1003100231313121-1032130002221132-3301112102311102-0022133230122013)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.default_route

<a id="canonical-1212010011222211-2212103032021131-3201011001020212-3010130323013203-1111112222000103-0022101023232120-1012023302220202-3201113223311101"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for default route.

Additional upstream details:

Default route matching all APIs.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

<a id="canonical-1031010131233220-3200322011302321-0130102233221110-3300020321111130-2203103001033223-3200213211120330-1311201100122012-2233011332120021"></a>

### Direct properties for `stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.default_route`

- [auto_host_rewrite](resources--workload--reference--group-021.md#canonical-0131313021021130-1003033233323022-3003300223031111-3220020102201023-0330203300233202-2300322312202230-0103302020020100-3121020133022131): complete subsection reference.

- [disable_host_rewrite](resources--workload--reference--group-021.md#canonical-1320331112210001-0302023001023233-2122203233302223-1303030201212232-1022110321232010-3221213330013131-2033213202300013-1103031030113031): complete subsection reference.

<a id="canonical-0013222013221200-0211030203201121-2231200023023102-1122330332310300-3112331230301231-0031020333303233-1032121311302331-2220021011201003"></a>

<a id="canonical-1030010230133031-2212121112120301-1001332311033010-0030231101132132-1312101100011332-0321231033201310-2023322032111111-1001331000020120"></a>

#### `stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.default_route.host_rewrite` property

Type: `"string"`. Optional.

Exclusive with \[auto\_host\_rewrite disable\_host\_rewrite\] Host header will be swapped with this
value.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-0131313021021130-1003033233323022-3003300223031111-3220020102201023-0330203300233202-2300322312202230-0103302020020100-3121020133022131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.default_route.auto_host_rewrite` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-017.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-021.md#canonical-3300101333202310-3223231311312310-2030111033033203-1011103231310023-1313110120311121-1221031312313020-1210302113302033-2110231331000202)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-021.md#canonical-1100112212100001-0021111102323333-1323322022200132-0012103312111033-0030020021321023-2123212213221211-2130231010230301-0201331202212302)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-021.md#canonical-3213200112122133-3131300232313020-2123002301331230-0311302300110122-1031102122031002-1200202213022002-0022210022203323-0120011222030133)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](resources--workload--reference--group-021.md#canonical-3321321030220132-0123111201021120-1223031102031120-2000323012112212-1003100231313121-1032130002221132-3301112102311102-0022133230122013)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.default_route](resources--workload--reference--group-021.md#canonical-2103121200313112-2212331322311120-3212220213031300-2000233130310232-1310210011012230-3212322000313311-3302122130101033-3010210130203121)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.default_route.auto_host_rewrite

<a id="canonical-3301210220233201-0011310333322200-1331012320123202-3220300130112121-3221220332003132-0112102221120103-1301310110100313-2202110313221033"></a>

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

<a id="canonical-1320331112210001-0302023001023233-2122203233302223-1303030201212232-1022110321232010-3221213330013131-2033213202300013-1103031030113031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.default_route.disable_host_rewrite` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-017.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-021.md#canonical-3300101333202310-3223231311312310-2030111033033203-1011103231310023-1313110120311121-1221031312313020-1210302113302033-2110231331000202)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-021.md#canonical-1100112212100001-0021111102323333-1323322022200132-0012103312111033-0030020021321023-2123212213221211-2130231010230301-0201331202212302)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-021.md#canonical-3213200112122133-3131300232313020-2123002301331230-0311302300110122-1031102122031002-1200202213022002-0022210022203323-0120011222030133)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](resources--workload--reference--group-021.md#canonical-3321321030220132-0123111201021120-1223031102031120-2000323012112212-1003100231313121-1032130002221132-3301112102311102-0022133230122013)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.default_route](resources--workload--reference--group-021.md#canonical-2103121200313112-2212331322311120-3212220213031300-2000233130310232-1310210011012230-3212322000313311-3302122130101033-3010210130203121)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.default_route.disable_host_rewrite

<a id="canonical-3212322111312320-2011113111321213-3221131112020012-0223111012030313-0133032113223012-0202020033021000-2213111320200013-3202213321022132"></a>

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

<a id="canonical-1112300021023200-3003230202133313-1122112123303221-2313230112331002-2020220332001032-3122132320101231-0023122110201233-2201313013302101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.http` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-017.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-021.md#canonical-3300101333202310-3223231311312310-2030111033033203-1011103231310023-1313110120311121-1221031312313020-1210302113302033-2110231331000202)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-021.md#canonical-1100112212100001-0021111102323333-1323322022200132-0012103312111033-0030020021321023-2123212213221211-2130231010230301-0201331202212302)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-021.md#canonical-3213200112122133-3131300232313020-2123002301331230-0311302300110122-1031102122031002-1200202213022002-0022210022203323-0120011222030133)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](resources--workload--reference--group-021.md#canonical-3321321030220132-0123111201021120-1223031102031120-2000323012112212-1003100231313121-1032130002221132-3301112102311102-0022133230122013)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.http

<a id="canonical-2003130120232333-3213011112103023-2300302212020013-2123320011031030-2302320322012031-2220102102303020-2303312333111020-2323133212131033"></a>

Type: `"object"`. single nested block, Optional.

HTTP Choice. Choice for selecting HTTP proxy.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

<a id="canonical-3323231200131100-1020111220031112-3001330002213310-3021001211012130-0233220230323320-1202011323222000-3133301230121111-0133012031012301"></a>

### Direct properties for `stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.http`

<a id="canonical-3113300303222103-3002033230020120-1213200200322022-0332022220001113-0022210211010301-2111002001333021-0021200131023122-1111010321010022"></a>

#### `stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.http.dns_volterra_managed` property

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

<a id="canonical-0203023220302021-2211102012013200-0021310332031323-3323232130131033-2021132301323030-1001130330230201-3233033022203100-2231130000003121"></a>

<a id="canonical-1220312222121223-2302321100212213-0232020203111212-3302311123133010-0213033331230000-1322322233302201-1301101123200031-0110130313020122"></a>

#### `stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.http.port` property

Type: `"number"`. Optional.

Exclusive with \[port\_ranges\] HTTP port to Listen.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-3023001303123220-0303231031303120-0313123313202022-1312103123330203-1001132331002122-3100230123022113-0220023301103321-3230310101013121"></a>

<a id="canonical-1322103013023000-0102320203110302-2022300112131003-1312222102001003-3120122111210213-0332101213112321-2100230312003012-3313300230203113"></a>

#### `stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.http.port_ranges` property

Type: `"string"`. Optional.

Exclusive with \[port\] A string containing a comma separated list of port ranges. Each port range
consists of a single port or two ports separated by '-'.

Additional upstream details:

Each port range consists of a single port or two ports separated by "-".

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-0321203222112202-1111303122230123-2002231022310211-1312303013230233-2031132232132311-1111220203200233-0202310101333222-2011201032133231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-017.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-021.md#canonical-3300101333202310-3223231311312310-2030111033033203-1011103231310023-1313110120311121-1221031312313020-1210302113302033-2110231331000202)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-021.md#canonical-1100112212100001-0021111102323333-1323322022200132-0012103312111033-0030020021321023-2123212213221211-2130231010230301-0201331202212302)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-021.md#canonical-3213200112122133-3131300232313020-2123002301331230-0311302300110122-1031102122031002-1200202213022002-0022210022203323-0120011222030133)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](resources--workload--reference--group-021.md#canonical-3321321030220132-0123111201021120-1223031102031120-2000323012112212-1003100231313121-1032130002221132-3301112102311102-0022133230122013)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https

<a id="canonical-1103313320332021-3330232232003103-1203221233230222-0210132111132001-3203231033220131-0111321211211312-1220022300133122-2301012332230322"></a>

Type: `"object"`. single nested block, Optional.

Choice for selecting HTTP proxy with bring your own certificates.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

<a id="canonical-1220201001202010-3230101011203232-3213311023320120-0231310012230201-2323111231020013-3213312111332303-1211232201033322-1223002313010133"></a>

### Direct properties for `stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https`

<a id="canonical-1030220312311122-1303221323113013-3322313022231100-3232233123113023-3022010323200001-3231031302320121-3202123220320130-3303120121232121"></a>

#### `stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.add_hsts` property

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

<a id="canonical-3103333010112203-3300120021000302-3113102030112311-3211222331120302-3032301102212233-0023300312033221-1113110102130313-2011023321122211"></a>

<a id="canonical-1021230002132201-3012331021333130-0202121311000022-3131001111210003-3332213211100101-3020201220021330-1021003113110232-2300220312203121"></a>

#### `stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.append_server_name` property

Type: `"string"`. Optional.

Exclusive with \[default\_header pass\_through server\_name\] Define the header value for the header
name “server”. If header value is already present, it is not overwritten and passed as-is.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

- [coalescing_options](resources--workload--reference--group-021.md#canonical-0032322300312032-3133333133020222-1033323001010103-0033301033012313-0010011013031032-1033312213230102-1020012313332133-0101232222122320): complete subsection reference.

<a id="canonical-0010112023200002-2211331312331221-2303300101112123-0223011320301132-3033132031021323-2033111221021302-2102322203030010-0133103130310021"></a>

<a id="canonical-2302312212001103-2323303111232211-0111013112030331-0233130021023222-3101232101020032-0312013312200032-1010331130321223-0130330121023020"></a>

#### `stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.connection_idle_timeout` property

Type: `"number"`. Optional.

The idle timeout for downstream connections. The idle timeout is defined as the period in which
there are no active requests. When the idle timeout is reached the connection will be closed. Note
that request based timeouts mean that HTTP/2 PINGs will not keep the connection alive. This is
specified in milliseconds. The default value is 2 minutes.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

- [default_header](resources--workload--reference--group-021.md#canonical-3302211012301032-2012330330011310-0113010232202002-2201321200312113-1221030223303330-2303321333100111-1221003222122203-3231313130213222): complete subsection reference.

- [default_loadbalancer](resources--workload--reference--group-021.md#canonical-1011210312010101-1020002203221023-3210133322003011-0021123113300022-1100330110303301-3110133100300213-0233300110131103-2212321311032331): complete subsection reference.

- [disable_path_normalize](resources--workload--reference--group-021.md#canonical-2100133311003320-1210320110323011-2203203113023312-0222003220120320-0300232211031303-0013132100002101-3100113332320331-1023223233202322): complete subsection reference.

- [enable_path_normalize](resources--workload--reference--group-021.md#canonical-0013221021222113-2311120333122113-2222331121011033-1202110232313200-1112103312000323-2202323130313102-2021301210113102-2303200031021200): complete subsection reference.

- [http_protocol_options](resources--workload--reference--group-022.md#canonical-2011301022111013-3131201322230120-2212231123222010-0001020033322333-2200030310322031-0111211033010310-3021202130330331-2011101110331100): complete subsection reference.

<a id="canonical-3103210013210302-0331030223013121-0321120322310220-1231020132200130-1020123002303311-2012200323123311-3131031112310212-1220023313200103"></a>

<a id="canonical-2103031133121211-2103232111210001-3313110012022312-2302131310023013-2300000222022221-2311310112120200-1131200210021020-2311230210333103"></a>

#### `stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_redirect` property

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

- [non_default_loadbalancer](resources--workload--reference--group-022.md#canonical-1211311132330022-0311311203000110-3311213030033001-3321132331311013-2010301333331322-0033222122122130-0003011001003113-1310232321210203): complete subsection reference.

- [pass_through](resources--workload--reference--group-022.md#canonical-2303013131220012-0130102213302222-0303330312111033-2230221332330230-1112123221011121-3031221011133200-0211002232330312-1323232302313022): complete subsection reference.

<a id="canonical-3312002301312330-0303202330003133-2330010233302110-0203303323232210-3223012232231110-1202032102233320-3102103113210111-0211033021133221"></a>

<a id="canonical-2111133221220333-3003133010133202-3023100013013222-0010131113313122-2310120033003323-2311010103013312-1121110330133001-3000200123233103"></a>

#### `stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.port` property

Type: `"number"`. Optional.

Exclusive with \[port\_ranges\] HTTPS port to Listen.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-2212201332222023-0231302020233200-1331313321202011-0103023102033211-0312212211220313-2230103021031002-2023232201123323-0223131303302020"></a>

<a id="canonical-2012310033030002-2033013212120233-1001022022131320-3312311121013230-1112100022220301-0022330333020021-1321311233132223-3031131011221323"></a>

#### `stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.port_ranges` property

Type: `"string"`. Optional.

Exclusive with \[port\] A string containing a comma separated list of port ranges. Each port range
consists of a single port or two ports separated by '-'.

Additional upstream details:

Each port range consists of a single port or two ports separated by "-".

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-3003122203202022-3213231133123130-1331031303122331-1033321200231201-1333332001002203-0332113323020310-2302302231100222-2331213313222021"></a>

<a id="canonical-3002312123301332-2113011301120202-1030202311231221-0312210301130211-1121102111033322-2330223332212230-2111301101221010-1230110301000101"></a>

#### `stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.server_name` property

Type: `"string"`. Optional.

Exclusive with \[append\_server\_name default\_header pass\_through\] Define the header value for
the header name “server”. This will overwrite existing values, if any, for the server header.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

- [tls_cert_params](resources--workload--reference--group-022.md#canonical-3122033231332110-3033122223220012-3102111111322223-1112011032012130-2022100233000021-3313120013112202-0130032212223212-1302033022300020): complete subsection reference.

- [tls_parameters](resources--workload--reference--group-022.md#canonical-0313313011110003-2213210313322220-3303210132002202-0210213302202103-0020220120002300-2001021203230113-1010112211132212-1120303303303311): complete subsection reference.

<a id="canonical-0032322300312032-3133333133020222-1033323001010103-0033301033012313-0010011013031032-1033312213230102-1020012313332133-0101232222122320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.coalescing_options` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-017.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-021.md#canonical-3300101333202310-3223231311312310-2030111033033203-1011103231310023-1313110120311121-1221031312313020-1210302113302033-2110231331000202)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-021.md#canonical-1100112212100001-0021111102323333-1323322022200132-0012103312111033-0030020021321023-2123212213221211-2130231010230301-0201331202212302)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-021.md#canonical-3213200112122133-3131300232313020-2123002301331230-0311302300110122-1031102122031002-1200202213022002-0022210022203323-0120011222030133)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](resources--workload--reference--group-021.md#canonical-3321321030220132-0123111201021120-1223031102031120-2000323012112212-1003100231313121-1032130002221132-3301112102311102-0022133230122013)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](resources--workload--reference--group-021.md#canonical-0321203222112202-1111303122230123-2002231022310211-1312303013230233-2031132232132311-1111220203200233-0202310101333222-2011201032133231)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.coalescing_options

<a id="canonical-2023130222210133-2222322302133031-2111331313002301-2010300110033000-2111233110222212-2101031200021200-0030123103311010-1211220100311013"></a>

Type: `"object"`. single nested block, Optional.

TLS connection coalescing configuration (not compatible with mTLS).

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("default_coalescing",
    "strict_coalescing")}
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
  "x-ves-oneof-field-coalescing_choice": "[\"default_coalescing\",\"strict_coalescing\"]"
}
```

Terraform syntax:

```terraform
coalescing_options {
  # Configure direct properties listed below.
}
```

<a id="canonical-1103220322211033-3131330001123001-2103201200331110-1122332112100121-1101112030203130-0322312322313313-3203320231231131-3201001232103300"></a>

### Direct properties for `stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.coalescing_options`

- [default_coalescing](resources--workload--reference--group-021.md#canonical-0323013230213031-1132102200313100-1303100223123121-2002202111231000-2211112122212310-1320302113220323-3130102212301011-1031100201113000): complete subsection reference.

- [strict_coalescing](resources--workload--reference--group-021.md#canonical-2333002032030300-3101023000031212-2331322223302321-3120013331123120-0223120020320211-2311020321211220-3013133120013310-3230310210211321): complete subsection reference.

<a id="canonical-0323013230213031-1132102200313100-1303100223123121-2002202111231000-2211112122212310-1320302113220323-3130102212301011-1031100201113000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.coalescing_options.default_coalescing` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-017.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-021.md#canonical-3300101333202310-3223231311312310-2030111033033203-1011103231310023-1313110120311121-1221031312313020-1210302113302033-2110231331000202)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-021.md#canonical-1100112212100001-0021111102323333-1323322022200132-0012103312111033-0030020021321023-2123212213221211-2130231010230301-0201331202212302)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-021.md#canonical-3213200112122133-3131300232313020-2123002301331230-0311302300110122-1031102122031002-1200202213022002-0022210022203323-0120011222030133)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](resources--workload--reference--group-021.md#canonical-3321321030220132-0123111201021120-1223031102031120-2000323012112212-1003100231313121-1032130002221132-3301112102311102-0022133230122013)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](resources--workload--reference--group-021.md#canonical-0321203222112202-1111303122230123-2002231022310211-1312303013230233-2031132232132311-1111220203200233-0202310101333222-2011201032133231)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.coalescing_options](resources--workload--reference--group-021.md#canonical-0032322300312032-3133333133020222-1033323001010103-0033301033012313-0010011013031032-1033312213230102-1020012313332133-0101232222122320)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.coalescing_options.default_coalescing

<a id="canonical-1011110323020301-2211203123121001-3211000010212012-1030200010323220-2320002003120310-0333301212002002-0300133221002322-0302231022002310"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
default_coalescing = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2333002032030300-3101023000031212-2331322223302321-3120013331123120-0223120020320211-2311020321211220-3013133120013310-3230310210211321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.coalescing_options.strict_coalescing` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-017.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-021.md#canonical-3300101333202310-3223231311312310-2030111033033203-1011103231310023-1313110120311121-1221031312313020-1210302113302033-2110231331000202)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-021.md#canonical-1100112212100001-0021111102323333-1323322022200132-0012103312111033-0030020021321023-2123212213221211-2130231010230301-0201331202212302)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-021.md#canonical-3213200112122133-3131300232313020-2123002301331230-0311302300110122-1031102122031002-1200202213022002-0022210022203323-0120011222030133)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](resources--workload--reference--group-021.md#canonical-3321321030220132-0123111201021120-1223031102031120-2000323012112212-1003100231313121-1032130002221132-3301112102311102-0022133230122013)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](resources--workload--reference--group-021.md#canonical-0321203222112202-1111303122230123-2002231022310211-1312303013230233-2031132232132311-1111220203200233-0202310101333222-2011201032133231)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.coalescing_options](resources--workload--reference--group-021.md#canonical-0032322300312032-3133333133020222-1033323001010103-0033301033012313-0010011013031032-1033312213230102-1020012313332133-0101232222122320)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.coalescing_options.strict_coalescing

<a id="canonical-3310230300031301-3132011210333300-2310111001021033-1322031020111133-1321110303032321-1020232001122101-0213212110120211-2131103233021100"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
strict_coalescing = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3302211012301032-2012330330011310-0113010232202002-2201321200312113-1221030223303330-2303321333100111-1221003222122203-3231313130213222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.default_header` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-017.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-021.md#canonical-3300101333202310-3223231311312310-2030111033033203-1011103231310023-1313110120311121-1221031312313020-1210302113302033-2110231331000202)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-021.md#canonical-1100112212100001-0021111102323333-1323322022200132-0012103312111033-0030020021321023-2123212213221211-2130231010230301-0201331202212302)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-021.md#canonical-3213200112122133-3131300232313020-2123002301331230-0311302300110122-1031102122031002-1200202213022002-0022210022203323-0120011222030133)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](resources--workload--reference--group-021.md#canonical-3321321030220132-0123111201021120-1223031102031120-2000323012112212-1003100231313121-1032130002221132-3301112102311102-0022133230122013)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](resources--workload--reference--group-021.md#canonical-0321203222112202-1111303122230123-2002231022310211-1312303013230233-2031132232132311-1111220203200233-0202310101333222-2011201032133231)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.default_header

<a id="canonical-3112113313001000-0021232202333011-0032111130333333-2113212131211302-3332321112100002-3013123030023320-1032302123331233-1102010100210320"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
default_header = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1011210312010101-1020002203221023-3210133322003011-0021123113300022-1100330110303301-3110133100300213-0233300110131103-2212321311032331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.default_loadbalancer` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-017.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-021.md#canonical-3300101333202310-3223231311312310-2030111033033203-1011103231310023-1313110120311121-1221031312313020-1210302113302033-2110231331000202)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-021.md#canonical-1100112212100001-0021111102323333-1323322022200132-0012103312111033-0030020021321023-2123212213221211-2130231010230301-0201331202212302)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-021.md#canonical-3213200112122133-3131300232313020-2123002301331230-0311302300110122-1031102122031002-1200202213022002-0022210022203323-0120011222030133)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](resources--workload--reference--group-021.md#canonical-3321321030220132-0123111201021120-1223031102031120-2000323012112212-1003100231313121-1032130002221132-3301112102311102-0022133230122013)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](resources--workload--reference--group-021.md#canonical-0321203222112202-1111303122230123-2002231022310211-1312303013230233-2031132232132311-1111220203200233-0202310101333222-2011201032133231)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.default_loadbalancer

<a id="canonical-1331320110212311-3032323111202302-1113023221223032-2312102233222332-0032233010000123-0332133233110203-2303010103230033-3010332010132010"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
default_loadbalancer = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2100133311003320-1210320110323011-2203203113023312-0222003220120320-0300232211031303-0013132100002101-3100113332320331-1023223233202322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.disable_path_normalize` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-017.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-021.md#canonical-3300101333202310-3223231311312310-2030111033033203-1011103231310023-1313110120311121-1221031312313020-1210302113302033-2110231331000202)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-021.md#canonical-1100112212100001-0021111102323333-1323322022200132-0012103312111033-0030020021321023-2123212213221211-2130231010230301-0201331202212302)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-021.md#canonical-3213200112122133-3131300232313020-2123002301331230-0311302300110122-1031102122031002-1200202213022002-0022210022203323-0120011222030133)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](resources--workload--reference--group-021.md#canonical-3321321030220132-0123111201021120-1223031102031120-2000323012112212-1003100231313121-1032130002221132-3301112102311102-0022133230122013)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](resources--workload--reference--group-021.md#canonical-0321203222112202-1111303122230123-2002231022310211-1312303013230233-2031132232132311-1111220203200233-0202310101333222-2011201032133231)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.disable_path_normalize

<a id="canonical-0333011222200203-0123021022332031-0331202210313312-1313033310022203-3313131232320212-1121201112202203-2313312210020230-0000323230113011"></a>

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
disable_path_normalize = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0013221021222113-2311120333122113-2222331121011033-1202110232313200-1112103312000323-2202323130313102-2021301210113102-2303200031021200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.enable_path_normalize` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-017.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-021.md#canonical-3300101333202310-3223231311312310-2030111033033203-1011103231310023-1313110120311121-1221031312313020-1210302113302033-2110231331000202)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-021.md#canonical-1100112212100001-0021111102323333-1323322022200132-0012103312111033-0030020021321023-2123212213221211-2130231010230301-0201331202212302)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-021.md#canonical-3213200112122133-3131300232313020-2123002301331230-0311302300110122-1031102122031002-1200202213022002-0022210022203323-0120011222030133)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](resources--workload--reference--group-021.md#canonical-3321321030220132-0123111201021120-1223031102031120-2000323012112212-1003100231313121-1032130002221132-3301112102311102-0022133230122013)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](resources--workload--reference--group-021.md#canonical-0321203222112202-1111303122230123-2002231022310211-1312303013230233-2031132232132311-1111220203200233-0202310101333222-2011201032133231)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.enable_path_normalize

<a id="canonical-0101220223231110-3130031001223312-0102020031233102-0131230220012022-3100332031302231-0300313122231023-3100123021020333-2021130033233101"></a>

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
enable_path_normalize = {}
```

This is an empty object or choice marker. It has no direct properties.
