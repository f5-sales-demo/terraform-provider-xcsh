---
page_title: "xcsh_workload reference"
subcategory: "Container"
description: "Complete grouped canonical reference for xcsh_workload reference."
---

# xcsh_workload reference

<a id="canonical-3102103313032210-0102301102011232-1303110012101212-3032312111101121-2212320022233233-2110332122303130-0322022331103012-2002002322210003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.default_route.auto_host_rewrite` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-018.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-018.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_custom](resources--workload--reference--group-018.md#canonical-0203102001320123-1120130012130212-0111101220231333-3101331123203121-1110000311003232-1321231201202312-3200200021132112-0200303231330002)
- [stateful_service.advertise_options.advertise_custom.ports](resources--workload--reference--group-018.md#canonical-3111311322332303-1201030020021201-3133021233000313-2022333331131202-3221312032323033-3133120120113232-3330021120310331-2331303302032231)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-018.md#canonical-0313222022003221-2301103230203023-3120322210033100-2121220010332201-3133231021021220-1313001311100021-0313102021211212-0331322021202002)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.default_route](resources--workload--reference--group-018.md#canonical-2031112003330101-0310120302332023-3120021113320130-3110000310223201-1012000003321130-2312213320212011-0022300121030033-3100123121103201)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.default_route.auto_host_rewrite

<a id="canonical-1311003302002132-1230121322200002-0000303102020032-0022020202222200-1002133203232000-0213103331223022-3020000310300020-3212220023302133"></a>

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

<a id="canonical-2233022211333231-1323213121201103-1312330322030123-0003203301103101-0332011031223232-2322233001313302-0312301223030330-1003213211032213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.default_route.disable_host_rewrite` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-018.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-018.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_custom](resources--workload--reference--group-018.md#canonical-0203102001320123-1120130012130212-0111101220231333-3101331123203121-1110000311003232-1321231201202312-3200200021132112-0200303231330002)
- [stateful_service.advertise_options.advertise_custom.ports](resources--workload--reference--group-018.md#canonical-3111311322332303-1201030020021201-3133021233000313-2022333331131202-3221312032323033-3133120120113232-3330021120310331-2331303302032231)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-018.md#canonical-0313222022003221-2301103230203023-3120322210033100-2121220010332201-3133231021021220-1313001311100021-0313102021211212-0331322021202002)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.default_route](resources--workload--reference--group-018.md#canonical-2031112003330101-0310120302332023-3120021113320130-3110000310223201-1012000003321130-2312213320212011-0022300121030033-3100123121103201)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.default_route.disable_host_rewrite

<a id="canonical-2331210223020330-0111122131202330-0232222202202333-2030003332022000-3111320123222112-0130000320131230-2201210203213100-1322021112231101"></a>

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

<a id="canonical-0001001230203210-1010213221120331-1102200030010210-1101322011000303-0020221022232213-1231230310223023-2110212231011313-1321232000313021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-018.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-018.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_custom](resources--workload--reference--group-018.md#canonical-0203102001320123-1120130012130212-0111101220231333-3101331123203121-1110000311003232-1321231201202312-3200200021132112-0200303231330002)
- [stateful_service.advertise_options.advertise_custom.ports](resources--workload--reference--group-018.md#canonical-3111311322332303-1201030020021201-3133021233000313-2022333331131202-3221312032323033-3133120120113232-3330021120310331-2331303302032231)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-018.md#canonical-0313222022003221-2301103230203023-3120322210033100-2121220010332201-3133231021021220-1313001311100021-0313102021211212-0331322021202002)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http

<a id="canonical-1110010121201202-1200123020202333-0213020213313231-3331333211212122-3311113032131021-2220220212101202-3203320300103132-1033101012131210"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
http {
  # Configure direct properties listed below.
}
```

<a id="canonical-0122003233313031-2100112001113003-1333301000020100-3130301200021120-3123132021100133-2130000113020010-1033322110202211-2323200202213201"></a>

### Direct properties for `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http`

<a id="canonical-0222203302032033-3211213122001230-3222112101233100-3331132221312213-0101320022331313-3033221210233011-0232030013232202-0223100221222112"></a>

#### `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http.dns_volterra_managed` property

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

<a id="canonical-3233133121321313-3131311031020212-0233023032222220-3320131022313222-0132322133332210-2202012021100232-3111201321211231-1331312211322331"></a>

<a id="canonical-1230203211033030-0011322332012013-0032230130012000-3100022300013311-0331013211013123-1111332113011231-2312011032102231-0300123222012231"></a>

#### `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http.port` property

Type: `"number"`. Optional.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-1132123003123100-1110021101021320-2323123022113323-2302013133330103-2221303032231113-2022023132011023-3133131231013020-3331003223302212"></a>

<a id="canonical-3112032222022020-2230010311333113-2130313211232102-0222103322230113-2222332313311002-0133001222033113-0131321122320102-3103223232221331"></a>

#### `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http.port_ranges` property

Type: `"string"`. Optional.

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

<a id="canonical-0112020200111022-0013103013103331-1110301001233101-1300202101312213-3301031210112102-1221232301300211-0231000232130131-3132113303133333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-018.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-018.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_custom](resources--workload--reference--group-018.md#canonical-0203102001320123-1120130012130212-0111101220231333-3101331123203121-1110000311003232-1321231201202312-3200200021132112-0200303231330002)
- [stateful_service.advertise_options.advertise_custom.ports](resources--workload--reference--group-018.md#canonical-3111311322332303-1201030020021201-3133021233000313-2022333331131202-3221312032323033-3133120120113232-3330021120310331-2331303302032231)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-018.md#canonical-0313222022003221-2301103230203023-3120322210033100-2121220010332201-3133231021021220-1313001311100021-0313102021211212-0331322021202002)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https

<a id="canonical-0120310310023211-3131320332233110-0100310123030011-2330301013323021-2112320302331231-0001123023313213-2103223200103222-2202012223022002"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
https {
  # Configure direct properties listed below.
}
```

<a id="canonical-1001301110111213-0011102231133333-3231032010200332-2311031013010213-0201332020221131-3031001111202011-0123122011003020-1122020302030120"></a>

### Direct properties for `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https`

<a id="canonical-1033333300330030-2322202112221223-1320223130030100-0322331131103210-3030010322113123-3220130223002113-0321201330133200-1231213032232000"></a>

#### `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.add_hsts` property

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

<a id="canonical-3012132302302101-0102220001313202-1321133011031313-1103032202231023-1320213030121332-1011112122102012-2031000110001010-1011110031233113"></a>

<a id="canonical-2200330103023300-2332121300101302-1132110210212232-3201122323133232-1331031010001221-2201311100100130-1022111232323110-0013210310001012"></a>

#### `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.append_server_name` property

Type: `"string"`. Optional.

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
    "ves.io.schema.rules.string.max_len": "8096"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "8096"
  }
}
```

- [coalescing_options](resources--workload--reference--group-019.md#canonical-2032232023103112-3132002323011100-2333333333213030-0131331001332133-3231230321300001-3121020113202300-0031031222210022-3100303312203312): complete subsection reference.

<a id="canonical-1301231032211202-1002223130012320-0122120200332003-3301133012231320-1202122201233313-2332303300230113-0222123310103033-3123211210102003"></a>

<a id="canonical-2110123330232132-3220330120100133-2121033310331101-2220131333100032-3210220020133332-1302200102301331-0333021012222022-2130301132210112"></a>

#### `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.connection_idle_timeout` property

Type: `"number"`. Optional.

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
    "ves.io.schema.rules.uint32.lte": "600000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "600000"
  }
}
```

- [default_header](resources--workload--reference--group-019.md#canonical-0003223323303031-3311332233132213-3000212121311110-0322333000010000-2012102033100303-2333200110221100-3023101230212212-0321110100231233): complete subsection reference.

- [default_loadbalancer](resources--workload--reference--group-019.md#canonical-2301030210312013-3330212130031303-3001333222011323-0133033200232003-2033031011231022-1331312210302033-3202203112021311-1123211231122020): complete subsection reference.

- [disable_path_normalize](resources--workload--reference--group-019.md#canonical-0123111002221221-3000302331312022-3020301301022002-2130333003203132-0322310312020311-1033223110102102-1222203300201132-0210022222201102): complete subsection reference.

- [enable_path_normalize](resources--workload--reference--group-019.md#canonical-2130133323100220-2131101221132301-1100311112031132-2020301212111213-0130001222222232-1121312101101020-1322130333030133-2322303321010132): complete subsection reference.

- [http_protocol_options](resources--workload--reference--group-019.md#canonical-3202212001033033-3133231302102230-0333203123220020-0022212313320130-3032131111301011-1301121203330231-2010200210001321-0233311332222011): complete subsection reference.

<a id="canonical-3203031123210120-3331303131332232-2202121111221001-2002122020201310-0231312302131301-3002031022333022-3010311130231120-1103011330100100"></a>

<a id="canonical-0012230202010331-0222021300022322-3131102233023301-2122120231302113-0322202020220120-2230303001321002-3221012313220121-3110010100330001"></a>

#### `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_redirect` property

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

- [non_default_loadbalancer](resources--workload--reference--group-019.md#canonical-1103323231113001-2101233200313132-0113203102011131-2310133213023120-2300020303303112-1311213222021323-3201102120300132-1030302220012022): complete subsection reference.

- [pass_through](resources--workload--reference--group-019.md#canonical-3122320112122220-2222123003111322-1032303200202320-0301311021210013-0010110320313023-3321020211021302-2302303311133010-3231010220310031): complete subsection reference.

<a id="canonical-3021222313010203-3012031133311010-2131100123333222-2230203033211331-0303021120033303-2133122111330231-0130130321021123-0232202320231120"></a>

<a id="canonical-1221201210020333-0003230331311323-0322100231303101-0012303311320032-2010233010223301-2300203302312113-1010012031300312-2320300113123232"></a>

#### `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.port` property

Type: `"number"`. Optional.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-1131101113301310-2303020113103231-2011020111322200-1330213212013100-3130100212000123-3120233111023031-2103200012201321-2003210323023111"></a>

<a id="canonical-1203022202121021-0213001202330223-3212220120301220-3131323030203121-2031020120310021-1120233121301231-2030200132123203-2203030133020231"></a>

#### `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.port_ranges` property

Type: `"string"`. Optional.

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

<a id="canonical-0033303101323213-3002212331203102-2200312221220002-3030312323220031-1011222110203111-0100002223222121-1223112020222211-0030231021011210"></a>

<a id="canonical-1001011313010210-1223121123022323-2322133220313130-2121032123111212-1313223012333202-0001323120033103-1002102003123101-2311333220121330"></a>

#### `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.server_name` property

Type: `"string"`. Optional.

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
    "ves.io.schema.rules.string.max_len": "8096"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "8096"
  }
}
```

- [tls_cert_params](resources--workload--reference--group-019.md#canonical-1330323120222300-2123201010331320-0300232300332322-0010223222211002-1032302312300112-0233320002220301-1020331113013010-3202200223121221): complete subsection reference.

- [tls_parameters](resources--workload--reference--group-019.md#canonical-2020320003013033-2231100311132300-2132132133022132-3002120002132203-3131231201330331-0231012122032023-1202211010110230-0230013122122221): complete subsection reference.

<a id="canonical-2032232023103112-3132002323011100-2333333333213030-0131331001332133-3231230321300001-3121020113202300-0031031222210022-3100303312203312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.coalescing_options` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-018.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-018.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_custom](resources--workload--reference--group-018.md#canonical-0203102001320123-1120130012130212-0111101220231333-3101331123203121-1110000311003232-1321231201202312-3200200021132112-0200303231330002)
- [stateful_service.advertise_options.advertise_custom.ports](resources--workload--reference--group-018.md#canonical-3111311322332303-1201030020021201-3133021233000313-2022333331131202-3221312032323033-3133120120113232-3330021120310331-2331303302032231)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-018.md#canonical-0313222022003221-2301103230203023-3120322210033100-2121220010332201-3133231021021220-1313001311100021-0313102021211212-0331322021202002)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-019.md#canonical-0112020200111022-0013103013103331-1110301001233101-1300202101312213-3301031210112102-1221232301300211-0231000232130131-3132113303133333)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.coalescing_options

<a id="canonical-1301323033020202-2133133302210233-0023231001201311-3133012110211023-3202233120200201-2233031330231232-0101113031210003-0301211201220002"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
coalescing_options {
  # Configure direct properties listed below.
}
```

<a id="canonical-3333013131220330-3302021121010323-2313231211133313-0103212200102031-2231211030310011-3011302222023203-3002310123102313-2332321103213301"></a>

### Direct properties for `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.coalescing_options`

- [default_coalescing](resources--workload--reference--group-019.md#canonical-0013002130222122-0232011323323313-0103322131310223-0300131033232133-0103223223322302-2013100033113013-3320203231321222-3310130113020200): complete subsection reference.

- [strict_coalescing](resources--workload--reference--group-019.md#canonical-2010111233322110-2233220133212231-3201112221030333-3110010013300201-3331022031103023-3021212212121321-2113101133100103-1110021210203001): complete subsection reference.

<a id="canonical-0013002130222122-0232011323323313-0103322131310223-0300131033232133-0103223223322302-2013100033113013-3320203231321222-3310130113020200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.coalescing_options.default_coalescing` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-018.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-018.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_custom](resources--workload--reference--group-018.md#canonical-0203102001320123-1120130012130212-0111101220231333-3101331123203121-1110000311003232-1321231201202312-3200200021132112-0200303231330002)
- [stateful_service.advertise_options.advertise_custom.ports](resources--workload--reference--group-018.md#canonical-3111311322332303-1201030020021201-3133021233000313-2022333331131202-3221312032323033-3133120120113232-3330021120310331-2331303302032231)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-018.md#canonical-0313222022003221-2301103230203023-3120322210033100-2121220010332201-3133231021021220-1313001311100021-0313102021211212-0331322021202002)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-019.md#canonical-0112020200111022-0013103013103331-1110301001233101-1300202101312213-3301031210112102-1221232301300211-0231000232130131-3132113303133333)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.coalescing_options](resources--workload--reference--group-019.md#canonical-2032232023103112-3132002323011100-2333333333213030-0131331001332133-3231230321300001-3121020113202300-0031031222210022-3100303312203312)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.coalescing_options.default_coalescing

<a id="canonical-1010313023223013-3202331221021201-2133332301013320-1123223321331313-0300310030002223-3202231322031330-3033132122220213-0302112101310023"></a>

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

<a id="canonical-2010111233322110-2233220133212231-3201112221030333-3110010013300201-3331022031103023-3021212212121321-2113101133100103-1110021210203001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.coalescing_options.strict_coalescing` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-018.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-018.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_custom](resources--workload--reference--group-018.md#canonical-0203102001320123-1120130012130212-0111101220231333-3101331123203121-1110000311003232-1321231201202312-3200200021132112-0200303231330002)
- [stateful_service.advertise_options.advertise_custom.ports](resources--workload--reference--group-018.md#canonical-3111311322332303-1201030020021201-3133021233000313-2022333331131202-3221312032323033-3133120120113232-3330021120310331-2331303302032231)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-018.md#canonical-0313222022003221-2301103230203023-3120322210033100-2121220010332201-3133231021021220-1313001311100021-0313102021211212-0331322021202002)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-019.md#canonical-0112020200111022-0013103013103331-1110301001233101-1300202101312213-3301031210112102-1221232301300211-0231000232130131-3132113303133333)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.coalescing_options](resources--workload--reference--group-019.md#canonical-2032232023103112-3132002323011100-2333333333213030-0131331001332133-3231230321300001-3121020113202300-0031031222210022-3100303312203312)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.coalescing_options.strict_coalescing

<a id="canonical-1112131323201332-2220300232102333-1021030130033221-2123002313112331-2033300302211101-3132000000031323-2111323322302100-3103312021101132"></a>

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

<a id="canonical-0003223323303031-3311332233132213-3000212121311110-0322333000010000-2012102033100303-2333200110221100-3023101230212212-0321110100231233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.default_header` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-018.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-018.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_custom](resources--workload--reference--group-018.md#canonical-0203102001320123-1120130012130212-0111101220231333-3101331123203121-1110000311003232-1321231201202312-3200200021132112-0200303231330002)
- [stateful_service.advertise_options.advertise_custom.ports](resources--workload--reference--group-018.md#canonical-3111311322332303-1201030020021201-3133021233000313-2022333331131202-3221312032323033-3133120120113232-3330021120310331-2331303302032231)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-018.md#canonical-0313222022003221-2301103230203023-3120322210033100-2121220010332201-3133231021021220-1313001311100021-0313102021211212-0331322021202002)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-019.md#canonical-0112020200111022-0013103013103331-1110301001233101-1300202101312213-3301031210112102-1221232301300211-0231000232130131-3132113303133333)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.default_header

<a id="canonical-1311321033110211-1331203022223231-1103201331312112-1221301011332331-3223213201302013-2332101111113333-2230122201122231-1021211030013202"></a>

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

<a id="canonical-2301030210312013-3330212130031303-3001333222011323-0133033200232003-2033031011231022-1331312210302033-3202203112021311-1123211231122020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.default_loadbalancer` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-018.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-018.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_custom](resources--workload--reference--group-018.md#canonical-0203102001320123-1120130012130212-0111101220231333-3101331123203121-1110000311003232-1321231201202312-3200200021132112-0200303231330002)
- [stateful_service.advertise_options.advertise_custom.ports](resources--workload--reference--group-018.md#canonical-3111311322332303-1201030020021201-3133021233000313-2022333331131202-3221312032323033-3133120120113232-3330021120310331-2331303302032231)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-018.md#canonical-0313222022003221-2301103230203023-3120322210033100-2121220010332201-3133231021021220-1313001311100021-0313102021211212-0331322021202002)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-019.md#canonical-0112020200111022-0013103013103331-1110301001233101-1300202101312213-3301031210112102-1221232301300211-0231000232130131-3132113303133333)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.default_loadbalancer

<a id="canonical-1100201303101031-0330332003310211-1301330231102320-0003020212220130-3022331033000110-1102212122302333-3123202330123210-0010220123232113"></a>

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

<a id="canonical-0123111002221221-3000302331312022-3020301301022002-2130333003203132-0322310312020311-1033223110102102-1222203300201132-0210022222201102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.disable_path_normalize` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-018.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-018.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_custom](resources--workload--reference--group-018.md#canonical-0203102001320123-1120130012130212-0111101220231333-3101331123203121-1110000311003232-1321231201202312-3200200021132112-0200303231330002)
- [stateful_service.advertise_options.advertise_custom.ports](resources--workload--reference--group-018.md#canonical-3111311322332303-1201030020021201-3133021233000313-2022333331131202-3221312032323033-3133120120113232-3330021120310331-2331303302032231)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-018.md#canonical-0313222022003221-2301103230203023-3120322210033100-2121220010332201-3133231021021220-1313001311100021-0313102021211212-0331322021202002)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-019.md#canonical-0112020200111022-0013103013103331-1110301001233101-1300202101312213-3301031210112102-1221232301300211-0231000232130131-3132113303133333)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.disable_path_normalize

<a id="canonical-0332320220302323-1333132203020011-0231220310221030-0133120031332303-0331012211323300-3112322000202201-0102021333203330-0233110013201222"></a>

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

<a id="canonical-2130133323100220-2131101221132301-1100311112031132-2020301212111213-0130001222222232-1121312101101020-1322130333030133-2322303321010132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.enable_path_normalize` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-018.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-018.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_custom](resources--workload--reference--group-018.md#canonical-0203102001320123-1120130012130212-0111101220231333-3101331123203121-1110000311003232-1321231201202312-3200200021132112-0200303231330002)
- [stateful_service.advertise_options.advertise_custom.ports](resources--workload--reference--group-018.md#canonical-3111311322332303-1201030020021201-3133021233000313-2022333331131202-3221312032323033-3133120120113232-3330021120310331-2331303302032231)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-018.md#canonical-0313222022003221-2301103230203023-3120322210033100-2121220010332201-3133231021021220-1313001311100021-0313102021211212-0331322021202002)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-019.md#canonical-0112020200111022-0013103013103331-1110301001233101-1300202101312213-3301031210112102-1221232301300211-0231000232130131-3132113303133333)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.enable_path_normalize

<a id="canonical-0210301332321213-2230033213320232-1220132102303200-2301330000133313-2332021300203212-3120113102211321-2020232230032222-3313310331100123"></a>

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

<a id="canonical-3202212001033033-3133231302102230-0333203123220020-0022212313320130-3032131111301011-1301121203330231-2010200210001321-0233311332222011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-018.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-018.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_custom](resources--workload--reference--group-018.md#canonical-0203102001320123-1120130012130212-0111101220231333-3101331123203121-1110000311003232-1321231201202312-3200200021132112-0200303231330002)
- [stateful_service.advertise_options.advertise_custom.ports](resources--workload--reference--group-018.md#canonical-3111311322332303-1201030020021201-3133021233000313-2022333331131202-3221312032323033-3133120120113232-3330021120310331-2331303302032231)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-018.md#canonical-0313222022003221-2301103230203023-3120322210033100-2121220010332201-3133231021021220-1313001311100021-0313102021211212-0331322021202002)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-019.md#canonical-0112020200111022-0013103013103331-1110301001233101-1300202101312213-3301031210112102-1221232301300211-0231000232130131-3132113303133333)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options

<a id="canonical-2211010311233122-1100302100101032-2031332131201001-2231020212000312-2011210133003131-2210333033303101-3223311132302221-0011102101012321"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
http_protocol_options {
  # Configure direct properties listed below.
}
```

<a id="canonical-2101323120122030-0230220032122033-1330132001132222-3131121330223101-0112122303122113-0112103000333103-3313313123021213-2310001310303113"></a>

### Direct properties for `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options`

- [http_protocol_enable_v1_only](resources--workload--reference--group-019.md#canonical-0101121203223101-2303112021302213-2210233131121322-1010121010210301-2330311311003110-2202230130203213-2331310302232000-1213011200013332): complete subsection reference.

- [http_protocol_enable_v1_v2](resources--workload--reference--group-019.md#canonical-1233012312003333-2101101012232120-0100310010232311-0011212101122211-2001031003213123-1031101223300032-3223031101001303-3201303031321011): complete subsection reference.

- [http_protocol_enable_v2_only](resources--workload--reference--group-019.md#canonical-2120212030201112-1220311010203002-0121211322222323-2032310203313301-2111110222212103-0222302013331310-1331023200202002-1220000133031023): complete subsection reference.

<a id="canonical-0101121203223101-2303112021302213-2210233131121322-1010121010210301-2330311311003110-2202230130203213-2331310302232000-1213011200013332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-018.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-018.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_custom](resources--workload--reference--group-018.md#canonical-0203102001320123-1120130012130212-0111101220231333-3101331123203121-1110000311003232-1321231201202312-3200200021132112-0200303231330002)
- [stateful_service.advertise_options.advertise_custom.ports](resources--workload--reference--group-018.md#canonical-3111311322332303-1201030020021201-3133021233000313-2022333331131202-3221312032323033-3133120120113232-3330021120310331-2331303302032231)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-018.md#canonical-0313222022003221-2301103230203023-3120322210033100-2121220010332201-3133231021021220-1313001311100021-0313102021211212-0331322021202002)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-019.md#canonical-0112020200111022-0013103013103331-1110301001233101-1300202101312213-3301031210112102-1221232301300211-0231000232130131-3132113303133333)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options](resources--workload--reference--group-019.md#canonical-3202212001033033-3133231302102230-0333203123220020-0022212313320130-3032131111301011-1301121203330231-2010200210001321-0233311332222011)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only

<a id="canonical-0310020332231110-3302210320021103-2213320123010132-2331123301330000-2320322202331321-3103223113000222-1102210330023311-3311021210111002"></a>

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

<a id="canonical-3102203310121202-0131002312332030-3031120221131020-3332031221221132-3213033333020310-3031123132211312-2320111102102102-3322123322331310"></a>

### Direct properties for `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only`

- [header_transformation](resources--workload--reference--group-019.md#canonical-3020121301012012-1323210110303103-1332333110211232-3101112232311132-3303103303122230-0233302131022013-3203113021013212-1223222232202303): complete subsection reference.

<a id="canonical-3020121301012012-1323210110303103-1332333110211232-3101112232311132-3303103303122230-0233302131022013-3203113021013212-1223222232202303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-018.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-018.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_custom](resources--workload--reference--group-018.md#canonical-0203102001320123-1120130012130212-0111101220231333-3101331123203121-1110000311003232-1321231201202312-3200200021132112-0200303231330002)
- [stateful_service.advertise_options.advertise_custom.ports](resources--workload--reference--group-018.md#canonical-3111311322332303-1201030020021201-3133021233000313-2022333331131202-3221312032323033-3133120120113232-3330021120310331-2331303302032231)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-018.md#canonical-0313222022003221-2301103230203023-3120322210033100-2121220010332201-3133231021021220-1313001311100021-0313102021211212-0331322021202002)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-019.md#canonical-0112020200111022-0013103013103331-1110301001233101-1300202101312213-3301031210112102-1221232301300211-0231000232130131-3132113303133333)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options](resources--workload--reference--group-019.md#canonical-3202212001033033-3133231302102230-0333203123220020-0022212313320130-3032131111301011-1301121203330231-2010200210001321-0233311332222011)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only](resources--workload--reference--group-019.md#canonical-0101121203223101-2303112021302213-2210233131121322-1010121010210301-2330311311003110-2202230130203213-2331310302232000-1213011200013332)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation

<a id="canonical-2310312012330001-3120203203330223-3010203020020200-1223111110000012-0110321020220022-0123212133301001-3301102112333013-1323310031230323"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
header_transformation {
  # Configure direct properties listed below.
}
```

<a id="canonical-3303003333022013-2311101013010221-2310012011301021-0213111113020312-0231221011102330-1301311200322221-3131232201100330-3212030200300313"></a>

### Direct properties for `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation`

- [default_header_transformation](resources--workload--reference--group-019.md#canonical-0010231113030312-1033233101131032-3112201133013213-3303112230223301-2331322231322331-0301320223323310-3212003203010320-1100132021202010): complete subsection reference.

- [preserve_case_header_transformation](resources--workload--reference--group-019.md#canonical-0122312232023030-1012112022323310-2200001103311021-3013321333232030-0133230321231123-1322213203213000-3031310310221221-0023133221022231): complete subsection reference.

- [proper_case_header_transformation](resources--workload--reference--group-019.md#canonical-3231002321200010-2030232222102330-2133211301012130-2211121233102332-3130131112312303-0001021113312310-1220012102302231-2113230131033030): complete subsection reference.

<a id="canonical-0010231113030312-1033233101131032-3112201133013213-3303112230223301-2331322231322331-0301320223323310-3212003203010320-1100132021202010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.default_header_transformation` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-018.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-018.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_custom](resources--workload--reference--group-018.md#canonical-0203102001320123-1120130012130212-0111101220231333-3101331123203121-1110000311003232-1321231201202312-3200200021132112-0200303231330002)
- [stateful_service.advertise_options.advertise_custom.ports](resources--workload--reference--group-018.md#canonical-3111311322332303-1201030020021201-3133021233000313-2022333331131202-3221312032323033-3133120120113232-3330021120310331-2331303302032231)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-018.md#canonical-0313222022003221-2301103230203023-3120322210033100-2121220010332201-3133231021021220-1313001311100021-0313102021211212-0331322021202002)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-019.md#canonical-0112020200111022-0013103013103331-1110301001233101-1300202101312213-3301031210112102-1221232301300211-0231000232130131-3132113303133333)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options](resources--workload--reference--group-019.md#canonical-3202212001033033-3133231302102230-0333203123220020-0022212313320130-3032131111301011-1301121203330231-2010200210001321-0233311332222011)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only](resources--workload--reference--group-019.md#canonical-0101121203223101-2303112021302213-2210233131121322-1010121010210301-2330311311003110-2202230130203213-2331310302232000-1213011200013332)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation](resources--workload--reference--group-019.md#canonical-3020121301012012-1323210110303103-1332333110211232-3101112232311132-3303103303122230-0233302131022013-3203113021013212-1223222232202303)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.default_header_transformation

<a id="canonical-2131100113130231-0110302303022121-1213010031212022-1130110233110322-1103231303023131-0112313233000332-2323111113120211-1010033320121332"></a>

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0122312232023030-1012112022323310-2200001103311021-3013321333232030-0133230321231123-1322213203213000-3031310310221221-0023133221022231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.preserve_case_header_transformation` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-018.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-018.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_custom](resources--workload--reference--group-018.md#canonical-0203102001320123-1120130012130212-0111101220231333-3101331123203121-1110000311003232-1321231201202312-3200200021132112-0200303231330002)
- [stateful_service.advertise_options.advertise_custom.ports](resources--workload--reference--group-018.md#canonical-3111311322332303-1201030020021201-3133021233000313-2022333331131202-3221312032323033-3133120120113232-3330021120310331-2331303302032231)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-018.md#canonical-0313222022003221-2301103230203023-3120322210033100-2121220010332201-3133231021021220-1313001311100021-0313102021211212-0331322021202002)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-019.md#canonical-0112020200111022-0013103013103331-1110301001233101-1300202101312213-3301031210112102-1221232301300211-0231000232130131-3132113303133333)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options](resources--workload--reference--group-019.md#canonical-3202212001033033-3133231302102230-0333203123220020-0022212313320130-3032131111301011-1301121203330231-2010200210001321-0233311332222011)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only](resources--workload--reference--group-019.md#canonical-0101121203223101-2303112021302213-2210233131121322-1010121010210301-2330311311003110-2202230130203213-2331310302232000-1213011200013332)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation](resources--workload--reference--group-019.md#canonical-3020121301012012-1323210110303103-1332333110211232-3101112232311132-3303103303122230-0233302131022013-3203113021013212-1223222232202303)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.preserve_case_header_transformation

<a id="canonical-0133110322320301-2313113013313123-3011030312131001-3321110022230021-3000231211100131-3021030033222202-0320322123332220-1131300102123232"></a>

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3231002321200010-2030232222102330-2133211301012130-2211121233102332-3130131112312303-0001021113312310-1220012102302231-2113230131033030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.proper_case_header_transformation` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-018.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-018.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_custom](resources--workload--reference--group-018.md#canonical-0203102001320123-1120130012130212-0111101220231333-3101331123203121-1110000311003232-1321231201202312-3200200021132112-0200303231330002)
- [stateful_service.advertise_options.advertise_custom.ports](resources--workload--reference--group-018.md#canonical-3111311322332303-1201030020021201-3133021233000313-2022333331131202-3221312032323033-3133120120113232-3330021120310331-2331303302032231)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-018.md#canonical-0313222022003221-2301103230203023-3120322210033100-2121220010332201-3133231021021220-1313001311100021-0313102021211212-0331322021202002)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-019.md#canonical-0112020200111022-0013103013103331-1110301001233101-1300202101312213-3301031210112102-1221232301300211-0231000232130131-3132113303133333)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options](resources--workload--reference--group-019.md#canonical-3202212001033033-3133231302102230-0333203123220020-0022212313320130-3032131111301011-1301121203330231-2010200210001321-0233311332222011)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only](resources--workload--reference--group-019.md#canonical-0101121203223101-2303112021302213-2210233131121322-1010121010210301-2330311311003110-2202230130203213-2331310302232000-1213011200013332)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation](resources--workload--reference--group-019.md#canonical-3020121301012012-1323210110303103-1332333110211232-3101112232311132-3303103303122230-0233302131022013-3203113021013212-1223222232202303)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.proper_case_header_transformation

<a id="canonical-2331102100333133-0321211202333023-3133311001211303-3110123300021221-3103020013132310-1213300013331122-1322113320022023-3010303201121202"></a>

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1233012312003333-2101101012232120-0100310010232311-0011212101122211-2001031003213123-1031101223300032-3223031101001303-3201303031321011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_v2` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-018.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-018.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_custom](resources--workload--reference--group-018.md#canonical-0203102001320123-1120130012130212-0111101220231333-3101331123203121-1110000311003232-1321231201202312-3200200021132112-0200303231330002)
- [stateful_service.advertise_options.advertise_custom.ports](resources--workload--reference--group-018.md#canonical-3111311322332303-1201030020021201-3133021233000313-2022333331131202-3221312032323033-3133120120113232-3330021120310331-2331303302032231)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-018.md#canonical-0313222022003221-2301103230203023-3120322210033100-2121220010332201-3133231021021220-1313001311100021-0313102021211212-0331322021202002)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-019.md#canonical-0112020200111022-0013103013103331-1110301001233101-1300202101312213-3301031210112102-1221232301300211-0231000232130131-3132113303133333)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options](resources--workload--reference--group-019.md#canonical-3202212001033033-3133231302102230-0333203123220020-0022212313320130-3032131111301011-1301121203330231-2010200210001321-0233311332222011)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_v2

<a id="canonical-3232122132023320-2310202332000210-1221133012033130-2012102023022003-1202132100110322-3121022121232113-3130131200000231-1331303031332212"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
http_protocol_enable_v1_v2 = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2120212030201112-1220311010203002-0121211322222323-2032310203313301-2111110222212103-0222302013331310-1331023200202002-1220000133031023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v2_only` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-018.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-018.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_custom](resources--workload--reference--group-018.md#canonical-0203102001320123-1120130012130212-0111101220231333-3101331123203121-1110000311003232-1321231201202312-3200200021132112-0200303231330002)
- [stateful_service.advertise_options.advertise_custom.ports](resources--workload--reference--group-018.md#canonical-3111311322332303-1201030020021201-3133021233000313-2022333331131202-3221312032323033-3133120120113232-3330021120310331-2331303302032231)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-018.md#canonical-0313222022003221-2301103230203023-3120322210033100-2121220010332201-3133231021021220-1313001311100021-0313102021211212-0331322021202002)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-019.md#canonical-0112020200111022-0013103013103331-1110301001233101-1300202101312213-3301031210112102-1221232301300211-0231000232130131-3132113303133333)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options](resources--workload--reference--group-019.md#canonical-3202212001033033-3133231302102230-0333203123220020-0022212313320130-3032131111301011-1301121203330231-2010200210001321-0233311332222011)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v2_only

<a id="canonical-0110130011211221-1333210130313300-2231302200023013-1010323110330313-3233121311113232-3200313112302321-0312011303003212-2022003020010222"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
http_protocol_enable_v2_only = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1103323231113001-2101233200313132-0113203102011131-2310133213023120-2300020303303112-1311213222021323-3201102120300132-1030302220012022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.non_default_loadbalancer` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-018.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-018.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_custom](resources--workload--reference--group-018.md#canonical-0203102001320123-1120130012130212-0111101220231333-3101331123203121-1110000311003232-1321231201202312-3200200021132112-0200303231330002)
- [stateful_service.advertise_options.advertise_custom.ports](resources--workload--reference--group-018.md#canonical-3111311322332303-1201030020021201-3133021233000313-2022333331131202-3221312032323033-3133120120113232-3330021120310331-2331303302032231)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-018.md#canonical-0313222022003221-2301103230203023-3120322210033100-2121220010332201-3133231021021220-1313001311100021-0313102021211212-0331322021202002)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-019.md#canonical-0112020200111022-0013103013103331-1110301001233101-1300202101312213-3301031210112102-1221232301300211-0231000232130131-3132113303133333)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.non_default_loadbalancer

<a id="canonical-3032211130222100-1112133231231021-1132323102330032-0130213202002000-3031110302313313-2002121232031213-0221332011303220-0223121332103010"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
non_default_loadbalancer = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3122320112122220-2222123003111322-1032303200202320-0301311021210013-0010110320313023-3321020211021302-2302303311133010-3231010220310031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.pass_through` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-018.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-018.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_custom](resources--workload--reference--group-018.md#canonical-0203102001320123-1120130012130212-0111101220231333-3101331123203121-1110000311003232-1321231201202312-3200200021132112-0200303231330002)
- [stateful_service.advertise_options.advertise_custom.ports](resources--workload--reference--group-018.md#canonical-3111311322332303-1201030020021201-3133021233000313-2022333331131202-3221312032323033-3133120120113232-3330021120310331-2331303302032231)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-018.md#canonical-0313222022003221-2301103230203023-3120322210033100-2121220010332201-3133231021021220-1313001311100021-0313102021211212-0331322021202002)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-019.md#canonical-0112020200111022-0013103013103331-1110301001233101-1300202101312213-3301031210112102-1221232301300211-0231000232130131-3132113303133333)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.pass_through

<a id="canonical-3003323100213101-3220012012330230-3121001230332003-3222010132233323-1012330220311013-3322212231213222-1111331303201233-3100221320112220"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
pass_through = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1330323120222300-2123201010331320-0300232300332322-0010223222211002-1032302312300112-0233320002220301-1020331113013010-3202200223121221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-018.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-018.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_custom](resources--workload--reference--group-018.md#canonical-0203102001320123-1120130012130212-0111101220231333-3101331123203121-1110000311003232-1321231201202312-3200200021132112-0200303231330002)
- [stateful_service.advertise_options.advertise_custom.ports](resources--workload--reference--group-018.md#canonical-3111311322332303-1201030020021201-3133021233000313-2022333331131202-3221312032323033-3133120120113232-3330021120310331-2331303302032231)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-018.md#canonical-0313222022003221-2301103230203023-3120322210033100-2121220010332201-3133231021021220-1313001311100021-0313102021211212-0331322021202002)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-019.md#canonical-0112020200111022-0013103013103331-1110301001233101-1300202101312213-3301031210112102-1221232301300211-0231000232130131-3132113303133333)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params

<a id="canonical-2013311010020320-2121113203323031-0200002033113131-0312102310321003-1232023020001013-1303232221320223-1300230002113120-2202133130120031"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
tls_cert_params {
  # Configure direct properties listed below.
}
```

<a id="canonical-2312210020320113-0110232123202223-1200202221310021-2313230111201022-2200101203022023-3030133301312320-1222303002323232-0333202020321323"></a>

### Direct properties for `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params`

- [certificates](resources--workload--reference--group-019.md#canonical-2302321201330300-1112012230113032-0022112020230310-3222302030221312-2011131033323122-1230110201111231-3102300001021111-1202003110201010): complete subsection reference.

- [no_mtls](resources--workload--reference--group-019.md#canonical-3201210010322203-1131010200302302-1323313133302012-2032121310301010-1331220210000022-2132221211132332-1313000211301001-0210022102033110): complete subsection reference.

- [tls_config](resources--workload--reference--group-019.md#canonical-0111002012300220-0123100132331012-0322301201120202-3013312031133201-3021203131320032-0333213333200111-2120023323002022-1203103032321031): complete subsection reference.

- [use_mtls](resources--workload--reference--group-019.md#canonical-1232000222132210-1231321130032102-2312300020102111-2032003122222202-0033301200101032-3330213032031032-2302102033201003-3021232302102100): complete subsection reference.

<a id="canonical-2302321201330300-1112012230113032-0022112020230310-3222302030221312-2011131033323122-1230110201111231-3102300001021111-1202003110201010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.certificates` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-018.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-018.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_custom](resources--workload--reference--group-018.md#canonical-0203102001320123-1120130012130212-0111101220231333-3101331123203121-1110000311003232-1321231201202312-3200200021132112-0200303231330002)
- [stateful_service.advertise_options.advertise_custom.ports](resources--workload--reference--group-018.md#canonical-3111311322332303-1201030020021201-3133021233000313-2022333331131202-3221312032323033-3133120120113232-3330021120310331-2331303302032231)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-018.md#canonical-0313222022003221-2301103230203023-3120322210033100-2121220010332201-3133231021021220-1313001311100021-0313102021211212-0331322021202002)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-019.md#canonical-0112020200111022-0013103013103331-1110301001233101-1300202101312213-3301031210112102-1221232301300211-0231000232130131-3132113303133333)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params](resources--workload--reference--group-019.md#canonical-1330323120222300-2123201010331320-0300232300332322-0010223222211002-1032302312300112-0233320002220301-1020331113013010-3202200223121221)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.certificates

<a id="canonical-0201121111210100-0200012131011222-2120033303323132-0301030021231123-0130003111322332-2100100113012133-1320110123020220-2000302112313322"></a>

Type: `"object"`. list nested block, Optional.

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

<a id="canonical-1120323113223010-0222330233330223-1203321001212003-0303113201203231-1033310103333223-2101032231131133-1021002230300000-2030332020302230"></a>

### Direct properties for `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.certificates`

<a id="canonical-1102121100221310-1132233333131331-3023321201120332-1033322132322113-0113110020011011-0002002321232103-0020121313212230-0113310223010203"></a>

#### `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.certificates.name` property

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

<a id="canonical-1032112101323332-0033012000013202-1000321222033011-2232333312012300-3321233231101222-3123131131301103-3110121000330202-2301013030022230"></a>

<a id="canonical-0030000120300230-2031003032002120-0133101312233003-0111233131230211-3301331333130103-3133110132130113-3220031011303100-0102312212113120"></a>

#### `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.certificates.namespace` property

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

<a id="canonical-2303331333212212-3221121001230000-3133122231211102-3231201230333012-0210010032210313-2032203321322311-3013220121131313-1201010310013212"></a>

<a id="canonical-2002130230331223-0311122122112311-2122332022102313-3312111322030010-2133313223002303-3213222300031210-0101031132303302-3201222202223311"></a>

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

<a id="canonical-3201210010322203-1131010200302302-1323313133302012-2032121310301010-1331220210000022-2132221211132332-1313000211301001-0210022102033110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.no_mtls` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-018.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-018.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_custom](resources--workload--reference--group-018.md#canonical-0203102001320123-1120130012130212-0111101220231333-3101331123203121-1110000311003232-1321231201202312-3200200021132112-0200303231330002)
- [stateful_service.advertise_options.advertise_custom.ports](resources--workload--reference--group-018.md#canonical-3111311322332303-1201030020021201-3133021233000313-2022333331131202-3221312032323033-3133120120113232-3330021120310331-2331303302032231)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-018.md#canonical-0313222022003221-2301103230203023-3120322210033100-2121220010332201-3133231021021220-1313001311100021-0313102021211212-0331322021202002)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-019.md#canonical-0112020200111022-0013103013103331-1110301001233101-1300202101312213-3301031210112102-1221232301300211-0231000232130131-3132113303133333)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params](resources--workload--reference--group-019.md#canonical-1330323120222300-2123201010331320-0300232300332322-0010223222211002-1032302312300112-0233320002220301-1020331113013010-3202200223121221)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.no_mtls

<a id="canonical-2122220133223111-3220220211011110-3330110031130220-2100211312030303-0320001302013310-1011310112021021-1230300220202030-1110233130110302"></a>

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
no_mtls = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0111002012300220-0123100132331012-0322301201120202-3013312031133201-3021203131320032-0333213333200111-2120023323002022-1203103032321031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.tls_config` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-018.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-018.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_custom](resources--workload--reference--group-018.md#canonical-0203102001320123-1120130012130212-0111101220231333-3101331123203121-1110000311003232-1321231201202312-3200200021132112-0200303231330002)
- [stateful_service.advertise_options.advertise_custom.ports](resources--workload--reference--group-018.md#canonical-3111311322332303-1201030020021201-3133021233000313-2022333331131202-3221312032323033-3133120120113232-3330021120310331-2331303302032231)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-018.md#canonical-0313222022003221-2301103230203023-3120322210033100-2121220010332201-3133231021021220-1313001311100021-0313102021211212-0331322021202002)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-019.md#canonical-0112020200111022-0013103013103331-1110301001233101-1300202101312213-3301031210112102-1221232301300211-0231000232130131-3132113303133333)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params](resources--workload--reference--group-019.md#canonical-1330323120222300-2123201010331320-0300232300332322-0010223222211002-1032302312300112-0233320002220301-1020331113013010-3202200223121221)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.tls_config

<a id="canonical-1330121113200001-0121232310331130-1023030122233102-2022222223011132-2230030011131320-3303220033303201-1200131312102033-1201310202111312"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
tls_config {
  # Configure direct properties listed below.
}
```

<a id="canonical-0003213032011103-3210202201110322-1033303301010311-2022113230212010-1020032123012220-2000310331023033-0332220313130332-0201300221132331"></a>

### Direct properties for `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.tls_config`

- [custom_security](resources--workload--reference--group-019.md#canonical-0213222020222022-2033230011011000-3133220133331122-2313032021200230-1233322200222033-1202020120332300-3221303013103010-1321233123110113): complete subsection reference.

- [default_security](resources--workload--reference--group-019.md#canonical-3113133100131011-1010333021003230-1123210303221313-1323213322031002-2013211001200120-1123330132001213-0132320122300003-0320201222013110): complete subsection reference.

- [low_security](resources--workload--reference--group-019.md#canonical-0213030220123100-3120123030111302-2220121322132023-3231112302000021-2030031312331030-3322131203122302-0220131033022120-2131002102103231): complete subsection reference.

- [medium_security](resources--workload--reference--group-019.md#canonical-3333323101212103-0100322011331322-0122123210000023-0320312011231232-2120011020130202-0223332023111032-1200223021123321-1001112202313303): complete subsection reference.

<a id="canonical-0213222020222022-2033230011011000-3133220133331122-2313032021200230-1233322200222033-1202020120332300-3221303013103010-1321233123110113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.tls_config.custom_security` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-018.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-018.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_custom](resources--workload--reference--group-018.md#canonical-0203102001320123-1120130012130212-0111101220231333-3101331123203121-1110000311003232-1321231201202312-3200200021132112-0200303231330002)
- [stateful_service.advertise_options.advertise_custom.ports](resources--workload--reference--group-018.md#canonical-3111311322332303-1201030020021201-3133021233000313-2022333331131202-3221312032323033-3133120120113232-3330021120310331-2331303302032231)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-018.md#canonical-0313222022003221-2301103230203023-3120322210033100-2121220010332201-3133231021021220-1313001311100021-0313102021211212-0331322021202002)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-019.md#canonical-0112020200111022-0013103013103331-1110301001233101-1300202101312213-3301031210112102-1221232301300211-0231000232130131-3132113303133333)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params](resources--workload--reference--group-019.md#canonical-1330323120222300-2123201010331320-0300232300332322-0010223222211002-1032302312300112-0233320002220301-1020331113013010-3202200223121221)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.tls_config](resources--workload--reference--group-019.md#canonical-0111002012300220-0123100132331012-0322301201120202-3013312031133201-3021203131320032-0333213333200111-2120023323002022-1203103032321031)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.tls_config.custom_security

<a id="canonical-0030022202022032-0123331111102011-0113102120202322-0023213302201331-3110330131200323-2103030111223001-3300002033120300-0132122322313201"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
custom_security {
  # Configure direct properties listed below.
}
```

<a id="canonical-0001323231110030-1220031112200313-2330000231121120-1220022003321103-1030020222201311-3222222111010101-3023202322130221-2132012112302300"></a>

### Direct properties for `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.tls_config.custom_security`

<a id="canonical-3011330023302003-0030121312021303-2023032112023302-3312200322023013-0232121032123200-2011120023230000-3020223230212321-2001222022222323"></a>

#### `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.tls_config.custom_security.cipher_suites` property

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

<a id="canonical-3023321230220101-2000011132122230-3220333001032012-0113300022031231-3012010231233013-1320102002003212-2101132102312301-0212130100203313"></a>

<a id="canonical-2020301132101202-3100031233322132-1323320120223322-3030003133131212-2311000213311320-2002120121121002-3332032211132330-2013301230120000"></a>

#### `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.tls_config.custom_security.max_version` property

Type: `"string"`. Optional.

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

<a id="canonical-2001313221103323-1222010332312230-2120231110223232-3210210032310210-2003130110220232-1202022332200030-2022310333120321-3033012230012131"></a>

<a id="canonical-1230020232210133-0123003332321330-1011111321220030-2133223300231002-3320321033000122-2333010112002303-1032323210313111-1121201322333020"></a>

#### `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.tls_config.custom_security.min_version` property

Type: `"string"`. Optional.

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

<a id="canonical-3113133100131011-1010333021003230-1123210303221313-1323213322031002-2013211001200120-1123330132001213-0132320122300003-0320201222013110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.tls_config.default_security` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-018.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-018.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_custom](resources--workload--reference--group-018.md#canonical-0203102001320123-1120130012130212-0111101220231333-3101331123203121-1110000311003232-1321231201202312-3200200021132112-0200303231330002)
- [stateful_service.advertise_options.advertise_custom.ports](resources--workload--reference--group-018.md#canonical-3111311322332303-1201030020021201-3133021233000313-2022333331131202-3221312032323033-3133120120113232-3330021120310331-2331303302032231)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-018.md#canonical-0313222022003221-2301103230203023-3120322210033100-2121220010332201-3133231021021220-1313001311100021-0313102021211212-0331322021202002)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-019.md#canonical-0112020200111022-0013103013103331-1110301001233101-1300202101312213-3301031210112102-1221232301300211-0231000232130131-3132113303133333)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params](resources--workload--reference--group-019.md#canonical-1330323120222300-2123201010331320-0300232300332322-0010223222211002-1032302312300112-0233320002220301-1020331113013010-3202200223121221)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.tls_config](resources--workload--reference--group-019.md#canonical-0111002012300220-0123100132331012-0322301201120202-3013312031133201-3021203131320032-0333213333200111-2120023323002022-1203103032321031)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.tls_config.default_security

<a id="canonical-1120332330303222-1211022002233011-2212001221321301-2001201030221110-1031003032332001-0123203021223301-0130023333232321-1330333131133033"></a>

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
default_security = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0213030220123100-3120123030111302-2220121322132023-3231112302000021-2030031312331030-3322131203122302-0220131033022120-2131002102103231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.tls_config.low_security` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-018.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-018.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_custom](resources--workload--reference--group-018.md#canonical-0203102001320123-1120130012130212-0111101220231333-3101331123203121-1110000311003232-1321231201202312-3200200021132112-0200303231330002)
- [stateful_service.advertise_options.advertise_custom.ports](resources--workload--reference--group-018.md#canonical-3111311322332303-1201030020021201-3133021233000313-2022333331131202-3221312032323033-3133120120113232-3330021120310331-2331303302032231)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-018.md#canonical-0313222022003221-2301103230203023-3120322210033100-2121220010332201-3133231021021220-1313001311100021-0313102021211212-0331322021202002)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-019.md#canonical-0112020200111022-0013103013103331-1110301001233101-1300202101312213-3301031210112102-1221232301300211-0231000232130131-3132113303133333)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params](resources--workload--reference--group-019.md#canonical-1330323120222300-2123201010331320-0300232300332322-0010223222211002-1032302312300112-0233320002220301-1020331113013010-3202200223121221)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.tls_config](resources--workload--reference--group-019.md#canonical-0111002012300220-0123100132331012-0322301201120202-3013312031133201-3021203131320032-0333213333200111-2120023323002022-1203103032321031)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.tls_config.low_security

<a id="canonical-3000002113321322-2200020133030123-3220023002130133-0322031331211321-1303302023311332-3101230312132121-2021332322202013-0231330331210323"></a>

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
low_security = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3333323101212103-0100322011331322-0122123210000023-0320312011231232-2120011020130202-0223332023111032-1200223021123321-1001112202313303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.tls_config.medium_security` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-018.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-018.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_custom](resources--workload--reference--group-018.md#canonical-0203102001320123-1120130012130212-0111101220231333-3101331123203121-1110000311003232-1321231201202312-3200200021132112-0200303231330002)
- [stateful_service.advertise_options.advertise_custom.ports](resources--workload--reference--group-018.md#canonical-3111311322332303-1201030020021201-3133021233000313-2022333331131202-3221312032323033-3133120120113232-3330021120310331-2331303302032231)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-018.md#canonical-0313222022003221-2301103230203023-3120322210033100-2121220010332201-3133231021021220-1313001311100021-0313102021211212-0331322021202002)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-019.md#canonical-0112020200111022-0013103013103331-1110301001233101-1300202101312213-3301031210112102-1221232301300211-0231000232130131-3132113303133333)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params](resources--workload--reference--group-019.md#canonical-1330323120222300-2123201010331320-0300232300332322-0010223222211002-1032302312300112-0233320002220301-1020331113013010-3202200223121221)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.tls_config](resources--workload--reference--group-019.md#canonical-0111002012300220-0123100132331012-0322301201120202-3013312031133201-3021203131320032-0333213333200111-2120023323002022-1203103032321031)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.tls_config.medium_security

<a id="canonical-2120002112210330-1120221222300022-2201212322301111-1223001333213230-1233031330312310-1320022001032331-3311030203210332-0013331202330113"></a>

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
medium_security = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1232000222132210-1231321130032102-2312300020102111-2032003122222202-0033301200101032-3330213032031032-2302102033201003-3021232302102100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.use_mtls` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-018.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-018.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_custom](resources--workload--reference--group-018.md#canonical-0203102001320123-1120130012130212-0111101220231333-3101331123203121-1110000311003232-1321231201202312-3200200021132112-0200303231330002)
- [stateful_service.advertise_options.advertise_custom.ports](resources--workload--reference--group-018.md#canonical-3111311322332303-1201030020021201-3133021233000313-2022333331131202-3221312032323033-3133120120113232-3330021120310331-2331303302032231)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-018.md#canonical-0313222022003221-2301103230203023-3120322210033100-2121220010332201-3133231021021220-1313001311100021-0313102021211212-0331322021202002)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-019.md#canonical-0112020200111022-0013103013103331-1110301001233101-1300202101312213-3301031210112102-1221232301300211-0231000232130131-3132113303133333)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params](resources--workload--reference--group-019.md#canonical-1330323120222300-2123201010331320-0300232300332322-0010223222211002-1032302312300112-0233320002220301-1020331113013010-3202200223121221)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.use_mtls

<a id="canonical-0200231003322122-0230312323230201-3102330221133103-2121230332220211-0332123200012023-2203231013112201-1210210230122033-0133013330010313"></a>

Type: `"object"`. single nested block, Optional.

Validation context for downstream client TLS connections.

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

<a id="canonical-3301003001103031-1003023123323203-2213310032133331-3102013110232211-3010133310223113-2131123012200032-3231320301300223-0032331121303230"></a>

### Direct properties for `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.use_mtls`

<a id="canonical-1131032211210233-0312011021320213-0223133032310020-1223300131310023-1001100313332102-0230121100032000-1110320101222010-2122311302332303"></a>

#### `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.use_mtls.client_certificate_optional` property

Type: `"bool"`. Optional.

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

- [crl](resources--workload--reference--group-019.md#canonical-1023310123331030-3330002101023131-2133232213203312-1203130332202311-2100220033333301-3323102331333230-1211233121310122-0121013103030011): complete subsection reference.

- [no_crl](resources--workload--reference--group-019.md#canonical-2233232331221312-2213100000032221-0112211013300212-2030330022220310-1331233012301211-2010110032211230-3223302023221310-3002312000203332): complete subsection reference.

- [trusted_ca](resources--workload--reference--group-019.md#canonical-2123331130310033-0203031003112001-0331233220320302-3230103032202121-3311302330231202-3002332101200103-3003111103223202-0222221332311201): complete subsection reference.

<a id="canonical-1222233131000331-3300123122231212-1022112232101120-1122011121033101-3100110122102231-2030312130221031-0013203133031000-1320310001313232"></a>

<a id="canonical-3012130210321300-2310021000001212-2200213232221030-2021231002012222-0100110202012213-0320223320010023-0131101301201001-2023212102202300"></a>

#### `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.use_mtls.trusted_ca_url` property

Type: `"string"`. Optional.

Exclusive with \[trusted\_ca\] Upload a Root CA Certificate specifically for this Load Balancer.

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

- [xfcc_disabled](resources--workload--reference--group-019.md#canonical-2113333000131311-3302120001323220-1200201202012001-1210000301020103-3310002113310230-3013100321100011-0231203223000031-3313332002011331): complete subsection reference.

- [xfcc_options](resources--workload--reference--group-019.md#canonical-3133132021133213-3020123203123012-3203121321220120-1000031311310301-2322312002221310-3110133230331222-1103100011120133-0000230331013210): complete subsection reference.

<a id="canonical-1023310123331030-3330002101023131-2133232213203312-1203130332202311-2100220033333301-3323102331333230-1211233121310122-0121013103030011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.use_mtls.crl` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-018.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-018.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_custom](resources--workload--reference--group-018.md#canonical-0203102001320123-1120130012130212-0111101220231333-3101331123203121-1110000311003232-1321231201202312-3200200021132112-0200303231330002)
- [stateful_service.advertise_options.advertise_custom.ports](resources--workload--reference--group-018.md#canonical-3111311322332303-1201030020021201-3133021233000313-2022333331131202-3221312032323033-3133120120113232-3330021120310331-2331303302032231)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-018.md#canonical-0313222022003221-2301103230203023-3120322210033100-2121220010332201-3133231021021220-1313001311100021-0313102021211212-0331322021202002)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-019.md#canonical-0112020200111022-0013103013103331-1110301001233101-1300202101312213-3301031210112102-1221232301300211-0231000232130131-3132113303133333)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params](resources--workload--reference--group-019.md#canonical-1330323120222300-2123201010331320-0300232300332322-0010223222211002-1032302312300112-0233320002220301-1020331113013010-3202200223121221)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.use_mtls](resources--workload--reference--group-019.md#canonical-1232000222132210-1231321130032102-2312300020102111-2032003122222202-0033301200101032-3330213032031032-2302102033201003-3021232302102100)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.use_mtls.crl

<a id="canonical-2230203333012010-3231102212313302-1013322220223002-0312203032323320-1023030110102233-0033231132231211-1303103012310330-0220312322301333"></a>

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
crl {
  # Configure direct properties listed below.
}
```

<a id="canonical-1321012310323123-2012002112120333-1201203203332320-1102030032032202-2210101221001230-0302132321330032-1133003301130003-2200131220211233"></a>

### Direct properties for `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.use_mtls.crl`

<a id="canonical-2321122310000133-0131130101331030-3102020220222102-3031313102020222-1321101130221213-1210303010033033-1020310030001123-3012330110112030"></a>

#### `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.use_mtls.crl.name` property

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

<a id="canonical-3302230031301102-1300311301200320-3221120002111012-3012313221230300-0111222302132120-0022211232132000-0200232101133232-2211321012300033"></a>

<a id="canonical-3020332113013000-0102131222010033-1000030100110031-3000221011213230-1133311130233200-1310331012221011-1221020022231102-3000100121113020"></a>

#### `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.use_mtls.crl.namespace` property

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

<a id="canonical-2211030221310232-0102111113321321-3231113132231012-3231323022213223-3320202002122123-1302033310333301-3010020102100102-3110300120333033"></a>

<a id="canonical-0331003131101212-0133132021310101-2030020321322100-3130001000100011-3302332133330313-1121031311122322-2023110331332001-2211030013002000"></a>

#### `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.use_mtls.crl.tenant` property

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

<a id="canonical-2233232331221312-2213100000032221-0112211013300212-2030330022220310-1331233012301211-2010110032211230-3223302023221310-3002312000203332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.use_mtls.no_crl` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-018.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-018.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_custom](resources--workload--reference--group-018.md#canonical-0203102001320123-1120130012130212-0111101220231333-3101331123203121-1110000311003232-1321231201202312-3200200021132112-0200303231330002)
- [stateful_service.advertise_options.advertise_custom.ports](resources--workload--reference--group-018.md#canonical-3111311322332303-1201030020021201-3133021233000313-2022333331131202-3221312032323033-3133120120113232-3330021120310331-2331303302032231)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-018.md#canonical-0313222022003221-2301103230203023-3120322210033100-2121220010332201-3133231021021220-1313001311100021-0313102021211212-0331322021202002)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-019.md#canonical-0112020200111022-0013103013103331-1110301001233101-1300202101312213-3301031210112102-1221232301300211-0231000232130131-3132113303133333)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params](resources--workload--reference--group-019.md#canonical-1330323120222300-2123201010331320-0300232300332322-0010223222211002-1032302312300112-0233320002220301-1020331113013010-3202200223121221)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.use_mtls](resources--workload--reference--group-019.md#canonical-1232000222132210-1231321130032102-2312300020102111-2032003122222202-0033301200101032-3330213032031032-2302102033201003-3021232302102100)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.use_mtls.no_crl

<a id="canonical-1032012311202203-1012303311222232-1201021102130231-0020033303203321-1321102133300202-0232310103010112-1323221230232030-2030202222223021"></a>

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
no_crl = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2123331130310033-0203031003112001-0331233220320302-3230103032202121-3311302330231202-3002332101200103-3003111103223202-0222221332311201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.use_mtls.trusted_ca` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-018.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-018.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_custom](resources--workload--reference--group-018.md#canonical-0203102001320123-1120130012130212-0111101220231333-3101331123203121-1110000311003232-1321231201202312-3200200021132112-0200303231330002)
- [stateful_service.advertise_options.advertise_custom.ports](resources--workload--reference--group-018.md#canonical-3111311322332303-1201030020021201-3133021233000313-2022333331131202-3221312032323033-3133120120113232-3330021120310331-2331303302032231)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-018.md#canonical-0313222022003221-2301103230203023-3120322210033100-2121220010332201-3133231021021220-1313001311100021-0313102021211212-0331322021202002)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-019.md#canonical-0112020200111022-0013103013103331-1110301001233101-1300202101312213-3301031210112102-1221232301300211-0231000232130131-3132113303133333)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params](resources--workload--reference--group-019.md#canonical-1330323120222300-2123201010331320-0300232300332322-0010223222211002-1032302312300112-0233320002220301-1020331113013010-3202200223121221)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.use_mtls](resources--workload--reference--group-019.md#canonical-1232000222132210-1231321130032102-2312300020102111-2032003122222202-0033301200101032-3330213032031032-2302102033201003-3021232302102100)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.use_mtls.trusted_ca

<a id="canonical-1122313111020022-0223202120313130-3213022010133302-0023322122301201-2232213211003010-3301112330033302-2030111212102220-1000210113130303"></a>

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
trusted_ca {
  # Configure direct properties listed below.
}
```

<a id="canonical-3300002012312013-3230013012012021-1231301221310033-2300302122230103-3132113120220220-3003300333222100-2211233122301023-3323000020333011"></a>

### Direct properties for `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.use_mtls.trusted_ca`

<a id="canonical-2001032230222003-1203211323032332-2331102213013121-0310012213332221-1120023011120131-2311112010103312-0320213303323313-2201321221321123"></a>

#### `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.use_mtls.trusted_ca.name` property

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

<a id="canonical-3130202023332312-3132010130011322-2003202220301010-1101233313120120-3113032221302330-2322302313101222-2301031012212321-3110312131333232"></a>

<a id="canonical-0323203021001210-1300132331203322-0303003113300130-1023031100132210-2320321120032003-1203333201030322-2023113233310213-0002010201332022"></a>

#### `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.use_mtls.trusted_ca.namespace` property

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

<a id="canonical-2221322223101212-1232132201202231-2102322232302130-3221313322011001-0201312030333303-0302102021130121-0330132021331113-3302330332033002"></a>

<a id="canonical-1220112001032103-1023310213323232-1303332130123101-3331223211002302-1122031310201032-2133112130022221-2323232001202012-3122322001311022"></a>

#### `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.use_mtls.trusted_ca.tenant` property

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

<a id="canonical-2113333000131311-3302120001323220-1200201202012001-1210000301020103-3310002113310230-3013100321100011-0231203223000031-3313332002011331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.use_mtls.xfcc_disabled` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-018.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-018.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_custom](resources--workload--reference--group-018.md#canonical-0203102001320123-1120130012130212-0111101220231333-3101331123203121-1110000311003232-1321231201202312-3200200021132112-0200303231330002)
- [stateful_service.advertise_options.advertise_custom.ports](resources--workload--reference--group-018.md#canonical-3111311322332303-1201030020021201-3133021233000313-2022333331131202-3221312032323033-3133120120113232-3330021120310331-2331303302032231)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-018.md#canonical-0313222022003221-2301103230203023-3120322210033100-2121220010332201-3133231021021220-1313001311100021-0313102021211212-0331322021202002)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-019.md#canonical-0112020200111022-0013103013103331-1110301001233101-1300202101312213-3301031210112102-1221232301300211-0231000232130131-3132113303133333)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params](resources--workload--reference--group-019.md#canonical-1330323120222300-2123201010331320-0300232300332322-0010223222211002-1032302312300112-0233320002220301-1020331113013010-3202200223121221)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.use_mtls](resources--workload--reference--group-019.md#canonical-1232000222132210-1231321130032102-2312300020102111-2032003122222202-0033301200101032-3330213032031032-2302102033201003-3021232302102100)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.use_mtls.xfcc_disabled

<a id="canonical-2013230222330223-1232310232031303-2210223232111230-3023313032021203-3202113100130030-3322100032223332-3102222123300200-3011113132220111"></a>

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
xfcc_disabled = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3133132021133213-3020123203123012-3203121321220120-1000031311310301-2322312002221310-3110133230331222-1103100011120133-0000230331013210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.use_mtls.xfcc_options` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-018.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-018.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_custom](resources--workload--reference--group-018.md#canonical-0203102001320123-1120130012130212-0111101220231333-3101331123203121-1110000311003232-1321231201202312-3200200021132112-0200303231330002)
- [stateful_service.advertise_options.advertise_custom.ports](resources--workload--reference--group-018.md#canonical-3111311322332303-1201030020021201-3133021233000313-2022333331131202-3221312032323033-3133120120113232-3330021120310331-2331303302032231)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-018.md#canonical-0313222022003221-2301103230203023-3120322210033100-2121220010332201-3133231021021220-1313001311100021-0313102021211212-0331322021202002)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-019.md#canonical-0112020200111022-0013103013103331-1110301001233101-1300202101312213-3301031210112102-1221232301300211-0231000232130131-3132113303133333)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params](resources--workload--reference--group-019.md#canonical-1330323120222300-2123201010331320-0300232300332322-0010223222211002-1032302312300112-0233320002220301-1020331113013010-3202200223121221)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.use_mtls](resources--workload--reference--group-019.md#canonical-1232000222132210-1231321130032102-2312300020102111-2032003122222202-0033301200101032-3330213032031032-2302102033201003-3021232302102100)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.use_mtls.xfcc_options

<a id="canonical-0023100130303320-3111213311322223-1222231020023313-1032333210231030-0132132122101302-2322303013222203-1232320001001233-0222220022301103"></a>

Type: `"object"`. single nested block, Optional.

X-Forwarded-Client-Cert header elements to be added to requests.

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
xfcc_options {
  # Configure direct properties listed below.
}
```

<a id="canonical-3303211333230103-0201020033112103-0133213333122221-2033013021101122-3010310213110111-1133030110202133-1111103003113111-2113222201031231"></a>

### Direct properties for `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.use_mtls.xfcc_options`

<a id="canonical-0321101211132130-2323123222033122-1103021101221333-2300311330111001-2230231231312021-0201112012132101-1120320203210012-0330222110232311"></a>

#### `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.use_mtls.xfcc_options.xfcc_header_elements` property

Type: `["list", "string"]`. Optional.

\[Enum: XFCC\_NONE|XFCC\_CERT|XFCC\_CHAIN|XFCC\_SUBJECT|XFCC\_URI|XFCC\_DNS\]
X-Forwarded-Client-Cert header elements to be added to requests. Possible values are \`XFCC\_NONE\`,
\`XFCC\_CERT\`, \`XFCC\_CHAIN\`, \`XFCC\_SUBJECT\`, \`XFCC\_URI\`, \`XFCC\_DNS\`. Defaults to
\`XFCC\_NONE\`.

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
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.items.enum.not_in": "[0]"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.items.enum.not_in": "[0]"
  }
}
```

<a id="canonical-2020320003013033-2231100311132300-2132132133022132-3002120002132203-3131231201330331-0231012122032023-1202211010110230-0230013122122221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-018.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-018.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_custom](resources--workload--reference--group-018.md#canonical-0203102001320123-1120130012130212-0111101220231333-3101331123203121-1110000311003232-1321231201202312-3200200021132112-0200303231330002)
- [stateful_service.advertise_options.advertise_custom.ports](resources--workload--reference--group-018.md#canonical-3111311322332303-1201030020021201-3133021233000313-2022333331131202-3221312032323033-3133120120113232-3330021120310331-2331303302032231)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-018.md#canonical-0313222022003221-2301103230203023-3120322210033100-2121220010332201-3133231021021220-1313001311100021-0313102021211212-0331322021202002)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-019.md#canonical-0112020200111022-0013103013103331-1110301001233101-1300202101312213-3301031210112102-1221232301300211-0231000232130131-3132113303133333)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters

<a id="canonical-3012320331111130-3320032303011332-3210332231220013-0033203303111011-0320003331320201-0200102302301013-0200000133221302-3210323313131221"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for tls parameters.

Additional upstream details:

Inline TLS parameters.

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
tls_parameters {
  # Configure direct properties listed below.
}
```

<a id="canonical-3021220313010201-1223103122123332-1031102111313231-2122011022023201-2213103221131232-2133032132320013-1310322121312203-2210313202013202"></a>

### Direct properties for `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters`

- [no_mtls](resources--workload--reference--group-019.md#canonical-2033222013223233-0320010131321230-3001211111320031-2013111220310333-2201211223101123-2233002331000331-0130220100213002-1313101033313212): complete subsection reference.

- [tls_certificates](resources--workload--reference--group-019.md#canonical-3332313302002330-1333101212032003-2203023132232213-0001033110210120-3121103210011323-2020220120012213-3220330021330022-1301300222033110): complete subsection reference.

- [tls_config](resources--workload--reference--group-019.md#canonical-3212132103211102-3103133011133021-3313221110320101-2111103132233113-1232201321022230-0232111013133132-1231122000001132-3012311030121223): complete subsection reference.

- [use_mtls](resources--workload--reference--group-020.md#canonical-0002002033132331-3111321121121321-2300013122120330-2032012003011010-3002010130131223-0102231301201330-0023121332032121-3203220222232210): complete subsection reference.

<a id="canonical-2033222013223233-0320010131321230-3001211111320031-2013111220310333-2201211223101123-2233002331000331-0130220100213002-1313101033313212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.no_mtls` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-018.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-018.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_custom](resources--workload--reference--group-018.md#canonical-0203102001320123-1120130012130212-0111101220231333-3101331123203121-1110000311003232-1321231201202312-3200200021132112-0200303231330002)
- [stateful_service.advertise_options.advertise_custom.ports](resources--workload--reference--group-018.md#canonical-3111311322332303-1201030020021201-3133021233000313-2022333331131202-3221312032323033-3133120120113232-3330021120310331-2331303302032231)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-018.md#canonical-0313222022003221-2301103230203023-3120322210033100-2121220010332201-3133231021021220-1313001311100021-0313102021211212-0331322021202002)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-019.md#canonical-0112020200111022-0013103013103331-1110301001233101-1300202101312213-3301031210112102-1221232301300211-0231000232130131-3132113303133333)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-019.md#canonical-2020320003013033-2231100311132300-2132132133022132-3002120002132203-3131231201330331-0231012122032023-1202211010110230-0230013122122221)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.no_mtls

<a id="canonical-1320330130121223-1032212223333201-3011012132102302-3202200020301201-0223112200332020-3223300112012130-3303311122322021-2133133022003020"></a>

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
no_mtls = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3332313302002330-1333101212032003-2203023132232213-0001033110210120-3121103210011323-2020220120012213-3220330021330022-1301300222033110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-018.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-018.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_custom](resources--workload--reference--group-018.md#canonical-0203102001320123-1120130012130212-0111101220231333-3101331123203121-1110000311003232-1321231201202312-3200200021132112-0200303231330002)
- [stateful_service.advertise_options.advertise_custom.ports](resources--workload--reference--group-018.md#canonical-3111311322332303-1201030020021201-3133021233000313-2022333331131202-3221312032323033-3133120120113232-3330021120310331-2331303302032231)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-018.md#canonical-0313222022003221-2301103230203023-3120322210033100-2121220010332201-3133231021021220-1313001311100021-0313102021211212-0331322021202002)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-019.md#canonical-0112020200111022-0013103013103331-1110301001233101-1300202101312213-3301031210112102-1221232301300211-0231000232130131-3132113303133333)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-019.md#canonical-2020320003013033-2231100311132300-2132132133022132-3002120002132203-3131231201330331-0231012122032023-1202211010110230-0230013122122221)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates

<a id="canonical-0000202233313232-3310231000020221-2113121311203110-2313010331022330-0111010012233330-1032332113131122-1302210011112001-2020133212111213"></a>

Type: `"object"`. list nested block, Optional.

Users can add one or more certificates that share the same set of domains. For example, domain.com
and \*.domain.com - but use different signature algorithms.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 16,
  "minItems": 1,
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
    "minItems": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

Terraform syntax:

```terraform
tls_certificates {
  # Configure direct properties listed below.
}
```

<a id="canonical-3232112013302102-1231020222223031-1211022303112121-1313031101312302-0120313310110301-1322230221300202-0131212231201212-1213030213320232"></a>

### Direct properties for `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates`

- [blindfold](resources--workload--reference--group-019.md#canonical-0311333301123130-3101332300002121-3010331213132202-0223222131101322-1123330112003000-2133011022302032-1130000212313030-1321200033013210): complete subsection reference.

<a id="canonical-2203023033022223-3132313210333231-3111113313203020-2121330033103322-3333001110122233-3001201213200332-1332110202313010-0202133311321132"></a>

<a id="canonical-0213103000210133-0233201031320130-3123323333002013-3122200132313020-1301303123100323-2001031021132032-3121311331223023-0011331123220102"></a>

#### `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates.certificate_url` property

Type: `"string"`. Optional, Computed.

TLS certificate. Certificate or certificate chain in PEM format including the PEM headers.

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
    "ves.io.schema.rules.string.certificate_url": "true",
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.min_bytes": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.certificate_url": "true",
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.min_bytes": "1"
  }
}
```

- [custom_hash_algorithms](resources--workload--reference--group-019.md#canonical-0321322131200210-0113302313033133-1020201133022012-0110122211211203-0033010232003213-2113013333102202-0333322013220101-1133131233131021): complete subsection reference.

<a id="canonical-0212320032222011-2013032302300231-2202033111221110-1202301023300030-0013200020201022-2031003211330001-0223022110032213-0112011223123113"></a>

<a id="canonical-1311331133032212-1330301333210011-0222322132112030-2001020110323130-3020133011311221-1131203221310210-0022111210312212-2021030033301131"></a>

#### `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates.description_spec` property

Type: `"string"`. Optional.

Description. Description for the certificate.

- [disable_ocsp_stapling](resources--workload--reference--group-019.md#canonical-3301303000233223-1331301220303233-2130110233111130-1022311112112031-0010323212302010-2320331230300113-1233002033103101-1021212200303002): complete subsection reference.

- [private_key](resources--workload--reference--group-019.md#canonical-3202322131133212-3233333112022031-2012323022120230-3213102233121101-2023030003200212-0221230103200223-2321231323033110-1310133003012200): complete subsection reference.

- [use_system_defaults](resources--workload--reference--group-019.md#canonical-0322223320021121-2322111312221300-0121032323323102-3201122311322132-2130230200033330-3132003303112111-0313001301113023-1022311120311313): complete subsection reference.

<a id="canonical-0311333301123130-3101332300002121-3010331213132202-0223222131101322-1123330112003000-2133011022302032-1130000212313030-1321200033013210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates.blindfold` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-018.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-018.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_custom](resources--workload--reference--group-018.md#canonical-0203102001320123-1120130012130212-0111101220231333-3101331123203121-1110000311003232-1321231201202312-3200200021132112-0200303231330002)
- [stateful_service.advertise_options.advertise_custom.ports](resources--workload--reference--group-018.md#canonical-3111311322332303-1201030020021201-3133021233000313-2022333331131202-3221312032323033-3133120120113232-3330021120310331-2331303302032231)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-018.md#canonical-0313222022003221-2301103230203023-3120322210033100-2121220010332201-3133231021021220-1313001311100021-0313102021211212-0331322021202002)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-019.md#canonical-0112020200111022-0013103013103331-1110301001233101-1300202101312213-3301031210112102-1221232301300211-0231000232130131-3132113303133333)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-019.md#canonical-2020320003013033-2231100311132300-2132132133022132-3002120002132203-3131231201330331-0231012122032023-1202211010110230-0230013122122221)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates](resources--workload--reference--group-019.md#canonical-3332313302002330-1333101212032003-2203023132232213-0001033110210120-3121103210011323-2020220120012213-3220330021330022-1301300222033110)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates.blindfold

<a id="canonical-3333322221003211-1031233130000103-2113010232211202-2220330310013112-1201001203030300-3300332233301130-0323231103330212-1033022223113300"></a>

Type: `"single"`. Optional.

Native certificate preparation. Use PEM files or a P12 file, or write-only key/bundle values with
material\_version (Terraform 1.11+). Defaults to shared/ves-io-allow-volterra. Inline certificates
require unique IDs. Private inputs are never stored.

<a id="canonical-3031310330131111-1120003332001323-0121011110203323-3103113021212100-0311132021021132-2200210113133323-0223100303102113-3011331133313231"></a>

### Direct properties for `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates.blindfold`

<a id="canonical-2210012232210111-3112032323113212-2111201121122033-2033130013122303-0312130002022302-0033322303323011-3310131232023022-3211321122002112"></a>

#### `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates.blindfold.algorithm` property

Type: `"string"`. Computed.

<a id="canonical-0311233322213101-0233102120222020-1203010031022010-3220102320010231-2312200110000033-3200212101210232-1300132310313110-2322203322330323"></a>

<a id="canonical-2133000310202312-2021022130033100-0000331102320221-0031010222233021-2231303121203312-2123230011011303-1233331322111123-1330020002120133"></a>

#### `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates.blindfold.certificate_file` property

Type: `"string"`. Optional.

<a id="canonical-2112213330212122-3331020030322102-1102123031312002-0001011200013313-2220113313220322-0220111022212200-3100220110322113-2213021333220322"></a>

<a id="canonical-3111222210001102-3320011212010231-2222133021013223-2033110311222333-0312122220322022-1010202130001102-2310133122201123-3103132220033203"></a>

#### `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates.blindfold.certificate_pem` property

Type: `"string"`. Optional.

<a id="canonical-0302031022312131-2032133011332001-1222113022210030-3330210113121133-0103011022221330-1000030032312321-2100232202333001-2222322012321300"></a>

<a id="canonical-0101010030010120-3312122112132022-3312001102031231-2022230022123200-2333223002130023-0032102112233011-2130230102030333-3012013312002232"></a>

#### `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates.blindfold.chain_identity` property

Type: `"string"`. Computed.

<a id="canonical-0323301330320323-0113310313203331-0310322002233210-2210230033100101-1133303331330011-1110333013213211-0300011102000022-0022212020121020"></a>

<a id="canonical-3003012011121310-1323332320011301-0111313310212003-2233331133112123-0020130010213003-1113301332233201-0002111110211332-3013133130122000"></a>

#### `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates.blindfold.context_digest` property

Type: `"string"`. Computed.

<a id="canonical-2000321310322201-0310130132120103-0201301320312112-1120100201201020-3111011121212021-0102003112032033-2001000132023310-1322033021030323"></a>

<a id="canonical-3133211220010032-3312331001232120-3020220221310000-0023322302112133-0022112320302133-2330232221221020-1013101331320311-1202010313300333"></a>

#### `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates.blindfold.encrypted_location` property

Type: `"string"`. Computed, Sensitive.

<a id="canonical-0010310022032123-3333200311110330-2302111101323311-1020212120113010-3110030333302333-2303033232232133-1303031301000111-0310230203321222"></a>

<a id="canonical-3011100321020230-1110033032210100-0330133322100333-1113132232313100-1102001001313023-3210032201320103-1220011201112221-1232010212311130"></a>

#### `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates.blindfold.expires_at` property

Type: `"string"`. Computed.

<a id="canonical-1211121011310232-1013320113321022-0200031033223321-0320133003122120-0230302033113221-0123101022302012-0032102323303320-3201020300220203"></a>

<a id="canonical-1101223201102022-3133123122030133-0312100131300301-0301012101133123-3032233002333201-3131011310030222-3222222212221103-3202031313221212"></a>

#### `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates.blindfold.fingerprint` property

Type: `"string"`. Computed.

<a id="canonical-2200231320200111-1001222301330311-1233100111120113-3113133131221200-3120203033211312-0010133301331322-1013033123302211-0201223313030031"></a>

<a id="canonical-2210202302101030-0010133130022323-2102320103002001-2022100101130220-3121030321201312-1033102030000122-2220221333201013-2012230320122100"></a>

#### `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates.blindfold.id` property

Type: `"string"`. Optional.

<a id="canonical-0100331212011213-1112302312232100-1331013223201320-0301112201003233-1330012021112311-2121111101200212-0301332123330102-2013131102310313"></a>

<a id="canonical-2100210322030111-0012222231332221-3201130102121112-0122133320101312-2102211221331111-2120000203230301-2310131021213101-1010201320213002"></a>

#### `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates.blindfold.material_version` property

Type: `"string"`. Optional.

<a id="canonical-3213321010220200-3222131212120103-0330110000231320-3123011231200330-2311223131322121-0302213230033130-0313212320201230-3130333202213321"></a>

<a id="canonical-1302003033113332-3000011303011010-3202323221300131-1033321222021123-1221123301313112-1103030203131200-3113001303123302-3302202321320021"></a>

#### `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates.blindfold.passphrase_env` property

Type: `"string"`. Optional.

<a id="canonical-0303100332102313-3331310031300012-2110313000131020-3020013111211012-1130012200311122-3121222111100120-2023312103030231-0003313132120333"></a>

<a id="canonical-0011302010203122-1301221131210120-3022113223201230-1232111011101212-2012021233202330-0213313313302233-2012100202320032-0100322121202023"></a>

#### `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates.blindfold.passphrase_wo` property

Type: `"string"`. Optional, Sensitive, Write_only.

<a id="canonical-3222010102002003-3331221030213130-1233333232030112-2120033230023101-0201320202231201-1110100313231110-0000000233312332-2222331120112221"></a>

<a id="canonical-2311030000103321-0323001020302232-2111300031130301-0112232032311211-3110200022033032-2330200322132110-3020100233321203-1203023302322203"></a>

#### `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates.blindfold.pkcs12_file` property

Type: `"string"`. Optional.

<a id="canonical-2123110321211202-1321302211302033-3132330322231001-0301220322220233-1202201112321230-3233311303103223-1212020312032013-3201202000202221"></a>

<a id="canonical-0033110102333313-3203330330001011-3030022313232300-3033230000020233-2123221332323003-3033022333203113-1232332330302122-2000202233022002"></a>

#### `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates.blindfold.pkcs12_wo` property

Type: `"string"`. Optional, Sensitive, Write_only.

<a id="canonical-3322330323313330-1101233011312022-3302210010222222-0011200202323222-3323031023023201-1002022122230322-0322110012111000-1003201112213122"></a>

<a id="canonical-3101333131032332-1122302202131110-2223213130033300-3130002300303211-3112122121322302-3101210101213012-0033113211033332-1200102102211232"></a>

#### `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates.blindfold.policy` property

Type: `"string"`. Optional.

<a id="canonical-3123310223323212-0303320023321022-1032012313111021-0102201302310102-3332020123110001-1301230030131300-2203101233310200-2000031021300323"></a>

<a id="canonical-3311112013132133-1320320303300302-1211231023011212-1032232012001002-2130102201112110-3032333121232332-0312120131203321-0220130223001221"></a>

#### `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates.blindfold.prepared_identity` property

Type: `"string"`. Computed.

<a id="canonical-0210101202022022-3211033202023101-1233331212110203-2033032020113210-3210221131000033-0132121130322210-1132232330301120-2210102211131233"></a>

<a id="canonical-0012332203132213-1031202302112300-1330200130311130-3133122312102023-1221121201322120-3302213330110330-0331000321331133-2330101313112111"></a>

#### `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates.blindfold.private_key_file` property

Type: `"string"`. Optional.

<a id="canonical-3021010302321233-0330003322100120-2223103013301222-2211122133200123-2032321001312103-0321011333003221-0120210023013333-3001013201110112"></a>

<a id="canonical-2302022010101010-2011211312133010-2110001033210301-1333202133013302-2023332312020303-1321312100210312-0301211313133030-1321301210021100"></a>

#### `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates.blindfold.private_key_wo` property

Type: `"string"`. Optional, Sensitive, Write_only.

<a id="canonical-2011203003331332-2003222331311322-1323023332133313-3333202332130232-1200203013021202-1310301113032310-0301322010210110-2100213101220303"></a>

<a id="canonical-1021221111213301-1123023321210133-2230312320003110-0110323100310002-3311111211320113-0030001030332002-3332023210330220-2101131313032231"></a>

#### `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates.blindfold.spki_identity` property

Type: `"string"`. Computed.

<a id="canonical-0321322131200210-0113302313033133-1020201133022012-0110122211211203-0033010232003213-2113013333102202-0333322013220101-1133131233131021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates.custom_hash_algorithms` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-018.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-018.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_custom](resources--workload--reference--group-018.md#canonical-0203102001320123-1120130012130212-0111101220231333-3101331123203121-1110000311003232-1321231201202312-3200200021132112-0200303231330002)
- [stateful_service.advertise_options.advertise_custom.ports](resources--workload--reference--group-018.md#canonical-3111311322332303-1201030020021201-3133021233000313-2022333331131202-3221312032323033-3133120120113232-3330021120310331-2331303302032231)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-018.md#canonical-0313222022003221-2301103230203023-3120322210033100-2121220010332201-3133231021021220-1313001311100021-0313102021211212-0331322021202002)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-019.md#canonical-0112020200111022-0013103013103331-1110301001233101-1300202101312213-3301031210112102-1221232301300211-0231000232130131-3132113303133333)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-019.md#canonical-2020320003013033-2231100311132300-2132132133022132-3002120002132203-3131231201330331-0231012122032023-1202211010110230-0230013122122221)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates](resources--workload--reference--group-019.md#canonical-3332313302002330-1333101212032003-2203023132232213-0001033110210120-3121103210011323-2020220120012213-3220330021330022-1301300222033110)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates.custom_hash_algorithms

<a id="canonical-1001302023300321-0001301313320231-1013111022332231-1223020300102221-3222122300212213-2001313310231022-1001200201033033-0030020033200201"></a>

Type: `"object"`. single nested block, Optional.

Specifies the hash algorithms to be used.

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
custom_hash_algorithms {
  # Configure direct properties listed below.
}
```

<a id="canonical-3300203310020332-2030011102212210-3001001100100032-1110002323130123-0100222312320221-2002233132030332-2203330233001023-3321230223322030"></a>

### Direct properties for `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates.custom_hash_algorithms`

<a id="canonical-1022223030202213-0201330303113123-2321020103113223-3230321123203323-3200222320023312-0311031030213222-0122002102302133-2230013312102101"></a>

#### `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates.custom_hash_algorithms.hash_algorithms` property

Type: `["list", "string"]`. Optional.

\[Enum: INVALID\_HASH\_ALGORITHM|SHA256|SHA1\] Ordered list of hash algorithms to be used. Possible
values are \`INVALID\_HASH\_ALGORITHM\`, \`SHA256\`, \`SHA1\`. Defaults to
\`INVALID\_HASH\_ALGORITHM\`.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 4,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 4,
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
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3301303000233223-1331301220303233-2130110233111130-1022311112112031-0010323212302010-2320331230300113-1233002033103101-1021212200303002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates.disable_ocsp_stapling` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-018.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-018.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_custom](resources--workload--reference--group-018.md#canonical-0203102001320123-1120130012130212-0111101220231333-3101331123203121-1110000311003232-1321231201202312-3200200021132112-0200303231330002)
- [stateful_service.advertise_options.advertise_custom.ports](resources--workload--reference--group-018.md#canonical-3111311322332303-1201030020021201-3133021233000313-2022333331131202-3221312032323033-3133120120113232-3330021120310331-2331303302032231)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-018.md#canonical-0313222022003221-2301103230203023-3120322210033100-2121220010332201-3133231021021220-1313001311100021-0313102021211212-0331322021202002)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-019.md#canonical-0112020200111022-0013103013103331-1110301001233101-1300202101312213-3301031210112102-1221232301300211-0231000232130131-3132113303133333)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-019.md#canonical-2020320003013033-2231100311132300-2132132133022132-3002120002132203-3131231201330331-0231012122032023-1202211010110230-0230013122122221)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates](resources--workload--reference--group-019.md#canonical-3332313302002330-1333101212032003-2203023132232213-0001033110210120-3121103210011323-2020220120012213-3220330021330022-1301300222033110)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates.disable_ocsp_stapling

<a id="canonical-0222113132112232-0122232130220212-2021220023003301-1210312110202323-3320130112132213-1101012122300211-1033120303032213-2031230212323203"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disable ocsp stapling.

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
disable_ocsp_stapling = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3202322131133212-3233333112022031-2012323022120230-3213102233121101-2023030003200212-0221230103200223-2321231323033110-1310133003012200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates.private_key` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-018.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-018.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_custom](resources--workload--reference--group-018.md#canonical-0203102001320123-1120130012130212-0111101220231333-3101331123203121-1110000311003232-1321231201202312-3200200021132112-0200303231330002)
- [stateful_service.advertise_options.advertise_custom.ports](resources--workload--reference--group-018.md#canonical-3111311322332303-1201030020021201-3133021233000313-2022333331131202-3221312032323033-3133120120113232-3330021120310331-2331303302032231)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-018.md#canonical-0313222022003221-2301103230203023-3120322210033100-2121220010332201-3133231021021220-1313001311100021-0313102021211212-0331322021202002)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-019.md#canonical-0112020200111022-0013103013103331-1110301001233101-1300202101312213-3301031210112102-1221232301300211-0231000232130131-3132113303133333)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-019.md#canonical-2020320003013033-2231100311132300-2132132133022132-3002120002132203-3131231201330331-0231012122032023-1202211010110230-0230013122122221)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates](resources--workload--reference--group-019.md#canonical-3332313302002330-1333101212032003-2203023132232213-0001033110210120-3121103210011323-2020220120012213-3220330021330022-1301300222033110)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates.private_key

<a id="canonical-1021102231312001-0330333033133100-1010222301102031-2011201131310331-3202312221013120-2131200330023123-1301313002120102-1030213103000222"></a>

Type: `"object"`. single nested block, Optional.

SecretType is used in an object to indicate a sensitive/confidential field.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-secret_info_oneof": "[\"blindfold_secret_info\",\"clear_secret_info\"]"
}
```

Terraform syntax:

```terraform
private_key {
  # Configure direct properties listed below.
}
```

<a id="canonical-0323200010222110-0012031002202303-0031331233013211-3102013303111231-0221301031002102-2233230010012210-1302300010200211-0321123210003133"></a>

### Direct properties for `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates.private_key`

- [blindfold_secret_info](resources--workload--reference--group-019.md#canonical-2103032132323102-3023011022021303-2201010033131212-2030111103320103-1120033321203121-3003213000323203-0323313000031023-2101102112002303): complete subsection reference.

- [clear_secret_info](resources--workload--reference--group-019.md#canonical-2331030102322003-1213200013022113-0001131012321023-1132102120312221-3220112212333213-0022231333213120-3102111000212031-3101013002021023): complete subsection reference.

<a id="canonical-2103032132323102-3023011022021303-2201010033131212-2030111103320103-1120033321203121-3003213000323203-0323313000031023-2101102112002303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates.private_key.blindfold_secret_info` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-018.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-018.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_custom](resources--workload--reference--group-018.md#canonical-0203102001320123-1120130012130212-0111101220231333-3101331123203121-1110000311003232-1321231201202312-3200200021132112-0200303231330002)
- [stateful_service.advertise_options.advertise_custom.ports](resources--workload--reference--group-018.md#canonical-3111311322332303-1201030020021201-3133021233000313-2022333331131202-3221312032323033-3133120120113232-3330021120310331-2331303302032231)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-018.md#canonical-0313222022003221-2301103230203023-3120322210033100-2121220010332201-3133231021021220-1313001311100021-0313102021211212-0331322021202002)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-019.md#canonical-0112020200111022-0013103013103331-1110301001233101-1300202101312213-3301031210112102-1221232301300211-0231000232130131-3132113303133333)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-019.md#canonical-2020320003013033-2231100311132300-2132132133022132-3002120002132203-3131231201330331-0231012122032023-1202211010110230-0230013122122221)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates](resources--workload--reference--group-019.md#canonical-3332313302002330-1333101212032003-2203023132232213-0001033110210120-3121103210011323-2020220120012213-3220330021330022-1301300222033110)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates.private_key](resources--workload--reference--group-019.md#canonical-3202322131133212-3233333112022031-2012323022120230-3213102233121101-2023030003200212-0221230103200223-2321231323033110-1310133003012200)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates.private_key.blindfold_secret_info

<a id="canonical-2223211110110010-3002020003233223-3022122101231222-0222203203103331-2031220131320011-2333131123323011-1103221300333131-0120121130332331"></a>

Type: `"object"`. single nested block, Optional.

BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management.

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
blindfold_secret_info {
  # Configure direct properties listed below.
}
```

<a id="canonical-0322321100211100-3302111131333022-3033223123000010-0331310023320111-2102313332011202-1221202020232022-3330232121331213-1102101223000030"></a>

### Direct properties for `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates.private_key.blindfold_secret_info`

<a id="canonical-0222201130033332-3000103022310312-3031203010230210-0330121313202330-3331332003020311-0130002221132021-2130312110020300-1233231112221213"></a>

#### `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates.private_key.blindfold_secret_info.decryption_provider` property

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the backend Secret
Management service.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-10-09T12:34:59+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1333011002321003-2130320133121002-0222221213323210-1021301222323220-3212020121320232-3030311220103120-2121212121232100-1333010232231002"></a>

<a id="canonical-0002201223302300-3030021201232333-3010333301320203-0021312010023222-2120332330001312-1331020332210103-2030330300120100-0032132313102022"></a>

#### `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates.private_key.blindfold_secret_info.location` property

Type: `"string"`. Optional, Sensitive.

Location is the URI\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "content",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "maxLength": 131072,
    "metadata": {
      "category": "content",
      "confidence": 1.0,
      "note": "Blindfold envelope encryption (AES-256-GCM + RSA-OAEP) of an RSA-2048 TLS private key produces ~3700 char string:/// URL. 128KB max secret size = ~175KB base64. Discovery reported 1024 which is incorrect.",
      "source": "manual-override",
      "validatedAt": "2026-10-09T12:34:59+00:00"
    },
    "minLength": 4
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-f5xc-sensitive": true,
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

<a id="canonical-1013002332022322-0023210323011333-3003123222023230-0011023301013023-0111230303223131-1332220303222300-1010313000002122-3031100123100013"></a>

<a id="canonical-3212011323102303-2023321110010132-2123331122033202-2233212320031013-2003101330112103-1000021230032221-2121101013323103-1001103003212200"></a>

#### `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates.private_key.blindfold_secret_info.store_provider` property

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-10-09T12:34:59+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2331030102322003-1213200013022113-0001131012321023-1132102120312221-3220112212333213-0022231333213120-3102111000212031-3101013002021023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates.private_key.clear_secret_info` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-018.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-018.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_custom](resources--workload--reference--group-018.md#canonical-0203102001320123-1120130012130212-0111101220231333-3101331123203121-1110000311003232-1321231201202312-3200200021132112-0200303231330002)
- [stateful_service.advertise_options.advertise_custom.ports](resources--workload--reference--group-018.md#canonical-3111311322332303-1201030020021201-3133021233000313-2022333331131202-3221312032323033-3133120120113232-3330021120310331-2331303302032231)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-018.md#canonical-0313222022003221-2301103230203023-3120322210033100-2121220010332201-3133231021021220-1313001311100021-0313102021211212-0331322021202002)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-019.md#canonical-0112020200111022-0013103013103331-1110301001233101-1300202101312213-3301031210112102-1221232301300211-0231000232130131-3132113303133333)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-019.md#canonical-2020320003013033-2231100311132300-2132132133022132-3002120002132203-3131231201330331-0231012122032023-1202211010110230-0230013122122221)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates](resources--workload--reference--group-019.md#canonical-3332313302002330-1333101212032003-2203023132232213-0001033110210120-3121103210011323-2020220120012213-3220330021330022-1301300222033110)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates.private_key](resources--workload--reference--group-019.md#canonical-3202322131133212-3233333112022031-2012323022120230-3213102233121101-2023030003200212-0221230103200223-2321231323033110-1310133003012200)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates.private_key.clear_secret_info

<a id="canonical-3211312332132010-1230022330210303-3132011203230101-0122001123300233-1323203102301023-2023023112001021-3200002131133313-3223303121120301"></a>

Type: `"object"`. single nested block, Optional.

ClearSecretInfoType specifies information about the Secret that is not encrypted.

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
clear_secret_info {
  # Configure direct properties listed below.
}
```

<a id="canonical-0200213200302330-0111101101002133-0113332332323301-2321333213112013-0331111301301323-2101332112322122-2111101213112013-2222112310313320"></a>

### Direct properties for `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates.private_key.clear_secret_info`

<a id="canonical-3132101211300201-1212122333202112-0010032312013230-1000132000231222-1313201221321130-1021000133021213-3010103032011103-3121122212101022"></a>

#### `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates.private_key.clear_secret_info.provider_ref` property

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-0103212110112122-2323210303112123-1111033301312133-1023222133302101-0113211321003112-2033332001030012-1112010300000300-0012322220231130"></a>

<a id="canonical-2001011023000210-0112131012122103-3330212012133201-2230311322002033-0320103101033220-3230323312210011-2012002222123311-0113020303123302"></a>

#### `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates.private_key.clear_secret_info.url` property

Type: `"string"`. Optional, Sensitive.

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded base64 format. When asked for this secret, caller will GET Secret bytes after
base64 decoding.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 131072,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 131072
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "formatDescription": "RFC 3986 URI with scheme (http, https, ftp)",
    "maxLength": 131072,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-09T12:34:59+00:00"
    },
    "minLength": 1,
    "pattern": "^(https?|ftp)://[^\\s/$.?#].[^\\s]*$",
    "validation": {
      "rfc": "RFC 3986"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-f5xc-sensitive": true,
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

<a id="canonical-0322223320021121-2322111312221300-0121032323323102-3201122311322132-2130230200033330-3132003303112111-0313001301113023-1022311120311313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates.use_system_defaults` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-018.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-018.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_custom](resources--workload--reference--group-018.md#canonical-0203102001320123-1120130012130212-0111101220231333-3101331123203121-1110000311003232-1321231201202312-3200200021132112-0200303231330002)
- [stateful_service.advertise_options.advertise_custom.ports](resources--workload--reference--group-018.md#canonical-3111311322332303-1201030020021201-3133021233000313-2022333331131202-3221312032323033-3133120120113232-3330021120310331-2331303302032231)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-018.md#canonical-0313222022003221-2301103230203023-3120322210033100-2121220010332201-3133231021021220-1313001311100021-0313102021211212-0331322021202002)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-019.md#canonical-0112020200111022-0013103013103331-1110301001233101-1300202101312213-3301031210112102-1221232301300211-0231000232130131-3132113303133333)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-019.md#canonical-2020320003013033-2231100311132300-2132132133022132-3002120002132203-3131231201330331-0231012122032023-1202211010110230-0230013122122221)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates](resources--workload--reference--group-019.md#canonical-3332313302002330-1333101212032003-2203023132232213-0001033110210120-3121103210011323-2020220120012213-3220330021330022-1301300222033110)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates.use_system_defaults

<a id="canonical-1233110211312100-1323000022012010-1133302103130023-3200120131010032-1020200100012130-0300112110201103-3322333332031330-0313111113132123"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for use system defaults.

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
use_system_defaults = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3212132103211102-3103133011133021-3313221110320101-2111103132233113-1232201321022230-0232111013133132-1231122000001132-3012311030121223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_config` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-018.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-018.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_custom](resources--workload--reference--group-018.md#canonical-0203102001320123-1120130012130212-0111101220231333-3101331123203121-1110000311003232-1321231201202312-3200200021132112-0200303231330002)
- [stateful_service.advertise_options.advertise_custom.ports](resources--workload--reference--group-018.md#canonical-3111311322332303-1201030020021201-3133021233000313-2022333331131202-3221312032323033-3133120120113232-3330021120310331-2331303302032231)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-018.md#canonical-0313222022003221-2301103230203023-3120322210033100-2121220010332201-3133231021021220-1313001311100021-0313102021211212-0331322021202002)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-019.md#canonical-0112020200111022-0013103013103331-1110301001233101-1300202101312213-3301031210112102-1221232301300211-0231000232130131-3132113303133333)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-019.md#canonical-2020320003013033-2231100311132300-2132132133022132-3002120002132203-3131231201330331-0231012122032023-1202211010110230-0230013122122221)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_config

<a id="canonical-1302313322112233-2103110020213110-3111100001201221-1102212003220303-3110220222030231-0232211313200331-3322003300321221-1133012303222032"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
tls_config {
  # Configure direct properties listed below.
}
```
