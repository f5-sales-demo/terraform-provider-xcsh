---
page_title: "xcsh_workload reference"
subcategory: "Container"
description: "Complete grouped canonical reference for xcsh_workload reference."
---

# xcsh_workload reference

<a id="canonical-0112020200111022-0013103013103331-1110301001233101-1300202101312213-3301031210112102-1221232301300211-0231000232130131-3132113303133333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-017.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_custom](resources--workload--reference--group-017.md#canonical-0203102001320123-1120130012130212-0111101220231333-3101331123203121-1110000311003232-1321231201202312-3200200021132112-0200303231330002)
- [stateful_service.advertise_options.advertise_custom.ports](resources--workload--reference--group-017.md#canonical-3111311322332303-1201030020021201-3133021233000313-2022333331131202-3221312032323033-3133120120113232-3330021120310331-2331303302032231)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-017.md#canonical-0313222022003221-2301103230203023-3120322210033100-2121220010332201-3133231021021220-1313001311100021-0313102021211212-0331322021202002)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https

<a id="canonical-0120310310023211-3131320332233110-0100310123030011-2330301013323021-2112320302331231-0001123023313213-2103223200103222-2202012223022002"></a>

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

- [coalescing_options](resources--workload--reference--group-018.md#canonical-2032232023103112-3132002323011100-2333333333213030-0131331001332133-3231230321300001-3121020113202300-0031031222210022-3100303312203312): complete subsection reference.

<a id="canonical-1301231032211202-1002223130012320-0122120200332003-3301133012231320-1202122201233313-2332303300230113-0222123310103033-3123211210102003"></a>

<a id="canonical-2110123330232132-3220330120100133-2121033310331101-2220131333100032-3210220020133332-1302200102301331-0333021012222022-2130301132210112"></a>

#### `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.connection_idle_timeout` property

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

- [default_header](resources--workload--reference--group-018.md#canonical-0003223323303031-3311332233132213-3000212121311110-0322333000010000-2012102033100303-2333200110221100-3023101230212212-0321110100231233): complete subsection reference.

- [default_loadbalancer](resources--workload--reference--group-018.md#canonical-2301030210312013-3330212130031303-3001333222011323-0133033200232003-2033031011231022-1331312210302033-3202203112021311-1123211231122020): complete subsection reference.

- [disable_path_normalize](resources--workload--reference--group-018.md#canonical-0123111002221221-3000302331312022-3020301301022002-2130333003203132-0322310312020311-1033223110102102-1222203300201132-0210022222201102): complete subsection reference.

- [enable_path_normalize](resources--workload--reference--group-018.md#canonical-2130133323100220-2131101221132301-1100311112031132-2020301212111213-0130001222222232-1121312101101020-1322130333030133-2322303321010132): complete subsection reference.

- [http_protocol_options](resources--workload--reference--group-018.md#canonical-3202212001033033-3133231302102230-0333203123220020-0022212313320130-3032131111301011-1301121203330231-2010200210001321-0233311332222011): complete subsection reference.

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

- [non_default_loadbalancer](resources--workload--reference--group-018.md#canonical-1103323231113001-2101233200313132-0113203102011131-2310133213023120-2300020303303112-1311213222021323-3201102120300132-1030302220012022): complete subsection reference.

- [pass_through](resources--workload--reference--group-018.md#canonical-3122320112122220-2222123003111322-1032303200202320-0301311021210013-0010110320313023-3321020211021302-2302303311133010-3231010220310031): complete subsection reference.

<a id="canonical-3021222313010203-3012031133311010-2131100123333222-2230203033211331-0303021120033303-2133122111330231-0130130321021123-0232202320231120"></a>

<a id="canonical-1221201210020333-0003230331311323-0322100231303101-0012303311320032-2010233010223301-2300203302312113-1010012031300312-2320300113123232"></a>

#### `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.port` property

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

<a id="canonical-1131101113301310-2303020113103231-2011020111322200-1330213212013100-3130100212000123-3120233111023031-2103200012201321-2003210323023111"></a>

<a id="canonical-1203022202121021-0213001202330223-3212220120301220-3131323030203121-2031020120310021-1120233121301231-2030200132123203-2203030133020231"></a>

#### `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.port_ranges` property

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

<a id="canonical-0033303101323213-3002212331203102-2200312221220002-3030312323220031-1011222110203111-0100002223222121-1223112020222211-0030231021011210"></a>

<a id="canonical-1001011313010210-1223121123022323-2322133220313130-2121032123111212-1313223012333202-0001323120033103-1002102003123101-2311333220121330"></a>

#### `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.server_name` property

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

- [tls_cert_params](resources--workload--reference--group-018.md#canonical-1330323120222300-2123201010331320-0300232300332322-0010223222211002-1032302312300112-0233320002220301-1020331113013010-3202200223121221): complete subsection reference.

- [tls_parameters](resources--workload--reference--group-018.md#canonical-2020320003013033-2231100311132300-2132132133022132-3002120002132203-3131231201330331-0231012122032023-1202211010110230-0230013122122221): complete subsection reference.

<a id="canonical-2032232023103112-3132002323011100-2333333333213030-0131331001332133-3231230321300001-3121020113202300-0031031222210022-3100303312203312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.coalescing_options` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-017.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_custom](resources--workload--reference--group-017.md#canonical-0203102001320123-1120130012130212-0111101220231333-3101331123203121-1110000311003232-1321231201202312-3200200021132112-0200303231330002)
- [stateful_service.advertise_options.advertise_custom.ports](resources--workload--reference--group-017.md#canonical-3111311322332303-1201030020021201-3133021233000313-2022333331131202-3221312032323033-3133120120113232-3330021120310331-2331303302032231)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-017.md#canonical-0313222022003221-2301103230203023-3120322210033100-2121220010332201-3133231021021220-1313001311100021-0313102021211212-0331322021202002)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-018.md#canonical-0112020200111022-0013103013103331-1110301001233101-1300202101312213-3301031210112102-1221232301300211-0231000232130131-3132113303133333)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.coalescing_options

<a id="canonical-1301323033020202-2133133302210233-0023231001201311-3133012110211023-3202233120200201-2233031330231232-0101113031210003-0301211201220002"></a>

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

<a id="canonical-3333013131220330-3302021121010323-2313231211133313-0103212200102031-2231211030310011-3011302222023203-3002310123102313-2332321103213301"></a>

### Direct properties for `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.coalescing_options`

- [default_coalescing](resources--workload--reference--group-018.md#canonical-0013002130222122-0232011323323313-0103322131310223-0300131033232133-0103223223322302-2013100033113013-3320203231321222-3310130113020200): complete subsection reference.

- [strict_coalescing](resources--workload--reference--group-018.md#canonical-2010111233322110-2233220133212231-3201112221030333-3110010013300201-3331022031103023-3021212212121321-2113101133100103-1110021210203001): complete subsection reference.

<a id="canonical-0013002130222122-0232011323323313-0103322131310223-0300131033232133-0103223223322302-2013100033113013-3320203231321222-3310130113020200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.coalescing_options.default_coalescing` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-017.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_custom](resources--workload--reference--group-017.md#canonical-0203102001320123-1120130012130212-0111101220231333-3101331123203121-1110000311003232-1321231201202312-3200200021132112-0200303231330002)
- [stateful_service.advertise_options.advertise_custom.ports](resources--workload--reference--group-017.md#canonical-3111311322332303-1201030020021201-3133021233000313-2022333331131202-3221312032323033-3133120120113232-3330021120310331-2331303302032231)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-017.md#canonical-0313222022003221-2301103230203023-3120322210033100-2121220010332201-3133231021021220-1313001311100021-0313102021211212-0331322021202002)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-018.md#canonical-0112020200111022-0013103013103331-1110301001233101-1300202101312213-3301031210112102-1221232301300211-0231000232130131-3132113303133333)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.coalescing_options](resources--workload--reference--group-018.md#canonical-2032232023103112-3132002323011100-2333333333213030-0131331001332133-3231230321300001-3121020113202300-0031031222210022-3100303312203312)
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
- [stateful_service](resources--workload--reference--group-017.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_custom](resources--workload--reference--group-017.md#canonical-0203102001320123-1120130012130212-0111101220231333-3101331123203121-1110000311003232-1321231201202312-3200200021132112-0200303231330002)
- [stateful_service.advertise_options.advertise_custom.ports](resources--workload--reference--group-017.md#canonical-3111311322332303-1201030020021201-3133021233000313-2022333331131202-3221312032323033-3133120120113232-3330021120310331-2331303302032231)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-017.md#canonical-0313222022003221-2301103230203023-3120322210033100-2121220010332201-3133231021021220-1313001311100021-0313102021211212-0331322021202002)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-018.md#canonical-0112020200111022-0013103013103331-1110301001233101-1300202101312213-3301031210112102-1221232301300211-0231000232130131-3132113303133333)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.coalescing_options](resources--workload--reference--group-018.md#canonical-2032232023103112-3132002323011100-2333333333213030-0131331001332133-3231230321300001-3121020113202300-0031031222210022-3100303312203312)
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
- [stateful_service](resources--workload--reference--group-017.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_custom](resources--workload--reference--group-017.md#canonical-0203102001320123-1120130012130212-0111101220231333-3101331123203121-1110000311003232-1321231201202312-3200200021132112-0200303231330002)
- [stateful_service.advertise_options.advertise_custom.ports](resources--workload--reference--group-017.md#canonical-3111311322332303-1201030020021201-3133021233000313-2022333331131202-3221312032323033-3133120120113232-3330021120310331-2331303302032231)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-017.md#canonical-0313222022003221-2301103230203023-3120322210033100-2121220010332201-3133231021021220-1313001311100021-0313102021211212-0331322021202002)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-018.md#canonical-0112020200111022-0013103013103331-1110301001233101-1300202101312213-3301031210112102-1221232301300211-0231000232130131-3132113303133333)
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
- [stateful_service](resources--workload--reference--group-017.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_custom](resources--workload--reference--group-017.md#canonical-0203102001320123-1120130012130212-0111101220231333-3101331123203121-1110000311003232-1321231201202312-3200200021132112-0200303231330002)
- [stateful_service.advertise_options.advertise_custom.ports](resources--workload--reference--group-017.md#canonical-3111311322332303-1201030020021201-3133021233000313-2022333331131202-3221312032323033-3133120120113232-3330021120310331-2331303302032231)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-017.md#canonical-0313222022003221-2301103230203023-3120322210033100-2121220010332201-3133231021021220-1313001311100021-0313102021211212-0331322021202002)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-018.md#canonical-0112020200111022-0013103013103331-1110301001233101-1300202101312213-3301031210112102-1221232301300211-0231000232130131-3132113303133333)
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
- [stateful_service](resources--workload--reference--group-017.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_custom](resources--workload--reference--group-017.md#canonical-0203102001320123-1120130012130212-0111101220231333-3101331123203121-1110000311003232-1321231201202312-3200200021132112-0200303231330002)
- [stateful_service.advertise_options.advertise_custom.ports](resources--workload--reference--group-017.md#canonical-3111311322332303-1201030020021201-3133021233000313-2022333331131202-3221312032323033-3133120120113232-3330021120310331-2331303302032231)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-017.md#canonical-0313222022003221-2301103230203023-3120322210033100-2121220010332201-3133231021021220-1313001311100021-0313102021211212-0331322021202002)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-018.md#canonical-0112020200111022-0013103013103331-1110301001233101-1300202101312213-3301031210112102-1221232301300211-0231000232130131-3132113303133333)
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
- [stateful_service](resources--workload--reference--group-017.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_custom](resources--workload--reference--group-017.md#canonical-0203102001320123-1120130012130212-0111101220231333-3101331123203121-1110000311003232-1321231201202312-3200200021132112-0200303231330002)
- [stateful_service.advertise_options.advertise_custom.ports](resources--workload--reference--group-017.md#canonical-3111311322332303-1201030020021201-3133021233000313-2022333331131202-3221312032323033-3133120120113232-3330021120310331-2331303302032231)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-017.md#canonical-0313222022003221-2301103230203023-3120322210033100-2121220010332201-3133231021021220-1313001311100021-0313102021211212-0331322021202002)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-018.md#canonical-0112020200111022-0013103013103331-1110301001233101-1300202101312213-3301031210112102-1221232301300211-0231000232130131-3132113303133333)
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
- [stateful_service](resources--workload--reference--group-017.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_custom](resources--workload--reference--group-017.md#canonical-0203102001320123-1120130012130212-0111101220231333-3101331123203121-1110000311003232-1321231201202312-3200200021132112-0200303231330002)
- [stateful_service.advertise_options.advertise_custom.ports](resources--workload--reference--group-017.md#canonical-3111311322332303-1201030020021201-3133021233000313-2022333331131202-3221312032323033-3133120120113232-3330021120310331-2331303302032231)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-017.md#canonical-0313222022003221-2301103230203023-3120322210033100-2121220010332201-3133231021021220-1313001311100021-0313102021211212-0331322021202002)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-018.md#canonical-0112020200111022-0013103013103331-1110301001233101-1300202101312213-3301031210112102-1221232301300211-0231000232130131-3132113303133333)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options

<a id="canonical-2211010311233122-1100302100101032-2031332131201001-2231020212000312-2011210133003131-2210333033303101-3223311132302221-0011102101012321"></a>

Type: `"object"`. single nested block, Optional.

HTTP protocol configuration OPTIONS for downstream connections.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

<a id="canonical-2101323120122030-0230220032122033-1330132001132222-3131121330223101-0112122303122113-0112103000333103-3313313123021213-2310001310303113"></a>

### Direct properties for `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options`

- [http_protocol_enable_v1_only](resources--workload--reference--group-018.md#canonical-0101121203223101-2303112021302213-2210233131121322-1010121010210301-2330311311003110-2202230130203213-2331310302232000-1213011200013332): complete subsection reference.

- [http_protocol_enable_v1_v2](resources--workload--reference--group-018.md#canonical-1233012312003333-2101101012232120-0100310010232311-0011212101122211-2001031003213123-1031101223300032-3223031101001303-3201303031321011): complete subsection reference.

- [http_protocol_enable_v2_only](resources--workload--reference--group-018.md#canonical-2120212030201112-1220311010203002-0121211322222323-2032310203313301-2111110222212103-0222302013331310-1331023200202002-1220000133031023): complete subsection reference.

<a id="canonical-0101121203223101-2303112021302213-2210233131121322-1010121010210301-2330311311003110-2202230130203213-2331310302232000-1213011200013332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-017.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_custom](resources--workload--reference--group-017.md#canonical-0203102001320123-1120130012130212-0111101220231333-3101331123203121-1110000311003232-1321231201202312-3200200021132112-0200303231330002)
- [stateful_service.advertise_options.advertise_custom.ports](resources--workload--reference--group-017.md#canonical-3111311322332303-1201030020021201-3133021233000313-2022333331131202-3221312032323033-3133120120113232-3330021120310331-2331303302032231)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-017.md#canonical-0313222022003221-2301103230203023-3120322210033100-2121220010332201-3133231021021220-1313001311100021-0313102021211212-0331322021202002)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-018.md#canonical-0112020200111022-0013103013103331-1110301001233101-1300202101312213-3301031210112102-1221232301300211-0231000232130131-3132113303133333)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options](resources--workload--reference--group-018.md#canonical-3202212001033033-3133231302102230-0333203123220020-0022212313320130-3032131111301011-1301121203330231-2010200210001321-0233311332222011)
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

- [header_transformation](resources--workload--reference--group-018.md#canonical-3020121301012012-1323210110303103-1332333110211232-3101112232311132-3303103303122230-0233302131022013-3203113021013212-1223222232202303): complete subsection reference.

<a id="canonical-3020121301012012-1323210110303103-1332333110211232-3101112232311132-3303103303122230-0233302131022013-3203113021013212-1223222232202303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-017.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_custom](resources--workload--reference--group-017.md#canonical-0203102001320123-1120130012130212-0111101220231333-3101331123203121-1110000311003232-1321231201202312-3200200021132112-0200303231330002)
- [stateful_service.advertise_options.advertise_custom.ports](resources--workload--reference--group-017.md#canonical-3111311322332303-1201030020021201-3133021233000313-2022333331131202-3221312032323033-3133120120113232-3330021120310331-2331303302032231)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-017.md#canonical-0313222022003221-2301103230203023-3120322210033100-2121220010332201-3133231021021220-1313001311100021-0313102021211212-0331322021202002)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-018.md#canonical-0112020200111022-0013103013103331-1110301001233101-1300202101312213-3301031210112102-1221232301300211-0231000232130131-3132113303133333)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options](resources--workload--reference--group-018.md#canonical-3202212001033033-3133231302102230-0333203123220020-0022212313320130-3032131111301011-1301121203330231-2010200210001321-0233311332222011)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only](resources--workload--reference--group-018.md#canonical-0101121203223101-2303112021302213-2210233131121322-1010121010210301-2330311311003110-2202230130203213-2331310302232000-1213011200013332)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation

<a id="canonical-2310312012330001-3120203203330223-3010203020020200-1223111110000012-0110321020220022-0123212133301001-3301102112333013-1323310031230323"></a>

Type: `"object"`. single nested block, Optional.

Header Transformation OPTIONS for HTTP/1.1 request/response headers.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

<a id="canonical-3303003333022013-2311101013010221-2310012011301021-0213111113020312-0231221011102330-1301311200322221-3131232201100330-3212030200300313"></a>

### Direct properties for `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation`

- [default_header_transformation](resources--workload--reference--group-018.md#canonical-0010231113030312-1033233101131032-3112201133013213-3303112230223301-2331322231322331-0301320223323310-3212003203010320-1100132021202010): complete subsection reference.

- [preserve_case_header_transformation](resources--workload--reference--group-018.md#canonical-0122312232023030-1012112022323310-2200001103311021-3013321333232030-0133230321231123-1322213203213000-3031310310221221-0023133221022231): complete subsection reference.

- [proper_case_header_transformation](resources--workload--reference--group-018.md#canonical-3231002321200010-2030232222102330-2133211301012130-2211121233102332-3130131112312303-0001021113312310-1220012102302231-2113230131033030): complete subsection reference.

<a id="canonical-0010231113030312-1033233101131032-3112201133013213-3303112230223301-2331322231322331-0301320223323310-3212003203010320-1100132021202010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.default_header_transformation` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-017.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_custom](resources--workload--reference--group-017.md#canonical-0203102001320123-1120130012130212-0111101220231333-3101331123203121-1110000311003232-1321231201202312-3200200021132112-0200303231330002)
- [stateful_service.advertise_options.advertise_custom.ports](resources--workload--reference--group-017.md#canonical-3111311322332303-1201030020021201-3133021233000313-2022333331131202-3221312032323033-3133120120113232-3330021120310331-2331303302032231)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-017.md#canonical-0313222022003221-2301103230203023-3120322210033100-2121220010332201-3133231021021220-1313001311100021-0313102021211212-0331322021202002)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-018.md#canonical-0112020200111022-0013103013103331-1110301001233101-1300202101312213-3301031210112102-1221232301300211-0231000232130131-3132113303133333)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options](resources--workload--reference--group-018.md#canonical-3202212001033033-3133231302102230-0333203123220020-0022212313320130-3032131111301011-1301121203330231-2010200210001321-0233311332222011)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only](resources--workload--reference--group-018.md#canonical-0101121203223101-2303112021302213-2210233131121322-1010121010210301-2330311311003110-2202230130203213-2331310302232000-1213011200013332)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation](resources--workload--reference--group-018.md#canonical-3020121301012012-1323210110303103-1332333110211232-3101112232311132-3303103303122230-0233302131022013-3203113021013212-1223222232202303)
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
- [stateful_service](resources--workload--reference--group-017.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_custom](resources--workload--reference--group-017.md#canonical-0203102001320123-1120130012130212-0111101220231333-3101331123203121-1110000311003232-1321231201202312-3200200021132112-0200303231330002)
- [stateful_service.advertise_options.advertise_custom.ports](resources--workload--reference--group-017.md#canonical-3111311322332303-1201030020021201-3133021233000313-2022333331131202-3221312032323033-3133120120113232-3330021120310331-2331303302032231)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-017.md#canonical-0313222022003221-2301103230203023-3120322210033100-2121220010332201-3133231021021220-1313001311100021-0313102021211212-0331322021202002)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-018.md#canonical-0112020200111022-0013103013103331-1110301001233101-1300202101312213-3301031210112102-1221232301300211-0231000232130131-3132113303133333)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options](resources--workload--reference--group-018.md#canonical-3202212001033033-3133231302102230-0333203123220020-0022212313320130-3032131111301011-1301121203330231-2010200210001321-0233311332222011)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only](resources--workload--reference--group-018.md#canonical-0101121203223101-2303112021302213-2210233131121322-1010121010210301-2330311311003110-2202230130203213-2331310302232000-1213011200013332)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation](resources--workload--reference--group-018.md#canonical-3020121301012012-1323210110303103-1332333110211232-3101112232311132-3303103303122230-0233302131022013-3203113021013212-1223222232202303)
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
- [stateful_service](resources--workload--reference--group-017.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_custom](resources--workload--reference--group-017.md#canonical-0203102001320123-1120130012130212-0111101220231333-3101331123203121-1110000311003232-1321231201202312-3200200021132112-0200303231330002)
- [stateful_service.advertise_options.advertise_custom.ports](resources--workload--reference--group-017.md#canonical-3111311322332303-1201030020021201-3133021233000313-2022333331131202-3221312032323033-3133120120113232-3330021120310331-2331303302032231)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-017.md#canonical-0313222022003221-2301103230203023-3120322210033100-2121220010332201-3133231021021220-1313001311100021-0313102021211212-0331322021202002)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-018.md#canonical-0112020200111022-0013103013103331-1110301001233101-1300202101312213-3301031210112102-1221232301300211-0231000232130131-3132113303133333)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options](resources--workload--reference--group-018.md#canonical-3202212001033033-3133231302102230-0333203123220020-0022212313320130-3032131111301011-1301121203330231-2010200210001321-0233311332222011)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only](resources--workload--reference--group-018.md#canonical-0101121203223101-2303112021302213-2210233131121322-1010121010210301-2330311311003110-2202230130203213-2331310302232000-1213011200013332)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation](resources--workload--reference--group-018.md#canonical-3020121301012012-1323210110303103-1332333110211232-3101112232311132-3303103303122230-0233302131022013-3203113021013212-1223222232202303)
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
- [stateful_service](resources--workload--reference--group-017.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_custom](resources--workload--reference--group-017.md#canonical-0203102001320123-1120130012130212-0111101220231333-3101331123203121-1110000311003232-1321231201202312-3200200021132112-0200303231330002)
- [stateful_service.advertise_options.advertise_custom.ports](resources--workload--reference--group-017.md#canonical-3111311322332303-1201030020021201-3133021233000313-2022333331131202-3221312032323033-3133120120113232-3330021120310331-2331303302032231)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-017.md#canonical-0313222022003221-2301103230203023-3120322210033100-2121220010332201-3133231021021220-1313001311100021-0313102021211212-0331322021202002)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-018.md#canonical-0112020200111022-0013103013103331-1110301001233101-1300202101312213-3301031210112102-1221232301300211-0231000232130131-3132113303133333)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options](resources--workload--reference--group-018.md#canonical-3202212001033033-3133231302102230-0333203123220020-0022212313320130-3032131111301011-1301121203330231-2010200210001321-0233311332222011)
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
- [stateful_service](resources--workload--reference--group-017.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_custom](resources--workload--reference--group-017.md#canonical-0203102001320123-1120130012130212-0111101220231333-3101331123203121-1110000311003232-1321231201202312-3200200021132112-0200303231330002)
- [stateful_service.advertise_options.advertise_custom.ports](resources--workload--reference--group-017.md#canonical-3111311322332303-1201030020021201-3133021233000313-2022333331131202-3221312032323033-3133120120113232-3330021120310331-2331303302032231)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-017.md#canonical-0313222022003221-2301103230203023-3120322210033100-2121220010332201-3133231021021220-1313001311100021-0313102021211212-0331322021202002)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-018.md#canonical-0112020200111022-0013103013103331-1110301001233101-1300202101312213-3301031210112102-1221232301300211-0231000232130131-3132113303133333)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options](resources--workload--reference--group-018.md#canonical-3202212001033033-3133231302102230-0333203123220020-0022212313320130-3032131111301011-1301121203330231-2010200210001321-0233311332222011)
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
- [stateful_service](resources--workload--reference--group-017.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_custom](resources--workload--reference--group-017.md#canonical-0203102001320123-1120130012130212-0111101220231333-3101331123203121-1110000311003232-1321231201202312-3200200021132112-0200303231330002)
- [stateful_service.advertise_options.advertise_custom.ports](resources--workload--reference--group-017.md#canonical-3111311322332303-1201030020021201-3133021233000313-2022333331131202-3221312032323033-3133120120113232-3330021120310331-2331303302032231)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-017.md#canonical-0313222022003221-2301103230203023-3120322210033100-2121220010332201-3133231021021220-1313001311100021-0313102021211212-0331322021202002)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-018.md#canonical-0112020200111022-0013103013103331-1110301001233101-1300202101312213-3301031210112102-1221232301300211-0231000232130131-3132113303133333)
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
- [stateful_service](resources--workload--reference--group-017.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_custom](resources--workload--reference--group-017.md#canonical-0203102001320123-1120130012130212-0111101220231333-3101331123203121-1110000311003232-1321231201202312-3200200021132112-0200303231330002)
- [stateful_service.advertise_options.advertise_custom.ports](resources--workload--reference--group-017.md#canonical-3111311322332303-1201030020021201-3133021233000313-2022333331131202-3221312032323033-3133120120113232-3330021120310331-2331303302032231)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-017.md#canonical-0313222022003221-2301103230203023-3120322210033100-2121220010332201-3133231021021220-1313001311100021-0313102021211212-0331322021202002)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-018.md#canonical-0112020200111022-0013103013103331-1110301001233101-1300202101312213-3301031210112102-1221232301300211-0231000232130131-3132113303133333)
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
- [stateful_service](resources--workload--reference--group-017.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_custom](resources--workload--reference--group-017.md#canonical-0203102001320123-1120130012130212-0111101220231333-3101331123203121-1110000311003232-1321231201202312-3200200021132112-0200303231330002)
- [stateful_service.advertise_options.advertise_custom.ports](resources--workload--reference--group-017.md#canonical-3111311322332303-1201030020021201-3133021233000313-2022333331131202-3221312032323033-3133120120113232-3330021120310331-2331303302032231)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-017.md#canonical-0313222022003221-2301103230203023-3120322210033100-2121220010332201-3133231021021220-1313001311100021-0313102021211212-0331322021202002)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-018.md#canonical-0112020200111022-0013103013103331-1110301001233101-1300202101312213-3301031210112102-1221232301300211-0231000232130131-3132113303133333)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params

<a id="canonical-2013311010020320-2121113203323031-0200002033113131-0312102310321003-1232023020001013-1303232221320223-1300230002113120-2202133130120031"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for tls cert params.

Additional upstream details:

Select TLS Parameters and Certificates.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

<a id="canonical-2312210020320113-0110232123202223-1200202221310021-2313230111201022-2200101203022023-3030133301312320-1222303002323232-0333202020321323"></a>

### Direct properties for `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params`

- [certificates](resources--workload--reference--group-018.md#canonical-2302321201330300-1112012230113032-0022112020230310-3222302030221312-2011131033323122-1230110201111231-3102300001021111-1202003110201010): complete subsection reference.

- [no_mtls](resources--workload--reference--group-018.md#canonical-3201210010322203-1131010200302302-1323313133302012-2032121310301010-1331220210000022-2132221211132332-1313000211301001-0210022102033110): complete subsection reference.

- [tls_config](resources--workload--reference--group-018.md#canonical-0111002012300220-0123100132331012-0322301201120202-3013312031133201-3021203131320032-0333213333200111-2120023323002022-1203103032321031): complete subsection reference.

- [use_mtls](resources--workload--reference--group-018.md#canonical-1232000222132210-1231321130032102-2312300020102111-2032003122222202-0033301200101032-3330213032031032-2302102033201003-3021232302102100): complete subsection reference.

<a id="canonical-2302321201330300-1112012230113032-0022112020230310-3222302030221312-2011131033323122-1230110201111231-3102300001021111-1202003110201010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.certificates` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-017.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_custom](resources--workload--reference--group-017.md#canonical-0203102001320123-1120130012130212-0111101220231333-3101331123203121-1110000311003232-1321231201202312-3200200021132112-0200303231330002)
- [stateful_service.advertise_options.advertise_custom.ports](resources--workload--reference--group-017.md#canonical-3111311322332303-1201030020021201-3133021233000313-2022333331131202-3221312032323033-3133120120113232-3330021120310331-2331303302032231)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-017.md#canonical-0313222022003221-2301103230203023-3120322210033100-2121220010332201-3133231021021220-1313001311100021-0313102021211212-0331322021202002)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-018.md#canonical-0112020200111022-0013103013103331-1110301001233101-1300202101312213-3301031210112102-1221232301300211-0231000232130131-3132113303133333)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params](resources--workload--reference--group-018.md#canonical-1330323120222300-2123201010331320-0300232300332322-0010223222211002-1032302312300112-0233320002220301-1020331113013010-3202200223121221)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.certificates

<a id="canonical-0201121111210100-0200012131011222-2120033303323132-0301030021231123-0130003111322332-2100100113012133-1320110123020220-2000302112313322"></a>

Type: `"object"`. list nested block, Optional.

Select one or more certificates with any domain names.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

<a id="canonical-1120323113223010-0222330233330223-1203321001212003-0303113201203231-1033310103333223-2101032231131133-1021002230300000-2030332020302230"></a>

### Direct properties for `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.certificates`

<a id="canonical-1102121100221310-1132233333131331-3023321201120332-1033322132322113-0113110020011011-0002002321232103-0020121313212230-0113310223010203"></a>

#### `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.certificates.name` property

Type: `"string"`. Optional.

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

<a id="canonical-1032112101323332-0033012000013202-1000321222033011-2232333312012300-3321233231101222-3123131131301103-3110121000330202-2301013030022230"></a>

<a id="canonical-0030000120300230-2031003032002120-0133101312233003-0111233131230211-3301331333130103-3133110132130113-3220031011303100-0102312212113120"></a>

#### `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.certificates.namespace` property

Type: `"string"`. Optional, Computed.

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

<a id="canonical-2303331333212212-3221121001230000-3133122231211102-3231201230333012-0210010032210313-2032203321322311-3013220121131313-1201010310013212"></a>

<a id="canonical-2002130230331223-0311122122112311-2122332022102313-3312111322030010-2133313223002303-3213222300031210-0101031132303302-3201222202223311"></a>

#### `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.certificates.tenant` property

Type: `"string"`. Computed.

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

<a id="canonical-3201210010322203-1131010200302302-1323313133302012-2032121310301010-1331220210000022-2132221211132332-1313000211301001-0210022102033110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.no_mtls` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-017.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_custom](resources--workload--reference--group-017.md#canonical-0203102001320123-1120130012130212-0111101220231333-3101331123203121-1110000311003232-1321231201202312-3200200021132112-0200303231330002)
- [stateful_service.advertise_options.advertise_custom.ports](resources--workload--reference--group-017.md#canonical-3111311322332303-1201030020021201-3133021233000313-2022333331131202-3221312032323033-3133120120113232-3330021120310331-2331303302032231)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-017.md#canonical-0313222022003221-2301103230203023-3120322210033100-2121220010332201-3133231021021220-1313001311100021-0313102021211212-0331322021202002)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-018.md#canonical-0112020200111022-0013103013103331-1110301001233101-1300202101312213-3301031210112102-1221232301300211-0231000232130131-3132113303133333)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params](resources--workload--reference--group-018.md#canonical-1330323120222300-2123201010331320-0300232300332322-0010223222211002-1032302312300112-0233320002220301-1020331113013010-3202200223121221)
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
- [stateful_service](resources--workload--reference--group-017.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_custom](resources--workload--reference--group-017.md#canonical-0203102001320123-1120130012130212-0111101220231333-3101331123203121-1110000311003232-1321231201202312-3200200021132112-0200303231330002)
- [stateful_service.advertise_options.advertise_custom.ports](resources--workload--reference--group-017.md#canonical-3111311322332303-1201030020021201-3133021233000313-2022333331131202-3221312032323033-3133120120113232-3330021120310331-2331303302032231)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-017.md#canonical-0313222022003221-2301103230203023-3120322210033100-2121220010332201-3133231021021220-1313001311100021-0313102021211212-0331322021202002)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-018.md#canonical-0112020200111022-0013103013103331-1110301001233101-1300202101312213-3301031210112102-1221232301300211-0231000232130131-3132113303133333)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params](resources--workload--reference--group-018.md#canonical-1330323120222300-2123201010331320-0300232300332322-0010223222211002-1032302312300112-0233320002220301-1020331113013010-3202200223121221)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.tls_config

<a id="canonical-1330121113200001-0121232310331130-1023030122233102-2022222223011132-2230030011131320-3303220033303201-1200131312102033-1201310202111312"></a>

Type: `"object"`. single nested block, Optional.

This defines various OPTIONS to configure TLS configuration parameters.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

<a id="canonical-0003213032011103-3210202201110322-1033303301010311-2022113230212010-1020032123012220-2000310331023033-0332220313130332-0201300221132331"></a>

### Direct properties for `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.tls_config`

- [custom_security](resources--workload--reference--group-018.md#canonical-0213222020222022-2033230011011000-3133220133331122-2313032021200230-1233322200222033-1202020120332300-3221303013103010-1321233123110113): complete subsection reference.

- [default_security](resources--workload--reference--group-018.md#canonical-3113133100131011-1010333021003230-1123210303221313-1323213322031002-2013211001200120-1123330132001213-0132320122300003-0320201222013110): complete subsection reference.

- [low_security](resources--workload--reference--group-018.md#canonical-0213030220123100-3120123030111302-2220121322132023-3231112302000021-2030031312331030-3322131203122302-0220131033022120-2131002102103231): complete subsection reference.

- [medium_security](resources--workload--reference--group-018.md#canonical-3333323101212103-0100322011331322-0122123210000023-0320312011231232-2120011020130202-0223332023111032-1200223021123321-1001112202313303): complete subsection reference.

<a id="canonical-0213222020222022-2033230011011000-3133220133331122-2313032021200230-1233322200222033-1202020120332300-3221303013103010-1321233123110113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.tls_config.custom_security` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-017.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_custom](resources--workload--reference--group-017.md#canonical-0203102001320123-1120130012130212-0111101220231333-3101331123203121-1110000311003232-1321231201202312-3200200021132112-0200303231330002)
- [stateful_service.advertise_options.advertise_custom.ports](resources--workload--reference--group-017.md#canonical-3111311322332303-1201030020021201-3133021233000313-2022333331131202-3221312032323033-3133120120113232-3330021120310331-2331303302032231)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-017.md#canonical-0313222022003221-2301103230203023-3120322210033100-2121220010332201-3133231021021220-1313001311100021-0313102021211212-0331322021202002)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-018.md#canonical-0112020200111022-0013103013103331-1110301001233101-1300202101312213-3301031210112102-1221232301300211-0231000232130131-3132113303133333)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params](resources--workload--reference--group-018.md#canonical-1330323120222300-2123201010331320-0300232300332322-0010223222211002-1032302312300112-0233320002220301-1020331113013010-3202200223121221)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.tls_config](resources--workload--reference--group-018.md#canonical-0111002012300220-0123100132331012-0322301201120202-3013312031133201-3021203131320032-0333213333200111-2120023323002022-1203103032321031)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.tls_config.custom_security

<a id="canonical-0030022202022032-0123331111102011-0113102120202322-0023213302201331-3110330131200323-2103030111223001-3300002033120300-0132122322313201"></a>

Type: `"object"`. single nested block, Optional.

This defines TLS protocol config including min/max versions and allowed ciphers.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

<a id="canonical-3023321230220101-2000011132122230-3220333001032012-0113300022031231-3012010231233013-1320102002003212-2101132102312301-0212130100203313"></a>

<a id="canonical-2020301132101202-3100031233322132-1323320120223322-3030003133131212-2311000213311320-2002120121121002-3332032211132330-2013301230120000"></a>

#### `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.tls_config.custom_security.max_version` property

Type: `"string"`. Optional.

\[Enum: TLS\_AUTO|TLSv1\_0|TLSv1\_1|TLSv1\_2|TLSv1\_3\] TlsProtocol is enumeration of supported TLS
versions F5 Distributed Cloud will choose the optimal TLS version. Possible values are
\`TLS\_AUTO\`, \`TLSv1\_0\`, \`TLSv1\_1\`, \`TLSv1\_2\`, \`TLSv1\_3\`. Defaults to \`TLS\_AUTO\`.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["TLS_AUTO","TLSv1_0","TLSv1_1","TLSv1_2","TLSv1_3"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
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

<a id="canonical-2001313221103323-1222010332312230-2120231110223232-3210210032310210-2003130110220232-1202022332200030-2022310333120321-3033012230012131"></a>

<a id="canonical-1230020232210133-0123003332321330-1011111321220030-2133223300231002-3320321033000122-2333010112002303-1032323210313111-1121201322333020"></a>

#### `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.tls_config.custom_security.min_version` property

Type: `"string"`. Optional.

\[Enum: TLS\_AUTO|TLSv1\_0|TLSv1\_1|TLSv1\_2|TLSv1\_3\] TlsProtocol is enumeration of supported TLS
versions F5 Distributed Cloud will choose the optimal TLS version. Possible values are
\`TLS\_AUTO\`, \`TLSv1\_0\`, \`TLSv1\_1\`, \`TLSv1\_2\`, \`TLSv1\_3\`. Defaults to \`TLS\_AUTO\`.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["TLS_AUTO","TLSv1_0","TLSv1_1","TLSv1_2","TLSv1_3"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
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

<a id="canonical-3113133100131011-1010333021003230-1123210303221313-1323213322031002-2013211001200120-1123330132001213-0132320122300003-0320201222013110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.tls_config.default_security` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-017.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_custom](resources--workload--reference--group-017.md#canonical-0203102001320123-1120130012130212-0111101220231333-3101331123203121-1110000311003232-1321231201202312-3200200021132112-0200303231330002)
- [stateful_service.advertise_options.advertise_custom.ports](resources--workload--reference--group-017.md#canonical-3111311322332303-1201030020021201-3133021233000313-2022333331131202-3221312032323033-3133120120113232-3330021120310331-2331303302032231)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-017.md#canonical-0313222022003221-2301103230203023-3120322210033100-2121220010332201-3133231021021220-1313001311100021-0313102021211212-0331322021202002)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-018.md#canonical-0112020200111022-0013103013103331-1110301001233101-1300202101312213-3301031210112102-1221232301300211-0231000232130131-3132113303133333)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params](resources--workload--reference--group-018.md#canonical-1330323120222300-2123201010331320-0300232300332322-0010223222211002-1032302312300112-0233320002220301-1020331113013010-3202200223121221)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.tls_config](resources--workload--reference--group-018.md#canonical-0111002012300220-0123100132331012-0322301201120202-3013312031133201-3021203131320032-0333213333200111-2120023323002022-1203103032321031)
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
- [stateful_service](resources--workload--reference--group-017.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_custom](resources--workload--reference--group-017.md#canonical-0203102001320123-1120130012130212-0111101220231333-3101331123203121-1110000311003232-1321231201202312-3200200021132112-0200303231330002)
- [stateful_service.advertise_options.advertise_custom.ports](resources--workload--reference--group-017.md#canonical-3111311322332303-1201030020021201-3133021233000313-2022333331131202-3221312032323033-3133120120113232-3330021120310331-2331303302032231)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-017.md#canonical-0313222022003221-2301103230203023-3120322210033100-2121220010332201-3133231021021220-1313001311100021-0313102021211212-0331322021202002)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-018.md#canonical-0112020200111022-0013103013103331-1110301001233101-1300202101312213-3301031210112102-1221232301300211-0231000232130131-3132113303133333)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params](resources--workload--reference--group-018.md#canonical-1330323120222300-2123201010331320-0300232300332322-0010223222211002-1032302312300112-0233320002220301-1020331113013010-3202200223121221)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.tls_config](resources--workload--reference--group-018.md#canonical-0111002012300220-0123100132331012-0322301201120202-3013312031133201-3021203131320032-0333213333200111-2120023323002022-1203103032321031)
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
- [stateful_service](resources--workload--reference--group-017.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_custom](resources--workload--reference--group-017.md#canonical-0203102001320123-1120130012130212-0111101220231333-3101331123203121-1110000311003232-1321231201202312-3200200021132112-0200303231330002)
- [stateful_service.advertise_options.advertise_custom.ports](resources--workload--reference--group-017.md#canonical-3111311322332303-1201030020021201-3133021233000313-2022333331131202-3221312032323033-3133120120113232-3330021120310331-2331303302032231)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-017.md#canonical-0313222022003221-2301103230203023-3120322210033100-2121220010332201-3133231021021220-1313001311100021-0313102021211212-0331322021202002)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-018.md#canonical-0112020200111022-0013103013103331-1110301001233101-1300202101312213-3301031210112102-1221232301300211-0231000232130131-3132113303133333)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params](resources--workload--reference--group-018.md#canonical-1330323120222300-2123201010331320-0300232300332322-0010223222211002-1032302312300112-0233320002220301-1020331113013010-3202200223121221)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.tls_config](resources--workload--reference--group-018.md#canonical-0111002012300220-0123100132331012-0322301201120202-3013312031133201-3021203131320032-0333213333200111-2120023323002022-1203103032321031)
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
- [stateful_service](resources--workload--reference--group-017.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_custom](resources--workload--reference--group-017.md#canonical-0203102001320123-1120130012130212-0111101220231333-3101331123203121-1110000311003232-1321231201202312-3200200021132112-0200303231330002)
- [stateful_service.advertise_options.advertise_custom.ports](resources--workload--reference--group-017.md#canonical-3111311322332303-1201030020021201-3133021233000313-2022333331131202-3221312032323033-3133120120113232-3330021120310331-2331303302032231)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-017.md#canonical-0313222022003221-2301103230203023-3120322210033100-2121220010332201-3133231021021220-1313001311100021-0313102021211212-0331322021202002)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-018.md#canonical-0112020200111022-0013103013103331-1110301001233101-1300202101312213-3301031210112102-1221232301300211-0231000232130131-3132113303133333)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params](resources--workload--reference--group-018.md#canonical-1330323120222300-2123201010331320-0300232300332322-0010223222211002-1032302312300112-0233320002220301-1020331113013010-3202200223121221)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.use_mtls

<a id="canonical-0200231003322122-0230312323230201-3102330221133103-2121230332220211-0332123200012023-2203231013112201-1210210230122033-0133013330010313"></a>

Type: `"object"`. single nested block, Optional.

Validation context for downstream client TLS connections.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

- [crl](resources--workload--reference--group-018.md#canonical-1023310123331030-3330002101023131-2133232213203312-1203130332202311-2100220033333301-3323102331333230-1211233121310122-0121013103030011): complete subsection reference.

- [no_crl](resources--workload--reference--group-018.md#canonical-2233232331221312-2213100000032221-0112211013300212-2030330022220310-1331233012301211-2010110032211230-3223302023221310-3002312000203332): complete subsection reference.

- [trusted_ca](resources--workload--reference--group-018.md#canonical-2123331130310033-0203031003112001-0331233220320302-3230103032202121-3311302330231202-3002332101200103-3003111103223202-0222221332311201): complete subsection reference.

<a id="canonical-1222233131000331-3300123122231212-1022112232101120-1122011121033101-3100110122102231-2030312130221031-0013203133031000-1320310001313232"></a>

<a id="canonical-3012130210321300-2310021000001212-2200213232221030-2021231002012222-0100110202012213-0320223320010023-0131101301201001-2023212102202300"></a>

#### `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.use_mtls.trusted_ca_url` property

Type: `"string"`. Optional.

Exclusive with \[trusted\_ca\] Upload a Root CA Certificate specifically for this Load Balancer.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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

- [xfcc_disabled](resources--workload--reference--group-018.md#canonical-2113333000131311-3302120001323220-1200201202012001-1210000301020103-3310002113310230-3013100321100011-0231203223000031-3313332002011331): complete subsection reference.

- [xfcc_options](resources--workload--reference--group-018.md#canonical-3133132021133213-3020123203123012-3203121321220120-1000031311310301-2322312002221310-3110133230331222-1103100011120133-0000230331013210): complete subsection reference.

<a id="canonical-1023310123331030-3330002101023131-2133232213203312-1203130332202311-2100220033333301-3323102331333230-1211233121310122-0121013103030011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.use_mtls.crl` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-017.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_custom](resources--workload--reference--group-017.md#canonical-0203102001320123-1120130012130212-0111101220231333-3101331123203121-1110000311003232-1321231201202312-3200200021132112-0200303231330002)
- [stateful_service.advertise_options.advertise_custom.ports](resources--workload--reference--group-017.md#canonical-3111311322332303-1201030020021201-3133021233000313-2022333331131202-3221312032323033-3133120120113232-3330021120310331-2331303302032231)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-017.md#canonical-0313222022003221-2301103230203023-3120322210033100-2121220010332201-3133231021021220-1313001311100021-0313102021211212-0331322021202002)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-018.md#canonical-0112020200111022-0013103013103331-1110301001233101-1300202101312213-3301031210112102-1221232301300211-0231000232130131-3132113303133333)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params](resources--workload--reference--group-018.md#canonical-1330323120222300-2123201010331320-0300232300332322-0010223222211002-1032302312300112-0233320002220301-1020331113013010-3202200223121221)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.use_mtls](resources--workload--reference--group-018.md#canonical-1232000222132210-1231321130032102-2312300020102111-2032003122222202-0033301200101032-3330213032031032-2302102033201003-3021232302102100)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.use_mtls.crl

<a id="canonical-2230203333012010-3231102212313302-1013322220223002-0312203032323320-1023030110102233-0033231132231211-1303103012310330-0220312322301333"></a>

Type: `"object"`. single nested block, Optional.

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

<a id="canonical-3302230031301102-1300311301200320-3221120002111012-3012313221230300-0111222302132120-0022211232132000-0200232101133232-2211321012300033"></a>

<a id="canonical-3020332113013000-0102131222010033-1000030100110031-3000221011213230-1133311130233200-1310331012221011-1221020022231102-3000100121113020"></a>

#### `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.use_mtls.crl.namespace` property

Type: `"string"`. Optional, Computed.

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

<a id="canonical-2211030221310232-0102111113321321-3231113132231012-3231323022213223-3320202002122123-1302033310333301-3010020102100102-3110300120333033"></a>

<a id="canonical-0331003131101212-0133132021310101-2030020321322100-3130001000100011-3302332133330313-1121031311122322-2023110331332001-2211030013002000"></a>

#### `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.use_mtls.crl.tenant` property

Type: `"string"`. Computed.

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

<a id="canonical-2233232331221312-2213100000032221-0112211013300212-2030330022220310-1331233012301211-2010110032211230-3223302023221310-3002312000203332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.use_mtls.no_crl` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-017.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_custom](resources--workload--reference--group-017.md#canonical-0203102001320123-1120130012130212-0111101220231333-3101331123203121-1110000311003232-1321231201202312-3200200021132112-0200303231330002)
- [stateful_service.advertise_options.advertise_custom.ports](resources--workload--reference--group-017.md#canonical-3111311322332303-1201030020021201-3133021233000313-2022333331131202-3221312032323033-3133120120113232-3330021120310331-2331303302032231)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-017.md#canonical-0313222022003221-2301103230203023-3120322210033100-2121220010332201-3133231021021220-1313001311100021-0313102021211212-0331322021202002)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-018.md#canonical-0112020200111022-0013103013103331-1110301001233101-1300202101312213-3301031210112102-1221232301300211-0231000232130131-3132113303133333)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params](resources--workload--reference--group-018.md#canonical-1330323120222300-2123201010331320-0300232300332322-0010223222211002-1032302312300112-0233320002220301-1020331113013010-3202200223121221)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.use_mtls](resources--workload--reference--group-018.md#canonical-1232000222132210-1231321130032102-2312300020102111-2032003122222202-0033301200101032-3330213032031032-2302102033201003-3021232302102100)
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
- [stateful_service](resources--workload--reference--group-017.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_custom](resources--workload--reference--group-017.md#canonical-0203102001320123-1120130012130212-0111101220231333-3101331123203121-1110000311003232-1321231201202312-3200200021132112-0200303231330002)
- [stateful_service.advertise_options.advertise_custom.ports](resources--workload--reference--group-017.md#canonical-3111311322332303-1201030020021201-3133021233000313-2022333331131202-3221312032323033-3133120120113232-3330021120310331-2331303302032231)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-017.md#canonical-0313222022003221-2301103230203023-3120322210033100-2121220010332201-3133231021021220-1313001311100021-0313102021211212-0331322021202002)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-018.md#canonical-0112020200111022-0013103013103331-1110301001233101-1300202101312213-3301031210112102-1221232301300211-0231000232130131-3132113303133333)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params](resources--workload--reference--group-018.md#canonical-1330323120222300-2123201010331320-0300232300332322-0010223222211002-1032302312300112-0233320002220301-1020331113013010-3202200223121221)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.use_mtls](resources--workload--reference--group-018.md#canonical-1232000222132210-1231321130032102-2312300020102111-2032003122222202-0033301200101032-3330213032031032-2302102033201003-3021232302102100)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.use_mtls.trusted_ca

<a id="canonical-1122313111020022-0223202120313130-3213022010133302-0023322122301201-2232213211003010-3301112330033302-2030111212102220-1000210113130303"></a>

Type: `"object"`. single nested block, Optional.

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

<a id="canonical-3130202023332312-3132010130011322-2003202220301010-1101233313120120-3113032221302330-2322302313101222-2301031012212321-3110312131333232"></a>

<a id="canonical-0323203021001210-1300132331203322-0303003113300130-1023031100132210-2320321120032003-1203333201030322-2023113233310213-0002010201332022"></a>

#### `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.use_mtls.trusted_ca.namespace` property

Type: `"string"`. Optional, Computed.

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

<a id="canonical-2221322223101212-1232132201202231-2102322232302130-3221313322011001-0201312030333303-0302102021130121-0330132021331113-3302330332033002"></a>

<a id="canonical-1220112001032103-1023310213323232-1303332130123101-3331223211002302-1122031310201032-2133112130022221-2323232001202012-3122322001311022"></a>

#### `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.use_mtls.trusted_ca.tenant` property

Type: `"string"`. Computed.

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

<a id="canonical-2113333000131311-3302120001323220-1200201202012001-1210000301020103-3310002113310230-3013100321100011-0231203223000031-3313332002011331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.use_mtls.xfcc_disabled` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-017.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_custom](resources--workload--reference--group-017.md#canonical-0203102001320123-1120130012130212-0111101220231333-3101331123203121-1110000311003232-1321231201202312-3200200021132112-0200303231330002)
- [stateful_service.advertise_options.advertise_custom.ports](resources--workload--reference--group-017.md#canonical-3111311322332303-1201030020021201-3133021233000313-2022333331131202-3221312032323033-3133120120113232-3330021120310331-2331303302032231)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-017.md#canonical-0313222022003221-2301103230203023-3120322210033100-2121220010332201-3133231021021220-1313001311100021-0313102021211212-0331322021202002)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-018.md#canonical-0112020200111022-0013103013103331-1110301001233101-1300202101312213-3301031210112102-1221232301300211-0231000232130131-3132113303133333)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params](resources--workload--reference--group-018.md#canonical-1330323120222300-2123201010331320-0300232300332322-0010223222211002-1032302312300112-0233320002220301-1020331113013010-3202200223121221)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.use_mtls](resources--workload--reference--group-018.md#canonical-1232000222132210-1231321130032102-2312300020102111-2032003122222202-0033301200101032-3330213032031032-2302102033201003-3021232302102100)
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
- [stateful_service](resources--workload--reference--group-017.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_custom](resources--workload--reference--group-017.md#canonical-0203102001320123-1120130012130212-0111101220231333-3101331123203121-1110000311003232-1321231201202312-3200200021132112-0200303231330002)
- [stateful_service.advertise_options.advertise_custom.ports](resources--workload--reference--group-017.md#canonical-3111311322332303-1201030020021201-3133021233000313-2022333331131202-3221312032323033-3133120120113232-3330021120310331-2331303302032231)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-017.md#canonical-0313222022003221-2301103230203023-3120322210033100-2121220010332201-3133231021021220-1313001311100021-0313102021211212-0331322021202002)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-018.md#canonical-0112020200111022-0013103013103331-1110301001233101-1300202101312213-3301031210112102-1221232301300211-0231000232130131-3132113303133333)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params](resources--workload--reference--group-018.md#canonical-1330323120222300-2123201010331320-0300232300332322-0010223222211002-1032302312300112-0233320002220301-1020331113013010-3202200223121221)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.use_mtls](resources--workload--reference--group-018.md#canonical-1232000222132210-1231321130032102-2312300020102111-2032003122222202-0033301200101032-3330213032031032-2302102033201003-3021232302102100)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.use_mtls.xfcc_options

<a id="canonical-0023100130303320-3111213311322223-1222231020023313-1032333210231030-0132132122101302-2322303013222203-1232320001001233-0222220022301103"></a>

Type: `"object"`. single nested block, Optional.

X-Forwarded-Client-Cert header elements to be added to requests.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("xfcc_header_elements")}
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
- [stateful_service](resources--workload--reference--group-017.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_custom](resources--workload--reference--group-017.md#canonical-0203102001320123-1120130012130212-0111101220231333-3101331123203121-1110000311003232-1321231201202312-3200200021132112-0200303231330002)
- [stateful_service.advertise_options.advertise_custom.ports](resources--workload--reference--group-017.md#canonical-3111311322332303-1201030020021201-3133021233000313-2022333331131202-3221312032323033-3133120120113232-3330021120310331-2331303302032231)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-017.md#canonical-0313222022003221-2301103230203023-3120322210033100-2121220010332201-3133231021021220-1313001311100021-0313102021211212-0331322021202002)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-018.md#canonical-0112020200111022-0013103013103331-1110301001233101-1300202101312213-3301031210112102-1221232301300211-0231000232130131-3132113303133333)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters

<a id="canonical-3012320331111130-3320032303011332-3210332231220013-0033203303111011-0320003331320201-0200102302301013-0200000133221302-3210323313131221"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for tls parameters.

Additional upstream details:

Inline TLS parameters.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("tls_certificates"),
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
tls_parameters {
  # Configure direct properties listed below.
}
```

<a id="canonical-3021220313010201-1223103122123332-1031102111313231-2122011022023201-2213103221131232-2133032132320013-1310322121312203-2210313202013202"></a>

### Direct properties for `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters`

- [no_mtls](resources--workload--reference--group-018.md#canonical-2033222013223233-0320010131321230-3001211111320031-2013111220310333-2201211223101123-2233002331000331-0130220100213002-1313101033313212): complete subsection reference.

- [tls_certificates](resources--workload--reference--group-018.md#canonical-3332313302002330-1333101212032003-2203023132232213-0001033110210120-3121103210011323-2020220120012213-3220330021330022-1301300222033110): complete subsection reference.

- [tls_config](resources--workload--reference--group-018.md#canonical-3212132103211102-3103133011133021-3313221110320101-2111103132233113-1232201321022230-0232111013133132-1231122000001132-3012311030121223): complete subsection reference.

- [use_mtls](resources--workload--reference--group-019.md#canonical-0002002033132331-3111321121121321-2300013122120330-2032012003011010-3002010130131223-0102231301201330-0023121332032121-3203220222232210): complete subsection reference.

<a id="canonical-2033222013223233-0320010131321230-3001211111320031-2013111220310333-2201211223101123-2233002331000331-0130220100213002-1313101033313212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.no_mtls` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-017.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_custom](resources--workload--reference--group-017.md#canonical-0203102001320123-1120130012130212-0111101220231333-3101331123203121-1110000311003232-1321231201202312-3200200021132112-0200303231330002)
- [stateful_service.advertise_options.advertise_custom.ports](resources--workload--reference--group-017.md#canonical-3111311322332303-1201030020021201-3133021233000313-2022333331131202-3221312032323033-3133120120113232-3330021120310331-2331303302032231)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-017.md#canonical-0313222022003221-2301103230203023-3120322210033100-2121220010332201-3133231021021220-1313001311100021-0313102021211212-0331322021202002)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-018.md#canonical-0112020200111022-0013103013103331-1110301001233101-1300202101312213-3301031210112102-1221232301300211-0231000232130131-3132113303133333)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-018.md#canonical-2020320003013033-2231100311132300-2132132133022132-3002120002132203-3131231201330331-0231012122032023-1202211010110230-0230013122122221)
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
- [stateful_service](resources--workload--reference--group-017.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_custom](resources--workload--reference--group-017.md#canonical-0203102001320123-1120130012130212-0111101220231333-3101331123203121-1110000311003232-1321231201202312-3200200021132112-0200303231330002)
- [stateful_service.advertise_options.advertise_custom.ports](resources--workload--reference--group-017.md#canonical-3111311322332303-1201030020021201-3133021233000313-2022333331131202-3221312032323033-3133120120113232-3330021120310331-2331303302032231)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-017.md#canonical-0313222022003221-2301103230203023-3120322210033100-2121220010332201-3133231021021220-1313001311100021-0313102021211212-0331322021202002)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-018.md#canonical-0112020200111022-0013103013103331-1110301001233101-1300202101312213-3301031210112102-1221232301300211-0231000232130131-3132113303133333)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-018.md#canonical-2020320003013033-2231100311132300-2132132133022132-3002120002132203-3131231201330331-0231012122032023-1202211010110230-0230013122122221)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates

<a id="canonical-0000202233313232-3310231000020221-2113121311203110-2313010331022330-0111010012233330-1032332113131122-1302210011112001-2020133212111213"></a>

Type: `"object"`. list nested block, Optional.

Users can add one or more certificates that share the same set of domains. For example, domain.com
and \*.domain.com - but use different signature algorithms.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.RequiredListObjectAttributes("certificate_url"),
  validators.ConflictingListObjectAttributes("custom_hash_algorithms",
    "disable_ocsp_stapling"),
  validators.ConflictingListObjectAttributes("custom_hash_algorithms",
    "use_system_defaults"),
  validators.ConflictingListObjectAttributes("disable_ocsp_stapling",
    "use_system_defaults")}
```

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-2203023033022223-3132313210333231-3111113313203020-2121330033103322-3333001110122233-3001201213200332-1332110202313010-0202133311321132"></a>

#### `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates.certificate_url` property

Type: `"string"`. Optional.

TLS certificate. Certificate or certificate chain in PEM format including the PEM headers.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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

- [custom_hash_algorithms](resources--workload--reference--group-018.md#canonical-0321322131200210-0113302313033133-1020201133022012-0110122211211203-0033010232003213-2113013333102202-0333322013220101-1133131233131021): complete subsection reference.

<a id="canonical-0212320032222011-2013032302300231-2202033111221110-1202301023300030-0013200020201022-2031003211330001-0223022110032213-0112011223123113"></a>

<a id="canonical-0213103000210133-0233201031320130-3123323333002013-3122200132313020-1301303123100323-2001031021132032-3121311331223023-0011331123220102"></a>

#### `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates.description_spec` property

Type: `"string"`. Optional.

Description. Description for the certificate.

- [disable_ocsp_stapling](resources--workload--reference--group-018.md#canonical-3301303000233223-1331301220303233-2130110233111130-1022311112112031-0010323212302010-2320331230300113-1233002033103101-1021212200303002): complete subsection reference.

- [private_key](resources--workload--reference--group-018.md#canonical-3202322131133212-3233333112022031-2012323022120230-3213102233121101-2023030003200212-0221230103200223-2321231323033110-1310133003012200): complete subsection reference.

- [use_system_defaults](resources--workload--reference--group-018.md#canonical-0322223320021121-2322111312221300-0121032323323102-3201122311322132-2130230200033330-3132003303112111-0313001301113023-1022311120311313): complete subsection reference.

<a id="canonical-0321322131200210-0113302313033133-1020201133022012-0110122211211203-0033010232003213-2113013333102202-0333322013220101-1133131233131021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates.custom_hash_algorithms` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-017.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_custom](resources--workload--reference--group-017.md#canonical-0203102001320123-1120130012130212-0111101220231333-3101331123203121-1110000311003232-1321231201202312-3200200021132112-0200303231330002)
- [stateful_service.advertise_options.advertise_custom.ports](resources--workload--reference--group-017.md#canonical-3111311322332303-1201030020021201-3133021233000313-2022333331131202-3221312032323033-3133120120113232-3330021120310331-2331303302032231)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-017.md#canonical-0313222022003221-2301103230203023-3120322210033100-2121220010332201-3133231021021220-1313001311100021-0313102021211212-0331322021202002)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-018.md#canonical-0112020200111022-0013103013103331-1110301001233101-1300202101312213-3301031210112102-1221232301300211-0231000232130131-3132113303133333)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-018.md#canonical-2020320003013033-2231100311132300-2132132133022132-3002120002132203-3131231201330331-0231012122032023-1202211010110230-0230013122122221)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates](resources--workload--reference--group-018.md#canonical-3332313302002330-1333101212032003-2203023132232213-0001033110210120-3121103210011323-2020220120012213-3220330021330022-1301300222033110)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates.custom_hash_algorithms

<a id="canonical-1001302023300321-0001301313320231-1013111022332231-1223020300102221-3222122300212213-2001313310231022-1001200201033033-0030020033200201"></a>

Type: `"object"`. single nested block, Optional.

Specifies the hash algorithms to be used.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("hash_algorithms")}
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

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeBetween(1, 4),
}
```

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
- [stateful_service](resources--workload--reference--group-017.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_custom](resources--workload--reference--group-017.md#canonical-0203102001320123-1120130012130212-0111101220231333-3101331123203121-1110000311003232-1321231201202312-3200200021132112-0200303231330002)
- [stateful_service.advertise_options.advertise_custom.ports](resources--workload--reference--group-017.md#canonical-3111311322332303-1201030020021201-3133021233000313-2022333331131202-3221312032323033-3133120120113232-3330021120310331-2331303302032231)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-017.md#canonical-0313222022003221-2301103230203023-3120322210033100-2121220010332201-3133231021021220-1313001311100021-0313102021211212-0331322021202002)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-018.md#canonical-0112020200111022-0013103013103331-1110301001233101-1300202101312213-3301031210112102-1221232301300211-0231000232130131-3132113303133333)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-018.md#canonical-2020320003013033-2231100311132300-2132132133022132-3002120002132203-3131231201330331-0231012122032023-1202211010110230-0230013122122221)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates](resources--workload--reference--group-018.md#canonical-3332313302002330-1333101212032003-2203023132232213-0001033110210120-3121103210011323-2020220120012213-3220330021330022-1301300222033110)
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
- [stateful_service](resources--workload--reference--group-017.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_custom](resources--workload--reference--group-017.md#canonical-0203102001320123-1120130012130212-0111101220231333-3101331123203121-1110000311003232-1321231201202312-3200200021132112-0200303231330002)
- [stateful_service.advertise_options.advertise_custom.ports](resources--workload--reference--group-017.md#canonical-3111311322332303-1201030020021201-3133021233000313-2022333331131202-3221312032323033-3133120120113232-3330021120310331-2331303302032231)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-017.md#canonical-0313222022003221-2301103230203023-3120322210033100-2121220010332201-3133231021021220-1313001311100021-0313102021211212-0331322021202002)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-018.md#canonical-0112020200111022-0013103013103331-1110301001233101-1300202101312213-3301031210112102-1221232301300211-0231000232130131-3132113303133333)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-018.md#canonical-2020320003013033-2231100311132300-2132132133022132-3002120002132203-3131231201330331-0231012122032023-1202211010110230-0230013122122221)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates](resources--workload--reference--group-018.md#canonical-3332313302002330-1333101212032003-2203023132232213-0001033110210120-3121103210011323-2020220120012213-3220330021330022-1301300222033110)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates.private_key

<a id="canonical-1021102231312001-0330333033133100-1010222301102031-2011201131310331-3202312221013120-2131200330023123-1301313002120102-1030213103000222"></a>

Type: `"object"`. single nested block, Optional.

SecretType is used in an object to indicate a sensitive/confidential field.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("blindfold_secret_info",
    "clear_secret_info")}
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

- [blindfold_secret_info](resources--workload--reference--group-018.md#canonical-2103032132323102-3023011022021303-2201010033131212-2030111103320103-1120033321203121-3003213000323203-0323313000031023-2101102112002303): complete subsection reference.

- [clear_secret_info](resources--workload--reference--group-018.md#canonical-2331030102322003-1213200013022113-0001131012321023-1132102120312221-3220112212333213-0022231333213120-3102111000212031-3101013002021023): complete subsection reference.

<a id="canonical-2103032132323102-3023011022021303-2201010033131212-2030111103320103-1120033321203121-3003213000323203-0323313000031023-2101102112002303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates.private_key.blindfold_secret_info` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-017.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_custom](resources--workload--reference--group-017.md#canonical-0203102001320123-1120130012130212-0111101220231333-3101331123203121-1110000311003232-1321231201202312-3200200021132112-0200303231330002)
- [stateful_service.advertise_options.advertise_custom.ports](resources--workload--reference--group-017.md#canonical-3111311322332303-1201030020021201-3133021233000313-2022333331131202-3221312032323033-3133120120113232-3330021120310331-2331303302032231)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-017.md#canonical-0313222022003221-2301103230203023-3120322210033100-2121220010332201-3133231021021220-1313001311100021-0313102021211212-0331322021202002)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-018.md#canonical-0112020200111022-0013103013103331-1110301001233101-1300202101312213-3301031210112102-1221232301300211-0231000232130131-3132113303133333)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-018.md#canonical-2020320003013033-2231100311132300-2132132133022132-3002120002132203-3131231201330331-0231012122032023-1202211010110230-0230013122122221)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates](resources--workload--reference--group-018.md#canonical-3332313302002330-1333101212032003-2203023132232213-0001033110210120-3121103210011323-2020220120012213-3220330021330022-1301300222033110)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates.private_key](resources--workload--reference--group-018.md#canonical-3202322131133212-3233333112022031-2012323022120230-3213102233121101-2023030003200212-0221230103200223-2321231323033110-1310133003012200)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates.private_key.blindfold_secret_info

<a id="canonical-2223211110110010-3002020003233223-3022122101231222-0222203203103331-2031220131320011-2333131123323011-1103221300333131-0120121130332331"></a>

Type: `"object"`. single nested block, Optional.

BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("location")}
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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(4, 131072),
}
```

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
- [stateful_service](resources--workload--reference--group-017.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_custom](resources--workload--reference--group-017.md#canonical-0203102001320123-1120130012130212-0111101220231333-3101331123203121-1110000311003232-1321231201202312-3200200021132112-0200303231330002)
- [stateful_service.advertise_options.advertise_custom.ports](resources--workload--reference--group-017.md#canonical-3111311322332303-1201030020021201-3133021233000313-2022333331131202-3221312032323033-3133120120113232-3330021120310331-2331303302032231)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-017.md#canonical-0313222022003221-2301103230203023-3120322210033100-2121220010332201-3133231021021220-1313001311100021-0313102021211212-0331322021202002)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-018.md#canonical-0112020200111022-0013103013103331-1110301001233101-1300202101312213-3301031210112102-1221232301300211-0231000232130131-3132113303133333)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-018.md#canonical-2020320003013033-2231100311132300-2132132133022132-3002120002132203-3131231201330331-0231012122032023-1202211010110230-0230013122122221)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates](resources--workload--reference--group-018.md#canonical-3332313302002330-1333101212032003-2203023132232213-0001033110210120-3121103210011323-2020220120012213-3220330021330022-1301300222033110)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates.private_key](resources--workload--reference--group-018.md#canonical-3202322131133212-3233333112022031-2012323022120230-3213102233121101-2023030003200212-0221230103200223-2321231323033110-1310133003012200)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates.private_key.clear_secret_info

<a id="canonical-3211312332132010-1230022330210303-3132011203230101-0122001123300233-1323203102301023-2023023112001021-3200002131133313-3223303121120301"></a>

Type: `"object"`. single nested block, Optional.

ClearSecretInfoType specifies information about the Secret that is not encrypted.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("url")}
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

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 131072),
}
```

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
- [stateful_service](resources--workload--reference--group-017.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_custom](resources--workload--reference--group-017.md#canonical-0203102001320123-1120130012130212-0111101220231333-3101331123203121-1110000311003232-1321231201202312-3200200021132112-0200303231330002)
- [stateful_service.advertise_options.advertise_custom.ports](resources--workload--reference--group-017.md#canonical-3111311322332303-1201030020021201-3133021233000313-2022333331131202-3221312032323033-3133120120113232-3330021120310331-2331303302032231)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-017.md#canonical-0313222022003221-2301103230203023-3120322210033100-2121220010332201-3133231021021220-1313001311100021-0313102021211212-0331322021202002)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-018.md#canonical-0112020200111022-0013103013103331-1110301001233101-1300202101312213-3301031210112102-1221232301300211-0231000232130131-3132113303133333)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-018.md#canonical-2020320003013033-2231100311132300-2132132133022132-3002120002132203-3131231201330331-0231012122032023-1202211010110230-0230013122122221)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates](resources--workload--reference--group-018.md#canonical-3332313302002330-1333101212032003-2203023132232213-0001033110210120-3121103210011323-2020220120012213-3220330021330022-1301300222033110)
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
- [stateful_service](resources--workload--reference--group-017.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_custom](resources--workload--reference--group-017.md#canonical-0203102001320123-1120130012130212-0111101220231333-3101331123203121-1110000311003232-1321231201202312-3200200021132112-0200303231330002)
- [stateful_service.advertise_options.advertise_custom.ports](resources--workload--reference--group-017.md#canonical-3111311322332303-1201030020021201-3133021233000313-2022333331131202-3221312032323033-3133120120113232-3330021120310331-2331303302032231)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-017.md#canonical-0313222022003221-2301103230203023-3120322210033100-2121220010332201-3133231021021220-1313001311100021-0313102021211212-0331322021202002)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-018.md#canonical-0112020200111022-0013103013103331-1110301001233101-1300202101312213-3301031210112102-1221232301300211-0231000232130131-3132113303133333)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-018.md#canonical-2020320003013033-2231100311132300-2132132133022132-3002120002132203-3131231201330331-0231012122032023-1202211010110230-0230013122122221)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_config

<a id="canonical-1302313322112233-2103110020213110-3111100001201221-1102212003220303-3110220222030231-0232211313200331-3322003300321221-1133012303222032"></a>

Type: `"object"`. single nested block, Optional.

This defines various OPTIONS to configure TLS configuration parameters.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

<a id="canonical-1311212113300322-0102002323021032-0203301133130300-0030302223113231-2101320100111213-1210030130021003-2321112222233022-0321301330221212"></a>

### Direct properties for `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_config`

- [custom_security](resources--workload--reference--group-018.md#canonical-0000213222210123-2112221000103122-2120010312210201-0232033020303101-3231232002331120-3333003012123320-2011132231010301-2031302332130331): complete subsection reference.

- [default_security](resources--workload--reference--group-018.md#canonical-2223003033100203-3000223312312003-1133233232230033-0200031323221222-0300203200112303-2102132021102133-2121330003013210-3332001310120231): complete subsection reference.

- [low_security](resources--workload--reference--group-019.md#canonical-1202200110030311-0313232213230033-3120322331330020-3321301300011032-3302110110102213-1220303321101321-0203103112022121-3122201020113110): complete subsection reference.

- [medium_security](resources--workload--reference--group-019.md#canonical-1122130230131030-2322031210223132-1123333221000002-2112312320202000-1231130101112221-0213120331320322-2231010023221002-3103321103313021): complete subsection reference.

<a id="canonical-0000213222210123-2112221000103122-2120010312210201-0232033020303101-3231232002331120-3333003012123320-2011132231010301-2031302332130331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_config.custom_security` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-017.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_custom](resources--workload--reference--group-017.md#canonical-0203102001320123-1120130012130212-0111101220231333-3101331123203121-1110000311003232-1321231201202312-3200200021132112-0200303231330002)
- [stateful_service.advertise_options.advertise_custom.ports](resources--workload--reference--group-017.md#canonical-3111311322332303-1201030020021201-3133021233000313-2022333331131202-3221312032323033-3133120120113232-3330021120310331-2331303302032231)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-017.md#canonical-0313222022003221-2301103230203023-3120322210033100-2121220010332201-3133231021021220-1313001311100021-0313102021211212-0331322021202002)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-018.md#canonical-0112020200111022-0013103013103331-1110301001233101-1300202101312213-3301031210112102-1221232301300211-0231000232130131-3132113303133333)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-018.md#canonical-2020320003013033-2231100311132300-2132132133022132-3002120002132203-3131231201330331-0231012122032023-1202211010110230-0230013122122221)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_config](resources--workload--reference--group-018.md#canonical-3212132103211102-3103133011133021-3313221110320101-2111103132233113-1232201321022230-0232111013133132-1231122000001132-3012311030121223)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_config.custom_security

<a id="canonical-2330013133222003-1310131020312131-3033002322313122-3303123000333313-1301332032230331-2033123321230202-3121323310132212-3012203222320002"></a>

Type: `"object"`. single nested block, Optional.

This defines TLS protocol config including min/max versions and allowed ciphers.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

<a id="canonical-0002020032133201-1001120213201020-1320212030102130-3030221123222010-3120102320133122-0311221011122113-1103021102031132-2111003232032133"></a>

### Direct properties for `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_config.custom_security`

<a id="canonical-1331331230031221-1212333010101010-1003331301322130-3101323101001232-0332221003022121-3223001100033101-2311033222213110-0302030210000211"></a>

#### `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_config.custom_security.cipher_suites` property

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

<a id="canonical-0311022112320202-3021112002111211-0231330031223123-2302210201023100-0101310230130030-1333221002111121-3003002012303011-0001311322212032"></a>

<a id="canonical-3023103022103330-2231020221321122-0303201332223020-1312213121120111-0331013333213323-3101213330131202-3030202313231212-3330121113202031"></a>

#### `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_config.custom_security.max_version` property

Type: `"string"`. Optional.

\[Enum: TLS\_AUTO|TLSv1\_0|TLSv1\_1|TLSv1\_2|TLSv1\_3\] TlsProtocol is enumeration of supported TLS
versions F5 Distributed Cloud will choose the optimal TLS version. Possible values are
\`TLS\_AUTO\`, \`TLSv1\_0\`, \`TLSv1\_1\`, \`TLSv1\_2\`, \`TLSv1\_3\`. Defaults to \`TLS\_AUTO\`.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["TLS_AUTO","TLSv1_0","TLSv1_1","TLSv1_2","TLSv1_3"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
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

<a id="canonical-3311130020020310-2131301000231321-3110011202110111-2333220201010202-0011122231120332-2032013131312302-2021122220131031-1102300212022011"></a>

<a id="canonical-1323002101132032-1203331110321022-1232113120222310-2233213122132130-1120023102301223-0333313011023221-2330010032312333-1003201322002100"></a>

#### `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_config.custom_security.min_version` property

Type: `"string"`. Optional.

\[Enum: TLS\_AUTO|TLSv1\_0|TLSv1\_1|TLSv1\_2|TLSv1\_3\] TlsProtocol is enumeration of supported TLS
versions F5 Distributed Cloud will choose the optimal TLS version. Possible values are
\`TLS\_AUTO\`, \`TLSv1\_0\`, \`TLSv1\_1\`, \`TLSv1\_2\`, \`TLSv1\_3\`. Defaults to \`TLS\_AUTO\`.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["TLS_AUTO","TLSv1_0","TLSv1_1","TLSv1_2","TLSv1_3"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
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

<a id="canonical-2223003033100203-3000223312312003-1133233232230033-0200031323221222-0300203200112303-2102132021102133-2121330003013210-3332001310120231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_config.default_security` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-017.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_custom](resources--workload--reference--group-017.md#canonical-0203102001320123-1120130012130212-0111101220231333-3101331123203121-1110000311003232-1321231201202312-3200200021132112-0200303231330002)
- [stateful_service.advertise_options.advertise_custom.ports](resources--workload--reference--group-017.md#canonical-3111311322332303-1201030020021201-3133021233000313-2022333331131202-3221312032323033-3133120120113232-3330021120310331-2331303302032231)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-017.md#canonical-0313222022003221-2301103230203023-3120322210033100-2121220010332201-3133231021021220-1313001311100021-0313102021211212-0331322021202002)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-018.md#canonical-0112020200111022-0013103013103331-1110301001233101-1300202101312213-3301031210112102-1221232301300211-0231000232130131-3132113303133333)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-018.md#canonical-2020320003013033-2231100311132300-2132132133022132-3002120002132203-3131231201330331-0231012122032023-1202211010110230-0230013122122221)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_config](resources--workload--reference--group-018.md#canonical-3212132103211102-3103133011133021-3313221110320101-2111103132233113-1232201321022230-0232111013133132-1231122000001132-3012311030121223)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_config.default_security

<a id="canonical-1231213210333320-0132002130001221-1010211322123202-0122022311000320-2133303021303101-0001211031200122-3002333323030303-1133233200011323"></a>

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
