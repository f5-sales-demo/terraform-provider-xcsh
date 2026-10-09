---
page_title: "xcsh_workload reference"
subcategory: "Container"
description: "Complete grouped canonical reference for xcsh_workload reference."
---

# xcsh_workload reference

<a id="canonical-2123321320011130-2322202022130203-1232022332123012-3322130130231120-1012322033132322-3123210033203303-0032110022320313-1100332130133333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_config.default_security` properties

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
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_config.default_security

<a id="canonical-2011333210031312-3210311313110201-0133011113201113-1033303110023110-0230002201302012-3221130110330333-1031330011210031-1333012002102010"></a>

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

<a id="canonical-1330120123231220-1110203203033310-3302303322300012-3132311131231021-1132323111013320-3011203111333023-0122222302023102-3133133022130130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_config.low_security` properties

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
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_config.low_security

<a id="canonical-1100333012202101-0003303323302133-1121022202200203-1010331121110322-3232132201330030-3121010233032211-1210031232011000-1203221011101310"></a>

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

<a id="canonical-0223331203323102-0130301232100223-1233320210122122-1101303101322110-0130000102120321-3132233132301003-3031301003211113-2332000311331232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_config.medium_security` properties

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
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_config.medium_security

<a id="canonical-2203110323022321-2013221301131223-2331110313201310-3322001233132313-0321303020332201-3301111320101223-3323230202110302-2201112113102120"></a>

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

<a id="canonical-2212033132123331-1022023101100202-0012200211203323-3100333332232032-2000331210002203-3023131312000213-0233333201322123-2002202002102311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.use_mtls` properties

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
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.use_mtls

<a id="canonical-0221130110202120-3122233301112010-3300211002230021-0232132113300323-0012220211301213-2231113010200310-1221023103033323-0313032201120333"></a>

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

<a id="canonical-2322213233222312-3302200011233123-0021300300131031-1331020200320322-3222310300213221-1213320000003122-0221330030103132-2221110212120003"></a>

### Direct properties for `service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.use_mtls`

<a id="canonical-0203100121113031-1102133130002333-0310230020211132-3211332213223113-3320231321023310-0111102101013310-2320100332123202-0111203003132200"></a>

#### `service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.use_mtls.client_certificate_optional` property

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

- [crl](resources--workload--reference--group-014.md#canonical-2210332213122022-0313313112222210-1221311123222020-2103131201012322-0032232132103120-0020300233222020-2132032313233031-1030212012211222): complete subsection reference.

- [no_crl](resources--workload--reference--group-014.md#canonical-2202333001101320-0330320302311032-0333132202300303-2221212212201311-1130133000230200-3022121231230213-2101323223321010-0233222131221003): complete subsection reference.

- [trusted_ca](resources--workload--reference--group-014.md#canonical-0223032130012113-0123222002030201-2311113110033213-1322132113003023-2022310032321111-3112011122232333-2312202332122021-1031023001232131): complete subsection reference.

<a id="canonical-0200032010330221-2331033211310121-3003023032231111-0003120003233102-0223232021221023-2003120112322303-2013222003102312-1311011010222322"></a>

<a id="canonical-2231312123131331-0131212030320302-2332111010132310-1031111323111032-3220330013101021-3232100233020022-2100320333330022-3032000100310131"></a>

#### `service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.use_mtls.trusted_ca_url` property

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

- [xfcc_disabled](resources--workload--reference--group-014.md#canonical-0222231130000212-1111110033320213-2303103222021013-1223133022333133-3231123313100031-2033220320122033-0223201110120033-2313130203333122): complete subsection reference.

- [xfcc_options](resources--workload--reference--group-014.md#canonical-3033121221321332-3202231033112132-0211011302031212-1031131103230031-0220211002000010-3103211121112012-3312230113220012-2010011332222202): complete subsection reference.

<a id="canonical-2210332213122022-0313313112222210-1221311123222020-2103131201012322-0032232132103120-0020300233222020-2132032313233031-1030212012211222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.use_mtls.crl` properties

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
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.use_mtls](resources--workload--reference--group-014.md#canonical-2212033132123331-1022023101100202-0012200211203323-3100333332232032-2000331210002203-3023131312000213-0233333201322123-2002202002102311)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.use_mtls.crl

<a id="canonical-2120122230201011-2131101221130032-2023201200221322-0021210021021222-3013110123322123-3313323331023313-0302030232220100-0231203221100110"></a>

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

<a id="canonical-3330301302303302-0232012013111021-3113201232320001-1022310021010233-0233203333210330-0323032300313222-0301300033231211-1313003032211312"></a>

### Direct properties for `service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.use_mtls.crl`

<a id="canonical-2231103112032311-2101321030220123-3112022322200210-2123302130020101-2123001021111230-3200330202033320-0013030113133112-2211121032223211"></a>

#### `service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.use_mtls.crl.name` property

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

<a id="canonical-0202122011332100-0130021003213123-0233213101031231-2233322210233112-3331223302010030-0213100303212131-2032212212321231-0300221323330320"></a>

<a id="canonical-0333313322313312-2121103320033130-0320323133001322-1201020200231000-2132011322121112-0200312112230333-0222012101221231-0300013213310331"></a>

#### `service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.use_mtls.crl.namespace` property

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

<a id="canonical-0111010320132321-3232310121021122-0302121132331020-0003113123310311-2211321001220011-1232211033202311-1120310133222131-2322102102030101"></a>

<a id="canonical-0223030301332111-2321021230102310-3232002130310330-0203210003323033-0320101202233221-2210033212000003-3001321031023201-1330122121020133"></a>

#### `service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.use_mtls.crl.tenant` property

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

<a id="canonical-2202333001101320-0330320302311032-0333132202300303-2221212212201311-1130133000230200-3022121231230213-2101323223321010-0233222131221003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.use_mtls.no_crl` properties

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
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.use_mtls](resources--workload--reference--group-014.md#canonical-2212033132123331-1022023101100202-0012200211203323-3100333332232032-2000331210002203-3023131312000213-0233333201322123-2002202002102311)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.use_mtls.no_crl

<a id="canonical-1011102032321323-3331031331231020-1003313311112023-0211113233003002-1010123112213012-0030202010213013-3321013301122302-3032012222123000"></a>

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

<a id="canonical-0223032130012113-0123222002030201-2311113110033213-1322132113003023-2022310032321111-3112011122232333-2312202332122021-1031023001232131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.use_mtls.trusted_ca` properties

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
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.use_mtls](resources--workload--reference--group-014.md#canonical-2212033132123331-1022023101100202-0012200211203323-3100333332232032-2000331210002203-3023131312000213-0233333201322123-2002202002102311)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.use_mtls.trusted_ca

<a id="canonical-0211033101131210-1200021003202101-0121301320302320-1110101213011321-1030223000321001-3013211312002111-3322313301220111-2011031020331333"></a>

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

<a id="canonical-3010010123102100-1201201103023013-2213300020102111-3312202232121030-1002120231013323-0231111030223121-3101323301123322-0001101210333110"></a>

### Direct properties for `service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.use_mtls.trusted_ca`

<a id="canonical-1230321322100103-1233200320032321-0203332311100032-0122210333333321-2102322332312210-0233303021121320-3220131112202300-1102031212000230"></a>

#### `service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.use_mtls.trusted_ca.name` property

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

<a id="canonical-1320012203000221-3030132022303321-3021011322331033-1131203011021011-3103110221001221-2213112231332020-1220220303331310-3231023321232100"></a>

<a id="canonical-2333033021203011-3103132011113011-2103022123212203-3330310210113212-2121322330100000-2100113231121201-2302000201202310-0101120222213032"></a>

#### `service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.use_mtls.trusted_ca.namespace` property

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

<a id="canonical-0202213302021123-1332132132120202-2213013233321031-0302032223320100-3301132012033300-2131331221330323-1302201213320233-3013030132333120"></a>

<a id="canonical-1202300211130113-2022013113121131-3000311330121013-2222011223110103-1222233301230221-2323001333212002-1301212232010320-0220310012220223"></a>

#### `service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.use_mtls.trusted_ca.tenant` property

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

<a id="canonical-0222231130000212-1111110033320213-2303103222021013-1223133022333133-3231123313100031-2033220320122033-0223201110120033-2313130203333122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.use_mtls.xfcc_disabled` properties

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
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.use_mtls](resources--workload--reference--group-014.md#canonical-2212033132123331-1022023101100202-0012200211203323-3100333332232032-2000331210002203-3023131312000213-0233333201322123-2002202002102311)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.use_mtls.xfcc_disabled

<a id="canonical-2322232301122111-2313031321130122-2321223110002000-1003000123113332-2231122002230303-3233310022230210-0231031321010120-0322220111122322"></a>

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

<a id="canonical-3033121221321332-3202231033112132-0211011302031212-1031131103230031-0220211002000010-3103211121112012-3312230113220012-2010011332222202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.use_mtls.xfcc_options` properties

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
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.use_mtls](resources--workload--reference--group-014.md#canonical-2212033132123331-1022023101100202-0012200211203323-3100333332232032-2000331210002203-3023131312000213-0233333201322123-2002202002102311)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.use_mtls.xfcc_options

<a id="canonical-3200113033013223-1222011112122210-0311211133200113-1111200231332322-3232101202312212-0122102023010030-2122312220022303-3100201102012333"></a>

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

<a id="canonical-2122200222300320-1311321221220133-1201031112310121-3023322300133230-2201323222133223-3303011323211031-2313023320003020-0310113003001111"></a>

### Direct properties for `service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.use_mtls.xfcc_options`

<a id="canonical-0321331230223131-3300323320211232-3003313233000310-0100232201112300-2012120103101131-3212010331003212-2333010313332003-1221230022133302"></a>

#### `service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.use_mtls.xfcc_options.xfcc_header_elements` property

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

<a id="canonical-2233002113231233-1201011231003023-2310112000233003-0120011321233313-3122132210210203-2300131312231320-2323313110111333-1012021022013033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-006.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-006.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-009.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.port](resources--workload--reference--group-012.md#canonical-3013031222022030-2021222233122213-2220022232301003-3102220000332302-0311213311310132-3133321213021123-0133300222212330-1122131003023301)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-012.md#canonical-0101000213001300-0222020002101313-1020132032013321-2231011111202333-2122231103331210-3202123132002020-0313332003123002-3101311011132330)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert

<a id="canonical-2231322012213232-0101322002011231-3121130002130200-3331233223021112-0231223200303212-2302001200013232-3320333001313032-1233123330210211"></a>

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
  "x-ves-oneof-field-mtls_choice": "[\"no_mtls\",\"use_mtls\"]",
  "x-ves-oneof-field-path_normalize_choice": "[\"disable_path_normalize\",\"enable_path_normalize\"]",
  "x-ves-oneof-field-port_choice": "[\"port\",\"port_ranges\"]",
  "x-ves-oneof-field-server_header_choice": "[\"append_server_name\",\"default_header\",\"pass_through\",\"server_name\"]"
}
```

Terraform syntax:

```terraform
https_auto_cert {
  # Configure direct properties listed below.
}
```

<a id="canonical-0303202122211200-2222330031310331-1202303130231132-3101330112311030-3210333002010100-2100330303302032-0111202312001120-0122110002331220"></a>

### Direct properties for `service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert`

<a id="canonical-2102103231130022-2322001103332030-2003322133100213-0132231303100223-1023202010331113-1002230210110322-3101130022201011-1020120032012033"></a>

#### `service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.add_hsts` property

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

<a id="canonical-2103310210113331-3003112330032331-0023013123232311-3011200201013232-0133313020220302-3111103012232202-0130331121313102-1030133321221330"></a>

<a id="canonical-1303111101001133-0101321233000221-0021211313133223-2323110010102200-3033233102103203-1123132012212203-1231132012003110-3030112222021321"></a>

#### `service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.append_server_name` property

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

- [coalescing_options](resources--workload--reference--group-014.md#canonical-1003313312003320-1310013331003020-2311011202322301-3113122312223102-2131103011233132-3111030113231220-3122033212120010-1013330233300210): complete subsection reference.

<a id="canonical-1331233322220122-3200130201111002-3010332231000320-2132011312222113-3122120221310232-2111122300322120-2021103220233021-0333001120300010"></a>

<a id="canonical-1323332313332332-3103112032100212-0202330303032300-0311201223202212-1313102103031322-2233222020303231-0313222101033122-3021303130001010"></a>

#### `service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.connection_idle_timeout` property

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

- [default_header](resources--workload--reference--group-014.md#canonical-2210232101332333-2320302211212012-0121133211031312-1020002112100312-0231313002212211-0130311101101123-1130132213302100-1033033032022003): complete subsection reference.

- [default_loadbalancer](resources--workload--reference--group-014.md#canonical-1013031323132002-0301102203102311-1231111002032231-1330330000313013-2022121003132201-0311130220212331-3310111013003313-2323330120111211): complete subsection reference.

- [disable_path_normalize](resources--workload--reference--group-014.md#canonical-0102331301232000-2120130002202233-3301201131102112-2211202113012002-2103232303302002-2123020300003220-0103130020112000-3322111203012123): complete subsection reference.

- [enable_path_normalize](resources--workload--reference--group-014.md#canonical-0010230120211331-3320303220201011-3112231022311202-3011321001310331-3321100211030313-2303112320201212-0110021020322200-2022222330100133): complete subsection reference.

- [http_protocol_options](resources--workload--reference--group-014.md#canonical-1001201213213330-3123312003110132-1101112033021031-2320100223111332-3200320010031233-3303131213230332-0010000311321001-0111112233232201): complete subsection reference.

<a id="canonical-1121130100212011-1333231223310311-2231031231311023-0131132112200213-1030101200220210-1222232120101003-1213223213333232-1223132323302032"></a>

<a id="canonical-2012013122310022-0120310112112113-1203010111322013-1122200013100213-2033230013111320-1013133131332110-0332021130311203-3222030030333000"></a>

#### `service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.http_redirect` property

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

- [no_mtls](resources--workload--reference--group-014.md#canonical-2101100100013330-2110303301022210-3033032330202132-3323133032033332-1231211022300223-0222213220011312-0131120300002112-3222311031011122): complete subsection reference.

- [non_default_loadbalancer](resources--workload--reference--group-014.md#canonical-0311111211233021-1220112303132010-3020331022211111-3221012030333220-2111302132100212-3232300220311233-2031302323101202-3220221331130010): complete subsection reference.

- [pass_through](resources--workload--reference--group-014.md#canonical-0330111313320232-0201321223322103-0200232011221321-3232213133200013-3010132211110221-2131111101100033-0103232200333220-1131033112223003): complete subsection reference.

<a id="canonical-3120133322303013-0012123022020201-1112201331330002-1203000213321322-2102013333201132-0001113313122001-1332003302302001-3303122232232232"></a>

<a id="canonical-2321222233232330-1022312310120131-2132300320013112-0121210302333123-2221303323333333-3210302012113101-3331103122122020-1020133321212201"></a>

#### `service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.port` property

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

<a id="canonical-3303213303331131-3203132210231023-0000200032033201-1122020212031333-3002320131333212-1012230232221010-2212121331100033-0231001332222320"></a>

<a id="canonical-0312200312030131-3211312233323012-2321033220000302-0201213303002331-0120321213223031-0231321321210203-0222032323023300-1301112203200212"></a>

#### `service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.port_ranges` property

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

<a id="canonical-3003100321121133-2331330313030003-2032130222130102-1333311023012020-2031232200021322-2110133123232310-0233023310230310-3022233113331232"></a>

<a id="canonical-0131321030103330-2222011000312123-3230332313032033-0112222120102130-0110202103310111-2101010131210232-2203311133101330-2032233333100302"></a>

#### `service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.server_name` property

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

- [tls_config](resources--workload--reference--group-014.md#canonical-3332002222313202-3123321100213021-1310123312032013-0101102313123323-2003030031022013-2212223201221323-0203203103020311-2032023003330023): complete subsection reference.

- [use_mtls](resources--workload--reference--group-014.md#canonical-1101302220013000-1130323120002221-2213110022303302-1002100102000211-3203030111202113-1021233103221020-3020210323320220-3011000311311022): complete subsection reference.

<a id="canonical-1003313312003320-1310013331003020-2311011202322301-3113122312223102-2131103011233132-3111030113231220-3122033212120010-1013330233300210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.coalescing_options` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-006.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-006.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-009.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.port](resources--workload--reference--group-012.md#canonical-3013031222022030-2021222233122213-2220022232301003-3102220000332302-0311213311310132-3133321213021123-0133300222212330-1122131003023301)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-012.md#canonical-0101000213001300-0222020002101313-1020132032013321-2231011111202333-2122231103331210-3202123132002020-0313332003123002-3101311011132330)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert](resources--workload--reference--group-014.md#canonical-2233002113231233-1201011231003023-2310112000233003-0120011321233313-3122132210210203-2300131312231320-2323313110111333-1012021022013033)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.coalescing_options

<a id="canonical-2132121232223132-2331212203333322-3122201000100112-1201322331213032-1322001131231033-0122113320223031-0312301212001232-2313300032133130"></a>

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

<a id="canonical-1312013211203021-0112122302031120-3320010201203100-3100100211122001-1121312130321001-2221110313100010-2110002203301020-2333212123000312"></a>

### Direct properties for `service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.coalescing_options`

- [default_coalescing](resources--workload--reference--group-014.md#canonical-1002111202130012-2213212112200030-2231331233321001-3123133202012110-2111301012221231-0320131333132230-0113310203203222-2311213030231001): complete subsection reference.

- [strict_coalescing](resources--workload--reference--group-014.md#canonical-1211101112303033-2232132010132032-0013102332121220-2213112003331131-2233202030022220-1102233032130032-2210312030103100-3232211103332011): complete subsection reference.

<a id="canonical-1002111202130012-2213212112200030-2231331233321001-3123133202012110-2111301012221231-0320131333132230-0113310203203222-2311213030231001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.coalescing_options.default_coalescing` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-006.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-006.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-009.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.port](resources--workload--reference--group-012.md#canonical-3013031222022030-2021222233122213-2220022232301003-3102220000332302-0311213311310132-3133321213021123-0133300222212330-1122131003023301)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-012.md#canonical-0101000213001300-0222020002101313-1020132032013321-2231011111202333-2122231103331210-3202123132002020-0313332003123002-3101311011132330)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert](resources--workload--reference--group-014.md#canonical-2233002113231233-1201011231003023-2310112000233003-0120011321233313-3122132210210203-2300131312231320-2323313110111333-1012021022013033)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.coalescing_options](resources--workload--reference--group-014.md#canonical-1003313312003320-1310013331003020-2311011202322301-3113122312223102-2131103011233132-3111030113231220-3122033212120010-1013330233300210)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.coalescing_options.default_coalescing

<a id="canonical-3303213230301231-1203323003222010-3121120112110313-0012300131203201-2201210100022110-2232233313220301-1123030122323113-0122211132120030"></a>

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

<a id="canonical-1211101112303033-2232132010132032-0013102332121220-2213112003331131-2233202030022220-1102233032130032-2210312030103100-3232211103332011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.coalescing_options.strict_coalescing` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-006.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-006.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-009.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.port](resources--workload--reference--group-012.md#canonical-3013031222022030-2021222233122213-2220022232301003-3102220000332302-0311213311310132-3133321213021123-0133300222212330-1122131003023301)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-012.md#canonical-0101000213001300-0222020002101313-1020132032013321-2231011111202333-2122231103331210-3202123132002020-0313332003123002-3101311011132330)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert](resources--workload--reference--group-014.md#canonical-2233002113231233-1201011231003023-2310112000233003-0120011321233313-3122132210210203-2300131312231320-2323313110111333-1012021022013033)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.coalescing_options](resources--workload--reference--group-014.md#canonical-1003313312003320-1310013331003020-2311011202322301-3113122312223102-2131103011233132-3111030113231220-3122033212120010-1013330233300210)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.coalescing_options.strict_coalescing

<a id="canonical-3020110320132020-3132013233302021-3013211103311022-3012120012223312-0003001321332231-0130132233210203-0002032002133010-0133103013001333"></a>

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

<a id="canonical-2210232101332333-2320302211212012-0121133211031312-1020002112100312-0231313002212211-0130311101101123-1130132213302100-1033033032022003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.default_header` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-006.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-006.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-009.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.port](resources--workload--reference--group-012.md#canonical-3013031222022030-2021222233122213-2220022232301003-3102220000332302-0311213311310132-3133321213021123-0133300222212330-1122131003023301)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-012.md#canonical-0101000213001300-0222020002101313-1020132032013321-2231011111202333-2122231103331210-3202123132002020-0313332003123002-3101311011132330)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert](resources--workload--reference--group-014.md#canonical-2233002113231233-1201011231003023-2310112000233003-0120011321233313-3122132210210203-2300131312231320-2323313110111333-1012021022013033)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.default_header

<a id="canonical-0113000302100031-0012033011223031-3003332313232230-1211201202231220-3010013310113230-2311230010332312-1312102011211110-0322100020231222"></a>

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

<a id="canonical-1013031323132002-0301102203102311-1231111002032231-1330330000313013-2022121003132201-0311130220212331-3310111013003313-2323330120111211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.default_loadbalancer` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-006.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-006.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-009.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.port](resources--workload--reference--group-012.md#canonical-3013031222022030-2021222233122213-2220022232301003-3102220000332302-0311213311310132-3133321213021123-0133300222212330-1122131003023301)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-012.md#canonical-0101000213001300-0222020002101313-1020132032013321-2231011111202333-2122231103331210-3202123132002020-0313332003123002-3101311011132330)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert](resources--workload--reference--group-014.md#canonical-2233002113231233-1201011231003023-2310112000233003-0120011321233313-3122132210210203-2300131312231320-2323313110111333-1012021022013033)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.default_loadbalancer

<a id="canonical-2103013032030320-1001202221111330-1310133132121013-0311032113112212-1002223201130032-1001231320231233-3332023202333313-3131120012310011"></a>

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

<a id="canonical-0102331301232000-2120130002202233-3301201131102112-2211202113012002-2103232303302002-2123020300003220-0103130020112000-3322111203012123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.disable_path_normalize` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-006.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-006.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-009.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.port](resources--workload--reference--group-012.md#canonical-3013031222022030-2021222233122213-2220022232301003-3102220000332302-0311213311310132-3133321213021123-0133300222212330-1122131003023301)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-012.md#canonical-0101000213001300-0222020002101313-1020132032013321-2231011111202333-2122231103331210-3202123132002020-0313332003123002-3101311011132330)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert](resources--workload--reference--group-014.md#canonical-2233002113231233-1201011231003023-2310112000233003-0120011321233313-3122132210210203-2300131312231320-2323313110111333-1012021022013033)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.disable_path_normalize

<a id="canonical-1320022301001222-1331201122012320-1220013202011023-3130333132331202-1111331132130002-0330303020200220-3301121212133210-0200223310303030"></a>

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

<a id="canonical-0010230120211331-3320303220201011-3112231022311202-3011321001310331-3321100211030313-2303112320201212-0110021020322200-2022222330100133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.enable_path_normalize` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-006.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-006.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-009.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.port](resources--workload--reference--group-012.md#canonical-3013031222022030-2021222233122213-2220022232301003-3102220000332302-0311213311310132-3133321213021123-0133300222212330-1122131003023301)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-012.md#canonical-0101000213001300-0222020002101313-1020132032013321-2231011111202333-2122231103331210-3202123132002020-0313332003123002-3101311011132330)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert](resources--workload--reference--group-014.md#canonical-2233002113231233-1201011231003023-2310112000233003-0120011321233313-3122132210210203-2300131312231320-2323313110111333-1012021022013033)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.enable_path_normalize

<a id="canonical-2323001212202300-1010122003212013-1123203112012013-1220012331331200-0221102331012033-0101230200313103-0122200023113331-1031012300121112"></a>

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

<a id="canonical-1001201213213330-3123312003110132-1101112033021031-2320100223111332-3200320010031233-3303131213230332-0010000311321001-0111112233232201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.http_protocol_options` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-006.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-006.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-009.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.port](resources--workload--reference--group-012.md#canonical-3013031222022030-2021222233122213-2220022232301003-3102220000332302-0311213311310132-3133321213021123-0133300222212330-1122131003023301)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-012.md#canonical-0101000213001300-0222020002101313-1020132032013321-2231011111202333-2122231103331210-3202123132002020-0313332003123002-3101311011132330)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert](resources--workload--reference--group-014.md#canonical-2233002113231233-1201011231003023-2310112000233003-0120011321233313-3122132210210203-2300131312231320-2323313110111333-1012021022013033)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.http_protocol_options

<a id="canonical-1203032030111022-3303200022323320-0130230221323010-2300100333012232-0323122231303330-2300122000113111-2321003310303112-1222333121023011"></a>

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

<a id="canonical-0030313232112223-1133322030103032-3130210033231133-2310120023131201-0002000320023302-2100013211101020-2212302330201300-3303003213320302"></a>

### Direct properties for `service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.http_protocol_options`

- [http_protocol_enable_v1_only](resources--workload--reference--group-014.md#canonical-2101100120112220-3011131101223030-0201203113223302-2210200120000233-3120203233011300-3231001202310203-3003232210201001-1103101233013201): complete subsection reference.

- [http_protocol_enable_v1_v2](resources--workload--reference--group-014.md#canonical-3300220010231100-0213320032310313-1312000010103323-3231331211200020-1101311002031212-0233311330233132-1001213333023313-1021321003320213): complete subsection reference.

- [http_protocol_enable_v2_only](resources--workload--reference--group-014.md#canonical-0100221013330112-0132313101223121-0222202123322200-0330111332333000-0003030211321311-0003102020222111-0311110033003233-3122003023021003): complete subsection reference.

<a id="canonical-2101100120112220-3011131101223030-0201203113223302-2210200120000233-3120203233011300-3231001202310203-3003232210201001-1103101233013201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-006.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-006.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-009.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.port](resources--workload--reference--group-012.md#canonical-3013031222022030-2021222233122213-2220022232301003-3102220000332302-0311213311310132-3133321213021123-0133300222212330-1122131003023301)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-012.md#canonical-0101000213001300-0222020002101313-1020132032013321-2231011111202333-2122231103331210-3202123132002020-0313332003123002-3101311011132330)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert](resources--workload--reference--group-014.md#canonical-2233002113231233-1201011231003023-2310112000233003-0120011321233313-3122132210210203-2300131312231320-2323313110111333-1012021022013033)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.http_protocol_options](resources--workload--reference--group-014.md#canonical-1001201213213330-3123312003110132-1101112033021031-2320100223111332-3200320010031233-3303131213230332-0010000311321001-0111112233232201)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only

<a id="canonical-0120122331221022-1311333130022211-0033003333112202-0302210000212203-0331332300211202-2321231332212112-1231213133021320-1310000110312120"></a>

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

<a id="canonical-1110321312003300-0030303023102002-2031313003302312-3332231301030131-0123222230010230-3110301203222331-3311332232032002-2123302330012303"></a>

### Direct properties for `service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only`

- [header_transformation](resources--workload--reference--group-014.md#canonical-1122230222212111-3313211200333020-1030101021020320-1323301220200121-1223222331313330-3120223220212022-1213020202330022-1320022103011012): complete subsection reference.

<a id="canonical-1122230222212111-3313211200333020-1030101021020320-1323301220200121-1223222331313330-3120223220212022-1213020202330022-1320022103011012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-006.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-006.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-009.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.port](resources--workload--reference--group-012.md#canonical-3013031222022030-2021222233122213-2220022232301003-3102220000332302-0311213311310132-3133321213021123-0133300222212330-1122131003023301)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-012.md#canonical-0101000213001300-0222020002101313-1020132032013321-2231011111202333-2122231103331210-3202123132002020-0313332003123002-3101311011132330)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert](resources--workload--reference--group-014.md#canonical-2233002113231233-1201011231003023-2310112000233003-0120011321233313-3122132210210203-2300131312231320-2323313110111333-1012021022013033)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.http_protocol_options](resources--workload--reference--group-014.md#canonical-1001201213213330-3123312003110132-1101112033021031-2320100223111332-3200320010031233-3303131213230332-0010000311321001-0111112233232201)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only](resources--workload--reference--group-014.md#canonical-2101100120112220-3011131101223030-0201203113223302-2210200120000233-3120203233011300-3231001202310203-3003232210201001-1103101233013201)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation

<a id="canonical-2020220033023231-1330130031300202-1233011111323012-0022331121100001-0012333103322003-3102222202302013-3003300211120213-3111210312221013"></a>

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

<a id="canonical-3202123302302032-1232021311232102-1002100010020112-2120213011021003-2111130330121210-0300211312210000-3113111321031312-0223113123021111"></a>

### Direct properties for `service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation`

- [default_header_transformation](resources--workload--reference--group-014.md#canonical-2123222033112002-0221011213330211-1210102322220233-0233131030212333-1303300203220230-3012213113312211-1303121133030122-3112022103110210): complete subsection reference.

- [preserve_case_header_transformation](resources--workload--reference--group-014.md#canonical-2202321301212202-0233011133031113-0210330212201231-0002133033002100-3313331203010213-1102020022333233-0220222222222020-1202102303331020): complete subsection reference.

- [proper_case_header_transformation](resources--workload--reference--group-014.md#canonical-1020132113133332-0001332222021231-2200322331103100-0101230301313111-2011121330231311-3313310331210031-3200223013001002-2213300000200111): complete subsection reference.

<a id="canonical-2123222033112002-0221011213330211-1210102322220233-0233131030212333-1303300203220230-3012213113312211-1303121133030122-3112022103110210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.default_header_transformation` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-006.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-006.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-009.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.port](resources--workload--reference--group-012.md#canonical-3013031222022030-2021222233122213-2220022232301003-3102220000332302-0311213311310132-3133321213021123-0133300222212330-1122131003023301)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-012.md#canonical-0101000213001300-0222020002101313-1020132032013321-2231011111202333-2122231103331210-3202123132002020-0313332003123002-3101311011132330)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert](resources--workload--reference--group-014.md#canonical-2233002113231233-1201011231003023-2310112000233003-0120011321233313-3122132210210203-2300131312231320-2323313110111333-1012021022013033)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.http_protocol_options](resources--workload--reference--group-014.md#canonical-1001201213213330-3123312003110132-1101112033021031-2320100223111332-3200320010031233-3303131213230332-0010000311321001-0111112233232201)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only](resources--workload--reference--group-014.md#canonical-2101100120112220-3011131101223030-0201203113223302-2210200120000233-3120203233011300-3231001202310203-3003232210201001-1103101233013201)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation](resources--workload--reference--group-014.md#canonical-1122230222212111-3313211200333020-1030101021020320-1323301220200121-1223222331313330-3120223220212022-1213020202330022-1320022103011012)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.default_header_transformation

<a id="canonical-0303320201123312-2200112333233311-1213110313223112-1031323223021233-0223331221300232-2330013101210002-2130030301102102-0102322232011210"></a>

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

<a id="canonical-2202321301212202-0233011133031113-0210330212201231-0002133033002100-3313331203010213-1102020022333233-0220222222222020-1202102303331020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.preserve_case_header_transformation` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-006.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-006.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-009.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.port](resources--workload--reference--group-012.md#canonical-3013031222022030-2021222233122213-2220022232301003-3102220000332302-0311213311310132-3133321213021123-0133300222212330-1122131003023301)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-012.md#canonical-0101000213001300-0222020002101313-1020132032013321-2231011111202333-2122231103331210-3202123132002020-0313332003123002-3101311011132330)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert](resources--workload--reference--group-014.md#canonical-2233002113231233-1201011231003023-2310112000233003-0120011321233313-3122132210210203-2300131312231320-2323313110111333-1012021022013033)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.http_protocol_options](resources--workload--reference--group-014.md#canonical-1001201213213330-3123312003110132-1101112033021031-2320100223111332-3200320010031233-3303131213230332-0010000311321001-0111112233232201)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only](resources--workload--reference--group-014.md#canonical-2101100120112220-3011131101223030-0201203113223302-2210200120000233-3120203233011300-3231001202310203-3003232210201001-1103101233013201)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation](resources--workload--reference--group-014.md#canonical-1122230222212111-3313211200333020-1030101021020320-1323301220200121-1223222331313330-3120223220212022-1213020202330022-1320022103011012)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.preserve_case_header_transformation

<a id="canonical-3220332023312330-1300312023103323-3311120210230132-0301203030002111-1010310011201332-0231001211030230-0112333203030320-2313102322100121"></a>

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

<a id="canonical-1020132113133332-0001332222021231-2200322331103100-0101230301313111-2011121330231311-3313310331210031-3200223013001002-2213300000200111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.proper_case_header_transformation` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-006.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-006.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-009.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.port](resources--workload--reference--group-012.md#canonical-3013031222022030-2021222233122213-2220022232301003-3102220000332302-0311213311310132-3133321213021123-0133300222212330-1122131003023301)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-012.md#canonical-0101000213001300-0222020002101313-1020132032013321-2231011111202333-2122231103331210-3202123132002020-0313332003123002-3101311011132330)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert](resources--workload--reference--group-014.md#canonical-2233002113231233-1201011231003023-2310112000233003-0120011321233313-3122132210210203-2300131312231320-2323313110111333-1012021022013033)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.http_protocol_options](resources--workload--reference--group-014.md#canonical-1001201213213330-3123312003110132-1101112033021031-2320100223111332-3200320010031233-3303131213230332-0010000311321001-0111112233232201)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only](resources--workload--reference--group-014.md#canonical-2101100120112220-3011131101223030-0201203113223302-2210200120000233-3120203233011300-3231001202310203-3003232210201001-1103101233013201)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation](resources--workload--reference--group-014.md#canonical-1122230222212111-3313211200333020-1030101021020320-1323301220200121-1223222331313330-3120223220212022-1213020202330022-1320022103011012)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.proper_case_header_transformation

<a id="canonical-0031330011302102-0112030312013220-0023023322003023-2232000323313010-1123122220210010-0013221200332100-3203110210222220-3113003202012132"></a>

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

<a id="canonical-3300220010231100-0213320032310313-1312000010103323-3231331211200020-1101311002031212-0233311330233132-1001213333023313-1021321003320213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_v2` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-006.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-006.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-009.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.port](resources--workload--reference--group-012.md#canonical-3013031222022030-2021222233122213-2220022232301003-3102220000332302-0311213311310132-3133321213021123-0133300222212330-1122131003023301)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-012.md#canonical-0101000213001300-0222020002101313-1020132032013321-2231011111202333-2122231103331210-3202123132002020-0313332003123002-3101311011132330)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert](resources--workload--reference--group-014.md#canonical-2233002113231233-1201011231003023-2310112000233003-0120011321233313-3122132210210203-2300131312231320-2323313110111333-1012021022013033)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.http_protocol_options](resources--workload--reference--group-014.md#canonical-1001201213213330-3123312003110132-1101112033021031-2320100223111332-3200320010031233-3303131213230332-0010000311321001-0111112233232201)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_v2

<a id="canonical-1202023232133131-2101312001333320-1223211032031121-0333113301213303-3332021300300111-3230220200220103-0321011011200233-0131013001032301"></a>

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

<a id="canonical-0100221013330112-0132313101223121-0222202123322200-0330111332333000-0003030211321311-0003102020222111-0311110033003233-3122003023021003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v2_only` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-006.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-006.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-009.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.port](resources--workload--reference--group-012.md#canonical-3013031222022030-2021222233122213-2220022232301003-3102220000332302-0311213311310132-3133321213021123-0133300222212330-1122131003023301)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-012.md#canonical-0101000213001300-0222020002101313-1020132032013321-2231011111202333-2122231103331210-3202123132002020-0313332003123002-3101311011132330)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert](resources--workload--reference--group-014.md#canonical-2233002113231233-1201011231003023-2310112000233003-0120011321233313-3122132210210203-2300131312231320-2323313110111333-1012021022013033)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.http_protocol_options](resources--workload--reference--group-014.md#canonical-1001201213213330-3123312003110132-1101112033021031-2320100223111332-3200320010031233-3303131213230332-0010000311321001-0111112233232201)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v2_only

<a id="canonical-1102110020203231-3203031030010023-2121120030211132-2120200210013132-0333111001302111-2100232313202131-0021021100321233-0002001331331022"></a>

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

<a id="canonical-2101100100013330-2110303301022210-3033032330202132-3323133032033332-1231211022300223-0222213220011312-0131120300002112-3222311031011122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.no_mtls` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-006.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-006.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-009.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.port](resources--workload--reference--group-012.md#canonical-3013031222022030-2021222233122213-2220022232301003-3102220000332302-0311213311310132-3133321213021123-0133300222212330-1122131003023301)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-012.md#canonical-0101000213001300-0222020002101313-1020132032013321-2231011111202333-2122231103331210-3202123132002020-0313332003123002-3101311011132330)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert](resources--workload--reference--group-014.md#canonical-2233002113231233-1201011231003023-2310112000233003-0120011321233313-3122132210210203-2300131312231320-2323313110111333-1012021022013033)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.no_mtls

<a id="canonical-0212011202210030-3121231021310322-0321212311213022-2230202221321210-3200033320131312-0221012323303030-3122121233301110-0322032123312103"></a>

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

<a id="canonical-0311111211233021-1220112303132010-3020331022211111-3221012030333220-2111302132100212-3232300220311233-2031302323101202-3220221331130010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.non_default_loadbalancer` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-006.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-006.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-009.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.port](resources--workload--reference--group-012.md#canonical-3013031222022030-2021222233122213-2220022232301003-3102220000332302-0311213311310132-3133321213021123-0133300222212330-1122131003023301)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-012.md#canonical-0101000213001300-0222020002101313-1020132032013321-2231011111202333-2122231103331210-3202123132002020-0313332003123002-3101311011132330)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert](resources--workload--reference--group-014.md#canonical-2233002113231233-1201011231003023-2310112000233003-0120011321233313-3122132210210203-2300131312231320-2323313110111333-1012021022013033)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.non_default_loadbalancer

<a id="canonical-0213123000313321-2233023110111210-0322201323301222-0312201123002201-2323113121322010-1212313211031002-0123132102222223-2013013021122033"></a>

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

<a id="canonical-0330111313320232-0201321223322103-0200232011221321-3232213133200013-3010132211110221-2131111101100033-0103232200333220-1131033112223003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.pass_through` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-006.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-006.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-009.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.port](resources--workload--reference--group-012.md#canonical-3013031222022030-2021222233122213-2220022232301003-3102220000332302-0311213311310132-3133321213021123-0133300222212330-1122131003023301)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-012.md#canonical-0101000213001300-0222020002101313-1020132032013321-2231011111202333-2122231103331210-3202123132002020-0313332003123002-3101311011132330)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert](resources--workload--reference--group-014.md#canonical-2233002113231233-1201011231003023-2310112000233003-0120011321233313-3122132210210203-2300131312231320-2323313110111333-1012021022013033)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.pass_through

<a id="canonical-0213303320020330-0131333102301102-1333301013232233-3011322323231300-3032220100231111-3232330202233032-1121111132220000-2131121020303302"></a>

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

<a id="canonical-3332002222313202-3123321100213021-1310123312032013-0101102313123323-2003030031022013-2212223201221323-0203203103020311-2032023003330023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.tls_config` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-006.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-006.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-009.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.port](resources--workload--reference--group-012.md#canonical-3013031222022030-2021222233122213-2220022232301003-3102220000332302-0311213311310132-3133321213021123-0133300222212330-1122131003023301)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-012.md#canonical-0101000213001300-0222020002101313-1020132032013321-2231011111202333-2122231103331210-3202123132002020-0313332003123002-3101311011132330)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert](resources--workload--reference--group-014.md#canonical-2233002113231233-1201011231003023-2310112000233003-0120011321233313-3122132210210203-2300131312231320-2323313110111333-1012021022013033)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.tls_config

<a id="canonical-0031130232033031-1030133010310200-3220303202232022-1010202200221312-2323021121333100-0230010000010130-3130322131001122-3102330220021110"></a>

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

<a id="canonical-2331320301203031-1332311001102320-0023020112300130-2311320112001131-0220113331323003-1011200233313203-0203211132303213-2302313121311121"></a>

### Direct properties for `service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.tls_config`

- [custom_security](resources--workload--reference--group-014.md#canonical-0020310132012230-3303200221231333-1013303312100222-1023322323103110-2223213002033200-2231010122332102-3301210212100130-3130201203103332): complete subsection reference.

- [default_security](resources--workload--reference--group-014.md#canonical-2231213032120223-2103333333331022-2210201121323020-1012122232201112-2313311220202033-2230013022032001-0310113023020231-0100033120000023): complete subsection reference.

- [low_security](resources--workload--reference--group-014.md#canonical-0310211210032321-1233100200313323-1010013110311310-1031120212210013-3112023103300030-1033312233001333-3320310221100303-2133323030122033): complete subsection reference.

- [medium_security](resources--workload--reference--group-014.md#canonical-0002221021001331-3011122023010100-1022102030233023-0012111333012011-3121033122133231-3101112330101330-2232333033132302-0312300130202112): complete subsection reference.

<a id="canonical-0020310132012230-3303200221231333-1013303312100222-1023322323103110-2223213002033200-2231010122332102-3301210212100130-3130201203103332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.tls_config.custom_security` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-006.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-006.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-009.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.port](resources--workload--reference--group-012.md#canonical-3013031222022030-2021222233122213-2220022232301003-3102220000332302-0311213311310132-3133321213021123-0133300222212330-1122131003023301)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-012.md#canonical-0101000213001300-0222020002101313-1020132032013321-2231011111202333-2122231103331210-3202123132002020-0313332003123002-3101311011132330)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert](resources--workload--reference--group-014.md#canonical-2233002113231233-1201011231003023-2310112000233003-0120011321233313-3122132210210203-2300131312231320-2323313110111333-1012021022013033)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.tls_config](resources--workload--reference--group-014.md#canonical-3332002222313202-3123321100213021-1310123312032013-0101102313123323-2003030031022013-2212223201221323-0203203103020311-2032023003330023)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.tls_config.custom_security

<a id="canonical-0112030200133032-0033211000011021-1111003102313000-3101122302011032-1230033210021022-2333233002210033-0200020321023310-1221302311303021"></a>

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

<a id="canonical-3023120313323213-1113030111323111-0331013211313031-0003101321121020-3133202131233101-1323321013212002-2121301021302102-0031211113322021"></a>

### Direct properties for `service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.tls_config.custom_security`

<a id="canonical-0331220023301003-0121120021013212-0213323203303021-0331131012231100-1210032323013022-3020201213100332-1233012223131231-2121233221012201"></a>

#### `service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.tls_config.custom_security.cipher_suites` property

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

<a id="canonical-3100021003113331-1130222001132313-0111030230022113-2300122132322021-3100033303300100-2013203120101213-3333323001033011-2310320330112110"></a>

<a id="canonical-0012223121233130-3112101123221013-2333033301221311-3201132321002320-1310001133003110-2130222012322322-1232023021021212-0202000110303030"></a>

#### `service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.tls_config.custom_security.max_version` property

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

<a id="canonical-2013020123223310-0133012321111131-3030030012122023-0332210311323102-1212001221301021-2130330001031232-3221322322133011-2121033232113211"></a>

<a id="canonical-0123133031023131-1320123022103233-2010102223330211-1010201113020033-0322331202200021-2221203132012321-0131011322302202-2003101213112111"></a>

#### `service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.tls_config.custom_security.min_version` property

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

<a id="canonical-2231213032120223-2103333333331022-2210201121323020-1012122232201112-2313311220202033-2230013022032001-0310113023020231-0100033120000023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.tls_config.default_security` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-006.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-006.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-009.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.port](resources--workload--reference--group-012.md#canonical-3013031222022030-2021222233122213-2220022232301003-3102220000332302-0311213311310132-3133321213021123-0133300222212330-1122131003023301)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-012.md#canonical-0101000213001300-0222020002101313-1020132032013321-2231011111202333-2122231103331210-3202123132002020-0313332003123002-3101311011132330)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert](resources--workload--reference--group-014.md#canonical-2233002113231233-1201011231003023-2310112000233003-0120011321233313-3122132210210203-2300131312231320-2323313110111333-1012021022013033)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.tls_config](resources--workload--reference--group-014.md#canonical-3332002222313202-3123321100213021-1310123312032013-0101102313123323-2003030031022013-2212223201221323-0203203103020311-2032023003330023)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.tls_config.default_security

<a id="canonical-2033233022311233-1013223122321332-0002312132210121-2023002302120223-0310021022020232-1033011103003113-3010120021122020-2330010133213013"></a>

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

<a id="canonical-0310211210032321-1233100200313323-1010013110311310-1031120212210013-3112023103300030-1033312233001333-3320310221100303-2133323030122033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.tls_config.low_security` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-006.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-006.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-009.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.port](resources--workload--reference--group-012.md#canonical-3013031222022030-2021222233122213-2220022232301003-3102220000332302-0311213311310132-3133321213021123-0133300222212330-1122131003023301)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-012.md#canonical-0101000213001300-0222020002101313-1020132032013321-2231011111202333-2122231103331210-3202123132002020-0313332003123002-3101311011132330)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert](resources--workload--reference--group-014.md#canonical-2233002113231233-1201011231003023-2310112000233003-0120011321233313-3122132210210203-2300131312231320-2323313110111333-1012021022013033)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.tls_config](resources--workload--reference--group-014.md#canonical-3332002222313202-3123321100213021-1310123312032013-0101102313123323-2003030031022013-2212223201221323-0203203103020311-2032023003330023)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.tls_config.low_security

<a id="canonical-1322203330111003-1010302220201303-1030332021321320-1102210032331320-0030202030012020-1302013000131032-0232011201223322-2133210001311133"></a>

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

<a id="canonical-0002221021001331-3011122023010100-1022102030233023-0012111333012011-3121033122133231-3101112330101330-2232333033132302-0312300130202112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.tls_config.medium_security` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-006.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-006.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-009.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.port](resources--workload--reference--group-012.md#canonical-3013031222022030-2021222233122213-2220022232301003-3102220000332302-0311213311310132-3133321213021123-0133300222212330-1122131003023301)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-012.md#canonical-0101000213001300-0222020002101313-1020132032013321-2231011111202333-2122231103331210-3202123132002020-0313332003123002-3101311011132330)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert](resources--workload--reference--group-014.md#canonical-2233002113231233-1201011231003023-2310112000233003-0120011321233313-3122132210210203-2300131312231320-2323313110111333-1012021022013033)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.tls_config](resources--workload--reference--group-014.md#canonical-3332002222313202-3123321100213021-1310123312032013-0101102313123323-2003030031022013-2212223201221323-0203203103020311-2032023003330023)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.tls_config.medium_security

<a id="canonical-0112202220231030-1010202120111031-0033011323220121-2230121000132201-3230000013300122-3133201311013111-0212313100023111-1103100303123100"></a>

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

<a id="canonical-1101302220013000-1130323120002221-2213110022303302-1002100102000211-3203030111202113-1021233103221020-3020210323320220-3011000311311022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.use_mtls` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-006.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-006.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-009.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.port](resources--workload--reference--group-012.md#canonical-3013031222022030-2021222233122213-2220022232301003-3102220000332302-0311213311310132-3133321213021123-0133300222212330-1122131003023301)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-012.md#canonical-0101000213001300-0222020002101313-1020132032013321-2231011111202333-2122231103331210-3202123132002020-0313332003123002-3101311011132330)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert](resources--workload--reference--group-014.md#canonical-2233002113231233-1201011231003023-2310112000233003-0120011321233313-3122132210210203-2300131312231320-2323313110111333-1012021022013033)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.use_mtls

<a id="canonical-0313002030203123-3300220111123012-1130333122103202-2311223031332133-2320021302223213-1231303133202110-0112000003201322-3331123033001001"></a>

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

<a id="canonical-2331012033020012-2023101132030202-0011230121130010-1332210221012011-2030113201310211-3310020210200021-1321001211210022-2113300330222011"></a>

### Direct properties for `service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.use_mtls`

<a id="canonical-3302311300023032-3311313122033000-1202101113311212-3211132002131213-1021100230022203-3322211333101232-2311320321213103-0021102030200223"></a>

#### `service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.use_mtls.client_certificate_optional` property

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

- [crl](resources--workload--reference--group-014.md#canonical-3231001020202230-1311220132312310-0331130202123200-2211210110013210-0020001300323200-1002231323230303-3010330121132012-0133230312011032): complete subsection reference.

- [no_crl](resources--workload--reference--group-014.md#canonical-2132123100221201-0222212200310002-2212103010111231-0023100022313322-2013222101300233-3312221011213301-3303211220320111-3011321011113110): complete subsection reference.

- [trusted_ca](resources--workload--reference--group-014.md#canonical-2002231111221303-3223031203010300-3211032220031330-2221230212223312-1332123310322112-2330211323202210-2323130130021110-1012211031132101): complete subsection reference.

<a id="canonical-1012023111011201-3332132332310330-3000301211211123-0232221220232030-2212331212131212-0000302321000103-1322013300201303-1233233131230332"></a>

<a id="canonical-3133121223333011-1302233203320233-2010313121231313-2202233332332130-2011322031213113-1301133312102010-0113032000210302-2333111123122022"></a>

#### `service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.use_mtls.trusted_ca_url` property

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

- [xfcc_disabled](resources--workload--reference--group-014.md#canonical-2000220010332232-1002031302003133-0211300231112220-1010100213023031-3023123211220222-1210302312022033-1103002223021132-2133001333112230): complete subsection reference.

- [xfcc_options](resources--workload--reference--group-014.md#canonical-1232003011022332-1003223100200021-2212320311030031-3300123212003120-2202111030100301-2301302312313121-2312211132201022-2311103031021012): complete subsection reference.

<a id="canonical-3231001020202230-1311220132312310-0331130202123200-2211210110013210-0020001300323200-1002231323230303-3010330121132012-0133230312011032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.use_mtls.crl` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-006.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-006.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-009.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.port](resources--workload--reference--group-012.md#canonical-3013031222022030-2021222233122213-2220022232301003-3102220000332302-0311213311310132-3133321213021123-0133300222212330-1122131003023301)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-012.md#canonical-0101000213001300-0222020002101313-1020132032013321-2231011111202333-2122231103331210-3202123132002020-0313332003123002-3101311011132330)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert](resources--workload--reference--group-014.md#canonical-2233002113231233-1201011231003023-2310112000233003-0120011321233313-3122132210210203-2300131312231320-2323313110111333-1012021022013033)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.use_mtls](resources--workload--reference--group-014.md#canonical-1101302220013000-1130323120002221-2213110022303302-1002100102000211-3203030111202113-1021233103221020-3020210323320220-3011000311311022)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.use_mtls.crl

<a id="canonical-3032021211121111-2023111013333222-0020122330000133-2103330322333023-0300211033210320-3201321321102323-0300102000333222-3012311232331232"></a>

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

<a id="canonical-2013202301321101-0022301333303123-3223230233110310-0333121110212333-0010003101100030-2122031010133323-1300113322210032-1201110330031030"></a>

### Direct properties for `service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.use_mtls.crl`

<a id="canonical-3300112023300331-1231200011022311-2002033301020232-3221202002332032-3331320223331223-0032312020102330-1223321111110202-3223003202201120"></a>

#### `service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.use_mtls.crl.name` property

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

<a id="canonical-3133302032020131-0212020312020112-1131300232303320-2133032210031123-0212002212120221-2202331012202121-2022230012022303-2021323011232013"></a>

<a id="canonical-0202220122301321-2132010323123111-3120123003020301-2311201210102112-0133312112022103-2302210311131233-3320213110213132-1021000032120122"></a>

#### `service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.use_mtls.crl.namespace` property

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

<a id="canonical-3131020000110303-0003000023120123-0320313310311223-1132300132232332-0021103101023121-1020212232133331-3330212103322020-1200121223302010"></a>

<a id="canonical-1331213113033322-2200112203210030-1013220300302212-0121123020200001-2330032211301133-0320311022232232-3030103301100021-1000322330303321"></a>

#### `service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.use_mtls.crl.tenant` property

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

<a id="canonical-2132123100221201-0222212200310002-2212103010111231-0023100022313322-2013222101300233-3312221011213301-3303211220320111-3011321011113110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.use_mtls.no_crl` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-006.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-006.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-009.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.port](resources--workload--reference--group-012.md#canonical-3013031222022030-2021222233122213-2220022232301003-3102220000332302-0311213311310132-3133321213021123-0133300222212330-1122131003023301)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-012.md#canonical-0101000213001300-0222020002101313-1020132032013321-2231011111202333-2122231103331210-3202123132002020-0313332003123002-3101311011132330)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert](resources--workload--reference--group-014.md#canonical-2233002113231233-1201011231003023-2310112000233003-0120011321233313-3122132210210203-2300131312231320-2323313110111333-1012021022013033)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.use_mtls](resources--workload--reference--group-014.md#canonical-1101302220013000-1130323120002221-2213110022303302-1002100102000211-3203030111202113-1021233103221020-3020210323320220-3011000311311022)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.use_mtls.no_crl

<a id="canonical-3111323200130133-2210012011121111-0030223030003222-1323113322222330-2112323031200233-3311133102130301-0131323020301033-2121020230233023"></a>

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

<a id="canonical-2002231111221303-3223031203010300-3211032220031330-2221230212223312-1332123310322112-2330211323202210-2323130130021110-1012211031132101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.use_mtls.trusted_ca` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-006.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-006.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-009.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.port](resources--workload--reference--group-012.md#canonical-3013031222022030-2021222233122213-2220022232301003-3102220000332302-0311213311310132-3133321213021123-0133300222212330-1122131003023301)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-012.md#canonical-0101000213001300-0222020002101313-1020132032013321-2231011111202333-2122231103331210-3202123132002020-0313332003123002-3101311011132330)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert](resources--workload--reference--group-014.md#canonical-2233002113231233-1201011231003023-2310112000233003-0120011321233313-3122132210210203-2300131312231320-2323313110111333-1012021022013033)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.use_mtls](resources--workload--reference--group-014.md#canonical-1101302220013000-1130323120002221-2213110022303302-1002100102000211-3203030111202113-1021233103221020-3020210323320220-3011000311311022)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.use_mtls.trusted_ca

<a id="canonical-2133320023111123-0001321323121023-3232000110313200-3320032022313330-1222302220132002-2011202320333200-0311212311000012-3202330110130222"></a>

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

<a id="canonical-2130100320112001-3202023023222220-3030210132300311-2332301312102133-2333023311332100-3031032023300211-1000121333032132-2322300223101022"></a>

### Direct properties for `service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.use_mtls.trusted_ca`

<a id="canonical-0221112200233121-2010010130010332-1320221123320210-2101011330131332-1331332110221120-1100100031210032-0211031103223223-0213011321003231"></a>

#### `service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.use_mtls.trusted_ca.name` property

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

<a id="canonical-0213301231111103-1103100030220021-3002231013133102-2120003332232111-2100100323012133-1312010120102330-0303313012133000-2012221010022011"></a>

<a id="canonical-0130201321002130-1223331223231233-1020030302320203-3301120130332010-2103212113002310-3122123320233313-2332012323133321-2303123010120123"></a>

#### `service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.use_mtls.trusted_ca.namespace` property

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

<a id="canonical-1210131310010231-1110031210013220-2230000311133233-3201003302111010-2203211223330022-2201103131002333-1300323010300132-2111001131103223"></a>

<a id="canonical-2130320032112203-3330232123302013-1203223231023121-2203032303312302-3201022113013320-1320232220033323-1303302203310003-0012322223332222"></a>

#### `service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.use_mtls.trusted_ca.tenant` property

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

<a id="canonical-2000220010332232-1002031302003133-0211300231112220-1010100213023031-3023123211220222-1210302312022033-1103002223021132-2133001333112230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.use_mtls.xfcc_disabled` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-006.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-006.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-009.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.port](resources--workload--reference--group-012.md#canonical-3013031222022030-2021222233122213-2220022232301003-3102220000332302-0311213311310132-3133321213021123-0133300222212330-1122131003023301)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-012.md#canonical-0101000213001300-0222020002101313-1020132032013321-2231011111202333-2122231103331210-3202123132002020-0313332003123002-3101311011132330)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert](resources--workload--reference--group-014.md#canonical-2233002113231233-1201011231003023-2310112000233003-0120011321233313-3122132210210203-2300131312231320-2323313110111333-1012021022013033)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.use_mtls](resources--workload--reference--group-014.md#canonical-1101302220013000-1130323120002221-2213110022303302-1002100102000211-3203030111202113-1021233103221020-3020210323320220-3011000311311022)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.use_mtls.xfcc_disabled

<a id="canonical-1013213023100230-1130101112001032-1332111123302313-1033202020230132-1302131213202002-2322303122132121-2212231103032100-1213120003123110"></a>

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

<a id="canonical-1232003011022332-1003223100200021-2212320311030031-3300123212003120-2202111030100301-2301302312313121-2312211132201022-2311103031021012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.use_mtls.xfcc_options` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-006.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-006.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-009.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.port](resources--workload--reference--group-012.md#canonical-3013031222022030-2021222233122213-2220022232301003-3102220000332302-0311213311310132-3133321213021123-0133300222212330-1122131003023301)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-012.md#canonical-0101000213001300-0222020002101313-1020132032013321-2231011111202333-2122231103331210-3202123132002020-0313332003123002-3101311011132330)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert](resources--workload--reference--group-014.md#canonical-2233002113231233-1201011231003023-2310112000233003-0120011321233313-3122132210210203-2300131312231320-2323313110111333-1012021022013033)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.use_mtls](resources--workload--reference--group-014.md#canonical-1101302220013000-1130323120002221-2213110022303302-1002100102000211-3203030111202113-1021233103221020-3020210323320220-3011000311311022)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.use_mtls.xfcc_options

<a id="canonical-3231331312011111-3123311310220331-2301022131223110-2231313300122032-2102033311110022-0032302002200211-3003031220003130-2332010131210003"></a>

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

<a id="canonical-0312020322231230-1333210320033311-1010312002121303-0310132010300103-3002333200123132-0003322133303201-0023211013332202-2122211111133301"></a>

### Direct properties for `service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.use_mtls.xfcc_options`

<a id="canonical-1320103023112300-3003221120012021-1023323222100233-2321313232031323-2000101000030303-1333222003213200-0300302000002310-3221300022222202"></a>

#### `service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.use_mtls.xfcc_options.xfcc_header_elements` property

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

<a id="canonical-1223131333131301-1101202200212003-2303000010310012-3003131221200020-0201223130203322-2303023001222013-1231100011211030-2032330121100311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-006.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-006.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-009.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.port](resources--workload--reference--group-012.md#canonical-3013031222022030-2021222233122213-2220022232301003-3102220000332302-0311213311310132-3133321213021123-0133300222212330-1122131003023301)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-012.md#canonical-0101000213001300-0222020002101313-1020132032013321-2231011111202333-2122231103331210-3202123132002020-0313332003123002-3101311011132330)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes

<a id="canonical-1120303110123101-1303103010332000-2200322120000300-1032013130201110-1110312101200102-2213232203101011-0113103033001031-3020222323212030"></a>

Type: `"object"`. single nested block, Optional.

This defines various OPTIONS to define a route.

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
specific_routes {
  # Configure direct properties listed below.
}
```

<a id="canonical-3223203011102233-2100300021001230-3002023010133220-0003113301112303-3302213320301030-2333023110231312-3320132221203002-3301311301300110"></a>

### Direct properties for `service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes`

- [routes](resources--workload--reference--group-014.md#canonical-1320302000201200-1013121233121332-1200303031313333-2131200212100310-2031222303000133-0013120300210320-3032201103202020-0103102312330232): complete subsection reference.

<a id="canonical-1320302000201200-1013121233121332-1200303031313333-2131200212100310-2031222303000133-0013120300210320-3032201103202020-0103102312330232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-006.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-006.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-009.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.port](resources--workload--reference--group-012.md#canonical-3013031222022030-2021222233122213-2220022232301003-3102220000332302-0311213311310132-3133321213021123-0133300222212330-1122131003023301)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-012.md#canonical-0101000213001300-0222020002101313-1020132032013321-2231011111202333-2122231103331210-3202123132002020-0313332003123002-3101311011132330)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes](resources--workload--reference--group-014.md#canonical-1223131333131301-1101202200212003-2303000010310012-3003131221200020-0201223130203322-2303023001222013-1231100011211030-2032330121100311)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes

<a id="canonical-0122003303322123-0131032103031322-2032320013303211-1333013223303210-1312020211223331-1103101123233312-2030112302031232-2101121111021302"></a>

Type: `"object"`. list nested block, Optional.

Routes. Routes for this loadbalancer.

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
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "32"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "32"
  }
}
```

Terraform syntax:

```terraform
routes {
  # Configure direct properties listed below.
}
```

<a id="canonical-1333301300012222-3332311202103323-3220111112033013-2323202001312333-3330221000333320-1323211003133210-2030321020031122-2012011013223233"></a>

### Direct properties for `service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes`

- [custom_route_object](resources--workload--reference--group-014.md#canonical-0210032211120033-1031103021121212-2103210321113312-3223110300113221-2120130213202222-0311013133210132-3032003113103210-0013130303202222): complete subsection reference.

- [direct_response_route](resources--workload--reference--group-014.md#canonical-3323201111322002-1033013210321030-1202233310120130-0103032213302001-0022102310321331-0300303311223133-3133223031000013-0100120001003301): complete subsection reference.

- [redirect_route](resources--workload--reference--group-015.md#canonical-3222223113112223-3211013120021223-1311101201320313-1332203113201030-0200233021032333-2313310201233223-3321131333112311-2102210333302332): complete subsection reference.

- [simple_route](resources--workload--reference--group-015.md#canonical-3130302330333200-3231020322000031-3022221013222100-1033010302022131-0310323013110100-0001322111112133-1103302220233331-0223232310322020): complete subsection reference.

<a id="canonical-0210032211120033-1031103021121212-2103210321113312-3223110300113221-2120130213202222-0311013133210132-3032003113103210-0013130303202222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.custom_route_object` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-006.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-006.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-009.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.port](resources--workload--reference--group-012.md#canonical-3013031222022030-2021222233122213-2220022232301003-3102220000332302-0311213311310132-3133321213021123-0133300222212330-1122131003023301)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-012.md#canonical-0101000213001300-0222020002101313-1020132032013321-2231011111202333-2122231103331210-3202123132002020-0313332003123002-3101311011132330)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes](resources--workload--reference--group-014.md#canonical-1223131333131301-1101202200212003-2303000010310012-3003131221200020-0201223130203322-2303023001222013-1231100011211030-2032330121100311)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes](resources--workload--reference--group-014.md#canonical-1320302000201200-1013121233121332-1200303031313333-2131200212100310-2031222303000133-0013120300210320-3032201103202020-0103102312330232)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.custom_route_object

<a id="canonical-0213233101120032-0100322222331032-1001232220331030-2021022030220013-3312303033130233-1232220132120320-2231110210000330-2300213020120032"></a>

Type: `"object"`. single nested block, Optional.

A custom route uses a route object created outside of this view.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-caching": "[\"caching_disable\",\"caching_inherit\"]"
}
```

Terraform syntax:

```terraform
custom_route_object {
  # Configure direct properties listed below.
}
```

<a id="canonical-1202210033302313-3121121132013200-3003330303131332-2323010210222022-3120010331032000-1311131012223130-0021000111230320-1030103111331032"></a>

### Direct properties for `service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.custom_route_object`

- [caching_disable](resources--workload--reference--group-014.md#canonical-3111222103212100-1113220231310310-0101303230011001-3221133320310301-3202210021010200-1331123001332300-1030113101021211-1112003311030112): complete subsection reference.

- [caching_inherit](resources--workload--reference--group-014.md#canonical-3101332020003300-1110021321130122-3201102111322210-2331230303231302-1021322212032013-3302320312032331-1123231023311003-1003000310200203): complete subsection reference.

- [route_ref](resources--workload--reference--group-014.md#canonical-0320120120232103-2132323323112122-2222100123100112-2232031322203123-3303103022123020-3121012232332000-1131303033022210-1303011201223013): complete subsection reference.

<a id="canonical-3111222103212100-1113220231310310-0101303230011001-3221133320310301-3202210021010200-1331123001332300-1030113101021211-1112003311030112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.custom_route_object.caching_disable` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-006.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-006.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-009.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.port](resources--workload--reference--group-012.md#canonical-3013031222022030-2021222233122213-2220022232301003-3102220000332302-0311213311310132-3133321213021123-0133300222212330-1122131003023301)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-012.md#canonical-0101000213001300-0222020002101313-1020132032013321-2231011111202333-2122231103331210-3202123132002020-0313332003123002-3101311011132330)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes](resources--workload--reference--group-014.md#canonical-1223131333131301-1101202200212003-2303000010310012-3003131221200020-0201223130203322-2303023001222013-1231100011211030-2032330121100311)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes](resources--workload--reference--group-014.md#canonical-1320302000201200-1013121233121332-1200303031313333-2131200212100310-2031222303000133-0013120300210320-3032201103202020-0103102312330232)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.custom_route_object](resources--workload--reference--group-014.md#canonical-0210032211120033-1031103021121212-2103210321113312-3223110300113221-2120130213202222-0311013133210132-3032003113103210-0013130303202222)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.custom_route_object.caching_disable

<a id="canonical-2113213112211230-1031311220123122-3332012323131323-3312101203323110-1231020033010311-2002130023032201-1132111132102131-1203102332332201"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for caching disable.

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
caching_disable = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3101332020003300-1110021321130122-3201102111322210-2331230303231302-1021322212032013-3302320312032331-1123231023311003-1003000310200203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.custom_route_object.caching_inherit` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-006.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-006.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-009.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.port](resources--workload--reference--group-012.md#canonical-3013031222022030-2021222233122213-2220022232301003-3102220000332302-0311213311310132-3133321213021123-0133300222212330-1122131003023301)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-012.md#canonical-0101000213001300-0222020002101313-1020132032013321-2231011111202333-2122231103331210-3202123132002020-0313332003123002-3101311011132330)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes](resources--workload--reference--group-014.md#canonical-1223131333131301-1101202200212003-2303000010310012-3003131221200020-0201223130203322-2303023001222013-1231100011211030-2032330121100311)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes](resources--workload--reference--group-014.md#canonical-1320302000201200-1013121233121332-1200303031313333-2131200212100310-2031222303000133-0013120300210320-3032201103202020-0103102312330232)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.custom_route_object](resources--workload--reference--group-014.md#canonical-0210032211120033-1031103021121212-2103210321113312-3223110300113221-2120130213202222-0311013133210132-3032003113103210-0013130303202222)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.custom_route_object.caching_inherit

<a id="canonical-0212112011230133-2331011032033023-2020000301220221-2321022023231132-0011121033320001-1000323230310030-2103101110201132-0133212133003222"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for caching inherit.

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
caching_inherit = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0320120120232103-2132323323112122-2222100123100112-2232031322203123-3303103022123020-3121012232332000-1131303033022210-1303011201223013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.custom_route_object.route_ref` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-006.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-006.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-009.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.port](resources--workload--reference--group-012.md#canonical-3013031222022030-2021222233122213-2220022232301003-3102220000332302-0311213311310132-3133321213021123-0133300222212330-1122131003023301)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-012.md#canonical-0101000213001300-0222020002101313-1020132032013321-2231011111202333-2122231103331210-3202123132002020-0313332003123002-3101311011132330)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes](resources--workload--reference--group-014.md#canonical-1223131333131301-1101202200212003-2303000010310012-3003131221200020-0201223130203322-2303023001222013-1231100011211030-2032330121100311)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes](resources--workload--reference--group-014.md#canonical-1320302000201200-1013121233121332-1200303031313333-2131200212100310-2031222303000133-0013120300210320-3032201103202020-0103102312330232)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.custom_route_object](resources--workload--reference--group-014.md#canonical-0210032211120033-1031103021121212-2103210321113312-3223110300113221-2120130213202222-0311013133210132-3032003113103210-0013130303202222)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.custom_route_object.route_ref

<a id="canonical-1303032312321031-1213202032221123-0002122031132223-0221332131013001-3321202021310221-2032013100230322-2332031313323321-0220102213003113"></a>

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
route_ref {
  # Configure direct properties listed below.
}
```

<a id="canonical-2101132212110231-0100130201103321-1230203211130011-1332013010313033-2123103032122013-1103020313222321-0020303313220202-3101301332101020"></a>

### Direct properties for `service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.custom_route_object.route_ref`

<a id="canonical-2020300131202133-0310233031230020-1203211310332121-2001311123222110-3112010321003323-1022030233322023-2020213230230111-0120330222312100"></a>

#### `service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.custom_route_object.route_ref.name` property

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

<a id="canonical-3310332123100231-0232030220200033-1220331232111112-2231111331120203-0101322132313322-0311122201122310-3003011020213020-3132120101331130"></a>

<a id="canonical-2110011132333313-0100313032003130-1013213133230132-3223101331032110-1320103012002312-2012233221220330-0201110132002110-0212000313333031"></a>

#### `service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.custom_route_object.route_ref.namespace` property

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

<a id="canonical-1232021013001100-2133221120112222-1213230223221313-3010201013013223-0003010013221302-2031233310121123-3231031233212203-1102303133331220"></a>

<a id="canonical-0303022002030012-0313022202210200-2010112110013003-0212302303002203-1003032011302311-3133003310033120-0333030211122033-2111222322321320"></a>

#### `service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.custom_route_object.route_ref.tenant` property

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

<a id="canonical-3323201111322002-1033013210321030-1202233310120130-0103032213302001-0022102310321331-0300303311223133-3133223031000013-0100120001003301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.direct_response_route` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-006.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-006.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-009.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.port](resources--workload--reference--group-012.md#canonical-3013031222022030-2021222233122213-2220022232301003-3102220000332302-0311213311310132-3133321213021123-0133300222212330-1122131003023301)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-012.md#canonical-0101000213001300-0222020002101313-1020132032013321-2231011111202333-2122231103331210-3202123132002020-0313332003123002-3101311011132330)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes](resources--workload--reference--group-014.md#canonical-1223131333131301-1101202200212003-2303000010310012-3003131221200020-0201223130203322-2303023001222013-1231100011211030-2032330121100311)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes](resources--workload--reference--group-014.md#canonical-1320302000201200-1013121233121332-1200303031313333-2131200212100310-2031222303000133-0013120300210320-3032201103202020-0103102312330232)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.direct_response_route

<a id="canonical-3101311033312122-0300131220330000-3333123313012332-1001000120013113-1131000023222132-3332001120223321-0323221202301310-1211233233320013"></a>

Type: `"object"`. single nested block, Optional.

A direct response route matches on path, incoming header, incoming port and/or HTTP method and
responds directly to the matching traffic.

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
direct_response_route {
  # Configure direct properties listed below.
}
```

<a id="canonical-0211030210013210-0220021303011123-0132131122212311-1213111201010312-1001101312011121-3031103113112330-1032000122210321-0112030013303122"></a>

### Direct properties for `service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.direct_response_route`

- [headers](resources--workload--reference--group-014.md#canonical-0222010010121030-2002121221111330-0210203213300203-0132213001033020-0012203131231310-1223010112313230-3120022201313120-1312230321223032): complete subsection reference.

<a id="canonical-1323321221012323-3130301010301310-1130333030332102-3301023332031123-0302131331023113-2130312310001211-2103331120110103-3022131022222202"></a>

<a id="canonical-2230123102311212-2212303203322332-0323211133313010-1201213211130300-2320321033021231-2231320021032320-2220010330000011-2323322210313021"></a>

#### `service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.direct_response_route.http_method` property

Type: `"string"`. Optional.

\[Enum: ANY|GET|HEAD|POST|PUT|DELETE|CONNECT|OPTIONS|TRACE|PATCH|COPY\] Specifies the HTTP method
used to access a resource. Any HTTP Method. Possible values are \`ANY\`, \`GET\`, \`HEAD\`,
\`POST\`, \`PUT\`, \`DELETE\`, \`CONNECT\`, \`OPTIONS\`, \`TRACE\`, \`PATCH\`, \`COPY\`. Defaults to
\`ANY\`.

Receipt-pinned upstream constraints:

```json
{
  "default": "ANY",
  "enum": [
    "ANY",
    "GET",
    "HEAD",
    "POST",
    "PUT",
    "DELETE",
    "CONNECT",
    "OPTIONS",
    "TRACE",
    "PATCH",
    "COPY"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [incoming_port](resources--workload--reference--group-015.md#canonical-0120013202023333-3320333210330311-2301200302022321-2012232001023312-0310301031010320-1321213323003001-3100201302222113-1201220311213122): complete subsection reference.

- [path](resources--workload--reference--group-015.md#canonical-2011030213320012-1221020120303121-1223221102302032-3020312130010032-3003103231033033-3301032122222001-0003232231020002-1330331122201233): complete subsection reference.

- [route_direct_response](resources--workload--reference--group-015.md#canonical-2103100132103010-2113121013301320-2011101010033333-3311112210230100-3310002131310333-2000321030022310-2031303013103100-1311011120013331): complete subsection reference.

<a id="canonical-0222010010121030-2002121221111330-0210203213300203-0132213001033020-0012203131231310-1223010112313230-3120022201313120-1312230321223032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.direct_response_route.headers` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-006.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-006.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-009.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.port](resources--workload--reference--group-012.md#canonical-3013031222022030-2021222233122213-2220022232301003-3102220000332302-0311213311310132-3133321213021123-0133300222212330-1122131003023301)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-012.md#canonical-0101000213001300-0222020002101313-1020132032013321-2231011111202333-2122231103331210-3202123132002020-0313332003123002-3101311011132330)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes](resources--workload--reference--group-014.md#canonical-1223131333131301-1101202200212003-2303000010310012-3003131221200020-0201223130203322-2303023001222013-1231100011211030-2032330121100311)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes](resources--workload--reference--group-014.md#canonical-1320302000201200-1013121233121332-1200303031313333-2131200212100310-2031222303000133-0013120300210320-3032201103202020-0103102312330232)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.direct_response_route](resources--workload--reference--group-014.md#canonical-3323201111322002-1033013210321030-1202233310120130-0103032213302001-0022102310321331-0300303311223133-3133223031000013-0100120001003301)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.direct_response_route.headers

<a id="canonical-1221132100300322-0110112211003300-3102333010332321-2330232102331133-3131211321233130-2003302101213102-3300031330002310-3100003012220203"></a>

Type: `"object"`. list nested block, Optional.

Headers. List of (key, value) headers.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
    },
    "minItems": 0,
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
headers {
  # Configure direct properties listed below.
}
```

<a id="canonical-1222230032112231-3202010101001212-3110110022013321-0303330203132021-1120012222122013-2221211232200021-0311303300113213-1021313003231013"></a>

### Direct properties for `service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.direct_response_route.headers`

<a id="canonical-1121110301311130-0322031210102022-1120000123123323-3200302210110233-3303031201302311-1332331020331021-3300331010113010-3123230102102230"></a>

#### `service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.direct_response_route.headers.exact` property

Type: `"string"`. Optional.

Exclusive with \[presence regular expression\] Header value to match exactly.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 256
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
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
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.not_empty": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.not_empty": "true"
  }
}
```

<a id="canonical-2330101220310310-3132123211103330-2033022332201013-0013120310021223-3000131130111031-2133323311333321-0201022021331132-2211312330133232"></a>

<a id="canonical-2013211010031210-0011101210323023-3021022021310332-3030010233202320-0133210303302110-3032223231131101-1311303220330001-2120332311321111"></a>

#### `service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.direct_response_route.headers.invert_match` property

Type: `"bool"`. Optional.

Invert the result of the match to detect missing header or non-matching value.

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

<a id="canonical-1133301330221032-2022221233120110-2010002133302031-1012223210201232-1103120313132320-3111132002121021-1320313222310303-1203202220322000"></a>

<a id="canonical-3233231013322212-3130221123311003-3200022021031210-2013332202031212-3001302111332101-1231302222011223-1220031323122302-0030022223313302"></a>

#### `service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.direct_response_route.headers.name` property

Type: `"string"`. Optional.

Name. Name of the header.

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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.http_header_field": "true",
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.http_header_field": "true",
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1"
  }
}
```

<a id="canonical-2320110202212101-2000000011121130-0332111221031230-1320210233123130-3203330003221022-1011133311133302-0202111211001011-3301000222232230"></a>

<a id="canonical-2222130211030310-1020332023112223-3021212031232221-2202233101100001-2131002032211002-3101031130223032-2031030133131331-3313001130313303"></a>

#### `service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.direct_response_route.headers.presence` property

Type: `"bool"`. Optional.

Exclusive with \[exact regular expression\] If true, check for presence of header.

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

<a id="canonical-1323022201222333-0120031232122022-2100213333130311-0320133303231011-1132032333303132-0212303210200330-0221030131031110-0110320203032310"></a>

<a id="canonical-0010130202222012-1313112030131221-2022102311110323-3103300100033133-0023113032000322-1222112123023033-1212003031332113-0012310101221310"></a>

#### `service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.direct_response_route.headers.regex` property

Type: `"string"`. Optional.

Exclusive with \[exact presence\] regular expression match of the header value in re2 format.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 256
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
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
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.regex": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.regex": "true"
  }
}
```
