---
page_title: "xcsh_workload reference"
subcategory: "Container"
description: "Complete grouped canonical reference for xcsh_workload reference."
---

# xcsh_workload reference

<a id="canonical-3312331021300201-3231310200321032-2100202230331020-3132212022330212-1013202211321002-0130033332232313-1233112222032030-2320100210130312"></a>

## domains property — http_loadbalancer / 321311303010 / 4

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

- [http](resources--workload--reference--group-009.md#canonical-2032123312300020-0313233113231322-1312210332031202-2133210330020310-0110031022301023-1212321320020231-2033231203003230-2310210201131220): complete subsection reference.

- [https](resources--workload--reference--group-009.md#canonical-3212310132100213-2000301220120030-0112323230310220-3200032133332210-3223310210320320-2331013310222010-2010102133130322-2111123310321223): complete subsection reference.

- [https_auto_cert](resources--workload--reference--group-010.md#canonical-0123203011200310-2323132333110331-0320211012002122-2323013220112302-2000030132131023-1312320130303213-2223130103233300-3130200202200302): complete subsection reference.

- [specific_routes](resources--workload--reference--group-011.md#canonical-3333233100013331-2322011200101232-2320001330000031-2000212200300310-3210113330031012-2012103222021320-0322112233332312-3331120230102200): complete subsection reference.

<a id="canonical-3200310312213013-2020102312002110-0000111311313201-0300102203212231-2132210100223000-2200112333021320-3300010001101103-0201320130322020"></a>

## Next pages — http_loadbalancer / 321311303010 / 5

- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.default_route](resources--workload--reference--group-009.md#canonical-0232013311111022-0003320023211122-1321321011010330-1223120121220333-0213130300023300-2001123101323323-3002132100221212-2311012111330101)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.http](resources--workload--reference--group-009.md#canonical-2032123312300020-0313233113231322-1312210332031202-2133210330020310-0110031022301023-1212321320020231-2033231203003230-2310210201131220)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](resources--workload--reference--group-009.md#canonical-3212310132100213-2000301220120030-0112323230310220-3200032133332210-3223310210320320-2331013310222010-2010102133130322-2111123310321223)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert](resources--workload--reference--group-010.md#canonical-0123203011200310-2323132333110331-0320211012002122-2323013220112302-2000030132131023-1312320130303213-2223130103233300-3130200202200302)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes](resources--workload--reference--group-011.md#canonical-3333233100013331-2322011200101232-2320001330000031-2000212200300310-3210113330031012-2012103222021320-0322112233332312-3331120230102200)
- [service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-008.md#canonical-2311132233010120-0223230230122210-0332200203312210-3000132111100333-1133101301233010-3211310033122333-2323002110210000-1300012002110122)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-0232013311111022-0003320023211122-1321321011010330-1223120121220333-0213130300023300-2001123101323323-3002132100221212-2311012111330101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1013322212100213-1012022103231212-3012110111323301-2002032133133031-2011310302320220-3102202010323302-3022231120101111-1002002132331203"></a>

## service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.default_route — default_route / 110120030200 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-008.md#canonical-3210102113031103-2332201101220100-1103122221002212-3002101012001000-1332031130002011-0021231133303122-2121222212212213-1123320220010203)
- [service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-008.md#canonical-2311132233010120-0223230230122210-0332200203312210-3000132111100333-1133101301233010-3211310033122333-2323002110210000-1300012002110122)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](resources--workload--reference--group-008.md#canonical-1202020123020333-2332300320113330-1321112113331202-2112220001121101-3022001013310321-0020100001212100-0032232320012323-1311122212310031)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.default_route

<a id="canonical-0032213203000102-1303101223331031-3301022322131211-0111132120030201-2201330132133323-3013023020210223-3102103202021013-1302312200212130"></a>

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

<a id="canonical-3201323120202231-2333133312312301-0030113110002131-0000011323313230-3100210131221011-2023000122233320-1331212320113022-3130132031111121"></a>

## Direct properties — default_route / 110120030200 / 3

- [auto_host_rewrite](resources--workload--reference--group-009.md#canonical-2322120332220030-0300010303222132-0231233230310120-3312130212030203-3323022301121303-1300013301332323-1310232112322310-3321122202110221): complete subsection reference.

- [disable_host_rewrite](resources--workload--reference--group-009.md#canonical-1211002103313333-1231121300012003-1312223130210212-0002230330122201-0012213033030211-0020320230121233-2113131113313320-3110331123031333): complete subsection reference.

<a id="canonical-1333232102233301-2200021232321101-1212202010312032-0001211311320112-3010213230112123-0010130011022311-1303013101323111-3230333320313332"></a>

<a id="canonical-0002031233210033-0203110123122213-3000312332132331-0003311300003110-0123033021102112-1021211212220301-0220001300100113-1322112123120110"></a>

## host_rewrite property — default_route / 110120030200 / 4

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
    "ves.io.schema.rules.string.hostname": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true"
  }
}
```

<a id="canonical-3022122311303332-1120022101133130-1301002103000033-2012120120132120-3033103020210200-2223312312202222-1023111123013200-1222201302110100"></a>

## Next pages — default_route / 110120030200 / 5

- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.default_route.auto_host_rewrite](resources--workload--reference--group-009.md#canonical-2322120332220030-0300010303222132-0231233230310120-3312130212030203-3323022301121303-1300013301332323-1310232112322310-3321122202110221)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.default_route.disable_host_rewrite](resources--workload--reference--group-009.md#canonical-1211002103313333-1231121300012003-1312223130210212-0002230330122201-0012213033030211-0020320230121233-2113131113313320-3110331123031333)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](resources--workload--reference--group-008.md#canonical-1202020123020333-2332300320113330-1321112113331202-2112220001121101-3022001013310321-0020100001212100-0032232320012323-1311122212310031)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-2322120332220030-0300010303222132-0231233230310120-3312130212030203-3323022301121303-1300013301332323-1310232112322310-3321122202110221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0220210230201213-0011223021221210-3200130111223001-1233132202020013-0212023302122012-2021321132120301-0211210302110010-3103313013003222"></a>

## service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.default_route.auto_host_rewrite — auto_host_rewrite / 012012003322 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-008.md#canonical-3210102113031103-2332201101220100-1103122221002212-3002101012001000-1332031130002011-0021231133303122-2121222212212213-1123320220010203)
- [service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-008.md#canonical-2311132233010120-0223230230122210-0332200203312210-3000132111100333-1133101301233010-3211310033122333-2323002110210000-1300012002110122)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](resources--workload--reference--group-008.md#canonical-1202020123020333-2332300320113330-1321112113331202-2112220001121101-3022001013310321-0020100001212100-0032232320012323-1311122212310031)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.default_route](resources--workload--reference--group-009.md#canonical-0232013311111022-0003320023211122-1321321011010330-1223120121220333-0213130300023300-2001123101323323-3002132100221212-2311012111330101)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.default_route.auto_host_rewrite

<a id="canonical-2021331302113012-2002312120020201-1201102220322311-3332112321330033-2032311123000131-3103221031211122-2111223213330233-0310230203113012"></a>

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

<a id="canonical-2000330302120210-2022012123030333-3333312010010032-3301123223201301-1010310232221033-2031221122032003-2210003122301330-2002033022003101"></a>

## Direct properties — auto_host_rewrite / 012012003322 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1003112113133013-0301300121321221-3130210023033331-1310001102112111-0010013002333230-0210002122303033-3013203131301200-0031031301323020"></a>

## Next pages — auto_host_rewrite / 012012003322 / 4

- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.default_route](resources--workload--reference--group-009.md#canonical-0232013311111022-0003320023211122-1321321011010330-1223120121220333-0213130300023300-2001123101323323-3002132100221212-2311012111330101)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-1211002103313333-1231121300012003-1312223130210212-0002230330122201-0012213033030211-0020320230121233-2113131113313320-3110331123031333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2113023020230223-2321131003210322-1113313212201101-1232021101310121-0132023330020202-2003213223000103-1233033100032223-1022200111132130"></a>

## service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.default_route.disable_host_rewrite — disable_host_rewrite / 303130201020 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-008.md#canonical-3210102113031103-2332201101220100-1103122221002212-3002101012001000-1332031130002011-0021231133303122-2121222212212213-1123320220010203)
- [service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-008.md#canonical-2311132233010120-0223230230122210-0332200203312210-3000132111100333-1133101301233010-3211310033122333-2323002110210000-1300012002110122)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](resources--workload--reference--group-008.md#canonical-1202020123020333-2332300320113330-1321112113331202-2112220001121101-3022001013310321-0020100001212100-0032232320012323-1311122212310031)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.default_route](resources--workload--reference--group-009.md#canonical-0232013311111022-0003320023211122-1321321011010330-1223120121220333-0213130300023300-2001123101323323-3002132100221212-2311012111330101)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.default_route.disable_host_rewrite

<a id="canonical-2212013233013002-1312320020100322-3302231222322102-0110111123312230-3100011031213210-0312300200200023-2001211011112031-2322002010130110"></a>

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

<a id="canonical-3312011013022331-3130101001002131-2101223131122101-0333130132310212-2032012331211200-1023232231120300-2012313111313213-1132231002333032"></a>

## Direct properties — disable_host_rewrite / 303130201020 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1103132031211030-2111202320230120-3321110330301233-1022230131121030-3231033011023112-1303112322303032-1012201020200031-1130200331020222"></a>

## Next pages — disable_host_rewrite / 303130201020 / 4

- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.default_route](resources--workload--reference--group-009.md#canonical-0232013311111022-0003320023211122-1321321011010330-1223120121220333-0213130300023300-2001123101323323-3002132100221212-2311012111330101)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-2032123312300020-0313233113231322-1312210332031202-2133210330020310-0110031022301023-1212321320020231-2033231203003230-2310210201131220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3312112202231101-2220333012310323-3011311002330012-1030000321200021-3230021103122212-2133320311303231-2103113333321120-2013102332120022"></a>

## service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.http — http / 031103110200 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-008.md#canonical-3210102113031103-2332201101220100-1103122221002212-3002101012001000-1332031130002011-0021231133303122-2121222212212213-1123320220010203)
- [service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-008.md#canonical-2311132233010120-0223230230122210-0332200203312210-3000132111100333-1133101301233010-3211310033122333-2323002110210000-1300012002110122)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](resources--workload--reference--group-008.md#canonical-1202020123020333-2332300320113330-1321112113331202-2112220001121101-3022001013310321-0020100001212100-0032232320012323-1311122212310031)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.http

<a id="canonical-2332011233212301-1201303130312203-2102323333231203-0321211302131113-0023013221221131-0220013031113301-0332003232032130-2102031322312301"></a>

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

<a id="canonical-3322212020322231-1032020312313301-3032201323220012-3302223323123233-2310022110110332-1110020013222312-1230021222323331-0201102013233323"></a>

## Direct properties — http / 031103110200 / 3

<a id="canonical-1223312120322332-2111100302101223-3230113112302323-3203010232202120-2030321301012203-1133302133122223-0202313321003212-1200021000203000"></a>

<a id="canonical-0331110211203033-0210332113310321-3021010133003201-3300330033322210-1211302000032133-0002321000023033-3210212130211321-2032120012320320"></a>

## dns_volterra_managed property — http / 031103110200 / 4

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

<a id="canonical-1131223103331013-3031300201323330-2330032102302033-3322012101133320-2321021330220033-0100013222300132-0210313232132210-1012102033112323"></a>

<a id="canonical-3301220020011301-2032201212233021-2022320110121123-1021013221331301-0002201202323122-1130300030011300-3210330232003313-2200023201020111"></a>

## port property — http / 031103110200 / 5

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

<a id="canonical-2003312200311022-1010100231102332-0320022020321302-0020023001130110-0322023012202103-2211132033321131-2130310322303022-1110332121232311"></a>

<a id="canonical-0232123213222222-0102232201211132-3002221230120112-0302202222231032-1002333003222321-1132112330111220-3312132302023303-1203211233031202"></a>

## port_ranges property — http / 031103110200 / 6

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

<a id="canonical-0201103001020203-1112012100232212-3230331303212222-0310102011330031-0011303201313131-0323102232211123-3233031330121113-2211012310330330"></a>

## Next pages — http / 031103110200 / 7

- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](resources--workload--reference--group-008.md#canonical-1202020123020333-2332300320113330-1321112113331202-2112220001121101-3022001013310321-0020100001212100-0032232320012323-1311122212310031)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-3212310132100213-2000301220120030-0112323230310220-3200032133332210-3223310210320320-2331013310222010-2010102133130322-2111123310321223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1200021102301023-1200220023333111-3011331033323131-1322220311203021-0331111212111132-1001130212323221-3211023132322032-2200002030230031"></a>

## service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https — https / 222303210330 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-008.md#canonical-3210102113031103-2332201101220100-1103122221002212-3002101012001000-1332031130002011-0021231133303122-2121222212212213-1123320220010203)
- [service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-008.md#canonical-2311132233010120-0223230230122210-0332200203312210-3000132111100333-1133101301233010-3211310033122333-2323002110210000-1300012002110122)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](resources--workload--reference--group-008.md#canonical-1202020123020333-2332300320113330-1321112113331202-2112220001121101-3022001013310321-0020100001212100-0032232320012323-1311122212310031)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https

<a id="canonical-3130003112033033-3220213301212302-1322310013120233-2201132331312332-0223301310331120-0233313230022200-3203201002323001-2131102210122102"></a>

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

<a id="canonical-3023031033023102-3313000131032031-2013111111011033-3131002133321002-2031030100333131-3133001300112130-2311033220011110-2030323103331200"></a>

## Direct properties — https / 222303210330 / 3

<a id="canonical-2100221011330111-1032323112010021-3000201322301230-1222103123300302-1030301203020232-1303021310232132-3330130113033132-3232223310012012"></a>

<a id="canonical-1020133232212231-1001132210210312-2302013312321033-0203222132120302-0313223133332121-2233013101120030-0001220313212223-0003221122123133"></a>

## add_hsts property — https / 222303210330 / 4

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

<a id="canonical-3122313213122211-2012200202101300-3200121133112200-2311201312030301-0322012303131331-0003233322210012-2121103010012322-2000312111211123"></a>

<a id="canonical-3321022203023313-2321022033311012-0101311110200322-3031130312332333-1130203132131000-1033102132102131-3103212211123321-1212103223012120"></a>

## append_server_name property — https / 222303210330 / 5

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

- [coalescing_options](resources--workload--reference--group-009.md#canonical-2300311201121333-1311112110203331-3310212300212011-2313230012130122-1332203113120320-3020302213011022-2331223222320311-0303223313003001): complete subsection reference.

<a id="canonical-2133123010310131-3103011110210332-3231202110332310-3210123230300111-3013212133112123-2000031202301230-2002213013121130-0331023133021323"></a>

<a id="canonical-3011103101113012-2321130301110121-1201002010202112-1330312232201330-2002333222203331-2301321203020311-1010211322003233-0221121000022300"></a>

## connection_idle_timeout property — https / 222303210330 / 6

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

- [default_header](resources--workload--reference--group-009.md#canonical-0001220200312001-1113113220200311-2121232002032023-3311312103002000-0121010101202330-2111213223210002-2321330322111220-2331333113001213): complete subsection reference.

- [default_loadbalancer](resources--workload--reference--group-009.md#canonical-0030131323133032-3230131331021321-0120331323111323-2311202313312123-2110223231212033-3320331321232223-0102313012203303-2311102313220320): complete subsection reference.

- [disable_path_normalize](resources--workload--reference--group-009.md#canonical-0100323020130132-1101221000003031-1021130013322030-0031130202003212-2332220321111301-0033022220301031-1033200212220030-3302131102131121): complete subsection reference.

- [enable_path_normalize](resources--workload--reference--group-009.md#canonical-1233033122010133-0313233003202220-3123310102130302-0030010200322301-1302221332333123-3300232123111230-0220102032030330-3101210201002321): complete subsection reference.

- [http_protocol_options](resources--workload--reference--group-009.md#canonical-2233110201213030-3333133212100020-2011121111212120-0113211322030320-0002310120033130-3322203130310011-3012120201131022-3120322301300011): complete subsection reference.

<a id="canonical-1312222321012332-0222033323033123-1213012320332303-0120100122120323-2302231202232322-1201102202332332-1013010130022332-2000111332110303"></a>

<a id="canonical-3313102331121203-1332131230230030-2021303110122031-1301331021233113-1300221310010120-3212300230033311-1120010222023020-1000110011133011"></a>

## http_redirect property — https / 222303210330 / 7

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

- [non_default_loadbalancer](resources--workload--reference--group-009.md#canonical-3033302113113321-1000032130332122-0133011213100323-3323120132112130-1222100231201310-3230202123223130-2103032130300101-1013121302101213): complete subsection reference.

- [pass_through](resources--workload--reference--group-009.md#canonical-0203331121231312-0110013320313123-1233303021130123-3331222120302133-1332231120213020-3203030031311331-2322322120023130-0002030203333232): complete subsection reference.

<a id="canonical-1022031321301103-3030302301310003-2010203122211331-0031323101330311-1212030030232202-1313220321200231-3032231202031301-0121200210022301"></a>

<a id="canonical-3033030112332011-1311022031211100-0333011322201332-3321202230002302-2122001232220120-3132232020233032-3213101313300003-0112222320133103"></a>

## port property — https / 222303210330 / 8

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

<a id="canonical-0320321213203201-0302112230213223-2110302210231121-3301023130033122-3013331121231203-2213112331210023-1310300113010311-2210210222310101"></a>

<a id="canonical-0023213211002001-1201322031323301-0332002110213010-2221020203113200-1303021333313301-1000310101003131-2011221221222133-2301303131312223"></a>

## port_ranges property — https / 222303210330 / 9

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

<a id="canonical-0001021113322132-3212310031122323-3111002200002300-2020310311331013-0303300031233100-2203132300101303-1031121301012221-1120132200020023"></a>

<a id="canonical-2320200323220133-2113302210210320-0000320100031001-1322232210321301-2311220111011131-1201133223332111-0022102100012023-0202212332313332"></a>

## server_name property — https / 222303210330 / 10

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

- [tls_cert_params](resources--workload--reference--group-009.md#canonical-0212122113133333-1000102320023221-2003032220233320-0021231111201312-3001031012012332-3313221303003312-0112011332100001-1120030302022032): complete subsection reference.

- [tls_parameters](resources--workload--reference--group-010.md#canonical-3230132323111312-0132020110123203-0222302101120231-0202012321012303-1213310020220000-0101203211110131-1303320330302033-3001313221002112): complete subsection reference.

<a id="canonical-1121302102312202-1000012213123000-1302320200200323-1300222210000222-1110322212313123-0113233232213201-0300221323212313-3332201222323100"></a>

## Next pages — https / 222303210330 / 11

- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.coalescing_options](resources--workload--reference--group-009.md#canonical-2300311201121333-1311112110203331-3310212300212011-2313230012130122-1332203113120320-3020302213011022-2331223222320311-0303223313003001)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.default_header](resources--workload--reference--group-009.md#canonical-0001220200312001-1113113220200311-2121232002032023-3311312103002000-0121010101202330-2111213223210002-2321330322111220-2331333113001213)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.default_loadbalancer](resources--workload--reference--group-009.md#canonical-0030131323133032-3230131331021321-0120331323111323-2311202313312123-2110223231212033-3320331321232223-0102313012203303-2311102313220320)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.disable_path_normalize](resources--workload--reference--group-009.md#canonical-0100323020130132-1101221000003031-1021130013322030-0031130202003212-2332220321111301-0033022220301031-1033200212220030-3302131102131121)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.enable_path_normalize](resources--workload--reference--group-009.md#canonical-1233033122010133-0313233003202220-3123310102130302-0030010200322301-1302221332333123-3300232123111230-0220102032030330-3101210201002321)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options](resources--workload--reference--group-009.md#canonical-2233110201213030-3333133212100020-2011121111212120-0113211322030320-0002310120033130-3322203130310011-3012120201131022-3120322301300011)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.non_default_loadbalancer](resources--workload--reference--group-009.md#canonical-3033302113113321-1000032130332122-0133011213100323-3323120132112130-1222100231201310-3230202123223130-2103032130300101-1013121302101213)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.pass_through](resources--workload--reference--group-009.md#canonical-0203331121231312-0110013320313123-1233303021130123-3331222120302133-1332231120213020-3203030031311331-2322322120023130-0002030203333232)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params](resources--workload--reference--group-009.md#canonical-0212122113133333-1000102320023221-2003032220233320-0021231111201312-3001031012012332-3313221303003312-0112011332100001-1120030302022032)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-010.md#canonical-3230132323111312-0132020110123203-0222302101120231-0202012321012303-1213310020220000-0101203211110131-1303320330302033-3001313221002112)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](resources--workload--reference--group-008.md#canonical-1202020123020333-2332300320113330-1321112113331202-2112220001121101-3022001013310321-0020100001212100-0032232320012323-1311122212310031)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-2300311201121333-1311112110203331-3310212300212011-2313230012130122-1332203113120320-3020302213011022-2331223222320311-0303223313003001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2301032132102101-0122131033123303-3213021022322133-2022012221103030-3130023003030310-2311233133001023-0231301210320002-2022132313112231"></a>

## service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.coalescing_options — coalescing_options / 231312133201 / 2

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
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.coalescing_options

<a id="canonical-0301030130123201-0123110122221103-1123130033220232-3030323332211201-2230120122022033-3330322212223211-3233031112320211-0033320121311223"></a>

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

<a id="canonical-3002221332121322-2100100321002303-0023030330311002-1333233303001101-1101123002120013-3121032300030231-0101202322002323-1103320100021132"></a>

## Direct properties — coalescing_options / 231312133201 / 3

- [default_coalescing](resources--workload--reference--group-009.md#canonical-1330132010122313-2131032322321333-0113223202132310-3231033002211120-0202003003231020-2222210013210331-1121322130003131-3311231033232132): complete subsection reference.

- [strict_coalescing](resources--workload--reference--group-009.md#canonical-3101101001110123-3200031233033322-3101313011322200-2120330033233331-1002233131222001-0321331321022132-1122231221330110-0310103033212031): complete subsection reference.

<a id="canonical-3221301112230203-1213213300233031-3031133301100200-1312313313213221-3320233010331013-3313112200023211-2232320121023023-3112122223323330"></a>

## Next pages — coalescing_options / 231312133201 / 4

- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.coalescing_options.default_coalescing](resources--workload--reference--group-009.md#canonical-1330132010122313-2131032322321333-0113223202132310-3231033002211120-0202003003231020-2222210013210331-1121322130003131-3311231033232132)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.coalescing_options.strict_coalescing](resources--workload--reference--group-009.md#canonical-3101101001110123-3200031233033322-3101313011322200-2120330033233331-1002233131222001-0321331321022132-1122231221330110-0310103033212031)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](resources--workload--reference--group-009.md#canonical-3212310132100213-2000301220120030-0112323230310220-3200032133332210-3223310210320320-2331013310222010-2010102133130322-2111123310321223)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-1330132010122313-2131032322321333-0113223202132310-3231033002211120-0202003003231020-2222210013210331-1121322130003131-3311231033232132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1322022211200223-1032302113133232-1321230331223010-1302122221301301-3101320312200111-1303031313232112-0223122102332000-1003332200201133"></a>

## service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.coalescing_options.default_coalescing — default_coalescing / 222103130312 / 2

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
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.coalescing_options](resources--workload--reference--group-009.md#canonical-2300311201121333-1311112110203331-3310212300212011-2313230012130122-1332203113120320-3020302213011022-2331223222320311-0303223313003001)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.coalescing_options.default_coalescing

<a id="canonical-2101301000313230-3301202001110210-3123020132323123-2220023123301100-2102320312331010-1132212312330303-2202321020332302-2211200122102301"></a>

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

<a id="canonical-2230320012002321-3222323100310221-3310233300003221-2202310002312313-2300300302320030-2233122121122021-0213333102233122-3111222331311220"></a>

## Direct properties — default_coalescing / 222103130312 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3011323331100121-3310133031001030-2232310002113321-2012000030002332-0101022033023133-2100010111301231-2312010203030222-2210300331223022"></a>

## Next pages — default_coalescing / 222103130312 / 4

- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.coalescing_options](resources--workload--reference--group-009.md#canonical-2300311201121333-1311112110203331-3310212300212011-2313230012130122-1332203113120320-3020302213011022-2331223222320311-0303223313003001)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-3101101001110123-3200031233033322-3101313011322200-2120330033233331-1002233131222001-0321331321022132-1122231221330110-0310103033212031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1132112330203002-1233232022000233-0300132023111212-2230101332031203-1131300320301010-0110222332212303-2210223023121131-3001022213120330"></a>

## service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.coalescing_options.strict_coalescing — strict_coalescing / 121130230132 / 2

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
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.coalescing_options](resources--workload--reference--group-009.md#canonical-2300311201121333-1311112110203331-3310212300212011-2313230012130122-1332203113120320-3020302213011022-2331223222320311-0303223313003001)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.coalescing_options.strict_coalescing

<a id="canonical-2332112001133111-2302002030133202-2233010233102222-1320102321130021-2331113001233201-3302003130233020-1312213001103233-2232213303132011"></a>

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

<a id="canonical-0031220010330312-0102112012213313-2100212121103311-3222322032130123-0221000013222221-0120301011311120-0322100130221003-3302212003010300"></a>

## Direct properties — strict_coalescing / 121130230132 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0233202103212011-1112310211201110-2033313203332213-0021101202311103-2021100021320023-1230110131212223-3122033223131200-2311001210032322"></a>

## Next pages — strict_coalescing / 121130230132 / 4

- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.coalescing_options](resources--workload--reference--group-009.md#canonical-2300311201121333-1311112110203331-3310212300212011-2313230012130122-1332203113120320-3020302213011022-2331223222320311-0303223313003001)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-0001220200312001-1113113220200311-2121232002032023-3311312103002000-0121010101202330-2111213223210002-2321330322111220-2331333113001213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3032211211133023-3231030320020023-0210331201013301-1220213303200232-1010113003321330-0101331213112203-0021312310131330-1203221333110212"></a>

## service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.default_header — default_header / 333022201101 / 2

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
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.default_header

<a id="canonical-3102330103210100-1201213022320013-0022310131201323-0301003303330002-2002331031001333-3032230322131201-2003013030230213-3113132223212313"></a>

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

<a id="canonical-1223002222232013-2111220131132233-3311113331003332-1100213201223012-2113312020212201-1010132131113021-1212131021311210-0111123203000133"></a>

## Direct properties — default_header / 333022201101 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2013020221202111-0010230213231131-3122321123303121-2310010131310130-1303131201333120-0302131011303001-1301230320323132-0332133003311311"></a>

## Next pages — default_header / 333022201101 / 4

- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](resources--workload--reference--group-009.md#canonical-3212310132100213-2000301220120030-0112323230310220-3200032133332210-3223310210320320-2331013310222010-2010102133130322-2111123310321223)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-0030131323133032-3230131331021321-0120331323111323-2311202313312123-2110223231212033-3320331321232223-0102313012203303-2311102313220320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2321022011133120-1212012131032012-2223132323231000-1213312200003212-1003330233012100-1311022212001132-1032331123311213-0011312310121303"></a>

## service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.default_loadbalancer — default_loadbalancer / 300022101000 / 2

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
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.default_loadbalancer

<a id="canonical-1010303023320330-0020320131212121-1131121101201101-1133122222102302-2030123001030222-1232121133021000-2311303222010020-0023321010123012"></a>

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

<a id="canonical-0232323010200012-3311002313301112-3122112002030122-1322121322100322-1112100212210203-3103030101223102-3100120211201023-0133130200233013"></a>

## Direct properties — default_loadbalancer / 300022101000 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2002121303113200-0030322321322320-3301333132000100-0333103011001330-0232200001323233-1120002022003031-0331323303102103-1210101311300102"></a>

## Next pages — default_loadbalancer / 300022101000 / 4

- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](resources--workload--reference--group-009.md#canonical-3212310132100213-2000301220120030-0112323230310220-3200032133332210-3223310210320320-2331013310222010-2010102133130322-2111123310321223)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-0100323020130132-1101221000003031-1021130013322030-0031130202003212-2332220321111301-0033022220301031-1033200212220030-3302131102131121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3333221210221011-1333031132021100-0321123001112020-1302010303110220-0110121103220321-3300313300231303-0231230331333112-2101100213123032"></a>

## service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.disable_path_normalize — disable_path_normalize / 320221131203 / 2

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
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.disable_path_normalize

<a id="canonical-0203022231323222-1112220200032331-3300021223210223-0321312110323130-1233011203112122-3100031033320002-2210021312222200-1130103013000010"></a>

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

<a id="canonical-3233132330222203-1021222120103200-0013322103332122-2120210222013321-0100233321103132-3112110210322020-2033020302113222-0322020301210031"></a>

## Direct properties — disable_path_normalize / 320221131203 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1011320102120301-2302110321230320-3032332013312301-1133121112023100-0213000111030222-3211003122123302-1022311301202221-1101213132303033"></a>

## Next pages — disable_path_normalize / 320221131203 / 4

- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](resources--workload--reference--group-009.md#canonical-3212310132100213-2000301220120030-0112323230310220-3200032133332210-3223310210320320-2331013310222010-2010102133130322-2111123310321223)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-1233033122010133-0313233003202220-3123310102130302-0030010200322301-1302221332333123-3300232123111230-0220102032030330-3101210201002321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0032310203311102-0100201321331232-1303302002201330-3223211333322213-1213321330301023-0330233001321201-0022101111120021-2132113222313231"></a>

## service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.enable_path_normalize — enable_path_normalize / 321302320100 / 2

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
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.enable_path_normalize

<a id="canonical-0023023010020213-1023331030011022-1222233322231311-3111230323132022-3113022101300313-0313301310121321-3002310011122002-0201103131311230"></a>

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

<a id="canonical-3022133200100203-1213011300033223-1121003210132202-0122210100103112-0313332300000203-1011122000122312-2203031020021013-1233223122111210"></a>

## Direct properties — enable_path_normalize / 321302320100 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2300322001311232-3333311102002100-1001123333212133-1230233311133033-3020133200201321-2113200032301201-1123230133103221-2320221020312130"></a>

## Next pages — enable_path_normalize / 321302320100 / 4

- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](resources--workload--reference--group-009.md#canonical-3212310132100213-2000301220120030-0112323230310220-3200032133332210-3223310210320320-2331013310222010-2010102133130322-2111123310321223)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-2233110201213030-3333133212100020-2011121111212120-0113211322030320-0002310120033130-3322203130310011-3012120201131022-3120322301300011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3313121300022021-2011030032000032-2030330013212311-2231230200100210-3203102001002112-3322123310233303-0121113000101020-3103130301123111"></a>

## service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options — http_protocol_options / 332223000313 / 2

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
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options

<a id="canonical-2001023211212122-3221111121102020-2332133221123301-0232003221210330-2033303213001110-2110023132120331-1001023221013213-0110303321013112"></a>

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

<a id="canonical-2021131123021313-0022123121011030-0220133301211222-1101120231133113-0031301222330001-0220331011032000-0121032312120133-0113011200010300"></a>

## Direct properties — http_protocol_options / 332223000313 / 3

- [http_protocol_enable_v1_only](resources--workload--reference--group-009.md#canonical-0012002132023001-3102332203330311-2031320232332313-3133033320123023-0311213000321130-1212202333233322-1321220102022202-2200223332220132): complete subsection reference.

- [http_protocol_enable_v1_v2](resources--workload--reference--group-009.md#canonical-2012133020100113-0113330111002031-2102300212021110-1203310101012003-1010321133233223-0022221313310211-3122022111202331-1311120331111232): complete subsection reference.

- [http_protocol_enable_v2_only](resources--workload--reference--group-009.md#canonical-2332302311023210-2101323320333312-0211130201231333-0130102032000003-1021110113031000-1320203320033021-3011301022231203-3011013311030101): complete subsection reference.

<a id="canonical-2122301021203300-3210113002333020-1032211302211122-3220221023221031-2303113133012321-1203323111220030-0001300232231013-1113000222323313"></a>

## Next pages — http_protocol_options / 332223000313 / 4

- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only](resources--workload--reference--group-009.md#canonical-0012002132023001-3102332203330311-2031320232332313-3133033320123023-0311213000321130-1212202333233322-1321220102022202-2200223332220132)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_v2](resources--workload--reference--group-009.md#canonical-2012133020100113-0113330111002031-2102300212021110-1203310101012003-1010321133233223-0022221313310211-3122022111202331-1311120331111232)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v2_only](resources--workload--reference--group-009.md#canonical-2332302311023210-2101323320333312-0211130201231333-0130102032000003-1021110113031000-1320203320033021-3011301022231203-3011013311030101)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](resources--workload--reference--group-009.md#canonical-3212310132100213-2000301220120030-0112323230310220-3200032133332210-3223310210320320-2331013310222010-2010102133130322-2111123310321223)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-0012002132023001-3102332203330311-2031320232332313-3133033320123023-0311213000321130-1212202333233322-1321220102022202-2200223332220132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3001211300320322-2123301212323303-3000131200330320-2233201303303333-1021120033331023-3030312101123133-3100121030221202-2030120012013320"></a>

## service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only — http_protocol_enable_v1_only / 233230202212 / 2

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
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options](resources--workload--reference--group-009.md#canonical-2233110201213030-3333133212100020-2011121111212120-0113211322030320-0002310120033130-3322203130310011-3012120201131022-3120322301300011)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only

<a id="canonical-1311331032302211-1003202220120132-2133021301301132-1202321301010330-0313312212103133-0012011213003020-1020231131212200-1012022010010111"></a>

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

<a id="canonical-3001123323100021-2322032221232003-0301121111033302-1330012233230001-2113313130201221-3121330133223222-2210233220200001-1301233313203010"></a>

## Direct properties — http_protocol_enable_v1_only / 233230202212 / 3

- [header_transformation](resources--workload--reference--group-009.md#canonical-2131010203232011-1010223103203021-2200322331110300-1000221121131001-3030013212001012-2231211222131123-2032010120032012-2122220230333301): complete subsection reference.

<a id="canonical-1101221110120211-1202311321111231-3000023030212113-2101102133121001-0322320001302032-1223331012320101-2321310312201221-1033201030101302"></a>

## Next pages — http_protocol_enable_v1_only / 233230202212 / 4

- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation](resources--workload--reference--group-009.md#canonical-2131010203232011-1010223103203021-2200322331110300-1000221121131001-3030013212001012-2231211222131123-2032010120032012-2122220230333301)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options](resources--workload--reference--group-009.md#canonical-2233110201213030-3333133212100020-2011121111212120-0113211322030320-0002310120033130-3322203130310011-3012120201131022-3120322301300011)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-2131010203232011-1010223103203021-2200322331110300-1000221121131001-3030013212001012-2231211222131123-2032010120032012-2122220230333301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3330100121131021-1303302321103120-0120321013312203-2000321113130002-2210223012301113-0323100213110013-1000123302122012-2300130002132232"></a>

## service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation — header_transformation / 321200021230 / 2

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
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options](resources--workload--reference--group-009.md#canonical-2233110201213030-3333133212100020-2011121111212120-0113211322030320-0002310120033130-3322203130310011-3012120201131022-3120322301300011)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only](resources--workload--reference--group-009.md#canonical-0012002132023001-3102332203330311-2031320232332313-3133033320123023-0311213000321130-1212202333233322-1321220102022202-2200223332220132)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation

<a id="canonical-0321002301111302-3030231012120330-2103313301003111-0112032311031301-3213123220300002-0032301231022232-0330033300231120-2212201201013221"></a>

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

<a id="canonical-0232123301230130-3332103131100212-0222303010133313-3122321221331302-2121113020320330-3012320123023131-2311323022320121-0100332032000113"></a>

## Direct properties — header_transformation / 321200021230 / 3

- [default_header_transformation](resources--workload--reference--group-009.md#canonical-2033033123111213-1201102110100221-0133112202032113-1222000213030313-0033233202213203-2031232211232201-0331013332021011-3321322110201221): complete subsection reference.

- [preserve_case_header_transformation](resources--workload--reference--group-009.md#canonical-3110302020331233-1010323211011201-0221331003303313-1101103312210013-0231003021332000-3010312033120212-0320213103212303-2102023031003332): complete subsection reference.

- [proper_case_header_transformation](resources--workload--reference--group-009.md#canonical-1012233112010230-1121023211323330-2113213231311301-2012120232100223-2233312022332113-0230033111000020-2323132013222013-0100222010212133): complete subsection reference.

<a id="canonical-1213220133020000-0030031202101132-1101100200210202-3102331212101120-0320120023222301-1331112203323211-0003313133213030-0221230032223122"></a>

## Next pages — header_transformation / 321200021230 / 4

- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.default_header_transformation](resources--workload--reference--group-009.md#canonical-2033033123111213-1201102110100221-0133112202032113-1222000213030313-0033233202213203-2031232211232201-0331013332021011-3321322110201221)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.preserve_case_header_transformation](resources--workload--reference--group-009.md#canonical-3110302020331233-1010323211011201-0221331003303313-1101103312210013-0231003021332000-3010312033120212-0320213103212303-2102023031003332)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.proper_case_header_transformation](resources--workload--reference--group-009.md#canonical-1012233112010230-1121023211323330-2113213231311301-2012120232100223-2233312022332113-0230033111000020-2323132013222013-0100222010212133)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only](resources--workload--reference--group-009.md#canonical-0012002132023001-3102332203330311-2031320232332313-3133033320123023-0311213000321130-1212202333233322-1321220102022202-2200223332220132)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-2033033123111213-1201102110100221-0133112202032113-1222000213030313-0033233202213203-2031232211232201-0331013332021011-3321322110201221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3132103121303212-1131222211313002-1103322212322333-2320103223331113-1102031002033333-0111230212033022-1233020001333013-2232300202310132"></a>

## service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.default_header_transformation — default_header_transformation / 121130123203 / 2

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
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options](resources--workload--reference--group-009.md#canonical-2233110201213030-3333133212100020-2011121111212120-0113211322030320-0002310120033130-3322203130310011-3012120201131022-3120322301300011)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only](resources--workload--reference--group-009.md#canonical-0012002132023001-3102332203330311-2031320232332313-3133033320123023-0311213000321130-1212202333233322-1321220102022202-2200223332220132)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation](resources--workload--reference--group-009.md#canonical-2131010203232011-1010223103203021-2200322331110300-1000221121131001-3030013212001012-2231211222131123-2032010120032012-2122220230333301)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.default_header_transformation

<a id="canonical-1113310322232301-1323110103212212-3031323223222112-0101321120001122-2111320331130220-2313310313330202-1201100131231211-3113310221322211"></a>

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

<a id="canonical-3112021322033100-2013012030303121-2322032020233122-3223231220030223-0030222110333103-2101033333031221-0231113101121002-2112302030331013"></a>

## Direct properties — default_header_transformation / 121130123203 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0001010221331321-3003330032330012-1210033311130212-1021222002121003-3323022101033301-2233112331131021-2013112001012320-3312010030003223"></a>

## Next pages — default_header_transformation / 121130123203 / 4

- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation](resources--workload--reference--group-009.md#canonical-2131010203232011-1010223103203021-2200322331110300-1000221121131001-3030013212001012-2231211222131123-2032010120032012-2122220230333301)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-3110302020331233-1010323211011201-0221331003303313-1101103312210013-0231003021332000-3010312033120212-0320213103212303-2102023031003332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3112022101300031-2001010223130302-0303200221223022-0111302312102033-3130123003132311-3213323031320332-1001313232221031-0030323102030122"></a>

## service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.preserve_case_header_transformation — preserve_case_header_transformation / 130001110021 / 2

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
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options](resources--workload--reference--group-009.md#canonical-2233110201213030-3333133212100020-2011121111212120-0113211322030320-0002310120033130-3322203130310011-3012120201131022-3120322301300011)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only](resources--workload--reference--group-009.md#canonical-0012002132023001-3102332203330311-2031320232332313-3133033320123023-0311213000321130-1212202333233322-1321220102022202-2200223332220132)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation](resources--workload--reference--group-009.md#canonical-2131010203232011-1010223103203021-2200322331110300-1000221121131001-3030013212001012-2231211222131123-2032010120032012-2122220230333301)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.preserve_case_header_transformation

<a id="canonical-1323312022100200-2030123002002321-0000231231232132-3213312222332102-1203212023301221-2012133023213321-2203333321213023-0023210221133200"></a>

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

<a id="canonical-3131033332321222-0222120330121210-2030021123013123-1020131113000033-0221011101210232-2213212110333031-1203300310221033-0321110130222011"></a>

## Direct properties — preserve_case_header_transformation / 130001110021 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3031010221102102-2311131231120002-3232020031300322-0233210000003023-3203103120023303-2333131313103132-3110210300210132-0000210110230133"></a>

## Next pages — preserve_case_header_transformation / 130001110021 / 4

- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation](resources--workload--reference--group-009.md#canonical-2131010203232011-1010223103203021-2200322331110300-1000221121131001-3030013212001012-2231211222131123-2032010120032012-2122220230333301)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-1012233112010230-1121023211323330-2113213231311301-2012120232100223-2233312022332113-0230033111000020-2323132013222013-0100222010212133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1323103132013101-2031322013303101-1113031033333122-2003121130023233-3030302022122020-3013233211333103-1211121000203303-3231213013313321"></a>

## service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.proper_case_header_transformation — proper_case_header_transformation / 332332231111 / 2

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
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options](resources--workload--reference--group-009.md#canonical-2233110201213030-3333133212100020-2011121111212120-0113211322030320-0002310120033130-3322203130310011-3012120201131022-3120322301300011)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only](resources--workload--reference--group-009.md#canonical-0012002132023001-3102332203330311-2031320232332313-3133033320123023-0311213000321130-1212202333233322-1321220102022202-2200223332220132)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation](resources--workload--reference--group-009.md#canonical-2131010203232011-1010223103203021-2200322331110300-1000221121131001-3030013212001012-2231211222131123-2032010120032012-2122220230333301)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.proper_case_header_transformation

<a id="canonical-1000011110333213-1032000020032110-2003223111101031-3123332322212012-2001110012102331-3201131103121022-2330322231013232-3230011231122023"></a>

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

<a id="canonical-2321310133202032-2222313312220233-1200301303112332-0203132120321022-0303310112331112-0132003122020012-3332023000202212-0300121303002010"></a>

## Direct properties — proper_case_header_transformation / 332332231111 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0122023303031233-3203030021012112-1312210300311110-0123022122030123-3202301013103221-3111320022013221-0301332201230001-2230231031030203"></a>

## Next pages — proper_case_header_transformation / 332332231111 / 4

- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation](resources--workload--reference--group-009.md#canonical-2131010203232011-1010223103203021-2200322331110300-1000221121131001-3030013212001012-2231211222131123-2032010120032012-2122220230333301)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-2012133020100113-0113330111002031-2102300212021110-1203310101012003-1010321133233223-0022221313310211-3122022111202331-1311120331111232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2233000111111113-0230323012331331-1322011112232233-2230031010221111-1013000303333212-1222021321220332-2001323222112311-3211320322020211"></a>

## service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_v2 — http_protocol_enable_v1_v2 / 303201300121 / 2

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
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options](resources--workload--reference--group-009.md#canonical-2233110201213030-3333133212100020-2011121111212120-0113211322030320-0002310120033130-3322203130310011-3012120201131022-3120322301300011)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_v2

<a id="canonical-0110320031303031-3231120110213321-1122130230233113-3033121333032223-3322120221323301-0332302110220122-0221133113012213-1331101233020101"></a>

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

<a id="canonical-3131332313132300-0222101110123213-0021122111231133-1222103033030001-3133201232103103-0201220221210220-2000022023023230-3300002123232223"></a>

## Direct properties — http_protocol_enable_v1_v2 / 303201300121 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2133112212013321-3320003112230022-3332130203111210-3032332020021121-3212022303310012-2001131300012211-3331112333103133-1032030210213010"></a>

## Next pages — http_protocol_enable_v1_v2 / 303201300121 / 4

- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options](resources--workload--reference--group-009.md#canonical-2233110201213030-3333133212100020-2011121111212120-0113211322030320-0002310120033130-3322203130310011-3012120201131022-3120322301300011)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-2332302311023210-2101323320333312-0211130201231333-0130102032000003-1021110113031000-1320203320033021-3011301022231203-3011013311030101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0203231212312020-0312013223233130-2132110130103201-0021203303103032-2211202203200202-1222130101002102-1202112110032201-0203200202231002"></a>

## service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v2_only — http_protocol_enable_v2_only / 131323123023 / 2

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
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options](resources--workload--reference--group-009.md#canonical-2233110201213030-3333133212100020-2011121111212120-0113211322030320-0002310120033130-3322203130310011-3012120201131022-3120322301300011)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v2_only

<a id="canonical-3001111110123331-3232032313110033-0001311130111330-0322133322030322-0213112322030331-0123230102100300-2302122232120321-3233332111111133"></a>

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

<a id="canonical-1302133230200221-3203231200222122-3201331030111300-1231130131332320-3300132331123232-3233010202333010-2013122220012001-2313120130123221"></a>

## Direct properties — http_protocol_enable_v2_only / 131323123023 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2211223300101002-1220333123323012-3230333110102021-0332201222212212-3303031110301323-2211210112212020-2102223232210110-0030111210211212"></a>

## Next pages — http_protocol_enable_v2_only / 131323123023 / 4

- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options](resources--workload--reference--group-009.md#canonical-2233110201213030-3333133212100020-2011121111212120-0113211322030320-0002310120033130-3322203130310011-3012120201131022-3120322301300011)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-3033302113113321-1000032130332122-0133011213100323-3323120132112130-1222100231201310-3230202123223130-2103032130300101-1013121302101213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1022311023100212-3132322033312203-3120002200200323-0031100020331301-1321033121303111-3013003123320020-0131020323211021-2003222331321213"></a>

## service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.non_default_loadbalancer — non_default_loadbalancer / 310212323200 / 2

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
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.non_default_loadbalancer

<a id="canonical-1100002130020000-3100100133013022-3310030112322201-0333322231311011-1212230013131120-3322113323133002-3310232000023131-0022333320221320"></a>

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

<a id="canonical-2223130113032002-0012102221030302-2033222203002223-2330202322201030-0303121103121231-1320203200020000-2213030112012121-0232330032321000"></a>

## Direct properties — non_default_loadbalancer / 310212323200 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2302113230012231-1130033112133013-0113031121213223-0223121221301112-2023013111101321-1120200131313022-2023011212223221-2001030002231000"></a>

## Next pages — non_default_loadbalancer / 310212323200 / 4

- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](resources--workload--reference--group-009.md#canonical-3212310132100213-2000301220120030-0112323230310220-3200032133332210-3223310210320320-2331013310222010-2010102133130322-2111123310321223)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-0203331121231312-0110013320313123-1233303021130123-3331222120302133-1332231120213020-3203030031311331-2322322120023130-0002030203333232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0212223022302321-1102003223130110-0010302230203231-0010121121131012-1003001011013210-1010223212111113-2300133333312321-0022132323231033"></a>

## service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.pass_through — pass_through / 300300132233 / 2

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
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.pass_through

<a id="canonical-3312121121202313-3021102003300101-0031003120323223-3200212320221201-0202232031023122-2221000302011133-2123321012200013-2112320303333303"></a>

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

<a id="canonical-2130102023023211-1130110001103200-3001212120223103-2310012031033211-0223112010232231-0231120021100331-1331312211021321-0223033012113330"></a>

## Direct properties — pass_through / 300300132233 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2021112110102000-0002220130331321-1020023322123001-1301132132130230-1332222033102112-2323020311231230-1332002032131202-1120003123010000"></a>

## Next pages — pass_through / 300300132233 / 4

- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](resources--workload--reference--group-009.md#canonical-3212310132100213-2000301220120030-0112323230310220-3200032133332210-3223310210320320-2331013310222010-2010102133130322-2111123310321223)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-0212122113133333-1000102320023221-2003032220233320-0021231111201312-3001031012012332-3313221303003312-0112011332100001-1120030302022032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2000010102213200-2321220202233011-0102020220230213-2212011112202110-0120331221200023-0203021023211332-2312111230312313-1210203130322133"></a>

## service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params — tls_cert_params / 303233101113 / 2

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
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params

<a id="canonical-0203231220123023-0131210320013100-2010302310111002-0332203130120111-1220221311330203-3220221232110132-2310031322131010-2010312100332302"></a>

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

<a id="canonical-3003032320333333-1333222322322122-0222031233002312-1012012230132202-3230100030010220-3213300002222231-2330131122122010-3203220130023211"></a>

## Direct properties — tls_cert_params / 303233101113 / 3

- [certificates](resources--workload--reference--group-009.md#canonical-0300303023201031-3212301123330130-3230033101113121-2220123322220331-2203113130321221-3213330030001121-1230230032210300-3111032223221233): complete subsection reference.

- [no_mtls](resources--workload--reference--group-009.md#canonical-2321102223232133-3122103122030313-3000312222131221-0330000230100223-1112312313113121-1022232301212322-1110312032102333-3110000312101002): complete subsection reference.

- [tls_config](resources--workload--reference--group-009.md#canonical-0233232311020320-2311233031001121-2321212112032013-2311011011131031-3331200111012210-2210033000032303-2231212213212012-3021103133101231): complete subsection reference.

- [use_mtls](resources--workload--reference--group-009.md#canonical-2230323133110331-2101302313123301-2033301333010232-3320123100210000-1220000110330331-3301011020212033-1110022310322130-3132201012030322): complete subsection reference.

<a id="canonical-0130100332223121-1010010010202133-3222103030230022-3001023021030230-0032233201311132-0221103200211011-2202102220330323-1312001021332231"></a>

## Next pages — tls_cert_params / 303233101113 / 4

- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.certificates](resources--workload--reference--group-009.md#canonical-0300303023201031-3212301123330130-3230033101113121-2220123322220331-2203113130321221-3213330030001121-1230230032210300-3111032223221233)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.no_mtls](resources--workload--reference--group-009.md#canonical-2321102223232133-3122103122030313-3000312222131221-0330000230100223-1112312313113121-1022232301212322-1110312032102333-3110000312101002)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.tls_config](resources--workload--reference--group-009.md#canonical-0233232311020320-2311233031001121-2321212112032013-2311011011131031-3331200111012210-2210033000032303-2231212213212012-3021103133101231)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.use_mtls](resources--workload--reference--group-009.md#canonical-2230323133110331-2101302313123301-2033301333010232-3320123100210000-1220000110330331-3301011020212033-1110022310322130-3132201012030322)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](resources--workload--reference--group-009.md#canonical-3212310132100213-2000301220120030-0112323230310220-3200032133332210-3223310210320320-2331013310222010-2010102133130322-2111123310321223)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-0300303023201031-3212301123330130-3230033101113121-2220123322220331-2203113130321221-3213330030001121-1230230032210300-3111032223221233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0133112113023120-2310321331130010-0332103011311313-0313030312031211-1021103112210223-1002120310103123-3110330033113132-2302233302300031"></a>

## service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.certificates — certificates / 120301322103 / 2

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
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.certificates

<a id="canonical-0221003312113312-3002001323133322-1133122100230200-0333130303010220-3032122110112212-1033100212221303-1303232123013112-3231301120330232"></a>

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

<a id="canonical-1022220213122200-2023110131330103-0333231113232321-2001000222233100-1002211000201321-3202000200130101-2022302311002232-0300203000103101"></a>

## Direct properties — certificates / 120301322103 / 3

<a id="canonical-2032203010221333-0023132223002322-2123030210113222-3121202201121131-3211031300321002-2211022111331020-1313013220111230-3031330213310123"></a>

<a id="canonical-2123013021230332-0203300100312101-1102303010112332-3322111001003012-2323133113232133-0101301301221021-1132211010223000-3013223301023020"></a>

## name property — certificates / 120301322103 / 4

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

<a id="canonical-1111233330212032-2301332222033112-1323300011333023-3123132203320000-2002020201011103-0033012133302133-1100233202312021-0123113203033302"></a>

<a id="canonical-1223032013333121-1333233203002100-1333231200213002-1023110103331113-2101222323012103-0211221033122332-2113020021103320-3120330110033010"></a>

## namespace property — certificates / 120301322103 / 5

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

<a id="canonical-2322303312100231-3033121033011033-0210130300203123-0310123200100322-1100211313113202-0231131130230021-2210112302110123-0311233003333102"></a>

<a id="canonical-2220301233102301-1022302101002302-2231210130122023-2331300233333030-0213100212301031-1302321233211212-3120211103102031-0220131123130302"></a>

## tenant property — certificates / 120301322103 / 6

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

<a id="canonical-0101222023300102-2322230222033113-2323023112302122-2032021110020332-1212133302112212-0110230322020013-2120102322233231-3210001331010211"></a>

## Next pages — certificates / 120301322103 / 7

- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params](resources--workload--reference--group-009.md#canonical-0212122113133333-1000102320023221-2003032220233320-0021231111201312-3001031012012332-3313221303003312-0112011332100001-1120030302022032)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-2321102223232133-3122103122030313-3000312222131221-0330000230100223-1112312313113121-1022232301212322-1110312032102333-3110000312101002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3203123231222300-1320031231230302-1000301210212203-3200123101103033-0111022112213013-0132111223031233-2110303233002030-3211233323320122"></a>

## service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.no_mtls — no_mtls / 313112103231 / 2

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
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.no_mtls

<a id="canonical-2101113311022121-3012213131030113-2131031103311033-2032132030302110-3313032011330322-2222010133023000-2003110301310211-1112221120330210"></a>

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

<a id="canonical-1331331232323301-3101122113202020-0132220013313103-0110232232223213-0322200302013221-2102013101303210-1220201030202321-0023233003201310"></a>

## Direct properties — no_mtls / 313112103231 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0113012033322311-1113300323121111-1011301031121122-2001321010121101-2213300011223200-0233211332000211-3013003010013130-3311322212021002"></a>

## Next pages — no_mtls / 313112103231 / 4

- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params](resources--workload--reference--group-009.md#canonical-0212122113133333-1000102320023221-2003032220233320-0021231111201312-3001031012012332-3313221303003312-0112011332100001-1120030302022032)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-0233232311020320-2311233031001121-2321212112032013-2311011011131031-3331200111012210-2210033000032303-2231212213212012-3021103133101231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3033013001132203-0022313230230002-1231022311313303-2022312012112203-1300210321331321-2133112320202212-0211300220033120-0102012120203231"></a>

## service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.tls_config — tls_config / 000013202130 / 2

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
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.tls_config

<a id="canonical-3330210102001122-0121030132111101-0103012123231201-1132311021003202-2312010110221020-2231323332031210-0130230021031130-2112013223213113"></a>

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

<a id="canonical-2213031132230112-2032310303311331-2103201022030032-1322201212320231-2333230232303330-0123012310001100-2011101023332012-1100233100203023"></a>

## Direct properties — tls_config / 000013202130 / 3

- [custom_security](resources--workload--reference--group-009.md#canonical-3020013013223000-2110120013030311-2212133321100033-2022220223011323-3010320103020332-1322211301303200-1110313121200201-2001011332312110): complete subsection reference.

- [default_security](resources--workload--reference--group-009.md#canonical-3030121303202200-1013020323202321-0102332110322111-2122112313202323-3111001100212002-2033221213301031-2302102011223320-2120300311032002): complete subsection reference.

- [low_security](resources--workload--reference--group-009.md#canonical-0330222201123101-1213033123311331-2322021100011103-3223323201203031-1200213201000222-0133010010203333-0020103200333230-1330302322313320): complete subsection reference.

- [medium_security](resources--workload--reference--group-009.md#canonical-0221122010013203-1101312112231103-1122201013012202-2113123321003232-3100110031322111-1210023032023330-3003232102310013-0133230222330330): complete subsection reference.

<a id="canonical-1233212131133132-2013100213312101-0020103310312312-0300003320021233-2203013000303021-0210332203330102-2303233012321212-2032113003220301"></a>

## Next pages — tls_config / 000013202130 / 4

- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.tls_config.custom_security](resources--workload--reference--group-009.md#canonical-3020013013223000-2110120013030311-2212133321100033-2022220223011323-3010320103020332-1322211301303200-1110313121200201-2001011332312110)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.tls_config.default_security](resources--workload--reference--group-009.md#canonical-3030121303202200-1013020323202321-0102332110322111-2122112313202323-3111001100212002-2033221213301031-2302102011223320-2120300311032002)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.tls_config.low_security](resources--workload--reference--group-009.md#canonical-0330222201123101-1213033123311331-2322021100011103-3223323201203031-1200213201000222-0133010010203333-0020103200333230-1330302322313320)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.tls_config.medium_security](resources--workload--reference--group-009.md#canonical-0221122010013203-1101312112231103-1122201013012202-2113123321003232-3100110031322111-1210023032023330-3003232102310013-0133230222330330)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params](resources--workload--reference--group-009.md#canonical-0212122113133333-1000102320023221-2003032220233320-0021231111201312-3001031012012332-3313221303003312-0112011332100001-1120030302022032)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-3020013013223000-2110120013030311-2212133321100033-2022220223011323-3010320103020332-1322211301303200-1110313121200201-2001011332312110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0011103102111110-3131220222232121-1113320222312300-1123112201303331-2203103202313011-2322210003213023-0303023102002220-2032220100112003"></a>

## service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.tls_config.custom_security — custom_security / 210103230313 / 2

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
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.tls_config](resources--workload--reference--group-009.md#canonical-0233232311020320-2311233031001121-2321212112032013-2311011011131031-3331200111012210-2210033000032303-2231212213212012-3021103133101231)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.tls_config.custom_security

<a id="canonical-1021212020022100-2103303022032332-1233103303020113-0333033301210210-3120233113013003-2320032221211022-1201101200310233-3123010320110211"></a>

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

<a id="canonical-1311202311132222-3301000220133311-3000303220123012-1333222130220032-3121232103330111-0223000032012112-2311232203220110-0133333023331131"></a>

## Direct properties — custom_security / 210103230313 / 3

<a id="canonical-2301022101232221-2320130213303113-2323313003011322-2210102313101223-0320020203013023-1030111231202223-0010232211012322-1031231312021130"></a>

<a id="canonical-0132112301202231-2113112210122311-2313210200323233-1302222022003223-3231020203221220-0312322100023233-3323310212231021-0002030020302300"></a>

## cipher_suites property — custom_security / 210103230313 / 4

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

<a id="canonical-0233021321321020-0200131023022320-1233033022002202-1022032133130231-0211031010130020-0112100313131301-0322213303111223-0110022111203022"></a>

<a id="canonical-1222033102313003-1321313133131323-3332203130032013-3203101022303031-2110100200100213-2312211230302003-1031021321131211-3121220222022223"></a>

## max_version property — custom_security / 210103230313 / 5

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

<a id="canonical-1112333112001000-3222122102232300-1031223133103102-3210110101203032-0103212223033330-2120222020110122-2130112322231022-2333122230031102"></a>

<a id="canonical-0322132122333002-3212111303012031-1312212213120312-1212230002022231-0010032301122100-2123231121001321-3023310312112302-2013022321120300"></a>

## min_version property — custom_security / 210103230313 / 6

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

<a id="canonical-1311020223322003-2222231023320333-2122202033133332-1023020030111330-1200021112230022-1221203123130212-1113233113322133-1123301211133320"></a>

## Next pages — custom_security / 210103230313 / 7

- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.tls_config](resources--workload--reference--group-009.md#canonical-0233232311020320-2311233031001121-2321212112032013-2311011011131031-3331200111012210-2210033000032303-2231212213212012-3021103133101231)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-3030121303202200-1013020323202321-0102332110322111-2122112313202323-3111001100212002-2033221213301031-2302102011223320-2120300311032002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3201313303023201-2223000331030331-0021231231112130-0331013001310032-3122322030231131-2212011331113021-3302203011221123-1120303301111222"></a>

## service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.tls_config.default_security — default_security / 313323022023 / 2

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
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.tls_config](resources--workload--reference--group-009.md#canonical-0233232311020320-2311233031001121-2321212112032013-2311011011131031-3331200111012210-2210033000032303-2231212213212012-3021103133101231)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.tls_config.default_security

<a id="canonical-3302320001133300-3200231101233233-0120222103203033-1101221320301230-3112013230010120-1201211310313203-3030031220231202-1002001232323100"></a>

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

<a id="canonical-0001011321003303-0213331102111120-1022120321203102-3300121223011133-2131303332213133-1100033312113112-0210303223131130-0212301101030122"></a>

## Direct properties — default_security / 313323022023 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3101103223331100-1022111232120310-2000232322212032-2333333013133223-2230212100133100-1332123231131213-2320230132113031-3322112000110132"></a>

## Next pages — default_security / 313323022023 / 4

- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.tls_config](resources--workload--reference--group-009.md#canonical-0233232311020320-2311233031001121-2321212112032013-2311011011131031-3331200111012210-2210033000032303-2231212213212012-3021103133101231)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-0330222201123101-1213033123311331-2322021100011103-3223323201203031-1200213201000222-0133010010203333-0020103200333230-1330302322313320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0221133112313222-1203102231203012-2033301012112220-2102011011303030-0130103110230220-0021023011332232-3122213021313230-1031302031322211"></a>

## service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.tls_config.low_security — low_security / 213013103200 / 2

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
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.tls_config](resources--workload--reference--group-009.md#canonical-0233232311020320-2311233031001121-2321212112032013-2311011011131031-3331200111012210-2210033000032303-2231212213212012-3021103133101231)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.tls_config.low_security

<a id="canonical-1033322111310312-3120222012232113-0212211300313000-3312003032123123-0233100312333203-0312022202021003-2332220221320312-2212100130003000"></a>

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

<a id="canonical-0111122100301223-2310300010110203-1031211211302311-0303110000323222-3322210320220021-3332120111113001-2222121332321233-3211001132201331"></a>

## Direct properties — low_security / 213013103200 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0132222021212003-3221111102211101-2321222212220022-2133033200203130-0201230332111022-2002001323310313-0023111233023102-0310312221132122"></a>

## Next pages — low_security / 213013103200 / 4

- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.tls_config](resources--workload--reference--group-009.md#canonical-0233232311020320-2311233031001121-2321212112032013-2311011011131031-3331200111012210-2210033000032303-2231212213212012-3021103133101231)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-0221122010013203-1101312112231103-1122201013012202-2113123321003232-3100110031322111-1210023032023330-3003232102310013-0133230222330330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3303131232323020-1301013131001113-1122232330231311-2022131123321202-1022311132332212-2312233030132300-1302333330002310-1202121101313113"></a>

## service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.tls_config.medium_security — medium_security / 120121212201 / 2

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
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.tls_config](resources--workload--reference--group-009.md#canonical-0233232311020320-2311233031001121-2321212112032013-2311011011131031-3331200111012210-2210033000032303-2231212213212012-3021103133101231)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.tls_config.medium_security

<a id="canonical-3221300321233321-2013123203121313-0130300103331331-1010310113023003-0032221201133323-0220021322302221-3120213330333300-2312222332023100"></a>

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

<a id="canonical-3023202100121122-0200022313012001-2103021123001102-2230131320131133-2020302312203102-1122201010320231-1101023132012130-1311021302331023"></a>

## Direct properties — medium_security / 120121212201 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2310321122001331-3132203213133300-2121031131031103-1333213332330332-1133211233202032-3221101132113121-1101303033101132-1101303030003020"></a>

## Next pages — medium_security / 120121212201 / 4

- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.tls_config](resources--workload--reference--group-009.md#canonical-0233232311020320-2311233031001121-2321212112032013-2311011011131031-3331200111012210-2210033000032303-2231212213212012-3021103133101231)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-2230323133110331-2101302313123301-2033301333010232-3320123100210000-1220000110330331-3301011020212033-1110022310322130-3132201012030322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2212122111201232-1201202322112011-0003130002230023-2100300210303232-1010010300003321-3322332311332310-3210003022113302-1013322111120321"></a>

## service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.use_mtls — use_mtls / 311330332103 / 2

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
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.use_mtls

<a id="canonical-3020310223313132-1220131033011120-3320103130002301-0212221330222331-3032030111201231-0303132222201002-0002001231233212-0210230212130230"></a>

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

<a id="canonical-2002213110130003-0202113310220112-1111230211100010-0122032101322331-2232330223320200-2002120013210121-1113130332231211-0311312030123133"></a>

## Direct properties — use_mtls / 311330332103 / 3

<a id="canonical-3123223231111100-2101231111102100-1030211012101103-1102012323000323-2312030233020332-1203002321311011-1200131211302331-1000221331113031"></a>

<a id="canonical-3223112232312332-1312123203131331-1023132333211032-2133311221231320-1010012031211223-1010012100010330-1232230312133313-3210122122022012"></a>

## client_certificate_optional property — use_mtls / 311330332103 / 4

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

- [crl](resources--workload--reference--group-009.md#canonical-1320000300021301-1003102121113101-3211330020132320-0212300201222232-2200133311332000-1330133110201130-0203323002103033-2102313032131103): complete subsection reference.

- [no_crl](resources--workload--reference--group-009.md#canonical-2312102003100320-3030200222022003-0230322030313121-0132210103012320-3120321210323130-1223312031332031-2301001023112013-3122131113002223): complete subsection reference.

- [trusted_ca](resources--workload--reference--group-009.md#canonical-3133110222330002-2303022310302013-0213311110032331-1220002033020310-2002101103301000-0332331020320203-1201313031011313-0231020301202211): complete subsection reference.

<a id="canonical-0220330201230303-1103121332313100-0232102321322011-3130232013221032-3011202223220313-3012331030132023-0030121121201303-2131221301232003"></a>

<a id="canonical-1201112322033221-1321121202310102-2121113102322120-2231210011323320-0333123121012201-0013202323102120-3111123233313011-0203222022323332"></a>

## trusted_ca_url property — use_mtls / 311330332103 / 5

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

- [xfcc_disabled](resources--workload--reference--group-010.md#canonical-0210002101312203-3033320003012002-3021222323013232-0223230020202010-0000203131132121-0231232312131022-1001001110320120-2033301233303132): complete subsection reference.

- [xfcc_options](resources--workload--reference--group-010.md#canonical-3331213103211210-0022333013121122-0010123223112003-0121131311131120-3120203200321230-0032102133022223-2031223033331220-3223022211203201): complete subsection reference.

<a id="canonical-0011002102301310-2221013103130232-2212312003223013-2123232122333233-2320330022223113-2001131022301010-2002321313301120-0200110021312220"></a>

## Next pages — use_mtls / 311330332103 / 6

- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.use_mtls.crl](resources--workload--reference--group-009.md#canonical-1320000300021301-1003102121113101-3211330020132320-0212300201222232-2200133311332000-1330133110201130-0203323002103033-2102313032131103)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.use_mtls.no_crl](resources--workload--reference--group-009.md#canonical-2312102003100320-3030200222022003-0230322030313121-0132210103012320-3120321210323130-1223312031332031-2301001023112013-3122131113002223)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.use_mtls.trusted_ca](resources--workload--reference--group-009.md#canonical-3133110222330002-2303022310302013-0213311110032331-1220002033020310-2002101103301000-0332331020320203-1201313031011313-0231020301202211)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.use_mtls.xfcc_disabled](resources--workload--reference--group-010.md#canonical-0210002101312203-3033320003012002-3021222323013232-0223230020202010-0000203131132121-0231232312131022-1001001110320120-2033301233303132)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.use_mtls.xfcc_options](resources--workload--reference--group-010.md#canonical-3331213103211210-0022333013121122-0010123223112003-0121131311131120-3120203200321230-0032102133022223-2031223033331220-3223022211203201)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params](resources--workload--reference--group-009.md#canonical-0212122113133333-1000102320023221-2003032220233320-0021231111201312-3001031012012332-3313221303003312-0112011332100001-1120030302022032)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-1320000300021301-1003102121113101-3211330020132320-0212300201222232-2200133311332000-1330133110201130-0203323002103033-2102313032131103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2303001320310320-2023110332333310-2021101100300113-0132033031122020-3002123013300101-1000202310121103-0110033212131010-1201130313022130"></a>

## service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.use_mtls.crl — crl / 330131213323 / 2

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
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.use_mtls.crl

<a id="canonical-1013232003311222-3223211220202002-3021122031331100-0021003020123332-0130222310022223-1022322202222200-2213133303131223-2133303233210322"></a>

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

<a id="canonical-2313113213202330-0213110222003303-1113013202102110-1232001322022213-1313131220330110-3323010021311223-2022011320121112-3223020132101332"></a>

## Direct properties — crl / 330131213323 / 3

<a id="canonical-2120210022321203-2123220310312330-2200033030032202-1211312333022032-0102112023111303-1020122121122122-3130230112333303-0122122232032331"></a>

<a id="canonical-2122220102301231-3312321032123122-0210012313020101-2221030301113123-2222311300103100-3123130223022322-3023012101312101-3223302012120230"></a>

## name property — crl / 330131213323 / 4

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

<a id="canonical-1001102232230133-1220133311130330-1111032021020322-2001232222322232-0312101301103321-3322331200323010-2300330112213100-0232001011232201"></a>

<a id="canonical-2002020010021121-3233213001110303-0322023322020313-0323312323213110-1112021221100131-2033233320201131-3230311303130012-1230311032121023"></a>

## namespace property — crl / 330131213323 / 5

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

<a id="canonical-1322021301100122-2312011010000321-1131132230231001-2122200031230300-1313133113331122-0002011101301232-1323133332210211-3122101232131233"></a>

<a id="canonical-1222021102012302-1302023002013330-1020110010322220-0012231002211301-2211301033000002-1323203331122303-2210301132031032-0021020133021310"></a>

## tenant property — crl / 330131213323 / 6

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

<a id="canonical-1202110020310023-1110302311113333-3132133010232230-0320330123312003-0233021132311230-1211322310103002-0313231312232132-1110033010322032"></a>

## Next pages — crl / 330131213323 / 7

- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.use_mtls](resources--workload--reference--group-009.md#canonical-2230323133110331-2101302313123301-2033301333010232-3320123100210000-1220000110330331-3301011020212033-1110022310322130-3132201012030322)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-2312102003100320-3030200222022003-0230322030313121-0132210103012320-3120321210323130-1223312031332031-2301001023112013-3122131113002223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0322212103123002-2330033212311112-1232133122103321-0102130300132233-2133230222201022-3131133302122233-3213323203032113-3221000110311013"></a>

## service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.use_mtls.no_crl — no_crl / 012101320230 / 2

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
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.use_mtls.no_crl

<a id="canonical-0323000002010120-1202102130300312-2131033330212331-2133031233011333-2022331110131100-1320002131002322-0210232022311010-1000202001132103"></a>

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

<a id="canonical-2000111023103320-1302210111312301-2030130020231011-2330210320121000-0200100111211021-2232233332032232-2030020231213232-2103033213220233"></a>

## Direct properties — no_crl / 012101320230 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1233201130003120-3303203101120333-2022220112132310-3221302013103013-3132321120002222-0222331322113311-1030222213021113-2201132230312310"></a>

## Next pages — no_crl / 012101320230 / 4

- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.use_mtls](resources--workload--reference--group-009.md#canonical-2230323133110331-2101302313123301-2033301333010232-3320123100210000-1220000110330331-3301011020212033-1110022310322130-3132201012030322)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-3133110222330002-2303022310302013-0213311110032331-1220002033020310-2002101103301000-0332331020320203-1201313031011313-0231020301202211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->
