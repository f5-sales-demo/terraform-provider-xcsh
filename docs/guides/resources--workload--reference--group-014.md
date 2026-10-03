---
page_title: "xcsh_workload reference"
subcategory: "Container"
description: "Complete grouped canonical reference for xcsh_workload reference."
---

# xcsh_workload reference

<a id="canonical-0332211232202331-1302022303210300-0322223222012321-3213330233033021-2230212333110013-3001131201030222-3020100111330130-1333203123313003"></a>

## client_certificate_optional property — use_mtls / 322232012023 / 4

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

- [crl](resources--workload--reference--group-014.md#canonical-2210332213122022-0313313112222210-1221311123222020-2103131201012322-0032232132103120-0020300233222020-2132032313233031-1030212012211222): complete subsection reference.

- [no_crl](resources--workload--reference--group-014.md#canonical-2202333001101320-0330320302311032-0333132202300303-2221212212201311-1130133000230200-3022121231230213-2101323223321010-0233222131221003): complete subsection reference.

- [trusted_ca](resources--workload--reference--group-014.md#canonical-0223032130012113-0123222002030201-2311113110033213-1322132113003023-2022310032321111-3112011122232333-2312202332122021-1031023001232131): complete subsection reference.

<a id="canonical-0200032010330221-2331033211310121-3003023032231111-0003120003233102-0223232021221023-2003120112322303-2013222003102312-1311011010222322"></a>

<a id="canonical-3113201210233100-0312113201030311-0210212303323211-0201311000211102-3200021122333321-0210033311002311-1302302031031023-2202320330010233"></a>

## trusted_ca_url property — use_mtls / 322232012023 / 5

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

- [xfcc_disabled](resources--workload--reference--group-014.md#canonical-0222231130000212-1111110033320213-2303103222021013-1223133022333133-3231123313100031-2033220320122033-0223201110120033-2313130203333122): complete subsection reference.

- [xfcc_options](resources--workload--reference--group-014.md#canonical-3033121221321332-3202231033112132-0211011302031212-1031131103230031-0220211002000010-3103211121112012-3312230113220012-2010011332222202): complete subsection reference.

<a id="canonical-1100112013030223-3322120203103212-1213012132013103-1212123121020231-3123201202031220-2121000120033121-2303110123001132-1102211211100311"></a>

## Next pages — use_mtls / 322232012023 / 6

- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.use_mtls.crl](resources--workload--reference--group-014.md#canonical-2210332213122022-0313313112222210-1221311123222020-2103131201012322-0032232132103120-0020300233222020-2132032313233031-1030212012211222)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.use_mtls.no_crl](resources--workload--reference--group-014.md#canonical-2202333001101320-0330320302311032-0333132202300303-2221212212201311-1130133000230200-3022121231230213-2101323223321010-0233222131221003)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.use_mtls.trusted_ca](resources--workload--reference--group-014.md#canonical-0223032130012113-0123222002030201-2311113110033213-1322132113003023-2022310032321111-3112011122232333-2312202332122021-1031023001232131)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.use_mtls.xfcc_disabled](resources--workload--reference--group-014.md#canonical-0222231130000212-1111110033320213-2303103222021013-1223133022333133-3231123313100031-2033220320122033-0223201110120033-2313130203333122)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.use_mtls.xfcc_options](resources--workload--reference--group-014.md#canonical-3033121221321332-3202231033112132-0211011302031212-1031131103230031-0220211002000010-3103211121112012-3312230113220012-2010011332222202)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-013.md#canonical-1231003120133112-2013110130100211-2233121010211001-2121031322013232-2232220303133231-0300100233200310-0003012301120011-3102210111030310)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-2210332213122022-0313313112222210-1221311123222020-2103131201012322-0032232132103120-0020300233222020-2132032313233031-1030212012211222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3330301302303302-0232012013111021-3113201232320001-1022310021010233-0233203333210330-0323032300313222-0301300033231211-1313003032211312"></a>

## service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.use_mtls.crl — crl / 323320110131 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.port](resources--workload--reference--group-012.md#canonical-3013031222022030-2021222233122213-2220022232301003-3102220000332302-0311213311310132-3133321213021123-0133300222212330-1122131003023301)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-012.md#canonical-0101000213001300-0222020002101313-1020132032013321-2231011111202333-2122231103331210-3202123132002020-0313332003123002-3101311011132330)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https](resources--workload--reference--group-012.md#canonical-2221000320213131-2001211030022032-2023131131113110-1000133011113313-1003112332120000-2102122231000321-0211212202023222-1022013120231111)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-013.md#canonical-1231003120133112-2013110130100211-2233121010211001-2121031322013232-2232220303133231-0300100233200310-0003012301120011-3102210111030310)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.use_mtls](resources--workload--reference--group-013.md#canonical-2212033132123331-1022023101100202-0012200211203323-3100333332232032-2000331210002203-3023131312000213-0233333201322123-2002202002102311)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.use_mtls.crl

<a id="canonical-2120122230201011-2131101221130032-2023201200221322-0021210021021222-3013110123322123-3313323331023313-0302030232220100-0231203221100110"></a>

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

<a id="canonical-0333313322313312-2121103320033130-0320323133001322-1201020200231000-2132011322121112-0200312112230333-0222012101221231-0300013213310331"></a>

## Direct properties — crl / 323320110131 / 3

<a id="canonical-2231103112032311-2101321030220123-3112022322200210-2123302130020101-2123001021111230-3200330202033320-0013030113133112-2211121032223211"></a>

<a id="canonical-0223030301332111-2321021230102310-3232002130310330-0203210003323033-0320101202233221-2210033212000003-3001321031023201-1330122121020133"></a>

## name property — crl / 323320110131 / 4

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

<a id="canonical-0202122011332100-0130021003213123-0233213101031231-2233322210233112-3331223302010030-0213100303212131-2032212212321231-0300221323330320"></a>

<a id="canonical-0001111010312201-3100221230230100-0031121311033023-3012300121200012-3210210132023333-3112332321100100-1130210321120332-2120111002021232"></a>

## namespace property — crl / 323320110131 / 5

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

<a id="canonical-0111010320132321-3232310121021122-0302121132331020-0003113123310311-2211321001220011-1232211033202311-1120310133222131-2322102102030101"></a>

<a id="canonical-2211232021311021-3010313133222012-0232112322121212-0010131021320001-0333131123321321-3132132233103310-2321201201011120-1321311031032121"></a>

## tenant property — crl / 323320110131 / 6

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

<a id="canonical-3200303302132321-2120113222032103-3233222003002013-1010012300130033-3022103100012203-1231300121323202-1222330230013112-0021133112100220"></a>

## Next pages — crl / 323320110131 / 7

- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.use_mtls](resources--workload--reference--group-013.md#canonical-2212033132123331-1022023101100202-0012200211203323-3100333332232032-2000331210002203-3023131312000213-0233333201322123-2002202002102311)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-2202333001101320-0330320302311032-0333132202300303-2221212212201311-1130133000230200-3022121231230213-2101323223321010-0233222131221003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0103303132320021-2000331201312111-3313132113333003-3221222233233112-1130023322021001-2231111221101200-2221212030133122-3023301201000021"></a>

## service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.use_mtls.no_crl — no_crl / 212120301002 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.port](resources--workload--reference--group-012.md#canonical-3013031222022030-2021222233122213-2220022232301003-3102220000332302-0311213311310132-3133321213021123-0133300222212330-1122131003023301)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-012.md#canonical-0101000213001300-0222020002101313-1020132032013321-2231011111202333-2122231103331210-3202123132002020-0313332003123002-3101311011132330)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https](resources--workload--reference--group-012.md#canonical-2221000320213131-2001211030022032-2023131131113110-1000133011113313-1003112332120000-2102122231000321-0211212202023222-1022013120231111)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-013.md#canonical-1231003120133112-2013110130100211-2233121010211001-2121031322013232-2232220303133231-0300100233200310-0003012301120011-3102210111030310)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.use_mtls](resources--workload--reference--group-013.md#canonical-2212033132123331-1022023101100202-0012200211203323-3100333332232032-2000331210002203-3023131312000213-0233333201322123-2002202002102311)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.use_mtls.no_crl

<a id="canonical-1011102032321323-3331031331231020-1003313311112023-0211113233003002-1010123112213012-0030202010213013-3321013301122302-3032012222123000"></a>

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

<a id="canonical-3301302331122323-1322230221031103-2103030032312221-3300313023101012-1321120112011133-2201333020223222-1330312011110033-0121332000200231"></a>

## Direct properties — no_crl / 212120301002 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1311302311302111-1301202021113321-2003223120013233-1100030102000112-3021031121022100-1321322220213223-1132123220120322-2133221012333201"></a>

## Next pages — no_crl / 212120301002 / 4

- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.use_mtls](resources--workload--reference--group-013.md#canonical-2212033132123331-1022023101100202-0012200211203323-3100333332232032-2000331210002203-3023131312000213-0233333201322123-2002202002102311)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-0223032130012113-0123222002030201-2311113110033213-1322132113003023-2022310032321111-3112011122232333-2312202332122021-1031023001232131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3010010123102100-1201201103023013-2213300020102111-3312202232121030-1002120231013323-0231111030223121-3101323301123322-0001101210333110"></a>

## service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.use_mtls.trusted_ca — trusted_ca / 331311030231 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.port](resources--workload--reference--group-012.md#canonical-3013031222022030-2021222233122213-2220022232301003-3102220000332302-0311213311310132-3133321213021123-0133300222212330-1122131003023301)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-012.md#canonical-0101000213001300-0222020002101313-1020132032013321-2231011111202333-2122231103331210-3202123132002020-0313332003123002-3101311011132330)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https](resources--workload--reference--group-012.md#canonical-2221000320213131-2001211030022032-2023131131113110-1000133011113313-1003112332120000-2102122231000321-0211212202023222-1022013120231111)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-013.md#canonical-1231003120133112-2013110130100211-2233121010211001-2121031322013232-2232220303133231-0300100233200310-0003012301120011-3102210111030310)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.use_mtls](resources--workload--reference--group-013.md#canonical-2212033132123331-1022023101100202-0012200211203323-3100333332232032-2000331210002203-3023131312000213-0233333201322123-2002202002102311)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.use_mtls.trusted_ca

<a id="canonical-0211033101131210-1200021003202101-0121301320302320-1110101213011321-1030223000321001-3013211312002111-3322313301220111-2011031020331333"></a>

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

<a id="canonical-2333033021203011-3103132011113011-2103022123212203-3330310210113212-2121322330100000-2100113231121201-2302000201202310-0101120222213032"></a>

## Direct properties — trusted_ca / 331311030231 / 3

<a id="canonical-1230321322100103-1233200320032321-0203332311100032-0122210333333321-2102322332312210-0233303021121320-3220131112202300-1102031212000230"></a>

<a id="canonical-1202300211130113-2022013113121131-3000311330121013-2222011223110103-1222233301230221-2323001333212002-1301212232010320-0220310012220223"></a>

## name property — trusted_ca / 331311030231 / 4

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

<a id="canonical-1320012203000221-3030132022303321-3021011322331033-1131203011021011-3103110221001221-2213112231332020-1220220303331310-3231023321232100"></a>

<a id="canonical-0311322231203011-3213320330232221-3202203333323110-1003102200202233-2211300013223333-0132302130323223-2222030022331311-0303021321123212"></a>

## namespace property — trusted_ca / 331311030231 / 5

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

<a id="canonical-0202213302021123-1332132132120202-2213013233321031-0302032223320100-3301132012033300-2131331221330323-1302201213320233-3013030132333120"></a>

<a id="canonical-2300310022303310-2302022200322303-0030312121203102-2302310100212301-1030013122110033-1113122000012022-1201232232203130-3132223223012230"></a>

## tenant property — trusted_ca / 331311030231 / 6

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

<a id="canonical-1332123032302002-3111110222313302-1030321310000311-1020112230231312-0011332110030213-2013032133212120-0333011130200302-3220000121002013"></a>

## Next pages — trusted_ca / 331311030231 / 7

- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.use_mtls](resources--workload--reference--group-013.md#canonical-2212033132123331-1022023101100202-0012200211203323-3100333332232032-2000331210002203-3023131312000213-0233333201322123-2002202002102311)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-0222231130000212-1111110033320213-2303103222021013-1223133022333133-3231123313100031-2033220320122033-0223201110120033-2313130203333122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3130031311330012-0003232113222310-1321112021130111-2101230031011213-2213030020022232-1210133031032231-3210031013100010-2033233220300032"></a>

## service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.use_mtls.xfcc_disabled — xfcc_disabled / 030303012310 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.port](resources--workload--reference--group-012.md#canonical-3013031222022030-2021222233122213-2220022232301003-3102220000332302-0311213311310132-3133321213021123-0133300222212330-1122131003023301)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-012.md#canonical-0101000213001300-0222020002101313-1020132032013321-2231011111202333-2122231103331210-3202123132002020-0313332003123002-3101311011132330)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https](resources--workload--reference--group-012.md#canonical-2221000320213131-2001211030022032-2023131131113110-1000133011113313-1003112332120000-2102122231000321-0211212202023222-1022013120231111)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-013.md#canonical-1231003120133112-2013110130100211-2233121010211001-2121031322013232-2232220303133231-0300100233200310-0003012301120011-3102210111030310)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.use_mtls](resources--workload--reference--group-013.md#canonical-2212033132123331-1022023101100202-0012200211203323-3100333332232032-2000331210002203-3023131312000213-0233333201322123-2002202002102311)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.use_mtls.xfcc_disabled

<a id="canonical-2322232301122111-2313031321130122-2321223110002000-1003000123113332-2231122002230303-3233310022230210-0231031321010120-0322220111122322"></a>

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

<a id="canonical-3222230122211121-2320001302113213-0211100303031201-1113332323121222-3003030223001031-3131223132322111-3320332000202303-2221021331231002"></a>

## Direct properties — xfcc_disabled / 030303012310 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3203223120310123-0103113003212113-2103203323300130-3213221110013132-1301121112200012-0221322320010013-1131023222032222-1231010212101213"></a>

## Next pages — xfcc_disabled / 030303012310 / 4

- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.use_mtls](resources--workload--reference--group-013.md#canonical-2212033132123331-1022023101100202-0012200211203323-3100333332232032-2000331210002203-3023131312000213-0233333201322123-2002202002102311)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-3033121221321332-3202231033112132-0211011302031212-1031131103230031-0220211002000010-3103211121112012-3312230113220012-2010011332222202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2122200222300320-1311321221220133-1201031112310121-3023322300133230-2201323222133223-3303011323211031-2313023320003020-0310113003001111"></a>

## service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.use_mtls.xfcc_options — xfcc_options / 302131111230 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.port](resources--workload--reference--group-012.md#canonical-3013031222022030-2021222233122213-2220022232301003-3102220000332302-0311213311310132-3133321213021123-0133300222212330-1122131003023301)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-012.md#canonical-0101000213001300-0222020002101313-1020132032013321-2231011111202333-2122231103331210-3202123132002020-0313332003123002-3101311011132330)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https](resources--workload--reference--group-012.md#canonical-2221000320213131-2001211030022032-2023131131113110-1000133011113313-1003112332120000-2102122231000321-0211212202023222-1022013120231111)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-013.md#canonical-1231003120133112-2013110130100211-2233121010211001-2121031322013232-2232220303133231-0300100233200310-0003012301120011-3102210111030310)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.use_mtls](resources--workload--reference--group-013.md#canonical-2212033132123331-1022023101100202-0012200211203323-3100333332232032-2000331210002203-3023131312000213-0233333201322123-2002202002102311)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.use_mtls.xfcc_options

<a id="canonical-3200113033013223-1222011112122210-0311211133200113-1111200231332322-3232101202312212-0122102023010030-2122312220022303-3100201102012333"></a>

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

<a id="canonical-2332021202201201-2201320300222223-2203101211312110-1232330112320012-3133031331332322-1201302033222203-1233211303010233-0200010333223031"></a>

## Direct properties — xfcc_options / 302131111230 / 3

<a id="canonical-0321331230223131-3300323320211232-3003313233000310-0100232201112300-2012120103101131-3212010331003212-2333010313332003-1221230022133302"></a>

<a id="canonical-3221103200211132-2301232332333030-1003331302000313-0302332231010113-0303320030312221-0131311211113231-2030313013200101-1112122302221301"></a>

## xfcc_header_elements property — xfcc_options / 302131111230 / 4

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

<a id="canonical-2233002231033202-0203032221321233-2331021032222332-1322120210200010-3130012033311033-1100222221332230-1023221020232101-0223220101033121"></a>

## Next pages — xfcc_options / 302131111230 / 5

- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.use_mtls](resources--workload--reference--group-013.md#canonical-2212033132123331-1022023101100202-0012200211203323-3100333332232032-2000331210002203-3023131312000213-0233333201322123-2002202002102311)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-2233002113231233-1201011231003023-2310112000233003-0120011321233313-3122132210210203-2300131312231320-2323313110111333-1012021022013033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0303202122211200-2222330031310331-1202303130231132-3101330112311030-3210333002010100-2100330303302032-0111202312001120-0122110002331220"></a>

## service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert — https_auto_cert / 031130333032 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.port](resources--workload--reference--group-012.md#canonical-3013031222022030-2021222233122213-2220022232301003-3102220000332302-0311213311310132-3133321213021123-0133300222212330-1122131003023301)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-012.md#canonical-0101000213001300-0222020002101313-1020132032013321-2231011111202333-2122231103331210-3202123132002020-0313332003123002-3101311011132330)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert

<a id="canonical-2231322012213232-0101322002011231-3121130002130200-3331233223021112-0231223200303212-2302001200013232-3320333001313032-1233123330210211"></a>

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
  validators.ConflictingObjectAttributes("no_mtls",
    "use_mtls"),
  validators.ConflictingObjectAttributes("pass_through",
    "server_name"),
  validators.ConflictingObjectAttributes("port",
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

<a id="canonical-1303111101001133-0101321233000221-0021211313133223-2323110010102200-3033233102103203-1123132012212203-1231132012003110-3030112222021321"></a>

## Direct properties — https_auto_cert / 031130333032 / 3

<a id="canonical-2102103231130022-2322001103332030-2003322133100213-0132231303100223-1023202010331113-1002230210110322-3101130022201011-1020120032012033"></a>

<a id="canonical-1323332313332332-3103112032100212-0202330303032300-0311201223202212-1313102103031322-2233222020303231-0313222101033122-3021303130001010"></a>

## add_hsts property — https_auto_cert / 031130333032 / 4

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

<a id="canonical-2012013122310022-0120310112112113-1203010111322013-1122200013100213-2033230013111320-1013133131332110-0332021130311203-3222030030333000"></a>

## append_server_name property — https_auto_cert / 031130333032 / 5

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

- [coalescing_options](resources--workload--reference--group-014.md#canonical-1003313312003320-1310013331003020-2311011202322301-3113122312223102-2131103011233132-3111030113231220-3122033212120010-1013330233300210): complete subsection reference.

<a id="canonical-1331233322220122-3200130201111002-3010332231000320-2132011312222113-3122120221310232-2111122300322120-2021103220233021-0333001120300010"></a>

<a id="canonical-2321222233232330-1022312310120131-2132300320013112-0121210302333123-2221303323333333-3210302012113101-3331103122122020-1020133321212201"></a>

## connection_idle_timeout property — https_auto_cert / 031130333032 / 6

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

- [default_header](resources--workload--reference--group-014.md#canonical-2210232101332333-2320302211212012-0121133211031312-1020002112100312-0231313002212211-0130311101101123-1130132213302100-1033033032022003): complete subsection reference.

- [default_loadbalancer](resources--workload--reference--group-014.md#canonical-1013031323132002-0301102203102311-1231111002032231-1330330000313013-2022121003132201-0311130220212331-3310111013003313-2323330120111211): complete subsection reference.

- [disable_path_normalize](resources--workload--reference--group-014.md#canonical-0102331301232000-2120130002202233-3301201131102112-2211202113012002-2103232303302002-2123020300003220-0103130020112000-3322111203012123): complete subsection reference.

- [enable_path_normalize](resources--workload--reference--group-014.md#canonical-0010230120211331-3320303220201011-3112231022311202-3011321001310331-3321100211030313-2303112320201212-0110021020322200-2022222330100133): complete subsection reference.

- [http_protocol_options](resources--workload--reference--group-014.md#canonical-1001201213213330-3123312003110132-1101112033021031-2320100223111332-3200320010031233-3303131213230332-0010000311321001-0111112233232201): complete subsection reference.

<a id="canonical-1121130100212011-1333231223310311-2231031231311023-0131132112200213-1030101200220210-1222232120101003-1213223213333232-1223132323302032"></a>

<a id="canonical-0312200312030131-3211312233323012-2321033220000302-0201213303002331-0120321213223031-0231321321210203-0222032323023300-1301112203200212"></a>

## http_redirect property — https_auto_cert / 031130333032 / 7

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

- [no_mtls](resources--workload--reference--group-014.md#canonical-2101100100013330-2110303301022210-3033032330202132-3323133032033332-1231211022300223-0222213220011312-0131120300002112-3222311031011122): complete subsection reference.

- [non_default_loadbalancer](resources--workload--reference--group-014.md#canonical-0311111211233021-1220112303132010-3020331022211111-3221012030333220-2111302132100212-3232300220311233-2031302323101202-3220221331130010): complete subsection reference.

- [pass_through](resources--workload--reference--group-014.md#canonical-0330111313320232-0201321223322103-0200232011221321-3232213133200013-3010132211110221-2131111101100033-0103232200333220-1131033112223003): complete subsection reference.

<a id="canonical-3120133322303013-0012123022020201-1112201331330002-1203000213321322-2102013333201132-0001113313122001-1332003302302001-3303122232232232"></a>

<a id="canonical-0131321030103330-2222011000312123-3230332313032033-0112222120102130-0110202103310111-2101010131210232-2203311133101330-2032233333100302"></a>

## port property — https_auto_cert / 031130333032 / 8

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

<a id="canonical-3303213303331131-3203132210231023-0000200032033201-1122020212031333-3002320131333212-1012230232221010-2212121331100033-0231001332222320"></a>

<a id="canonical-3122230311222112-3213033211322310-1223013022330022-3310231332211333-0322100310232223-1023331102000302-3313001330223010-3232301110020001"></a>

## port_ranges property — https_auto_cert / 031130333032 / 9

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

<a id="canonical-3003100321121133-2331330313030003-2032130222130102-1333311023012020-2031232200021322-2110133123232310-0233023310230310-3022233113331232"></a>

<a id="canonical-3313221111021102-1023033202002010-0202311112021222-2131323121200311-2223011302131132-2031213121320302-3002231003211121-3030111020002101"></a>

## server_name property — https_auto_cert / 031130333032 / 10

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

- [tls_config](resources--workload--reference--group-014.md#canonical-3332002222313202-3123321100213021-1310123312032013-0101102313123323-2003030031022013-2212223201221323-0203203103020311-2032023003330023): complete subsection reference.

- [use_mtls](resources--workload--reference--group-014.md#canonical-1101302220013000-1130323120002221-2213110022303302-1002100102000211-3203030111202113-1021233103221020-3020210323320220-3011000311311022): complete subsection reference.

<a id="canonical-1231121300312232-0231320130003112-1331221012023020-0132232002032231-1123112300013021-3213111101131012-1202332020212302-3332312211133320"></a>

## Next pages — https_auto_cert / 031130333032 / 11

- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.coalescing_options](resources--workload--reference--group-014.md#canonical-1003313312003320-1310013331003020-2311011202322301-3113122312223102-2131103011233132-3111030113231220-3122033212120010-1013330233300210)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.default_header](resources--workload--reference--group-014.md#canonical-2210232101332333-2320302211212012-0121133211031312-1020002112100312-0231313002212211-0130311101101123-1130132213302100-1033033032022003)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.default_loadbalancer](resources--workload--reference--group-014.md#canonical-1013031323132002-0301102203102311-1231111002032231-1330330000313013-2022121003132201-0311130220212331-3310111013003313-2323330120111211)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.disable_path_normalize](resources--workload--reference--group-014.md#canonical-0102331301232000-2120130002202233-3301201131102112-2211202113012002-2103232303302002-2123020300003220-0103130020112000-3322111203012123)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.enable_path_normalize](resources--workload--reference--group-014.md#canonical-0010230120211331-3320303220201011-3112231022311202-3011321001310331-3321100211030313-2303112320201212-0110021020322200-2022222330100133)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.http_protocol_options](resources--workload--reference--group-014.md#canonical-1001201213213330-3123312003110132-1101112033021031-2320100223111332-3200320010031233-3303131213230332-0010000311321001-0111112233232201)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.no_mtls](resources--workload--reference--group-014.md#canonical-2101100100013330-2110303301022210-3033032330202132-3323133032033332-1231211022300223-0222213220011312-0131120300002112-3222311031011122)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.non_default_loadbalancer](resources--workload--reference--group-014.md#canonical-0311111211233021-1220112303132010-3020331022211111-3221012030333220-2111302132100212-3232300220311233-2031302323101202-3220221331130010)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.pass_through](resources--workload--reference--group-014.md#canonical-0330111313320232-0201321223322103-0200232011221321-3232213133200013-3010132211110221-2131111101100033-0103232200333220-1131033112223003)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.tls_config](resources--workload--reference--group-014.md#canonical-3332002222313202-3123321100213021-1310123312032013-0101102313123323-2003030031022013-2212223201221323-0203203103020311-2032023003330023)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.use_mtls](resources--workload--reference--group-014.md#canonical-1101302220013000-1130323120002221-2213110022303302-1002100102000211-3203030111202113-1021233103221020-3020210323320220-3011000311311022)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-012.md#canonical-0101000213001300-0222020002101313-1020132032013321-2231011111202333-2122231103331210-3202123132002020-0313332003123002-3101311011132330)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-1003313312003320-1310013331003020-2311011202322301-3113122312223102-2131103011233132-3111030113231220-3122033212120010-1013330233300210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1312013211203021-0112122302031120-3320010201203100-3100100211122001-1121312130321001-2221110313100010-2110002203301020-2333212123000312"></a>

## service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.coalescing_options — coalescing_options / 023010303300 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.port](resources--workload--reference--group-012.md#canonical-3013031222022030-2021222233122213-2220022232301003-3102220000332302-0311213311310132-3133321213021123-0133300222212330-1122131003023301)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-012.md#canonical-0101000213001300-0222020002101313-1020132032013321-2231011111202333-2122231103331210-3202123132002020-0313332003123002-3101311011132330)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert](resources--workload--reference--group-014.md#canonical-2233002113231233-1201011231003023-2310112000233003-0120011321233313-3122132210210203-2300131312231320-2323313110111333-1012021022013033)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.coalescing_options

<a id="canonical-2132121232223132-2331212203333322-3122201000100112-1201322331213032-1322001131231033-0122113320223031-0312301212001232-2313300032133130"></a>

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

<a id="canonical-0203131201201301-1121003020021211-0313131102312213-1320203111211213-1212132300112013-2010131021001312-1302013310101310-1310210002131013"></a>

## Direct properties — coalescing_options / 023010303300 / 3

- [default_coalescing](resources--workload--reference--group-014.md#canonical-1002111202130012-2213212112200030-2231331233321001-3123133202012110-2111301012221231-0320131333132230-0113310203203222-2311213030231001): complete subsection reference.

- [strict_coalescing](resources--workload--reference--group-014.md#canonical-1211101112303033-2232132010132032-0013102332121220-2213112003331131-2233202030022220-1102233032130032-2210312030103100-3232211103332011): complete subsection reference.

<a id="canonical-1031300302330331-0033112113323033-3312313011220002-3133001203230202-2332330321212111-2102310311003110-1323300120231133-3130323012322302"></a>

## Next pages — coalescing_options / 023010303300 / 4

- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.coalescing_options.default_coalescing](resources--workload--reference--group-014.md#canonical-1002111202130012-2213212112200030-2231331233321001-3123133202012110-2111301012221231-0320131333132230-0113310203203222-2311213030231001)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.coalescing_options.strict_coalescing](resources--workload--reference--group-014.md#canonical-1211101112303033-2232132010132032-0013102332121220-2213112003331131-2233202030022220-1102233032130032-2210312030103100-3232211103332011)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert](resources--workload--reference--group-014.md#canonical-2233002113231233-1201011231003023-2310112000233003-0120011321233313-3122132210210203-2300131312231320-2323313110111333-1012021022013033)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-1002111202130012-2213212112200030-2231331233321001-3123133202012110-2111301012221231-0320131333132230-0113310203203222-2311213030231001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2211320200133233-1311333001113011-2312320130020300-0302120011101301-0210331032301020-3211011031321101-0322313022020113-0200112000120133"></a>

## service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.coalescing_options.default_coalescing — default_coalescing / 002300110330 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.port](resources--workload--reference--group-012.md#canonical-3013031222022030-2021222233122213-2220022232301003-3102220000332302-0311213311310132-3133321213021123-0133300222212330-1122131003023301)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-012.md#canonical-0101000213001300-0222020002101313-1020132032013321-2231011111202333-2122231103331210-3202123132002020-0313332003123002-3101311011132330)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert](resources--workload--reference--group-014.md#canonical-2233002113231233-1201011231003023-2310112000233003-0120011321233313-3122132210210203-2300131312231320-2323313110111333-1012021022013033)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.coalescing_options](resources--workload--reference--group-014.md#canonical-1003313312003320-1310013331003020-2311011202322301-3113122312223102-2131103011233132-3111030113231220-3122033212120010-1013330233300210)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.coalescing_options.default_coalescing

<a id="canonical-3303213230301231-1203323003222010-3121120112110313-0012300131203201-2201210100022110-2232233313220301-1123030122323113-0122211132120030"></a>

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

<a id="canonical-2202322321203113-2103322013223332-3201112032110120-1333112110130303-0231302231212120-3132322300230321-1301310211022023-3031101301131223"></a>

## Direct properties — default_coalescing / 002300110330 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1202100012333030-3121210121101123-0302230001130230-0312132322322001-1321121230213011-3031131121310201-3333303100122123-0330110321002003"></a>

## Next pages — default_coalescing / 002300110330 / 4

- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.coalescing_options](resources--workload--reference--group-014.md#canonical-1003313312003320-1310013331003020-2311011202322301-3113122312223102-2131103011233132-3111030113231220-3122033212120010-1013330233300210)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-1211101112303033-2232132010132032-0013102332121220-2213112003331131-2233202030022220-1102233032130032-2210312030103100-3232211103332011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1203130001032323-1103200030320301-0223100331312133-1002133203131331-2303310121120323-1122313320012130-1021231001301320-0123110100130100"></a>

## service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.coalescing_options.strict_coalescing — strict_coalescing / 113021301111 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.port](resources--workload--reference--group-012.md#canonical-3013031222022030-2021222233122213-2220022232301003-3102220000332302-0311213311310132-3133321213021123-0133300222212330-1122131003023301)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-012.md#canonical-0101000213001300-0222020002101313-1020132032013321-2231011111202333-2122231103331210-3202123132002020-0313332003123002-3101311011132330)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert](resources--workload--reference--group-014.md#canonical-2233002113231233-1201011231003023-2310112000233003-0120011321233313-3122132210210203-2300131312231320-2323313110111333-1012021022013033)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.coalescing_options](resources--workload--reference--group-014.md#canonical-1003313312003320-1310013331003020-2311011202322301-3113122312223102-2131103011233132-3111030113231220-3122033212120010-1013330233300210)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.coalescing_options.strict_coalescing

<a id="canonical-3020110320132020-3132013233302021-3013211103311022-3012120012223312-0003001321332231-0130132233210203-0002032002133010-0133103013001333"></a>

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

<a id="canonical-3323310122100211-2222002213202311-3122101012102001-2123111313001321-2103032131232022-1110110202331100-0120120212322031-3213230311303111"></a>

## Direct properties — strict_coalescing / 113021301111 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2110331103302121-3313002032120231-0310323023013230-3231021203102022-3133032010222303-3003133332022300-3013031200323031-2001203003210012"></a>

## Next pages — strict_coalescing / 113021301111 / 4

- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.coalescing_options](resources--workload--reference--group-014.md#canonical-1003313312003320-1310013331003020-2311011202322301-3113122312223102-2131103011233132-3111030113231220-3122033212120010-1013330233300210)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-2210232101332333-2320302211212012-0121133211031312-1020002112100312-0231313002212211-0130311101101123-1130132213302100-1033033032022003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2102233031121230-2012230302210020-0021212213212111-0101132113213031-3102130201103300-3200201210011011-2111202033010312-1130200101221103"></a>

## service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.default_header — default_header / 233132201321 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.port](resources--workload--reference--group-012.md#canonical-3013031222022030-2021222233122213-2220022232301003-3102220000332302-0311213311310132-3133321213021123-0133300222212330-1122131003023301)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-012.md#canonical-0101000213001300-0222020002101313-1020132032013321-2231011111202333-2122231103331210-3202123132002020-0313332003123002-3101311011132330)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert](resources--workload--reference--group-014.md#canonical-2233002113231233-1201011231003023-2310112000233003-0120011321233313-3122132210210203-2300131312231320-2323313110111333-1012021022013033)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.default_header

<a id="canonical-0113000302100031-0012033011223031-3003332313232230-1211201202231220-3010013310113230-2311230010332312-1312102011211110-0322100020231222"></a>

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

<a id="canonical-1221232323021332-0032201002000130-3232220111133112-1331320320301300-2310312111031000-2131022033303111-2032120010322333-3232032113110203"></a>

## Direct properties — default_header / 233132201321 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0112021323112311-0213201102331031-0222110133020021-1321312200020233-1022100012210301-3333103322001111-2111313333212002-0302211312300313"></a>

## Next pages — default_header / 233132201321 / 4

- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert](resources--workload--reference--group-014.md#canonical-2233002113231233-1201011231003023-2310112000233003-0120011321233313-3122132210210203-2300131312231320-2323313110111333-1012021022013033)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-1013031323132002-0301102203102311-1231111002032231-1330330000313013-2022121003132201-0311130220212331-3310111013003313-2323330120111211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2122023211103031-2122210330022333-1011112231110002-0303011330200022-3232211123131112-0103012233301030-2012122011231020-1233012033113001"></a>

## service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.default_loadbalancer — default_loadbalancer / 000221112011 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.port](resources--workload--reference--group-012.md#canonical-3013031222022030-2021222233122213-2220022232301003-3102220000332302-0311213311310132-3133321213021123-0133300222212330-1122131003023301)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-012.md#canonical-0101000213001300-0222020002101313-1020132032013321-2231011111202333-2122231103331210-3202123132002020-0313332003123002-3101311011132330)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert](resources--workload--reference--group-014.md#canonical-2233002113231233-1201011231003023-2310112000233003-0120011321233313-3122132210210203-2300131312231320-2323313110111333-1012021022013033)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.default_loadbalancer

<a id="canonical-2103013032030320-1001202221111330-1310133132121013-0311032113112212-1002223201130032-1001231320231233-3332023202333313-3131120012310011"></a>

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

<a id="canonical-1030023022333212-3101122020310221-2233220032032230-3300212003012312-2230213103212022-0311020220101000-3113211300211003-3102100102201202"></a>

## Direct properties — default_loadbalancer / 000221112011 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0123033330021013-0120301222121311-2113012001212120-0223030213303203-3112321311212210-2213332001032330-0233303313132013-1312101222102031"></a>

## Next pages — default_loadbalancer / 000221112011 / 4

- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert](resources--workload--reference--group-014.md#canonical-2233002113231233-1201011231003023-2310112000233003-0120011321233313-3122132210210203-2300131312231320-2323313110111333-1012021022013033)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-0102331301232000-2120130002202233-3301201131102112-2211202113012002-2103232303302002-2123020300003220-0103130020112000-3322111203012123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2013010012021100-2213023321302033-0322211120201113-2020330233303323-1133023220330001-1102131013010223-0210223113013321-3211210110121133"></a>

## service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.disable_path_normalize — disable_path_normalize / 011221031320 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.port](resources--workload--reference--group-012.md#canonical-3013031222022030-2021222233122213-2220022232301003-3102220000332302-0311213311310132-3133321213021123-0133300222212330-1122131003023301)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-012.md#canonical-0101000213001300-0222020002101313-1020132032013321-2231011111202333-2122231103331210-3202123132002020-0313332003123002-3101311011132330)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert](resources--workload--reference--group-014.md#canonical-2233002113231233-1201011231003023-2310112000233003-0120011321233313-3122132210210203-2300131312231320-2323313110111333-1012021022013033)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.disable_path_normalize

<a id="canonical-1320022301001222-1331201122012320-1220013202011023-3130333132331202-1111331132130002-0330303020200220-3301121212133210-0200223310303030"></a>

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

<a id="canonical-3332111331023323-1000103321000203-1202332323330231-2022012321110023-1211231031313012-0302230231022121-2302221030211323-1221312102002020"></a>

## Direct properties — disable_path_normalize / 011221031320 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3233121231000210-1330311220101222-3330223102311122-0113203000113320-2023132330001312-2112023301203310-2132210320121013-3000310110331221"></a>

## Next pages — disable_path_normalize / 011221031320 / 4

- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert](resources--workload--reference--group-014.md#canonical-2233002113231233-1201011231003023-2310112000233003-0120011321233313-3122132210210203-2300131312231320-2323313110111333-1012021022013033)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-0010230120211331-3320303220201011-3112231022311202-3011321001310331-3321100211030313-2303112320201212-0110021020322200-2022222330100133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3221132123233122-3233013012200023-2222011021121013-1010122233000302-3332132020302221-2011110300003322-1223303301232302-3330113131303222"></a>

## service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.enable_path_normalize — enable_path_normalize / 303300313012 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.port](resources--workload--reference--group-012.md#canonical-3013031222022030-2021222233122213-2220022232301003-3102220000332302-0311213311310132-3133321213021123-0133300222212330-1122131003023301)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-012.md#canonical-0101000213001300-0222020002101313-1020132032013321-2231011111202333-2122231103331210-3202123132002020-0313332003123002-3101311011132330)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert](resources--workload--reference--group-014.md#canonical-2233002113231233-1201011231003023-2310112000233003-0120011321233313-3122132210210203-2300131312231320-2323313110111333-1012021022013033)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.enable_path_normalize

<a id="canonical-2323001212202300-1010122003212013-1123203112012013-1220012331331200-0221102331012033-0101230200313103-0122200023113331-1031012300121112"></a>

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

<a id="canonical-0032303131331211-1022113321111013-2111230202323212-2111202223300132-0120313320012111-1111033101132310-2120030131121003-0011320103011000"></a>

## Direct properties — enable_path_normalize / 303300313012 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1320002011312323-3102233332232132-3233101231331221-3113230032332210-3101010231002100-2020031000110312-2022320212100201-3123231220102113"></a>

## Next pages — enable_path_normalize / 303300313012 / 4

- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert](resources--workload--reference--group-014.md#canonical-2233002113231233-1201011231003023-2310112000233003-0120011321233313-3122132210210203-2300131312231320-2323313110111333-1012021022013033)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-1001201213213330-3123312003110132-1101112033021031-2320100223111332-3200320010031233-3303131213230332-0010000311321001-0111112233232201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0030313232112223-1133322030103032-3130210033231133-2310120023131201-0002000320023302-2100013211101020-2212302330201300-3303003213320302"></a>

## service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.http_protocol_options — http_protocol_options / 111202302321 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.port](resources--workload--reference--group-012.md#canonical-3013031222022030-2021222233122213-2220022232301003-3102220000332302-0311213311310132-3133321213021123-0133300222212330-1122131003023301)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-012.md#canonical-0101000213001300-0222020002101313-1020132032013321-2231011111202333-2122231103331210-3202123132002020-0313332003123002-3101311011132330)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert](resources--workload--reference--group-014.md#canonical-2233002113231233-1201011231003023-2310112000233003-0120011321233313-3122132210210203-2300131312231320-2323313110111333-1012021022013033)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.http_protocol_options

<a id="canonical-1203032030111022-3303200022323320-0130230221323010-2300100333012232-0323122231303330-2300122000113111-2321003310303112-1222333121023011"></a>

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

<a id="canonical-2100231013022323-2123113222133101-0300021130100310-3110020222213231-2202203000333310-2111332010322013-2132210110232012-3001022213223130"></a>

## Direct properties — http_protocol_options / 111202302321 / 3

- [http_protocol_enable_v1_only](resources--workload--reference--group-014.md#canonical-2101100120112220-3011131101223030-0201203113223302-2210200120000233-3120203233011300-3231001202310203-3003232210201001-1103101233013201): complete subsection reference.

- [http_protocol_enable_v1_v2](resources--workload--reference--group-014.md#canonical-3300220010231100-0213320032310313-1312000010103323-3231331211200020-1101311002031212-0233311330233132-1001213333023313-1021321003320213): complete subsection reference.

- [http_protocol_enable_v2_only](resources--workload--reference--group-014.md#canonical-0100221013330112-0132313101223121-0222202123322200-0330111332333000-0003030211321311-0003102020222111-0311110033003233-3122003023021003): complete subsection reference.

<a id="canonical-0012022132103031-2021323313303032-2020313221332203-1303103123222102-1111132131301211-0232111202201110-2213021122121123-3130322311222100"></a>

## Next pages — http_protocol_options / 111202302321 / 4

- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only](resources--workload--reference--group-014.md#canonical-2101100120112220-3011131101223030-0201203113223302-2210200120000233-3120203233011300-3231001202310203-3003232210201001-1103101233013201)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_v2](resources--workload--reference--group-014.md#canonical-3300220010231100-0213320032310313-1312000010103323-3231331211200020-1101311002031212-0233311330233132-1001213333023313-1021321003320213)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v2_only](resources--workload--reference--group-014.md#canonical-0100221013330112-0132313101223121-0222202123322200-0330111332333000-0003030211321311-0003102020222111-0311110033003233-3122003023021003)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert](resources--workload--reference--group-014.md#canonical-2233002113231233-1201011231003023-2310112000233003-0120011321233313-3122132210210203-2300131312231320-2323313110111333-1012021022013033)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-2101100120112220-3011131101223030-0201203113223302-2210200120000233-3120203233011300-3231001202310203-3003232210201001-1103101233013201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1110321312003300-0030303023102002-2031313003302312-3332231301030131-0123222230010230-3110301203222331-3311332232032002-2123302330012303"></a>

## service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only — http_protocol_enable_v1_only / 231132201103 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
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

<a id="canonical-1001331203123312-2301213033203020-1303130010300030-3202103333011320-2111030311200203-0201320311103102-2323203033213121-3313011100110322"></a>

## Direct properties — http_protocol_enable_v1_only / 231132201103 / 3

- [header_transformation](resources--workload--reference--group-014.md#canonical-1122230222212111-3313211200333020-1030101021020320-1323301220200121-1223222331313330-3120223220212022-1213020202330022-1320022103011012): complete subsection reference.

<a id="canonical-2033112203020211-3303320313223221-2000103221113232-1002132221003021-2223020210012001-2310113020103202-0033322231212012-0032223203221113"></a>

## Next pages — http_protocol_enable_v1_only / 231132201103 / 4

- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation](resources--workload--reference--group-014.md#canonical-1122230222212111-3313211200333020-1030101021020320-1323301220200121-1223222331313330-3120223220212022-1213020202330022-1320022103011012)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.http_protocol_options](resources--workload--reference--group-014.md#canonical-1001201213213330-3123312003110132-1101112033021031-2320100223111332-3200320010031233-3303131213230332-0010000311321001-0111112233232201)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-1122230222212111-3313211200333020-1030101021020320-1323301220200121-1223222331313330-3120223220212022-1213020202330022-1320022103011012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3202123302302032-1232021311232102-1002100010020112-2120213011021003-2111130330121210-0300211312210000-3113111321031312-0223113123021111"></a>

## service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation — header_transformation / 200120030330 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.port](resources--workload--reference--group-012.md#canonical-3013031222022030-2021222233122213-2220022232301003-3102220000332302-0311213311310132-3133321213021123-0133300222212330-1122131003023301)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-012.md#canonical-0101000213001300-0222020002101313-1020132032013321-2231011111202333-2122231103331210-3202123132002020-0313332003123002-3101311011132330)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert](resources--workload--reference--group-014.md#canonical-2233002113231233-1201011231003023-2310112000233003-0120011321233313-3122132210210203-2300131312231320-2323313110111333-1012021022013033)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.http_protocol_options](resources--workload--reference--group-014.md#canonical-1001201213213330-3123312003110132-1101112033021031-2320100223111332-3200320010031233-3303131213230332-0010000311321001-0111112233232201)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only](resources--workload--reference--group-014.md#canonical-2101100120112220-3011131101223030-0201203113223302-2210200120000233-3120203233011300-3231001202310203-3003232210201001-1103101233013201)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation

<a id="canonical-2020220033023231-1330130031300202-1233011111323012-0022331121100001-0012333103322003-3102222202302013-3003300211120213-3111210312221013"></a>

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

<a id="canonical-1322303103330313-1230200232302201-2300102100121233-3030120130011100-2000322332303000-1113230002120022-0131231030313301-3303103022031230"></a>

## Direct properties — header_transformation / 200120030330 / 3

- [default_header_transformation](resources--workload--reference--group-014.md#canonical-2123222033112002-0221011213330211-1210102322220233-0233131030212333-1303300203220230-3012213113312211-1303121133030122-3112022103110210): complete subsection reference.

- [preserve_case_header_transformation](resources--workload--reference--group-014.md#canonical-2202321301212202-0233011133031113-0210330212201231-0002133033002100-3313331203010213-1102020022333233-0220222222222020-1202102303331020): complete subsection reference.

- [proper_case_header_transformation](resources--workload--reference--group-014.md#canonical-1020132113133332-0001332222021231-2200322331103100-0101230301313111-2011121330231311-3313310331210031-3200223013001002-2213300000200111): complete subsection reference.

<a id="canonical-3223023030132223-3110101211012223-0330112021332003-2312122101310203-1312011123300123-0022303010230231-0133100110302000-2310010230312102"></a>

## Next pages — header_transformation / 200120030330 / 4

- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.default_header_transformation](resources--workload--reference--group-014.md#canonical-2123222033112002-0221011213330211-1210102322220233-0233131030212333-1303300203220230-3012213113312211-1303121133030122-3112022103110210)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.preserve_case_header_transformation](resources--workload--reference--group-014.md#canonical-2202321301212202-0233011133031113-0210330212201231-0002133033002100-3313331203010213-1102020022333233-0220222222222020-1202102303331020)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.proper_case_header_transformation](resources--workload--reference--group-014.md#canonical-1020132113133332-0001332222021231-2200322331103100-0101230301313111-2011121330231311-3313310331210031-3200223013001002-2213300000200111)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only](resources--workload--reference--group-014.md#canonical-2101100120112220-3011131101223030-0201203113223302-2210200120000233-3120203233011300-3231001202310203-3003232210201001-1103101233013201)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-2123222033112002-0221011213330211-1210102322220233-0233131030212333-1303300203220230-3012213113312211-1303121133030122-3112022103110210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1102221333322111-2320003213221033-0000323121010223-3313012012201011-3122203320313220-3223333013001111-1233101023102302-0211023310222003"></a>

## service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.default_header_transformation — default_header_transformation / 332120232313 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
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

<a id="canonical-3213220321300303-0300322312133222-2213103013103011-3203003210020002-1012112201233221-0302113212232310-2102003330001031-0201022230102221"></a>

## Direct properties — default_header_transformation / 332120232313 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0101202133013002-2333330031130213-3113003100222233-2113231000020320-3203331313220210-1133323223031001-2123312221210302-0111300231111003"></a>

## Next pages — default_header_transformation / 332120232313 / 4

- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation](resources--workload--reference--group-014.md#canonical-1122230222212111-3313211200333020-1030101021020320-1323301220200121-1223222331313330-3120223220212022-1213020202330022-1320022103011012)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-2202321301212202-0233011133031113-0210330212201231-0002133033002100-3313331203010213-1102020022333233-0220222222222020-1202102303331020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2122012202213212-3311020100031123-3121102010131210-2332230231131121-1302013313210321-1321101211230020-1323020210331012-0013102201003122"></a>

## service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.preserve_case_header_transformation — preserve_case_header_transformation / 130033031112 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
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

<a id="canonical-0320012113003103-0312321200212032-3010312000331330-0011000313230101-2012020203221220-3121200200032303-2322010331121231-1102322221001332"></a>

## Direct properties — preserve_case_header_transformation / 130033031112 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1333011303203133-0323303131120213-3022100010131210-0201212332233113-0112223032230231-3120110023030323-0010223323103212-1200023132223022"></a>

## Next pages — preserve_case_header_transformation / 130033031112 / 4

- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation](resources--workload--reference--group-014.md#canonical-1122230222212111-3313211200333020-1030101021020320-1323301220200121-1223222331313330-3120223220212022-1213020202330022-1320022103011012)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-1020132113133332-0001332222021231-2200322331103100-0101230301313111-2011121330231311-3313310331210031-3200223013001002-2213300000200111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0203032230131220-2200002121123201-1110103313032013-0323132133303322-2033030022333012-1210223211231003-1300123122333213-2231131003023010"></a>

## service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.proper_case_header_transformation — proper_case_header_transformation / 023130012202 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
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

<a id="canonical-0232201301023031-0021211111323032-0323033011322111-3130130021302230-1300121221200123-1131213112222311-1310003133001120-3010101031123303"></a>

## Direct properties — proper_case_header_transformation / 023130012202 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2011101121002033-2312133221031122-1111300203200111-0202101311232303-1300023201300123-2312012203300121-3011110121112312-1003102322211003"></a>

## Next pages — proper_case_header_transformation / 023130012202 / 4

- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation](resources--workload--reference--group-014.md#canonical-1122230222212111-3313211200333020-1030101021020320-1323301220200121-1223222331313330-3120223220212022-1213020202330022-1320022103011012)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-3300220010231100-0213320032310313-1312000010103323-3231331211200020-1101311002031212-0233311330233132-1001213333023313-1021321003320213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2110303002201032-0302213113311031-3022130302011110-3010030122120121-0213332212212313-2023333300001013-2010211310132113-1131230313110310"></a>

## service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_v2 — http_protocol_enable_v1_v2 / 032013113323 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.port](resources--workload--reference--group-012.md#canonical-3013031222022030-2021222233122213-2220022232301003-3102220000332302-0311213311310132-3133321213021123-0133300222212330-1122131003023301)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-012.md#canonical-0101000213001300-0222020002101313-1020132032013321-2231011111202333-2122231103331210-3202123132002020-0313332003123002-3101311011132330)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert](resources--workload--reference--group-014.md#canonical-2233002113231233-1201011231003023-2310112000233003-0120011321233313-3122132210210203-2300131312231320-2323313110111333-1012021022013033)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.http_protocol_options](resources--workload--reference--group-014.md#canonical-1001201213213330-3123312003110132-1101112033021031-2320100223111332-3200320010031233-3303131213230332-0010000311321001-0111112233232201)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_v2

<a id="canonical-1202023232133131-2101312001333320-1223211032031121-0333113301213303-3332021300300111-3230220200220103-0321011011200233-0131013001032301"></a>

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

<a id="canonical-3021113121301120-3020033211313110-0102303112102123-3022020130301002-2001321221311221-1003300212223023-0132022010102131-1220331000220203"></a>

## Direct properties — http_protocol_enable_v1_v2 / 032013113323 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1222031130122321-0001123310032020-1320002033101333-3222230303102200-1203220322220312-2233022011013233-1033200313210200-2032213000200102"></a>

## Next pages — http_protocol_enable_v1_v2 / 032013113323 / 4

- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.http_protocol_options](resources--workload--reference--group-014.md#canonical-1001201213213330-3123312003110132-1101112033021031-2320100223111332-3200320010031233-3303131213230332-0010000311321001-0111112233232201)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-0100221013330112-0132313101223121-0222202123322200-0330111332333000-0003030211321311-0003102020222111-0311110033003233-3122003023021003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2020302110300233-2023023311312130-3212032212302113-3230332123003031-0222101030121111-0302021133002223-3030311213201332-3131233333012010"></a>

## service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v2_only — http_protocol_enable_v2_only / 013012021212 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.port](resources--workload--reference--group-012.md#canonical-3013031222022030-2021222233122213-2220022232301003-3102220000332302-0311213311310132-3133321213021123-0133300222212330-1122131003023301)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-012.md#canonical-0101000213001300-0222020002101313-1020132032013321-2231011111202333-2122231103331210-3202123132002020-0313332003123002-3101311011132330)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert](resources--workload--reference--group-014.md#canonical-2233002113231233-1201011231003023-2310112000233003-0120011321233313-3122132210210203-2300131312231320-2323313110111333-1012021022013033)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.http_protocol_options](resources--workload--reference--group-014.md#canonical-1001201213213330-3123312003110132-1101112033021031-2320100223111332-3200320010031233-3303131213230332-0010000311321001-0111112233232201)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v2_only

<a id="canonical-1102110020203231-3203031030010023-2121120030211132-2120200210013132-0333111001302111-2100232313202131-0021021100321233-0002001331331022"></a>

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

<a id="canonical-1022200203301121-1333121331311203-0213131020302210-0011223133010222-1312101122220310-2020031223021022-1000230220130312-3121310012100303"></a>

## Direct properties — http_protocol_enable_v2_only / 013012021212 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0010130311120311-1202203202212131-0231333030131122-2123232211213112-3223113103030000-1220113112331130-1121210031032012-2023220101130301"></a>

## Next pages — http_protocol_enable_v2_only / 013012021212 / 4

- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.http_protocol_options](resources--workload--reference--group-014.md#canonical-1001201213213330-3123312003110132-1101112033021031-2320100223111332-3200320010031233-3303131213230332-0010000311321001-0111112233232201)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-2101100100013330-2110303301022210-3033032330202132-3323133032033332-1231211022300223-0222213220011312-0131120300002112-3222311031011122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2011333333022102-3213002012203312-0003323211332210-3222020313301303-1232100112003130-3320132323220203-2233221113011130-0211011121010112"></a>

## service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.no_mtls — no_mtls / 223010223222 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.port](resources--workload--reference--group-012.md#canonical-3013031222022030-2021222233122213-2220022232301003-3102220000332302-0311213311310132-3133321213021123-0133300222212330-1122131003023301)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-012.md#canonical-0101000213001300-0222020002101313-1020132032013321-2231011111202333-2122231103331210-3202123132002020-0313332003123002-3101311011132330)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert](resources--workload--reference--group-014.md#canonical-2233002113231233-1201011231003023-2310112000233003-0120011321233313-3122132210210203-2300131312231320-2323313110111333-1012021022013033)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.no_mtls

<a id="canonical-0212011202210030-3121231021310322-0321212311213022-2230202221321210-3200033320131312-0221012323303030-3122121233301110-0322032123312103"></a>

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

<a id="canonical-0131202210303011-0031333212222311-0322020313333323-0200211123032333-3211232120013201-0302133200031232-1031033033020101-1312132310131311"></a>

## Direct properties — no_mtls / 223010223222 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2222210121321333-0010313123203330-3232013323233101-0131002132301210-0111100111220322-0002002121111122-2230101302101333-3112120232322222"></a>

## Next pages — no_mtls / 223010223222 / 4

- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert](resources--workload--reference--group-014.md#canonical-2233002113231233-1201011231003023-2310112000233003-0120011321233313-3122132210210203-2300131312231320-2323313110111333-1012021022013033)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-0311111211233021-1220112303132010-3020331022211111-3221012030333220-2111302132100212-3232300220311233-2031302323101202-3220221331130010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2013321322011331-0212112030102303-1111303221023222-1232013121111113-3020121002323200-2112302230230322-2313310023103211-1330103100102301"></a>

## service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.non_default_loadbalancer — non_default_loadbalancer / 220113102310 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.port](resources--workload--reference--group-012.md#canonical-3013031222022030-2021222233122213-2220022232301003-3102220000332302-0311213311310132-3133321213021123-0133300222212330-1122131003023301)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-012.md#canonical-0101000213001300-0222020002101313-1020132032013321-2231011111202333-2122231103331210-3202123132002020-0313332003123002-3101311011132330)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert](resources--workload--reference--group-014.md#canonical-2233002113231233-1201011231003023-2310112000233003-0120011321233313-3122132210210203-2300131312231320-2323313110111333-1012021022013033)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.non_default_loadbalancer

<a id="canonical-0213123000313321-2233023110111210-0322201323301222-0312201123002201-2323113121322010-1212313211031002-0123132102222223-2013013021122033"></a>

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

<a id="canonical-1032111101311210-1013112233213111-3320230303032333-2320210311032102-2200333230123233-2112033013133012-1122110023302033-1010302311023332"></a>

## Direct properties — non_default_loadbalancer / 220113102310 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2321312212210310-3030110203332002-1222033112322330-2112331113103131-0301020321310111-0331231211033003-2230300213313001-1101101130000023"></a>

## Next pages — non_default_loadbalancer / 220113102310 / 4

- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert](resources--workload--reference--group-014.md#canonical-2233002113231233-1201011231003023-2310112000233003-0120011321233313-3122132210210203-2300131312231320-2323313110111333-1012021022013033)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-0330111313320232-0201321223322103-0200232011221321-3232213133200013-3010132211110221-2131111101100033-0103232200333220-1131033112223003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1010322012302002-0021032222222002-1332220111312011-1322222313202120-0031332213333023-2023010110332031-3010021010020020-0232302220231001"></a>

## service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.pass_through — pass_through / 203111112021 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.port](resources--workload--reference--group-012.md#canonical-3013031222022030-2021222233122213-2220022232301003-3102220000332302-0311213311310132-3133321213021123-0133300222212330-1122131003023301)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-012.md#canonical-0101000213001300-0222020002101313-1020132032013321-2231011111202333-2122231103331210-3202123132002020-0313332003123002-3101311011132330)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert](resources--workload--reference--group-014.md#canonical-2233002113231233-1201011231003023-2310112000233003-0120011321233313-3122132210210203-2300131312231320-2323313110111333-1012021022013033)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.pass_through

<a id="canonical-0213303320020330-0131333102301102-1333301013232233-3011322323231300-3032220100231111-3232330202233032-1121111132220000-2131121020303302"></a>

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

<a id="canonical-0221323323032313-0000333223011331-3112322313200003-2001223023313130-2030112010101023-1001123301331103-1101231030131131-3020113122012120"></a>

## Direct properties — pass_through / 203111112021 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1102202223020330-3320113001021312-3321203203300222-3022022023230312-3123321130100303-2213310232210133-2201113212212200-0102121132012130"></a>

## Next pages — pass_through / 203111112021 / 4

- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert](resources--workload--reference--group-014.md#canonical-2233002113231233-1201011231003023-2310112000233003-0120011321233313-3122132210210203-2300131312231320-2323313110111333-1012021022013033)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-3332002222313202-3123321100213021-1310123312032013-0101102313123323-2003030031022013-2212223201221323-0203203103020311-2032023003330023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2331320301203031-1332311001102320-0023020112300130-2311320112001131-0220113331323003-1011200233313203-0203211132303213-2302313121311121"></a>

## service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.tls_config — tls_config / 000120030321 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.port](resources--workload--reference--group-012.md#canonical-3013031222022030-2021222233122213-2220022232301003-3102220000332302-0311213311310132-3133321213021123-0133300222212330-1122131003023301)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-012.md#canonical-0101000213001300-0222020002101313-1020132032013321-2231011111202333-2122231103331210-3202123132002020-0313332003123002-3101311011132330)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert](resources--workload--reference--group-014.md#canonical-2233002113231233-1201011231003023-2310112000233003-0120011321233313-3122132210210203-2300131312231320-2323313110111333-1012021022013033)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.tls_config

<a id="canonical-0031130232033031-1030133010310200-3220303202232022-1010202200221312-2323021121333100-0230010000010130-3130322131001122-3102330220021110"></a>

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

<a id="canonical-1300233000011100-0131221233221312-0323311111123022-3021000303220130-1310313022211133-0311012031200212-0201113113111320-2231332301022303"></a>

## Direct properties — tls_config / 000120030321 / 3

- [custom_security](resources--workload--reference--group-014.md#canonical-0020310132012230-3303200221231333-1013303312100222-1023322323103110-2223213002033200-2231010122332102-3301210212100130-3130201203103332): complete subsection reference.

- [default_security](resources--workload--reference--group-014.md#canonical-2231213032120223-2103333333331022-2210201121323020-1012122232201112-2313311220202033-2230013022032001-0310113023020231-0100033120000023): complete subsection reference.

- [low_security](resources--workload--reference--group-014.md#canonical-0310211210032321-1233100200313323-1010013110311310-1031120212210013-3112023103300030-1033312233001333-3320310221100303-2133323030122033): complete subsection reference.

- [medium_security](resources--workload--reference--group-014.md#canonical-0002221021001331-3011122023010100-1022102030233023-0012111333012011-3121033122133231-3101112330101330-2232333033132302-0312300130202112): complete subsection reference.

<a id="canonical-0023322213031203-0101301200022323-1131313312121201-3100030120323311-2301133121302332-0112111232000120-1110310123132012-1102001322030012"></a>

## Next pages — tls_config / 000120030321 / 4

- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.tls_config.custom_security](resources--workload--reference--group-014.md#canonical-0020310132012230-3303200221231333-1013303312100222-1023322323103110-2223213002033200-2231010122332102-3301210212100130-3130201203103332)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.tls_config.default_security](resources--workload--reference--group-014.md#canonical-2231213032120223-2103333333331022-2210201121323020-1012122232201112-2313311220202033-2230013022032001-0310113023020231-0100033120000023)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.tls_config.low_security](resources--workload--reference--group-014.md#canonical-0310211210032321-1233100200313323-1010013110311310-1031120212210013-3112023103300030-1033312233001333-3320310221100303-2133323030122033)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.tls_config.medium_security](resources--workload--reference--group-014.md#canonical-0002221021001331-3011122023010100-1022102030233023-0012111333012011-3121033122133231-3101112330101330-2232333033132302-0312300130202112)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert](resources--workload--reference--group-014.md#canonical-2233002113231233-1201011231003023-2310112000233003-0120011321233313-3122132210210203-2300131312231320-2323313110111333-1012021022013033)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-0020310132012230-3303200221231333-1013303312100222-1023322323103110-2223213002033200-2231010122332102-3301210212100130-3130201203103332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3023120313323213-1113030111323111-0331013211313031-0003101321121020-3133202131233101-1323321013212002-2121301021302102-0031211113322021"></a>

## service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.tls_config.custom_security — custom_security / 300200030023 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.port](resources--workload--reference--group-012.md#canonical-3013031222022030-2021222233122213-2220022232301003-3102220000332302-0311213311310132-3133321213021123-0133300222212330-1122131003023301)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-012.md#canonical-0101000213001300-0222020002101313-1020132032013321-2231011111202333-2122231103331210-3202123132002020-0313332003123002-3101311011132330)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert](resources--workload--reference--group-014.md#canonical-2233002113231233-1201011231003023-2310112000233003-0120011321233313-3122132210210203-2300131312231320-2323313110111333-1012021022013033)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.tls_config](resources--workload--reference--group-014.md#canonical-3332002222313202-3123321100213021-1310123312032013-0101102313123323-2003030031022013-2212223201221323-0203203103020311-2032023003330023)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.tls_config.custom_security

<a id="canonical-0112030200133032-0033211000011021-1111003102313000-3101122302011032-1230033210021022-2333233002210033-0200020321023310-1221302311303021"></a>

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

<a id="canonical-0012223121233130-3112101123221013-2333033301221311-3201132321002320-1310001133003110-2130222012322322-1232023021021212-0202000110303030"></a>

## Direct properties — custom_security / 300200030023 / 3

<a id="canonical-0331220023301003-0121120021013212-0213323203303021-0331131012231100-1210032323013022-3020201213100332-1233012223131231-2121233221012201"></a>

<a id="canonical-0123133031023131-1320123022103233-2010102223330211-1010201113020033-0322331202200021-2221203132012321-0131011322302202-2003101213112111"></a>

## cipher_suites property — custom_security / 300200030023 / 4

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

<a id="canonical-3100021003113331-1130222001132313-0111030230022113-2300122132322021-3100033303300100-2013203120101213-3333323001033011-2310320330112110"></a>

<a id="canonical-2303032221302213-0211310200222013-3121220201033002-0313003203302130-0322200000120333-0132310132002313-1012223100230320-2032332201213223"></a>

## max_version property — custom_security / 300200030023 / 5

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

<a id="canonical-2013020123223310-0133012321111131-3030030012122023-0332210311323102-1212001221301021-2130330001031232-3221322322133011-2121033232113211"></a>

<a id="canonical-2310000303030311-2011101003002321-0103033321230233-0030203210222112-3032320023312012-2023322022010001-3333122230213200-3210300231213212"></a>

## min_version property — custom_security / 300200030023 / 6

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

<a id="canonical-0123002022121113-0100333202303101-3000330123301002-2302011301123000-2311121120030203-1000002322232311-3102231021122011-2022322012103303"></a>

## Next pages — custom_security / 300200030023 / 7

- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.tls_config](resources--workload--reference--group-014.md#canonical-3332002222313202-3123321100213021-1310123312032013-0101102313123323-2003030031022013-2212223201221323-0203203103020311-2032023003330023)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-2231213032120223-2103333333331022-2210201121323020-1012122232201112-2313311220202033-2230013022032001-0310113023020231-0100033120000023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2213130131201331-0132131113313310-2023303333332303-1032133122132310-1313122102030130-2023122223321202-0111230003233312-1221232302100302"></a>

## service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.tls_config.default_security — default_security / 122231001102 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.port](resources--workload--reference--group-012.md#canonical-3013031222022030-2021222233122213-2220022232301003-3102220000332302-0311213311310132-3133321213021123-0133300222212330-1122131003023301)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-012.md#canonical-0101000213001300-0222020002101313-1020132032013321-2231011111202333-2122231103331210-3202123132002020-0313332003123002-3101311011132330)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert](resources--workload--reference--group-014.md#canonical-2233002113231233-1201011231003023-2310112000233003-0120011321233313-3122132210210203-2300131312231320-2323313110111333-1012021022013033)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.tls_config](resources--workload--reference--group-014.md#canonical-3332002222313202-3123321100213021-1310123312032013-0101102313123323-2003030031022013-2212223201221323-0203203103020311-2032023003330023)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.tls_config.default_security

<a id="canonical-2033233022311233-1013223122321332-0002312132210121-2023002302120223-0310021022020232-1033011103003113-3010120021122020-2330010133213013"></a>

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

<a id="canonical-3322212301013310-0223233331320122-1130121132021021-2121132112011212-0232133132100103-2203031230200201-0000303323233323-3321212310300302"></a>

## Direct properties — default_security / 122231001102 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3100020021301310-2101010100312303-3131003210213013-3210223023132301-2223010022102131-0112120221021031-0120221013301023-1322231101201303"></a>

## Next pages — default_security / 122231001102 / 4

- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.tls_config](resources--workload--reference--group-014.md#canonical-3332002222313202-3123321100213021-1310123312032013-0101102313123323-2003030031022013-2212223201221323-0203203103020311-2032023003330023)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-0310211210032321-1233100200313323-1010013110311310-1031120212210013-3112023103300030-1033312233001333-3320310221100303-2133323030122033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3111230202032120-0001332010231102-1133302332232230-0102303033000033-0203001322120230-1220323132001020-2201212110233111-2233032222120013"></a>

## service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.tls_config.low_security — low_security / 100203121003 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.port](resources--workload--reference--group-012.md#canonical-3013031222022030-2021222233122213-2220022232301003-3102220000332302-0311213311310132-3133321213021123-0133300222212330-1122131003023301)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-012.md#canonical-0101000213001300-0222020002101313-1020132032013321-2231011111202333-2122231103331210-3202123132002020-0313332003123002-3101311011132330)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert](resources--workload--reference--group-014.md#canonical-2233002113231233-1201011231003023-2310112000233003-0120011321233313-3122132210210203-2300131312231320-2323313110111333-1012021022013033)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.tls_config](resources--workload--reference--group-014.md#canonical-3332002222313202-3123321100213021-1310123312032013-0101102313123323-2003030031022013-2212223201221323-0203203103020311-2032023003330023)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.tls_config.low_security

<a id="canonical-1322203330111003-1010302220201303-1030332021321320-1102210032331320-0030202030012020-1302013000131032-0232011201223322-2133210001311133"></a>

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

<a id="canonical-0200111200010312-0001313012332203-3013331222121131-1330333321210112-0312301113031301-2021023302322033-1023222001010312-0303120320313210"></a>

## Direct properties — low_security / 100203121003 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1110322201332321-1110220023323133-1300122012320010-2021220201300131-0202213331012132-0022113231210133-3311100221230011-1311333121202013"></a>

## Next pages — low_security / 100203121003 / 4

- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.tls_config](resources--workload--reference--group-014.md#canonical-3332002222313202-3123321100213021-1310123312032013-0101102313123323-2003030031022013-2212223201221323-0203203103020311-2032023003330023)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-0002221021001331-3011122023010100-1022102030233023-0012111333012011-3121033122133231-3101112330101330-2232333033132302-0312300130202112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3222010010213002-1222101100120202-2000303301031211-1203303321223113-3203320320003220-0133032200111130-3311120331323132-0033113232213200"></a>

## service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.tls_config.medium_security — medium_security / 230311333002 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.port](resources--workload--reference--group-012.md#canonical-3013031222022030-2021222233122213-2220022232301003-3102220000332302-0311213311310132-3133321213021123-0133300222212330-1122131003023301)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-012.md#canonical-0101000213001300-0222020002101313-1020132032013321-2231011111202333-2122231103331210-3202123132002020-0313332003123002-3101311011132330)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert](resources--workload--reference--group-014.md#canonical-2233002113231233-1201011231003023-2310112000233003-0120011321233313-3122132210210203-2300131312231320-2323313110111333-1012021022013033)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.tls_config](resources--workload--reference--group-014.md#canonical-3332002222313202-3123321100213021-1310123312032013-0101102313123323-2003030031022013-2212223201221323-0203203103020311-2032023003330023)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.tls_config.medium_security

<a id="canonical-0112202220231030-1010202120111031-0033011323220121-2230121000132201-3230000013300122-3133201311013111-0212313100023111-1103100303123100"></a>

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

<a id="canonical-3031231112020120-1122010022231123-0023312012312303-1132033130213032-3130210001210113-2233022300021331-3222010200022021-2122323110310131"></a>

## Direct properties — medium_security / 230311333002 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0312021312123013-0131222102011320-2102000031112211-2230313313123221-1030233100101230-1332302212313231-2112013220201233-2310232022321201"></a>

## Next pages — medium_security / 230311333002 / 4

- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.tls_config](resources--workload--reference--group-014.md#canonical-3332002222313202-3123321100213021-1310123312032013-0101102313123323-2003030031022013-2212223201221323-0203203103020311-2032023003330023)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-1101302220013000-1130323120002221-2213110022303302-1002100102000211-3203030111202113-1021233103221020-3020210323320220-3011000311311022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2331012033020012-2023101132030202-0011230121130010-1332210221012011-2030113201310211-3310020210200021-1321001211210022-2113300330222011"></a>

## service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.use_mtls — use_mtls / 331123330022 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.port](resources--workload--reference--group-012.md#canonical-3013031222022030-2021222233122213-2220022232301003-3102220000332302-0311213311310132-3133321213021123-0133300222212330-1122131003023301)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-012.md#canonical-0101000213001300-0222020002101313-1020132032013321-2231011111202333-2122231103331210-3202123132002020-0313332003123002-3101311011132330)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert](resources--workload--reference--group-014.md#canonical-2233002113231233-1201011231003023-2310112000233003-0120011321233313-3122132210210203-2300131312231320-2323313110111333-1012021022013033)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.use_mtls

<a id="canonical-0313002030203123-3300220111123012-1130333122103202-2311223031332133-2320021302223213-1231303133202110-0112000003201322-3331123033001001"></a>

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

<a id="canonical-3133121223333011-1302233203320233-2010313121231313-2202233332332130-2011322031213113-1301133312102010-0113032000210302-2333111123122022"></a>

## Direct properties — use_mtls / 331123330022 / 3

<a id="canonical-3302311300023032-3311313122033000-1202101113311212-3211132002131213-1021100230022203-3322211333101232-2311320321213103-0021102030200223"></a>

<a id="canonical-3032122030223011-2200011001310020-1210013230202021-3001200120111110-0311031123123000-1330301132220323-0220321100332333-1021332020013230"></a>

## client_certificate_optional property — use_mtls / 331123330022 / 4

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

- [crl](resources--workload--reference--group-014.md#canonical-3231001020202230-1311220132312310-0331130202123200-2211210110013210-0020001300323200-1002231323230303-3010330121132012-0133230312011032): complete subsection reference.

- [no_crl](resources--workload--reference--group-014.md#canonical-2132123100221201-0222212200310002-2212103010111231-0023100022313322-2013222101300233-3312221011213301-3303211220320111-3011321011113110): complete subsection reference.

- [trusted_ca](resources--workload--reference--group-014.md#canonical-2002231111221303-3223031203010300-3211032220031330-2221230212223312-1332123310322112-2330211323202210-2323130130021110-1012211031132101): complete subsection reference.

<a id="canonical-1012023111011201-3332132332310330-3000301211211123-0232221220232030-2212331212131212-0000302321000103-1322013300201303-1233233131230332"></a>

<a id="canonical-2030313003130203-0030310222231020-2131223031323330-2200211331201231-0232301120111121-0100330313113312-0101320211321013-3210002033000101"></a>

## trusted_ca_url property — use_mtls / 331123330022 / 5

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

- [xfcc_disabled](resources--workload--reference--group-014.md#canonical-2000220010332232-1002031302003133-0211300231112220-1010100213023031-3023123211220222-1210302312022033-1103002223021132-2133001333112230): complete subsection reference.

- [xfcc_options](resources--workload--reference--group-014.md#canonical-1232003011022332-1003223100200021-2212320311030031-3300123212003120-2202111030100301-2301302312313121-2312211132201022-2311103031021012): complete subsection reference.

<a id="canonical-3312123132032302-3323113323300223-1201201023303123-1013012131021332-0132011333223003-1130003110321212-1113130123322311-0320112100220230"></a>

## Next pages — use_mtls / 331123330022 / 6

- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.use_mtls.crl](resources--workload--reference--group-014.md#canonical-3231001020202230-1311220132312310-0331130202123200-2211210110013210-0020001300323200-1002231323230303-3010330121132012-0133230312011032)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.use_mtls.no_crl](resources--workload--reference--group-014.md#canonical-2132123100221201-0222212200310002-2212103010111231-0023100022313322-2013222101300233-3312221011213301-3303211220320111-3011321011113110)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.use_mtls.trusted_ca](resources--workload--reference--group-014.md#canonical-2002231111221303-3223031203010300-3211032220031330-2221230212223312-1332123310322112-2330211323202210-2323130130021110-1012211031132101)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.use_mtls.xfcc_disabled](resources--workload--reference--group-014.md#canonical-2000220010332232-1002031302003133-0211300231112220-1010100213023031-3023123211220222-1210302312022033-1103002223021132-2133001333112230)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.use_mtls.xfcc_options](resources--workload--reference--group-014.md#canonical-1232003011022332-1003223100200021-2212320311030031-3300123212003120-2202111030100301-2301302312313121-2312211132201022-2311103031021012)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert](resources--workload--reference--group-014.md#canonical-2233002113231233-1201011231003023-2310112000233003-0120011321233313-3122132210210203-2300131312231320-2323313110111333-1012021022013033)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-3231001020202230-1311220132312310-0331130202123200-2211210110013210-0020001300323200-1002231323230303-3010330121132012-0133230312011032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2013202301321101-0022301333303123-3223230233110310-0333121110212333-0010003101100030-2122031010133323-1300113322210032-1201110330031030"></a>

## service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.use_mtls.crl — crl / 001311333111 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.port](resources--workload--reference--group-012.md#canonical-3013031222022030-2021222233122213-2220022232301003-3102220000332302-0311213311310132-3133321213021123-0133300222212330-1122131003023301)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-012.md#canonical-0101000213001300-0222020002101313-1020132032013321-2231011111202333-2122231103331210-3202123132002020-0313332003123002-3101311011132330)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert](resources--workload--reference--group-014.md#canonical-2233002113231233-1201011231003023-2310112000233003-0120011321233313-3122132210210203-2300131312231320-2323313110111333-1012021022013033)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.use_mtls](resources--workload--reference--group-014.md#canonical-1101302220013000-1130323120002221-2213110022303302-1002100102000211-3203030111202113-1021233103221020-3020210323320220-3011000311311022)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.use_mtls.crl

<a id="canonical-3032021211121111-2023111013333222-0020122330000133-2103330322333023-0300211033210320-3201321321102323-0300102000333222-3012311232331232"></a>

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

<a id="canonical-0202220122301321-2132010323123111-3120123003020301-2311201210102112-0133312112022103-2302210311131233-3320213110213132-1021000032120122"></a>

## Direct properties — crl / 001311333111 / 3

<a id="canonical-3300112023300331-1231200011022311-2002033301020232-3221202002332032-3331320223331223-0032312020102330-1223321111110202-3223003202201120"></a>

<a id="canonical-1331213113033322-2200112203210030-1013220300302212-0121123020200001-2330032211301133-0320311022232232-3030103301100021-1000322330303321"></a>

## name property — crl / 001311333111 / 4

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

<a id="canonical-3133302032020131-0212020312020112-1131300232303320-2133032210031123-0212002212120221-2202331012202121-2022230012022303-2021323011232013"></a>

<a id="canonical-3130103111112021-1022221022331332-3300101131200133-0021330101323010-0310011331330111-0031001002012200-3310301322203231-2201021232200002"></a>

## namespace property — crl / 001311333111 / 5

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

<a id="canonical-3131020000110303-0003000023120123-0320313310311223-1132300132232332-0021103101023121-1020212232133331-3330212103322020-1200121223302010"></a>

<a id="canonical-1021011202311032-2002121121010202-0002210022102023-2013003023102010-1213110003013103-2002020022031212-2033130231202333-0230311110132330"></a>

## tenant property — crl / 001311333111 / 6

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

<a id="canonical-1031230003331312-1302031210231013-1032131303323210-1121333221001220-1030122112231313-0210203310123322-0311332123301300-1232202030323001"></a>

## Next pages — crl / 001311333111 / 7

- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.use_mtls](resources--workload--reference--group-014.md#canonical-1101302220013000-1130323120002221-2213110022303302-1002100102000211-3203030111202113-1021233103221020-3020210323320220-3011000311311022)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-2132123100221201-0222212200310002-2212103010111231-0023100022313322-2013222101300233-3312221011213301-3303211220320111-3011321011113110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2312300033202002-0022032130132213-1210122123203020-1212203101230233-1113000212202131-3200322322311232-0232132323022021-1010122103302303"></a>

## service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.use_mtls.no_crl — no_crl / 321321323232 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.port](resources--workload--reference--group-012.md#canonical-3013031222022030-2021222233122213-2220022232301003-3102220000332302-0311213311310132-3133321213021123-0133300222212330-1122131003023301)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-012.md#canonical-0101000213001300-0222020002101313-1020132032013321-2231011111202333-2122231103331210-3202123132002020-0313332003123002-3101311011132330)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert](resources--workload--reference--group-014.md#canonical-2233002113231233-1201011231003023-2310112000233003-0120011321233313-3122132210210203-2300131312231320-2323313110111333-1012021022013033)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.use_mtls](resources--workload--reference--group-014.md#canonical-1101302220013000-1130323120002221-2213110022303302-1002100102000211-3203030111202113-1021233103221020-3020210323320220-3011000311311022)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.use_mtls.no_crl

<a id="canonical-3111323200130133-2210012011121111-0030223030003222-1323113322222330-2112323031200233-3311133102130301-0131323020301033-2121020230233023"></a>

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

<a id="canonical-1010200131331330-1120210103331312-2230022102131212-0113231113321320-0320001132012332-3301010212023112-0332020221010022-0320210033113221"></a>

## Direct properties — no_crl / 321321323232 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1322202203023033-0132212231200303-0320122010222031-3203123131311112-1220112312320011-3032310312023121-1113131332133211-2301012103121321"></a>

## Next pages — no_crl / 321321323232 / 4

- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.use_mtls](resources--workload--reference--group-014.md#canonical-1101302220013000-1130323120002221-2213110022303302-1002100102000211-3203030111202113-1021233103221020-3020210323320220-3011000311311022)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-2002231111221303-3223031203010300-3211032220031330-2221230212223312-1332123310322112-2330211323202210-2323130130021110-1012211031132101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2130100320112001-3202023023222220-3030210132300311-2332301312102133-2333023311332100-3031032023300211-1000121333032132-2322300223101022"></a>

## service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.use_mtls.trusted_ca — trusted_ca / 210120300100 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.port](resources--workload--reference--group-012.md#canonical-3013031222022030-2021222233122213-2220022232301003-3102220000332302-0311213311310132-3133321213021123-0133300222212330-1122131003023301)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-012.md#canonical-0101000213001300-0222020002101313-1020132032013321-2231011111202333-2122231103331210-3202123132002020-0313332003123002-3101311011132330)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert](resources--workload--reference--group-014.md#canonical-2233002113231233-1201011231003023-2310112000233003-0120011321233313-3122132210210203-2300131312231320-2323313110111333-1012021022013033)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.use_mtls](resources--workload--reference--group-014.md#canonical-1101302220013000-1130323120002221-2213110022303302-1002100102000211-3203030111202113-1021233103221020-3020210323320220-3011000311311022)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.use_mtls.trusted_ca

<a id="canonical-2133320023111123-0001321323121023-3232000110313200-3320032022313330-1222302220132002-2011202320333200-0311212311000012-3202330110130222"></a>

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

<a id="canonical-0130201321002130-1223331223231233-1020030302320203-3301120130332010-2103212113002310-3122123320233313-2332012323133321-2303123010120123"></a>

## Direct properties — trusted_ca / 210120300100 / 3

<a id="canonical-0221112200233121-2010010130010332-1320221123320210-2101011330131332-1331332110221120-1100100031210032-0211031103223223-0213011321003231"></a>

<a id="canonical-2130320032112203-3330232123302013-1203223231023121-2203032303312302-3201022113013320-1320232220033323-1303302203310003-0012322223332222"></a>

## name property — trusted_ca / 210120300100 / 4

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

<a id="canonical-0213301231111103-1103100030220021-3002231013133102-2120003332232111-2100100323012133-1312010120102330-0303313012133000-2012221010022011"></a>

<a id="canonical-1310021130102132-2121300033323001-2223003031230031-2021012032332030-2300133331231213-1103332110201212-2321121133101102-0132011010121120"></a>

## namespace property — trusted_ca / 210120300100 / 5

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

<a id="canonical-1210131310010231-1110031210013220-2230000311133233-3201003302111010-2203211223330022-2201103131002333-1300323010300132-2111001131103223"></a>

<a id="canonical-0323011103200100-1112221011320111-0133102102313311-3221201011001210-0233300203321132-1311331031231101-2133000100300130-3323030122333231"></a>

## tenant property — trusted_ca / 210120300100 / 6

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

<a id="canonical-0031232210020201-3111222223212003-3321333320233201-2033030122331021-3232303310200301-3120103023021122-3312122303211230-2033020310220023"></a>

## Next pages — trusted_ca / 210120300100 / 7

- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.use_mtls](resources--workload--reference--group-014.md#canonical-1101302220013000-1130323120002221-2213110022303302-1002100102000211-3203030111202113-1021233103221020-3020210323320220-3011000311311022)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-2000220010332232-1002031302003133-0211300231112220-1010100213023031-3023123211220222-1210302312022033-1103002223021132-2133001333112230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1310022322010130-3021211021101001-1201233301132020-1302010332020111-0002311220211222-2011120332330201-3233212023310332-1122201221201222"></a>

## service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.use_mtls.xfcc_disabled — xfcc_disabled / 220210311232 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.port](resources--workload--reference--group-012.md#canonical-3013031222022030-2021222233122213-2220022232301003-3102220000332302-0311213311310132-3133321213021123-0133300222212330-1122131003023301)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-012.md#canonical-0101000213001300-0222020002101313-1020132032013321-2231011111202333-2122231103331210-3202123132002020-0313332003123002-3101311011132330)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert](resources--workload--reference--group-014.md#canonical-2233002113231233-1201011231003023-2310112000233003-0120011321233313-3122132210210203-2300131312231320-2323313110111333-1012021022013033)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.use_mtls](resources--workload--reference--group-014.md#canonical-1101302220013000-1130323120002221-2213110022303302-1002100102000211-3203030111202113-1021233103221020-3020210323320220-3011000311311022)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.use_mtls.xfcc_disabled

<a id="canonical-1013213023100230-1130101112001032-1332111123302313-1033202020230132-1302131213202002-2322303122132121-2212231103032100-1213120003123110"></a>

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

<a id="canonical-3112113112013200-0202213312022032-1000110302200012-3221112111100102-0121100110103202-3201311032203332-1331021212021203-3022031303111311"></a>

## Direct properties — xfcc_disabled / 220210311232 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0022323220111102-0012110100103101-1331033212022020-1302103231130133-0332202311213203-1201313201311212-0300111012023113-3201030001002222"></a>

## Next pages — xfcc_disabled / 220210311232 / 4

- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.use_mtls](resources--workload--reference--group-014.md#canonical-1101302220013000-1130323120002221-2213110022303302-1002100102000211-3203030111202113-1021233103221020-3020210323320220-3011000311311022)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-1232003011022332-1003223100200021-2212320311030031-3300123212003120-2202111030100301-2301302312313121-2312211132201022-2311103031021012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0312020322231230-1333210320033311-1010312002121303-0310132010300103-3002333200123132-0003322133303201-0023211013332202-2122211111133301"></a>

## service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.use_mtls.xfcc_options — xfcc_options / 022001111201 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.port](resources--workload--reference--group-012.md#canonical-3013031222022030-2021222233122213-2220022232301003-3102220000332302-0311213311310132-3133321213021123-0133300222212330-1122131003023301)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-012.md#canonical-0101000213001300-0222020002101313-1020132032013321-2231011111202333-2122231103331210-3202123132002020-0313332003123002-3101311011132330)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert](resources--workload--reference--group-014.md#canonical-2233002113231233-1201011231003023-2310112000233003-0120011321233313-3122132210210203-2300131312231320-2323313110111333-1012021022013033)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.use_mtls](resources--workload--reference--group-014.md#canonical-1101302220013000-1130323120002221-2213110022303302-1002100102000211-3203030111202113-1021233103221020-3020210323320220-3011000311311022)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.use_mtls.xfcc_options

<a id="canonical-3231331312011111-3123311310220331-2301022131223110-2231313300122032-2102033311110022-0032302002200211-3003031220003130-2332010131210003"></a>

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

<a id="canonical-0112202123312102-0201322111003203-2010201131223130-3021313133220211-2333220232331120-3110322203011210-0022132331101301-0201202000300123"></a>

## Direct properties — xfcc_options / 022001111201 / 3

<a id="canonical-1320103023112300-3003221120012021-1023323222100233-2321313232031323-2000101000030303-1333222003213200-0300302000002310-3221300022222202"></a>

<a id="canonical-0203001023200023-1101312102003013-0033201130032302-3001101331033323-0232230320210311-0013310021133321-0122130230202021-2023223032230013"></a>

## xfcc_header_elements property — xfcc_options / 022001111201 / 4

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

<a id="canonical-2310131101230312-0012320123023120-0320203221221122-0210233012012212-0301122020301211-3111310320022200-3021001313023323-1321321312001131"></a>

## Next pages — xfcc_options / 022001111201 / 5

- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.use_mtls](resources--workload--reference--group-014.md#canonical-1101302220013000-1130323120002221-2213110022303302-1002100102000211-3203030111202113-1021233103221020-3020210323320220-3011000311311022)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-1223131333131301-1101202200212003-2303000010310012-3003131221200020-0201223130203322-2303023001222013-1231100011211030-2032330121100311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3223203011102233-2100300021001230-3002023010133220-0003113301112303-3302213320301030-2333023110231312-3320132221203002-3301311301300110"></a>

## service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes — specific_routes / 123102030120 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.port](resources--workload--reference--group-012.md#canonical-3013031222022030-2021222233122213-2220022232301003-3102220000332302-0311213311310132-3133321213021123-0133300222212330-1122131003023301)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-012.md#canonical-0101000213001300-0222020002101313-1020132032013321-2231011111202333-2122231103331210-3202123132002020-0313332003123002-3101311011132330)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes

<a id="canonical-1120303110123101-1303103010332000-2200322120000300-1032013130201110-1110312101200102-2213232203101011-0113103033001031-3020222323212030"></a>

Type: `"object"`. single nested block, Optional.

Defines various OPTIONS to define a route.

Upstream description:

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

<a id="canonical-1212332311021120-3112221322113000-0023201101211320-2333022133023133-2211022230303202-0330000303132032-2201222230300001-3001103003133031"></a>

## Direct properties — specific_routes / 123102030120 / 3

- [routes](resources--workload--reference--group-014.md#canonical-1320302000201200-1013121233121332-1200303031313333-2131200212100310-2031222303000133-0013120300210320-3032201103202020-0103102312330232): complete subsection reference.

<a id="canonical-2211211031332200-2011221313101132-1021310011103300-2111301211103310-1311011101011032-1203203110110310-3010031100311212-0132301030202302"></a>

## Next pages — specific_routes / 123102030120 / 4

- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes](resources--workload--reference--group-014.md#canonical-1320302000201200-1013121233121332-1200303031313333-2131200212100310-2031222303000133-0013120300210320-3032201103202020-0103102312330232)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-012.md#canonical-0101000213001300-0222020002101313-1020132032013321-2231011111202333-2122231103331210-3202123132002020-0313332003123002-3101311011132330)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-1320302000201200-1013121233121332-1200303031313333-2131200212100310-2031222303000133-0013120300210320-3032201103202020-0103102312330232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->
