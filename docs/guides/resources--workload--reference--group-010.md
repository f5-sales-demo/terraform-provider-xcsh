---
page_title: "xcsh_workload reference"
subcategory: "Container"
description: "Complete grouped canonical reference for xcsh_workload reference."
---

# xcsh_workload reference

<a id="canonical-2200102323322311-2321310211002100-0110312020033211-1030231203313322-0021230222030221-0022320300100102-1111001203230132-0333002002123010"></a>

## service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.use_mtls.trusted_ca — trusted_ca / 220032302111 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-008.md#canonical-3210102113031103-2332201101220100-1103122221002212-3002101012001000-1332031130002011-0021231133303122-2121222212212213-1123320220010203)
- [service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-008.md#canonical-2311132233010120-0223230230122210-0332200203312210-3000132111100333-1133101301233010-3211310033122333-2323002110210000-1300012002110122)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](resources--workload--reference--group-008.md#canonical-1202020123020333-2332300320113330-1321112113331202-2112220001121101-3022001013310321-0020100001212100-0032232320012323-1311122212310031)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](resources--workload--reference--group-009.md#canonical-3212310132100213-2000301220120030-0112323230310220-3200032133332210-3223310210320320-2331013310222010-2010102133130322-2111123310321223)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params](resources--workload--reference--group-009.md#canonical-0212122113133333-1000102320023221-2003032220233320-0021231111201312-3001031012012332-3313221303003312-0112011332100001-1120030302022032)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.use_mtls](resources--workload--reference--group-009.md#canonical-2230323133110331-2101302313123301-2033301333010232-3320123100210000-1220000110330331-3301011020212033-1110022310322130-3132201012030322)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.use_mtls.trusted_ca

<a id="canonical-1132213132202022-2000320103321320-3222222210333320-2100201312312222-2010123333332001-1211020003022032-1030320110300002-2103213033011131"></a>

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

<a id="canonical-2022212200321030-3001211230223223-2132321221220000-2233320322302313-1213320123010211-0302132120120303-0222200110122212-1031232130201311"></a>

## Direct properties — trusted_ca / 220032302111 / 3

<a id="canonical-2002121112023323-1122212212111203-0032121120213201-3130310213001232-2332200222211102-1032111200323333-2221023230120333-3200212100200130"></a>

<a id="canonical-2033033123310133-1030012020203323-3130213132120221-1213102013002101-2002011202232021-2220033230211031-2320332023300220-1220121311112313"></a>

## name property — trusted_ca / 220032302111 / 4

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-3212101011021021-1130332313301233-0122131033000110-0122100323033123-0131331111332320-3203133302031311-3101310123232330-3310123120312131"></a>

<a id="canonical-3111323100013102-0032131333100210-0023023112320312-3322303112200212-0202032333221223-2002110210310322-0132201101001131-2031033313122010"></a>

## namespace property — trusted_ca / 220032302111 / 5

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-3232121010101110-1032201133011232-1122011010233212-1220211331101331-1212231301023132-3110211322211223-0312312213232112-2301031022310023"></a>

<a id="canonical-2202123100231132-0112012012010123-2131002100202320-1033330210313222-3232310313003211-3033122300303010-2023213112221333-0331030212120021"></a>

## tenant property — trusted_ca / 220032302111 / 6

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-2001312303101201-3231223103332033-2033112113330101-1133323331110310-2302132201022230-0010213111120313-1011013321313012-2100132202211323"></a>

## Next pages — trusted_ca / 220032302111 / 7

- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.use_mtls](resources--workload--reference--group-009.md#canonical-2230323133110331-2101302313123301-2033301333010232-3320123100210000-1220000110330331-3301011020212033-1110022310322130-3132201012030322)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-0210002101312203-3033320003012002-3021222323013232-0223230020202010-0000203131132121-0231232312131022-1001001110320120-2033301233303132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2001201032211122-3020030003033113-3220310022003302-2020130111302233-3020122032102312-1130330330302123-3321031300230320-3011011120231223"></a>

## service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.use_mtls.xfcc_disabled — xfcc_disabled / 320103230023 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-008.md#canonical-3210102113031103-2332201101220100-1103122221002212-3002101012001000-1332031130002011-0021231133303122-2121222212212213-1123320220010203)
- [service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-008.md#canonical-2311132233010120-0223230230122210-0332200203312210-3000132111100333-1133101301233010-3211310033122333-2323002110210000-1300012002110122)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](resources--workload--reference--group-008.md#canonical-1202020123020333-2332300320113330-1321112113331202-2112220001121101-3022001013310321-0020100001212100-0032232320012323-1311122212310031)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](resources--workload--reference--group-009.md#canonical-3212310132100213-2000301220120030-0112323230310220-3200032133332210-3223310210320320-2331013310222010-2010102133130322-2111123310321223)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params](resources--workload--reference--group-009.md#canonical-0212122113133333-1000102320023221-2003032220233320-0021231111201312-3001031012012332-3313221303003312-0112011332100001-1120030302022032)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.use_mtls](resources--workload--reference--group-009.md#canonical-2230323133110331-2101302313123301-2033301333010232-3320123100210000-1220000110330331-3301011020212033-1110022310322130-3132201012030322)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.use_mtls.xfcc_disabled

<a id="canonical-0130022110301111-0102021021020122-0013330211120122-3030230021310301-0013311320221001-1131203031233101-3100132011023102-2220033332330021"></a>

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

<a id="canonical-3121033021001102-0100120000300013-2201303210322000-2302012331322031-3211221122120000-0120110120103003-2213230311122312-0202321013332011"></a>

## Direct properties — xfcc_disabled / 320103230023 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2201022100101220-2311232002122023-0213003310130003-3021111203023130-0112212202310212-3020330122020001-3201021302310020-0212222101012322"></a>

## Next pages — xfcc_disabled / 320103230023 / 4

- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.use_mtls](resources--workload--reference--group-009.md#canonical-2230323133110331-2101302313123301-2033301333010232-3320123100210000-1220000110330331-3301011020212033-1110022310322130-3132201012030322)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-3331213103211210-0022333013121122-0010123223112003-0121131311131120-3120203200321230-0032102133022223-2031223033331220-3223022211203201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2120312301300110-0111222332020212-3021120020001220-1100121203302112-1223302001323300-0331303223101021-0233130321002221-0230031310013002"></a>

## service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.use_mtls.xfcc_options — xfcc_options / 011310002333 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-008.md#canonical-3210102113031103-2332201101220100-1103122221002212-3002101012001000-1332031130002011-0021231133303122-2121222212212213-1123320220010203)
- [service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-008.md#canonical-2311132233010120-0223230230122210-0332200203312210-3000132111100333-1133101301233010-3211310033122333-2323002110210000-1300012002110122)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](resources--workload--reference--group-008.md#canonical-1202020123020333-2332300320113330-1321112113331202-2112220001121101-3022001013310321-0020100001212100-0032232320012323-1311122212310031)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](resources--workload--reference--group-009.md#canonical-3212310132100213-2000301220120030-0112323230310220-3200032133332210-3223310210320320-2331013310222010-2010102133130322-2111123310321223)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params](resources--workload--reference--group-009.md#canonical-0212122113133333-1000102320023221-2003032220233320-0021231111201312-3001031012012332-3313221303003312-0112011332100001-1120030302022032)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.use_mtls](resources--workload--reference--group-009.md#canonical-2230323133110331-2101302313123301-2033301333010232-3320123100210000-1220000110330331-3301011020212033-1110022310322130-3132201012030322)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.use_mtls.xfcc_options

<a id="canonical-2033201230301212-2003010030100011-2130122112002112-0102221021101210-1012030000300223-0021111100213012-0302121223031330-0311333100231122"></a>

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

<a id="canonical-3311332331021311-3202021032233020-2210021312123303-3030330220231332-3203002311100022-2322323100230013-3313121123031331-0100132222221222"></a>

## Direct properties — xfcc_options / 011310002333 / 3

<a id="canonical-0331131033231213-1031311132231100-2213121003312222-2233222200103022-0221111000213322-3232312130122022-2011133003111212-2333300332210233"></a>

<a id="canonical-3110123320203113-2320030100001223-0210320311012020-0331200113120100-3330123032203203-2112200330321000-2303312121203310-0121030233320131"></a>

## xfcc_header_elements property — xfcc_options / 011310002333 / 4

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

<a id="canonical-0300230123322101-1113310221010120-2103202313320201-3010002121131220-3000012231200231-1102122020020113-0001132211003201-3023221011220013"></a>

## Next pages — xfcc_options / 011310002333 / 5

- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.use_mtls](resources--workload--reference--group-009.md#canonical-2230323133110331-2101302313123301-2033301333010232-3320123100210000-1220000110330331-3301011020212033-1110022310322130-3132201012030322)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-3230132323111312-0132020110123203-0222302101120231-0202012321012303-1213310020220000-0101203211110131-1303320330302033-3001313221002112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0122001220002331-0233031101011330-2123020231001301-2111121233022221-2220012021110122-2330122213033202-0100133123033333-3111331001302201"></a>

## service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters — tls_parameters / 023103013001 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-008.md#canonical-3210102113031103-2332201101220100-1103122221002212-3002101012001000-1332031130002011-0021231133303122-2121222212212213-1123320220010203)
- [service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-008.md#canonical-2311132233010120-0223230230122210-0332200203312210-3000132111100333-1133101301233010-3211310033122333-2323002110210000-1300012002110122)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](resources--workload--reference--group-008.md#canonical-1202020123020333-2332300320113330-1321112113331202-2112220001121101-3022001013310321-0020100001212100-0032232320012323-1311122212310031)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](resources--workload--reference--group-009.md#canonical-3212310132100213-2000301220120030-0112323230310220-3200032133332210-3223310210320320-2331013310222010-2010102133130322-2111123310321223)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters

<a id="canonical-0230231021123011-1203213213330303-0232233313101130-1020223203331011-0222112301130223-3030232300333010-3301102100020213-1220120311111332"></a>

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

<a id="canonical-1200100003131133-1123023132332322-2233330300231200-1000310032200100-3331121032022232-1033121020100202-0301003332112232-2011302201002133"></a>

## Direct properties — tls_parameters / 023103013001 / 3

- [no_mtls](resources--workload--reference--group-010.md#canonical-1100310113021011-2300121112223230-3322231112002312-1010012320233101-3131033110123112-0010231221033102-2012213121201332-2321131202330003): complete subsection reference.

- [tls_certificates](resources--workload--reference--group-010.md#canonical-0220313011200113-2210220130122300-1020333311303231-0221222302220012-3120013102133220-2132302001202230-2322133133130313-1313320213312221): complete subsection reference.

- [tls_config](resources--workload--reference--group-010.md#canonical-1233203302211212-3300120033202101-0200330001312321-1222133023031223-1233313002123211-3203212322123012-1131031112003021-1331232103023221): complete subsection reference.

- [use_mtls](resources--workload--reference--group-010.md#canonical-3310220202003120-0020232310001311-1230011321302031-3002200220020220-1002102031303233-0322332031301112-0031333023233103-0133102211311201): complete subsection reference.

<a id="canonical-1333203233122311-1330020103223011-0122122131033211-0210222131331021-2011232103310103-2123223033213310-3320112121312130-0031021113021320"></a>

## Next pages — tls_parameters / 023103013001 / 4

- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.no_mtls](resources--workload--reference--group-010.md#canonical-1100310113021011-2300121112223230-3322231112002312-1010012320233101-3131033110123112-0010231221033102-2012213121201332-2321131202330003)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_certificates](resources--workload--reference--group-010.md#canonical-0220313011200113-2210220130122300-1020333311303231-0221222302220012-3120013102133220-2132302001202230-2322133133130313-1313320213312221)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_config](resources--workload--reference--group-010.md#canonical-1233203302211212-3300120033202101-0200330001312321-1222133023031223-1233313002123211-3203212322123012-1131031112003021-1331232103023221)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.use_mtls](resources--workload--reference--group-010.md#canonical-3310220202003120-0020232310001311-1230011321302031-3002200220020220-1002102031303233-0322332031301112-0031333023233103-0133102211311201)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](resources--workload--reference--group-009.md#canonical-3212310132100213-2000301220120030-0112323230310220-3200032133332210-3223310210320320-2331013310222010-2010102133130322-2111123310321223)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-1100310113021011-2300121112223230-3322231112002312-1010012320233101-3131033110123112-0010231221033102-2012213121201332-2321131202330003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3030200101032020-3221302111132201-3010100000232322-2021302102133013-2230131213033023-0132302300131221-0111332003130002-0323122301320102"></a>

## service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.no_mtls — no_mtls / 110223233023 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-008.md#canonical-3210102113031103-2332201101220100-1103122221002212-3002101012001000-1332031130002011-0021231133303122-2121222212212213-1123320220010203)
- [service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-008.md#canonical-2311132233010120-0223230230122210-0332200203312210-3000132111100333-1133101301233010-3211310033122333-2323002110210000-1300012002110122)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](resources--workload--reference--group-008.md#canonical-1202020123020333-2332300320113330-1321112113331202-2112220001121101-3022001013310321-0020100001212100-0032232320012323-1311122212310031)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](resources--workload--reference--group-009.md#canonical-3212310132100213-2000301220120030-0112323230310220-3200032133332210-3223310210320320-2331013310222010-2010102133130322-2111123310321223)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-010.md#canonical-3230132323111312-0132020110123203-0222302101120231-0202012321012303-1213310020220000-0101203211110131-1303320330302033-3001313221002112)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.no_mtls

<a id="canonical-0200002222113033-0202331111220101-1122122022313311-0012110333111330-0023300203321132-2013031310031333-0331302322210222-2030201112202301"></a>

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

<a id="canonical-1331323322231100-1013212213100033-1023021223112301-3320132002233012-2033013100310220-1003102332322202-2212212101121011-0231100033212010"></a>

## Direct properties — no_mtls / 110223233023 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3020232220032321-2313320333031101-2201232300312303-1013331020220023-0132102111031112-1012121110222331-1221032012120012-1212201130300102"></a>

## Next pages — no_mtls / 110223233023 / 4

- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-010.md#canonical-3230132323111312-0132020110123203-0222302101120231-0202012321012303-1213310020220000-0101203211110131-1303320330302033-3001313221002112)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-0220313011200113-2210220130122300-1020333311303231-0221222302220012-3120013102133220-2132302001202230-2322133133130313-1313320213312221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2310111233113013-1012032030330032-2001031233320301-0120212223302220-0212231210110121-2112230133122213-0010200121002100-3121331123130332"></a>

## service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_certificates — tls_certificates / 012311030300 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-008.md#canonical-3210102113031103-2332201101220100-1103122221002212-3002101012001000-1332031130002011-0021231133303122-2121222212212213-1123320220010203)
- [service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-008.md#canonical-2311132233010120-0223230230122210-0332200203312210-3000132111100333-1133101301233010-3211310033122333-2323002110210000-1300012002110122)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](resources--workload--reference--group-008.md#canonical-1202020123020333-2332300320113330-1321112113331202-2112220001121101-3022001013310321-0020100001212100-0032232320012323-1311122212310031)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](resources--workload--reference--group-009.md#canonical-3212310132100213-2000301220120030-0112323230310220-3200032133332210-3223310210320320-2331013310222010-2010102133130322-2111123310321223)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-010.md#canonical-3230132323111312-0132020110123203-0222302101120231-0202012321012303-1213310020220000-0101203211110131-1303320330302033-3001313221002112)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_certificates

<a id="canonical-1232300110032012-1132221313133331-3020221130222001-1202021230222110-1030010330221203-3012013321232030-1020030202022211-2022301031112332"></a>

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-0231210012313123-3313321210320132-3131330223230032-1122211231323203-2322033223031313-2021213211210212-0101213122000032-2323112333002211"></a>

## Direct properties — tls_certificates / 012311030300 / 3

<a id="canonical-0310302211003333-0300220330123103-3301000231222130-1301000312332033-1332333022322220-2113020203133030-0211110101302230-3220111212103022"></a>

<a id="canonical-3020200213131103-3001011211111011-1313222310013102-2023303103313213-0301203232000223-1003012300033303-2302230031311003-1010231321201101"></a>

## certificate_url property — tls_certificates / 012311030300 / 4

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

- [custom_hash_algorithms](resources--workload--reference--group-010.md#canonical-1100332230200022-1320301123211133-3301233012201000-2212312020330103-1313303313301322-2011022030012011-3021030200200130-2121002333332023): complete subsection reference.

<a id="canonical-3201312233121311-1312222330331122-3231212203323222-3100311303311231-0330313131011003-0000031121002310-3023220002310322-1211020202210301"></a>

<a id="canonical-3113110230310000-0132112232133120-1033013331332121-3002220121323313-1223202322302230-1100201312231300-2333103123123312-0112133213233200"></a>

## description_spec property — tls_certificates / 012311030300 / 5

Type: `"string"`. Optional.

Description. Description for the certificate.

- [disable_ocsp_stapling](resources--workload--reference--group-010.md#canonical-1120200202033020-3301221130122030-1333020301210113-2112321001230103-1131212222223230-2211110200221131-2121200222211120-3011112002022220): complete subsection reference.

- [private_key](resources--workload--reference--group-010.md#canonical-0000103021020013-2333222020321111-2221322222122222-1012301222211232-3331320010202111-1101320110232221-1222102310300230-3123220221113213): complete subsection reference.

- [use_system_defaults](resources--workload--reference--group-010.md#canonical-2222002321020303-1031121311231212-3123122030010220-0230112323220000-1200102022001220-3002100212022330-3331312123031020-2101231133103232): complete subsection reference.

<a id="canonical-3002233103302112-1231012300211200-3301330231223010-2121220130112302-0230202020131233-3131222231022113-1101233013331323-2122023302201323"></a>

## Next pages — tls_certificates / 012311030300 / 6

- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_certificates.custom_hash_algorithms](resources--workload--reference--group-010.md#canonical-1100332230200022-1320301123211133-3301233012201000-2212312020330103-1313303313301322-2011022030012011-3021030200200130-2121002333332023)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_certificates.disable_ocsp_stapling](resources--workload--reference--group-010.md#canonical-1120200202033020-3301221130122030-1333020301210113-2112321001230103-1131212222223230-2211110200221131-2121200222211120-3011112002022220)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_certificates.private_key](resources--workload--reference--group-010.md#canonical-0000103021020013-2333222020321111-2221322222122222-1012301222211232-3331320010202111-1101320110232221-1222102310300230-3123220221113213)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_certificates.use_system_defaults](resources--workload--reference--group-010.md#canonical-2222002321020303-1031121311231212-3123122030010220-0230112323220000-1200102022001220-3002100212022330-3331312123031020-2101231133103232)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-010.md#canonical-3230132323111312-0132020110123203-0222302101120231-0202012321012303-1213310020220000-0101203211110131-1303320330302033-3001313221002112)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-1100332230200022-1320301123211133-3301233012201000-2212312020330103-1313303313301322-2011022030012011-3021030200200130-2121002333332023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1130321020000202-3200233310331000-2220110322221000-1001010003232030-2102233111002233-0123030111000322-1220203130120213-0212000021231002"></a>

## service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_certificates.custom_hash_algorithms — custom_hash_algorithms / 232012332032 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-008.md#canonical-3210102113031103-2332201101220100-1103122221002212-3002101012001000-1332031130002011-0021231133303122-2121222212212213-1123320220010203)
- [service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-008.md#canonical-2311132233010120-0223230230122210-0332200203312210-3000132111100333-1133101301233010-3211310033122333-2323002110210000-1300012002110122)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](resources--workload--reference--group-008.md#canonical-1202020123020333-2332300320113330-1321112113331202-2112220001121101-3022001013310321-0020100001212100-0032232320012323-1311122212310031)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](resources--workload--reference--group-009.md#canonical-3212310132100213-2000301220120030-0112323230310220-3200032133332210-3223310210320320-2331013310222010-2010102133130322-2111123310321223)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-010.md#canonical-3230132323111312-0132020110123203-0222302101120231-0202012321012303-1213310020220000-0101203211110131-1303320330302033-3001313221002112)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_certificates](resources--workload--reference--group-010.md#canonical-0220313011200113-2210220130122300-1020333311303231-0221222302220012-3120013102133220-2132302001202230-2322133133130313-1313320213312221)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_certificates.custom_hash_algorithms

<a id="canonical-1201100322031311-1210031320031130-2311210121003131-2102020203201200-0012232131320222-1032231212131200-2133123212102023-1111210230313121"></a>

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

<a id="canonical-3030322023311002-0303322033120031-3010032321000312-1332230113132210-1333313320010113-2312230113210033-0311332023101210-1021233310120002"></a>

## Direct properties — custom_hash_algorithms / 232012332032 / 3

<a id="canonical-0230312223102321-0200020012102313-1330123232303210-2101130210023032-2132311302221120-1300201112200012-3202011020000331-0131101221000300"></a>

<a id="canonical-1000212221002020-1010102313322101-3320310232313130-2223023222002330-3020322011200313-3120322211232302-0031213210201331-0123103020322101"></a>

## hash_algorithms property — custom_hash_algorithms / 232012332032 / 4

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-1023130010210012-3233010001031021-2131012221323133-2123311320020202-3312031023102031-2113323232022023-2021202020202002-3231112311000000"></a>

## Next pages — custom_hash_algorithms / 232012332032 / 5

- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_certificates](resources--workload--reference--group-010.md#canonical-0220313011200113-2210220130122300-1020333311303231-0221222302220012-3120013102133220-2132302001202230-2322133133130313-1313320213312221)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-1120200202033020-3301221130122030-1333020301210113-2112321001230103-1131212222223230-2211110200221131-2121200222211120-3011112002022220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0033322222311222-1233220101111113-0110022323102022-0301100002002231-3230102320210003-3331013013011002-0033012103112230-3212310330013101"></a>

## service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_certificates.disable_ocsp_stapling — disable_ocsp_stapling / 133230010321 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-008.md#canonical-3210102113031103-2332201101220100-1103122221002212-3002101012001000-1332031130002011-0021231133303122-2121222212212213-1123320220010203)
- [service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-008.md#canonical-2311132233010120-0223230230122210-0332200203312210-3000132111100333-1133101301233010-3211310033122333-2323002110210000-1300012002110122)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](resources--workload--reference--group-008.md#canonical-1202020123020333-2332300320113330-1321112113331202-2112220001121101-3022001013310321-0020100001212100-0032232320012323-1311122212310031)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](resources--workload--reference--group-009.md#canonical-3212310132100213-2000301220120030-0112323230310220-3200032133332210-3223310210320320-2331013310222010-2010102133130322-2111123310321223)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-010.md#canonical-3230132323111312-0132020110123203-0222302101120231-0202012321012303-1213310020220000-0101203211110131-1303320330302033-3001313221002112)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_certificates](resources--workload--reference--group-010.md#canonical-0220313011200113-2210220130122300-1020333311303231-0221222302220012-3120013102133220-2132302001202230-2322133133130313-1313320213312221)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_certificates.disable_ocsp_stapling

<a id="canonical-1030220020203112-0311031231230022-1302213103223331-0001313221301112-0112021111302031-1211110320220031-1223331302210211-0232032221030011"></a>

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

<a id="canonical-3221120010030133-2221300312112131-3121202022301321-0203123222111313-2112102110032022-3031010013232131-0232030133012233-1003213223112233"></a>

## Direct properties — disable_ocsp_stapling / 133230010321 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1112222332211000-0110010323103130-3022010033002021-2211103101131132-1110201031313023-3033330313331320-1032332102313300-1300102111013021"></a>

## Next pages — disable_ocsp_stapling / 133230010321 / 4

- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_certificates](resources--workload--reference--group-010.md#canonical-0220313011200113-2210220130122300-1020333311303231-0221222302220012-3120013102133220-2132302001202230-2322133133130313-1313320213312221)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-0000103021020013-2333222020321111-2221322222122222-1012301222211232-3331320010202111-1101320110232221-1222102310300230-3123220221113213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3112233031211311-0001311312303111-3133332002131310-1023003022311112-0301122331230211-3222103323023121-1221021030330310-1333220202333113"></a>

## service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_certificates.private_key — private_key / 031201033033 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-008.md#canonical-3210102113031103-2332201101220100-1103122221002212-3002101012001000-1332031130002011-0021231133303122-2121222212212213-1123320220010203)
- [service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-008.md#canonical-2311132233010120-0223230230122210-0332200203312210-3000132111100333-1133101301233010-3211310033122333-2323002110210000-1300012002110122)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](resources--workload--reference--group-008.md#canonical-1202020123020333-2332300320113330-1321112113331202-2112220001121101-3022001013310321-0020100001212100-0032232320012323-1311122212310031)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](resources--workload--reference--group-009.md#canonical-3212310132100213-2000301220120030-0112323230310220-3200032133332210-3223310210320320-2331013310222010-2010102133130322-2111123310321223)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-010.md#canonical-3230132323111312-0132020110123203-0222302101120231-0202012321012303-1213310020220000-0101203211110131-1303320330302033-3001313221002112)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_certificates](resources--workload--reference--group-010.md#canonical-0220313011200113-2210220130122300-1020333311303231-0221222302220012-3120013102133220-2132302001202230-2322133133130313-1313320213312221)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_certificates.private_key

<a id="canonical-2113110101032232-2113221003212123-3132210101331223-1311302301300132-2302010020330122-1010330132113003-1332321311222100-2030003311132013"></a>

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

<a id="canonical-1011330010320132-2101100331000223-0221330022300230-3323101002113110-0020113120101311-3132031010203220-1010231111312333-0033332030202003"></a>

## Direct properties — private_key / 031201033033 / 3

- [blindfold_secret_info](resources--workload--reference--group-010.md#canonical-3220011020230211-2020302110311231-2301203030223111-1003113120302033-2030333022001033-0110001023230212-3100220220312002-1330303033123010): complete subsection reference.

- [clear_secret_info](resources--workload--reference--group-010.md#canonical-3233131032002012-1023332003300231-2203300311232200-1002033210211020-3013320211001002-3112231303133232-2033211123101332-2320202022010120): complete subsection reference.

<a id="canonical-3031233332023111-2103221031220332-2031322030033230-3332311032220221-1033123120100003-2330120220133200-2120321211103112-2231213011203210"></a>

## Next pages — private_key / 031201033033 / 4

- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_certificates.private_key.blindfold_secret_info](resources--workload--reference--group-010.md#canonical-3220011020230211-2020302110311231-2301203030223111-1003113120302033-2030333022001033-0110001023230212-3100220220312002-1330303033123010)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_certificates.private_key.clear_secret_info](resources--workload--reference--group-010.md#canonical-3233131032002012-1023332003300231-2203300311232200-1002033210211020-3013320211001002-3112231303133232-2033211123101332-2320202022010120)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_certificates](resources--workload--reference--group-010.md#canonical-0220313011200113-2210220130122300-1020333311303231-0221222302220012-3120013102133220-2132302001202230-2322133133130313-1313320213312221)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-3220011020230211-2020302110311231-2301203030223111-1003113120302033-2030333022001033-0110001023230212-3100220220312002-1330303033123010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1011103032320213-2230203300303212-1133310213201313-3310323102100313-2331321102223311-3111111222310301-2110013102212102-2221102130331300"></a>

## service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_certificates.private_key.blindfold_secret_info — blindfold_secret_info / 312123110123 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-008.md#canonical-3210102113031103-2332201101220100-1103122221002212-3002101012001000-1332031130002011-0021231133303122-2121222212212213-1123320220010203)
- [service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-008.md#canonical-2311132233010120-0223230230122210-0332200203312210-3000132111100333-1133101301233010-3211310033122333-2323002110210000-1300012002110122)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](resources--workload--reference--group-008.md#canonical-1202020123020333-2332300320113330-1321112113331202-2112220001121101-3022001013310321-0020100001212100-0032232320012323-1311122212310031)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](resources--workload--reference--group-009.md#canonical-3212310132100213-2000301220120030-0112323230310220-3200032133332210-3223310210320320-2331013310222010-2010102133130322-2111123310321223)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-010.md#canonical-3230132323111312-0132020110123203-0222302101120231-0202012321012303-1213310020220000-0101203211110131-1303320330302033-3001313221002112)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_certificates](resources--workload--reference--group-010.md#canonical-0220313011200113-2210220130122300-1020333311303231-0221222302220012-3120013102133220-2132302001202230-2322133133130313-1313320213312221)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_certificates.private_key](resources--workload--reference--group-010.md#canonical-0000103021020013-2333222020321111-2221322222122222-1012301222211232-3331320010202111-1101320110232221-1222102310300230-3123220221113213)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_certificates.private_key.blindfold_secret_info

<a id="canonical-0230332201020213-1202020121132212-0123032331322123-2001302031201123-0303033211333221-3311233212330310-0032301301112310-0010210300312031"></a>

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

<a id="canonical-0112223133110330-1133203012100322-2022233212303311-3200222011303212-3101231321323020-3113320221123001-0222330330321130-3332300321122103"></a>

## Direct properties — blindfold_secret_info / 312123110123 / 3

<a id="canonical-1201232123322310-1020131323012130-1323323122312210-0203233331133312-3212323211123221-0200333120303222-1123321220111032-0001323013221323"></a>

<a id="canonical-3011310332123200-2211002122232033-0330331222313211-3023103220011100-3231313121213210-2011001131303010-2113033003312331-2123323201013031"></a>

## decryption_provider property — blindfold_secret_info / 312123110123 / 4

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-2213220003130110-2313200223111100-1122010122032330-2332200110000332-3320201020100012-0110300130202320-3022231220011021-3323322220111112"></a>

<a id="canonical-2203002020020121-3103131122030313-0120022331313333-2223132012222221-0012332212103110-1132201130212302-2023113231101021-2323322003202030"></a>

## location property — blindfold_secret_info / 312123110123 / 5

Type: `"string"`. Optional, Sensitive.

Location is the uri\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Upstream description:

Location is the uri\_ref. It could be in URL format for string:/// Or it could be a path if the
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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-2013110012233133-0300122021330213-1020131100122002-1320020303122133-0022230330001020-2203030031211312-3303001331301032-2323002021230131"></a>

<a id="canonical-1313310132330210-2111002013301102-1131313110110111-2321302232100003-3101332100001203-0131230130123213-3333323231210201-3300323111302221"></a>

## store_provider property — blindfold_secret_info / 312123110123 / 6

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-2333223312113220-0233012023310123-1110110120233110-0002022303220000-0112110230003300-0133311311001003-0012232113321101-1323303333022110"></a>

## Next pages — blindfold_secret_info / 312123110123 / 7

- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_certificates.private_key](resources--workload--reference--group-010.md#canonical-0000103021020013-2333222020321111-2221322222122222-1012301222211232-3331320010202111-1101320110232221-1222102310300230-3123220221113213)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-3233131032002012-1023332003300231-2203300311232200-1002033210211020-3013320211001002-3112231303133232-2033211123101332-2320202022010120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3321020000223333-1230220030113122-3201331202013231-0211002021300200-1300131330323220-1223220311332313-1120132313211312-2113233313230103"></a>

## service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_certificates.private_key.clear_secret_info — clear_secret_info / 031201313010 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-008.md#canonical-3210102113031103-2332201101220100-1103122221002212-3002101012001000-1332031130002011-0021231133303122-2121222212212213-1123320220010203)
- [service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-008.md#canonical-2311132233010120-0223230230122210-0332200203312210-3000132111100333-1133101301233010-3211310033122333-2323002110210000-1300012002110122)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](resources--workload--reference--group-008.md#canonical-1202020123020333-2332300320113330-1321112113331202-2112220001121101-3022001013310321-0020100001212100-0032232320012323-1311122212310031)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](resources--workload--reference--group-009.md#canonical-3212310132100213-2000301220120030-0112323230310220-3200032133332210-3223310210320320-2331013310222010-2010102133130322-2111123310321223)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-010.md#canonical-3230132323111312-0132020110123203-0222302101120231-0202012321012303-1213310020220000-0101203211110131-1303320330302033-3001313221002112)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_certificates](resources--workload--reference--group-010.md#canonical-0220313011200113-2210220130122300-1020333311303231-0221222302220012-3120013102133220-2132302001202230-2322133133130313-1313320213312221)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_certificates.private_key](resources--workload--reference--group-010.md#canonical-0000103021020013-2333222020321111-2221322222122222-1012301222211232-3331320010202111-1101320110232221-1222102310300230-3123220221113213)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_certificates.private_key.clear_secret_info

<a id="canonical-0022223303312122-1201000322010330-1000333312112310-3233230113320233-3222220010230201-3312000301330230-1321003131333031-1003121321021212"></a>

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

<a id="canonical-3123322223312203-1021030220001012-0001123211301201-1121330030321012-3030211200002023-3321123203320131-1211030201120212-1233210303030131"></a>

## Direct properties — clear_secret_info / 031201313010 / 3

<a id="canonical-3301320302110203-3032323200013001-1121032012303003-0301211332201003-0031120020332000-2230212123012231-2001310323310301-3100212001130230"></a>

<a id="canonical-3030113311230330-1331313330000001-3221003032102313-0300201000112203-3210212120032101-0130102102312212-1120031320223322-0031010313211232"></a>

## provider_ref property — clear_secret_info / 031201313010 / 4

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-2033232020213113-3332123103201231-3112211101031312-0332301330121213-1111313132133203-1220123320232010-0230223330323231-3122322213322230"></a>

<a id="canonical-2132102322031213-0202030331210010-3331310031323103-1210011003121231-2231133200320321-0122112111031300-2231203113203322-2203222103031101"></a>

## URL property — clear_secret_info / 031201313010 / 5

Type: `"string"`. Optional, Sensitive.

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded Base64 format. When asked for this secret, caller will GET Secret bytes after
Base64 decoding.

Upstream description:

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded Base64 format. When asked for this secret, caller will GET Secret bytes after
Base64 decoding.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-1002312323301003-3332112222331210-1202313323131212-3122011223212303-3123120002111003-3200200112231320-3133133011033333-2012003331111131"></a>

## Next pages — clear_secret_info / 031201313010 / 6

- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_certificates.private_key](resources--workload--reference--group-010.md#canonical-0000103021020013-2333222020321111-2221322222122222-1012301222211232-3331320010202111-1101320110232221-1222102310300230-3123220221113213)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-2222002321020303-1031121311231212-3123122030010220-0230112323220000-1200102022001220-3002100212022330-3331312123031020-2101231133103232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0131221223303300-3021223133220003-0301031231332312-2212102121222022-3232030101010331-1212010003333031-0203321332201111-2331131102132112"></a>

## service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_certificates.use_system_defaults — use_system_defaults / 122123031200 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-008.md#canonical-3210102113031103-2332201101220100-1103122221002212-3002101012001000-1332031130002011-0021231133303122-2121222212212213-1123320220010203)
- [service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-008.md#canonical-2311132233010120-0223230230122210-0332200203312210-3000132111100333-1133101301233010-3211310033122333-2323002110210000-1300012002110122)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](resources--workload--reference--group-008.md#canonical-1202020123020333-2332300320113330-1321112113331202-2112220001121101-3022001013310321-0020100001212100-0032232320012323-1311122212310031)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](resources--workload--reference--group-009.md#canonical-3212310132100213-2000301220120030-0112323230310220-3200032133332210-3223310210320320-2331013310222010-2010102133130322-2111123310321223)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-010.md#canonical-3230132323111312-0132020110123203-0222302101120231-0202012321012303-1213310020220000-0101203211110131-1303320330302033-3001313221002112)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_certificates](resources--workload--reference--group-010.md#canonical-0220313011200113-2210220130122300-1020333311303231-0221222302220012-3120013102133220-2132302001202230-2322133133130313-1313320213312221)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_certificates.use_system_defaults

<a id="canonical-1000220113001030-0031103113122231-2110003022203301-3222221222302210-2101133210332011-3330113010321002-1101230212102000-1222312202230130"></a>

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

<a id="canonical-3220133213001220-1011023210012323-0010231321023211-0013102333331012-0033113100300323-3233301233103212-0013112122332101-2130311233021111"></a>

## Direct properties — use_system_defaults / 122123031200 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1012200213201022-2222200300021131-2233012003112223-0213033231122331-3232120122130111-3121212112222300-1112301013332322-3320001103202301"></a>

## Next pages — use_system_defaults / 122123031200 / 4

- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_certificates](resources--workload--reference--group-010.md#canonical-0220313011200113-2210220130122300-1020333311303231-0221222302220012-3120013102133220-2132302001202230-2322133133130313-1313320213312221)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-1233203302211212-3300120033202101-0200330001312321-1222133023031223-1233313002123211-3203212322123012-1131031112003021-1331232103023221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1001203233221300-2321321002223122-3211002031201202-1022002330220201-0311101123211321-3120012202131120-0310001213012130-0013022313101120"></a>

## service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_config — tls_config / 030231132310 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-008.md#canonical-3210102113031103-2332201101220100-1103122221002212-3002101012001000-1332031130002011-0021231133303122-2121222212212213-1123320220010203)
- [service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-008.md#canonical-2311132233010120-0223230230122210-0332200203312210-3000132111100333-1133101301233010-3211310033122333-2323002110210000-1300012002110122)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](resources--workload--reference--group-008.md#canonical-1202020123020333-2332300320113330-1321112113331202-2112220001121101-3022001013310321-0020100001212100-0032232320012323-1311122212310031)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](resources--workload--reference--group-009.md#canonical-3212310132100213-2000301220120030-0112323230310220-3200032133332210-3223310210320320-2331013310222010-2010102133130322-2111123310321223)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-010.md#canonical-3230132323111312-0132020110123203-0222302101120231-0202012321012303-1213310020220000-0101203211110131-1303320330302033-3001313221002112)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_config

<a id="canonical-1021112010022003-1232212023220022-1131231223030003-3001312101202103-0020312131033210-1201003023023020-2311332101003112-0213010211311223"></a>

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

<a id="canonical-0013210030000010-2323130302203210-0310110130330120-1102100020202001-3233233300000222-1100312231011233-3320101000122033-1310132012220330"></a>

## Direct properties — tls_config / 030231132310 / 3

- [custom_security](resources--workload--reference--group-010.md#canonical-1231032221230000-1010100200221203-1121220011012313-3332330220020233-1100222130021303-0203321102313101-0211023223331023-0011323212302011): complete subsection reference.

- [default_security](resources--workload--reference--group-010.md#canonical-2212231013132322-1113323133200030-3133300033021130-0100211301333112-3133022222320232-3120322131131110-3113011032331223-3312133132103033): complete subsection reference.

- [low_security](resources--workload--reference--group-010.md#canonical-3230331013322123-1233332021013212-2301010301100300-2130203133232103-2300113022202231-3013013233020330-0323212301221010-3003032120203312): complete subsection reference.

- [medium_security](resources--workload--reference--group-010.md#canonical-3221312221332112-3133132000110333-0021212030322333-1222322221303102-2111122201230103-3320102333230330-2101323311022232-0032302321133232): complete subsection reference.

<a id="canonical-1132200033120032-0120213223110233-2212221320202203-1212003003123032-0311033102323120-0301313221311103-3031303301210231-0133232111202331"></a>

## Next pages — tls_config / 030231132310 / 4

- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_config.custom_security](resources--workload--reference--group-010.md#canonical-1231032221230000-1010100200221203-1121220011012313-3332330220020233-1100222130021303-0203321102313101-0211023223331023-0011323212302011)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_config.default_security](resources--workload--reference--group-010.md#canonical-2212231013132322-1113323133200030-3133300033021130-0100211301333112-3133022222320232-3120322131131110-3113011032331223-3312133132103033)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_config.low_security](resources--workload--reference--group-010.md#canonical-3230331013322123-1233332021013212-2301010301100300-2130203133232103-2300113022202231-3013013233020330-0323212301221010-3003032120203312)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_config.medium_security](resources--workload--reference--group-010.md#canonical-3221312221332112-3133132000110333-0021212030322333-1222322221303102-2111122201230103-3320102333230330-2101323311022232-0032302321133232)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-010.md#canonical-3230132323111312-0132020110123203-0222302101120231-0202012321012303-1213310020220000-0101203211110131-1303320330302033-3001313221002112)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-1231032221230000-1010100200221203-1121220011012313-3332330220020233-1100222130021303-0203321102313101-0211023223331023-0011323212302011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1332033010032010-1223030230102022-2223100313021220-3313213311103023-2232313033000123-2131123033323020-1020012213332303-2211220110032001"></a>

## service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_config.custom_security — custom_security / 232011303122 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-008.md#canonical-3210102113031103-2332201101220100-1103122221002212-3002101012001000-1332031130002011-0021231133303122-2121222212212213-1123320220010203)
- [service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-008.md#canonical-2311132233010120-0223230230122210-0332200203312210-3000132111100333-1133101301233010-3211310033122333-2323002110210000-1300012002110122)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](resources--workload--reference--group-008.md#canonical-1202020123020333-2332300320113330-1321112113331202-2112220001121101-3022001013310321-0020100001212100-0032232320012323-1311122212310031)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](resources--workload--reference--group-009.md#canonical-3212310132100213-2000301220120030-0112323230310220-3200032133332210-3223310210320320-2331013310222010-2010102133130322-2111123310321223)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-010.md#canonical-3230132323111312-0132020110123203-0222302101120231-0202012321012303-1213310020220000-0101203211110131-1303320330302033-3001313221002112)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_config](resources--workload--reference--group-010.md#canonical-1233203302211212-3300120033202101-0200330001312321-1222133023031223-1233313002123211-3203212322123012-1131031112003021-1331232103023221)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_config.custom_security

<a id="canonical-0001022220131201-0012010223003002-0221120022222010-0123320300102022-3333213113311320-2300322103310123-1111331333032030-3033012030003302"></a>

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

<a id="canonical-3303300121113103-2022131022303001-3231321111333111-3212331202023231-2230032312030333-2311233000223203-2132011031020130-3123033111202302"></a>

## Direct properties — custom_security / 232011303122 / 3

<a id="canonical-1220130030100132-2211000312211001-1022020203120303-2303331110021232-1130232323231123-1023131032100122-3303011001120332-1031322313223322"></a>

<a id="canonical-0211032012300033-2331221231303131-0002023230321033-2031122112130201-0020022122310222-1101100021200322-2022232010121003-0223322003233310"></a>

## cipher_suites property — custom_security / 232011303122 / 4

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-0031122200122322-1113230120223003-2012010012000223-0322300032010113-0312201130122301-2201311123100002-0111112201233332-2123123112111122"></a>

<a id="canonical-3303320033112201-0211130212313210-2031232213202320-0232301033333100-2312012032330102-2020213122022302-3030321023000133-2110210213303212"></a>

## max_version property — custom_security / 232011303122 / 5

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

<a id="canonical-0030220231123331-2301102210220311-3001033321122131-2131300100302211-2323133020322010-1130320110131120-0301332332311303-2331011310120201"></a>

<a id="canonical-1223030303330312-3222112202303201-0232313120303030-1210322131103131-0132311133102333-2210013011230132-2300330213322311-0101311101003021"></a>

## min_version property — custom_security / 232011303122 / 6

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

<a id="canonical-2003130211322023-2010023022303020-2102100223332021-1120121320312221-0010211321010013-1203113213021231-1103322010103033-0301322233101330"></a>

## Next pages — custom_security / 232011303122 / 7

- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_config](resources--workload--reference--group-010.md#canonical-1233203302211212-3300120033202101-0200330001312321-1222133023031223-1233313002123211-3203212322123012-1131031112003021-1331232103023221)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-2212231013132322-1113323133200030-3133300033021130-0100211301333112-3133022222320232-3120322131131110-3113011032331223-3312133132103033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1101122012220301-2331123210020033-3032113232320122-0331220233232021-0323103201222022-0020331202011332-1220122221032332-2000201033020220"></a>

## service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_config.default_security — default_security / 321301211020 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-008.md#canonical-3210102113031103-2332201101220100-1103122221002212-3002101012001000-1332031130002011-0021231133303122-2121222212212213-1123320220010203)
- [service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-008.md#canonical-2311132233010120-0223230230122210-0332200203312210-3000132111100333-1133101301233010-3211310033122333-2323002110210000-1300012002110122)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](resources--workload--reference--group-008.md#canonical-1202020123020333-2332300320113330-1321112113331202-2112220001121101-3022001013310321-0020100001212100-0032232320012323-1311122212310031)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](resources--workload--reference--group-009.md#canonical-3212310132100213-2000301220120030-0112323230310220-3200032133332210-3223310210320320-2331013310222010-2010102133130322-2111123310321223)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-010.md#canonical-3230132323111312-0132020110123203-0222302101120231-0202012321012303-1213310020220000-0101203211110131-1303320330302033-3001313221002112)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_config](resources--workload--reference--group-010.md#canonical-1233203302211212-3300120033202101-0200330001312321-1222133023031223-1233313002123211-3203212322123012-1131031112003021-1331232103023221)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_config.default_security

<a id="canonical-2223030220201032-3013130030311211-2113222113102022-0222001001021223-2330130220300120-3313303131230030-3200232010321112-2020032122302320"></a>

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

<a id="canonical-1220301213213121-1010321020330223-2032031110131213-3013103301033212-2333303133120110-3130030122130231-2200323112223300-2022002133020102"></a>

## Direct properties — default_security / 321301211020 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0032321103023020-0003322102302303-1301210100222222-2000202032222221-2102330220011303-1333113000311012-2300222333323321-3201003103120120"></a>

## Next pages — default_security / 321301211020 / 4

- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_config](resources--workload--reference--group-010.md#canonical-1233203302211212-3300120033202101-0200330001312321-1222133023031223-1233313002123211-3203212322123012-1131031112003021-1331232103023221)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-3230331013322123-1233332021013212-2301010301100300-2130203133232103-2300113022202231-3013013233020330-0323212301221010-3003032120203312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0310111002033301-3102131021310202-1130322023203202-1101113132132211-0330213120020100-0322333203102012-1312133310133112-0121303122230131"></a>

## service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_config.low_security — low_security / 322132102131 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-008.md#canonical-3210102113031103-2332201101220100-1103122221002212-3002101012001000-1332031130002011-0021231133303122-2121222212212213-1123320220010203)
- [service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-008.md#canonical-2311132233010120-0223230230122210-0332200203312210-3000132111100333-1133101301233010-3211310033122333-2323002110210000-1300012002110122)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](resources--workload--reference--group-008.md#canonical-1202020123020333-2332300320113330-1321112113331202-2112220001121101-3022001013310321-0020100001212100-0032232320012323-1311122212310031)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](resources--workload--reference--group-009.md#canonical-3212310132100213-2000301220120030-0112323230310220-3200032133332210-3223310210320320-2331013310222010-2010102133130322-2111123310321223)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-010.md#canonical-3230132323111312-0132020110123203-0222302101120231-0202012321012303-1213310020220000-0101203211110131-1303320330302033-3001313221002112)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_config](resources--workload--reference--group-010.md#canonical-1233203302211212-3300120033202101-0200330001312321-1222133023031223-1233313002123211-3203212322123012-1131031112003021-1331232103023221)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_config.low_security

<a id="canonical-2212131332110032-2020021232220111-2032120221213200-1020103203031131-0022311001332201-2113113111130202-3230212233210230-3333223313022301"></a>

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

<a id="canonical-0113213232113300-0302230002312123-2030230102110312-3302133111203110-2300120103301201-1102311301211322-2112211313103123-3311332331232233"></a>

## Direct properties — low_security / 322132102131 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0123201010003010-1212011310132310-3300021001130233-1303131203232022-2333133010102333-1313011122313230-1121323202330331-0011003201003011"></a>

## Next pages — low_security / 322132102131 / 4

- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_config](resources--workload--reference--group-010.md#canonical-1233203302211212-3300120033202101-0200330001312321-1222133023031223-1233313002123211-3203212322123012-1131031112003021-1331232103023221)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-3221312221332112-3133132000110333-0021212030322333-1222322221303102-2111122201230103-3320102333230330-2101323311022232-0032302321133232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2012020120100021-1303220320333110-3121302332102211-1132231030211012-2001022010333333-1303003312331303-3100220023210232-2003123022102303"></a>

## service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_config.medium_security — medium_security / 303311230132 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-008.md#canonical-3210102113031103-2332201101220100-1103122221002212-3002101012001000-1332031130002011-0021231133303122-2121222212212213-1123320220010203)
- [service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-008.md#canonical-2311132233010120-0223230230122210-0332200203312210-3000132111100333-1133101301233010-3211310033122333-2323002110210000-1300012002110122)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](resources--workload--reference--group-008.md#canonical-1202020123020333-2332300320113330-1321112113331202-2112220001121101-3022001013310321-0020100001212100-0032232320012323-1311122212310031)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](resources--workload--reference--group-009.md#canonical-3212310132100213-2000301220120030-0112323230310220-3200032133332210-3223310210320320-2331013310222010-2010102133130322-2111123310321223)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-010.md#canonical-3230132323111312-0132020110123203-0222302101120231-0202012321012303-1213310020220000-0101203211110131-1303320330302033-3001313221002112)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_config](resources--workload--reference--group-010.md#canonical-1233203302211212-3300120033202101-0200330001312321-1222133023031223-1233313002123211-3203212322123012-1131031112003021-1331232103023221)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_config.medium_security

<a id="canonical-1332123033032010-1110111113002120-2121032330331233-0231330000322133-3011120200122021-0012301213201331-0122331203100210-1003120123132312"></a>

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

<a id="canonical-2022333332332330-1202303303312000-2011302111123223-1020012030011103-1131011211012001-2032201321021110-0323022231302210-3132313222230101"></a>

## Direct properties — medium_security / 303311230132 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2020021222131123-2323213221222133-3000001102320112-0122211031101330-0213131112233301-1121210011200013-0300302110022133-1013321012211303"></a>

## Next pages — medium_security / 303311230132 / 4

- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_config](resources--workload--reference--group-010.md#canonical-1233203302211212-3300120033202101-0200330001312321-1222133023031223-1233313002123211-3203212322123012-1131031112003021-1331232103023221)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-3310220202003120-0020232310001311-1230011321302031-3002200220020220-1002102031303233-0322332031301112-0031333023233103-0133102211311201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0321302212210203-1311333302222202-1313221013001101-2232013232302232-1221311201301321-2013021033030122-2322001211210211-1113022202223222"></a>

## service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.use_mtls — use_mtls / 110231132032 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-008.md#canonical-3210102113031103-2332201101220100-1103122221002212-3002101012001000-1332031130002011-0021231133303122-2121222212212213-1123320220010203)
- [service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-008.md#canonical-2311132233010120-0223230230122210-0332200203312210-3000132111100333-1133101301233010-3211310033122333-2323002110210000-1300012002110122)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](resources--workload--reference--group-008.md#canonical-1202020123020333-2332300320113330-1321112113331202-2112220001121101-3022001013310321-0020100001212100-0032232320012323-1311122212310031)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](resources--workload--reference--group-009.md#canonical-3212310132100213-2000301220120030-0112323230310220-3200032133332210-3223310210320320-2331013310222010-2010102133130322-2111123310321223)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-010.md#canonical-3230132323111312-0132020110123203-0222302101120231-0202012321012303-1213310020220000-0101203211110131-1303320330302033-3001313221002112)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.use_mtls

<a id="canonical-2120111023132302-1210312111201023-1030011013200013-2033120121331120-0021231020003133-2022012201133213-3121202030132031-2101000020232022"></a>

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

<a id="canonical-1111001303302331-0221020201113320-0303031221312211-0302202102321030-2133020220331011-1233220133300332-0313103222120010-3113323003030202"></a>

## Direct properties — use_mtls / 110231132032 / 3

<a id="canonical-3103233303331132-3313212021122001-2120312023233322-2233212330310132-1222001011322332-3230211232120112-1022100200030211-3210023012321102"></a>

<a id="canonical-1021110110132210-0103330310011131-0302212102032312-3322000200123213-2100323223120231-1001201112110121-3012102133013102-0332021321300011"></a>

## client_certificate_optional property — use_mtls / 110231132032 / 4

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

- [crl](resources--workload--reference--group-010.md#canonical-2001333030330320-2023002103031001-3002232311311032-3022320130133223-2220003002311012-2200213221110010-2011130013200033-3312301120331202): complete subsection reference.

- [no_crl](resources--workload--reference--group-010.md#canonical-1311102021032031-3333223133333320-1000312233313312-1201011122102200-2120123032230012-2121221321023031-2121231012212103-3112303201103323): complete subsection reference.

- [trusted_ca](resources--workload--reference--group-010.md#canonical-1320023331013030-2310123101313321-2010323333101333-1131332221202102-1130032102133222-2111333301023022-0032001012303232-0102013111132220): complete subsection reference.

<a id="canonical-3302021222231310-0012120021310322-2231202120121221-0130002022333320-2021331213313133-1333002303230003-1330032010330201-3111023010021133"></a>

<a id="canonical-3301132201332200-0133010310023010-3133113223032322-2323230210022103-2030312002011112-1220320010221001-2002332223001322-3023010110021110"></a>

## trusted_ca_url property — use_mtls / 110231132032 / 5

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

- [xfcc_disabled](resources--workload--reference--group-010.md#canonical-1123230303002221-0101211213322212-1222302011111321-1110212023133202-0122011213310112-0310310110213111-2023030303012010-1310313101030302): complete subsection reference.

- [xfcc_options](resources--workload--reference--group-010.md#canonical-1022001030312122-0110021131331203-1321111312211223-0231321110200201-2232010002033222-1321002220133122-2012110202313210-2130032010303332): complete subsection reference.

<a id="canonical-2302001321031021-0302302112133011-3222122102012031-1031212321022222-0330001020320230-1130111233132131-1100010232030130-2003300010003300"></a>

## Next pages — use_mtls / 110231132032 / 6

- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.use_mtls.crl](resources--workload--reference--group-010.md#canonical-2001333030330320-2023002103031001-3002232311311032-3022320130133223-2220003002311012-2200213221110010-2011130013200033-3312301120331202)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.use_mtls.no_crl](resources--workload--reference--group-010.md#canonical-1311102021032031-3333223133333320-1000312233313312-1201011122102200-2120123032230012-2121221321023031-2121231012212103-3112303201103323)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.use_mtls.trusted_ca](resources--workload--reference--group-010.md#canonical-1320023331013030-2310123101313321-2010323333101333-1131332221202102-1130032102133222-2111333301023022-0032001012303232-0102013111132220)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.use_mtls.xfcc_disabled](resources--workload--reference--group-010.md#canonical-1123230303002221-0101211213322212-1222302011111321-1110212023133202-0122011213310112-0310310110213111-2023030303012010-1310313101030302)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.use_mtls.xfcc_options](resources--workload--reference--group-010.md#canonical-1022001030312122-0110021131331203-1321111312211223-0231321110200201-2232010002033222-1321002220133122-2012110202313210-2130032010303332)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-010.md#canonical-3230132323111312-0132020110123203-0222302101120231-0202012321012303-1213310020220000-0101203211110131-1303320330302033-3001313221002112)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-2001333030330320-2023002103031001-3002232311311032-3022320130133223-2220003002311012-2200213221110010-2011130013200033-3312301120331202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0232133120222131-2222211001211133-2011230220131222-1220110303322100-3113031202110202-3010330322110333-3313023221010321-2022103130330230"></a>

## service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.use_mtls.crl — crl / 130011311111 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-008.md#canonical-3210102113031103-2332201101220100-1103122221002212-3002101012001000-1332031130002011-0021231133303122-2121222212212213-1123320220010203)
- [service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-008.md#canonical-2311132233010120-0223230230122210-0332200203312210-3000132111100333-1133101301233010-3211310033122333-2323002110210000-1300012002110122)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](resources--workload--reference--group-008.md#canonical-1202020123020333-2332300320113330-1321112113331202-2112220001121101-3022001013310321-0020100001212100-0032232320012323-1311122212310031)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](resources--workload--reference--group-009.md#canonical-3212310132100213-2000301220120030-0112323230310220-3200032133332210-3223310210320320-2331013310222010-2010102133130322-2111123310321223)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-010.md#canonical-3230132323111312-0132020110123203-0222302101120231-0202012321012303-1213310020220000-0101203211110131-1303320330302033-3001313221002112)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.use_mtls](resources--workload--reference--group-010.md#canonical-3310220202003120-0020232310001311-1230011321302031-3002200220020220-1002102031303233-0322332031301112-0031333023233103-0133102211311201)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.use_mtls.crl

<a id="canonical-2130211313312130-1210132332123120-0122213002122302-3031021020233210-0102033001112103-2113103031332330-2102000133101013-3020033001011101"></a>

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

<a id="canonical-2310233220231321-1110332011200312-1130003200020223-2123113121030033-3210030022132303-0012112102003132-3001023110013103-3011210021113002"></a>

## Direct properties — crl / 130011311111 / 3

<a id="canonical-0022102312301022-0203101333013233-0320020132320212-2100030112033000-2021000100303330-0100023102213030-0221212001012333-3333323101120222"></a>

<a id="canonical-1113110210010230-1303112000002031-3103311022202201-0120102312021003-0321303233321111-1010333000220011-0200300223300121-2320023031132103"></a>

## name property — crl / 130011311111 / 4

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-3322203013220313-0120313100003113-3222233102011012-2110013131100303-0012133222332101-0023323230211022-2311123202023103-0200033320000113"></a>

<a id="canonical-1130102022311301-2101320130300211-3232211101213220-0112003031121230-3232003301130031-3233223301313311-1323222031203201-0011312132312320"></a>

## namespace property — crl / 130011311111 / 5

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-3222222133021332-0222232322230332-0030001310233100-3231022000220032-0312321103123202-1210233120300110-0030212131302202-0130333302011010"></a>

<a id="canonical-0302010023102302-2311200211032101-0323103100203112-1122211011321111-2131101020311120-2310103130013020-0120111032221310-0120321310311130"></a>

## tenant property — crl / 130011311111 / 6

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-1313121132221321-2200111120312111-2133012120013111-2301332323330011-2202202201013233-0232301210303200-0231310032233132-1222121001313122"></a>

## Next pages — crl / 130011311111 / 7

- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.use_mtls](resources--workload--reference--group-010.md#canonical-3310220202003120-0020232310001311-1230011321302031-3002200220020220-1002102031303233-0322332031301112-0031333023233103-0133102211311201)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-1311102021032031-3333223133333320-1000312233313312-1201011122102200-2120123032230012-2121221321023031-2121231012212103-3112303201103323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0300201201332231-0132300213111110-1301022123312330-2231011301033302-1301102323001231-1112112133122321-0111212011322030-0230022101000010"></a>

## service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.use_mtls.no_crl — no_crl / 201123000110 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-008.md#canonical-3210102113031103-2332201101220100-1103122221002212-3002101012001000-1332031130002011-0021231133303122-2121222212212213-1123320220010203)
- [service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-008.md#canonical-2311132233010120-0223230230122210-0332200203312210-3000132111100333-1133101301233010-3211310033122333-2323002110210000-1300012002110122)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](resources--workload--reference--group-008.md#canonical-1202020123020333-2332300320113330-1321112113331202-2112220001121101-3022001013310321-0020100001212100-0032232320012323-1311122212310031)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](resources--workload--reference--group-009.md#canonical-3212310132100213-2000301220120030-0112323230310220-3200032133332210-3223310210320320-2331013310222010-2010102133130322-2111123310321223)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-010.md#canonical-3230132323111312-0132020110123203-0222302101120231-0202012321012303-1213310020220000-0101203211110131-1303320330302033-3001313221002112)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.use_mtls](resources--workload--reference--group-010.md#canonical-3310220202003120-0020232310001311-1230011321302031-3002200220020220-1002102031303233-0322332031301112-0031333023233103-0133102211311201)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.use_mtls.no_crl

<a id="canonical-2123102233120113-3211220313301001-1211211111130202-1032110131233333-2123300202023322-0220110020212130-2221323233002232-0311013320202032"></a>

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

<a id="canonical-1120220301100230-0110113300311203-1220332102033031-1301322132323031-0312000332132011-0012023133323332-2022220213020011-3000103002303100"></a>

## Direct properties — no_crl / 201123000110 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2223100322310333-1312111003022200-2201233233321131-1200202322322012-1001120210321130-2332223231210210-1121233013311201-0031102200233230"></a>

## Next pages — no_crl / 201123000110 / 4

- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.use_mtls](resources--workload--reference--group-010.md#canonical-3310220202003120-0020232310001311-1230011321302031-3002200220020220-1002102031303233-0322332031301112-0031333023233103-0133102211311201)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-1320023331013030-2310123101313321-2010323333101333-1131332221202102-1130032102133222-2111333301023022-0032001012303232-0102013111132220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3310122101120013-2102323003313100-3032331213220020-0201011120131100-2322031133012022-3300332332323221-3023313111011102-0233213320013133"></a>

## service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.use_mtls.trusted_ca — trusted_ca / 220003102132 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-008.md#canonical-3210102113031103-2332201101220100-1103122221002212-3002101012001000-1332031130002011-0021231133303122-2121222212212213-1123320220010203)
- [service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-008.md#canonical-2311132233010120-0223230230122210-0332200203312210-3000132111100333-1133101301233010-3211310033122333-2323002110210000-1300012002110122)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](resources--workload--reference--group-008.md#canonical-1202020123020333-2332300320113330-1321112113331202-2112220001121101-3022001013310321-0020100001212100-0032232320012323-1311122212310031)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](resources--workload--reference--group-009.md#canonical-3212310132100213-2000301220120030-0112323230310220-3200032133332210-3223310210320320-2331013310222010-2010102133130322-2111123310321223)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-010.md#canonical-3230132323111312-0132020110123203-0222302101120231-0202012321012303-1213310020220000-0101203211110131-1303320330302033-3001313221002112)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.use_mtls](resources--workload--reference--group-010.md#canonical-3310220202003120-0020232310001311-1230011321302031-3002200220020220-1002102031303233-0322332031301112-0031333023233103-0133102211311201)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.use_mtls.trusted_ca

<a id="canonical-2010321033203023-0122011110100121-1221203302303013-3102222030211032-3030202121203130-0122110101310123-3222230333023203-3331213210023300"></a>

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

<a id="canonical-0023312012012333-2323101002212131-0213300030200123-0200313001021302-3132103000032310-1221331321202031-3012331333232122-3112332222233211"></a>

## Direct properties — trusted_ca / 220003102132 / 3

<a id="canonical-3230313123330033-2123122201133011-0233321322211220-1320333013023100-2102111330103103-2310232112221100-3013013300220212-2002100002020130"></a>

<a id="canonical-2312211211301130-0322112211322030-1331030111322120-2311313030000102-2303013333212020-0131003313222212-1120211130132221-3221211022312121"></a>

## name property — trusted_ca / 220003102132 / 4

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-2301222132203210-3313223021320213-1333012323102103-2321110102103231-1220011010000011-2102232120202323-2211302320121130-2223301103311000"></a>

<a id="canonical-3102010333310032-0122102323103002-0110103320020223-1310010312002213-3020101300230232-3333230230323012-0030332203111313-0030312301103202"></a>

## namespace property — trusted_ca / 220003102132 / 5

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-1032123320321033-0200312331001112-2210300102020101-0103021003000132-0330023033121112-0131332331122102-1321233223133001-0203221212303122"></a>

<a id="canonical-1210012102201301-3102000223210321-3313232010202022-3200211311201100-0003213102030011-0101122001131031-3311131021223222-2320021021302321"></a>

## tenant property — trusted_ca / 220003102132 / 6

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-3111311033030101-1223301022010303-0211201223301123-3003231313022223-1201102323321331-1300303203200210-2022321112110321-3121133011230103"></a>

## Next pages — trusted_ca / 220003102132 / 7

- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.use_mtls](resources--workload--reference--group-010.md#canonical-3310220202003120-0020232310001311-1230011321302031-3002200220020220-1002102031303233-0322332031301112-0031333023233103-0133102211311201)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-1123230303002221-0101211213322212-1222302011111321-1110212023133202-0122011213310112-0310310110213111-2023030303012010-1310313101030302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1033321013230102-3212202313123010-3112101101032000-0321301322213333-1331323312323030-1100101121033221-1113321223330003-2120311300100121"></a>

## service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.use_mtls.xfcc_disabled — xfcc_disabled / 012320230021 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-008.md#canonical-3210102113031103-2332201101220100-1103122221002212-3002101012001000-1332031130002011-0021231133303122-2121222212212213-1123320220010203)
- [service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-008.md#canonical-2311132233010120-0223230230122210-0332200203312210-3000132111100333-1133101301233010-3211310033122333-2323002110210000-1300012002110122)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](resources--workload--reference--group-008.md#canonical-1202020123020333-2332300320113330-1321112113331202-2112220001121101-3022001013310321-0020100001212100-0032232320012323-1311122212310031)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](resources--workload--reference--group-009.md#canonical-3212310132100213-2000301220120030-0112323230310220-3200032133332210-3223310210320320-2331013310222010-2010102133130322-2111123310321223)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-010.md#canonical-3230132323111312-0132020110123203-0222302101120231-0202012321012303-1213310020220000-0101203211110131-1303320330302033-3001313221002112)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.use_mtls](resources--workload--reference--group-010.md#canonical-3310220202003120-0020232310001311-1230011321302031-3002200220020220-1002102031303233-0322332031301112-0031333023233103-0133102211311201)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.use_mtls.xfcc_disabled

<a id="canonical-3111330312223100-3212113002300021-3011312000233330-2033302013303200-3130213332312211-0221032030232231-3311213300212120-0332001223203000"></a>

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

<a id="canonical-0030323333121130-2301112021103303-2010123210010212-0311303302220100-3022130132021112-2223121101001310-2000021123013231-2130112232330203"></a>

## Direct properties — xfcc_disabled / 012320230021 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3113322133112103-1303232131121213-2020230232002313-2211333223221032-2230230110123120-1103010002111021-1332321103110310-0023322033112220"></a>

## Next pages — xfcc_disabled / 012320230021 / 4

- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.use_mtls](resources--workload--reference--group-010.md#canonical-3310220202003120-0020232310001311-1230011321302031-3002200220020220-1002102031303233-0322332031301112-0031333023233103-0133102211311201)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-1022001030312122-0110021131331203-1321111312211223-0231321110200201-2232010002033222-1321002220133122-2012110202313210-2130032010303332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3023103223102031-0302023011123113-0102310230221103-1200300312130123-3230220302122220-0300223320112130-0311030030301122-1331003230323133"></a>

## service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.use_mtls.xfcc_options — xfcc_options / 103122231013 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-008.md#canonical-3210102113031103-2332201101220100-1103122221002212-3002101012001000-1332031130002011-0021231133303122-2121222212212213-1123320220010203)
- [service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-008.md#canonical-2311132233010120-0223230230122210-0332200203312210-3000132111100333-1133101301233010-3211310033122333-2323002110210000-1300012002110122)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](resources--workload--reference--group-008.md#canonical-1202020123020333-2332300320113330-1321112113331202-2112220001121101-3022001013310321-0020100001212100-0032232320012323-1311122212310031)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](resources--workload--reference--group-009.md#canonical-3212310132100213-2000301220120030-0112323230310220-3200032133332210-3223310210320320-2331013310222010-2010102133130322-2111123310321223)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-010.md#canonical-3230132323111312-0132020110123203-0222302101120231-0202012321012303-1213310020220000-0101203211110131-1303320330302033-3001313221002112)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.use_mtls](resources--workload--reference--group-010.md#canonical-3310220202003120-0020232310001311-1230011321302031-3002200220020220-1002102031303233-0322332031301112-0031333023233103-0133102211311201)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.use_mtls.xfcc_options

<a id="canonical-3030102300301111-2132021123223002-2331131033102101-0113203222021332-2122131110023202-3302330233311223-1133210311301333-0202330000301013"></a>

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

<a id="canonical-0332202021122223-2111130213212132-2023323031132312-2211023310303212-1213311322313001-1012110230200333-1323223302020111-3032023032323023"></a>

## Direct properties — xfcc_options / 103122231013 / 3

<a id="canonical-3021102022303222-3100111111023010-1211032231033322-3303321133211131-2113102122223303-2001111131202233-2120211303313320-3020132200313320"></a>

<a id="canonical-3302203010302032-1222201213322102-2230033012002021-3000211032312210-0200312000033032-0203230200112311-1303131022332323-0202321223002322"></a>

## xfcc_header_elements property — xfcc_options / 103122231013 / 4

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

<a id="canonical-0310131331310132-1112030022313220-1202301002330233-0001211113011122-3331220310100201-3113002031221123-0013122130221112-3000311302202111"></a>

## Next pages — xfcc_options / 103122231013 / 5

- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.use_mtls](resources--workload--reference--group-010.md#canonical-3310220202003120-0020232310001311-1230011321302031-3002200220020220-1002102031303233-0322332031301112-0031333023233103-0133102211311201)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-0123203011200310-2323132333110331-0320211012002122-2323013220112302-2000030132131023-1312320130303213-2223130103233300-3130200202200302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0110321121232120-3321012331231232-1221301112323021-1322110330200133-1302112320200223-2030312221011202-1100310303112130-2000332331123331"></a>

## service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert — https_auto_cert / 120012321113 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-008.md#canonical-3210102113031103-2332201101220100-1103122221002212-3002101012001000-1332031130002011-0021231133303122-2121222212212213-1123320220010203)
- [service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-008.md#canonical-2311132233010120-0223230230122210-0332200203312210-3000132111100333-1133101301233010-3211310033122333-2323002110210000-1300012002110122)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](resources--workload--reference--group-008.md#canonical-1202020123020333-2332300320113330-1321112113331202-2112220001121101-3022001013310321-0020100001212100-0032232320012323-1311122212310031)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert

<a id="canonical-1232112013302333-0303103031303231-1222031113202031-3331003311332113-0013301133323011-0302111121301200-1222300233002111-1220032321103222"></a>

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

<a id="canonical-0300031012231133-2030221103321213-2110200132332111-3022323322212321-3121222123021120-2010001320103002-2330010233122300-3020332003321012"></a>

## Direct properties — https_auto_cert / 120012321113 / 3

<a id="canonical-1323133111232320-3112232100031121-2202021111201223-2133333211101030-2131122303022111-0133312013012322-0101233220003322-1023101020331222"></a>

<a id="canonical-1022103211231020-1220021230020200-0210132111310031-2331031021102330-0132310320330221-3333313133333010-2210011310213012-2300002012131321"></a>

## add_hsts property — https_auto_cert / 120012321113 / 4

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

<a id="canonical-3123302120010331-1230221202032032-1212012102323313-0012100223331031-3133022231220003-1233020310102221-3302210300301033-0021230202220333"></a>

<a id="canonical-1030213122030302-1322000300101220-0310301311103201-0311220201211230-2221300131101000-2131212131203303-1022323032210201-3321000330302323"></a>

## append_server_name property — https_auto_cert / 120012321113 / 5

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

- [coalescing_options](resources--workload--reference--group-010.md#canonical-3210303103310002-2310113031123222-2101231231122213-2233021130122303-0233321001201002-0330322101202322-2003321331113003-0022133312300232): complete subsection reference.

<a id="canonical-3132312010312213-2323033210000301-0012312330222002-0301031213332012-2211201112113331-0311231122310110-0302303011200231-2012112203222211"></a>

<a id="canonical-0232313132020100-0100132233321010-1002012223332021-1101220230133230-3121201030002023-0003332012202010-1103011122121012-3001013303032033"></a>

## connection_idle_timeout property — https_auto_cert / 120012321113 / 6

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

- [default_header](resources--workload--reference--group-010.md#canonical-2212010221300233-2201101023002210-3010002022323233-3113011212030131-3032203233302230-0231011311211003-2200023023000231-0320202103113033): complete subsection reference.

- [default_loadbalancer](resources--workload--reference--group-010.md#canonical-0331230213133123-1130310110210111-1321130212223101-0320232323122321-3111200212112300-3031202031023212-0122001001002002-3111023022301313): complete subsection reference.

- [disable_path_normalize](resources--workload--reference--group-010.md#canonical-0123302220301321-0132211023131201-3220202102311211-2131030202102302-0221310222332302-3321210201332223-3210110121201203-0011033011213313): complete subsection reference.

- [enable_path_normalize](resources--workload--reference--group-010.md#canonical-2301110231331322-3310221221300232-1201002123223112-0122032021031011-0101203313332122-1312001131032013-1101002312122203-1030033232131121): complete subsection reference.

- [http_protocol_options](resources--workload--reference--group-010.md#canonical-3321233023232103-1212213330303210-1111113130313312-1323223120210100-2202312311101302-1303212320303032-3030011002101032-0231021313323001): complete subsection reference.

<a id="canonical-3112313333123023-0132221002100312-1023212222003312-3132111201200113-1311223332010320-0221013032333202-2012300020320102-2320003101330223"></a>

<a id="canonical-2030300301003203-3103311001120031-2303320201012331-0011211001022013-0013210121110012-1000220110203021-0023301113212213-3210012200011202"></a>

## http_redirect property — https_auto_cert / 120012321113 / 7

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

- [no_mtls](resources--workload--reference--group-011.md#canonical-1222331122321301-0002020313232131-2111200221222211-0133120203310331-3111211101033101-3212020231221033-2003023203022002-0202201133311302): complete subsection reference.

- [non_default_loadbalancer](resources--workload--reference--group-011.md#canonical-2212121030220022-1122202123012230-2320133132212021-1010311121113312-3023212110310032-2101332003123313-0112003011103233-1110123011301312): complete subsection reference.

- [pass_through](resources--workload--reference--group-011.md#canonical-0011102022130121-2312331112021323-3112133222310201-1022122230332333-1001102130231330-0303221022113011-0131100333200021-3212121001001113): complete subsection reference.

<a id="canonical-2312202332200003-2311111301131023-1102300300300101-3300201213203023-2111201211122230-0311033133113211-1201121032223203-2311221110121113"></a>

<a id="canonical-2002030231220330-1203111033322110-1011101320032303-0300131021032031-0001131301322311-1110023012130210-3111300233131220-2112110301331112"></a>

## port property — https_auto_cert / 120012321113 / 8

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-1212100333110003-0010030003023213-1331133201222211-0023101302331222-2113003102130220-1211023021321100-1223122303012002-2221310200100222"></a>

<a id="canonical-1230013101201011-1120001113210220-0020331232230000-0030022131110332-0100013300213001-3223231102231132-0032220201220330-2323311212030332"></a>

## port_ranges property — https_auto_cert / 120012321113 / 9

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-1010030321220303-1302202103201122-3302102232200220-1322202103121200-2303220033323201-1131100210033120-1123130023323033-1212311121112030"></a>

<a id="canonical-2123022310231011-2122333303120003-3033012301220103-2002012200333103-2222201231001311-3102333012002332-0223331231130030-3110120313330021"></a>

## server_name property — https_auto_cert / 120012321113 / 10

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

- [tls_config](resources--workload--reference--group-011.md#canonical-3113321322213012-1133311332223211-0212301313120320-2200201201212023-3113301313231213-2111031111233222-3330001032333212-3112113321000021): complete subsection reference.

- [use_mtls](resources--workload--reference--group-011.md#canonical-1112020130200033-1212030123322311-1103211123233210-3303131123200003-0033323320312331-1303132312122133-3032112313313232-1220000112111220): complete subsection reference.

<a id="canonical-0301023200210023-3330032011000323-0303322100101132-1330323122231200-2332133112010210-2211110101000033-3010003022300032-0300023003233123"></a>

## Next pages — https_auto_cert / 120012321113 / 11

- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.coalescing_options](resources--workload--reference--group-010.md#canonical-3210303103310002-2310113031123222-2101231231122213-2233021130122303-0233321001201002-0330322101202322-2003321331113003-0022133312300232)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.default_header](resources--workload--reference--group-010.md#canonical-2212010221300233-2201101023002210-3010002022323233-3113011212030131-3032203233302230-0231011311211003-2200023023000231-0320202103113033)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.default_loadbalancer](resources--workload--reference--group-010.md#canonical-0331230213133123-1130310110210111-1321130212223101-0320232323122321-3111200212112300-3031202031023212-0122001001002002-3111023022301313)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.disable_path_normalize](resources--workload--reference--group-010.md#canonical-0123302220301321-0132211023131201-3220202102311211-2131030202102302-0221310222332302-3321210201332223-3210110121201203-0011033011213313)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.enable_path_normalize](resources--workload--reference--group-010.md#canonical-2301110231331322-3310221221300232-1201002123223112-0122032021031011-0101203313332122-1312001131032013-1101002312122203-1030033232131121)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.http_protocol_options](resources--workload--reference--group-010.md#canonical-3321233023232103-1212213330303210-1111113130313312-1323223120210100-2202312311101302-1303212320303032-3030011002101032-0231021313323001)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.no_mtls](resources--workload--reference--group-011.md#canonical-1222331122321301-0002020313232131-2111200221222211-0133120203310331-3111211101033101-3212020231221033-2003023203022002-0202201133311302)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.non_default_loadbalancer](resources--workload--reference--group-011.md#canonical-2212121030220022-1122202123012230-2320133132212021-1010311121113312-3023212110310032-2101332003123313-0112003011103233-1110123011301312)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.pass_through](resources--workload--reference--group-011.md#canonical-0011102022130121-2312331112021323-3112133222310201-1022122230332333-1001102130231330-0303221022113011-0131100333200021-3212121001001113)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.tls_config](resources--workload--reference--group-011.md#canonical-3113321322213012-1133311332223211-0212301313120320-2200201201212023-3113301313231213-2111031111233222-3330001032333212-3112113321000021)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.use_mtls](resources--workload--reference--group-011.md#canonical-1112020130200033-1212030123322311-1103211123233210-3303131123200003-0033323320312331-1303132312122133-3032112313313232-1220000112111220)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](resources--workload--reference--group-008.md#canonical-1202020123020333-2332300320113330-1321112113331202-2112220001121101-3022001013310321-0020100001212100-0032232320012323-1311122212310031)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-3210303103310002-2310113031123222-2101231231122213-2233021130122303-0233321001201002-0330322101202322-2003321331113003-0022133312300232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2301010322200001-2132031111023300-3120222231010330-0210332100201033-1222211203032211-1321201232200231-2311003000102333-3210133003000130"></a>

## service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.coalescing_options — coalescing_options / 130030300120 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-008.md#canonical-3210102113031103-2332201101220100-1103122221002212-3002101012001000-1332031130002011-0021231133303122-2121222212212213-1123320220010203)
- [service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-008.md#canonical-2311132233010120-0223230230122210-0332200203312210-3000132111100333-1133101301233010-3211310033122333-2323002110210000-1300012002110122)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](resources--workload--reference--group-008.md#canonical-1202020123020333-2332300320113330-1321112113331202-2112220001121101-3022001013310321-0020100001212100-0032232320012323-1311122212310031)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert](resources--workload--reference--group-010.md#canonical-0123203011200310-2323132333110331-0320211012002122-2323013220112302-2000030132131023-1312320130303213-2223130103233300-3130200202200302)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.coalescing_options

<a id="canonical-0032231120300313-2312012110303100-0303011330133331-2312133233023020-1302212330033032-1010202323322031-0202330220101333-2223023023312000"></a>

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

<a id="canonical-3330310010333102-2230201000310223-0311233303231200-2021021221013322-0313020312013313-3301113002031202-1310332213113321-3102102211233022"></a>

## Direct properties — coalescing_options / 130030300120 / 3

- [default_coalescing](resources--workload--reference--group-010.md#canonical-3001201313233020-0031113321003221-1223021203033001-0303222112103221-2020323230012131-1120012022332202-2033301103122200-3322012122333032): complete subsection reference.

- [strict_coalescing](resources--workload--reference--group-010.md#canonical-3001101301333012-1001032231123013-0022102001110211-2121001331021222-2121200030001103-3010003322331333-3202230111110223-0131101301031023): complete subsection reference.

<a id="canonical-0000212100031221-1120211331323223-1321112021132123-1102312011002033-2310333210020222-1133332113312113-2203121232031333-1322300212330131"></a>

## Next pages — coalescing_options / 130030300120 / 4

- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.coalescing_options.default_coalescing](resources--workload--reference--group-010.md#canonical-3001201313233020-0031113321003221-1223021203033001-0303222112103221-2020323230012131-1120012022332202-2033301103122200-3322012122333032)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.coalescing_options.strict_coalescing](resources--workload--reference--group-010.md#canonical-3001101301333012-1001032231123013-0022102001110211-2121001331021222-2121200030001103-3010003322331333-3202230111110223-0131101301031023)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert](resources--workload--reference--group-010.md#canonical-0123203011200310-2323132333110331-0320211012002122-2323013220112302-2000030132131023-1312320130303213-2223130103233300-3130200202200302)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-3001201313233020-0031113321003221-1223021203033001-0303222112103221-2020323230012131-1120012022332202-2033301103122200-3322012122333032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3030100310102103-3212321112131321-1231132312120200-1133312113130321-1111002301230322-3111111012210233-2211320012112112-3033302320220000"></a>

## service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.coalescing_options.default_coalescing — default_coalescing / 030133211331 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-008.md#canonical-3210102113031103-2332201101220100-1103122221002212-3002101012001000-1332031130002011-0021231133303122-2121222212212213-1123320220010203)
- [service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-008.md#canonical-2311132233010120-0223230230122210-0332200203312210-3000132111100333-1133101301233010-3211310033122333-2323002110210000-1300012002110122)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](resources--workload--reference--group-008.md#canonical-1202020123020333-2332300320113330-1321112113331202-2112220001121101-3022001013310321-0020100001212100-0032232320012323-1311122212310031)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert](resources--workload--reference--group-010.md#canonical-0123203011200310-2323132333110331-0320211012002122-2323013220112302-2000030132131023-1312320130303213-2223130103233300-3130200202200302)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.coalescing_options](resources--workload--reference--group-010.md#canonical-3210303103310002-2310113031123222-2101231231122213-2233021130122303-0233321001201002-0330322101202322-2003321331113003-0022133312300232)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.coalescing_options.default_coalescing

<a id="canonical-1113301221232113-0231122113222320-3111202102101023-2202023233220123-3301331020333033-0213131031312232-1000031310013100-0323313131222200"></a>

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

<a id="canonical-0232311221031020-3230300310313113-0323123001130113-1223330302312331-1221301232120303-3123311131002032-2101131333133223-0110310101103310"></a>

## Direct properties — default_coalescing / 030133211331 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0131332030333112-3200200002332222-1200322313003022-1120011320113000-0131333323032120-0213202010032133-0031311213031322-3131002200010023"></a>

## Next pages — default_coalescing / 030133211331 / 4

- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.coalescing_options](resources--workload--reference--group-010.md#canonical-3210303103310002-2310113031123222-2101231231122213-2233021130122303-0233321001201002-0330322101202322-2003321331113003-0022133312300232)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-3001101301333012-1001032231123013-0022102001110211-2121001331021222-2121200030001103-3010003322331333-3202230111110223-0131101301031023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3200311221203321-1021211032201222-3303100003203011-2300133333303022-0013313223321302-2020003102132112-0000330023212312-1110321200110013"></a>

## service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.coalescing_options.strict_coalescing — strict_coalescing / 231022011030 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-008.md#canonical-3210102113031103-2332201101220100-1103122221002212-3002101012001000-1332031130002011-0021231133303122-2121222212212213-1123320220010203)
- [service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-008.md#canonical-2311132233010120-0223230230122210-0332200203312210-3000132111100333-1133101301233010-3211310033122333-2323002110210000-1300012002110122)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](resources--workload--reference--group-008.md#canonical-1202020123020333-2332300320113330-1321112113331202-2112220001121101-3022001013310321-0020100001212100-0032232320012323-1311122212310031)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert](resources--workload--reference--group-010.md#canonical-0123203011200310-2323132333110331-0320211012002122-2323013220112302-2000030132131023-1312320130303213-2223130103233300-3130200202200302)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.coalescing_options](resources--workload--reference--group-010.md#canonical-3210303103310002-2310113031123222-2101231231122213-2233021130122303-0233321001201002-0330322101202322-2003321331113003-0022133312300232)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.coalescing_options.strict_coalescing

<a id="canonical-3122312100033321-3122201130333033-2002021320033212-3030003021123132-2211100012203200-1221200031113113-3110022032332232-1203313100200120"></a>

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

<a id="canonical-2112130200132303-3100011221301130-0031112203132211-1023012201330111-1000202012332133-3123221130212323-0302220001120331-2220233211100022"></a>

## Direct properties — strict_coalescing / 231022011030 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2022221111232000-1313010123030020-0222202230130233-0220033200301313-1113013120330333-3331332333330202-2011112100301103-0310110130131013"></a>

## Next pages — strict_coalescing / 231022011030 / 4

- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.coalescing_options](resources--workload--reference--group-010.md#canonical-3210303103310002-2310113031123222-2101231231122213-2233021130122303-0233321001201002-0330322101202322-2003321331113003-0022133312300232)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-2212010221300233-2201101023002210-3010002022323233-3113011212030131-3032203233302230-0231011311211003-2200023023000231-0320202103113033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2032212232021303-1202033030021312-0322310000022012-2111233122121102-1233320320020030-3023332113331223-1302103131023313-1122110113332113"></a>

## service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.default_header — default_header / 033310330300 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-008.md#canonical-3210102113031103-2332201101220100-1103122221002212-3002101012001000-1332031130002011-0021231133303122-2121222212212213-1123320220010203)
- [service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-008.md#canonical-2311132233010120-0223230230122210-0332200203312210-3000132111100333-1133101301233010-3211310033122333-2323002110210000-1300012002110122)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](resources--workload--reference--group-008.md#canonical-1202020123020333-2332300320113330-1321112113331202-2112220001121101-3022001013310321-0020100001212100-0032232320012323-1311122212310031)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert](resources--workload--reference--group-010.md#canonical-0123203011200310-2323132333110331-0320211012002122-2323013220112302-2000030132131023-1312320130303213-2223130103233300-3130200202200302)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.default_header

<a id="canonical-2103110332300320-1130103030222311-3000012332103200-3003311111112302-0132003013003102-1321331322123002-0122231313232121-3331000003333123"></a>

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

<a id="canonical-0221221223003130-0323110130331300-3131322122223203-0201002231001003-1212133100013010-0321322001110131-3030011102210003-1100002102301001"></a>

## Direct properties — default_header / 033310330300 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1232122102322130-3132220221231120-3310013001022023-2102102232310020-3112013233300102-0102012310130011-3332003003012211-1121313123000203"></a>

## Next pages — default_header / 033310330300 / 4

- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert](resources--workload--reference--group-010.md#canonical-0123203011200310-2323132333110331-0320211012002122-2323013220112302-2000030132131023-1312320130303213-2223130103233300-3130200202200302)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-0331230213133123-1130310110210111-1321130212223101-0320232323122321-3111200212112300-3031202031023212-0122001001002002-3111023022301313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1123133023112010-0212313321202211-0213302311312012-1112332312220103-2312131303330323-0021122311220101-3212100131302220-1110122102010102"></a>

## service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.default_loadbalancer — default_loadbalancer / 233031231321 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-008.md#canonical-3210102113031103-2332201101220100-1103122221002212-3002101012001000-1332031130002011-0021231133303122-2121222212212213-1123320220010203)
- [service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-008.md#canonical-2311132233010120-0223230230122210-0332200203312210-3000132111100333-1133101301233010-3211310033122333-2323002110210000-1300012002110122)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](resources--workload--reference--group-008.md#canonical-1202020123020333-2332300320113330-1321112113331202-2112220001121101-3022001013310321-0020100001212100-0032232320012323-1311122212310031)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert](resources--workload--reference--group-010.md#canonical-0123203011200310-2323132333110331-0320211012002122-2323013220112302-2000030132131023-1312320130303213-2223130103233300-3130200202200302)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.default_loadbalancer

<a id="canonical-1101130120102203-1021132231003031-0031100212231322-1302322322203203-2212220310311332-0311033330211021-3120301012231111-0123122231201230"></a>

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

<a id="canonical-3012030113120130-1111222030223000-0302201202011113-2210112302320100-1030020112112001-0330212010032303-1120120312122331-2302303002000111"></a>

## Direct properties — default_loadbalancer / 233031231321 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2230320231001333-0320123003030120-1032230032133321-3212313332030321-0323033220003122-2100221001322201-1312330013023203-2213121120022200"></a>

## Next pages — default_loadbalancer / 233031231321 / 4

- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert](resources--workload--reference--group-010.md#canonical-0123203011200310-2323132333110331-0320211012002122-2323013220112302-2000030132131023-1312320130303213-2223130103233300-3130200202200302)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-0123302220301321-0132211023131201-3220202102311211-2131030202102302-0221310222332302-3321210201332223-3210110121201203-0011033011213313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1111101303330133-0220002310011323-3132331322300132-1202320232002131-2231233022210203-1203111303133013-0322022012130133-3103321000031331"></a>

## service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.disable_path_normalize — disable_path_normalize / 120311201013 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-008.md#canonical-3210102113031103-2332201101220100-1103122221002212-3002101012001000-1332031130002011-0021231133303122-2121222212212213-1123320220010203)
- [service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-008.md#canonical-2311132233010120-0223230230122210-0332200203312210-3000132111100333-1133101301233010-3211310033122333-2323002110210000-1300012002110122)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](resources--workload--reference--group-008.md#canonical-1202020123020333-2332300320113330-1321112113331202-2112220001121101-3022001013310321-0020100001212100-0032232320012323-1311122212310031)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert](resources--workload--reference--group-010.md#canonical-0123203011200310-2323132333110331-0320211012002122-2323013220112302-2000030132131023-1312320130303213-2223130103233300-3130200202200302)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.disable_path_normalize

<a id="canonical-2313312033200311-2020223013103232-3320231001211002-1200220310100110-0213221001030311-2322021201201020-3021221020012103-0233331202231021"></a>

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

<a id="canonical-0020000032013122-0232313300001311-1121123011302010-2111113013203113-0303230210200032-2222211320232030-3200201031330010-0011210130310322"></a>

## Direct properties — disable_path_normalize / 120311201013 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2200211123003002-3103313013223131-0131122101020300-2001212300201211-0203000011323131-0320130211231220-2103313323101332-3213301333201213"></a>

## Next pages — disable_path_normalize / 120311201013 / 4

- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert](resources--workload--reference--group-010.md#canonical-0123203011200310-2323132333110331-0320211012002122-2323013220112302-2000030132131023-1312320130303213-2223130103233300-3130200202200302)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-2301110231331322-3310221221300232-1201002123223112-0122032021031011-0101203313332122-1312001131032013-1101002312122203-1030033232131121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1203011320312122-1113313322331023-1021332101212113-2321200233102112-2323020331113323-2110010021001311-3033220120002132-3231123230123222"></a>

## service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.enable_path_normalize — enable_path_normalize / 023323020100 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-008.md#canonical-3210102113031103-2332201101220100-1103122221002212-3002101012001000-1332031130002011-0021231133303122-2121222212212213-1123320220010203)
- [service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-008.md#canonical-2311132233010120-0223230230122210-0332200203312210-3000132111100333-1133101301233010-3211310033122333-2323002110210000-1300012002110122)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](resources--workload--reference--group-008.md#canonical-1202020123020333-2332300320113330-1321112113331202-2112220001121101-3022001013310321-0020100001212100-0032232320012323-1311122212310031)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert](resources--workload--reference--group-010.md#canonical-0123203011200310-2323132333110331-0320211012002122-2323013220112302-2000030132131023-1312320130303213-2223130103233300-3130200202200302)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.enable_path_normalize

<a id="canonical-2320121220311111-3333022012222323-3200311102331320-0302212003020121-3123300021110301-3131130203321233-1110232010321323-2202110223300331"></a>

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

<a id="canonical-3000223032001202-1000320312312302-3312111213230021-2312003310310210-1131130000321213-3021222111222300-3322330313201221-3122211322101110"></a>

## Direct properties — enable_path_normalize / 023323020100 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3332131022230222-1012312013020131-0120332221331332-2312321100001110-2110011010233313-1001311001202023-1312333023132133-1032221232220112"></a>

## Next pages — enable_path_normalize / 023323020100 / 4

- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert](resources--workload--reference--group-010.md#canonical-0123203011200310-2323132333110331-0320211012002122-2323013220112302-2000030132131023-1312320130303213-2223130103233300-3130200202200302)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-3321233023232103-1212213330303210-1111113130313312-1323223120210100-2202312311101302-1303212320303032-3030011002101032-0231021313323001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0123113233210311-1111231200031012-2121220211312322-1211310332331021-2103110103213001-3202222132332230-3310031333220301-3012222122332231"></a>

## service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.http_protocol_options — http_protocol_options / 000103230110 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-008.md#canonical-3210102113031103-2332201101220100-1103122221002212-3002101012001000-1332031130002011-0021231133303122-2121222212212213-1123320220010203)
- [service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-008.md#canonical-2311132233010120-0223230230122210-0332200203312210-3000132111100333-1133101301233010-3211310033122333-2323002110210000-1300012002110122)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](resources--workload--reference--group-008.md#canonical-1202020123020333-2332300320113330-1321112113331202-2112220001121101-3022001013310321-0020100001212100-0032232320012323-1311122212310031)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert](resources--workload--reference--group-010.md#canonical-0123203011200310-2323132333110331-0320211012002122-2323013220112302-2000030132131023-1312320130303213-2223130103233300-3130200202200302)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.http_protocol_options

<a id="canonical-1223310002012002-0110033021102331-3310330013101112-0002100221312013-0322312333233032-3013230310211102-0331132200301012-3212123311222202"></a>

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

<a id="canonical-2201201030331322-1223120330203030-0311032321310330-0132231213113101-1231033330311112-1313231310023130-3202133023310130-3100313202120002"></a>

## Direct properties — http_protocol_options / 000103230110 / 3

- [http_protocol_enable_v1_only](resources--workload--reference--group-011.md#canonical-3131320122021222-1122202213031111-0311122112320212-1002111302121300-1312200303033032-1020201110003003-0211332003112122-2113212112233200): complete subsection reference.

- [http_protocol_enable_v1_v2](resources--workload--reference--group-011.md#canonical-1223222121310230-1300021101211033-1132021133012330-1010020023021121-1313030322131100-1023321112331222-2231313112223331-3321101313100120): complete subsection reference.

- [http_protocol_enable_v2_only](resources--workload--reference--group-011.md#canonical-3030001013313220-3110300130210113-1031003122323110-3321330320103001-3020233032232131-1322033202000033-3133030023223232-2200133133102122): complete subsection reference.

<a id="canonical-1302000321312320-2211010320100220-1133011233132313-0103213003220230-1201010001201122-3312323013232031-2213013220310021-3120010202313130"></a>

## Next pages — http_protocol_options / 000103230110 / 4

- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only](resources--workload--reference--group-011.md#canonical-3131320122021222-1122202213031111-0311122112320212-1002111302121300-1312200303033032-1020201110003003-0211332003112122-2113212112233200)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_v2](resources--workload--reference--group-011.md#canonical-1223222121310230-1300021101211033-1132021133012330-1010020023021121-1313030322131100-1023321112331222-2231313112223331-3321101313100120)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v2_only](resources--workload--reference--group-011.md#canonical-3030001013313220-3110300130210113-1031003122323110-3321330320103001-3020233032232131-1322033202000033-3133030023223232-2200133133102122)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert](resources--workload--reference--group-010.md#canonical-0123203011200310-2323132333110331-0320211012002122-2323013220112302-2000030132131023-1312320130303213-2223130103233300-3130200202200302)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
