---
page_title: "xcsh_workload reference"
subcategory: "Container"
description: "Complete grouped canonical reference for xcsh_workload reference."
---

# xcsh_workload reference

<a id="canonical-1320030113012233-1022032312211012-1001221123023230-2222203332110111-0020322321131123-0203312222011111-2231002231202012-3102013312011000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.port.http_loadbalancer.http` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-006.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-006.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-009.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.port](resources--workload--reference--group-012.md#canonical-3013031222022030-2021222233122213-2220022232301003-3102220000332302-0311213311310132-3133321213021123-0133300222212330-1122131003023301)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-012.md#canonical-0101000213001300-0222020002101313-1020132032013321-2231011111202333-2122231103331210-3202123132002020-0313332003123002-3101311011132330)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.http

<a id="canonical-0031323223002202-2331221023220310-3230101122010002-2013231323222002-1021232213230322-2200232013201010-3133131323122233-0032231203301330"></a>

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

<a id="canonical-3102210300110113-3110031301220222-3122213013333023-3223021221120030-1323332130111300-0332202021001113-1333110110113130-0210310313000233"></a>

### Direct properties for `service.advertise_options.advertise_on_public.port.http_loadbalancer.http`

<a id="canonical-2022110331032023-3322203310312202-2011220323021331-1020201201222113-1121032320000103-3321232331313323-2210133233203231-2021120220203323"></a>

#### `service.advertise_options.advertise_on_public.port.http_loadbalancer.http.dns_volterra_managed` property

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

<a id="canonical-3200023213213032-0313332222222111-1301203223203330-2232103011302323-0100333330233113-3001021100200031-2203233232232113-1113220233220023"></a>

<a id="canonical-0302122232131130-3202132220012031-1002330122030102-0312211002112032-1101232323310232-3323330103331003-1022301120021200-3322211012322301"></a>

#### `service.advertise_options.advertise_on_public.port.http_loadbalancer.http.port` property

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

<a id="canonical-0323300120311121-1303132220103203-1030231320231002-3203300313323002-0233321313201111-1130010102200120-0003211011003233-0012011203120213"></a>

<a id="canonical-2332200110133030-0232132332203121-0032213221002032-1112023033003221-0102000211321030-2303030120003003-0121120001313233-1001010021232331"></a>

#### `service.advertise_options.advertise_on_public.port.http_loadbalancer.http.port_ranges` property

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

<a id="canonical-2221000320213131-2001211030022032-2023131131113110-1000133011113313-1003112332120000-2102122231000321-0211212202023222-1022013120231111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.port.http_loadbalancer.https` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-006.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-006.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-009.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.port](resources--workload--reference--group-012.md#canonical-3013031222022030-2021222233122213-2220022232301003-3102220000332302-0311213311310132-3133321213021123-0133300222212330-1122131003023301)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-012.md#canonical-0101000213001300-0222020002101313-1020132032013321-2231011111202333-2122231103331210-3202123132002020-0313332003123002-3101311011132330)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https

<a id="canonical-2323030002233233-0202303330123212-1312030120031331-2212130310220101-1310013333033032-3213333002010230-0330122021123221-1323113333302131"></a>

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

<a id="canonical-1300312333100200-1320332133000212-3311313023132012-1312212000000332-1033213322002102-0022130121310302-3230022303101121-3132210212113123"></a>

### Direct properties for `service.advertise_options.advertise_on_public.port.http_loadbalancer.https`

<a id="canonical-0212012023333000-1310321110323210-1100103330131330-3322103010030012-2223211311312121-2020202133130022-1012200222322122-2103312233032202"></a>

#### `service.advertise_options.advertise_on_public.port.http_loadbalancer.https.add_hsts` property

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

<a id="canonical-0103233010332103-3130122311312311-0211002100312231-0113122232233110-0012103312322331-0131001121300130-2112230130032232-1210131211110333"></a>

<a id="canonical-3220123000033033-1320233311322003-0111202132331223-2312122123322100-3021033132130031-3132013022320132-1212230022030323-1303220301223222"></a>

#### `service.advertise_options.advertise_on_public.port.http_loadbalancer.https.append_server_name` property

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

- [coalescing_options](resources--workload--reference--group-013.md#canonical-3131021203100301-0323303103132212-0030113113022103-2110220112133323-3002212031312033-2302323113223221-2022010011303303-1311203212022220): complete subsection reference.

<a id="canonical-2331333301130023-0202102221102332-1112112130200011-2233113133230213-1220133210120021-3331233030022013-1002211022232210-2201100120301313"></a>

<a id="canonical-1211022330113313-1322013002300202-3302121123220320-1011222002101131-0102221130110313-1112303322022031-3211020320331203-1130221100212111"></a>

#### `service.advertise_options.advertise_on_public.port.http_loadbalancer.https.connection_idle_timeout` property

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

- [default_header](resources--workload--reference--group-013.md#canonical-3321310121130013-0222122313201310-2213231322223211-3200100220211010-1310022022201332-0101103120113210-0221322120201202-2300120220310301): complete subsection reference.

- [default_loadbalancer](resources--workload--reference--group-013.md#canonical-1320132313310100-3321223020113201-0130032123021103-1131232332230310-3230213202011331-3201311113003121-3211003013011113-1033213131023001): complete subsection reference.

- [disable_path_normalize](resources--workload--reference--group-013.md#canonical-2221311312002300-1300300121022222-0232321211132023-0311332020111332-0022132113301000-3032221303002122-1123313231221231-1200113120011331): complete subsection reference.

- [enable_path_normalize](resources--workload--reference--group-013.md#canonical-3233203303111321-3231322112033101-2221200310110223-1202313122212102-1223331231202221-2121122131101331-1010313202012022-2203102032221023): complete subsection reference.

- [http_protocol_options](resources--workload--reference--group-013.md#canonical-0113021231303233-2111320110111221-1302231000110011-2320311231321233-2223201132301001-2100032112123121-3000013033320312-1132003001210123): complete subsection reference.

<a id="canonical-3001021302323111-1121110212031203-2333221133332010-0123122232333320-0312103321323113-3323320012200332-3212000123333021-1323300222132132"></a>

<a id="canonical-3210310113120233-3023102223330312-2013112000322230-2020200233211320-0230010200303112-2230102133012023-3121220123221300-3302233032220100"></a>

#### `service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_redirect` property

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

- [non_default_loadbalancer](resources--workload--reference--group-013.md#canonical-0201012102332122-3021012133330000-1212133112031131-3000111023333013-1113210313110021-3303133302312331-3303212013301222-0032301130003232): complete subsection reference.

- [pass_through](resources--workload--reference--group-013.md#canonical-1331231121223122-3020333331231320-2222002002323013-0001220313222302-3223121012120332-1013121312110111-0010212110113013-0200301020223132): complete subsection reference.

<a id="canonical-1221312012021113-2113012111120323-0111122320323111-3322130301332232-2013001222033133-1100110230233313-0021003312103233-1231322111233210"></a>

<a id="canonical-2203301100203103-2100111331020122-1223013321331003-0223232110312122-2310120210030110-3023303133212231-2211032300221121-3033231330103313"></a>

#### `service.advertise_options.advertise_on_public.port.http_loadbalancer.https.port` property

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

<a id="canonical-0020010320031323-2320102133132133-3011112100031030-1331213323000212-0020023101211221-3331213010111021-3300000212221120-0021300031121300"></a>

<a id="canonical-1003010123131000-3232302321123130-1130002102011101-3211020121332320-1321113032221032-0111032203220332-3220332301313020-3102311110332302"></a>

#### `service.advertise_options.advertise_on_public.port.http_loadbalancer.https.port_ranges` property

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

<a id="canonical-1210323300033212-2202121030312121-3220130311313001-1222100013233030-0303112330030301-3330130203223223-3332121011013002-3230311102223213"></a>

<a id="canonical-0103202321303323-3221212100233033-0001302231223322-1203031003101021-0231032123031201-2001232333011032-0130230030001211-1133023031132211"></a>

#### `service.advertise_options.advertise_on_public.port.http_loadbalancer.https.server_name` property

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

- [tls_cert_params](resources--workload--reference--group-013.md#canonical-2021233301230331-3111000132213023-0320231030332320-2213232031303212-1323212201220200-0313001122300320-3212003130200223-2033202033223233): complete subsection reference.

- [tls_parameters](resources--workload--reference--group-013.md#canonical-1231003120133112-2013110130100211-2233121010211001-2121031322013232-2232220303133231-0300100233200310-0003012301120011-3102210111030310): complete subsection reference.

<a id="canonical-3131021203100301-0323303103132212-0030113113022103-2110220112133323-3002212031312033-2302323113223221-2022010011303303-1311203212022220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.port.http_loadbalancer.https.coalescing_options` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-006.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-006.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-009.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.port](resources--workload--reference--group-012.md#canonical-3013031222022030-2021222233122213-2220022232301003-3102220000332302-0311213311310132-3133321213021123-0133300222212330-1122131003023301)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-012.md#canonical-0101000213001300-0222020002101313-1020132032013321-2231011111202333-2122231103331210-3202123132002020-0313332003123002-3101311011132330)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https](resources--workload--reference--group-013.md#canonical-2221000320213131-2001211030022032-2023131131113110-1000133011113313-1003112332120000-2102122231000321-0211212202023222-1022013120231111)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https.coalescing_options

<a id="canonical-2122121130121203-0321200021221233-0223331132131030-1020322011310321-3033000311102311-3313110211130202-0002320202230111-0111212110301122"></a>

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

<a id="canonical-3231022031210221-2132223013011320-1332100023002100-3132322231330332-1020201013131302-1030113321210000-3330333333210213-0202102102033101"></a>

### Direct properties for `service.advertise_options.advertise_on_public.port.http_loadbalancer.https.coalescing_options`

- [default_coalescing](resources--workload--reference--group-013.md#canonical-3102301230021131-0233123330123231-3033332321113121-2100323030123023-1101030031021131-3311221233310211-0332320212031033-3032100121121121): complete subsection reference.

- [strict_coalescing](resources--workload--reference--group-013.md#canonical-2212331311020211-3011111311300012-3012210221213023-0103033200211213-2102032221201321-0123023212112023-3323000311303320-0100230011033011): complete subsection reference.

<a id="canonical-3102301230021131-0233123330123231-3033332321113121-2100323030123023-1101030031021131-3311221233310211-0332320212031033-3032100121121121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.port.http_loadbalancer.https.coalescing_options.default_coalescing` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-006.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-006.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-009.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.port](resources--workload--reference--group-012.md#canonical-3013031222022030-2021222233122213-2220022232301003-3102220000332302-0311213311310132-3133321213021123-0133300222212330-1122131003023301)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-012.md#canonical-0101000213001300-0222020002101313-1020132032013321-2231011111202333-2122231103331210-3202123132002020-0313332003123002-3101311011132330)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https](resources--workload--reference--group-013.md#canonical-2221000320213131-2001211030022032-2023131131113110-1000133011113313-1003112332120000-2102122231000321-0211212202023222-1022013120231111)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.coalescing_options](resources--workload--reference--group-013.md#canonical-3131021203100301-0323303103132212-0030113113022103-2110220112133323-3002212031312033-2302323113223221-2022010011303303-1311203212022220)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https.coalescing_options.default_coalescing

<a id="canonical-2101300123311302-2121001131332011-2323220003301222-2301120332003201-0233122230321310-0122300233333000-3132311301002231-2213130121033112"></a>

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

<a id="canonical-2212331311020211-3011111311300012-3012210221213023-0103033200211213-2102032221201321-0123023212112023-3323000311303320-0100230011033011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.port.http_loadbalancer.https.coalescing_options.strict_coalescing` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-006.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-006.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-009.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.port](resources--workload--reference--group-012.md#canonical-3013031222022030-2021222233122213-2220022232301003-3102220000332302-0311213311310132-3133321213021123-0133300222212330-1122131003023301)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-012.md#canonical-0101000213001300-0222020002101313-1020132032013321-2231011111202333-2122231103331210-3202123132002020-0313332003123002-3101311011132330)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https](resources--workload--reference--group-013.md#canonical-2221000320213131-2001211030022032-2023131131113110-1000133011113313-1003112332120000-2102122231000321-0211212202023222-1022013120231111)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.coalescing_options](resources--workload--reference--group-013.md#canonical-3131021203100301-0323303103132212-0030113113022103-2110220112133323-3002212031312033-2302323113223221-2022010011303303-1311203212022220)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https.coalescing_options.strict_coalescing

<a id="canonical-2030113112032200-2111122320102223-0213232300131302-1210211133203233-3132322122120133-0211330202323201-1032202132203212-2212022100022010"></a>

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

<a id="canonical-3321310121130013-0222122313201310-2213231322223211-3200100220211010-1310022022201332-0101103120113210-0221322120201202-2300120220310301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.port.http_loadbalancer.https.default_header` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-006.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-006.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-009.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.port](resources--workload--reference--group-012.md#canonical-3013031222022030-2021222233122213-2220022232301003-3102220000332302-0311213311310132-3133321213021123-0133300222212330-1122131003023301)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-012.md#canonical-0101000213001300-0222020002101313-1020132032013321-2231011111202333-2122231103331210-3202123132002020-0313332003123002-3101311011132330)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https](resources--workload--reference--group-013.md#canonical-2221000320213131-2001211030022032-2023131131113110-1000133011113313-1003112332120000-2102122231000321-0211212202023222-1022013120231111)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https.default_header

<a id="canonical-0013231333303232-2031100303331320-0302231322300032-0331100111103323-2101032013101213-2210022103023323-1012003332000313-1113132123331010"></a>

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

<a id="canonical-1320132313310100-3321223020113201-0130032123021103-1131232332230310-3230213202011331-3201311113003121-3211003013011113-1033213131023001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.port.http_loadbalancer.https.default_loadbalancer` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-006.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-006.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-009.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.port](resources--workload--reference--group-012.md#canonical-3013031222022030-2021222233122213-2220022232301003-3102220000332302-0311213311310132-3133321213021123-0133300222212330-1122131003023301)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-012.md#canonical-0101000213001300-0222020002101313-1020132032013321-2231011111202333-2122231103331210-3202123132002020-0313332003123002-3101311011132330)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https](resources--workload--reference--group-013.md#canonical-2221000320213131-2001211030022032-2023131131113110-1000133011113313-1003112332120000-2102122231000321-0211212202023222-1022013120231111)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https.default_loadbalancer

<a id="canonical-0001202120301200-1032111230300210-1033222203030230-2331313312032003-3310100323232320-2031331131022132-2002030330302302-1301112233003230"></a>

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

<a id="canonical-2221311312002300-1300300121022222-0232321211132023-0311332020111332-0022132113301000-3032221303002122-1123313231221231-1200113120011331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.port.http_loadbalancer.https.disable_path_normalize` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-006.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-006.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-009.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.port](resources--workload--reference--group-012.md#canonical-3013031222022030-2021222233122213-2220022232301003-3102220000332302-0311213311310132-3133321213021123-0133300222212330-1122131003023301)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-012.md#canonical-0101000213001300-0222020002101313-1020132032013321-2231011111202333-2122231103331210-3202123132002020-0313332003123002-3101311011132330)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https](resources--workload--reference--group-013.md#canonical-2221000320213131-2001211030022032-2023131131113110-1000133011113313-1003112332120000-2102122231000321-0211212202023222-1022013120231111)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https.disable_path_normalize

<a id="canonical-0301122012020302-0332113222320110-3301220322221002-3130120310010222-1020321103320202-3222120112233202-2011222123230111-1001200311121121"></a>

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

<a id="canonical-3233203303111321-3231322112033101-2221200310110223-1202313122212102-1223331231202221-2121122131101331-1010313202012022-2203102032221023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.port.http_loadbalancer.https.enable_path_normalize` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-006.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-006.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-009.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.port](resources--workload--reference--group-012.md#canonical-3013031222022030-2021222233122213-2220022232301003-3102220000332302-0311213311310132-3133321213021123-0133300222212330-1122131003023301)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-012.md#canonical-0101000213001300-0222020002101313-1020132032013321-2231011111202333-2122231103331210-3202123132002020-0313332003123002-3101311011132330)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https](resources--workload--reference--group-013.md#canonical-2221000320213131-2001211030022032-2023131131113110-1000133011113313-1003112332120000-2102122231000321-0211212202023222-1022013120231111)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https.enable_path_normalize

<a id="canonical-3313330222303111-0130121131122000-3133312212232331-3320033001102002-3103320033332233-3012003133000121-0101022022100302-3100303002002233"></a>

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

<a id="canonical-0113021231303233-2111320110111221-1302231000110011-2320311231321233-2223201132301001-2100032112123121-3000013033320312-1132003001210123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-006.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-006.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-009.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.port](resources--workload--reference--group-012.md#canonical-3013031222022030-2021222233122213-2220022232301003-3102220000332302-0311213311310132-3133321213021123-0133300222212330-1122131003023301)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-012.md#canonical-0101000213001300-0222020002101313-1020132032013321-2231011111202333-2122231103331210-3202123132002020-0313332003123002-3101311011132330)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https](resources--workload--reference--group-013.md#canonical-2221000320213131-2001211030022032-2023131131113110-1000133011113313-1003112332120000-2102122231000321-0211212202023222-1022013120231111)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options

<a id="canonical-1100331302233013-1321201200300220-3112100011031002-2121120332110332-1122310023001102-1100330103331332-2220030113231332-1323021230122122"></a>

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

<a id="canonical-3110100312010130-2123212030300232-1023300320003213-2011330212320233-2300022031113202-1230323022223021-1333231230033322-0301101020313330"></a>

### Direct properties for `service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options`

- [http_protocol_enable_v1_only](resources--workload--reference--group-013.md#canonical-2303202330210310-1122330231233312-2232130211120231-2331021321111121-0322101110222330-3032111122112023-1120023301030130-0310222131011321): complete subsection reference.

- [http_protocol_enable_v1_v2](resources--workload--reference--group-013.md#canonical-3001301323103222-2121231111201202-3033021201331202-2013203013231122-1022130313032203-3102302300331331-2322130230020110-1211323202320233): complete subsection reference.

- [http_protocol_enable_v2_only](resources--workload--reference--group-013.md#canonical-2233033312213312-3211033203222131-0330233230121212-1311330333020333-0230212230320330-1231010012113223-2230220213021211-0121110031221132): complete subsection reference.

<a id="canonical-2303202330210310-1122330231233312-2232130211120231-2331021321111121-0322101110222330-3032111122112023-1120023301030130-0310222131011321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-006.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-006.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-009.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.port](resources--workload--reference--group-012.md#canonical-3013031222022030-2021222233122213-2220022232301003-3102220000332302-0311213311310132-3133321213021123-0133300222212330-1122131003023301)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-012.md#canonical-0101000213001300-0222020002101313-1020132032013321-2231011111202333-2122231103331210-3202123132002020-0313332003123002-3101311011132330)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https](resources--workload--reference--group-013.md#canonical-2221000320213131-2001211030022032-2023131131113110-1000133011113313-1003112332120000-2102122231000321-0211212202023222-1022013120231111)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options](resources--workload--reference--group-013.md#canonical-0113021231303233-2111320110111221-1302231000110011-2320311231321233-2223201132301001-2100032112123121-3000013033320312-1132003001210123)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only

<a id="canonical-0112211232330301-0132003111330112-0121020101031312-2301131220122310-1203223132033013-3301122310011133-2030123113103303-2030211200033030"></a>

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

<a id="canonical-2031221311123130-3201011022003113-1003312123322100-0310132221112232-1122213132122110-1000220131223212-3101332222233301-0031010210013313"></a>

### Direct properties for `service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only`

- [header_transformation](resources--workload--reference--group-013.md#canonical-3131121300332222-1031021303200010-1332012300102232-0310333321212300-2003222110311102-0132132012320020-2312120300330300-1332031233133203): complete subsection reference.

<a id="canonical-3131121300332222-1031021303200010-1332012300102232-0310333321212300-2003222110311102-0132132012320020-2312120300330300-1332031233133203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-006.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-006.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-009.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.port](resources--workload--reference--group-012.md#canonical-3013031222022030-2021222233122213-2220022232301003-3102220000332302-0311213311310132-3133321213021123-0133300222212330-1122131003023301)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-012.md#canonical-0101000213001300-0222020002101313-1020132032013321-2231011111202333-2122231103331210-3202123132002020-0313332003123002-3101311011132330)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https](resources--workload--reference--group-013.md#canonical-2221000320213131-2001211030022032-2023131131113110-1000133011113313-1003112332120000-2102122231000321-0211212202023222-1022013120231111)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options](resources--workload--reference--group-013.md#canonical-0113021231303233-2111320110111221-1302231000110011-2320311231321233-2223201132301001-2100032112123121-3000013033320312-1132003001210123)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only](resources--workload--reference--group-013.md#canonical-2303202330210310-1122330231233312-2232130211120231-2331021321111121-0322101110222330-3032111122112023-1120023301030130-0310222131011321)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation

<a id="canonical-1213001220033223-1222233021101330-2010103112302322-3233302001000300-3212100230133122-0102023213113033-1103332311320333-1033102002230310"></a>

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

<a id="canonical-1223032131002223-3233301103201301-1231313122121122-2200123232323011-3202012232231111-1303232232320302-2010111032133121-0201320321212333"></a>

### Direct properties for `service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation`

- [default_header_transformation](resources--workload--reference--group-013.md#canonical-1231232201100230-3232031132200011-2030110220331233-0022223202311032-2332231110211321-0022003220123032-2012102022313233-1033100030111223): complete subsection reference.

- [preserve_case_header_transformation](resources--workload--reference--group-013.md#canonical-2013011031103030-1210030331231320-3122133311033210-2310331322212101-1000121001321020-2001332010221033-2232310222330223-1320311010132030): complete subsection reference.

- [proper_case_header_transformation](resources--workload--reference--group-013.md#canonical-3132322022001003-1100210323013220-2022130011201130-3330331210103031-2202332123113302-3030130333133213-0131322320120121-2302011103131201): complete subsection reference.

<a id="canonical-1231232201100230-3232031132200011-2030110220331233-0022223202311032-2332231110211321-0022003220123032-2012102022313233-1033100030111223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.default_header_transformation` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-006.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-006.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-009.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.port](resources--workload--reference--group-012.md#canonical-3013031222022030-2021222233122213-2220022232301003-3102220000332302-0311213311310132-3133321213021123-0133300222212330-1122131003023301)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-012.md#canonical-0101000213001300-0222020002101313-1020132032013321-2231011111202333-2122231103331210-3202123132002020-0313332003123002-3101311011132330)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https](resources--workload--reference--group-013.md#canonical-2221000320213131-2001211030022032-2023131131113110-1000133011113313-1003112332120000-2102122231000321-0211212202023222-1022013120231111)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options](resources--workload--reference--group-013.md#canonical-0113021231303233-2111320110111221-1302231000110011-2320311231321233-2223201132301001-2100032112123121-3000013033320312-1132003001210123)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only](resources--workload--reference--group-013.md#canonical-2303202330210310-1122330231233312-2232130211120231-2331021321111121-0322101110222330-3032111122112023-1120023301030130-0310222131011321)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation](resources--workload--reference--group-013.md#canonical-3131121300332222-1031021303200010-1332012300102232-0310333321212300-2003222110311102-0132132012320020-2312120300330300-1332031233133203)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.default_header_transformation

<a id="canonical-3100313100103200-3020232211023322-3130220033333001-3112122301011222-0123130010011103-2032212303233323-2202201303323300-0103223133101330"></a>

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

<a id="canonical-2013011031103030-1210030331231320-3122133311033210-2310331322212101-1000121001321020-2001332010221033-2232310222330223-1320311010132030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.preserve_case_header_transformation` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-006.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-006.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-009.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.port](resources--workload--reference--group-012.md#canonical-3013031222022030-2021222233122213-2220022232301003-3102220000332302-0311213311310132-3133321213021123-0133300222212330-1122131003023301)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-012.md#canonical-0101000213001300-0222020002101313-1020132032013321-2231011111202333-2122231103331210-3202123132002020-0313332003123002-3101311011132330)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https](resources--workload--reference--group-013.md#canonical-2221000320213131-2001211030022032-2023131131113110-1000133011113313-1003112332120000-2102122231000321-0211212202023222-1022013120231111)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options](resources--workload--reference--group-013.md#canonical-0113021231303233-2111320110111221-1302231000110011-2320311231321233-2223201132301001-2100032112123121-3000013033320312-1132003001210123)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only](resources--workload--reference--group-013.md#canonical-2303202330210310-1122330231233312-2232130211120231-2331021321111121-0322101110222330-3032111122112023-1120023301030130-0310222131011321)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation](resources--workload--reference--group-013.md#canonical-3131121300332222-1031021303200010-1332012300102232-0310333321212300-2003222110311102-0132132012320020-2312120300330300-1332031233133203)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.preserve_case_header_transformation

<a id="canonical-3123232131213013-1000300210303221-2231200322312231-3301203101210200-1110312230320001-2203121111223030-2121023132221002-2202200322200020"></a>

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

<a id="canonical-3132322022001003-1100210323013220-2022130011201130-3330331210103031-2202332123113302-3030130333133213-0131322320120121-2302011103131201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.proper_case_header_transformation` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-006.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-006.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-009.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.port](resources--workload--reference--group-012.md#canonical-3013031222022030-2021222233122213-2220022232301003-3102220000332302-0311213311310132-3133321213021123-0133300222212330-1122131003023301)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-012.md#canonical-0101000213001300-0222020002101313-1020132032013321-2231011111202333-2122231103331210-3202123132002020-0313332003123002-3101311011132330)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https](resources--workload--reference--group-013.md#canonical-2221000320213131-2001211030022032-2023131131113110-1000133011113313-1003112332120000-2102122231000321-0211212202023222-1022013120231111)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options](resources--workload--reference--group-013.md#canonical-0113021231303233-2111320110111221-1302231000110011-2320311231321233-2223201132301001-2100032112123121-3000013033320312-1132003001210123)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only](resources--workload--reference--group-013.md#canonical-2303202330210310-1122330231233312-2232130211120231-2331021321111121-0322101110222330-3032111122112023-1120023301030130-0310222131011321)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation](resources--workload--reference--group-013.md#canonical-3131121300332222-1031021303200010-1332012300102232-0310333321212300-2003222110311102-0132132012320020-2312120300330300-1332031233133203)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.proper_case_header_transformation

<a id="canonical-3220002302232011-2121200101312332-2222110322002030-2023220223233303-2222111301120111-3032303021101211-0223322330010221-3223321220232112"></a>

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

<a id="canonical-3001301323103222-2121231111201202-3033021201331202-2013203013231122-1022130313032203-3102302300331331-2322130230020110-1211323202320233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_v2` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-006.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-006.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-009.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.port](resources--workload--reference--group-012.md#canonical-3013031222022030-2021222233122213-2220022232301003-3102220000332302-0311213311310132-3133321213021123-0133300222212330-1122131003023301)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-012.md#canonical-0101000213001300-0222020002101313-1020132032013321-2231011111202333-2122231103331210-3202123132002020-0313332003123002-3101311011132330)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https](resources--workload--reference--group-013.md#canonical-2221000320213131-2001211030022032-2023131131113110-1000133011113313-1003112332120000-2102122231000321-0211212202023222-1022013120231111)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options](resources--workload--reference--group-013.md#canonical-0113021231303233-2111320110111221-1302231000110011-2320311231321233-2223201132301001-2100032112123121-3000013033320312-1132003001210123)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_v2

<a id="canonical-3201013323102020-0331223002033113-1100302003233221-2323222203220112-3023112202011121-1021312121132231-3012110112003010-2010133023233132"></a>

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

<a id="canonical-2233033312213312-3211033203222131-0330233230121212-1311330333020333-0230212230320330-1231010012113223-2230220213021211-0121110031221132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v2_only` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-006.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-006.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-009.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.port](resources--workload--reference--group-012.md#canonical-3013031222022030-2021222233122213-2220022232301003-3102220000332302-0311213311310132-3133321213021123-0133300222212330-1122131003023301)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-012.md#canonical-0101000213001300-0222020002101313-1020132032013321-2231011111202333-2122231103331210-3202123132002020-0313332003123002-3101311011132330)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https](resources--workload--reference--group-013.md#canonical-2221000320213131-2001211030022032-2023131131113110-1000133011113313-1003112332120000-2102122231000321-0211212202023222-1022013120231111)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options](resources--workload--reference--group-013.md#canonical-0113021231303233-2111320110111221-1302231000110011-2320311231321233-2223201132301001-2100032112123121-3000013033320312-1132003001210123)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v2_only

<a id="canonical-2023310001333321-0031321331010120-2233032001100031-0202222123231032-3010133133010012-2203121013203313-3310331032211203-3230031210112000"></a>

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

<a id="canonical-0201012102332122-3021012133330000-1212133112031131-3000111023333013-1113210313110021-3303133302312331-3303212013301222-0032301130003232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.port.http_loadbalancer.https.non_default_loadbalancer` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-006.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-006.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-009.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.port](resources--workload--reference--group-012.md#canonical-3013031222022030-2021222233122213-2220022232301003-3102220000332302-0311213311310132-3133321213021123-0133300222212330-1122131003023301)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-012.md#canonical-0101000213001300-0222020002101313-1020132032013321-2231011111202333-2122231103331210-3202123132002020-0313332003123002-3101311011132330)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https](resources--workload--reference--group-013.md#canonical-2221000320213131-2001211030022032-2023131131113110-1000133011113313-1003112332120000-2102122231000321-0211212202023222-1022013120231111)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https.non_default_loadbalancer

<a id="canonical-1130301322212311-0300222332010222-0333131130110031-0123132020222223-3321023303033030-1211130303102002-1211103300320001-2202002232130122"></a>

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

<a id="canonical-1331231121223122-3020333331231320-2222002002323013-0001220313222302-3223121012120332-1013121312110111-0010212110113013-0200301020223132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.port.http_loadbalancer.https.pass_through` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-006.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-006.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-009.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.port](resources--workload--reference--group-012.md#canonical-3013031222022030-2021222233122213-2220022232301003-3102220000332302-0311213311310132-3133321213021123-0133300222212330-1122131003023301)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-012.md#canonical-0101000213001300-0222020002101313-1020132032013321-2231011111202333-2122231103331210-3202123132002020-0313332003123002-3101311011132330)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https](resources--workload--reference--group-013.md#canonical-2221000320213131-2001211030022032-2023131131113110-1000133011113313-1003112332120000-2102122231000321-0211212202023222-1022013120231111)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https.pass_through

<a id="canonical-2223121310220231-3320123201330011-1301123000000302-3321103021232332-0322320221122123-0213301112300013-3030010213011223-1013221023312121"></a>

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

<a id="canonical-2021233301230331-3111000132213023-0320231030332320-2213232031303212-1323212201220200-0313001122300320-3212003130200223-2033202033223233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-006.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-006.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-009.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.port](resources--workload--reference--group-012.md#canonical-3013031222022030-2021222233122213-2220022232301003-3102220000332302-0311213311310132-3133321213021123-0133300222212330-1122131003023301)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-012.md#canonical-0101000213001300-0222020002101313-1020132032013321-2231011111202333-2122231103331210-3202123132002020-0313332003123002-3101311011132330)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https](resources--workload--reference--group-013.md#canonical-2221000320213131-2001211030022032-2023131131113110-1000133011113313-1003112332120000-2102122231000321-0211212202023222-1022013120231111)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params

<a id="canonical-2232302231231020-1012020022003231-2301012121212323-0102310100012023-3230003020220022-3001332033133231-0021202313311010-0222221310010230"></a>

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

<a id="canonical-2130303001020132-2013312322023032-0121100101222322-0013221312303110-3011030202220231-1002002121323231-1200221332121330-3233232231023322"></a>

### Direct properties for `service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params`

- [certificates](resources--workload--reference--group-013.md#canonical-1332231233330232-2212330030000031-1031123010312231-3301132130000032-0112000002233330-2113113200001013-3133013131213231-2221113021013102): complete subsection reference.

- [no_mtls](resources--workload--reference--group-013.md#canonical-0011012023232223-0000230302312132-2000032202200320-3213023310012111-3201230233220100-3322022110223101-0231301112102321-1211133300320132): complete subsection reference.

- [tls_config](resources--workload--reference--group-013.md#canonical-3332312101110311-0230311121302103-3013222010012313-2221303120311121-1220101010231203-2012201122333230-0101002231102123-0102233012133010): complete subsection reference.

- [use_mtls](resources--workload--reference--group-013.md#canonical-1133110233020021-3213122111302023-0212200012203202-0032103103301103-2223223111202132-3010231300103022-0312330001310030-0131021213101310): complete subsection reference.

<a id="canonical-1332231233330232-2212330030000031-1031123010312231-3301132130000032-0112000002233330-2113113200001013-3133013131213231-2221113021013102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.certificates` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-006.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-006.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-009.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.port](resources--workload--reference--group-012.md#canonical-3013031222022030-2021222233122213-2220022232301003-3102220000332302-0311213311310132-3133321213021123-0133300222212330-1122131003023301)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-012.md#canonical-0101000213001300-0222020002101313-1020132032013321-2231011111202333-2122231103331210-3202123132002020-0313332003123002-3101311011132330)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https](resources--workload--reference--group-013.md#canonical-2221000320213131-2001211030022032-2023131131113110-1000133011113313-1003112332120000-2102122231000321-0211212202023222-1022013120231111)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params](resources--workload--reference--group-013.md#canonical-2021233301230331-3111000132213023-0320231030332320-2213232031303212-1323212201220200-0313001122300320-3212003130200223-2033202033223233)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.certificates

<a id="canonical-3130220323130012-1213101133320012-0031323302123010-2320212133032230-1010332313200230-0220132011221222-3010121110330122-1102331322223302"></a>

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

Terraform syntax:

```terraform
certificates {
  # Configure direct properties listed below.
}
```

<a id="canonical-0100303022011130-2212200320031120-0133010202012330-0010312311122223-3331030022332232-3220311102330130-2101031113130212-0023023312012003"></a>

### Direct properties for `service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.certificates`

<a id="canonical-2310323232130203-2200020301202001-3222310211013331-1211131012220330-0112212122222322-0122203101333113-1031120110033320-3032301112310230"></a>

#### `service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.certificates.name` property

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

<a id="canonical-3121302333331222-1201133110132323-2122321232131303-3301203220303131-0002300001023200-1320100333121213-0311202322131213-0111000122020333"></a>

<a id="canonical-2230032303120231-2122220101033103-2333230020232021-3131130222210101-0323130332121230-1333220121310223-1001012202002002-1003223313012323"></a>

#### `service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.certificates.namespace` property

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

<a id="canonical-2132130223133022-1101130222111233-2232223001012301-0213313002033333-2300110111233213-1220312131232132-1121301012212322-3202301300023211"></a>

<a id="canonical-1033223223300203-3131232221233323-3221232103321213-2002010312303032-3202111311231101-0002000122322203-1023313122000001-0021020002303030"></a>

#### `service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.certificates.tenant` property

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

<a id="canonical-0011012023232223-0000230302312132-2000032202200320-3213023310012111-3201230233220100-3322022110223101-0231301112102321-1211133300320132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.no_mtls` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-006.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-006.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-009.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.port](resources--workload--reference--group-012.md#canonical-3013031222022030-2021222233122213-2220022232301003-3102220000332302-0311213311310132-3133321213021123-0133300222212330-1122131003023301)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-012.md#canonical-0101000213001300-0222020002101313-1020132032013321-2231011111202333-2122231103331210-3202123132002020-0313332003123002-3101311011132330)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https](resources--workload--reference--group-013.md#canonical-2221000320213131-2001211030022032-2023131131113110-1000133011113313-1003112332120000-2102122231000321-0211212202023222-1022013120231111)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params](resources--workload--reference--group-013.md#canonical-2021233301230331-3111000132213023-0320231030332320-2213232031303212-1323212201220200-0313001122300320-3212003130200223-2033202033223233)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.no_mtls

<a id="canonical-3212012001132331-2122032222021103-1002103202320000-0331323330311323-3222213032001310-3130223313321312-2320201132310022-2120313012303032"></a>

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

<a id="canonical-3332312101110311-0230311121302103-3013222010012313-2221303120311121-1220101010231203-2012201122333230-0101002231102123-0102233012133010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.tls_config` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-006.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-006.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-009.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.port](resources--workload--reference--group-012.md#canonical-3013031222022030-2021222233122213-2220022232301003-3102220000332302-0311213311310132-3133321213021123-0133300222212330-1122131003023301)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-012.md#canonical-0101000213001300-0222020002101313-1020132032013321-2231011111202333-2122231103331210-3202123132002020-0313332003123002-3101311011132330)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https](resources--workload--reference--group-013.md#canonical-2221000320213131-2001211030022032-2023131131113110-1000133011113313-1003112332120000-2102122231000321-0211212202023222-1022013120231111)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params](resources--workload--reference--group-013.md#canonical-2021233301230331-3111000132213023-0320231030332320-2213232031303212-1323212201220200-0313001122300320-3212003130200223-2033202033223233)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.tls_config

<a id="canonical-1332121111112033-0001032322323010-2232211000031330-0013013230133300-3310001321103311-0122031002132230-2211332331312221-1011030020222101"></a>

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

<a id="canonical-2012102303301013-1113112231032110-2313011212013002-2313211221102302-3302220302203331-0011031031000131-0111100120221230-0033122113333302"></a>

### Direct properties for `service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.tls_config`

- [custom_security](resources--workload--reference--group-013.md#canonical-0010323001301013-3113303210311201-2210331210100200-3031020330333100-3011102312210201-1202000121203303-1212330102321232-2320002332131212): complete subsection reference.

- [default_security](resources--workload--reference--group-013.md#canonical-2223331133123222-1113301202210223-3321233300000132-2332203313133023-0300110100001013-0030232000320233-3212301101111330-2000131332303223): complete subsection reference.

- [low_security](resources--workload--reference--group-013.md#canonical-0001323101131312-1121011002323102-3121112301201230-2202312031210220-1013331211121311-1320211123310020-0333311322201322-3102111223301103): complete subsection reference.

- [medium_security](resources--workload--reference--group-013.md#canonical-2121203033100222-2110122010211331-2322110121120303-1312221312130012-1311210200020333-3201330203100312-1220323322221022-2313030000323310): complete subsection reference.

<a id="canonical-0010323001301013-3113303210311201-2210331210100200-3031020330333100-3011102312210201-1202000121203303-1212330102321232-2320002332131212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.tls_config.custom_security` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-006.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-006.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-009.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.port](resources--workload--reference--group-012.md#canonical-3013031222022030-2021222233122213-2220022232301003-3102220000332302-0311213311310132-3133321213021123-0133300222212330-1122131003023301)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-012.md#canonical-0101000213001300-0222020002101313-1020132032013321-2231011111202333-2122231103331210-3202123132002020-0313332003123002-3101311011132330)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https](resources--workload--reference--group-013.md#canonical-2221000320213131-2001211030022032-2023131131113110-1000133011113313-1003112332120000-2102122231000321-0211212202023222-1022013120231111)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params](resources--workload--reference--group-013.md#canonical-2021233301230331-3111000132213023-0320231030332320-2213232031303212-1323212201220200-0313001122300320-3212003130200223-2033202033223233)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.tls_config](resources--workload--reference--group-013.md#canonical-3332312101110311-0230311121302103-3013222010012313-2221303120311121-1220101010231203-2012201122333230-0101002231102123-0102233012133010)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.tls_config.custom_security

<a id="canonical-3132301202123223-3202311303032010-3112131010000221-3300120123132020-2302230200231300-1121210233301121-0210212213000310-0122032332102321"></a>

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

<a id="canonical-0213322202331020-0102111332212020-3133103020221011-1113123100231301-2110233201022130-3101201301220112-1303033222000202-0223002333211331"></a>

### Direct properties for `service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.tls_config.custom_security`

<a id="canonical-1012000121122132-1203033301133003-1233231233332313-1101121321231031-2012010302100223-2011313333301020-1200201310003033-1331310011200303"></a>

#### `service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.tls_config.custom_security.cipher_suites` property

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

<a id="canonical-0111031301122010-3303302323032101-3122022033322232-1212120110111010-2310233121032222-3031222103220223-1033300233232111-1020003322013210"></a>

<a id="canonical-2233010310312312-3221233210301013-2212210203201110-1111212320322203-3312221210012010-1001131100322233-0101212322311212-0303122312230122"></a>

#### `service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.tls_config.custom_security.max_version` property

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

<a id="canonical-2332211013110202-3333020023100331-2332211232133122-2233303112032310-3110031201320333-1301011121003232-2223313320113312-3222122120332222"></a>

<a id="canonical-0130012100220032-0202300133033101-0013111202000303-1132313022201230-1133201310220223-1020130120020130-1131023232223132-3332332023222011"></a>

#### `service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.tls_config.custom_security.min_version` property

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

<a id="canonical-2223331133123222-1113301202210223-3321233300000132-2332203313133023-0300110100001013-0030232000320233-3212301101111330-2000131332303223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.tls_config.default_security` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-006.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-006.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-009.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.port](resources--workload--reference--group-012.md#canonical-3013031222022030-2021222233122213-2220022232301003-3102220000332302-0311213311310132-3133321213021123-0133300222212330-1122131003023301)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-012.md#canonical-0101000213001300-0222020002101313-1020132032013321-2231011111202333-2122231103331210-3202123132002020-0313332003123002-3101311011132330)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https](resources--workload--reference--group-013.md#canonical-2221000320213131-2001211030022032-2023131131113110-1000133011113313-1003112332120000-2102122231000321-0211212202023222-1022013120231111)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params](resources--workload--reference--group-013.md#canonical-2021233301230331-3111000132213023-0320231030332320-2213232031303212-1323212201220200-0313001122300320-3212003130200223-2033202033223233)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.tls_config](resources--workload--reference--group-013.md#canonical-3332312101110311-0230311121302103-3013222010012313-2221303120311121-1220101010231203-2012201122333230-0101002231102123-0102233012133010)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.tls_config.default_security

<a id="canonical-0322130130323020-1310213213200220-2322323212000320-3132130220001132-3300110033133023-1001200012221022-2210230130013133-2301221223233323"></a>

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

<a id="canonical-0001323101131312-1121011002323102-3121112301201230-2202312031210220-1013331211121311-1320211123310020-0333311322201322-3102111223301103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.tls_config.low_security` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-006.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-006.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-009.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.port](resources--workload--reference--group-012.md#canonical-3013031222022030-2021222233122213-2220022232301003-3102220000332302-0311213311310132-3133321213021123-0133300222212330-1122131003023301)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-012.md#canonical-0101000213001300-0222020002101313-1020132032013321-2231011111202333-2122231103331210-3202123132002020-0313332003123002-3101311011132330)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https](resources--workload--reference--group-013.md#canonical-2221000320213131-2001211030022032-2023131131113110-1000133011113313-1003112332120000-2102122231000321-0211212202023222-1022013120231111)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params](resources--workload--reference--group-013.md#canonical-2021233301230331-3111000132213023-0320231030332320-2213232031303212-1323212201220200-0313001122300320-3212003130200223-2033202033223233)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.tls_config](resources--workload--reference--group-013.md#canonical-3332312101110311-0230311121302103-3013222010012313-2221303120311121-1220101010231203-2012201122333230-0101002231102123-0102233012133010)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.tls_config.low_security

<a id="canonical-2130120332032301-3222101331013232-0220132220331120-2201323221133312-3300330123301232-2211302030133313-1331003122111212-3011332012231311"></a>

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

<a id="canonical-2121203033100222-2110122010211331-2322110121120303-1312221312130012-1311210200020333-3201330203100312-1220323322221022-2313030000323310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.tls_config.medium_security` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-006.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-006.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-009.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.port](resources--workload--reference--group-012.md#canonical-3013031222022030-2021222233122213-2220022232301003-3102220000332302-0311213311310132-3133321213021123-0133300222212330-1122131003023301)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-012.md#canonical-0101000213001300-0222020002101313-1020132032013321-2231011111202333-2122231103331210-3202123132002020-0313332003123002-3101311011132330)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https](resources--workload--reference--group-013.md#canonical-2221000320213131-2001211030022032-2023131131113110-1000133011113313-1003112332120000-2102122231000321-0211212202023222-1022013120231111)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params](resources--workload--reference--group-013.md#canonical-2021233301230331-3111000132213023-0320231030332320-2213232031303212-1323212201220200-0313001122300320-3212003130200223-2033202033223233)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.tls_config](resources--workload--reference--group-013.md#canonical-3332312101110311-0230311121302103-3013222010012313-2221303120311121-1220101010231203-2012201122333230-0101002231102123-0102233012133010)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.tls_config.medium_security

<a id="canonical-1132333331231331-0023303232311331-0131323203023103-1233030333331033-1003001300100132-2033302020103023-2300021312222323-2231311303221202"></a>

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

<a id="canonical-1133110233020021-3213122111302023-0212200012203202-0032103103301103-2223223111202132-3010231300103022-0312330001310030-0131021213101310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.use_mtls` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-006.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-006.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-009.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.port](resources--workload--reference--group-012.md#canonical-3013031222022030-2021222233122213-2220022232301003-3102220000332302-0311213311310132-3133321213021123-0133300222212330-1122131003023301)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-012.md#canonical-0101000213001300-0222020002101313-1020132032013321-2231011111202333-2122231103331210-3202123132002020-0313332003123002-3101311011132330)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https](resources--workload--reference--group-013.md#canonical-2221000320213131-2001211030022032-2023131131113110-1000133011113313-1003112332120000-2102122231000321-0211212202023222-1022013120231111)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params](resources--workload--reference--group-013.md#canonical-2021233301230331-3111000132213023-0320231030332320-2213232031303212-1323212201220200-0313001122300320-3212003130200223-2033202033223233)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.use_mtls

<a id="canonical-0323113333212132-2301033110122033-1313323231320101-3303232333013202-0302311312011203-2130211331033030-1013322013102201-0013012103332210"></a>

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

<a id="canonical-0112232221323200-2010123303003223-0123110013020221-3222112112222320-1330033210131211-2233113210330110-3202011112321202-1233333100003012"></a>

### Direct properties for `service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.use_mtls`

<a id="canonical-3231121221211321-1132001220023120-0300222123113232-3210020000131210-3000021122103033-3233233221312120-3232100012110320-3011101120302011"></a>

#### `service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.use_mtls.client_certificate_optional` property

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

- [crl](resources--workload--reference--group-013.md#canonical-2333013222311013-3013121331033312-0202001102221100-0330100003300031-1003032031202112-2211332203032212-2113211110200013-1101222300232333): complete subsection reference.

- [no_crl](resources--workload--reference--group-013.md#canonical-0012123301120011-3122210100113032-0303230320321120-1120322332211203-2000113033001313-1313320031110202-2030323320212332-1000302013110313): complete subsection reference.

- [trusted_ca](resources--workload--reference--group-013.md#canonical-2202303230230010-3022110330201013-3122131132331110-2011001001002012-0310011102200011-2330222123002012-0332132012022001-1231231222202021): complete subsection reference.

<a id="canonical-0022013230220221-0110131021033103-0221232322112130-1101112103331230-3020312312330031-1220221212311013-3330131312011303-3223310111213110"></a>

<a id="canonical-1001021111222132-1310322020213120-1232013013013030-0021133330322121-1300311211100202-2322330311013101-1031310330223022-1030312130322133"></a>

#### `service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.use_mtls.trusted_ca_url` property

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

- [xfcc_disabled](resources--workload--reference--group-013.md#canonical-0321203011031320-1300132222000301-1031301210023313-3103100202333322-2202312022030332-3131233213311012-0010231212231032-1202221221023333): complete subsection reference.

- [xfcc_options](resources--workload--reference--group-013.md#canonical-1001300333302321-2023302021011032-2333303131002123-0331011331023021-1132311021112011-0230213112023320-2231333100320120-3111331001001210): complete subsection reference.

<a id="canonical-2333013222311013-3013121331033312-0202001102221100-0330100003300031-1003032031202112-2211332203032212-2113211110200013-1101222300232333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.use_mtls.crl` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-006.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-006.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-009.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.port](resources--workload--reference--group-012.md#canonical-3013031222022030-2021222233122213-2220022232301003-3102220000332302-0311213311310132-3133321213021123-0133300222212330-1122131003023301)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-012.md#canonical-0101000213001300-0222020002101313-1020132032013321-2231011111202333-2122231103331210-3202123132002020-0313332003123002-3101311011132330)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https](resources--workload--reference--group-013.md#canonical-2221000320213131-2001211030022032-2023131131113110-1000133011113313-1003112332120000-2102122231000321-0211212202023222-1022013120231111)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params](resources--workload--reference--group-013.md#canonical-2021233301230331-3111000132213023-0320231030332320-2213232031303212-1323212201220200-0313001122300320-3212003130200223-2033202033223233)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.use_mtls](resources--workload--reference--group-013.md#canonical-1133110233020021-3213122111302023-0212200012203202-0032103103301103-2223223111202132-3010231300103022-0312330001310030-0131021213101310)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.use_mtls.crl

<a id="canonical-2200012113322032-1032320333303021-1213023021033302-0100313332333023-0300333032202131-3021201211300130-1312321211221230-2203202220210120"></a>

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

<a id="canonical-3001220231331123-3100300211213033-3033120000010100-1013033102201203-3111122011331321-1201203233331112-0322131220100202-3111223330002221"></a>

### Direct properties for `service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.use_mtls.crl`

<a id="canonical-2130122312130000-2020303313101003-3202000132132100-3003322320103332-1113303311021201-1232112210300203-0330010303222232-3310302330223331"></a>

#### `service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.use_mtls.crl.name` property

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

<a id="canonical-2330102002322121-0002222002021311-1001302220133000-0321000303103331-3112030300223222-2101310300301002-1300313123032213-2110231230231011"></a>

<a id="canonical-1331212321212233-2102301001321022-2223301333312323-2111030301002313-1010022211130033-1313123311331110-1313023312213103-3012032132113132"></a>

#### `service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.use_mtls.crl.namespace` property

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

<a id="canonical-0122121120323132-1033012021230311-1103000303100313-2311320100111031-2120001213232123-0312120031111212-2111031300022231-0310210223021102"></a>

<a id="canonical-1101231003333123-3230320221020232-0330130310301033-2202311303132003-0312001131120013-3200130032010211-2203113021331322-0322203103220203"></a>

#### `service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.use_mtls.crl.tenant` property

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

<a id="canonical-0012123301120011-3122210100113032-0303230320321120-1120322332211203-2000113033001313-1313320031110202-2030323320212332-1000302013110313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.use_mtls.no_crl` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-006.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-006.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-009.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.port](resources--workload--reference--group-012.md#canonical-3013031222022030-2021222233122213-2220022232301003-3102220000332302-0311213311310132-3133321213021123-0133300222212330-1122131003023301)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-012.md#canonical-0101000213001300-0222020002101313-1020132032013321-2231011111202333-2122231103331210-3202123132002020-0313332003123002-3101311011132330)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https](resources--workload--reference--group-013.md#canonical-2221000320213131-2001211030022032-2023131131113110-1000133011113313-1003112332120000-2102122231000321-0211212202023222-1022013120231111)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params](resources--workload--reference--group-013.md#canonical-2021233301230331-3111000132213023-0320231030332320-2213232031303212-1323212201220200-0313001122300320-3212003130200223-2033202033223233)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.use_mtls](resources--workload--reference--group-013.md#canonical-1133110233020021-3213122111302023-0212200012203202-0032103103301103-2223223111202132-3010231300103022-0312330001310030-0131021213101310)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.use_mtls.no_crl

<a id="canonical-2232020112133312-1231021002212112-0033113321000213-1330031231021310-2322033310202231-1222212030333003-0220203022001303-2203221022000031"></a>

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

<a id="canonical-2202303230230010-3022110330201013-3122131132331110-2011001001002012-0310011102200011-2330222123002012-0332132012022001-1231231222202021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.use_mtls.trusted_ca` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-006.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-006.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-009.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.port](resources--workload--reference--group-012.md#canonical-3013031222022030-2021222233122213-2220022232301003-3102220000332302-0311213311310132-3133321213021123-0133300222212330-1122131003023301)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-012.md#canonical-0101000213001300-0222020002101313-1020132032013321-2231011111202333-2122231103331210-3202123132002020-0313332003123002-3101311011132330)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https](resources--workload--reference--group-013.md#canonical-2221000320213131-2001211030022032-2023131131113110-1000133011113313-1003112332120000-2102122231000321-0211212202023222-1022013120231111)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params](resources--workload--reference--group-013.md#canonical-2021233301230331-3111000132213023-0320231030332320-2213232031303212-1323212201220200-0313001122300320-3212003130200223-2033202033223233)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.use_mtls](resources--workload--reference--group-013.md#canonical-1133110233020021-3213122111302023-0212200012203202-0032103103301103-2223223111202132-3010231300103022-0312330001310030-0131021213101310)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.use_mtls.trusted_ca

<a id="canonical-3230310323031112-0303113313033013-3201003313222203-0320231030120312-1330020302112003-2322221320211133-0033112133200232-3330231331201313"></a>

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

<a id="canonical-3110003032001012-2022032200131321-1031010023110201-3233311223031122-2001331233202000-1103311202010233-0331033013030301-3321002211322220"></a>

### Direct properties for `service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.use_mtls.trusted_ca`

<a id="canonical-0200211103333031-2003000332013110-0202002031333133-3221302011231200-1302122002311111-2310001021331213-3123321002203020-1020300312112302"></a>

#### `service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.use_mtls.trusted_ca.name` property

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

<a id="canonical-1102023102131203-1302221033303213-2012201123120003-2102200022202023-0231102033132323-2013021200332301-3110030002013333-2302323133100120"></a>

<a id="canonical-2103002002311131-0230203020113121-1023332111312321-3212000022220232-1200011320031013-0302330033112102-0133132322030321-1132020313330133"></a>

#### `service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.use_mtls.trusted_ca.namespace` property

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

<a id="canonical-3010233113130320-0110330323100322-2100021013021032-1102002022221202-3101233320222331-2213002222011121-3032312333002110-1330320033330123"></a>

<a id="canonical-3233022313330230-0022132002230123-2101201012032031-2022333223221032-3110003332000013-3320322212203203-2120110132300133-1111010320030021"></a>

#### `service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.use_mtls.trusted_ca.tenant` property

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

<a id="canonical-0321203011031320-1300132222000301-1031301210023313-3103100202333322-2202312022030332-3131233213311012-0010231212231032-1202221221023333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.use_mtls.xfcc_disabled` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-006.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-006.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-009.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.port](resources--workload--reference--group-012.md#canonical-3013031222022030-2021222233122213-2220022232301003-3102220000332302-0311213311310132-3133321213021123-0133300222212330-1122131003023301)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-012.md#canonical-0101000213001300-0222020002101313-1020132032013321-2231011111202333-2122231103331210-3202123132002020-0313332003123002-3101311011132330)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https](resources--workload--reference--group-013.md#canonical-2221000320213131-2001211030022032-2023131131113110-1000133011113313-1003112332120000-2102122231000321-0211212202023222-1022013120231111)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params](resources--workload--reference--group-013.md#canonical-2021233301230331-3111000132213023-0320231030332320-2213232031303212-1323212201220200-0313001122300320-3212003130200223-2033202033223233)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.use_mtls](resources--workload--reference--group-013.md#canonical-1133110233020021-3213122111302023-0212200012203202-0032103103301103-2223223111202132-3010231300103022-0312330001310030-0131021213101310)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.use_mtls.xfcc_disabled

<a id="canonical-0202010130132331-1333012230200010-0210001113210131-1222331010010223-0313022301233323-3001210000022023-0300102232113300-2002000000330113"></a>

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

<a id="canonical-1001300333302321-2023302021011032-2333303131002123-0331011331023021-1132311021112011-0230213112023320-2231333100320120-3111331001001210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.use_mtls.xfcc_options` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-006.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-006.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-009.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.port](resources--workload--reference--group-012.md#canonical-3013031222022030-2021222233122213-2220022232301003-3102220000332302-0311213311310132-3133321213021123-0133300222212330-1122131003023301)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-012.md#canonical-0101000213001300-0222020002101313-1020132032013321-2231011111202333-2122231103331210-3202123132002020-0313332003123002-3101311011132330)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https](resources--workload--reference--group-013.md#canonical-2221000320213131-2001211030022032-2023131131113110-1000133011113313-1003112332120000-2102122231000321-0211212202023222-1022013120231111)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params](resources--workload--reference--group-013.md#canonical-2021233301230331-3111000132213023-0320231030332320-2213232031303212-1323212201220200-0313001122300320-3212003130200223-2033202033223233)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.use_mtls](resources--workload--reference--group-013.md#canonical-1133110233020021-3213122111302023-0212200012203202-0032103103301103-2223223111202132-3010231300103022-0312330001310030-0131021213101310)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.use_mtls.xfcc_options

<a id="canonical-2032031310113231-3003012220130103-2321030121200231-0003220112331210-2122301032331130-2013233311331222-0002122233223310-0332120121133031"></a>

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

<a id="canonical-3113121222213112-1331212102002311-1220230011301313-2203223232232213-1022112211320001-0122311031100122-0303311301301122-0301301130220122"></a>

### Direct properties for `service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.use_mtls.xfcc_options`

<a id="canonical-2122302232312113-2330112233203012-3320111131000132-2233110123201112-2200012230211132-0203222012303222-2210230100110132-1021031202021332"></a>

#### `service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.use_mtls.xfcc_options.xfcc_header_elements` property

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

<a id="canonical-1231003120133112-2013110130100211-2233121010211001-2121031322013232-2232220303133231-0300100233200310-0003012301120011-3102210111030310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-006.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-006.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-009.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.port](resources--workload--reference--group-012.md#canonical-3013031222022030-2021222233122213-2220022232301003-3102220000332302-0311213311310132-3133321213021123-0133300222212330-1122131003023301)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-012.md#canonical-0101000213001300-0222020002101313-1020132032013321-2231011111202333-2122231103331210-3202123132002020-0313332003123002-3101311011132330)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https](resources--workload--reference--group-013.md#canonical-2221000320213131-2001211030022032-2023131131113110-1000133011113313-1003112332120000-2102122231000321-0211212202023222-1022013120231111)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters

<a id="canonical-3102300333113320-2211111332222003-1122230313030302-0312203003210113-3231203320032331-1333131110100203-3313023220221332-1233120103003331"></a>

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

<a id="canonical-1112031020011221-0011001201302133-3312212330313001-0131102102113113-1102002131013202-2311102231301002-0330301033330203-1133221333302023"></a>

### Direct properties for `service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters`

- [no_mtls](resources--workload--reference--group-013.md#canonical-0212021031112230-1222201120330130-3021011223220013-1200111013202002-2020013232230112-2333312002003133-3231113111302102-2300110133232223): complete subsection reference.

- [tls_certificates](resources--workload--reference--group-013.md#canonical-1222323211330211-1131011013013011-0010113301022222-0300023112121103-0132020001012213-1113030220101112-1233001112311310-2321032212303310): complete subsection reference.

- [tls_config](resources--workload--reference--group-013.md#canonical-3232132023002231-3222012312123031-0032313002232322-3220233203030300-3132010213033111-2003310330231321-1313101010130131-3002200003132000): complete subsection reference.

- [use_mtls](resources--workload--reference--group-014.md#canonical-2212033132123331-1022023101100202-0012200211203323-3100333332232032-2000331210002203-3023131312000213-0233333201322123-2002202002102311): complete subsection reference.

<a id="canonical-0212021031112230-1222201120330130-3021011223220013-1200111013202002-2020013232230112-2333312002003133-3231113111302102-2300110133232223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.no_mtls` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-006.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-006.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-009.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.port](resources--workload--reference--group-012.md#canonical-3013031222022030-2021222233122213-2220022232301003-3102220000332302-0311213311310132-3133321213021123-0133300222212330-1122131003023301)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-012.md#canonical-0101000213001300-0222020002101313-1020132032013321-2231011111202333-2122231103331210-3202123132002020-0313332003123002-3101311011132330)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https](resources--workload--reference--group-013.md#canonical-2221000320213131-2001211030022032-2023131131113110-1000133011113313-1003112332120000-2102122231000321-0211212202023222-1022013120231111)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-013.md#canonical-1231003120133112-2013110130100211-2233121010211001-2121031322013232-2232220303133231-0300100233200310-0003012301120011-3102210111030310)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.no_mtls

<a id="canonical-2101122322023123-1102322233212111-0201012223223121-0220310130211110-1202313033030110-3101331313030012-2130303302032220-3313112331013201"></a>

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

<a id="canonical-1222323211330211-1131011013013011-0010113301022222-0300023112121103-0132020001012213-1113030220101112-1233001112311310-2321032212303310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-006.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-006.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-009.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.port](resources--workload--reference--group-012.md#canonical-3013031222022030-2021222233122213-2220022232301003-3102220000332302-0311213311310132-3133321213021123-0133300222212330-1122131003023301)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-012.md#canonical-0101000213001300-0222020002101313-1020132032013321-2231011111202333-2122231103331210-3202123132002020-0313332003123002-3101311011132330)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https](resources--workload--reference--group-013.md#canonical-2221000320213131-2001211030022032-2023131131113110-1000133011113313-1003112332120000-2102122231000321-0211212202023222-1022013120231111)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-013.md#canonical-1231003120133112-2013110130100211-2233121010211001-2121031322013232-2232220303133231-0300100233200310-0003012301120011-3102210111030310)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates

<a id="canonical-1203212033321112-1102222133031123-2232102200223223-0221102223013300-3203131013002233-2022103332300000-0202111131002110-2201002233112213"></a>

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-1321012002122000-2122021302022032-2322202000033312-3103330132022101-1220321112021032-2333010112323210-1111300112202211-1323023220021201"></a>

### Direct properties for `service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates`

- [blindfold](resources--workload--reference--group-013.md#canonical-3311333130231231-0022301303130123-0022300303332120-3121032120130122-1330303210120210-0323223123310302-0122200021200113-3210213033301310): complete subsection reference.

<a id="canonical-0312220001110121-3321000330012120-2300103313113122-1203010010202021-1023330300320220-3203112223021000-3023100320031123-0222300000031123"></a>

<a id="canonical-2001233022320032-0300022333311113-1100321023021312-1213012022102010-2330003323021200-3303220322020203-3203123012032033-3000102231233230"></a>

#### `service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates.certificate_url` property

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

- [custom_hash_algorithms](resources--workload--reference--group-013.md#canonical-0101000310022230-0310111032313202-3003331220301302-2201020020020300-1131310023233313-2302331220210031-3330130102113310-3030203030003013): complete subsection reference.

<a id="canonical-2233311000001232-0121320113210133-2022122110201113-1112131130113232-0013123133120012-3031213133202121-3103333020320121-0323332222132101"></a>

<a id="canonical-0030310012320001-1302111313121033-3320213300133222-0203101021222332-2211333212223001-2203310313101003-2122111301213230-0221033011223031"></a>

#### `service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates.description_spec` property

Type: `"string"`. Optional.

Description. Description for the certificate.

- [disable_ocsp_stapling](resources--workload--reference--group-013.md#canonical-1111013233010131-2221110113130120-2312121320312022-2303321110232001-2311233333330203-2303313001013112-3232103110122111-1013303023313130): complete subsection reference.

- [private_key](resources--workload--reference--group-013.md#canonical-2033112022013132-3022203212311013-1130313200322231-2212013001100010-0011112031131213-3002001011030201-0121133021012111-1012222332230300): complete subsection reference.

- [use_system_defaults](resources--workload--reference--group-013.md#canonical-3220122011200333-3001230313130220-3013223100030322-1021301010012321-0032223013233321-2102220020230330-1113232313100312-0123202032102002): complete subsection reference.

<a id="canonical-3311333130231231-0022301303130123-0022300303332120-3121032120130122-1330303210120210-0323223123310302-0122200021200113-3210213033301310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates.blindfold` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-006.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-006.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-009.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.port](resources--workload--reference--group-012.md#canonical-3013031222022030-2021222233122213-2220022232301003-3102220000332302-0311213311310132-3133321213021123-0133300222212330-1122131003023301)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-012.md#canonical-0101000213001300-0222020002101313-1020132032013321-2231011111202333-2122231103331210-3202123132002020-0313332003123002-3101311011132330)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https](resources--workload--reference--group-013.md#canonical-2221000320213131-2001211030022032-2023131131113110-1000133011113313-1003112332120000-2102122231000321-0211212202023222-1022013120231111)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-013.md#canonical-1231003120133112-2013110130100211-2233121010211001-2121031322013232-2232220303133231-0300100233200310-0003012301120011-3102210111030310)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates](resources--workload--reference--group-013.md#canonical-1222323211330211-1131011013013011-0010113301022222-0300023112121103-0132020001012213-1113030220101112-1233001112311310-2321032212303310)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates.blindfold

<a id="canonical-0111220320222311-1030100330113201-1221331100123333-1122100213112013-3022320223121113-0310030131130121-1123301313130203-3030323201220012"></a>

Type: `"single"`. Optional.

Native certificate preparation. Use PEM files or a P12 file, or write-only key/bundle values with
material\_version (Terraform 1.11+). Defaults to shared/ves-io-allow-volterra. Inline certificates
require unique IDs. Private inputs are never stored.

<a id="canonical-2030013110122323-2021333121103122-3222332121002332-1312232030111003-3322133331120103-1311122011032123-1103331323301101-0201123301020011"></a>

### Direct properties for `service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates.blindfold`

<a id="canonical-0123003212103003-2333011232122100-2022300001311003-0131130233011112-1321201001320013-1232130202321021-0223032310030302-3321020211201130"></a>

#### `service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates.blindfold.algorithm` property

Type: `"string"`. Computed.

<a id="canonical-3332222201003212-1232330123203230-1132310011100121-1211022323111310-1023103213033001-3211121311013132-1331122131202103-3032332211321110"></a>

<a id="canonical-2323333213010101-2311131331323012-1131131332232233-3301222032101213-0220010321021011-2103112111332013-0210013132011130-3033302020230221"></a>

#### `service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates.blindfold.certificate_file` property

Type: `"string"`. Optional.

<a id="canonical-2331331212223212-3003310003002210-3001022301220232-1120200130023212-1010330331113012-3220013300322110-1133331211210002-3233010313310001"></a>

<a id="canonical-1203100220021133-2022321111102000-1231303333311200-0213223302023103-0001223331313220-2101110101330233-3333013202020212-0232201113322322"></a>

#### `service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates.blindfold.certificate_pem` property

Type: `"string"`. Optional.

<a id="canonical-0312132202112102-0133112210012133-1120003101133212-2102230220002102-3311010010203333-0311230100100312-2001303220213310-1013110212231123"></a>

<a id="canonical-1220103320200132-2023031132112112-3323211311320000-0211223321203322-0122121330302202-2020231032200020-3002201221030020-3111212322312313"></a>

#### `service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates.blindfold.chain_identity` property

Type: `"string"`. Computed.

<a id="canonical-0100001213233223-1303023121202132-2102122313021003-0322112220223033-3332121201131011-2132113011020030-0311101012223033-0101022011203003"></a>

<a id="canonical-2303222313330232-3212113022002131-3200111100330212-1123322333333233-3212023201100233-3301131100112231-1030312123033223-0310310112121333"></a>

#### `service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates.blindfold.context_digest` property

Type: `"string"`. Computed.

<a id="canonical-2320032012222030-3033003110022013-1102230120223323-3331020331132333-0313102223113102-0233300130002210-0113002023333010-2130203223201110"></a>

<a id="canonical-0023010133301232-1200321003322333-3103222303131330-3303000113213013-1332330322100210-1301201120022032-3002212323201300-3201013330111000"></a>

#### `service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates.blindfold.encrypted_location` property

Type: `"string"`. Computed, Sensitive.

<a id="canonical-2223100002320003-2112023011023200-0313313112113123-1010312302323222-3113221113033011-3320333211302123-2132132010222110-2331320130002131"></a>

<a id="canonical-1013323133010222-0111311031232021-3023333312132232-1312131220303221-3313120000113022-3010011123320212-0210032102133330-0013230102111313"></a>

#### `service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates.blindfold.expires_at` property

Type: `"string"`. Computed.

<a id="canonical-0013130300210101-1121200333132010-2333212100032220-2223320111220133-0202030000202311-1121112100312112-1322020323213200-1233330323310110"></a>

<a id="canonical-0103230333212300-1123302130100313-3130030000201310-3010100230132220-3001213131132111-0030213333033121-1333301232320321-3020201013010233"></a>

#### `service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates.blindfold.fingerprint` property

Type: `"string"`. Computed.

<a id="canonical-1220100131233022-3201112122220033-1323013301220232-2330313120213133-1030331003321202-2312033021331002-1202100320200201-0232230033300202"></a>

<a id="canonical-1013001130020111-0120122003312232-2311300032320011-3303313020013001-2323000323221133-2130320222220303-3032113310310100-3203233230031111"></a>

#### `service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates.blindfold.id` property

Type: `"string"`. Optional.

<a id="canonical-3211103200000313-0131013001201123-3300033230122121-1010213030213111-2031313333210110-3033102121113221-3312201223233020-3010032013131012"></a>

<a id="canonical-2200313022200020-3100312122101201-2133230030002220-3332001001223013-1232021311133112-1023112212332321-0211030103333102-0021022213330311"></a>

#### `service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates.blindfold.material_version` property

Type: `"string"`. Optional.

<a id="canonical-3311303000320321-1210010111000112-2031302031330200-2003201313133113-2102211120110210-1120001012121212-1111203013213130-0010220303311233"></a>

<a id="canonical-2110230310213023-0002332022202111-0232203012211230-2213010112112031-1311113013332323-3030011032300131-3133230321230030-1331320233203010"></a>

#### `service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates.blindfold.passphrase_env` property

Type: `"string"`. Optional.

<a id="canonical-1320213200303000-0003202313222321-2103030300232023-3100311320132021-2122210201002122-3003112302102002-0230312203122033-2222202130132222"></a>

<a id="canonical-1302110111333100-1101220220032130-2223220033122013-1223012230221131-1213201132213022-1333310020013313-3122330023201012-1113320021332102"></a>

#### `service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates.blindfold.passphrase_wo` property

Type: `"string"`. Optional, Sensitive, Write_only.

<a id="canonical-1020201210002033-1300231311202320-1001300200202320-2023220230132212-2110013031000112-0321131331231121-1310333302122032-2121031000013023"></a>

<a id="canonical-1201111113201011-2210200033022313-3310021012130320-0332303033200013-0131300031033100-0201220022001202-1022222131312111-3003312113311113"></a>

#### `service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates.blindfold.pkcs12_file` property

Type: `"string"`. Optional.

<a id="canonical-3332311123323321-2022312111203123-3101033220000222-0130013032312110-3002130002002233-2123123120001103-1021300300233331-0333320110331333"></a>

<a id="canonical-0323101200122112-0103110101213021-2030201103202200-2212303201021233-3201210311321210-2211233231103100-1010020333321212-1111021231022103"></a>

#### `service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates.blindfold.pkcs12_wo` property

Type: `"string"`. Optional, Sensitive, Write_only.

<a id="canonical-1211031331033021-1100303231032301-2232203011012101-3032231210201230-0322210133333221-3201312130300303-3320200121222212-1030020330200011"></a>

<a id="canonical-1323203212002002-0221313310203023-0122102010130012-1032231313022123-3232031231302313-2122001320022103-2032033111020312-3131122033330101"></a>

#### `service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates.blindfold.policy` property

Type: `"string"`. Optional.

<a id="canonical-2230212031321301-0130313230012303-1310112013221233-3001121213121033-2311333101112013-2303231031020031-3023003033231021-2223010330202010"></a>

<a id="canonical-2320133211100131-3213320330211321-3010033012033312-2131221131131210-0223001333103201-0210330303001023-2032102231120221-3311232203333133"></a>

#### `service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates.blindfold.prepared_identity` property

Type: `"string"`. Computed.

<a id="canonical-0002300310102303-0222130333012232-0110110202322232-3131322322231032-2130302012302213-1113111010112133-0311220021203223-3130202211200322"></a>

<a id="canonical-0220012013203231-2233212211213201-3223023020323001-3322032312133100-3120113221111012-3002222013211233-1233133020032131-3330333111203333"></a>

#### `service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates.blindfold.private_key_file` property

Type: `"string"`. Optional.

<a id="canonical-3130102221221130-2233303132202213-3123113320133123-2023203102321033-3312322110211032-2110113331023021-3010010121320310-0102010122212032"></a>

<a id="canonical-3002031210220100-2303031210210313-1031110131022102-2321332130002332-2111102312011200-1012211001003022-1000300220113112-0031222113231122"></a>

#### `service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates.blindfold.private_key_wo` property

Type: `"string"`. Optional, Sensitive, Write_only.

<a id="canonical-1132121032320331-2200311122301311-0312322112012031-0311121000103113-1210201330011313-3321033230210310-0022031233012221-2230221132222310"></a>

<a id="canonical-2000232002032001-3333301002212132-3133123202202002-1231113221200303-0303200320230011-1330030321131232-1013330111200222-0221012010001033"></a>

#### `service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates.blindfold.spki_identity` property

Type: `"string"`. Computed.

<a id="canonical-0101000310022230-0310111032313202-3003331220301302-2201020020020300-1131310023233313-2302331220210031-3330130102113310-3030203030003013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates.custom_hash_algorithms` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-006.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-006.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-009.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.port](resources--workload--reference--group-012.md#canonical-3013031222022030-2021222233122213-2220022232301003-3102220000332302-0311213311310132-3133321213021123-0133300222212330-1122131003023301)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-012.md#canonical-0101000213001300-0222020002101313-1020132032013321-2231011111202333-2122231103331210-3202123132002020-0313332003123002-3101311011132330)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https](resources--workload--reference--group-013.md#canonical-2221000320213131-2001211030022032-2023131131113110-1000133011113313-1003112332120000-2102122231000321-0211212202023222-1022013120231111)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-013.md#canonical-1231003120133112-2013110130100211-2233121010211001-2121031322013232-2232220303133231-0300100233200310-0003012301120011-3102210111030310)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates](resources--workload--reference--group-013.md#canonical-1222323211330211-1131011013013011-0010113301022222-0300023112121103-0132020001012213-1113030220101112-1233001112311310-2321032212303310)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates.custom_hash_algorithms

<a id="canonical-0321231331323220-2102020231110333-1000113022002310-1202120210310021-1010001222300030-2321223310220230-3213000213113330-2323132233313213"></a>

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

<a id="canonical-0332030012233003-2201323100102303-1223331101223203-0001022030023131-0010300201321302-3012000200330002-1210100233113101-1310031203023020"></a>

### Direct properties for `service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates.custom_hash_algorithms`

<a id="canonical-0122203201120032-1232010232003320-2101320223011330-1203131002011002-3023311032102220-2013211100012013-0312121013023312-1331302010302101"></a>

#### `service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates.custom_hash_algorithms.hash_algorithms` property

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

<a id="canonical-1111013233010131-2221110113130120-2312121320312022-2303321110232001-2311233333330203-2303313001013112-3232103110122111-1013303023313130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates.disable_ocsp_stapling` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-006.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-006.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-009.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.port](resources--workload--reference--group-012.md#canonical-3013031222022030-2021222233122213-2220022232301003-3102220000332302-0311213311310132-3133321213021123-0133300222212330-1122131003023301)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-012.md#canonical-0101000213001300-0222020002101313-1020132032013321-2231011111202333-2122231103331210-3202123132002020-0313332003123002-3101311011132330)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https](resources--workload--reference--group-013.md#canonical-2221000320213131-2001211030022032-2023131131113110-1000133011113313-1003112332120000-2102122231000321-0211212202023222-1022013120231111)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-013.md#canonical-1231003120133112-2013110130100211-2233121010211001-2121031322013232-2232220303133231-0300100233200310-0003012301120011-3102210111030310)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates](resources--workload--reference--group-013.md#canonical-1222323211330211-1131011013013011-0010113301022222-0300023112121103-0132020001012213-1113030220101112-1233001112311310-2321032212303310)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates.disable_ocsp_stapling

<a id="canonical-1322110120001311-0321331020323322-0012332300033030-0000313332213001-3030003311322010-0111012213001112-2232321020211012-2221330231313010"></a>

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

<a id="canonical-2033112022013132-3022203212311013-1130313200322231-2212013001100010-0011112031131213-3002001011030201-0121133021012111-1012222332230300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates.private_key` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-006.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-006.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-009.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.port](resources--workload--reference--group-012.md#canonical-3013031222022030-2021222233122213-2220022232301003-3102220000332302-0311213311310132-3133321213021123-0133300222212330-1122131003023301)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-012.md#canonical-0101000213001300-0222020002101313-1020132032013321-2231011111202333-2122231103331210-3202123132002020-0313332003123002-3101311011132330)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https](resources--workload--reference--group-013.md#canonical-2221000320213131-2001211030022032-2023131131113110-1000133011113313-1003112332120000-2102122231000321-0211212202023222-1022013120231111)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-013.md#canonical-1231003120133112-2013110130100211-2233121010211001-2121031322013232-2232220303133231-0300100233200310-0003012301120011-3102210111030310)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates](resources--workload--reference--group-013.md#canonical-1222323211330211-1131011013013011-0010113301022222-0300023112121103-0132020001012213-1113030220101112-1233001112311310-2321032212303310)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates.private_key

<a id="canonical-0100320110120331-2231122221211030-1213230011202010-3120111330312121-0223111130233320-2013202110030112-3103120020333111-3032310132012200"></a>

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

<a id="canonical-3231333011102101-1012023223131010-0211021220000133-2210111130102021-1033021100121221-0321330130010111-1323133330231333-2001202220110310"></a>

### Direct properties for `service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates.private_key`

- [blindfold_secret_info](resources--workload--reference--group-013.md#canonical-3033213033331102-2331212230121200-2102112010212303-3211122333220203-3333010312010310-2122303011322130-1231220012011021-3322101003130221): complete subsection reference.

- [clear_secret_info](resources--workload--reference--group-013.md#canonical-1333331310102313-2013030320202312-1010332311033031-0111000100202230-0203010031133012-0220121113111321-2312030031200313-3021110122100222): complete subsection reference.

<a id="canonical-3033213033331102-2331212230121200-2102112010212303-3211122333220203-3333010312010310-2122303011322130-1231220012011021-3322101003130221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates.private_key.blindfold_secret_info` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-006.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-006.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-009.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.port](resources--workload--reference--group-012.md#canonical-3013031222022030-2021222233122213-2220022232301003-3102220000332302-0311213311310132-3133321213021123-0133300222212330-1122131003023301)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-012.md#canonical-0101000213001300-0222020002101313-1020132032013321-2231011111202333-2122231103331210-3202123132002020-0313332003123002-3101311011132330)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https](resources--workload--reference--group-013.md#canonical-2221000320213131-2001211030022032-2023131131113110-1000133011113313-1003112332120000-2102122231000321-0211212202023222-1022013120231111)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-013.md#canonical-1231003120133112-2013110130100211-2233121010211001-2121031322013232-2232220303133231-0300100233200310-0003012301120011-3102210111030310)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates](resources--workload--reference--group-013.md#canonical-1222323211330211-1131011013013011-0010113301022222-0300023112121103-0132020001012213-1113030220101112-1233001112311310-2321032212303310)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates.private_key](resources--workload--reference--group-013.md#canonical-2033112022013132-3022203212311013-1130313200322231-2212013001100010-0011112031131213-3002001011030201-0121133021012111-1012222332230300)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates.private_key.blindfold_secret_info

<a id="canonical-1010120210230320-1210233010222103-1303203103320211-2132222100002311-3023013201320130-0102320200221203-0010123023031111-0300202122033221"></a>

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

<a id="canonical-2322032231331012-1122201003313022-0033022202013122-1130110230202103-0010101131032111-1030123121322003-0323123101033312-3011321011031022"></a>

### Direct properties for `service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates.private_key.blindfold_secret_info`

<a id="canonical-0120123203302130-2301122230030030-2020201322210330-2022022232101023-1200331231213332-1031210302100312-1201000122323223-0311322022211132"></a>

#### `service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates.private_key.blindfold_secret_info.decryption_provider` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-1210201031322221-3201331023221121-0310023102110310-3302002220321301-1203120023201132-2010213010113333-3331121103232131-1201021002112221"></a>

<a id="canonical-3322321211331003-1330110020320321-2212030010233313-3012102022332010-3110003303111332-0133231222200302-1330323110031022-0012032130203200"></a>

#### `service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates.private_key.blindfold_secret_info.location` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-0131232003200112-2022023200132132-0321132123221302-3021103303102010-3031021323010130-3231112323000220-2023002100033132-0120331101122332"></a>

<a id="canonical-1222023333220023-1303001123120133-0231001302222321-1321001301220023-3102212321120332-0133112132020331-3320333002132122-0213120330123002"></a>

#### `service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates.private_key.blindfold_secret_info.store_provider` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-1333331310102313-2013030320202312-1010332311033031-0111000100202230-0203010031133012-0220121113111321-2312030031200313-3021110122100222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates.private_key.clear_secret_info` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-006.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-006.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-009.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.port](resources--workload--reference--group-012.md#canonical-3013031222022030-2021222233122213-2220022232301003-3102220000332302-0311213311310132-3133321213021123-0133300222212330-1122131003023301)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-012.md#canonical-0101000213001300-0222020002101313-1020132032013321-2231011111202333-2122231103331210-3202123132002020-0313332003123002-3101311011132330)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https](resources--workload--reference--group-013.md#canonical-2221000320213131-2001211030022032-2023131131113110-1000133011113313-1003112332120000-2102122231000321-0211212202023222-1022013120231111)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-013.md#canonical-1231003120133112-2013110130100211-2233121010211001-2121031322013232-2232220303133231-0300100233200310-0003012301120011-3102210111030310)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates](resources--workload--reference--group-013.md#canonical-1222323211330211-1131011013013011-0010113301022222-0300023112121103-0132020001012213-1113030220101112-1233001112311310-2321032212303310)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates.private_key](resources--workload--reference--group-013.md#canonical-2033112022013132-3022203212311013-1130313200322231-2212013001100010-0011112031131213-3002001011030201-0121133021012111-1012222332230300)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates.private_key.clear_secret_info

<a id="canonical-0332302333333323-0230103230300123-2213023133023120-1310133033223102-0332132030020232-1221121200031123-3122121322003130-2133113030230303"></a>

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

<a id="canonical-0002330003120200-3201302021100320-0230232130010313-0133012210320030-1023000022032022-1301313312112022-0100331033320220-0022312210132113"></a>

### Direct properties for `service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates.private_key.clear_secret_info`

<a id="canonical-2011231103212000-2310132112130030-1032011321022003-2330302133023330-1113112012210233-2122133021033302-3203331302231010-3312000103303101"></a>

#### `service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates.private_key.clear_secret_info.provider_ref` property

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-1101102233230223-3230103011322120-0010112020123200-3230133313322202-1030233111032110-0232310012020022-0011032133323322-3303312103330100"></a>

<a id="canonical-1322102312320100-1123200113023332-0312122031011232-0302213112103021-0233232012131212-2022203313031101-3101332303000110-1332210213020123"></a>

#### `service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates.private_key.clear_secret_info.url` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-3220122011200333-3001230313130220-3013223100030322-1021301010012321-0032223013233321-2102220020230330-1113232313100312-0123202032102002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates.use_system_defaults` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-006.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-006.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-009.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.port](resources--workload--reference--group-012.md#canonical-3013031222022030-2021222233122213-2220022232301003-3102220000332302-0311213311310132-3133321213021123-0133300222212330-1122131003023301)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-012.md#canonical-0101000213001300-0222020002101313-1020132032013321-2231011111202333-2122231103331210-3202123132002020-0313332003123002-3101311011132330)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https](resources--workload--reference--group-013.md#canonical-2221000320213131-2001211030022032-2023131131113110-1000133011113313-1003112332120000-2102122231000321-0211212202023222-1022013120231111)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-013.md#canonical-1231003120133112-2013110130100211-2233121010211001-2121031322013232-2232220303133231-0300100233200310-0003012301120011-3102210111030310)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates](resources--workload--reference--group-013.md#canonical-1222323211330211-1131011013013011-0010113301022222-0300023112121103-0132020001012213-1113030220101112-1233001112311310-2321032212303310)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates.use_system_defaults

<a id="canonical-3230123233203032-2313102000011032-0213222312012113-1232123332020322-0102302302120211-3001030220023113-2023212331011003-0202202012103310"></a>

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

<a id="canonical-3232132023002231-3222012312123031-0032313002232322-3220233203030300-3132010213033111-2003310330231321-1313101010130131-3002200003132000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_config` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-006.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-006.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-009.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.port](resources--workload--reference--group-012.md#canonical-3013031222022030-2021222233122213-2220022232301003-3102220000332302-0311213311310132-3133321213021123-0133300222212330-1122131003023301)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-012.md#canonical-0101000213001300-0222020002101313-1020132032013321-2231011111202333-2122231103331210-3202123132002020-0313332003123002-3101311011132330)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https](resources--workload--reference--group-013.md#canonical-2221000320213131-2001211030022032-2023131131113110-1000133011113313-1003112332120000-2102122231000321-0211212202023222-1022013120231111)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-013.md#canonical-1231003120133112-2013110130100211-2233121010211001-2121031322013232-2232220303133231-0300100233200310-0003012301120011-3102210111030310)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_config

<a id="canonical-1220332301003312-1233122021323002-0303211231320110-3323031021333300-3220131222203012-3223103003320023-3131112011213013-2230032033200002"></a>

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

<a id="canonical-0001112022133220-0000221122020330-2201232002111330-0031301213302330-3331211311302233-1101312303223202-1333033103003231-0311203323323011"></a>

### Direct properties for `service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_config`

- [custom_security](resources--workload--reference--group-013.md#canonical-0211300122120131-1230103220220122-3101233300131010-3323310312211331-0130312002033003-0121133203302202-2223123120303100-3132030100222223): complete subsection reference.

- [default_security](resources--workload--reference--group-014.md#canonical-2123321320011130-2322202022130203-1232022332123012-3322130130231120-1012322033132322-3123210033203303-0032110022320313-1100332130133333): complete subsection reference.

- [low_security](resources--workload--reference--group-014.md#canonical-1330120123231220-1110203203033310-3302303322300012-3132311131231021-1132323111013320-3011203111333023-0122222302023102-3133133022130130): complete subsection reference.

- [medium_security](resources--workload--reference--group-014.md#canonical-0223331203323102-0130301232100223-1233320210122122-1101303101322110-0130000102120321-3132233132301003-3031301003211113-2332000311331232): complete subsection reference.

<a id="canonical-0211300122120131-1230103220220122-3101233300131010-3323310312211331-0130312002033003-0121133203302202-2223123120303100-3132030100222223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_config.custom_security` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-006.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-006.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-009.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.port](resources--workload--reference--group-012.md#canonical-3013031222022030-2021222233122213-2220022232301003-3102220000332302-0311213311310132-3133321213021123-0133300222212330-1122131003023301)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-012.md#canonical-0101000213001300-0222020002101313-1020132032013321-2231011111202333-2122231103331210-3202123132002020-0313332003123002-3101311011132330)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https](resources--workload--reference--group-013.md#canonical-2221000320213131-2001211030022032-2023131131113110-1000133011113313-1003112332120000-2102122231000321-0211212202023222-1022013120231111)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-013.md#canonical-1231003120133112-2013110130100211-2233121010211001-2121031322013232-2232220303133231-0300100233200310-0003012301120011-3102210111030310)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_config](resources--workload--reference--group-013.md#canonical-3232132023002231-3222012312123031-0032313002232322-3220233203030300-3132010213033111-2003310330231321-1313101010130131-3002200003132000)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_config.custom_security

<a id="canonical-0012132312321120-0003111112012221-3203222110120002-2002030033113232-0023201213112103-3233331233311122-3131322210231213-1111021132321131"></a>

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

<a id="canonical-3321103132030110-2233133030013002-2333022003323302-3112302023233221-1223200133303321-3332301311223333-2000123311110330-3122310011112220"></a>

### Direct properties for `service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_config.custom_security`

<a id="canonical-3222013130101333-3011223233021300-3111113132321130-3231212111301212-3313330132312312-3133202320230321-3322311300213310-0110120022131233"></a>

#### `service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_config.custom_security.cipher_suites` property

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

<a id="canonical-1002120020111332-2111120220130021-1001211330332310-0300232030323002-3002312320131230-2133002213301313-1102230213121003-2111122011033110"></a>

<a id="canonical-1120220022332000-2133132111000022-3331303223020020-1303103113102330-3321311330023121-2121031200032212-0023221111322000-0122230311233001"></a>

#### `service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_config.custom_security.max_version` property

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

<a id="canonical-2320310022330331-2233031022101102-2010121133100230-3321201200323223-0030221103230210-2010111023103010-3103230301223001-3233223332123130"></a>

<a id="canonical-0020100212100002-0330230101113220-2303222130210111-1010131010011030-0233311220102221-0002023132011002-1111130302233020-3231301232113300"></a>

#### `service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_config.custom_security.min_version` property

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
