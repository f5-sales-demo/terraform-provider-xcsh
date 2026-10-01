---
page_title: "xcsh_azure_vnet_site reference"
subcategory: "Infrastructure"
description: "Complete grouped canonical reference for xcsh_azure_vnet_site reference."
---

# xcsh_azure_vnet_site reference

<a id="canonical-3102101320030322-0100202030102133-2030323312012031-0210221210232003-1120021331032030-3011320332032301-3222203100202321-1223022121211221"></a>

## ingress_egress_gw_ar.hub.express_route_disabled — express_route_disabled / 203203030103 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-3012120331203201-1112000130332131-3202332120003001-0123013103333002-1112013103100001-2021323312110130-1112011111221321-2323233113231232)
- [ingress_egress_gw_ar](resources--azure_vnet_site--reference--group-005.md#canonical-0312012003033102-3210313212120203-0230331020120000-0030121220200120-2223003302123103-2232111103103011-2023112221000212-0212333313012311)
- [ingress_egress_gw_ar.hub](resources--azure_vnet_site--reference--group-005.md#canonical-2322310312311313-3121121220121201-2231223330120013-2030033320203223-3023310333201212-1221020313120000-1032200112132013-1301213311322012)
- ingress_egress_gw_ar.hub.express_route_disabled

<a id="canonical-3013130123120301-2333111013003130-0013012120123000-0133211020132030-0231220131330200-2323032010110030-2130230023201021-3130221220312103"></a>

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
express_route_disabled = {}
```

<a id="canonical-3120101133101312-2122312001220221-1130110100021113-1110120232003100-1331130011223201-2201323132321032-2001021302001303-0231123331200232"></a>

## Direct properties — express_route_disabled / 203203030103 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3220111231310211-0230121311311201-0222031033023200-3233311302033212-1023213212313021-3010211331302130-3233333001121221-1213100323132332"></a>

## Next pages — express_route_disabled / 203203030103 / 4

- [ingress_egress_gw_ar.hub](resources--azure_vnet_site--reference--group-005.md#canonical-2322310312311313-3121121220121201-2231223330120013-2030033320203223-3023310333201212-1221020313120000-1032200112132013-1301213311322012)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)

<a id="canonical-2120301311230123-0230132233331312-1221203213113122-2231210212002002-1322021311302300-2201203011231232-1112123300032202-0231011112312022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0303121133233311-1011113131223330-2333201320220002-1221122030213123-2103331310103203-0103213323221123-3230220303013330-3323203002001100"></a>

## ingress_egress_gw_ar.hub.express_route_enabled — express_route_enabled / 212303311120 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-3012120331203201-1112000130332131-3202332120003001-0123013103333002-1112013103100001-2021323312110130-1112011111221321-2323233113231232)
- [ingress_egress_gw_ar](resources--azure_vnet_site--reference--group-005.md#canonical-0312012003033102-3210313212120203-0230331020120000-0030121220200120-2223003302123103-2232111103103011-2023112221000212-0212333313012311)
- [ingress_egress_gw_ar.hub](resources--azure_vnet_site--reference--group-005.md#canonical-2322310312311313-3121121220121201-2231223330120013-2030033320203223-3023310333201212-1221020313120000-1032200112132013-1301213311322012)
- ingress_egress_gw_ar.hub.express_route_enabled

<a id="canonical-2023312320010313-2033022331222013-3321222201032202-0200121300202321-3003012312003123-3311131330210323-1301333000010132-3120030330010121"></a>

Type: `"object"`. single nested block, Optional.

Express Route Configuration. Express Route Configuration.

Upstream description:

Express Route Configuration.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("connections"),
  validators.ConflictingObjectAttributes("advertise_to_route_server",
    "do_not_advertise_to_route_server"),
  validators.ConflictingObjectAttributes("auto_asn",
    "custom_asn"),
  validators.ConflictingObjectAttributes("site_registration_over_express_route",
    "site_registration_over_internet"),
  validators.ConflictingObjectAttributes("sku_ergw1az",
    "sku_ergw2az"),
  validators.ConflictingObjectAttributes("sku_ergw1az",
    "sku_high_perf"),
  validators.ConflictingObjectAttributes("sku_ergw1az",
    "sku_standard"),
  validators.ConflictingObjectAttributes("sku_ergw2az",
    "sku_high_perf"),
  validators.ConflictingObjectAttributes("sku_ergw2az",
    "sku_standard"),
  validators.ConflictingObjectAttributes("sku_high_perf",
    "sku_standard")}
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
  "x-ves-oneof-field-asn_choice": "[\"auto_asn\",\"custom_asn\"]",
  "x-ves-oneof-field-connectivity_options": "[\"site_registration_over_express_route\",\"site_registration_over_internet\"]",
  "x-ves-oneof-field-sku_choice": "[\"sku_ergw1az\",\"sku_ergw2az\",\"sku_high_perf\",\"sku_standard\"]",
  "x-ves-oneof-field-spoke_vnet_routes": "[\"advertise_to_route_server\",\"do_not_advertise_to_route_server\"]"
}
```

Terraform syntax:

```terraform
express_route_enabled {
  # Configure direct properties listed below.
}
```

<a id="canonical-3312333203122002-1331200222201110-0200022112303313-3111312021112213-3120031101110301-3132221201312213-1001022010200003-3332223022223022"></a>

## Direct properties — express_route_enabled / 212303311120 / 3

- [advertise_to_route_server](resources--azure_vnet_site--reference--group-006.md#canonical-2103233323022023-0322230231202222-3033110030100110-2112002213121110-1003123221133021-0202230303301212-0001330020113210-0311212130032011): complete subsection reference.

- [auto_asn](resources--azure_vnet_site--reference--group-006.md#canonical-3102101121113032-0002123112000112-1300232323112221-2003321033112011-1312112131302131-1010300011221121-1330103220321010-0202023220120120): complete subsection reference.

- [connections](resources--azure_vnet_site--reference--group-006.md#canonical-3330313320203323-2202333121331103-2233202311223312-2131323301311121-3310322231332133-0213233000211231-2230011313311233-1130001123332103): complete subsection reference.

<a id="canonical-3330320312313010-1032210012011113-2212012133331212-3021021203033303-0312013311033102-0201323130002232-3301001220320031-2310133232230220"></a>

<a id="canonical-0003123130233013-0003330220303012-0033112313030111-3132101103211033-3301222303031200-0101001300030032-1323323032131331-3013213120221230"></a>

## custom_asn property — express_route_enabled / 212303311120 / 4

Type: `"number"`. Optional.

Exclusive with \[auto\_asn\] Set custom ASN for F5XC Site.

Upstream description:

Exclusive with \[auto\_asn\] Set custom ASN for F5XC Site.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(2, 65535),
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
    "minimum": 2
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gt": "1",
    "ves.io.schema.rules.uint32.lte": "65535",
    "ves.io.schema.rules.uint32.not_in_ranges": "65515,65517,65518,65519,65520,8074,8075,12076,23456"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gt": "1",
    "ves.io.schema.rules.uint32.lte": "65535",
    "ves.io.schema.rules.uint32.not_in_ranges": "65515,65517,65518,65519,65520,8074,8075,12076,23456"
  }
}
```

- [do_not_advertise_to_route_server](resources--azure_vnet_site--reference--group-006.md#canonical-0020001210333110-3111230320332332-1011223222333222-2300033323111201-1000321333302132-3313113230331011-3110302333032230-3003300321212310): complete subsection reference.

- [gateway_subnet](resources--azure_vnet_site--reference--group-006.md#canonical-3213100012220112-3132130321000300-2200303301211013-0231210003133130-3320101132223312-3113123112213233-0120032210312300-3200212200020121): complete subsection reference.

- [route_server_subnet](resources--azure_vnet_site--reference--group-006.md#canonical-3011112311323102-0020113321321123-1001313030011200-3012033203133312-2111223131023023-1310233212123020-0020332113130102-1233133222033031): complete subsection reference.

- [site_registration_over_express_route](resources--azure_vnet_site--reference--group-006.md#canonical-0133201130033011-2110032111102032-1012313112111323-3032132132010310-0032213131123113-3123202230132110-0031111233020122-1003220012332110): complete subsection reference.

- [site_registration_over_internet](resources--azure_vnet_site--reference--group-006.md#canonical-0333100230231012-1313010222233022-1031333211333320-2212203101030310-3133210201222203-2200012233121222-2311220000321221-3322132022030133): complete subsection reference.

- [sku_ergw1az](resources--azure_vnet_site--reference--group-006.md#canonical-1300011101023331-2113300100003220-0100302131323123-3030131232113322-2000132030102011-0313030202130232-1313001311210120-3023100021220321): complete subsection reference.

- [sku_ergw2az](resources--azure_vnet_site--reference--group-006.md#canonical-2000033202113202-0130021211120213-3312320200331023-0322302332202000-0332232230313000-2110112221223213-1013331130001101-3003110002112002): complete subsection reference.

- [sku_high_perf](resources--azure_vnet_site--reference--group-006.md#canonical-1011002021003211-2013232003011230-3121330212102131-1221212122220002-1220230330222032-2300322211230103-1011330310313200-0131030132313020): complete subsection reference.

- [sku_standard](resources--azure_vnet_site--reference--group-006.md#canonical-0313103132221000-2312202121021330-1300303011001200-2121221233222122-2101133212022013-1232020120010110-3010232113230302-1222000133022122): complete subsection reference.

<a id="canonical-2320112200203213-3033300033301021-2333203322132201-0101222213213203-3030231200313031-2333100121210112-2022133130200220-2112022330321302"></a>

## Next pages — express_route_enabled / 212303311120 / 5

- [ingress_egress_gw_ar.hub.express_route_enabled.advertise_to_route_server](resources--azure_vnet_site--reference--group-006.md#canonical-2103233323022023-0322230231202222-3033110030100110-2112002213121110-1003123221133021-0202230303301212-0001330020113210-0311212130032011)
- [ingress_egress_gw_ar.hub.express_route_enabled.auto_asn](resources--azure_vnet_site--reference--group-006.md#canonical-3102101121113032-0002123112000112-1300232323112221-2003321033112011-1312112131302131-1010300011221121-1330103220321010-0202023220120120)
- [ingress_egress_gw_ar.hub.express_route_enabled.connections](resources--azure_vnet_site--reference--group-006.md#canonical-3330313320203323-2202333121331103-2233202311223312-2131323301311121-3310322231332133-0213233000211231-2230011313311233-1130001123332103)
- [ingress_egress_gw_ar.hub.express_route_enabled.do_not_advertise_to_route_server](resources--azure_vnet_site--reference--group-006.md#canonical-0020001210333110-3111230320332332-1011223222333222-2300033323111201-1000321333302132-3313113230331011-3110302333032230-3003300321212310)
- [ingress_egress_gw_ar.hub.express_route_enabled.gateway_subnet](resources--azure_vnet_site--reference--group-006.md#canonical-3213100012220112-3132130321000300-2200303301211013-0231210003133130-3320101132223312-3113123112213233-0120032210312300-3200212200020121)
- [ingress_egress_gw_ar.hub.express_route_enabled.route_server_subnet](resources--azure_vnet_site--reference--group-006.md#canonical-3011112311323102-0020113321321123-1001313030011200-3012033203133312-2111223131023023-1310233212123020-0020332113130102-1233133222033031)
- [ingress_egress_gw_ar.hub.express_route_enabled.site_registration_over_express_route](resources--azure_vnet_site--reference--group-006.md#canonical-0133201130033011-2110032111102032-1012313112111323-3032132132010310-0032213131123113-3123202230132110-0031111233020122-1003220012332110)
- [ingress_egress_gw_ar.hub.express_route_enabled.site_registration_over_internet](resources--azure_vnet_site--reference--group-006.md#canonical-0333100230231012-1313010222233022-1031333211333320-2212203101030310-3133210201222203-2200012233121222-2311220000321221-3322132022030133)
- [ingress_egress_gw_ar.hub.express_route_enabled.sku_ergw1az](resources--azure_vnet_site--reference--group-006.md#canonical-1300011101023331-2113300100003220-0100302131323123-3030131232113322-2000132030102011-0313030202130232-1313001311210120-3023100021220321)
- [ingress_egress_gw_ar.hub.express_route_enabled.sku_ergw2az](resources--azure_vnet_site--reference--group-006.md#canonical-2000033202113202-0130021211120213-3312320200331023-0322302332202000-0332232230313000-2110112221223213-1013331130001101-3003110002112002)
- [ingress_egress_gw_ar.hub.express_route_enabled.sku_high_perf](resources--azure_vnet_site--reference--group-006.md#canonical-1011002021003211-2013232003011230-3121330212102131-1221212122220002-1220230330222032-2300322211230103-1011330310313200-0131030132313020)
- [ingress_egress_gw_ar.hub.express_route_enabled.sku_standard](resources--azure_vnet_site--reference--group-006.md#canonical-0313103132221000-2312202121021330-1300303011001200-2121221233222122-2101133212022013-1232020120010110-3010232113230302-1222000133022122)
- [ingress_egress_gw_ar.hub](resources--azure_vnet_site--reference--group-005.md#canonical-2322310312311313-3121121220121201-2231223330120013-2030033320203223-3023310333201212-1221020313120000-1032200112132013-1301213311322012)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)

<a id="canonical-2103233323022023-0322230231202222-3033110030100110-2112002213121110-1003123221133021-0202230303301212-0001330020113210-0311212130032011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0000123321031301-3322032200223323-1132302302232213-0122120033002230-3111033110321332-1332012212121312-2313303123301210-1202321213020031"></a>

## ingress_egress_gw_ar.hub.express_route_enabled.advertise_to_route_server — advertise_to_route_server / 023212003103 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-3012120331203201-1112000130332131-3202332120003001-0123013103333002-1112013103100001-2021323312110130-1112011111221321-2323233113231232)
- [ingress_egress_gw_ar](resources--azure_vnet_site--reference--group-005.md#canonical-0312012003033102-3210313212120203-0230331020120000-0030121220200120-2223003302123103-2232111103103011-2023112221000212-0212333313012311)
- [ingress_egress_gw_ar.hub](resources--azure_vnet_site--reference--group-005.md#canonical-2322310312311313-3121121220121201-2231223330120013-2030033320203223-3023310333201212-1221020313120000-1032200112132013-1301213311322012)
- [ingress_egress_gw_ar.hub.express_route_enabled](resources--azure_vnet_site--reference--group-006.md#canonical-2120301311230123-0230132233331312-1221203213113122-2231210212002002-1322021311302300-2201203011231232-1112123300032202-0231011112312022)
- ingress_egress_gw_ar.hub.express_route_enabled.advertise_to_route_server

<a id="canonical-1020213111102002-1332222310020013-1122132200221332-3200110303300313-2010311302211022-1111303230030322-3310300313010011-0311003033322120"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for advertise to route server.

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
advertise_to_route_server = {}
```

<a id="canonical-3230110201101022-1303113122302231-1213033000110321-0133312311303323-1133233001213130-1230313013210013-3103031121321302-2011221001230011"></a>

## Direct properties — advertise_to_route_server / 023212003103 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3320212101103200-3000232020332322-2201012330320012-0023130022230320-2312301113123202-2031332203301331-1020131022113213-3333321130130112"></a>

## Next pages — advertise_to_route_server / 023212003103 / 4

- [ingress_egress_gw_ar.hub.express_route_enabled](resources--azure_vnet_site--reference--group-006.md#canonical-2120301311230123-0230132233331312-1221203213113122-2231210212002002-1322021311302300-2201203011231232-1112123300032202-0231011112312022)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)

<a id="canonical-3102101121113032-0002123112000112-1300232323112221-2003321033112011-1312112131302131-1010300011221121-1330103220321010-0202023220120120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2310110131323231-0331220223322120-0223013020231002-1333002012100202-3211112231212302-3203332032021032-2032313001032130-0030300013333302"></a>

## ingress_egress_gw_ar.hub.express_route_enabled.auto_asn — auto_asn / 230210203131 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-3012120331203201-1112000130332131-3202332120003001-0123013103333002-1112013103100001-2021323312110130-1112011111221321-2323233113231232)
- [ingress_egress_gw_ar](resources--azure_vnet_site--reference--group-005.md#canonical-0312012003033102-3210313212120203-0230331020120000-0030121220200120-2223003302123103-2232111103103011-2023112221000212-0212333313012311)
- [ingress_egress_gw_ar.hub](resources--azure_vnet_site--reference--group-005.md#canonical-2322310312311313-3121121220121201-2231223330120013-2030033320203223-3023310333201212-1221020313120000-1032200112132013-1301213311322012)
- [ingress_egress_gw_ar.hub.express_route_enabled](resources--azure_vnet_site--reference--group-006.md#canonical-2120301311230123-0230132233331312-1221203213113122-2231210212002002-1322021311302300-2201203011231232-1112123300032202-0231011112312022)
- ingress_egress_gw_ar.hub.express_route_enabled.auto_asn

<a id="canonical-1321210331300013-2113222032303003-2223223233120323-1220002133331201-2222201003313121-3033100011023333-0303030220030120-1200110133130200"></a>

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
auto_asn = {}
```

<a id="canonical-0223011032100130-1200032013321211-3012203101312232-1032101001310200-3223222131330322-0112222011003101-1110101001233023-0022300333121111"></a>

## Direct properties — auto_asn / 230210203131 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0310000003221123-1001003111033002-0123010213200300-2302322123300230-0203322332122030-0110121322332123-0202101112322013-1303032312120132"></a>

## Next pages — auto_asn / 230210203131 / 4

- [ingress_egress_gw_ar.hub.express_route_enabled](resources--azure_vnet_site--reference--group-006.md#canonical-2120301311230123-0230132233331312-1221203213113122-2231210212002002-1322021311302300-2201203011231232-1112123300032202-0231011112312022)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)

<a id="canonical-3330313320203323-2202333121331103-2233202311223312-2131323301311121-3310322231332133-0213233000211231-2230011313311233-1130001123332103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0232230231221031-2320113021001021-3112213111333023-3121202101302223-0132201200332122-3103302322123120-3120020023232013-3120220110111200"></a>

## ingress_egress_gw_ar.hub.express_route_enabled.connections — connections / 011012133013 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-3012120331203201-1112000130332131-3202332120003001-0123013103333002-1112013103100001-2021323312110130-1112011111221321-2323233113231232)
- [ingress_egress_gw_ar](resources--azure_vnet_site--reference--group-005.md#canonical-0312012003033102-3210313212120203-0230331020120000-0030121220200120-2223003302123103-2232111103103011-2023112221000212-0212333313012311)
- [ingress_egress_gw_ar.hub](resources--azure_vnet_site--reference--group-005.md#canonical-2322310312311313-3121121220121201-2231223330120013-2030033320203223-3023310333201212-1221020313120000-1032200112132013-1301213311322012)
- [ingress_egress_gw_ar.hub.express_route_enabled](resources--azure_vnet_site--reference--group-006.md#canonical-2120301311230123-0230132233331312-1221203213113122-2231210212002002-1322021311302300-2201203011231232-1112123300032202-0231011112312022)
- ingress_egress_gw_ar.hub.express_route_enabled.connections

<a id="canonical-0032021321330131-3101211200311100-1331022210331310-3130310020023112-0122213210203303-3030330310101321-0302331320322300-3233203322020001"></a>

Type: `"object"`. list nested block, Optional.

Add the ExpressRoute Circuit Connections to this site.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("circuit_id",
    "other_subscription")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 8,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 8,
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
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

Terraform syntax:

```terraform
connections {
  # Configure direct properties listed below.
}
```

<a id="canonical-0012301202233012-1213332030233130-1331103331202131-3301022211231120-3103303332201103-2030302212020103-1012311300202310-1232023202231231"></a>

## Direct properties — connections / 011012133013 / 3

<a id="canonical-3101333003113230-2000122103003312-1102222101220202-0320032101321030-0023321202302023-2032120223011133-3102302301310302-2333002030333131"></a>

<a id="canonical-1103223022032311-2121322200103121-2133003310203210-0011120203120111-0210301111223021-2101023322330313-1320311303020212-3223123330333021"></a>

## circuit_id property — connections / 011012133013 / 4

Type: `"string"`. Optional.

Exclusive with \[other\_subscription\] ExpressRoute Circuit is in same subscription as the site.

Upstream description:

Exclusive with \[other\_subscription\] ExpressRoute Circuit is in same subscription as the site.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(512),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 512,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 512,
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
    "ves.io.schema.rules.string.max_len": "512"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "512"
  }
}
```

- [metadata](resources--azure_vnet_site--reference--group-006.md#canonical-3113310100300132-2231320012030310-2323033133301311-0030312031032313-0322231120123300-2303333303033120-3130101203021120-1133203033023203): complete subsection reference.

- [other_subscription](resources--azure_vnet_site--reference--group-006.md#canonical-1013331330003322-3012211220211313-3023223112302312-0232301022021313-0313112000331101-3202120021200223-0203200133031122-1032131001113013): complete subsection reference.

<a id="canonical-3332013121111101-0212113111231010-1220310211211333-3112210203302102-0301312200032012-1322023010312110-1113103200221330-2033001313333101"></a>

<a id="canonical-3311333221201103-1213001313011101-3230122033321302-0322211211001101-1132312130213322-0111033123111001-0212022011121000-1111220212321222"></a>

## weight property — connections / 011012133013 / 5

Type: `"number"`. Optional.

The weight (or priority) for the routes received from this connection. The. Defaults to \`10\`.

Upstream description:

The weight (or priority) for the routes received from this connection. The default value is 10.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "load-balancing",
    "constraintType": "number",
    "maximum": 100,
    "metadata": {
      "confidence": 0.8,
      "source": "inferred",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 0,
    "multipleOf": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2112322333211103-2300333121113233-1000033333113011-0303332333310023-2201301223021322-3203202203030213-1223010103132130-0011330213203033"></a>

## Next pages — connections / 011012133013 / 6

- [ingress_egress_gw_ar.hub.express_route_enabled.connections.metadata](resources--azure_vnet_site--reference--group-006.md#canonical-3113310100300132-2231320012030310-2323033133301311-0030312031032313-0322231120123300-2303333303033120-3130101203021120-1133203033023203)
- [ingress_egress_gw_ar.hub.express_route_enabled.connections.other_subscription](resources--azure_vnet_site--reference--group-006.md#canonical-1013331330003322-3012211220211313-3023223112302312-0232301022021313-0313112000331101-3202120021200223-0203200133031122-1032131001113013)
- [ingress_egress_gw_ar.hub.express_route_enabled](resources--azure_vnet_site--reference--group-006.md#canonical-2120301311230123-0230132233331312-1221203213113122-2231210212002002-1322021311302300-2201203011231232-1112123300032202-0231011112312022)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)

<a id="canonical-3113310100300132-2231320012030310-2323033133301311-0030312031032313-0322231120123300-2303333303033120-3130101203021120-1133203033023203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3032033212221222-3212302310102100-2012200102300230-1111021332321213-0100332013110010-0311233331131233-0231223121021310-0332210330001201"></a>

## ingress_egress_gw_ar.hub.express_route_enabled.connections.metadata — metadata / 110230232021 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-3012120331203201-1112000130332131-3202332120003001-0123013103333002-1112013103100001-2021323312110130-1112011111221321-2323233113231232)
- [ingress_egress_gw_ar](resources--azure_vnet_site--reference--group-005.md#canonical-0312012003033102-3210313212120203-0230331020120000-0030121220200120-2223003302123103-2232111103103011-2023112221000212-0212333313012311)
- [ingress_egress_gw_ar.hub](resources--azure_vnet_site--reference--group-005.md#canonical-2322310312311313-3121121220121201-2231223330120013-2030033320203223-3023310333201212-1221020313120000-1032200112132013-1301213311322012)
- [ingress_egress_gw_ar.hub.express_route_enabled](resources--azure_vnet_site--reference--group-006.md#canonical-2120301311230123-0230132233331312-1221203213113122-2231210212002002-1322021311302300-2201203011231232-1112123300032202-0231011112312022)
- [ingress_egress_gw_ar.hub.express_route_enabled.connections](resources--azure_vnet_site--reference--group-006.md#canonical-3330313320203323-2202333121331103-2233202311223312-2131323301311121-3310322231332133-0213233000211231-2230011313311233-1130001123332103)
- ingress_egress_gw_ar.hub.express_route_enabled.connections.metadata

<a id="canonical-0222123011310320-2331122013221320-2331302113320020-3111222130012130-1131100022300311-0033010222121011-2231121100312113-0231000200211012"></a>

Type: `"object"`. single nested block, Optional.

MessageMetaType is metadata (common attributes) of a message that only certain messages have. This
information is propagated to the metadata of a child object that gets created from the containing
message during view processing. The information in this type can be specified by user during
create..

Upstream description:

MessageMetaType is metadata (common attributes) of a message that only certain messages have. This
information is propagated to the metadata of a child object that gets created from the containing
message during view processing. The information in this type can be specified by user during create
and replace APIs.

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
metadata {
  # Configure direct properties listed below.
}
```

<a id="canonical-3210023133122210-1010132123110103-1012202201212102-2010200022023012-2232312311231113-2101101212303322-2231120200021301-2202222133030221"></a>

## Direct properties — metadata / 110230232021 / 3

<a id="canonical-2332211233013231-0013233132323223-2120211023332330-1212011020312312-1230322231000321-1222011203330021-0003302333303110-1203310312211320"></a>

<a id="canonical-1213331223301101-1103003113310013-2011130330003330-0111322101130310-3022121322201000-1123121301222032-0132100030022112-2313311130011212"></a>

## description_spec property — metadata / 110230232021 / 4

Type: `"string"`. Optional.

Description. Human readable description.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

<a id="canonical-1320032132122232-2003202002312030-1100013230123222-0212322313222011-2332230031233212-0333300233210323-1332313301203322-0302111200102012"></a>

<a id="canonical-1231332212202323-1120311210132231-1212212130313122-3033230330032231-3200220020031002-3303031013113120-3132233000300322-1330113233321011"></a>

## name property — metadata / 110230232021 / 5

Type: `"string"`. Optional.

Name of the message. The value of name has to follow DNS-1035 format.

Upstream description:

This is the name of the message. The value of name has to follow DNS-1035 format.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

Receipt-pinned upstream constraints:

```json
{
  "minLength": 1,
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
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.ves_object_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.ves_object_name": "true"
  }
}
```

<a id="canonical-2101103201223032-1112233000102013-2323320032130131-3323121213210201-1200223020202212-1103213300322303-3323313302002102-2013021320110011"></a>

## Next pages — metadata / 110230232021 / 6

- [ingress_egress_gw_ar.hub.express_route_enabled.connections](resources--azure_vnet_site--reference--group-006.md#canonical-3330313320203323-2202333121331103-2233202311223312-2131323301311121-3310322231332133-0213233000211231-2230011313311233-1130001123332103)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)

<a id="canonical-1013331330003322-3012211220211313-3023223112302312-0232301022021313-0313112000331101-3202120021200223-0203200133031122-1032131001113013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2223131223310213-1302313003123211-0332112312233323-0120200303311010-2233333300102022-3001133312301221-0332312100003032-2220033110031111"></a>

## ingress_egress_gw_ar.hub.express_route_enabled.connections.other_subscription — other_subscription / 030131011023 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-3012120331203201-1112000130332131-3202332120003001-0123013103333002-1112013103100001-2021323312110130-1112011111221321-2323233113231232)
- [ingress_egress_gw_ar](resources--azure_vnet_site--reference--group-005.md#canonical-0312012003033102-3210313212120203-0230331020120000-0030121220200120-2223003302123103-2232111103103011-2023112221000212-0212333313012311)
- [ingress_egress_gw_ar.hub](resources--azure_vnet_site--reference--group-005.md#canonical-2322310312311313-3121121220121201-2231223330120013-2030033320203223-3023310333201212-1221020313120000-1032200112132013-1301213311322012)
- [ingress_egress_gw_ar.hub.express_route_enabled](resources--azure_vnet_site--reference--group-006.md#canonical-2120301311230123-0230132233331312-1221203213113122-2231210212002002-1322021311302300-2201203011231232-1112123300032202-0231011112312022)
- [ingress_egress_gw_ar.hub.express_route_enabled.connections](resources--azure_vnet_site--reference--group-006.md#canonical-3330313320203323-2202333121331103-2233202311223312-2131323301311121-3310322231332133-0213233000211231-2230011313311233-1130001123332103)
- ingress_egress_gw_ar.hub.express_route_enabled.connections.other_subscription

<a id="canonical-0233020212111130-1211221211221001-0013120132102300-3312033200000120-1200302033112012-3203002213011310-2121123203233102-1030021130001113"></a>

Type: `"object"`. single nested block, Optional.

Express Route Circuit Config From Other Subscription.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
other_subscription {
  # Configure direct properties listed below.
}
```

<a id="canonical-2103031000312103-1002112102200321-2120301131123030-3003232220330203-0100121131130211-2021221211211300-0002330001022121-1003110323300300"></a>

## Direct properties — other_subscription / 030131011023 / 3

- [authorized_key](resources--azure_vnet_site--reference--group-006.md#canonical-3232003323122302-1230000103301323-3301112210011332-2310131213033333-3311200031221221-3202031103221200-3013310203220232-1213132032323203): complete subsection reference.

<a id="canonical-1110230032032302-1010120302220120-1331311303130101-1110323211100131-0020111113110131-2210031201132300-1333102323301223-3031320310302021"></a>

<a id="canonical-1303320310112032-2133130232112201-0031121000121220-3030330010000122-3131001011222033-0212231130031131-1201202321313323-3323330001201110"></a>

## circuit_id property — other_subscription / 030131011023 / 4

Type: `"string"`. Optional.

Circuit ID. Circuit ID.

Upstream description:

Circuit ID.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(512),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 512,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 512,
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
    "ves.io.schema.rules.string.max_len": "512"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "512"
  }
}
```

<a id="canonical-2233002322130331-3123222100110213-3020130231230222-3122131222031103-3331330023310001-0222233223301120-1021033222333122-3010313121012200"></a>

## Next pages — other_subscription / 030131011023 / 5

- [ingress_egress_gw_ar.hub.express_route_enabled.connections.other_subscription.authorized_key](resources--azure_vnet_site--reference--group-006.md#canonical-3232003323122302-1230000103301323-3301112210011332-2310131213033333-3311200031221221-3202031103221200-3013310203220232-1213132032323203)
- [ingress_egress_gw_ar.hub.express_route_enabled.connections](resources--azure_vnet_site--reference--group-006.md#canonical-3330313320203323-2202333121331103-2233202311223312-2131323301311121-3310322231332133-0213233000211231-2230011313311233-1130001123332103)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)

<a id="canonical-3232003323122302-1230000103301323-3301112210011332-2310131213033333-3311200031221221-3202031103221200-3013310203220232-1213132032323203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2303210220100220-1231003321210323-1013303012101102-0313031032322110-0133001210112220-2120212323111320-2213021003210133-2200210232312102"></a>

## ingress_egress_gw_ar.hub.express_route_enabled.connections.other_subscription.authorized_key — authorized_key / 232130131220 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-3012120331203201-1112000130332131-3202332120003001-0123013103333002-1112013103100001-2021323312110130-1112011111221321-2323233113231232)
- [ingress_egress_gw_ar](resources--azure_vnet_site--reference--group-005.md#canonical-0312012003033102-3210313212120203-0230331020120000-0030121220200120-2223003302123103-2232111103103011-2023112221000212-0212333313012311)
- [ingress_egress_gw_ar.hub](resources--azure_vnet_site--reference--group-005.md#canonical-2322310312311313-3121121220121201-2231223330120013-2030033320203223-3023310333201212-1221020313120000-1032200112132013-1301213311322012)
- [ingress_egress_gw_ar.hub.express_route_enabled](resources--azure_vnet_site--reference--group-006.md#canonical-2120301311230123-0230132233331312-1221203213113122-2231210212002002-1322021311302300-2201203011231232-1112123300032202-0231011112312022)
- [ingress_egress_gw_ar.hub.express_route_enabled.connections](resources--azure_vnet_site--reference--group-006.md#canonical-3330313320203323-2202333121331103-2233202311223312-2131323301311121-3310322231332133-0213233000211231-2230011313311233-1130001123332103)
- [ingress_egress_gw_ar.hub.express_route_enabled.connections.other_subscription](resources--azure_vnet_site--reference--group-006.md#canonical-1013331330003322-3012211220211313-3023223112302312-0232301022021313-0313112000331101-3202120021200223-0203200133031122-1032131001113013)
- ingress_egress_gw_ar.hub.express_route_enabled.connections.other_subscription.authorized_key

<a id="canonical-3311322031001332-2023121023103213-0111202230102332-1213201032233233-2020302023230220-2121321301320332-3132003032131331-1003102211322210"></a>

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
authorized_key {
  # Configure direct properties listed below.
}
```

<a id="canonical-3330130333312001-3120221231122312-3012222133121101-2321131301222001-0311011201002002-2231011111200322-0201133232110032-3023101300311002"></a>

## Direct properties — authorized_key / 232130131220 / 3

- [blindfold_secret_info](resources--azure_vnet_site--reference--group-006.md#canonical-2211011302231302-2030022021201211-1312333212233012-3122033302230111-1033221011301200-2111010311232132-1332001000033111-3033310200300212): complete subsection reference.

- [clear_secret_info](resources--azure_vnet_site--reference--group-006.md#canonical-2131311333301230-0130112200332310-1000002231131101-1220003002101232-0030221232000202-2300100003300020-3020222113203203-3110130212013010): complete subsection reference.

<a id="canonical-0223231312202112-3213312301321110-2220110131303022-0011203301300131-1330012231021332-1200111200330013-2003002101030011-3310320132012202"></a>

## Next pages — authorized_key / 232130131220 / 4

- [ingress_egress_gw_ar.hub.express_route_enabled.connections.other_subscription.authorized_key.blindfold_secret_info](resources--azure_vnet_site--reference--group-006.md#canonical-2211011302231302-2030022021201211-1312333212233012-3122033302230111-1033221011301200-2111010311232132-1332001000033111-3033310200300212)
- [ingress_egress_gw_ar.hub.express_route_enabled.connections.other_subscription.authorized_key.clear_secret_info](resources--azure_vnet_site--reference--group-006.md#canonical-2131311333301230-0130112200332310-1000002231131101-1220003002101232-0030221232000202-2300100003300020-3020222113203203-3110130212013010)
- [ingress_egress_gw_ar.hub.express_route_enabled.connections.other_subscription](resources--azure_vnet_site--reference--group-006.md#canonical-1013331330003322-3012211220211313-3023223112302312-0232301022021313-0313112000331101-3202120021200223-0203200133031122-1032131001113013)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)

<a id="canonical-2211011302231302-2030022021201211-1312333212233012-3122033302230111-1033221011301200-2111010311232132-1332001000033111-3033310200300212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2303012020201130-0322102022121020-2113112222100003-0320222223331023-2311201233013212-3223123101110012-2310023233103032-2233012201112310"></a>

## ingress_egress_gw_ar.hub.express_route_enabled.connections.other_subscription.authorized_key.blindfold_secret_info — blindfold_secret_info / 302010023310 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-3012120331203201-1112000130332131-3202332120003001-0123013103333002-1112013103100001-2021323312110130-1112011111221321-2323233113231232)
- [ingress_egress_gw_ar](resources--azure_vnet_site--reference--group-005.md#canonical-0312012003033102-3210313212120203-0230331020120000-0030121220200120-2223003302123103-2232111103103011-2023112221000212-0212333313012311)
- [ingress_egress_gw_ar.hub](resources--azure_vnet_site--reference--group-005.md#canonical-2322310312311313-3121121220121201-2231223330120013-2030033320203223-3023310333201212-1221020313120000-1032200112132013-1301213311322012)
- [ingress_egress_gw_ar.hub.express_route_enabled](resources--azure_vnet_site--reference--group-006.md#canonical-2120301311230123-0230132233331312-1221203213113122-2231210212002002-1322021311302300-2201203011231232-1112123300032202-0231011112312022)
- [ingress_egress_gw_ar.hub.express_route_enabled.connections](resources--azure_vnet_site--reference--group-006.md#canonical-3330313320203323-2202333121331103-2233202311223312-2131323301311121-3310322231332133-0213233000211231-2230011313311233-1130001123332103)
- [ingress_egress_gw_ar.hub.express_route_enabled.connections.other_subscription](resources--azure_vnet_site--reference--group-006.md#canonical-1013331330003322-3012211220211313-3023223112302312-0232301022021313-0313112000331101-3202120021200223-0203200133031122-1032131001113013)
- [ingress_egress_gw_ar.hub.express_route_enabled.connections.other_subscription.authorized_key](resources--azure_vnet_site--reference--group-006.md#canonical-3232003323122302-1230000103301323-3301112210011332-2310131213033333-3311200031221221-3202031103221200-3013310203220232-1213132032323203)
- ingress_egress_gw_ar.hub.express_route_enabled.connections.other_subscription.authorized_key.blindfold_secret_info

<a id="canonical-0220231033003300-2133130011202003-0101032021023120-2313012123222123-2231030323030202-0223010000112313-1313220010301133-2301132332213001"></a>

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

<a id="canonical-1021220100213102-3033023332113021-3333012202230202-3200322202003213-0300010100312130-0033000313022123-0310010302221200-0111203020113102"></a>

## Direct properties — blindfold_secret_info / 302010023310 / 3

<a id="canonical-0303310230321330-2022301132223330-2220102012012112-3301220301130003-0311220013032122-0332022032311113-3003320323221320-1200111110010233"></a>

<a id="canonical-2120210000310320-2213220230223103-2011101220202303-0303312030222001-1023001013223033-2020320130332303-0321113011220131-3313302123332133"></a>

## decryption_provider property — blindfold_secret_info / 302010023310 / 4

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

<a id="canonical-0110103013111312-2223220003001302-0312213131031203-3101230320001130-1103110131003222-0233022230322211-3110010101303221-3103320323233231"></a>

<a id="canonical-3030130331112033-1221213202030120-1120120131330111-0212013033220321-0320110321030121-0220031201232220-1131001311001103-2223313220103211"></a>

## location property — blindfold_secret_info / 302010023310 / 5

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

<a id="canonical-0323312021123300-0213032030212111-1312311213312010-3120032012011301-0103133300130103-3223021003323223-1012212101030202-3102203302320003"></a>

<a id="canonical-3212122210111110-3210021333003212-3001132011211301-1313132133032012-0222222122032320-0213012313131320-3020212310312030-1202101132303302"></a>

## store_provider property — blindfold_secret_info / 302010023310 / 6

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

<a id="canonical-0000332012212320-1130320213213001-2133021100212121-1300330233113301-2122101102132131-2120033331122201-1230100311112133-1311323133212231"></a>

## Next pages — blindfold_secret_info / 302010023310 / 7

- [ingress_egress_gw_ar.hub.express_route_enabled.connections.other_subscription.authorized_key](resources--azure_vnet_site--reference--group-006.md#canonical-3232003323122302-1230000103301323-3301112210011332-2310131213033333-3311200031221221-3202031103221200-3013310203220232-1213132032323203)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)

<a id="canonical-2131311333301230-0130112200332310-1000002231131101-1220003002101232-0030221232000202-2300100003300020-3020222113203203-3110130212013010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2202212311211032-3322230022222230-3132331123022132-3202233210103313-0103021121303230-2031200013032101-3111201032022323-0203300111231201"></a>

## ingress_egress_gw_ar.hub.express_route_enabled.connections.other_subscription.authorized_key.clear_secret_info — clear_secret_info / 210323300112 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-3012120331203201-1112000130332131-3202332120003001-0123013103333002-1112013103100001-2021323312110130-1112011111221321-2323233113231232)
- [ingress_egress_gw_ar](resources--azure_vnet_site--reference--group-005.md#canonical-0312012003033102-3210313212120203-0230331020120000-0030121220200120-2223003302123103-2232111103103011-2023112221000212-0212333313012311)
- [ingress_egress_gw_ar.hub](resources--azure_vnet_site--reference--group-005.md#canonical-2322310312311313-3121121220121201-2231223330120013-2030033320203223-3023310333201212-1221020313120000-1032200112132013-1301213311322012)
- [ingress_egress_gw_ar.hub.express_route_enabled](resources--azure_vnet_site--reference--group-006.md#canonical-2120301311230123-0230132233331312-1221203213113122-2231210212002002-1322021311302300-2201203011231232-1112123300032202-0231011112312022)
- [ingress_egress_gw_ar.hub.express_route_enabled.connections](resources--azure_vnet_site--reference--group-006.md#canonical-3330313320203323-2202333121331103-2233202311223312-2131323301311121-3310322231332133-0213233000211231-2230011313311233-1130001123332103)
- [ingress_egress_gw_ar.hub.express_route_enabled.connections.other_subscription](resources--azure_vnet_site--reference--group-006.md#canonical-1013331330003322-3012211220211313-3023223112302312-0232301022021313-0313112000331101-3202120021200223-0203200133031122-1032131001113013)
- [ingress_egress_gw_ar.hub.express_route_enabled.connections.other_subscription.authorized_key](resources--azure_vnet_site--reference--group-006.md#canonical-3232003323122302-1230000103301323-3301112210011332-2310131213033333-3311200031221221-3202031103221200-3013310203220232-1213132032323203)
- ingress_egress_gw_ar.hub.express_route_enabled.connections.other_subscription.authorized_key.clear_secret_info

<a id="canonical-3231133022313102-2220122211210332-1100011110321312-0211012101003023-2302123113023020-1033130123001103-1310312223201211-3111312021110010"></a>

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

<a id="canonical-1130223100310020-2132210212213003-3203233121102213-3033302202020000-1303330302021011-3220111210310301-1022233122033202-3100320111120002"></a>

## Direct properties — clear_secret_info / 210323300112 / 3

<a id="canonical-0112020232301011-0302210213201233-3000313233213213-2201001230330023-0232330330022222-2130112113110222-1310333011230100-1211113113311330"></a>

<a id="canonical-1020301200320013-3332131110310122-0133332300111223-3310012122202301-3213312111211300-0031033102003310-0313310200010130-3021310100302131"></a>

## provider_ref property — clear_secret_info / 210323300112 / 4

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-3310200101031302-0001310030132120-1321333202300213-2101110000023200-3221310010330303-3221123300112233-2230011121230111-2132202111110303"></a>

<a id="canonical-1003030130013202-3321020121310331-3103101122020201-2111201222210332-0330332223313221-0310321203100322-1133333022311001-2213030321330213"></a>

## URL property — clear_secret_info / 210323300112 / 5

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

<a id="canonical-3213122131013202-2000213210120032-3201130322311012-0111132020233300-2020120101321302-1210220023020002-2220313301031201-2210303313133313"></a>

## Next pages — clear_secret_info / 210323300112 / 6

- [ingress_egress_gw_ar.hub.express_route_enabled.connections.other_subscription.authorized_key](resources--azure_vnet_site--reference--group-006.md#canonical-3232003323122302-1230000103301323-3301112210011332-2310131213033333-3311200031221221-3202031103221200-3013310203220232-1213132032323203)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)

<a id="canonical-0020001210333110-3111230320332332-1011223222333222-2300033323111201-1000321333302132-3313113230331011-3110302333032230-3003300321212310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1020013002332110-0220331113002203-0312322120131221-1200222222002222-2311333233022123-2110323111210201-0222013113123312-0112311121200213"></a>

## ingress_egress_gw_ar.hub.express_route_enabled.do_not_advertise_to_route_server — do_not_advertise_to_route_server / 220213111200 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-3012120331203201-1112000130332131-3202332120003001-0123013103333002-1112013103100001-2021323312110130-1112011111221321-2323233113231232)
- [ingress_egress_gw_ar](resources--azure_vnet_site--reference--group-005.md#canonical-0312012003033102-3210313212120203-0230331020120000-0030121220200120-2223003302123103-2232111103103011-2023112221000212-0212333313012311)
- [ingress_egress_gw_ar.hub](resources--azure_vnet_site--reference--group-005.md#canonical-2322310312311313-3121121220121201-2231223330120013-2030033320203223-3023310333201212-1221020313120000-1032200112132013-1301213311322012)
- [ingress_egress_gw_ar.hub.express_route_enabled](resources--azure_vnet_site--reference--group-006.md#canonical-2120301311230123-0230132233331312-1221203213113122-2231210212002002-1322021311302300-2201203011231232-1112123300032202-0231011112312022)
- ingress_egress_gw_ar.hub.express_route_enabled.do_not_advertise_to_route_server

<a id="canonical-1300232233013221-0113212301200231-3012231231332330-3330120330001333-0002221320030311-3223033013102112-1213011323213000-2032322303032203"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for do not advertise to route server.

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
do_not_advertise_to_route_server = {}
```

<a id="canonical-1033021232330323-1011332330120321-1213212221131012-1210312200230010-3120133213212332-2032033312031100-1302322320121233-0321030300203003"></a>

## Direct properties — do_not_advertise_to_route_server / 220213111200 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0333131322001322-3003122230311233-1212132110203322-0322021312231300-0003012003113001-2302112330100103-2331113301213101-3011013213122222"></a>

## Next pages — do_not_advertise_to_route_server / 220213111200 / 4

- [ingress_egress_gw_ar.hub.express_route_enabled](resources--azure_vnet_site--reference--group-006.md#canonical-2120301311230123-0230132233331312-1221203213113122-2231210212002002-1322021311302300-2201203011231232-1112123300032202-0231011112312022)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)

<a id="canonical-3213100012220112-3132130321000300-2200303301211013-0231210003133130-3320101132223312-3113123112213233-0120032210312300-3200212200020121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2110302212303012-2102213330311332-3131030312130001-3000000223312232-1121020230110123-3103021233230123-3232033213221012-3121002001213123"></a>

## ingress_egress_gw_ar.hub.express_route_enabled.gateway_subnet — gateway_subnet / 120022222230 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-3012120331203201-1112000130332131-3202332120003001-0123013103333002-1112013103100001-2021323312110130-1112011111221321-2323233113231232)
- [ingress_egress_gw_ar](resources--azure_vnet_site--reference--group-005.md#canonical-0312012003033102-3210313212120203-0230331020120000-0030121220200120-2223003302123103-2232111103103011-2023112221000212-0212333313012311)
- [ingress_egress_gw_ar.hub](resources--azure_vnet_site--reference--group-005.md#canonical-2322310312311313-3121121220121201-2231223330120013-2030033320203223-3023310333201212-1221020313120000-1032200112132013-1301213311322012)
- [ingress_egress_gw_ar.hub.express_route_enabled](resources--azure_vnet_site--reference--group-006.md#canonical-2120301311230123-0230132233331312-1221203213113122-2231210212002002-1322021311302300-2201203011231232-1112123300032202-0231011112312022)
- ingress_egress_gw_ar.hub.express_route_enabled.gateway_subnet

<a id="canonical-3002132131203222-1133113133020112-3202020011313103-0103022102120202-1013320223311112-2112112233201100-1123011130011212-0201201021021223"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for gateway subnet.

Upstream description:

Parameters for Azure subnet.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("auto",
    "subnet"),
  validators.ConflictingObjectAttributes("auto",
    "subnet_param"),
  validators.ConflictingObjectAttributes("subnet",
    "subnet_param")}
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
  "x-ves-oneof-field-choice": "[\"auto\",\"subnet\",\"subnet_param\"]"
}
```

Terraform syntax:

```terraform
gateway_subnet {
  # Configure direct properties listed below.
}
```

<a id="canonical-0232322030220200-2320323323203320-3100300230013212-3120011210202302-3330312311301322-1121001213331300-3103320023312332-0323131123130233"></a>

## Direct properties — gateway_subnet / 120022222230 / 3

- [auto](resources--azure_vnet_site--reference--group-006.md#canonical-3311103010332020-3022202223332330-1321321233233222-3031221102231022-2110033101301123-0031303033231212-1330202033131203-0333011321113111): complete subsection reference.

- [subnet](resources--azure_vnet_site--reference--group-006.md#canonical-2132320020301021-2213302323331120-1000012100310233-0311313003113331-0133233102301121-2102000200203332-0012223312121212-1121303121030222): complete subsection reference.

- [subnet_param](resources--azure_vnet_site--reference--group-006.md#canonical-2132221021020220-1020302012003132-0003131003131010-1000022013333313-2200030021322212-2022320323321332-1311232013233313-3011033313031021): complete subsection reference.

<a id="canonical-1130222130013012-0221111203030200-2330303203111001-2111130110010203-3111001330121222-1323002032323210-2231123232002120-0110333030300321"></a>

## Next pages — gateway_subnet / 120022222230 / 4

- [ingress_egress_gw_ar.hub.express_route_enabled.gateway_subnet.auto](resources--azure_vnet_site--reference--group-006.md#canonical-3311103010332020-3022202223332330-1321321233233222-3031221102231022-2110033101301123-0031303033231212-1330202033131203-0333011321113111)
- [ingress_egress_gw_ar.hub.express_route_enabled.gateway_subnet.subnet](resources--azure_vnet_site--reference--group-006.md#canonical-2132320020301021-2213302323331120-1000012100310233-0311313003113331-0133233102301121-2102000200203332-0012223312121212-1121303121030222)
- [ingress_egress_gw_ar.hub.express_route_enabled.gateway_subnet.subnet_param](resources--azure_vnet_site--reference--group-006.md#canonical-2132221021020220-1020302012003132-0003131003131010-1000022013333313-2200030021322212-2022320323321332-1311232013233313-3011033313031021)
- [ingress_egress_gw_ar.hub.express_route_enabled](resources--azure_vnet_site--reference--group-006.md#canonical-2120301311230123-0230132233331312-1221203213113122-2231210212002002-1322021311302300-2201203011231232-1112123300032202-0231011112312022)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)

<a id="canonical-3311103010332020-3022202223332330-1321321233233222-3031221102231022-2110033101301123-0031303033231212-1330202033131203-0333011321113111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2033121112012222-2221031031021121-0131021111133030-0333112323010302-3210232101202302-3303320120001002-1332121320231022-2033331121331013"></a>

## ingress_egress_gw_ar.hub.express_route_enabled.gateway_subnet.auto — auto / 201312032021 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-3012120331203201-1112000130332131-3202332120003001-0123013103333002-1112013103100001-2021323312110130-1112011111221321-2323233113231232)
- [ingress_egress_gw_ar](resources--azure_vnet_site--reference--group-005.md#canonical-0312012003033102-3210313212120203-0230331020120000-0030121220200120-2223003302123103-2232111103103011-2023112221000212-0212333313012311)
- [ingress_egress_gw_ar.hub](resources--azure_vnet_site--reference--group-005.md#canonical-2322310312311313-3121121220121201-2231223330120013-2030033320203223-3023310333201212-1221020313120000-1032200112132013-1301213311322012)
- [ingress_egress_gw_ar.hub.express_route_enabled](resources--azure_vnet_site--reference--group-006.md#canonical-2120301311230123-0230132233331312-1221203213113122-2231210212002002-1322021311302300-2201203011231232-1112123300032202-0231011112312022)
- [ingress_egress_gw_ar.hub.express_route_enabled.gateway_subnet](resources--azure_vnet_site--reference--group-006.md#canonical-3213100012220112-3132130321000300-2200303301211013-0231210003133130-3320101132223312-3113123112213233-0120032210312300-3200212200020121)
- ingress_egress_gw_ar.hub.express_route_enabled.gateway_subnet.auto

<a id="canonical-2201021312303300-0032223321303313-0130110200332232-0100023001002320-3203102130222131-3311302310023213-3300101303201022-1001012312322031"></a>

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
auto = {}
```

<a id="canonical-0102332011002032-0022201022303330-2111032221102111-0310232230013301-1022213112113222-0332211102210032-1112311133232232-2133020202031330"></a>

## Direct properties — auto / 201312032021 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2033203220121332-3023203032033220-2321213010311211-1133210302013122-0312231122323000-1311201203133301-1132133011212330-2011031101102333"></a>

## Next pages — auto / 201312032021 / 4

- [ingress_egress_gw_ar.hub.express_route_enabled.gateway_subnet](resources--azure_vnet_site--reference--group-006.md#canonical-3213100012220112-3132130321000300-2200303301211013-0231210003133130-3320101132223312-3113123112213233-0120032210312300-3200212200020121)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)

<a id="canonical-2132320020301021-2213302323331120-1000012100310233-0311313003113331-0133233102301121-2102000200203332-0012223312121212-1121303121030222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2221223322020013-0011032022321300-1210112113311002-0003021211133012-0223000023331322-0201321311232022-1122013113031321-3011301231231010"></a>

## ingress_egress_gw_ar.hub.express_route_enabled.gateway_subnet.subnet — subnet / 013002032002 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-3012120331203201-1112000130332131-3202332120003001-0123013103333002-1112013103100001-2021323312110130-1112011111221321-2323233113231232)
- [ingress_egress_gw_ar](resources--azure_vnet_site--reference--group-005.md#canonical-0312012003033102-3210313212120203-0230331020120000-0030121220200120-2223003302123103-2232111103103011-2023112221000212-0212333313012311)
- [ingress_egress_gw_ar.hub](resources--azure_vnet_site--reference--group-005.md#canonical-2322310312311313-3121121220121201-2231223330120013-2030033320203223-3023310333201212-1221020313120000-1032200112132013-1301213311322012)
- [ingress_egress_gw_ar.hub.express_route_enabled](resources--azure_vnet_site--reference--group-006.md#canonical-2120301311230123-0230132233331312-1221203213113122-2231210212002002-1322021311302300-2201203011231232-1112123300032202-0231011112312022)
- [ingress_egress_gw_ar.hub.express_route_enabled.gateway_subnet](resources--azure_vnet_site--reference--group-006.md#canonical-3213100012220112-3132130321000300-2200303301211013-0231210003133130-3320101132223312-3113123112213233-0120032210312300-3200212200020121)
- ingress_egress_gw_ar.hub.express_route_enabled.gateway_subnet.subnet

<a id="canonical-0331331111023320-3113203012030010-0223023203023221-2303200131022101-3030311102113333-1133130320022022-0123013112031230-1110202303310132"></a>

Type: `"object"`. single nested block, Optional.

Parameters for Azure special subnet which name is reserved. (i.e GatewaySubnet or
RouteServerSubnet).

Upstream description:

Parameters for Azure special subnet which name is reserved. (i.e GatewaySubnet or RouteServerSubnet)

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("subnet_resource_grp",
    "vnet_resource_group")}
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
  "x-ves-oneof-field-resource_group_choice": "[\"subnet_resource_grp\",\"vnet_resource_group\"]"
}
```

Terraform syntax:

```terraform
subnet {
  # Configure direct properties listed below.
}
```

<a id="canonical-0021130112021120-3313122220130313-3331120333201103-3221311022330331-0121122221012112-2201100210203212-0233001300030221-0100331130330212"></a>

## Direct properties — subnet / 013002032002 / 3

<a id="canonical-0111203231130121-1223321202110202-3301130100221231-3011003231032312-3131230302230203-1032320200022231-3320032302300323-3310133220000212"></a>

<a id="canonical-2023222123020221-2230233133033123-0333300333321003-0201003310121313-1313021100031113-3331020333313111-2330112133322032-2210001000130020"></a>

## subnet_resource_grp property — subnet / 013002032002 / 4

Type: `"string"`. Optional.

Exclusive with \[vnet\_resource\_group\] Specify name of Resource Group.

Upstream description:

Exclusive with \[vnet\_resource\_group\] Specify name of Resource Group.

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
    "ves.io.schema.rules.string.max_len": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64"
  }
}
```

- [vnet_resource_group](resources--azure_vnet_site--reference--group-006.md#canonical-2302133312113122-2103131230100032-2022320201011133-1233322313100303-1311332312100000-1300030220033203-2311003202110111-2110003001003333): complete subsection reference.

<a id="canonical-3013210223313111-0221033213210333-2133113312003031-2001020032222102-2031002321132033-2130010101221332-3122101120201030-1322030030232230"></a>

## Next pages — subnet / 013002032002 / 5

- [ingress_egress_gw_ar.hub.express_route_enabled.gateway_subnet.subnet.vnet_resource_group](resources--azure_vnet_site--reference--group-006.md#canonical-2302133312113122-2103131230100032-2022320201011133-1233322313100303-1311332312100000-1300030220033203-2311003202110111-2110003001003333)
- [ingress_egress_gw_ar.hub.express_route_enabled.gateway_subnet](resources--azure_vnet_site--reference--group-006.md#canonical-3213100012220112-3132130321000300-2200303301211013-0231210003133130-3320101132223312-3113123112213233-0120032210312300-3200212200020121)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)

<a id="canonical-2302133312113122-2103131230100032-2022320201011133-1233322313100303-1311332312100000-1300030220033203-2311003202110111-2110003001003333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3031203120331202-3332333133321300-1300222221230021-1013320101232011-0321311001201330-1320321111301032-0210322132000130-3332201200113003"></a>

## ingress_egress_gw_ar.hub.express_route_enabled.gateway_subnet.subnet.vnet_resource_group — vnet_resource_group / 003333033032 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-3012120331203201-1112000130332131-3202332120003001-0123013103333002-1112013103100001-2021323312110130-1112011111221321-2323233113231232)
- [ingress_egress_gw_ar](resources--azure_vnet_site--reference--group-005.md#canonical-0312012003033102-3210313212120203-0230331020120000-0030121220200120-2223003302123103-2232111103103011-2023112221000212-0212333313012311)
- [ingress_egress_gw_ar.hub](resources--azure_vnet_site--reference--group-005.md#canonical-2322310312311313-3121121220121201-2231223330120013-2030033320203223-3023310333201212-1221020313120000-1032200112132013-1301213311322012)
- [ingress_egress_gw_ar.hub.express_route_enabled](resources--azure_vnet_site--reference--group-006.md#canonical-2120301311230123-0230132233331312-1221203213113122-2231210212002002-1322021311302300-2201203011231232-1112123300032202-0231011112312022)
- [ingress_egress_gw_ar.hub.express_route_enabled.gateway_subnet](resources--azure_vnet_site--reference--group-006.md#canonical-3213100012220112-3132130321000300-2200303301211013-0231210003133130-3320101132223312-3113123112213233-0120032210312300-3200212200020121)
- [ingress_egress_gw_ar.hub.express_route_enabled.gateway_subnet.subnet](resources--azure_vnet_site--reference--group-006.md#canonical-2132320020301021-2213302323331120-1000012100310233-0311313003113331-0133233102301121-2102000200203332-0012223312121212-1121303121030222)
- ingress_egress_gw_ar.hub.express_route_enabled.gateway_subnet.subnet.vnet_resource_group

<a id="canonical-3132030220311022-1311120301113013-3333021200121332-0123022102322221-3011111020101112-2003132230320213-1232113222333131-0220021213000312"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for vnet resource group.

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
vnet_resource_group = {}
```

<a id="canonical-3220113112100203-1122323021031023-1031210122131220-0020110233232211-3002231103033122-0033320103303021-0010212313133111-1332321032233333"></a>

## Direct properties — vnet_resource_group / 003333033032 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2320123333203030-0112223111132130-1012200103321020-3020033201311033-3011311300230223-2110313320132213-0313202203200023-3113032012331300"></a>

## Next pages — vnet_resource_group / 003333033032 / 4

- [ingress_egress_gw_ar.hub.express_route_enabled.gateway_subnet.subnet](resources--azure_vnet_site--reference--group-006.md#canonical-2132320020301021-2213302323331120-1000012100310233-0311313003113331-0133233102301121-2102000200203332-0012223312121212-1121303121030222)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)

<a id="canonical-2132221021020220-1020302012003132-0003131003131010-1000022013333313-2200030021322212-2022320323321332-1311232013233313-3011033313031021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1131203212010001-0200300102031100-3003013321202102-3103132313013231-0221321222200120-2133021203220211-2103001333023232-1111102011033233"></a>

## ingress_egress_gw_ar.hub.express_route_enabled.gateway_subnet.subnet_param — subnet_param / 200321222132 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-3012120331203201-1112000130332131-3202332120003001-0123013103333002-1112013103100001-2021323312110130-1112011111221321-2323233113231232)
- [ingress_egress_gw_ar](resources--azure_vnet_site--reference--group-005.md#canonical-0312012003033102-3210313212120203-0230331020120000-0030121220200120-2223003302123103-2232111103103011-2023112221000212-0212333313012311)
- [ingress_egress_gw_ar.hub](resources--azure_vnet_site--reference--group-005.md#canonical-2322310312311313-3121121220121201-2231223330120013-2030033320203223-3023310333201212-1221020313120000-1032200112132013-1301213311322012)
- [ingress_egress_gw_ar.hub.express_route_enabled](resources--azure_vnet_site--reference--group-006.md#canonical-2120301311230123-0230132233331312-1221203213113122-2231210212002002-1322021311302300-2201203011231232-1112123300032202-0231011112312022)
- [ingress_egress_gw_ar.hub.express_route_enabled.gateway_subnet](resources--azure_vnet_site--reference--group-006.md#canonical-3213100012220112-3132130321000300-2200303301211013-0231210003133130-3320101132223312-3113123112213233-0120032210312300-3200212200020121)
- ingress_egress_gw_ar.hub.express_route_enabled.gateway_subnet.subnet_param

<a id="canonical-2112330331001220-0110131233131211-0201011032311003-1230302000331212-3312223023313102-3320023301303033-3111203111131202-3021033211323122"></a>

Type: `"object"`. single nested block, Optional.

Parameters for creating a new cloud subnet.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("ipv4")}
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
subnet_param {
  # Configure direct properties listed below.
}
```

<a id="canonical-0011230033221010-3330232023013321-3211200021123320-2330303302231123-2322332221233231-2213310210211221-1313121122111200-3210002120332303"></a>

## Direct properties — subnet_param / 200321222132 / 3

<a id="canonical-2122300030223130-2001000102012332-3232220212220202-0210002312132122-3131230013111333-1033303331113221-1313200002110021-2202221130000001"></a>

<a id="canonical-1200013300203123-1321333201322102-0022221033313313-0202302033310320-1230002131310300-3332203311323321-3300332203031203-3302222030311011"></a>

## IPv4 property — subnet_param / 200321222132 / 4

Type: `"string"`. Optional.

IPv4 Subnet. IPv4 subnet prefix for this subnet.

Upstream description:

IPv4 subnet prefix for this subnet.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "format": "ipv4",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
    "ves.io.schema.rules.string.ipv4_prefix": "true",
    "ves.io.schema.rules.string.max_ip_prefix_length": "28"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.ipv4_prefix": "true",
    "ves.io.schema.rules.string.max_ip_prefix_length": "28"
  }
}
```

<a id="canonical-1200220011311201-0133110003032232-2231023233330121-0121122320002331-2313321312202133-1131132300102322-2201012332131302-1300212333223023"></a>

## Next pages — subnet_param / 200321222132 / 5

- [ingress_egress_gw_ar.hub.express_route_enabled.gateway_subnet](resources--azure_vnet_site--reference--group-006.md#canonical-3213100012220112-3132130321000300-2200303301211013-0231210003133130-3320101132223312-3113123112213233-0120032210312300-3200212200020121)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)

<a id="canonical-3011112311323102-0020113321321123-1001313030011200-3012033203133312-2111223131023023-1310233212123020-0020332113130102-1233133222033031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3311323302112121-3311113233200121-1100302221200303-1223320332111012-1223102233231201-0131332320120031-3131010322031310-1230003012032012"></a>

## ingress_egress_gw_ar.hub.express_route_enabled.route_server_subnet — route_server_subnet / 010223230222 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-3012120331203201-1112000130332131-3202332120003001-0123013103333002-1112013103100001-2021323312110130-1112011111221321-2323233113231232)
- [ingress_egress_gw_ar](resources--azure_vnet_site--reference--group-005.md#canonical-0312012003033102-3210313212120203-0230331020120000-0030121220200120-2223003302123103-2232111103103011-2023112221000212-0212333313012311)
- [ingress_egress_gw_ar.hub](resources--azure_vnet_site--reference--group-005.md#canonical-2322310312311313-3121121220121201-2231223330120013-2030033320203223-3023310333201212-1221020313120000-1032200112132013-1301213311322012)
- [ingress_egress_gw_ar.hub.express_route_enabled](resources--azure_vnet_site--reference--group-006.md#canonical-2120301311230123-0230132233331312-1221203213113122-2231210212002002-1322021311302300-2201203011231232-1112123300032202-0231011112312022)
- ingress_egress_gw_ar.hub.express_route_enabled.route_server_subnet

<a id="canonical-3012232210220231-1320200110002303-0023210000203001-2220031033301211-3121322213221020-3031233111101020-3332022332200030-0330303211120102"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for route server subnet.

Upstream description:

Parameters for Azure subnet.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("auto",
    "subnet"),
  validators.ConflictingObjectAttributes("auto",
    "subnet_param"),
  validators.ConflictingObjectAttributes("subnet",
    "subnet_param")}
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
  "x-ves-oneof-field-choice": "[\"auto\",\"subnet\",\"subnet_param\"]"
}
```

Terraform syntax:

```terraform
route_server_subnet {
  # Configure direct properties listed below.
}
```

<a id="canonical-2200102120111202-0112212303012030-2010001330201332-2130221213112122-3233212313003200-0322200223032111-1303102132112333-1113031230001131"></a>

## Direct properties — route_server_subnet / 010223230222 / 3

- [auto](resources--azure_vnet_site--reference--group-006.md#canonical-3201332311112230-0202020022231120-3333133221103121-3303221011020030-0121202030300002-2133110110103200-3322033231330111-3232333203212200): complete subsection reference.

- [subnet](resources--azure_vnet_site--reference--group-006.md#canonical-2022221232103212-0230210321021132-3301031232111323-2322002021020132-2203203322110111-0223333312131010-1231023220032333-0220031302130302): complete subsection reference.

- [subnet_param](resources--azure_vnet_site--reference--group-006.md#canonical-0012302332002322-1132333113222111-1013231202010033-0122010120202302-1103030313133233-2201032323103102-1012230203013021-0312030323311232): complete subsection reference.

<a id="canonical-0232012220033002-0103103301203232-3022133023131322-0133132220103132-3010131323211002-0222011100221010-2003223003202002-0101033213222223"></a>

## Next pages — route_server_subnet / 010223230222 / 4

- [ingress_egress_gw_ar.hub.express_route_enabled.route_server_subnet.auto](resources--azure_vnet_site--reference--group-006.md#canonical-3201332311112230-0202020022231120-3333133221103121-3303221011020030-0121202030300002-2133110110103200-3322033231330111-3232333203212200)
- [ingress_egress_gw_ar.hub.express_route_enabled.route_server_subnet.subnet](resources--azure_vnet_site--reference--group-006.md#canonical-2022221232103212-0230210321021132-3301031232111323-2322002021020132-2203203322110111-0223333312131010-1231023220032333-0220031302130302)
- [ingress_egress_gw_ar.hub.express_route_enabled.route_server_subnet.subnet_param](resources--azure_vnet_site--reference--group-006.md#canonical-0012302332002322-1132333113222111-1013231202010033-0122010120202302-1103030313133233-2201032323103102-1012230203013021-0312030323311232)
- [ingress_egress_gw_ar.hub.express_route_enabled](resources--azure_vnet_site--reference--group-006.md#canonical-2120301311230123-0230132233331312-1221203213113122-2231210212002002-1322021311302300-2201203011231232-1112123300032202-0231011112312022)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)

<a id="canonical-3201332311112230-0202020022231120-3333133221103121-3303221011020030-0121202030300002-2133110110103200-3322033231330111-3232333203212200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1031120302032330-0233121312003000-0022133120230220-1201020010023113-0132030021230111-3311322022132323-1200112303110032-1011203303121221"></a>

## ingress_egress_gw_ar.hub.express_route_enabled.route_server_subnet.auto — auto / 110013201110 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-3012120331203201-1112000130332131-3202332120003001-0123013103333002-1112013103100001-2021323312110130-1112011111221321-2323233113231232)
- [ingress_egress_gw_ar](resources--azure_vnet_site--reference--group-005.md#canonical-0312012003033102-3210313212120203-0230331020120000-0030121220200120-2223003302123103-2232111103103011-2023112221000212-0212333313012311)
- [ingress_egress_gw_ar.hub](resources--azure_vnet_site--reference--group-005.md#canonical-2322310312311313-3121121220121201-2231223330120013-2030033320203223-3023310333201212-1221020313120000-1032200112132013-1301213311322012)
- [ingress_egress_gw_ar.hub.express_route_enabled](resources--azure_vnet_site--reference--group-006.md#canonical-2120301311230123-0230132233331312-1221203213113122-2231210212002002-1322021311302300-2201203011231232-1112123300032202-0231011112312022)
- [ingress_egress_gw_ar.hub.express_route_enabled.route_server_subnet](resources--azure_vnet_site--reference--group-006.md#canonical-3011112311323102-0020113321321123-1001313030011200-3012033203133312-2111223131023023-1310233212123020-0020332113130102-1233133222033031)
- ingress_egress_gw_ar.hub.express_route_enabled.route_server_subnet.auto

<a id="canonical-3231023222310222-2333200221210301-0333101321011313-0230122011211122-3002330003230002-0230300030103131-2302101130330111-1122120303211321"></a>

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
auto = {}
```

<a id="canonical-3332121222121102-2033122013332332-1230002202132133-2313123121201113-2312130013303112-1213323100213122-0233013022222130-3313110321121332"></a>

## Direct properties — auto / 110013201110 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0323312201101031-3322010301202233-1222000023122022-2010320222212233-1131003011332213-1112300320203021-3200321031321033-1210332221322302"></a>

## Next pages — auto / 110013201110 / 4

- [ingress_egress_gw_ar.hub.express_route_enabled.route_server_subnet](resources--azure_vnet_site--reference--group-006.md#canonical-3011112311323102-0020113321321123-1001313030011200-3012033203133312-2111223131023023-1310233212123020-0020332113130102-1233133222033031)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)

<a id="canonical-2022221232103212-0230210321021132-3301031232111323-2322002021020132-2203203322110111-0223333312131010-1231023220032333-0220031302130302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2230021102013133-2123312232012300-1002301021312313-0132200313103011-1131301331203322-1322033300022002-0232203312213302-1331332230223220"></a>

## ingress_egress_gw_ar.hub.express_route_enabled.route_server_subnet.subnet — subnet / 033303013031 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-3012120331203201-1112000130332131-3202332120003001-0123013103333002-1112013103100001-2021323312110130-1112011111221321-2323233113231232)
- [ingress_egress_gw_ar](resources--azure_vnet_site--reference--group-005.md#canonical-0312012003033102-3210313212120203-0230331020120000-0030121220200120-2223003302123103-2232111103103011-2023112221000212-0212333313012311)
- [ingress_egress_gw_ar.hub](resources--azure_vnet_site--reference--group-005.md#canonical-2322310312311313-3121121220121201-2231223330120013-2030033320203223-3023310333201212-1221020313120000-1032200112132013-1301213311322012)
- [ingress_egress_gw_ar.hub.express_route_enabled](resources--azure_vnet_site--reference--group-006.md#canonical-2120301311230123-0230132233331312-1221203213113122-2231210212002002-1322021311302300-2201203011231232-1112123300032202-0231011112312022)
- [ingress_egress_gw_ar.hub.express_route_enabled.route_server_subnet](resources--azure_vnet_site--reference--group-006.md#canonical-3011112311323102-0020113321321123-1001313030011200-3012033203133312-2111223131023023-1310233212123020-0020332113130102-1233133222033031)
- ingress_egress_gw_ar.hub.express_route_enabled.route_server_subnet.subnet

<a id="canonical-2030320222002213-1000210122233313-3321203301321323-0221131122221203-2231200120033131-1111002020103202-3113200002213011-0110003121301121"></a>

Type: `"object"`. single nested block, Optional.

Parameters for Azure special subnet which name is reserved. (i.e GatewaySubnet or
RouteServerSubnet).

Upstream description:

Parameters for Azure special subnet which name is reserved. (i.e GatewaySubnet or RouteServerSubnet)

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("subnet_resource_grp",
    "vnet_resource_group")}
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
  "x-ves-oneof-field-resource_group_choice": "[\"subnet_resource_grp\",\"vnet_resource_group\"]"
}
```

Terraform syntax:

```terraform
subnet {
  # Configure direct properties listed below.
}
```

<a id="canonical-1321220311222022-2120233212011222-3200011111222122-3023001003001031-0010032032020212-2312130323210301-0103003310321203-0300332202101322"></a>

## Direct properties — subnet / 033303013031 / 3

<a id="canonical-3010311020230322-2132221203002303-2132131233003002-2122002121232020-3302031120102012-0300021312120010-3132331213231013-1121310332030212"></a>

<a id="canonical-0010202221100233-2030332110133031-0212110122122200-2203121223320020-0300123020231210-1301212013321300-1020222200023130-0323020131332130"></a>

## subnet_resource_grp property — subnet / 033303013031 / 4

Type: `"string"`. Optional.

Exclusive with \[vnet\_resource\_group\] Specify name of Resource Group.

Upstream description:

Exclusive with \[vnet\_resource\_group\] Specify name of Resource Group.

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
    "ves.io.schema.rules.string.max_len": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64"
  }
}
```

- [vnet_resource_group](resources--azure_vnet_site--reference--group-006.md#canonical-3303222011003001-1132100123303230-1200222323010123-3231112031002211-1130131320301120-0131322032021220-3000302312232122-3000120032313322): complete subsection reference.

<a id="canonical-0023233213323102-1100030212010101-0003222230232222-0211102320213333-1222022203202233-1320120303030302-1200103123011222-2010011113323202"></a>

## Next pages — subnet / 033303013031 / 5

- [ingress_egress_gw_ar.hub.express_route_enabled.route_server_subnet.subnet.vnet_resource_group](resources--azure_vnet_site--reference--group-006.md#canonical-3303222011003001-1132100123303230-1200222323010123-3231112031002211-1130131320301120-0131322032021220-3000302312232122-3000120032313322)
- [ingress_egress_gw_ar.hub.express_route_enabled.route_server_subnet](resources--azure_vnet_site--reference--group-006.md#canonical-3011112311323102-0020113321321123-1001313030011200-3012033203133312-2111223131023023-1310233212123020-0020332113130102-1233133222033031)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)

<a id="canonical-3303222011003001-1132100123303230-1200222323010123-3231112031002211-1130131320301120-0131322032021220-3000302312232122-3000120032313322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1131020011131332-0330223313122200-3032132202123021-2012233233012120-3110120220102321-0312130302123002-0001212333010302-2210103013000011"></a>

## ingress_egress_gw_ar.hub.express_route_enabled.route_server_subnet.subnet.vnet_resource_group — vnet_resource_group / 130220322322 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-3012120331203201-1112000130332131-3202332120003001-0123013103333002-1112013103100001-2021323312110130-1112011111221321-2323233113231232)
- [ingress_egress_gw_ar](resources--azure_vnet_site--reference--group-005.md#canonical-0312012003033102-3210313212120203-0230331020120000-0030121220200120-2223003302123103-2232111103103011-2023112221000212-0212333313012311)
- [ingress_egress_gw_ar.hub](resources--azure_vnet_site--reference--group-005.md#canonical-2322310312311313-3121121220121201-2231223330120013-2030033320203223-3023310333201212-1221020313120000-1032200112132013-1301213311322012)
- [ingress_egress_gw_ar.hub.express_route_enabled](resources--azure_vnet_site--reference--group-006.md#canonical-2120301311230123-0230132233331312-1221203213113122-2231210212002002-1322021311302300-2201203011231232-1112123300032202-0231011112312022)
- [ingress_egress_gw_ar.hub.express_route_enabled.route_server_subnet](resources--azure_vnet_site--reference--group-006.md#canonical-3011112311323102-0020113321321123-1001313030011200-3012033203133312-2111223131023023-1310233212123020-0020332113130102-1233133222033031)
- [ingress_egress_gw_ar.hub.express_route_enabled.route_server_subnet.subnet](resources--azure_vnet_site--reference--group-006.md#canonical-2022221232103212-0230210321021132-3301031232111323-2322002021020132-2203203322110111-0223333312131010-1231023220032333-0220031302130302)
- ingress_egress_gw_ar.hub.express_route_enabled.route_server_subnet.subnet.vnet_resource_group

<a id="canonical-3313002300321133-1330122033023001-3013323100300132-0021332212010200-2120101301120021-3321322003221332-2133203121113033-0021311211211031"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for vnet resource group.

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
vnet_resource_group = {}
```

<a id="canonical-2032010332022132-3011131323300113-2003100330312120-0132212020212130-2231331320233302-3103310100300133-2221023113120323-2021331313011302"></a>

## Direct properties — vnet_resource_group / 130220322322 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0112333121021003-1022122322013222-0213221220101300-0220301010233320-1032211031320221-1033003123200032-0331322031021123-0010033303213203"></a>

## Next pages — vnet_resource_group / 130220322322 / 4

- [ingress_egress_gw_ar.hub.express_route_enabled.route_server_subnet.subnet](resources--azure_vnet_site--reference--group-006.md#canonical-2022221232103212-0230210321021132-3301031232111323-2322002021020132-2203203322110111-0223333312131010-1231023220032333-0220031302130302)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)

<a id="canonical-0012302332002322-1132333113222111-1013231202010033-0122010120202302-1103030313133233-2201032323103102-1012230203013021-0312030323311232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0322033123200001-0330033122022313-0012100302212310-0122003021302311-3020301301303021-2030233000100323-3010002301231213-2130213103131230"></a>

## ingress_egress_gw_ar.hub.express_route_enabled.route_server_subnet.subnet_param — subnet_param / 303111130201 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-3012120331203201-1112000130332131-3202332120003001-0123013103333002-1112013103100001-2021323312110130-1112011111221321-2323233113231232)
- [ingress_egress_gw_ar](resources--azure_vnet_site--reference--group-005.md#canonical-0312012003033102-3210313212120203-0230331020120000-0030121220200120-2223003302123103-2232111103103011-2023112221000212-0212333313012311)
- [ingress_egress_gw_ar.hub](resources--azure_vnet_site--reference--group-005.md#canonical-2322310312311313-3121121220121201-2231223330120013-2030033320203223-3023310333201212-1221020313120000-1032200112132013-1301213311322012)
- [ingress_egress_gw_ar.hub.express_route_enabled](resources--azure_vnet_site--reference--group-006.md#canonical-2120301311230123-0230132233331312-1221203213113122-2231210212002002-1322021311302300-2201203011231232-1112123300032202-0231011112312022)
- [ingress_egress_gw_ar.hub.express_route_enabled.route_server_subnet](resources--azure_vnet_site--reference--group-006.md#canonical-3011112311323102-0020113321321123-1001313030011200-3012033203133312-2111223131023023-1310233212123020-0020332113130102-1233133222033031)
- ingress_egress_gw_ar.hub.express_route_enabled.route_server_subnet.subnet_param

<a id="canonical-0021223000123112-0202333333103033-0011033022131332-1021300221211030-1002100111220023-1000122333202331-3023202021212032-3333212113221323"></a>

Type: `"object"`. single nested block, Optional.

Parameters for creating a new cloud subnet.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("ipv4")}
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
subnet_param {
  # Configure direct properties listed below.
}
```

<a id="canonical-3110103130312323-3110122021211100-1121323322300031-2333202112003310-2032211131002012-0112311213230112-0112002130100002-2000223022213031"></a>

## Direct properties — subnet_param / 303111130201 / 3

<a id="canonical-1311320333011310-3210321103022301-2021000220313110-3300330312211112-1102002100221222-3201211030320222-0220201003313301-3111331111323123"></a>

<a id="canonical-3322110323000320-1111031112033221-0201333100201222-2030212310003202-0032120000110201-2122013001102203-3202020222301222-1100121032031332"></a>

## IPv4 property — subnet_param / 303111130201 / 4

Type: `"string"`. Optional.

IPv4 Subnet. IPv4 subnet prefix for this subnet.

Upstream description:

IPv4 subnet prefix for this subnet.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "format": "ipv4",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
    "ves.io.schema.rules.string.ipv4_prefix": "true",
    "ves.io.schema.rules.string.max_ip_prefix_length": "28"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.ipv4_prefix": "true",
    "ves.io.schema.rules.string.max_ip_prefix_length": "28"
  }
}
```

<a id="canonical-0313021213103122-1032201312203031-3331221020332002-0220010013222212-1023013230222110-1232002212201230-0103303203113010-2133302211323212"></a>

## Next pages — subnet_param / 303111130201 / 5

- [ingress_egress_gw_ar.hub.express_route_enabled.route_server_subnet](resources--azure_vnet_site--reference--group-006.md#canonical-3011112311323102-0020113321321123-1001313030011200-3012033203133312-2111223131023023-1310233212123020-0020332113130102-1233133222033031)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)

<a id="canonical-0133201130033011-2110032111102032-1012313112111323-3032132132010310-0032213131123113-3123202230132110-0031111233020122-1003220012332110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3112231133333333-3012230213032322-3231002013333323-1111113232203013-3111130130113133-0122021033020133-2123301233310103-2230302000000122"></a>

## ingress_egress_gw_ar.hub.express_route_enabled.site_registration_over_express_route — site_registration_over_express_route / 220303001300 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-3012120331203201-1112000130332131-3202332120003001-0123013103333002-1112013103100001-2021323312110130-1112011111221321-2323233113231232)
- [ingress_egress_gw_ar](resources--azure_vnet_site--reference--group-005.md#canonical-0312012003033102-3210313212120203-0230331020120000-0030121220200120-2223003302123103-2232111103103011-2023112221000212-0212333313012311)
- [ingress_egress_gw_ar.hub](resources--azure_vnet_site--reference--group-005.md#canonical-2322310312311313-3121121220121201-2231223330120013-2030033320203223-3023310333201212-1221020313120000-1032200112132013-1301213311322012)
- [ingress_egress_gw_ar.hub.express_route_enabled](resources--azure_vnet_site--reference--group-006.md#canonical-2120301311230123-0230132233331312-1221203213113122-2231210212002002-1322021311302300-2201203011231232-1112123300032202-0231011112312022)
- ingress_egress_gw_ar.hub.express_route_enabled.site_registration_over_express_route

<a id="canonical-2003122221210002-3132122113030022-3332003230112111-2011100203200312-1232013121101323-2312221232131312-2232322312123000-1002201321103303"></a>

Type: `"object"`. single nested block, Optional.

CloudLink ADN Network Config.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("cloudlink_network_name")}
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
site_registration_over_express_route {
  # Configure direct properties listed below.
}
```

<a id="canonical-2301133021102020-3202002023220222-2321012311332022-1102112330333301-3113121010310211-2321113323202212-0010003332302012-3223030010211220"></a>

## Direct properties — site_registration_over_express_route / 220303001300 / 3

<a id="canonical-3201013233202032-1212013223221303-0203311001310231-2202002223330031-0333120033332222-3031320100232132-1320301202323322-1023130002220232"></a>

<a id="canonical-2321013303212032-2103100303230202-1302213030010002-0101031000120013-1300321032220001-1300300132203330-3020130310322230-0011202032330232"></a>

## cloudlink_network_name property — site_registration_over_express_route / 220303001300 / 4

Type: `"string"`. Optional.

Establish private connectivity with the F5 Distributed Cloud Global Network using a Private ADN
network. To provision a Private ADN network, please contact F5 Distributed Cloud support.

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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-2232322000103311-0022213202312133-0101012102212100-1032313220332320-2100000020220213-2211013233122111-0113020121211333-2323131202001001"></a>

## Next pages — site_registration_over_express_route / 220303001300 / 5

- [ingress_egress_gw_ar.hub.express_route_enabled](resources--azure_vnet_site--reference--group-006.md#canonical-2120301311230123-0230132233331312-1221203213113122-2231210212002002-1322021311302300-2201203011231232-1112123300032202-0231011112312022)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)

<a id="canonical-0333100230231012-1313010222233022-1031333211333320-2212203101030310-3133210201222203-2200012233121222-2311220000321221-3322132022030133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0032113302100022-0131332212111330-2200223201222203-1131011313010231-0312003132232101-1323001012303322-1023022103123230-1220103021110232"></a>

## ingress_egress_gw_ar.hub.express_route_enabled.site_registration_over_internet — site_registration_over_internet / 101111133320 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-3012120331203201-1112000130332131-3202332120003001-0123013103333002-1112013103100001-2021323312110130-1112011111221321-2323233113231232)
- [ingress_egress_gw_ar](resources--azure_vnet_site--reference--group-005.md#canonical-0312012003033102-3210313212120203-0230331020120000-0030121220200120-2223003302123103-2232111103103011-2023112221000212-0212333313012311)
- [ingress_egress_gw_ar.hub](resources--azure_vnet_site--reference--group-005.md#canonical-2322310312311313-3121121220121201-2231223330120013-2030033320203223-3023310333201212-1221020313120000-1032200112132013-1301213311322012)
- [ingress_egress_gw_ar.hub.express_route_enabled](resources--azure_vnet_site--reference--group-006.md#canonical-2120301311230123-0230132233331312-1221203213113122-2231210212002002-1322021311302300-2201203011231232-1112123300032202-0231011112312022)
- ingress_egress_gw_ar.hub.express_route_enabled.site_registration_over_internet

<a id="canonical-2122213110121302-2022120022221213-0120030323233120-3033202313101123-0031232301003230-2231032200323030-0121301101221303-2320111033223321"></a>

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
site_registration_over_internet = {}
```

<a id="canonical-3022013003220301-3331202212200130-3232320231113012-1331213100321003-1302121211120100-0121233230102032-2033223122220120-0230213332020112"></a>

## Direct properties — site_registration_over_internet / 101111133320 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1131231203220133-1331321023301322-3320123331321032-2131101123122111-0101101110123132-3102213300121331-0330210112231200-0202332232212111"></a>

## Next pages — site_registration_over_internet / 101111133320 / 4

- [ingress_egress_gw_ar.hub.express_route_enabled](resources--azure_vnet_site--reference--group-006.md#canonical-2120301311230123-0230132233331312-1221203213113122-2231210212002002-1322021311302300-2201203011231232-1112123300032202-0231011112312022)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)

<a id="canonical-1300011101023331-2113300100003220-0100302131323123-3030131232113322-2000132030102011-0313030202130232-1313001311210120-3023100021220321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0111112212002212-2112332100030201-0230131103102330-0200302013113031-2030020110211103-1213231022111130-0013032020011022-0333320102011010"></a>

## ingress_egress_gw_ar.hub.express_route_enabled.sku_ergw1az — sku_ergw1az / 112123302303 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-3012120331203201-1112000130332131-3202332120003001-0123013103333002-1112013103100001-2021323312110130-1112011111221321-2323233113231232)
- [ingress_egress_gw_ar](resources--azure_vnet_site--reference--group-005.md#canonical-0312012003033102-3210313212120203-0230331020120000-0030121220200120-2223003302123103-2232111103103011-2023112221000212-0212333313012311)
- [ingress_egress_gw_ar.hub](resources--azure_vnet_site--reference--group-005.md#canonical-2322310312311313-3121121220121201-2231223330120013-2030033320203223-3023310333201212-1221020313120000-1032200112132013-1301213311322012)
- [ingress_egress_gw_ar.hub.express_route_enabled](resources--azure_vnet_site--reference--group-006.md#canonical-2120301311230123-0230132233331312-1221203213113122-2231210212002002-1322021311302300-2201203011231232-1112123300032202-0231011112312022)
- ingress_egress_gw_ar.hub.express_route_enabled.sku_ergw1az

<a id="canonical-0201022330020013-2232322021231132-0223313223122302-0310203303201220-2333013120221202-1102031122320312-3320131020001310-2201012130223312"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for sku ergw1az.

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
sku_ergw1az = {}
```

<a id="canonical-1001312123311131-3332123212103220-1310200330310120-0321123132313013-1130012322003110-1122102131023013-1230323022102323-1003131211123131"></a>

## Direct properties — sku_ergw1az / 112123302303 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1010331323231132-2213032212200301-2321013013233033-2203122203132211-2130111311102230-3020203013320003-1213032323302333-0200111012103313"></a>

## Next pages — sku_ergw1az / 112123302303 / 4

- [ingress_egress_gw_ar.hub.express_route_enabled](resources--azure_vnet_site--reference--group-006.md#canonical-2120301311230123-0230132233331312-1221203213113122-2231210212002002-1322021311302300-2201203011231232-1112123300032202-0231011112312022)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)

<a id="canonical-2000033202113202-0130021211120213-3312320200331023-0322302332202000-0332232230313000-2110112221223213-1013331130001101-3003110002112002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0131001203012230-2220313013313003-1320301003002212-3021232113221331-3210300130300020-3113101112301333-2313112212131201-2100332121211320"></a>

## ingress_egress_gw_ar.hub.express_route_enabled.sku_ergw2az — sku_ergw2az / 003020133123 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-3012120331203201-1112000130332131-3202332120003001-0123013103333002-1112013103100001-2021323312110130-1112011111221321-2323233113231232)
- [ingress_egress_gw_ar](resources--azure_vnet_site--reference--group-005.md#canonical-0312012003033102-3210313212120203-0230331020120000-0030121220200120-2223003302123103-2232111103103011-2023112221000212-0212333313012311)
- [ingress_egress_gw_ar.hub](resources--azure_vnet_site--reference--group-005.md#canonical-2322310312311313-3121121220121201-2231223330120013-2030033320203223-3023310333201212-1221020313120000-1032200112132013-1301213311322012)
- [ingress_egress_gw_ar.hub.express_route_enabled](resources--azure_vnet_site--reference--group-006.md#canonical-2120301311230123-0230132233331312-1221203213113122-2231210212002002-1322021311302300-2201203011231232-1112123300032202-0231011112312022)
- ingress_egress_gw_ar.hub.express_route_enabled.sku_ergw2az

<a id="canonical-1211230320320023-2113123211012111-1230102233222112-3213222031333300-3231122003021032-2100122310210213-1232033100022013-2113301233112323"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for sku ergw2az.

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
sku_ergw2az = {}
```

<a id="canonical-1211323211000033-2300011310031101-2211331222232331-1101010032130331-0320020330332231-3322303031320233-3211132302232131-2331330201200321"></a>

## Direct properties — sku_ergw2az / 003020133123 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0123213112220102-0022310131031221-1032221311121002-0030113103332212-3321321033020232-2103000012230200-2231013213323011-1220103132202033"></a>

## Next pages — sku_ergw2az / 003020133123 / 4

- [ingress_egress_gw_ar.hub.express_route_enabled](resources--azure_vnet_site--reference--group-006.md#canonical-2120301311230123-0230132233331312-1221203213113122-2231210212002002-1322021311302300-2201203011231232-1112123300032202-0231011112312022)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)

<a id="canonical-1011002021003211-2013232003011230-3121330212102131-1221212122220002-1220230330222032-2300322211230103-1011330310313200-0131030132313020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1100003020321011-3033310122112223-3210201112300131-2303311230301230-2023033012122220-3302301111013320-3322300100103101-1230031123310320"></a>

## ingress_egress_gw_ar.hub.express_route_enabled.sku_high_perf — sku_high_perf / 330111020010 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-3012120331203201-1112000130332131-3202332120003001-0123013103333002-1112013103100001-2021323312110130-1112011111221321-2323233113231232)
- [ingress_egress_gw_ar](resources--azure_vnet_site--reference--group-005.md#canonical-0312012003033102-3210313212120203-0230331020120000-0030121220200120-2223003302123103-2232111103103011-2023112221000212-0212333313012311)
- [ingress_egress_gw_ar.hub](resources--azure_vnet_site--reference--group-005.md#canonical-2322310312311313-3121121220121201-2231223330120013-2030033320203223-3023310333201212-1221020313120000-1032200112132013-1301213311322012)
- [ingress_egress_gw_ar.hub.express_route_enabled](resources--azure_vnet_site--reference--group-006.md#canonical-2120301311230123-0230132233331312-1221203213113122-2231210212002002-1322021311302300-2201203011231232-1112123300032202-0231011112312022)
- ingress_egress_gw_ar.hub.express_route_enabled.sku_high_perf

<a id="canonical-1131011133220120-0012210121233220-0331230033333223-3122113333312312-3230322301230102-1330012011300013-0121320201022311-3121112010212030"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for sku high perf.

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
sku_high_perf = {}
```

<a id="canonical-3022000020033302-0103111113132211-3132013100211311-1210023123230300-0302122333200100-2033031222202312-2300200301013222-0130322302031300"></a>

## Direct properties — sku_high_perf / 330111020010 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1310312232212000-3131331301013300-2320030033222000-1010132330121101-1332013033320202-0032010200210031-0013330311003232-1233323312120113"></a>

## Next pages — sku_high_perf / 330111020010 / 4

- [ingress_egress_gw_ar.hub.express_route_enabled](resources--azure_vnet_site--reference--group-006.md#canonical-2120301311230123-0230132233331312-1221203213113122-2231210212002002-1322021311302300-2201203011231232-1112123300032202-0231011112312022)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)

<a id="canonical-0313103132221000-2312202121021330-1300303011001200-2121221233222122-2101133212022013-1232020120010110-3010232113230302-1222000133022122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2320231300011031-2021302010132231-1012211131313013-3010332302213131-2021111001311122-1233213202132233-0103301330111132-3223023131123312"></a>

## ingress_egress_gw_ar.hub.express_route_enabled.sku_standard — sku_standard / 311321023303 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-3012120331203201-1112000130332131-3202332120003001-0123013103333002-1112013103100001-2021323312110130-1112011111221321-2323233113231232)
- [ingress_egress_gw_ar](resources--azure_vnet_site--reference--group-005.md#canonical-0312012003033102-3210313212120203-0230331020120000-0030121220200120-2223003302123103-2232111103103011-2023112221000212-0212333313012311)
- [ingress_egress_gw_ar.hub](resources--azure_vnet_site--reference--group-005.md#canonical-2322310312311313-3121121220121201-2231223330120013-2030033320203223-3023310333201212-1221020313120000-1032200112132013-1301213311322012)
- [ingress_egress_gw_ar.hub.express_route_enabled](resources--azure_vnet_site--reference--group-006.md#canonical-2120301311230123-0230132233331312-1221203213113122-2231210212002002-1322021311302300-2201203011231232-1112123300032202-0231011112312022)
- ingress_egress_gw_ar.hub.express_route_enabled.sku_standard

<a id="canonical-0122120000211033-1312230310312100-0111020303311012-1122021332121012-0112211110220123-1312230011223220-1310301222333321-2023101023233222"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for sku standard.

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
sku_standard = {}
```

<a id="canonical-0220103120023310-1213131322222110-1111222111111331-2310021311113033-0130302003303212-0133202112002010-3310110301202233-3210031030313100"></a>

## Direct properties — sku_standard / 311321023303 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3012132133333121-0031113010011001-0332330000101231-1101223321112333-3303210313200011-0021303313000122-0110303202322113-3312101131301100"></a>

## Next pages — sku_standard / 311321023303 / 4

- [ingress_egress_gw_ar.hub.express_route_enabled](resources--azure_vnet_site--reference--group-006.md#canonical-2120301311230123-0230132233331312-1221203213113122-2231210212002002-1322021311302300-2201203011231232-1112123300032202-0231011112312022)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)

<a id="canonical-1003221131012101-2320323232111123-2221231310213210-0221032001213103-3201023212230231-3100210102113322-1112302331111002-0102011111130123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1020113033001102-2332110220121131-2111203210112103-2030312111120022-2110021320010323-1130012200011330-2113212303031330-0333111033313231"></a>

## ingress_egress_gw_ar.hub.spoke_vnets — spoke_vnets / 113113013031 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-3012120331203201-1112000130332131-3202332120003001-0123013103333002-1112013103100001-2021323312110130-1112011111221321-2323233113231232)
- [ingress_egress_gw_ar](resources--azure_vnet_site--reference--group-005.md#canonical-0312012003033102-3210313212120203-0230331020120000-0030121220200120-2223003302123103-2232111103103011-2023112221000212-0212333313012311)
- [ingress_egress_gw_ar.hub](resources--azure_vnet_site--reference--group-005.md#canonical-2322310312311313-3121121220121201-2231223330120013-2030033320203223-3023310333201212-1221020313120000-1032200112132013-1301213311322012)
- ingress_egress_gw_ar.hub.spoke_vnets

<a id="canonical-3202320130321130-2013310101000022-3011310223303330-0212102300200013-0313113030131132-1213001210322120-1132011123133332-3200123313231312"></a>

Type: `"object"`. list nested block, Optional.

Spoke VNet Peering (Legacy). Spoke VNet Peering.

Upstream description:

Spoke VNet Peering.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("auto",
    "manual")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 100,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 100,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
spoke_vnets {
  # Configure direct properties listed below.
}
```

<a id="canonical-0212001313200320-1333313203321003-0331223310113120-3133002300021111-3121320120102312-1322022302023023-3012001030123301-2221310211331333"></a>

## Direct properties — spoke_vnets / 113113013031 / 3

- [auto](resources--azure_vnet_site--reference--group-006.md#canonical-1202100223030132-0321011023032203-0300221021020232-0112003103130013-3231221233301231-3213022003112012-0320110210233331-1120311232022112): complete subsection reference.

- [labels](resources--azure_vnet_site--reference--group-006.md#canonical-2100222312312111-3330330012303203-0212131332120312-2120122301132132-0110322131121223-0112320031230031-3300300331021002-1002130120231123): complete subsection reference.

- [manual](resources--azure_vnet_site--reference--group-006.md#canonical-3221233230210310-3312133020303113-2230000322211102-2121302331230033-2012010110331110-2222330232130210-3221230333123022-0200222321231310): complete subsection reference.

- [vnet](resources--azure_vnet_site--reference--group-006.md#canonical-0321211301213312-3123000213003232-1012001312133100-1113100030130332-1110120230200011-3033023102303330-0333331003320112-2101200221302021): complete subsection reference.

<a id="canonical-3132023012032202-3113012102011101-1012020201003220-1231131002111200-0313012333112033-1111221223011310-3100220312212022-2101331131103031"></a>

## Next pages — spoke_vnets / 113113013031 / 4

- [ingress_egress_gw_ar.hub.spoke_vnets.auto](resources--azure_vnet_site--reference--group-006.md#canonical-1202100223030132-0321011023032203-0300221021020232-0112003103130013-3231221233301231-3213022003112012-0320110210233331-1120311232022112)
- [ingress_egress_gw_ar.hub.spoke_vnets.labels](resources--azure_vnet_site--reference--group-006.md#canonical-2100222312312111-3330330012303203-0212131332120312-2120122301132132-0110322131121223-0112320031230031-3300300331021002-1002130120231123)
- [ingress_egress_gw_ar.hub.spoke_vnets.manual](resources--azure_vnet_site--reference--group-006.md#canonical-3221233230210310-3312133020303113-2230000322211102-2121302331230033-2012010110331110-2222330232130210-3221230333123022-0200222321231310)
- [ingress_egress_gw_ar.hub.spoke_vnets.vnet](resources--azure_vnet_site--reference--group-006.md#canonical-0321211301213312-3123000213003232-1012001312133100-1113100030130332-1110120230200011-3033023102303330-0333331003320112-2101200221302021)
- [ingress_egress_gw_ar.hub](resources--azure_vnet_site--reference--group-005.md#canonical-2322310312311313-3121121220121201-2231223330120013-2030033320203223-3023310333201212-1221020313120000-1032200112132013-1301213311322012)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)

<a id="canonical-1202100223030132-0321011023032203-0300221021020232-0112003103130013-3231221233301231-3213022003112012-0320110210233331-1120311232022112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3103212033103022-1232002310222230-3223201102221303-0223133221033110-3323120020030301-2302131132210121-2321231031233210-3111321133333300"></a>

## ingress_egress_gw_ar.hub.spoke_vnets.auto — auto / 132103213012 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-3012120331203201-1112000130332131-3202332120003001-0123013103333002-1112013103100001-2021323312110130-1112011111221321-2323233113231232)
- [ingress_egress_gw_ar](resources--azure_vnet_site--reference--group-005.md#canonical-0312012003033102-3210313212120203-0230331020120000-0030121220200120-2223003302123103-2232111103103011-2023112221000212-0212333313012311)
- [ingress_egress_gw_ar.hub](resources--azure_vnet_site--reference--group-005.md#canonical-2322310312311313-3121121220121201-2231223330120013-2030033320203223-3023310333201212-1221020313120000-1032200112132013-1301213311322012)
- [ingress_egress_gw_ar.hub.spoke_vnets](resources--azure_vnet_site--reference--group-006.md#canonical-1003221131012101-2320323232111123-2221231310213210-0221032001213103-3201023212230231-3100210102113322-1112302331111002-0102011111130123)
- ingress_egress_gw_ar.hub.spoke_vnets.auto

<a id="canonical-0031022310002022-3213132221130331-3100023013230222-1133220313101313-3333233222113232-3020030020100232-0222120122313211-0031130320120202"></a>

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
auto = {}
```

<a id="canonical-1100233020310210-0210301332303221-1010131313313231-1222311312300110-1210102333222120-0103232200133123-2223222333032001-1022301203133202"></a>

## Direct properties — auto / 132103213012 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2100321323213201-0002233032133003-2302232012020121-0132020320232033-3031313220223103-2233202330012332-2303001023022200-2230131133210323"></a>

## Next pages — auto / 132103213012 / 4

- [ingress_egress_gw_ar.hub.spoke_vnets](resources--azure_vnet_site--reference--group-006.md#canonical-1003221131012101-2320323232111123-2221231310213210-0221032001213103-3201023212230231-3100210102113322-1112302331111002-0102011111130123)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)

<a id="canonical-2100222312312111-3330330012303203-0212131332120312-2120122301132132-0110322131121223-0112320031230031-3300300331021002-1002130120231123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1330332111023322-3233010132233130-3032233123003020-0322332320311211-2231333210220212-1231200100220002-3012300120001133-1113013233310113"></a>

## ingress_egress_gw_ar.hub.spoke_vnets.labels — labels / 202221110123 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-3012120331203201-1112000130332131-3202332120003001-0123013103333002-1112013103100001-2021323312110130-1112011111221321-2323233113231232)
- [ingress_egress_gw_ar](resources--azure_vnet_site--reference--group-005.md#canonical-0312012003033102-3210313212120203-0230331020120000-0030121220200120-2223003302123103-2232111103103011-2023112221000212-0212333313012311)
- [ingress_egress_gw_ar.hub](resources--azure_vnet_site--reference--group-005.md#canonical-2322310312311313-3121121220121201-2231223330120013-2030033320203223-3023310333201212-1221020313120000-1032200112132013-1301213311322012)
- [ingress_egress_gw_ar.hub.spoke_vnets](resources--azure_vnet_site--reference--group-006.md#canonical-1003221131012101-2320323232111123-2221231310213210-0221032001213103-3201023212230231-3100210102113322-1112302331111002-0102011111130123)
- ingress_egress_gw_ar.hub.spoke_vnets.labels

<a id="canonical-2203013121213133-0131320313122333-3210332111001322-3220002232303031-1311221103202102-2220103331210200-0321102230310020-3022223100321302"></a>

Type: `"object"`. single nested block, Optional.

Add Labels for each of the VNets peered with transit VNet, these labels can be used in firewall
policy These labels used must be from known key and label defined in shared namespace.

Upstream description:

Add Labels for each of the VNets peered with transit VNet, these labels can be used in firewall
policy These labels used must be from known key and label defined in shared namespace.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
labels {}
```

<a id="canonical-2013123311000310-2310033031132333-0111200121233302-2031121332332121-1210003322300130-3000201021031230-1233033332330013-1312210310030123"></a>

## Direct properties — labels / 202221110123 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3031211021000120-2333323020122211-2130130333222030-1220211132131012-2031021112322003-2120301032211030-1232202110232302-1332022023113222"></a>

## Next pages — labels / 202221110123 / 4

- [ingress_egress_gw_ar.hub.spoke_vnets](resources--azure_vnet_site--reference--group-006.md#canonical-1003221131012101-2320323232111123-2221231310213210-0221032001213103-3201023212230231-3100210102113322-1112302331111002-0102011111130123)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)

<a id="canonical-3221233230210310-3312133020303113-2230000322211102-2121302331230033-2012010110331110-2222330232130210-3221230333123022-0200222321231310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1223033013211202-0123130100120101-2101210132122331-1101002020211001-0222103211032311-0111212201003303-0313001113213200-0232131211031001"></a>

## ingress_egress_gw_ar.hub.spoke_vnets.manual — manual / 123121033323 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-3012120331203201-1112000130332131-3202332120003001-0123013103333002-1112013103100001-2021323312110130-1112011111221321-2323233113231232)
- [ingress_egress_gw_ar](resources--azure_vnet_site--reference--group-005.md#canonical-0312012003033102-3210313212120203-0230331020120000-0030121220200120-2223003302123103-2232111103103011-2023112221000212-0212333313012311)
- [ingress_egress_gw_ar.hub](resources--azure_vnet_site--reference--group-005.md#canonical-2322310312311313-3121121220121201-2231223330120013-2030033320203223-3023310333201212-1221020313120000-1032200112132013-1301213311322012)
- [ingress_egress_gw_ar.hub.spoke_vnets](resources--azure_vnet_site--reference--group-006.md#canonical-1003221131012101-2320323232111123-2221231310213210-0221032001213103-3201023212230231-3100210102113322-1112302331111002-0102011111130123)
- ingress_egress_gw_ar.hub.spoke_vnets.manual

<a id="canonical-2123331221203112-3120022113313312-0023210100210021-3130130030112302-0013313212313231-2101311300001311-2113100223030320-0332220013021233"></a>

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
manual = {}
```

<a id="canonical-2201201320201302-1020113313230222-0000022301332001-2010100000330033-0123022203223033-2133330313303323-3201312332301133-0201210331333232"></a>

## Direct properties — manual / 123121033323 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0022210203121213-0103023110312013-1032031011113011-1002331202111323-3330330032201232-0103200313020123-0133210231112320-2231112012233113"></a>

## Next pages — manual / 123121033323 / 4

- [ingress_egress_gw_ar.hub.spoke_vnets](resources--azure_vnet_site--reference--group-006.md#canonical-1003221131012101-2320323232111123-2221231310213210-0221032001213103-3201023212230231-3100210102113322-1112302331111002-0102011111130123)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)

<a id="canonical-0321211301213312-3123000213003232-1012001312133100-1113100030130332-1110120230200011-3033023102303330-0333331003320112-2101200221302021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0130302003120221-0201220323122213-3233111122131111-0222023000111110-0331101013002023-1323112212103130-0101030233131200-3122213002013212"></a>

## ingress_egress_gw_ar.hub.spoke_vnets.vnet — vnet / 022002021132 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-3012120331203201-1112000130332131-3202332120003001-0123013103333002-1112013103100001-2021323312110130-1112011111221321-2323233113231232)
- [ingress_egress_gw_ar](resources--azure_vnet_site--reference--group-005.md#canonical-0312012003033102-3210313212120203-0230331020120000-0030121220200120-2223003302123103-2232111103103011-2023112221000212-0212333313012311)
- [ingress_egress_gw_ar.hub](resources--azure_vnet_site--reference--group-005.md#canonical-2322310312311313-3121121220121201-2231223330120013-2030033320203223-3023310333201212-1221020313120000-1032200112132013-1301213311322012)
- [ingress_egress_gw_ar.hub.spoke_vnets](resources--azure_vnet_site--reference--group-006.md#canonical-1003221131012101-2320323232111123-2221231310213210-0221032001213103-3201023212230231-3100210102113322-1112302331111002-0102011111130123)
- ingress_egress_gw_ar.hub.spoke_vnets.vnet

<a id="canonical-2301032013100101-2311222130333331-3101230210012023-2010302203033113-0122133223120223-0010211223132001-2120302300002013-1300231103221330"></a>

Type: `"object"`. single nested block, Optional.

Resource group and name of existing Azure VNet.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("resource_group",
    "vnet_name"),
  validators.ConflictingObjectAttributes("f5_orchestrated_routing",
    "manual_routing")}
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
  "x-ves-oneof-field-routing_type": "[\"f5_orchestrated_routing\",\"manual_routing\"]"
}
```

Terraform syntax:

```terraform
vnet {
  # Configure direct properties listed below.
}
```

<a id="canonical-1221020033220133-1030311010222210-1012333023213022-3101023122121332-0033221012300230-1301122231202323-2130303100022220-0001313233231130"></a>

## Direct properties — vnet / 022002021132 / 3

- [f5_orchestrated_routing](resources--azure_vnet_site--reference--group-006.md#canonical-1233021211012202-1002121232012302-2223121130102003-1303320030320330-2232220111312023-1311302112131030-1333332023222002-1003110001012233): complete subsection reference.

- [manual_routing](resources--azure_vnet_site--reference--group-006.md#canonical-2322221023101002-3122220231301032-2231220102100210-1020133331212312-1202103331301010-3212131332123122-3311211110210330-0320002302132331): complete subsection reference.

<a id="canonical-0320122022303232-2212223020111201-1032020102123021-0013330131131112-3111110113330103-2201002322211201-3001223010200013-1022333211203003"></a>

<a id="canonical-0222122323212121-2112113301310001-1333310021313022-0313021020331020-3031022211211231-2101110022233332-2012031022323133-2313210213132312"></a>

## resource_group property — vnet / 022002021132 / 4

Type: `"string"`. Optional.

Existing VNet Resource Group. Resource group of existing VNet.

Upstream description:

Resource group of existing VNet.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 64),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
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
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-0002021202321301-0020212110323213-0130320131231033-0233011310232333-3133033000302132-2321211333021011-1230031103121223-0330110321013120"></a>

<a id="canonical-1212123010130021-1112131100231003-3131230020101203-1221031110301103-2332020203232020-0303331333122330-3030001232311112-2332303113123222"></a>

## vnet_name property — vnet / 022002021132 / 5

Type: `"string"`. Optional.

Existing VNet Name. Name of existing VNet.

Upstream description:

Name of existing VNet.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 64),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
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
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-2032103213132130-0010023311202321-0221211031012312-2021111130102121-0110231010301331-2121033100100123-3230133010023322-3223221103123200"></a>

## Next pages — vnet / 022002021132 / 6

- [ingress_egress_gw_ar.hub.spoke_vnets.vnet.f5_orchestrated_routing](resources--azure_vnet_site--reference--group-006.md#canonical-1233021211012202-1002121232012302-2223121130102003-1303320030320330-2232220111312023-1311302112131030-1333332023222002-1003110001012233)
- [ingress_egress_gw_ar.hub.spoke_vnets.vnet.manual_routing](resources--azure_vnet_site--reference--group-006.md#canonical-2322221023101002-3122220231301032-2231220102100210-1020133331212312-1202103331301010-3212131332123122-3311211110210330-0320002302132331)
- [ingress_egress_gw_ar.hub.spoke_vnets](resources--azure_vnet_site--reference--group-006.md#canonical-1003221131012101-2320323232111123-2221231310213210-0221032001213103-3201023212230231-3100210102113322-1112302331111002-0102011111130123)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)

<a id="canonical-1233021211012202-1002121232012302-2223121130102003-1303320030320330-2232220111312023-1311302112131030-1333332023222002-1003110001012233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3033100110311010-1121033132320302-0312302310030233-0311223002223101-3322321231331310-1303110013312211-2333023212032133-0332202202111012"></a>

## ingress_egress_gw_ar.hub.spoke_vnets.vnet.f5_orchestrated_routing — f5_orchestrated_routing / 333010301213 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-3012120331203201-1112000130332131-3202332120003001-0123013103333002-1112013103100001-2021323312110130-1112011111221321-2323233113231232)
- [ingress_egress_gw_ar](resources--azure_vnet_site--reference--group-005.md#canonical-0312012003033102-3210313212120203-0230331020120000-0030121220200120-2223003302123103-2232111103103011-2023112221000212-0212333313012311)
- [ingress_egress_gw_ar.hub](resources--azure_vnet_site--reference--group-005.md#canonical-2322310312311313-3121121220121201-2231223330120013-2030033320203223-3023310333201212-1221020313120000-1032200112132013-1301213311322012)
- [ingress_egress_gw_ar.hub.spoke_vnets](resources--azure_vnet_site--reference--group-006.md#canonical-1003221131012101-2320323232111123-2221231310213210-0221032001213103-3201023212230231-3100210102113322-1112302331111002-0102011111130123)
- [ingress_egress_gw_ar.hub.spoke_vnets.vnet](resources--azure_vnet_site--reference--group-006.md#canonical-0321211301213312-3123000213003232-1012001312133100-1113100030130332-1110120230200011-3033023102303330-0333331003320112-2101200221302021)
- ingress_egress_gw_ar.hub.spoke_vnets.vnet.f5_orchestrated_routing

<a id="canonical-1123001120030110-2312311210332033-0222213212130133-1000100113100002-2322231312010111-2032300133132203-3130013310021102-1003320233120023"></a>

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
f5_orchestrated_routing = {}
```

<a id="canonical-1012112122320223-1021031030113003-2123030121123313-0302030012000302-1320331001122301-3300230223122111-3033131201011222-0022120333203132"></a>

## Direct properties — f5_orchestrated_routing / 333010301213 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0201333301203332-0030130133102231-3111300310133220-1112123323233313-3111023011233011-1133213232110322-0213012301110010-2112110232303310"></a>

## Next pages — f5_orchestrated_routing / 333010301213 / 4

- [ingress_egress_gw_ar.hub.spoke_vnets.vnet](resources--azure_vnet_site--reference--group-006.md#canonical-0321211301213312-3123000213003232-1012001312133100-1113100030130332-1110120230200011-3033023102303330-0333331003320112-2101200221302021)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)

<a id="canonical-2322221023101002-3122220231301032-2231220102100210-1020133331212312-1202103331301010-3212131332123122-3311211110210330-0320002302132331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3202130321121331-2231232201032220-2201233210130321-0123131310120020-3300033222311320-2310323031212111-3310201003131200-0012131201320321"></a>

## ingress_egress_gw_ar.hub.spoke_vnets.vnet.manual_routing — manual_routing / 322003120110 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-3012120331203201-1112000130332131-3202332120003001-0123013103333002-1112013103100001-2021323312110130-1112011111221321-2323233113231232)
- [ingress_egress_gw_ar](resources--azure_vnet_site--reference--group-005.md#canonical-0312012003033102-3210313212120203-0230331020120000-0030121220200120-2223003302123103-2232111103103011-2023112221000212-0212333313012311)
- [ingress_egress_gw_ar.hub](resources--azure_vnet_site--reference--group-005.md#canonical-2322310312311313-3121121220121201-2231223330120013-2030033320203223-3023310333201212-1221020313120000-1032200112132013-1301213311322012)
- [ingress_egress_gw_ar.hub.spoke_vnets](resources--azure_vnet_site--reference--group-006.md#canonical-1003221131012101-2320323232111123-2221231310213210-0221032001213103-3201023212230231-3100210102113322-1112302331111002-0102011111130123)
- [ingress_egress_gw_ar.hub.spoke_vnets.vnet](resources--azure_vnet_site--reference--group-006.md#canonical-0321211301213312-3123000213003232-1012001312133100-1113100030130332-1110120230200011-3033023102303330-0333331003320112-2101200221302021)
- ingress_egress_gw_ar.hub.spoke_vnets.vnet.manual_routing

<a id="canonical-3200310222103210-2013202332130212-0332333323321000-3311023232300310-0301121101112211-2111020010022132-0033020312232230-0220220222033321"></a>

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
manual_routing = {}
```

<a id="canonical-0013101330323100-2201200132220012-0212332331002202-2131323203321000-0012101121220303-3023132321021023-0300221100311320-3210311213022130"></a>

## Direct properties — manual_routing / 322003120110 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0011123222033002-2213220331000322-2303010032132311-0123331122221203-1200331300100320-0023032221133200-1311333231232322-1310010220201231"></a>

## Next pages — manual_routing / 322003120110 / 4

- [ingress_egress_gw_ar.hub.spoke_vnets.vnet](resources--azure_vnet_site--reference--group-006.md#canonical-0321211301213312-3123000213003232-1012001312133100-1113100030130332-1110120230200011-3033023102303330-0333331003320112-2101200221302021)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)

<a id="canonical-2003312233332113-0111122021131232-0103020222220000-1131133231223012-3330001301220131-1222310013113030-1013300112132003-0330213300002202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3121300113113231-0321133021030133-3111330103213103-2000222021013201-1232122110021131-1221012122331233-3012100231332121-0100311310322301"></a>

## ingress_egress_gw_ar.inside_static_routes — inside_static_routes / 331213211020 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-3012120331203201-1112000130332131-3202332120003001-0123013103333002-1112013103100001-2021323312110130-1112011111221321-2323233113231232)
- [ingress_egress_gw_ar](resources--azure_vnet_site--reference--group-005.md#canonical-0312012003033102-3210313212120203-0230331020120000-0030121220200120-2223003302123103-2232111103103011-2023112221000212-0212333313012311)
- ingress_egress_gw_ar.inside_static_routes

<a id="canonical-0110032200312013-1332021111122011-2111232111213101-2031231210020232-3321202200121333-1231323333032123-1122013032002210-2113332020021333"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for inside static routes.

Upstream description:

List of static routes.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("static_route_list")}
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
inside_static_routes {
  # Configure direct properties listed below.
}
```

<a id="canonical-3321313003200020-0011000133210130-3213132323132323-0132300213102231-2032320123331022-0303310111301013-0313213211323020-1200332323011130"></a>

## Direct properties — inside_static_routes / 331213211020 / 3

- [static_route_list](resources--azure_vnet_site--reference--group-006.md#canonical-2323231232330320-3210301011110220-1130100131330230-2113013022002303-1123030101333321-1331230102100131-3002210131120221-3010111311031010): complete subsection reference.

<a id="canonical-0023222132132110-1010233330012200-2033221022232023-1011131231122101-3212300311311033-0132212323111121-2311123303013120-1223002320001212"></a>

## Next pages — inside_static_routes / 331213211020 / 4

- [ingress_egress_gw_ar.inside_static_routes.static_route_list](resources--azure_vnet_site--reference--group-006.md#canonical-2323231232330320-3210301011110220-1130100131330230-2113013022002303-1123030101333321-1331230102100131-3002210131120221-3010111311031010)
- [ingress_egress_gw_ar](resources--azure_vnet_site--reference--group-005.md#canonical-0312012003033102-3210313212120203-0230331020120000-0030121220200120-2223003302123103-2232111103103011-2023112221000212-0212333313012311)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)

<a id="canonical-2323231232330320-3210301011110220-1130100131330230-2113013022002303-1123030101333321-1331230102100131-3002210131120221-3010111311031010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2321030200202123-1102131201121210-3001003221322303-0013003233120031-3010313130130133-2310210132202212-3032330302111203-0330010120120120"></a>

## ingress_egress_gw_ar.inside_static_routes.static_route_list — static_route_list / 011100302212 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-3012120331203201-1112000130332131-3202332120003001-0123013103333002-1112013103100001-2021323312110130-1112011111221321-2323233113231232)
- [ingress_egress_gw_ar](resources--azure_vnet_site--reference--group-005.md#canonical-0312012003033102-3210313212120203-0230331020120000-0030121220200120-2223003302123103-2232111103103011-2023112221000212-0212333313012311)
- [ingress_egress_gw_ar.inside_static_routes](resources--azure_vnet_site--reference--group-006.md#canonical-2003312233332113-0111122021131232-0103020222220000-1131133231223012-3330001301220131-1222310013113030-1013300112132003-0330213300002202)
- ingress_egress_gw_ar.inside_static_routes.static_route_list

<a id="canonical-3220003211323002-2301333311201211-0101231031323332-2300311212211010-0111332303123003-1033022233231112-1101230102200232-1032310120320012"></a>

Type: `"object"`. list nested block, Optional.

List of Static Routes. List of Static routes.

Upstream description:

List of Static routes.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("custom_static_route",
    "simple_static_route")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 64,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 64,
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
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

Terraform syntax:

```terraform
static_route_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-2233203212233013-3210003113311123-2003230100103020-1203333101302201-3320112110210232-1330232203322330-0002001013001310-3323010001332121"></a>

## Direct properties — static_route_list / 011100302212 / 3

- [custom_static_route](resources--azure_vnet_site--reference--group-006.md#canonical-2222001221311112-1030220102011022-1210123221302232-2311022330031311-2102131031323330-0030302301313301-2121330312122122-2133123231233022): complete subsection reference.

<a id="canonical-2201032231202203-1022212003123003-2113231013011320-2001310132302201-0103021330030032-0130223213213223-3233211000223002-2232013313333330"></a>

<a id="canonical-0313311311011021-2323113012003032-0103011001203000-0013110212033232-0001232311031321-0323322012130323-2132010221130323-0333002122011312"></a>

## simple_static_route property — static_route_list / 011100302212 / 4

Type: `"string"`. Optional.

Exclusive with \[custom\_static\_route\] Use simple static route for prefix pointing to single
interface in the network.

Upstream description:

Exclusive with \[custom\_static\_route\] Use simple static route for prefix pointing to single
interface in the network.

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
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.ipv4_prefix": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv4_prefix": "true"
  }
}
```

<a id="canonical-0220000311222111-1203321320313023-0300323213130230-1232313110311231-3223111331311033-3210012210012223-0333312012331001-0120301320212012"></a>

## Next pages — static_route_list / 011100302212 / 5

- [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route](resources--azure_vnet_site--reference--group-006.md#canonical-2222001221311112-1030220102011022-1210123221302232-2311022330031311-2102131031323330-0030302301313301-2121330312122122-2133123231233022)
- [ingress_egress_gw_ar.inside_static_routes](resources--azure_vnet_site--reference--group-006.md#canonical-2003312233332113-0111122021131232-0103020222220000-1131133231223012-3330001301220131-1222310013113030-1013300112132003-0330213300002202)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)

<a id="canonical-2222001221311112-1030220102011022-1210123221302232-2311022330031311-2102131031323330-0030302301313301-2121330312122122-2133123231233022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1313221303320121-1012311111032023-2003003323103330-2233312332212310-2001312133303012-1120032123312210-1231223000330203-1113120122030333"></a>

## ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route — custom_static_route / 133102201233 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-3012120331203201-1112000130332131-3202332120003001-0123013103333002-1112013103100001-2021323312110130-1112011111221321-2323233113231232)
- [ingress_egress_gw_ar](resources--azure_vnet_site--reference--group-005.md#canonical-0312012003033102-3210313212120203-0230331020120000-0030121220200120-2223003302123103-2232111103103011-2023112221000212-0212333313012311)
- [ingress_egress_gw_ar.inside_static_routes](resources--azure_vnet_site--reference--group-006.md#canonical-2003312233332113-0111122021131232-0103020222220000-1131133231223012-3330001301220131-1222310013113030-1013300112132003-0330213300002202)
- [ingress_egress_gw_ar.inside_static_routes.static_route_list](resources--azure_vnet_site--reference--group-006.md#canonical-2323231232330320-3210301011110220-1130100131330230-2113013022002303-1123030101333321-1331230102100131-3002210131120221-3010111311031010)
- ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route

<a id="canonical-1022100223300302-3232121222011122-1022300321032030-1213032322112023-0303203103232121-3310011031013130-2321322103210112-1210131210123011"></a>

Type: `"object"`. single nested block, Optional.

Defines a static route, configuring a list of prefixes and a next-hop to be used for them.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("subnets")}
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
custom_static_route {
  # Configure direct properties listed below.
}
```

<a id="canonical-1333302221202220-3033201100103121-2033231022312301-1102213000002111-0311010313110311-3001101232103203-1303323000031203-3132333323211101"></a>

## Direct properties — custom_static_route / 133102201233 / 3

<a id="canonical-0000211120322122-0210313320223131-3323220221000313-3330311100032223-2223103300300301-3201312232030110-2010222310201301-3233111302323232"></a>

<a id="canonical-2302222201301032-0012230012232301-2133223010032222-3330311130302102-3033032122132110-1010222120101211-0133310031323202-1030203302033220"></a>

## attrs property — custom_static_route / 133102201233 / 4

Type: `["list", "string"]`. Optional.

\[Enum:
ROUTE\_ATTR\_NO\_OP|ROUTE\_ATTR\_ADVERTISE|ROUTE\_ATTR\_INSTALL\_HOST|ROUTE\_ATTR\_INSTALL\_FORWARDING|ROUTE\_ATTR\_MERGE\_ONLY\]
List of route attributes associated with the static route. Possible values are
\`ROUTE\_ATTR\_NO\_OP\`, \`ROUTE\_ATTR\_ADVERTISE\`, \`ROUTE\_ATTR\_INSTALL\_HOST\`,
\`ROUTE\_ATTR\_INSTALL\_FORWARDING\`, \`ROUTE\_ATTR\_MERGE\_ONLY\`. Defaults to
\`ROUTE\_ATTR\_NO\_OP\`.

Upstream description:

List of route attributes associated with the static route.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(4),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 4,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 4,
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
    "ves.io.schema.rules.repeated.max_items": "4"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "4"
  }
}
```

- [labels](resources--azure_vnet_site--reference--group-006.md#canonical-0232031031011001-0331221312132113-0203011323123313-0323231032032212-1233301233232222-0112013302322323-2103002031023022-0222203322211000): complete subsection reference.

- [nexthop](resources--azure_vnet_site--reference--group-006.md#canonical-0302121313113312-3323313023020010-3032301101221202-3210120223011032-2311022102303022-3322330023021203-1013222221302132-2110003320121013): complete subsection reference.

- [subnets](resources--azure_vnet_site--reference--group-006.md#canonical-1331313323323133-2232110222031313-3303122302120020-2121113000333222-1203300200213111-2012310222120112-0021020033323032-0223130310030220): complete subsection reference.

<a id="canonical-2031121303322002-1220002111020000-1223320200320010-2220203120313213-1203322103222011-3000302022133002-1313212322002030-3322132302310210"></a>

## Next pages — custom_static_route / 133102201233 / 5

- [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.labels](resources--azure_vnet_site--reference--group-006.md#canonical-0232031031011001-0331221312132113-0203011323123313-0323231032032212-1233301233232222-0112013302322323-2103002031023022-0222203322211000)
- [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop](resources--azure_vnet_site--reference--group-006.md#canonical-0302121313113312-3323313023020010-3032301101221202-3210120223011032-2311022102303022-3322330023021203-1013222221302132-2110003320121013)
- [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.subnets](resources--azure_vnet_site--reference--group-006.md#canonical-1331313323323133-2232110222031313-3303122302120020-2121113000333222-1203300200213111-2012310222120112-0021020033323032-0223130310030220)
- [ingress_egress_gw_ar.inside_static_routes.static_route_list](resources--azure_vnet_site--reference--group-006.md#canonical-2323231232330320-3210301011110220-1130100131330230-2113013022002303-1123030101333321-1331230102100131-3002210131120221-3010111311031010)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)

<a id="canonical-0232031031011001-0331221312132113-0203011323123313-0323231032032212-1233301233232222-0112013302322323-2103002031023022-0222203322211000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1030311200311111-1213131110312223-2302211010310022-0311210211130233-1120310022032103-1003100320100320-1001120130202201-3200033313031232"></a>

## ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.labels — labels / 210232232300 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-3012120331203201-1112000130332131-3202332120003001-0123013103333002-1112013103100001-2021323312110130-1112011111221321-2323233113231232)
- [ingress_egress_gw_ar](resources--azure_vnet_site--reference--group-005.md#canonical-0312012003033102-3210313212120203-0230331020120000-0030121220200120-2223003302123103-2232111103103011-2023112221000212-0212333313012311)
- [ingress_egress_gw_ar.inside_static_routes](resources--azure_vnet_site--reference--group-006.md#canonical-2003312233332113-0111122021131232-0103020222220000-1131133231223012-3330001301220131-1222310013113030-1013300112132003-0330213300002202)
- [ingress_egress_gw_ar.inside_static_routes.static_route_list](resources--azure_vnet_site--reference--group-006.md#canonical-2323231232330320-3210301011110220-1130100131330230-2113013022002303-1123030101333321-1331230102100131-3002210131120221-3010111311031010)
- [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route](resources--azure_vnet_site--reference--group-006.md#canonical-2222001221311112-1030220102011022-1210123221302232-2311022330031311-2102131031323330-0030302301313301-2121330312122122-2133123231233022)
- ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.labels

<a id="canonical-0320233011012131-2100022013122121-3232322001120201-0312121201003021-0013012300233122-2201100031010210-1020203101112212-3120110111132230"></a>

Type: `"object"`. single nested block, Optional.

Add Labels for this Static Route, these labels can be used in network policy.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
labels {}
```

<a id="canonical-0220123132021133-2122220102120202-1200012230023333-1202321121233121-0132033203022000-3310320310232213-1333321210123231-3333111000212231"></a>

## Direct properties — labels / 210232232300 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0132200010321221-3132322113132001-3310102311031023-2022203131231321-0012033132031332-2330201022212122-2010212222302202-0031203323311323"></a>

## Next pages — labels / 210232232300 / 4

- [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route](resources--azure_vnet_site--reference--group-006.md#canonical-2222001221311112-1030220102011022-1210123221302232-2311022330031311-2102131031323330-0030302301313301-2121330312122122-2133123231233022)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)

<a id="canonical-0302121313113312-3323313023020010-3032301101221202-3210120223011032-2311022102303022-3322330023021203-1013222221302132-2110003320121013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1210133302210223-2131202211121020-2133113033102010-3031011121332203-1302302111231003-1121103121130330-2112330102313022-3122121003113033"></a>

## ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop — nexthop / 102230201213 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-3012120331203201-1112000130332131-3202332120003001-0123013103333002-1112013103100001-2021323312110130-1112011111221321-2323233113231232)
- [ingress_egress_gw_ar](resources--azure_vnet_site--reference--group-005.md#canonical-0312012003033102-3210313212120203-0230331020120000-0030121220200120-2223003302123103-2232111103103011-2023112221000212-0212333313012311)
- [ingress_egress_gw_ar.inside_static_routes](resources--azure_vnet_site--reference--group-006.md#canonical-2003312233332113-0111122021131232-0103020222220000-1131133231223012-3330001301220131-1222310013113030-1013300112132003-0330213300002202)
- [ingress_egress_gw_ar.inside_static_routes.static_route_list](resources--azure_vnet_site--reference--group-006.md#canonical-2323231232330320-3210301011110220-1130100131330230-2113013022002303-1123030101333321-1331230102100131-3002210131120221-3010111311031010)
- [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route](resources--azure_vnet_site--reference--group-006.md#canonical-2222001221311112-1030220102011022-1210123221302232-2311022330031311-2102131031323330-0030302301313301-2121330312122122-2133123231233022)
- ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop

<a id="canonical-0033330300331211-3013121330020102-0202223300111022-3110031031130213-1203123211223312-1133111333211212-0033333301223321-3101201301332002"></a>

Type: `"object"`. single nested block, Optional.

Nexthop. Identifies the next-hop for a route.

Upstream description:

Identifies the next-hop for a route.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
nexthop {
  # Configure direct properties listed below.
}
```

<a id="canonical-3202232011200312-0020021112013022-0103111003333311-2002030113133331-3330312003300223-3201331211001321-3021220101133222-0202333320211100"></a>

## Direct properties — nexthop / 102230201213 / 3

- [interface](resources--azure_vnet_site--reference--group-006.md#canonical-1101033021223122-2120100010030203-2022230113103332-0322110301022301-0233333031222310-3103302010130020-1202133300100212-3120202220122221): complete subsection reference.

- [nexthop_address](resources--azure_vnet_site--reference--group-006.md#canonical-0203210121300221-3132113002113110-1022021220213313-3302221233133232-3320322013233011-0112300223011112-2322321033103202-3003131311303112): complete subsection reference.

<a id="canonical-1001102333002031-3301133330302320-2020203102203001-0002301122213013-1201213202131333-0211333012322233-1030222200321030-1222101203121303"></a>

<a id="canonical-3120001022320312-2320211333120001-1223333112011311-0333002131233022-0132011221223013-3213112211232310-3120231021310311-1022303033230120"></a>

## type property — nexthop / 102230201213 / 4

Type: `"string"`. Optional.

\[Enum: NEXT\_HOP\_DEFAULT\_GATEWAY|NEXT\_HOP\_USE\_CONFIGURED|NEXT\_HOP\_NETWORK\_INTERFACE\]
Defines types of next-hop Use default gateway on the local interface as gateway for route. Assumes
there is only one local interface on the virtual network. Use the specified address as nexthop Use
the network interface as nexthop Discard nexthop, used when attr type is Advertise Used in VoltADN..
Possible values are \`NEXT\_HOP\_DEFAULT\_GATEWAY\`, \`NEXT\_HOP\_USE\_CONFIGURED\`,
\`NEXT\_HOP\_NETWORK\_INTERFACE\`. Defaults to \`NEXT\_HOP\_DEFAULT\_GATEWAY\`.

Upstream description:

Defines types of next-hop

Use default gateway on the local interface as gateway for route. Assumes there is only one local
interface on the virtual network. Use the specified address as nexthop Use the network interface as
nexthop Discard nexthop, used when attr type is Advertise Used in VoltADN private virtual network.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("NEXT_HOP_DEFAULT_GATEWAY",
    "NEXT_HOP_USE_CONFIGURED",
    "NEXT_HOP_NETWORK_INTERFACE"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "NEXT_HOP_DEFAULT_GATEWAY",
  "enum": [
    "NEXT_HOP_DEFAULT_GATEWAY",
    "NEXT_HOP_USE_CONFIGURED",
    "NEXT_HOP_NETWORK_INTERFACE"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0011203313213310-1323030113000311-3203012002323223-3330230011002201-2201011133121101-0021033020230310-0201222201102111-1230113103101121"></a>

## Next pages — nexthop / 102230201213 / 5

- [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop.interface](resources--azure_vnet_site--reference--group-006.md#canonical-1101033021223122-2120100010030203-2022230113103332-0322110301022301-0233333031222310-3103302010130020-1202133300100212-3120202220122221)
- [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](resources--azure_vnet_site--reference--group-006.md#canonical-0203210121300221-3132113002113110-1022021220213313-3302221233133232-3320322013233011-0112300223011112-2322321033103202-3003131311303112)
- [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route](resources--azure_vnet_site--reference--group-006.md#canonical-2222001221311112-1030220102011022-1210123221302232-2311022330031311-2102131031323330-0030302301313301-2121330312122122-2133123231233022)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)

<a id="canonical-1101033021223122-2120100010030203-2022230113103332-0322110301022301-0233333031222310-3103302010130020-1202133300100212-3120202220122221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3310030323202020-0112221313203011-1311212110203003-0000220021221111-2232330321321123-3023133233120301-1003211300322033-0332031112310032"></a>

## ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop.interface — interface / 100300130203 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-3012120331203201-1112000130332131-3202332120003001-0123013103333002-1112013103100001-2021323312110130-1112011111221321-2323233113231232)
- [ingress_egress_gw_ar](resources--azure_vnet_site--reference--group-005.md#canonical-0312012003033102-3210313212120203-0230331020120000-0030121220200120-2223003302123103-2232111103103011-2023112221000212-0212333313012311)
- [ingress_egress_gw_ar.inside_static_routes](resources--azure_vnet_site--reference--group-006.md#canonical-2003312233332113-0111122021131232-0103020222220000-1131133231223012-3330001301220131-1222310013113030-1013300112132003-0330213300002202)
- [ingress_egress_gw_ar.inside_static_routes.static_route_list](resources--azure_vnet_site--reference--group-006.md#canonical-2323231232330320-3210301011110220-1130100131330230-2113013022002303-1123030101333321-1331230102100131-3002210131120221-3010111311031010)
- [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route](resources--azure_vnet_site--reference--group-006.md#canonical-2222001221311112-1030220102011022-1210123221302232-2311022330031311-2102131031323330-0030302301313301-2121330312122122-2133123231233022)
- [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop](resources--azure_vnet_site--reference--group-006.md#canonical-0302121313113312-3323313023020010-3032301101221202-3210120223011032-2311022102303022-3322330023021203-1013222221302132-2110003320121013)
- ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop.interface

<a id="canonical-3113101033103313-3000233112002133-0211112022231103-2332312222022111-2303021301312332-1002232303221012-2233232221223211-1311210122003321"></a>

Type: `"object"`. list nested block, Optional.

Nexthop is network interface when type is 'Network-Interface'.

Upstream description:

Nexthop is network interface when type is "Network-Interface"

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 1,
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
    "ves.io.schema.rules.repeated.max_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1"
  }
}
```

Terraform syntax:

```terraform
interface {
  # Configure direct properties listed below.
}
```

<a id="canonical-0310113301210313-0231100332330300-0000122232101332-0133020032332333-0010000103033010-0122033312300100-2001003312032312-1301013002201322"></a>

## Direct properties — interface / 100300130203 / 3

<a id="canonical-3223213333322231-1133033302020001-0232102133201311-2122001120232330-0113122313121101-3301010110303331-3110013303302011-0012213332033101"></a>

<a id="canonical-2110320000101300-3313121031203323-1022211013112311-1321020123100032-3223030332021311-2330033303113102-0011231231032212-2322032232302033"></a>

## kind property — interface / 100300130203 / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. 'route').

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. "route")

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

<a id="canonical-1301333003322310-1231000223303312-0301131031310210-2130121232222003-3133131200111110-0021112032232133-0221221110303310-2112123323030233"></a>

<a id="canonical-2323332230021033-2231220111221023-1120310013231221-1323222303112012-2313320330212230-0212132002113033-3021003203102220-1332013210333301"></a>

## name property — interface / 100300130203 / 5

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

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

<a id="canonical-3112202211301111-3302003233231021-1213302232211122-3103031222031010-2312322021200313-0300333203122321-3003013211232232-2311203111333121"></a>

<a id="canonical-0102101301120332-3013201030330331-3231112023132212-0321111001121320-3032232032112012-3321300100202232-0021133222133231-3233202130120113"></a>

## namespace property — interface / 100300130203 / 6

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
    "formatDescription": "DNS-1035 label: must start with a lowercase letter",
    "maxLength": 63,
    "metadata": {
      "confidence": 0.99,
      "source": "inferred",
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
  }
}
```

<a id="canonical-3003030103310120-1203113011310301-1320031102113231-1013000322232102-2200322332001300-3331313232233322-2123232210220302-2003223121003021"></a>

<a id="canonical-3110323322103220-3313132123323301-2033103202212310-0322223322131222-0013122210113010-3021033333202301-0120113103131120-2113330220013131"></a>

## tenant property — interface / 100300130203 / 7

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

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

<a id="canonical-2021321101132222-3320123113210123-1213223301032031-2123313100111300-2112320112202131-3333331210331300-1001220102000333-1023231003211232"></a>

<a id="canonical-1202010232031003-2202222201110021-0102320201320000-0211313032303122-2022233232222300-1331113201303320-2012231300322230-2012232102032122"></a>

## uid property — interface / 100300130203 / 8

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then uid will hold the
referred object's(e.g. Route's) uid.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then uid will hold the
referred object's(e.g. Route's) uid.

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

<a id="canonical-3021000230001131-1213103120221111-3322103032133203-0022213100221010-1002011012211220-2311330302102100-1031203200022023-0002311020110232"></a>

## Next pages — interface / 100300130203 / 9

- [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop](resources--azure_vnet_site--reference--group-006.md#canonical-0302121313113312-3323313023020010-3032301101221202-3210120223011032-2311022102303022-3322330023021203-1013222221302132-2110003320121013)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)

<a id="canonical-0203210121300221-3132113002113110-1022021220213313-3302221233133232-3320322013233011-0112300223011112-2322321033103202-3003131311303112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3211003233000032-1021110100121310-0301203010120320-2020332233032320-1231002302203233-0321123300132210-2121110303332330-0232312122032303"></a>

## ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address — nexthop_address / 000103333101 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-3012120331203201-1112000130332131-3202332120003001-0123013103333002-1112013103100001-2021323312110130-1112011111221321-2323233113231232)
- [ingress_egress_gw_ar](resources--azure_vnet_site--reference--group-005.md#canonical-0312012003033102-3210313212120203-0230331020120000-0030121220200120-2223003302123103-2232111103103011-2023112221000212-0212333313012311)
- [ingress_egress_gw_ar.inside_static_routes](resources--azure_vnet_site--reference--group-006.md#canonical-2003312233332113-0111122021131232-0103020222220000-1131133231223012-3330001301220131-1222310013113030-1013300112132003-0330213300002202)
- [ingress_egress_gw_ar.inside_static_routes.static_route_list](resources--azure_vnet_site--reference--group-006.md#canonical-2323231232330320-3210301011110220-1130100131330230-2113013022002303-1123030101333321-1331230102100131-3002210131120221-3010111311031010)
- [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route](resources--azure_vnet_site--reference--group-006.md#canonical-2222001221311112-1030220102011022-1210123221302232-2311022330031311-2102131031323330-0030302301313301-2121330312122122-2133123231233022)
- [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop](resources--azure_vnet_site--reference--group-006.md#canonical-0302121313113312-3323313023020010-3032301101221202-3210120223011032-2311022102303022-3322330023021203-1013222221302132-2110003320121013)
- ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address

<a id="canonical-2023303000210112-0111011312012313-0030202203120110-2311113200300300-1101303131003200-2001103021120311-2012332332021020-0010112323012032"></a>

Type: `"object"`. single nested block, Optional.

IP Address used to specify an IPv4 or IPv6 address.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("dual_stack",
    "ipv4"),
  validators.ConflictingObjectAttributes("dual_stack",
    "ipv6"),
  validators.ConflictingObjectAttributes("ipv4",
    "ipv6")}
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
  "x-ves-oneof-field-ver": "[\"dual_stack\",\"ipv4\",\"ipv6\"]"
}
```

Terraform syntax:

```terraform
nexthop_address {
  # Configure direct properties listed below.
}
```

<a id="canonical-1121132032313233-3220330112213300-2322111332002012-2023010001003131-3220200301202013-1010332103001331-2113303103201121-3223230221322210"></a>

## Direct properties — nexthop_address / 000103333101 / 3

- [dual_stack](resources--azure_vnet_site--reference--group-006.md#canonical-3230032332010322-0301112330200010-0113101221002013-2322300022011303-2311031013310230-3321331330013123-0311333012101322-2000133320213102): complete subsection reference.

- [IPv4](resources--azure_vnet_site--reference--group-006.md#canonical-2102121133012312-2323032101010113-1013321122322122-0333121222231130-3233112012320033-2222322231322232-3113302210203203-0202031333030113): complete subsection reference.

- [IPv6](resources--azure_vnet_site--reference--group-006.md#canonical-2011200011020121-2203301302100000-2032321313303212-3313012331202113-3013332323130303-0113301202311031-2101002201213032-3122213310120121): complete subsection reference.

<a id="canonical-2313130011311033-2010000212312103-2101113302003033-1022321031100112-2013221320321322-0201331301112220-2123013203210030-3020201333223100"></a>

## Next pages — nexthop_address / 000103333101 / 4

- [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](resources--azure_vnet_site--reference--group-006.md#canonical-3230032332010322-0301112330200010-0113101221002013-2322300022011303-2311031013310230-3321331330013123-0311333012101322-2000133320213102)
- [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4](resources--azure_vnet_site--reference--group-006.md#canonical-2102121133012312-2323032101010113-1013321122322122-0333121222231130-3233112012320033-2222322231322232-3113302210203203-0202031333030113)
- [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6](resources--azure_vnet_site--reference--group-006.md#canonical-2011200011020121-2203301302100000-2032321313303212-3313012331202113-3013332323130303-0113301202311031-2101002201213032-3122213310120121)
- [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop](resources--azure_vnet_site--reference--group-006.md#canonical-0302121313113312-3323313023020010-3032301101221202-3210120223011032-2311022102303022-3322330023021203-1013222221302132-2110003320121013)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)

<a id="canonical-3230032332010322-0301112330200010-0113101221002013-2322300022011303-2311031013310230-3321331330013123-0311333012101322-2000133320213102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1032323110101122-0322121202133023-3220110013023003-3102032032323322-2212103321222230-2010332222132133-2120013031230110-1023200211101000"></a>

## ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack — dual_stack / 103120003231 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-3012120331203201-1112000130332131-3202332120003001-0123013103333002-1112013103100001-2021323312110130-1112011111221321-2323233113231232)
- [ingress_egress_gw_ar](resources--azure_vnet_site--reference--group-005.md#canonical-0312012003033102-3210313212120203-0230331020120000-0030121220200120-2223003302123103-2232111103103011-2023112221000212-0212333313012311)
- [ingress_egress_gw_ar.inside_static_routes](resources--azure_vnet_site--reference--group-006.md#canonical-2003312233332113-0111122021131232-0103020222220000-1131133231223012-3330001301220131-1222310013113030-1013300112132003-0330213300002202)
- [ingress_egress_gw_ar.inside_static_routes.static_route_list](resources--azure_vnet_site--reference--group-006.md#canonical-2323231232330320-3210301011110220-1130100131330230-2113013022002303-1123030101333321-1331230102100131-3002210131120221-3010111311031010)
- [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route](resources--azure_vnet_site--reference--group-006.md#canonical-2222001221311112-1030220102011022-1210123221302232-2311022330031311-2102131031323330-0030302301313301-2121330312122122-2133123231233022)
- [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop](resources--azure_vnet_site--reference--group-006.md#canonical-0302121313113312-3323313023020010-3032301101221202-3210120223011032-2311022102303022-3322330023021203-1013222221302132-2110003320121013)
- [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](resources--azure_vnet_site--reference--group-006.md#canonical-0203210121300221-3132113002113110-1022021220213313-3302221233133232-3320322013233011-0112300223011112-2322321033103202-3003131311303112)
- ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack

<a id="canonical-1113032232032022-2012103011100021-3010301001232112-1301321102230230-3230013302123303-0133202121302000-2033301320211132-1311301220313120"></a>

Type: `"object"`. single nested block, Optional.

DualStackAddressType represents both IPv4 and IPv6 together.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
dual_stack {
  # Configure direct properties listed below.
}
```

<a id="canonical-1111031031013230-2210132323002323-0010231313323331-0101003001103213-3211110001011021-2022221123310302-1011302022202000-3121232031213110"></a>

## Direct properties — dual_stack / 103120003231 / 3

- [IPv4](resources--azure_vnet_site--reference--group-006.md#canonical-3222302201012022-3201033131001122-3002033130223303-0101330210301230-1322023111300202-3303323120210210-0102121133203010-2210301330220311): complete subsection reference.

- [IPv6](resources--azure_vnet_site--reference--group-006.md#canonical-3032202000210303-0020233320303132-1323103311010212-1222012112223220-1131113303230132-3013311321223300-1301220010212033-0203002101320211): complete subsection reference.

<a id="canonical-3213221022102030-1222033331310310-0101211013031012-1332021311003030-0122102301330013-0323000100112132-1230000303112000-3133333323233000"></a>

## Next pages — dual_stack / 103120003231 / 4

- [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4](resources--azure_vnet_site--reference--group-006.md#canonical-3222302201012022-3201033131001122-3002033130223303-0101330210301230-1322023111300202-3303323120210210-0102121133203010-2210301330220311)
- [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6](resources--azure_vnet_site--reference--group-006.md#canonical-3032202000210303-0020233320303132-1323103311010212-1222012112223220-1131113303230132-3013311321223300-1301220010212033-0203002101320211)
- [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](resources--azure_vnet_site--reference--group-006.md#canonical-0203210121300221-3132113002113110-1022021220213313-3302221233133232-3320322013233011-0112300223011112-2322321033103202-3003131311303112)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)

<a id="canonical-3222302201012022-3201033131001122-3002033130223303-0101330210301230-1322023111300202-3303323120210210-0102121133203010-2210301330220311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2122312013111113-3322120112121310-0333122313300130-2211232122203232-1323312100120031-3121032222230000-1032102030212313-0013202133330022"></a>

## ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.IPv4 — IPv4 / 113313210102 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-3012120331203201-1112000130332131-3202332120003001-0123013103333002-1112013103100001-2021323312110130-1112011111221321-2323233113231232)
- [ingress_egress_gw_ar](resources--azure_vnet_site--reference--group-005.md#canonical-0312012003033102-3210313212120203-0230331020120000-0030121220200120-2223003302123103-2232111103103011-2023112221000212-0212333313012311)
- [ingress_egress_gw_ar.inside_static_routes](resources--azure_vnet_site--reference--group-006.md#canonical-2003312233332113-0111122021131232-0103020222220000-1131133231223012-3330001301220131-1222310013113030-1013300112132003-0330213300002202)
- [ingress_egress_gw_ar.inside_static_routes.static_route_list](resources--azure_vnet_site--reference--group-006.md#canonical-2323231232330320-3210301011110220-1130100131330230-2113013022002303-1123030101333321-1331230102100131-3002210131120221-3010111311031010)
- [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route](resources--azure_vnet_site--reference--group-006.md#canonical-2222001221311112-1030220102011022-1210123221302232-2311022330031311-2102131031323330-0030302301313301-2121330312122122-2133123231233022)
- [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop](resources--azure_vnet_site--reference--group-006.md#canonical-0302121313113312-3323313023020010-3032301101221202-3210120223011032-2311022102303022-3322330023021203-1013222221302132-2110003320121013)
- [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](resources--azure_vnet_site--reference--group-006.md#canonical-0203210121300221-3132113002113110-1022021220213313-3302221233133232-3320322013233011-0112300223011112-2322321033103202-3003131311303112)
- [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](resources--azure_vnet_site--reference--group-006.md#canonical-3230032332010322-0301112330200010-0113101221002013-2322300022011303-2311031013310230-3321331330013123-0311333012101322-2000133320213102)
- ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.IPv4

<a id="canonical-0203230023032110-2122132300312231-3222200221121211-2230300313220020-0033023121121223-2023112323012003-1131031332112310-3210220333310030"></a>

Type: `"object"`. single nested block, Optional.

IPv4 address in dotted decimal notation (e.g., 192.0.2.1).

Upstream description:

IPv4 Address in dot-decimal notation.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
ipv4 {
  # Configure direct properties listed below.
}
```

<a id="canonical-3321010303222233-1311323112033002-0222313322201223-1011022203101100-0020230100301211-1203202112001222-1122302223211100-0300010300210100"></a>

## Direct properties — IPv4 / 113313210102 / 3

<a id="canonical-1211210211101122-0012012133303200-0332322121301220-3233022300113023-2320120133233231-1313022120131212-3320323003122132-0121122011202321"></a>

<a id="canonical-0131202201011233-3031013020022000-3002330303331031-2130203103030033-0200331201030321-3202310110331233-1003111322013313-3222302101300013"></a>

## addr property — IPv4 / 113313210102 / 4

Type: `"string"`. Optional.

IPv4 Address in string form with dot-decimal notation.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
  validators.IPv4Validator(),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ipv4",
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
    "ves.io.schema.rules.string.ipv4": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv4": "true"
  }
}
```

<a id="canonical-3112001103030020-3333203020022033-3201201002322100-1021001232231130-2310123333010331-3330300312003222-2131033030122131-3212320320111320"></a>

## Next pages — IPv4 / 113313210102 / 5

- [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](resources--azure_vnet_site--reference--group-006.md#canonical-3230032332010322-0301112330200010-0113101221002013-2322300022011303-2311031013310230-3321331330013123-0311333012101322-2000133320213102)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)

<a id="canonical-3032202000210303-0020233320303132-1323103311010212-1222012112223220-1131113303230132-3013311321223300-1301220010212033-0203002101320211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3320310110002333-0022100233011100-0333221330120300-2101223303233211-0110112010323312-3113033230122030-2132200203123203-3302333131012033"></a>

## ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.IPv6 — IPv6 / 113302213312 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-3012120331203201-1112000130332131-3202332120003001-0123013103333002-1112013103100001-2021323312110130-1112011111221321-2323233113231232)
- [ingress_egress_gw_ar](resources--azure_vnet_site--reference--group-005.md#canonical-0312012003033102-3210313212120203-0230331020120000-0030121220200120-2223003302123103-2232111103103011-2023112221000212-0212333313012311)
- [ingress_egress_gw_ar.inside_static_routes](resources--azure_vnet_site--reference--group-006.md#canonical-2003312233332113-0111122021131232-0103020222220000-1131133231223012-3330001301220131-1222310013113030-1013300112132003-0330213300002202)
- [ingress_egress_gw_ar.inside_static_routes.static_route_list](resources--azure_vnet_site--reference--group-006.md#canonical-2323231232330320-3210301011110220-1130100131330230-2113013022002303-1123030101333321-1331230102100131-3002210131120221-3010111311031010)
- [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route](resources--azure_vnet_site--reference--group-006.md#canonical-2222001221311112-1030220102011022-1210123221302232-2311022330031311-2102131031323330-0030302301313301-2121330312122122-2133123231233022)
- [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop](resources--azure_vnet_site--reference--group-006.md#canonical-0302121313113312-3323313023020010-3032301101221202-3210120223011032-2311022102303022-3322330023021203-1013222221302132-2110003320121013)
- [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](resources--azure_vnet_site--reference--group-006.md#canonical-0203210121300221-3132113002113110-1022021220213313-3302221233133232-3320322013233011-0112300223011112-2322321033103202-3003131311303112)
- [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](resources--azure_vnet_site--reference--group-006.md#canonical-3230032332010322-0301112330200010-0113101221002013-2322300022011303-2311031013310230-3321331330013123-0311333012101322-2000133320213102)
- ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.IPv6

<a id="canonical-1020112201101111-0221111310230102-2121312331002030-1332013233030021-2112033122120030-1130332031123113-0331000121333323-3232331002211011"></a>

Type: `"object"`. single nested block, Optional.

IPv6 Address specified as hexadecimal numbers separated by ':'.

Upstream description:

IPv6 Address specified as hexadecimal numbers separated by ':'

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
ipv6 {
  # Configure direct properties listed below.
}
```

<a id="canonical-3131013003300132-0032011033310213-3203313210203310-0303202222213111-0332010003210212-3301031333230311-0013231312011123-3103012123310033"></a>

## Direct properties — IPv6 / 113302213312 / 3

<a id="canonical-2222321020001113-2300220210121023-2312330312222313-0320312233103211-3033310111020211-1303213210013212-3020213021230313-3122020202212013"></a>

<a id="canonical-2312321322030201-1131122103023103-2132332330232000-0233000223311213-1021323220000321-0210301121110301-0132201312131001-1020230230233231"></a>

## addr property — IPv6 / 113302213312 / 4

Type: `"string"`. Optional.

IPv6 Address in form of string. IPv6 address must be specified as hexadecimal numbers separated by
':' The address can be compacted by suppressing zeros e.g. '2001:db8:0:0:0:0:2:1' becomes
'2001:db8::2:1' or '2001:db8:0:0:0:2:0:0' becomes '2001:db8::2::'.

Upstream description:

IPv6 Address in form of string. IPv6 address must be specified as hexadecimal numbers separated by
':' The address can be compacted by suppressing zeros e.g. '2001:db8:0:0:0:0:2:1' becomes
'2001:db8::2:1' or '2001:db8:0:0:0:2:0:0' becomes '2001:db8::2::'

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
  validators.IPv6Validator(),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ipv6",
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
    "ves.io.schema.rules.string.ipv6": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv6": "true"
  }
}
```

<a id="canonical-2331323303213033-3312231102112311-3310133020231030-1111132222203212-1030100232011030-0301223221322113-2313133012202113-1103122202032211"></a>

## Next pages — IPv6 / 113302213312 / 5

- [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](resources--azure_vnet_site--reference--group-006.md#canonical-3230032332010322-0301112330200010-0113101221002013-2322300022011303-2311031013310230-3321331330013123-0311333012101322-2000133320213102)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)

<a id="canonical-2102121133012312-2323032101010113-1013321122322122-0333121222231130-3233112012320033-2222322231322232-3113302210203203-0202031333030113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0303031110223130-2310033033210300-1201100202122120-0320013101102003-2331131123010030-1131003133310312-2322301323312331-0210000203111321"></a>

## ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.IPv4 — IPv4 / 212032001303 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-3012120331203201-1112000130332131-3202332120003001-0123013103333002-1112013103100001-2021323312110130-1112011111221321-2323233113231232)
- [ingress_egress_gw_ar](resources--azure_vnet_site--reference--group-005.md#canonical-0312012003033102-3210313212120203-0230331020120000-0030121220200120-2223003302123103-2232111103103011-2023112221000212-0212333313012311)
- [ingress_egress_gw_ar.inside_static_routes](resources--azure_vnet_site--reference--group-006.md#canonical-2003312233332113-0111122021131232-0103020222220000-1131133231223012-3330001301220131-1222310013113030-1013300112132003-0330213300002202)
- [ingress_egress_gw_ar.inside_static_routes.static_route_list](resources--azure_vnet_site--reference--group-006.md#canonical-2323231232330320-3210301011110220-1130100131330230-2113013022002303-1123030101333321-1331230102100131-3002210131120221-3010111311031010)
- [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route](resources--azure_vnet_site--reference--group-006.md#canonical-2222001221311112-1030220102011022-1210123221302232-2311022330031311-2102131031323330-0030302301313301-2121330312122122-2133123231233022)
- [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop](resources--azure_vnet_site--reference--group-006.md#canonical-0302121313113312-3323313023020010-3032301101221202-3210120223011032-2311022102303022-3322330023021203-1013222221302132-2110003320121013)
- [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](resources--azure_vnet_site--reference--group-006.md#canonical-0203210121300221-3132113002113110-1022021220213313-3302221233133232-3320322013233011-0112300223011112-2322321033103202-3003131311303112)
- ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.IPv4

<a id="canonical-0232112002123011-2332113303212113-3200302033012323-0012311000211112-2323112230121323-0122012222103123-2200100302103120-3033132323013323"></a>

Type: `"object"`. single nested block, Optional.

IPv4 address in dotted decimal notation (e.g., 192.0.2.1).

Upstream description:

IPv4 Address in dot-decimal notation.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
ipv4 {
  # Configure direct properties listed below.
}
```

<a id="canonical-2131033012201011-2210001112213312-0203112331001232-0333111033111103-0330113302132032-1032101313101210-1232221310300012-3213232003112213"></a>

## Direct properties — IPv4 / 212032001303 / 3

<a id="canonical-0213003302333302-0013233302220332-0330321132113211-2212131331203000-0133213211003202-1000122321113003-2133210002213302-2011211003033331"></a>

<a id="canonical-1130320012233300-1133223223302033-0203100031033302-1321233212122220-3132311130201033-0211000333132211-1332121332233230-0000302000011002"></a>

## addr property — IPv4 / 212032001303 / 4

Type: `"string"`. Optional.

IPv4 Address in string form with dot-decimal notation.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
  validators.IPv4Validator(),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ipv4",
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
    "ves.io.schema.rules.string.ipv4": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv4": "true"
  }
}
```

<a id="canonical-1132210013003221-0232130313230021-0033232303230313-0110300101100312-2312021033121333-0011222003213221-2032212010001301-2321111132201111"></a>

## Next pages — IPv4 / 212032001303 / 5

- [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](resources--azure_vnet_site--reference--group-006.md#canonical-0203210121300221-3132113002113110-1022021220213313-3302221233133232-3320322013233011-0112300223011112-2322321033103202-3003131311303112)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)

<a id="canonical-2011200011020121-2203301302100000-2032321313303212-3313012331202113-3013332323130303-0113301202311031-2101002201213032-3122213310120121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2101101211013033-2211113112231223-1120130313123112-2020210210201113-1023223233012200-1332312013112331-2131310102223110-1101103100112133"></a>

## ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.IPv6 — IPv6 / 001310331103 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-3012120331203201-1112000130332131-3202332120003001-0123013103333002-1112013103100001-2021323312110130-1112011111221321-2323233113231232)
- [ingress_egress_gw_ar](resources--azure_vnet_site--reference--group-005.md#canonical-0312012003033102-3210313212120203-0230331020120000-0030121220200120-2223003302123103-2232111103103011-2023112221000212-0212333313012311)
- [ingress_egress_gw_ar.inside_static_routes](resources--azure_vnet_site--reference--group-006.md#canonical-2003312233332113-0111122021131232-0103020222220000-1131133231223012-3330001301220131-1222310013113030-1013300112132003-0330213300002202)
- [ingress_egress_gw_ar.inside_static_routes.static_route_list](resources--azure_vnet_site--reference--group-006.md#canonical-2323231232330320-3210301011110220-1130100131330230-2113013022002303-1123030101333321-1331230102100131-3002210131120221-3010111311031010)
- [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route](resources--azure_vnet_site--reference--group-006.md#canonical-2222001221311112-1030220102011022-1210123221302232-2311022330031311-2102131031323330-0030302301313301-2121330312122122-2133123231233022)
- [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop](resources--azure_vnet_site--reference--group-006.md#canonical-0302121313113312-3323313023020010-3032301101221202-3210120223011032-2311022102303022-3322330023021203-1013222221302132-2110003320121013)
- [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](resources--azure_vnet_site--reference--group-006.md#canonical-0203210121300221-3132113002113110-1022021220213313-3302221233133232-3320322013233011-0112300223011112-2322321033103202-3003131311303112)
- ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.IPv6

<a id="canonical-0002321322100100-0030000201113222-2030003122313321-0320212233112120-0302102121132023-3102131033013010-1211231102230111-2322020233202103"></a>

Type: `"object"`. single nested block, Optional.

IPv6 Address specified as hexadecimal numbers separated by ':'.

Upstream description:

IPv6 Address specified as hexadecimal numbers separated by ':'

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
ipv6 {
  # Configure direct properties listed below.
}
```

<a id="canonical-1323031203320000-1333332311312230-2331300230201001-2313130033021003-1221101012130131-2001331303213023-1313230221323132-0321001232223202"></a>

## Direct properties — IPv6 / 001310331103 / 3

<a id="canonical-3303210202213301-1212031233012210-0330122230110110-0332011303230031-0210230302301122-0121212021332101-1213113323002103-2311221231120020"></a>

<a id="canonical-3020222132121201-2330311130300332-1300332110300210-3223313333211021-3311101000302222-3311311322131130-1132022023030113-2331121332020220"></a>

## addr property — IPv6 / 001310331103 / 4

Type: `"string"`. Optional.

IPv6 Address in form of string. IPv6 address must be specified as hexadecimal numbers separated by
':' The address can be compacted by suppressing zeros e.g. '2001:db8:0:0:0:0:2:1' becomes
'2001:db8::2:1' or '2001:db8:0:0:0:2:0:0' becomes '2001:db8::2::'.

Upstream description:

IPv6 Address in form of string. IPv6 address must be specified as hexadecimal numbers separated by
':' The address can be compacted by suppressing zeros e.g. '2001:db8:0:0:0:0:2:1' becomes
'2001:db8::2:1' or '2001:db8:0:0:0:2:0:0' becomes '2001:db8::2::'

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
  validators.IPv6Validator(),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ipv6",
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
    "ves.io.schema.rules.string.ipv6": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv6": "true"
  }
}
```

<a id="canonical-2223221100102103-0333311300012003-0300200012132231-1021232012231111-2211130211211021-2322333301220321-3030113331301022-2101313310220312"></a>

## Next pages — IPv6 / 001310331103 / 5

- [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](resources--azure_vnet_site--reference--group-006.md#canonical-0203210121300221-3132113002113110-1022021220213313-3302221233133232-3320322013233011-0112300223011112-2322321033103202-3003131311303112)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)

<a id="canonical-1331313323323133-2232110222031313-3303122302120020-2121113000333222-1203300200213111-2012310222120112-0021020033323032-0223130310030220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3100221012203302-0113200300210200-2203000233203200-1323322332013331-2110232101031213-1133300332120011-3122231122003200-2220103303313322"></a>

## ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.subnets — subnets / 233200210010 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-3012120331203201-1112000130332131-3202332120003001-0123013103333002-1112013103100001-2021323312110130-1112011111221321-2323233113231232)
- [ingress_egress_gw_ar](resources--azure_vnet_site--reference--group-005.md#canonical-0312012003033102-3210313212120203-0230331020120000-0030121220200120-2223003302123103-2232111103103011-2023112221000212-0212333313012311)
- [ingress_egress_gw_ar.inside_static_routes](resources--azure_vnet_site--reference--group-006.md#canonical-2003312233332113-0111122021131232-0103020222220000-1131133231223012-3330001301220131-1222310013113030-1013300112132003-0330213300002202)
- [ingress_egress_gw_ar.inside_static_routes.static_route_list](resources--azure_vnet_site--reference--group-006.md#canonical-2323231232330320-3210301011110220-1130100131330230-2113013022002303-1123030101333321-1331230102100131-3002210131120221-3010111311031010)
- [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route](resources--azure_vnet_site--reference--group-006.md#canonical-2222001221311112-1030220102011022-1210123221302232-2311022330031311-2102131031323330-0030302301313301-2121330312122122-2133123231233022)
- ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.subnets

<a id="canonical-3101003130312131-1020130133033233-1300331212000203-0020203231122030-1303001310323101-3120010302312203-0113232022210232-3110131322330121"></a>

Type: `"object"`. list nested block, Optional.

Subnets. List of route prefixes.

Upstream description:

List of route prefixes.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("ipv4",
    "ipv6")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "256"
  }
}
```

Terraform syntax:

```terraform
subnets {
  # Configure direct properties listed below.
}
```

<a id="canonical-0111033113321001-0013010021003320-3030202030233102-1032022312021220-1113221130333120-0030301210333233-2233331210032033-1030210211023232"></a>

## Direct properties — subnets / 233200210010 / 3

- [IPv4](resources--azure_vnet_site--reference--group-006.md#canonical-2213010313120321-0001220212030220-1020113112033110-1203003311010322-3333110333132122-3120010203112122-0132020130231300-3031213120111302): complete subsection reference.

- [IPv6](resources--azure_vnet_site--reference--group-007.md#canonical-2021131022311213-1123101110132000-1122210101013221-1313333311323200-1001322003031000-2000000333030131-2023321003113112-1220030211210010): complete subsection reference.

<a id="canonical-1303113213222001-3221321020102110-3333121222100130-1332110133303330-1321031213332123-1003023201020030-3212133003231003-2330022122012102"></a>

## Next pages — subnets / 233200210010 / 4

- [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.subnets.ipv4](resources--azure_vnet_site--reference--group-006.md#canonical-2213010313120321-0001220212030220-1020113112033110-1203003311010322-3333110333132122-3120010203112122-0132020130231300-3031213120111302)
- [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.subnets.ipv6](resources--azure_vnet_site--reference--group-007.md#canonical-2021131022311213-1123101110132000-1122210101013221-1313333311323200-1001322003031000-2000000333030131-2023321003113112-1220030211210010)
- [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route](resources--azure_vnet_site--reference--group-006.md#canonical-2222001221311112-1030220102011022-1210123221302232-2311022330031311-2102131031323330-0030302301313301-2121330312122122-2133123231233022)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)

<a id="canonical-2213010313120321-0001220212030220-1020113112033110-1203003311010322-3333110333132122-3120010203112122-0132020130231300-3031213120111302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->
