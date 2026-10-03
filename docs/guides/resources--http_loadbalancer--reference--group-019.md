---
page_title: "xcsh_http_loadbalancer reference"
subcategory: "Load Balancing"
description: "Complete grouped canonical reference for xcsh_http_loadbalancer reference."
---

# xcsh_http_loadbalancer reference

<a id="canonical-3223001213213131-2023230002332210-1130133201232212-3333023322332202-0202232321220020-0122301012033210-0211231001312122-2333001231332001"></a>

## port_ranges property — http / 302213323300 / 6

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

<a id="canonical-0332001113302030-1003213113121022-3310013223333311-0310103201331010-3230322321003220-3210103022001321-3321202122212010-2122111111120313"></a>

## Next pages — http / 302213323300 / 7

- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-0321232101200210-0011302000022301-2002321223210110-0303132310312201-0222311023330031-3231011200231021-2012211321122303-0001332120000102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3301011010210201-2311221033333333-0210103131302100-1011003030320220-0102301331310131-1222220220330321-3001111333100331-2321110313322022"></a>

## https — https / 333233133300 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- https

<a id="canonical-1300211231013103-1001303130221323-0000301002110021-2210133210111010-1323100310232132-0101111330112313-3100110313000100-1120103110300213"></a>

Type: `"object"`. single nested block, Optional.

Choice for selecting HTTP proxy with bring your own certificates. Changing this type selection
requires recreation and may interrupt service. Supported settings within the same selected type
remain updatable.

Upstream description:

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

<a id="canonical-2320200110221002-2012210031021230-1123213030020310-1201210131030230-3300120300003001-2023100221323220-2302331120020202-2020113202232000"></a>

## Direct properties — https / 333233133300 / 3

<a id="canonical-2100112122110321-2010303103220203-3232301112203331-3002220221212210-1211113223133110-0222230200022231-0100321033103332-3100303201111321"></a>

<a id="canonical-1210100020203220-2101013310331110-0222333233330000-1102112120120032-1230002232330313-2222123133200202-1002113031020312-0020100101220210"></a>

## add_hsts property — https / 333233133300 / 4

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

<a id="canonical-0223011010203002-1202313223300011-2101003030311003-3232233332010130-2020301123123231-2022113312123232-0102312220302133-3000331033003222"></a>

<a id="canonical-1310330333221301-0332220011221100-2031133121001012-3312000023110012-3033211011013030-2231012312323312-3003320001010210-0210333002032032"></a>

## append_server_name property — https / 333233133300 / 5

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

- [coalescing_options](resources--http_loadbalancer--reference--group-019.md#canonical-1002313102323120-3123301033030000-3103101001113132-3133202100213312-1322100132122311-0323121030233012-0112120332223120-3311030210023130): complete subsection reference.

<a id="canonical-3330133120100102-3202120011033200-1133033320331101-0101113002131000-3321120023011002-3003203320301323-0003132233122221-1312300010022102"></a>

<a id="canonical-1300300210300231-1123312313101201-0213302330310022-2312330320110211-0330011033333331-2101123220121121-2333001130320020-2332320130203102"></a>

## connection_idle_timeout property — https / 333233133300 / 6

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

- [default_header](resources--http_loadbalancer--reference--group-019.md#canonical-0001322102003113-3211120130133321-0221301230112220-0231230323033033-1020122212100120-0231201121003012-2223220021021011-0312001323312113): complete subsection reference.

- [default_loadbalancer](resources--http_loadbalancer--reference--group-019.md#canonical-1202321302320013-2012321213132330-1020322133333023-2033121332202310-2201003010201233-0330000122003010-2123113120113011-3332022131113310): complete subsection reference.

- [disable_path_normalize](resources--http_loadbalancer--reference--group-019.md#canonical-3031231302133033-2202202232302122-1011101311023032-1222210110212223-0203220300212131-1233112202010222-0321013031023321-3233303101212302): complete subsection reference.

- [enable_path_normalize](resources--http_loadbalancer--reference--group-019.md#canonical-3310310330000320-0012333332303330-0013011232220321-3101220020113220-1031033213202333-2022111013202310-1321200233213200-1132333020110331): complete subsection reference.

- [http_protocol_options](resources--http_loadbalancer--reference--group-019.md#canonical-1231312213210133-3133021101233212-3301313120010020-3113001231200211-0223302012033313-3000301122212000-3131320032301320-1211302233030222): complete subsection reference.

<a id="canonical-1122213102031302-3000123320103002-2220321030122221-3300030332012212-2200311010020010-1203313121332311-0313022120033201-2022130310223313"></a>

<a id="canonical-0003203302200310-3112131211002331-2033303323132322-1132121323130333-2100012010222333-3010232121133220-1212221000101113-2100002233033330"></a>

## http_redirect property — https / 333233133300 / 7

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

- [non_default_loadbalancer](resources--http_loadbalancer--reference--group-019.md#canonical-3122121110210020-0112331300201220-2103232132122022-2101302110211121-3013232333012111-3310330202000103-2001120322233003-2030011312230303): complete subsection reference.

- [pass_through](resources--http_loadbalancer--reference--group-019.md#canonical-1320021323133122-2131132203230011-0222102300201203-2133211311213200-0330211132231220-3302123120111202-3201010230103132-0121030000310101): complete subsection reference.

<a id="canonical-1302033032112013-3322330303202112-3313300300310301-3003312303320013-3303031300311210-2200000003003103-1302300120211023-2200011200123212"></a>

<a id="canonical-2210032012020322-2033301002231301-2203200033311233-1321312331010021-2220121133120030-2231120232223212-0011120230122031-1210032233332003"></a>

## port property — https / 333233133300 / 8

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

<a id="canonical-3231112133022202-3300033010201033-2020210010200303-1202020102220022-3303313110232011-0213100130310301-2131103333203233-1202201220000210"></a>

<a id="canonical-1322213121331332-1301122021101301-0022100022233203-2000130011213022-1201101313233302-0322231013333013-3112313111302202-0103132223121001"></a>

## port_ranges property — https / 333233133300 / 9

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

<a id="canonical-3202003303131203-0030300311313211-2013303221213131-1201101012311131-2210131023202330-1332123303023312-3133302033232131-2023331212101320"></a>

<a id="canonical-0200313103232301-2131331110212322-3003223332223112-0133133103222030-3032303131313111-0202011300021300-2030302210001300-3220300331020222"></a>

## server_name property — https / 333233133300 / 10

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

- [tls_cert_params](resources--http_loadbalancer--reference--group-019.md#canonical-0301302212033322-0103133010132330-1223032330030010-2130122212120320-0130211111203213-1023220310223110-3201023013223032-1132311122210133): complete subsection reference.

- [tls_parameters](resources--http_loadbalancer--reference--group-019.md#canonical-1231113013303201-3233222000300212-2211131130013302-2322211003010221-0320233103301130-2102232111212102-2321212333113100-3130203302022212): complete subsection reference.

<a id="canonical-3321230321212101-3023333312301221-0210112022101320-1023013300332212-2202013233230111-0132112313222300-3121233101022312-0103133023030333"></a>

## Next pages — https / 333233133300 / 11

- [https.coalescing_options](resources--http_loadbalancer--reference--group-019.md#canonical-1002313102323120-3123301033030000-3103101001113132-3133202100213312-1322100132122311-0323121030233012-0112120332223120-3311030210023130)
- [https.default_header](resources--http_loadbalancer--reference--group-019.md#canonical-0001322102003113-3211120130133321-0221301230112220-0231230323033033-1020122212100120-0231201121003012-2223220021021011-0312001323312113)
- [https.default_loadbalancer](resources--http_loadbalancer--reference--group-019.md#canonical-1202321302320013-2012321213132330-1020322133333023-2033121332202310-2201003010201233-0330000122003010-2123113120113011-3332022131113310)
- [https.disable_path_normalize](resources--http_loadbalancer--reference--group-019.md#canonical-3031231302133033-2202202232302122-1011101311023032-1222210110212223-0203220300212131-1233112202010222-0321013031023321-3233303101212302)
- [https.enable_path_normalize](resources--http_loadbalancer--reference--group-019.md#canonical-3310310330000320-0012333332303330-0013011232220321-3101220020113220-1031033213202333-2022111013202310-1321200233213200-1132333020110331)
- [https.http_protocol_options](resources--http_loadbalancer--reference--group-019.md#canonical-1231312213210133-3133021101233212-3301313120010020-3113001231200211-0223302012033313-3000301122212000-3131320032301320-1211302233030222)
- [https.non_default_loadbalancer](resources--http_loadbalancer--reference--group-019.md#canonical-3122121110210020-0112331300201220-2103232132122022-2101302110211121-3013232333012111-3310330202000103-2001120322233003-2030011312230303)
- [https.pass_through](resources--http_loadbalancer--reference--group-019.md#canonical-1320021323133122-2131132203230011-0222102300201203-2133211311213200-0330211132231220-3302123120111202-3201010230103132-0121030000310101)
- [https.tls_cert_params](resources--http_loadbalancer--reference--group-019.md#canonical-0301302212033322-0103133010132330-1223032330030010-2130122212120320-0130211111203213-1023220310223110-3201023013223032-1132311122210133)
- [https.tls_parameters](resources--http_loadbalancer--reference--group-019.md#canonical-1231113013303201-3233222000300212-2211131130013302-2322211003010221-0320233103301130-2102232111212102-2321212333113100-3130203302022212)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-1002313102323120-3123301033030000-3103101001113132-3133202100213312-1322100132122311-0323121030233012-0112120332223120-3311030210023130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0303130310112310-0301312332231220-3232233113110213-2210012210233323-0010022233332331-3111233112113323-1210100203100222-2211312231130301"></a>

## https.coalescing_options — coalescing_options / 312101202011 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [https](resources--http_loadbalancer--reference--group-019.md#canonical-0321232101200210-0011302000022301-2002321223210110-0303132310312201-0222311023330031-3231011200231021-2012211321122303-0001332120000102)
- https.coalescing_options

<a id="canonical-1122033020202321-1320331133312222-2022311310023123-2233110100233032-2021032102101311-0311223200032323-2031301023033002-0130312111330333"></a>

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

<a id="canonical-0033132133101330-2131311002300000-0202121300132332-2003001113200331-2320123132312200-2201132231220011-0131231222133201-1022232320322322"></a>

## Direct properties — coalescing_options / 312101202011 / 3

- [default_coalescing](resources--http_loadbalancer--reference--group-019.md#canonical-3320122023030300-1303303301013333-3300220311333112-0311030201220313-3212203012212033-3110331321232100-3102232103221123-1101122320230000): complete subsection reference.

- [strict_coalescing](resources--http_loadbalancer--reference--group-019.md#canonical-2303122331132210-0313232231022031-0320321031222022-3322112320123101-0222002321122103-2202230132200132-2201300330300120-3000003023202202): complete subsection reference.

<a id="canonical-3230122011112200-0023200213201031-3220110023231012-1113123120120122-0311323132120112-3113203102210023-2100330030110022-2021022203201331"></a>

## Next pages — coalescing_options / 312101202011 / 4

- [https.coalescing_options.default_coalescing](resources--http_loadbalancer--reference--group-019.md#canonical-3320122023030300-1303303301013333-3300220311333112-0311030201220313-3212203012212033-3110331321232100-3102232103221123-1101122320230000)
- [https.coalescing_options.strict_coalescing](resources--http_loadbalancer--reference--group-019.md#canonical-2303122331132210-0313232231022031-0320321031222022-3322112320123101-0222002321122103-2202230132200132-2201300330300120-3000003023202202)
- [https](resources--http_loadbalancer--reference--group-019.md#canonical-0321232101200210-0011302000022301-2002321223210110-0303132310312201-0222311023330031-3231011200231021-2012211321122303-0001332120000102)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-3320122023030300-1303303301013333-3300220311333112-0311030201220313-3212203012212033-3110331321232100-3102232103221123-1101122320230000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3010120100311230-3201310232131022-1330000112110122-3110011033032213-3122030230231202-0113103002213012-0101303101010133-2331132130012220"></a>

## https.coalescing_options.default_coalescing — default_coalescing / 320101011020 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [https](resources--http_loadbalancer--reference--group-019.md#canonical-0321232101200210-0011302000022301-2002321223210110-0303132310312201-0222311023330031-3231011200231021-2012211321122303-0001332120000102)
- [https.coalescing_options](resources--http_loadbalancer--reference--group-019.md#canonical-1002313102323120-3123301033030000-3103101001113132-3133202100213312-1322100132122311-0323121030233012-0112120332223120-3311030210023130)
- https.coalescing_options.default_coalescing

<a id="canonical-1030032320202301-1302021210011012-3321023313002201-1010010232012032-1213120023321213-0032121211132230-2213123203001321-2131233210200211"></a>

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

<a id="canonical-2133001031233122-2210231220002032-0102000222123002-2320302101110110-2301223013223302-1133321120003121-0320113321333233-1002031102013123"></a>

## Direct properties — default_coalescing / 320101011020 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2233020323023032-1322101233133310-0320102020001033-2123121330332333-0201021120133112-3122013303101210-2102132211013130-3122103232131130"></a>

## Next pages — default_coalescing / 320101011020 / 4

- [https.coalescing_options](resources--http_loadbalancer--reference--group-019.md#canonical-1002313102323120-3123301033030000-3103101001113132-3133202100213312-1322100132122311-0323121030233012-0112120332223120-3311030210023130)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-2303122331132210-0313232231022031-0320321031222022-3322112320123101-0222002321122103-2202230132200132-2201300330300120-3000003023202202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1111230032201230-1333203002113200-2233033200123202-1112303321200233-1123101001122320-1320221200212311-0100230232111003-1113303031123331"></a>

## https.coalescing_options.strict_coalescing — strict_coalescing / 011332103303 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [https](resources--http_loadbalancer--reference--group-019.md#canonical-0321232101200210-0011302000022301-2002321223210110-0303132310312201-0222311023330031-3231011200231021-2012211321122303-0001332120000102)
- [https.coalescing_options](resources--http_loadbalancer--reference--group-019.md#canonical-1002313102323120-3123301033030000-3103101001113132-3133202100213312-1322100132122311-0323121030233012-0112120332223120-3311030210023130)
- https.coalescing_options.strict_coalescing

<a id="canonical-0232130330011310-0030113300302203-1211302133103000-0200322003332102-3000210322321220-0231331213232333-1130230130012100-0021332103112130"></a>

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

<a id="canonical-0111211133012133-0013232122010003-2301322300011010-2323003012322230-3332300123333311-3222100232301300-0120221310200330-0130322033201301"></a>

## Direct properties — strict_coalescing / 011332103303 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2321110312132103-1202012303302000-3231230313130100-3031012300112130-3031000011002323-2023011023131333-0130202300302102-1300133300022100"></a>

## Next pages — strict_coalescing / 011332103303 / 4

- [https.coalescing_options](resources--http_loadbalancer--reference--group-019.md#canonical-1002313102323120-3123301033030000-3103101001113132-3133202100213312-1322100132122311-0323121030233012-0112120332223120-3311030210023130)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-0001322102003113-3211120130133321-0221301230112220-0231230323033033-1020122212100120-0231201121003012-2223220021021011-0312001323312113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3302213013111210-0023023310222321-3003320301330113-1002213313310112-1000232232121231-2032212221021021-2000232023231010-3021200030203133"></a>

## https.default_header — default_header / 001003233100 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [https](resources--http_loadbalancer--reference--group-019.md#canonical-0321232101200210-0011302000022301-2002321223210110-0303132310312201-0222311023330031-3231011200231021-2012211321122303-0001332120000102)
- https.default_header

<a id="canonical-3230323130312231-0102321210323322-1211022021212320-3332001332001221-2020001211132121-0300123221011312-1023130131203111-2202030020333100"></a>

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

<a id="canonical-2333210003310330-0211231212000321-2131000211131000-1032000332033233-2021020002020011-3033000003311223-0112133003331230-1211103222330332"></a>

## Direct properties — default_header / 001003233100 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3233213331210133-3303201000221133-1103322031031020-2220303322000313-2300101330031111-2131320203131001-1023103322202322-3100321223311021"></a>

## Next pages — default_header / 001003233100 / 4

- [https](resources--http_loadbalancer--reference--group-019.md#canonical-0321232101200210-0011302000022301-2002321223210110-0303132310312201-0222311023330031-3231011200231021-2012211321122303-0001332120000102)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-1202321302320013-2012321213132330-1020322133333023-2033121332202310-2201003010201233-0330000122003010-2123113120113011-3332022131113310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0313000212223221-0110111031213300-2013213211230330-3011332302233030-2133110330022302-3311211230201302-1103130212303213-0323003301013321"></a>

## https.default_loadbalancer — default_loadbalancer / 011013212321 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [https](resources--http_loadbalancer--reference--group-019.md#canonical-0321232101200210-0011302000022301-2002321223210110-0303132310312201-0222311023330031-3231011200231021-2012211321122303-0001332120000102)
- https.default_loadbalancer

<a id="canonical-1223202332122332-3331232123313010-1302023121201110-0301030031023213-1003213310022000-2222101333101102-1231313122130323-2321320032330220"></a>

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

<a id="canonical-2013220033332031-0310002110330333-3300331203302321-3032111003323302-2323131200120311-0230121233331102-0121321332031010-3113202123113022"></a>

## Direct properties — default_loadbalancer / 011013212321 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1232133332231013-1002311131200202-0023110233301030-1001231011223011-0013213300321201-1310123213211020-0330103322122210-2223013230022021"></a>

## Next pages — default_loadbalancer / 011013212321 / 4

- [https](resources--http_loadbalancer--reference--group-019.md#canonical-0321232101200210-0011302000022301-2002321223210110-0303132310312201-0222311023330031-3231011200231021-2012211321122303-0001332120000102)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-3031231302133033-2202202232302122-1011101311023032-1222210110212223-0203220300212131-1233112202010222-0321013031023321-3233303101212302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3321012103202302-3002211120200203-3102022023010213-1233121203211010-0021022231221101-0223011210000030-1020100203032011-2032232312133003"></a>

## https.disable_path_normalize — disable_path_normalize / 021231130230 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [https](resources--http_loadbalancer--reference--group-019.md#canonical-0321232101200210-0011302000022301-2002321223210110-0303132310312201-0222311023330031-3231011200231021-2012211321122303-0001332120000102)
- https.disable_path_normalize

<a id="canonical-1031230121302321-1111221210301321-2000330132301000-3223120331030321-1222001201331133-2011021030120230-1300312232230230-1022100223320023"></a>

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

<a id="canonical-1320002021230202-3320033213132120-2010000003210012-0000131233131010-1113021011113232-0210111310330032-0232100330021132-0030031110212012"></a>

## Direct properties — disable_path_normalize / 021231130230 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1322000112133010-1120332232322330-3300101331100123-2001312321211030-2221313321231222-3131230101233332-0023033221223320-3100210221202301"></a>

## Next pages — disable_path_normalize / 021231130230 / 4

- [https](resources--http_loadbalancer--reference--group-019.md#canonical-0321232101200210-0011302000022301-2002321223210110-0303132310312201-0222311023330031-3231011200231021-2012211321122303-0001332120000102)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-3310310330000320-0012333332303330-0013011232220321-3101220020113220-1031033213202333-2022111013202310-1321200233213200-1132333020110331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2202313112233220-1201222123032212-0012212321321012-1011101030100131-1003002203220010-2023323320022133-3123313121020000-3200020212333121"></a>

## https.enable_path_normalize — enable_path_normalize / 101101320111 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [https](resources--http_loadbalancer--reference--group-019.md#canonical-0321232101200210-0011302000022301-2002321223210110-0303132310312201-0222311023330031-3231011200231021-2012211321122303-0001332120000102)
- https.enable_path_normalize

<a id="canonical-3100122030321033-1011002222030332-2102213021201023-1033232022211120-1202231030330103-1010200200103322-2213201221032331-3310130202302031"></a>

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

<a id="canonical-3101012200303123-2020113312313201-1031012110230322-3121223030323131-0113302112030301-3120023211331213-3122131310201320-3112332133312001"></a>

## Direct properties — enable_path_normalize / 101101320111 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0011022110120100-0122313111110222-3022123021103203-0013033111331013-0310322322031111-2023200210001223-3102222330331310-0210033301223200"></a>

## Next pages — enable_path_normalize / 101101320111 / 4

- [https](resources--http_loadbalancer--reference--group-019.md#canonical-0321232101200210-0011302000022301-2002321223210110-0303132310312201-0222311023330031-3231011200231021-2012211321122303-0001332120000102)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-1231312213210133-3133021101233212-3301313120010020-3113001231200211-0223302012033313-3000301122212000-3131320032301320-1211302233030222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0323113220021300-3232330032100120-1212201030311233-0003022300302012-3121333212300030-0131302023320303-0032330200121322-3000320201013321"></a>

## https.http_protocol_options — http_protocol_options / 313313323310 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [https](resources--http_loadbalancer--reference--group-019.md#canonical-0321232101200210-0011302000022301-2002321223210110-0303132310312201-0222311023330031-3231011200231021-2012211321122303-0001332120000102)
- https.http_protocol_options

<a id="canonical-0220313022230110-3030031100202020-0210331332011031-3110320012030313-3033131313020133-0113330202331320-2023131312002230-1201202131331232"></a>

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

<a id="canonical-0232130131301011-1310032303201010-0011322222311033-3312133130201311-0013130031022323-0232233101113130-3233003111123031-2321200033111103"></a>

## Direct properties — http_protocol_options / 313313323310 / 3

- [http_protocol_enable_v1_only](resources--http_loadbalancer--reference--group-019.md#canonical-2133101003321303-3101213211122210-2233132021020222-3201312321020100-0102323302003303-2213010100210213-0212031113232031-0221300110230303): complete subsection reference.

- [http_protocol_enable_v1_v2](resources--http_loadbalancer--reference--group-019.md#canonical-0010313222101213-2030222023120120-3033111100332312-3000213103211301-1002213101332020-1012332212021033-2333200203003133-1002230333101111): complete subsection reference.

- [http_protocol_enable_v2_only](resources--http_loadbalancer--reference--group-019.md#canonical-0213320000020033-2120001112320123-2330332010013101-3101232320223210-0223322031210210-3112031102302331-0322121121003330-1000320310232102): complete subsection reference.

<a id="canonical-1320021020321103-2310200231233300-0030201220331100-1021020002222321-3031232332012201-0131131221122101-3012000103221322-2303020321100330"></a>

## Next pages — http_protocol_options / 313313323310 / 4

- [https.http_protocol_options.http_protocol_enable_v1_only](resources--http_loadbalancer--reference--group-019.md#canonical-2133101003321303-3101213211122210-2233132021020222-3201312321020100-0102323302003303-2213010100210213-0212031113232031-0221300110230303)
- [https.http_protocol_options.http_protocol_enable_v1_v2](resources--http_loadbalancer--reference--group-019.md#canonical-0010313222101213-2030222023120120-3033111100332312-3000213103211301-1002213101332020-1012332212021033-2333200203003133-1002230333101111)
- [https.http_protocol_options.http_protocol_enable_v2_only](resources--http_loadbalancer--reference--group-019.md#canonical-0213320000020033-2120001112320123-2330332010013101-3101232320223210-0223322031210210-3112031102302331-0322121121003330-1000320310232102)
- [https](resources--http_loadbalancer--reference--group-019.md#canonical-0321232101200210-0011302000022301-2002321223210110-0303132310312201-0222311023330031-3231011200231021-2012211321122303-0001332120000102)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-2133101003321303-3101213211122210-2233132021020222-3201312321020100-0102323302003303-2213010100210213-0212031113232031-0221300110230303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3231001221001003-1012302231210010-0222303210010002-1303000103323022-1131122030012031-3301302233133020-1103230133231120-0202132201203100"></a>

## https.http_protocol_options.http_protocol_enable_v1_only — http_protocol_enable_v1_only / 030233020133 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [https](resources--http_loadbalancer--reference--group-019.md#canonical-0321232101200210-0011302000022301-2002321223210110-0303132310312201-0222311023330031-3231011200231021-2012211321122303-0001332120000102)
- [https.http_protocol_options](resources--http_loadbalancer--reference--group-019.md#canonical-1231312213210133-3133021101233212-3301313120010020-3113001231200211-0223302012033313-3000301122212000-3131320032301320-1211302233030222)
- https.http_protocol_options.http_protocol_enable_v1_only

<a id="canonical-1021122103311221-0102212302120022-3033113002220000-1131110313333032-3200201221221333-3033000300212330-0023203130202121-2121012200222320"></a>

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

<a id="canonical-0011111111132101-0131331023130000-3110201301201021-0220112021100202-1120003133133322-3023011010321222-0121000231001001-1333321200231233"></a>

## Direct properties — http_protocol_enable_v1_only / 030233020133 / 3

- [header_transformation](resources--http_loadbalancer--reference--group-019.md#canonical-3323230131311000-3101333122133201-3231223220131130-0130233100111121-0200300231200030-0320231233333003-1130030000311312-3321123212131300): complete subsection reference.

<a id="canonical-2213310210220301-1012003302013122-2323102100000023-2213221103012112-1020222112330230-0113023233121200-1302100103030110-2310003002210320"></a>

## Next pages — http_protocol_enable_v1_only / 030233020133 / 4

- [https.http_protocol_options.http_protocol_enable_v1_only.header_transformation](resources--http_loadbalancer--reference--group-019.md#canonical-3323230131311000-3101333122133201-3231223220131130-0130233100111121-0200300231200030-0320231233333003-1130030000311312-3321123212131300)
- [https.http_protocol_options](resources--http_loadbalancer--reference--group-019.md#canonical-1231312213210133-3133021101233212-3301313120010020-3113001231200211-0223302012033313-3000301122212000-3131320032301320-1211302233030222)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-3323230131311000-3101333122133201-3231223220131130-0130233100111121-0200300231200030-0320231233333003-1130030000311312-3321123212131300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1030000001231131-1113000002213120-1320230101120302-2303120012031123-2102021110213032-0033121101013113-3131331102330122-3230201331111330"></a>

## https.http_protocol_options.http_protocol_enable_v1_only.header_transformation — header_transformation / 203312032220 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [https](resources--http_loadbalancer--reference--group-019.md#canonical-0321232101200210-0011302000022301-2002321223210110-0303132310312201-0222311023330031-3231011200231021-2012211321122303-0001332120000102)
- [https.http_protocol_options](resources--http_loadbalancer--reference--group-019.md#canonical-1231312213210133-3133021101233212-3301313120010020-3113001231200211-0223302012033313-3000301122212000-3131320032301320-1211302233030222)
- [https.http_protocol_options.http_protocol_enable_v1_only](resources--http_loadbalancer--reference--group-019.md#canonical-2133101003321303-3101213211122210-2233132021020222-3201312321020100-0102323302003303-2213010100210213-0212031113232031-0221300110230303)
- https.http_protocol_options.http_protocol_enable_v1_only.header_transformation

<a id="canonical-3011111120212220-0130201303013010-1120233303211003-2312201013201301-1101310211112331-0103213132212132-1102111023303323-2200013220312000"></a>

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

<a id="canonical-1011332000102023-0222001000003012-0002202010032331-1321232103010021-2221321320231003-0021310231030223-2113111021121333-3132021022030011"></a>

## Direct properties — header_transformation / 203312032220 / 3

- [default_header_transformation](resources--http_loadbalancer--reference--group-019.md#canonical-1323330232110033-0211122003202001-1332332230121232-0323200030122223-0112103231321211-2233120020200130-3120021013232110-0310201230230001): complete subsection reference.

- [preserve_case_header_transformation](resources--http_loadbalancer--reference--group-019.md#canonical-1231233310223202-3233312011132312-2210010203222120-3313321120312201-2101001212301223-3032231122301021-2202120302010301-3223220200330223): complete subsection reference.

- [proper_case_header_transformation](resources--http_loadbalancer--reference--group-019.md#canonical-0000200030102031-0302100111210302-2323113211322313-2133230131002012-0312030000211033-1300200221332213-3313210133020021-3201321213311331): complete subsection reference.

<a id="canonical-2230301020031233-2111301113121030-2223303033233002-3322310022210332-1212212332312102-3312203103131132-0013032012311113-2320330321231010"></a>

## Next pages — header_transformation / 203312032220 / 4

- [https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.default_header_transformation](resources--http_loadbalancer--reference--group-019.md#canonical-1323330232110033-0211122003202001-1332332230121232-0323200030122223-0112103231321211-2233120020200130-3120021013232110-0310201230230001)
- [https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.preserve_case_header_transformation](resources--http_loadbalancer--reference--group-019.md#canonical-1231233310223202-3233312011132312-2210010203222120-3313321120312201-2101001212301223-3032231122301021-2202120302010301-3223220200330223)
- [https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.proper_case_header_transformation](resources--http_loadbalancer--reference--group-019.md#canonical-0000200030102031-0302100111210302-2323113211322313-2133230131002012-0312030000211033-1300200221332213-3313210133020021-3201321213311331)
- [https.http_protocol_options.http_protocol_enable_v1_only](resources--http_loadbalancer--reference--group-019.md#canonical-2133101003321303-3101213211122210-2233132021020222-3201312321020100-0102323302003303-2213010100210213-0212031113232031-0221300110230303)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-1323330232110033-0211122003202001-1332332230121232-0323200030122223-0112103231321211-2233120020200130-3120021013232110-0310201230230001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0200033012000331-3111032131332211-2011000020201330-3211221211121313-2101113122000122-2311103312102111-0211101011211231-1010210020300223"></a>

## https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.default_header_transformation — default_header_transformation / 211213031023 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [https](resources--http_loadbalancer--reference--group-019.md#canonical-0321232101200210-0011302000022301-2002321223210110-0303132310312201-0222311023330031-3231011200231021-2012211321122303-0001332120000102)
- [https.http_protocol_options](resources--http_loadbalancer--reference--group-019.md#canonical-1231312213210133-3133021101233212-3301313120010020-3113001231200211-0223302012033313-3000301122212000-3131320032301320-1211302233030222)
- [https.http_protocol_options.http_protocol_enable_v1_only](resources--http_loadbalancer--reference--group-019.md#canonical-2133101003321303-3101213211122210-2233132021020222-3201312321020100-0102323302003303-2213010100210213-0212031113232031-0221300110230303)
- [https.http_protocol_options.http_protocol_enable_v1_only.header_transformation](resources--http_loadbalancer--reference--group-019.md#canonical-3323230131311000-3101333122133201-3231223220131130-0130233100111121-0200300231200030-0320231233333003-1130030000311312-3321123212131300)
- https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.default_header_transformation

<a id="canonical-3022103032321211-0101203203202213-1212322322013033-0311313103002331-3322100203131032-1201302203002202-0111012332331120-2233233302033312"></a>

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

<a id="canonical-0031333202033001-3112220333113022-3112122230100013-2102203021012131-1100310132203323-2211122220110230-2112230120102020-3213023221311333"></a>

## Direct properties — default_header_transformation / 211213031023 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3000123013103220-3003103023131023-0230213230233222-2120120323100303-1213221031130001-0220030231231302-1322332030311213-1321113110010031"></a>

## Next pages — default_header_transformation / 211213031023 / 4

- [https.http_protocol_options.http_protocol_enable_v1_only.header_transformation](resources--http_loadbalancer--reference--group-019.md#canonical-3323230131311000-3101333122133201-3231223220131130-0130233100111121-0200300231200030-0320231233333003-1130030000311312-3321123212131300)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-1231233310223202-3233312011132312-2210010203222120-3313321120312201-2101001212301223-3032231122301021-2202120302010301-3223220200330223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2030311211023301-1021100003212311-1122300112132131-2332302103112030-1002222201112203-3102000001202230-3231013231010030-3220201132101033"></a>

## https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.preserve_case_header_transformation — preserve_case_header_transformation / 001022102122 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [https](resources--http_loadbalancer--reference--group-019.md#canonical-0321232101200210-0011302000022301-2002321223210110-0303132310312201-0222311023330031-3231011200231021-2012211321122303-0001332120000102)
- [https.http_protocol_options](resources--http_loadbalancer--reference--group-019.md#canonical-1231312213210133-3133021101233212-3301313120010020-3113001231200211-0223302012033313-3000301122212000-3131320032301320-1211302233030222)
- [https.http_protocol_options.http_protocol_enable_v1_only](resources--http_loadbalancer--reference--group-019.md#canonical-2133101003321303-3101213211122210-2233132021020222-3201312321020100-0102323302003303-2213010100210213-0212031113232031-0221300110230303)
- [https.http_protocol_options.http_protocol_enable_v1_only.header_transformation](resources--http_loadbalancer--reference--group-019.md#canonical-3323230131311000-3101333122133201-3231223220131130-0130233100111121-0200300231200030-0320231233333003-1130030000311312-3321123212131300)
- https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.preserve_case_header_transformation

<a id="canonical-2002111002031003-1221032101222103-0011310033300301-0331121222133213-0311100013120003-2023311303203001-2101203020310203-1002202210203132"></a>

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

<a id="canonical-1121103330300302-3022201113332023-3302212032133321-1123203030230313-2032111301103132-2012120130131222-0031201323200012-3002133233110110"></a>

## Direct properties — preserve_case_header_transformation / 001022102122 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3130210111032013-1101131221323103-0011223233223221-2322333013222003-3312212222220301-2112332332233013-3310110221110033-0122213201001321"></a>

## Next pages — preserve_case_header_transformation / 001022102122 / 4

- [https.http_protocol_options.http_protocol_enable_v1_only.header_transformation](resources--http_loadbalancer--reference--group-019.md#canonical-3323230131311000-3101333122133201-3231223220131130-0130233100111121-0200300231200030-0320231233333003-1130030000311312-3321123212131300)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-0000200030102031-0302100111210302-2323113211322313-2133230131002012-0312030000211033-1300200221332213-3313210133020021-3201321213311331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0023102310231331-1010110132310233-1013313303300310-0011312220121232-2131003032330111-3321123332212002-2112111100100313-2223002233302303"></a>

## https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.proper_case_header_transformation — proper_case_header_transformation / 312322023203 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [https](resources--http_loadbalancer--reference--group-019.md#canonical-0321232101200210-0011302000022301-2002321223210110-0303132310312201-0222311023330031-3231011200231021-2012211321122303-0001332120000102)
- [https.http_protocol_options](resources--http_loadbalancer--reference--group-019.md#canonical-1231312213210133-3133021101233212-3301313120010020-3113001231200211-0223302012033313-3000301122212000-3131320032301320-1211302233030222)
- [https.http_protocol_options.http_protocol_enable_v1_only](resources--http_loadbalancer--reference--group-019.md#canonical-2133101003321303-3101213211122210-2233132021020222-3201312321020100-0102323302003303-2213010100210213-0212031113232031-0221300110230303)
- [https.http_protocol_options.http_protocol_enable_v1_only.header_transformation](resources--http_loadbalancer--reference--group-019.md#canonical-3323230131311000-3101333122133201-3231223220131130-0130233100111121-0200300231200030-0320231233333003-1130030000311312-3321123212131300)
- https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.proper_case_header_transformation

<a id="canonical-1003333201033021-0302001102323231-0012211203023021-3003110100310220-3203311303332013-3133000310222133-1112112213012132-2131030310011020"></a>

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

<a id="canonical-1133130012101031-0201031232003130-3231330310221321-3331012002323221-1020121311122231-3022110213202030-3012031213022330-1222330032030330"></a>

## Direct properties — proper_case_header_transformation / 312322023203 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3212300320121033-1232131013233211-0001202233013312-0011100232011211-3302203002212303-0100023012202310-2312002311233223-3323132303330123"></a>

## Next pages — proper_case_header_transformation / 312322023203 / 4

- [https.http_protocol_options.http_protocol_enable_v1_only.header_transformation](resources--http_loadbalancer--reference--group-019.md#canonical-3323230131311000-3101333122133201-3231223220131130-0130233100111121-0200300231200030-0320231233333003-1130030000311312-3321123212131300)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-0010313222101213-2030222023120120-3033111100332312-3000213103211301-1002213101332020-1012332212021033-2333200203003133-1002230333101111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1111310112231103-2321001000323131-0332130312011232-2333033302231003-2033120311321330-3222101000020302-2301030302123331-2000001221110320"></a>

## https.http_protocol_options.http_protocol_enable_v1_v2 — http_protocol_enable_v1_v2 / 220331303020 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [https](resources--http_loadbalancer--reference--group-019.md#canonical-0321232101200210-0011302000022301-2002321223210110-0303132310312201-0222311023330031-3231011200231021-2012211321122303-0001332120000102)
- [https.http_protocol_options](resources--http_loadbalancer--reference--group-019.md#canonical-1231312213210133-3133021101233212-3301313120010020-3113001231200211-0223302012033313-3000301122212000-3131320032301320-1211302233030222)
- https.http_protocol_options.http_protocol_enable_v1_v2

<a id="canonical-2320121312112011-0031310232210001-3110210310120223-0323101321133103-0111033130011101-3310112330111321-1301222231113020-0302000132233123"></a>

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

<a id="canonical-2231110201102032-2333003102000200-0000302221001011-3123112220113233-2323123020201122-3330003000133220-2210100123202212-3302230021131012"></a>

## Direct properties — http_protocol_enable_v1_v2 / 220331303020 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2322110131221100-0113120000001120-2030020030301220-1300310300110211-2102210330203212-2220210031120123-3321010311232130-1100221301200312"></a>

## Next pages — http_protocol_enable_v1_v2 / 220331303020 / 4

- [https.http_protocol_options](resources--http_loadbalancer--reference--group-019.md#canonical-1231312213210133-3133021101233212-3301313120010020-3113001231200211-0223302012033313-3000301122212000-3131320032301320-1211302233030222)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-0213320000020033-2120001112320123-2330332010013101-3101232320223210-0223322031210210-3112031102302331-0322121121003330-1000320310232102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0102112111132232-2203313222011110-2021222330313020-1323221323021313-2013303222121022-3322313232303022-2031231010002211-3133212122321111"></a>

## https.http_protocol_options.http_protocol_enable_v2_only — http_protocol_enable_v2_only / 101231303303 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [https](resources--http_loadbalancer--reference--group-019.md#canonical-0321232101200210-0011302000022301-2002321223210110-0303132310312201-0222311023330031-3231011200231021-2012211321122303-0001332120000102)
- [https.http_protocol_options](resources--http_loadbalancer--reference--group-019.md#canonical-1231312213210133-3133021101233212-3301313120010020-3113001231200211-0223302012033313-3000301122212000-3131320032301320-1211302233030222)
- https.http_protocol_options.http_protocol_enable_v2_only

<a id="canonical-2302203231130002-1022211330223011-3312101311123222-2132331300322312-3030300300310111-1023321333213132-3020130320230010-3202202312010113"></a>

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

<a id="canonical-1331120321033200-1121033300322000-3023001120112233-3213130331022221-0231310031203020-1332123302223332-0312101022232020-1001001021111012"></a>

## Direct properties — http_protocol_enable_v2_only / 101231303303 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2012033233203011-3112222222122132-0330122233332222-1001102133311312-2300201211230220-0023300332113103-1323200223102113-0300130211000202"></a>

## Next pages — http_protocol_enable_v2_only / 101231303303 / 4

- [https.http_protocol_options](resources--http_loadbalancer--reference--group-019.md#canonical-1231312213210133-3133021101233212-3301313120010020-3113001231200211-0223302012033313-3000301122212000-3131320032301320-1211302233030222)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-3122121110210020-0112331300201220-2103232132122022-2101302110211121-3013232333012111-3310330202000103-2001120322233003-2030011312230303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0031011221000202-2121310012031123-2111033302030310-1310231323120002-2303102233211123-3010320100130231-2300231213310302-0202233101321302"></a>

## https.non_default_loadbalancer — non_default_loadbalancer / 303330330033 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [https](resources--http_loadbalancer--reference--group-019.md#canonical-0321232101200210-0011302000022301-2002321223210110-0303132310312201-0222311023330031-3231011200231021-2012211321122303-0001332120000102)
- https.non_default_loadbalancer

<a id="canonical-2322103023232120-0303332033010212-1010213033132213-3232003301130121-2322230231112232-3232202311023330-0332322233332133-3010110313023103"></a>

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

<a id="canonical-2202330222201003-3121230231011230-0302023012233211-0330010221302312-1322303223012020-0200123031320112-1023033112030201-0331312232010133"></a>

## Direct properties — non_default_loadbalancer / 303330330033 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3333213031003032-2333012331003003-1100122000113222-3220221011231213-3312222330202120-3230122333201120-1032202300112322-2233010010110333"></a>

## Next pages — non_default_loadbalancer / 303330330033 / 4

- [https](resources--http_loadbalancer--reference--group-019.md#canonical-0321232101200210-0011302000022301-2002321223210110-0303132310312201-0222311023330031-3231011200231021-2012211321122303-0001332120000102)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-1320021323133122-2131132203230011-0222102300201203-2133211311213200-0330211132231220-3302123120111202-3201010230103132-0121030000310101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3130321133013101-0103130203002111-0112232213123033-0102122111021303-2233010211303113-0021220203023232-1222300323313321-1302023321123111"></a>

## https.pass_through — pass_through / 033113002323 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [https](resources--http_loadbalancer--reference--group-019.md#canonical-0321232101200210-0011302000022301-2002321223210110-0303132310312201-0222311023330031-3231011200231021-2012211321122303-0001332120000102)
- https.pass_through

<a id="canonical-1233131321112112-2023201230013123-3111331122222310-1320121132002322-2001303103033322-0312232120032130-3311231213003322-0001211320110232"></a>

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

<a id="canonical-3233223020220013-0313100020022102-2232333330310120-1312310023210311-3112310210220220-0001022121112300-0333210101301210-2112022113112122"></a>

## Direct properties — pass_through / 033113002323 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1003112210103302-2030032302232211-2301322233233111-0132333301121012-3303231102023032-0222001123320210-2203132202132033-3123122300111132"></a>

## Next pages — pass_through / 033113002323 / 4

- [https](resources--http_loadbalancer--reference--group-019.md#canonical-0321232101200210-0011302000022301-2002321223210110-0303132310312201-0222311023330031-3231011200231021-2012211321122303-0001332120000102)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-0301302212033322-0103133010132330-1223032330030010-2130122212120320-0130211111203213-1023220310223110-3201023013223032-1132311122210133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1101021200030223-1303021110003002-0103231321010202-3010302220110213-1212200103020302-3201223313203002-3130132130312021-1111200010333310"></a>

## https.tls_cert_params — tls_cert_params / 321301030001 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [https](resources--http_loadbalancer--reference--group-019.md#canonical-0321232101200210-0011302000022301-2002321223210110-0303132310312201-0222311023330031-3231011200231021-2012211321122303-0001332120000102)
- https.tls_cert_params

<a id="canonical-3233301011213101-1000031112213101-0010332210132112-2303320013132133-0320331321232100-3331323331331002-1320220313310221-2033203310031223"></a>

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

<a id="canonical-3331120300330213-3210313132021301-1311333021100100-0111301010210210-3201331321113333-3121022020330332-2022102122110230-1330302031300131"></a>

## Direct properties — tls_cert_params / 321301030001 / 3

- [certificates](resources--http_loadbalancer--reference--group-019.md#canonical-2023212112103030-3233311000100230-1003201100002300-0233320011103333-1010103112010201-0010203030023201-3303221000120001-0300233313212031): complete subsection reference.

- [no_mtls](resources--http_loadbalancer--reference--group-019.md#canonical-2122111201321030-0111221001022223-0300313130002313-2221332311201220-3312021030230011-2132312132200123-0003123331003331-3221021123223023): complete subsection reference.

- [tls_config](resources--http_loadbalancer--reference--group-019.md#canonical-0231100000113102-2003020103200311-0230123100111320-1103010233332230-2022220233323033-3002231111332302-3100032303201310-2322332031221120): complete subsection reference.

- [use_mtls](resources--http_loadbalancer--reference--group-019.md#canonical-2030002231213113-2332133122130230-0330233230200231-0121211003120103-0333200112230101-0130122032110113-2313032331321301-3010222210112230): complete subsection reference.

<a id="canonical-3203021231211300-3112212011201011-0001233212301333-2321002312030101-1122331120122131-0323311031101312-1003332110210231-1212103212313302"></a>

## Next pages — tls_cert_params / 321301030001 / 4

- [https.tls_cert_params.certificates](resources--http_loadbalancer--reference--group-019.md#canonical-2023212112103030-3233311000100230-1003201100002300-0233320011103333-1010103112010201-0010203030023201-3303221000120001-0300233313212031)
- [https.tls_cert_params.no_mtls](resources--http_loadbalancer--reference--group-019.md#canonical-2122111201321030-0111221001022223-0300313130002313-2221332311201220-3312021030230011-2132312132200123-0003123331003331-3221021123223023)
- [https.tls_cert_params.tls_config](resources--http_loadbalancer--reference--group-019.md#canonical-0231100000113102-2003020103200311-0230123100111320-1103010233332230-2022220233323033-3002231111332302-3100032303201310-2322332031221120)
- [https.tls_cert_params.use_mtls](resources--http_loadbalancer--reference--group-019.md#canonical-2030002231213113-2332133122130230-0330233230200231-0121211003120103-0333200112230101-0130122032110113-2313032331321301-3010222210112230)
- [https](resources--http_loadbalancer--reference--group-019.md#canonical-0321232101200210-0011302000022301-2002321223210110-0303132310312201-0222311023330031-3231011200231021-2012211321122303-0001332120000102)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-2023212112103030-3233311000100230-1003201100002300-0233320011103333-1010103112010201-0010203030023201-3303221000120001-0300233313212031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3000030212020210-2313032303222103-0030231022011211-0311011300312111-2203011132313220-3030100311301321-1301332221021121-0222013023003003"></a>

## https.tls_cert_params.certificates — certificates / 000302202010 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [https](resources--http_loadbalancer--reference--group-019.md#canonical-0321232101200210-0011302000022301-2002321223210110-0303132310312201-0222311023330031-3231011200231021-2012211321122303-0001332120000102)
- [https.tls_cert_params](resources--http_loadbalancer--reference--group-019.md#canonical-0301302212033322-0103133010132330-1223032330030010-2130122212120320-0130211111203213-1023220310223110-3201023013223032-1132311122210133)
- https.tls_cert_params.certificates

<a id="canonical-1032321203121323-0030221302320223-2300110232011230-2111033221000320-3122321032020332-0110020132032110-2231303202301120-1033332113211033"></a>

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

<a id="canonical-2230010231013023-2123130222012233-1202332303303311-3211312300110013-2220001003111311-0301312303211213-3011213113212013-1022300313202003"></a>

## Direct properties — certificates / 000302202010 / 3

<a id="canonical-3003210313232232-0003212302213011-1320123112133003-1331201302133211-1331111202230012-2113013323202112-0321112103213222-0022032112123211"></a>

<a id="canonical-3222211220021030-0211320130312031-1101211131120132-3121302100310032-0120130302210123-3020310023022112-2203130130123102-2110111020303021"></a>

## name property — certificates / 000302202010 / 4

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

<a id="canonical-1000102132201202-2212033332103002-3302230023111010-0030210223232202-0321211230322333-2321232130120130-3230311332303120-0200201021102303"></a>

<a id="canonical-3313321130331203-2123110013201231-2202031110030130-2123100110013332-1223200320123211-0233001002003221-2321202020302001-3313100132112031"></a>

## namespace property — certificates / 000302202010 / 5

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

<a id="canonical-3121220331101020-3032301300311323-1003312030302301-2010220122330121-3212133211323013-2123121121310320-3033300321021213-2302110033100023"></a>

<a id="canonical-1222303202323113-3211300330201032-1110303133110000-2002210023102222-3012230302122331-0301021302113312-3301330202111123-1131001310020113"></a>

## tenant property — certificates / 000302202010 / 6

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

<a id="canonical-0000220113330323-2012213003220200-1232101203011110-1221010202103320-0122230112213012-1123300002113213-1133023102123113-2011133100013300"></a>

## Next pages — certificates / 000302202010 / 7

- [https.tls_cert_params](resources--http_loadbalancer--reference--group-019.md#canonical-0301302212033322-0103133010132330-1223032330030010-2130122212120320-0130211111203213-1023220310223110-3201023013223032-1132311122210133)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-2122111201321030-0111221001022223-0300313130002313-2221332311201220-3312021030230011-2132312132200123-0003123331003331-3221021123223023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1131210021301331-1213012112212012-2303303201031031-3312103033030030-1311210133032233-3212221132012103-0011000121113112-1012210031310210"></a>

## https.tls_cert_params.no_mtls — no_mtls / 000321131001 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [https](resources--http_loadbalancer--reference--group-019.md#canonical-0321232101200210-0011302000022301-2002321223210110-0303132310312201-0222311023330031-3231011200231021-2012211321122303-0001332120000102)
- [https.tls_cert_params](resources--http_loadbalancer--reference--group-019.md#canonical-0301302212033322-0103133010132330-1223032330030010-2130122212120320-0130211111203213-1023220310223110-3201023013223032-1132311122210133)
- https.tls_cert_params.no_mtls

<a id="canonical-1132213030312033-2222310202131003-2030022120130213-3113032122013002-2002331221202221-1232033023100201-3110110303212201-2332102001030130"></a>

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

<a id="canonical-2121130131030200-0100101121311110-0123202000113022-0231322300101302-1313022203100322-2002221122233310-3030011121012322-1031222112210023"></a>

## Direct properties — no_mtls / 000321131001 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1030310120232323-0311110120301303-1023320012113102-2133332121211020-0112012032331212-2203033011032311-0303103322330001-0220220113120121"></a>

## Next pages — no_mtls / 000321131001 / 4

- [https.tls_cert_params](resources--http_loadbalancer--reference--group-019.md#canonical-0301302212033322-0103133010132330-1223032330030010-2130122212120320-0130211111203213-1023220310223110-3201023013223032-1132311122210133)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-0231100000113102-2003020103200311-0230123100111320-1103010233332230-2022220233323033-3002231111332302-3100032303201310-2322332031221120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0331310213322130-2300122031002113-3302120330110011-0303021100033212-2112123313220232-2233232331122001-2223103221011112-1200213113200123"></a>

## https.tls_cert_params.tls_config — tls_config / 322013130221 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [https](resources--http_loadbalancer--reference--group-019.md#canonical-0321232101200210-0011302000022301-2002321223210110-0303132310312201-0222311023330031-3231011200231021-2012211321122303-0001332120000102)
- [https.tls_cert_params](resources--http_loadbalancer--reference--group-019.md#canonical-0301302212033322-0103133010132330-1223032330030010-2130122212120320-0130211111203213-1023220310223110-3201023013223032-1132311122210133)
- https.tls_cert_params.tls_config

<a id="canonical-2222312213321130-2103021001100102-0322033023130312-0033331332210103-3110331311312231-2331213002030220-0323331203011123-0201100130130232"></a>

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

<a id="canonical-1222110030011101-1211310110221000-0332120012222200-2002202233023320-2230120231330123-0303001011003123-1231133203201002-0302213300100212"></a>

## Direct properties — tls_config / 322013130221 / 3

- [custom_security](resources--http_loadbalancer--reference--group-019.md#canonical-0330033022331213-2311312122223110-2312120231032020-3030213213300212-2303010203200323-2331201011131322-2110333211110310-0101332020332220): complete subsection reference.

- [default_security](resources--http_loadbalancer--reference--group-019.md#canonical-0213322030003310-2001113002220300-3012211332302111-0131203303321231-0300333223202100-0212011001331133-3303230010120130-0023111213123200): complete subsection reference.

- [low_security](resources--http_loadbalancer--reference--group-019.md#canonical-0001210013221011-1011323010230003-2221332100031113-1312303212000221-1223011221210003-3302203201110333-0233010233220032-1200330303230322): complete subsection reference.

- [medium_security](resources--http_loadbalancer--reference--group-019.md#canonical-2321013122303030-1033102001121011-0031020310333012-0130231021202011-1012033131102332-0003111330102131-3022100210113001-1333013311203322): complete subsection reference.

<a id="canonical-1110111021302203-2201103012032202-2000323022320332-0300121221201121-1123011333321201-1122030202331322-1331031000300310-1233130101311201"></a>

## Next pages — tls_config / 322013130221 / 4

- [https.tls_cert_params.tls_config.custom_security](resources--http_loadbalancer--reference--group-019.md#canonical-0330033022331213-2311312122223110-2312120231032020-3030213213300212-2303010203200323-2331201011131322-2110333211110310-0101332020332220)
- [https.tls_cert_params.tls_config.default_security](resources--http_loadbalancer--reference--group-019.md#canonical-0213322030003310-2001113002220300-3012211332302111-0131203303321231-0300333223202100-0212011001331133-3303230010120130-0023111213123200)
- [https.tls_cert_params.tls_config.low_security](resources--http_loadbalancer--reference--group-019.md#canonical-0001210013221011-1011323010230003-2221332100031113-1312303212000221-1223011221210003-3302203201110333-0233010233220032-1200330303230322)
- [https.tls_cert_params.tls_config.medium_security](resources--http_loadbalancer--reference--group-019.md#canonical-2321013122303030-1033102001121011-0031020310333012-0130231021202011-1012033131102332-0003111330102131-3022100210113001-1333013311203322)
- [https.tls_cert_params](resources--http_loadbalancer--reference--group-019.md#canonical-0301302212033322-0103133010132330-1223032330030010-2130122212120320-0130211111203213-1023220310223110-3201023013223032-1132311122210133)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-0330033022331213-2311312122223110-2312120231032020-3030213213300212-2303010203200323-2331201011131322-2110333211110310-0101332020332220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0131101023121210-3313202201213130-1232032011111013-0001013121021220-3102031311031000-0111131233303222-2030100030022002-2222333111213231"></a>

## https.tls_cert_params.tls_config.custom_security — custom_security / 303132213020 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [https](resources--http_loadbalancer--reference--group-019.md#canonical-0321232101200210-0011302000022301-2002321223210110-0303132310312201-0222311023330031-3231011200231021-2012211321122303-0001332120000102)
- [https.tls_cert_params](resources--http_loadbalancer--reference--group-019.md#canonical-0301302212033322-0103133010132330-1223032330030010-2130122212120320-0130211111203213-1023220310223110-3201023013223032-1132311122210133)
- [https.tls_cert_params.tls_config](resources--http_loadbalancer--reference--group-019.md#canonical-0231100000113102-2003020103200311-0230123100111320-1103010233332230-2022220233323033-3002231111332302-3100032303201310-2322332031221120)
- https.tls_cert_params.tls_config.custom_security

<a id="canonical-1222221210221121-1113023102011003-0013031122311030-2131032113032323-1210031111122202-0013031231201102-0220123202113101-3122222121311330"></a>

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

<a id="canonical-0132321021230311-3313310030223100-1222323030000023-1030022031131223-2003323033331100-0231132310233221-2100330100230003-0120112103311133"></a>

## Direct properties — custom_security / 303132213020 / 3

<a id="canonical-2022231313033211-3113221112313301-3100231101122212-3031221011103202-2010201300012223-2030302011303032-0002012311320100-0310312103021210"></a>

<a id="canonical-0121131312131020-3210103311110102-0212113013131031-2221032002303211-2133322111300302-3002120331113131-0312123110223000-3023332031100100"></a>

## cipher_suites property — custom_security / 303132213020 / 4

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

<a id="canonical-3233223030231011-1313212130300210-3331331331222032-1213021221301123-3333121101030022-1322311002201211-0202100211111322-0101010201120210"></a>

<a id="canonical-2020331100333113-3000310221123110-2001321002233112-0010010202220220-2301100131222033-2332113233202302-3132220311201133-1033312133200212"></a>

## max_version property — custom_security / 303132213020 / 5

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

<a id="canonical-3102301103033132-2021232310121220-3000023112230110-1123031233033101-2321303313230212-2233003120032122-3100311310201121-1310021221133333"></a>

<a id="canonical-1123122333013321-2031100211302220-2212110110212121-0311010230321320-3201022213332321-0220122013031301-3121022112233331-3312221331120031"></a>

## min_version property — custom_security / 303132213020 / 6

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

<a id="canonical-1002002013320202-3010321201132021-2020011221110132-0311002113320202-2102220101111220-1201313022020133-2022133032320322-1002123311220233"></a>

## Next pages — custom_security / 303132213020 / 7

- [https.tls_cert_params.tls_config](resources--http_loadbalancer--reference--group-019.md#canonical-0231100000113102-2003020103200311-0230123100111320-1103010233332230-2022220233323033-3002231111332302-3100032303201310-2322332031221120)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-0213322030003310-2001113002220300-3012211332302111-0131203303321231-0300333223202100-0212011001331133-3303230010120130-0023111213123200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3032201222022101-0020130002120000-0132321013002110-1133200103111200-1122232300113320-2321323011322011-1330121203300313-0303230030320132"></a>

## https.tls_cert_params.tls_config.default_security — default_security / 002213202333 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [https](resources--http_loadbalancer--reference--group-019.md#canonical-0321232101200210-0011302000022301-2002321223210110-0303132310312201-0222311023330031-3231011200231021-2012211321122303-0001332120000102)
- [https.tls_cert_params](resources--http_loadbalancer--reference--group-019.md#canonical-0301302212033322-0103133010132330-1223032330030010-2130122212120320-0130211111203213-1023220310223110-3201023013223032-1132311122210133)
- [https.tls_cert_params.tls_config](resources--http_loadbalancer--reference--group-019.md#canonical-0231100000113102-2003020103200311-0230123100111320-1103010233332230-2022220233323033-3002231111332302-3100032303201310-2322332031221120)
- https.tls_cert_params.tls_config.default_security

<a id="canonical-3003231122102310-3332132322221321-3201111122100322-2212332303132201-3130103110310121-0112000022113110-1133332030013010-3103111121103000"></a>

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

<a id="canonical-3331332103201002-1332223103131113-3333212101032131-1220101313230021-0101130231222321-1220211212021131-1231133103121102-0222121310202211"></a>

## Direct properties — default_security / 002213202333 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1220310312002310-1033122111311311-2223122020131330-3110103001123233-0013330112010130-1221231320320012-3010002212330322-0102311112201102"></a>

## Next pages — default_security / 002213202333 / 4

- [https.tls_cert_params.tls_config](resources--http_loadbalancer--reference--group-019.md#canonical-0231100000113102-2003020103200311-0230123100111320-1103010233332230-2022220233323033-3002231111332302-3100032303201310-2322332031221120)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-0001210013221011-1011323010230003-2221332100031113-1312303212000221-1223011221210003-3302203201110333-0233010233220032-1200330303230322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0221032311321031-1112132100112222-0003132211012101-2112103122301311-3203321213223110-1311100001322221-3113222002133321-0010023010313320"></a>

## https.tls_cert_params.tls_config.low_security — low_security / 312330113123 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [https](resources--http_loadbalancer--reference--group-019.md#canonical-0321232101200210-0011302000022301-2002321223210110-0303132310312201-0222311023330031-3231011200231021-2012211321122303-0001332120000102)
- [https.tls_cert_params](resources--http_loadbalancer--reference--group-019.md#canonical-0301302212033322-0103133010132330-1223032330030010-2130122212120320-0130211111203213-1023220310223110-3201023013223032-1132311122210133)
- [https.tls_cert_params.tls_config](resources--http_loadbalancer--reference--group-019.md#canonical-0231100000113102-2003020103200311-0230123100111320-1103010233332230-2022220233323033-3002231111332302-3100032303201310-2322332031221120)
- https.tls_cert_params.tls_config.low_security

<a id="canonical-0212033313212011-0231033103121300-0302030111213210-1132232030332200-0210122303200121-0121021212333301-3221330301002111-0311120121213023"></a>

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

<a id="canonical-2203003033120123-3110101320001031-3031331333102003-0213312303221212-2130323202131210-3230113003012223-3323310100313102-2322100310333232"></a>

## Direct properties — low_security / 312330113123 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0112300223032110-3232122233221323-3132031112013302-3133211221333123-0313310110012230-3000112320202311-3220012213300003-1031203121002103"></a>

## Next pages — low_security / 312330113123 / 4

- [https.tls_cert_params.tls_config](resources--http_loadbalancer--reference--group-019.md#canonical-0231100000113102-2003020103200311-0230123100111320-1103010233332230-2022220233323033-3002231111332302-3100032303201310-2322332031221120)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-2321013122303030-1033102001121011-0031020310333012-0130231021202011-1012033131102332-0003111330102131-3022100210113001-1333013311203322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3221111311031321-2202133333021013-1120300033100300-0232210130300001-2131033033222203-1131312003102311-0011033210011012-1221131103010303"></a>

## https.tls_cert_params.tls_config.medium_security — medium_security / 113313133002 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [https](resources--http_loadbalancer--reference--group-019.md#canonical-0321232101200210-0011302000022301-2002321223210110-0303132310312201-0222311023330031-3231011200231021-2012211321122303-0001332120000102)
- [https.tls_cert_params](resources--http_loadbalancer--reference--group-019.md#canonical-0301302212033322-0103133010132330-1223032330030010-2130122212120320-0130211111203213-1023220310223110-3201023013223032-1132311122210133)
- [https.tls_cert_params.tls_config](resources--http_loadbalancer--reference--group-019.md#canonical-0231100000113102-2003020103200311-0230123100111320-1103010233332230-2022220233323033-3002231111332302-3100032303201310-2322332031221120)
- https.tls_cert_params.tls_config.medium_security

<a id="canonical-1210202102322020-2331320101222301-3032232132000320-0201320103213320-1322101000131111-2213012020203223-1202302113230331-1210002311011200"></a>

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

<a id="canonical-0333020002123003-1103323120001201-1132301322330100-1022321110310203-1233223110021321-3133030132211032-2222020132210010-1022331213131020"></a>

## Direct properties — medium_security / 113313133002 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2011313113021003-2021113200301230-3033200132223001-2220121323232312-0103023213031300-0333333010130333-2301213320030123-1100231122320100"></a>

## Next pages — medium_security / 113313133002 / 4

- [https.tls_cert_params.tls_config](resources--http_loadbalancer--reference--group-019.md#canonical-0231100000113102-2003020103200311-0230123100111320-1103010233332230-2022220233323033-3002231111332302-3100032303201310-2322332031221120)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-2030002231213113-2332133122130230-0330233230200231-0121211003120103-0333200112230101-0130122032110113-2313032331321301-3010222210112230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1113300110111311-1003310131102232-0231132121222022-1130320210100022-2200303003020312-0101132210212013-2230300203130030-1013132121201010"></a>

## https.tls_cert_params.use_mtls — use_mtls / 102023302100 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [https](resources--http_loadbalancer--reference--group-019.md#canonical-0321232101200210-0011302000022301-2002321223210110-0303132310312201-0222311023330031-3231011200231021-2012211321122303-0001332120000102)
- [https.tls_cert_params](resources--http_loadbalancer--reference--group-019.md#canonical-0301302212033322-0103133010132330-1223032330030010-2130122212120320-0130211111203213-1023220310223110-3201023013223032-1132311122210133)
- https.tls_cert_params.use_mtls

<a id="canonical-0120233131212230-3003131013010111-2010100333223103-0002023133011232-1322001121301002-1230111331101223-1132212110333002-0112201203133303"></a>

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

<a id="canonical-2220130211031222-2330023231111203-1322210132320012-0213123131233133-1313310122111132-3032031220100331-2120031133120320-2020222101012100"></a>

## Direct properties — use_mtls / 102023302100 / 3

<a id="canonical-0013220031312131-0232223031233312-1211113323013022-0301321200023000-0323112322030303-3303302231322331-2001302032003231-0020000303030101"></a>

<a id="canonical-1113333211202313-3133121111031313-0003113302332230-3302211310103213-3302220321321212-0320330333202120-3323311013301022-2233103232102303"></a>

## client_certificate_optional property — use_mtls / 102023302100 / 4

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

- [crl](resources--http_loadbalancer--reference--group-019.md#canonical-0213223013112213-3232101122221211-0302300301322233-0031120132110111-3031133310101232-3111303310100122-1123231222120130-3202010133333203): complete subsection reference.

- [no_crl](resources--http_loadbalancer--reference--group-019.md#canonical-1201110112013000-3320022111210132-2022311323333201-2210203323233221-1221331122000310-3102300130213213-2121321222333130-1001100110002213): complete subsection reference.

- [trusted_ca](resources--http_loadbalancer--reference--group-019.md#canonical-2300023113211202-1232310031212033-1103320000111301-2023123322032203-0023301230133132-3222032033210003-3122212131020030-0331232002312013): complete subsection reference.

<a id="canonical-3202122212331120-2101031012331230-3020130231223023-1132012101013212-2333201012013333-3330120320222233-0201120123332131-3111213202221010"></a>

<a id="canonical-0202003302020331-3211100232211302-3333321133023231-0301102130022303-0002301200323322-2202233301013132-1210001201330122-2331021030220331"></a>

## trusted_ca_url property — use_mtls / 102023302100 / 5

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

- [xfcc_disabled](resources--http_loadbalancer--reference--group-019.md#canonical-1023120331133303-0110123333003303-3223301012320202-0022201331221332-2233231110010223-0003022131301110-1331332021011132-0010020012321330): complete subsection reference.

- [xfcc_options](resources--http_loadbalancer--reference--group-019.md#canonical-2030032300301111-0230222233121102-1221202032221023-0022320332212033-3312010223022100-2112211333011132-3103013020023032-1101213202000203): complete subsection reference.

<a id="canonical-1123123203022332-2332133010233003-3221311203100132-2212012110003003-3002233033011133-3201103202010021-0102122001312013-2313131210030320"></a>

## Next pages — use_mtls / 102023302100 / 6

- [https.tls_cert_params.use_mtls.crl](resources--http_loadbalancer--reference--group-019.md#canonical-0213223013112213-3232101122221211-0302300301322233-0031120132110111-3031133310101232-3111303310100122-1123231222120130-3202010133333203)
- [https.tls_cert_params.use_mtls.no_crl](resources--http_loadbalancer--reference--group-019.md#canonical-1201110112013000-3320022111210132-2022311323333201-2210203323233221-1221331122000310-3102300130213213-2121321222333130-1001100110002213)
- [https.tls_cert_params.use_mtls.trusted_ca](resources--http_loadbalancer--reference--group-019.md#canonical-2300023113211202-1232310031212033-1103320000111301-2023123322032203-0023301230133132-3222032033210003-3122212131020030-0331232002312013)
- [https.tls_cert_params.use_mtls.xfcc_disabled](resources--http_loadbalancer--reference--group-019.md#canonical-1023120331133303-0110123333003303-3223301012320202-0022201331221332-2233231110010223-0003022131301110-1331332021011132-0010020012321330)
- [https.tls_cert_params.use_mtls.xfcc_options](resources--http_loadbalancer--reference--group-019.md#canonical-2030032300301111-0230222233121102-1221202032221023-0022320332212033-3312010223022100-2112211333011132-3103013020023032-1101213202000203)
- [https.tls_cert_params](resources--http_loadbalancer--reference--group-019.md#canonical-0301302212033322-0103133010132330-1223032330030010-2130122212120320-0130211111203213-1023220310223110-3201023013223032-1132311122210133)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-0213223013112213-3232101122221211-0302300301322233-0031120132110111-3031133310101232-3111303310100122-1123231222120130-3202010133333203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2011012022330111-3013311313102302-2000111132020023-2030203312111233-1330122120033323-2013211111031130-3222212033231313-2101122330122003"></a>

## https.tls_cert_params.use_mtls.crl — crl / 003221231110 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [https](resources--http_loadbalancer--reference--group-019.md#canonical-0321232101200210-0011302000022301-2002321223210110-0303132310312201-0222311023330031-3231011200231021-2012211321122303-0001332120000102)
- [https.tls_cert_params](resources--http_loadbalancer--reference--group-019.md#canonical-0301302212033322-0103133010132330-1223032330030010-2130122212120320-0130211111203213-1023220310223110-3201023013223032-1132311122210133)
- [https.tls_cert_params.use_mtls](resources--http_loadbalancer--reference--group-019.md#canonical-2030002231213113-2332133122130230-0330233230200231-0121211003120103-0333200112230101-0130122032110113-2313032331321301-3010222210112230)
- https.tls_cert_params.use_mtls.crl

<a id="canonical-2012322203013010-3010331231321120-2300333231330332-3220230120202321-2233321211203132-2320211021000231-3203033120001300-2011311302311301"></a>

Type: `"object"`. single nested block, Optional.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

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
crl {
  # Configure direct properties listed below.
}
```

<a id="canonical-0120313232223033-3113223310300301-0122301312311201-1101213002012113-1132303221200301-3200221033331113-1102100100020333-2223033030331120"></a>

## Direct properties — crl / 003221231110 / 3

<a id="canonical-2323023203333030-3300320203111331-3101332120001301-1113330122033111-1303132332223132-1202133101110123-1110003112003101-0233033001011232"></a>

<a id="canonical-3023003030203021-1330230232030121-3100102311031130-0333203332220021-1212211303022201-0012203230311022-1303320110131230-1202130300110023"></a>

## name property — crl / 003221231110 / 4

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

<a id="canonical-1332022311011031-2012031300021312-2200303101003003-2031012221120201-2030322220312103-0331110303302100-1311321020210213-1322003333303310"></a>

<a id="canonical-0210123001013303-3032301213102130-3132111201223002-0100131222222121-1213301323210232-0110323331220031-2303211212101311-3121231203113002"></a>

## namespace property — crl / 003221231110 / 5

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

<a id="canonical-0130011110210333-1211203010313322-0123233313021320-2212211222231333-3231231301212320-1300133112313332-0332013231101122-1322231330001101"></a>

<a id="canonical-1021332202333220-3232333120301010-0003130302210013-3303130022002130-2001303110232130-1313011302001333-1222012022212221-0211220200310030"></a>

## tenant property — crl / 003221231110 / 6

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

<a id="canonical-1111211210201221-0033123233303203-0102332112010300-0123231322211123-2332010132000231-2223202111212320-0032000310130221-2232201330010120"></a>

## Next pages — crl / 003221231110 / 7

- [https.tls_cert_params.use_mtls](resources--http_loadbalancer--reference--group-019.md#canonical-2030002231213113-2332133122130230-0330233230200231-0121211003120103-0333200112230101-0130122032110113-2313032331321301-3010222210112230)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-1201110112013000-3320022111210132-2022311323333201-2210203323233221-1221331122000310-3102300130213213-2121321222333130-1001100110002213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0003321211330103-2023013022121323-3102320231212131-0333321131112123-1000032020123011-1330120230310323-1112330201121021-1111200220200212"></a>

## https.tls_cert_params.use_mtls.no_crl — no_crl / 231210131013 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [https](resources--http_loadbalancer--reference--group-019.md#canonical-0321232101200210-0011302000022301-2002321223210110-0303132310312201-0222311023330031-3231011200231021-2012211321122303-0001332120000102)
- [https.tls_cert_params](resources--http_loadbalancer--reference--group-019.md#canonical-0301302212033322-0103133010132330-1223032330030010-2130122212120320-0130211111203213-1023220310223110-3201023013223032-1132311122210133)
- [https.tls_cert_params.use_mtls](resources--http_loadbalancer--reference--group-019.md#canonical-2030002231213113-2332133122130230-0330233230200231-0121211003120103-0333200112230101-0130122032110113-2313032331321301-3010222210112230)
- https.tls_cert_params.use_mtls.no_crl

<a id="canonical-2311031021103201-2222320202010023-3120221323230223-2013113130330103-0003202112201030-2110022012130133-1333322121111313-0220202231131312"></a>

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
no_crl = {}
```

<a id="canonical-2101222212222122-3322303020031020-2011333212000020-2321303321233133-0210223223332132-0332212013123110-1110331113302310-3310211131232300"></a>

## Direct properties — no_crl / 231210131013 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3030130031302231-0012332333301233-1001033231023313-3321022103333221-1133312010320102-1131112232310223-3311011131220200-3023300232013000"></a>

## Next pages — no_crl / 231210131013 / 4

- [https.tls_cert_params.use_mtls](resources--http_loadbalancer--reference--group-019.md#canonical-2030002231213113-2332133122130230-0330233230200231-0121211003120103-0333200112230101-0130122032110113-2313032331321301-3010222210112230)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-2300023113211202-1232310031212033-1103320000111301-2023123322032203-0023301230133132-3222032033210003-3122212131020030-0331232002312013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0121322122313223-1213002112210202-0323103312333012-3212130321201202-2211210130213030-1212200200311113-1202323030001201-2233110330003030"></a>

## https.tls_cert_params.use_mtls.trusted_ca — trusted_ca / 100303312120 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [https](resources--http_loadbalancer--reference--group-019.md#canonical-0321232101200210-0011302000022301-2002321223210110-0303132310312201-0222311023330031-3231011200231021-2012211321122303-0001332120000102)
- [https.tls_cert_params](resources--http_loadbalancer--reference--group-019.md#canonical-0301302212033322-0103133010132330-1223032330030010-2130122212120320-0130211111203213-1023220310223110-3201023013223032-1132311122210133)
- [https.tls_cert_params.use_mtls](resources--http_loadbalancer--reference--group-019.md#canonical-2030002231213113-2332133122130230-0330233230200231-0121211003120103-0333200112230101-0130122032110113-2313032331321301-3010222210112230)
- https.tls_cert_params.use_mtls.trusted_ca

<a id="canonical-0200202013332301-3011333303012332-3110310021230312-0030311033231133-0130122031202321-1221020333030122-3103221010320211-2302322101210131"></a>

Type: `"object"`. single nested block, Optional.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

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
trusted_ca {
  # Configure direct properties listed below.
}
```

<a id="canonical-0313231211212222-3233101131113300-2133000020003001-3213202003000032-1020331210121030-0331333000002002-0003021310311123-2101031032000320"></a>

## Direct properties — trusted_ca / 100303312120 / 3

<a id="canonical-3200201120220100-3223330113010000-1312013222111231-0101003002332103-3020011013000010-3001111030332103-3111230000102020-2130212133323022"></a>

<a id="canonical-1320221311330332-1221013100211323-1000012303021121-3222230301010121-2222311030013210-1102131010103030-1312001010022003-3022121021010013"></a>

## name property — trusted_ca / 100303312120 / 4

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

<a id="canonical-0113001120032211-2031112100021103-3122110213031220-2033200231322322-0213200320012023-2321031233331021-2102231313223001-3203221310112210"></a>

<a id="canonical-0002213300010311-2000231021000221-1220113310212330-1300032033100120-2203222203011101-0133232333320033-1311301032103210-3130000120103230"></a>

## namespace property — trusted_ca / 100303312120 / 5

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

<a id="canonical-1121231212320221-1223130121131133-3320131000311312-0202331210313102-3103320203022131-1311321023032312-2121132332000133-3022113300030203"></a>

<a id="canonical-0122320223002203-3312302011001111-1123212102330301-0030213323333230-2133001231332130-2031123311133113-0332023001233231-1323012221112001"></a>

## tenant property — trusted_ca / 100303312120 / 6

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

<a id="canonical-2103013203202320-3022220222310003-3112211102013103-2012212121201201-2031122121203222-0112302330131101-2110200223333112-3322011021201213"></a>

## Next pages — trusted_ca / 100303312120 / 7

- [https.tls_cert_params.use_mtls](resources--http_loadbalancer--reference--group-019.md#canonical-2030002231213113-2332133122130230-0330233230200231-0121211003120103-0333200112230101-0130122032110113-2313032331321301-3010222210112230)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-1023120331133303-0110123333003303-3223301012320202-0022201331221332-2233231110010223-0003022131301110-1331332021011132-0010020012321330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0201110031233002-3002303002213012-2222013130312012-0213313033003331-0232010300311023-1001031003122202-1123101332312031-1202113001012300"></a>

## https.tls_cert_params.use_mtls.xfcc_disabled — xfcc_disabled / 330310321033 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [https](resources--http_loadbalancer--reference--group-019.md#canonical-0321232101200210-0011302000022301-2002321223210110-0303132310312201-0222311023330031-3231011200231021-2012211321122303-0001332120000102)
- [https.tls_cert_params](resources--http_loadbalancer--reference--group-019.md#canonical-0301302212033322-0103133010132330-1223032330030010-2130122212120320-0130211111203213-1023220310223110-3201023013223032-1132311122210133)
- [https.tls_cert_params.use_mtls](resources--http_loadbalancer--reference--group-019.md#canonical-2030002231213113-2332133122130230-0330233230200231-0121211003120103-0333200112230101-0130122032110113-2313032331321301-3010222210112230)
- https.tls_cert_params.use_mtls.xfcc_disabled

<a id="canonical-1233213323033011-0102231200002012-1111301021122222-1122210222003210-1222222010232023-1310201213310131-0331023203020220-1221103101031112"></a>

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
xfcc_disabled = {}
```

<a id="canonical-1331231032231003-3330112200032202-2101101332212320-3212102102311011-3021123200202020-3003112211331330-0331022021132303-2301121332002232"></a>

## Direct properties — xfcc_disabled / 330310321033 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2001121023131021-1001113223213320-0123312012312320-1333132223110002-3303213303203032-0103021033203032-1101333233112313-2133000313232101"></a>

## Next pages — xfcc_disabled / 330310321033 / 4

- [https.tls_cert_params.use_mtls](resources--http_loadbalancer--reference--group-019.md#canonical-2030002231213113-2332133122130230-0330233230200231-0121211003120103-0333200112230101-0130122032110113-2313032331321301-3010222210112230)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-2030032300301111-0230222233121102-1221202032221023-0022320332212033-3312010223022100-2112211333011132-3103013020023032-1101213202000203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3011023323211321-2030110300210213-2000211030210213-1100003031310300-2003203033323020-3320210303310121-0220020131231031-3303213303300032"></a>

## https.tls_cert_params.use_mtls.xfcc_options — xfcc_options / 011313332110 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [https](resources--http_loadbalancer--reference--group-019.md#canonical-0321232101200210-0011302000022301-2002321223210110-0303132310312201-0222311023330031-3231011200231021-2012211321122303-0001332120000102)
- [https.tls_cert_params](resources--http_loadbalancer--reference--group-019.md#canonical-0301302212033322-0103133010132330-1223032330030010-2130122212120320-0130211111203213-1023220310223110-3201023013223032-1132311122210133)
- [https.tls_cert_params.use_mtls](resources--http_loadbalancer--reference--group-019.md#canonical-2030002231213113-2332133122130230-0330233230200231-0121211003120103-0333200112230101-0130122032110113-2313032331321301-3010222210112230)
- https.tls_cert_params.use_mtls.xfcc_options

<a id="canonical-0312131213031111-2232231320233001-3131132311033233-0222321311133113-1131001222122010-1232231022100332-1321031333120300-2000123300111022"></a>

Type: `"object"`. single nested block, Optional.

X-Forwarded-Client-Cert header elements to be added to requests.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-2000111203210211-1103303200132300-1010220232020331-1212221212012011-0011013030101321-1300022202022232-1313032233221030-3202000220213322"></a>

## Direct properties — xfcc_options / 011313332110 / 3

<a id="canonical-1203012023003231-3330333302122121-3123033131030022-1220101313111033-3100322222301210-3320002210023233-1112202223131221-3202231133212330"></a>

<a id="canonical-1333033112020022-0023332133213312-0230332113022003-2000201321323120-1121102232313202-2333101121101121-1120033220032232-3310003232330202"></a>

## xfcc_header_elements property — xfcc_options / 011313332110 / 4

Type: `["list", "string"]`. Optional.

\[Enum: XFCC\_NONE|XFCC\_CERT|XFCC\_CHAIN|XFCC\_SUBJECT|XFCC\_URI|XFCC\_DNS\]
X-Forwarded-Client-Cert header elements to be added to requests. Possible values are \`XFCC\_NONE\`,
\`XFCC\_CERT\`, \`XFCC\_CHAIN\`, \`XFCC\_SUBJECT\`, \`XFCC\_URI\`, \`XFCC\_DNS\`. Defaults to
\`XFCC\_NONE\`.

Upstream description:

X-Forwarded-Client-Cert header elements to be added to requests.

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

<a id="canonical-2322312211130110-1230121332102312-0123131132122221-1221111120101011-2233012132222110-3122321030230123-2011221322303030-2012122133023023"></a>

## Next pages — xfcc_options / 011313332110 / 5

- [https.tls_cert_params.use_mtls](resources--http_loadbalancer--reference--group-019.md#canonical-2030002231213113-2332133122130230-0330233230200231-0121211003120103-0333200112230101-0130122032110113-2313032331321301-3010222210112230)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-1231113013303201-3233222000300212-2211131130013302-2322211003010221-0320233103301130-2102232111212102-2321212333113100-3130203302022212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0122130011222301-2222133100100312-1232321012221031-1121300132111231-0023011203230102-1200102030003223-2322122121020120-0032332323330020"></a>

## https.tls_parameters — tls_parameters / 302333123000 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [https](resources--http_loadbalancer--reference--group-019.md#canonical-0321232101200210-0011302000022301-2002321223210110-0303132310312201-0222311023330031-3231011200231021-2012211321122303-0001332120000102)
- https.tls_parameters

<a id="canonical-2303112103210221-3112000301021210-3001130121323132-1023131031001322-2231111331031201-2200303303211200-0001320332231020-1131213312202302"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for tls parameters.

Upstream description:

Inline TLS parameters.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-0313123321320321-2130123102113003-2201103333023310-2330130203010011-0103131232120100-0020032213023001-0322220133312032-2233113322012122"></a>

## Direct properties — tls_parameters / 302333123000 / 3

- [no_mtls](resources--http_loadbalancer--reference--group-019.md#canonical-0310032130223213-0000030332120330-3331033231032110-0012003112113033-3121211021200000-3321212301033033-0111112121023132-1102003021003133): complete subsection reference.

- [tls_certificates](resources--http_loadbalancer--reference--group-019.md#canonical-2333201003023222-1221313103330132-3312310111023023-3001030030313321-2220312323120110-2133112333103203-1301231023103233-1333103021220231): complete subsection reference.

- [tls_config](resources--http_loadbalancer--reference--group-019.md#canonical-1331301222322132-3310303011021332-0330022222232121-3232113210300302-3100323021101302-2323203211312232-2232012310223322-1333222331330000): complete subsection reference.

- [use_mtls](resources--http_loadbalancer--reference--group-019.md#canonical-1321332101331311-1023112313030210-2312030203110123-1202033022020030-0101101100102110-2022201111233012-3210333032232312-3311120222312230): complete subsection reference.

<a id="canonical-0132103022000322-0331103203211212-2021110133031331-0302013321111010-2133133302113110-3211000301130323-0333223120221111-3223322132131012"></a>

## Next pages — tls_parameters / 302333123000 / 4

- [https.tls_parameters.no_mtls](resources--http_loadbalancer--reference--group-019.md#canonical-0310032130223213-0000030332120330-3331033231032110-0012003112113033-3121211021200000-3321212301033033-0111112121023132-1102003021003133)
- [https.tls_parameters.tls_certificates](resources--http_loadbalancer--reference--group-019.md#canonical-2333201003023222-1221313103330132-3312310111023023-3001030030313321-2220312323120110-2133112333103203-1301231023103233-1333103021220231)
- [https.tls_parameters.tls_config](resources--http_loadbalancer--reference--group-019.md#canonical-1331301222322132-3310303011021332-0330022222232121-3232113210300302-3100323021101302-2323203211312232-2232012310223322-1333222331330000)
- [https.tls_parameters.use_mtls](resources--http_loadbalancer--reference--group-019.md#canonical-1321332101331311-1023112313030210-2312030203110123-1202033022020030-0101101100102110-2022201111233012-3210333032232312-3311120222312230)
- [https](resources--http_loadbalancer--reference--group-019.md#canonical-0321232101200210-0011302000022301-2002321223210110-0303132310312201-0222311023330031-3231011200231021-2012211321122303-0001332120000102)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-0310032130223213-0000030332120330-3331033231032110-0012003112113033-3121211021200000-3321212301033033-0111112121023132-1102003021003133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3323122011333122-2033203110001332-2313223112321113-0221101010222213-0210103000112121-3212212000120322-0012313322222133-1122300233123032"></a>

## https.tls_parameters.no_mtls — no_mtls / 031221130110 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [https](resources--http_loadbalancer--reference--group-019.md#canonical-0321232101200210-0011302000022301-2002321223210110-0303132310312201-0222311023330031-3231011200231021-2012211321122303-0001332120000102)
- [https.tls_parameters](resources--http_loadbalancer--reference--group-019.md#canonical-1231113013303201-3233222000300212-2211131130013302-2322211003010221-0320233103301130-2102232111212102-2321212333113100-3130203302022212)
- https.tls_parameters.no_mtls

<a id="canonical-3112333032131022-0330032203030032-0321300033001222-3112310130230023-2333121000023320-3320220222232010-0022201101303002-3003333332131023"></a>

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

<a id="canonical-2133221121031310-0301212233030301-2211132133210003-1300220010231103-0122011012032000-0120130100323022-2013321222000233-3011203022013202"></a>

## Direct properties — no_mtls / 031221130110 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3101230023011102-0133221312301010-3131212331010102-3020310322202022-0000132022230321-2333232001030122-2013201313312131-0101330103201120"></a>

## Next pages — no_mtls / 031221130110 / 4

- [https.tls_parameters](resources--http_loadbalancer--reference--group-019.md#canonical-1231113013303201-3233222000300212-2211131130013302-2322211003010221-0320233103301130-2102232111212102-2321212333113100-3130203302022212)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-2333201003023222-1221313103330132-3312310111023023-3001030030313321-2220312323120110-2133112333103203-1301231023103233-1333103021220231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1212133100310101-2320231220030010-3131130030322002-1223122330031020-3112220221200203-0011032302212100-3010311001032010-2231220033211202"></a>

## https.tls_parameters.tls_certificates — tls_certificates / 331222321232 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [https](resources--http_loadbalancer--reference--group-019.md#canonical-0321232101200210-0011302000022301-2002321223210110-0303132310312201-0222311023330031-3231011200231021-2012211321122303-0001332120000102)
- [https.tls_parameters](resources--http_loadbalancer--reference--group-019.md#canonical-1231113013303201-3233222000300212-2211131130013302-2322211003010221-0320233103301130-2102232111212102-2321212333113100-3130203302022212)
- https.tls_parameters.tls_certificates

<a id="canonical-2003103320220312-0010201010022112-1203212022022000-1013230000322033-3301303303232123-0011213131220222-0101212332232113-0122310330203310"></a>

Type: `"object"`. list nested block, Optional.

Users can add one or more certificates that share the same set of domains. For example, domain.com
and \*.domain.com - but use different signature algorithms.

Upstream description:

Users can add one or more certificates that share the same set of domains. For example, domain.com
and \*.domain.com - but use different signature algorithms.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-1320200323030202-0212322011233300-2310300223133220-0002030310100321-3130220131113130-2003100121032321-3010321330122310-2032223230200311"></a>

## Direct properties — tls_certificates / 331222321232 / 3

<a id="canonical-3302233212023301-2110131231321233-3312333133312012-0311110123023033-1331012303122203-0111101333303302-2222333133223301-3303102223232332"></a>

<a id="canonical-0101323223303113-1221013001102000-3202102300311203-3121230112302220-3212130010200333-0113112110230303-1212312013333032-0211110120020203"></a>

## certificate_url property — tls_certificates / 331222321232 / 4

Type: `"string"`. Optional.

TLS certificate. Certificate or certificate chain in PEM format including the PEM headers.

Upstream description:

TLS certificate. Certificate or certificate chain in PEM format including the PEM headers.

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

- [custom_hash_algorithms](resources--http_loadbalancer--reference--group-019.md#canonical-2013303333001103-3011311020010221-1123023321112203-2333031023203133-3321020113223312-0221233102012220-0130102030032221-1231011223320321): complete subsection reference.

<a id="canonical-1101011030321211-1333311312123032-3113323312021121-1010202213030322-1323330223313203-2212221213001312-2003320020000020-2331001021222132"></a>

<a id="canonical-2031113121121220-2311123003222010-1232112020212312-2231203103232002-1123011221232320-1112233030011321-0020231022111011-2233010322123201"></a>

## description_spec property — tls_certificates / 331222321232 / 5

Type: `"string"`. Optional.

Description. Description for the certificate.

- [disable_ocsp_stapling](resources--http_loadbalancer--reference--group-019.md#canonical-1302212313232323-3131331323202201-0203000320021212-0032301212020221-1120030111300331-3203132311220220-0202202133102022-1033120123120132): complete subsection reference.

- [private_key](resources--http_loadbalancer--reference--group-019.md#canonical-2012002302202202-2020122200300213-3021321323122130-1322013120010111-0102301303322223-1233213233122333-3123201031333301-3110202011013000): complete subsection reference.

- [use_system_defaults](resources--http_loadbalancer--reference--group-019.md#canonical-1233033031232212-2012200221121213-2333013220123312-1032011003203332-2330130201010101-1103323120112201-2030313103301202-1202301301322103): complete subsection reference.

<a id="canonical-0023210201013302-1031100002313131-2133101320323311-0300102132112213-3210210111001331-0232202212203220-1330232113330322-3211302210330031"></a>

## Next pages — tls_certificates / 331222321232 / 6

- [https.tls_parameters.tls_certificates.custom_hash_algorithms](resources--http_loadbalancer--reference--group-019.md#canonical-2013303333001103-3011311020010221-1123023321112203-2333031023203133-3321020113223312-0221233102012220-0130102030032221-1231011223320321)
- [https.tls_parameters.tls_certificates.disable_ocsp_stapling](resources--http_loadbalancer--reference--group-019.md#canonical-1302212313232323-3131331323202201-0203000320021212-0032301212020221-1120030111300331-3203132311220220-0202202133102022-1033120123120132)
- [https.tls_parameters.tls_certificates.private_key](resources--http_loadbalancer--reference--group-019.md#canonical-2012002302202202-2020122200300213-3021321323122130-1322013120010111-0102301303322223-1233213233122333-3123201031333301-3110202011013000)
- [https.tls_parameters.tls_certificates.use_system_defaults](resources--http_loadbalancer--reference--group-019.md#canonical-1233033031232212-2012200221121213-2333013220123312-1032011003203332-2330130201010101-1103323120112201-2030313103301202-1202301301322103)
- [https.tls_parameters](resources--http_loadbalancer--reference--group-019.md#canonical-1231113013303201-3233222000300212-2211131130013302-2322211003010221-0320233103301130-2102232111212102-2321212333113100-3130203302022212)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-2013303333001103-3011311020010221-1123023321112203-2333031023203133-3321020113223312-0221233102012220-0130102030032221-1231011223320321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3310120332111103-1210313020332313-1002213313302201-1232020103200011-1322322111002000-1300130131332332-0011303300100100-3011300001220221"></a>

## https.tls_parameters.tls_certificates.custom_hash_algorithms — custom_hash_algorithms / 302213330302 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [https](resources--http_loadbalancer--reference--group-019.md#canonical-0321232101200210-0011302000022301-2002321223210110-0303132310312201-0222311023330031-3231011200231021-2012211321122303-0001332120000102)
- [https.tls_parameters](resources--http_loadbalancer--reference--group-019.md#canonical-1231113013303201-3233222000300212-2211131130013302-2322211003010221-0320233103301130-2102232111212102-2321212333113100-3130203302022212)
- [https.tls_parameters.tls_certificates](resources--http_loadbalancer--reference--group-019.md#canonical-2333201003023222-1221313103330132-3312310111023023-3001030030313321-2220312323120110-2133112333103203-1301231023103233-1333103021220231)
- https.tls_parameters.tls_certificates.custom_hash_algorithms

<a id="canonical-2132020131303202-0112110030003012-0313110103312311-2123301033020232-3132033201212000-0223130230013212-0122002033310030-3312321022223231"></a>

Type: `"object"`. single nested block, Optional.

Specifies the hash algorithms to be used.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-2211112330030033-3120131211301112-3210211333110132-1113212121133233-2302200231333112-1232000313322030-1031030021002210-3332330211022110"></a>

## Direct properties — custom_hash_algorithms / 302213330302 / 3

<a id="canonical-2233030010212333-2000232020112001-3111221333122220-2011101112300103-3133312203222101-1110310232210001-1333300033333120-0202032200330013"></a>

<a id="canonical-1203031213133111-2321213021022210-0121223031322012-2020230121212332-0121101222320321-3031000232220112-3131311310332202-2122112103121220"></a>

## hash_algorithms property — custom_hash_algorithms / 302213330302 / 4

Type: `["list", "string"]`. Optional.

\[Enum: INVALID\_HASH\_ALGORITHM|SHA256|SHA1\] Ordered list of hash algorithms to be used. Possible
values are \`INVALID\_HASH\_ALGORITHM\`, \`SHA256\`, \`SHA1\`. Defaults to
\`INVALID\_HASH\_ALGORITHM\`.

Upstream description:

Ordered list of hash algorithms to be used.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-0312330313101100-1002221131300332-1320320330312300-1231003103223330-0130310122012311-3132331322102312-1101303333000011-2101231021203122"></a>

## Next pages — custom_hash_algorithms / 302213330302 / 5

- [https.tls_parameters.tls_certificates](resources--http_loadbalancer--reference--group-019.md#canonical-2333201003023222-1221313103330132-3312310111023023-3001030030313321-2220312323120110-2133112333103203-1301231023103233-1333103021220231)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-1302212313232323-3131331323202201-0203000320021212-0032301212020221-1120030111300331-3203132311220220-0202202133102022-1033120123120132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2120332020113211-3300013121000201-2123322100221310-0323232111103103-0303313112201113-0200313001110230-2122203311131301-1330101101230332"></a>

## https.tls_parameters.tls_certificates.disable_ocsp_stapling — disable_ocsp_stapling / 233202202202 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [https](resources--http_loadbalancer--reference--group-019.md#canonical-0321232101200210-0011302000022301-2002321223210110-0303132310312201-0222311023330031-3231011200231021-2012211321122303-0001332120000102)
- [https.tls_parameters](resources--http_loadbalancer--reference--group-019.md#canonical-1231113013303201-3233222000300212-2211131130013302-2322211003010221-0320233103301130-2102232111212102-2321212333113100-3130203302022212)
- [https.tls_parameters.tls_certificates](resources--http_loadbalancer--reference--group-019.md#canonical-2333201003023222-1221313103330132-3312310111023023-3001030030313321-2220312323120110-2133112333103203-1301231023103233-1333103021220231)
- https.tls_parameters.tls_certificates.disable_ocsp_stapling

<a id="canonical-3121222102322021-2120213120022200-1323300213021113-1302123231233001-0120011301230031-3101120130202101-1330022222101011-3130213121132303"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disable ocsp stapling.

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
disable_ocsp_stapling = {}
```

<a id="canonical-1311332111003310-2320301303311310-2311232012001232-0023002210012203-1302100210031233-0313010030021221-1201030320030132-3313300020101222"></a>

## Direct properties — disable_ocsp_stapling / 233202202202 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0013320211002000-3200300203013003-0131331213313001-2101211302333210-2301333232110113-1100000100032110-1200301120212030-1132131222133331"></a>

## Next pages — disable_ocsp_stapling / 233202202202 / 4

- [https.tls_parameters.tls_certificates](resources--http_loadbalancer--reference--group-019.md#canonical-2333201003023222-1221313103330132-3312310111023023-3001030030313321-2220312323120110-2133112333103203-1301231023103233-1333103021220231)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-2012002302202202-2020122200300213-3021321323122130-1322013120010111-0102301303322223-1233213233122333-3123201031333301-3110202011013000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0131001013032222-1101023113003003-1301130313131211-2021301330013112-2033221303100231-1323200322132023-0310130220302223-1133211011222131"></a>

## https.tls_parameters.tls_certificates.private_key — private_key / 031112100221 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [https](resources--http_loadbalancer--reference--group-019.md#canonical-0321232101200210-0011302000022301-2002321223210110-0303132310312201-0222311023330031-3231011200231021-2012211321122303-0001332120000102)
- [https.tls_parameters](resources--http_loadbalancer--reference--group-019.md#canonical-1231113013303201-3233222000300212-2211131130013302-2322211003010221-0320233103301130-2102232111212102-2321212333113100-3130203302022212)
- [https.tls_parameters.tls_certificates](resources--http_loadbalancer--reference--group-019.md#canonical-2333201003023222-1221313103330132-3312310111023023-3001030030313321-2220312323120110-2133112333103203-1301231023103233-1333103021220231)
- https.tls_parameters.tls_certificates.private_key

<a id="canonical-1322303213121132-1322202213231111-0011300132033000-0201213301131201-2121001010203001-1232220000311021-1231110000100022-0032333011030133"></a>

Type: `"object"`. single nested block, Optional.

SecretType is used in an object to indicate a sensitive/confidential field.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-3223222331013002-1222001100011222-2302312013221031-0133021301031230-3121010333003020-3020233221023212-3010322031102032-0200222303230111"></a>

## Direct properties — private_key / 031112100221 / 3

- [blindfold_secret_info](resources--http_loadbalancer--reference--group-019.md#canonical-2101021213111133-3000323023020322-3313130212120212-0320323223112323-1331130021032013-0032110321320020-3200120311301200-0031111303303012): complete subsection reference.

- [clear_secret_info](resources--http_loadbalancer--reference--group-019.md#canonical-1211202021023200-3331130303131220-3300220332121300-0233010100332222-2012303203323012-0010212212001101-0120013221322220-2101311321110123): complete subsection reference.

<a id="canonical-0012313013131232-1100020001210320-0230202322022012-2102230300211300-3122001010030031-0210313030322323-1230131210112002-1010033001000310"></a>

## Next pages — private_key / 031112100221 / 4

- [https.tls_parameters.tls_certificates.private_key.blindfold_secret_info](resources--http_loadbalancer--reference--group-019.md#canonical-2101021213111133-3000323023020322-3313130212120212-0320323223112323-1331130021032013-0032110321320020-3200120311301200-0031111303303012)
- [https.tls_parameters.tls_certificates.private_key.clear_secret_info](resources--http_loadbalancer--reference--group-019.md#canonical-1211202021023200-3331130303131220-3300220332121300-0233010100332222-2012303203323012-0010212212001101-0120013221322220-2101311321110123)
- [https.tls_parameters.tls_certificates](resources--http_loadbalancer--reference--group-019.md#canonical-2333201003023222-1221313103330132-3312310111023023-3001030030313321-2220312323120110-2133112333103203-1301231023103233-1333103021220231)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-2101021213111133-3000323023020322-3313130212120212-0320323223112323-1331130021032013-0032110321320020-3200120311301200-0031111303303012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3000013222031032-2332100210011102-2310313211303310-2113300032221120-2202120020120023-0023122232320003-0322113312310132-1320101333011232"></a>

## https.tls_parameters.tls_certificates.private_key.blindfold_secret_info — blindfold_secret_info / 022023122200 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [https](resources--http_loadbalancer--reference--group-019.md#canonical-0321232101200210-0011302000022301-2002321223210110-0303132310312201-0222311023330031-3231011200231021-2012211321122303-0001332120000102)
- [https.tls_parameters](resources--http_loadbalancer--reference--group-019.md#canonical-1231113013303201-3233222000300212-2211131130013302-2322211003010221-0320233103301130-2102232111212102-2321212333113100-3130203302022212)
- [https.tls_parameters.tls_certificates](resources--http_loadbalancer--reference--group-019.md#canonical-2333201003023222-1221313103330132-3312310111023023-3001030030313321-2220312323120110-2133112333103203-1301231023103233-1333103021220231)
- [https.tls_parameters.tls_certificates.private_key](resources--http_loadbalancer--reference--group-019.md#canonical-2012002302202202-2020122200300213-3021321323122130-1322013120010111-0102301303322223-1233213233122333-3123201031333301-3110202011013000)
- https.tls_parameters.tls_certificates.private_key.blindfold_secret_info

<a id="canonical-2222102202303310-3013113203331333-0330133331202020-0010122023130210-1320231131210132-0101013231230012-3111220322002202-3330111200013100"></a>

Type: `"object"`. single nested block, Optional.

BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-2121100133023100-3110321220001003-3022031031013111-1122023001213110-3033002323200021-2002103210133322-0233301300330020-1300120321310313"></a>

## Direct properties — blindfold_secret_info / 022023122200 / 3

<a id="canonical-1330321003211302-3233033103213131-2200311310003203-0131033203130001-3233313123112131-1300103031210123-1331330003313311-3332102132213130"></a>

<a id="canonical-2212112303232333-0131231222202031-0032011010311121-0131103223031001-3002002013021320-3033330231311001-2103122113333001-3001000012311122"></a>

## decryption_provider property — blindfold_secret_info / 022023122200 / 4

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

<a id="canonical-1233033121221113-1210103031333132-3013023320032200-1003201221013310-1220131023031320-0002313122031331-3203233023030311-1112122221211300"></a>

<a id="canonical-3023312132021313-3122212212110101-2233302133110111-1122011100311200-1103100013333120-1020230001332113-3322113231102301-1222131233321030"></a>

## location property — blindfold_secret_info / 022023122200 / 5

Type: `"string"`. Optional, Sensitive.

Location is the URI\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Upstream description:

Location is the URI\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-1233123331103103-2001030332101203-3033230023301013-0100232030132100-3233303201210012-0313221112310010-3120000221000212-0302013231131013"></a>

<a id="canonical-0202222101123312-2022333031020323-3002021000210220-2323111130203120-2101201013121130-1320022321222331-0030320203133312-0000201101003211"></a>

## store_provider property — blindfold_secret_info / 022023122200 / 6

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

Upstream description:

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

<a id="canonical-2312231212013332-1231123321312110-3100112101221111-2020311132031123-3111110211330123-3330221303333011-3032022203020202-1003112310113232"></a>

## Next pages — blindfold_secret_info / 022023122200 / 7

- [https.tls_parameters.tls_certificates.private_key](resources--http_loadbalancer--reference--group-019.md#canonical-2012002302202202-2020122200300213-3021321323122130-1322013120010111-0102301303322223-1233213233122333-3123201031333301-3110202011013000)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-1211202021023200-3331130303131220-3300220332121300-0233010100332222-2012303203323012-0010212212001101-0120013221322220-2101311321110123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1130003131310033-2003130330020021-3323300012210032-0130102123131002-1032002212332221-2231013331113312-0123022133211011-1010002213320222"></a>

## https.tls_parameters.tls_certificates.private_key.clear_secret_info — clear_secret_info / 123112210211 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [https](resources--http_loadbalancer--reference--group-019.md#canonical-0321232101200210-0011302000022301-2002321223210110-0303132310312201-0222311023330031-3231011200231021-2012211321122303-0001332120000102)
- [https.tls_parameters](resources--http_loadbalancer--reference--group-019.md#canonical-1231113013303201-3233222000300212-2211131130013302-2322211003010221-0320233103301130-2102232111212102-2321212333113100-3130203302022212)
- [https.tls_parameters.tls_certificates](resources--http_loadbalancer--reference--group-019.md#canonical-2333201003023222-1221313103330132-3312310111023023-3001030030313321-2220312323120110-2133112333103203-1301231023103233-1333103021220231)
- [https.tls_parameters.tls_certificates.private_key](resources--http_loadbalancer--reference--group-019.md#canonical-2012002302202202-2020122200300213-3021321323122130-1322013120010111-0102301303322223-1233213233122333-3123201031333301-3110202011013000)
- https.tls_parameters.tls_certificates.private_key.clear_secret_info

<a id="canonical-3123133212313302-1230231030231300-1003231202333100-0012002313302231-2202303131131232-0133321303231133-1210123030030003-1301333011100022"></a>

Type: `"object"`. single nested block, Optional.

ClearSecretInfoType specifies information about the Secret that is not encrypted.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-3021120112102131-1302233323012222-0233001232231011-1323123323023120-0231111110322312-2100133031102120-0321231002220002-2013131201012000"></a>

## Direct properties — clear_secret_info / 123112210211 / 3

<a id="canonical-2222310232111002-1001113012230220-1223003300001212-2221303200302023-2033332013332311-3021230101021331-2000013131323133-0000021133021233"></a>

<a id="canonical-1203220002001111-2133323233323212-3213002323212010-0031311001310013-0320211023023203-0211130302320303-1210231100331021-2110130023322312"></a>

## provider_ref property — clear_secret_info / 123112210211 / 4

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-2120020221330001-1101321231110023-1012133330133301-0103222123112300-2332103202220001-0203303333333333-3330102233102331-0133320210033303"></a>

<a id="canonical-2032131130002021-3232301312021201-0131020231013233-1100202333101103-3333011032101233-2320023120120313-1102130202231020-2110002222312201"></a>

## URL property — clear_secret_info / 123112210211 / 5

Type: `"string"`. Optional, Sensitive.

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded base64 format. When asked for this secret, caller will GET Secret bytes after
base64 decoding.

Upstream description:

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded base64 format. When asked for this secret, caller will GET Secret bytes after
base64 decoding.

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

<a id="canonical-1120211120001310-2300012132111212-1321010312010203-1113210032022123-0113100201110311-3331231233301010-0312302112203231-3301110132303303"></a>

## Next pages — clear_secret_info / 123112210211 / 6

- [https.tls_parameters.tls_certificates.private_key](resources--http_loadbalancer--reference--group-019.md#canonical-2012002302202202-2020122200300213-3021321323122130-1322013120010111-0102301303322223-1233213233122333-3123201031333301-3110202011013000)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-1233033031232212-2012200221121213-2333013220123312-1032011003203332-2330130201010101-1103323120112201-2030313103301202-1202301301322103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0212223300032131-2131131010030010-0013131320220322-0000313202132201-2001001132033023-2320221113103123-2110323323010330-2200133213012233"></a>

## https.tls_parameters.tls_certificates.use_system_defaults — use_system_defaults / 213210021233 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [https](resources--http_loadbalancer--reference--group-019.md#canonical-0321232101200210-0011302000022301-2002321223210110-0303132310312201-0222311023330031-3231011200231021-2012211321122303-0001332120000102)
- [https.tls_parameters](resources--http_loadbalancer--reference--group-019.md#canonical-1231113013303201-3233222000300212-2211131130013302-2322211003010221-0320233103301130-2102232111212102-2321212333113100-3130203302022212)
- [https.tls_parameters.tls_certificates](resources--http_loadbalancer--reference--group-019.md#canonical-2333201003023222-1221313103330132-3312310111023023-3001030030313321-2220312323120110-2133112333103203-1301231023103233-1333103021220231)
- https.tls_parameters.tls_certificates.use_system_defaults

<a id="canonical-2303201310310112-0220203301130313-3130311303212213-2100213201132333-3131301020111101-0123031313311110-1330330220011222-0222223112132213"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for use system defaults.

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
use_system_defaults = {}
```

<a id="canonical-0000323020330012-1023330013030032-2121203102110200-0001100301223103-2102002103100102-2103310331130212-3113123102313322-3010213013213121"></a>

## Direct properties — use_system_defaults / 213210021233 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0021012110211020-1113010113131202-1033110312120312-1213123212122111-2131201003120213-3031312112011020-1323200123300322-1221032320331022"></a>

## Next pages — use_system_defaults / 213210021233 / 4

- [https.tls_parameters.tls_certificates](resources--http_loadbalancer--reference--group-019.md#canonical-2333201003023222-1221313103330132-3312310111023023-3001030030313321-2220312323120110-2133112333103203-1301231023103233-1333103021220231)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-1331301222322132-3310303011021332-0330022222232121-3232113210300302-3100323021101302-2323203211312232-2232012310223322-1333222331330000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1210312201021122-3101123122330010-2312303002100032-1201013100120010-1103300202312300-0213110220231331-0102112321130021-0313330323120330"></a>

## https.tls_parameters.tls_config — tls_config / 322213231311 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [https](resources--http_loadbalancer--reference--group-019.md#canonical-0321232101200210-0011302000022301-2002321223210110-0303132310312201-0222311023330031-3231011200231021-2012211321122303-0001332120000102)
- [https.tls_parameters](resources--http_loadbalancer--reference--group-019.md#canonical-1231113013303201-3233222000300212-2211131130013302-2322211003010221-0320233103301130-2102232111212102-2321212333113100-3130203302022212)
- https.tls_parameters.tls_config

<a id="canonical-2213301003123031-0130321302323013-3311012000232203-2230100300311322-2033012301001033-0033121320120122-0120301330010130-0101302103012100"></a>

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

<a id="canonical-0330302213133021-2330312102332111-1213031202020031-0132110331131110-3111201333331012-3103033020332322-1332231003100332-0111021122301213"></a>

## Direct properties — tls_config / 322213231311 / 3

- [custom_security](resources--http_loadbalancer--reference--group-019.md#canonical-2231332332000212-1131230233023220-0003301313231321-0020100131230112-3312121012300032-2131021102332211-0300110333110231-3202232032002221): complete subsection reference.

- [default_security](resources--http_loadbalancer--reference--group-019.md#canonical-1013310333012201-2000320121203322-2303222121322232-3013000302333330-3013132333000123-3030301213012232-2330301001000201-1011322002113213): complete subsection reference.

- [low_security](resources--http_loadbalancer--reference--group-019.md#canonical-3212011310330030-2002101322323302-0030022003030222-1222111202312132-0130202122103133-3211111312131222-3033100131100311-3013030233002023): complete subsection reference.

- [medium_security](resources--http_loadbalancer--reference--group-019.md#canonical-1101201203013130-0213013122203320-3230310031103130-2222130112000203-0312012200013133-1200120013233113-0000122211323131-0113321330233213): complete subsection reference.

<a id="canonical-2300311131212303-3212121133132020-1012203200023231-2233033321221130-1200010020322013-1332203301321132-2223221010300320-3232010320330223"></a>

## Next pages — tls_config / 322213231311 / 4

- [https.tls_parameters.tls_config.custom_security](resources--http_loadbalancer--reference--group-019.md#canonical-2231332332000212-1131230233023220-0003301313231321-0020100131230112-3312121012300032-2131021102332211-0300110333110231-3202232032002221)
- [https.tls_parameters.tls_config.default_security](resources--http_loadbalancer--reference--group-019.md#canonical-1013310333012201-2000320121203322-2303222121322232-3013000302333330-3013132333000123-3030301213012232-2330301001000201-1011322002113213)
- [https.tls_parameters.tls_config.low_security](resources--http_loadbalancer--reference--group-019.md#canonical-3212011310330030-2002101322323302-0030022003030222-1222111202312132-0130202122103133-3211111312131222-3033100131100311-3013030233002023)
- [https.tls_parameters.tls_config.medium_security](resources--http_loadbalancer--reference--group-019.md#canonical-1101201203013130-0213013122203320-3230310031103130-2222130112000203-0312012200013133-1200120013233113-0000122211323131-0113321330233213)
- [https.tls_parameters](resources--http_loadbalancer--reference--group-019.md#canonical-1231113013303201-3233222000300212-2211131130013302-2322211003010221-0320233103301130-2102232111212102-2321212333113100-3130203302022212)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-2231332332000212-1131230233023220-0003301313231321-0020100131230112-3312121012300032-2131021102332211-0300110333110231-3202232032002221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0122321230322321-2123021313020232-0332130102201203-1301023323222213-1231300120003022-0323321030022132-2121022122210122-0023122300130300"></a>

## https.tls_parameters.tls_config.custom_security — custom_security / 122221311213 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [https](resources--http_loadbalancer--reference--group-019.md#canonical-0321232101200210-0011302000022301-2002321223210110-0303132310312201-0222311023330031-3231011200231021-2012211321122303-0001332120000102)
- [https.tls_parameters](resources--http_loadbalancer--reference--group-019.md#canonical-1231113013303201-3233222000300212-2211131130013302-2322211003010221-0320233103301130-2102232111212102-2321212333113100-3130203302022212)
- [https.tls_parameters.tls_config](resources--http_loadbalancer--reference--group-019.md#canonical-1331301222322132-3310303011021332-0330022222232121-3232113210300302-3100323021101302-2323203211312232-2232012310223322-1333222331330000)
- https.tls_parameters.tls_config.custom_security

<a id="canonical-0210323020112013-2101101000100013-2232013332331201-3010020000103012-1222231230320130-0102201010301131-3012303023011123-2230232323322201"></a>

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

<a id="canonical-1021010133210320-2131012022301011-0102333013301123-0311231120102303-3010003031120132-1003020321033312-1030101023002101-1032200111131113"></a>

## Direct properties — custom_security / 122221311213 / 3

<a id="canonical-3330312120100202-3302032321032313-2123111232112303-1002033221302200-2221212101010132-2030032310321212-2320233000211221-0032323022332312"></a>

<a id="canonical-1331303320031232-0211102103222020-1322111322121201-3032330103121222-1303300212302211-2330133012012100-1221113311322001-2232323310230312"></a>

## cipher_suites property — custom_security / 122221311213 / 4

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

<a id="canonical-1330311010211003-0033012023003220-3022130201001323-2031132033210212-0200313333101030-2123100133112111-1102221331033130-2323133323112301"></a>

<a id="canonical-1000220023131121-0113031113210100-3013213021111003-3201201322102232-0111312022131300-0101030231313102-0321131022131321-2201100220213331"></a>

## max_version property — custom_security / 122221311213 / 5

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

<a id="canonical-0113232120230013-2023122230132020-3330112120301021-3110323303332323-3221311013331011-3232113300012321-1002013101113033-0332031120213020"></a>

<a id="canonical-1101322000000111-2133002300311323-0021220201101331-2203131233002021-1220333202101033-1010232330333133-1213030001102331-0300312220000320"></a>

## min_version property — custom_security / 122221311213 / 6

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

<a id="canonical-3032030122023130-3130002322210320-1312012300020220-1001110212123110-3121012021101333-3003210322030233-2111330203321222-0330203322100111"></a>

## Next pages — custom_security / 122221311213 / 7

- [https.tls_parameters.tls_config](resources--http_loadbalancer--reference--group-019.md#canonical-1331301222322132-3310303011021332-0330022222232121-3232113210300302-3100323021101302-2323203211312232-2232012310223322-1333222331330000)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-1013310333012201-2000320121203322-2303222121322232-3013000302333330-3013132333000123-3030301213012232-2330301001000201-1011322002113213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2112101013102301-1113332212022221-1333133320111003-2220320202231303-2001330221133230-2300332233133231-1211002203131013-0301322311102310"></a>

## https.tls_parameters.tls_config.default_security — default_security / 312002001110 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [https](resources--http_loadbalancer--reference--group-019.md#canonical-0321232101200210-0011302000022301-2002321223210110-0303132310312201-0222311023330031-3231011200231021-2012211321122303-0001332120000102)
- [https.tls_parameters](resources--http_loadbalancer--reference--group-019.md#canonical-1231113013303201-3233222000300212-2211131130013302-2322211003010221-0320233103301130-2102232111212102-2321212333113100-3130203302022212)
- [https.tls_parameters.tls_config](resources--http_loadbalancer--reference--group-019.md#canonical-1331301222322132-3310303011021332-0330022222232121-3232113210300302-3100323021101302-2323203211312232-2232012310223322-1333222331330000)
- https.tls_parameters.tls_config.default_security

<a id="canonical-3002000320000211-2203212130321103-1132203001131123-3112131110030012-3100011202002200-0232033212021331-3220103323331113-1033311330303210"></a>

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

<a id="canonical-0331202211010023-2123131101213133-0003212331233313-3330012130232200-2030212003021321-1113102032133300-2121201213333001-2223322330212120"></a>

## Direct properties — default_security / 312002001110 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0103020211222302-0032121310221321-0122230332200103-1331001110133322-3123000002011113-3203033333231103-2222212032221132-3330111030100301"></a>

## Next pages — default_security / 312002001110 / 4

- [https.tls_parameters.tls_config](resources--http_loadbalancer--reference--group-019.md#canonical-1331301222322132-3310303011021332-0330022222232121-3232113210300302-3100323021101302-2323203211312232-2232012310223322-1333222331330000)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-3212011310330030-2002101322323302-0030022003030222-1222111202312132-0130202122103133-3211111312131222-3033100131100311-3013030233002023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0203021210022002-1222311120110232-2030220011121212-3030102103101333-1203313021311010-3112210021213012-1301210020232030-1123200031012213"></a>

## https.tls_parameters.tls_config.low_security — low_security / 230323203302 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [https](resources--http_loadbalancer--reference--group-019.md#canonical-0321232101200210-0011302000022301-2002321223210110-0303132310312201-0222311023330031-3231011200231021-2012211321122303-0001332120000102)
- [https.tls_parameters](resources--http_loadbalancer--reference--group-019.md#canonical-1231113013303201-3233222000300212-2211131130013302-2322211003010221-0320233103301130-2102232111212102-2321212333113100-3130203302022212)
- [https.tls_parameters.tls_config](resources--http_loadbalancer--reference--group-019.md#canonical-1331301222322132-3310303011021332-0330022222232121-3232113210300302-3100323021101302-2323203211312232-2232012310223322-1333222331330000)
- https.tls_parameters.tls_config.low_security

<a id="canonical-1101312020213220-3310231023310023-3201323301210000-1023120131201012-1302111113230103-0123211203023122-2122331310132221-3100032322213333"></a>

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

<a id="canonical-3121201030001201-0012311302133300-3100233221011112-1313000233002310-0230221213212121-0103003113121313-1303222121320113-3021312221022212"></a>

## Direct properties — low_security / 230323203302 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0132311120301302-0323213200100313-2312200223031113-3322212121300222-3311331003001312-2103133322312023-2323031232311110-0033321030203213"></a>

## Next pages — low_security / 230323203302 / 4

- [https.tls_parameters.tls_config](resources--http_loadbalancer--reference--group-019.md#canonical-1331301222322132-3310303011021332-0330022222232121-3232113210300302-3100323021101302-2323203211312232-2232012310223322-1333222331330000)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-1101201203013130-0213013122203320-3230310031103130-2222130112000203-0312012200013133-1200120013233113-0000122211323131-0113321330233213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1123033001030212-1221002100330311-3003320121133322-3033200032103033-1210303120232223-3011201123022022-3121321012011230-3101033331332232"></a>

## https.tls_parameters.tls_config.medium_security — medium_security / 112230303223 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [https](resources--http_loadbalancer--reference--group-019.md#canonical-0321232101200210-0011302000022301-2002321223210110-0303132310312201-0222311023330031-3231011200231021-2012211321122303-0001332120000102)
- [https.tls_parameters](resources--http_loadbalancer--reference--group-019.md#canonical-1231113013303201-3233222000300212-2211131130013302-2322211003010221-0320233103301130-2102232111212102-2321212333113100-3130203302022212)
- [https.tls_parameters.tls_config](resources--http_loadbalancer--reference--group-019.md#canonical-1331301222322132-3310303011021332-0330022222232121-3232113210300302-3100323021101302-2323203211312232-2232012310223322-1333222331330000)
- https.tls_parameters.tls_config.medium_security

<a id="canonical-2001320201021031-0312310022032331-2310001011010332-3122300323131112-0223322130313023-3131321201021213-2021202111101230-1210302130001202"></a>

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

<a id="canonical-2332101300013320-2110032322112322-2012211022310302-2000013202223333-3331120110022032-3123022101131113-3023233201120103-2002232333311312"></a>

## Direct properties — medium_security / 112230303223 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0002333010320330-2202222103230131-0011313031323330-2332330032010310-1313012021331122-1003121210111111-3123033010000230-1202000323021131"></a>

## Next pages — medium_security / 112230303223 / 4

- [https.tls_parameters.tls_config](resources--http_loadbalancer--reference--group-019.md#canonical-1331301222322132-3310303011021332-0330022222232121-3232113210300302-3100323021101302-2323203211312232-2232012310223322-1333222331330000)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-1321332101331311-1023112313030210-2312030203110123-1202033022020030-0101101100102110-2022201111233012-3210333032232312-3311120222312230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1030202023210321-1302011312202023-2231333100222220-0021333002231013-2033233120112203-3023033133330101-3232010302230301-2133230003131200"></a>

## https.tls_parameters.use_mtls — use_mtls / 230030021102 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [https](resources--http_loadbalancer--reference--group-019.md#canonical-0321232101200210-0011302000022301-2002321223210110-0303132310312201-0222311023330031-3231011200231021-2012211321122303-0001332120000102)
- [https.tls_parameters](resources--http_loadbalancer--reference--group-019.md#canonical-1231113013303201-3233222000300212-2211131130013302-2322211003010221-0320233103301130-2102232111212102-2321212333113100-3130203302022212)
- https.tls_parameters.use_mtls

<a id="canonical-3022100120320223-3301233221233120-1330113123121322-0233233013022122-2111132030111133-2120321302200233-3020130302311230-0010000300000202"></a>

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

<a id="canonical-3103033200310011-1033003331331221-3020333022021221-3313123231002321-3101013113200102-1011131123213310-1033110113020100-0132300300330232"></a>

## Direct properties — use_mtls / 230030021102 / 3

<a id="canonical-0031002303212110-3011330210123311-3233330331212103-3023333120003131-0112021033111132-2201320121131003-1303120320122221-1101220310203330"></a>

<a id="canonical-1023012231203213-3130221332200133-1231200320301213-1030323011022312-3331213001223221-2200121103102321-2323200223011130-0312000233110210"></a>

## client_certificate_optional property — use_mtls / 230030021102 / 4

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

- [crl](resources--http_loadbalancer--reference--group-019.md#canonical-0222010121130210-2230331320103033-1003331213102030-1231111322120121-1013331102203000-3000212313333123-3110020103321020-0000222200021202): complete subsection reference.

- [no_crl](resources--http_loadbalancer--reference--group-019.md#canonical-1030111002012120-3333122220103311-3320230001320012-1232020311120002-0232303302022021-1113121110000110-0020100233000213-2102200103102111): complete subsection reference.

- [trusted_ca](resources--http_loadbalancer--reference--group-019.md#canonical-2322231221033132-3033011200232320-1122002210223211-3311212321020032-2112201023110101-2023322002222213-0110121233231212-1322321203331110): complete subsection reference.

<a id="canonical-2010320113301112-1322313133011222-2233233032232102-3110322013212011-3300113133332013-0100213300221332-2321100130213023-3212112123120313"></a>

<a id="canonical-1113103222312012-1113001202002022-2300011321302322-2312113202130021-3001112012231133-3302131223121330-3211120130100322-0022032003212011"></a>

## trusted_ca_url property — use_mtls / 230030021102 / 5

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

- [xfcc_disabled](resources--http_loadbalancer--reference--group-019.md#canonical-2322002130002320-1330120021101020-1200130232312020-0133311022301011-3202331222300110-3322232232122000-1021003111002231-2232030010220020): complete subsection reference.

- [xfcc_options](resources--http_loadbalancer--reference--group-019.md#canonical-0133133030323002-3012021101322010-0220200122033323-3222120033110132-3121211121220001-0131113203131111-0020130103231221-1131302303211333): complete subsection reference.

<a id="canonical-2111112232213121-2011222312001200-3302222021013121-1113220031300032-1321302212032003-0002312020111131-0031231223121121-2301031023202231"></a>

## Next pages — use_mtls / 230030021102 / 6

- [https.tls_parameters.use_mtls.crl](resources--http_loadbalancer--reference--group-019.md#canonical-0222010121130210-2230331320103033-1003331213102030-1231111322120121-1013331102203000-3000212313333123-3110020103321020-0000222200021202)
- [https.tls_parameters.use_mtls.no_crl](resources--http_loadbalancer--reference--group-019.md#canonical-1030111002012120-3333122220103311-3320230001320012-1232020311120002-0232303302022021-1113121110000110-0020100233000213-2102200103102111)
- [https.tls_parameters.use_mtls.trusted_ca](resources--http_loadbalancer--reference--group-019.md#canonical-2322231221033132-3033011200232320-1122002210223211-3311212321020032-2112201023110101-2023322002222213-0110121233231212-1322321203331110)
- [https.tls_parameters.use_mtls.xfcc_disabled](resources--http_loadbalancer--reference--group-019.md#canonical-2322002130002320-1330120021101020-1200130232312020-0133311022301011-3202331222300110-3322232232122000-1021003111002231-2232030010220020)
- [https.tls_parameters.use_mtls.xfcc_options](resources--http_loadbalancer--reference--group-019.md#canonical-0133133030323002-3012021101322010-0220200122033323-3222120033110132-3121211121220001-0131113203131111-0020130103231221-1131302303211333)
- [https.tls_parameters](resources--http_loadbalancer--reference--group-019.md#canonical-1231113013303201-3233222000300212-2211131130013302-2322211003010221-0320233103301130-2102232111212102-2321212333113100-3130203302022212)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-0222010121130210-2230331320103033-1003331213102030-1231111322120121-1013331102203000-3000212313333123-3110020103321020-0000222200021202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0302122203021023-0010322301021102-3303011001110013-2103011002312131-2210003120033302-2131113310230033-1303101023330123-1112002022301110"></a>

## https.tls_parameters.use_mtls.crl — crl / 112123113210 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [https](resources--http_loadbalancer--reference--group-019.md#canonical-0321232101200210-0011302000022301-2002321223210110-0303132310312201-0222311023330031-3231011200231021-2012211321122303-0001332120000102)
- [https.tls_parameters](resources--http_loadbalancer--reference--group-019.md#canonical-1231113013303201-3233222000300212-2211131130013302-2322211003010221-0320233103301130-2102232111212102-2321212333113100-3130203302022212)
- [https.tls_parameters.use_mtls](resources--http_loadbalancer--reference--group-019.md#canonical-1321332101331311-1023112313030210-2312030203110123-1202033022020030-0101101100102110-2022201111233012-3210333032232312-3311120222312230)
- https.tls_parameters.use_mtls.crl

<a id="canonical-2331030020221031-1231123300213211-1012322130331331-0101122012120301-2003231113020203-3221230223233303-3332100111121312-3102300203320223"></a>

Type: `"object"`. single nested block, Optional.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

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
crl {
  # Configure direct properties listed below.
}
```

<a id="canonical-2303011120122132-3330113223011031-2332333012130030-2221002111023031-0221220113213233-3231000133022101-0011303203011311-3100021223113212"></a>

## Direct properties — crl / 112123113210 / 3

<a id="canonical-2212122201022003-2210020203120233-0132023213233133-1313101202002320-3203031031321323-1130200013021203-1100322100022103-0123123031112023"></a>

<a id="canonical-1222312011220331-0130210021200112-1130013213322021-1120123101102213-3332030101131300-3302102001130233-3022201232202230-1011313023201021"></a>

## name property — crl / 112123113210 / 4

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

<a id="canonical-2011322131222112-3321132200030332-1312322011110303-3032022232111231-0323213130102233-3130302131213123-3313300012332021-1001112223310022"></a>

<a id="canonical-2303322013312032-0010011221300003-1121213230232203-0013302312222301-0001030200202313-2121232101220003-2113321003021023-3221120331313320"></a>

## namespace property — crl / 112123113210 / 5

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

<a id="canonical-3332010102322023-1130112031100211-3133131210210121-2032033211200310-3032212313032231-0201030223320303-3003023312331123-0002301120311201"></a>

<a id="canonical-0112031133131210-0120131333233220-0013101120002230-1031330320313330-3031120321102021-2312221213110112-3303032302110311-1333303030310110"></a>

## tenant property — crl / 112123113210 / 6

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

<a id="canonical-0310121210320320-2312000010030002-3111133020001101-2112110011333122-0001332333211100-3133211231032332-3230030331330000-1003002132030321"></a>

## Next pages — crl / 112123113210 / 7

- [https.tls_parameters.use_mtls](resources--http_loadbalancer--reference--group-019.md#canonical-1321332101331311-1023112313030210-2312030203110123-1202033022020030-0101101100102110-2022201111233012-3210333032232312-3311120222312230)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-1030111002012120-3333122220103311-3320230001320012-1232020311120002-0232303302022021-1113121110000110-0020100233000213-2102200103102111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1000012313203211-3311022000102313-3330113333200022-0013220113310322-2111132333303022-3301130133312121-3202111122302301-2210320032000013"></a>

## https.tls_parameters.use_mtls.no_crl — no_crl / 002302130223 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [https](resources--http_loadbalancer--reference--group-019.md#canonical-0321232101200210-0011302000022301-2002321223210110-0303132310312201-0222311023330031-3231011200231021-2012211321122303-0001332120000102)
- [https.tls_parameters](resources--http_loadbalancer--reference--group-019.md#canonical-1231113013303201-3233222000300212-2211131130013302-2322211003010221-0320233103301130-2102232111212102-2321212333113100-3130203302022212)
- [https.tls_parameters.use_mtls](resources--http_loadbalancer--reference--group-019.md#canonical-1321332101331311-1023112313030210-2312030203110123-1202033022020030-0101101100102110-2022201111233012-3210333032232312-3311120222312230)
- https.tls_parameters.use_mtls.no_crl

<a id="canonical-1230002332301133-0313013212112020-0323120312100111-0310012323202031-1112132103230200-1321203313230310-0213113231131302-1231323322211312"></a>

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
no_crl = {}
```

<a id="canonical-3000033120121212-2201113312111123-0010023103331331-0313313002103023-2303020331301210-3111012131011232-0133230303010131-1123312203012111"></a>

## Direct properties — no_crl / 002302130223 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1302023201203120-0311321010112301-3133020301100122-0121221030323031-2130300330133220-3133102131122312-0102200122132031-2112112231031120"></a>

## Next pages — no_crl / 002302130223 / 4

- [https.tls_parameters.use_mtls](resources--http_loadbalancer--reference--group-019.md#canonical-1321332101331311-1023112313030210-2312030203110123-1202033022020030-0101101100102110-2022201111233012-3210333032232312-3311120222312230)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-2322231221033132-3033011200232320-1122002210223211-3311212321020032-2112201023110101-2023322002222213-0110121233231212-1322321203331110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0121212130323023-0101122011111203-3321231021310302-0020303121102233-1021001202033022-0120331301101131-2330211002201122-3220320130202313"></a>

## https.tls_parameters.use_mtls.trusted_ca — trusted_ca / 212320010010 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [https](resources--http_loadbalancer--reference--group-019.md#canonical-0321232101200210-0011302000022301-2002321223210110-0303132310312201-0222311023330031-3231011200231021-2012211321122303-0001332120000102)
- [https.tls_parameters](resources--http_loadbalancer--reference--group-019.md#canonical-1231113013303201-3233222000300212-2211131130013302-2322211003010221-0320233103301130-2102232111212102-2321212333113100-3130203302022212)
- [https.tls_parameters.use_mtls](resources--http_loadbalancer--reference--group-019.md#canonical-1321332101331311-1023112313030210-2312030203110123-1202033022020030-0101101100102110-2022201111233012-3210333032232312-3311120222312230)
- https.tls_parameters.use_mtls.trusted_ca

<a id="canonical-2122333012202002-3210202000122111-2111210002202102-1002320330333210-2200110010013233-2310133101231233-0221320203011313-2221020100310333"></a>

Type: `"object"`. single nested block, Optional.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

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
trusted_ca {
  # Configure direct properties listed below.
}
```

<a id="canonical-0320210211100310-0111100032003203-0331133020302300-3323012030223033-2230112121030323-2232031030030130-2310002100023013-1301110320321130"></a>

## Direct properties — trusted_ca / 212320010010 / 3

<a id="canonical-1223133011301133-0120233310020313-0122030003032231-1123032031101223-0100001212120012-1103311333101000-1310020032121123-1331131122303102"></a>

<a id="canonical-1122323132322330-3013112213022010-2321111233122113-3210120223203223-0301311330121031-3022001310010323-2303001332100010-2023322230223011"></a>

## name property — trusted_ca / 212320010010 / 4

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

<a id="canonical-1211333230221201-2320332332322311-2331010033133022-3122020323223032-0213013030132032-0210122111303321-2120023203210002-2010120011331321"></a>

<a id="canonical-2032002210012021-2030222012330201-3001332023112122-0011202012212111-1010212230312332-0023131012311020-1132222033113302-2122102132233110"></a>

## namespace property — trusted_ca / 212320010010 / 5

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

<a id="canonical-3120200012201003-1112212113021130-3212303133303103-0333031112000223-2002212020300233-0102323200212233-0323212300300113-2310320122310302"></a>

<a id="canonical-3121230000120021-2200331210000231-3221011300122122-1132120022220220-0010213223012032-0333231320310321-2221000232223233-0220310021320220"></a>

## tenant property — trusted_ca / 212320010010 / 6

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

<a id="canonical-1300110120102033-0323330303303112-3302321003001130-0203301002012131-2203213200112102-3303133111100000-2031232320121132-1301231030312022"></a>

## Next pages — trusted_ca / 212320010010 / 7

- [https.tls_parameters.use_mtls](resources--http_loadbalancer--reference--group-019.md#canonical-1321332101331311-1023112313030210-2312030203110123-1202033022020030-0101101100102110-2022201111233012-3210333032232312-3311120222312230)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-2322002130002320-1330120021101020-1200130232312020-0133311022301011-3202331222300110-3322232232122000-1021003111002231-2232030010220020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0302123001113332-3230001222023103-3110113000322213-1231000013133322-1133231230022113-1112003023032011-2233202112231200-0111001032333333"></a>

## https.tls_parameters.use_mtls.xfcc_disabled — xfcc_disabled / 120001233100 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [https](resources--http_loadbalancer--reference--group-019.md#canonical-0321232101200210-0011302000022301-2002321223210110-0303132310312201-0222311023330031-3231011200231021-2012211321122303-0001332120000102)
- [https.tls_parameters](resources--http_loadbalancer--reference--group-019.md#canonical-1231113013303201-3233222000300212-2211131130013302-2322211003010221-0320233103301130-2102232111212102-2321212333113100-3130203302022212)
- [https.tls_parameters.use_mtls](resources--http_loadbalancer--reference--group-019.md#canonical-1321332101331311-1023112313030210-2312030203110123-1202033022020030-0101101100102110-2022201111233012-3210333032232312-3311120222312230)
- https.tls_parameters.use_mtls.xfcc_disabled

<a id="canonical-2001330313213112-2003121033010131-1233202331001132-1031030001011100-3332012002200003-0310121321203212-1121303033022021-1113231103133003"></a>

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
xfcc_disabled = {}
```

<a id="canonical-0120031033322300-2303200222121212-3131323022120020-1331112113102312-0230103121221331-1220112311201313-1322131023133300-1100331121101310"></a>

## Direct properties — xfcc_disabled / 120001233100 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2132321011232113-2310213033331113-2113200121221322-3330101333132332-1030332110102312-3313111302313032-2113122012010220-1213033223010213"></a>

## Next pages — xfcc_disabled / 120001233100 / 4

- [https.tls_parameters.use_mtls](resources--http_loadbalancer--reference--group-019.md#canonical-1321332101331311-1023112313030210-2312030203110123-1202033022020030-0101101100102110-2022201111233012-3210333032232312-3311120222312230)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-0133133030323002-3012021101322010-0220200122033323-3222120033110132-3121211121220001-0131113203131111-0020130103231221-1131302303211333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->
