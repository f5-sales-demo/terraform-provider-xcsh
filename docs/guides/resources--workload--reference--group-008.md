---
page_title: "xcsh_workload reference"
subcategory: "Container"
description: "Complete grouped canonical reference for xcsh_workload reference."
---

# xcsh_workload reference

<a id="canonical-3011211302110102-1110003330130012-2330003231002123-1333212203033211-1212101000211203-3333313212202213-1122232332203332-0220110332323232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-2121230312030330-3010113231323001-1233013032320322-1032203002210311-2022021002113232-1123312032112001-0012103213230102-1112311001110020)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-006.md#canonical-0133221101333210-0222103021001232-3212010032013103-3331120032012221-1012203003232110-1233102323001233-2233321121201101-0122313122210200)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-006.md#canonical-2311001323203002-2023333020201320-1330320322101233-1033123102030111-2013210002333120-2130122003123103-1133301323123220-2302130221012102)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes

<a id="canonical-1220233233300032-0113322210322011-0123210001203323-3232103110313331-3330130123333031-1010112103100003-2011332201212333-2120002321201123"></a>

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

<a id="canonical-3101000322210023-1012323321311133-1001221313332121-2022013022303223-3320203323232203-0200212220111202-2102110021333223-2220202203300212"></a>

### Direct properties for `service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes`

- [routes](resources--workload--reference--group-008.md#canonical-3101223212120012-3332300223001100-0333220310012333-1031332113331120-2210213133122022-1333032223122231-2030231210333100-1103221131133311): complete subsection reference.

<a id="canonical-3101223212120012-3332300223001100-0333220310012333-1031332113331120-2210213133122022-1333032223122231-2030231210333100-1103221131133311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-2121230312030330-3010113231323001-1233013032320322-1032203002210311-2022021002113232-1123312032112001-0012103213230102-1112311001110020)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-006.md#canonical-0133221101333210-0222103021001232-3212010032013103-3331120032012221-1012203003232110-1233102323001233-2233321121201101-0122313122210200)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-006.md#canonical-2311001323203002-2023333020201320-1330320322101233-1033123102030111-2013210002333120-2130122003123103-1133301323123220-2302130221012102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes](resources--workload--reference--group-008.md#canonical-3011211302110102-1110003330130012-2330003231002123-1333212203033211-1212101000211203-3333313212202213-1122232332203332-0220110332323232)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes

<a id="canonical-0320120233011301-2200031330002110-3221121311302333-3003121232230310-3332323122323322-2022332313211203-2001121302330031-1311110103303222"></a>

Type: `"object"`. list nested block, Optional.

Routes. Routes for this loadbalancer.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-1310221200033313-0113301012303323-1000120311123330-3032330121123321-0311003010331312-0030322211311103-3100313303303112-1020220203323213"></a>

### Direct properties for `service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes`

- [custom_route_object](resources--workload--reference--group-008.md#canonical-2000122322012020-3202320303320213-0231013331320221-2023031002230033-2112230020111131-0300200113021323-3020232330312200-0223013321012112): complete subsection reference.

- [direct_response_route](resources--workload--reference--group-008.md#canonical-3231322330202023-0011210010031213-0320313110113210-0323302203013003-1003131123333113-2031030012023021-3201100012202203-0030301312131222): complete subsection reference.

- [redirect_route](resources--workload--reference--group-008.md#canonical-0003020102311303-0233233032322332-3002231010033312-1313303221323312-0103013131011120-3102111011102132-1313020300032233-1111311233230113): complete subsection reference.

- [simple_route](resources--workload--reference--group-008.md#canonical-2332230303301330-0120223130133121-2131103033210110-0000331031322110-2102032123012022-0000121011033231-1222320000321202-3323210133011002): complete subsection reference.

<a id="canonical-2000122322012020-3202320303320213-0231013331320221-2023031002230033-2112230020111131-0300200113021323-3020232330312200-0223013321012112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.custom_route_object` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-2121230312030330-3010113231323001-1233013032320322-1032203002210311-2022021002113232-1123312032112001-0012103213230102-1112311001110020)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-006.md#canonical-0133221101333210-0222103021001232-3212010032013103-3331120032012221-1012203003232110-1233102323001233-2233321121201101-0122313122210200)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-006.md#canonical-2311001323203002-2023333020201320-1330320322101233-1033123102030111-2013210002333120-2130122003123103-1133301323123220-2302130221012102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes](resources--workload--reference--group-008.md#canonical-3011211302110102-1110003330130012-2330003231002123-1333212203033211-1212101000211203-3333313212202213-1122232332203332-0220110332323232)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes](resources--workload--reference--group-008.md#canonical-3101223212120012-3332300223001100-0333220310012333-1031332113331120-2210213133122022-1333032223122231-2030231210333100-1103221131133311)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.custom_route_object

<a id="canonical-1213010101330121-1132211303120323-2212112310000322-0211021200113210-1032301122302013-0020112303011230-0001012023302322-2121330103223001"></a>

Type: `"object"`. single nested block, Optional.

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

<a id="canonical-3200032221212301-0002223032022301-3231323220302100-3023222222100133-2322030302031101-1320220222110013-0132110212122031-0123002222131110"></a>

### Direct properties for `service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.custom_route_object`

- [caching_disable](resources--workload--reference--group-008.md#canonical-2013012230220021-3011122033311001-0101101133003320-1011321003233233-1110102020203011-3021221203323112-2313303310222023-0000221032002120): complete subsection reference.

- [caching_inherit](resources--workload--reference--group-008.md#canonical-3230112111031231-2002111310330022-0223003233132003-2212133100121201-3032223120100312-2001123022122023-2121112123003212-3102202133020102): complete subsection reference.

- [route_ref](resources--workload--reference--group-008.md#canonical-3113101301231330-1201332333231130-2122102231230121-2033310222131210-0021302132233220-0032213131220003-2201110210212221-1233123001332220): complete subsection reference.

<a id="canonical-2013012230220021-3011122033311001-0101101133003320-1011321003233233-1110102020203011-3021221203323112-2313303310222023-0000221032002120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.custom_route_object.caching_disable` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-2121230312030330-3010113231323001-1233013032320322-1032203002210311-2022021002113232-1123312032112001-0012103213230102-1112311001110020)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-006.md#canonical-0133221101333210-0222103021001232-3212010032013103-3331120032012221-1012203003232110-1233102323001233-2233321121201101-0122313122210200)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-006.md#canonical-2311001323203002-2023333020201320-1330320322101233-1033123102030111-2013210002333120-2130122003123103-1133301323123220-2302130221012102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes](resources--workload--reference--group-008.md#canonical-3011211302110102-1110003330130012-2330003231002123-1333212203033211-1212101000211203-3333313212202213-1122232332203332-0220110332323232)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes](resources--workload--reference--group-008.md#canonical-3101223212120012-3332300223001100-0333220310012333-1031332113331120-2210213133122022-1333032223122231-2030231210333100-1103221131133311)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.custom_route_object](resources--workload--reference--group-008.md#canonical-2000122322012020-3202320303320213-0231013331320221-2023031002230033-2112230020111131-0300200113021323-3020232330312200-0223013321012112)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.custom_route_object.caching_disable

<a id="canonical-3101000232011003-1012310213111031-2311001000121111-2022312233213323-0002322333112000-2323211101031320-0102233320331323-2312021121110303"></a>

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

<a id="canonical-3230112111031231-2002111310330022-0223003233132003-2212133100121201-3032223120100312-2001123022122023-2121112123003212-3102202133020102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.custom_route_object.caching_inherit` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-2121230312030330-3010113231323001-1233013032320322-1032203002210311-2022021002113232-1123312032112001-0012103213230102-1112311001110020)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-006.md#canonical-0133221101333210-0222103021001232-3212010032013103-3331120032012221-1012203003232110-1233102323001233-2233321121201101-0122313122210200)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-006.md#canonical-2311001323203002-2023333020201320-1330320322101233-1033123102030111-2013210002333120-2130122003123103-1133301323123220-2302130221012102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes](resources--workload--reference--group-008.md#canonical-3011211302110102-1110003330130012-2330003231002123-1333212203033211-1212101000211203-3333313212202213-1122232332203332-0220110332323232)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes](resources--workload--reference--group-008.md#canonical-3101223212120012-3332300223001100-0333220310012333-1031332113331120-2210213133122022-1333032223122231-2030231210333100-1103221131133311)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.custom_route_object](resources--workload--reference--group-008.md#canonical-2000122322012020-3202320303320213-0231013331320221-2023031002230033-2112230020111131-0300200113021323-3020232330312200-0223013321012112)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.custom_route_object.caching_inherit

<a id="canonical-1000031002120023-3010210303203322-2011010020333021-3322102033332322-2121201032023213-2032220013121012-2022213321023310-0012022223202330"></a>

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

<a id="canonical-3113101301231330-1201332333231130-2122102231230121-2033310222131210-0021302132233220-0032213131220003-2201110210212221-1233123001332220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.custom_route_object.route_ref` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-2121230312030330-3010113231323001-1233013032320322-1032203002210311-2022021002113232-1123312032112001-0012103213230102-1112311001110020)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-006.md#canonical-0133221101333210-0222103021001232-3212010032013103-3331120032012221-1012203003232110-1233102323001233-2233321121201101-0122313122210200)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-006.md#canonical-2311001323203002-2023333020201320-1330320322101233-1033123102030111-2013210002333120-2130122003123103-1133301323123220-2302130221012102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes](resources--workload--reference--group-008.md#canonical-3011211302110102-1110003330130012-2330003231002123-1333212203033211-1212101000211203-3333313212202213-1122232332203332-0220110332323232)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes](resources--workload--reference--group-008.md#canonical-3101223212120012-3332300223001100-0333220310012333-1031332113331120-2210213133122022-1333032223122231-2030231210333100-1103221131133311)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.custom_route_object](resources--workload--reference--group-008.md#canonical-2000122322012020-3202320303320213-0231013331320221-2023031002230033-2112230020111131-0300200113021323-3020232330312200-0223013321012112)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.custom_route_object.route_ref

<a id="canonical-3101000000003223-3313301122313200-3212222121201112-0210333232322233-3010203312313022-0013101003011202-1021033020321020-0031120202012323"></a>

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
route_ref {
  # Configure direct properties listed below.
}
```

<a id="canonical-2232130131331230-2011200302201201-2031230111101332-2311203321212110-2102210102103111-0100230202002332-2130031033013311-0301321003113110"></a>

### Direct properties for `service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.custom_route_object.route_ref`

<a id="canonical-1300230132032100-1233103312230121-1330231011230320-0231131132022231-0331302321233111-0212002322112220-0213231210032131-0103002031320303"></a>

#### `service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.custom_route_object.route_ref.name` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-1202213111331013-2010300002022202-1021233331003213-1201120133010331-3313133011312223-1313021222100200-1212010203231000-2120022111021120"></a>

<a id="canonical-3223332130003210-3211330203101122-2033212123303302-2221013013132202-3321213223103111-2022130202003033-0231022332000333-0221003111310213"></a>

#### `service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.custom_route_object.route_ref.namespace` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-0321133221100302-1132230203023302-3023002002201200-1231302212232011-1203132231333133-0211102202302333-1001303003120222-0232222330130212"></a>

<a id="canonical-2020312123222020-1131223102001003-1333230333231022-0322203200021203-2121222131002031-0110220323130322-2132023223130223-2003020013000311"></a>

#### `service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.custom_route_object.route_ref.tenant` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-3231322330202023-0011210010031213-0320313110113210-0323302203013003-1003131123333113-2031030012023021-3201100012202203-0030301312131222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.direct_response_route` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-2121230312030330-3010113231323001-1233013032320322-1032203002210311-2022021002113232-1123312032112001-0012103213230102-1112311001110020)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-006.md#canonical-0133221101333210-0222103021001232-3212010032013103-3331120032012221-1012203003232110-1233102323001233-2233321121201101-0122313122210200)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-006.md#canonical-2311001323203002-2023333020201320-1330320322101233-1033123102030111-2013210002333120-2130122003123103-1133301323123220-2302130221012102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes](resources--workload--reference--group-008.md#canonical-3011211302110102-1110003330130012-2330003231002123-1333212203033211-1212101000211203-3333313212202213-1122232332203332-0220110332323232)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes](resources--workload--reference--group-008.md#canonical-3101223212120012-3332300223001100-0333220310012333-1031332113331120-2210213133122022-1333032223122231-2030231210333100-1103221131133311)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.direct_response_route

<a id="canonical-3211110021100023-0221220022131132-0122221001312113-3202303022010020-0121031222332001-1301112220120032-3030123031120313-1211202222001000"></a>

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

<a id="canonical-0332323123102330-2222333310030030-2212220003112233-3021333000131003-2233231010330321-0032312103321100-3131230331230022-2322321003122330"></a>

### Direct properties for `service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.direct_response_route`

- [headers](resources--workload--reference--group-008.md#canonical-2020230033132332-1001001321102332-2222213203320321-0201222313012110-2023333030020232-0311113122030231-2002113032210232-0123330331111230): complete subsection reference.

<a id="canonical-1102233132121102-0010213332123220-1012302321020230-3302321311110122-3032321111113321-1320300233212132-3013232123332321-1101230310110020"></a>

<a id="canonical-2211002031011321-0323012001332133-1203113000002310-1012021101100311-1221023332100013-0233132332013031-2001110030230023-1332001323103302"></a>

#### `service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.direct_response_route.http_method` property

Type: `"string"`. Optional.

\[Enum: ANY|GET|HEAD|POST|PUT|DELETE|CONNECT|OPTIONS|TRACE|PATCH|COPY\] Specifies the HTTP method
used to access a resource. Any HTTP Method. Possible values are \`ANY\`, \`GET\`, \`HEAD\`,
\`POST\`, \`PUT\`, \`DELETE\`, \`CONNECT\`, \`OPTIONS\`, \`TRACE\`, \`PATCH\`, \`COPY\`. Defaults to
\`ANY\`.

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

- [incoming_port](resources--workload--reference--group-008.md#canonical-2123033123102120-2331000002131212-0121121320112021-3030313331301033-0030311130230101-3013220213011320-3010213301303110-2332103131000202): complete subsection reference.

- [path](resources--workload--reference--group-008.md#canonical-2101000200001200-2113213200123021-3031201121320103-2232230132032332-0021113020023231-1302312203103121-1220332011010133-0233330320122213): complete subsection reference.

- [route_direct_response](resources--workload--reference--group-008.md#canonical-0013230300013321-2022130123233200-1323122002332002-1233012001231310-3023020323223100-3202010020323203-3302111201322232-0013031221012013): complete subsection reference.

<a id="canonical-2020230033132332-1001001321102332-2222213203320321-0201222313012110-2023333030020232-0311113122030231-2002113032210232-0123330331111230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.direct_response_route.headers` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-2121230312030330-3010113231323001-1233013032320322-1032203002210311-2022021002113232-1123312032112001-0012103213230102-1112311001110020)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-006.md#canonical-0133221101333210-0222103021001232-3212010032013103-3331120032012221-1012203003232110-1233102323001233-2233321121201101-0122313122210200)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-006.md#canonical-2311001323203002-2023333020201320-1330320322101233-1033123102030111-2013210002333120-2130122003123103-1133301323123220-2302130221012102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes](resources--workload--reference--group-008.md#canonical-3011211302110102-1110003330130012-2330003231002123-1333212203033211-1212101000211203-3333313212202213-1122232332203332-0220110332323232)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes](resources--workload--reference--group-008.md#canonical-3101223212120012-3332300223001100-0333220310012333-1031332113331120-2210213133122022-1333032223122231-2030231210333100-1103221131133311)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.direct_response_route](resources--workload--reference--group-008.md#canonical-3231322330202023-0011210010031213-0320313110113210-0323302203013003-1003131123333113-2031030012023021-3201100012202203-0030301312131222)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.direct_response_route.headers

<a id="canonical-1201132131321133-1010222223133220-1113210112331222-1013210002320310-2212113022030213-0023323023211321-1200313112202322-3103111213022121"></a>

Type: `"object"`. list nested block, Optional.

Headers. List of (key, value) headers.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-2032002132121222-3321010003300303-0110302111000031-1100122133222111-3000313330113100-2011013233200330-1010323222131112-2322021300211130"></a>

### Direct properties for `service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.direct_response_route.headers`

<a id="canonical-0302110113300222-0102003031102102-1023303113110303-3333000033010012-2230111000132311-2110113322333213-2021023030221213-1302100310011112"></a>

#### `service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.direct_response_route.headers.exact` property

Type: `"string"`. Optional.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-0122113331211323-2112021221312330-2033021123313011-0200100010000223-1200333121101031-2203310112332023-1320231211202031-2221122310313330"></a>

<a id="canonical-3233220223302222-0333000220331013-1010230220111032-1313100002021133-0032322012023233-1010010321221013-3212323000110130-2200022321100120"></a>

#### `service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.direct_response_route.headers.invert_match` property

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

<a id="canonical-3220211321212022-0322132130213332-0022222120330203-1303210220002131-1330131200333233-2013311313322121-2323013102313101-3300223113030231"></a>

<a id="canonical-3001011102222013-0320121221032200-2323333131210220-0021231312032133-2013123230122203-2222032002323123-0322030202201321-0113122012121202"></a>

#### `service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.direct_response_route.headers.name` property

Type: `"string"`. Optional.

Name. Name of the header.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-1220321221031302-2312301220030020-1303011010001321-3103123111120022-3220010012021213-1030012122323132-0303132331320112-0000201111333011"></a>

<a id="canonical-2311111130233230-1033303320320133-3322130301022230-1211021302203123-0130030121213121-2202130000013222-1120101131313222-0002202230003303"></a>

#### `service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.direct_response_route.headers.presence` property

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

<a id="canonical-2300220201201032-0230130220330120-2103001332310311-0201113322002022-1222202132220001-3032123201303030-2011203000322310-3322233010222221"></a>

<a id="canonical-0011033133302130-0323012332201301-0320321132332131-0112023020233202-3213133213231002-1323103102001131-3013032103210230-2131230201221010"></a>

#### `service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.direct_response_route.headers.regex` property

Type: `"string"`. Optional.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-2123033123102120-2331000002131212-0121121320112021-3030313331301033-0030311130230101-3013220213011320-3010213301303110-2332103131000202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.direct_response_route.incoming_port` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-2121230312030330-3010113231323001-1233013032320322-1032203002210311-2022021002113232-1123312032112001-0012103213230102-1112311001110020)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-006.md#canonical-0133221101333210-0222103021001232-3212010032013103-3331120032012221-1012203003232110-1233102323001233-2233321121201101-0122313122210200)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-006.md#canonical-2311001323203002-2023333020201320-1330320322101233-1033123102030111-2013210002333120-2130122003123103-1133301323123220-2302130221012102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes](resources--workload--reference--group-008.md#canonical-3011211302110102-1110003330130012-2330003231002123-1333212203033211-1212101000211203-3333313212202213-1122232332203332-0220110332323232)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes](resources--workload--reference--group-008.md#canonical-3101223212120012-3332300223001100-0333220310012333-1031332113331120-2210213133122022-1333032223122231-2030231210333100-1103221131133311)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.direct_response_route](resources--workload--reference--group-008.md#canonical-3231322330202023-0011210010031213-0320313110113210-0323302203013003-1003131123333113-2031030012023021-3201100012202203-0030301312131222)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.direct_response_route.incoming_port

<a id="canonical-2303132103213311-1000231102033103-2212303032001133-2302312321333023-1032012131222302-2110211220131001-1131230230222003-3101220122221132"></a>

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

<a id="canonical-3303031000131303-3032101333222203-0213220020021322-2312000203302031-1020032210101020-2332122033232220-3021201100321202-3303233231131310"></a>

### Direct properties for `service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.direct_response_route.incoming_port`

- [no_port_match](resources--workload--reference--group-008.md#canonical-3131312302120023-0023213202322211-3320302301002230-2100031001222002-3123021130020023-0020201232031211-0032120330112102-1013001332210322): complete subsection reference.

<a id="canonical-0220033222222000-1323300001310133-3330112331112321-2031010331111313-1211312032223301-2300110222320321-0001323201121223-3313001013220020"></a>

<a id="canonical-1032311211132100-2001321201031012-2110322323011003-0202112230022112-1201332012030001-2220132131332013-1231102300321010-0000102022001000"></a>

#### `service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.direct_response_route.incoming_port.port` property

Type: `"number"`. Optional.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-0312013231320213-2103301002122100-0323001120010333-0311312112333210-0011303200203023-3020101130031203-3202303233212233-2211002130120123"></a>

<a id="canonical-1031122121302213-2103320313121202-0100323122021230-1011331210202322-0022002112023023-0233003201033030-0110202100020220-3111300210210012"></a>

#### `service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.direct_response_route.incoming_port.port_ranges` property

Type: `"string"`. Optional.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-3131312302120023-0023213202322211-3320302301002230-2100031001222002-3123021130020023-0020201232031211-0032120330112102-1013001332210322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.direct_response_route.incoming_port.no_port_match` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-2121230312030330-3010113231323001-1233013032320322-1032203002210311-2022021002113232-1123312032112001-0012103213230102-1112311001110020)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-006.md#canonical-0133221101333210-0222103021001232-3212010032013103-3331120032012221-1012203003232110-1233102323001233-2233321121201101-0122313122210200)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-006.md#canonical-2311001323203002-2023333020201320-1330320322101233-1033123102030111-2013210002333120-2130122003123103-1133301323123220-2302130221012102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes](resources--workload--reference--group-008.md#canonical-3011211302110102-1110003330130012-2330003231002123-1333212203033211-1212101000211203-3333313212202213-1122232332203332-0220110332323232)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes](resources--workload--reference--group-008.md#canonical-3101223212120012-3332300223001100-0333220310012333-1031332113331120-2210213133122022-1333032223122231-2030231210333100-1103221131133311)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.direct_response_route](resources--workload--reference--group-008.md#canonical-3231322330202023-0011210010031213-0320313110113210-0323302203013003-1003131123333113-2031030012023021-3201100012202203-0030301312131222)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.direct_response_route.incoming_port](resources--workload--reference--group-008.md#canonical-2123033123102120-2331000002131212-0121121320112021-3030313331301033-0030311130230101-3013220213011320-3010213301303110-2332103131000202)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.direct_response_route.incoming_port.no_port_match

<a id="canonical-2230132213301013-0331122303010132-2210010330213301-3333120202021203-3213003313031321-1103313030012031-2300012301203332-1211130001113312"></a>

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
no_port_match = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2101000200001200-2113213200123021-3031201121320103-2232230132032332-0021113020023231-1302312203103121-1220332011010133-0233330320122213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.direct_response_route.path` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-2121230312030330-3010113231323001-1233013032320322-1032203002210311-2022021002113232-1123312032112001-0012103213230102-1112311001110020)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-006.md#canonical-0133221101333210-0222103021001232-3212010032013103-3331120032012221-1012203003232110-1233102323001233-2233321121201101-0122313122210200)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-006.md#canonical-2311001323203002-2023333020201320-1330320322101233-1033123102030111-2013210002333120-2130122003123103-1133301323123220-2302130221012102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes](resources--workload--reference--group-008.md#canonical-3011211302110102-1110003330130012-2330003231002123-1333212203033211-1212101000211203-3333313212202213-1122232332203332-0220110332323232)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes](resources--workload--reference--group-008.md#canonical-3101223212120012-3332300223001100-0333220310012333-1031332113331120-2210213133122022-1333032223122231-2030231210333100-1103221131133311)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.direct_response_route](resources--workload--reference--group-008.md#canonical-3231322330202023-0011210010031213-0320313110113210-0323302203013003-1003131123333113-2031030012023021-3201100012202203-0030301312131222)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.direct_response_route.path

<a id="canonical-2123020131221121-1002300202330020-0000013113231200-1033021112303311-1132333032120223-3332320321133223-1322330303201322-1321100332211120"></a>

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

<a id="canonical-1000210223301103-1201123230002021-1310013033311033-2103020310113300-0312302301032210-3220210032022011-2321133221132201-1132132101220200"></a>

### Direct properties for `service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.direct_response_route.path`

<a id="canonical-3001312333100113-0112003012122020-0321331121323200-1103223001332233-3112321211302013-0223100100133333-0000013300323300-2230321231131113"></a>

#### `service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.direct_response_route.path.path` property

Type: `"string"`. Optional.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-1223222002131121-3122213232100001-3223213030011320-1133033220022231-3113201221302313-3201311111333223-3213303022230022-1201223220321330"></a>

<a id="canonical-0333020300300031-2313130131010211-1103222221310021-3113023031123132-0102121110013302-1222303120000021-1023301212013002-2022010122201221"></a>

#### `service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.direct_response_route.path.prefix` property

Type: `"string"`. Optional.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-1313202203022221-2331233031230033-0012032030221013-1220313132130013-3110102203213232-0033300202013313-2132032103130113-3103312220330100"></a>

<a id="canonical-1000220132232201-1003200212323233-3012202102311333-3130022032133131-1022313231320013-0310213303333220-0112333103212100-0131102200132133"></a>

#### `service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.direct_response_route.path.regex` property

Type: `"string"`. Optional.

Exclusive with \[path prefix\] Regular expression of path match (e.g. The value .\* will match on
all paths).

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-0013230300013321-2022130123233200-1323122002332002-1233012001231310-3023020323223100-3202010020323203-3302111201322232-0013031221012013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.direct_response_route.route_direct_response` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-2121230312030330-3010113231323001-1233013032320322-1032203002210311-2022021002113232-1123312032112001-0012103213230102-1112311001110020)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-006.md#canonical-0133221101333210-0222103021001232-3212010032013103-3331120032012221-1012203003232110-1233102323001233-2233321121201101-0122313122210200)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-006.md#canonical-2311001323203002-2023333020201320-1330320322101233-1033123102030111-2013210002333120-2130122003123103-1133301323123220-2302130221012102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes](resources--workload--reference--group-008.md#canonical-3011211302110102-1110003330130012-2330003231002123-1333212203033211-1212101000211203-3333313212202213-1122232332203332-0220110332323232)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes](resources--workload--reference--group-008.md#canonical-3101223212120012-3332300223001100-0333220310012333-1031332113331120-2210213133122022-1333032223122231-2030231210333100-1103221131133311)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.direct_response_route](resources--workload--reference--group-008.md#canonical-3231322330202023-0011210010031213-0320313110113210-0323302203013003-1003131123333113-2031030012023021-3201100012202203-0030301312131222)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.direct_response_route.route_direct_response

<a id="canonical-0233203220000011-0001130103003212-2000231212202110-1010031332213132-1213233322131233-0130322113312221-0222120101111203-1211220202333001"></a>

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

<a id="canonical-1222023130000213-0011112130221320-1031002023111002-2003022213002210-3311132101231031-0203222011323302-3112101002221212-0103121202313231"></a>

### Direct properties for `service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.direct_response_route.route_direct_response`

<a id="canonical-1220302320123133-3233101031011110-3032023332211211-2132212232323223-1322111230301001-1320130103111010-3003223033321213-0331310022132001"></a>

#### `service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.direct_response_route.route_direct_response.response_body_encoded` property

Type: `"string"`. Optional.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-0023320222212213-1203202131121113-1202130100023100-1021102331210011-2222130120313320-0333011001303300-1312321201122211-1303330213202220"></a>

<a id="canonical-1033233032333023-0131330313032123-2120323100212323-3311022112130310-2320211320333311-0001131313302301-2032032023203121-0233213323330110"></a>

#### `service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.direct_response_route.route_direct_response.response_code` property

Type: `"number"`. Optional.

Response Code. Response code to send.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-0003020102311303-0233233032322332-3002231010033312-1313303221323312-0103013131011120-3102111011102132-1313020300032233-1111311233230113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-2121230312030330-3010113231323001-1233013032320322-1032203002210311-2022021002113232-1123312032112001-0012103213230102-1112311001110020)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-006.md#canonical-0133221101333210-0222103021001232-3212010032013103-3331120032012221-1012203003232110-1233102323001233-2233321121201101-0122313122210200)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-006.md#canonical-2311001323203002-2023333020201320-1330320322101233-1033123102030111-2013210002333120-2130122003123103-1133301323123220-2302130221012102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes](resources--workload--reference--group-008.md#canonical-3011211302110102-1110003330130012-2330003231002123-1333212203033211-1212101000211203-3333313212202213-1122232332203332-0220110332323232)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes](resources--workload--reference--group-008.md#canonical-3101223212120012-3332300223001100-0333220310012333-1031332113331120-2210213133122022-1333032223122231-2030231210333100-1103221131133311)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route

<a id="canonical-3020220301312121-3020000100131011-2310331123013200-0033301202202220-1333001211232333-0213113013311103-2220213123023211-1002120121331021"></a>

Type: `"object"`. single nested block, Optional.

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

<a id="canonical-2102113131010222-3033010133310113-2022311223023300-3231301001313220-2302111330113133-2010302111333000-1011203003100213-2012220102022100"></a>

### Direct properties for `service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route`

- [headers](resources--workload--reference--group-008.md#canonical-1232120323103103-1303003111002020-3203221110100320-2200310211301223-0000121223301023-2221011312120032-0212311100023331-1121330022320033): complete subsection reference.

<a id="canonical-1231033200123312-0233200123202200-1030020103202010-0230311302020000-3222222312320013-3332132222323002-3323211223013113-1202103222220312"></a>

<a id="canonical-1122210023323021-0301112002000111-0013330200312300-0130001033131133-0333112021132010-1123132301333123-3202101302303010-2311200330332223"></a>

#### `service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route.http_method` property

Type: `"string"`. Optional.

\[Enum: ANY|GET|HEAD|POST|PUT|DELETE|CONNECT|OPTIONS|TRACE|PATCH|COPY\] Specifies the HTTP method
used to access a resource. Any HTTP Method. Possible values are \`ANY\`, \`GET\`, \`HEAD\`,
\`POST\`, \`PUT\`, \`DELETE\`, \`CONNECT\`, \`OPTIONS\`, \`TRACE\`, \`PATCH\`, \`COPY\`. Defaults to
\`ANY\`.

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

- [incoming_port](resources--workload--reference--group-008.md#canonical-1020323201110112-2003203023313010-0130110100322020-2100332301223310-0121230331313023-1013011231112133-2010303211002033-3100032033120011): complete subsection reference.

- [path](resources--workload--reference--group-008.md#canonical-3233030021033313-3031030311221020-0200331022102113-1313122110031101-2210311230313221-2020211211322311-2201003031033002-3320231130112311): complete subsection reference.

- [route_redirect](resources--workload--reference--group-008.md#canonical-1013022013311110-1211131313223021-1020032322323123-0111133101310103-2132011220312120-1323130111012310-3301103301322333-3232211311000323): complete subsection reference.

<a id="canonical-1232120323103103-1303003111002020-3203221110100320-2200310211301223-0000121223301023-2221011312120032-0212311100023331-1121330022320033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route.headers` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-2121230312030330-3010113231323001-1233013032320322-1032203002210311-2022021002113232-1123312032112001-0012103213230102-1112311001110020)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-006.md#canonical-0133221101333210-0222103021001232-3212010032013103-3331120032012221-1012203003232110-1233102323001233-2233321121201101-0122313122210200)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-006.md#canonical-2311001323203002-2023333020201320-1330320322101233-1033123102030111-2013210002333120-2130122003123103-1133301323123220-2302130221012102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes](resources--workload--reference--group-008.md#canonical-3011211302110102-1110003330130012-2330003231002123-1333212203033211-1212101000211203-3333313212202213-1122232332203332-0220110332323232)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes](resources--workload--reference--group-008.md#canonical-3101223212120012-3332300223001100-0333220310012333-1031332113331120-2210213133122022-1333032223122231-2030231210333100-1103221131133311)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route](resources--workload--reference--group-008.md#canonical-0003020102311303-0233233032322332-3002231010033312-1313303221323312-0103013131011120-3102111011102132-1313020300032233-1111311233230113)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route.headers

<a id="canonical-3300232130201320-1112121222331233-1020330031322302-2000203110021332-2212011201312333-0100201203030132-2011302303300311-1013323111213331"></a>

Type: `"object"`. list nested block, Optional.

Headers. List of (key, value) headers.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-1012321133122110-1023121321323131-2302313123330122-3022031122030011-3201100132002210-1120233201213022-0023132332101302-3133333132311033"></a>

### Direct properties for `service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route.headers`

<a id="canonical-3212112000010320-3011200103333001-0111103201220322-0222220231001332-1111113122302102-0211003312022131-1310022122020323-2322123300301020"></a>

#### `service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route.headers.exact` property

Type: `"string"`. Optional.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-3112213101221221-1001120321012132-0010131211003220-1313002212110120-2322033231301033-3101003100132130-1301021302001031-3310221203023320"></a>

<a id="canonical-3223133032332120-2113131313232331-0231013020323020-0021001232202231-0010331302221103-2210002113001302-3302200112330332-3023320010003122"></a>

#### `service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route.headers.invert_match` property

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

<a id="canonical-1220032233020333-1123023303203230-3333101330010201-2331222113130322-0013210030313211-1312330221313312-1320221311012312-0301310311203333"></a>

<a id="canonical-0110023202032333-2111011210311320-3111023113100302-2231132032012331-3110003323320210-2031210331122331-2200133303132330-0201330102303012"></a>

#### `service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route.headers.name` property

Type: `"string"`. Optional.

Name. Name of the header.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-3212221311230031-0233123310132321-0223323223232010-1001303022130020-3112121230102100-1100221312002210-1332220331023103-0103003132022231"></a>

<a id="canonical-1122312102121322-0310012300202002-3020311312331230-3301103111031131-3002101023011332-0322230222200230-1012232322301111-1220212330200022"></a>

#### `service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route.headers.presence` property

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

<a id="canonical-0300211331230123-2321202210310200-0012123110120220-2010333133101023-2332002010203100-3300310113030002-2323010312023320-0332213200223301"></a>

<a id="canonical-0120233030132311-2001013211222203-0311110132021120-2303302311101011-2132201221203230-2311011202010232-1220102110220123-2310000123311022"></a>

#### `service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route.headers.regex` property

Type: `"string"`. Optional.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-1020323201110112-2003203023313010-0130110100322020-2100332301223310-0121230331313023-1013011231112133-2010303211002033-3100032033120011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route.incoming_port` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-2121230312030330-3010113231323001-1233013032320322-1032203002210311-2022021002113232-1123312032112001-0012103213230102-1112311001110020)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-006.md#canonical-0133221101333210-0222103021001232-3212010032013103-3331120032012221-1012203003232110-1233102323001233-2233321121201101-0122313122210200)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-006.md#canonical-2311001323203002-2023333020201320-1330320322101233-1033123102030111-2013210002333120-2130122003123103-1133301323123220-2302130221012102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes](resources--workload--reference--group-008.md#canonical-3011211302110102-1110003330130012-2330003231002123-1333212203033211-1212101000211203-3333313212202213-1122232332203332-0220110332323232)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes](resources--workload--reference--group-008.md#canonical-3101223212120012-3332300223001100-0333220310012333-1031332113331120-2210213133122022-1333032223122231-2030231210333100-1103221131133311)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route](resources--workload--reference--group-008.md#canonical-0003020102311303-0233233032322332-3002231010033312-1313303221323312-0103013131011120-3102111011102132-1313020300032233-1111311233230113)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route.incoming_port

<a id="canonical-2023130030033230-2302112202133011-0103103122313331-1223333110131013-2303331011030130-2202123002310330-2020312120002102-3312323311302120"></a>

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

<a id="canonical-0120212032122122-0011232222133211-0121223221123301-2332101332120100-0001311313310233-0031021200211223-3202213110131032-0203032132012232"></a>

### Direct properties for `service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route.incoming_port`

- [no_port_match](resources--workload--reference--group-008.md#canonical-1100123212233303-1101022230112200-1321101001212020-2330032102000112-0222001033210103-1132102002033200-3023130003310230-0112321320301123): complete subsection reference.

<a id="canonical-0013110110022131-3013002320130220-1121210213212330-2201223303100012-1321301110120101-2200113222213213-2000333331311203-3132303230113013"></a>

<a id="canonical-3332112320113211-3201002021023302-0321332203200301-1213022102102031-0001211302233211-3322011213323322-1000022013202311-0021102302331100"></a>

#### `service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route.incoming_port.port` property

Type: `"number"`. Optional.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-0213023023113131-3020210300100010-0122000130021212-1113031120321223-3033021131110001-3303202003130122-1031313333300312-1233211223220000"></a>

<a id="canonical-3212201221330123-2233200200220321-1322320000123220-0301323230101312-0313212300130212-1001203332213303-3223212112201113-3022001133332300"></a>

#### `service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route.incoming_port.port_ranges` property

Type: `"string"`. Optional.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-1100123212233303-1101022230112200-1321101001212020-2330032102000112-0222001033210103-1132102002033200-3023130003310230-0112321320301123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route.incoming_port.no_port_match` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-2121230312030330-3010113231323001-1233013032320322-1032203002210311-2022021002113232-1123312032112001-0012103213230102-1112311001110020)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-006.md#canonical-0133221101333210-0222103021001232-3212010032013103-3331120032012221-1012203003232110-1233102323001233-2233321121201101-0122313122210200)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-006.md#canonical-2311001323203002-2023333020201320-1330320322101233-1033123102030111-2013210002333120-2130122003123103-1133301323123220-2302130221012102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes](resources--workload--reference--group-008.md#canonical-3011211302110102-1110003330130012-2330003231002123-1333212203033211-1212101000211203-3333313212202213-1122232332203332-0220110332323232)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes](resources--workload--reference--group-008.md#canonical-3101223212120012-3332300223001100-0333220310012333-1031332113331120-2210213133122022-1333032223122231-2030231210333100-1103221131133311)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route](resources--workload--reference--group-008.md#canonical-0003020102311303-0233233032322332-3002231010033312-1313303221323312-0103013131011120-3102111011102132-1313020300032233-1111311233230113)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route.incoming_port](resources--workload--reference--group-008.md#canonical-1020323201110112-2003203023313010-0130110100322020-2100332301223310-0121230331313023-1013011231112133-2010303211002033-3100032033120011)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route.incoming_port.no_port_match

<a id="canonical-3311220033101312-1121013033131013-2023131000211233-2201120223010233-1231220211121201-2101233232311201-2332330132201121-0013203213030112"></a>

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
no_port_match = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3233030021033313-3031030311221020-0200331022102113-1313122110031101-2210311230313221-2020211211322311-2201003031033002-3320231130112311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route.path` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-2121230312030330-3010113231323001-1233013032320322-1032203002210311-2022021002113232-1123312032112001-0012103213230102-1112311001110020)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-006.md#canonical-0133221101333210-0222103021001232-3212010032013103-3331120032012221-1012203003232110-1233102323001233-2233321121201101-0122313122210200)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-006.md#canonical-2311001323203002-2023333020201320-1330320322101233-1033123102030111-2013210002333120-2130122003123103-1133301323123220-2302130221012102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes](resources--workload--reference--group-008.md#canonical-3011211302110102-1110003330130012-2330003231002123-1333212203033211-1212101000211203-3333313212202213-1122232332203332-0220110332323232)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes](resources--workload--reference--group-008.md#canonical-3101223212120012-3332300223001100-0333220310012333-1031332113331120-2210213133122022-1333032223122231-2030231210333100-1103221131133311)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route](resources--workload--reference--group-008.md#canonical-0003020102311303-0233233032322332-3002231010033312-1313303221323312-0103013131011120-3102111011102132-1313020300032233-1111311233230113)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route.path

<a id="canonical-0120033123301021-3133200221012111-0321320333000033-1013332123032221-1113311120100233-1121220111011013-2313112202122311-1332121210312202"></a>

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

<a id="canonical-1330230210103312-1322032023030231-1202033120320302-3123321303211012-3201333230233113-1032003321320333-3000321112333020-2001103021013300"></a>

### Direct properties for `service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route.path`

<a id="canonical-2200000120222320-1310132330012212-0013312100203003-3023131023213101-1230210130123123-1020303113020301-1122320023233333-0213110300302203"></a>

#### `service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route.path.path` property

Type: `"string"`. Optional.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-0022131321030221-1232010312212320-0100133232220113-1321111223130103-2232213313110202-3002021202021223-1101331303023120-1323121000212131"></a>

<a id="canonical-1212123300002001-3121002212322030-3031200333302133-2231310101000202-0021020300020013-1031221002130101-2232010132000111-0103030031232211"></a>

#### `service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route.path.prefix` property

Type: `"string"`. Optional.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-1120221200003332-0202310201301203-3122210002012203-0320113213212030-2130231211000232-2230121020003001-1202300222123010-0211200033323233"></a>

<a id="canonical-3131223022330200-1221300222203320-2322330031220210-2221202310122120-0120033321003010-1213133022301113-3013011021202211-0233200222013220"></a>

#### `service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route.path.regex` property

Type: `"string"`. Optional.

Exclusive with \[path prefix\] Regular expression of path match (e.g. The value .\* will match on
all paths).

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-1013022013311110-1211131313223021-1020032322323123-0111133101310103-2132011220312120-1323130111012310-3301103301322333-3232211311000323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-2121230312030330-3010113231323001-1233013032320322-1032203002210311-2022021002113232-1123312032112001-0012103213230102-1112311001110020)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-006.md#canonical-0133221101333210-0222103021001232-3212010032013103-3331120032012221-1012203003232110-1233102323001233-2233321121201101-0122313122210200)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-006.md#canonical-2311001323203002-2023333020201320-1330320322101233-1033123102030111-2013210002333120-2130122003123103-1133301323123220-2302130221012102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes](resources--workload--reference--group-008.md#canonical-3011211302110102-1110003330130012-2330003231002123-1333212203033211-1212101000211203-3333313212202213-1122232332203332-0220110332323232)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes](resources--workload--reference--group-008.md#canonical-3101223212120012-3332300223001100-0333220310012333-1031332113331120-2210213133122022-1333032223122231-2030231210333100-1103221131133311)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route](resources--workload--reference--group-008.md#canonical-0003020102311303-0233233032322332-3002231010033312-1313303221323312-0103013131011120-3102111011102132-1313020300032233-1111311233230113)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect

<a id="canonical-2113331313112203-0021303132111332-0203321210123120-2333100221110303-2312000210302001-3312133100130211-1113023103323003-1101231023302023"></a>

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

<a id="canonical-3211012302023200-1000003221221022-1323210223110302-1330021313133330-2101332131321120-3013203131230121-3131220132021112-0022302220123321"></a>

### Direct properties for `service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect`

<a id="canonical-3103123231213210-0310030331101113-1001212303223023-3020120122010002-2113130022110210-3203031003120003-3331223123120312-3202302133333210"></a>

#### `service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect.host_redirect` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-3320130001001313-2003102213210111-3310312333220223-1130322220331132-0013120023123311-1023102123220030-2102222031300333-2132330202230001"></a>

<a id="canonical-0102102310033320-2200332023323322-0112301331120030-0213231230302303-0133021113311201-1013032323200320-1310021111313200-3212221000023122"></a>

#### `service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect.path_redirect` property

Type: `"string"`. Optional.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-2313333213221213-3131303211331333-0102311322022132-2201131012032031-1231103331101113-3223312003123332-0200121230011310-2031021133211120"></a>

<a id="canonical-2010330220200220-3320321211123003-3100201120231031-2323000120321030-3232103201310332-3032131201203333-3322000323001100-3033203021021022"></a>

#### `service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect.prefix_rewrite` property

Type: `"string"`. Optional.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-2100221112111033-1003203221302303-1022303301201101-0123200012310232-1230020020313232-0111133332212020-2221022331022010-2121201123222330"></a>

<a id="canonical-3233211020030032-3222123023110302-0203200202020213-2111220102210102-1033303110123233-1001132033010013-2210322030022033-0010223230011313"></a>

#### `service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect.proto_redirect` property

Type: `"string"`. Optional.

\[Enum: incoming-proto|http|https\] Swap protocol part of incoming URL in redirect URL The protocol
can be swapped with either HTTP or HTTPS When incoming-proto option is specified, swapping of
protocol is not done. Possible values are \`incoming-proto\`, \`http\`, \`https\`.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

- [remove_all_params](resources--workload--reference--group-008.md#canonical-3030001233323101-1213321031223203-2130331003313333-0232203302320112-3010210113021201-3102332130322022-3122221213011101-1030220123101010): complete subsection reference.

<a id="canonical-1100210330113232-3010232321203301-2323113113033220-1131000131132311-0030102202313011-3023323003223211-2310203033120000-2201103232220231"></a>

<a id="canonical-0000032111220311-2310121330212303-0113122233032123-3322011112202130-0110120301302100-3231133310212202-1010003230011133-0023030303213010"></a>

#### `service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect.replace_params` property

Type: `"string"`. Optional.

Exclusive with \[remove\_all\_params retain\_all\_params\].

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-3222332031023323-0201031112230132-2302112300233313-1300100032101102-0103002001002111-2020020113121110-1302022233002331-3132031000302310"></a>

<a id="canonical-0021232132232310-2310131322230033-2221231103213120-1022200012022000-0123321212322302-1130102301323111-0331003133012331-2230113022211332"></a>

#### `service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect.response_code` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

- [retain_all_params](resources--workload--reference--group-008.md#canonical-3322210232313033-0233122011002301-0322211331113203-1122233332212220-2200101030301312-1000312321303020-1331011230331221-0133023110033110): complete subsection reference.

<a id="canonical-3030001233323101-1213321031223203-2130331003313333-0232203302320112-3010210113021201-3102332130322022-3122221213011101-1030220123101010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect.remove_all_params` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-2121230312030330-3010113231323001-1233013032320322-1032203002210311-2022021002113232-1123312032112001-0012103213230102-1112311001110020)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-006.md#canonical-0133221101333210-0222103021001232-3212010032013103-3331120032012221-1012203003232110-1233102323001233-2233321121201101-0122313122210200)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-006.md#canonical-2311001323203002-2023333020201320-1330320322101233-1033123102030111-2013210002333120-2130122003123103-1133301323123220-2302130221012102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes](resources--workload--reference--group-008.md#canonical-3011211302110102-1110003330130012-2330003231002123-1333212203033211-1212101000211203-3333313212202213-1122232332203332-0220110332323232)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes](resources--workload--reference--group-008.md#canonical-3101223212120012-3332300223001100-0333220310012333-1031332113331120-2210213133122022-1333032223122231-2030231210333100-1103221131133311)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route](resources--workload--reference--group-008.md#canonical-0003020102311303-0233233032322332-3002231010033312-1313303221323312-0103013131011120-3102111011102132-1313020300032233-1111311233230113)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect](resources--workload--reference--group-008.md#canonical-1013022013311110-1211131313223021-1020032322323123-0111133101310103-2132011220312120-1323130111012310-3301103301322333-3232211311000323)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect.remove_all_params

<a id="canonical-3312213121211102-2010303012122113-2002103012311003-0330221202301013-3221322200000333-0112000212313210-0220113131012032-2201233111102133"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for remove all params.

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
remove_all_params = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3322210232313033-0233122011002301-0322211331113203-1122233332212220-2200101030301312-1000312321303020-1331011230331221-0133023110033110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect.retain_all_params` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-2121230312030330-3010113231323001-1233013032320322-1032203002210311-2022021002113232-1123312032112001-0012103213230102-1112311001110020)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-006.md#canonical-0133221101333210-0222103021001232-3212010032013103-3331120032012221-1012203003232110-1233102323001233-2233321121201101-0122313122210200)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-006.md#canonical-2311001323203002-2023333020201320-1330320322101233-1033123102030111-2013210002333120-2130122003123103-1133301323123220-2302130221012102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes](resources--workload--reference--group-008.md#canonical-3011211302110102-1110003330130012-2330003231002123-1333212203033211-1212101000211203-3333313212202213-1122232332203332-0220110332323232)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes](resources--workload--reference--group-008.md#canonical-3101223212120012-3332300223001100-0333220310012333-1031332113331120-2210213133122022-1333032223122231-2030231210333100-1103221131133311)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route](resources--workload--reference--group-008.md#canonical-0003020102311303-0233233032322332-3002231010033312-1313303221323312-0103013131011120-3102111011102132-1313020300032233-1111311233230113)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect](resources--workload--reference--group-008.md#canonical-1013022013311110-1211131313223021-1020032322323123-0111133101310103-2132011220312120-1323130111012310-3301103301322333-3232211311000323)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect.retain_all_params

<a id="canonical-1033300012131220-2212333233000232-1323022303323222-3302213222300010-2312212123220220-0010032002033102-3321130203312311-0230133002213220"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for retain all params.

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
retain_all_params = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2332230303301330-0120223130133121-2131103033210110-0000331031322110-2102032123012022-0000121011033231-1222320000321202-3323210133011002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.simple_route` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-2121230312030330-3010113231323001-1233013032320322-1032203002210311-2022021002113232-1123312032112001-0012103213230102-1112311001110020)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-006.md#canonical-0133221101333210-0222103021001232-3212010032013103-3331120032012221-1012203003232110-1233102323001233-2233321121201101-0122313122210200)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-006.md#canonical-2311001323203002-2023333020201320-1330320322101233-1033123102030111-2013210002333120-2130122003123103-1133301323123220-2302130221012102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes](resources--workload--reference--group-008.md#canonical-3011211302110102-1110003330130012-2330003231002123-1333212203033211-1212101000211203-3333313212202213-1122232332203332-0220110332323232)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes](resources--workload--reference--group-008.md#canonical-3101223212120012-3332300223001100-0333220310012333-1031332113331120-2210213133122022-1333032223122231-2030231210333100-1103221131133311)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.simple_route

<a id="canonical-0023301100111202-2010221211111300-1003211011133233-0220201200002312-3030033223001302-0230022332133220-2303121213010112-2230022002133311"></a>

Type: `"object"`. single nested block, Optional.

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

<a id="canonical-3233202133222313-0320003011311111-0201130303022111-3112213301312133-3103211010002030-0221200123111133-3222101332030210-3123023331202032"></a>

### Direct properties for `service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.simple_route`

- [auto_host_rewrite](resources--workload--reference--group-008.md#canonical-0112230132313323-2122103132021131-1012002001203012-1000103302310330-0032032123212331-0123013323310310-2230212223000111-0220101331331003): complete subsection reference.

- [disable_host_rewrite](resources--workload--reference--group-008.md#canonical-1111112032212222-1331322310310203-3123130031332003-3000123221101021-0102220030220211-3130102131232020-3010130002220003-2212212321323022): complete subsection reference.

<a id="canonical-2213103011033112-1101103313232020-1102322303013111-0022213231232323-1122101321303333-2101133033021000-0222223213313331-2220330213102122"></a>

<a id="canonical-1222030011031121-1330300131103121-0221123001322330-1230303003022011-0232212313213302-1020321223013011-2303011331132012-2110210111202313"></a>

#### `service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.simple_route.host_rewrite` property

Type: `"string"`. Optional.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-2202221103103213-2103013011301221-2213230211133121-3111102313030313-3021002102021320-2023100233333023-2132231232211010-0133020122001101"></a>

<a id="canonical-1022000103212233-0220012120320133-3011321231011022-3210223002200122-3002023200023133-3100202321310312-0103202032100333-3002132010031301"></a>

#### `service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.simple_route.http_method` property

Type: `"string"`. Optional.

\[Enum: ANY|GET|HEAD|POST|PUT|DELETE|CONNECT|OPTIONS|TRACE|PATCH|COPY\] Specifies the HTTP method
used to access a resource. Any HTTP Method. Possible values are \`ANY\`, \`GET\`, \`HEAD\`,
\`POST\`, \`PUT\`, \`DELETE\`, \`CONNECT\`, \`OPTIONS\`, \`TRACE\`, \`PATCH\`, \`COPY\`. Defaults to
\`ANY\`.

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

- [path](resources--workload--reference--group-008.md#canonical-3221021331221300-2122202121211303-0333120310113302-2003110323312131-1320212313100121-1101212002232103-1001323321022113-3210123212220031): complete subsection reference.

<a id="canonical-0112230132313323-2122103132021131-1012002001203012-1000103302310330-0032032123212331-0123013323310310-2230212223000111-0220101331331003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.simple_route.auto_host_rewrite` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-2121230312030330-3010113231323001-1233013032320322-1032203002210311-2022021002113232-1123312032112001-0012103213230102-1112311001110020)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-006.md#canonical-0133221101333210-0222103021001232-3212010032013103-3331120032012221-1012203003232110-1233102323001233-2233321121201101-0122313122210200)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-006.md#canonical-2311001323203002-2023333020201320-1330320322101233-1033123102030111-2013210002333120-2130122003123103-1133301323123220-2302130221012102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes](resources--workload--reference--group-008.md#canonical-3011211302110102-1110003330130012-2330003231002123-1333212203033211-1212101000211203-3333313212202213-1122232332203332-0220110332323232)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes](resources--workload--reference--group-008.md#canonical-3101223212120012-3332300223001100-0333220310012333-1031332113331120-2210213133122022-1333032223122231-2030231210333100-1103221131133311)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.simple_route](resources--workload--reference--group-008.md#canonical-2332230303301330-0120223130133121-2131103033210110-0000331031322110-2102032123012022-0000121011033231-1222320000321202-3323210133011002)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.simple_route.auto_host_rewrite

<a id="canonical-3020001302123220-3032021000301010-0312301322220200-3332312312310232-2212301110110320-0333302321001121-2030223301130312-3223212123122120"></a>

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
auto_host_rewrite = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1111112032212222-1331322310310203-3123130031332003-3000123221101021-0102220030220211-3130102131232020-3010130002220003-2212212321323022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.simple_route.disable_host_rewrite` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-2121230312030330-3010113231323001-1233013032320322-1032203002210311-2022021002113232-1123312032112001-0012103213230102-1112311001110020)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-006.md#canonical-0133221101333210-0222103021001232-3212010032013103-3331120032012221-1012203003232110-1233102323001233-2233321121201101-0122313122210200)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-006.md#canonical-2311001323203002-2023333020201320-1330320322101233-1033123102030111-2013210002333120-2130122003123103-1133301323123220-2302130221012102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes](resources--workload--reference--group-008.md#canonical-3011211302110102-1110003330130012-2330003231002123-1333212203033211-1212101000211203-3333313212202213-1122232332203332-0220110332323232)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes](resources--workload--reference--group-008.md#canonical-3101223212120012-3332300223001100-0333220310012333-1031332113331120-2210213133122022-1333032223122231-2030231210333100-1103221131133311)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.simple_route](resources--workload--reference--group-008.md#canonical-2332230303301330-0120223130133121-2131103033210110-0000331031322110-2102032123012022-0000121011033231-1222320000321202-3323210133011002)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.simple_route.disable_host_rewrite

<a id="canonical-3013131201302111-3200103322030120-3211031323212022-2023230122322203-0132210130102222-3122311100232130-0100223111333101-3002100012031012"></a>

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
disable_host_rewrite = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3221021331221300-2122202121211303-0333120310113302-2003110323312131-1320212313100121-1101212002232103-1001323321022113-3210123212220031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.simple_route.path` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-2121230312030330-3010113231323001-1233013032320322-1032203002210311-2022021002113232-1123312032112001-0012103213230102-1112311001110020)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-006.md#canonical-0133221101333210-0222103021001232-3212010032013103-3331120032012221-1012203003232110-1233102323001233-2233321121201101-0122313122210200)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-006.md#canonical-2311001323203002-2023333020201320-1330320322101233-1033123102030111-2013210002333120-2130122003123103-1133301323123220-2302130221012102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes](resources--workload--reference--group-008.md#canonical-3011211302110102-1110003330130012-2330003231002123-1333212203033211-1212101000211203-3333313212202213-1122232332203332-0220110332323232)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes](resources--workload--reference--group-008.md#canonical-3101223212120012-3332300223001100-0333220310012333-1031332113331120-2210213133122022-1333032223122231-2030231210333100-1103221131133311)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.simple_route](resources--workload--reference--group-008.md#canonical-2332230303301330-0120223130133121-2131103033210110-0000331031322110-2102032123012022-0000121011033231-1222320000321202-3323210133011002)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.simple_route.path

<a id="canonical-2233001333001332-2303111302302332-0202302322123220-1130113312221013-1332110210121221-3220120321013001-2321331130311303-3113023221202222"></a>

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

<a id="canonical-2232330312223222-2213113302122220-2103111133032330-3331102022332003-3023203331021123-1133000103323030-2123230211331013-3113022313033223"></a>

### Direct properties for `service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.simple_route.path`

<a id="canonical-0322133032200232-2132032232022301-3002330223220223-2230033003112322-3232112213331330-0321002121010012-3010130310101133-3330032230110120"></a>

#### `service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.simple_route.path.path` property

Type: `"string"`. Optional.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-3033133211102031-0210002131003012-1001020331002313-1122212103103133-1230012021113022-0330122332123010-3113331230000132-0232000002301211"></a>

<a id="canonical-1321313232101020-0031333013231121-1333102202122232-3101221012131022-3112311223023310-3221323032200110-0200000310122233-2312203020111102"></a>

#### `service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.simple_route.path.prefix` property

Type: `"string"`. Optional.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-0330323132022321-3203032310233020-0003231130130101-3002120102210110-1203123113113213-3220212033011012-0113210033132000-3121231311030012"></a>

<a id="canonical-2210030130330012-3112231310121032-2123030121102123-2311201333020032-0301003130332123-2101103113010123-1002200300220302-2102032001002103"></a>

#### `service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.simple_route.path.regex` property

Type: `"string"`. Optional.

Exclusive with \[path prefix\] Regular expression of path match (e.g. The value .\* will match on
all paths).

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-2011231233103202-1130133301231221-3112112222130011-0133100320022222-3002023210301231-3323320331313221-3230300303311111-0302223303200123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_custom.ports.port` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-2121230312030330-3010113231323001-1233013032320322-1032203002210311-2022021002113232-1123312032112001-0012103213230102-1112311001110020)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-006.md#canonical-0133221101333210-0222103021001232-3212010032013103-3331120032012221-1012203003232110-1233102323001233-2233321121201101-0122313122210200)
- service.advertise_options.advertise_custom.ports.port

<a id="canonical-0303220310031220-2231211003010121-2130101032232221-3001101213022122-0330311011320112-2211322120222030-2133311323320300-1103211130031312"></a>

Type: `"object"`. single nested block, Optional.

Port. Port of the workload.

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
port {
  # Configure direct properties listed below.
}
```

<a id="canonical-2000312131021111-1322002100000322-3001222320013321-0010203332221311-1033201021322303-0201313210131122-2133003110330332-1313131300231110"></a>

### Direct properties for `service.advertise_options.advertise_custom.ports.port`

- [info](resources--workload--reference--group-008.md#canonical-1122232230031120-0022203213323110-0233011123013101-3033030322231311-0002300310301212-3213332001002102-0103000132313210-3031331333211321): complete subsection reference.

<a id="canonical-0023000021233322-3210122310230003-3133020310012111-0330331320122320-3221113220030013-3021322010001231-2202303311022010-0322231210130020"></a>

<a id="canonical-3033303032032110-2002330332133132-2020011122101331-1320002101313011-1301030222100212-2301011031221323-3310011230022130-3313211233330230"></a>

#### `service.advertise_options.advertise_custom.ports.port.name` property

Type: `"string"`. Optional.

Name. Name of the Port.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
  stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z]([-a-z0-9]*[a-z0-9])?$`),
    ""),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "naming",
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
      "source": "inferred",
      "validatedAt": "2026-10-07T18:46:18+00:00"
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
    "ves.io.schema.rules.string.iana_svc_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.iana_svc_name": "true"
  }
}
```

<a id="canonical-1122232230031120-0022203213323110-0233011123013101-3033030322231311-0002300310301212-3213332001002102-0103000132313210-3031331333211321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_custom.ports.port.info` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-2121230312030330-3010113231323001-1233013032320322-1032203002210311-2022021002113232-1123312032112001-0012103213230102-1112311001110020)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-006.md#canonical-0133221101333210-0222103021001232-3212010032013103-3331120032012221-1012203003232110-1233102323001233-2233321121201101-0122313122210200)
- [service.advertise_options.advertise_custom.ports.port](resources--workload--reference--group-008.md#canonical-2011231233103202-1130133301231221-3112112222130011-0133100320022222-3002023210301231-3323320331313221-3230300303311111-0302223303200123)
- service.advertise_options.advertise_custom.ports.port.info

<a id="canonical-0213020202001320-1303321131021231-0321003301313021-1203023131303220-2300321033130002-0332110330301121-0201213033203212-1100001020102031"></a>

Type: `"object"`. single nested block, Optional.

Port Information. Port information.

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

<a id="canonical-3113332212100112-2202131001232123-1111100030312232-0030012321033322-0031010202110332-2033123322130221-0012023323232032-3030322332021120"></a>

### Direct properties for `service.advertise_options.advertise_custom.ports.port.info`

<a id="canonical-3121333321222000-2003212003020111-1030002011031030-3102203233310312-2211330113133222-1230013222032000-2201011323220120-0031331310323333"></a>

#### `service.advertise_options.advertise_custom.ports.port.info.port` property

Type: `"number"`. Optional.

Port. Port the workload can be reached on.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-2130113332133012-3321112320130113-2221132013233301-1223132011211110-3311213231113003-2022331022001012-3213320010101102-3110111032013302"></a>

<a id="canonical-2313033323023301-0001202030113222-3133111022302303-2011003220013202-0011202322001022-2123330303223032-2332022302311012-0322303333033020"></a>

#### `service.advertise_options.advertise_custom.ports.port.info.protocol` property

Type: `"string"`. Optional.

\[Enum: PROTOCOL\_TCP|PROTOCOL\_HTTP|PROTOCOL\_HTTP2|PROTOCOL\_TLS\_WITH\_SNI|PROTOCOL\_UDP\] Type
of protocol - PROTOCOL\_TCP: TCP TCP - PROTOCOL\_HTTP: HTTP HTTP - PROTOCOL\_HTTP2: HTTP2 HTTP2 -
PROTOCOL\_TLS\_WITH\_SNI: TLS with SNI TLS with SNI - PROTOCOL\_UDP: UDP UDP. Possible values are
\`PROTOCOL\_TCP\`, \`PROTOCOL\_HTTP\`, \`PROTOCOL\_HTTP2\`, \`PROTOCOL\_TLS\_WITH\_SNI\`,
\`PROTOCOL\_UDP\`. Defaults to \`PROTOCOL\_TCP\`.

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

- [same_as_port](resources--workload--reference--group-008.md#canonical-1110030120221301-2033120210021001-0101031112323320-0213132221023103-0223022301323322-1022213103310103-2011031302113000-1011103331323132): complete subsection reference.

<a id="canonical-1213323310203333-3101130010212232-1133311103130032-3022003211030311-1013000230233330-0313121111120121-0021302223010123-0131032002020310"></a>

<a id="canonical-0311210300001112-0221100103100020-1220232032013100-2221203333321000-3312220011121033-0030232022031123-3021133203031021-0130011133231100"></a>

#### `service.advertise_options.advertise_custom.ports.port.info.target_port` property

Type: `"number"`. Optional.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-1110030120221301-2033120210021001-0101031112323320-0213132221023103-0223022301323322-1022213103310103-2011031302113000-1011103331323132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_custom.ports.port.info.same_as_port` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-2121230312030330-3010113231323001-1233013032320322-1032203002210311-2022021002113232-1123312032112001-0012103213230102-1112311001110020)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-006.md#canonical-0133221101333210-0222103021001232-3212010032013103-3331120032012221-1012203003232110-1233102323001233-2233321121201101-0122313122210200)
- [service.advertise_options.advertise_custom.ports.port](resources--workload--reference--group-008.md#canonical-2011231233103202-1130133301231221-3112112222130011-0133100320022222-3002023210301231-3323320331313221-3230300303311111-0302223303200123)
- [service.advertise_options.advertise_custom.ports.port.info](resources--workload--reference--group-008.md#canonical-1122232230031120-0022203213323110-0233011123013101-3033030322231311-0002300310301212-3213332001002102-0103000132313210-3031331333211321)
- service.advertise_options.advertise_custom.ports.port.info.same_as_port

<a id="canonical-1330133210030312-0100112101001102-0012030120131323-2213310121331210-3332013323130330-1232301301130103-1000010033011211-1131010012013133"></a>

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
same_as_port = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2333000323211321-3320220213111111-3000113231121121-2211220223011110-3303030010002012-3322030301222113-2111223031202010-0320112332022000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_custom.ports.tcp_loadbalancer` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-2121230312030330-3010113231323001-1233013032320322-1032203002210311-2022021002113232-1123312032112001-0012103213230102-1112311001110020)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-006.md#canonical-0133221101333210-0222103021001232-3212010032013103-3331120032012221-1012203003232110-1233102323001233-2233321121201101-0122313122210200)
- service.advertise_options.advertise_custom.ports.tcp_loadbalancer

<a id="canonical-2100023101202221-1121112011012321-1010332030111311-1232012122332331-2011223133312332-3311201231213313-0000332123100013-3301321001033101"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for tcp loadbalancer.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
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

<a id="canonical-0120222130300023-3233330321032112-0333212200202303-0002310313302132-3223312022302230-0022010310231303-1003222211320013-2311002003030022"></a>

### Direct properties for `service.advertise_options.advertise_custom.ports.tcp_loadbalancer`

<a id="canonical-3300103310101212-3023222200221303-0010222232310312-1322201021232112-2110012302301311-3113012120331232-1231122021233310-2333330130101200"></a>

#### `service.advertise_options.advertise_custom.ports.tcp_loadbalancer.domains` property

Type: `["list", "string"]`. Optional.

List of additional domains (host/authority header) that will be matched to this loadbalancer.
Domains are also used for SNI matching if the is true Domains also indicate the list of names for
which DNS resolution will be done by VER.

Additional upstream details:

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-1300131122302020-0330022100313311-2012212200121003-1001101012200320-1333123111120300-0020003323023132-0102203001103331-1113211132013111"></a>

<a id="canonical-3021032321001211-2001333313320310-2122202330333031-1203121002002122-2000100102012303-2011000232010101-2332313021003231-2312320313203102"></a>

#### `service.advertise_options.advertise_custom.ports.tcp_loadbalancer.with_sni` property

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

<a id="canonical-0330023301131021-0130211002331300-3002122032320301-2112231133101233-3303101211030122-3300030030023300-1213020320030020-0213001312211213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_in_cluster` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- service.advertise_options.advertise_in_cluster

<a id="canonical-3303010210113203-0000123133102021-2220101212333303-2022210232202312-1121210030320001-2122031211022321-0021132013231232-2231133202303313"></a>

Type: `"object"`. single nested block, Optional.

Advertise the workload locally in-cluster.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("multi_ports",
    "port")}
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
  "x-ves-oneof-field-port_choice": "[\"multi_ports\",\"port\"]"
}
```

Terraform syntax:

```terraform
advertise_in_cluster {
  # Configure direct properties listed below.
}
```

<a id="canonical-0330220031113300-0003030303312010-3230010030031303-2330233220133033-2012210020230323-3020122113211001-3033031213012202-0002313032203301"></a>

### Direct properties for `service.advertise_options.advertise_in_cluster`

- [multi_ports](resources--workload--reference--group-008.md#canonical-3233102030323303-0010001111332222-2221330300220032-2322233002300322-1030033321333231-1200300001103302-1331010302322302-2231113103111232): complete subsection reference.

- [port](resources--workload--reference--group-008.md#canonical-2032101121303011-1103222233000330-0203232311001010-2322120323102000-3202033332221123-2310113223210032-0131222201221033-2112013321213130): complete subsection reference.

<a id="canonical-3233102030323303-0010001111332222-2221330300220032-2322233002300322-1030033321333231-1200300001103302-1331010302322302-2231113103111232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_in_cluster.multi_ports` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_in_cluster](resources--workload--reference--group-008.md#canonical-0330023301131021-0130211002331300-3002122032320301-2112231133101233-3303101211030122-3300030030023300-1213020320030020-0213001312211213)
- service.advertise_options.advertise_in_cluster.multi_ports

<a id="canonical-0010032311313302-2100112001110300-0332012333311013-1301111313330310-3033232010320211-1312310313022101-1123110330123221-0022210010330200"></a>

Type: `"object"`. single nested block, Optional.

Multiple Ports. Multiple ports.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("ports")}
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
multi_ports {
  # Configure direct properties listed below.
}
```

<a id="canonical-2121202131001323-2330111213331230-0212203132030221-2132211323022111-3310123332033313-0010331001230213-1122020300200303-2032130221210202"></a>

### Direct properties for `service.advertise_options.advertise_in_cluster.multi_ports`

- [ports](resources--workload--reference--group-008.md#canonical-0302230202101333-3033110013223110-0210311023223222-0203030332102133-2201333301103120-1000302302100130-2222303013302022-1022120223133312): complete subsection reference.

<a id="canonical-0302230202101333-3033110013223110-0210311023223222-0203030332102133-2201333301103120-1000302302100130-2222303013302022-1022120223133312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_in_cluster.multi_ports.ports` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_in_cluster](resources--workload--reference--group-008.md#canonical-0330023301131021-0130211002331300-3002122032320301-2112231133101233-3303101211030122-3300030030023300-1213020320030020-0213001312211213)
- [service.advertise_options.advertise_in_cluster.multi_ports](resources--workload--reference--group-008.md#canonical-3233102030323303-0010001111332222-2221330300220032-2322233002300322-1030033321333231-1200300001103302-1331010302322302-2231113103111232)
- service.advertise_options.advertise_in_cluster.multi_ports.ports

<a id="canonical-3310220231101131-2122202103203132-2032032312323112-0111111123010110-0210101110331033-3311220102322010-2300011333310100-3021201133033222"></a>

Type: `"object"`. list nested block, Optional.

Ports. Ports to advertise.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.RequiredListObjectAttributes("name")}
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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "16"
  }
}
```

Terraform syntax:

```terraform
ports {
  # Configure direct properties listed below.
}
```

<a id="canonical-3101111113211202-3212031000032220-2230230213230302-1312223130030012-2231222033032331-1101133300332102-2330100330301210-2223123203000103"></a>

### Direct properties for `service.advertise_options.advertise_in_cluster.multi_ports.ports`

- [info](resources--workload--reference--group-008.md#canonical-1032330310201200-3101102022121101-0202311121001130-1310203222000011-1113321122130331-0113223320013030-1012012320323211-2111201022201331): complete subsection reference.

<a id="canonical-0211032121020212-0320101100100331-2030002032011211-2133120033210201-1201202203121000-2303322233330300-3011300100221030-1302003300331112"></a>

<a id="canonical-2033313000001030-3130111022233102-3310333033111030-1130001320101323-3201202122330332-0312022313103002-1302023022033322-0321000210330333"></a>

#### `service.advertise_options.advertise_in_cluster.multi_ports.ports.name` property

Type: `"string"`. Optional.

Name. Name of the Port.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
  stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z]([-a-z0-9]*[a-z0-9])?$`),
    ""),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "naming",
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
      "source": "inferred",
      "validatedAt": "2026-10-07T18:46:18+00:00"
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
    "ves.io.schema.rules.string.iana_svc_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.iana_svc_name": "true"
  }
}
```

<a id="canonical-1032330310201200-3101102022121101-0202311121001130-1310203222000011-1113321122130331-0113223320013030-1012012320323211-2111201022201331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_in_cluster.multi_ports.ports.info` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_in_cluster](resources--workload--reference--group-008.md#canonical-0330023301131021-0130211002331300-3002122032320301-2112231133101233-3303101211030122-3300030030023300-1213020320030020-0213001312211213)
- [service.advertise_options.advertise_in_cluster.multi_ports](resources--workload--reference--group-008.md#canonical-3233102030323303-0010001111332222-2221330300220032-2322233002300322-1030033321333231-1200300001103302-1331010302322302-2231113103111232)
- [service.advertise_options.advertise_in_cluster.multi_ports.ports](resources--workload--reference--group-008.md#canonical-0302230202101333-3033110013223110-0210311023223222-0203030332102133-2201333301103120-1000302302100130-2222303013302022-1022120223133312)
- service.advertise_options.advertise_in_cluster.multi_ports.ports.info

<a id="canonical-2212212020221220-2223322222013330-3113113200233311-3132033221200000-1010220031230330-2100032331322220-0102331032210211-0032331101303110"></a>

Type: `"object"`. single nested block, Optional.

Port Information. Port information.

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

<a id="canonical-2310030110231120-1232203023211031-3310322122130032-3112203112130110-0001111211203312-2122011230031300-3013313101000023-2102213112030202"></a>

### Direct properties for `service.advertise_options.advertise_in_cluster.multi_ports.ports.info`

<a id="canonical-3121101201110211-1101213022120113-0211330330113021-0100031313112122-1323300222332221-1331102122010211-3001012123210231-3303321312013331"></a>

#### `service.advertise_options.advertise_in_cluster.multi_ports.ports.info.port` property

Type: `"number"`. Optional.

Port. Port the workload can be reached on.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-3203030333002131-3013112331220123-1033102222310020-2110213032002311-2312312312003330-1222331111031320-2010021123002110-1000203203323021"></a>

<a id="canonical-0110200323111113-2113210211213103-2331313020320230-0222302132313120-0011030112131011-2201232203220200-0302222230003311-1323210213013110"></a>

#### `service.advertise_options.advertise_in_cluster.multi_ports.ports.info.protocol` property

Type: `"string"`. Optional.

\[Enum: PROTOCOL\_TCP|PROTOCOL\_HTTP|PROTOCOL\_HTTP2|PROTOCOL\_TLS\_WITH\_SNI|PROTOCOL\_UDP\] Type
of protocol - PROTOCOL\_TCP: TCP TCP - PROTOCOL\_HTTP: HTTP HTTP - PROTOCOL\_HTTP2: HTTP2 HTTP2 -
PROTOCOL\_TLS\_WITH\_SNI: TLS with SNI TLS with SNI - PROTOCOL\_UDP: UDP UDP. Possible values are
\`PROTOCOL\_TCP\`, \`PROTOCOL\_HTTP\`, \`PROTOCOL\_HTTP2\`, \`PROTOCOL\_TLS\_WITH\_SNI\`,
\`PROTOCOL\_UDP\`. Defaults to \`PROTOCOL\_TCP\`.

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

- [same_as_port](resources--workload--reference--group-008.md#canonical-3203022201002002-3230201120302112-2201031221333112-2131003320203232-0313201303030011-2101023111002031-2322011130001010-0231013131320131): complete subsection reference.

<a id="canonical-1032202030013211-0202010310302001-1323333032331310-0323120012210031-2012210221013002-3011132221003323-0110111333003113-1202131203000022"></a>

<a id="canonical-1220030212122020-0211033113221113-1100211031233331-0332230221210223-1201132023330321-2020330022233032-3313321321010203-1211200122300001"></a>

#### `service.advertise_options.advertise_in_cluster.multi_ports.ports.info.target_port` property

Type: `"number"`. Optional.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-3203022201002002-3230201120302112-2201031221333112-2131003320203232-0313201303030011-2101023111002031-2322011130001010-0231013131320131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_in_cluster.multi_ports.ports.info.same_as_port` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_in_cluster](resources--workload--reference--group-008.md#canonical-0330023301131021-0130211002331300-3002122032320301-2112231133101233-3303101211030122-3300030030023300-1213020320030020-0213001312211213)
- [service.advertise_options.advertise_in_cluster.multi_ports](resources--workload--reference--group-008.md#canonical-3233102030323303-0010001111332222-2221330300220032-2322233002300322-1030033321333231-1200300001103302-1331010302322302-2231113103111232)
- [service.advertise_options.advertise_in_cluster.multi_ports.ports](resources--workload--reference--group-008.md#canonical-0302230202101333-3033110013223110-0210311023223222-0203030332102133-2201333301103120-1000302302100130-2222303013302022-1022120223133312)
- [service.advertise_options.advertise_in_cluster.multi_ports.ports.info](resources--workload--reference--group-008.md#canonical-1032330310201200-3101102022121101-0202311121001130-1310203222000011-1113321122130331-0113223320013030-1012012320323211-2111201022201331)
- service.advertise_options.advertise_in_cluster.multi_ports.ports.info.same_as_port

<a id="canonical-1100100120332232-2031310231312130-3033131012331023-2331230011302021-3233033232202301-0010313031313132-0333012102302210-3033111123331223"></a>

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
same_as_port = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2032101121303011-1103222233000330-0203232311001010-2322120323102000-3202033332221123-2310113223210032-0131222201221033-2112013321213130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_in_cluster.port` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_in_cluster](resources--workload--reference--group-008.md#canonical-0330023301131021-0130211002331300-3002122032320301-2112231133101233-3303101211030122-3300030030023300-1213020320030020-0213001312211213)
- service.advertise_options.advertise_in_cluster.port

<a id="canonical-2101101323122211-1123110212023233-2013223033212303-1110022300332000-3202333301102030-1131030212001331-2311013113311312-3003012311022321"></a>

Type: `"object"`. single nested block, Optional.

Port. Single port.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
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

<a id="canonical-3203221012003020-3303032232023210-1103311231233011-2201302023210110-3320032110122312-0110020010011320-2203303121312013-2012002300030103"></a>

### Direct properties for `service.advertise_options.advertise_in_cluster.port`

- [info](resources--workload--reference--group-008.md#canonical-3000021121133300-0112013033122013-0332121130331320-2320022103131303-3211132020321221-0102021013132121-3121210012321313-1032112131200133): complete subsection reference.

<a id="canonical-3000021121133300-0112013033122013-0332121130331320-2320022103131303-3211132020321221-0102021013132121-3121210012321313-1032112131200133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_in_cluster.port.info` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_in_cluster](resources--workload--reference--group-008.md#canonical-0330023301131021-0130211002331300-3002122032320301-2112231133101233-3303101211030122-3300030030023300-1213020320030020-0213001312211213)
- [service.advertise_options.advertise_in_cluster.port](resources--workload--reference--group-008.md#canonical-2032101121303011-1103222233000330-0203232311001010-2322120323102000-3202033332221123-2310113223210032-0131222201221033-2112013321213130)
- service.advertise_options.advertise_in_cluster.port.info

<a id="canonical-3102321302002202-1313330301020233-3220111220002100-1000101110111200-2210122331010211-2330233022003102-3221333221311132-3122313223223010"></a>

Type: `"object"`. single nested block, Optional.

Port Information. Port information.

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

<a id="canonical-3232012010303303-1212111321313200-1131303132331232-3121122131202110-3101232121231211-1301222130220132-0322102111311210-2321100101031010"></a>

### Direct properties for `service.advertise_options.advertise_in_cluster.port.info`

<a id="canonical-0203013130130121-2211213131321210-0110230102033000-3202003111200310-3231033231311301-2201220120101032-2211032220132101-2133230030021321"></a>

#### `service.advertise_options.advertise_in_cluster.port.info.port` property

Type: `"number"`. Optional.

Port. Port the workload can be reached on.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-2303100213300302-0020323211102102-0022222100230221-3213010211333131-0122201230303120-2003100210023103-1000330010313301-3303333003122111"></a>

<a id="canonical-3211101122322133-1301203032331312-0201203011220322-0331300221020323-2033302021023123-1102203102201031-0311302013023203-1032001022031212"></a>

#### `service.advertise_options.advertise_in_cluster.port.info.protocol` property

Type: `"string"`. Optional.

\[Enum: PROTOCOL\_TCP|PROTOCOL\_HTTP|PROTOCOL\_HTTP2|PROTOCOL\_TLS\_WITH\_SNI|PROTOCOL\_UDP\] Type
of protocol - PROTOCOL\_TCP: TCP TCP - PROTOCOL\_HTTP: HTTP HTTP - PROTOCOL\_HTTP2: HTTP2 HTTP2 -
PROTOCOL\_TLS\_WITH\_SNI: TLS with SNI TLS with SNI - PROTOCOL\_UDP: UDP UDP. Possible values are
\`PROTOCOL\_TCP\`, \`PROTOCOL\_HTTP\`, \`PROTOCOL\_HTTP2\`, \`PROTOCOL\_TLS\_WITH\_SNI\`,
\`PROTOCOL\_UDP\`. Defaults to \`PROTOCOL\_TCP\`.

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

- [same_as_port](resources--workload--reference--group-008.md#canonical-0331030002203301-2000202112302103-1121310003132201-1133302201230133-3333102123022321-3023221303113110-3202332010212000-1012220001110232): complete subsection reference.

<a id="canonical-2313132133230101-0113333300123012-3201210202101322-2230231221112011-0202100311201003-0233323033131220-2010111222221211-1120311323001122"></a>

<a id="canonical-1010300213103323-0233102322203322-0211313221110120-3021211003003301-0102233000313232-1213033220201302-2000103132121122-2102211021321231"></a>

#### `service.advertise_options.advertise_in_cluster.port.info.target_port` property

Type: `"number"`. Optional.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-0331030002203301-2000202112302103-1121310003132201-1133302201230133-3333102123022321-3023221303113110-3202332010212000-1012220001110232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_in_cluster.port.info.same_as_port` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_in_cluster](resources--workload--reference--group-008.md#canonical-0330023301131021-0130211002331300-3002122032320301-2112231133101233-3303101211030122-3300030030023300-1213020320030020-0213001312211213)
- [service.advertise_options.advertise_in_cluster.port](resources--workload--reference--group-008.md#canonical-2032101121303011-1103222233000330-0203232311001010-2322120323102000-3202033332221123-2310113223210032-0131222201221033-2112013321213130)
- [service.advertise_options.advertise_in_cluster.port.info](resources--workload--reference--group-008.md#canonical-3000021121133300-0112013033122013-0332121130331320-2320022103131303-3211132020321221-0102021013132121-3121210012321313-1032112131200133)
- service.advertise_options.advertise_in_cluster.port.info.same_as_port

<a id="canonical-2310130320320001-0321130312020100-1102312221011200-0331132113122323-0220233323202133-2231321002231203-2232231322102110-1233331301201022"></a>

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
same_as_port = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- service.advertise_options.advertise_on_public

<a id="canonical-3101011112220000-1112230302331111-0310010332131220-0300200300310022-0030222321031212-1120302233113133-0133032102333001-0032321033213020"></a>

Type: `"object"`. single nested block, Optional.

Advertise this workload via loadbalancer on internet with default VIP.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("multi_ports",
    "port")}
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
  "x-ves-oneof-field-advertise_choice": "[\"multi_ports\",\"port\"]"
}
```

Terraform syntax:

```terraform
advertise_on_public {
  # Configure direct properties listed below.
}
```

<a id="canonical-3232101100321100-3112220210213311-2300313113013112-3230031232321011-3312023000200031-2312310003001312-1030213202230022-1203200031222110"></a>

### Direct properties for `service.advertise_options.advertise_on_public`

- [multi_ports](resources--workload--reference--group-008.md#canonical-3210102113031103-2332201101220100-1103122221002212-3002101012001000-1332031130002011-0021231133303122-2121222212212213-1123320220010203): complete subsection reference.

- [port](resources--workload--reference--group-011.md#canonical-3013031222022030-2021222233122213-2220022232301003-3102220000332302-0311213311310132-3133321213021123-0133300222212330-1122131003023301): complete subsection reference.

<a id="canonical-3210102113031103-2332201101220100-1103122221002212-3002101012001000-1332031130002011-0021231133303122-2121222212212213-1123320220010203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.multi_ports` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- service.advertise_options.advertise_on_public.multi_ports

<a id="canonical-1030002101120112-1320101011031223-2332221312121111-3230202222201310-1320001001201120-0212221333123301-0013210013013102-2130112130310310"></a>

Type: `"object"`. single nested block, Optional.

Advertise Multiple Ports. Advertise multiple ports.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("ports")}
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
multi_ports {
  # Configure direct properties listed below.
}
```

<a id="canonical-2223010133032312-2020301302212331-2031232212223010-3000200230011323-3303120020320330-3003300023223302-2011013032101022-2310020121221012"></a>

### Direct properties for `service.advertise_options.advertise_on_public.multi_ports`

- [ports](resources--workload--reference--group-008.md#canonical-2311132233010120-0223230230122210-0332200203312210-3000132111100333-1133101301233010-3211310033122333-2323002110210000-1300012002110122): complete subsection reference.

<a id="canonical-2311132233010120-0223230230122210-0332200203312210-3000132111100333-1133101301233010-3211310033122333-2323002110210000-1300012002110122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.multi_ports.ports` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-008.md#canonical-3210102113031103-2332201101220100-1103122221002212-3002101012001000-1332031130002011-0021231133303122-2121222212212213-1123320220010203)
- service.advertise_options.advertise_on_public.multi_ports.ports

<a id="canonical-1212022013113012-1020030022221212-3120222021020231-2303332232323312-0301331020110111-2000302130200213-3322222032321023-1302033110100212"></a>

Type: `"object"`. list nested block, Optional.

Ports. Ports to advertise.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.ConflictingListObjectAttributes("http_loadbalancer",
    "tcp_loadbalancer")}
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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "16"
  }
}
```

Terraform syntax:

```terraform
ports {
  # Configure direct properties listed below.
}
```

<a id="canonical-0031032302032013-0102131002121322-0101100201033203-3312221312032212-2010120011031222-2022111120302331-3213013002033322-1130001331322301"></a>

### Direct properties for `service.advertise_options.advertise_on_public.multi_ports.ports`

- [http_loadbalancer](resources--workload--reference--group-008.md#canonical-1202020123020333-2332300320113330-1321112113331202-2112220001121101-3022001013310321-0020100001212100-0032232320012323-1311122212310031): complete subsection reference.

- [port](resources--workload--reference--group-011.md#canonical-0331331222102030-2203010231311222-1103320331132010-0113201101003101-3301201000333230-1221211320313131-1102020121310021-1030210213331030): complete subsection reference.

- [tcp_loadbalancer](resources--workload--reference--group-011.md#canonical-1201000230031032-2312033000023331-1111332222123100-2001203111302231-0133102032000130-2121122320101311-2223311020113123-1220222002313220): complete subsection reference.

<a id="canonical-1202020123020333-2332300320113330-1321112113331202-2112220001121101-3022001013310321-0020100001212100-0032232320012323-1311122212310031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-008.md#canonical-3210102113031103-2332201101220100-1103122221002212-3002101012001000-1332031130002011-0021231133303122-2121222212212213-1123320220010203)
- [service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-008.md#canonical-2311132233010120-0223230230122210-0332200203312210-3000132111100333-1133101301233010-3211310033122333-2323002110210000-1300012002110122)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer

<a id="canonical-0212010320303113-1320002303012030-3100331023002103-2332010133112122-0022210313302310-1220022300232120-2010322322321312-3321002321021320"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for http loadbalancer.

Additional upstream details:

HTTP/HTTPS Load balancer.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("domains"),
  validators.ConflictingObjectAttributes("default_route",
    "specific_routes"),
  validators.ConflictingObjectAttributes("http",
    "https"),
  validators.ConflictingObjectAttributes("http",
    "https_auto_cert"),
  validators.ConflictingObjectAttributes("https",
    "https_auto_cert")}
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
  "x-ves-oneof-field-loadbalancer_type": "[\"http\",\"https\",\"https_auto_cert\"]",
  "x-ves-oneof-field-route_choice": "[\"default_route\",\"specific_routes\"]"
}
```

Terraform syntax:

```terraform
http_loadbalancer {
  # Configure direct properties listed below.
}
```

<a id="canonical-2200331202212111-0230313203121013-2313211220233030-1303113332213013-1033311013133232-3100033303311300-1003020213330332-0121332213010000"></a>

### Direct properties for `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer`

- [default_route](resources--workload--reference--group-008.md#canonical-0232013311111022-0003320023211122-1321321011010330-1223120121220333-0213130300023300-2001123101323323-3002132100221212-2311012111330101): complete subsection reference.

<a id="canonical-2032303100131001-3230001322010003-3003202000122332-3011000120123020-2101013013103112-0132003102012100-3331201111300130-1220333313331030"></a>

<a id="canonical-2333303023130320-3320030321000323-1202033112210310-2201320303010331-1330301300212321-3120010101013112-1011032130131322-0233101013331211"></a>

#### `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.domains` property

Type: `["list", "string"]`. Optional.

List of domains (host/authority header) that will be matched to loadbalancer. Wildcard hosts are
supported in the suffix or prefix form Domain search order: 1. Exact domain names: \`\` is invalid
Domains are also used for SNI matching if the loadbalancer type is HTTPS Domains also indicate the..

Additional upstream details:

A list of domains (host/authority header) that will be matched to loadbalancer. Exact domain names:
\`\`www&#46;example.com\`\`. &#8203;2. Prefix domain wildcards: \`\`\*.example.com\`\` or
\`\`\*.bar.example.com\`\`. &#8203;3. Special wildcard \`\`\*\`\` matching any domain. Wildcard will
not match empty string. E.g. \`\`\*.example.com\`\` will match \`\`bar.example.com\`\` and
\`\`baz-bar.example.com\`\` but not \`\`.example.com\`\`. The longest wildcards match first.
Wildcards must match a whole DNS label. E.g. \`\`\*.example.com\`\` and \*.bar.example.com are
valid, however \`\`\*bar.example.com\`\` or \`\`\*-bar.example.com\`\` is invalid

Domains are also used for SNI matching if the loadbalancer type is HTTPS Domains also indicate the
list of names for which DNS resolution will be done by VER.

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
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 32,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

- [http](resources--workload--reference--group-008.md#canonical-2032123312300020-0313233113231322-1312210332031202-2133210330020310-0110031022301023-1212321320020231-2033231203003230-2310210201131220): complete subsection reference.

- [https](resources--workload--reference--group-008.md#canonical-3212310132100213-2000301220120030-0112323230310220-3200032133332210-3223310210320320-2331013310222010-2010102133130322-2111123310321223): complete subsection reference.

- [https_auto_cert](resources--workload--reference--group-010.md#canonical-0123203011200310-2323132333110331-0320211012002122-2323013220112302-2000030132131023-1312320130303213-2223130103233300-3130200202200302): complete subsection reference.

- [specific_routes](resources--workload--reference--group-010.md#canonical-3333233100013331-2322011200101232-2320001330000031-2000212200300310-3210113330031012-2012103222021320-0322112233332312-3331120230102200): complete subsection reference.

<a id="canonical-0232013311111022-0003320023211122-1321321011010330-1223120121220333-0213130300023300-2001123101323323-3002132100221212-2311012111330101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.default_route` properties

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

Additional upstream details:

Default route matching all APIs.

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
default_route {
  # Configure direct properties listed below.
}
```

<a id="canonical-1013322212100213-1012022103231212-3012110111323301-2002032133133031-2011310302320220-3102202010323302-3022231120101111-1002002132331203"></a>

### Direct properties for `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.default_route`

- [auto_host_rewrite](resources--workload--reference--group-008.md#canonical-2322120332220030-0300010303222132-0231233230310120-3312130212030203-3323022301121303-1300013301332323-1310232112322310-3321122202110221): complete subsection reference.

- [disable_host_rewrite](resources--workload--reference--group-008.md#canonical-1211002103313333-1231121300012003-1312223130210212-0002230330122201-0012213033030211-0020320230121233-2113131113313320-3110331123031333): complete subsection reference.

<a id="canonical-1333232102233301-2200021232321101-1212202010312032-0001211311320112-3010213230112123-0010130011022311-1303013101323111-3230333320313332"></a>

<a id="canonical-3201323120202231-2333133312312301-0030113110002131-0000011323313230-3100210131221011-2023000122233320-1331212320113022-3130132031111121"></a>

#### `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.default_route.host_rewrite` property

Type: `"string"`. Optional.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-2322120332220030-0300010303222132-0231233230310120-3312130212030203-3323022301121303-1300013301332323-1310232112322310-3321122202110221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.default_route.auto_host_rewrite` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-008.md#canonical-3210102113031103-2332201101220100-1103122221002212-3002101012001000-1332031130002011-0021231133303122-2121222212212213-1123320220010203)
- [service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-008.md#canonical-2311132233010120-0223230230122210-0332200203312210-3000132111100333-1133101301233010-3211310033122333-2323002110210000-1300012002110122)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](resources--workload--reference--group-008.md#canonical-1202020123020333-2332300320113330-1321112113331202-2112220001121101-3022001013310321-0020100001212100-0032232320012323-1311122212310031)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.default_route](resources--workload--reference--group-008.md#canonical-0232013311111022-0003320023211122-1321321011010330-1223120121220333-0213130300023300-2001123101323323-3002132100221212-2311012111330101)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.default_route.auto_host_rewrite

<a id="canonical-2021331302113012-2002312120020201-1201102220322311-3332112321330033-2032311123000131-3103221031211122-2111223213330233-0310230203113012"></a>

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
auto_host_rewrite = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1211002103313333-1231121300012003-1312223130210212-0002230330122201-0012213033030211-0020320230121233-2113131113313320-3110331123031333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.default_route.disable_host_rewrite` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-008.md#canonical-3210102113031103-2332201101220100-1103122221002212-3002101012001000-1332031130002011-0021231133303122-2121222212212213-1123320220010203)
- [service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-008.md#canonical-2311132233010120-0223230230122210-0332200203312210-3000132111100333-1133101301233010-3211310033122333-2323002110210000-1300012002110122)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](resources--workload--reference--group-008.md#canonical-1202020123020333-2332300320113330-1321112113331202-2112220001121101-3022001013310321-0020100001212100-0032232320012323-1311122212310031)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.default_route](resources--workload--reference--group-008.md#canonical-0232013311111022-0003320023211122-1321321011010330-1223120121220333-0213130300023300-2001123101323323-3002132100221212-2311012111330101)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.default_route.disable_host_rewrite

<a id="canonical-2212013233013002-1312320020100322-3302231222322102-0110111123312230-3100011031213210-0312300200200023-2001211011112031-2322002010130110"></a>

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
disable_host_rewrite = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2032123312300020-0313233113231322-1312210332031202-2133210330020310-0110031022301023-1212321320020231-2033231203003230-2310210201131220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.http` properties

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

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

<a id="canonical-3312112202231101-2220333012310323-3011311002330012-1030000321200021-3230021103122212-2133320311303231-2103113333321120-2013102332120022"></a>

### Direct properties for `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.http`

<a id="canonical-1223312120322332-2111100302101223-3230113112302323-3203010232202120-2030321301012203-1133302133122223-0202313321003212-1200021000203000"></a>

#### `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.http.dns_volterra_managed` property

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

<a id="canonical-1131223103331013-3031300201323330-2330032102302033-3322012101133320-2321021330220033-0100013222300132-0210313232132210-1012102033112323"></a>

<a id="canonical-3322212020322231-1032020312313301-3032201323220012-3302223323123233-2310022110110332-1110020013222312-1230021222323331-0201102013233323"></a>

#### `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.http.port` property

Type: `"number"`. Optional.

Exclusive with \[port\_ranges\] HTTP port to Listen.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-0331110211203033-0210332113310321-3021010133003201-3300330033322210-1211302000032133-0002321000023033-3210212130211321-2032120012320320"></a>

#### `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.http.port_ranges` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-3212310132100213-2000301220120030-0112323230310220-3200032133332210-3223310210320320-2331013310222010-2010102133130322-2111123310321223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https` properties

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

<a id="canonical-1200021102301023-1200220023333111-3011331033323131-1322220311203021-0331111212111132-1001130212323221-3211023132322032-2200002030230031"></a>

### Direct properties for `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https`

<a id="canonical-2100221011330111-1032323112010021-3000201322301230-1222103123300302-1030301203020232-1303021310232132-3330130113033132-3232223310012012"></a>

#### `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.add_hsts` property

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
