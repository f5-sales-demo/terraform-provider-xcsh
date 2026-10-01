---
page_title: "xcsh_workload reference"
subcategory: "Container"
description: "Complete grouped canonical reference for xcsh_workload reference."
---

# xcsh_workload reference

<a id="canonical-1111300011010103-0303311220331231-2130213331122303-1131333321003003-3320232132231223-1331132232000130-1123103133300310-3021320222230020"></a>

## stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.http_protocol_options — http_protocol_options / 321003213320 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-017.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-021.md#canonical-3300101333202310-3223231311312310-2030111033033203-1011103231310023-1313110120311121-1221031312313020-1210302113302033-2110231331000202)
- [stateful_service.advertise_options.advertise_on_public.port](resources--workload--reference--group-025.md#canonical-0120012201110120-2320030110233132-3233102311321321-2202232022223030-2133131131013111-3231210131302330-1022020011332122-3111230313032011)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-025.md#canonical-2203002110120031-3203013111221013-1321101003110222-2103121303322022-0021002301311230-0203100203111101-0321302220100020-1320220233023212)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert](resources--workload--reference--group-026.md#canonical-2220320330100322-1330210101012132-0301302331113332-3313313132203223-2330123211123011-0222332113113122-0203030132101203-1231012010033022)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.http_protocol_options

<a id="canonical-2323323032102323-1331333302231213-0032123301020103-1212202100113133-3102230020230221-3203333020333023-1202130102311012-2311211220000111"></a>

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

<a id="canonical-1113321110303220-0303121120330223-0012111103101201-2020200123110100-3301321123022002-3122202011033103-0321311232303010-1032210310013102"></a>

## Direct properties — http_protocol_options / 321003213320 / 3

- [http_protocol_enable_v1_only](resources--workload--reference--group-027.md#canonical-3332103310011000-1311211312111102-0121211013231233-1322332000022223-0222121321121201-1212103101231101-2101201332230012-2013111011231320): complete subsection reference.

- [http_protocol_enable_v1_v2](resources--workload--reference--group-027.md#canonical-2113033122121112-2210020003032101-1203113022113120-1030230321113231-0031031310210330-0320113301220233-2023223013113132-2011023132302123): complete subsection reference.

- [http_protocol_enable_v2_only](resources--workload--reference--group-027.md#canonical-2032023010031211-3013130120220333-0121302223022100-0002211223231120-3123120032122221-3321002012223002-2223302120021232-3301333102312203): complete subsection reference.

<a id="canonical-0310330200232021-0212102200022102-3201232320210313-0200130002000133-3232012100031110-2233131330122302-3120000303112230-3313321113120133"></a>

## Next pages — http_protocol_options / 321003213320 / 4

- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only](resources--workload--reference--group-027.md#canonical-3332103310011000-1311211312111102-0121211013231233-1322332000022223-0222121321121201-1212103101231101-2101201332230012-2013111011231320)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_v2](resources--workload--reference--group-027.md#canonical-2113033122121112-2210020003032101-1203113022113120-1030230321113231-0031031310210330-0320113301220233-2023223013113132-2011023132302123)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v2_only](resources--workload--reference--group-027.md#canonical-2032023010031211-3013130120220333-0121302223022100-0002211223231120-3123120032122221-3321002012223002-2223302120021232-3301333102312203)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert](resources--workload--reference--group-026.md#canonical-2220320330100322-1330210101012132-0301302331113332-3313313132203223-2330123211123011-0222332113113122-0203030132101203-1231012010033022)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-3332103310011000-1311211312111102-0121211013231233-1322332000022223-0222121321121201-1212103101231101-2101201332230012-2013111011231320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2103210033320103-2032212223012112-3310020332201331-1001200001320021-0201333311111032-3313213222122310-2033303020322220-1231301031210202"></a>

## stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only — http_protocol_enable_v1_only / 232022230221 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-017.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-021.md#canonical-3300101333202310-3223231311312310-2030111033033203-1011103231310023-1313110120311121-1221031312313020-1210302113302033-2110231331000202)
- [stateful_service.advertise_options.advertise_on_public.port](resources--workload--reference--group-025.md#canonical-0120012201110120-2320030110233132-3233102311321321-2202232022223030-2133131131013111-3231210131302330-1022020011332122-3111230313032011)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-025.md#canonical-2203002110120031-3203013111221013-1321101003110222-2103121303322022-0021002301311230-0203100203111101-0321302220100020-1320220233023212)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert](resources--workload--reference--group-026.md#canonical-2220320330100322-1330210101012132-0301302331113332-3313313132203223-2330123211123011-0222332113113122-0203030132101203-1231012010033022)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.http_protocol_options](resources--workload--reference--group-026.md#canonical-3111212033013022-0013130201203321-3112121303030312-1000312300110101-1302000111133002-2123102233332323-3310121102112230-0011132102231313)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only

<a id="canonical-3101213231330030-0212023113130030-1201110011203123-2131302232302302-2331331112013013-2332000213023222-3331311120323110-3211032201232100"></a>

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

<a id="canonical-2311232330103301-3330333301011312-2123201232122201-3001323213120223-0232220113301120-2221231101003201-1103032033311020-3212301012120300"></a>

## Direct properties — http_protocol_enable_v1_only / 232022230221 / 3

- [header_transformation](resources--workload--reference--group-027.md#canonical-1030200033313230-3312300021230030-1032121021122312-2023130001230122-3201212000330133-1313231032211023-0100323311112332-1101000222231320): complete subsection reference.

<a id="canonical-1122103002210333-2213323301013210-1130102311232311-2030033233030032-1101300000322022-2103321301312211-1023113002322301-0031321000011213"></a>

## Next pages — http_protocol_enable_v1_only / 232022230221 / 4

- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation](resources--workload--reference--group-027.md#canonical-1030200033313230-3312300021230030-1032121021122312-2023130001230122-3201212000330133-1313231032211023-0100323311112332-1101000222231320)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.http_protocol_options](resources--workload--reference--group-026.md#canonical-3111212033013022-0013130201203321-3112121303030312-1000312300110101-1302000111133002-2123102233332323-3310121102112230-0011132102231313)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-1030200033313230-3312300021230030-1032121021122312-2023130001230122-3201212000330133-1313231032211023-0100323311112332-1101000222231320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1233220012233101-3130013022030033-2222113102330123-2030103213320132-0201113131213123-3232031213112022-2201130023020122-1300333022223000"></a>

## stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation — header_transformation / 123220323002 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-017.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-021.md#canonical-3300101333202310-3223231311312310-2030111033033203-1011103231310023-1313110120311121-1221031312313020-1210302113302033-2110231331000202)
- [stateful_service.advertise_options.advertise_on_public.port](resources--workload--reference--group-025.md#canonical-0120012201110120-2320030110233132-3233102311321321-2202232022223030-2133131131013111-3231210131302330-1022020011332122-3111230313032011)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-025.md#canonical-2203002110120031-3203013111221013-1321101003110222-2103121303322022-0021002301311230-0203100203111101-0321302220100020-1320220233023212)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert](resources--workload--reference--group-026.md#canonical-2220320330100322-1330210101012132-0301302331113332-3313313132203223-2330123211123011-0222332113113122-0203030132101203-1231012010033022)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.http_protocol_options](resources--workload--reference--group-026.md#canonical-3111212033013022-0013130201203321-3112121303030312-1000312300110101-1302000111133002-2123102233332323-3310121102112230-0011132102231313)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only](resources--workload--reference--group-027.md#canonical-3332103310011000-1311211312111102-0121211013231233-1322332000022223-0222121321121201-1212103101231101-2101201332230012-2013111011231320)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation

<a id="canonical-1222233331310213-2031113332312022-0012032020023312-3111001101203022-0230132312230212-2302133132212011-2030033122220332-2023033022220101"></a>

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

<a id="canonical-3211332032033120-0210001222012311-2230233220202001-3312011331000312-1130300220230022-2331031122201113-3122321111123331-0030012130320001"></a>

## Direct properties — header_transformation / 123220323002 / 3

- [default_header_transformation](resources--workload--reference--group-027.md#canonical-1223310022120000-2011202203030332-2002000020300123-0130313133130233-1132011210323302-0023110203023211-0330130230111200-0311203212330133): complete subsection reference.

- [preserve_case_header_transformation](resources--workload--reference--group-027.md#canonical-3030332312112321-3122101022320213-2312112132213201-0211103312121321-2330003221313310-3303001101330110-0212310103123011-1030022012220333): complete subsection reference.

- [proper_case_header_transformation](resources--workload--reference--group-027.md#canonical-0231023032310213-1020012112220312-1232100002330131-0013300111023230-0323230111302231-0001130032230001-3001002332222131-1232110010030230): complete subsection reference.

<a id="canonical-3011002033103220-1210102331333120-0323213233213313-1222212111002231-3300222032133031-3213020123223003-3311113020213103-2012032201021203"></a>

## Next pages — header_transformation / 123220323002 / 4

- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.default_header_transformation](resources--workload--reference--group-027.md#canonical-1223310022120000-2011202203030332-2002000020300123-0130313133130233-1132011210323302-0023110203023211-0330130230111200-0311203212330133)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.preserve_case_header_transformation](resources--workload--reference--group-027.md#canonical-3030332312112321-3122101022320213-2312112132213201-0211103312121321-2330003221313310-3303001101330110-0212310103123011-1030022012220333)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.proper_case_header_transformation](resources--workload--reference--group-027.md#canonical-0231023032310213-1020012112220312-1232100002330131-0013300111023230-0323230111302231-0001130032230001-3001002332222131-1232110010030230)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only](resources--workload--reference--group-027.md#canonical-3332103310011000-1311211312111102-0121211013231233-1322332000022223-0222121321121201-1212103101231101-2101201332230012-2013111011231320)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-1223310022120000-2011202203030332-2002000020300123-0130313133130233-1132011210323302-0023110203023211-0330130230111200-0311203212330133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3033101331321113-0311132030232103-0312030111201100-0221233021331231-2332220303232222-1102232322200222-3012230020021031-2033211202222230"></a>

## stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.default_header_transformation — default_header_transformation / 302332000332 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-017.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-021.md#canonical-3300101333202310-3223231311312310-2030111033033203-1011103231310023-1313110120311121-1221031312313020-1210302113302033-2110231331000202)
- [stateful_service.advertise_options.advertise_on_public.port](resources--workload--reference--group-025.md#canonical-0120012201110120-2320030110233132-3233102311321321-2202232022223030-2133131131013111-3231210131302330-1022020011332122-3111230313032011)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-025.md#canonical-2203002110120031-3203013111221013-1321101003110222-2103121303322022-0021002301311230-0203100203111101-0321302220100020-1320220233023212)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert](resources--workload--reference--group-026.md#canonical-2220320330100322-1330210101012132-0301302331113332-3313313132203223-2330123211123011-0222332113113122-0203030132101203-1231012010033022)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.http_protocol_options](resources--workload--reference--group-026.md#canonical-3111212033013022-0013130201203321-3112121303030312-1000312300110101-1302000111133002-2123102233332323-3310121102112230-0011132102231313)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only](resources--workload--reference--group-027.md#canonical-3332103310011000-1311211312111102-0121211013231233-1322332000022223-0222121321121201-1212103101231101-2101201332230012-2013111011231320)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation](resources--workload--reference--group-027.md#canonical-1030200033313230-3312300021230030-1032121021122312-2023130001230122-3201212000330133-1313231032211023-0100323311112332-1101000222231320)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.default_header_transformation

<a id="canonical-1213231320302131-3212111022321210-1212223230310202-3312030323131032-0120203301321232-1033211110300333-2310120203202003-2312110110133121"></a>

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

<a id="canonical-0100302221120312-1212112332003320-0331332222221223-0301210110303323-2022133101020331-0022131023131120-3302301332100313-2022122221102002"></a>

## Direct properties — default_header_transformation / 302332000332 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0203020000321310-1203032012100032-3301020100321110-1301233202301011-3001000323121010-2310130031230321-0123032032013330-0011223031022302"></a>

## Next pages — default_header_transformation / 302332000332 / 4

- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation](resources--workload--reference--group-027.md#canonical-1030200033313230-3312300021230030-1032121021122312-2023130001230122-3201212000330133-1313231032211023-0100323311112332-1101000222231320)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-3030332312112321-3122101022320213-2312112132213201-0211103312121321-2330003221313310-3303001101330110-0212310103123011-1030022012220333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3031003223232100-2011200111323102-2002021113212110-0300103030002003-0200312121130102-2233001313230321-1302111223102210-2313101103023303"></a>

## stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.preserve_case_header_transformation — preserve_case_header_transformation / 212030231012 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-017.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-021.md#canonical-3300101333202310-3223231311312310-2030111033033203-1011103231310023-1313110120311121-1221031312313020-1210302113302033-2110231331000202)
- [stateful_service.advertise_options.advertise_on_public.port](resources--workload--reference--group-025.md#canonical-0120012201110120-2320030110233132-3233102311321321-2202232022223030-2133131131013111-3231210131302330-1022020011332122-3111230313032011)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-025.md#canonical-2203002110120031-3203013111221013-1321101003110222-2103121303322022-0021002301311230-0203100203111101-0321302220100020-1320220233023212)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert](resources--workload--reference--group-026.md#canonical-2220320330100322-1330210101012132-0301302331113332-3313313132203223-2330123211123011-0222332113113122-0203030132101203-1231012010033022)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.http_protocol_options](resources--workload--reference--group-026.md#canonical-3111212033013022-0013130201203321-3112121303030312-1000312300110101-1302000111133002-2123102233332323-3310121102112230-0011132102231313)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only](resources--workload--reference--group-027.md#canonical-3332103310011000-1311211312111102-0121211013231233-1322332000022223-0222121321121201-1212103101231101-2101201332230012-2013111011231320)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation](resources--workload--reference--group-027.md#canonical-1030200033313230-3312300021230030-1032121021122312-2023130001230122-3201212000330133-1313231032211023-0100323311112332-1101000222231320)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.preserve_case_header_transformation

<a id="canonical-1003121103321303-1020220203301022-2002000131211011-0203002023330332-2321331312113330-2300331202033330-1333020303132230-1213333003011013"></a>

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

<a id="canonical-1321331313111120-3322032111303331-2201113103203032-3023303012312112-0033021101203320-0121202212000203-1323012001320222-3132001313231003"></a>

## Direct properties — preserve_case_header_transformation / 212030231012 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2121311200112333-2001131310231323-3202010130323222-0122130113333132-1203323313102320-3311021031023220-1113230103303203-1200022211123212"></a>

## Next pages — preserve_case_header_transformation / 212030231012 / 4

- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation](resources--workload--reference--group-027.md#canonical-1030200033313230-3312300021230030-1032121021122312-2023130001230122-3201212000330133-1313231032211023-0100323311112332-1101000222231320)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-0231023032310213-1020012112220312-1232100002330131-0013300111023230-0323230111302231-0001130032230001-3001002332222131-1232110010030230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3110322132211302-1322032332323322-2332001213020312-3231300120301123-0322131103131202-3120111200032321-3213312002330022-2223200102111021"></a>

## stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.proper_case_header_transformation — proper_case_header_transformation / 322310102301 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-017.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-021.md#canonical-3300101333202310-3223231311312310-2030111033033203-1011103231310023-1313110120311121-1221031312313020-1210302113302033-2110231331000202)
- [stateful_service.advertise_options.advertise_on_public.port](resources--workload--reference--group-025.md#canonical-0120012201110120-2320030110233132-3233102311321321-2202232022223030-2133131131013111-3231210131302330-1022020011332122-3111230313032011)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-025.md#canonical-2203002110120031-3203013111221013-1321101003110222-2103121303322022-0021002301311230-0203100203111101-0321302220100020-1320220233023212)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert](resources--workload--reference--group-026.md#canonical-2220320330100322-1330210101012132-0301302331113332-3313313132203223-2330123211123011-0222332113113122-0203030132101203-1231012010033022)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.http_protocol_options](resources--workload--reference--group-026.md#canonical-3111212033013022-0013130201203321-3112121303030312-1000312300110101-1302000111133002-2123102233332323-3310121102112230-0011132102231313)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only](resources--workload--reference--group-027.md#canonical-3332103310011000-1311211312111102-0121211013231233-1322332000022223-0222121321121201-1212103101231101-2101201332230012-2013111011231320)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation](resources--workload--reference--group-027.md#canonical-1030200033313230-3312300021230030-1032121021122312-2023130001230122-3201212000330133-1313231032211023-0100323311112332-1101000222231320)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.proper_case_header_transformation

<a id="canonical-2313200203212000-0010131120301322-1222331120020301-2223121003332213-2333212202003333-1201230203210322-0020330113302023-1032020331230112"></a>

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

<a id="canonical-2131022033023330-0213122222203211-1102103122112030-3121310312221321-1112133332220202-2213211211323231-2312323110132013-3113213001011302"></a>

## Direct properties — proper_case_header_transformation / 322310102301 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3123201322230210-2203202130032012-3110302022203100-3310313120021121-0220201333112033-2232222232030100-0230322233310123-2201201231012020"></a>

## Next pages — proper_case_header_transformation / 322310102301 / 4

- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation](resources--workload--reference--group-027.md#canonical-1030200033313230-3312300021230030-1032121021122312-2023130001230122-3201212000330133-1313231032211023-0100323311112332-1101000222231320)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-2113033122121112-2210020003032101-1203113022113120-1030230321113231-0031031310210330-0320113301220233-2023223013113132-2011023132302123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0313232011011011-0230221220230033-0003313303222211-0010031212310331-1030023031001012-3122311133332100-0331313030123320-0211211021312302"></a>

## stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_v2 — http_protocol_enable_v1_v2 / 203010032030 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-017.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-021.md#canonical-3300101333202310-3223231311312310-2030111033033203-1011103231310023-1313110120311121-1221031312313020-1210302113302033-2110231331000202)
- [stateful_service.advertise_options.advertise_on_public.port](resources--workload--reference--group-025.md#canonical-0120012201110120-2320030110233132-3233102311321321-2202232022223030-2133131131013111-3231210131302330-1022020011332122-3111230313032011)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-025.md#canonical-2203002110120031-3203013111221013-1321101003110222-2103121303322022-0021002301311230-0203100203111101-0321302220100020-1320220233023212)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert](resources--workload--reference--group-026.md#canonical-2220320330100322-1330210101012132-0301302331113332-3313313132203223-2330123211123011-0222332113113122-0203030132101203-1231012010033022)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.http_protocol_options](resources--workload--reference--group-026.md#canonical-3111212033013022-0013130201203321-3112121303030312-1000312300110101-1302000111133002-2123102233332323-3310121102112230-0011132102231313)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_v2

<a id="canonical-2211101303112231-2300121333132003-2133311211112012-3033220203122013-0312131221000102-2213021303133333-0133310033211121-2323210000020322"></a>

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

<a id="canonical-0031221200332212-0333101222203003-0101023222123201-3312311232310100-0021300221001300-2211231010302111-1212112102131330-0121200310113321"></a>

## Direct properties — http_protocol_enable_v1_v2 / 203010032030 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1022111111020330-3203330133120003-0301002021021300-0130233122300230-3213000100013211-2031011322320001-0123032003313322-0321202332021012"></a>

## Next pages — http_protocol_enable_v1_v2 / 203010032030 / 4

- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.http_protocol_options](resources--workload--reference--group-026.md#canonical-3111212033013022-0013130201203321-3112121303030312-1000312300110101-1302000111133002-2123102233332323-3310121102112230-0011132102231313)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-2032023010031211-3013130120220333-0121302223022100-0002211223231120-3123120032122221-3321002012223002-2223302120021232-3301333102312203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0032302222102302-1002103123202202-2113002103103033-0222310030132121-1130233322333322-0330113300302022-0033000110220211-0102320233013220"></a>

## stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v2_only — http_protocol_enable_v2_only / 000212131121 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-017.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-021.md#canonical-3300101333202310-3223231311312310-2030111033033203-1011103231310023-1313110120311121-1221031312313020-1210302113302033-2110231331000202)
- [stateful_service.advertise_options.advertise_on_public.port](resources--workload--reference--group-025.md#canonical-0120012201110120-2320030110233132-3233102311321321-2202232022223030-2133131131013111-3231210131302330-1022020011332122-3111230313032011)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-025.md#canonical-2203002110120031-3203013111221013-1321101003110222-2103121303322022-0021002301311230-0203100203111101-0321302220100020-1320220233023212)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert](resources--workload--reference--group-026.md#canonical-2220320330100322-1330210101012132-0301302331113332-3313313132203223-2330123211123011-0222332113113122-0203030132101203-1231012010033022)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.http_protocol_options](resources--workload--reference--group-026.md#canonical-3111212033013022-0013130201203321-3112121303030312-1000312300110101-1302000111133002-2123102233332323-3310121102112230-0011132102231313)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v2_only

<a id="canonical-0033122212312123-0213003010200122-1302303021230111-0220121130102321-0331120303000123-3320310221322201-3022202131010112-1200012032002312"></a>

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

<a id="canonical-3021320013203130-2011321032030222-3032331113111321-1110133303311223-0011023302330212-2220230121233101-2033120112331230-3203002323202221"></a>

## Direct properties — http_protocol_enable_v2_only / 000212131121 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2021001233130232-3311010101233231-3023023310212012-1221123033130201-2003001123020123-0323222223020233-0321322230003230-3030313012332030"></a>

## Next pages — http_protocol_enable_v2_only / 000212131121 / 4

- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.http_protocol_options](resources--workload--reference--group-026.md#canonical-3111212033013022-0013130201203321-3112121303030312-1000312300110101-1302000111133002-2123102233332323-3310121102112230-0011132102231313)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-1101112113010123-0310031231111212-1320223011302301-2010232232211320-2000100120312320-1231222013230310-1013330031330010-2230323021212212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1212223103011130-3022030010212033-3113130130120002-1321212233013021-1203113300321013-3133312302330001-2213120033321122-3113201133100202"></a>

## stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.no_mtls — no_mtls / 330110231323 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-017.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-021.md#canonical-3300101333202310-3223231311312310-2030111033033203-1011103231310023-1313110120311121-1221031312313020-1210302113302033-2110231331000202)
- [stateful_service.advertise_options.advertise_on_public.port](resources--workload--reference--group-025.md#canonical-0120012201110120-2320030110233132-3233102311321321-2202232022223030-2133131131013111-3231210131302330-1022020011332122-3111230313032011)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-025.md#canonical-2203002110120031-3203013111221013-1321101003110222-2103121303322022-0021002301311230-0203100203111101-0321302220100020-1320220233023212)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert](resources--workload--reference--group-026.md#canonical-2220320330100322-1330210101012132-0301302331113332-3313313132203223-2330123211123011-0222332113113122-0203030132101203-1231012010033022)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.no_mtls

<a id="canonical-1313323221203032-3121121001111310-0012131103211001-0132130332233110-3203023011212223-1311331021231322-0001033212120123-0200121330330113"></a>

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

<a id="canonical-3010032002100310-0221220310232223-0001302333130331-0321120223201202-3033311033211323-1103302302030131-2112302232221023-3213010132333030"></a>

## Direct properties — no_mtls / 330110231323 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3030113302311131-0102303130122023-2202213102233131-1022331230313201-3233001223311132-1123230000003130-2323000212300311-2322122000001310"></a>

## Next pages — no_mtls / 330110231323 / 4

- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert](resources--workload--reference--group-026.md#canonical-2220320330100322-1330210101012132-0301302331113332-3313313132203223-2330123211123011-0222332113113122-0203030132101203-1231012010033022)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-2213221202020111-1230222220230033-2233220322223202-0002031231323322-2110220031333031-1202223320223203-2023323103203222-1310012303101112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1003032321130231-0300330313013023-1301020010011211-1302110111322303-3102011132131232-0122223022033302-2211322133012113-2102033100323031"></a>

## stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.non_default_loadbalancer — non_default_loadbalancer / 032201202300 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-017.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-021.md#canonical-3300101333202310-3223231311312310-2030111033033203-1011103231310023-1313110120311121-1221031312313020-1210302113302033-2110231331000202)
- [stateful_service.advertise_options.advertise_on_public.port](resources--workload--reference--group-025.md#canonical-0120012201110120-2320030110233132-3233102311321321-2202232022223030-2133131131013111-3231210131302330-1022020011332122-3111230313032011)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-025.md#canonical-2203002110120031-3203013111221013-1321101003110222-2103121303322022-0021002301311230-0203100203111101-0321302220100020-1320220233023212)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert](resources--workload--reference--group-026.md#canonical-2220320330100322-1330210101012132-0301302331113332-3313313132203223-2330123211123011-0222332113113122-0203030132101203-1231012010033022)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.non_default_loadbalancer

<a id="canonical-2123132031213331-0221232100122223-3231131233003120-3230321030203020-2002311221003322-2322002122031000-2312201323203112-1103102102211211"></a>

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

<a id="canonical-1333201110002111-2332331313233303-3311030012210001-3003303321022322-3130332303012020-1323111212232032-1301212130112130-2220221202120120"></a>

## Direct properties — non_default_loadbalancer / 032201202300 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2020301033021201-2310301323232030-2120330333013222-0003331123203320-3000033112300032-3310002022230202-0000313133232032-0011130200201023"></a>

## Next pages — non_default_loadbalancer / 032201202300 / 4

- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert](resources--workload--reference--group-026.md#canonical-2220320330100322-1330210101012132-0301302331113332-3313313132203223-2330123211123011-0222332113113122-0203030132101203-1231012010033022)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-1210221231112112-2312202201021312-0031013011120123-2011312030013201-0020033201103032-3310112223133211-0131332330032312-0201010032013222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3032130222001013-1000110311100002-0210300231221121-1002331311320333-3322013032013330-1303112121313122-3201210123300213-2131103022011100"></a>

## stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.pass_through — pass_through / 230130000122 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-017.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-021.md#canonical-3300101333202310-3223231311312310-2030111033033203-1011103231310023-1313110120311121-1221031312313020-1210302113302033-2110231331000202)
- [stateful_service.advertise_options.advertise_on_public.port](resources--workload--reference--group-025.md#canonical-0120012201110120-2320030110233132-3233102311321321-2202232022223030-2133131131013111-3231210131302330-1022020011332122-3111230313032011)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-025.md#canonical-2203002110120031-3203013111221013-1321101003110222-2103121303322022-0021002301311230-0203100203111101-0321302220100020-1320220233023212)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert](resources--workload--reference--group-026.md#canonical-2220320330100322-1330210101012132-0301302331113332-3313313132203223-2330123211123011-0222332113113122-0203030132101203-1231012010033022)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.pass_through

<a id="canonical-2223130122301021-1023220213121201-0233303001333203-2300011310031211-1203221003200201-2212111131010201-1112002013200012-1323201223021013"></a>

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

<a id="canonical-3133012111011012-2311222233021222-1100001323121320-1332332023212321-2113202200021200-1311200111020121-0330310231233122-1121020202322223"></a>

## Direct properties — pass_through / 230130000122 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1122000111221213-2002210222121110-0212222203132113-0212212121033003-3232301301010321-3321220220101003-1320003113033101-2320233202333031"></a>

## Next pages — pass_through / 230130000122 / 4

- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert](resources--workload--reference--group-026.md#canonical-2220320330100322-1330210101012132-0301302331113332-3313313132203223-2330123211123011-0222332113113122-0203030132101203-1231012010033022)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-0310210102013132-3210033020300001-1012032303002020-3312202003132121-1201330230102320-2300111322301332-2323200221333120-1102213222113300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0032231132012100-0201221231231232-2201131231122100-3303120220002111-3333201132312203-1110331023313223-0213310023133312-2203032103101003"></a>

## stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.tls_config — tls_config / 322032333031 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-017.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-021.md#canonical-3300101333202310-3223231311312310-2030111033033203-1011103231310023-1313110120311121-1221031312313020-1210302113302033-2110231331000202)
- [stateful_service.advertise_options.advertise_on_public.port](resources--workload--reference--group-025.md#canonical-0120012201110120-2320030110233132-3233102311321321-2202232022223030-2133131131013111-3231210131302330-1022020011332122-3111230313032011)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-025.md#canonical-2203002110120031-3203013111221013-1321101003110222-2103121303322022-0021002301311230-0203100203111101-0321302220100020-1320220233023212)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert](resources--workload--reference--group-026.md#canonical-2220320330100322-1330210101012132-0301302331113332-3313313132203223-2330123211123011-0222332113113122-0203030132101203-1231012010033022)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.tls_config

<a id="canonical-1003203020210201-0012321012313230-3200010222010033-3021021003331201-0302131011233023-1120021331201102-1100301122110032-2103032010013132"></a>

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

<a id="canonical-2033110200013203-2122101221313212-1111110211111223-2031030223221301-1301332213022010-3111013013302310-1102322230120012-1202323021013223"></a>

## Direct properties — tls_config / 322032333031 / 3

- [custom_security](resources--workload--reference--group-027.md#canonical-1022232031033113-0333013233200110-2021011102131110-1032203003230131-3200312130221102-0122121332200201-1321212112003102-2211323333211213): complete subsection reference.

- [default_security](resources--workload--reference--group-027.md#canonical-0230012300220001-0032001312133030-1231233122323301-0113220011133023-2101210122000313-1013130220113300-0322332220213212-2010303010113222): complete subsection reference.

- [low_security](resources--workload--reference--group-027.md#canonical-0332112101103122-3033313112123020-0023133013113130-1331201201303210-3031132001301302-0103302331121233-0333311211313312-1301213312112022): complete subsection reference.

- [medium_security](resources--workload--reference--group-027.md#canonical-3300200023030022-1333021221001122-0231202113220012-2003212331133220-2210100333302300-0121332112120002-0113221110212130-1311010302020301): complete subsection reference.

<a id="canonical-2031302310313323-0202103130320311-3010202102031032-2312232111101303-3002222032122230-1011301113232100-2020021331212133-0301103201121310"></a>

## Next pages — tls_config / 322032333031 / 4

- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.tls_config.custom_security](resources--workload--reference--group-027.md#canonical-1022232031033113-0333013233200110-2021011102131110-1032203003230131-3200312130221102-0122121332200201-1321212112003102-2211323333211213)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.tls_config.default_security](resources--workload--reference--group-027.md#canonical-0230012300220001-0032001312133030-1231233122323301-0113220011133023-2101210122000313-1013130220113300-0322332220213212-2010303010113222)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.tls_config.low_security](resources--workload--reference--group-027.md#canonical-0332112101103122-3033313112123020-0023133013113130-1331201201303210-3031132001301302-0103302331121233-0333311211313312-1301213312112022)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.tls_config.medium_security](resources--workload--reference--group-027.md#canonical-3300200023030022-1333021221001122-0231202113220012-2003212331133220-2210100333302300-0121332112120002-0113221110212130-1311010302020301)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert](resources--workload--reference--group-026.md#canonical-2220320330100322-1330210101012132-0301302331113332-3313313132203223-2330123211123011-0222332113113122-0203030132101203-1231012010033022)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-1022232031033113-0333013233200110-2021011102131110-1032203003230131-3200312130221102-0122121332200201-1321212112003102-2211323333211213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0331213113320032-0112112332310130-1331120023230013-1230323220203002-0102000030300223-2021221232022031-1323030122103301-1102203311312233"></a>

## stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.tls_config.custom_security — custom_security / 303303121310 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-017.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-021.md#canonical-3300101333202310-3223231311312310-2030111033033203-1011103231310023-1313110120311121-1221031312313020-1210302113302033-2110231331000202)
- [stateful_service.advertise_options.advertise_on_public.port](resources--workload--reference--group-025.md#canonical-0120012201110120-2320030110233132-3233102311321321-2202232022223030-2133131131013111-3231210131302330-1022020011332122-3111230313032011)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-025.md#canonical-2203002110120031-3203013111221013-1321101003110222-2103121303322022-0021002301311230-0203100203111101-0321302220100020-1320220233023212)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert](resources--workload--reference--group-026.md#canonical-2220320330100322-1330210101012132-0301302331113332-3313313132203223-2330123211123011-0222332113113122-0203030132101203-1231012010033022)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.tls_config](resources--workload--reference--group-027.md#canonical-0310210102013132-3210033020300001-1012032303002020-3312202003132121-1201330230102320-2300111322301332-2323200221333120-1102213222113300)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.tls_config.custom_security

<a id="canonical-0000321210212113-1210223033333002-1101203030123311-3112203221133201-2310331012200300-1210303310210020-2132101010323022-0333232113022333"></a>

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

<a id="canonical-1221000213020100-3112323310312222-1103010211130303-0200301013320111-1323103200110200-1110132121213213-3323033103310333-0111120232133213"></a>

## Direct properties — custom_security / 303303121310 / 3

<a id="canonical-1130123030331102-2022130233110013-3132110222022023-1202012110300130-3203002101202131-1232033310332101-3233111303032022-3330011301000101"></a>

<a id="canonical-0320102220323100-3232331021111111-0002211302231031-0100301231103203-2313100310302310-2330133022330023-3233203222332033-0311200133331102"></a>

## cipher_suites property — custom_security / 303303121310 / 4

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

<a id="canonical-3213102012032332-0313332003331230-0203012333011111-1002333220203203-0222202230010133-3230023000012221-1133303230230001-3201233320103121"></a>

<a id="canonical-2211133123313222-0132031011221321-1322333201332220-3113003123103122-1203332310222112-2111302203033300-1031001113232112-0103223330121113"></a>

## max_version property — custom_security / 303303121310 / 5

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

<a id="canonical-2020202113213112-3321122202033232-0022302013030112-1212310101201232-2113032122030022-0003101230032230-2332001002022023-2232211003133002"></a>

<a id="canonical-3121031201121233-0013133122020333-3232111102201203-3331312121312203-2113131131230030-3310133030230030-0323213200233110-2332023110211331"></a>

## min_version property — custom_security / 303303121310 / 6

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

<a id="canonical-0133301300233030-3031203322102121-3320123332200330-3302230333130323-1303120210311331-0111212113321233-3020010312120032-2203322112210303"></a>

## Next pages — custom_security / 303303121310 / 7

- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.tls_config](resources--workload--reference--group-027.md#canonical-0310210102013132-3210033020300001-1012032303002020-3312202003132121-1201330230102320-2300111322301332-2323200221333120-1102213222113300)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-0230012300220001-0032001312133030-1231233122323301-0113220011133023-2101210122000313-1013130220113300-0322332220213212-2010303010113222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3132331222110101-1313330303212203-1111212330332021-0112331330201000-2303210103012111-0020203033033110-3000013103122011-0223133002323301"></a>

## stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.tls_config.default_security — default_security / 022332321221 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-017.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-021.md#canonical-3300101333202310-3223231311312310-2030111033033203-1011103231310023-1313110120311121-1221031312313020-1210302113302033-2110231331000202)
- [stateful_service.advertise_options.advertise_on_public.port](resources--workload--reference--group-025.md#canonical-0120012201110120-2320030110233132-3233102311321321-2202232022223030-2133131131013111-3231210131302330-1022020011332122-3111230313032011)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-025.md#canonical-2203002110120031-3203013111221013-1321101003110222-2103121303322022-0021002301311230-0203100203111101-0321302220100020-1320220233023212)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert](resources--workload--reference--group-026.md#canonical-2220320330100322-1330210101012132-0301302331113332-3313313132203223-2330123211123011-0222332113113122-0203030132101203-1231012010033022)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.tls_config](resources--workload--reference--group-027.md#canonical-0310210102013132-3210033020300001-1012032303002020-3312202003132121-1201330230102320-2300111322301332-2323200221333120-1102213222113300)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.tls_config.default_security

<a id="canonical-0033031022222323-2203111020100000-1013122103332020-2121011221130301-3333120032320032-0231301331311021-1312031201323032-0320101312000033"></a>

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

<a id="canonical-3223222003203312-2122311111021132-2323232303133100-3032122322011131-1031003331010013-3013120023331011-3220121130222002-3121012230003012"></a>

## Direct properties — default_security / 022332321221 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1132313011133001-3201311203313201-1310220330003213-3112101020101210-3320220112333110-0222233202002101-3012020131230122-0001320120223213"></a>

## Next pages — default_security / 022332321221 / 4

- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.tls_config](resources--workload--reference--group-027.md#canonical-0310210102013132-3210033020300001-1012032303002020-3312202003132121-1201330230102320-2300111322301332-2323200221333120-1102213222113300)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-0332112101103122-3033313112123020-0023133013113130-1331201201303210-3031132001301302-0103302331121233-0333311211313312-1301213312112022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3010210210122110-0211203210020302-2312300010001301-1300021023301333-1213201313203020-0032130103322100-2321201310321202-0100110331332030"></a>

## stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.tls_config.low_security — low_security / 222323332031 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-017.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-021.md#canonical-3300101333202310-3223231311312310-2030111033033203-1011103231310023-1313110120311121-1221031312313020-1210302113302033-2110231331000202)
- [stateful_service.advertise_options.advertise_on_public.port](resources--workload--reference--group-025.md#canonical-0120012201110120-2320030110233132-3233102311321321-2202232022223030-2133131131013111-3231210131302330-1022020011332122-3111230313032011)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-025.md#canonical-2203002110120031-3203013111221013-1321101003110222-2103121303322022-0021002301311230-0203100203111101-0321302220100020-1320220233023212)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert](resources--workload--reference--group-026.md#canonical-2220320330100322-1330210101012132-0301302331113332-3313313132203223-2330123211123011-0222332113113122-0203030132101203-1231012010033022)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.tls_config](resources--workload--reference--group-027.md#canonical-0310210102013132-3210033020300001-1012032303002020-3312202003132121-1201330230102320-2300111322301332-2323200221333120-1102213222113300)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.tls_config.low_security

<a id="canonical-1112001320000303-2320321012001310-2232301020310330-3133311321322320-3001033002202313-0031303333223122-1313311210102231-3112110010032312"></a>

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

<a id="canonical-0132112013201322-2012321130111112-0032320303212222-3233120012233032-0031302301002002-2202220130332201-2210223110312212-2020100002100023"></a>

## Direct properties — low_security / 222323332031 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3000332332211202-0123122200231020-2223332233133210-3112110111223330-0232022332030201-0321320000011012-1223222201332331-3223200013202333"></a>

## Next pages — low_security / 222323332031 / 4

- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.tls_config](resources--workload--reference--group-027.md#canonical-0310210102013132-3210033020300001-1012032303002020-3312202003132121-1201330230102320-2300111322301332-2323200221333120-1102213222113300)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-3300200023030022-1333021221001122-0231202113220012-2003212331133220-2210100333302300-0121332112120002-0113221110212130-1311010302020301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1300112101131310-0012330101123230-1103021323022003-3121133120201013-1032121233210001-3322120021211313-3201302223033203-3010323330303030"></a>

## stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.tls_config.medium_security — medium_security / 013230200313 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-017.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-021.md#canonical-3300101333202310-3223231311312310-2030111033033203-1011103231310023-1313110120311121-1221031312313020-1210302113302033-2110231331000202)
- [stateful_service.advertise_options.advertise_on_public.port](resources--workload--reference--group-025.md#canonical-0120012201110120-2320030110233132-3233102311321321-2202232022223030-2133131131013111-3231210131302330-1022020011332122-3111230313032011)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-025.md#canonical-2203002110120031-3203013111221013-1321101003110222-2103121303322022-0021002301311230-0203100203111101-0321302220100020-1320220233023212)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert](resources--workload--reference--group-026.md#canonical-2220320330100322-1330210101012132-0301302331113332-3313313132203223-2330123211123011-0222332113113122-0203030132101203-1231012010033022)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.tls_config](resources--workload--reference--group-027.md#canonical-0310210102013132-3210033020300001-1012032303002020-3312202003132121-1201330230102320-2300111322301332-2323200221333120-1102213222113300)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.tls_config.medium_security

<a id="canonical-1200301321330000-2131202333201112-0210132332333233-3021121001303303-3113232030110111-1110332120320132-3020133303101113-1330310323301030"></a>

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

<a id="canonical-2331003122002300-3211011301010203-3033232031311302-3332013223033232-1100203220123010-2122220112220132-0333120110230220-1031210103100012"></a>

## Direct properties — medium_security / 013230200313 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3031002333031021-0301103010213222-1203003102100200-0103011012112232-1011312013000303-1210302301130300-3312202111211332-2010122213203102"></a>

## Next pages — medium_security / 013230200313 / 4

- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.tls_config](resources--workload--reference--group-027.md#canonical-0310210102013132-3210033020300001-1012032303002020-3312202003132121-1201330230102320-2300111322301332-2323200221333120-1102213222113300)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-2333133220013023-2213202103321302-2200200031101200-1110110132122230-2112303211121122-3013330333300103-3000201200111130-2002030023333122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2123111213202333-2322132020032122-3212301321302221-3000200001032122-1310000310213032-1302122013200133-1102320222021213-1121230020003330"></a>

## stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.use_mtls — use_mtls / 023302232330 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-017.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-021.md#canonical-3300101333202310-3223231311312310-2030111033033203-1011103231310023-1313110120311121-1221031312313020-1210302113302033-2110231331000202)
- [stateful_service.advertise_options.advertise_on_public.port](resources--workload--reference--group-025.md#canonical-0120012201110120-2320030110233132-3233102311321321-2202232022223030-2133131131013111-3231210131302330-1022020011332122-3111230313032011)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-025.md#canonical-2203002110120031-3203013111221013-1321101003110222-2103121303322022-0021002301311230-0203100203111101-0321302220100020-1320220233023212)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert](resources--workload--reference--group-026.md#canonical-2220320330100322-1330210101012132-0301302331113332-3313313132203223-2330123211123011-0222332113113122-0203030132101203-1231012010033022)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.use_mtls

<a id="canonical-3321121132231013-3102130011311120-1301202221120102-0312313212102030-2000033333012132-3302301213310223-1200001110012123-1322130222003220"></a>

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

<a id="canonical-2312010112322220-2000330332213231-1123111120000112-1321322231301001-1231010011121102-3123132113022112-1010311103301120-1313331211322220"></a>

## Direct properties — use_mtls / 023302232330 / 3

<a id="canonical-3022230003322010-1310112132032002-2232000321211311-3322310332010132-2322221120221102-1302232211023111-2211212121230232-3002232321100022"></a>

<a id="canonical-1023310130322300-3302022130310321-0110310313012021-1132323100302200-1033233322300012-3010300103223311-0013021222200223-1312322013310111"></a>

## client_certificate_optional property — use_mtls / 023302232330 / 4

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

- [crl](resources--workload--reference--group-027.md#canonical-3213131122113303-2211101302100231-3332202130130031-3330311231120022-2313033321113000-2101112102013233-2310302313131332-3113221010301121): complete subsection reference.

- [no_crl](resources--workload--reference--group-027.md#canonical-2301132021033130-2020010200203112-2120302331033211-1031013320121032-0323111301011033-3231320011302031-2100233020331213-2003131221313023): complete subsection reference.

- [trusted_ca](resources--workload--reference--group-027.md#canonical-1033100210320033-3213330302001210-3310200231020103-1222200233302232-0223030203131303-2133223033022033-3223110313032333-0000030323322001): complete subsection reference.

<a id="canonical-1010012221012110-1011032220103013-3130312221031221-0202222123333221-3102212220011213-2012122210211113-3012121112120231-2013312233000200"></a>

<a id="canonical-2330103001232031-1313333311030332-3100301220000102-0122312032322023-2211133333031121-0323233311003223-2311221233030110-0331011202203233"></a>

## trusted_ca_url property — use_mtls / 023302232330 / 5

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

- [xfcc_disabled](resources--workload--reference--group-027.md#canonical-1311000012110300-3021300302022121-0000332110111020-3022011202032132-3332303330133301-0230122121130100-1032033223223233-1120121211030111): complete subsection reference.

- [xfcc_options](resources--workload--reference--group-027.md#canonical-3300030331122223-1100201033312313-3322122222110310-3301320232031232-0220110220002302-0330111233101112-0033232132302121-2203330131000032): complete subsection reference.

<a id="canonical-2120310221322021-3031230002110031-3113130332112223-3011310320230223-0232333303002011-1103220233301301-2312123120201100-1101100300312010"></a>

## Next pages — use_mtls / 023302232330 / 6

- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.use_mtls.crl](resources--workload--reference--group-027.md#canonical-3213131122113303-2211101302100231-3332202130130031-3330311231120022-2313033321113000-2101112102013233-2310302313131332-3113221010301121)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.use_mtls.no_crl](resources--workload--reference--group-027.md#canonical-2301132021033130-2020010200203112-2120302331033211-1031013320121032-0323111301011033-3231320011302031-2100233020331213-2003131221313023)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.use_mtls.trusted_ca](resources--workload--reference--group-027.md#canonical-1033100210320033-3213330302001210-3310200231020103-1222200233302232-0223030203131303-2133223033022033-3223110313032333-0000030323322001)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.use_mtls.xfcc_disabled](resources--workload--reference--group-027.md#canonical-1311000012110300-3021300302022121-0000332110111020-3022011202032132-3332303330133301-0230122121130100-1032033223223233-1120121211030111)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.use_mtls.xfcc_options](resources--workload--reference--group-027.md#canonical-3300030331122223-1100201033312313-3322122222110310-3301320232031232-0220110220002302-0330111233101112-0033232132302121-2203330131000032)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert](resources--workload--reference--group-026.md#canonical-2220320330100322-1330210101012132-0301302331113332-3313313132203223-2330123211123011-0222332113113122-0203030132101203-1231012010033022)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-3213131122113303-2211101302100231-3332202130130031-3330311231120022-2313033321113000-2101112102013233-2310302313131332-3113221010301121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2313022213002333-0232021120133131-0122122020023200-3323200232120223-2113313010013201-0121120110100201-2310211323030200-0211333202223221"></a>

## stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.use_mtls.crl — crl / 010233002032 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-017.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-021.md#canonical-3300101333202310-3223231311312310-2030111033033203-1011103231310023-1313110120311121-1221031312313020-1210302113302033-2110231331000202)
- [stateful_service.advertise_options.advertise_on_public.port](resources--workload--reference--group-025.md#canonical-0120012201110120-2320030110233132-3233102311321321-2202232022223030-2133131131013111-3231210131302330-1022020011332122-3111230313032011)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-025.md#canonical-2203002110120031-3203013111221013-1321101003110222-2103121303322022-0021002301311230-0203100203111101-0321302220100020-1320220233023212)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert](resources--workload--reference--group-026.md#canonical-2220320330100322-1330210101012132-0301302331113332-3313313132203223-2330123211123011-0222332113113122-0203030132101203-1231012010033022)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.use_mtls](resources--workload--reference--group-027.md#canonical-2333133220013023-2213202103321302-2200200031101200-1110110132122230-2112303211121122-3013330333300103-3000201200111130-2002030023333122)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.use_mtls.crl

<a id="canonical-0232310210313323-1000210332212210-2210023213313131-3332310202330213-0211110312221311-2131120021010201-0020101133201301-1313232122020012"></a>

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

<a id="canonical-0103131302210122-3120100320202231-1003101013323001-1003320133200211-1300212230202220-1012001321200201-3120210232310233-0020303310003031"></a>

## Direct properties — crl / 010233002032 / 3

<a id="canonical-3123012221323001-2131203323220221-0010212022031130-3311013112012213-0002012011303202-0130133102323121-3132101100131102-0103131123332223"></a>

<a id="canonical-2123331100312203-2233100231320120-3220012321100033-1213002203333233-3103022202300210-1211313023223232-1332203222022322-3123233122210210"></a>

## name property — crl / 010233002032 / 4

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

<a id="canonical-2130011221132131-2130123111311120-3102202112100322-0101112323130010-0101122120201002-3023302310020232-3302221230032213-3201011102020130"></a>

<a id="canonical-0132222212320333-3010013012011221-1010111103302321-2332231321320103-3302120121133230-3213000113323120-0000121102010113-0033100232200120"></a>

## namespace property — crl / 010233002032 / 5

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

<a id="canonical-0110321130010300-2113111011231333-3202020230120301-2001012111202233-3120301330333212-3303312203202201-1022211012123103-2322121013123332"></a>

<a id="canonical-0323323321102132-3133031102023332-3122310233210200-2020220203310020-0011011121031000-1013010020221003-3132320332133023-0132033231230003"></a>

## tenant property — crl / 010233002032 / 6

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

<a id="canonical-1300011003311210-0213030311330310-1100133101202303-3213323310103130-2032203303302003-2331320230330233-1112330101013301-2132000301020132"></a>

## Next pages — crl / 010233002032 / 7

- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.use_mtls](resources--workload--reference--group-027.md#canonical-2333133220013023-2213202103321302-2200200031101200-1110110132122230-2112303211121122-3013330333300103-3000201200111130-2002030023333122)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-2301132021033130-2020010200203112-2120302331033211-1031013320121032-0323111301011033-3231320011302031-2100233020331213-2003131221313023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0110323133133212-2330310312132031-3221120302003330-3231231123101230-0210130313110222-0200121003203011-3323133302302223-3112012223221300"></a>

## stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.use_mtls.no_crl — no_crl / 123003310213 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-017.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-021.md#canonical-3300101333202310-3223231311312310-2030111033033203-1011103231310023-1313110120311121-1221031312313020-1210302113302033-2110231331000202)
- [stateful_service.advertise_options.advertise_on_public.port](resources--workload--reference--group-025.md#canonical-0120012201110120-2320030110233132-3233102311321321-2202232022223030-2133131131013111-3231210131302330-1022020011332122-3111230313032011)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-025.md#canonical-2203002110120031-3203013111221013-1321101003110222-2103121303322022-0021002301311230-0203100203111101-0321302220100020-1320220233023212)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert](resources--workload--reference--group-026.md#canonical-2220320330100322-1330210101012132-0301302331113332-3313313132203223-2330123211123011-0222332113113122-0203030132101203-1231012010033022)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.use_mtls](resources--workload--reference--group-027.md#canonical-2333133220013023-2213202103321302-2200200031101200-1110110132122230-2112303211121122-3013330333300103-3000201200111130-2002030023333122)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.use_mtls.no_crl

<a id="canonical-0313322012222123-1221130120133322-1220022322223031-0031022302030030-0031230212211220-3011223230000001-3003021220303111-3123201323303023"></a>

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

<a id="canonical-1220120110131103-0222020132312231-1032301020023203-0130232130100133-0202002020202211-0303012302032112-1002312203103302-1130030321303011"></a>

## Direct properties — no_crl / 123003310213 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1323331213012310-2311131001113122-3102100313032032-1010022301102332-1203030031013302-1301203230132213-2333322020030131-0103300233112122"></a>

## Next pages — no_crl / 123003310213 / 4

- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.use_mtls](resources--workload--reference--group-027.md#canonical-2333133220013023-2213202103321302-2200200031101200-1110110132122230-2112303211121122-3013330333300103-3000201200111130-2002030023333122)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-1033100210320033-3213330302001210-3310200231020103-1222200233302232-0223030203131303-2133223033022033-3223110313032333-0000030323322001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3031301323203230-2322201121003021-1223303201323233-0002133321213120-3111102022000330-1313203133330221-3111030200021013-3133203311222020"></a>

## stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.use_mtls.trusted_ca — trusted_ca / 300313211012 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-017.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-021.md#canonical-3300101333202310-3223231311312310-2030111033033203-1011103231310023-1313110120311121-1221031312313020-1210302113302033-2110231331000202)
- [stateful_service.advertise_options.advertise_on_public.port](resources--workload--reference--group-025.md#canonical-0120012201110120-2320030110233132-3233102311321321-2202232022223030-2133131131013111-3231210131302330-1022020011332122-3111230313032011)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-025.md#canonical-2203002110120031-3203013111221013-1321101003110222-2103121303322022-0021002301311230-0203100203111101-0321302220100020-1320220233023212)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert](resources--workload--reference--group-026.md#canonical-2220320330100322-1330210101012132-0301302331113332-3313313132203223-2330123211123011-0222332113113122-0203030132101203-1231012010033022)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.use_mtls](resources--workload--reference--group-027.md#canonical-2333133220013023-2213202103321302-2200200031101200-1110110132122230-2112303211121122-3013330333300103-3000201200111130-2002030023333122)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.use_mtls.trusted_ca

<a id="canonical-2022131132213102-2013002113103310-0231031112030020-0001101023100111-1223312201220112-1023202103330233-0230132230130310-3321331002301112"></a>

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

<a id="canonical-2100120221003332-0310110233330213-2111012331132330-3322232112113233-0310221102330311-2000321002221303-0301220031110010-0320120312320131"></a>

## Direct properties — trusted_ca / 300313211012 / 3

<a id="canonical-2212232320132300-1123223213122222-3023221200032200-1032022233003301-1120213010120231-0033211233120021-1312322133330110-3331110101033311"></a>

<a id="canonical-3011313011312312-2231001122302323-1030131121033201-2130303102131113-3200002101000321-1303311103031333-0121213321001133-0303332223110130"></a>

## name property — trusted_ca / 300313211012 / 4

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

<a id="canonical-1330221123022323-2133021123031202-0213133122310300-3000121023130331-2310000012113103-1331002223033223-0302200232131232-1322131003210032"></a>

<a id="canonical-3103003123000013-2220301102112020-0133200332120212-0002023321021203-3322211333212131-0020321103200330-2220102122011100-3302312200111021"></a>

## namespace property — trusted_ca / 300313211012 / 5

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

<a id="canonical-0131022110202330-1010231023113111-2222103002020033-3200201313133023-2303101330332102-1100010011330131-3033012022011200-3023211030301313"></a>

<a id="canonical-1102201030330303-1030322313210303-2320302132120330-0112103332132210-1013222102110131-0330122001120023-3302211213221302-3132200011023201"></a>

## tenant property — trusted_ca / 300313211012 / 6

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

<a id="canonical-0003302310031212-3320123201103303-1131002021113332-3121001310213130-1201333301101330-3200332021000121-3221103131011333-2120012022110131"></a>

## Next pages — trusted_ca / 300313211012 / 7

- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.use_mtls](resources--workload--reference--group-027.md#canonical-2333133220013023-2213202103321302-2200200031101200-1110110132122230-2112303211121122-3013330333300103-3000201200111130-2002030023333122)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-1311000012110300-3021300302022121-0000332110111020-3022011202032132-3332303330133301-0230122121130100-1032033223223233-1120121211030111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0000103121231323-3001001111123232-1321232201111321-2022330123110300-3121110002201211-1323011110032022-2302222231211100-2203201033131232"></a>

## stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.use_mtls.xfcc_disabled — xfcc_disabled / 000133133130 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-017.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-021.md#canonical-3300101333202310-3223231311312310-2030111033033203-1011103231310023-1313110120311121-1221031312313020-1210302113302033-2110231331000202)
- [stateful_service.advertise_options.advertise_on_public.port](resources--workload--reference--group-025.md#canonical-0120012201110120-2320030110233132-3233102311321321-2202232022223030-2133131131013111-3231210131302330-1022020011332122-3111230313032011)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-025.md#canonical-2203002110120031-3203013111221013-1321101003110222-2103121303322022-0021002301311230-0203100203111101-0321302220100020-1320220233023212)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert](resources--workload--reference--group-026.md#canonical-2220320330100322-1330210101012132-0301302331113332-3313313132203223-2330123211123011-0222332113113122-0203030132101203-1231012010033022)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.use_mtls](resources--workload--reference--group-027.md#canonical-2333133220013023-2213202103321302-2200200031101200-1110110132122230-2112303211121122-3013330333300103-3000201200111130-2002030023333122)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.use_mtls.xfcc_disabled

<a id="canonical-1111120000100121-3032020303322033-0031230233111321-3302002010300002-3112211123213220-3313020232120211-2302103020102130-3323231012333022"></a>

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

<a id="canonical-3022132212022020-3223133101120322-2133201102223223-3003210202013022-1312332012132302-3302302321210100-0301201122030330-2122103030301002"></a>

## Direct properties — xfcc_disabled / 000133133130 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0211023013033230-2300001003033223-3332320330032300-2332122022300132-0033011231323313-1111110201312013-3332223110120200-2100303023020231"></a>

## Next pages — xfcc_disabled / 000133133130 / 4

- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.use_mtls](resources--workload--reference--group-027.md#canonical-2333133220013023-2213202103321302-2200200031101200-1110110132122230-2112303211121122-3013330333300103-3000201200111130-2002030023333122)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-3300030331122223-1100201033312313-3322122222110310-3301320232031232-0220110220002302-0330111233101112-0033232132302121-2203330131000032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2133023303133220-1233210322021111-2133112011100230-0333303033230120-3322011222221112-0231023312013313-1203320000031321-1130123302221333"></a>

## stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.use_mtls.xfcc_options — xfcc_options / 300132023000 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-017.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-021.md#canonical-3300101333202310-3223231311312310-2030111033033203-1011103231310023-1313110120311121-1221031312313020-1210302113302033-2110231331000202)
- [stateful_service.advertise_options.advertise_on_public.port](resources--workload--reference--group-025.md#canonical-0120012201110120-2320030110233132-3233102311321321-2202232022223030-2133131131013111-3231210131302330-1022020011332122-3111230313032011)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-025.md#canonical-2203002110120031-3203013111221013-1321101003110222-2103121303322022-0021002301311230-0203100203111101-0321302220100020-1320220233023212)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert](resources--workload--reference--group-026.md#canonical-2220320330100322-1330210101012132-0301302331113332-3313313132203223-2330123211123011-0222332113113122-0203030132101203-1231012010033022)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.use_mtls](resources--workload--reference--group-027.md#canonical-2333133220013023-2213202103321302-2200200031101200-1110110132122230-2112303211121122-3013330333300103-3000201200111130-2002030023333122)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.use_mtls.xfcc_options

<a id="canonical-1003302023031033-2231100322330230-2331203020231301-0213220330021011-2202301233303302-2103100033133011-0020131102302102-3312010101023023"></a>

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

<a id="canonical-1133203323111022-0222202311201112-1132021323232210-0231202132303300-3003132020002330-0222101101201221-1010103132100323-1003230111313002"></a>

## Direct properties — xfcc_options / 300132023000 / 3

<a id="canonical-3202030303203210-1011001101331233-0213333220221030-3113233000022030-1021233002212213-0001102220133201-1010222010222113-1212000113100202"></a>

<a id="canonical-1221002030031033-0232311132323221-0003222033200022-0121011322122111-1212312202022031-0231223032213120-0202200232222323-2020202302203010"></a>

## xfcc_header_elements property — xfcc_options / 300132023000 / 4

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

<a id="canonical-3021133112001303-3030312221031200-3222130321132223-1012312233110033-2233333200223032-1323330101233121-3200223213330122-2202201002000203"></a>

## Next pages — xfcc_options / 300132023000 / 5

- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.use_mtls](resources--workload--reference--group-027.md#canonical-2333133220013023-2213202103321302-2200200031101200-1110110132122230-2112303211121122-3013330333300103-3000201200111130-2002030023333122)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-2300220120202032-2120133330213102-2332100100003133-0101032330232302-2302101320100002-1302333333023222-3232223323232230-0320110220223201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2213211230102022-3110310021031211-0233303003001000-1312022101322333-0131333101230010-2131022031231103-0103203002123101-1112303002033102"></a>

## stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes — specific_routes / 310301031332 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-017.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-021.md#canonical-3300101333202310-3223231311312310-2030111033033203-1011103231310023-1313110120311121-1221031312313020-1210302113302033-2110231331000202)
- [stateful_service.advertise_options.advertise_on_public.port](resources--workload--reference--group-025.md#canonical-0120012201110120-2320030110233132-3233102311321321-2202232022223030-2133131131013111-3231210131302330-1022020011332122-3111230313032011)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-025.md#canonical-2203002110120031-3203013111221013-1321101003110222-2103121303322022-0021002301311230-0203100203111101-0321302220100020-1320220233023212)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes

<a id="canonical-0000223113321033-0322201213201322-0112013210301131-3101333232130311-1320312303102212-1013201303003203-2311330011312213-2331013133021201"></a>

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

<a id="canonical-1232323203321002-0022330323311023-2022022203232031-3102030022310312-2222031133203311-0132211202130020-3021010030221232-1312223211213110"></a>

## Direct properties — specific_routes / 310301031332 / 3

- [routes](resources--workload--reference--group-027.md#canonical-2302003233200103-3312120222212200-1000311013000201-0201011231133032-2032003100312302-3001032323113031-1031222100120212-0102213202320111): complete subsection reference.

<a id="canonical-3021311001213310-3333331133032323-1321002212021022-3321202212022033-3300100312332033-0031121020102130-2211332233211212-1101213222312323"></a>

## Next pages — specific_routes / 310301031332 / 4

- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes](resources--workload--reference--group-027.md#canonical-2302003233200103-3312120222212200-1000311013000201-0201011231133032-2032003100312302-3001032323113031-1031222100120212-0102213202320111)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-025.md#canonical-2203002110120031-3203013111221013-1321101003110222-2103121303322022-0021002301311230-0203100203111101-0321302220100020-1320220233023212)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-2302003233200103-3312120222212200-1000311013000201-0201011231133032-2032003100312302-3001032323113031-1031222100120212-0102213202320111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3231222033222322-1111031221132223-2321222112112313-0321010311032030-2210313331220032-0332103100230320-0003212132223220-3232231132132313"></a>

## stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes — routes / 100301013012 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-017.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-021.md#canonical-3300101333202310-3223231311312310-2030111033033203-1011103231310023-1313110120311121-1221031312313020-1210302113302033-2110231331000202)
- [stateful_service.advertise_options.advertise_on_public.port](resources--workload--reference--group-025.md#canonical-0120012201110120-2320030110233132-3233102311321321-2202232022223030-2133131131013111-3231210131302330-1022020011332122-3111230313032011)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-025.md#canonical-2203002110120031-3203013111221013-1321101003110222-2103121303322022-0021002301311230-0203100203111101-0321302220100020-1320220233023212)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes](resources--workload--reference--group-027.md#canonical-2300220120202032-2120133330213102-2332100100003133-0101032330232302-2302101320100002-1302333333023222-3232223323232230-0320110220223201)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes

<a id="canonical-3212030312102320-2020021230002210-0313233331321122-1320311300013320-0303211203231223-1322322131000110-1133321123323313-3100213123233232"></a>

Type: `"object"`. list nested block, Optional.

Routes. Routes for this loadbalancer.

Upstream description:

Routes for this loadbalancer.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("custom_route_object",
    "direct_response_route"),
  validators.ConflictingListObjectAttributes("custom_route_object",
    "redirect_route"),
  validators.ConflictingListObjectAttributes("custom_route_object",
    "simple_route"),
  validators.ConflictingListObjectAttributes("direct_response_route",
    "redirect_route"),
  validators.ConflictingListObjectAttributes("direct_response_route",
    "simple_route"),
  validators.ConflictingListObjectAttributes("redirect_route",
    "simple_route")}
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

<a id="canonical-2322300210013003-0313330121021211-2232001231033311-0212310320013001-0222311220311322-1133133001301130-1113332120210300-0123112212302032"></a>

## Direct properties — routes / 100301013012 / 3

- [custom_route_object](resources--workload--reference--group-027.md#canonical-3220230303101313-0122201023323032-0302022132031233-3332233000011332-2123202133102020-1133230203021321-3311110223100121-0113002202213112): complete subsection reference.

- [direct_response_route](resources--workload--reference--group-027.md#canonical-1121000311012033-0022112012130003-3211200311100011-1032311002323322-2310120120102220-2103030310121321-2022131002003111-2221012223020220): complete subsection reference.

- [redirect_route](resources--workload--reference--group-027.md#canonical-0330013222233302-2210122030021330-0131201133002112-2132231131331102-2322210303201110-1121010222301222-3302213332031013-0122333212301102): complete subsection reference.

- [simple_route](resources--workload--reference--group-028.md#canonical-3223313232022132-1300023030303322-3120001113102132-2101020221132002-0011333231033012-1032011333210012-0303212231230122-1132022333212200): complete subsection reference.

<a id="canonical-2322030201302321-3210321221133232-0311311010320320-0233203120302133-1320200231101111-2120333310231113-2222302231232221-0203113203032131"></a>

## Next pages — routes / 100301013012 / 4

- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.custom_route_object](resources--workload--reference--group-027.md#canonical-3220230303101313-0122201023323032-0302022132031233-3332233000011332-2123202133102020-1133230203021321-3311110223100121-0113002202213112)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.direct_response_route](resources--workload--reference--group-027.md#canonical-1121000311012033-0022112012130003-3211200311100011-1032311002323322-2310120120102220-2103030310121321-2022131002003111-2221012223020220)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.redirect_route](resources--workload--reference--group-027.md#canonical-0330013222233302-2210122030021330-0131201133002112-2132231131331102-2322210303201110-1121010222301222-3302213332031013-0122333212301102)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.simple_route](resources--workload--reference--group-028.md#canonical-3223313232022132-1300023030303322-3120001113102132-2101020221132002-0011333231033012-1032011333210012-0303212231230122-1132022333212200)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes](resources--workload--reference--group-027.md#canonical-2300220120202032-2120133330213102-2332100100003133-0101032330232302-2302101320100002-1302333333023222-3232223323232230-0320110220223201)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-3220230303101313-0122201023323032-0302022132031233-3332233000011332-2123202133102020-1133230203021321-3311110223100121-0113002202213112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3013332033012331-0331331020000022-3302100001321211-0030211132323120-1022223132001330-1320313133211111-3033021200221333-3121323210011123"></a>

## stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.custom_route_object — custom_route_object / 120211221212 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-017.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-021.md#canonical-3300101333202310-3223231311312310-2030111033033203-1011103231310023-1313110120311121-1221031312313020-1210302113302033-2110231331000202)
- [stateful_service.advertise_options.advertise_on_public.port](resources--workload--reference--group-025.md#canonical-0120012201110120-2320030110233132-3233102311321321-2202232022223030-2133131131013111-3231210131302330-1022020011332122-3111230313032011)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-025.md#canonical-2203002110120031-3203013111221013-1321101003110222-2103121303322022-0021002301311230-0203100203111101-0321302220100020-1320220233023212)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes](resources--workload--reference--group-027.md#canonical-2300220120202032-2120133330213102-2332100100003133-0101032330232302-2302101320100002-1302333333023222-3232223323232230-0320110220223201)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes](resources--workload--reference--group-027.md#canonical-2302003233200103-3312120222212200-1000311013000201-0201011231133032-2032003100312302-3001032323113031-1031222100120212-0102213202320111)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.custom_route_object

<a id="canonical-1322223132100011-2232031123032133-0022300303021130-3031221233012332-2100130120311211-2132333122101000-2002201200022330-2213021202320031"></a>

Type: `"object"`. single nested block, Optional.

Custom route uses a route object created outside of this view.

Upstream description:

A custom route uses a route object created outside of this view.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("caching_disable",
    "caching_inherit")}
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
  "x-ves-oneof-field-caching": "[\"caching_disable\",\"caching_inherit\"]"
}
```

Terraform syntax:

```terraform
custom_route_object {
  # Configure direct properties listed below.
}
```

<a id="canonical-2021313302321330-3123220023210130-0130332102030101-2230021121302232-2330020203022020-3022201103310232-2233311200300330-0312201021230210"></a>

## Direct properties — custom_route_object / 120211221212 / 3

- [caching_disable](resources--workload--reference--group-027.md#canonical-0321331020112113-2221312003120313-3000223222202031-3322213033010313-1221221212232312-1301221220313131-2021320313300001-1303011223112020): complete subsection reference.

- [caching_inherit](resources--workload--reference--group-027.md#canonical-2003331110010111-3301013223230120-0320023103322321-0011200211020130-3112200312302301-1302012203123032-1122100012100000-1023130211002222): complete subsection reference.

- [route_ref](resources--workload--reference--group-027.md#canonical-2310202012101022-3110000233211021-3211130003232201-1221333012332102-0200231302232002-3020101320310313-2320213220313103-0220003113022302): complete subsection reference.

<a id="canonical-0312310202333231-1130012113211030-3111330103130211-1223012213222021-0301020222031123-2111123131102002-0301322131013310-3001031131120333"></a>

## Next pages — custom_route_object / 120211221212 / 4

- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.custom_route_object.caching_disable](resources--workload--reference--group-027.md#canonical-0321331020112113-2221312003120313-3000223222202031-3322213033010313-1221221212232312-1301221220313131-2021320313300001-1303011223112020)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.custom_route_object.caching_inherit](resources--workload--reference--group-027.md#canonical-2003331110010111-3301013223230120-0320023103322321-0011200211020130-3112200312302301-1302012203123032-1122100012100000-1023130211002222)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.custom_route_object.route_ref](resources--workload--reference--group-027.md#canonical-2310202012101022-3110000233211021-3211130003232201-1221333012332102-0200231302232002-3020101320310313-2320213220313103-0220003113022302)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes](resources--workload--reference--group-027.md#canonical-2302003233200103-3312120222212200-1000311013000201-0201011231133032-2032003100312302-3001032323113031-1031222100120212-0102213202320111)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-0321331020112113-2221312003120313-3000223222202031-3322213033010313-1221221212232312-1301221220313131-2021320313300001-1303011223112020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1020001032332123-0233231113030133-1122022222312331-0121312111200322-2322233211203111-1231030311320312-0201112331222323-0203110312212100"></a>

## stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.custom_route_object.caching_disable — caching_disable / 203010132321 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-017.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-021.md#canonical-3300101333202310-3223231311312310-2030111033033203-1011103231310023-1313110120311121-1221031312313020-1210302113302033-2110231331000202)
- [stateful_service.advertise_options.advertise_on_public.port](resources--workload--reference--group-025.md#canonical-0120012201110120-2320030110233132-3233102311321321-2202232022223030-2133131131013111-3231210131302330-1022020011332122-3111230313032011)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-025.md#canonical-2203002110120031-3203013111221013-1321101003110222-2103121303322022-0021002301311230-0203100203111101-0321302220100020-1320220233023212)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes](resources--workload--reference--group-027.md#canonical-2300220120202032-2120133330213102-2332100100003133-0101032330232302-2302101320100002-1302333333023222-3232223323232230-0320110220223201)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes](resources--workload--reference--group-027.md#canonical-2302003233200103-3312120222212200-1000311013000201-0201011231133032-2032003100312302-3001032323113031-1031222100120212-0102213202320111)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.custom_route_object](resources--workload--reference--group-027.md#canonical-3220230303101313-0122201023323032-0302022132031233-3332233000011332-2123202133102020-1133230203021321-3311110223100121-0113002202213112)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.custom_route_object.caching_disable

<a id="canonical-0020301130121102-0222002103032100-3002002213332102-1012033130010321-2331321130003320-1133202302013133-0233200010203022-1032221300031101"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for caching disable.

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
caching_disable = {}
```

<a id="canonical-3103002302031230-2013003201112231-1020221301203033-0212003133312000-0313031232132302-3103332212212010-0130312201221113-2013331330123302"></a>

## Direct properties — caching_disable / 203010132321 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3130130111211322-2133232210013323-0130020100231202-1232111332100130-0112303323013211-1112322221121221-3030303110133302-1132131020213103"></a>

## Next pages — caching_disable / 203010132321 / 4

- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.custom_route_object](resources--workload--reference--group-027.md#canonical-3220230303101313-0122201023323032-0302022132031233-3332233000011332-2123202133102020-1133230203021321-3311110223100121-0113002202213112)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-2003331110010111-3301013223230120-0320023103322321-0011200211020130-3112200312302301-1302012203123032-1122100012100000-1023130211002222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0013133101123020-2020030210031012-2013101210112013-1103202032221010-3301011012010020-1210122330131001-2313023200113313-1010203231010213"></a>

## stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.custom_route_object.caching_inherit — caching_inherit / 012231202002 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-017.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-021.md#canonical-3300101333202310-3223231311312310-2030111033033203-1011103231310023-1313110120311121-1221031312313020-1210302113302033-2110231331000202)
- [stateful_service.advertise_options.advertise_on_public.port](resources--workload--reference--group-025.md#canonical-0120012201110120-2320030110233132-3233102311321321-2202232022223030-2133131131013111-3231210131302330-1022020011332122-3111230313032011)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-025.md#canonical-2203002110120031-3203013111221013-1321101003110222-2103121303322022-0021002301311230-0203100203111101-0321302220100020-1320220233023212)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes](resources--workload--reference--group-027.md#canonical-2300220120202032-2120133330213102-2332100100003133-0101032330232302-2302101320100002-1302333333023222-3232223323232230-0320110220223201)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes](resources--workload--reference--group-027.md#canonical-2302003233200103-3312120222212200-1000311013000201-0201011231133032-2032003100312302-3001032323113031-1031222100120212-0102213202320111)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.custom_route_object](resources--workload--reference--group-027.md#canonical-3220230303101313-0122201023323032-0302022132031233-3332233000011332-2123202133102020-1133230203021321-3311110223100121-0113002202213112)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.custom_route_object.caching_inherit

<a id="canonical-0311212011023023-0223023213111101-0121313213001313-2000103131201321-0313202023211222-0222111033323101-1010301032123020-0230330323321131"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for caching inherit.

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
caching_inherit = {}
```

<a id="canonical-2001323230323303-2302210222000210-2331103213213231-0322110020113103-2020332001200301-3102213210111211-2213012223213013-1211310102122220"></a>

## Direct properties — caching_inherit / 012231202002 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1022300120122030-1022222220112102-1201100130020132-3033313222233110-2331303110231312-2302001101220003-0021110113113221-3333331303131211"></a>

## Next pages — caching_inherit / 012231202002 / 4

- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.custom_route_object](resources--workload--reference--group-027.md#canonical-3220230303101313-0122201023323032-0302022132031233-3332233000011332-2123202133102020-1133230203021321-3311110223100121-0113002202213112)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-2310202012101022-3110000233211021-3211130003232201-1221333012332102-0200231302232002-3020101320310313-2320213220313103-0220003113022302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2023120023032320-0301321230113203-1130303223131100-3320321002301101-0020031313020012-0112012100020233-0221021102011310-2223021031120023"></a>

## stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.custom_route_object.route_ref — route_ref / 022300230221 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-017.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-021.md#canonical-3300101333202310-3223231311312310-2030111033033203-1011103231310023-1313110120311121-1221031312313020-1210302113302033-2110231331000202)
- [stateful_service.advertise_options.advertise_on_public.port](resources--workload--reference--group-025.md#canonical-0120012201110120-2320030110233132-3233102311321321-2202232022223030-2133131131013111-3231210131302330-1022020011332122-3111230313032011)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-025.md#canonical-2203002110120031-3203013111221013-1321101003110222-2103121303322022-0021002301311230-0203100203111101-0321302220100020-1320220233023212)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes](resources--workload--reference--group-027.md#canonical-2300220120202032-2120133330213102-2332100100003133-0101032330232302-2302101320100002-1302333333023222-3232223323232230-0320110220223201)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes](resources--workload--reference--group-027.md#canonical-2302003233200103-3312120222212200-1000311013000201-0201011231133032-2032003100312302-3001032323113031-1031222100120212-0102213202320111)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.custom_route_object](resources--workload--reference--group-027.md#canonical-3220230303101313-0122201023323032-0302022132031233-3332233000011332-2123202133102020-1133230203021321-3311110223100121-0113002202213112)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.custom_route_object.route_ref

<a id="canonical-3002202002323332-0011300000222322-3123113023012233-1111321231133000-1231312123021212-3132012231323223-1230003100100001-1231311133213133"></a>

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
route_ref {
  # Configure direct properties listed below.
}
```

<a id="canonical-0120012131212033-1021032003302122-3030030222111102-1211011121131103-2212112030112100-0102011012310022-0013331233223021-3212113302203120"></a>

## Direct properties — route_ref / 022300230221 / 3

<a id="canonical-0102110302110110-1200123311330000-1303011211133122-2302013002322220-2302233213202203-3120033111210310-1120233001311131-1202233111221032"></a>

<a id="canonical-3111123132100211-0111321202201112-1200232203332122-1202130313210230-1303201300002020-2201223033302003-1300103311303302-1033320210100011"></a>

## name property — route_ref / 022300230221 / 4

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

<a id="canonical-0322113102320100-3012133333101230-1211313321013320-0300213002120012-3031321031320002-0020203232301310-2111210232003000-3012133103121210"></a>

<a id="canonical-1020210110122003-1111110130302101-0320112211011012-3032113023131031-2320122133311330-2003233102211212-2131201123333323-0011232102133121"></a>

## namespace property — route_ref / 022300230221 / 5

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

<a id="canonical-3032330300312310-1321331030303230-0001032030133213-2031033211000020-3020020132012021-1120112101130333-0300120300033213-2221301011211101"></a>

<a id="canonical-2030302230212303-0122100030013011-3011200103223211-3211321331300002-3131331120333000-2212330013303303-1311111200303030-0100122123300323"></a>

## tenant property — route_ref / 022300230221 / 6

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

<a id="canonical-1223012210321010-2201000230200000-2233020200130323-2221122111331133-3020233030321311-1023323232310302-0233111210022300-2323112130030212"></a>

## Next pages — route_ref / 022300230221 / 7

- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.custom_route_object](resources--workload--reference--group-027.md#canonical-3220230303101313-0122201023323032-0302022132031233-3332233000011332-2123202133102020-1133230203021321-3311110223100121-0113002202213112)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-1121000311012033-0022112012130003-3211200311100011-1032311002323322-2310120120102220-2103030310121321-2022131002003111-2221012223020220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0013202311100212-2100202111230002-3223133310211102-1013113212100032-3233203231132121-0022211322313120-3002233300132322-3311201333101020"></a>

## stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.direct_response_route — direct_response_route / 030330221211 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-017.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-021.md#canonical-3300101333202310-3223231311312310-2030111033033203-1011103231310023-1313110120311121-1221031312313020-1210302113302033-2110231331000202)
- [stateful_service.advertise_options.advertise_on_public.port](resources--workload--reference--group-025.md#canonical-0120012201110120-2320030110233132-3233102311321321-2202232022223030-2133131131013111-3231210131302330-1022020011332122-3111230313032011)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-025.md#canonical-2203002110120031-3203013111221013-1321101003110222-2103121303322022-0021002301311230-0203100203111101-0321302220100020-1320220233023212)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes](resources--workload--reference--group-027.md#canonical-2300220120202032-2120133330213102-2332100100003133-0101032330232302-2302101320100002-1302333333023222-3232223323232230-0320110220223201)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes](resources--workload--reference--group-027.md#canonical-2302003233200103-3312120222212200-1000311013000201-0201011231133032-2032003100312302-3001032323113031-1031222100120212-0102213202320111)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.direct_response_route

<a id="canonical-0020233301221332-1103333033222330-2212002021011011-1011122233122031-1120011122311031-3001200100203012-1110113100310101-2220121003221103"></a>

Type: `"object"`. single nested block, Optional.

Direct response route matches on path, incoming header, incoming port and/or HTTP method and
responds directly to the matching traffic.

Upstream description:

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

<a id="canonical-3203002223122210-2020202320010300-3021221002123031-1113303230103312-1021022233332030-1003113202320113-0030201133322122-3031332332321330"></a>

## Direct properties — direct_response_route / 030330221211 / 3

- [headers](resources--workload--reference--group-027.md#canonical-2021001211121212-2211010221133110-2312002031033301-2102210013111113-3222011313032323-3023121220121001-3011033011333100-3310033110322011): complete subsection reference.

<a id="canonical-1002010032022320-1313220332020020-1321222312311230-3310221022030003-0331331101122001-1230113030002000-0230310231332210-3023200011133220"></a>

<a id="canonical-1001300020110310-0203133013020010-1332100032322211-0202101303112003-1133021113231302-3113311003302302-3210030231333130-3123233001333012"></a>

## http_method property — direct_response_route / 030330221211 / 4

Type: `"string"`. Optional.

\[Enum: ANY|GET|HEAD|POST|PUT|DELETE|CONNECT|OPTIONS|TRACE|PATCH|COPY\] Specifies the HTTP method
used to access a resource. Any HTTP Method. Possible values are \`ANY\`, \`GET\`, \`HEAD\`,
\`POST\`, \`PUT\`, \`DELETE\`, \`CONNECT\`, \`OPTIONS\`, \`TRACE\`, \`PATCH\`, \`COPY\`. Defaults to
\`ANY\`.

Upstream description:

Specifies the HTTP method used to access a resource.

Any HTTP Method.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("ANY",
    "GET",
    "HEAD",
    "POST",
    "PUT",
    "DELETE",
    "CONNECT",
    "OPTIONS",
    "TRACE",
    "PATCH",
    "COPY"),
}
```

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

- [incoming_port](resources--workload--reference--group-027.md#canonical-1322002301102013-1102013021110231-3312310232130131-3001132310203023-3321020001333012-2313103303323110-3000010233220120-0020103002001122): complete subsection reference.

- [path](resources--workload--reference--group-027.md#canonical-0302210230132021-2032120122300030-3003211010301112-0132311210031203-3133212001221213-2332320330003331-0000231001010103-3221003223021201): complete subsection reference.

- [route_direct_response](resources--workload--reference--group-027.md#canonical-0003100321020112-0121213303301113-0033102000000210-3100212332120000-0223312333312021-0230032032321130-0220300132110301-2121020201013201): complete subsection reference.

<a id="canonical-3223033020012001-0212320000313033-1220111220012103-2302111131000120-1211133021233331-3112310033121012-0221220333312031-2010010023033000"></a>

## Next pages — direct_response_route / 030330221211 / 5

- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.direct_response_route.headers](resources--workload--reference--group-027.md#canonical-2021001211121212-2211010221133110-2312002031033301-2102210013111113-3222011313032323-3023121220121001-3011033011333100-3310033110322011)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.direct_response_route.incoming_port](resources--workload--reference--group-027.md#canonical-1322002301102013-1102013021110231-3312310232130131-3001132310203023-3321020001333012-2313103303323110-3000010233220120-0020103002001122)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.direct_response_route.path](resources--workload--reference--group-027.md#canonical-0302210230132021-2032120122300030-3003211010301112-0132311210031203-3133212001221213-2332320330003331-0000231001010103-3221003223021201)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.direct_response_route.route_direct_response](resources--workload--reference--group-027.md#canonical-0003100321020112-0121213303301113-0033102000000210-3100212332120000-0223312333312021-0230032032321130-0220300132110301-2121020201013201)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes](resources--workload--reference--group-027.md#canonical-2302003233200103-3312120222212200-1000311013000201-0201011231133032-2032003100312302-3001032323113031-1031222100120212-0102213202320111)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-2021001211121212-2211010221133110-2312002031033301-2102210013111113-3222011313032323-3023121220121001-3011033011333100-3310033110322011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2310222213031200-2120232022020320-3100132203312331-2223210101113100-1210331212302220-2213313212133333-2232132002223320-0002131120311100"></a>

## stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.direct_response_route.headers — headers / 133031313211 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-017.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-021.md#canonical-3300101333202310-3223231311312310-2030111033033203-1011103231310023-1313110120311121-1221031312313020-1210302113302033-2110231331000202)
- [stateful_service.advertise_options.advertise_on_public.port](resources--workload--reference--group-025.md#canonical-0120012201110120-2320030110233132-3233102311321321-2202232022223030-2133131131013111-3231210131302330-1022020011332122-3111230313032011)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-025.md#canonical-2203002110120031-3203013111221013-1321101003110222-2103121303322022-0021002301311230-0203100203111101-0321302220100020-1320220233023212)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes](resources--workload--reference--group-027.md#canonical-2300220120202032-2120133330213102-2332100100003133-0101032330232302-2302101320100002-1302333333023222-3232223323232230-0320110220223201)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes](resources--workload--reference--group-027.md#canonical-2302003233200103-3312120222212200-1000311013000201-0201011231133032-2032003100312302-3001032323113031-1031222100120212-0102213202320111)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.direct_response_route](resources--workload--reference--group-027.md#canonical-1121000311012033-0022112012130003-3211200311100011-1032311002323322-2310120120102220-2103030310121321-2022131002003111-2221012223020220)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.direct_response_route.headers

<a id="canonical-0301000221020101-3130013120313002-1200133300301222-2302320232003130-1210023132310330-2032301112211321-1002230010102022-3103322020212110"></a>

Type: `"object"`. list nested block, Optional.

Headers. List of (key, value) headers.

Upstream description:

List of (key, value) headers.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("name"),
  validators.ConflictingListObjectAttributes("exact",
    "presence"),
  validators.ConflictingListObjectAttributes("exact",
    "regex"),
  validators.ConflictingListObjectAttributes("presence",
    "regex")}
```

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-2330120110311131-1102031031032213-0022103131202021-0033021312003220-1111111000131112-0333312203203003-2220231221012022-3031122100023023"></a>

## Direct properties — headers / 133031313211 / 3

<a id="canonical-3021333131322131-1001211231330312-1210211201303110-2003001201311232-1202102220000301-0330103202221333-1033232231322220-3133201133003312"></a>

<a id="canonical-3223221320121120-0133213322302322-1000212332002011-3120212000331220-1013202003323131-2303022132013000-1202130021220121-1030222302221222"></a>

## exact property — headers / 133031313211 / 4

Type: `"string"`. Optional.

Exclusive with \[presence regular expression\] Header value to match exactly.

Upstream description:

Exclusive with \[presence regular expression\] Header value to match exactly.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

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
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.not_empty": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.not_empty": "true"
  }
}
```

<a id="canonical-2323001203331312-1112212233122320-2323231220010111-3211113332322200-0221120222332000-0013103100321310-1233313233311033-2213321330113112"></a>

<a id="canonical-1223303212232223-2211122212020201-0012030210113321-0033011021203311-2123132230100001-1111202313211130-2030001333030202-2023022300123221"></a>

## invert_match property — headers / 133031313211 / 5

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

<a id="canonical-2000211001123300-3100022222221310-1033323023232202-1030231233202300-3212320300002222-1331131220130301-2201010230033202-3332132232320021"></a>

<a id="canonical-3313320332222121-3113110312213211-0313131022213132-2023031123032333-0022033302101213-2132013222121023-2321320201210203-0313301221201330"></a>

## name property — headers / 133031313211 / 6

Type: `"string"`. Optional.

Name. Name of the header.

Upstream description:

Name of the header.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

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

<a id="canonical-0132122300102221-0212111201112222-0231122011112323-2222303120320312-1203221122112111-0123200003320310-0033220323023222-2111110212312130"></a>

<a id="canonical-0232131313310302-0123113011003211-3101020103022221-2202221103203312-3012311012032331-1012210002232102-0201112322300123-0323013303211321"></a>

## presence property — headers / 133031313211 / 7

Type: `"bool"`. Optional.

Exclusive with \[exact regular expression\] If true, check for presence of header.

Upstream description:

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

<a id="canonical-0310121300132010-0232133103111132-1213121201203101-1312033231220222-1323132030303110-0100022021220220-3232020021200001-3311333302032110"></a>

<a id="canonical-2330211131022330-0011233300303311-3120021002212302-3230303303320222-0232332021231310-0110013132202133-1200120311002302-3020313131200200"></a>

## regular expression property — headers / 133031313211 / 8

Type: `"string"`. Optional.

Exclusive with \[exact presence\] regular expression match of the header value in re2 format.

Upstream description:

Exclusive with \[exact presence\] regular expression match of the header value in re2 format.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

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
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.regex": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.regex": "true"
  }
}
```

<a id="canonical-1011231013020203-2023122132303302-0113203012330232-1010113133030022-0012031322221300-1110212223112231-0120112230020333-1312120320303320"></a>

## Next pages — headers / 133031313211 / 9

- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.direct_response_route](resources--workload--reference--group-027.md#canonical-1121000311012033-0022112012130003-3211200311100011-1032311002323322-2310120120102220-2103030310121321-2022131002003111-2221012223020220)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-1322002301102013-1102013021110231-3312310232130131-3001132310203023-3321020001333012-2313103303323110-3000010233220120-0020103002001122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0311132210200003-2223111311010300-3131101210031133-0030211202112113-1103012000231220-1313312020033212-1211233001211001-2311021302000232"></a>

## stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.direct_response_route.incoming_port — incoming_port / 122311002122 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-017.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-021.md#canonical-3300101333202310-3223231311312310-2030111033033203-1011103231310023-1313110120311121-1221031312313020-1210302113302033-2110231331000202)
- [stateful_service.advertise_options.advertise_on_public.port](resources--workload--reference--group-025.md#canonical-0120012201110120-2320030110233132-3233102311321321-2202232022223030-2133131131013111-3231210131302330-1022020011332122-3111230313032011)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-025.md#canonical-2203002110120031-3203013111221013-1321101003110222-2103121303322022-0021002301311230-0203100203111101-0321302220100020-1320220233023212)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes](resources--workload--reference--group-027.md#canonical-2300220120202032-2120133330213102-2332100100003133-0101032330232302-2302101320100002-1302333333023222-3232223323232230-0320110220223201)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes](resources--workload--reference--group-027.md#canonical-2302003233200103-3312120222212200-1000311013000201-0201011231133032-2032003100312302-3001032323113031-1031222100120212-0102213202320111)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.direct_response_route](resources--workload--reference--group-027.md#canonical-1121000311012033-0022112012130003-3211200311100011-1032311002323322-2310120120102220-2103030310121321-2022131002003111-2221012223020220)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.direct_response_route.incoming_port

<a id="canonical-3020331131013311-3000320002010121-3020322300011020-3021212130221223-2021203131212312-2003321322211033-0103111301022222-1123210023232202"></a>

Type: `"object"`. single nested block, Optional.

Port match of the request can be a range or a specific port.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("no_port_match",
    "port"),
  validators.ConflictingObjectAttributes("no_port_match",
    "port_ranges"),
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
  "x-ves-oneof-field-port_match": "[\"no_port_match\",\"port\",\"port_ranges\"]"
}
```

Terraform syntax:

```terraform
incoming_port {
  # Configure direct properties listed below.
}
```

<a id="canonical-3313230233001103-0030331031223121-0321202212310033-0233301311000200-2213232011032221-0312312303001013-2010021021213213-0203100320033102"></a>

## Direct properties — incoming_port / 122311002122 / 3

- [no_port_match](resources--workload--reference--group-027.md#canonical-1203101010132132-3031330103202032-0012011003130331-3013230313320033-0202220132002032-2323200020333021-1001022101003003-1123001003302100): complete subsection reference.

<a id="canonical-0201212020220333-2233121003033012-3310201122310033-1132230121023212-1001131203213303-3332032002003010-3211211333011002-3222113000221131"></a>

<a id="canonical-0231120321033330-0312303212102121-2023103010130332-0213320011133303-1300202233202332-3310332110120201-2133103012220102-3320202113003023"></a>

## port property — incoming_port / 122311002122 / 4

Type: `"number"`. Optional.

Exclusive with \[no\_port\_match port\_ranges\] Exact Port to match.

Upstream description:

Exclusive with \[no\_port\_match port\_ranges\] Exact Port to match.

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

<a id="canonical-1031300132322131-3313133330223211-0033320023212201-2212323200313210-3120330223021221-0130231102021320-3031020323220121-0210300000222111"></a>

<a id="canonical-1320122321211103-0113133332223200-1221101122111222-0122121121330211-3113233032321021-0121303121331123-1002202202122021-2202233022110021"></a>

## port_ranges property — incoming_port / 122311002122 / 5

Type: `"string"`. Optional.

Exclusive with \[no\_port\_match port\] Port range to match.

Upstream description:

Exclusive with \[no\_port\_match port\] Port range to match.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 32),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 32,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 32,
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
    "ves.io.schema.rules.string.max_len": "32",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.port_range": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "32",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.port_range": "true"
  }
}
```

<a id="canonical-1122233301123133-3313332303313203-0322001230331113-1213002300331203-2213000300110211-3302200123000301-3331302322230200-0220112202023212"></a>

## Next pages — incoming_port / 122311002122 / 6

- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.direct_response_route.incoming_port.no_port_match](resources--workload--reference--group-027.md#canonical-1203101010132132-3031330103202032-0012011003130331-3013230313320033-0202220132002032-2323200020333021-1001022101003003-1123001003302100)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.direct_response_route](resources--workload--reference--group-027.md#canonical-1121000311012033-0022112012130003-3211200311100011-1032311002323322-2310120120102220-2103030310121321-2022131002003111-2221012223020220)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-1203101010132132-3031330103202032-0012011003130331-3013230313320033-0202220132002032-2323200020333021-1001022101003003-1123001003302100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1222003322221122-1000302100301321-0332220223131021-1012123112233132-0123221301232331-3011033223333230-2011030131322122-1332323120002233"></a>

## stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.direct_response_route.incoming_port.no_port_match — no_port_match / 311132301023 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-017.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-021.md#canonical-3300101333202310-3223231311312310-2030111033033203-1011103231310023-1313110120311121-1221031312313020-1210302113302033-2110231331000202)
- [stateful_service.advertise_options.advertise_on_public.port](resources--workload--reference--group-025.md#canonical-0120012201110120-2320030110233132-3233102311321321-2202232022223030-2133131131013111-3231210131302330-1022020011332122-3111230313032011)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-025.md#canonical-2203002110120031-3203013111221013-1321101003110222-2103121303322022-0021002301311230-0203100203111101-0321302220100020-1320220233023212)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes](resources--workload--reference--group-027.md#canonical-2300220120202032-2120133330213102-2332100100003133-0101032330232302-2302101320100002-1302333333023222-3232223323232230-0320110220223201)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes](resources--workload--reference--group-027.md#canonical-2302003233200103-3312120222212200-1000311013000201-0201011231133032-2032003100312302-3001032323113031-1031222100120212-0102213202320111)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.direct_response_route](resources--workload--reference--group-027.md#canonical-1121000311012033-0022112012130003-3211200311100011-1032311002323322-2310120120102220-2103030310121321-2022131002003111-2221012223020220)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.direct_response_route.incoming_port](resources--workload--reference--group-027.md#canonical-1322002301102013-1102013021110231-3312310232130131-3001132310203023-3321020001333012-2313103303323110-3000010233220120-0020103002001122)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.direct_response_route.incoming_port.no_port_match

<a id="canonical-0131321222211021-0221233311333010-3323130123221323-2130010030000202-3212213112331212-0313131110122101-1020230213330121-2022133010212131"></a>

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
no_port_match = {}
```

<a id="canonical-2330212320123102-1121202220203100-2323101212121033-3111212133102003-3122330332213222-3001110103330113-3131310221111013-2321333000331010"></a>

## Direct properties — no_port_match / 311132301023 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1001003233223310-3012133122321302-2302123332312213-0003132333321312-3101333132133101-1222311110301331-1210233310031220-1232302210112123"></a>

## Next pages — no_port_match / 311132301023 / 4

- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.direct_response_route.incoming_port](resources--workload--reference--group-027.md#canonical-1322002301102013-1102013021110231-3312310232130131-3001132310203023-3321020001333012-2313103303323110-3000010233220120-0020103002001122)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-0302210230132021-2032120122300030-3003211010301112-0132311210031203-3133212001221213-2332320330003331-0000231001010103-3221003223021201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0133301033321220-1221230003321021-1020022303232213-2023310303021022-2211030222220233-1122231321130213-3021203022320221-2111023302320120"></a>

## stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.direct_response_route.path — path / 230220123231 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-017.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-021.md#canonical-3300101333202310-3223231311312310-2030111033033203-1011103231310023-1313110120311121-1221031312313020-1210302113302033-2110231331000202)
- [stateful_service.advertise_options.advertise_on_public.port](resources--workload--reference--group-025.md#canonical-0120012201110120-2320030110233132-3233102311321321-2202232022223030-2133131131013111-3231210131302330-1022020011332122-3111230313032011)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-025.md#canonical-2203002110120031-3203013111221013-1321101003110222-2103121303322022-0021002301311230-0203100203111101-0321302220100020-1320220233023212)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes](resources--workload--reference--group-027.md#canonical-2300220120202032-2120133330213102-2332100100003133-0101032330232302-2302101320100002-1302333333023222-3232223323232230-0320110220223201)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes](resources--workload--reference--group-027.md#canonical-2302003233200103-3312120222212200-1000311013000201-0201011231133032-2032003100312302-3001032323113031-1031222100120212-0102213202320111)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.direct_response_route](resources--workload--reference--group-027.md#canonical-1121000311012033-0022112012130003-3211200311100011-1032311002323322-2310120120102220-2103030310121321-2022131002003111-2221012223020220)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.direct_response_route.path

<a id="canonical-0123100122330201-1001300022033332-1033103010211233-2103031231010121-1333322220002111-2101200103201321-1013300122033322-3233212110312102"></a>

Type: `"object"`. single nested block, Optional.

Path match of the URI can be either be, Prefix match or exact match or regular expression match.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("path",
    "prefix"),
  validators.ConflictingObjectAttributes("path",
    "regex"),
  validators.ConflictingObjectAttributes("prefix",
    "regex")}
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
  "x-ves-oneof-field-path_match": "[\"path\",\"prefix\",\"regex\"]"
}
```

Terraform syntax:

```terraform
path {
  # Configure direct properties listed below.
}
```

<a id="canonical-2311033001023211-0031212001023130-3321011121322333-0203223231021321-3312003032220213-0301232323003013-1020103233031032-0220103331221101"></a>

## Direct properties — path / 230220123231 / 3

<a id="canonical-3301302122202013-0333111113223331-2232100311012002-3203232002003032-3111313023013002-2312231220102123-2230002233120322-1220001000101323"></a>

<a id="canonical-1121322301203310-3123220100001322-1303210231110301-3300013101303133-1123322122111301-0010030102300223-0010023101031010-1102112303131302"></a>

## path property — path / 230220123231 / 4

Type: `"string"`. Optional.

Exclusive with \[prefix regular expression\] Exact path value to match.

Upstream description:

Exclusive with \[prefix regular expression\] Exact path value to match.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^[/a-zA-Z0-9._-]+$"
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-2032131313302103-0011131102023110-2220031102221110-3302103200222233-0013002303211131-0310312220000003-3233110121302110-2123202321210230"></a>

<a id="canonical-3300310312001103-2213210111301303-3223103021220301-3032231233223102-2330111023113331-3231202332112220-0021132303210002-2030101023123013"></a>

## prefix property — path / 230220123231 / 5

Type: `"string"`. Optional.

Exclusive with \[path regular expression\] Path prefix to match (e.g. The value / will match on all paths)

Upstream description:

Exclusive with \[path regular expression\] Path prefix to match (e.g. The value / will match on all paths)

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
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
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-2321333222212103-3110202301300011-3223221113001011-3303310312322102-3212232123032200-0030130023320312-2022213101131223-1232213030331130"></a>

<a id="canonical-1210112021312022-2330221013100200-2031033300100100-0310130030201121-2330012232020330-0321200333000332-2132111213321322-0200030013233033"></a>

## regular expression property — path / 230220123231 / 6

Type: `"string"`. Optional.

Exclusive with \[path prefix\] Regular expression of path match (e.g. The value .\* will match on
all paths).

Upstream description:

Exclusive with \[path prefix\] Regular expression of path match (e.g. The value .\* will match on
all paths)

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

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
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
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
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.regex": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.regex": "true"
  }
}
```

<a id="canonical-3130222301113130-0100332310033023-3333220110322320-2022313221200231-1121222000322230-2000002113311002-1012202020101333-1033312021222332"></a>

## Next pages — path / 230220123231 / 7

- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.direct_response_route](resources--workload--reference--group-027.md#canonical-1121000311012033-0022112012130003-3211200311100011-1032311002323322-2310120120102220-2103030310121321-2022131002003111-2221012223020220)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-0003100321020112-0121213303301113-0033102000000210-3100212332120000-0223312333312021-0230032032321130-0220300132110301-2121020201013201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1200331223321302-3330032200311021-1232313211231031-3021311322222101-0310030200212122-2113121031210110-0212101322103232-1220130320300001"></a>

## stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.direct_response_route.route_direct_response — route_direct_response / 022110301101 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-017.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-021.md#canonical-3300101333202310-3223231311312310-2030111033033203-1011103231310023-1313110120311121-1221031312313020-1210302113302033-2110231331000202)
- [stateful_service.advertise_options.advertise_on_public.port](resources--workload--reference--group-025.md#canonical-0120012201110120-2320030110233132-3233102311321321-2202232022223030-2133131131013111-3231210131302330-1022020011332122-3111230313032011)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-025.md#canonical-2203002110120031-3203013111221013-1321101003110222-2103121303322022-0021002301311230-0203100203111101-0321302220100020-1320220233023212)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes](resources--workload--reference--group-027.md#canonical-2300220120202032-2120133330213102-2332100100003133-0101032330232302-2302101320100002-1302333333023222-3232223323232230-0320110220223201)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes](resources--workload--reference--group-027.md#canonical-2302003233200103-3312120222212200-1000311013000201-0201011231133032-2032003100312302-3001032323113031-1031222100120212-0102213202320111)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.direct_response_route](resources--workload--reference--group-027.md#canonical-1121000311012033-0022112012130003-3211200311100011-1032311002323322-2310120120102220-2103030310121321-2022131002003111-2221012223020220)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.direct_response_route.route_direct_response

<a id="canonical-1233001121320210-3023033002010212-3010202200130011-3303223131311023-0331032113312133-2220303102320103-3031001120312211-0110013320333300"></a>

Type: `"object"`. single nested block, Optional.

Send this direct response in case of route match action is direct response.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("response_code")}
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
route_direct_response {
  # Configure direct properties listed below.
}
```

<a id="canonical-1111233002330111-1133010203011031-0203031121010313-2303030000333223-2220003131311201-1313101323133332-0222003120203100-2103033123302023"></a>

## Direct properties — route_direct_response / 022110301101 / 3

<a id="canonical-2313132133231321-3211332030310012-3021102201023000-1030302311333210-3302030010321113-0012133012323220-3321000031122320-2000131333002223"></a>

<a id="canonical-0333310302111302-3020221102012322-2033021003331333-2233303222220101-2302130332030130-1133012312000330-2033211202000220-1122223021011132"></a>

## response_body_encoded property — route_direct_response / 022110301101 / 4

Type: `"string"`. Optional.

Response body to send. Currently supported URL schemes is string:/// for which message should be
encoded in base64 format. The message can be either plain text or HTML.

Upstream description:

Response body to send. Currently supported URL schemes is string:/// for which message should be
encoded in base64 format. The message can be either plain text or HTML. E.g. "&lt;p&gt; Access
Denied &lt;/p&gt;". base64 encoded string URL for this is
string:///PHA+IEFjY2VzcyBEZW5pZWQgPC9wPg==.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(65536),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 65536,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 65536
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "maxLength": 65536,
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
    "ves.io.schema.rules.string.max_bytes": "65536",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "65536",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

<a id="canonical-1113331310202020-1130021103303120-0230331212311221-3332000203322321-2122003103333103-3030002122130020-3300233010200120-2312222222023302"></a>

<a id="canonical-2131123113023111-2101101131030103-0032032222013030-3101032103313313-3312033311013023-2321320002332032-3010032012211220-3221021210311212"></a>

## response_code property — route_direct_response / 022110301101 / 5

Type: `"number"`. Optional.

Response Code. Response code to send.

Upstream description:

Response code to send.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(100, 599),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 599,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 100
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "100",
    "ves.io.schema.rules.uint32.lte": "599"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "100",
    "ves.io.schema.rules.uint32.lte": "599"
  }
}
```

<a id="canonical-0311123013111203-3012202003221131-3012311203023332-3123011331322120-0232223220303010-0333330222202013-3033013233032132-2100302213020231"></a>

## Next pages — route_direct_response / 022110301101 / 6

- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.direct_response_route](resources--workload--reference--group-027.md#canonical-1121000311012033-0022112012130003-3211200311100011-1032311002323322-2310120120102220-2103030310121321-2022131002003111-2221012223020220)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-0330013222233302-2210122030021330-0131201133002112-2132231131331102-2322210303201110-1121010222301222-3302213332031013-0122333212301102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3333300320233210-2032121031130010-3210303130112200-3212001001010303-0131213020232313-0002130013333320-0110312030231032-0110023103001021"></a>

## stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.redirect_route — redirect_route / 030030011103 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-017.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-021.md#canonical-3300101333202310-3223231311312310-2030111033033203-1011103231310023-1313110120311121-1221031312313020-1210302113302033-2110231331000202)
- [stateful_service.advertise_options.advertise_on_public.port](resources--workload--reference--group-025.md#canonical-0120012201110120-2320030110233132-3233102311321321-2202232022223030-2133131131013111-3231210131302330-1022020011332122-3111230313032011)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-025.md#canonical-2203002110120031-3203013111221013-1321101003110222-2103121303322022-0021002301311230-0203100203111101-0321302220100020-1320220233023212)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes](resources--workload--reference--group-027.md#canonical-2300220120202032-2120133330213102-2332100100003133-0101032330232302-2302101320100002-1302333333023222-3232223323232230-0320110220223201)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes](resources--workload--reference--group-027.md#canonical-2302003233200103-3312120222212200-1000311013000201-0201011231133032-2032003100312302-3001032323113031-1031222100120212-0102213202320111)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.redirect_route

<a id="canonical-0113001320230030-3030303002211321-1203330101020331-2313322033321232-2133303312122002-2002301002223103-0132010320301211-1112230210333302"></a>

Type: `"object"`. single nested block, Optional.

Redirect route matches on path, incoming header, incoming port and/or HTTP method and redirects the
matching traffic to a different URL.

Upstream description:

A redirect route matches on path, incoming header, incoming port and/or HTTP method and redirects
the matching traffic to a different URL.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
redirect_route {
  # Configure direct properties listed below.
}
```

<a id="canonical-0102122220010120-3323032332231221-0201011322121231-3000222211110312-3212020203322030-2000103310323231-1310301213022200-2301303112032032"></a>

## Direct properties — redirect_route / 030030011103 / 3

- [headers](resources--workload--reference--group-028.md#canonical-0032131122323001-1033332220221013-0310210230202112-2321033223002123-2113001013123333-2232332230200021-3130102220022032-3130121230320302): complete subsection reference.

<a id="canonical-3133311121202300-1022113111031120-0012330113203203-2223101000012203-3213320220131231-0002331330001121-1011202203303122-1221102113313321"></a>

<a id="canonical-1223023331000300-2230303002132312-1120323113331103-2313331302120210-1313322121322210-2313203101333120-2032010113220203-0102110012132321"></a>

## http_method property — redirect_route / 030030011103 / 4

Type: `"string"`. Optional.

\[Enum: ANY|GET|HEAD|POST|PUT|DELETE|CONNECT|OPTIONS|TRACE|PATCH|COPY\] Specifies the HTTP method
used to access a resource. Any HTTP Method. Possible values are \`ANY\`, \`GET\`, \`HEAD\`,
\`POST\`, \`PUT\`, \`DELETE\`, \`CONNECT\`, \`OPTIONS\`, \`TRACE\`, \`PATCH\`, \`COPY\`. Defaults to
\`ANY\`.

Upstream description:

Specifies the HTTP method used to access a resource.

Any HTTP Method.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("ANY",
    "GET",
    "HEAD",
    "POST",
    "PUT",
    "DELETE",
    "CONNECT",
    "OPTIONS",
    "TRACE",
    "PATCH",
    "COPY"),
}
```

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

- [incoming_port](resources--workload--reference--group-028.md#canonical-1303023022032323-0011211310021103-2302310103100323-1020331133202010-1211220133311022-0313230133233210-0330011000132000-0200302302220132): complete subsection reference.

- [path](resources--workload--reference--group-028.md#canonical-3012300121100130-2123213313031300-1111323022000031-3200021200320130-3002333023012223-3010202011010023-0133110113330213-0102003333203322): complete subsection reference.

- [route_redirect](resources--workload--reference--group-028.md#canonical-3302110132323220-0130312011310200-3213310121101213-2332012011222213-2322213101231122-2221221001002212-3303021101022230-2212202010103100): complete subsection reference.
