---
page_title: "xcsh_workload reference"
subcategory: "Container"
description: "Complete grouped canonical reference for xcsh_workload reference."
---

# xcsh_workload reference

<a id="canonical-0011030022100032-1001121231022101-1312233123123212-1230321111212001-1202321021011311-2221131013321320-0230110120331322-0211202331210123"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v2_only — http_protocol_enable_v2_only / 111203212222 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-2121230312030330-3010113231323001-1233013032320322-1032203002210311-2022021002113232-1123312032112001-0012103213230102-1112311001110020)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-005.md#canonical-0133221101333210-0222103021001232-3212010032013103-3331120032012221-1012203003232110-1233102323001233-2233321121201101-0122313122210200)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-005.md#canonical-2311001323203002-2023333020201320-1330320322101233-1033123102030111-2013210002333120-2130122003123103-1133301323123220-2302130221012102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-005.md#canonical-2330031030101220-3210010323021133-3322231130101013-1101131102102332-3121022130002321-1233200100132332-0213220230130033-0110233032331011)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options](resources--workload--reference--group-005.md#canonical-1302010300230312-3111312201201113-2013133300112103-0112233323311000-0220233130302132-1102033102333132-1100222323112330-1222302133120000)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v2_only

<a id="canonical-1100032013111201-1122110323032123-2011200021300222-2331213002212120-2021310113102122-3213220202023332-2022201323131303-0012220033002330"></a>

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

<a id="canonical-1030222122013221-2131022030112011-1221221223302320-3113100030302123-3212133220010211-1300031302101122-1032001322111311-1003111321113023"></a>

## Direct properties — http_protocol_enable_v2_only / 111203212222 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3310220320230003-2100311231202303-2301112013102321-3201030033023120-2330302301303030-2132003102300230-3322002110102302-2233033000201023"></a>

## Next pages — http_protocol_enable_v2_only / 111203212222 / 4

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options](resources--workload--reference--group-005.md#canonical-1302010300230312-3111312201201113-2013133300112103-0112233323311000-0220233130302132-1102033102333132-1100222323112330-1222302133120000)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-3032202302032213-0303131002200121-0310313203203001-0303223311331203-3110132231230001-3300021313312233-0130110330103111-2132202022101200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2133203120302301-1123213111213313-2030113121232103-2031331120213213-0313302012120002-0221300002332310-0233201333002320-3201310232023300"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https.non_default_loadbalancer — non_default_loadbalancer / 201202101301 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-2121230312030330-3010113231323001-1233013032320322-1032203002210311-2022021002113232-1123312032112001-0012103213230102-1112311001110020)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-005.md#canonical-0133221101333210-0222103021001232-3212010032013103-3331120032012221-1012203003232110-1233102323001233-2233321121201101-0122313122210200)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-005.md#canonical-2311001323203002-2023333020201320-1330320322101233-1033123102030111-2013210002333120-2130122003123103-1133301323123220-2302130221012102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-005.md#canonical-2330031030101220-3210010323021133-3322231130101013-1101131102102332-3121022130002321-1233200100132332-0213220230130033-0110233032331011)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https.non_default_loadbalancer

<a id="canonical-3103032201033020-0203132310210233-2101101023202110-3200021203333030-1332101230021223-3301133001100200-1230323000020122-3030321302230000"></a>

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

<a id="canonical-1201221130122020-1021312223101010-3233332001233231-2130333323032113-3000220333311333-3210110211103000-3010311002221331-1103132212133132"></a>

## Direct properties — non_default_loadbalancer / 201202101301 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0313031213123223-2001021101002130-0000120031103311-2001121210223333-1313030003122031-1300000110232022-2033211211123231-1122221131032332"></a>

## Next pages — non_default_loadbalancer / 201202101301 / 4

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-005.md#canonical-2330031030101220-3210010323021133-3322231130101013-1101131102102332-3121022130002321-1233200100132332-0213220230130033-0110233032331011)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-1020131000000021-0213323132231310-3302312202012121-1100210321023123-0023331100011120-0223121031212023-2031020223031102-0220123120021331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1013302131100312-0313112311302013-3232023132233021-1210320331013112-3120010313020120-2103113112232321-3101033113223113-3102003323110113"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https.pass_through — pass_through / 132002020031 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-2121230312030330-3010113231323001-1233013032320322-1032203002210311-2022021002113232-1123312032112001-0012103213230102-1112311001110020)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-005.md#canonical-0133221101333210-0222103021001232-3212010032013103-3331120032012221-1012203003232110-1233102323001233-2233321121201101-0122313122210200)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-005.md#canonical-2311001323203002-2023333020201320-1330320322101233-1033123102030111-2013210002333120-2130122003123103-1133301323123220-2302130221012102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-005.md#canonical-2330031030101220-3210010323021133-3322231130101013-1101131102102332-3121022130002321-1233200100132332-0213220230130033-0110233032331011)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https.pass_through

<a id="canonical-0012000133033001-3202221220031131-2130103212230232-0003312310011302-0100110120230003-2003211332330322-0113221002033311-3022310322311102"></a>

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

<a id="canonical-0022221023322120-1032003230330002-3210113011033111-0222331203122033-0121211301221210-1320101200220100-3211121110102102-2023311002212222"></a>

## Direct properties — pass_through / 132002020031 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3312112111210131-3202021112100110-1201210300230002-1123020121203322-0003111032200121-3022010103313113-0100012301031012-3121003202011121"></a>

## Next pages — pass_through / 132002020031 / 4

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-005.md#canonical-2330031030101220-3210010323021133-3322231130101013-1101131102102332-3121022130002321-1233200100132332-0213220230130033-0110233032331011)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-0103133333202221-1131021113321323-2122021221232201-1013312322320320-2313333220100330-3011001333202103-0121120020303322-3332000000101010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1031031201233320-0101221101120100-3021112320221331-3210132010112032-0011121301121003-2210121121323121-1011121210000310-1003102333312212"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params — tls_cert_params / 103202222010 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-2121230312030330-3010113231323001-1233013032320322-1032203002210311-2022021002113232-1123312032112001-0012103213230102-1112311001110020)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-005.md#canonical-0133221101333210-0222103021001232-3212010032013103-3331120032012221-1012203003232110-1233102323001233-2233321121201101-0122313122210200)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-005.md#canonical-2311001323203002-2023333020201320-1330320322101233-1033123102030111-2013210002333120-2130122003123103-1133301323123220-2302130221012102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-005.md#canonical-2330031030101220-3210010323021133-3322231130101013-1101131102102332-3121022130002321-1233200100132332-0213220230130033-0110233032331011)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params

<a id="canonical-2200202311313013-3133100132100201-3130132133313203-0112330013112121-0122122232113111-0133202103101020-0113201332113211-1113311111212200"></a>

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

<a id="canonical-2120332221020100-1012003131121210-0203131021331103-3311113232222112-3222203130223223-3011032033122232-2311313312011021-3220021100303033"></a>

## Direct properties — tls_cert_params / 103202222010 / 3

- [certificates](resources--workload--reference--group-006.md#canonical-3223230020210003-0030232022001231-0012133222023300-1003220210131222-0101131033021303-0211011220220223-2000202120100222-2231202103122323): complete subsection reference.

- [no_mtls](resources--workload--reference--group-006.md#canonical-3302203010133000-2002030332212231-0300330233320313-3101023301132322-1133333000321303-0313312031033113-1222111302213030-0120213003332202): complete subsection reference.

- [tls_config](resources--workload--reference--group-006.md#canonical-1310201203210320-3212123002132201-3311120222033220-1211003021103131-0233322212003321-2231310022001131-2023122323232123-1320232213200223): complete subsection reference.

- [use_mtls](resources--workload--reference--group-006.md#canonical-1101333300120120-3120122223232231-1210100331033331-3010100230302312-0312021231303102-3112111210033200-3310010333223022-3120311223122211): complete subsection reference.

<a id="canonical-1232102010330303-0101021112330300-2011003130012000-0220222202132233-1210321030011011-3010222101000022-1233033123130033-1311003221333002"></a>

## Next pages — tls_cert_params / 103202222010 / 4

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.certificates](resources--workload--reference--group-006.md#canonical-3223230020210003-0030232022001231-0012133222023300-1003220210131222-0101131033021303-0211011220220223-2000202120100222-2231202103122323)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.no_mtls](resources--workload--reference--group-006.md#canonical-3302203010133000-2002030332212231-0300330233320313-3101023301132322-1133333000321303-0313312031033113-1222111302213030-0120213003332202)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.tls_config](resources--workload--reference--group-006.md#canonical-1310201203210320-3212123002132201-3311120222033220-1211003021103131-0233322212003321-2231310022001131-2023122323232123-1320232213200223)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.use_mtls](resources--workload--reference--group-006.md#canonical-1101333300120120-3120122223232231-1210100331033331-3010100230302312-0312021231303102-3112111210033200-3310010333223022-3120311223122211)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-005.md#canonical-2330031030101220-3210010323021133-3322231130101013-1101131102102332-3121022130002321-1233200100132332-0213220230130033-0110233032331011)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-3223230020210003-0030232022001231-0012133222023300-1003220210131222-0101131033021303-0211011220220223-2000202120100222-2231202103122323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1220321121222023-2130232020121231-0222100203203233-0001233300313101-1310033213030112-0212300211333332-1210230330301110-3210100212032223"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.certificates — certificates / 223321030121 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-2121230312030330-3010113231323001-1233013032320322-1032203002210311-2022021002113232-1123312032112001-0012103213230102-1112311001110020)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-005.md#canonical-0133221101333210-0222103021001232-3212010032013103-3331120032012221-1012203003232110-1233102323001233-2233321121201101-0122313122210200)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-005.md#canonical-2311001323203002-2023333020201320-1330320322101233-1033123102030111-2013210002333120-2130122003123103-1133301323123220-2302130221012102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-005.md#canonical-2330031030101220-3210010323021133-3322231130101013-1101131102102332-3121022130002321-1233200100132332-0213220230130033-0110233032331011)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params](resources--workload--reference--group-006.md#canonical-0103133333202221-1131021113321323-2122021221232201-1013312322320320-2313333220100330-3011001333202103-0121120020303322-3332000000101010)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.certificates

<a id="canonical-0101200221021110-3203213023100111-1121300031223331-2110030312203030-3212111231030303-3200313330310213-0330311011310212-2203200003131120"></a>

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

<a id="canonical-3013132100210021-3331111332023333-2303321030100302-3202312001211331-3031230011222202-1111203303102120-3032200200132002-1332233032130101"></a>

## Direct properties — certificates / 223321030121 / 3

<a id="canonical-0000101201001211-0122133102113012-2333112200231031-2000100201121100-3233012131231330-2332223021322331-1122302231233303-0210010232131233"></a>

<a id="canonical-1012332013210022-1232023301131202-1110030222010202-2100200000122022-2223131001222313-2320010012233012-0032210122012123-1110302323121323"></a>

## name property — certificates / 223321030121 / 4

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

<a id="canonical-1011032120111210-3200322320113220-0310213130301310-0333300121011221-0210012222302202-2103130001230202-0211013230103230-1311223002112001"></a>

<a id="canonical-2200032123023012-1030330301131301-2302312302223112-0113231300103320-0021233301112310-1201001310220103-2330112200230003-2131011103010030"></a>

## namespace property — certificates / 223321030121 / 5

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

<a id="canonical-1303010110300320-0220032233313230-3322132333202131-3220212011220010-3132013002232103-3310320030113013-2102203012132032-1000211320210021"></a>

<a id="canonical-3132121020211321-3313220012220322-1212132121220102-2000133232310023-3213313202332322-0020210322230202-3313121210021022-1033330013321310"></a>

## tenant property — certificates / 223321030121 / 6

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

<a id="canonical-2223020002310202-0232111331120330-3103000022121300-2112103032301313-0313212122022101-1321011103223220-2100211013123233-3002220332130023"></a>

## Next pages — certificates / 223321030121 / 7

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params](resources--workload--reference--group-006.md#canonical-0103133333202221-1131021113321323-2122021221232201-1013312322320320-2313333220100330-3011001333202103-0121120020303322-3332000000101010)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-3302203010133000-2002030332212231-0300330233320313-3101023301132322-1133333000321303-0313312031033113-1222111302213030-0120213003332202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3110222011023212-3000321330230203-0111320201310332-0131312022103032-3232331211330121-2032233231302203-2130032203122212-3133021102012311"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.no_mtls — no_mtls / 012102012333 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-2121230312030330-3010113231323001-1233013032320322-1032203002210311-2022021002113232-1123312032112001-0012103213230102-1112311001110020)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-005.md#canonical-0133221101333210-0222103021001232-3212010032013103-3331120032012221-1012203003232110-1233102323001233-2233321121201101-0122313122210200)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-005.md#canonical-2311001323203002-2023333020201320-1330320322101233-1033123102030111-2013210002333120-2130122003123103-1133301323123220-2302130221012102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-005.md#canonical-2330031030101220-3210010323021133-3322231130101013-1101131102102332-3121022130002321-1233200100132332-0213220230130033-0110233032331011)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params](resources--workload--reference--group-006.md#canonical-0103133333202221-1131021113321323-2122021221232201-1013312322320320-2313333220100330-3011001333202103-0121120020303322-3332000000101010)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.no_mtls

<a id="canonical-3332331232230320-1321023103101302-0103111003032212-0210022200003102-3323301210102132-0111320130321132-0233323322302231-1022210232000210"></a>

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

<a id="canonical-3130112312031130-0213220300232123-2210002000020332-1011320223003303-1110301323310030-3300032131111021-3200102200010103-3331103131100213"></a>

## Direct properties — no_mtls / 012102012333 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2321210211313003-0030122131130232-3313120101131323-0230220302232220-1100120030101013-1112100321002102-1122022331100301-1112113111222310"></a>

## Next pages — no_mtls / 012102012333 / 4

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params](resources--workload--reference--group-006.md#canonical-0103133333202221-1131021113321323-2122021221232201-1013312322320320-2313333220100330-3011001333202103-0121120020303322-3332000000101010)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-1310201203210320-3212123002132201-3311120222033220-1211003021103131-0233322212003321-2231310022001131-2023122323232123-1320232213200223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2021301333312130-3312313333230223-2202103212222023-0312100110330031-2032103030032032-2301031021031300-0133110112003233-2211133231113013"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.tls_config — tls_config / 231003030121 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-2121230312030330-3010113231323001-1233013032320322-1032203002210311-2022021002113232-1123312032112001-0012103213230102-1112311001110020)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-005.md#canonical-0133221101333210-0222103021001232-3212010032013103-3331120032012221-1012203003232110-1233102323001233-2233321121201101-0122313122210200)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-005.md#canonical-2311001323203002-2023333020201320-1330320322101233-1033123102030111-2013210002333120-2130122003123103-1133301323123220-2302130221012102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-005.md#canonical-2330031030101220-3210010323021133-3322231130101013-1101131102102332-3121022130002321-1233200100132332-0213220230130033-0110233032331011)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params](resources--workload--reference--group-006.md#canonical-0103133333202221-1131021113321323-2122021221232201-1013312322320320-2313333220100330-3011001333202103-0121120020303322-3332000000101010)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.tls_config

<a id="canonical-0202020303211030-1213312230011103-3210313112321222-3213001000131320-1311000010223033-0301013202020121-0010231203030211-0030120113011112"></a>

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

<a id="canonical-0331211230200011-0013132101302012-2011223102212311-0031001120322022-0121232001032031-2301321201311122-2221133320312333-1203023300012031"></a>

## Direct properties — tls_config / 231003030121 / 3

- [custom_security](resources--workload--reference--group-006.md#canonical-3002103320031123-2303221012231133-2213130122010133-1233020223023131-1333133003213303-2003021232033113-0213201000022122-2332103021111010): complete subsection reference.

- [default_security](resources--workload--reference--group-006.md#canonical-2113103022200220-2221201030103030-0101303320200030-1310021212011321-0223013312300202-0231021322232232-3131202122200031-2101103220100211): complete subsection reference.

- [low_security](resources--workload--reference--group-006.md#canonical-1313320312301230-1321010323120200-0002332120223322-2133010302312133-3312210332020100-2211121311010221-2323003003120221-1322330302012021): complete subsection reference.

- [medium_security](resources--workload--reference--group-006.md#canonical-1203201320130323-1302320132132110-0032113120132211-0323100112223112-0002323201220000-1122103001023320-2200233202221232-2001321112110112): complete subsection reference.

<a id="canonical-0321303323331201-2322032113310112-2110031020132012-3133330103012221-1211303013021212-0302310310020202-2210032223223221-0212010030310030"></a>

## Next pages — tls_config / 231003030121 / 4

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.tls_config.custom_security](resources--workload--reference--group-006.md#canonical-3002103320031123-2303221012231133-2213130122010133-1233020223023131-1333133003213303-2003021232033113-0213201000022122-2332103021111010)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.tls_config.default_security](resources--workload--reference--group-006.md#canonical-2113103022200220-2221201030103030-0101303320200030-1310021212011321-0223013312300202-0231021322232232-3131202122200031-2101103220100211)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.tls_config.low_security](resources--workload--reference--group-006.md#canonical-1313320312301230-1321010323120200-0002332120223322-2133010302312133-3312210332020100-2211121311010221-2323003003120221-1322330302012021)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.tls_config.medium_security](resources--workload--reference--group-006.md#canonical-1203201320130323-1302320132132110-0032113120132211-0323100112223112-0002323201220000-1122103001023320-2200233202221232-2001321112110112)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params](resources--workload--reference--group-006.md#canonical-0103133333202221-1131021113321323-2122021221232201-1013312322320320-2313333220100330-3011001333202103-0121120020303322-3332000000101010)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-3002103320031123-2303221012231133-2213130122010133-1233020223023131-1333133003213303-2003021232033113-0213201000022122-2332103021111010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0030213211331120-2211010332213232-3030232113031122-2320300213330301-2023021212210223-2010001003023232-3103002231322130-1112001033231312"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.tls_config.custom_security — custom_security / 220210202122 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-2121230312030330-3010113231323001-1233013032320322-1032203002210311-2022021002113232-1123312032112001-0012103213230102-1112311001110020)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-005.md#canonical-0133221101333210-0222103021001232-3212010032013103-3331120032012221-1012203003232110-1233102323001233-2233321121201101-0122313122210200)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-005.md#canonical-2311001323203002-2023333020201320-1330320322101233-1033123102030111-2013210002333120-2130122003123103-1133301323123220-2302130221012102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-005.md#canonical-2330031030101220-3210010323021133-3322231130101013-1101131102102332-3121022130002321-1233200100132332-0213220230130033-0110233032331011)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params](resources--workload--reference--group-006.md#canonical-0103133333202221-1131021113321323-2122021221232201-1013312322320320-2313333220100330-3011001333202103-0121120020303322-3332000000101010)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.tls_config](resources--workload--reference--group-006.md#canonical-1310201203210320-3212123002132201-3311120222033220-1211003021103131-0233322212003321-2231310022001131-2023122323232123-1320232213200223)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.tls_config.custom_security

<a id="canonical-2123333012222132-1202312102300033-3032023003303002-0323230323130230-2012030212033233-2311123330313110-3222332012112321-1332331100110303"></a>

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

<a id="canonical-1120301013000312-0101322102230201-2103132322230203-0033020033311231-0120131103132210-1320220200302022-1023110021003322-1123310210032101"></a>

## Direct properties — custom_security / 220210202122 / 3

<a id="canonical-1230330121022333-1233222011200301-2103132303010203-1302233021230101-1030030001322330-3322122031120101-1330312012202231-0322303010332133"></a>

<a id="canonical-0130323032332311-1003100103030010-2322102130022200-2202211322013032-2020303200313201-0133002021012020-2000300130202023-0201113103121033"></a>

## cipher_suites property — custom_security / 220210202122 / 4

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

<a id="canonical-0000333001131031-0133212022331013-0023323202313102-0033100310032330-3100202031003330-3312102123201110-1020100200000121-2020130130300322"></a>

<a id="canonical-2332020013030330-3202021100122122-2102030200112111-1011213000301122-3111230110220113-2020102011033210-2200031122203001-3000211013300012"></a>

## max_version property — custom_security / 220210202122 / 5

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

<a id="canonical-1233321303331210-0101203320100012-1231121301332022-3012231023012113-1002001032313231-0331112101030231-2031333102302003-2122031103211213"></a>

<a id="canonical-1221130122012131-1220001321210331-0233303233230310-1021222013101113-0332231101311030-3000022120323003-3331203122012223-2020322111130210"></a>

## min_version property — custom_security / 220210202122 / 6

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

<a id="canonical-2233310011133111-2322223332032103-2221221303202103-3031323021220103-3332332322232102-2211112203133033-1221001322300133-2222310210303133"></a>

## Next pages — custom_security / 220210202122 / 7

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.tls_config](resources--workload--reference--group-006.md#canonical-1310201203210320-3212123002132201-3311120222033220-1211003021103131-0233322212003321-2231310022001131-2023122323232123-1320232213200223)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-2113103022200220-2221201030103030-0101303320200030-1310021212011321-0223013312300202-0231021322232232-3131202122200031-2101103220100211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0333101120222112-2003332232321303-2031023300222321-3311311011030330-2210223222302220-0022101331102330-2201211302133311-3121112201303123"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.tls_config.default_security — default_security / 310222023312 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-2121230312030330-3010113231323001-1233013032320322-1032203002210311-2022021002113232-1123312032112001-0012103213230102-1112311001110020)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-005.md#canonical-0133221101333210-0222103021001232-3212010032013103-3331120032012221-1012203003232110-1233102323001233-2233321121201101-0122313122210200)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-005.md#canonical-2311001323203002-2023333020201320-1330320322101233-1033123102030111-2013210002333120-2130122003123103-1133301323123220-2302130221012102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-005.md#canonical-2330031030101220-3210010323021133-3322231130101013-1101131102102332-3121022130002321-1233200100132332-0213220230130033-0110233032331011)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params](resources--workload--reference--group-006.md#canonical-0103133333202221-1131021113321323-2122021221232201-1013312322320320-2313333220100330-3011001333202103-0121120020303322-3332000000101010)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.tls_config](resources--workload--reference--group-006.md#canonical-1310201203210320-3212123002132201-3311120222033220-1211003021103131-0233322212003321-2231310022001131-2023122323232123-1320232213200223)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.tls_config.default_security

<a id="canonical-2330120020203132-1101323232300113-2210021233013123-2333000321111112-1213301332113211-0333111202321301-1333300200033230-3030133230312011"></a>

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

<a id="canonical-2123110313201320-0101103223231110-2101232203203333-3320032223333230-3201233200331111-1001302322300312-0111212200112121-3013210320101031"></a>

## Direct properties — default_security / 310222023312 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1101202233221131-0311322212123322-0203130011212131-0231222221123120-1012333002330033-0311132212230132-2100301213112322-3000101231212322"></a>

## Next pages — default_security / 310222023312 / 4

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.tls_config](resources--workload--reference--group-006.md#canonical-1310201203210320-3212123002132201-3311120222033220-1211003021103131-0233322212003321-2231310022001131-2023122323232123-1320232213200223)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-1313320312301230-1321010323120200-0002332120223322-2133010302312133-3312210332020100-2211121311010221-2323003003120221-1322330302012021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3213000132110133-2211022022210102-3323232003200013-1021201212331010-2120010020132213-0313010300322012-1123201102102020-3121210111100232"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.tls_config.low_security — low_security / 322203000321 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-2121230312030330-3010113231323001-1233013032320322-1032203002210311-2022021002113232-1123312032112001-0012103213230102-1112311001110020)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-005.md#canonical-0133221101333210-0222103021001232-3212010032013103-3331120032012221-1012203003232110-1233102323001233-2233321121201101-0122313122210200)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-005.md#canonical-2311001323203002-2023333020201320-1330320322101233-1033123102030111-2013210002333120-2130122003123103-1133301323123220-2302130221012102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-005.md#canonical-2330031030101220-3210010323021133-3322231130101013-1101131102102332-3121022130002321-1233200100132332-0213220230130033-0110233032331011)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params](resources--workload--reference--group-006.md#canonical-0103133333202221-1131021113321323-2122021221232201-1013312322320320-2313333220100330-3011001333202103-0121120020303322-3332000000101010)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.tls_config](resources--workload--reference--group-006.md#canonical-1310201203210320-3212123002132201-3311120222033220-1211003021103131-0233322212003321-2231310022001131-2023122323232123-1320232213200223)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.tls_config.low_security

<a id="canonical-3122233020332300-1011323220102121-3213212323131332-2333212110211303-3203132100121101-2323211300311023-3200123030222131-0021311131330021"></a>

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

<a id="canonical-0210010012101320-2021012303022303-0310323223300211-0022230221122103-3113032023100222-2031021131031302-0021021203233230-3131311311312121"></a>

## Direct properties — low_security / 322203000321 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1333030213210220-2203112113300220-1203103132211331-2110000131201100-3321323131201221-0313113202312030-3133012332122223-3001312201123113"></a>

## Next pages — low_security / 322203000321 / 4

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.tls_config](resources--workload--reference--group-006.md#canonical-1310201203210320-3212123002132201-3311120222033220-1211003021103131-0233322212003321-2231310022001131-2023122323232123-1320232213200223)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-1203201320130323-1302320132132110-0032113120132211-0323100112223112-0002323201220000-1122103001023320-2200233202221232-2001321112110112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0032202011021013-1103313103032302-2210011021330100-0010112030023020-2302231103212311-3122013311302313-3301320221202220-3331330303221202"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.tls_config.medium_security — medium_security / 102223330231 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-2121230312030330-3010113231323001-1233013032320322-1032203002210311-2022021002113232-1123312032112001-0012103213230102-1112311001110020)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-005.md#canonical-0133221101333210-0222103021001232-3212010032013103-3331120032012221-1012203003232110-1233102323001233-2233321121201101-0122313122210200)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-005.md#canonical-2311001323203002-2023333020201320-1330320322101233-1033123102030111-2013210002333120-2130122003123103-1133301323123220-2302130221012102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-005.md#canonical-2330031030101220-3210010323021133-3322231130101013-1101131102102332-3121022130002321-1233200100132332-0213220230130033-0110233032331011)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params](resources--workload--reference--group-006.md#canonical-0103133333202221-1131021113321323-2122021221232201-1013312322320320-2313333220100330-3011001333202103-0121120020303322-3332000000101010)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.tls_config](resources--workload--reference--group-006.md#canonical-1310201203210320-3212123002132201-3311120222033220-1211003021103131-0233322212003321-2231310022001131-2023122323232123-1320232213200223)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.tls_config.medium_security

<a id="canonical-3200132312110323-0001300123213120-0202213001322210-0123033133100021-3232313132011133-1230131312021120-0130300310110202-1101130201121211"></a>

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

<a id="canonical-1311330101112031-2320300003211312-3110230320133330-2331030323213131-3303003132031023-3211032331001223-1023023302102223-0112022211223003"></a>

## Direct properties — medium_security / 102223330231 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0213201311200202-1013310121012312-3113013302232032-0332202103100221-2232331000230201-0003022232121201-3303112233013301-1001300320013111"></a>

## Next pages — medium_security / 102223330231 / 4

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.tls_config](resources--workload--reference--group-006.md#canonical-1310201203210320-3212123002132201-3311120222033220-1211003021103131-0233322212003321-2231310022001131-2023122323232123-1320232213200223)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-1101333300120120-3120122223232231-1210100331033331-3010100230302312-0312021231303102-3112111210033200-3310010333223022-3120311223122211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1320101222200032-1133311121232122-3333202321110230-1320310102132310-2313233213202011-1132113210000322-1012123331022103-2021233022321312"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.use_mtls — use_mtls / 131320321333 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-2121230312030330-3010113231323001-1233013032320322-1032203002210311-2022021002113232-1123312032112001-0012103213230102-1112311001110020)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-005.md#canonical-0133221101333210-0222103021001232-3212010032013103-3331120032012221-1012203003232110-1233102323001233-2233321121201101-0122313122210200)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-005.md#canonical-2311001323203002-2023333020201320-1330320322101233-1033123102030111-2013210002333120-2130122003123103-1133301323123220-2302130221012102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-005.md#canonical-2330031030101220-3210010323021133-3322231130101013-1101131102102332-3121022130002321-1233200100132332-0213220230130033-0110233032331011)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params](resources--workload--reference--group-006.md#canonical-0103133333202221-1131021113321323-2122021221232201-1013312322320320-2313333220100330-3011001333202103-0121120020303322-3332000000101010)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.use_mtls

<a id="canonical-3301111132132111-1131322222100133-0100230010002201-0103112322233200-0131300202121123-2023232122012102-2333220220333311-2220023120011223"></a>

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

<a id="canonical-1123312202312100-2321333003021010-2330303321302230-0000233131112020-3022001332332103-1213322111011330-1132330323013210-3030102220331032"></a>

## Direct properties — use_mtls / 131320321333 / 3

<a id="canonical-3223323321033131-3110331101202113-2312232311100131-3003003331231320-1302121320202012-3021002202003100-1111110311320012-1212023312121133"></a>

<a id="canonical-2302321231320213-2331100020111112-0320313012003112-1110003021021311-3120203302201002-2322132100213303-0321123302001210-2310211132331311"></a>

## client_certificate_optional property — use_mtls / 131320321333 / 4

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

- [crl](resources--workload--reference--group-006.md#canonical-3330320011212133-0132212322323321-3111033020332330-3120021223213203-1130322223011331-0120222032003120-0021111220013002-2110132100001201): complete subsection reference.

- [no_crl](resources--workload--reference--group-006.md#canonical-3132232133022013-0021101022010033-1023113222130002-1023230121303133-1030320221221321-3030110332130323-2110003122131202-2231113210212233): complete subsection reference.

- [trusted_ca](resources--workload--reference--group-006.md#canonical-2102323330320121-0322003333331211-1312032013333003-3202230030023233-3131111310133110-0000321133122312-3303301110000232-2013002001331313): complete subsection reference.

<a id="canonical-0032032011222211-2100232033030233-2001122321201333-1210231033120002-0110033111322013-2100322110022012-3012233010332123-3332010021202012"></a>

<a id="canonical-3130302012213100-0032323100100031-2311131013012310-3331012101130123-2312233112302221-2200002211123300-2333110103201012-3010332303233123"></a>

## trusted_ca_url property — use_mtls / 131320321333 / 5

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

- [xfcc_disabled](resources--workload--reference--group-006.md#canonical-2333232120322321-2130320212013102-0032111131233131-2021102130103132-1021322121001131-3211301012132310-1233002102010132-1200020200232221): complete subsection reference.

- [xfcc_options](resources--workload--reference--group-006.md#canonical-0002132221020332-1110113000033220-3131330132113303-0331022100321200-1221302021011012-3300210130222002-3021131021010232-1031120312220302): complete subsection reference.

<a id="canonical-1202212203010322-3012103203330111-1030102031011230-0213232132031331-2101101232130020-2032232022303233-1100300120120210-0100330002211233"></a>

## Next pages — use_mtls / 131320321333 / 6

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.use_mtls.crl](resources--workload--reference--group-006.md#canonical-3330320011212133-0132212322323321-3111033020332330-3120021223213203-1130322223011331-0120222032003120-0021111220013002-2110132100001201)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.use_mtls.no_crl](resources--workload--reference--group-006.md#canonical-3132232133022013-0021101022010033-1023113222130002-1023230121303133-1030320221221321-3030110332130323-2110003122131202-2231113210212233)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.use_mtls.trusted_ca](resources--workload--reference--group-006.md#canonical-2102323330320121-0322003333331211-1312032013333003-3202230030023233-3131111310133110-0000321133122312-3303301110000232-2013002001331313)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.use_mtls.xfcc_disabled](resources--workload--reference--group-006.md#canonical-2333232120322321-2130320212013102-0032111131233131-2021102130103132-1021322121001131-3211301012132310-1233002102010132-1200020200232221)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.use_mtls.xfcc_options](resources--workload--reference--group-006.md#canonical-0002132221020332-1110113000033220-3131330132113303-0331022100321200-1221302021011012-3300210130222002-3021131021010232-1031120312220302)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params](resources--workload--reference--group-006.md#canonical-0103133333202221-1131021113321323-2122021221232201-1013312322320320-2313333220100330-3011001333202103-0121120020303322-3332000000101010)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-3330320011212133-0132212322323321-3111033020332330-3120021223213203-1130322223011331-0120222032003120-0021111220013002-2110132100001201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2012330100030110-1330013022233233-0221200123130323-3201112111110222-1021312023323022-0313133221021330-2232200120230302-1211130013003223"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.use_mtls.crl — crl / 303230201232 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-2121230312030330-3010113231323001-1233013032320322-1032203002210311-2022021002113232-1123312032112001-0012103213230102-1112311001110020)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-005.md#canonical-0133221101333210-0222103021001232-3212010032013103-3331120032012221-1012203003232110-1233102323001233-2233321121201101-0122313122210200)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-005.md#canonical-2311001323203002-2023333020201320-1330320322101233-1033123102030111-2013210002333120-2130122003123103-1133301323123220-2302130221012102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-005.md#canonical-2330031030101220-3210010323021133-3322231130101013-1101131102102332-3121022130002321-1233200100132332-0213220230130033-0110233032331011)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params](resources--workload--reference--group-006.md#canonical-0103133333202221-1131021113321323-2122021221232201-1013312322320320-2313333220100330-3011001333202103-0121120020303322-3332000000101010)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.use_mtls](resources--workload--reference--group-006.md#canonical-1101333300120120-3120122223232231-1210100331033331-3010100230302312-0312021231303102-3112111210033200-3310010333223022-3120311223122211)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.use_mtls.crl

<a id="canonical-1021320010031111-1111302122122003-2123201022322211-1302030003330121-1003301311303032-3003020221312010-2030121330013001-1130203231201121"></a>

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

<a id="canonical-0233203033002223-0030132332011002-2111231000003111-3221023222110113-3310103103312233-0110323222120221-0020223333103300-2301110132223211"></a>

## Direct properties — crl / 303230201232 / 3

<a id="canonical-1333032231131301-0022333331023030-1220212310001110-1212232033011210-3111013010230003-1133111011223131-1012110232131010-1013210331111312"></a>

<a id="canonical-1300331233231322-0011031110010021-2233103221112331-1312321130130220-1230210323213003-1013010103212202-1021112122230021-1202331321130310"></a>

## name property — crl / 303230201232 / 4

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

<a id="canonical-0200220011113023-2101110001311210-0300122100031333-1201122200113002-1300200100223010-1223021022122130-2330201000211202-0330323332223102"></a>

<a id="canonical-2030133122033310-2111111221103330-0113310031011302-1013101013222103-3120102033130302-3333202100230112-0322312033102322-2011011330001301"></a>

## namespace property — crl / 303230201232 / 5

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

<a id="canonical-0223330311203131-0113322012200000-0331100033200233-2123311023203030-1132002032031212-0332111202311221-2110313201313212-2210330133323302"></a>

<a id="canonical-2110202002322100-0303130300322122-0331310112203202-2031100212321101-1100112300333012-0333101132003131-0313332000121120-2211233212020132"></a>

## tenant property — crl / 303230201232 / 6

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

<a id="canonical-1131121002203121-2021202112202330-0003200321101122-1113013012023200-3021202010030011-2331131322032101-2223120103011131-1223013233032103"></a>

## Next pages — crl / 303230201232 / 7

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.use_mtls](resources--workload--reference--group-006.md#canonical-1101333300120120-3120122223232231-1210100331033331-3010100230302312-0312021231303102-3112111210033200-3310010333223022-3120311223122211)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-3132232133022013-0021101022010033-1023113222130002-1023230121303133-1030320221221321-3030110332130323-2110003122131202-2231113210212233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1321021301000110-1321113111102212-1320203321203013-2130123133022331-2021311122010023-3132201113000101-2313032330320322-2312233210002133"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.use_mtls.no_crl — no_crl / 202223230113 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-2121230312030330-3010113231323001-1233013032320322-1032203002210311-2022021002113232-1123312032112001-0012103213230102-1112311001110020)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-005.md#canonical-0133221101333210-0222103021001232-3212010032013103-3331120032012221-1012203003232110-1233102323001233-2233321121201101-0122313122210200)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-005.md#canonical-2311001323203002-2023333020201320-1330320322101233-1033123102030111-2013210002333120-2130122003123103-1133301323123220-2302130221012102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-005.md#canonical-2330031030101220-3210010323021133-3322231130101013-1101131102102332-3121022130002321-1233200100132332-0213220230130033-0110233032331011)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params](resources--workload--reference--group-006.md#canonical-0103133333202221-1131021113321323-2122021221232201-1013312322320320-2313333220100330-3011001333202103-0121120020303322-3332000000101010)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.use_mtls](resources--workload--reference--group-006.md#canonical-1101333300120120-3120122223232231-1210100331033331-3010100230302312-0312021231303102-3112111210033200-3310010333223022-3120311223122211)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.use_mtls.no_crl

<a id="canonical-1000111120300012-3100232222021002-1202211133301211-2300313021000310-1121203330203230-0122102203000111-1132220233002010-3203203313132110"></a>

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

<a id="canonical-0221123221122312-0030113122332211-2201023003221203-1100320121103313-1320222001222212-2330112023321233-0110203123312321-1333020112213313"></a>

## Direct properties — no_crl / 202223230113 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3123100011020300-3002001123200300-2030211012130233-1313211001103211-2012222010020213-0122111102022003-2133122112230130-0223030320230212"></a>

## Next pages — no_crl / 202223230113 / 4

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.use_mtls](resources--workload--reference--group-006.md#canonical-1101333300120120-3120122223232231-1210100331033331-3010100230302312-0312021231303102-3112111210033200-3310010333223022-3120311223122211)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-2102323330320121-0322003333331211-1312032013333003-3202230030023233-3131111310133110-0000321133122312-3303301110000232-2013002001331313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1003020300321132-1230020010020311-3221023003103232-3323030031102321-3323312323112221-3012231232020022-3322021013330012-2023331101212033"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.use_mtls.trusted_ca — trusted_ca / 320122233021 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-2121230312030330-3010113231323001-1233013032320322-1032203002210311-2022021002113232-1123312032112001-0012103213230102-1112311001110020)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-005.md#canonical-0133221101333210-0222103021001232-3212010032013103-3331120032012221-1012203003232110-1233102323001233-2233321121201101-0122313122210200)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-005.md#canonical-2311001323203002-2023333020201320-1330320322101233-1033123102030111-2013210002333120-2130122003123103-1133301323123220-2302130221012102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-005.md#canonical-2330031030101220-3210010323021133-3322231130101013-1101131102102332-3121022130002321-1233200100132332-0213220230130033-0110233032331011)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params](resources--workload--reference--group-006.md#canonical-0103133333202221-1131021113321323-2122021221232201-1013312322320320-2313333220100330-3011001333202103-0121120020303322-3332000000101010)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.use_mtls](resources--workload--reference--group-006.md#canonical-1101333300120120-3120122223232231-1210100331033331-3010100230302312-0312021231303102-3112111210033200-3310010333223022-3120311223122211)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.use_mtls.trusted_ca

<a id="canonical-1232030221300303-0133332310002011-1302320110311201-1201220323230230-0313300100222330-2302202311002133-3131030111332001-0333020031003111"></a>

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

<a id="canonical-0231100020202133-3231310133002110-2330332013121002-3003030302133202-2110300200301222-0120021111332023-2333021102313222-3112003130122201"></a>

## Direct properties — trusted_ca / 320122233021 / 3

<a id="canonical-3012312130000001-0202201232020112-0223113023030300-2001121320111030-3102022003102301-2311210020020302-2133003120020012-2103132303002300"></a>

<a id="canonical-1033102131203020-1132131111211023-1031302132031130-2211120131112102-1322200201201111-2131102302020123-1013103202001330-3222102201122320"></a>

## name property — trusted_ca / 320122233021 / 4

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

<a id="canonical-3103300333020212-0200003201220212-0303220101331002-2312121001201333-2311122031331131-0220311020121203-0110120011023131-3122030211011030"></a>

<a id="canonical-0333032332221021-2022331330301312-3330333222323330-2303022331320320-1013002130031202-2323132322222211-0222010322021201-2222213223312130"></a>

## namespace property — trusted_ca / 320122233021 / 5

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

<a id="canonical-0133130001101122-3323100013020112-3203220122031333-1032212022201113-0123233230022221-1030021003013010-0121122131211112-3332112232010312"></a>

<a id="canonical-3302313022320330-3333203220012233-1323113011211010-1132310322330221-0003031310122202-1211112303103112-3210331302302033-1331301200220212"></a>

## tenant property — trusted_ca / 320122233021 / 6

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

<a id="canonical-2333301211302013-1012021133300001-2101232230333122-0201001010002111-1032322002200322-3233100232123123-3312303110021122-2312303021332330"></a>

## Next pages — trusted_ca / 320122233021 / 7

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.use_mtls](resources--workload--reference--group-006.md#canonical-1101333300120120-3120122223232231-1210100331033331-3010100230302312-0312021231303102-3112111210033200-3310010333223022-3120311223122211)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-2333232120322321-2130320212013102-0032111131233131-2021102130103132-1021322121001131-3211301012132310-1233002102010132-1200020200232221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2233230121231132-3301032221130203-1011013020102322-1133122323011302-1132031220320330-0233113033101110-3232311223312022-1332030232123000"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.use_mtls.xfcc_disabled — xfcc_disabled / 101122120131 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-2121230312030330-3010113231323001-1233013032320322-1032203002210311-2022021002113232-1123312032112001-0012103213230102-1112311001110020)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-005.md#canonical-0133221101333210-0222103021001232-3212010032013103-3331120032012221-1012203003232110-1233102323001233-2233321121201101-0122313122210200)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-005.md#canonical-2311001323203002-2023333020201320-1330320322101233-1033123102030111-2013210002333120-2130122003123103-1133301323123220-2302130221012102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-005.md#canonical-2330031030101220-3210010323021133-3322231130101013-1101131102102332-3121022130002321-1233200100132332-0213220230130033-0110233032331011)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params](resources--workload--reference--group-006.md#canonical-0103133333202221-1131021113321323-2122021221232201-1013312322320320-2313333220100330-3011001333202103-0121120020303322-3332000000101010)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.use_mtls](resources--workload--reference--group-006.md#canonical-1101333300120120-3120122223232231-1210100331033331-3010100230302312-0312021231303102-3112111210033200-3310010333223022-3120311223122211)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.use_mtls.xfcc_disabled

<a id="canonical-0222132333331121-0010012010230031-1230201012133123-0000312031332032-2103111131130011-1231222313030010-2320300013033120-1130232303002220"></a>

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

<a id="canonical-0320200211112101-3011023101000110-0211121012123032-0300120021231032-3323100000033201-3320323222303303-1001321031210123-1030132031120312"></a>

## Direct properties — xfcc_disabled / 101122120131 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2021310112212322-1211031123132323-0311210031332223-1023022231213231-0011213033310120-0220012133230303-0113332010032223-0103312031211212"></a>

## Next pages — xfcc_disabled / 101122120131 / 4

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.use_mtls](resources--workload--reference--group-006.md#canonical-1101333300120120-3120122223232231-1210100331033331-3010100230302312-0312021231303102-3112111210033200-3310010333223022-3120311223122211)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-0002132221020332-1110113000033220-3131330132113303-0331022100321200-1221302021011012-3300210130222002-3021131021010232-1031120312220302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0322112322321302-0201021233231230-0032122330033332-0020203011333223-1311013223320210-1321331310203230-1210022111301132-0202210102022223"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.use_mtls.xfcc_options — xfcc_options / 101120121223 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-2121230312030330-3010113231323001-1233013032320322-1032203002210311-2022021002113232-1123312032112001-0012103213230102-1112311001110020)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-005.md#canonical-0133221101333210-0222103021001232-3212010032013103-3331120032012221-1012203003232110-1233102323001233-2233321121201101-0122313122210200)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-005.md#canonical-2311001323203002-2023333020201320-1330320322101233-1033123102030111-2013210002333120-2130122003123103-1133301323123220-2302130221012102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-005.md#canonical-2330031030101220-3210010323021133-3322231130101013-1101131102102332-3121022130002321-1233200100132332-0213220230130033-0110233032331011)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params](resources--workload--reference--group-006.md#canonical-0103133333202221-1131021113321323-2122021221232201-1013312322320320-2313333220100330-3011001333202103-0121120020303322-3332000000101010)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.use_mtls](resources--workload--reference--group-006.md#canonical-1101333300120120-3120122223232231-1210100331033331-3010100230302312-0312021231303102-3112111210033200-3310010333223022-3120311223122211)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.use_mtls.xfcc_options

<a id="canonical-2130133020111311-0303020021310112-2312333003023103-2233000222320312-0222310132323121-1333212110011201-0231301123233113-2013102221322320"></a>

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

<a id="canonical-2210310020223110-1120133111120032-0023032313023331-3110333221210032-1231333301010013-1022110123133010-3131012312303012-2012100303222011"></a>

## Direct properties — xfcc_options / 101120121223 / 3

<a id="canonical-2130130231311312-2222012232300120-2222011103313000-2210301102123113-2322301302300022-1333333321313313-2033211212303211-3102111033222122"></a>

<a id="canonical-1002333022133220-0102011211122002-3113320303300002-2311020301301211-0113312002131321-0231210332221023-2312031122232002-2122123130212123"></a>

## xfcc_header_elements property — xfcc_options / 101120121223 / 4

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

<a id="canonical-1300313213010233-0011110033212321-1112003301003021-0001023111122320-0331201133200111-2123010102212222-0102021003000110-1021311132300331"></a>

## Next pages — xfcc_options / 101120121223 / 5

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.use_mtls](resources--workload--reference--group-006.md#canonical-1101333300120120-3120122223232231-1210100331033331-3010100230302312-0312021231303102-3112111210033200-3310010333223022-3120311223122211)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-3023301020330331-1221233122321220-2011202113013003-2331003313323320-0122021000022233-0121202301001222-1011313113011301-2103023311112320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0103331231200020-2233312333121302-1303320103012322-2111130000323012-2001000130313200-3332331132011021-0101020102220201-0002221231000302"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters — tls_parameters / 112130000130 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-2121230312030330-3010113231323001-1233013032320322-1032203002210311-2022021002113232-1123312032112001-0012103213230102-1112311001110020)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-005.md#canonical-0133221101333210-0222103021001232-3212010032013103-3331120032012221-1012203003232110-1233102323001233-2233321121201101-0122313122210200)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-005.md#canonical-2311001323203002-2023333020201320-1330320322101233-1033123102030111-2013210002333120-2130122003123103-1133301323123220-2302130221012102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-005.md#canonical-2330031030101220-3210010323021133-3322231130101013-1101131102102332-3121022130002321-1233200100132332-0213220230130033-0110233032331011)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters

<a id="canonical-0303010130312230-3021223130110321-0021221232100311-2302032023332130-1302330232313002-1110301000333030-0113030001002110-2032302021021100"></a>

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

<a id="canonical-1312323032222031-2302022032013121-3322120021002232-2100320303200223-2323113302330123-1330223332002230-2332221011110330-1233201001010212"></a>

## Direct properties — tls_parameters / 112130000130 / 3

- [no_mtls](resources--workload--reference--group-006.md#canonical-1222103023212012-1003000121120210-3010300321130303-2020133030103020-0012133331000201-3033100121103313-3021202002302001-2031200021023212): complete subsection reference.

- [tls_certificates](resources--workload--reference--group-006.md#canonical-2023030202313320-2120223312023223-2011313021311023-1032020310121131-3203020102320011-3102231323232010-2320332011330312-2102320000023322): complete subsection reference.

- [tls_config](resources--workload--reference--group-006.md#canonical-0131213121022010-2302331123122311-0333232020222202-2110030111033001-1010012332232032-2230033302023213-0102002320332232-1002130120300231): complete subsection reference.

- [use_mtls](resources--workload--reference--group-006.md#canonical-1201123022002000-0122023033012312-0210013220112133-1220130111331213-0321022230013202-3213120123303001-1002300100311330-1023002312220123): complete subsection reference.

<a id="canonical-1210010333211300-3033200201230322-3123320330201300-0101213300332313-1313012210131312-3011222232110030-2111313010012123-1200211000011132"></a>

## Next pages — tls_parameters / 112130000130 / 4

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.no_mtls](resources--workload--reference--group-006.md#canonical-1222103023212012-1003000121120210-3010300321130303-2020133030103020-0012133331000201-3033100121103313-3021202002302001-2031200021023212)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates](resources--workload--reference--group-006.md#canonical-2023030202313320-2120223312023223-2011313021311023-1032020310121131-3203020102320011-3102231323232010-2320332011330312-2102320000023322)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_config](resources--workload--reference--group-006.md#canonical-0131213121022010-2302331123122311-0333232020222202-2110030111033001-1010012332232032-2230033302023213-0102002320332232-1002130120300231)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.use_mtls](resources--workload--reference--group-006.md#canonical-1201123022002000-0122023033012312-0210013220112133-1220130111331213-0321022230013202-3213120123303001-1002300100311330-1023002312220123)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-005.md#canonical-2330031030101220-3210010323021133-3322231130101013-1101131102102332-3121022130002321-1233200100132332-0213220230130033-0110233032331011)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-1222103023212012-1003000121120210-3010300321130303-2020133030103020-0012133331000201-3033100121103313-3021202002302001-2031200021023212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3202301210203211-3233021202223033-3001111330023320-0101212010112123-1231132212303001-0000110122313332-2121130220031323-0322201000230332"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.no_mtls — no_mtls / 233023322123 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-2121230312030330-3010113231323001-1233013032320322-1032203002210311-2022021002113232-1123312032112001-0012103213230102-1112311001110020)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-005.md#canonical-0133221101333210-0222103021001232-3212010032013103-3331120032012221-1012203003232110-1233102323001233-2233321121201101-0122313122210200)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-005.md#canonical-2311001323203002-2023333020201320-1330320322101233-1033123102030111-2013210002333120-2130122003123103-1133301323123220-2302130221012102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-005.md#canonical-2330031030101220-3210010323021133-3322231130101013-1101131102102332-3121022130002321-1233200100132332-0213220230130033-0110233032331011)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-006.md#canonical-3023301020330331-1221233122321220-2011202113013003-2331003313323320-0122021000022233-0121202301001222-1011313113011301-2103023311112320)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.no_mtls

<a id="canonical-1312321001311222-2303333000302303-2312120103110312-0222200011221232-2133123101010231-1010223221213211-1330232002321201-0010132032031221"></a>

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

<a id="canonical-2132130121120001-0303122022223101-0301313322300030-1333231210122202-1012302103111221-3010323111300323-2321033332231302-2113311131220221"></a>

## Direct properties — no_mtls / 233023322123 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3210201310132111-3303311302331203-0021331212121020-0111122231033330-3120101302020330-3232303132133011-2101002320202031-2103201320022201"></a>

## Next pages — no_mtls / 233023322123 / 4

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-006.md#canonical-3023301020330331-1221233122321220-2011202113013003-2331003313323320-0122021000022233-0121202301001222-1011313113011301-2103023311112320)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-2023030202313320-2120223312023223-2011313021311023-1032020310121131-3203020102320011-3102231323232010-2320332011330312-2102320000023322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0122111332203003-1133022203001132-2021131030301330-0003102230120132-1133321111323321-3223223130200020-1122012111121132-2203331011101322"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates — tls_certificates / 131013200021 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-2121230312030330-3010113231323001-1233013032320322-1032203002210311-2022021002113232-1123312032112001-0012103213230102-1112311001110020)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-005.md#canonical-0133221101333210-0222103021001232-3212010032013103-3331120032012221-1012203003232110-1233102323001233-2233321121201101-0122313122210200)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-005.md#canonical-2311001323203002-2023333020201320-1330320322101233-1033123102030111-2013210002333120-2130122003123103-1133301323123220-2302130221012102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-005.md#canonical-2330031030101220-3210010323021133-3322231130101013-1101131102102332-3121022130002321-1233200100132332-0213220230130033-0110233032331011)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-006.md#canonical-3023301020330331-1221233122321220-2011202113013003-2331003313323320-0122021000022233-0121202301001222-1011313113011301-2103023311112320)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates

<a id="canonical-1303222310312331-1213111002223313-3012320120012111-3002031221013202-1120012100003010-2313301232113010-3013132310003200-0022230211022323"></a>

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

<a id="canonical-3313200022332211-1022313030032120-0030001333123310-1133223320102211-3210221321102123-0111301031310313-0232232102213130-3103213020033010"></a>

## Direct properties — tls_certificates / 131013200021 / 3

<a id="canonical-1020103102133130-3300223001322123-0323123031022110-2122000202211320-3300303330333222-3132031123131332-2232332320130311-0200021210123122"></a>

<a id="canonical-3212002001012332-2323330221132220-1120320310030120-2230210001323010-3030010210112221-3102003233120103-3112023300131231-0111122032223020"></a>

## certificate_url property — tls_certificates / 131013200021 / 4

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

- [custom_hash_algorithms](resources--workload--reference--group-006.md#canonical-2231110321222212-0201111132112130-0013003032121100-2012323122331020-1223032133221000-1031230012022333-1331123330031200-2032123003310201): complete subsection reference.

<a id="canonical-3023210213100013-3011032201000131-1331203203110203-0302222100102032-0230331123023113-2102221201132202-2312130212212223-1220311230231300"></a>

<a id="canonical-1110013012010323-2102223301021321-0310202311222031-2011323332210121-3031003121022102-3331222200302322-3221201111310200-1302211303123110"></a>

## description_spec property — tls_certificates / 131013200021 / 5

Type: `"string"`. Optional.

Description. Description for the certificate.

- [disable_ocsp_stapling](resources--workload--reference--group-006.md#canonical-3330313200200023-1001120330111313-1100202000301111-3101101000000111-1231221223021002-0210132300321003-0112033022321102-3011223211131321): complete subsection reference.

- [private_key](resources--workload--reference--group-006.md#canonical-3031113033003222-0302203112110233-3013320030113033-2320330012301030-0323112210133122-1323133221322000-1301022232122020-1023332303223000): complete subsection reference.

- [use_system_defaults](resources--workload--reference--group-006.md#canonical-3130200102313022-1012211213023033-1232022333113120-0320033120300203-0012330030113331-1331132310230211-2022322233100320-3323012232012310): complete subsection reference.

<a id="canonical-2232311303210221-3030013200301001-2230002212313301-0010333203113111-2112021133130310-1010132013302300-1102320203003020-3232022020030321"></a>

## Next pages — tls_certificates / 131013200021 / 6

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates.custom_hash_algorithms](resources--workload--reference--group-006.md#canonical-2231110321222212-0201111132112130-0013003032121100-2012323122331020-1223032133221000-1031230012022333-1331123330031200-2032123003310201)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates.disable_ocsp_stapling](resources--workload--reference--group-006.md#canonical-3330313200200023-1001120330111313-1100202000301111-3101101000000111-1231221223021002-0210132300321003-0112033022321102-3011223211131321)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates.private_key](resources--workload--reference--group-006.md#canonical-3031113033003222-0302203112110233-3013320030113033-2320330012301030-0323112210133122-1323133221322000-1301022232122020-1023332303223000)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates.use_system_defaults](resources--workload--reference--group-006.md#canonical-3130200102313022-1012211213023033-1232022333113120-0320033120300203-0012330030113331-1331132310230211-2022322233100320-3323012232012310)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-006.md#canonical-3023301020330331-1221233122321220-2011202113013003-2331003313323320-0122021000022233-0121202301001222-1011313113011301-2103023311112320)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-2231110321222212-0201111132112130-0013003032121100-2012323122331020-1223032133221000-1031230012022333-1331123330031200-2032123003310201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2120303203123020-0110121000133132-1212110310223313-2032020010030120-0021012221002232-0303323313020123-2323333302000232-0222101010201012"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates.custom_hash_algorithms — custom_hash_algorithms / 331013000301 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-2121230312030330-3010113231323001-1233013032320322-1032203002210311-2022021002113232-1123312032112001-0012103213230102-1112311001110020)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-005.md#canonical-0133221101333210-0222103021001232-3212010032013103-3331120032012221-1012203003232110-1233102323001233-2233321121201101-0122313122210200)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-005.md#canonical-2311001323203002-2023333020201320-1330320322101233-1033123102030111-2013210002333120-2130122003123103-1133301323123220-2302130221012102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-005.md#canonical-2330031030101220-3210010323021133-3322231130101013-1101131102102332-3121022130002321-1233200100132332-0213220230130033-0110233032331011)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-006.md#canonical-3023301020330331-1221233122321220-2011202113013003-2331003313323320-0122021000022233-0121202301001222-1011313113011301-2103023311112320)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates](resources--workload--reference--group-006.md#canonical-2023030202313320-2120223312023223-2011313021311023-1032020310121131-3203020102320011-3102231323232010-2320332011330312-2102320000023322)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates.custom_hash_algorithms

<a id="canonical-3110233122021230-1023230203012113-2213003011302032-2213012301133122-0121032112000233-2323013302231232-0013330013331022-2132131202220023"></a>

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

<a id="canonical-0030030223113220-0210211102100011-0000030313210212-1000012111112012-1102332032031111-2112120101001001-0302001002101002-1321020202300130"></a>

## Direct properties — custom_hash_algorithms / 331013000301 / 3

<a id="canonical-3231200230011312-0303101230120330-0133022312213233-0133021032231031-0332001020233201-3322010212231202-0233001121213000-0321020201322322"></a>

<a id="canonical-1022131000103031-1133130323122111-1001020221013111-2330101321220210-3203322220003111-2230100133232220-3122331230133111-2200032201122111"></a>

## hash_algorithms property — custom_hash_algorithms / 331013000301 / 4

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

<a id="canonical-1113220321222110-3032001331021313-0002201100302323-1313321110132322-0103201231111033-1131333223211023-0030031020000310-0011112310302212"></a>

## Next pages — custom_hash_algorithms / 331013000301 / 5

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates](resources--workload--reference--group-006.md#canonical-2023030202313320-2120223312023223-2011313021311023-1032020310121131-3203020102320011-3102231323232010-2320332011330312-2102320000023322)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-3330313200200023-1001120330111313-1100202000301111-3101101000000111-1231221223021002-0210132300321003-0112033022321102-3011223211131321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2232002313221111-3032010031133332-1100122233310210-1101102111233312-3330122311013331-1310220310030013-3120021013320110-0121332312123233"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates.disable_ocsp_stapling — disable_ocsp_stapling / 031210220313 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-2121230312030330-3010113231323001-1233013032320322-1032203002210311-2022021002113232-1123312032112001-0012103213230102-1112311001110020)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-005.md#canonical-0133221101333210-0222103021001232-3212010032013103-3331120032012221-1012203003232110-1233102323001233-2233321121201101-0122313122210200)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-005.md#canonical-2311001323203002-2023333020201320-1330320322101233-1033123102030111-2013210002333120-2130122003123103-1133301323123220-2302130221012102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-005.md#canonical-2330031030101220-3210010323021133-3322231130101013-1101131102102332-3121022130002321-1233200100132332-0213220230130033-0110233032331011)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-006.md#canonical-3023301020330331-1221233122321220-2011202113013003-2331003313323320-0122021000022233-0121202301001222-1011313113011301-2103023311112320)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates](resources--workload--reference--group-006.md#canonical-2023030202313320-2120223312023223-2011313021311023-1032020310121131-3203020102320011-3102231323232010-2320332011330312-2102320000023322)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates.disable_ocsp_stapling

<a id="canonical-1232102002022032-2021232312323333-3213000012102220-2231332002220320-3120002010313001-1120323111120100-0303101131312321-3301100110102221"></a>

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

<a id="canonical-0312201033110121-3230101111221022-2121122230013033-1210030013131310-0332213213130210-0303200303300101-0331132231133302-3032223023023032"></a>

## Direct properties — disable_ocsp_stapling / 031210220313 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2023123101123222-0031222203301213-0131310331000122-2333230133211301-2320112312221111-2021121101313220-1203033330123033-0133021100310011"></a>

## Next pages — disable_ocsp_stapling / 031210220313 / 4

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates](resources--workload--reference--group-006.md#canonical-2023030202313320-2120223312023223-2011313021311023-1032020310121131-3203020102320011-3102231323232010-2320332011330312-2102320000023322)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-3031113033003222-0302203112110233-3013320030113033-2320330012301030-0323112210133122-1323133221322000-1301022232122020-1023332303223000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1300312320112010-3113310031210101-1032111110102221-3232111131131311-2020213100200010-3033221030032132-3301210210323001-1012301023112020"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates.private_key — private_key / 001332131332 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-2121230312030330-3010113231323001-1233013032320322-1032203002210311-2022021002113232-1123312032112001-0012103213230102-1112311001110020)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-005.md#canonical-0133221101333210-0222103021001232-3212010032013103-3331120032012221-1012203003232110-1233102323001233-2233321121201101-0122313122210200)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-005.md#canonical-2311001323203002-2023333020201320-1330320322101233-1033123102030111-2013210002333120-2130122003123103-1133301323123220-2302130221012102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-005.md#canonical-2330031030101220-3210010323021133-3322231130101013-1101131102102332-3121022130002321-1233200100132332-0213220230130033-0110233032331011)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-006.md#canonical-3023301020330331-1221233122321220-2011202113013003-2331003313323320-0122021000022233-0121202301001222-1011313113011301-2103023311112320)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates](resources--workload--reference--group-006.md#canonical-2023030202313320-2120223312023223-2011313021311023-1032020310121131-3203020102320011-3102231323232010-2320332011330312-2102320000023322)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates.private_key

<a id="canonical-3032331132111332-3110123301301303-2330021323310103-1033321122112002-3220032321111023-0032100310102303-2233112220131101-2331000020032120"></a>

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

<a id="canonical-3112312132023301-0103122303132012-2022333011031133-0002023020210133-3312000200123300-3321331202022211-2202330100033130-2021221212000213"></a>

## Direct properties — private_key / 001332131332 / 3

- [blindfold_secret_info](resources--workload--reference--group-006.md#canonical-3132102231030032-3330311201300002-0000231202313011-2001232130310101-3030010303022222-2202323223113203-3103110310312321-1230020033301332): complete subsection reference.

- [clear_secret_info](resources--workload--reference--group-006.md#canonical-3130130300133132-2121122113102301-0310101220022230-0100033212022002-2023020101311120-0231132130033203-1013311133332302-2003112103132002): complete subsection reference.

<a id="canonical-3232022012331100-0231030333002102-0130312332011101-0232033130312010-1010221010321102-3122211030020232-2330021320203323-1203201103220231"></a>

## Next pages — private_key / 001332131332 / 4

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates.private_key.blindfold_secret_info](resources--workload--reference--group-006.md#canonical-3132102231030032-3330311201300002-0000231202313011-2001232130310101-3030010303022222-2202323223113203-3103110310312321-1230020033301332)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates.private_key.clear_secret_info](resources--workload--reference--group-006.md#canonical-3130130300133132-2121122113102301-0310101220022230-0100033212022002-2023020101311120-0231132130033203-1013311133332302-2003112103132002)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates](resources--workload--reference--group-006.md#canonical-2023030202313320-2120223312023223-2011313021311023-1032020310121131-3203020102320011-3102231323232010-2320332011330312-2102320000023322)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-3132102231030032-3330311201300002-0000231202313011-2001232130310101-3030010303022222-2202323223113203-3103110310312321-1230020033301332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3212232103200201-0200111302302132-1012233333223122-0232130233100221-0113211222123031-0013112001312230-3312201102013213-0102313133033102"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates.private_key.blindfold_secret_info — blindfold_secret_info / 332201213323 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-2121230312030330-3010113231323001-1233013032320322-1032203002210311-2022021002113232-1123312032112001-0012103213230102-1112311001110020)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-005.md#canonical-0133221101333210-0222103021001232-3212010032013103-3331120032012221-1012203003232110-1233102323001233-2233321121201101-0122313122210200)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-005.md#canonical-2311001323203002-2023333020201320-1330320322101233-1033123102030111-2013210002333120-2130122003123103-1133301323123220-2302130221012102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-005.md#canonical-2330031030101220-3210010323021133-3322231130101013-1101131102102332-3121022130002321-1233200100132332-0213220230130033-0110233032331011)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-006.md#canonical-3023301020330331-1221233122321220-2011202113013003-2331003313323320-0122021000022233-0121202301001222-1011313113011301-2103023311112320)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates](resources--workload--reference--group-006.md#canonical-2023030202313320-2120223312023223-2011313021311023-1032020310121131-3203020102320011-3102231323232010-2320332011330312-2102320000023322)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates.private_key](resources--workload--reference--group-006.md#canonical-3031113033003222-0302203112110233-3013320030113033-2320330012301030-0323112210133122-1323133221322000-1301022232122020-1023332303223000)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates.private_key.blindfold_secret_info

<a id="canonical-3201301112233220-2332003210011130-0011021112111211-3002111010223003-1130013110130232-2031333133113222-2033322000313002-2223030233131331"></a>

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

<a id="canonical-0302320102210321-2211311203113330-2310020103021131-3200102133300233-0312330001132023-1310013201221322-3022322303313231-0002011123233023"></a>

## Direct properties — blindfold_secret_info / 332201213323 / 3

<a id="canonical-3010320011131100-0300311220302012-0330013030301212-0031013113332231-2231223020301220-2111203323320130-2212313300232121-0303112120201122"></a>

<a id="canonical-2132201102131103-3332132220203111-2032221221331333-3001232311311211-0332021123110030-2200211321320012-2332031230202321-3121220211132111"></a>

## decryption_provider property — blindfold_secret_info / 332201213323 / 4

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

<a id="canonical-1001223012313030-1302132323133100-2302320310013220-3033312113000233-1333333223002322-0002000011033122-2333210102233200-3333012312101213"></a>

<a id="canonical-0213330330112331-1103322001313133-3020220331120210-1010133303132120-2320220132221013-0120133311131020-0121230112121303-1112312021132212"></a>

## location property — blindfold_secret_info / 332201213323 / 5

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

<a id="canonical-2132111202033221-3210223321303303-1011311132022112-1320301223123100-1130111233312013-2111201012312020-1323332213010131-3203203112002313"></a>

<a id="canonical-0330030201303200-1330013031023313-1233020200132310-3013030002020012-1303212033123213-3030122020323322-0132030211102031-3333302010221333"></a>

## store_provider property — blindfold_secret_info / 332201213323 / 6

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

<a id="canonical-2300213333203001-1110300013033220-2302313301200013-3213132131312020-1013220032230301-0210222132310321-1112132310113131-3232001221023222"></a>

## Next pages — blindfold_secret_info / 332201213323 / 7

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates.private_key](resources--workload--reference--group-006.md#canonical-3031113033003222-0302203112110233-3013320030113033-2320330012301030-0323112210133122-1323133221322000-1301022232122020-1023332303223000)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-3130130300133132-2121122113102301-0310101220022230-0100033212022002-2023020101311120-0231132130033203-1013311133332302-2003112103132002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1303001202012221-3232302212031011-1213331033030030-1131022320300220-1010021122010110-2113201031322322-0332211233012131-0013310122033312"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates.private_key.clear_secret_info — clear_secret_info / 222221213212 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-2121230312030330-3010113231323001-1233013032320322-1032203002210311-2022021002113232-1123312032112001-0012103213230102-1112311001110020)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-005.md#canonical-0133221101333210-0222103021001232-3212010032013103-3331120032012221-1012203003232110-1233102323001233-2233321121201101-0122313122210200)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-005.md#canonical-2311001323203002-2023333020201320-1330320322101233-1033123102030111-2013210002333120-2130122003123103-1133301323123220-2302130221012102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-005.md#canonical-2330031030101220-3210010323021133-3322231130101013-1101131102102332-3121022130002321-1233200100132332-0213220230130033-0110233032331011)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-006.md#canonical-3023301020330331-1221233122321220-2011202113013003-2331003313323320-0122021000022233-0121202301001222-1011313113011301-2103023311112320)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates](resources--workload--reference--group-006.md#canonical-2023030202313320-2120223312023223-2011313021311023-1032020310121131-3203020102320011-3102231323232010-2320332011330312-2102320000023322)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates.private_key](resources--workload--reference--group-006.md#canonical-3031113033003222-0302203112110233-3013320030113033-2320330012301030-0323112210133122-1323133221322000-1301022232122020-1023332303223000)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates.private_key.clear_secret_info

<a id="canonical-3311001211021322-0330301013021103-1223122111101101-1323032032110211-0023021321301032-2031302200320033-2202330130130201-3300123210030120"></a>

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

<a id="canonical-3213023003220230-1131333003131330-1132000111131022-2223302031023220-3021232222202100-2232021133113130-1103002300232032-0233120120012003"></a>

## Direct properties — clear_secret_info / 222221213212 / 3

<a id="canonical-0002030111112112-0132300301311222-0103120211133132-0020312221231130-0130222202301200-2232233213222322-2111133113112001-1022110300303320"></a>

<a id="canonical-1320103320200011-3333201000133330-1133213021200312-1223013211111220-0130330012102021-2010121213030330-3210222112311333-1023333023222300"></a>

## provider_ref property — clear_secret_info / 222221213212 / 4

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-1021211320210000-2231310021023122-2220011011120203-1120231212230002-1110211012121320-2302310120310122-3102203121200001-0111002202332020"></a>

<a id="canonical-3123302300333000-0200212211331013-1220302031210130-1332312230223233-3221030213223332-1133212011331320-1231132130222021-0113020120332221"></a>

## URL property — clear_secret_info / 222221213212 / 5

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

<a id="canonical-1001320000032231-2211203303311122-2302301012213301-2230013302130220-2213131131121201-3032111112230023-2333302202112322-3121333333211233"></a>

## Next pages — clear_secret_info / 222221213212 / 6

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates.private_key](resources--workload--reference--group-006.md#canonical-3031113033003222-0302203112110233-3013320030113033-2320330012301030-0323112210133122-1323133221322000-1301022232122020-1023332303223000)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-3130200102313022-1012211213023033-1232022333113120-0320033120300203-0012330030113331-1331132310230211-2022322233100320-3323012232012310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0320230000322133-2332030223333013-3130120021330201-1312101112311131-2201320233120021-0302133200322021-0322131212020213-0311111233121203"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates.use_system_defaults — use_system_defaults / 013232101322 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-2121230312030330-3010113231323001-1233013032320322-1032203002210311-2022021002113232-1123312032112001-0012103213230102-1112311001110020)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-005.md#canonical-0133221101333210-0222103021001232-3212010032013103-3331120032012221-1012203003232110-1233102323001233-2233321121201101-0122313122210200)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-005.md#canonical-2311001323203002-2023333020201320-1330320322101233-1033123102030111-2013210002333120-2130122003123103-1133301323123220-2302130221012102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-005.md#canonical-2330031030101220-3210010323021133-3322231130101013-1101131102102332-3121022130002321-1233200100132332-0213220230130033-0110233032331011)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-006.md#canonical-3023301020330331-1221233122321220-2011202113013003-2331003313323320-0122021000022233-0121202301001222-1011313113011301-2103023311112320)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates](resources--workload--reference--group-006.md#canonical-2023030202313320-2120223312023223-2011313021311023-1032020310121131-3203020102320011-3102231323232010-2320332011330312-2102320000023322)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates.use_system_defaults

<a id="canonical-0212312021133232-3122103133221113-0200111202001131-0201223213231331-2223132121233232-1230222320323211-0303033110121101-1023212132032110"></a>

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

<a id="canonical-2121333112311323-0311123021222131-0301212113112320-3121312133231001-3110311332201012-0230210230210211-1332323302201231-3003113133113331"></a>

## Direct properties — use_system_defaults / 013232101322 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2213231120321030-1200033332332111-1113103010003300-0201112023111221-2113122131223200-0031111202023303-0223202101323132-2111301011300203"></a>

## Next pages — use_system_defaults / 013232101322 / 4

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates](resources--workload--reference--group-006.md#canonical-2023030202313320-2120223312023223-2011313021311023-1032020310121131-3203020102320011-3102231323232010-2320332011330312-2102320000023322)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-0131213121022010-2302331123122311-0333232020222202-2110030111033001-1010012332232032-2230033302023213-0102002320332232-1002130120300231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3010322103321022-1111323110031131-0131110022033201-2022302121000233-3313121031301232-1312330010222200-3111201110210013-0022231333130131"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_config — tls_config / 332131222113 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-2121230312030330-3010113231323001-1233013032320322-1032203002210311-2022021002113232-1123312032112001-0012103213230102-1112311001110020)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-005.md#canonical-0133221101333210-0222103021001232-3212010032013103-3331120032012221-1012203003232110-1233102323001233-2233321121201101-0122313122210200)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-005.md#canonical-2311001323203002-2023333020201320-1330320322101233-1033123102030111-2013210002333120-2130122003123103-1133301323123220-2302130221012102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-005.md#canonical-2330031030101220-3210010323021133-3322231130101013-1101131102102332-3121022130002321-1233200100132332-0213220230130033-0110233032331011)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-006.md#canonical-3023301020330331-1221233122321220-2011202113013003-2331003313323320-0122021000022233-0121202301001222-1011313113011301-2103023311112320)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_config

<a id="canonical-0111012211012203-2321213132300010-2221312320221011-0311121122320331-1330100230031101-2210130232223103-2322032012022323-1323223230332312"></a>

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

<a id="canonical-3232230221031230-0133032333011102-2302220003023130-1100123133310122-1302031121032020-3011233123331133-1031102013213030-1301033233020300"></a>

## Direct properties — tls_config / 332131222113 / 3

- [custom_security](resources--workload--reference--group-006.md#canonical-3300213330112100-0310203322220231-0033200003022313-1100221202220000-1331112022112312-0020220210300331-2311111202203022-1000130320221013): complete subsection reference.

- [default_security](resources--workload--reference--group-006.md#canonical-2011233321321130-3201322122230032-0002032010203313-0313232221033211-2000230233313320-1331303012122032-2010300331001013-3220222100022011): complete subsection reference.

- [low_security](resources--workload--reference--group-006.md#canonical-2012232121033232-1130011220011211-2332303232301121-0200323121313003-3332332103310330-1332020100210301-0322101200333112-1032002123211110): complete subsection reference.

- [medium_security](resources--workload--reference--group-006.md#canonical-0003320232020300-1300231123101010-0003320113021101-2223321032212013-1103032100102230-2000203322111131-1301130203232031-3122230202303302): complete subsection reference.

<a id="canonical-0233133212002222-1313003000001010-3212223010001022-0231300131311330-0213033000011102-2212033131331202-3122103110112203-0302311032310111"></a>

## Next pages — tls_config / 332131222113 / 4

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_config.custom_security](resources--workload--reference--group-006.md#canonical-3300213330112100-0310203322220231-0033200003022313-1100221202220000-1331112022112312-0020220210300331-2311111202203022-1000130320221013)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_config.default_security](resources--workload--reference--group-006.md#canonical-2011233321321130-3201322122230032-0002032010203313-0313232221033211-2000230233313320-1331303012122032-2010300331001013-3220222100022011)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_config.low_security](resources--workload--reference--group-006.md#canonical-2012232121033232-1130011220011211-2332303232301121-0200323121313003-3332332103310330-1332020100210301-0322101200333112-1032002123211110)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_config.medium_security](resources--workload--reference--group-006.md#canonical-0003320232020300-1300231123101010-0003320113021101-2223321032212013-1103032100102230-2000203322111131-1301130203232031-3122230202303302)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-006.md#canonical-3023301020330331-1221233122321220-2011202113013003-2331003313323320-0122021000022233-0121202301001222-1011313113011301-2103023311112320)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-3300213330112100-0310203322220231-0033200003022313-1100221202220000-1331112022112312-0020220210300331-2311111202203022-1000130320221013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3303210202131003-2212332103021113-0210333220130320-1102033122133023-3321200331132303-3013301211000111-0300330133311030-3300000233312201"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_config.custom_security — custom_security / 121103223311 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-2121230312030330-3010113231323001-1233013032320322-1032203002210311-2022021002113232-1123312032112001-0012103213230102-1112311001110020)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-005.md#canonical-0133221101333210-0222103021001232-3212010032013103-3331120032012221-1012203003232110-1233102323001233-2233321121201101-0122313122210200)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-005.md#canonical-2311001323203002-2023333020201320-1330320322101233-1033123102030111-2013210002333120-2130122003123103-1133301323123220-2302130221012102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-005.md#canonical-2330031030101220-3210010323021133-3322231130101013-1101131102102332-3121022130002321-1233200100132332-0213220230130033-0110233032331011)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-006.md#canonical-3023301020330331-1221233122321220-2011202113013003-2331003313323320-0122021000022233-0121202301001222-1011313113011301-2103023311112320)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_config](resources--workload--reference--group-006.md#canonical-0131213121022010-2302331123122311-0333232020222202-2110030111033001-1010012332232032-2230033302023213-0102002320332232-1002130120300231)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_config.custom_security

<a id="canonical-1110103000221223-2121210100122110-3001133220000121-3132023001201301-3303332230200333-1012300020232000-0003203331222013-2330303011020321"></a>

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

<a id="canonical-1000112120121113-2100311131221003-3230322012300221-3113032210111330-3331322030102021-2121021022013131-0121002231023033-1323203313012001"></a>

## Direct properties — custom_security / 121103223311 / 3

<a id="canonical-1112011223031100-1111000300130133-3121021002011110-3212231230132012-1112000232322220-2021302200211101-0011012232300302-1310022232200013"></a>

<a id="canonical-2312222033002021-3330332302323112-3330201002010100-0122102302003312-3301031102033222-2223023312231333-1031102222201031-1023000313221131"></a>

## cipher_suites property — custom_security / 121103223311 / 4

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

<a id="canonical-0222233211103301-2023021131110221-1131132211001103-3233212033320000-0101212211330221-1323233233022303-0200223310321033-0320011122312311"></a>

<a id="canonical-0022202210033130-1001323133222210-1210213300210112-3100110302231313-3221202311211030-0303133131223032-0220332123112122-3121103013303011"></a>

## max_version property — custom_security / 121103223311 / 5

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

<a id="canonical-1210112122202103-1030122112310233-1220330113301232-1202222113210112-2213101331200232-0310022321320230-2303212201012213-2201231320320322"></a>

<a id="canonical-2101020120322331-3011220003223020-0201313033132121-3000202110313131-2020103030231223-0300103221013120-0220133123310013-3000122333102112"></a>

## min_version property — custom_security / 121103223311 / 6

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

<a id="canonical-2130021021321211-2022022131222210-1230021301121232-2113020312000101-0202120222101032-1301211221111133-3110313111302312-3013322103022001"></a>

## Next pages — custom_security / 121103223311 / 7

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_config](resources--workload--reference--group-006.md#canonical-0131213121022010-2302331123122311-0333232020222202-2110030111033001-1010012332232032-2230033302023213-0102002320332232-1002130120300231)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-2011233321321130-3201322122230032-0002032010203313-0313232221033211-2000230233313320-1331303012122032-2010300331001013-3220222100022011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0103232303300300-1021103000103302-1332312300322213-2101010230230102-2033123332300102-3302210003230000-3312330020003131-3023023313122113"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_config.default_security — default_security / 232200101133 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-2121230312030330-3010113231323001-1233013032320322-1032203002210311-2022021002113232-1123312032112001-0012103213230102-1112311001110020)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-005.md#canonical-0133221101333210-0222103021001232-3212010032013103-3331120032012221-1012203003232110-1233102323001233-2233321121201101-0122313122210200)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-005.md#canonical-2311001323203002-2023333020201320-1330320322101233-1033123102030111-2013210002333120-2130122003123103-1133301323123220-2302130221012102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-005.md#canonical-2330031030101220-3210010323021133-3322231130101013-1101131102102332-3121022130002321-1233200100132332-0213220230130033-0110233032331011)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-006.md#canonical-3023301020330331-1221233122321220-2011202113013003-2331003313323320-0122021000022233-0121202301001222-1011313113011301-2103023311112320)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_config](resources--workload--reference--group-006.md#canonical-0131213121022010-2302331123122311-0333232020222202-2110030111033001-1010012332232032-2230033302023213-0102002320332232-1002130120300231)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_config.default_security

<a id="canonical-3211121100023102-3203103102232301-1200221330120000-2133311031220031-0010032212313311-0300122113331022-3011322020313001-1012130012021312"></a>

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

<a id="canonical-3033001011232211-1303221322221100-1002120001331023-2120331220111130-3010232122220003-3033033133021222-0332032000332310-2121012330223323"></a>

## Direct properties — default_security / 232200101133 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3000101321300031-3202302333302302-3322200231212020-2103211311331322-3113302122013123-0231333300220303-0202032023013200-2011300233230210"></a>

## Next pages — default_security / 232200101133 / 4

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_config](resources--workload--reference--group-006.md#canonical-0131213121022010-2302331123122311-0333232020222202-2110030111033001-1010012332232032-2230033302023213-0102002320332232-1002130120300231)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-2012232121033232-1130011220011211-2332303232301121-0200323121313003-3332332103310330-1332020100210301-0322101200333112-1032002123211110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2130023002130231-2132221200311231-0230203000011233-2330210223332033-3213333303212122-1231212102203232-2020302122030323-2303322220032221"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_config.low_security — low_security / 322120110200 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-2121230312030330-3010113231323001-1233013032320322-1032203002210311-2022021002113232-1123312032112001-0012103213230102-1112311001110020)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-005.md#canonical-0133221101333210-0222103021001232-3212010032013103-3331120032012221-1012203003232110-1233102323001233-2233321121201101-0122313122210200)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-005.md#canonical-2311001323203002-2023333020201320-1330320322101233-1033123102030111-2013210002333120-2130122003123103-1133301323123220-2302130221012102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-005.md#canonical-2330031030101220-3210010323021133-3322231130101013-1101131102102332-3121022130002321-1233200100132332-0213220230130033-0110233032331011)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-006.md#canonical-3023301020330331-1221233122321220-2011202113013003-2331003313323320-0122021000022233-0121202301001222-1011313113011301-2103023311112320)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_config](resources--workload--reference--group-006.md#canonical-0131213121022010-2302331123122311-0333232020222202-2110030111033001-1010012332232032-2230033302023213-0102002320332232-1002130120300231)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_config.low_security

<a id="canonical-2302223013322231-1122200010000322-1313102031331123-0321303123102203-3223301221033231-2012020100031013-2120200103031302-2330231321321010"></a>

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

<a id="canonical-0332302020002321-3132010103203211-2210033300300300-1323021112011011-1130313123030232-2220103103312300-2221302100133110-0122310323010023"></a>

## Direct properties — low_security / 322120110200 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1333122101321122-0220011101213230-0301330111232112-1121320231012130-1313211301301200-1210001322320302-3013110311113030-2020011033220033"></a>

## Next pages — low_security / 322120110200 / 4

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_config](resources--workload--reference--group-006.md#canonical-0131213121022010-2302331123122311-0333232020222202-2110030111033001-1010012332232032-2230033302023213-0102002320332232-1002130120300231)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-0003320232020300-1300231123101010-0003320113021101-2223321032212013-1103032100102230-2000203322111131-1301130203232031-3122230202303302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1212013011232200-0122101310332102-0311320302200203-3011002301221003-1112311213121121-1211112331210001-0333123222202013-2303333202001220"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_config.medium_security — medium_security / 020010023102 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-2121230312030330-3010113231323001-1233013032320322-1032203002210311-2022021002113232-1123312032112001-0012103213230102-1112311001110020)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-005.md#canonical-0133221101333210-0222103021001232-3212010032013103-3331120032012221-1012203003232110-1233102323001233-2233321121201101-0122313122210200)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-005.md#canonical-2311001323203002-2023333020201320-1330320322101233-1033123102030111-2013210002333120-2130122003123103-1133301323123220-2302130221012102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-005.md#canonical-2330031030101220-3210010323021133-3322231130101013-1101131102102332-3121022130002321-1233200100132332-0213220230130033-0110233032331011)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-006.md#canonical-3023301020330331-1221233122321220-2011202113013003-2331003313323320-0122021000022233-0121202301001222-1011313113011301-2103023311112320)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_config](resources--workload--reference--group-006.md#canonical-0131213121022010-2302331123122311-0333232020222202-2110030111033001-1010012332232032-2230033302023213-0102002320332232-1002130120300231)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_config.medium_security

<a id="canonical-2010223101120332-2120331023133311-2112320303210101-3133232221333121-3201302333111111-1211030213100331-2100000022330010-2011120002110122"></a>

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

<a id="canonical-2323311120003103-2100210310100212-1023322013312302-1322310322212013-2311123002131132-3102203332010103-2232010123220310-2031010113132101"></a>

## Direct properties — medium_security / 020010023102 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3302103233220100-1220033313022222-2321230222332031-3232010003102102-3203301113112321-1232200033310100-3321130313323221-3002300000230210"></a>

## Next pages — medium_security / 020010023102 / 4

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_config](resources--workload--reference--group-006.md#canonical-0131213121022010-2302331123122311-0333232020222202-2110030111033001-1010012332232032-2230033302023213-0102002320332232-1002130120300231)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-1201123022002000-0122023033012312-0210013220112133-1220130111331213-0321022230013202-3213120123303001-1002300100311330-1023002312220123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2333030021202133-2133220110203203-3130220131203122-1133200120002333-0013311323110121-2201110003132103-2132001003003112-0312131110032203"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.use_mtls — use_mtls / 100122231000 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-2121230312030330-3010113231323001-1233013032320322-1032203002210311-2022021002113232-1123312032112001-0012103213230102-1112311001110020)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-005.md#canonical-0133221101333210-0222103021001232-3212010032013103-3331120032012221-1012203003232110-1233102323001233-2233321121201101-0122313122210200)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-005.md#canonical-2311001323203002-2023333020201320-1330320322101233-1033123102030111-2013210002333120-2130122003123103-1133301323123220-2302130221012102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-005.md#canonical-2330031030101220-3210010323021133-3322231130101013-1101131102102332-3121022130002321-1233200100132332-0213220230130033-0110233032331011)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-006.md#canonical-3023301020330331-1221233122321220-2011202113013003-2331003313323320-0122021000022233-0121202301001222-1011313113011301-2103023311112320)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.use_mtls

<a id="canonical-1121012203011210-3333202300113201-0012031100103303-2201301001222123-1013002103022211-1031111023310030-1223311113203020-3123203331330213"></a>

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

<a id="canonical-2102100031023323-0300300333010310-3232103102212212-1330232112020020-3230123033012122-0331213102212200-1333031022011222-0311233031312310"></a>

## Direct properties — use_mtls / 100122231000 / 3

<a id="canonical-1031022011120122-0113133132231320-1221003323110113-3123123212032232-2012213022312032-2331301301322133-2310203321221032-2232231103001101"></a>

<a id="canonical-2221311202020333-2123013230001002-0321122013322001-3012110232221322-0133121321200113-0210230010131010-2300133120220303-3320021020310133"></a>

## client_certificate_optional property — use_mtls / 100122231000 / 4

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

- [crl](resources--workload--reference--group-006.md#canonical-3300311222020333-1030030320231030-3301031313003330-1200212122310110-0210201220111113-3322220310010031-2211000131231301-2011123001002023): complete subsection reference.

- [no_crl](resources--workload--reference--group-006.md#canonical-0101031013110210-3301211201131132-2332013211122123-1203010200202230-0302122311231131-0000302303011203-1032133113001312-1102111020133131): complete subsection reference.

- [trusted_ca](resources--workload--reference--group-006.md#canonical-3012021133200132-0121233331132323-2001013132311003-1303021013222301-0222013312123101-0011132232102131-3221210133123022-0300110000120121): complete subsection reference.

<a id="canonical-3212011011113133-2220012321321223-2003020103102002-2030202132203222-2130033220313013-3200320303031000-1021300000330332-3131303112300012"></a>

<a id="canonical-2011022030133221-3310232001320332-3102001002020333-2323231311131310-2220001230002111-2303303103032223-3311301130131112-0331033220220311"></a>

## trusted_ca_url property — use_mtls / 100122231000 / 5

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

- [xfcc_disabled](resources--workload--reference--group-007.md#canonical-3101111211011003-0322113122333131-1320332000301323-2021001001032003-2302030313012323-0033322030001123-0200012121320010-1233222220022003): complete subsection reference.

- [xfcc_options](resources--workload--reference--group-007.md#canonical-2002102111223013-1130130103332311-0013000222311033-3323112200211201-3132200112231120-0221003231221002-3331103111300223-3222022211030303): complete subsection reference.

<a id="canonical-1212032331321223-3312203021002313-0312010110302201-3312333210302123-2101313131202312-1131210203322010-1100300201123033-3101233300023300"></a>

## Next pages — use_mtls / 100122231000 / 6

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.use_mtls.crl](resources--workload--reference--group-006.md#canonical-3300311222020333-1030030320231030-3301031313003330-1200212122310110-0210201220111113-3322220310010031-2211000131231301-2011123001002023)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.use_mtls.no_crl](resources--workload--reference--group-006.md#canonical-0101031013110210-3301211201131132-2332013211122123-1203010200202230-0302122311231131-0000302303011203-1032133113001312-1102111020133131)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.use_mtls.trusted_ca](resources--workload--reference--group-006.md#canonical-3012021133200132-0121233331132323-2001013132311003-1303021013222301-0222013312123101-0011132232102131-3221210133123022-0300110000120121)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.use_mtls.xfcc_disabled](resources--workload--reference--group-007.md#canonical-3101111211011003-0322113122333131-1320332000301323-2021001001032003-2302030313012323-0033322030001123-0200012121320010-1233222220022003)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.use_mtls.xfcc_options](resources--workload--reference--group-007.md#canonical-2002102111223013-1130130103332311-0013000222311033-3323112200211201-3132200112231120-0221003231221002-3331103111300223-3222022211030303)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-006.md#canonical-3023301020330331-1221233122321220-2011202113013003-2331003313323320-0122021000022233-0121202301001222-1011313113011301-2103023311112320)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-3300311222020333-1030030320231030-3301031313003330-1200212122310110-0210201220111113-3322220310010031-2211000131231301-2011123001002023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1332201330132021-1131002300122102-3331131202122302-0130220220003122-1023031330200110-0330202232100033-2303003211223210-0131331332322023"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.use_mtls.crl — crl / 011111330012 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-2121230312030330-3010113231323001-1233013032320322-1032203002210311-2022021002113232-1123312032112001-0012103213230102-1112311001110020)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-005.md#canonical-0133221101333210-0222103021001232-3212010032013103-3331120032012221-1012203003232110-1233102323001233-2233321121201101-0122313122210200)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-005.md#canonical-2311001323203002-2023333020201320-1330320322101233-1033123102030111-2013210002333120-2130122003123103-1133301323123220-2302130221012102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-005.md#canonical-2330031030101220-3210010323021133-3322231130101013-1101131102102332-3121022130002321-1233200100132332-0213220230130033-0110233032331011)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-006.md#canonical-3023301020330331-1221233122321220-2011202113013003-2331003313323320-0122021000022233-0121202301001222-1011313113011301-2103023311112320)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.use_mtls](resources--workload--reference--group-006.md#canonical-1201123022002000-0122023033012312-0210013220112133-1220130111331213-0321022230013202-3213120123303001-1002300100311330-1023002312220123)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.use_mtls.crl

<a id="canonical-3313200202322302-2110333300320130-3011011122312030-2231112130020323-3131010212102120-3113332331121231-2201111333021101-0031223123201222"></a>

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

<a id="canonical-3333012032322220-1011001100200322-3013020220210300-3232303200331122-3023103202221320-2332310213002032-0333300110303000-2003231022031230"></a>

## Direct properties — crl / 011111330012 / 3

<a id="canonical-2020010122010222-1000311321012023-3221001220003013-0103300333223203-0101102000112200-0221213002302120-3230110102111132-2012111132301303"></a>

<a id="canonical-2122102003120332-3202022123211201-0222133333130212-0233333310302133-0213220233312122-1111100321123303-3111001333232003-3131103012211211"></a>

## name property — crl / 011111330012 / 4

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

<a id="canonical-1111220130132322-1202231111300102-1211220130203000-3013231033311031-0232000311023022-3010301211121102-0301313300103232-3002211310133303"></a>

<a id="canonical-0131303223211331-1000310133332203-3101033020031031-1320220101131120-3210122032213330-3230311333210203-2122131031330031-1030233303100033"></a>

## namespace property — crl / 011111330012 / 5

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

<a id="canonical-2213203132022002-1121333013201200-2231132103120001-1013001210332122-0303222111333310-3001010100322232-2111230201133022-2100211001302301"></a>

<a id="canonical-2010220101101131-0122231323031211-1103001312321321-2100301011220020-3221233001022112-0030211111232331-1311010031223202-0300320301323221"></a>

## tenant property — crl / 011111330012 / 6

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

<a id="canonical-2101123200013311-2202013001200303-0322222203303101-1332332203202121-2312211110203213-3132100222121123-0010310301320223-3033323112211112"></a>

## Next pages — crl / 011111330012 / 7

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.use_mtls](resources--workload--reference--group-006.md#canonical-1201123022002000-0122023033012312-0210013220112133-1220130111331213-0321022230013202-3213120123303001-1002300100311330-1023002312220123)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-0101031013110210-3301211201131132-2332013211122123-1203010200202230-0302122311231131-0000302303011203-1032133113001312-1102111020133131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1011202233022013-0132121321120111-3000201132220331-3202323312132120-0002213112332200-1201223331200033-3201002221330223-2331302120202313"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.use_mtls.no_crl — no_crl / 111131323132 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-2121230312030330-3010113231323001-1233013032320322-1032203002210311-2022021002113232-1123312032112001-0012103213230102-1112311001110020)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-005.md#canonical-0133221101333210-0222103021001232-3212010032013103-3331120032012221-1012203003232110-1233102323001233-2233321121201101-0122313122210200)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-005.md#canonical-2311001323203002-2023333020201320-1330320322101233-1033123102030111-2013210002333120-2130122003123103-1133301323123220-2302130221012102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-005.md#canonical-2330031030101220-3210010323021133-3322231130101013-1101131102102332-3121022130002321-1233200100132332-0213220230130033-0110233032331011)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-006.md#canonical-3023301020330331-1221233122321220-2011202113013003-2331003313323320-0122021000022233-0121202301001222-1011313113011301-2103023311112320)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.use_mtls](resources--workload--reference--group-006.md#canonical-1201123022002000-0122023033012312-0210013220112133-1220130111331213-0321022230013202-3213120123303001-1002300100311330-1023002312220123)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.use_mtls.no_crl

<a id="canonical-0002213023230000-1232330203322310-1012131302020033-2303301010033000-2201131000031123-0203303210303013-1300030133120311-1020332223200001"></a>

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

<a id="canonical-1200023310322230-3223100100010333-0111012103311011-2320230021331012-3233123222230311-2033023323121200-1033212313323312-0133020133311231"></a>

## Direct properties — no_crl / 111131323132 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0031301300001331-2303120130331331-2122220011022120-1300221110011211-1312333130210231-2302321131121021-0311322120221233-3121133312130012"></a>

## Next pages — no_crl / 111131323132 / 4

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.use_mtls](resources--workload--reference--group-006.md#canonical-1201123022002000-0122023033012312-0210013220112133-1220130111331213-0321022230013202-3213120123303001-1002300100311330-1023002312220123)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-3012021133200132-0121233331132323-2001013132311003-1303021013222301-0222013312123101-0011132232102131-3221210133123022-0300110000120121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->
