---
page_title: "xcsh_workload reference"
subcategory: "Container"
description: "Complete grouped canonical reference for xcsh_workload reference."
---

# xcsh_workload reference

<a id="canonical-2100302100102130-3203111030030320-1021211121113332-2222133102210213-2311333302120120-2103110313130103-1010130203101033-3313322121031033"></a>

## name property — port / 232203231033 / 4

Type: `"string"`. Optional.

Name. Name of the Port.

Upstream description:

Name of the Port.

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

<a id="canonical-0330003030220232-0101123332202122-1032013321113013-0332002121002220-3320233211210100-3210211023123212-0321012002222113-1202031303210033"></a>

## Next pages — port / 232203231033 / 5

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.port.info](resources--workload--reference--group-025.md#canonical-0331001131000022-2323220112030233-1321313203210102-0000133120012301-1132201221311102-1120102021211023-0021031131323031-1301113313303210)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-021.md#canonical-3213200112122133-3131300232313020-2123002301331230-0311302300110122-1031102122031002-1200202213022002-0022210022203323-0120011222030133)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-0331001131000022-2323220112030233-1321313203210102-0000133120012301-1132201221311102-1120102021211023-0021031131323031-1301113313303210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1212321123332021-2310022320201231-3133100001302201-0011133011230112-2221101122012120-3012133201301130-2210203011122103-2031200231220002"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.port.info — info / 033012202203 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-017.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-021.md#canonical-3300101333202310-3223231311312310-2030111033033203-1011103231310023-1313110120311121-1221031312313020-1210302113302033-2110231331000202)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-021.md#canonical-1100112212100001-0021111102323333-1323322022200132-0012103312111033-0030020021321023-2123212213221211-2130231010230301-0201331202212302)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-021.md#canonical-3213200112122133-3131300232313020-2123002301331230-0311302300110122-1031102122031002-1200202213022002-0022210022203323-0120011222030133)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.port](resources--workload--reference--group-024.md#canonical-1033111330232111-3032000001323323-1021313331201023-3300200110033300-2331233133210221-0213013303022233-3320021122001110-3222333133221131)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.port.info

<a id="canonical-1123323130000311-3110132001030230-2011103330113022-2210322201013012-0020320122113211-1310132232311301-0313331123231112-2302001303302120"></a>

Type: `"object"`. single nested block, Optional.

Port Information. Port information.

Upstream description:

Port information.

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

<a id="canonical-0012031231221011-3001211212120133-3323321201300202-1212021030112011-2131310301321202-0001113223112023-0303322110323012-2300210033122231"></a>

## Direct properties — info / 033012202203 / 3

<a id="canonical-3020113232333120-1303210103032010-2231301311100010-0310122312313233-1122312121000212-2302103200011311-1110321303123001-0321200211032203"></a>

<a id="canonical-0001102332012203-1303203102202010-0100221223020331-1301010110112302-1230022333032230-3103103133100333-1111003100302201-2133223320102231"></a>

## port property — info / 033012202203 / 4

Type: `"number"`. Optional.

Port. Port the workload can be reached on.

Upstream description:

Port the workload can be reached on.

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

<a id="canonical-2210330032312323-2221013102113012-2333001223233203-3032101320212030-1131211123213000-2223031331013312-3011231300302312-3112130013102201"></a>

<a id="canonical-1022110230323210-0300001100003010-1230221010030012-2122211123020303-2310000332211223-1212200313123133-1020222010202022-2321002032131012"></a>

## protocol property — info / 033012202203 / 5

Type: `"string"`. Optional.

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

- [same_as_port](resources--workload--reference--group-025.md#canonical-3011211100011021-1020120303233022-2011202010033103-0303132202032222-3021031023333122-3012130222312132-3231112000123032-0022000112313330): complete subsection reference.

<a id="canonical-3133212011200202-3032211331303031-2031313312230222-3133313021021033-1320133311002211-1010323213021222-0031300031032310-0133030233322210"></a>

<a id="canonical-3310331320223330-1202211321020120-0130102130232010-3103101321023321-2220122333302032-2201002320130001-1300122212201222-2110330022200323"></a>

## target_port property — info / 033012202203 / 6

Type: `"number"`. Optional.

Exclusive with \[same\_as\_port\] Port the workload is listening on.

Upstream description:

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

<a id="canonical-2312000010323120-2311130302131000-2131323020112032-1023211132012120-0002010333210302-2000222200323101-2001102201302202-2220221003333233"></a>

## Next pages — info / 033012202203 / 7

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.port.info.same_as_port](resources--workload--reference--group-025.md#canonical-3011211100011021-1020120303233022-2011202010033103-0303132202032222-3021031023333122-3012130222312132-3231112000123032-0022000112313330)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.port](resources--workload--reference--group-024.md#canonical-1033111330232111-3032000001323323-1021313331201023-3300200110033300-2331233133210221-0213013303022233-3320021122001110-3222333133221131)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-3011211100011021-1020120303233022-2011202010033103-0303132202032222-3021031023333122-3012130222312132-3231112000123032-0022000112313330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3121010211111222-2112232122010132-0203103103200131-2212022211013021-1323033331301300-3020132111322112-1211102012202101-3210331211110003"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.port.info.same_as_port — same_as_port / 011220300313 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-017.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-021.md#canonical-3300101333202310-3223231311312310-2030111033033203-1011103231310023-1313110120311121-1221031312313020-1210302113302033-2110231331000202)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-021.md#canonical-1100112212100001-0021111102323333-1323322022200132-0012103312111033-0030020021321023-2123212213221211-2130231010230301-0201331202212302)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-021.md#canonical-3213200112122133-3131300232313020-2123002301331230-0311302300110122-1031102122031002-1200202213022002-0022210022203323-0120011222030133)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.port](resources--workload--reference--group-024.md#canonical-1033111330232111-3032000001323323-1021313331201023-3300200110033300-2331233133210221-0213013303022233-3320021122001110-3222333133221131)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.port.info](resources--workload--reference--group-025.md#canonical-0331001131000022-2323220112030233-1321313203210102-0000133120012301-1132201221311102-1120102021211023-0021031131323031-1301113313303210)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.port.info.same_as_port

<a id="canonical-0200011130320211-3331122231311320-3321203312133103-3021123200211033-1030213313011000-0311323100320312-2103333212113222-0020301322222111"></a>

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
same_as_port = {}
```

<a id="canonical-0011021113131310-2230302032123313-2222300000330333-2121132301211031-1330132001301313-3222321232300301-2123130202101132-3121023130223010"></a>

## Direct properties — same_as_port / 011220300313 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3022311313130013-1203032120031210-1021330210320232-0011020232332222-2100232222012110-2103300103233122-1220122110321102-3033332211111030"></a>

## Next pages — same_as_port / 011220300313 / 4

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.port.info](resources--workload--reference--group-025.md#canonical-0331001131000022-2323220112030233-1321313203210102-0000133120012301-1132201221311102-1120102021211023-0021031131323031-1301113313303210)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-0311120130302300-0212300231130012-0030210231023313-2321020221010333-0011113212201320-1011033002112333-1130110102303303-1322201032210213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3123320131311013-2330103210013221-2312030211311321-0023202321220113-0032003330321203-0023332101233331-2311102102210133-0320210011323031"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.tcp_loadbalancer — tcp_loadbalancer / 233202230323 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-017.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-021.md#canonical-3300101333202310-3223231311312310-2030111033033203-1011103231310023-1313110120311121-1221031312313020-1210302113302033-2110231331000202)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-021.md#canonical-1100112212100001-0021111102323333-1323322022200132-0012103312111033-0030020021321023-2123212213221211-2130231010230301-0201331202212302)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-021.md#canonical-3213200112122133-3131300232313020-2123002301331230-0311302300110122-1031102122031002-1200202213022002-0022210022203323-0120011222030133)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.tcp_loadbalancer

<a id="canonical-0330110220302002-0003010112300031-1231111222230301-3030002303002201-3000310320222220-0011111111203211-1313233223121002-1101311212000232"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
tcp_loadbalancer {
  # Configure direct properties listed below.
}
```

<a id="canonical-1111213201210222-3300313133333301-0112102223331333-0010101021223310-3301203032323232-1031221321100202-2302233210020203-0330200211230300"></a>

## Direct properties — tcp_loadbalancer / 233202230323 / 3

<a id="canonical-1213101321003113-3300310130213203-0221223012202320-3120131023203022-0230100111112232-3121220332223221-2001031133322130-1201010001021332"></a>

<a id="canonical-3320201011321001-2310220111033102-2102313232130213-0101203323333222-0113320202010102-3130101232203022-2133231111222303-0000232103330311"></a>

## domains property — tcp_loadbalancer / 233202230323 / 4

Type: `["list", "string"]`. Optional.

List of additional domains (host/authority header) that will be matched to this loadbalancer.
Domains are also used for SNI matching if the is true Domains also indicate the list of names for
which DNS resolution will be done by VER.

Upstream description:

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

<a id="canonical-2322222213321113-1002103222311201-3333323002133032-1001232003213133-0223202312131330-0022323030012012-2323133332330000-1003122130311102"></a>

<a id="canonical-1010101110023112-3020213333322200-2312101222222310-3223223230322232-0310322221120311-3023113001100332-0112000312002201-3212022131313122"></a>

## with_sni property — tcp_loadbalancer / 233202230323 / 5

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

<a id="canonical-0032303002123102-2030213232132003-1330201200130121-2131020323122132-3321012222320033-0130030221123111-0302311201110223-3132121321313200"></a>

## Next pages — tcp_loadbalancer / 233202230323 / 6

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-021.md#canonical-3213200112122133-3131300232313020-2123002301331230-0311302300110122-1031102122031002-1200202213022002-0022210022203323-0120011222030133)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-0120012201110120-2320030110233132-3233102311321321-2202232022223030-2133131131013111-3231210131302330-1022020011332122-3111230313032011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2230032002011301-2023022132322133-3131223023133011-2033131203111311-0222030022221130-3322132003313000-2112200121100231-2302232202321032"></a>

## stateful_service.advertise_options.advertise_on_public.port — port / 122321332300 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-017.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-021.md#canonical-3300101333202310-3223231311312310-2030111033033203-1011103231310023-1313110120311121-1221031312313020-1210302113302033-2110231331000202)
- stateful_service.advertise_options.advertise_on_public.port

<a id="canonical-2011211110022221-3331232033210213-0203111010001323-1303330023013311-1132011230330113-0112232301200021-3313222111003031-3220313302301232"></a>

Type: `"object"`. single nested block, Optional.

Advertise Port. Advertise single port.

Upstream description:

Advertise single port.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("http_loadbalancer",
    "tcp_loadbalancer")}
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
  "x-ves-oneof-field-advertise_choice": "[\"http_loadbalancer\",\"tcp_loadbalancer\"]"
}
```

Terraform syntax:

```terraform
port {
  # Configure direct properties listed below.
}
```

<a id="canonical-3111011112212331-3222130313300000-2312203223003103-2212123331133332-0133201122032120-2132323333032031-1013220320222321-0130202332012011"></a>

## Direct properties — port / 122321332300 / 3

- [http_loadbalancer](resources--workload--reference--group-025.md#canonical-2203002110120031-3203013111221013-1321101003110222-2103121303322022-0021002301311230-0203100203111101-0321302220100020-1320220233023212): complete subsection reference.

- [port](resources--workload--reference--group-028.md#canonical-2323101323000230-3130303013113131-1312013231212223-1212222320213220-2013000033112020-2320301032133202-3111123300030002-2013131210032122): complete subsection reference.

- [tcp_loadbalancer](resources--workload--reference--group-028.md#canonical-0201031010010211-0110122211331300-2100213123230032-2210131211202000-0123012121332012-2132032030312023-0111231211021031-2320223020033130): complete subsection reference.

<a id="canonical-2311013222111111-0301023101230323-1021121220320300-1101332322002323-1211031301021232-2122123001321312-2222302213220033-3322032033100330"></a>

## Next pages — port / 122321332300 / 4

- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-025.md#canonical-2203002110120031-3203013111221013-1321101003110222-2103121303322022-0021002301311230-0203100203111101-0321302220100020-1320220233023212)
- [stateful_service.advertise_options.advertise_on_public.port.port](resources--workload--reference--group-028.md#canonical-2323101323000230-3130303013113131-1312013231212223-1212222320213220-2013000033112020-2320301032133202-3111123300030002-2013131210032122)
- [stateful_service.advertise_options.advertise_on_public.port.tcp_loadbalancer](resources--workload--reference--group-028.md#canonical-0201031010010211-0110122211331300-2100213123230032-2210131211202000-0123012121332012-2132032030312023-0111231211021031-2320223020033130)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-021.md#canonical-3300101333202310-3223231311312310-2030111033033203-1011103231310023-1313110120311121-1221031312313020-1210302113302033-2110231331000202)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-2203002110120031-3203013111221013-1321101003110222-2103121303322022-0021002301311230-0203100203111101-0321302220100020-1320220233023212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0031022120323123-1300232231222300-0233120103131123-0011331220101221-3131031312123310-1100030303212222-1301220030333010-1213200200222302"></a>

## stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer — http_loadbalancer / 220221212310 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-017.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-021.md#canonical-3300101333202310-3223231311312310-2030111033033203-1011103231310023-1313110120311121-1221031312313020-1210302113302033-2110231331000202)
- [stateful_service.advertise_options.advertise_on_public.port](resources--workload--reference--group-025.md#canonical-0120012201110120-2320030110233132-3233102311321321-2202232022223030-2133131131013111-3231210131302330-1022020011332122-3111230313032011)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer

<a id="canonical-1110202200020203-0321211011012032-1211132013323221-1111222133120231-2133112223211221-0210333311131231-1312112020313312-3102213321231222"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for http loadbalancer.

Upstream description:

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

<a id="canonical-0212030113032331-1111120012300112-1333221130100231-2222303131123220-0102131021002331-0331010012021023-3213102313120001-2313310311131122"></a>

## Direct properties — http_loadbalancer / 220221212310 / 3

- [default_route](resources--workload--reference--group-025.md#canonical-1230222113003210-1302101103221220-0001131132130300-2031003120223011-0320021111213133-2113030112031032-1113221220203332-1122131330203202): complete subsection reference.

<a id="canonical-2323002313303031-3022033212003133-2020203101221231-0011011210103130-3222120112213021-3003113301023113-2033023133201301-3131131101213130"></a>

<a id="canonical-1023312231103221-2333002303003123-0203321232130001-2000022130333020-1222213300112032-1330000312213331-3300201031130100-0133100120302012"></a>

## domains property — http_loadbalancer / 220221212310 / 4

Type: `["list", "string"]`. Optional.

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

- [http](resources--workload--reference--group-025.md#canonical-2123333100103222-2211301302000212-3113312132112323-3231132113021003-3013302213020120-3013201031021330-2313212132310112-0011220113322010): complete subsection reference.

- [https](resources--workload--reference--group-025.md#canonical-2320131312110211-0233203100020221-1102210111313012-2101323032112112-2012023213113210-1213111131330011-1020010101131133-0020231202232230): complete subsection reference.

- [https_auto_cert](resources--workload--reference--group-026.md#canonical-2220320330100322-1330210101012132-0301302331113332-3313313132203223-2330123211123011-0222332113113122-0203030132101203-1231012010033022): complete subsection reference.

- [specific_routes](resources--workload--reference--group-027.md#canonical-2300220120202032-2120133330213102-2332100100003133-0101032330232302-2302101320100002-1302333333023222-3232223323232230-0320110220223201): complete subsection reference.

<a id="canonical-2113300023303132-0101001203302211-2211323013331213-2303110201330301-1320103000233203-3211233101011022-1113002001132203-1121320332201221"></a>

## Next pages — http_loadbalancer / 220221212310 / 5

- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.default_route](resources--workload--reference--group-025.md#canonical-1230222113003210-1302101103221220-0001131132130300-2031003120223011-0320021111213133-2113030112031032-1113221220203332-1122131330203202)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.http](resources--workload--reference--group-025.md#canonical-2123333100103222-2211301302000212-3113312132112323-3231132113021003-3013302213020120-3013201031021330-2313212132310112-0011220113322010)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https](resources--workload--reference--group-025.md#canonical-2320131312110211-0233203100020221-1102210111313012-2101323032112112-2012023213113210-1213111131330011-1020010101131133-0020231202232230)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert](resources--workload--reference--group-026.md#canonical-2220320330100322-1330210101012132-0301302331113332-3313313132203223-2330123211123011-0222332113113122-0203030132101203-1231012010033022)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes](resources--workload--reference--group-027.md#canonical-2300220120202032-2120133330213102-2332100100003133-0101032330232302-2302101320100002-1302333333023222-3232223323232230-0320110220223201)
- [stateful_service.advertise_options.advertise_on_public.port](resources--workload--reference--group-025.md#canonical-0120012201110120-2320030110233132-3233102311321321-2202232022223030-2133131131013111-3231210131302330-1022020011332122-3111230313032011)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-1230222113003210-1302101103221220-0001131132130300-2031003120223011-0320021111213133-2113030112031032-1113221220203332-1122131330203202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0021121310220000-3232333201333331-1011201110213121-3031123201102302-3013110100313201-1303102220011222-1220012213202231-2211130002003210"></a>

## stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.default_route — default_route / 232010021000 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-017.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-021.md#canonical-3300101333202310-3223231311312310-2030111033033203-1011103231310023-1313110120311121-1221031312313020-1210302113302033-2110231331000202)
- [stateful_service.advertise_options.advertise_on_public.port](resources--workload--reference--group-025.md#canonical-0120012201110120-2320030110233132-3233102311321321-2202232022223030-2133131131013111-3231210131302330-1022020011332122-3111230313032011)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-025.md#canonical-2203002110120031-3203013111221013-1321101003110222-2103121303322022-0021002301311230-0203100203111101-0321302220100020-1320220233023212)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.default_route

<a id="canonical-3020233130031022-0132200101301230-0001220300230213-0003131213313102-0322112331110202-0301001231012031-0212311220133332-2303312223332003"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for default route.

Upstream description:

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

<a id="canonical-1300122113230112-2203012130031231-3213333321021132-3301120110222220-1233303010130032-1213300322221333-0220330230202332-2233120112230212"></a>

## Direct properties — default_route / 232010021000 / 3

- [auto_host_rewrite](resources--workload--reference--group-025.md#canonical-0211201203322103-1222300012213023-2203220033331000-0122102330320103-2310130013032332-0000210131220332-2232223001320221-3020232120312020): complete subsection reference.

- [disable_host_rewrite](resources--workload--reference--group-025.md#canonical-1331012202203300-3020132310323120-2332120133122110-3100201302203100-1013233001123311-0010112022200322-2232132222022220-3301132322023332): complete subsection reference.

<a id="canonical-3301310010222332-1201033301211211-2233302102212123-2031232312300003-2323001123210322-1012013333033203-0001321020013322-2012302013202302"></a>

<a id="canonical-3321100102020130-2111021102203321-0212321120203312-2200120213020332-3211122031321103-2013310302113002-3123020112310011-2100032330222100"></a>

## host_rewrite property — default_route / 232010021000 / 4

Type: `"string"`. Optional.

Exclusive with \[auto\_host\_rewrite disable\_host\_rewrite\] Host header will be swapped with this
value.

Upstream description:

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

<a id="canonical-1230322301032221-1220103300003220-3003123102010322-3232200022123312-2101232013010331-2123133202010122-0201121130202122-3333033121311103"></a>

## Next pages — default_route / 232010021000 / 5

- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.default_route.auto_host_rewrite](resources--workload--reference--group-025.md#canonical-0211201203322103-1222300012213023-2203220033331000-0122102330320103-2310130013032332-0000210131220332-2232223001320221-3020232120312020)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.default_route.disable_host_rewrite](resources--workload--reference--group-025.md#canonical-1331012202203300-3020132310323120-2332120133122110-3100201302203100-1013233001123311-0010112022200322-2232132222022220-3301132322023332)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-025.md#canonical-2203002110120031-3203013111221013-1321101003110222-2103121303322022-0021002301311230-0203100203111101-0321302220100020-1320220233023212)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-0211201203322103-1222300012213023-2203220033331000-0122102330320103-2310130013032332-0000210131220332-2232223001320221-3020232120312020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3121211210333032-1011002123310223-0321120020202312-1220302202230132-0021120100322132-2231122322301310-2202313121302102-2131303023220012"></a>

## stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.default_route.auto_host_rewrite — auto_host_rewrite / 201310033031 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-017.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-021.md#canonical-3300101333202310-3223231311312310-2030111033033203-1011103231310023-1313110120311121-1221031312313020-1210302113302033-2110231331000202)
- [stateful_service.advertise_options.advertise_on_public.port](resources--workload--reference--group-025.md#canonical-0120012201110120-2320030110233132-3233102311321321-2202232022223030-2133131131013111-3231210131302330-1022020011332122-3111230313032011)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-025.md#canonical-2203002110120031-3203013111221013-1321101003110222-2103121303322022-0021002301311230-0203100203111101-0321302220100020-1320220233023212)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.default_route](resources--workload--reference--group-025.md#canonical-1230222113003210-1302101103221220-0001131132130300-2031003120223011-0320021111213133-2113030112031032-1113221220203332-1122131330203202)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.default_route.auto_host_rewrite

<a id="canonical-2121030212000213-0331321113101120-0001213201201231-0232221132310330-0033333002023113-3001003322232320-2203001312322123-1020222223111303"></a>

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
auto_host_rewrite = {}
```

<a id="canonical-3123000201203231-2230202301310113-1212113110322003-2322131221300020-3300013233133021-3103100311131101-0122222023200232-0222213102132310"></a>

## Direct properties — auto_host_rewrite / 201310033031 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1301010303222122-2312001311202022-3113322013230323-1301003100022313-3200110223111210-0031202030300113-2303132322310310-1000311212123212"></a>

## Next pages — auto_host_rewrite / 201310033031 / 4

- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.default_route](resources--workload--reference--group-025.md#canonical-1230222113003210-1302101103221220-0001131132130300-2031003120223011-0320021111213133-2113030112031032-1113221220203332-1122131330203202)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-1331012202203300-3020132310323120-2332120133122110-3100201302203100-1013233001123311-0010112022200322-2232132222022220-3301132322023332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3323122332102203-3303332133000311-3102111221033333-1103223000303300-2113213320301003-2232033322222300-3320212331231011-1003023013133213"></a>

## stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.default_route.disable_host_rewrite — disable_host_rewrite / 221213033230 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-017.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-021.md#canonical-3300101333202310-3223231311312310-2030111033033203-1011103231310023-1313110120311121-1221031312313020-1210302113302033-2110231331000202)
- [stateful_service.advertise_options.advertise_on_public.port](resources--workload--reference--group-025.md#canonical-0120012201110120-2320030110233132-3233102311321321-2202232022223030-2133131131013111-3231210131302330-1022020011332122-3111230313032011)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-025.md#canonical-2203002110120031-3203013111221013-1321101003110222-2103121303322022-0021002301311230-0203100203111101-0321302220100020-1320220233023212)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.default_route](resources--workload--reference--group-025.md#canonical-1230222113003210-1302101103221220-0001131132130300-2031003120223011-0320021111213133-2113030112031032-1113221220203332-1122131330203202)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.default_route.disable_host_rewrite

<a id="canonical-1212020223010100-2021310102010032-3231321210101001-3222332312210322-2301101102222211-0322013333303310-2321313200123121-2000033111210023"></a>

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
disable_host_rewrite = {}
```

<a id="canonical-2230320233130013-3133000300200002-3303220022111013-2122321102003020-1132131013301012-1000002010132321-3333220312303023-1110220200030301"></a>

## Direct properties — disable_host_rewrite / 221213033230 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0003030120311303-0230023010032220-2000323012320223-1020222131121202-3133101110320333-3010201212122022-1211120112221323-2033130321101310"></a>

## Next pages — disable_host_rewrite / 221213033230 / 4

- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.default_route](resources--workload--reference--group-025.md#canonical-1230222113003210-1302101103221220-0001131132130300-2031003120223011-0320021111213133-2113030112031032-1113221220203332-1122131330203202)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-2123333100103222-2211301302000212-3113312132112323-3231132113021003-3013302213020120-3013201031021330-2313212132310112-0011220113322010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0133322032311003-1001331202010013-0120311010222233-3131223101323030-2132100220330331-0222313313101103-0201021230333301-3121303222020021"></a>

## stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.http — http / 233101300301 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-017.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-021.md#canonical-3300101333202310-3223231311312310-2030111033033203-1011103231310023-1313110120311121-1221031312313020-1210302113302033-2110231331000202)
- [stateful_service.advertise_options.advertise_on_public.port](resources--workload--reference--group-025.md#canonical-0120012201110120-2320030110233132-3233102311321321-2202232022223030-2133131131013111-3231210131302330-1022020011332122-3111230313032011)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-025.md#canonical-2203002110120031-3203013111221013-1321101003110222-2103121303322022-0021002301311230-0203100203111101-0321302220100020-1320220233023212)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.http

<a id="canonical-0322012020210322-2131330111210220-1212032223332211-2332021003123010-3220000310311323-3131130203113311-1120133010130210-3231111001110002"></a>

Type: `"object"`. single nested block, Optional.

HTTP Choice. Choice for selecting HTTP proxy.

Upstream description:

Choice for selecting HTTP proxy.

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

<a id="canonical-3300032223120200-1231311110100310-1212120011213101-2202221110202331-3230203123031133-0033202231010030-1101203130102302-2203010021220311"></a>

## Direct properties — http / 233101300301 / 3

<a id="canonical-1032100330332330-2122013310132312-0012112031012231-2121022320223012-1222011013113123-1103000213203002-2302100231320100-3232221233100333"></a>

<a id="canonical-2223200212310211-3102302231113101-1320123222012000-2013033112002022-1023331320113313-1023202000200102-0333331233120002-0131231332022133"></a>

## dns_volterra_managed property — http / 233101300301 / 4

Type: `"bool"`. Optional.

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

<a id="canonical-3221020121122312-2010012313132010-2312133001132003-1122200030332301-2221123203333313-3333132321130030-0322220031113202-1301320201231022"></a>

<a id="canonical-1001021023323101-3111202331002221-0020332322310220-3313012101210201-3012210211322210-0130131312221311-0200000223033211-0312330021032322"></a>

## port property — http / 233101300301 / 5

Type: `"number"`. Optional.

Exclusive with \[port\_ranges\] HTTP port to Listen.

Upstream description:

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

<a id="canonical-0032301321233203-1213203323021222-1020230031020310-2102330100000033-1231033010310333-1332001033222121-1330132100013300-1222212103303332"></a>

<a id="canonical-3320201311001130-0300200103100222-2201033010101332-2133300002012220-1211232003023330-2003011230113021-3213100211321200-3003021112322313"></a>

## port_ranges property — http / 233101300301 / 6

Type: `"string"`. Optional.

Exclusive with \[port\] A string containing a comma separated list of port ranges. Each port range
consists of a single port or two ports separated by '-'.

Upstream description:

Exclusive with \[port\] A string containing a comma separated list of port ranges. Each port range
consists of a single port or two ports separated by "-".

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

<a id="canonical-2122122030002123-1013311222002312-3320120210302233-3200201203121320-3303102000010020-1033013232031122-1230102122130233-1333030013122223"></a>

## Next pages — http / 233101300301 / 7

- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-025.md#canonical-2203002110120031-3203013111221013-1321101003110222-2103121303322022-0021002301311230-0203100203111101-0321302220100020-1320220233023212)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-2320131312110211-0233203100020221-1102210111313012-2101323032112112-2012023213113210-1213111131330011-1020010101131133-0020231202232230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3120111323212122-0230000321022330-3312321312331212-2113133201201130-0303331213211231-3101301122010132-3113321302122200-0020130302103100"></a>

## stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https — https / 223300032330 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-017.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-021.md#canonical-3300101333202310-3223231311312310-2030111033033203-1011103231310023-1313110120311121-1221031312313020-1210302113302033-2110231331000202)
- [stateful_service.advertise_options.advertise_on_public.port](resources--workload--reference--group-025.md#canonical-0120012201110120-2320030110233132-3233102311321321-2202232022223030-2133131131013111-3231210131302330-1022020011332122-3111230313032011)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-025.md#canonical-2203002110120031-3203013111221013-1321101003110222-2103121303322022-0021002301311230-0203100203111101-0321302220100020-1320220233023212)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https

<a id="canonical-0312002103303103-1230220001210331-1001013131023300-1301200102212032-1310220322121300-1112020131212301-2312311231320131-0200332120311220"></a>

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

<a id="canonical-3230022120331131-0000113312323032-3330032003211030-2122131032330112-3303113010111121-2131001111200232-1123031303202012-2222111033321222"></a>

## Direct properties — https / 223300032330 / 3

<a id="canonical-1031101133001321-0022300211212132-1301221100221312-3111322332131002-3300111100100102-3021113103133001-0022311212033032-3313210020003101"></a>

<a id="canonical-2212012013132001-1022011311002333-0331331323230013-1031101223121133-2330010120323101-2312021103212023-0002011021021122-2231012321312112"></a>

## add_hsts property — https / 223300032330 / 4

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

<a id="canonical-1331101311031131-1120232331202032-1000223201110223-0122230101203320-2213000300223323-1312111211331012-2133101223200312-3120323031332331"></a>

<a id="canonical-1031103112312223-2201122121322013-1032320112312322-1200100321032101-0312220211101300-0021302031332200-2113211111023220-2021113303311311"></a>

## append_server_name property — https / 223300032330 / 5

Type: `"string"`. Optional.

Exclusive with \[default\_header pass\_through server\_name\] Define the header value for the header
name “server”. If header value is already present, it is not overwritten and passed as-is.

Upstream description:

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

- [coalescing_options](resources--workload--reference--group-025.md#canonical-0031220212102213-2023001210113030-3202110102131200-2313322210012222-2301222222313032-0202110131302200-0033121103230333-3321012333011031): complete subsection reference.

<a id="canonical-0033323120113032-2223233001300103-0013311033220111-1132031121203320-3003130132003211-0133103012023333-1233012200302320-0322333323311213"></a>

<a id="canonical-3022130111223311-3133310200110331-2031100123032233-0131103301322323-0110231213211301-1123010200213322-2331032103223022-0333232121233003"></a>

## connection_idle_timeout property — https / 223300032330 / 6

Type: `"number"`. Optional.

The idle timeout for downstream connections. The idle timeout is defined as the period in which
there are no active requests. When the idle timeout is reached the connection will be closed.

Upstream description:

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

- [default_header](resources--workload--reference--group-025.md#canonical-0301223300331201-1112110023010300-1312323213203202-3333312103222100-3201000010002020-1101302100100302-0321330302032101-2201131111132321): complete subsection reference.

- [default_loadbalancer](resources--workload--reference--group-025.md#canonical-3200003201211113-0121013200000330-3321230201201113-0023033221112202-1220201021002321-0230301203312103-1322301133203002-0312221012322010): complete subsection reference.

- [disable_path_normalize](resources--workload--reference--group-025.md#canonical-1012103331231202-0311003123012302-0121213110221111-3332110322300112-3202013213313023-3221131112232033-0021232000032032-3013030030333120): complete subsection reference.

- [enable_path_normalize](resources--workload--reference--group-025.md#canonical-2102000330000002-0320220020130102-2033201033101330-2120200210111302-1010231303001230-1223220311132011-0311113212000000-1232023322331131): complete subsection reference.

- [http_protocol_options](resources--workload--reference--group-025.md#canonical-2032300002331210-2003031033332223-3100202201220210-3023121111333223-0333211003120022-0012320301030022-0220213110021020-2200120102133013): complete subsection reference.

<a id="canonical-2300200230102202-2110322303021200-0022233013130231-3122110232100202-3332120233202022-2101233003132322-0330233312013320-2031303301121333"></a>

<a id="canonical-3200333332020010-2102323323312133-3110022011212021-0220122203010023-3330200221130033-0302333000200001-3132300122321331-1120331113133212"></a>

## http_redirect property — https / 223300032330 / 7

Type: `"bool"`. Optional.

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

- [non_default_loadbalancer](resources--workload--reference--group-025.md#canonical-2011020100023130-1310233221300113-1322210220022022-2221322110003020-3312332110011203-0203230122110211-1201131320233323-1200222013132021): complete subsection reference.

- [pass_through](resources--workload--reference--group-025.md#canonical-1323022111031301-3221111221131322-3020221302312032-0012113322331131-0113211232022110-3220023132120131-3313210220331333-1233013220122322): complete subsection reference.

<a id="canonical-3123213033013233-2323121333010112-3001031000032331-1132302032123032-0001333020112301-2120010102112020-1232220021231122-3230212223031102"></a>

<a id="canonical-0131100023021100-1232131310011230-3120311103333233-1232113031031320-1120123111113223-1311133121313332-0032023023302011-0032030310212133"></a>

## port property — https / 223300032330 / 8

Type: `"number"`. Optional.

Exclusive with \[port\_ranges\] HTTPS port to Listen.

Upstream description:

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

<a id="canonical-1100210331310030-1203112102231210-3000030011212223-1100313032111002-3223131231233103-0111031310112313-0132202123303112-3112123121220012"></a>

<a id="canonical-1101003203011033-1102110212013102-2031000030101313-3020233332110122-0211111113210122-1201131312223112-1031131101100123-3111200322023230"></a>

## port_ranges property — https / 223300032330 / 9

Type: `"string"`. Optional.

Exclusive with \[port\] A string containing a comma separated list of port ranges. Each port range
consists of a single port or two ports separated by '-'.

Upstream description:

Exclusive with \[port\] A string containing a comma separated list of port ranges. Each port range
consists of a single port or two ports separated by "-".

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

<a id="canonical-1213121312130101-0033103023031311-1331323331212123-1302333311022013-1313201310010313-2232102211120321-0123212220232201-2122311302213220"></a>

<a id="canonical-3320213032033121-0313313110032332-0102202011110222-1222233232323033-1220101311011331-2021311000000102-0323313023302113-0121130002102133"></a>

## server_name property — https / 223300032330 / 10

Type: `"string"`. Optional.

Exclusive with \[append\_server\_name default\_header pass\_through\] Define the header value for
the header name “server”. This will overwrite existing values, if any, for the server header.

Upstream description:

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

- [tls_cert_params](resources--workload--reference--group-025.md#canonical-0330120220221323-3001303122321320-1021021120022323-3033111012020201-2201223300221222-2120132123232130-3021131112021012-3210022322020231): complete subsection reference.

- [tls_parameters](resources--workload--reference--group-026.md#canonical-2220233222302121-1211133331002001-2312102020200100-3222223333223312-3022013120330322-2001232101131301-2201322330231121-2322331210022031): complete subsection reference.

<a id="canonical-2022001302132120-0102222111123100-2120032023213302-1002033132310223-1030233222031221-2211022200320000-0333102103331310-3020322321001231"></a>

## Next pages — https / 223300032330 / 11

- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.coalescing_options](resources--workload--reference--group-025.md#canonical-0031220212102213-2023001210113030-3202110102131200-2313322210012222-2301222222313032-0202110131302200-0033121103230333-3321012333011031)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.default_header](resources--workload--reference--group-025.md#canonical-0301223300331201-1112110023010300-1312323213203202-3333312103222100-3201000010002020-1101302100100302-0321330302032101-2201131111132321)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.default_loadbalancer](resources--workload--reference--group-025.md#canonical-3200003201211113-0121013200000330-3321230201201113-0023033221112202-1220201021002321-0230301203312103-1322301133203002-0312221012322010)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.disable_path_normalize](resources--workload--reference--group-025.md#canonical-1012103331231202-0311003123012302-0121213110221111-3332110322300112-3202013213313023-3221131112232033-0021232000032032-3013030030333120)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.enable_path_normalize](resources--workload--reference--group-025.md#canonical-2102000330000002-0320220020130102-2033201033101330-2120200210111302-1010231303001230-1223220311132011-0311113212000000-1232023322331131)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options](resources--workload--reference--group-025.md#canonical-2032300002331210-2003031033332223-3100202201220210-3023121111333223-0333211003120022-0012320301030022-0220213110021020-2200120102133013)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.non_default_loadbalancer](resources--workload--reference--group-025.md#canonical-2011020100023130-1310233221300113-1322210220022022-2221322110003020-3312332110011203-0203230122110211-1201131320233323-1200222013132021)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.pass_through](resources--workload--reference--group-025.md#canonical-1323022111031301-3221111221131322-3020221302312032-0012113322331131-0113211232022110-3220023132120131-3313210220331333-1233013220122322)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params](resources--workload--reference--group-025.md#canonical-0330120220221323-3001303122321320-1021021120022323-3033111012020201-2201223300221222-2120132123232130-3021131112021012-3210022322020231)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-026.md#canonical-2220233222302121-1211133331002001-2312102020200100-3222223333223312-3022013120330322-2001232101131301-2201322330231121-2322331210022031)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-025.md#canonical-2203002110120031-3203013111221013-1321101003110222-2103121303322022-0021002301311230-0203100203111101-0321302220100020-1320220233023212)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-0031220212102213-2023001210113030-3202110102131200-2313322210012222-2301222222313032-0202110131302200-0033121103230333-3321012333011031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3312331323110223-2233313213113031-1010300111121322-3103311322313022-0002103031201111-2013232132030213-1121222330121212-0022031321303120"></a>

## stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.coalescing_options — coalescing_options / 233010202322 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-017.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-021.md#canonical-3300101333202310-3223231311312310-2030111033033203-1011103231310023-1313110120311121-1221031312313020-1210302113302033-2110231331000202)
- [stateful_service.advertise_options.advertise_on_public.port](resources--workload--reference--group-025.md#canonical-0120012201110120-2320030110233132-3233102311321321-2202232022223030-2133131131013111-3231210131302330-1022020011332122-3111230313032011)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-025.md#canonical-2203002110120031-3203013111221013-1321101003110222-2103121303322022-0021002301311230-0203100203111101-0321302220100020-1320220233023212)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https](resources--workload--reference--group-025.md#canonical-2320131312110211-0233203100020221-1102210111313012-2101323032112112-2012023213113210-1213111131330011-1020010101131133-0020231202232230)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.coalescing_options

<a id="canonical-2212212320232333-1211300123310203-0120133332312202-1102312311112310-1211033000300112-2011023100102022-2211120220303130-2122310311210311"></a>

Type: `"object"`. single nested block, Optional.

TLS connection coalescing configuration (not compatible with mTLS).

Upstream description:

TLS connection coalescing configuration (not compatible with mTLS)

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-0201220233203023-0323133312110300-3310123000330312-2002100020110323-2223023301321301-1100102201020021-2230202100311313-2023302020031301"></a>

## Direct properties — coalescing_options / 233010202322 / 3

- [default_coalescing](resources--workload--reference--group-025.md#canonical-0231032213032202-3023222033123332-2220120212233320-3013302310102103-1302113032233133-0003321211021011-3012112321213223-2100000032231010): complete subsection reference.

- [strict_coalescing](resources--workload--reference--group-025.md#canonical-3002310231302102-3113101123310223-0321231212032311-0102313120022333-2112220032301222-3023031222111312-2123321112221122-3022221102231031): complete subsection reference.

<a id="canonical-3330111332320133-0002322103232123-1113001321330233-1322321223203111-2301231101102023-0000123302333002-0230203021021233-2322200221322210"></a>

## Next pages — coalescing_options / 233010202322 / 4

- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.coalescing_options.default_coalescing](resources--workload--reference--group-025.md#canonical-0231032213032202-3023222033123332-2220120212233320-3013302310102103-1302113032233133-0003321211021011-3012112321213223-2100000032231010)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.coalescing_options.strict_coalescing](resources--workload--reference--group-025.md#canonical-3002310231302102-3113101123310223-0321231212032311-0102313120022333-2112220032301222-3023031222111312-2123321112221122-3022221102231031)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https](resources--workload--reference--group-025.md#canonical-2320131312110211-0233203100020221-1102210111313012-2101323032112112-2012023213113210-1213111131330011-1020010101131133-0020231202232230)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-0231032213032202-3023222033123332-2220120212233320-3013302310102103-1302113032233133-0003321211021011-3012112321213223-2100000032231010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2302123312300231-0201123312022110-0310213323220303-2032210012131012-0201322032320113-0300331132312121-0100031233303022-0300022011132022"></a>

## stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.coalescing_options.default_coalescing — default_coalescing / 321331320201 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-017.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-021.md#canonical-3300101333202310-3223231311312310-2030111033033203-1011103231310023-1313110120311121-1221031312313020-1210302113302033-2110231331000202)
- [stateful_service.advertise_options.advertise_on_public.port](resources--workload--reference--group-025.md#canonical-0120012201110120-2320030110233132-3233102311321321-2202232022223030-2133131131013111-3231210131302330-1022020011332122-3111230313032011)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-025.md#canonical-2203002110120031-3203013111221013-1321101003110222-2103121303322022-0021002301311230-0203100203111101-0321302220100020-1320220233023212)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https](resources--workload--reference--group-025.md#canonical-2320131312110211-0233203100020221-1102210111313012-2101323032112112-2012023213113210-1213111131330011-1020010101131133-0020231202232230)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.coalescing_options](resources--workload--reference--group-025.md#canonical-0031220212102213-2023001210113030-3202110102131200-2313322210012222-2301222222313032-0202110131302200-0033121103230333-3321012333011031)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.coalescing_options.default_coalescing

<a id="canonical-1113013213230200-0101120300031312-3101333232231032-0311310330123032-0302120101333331-1311301210220001-2301132322230201-1323221300312221"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
default_coalescing = {}
```

<a id="canonical-1020320221020121-3110032320200212-3203011001202120-3303231210212230-0033010230030120-1001332300111223-3023011222203302-1033112132303220"></a>

## Direct properties — default_coalescing / 321331320201 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3210331121302122-3022100311103200-1031300221003033-2112030200331223-1010003013113302-0302303320112032-2303122032012322-0310110101032123"></a>

## Next pages — default_coalescing / 321331320201 / 4

- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.coalescing_options](resources--workload--reference--group-025.md#canonical-0031220212102213-2023001210113030-3202110102131200-2313322210012222-2301222222313032-0202110131302200-0033121103230333-3321012333011031)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-3002310231302102-3113101123310223-0321231212032311-0102313120022333-2112220032301222-3023031222111312-2123321112221122-3022221102231031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0212203130001022-2231110120301130-1313103323322130-2131211331223011-1331320032230020-2103231203033101-3212333011130031-2100211000023013"></a>

## stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.coalescing_options.strict_coalescing — strict_coalescing / 121130213021 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-017.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-021.md#canonical-3300101333202310-3223231311312310-2030111033033203-1011103231310023-1313110120311121-1221031312313020-1210302113302033-2110231331000202)
- [stateful_service.advertise_options.advertise_on_public.port](resources--workload--reference--group-025.md#canonical-0120012201110120-2320030110233132-3233102311321321-2202232022223030-2133131131013111-3231210131302330-1022020011332122-3111230313032011)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-025.md#canonical-2203002110120031-3203013111221013-1321101003110222-2103121303322022-0021002301311230-0203100203111101-0321302220100020-1320220233023212)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https](resources--workload--reference--group-025.md#canonical-2320131312110211-0233203100020221-1102210111313012-2101323032112112-2012023213113210-1213111131330011-1020010101131133-0020231202232230)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.coalescing_options](resources--workload--reference--group-025.md#canonical-0031220212102213-2023001210113030-3202110102131200-2313322210012222-2301222222313032-0202110131302200-0033121103230333-3321012333011031)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.coalescing_options.strict_coalescing

<a id="canonical-2103310220012230-2113203010132212-2020132221310011-0323201332202013-1001033022201102-0011010130131030-0232000201121020-2110310213220333"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
strict_coalescing = {}
```

<a id="canonical-3220312002222032-3311200330332200-2221233311302011-2332230300211111-3323121002303233-3011003233303102-2323013000133310-0323213100310233"></a>

## Direct properties — strict_coalescing / 121130213021 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3232221102100220-2320220221312002-0033130233033013-0031330133022330-2102020311012233-3101023103210333-1030113120222330-1321033313320332"></a>

## Next pages — strict_coalescing / 121130213021 / 4

- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.coalescing_options](resources--workload--reference--group-025.md#canonical-0031220212102213-2023001210113030-3202110102131200-2313322210012222-2301222222313032-0202110131302200-0033121103230333-3321012333011031)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-0301223300331201-1112110023010300-1312323213203202-3333312103222100-3201000010002020-1101302100100302-0321330302032101-2201131111132321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1321211101132203-0213121331200221-2202110221130120-1301122031232202-1302002331211132-0113222102032022-0222020002121201-1032132000000023"></a>

## stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.default_header — default_header / 222312303132 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-017.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-021.md#canonical-3300101333202310-3223231311312310-2030111033033203-1011103231310023-1313110120311121-1221031312313020-1210302113302033-2110231331000202)
- [stateful_service.advertise_options.advertise_on_public.port](resources--workload--reference--group-025.md#canonical-0120012201110120-2320030110233132-3233102311321321-2202232022223030-2133131131013111-3231210131302330-1022020011332122-3111230313032011)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-025.md#canonical-2203002110120031-3203013111221013-1321101003110222-2103121303322022-0021002301311230-0203100203111101-0321302220100020-1320220233023212)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https](resources--workload--reference--group-025.md#canonical-2320131312110211-0233203100020221-1102210111313012-2101323032112112-2012023213113210-1213111131330011-1020010101131133-0020231202232230)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.default_header

<a id="canonical-0211012120300020-3231231231032130-2230130202002323-0101120021213331-1330010200230300-2222020330100320-3330231002113213-2011222212213303"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
default_header = {}
```

<a id="canonical-0321222201101303-1220301311121232-2220020102121302-2203232100020222-1013212131300101-3232120221012013-3321001032333011-2022333310233031"></a>

## Direct properties — default_header / 222312303132 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0223003001221310-3121203133123023-3332230330310232-0101021233331003-0202022120201033-2102031012110101-2022212020212231-0111231211311102"></a>

## Next pages — default_header / 222312303132 / 4

- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https](resources--workload--reference--group-025.md#canonical-2320131312110211-0233203100020221-1102210111313012-2101323032112112-2012023213113210-1213111131330011-1020010101131133-0020231202232230)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-3200003201211113-0121013200000330-3321230201201113-0023033221112202-1220201021002321-0230301203312103-1322301133203002-0312221012322010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3131022123101021-0021332110312300-0232213232321013-0030221212122111-0012020222111111-1222112330020003-0033233022212123-2132102332133323"></a>

## stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.default_loadbalancer — default_loadbalancer / 110010322230 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-017.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-021.md#canonical-3300101333202310-3223231311312310-2030111033033203-1011103231310023-1313110120311121-1221031312313020-1210302113302033-2110231331000202)
- [stateful_service.advertise_options.advertise_on_public.port](resources--workload--reference--group-025.md#canonical-0120012201110120-2320030110233132-3233102311321321-2202232022223030-2133131131013111-3231210131302330-1022020011332122-3111230313032011)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-025.md#canonical-2203002110120031-3203013111221013-1321101003110222-2103121303322022-0021002301311230-0203100203111101-0321302220100020-1320220233023212)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https](resources--workload--reference--group-025.md#canonical-2320131312110211-0233203100020221-1102210111313012-2101323032112112-2012023213113210-1213111131330011-1020010101131133-0020231202232230)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.default_loadbalancer

<a id="canonical-3232100033113322-1110311021333111-0023313001131020-3033303022202021-0233302031122113-2102203331210132-2131220223000110-2000321323321230"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
default_loadbalancer = {}
```

<a id="canonical-3131232032033003-2111303202330230-1201122233330312-2203323013100010-2002121011333001-2022131032021002-2313312231310233-1310233213210223"></a>

## Direct properties — default_loadbalancer / 110010322230 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0211010122113023-0132213003301212-2111311000200122-2310233322323313-3023121332320123-2003323310020132-2300220001231332-0130201111320331"></a>

## Next pages — default_loadbalancer / 110010322230 / 4

- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https](resources--workload--reference--group-025.md#canonical-2320131312110211-0233203100020221-1102210111313012-2101323032112112-2012023213113210-1213111131330011-1020010101131133-0020231202232230)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-1012103331231202-0311003123012302-0121213110221111-3332110322300112-3202013213313023-3221131112232033-0021232000032032-3013030030333120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2322100231032020-1121312132113331-0232132210123000-0031302110012213-1232023111003120-2121031031113133-1003110322123000-0333320222303002"></a>

## stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.disable_path_normalize — disable_path_normalize / 002323320322 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-017.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-021.md#canonical-3300101333202310-3223231311312310-2030111033033203-1011103231310023-1313110120311121-1221031312313020-1210302113302033-2110231331000202)
- [stateful_service.advertise_options.advertise_on_public.port](resources--workload--reference--group-025.md#canonical-0120012201110120-2320030110233132-3233102311321321-2202232022223030-2133131131013111-3231210131302330-1022020011332122-3111230313032011)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-025.md#canonical-2203002110120031-3203013111221013-1321101003110222-2103121303322022-0021002301311230-0203100203111101-0321302220100020-1320220233023212)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https](resources--workload--reference--group-025.md#canonical-2320131312110211-0233203100020221-1102210111313012-2101323032112112-2012023213113210-1213111131330011-1020010101131133-0020231202232230)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.disable_path_normalize

<a id="canonical-3110230302202120-0312201032231132-3321200313223122-3231111203311312-3103320220100303-2031130320231231-1010333313103232-3202120212033323"></a>

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
disable_path_normalize = {}
```

<a id="canonical-3012113221201131-0133121100002001-3122203213200020-3212123020103120-0110120111200133-2213013013222312-0103121021120011-2010132110020203"></a>

## Direct properties — disable_path_normalize / 002323320322 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1132311323311313-0103013023020330-1213111023103023-2301323132323020-1201113310201233-1223130330011011-0312212213301022-2230222131030220"></a>

## Next pages — disable_path_normalize / 002323320322 / 4

- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https](resources--workload--reference--group-025.md#canonical-2320131312110211-0233203100020221-1102210111313012-2101323032112112-2012023213113210-1213111131330011-1020010101131133-0020231202232230)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-2102000330000002-0320220020130102-2033201033101330-2120200210111302-1010231303001230-1223220311132011-0311113212000000-1232023322331131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3000133131020112-0233030032002010-3001330022110130-0333120200321200-3233323310110330-1001000221113233-3012223321002122-3020011221123111"></a>

## stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.enable_path_normalize — enable_path_normalize / 311212233100 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-017.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-021.md#canonical-3300101333202310-3223231311312310-2030111033033203-1011103231310023-1313110120311121-1221031312313020-1210302113302033-2110231331000202)
- [stateful_service.advertise_options.advertise_on_public.port](resources--workload--reference--group-025.md#canonical-0120012201110120-2320030110233132-3233102311321321-2202232022223030-2133131131013111-3231210131302330-1022020011332122-3111230313032011)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-025.md#canonical-2203002110120031-3203013111221013-1321101003110222-2103121303322022-0021002301311230-0203100203111101-0321302220100020-1320220233023212)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https](resources--workload--reference--group-025.md#canonical-2320131312110211-0233203100020221-1102210111313012-2101323032112112-2012023213113210-1213111131330011-1020010101131133-0020231202232230)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.enable_path_normalize

<a id="canonical-0312012222011211-0121121331301113-1321221031330120-3302110233230232-0122032302203201-0313332232103310-0311000222232102-1123211212311323"></a>

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
enable_path_normalize = {}
```

<a id="canonical-0030330330121120-3302311330010000-1103332303001100-0112022222131000-1021112313303320-3122101310321003-3103320101202000-1301000200112200"></a>

## Direct properties — enable_path_normalize / 311212233100 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3313212201232300-0131020132331101-0030032301331132-0233233203030300-0211132020233132-3300223322112121-2210011300123331-0231131010111031"></a>

## Next pages — enable_path_normalize / 311212233100 / 4

- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https](resources--workload--reference--group-025.md#canonical-2320131312110211-0233203100020221-1102210111313012-2101323032112112-2012023213113210-1213111131330011-1020010101131133-0020231202232230)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-2032300002331210-2003031033332223-3100202201220210-3023121111333223-0333211003120022-0012320301030022-0220213110021020-2200120102133013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2022332122201032-3133311301301032-3132032100110213-2033002203301123-2200223202022020-0313110103321121-2310112233321023-2330101332323303"></a>

## stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options — http_protocol_options / 130102000221 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-017.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-021.md#canonical-3300101333202310-3223231311312310-2030111033033203-1011103231310023-1313110120311121-1221031312313020-1210302113302033-2110231331000202)
- [stateful_service.advertise_options.advertise_on_public.port](resources--workload--reference--group-025.md#canonical-0120012201110120-2320030110233132-3233102311321321-2202232022223030-2133131131013111-3231210131302330-1022020011332122-3111230313032011)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-025.md#canonical-2203002110120031-3203013111221013-1321101003110222-2103121303322022-0021002301311230-0203100203111101-0321302220100020-1320220233023212)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https](resources--workload--reference--group-025.md#canonical-2320131312110211-0233203100020221-1102210111313012-2101323032112112-2012023213113210-1213111131330011-1020010101131133-0020231202232230)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options

<a id="canonical-0011222003332132-2101321302211311-2012220113023031-3111300111200202-0023131132232130-1001100121001320-1312332033211122-1313213221110112"></a>

Type: `"object"`. single nested block, Optional.

HTTP protocol configuration OPTIONS for downstream connections.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("http_protocol_enable_v1_only",
    "http_protocol_enable_v1_v2"),
  validators.ConflictingObjectAttributes("http_protocol_enable_v1_only",
    "http_protocol_enable_v2_only"),
  validators.ConflictingObjectAttributes("http_protocol_enable_v1_v2",
    "http_protocol_enable_v2_only")}
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
  "x-ves-oneof-field-http_protocol_choice": "[\"http_protocol_enable_v1_only\",\"http_protocol_enable_v1_v2\",\"http_protocol_enable_v2_only\"]"
}
```

Terraform syntax:

```terraform
http_protocol_options {
  # Configure direct properties listed below.
}
```

<a id="canonical-0110131301003101-1032002220111110-0032313010122013-3111003002010032-1233120300202030-0103212032030032-0100102030330303-1112210113312013"></a>

## Direct properties — http_protocol_options / 130102000221 / 3

- [http_protocol_enable_v1_only](resources--workload--reference--group-025.md#canonical-2102101211002103-0300031120110111-1032032002310002-1333323130021301-3213223120022023-1323101210321110-3003121232011113-2220222031321111): complete subsection reference.

- [http_protocol_enable_v1_v2](resources--workload--reference--group-025.md#canonical-0111101001300011-2231203132330302-1313203121323310-3111210202230200-1302333130201212-0310200320232221-1022021022222020-1123121212030001): complete subsection reference.

- [http_protocol_enable_v2_only](resources--workload--reference--group-025.md#canonical-3333011020331002-3013313200123020-0213321221322000-1000020020313033-3223013300232100-0100110201022223-1231132313131321-0222302013301102): complete subsection reference.

<a id="canonical-0311331001230132-3313030223020301-3113220203213221-0212320020200021-0323012003330130-3133003300122320-1021120020323133-3322333000210311"></a>

## Next pages — http_protocol_options / 130102000221 / 4

- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only](resources--workload--reference--group-025.md#canonical-2102101211002103-0300031120110111-1032032002310002-1333323130021301-3213223120022023-1323101210321110-3003121232011113-2220222031321111)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_v2](resources--workload--reference--group-025.md#canonical-0111101001300011-2231203132330302-1313203121323310-3111210202230200-1302333130201212-0310200320232221-1022021022222020-1123121212030001)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v2_only](resources--workload--reference--group-025.md#canonical-3333011020331002-3013313200123020-0213321221322000-1000020020313033-3223013300232100-0100110201022223-1231132313131321-0222302013301102)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https](resources--workload--reference--group-025.md#canonical-2320131312110211-0233203100020221-1102210111313012-2101323032112112-2012023213113210-1213111131330011-1020010101131133-0020231202232230)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-2102101211002103-0300031120110111-1032032002310002-1333323130021301-3213223120022023-1323101210321110-3003121232011113-2220222031321111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0330000111012231-0001111133113122-0122301130130322-3231321012322022-2220302303131320-1122221111332300-2301322322303333-2023331201211220"></a>

## stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only — http_protocol_enable_v1_only / 333100313301 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-017.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-021.md#canonical-3300101333202310-3223231311312310-2030111033033203-1011103231310023-1313110120311121-1221031312313020-1210302113302033-2110231331000202)
- [stateful_service.advertise_options.advertise_on_public.port](resources--workload--reference--group-025.md#canonical-0120012201110120-2320030110233132-3233102311321321-2202232022223030-2133131131013111-3231210131302330-1022020011332122-3111230313032011)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-025.md#canonical-2203002110120031-3203013111221013-1321101003110222-2103121303322022-0021002301311230-0203100203111101-0321302220100020-1320220233023212)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https](resources--workload--reference--group-025.md#canonical-2320131312110211-0233203100020221-1102210111313012-2101323032112112-2012023213113210-1213111131330011-1020010101131133-0020231202232230)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options](resources--workload--reference--group-025.md#canonical-2032300002331210-2003031033332223-3100202201220210-3023121111333223-0333211003120022-0012320301030022-0220213110021020-2200120102133013)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only

<a id="canonical-2113313031021003-3302022032021021-3330332221312221-3010100003230300-3220020131133110-2120211220122201-0111103130022120-1331103200321012"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
http_protocol_enable_v1_only {
  # Configure direct properties listed below.
}
```

<a id="canonical-2211312223330111-1011231110123233-3113232002223333-0031203000131031-1021221333303232-1302023131233312-2300011212200332-3303221032220022"></a>

## Direct properties — http_protocol_enable_v1_only / 333100313301 / 3

- [header_transformation](resources--workload--reference--group-025.md#canonical-3102223033322310-1021310012201002-0131002033213211-3121032313323011-0321321130322221-3033031331010313-2221331122123131-0211313032102103): complete subsection reference.

<a id="canonical-0301233211302100-2022203110330212-3221313001110030-1022000302220212-1021300000020331-3131303000122232-1020212331220123-3330221133133123"></a>

## Next pages — http_protocol_enable_v1_only / 333100313301 / 4

- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation](resources--workload--reference--group-025.md#canonical-3102223033322310-1021310012201002-0131002033213211-3121032313323011-0321321130322221-3033031331010313-2221331122123131-0211313032102103)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options](resources--workload--reference--group-025.md#canonical-2032300002331210-2003031033332223-3100202201220210-3023121111333223-0333211003120022-0012320301030022-0220213110021020-2200120102133013)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-3102223033322310-1021310012201002-0131002033213211-3121032313323011-0321321130322221-3033031331010313-2221331122123131-0211313032102103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3313221102031121-3120132222122012-2220120301201300-0212211133202101-3112102332222123-2012310112132311-3020123010111013-0230230033000102"></a>

## stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation — header_transformation / 310320121101 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-017.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-021.md#canonical-3300101333202310-3223231311312310-2030111033033203-1011103231310023-1313110120311121-1221031312313020-1210302113302033-2110231331000202)
- [stateful_service.advertise_options.advertise_on_public.port](resources--workload--reference--group-025.md#canonical-0120012201110120-2320030110233132-3233102311321321-2202232022223030-2133131131013111-3231210131302330-1022020011332122-3111230313032011)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-025.md#canonical-2203002110120031-3203013111221013-1321101003110222-2103121303322022-0021002301311230-0203100203111101-0321302220100020-1320220233023212)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https](resources--workload--reference--group-025.md#canonical-2320131312110211-0233203100020221-1102210111313012-2101323032112112-2012023213113210-1213111131330011-1020010101131133-0020231202232230)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options](resources--workload--reference--group-025.md#canonical-2032300002331210-2003031033332223-3100202201220210-3023121111333223-0333211003120022-0012320301030022-0220213110021020-2200120102133013)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only](resources--workload--reference--group-025.md#canonical-2102101211002103-0300031120110111-1032032002310002-1333323130021301-3213223120022023-1323101210321110-3003121232011113-2220222031321111)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation

<a id="canonical-0003030220231033-0000231021011333-1320000213130101-3231133132033020-2311302333020201-0233001302201021-2231321330230210-0231323101030322"></a>

Type: `"object"`. single nested block, Optional.

Header Transformation OPTIONS for HTTP/1.1 request/response headers.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("default_header_transformation",
    "preserve_case_header_transformation"),
  validators.ConflictingObjectAttributes("default_header_transformation",
    "proper_case_header_transformation"),
  validators.ConflictingObjectAttributes("preserve_case_header_transformation",
    "proper_case_header_transformation")}
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
  "x-ves-oneof-field-header_transformation_choice": "[\"default_header_transformation\",\"preserve_case_header_transformation\",\"proper_case_header_transformation\"]"
}
```

Terraform syntax:

```terraform
header_transformation {
  # Configure direct properties listed below.
}
```

<a id="canonical-3230033123102001-0220210331200001-3221133122010211-0321313131021200-3232110033032231-3231010121001301-3211020231032122-2233021211200231"></a>

## Direct properties — header_transformation / 310320121101 / 3

- [default_header_transformation](resources--workload--reference--group-025.md#canonical-1321333132133201-3223311133333111-0232220030030200-0131032003022200-2310001112131233-1202132132001310-1203110210010332-2300303320031000): complete subsection reference.

- [preserve_case_header_transformation](resources--workload--reference--group-025.md#canonical-3022022121123222-1321130222221012-0222231300031321-0020022011021010-2332112330202130-2233120002222021-3120201113322101-1022010123212113): complete subsection reference.

- [proper_case_header_transformation](resources--workload--reference--group-025.md#canonical-0213010103031001-2123220203310031-3100010031213012-3011023033112201-2131021012113230-2313202123233030-1233030332131100-1101320201303112): complete subsection reference.

<a id="canonical-3100032022122222-3301322231033131-0330012003113203-2202202313313103-3112232311032131-1011323102221213-0321013212310200-1302110303032100"></a>

## Next pages — header_transformation / 310320121101 / 4

- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.default_header_transformation](resources--workload--reference--group-025.md#canonical-1321333132133201-3223311133333111-0232220030030200-0131032003022200-2310001112131233-1202132132001310-1203110210010332-2300303320031000)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.preserve_case_header_transformation](resources--workload--reference--group-025.md#canonical-3022022121123222-1321130222221012-0222231300031321-0020022011021010-2332112330202130-2233120002222021-3120201113322101-1022010123212113)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.proper_case_header_transformation](resources--workload--reference--group-025.md#canonical-0213010103031001-2123220203310031-3100010031213012-3011023033112201-2131021012113230-2313202123233030-1233030332131100-1101320201303112)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only](resources--workload--reference--group-025.md#canonical-2102101211002103-0300031120110111-1032032002310002-1333323130021301-3213223120022023-1323101210321110-3003121232011113-2220222031321111)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-1321333132133201-3223311133333111-0232220030030200-0131032003022200-2310001112131233-1202132132001310-1203110210010332-2300303320031000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1210133020301030-1012133121130213-3130323002023200-2003330200020121-3213320013033022-2032320000023231-2021331212121302-0122321310002301"></a>

## stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.default_header_transformation — default_header_transformation / 133201211023 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-017.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-021.md#canonical-3300101333202310-3223231311312310-2030111033033203-1011103231310023-1313110120311121-1221031312313020-1210302113302033-2110231331000202)
- [stateful_service.advertise_options.advertise_on_public.port](resources--workload--reference--group-025.md#canonical-0120012201110120-2320030110233132-3233102311321321-2202232022223030-2133131131013111-3231210131302330-1022020011332122-3111230313032011)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-025.md#canonical-2203002110120031-3203013111221013-1321101003110222-2103121303322022-0021002301311230-0203100203111101-0321302220100020-1320220233023212)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https](resources--workload--reference--group-025.md#canonical-2320131312110211-0233203100020221-1102210111313012-2101323032112112-2012023213113210-1213111131330011-1020010101131133-0020231202232230)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options](resources--workload--reference--group-025.md#canonical-2032300002331210-2003031033332223-3100202201220210-3023121111333223-0333211003120022-0012320301030022-0220213110021020-2200120102133013)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only](resources--workload--reference--group-025.md#canonical-2102101211002103-0300031120110111-1032032002310002-1333323130021301-3213223120022023-1323101210321110-3003121232011113-2220222031321111)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation](resources--workload--reference--group-025.md#canonical-3102223033322310-1021310012201002-0131002033213211-3121032313323011-0321321130322221-3033031331010313-2221331122123131-0211313032102103)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.default_header_transformation

<a id="canonical-3323000230332231-2321303213213310-3311033123022303-0212120312031120-0033203103212302-1122222102312310-3321113120131202-2010311020201302"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
default_header_transformation = {}
```

<a id="canonical-1303321330002031-2232300210131221-2122013013201011-2111000130221100-1211012123332212-0102023200221002-2331103220233220-2203111210321103"></a>

## Direct properties — default_header_transformation / 133201211023 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3232333223311011-1003001331332231-1203022123323311-2122002232221313-1112302211022223-0032131131213031-2311021122201033-3331201331101313"></a>

## Next pages — default_header_transformation / 133201211023 / 4

- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation](resources--workload--reference--group-025.md#canonical-3102223033322310-1021310012201002-0131002033213211-3121032313323011-0321321130322221-3033031331010313-2221331122123131-0211313032102103)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-3022022121123222-1321130222221012-0222231300031321-0020022011021010-2332112330202130-2233120002222021-3120201113322101-1022010123212113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2230111303220130-0303323123031301-0132021302322221-3202302213301011-0331220301333330-3130100331222201-2031030222001100-2222031013020313"></a>

## stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.preserve_case_header_transformation — preserve_case_header_transformation / 003232313200 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-017.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-021.md#canonical-3300101333202310-3223231311312310-2030111033033203-1011103231310023-1313110120311121-1221031312313020-1210302113302033-2110231331000202)
- [stateful_service.advertise_options.advertise_on_public.port](resources--workload--reference--group-025.md#canonical-0120012201110120-2320030110233132-3233102311321321-2202232022223030-2133131131013111-3231210131302330-1022020011332122-3111230313032011)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-025.md#canonical-2203002110120031-3203013111221013-1321101003110222-2103121303322022-0021002301311230-0203100203111101-0321302220100020-1320220233023212)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https](resources--workload--reference--group-025.md#canonical-2320131312110211-0233203100020221-1102210111313012-2101323032112112-2012023213113210-1213111131330011-1020010101131133-0020231202232230)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options](resources--workload--reference--group-025.md#canonical-2032300002331210-2003031033332223-3100202201220210-3023121111333223-0333211003120022-0012320301030022-0220213110021020-2200120102133013)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only](resources--workload--reference--group-025.md#canonical-2102101211002103-0300031120110111-1032032002310002-1333323130021301-3213223120022023-1323101210321110-3003121232011113-2220222031321111)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation](resources--workload--reference--group-025.md#canonical-3102223033322310-1021310012201002-0131002033213211-3121032313323011-0321321130322221-3033031331010313-2221331122123131-0211313032102103)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.preserve_case_header_transformation

<a id="canonical-1032203221320321-0302132233023101-1032232101111121-0121310023021323-0223022303031013-2121200223300011-3021030321020332-2320301201300302"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
preserve_case_header_transformation = {}
```

<a id="canonical-3332110201302032-0223130300320101-3130010230302000-0121002030221032-2331311000011131-3131331212102322-2300211212221330-3022311323123230"></a>

## Direct properties — preserve_case_header_transformation / 003232313200 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0202110031200021-2123210333002210-3230012202300211-3130030221032200-3221100110322000-1311110300321221-1002001023231222-2212120311213022"></a>

## Next pages — preserve_case_header_transformation / 003232313200 / 4

- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation](resources--workload--reference--group-025.md#canonical-3102223033322310-1021310012201002-0131002033213211-3121032313323011-0321321130322221-3033031331010313-2221331122123131-0211313032102103)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-0213010103031001-2123220203310031-3100010031213012-3011023033112201-2131021012113230-2313202123233030-1233030332131100-1101320201303112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1333113223322101-0020121203010021-3301100132013203-2031322233021232-0011121213201212-0021211221221313-1201033311312300-0333220332133133"></a>

## stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.proper_case_header_transformation — proper_case_header_transformation / 001322321102 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-017.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-021.md#canonical-3300101333202310-3223231311312310-2030111033033203-1011103231310023-1313110120311121-1221031312313020-1210302113302033-2110231331000202)
- [stateful_service.advertise_options.advertise_on_public.port](resources--workload--reference--group-025.md#canonical-0120012201110120-2320030110233132-3233102311321321-2202232022223030-2133131131013111-3231210131302330-1022020011332122-3111230313032011)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-025.md#canonical-2203002110120031-3203013111221013-1321101003110222-2103121303322022-0021002301311230-0203100203111101-0321302220100020-1320220233023212)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https](resources--workload--reference--group-025.md#canonical-2320131312110211-0233203100020221-1102210111313012-2101323032112112-2012023213113210-1213111131330011-1020010101131133-0020231202232230)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options](resources--workload--reference--group-025.md#canonical-2032300002331210-2003031033332223-3100202201220210-3023121111333223-0333211003120022-0012320301030022-0220213110021020-2200120102133013)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only](resources--workload--reference--group-025.md#canonical-2102101211002103-0300031120110111-1032032002310002-1333323130021301-3213223120022023-1323101210321110-3003121232011113-2220222031321111)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation](resources--workload--reference--group-025.md#canonical-3102223033322310-1021310012201002-0131002033213211-3121032313323011-0321321130322221-3033031331010313-2221331122123131-0211313032102103)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.proper_case_header_transformation

<a id="canonical-2103331033110130-0133321230032010-0130013221112231-3233111100122102-1133002131033221-3120020333331103-0122201113012301-0231322210020101"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
proper_case_header_transformation = {}
```

<a id="canonical-0122301100103223-1013221201000231-1013131302020023-2100100311323210-3131312122032131-0231201010313303-0113123331131203-2202333331220211"></a>

## Direct properties — proper_case_header_transformation / 001322321102 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3100332300023113-2010113031113323-3331023321023222-1012213001012030-0112130000311211-0200222030300321-0100311002013230-0232230122000032"></a>

## Next pages — proper_case_header_transformation / 001322321102 / 4

- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation](resources--workload--reference--group-025.md#canonical-3102223033322310-1021310012201002-0131002033213211-3121032313323011-0321321130322221-3033031331010313-2221331122123131-0211313032102103)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-0111101001300011-2231203132330302-1313203121323310-3111210202230200-1302333130201212-0310200320232221-1022021022222020-1123121212030001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2030121311003101-0302113330231200-2133220123310131-3300301030032101-2120020310010331-2130020100130132-0332112210320303-3113111122231123"></a>

## stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_v2 — http_protocol_enable_v1_v2 / 131331023232 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-017.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-021.md#canonical-3300101333202310-3223231311312310-2030111033033203-1011103231310023-1313110120311121-1221031312313020-1210302113302033-2110231331000202)
- [stateful_service.advertise_options.advertise_on_public.port](resources--workload--reference--group-025.md#canonical-0120012201110120-2320030110233132-3233102311321321-2202232022223030-2133131131013111-3231210131302330-1022020011332122-3111230313032011)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-025.md#canonical-2203002110120031-3203013111221013-1321101003110222-2103121303322022-0021002301311230-0203100203111101-0321302220100020-1320220233023212)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https](resources--workload--reference--group-025.md#canonical-2320131312110211-0233203100020221-1102210111313012-2101323032112112-2012023213113210-1213111131330011-1020010101131133-0020231202232230)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options](resources--workload--reference--group-025.md#canonical-2032300002331210-2003031033332223-3100202201220210-3023121111333223-0333211003120022-0012320301030022-0220213110021020-2200120102133013)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_v2

<a id="canonical-2111121330312302-3002133030101312-1233331032011220-0012322133202032-2332213102002023-2303301330133333-3132211320022220-3131222022030322"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for http protocol enable v1 v2.

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
http_protocol_enable_v1_v2 = {}
```

<a id="canonical-2323111221213123-1331113311031201-3111332213013200-3011220203333323-0323310123030202-0121001121031122-2102203310100101-3002210312132200"></a>

## Direct properties — http_protocol_enable_v1_v2 / 131331023232 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1213201120202032-1313200322102123-0333301012220031-1120101032302011-2100001200101330-3121002312123213-0003103011212201-1222311120202103"></a>

## Next pages — http_protocol_enable_v1_v2 / 131331023232 / 4

- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options](resources--workload--reference--group-025.md#canonical-2032300002331210-2003031033332223-3100202201220210-3023121111333223-0333211003120022-0012320301030022-0220213110021020-2200120102133013)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-3333011020331002-3013313200123020-0213321221322000-1000020020313033-3223013300232100-0100110201022223-1231132313131321-0222302013301102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1100000113220020-0300330313202031-3312231312020113-3131130012021310-3132002102132332-1221200210312331-2113211102131011-3032112120131332"></a>

## stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v2_only — http_protocol_enable_v2_only / 121012022011 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-017.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-021.md#canonical-3300101333202310-3223231311312310-2030111033033203-1011103231310023-1313110120311121-1221031312313020-1210302113302033-2110231331000202)
- [stateful_service.advertise_options.advertise_on_public.port](resources--workload--reference--group-025.md#canonical-0120012201110120-2320030110233132-3233102311321321-2202232022223030-2133131131013111-3231210131302330-1022020011332122-3111230313032011)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-025.md#canonical-2203002110120031-3203013111221013-1321101003110222-2103121303322022-0021002301311230-0203100203111101-0321302220100020-1320220233023212)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https](resources--workload--reference--group-025.md#canonical-2320131312110211-0233203100020221-1102210111313012-2101323032112112-2012023213113210-1213111131330011-1020010101131133-0020231202232230)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options](resources--workload--reference--group-025.md#canonical-2032300002331210-2003031033332223-3100202201220210-3023121111333223-0333211003120022-0012320301030022-0220213110021020-2200120102133013)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v2_only

<a id="canonical-0121102332130301-2231210003213312-0202201212130023-1001212113023133-2200011313131131-0131023331020002-2110210212221113-1201310030001230"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for http protocol enable v2 only.

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
http_protocol_enable_v2_only = {}
```

<a id="canonical-1121212032031202-1320230111302022-3023021101002003-3031331332020303-2010231111302102-2031111000120131-1200221112102332-0313030103300112"></a>

## Direct properties — http_protocol_enable_v2_only / 121012022011 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3112111220321300-0311000212130112-0102321303032322-3112112230202330-3332233002203121-1031332132110001-3032133232331332-2021202010002130"></a>

## Next pages — http_protocol_enable_v2_only / 121012022011 / 4

- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options](resources--workload--reference--group-025.md#canonical-2032300002331210-2003031033332223-3100202201220210-3023121111333223-0333211003120022-0012320301030022-0220213110021020-2200120102133013)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-2011020100023130-1310233221300113-1322210220022022-2221322110003020-3312332110011203-0203230122110211-1201131320233323-1200222013132021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0102321133320311-2033221332121023-2000220113312213-2001002110300213-0013003222023201-0330222003021121-0321133301221332-1103133211010230"></a>

## stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.non_default_loadbalancer — non_default_loadbalancer / 323301111311 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-017.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-021.md#canonical-3300101333202310-3223231311312310-2030111033033203-1011103231310023-1313110120311121-1221031312313020-1210302113302033-2110231331000202)
- [stateful_service.advertise_options.advertise_on_public.port](resources--workload--reference--group-025.md#canonical-0120012201110120-2320030110233132-3233102311321321-2202232022223030-2133131131013111-3231210131302330-1022020011332122-3111230313032011)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-025.md#canonical-2203002110120031-3203013111221013-1321101003110222-2103121303322022-0021002301311230-0203100203111101-0321302220100020-1320220233023212)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https](resources--workload--reference--group-025.md#canonical-2320131312110211-0233203100020221-1102210111313012-2101323032112112-2012023213113210-1213111131330011-1020010101131133-0020231202232230)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.non_default_loadbalancer

<a id="canonical-1331220312001122-1103212311001311-0032123112202331-3231132132313200-2130012130121203-2122311121210233-1131333130310212-2033022330332311"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for non default loadbalancer.

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
non_default_loadbalancer = {}
```

<a id="canonical-0001321130113122-3103222223002221-2223301321113112-0213231131312220-1331110131211212-0213300202020323-2003311011310023-3232022030111133"></a>

## Direct properties — non_default_loadbalancer / 323301111311 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2321312101202201-2122130130211121-0203220211002212-0110100300113231-3211223212211101-2212230230002212-2322021111230111-1120313331111313"></a>

## Next pages — non_default_loadbalancer / 323301111311 / 4

- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https](resources--workload--reference--group-025.md#canonical-2320131312110211-0233203100020221-1102210111313012-2101323032112112-2012023213113210-1213111131330011-1020010101131133-0020231202232230)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-1323022111031301-3221111221131322-3020221302312032-0012113322331131-0113211232022110-3220023132120131-3313210220331333-1233013220122322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1200230132120102-3332100022203012-1103113212000310-2200212001200212-2002302210022022-2103203113230012-1032331213333213-1112322032332133"></a>

## stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.pass_through — pass_through / 202003112120 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-017.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-021.md#canonical-3300101333202310-3223231311312310-2030111033033203-1011103231310023-1313110120311121-1221031312313020-1210302113302033-2110231331000202)
- [stateful_service.advertise_options.advertise_on_public.port](resources--workload--reference--group-025.md#canonical-0120012201110120-2320030110233132-3233102311321321-2202232022223030-2133131131013111-3231210131302330-1022020011332122-3111230313032011)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-025.md#canonical-2203002110120031-3203013111221013-1321101003110222-2103121303322022-0021002301311230-0203100203111101-0321302220100020-1320220233023212)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https](resources--workload--reference--group-025.md#canonical-2320131312110211-0233203100020221-1102210111313012-2101323032112112-2012023213113210-1213111131330011-1020010101131133-0020231202232230)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.pass_through

<a id="canonical-0301112220031103-2101211332011002-0002330032210110-2030301333001203-1200313031031230-2321101003332313-1103312023112330-3321031320111221"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for pass through.

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
pass_through = {}
```

<a id="canonical-0330030202022221-3132313032003123-2123122030112120-1120301333223010-1201222120111211-0233002331133022-1131213031312223-1033032031132123"></a>

## Direct properties — pass_through / 202003112120 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0221200203121113-0302332213221000-1111131210210102-0133010231121233-1122322323201210-3003331200001112-1322312030332323-3202330301210300"></a>

## Next pages — pass_through / 202003112120 / 4

- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https](resources--workload--reference--group-025.md#canonical-2320131312110211-0233203100020221-1102210111313012-2101323032112112-2012023213113210-1213111131330011-1020010101131133-0020231202232230)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-0330120220221323-3001303122321320-1021021120022323-3033111012020201-2201223300221222-2120132123232130-3021131112021012-3210022322020231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1132201131201221-1003331331012123-2211001203122312-0111033333330113-1223001022312220-1230032120022121-3111021131201022-1130321112333031"></a>

## stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params — tls_cert_params / 021022030032 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-017.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-021.md#canonical-3300101333202310-3223231311312310-2030111033033203-1011103231310023-1313110120311121-1221031312313020-1210302113302033-2110231331000202)
- [stateful_service.advertise_options.advertise_on_public.port](resources--workload--reference--group-025.md#canonical-0120012201110120-2320030110233132-3233102311321321-2202232022223030-2133131131013111-3231210131302330-1022020011332122-3111230313032011)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-025.md#canonical-2203002110120031-3203013111221013-1321101003110222-2103121303322022-0021002301311230-0203100203111101-0321302220100020-1320220233023212)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https](resources--workload--reference--group-025.md#canonical-2320131312110211-0233203100020221-1102210111313012-2101323032112112-2012023213113210-1213111131330011-1020010101131133-0020231202232230)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params

<a id="canonical-0232201331203310-2122202320032030-0223310010112021-0332133222010032-2221123102300201-0232102233220203-3000111320332100-0030002013203202"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for tls cert params.

Upstream description:

Select TLS Parameters and Certificates.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("certificates"),
  validators.ConflictingObjectAttributes("no_mtls",
    "use_mtls")}
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
  "x-ves-oneof-field-mtls_choice": "[\"no_mtls\",\"use_mtls\"]"
}
```

Terraform syntax:

```terraform
tls_cert_params {
  # Configure direct properties listed below.
}
```

<a id="canonical-2202101301223111-3310100101022132-0201120102213110-2222202101121020-3013321102323002-3123311233323122-0032133021332000-0110303203220303"></a>

## Direct properties — tls_cert_params / 021022030032 / 3

- [certificates](resources--workload--reference--group-025.md#canonical-1320320332332013-3000002322310202-0212222002210131-1200300012013103-1132121032133313-2310321030202102-2101301330012220-1211233231212313): complete subsection reference.

- [no_mtls](resources--workload--reference--group-025.md#canonical-3311301002213232-3032212233333300-1301201000110122-1211001010000011-2233001000223302-1111201011313003-0233200311222201-2203120030313212): complete subsection reference.

- [tls_config](resources--workload--reference--group-025.md#canonical-3011322100021322-2121021022310000-2011022031301301-3323233310022131-0113312101320323-0211100303013221-0023223321310100-2133210222331011): complete subsection reference.

- [use_mtls](resources--workload--reference--group-025.md#canonical-1023122310132111-2212330202030310-2112313020200023-3101321202230303-1220121020221312-1210310302131011-2132302200011221-1002023020300223): complete subsection reference.

<a id="canonical-3332103301023030-0321300201130020-2101221311303102-0320001220023221-0100231330111123-2233013222201112-2202223312312031-0110001021221213"></a>

## Next pages — tls_cert_params / 021022030032 / 4

- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.certificates](resources--workload--reference--group-025.md#canonical-1320320332332013-3000002322310202-0212222002210131-1200300012013103-1132121032133313-2310321030202102-2101301330012220-1211233231212313)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.no_mtls](resources--workload--reference--group-025.md#canonical-3311301002213232-3032212233333300-1301201000110122-1211001010000011-2233001000223302-1111201011313003-0233200311222201-2203120030313212)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.tls_config](resources--workload--reference--group-025.md#canonical-3011322100021322-2121021022310000-2011022031301301-3323233310022131-0113312101320323-0211100303013221-0023223321310100-2133210222331011)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.use_mtls](resources--workload--reference--group-025.md#canonical-1023122310132111-2212330202030310-2112313020200023-3101321202230303-1220121020221312-1210310302131011-2132302200011221-1002023020300223)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https](resources--workload--reference--group-025.md#canonical-2320131312110211-0233203100020221-1102210111313012-2101323032112112-2012023213113210-1213111131330011-1020010101131133-0020231202232230)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-1320320332332013-3000002322310202-0212222002210131-1200300012013103-1132121032133313-2310321030202102-2101301330012220-1211233231212313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2302113200202121-3122333120003133-2322322231102223-0021222231032121-0310303221003123-0121001113002310-3213031311023011-3123122032132112"></a>

## stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.certificates — certificates / 033100000021 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-017.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-021.md#canonical-3300101333202310-3223231311312310-2030111033033203-1011103231310023-1313110120311121-1221031312313020-1210302113302033-2110231331000202)
- [stateful_service.advertise_options.advertise_on_public.port](resources--workload--reference--group-025.md#canonical-0120012201110120-2320030110233132-3233102311321321-2202232022223030-2133131131013111-3231210131302330-1022020011332122-3111230313032011)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-025.md#canonical-2203002110120031-3203013111221013-1321101003110222-2103121303322022-0021002301311230-0203100203111101-0321302220100020-1320220233023212)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https](resources--workload--reference--group-025.md#canonical-2320131312110211-0233203100020221-1102210111313012-2101323032112112-2012023213113210-1213111131330011-1020010101131133-0020231202232230)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params](resources--workload--reference--group-025.md#canonical-0330120220221323-3001303122321320-1021021120022323-3033111012020201-2201223300221222-2120132123232130-3021131112021012-3210022322020231)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.certificates

<a id="canonical-0021213320200012-3030203022202131-3211330332001033-3310110312010132-3213310032200302-3201111301200332-3200200220201013-2330310033231023"></a>

Type: `"object"`. list nested block, Optional.

Select one or more certificates with any domain names.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("name")}
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

Terraform syntax:

```terraform
certificates {
  # Configure direct properties listed below.
}
```

<a id="canonical-2313202031202031-3313310013221003-0120001112112000-3231112130202223-3020031321113203-0332300032231112-1222321101323302-3022333010032113"></a>

## Direct properties — certificates / 033100000021 / 3

<a id="canonical-3333031100231023-3032331203322202-2100301110221133-1030110003033221-3113220132310202-1103021131000203-3221210001220010-1230003010300213"></a>

<a id="canonical-3232010003212322-0110002131221300-3213133233122201-0100133211333333-2101122132022023-3132111331022303-0103123113323300-3222330321030312"></a>

## name property — certificates / 033100000021 / 4

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

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

<a id="canonical-0211002010220322-2312020112323302-3330123323030020-2012311303133300-1321010230221200-3133223121320233-3103230301312022-2312311113331002"></a>

<a id="canonical-1331000131313032-0222122222133302-2133231213120233-2302303020221203-2330103232333002-2000202331132222-0002031211033010-1012301011011122"></a>

## namespace property — certificates / 033100000021 / 5

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

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

<a id="canonical-1033303230012311-0212221303211031-2200231210331023-3313102232303310-1030303300300222-0113020121000033-2301203023120303-1101022332303122"></a>

<a id="canonical-1020320222031101-3113231022233232-0213333303002021-0112113313200300-3123032130000132-0321122310213130-3330113123011123-1101322301103331"></a>

## tenant property — certificates / 033100000021 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

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

<a id="canonical-3112330111030223-3320223330022002-0302212312311130-1303112032331312-0201230111103303-0333322201013312-0022302232231032-3220110221200131"></a>

## Next pages — certificates / 033100000021 / 7

- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params](resources--workload--reference--group-025.md#canonical-0330120220221323-3001303122321320-1021021120022323-3033111012020201-2201223300221222-2120132123232130-3021131112021012-3210022322020231)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-3311301002213232-3032212233333300-1301201000110122-1211001010000011-2233001000223302-1111201011313003-0233200311222201-2203120030313212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1200303103002023-1211311000110232-3102121222220003-2122220312023001-0032012130212200-0203100102222122-0200110210200102-2301013233200232"></a>

## stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.no_mtls — no_mtls / 001023101100 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-017.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-021.md#canonical-3300101333202310-3223231311312310-2030111033033203-1011103231310023-1313110120311121-1221031312313020-1210302113302033-2110231331000202)
- [stateful_service.advertise_options.advertise_on_public.port](resources--workload--reference--group-025.md#canonical-0120012201110120-2320030110233132-3233102311321321-2202232022223030-2133131131013111-3231210131302330-1022020011332122-3111230313032011)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-025.md#canonical-2203002110120031-3203013111221013-1321101003110222-2103121303322022-0021002301311230-0203100203111101-0321302220100020-1320220233023212)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https](resources--workload--reference--group-025.md#canonical-2320131312110211-0233203100020221-1102210111313012-2101323032112112-2012023213113210-1213111131330011-1020010101131133-0020231202232230)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params](resources--workload--reference--group-025.md#canonical-0330120220221323-3001303122321320-1021021120022323-3033111012020201-2201223300221222-2120132123232130-3021131112021012-3210022322020231)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.no_mtls

<a id="canonical-3112031122103320-0100200120030033-2013332302100330-0002130201002130-2121303300103131-3321230122223003-3333031211202023-1201203322230330"></a>

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
no_mtls = {}
```

<a id="canonical-1023130203131023-0222011033031300-0211230010131312-3003123313231311-3301003000123013-0101110232211113-3303321013031310-1003212333021122"></a>

## Direct properties — no_mtls / 001023101100 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0031020221201323-2221133320133003-0211303130202331-2322130233320220-0221101003232013-2213320031200210-0122303030100321-0220212001310023"></a>

## Next pages — no_mtls / 001023101100 / 4

- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params](resources--workload--reference--group-025.md#canonical-0330120220221323-3001303122321320-1021021120022323-3033111012020201-2201223300221222-2120132123232130-3021131112021012-3210022322020231)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-3011322100021322-2121021022310000-2011022031301301-3323233310022131-0113312101320323-0211100303013221-0023223321310100-2133210222331011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1332121303330132-3310033221212200-3013132312331103-1201300203020321-0122201311131202-0300021023201213-2200111201313011-1022320113221311"></a>

## stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.tls_config — tls_config / 330213032300 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-017.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-021.md#canonical-3300101333202310-3223231311312310-2030111033033203-1011103231310023-1313110120311121-1221031312313020-1210302113302033-2110231331000202)
- [stateful_service.advertise_options.advertise_on_public.port](resources--workload--reference--group-025.md#canonical-0120012201110120-2320030110233132-3233102311321321-2202232022223030-2133131131013111-3231210131302330-1022020011332122-3111230313032011)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-025.md#canonical-2203002110120031-3203013111221013-1321101003110222-2103121303322022-0021002301311230-0203100203111101-0321302220100020-1320220233023212)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https](resources--workload--reference--group-025.md#canonical-2320131312110211-0233203100020221-1102210111313012-2101323032112112-2012023213113210-1213111131330011-1020010101131133-0020231202232230)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params](resources--workload--reference--group-025.md#canonical-0330120220221323-3001303122321320-1021021120022323-3033111012020201-2201223300221222-2120132123232130-3021131112021012-3210022322020231)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.tls_config

<a id="canonical-1101333300002323-3112031032301200-1320313003203320-2221230102000002-0113200310003323-3231202222203131-2132221330203010-1311221023122312"></a>

Type: `"object"`. single nested block, Optional.

Defines various OPTIONS to configure TLS configuration parameters.

Upstream description:

This defines various OPTIONS to configure TLS configuration parameters.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("custom_security",
    "default_security"),
  validators.ConflictingObjectAttributes("custom_security",
    "low_security"),
  validators.ConflictingObjectAttributes("custom_security",
    "medium_security"),
  validators.ConflictingObjectAttributes("default_security",
    "low_security"),
  validators.ConflictingObjectAttributes("default_security",
    "medium_security"),
  validators.ConflictingObjectAttributes("low_security",
    "medium_security")}
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
  "x-ves-oneof-field-choice": "[\"custom_security\",\"default_security\",\"low_security\",\"medium_security\"]"
}
```

Terraform syntax:

```terraform
tls_config {
  # Configure direct properties listed below.
}
```

<a id="canonical-0132001333023132-3132213333101032-1131001001102211-1110012321123013-2321100312313332-1011301111020331-3321310330013033-0102212030020111"></a>

## Direct properties — tls_config / 330213032300 / 3

- [custom_security](resources--workload--reference--group-025.md#canonical-2112322132321113-1232102323132021-1020022020231213-2110122023233001-1201220132211203-3303131030013123-2000311312312330-1122210312322313): complete subsection reference.

- [default_security](resources--workload--reference--group-025.md#canonical-2002033023012300-2201103310331010-3321121223323030-3300221332201020-1221032330232332-3030300222212231-0123220002311021-3131123133313102): complete subsection reference.

- [low_security](resources--workload--reference--group-025.md#canonical-3320330113031320-3032132133223001-2320100223311202-0100313133112322-1302013323031323-0332210132001002-1112100122032020-0132321130030032): complete subsection reference.

- [medium_security](resources--workload--reference--group-025.md#canonical-2323202010211333-0330330203002032-2112022123112223-1121301321231310-3203321223213302-3311222300113033-2303220103233010-1031231000201111): complete subsection reference.

<a id="canonical-1331200100300022-1233330111031303-2221123122221122-3031301113300231-3030113110322131-3301033321303133-1202322112030111-1032230213122020"></a>

## Next pages — tls_config / 330213032300 / 4

- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.tls_config.custom_security](resources--workload--reference--group-025.md#canonical-2112322132321113-1232102323132021-1020022020231213-2110122023233001-1201220132211203-3303131030013123-2000311312312330-1122210312322313)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.tls_config.default_security](resources--workload--reference--group-025.md#canonical-2002033023012300-2201103310331010-3321121223323030-3300221332201020-1221032330232332-3030300222212231-0123220002311021-3131123133313102)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.tls_config.low_security](resources--workload--reference--group-025.md#canonical-3320330113031320-3032132133223001-2320100223311202-0100313133112322-1302013323031323-0332210132001002-1112100122032020-0132321130030032)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.tls_config.medium_security](resources--workload--reference--group-025.md#canonical-2323202010211333-0330330203002032-2112022123112223-1121301321231310-3203321223213302-3311222300113033-2303220103233010-1031231000201111)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params](resources--workload--reference--group-025.md#canonical-0330120220221323-3001303122321320-1021021120022323-3033111012020201-2201223300221222-2120132123232130-3021131112021012-3210022322020231)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-2112322132321113-1232102323132021-1020022020231213-2110122023233001-1201220132211203-3303131030013123-2000311312312330-1122210312322313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3130203003102110-2001010022110023-2233122033333223-3100022223302230-0201321032130331-3333312210321130-0202111022123130-3130303301111001"></a>

## stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.tls_config.custom_security — custom_security / 001331001332 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-017.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-021.md#canonical-3300101333202310-3223231311312310-2030111033033203-1011103231310023-1313110120311121-1221031312313020-1210302113302033-2110231331000202)
- [stateful_service.advertise_options.advertise_on_public.port](resources--workload--reference--group-025.md#canonical-0120012201110120-2320030110233132-3233102311321321-2202232022223030-2133131131013111-3231210131302330-1022020011332122-3111230313032011)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-025.md#canonical-2203002110120031-3203013111221013-1321101003110222-2103121303322022-0021002301311230-0203100203111101-0321302220100020-1320220233023212)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https](resources--workload--reference--group-025.md#canonical-2320131312110211-0233203100020221-1102210111313012-2101323032112112-2012023213113210-1213111131330011-1020010101131133-0020231202232230)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params](resources--workload--reference--group-025.md#canonical-0330120220221323-3001303122321320-1021021120022323-3033111012020201-2201223300221222-2120132123232130-3021131112021012-3210022322020231)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.tls_config](resources--workload--reference--group-025.md#canonical-3011322100021322-2121021022310000-2011022031301301-3323233310022131-0113312101320323-0211100303013221-0023223321310100-2133210222331011)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.tls_config.custom_security

<a id="canonical-1001131210210310-2203312023020203-2111230312323301-3213302332221000-1233012112313123-3312103311001131-0201201113222003-2211210121201122"></a>

Type: `"object"`. single nested block, Optional.

Defines TLS protocol config including min/max versions and allowed ciphers.

Upstream description:

This defines TLS protocol config including min/max versions and allowed ciphers.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("cipher_suites")}
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
custom_security {
  # Configure direct properties listed below.
}
```

<a id="canonical-3101333223322102-2201220010231113-2003102330320111-1021133213230301-0101330313221022-0013132131210003-3123031201023113-0133331000020220"></a>

## Direct properties — custom_security / 001331001332 / 3

<a id="canonical-3120003201311231-2330321031033312-3223221300200312-3013332212311320-3321110121303110-3012110123120312-1103312021202131-2233310201203012"></a>

<a id="canonical-2132000003323030-0232220211313303-2002321220210033-2012113003012110-2002203133231232-0231133220231220-0333211231220320-3302233210010213"></a>

## cipher_suites property — custom_security / 001331001332 / 4

Type: `["list", "string"]`. Optional.

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

<a id="canonical-2232213312330312-2302102201032323-2002000230212133-2201123001101122-3011311203132210-1230313021013312-0232111123021301-0231312203120200"></a>

<a id="canonical-3311223332233311-1030303220001222-1320022131332321-3313132103320020-1300322120111101-3301030103032031-3021131330022300-2202112302002000"></a>

## max_version property — custom_security / 001331001332 / 5

Type: `"string"`. Optional.

\[Enum: TLS\_AUTO|TLSv1\_0|TLSv1\_1|TLSv1\_2|TLSv1\_3\] TlsProtocol is enumeration of supported TLS
versions F5 Distributed Cloud will choose the optimal TLS version. Possible values are
\`TLS\_AUTO\`, \`TLSv1\_0\`, \`TLSv1\_1\`, \`TLSv1\_2\`, \`TLSv1\_3\`. Defaults to \`TLS\_AUTO\`.

Upstream description:

TlsProtocol is enumeration of supported TLS versions

F5 Distributed Cloud will choose the optimal TLS version.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("TLS_AUTO",
    "TLSv1_0",
    "TLSv1_1",
    "TLSv1_2",
    "TLSv1_3"),
}
```

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

<a id="canonical-1130303310330021-2123210300000300-3031230320133022-0012001300003211-0212213131320233-1220020003102220-2123301332301332-3322222112033013"></a>

<a id="canonical-1233200311101023-0220002002322220-1121201313312303-1133321210030000-3203221203012110-2102111313200010-2223021322230211-1102131220122110"></a>

## min_version property — custom_security / 001331001332 / 6

Type: `"string"`. Optional.

\[Enum: TLS\_AUTO|TLSv1\_0|TLSv1\_1|TLSv1\_2|TLSv1\_3\] TlsProtocol is enumeration of supported TLS
versions F5 Distributed Cloud will choose the optimal TLS version. Possible values are
\`TLS\_AUTO\`, \`TLSv1\_0\`, \`TLSv1\_1\`, \`TLSv1\_2\`, \`TLSv1\_3\`. Defaults to \`TLS\_AUTO\`.

Upstream description:

TlsProtocol is enumeration of supported TLS versions

F5 Distributed Cloud will choose the optimal TLS version.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("TLS_AUTO",
    "TLSv1_0",
    "TLSv1_1",
    "TLSv1_2",
    "TLSv1_3"),
}
```

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

<a id="canonical-2210132113030230-3320213201210210-1022230010113333-1121220222113122-1332033011002011-2131031300111122-3203113020012021-2122231020120303"></a>

## Next pages — custom_security / 001331001332 / 7

- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.tls_config](resources--workload--reference--group-025.md#canonical-3011322100021322-2121021022310000-2011022031301301-3323233310022131-0113312101320323-0211100303013221-0023223321310100-2133210222331011)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-2002033023012300-2201103310331010-3321121223323030-3300221332201020-1221032330232332-3030300222212231-0123220002311021-3131123133313102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2323232312000023-2031011332011000-0322321121032301-3101103200112000-2130012032020121-1311023330101120-2030013311330201-1010200101110030"></a>

## stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.tls_config.default_security — default_security / 122203322101 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-017.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-021.md#canonical-3300101333202310-3223231311312310-2030111033033203-1011103231310023-1313110120311121-1221031312313020-1210302113302033-2110231331000202)
- [stateful_service.advertise_options.advertise_on_public.port](resources--workload--reference--group-025.md#canonical-0120012201110120-2320030110233132-3233102311321321-2202232022223030-2133131131013111-3231210131302330-1022020011332122-3111230313032011)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-025.md#canonical-2203002110120031-3203013111221013-1321101003110222-2103121303322022-0021002301311230-0203100203111101-0321302220100020-1320220233023212)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https](resources--workload--reference--group-025.md#canonical-2320131312110211-0233203100020221-1102210111313012-2101323032112112-2012023213113210-1213111131330011-1020010101131133-0020231202232230)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params](resources--workload--reference--group-025.md#canonical-0330120220221323-3001303122321320-1021021120022323-3033111012020201-2201223300221222-2120132123232130-3021131112021012-3210022322020231)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.tls_config](resources--workload--reference--group-025.md#canonical-3011322100021322-2121021022310000-2011022031301301-3323233310022131-0113312101320323-0211100303013221-0023223321310100-2133210222331011)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.tls_config.default_security

<a id="canonical-1133330002011220-2301010232112020-0023020101233333-0230000010133313-1323012303312101-1013011202121130-3310213021230210-0322313202032030"></a>

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
default_security = {}
```

<a id="canonical-1300012230033003-3231200300112323-3130212103122310-3333223101203231-0321102310010213-1031012303000232-3103303030231003-3023301203031130"></a>

## Direct properties — default_security / 122203322101 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2100121111112312-2021230220131123-2103330310002111-3012011120332200-0220120013303001-0121013231323210-1333221133031132-1022200321323323"></a>

## Next pages — default_security / 122203322101 / 4

- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.tls_config](resources--workload--reference--group-025.md#canonical-3011322100021322-2121021022310000-2011022031301301-3323233310022131-0113312101320323-0211100303013221-0023223321310100-2133210222331011)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-3320330113031320-3032132133223001-2320100223311202-0100313133112322-1302013323031323-0332210132001002-1112100122032020-0132321130030032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0200111003320021-3013332000231331-2022021113203223-2220320133323223-1113313233112333-1312123300230112-3102320113303003-0100012301212111"></a>

## stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.tls_config.low_security — low_security / 201103123030 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-017.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-021.md#canonical-3300101333202310-3223231311312310-2030111033033203-1011103231310023-1313110120311121-1221031312313020-1210302113302033-2110231331000202)
- [stateful_service.advertise_options.advertise_on_public.port](resources--workload--reference--group-025.md#canonical-0120012201110120-2320030110233132-3233102311321321-2202232022223030-2133131131013111-3231210131302330-1022020011332122-3111230313032011)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-025.md#canonical-2203002110120031-3203013111221013-1321101003110222-2103121303322022-0021002301311230-0203100203111101-0321302220100020-1320220233023212)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https](resources--workload--reference--group-025.md#canonical-2320131312110211-0233203100020221-1102210111313012-2101323032112112-2012023213113210-1213111131330011-1020010101131133-0020231202232230)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params](resources--workload--reference--group-025.md#canonical-0330120220221323-3001303122321320-1021021120022323-3033111012020201-2201223300221222-2120132123232130-3021131112021012-3210022322020231)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.tls_config](resources--workload--reference--group-025.md#canonical-3011322100021322-2121021022310000-2011022031301301-3323233310022131-0113312101320323-0211100303013221-0023223321310100-2133210222331011)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.tls_config.low_security

<a id="canonical-3121101131323330-3123123030233121-3121110210232030-1111013130101230-2333220320122133-3232203001113003-3232123003123031-2113012301213102"></a>

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
low_security = {}
```

<a id="canonical-3232002030231130-2222303133121030-2020310031223023-3301130311323312-1203310023131331-3021321211300131-0023302003120113-1003200000211220"></a>

## Direct properties — low_security / 201103123030 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2122132100012301-0013210130103021-3231211001102230-0023310002031211-3213211230100333-1013222230310001-2332232212001330-3311001112113202"></a>

## Next pages — low_security / 201103123030 / 4

- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.tls_config](resources--workload--reference--group-025.md#canonical-3011322100021322-2121021022310000-2011022031301301-3323233310022131-0113312101320323-0211100303013221-0023223321310100-2133210222331011)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-2323202010211333-0330330203002032-2112022123112223-1121301321231310-3203321223213302-3311222300113033-2303220103233010-1031231000201111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2023012011111311-2321312130332023-0030011211033222-2333130132203013-3232331202011111-0231022103232021-0331232233123201-2131312300201212"></a>

## stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.tls_config.medium_security — medium_security / 201002123222 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-017.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-021.md#canonical-3300101333202310-3223231311312310-2030111033033203-1011103231310023-1313110120311121-1221031312313020-1210302113302033-2110231331000202)
- [stateful_service.advertise_options.advertise_on_public.port](resources--workload--reference--group-025.md#canonical-0120012201110120-2320030110233132-3233102311321321-2202232022223030-2133131131013111-3231210131302330-1022020011332122-3111230313032011)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-025.md#canonical-2203002110120031-3203013111221013-1321101003110222-2103121303322022-0021002301311230-0203100203111101-0321302220100020-1320220233023212)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https](resources--workload--reference--group-025.md#canonical-2320131312110211-0233203100020221-1102210111313012-2101323032112112-2012023213113210-1213111131330011-1020010101131133-0020231202232230)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params](resources--workload--reference--group-025.md#canonical-0330120220221323-3001303122321320-1021021120022323-3033111012020201-2201223300221222-2120132123232130-3021131112021012-3210022322020231)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.tls_config](resources--workload--reference--group-025.md#canonical-3011322100021322-2121021022310000-2011022031301301-3323233310022131-0113312101320323-0211100303013221-0023223321310100-2133210222331011)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.tls_config.medium_security

<a id="canonical-2003233330130323-3031310330000310-3002233331211103-1020223031003002-2201020122103233-1023120222032131-0311100110122201-3102000212011033"></a>

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
medium_security = {}
```

<a id="canonical-1131013000121203-2002113001030021-2110010022301013-3010012022033033-0202122322122311-3313102332012330-1301323103023320-1321021000222121"></a>

## Direct properties — medium_security / 201002123222 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3210221213012012-3331200031031021-1320322330333202-0312310030323232-2332230231230002-0311312300333233-1333320321121131-1120110231230031"></a>

## Next pages — medium_security / 201002123222 / 4

- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.tls_config](resources--workload--reference--group-025.md#canonical-3011322100021322-2121021022310000-2011022031301301-3323233310022131-0113312101320323-0211100303013221-0023223321310100-2133210222331011)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-1023122310132111-2212330202030310-2112313020200023-3101321202230303-1220121020221312-1210310302131011-2132302200011221-1002023020300223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2302033101310120-1010131322120202-0302210013323130-2322310113310313-2303203302202020-2201003201310010-3011012030032301-3200122000210213"></a>

## stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.use_mtls — use_mtls / 001113030333 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-017.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-021.md#canonical-3300101333202310-3223231311312310-2030111033033203-1011103231310023-1313110120311121-1221031312313020-1210302113302033-2110231331000202)
- [stateful_service.advertise_options.advertise_on_public.port](resources--workload--reference--group-025.md#canonical-0120012201110120-2320030110233132-3233102311321321-2202232022223030-2133131131013111-3231210131302330-1022020011332122-3111230313032011)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-025.md#canonical-2203002110120031-3203013111221013-1321101003110222-2103121303322022-0021002301311230-0203100203111101-0321302220100020-1320220233023212)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https](resources--workload--reference--group-025.md#canonical-2320131312110211-0233203100020221-1102210111313012-2101323032112112-2012023213113210-1213111131330011-1020010101131133-0020231202232230)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params](resources--workload--reference--group-025.md#canonical-0330120220221323-3001303122321320-1021021120022323-3033111012020201-2201223300221222-2120132123232130-3021131112021012-3210022322020231)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.use_mtls

<a id="canonical-3313210303323203-0330000312130300-0200030100322103-0011110301211320-3011112323101300-1302232201312133-2122102231211210-3313310011130012"></a>

Type: `"object"`. single nested block, Optional.

Validation context for downstream client TLS connections.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("crl",
    "no_crl"),
  validators.ConflictingObjectAttributes("trusted_ca",
    "trusted_ca_url"),
  validators.ConflictingObjectAttributes("xfcc_disabled",
    "xfcc_options")}
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
  "x-ves-oneof-field-crl_choice": "[\"crl\",\"no_crl\"]",
  "x-ves-oneof-field-trusted_ca_choice": "[\"trusted_ca\",\"trusted_ca_url\"]",
  "x-ves-oneof-field-xfcc_header": "[\"xfcc_disabled\",\"xfcc_options\"]"
}
```

Terraform syntax:

```terraform
use_mtls {
  # Configure direct properties listed below.
}
```

<a id="canonical-3311131022023033-3233021113230312-0330003113003223-0123221012202120-3210220232020021-1210221332232331-1102103201211210-3320132101310331"></a>

## Direct properties — use_mtls / 001113030333 / 3

<a id="canonical-1203020312313132-1210030303233101-0000200123202031-3012033011101013-0310303222022300-3113212022323131-1111010123000103-0223031000302233"></a>

<a id="canonical-3000011130030111-2302101032021313-2233131120010202-1321023311003300-2200332302211020-0311232332001333-3222202220201323-0123133213233111"></a>

## client_certificate_optional property — use_mtls / 001113030333 / 4

Type: `"bool"`. Optional.

Client certificate is optional. If the client has provided a certificate, the load balancer will
verify it. If certification verification fails, the connection will be terminated.

Upstream description:

Client certificate is optional. If the client has provided a certificate, the load balancer will
verify it. If certification verification fails, the connection will be terminated. If the client
does not provide a certificate, the connection will be accepted.

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

- [crl](resources--workload--reference--group-026.md#canonical-0200123201103213-1003301322200021-1320312012332001-2121030121131323-0300131210310023-2310030212002320-3202012120203333-3232021020323031): complete subsection reference.

- [no_crl](resources--workload--reference--group-026.md#canonical-3300112301020211-2210203222231212-0132132310103013-1203301331232221-0133001332203333-1303011100201133-2303231103233310-1030022233323333): complete subsection reference.

- [trusted_ca](resources--workload--reference--group-026.md#canonical-0303222332201101-3103233323020220-3030332212202201-3121311320301310-1032020131001030-1220203212111223-3113303122010102-2323202133221113): complete subsection reference.

<a id="canonical-2232230213213303-2300233121301312-2112001023021233-3101123221010202-0322112233201313-1113101230030030-0301333320131213-3221021310002022"></a>

<a id="canonical-2003132210320110-0300032122133121-3022320210110002-3311003210112101-3011120231013032-3030122122311320-1100222310023232-0120331001101123"></a>

## trusted_ca_url property — use_mtls / 001113030333 / 5

Type: `"string"`. Optional.

Exclusive with \[trusted\_ca\] Upload a Root CA Certificate specifically for this Load Balancer.

Upstream description:

Exclusive with \[trusted\_ca\] Upload a Root CA Certificate specifically for this Load Balancer.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 131072),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 131072,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 131072,
      "min": 1
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 131072,
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
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.truststore_url": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.truststore_url": "true"
  }
}
```

- [xfcc_disabled](resources--workload--reference--group-026.md#canonical-3020100130232111-0122013010332131-1311000123113310-2202031223222010-1132311030133111-0013222211331001-1103222233123211-3123030123230302): complete subsection reference.

- [xfcc_options](resources--workload--reference--group-026.md#canonical-1221132012332310-1013300332321132-0031323313202001-0213002031331321-1232213321322120-1020311333231333-0231001200222130-3130102313201210): complete subsection reference.
