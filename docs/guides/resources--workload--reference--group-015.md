---
page_title: "xcsh_workload reference"
subcategory: "Container"
description: "Complete grouped canonical reference for xcsh_workload reference."
---

# xcsh_workload reference

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

- [routes](resources--workload--reference--group-015.md#canonical-1320302000201200-1013121233121332-1200303031313333-2131200212100310-2031222303000133-0013120300210320-3032201103202020-0103102312330232): complete subsection reference.

<a id="canonical-2211211031332200-2011221313101132-1021310011103300-2111301211103310-1311011101011032-1203203110110310-3010031100311212-0132301030202302"></a>

## Next pages — specific_routes / 123102030120 / 4

- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes](resources--workload--reference--group-015.md#canonical-1320302000201200-1013121233121332-1200303031313333-2131200212100310-2031222303000133-0013120300210320-3032201103202020-0103102312330232)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-012.md#canonical-0101000213001300-0222020002101313-1020132032013321-2231011111202333-2122231103331210-3202123132002020-0313332003123002-3101311011132330)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-1320302000201200-1013121233121332-1200303031313333-2131200212100310-2031222303000133-0013120300210320-3032201103202020-0103102312330232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1333301300012222-3332311202103323-3220111112033013-2323202001312333-3330221000333320-1323211003133210-2030321020031122-2012011013223233"></a>

## service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes — routes / 201130302313 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.port](resources--workload--reference--group-012.md#canonical-3013031222022030-2021222233122213-2220022232301003-3102220000332302-0311213311310132-3133321213021123-0133300222212330-1122131003023301)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-012.md#canonical-0101000213001300-0222020002101313-1020132032013321-2231011111202333-2122231103331210-3202123132002020-0313332003123002-3101311011132330)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes](resources--workload--reference--group-015.md#canonical-1223131333131301-1101202200212003-2303000010310012-3003131221200020-0201223130203322-2303023001222013-1231100011211030-2032330121100311)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes

<a id="canonical-0122003303322123-0131032103031322-2032320013303211-1333013223303210-1312020211223331-1103101123233312-2030112302031232-2101121111021302"></a>

Type: `"object"`. list nested block, Optional.

Routes. Routes for this loadbalancer.

Upstream description:

Routes for this loadbalancer.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

<a id="canonical-1113331332120302-2330310311021312-3211233300203311-1202233022111223-2122312130231222-2001020300223020-2213303310020323-0230311000033301"></a>

## Direct properties — routes / 201130302313 / 3

- [custom_route_object](resources--workload--reference--group-015.md#canonical-0210032211120033-1031103021121212-2103210321113312-3223110300113221-2120130213202222-0311013133210132-3032003113103210-0013130303202222): complete subsection reference.

- [direct_response_route](resources--workload--reference--group-015.md#canonical-3323201111322002-1033013210321030-1202233310120130-0103032213302001-0022102310321331-0300303311223133-3133223031000013-0100120001003301): complete subsection reference.

- [redirect_route](resources--workload--reference--group-015.md#canonical-3222223113112223-3211013120021223-1311101201320313-1332203113201030-0200233021032333-2313310201233223-3321131333112311-2102210333302332): complete subsection reference.

- [simple_route](resources--workload--reference--group-015.md#canonical-3130302330333200-3231020322000031-3022221013222100-1033010302022131-0310323013110100-0001322111112133-1103302220233331-0223232310322020): complete subsection reference.

<a id="canonical-1020003023200113-3021323110120311-2201231131132122-1120323012113123-1310203131131333-0033322220000211-0013133221011113-2213301023113302"></a>

## Next pages — routes / 201130302313 / 4

- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.custom_route_object](resources--workload--reference--group-015.md#canonical-0210032211120033-1031103021121212-2103210321113312-3223110300113221-2120130213202222-0311013133210132-3032003113103210-0013130303202222)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.direct_response_route](resources--workload--reference--group-015.md#canonical-3323201111322002-1033013210321030-1202233310120130-0103032213302001-0022102310321331-0300303311223133-3133223031000013-0100120001003301)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.redirect_route](resources--workload--reference--group-015.md#canonical-3222223113112223-3211013120021223-1311101201320313-1332203113201030-0200233021032333-2313310201233223-3321131333112311-2102210333302332)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.simple_route](resources--workload--reference--group-015.md#canonical-3130302330333200-3231020322000031-3022221013222100-1033010302022131-0310323013110100-0001322111112133-1103302220233331-0223232310322020)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes](resources--workload--reference--group-015.md#canonical-1223131333131301-1101202200212003-2303000010310012-3003131221200020-0201223130203322-2303023001222013-1231100011211030-2032330121100311)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-0210032211120033-1031103021121212-2103210321113312-3223110300113221-2120130213202222-0311013133210132-3032003113103210-0013130303202222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1202210033302313-3121121132013200-3003330303131332-2323010210222022-3120010331032000-1311131012223130-0021000111230320-1030103111331032"></a>

## service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.custom_route_object — custom_route_object / 332232300211 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.port](resources--workload--reference--group-012.md#canonical-3013031222022030-2021222233122213-2220022232301003-3102220000332302-0311213311310132-3133321213021123-0133300222212330-1122131003023301)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-012.md#canonical-0101000213001300-0222020002101313-1020132032013321-2231011111202333-2122231103331210-3202123132002020-0313332003123002-3101311011132330)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes](resources--workload--reference--group-015.md#canonical-1223131333131301-1101202200212003-2303000010310012-3003131221200020-0201223130203322-2303023001222013-1231100011211030-2032330121100311)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes](resources--workload--reference--group-015.md#canonical-1320302000201200-1013121233121332-1200303031313333-2131200212100310-2031222303000133-0013120300210320-3032201103202020-0103102312330232)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.custom_route_object

<a id="canonical-0213233101120032-0100322222331032-1001232220331030-2021022030220013-3312303033130233-1232220132120320-2231110210000330-2300213020120032"></a>

Type: `"object"`. single nested block, Optional.

Custom route uses a route object created outside of this view.

Upstream description:

A custom route uses a route object created outside of this view.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

<a id="canonical-2120231303110232-1223123101331103-3013130303022303-3001131200223022-3320103322301331-0020233101321222-0222133100310013-1102032231322010"></a>

## Direct properties — custom_route_object / 332232300211 / 3

- [caching_disable](resources--workload--reference--group-015.md#canonical-3111222103212100-1113220231310310-0101303230011001-3221133320310301-3202210021010200-1331123001332300-1030113101021211-1112003311030112): complete subsection reference.

- [caching_inherit](resources--workload--reference--group-015.md#canonical-3101332020003300-1110021321130122-3201102111322210-2331230303231302-1021322212032013-3302320312032331-1123231023311003-1003000310200203): complete subsection reference.

- [route_ref](resources--workload--reference--group-015.md#canonical-0320120120232103-2132323323112122-2222100123100112-2232031322203123-3303103022123020-3121012232332000-1131303033022210-1303011201223013): complete subsection reference.

<a id="canonical-1103232203222313-2222221121332320-2000001300202310-1302202012101330-2302331203300310-1000022023133031-0301320003313120-3203233111203110"></a>

## Next pages — custom_route_object / 332232300211 / 4

- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.custom_route_object.caching_disable](resources--workload--reference--group-015.md#canonical-3111222103212100-1113220231310310-0101303230011001-3221133320310301-3202210021010200-1331123001332300-1030113101021211-1112003311030112)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.custom_route_object.caching_inherit](resources--workload--reference--group-015.md#canonical-3101332020003300-1110021321130122-3201102111322210-2331230303231302-1021322212032013-3302320312032331-1123231023311003-1003000310200203)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.custom_route_object.route_ref](resources--workload--reference--group-015.md#canonical-0320120120232103-2132323323112122-2222100123100112-2232031322203123-3303103022123020-3121012232332000-1131303033022210-1303011201223013)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes](resources--workload--reference--group-015.md#canonical-1320302000201200-1013121233121332-1200303031313333-2131200212100310-2031222303000133-0013120300210320-3032201103202020-0103102312330232)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-3111222103212100-1113220231310310-0101303230011001-3221133320310301-3202210021010200-1331123001332300-1030113101021211-1112003311030112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1231023232301010-1111131321122311-0133203223131010-3330212302030012-0003221030022112-3211103233120120-3221113120112203-2211230302333123"></a>

## service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.custom_route_object.caching_disable — caching_disable / 331113133100 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.port](resources--workload--reference--group-012.md#canonical-3013031222022030-2021222233122213-2220022232301003-3102220000332302-0311213311310132-3133321213021123-0133300222212330-1122131003023301)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-012.md#canonical-0101000213001300-0222020002101313-1020132032013321-2231011111202333-2122231103331210-3202123132002020-0313332003123002-3101311011132330)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes](resources--workload--reference--group-015.md#canonical-1223131333131301-1101202200212003-2303000010310012-3003131221200020-0201223130203322-2303023001222013-1231100011211030-2032330121100311)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes](resources--workload--reference--group-015.md#canonical-1320302000201200-1013121233121332-1200303031313333-2131200212100310-2031222303000133-0013120300210320-3032201103202020-0103102312330232)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.custom_route_object](resources--workload--reference--group-015.md#canonical-0210032211120033-1031103021121212-2103210321113312-3223110300113221-2120130213202222-0311013133210132-3032003113103210-0013130303202222)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.custom_route_object.caching_disable

<a id="canonical-2113213112211230-1031311220123122-3332012323131323-3312101203323110-1231020033010311-2002130023032201-1132111132102131-1203102332332201"></a>

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

<a id="canonical-1231113213210211-3323120221203311-1133112001220310-3022310203103331-3321223302330212-2211132201123220-0032333310210001-0332322121311023"></a>

## Direct properties — caching_disable / 331113133100 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2100321023110310-0101332020033201-1212030131300003-1310313121231033-1102311022322323-2232030221021210-3000323030021221-2113000021032211"></a>

## Next pages — caching_disable / 331113133100 / 4

- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.custom_route_object](resources--workload--reference--group-015.md#canonical-0210032211120033-1031103021121212-2103210321113312-3223110300113221-2120130213202222-0311013133210132-3032003113103210-0013130303202222)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-3101332020003300-1110021321130122-3201102111322210-2331230303231302-1021322212032013-3302320312032331-1123231023311003-1003000310200203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0200211121302123-0011020330003323-3211101311002220-0012303313112230-0231001023212131-3111200121311330-0022321101011032-0133222021000323"></a>

## service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.custom_route_object.caching_inherit — caching_inherit / 120220100211 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.port](resources--workload--reference--group-012.md#canonical-3013031222022030-2021222233122213-2220022232301003-3102220000332302-0311213311310132-3133321213021123-0133300222212330-1122131003023301)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-012.md#canonical-0101000213001300-0222020002101313-1020132032013321-2231011111202333-2122231103331210-3202123132002020-0313332003123002-3101311011132330)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes](resources--workload--reference--group-015.md#canonical-1223131333131301-1101202200212003-2303000010310012-3003131221200020-0201223130203322-2303023001222013-1231100011211030-2032330121100311)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes](resources--workload--reference--group-015.md#canonical-1320302000201200-1013121233121332-1200303031313333-2131200212100310-2031222303000133-0013120300210320-3032201103202020-0103102312330232)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.custom_route_object](resources--workload--reference--group-015.md#canonical-0210032211120033-1031103021121212-2103210321113312-3223110300113221-2120130213202222-0311013133210132-3032003113103210-0013130303202222)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.custom_route_object.caching_inherit

<a id="canonical-0212112011230133-2331011032033023-2020000301220221-2321022023231132-0011121033320001-1000323230310030-2103101110201132-0133212133003222"></a>

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

<a id="canonical-2102022303021011-2223202300112231-3031310101101331-1320313030012011-2200020302210113-1231313021211220-3201123101102210-3100312333301131"></a>

## Direct properties — caching_inherit / 120220100211 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2200313221310323-1312101331323023-0113232301131113-3212101323023120-2220112320013001-2332023122032100-3121321133022003-1212212211110120"></a>

## Next pages — caching_inherit / 120220100211 / 4

- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.custom_route_object](resources--workload--reference--group-015.md#canonical-0210032211120033-1031103021121212-2103210321113312-3223110300113221-2120130213202222-0311013133210132-3032003113103210-0013130303202222)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-0320120120232103-2132323323112122-2222100123100112-2232031322203123-3303103022123020-3121012232332000-1131303033022210-1303011201223013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2101132212110231-0100130201103321-1230203211130011-1332013010313033-2123103032122013-1103020313222321-0020303313220202-3101301332101020"></a>

## service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.custom_route_object.route_ref — route_ref / 123331301211 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.port](resources--workload--reference--group-012.md#canonical-3013031222022030-2021222233122213-2220022232301003-3102220000332302-0311213311310132-3133321213021123-0133300222212330-1122131003023301)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-012.md#canonical-0101000213001300-0222020002101313-1020132032013321-2231011111202333-2122231103331210-3202123132002020-0313332003123002-3101311011132330)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes](resources--workload--reference--group-015.md#canonical-1223131333131301-1101202200212003-2303000010310012-3003131221200020-0201223130203322-2303023001222013-1231100011211030-2032330121100311)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes](resources--workload--reference--group-015.md#canonical-1320302000201200-1013121233121332-1200303031313333-2131200212100310-2031222303000133-0013120300210320-3032201103202020-0103102312330232)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.custom_route_object](resources--workload--reference--group-015.md#canonical-0210032211120033-1031103021121212-2103210321113312-3223110300113221-2120130213202222-0311013133210132-3032003113103210-0013130303202222)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.custom_route_object.route_ref

<a id="canonical-1303032312321031-1213202032221123-0002122031132223-0221332131013001-3321202021310221-2032013100230322-2332031313323321-0220102213003113"></a>

Type: `"object"`. single nested block, Optional.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

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
route_ref {
  # Configure direct properties listed below.
}
```

<a id="canonical-2110011132333313-0100313032003130-1013213133230132-3223101331032110-1320103012002312-2012233221220330-0201110132002110-0212000313333031"></a>

## Direct properties — route_ref / 123331301211 / 3

<a id="canonical-2020300131202133-0310233031230020-1203211310332121-2001311123222110-3112010321003323-1022030233322023-2020213230230111-0120330222312100"></a>

<a id="canonical-0303022002030012-0313022202210200-2010112110013003-0212302303002203-1003032011302311-3133003310033120-0333030211122033-2111222322321320"></a>

## name property — route_ref / 123331301211 / 4

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

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

<a id="canonical-3310332123100231-0232030220200033-1220331232111112-2231111331120203-0101322132313322-0311122201122310-3003011020213020-3132120101331130"></a>

<a id="canonical-0111002023233332-0132223221123231-0321022221001121-3022130013310103-0003122022033312-2230031121013102-0132230010131312-0320130211332123"></a>

## namespace property — route_ref / 123331301211 / 5

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

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

<a id="canonical-1232021013001100-2133221120112222-1213230223221313-3010201013013223-0003010013221302-2031233310121123-3231031233212203-1102303133331220"></a>

<a id="canonical-2322333300201122-3201223021112031-3221020331331220-1033032132023200-0313223012120201-1000222300230202-1203022323313033-3332012102033020"></a>

## tenant property — route_ref / 123331301211 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

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

<a id="canonical-1011200002031331-3021012220033132-3122110103030312-2031103322323301-3231313223203222-3103121102322303-2021131321322201-0010123233302110"></a>

## Next pages — route_ref / 123331301211 / 7

- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.custom_route_object](resources--workload--reference--group-015.md#canonical-0210032211120033-1031103021121212-2103210321113312-3223110300113221-2120130213202222-0311013133210132-3032003113103210-0013130303202222)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-3323201111322002-1033013210321030-1202233310120130-0103032213302001-0022102310321331-0300303311223133-3133223031000013-0100120001003301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0211030210013210-0220021303011123-0132131122212311-1213111201010312-1001101312011121-3031103113112330-1032000122210321-0112030013303122"></a>

## service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.direct_response_route — direct_response_route / 032221301031 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.port](resources--workload--reference--group-012.md#canonical-3013031222022030-2021222233122213-2220022232301003-3102220000332302-0311213311310132-3133321213021123-0133300222212330-1122131003023301)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-012.md#canonical-0101000213001300-0222020002101313-1020132032013321-2231011111202333-2122231103331210-3202123132002020-0313332003123002-3101311011132330)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes](resources--workload--reference--group-015.md#canonical-1223131333131301-1101202200212003-2303000010310012-3003131221200020-0201223130203322-2303023001222013-1231100011211030-2032330121100311)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes](resources--workload--reference--group-015.md#canonical-1320302000201200-1013121233121332-1200303031313333-2131200212100310-2031222303000133-0013120300210320-3032201103202020-0103102312330232)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.direct_response_route

<a id="canonical-3101311033312122-0300131220330000-3333123313012332-1001000120013113-1131000023222132-3332001120223321-0323221202301310-1211233233320013"></a>

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

<a id="canonical-2230123102311212-2212303203322332-0323211133313010-1201213211130300-2320321033021231-2231320021032320-2220010330000011-2323322210313021"></a>

## Direct properties — direct_response_route / 032221301031 / 3

- [headers](resources--workload--reference--group-015.md#canonical-0222010010121030-2002121221111330-0210203213300203-0132213001033020-0012203131231310-1223010112313230-3120022201313120-1312230321223032): complete subsection reference.

<a id="canonical-1323321221012323-3130301010301310-1130333030332102-3301023332031123-0302131331023113-2130312310001211-2103331120110103-3022131022222202"></a>

<a id="canonical-3111211201000331-0313332201233101-0321002031212211-3223113131222221-0203233320003202-2203001311002333-3001231032203003-2300022133231320"></a>

## http_method property — direct_response_route / 032221301031 / 4

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
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["ANY","CONNECT","COPY","DELETE","GET","HEAD","OPTIONS","PATCH","POST","PUT","TRACE"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
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

- [incoming_port](resources--workload--reference--group-015.md#canonical-0120013202023333-3320333210330311-2301200302022321-2012232001023312-0310301031010320-1321213323003001-3100201302222113-1201220311213122): complete subsection reference.

- [path](resources--workload--reference--group-015.md#canonical-2011030213320012-1221020120303121-1223221102302032-3020312130010032-3003103231033033-3301032122222001-0003232231020002-1330331122201233): complete subsection reference.

- [route_direct_response](resources--workload--reference--group-015.md#canonical-2103100132103010-2113121013301320-2011101010033333-3311112210230100-3310002131310333-2000321030022310-2031303013103100-1311011120013331): complete subsection reference.

<a id="canonical-0232112220110011-2221303120103220-0331113112021000-1133203200313221-3102001232013312-3020000302312232-3303212013000201-3333311202101100"></a>

## Next pages — direct_response_route / 032221301031 / 5

- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.direct_response_route.headers](resources--workload--reference--group-015.md#canonical-0222010010121030-2002121221111330-0210203213300203-0132213001033020-0012203131231310-1223010112313230-3120022201313120-1312230321223032)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.direct_response_route.incoming_port](resources--workload--reference--group-015.md#canonical-0120013202023333-3320333210330311-2301200302022321-2012232001023312-0310301031010320-1321213323003001-3100201302222113-1201220311213122)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.direct_response_route.path](resources--workload--reference--group-015.md#canonical-2011030213320012-1221020120303121-1223221102302032-3020312130010032-3003103231033033-3301032122222001-0003232231020002-1330331122201233)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.direct_response_route.route_direct_response](resources--workload--reference--group-015.md#canonical-2103100132103010-2113121013301320-2011101010033333-3311112210230100-3310002131310333-2000321030022310-2031303013103100-1311011120013331)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes](resources--workload--reference--group-015.md#canonical-1320302000201200-1013121233121332-1200303031313333-2131200212100310-2031222303000133-0013120300210320-3032201103202020-0103102312330232)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-0222010010121030-2002121221111330-0210203213300203-0132213001033020-0012203131231310-1223010112313230-3120022201313120-1312230321223032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1222230032112231-3202010101001212-3110110022013321-0303330203132021-1120012222122013-2221211232200021-0311303300113213-1021313003231013"></a>

## service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.direct_response_route.headers — headers / 033211213002 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.port](resources--workload--reference--group-012.md#canonical-3013031222022030-2021222233122213-2220022232301003-3102220000332302-0311213311310132-3133321213021123-0133300222212330-1122131003023301)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-012.md#canonical-0101000213001300-0222020002101313-1020132032013321-2231011111202333-2122231103331210-3202123132002020-0313332003123002-3101311011132330)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes](resources--workload--reference--group-015.md#canonical-1223131333131301-1101202200212003-2303000010310012-3003131221200020-0201223130203322-2303023001222013-1231100011211030-2032330121100311)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes](resources--workload--reference--group-015.md#canonical-1320302000201200-1013121233121332-1200303031313333-2131200212100310-2031222303000133-0013120300210320-3032201103202020-0103102312330232)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.direct_response_route](resources--workload--reference--group-015.md#canonical-3323201111322002-1033013210321030-1202233310120130-0103032213302001-0022102310321331-0300303311223133-3133223031000013-0100120001003301)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.direct_response_route.headers

<a id="canonical-1221132100300322-0110112211003300-3102333010332321-2330232102331133-3131211321233130-2003302101213102-3300031330002310-3100003012220203"></a>

Type: `"object"`. list nested block, Optional.

Headers. List of (key, value) headers.

Upstream description:

List of (key, value) headers.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-2013211010031210-0011101210323023-3021022021310332-3030010233202320-0133210303302110-3032223231131101-1311303220330001-2120332311321111"></a>

## Direct properties — headers / 033211213002 / 3

<a id="canonical-1121110301311130-0322031210102022-1120000123123323-3200302210110233-3303031201302311-1332331020331021-3300331010113010-3123230102102230"></a>

<a id="canonical-3233231013322212-3130221123311003-3200022021031210-2013332202031212-3001302111332101-1231302222011223-1220031323122302-0030022223313302"></a>

## exact property — headers / 033211213002 / 4

Type: `"string"`. Optional.

Exclusive with \[presence regular expression\] Header value to match exactly.

Upstream description:

Exclusive with \[presence regular expression\] Header value to match exactly.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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

<a id="canonical-2222130211030310-1020332023112223-3021212031232221-2202233101100001-2131002032211002-3101031130223032-2031030133131331-3313001130313303"></a>

## invert_match property — headers / 033211213002 / 5

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

<a id="canonical-0010130202222012-1313112030131221-2022102311110323-3103300100033133-0023113032000322-1222112123023033-1212003031332113-0012310101221310"></a>

## name property — headers / 033211213002 / 6

Type: `"string"`. Optional.

Name. Name of the header.

Upstream description:

Name of the header.

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

<a id="canonical-2333200322100331-1012023303113331-0033132310232021-2322022001222221-3310332312201030-0130033321031111-3301233103300333-1023002200301310"></a>

## presence property — headers / 033211213002 / 7

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

<a id="canonical-1323022201222333-0120031232122022-2100213333130311-0320133303231011-1132032333303132-0212303210200330-0221030131031110-0110320203032310"></a>

<a id="canonical-3003030133020331-3000131120211321-3011101313123202-3222033100133222-0111031233113123-3121001003002022-2132330001112220-1000102321330302"></a>

## regular expression property — headers / 033211213002 / 8

Type: `"string"`. Optional.

Exclusive with \[exact presence\] regular expression match of the header value in re2 format.

Upstream description:

Exclusive with \[exact presence\] regular expression match of the header value in re2 format.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.regex": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.regex": "true"
  }
}
```

<a id="canonical-1311302203333030-1033130313332203-1221030033200323-1300313310303212-3223132101132023-3132303011103322-1310333032201303-2032031233201011"></a>

## Next pages — headers / 033211213002 / 9

- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.direct_response_route](resources--workload--reference--group-015.md#canonical-3323201111322002-1033013210321030-1202233310120130-0103032213302001-0022102310321331-0300303311223133-3133223031000013-0100120001003301)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-0120013202023333-3320333210330311-2301200302022321-2012232001023312-0310301031010320-1321213323003001-3100201302222113-1201220311213122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0103303320021333-1133210312230220-2103122030113101-0000231102113000-0202100002033222-2330330122200122-1100013230321022-2033222033022011"></a>

## service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.direct_response_route.incoming_port — incoming_port / 012020202113 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.port](resources--workload--reference--group-012.md#canonical-3013031222022030-2021222233122213-2220022232301003-3102220000332302-0311213311310132-3133321213021123-0133300222212330-1122131003023301)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-012.md#canonical-0101000213001300-0222020002101313-1020132032013321-2231011111202333-2122231103331210-3202123132002020-0313332003123002-3101311011132330)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes](resources--workload--reference--group-015.md#canonical-1223131333131301-1101202200212003-2303000010310012-3003131221200020-0201223130203322-2303023001222013-1231100011211030-2032330121100311)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes](resources--workload--reference--group-015.md#canonical-1320302000201200-1013121233121332-1200303031313333-2131200212100310-2031222303000133-0013120300210320-3032201103202020-0103102312330232)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.direct_response_route](resources--workload--reference--group-015.md#canonical-3323201111322002-1033013210321030-1202233310120130-0103032213302001-0022102310321331-0300303311223133-3133223031000013-0100120001003301)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.direct_response_route.incoming_port

<a id="canonical-2313331130021030-1210113021321210-1130311231021221-3113123101332230-1302100221202120-0211003222001230-2221000133020032-3032020330103120"></a>

Type: `"object"`. single nested block, Optional.

Port match of the request can be a range or a specific port.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

<a id="canonical-2201010132100230-3020231303022301-0332300312130332-1101002313203033-2011212222303031-0120220020021122-2311012222332322-0021032211010333"></a>

## Direct properties — incoming_port / 012020202113 / 3

- [no_port_match](resources--workload--reference--group-015.md#canonical-1311320002023131-3132312210221310-1321203110030030-3130103133300111-1310031313023300-0212222322301103-3330001312210000-0331223001212130): complete subsection reference.

<a id="canonical-3323331021213303-1221203310100130-1221102331321121-3032232222220013-3200200011013202-1100301323300000-2230231211132323-3023312222303022"></a>

<a id="canonical-2203333122320111-1321211101011300-2120202132213200-2003101103003201-1322221023212322-1333330331133102-0031202223112320-3213003020330023"></a>

## port property — incoming_port / 012020202113 / 4

Type: `"number"`. Optional.

Exclusive with \[no\_port\_match port\_ranges\] Exact Port to match.

Upstream description:

Exclusive with \[no\_port\_match port\_ranges\] Exact Port to match.

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

<a id="canonical-3003103103333011-0201323321221310-3103011203102010-0103013313022020-1303220112320132-3112202131321232-1320102102331032-0330130312312311"></a>

<a id="canonical-1321303331331031-1301311122333322-3113011210321321-0101312331232331-1213322010100010-3203000323200332-0001102320111033-3122020322113212"></a>

## port_ranges property — incoming_port / 012020202113 / 5

Type: `"string"`. Optional.

Exclusive with \[no\_port\_match port\] Port range to match.

Upstream description:

Exclusive with \[no\_port\_match port\] Port range to match.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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

<a id="canonical-1112231322020222-1300201022313331-0003321230102212-1133213011031212-0100333023330022-1221311003301032-3332130113002033-1201001230330231"></a>

## Next pages — incoming_port / 012020202113 / 6

- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.direct_response_route.incoming_port.no_port_match](resources--workload--reference--group-015.md#canonical-1311320002023131-3132312210221310-1321203110030030-3130103133300111-1310031313023300-0212222322301103-3330001312210000-0331223001212130)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.direct_response_route](resources--workload--reference--group-015.md#canonical-3323201111322002-1033013210321030-1202233310120130-0103032213302001-0022102310321331-0300303311223133-3133223031000013-0100120001003301)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-1311320002023131-3132312210221310-1321203110030030-3130103133300111-1310031313023300-0212222322301103-3330001312210000-0331223001212130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2101122021220201-2001000300120302-3311312102010111-0303212231220123-2332201013000310-0002031111021312-1201011112112233-3312222310131223"></a>

## service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.direct_response_route.incoming_port.no_port_match — no_port_match / 023111131013 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.port](resources--workload--reference--group-012.md#canonical-3013031222022030-2021222233122213-2220022232301003-3102220000332302-0311213311310132-3133321213021123-0133300222212330-1122131003023301)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-012.md#canonical-0101000213001300-0222020002101313-1020132032013321-2231011111202333-2122231103331210-3202123132002020-0313332003123002-3101311011132330)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes](resources--workload--reference--group-015.md#canonical-1223131333131301-1101202200212003-2303000010310012-3003131221200020-0201223130203322-2303023001222013-1231100011211030-2032330121100311)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes](resources--workload--reference--group-015.md#canonical-1320302000201200-1013121233121332-1200303031313333-2131200212100310-2031222303000133-0013120300210320-3032201103202020-0103102312330232)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.direct_response_route](resources--workload--reference--group-015.md#canonical-3323201111322002-1033013210321030-1202233310120130-0103032213302001-0022102310321331-0300303311223133-3133223031000013-0100120001003301)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.direct_response_route.incoming_port](resources--workload--reference--group-015.md#canonical-0120013202023333-3320333210330311-2301200302022321-2012232001023312-0310301031010320-1321213323003001-3100201302222113-1201220311213122)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.direct_response_route.incoming_port.no_port_match

<a id="canonical-0332033020032133-3301312320331000-0101310013131031-1313333233012020-2311011021212330-0030111303030132-3211203131323031-0113123010132310"></a>

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

<a id="canonical-3013032033000223-2210332302001331-0011313012003233-0012010113020322-1123202023310111-1220010333231032-3201300322111313-0303130000310302"></a>

## Direct properties — no_port_match / 023111131013 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3330310300313121-2023201010300011-1020010201000110-0333033032002322-3131200002132202-1003330323100331-2113112312213221-1301333133213101"></a>

## Next pages — no_port_match / 023111131013 / 4

- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.direct_response_route.incoming_port](resources--workload--reference--group-015.md#canonical-0120013202023333-3320333210330311-2301200302022321-2012232001023312-0310301031010320-1321213323003001-3100201302222113-1201220311213122)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-2011030213320012-1221020120303121-1223221102302032-3020312130010032-3003103231033033-3301032122222001-0003232231020002-1330331122201233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2311232203302133-0233211311213202-2011100222131111-1102233322120323-2331111200221100-1121301003213333-3202332313313103-3210002023300332"></a>

## service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.direct_response_route.path — path / 311132233121 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.port](resources--workload--reference--group-012.md#canonical-3013031222022030-2021222233122213-2220022232301003-3102220000332302-0311213311310132-3133321213021123-0133300222212330-1122131003023301)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-012.md#canonical-0101000213001300-0222020002101313-1020132032013321-2231011111202333-2122231103331210-3202123132002020-0313332003123002-3101311011132330)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes](resources--workload--reference--group-015.md#canonical-1223131333131301-1101202200212003-2303000010310012-3003131221200020-0201223130203322-2303023001222013-1231100011211030-2032330121100311)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes](resources--workload--reference--group-015.md#canonical-1320302000201200-1013121233121332-1200303031313333-2131200212100310-2031222303000133-0013120300210320-3032201103202020-0103102312330232)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.direct_response_route](resources--workload--reference--group-015.md#canonical-3323201111322002-1033013210321030-1202233310120130-0103032213302001-0022102310321331-0300303311223133-3133223031000013-0100120001003301)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.direct_response_route.path

<a id="canonical-2320311000233310-0333003112221033-2332313020311102-1012223220031001-0302012132322111-0112031032212121-3132212332321120-2121220313010130"></a>

Type: `"object"`. single nested block, Optional.

Path match of the URI can be either be, Prefix match or exact match or regular expression match.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

<a id="canonical-0313112002331020-2133023231002302-1000003220103023-3020121102303330-2200210122021001-3131002301123333-1130302020321100-2320230301223201"></a>

## Direct properties — path / 311132233121 / 3

<a id="canonical-1322123331102032-1010333033230103-3213222312111101-3113033232120100-1131323101122300-2212202220112311-3121013303302131-0111311200222130"></a>

<a id="canonical-0113032030112031-2300322103021121-0111201122131312-0301003232331133-1200232001222100-1323230331301130-1331011023023202-0223210231330022"></a>

## path property — path / 311132233121 / 4

Type: `"string"`. Optional.

Exclusive with \[prefix regular expression\] Exact path value to match.

Upstream description:

Exclusive with \[prefix regular expression\] Exact path value to match.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-3232003123221132-3113222021222022-1233010322101022-3301322103331322-3131223203223313-3021130230323103-2231030313212233-2220031301202101"></a>

<a id="canonical-0103112221313313-1000221000331130-3300121322133210-1210311001013110-0201222013013203-0212110300131220-1203112310133121-3011110103130203"></a>

## prefix property — path / 311132233121 / 5

Type: `"string"`. Optional.

Exclusive with \[path regular expression\] Path prefix to match (e.g. The value / will match on all paths)

Upstream description:

Exclusive with \[path regular expression\] Path prefix to match (e.g. The value / will match on all paths)

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-3011031131212201-0323112022131011-3311022101220302-3032010131200223-3202311031210033-1033000110103112-3320232021110302-2022101130313211"></a>

<a id="canonical-2220110310021300-0031301211113130-3001110013301130-2320332133323300-3231211221002022-1202212230033221-0111200332311231-1023222011333120"></a>

## regular expression property — path / 311132233121 / 6

Type: `"string"`. Optional.

Exclusive with \[path prefix\] Regular expression of path match (e.g. The value .\* will match on
all paths).

Upstream description:

Exclusive with \[path prefix\] Regular expression of path match (e.g. The value .\* will match on
all paths)

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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

<a id="canonical-2231000222020333-3020011113022133-0212010013211203-1112222303022012-0021120211200110-1132310121012003-0101000321111220-0030222322303011"></a>

## Next pages — path / 311132233121 / 7

- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.direct_response_route](resources--workload--reference--group-015.md#canonical-3323201111322002-1033013210321030-1202233310120130-0103032213302001-0022102310321331-0300303311223133-3133223031000013-0100120001003301)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-2103100132103010-2113121013301320-2011101010033333-3311112210230100-3310002131310333-2000321030022310-2031303013103100-1311011120013331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0231320311122100-3012122102030202-2022100031010000-1032330102020320-0121121112202033-2203033320000233-0232303001023012-3313313222032231"></a>

## service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.direct_response_route.route_direct_response — route_direct_response / 201321331102 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.port](resources--workload--reference--group-012.md#canonical-3013031222022030-2021222233122213-2220022232301003-3102220000332302-0311213311310132-3133321213021123-0133300222212330-1122131003023301)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-012.md#canonical-0101000213001300-0222020002101313-1020132032013321-2231011111202333-2122231103331210-3202123132002020-0313332003123002-3101311011132330)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes](resources--workload--reference--group-015.md#canonical-1223131333131301-1101202200212003-2303000010310012-3003131221200020-0201223130203322-2303023001222013-1231100011211030-2032330121100311)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes](resources--workload--reference--group-015.md#canonical-1320302000201200-1013121233121332-1200303031313333-2131200212100310-2031222303000133-0013120300210320-3032201103202020-0103102312330232)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.direct_response_route](resources--workload--reference--group-015.md#canonical-3323201111322002-1033013210321030-1202233310120130-0103032213302001-0022102310321331-0300303311223133-3133223031000013-0100120001003301)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.direct_response_route.route_direct_response

<a id="canonical-0122221222222213-0300021333220311-0323321223100322-0332210321321323-3230223102131003-3213323323113203-0002021022301101-1212030330303313"></a>

Type: `"object"`. single nested block, Optional.

Send this direct response in case of route match action is direct response.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

<a id="canonical-0322223131013202-1212011300101333-0111013333021003-1133113233033102-1003101322033212-1201021123332013-0310111000010331-2120000310121130"></a>

## Direct properties — route_direct_response / 201321331102 / 3

<a id="canonical-3003112230222132-0023320301211221-3012100320313223-3232320011203233-3311312102100300-3030103100030310-1220310200131302-0122323112032332"></a>

<a id="canonical-1131033130012233-0012130120021311-0203031101213223-1130320331132101-0233030101230002-1022333233211111-3231323311013032-2212012111000332"></a>

## response_body_encoded property — route_direct_response / 201321331102 / 4

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
EnumExtractionComplete: true
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
    "ves.io.schema.rules.string.max_bytes": "65536",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "65536",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

<a id="canonical-1212021311111200-1000200002000211-2312131232113013-2132232130301113-3021223320133202-0300320210220023-2323230022030000-2131012313210030"></a>

<a id="canonical-0321012321111210-1310201032230322-3202210132131102-2223300313012130-2201330022222200-1220203303103100-2231023102021021-0331101131302020"></a>

## response_code property — route_direct_response / 201321331102 / 5

Type: `"number"`. Optional.

Response Code. Response code to send.

Upstream description:

Response code to send.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-0102002303302003-3002022321132301-2320302101030000-3030011010030020-0122203331023020-2323221211321100-3010101023230312-1333323102321212"></a>

## Next pages — route_direct_response / 201321331102 / 6

- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.direct_response_route](resources--workload--reference--group-015.md#canonical-3323201111322002-1033013210321030-1202233310120130-0103032213302001-0022102310321331-0300303311223133-3133223031000013-0100120001003301)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-3222223113112223-3211013120021223-1311101201320313-1332203113201030-0200233021032333-2313310201233223-3321131333112311-2102210333302332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0313030100332011-0000300031121212-2021222130013332-2122232201233312-1122320301003112-3121322233110100-0211101202202000-0211223112120202"></a>

## service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.redirect_route — redirect_route / 202113203220 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.port](resources--workload--reference--group-012.md#canonical-3013031222022030-2021222233122213-2220022232301003-3102220000332302-0311213311310132-3133321213021123-0133300222212330-1122131003023301)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-012.md#canonical-0101000213001300-0222020002101313-1020132032013321-2231011111202333-2122231103331210-3202123132002020-0313332003123002-3101311011132330)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes](resources--workload--reference--group-015.md#canonical-1223131333131301-1101202200212003-2303000010310012-3003131221200020-0201223130203322-2303023001222013-1231100011211030-2032330121100311)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes](resources--workload--reference--group-015.md#canonical-1320302000201200-1013121233121332-1200303031313333-2131200212100310-2031222303000133-0013120300210320-3032201103202020-0103102312330232)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.redirect_route

<a id="canonical-1203223022000110-0201101120320313-3210030021130112-2011211113112110-1032321012103213-3010010320232202-0332121002212322-2021101220220023"></a>

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

<a id="canonical-1332203011012221-3320013320302000-2332200012100131-0202300313231122-1123211200323030-3312331230203301-3201103033323122-1112320030131231"></a>

## Direct properties — redirect_route / 202113203220 / 3

- [headers](resources--workload--reference--group-015.md#canonical-0120310330212123-0131002110032000-2113000212111032-2123311103322122-3232323202130020-2023330331223030-2222031113120302-0111202302330203): complete subsection reference.

<a id="canonical-2231120302122033-3312020121121002-3302033210320133-2131212303033233-2131333331103112-3123020132011212-2002233022310120-1001200020123122"></a>

<a id="canonical-1331103013000211-3021321001323101-3113230232202223-0323111233203120-2312023120201222-2132332232032202-3302022201231022-2132322213010230"></a>

## http_method property — redirect_route / 202113203220 / 4

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
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["ANY","CONNECT","COPY","DELETE","GET","HEAD","OPTIONS","PATCH","POST","PUT","TRACE"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
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

- [incoming_port](resources--workload--reference--group-015.md#canonical-0202203213231020-2122031002322313-3212310123013132-1322003110030101-2112110002000211-1000200200323301-3200321110021210-0310101322222211): complete subsection reference.

- [path](resources--workload--reference--group-015.md#canonical-3320321133100032-2200123002110231-2021210213020230-2332232130131120-2311100230101112-0032020132101313-1231323003010230-0032332123032011): complete subsection reference.

- [route_redirect](resources--workload--reference--group-015.md#canonical-0122121022203023-2310120022312111-2130023032322231-3303312133103203-1102131032131111-0002311331200312-0133211102121210-3212220021101200): complete subsection reference.

<a id="canonical-1203010120110230-0212310201223011-2311123102233210-1132131103323230-1103121011021301-0203032020011112-1200332213321213-2013333130302031"></a>

## Next pages — redirect_route / 202113203220 / 5

- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.redirect_route.headers](resources--workload--reference--group-015.md#canonical-0120310330212123-0131002110032000-2113000212111032-2123311103322122-3232323202130020-2023330331223030-2222031113120302-0111202302330203)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.redirect_route.incoming_port](resources--workload--reference--group-015.md#canonical-0202203213231020-2122031002322313-3212310123013132-1322003110030101-2112110002000211-1000200200323301-3200321110021210-0310101322222211)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.redirect_route.path](resources--workload--reference--group-015.md#canonical-3320321133100032-2200123002110231-2021210213020230-2332232130131120-2311100230101112-0032020132101313-1231323003010230-0032332123032011)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect](resources--workload--reference--group-015.md#canonical-0122121022203023-2310120022312111-2130023032322231-3303312133103203-1102131032131111-0002311331200312-0133211102121210-3212220021101200)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes](resources--workload--reference--group-015.md#canonical-1320302000201200-1013121233121332-1200303031313333-2131200212100310-2031222303000133-0013120300210320-3032201103202020-0103102312330232)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-0120310330212123-0131002110032000-2113000212111032-2123311103322122-3232323202130020-2023330331223030-2222031113120302-0111202302330203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0322012011220233-3330202312303100-0223221112330121-0222010333021112-3200212111212201-1103310300322310-1103033200132331-0110120031001212"></a>

## service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.redirect_route.headers — headers / 221101023123 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.port](resources--workload--reference--group-012.md#canonical-3013031222022030-2021222233122213-2220022232301003-3102220000332302-0311213311310132-3133321213021123-0133300222212330-1122131003023301)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-012.md#canonical-0101000213001300-0222020002101313-1020132032013321-2231011111202333-2122231103331210-3202123132002020-0313332003123002-3101311011132330)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes](resources--workload--reference--group-015.md#canonical-1223131333131301-1101202200212003-2303000010310012-3003131221200020-0201223130203322-2303023001222013-1231100011211030-2032330121100311)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes](resources--workload--reference--group-015.md#canonical-1320302000201200-1013121233121332-1200303031313333-2131200212100310-2031222303000133-0013120300210320-3032201103202020-0103102312330232)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.redirect_route](resources--workload--reference--group-015.md#canonical-3222223113112223-3211013120021223-1311101201320313-1332203113201030-0200233021032333-2313310201233223-3321131333112311-2102210333302332)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.redirect_route.headers

<a id="canonical-3333013013033210-0330220010232121-3230200233130202-3023031313100020-1100222321231320-1101322333023230-2032230330002032-2330001013221013"></a>

Type: `"object"`. list nested block, Optional.

Headers. List of (key, value) headers.

Upstream description:

List of (key, value) headers.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-2003302332300021-0320112210133210-0231303100003003-0302013112302112-2130210102203022-1203100103000332-1201012022011330-2002022332001021"></a>

## Direct properties — headers / 221101023123 / 3

<a id="canonical-2111233313202112-0122200320133210-3100102001133021-1031301030133032-0013211103331231-1203321121211032-1300211110303202-2131311313310122"></a>

<a id="canonical-1201223222011122-1003211021102110-1120232232311013-3233012231002211-2231320323232310-2000100110131312-2031030103111113-3131212113332303"></a>

## exact property — headers / 221101023123 / 4

Type: `"string"`. Optional.

Exclusive with \[presence regular expression\] Header value to match exactly.

Upstream description:

Exclusive with \[presence regular expression\] Header value to match exactly.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.not_empty": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.not_empty": "true"
  }
}
```

<a id="canonical-2002010200201110-3321320233102221-0120133131023132-0031111212002233-0300302123323230-2033321220203201-0021201203301033-2023200011133032"></a>

<a id="canonical-1230221323222320-2110133230130030-2113103130032003-0011220110033310-3113033223330032-3111110121033020-1232033112220101-1120322033030300"></a>

## invert_match property — headers / 221101023123 / 5

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

<a id="canonical-2022302302302333-0133233112012002-1320023020130030-1330023002122221-3000010122232023-2333110233132022-0200130002133021-2021210011323102"></a>

<a id="canonical-2321230211132132-2301310000222301-3011200120131231-3231200012201010-0321313313123010-3333320303330110-3231303302002102-3311111201203310"></a>

## name property — headers / 221101023123 / 6

Type: `"string"`. Optional.

Name. Name of the header.

Upstream description:

Name of the header.

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

<a id="canonical-3200120313210212-2211212320310013-0301023322102122-1131300102322322-3022020211200033-0311100233223330-1132100200113031-2123001121120102"></a>

<a id="canonical-3031123201301213-2112111102012002-3031020020110101-3102301322123212-2011100312131010-0103031203313201-3220302020300002-0122220321303330"></a>

## presence property — headers / 221101023123 / 7

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

<a id="canonical-1303232012311010-0210232303113231-1312030010111010-2132122120233310-0303210322233322-0223031302103022-0123103121310300-3233330203032333"></a>

<a id="canonical-3203222211031010-3330012113031020-0221320010223120-1223210212302220-1100021121321002-3332302232010320-1220331121323013-2301320202121220"></a>

## regular expression property — headers / 221101023123 / 8

Type: `"string"`. Optional.

Exclusive with \[exact presence\] regular expression match of the header value in re2 format.

Upstream description:

Exclusive with \[exact presence\] regular expression match of the header value in re2 format.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.regex": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.regex": "true"
  }
}
```

<a id="canonical-0011103011320212-2303120301013230-3013010310101232-1200120322100330-3200231201003011-3003033201000000-3132211300312022-1221301321021010"></a>

## Next pages — headers / 221101023123 / 9

- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.redirect_route](resources--workload--reference--group-015.md#canonical-3222223113112223-3211013120021223-1311101201320313-1332203113201030-0200233021032333-2313310201233223-3321131333112311-2102210333302332)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-0202203213231020-2122031002322313-3212310123013132-1322003110030101-2112110002000211-1000200200323301-3200321110021210-0310101322222211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3312013002020333-1322212112113113-0002011130100221-0031030111120132-1312001103312120-1231230030030022-0202311310131033-3021333220211113"></a>

## service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.redirect_route.incoming_port — incoming_port / 223000211212 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.port](resources--workload--reference--group-012.md#canonical-3013031222022030-2021222233122213-2220022232301003-3102220000332302-0311213311310132-3133321213021123-0133300222212330-1122131003023301)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-012.md#canonical-0101000213001300-0222020002101313-1020132032013321-2231011111202333-2122231103331210-3202123132002020-0313332003123002-3101311011132330)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes](resources--workload--reference--group-015.md#canonical-1223131333131301-1101202200212003-2303000010310012-3003131221200020-0201223130203322-2303023001222013-1231100011211030-2032330121100311)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes](resources--workload--reference--group-015.md#canonical-1320302000201200-1013121233121332-1200303031313333-2131200212100310-2031222303000133-0013120300210320-3032201103202020-0103102312330232)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.redirect_route](resources--workload--reference--group-015.md#canonical-3222223113112223-3211013120021223-1311101201320313-1332203113201030-0200233021032333-2313310201233223-3321131333112311-2102210333302332)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.redirect_route.incoming_port

<a id="canonical-0223122233232232-0000132231030132-1211020320112033-0330001201032202-2331020121032123-2323003332003212-0101031023023332-2312110032101331"></a>

Type: `"object"`. single nested block, Optional.

Port match of the request can be a range or a specific port.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

<a id="canonical-0312213030322333-3121332310031222-3132320230203123-3012013322321103-3000021010130011-0012203310213133-3113100211122322-0130323112203210"></a>

## Direct properties — incoming_port / 223000211212 / 3

- [no_port_match](resources--workload--reference--group-015.md#canonical-2303131230322211-0012223231133110-3031131130321223-0121322232023000-1233020232220332-1332301302332130-1222200313201300-1221320121010003): complete subsection reference.

<a id="canonical-0213203310012210-3031101112203301-0003122133113121-1322210330200212-1210222202003331-0221220223222010-0122321012102100-3321300023310023"></a>

<a id="canonical-0133213100211110-2312120313001031-0120203013320202-2321011111203221-1013330221121013-0122010021302200-2321330300133310-1320220011331212"></a>

## port property — incoming_port / 223000211212 / 4

Type: `"number"`. Optional.

Exclusive with \[no\_port\_match port\_ranges\] Exact Port to match.

Upstream description:

Exclusive with \[no\_port\_match port\_ranges\] Exact Port to match.

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

<a id="canonical-3122022320221002-0210210311023301-3302321102030110-3112303211302300-1221103321100211-0210322323112032-1002232021011213-2132100203332231"></a>

<a id="canonical-2131312100221330-1133011002301320-3111103033311222-2313312330110331-1133232332213022-2202230213133130-0100300322331000-3011310101203202"></a>

## port_ranges property — incoming_port / 223000211212 / 5

Type: `"string"`. Optional.

Exclusive with \[no\_port\_match port\] Port range to match.

Upstream description:

Exclusive with \[no\_port\_match port\] Port range to match.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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

<a id="canonical-2120200010101332-1032031021221230-0323303301200100-0320002022332031-0100210130111231-1001313310102200-1313222103002223-1103213121303002"></a>

## Next pages — incoming_port / 223000211212 / 6

- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.redirect_route.incoming_port.no_port_match](resources--workload--reference--group-015.md#canonical-2303131230322211-0012223231133110-3031131130321223-0121322232023000-1233020232220332-1332301302332130-1222200313201300-1221320121010003)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.redirect_route](resources--workload--reference--group-015.md#canonical-3222223113112223-3211013120021223-1311101201320313-1332203113201030-0200233021032333-2313310201233223-3321131333112311-2102210333302332)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-2303131230322211-0012223231133110-3031131130321223-0121322232023000-1233020232220332-1332301302332130-1222200313201300-1221320121010003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1123323312311332-1321321331132102-0032000300030002-2202312231312303-3130120002230331-0312001301203323-0202332312113302-1111200012030302"></a>

## service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.redirect_route.incoming_port.no_port_match — no_port_match / 322000121223 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.port](resources--workload--reference--group-012.md#canonical-3013031222022030-2021222233122213-2220022232301003-3102220000332302-0311213311310132-3133321213021123-0133300222212330-1122131003023301)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-012.md#canonical-0101000213001300-0222020002101313-1020132032013321-2231011111202333-2122231103331210-3202123132002020-0313332003123002-3101311011132330)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes](resources--workload--reference--group-015.md#canonical-1223131333131301-1101202200212003-2303000010310012-3003131221200020-0201223130203322-2303023001222013-1231100011211030-2032330121100311)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes](resources--workload--reference--group-015.md#canonical-1320302000201200-1013121233121332-1200303031313333-2131200212100310-2031222303000133-0013120300210320-3032201103202020-0103102312330232)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.redirect_route](resources--workload--reference--group-015.md#canonical-3222223113112223-3211013120021223-1311101201320313-1332203113201030-0200233021032333-2313310201233223-3321131333112311-2102210333302332)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.redirect_route.incoming_port](resources--workload--reference--group-015.md#canonical-0202203213231020-2122031002322313-3212310123013132-1322003110030101-2112110002000211-1000200200323301-3200321110021210-0310101322222211)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.redirect_route.incoming_port.no_port_match

<a id="canonical-0320001103200303-3322222123233212-1020023131113303-1311331103022031-1310202200021103-3010311333002010-1302310221333110-0230223002221103"></a>

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

<a id="canonical-2203100120101321-0113010112102230-1313133212203231-3130020211330320-3101033332322222-3302123133102002-1133111332111310-1133031211010132"></a>

## Direct properties — no_port_match / 322000121223 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0310132331001322-1322101033000203-2010232300131330-2210220103301000-2131310113122002-0202031111002030-1201133103031221-0122220132123321"></a>

## Next pages — no_port_match / 322000121223 / 4

- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.redirect_route.incoming_port](resources--workload--reference--group-015.md#canonical-0202203213231020-2122031002322313-3212310123013132-1322003110030101-2112110002000211-1000200200323301-3200321110021210-0310101322222211)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-3320321133100032-2200123002110231-2021210213020230-2332232130131120-2311100230101112-0032020132101313-1231323003010230-0032332123032011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3120310121103213-0011202301232133-2321021030130121-3303212321303033-1200223020222200-0233033330101333-1331001013012301-1202312320300320"></a>

## service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.redirect_route.path — path / 103232331120 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.port](resources--workload--reference--group-012.md#canonical-3013031222022030-2021222233122213-2220022232301003-3102220000332302-0311213311310132-3133321213021123-0133300222212330-1122131003023301)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-012.md#canonical-0101000213001300-0222020002101313-1020132032013321-2231011111202333-2122231103331210-3202123132002020-0313332003123002-3101311011132330)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes](resources--workload--reference--group-015.md#canonical-1223131333131301-1101202200212003-2303000010310012-3003131221200020-0201223130203322-2303023001222013-1231100011211030-2032330121100311)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes](resources--workload--reference--group-015.md#canonical-1320302000201200-1013121233121332-1200303031313333-2131200212100310-2031222303000133-0013120300210320-3032201103202020-0103102312330232)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.redirect_route](resources--workload--reference--group-015.md#canonical-3222223113112223-3211013120021223-1311101201320313-1332203113201030-0200233021032333-2313310201233223-3321131333112311-2102210333302332)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.redirect_route.path

<a id="canonical-3033022122332111-2310323002230313-2131011201101311-2023222322233313-3230333012020213-0232231302130123-2131030012021213-1303322310210122"></a>

Type: `"object"`. single nested block, Optional.

Path match of the URI can be either be, Prefix match or exact match or regular expression match.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

<a id="canonical-3313303110323203-3003211101313223-0232121122223332-1212201100201023-3121202022200331-1232220122220333-3213000331130030-2233021313311221"></a>

## Direct properties — path / 103232331120 / 3

<a id="canonical-1320003313023302-3133301220100033-3133020332231020-3102221310013121-2132231122223210-2203303211021132-1013211323331210-1322212022202031"></a>

<a id="canonical-3000330203122211-3023133232312020-1113310331122011-3332322330220120-0202133023313313-1021321220010111-3020322310223013-1322331121200331"></a>

## path property — path / 103232331120 / 4

Type: `"string"`. Optional.

Exclusive with \[prefix regular expression\] Exact path value to match.

Upstream description:

Exclusive with \[prefix regular expression\] Exact path value to match.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-3323130030312300-0213203302033300-1311013132122232-2110012323311211-1223113213023010-2112311233313100-1212330332020311-1101201120132010"></a>

<a id="canonical-2311121123031200-3003301033012332-3223210032302120-1303020320130332-3311132022310030-2200102301001012-1323300132030030-0131023111331031"></a>

## prefix property — path / 103232331120 / 5

Type: `"string"`. Optional.

Exclusive with \[path regular expression\] Path prefix to match (e.g. The value / will match on all paths)

Upstream description:

Exclusive with \[path regular expression\] Path prefix to match (e.g. The value / will match on all paths)

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-0303230002020322-2023123232103233-2033122323121001-0022220322012133-0012200022021312-0222310003200211-1311003001101310-2221101301222011"></a>

<a id="canonical-1233200213220020-1232112303010112-0123211311201111-0230313202203203-2320110200121103-3310231101112330-3312031320302130-1102221310033010"></a>

## regular expression property — path / 103232331120 / 6

Type: `"string"`. Optional.

Exclusive with \[path prefix\] Regular expression of path match (e.g. The value .\* will match on
all paths).

Upstream description:

Exclusive with \[path prefix\] Regular expression of path match (e.g. The value .\* will match on
all paths)

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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

<a id="canonical-1101332330233222-2323001222312001-2321211120303202-2110221310031012-1221122201202113-2003001302333023-1201310022100310-1310103200031010"></a>

## Next pages — path / 103232331120 / 7

- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.redirect_route](resources--workload--reference--group-015.md#canonical-3222223113112223-3211013120021223-1311101201320313-1332203113201030-0200233021032333-2313310201233223-3321131333112311-2102210333302332)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-0122121022203023-2310120022312111-2130023032322231-3303312133103203-1102131032131111-0002311331200312-0133211102121210-3212220021101200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3132320212012310-0023022120230131-2200232002231200-3300322223003031-0001321223223310-3110131122323303-0320202313132132-3130310312102303"></a>

## service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect — route_redirect / 323110011000 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.port](resources--workload--reference--group-012.md#canonical-3013031222022030-2021222233122213-2220022232301003-3102220000332302-0311213311310132-3133321213021123-0133300222212330-1122131003023301)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-012.md#canonical-0101000213001300-0222020002101313-1020132032013321-2231011111202333-2122231103331210-3202123132002020-0313332003123002-3101311011132330)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes](resources--workload--reference--group-015.md#canonical-1223131333131301-1101202200212003-2303000010310012-3003131221200020-0201223130203322-2303023001222013-1231100011211030-2032330121100311)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes](resources--workload--reference--group-015.md#canonical-1320302000201200-1013121233121332-1200303031313333-2131200212100310-2031222303000133-0013120300210320-3032201103202020-0103102312330232)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.redirect_route](resources--workload--reference--group-015.md#canonical-3222223113112223-3211013120021223-1311101201320313-1332203113201030-0200233021032333-2313310201233223-3321131333112311-2102210333302332)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect

<a id="canonical-0100123001323211-2021303233303110-0133023200023222-3302132321202011-0010223232210332-0231320201023200-0130032120030123-0222333302313301"></a>

Type: `"object"`. single nested block, Optional.

Route redirect parameters when match action is redirect.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("path_redirect",
    "prefix_rewrite"),
  validators.ConflictingObjectAttributes("remove_all_params",
    "replace_params"),
  validators.ConflictingObjectAttributes("remove_all_params",
    "retain_all_params"),
  validators.ConflictingObjectAttributes("replace_params",
    "retain_all_params")}
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
  "x-ves-oneof-field-query_params": "[\"remove_all_params\",\"replace_params\",\"retain_all_params\"]",
  "x-ves-oneof-field-redirect_path_choice": "[\"path_redirect\",\"prefix_rewrite\"]"
}
```

Terraform syntax:

```terraform
route_redirect {
  # Configure direct properties listed below.
}
```

<a id="canonical-2130121110123212-1131330030300202-1301110230120033-0330220303103002-1311001002331133-0331222012301330-1221020231033101-0310123120010002"></a>

## Direct properties — route_redirect / 323110011000 / 3

<a id="canonical-2120233322312202-3330021321311222-2003132101211010-1131130102323322-2010322303301301-3232101320111203-2023220033130332-3220003333303013"></a>

<a id="canonical-3302302021130223-2112111333300030-2111321332112230-2023213112213202-2333112100310110-3001212133101310-2212312130233002-0033030013311110"></a>

## host_redirect property — route_redirect / 323110011000 / 4

Type: `"string"`. Optional.

Swap host part of incoming URL in redirect URL.

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

<a id="canonical-3131211113313123-0212320012202213-0323301033032213-0301121001321212-1331303003312311-1033133131300322-0223231202313101-0112303210301310"></a>

<a id="canonical-1200223333232123-0120112212033003-3132221023300031-3230330330013033-2130302130100232-1120203213001322-3121320103322113-2111210321012330"></a>

## path_redirect property — route_redirect / 323110011000 / 5

Type: `"string"`. Optional.

Exclusive with \[prefix\_rewrite\] swap path part of incoming URL in redirect URL.

Upstream description:

Exclusive with \[prefix\_rewrite\] swap path part of incoming URL in redirect URL.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-0112210013100300-0121323020131122-0230031110001101-1030033312202123-3331100333232220-1200130120011021-2123131302320101-3113300010303120"></a>

<a id="canonical-3311221110100231-1022002012012101-3220211221013211-1021031132030003-0201201102210013-0013032102222313-3322113012100020-1213031200211201"></a>

## prefix_rewrite property — route_redirect / 323110011000 / 6

Type: `"string"`. Optional.

Exclusive with \[path\_redirect\] In Redirect response, the matched prefix (or path) should be
swapped with this value. This option allows redirect URLs be dynamically created based on the
request.

Upstream description:

Exclusive with \[path\_redirect\] In Redirect response, the matched prefix (or path) should be
swapped with this value. This option allows redirect URLs be dynamically created based on the
request.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-2303113030211131-3102203302013223-2130331011033223-3100130121310210-0103330122302331-3123333132311213-0001031323030312-2123221331301221"></a>

<a id="canonical-3123020130312031-3321231032212012-1322130233213100-0130000012322113-3103100121320132-3021113111220011-1133333132232123-2200332310122120"></a>

## proto_redirect property — route_redirect / 323110011000 / 7

Type: `"string"`. Optional.

\[Enum: incoming-proto|http|https\] Swap protocol part of incoming URL in redirect URL The protocol
can be swapped with either HTTP or HTTPS When incoming-proto option is specified, swapping of
protocol is not done. Possible values are \`incoming-proto\`, \`http\`, \`https\`.

Upstream description:

Swap protocol part of incoming URL in redirect URL The protocol can be swapped with either HTTP or
HTTPS When incoming-proto option is specified, swapping of protocol is not done.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["http","https","incoming-proto"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
Validators: []validator.String{
  stringvalidator.OneOf("incoming-proto",
    "http",
    "https"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "enum": [
    "incoming-proto",
    "http",
    "https"
  ],
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
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.in": "[\\\"incoming-proto\\\",\\\"http\\\",\\\"https\\\"]"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.in": "[\\\"incoming-proto\\\",\\\"http\\\",\\\"https\\\"]"
  }
}
```

- [remove_all_params](resources--workload--reference--group-015.md#canonical-1110020011223023-2112122113030202-0201021233220013-0220023023323022-0011003210103110-2030130230202121-1320221013030322-3133121032300020): complete subsection reference.

<a id="canonical-0022313301121222-2221110312000231-2021310032133330-0002321322003123-1021010023232023-1132321133113113-3012003110033332-2012202213033231"></a>

<a id="canonical-1121031313132031-1020212203003222-2122232102000221-0023323332323023-2203223231110220-3203220032121222-0121031122313301-2223020121102112"></a>

## replace_params property — route_redirect / 323110011000 / 8

Type: `"string"`. Optional.

Exclusive with \[remove\_all\_params retain\_all\_params\].

Upstream description:

Exclusive with \[remove\_all\_params retain\_all\_params\]

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
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
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-1211312030213213-2312230102332101-3212230130320110-2122210003212211-1221003300221023-1302330211201113-1111033032102223-2110002300003331"></a>

<a id="canonical-1222213021001330-0330021020033202-2022030030330012-1233320020202031-0201330123110123-3002133222230300-3230331123221111-2211203223133332"></a>

## response_code property — route_redirect / 323110011000 / 9

Type: `"number"`. Optional.

The HTTP status code to use in the redirect response.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.AtMost(599),
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
    "ves.io.schema.rules.uint32.lte": "599"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "599"
  }
}
```

- [retain_all_params](resources--workload--reference--group-015.md#canonical-3221213320313200-3021122101031111-0331322202112113-3212021200122230-1130122201331103-2221123231010120-0200321220231320-1230120110311221): complete subsection reference.

<a id="canonical-0202030301011131-3131210221120133-2020201301123000-0132230310001322-1111100110001013-1023111300011203-0201303200200002-1130020200113013"></a>

## Next pages — route_redirect / 323110011000 / 10

- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect.remove_all_params](resources--workload--reference--group-015.md#canonical-1110020011223023-2112122113030202-0201021233220013-0220023023323022-0011003210103110-2030130230202121-1320221013030322-3133121032300020)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect.retain_all_params](resources--workload--reference--group-015.md#canonical-3221213320313200-3021122101031111-0331322202112113-3212021200122230-1130122201331103-2221123231010120-0200321220231320-1230120110311221)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.redirect_route](resources--workload--reference--group-015.md#canonical-3222223113112223-3211013120021223-1311101201320313-1332203113201030-0200233021032333-2313310201233223-3321131333112311-2102210333302332)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-1110020011223023-2112122113030202-0201021233220013-0220023023323022-0011003210103110-2030130230202121-1320221013030322-3133121032300020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3312130322201031-2231203333012102-1213030302210220-2132131313310032-0321231301201232-2002031313332131-2133100020223110-2331122301103321"></a>

## service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect.remove_all_params — remove_all_params / 202033220201 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.port](resources--workload--reference--group-012.md#canonical-3013031222022030-2021222233122213-2220022232301003-3102220000332302-0311213311310132-3133321213021123-0133300222212330-1122131003023301)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-012.md#canonical-0101000213001300-0222020002101313-1020132032013321-2231011111202333-2122231103331210-3202123132002020-0313332003123002-3101311011132330)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes](resources--workload--reference--group-015.md#canonical-1223131333131301-1101202200212003-2303000010310012-3003131221200020-0201223130203322-2303023001222013-1231100011211030-2032330121100311)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes](resources--workload--reference--group-015.md#canonical-1320302000201200-1013121233121332-1200303031313333-2131200212100310-2031222303000133-0013120300210320-3032201103202020-0103102312330232)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.redirect_route](resources--workload--reference--group-015.md#canonical-3222223113112223-3211013120021223-1311101201320313-1332203113201030-0200233021032333-2313310201233223-3321131333112311-2102210333302332)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect](resources--workload--reference--group-015.md#canonical-0122121022203023-2310120022312111-2130023032322231-3303312133103203-1102131032131111-0002311331200312-0133211102121210-3212220021101200)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect.remove_all_params

<a id="canonical-1032101101201122-2203110002332031-3203120230030223-3233312110120213-0133023132112221-3001132121003000-0322101002323233-3331133133331113"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for remove all params.

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
remove_all_params = {}
```

<a id="canonical-1011230322131033-2130331022001303-1033003103221032-2211301002111321-3331110231230011-1023330133322111-3101100213213210-2130013301211320"></a>

## Direct properties — remove_all_params / 202033220201 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3022301331001010-3130331020320221-0223200123321331-1220023132031222-3011332211303313-0301030200112322-2110011111133112-0133300332211133"></a>

## Next pages — remove_all_params / 202033220201 / 4

- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect](resources--workload--reference--group-015.md#canonical-0122121022203023-2310120022312111-2130023032322231-3303312133103203-1102131032131111-0002311331200312-0133211102121210-3212220021101200)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-3221213320313200-3021122101031111-0331322202112113-3212021200122230-1130122201331103-2221123231010120-0200321220231320-1230120110311221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2200121330100302-2231123332203212-3303021001321110-2131300033010130-1202313030032311-2132031313333302-3222211122011010-3312133230111232"></a>

## service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect.retain_all_params — retain_all_params / 322212112120 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.port](resources--workload--reference--group-012.md#canonical-3013031222022030-2021222233122213-2220022232301003-3102220000332302-0311213311310132-3133321213021123-0133300222212330-1122131003023301)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-012.md#canonical-0101000213001300-0222020002101313-1020132032013321-2231011111202333-2122231103331210-3202123132002020-0313332003123002-3101311011132330)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes](resources--workload--reference--group-015.md#canonical-1223131333131301-1101202200212003-2303000010310012-3003131221200020-0201223130203322-2303023001222013-1231100011211030-2032330121100311)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes](resources--workload--reference--group-015.md#canonical-1320302000201200-1013121233121332-1200303031313333-2131200212100310-2031222303000133-0013120300210320-3032201103202020-0103102312330232)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.redirect_route](resources--workload--reference--group-015.md#canonical-3222223113112223-3211013120021223-1311101201320313-1332203113201030-0200233021032333-2313310201233223-3321131333112311-2102210333302332)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect](resources--workload--reference--group-015.md#canonical-0122121022203023-2310120022312111-2130023032322231-3303312133103203-1102131032131111-0002311331200312-0133211102121210-3212220021101200)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect.retain_all_params

<a id="canonical-1013100031223022-1033132320003102-3333312122220332-0320300201013023-0303112033233213-1303310103020222-1200201203103123-1222300300113302"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for retain all params.

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
retain_all_params = {}
```

<a id="canonical-2330110101233200-2330320012110313-0312301030300132-1003312011003013-1003330023300100-3323330320330003-2121332111012101-3303131003220021"></a>

## Direct properties — retain_all_params / 322212112120 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1323112203001223-0203210111033311-3100002123102332-2013011120331031-3020112100011301-1112302121313003-2330020120322213-0033210130332101"></a>

## Next pages — retain_all_params / 322212112120 / 4

- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect](resources--workload--reference--group-015.md#canonical-0122121022203023-2310120022312111-2130023032322231-3303312133103203-1102131032131111-0002311331200312-0133211102121210-3212220021101200)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-3130302330333200-3231020322000031-3022221013222100-1033010302022131-0310323013110100-0001322111112133-1103302220233331-0223232310322020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3020030111222013-1100102113120011-2030113000001013-3113122111110221-0030200321330330-2110103312213011-2200302311300012-3223331031130320"></a>

## service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.simple_route — simple_route / 022033311123 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.port](resources--workload--reference--group-012.md#canonical-3013031222022030-2021222233122213-2220022232301003-3102220000332302-0311213311310132-3133321213021123-0133300222212330-1122131003023301)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-012.md#canonical-0101000213001300-0222020002101313-1020132032013321-2231011111202333-2122231103331210-3202123132002020-0313332003123002-3101311011132330)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes](resources--workload--reference--group-015.md#canonical-1223131333131301-1101202200212003-2303000010310012-3003131221200020-0201223130203322-2303023001222013-1231100011211030-2032330121100311)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes](resources--workload--reference--group-015.md#canonical-1320302000201200-1013121233121332-1200303031313333-2131200212100310-2031222303000133-0013120300210320-3032201103202020-0103102312330232)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.simple_route

<a id="canonical-2022000202121222-0023113002300013-0332232330021033-3230213300123121-1310212311002300-1113011333233032-0130221330112023-2312011202222331"></a>

Type: `"object"`. single nested block, Optional.

Simple route matches on path and/or HTTP method and forwards the matching traffic to the default
origin pool specified outside.

Upstream description:

A simple route matches on path and/or HTTP method and forwards the matching traffic to the default
origin pool specified outside.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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
simple_route {
  # Configure direct properties listed below.
}
```

<a id="canonical-2213310101123231-2013003030333113-1300031132312210-0221313032233232-1130210101310302-2310100210203210-3302130013322312-1233030002122120"></a>

## Direct properties — simple_route / 022033311123 / 3

- [auto_host_rewrite](resources--workload--reference--group-015.md#canonical-0332130002303332-1202033113212210-0103011102322313-1100000230231320-2131203100310213-1331302001220130-3230331103223221-1130202323103102): complete subsection reference.

- [disable_host_rewrite](resources--workload--reference--group-015.md#canonical-2012100332032223-0021323003310010-1332133220203303-0303302031012032-3120113101030130-3010210330110000-3303032231220332-1233113211233302): complete subsection reference.

<a id="canonical-1120123102110323-0220212021230131-2000330132100320-0302003323303232-0133120212131221-3301120313201233-1321000320131210-3013323200330333"></a>

<a id="canonical-3012332130233310-0131103301003223-1223132301221011-3001200311102111-2002320210020012-1203011322200130-1022130031001012-2033200303011233"></a>

## host_rewrite property — simple_route / 022033311123 / 4

Type: `"string"`. Optional.

Exclusive with \[auto\_host\_rewrite disable\_host\_rewrite\] Host header will be swapped with this
value.

Upstream description:

Exclusive with \[auto\_host\_rewrite disable\_host\_rewrite\] Host header will be swapped with this
value.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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
    "ves.io.schema.rules.string.hostname": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true"
  }
}
```

<a id="canonical-0131220110120322-2012221110200000-1233032022303001-1333123332013310-1031122133110030-0120330300012221-2330023102331323-0131101222103132"></a>

<a id="canonical-0300322111332303-1133122130220211-0133002230023020-2101022100021200-0220020222033131-2230210131011031-0213202001112312-3232200002002231"></a>

## http_method property — simple_route / 022033311123 / 5

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
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["ANY","CONNECT","COPY","DELETE","GET","HEAD","OPTIONS","PATCH","POST","PUT","TRACE"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
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

- [path](resources--workload--reference--group-015.md#canonical-3201300030230310-3010121013303332-1123003111111230-3302022111220332-1003323021023000-3102010013122101-3000011120200310-0333032310311233): complete subsection reference.

<a id="canonical-1021333202020010-0311003110303002-2111023111232233-2330223300311200-1001130311221323-0123123313002202-2220033321301212-1230132021211013"></a>

## Next pages — simple_route / 022033311123 / 6

- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.simple_route.auto_host_rewrite](resources--workload--reference--group-015.md#canonical-0332130002303332-1202033113212210-0103011102322313-1100000230231320-2131203100310213-1331302001220130-3230331103223221-1130202323103102)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.simple_route.disable_host_rewrite](resources--workload--reference--group-015.md#canonical-2012100332032223-0021323003310010-1332133220203303-0303302031012032-3120113101030130-3010210330110000-3303032231220332-1233113211233302)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.simple_route.path](resources--workload--reference--group-015.md#canonical-3201300030230310-3010121013303332-1123003111111230-3302022111220332-1003323021023000-3102010013122101-3000011120200310-0333032310311233)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes](resources--workload--reference--group-015.md#canonical-1320302000201200-1013121233121332-1200303031313333-2131200212100310-2031222303000133-0013120300210320-3032201103202020-0103102312330232)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-0332130002303332-1202033113212210-0103011102322313-1100000230231320-2131203100310213-1331302001220130-3230331103223221-1130202323103102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2233203023002102-0133123202130113-3020332203321120-2011212220020312-0020313023321202-2123332320301123-1130123010021333-3300110133211032"></a>

## service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.simple_route.auto_host_rewrite — auto_host_rewrite / 333233311123 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.port](resources--workload--reference--group-012.md#canonical-3013031222022030-2021222233122213-2220022232301003-3102220000332302-0311213311310132-3133321213021123-0133300222212330-1122131003023301)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-012.md#canonical-0101000213001300-0222020002101313-1020132032013321-2231011111202333-2122231103331210-3202123132002020-0313332003123002-3101311011132330)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes](resources--workload--reference--group-015.md#canonical-1223131333131301-1101202200212003-2303000010310012-3003131221200020-0201223130203322-2303023001222013-1231100011211030-2032330121100311)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes](resources--workload--reference--group-015.md#canonical-1320302000201200-1013121233121332-1200303031313333-2131200212100310-2031222303000133-0013120300210320-3032201103202020-0103102312330232)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.simple_route](resources--workload--reference--group-015.md#canonical-3130302330333200-3231020322000031-3022221013222100-1033010302022131-0310323013110100-0001322111112133-1103302220233331-0223232310322020)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.simple_route.auto_host_rewrite

<a id="canonical-3331320332231012-0010012233010323-1021001310120331-3301011032231011-1221100233321331-2220212100130212-0333212012133212-1322033200320122"></a>

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

<a id="canonical-3330103312203220-3110302201131122-3000012120321213-0331102113311232-1111103210321111-3311302333110003-0030323232023301-0003203220010002"></a>

## Direct properties — auto_host_rewrite / 333233311123 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2002330022112200-2210020012200330-1200323310302000-3331221121000132-1113132121032231-2121112323323032-2331012022211120-0200220002123132"></a>

## Next pages — auto_host_rewrite / 333233311123 / 4

- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.simple_route](resources--workload--reference--group-015.md#canonical-3130302330333200-3231020322000031-3022221013222100-1033010302022131-0310323013110100-0001322111112133-1103302220233331-0223232310322020)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-2012100332032223-0021323003310010-1332133220203303-0303302031012032-3120113101030130-3010210330110000-3303032231220332-1233113211233302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1120203003223211-1222213133331323-3201121333221023-2220132112013223-0303011013011311-1003120000230003-0212033001211303-2102132110211233"></a>

## service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.simple_route.disable_host_rewrite — disable_host_rewrite / 001310232120 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.port](resources--workload--reference--group-012.md#canonical-3013031222022030-2021222233122213-2220022232301003-3102220000332302-0311213311310132-3133321213021123-0133300222212330-1122131003023301)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-012.md#canonical-0101000213001300-0222020002101313-1020132032013321-2231011111202333-2122231103331210-3202123132002020-0313332003123002-3101311011132330)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes](resources--workload--reference--group-015.md#canonical-1223131333131301-1101202200212003-2303000010310012-3003131221200020-0201223130203322-2303023001222013-1231100011211030-2032330121100311)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes](resources--workload--reference--group-015.md#canonical-1320302000201200-1013121233121332-1200303031313333-2131200212100310-2031222303000133-0013120300210320-3032201103202020-0103102312330232)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.simple_route](resources--workload--reference--group-015.md#canonical-3130302330333200-3231020322000031-3022221013222100-1033010302022131-0310323013110100-0001322111112133-1103302220233331-0223232310322020)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.simple_route.disable_host_rewrite

<a id="canonical-0202321213023302-0121122212230003-1303323322112321-2032313131000131-3103213001323210-3021002333233112-2100330200031321-1110202000333223"></a>

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

<a id="canonical-0200200002302200-0231333233320030-1011010200203310-2311111010332011-2213001102210000-1311132233313231-3122202113133312-0320232321322330"></a>

## Direct properties — disable_host_rewrite / 001310232120 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3310313131021230-2023012210311303-0111303130100200-1333202012101033-1020011200220113-1000211200233301-3322211210202222-2233011133123101"></a>

## Next pages — disable_host_rewrite / 001310232120 / 4

- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.simple_route](resources--workload--reference--group-015.md#canonical-3130302330333200-3231020322000031-3022221013222100-1033010302022131-0310323013110100-0001322111112133-1103302220233331-0223232310322020)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-3201300030230310-3010121013303332-1123003111111230-3302022111220332-1003323021023000-3102010013122101-3000011120200310-0333032310311233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0103202032030133-1211110233301112-2200301303031131-1323010001111002-2002313021020013-3311300001223211-2010331113300001-1113211010221121"></a>

## service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.simple_route.path — path / 030031231120 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.port](resources--workload--reference--group-012.md#canonical-3013031222022030-2021222233122213-2220022232301003-3102220000332302-0311213311310132-3133321213021123-0133300222212330-1122131003023301)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-012.md#canonical-0101000213001300-0222020002101313-1020132032013321-2231011111202333-2122231103331210-3202123132002020-0313332003123002-3101311011132330)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes](resources--workload--reference--group-015.md#canonical-1223131333131301-1101202200212003-2303000010310012-3003131221200020-0201223130203322-2303023001222013-1231100011211030-2032330121100311)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes](resources--workload--reference--group-015.md#canonical-1320302000201200-1013121233121332-1200303031313333-2131200212100310-2031222303000133-0013120300210320-3032201103202020-0103102312330232)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.simple_route](resources--workload--reference--group-015.md#canonical-3130302330333200-3231020322000031-3022221013222100-1033010302022131-0310323013110100-0001322111112133-1103302220233331-0223232310322020)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.simple_route.path

<a id="canonical-1310201301332031-3122211333013013-1133211111200021-0221331220013032-0231001203303232-0030133010333101-3112313030331202-0020031322011230"></a>

Type: `"object"`. single nested block, Optional.

Path match of the URI can be either be, Prefix match or exact match or regular expression match.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

<a id="canonical-2323330011311303-0001002010301023-3113101201113102-0333211312130311-2133300022113023-2212322031221313-3033101001231303-2220102211130012"></a>

## Direct properties — path / 030031231120 / 3

<a id="canonical-2101013330202312-1300001331130130-1031101013321002-2222221130111021-3302333012232010-0321201031301003-0331212332301100-1200013123300320"></a>

<a id="canonical-3323211100220302-0320001122212203-2310332031310332-3033001203301223-3033012030313221-1110020002311331-0231322313203311-2003013012003001"></a>

## path property — path / 030031231120 / 4

Type: `"string"`. Optional.

Exclusive with \[prefix regular expression\] Exact path value to match.

Upstream description:

Exclusive with \[prefix regular expression\] Exact path value to match.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-1301321101132111-2202302330230020-0213203210332331-3211323222032203-2131031002031233-2020211130100000-2222032232133322-2131221210001302"></a>

<a id="canonical-2320130233111033-2220201103230322-2123313321203322-1003032230211202-2030203110100110-1200131330233221-0133021032101102-0121032002001200"></a>

## prefix property — path / 030031231120 / 5

Type: `"string"`. Optional.

Exclusive with \[path regular expression\] Path prefix to match (e.g. The value / will match on all paths)

Upstream description:

Exclusive with \[path regular expression\] Path prefix to match (e.g. The value / will match on all paths)

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-2000120331131133-1133000203313102-2311211010100121-1210002013213011-1133020222102131-0302130302032123-2312310101112131-2123000122210031"></a>

<a id="canonical-0110001221233320-2310132002221131-2331011212021201-1030133323031223-1013333123000111-1010001220010010-0320231210122331-1120110110331023"></a>

## regular expression property — path / 030031231120 / 6

Type: `"string"`. Optional.

Exclusive with \[path prefix\] Regular expression of path match (e.g. The value .\* will match on
all paths).

Upstream description:

Exclusive with \[path prefix\] Regular expression of path match (e.g. The value .\* will match on
all paths)

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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

<a id="canonical-0300122003003333-2320201213030210-1220303203202301-0120013133112232-0021300023033031-2210013020211101-3332032001303130-1323122313021013"></a>

## Next pages — path / 030031231120 / 7

- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.simple_route](resources--workload--reference--group-015.md#canonical-3130302330333200-3231020322000031-3022221013222100-1033010302022131-0310323013110100-0001322111112133-1103302220233331-0223232310322020)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-0101321212331202-1331113113220303-3201003220312100-1011223301302220-1321111200101302-1211203132222003-1303023311003132-2113020220211101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3123322113203011-3302232322312301-2331101111201120-0211020002310130-0233121013211323-1213331313313303-1021022010210113-2123020003010300"></a>

## service.advertise_options.advertise_on_public.port.port — port / 101011302130 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.port](resources--workload--reference--group-012.md#canonical-3013031222022030-2021222233122213-2220022232301003-3102220000332302-0311213311310132-3133321213021123-0133300222212330-1122131003023301)
- service.advertise_options.advertise_on_public.port.port

<a id="canonical-3033012213031133-3322332103030032-1103232201020110-2323102021223021-3313003021233223-3113122220103112-1321120212033333-2312301000322130"></a>

Type: `"object"`. single nested block, Optional.

Port. Single port.

Upstream description:

Single port.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
port {
  # Configure direct properties listed below.
}
```

<a id="canonical-0322200233313322-2133021222323031-0223020031022331-3311233123032002-0103222210132222-2002201210211211-3032203122133001-3311210010023230"></a>

## Direct properties — port / 101011302130 / 3

- [info](resources--workload--reference--group-015.md#canonical-2233321102202333-3221231301223203-0103122201230313-0002133112312132-3302302022032303-0131103100030133-2003311020201130-2300012000001331): complete subsection reference.

<a id="canonical-0303223123103213-3303033232111111-1332311133232331-0111031103333131-0023203010113031-0211313201023003-3313121133323022-3220110030131212"></a>

## Next pages — port / 101011302130 / 4

- [service.advertise_options.advertise_on_public.port.port.info](resources--workload--reference--group-015.md#canonical-2233321102202333-3221231301223203-0103122201230313-0002133112312132-3302302022032303-0131103100030133-2003311020201130-2300012000001331)
- [service.advertise_options.advertise_on_public.port](resources--workload--reference--group-012.md#canonical-3013031222022030-2021222233122213-2220022232301003-3102220000332302-0311213311310132-3133321213021123-0133300222212330-1122131003023301)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-2233321102202333-3221231301223203-0103122201230313-0002133112312132-3302302022032303-0131103100030133-2003311020201130-2300012000001331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1033223231322310-2101112111030030-1013321000122211-1132233030200321-2001303223001230-1230222223021030-0231000200202222-1322301223221012"></a>

## service.advertise_options.advertise_on_public.port.port.info — info / 001111111031 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.port](resources--workload--reference--group-012.md#canonical-3013031222022030-2021222233122213-2220022232301003-3102220000332302-0311213311310132-3133321213021123-0133300222212330-1122131003023301)
- [service.advertise_options.advertise_on_public.port.port](resources--workload--reference--group-015.md#canonical-0101321212331202-1331113113220303-3201003220312100-1011223301302220-1321111200101302-1211203132222003-1303023311003132-2113020220211101)
- service.advertise_options.advertise_on_public.port.port.info

<a id="canonical-1000300230122203-3111102032100021-0220300033301112-1031213321020113-2213322300133332-3212110111312221-3020303233332313-2332103213213300"></a>

Type: `"object"`. single nested block, Optional.

Port Information. Port information.

Upstream description:

Port information.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("port"),
  validators.ConflictingObjectAttributes("same_as_port",
    "target_port")}
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
  "x-ves-oneof-field-target_port_choice": "[\"same_as_port\",\"target_port\"]"
}
```

Terraform syntax:

```terraform
info {
  # Configure direct properties listed below.
}
```

<a id="canonical-1302320312200311-2331310300212003-2032013210130322-0232111230313101-2313013220303012-1023232231213323-0330201012322232-2232023111030131"></a>

## Direct properties — info / 001111111031 / 3

<a id="canonical-2122003003201301-0232021010310200-1233320113011131-0213031001011123-1132001202010003-1002230221223132-3330220113210032-3221010313021332"></a>

<a id="canonical-3322031200110203-3111213132322132-1122202113301020-0011221310322323-0302122013233303-0231021102020203-3332332322333020-1030132321033001"></a>

## port property — info / 001111111031 / 4

Type: `"number"`. Optional.

Port. Port the workload can be reached on.

Upstream description:

Port the workload can be reached on.

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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gt": "0",
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gt": "0",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-0113131200000020-2110222003111212-0223201111113301-0311021223022003-0132233200122113-1130220030001232-1120320212233322-3201103023322202"></a>

<a id="canonical-3201302001020321-2011301211022202-1101330000103310-2123110301323302-1021023112323132-2010202303011021-1221302222311332-1101112201023020"></a>

## protocol property — info / 001111111031 / 5

Type: `"string"`. Optional.

\[Enum: PROTOCOL\_TCP|PROTOCOL\_HTTP|PROTOCOL\_HTTP2|PROTOCOL\_TLS\_WITH\_SNI|PROTOCOL\_UDP\] Type
of protocol - PROTOCOL\_TCP: TCP TCP - PROTOCOL\_HTTP: HTTP HTTP - PROTOCOL\_HTTP2: HTTP2 HTTP2 -
PROTOCOL\_TLS\_WITH\_SNI: TLS with SNI TLS with SNI - PROTOCOL\_UDP: UDP UDP. Possible values are
\`PROTOCOL\_TCP\`, \`PROTOCOL\_HTTP\`, \`PROTOCOL\_HTTP2\`, \`PROTOCOL\_TLS\_WITH\_SNI\`,
\`PROTOCOL\_UDP\`. Defaults to \`PROTOCOL\_TCP\`.

Upstream description:

Type of protocol

&#8203;- PROTOCOL\_TCP: TCP

TCP &#8203;- PROTOCOL\_HTTP: HTTP

HTTP &#8203;- PROTOCOL\_HTTP2: HTTP2

HTTP2 &#8203;- PROTOCOL\_TLS\_WITH\_SNI: TLS with SNI

TLS with SNI &#8203;- PROTOCOL\_UDP: UDP

UDP.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["PROTOCOL_HTTP","PROTOCOL_HTTP2","PROTOCOL_TCP","PROTOCOL_TLS_WITH_SNI","PROTOCOL_UDP"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
Validators: []validator.String{
  stringvalidator.OneOf("PROTOCOL_TCP",
    "PROTOCOL_HTTP",
    "PROTOCOL_HTTP2",
    "PROTOCOL_TLS_WITH_SNI",
    "PROTOCOL_UDP"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "PROTOCOL_TCP",
  "enum": [
    "PROTOCOL_TCP",
    "PROTOCOL_HTTP",
    "PROTOCOL_HTTP2",
    "PROTOCOL_TLS_WITH_SNI",
    "PROTOCOL_UDP"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [same_as_port](resources--workload--reference--group-015.md#canonical-1030310030111023-2101222122002312-2130203313223212-2022320323000331-3133322111220221-1312021110032023-1123233103003313-3100123211310000): complete subsection reference.

<a id="canonical-3132021112233020-2003010122221310-1102013230111011-1001030022131001-0200032202120001-2233130011013221-3330303321002211-2131221111012202"></a>

<a id="canonical-1001002012300330-2100321313313011-0032202211133201-0031132022231223-0103221121022012-2100323202201012-2333110230332132-0010102032010212"></a>

## target_port property — info / 001111111031 / 6

Type: `"number"`. Optional.

Exclusive with \[same\_as\_port\] Port the workload is listening on.

Upstream description:

Exclusive with \[same\_as\_port\] Port the workload is listening on.

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
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gt": "0",
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gt": "0",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-2002032301131220-0202130021103132-3102300001203111-3201201001333001-2033020201301111-3013233232010223-3300303232220113-0111311312323223"></a>

## Next pages — info / 001111111031 / 7

- [service.advertise_options.advertise_on_public.port.port.info.same_as_port](resources--workload--reference--group-015.md#canonical-1030310030111023-2101222122002312-2130203313223212-2022320323000331-3133322111220221-1312021110032023-1123233103003313-3100123211310000)
- [service.advertise_options.advertise_on_public.port.port](resources--workload--reference--group-015.md#canonical-0101321212331202-1331113113220303-3201003220312100-1011223301302220-1321111200101302-1211203132222003-1303023311003132-2113020220211101)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-1030310030111023-2101222122002312-2130203313223212-2022320323000331-3133322111220221-1312021110032023-1123233103003313-3100123211310000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3022121120023213-2002312333221300-0023313311112203-0200210003112331-3313123003013221-2001111213102300-3211332210013021-2133321001030020"></a>

## service.advertise_options.advertise_on_public.port.port.info.same_as_port — same_as_port / 001031123321 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.port](resources--workload--reference--group-012.md#canonical-3013031222022030-2021222233122213-2220022232301003-3102220000332302-0311213311310132-3133321213021123-0133300222212330-1122131003023301)
- [service.advertise_options.advertise_on_public.port.port](resources--workload--reference--group-015.md#canonical-0101321212331202-1331113113220303-3201003220312100-1011223301302220-1321111200101302-1211203132222003-1303023311003132-2113020220211101)
- [service.advertise_options.advertise_on_public.port.port.info](resources--workload--reference--group-015.md#canonical-2233321102202333-3221231301223203-0103122201230313-0002133112312132-3302302022032303-0131103100030133-2003311020201130-2300012000001331)
- service.advertise_options.advertise_on_public.port.port.info.same_as_port

<a id="canonical-1203232311133200-1303110232313003-3032022132022011-0113320003100111-0003011210301210-3101233113012030-1330211020032120-2211133320311120"></a>

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
same_as_port = {}
```

<a id="canonical-1011030330313011-1131101223000103-3320300023212323-2231002202313113-1132203022021232-0301112030002220-0011233300003111-3032200000311221"></a>

## Direct properties — same_as_port / 001031123321 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1131012322200133-0230123022020021-3320021030232212-2213011032002322-1133323300110130-1231212002003131-2002200131321130-1102020101112201"></a>

## Next pages — same_as_port / 001031123321 / 4

- [service.advertise_options.advertise_on_public.port.port.info](resources--workload--reference--group-015.md#canonical-2233321102202333-3221231301223203-0103122201230313-0002133112312132-3302302022032303-0131103100030133-2003311020201130-2300012000001331)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-1002130311113332-2130231123003000-2113012220322103-2102212131300021-3233123213121332-3000203213103023-2322231100121132-3001110212313022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3120111322211202-3200223123302200-2322013020022011-1121103302013103-3331110331013200-2332103320320002-0310223230022222-1223232101031001"></a>

## service.advertise_options.advertise_on_public.port.tcp_loadbalancer — tcp_loadbalancer / 100121200131 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.port](resources--workload--reference--group-012.md#canonical-3013031222022030-2021222233122213-2220022232301003-3102220000332302-0311213311310132-3133321213021123-0133300222212330-1122131003023301)
- service.advertise_options.advertise_on_public.port.tcp_loadbalancer

<a id="canonical-0210102032213011-2123223030022110-2011102030200000-1231312323210001-0110300113312013-3111010003113333-2313133222323102-2113122110112213"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for tcp loadbalancer.

Upstream description:

TCP loadbalancer.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
tcp_loadbalancer {
  # Configure direct properties listed below.
}
```

<a id="canonical-3031000320111220-0133100313202130-3202302032220001-2332000010221322-0313110231021303-0022102320011303-1212200123121122-0131231002331213"></a>

## Direct properties — tcp_loadbalancer / 100121200131 / 3

<a id="canonical-0012232313331122-1002003310301323-2032113323223103-3112010003312312-2100222211001120-2302331300231121-1113300222203322-1110203131130121"></a>

<a id="canonical-0103300113021110-2103022231003103-3031003321022232-1323230212202213-0201012310221010-1010323220131123-3020031013232222-2200213300220211"></a>

## domains property — tcp_loadbalancer / 100121200131 / 4

Type: `["list", "string"]`. Optional.

List of additional domains (host/authority header) that will be matched to this loadbalancer.
Domains are also used for SNI matching if the is true Domains also indicate the list of names for
which DNS resolution will be done by VER.

Upstream description:

A list of additional domains (host/authority header) that will be matched to this loadbalancer.

Domains are also used for SNI matching if the \`with\_sni\` is true Domains also indicate the list
of names for which DNS resolution will be done by VER.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeBetween(1, 32),
}
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
    "minItems": 1,
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.hostname": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.hostname": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3221312030111230-0002220331033001-2212230231002311-1112120023103333-0213020313230023-0323331320100323-0231333210101033-0010302221301302"></a>

<a id="canonical-2333332233000300-0213232102113212-2330001310003302-1021010211011013-2110331003301202-0322331301003010-1012232222202320-2111232313232200"></a>

## with_sni property — tcp_loadbalancer / 100121200131 / 5

Type: `"bool"`. Optional.

Set to true to enable TCP loadbalancer with SNI.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2220231031231022-0033233201133200-1313303311112102-0313313012203111-0003321221120322-0331122211012012-3113223313012101-1333001322333122"></a>

## Next pages — tcp_loadbalancer / 100121200131 / 6

- [service.advertise_options.advertise_on_public.port](resources--workload--reference--group-012.md#canonical-3013031222022030-2021222233122213-2220022232301003-3102220000332302-0311213311310132-3133321213021123-0133300222212330-1122131003023301)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-0310302023323212-1013021312131132-1222231111312001-2121122023132203-1332202130321213-0133012220300033-3111032100113201-3031211330022333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2211200300331112-0121012000221020-2102102121013321-1330033202130213-0133122220310202-0211111211303123-1321022321112233-0133032130330220"></a>

## service.advertise_options.do_not_advertise — do_not_advertise / 311131132322 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- service.advertise_options.do_not_advertise

<a id="canonical-3113320122112331-3321113113220313-3321212300122110-3033313203323201-0123231221102313-1233322232202222-0231222002110203-2023132231122111"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for do not advertise.

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
do_not_advertise = {}
```

<a id="canonical-3021123203311001-3002331030131030-3312211212112200-3001001121220003-2003023122310323-0111131103201033-0131023201222100-2231000222121212"></a>

## Direct properties — do_not_advertise / 311131132322 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0133213211221301-3010123320103233-1233111233033101-0030322233110120-2311012023201103-3233320232101223-2210231121133113-0333210000213003"></a>

## Next pages — do_not_advertise / 311131132322 / 4

- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-0120300131213121-2030323030023013-0301331320133010-0123120231213121-1220222031210011-1001323001213020-0213022223133211-2110030001312333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1031223320111223-1112332120313322-0022003010010323-3331101103101030-2033333120313331-2223222202212220-2133303201032310-0112130303023220"></a>

## service.configuration — configuration / 333021211033 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- service.configuration

<a id="canonical-3133203200133120-2233323103122130-2111230332020000-0221132131010212-1033121330121223-1011213101302233-2333032213032013-0103222022102301"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameters of the workload.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
configuration {
  # Configure direct properties listed below.
}
```

<a id="canonical-0110332311312213-1201120011110333-0112031331332200-1002231300002311-1221231012031011-3223303301000031-0332011321203301-3020010321313033"></a>

## Direct properties — configuration / 333021211033 / 3

- [parameters](resources--workload--reference--group-015.md#canonical-2322332321021201-2221302122111313-1123221021001331-0221331123102121-2003331311023002-0203023003323011-2000101302102110-0101112103120131): complete subsection reference.

<a id="canonical-2100321010211213-3213313201100313-1232010113101321-3032222231330212-1230112200333011-1012012032303032-2213022003032131-0310331221000103"></a>

## Next pages — configuration / 333021211033 / 4

- [service.configuration.parameters](resources--workload--reference--group-015.md#canonical-2322332321021201-2221302122111313-1123221021001331-0221331123102121-2003331311023002-0203023003323011-2000101302102110-0101112103120131)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-2322332321021201-2221302122111313-1123221021001331-0221331123102121-2003331311023002-0203023003323011-2000101302102110-0101112103120131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1232233002203203-1112211120331313-0112123220232112-3212303000203331-3333321321313322-2201031010130323-0203003220123310-2310012233110212"></a>

## service.configuration.parameters — parameters / 231133121320 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.configuration](resources--workload--reference--group-015.md#canonical-0120300131213121-2030323030023013-0301331320133010-0123120231213121-1220222031210011-1001323001213020-0213022223133211-2110030001312333)
- service.configuration.parameters

<a id="canonical-3003313021012223-2121021311131113-1312321100131311-3210200313331131-0023122100032211-1121103323111313-3013210321310320-3103031010230011"></a>

Type: `"object"`. list nested block, Optional.

Parameters. Parameters for the workload.

Upstream description:

Parameters for the workload.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.ConflictingListObjectAttributes("env_var",
    "file")}
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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
parameters {
  # Configure direct properties listed below.
}
```

<a id="canonical-3312103310212212-3213110112023023-3011210313303310-3201213023132113-3111222300120311-0032132020310301-1001332311120211-3302113311020123"></a>

## Direct properties — parameters / 231133121320 / 3

- [env_var](resources--workload--reference--group-015.md#canonical-0002330111303113-2303101113120111-1131202032110112-1011310002132032-0223301320220200-0232133022022012-3233133112320313-3030030123312120): complete subsection reference.

- [file](resources--workload--reference--group-015.md#canonical-0021330302133213-0100230320121003-2021132302312123-1030102101111211-3003210002310101-2201203320301212-2213131323001333-0231100302122010): complete subsection reference.

<a id="canonical-2321221321132113-2101131101130122-3110020033333331-2200123303310202-0030223133121220-1300103033021320-3132020311233330-3220101100000200"></a>

## Next pages — parameters / 231133121320 / 4

- [service.configuration.parameters.env_var](resources--workload--reference--group-015.md#canonical-0002330111303113-2303101113120111-1131202032110112-1011310002132032-0223301320220200-0232133022022012-3233133112320313-3030030123312120)
- [service.configuration.parameters.file](resources--workload--reference--group-015.md#canonical-0021330302133213-0100230320121003-2021132302312123-1030102101111211-3003210002310101-2201203320301212-2213131323001333-0231100302122010)
- [service.configuration](resources--workload--reference--group-015.md#canonical-0120300131213121-2030323030023013-0301331320133010-0123120231213121-1220222031210011-1001323001213020-0213022223133211-2110030001312333)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-0002330111303113-2303101113120111-1131202032110112-1011310002132032-0223301320220200-0232133022022012-3233133112320313-3030030123312120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1200132012111203-2203110132102222-0213122123222021-3001320020010312-3002133112123221-0212311132333123-3301223221231303-3210122120103133"></a>

## service.configuration.parameters.env_var — env_var / 330103110202 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.configuration](resources--workload--reference--group-015.md#canonical-0120300131213121-2030323030023013-0301331320133010-0123120231213121-1220222031210011-1001323001213020-0213022223133211-2110030001312333)
- [service.configuration.parameters](resources--workload--reference--group-015.md#canonical-2322332321021201-2221302122111313-1123221021001331-0221331123102121-2003331311023002-0203023003323011-2000101302102110-0101112103120131)
- service.configuration.parameters.env_var

<a id="canonical-3322000301223221-0330030320130020-0332302031231003-3221312110123100-3213013112021132-0211032231032312-0300312022012220-1212221321002220"></a>

Type: `"object"`. single nested block, Optional.

Environment Variable. Environment Variable.

Upstream description:

Environment Variable.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
env_var {
  # Configure direct properties listed below.
}
```

<a id="canonical-0133221020030332-0332230303033331-3220211313310330-0312122230213022-0112022213110333-0033010321313031-0322230101102021-0102110102233300"></a>

## Direct properties — env_var / 330103110202 / 3

<a id="canonical-1300211131023212-1130123112220131-0123110011230031-2033120231313022-0031302232300233-2031121323211103-1103023110201020-3031233231020003"></a>

<a id="canonical-3023212001202301-2001200112110102-1003001221123223-1220033203321021-1133002130121221-0211000032023233-1303022312011231-1130132302032310"></a>

## name property — env_var / 330103110202 / 4

Type: `"string"`. Optional.

Name. Name of Environment Variable.

Upstream description:

Name of Environment Variable.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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
    "maxLength": 256,
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
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-1122003130121122-1213202012331021-2320233113012033-0212033031311321-0211013203332311-0302332321201000-1131300233032211-0302201223323202"></a>

<a id="canonical-1230110232011103-0032332312020222-1203100010311103-0213333322210300-2010112221300030-2032103032301123-2311032031203003-2313122102111100"></a>

## value property — env_var / 330103110202 / 5

Type: `"string"`. Optional.

Value. Value of Environment Variable.

Upstream description:

Value of Environment Variable.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-1101223032302221-0302200331113131-2111303223003102-1011321032302030-3303200112020312-2130121122132000-1202220111013310-0013331013112003"></a>

## Next pages — env_var / 330103110202 / 6

- [service.configuration.parameters](resources--workload--reference--group-015.md#canonical-2322332321021201-2221302122111313-1123221021001331-0221331123102121-2003331311023002-0203023003323011-2000101302102110-0101112103120131)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-0021330302133213-0100230320121003-2021132302312123-1030102101111211-3003210002310101-2201203320301212-2213131323001333-0231100302122010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3130120310020010-3331120221111102-0102031332010321-1132203103113111-1100002121023303-2110000020202201-0303020111133322-0101311110110030"></a>

## service.configuration.parameters.file — file / 200022002213 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.configuration](resources--workload--reference--group-015.md#canonical-0120300131213121-2030323030023013-0301331320133010-0123120231213121-1220222031210011-1001323001213020-0213022223133211-2110030001312333)
- [service.configuration.parameters](resources--workload--reference--group-015.md#canonical-2322332321021201-2221302122111313-1123221021001331-0221331123102121-2003331311023002-0203023003323011-2000101302102110-0101112103120131)
- service.configuration.parameters.file

<a id="canonical-0133202021223320-0003010301223023-0002022323323000-2210213031232232-2103123133222233-1222122200003311-3001323122303213-2200113233230320"></a>

Type: `"object"`. single nested block, Optional.

Configuration File. Configuration File for the workload.

Upstream description:

Configuration File for the workload.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("name",
    "volume_name")}
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
file {
  # Configure direct properties listed below.
}
```

<a id="canonical-3200222312313310-2121011213301030-3310223033133122-2202011230123121-2003002032212330-1320003233130300-0300132310212302-2210022013013303"></a>

## Direct properties — file / 200022002213 / 3

<a id="canonical-0201311023131330-0110210010021102-2120002100203320-2230213213223100-3002130302312310-2311012322213132-2133123312202111-0222212330310332"></a>

<a id="canonical-1031230013000210-1202210333030201-2200113101202223-0100233130203211-2221300221301232-2220302102003122-2011022302011021-2130120002323022"></a>

## data property — file / 200022002213 / 4

Type: `"string"`. Optional.

Data. File data

Upstream description:

File data

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(16384),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 16384,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "maxLength": 16384,
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
    "ves.io.schema.rules.string.max_len": "16384",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "16384",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

- [mount](resources--workload--reference--group-015.md#canonical-2031313210013212-2213211123312132-1101031213300321-1021132030122012-2322101220331311-2330202332012212-2201121213103233-2113120200323020): complete subsection reference.

<a id="canonical-2013331200302203-0310023312320111-2123312020121333-2130212313002010-1010331000230201-0231202300312000-0000010330130330-0320022002102203"></a>

<a id="canonical-2131200210121211-2231210020330300-3222121233021023-1221323010230331-2030121133103231-3230230232233022-2123332112111203-0302333122332100"></a>

## name property — file / 200022002213 / 5

Type: `"string"`. Optional.

Name. Name of the file.

Upstream description:

Name of the file.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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
    "maxLength": 256,
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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-1201200212330221-2210222111333302-1301322101121330-3112313323322032-0130321222220020-2010023000031123-2302022131312113-3332022002201033"></a>

<a id="canonical-3233332031211312-3213111030311032-0232001122313221-3310112203130002-3031313001013100-2232131132223223-3233022201320123-2311200310121222"></a>

## volume_name property — file / 200022002213 / 6

Type: `"string"`. Optional.

Volume Name. Name of the Volume.

Upstream description:

Name of the Volume.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-1233133132233002-2121310333102112-3011012112000113-1230223332200202-2300103310331322-2201202201022011-2001001230313021-1033301223302001"></a>

## Next pages — file / 200022002213 / 7

- [service.configuration.parameters.file.mount](resources--workload--reference--group-015.md#canonical-2031313210013212-2213211123312132-1101031213300321-1021132030122012-2322101220331311-2330202332012212-2201121213103233-2113120200323020)
- [service.configuration.parameters](resources--workload--reference--group-015.md#canonical-2322332321021201-2221302122111313-1123221021001331-0221331123102121-2003331311023002-0203023003323011-2000101302102110-0101112103120131)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-2031313210013212-2213211123312132-1101031213300321-1021132030122012-2322101220331311-2330202332012212-2201121213103233-2113120200323020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1302320220222022-3310121310111102-0102232031120122-2113131111023221-3330132222130233-1021131102313303-0210231111013103-3301031230232301"></a>

## service.configuration.parameters.file.mount — mount / 303322302031 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.configuration](resources--workload--reference--group-015.md#canonical-0120300131213121-2030323030023013-0301331320133010-0123120231213121-1220222031210011-1001323001213020-0213022223133211-2110030001312333)
- [service.configuration.parameters](resources--workload--reference--group-015.md#canonical-2322332321021201-2221302122111313-1123221021001331-0221331123102121-2003331311023002-0203023003323011-2000101302102110-0101112103120131)
- [service.configuration.parameters.file](resources--workload--reference--group-015.md#canonical-0021330302133213-0100230320121003-2021132302312123-1030102101111211-3003210002310101-2201203320301212-2213131323001333-0231100302122010)
- service.configuration.parameters.file.mount

<a id="canonical-0111202111330230-0110121212100030-0013211312323320-3212300213133331-2322221111200102-2232123333011002-0213123033221321-1330132102032232"></a>

Type: `"object"`. single nested block, Optional.

Volume mount describes how volume is mounted inside a workload.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("mount_path")}
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
mount {
  # Configure direct properties listed below.
}
```

<a id="canonical-1213333122031131-0323123122230130-1333210320022031-1001200010300101-3312013201112100-1300201231233022-0330322221300323-1011001012212320"></a>

## Direct properties — mount / 303322302031 / 3

<a id="canonical-2022021302222303-3133010331233203-1230120231221311-2023022123221130-3201333313132223-1221113030222110-2022220023301202-0310310120203022"></a>

<a id="canonical-0103011110232313-0012131333223120-2003303023100233-3233100301213100-2000102212220312-2132001210212231-1001200313131310-1111132312312102"></a>

## mode property — mount / 303322302031 / 4

Type: `"string"`. Optional.

\[Enum: VOLUME\_MOUNT\_READ\_ONLY|VOLUME\_MOUNT\_READ\_WRITE\] Mode in which the volume should be
mounted to the workload - VOLUME\_MOUNT\_READ\_ONLY: ReadOnly Mount the volume in read-only mode -
VOLUME\_MOUNT\_READ\_WRITE: Read Write Mount the volume in read-write mode. Possible values are
\`VOLUME\_MOUNT\_READ\_ONLY\`, \`VOLUME\_MOUNT\_READ\_WRITE\`. Defaults to
\`VOLUME\_MOUNT\_READ\_ONLY\`.

Upstream description:

Mode in which the volume should be mounted to the workload

&#8203;- VOLUME\_MOUNT\_READ\_ONLY: ReadOnly

Mount the volume in read-only mode &#8203;- VOLUME\_MOUNT\_READ\_WRITE: Read Write

Mount the volume in read-write mode.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["VOLUME_MOUNT_READ_ONLY","VOLUME_MOUNT_READ_WRITE"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
Validators: []validator.String{
  stringvalidator.OneOf("VOLUME_MOUNT_READ_ONLY",
    "VOLUME_MOUNT_READ_WRITE"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "VOLUME_MOUNT_READ_ONLY",
  "enum": [
    "VOLUME_MOUNT_READ_ONLY",
    "VOLUME_MOUNT_READ_WRITE"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0133301320110013-2202122301333013-3210022121102332-3320112011320030-1132112303333301-1110130310010101-3321223313231110-2200301203031130"></a>

<a id="canonical-3211322213123333-0213032020221023-2130231303313101-0133210113131203-3010211330222213-2201330301001300-3231101010021203-0300232311200223"></a>

## mount_path property — mount / 303322302031 / 5

Type: `"string"`. Optional.

Path within the workload container at which the volume should be mounted. Must not contain ':'.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "pattern": "^[^:]*$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.pattern": "^[^:]*$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.pattern": "^[^:]*$"
  }
}
```

<a id="canonical-2212103321130031-3101220312101210-0033010001332333-3033000201203131-2233211013132002-3001200132321220-3313212200000122-1200320213120001"></a>

<a id="canonical-1133301330010312-1323300202103333-2211312222122212-3320232232110300-1111121310010123-2321301212211021-1000313203211122-2201323013211020"></a>

## sub_path property — mount / 303322302031 / 6

Type: `"string"`. Optional.

Path within the volume from which the workload's volume should be mounted. Defaults to '' (volume's
root).

Upstream description:

Path within the volume from which the workload's volume should be mounted. Defaults to "" (volume's
root).

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```
