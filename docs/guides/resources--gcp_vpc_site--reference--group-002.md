---
page_title: "xcsh_gcp_vpc_site reference"
subcategory: "Infrastructure"
description: "Complete grouped canonical reference for xcsh_gcp_vpc_site reference."
---

# xcsh_gcp_vpc_site reference

<a id="canonical-2022310312302302-0330302223312311-3233202030200001-0232032033322121-1002023012333323-0310113333132200-1202220123113332-2133233211200220"></a>

## Direct properties — ingress_egress_gw / 331213233023 / 3

- [active_enhanced_firewall_policies](resources--gcp_vpc_site--reference--group-002.md#canonical-1301001123030321-3330020323110300-0010201202332110-0232111331033322-2223300110000312-0302201210302013-3313220231131310-2300313233100213): complete subsection reference.

- [active_forward_proxy_policies](resources--gcp_vpc_site--reference--group-002.md#canonical-2311231030030323-0303023002330311-2113312033103203-2330303331111123-2231102303100300-2002222211231230-3122033210231030-1103130311303323): complete subsection reference.

- [active_network_policies](resources--gcp_vpc_site--reference--group-002.md#canonical-0212013333022211-1001302012023020-3122102202102122-3103121111332212-0021111032333010-0002012012010100-1220132220131102-2001123001010132): complete subsection reference.

- [dc_cluster_group_inside_vn](resources--gcp_vpc_site--reference--group-002.md#canonical-0211003021310313-1300012031202133-2303213320131031-0133313313000021-2130221202103102-3201121233220023-1101033302303113-3120013303232300): complete subsection reference.

- [dc_cluster_group_outside_vn](resources--gcp_vpc_site--reference--group-002.md#canonical-2103132110012023-3322010132201321-3211232011002221-1330321223032322-2022233110013331-3121223121223102-1103013001232203-0303330031101022): complete subsection reference.

- [forward_proxy_allow_all](resources--gcp_vpc_site--reference--group-002.md#canonical-2011302000102201-3232023320300020-0303212120230300-3013012121333131-0322032223310021-3021011110023032-1022201123121123-1331123020220333): complete subsection reference.

<a id="canonical-3132033023002102-3210110010221302-1333203113210213-0220330112001232-0211312221003312-0210312312300022-1332313031203300-2131312302212000"></a>

<a id="canonical-2220111123131302-2201332111133322-1312020022322131-0302030231130200-0033212130101301-1023110311102023-3033011210033320-2222332032022230"></a>

## gcp_certified_hw property — ingress_egress_gw / 331213233023 / 4

Type: `"string"`. Optional.

\[Enum: gcp-byol-multi-nic-voltmesh\] GCP Certified Hardware. Name for GCP certified hardware. The
only possible value is \`gcp-byol-multi-nic-voltmesh\`.

Upstream description:

Name for GCP certified hardware.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(64),
  stringvalidator.OneOf("gcp-byol-multi-nic-voltmesh"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "enum": [
    "gcp-byol-multi-nic-voltmesh"
  ],
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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.in": "[\\\"gcp-byol-multi-nic-voltmesh\\\"]",
    "ves.io.schema.rules.string.max_len": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.in": "[\\\"gcp-byol-multi-nic-voltmesh\\\"]",
    "ves.io.schema.rules.string.max_len": "64"
  }
}
```

<a id="canonical-0133302222332222-2202131313133232-0021101120202220-2212120212111202-0002230322310133-0221122010301101-1320031321130032-1012101110332301"></a>

<a id="canonical-0012013222220231-2023011320220230-0030113031010310-0200301201011000-3120203101120323-2322100112213213-2112202023320220-1130031321202331"></a>

## gcp_zone_names property — ingress_egress_gw / 331213233023 / 5

Type: `["list", "string"]`. Optional.

X-required List of zones when instances will be created, needs to match with region selected.

Upstream description:

X-required List of zones when instances will be created, needs to match with region selected.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(3),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 3,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 3,
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
    "ves.io.schema.rules.repeated.items.string.max_len": "64",
    "ves.io.schema.rules.repeated.max_items": "3",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.max_len": "64",
    "ves.io.schema.rules.repeated.max_items": "3",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [global_network_list](resources--gcp_vpc_site--reference--group-002.md#canonical-1131133100211321-0133332003213130-3001023120110100-2112200313000011-1100333310030310-0010101033020232-0203330321121321-3012232113212212): complete subsection reference.

- [inside_network](resources--gcp_vpc_site--reference--group-002.md#canonical-2123301101220113-0123332320222111-0202203303002102-2012313302013323-3011301022113123-0301003102033320-2332330212323112-0113111321202221): complete subsection reference.

- [inside_static_routes](resources--gcp_vpc_site--reference--group-002.md#canonical-0220321000013112-2030300310000323-2202131321000302-0112111112120211-0223122220232223-0023132200212001-3230213111203011-1330123303013023): complete subsection reference.

- [inside_subnet](resources--gcp_vpc_site--reference--group-002.md#canonical-3222001132100012-1200312003312120-3302312333232311-2232110003311200-1302022010113013-1222012220132232-2322210000001333-2032110122033120): complete subsection reference.

- [no_dc_cluster_group](resources--gcp_vpc_site--reference--group-002.md#canonical-0223121212322031-2211320312222312-2300130323120123-2121231012031031-3333321103332333-2002332220233102-1122031312200130-2200321033313231): complete subsection reference.

- [no_forward_proxy](resources--gcp_vpc_site--reference--group-002.md#canonical-1121332000003013-1222000333333223-2121210310100010-2312003033100003-2111220103101112-0133232123131210-0311203213021300-2133223203320312): complete subsection reference.

- [no_global_network](resources--gcp_vpc_site--reference--group-002.md#canonical-3102021312020122-3021300021202033-0210332020111302-3000312313002100-3022110003022221-1123023001332030-0310210102222211-3103101010200133): complete subsection reference.

- [no_inside_static_routes](resources--gcp_vpc_site--reference--group-002.md#canonical-0201103101130210-0133000033010210-1231030122220020-3320020013012113-1203001131113000-1103210221011202-0331103332210322-3130012302231213): complete subsection reference.

- [no_network_policy](resources--gcp_vpc_site--reference--group-002.md#canonical-1302131011321212-1300221023200122-2300210033331311-0113012022102331-3012030100001333-1111010120111102-1013112203202320-0213131013301310): complete subsection reference.

- [no_outside_static_routes](resources--gcp_vpc_site--reference--group-002.md#canonical-1122203313010332-0123231301322313-2023101323010232-0202300003230303-1123123013113301-2230221312331320-2113021110010331-3223302020032210): complete subsection reference.

<a id="canonical-2122103103322113-1112020030111233-2333313300201131-0101133232122230-0322300033111101-2221322232023201-3021231232013113-3223222031212212"></a>

<a id="canonical-3300300102122210-1022011013031001-2020011101301033-1010113212103002-1022103110303222-3013011330011013-1110323110200310-2231212212313112"></a>

## node_number property — ingress_egress_gw / 331213233023 / 6

Type: `"number"`. Optional.

Number of main nodes to create, either 1 or 3.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.in": "[1,3]"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.in": "[1,3]"
  }
}
```

- [outside_network](resources--gcp_vpc_site--reference--group-002.md#canonical-1211123322330233-3312333012002031-1012211321132000-0012003132301223-2001313132331302-3033013021331012-0200130302231001-1003322233332332): complete subsection reference.

- [outside_static_routes](resources--gcp_vpc_site--reference--group-003.md#canonical-3203322020133213-3323210012231001-3313321013202200-3121201110320133-2111030212103130-3103303221100321-3331122303023000-0100331112221320): complete subsection reference.

- [outside_subnet](resources--gcp_vpc_site--reference--group-003.md#canonical-3331120001113100-3203200201122222-1030313310121000-2333133202033022-0330313230303203-2301303322122303-1233232211210020-1221010302223301): complete subsection reference.

- [performance_enhancement_mode](resources--gcp_vpc_site--reference--group-003.md#canonical-1323331230000321-0231112112000222-1211032031012330-0003021023030323-2032321210111021-2333011123121312-2110101123031310-1213232033121331): complete subsection reference.

- [sm_connection_public_ip](resources--gcp_vpc_site--reference--group-003.md#canonical-3131111033121221-1103120232232131-3021320321321320-2223311122003033-1101002000200113-0001303232031221-1022201020303213-1113232310103233): complete subsection reference.

- [sm_connection_pvt_ip](resources--gcp_vpc_site--reference--group-003.md#canonical-2202031012123202-3022021103302033-0232201010232321-0113232031103202-2011000013133001-0130113202011203-1013001020001232-3231022001303311): complete subsection reference.

<a id="canonical-0002331122132032-2032022013222312-1221000133123220-2132033001313023-1232033230123321-3110013102133321-1121333303110102-1100321222102233"></a>

## Next pages — ingress_egress_gw / 331213233023 / 7

- [ingress_egress_gw.active_enhanced_firewall_policies](resources--gcp_vpc_site--reference--group-002.md#canonical-1301001123030321-3330020323110300-0010201202332110-0232111331033322-2223300110000312-0302201210302013-3313220231131310-2300313233100213)
- [ingress_egress_gw.active_forward_proxy_policies](resources--gcp_vpc_site--reference--group-002.md#canonical-2311231030030323-0303023002330311-2113312033103203-2330303331111123-2231102303100300-2002222211231230-3122033210231030-1103130311303323)
- [ingress_egress_gw.active_network_policies](resources--gcp_vpc_site--reference--group-002.md#canonical-0212013333022211-1001302012023020-3122102202102122-3103121111332212-0021111032333010-0002012012010100-1220132220131102-2001123001010132)
- [ingress_egress_gw.dc_cluster_group_inside_vn](resources--gcp_vpc_site--reference--group-002.md#canonical-0211003021310313-1300012031202133-2303213320131031-0133313313000021-2130221202103102-3201121233220023-1101033302303113-3120013303232300)
- [ingress_egress_gw.dc_cluster_group_outside_vn](resources--gcp_vpc_site--reference--group-002.md#canonical-2103132110012023-3322010132201321-3211232011002221-1330321223032322-2022233110013331-3121223121223102-1103013001232203-0303330031101022)
- [ingress_egress_gw.forward_proxy_allow_all](resources--gcp_vpc_site--reference--group-002.md#canonical-2011302000102201-3232023320300020-0303212120230300-3013012121333131-0322032223310021-3021011110023032-1022201123121123-1331123020220333)
- [ingress_egress_gw.global_network_list](resources--gcp_vpc_site--reference--group-002.md#canonical-1131133100211321-0133332003213130-3001023120110100-2112200313000011-1100333310030310-0010101033020232-0203330321121321-3012232113212212)
- [ingress_egress_gw.inside_network](resources--gcp_vpc_site--reference--group-002.md#canonical-2123301101220113-0123332320222111-0202203303002102-2012313302013323-3011301022113123-0301003102033320-2332330212323112-0113111321202221)
- [ingress_egress_gw.inside_static_routes](resources--gcp_vpc_site--reference--group-002.md#canonical-0220321000013112-2030300310000323-2202131321000302-0112111112120211-0223122220232223-0023132200212001-3230213111203011-1330123303013023)
- [ingress_egress_gw.inside_subnet](resources--gcp_vpc_site--reference--group-002.md#canonical-3222001132100012-1200312003312120-3302312333232311-2232110003311200-1302022010113013-1222012220132232-2322210000001333-2032110122033120)
- [ingress_egress_gw.no_dc_cluster_group](resources--gcp_vpc_site--reference--group-002.md#canonical-0223121212322031-2211320312222312-2300130323120123-2121231012031031-3333321103332333-2002332220233102-1122031312200130-2200321033313231)
- [ingress_egress_gw.no_forward_proxy](resources--gcp_vpc_site--reference--group-002.md#canonical-1121332000003013-1222000333333223-2121210310100010-2312003033100003-2111220103101112-0133232123131210-0311203213021300-2133223203320312)
- [ingress_egress_gw.no_global_network](resources--gcp_vpc_site--reference--group-002.md#canonical-3102021312020122-3021300021202033-0210332020111302-3000312313002100-3022110003022221-1123023001332030-0310210102222211-3103101010200133)
- [ingress_egress_gw.no_inside_static_routes](resources--gcp_vpc_site--reference--group-002.md#canonical-0201103101130210-0133000033010210-1231030122220020-3320020013012113-1203001131113000-1103210221011202-0331103332210322-3130012302231213)
- [ingress_egress_gw.no_network_policy](resources--gcp_vpc_site--reference--group-002.md#canonical-1302131011321212-1300221023200122-2300210033331311-0113012022102331-3012030100001333-1111010120111102-1013112203202320-0213131013301310)
- [ingress_egress_gw.no_outside_static_routes](resources--gcp_vpc_site--reference--group-002.md#canonical-1122203313010332-0123231301322313-2023101323010232-0202300003230303-1123123013113301-2230221312331320-2113021110010331-3223302020032210)
- [ingress_egress_gw.outside_network](resources--gcp_vpc_site--reference--group-002.md#canonical-1211123322330233-3312333012002031-1012211321132000-0012003132301223-2001313132331302-3033013021331012-0200130302231001-1003322233332332)
- [ingress_egress_gw.outside_static_routes](resources--gcp_vpc_site--reference--group-003.md#canonical-3203322020133213-3323210012231001-3313321013202200-3121201110320133-2111030212103130-3103303221100321-3331122303023000-0100331112221320)
- [ingress_egress_gw.outside_subnet](resources--gcp_vpc_site--reference--group-003.md#canonical-3331120001113100-3203200201122222-1030313310121000-2333133202033022-0330313230303203-2301303322122303-1233232211210020-1221010302223301)
- [ingress_egress_gw.performance_enhancement_mode](resources--gcp_vpc_site--reference--group-003.md#canonical-1323331230000321-0231112112000222-1211032031012330-0003021023030323-2032321210111021-2333011123121312-2110101123031310-1213232033121331)
- [ingress_egress_gw.sm_connection_public_ip](resources--gcp_vpc_site--reference--group-003.md#canonical-3131111033121221-1103120232232131-3021320321321320-2223311122003033-1101002000200113-0001303232031221-1022201020303213-1113232310103233)
- [ingress_egress_gw.sm_connection_pvt_ip](resources--gcp_vpc_site--reference--group-003.md#canonical-2202031012123202-3022021103302033-0232201010232321-0113232031103202-2011000013133001-0130113202011203-1013001020001232-3231022001303311)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-2101111213103332-1031003102012122-3033333002321300-0212010323021100-0301331300303302-3321221003203222-1230133302203013-3232231332123212)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-0131000100012312-2121310130001102-2202230331103203-2230222003223001-1113101010113111-2131313030021223-2030122333203101-2213001121001122)

<a id="canonical-1301001123030321-3330020323110300-0010201202332110-0232111331033322-2223300110000312-0302201210302013-3313220231131310-2300313233100213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0033332023001131-0022123223103100-2222011303313000-3100212200003101-1102013000201111-2000222021300330-1320222323210033-2011110221013022"></a>

## ingress_egress_gw.active_enhanced_firewall_policies — active_enhanced_firewall_policies / 100223212302 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-0131000100012312-2121310130001102-2202230331103203-2230222003223001-1113101010113111-2131313030021223-2030122333203101-2213001121001122)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-2101111213103332-1031003102012122-3033333002321300-0212010323021100-0301331300303302-3321221003203222-1230133302203013-3232231332123212)
- [ingress_egress_gw](resources--gcp_vpc_site--reference--group-001.md#canonical-3123100130121022-2313112013232222-0213021230302120-3230300230213031-3101322300112331-0011013220032322-1311222312201210-2213020302010100)
- ingress_egress_gw.active_enhanced_firewall_policies

<a id="canonical-2333230122003302-3020023031221130-3122001123331011-2100011230312313-0223332121311211-3122201002010201-3312021123221012-2222313023210112"></a>

Type: `"object"`. single nested block, Optional.

List of Enhanced Firewall Policies These policies use session-based rules and provide all OPTIONS
available under firewall policies with an additional option for service insertion.

Upstream description:

List of Enhanced Firewall Policies These policies use session-based rules and provide all OPTIONS
available under firewall policies with an additional option for service insertion.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("enhanced_firewall_policies")}
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
active_enhanced_firewall_policies {
  # Configure direct properties listed below.
}
```

<a id="canonical-0311332132030002-2020223201021222-1312231000013123-1103222320302221-2123303320213031-1121023230202220-1223110232223010-2332200112212230"></a>

## Direct properties — active_enhanced_firewall_policies / 100223212302 / 3

- [enhanced_firewall_policies](resources--gcp_vpc_site--reference--group-002.md#canonical-1203000032130333-0003111030121022-0103112032121311-0223120301320201-2130200333231133-3201033322011232-3303103101000302-1213333303310132): complete subsection reference.

<a id="canonical-0012211221200200-2020212120203201-1200100001022022-2310121210201002-2212010211212101-2033201223022322-1003011300033300-3003322123321032"></a>

## Next pages — active_enhanced_firewall_policies / 100223212302 / 4

- [ingress_egress_gw.active_enhanced_firewall_policies.enhanced_firewall_policies](resources--gcp_vpc_site--reference--group-002.md#canonical-1203000032130333-0003111030121022-0103112032121311-0223120301320201-2130200333231133-3201033322011232-3303103101000302-1213333303310132)
- [ingress_egress_gw](resources--gcp_vpc_site--reference--group-001.md#canonical-3123100130121022-2313112013232222-0213021230302120-3230300230213031-3101322300112331-0011013220032322-1311222312201210-2213020302010100)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-0131000100012312-2121310130001102-2202230331103203-2230222003223001-1113101010113111-2131313030021223-2030122333203101-2213001121001122)

<a id="canonical-1203000032130333-0003111030121022-0103112032121311-0223120301320201-2130200333231133-3201033322011232-3303103101000302-1213333303310132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3013303230201330-2013333222133211-0310201213303210-0202320010303302-1100222113302011-2303130121023012-2312200323330100-1303202102002021"></a>

## ingress_egress_gw.active_enhanced_firewall_policies.enhanced_firewall_policies — enhanced_firewall_policies / 310303001312 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-0131000100012312-2121310130001102-2202230331103203-2230222003223001-1113101010113111-2131313030021223-2030122333203101-2213001121001122)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-2101111213103332-1031003102012122-3033333002321300-0212010323021100-0301331300303302-3321221003203222-1230133302203013-3232231332123212)
- [ingress_egress_gw](resources--gcp_vpc_site--reference--group-001.md#canonical-3123100130121022-2313112013232222-0213021230302120-3230300230213031-3101322300112331-0011013220032322-1311222312201210-2213020302010100)
- [ingress_egress_gw.active_enhanced_firewall_policies](resources--gcp_vpc_site--reference--group-002.md#canonical-1301001123030321-3330020323110300-0010201202332110-0232111331033322-2223300110000312-0302201210302013-3313220231131310-2300313233100213)
- ingress_egress_gw.active_enhanced_firewall_policies.enhanced_firewall_policies

<a id="canonical-0113111111031113-3011300103132031-0132203020202210-3302311221232030-3033120002030330-1020320223322022-1131311002022121-3310223231101212"></a>

Type: `"object"`. list nested block, Optional.

Ordered List of Enhanced Firewall Policies active.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("name")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 128,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 128,
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
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

Terraform syntax:

```terraform
enhanced_firewall_policies {
  # Configure direct properties listed below.
}
```

<a id="canonical-0212311011121230-1102112320023101-1123113013132120-2321113322031011-1113113301122303-3232023212001201-0021100111110110-2223202323102332"></a>

## Direct properties — enhanced_firewall_policies / 310303001312 / 3

<a id="canonical-0201032031003300-1221302300302013-2232300003231010-1233022122220330-0221020013303231-1010321303301100-1111022330112030-0221010022010011"></a>

<a id="canonical-2201323031231033-2210013231121311-0200000022110200-2201123001330122-2111130000220132-2021120002213230-0123003120332313-2132203202212323"></a>

## name property — enhanced_firewall_policies / 310303001312 / 4

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

<a id="canonical-1012123300312120-2203100130103212-2203120221010331-3110111320031002-1312231210023212-2003231202211203-3211212131212203-1012311300203313"></a>

<a id="canonical-0021321122212233-1322100211001101-2202203130132301-1030121111230210-1200103113032112-3130000123132321-2332233112211222-1221103222012330"></a>

## namespace property — enhanced_firewall_policies / 310303001312 / 5

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

<a id="canonical-1103020120020333-0133011132210120-2302023101322013-0031232111133002-3333320130303111-2030233003312203-2213211233233203-0302221200121003"></a>

<a id="canonical-3111222021000120-3221301233020100-0121331030130121-2030020033222103-1212303133130010-1313231011301030-0122300303202002-0001022021020231"></a>

## tenant property — enhanced_firewall_policies / 310303001312 / 6

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

<a id="canonical-1102112023200100-3321222033331233-2112210203313311-0200202331303033-0201100331032311-1201120332032110-2313123221023320-3133020100210330"></a>

## Next pages — enhanced_firewall_policies / 310303001312 / 7

- [ingress_egress_gw.active_enhanced_firewall_policies](resources--gcp_vpc_site--reference--group-002.md#canonical-1301001123030321-3330020323110300-0010201202332110-0232111331033322-2223300110000312-0302201210302013-3313220231131310-2300313233100213)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-0131000100012312-2121310130001102-2202230331103203-2230222003223001-1113101010113111-2131313030021223-2030122333203101-2213001121001122)

<a id="canonical-2311231030030323-0303023002330311-2113312033103203-2330303331111123-2231102303100300-2002222211231230-3122033210231030-1103130311303323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1230002103310231-3302311200001212-0133302211133131-0130032032323030-0022203211100230-0213103211121220-2132212200123131-3130102201013022"></a>

## ingress_egress_gw.active_forward_proxy_policies — active_forward_proxy_policies / 310222121311 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-0131000100012312-2121310130001102-2202230331103203-2230222003223001-1113101010113111-2131313030021223-2030122333203101-2213001121001122)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-2101111213103332-1031003102012122-3033333002321300-0212010323021100-0301331300303302-3321221003203222-1230133302203013-3232231332123212)
- [ingress_egress_gw](resources--gcp_vpc_site--reference--group-001.md#canonical-3123100130121022-2313112013232222-0213021230302120-3230300230213031-3101322300112331-0011013220032322-1311222312201210-2213020302010100)
- ingress_egress_gw.active_forward_proxy_policies

<a id="canonical-0332012032001200-3323002331020200-3302001010020303-2333021313131100-2002223302322031-2232112013300310-2102031102311101-0222300222211230"></a>

Type: `"object"`. single nested block, Optional.

Ordered List of Forward Proxy Policies active.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("forward_proxy_policies")}
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
active_forward_proxy_policies {
  # Configure direct properties listed below.
}
```

<a id="canonical-1000133231112032-1032213231312032-3322333003110301-1312132311322332-3121213200001112-3133310122021112-1310203032100000-1101232230322122"></a>

## Direct properties — active_forward_proxy_policies / 310222121311 / 3

- [forward_proxy_policies](resources--gcp_vpc_site--reference--group-002.md#canonical-3133222033121201-2303302021230103-2132130233111130-0212233200000231-2202001210130120-3312133203221231-2032322200003210-2203320103220133): complete subsection reference.

<a id="canonical-2230232020323131-2122003302030322-1333202300232013-0312123101110033-2322100033211323-2222133032220222-1011123103320331-0213332220231332"></a>

## Next pages — active_forward_proxy_policies / 310222121311 / 4

- [ingress_egress_gw.active_forward_proxy_policies.forward_proxy_policies](resources--gcp_vpc_site--reference--group-002.md#canonical-3133222033121201-2303302021230103-2132130233111130-0212233200000231-2202001210130120-3312133203221231-2032322200003210-2203320103220133)
- [ingress_egress_gw](resources--gcp_vpc_site--reference--group-001.md#canonical-3123100130121022-2313112013232222-0213021230302120-3230300230213031-3101322300112331-0011013220032322-1311222312201210-2213020302010100)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-0131000100012312-2121310130001102-2202230331103203-2230222003223001-1113101010113111-2131313030021223-2030122333203101-2213001121001122)

<a id="canonical-3133222033121201-2303302021230103-2132130233111130-0212233200000231-2202001210130120-3312133203221231-2032322200003210-2203320103220133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2203000121330110-3022302302211212-2102122110332133-1200221133222122-2310233222100023-3012132312121222-3123031010121131-1122303320203132"></a>

## ingress_egress_gw.active_forward_proxy_policies.forward_proxy_policies — forward_proxy_policies / 102203030023 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-0131000100012312-2121310130001102-2202230331103203-2230222003223001-1113101010113111-2131313030021223-2030122333203101-2213001121001122)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-2101111213103332-1031003102012122-3033333002321300-0212010323021100-0301331300303302-3321221003203222-1230133302203013-3232231332123212)
- [ingress_egress_gw](resources--gcp_vpc_site--reference--group-001.md#canonical-3123100130121022-2313112013232222-0213021230302120-3230300230213031-3101322300112331-0011013220032322-1311222312201210-2213020302010100)
- [ingress_egress_gw.active_forward_proxy_policies](resources--gcp_vpc_site--reference--group-002.md#canonical-2311231030030323-0303023002330311-2113312033103203-2330303331111123-2231102303100300-2002222211231230-3122033210231030-1103130311303323)
- ingress_egress_gw.active_forward_proxy_policies.forward_proxy_policies

<a id="canonical-3200013202113113-2230122330013220-1030201130320103-1103020032213200-1203011332201122-0301332103330332-2202320210013220-2302332003031232"></a>

Type: `"object"`. list nested block, Optional.

Ordered List of Forward Proxy Policies active.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("name")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 128,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 128,
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
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

Terraform syntax:

```terraform
forward_proxy_policies {
  # Configure direct properties listed below.
}
```

<a id="canonical-2203021301300003-0003032322222332-3310131020231010-1223332322210313-2201100222213031-0301200301310133-3132312222211022-0130011230300310"></a>

## Direct properties — forward_proxy_policies / 102203030023 / 3

<a id="canonical-0331113031323311-0302230323121120-2213023100102233-0011331333112332-1232311320303130-3231220331211232-2321100213312321-1222102122232302"></a>

<a id="canonical-0120131003202303-3033333332112133-3032102112310133-0102321103002300-2302001203200323-1010332130012020-1231133013010012-2111000113332322"></a>

## name property — forward_proxy_policies / 102203030023 / 4

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

<a id="canonical-3232103212300032-0221100302111201-2301210223120100-0022012022000102-1202310030031133-1303131012031213-0301202100110011-3230123000100300"></a>

<a id="canonical-1002233013230021-1130231303102111-3102302111032101-0101210232033103-2103033110012333-2003000032232011-0023303121331223-0200102013003100"></a>

## namespace property — forward_proxy_policies / 102203030023 / 5

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

<a id="canonical-1021102023131202-2003121213332013-0203031230211021-3320221303330120-1020233221320003-3123002101002233-2202020023130201-2032121003132333"></a>

<a id="canonical-1000323103213001-3332221013021212-1023200300211003-2033332332333001-3233103302001333-3200010221123132-2323113322130311-3130120223323001"></a>

## tenant property — forward_proxy_policies / 102203030023 / 6

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

<a id="canonical-1210020032123013-1123121132201222-2210310130033210-3021301300102123-0111233313302302-3002221200231110-2301311230311102-3311002101122020"></a>

## Next pages — forward_proxy_policies / 102203030023 / 7

- [ingress_egress_gw.active_forward_proxy_policies](resources--gcp_vpc_site--reference--group-002.md#canonical-2311231030030323-0303023002330311-2113312033103203-2330303331111123-2231102303100300-2002222211231230-3122033210231030-1103130311303323)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-0131000100012312-2121310130001102-2202230331103203-2230222003223001-1113101010113111-2131313030021223-2030122333203101-2213001121001122)

<a id="canonical-0212013333022211-1001302012023020-3122102202102122-3103121111332212-0021111032333010-0002012012010100-1220132220131102-2001123001010132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2303022030003211-1010123303321233-2030322203213232-1322301022231311-1110103330101122-1231323333121212-1333002223312113-3012210312210232"></a>

## ingress_egress_gw.active_network_policies — active_network_policies / 032010123220 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-0131000100012312-2121310130001102-2202230331103203-2230222003223001-1113101010113111-2131313030021223-2030122333203101-2213001121001122)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-2101111213103332-1031003102012122-3033333002321300-0212010323021100-0301331300303302-3321221003203222-1230133302203013-3232231332123212)
- [ingress_egress_gw](resources--gcp_vpc_site--reference--group-001.md#canonical-3123100130121022-2313112013232222-0213021230302120-3230300230213031-3101322300112331-0011013220032322-1311222312201210-2213020302010100)
- ingress_egress_gw.active_network_policies

<a id="canonical-2220200022201113-3011311120102223-0130202031103110-1033321231123130-2013113113011310-1131322021003020-3332222311223131-1011000203002013"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for active network policies.

Upstream description:

List of firewall policy views.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("network_policies")}
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
active_network_policies {
  # Configure direct properties listed below.
}
```

<a id="canonical-1232220332220210-2033111223111300-3020103321113310-0001001003103231-2323013020011001-3033233101133303-1001231000302301-0003101312033202"></a>

## Direct properties — active_network_policies / 032010123220 / 3

- [network_policies](resources--gcp_vpc_site--reference--group-002.md#canonical-0031111202020113-1101321003311103-2133321112320320-0033002103220300-1003111201113202-3112321030113331-2202313312133011-3133020120230002): complete subsection reference.

<a id="canonical-1113233110310012-3102200112232103-0212321230233012-0213301201203303-2200020111002132-0123133312023331-1030222213330111-3100103121223323"></a>

## Next pages — active_network_policies / 032010123220 / 4

- [ingress_egress_gw.active_network_policies.network_policies](resources--gcp_vpc_site--reference--group-002.md#canonical-0031111202020113-1101321003311103-2133321112320320-0033002103220300-1003111201113202-3112321030113331-2202313312133011-3133020120230002)
- [ingress_egress_gw](resources--gcp_vpc_site--reference--group-001.md#canonical-3123100130121022-2313112013232222-0213021230302120-3230300230213031-3101322300112331-0011013220032322-1311222312201210-2213020302010100)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-0131000100012312-2121310130001102-2202230331103203-2230222003223001-1113101010113111-2131313030021223-2030122333203101-2213001121001122)

<a id="canonical-0031111202020113-1101321003311103-2133321112320320-0033002103220300-1003111201113202-3112321030113331-2202313312133011-3133020120230002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1233120023212012-3313023332210132-1231320221021331-3132001010021200-3010000031333032-1112322331033121-3201100312323332-2023300320013011"></a>

## ingress_egress_gw.active_network_policies.network_policies — network_policies / 213000112213 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-0131000100012312-2121310130001102-2202230331103203-2230222003223001-1113101010113111-2131313030021223-2030122333203101-2213001121001122)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-2101111213103332-1031003102012122-3033333002321300-0212010323021100-0301331300303302-3321221003203222-1230133302203013-3232231332123212)
- [ingress_egress_gw](resources--gcp_vpc_site--reference--group-001.md#canonical-3123100130121022-2313112013232222-0213021230302120-3230300230213031-3101322300112331-0011013220032322-1311222312201210-2213020302010100)
- [ingress_egress_gw.active_network_policies](resources--gcp_vpc_site--reference--group-002.md#canonical-0212013333022211-1001302012023020-3122102202102122-3103121111332212-0021111032333010-0002012012010100-1220132220131102-2001123001010132)
- ingress_egress_gw.active_network_policies.network_policies

<a id="canonical-3211203032233303-1312101213323132-3103000223332010-0203121032311103-1133022121302033-0103203333123203-0033303123012302-1221203121113111"></a>

Type: `"object"`. list nested block, Optional.

Ordered List of Firewall Policies active for this network firewall.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("name")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 128,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 128,
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
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

Terraform syntax:

```terraform
network_policies {
  # Configure direct properties listed below.
}
```

<a id="canonical-3322033322022000-3003111332022033-1311000131102131-3200203211320313-3002312313230100-3102300003210123-2310022320102302-2130322101332022"></a>

## Direct properties — network_policies / 213000112213 / 3

<a id="canonical-0111022122330223-0130012111031011-0130201123231132-3302120203110221-0333230102232203-2303123121123122-0200033232330002-2232000320223121"></a>

<a id="canonical-0002023031110231-1320223002333033-3330112300311123-2303123110323221-0100120331200031-2131121321112221-2230220022022322-0320333303313333"></a>

## name property — network_policies / 213000112213 / 4

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

<a id="canonical-0122111121301330-2322033222320031-1021233101222112-1231133333110332-3012110132011100-0333212023212312-3221003232133201-1021013011113113"></a>

<a id="canonical-0111112230332210-1013311321132021-2001332121331313-0000100023121030-0100010031320030-3131112333122033-2010002023110000-2233020000302222"></a>

## namespace property — network_policies / 213000112213 / 5

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

<a id="canonical-0202300221000013-0001102023002211-0333110203002123-3003331330200021-1333032022212101-1212210213212330-2233111021101111-2122210233303103"></a>

<a id="canonical-2102112221122011-1312232030203200-2011030332123230-0212210320201002-0203122102221202-1210011123121013-1303031313213003-1222010032132110"></a>

## tenant property — network_policies / 213000112213 / 6

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

<a id="canonical-2122222232130102-1223022233123113-3000210110103202-0313223113310300-1010003331213331-1210210333111113-3301230313121021-1222332033001123"></a>

## Next pages — network_policies / 213000112213 / 7

- [ingress_egress_gw.active_network_policies](resources--gcp_vpc_site--reference--group-002.md#canonical-0212013333022211-1001302012023020-3122102202102122-3103121111332212-0021111032333010-0002012012010100-1220132220131102-2001123001010132)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-0131000100012312-2121310130001102-2202230331103203-2230222003223001-1113101010113111-2131313030021223-2030122333203101-2213001121001122)

<a id="canonical-0211003021310313-1300012031202133-2303213320131031-0133313313000021-2130221202103102-3201121233220023-1101033302303113-3120013303232300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2331022033031333-3233212123022200-1232030122202113-3232333322232332-0103101321012012-1002103323030222-3302102332002011-1003021201311321"></a>

## ingress_egress_gw.dc_cluster_group_inside_vn — dc_cluster_group_inside_vn / 132320021303 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-0131000100012312-2121310130001102-2202230331103203-2230222003223001-1113101010113111-2131313030021223-2030122333203101-2213001121001122)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-2101111213103332-1031003102012122-3033333002321300-0212010323021100-0301331300303302-3321221003203222-1230133302203013-3232231332123212)
- [ingress_egress_gw](resources--gcp_vpc_site--reference--group-001.md#canonical-3123100130121022-2313112013232222-0213021230302120-3230300230213031-3101322300112331-0011013220032322-1311222312201210-2213020302010100)
- ingress_egress_gw.dc_cluster_group_inside_vn

<a id="canonical-1100013301000323-0020330133001213-2320323001331120-0011221310033133-0111022021112033-2310022213130132-0021313213313132-2023113021333303"></a>

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
dc_cluster_group_inside_vn {
  # Configure direct properties listed below.
}
```

<a id="canonical-3113310012100132-1211100032022003-0102211233221020-3113313332021122-2112012002231021-3102021200222220-1030300123123222-2212113003001102"></a>

## Direct properties — dc_cluster_group_inside_vn / 132320021303 / 3

<a id="canonical-0000101232202101-0033232022203230-2333133131031103-1033333222310212-1223121201001300-2131031222231130-1312311310320320-1311221033030101"></a>

<a id="canonical-0312030231132322-3300031200030322-1021323011213010-3222111033021232-2123002100120013-0301203213220021-2213100313020230-2313233301003222"></a>

## name property — dc_cluster_group_inside_vn / 132320021303 / 4

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

<a id="canonical-0013111113313320-3310111221122302-2202023221202313-2302332301022330-0013330030302110-0013031303110330-1222223313020122-0033213302013022"></a>

<a id="canonical-0132011311210131-3021232203003333-3032130221333333-1113230301000031-0303023011212220-2122202110331011-3303011312002223-3302200010222130"></a>

## namespace property — dc_cluster_group_inside_vn / 132320021303 / 5

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

<a id="canonical-1220102300233212-0033102322103310-3110200212033121-1230001002231031-1112233232113020-2300302010202132-3102021322010220-3223121231021131"></a>

<a id="canonical-0330332133032023-2200121012032200-0101211112120131-1321323231202032-3311230133301111-2212230122111130-3122320330122011-1333131303112110"></a>

## tenant property — dc_cluster_group_inside_vn / 132320021303 / 6

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

<a id="canonical-3203302120121201-0301332230330221-1111130333232313-2101012231023102-0023312211030123-2113013302313032-3232021003322112-3113033101311211"></a>

## Next pages — dc_cluster_group_inside_vn / 132320021303 / 7

- [ingress_egress_gw](resources--gcp_vpc_site--reference--group-001.md#canonical-3123100130121022-2313112013232222-0213021230302120-3230300230213031-3101322300112331-0011013220032322-1311222312201210-2213020302010100)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-0131000100012312-2121310130001102-2202230331103203-2230222003223001-1113101010113111-2131313030021223-2030122333203101-2213001121001122)

<a id="canonical-2103132110012023-3322010132201321-3211232011002221-1330321223032322-2022233110013331-3121223121223102-1103013001232203-0303330031101022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0122222032023311-3332031122101211-0220303332231020-0300322023022101-0233010112231003-1202302133203112-3102310230303300-1002331100233320"></a>

## ingress_egress_gw.dc_cluster_group_outside_vn — dc_cluster_group_outside_vn / 221320102032 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-0131000100012312-2121310130001102-2202230331103203-2230222003223001-1113101010113111-2131313030021223-2030122333203101-2213001121001122)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-2101111213103332-1031003102012122-3033333002321300-0212010323021100-0301331300303302-3321221003203222-1230133302203013-3232231332123212)
- [ingress_egress_gw](resources--gcp_vpc_site--reference--group-001.md#canonical-3123100130121022-2313112013232222-0213021230302120-3230300230213031-3101322300112331-0011013220032322-1311222312201210-2213020302010100)
- ingress_egress_gw.dc_cluster_group_outside_vn

<a id="canonical-0102302312122200-0130122133223222-0301013031112303-1022122210320131-1311212023102213-2101332101103000-2221213310011022-1133103200213333"></a>

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
dc_cluster_group_outside_vn {
  # Configure direct properties listed below.
}
```

<a id="canonical-1313100111313121-0111201232310000-3223212020032310-3303022102202333-0301231010013012-0200000332023103-3030230011121202-3120203101110112"></a>

## Direct properties — dc_cluster_group_outside_vn / 221320102032 / 3

<a id="canonical-0131013121320031-1021330003231201-1321310010323212-2300212032010030-1202213001032133-0321200323001211-2132022122113302-2113333321012033"></a>

<a id="canonical-1231020221333312-2211123121201230-2301300212120302-3022122013110312-1121200203211213-3301010122130323-3101021303122120-1120030121131231"></a>

## name property — dc_cluster_group_outside_vn / 221320102032 / 4

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

<a id="canonical-1010002130330031-2232333321110213-3002332001001010-1331203303003311-3310031321312112-0010010131232030-0020110100113321-1232032013101202"></a>

<a id="canonical-3212011301210022-2302333303333210-1103002211012231-0201020111020123-0020300321101320-2320311311033333-2003130212321211-0203323323100203"></a>

## namespace property — dc_cluster_group_outside_vn / 221320102032 / 5

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

<a id="canonical-2130012201132203-1323110111213010-3212223302030333-0332220231202332-0021331212233330-0122111021203333-0131110331231320-2301222222033230"></a>

<a id="canonical-3312001222101033-2021030033230233-3212200033122323-1113320110023110-1101133120233132-1113220320311013-0223210320133120-1000330223313201"></a>

## tenant property — dc_cluster_group_outside_vn / 221320102032 / 6

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

<a id="canonical-1310022333210131-1302010220303011-3021300303200021-2021322212320110-2101013013013032-1003313323123103-1023021203132311-3310223020122320"></a>

## Next pages — dc_cluster_group_outside_vn / 221320102032 / 7

- [ingress_egress_gw](resources--gcp_vpc_site--reference--group-001.md#canonical-3123100130121022-2313112013232222-0213021230302120-3230300230213031-3101322300112331-0011013220032322-1311222312201210-2213020302010100)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-0131000100012312-2121310130001102-2202230331103203-2230222003223001-1113101010113111-2131313030021223-2030122333203101-2213001121001122)

<a id="canonical-2011302000102201-3232023320300020-0303212120230300-3013012121333131-0322032223310021-3021011110023032-1022201123121123-1331123020220333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1120232012332021-2221330012013031-0311212331311231-1101233120313032-3221000100011111-3123332201122000-2022012101332010-1313323032230002"></a>

## ingress_egress_gw.forward_proxy_allow_all — forward_proxy_allow_all / 020232030000 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-0131000100012312-2121310130001102-2202230331103203-2230222003223001-1113101010113111-2131313030021223-2030122333203101-2213001121001122)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-2101111213103332-1031003102012122-3033333002321300-0212010323021100-0301331300303302-3321221003203222-1230133302203013-3232231332123212)
- [ingress_egress_gw](resources--gcp_vpc_site--reference--group-001.md#canonical-3123100130121022-2313112013232222-0213021230302120-3230300230213031-3101322300112331-0011013220032322-1311222312201210-2213020302010100)
- ingress_egress_gw.forward_proxy_allow_all

<a id="canonical-1221312311231120-1300332032323121-0111021030023320-3013313102203111-1020231100321023-0210103213111301-0201113031230001-0321101131231311"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for forward proxy allow all.

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
forward_proxy_allow_all = {}
```

<a id="canonical-0122211331211301-2020133213333332-2013120013303003-1203311023011311-0021020221313112-2221201031113120-3312021212321012-3233220221300102"></a>

## Direct properties — forward_proxy_allow_all / 020232030000 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2200111300020020-3012322020320320-2221333222232301-2202330212021330-1010210323120220-2021111331031030-0330000013121230-0202201212200100"></a>

## Next pages — forward_proxy_allow_all / 020232030000 / 4

- [ingress_egress_gw](resources--gcp_vpc_site--reference--group-001.md#canonical-3123100130121022-2313112013232222-0213021230302120-3230300230213031-3101322300112331-0011013220032322-1311222312201210-2213020302010100)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-0131000100012312-2121310130001102-2202230331103203-2230222003223001-1113101010113111-2131313030021223-2030122333203101-2213001121001122)

<a id="canonical-1131133100211321-0133332003213130-3001023120110100-2112200313000011-1100333310030310-0010101033020232-0203330321121321-3012232113212212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0123120032210233-2312001313120131-3030031002010101-2102230220002133-0022111322210020-3321011311031212-0122210200101112-1031102130002102"></a>

## ingress_egress_gw.global_network_list — global_network_list / 120303221122 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-0131000100012312-2121310130001102-2202230331103203-2230222003223001-1113101010113111-2131313030021223-2030122333203101-2213001121001122)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-2101111213103332-1031003102012122-3033333002321300-0212010323021100-0301331300303302-3321221003203222-1230133302203013-3232231332123212)
- [ingress_egress_gw](resources--gcp_vpc_site--reference--group-001.md#canonical-3123100130121022-2313112013232222-0213021230302120-3230300230213031-3101322300112331-0011013220032322-1311222312201210-2213020302010100)
- ingress_egress_gw.global_network_list

<a id="canonical-3312020302221223-3233230233221013-1003302301021032-2102331210333330-3000322301230103-3111033220210302-0213333213102021-2012002203000303"></a>

Type: `"object"`. single nested block, Optional.

Global Network Connection List. List of global network connections.

Upstream description:

List of global network connections.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("global_network_connections")}
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
global_network_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-1310203233122220-2231321010113213-0111023001000311-1201230010322222-1231312302313331-2233310101112011-1303003220310021-3310002031210323"></a>

## Direct properties — global_network_list / 120303221122 / 3

- [global_network_connections](resources--gcp_vpc_site--reference--group-002.md#canonical-1110212202030220-3031323212301211-0222110300331222-2222131313230133-1133210212001220-1211230303131323-0012020231030221-0121301130002121): complete subsection reference.

<a id="canonical-1023223313032323-3333132111220222-1110021121312102-0202231021112212-3112012310000002-3320322033003211-2112131022210202-2202311221113232"></a>

## Next pages — global_network_list / 120303221122 / 4

- [ingress_egress_gw.global_network_list.global_network_connections](resources--gcp_vpc_site--reference--group-002.md#canonical-1110212202030220-3031323212301211-0222110300331222-2222131313230133-1133210212001220-1211230303131323-0012020231030221-0121301130002121)
- [ingress_egress_gw](resources--gcp_vpc_site--reference--group-001.md#canonical-3123100130121022-2313112013232222-0213021230302120-3230300230213031-3101322300112331-0011013220032322-1311222312201210-2213020302010100)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-0131000100012312-2121310130001102-2202230331103203-2230222003223001-1113101010113111-2131313030021223-2030122333203101-2213001121001122)

<a id="canonical-1110212202030220-3031323212301211-0222110300331222-2222131313230133-1133210212001220-1211230303131323-0012020231030221-0121301130002121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0101333111203220-3231313231023123-1020302013323013-1203030012233321-3320310021110300-1222013012323002-3103032013033030-3012120123223020"></a>

## ingress_egress_gw.global_network_list.global_network_connections — global_network_connections / 132201222303 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-0131000100012312-2121310130001102-2202230331103203-2230222003223001-1113101010113111-2131313030021223-2030122333203101-2213001121001122)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-2101111213103332-1031003102012122-3033333002321300-0212010323021100-0301331300303302-3321221003203222-1230133302203013-3232231332123212)
- [ingress_egress_gw](resources--gcp_vpc_site--reference--group-001.md#canonical-3123100130121022-2313112013232222-0213021230302120-3230300230213031-3101322300112331-0011013220032322-1311222312201210-2213020302010100)
- [ingress_egress_gw.global_network_list](resources--gcp_vpc_site--reference--group-002.md#canonical-1131133100211321-0133332003213130-3001023120110100-2112200313000011-1100333310030310-0010101033020232-0203330321121321-3012232113212212)
- ingress_egress_gw.global_network_list.global_network_connections

<a id="canonical-2202111301201000-0231131220223032-2332033001022031-1111112303303220-1331023030303003-0022010330032011-1110112300203011-1213120111001321"></a>

Type: `"object"`. list nested block, Optional.

Global Network Connections. Global network connections.

Upstream description:

Global network connections.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("sli_to_global_dr",
    "slo_to_global_dr")}
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
global_network_connections {
  # Configure direct properties listed below.
}
```

<a id="canonical-1020222203121021-0021011020331132-2100311223033100-0012103002203122-1200223313311202-1233321311022330-2201133110213203-2110311020200323"></a>

## Direct properties — global_network_connections / 132201222303 / 3

- [sli_to_global_dr](resources--gcp_vpc_site--reference--group-002.md#canonical-0110321302200003-1032031112011321-0120203001301132-1010201302323322-3011230311011323-1131303213312200-1123203110020133-3300132213331310): complete subsection reference.

- [slo_to_global_dr](resources--gcp_vpc_site--reference--group-002.md#canonical-2231211013211002-2311211020332223-3113311221212000-2213012222032211-2112021132203101-1320120001320113-0022230001233133-0013023202133303): complete subsection reference.

<a id="canonical-0222331023113101-2201032130101113-2010231003310022-2220213121032312-2303011321012100-0302100231101031-0031113111300111-1211002121333222"></a>

## Next pages — global_network_connections / 132201222303 / 4

- [ingress_egress_gw.global_network_list.global_network_connections.sli_to_global_dr](resources--gcp_vpc_site--reference--group-002.md#canonical-0110321302200003-1032031112011321-0120203001301132-1010201302323322-3011230311011323-1131303213312200-1123203110020133-3300132213331310)
- [ingress_egress_gw.global_network_list.global_network_connections.slo_to_global_dr](resources--gcp_vpc_site--reference--group-002.md#canonical-2231211013211002-2311211020332223-3113311221212000-2213012222032211-2112021132203101-1320120001320113-0022230001233133-0013023202133303)
- [ingress_egress_gw.global_network_list](resources--gcp_vpc_site--reference--group-002.md#canonical-1131133100211321-0133332003213130-3001023120110100-2112200313000011-1100333310030310-0010101033020232-0203330321121321-3012232113212212)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-0131000100012312-2121310130001102-2202230331103203-2230222003223001-1113101010113111-2131313030021223-2030122333203101-2213001121001122)

<a id="canonical-0110321302200003-1032031112011321-0120203001301132-1010201302323322-3011230311011323-1131303213312200-1123203110020133-3300132213331310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3333313020113201-3113002132332210-3120131010030003-1033323113322201-3100013032012202-1023222001122211-3003313302301222-3011020132322021"></a>

## ingress_egress_gw.global_network_list.global_network_connections.sli_to_global_dr — sli_to_global_dr / 232132300130 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-0131000100012312-2121310130001102-2202230331103203-2230222003223001-1113101010113111-2131313030021223-2030122333203101-2213001121001122)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-2101111213103332-1031003102012122-3033333002321300-0212010323021100-0301331300303302-3321221003203222-1230133302203013-3232231332123212)
- [ingress_egress_gw](resources--gcp_vpc_site--reference--group-001.md#canonical-3123100130121022-2313112013232222-0213021230302120-3230300230213031-3101322300112331-0011013220032322-1311222312201210-2213020302010100)
- [ingress_egress_gw.global_network_list](resources--gcp_vpc_site--reference--group-002.md#canonical-1131133100211321-0133332003213130-3001023120110100-2112200313000011-1100333310030310-0010101033020232-0203330321121321-3012232113212212)
- [ingress_egress_gw.global_network_list.global_network_connections](resources--gcp_vpc_site--reference--group-002.md#canonical-1110212202030220-3031323212301211-0222110300331222-2222131313230133-1133210212001220-1211230303131323-0012020231030221-0121301130002121)
- ingress_egress_gw.global_network_list.global_network_connections.sli_to_global_dr

<a id="canonical-1233223030331133-1100230033233331-1001233013013133-1003012020021032-2122130103312313-1010301213100133-3120220010232101-1112231033302002"></a>

Type: `"object"`. single nested block, Optional.

Global network reference for direct connection.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
sli_to_global_dr {
  # Configure direct properties listed below.
}
```

<a id="canonical-0111202211230202-3000202022132012-1002211202313000-1321021312330100-1023021223320103-2311233000131031-3102223133210003-3102122023100110"></a>

## Direct properties — sli_to_global_dr / 232132300130 / 3

- [global_vn](resources--gcp_vpc_site--reference--group-002.md#canonical-1131330010202133-1212333012120132-2303220322101013-3102220233032120-0010123020302010-2222211330233320-3230111031133113-0133130123100320): complete subsection reference.

<a id="canonical-1100130220023201-0233120221000013-0013122033303011-3013220120200313-3011022000103130-1231331322113302-1012111211313133-0222010333022110"></a>

## Next pages — sli_to_global_dr / 232132300130 / 4

- [ingress_egress_gw.global_network_list.global_network_connections.sli_to_global_dr.global_vn](resources--gcp_vpc_site--reference--group-002.md#canonical-1131330010202133-1212333012120132-2303220322101013-3102220233032120-0010123020302010-2222211330233320-3230111031133113-0133130123100320)
- [ingress_egress_gw.global_network_list.global_network_connections](resources--gcp_vpc_site--reference--group-002.md#canonical-1110212202030220-3031323212301211-0222110300331222-2222131313230133-1133210212001220-1211230303131323-0012020231030221-0121301130002121)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-0131000100012312-2121310130001102-2202230331103203-2230222003223001-1113101010113111-2131313030021223-2030122333203101-2213001121001122)

<a id="canonical-1131330010202133-1212333012120132-2303220322101013-3102220233032120-0010123020302010-2222211330233320-3230111031133113-0133130123100320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0101232230020313-0012210121122222-0311203112010013-0111203101113022-1222302103110110-1301113000121300-0220021311130000-2130013033323031"></a>

## ingress_egress_gw.global_network_list.global_network_connections.sli_to_global_dr.global_vn — global_vn / 322002012332 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-0131000100012312-2121310130001102-2202230331103203-2230222003223001-1113101010113111-2131313030021223-2030122333203101-2213001121001122)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-2101111213103332-1031003102012122-3033333002321300-0212010323021100-0301331300303302-3321221003203222-1230133302203013-3232231332123212)
- [ingress_egress_gw](resources--gcp_vpc_site--reference--group-001.md#canonical-3123100130121022-2313112013232222-0213021230302120-3230300230213031-3101322300112331-0011013220032322-1311222312201210-2213020302010100)
- [ingress_egress_gw.global_network_list](resources--gcp_vpc_site--reference--group-002.md#canonical-1131133100211321-0133332003213130-3001023120110100-2112200313000011-1100333310030310-0010101033020232-0203330321121321-3012232113212212)
- [ingress_egress_gw.global_network_list.global_network_connections](resources--gcp_vpc_site--reference--group-002.md#canonical-1110212202030220-3031323212301211-0222110300331222-2222131313230133-1133210212001220-1211230303131323-0012020231030221-0121301130002121)
- [ingress_egress_gw.global_network_list.global_network_connections.sli_to_global_dr](resources--gcp_vpc_site--reference--group-002.md#canonical-0110321302200003-1032031112011321-0120203001301132-1010201302323322-3011230311011323-1131303213312200-1123203110020133-3300132213331310)
- ingress_egress_gw.global_network_list.global_network_connections.sli_to_global_dr.global_vn

<a id="canonical-0300221312010132-1322212103213130-3221200232130010-3300021122310331-1221211231102011-1222021020101322-1210113033212310-3200310223310101"></a>

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
global_vn {
  # Configure direct properties listed below.
}
```

<a id="canonical-0032210021203221-3120223213200232-2031333233011312-1120110001312022-2013003030312333-0002233013032000-2033332200002330-3122012200213033"></a>

## Direct properties — global_vn / 322002012332 / 3

<a id="canonical-3311120220212112-0010221201133212-0000120210221010-1123302303223322-3202012032332231-2223203011110200-2310320130231032-3000122101212201"></a>

<a id="canonical-1131011200002223-1232303103202102-0132332310210122-3310213121033130-0232223032231120-0032020322320001-2130212232020102-3100133102003001"></a>

## name property — global_vn / 322002012332 / 4

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

<a id="canonical-3202301133200032-2201203030202231-1323012233113301-1311120121030032-1132321031312032-2022201311202033-0122011013333101-0020303001303331"></a>

<a id="canonical-1110210031313322-1211313222201311-2013130303132222-3321013211322112-1121133231110232-1300122202122210-0211011022002210-1013212033203302"></a>

## namespace property — global_vn / 322002012332 / 5

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

<a id="canonical-0103323230133113-2132332210202303-1022022133031312-1030222210312002-3033223022122003-0301202102231211-1023110303122313-2310320023131113"></a>

<a id="canonical-2233032310130321-2230031320333013-0123202012211122-1101230312012232-1220221330000003-2221201103311030-1231303321210200-0310132012333112"></a>

## tenant property — global_vn / 322002012332 / 6

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

<a id="canonical-0131021201333131-0232102222013323-3331221213103320-1221233303003132-2330210322313021-0103302021120322-2331222003100011-0332231013232212"></a>

## Next pages — global_vn / 322002012332 / 7

- [ingress_egress_gw.global_network_list.global_network_connections.sli_to_global_dr](resources--gcp_vpc_site--reference--group-002.md#canonical-0110321302200003-1032031112011321-0120203001301132-1010201302323322-3011230311011323-1131303213312200-1123203110020133-3300132213331310)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-0131000100012312-2121310130001102-2202230331103203-2230222003223001-1113101010113111-2131313030021223-2030122333203101-2213001121001122)

<a id="canonical-2231211013211002-2311211020332223-3113311221212000-2213012222032211-2112021132203101-1320120001320113-0022230001233133-0013023202133303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0121131222210231-0220000010010030-3120313100010223-1120012222033023-3023101230012101-3200320013032001-2322121103121111-1322310201222001"></a>

## ingress_egress_gw.global_network_list.global_network_connections.slo_to_global_dr — slo_to_global_dr / 020220212232 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-0131000100012312-2121310130001102-2202230331103203-2230222003223001-1113101010113111-2131313030021223-2030122333203101-2213001121001122)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-2101111213103332-1031003102012122-3033333002321300-0212010323021100-0301331300303302-3321221003203222-1230133302203013-3232231332123212)
- [ingress_egress_gw](resources--gcp_vpc_site--reference--group-001.md#canonical-3123100130121022-2313112013232222-0213021230302120-3230300230213031-3101322300112331-0011013220032322-1311222312201210-2213020302010100)
- [ingress_egress_gw.global_network_list](resources--gcp_vpc_site--reference--group-002.md#canonical-1131133100211321-0133332003213130-3001023120110100-2112200313000011-1100333310030310-0010101033020232-0203330321121321-3012232113212212)
- [ingress_egress_gw.global_network_list.global_network_connections](resources--gcp_vpc_site--reference--group-002.md#canonical-1110212202030220-3031323212301211-0222110300331222-2222131313230133-1133210212001220-1211230303131323-0012020231030221-0121301130002121)
- ingress_egress_gw.global_network_list.global_network_connections.slo_to_global_dr

<a id="canonical-2111110123220130-2110322123032033-0223001310130311-0310222031331011-0120130032231032-1333123202232231-1100231313020300-1031313211130303"></a>

Type: `"object"`. single nested block, Optional.

Global network reference for direct connection.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
slo_to_global_dr {
  # Configure direct properties listed below.
}
```

<a id="canonical-0323221123210323-2311212120331103-1111203301223023-0010301211100003-1032101133302322-1331213013223012-1111221103122001-0032321113202100"></a>

## Direct properties — slo_to_global_dr / 020220212232 / 3

- [global_vn](resources--gcp_vpc_site--reference--group-002.md#canonical-0132132122302310-0023111333330211-2121201112030110-3010232010102113-2232020131222132-0120120133130232-2220033223001202-3331302300121031): complete subsection reference.

<a id="canonical-0200211030211332-1022230100333230-1102032113212032-2313312311300031-0302123023100202-2123200310022020-1213103031110023-0303221320203103"></a>

## Next pages — slo_to_global_dr / 020220212232 / 4

- [ingress_egress_gw.global_network_list.global_network_connections.slo_to_global_dr.global_vn](resources--gcp_vpc_site--reference--group-002.md#canonical-0132132122302310-0023111333330211-2121201112030110-3010232010102113-2232020131222132-0120120133130232-2220033223001202-3331302300121031)
- [ingress_egress_gw.global_network_list.global_network_connections](resources--gcp_vpc_site--reference--group-002.md#canonical-1110212202030220-3031323212301211-0222110300331222-2222131313230133-1133210212001220-1211230303131323-0012020231030221-0121301130002121)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-0131000100012312-2121310130001102-2202230331103203-2230222003223001-1113101010113111-2131313030021223-2030122333203101-2213001121001122)

<a id="canonical-0132132122302310-0023111333330211-2121201112030110-3010232010102113-2232020131222132-0120120133130232-2220033223001202-3331302300121031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3210333333101022-1120231112232123-0120212303303222-3001110102322223-1202122333201000-3132231023313101-3302202221232112-3230122020023020"></a>

## ingress_egress_gw.global_network_list.global_network_connections.slo_to_global_dr.global_vn — global_vn / 203010113123 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-0131000100012312-2121310130001102-2202230331103203-2230222003223001-1113101010113111-2131313030021223-2030122333203101-2213001121001122)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-2101111213103332-1031003102012122-3033333002321300-0212010323021100-0301331300303302-3321221003203222-1230133302203013-3232231332123212)
- [ingress_egress_gw](resources--gcp_vpc_site--reference--group-001.md#canonical-3123100130121022-2313112013232222-0213021230302120-3230300230213031-3101322300112331-0011013220032322-1311222312201210-2213020302010100)
- [ingress_egress_gw.global_network_list](resources--gcp_vpc_site--reference--group-002.md#canonical-1131133100211321-0133332003213130-3001023120110100-2112200313000011-1100333310030310-0010101033020232-0203330321121321-3012232113212212)
- [ingress_egress_gw.global_network_list.global_network_connections](resources--gcp_vpc_site--reference--group-002.md#canonical-1110212202030220-3031323212301211-0222110300331222-2222131313230133-1133210212001220-1211230303131323-0012020231030221-0121301130002121)
- [ingress_egress_gw.global_network_list.global_network_connections.slo_to_global_dr](resources--gcp_vpc_site--reference--group-002.md#canonical-2231211013211002-2311211020332223-3113311221212000-2213012222032211-2112021132203101-1320120001320113-0022230001233133-0013023202133303)
- ingress_egress_gw.global_network_list.global_network_connections.slo_to_global_dr.global_vn

<a id="canonical-2323303322313122-3220012300033111-3131301210220300-2121331212111023-3331013303101023-2112221031200010-1121213130031122-2202311230310331"></a>

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
global_vn {
  # Configure direct properties listed below.
}
```

<a id="canonical-1003332321222110-1100120212321010-3221220001302103-1010203010220111-1333010001310230-2110101320223232-0212303012331210-2230221001023120"></a>

## Direct properties — global_vn / 203010113123 / 3

<a id="canonical-0302011021031323-1221202111210210-1011230303230023-2100223332111300-0032303131032232-2032130011212321-2110213301202330-2300312023303310"></a>

<a id="canonical-3002313322023331-0313320320132023-2102313303123230-3312202100222321-0210330013233100-3321021022211330-3230120012013000-1233003032231230"></a>

## name property — global_vn / 203010113123 / 4

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

<a id="canonical-2230032020231310-3313320233113112-0023123030323222-1222213111202122-1123231322020023-2131011002233322-2323022313212210-3320331123310111"></a>

<a id="canonical-3002320120021311-0303230111110321-1033113230222320-2330221022233331-3310230001201133-2213110213032113-2312132022322131-0121333332123030"></a>

## namespace property — global_vn / 203010113123 / 5

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

<a id="canonical-3121332013013313-1231213322230321-1001100031220322-2311211311310030-3222002332021302-1010320331200133-2030022203333112-3332211120312012"></a>

<a id="canonical-3333230220031002-0303111302231210-3322212121333320-1330203222322022-3212003120011203-0322121103212102-2002031320030332-1201332323132310"></a>

## tenant property — global_vn / 203010113123 / 6

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

<a id="canonical-1312311023121210-3303223301020012-0031331202200132-2013102110202213-2220302320011130-0112110321132201-2010312000013312-0111032310330102"></a>

## Next pages — global_vn / 203010113123 / 7

- [ingress_egress_gw.global_network_list.global_network_connections.slo_to_global_dr](resources--gcp_vpc_site--reference--group-002.md#canonical-2231211013211002-2311211020332223-3113311221212000-2213012222032211-2112021132203101-1320120001320113-0022230001233133-0013023202133303)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-0131000100012312-2121310130001102-2202230331103203-2230222003223001-1113101010113111-2131313030021223-2030122333203101-2213001121001122)

<a id="canonical-2123301101220113-0123332320222111-0202203303002102-2012313302013323-3011301022113123-0301003102033320-2332330212323112-0113111321202221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0001323310012332-0333032233323031-3302222011301121-1222012201103220-0230012120010033-0003133312221032-1133023323211112-1133301022331113"></a>

## ingress_egress_gw.inside_network — inside_network / 003323301313 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-0131000100012312-2121310130001102-2202230331103203-2230222003223001-1113101010113111-2131313030021223-2030122333203101-2213001121001122)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-2101111213103332-1031003102012122-3033333002321300-0212010323021100-0301331300303302-3321221003203222-1230133302203013-3232231332123212)
- [ingress_egress_gw](resources--gcp_vpc_site--reference--group-001.md#canonical-3123100130121022-2313112013232222-0213021230302120-3230300230213031-3101322300112331-0011013220032322-1311222312201210-2213020302010100)
- ingress_egress_gw.inside_network

<a id="canonical-1220201310013012-2332112312123201-1011202121130233-1121133233210021-1123221333300131-0011123322033122-1033122022013233-3332302113303311"></a>

Type: `"object"`. single nested block, Optional.

Defines choice about GCP VPC network for a view.

Upstream description:

This defines choice about GCP VPC network for a view.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("existing_network",
    "new_network"),
  validators.ConflictingObjectAttributes("existing_network",
    "new_network_autogenerate"),
  validators.ConflictingObjectAttributes("new_network",
    "new_network_autogenerate")}
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
  "x-ves-oneof-field-choice": "[\"existing_network\",\"new_network\",\"new_network_autogenerate\"]"
}
```

Terraform syntax:

```terraform
inside_network {
  # Configure direct properties listed below.
}
```

<a id="canonical-3201011111312012-1132301112322323-3012230211013212-1300123302333332-2331330123011030-1001200011020100-1331020131133003-0021313032132022"></a>

## Direct properties — inside_network / 003323301313 / 3

- [existing_network](resources--gcp_vpc_site--reference--group-002.md#canonical-3303301001023010-1012133310330102-3133132102212000-0011131322220230-1302320222102202-0113230121132032-2220200211312311-0022020332301100): complete subsection reference.

- [new_network](resources--gcp_vpc_site--reference--group-002.md#canonical-2211023030213201-3102003310000300-2232010001202320-3302130031231322-2322033103113232-2220123330203113-3231100310210120-0133113231122203): complete subsection reference.

- [new_network_autogenerate](resources--gcp_vpc_site--reference--group-002.md#canonical-3103131310011102-0121323220112003-1023322321313021-3033102110112311-0032112112123220-0113111213230213-1232230013022132-3031203023101300): complete subsection reference.

<a id="canonical-3302210321023231-0322213100031010-0031202101012011-1233222210323032-1231210301213023-1221233023130100-1010301212323033-3202101222002110"></a>

## Next pages — inside_network / 003323301313 / 4

- [ingress_egress_gw.inside_network.existing_network](resources--gcp_vpc_site--reference--group-002.md#canonical-3303301001023010-1012133310330102-3133132102212000-0011131322220230-1302320222102202-0113230121132032-2220200211312311-0022020332301100)
- [ingress_egress_gw.inside_network.new_network](resources--gcp_vpc_site--reference--group-002.md#canonical-2211023030213201-3102003310000300-2232010001202320-3302130031231322-2322033103113232-2220123330203113-3231100310210120-0133113231122203)
- [ingress_egress_gw.inside_network.new_network_autogenerate](resources--gcp_vpc_site--reference--group-002.md#canonical-3103131310011102-0121323220112003-1023322321313021-3033102110112311-0032112112123220-0113111213230213-1232230013022132-3031203023101300)
- [ingress_egress_gw](resources--gcp_vpc_site--reference--group-001.md#canonical-3123100130121022-2313112013232222-0213021230302120-3230300230213031-3101322300112331-0011013220032322-1311222312201210-2213020302010100)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-0131000100012312-2121310130001102-2202230331103203-2230222003223001-1113101010113111-2131313030021223-2030122333203101-2213001121001122)

<a id="canonical-3303301001023010-1012133310330102-3133132102212000-0011131322220230-1302320222102202-0113230121132032-2220200211312311-0022020332301100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1300000321012200-1210001013131332-3002110220021231-0320303130223011-0213011321121313-2231301301332202-2310102123203031-1101323211031103"></a>

## ingress_egress_gw.inside_network.existing_network — existing_network / 123311302021 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-0131000100012312-2121310130001102-2202230331103203-2230222003223001-1113101010113111-2131313030021223-2030122333203101-2213001121001122)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-2101111213103332-1031003102012122-3033333002321300-0212010323021100-0301331300303302-3321221003203222-1230133302203013-3232231332123212)
- [ingress_egress_gw](resources--gcp_vpc_site--reference--group-001.md#canonical-3123100130121022-2313112013232222-0213021230302120-3230300230213031-3101322300112331-0011013220032322-1311222312201210-2213020302010100)
- [ingress_egress_gw.inside_network](resources--gcp_vpc_site--reference--group-002.md#canonical-2123301101220113-0123332320222111-0202203303002102-2012313302013323-3011301022113123-0301003102033320-2332330212323112-0113111321202221)
- ingress_egress_gw.inside_network.existing_network

<a id="canonical-0203312210210220-1000023000210202-2032310331331032-0223233223221021-2232110121210222-1113101123210003-3112010313122223-3030333133022111"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for existing network.

Upstream description:

Name of existing VPC network.

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
  },
  "x-ves-oneof-field-routing_type": "[]"
}
```

Terraform syntax:

```terraform
existing_network {
  # Configure direct properties listed below.
}
```

<a id="canonical-0223330322132223-3021220332312100-3023021123210232-1221031231133310-3211110231131003-0102322221232220-3221021223021112-0011031101110021"></a>

## Direct properties — existing_network / 123311302021 / 3

<a id="canonical-1102021222030201-2030112310300131-3212121101223110-1302221003000022-3302023010322131-3001002020110100-3312331323230033-2121031003333120"></a>

<a id="canonical-1013203200011001-3111002112010022-0313300011231330-0103232220032102-2110001122213330-2111323322122323-3320200000321021-2311301302203133"></a>

## name property — existing_network / 123311302021 / 4

Type: `"string"`. Optional.

GCP VPC Network Name. Name for your GCP VPC Network.

Upstream description:

Name for your GCP VPC Network.

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
    "maxLength": 64,
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

<a id="canonical-3202033123032021-3210032000133121-0110223100100100-2211110321032103-1010000100032011-3303302033132101-1003211033320211-3020022000302222"></a>

## Next pages — existing_network / 123311302021 / 5

- [ingress_egress_gw.inside_network](resources--gcp_vpc_site--reference--group-002.md#canonical-2123301101220113-0123332320222111-0202203303002102-2012313302013323-3011301022113123-0301003102033320-2332330212323112-0113111321202221)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-0131000100012312-2121310130001102-2202230331103203-2230222003223001-1113101010113111-2131313030021223-2030122333203101-2213001121001122)

<a id="canonical-2211023030213201-3102003310000300-2232010001202320-3302130031231322-2322033103113232-2220123330203113-3231100310210120-0133113231122203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0232133312122032-3121212310232313-3002230003222230-0122233010003202-1120000122031123-0013123121000200-0123201221300333-1112002222033002"></a>

## ingress_egress_gw.inside_network.new_network — new_network / 131133123333 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-0131000100012312-2121310130001102-2202230331103203-2230222003223001-1113101010113111-2131313030021223-2030122333203101-2213001121001122)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-2101111213103332-1031003102012122-3033333002321300-0212010323021100-0301331300303302-3321221003203222-1230133302203013-3232231332123212)
- [ingress_egress_gw](resources--gcp_vpc_site--reference--group-001.md#canonical-3123100130121022-2313112013232222-0213021230302120-3230300230213031-3101322300112331-0011013220032322-1311222312201210-2213020302010100)
- [ingress_egress_gw.inside_network](resources--gcp_vpc_site--reference--group-002.md#canonical-2123301101220113-0123332320222111-0202203303002102-2012313302013323-3011301022113123-0301003102033320-2332330212323112-0113111321202221)
- ingress_egress_gw.inside_network.new_network

<a id="canonical-2012233322012120-1333113001332212-0223203230123313-1110122223023301-2103123213332021-2120221231112332-1322323100131322-0021202212122021"></a>

Type: `"object"`. single nested block, Optional.

Parameters to create a new GCP VPC Network.

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
new_network {
  # Configure direct properties listed below.
}
```

<a id="canonical-1013131231000012-1310133303002311-1302213122113133-3201033012223010-0102132112030311-0311103223221031-0113032112302230-3321321302020011"></a>

## Direct properties — new_network / 131133123333 / 3

<a id="canonical-0203212311021121-0203011233003001-0032001100200010-0301302000321122-1032202130030302-1113320233033331-3201000230120000-2203311203022210"></a>

<a id="canonical-2001233003113030-3222310030131233-3030223100110331-2000330112031133-1310222311023032-0333332213213020-0223313011323301-1132221312313221"></a>

## name property — new_network / 131133123333 / 4

Type: `"string"`. Optional.

GCP VPC Network Name. Name for your GCP VPC Network.

Upstream description:

Name for your GCP VPC Network.

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
    "maxLength": 64,
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

<a id="canonical-2203031313130022-0332130033021211-0310130203220313-0312203103201212-1100330232320101-0211130033332032-1210232133033111-2321201232102123"></a>

## Next pages — new_network / 131133123333 / 5

- [ingress_egress_gw.inside_network](resources--gcp_vpc_site--reference--group-002.md#canonical-2123301101220113-0123332320222111-0202203303002102-2012313302013323-3011301022113123-0301003102033320-2332330212323112-0113111321202221)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-0131000100012312-2121310130001102-2202230331103203-2230222003223001-1113101010113111-2131313030021223-2030122333203101-2213001121001122)

<a id="canonical-3103131310011102-0121323220112003-1023322321313021-3033102110112311-0032112112123220-0113111213230213-1232230013022132-3031203023101300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0112122011212030-0202031221233122-0333310103210221-1302123300231101-3200312032132003-0200003203030112-0202031202033201-3230310130032232"></a>

## ingress_egress_gw.inside_network.new_network_autogenerate — new_network_autogenerate / 100112303103 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-0131000100012312-2121310130001102-2202230331103203-2230222003223001-1113101010113111-2131313030021223-2030122333203101-2213001121001122)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-2101111213103332-1031003102012122-3033333002321300-0212010323021100-0301331300303302-3321221003203222-1230133302203013-3232231332123212)
- [ingress_egress_gw](resources--gcp_vpc_site--reference--group-001.md#canonical-3123100130121022-2313112013232222-0213021230302120-3230300230213031-3101322300112331-0011013220032322-1311222312201210-2213020302010100)
- [ingress_egress_gw.inside_network](resources--gcp_vpc_site--reference--group-002.md#canonical-2123301101220113-0123332320222111-0202203303002102-2012313302013323-3011301022113123-0301003102033320-2332330212323112-0113111321202221)
- ingress_egress_gw.inside_network.new_network_autogenerate

<a id="canonical-3130011201120212-3020230122023232-3222133003302012-3110231310322201-2100130230313123-1323231103022003-2213021332102100-2122202303123321"></a>

Type: `["object", {}]`. Optional.

Create a new GCP VPC Network with autogenerated name.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
new_network_autogenerate = {}
```

<a id="canonical-3232312003131330-0000200331310020-1030321202021011-2222203221021131-2203313120232113-2322332120300121-0101222230022123-0133130003021220"></a>

## Direct properties — new_network_autogenerate / 100112303103 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0201233130000313-2013301202331202-0330321232033322-2332312230020022-3002113100033301-3100222303032301-3003013011322101-1023231202331233"></a>

## Next pages — new_network_autogenerate / 100112303103 / 4

- [ingress_egress_gw.inside_network](resources--gcp_vpc_site--reference--group-002.md#canonical-2123301101220113-0123332320222111-0202203303002102-2012313302013323-3011301022113123-0301003102033320-2332330212323112-0113111321202221)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-0131000100012312-2121310130001102-2202230331103203-2230222003223001-1113101010113111-2131313030021223-2030122333203101-2213001121001122)

<a id="canonical-0220321000013112-2030300310000323-2202131321000302-0112111112120211-0223122220232223-0023132200212001-3230213111203011-1330123303013023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3102231303321223-1000120233320000-0302301131331212-2033100121203332-2202233021322221-0012121333020220-2223023201001331-0022021233331223"></a>

## ingress_egress_gw.inside_static_routes — inside_static_routes / 103232320312 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-0131000100012312-2121310130001102-2202230331103203-2230222003223001-1113101010113111-2131313030021223-2030122333203101-2213001121001122)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-2101111213103332-1031003102012122-3033333002321300-0212010323021100-0301331300303302-3321221003203222-1230133302203013-3232231332123212)
- [ingress_egress_gw](resources--gcp_vpc_site--reference--group-001.md#canonical-3123100130121022-2313112013232222-0213021230302120-3230300230213031-3101322300112331-0011013220032322-1311222312201210-2213020302010100)
- ingress_egress_gw.inside_static_routes

<a id="canonical-0133112210023230-3021200312011212-3313310011222011-2013120121231213-0103202113122323-3332103012111213-2313220033121303-2011020002310033"></a>

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

<a id="canonical-2030321202133103-2202030331002023-2001030133001120-0030021012221212-0331030233102312-1020231023013322-0113121213210013-3302130013202113"></a>

## Direct properties — inside_static_routes / 103232320312 / 3

- [static_route_list](resources--gcp_vpc_site--reference--group-002.md#canonical-3112002300031122-1033032203133130-3320033210101302-1013211120120100-3021001213301031-1131300302030232-1123312000113231-0331101202321111): complete subsection reference.

<a id="canonical-0212031002210311-1312312322202010-2300203023322121-2231101121111203-2030313111012313-2222032023001311-2332001321302201-3303023333021013"></a>

## Next pages — inside_static_routes / 103232320312 / 4

- [ingress_egress_gw.inside_static_routes.static_route_list](resources--gcp_vpc_site--reference--group-002.md#canonical-3112002300031122-1033032203133130-3320033210101302-1013211120120100-3021001213301031-1131300302030232-1123312000113231-0331101202321111)
- [ingress_egress_gw](resources--gcp_vpc_site--reference--group-001.md#canonical-3123100130121022-2313112013232222-0213021230302120-3230300230213031-3101322300112331-0011013220032322-1311222312201210-2213020302010100)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-0131000100012312-2121310130001102-2202230331103203-2230222003223001-1113101010113111-2131313030021223-2030122333203101-2213001121001122)

<a id="canonical-3112002300031122-1033032203133130-3320033210101302-1013211120120100-3021001213301031-1131300302030232-1123312000113231-0331101202321111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1001113003200330-1002013132331103-1103221113100022-0110323110101132-0020301033112131-0001022332323333-0202232201102331-1002112203000033"></a>

## ingress_egress_gw.inside_static_routes.static_route_list — static_route_list / 122033032002 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-0131000100012312-2121310130001102-2202230331103203-2230222003223001-1113101010113111-2131313030021223-2030122333203101-2213001121001122)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-2101111213103332-1031003102012122-3033333002321300-0212010323021100-0301331300303302-3321221003203222-1230133302203013-3232231332123212)
- [ingress_egress_gw](resources--gcp_vpc_site--reference--group-001.md#canonical-3123100130121022-2313112013232222-0213021230302120-3230300230213031-3101322300112331-0011013220032322-1311222312201210-2213020302010100)
- [ingress_egress_gw.inside_static_routes](resources--gcp_vpc_site--reference--group-002.md#canonical-0220321000013112-2030300310000323-2202131321000302-0112111112120211-0223122220232223-0023132200212001-3230213111203011-1330123303013023)
- ingress_egress_gw.inside_static_routes.static_route_list

<a id="canonical-3301310210223300-3213012303212002-3012220032223330-2123112133313103-0012000330012030-2101212303221100-0313131113012321-3020020001301102"></a>

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

<a id="canonical-1231003032130033-0132012301231312-2302203231222220-1023011102112302-3230331221233332-1113332021100033-3230331033331012-2321233222102131"></a>

## Direct properties — static_route_list / 122033032002 / 3

- [custom_static_route](resources--gcp_vpc_site--reference--group-002.md#canonical-3333011102112323-2110222010120203-0032223033000231-0112221131033331-2231222113200233-2301232032001332-0331231203133210-0030103320023020): complete subsection reference.

<a id="canonical-2012312322322001-0302332013100301-3233111022333101-0313032011313103-2012301130300033-0013122132213333-0122321123212023-2130320202002310"></a>

<a id="canonical-0123033121211101-1020113210301223-0010131022020302-2031111102003112-2313132323003310-0302132302331033-0031200210323021-3302001301101123"></a>

## simple_static_route property — static_route_list / 122033032002 / 4

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

<a id="canonical-2023121221313232-2321002113200112-3203332210331201-1302210033300120-2213032223100311-0222031010310332-2200022033133013-3303022313213031"></a>

## Next pages — static_route_list / 122033032002 / 5

- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route](resources--gcp_vpc_site--reference--group-002.md#canonical-3333011102112323-2110222010120203-0032223033000231-0112221131033331-2231222113200233-2301232032001332-0331231203133210-0030103320023020)
- [ingress_egress_gw.inside_static_routes](resources--gcp_vpc_site--reference--group-002.md#canonical-0220321000013112-2030300310000323-2202131321000302-0112111112120211-0223122220232223-0023132200212001-3230213111203011-1330123303013023)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-0131000100012312-2121310130001102-2202230331103203-2230222003223001-1113101010113111-2131313030021223-2030122333203101-2213001121001122)

<a id="canonical-3333011102112323-2110222010120203-0032223033000231-0112221131033331-2231222113200233-2301232032001332-0331231203133210-0030103320023020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1332222012230022-3000013112022320-1133031233121111-3312311211131101-2031320122133203-3320220322330113-3110231000032123-3300022213120220"></a>

## ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route — custom_static_route / 222103122023 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-0131000100012312-2121310130001102-2202230331103203-2230222003223001-1113101010113111-2131313030021223-2030122333203101-2213001121001122)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-2101111213103332-1031003102012122-3033333002321300-0212010323021100-0301331300303302-3321221003203222-1230133302203013-3232231332123212)
- [ingress_egress_gw](resources--gcp_vpc_site--reference--group-001.md#canonical-3123100130121022-2313112013232222-0213021230302120-3230300230213031-3101322300112331-0011013220032322-1311222312201210-2213020302010100)
- [ingress_egress_gw.inside_static_routes](resources--gcp_vpc_site--reference--group-002.md#canonical-0220321000013112-2030300310000323-2202131321000302-0112111112120211-0223122220232223-0023132200212001-3230213111203011-1330123303013023)
- [ingress_egress_gw.inside_static_routes.static_route_list](resources--gcp_vpc_site--reference--group-002.md#canonical-3112002300031122-1033032203133130-3320033210101302-1013211120120100-3021001213301031-1131300302030232-1123312000113231-0331101202321111)
- ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route

<a id="canonical-3200112120011022-1100223203333232-1301030102233010-0000332021011002-0321010111111302-1312030333132021-3322320000222103-3003002032000213"></a>

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

<a id="canonical-3120000332120301-1202021102111330-2302023102033110-1000010023300031-1201123002113222-2303132233213112-1130313021010031-1012322012323230"></a>

## Direct properties — custom_static_route / 222103122023 / 3

<a id="canonical-0001221013223201-0130023310210122-2210033111201221-1102302231311323-1123122102010233-3011132320310200-1321221231020012-2020022302112333"></a>

<a id="canonical-3000130322221331-2300103103200200-3112333303320112-2230021201112213-0332032221101200-1213011213102123-2231133001013301-0031233100003130"></a>

## attrs property — custom_static_route / 222103122023 / 4

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

- [labels](resources--gcp_vpc_site--reference--group-002.md#canonical-1211333312230230-2213210111312113-2120232002030200-1321130021300122-2033231300231332-2321002003002310-0101301123130022-1202003031003112): complete subsection reference.

- [nexthop](resources--gcp_vpc_site--reference--group-002.md#canonical-3300033221322203-3221030020310321-0300023333210301-3213020131322111-2123320301321122-0313303011200303-3120311221112230-3103213011123030): complete subsection reference.

- [subnets](resources--gcp_vpc_site--reference--group-002.md#canonical-0233220023132033-0120221223320323-0132122022200220-0021223222210311-3020301231112332-3111211310000102-2300300121313311-0203211200320112): complete subsection reference.

<a id="canonical-3023220113210300-3200230011211121-1100123322223200-1103133220012033-2121132113111010-3133033113022203-3333032113303012-2131031300232012"></a>

## Next pages — custom_static_route / 222103122023 / 5

- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.labels](resources--gcp_vpc_site--reference--group-002.md#canonical-1211333312230230-2213210111312113-2120232002030200-1321130021300122-2033231300231332-2321002003002310-0101301123130022-1202003031003112)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop](resources--gcp_vpc_site--reference--group-002.md#canonical-3300033221322203-3221030020310321-0300023333210301-3213020131322111-2123320301321122-0313303011200303-3120311221112230-3103213011123030)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets](resources--gcp_vpc_site--reference--group-002.md#canonical-0233220023132033-0120221223320323-0132122022200220-0021223222210311-3020301231112332-3111211310000102-2300300121313311-0203211200320112)
- [ingress_egress_gw.inside_static_routes.static_route_list](resources--gcp_vpc_site--reference--group-002.md#canonical-3112002300031122-1033032203133130-3320033210101302-1013211120120100-3021001213301031-1131300302030232-1123312000113231-0331101202321111)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-0131000100012312-2121310130001102-2202230331103203-2230222003223001-1113101010113111-2131313030021223-2030122333203101-2213001121001122)

<a id="canonical-1211333312230230-2213210111312113-2120232002030200-1321130021300122-2033231300231332-2321002003002310-0101301123130022-1202003031003112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0313000120302013-1331333213130033-0200101321020011-0310232003021321-3321310132220122-1103310310000101-2232012301210321-0022100203322211"></a>

## ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.labels — labels / 103100022111 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-0131000100012312-2121310130001102-2202230331103203-2230222003223001-1113101010113111-2131313030021223-2030122333203101-2213001121001122)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-2101111213103332-1031003102012122-3033333002321300-0212010323021100-0301331300303302-3321221003203222-1230133302203013-3232231332123212)
- [ingress_egress_gw](resources--gcp_vpc_site--reference--group-001.md#canonical-3123100130121022-2313112013232222-0213021230302120-3230300230213031-3101322300112331-0011013220032322-1311222312201210-2213020302010100)
- [ingress_egress_gw.inside_static_routes](resources--gcp_vpc_site--reference--group-002.md#canonical-0220321000013112-2030300310000323-2202131321000302-0112111112120211-0223122220232223-0023132200212001-3230213111203011-1330123303013023)
- [ingress_egress_gw.inside_static_routes.static_route_list](resources--gcp_vpc_site--reference--group-002.md#canonical-3112002300031122-1033032203133130-3320033210101302-1013211120120100-3021001213301031-1131300302030232-1123312000113231-0331101202321111)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route](resources--gcp_vpc_site--reference--group-002.md#canonical-3333011102112323-2110222010120203-0032223033000231-0112221131033331-2231222113200233-2301232032001332-0331231203133210-0030103320023020)
- ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.labels

<a id="canonical-3220211310333232-0011032101322303-3132022210100201-3110003020301013-1112031001220211-1023130223111302-2111020322202301-0030022231222133"></a>

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

<a id="canonical-0303200012212102-3033120300311112-0311100101221110-1022202202300233-1302001023000131-3000321220230202-0112102200130132-0022113320130211"></a>

## Direct properties — labels / 103100022111 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2233122212330010-3202223000133001-0121021011131113-2202333013032000-0113102101322133-3303302200333103-0321001310223333-3003201100111331"></a>

## Next pages — labels / 103100022111 / 4

- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route](resources--gcp_vpc_site--reference--group-002.md#canonical-3333011102112323-2110222010120203-0032223033000231-0112221131033331-2231222113200233-2301232032001332-0331231203133210-0030103320023020)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-0131000100012312-2121310130001102-2202230331103203-2230222003223001-1113101010113111-2131313030021223-2030122333203101-2213001121001122)

<a id="canonical-3300033221322203-3221030020310321-0300023333210301-3213020131322111-2123320301321122-0313303011200303-3120311221112230-3103213011123030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0112012011210023-0122123313202110-2031223200332032-1132022110323331-0120300213101322-0203301311320333-2220320210101010-2233221310133220"></a>

## ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop — nexthop / 022113202303 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-0131000100012312-2121310130001102-2202230331103203-2230222003223001-1113101010113111-2131313030021223-2030122333203101-2213001121001122)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-2101111213103332-1031003102012122-3033333002321300-0212010323021100-0301331300303302-3321221003203222-1230133302203013-3232231332123212)
- [ingress_egress_gw](resources--gcp_vpc_site--reference--group-001.md#canonical-3123100130121022-2313112013232222-0213021230302120-3230300230213031-3101322300112331-0011013220032322-1311222312201210-2213020302010100)
- [ingress_egress_gw.inside_static_routes](resources--gcp_vpc_site--reference--group-002.md#canonical-0220321000013112-2030300310000323-2202131321000302-0112111112120211-0223122220232223-0023132200212001-3230213111203011-1330123303013023)
- [ingress_egress_gw.inside_static_routes.static_route_list](resources--gcp_vpc_site--reference--group-002.md#canonical-3112002300031122-1033032203133130-3320033210101302-1013211120120100-3021001213301031-1131300302030232-1123312000113231-0331101202321111)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route](resources--gcp_vpc_site--reference--group-002.md#canonical-3333011102112323-2110222010120203-0032223033000231-0112221131033331-2231222113200233-2301232032001332-0331231203133210-0030103320023020)
- ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop

<a id="canonical-3332131301221222-3012201130130202-1133103213313201-0312321303200123-0131220230111001-0120312001101322-2313102011330023-3111313313312031"></a>

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

<a id="canonical-2212013131101022-1101210313112121-2033201330021220-1200313113102200-0302223011103203-0310312010211331-0101300232132300-3101331120122210"></a>

## Direct properties — nexthop / 022113202303 / 3

- [interface](resources--gcp_vpc_site--reference--group-002.md#canonical-1110211211332022-1321232003031103-0202002102022111-0110330213211122-3312320202211110-2032310102203123-3120213101230120-1131010331231021): complete subsection reference.

- [nexthop_address](resources--gcp_vpc_site--reference--group-002.md#canonical-0220310313220112-0023330303232131-2113103022003023-0220221322022133-1102332132322032-2023012010311002-0111003322022130-0232033233001031): complete subsection reference.

<a id="canonical-3022302103133000-0000101130203310-0013303120123331-2003203130330303-0313101313333203-1210001102303030-0323222202010231-1013132210132020"></a>

<a id="canonical-3020022210301230-1202213210110030-3332133303211302-2111131312220300-3023230130022120-1331323302310123-3313012120131012-3231010010033223"></a>

## type property — nexthop / 022113202303 / 4

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

<a id="canonical-3003230303302003-1222002310031302-3123203101233111-3111230320123203-2110323133300030-0123120031322120-1323231030321121-1121313230233023"></a>

## Next pages — nexthop / 022113202303 / 5

- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.interface](resources--gcp_vpc_site--reference--group-002.md#canonical-1110211211332022-1321232003031103-0202002102022111-0110330213211122-3312320202211110-2032310102203123-3120213101230120-1131010331231021)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](resources--gcp_vpc_site--reference--group-002.md#canonical-0220310313220112-0023330303232131-2113103022003023-0220221322022133-1102332132322032-2023012010311002-0111003322022130-0232033233001031)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route](resources--gcp_vpc_site--reference--group-002.md#canonical-3333011102112323-2110222010120203-0032223033000231-0112221131033331-2231222113200233-2301232032001332-0331231203133210-0030103320023020)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-0131000100012312-2121310130001102-2202230331103203-2230222003223001-1113101010113111-2131313030021223-2030122333203101-2213001121001122)

<a id="canonical-1110211211332022-1321232003031103-0202002102022111-0110330213211122-3312320202211110-2032310102203123-3120213101230120-1131010331231021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0112301021333002-2323022310102332-2013212131003000-1303012012303032-0122031011031120-2120201013311233-0120301223130232-2310313210133133"></a>

## ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.interface — interface / 303213233221 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-0131000100012312-2121310130001102-2202230331103203-2230222003223001-1113101010113111-2131313030021223-2030122333203101-2213001121001122)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-2101111213103332-1031003102012122-3033333002321300-0212010323021100-0301331300303302-3321221003203222-1230133302203013-3232231332123212)
- [ingress_egress_gw](resources--gcp_vpc_site--reference--group-001.md#canonical-3123100130121022-2313112013232222-0213021230302120-3230300230213031-3101322300112331-0011013220032322-1311222312201210-2213020302010100)
- [ingress_egress_gw.inside_static_routes](resources--gcp_vpc_site--reference--group-002.md#canonical-0220321000013112-2030300310000323-2202131321000302-0112111112120211-0223122220232223-0023132200212001-3230213111203011-1330123303013023)
- [ingress_egress_gw.inside_static_routes.static_route_list](resources--gcp_vpc_site--reference--group-002.md#canonical-3112002300031122-1033032203133130-3320033210101302-1013211120120100-3021001213301031-1131300302030232-1123312000113231-0331101202321111)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route](resources--gcp_vpc_site--reference--group-002.md#canonical-3333011102112323-2110222010120203-0032223033000231-0112221131033331-2231222113200233-2301232032001332-0331231203133210-0030103320023020)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop](resources--gcp_vpc_site--reference--group-002.md#canonical-3300033221322203-3221030020310321-0300023333210301-3213020131322111-2123320301321122-0313303011200303-3120311221112230-3103213011123030)
- ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.interface

<a id="canonical-0111101213002322-3133310002003222-0010033001100120-3123111231203210-2233030003030131-0331331131013122-3103301122123120-2023122233230121"></a>

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

<a id="canonical-0323121323210203-1332213020033323-3013131331023103-0223022000303313-3001202303100101-2132313300130113-1202313332111220-0312002233320000"></a>

## Direct properties — interface / 303213233221 / 3

<a id="canonical-3013210303222031-3120232331330302-3221123313132220-0332300033133000-2022333231211320-1220301202211311-2231310232112130-2323332121223103"></a>

<a id="canonical-3210110322012032-1113303130121220-3011302003210232-3003213330110113-2121321330231120-2301112103023102-1100022333301200-3323121202231333"></a>

## kind property — interface / 303213233221 / 4

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

<a id="canonical-0331312001212012-2310003003123021-2333113023103203-2132012132101201-2022101321122021-2200210003313033-0220122222222213-1210110123130322"></a>

<a id="canonical-3330013120133010-0222211030033232-0301020032301332-0110122120220110-1112301012122011-3032032331211221-0212313320231010-3121201020233331"></a>

## name property — interface / 303213233221 / 5

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

<a id="canonical-1222333002213221-3000202203331202-2322330222011323-3213233001221333-0231000101120320-3222200211333021-1212211020110032-1132221030121112"></a>

<a id="canonical-0301120322202201-1222313031110232-1330132301112333-3213331121313123-1301213130033113-3112331333011210-2221032332303203-3230221221000332"></a>

## namespace property — interface / 303213233221 / 6

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

<a id="canonical-2221030031221101-3010230313310201-0303300201332213-3223012002233331-0133031212010201-2033111201100103-0231000300322201-3310110020101111"></a>

<a id="canonical-0120221112002112-0313321211133112-1120220113021223-2301322131233131-1100321213011021-2122332313100102-2321320323203211-3103323212023003"></a>

## tenant property — interface / 303213233221 / 7

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

<a id="canonical-1230200212302333-3100323231300321-3032010200133010-3022013132210201-3311131233130023-1232103313301222-1012222210210113-0231002130123120"></a>

<a id="canonical-3030112302110222-2212210001003211-1310023030021022-0200122310131302-2222303303020312-1330121100231313-1111301201132021-0010021021212321"></a>

## uid property — interface / 303213233221 / 8

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

<a id="canonical-0321122301033130-2311000231131332-3200132202331233-2221232323132123-2202130202310010-0022033012313210-2230300312201112-2002200121323013"></a>

## Next pages — interface / 303213233221 / 9

- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop](resources--gcp_vpc_site--reference--group-002.md#canonical-3300033221322203-3221030020310321-0300023333210301-3213020131322111-2123320301321122-0313303011200303-3120311221112230-3103213011123030)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-0131000100012312-2121310130001102-2202230331103203-2230222003223001-1113101010113111-2131313030021223-2030122333203101-2213001121001122)

<a id="canonical-0220310313220112-0023330303232131-2113103022003023-0220221322022133-1102332132322032-2023012010311002-0111003322022130-0232033233001031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0020233132122113-2111332302010213-3102222321331003-0030300020331230-2103300112313221-2123100101022101-0232200322020120-0031100122313122"></a>

## ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address — nexthop_address / 211201201332 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-0131000100012312-2121310130001102-2202230331103203-2230222003223001-1113101010113111-2131313030021223-2030122333203101-2213001121001122)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-2101111213103332-1031003102012122-3033333002321300-0212010323021100-0301331300303302-3321221003203222-1230133302203013-3232231332123212)
- [ingress_egress_gw](resources--gcp_vpc_site--reference--group-001.md#canonical-3123100130121022-2313112013232222-0213021230302120-3230300230213031-3101322300112331-0011013220032322-1311222312201210-2213020302010100)
- [ingress_egress_gw.inside_static_routes](resources--gcp_vpc_site--reference--group-002.md#canonical-0220321000013112-2030300310000323-2202131321000302-0112111112120211-0223122220232223-0023132200212001-3230213111203011-1330123303013023)
- [ingress_egress_gw.inside_static_routes.static_route_list](resources--gcp_vpc_site--reference--group-002.md#canonical-3112002300031122-1033032203133130-3320033210101302-1013211120120100-3021001213301031-1131300302030232-1123312000113231-0331101202321111)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route](resources--gcp_vpc_site--reference--group-002.md#canonical-3333011102112323-2110222010120203-0032223033000231-0112221131033331-2231222113200233-2301232032001332-0331231203133210-0030103320023020)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop](resources--gcp_vpc_site--reference--group-002.md#canonical-3300033221322203-3221030020310321-0300023333210301-3213020131322111-2123320301321122-0313303011200303-3120311221112230-3103213011123030)
- ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address

<a id="canonical-2113123313021012-0022131113003100-1213131122123210-1302330231303231-1132020013210232-2000022222130322-0202122030221023-3203332120013221"></a>

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

<a id="canonical-1032123132131101-2310333313300002-1011203003001302-3130301103310113-2130301310113033-0233221121331100-1032210321113102-3000032102023323"></a>

## Direct properties — nexthop_address / 211201201332 / 3

- [dual_stack](resources--gcp_vpc_site--reference--group-002.md#canonical-1220333300233211-1300220223202303-2022300303031100-3331112022301020-1110130013231100-2232132020121122-0221101131030113-2213100031012120): complete subsection reference.

- [IPv4](resources--gcp_vpc_site--reference--group-002.md#canonical-3020020211002203-2130030000300210-1001213200123031-1123033133300031-3010030001213211-1300021202212001-0021313231021320-2302321210322022): complete subsection reference.

- [IPv6](resources--gcp_vpc_site--reference--group-002.md#canonical-2300321022112010-2020312230101012-3031112123323133-2331111212333221-0002300331222011-0001123221223232-0312020210211103-1333121000330312): complete subsection reference.

<a id="canonical-3010101023111200-1301203133212010-1003032212230133-0002110310221012-1303303133203211-3131133212033233-3033222131100331-3130332112010002"></a>

## Next pages — nexthop_address / 211201201332 / 4

- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](resources--gcp_vpc_site--reference--group-002.md#canonical-1220333300233211-1300220223202303-2022300303031100-3331112022301020-1110130013231100-2232132020121122-0221101131030113-2213100031012120)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4](resources--gcp_vpc_site--reference--group-002.md#canonical-3020020211002203-2130030000300210-1001213200123031-1123033133300031-3010030001213211-1300021202212001-0021313231021320-2302321210322022)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6](resources--gcp_vpc_site--reference--group-002.md#canonical-2300321022112010-2020312230101012-3031112123323133-2331111212333221-0002300331222011-0001123221223232-0312020210211103-1333121000330312)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop](resources--gcp_vpc_site--reference--group-002.md#canonical-3300033221322203-3221030020310321-0300023333210301-3213020131322111-2123320301321122-0313303011200303-3120311221112230-3103213011123030)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-0131000100012312-2121310130001102-2202230331103203-2230222003223001-1113101010113111-2131313030021223-2030122333203101-2213001121001122)

<a id="canonical-1220333300233211-1300220223202303-2022300303031100-3331112022301020-1110130013231100-2232132020121122-0221101131030113-2213100031012120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2202200323202130-1131223312322320-1023023113131311-1130302213202302-2010111301232001-1302220313210200-0031033033132112-3213130033313300"></a>

## ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack — dual_stack / 310001131201 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-0131000100012312-2121310130001102-2202230331103203-2230222003223001-1113101010113111-2131313030021223-2030122333203101-2213001121001122)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-2101111213103332-1031003102012122-3033333002321300-0212010323021100-0301331300303302-3321221003203222-1230133302203013-3232231332123212)
- [ingress_egress_gw](resources--gcp_vpc_site--reference--group-001.md#canonical-3123100130121022-2313112013232222-0213021230302120-3230300230213031-3101322300112331-0011013220032322-1311222312201210-2213020302010100)
- [ingress_egress_gw.inside_static_routes](resources--gcp_vpc_site--reference--group-002.md#canonical-0220321000013112-2030300310000323-2202131321000302-0112111112120211-0223122220232223-0023132200212001-3230213111203011-1330123303013023)
- [ingress_egress_gw.inside_static_routes.static_route_list](resources--gcp_vpc_site--reference--group-002.md#canonical-3112002300031122-1033032203133130-3320033210101302-1013211120120100-3021001213301031-1131300302030232-1123312000113231-0331101202321111)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route](resources--gcp_vpc_site--reference--group-002.md#canonical-3333011102112323-2110222010120203-0032223033000231-0112221131033331-2231222113200233-2301232032001332-0331231203133210-0030103320023020)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop](resources--gcp_vpc_site--reference--group-002.md#canonical-3300033221322203-3221030020310321-0300023333210301-3213020131322111-2123320301321122-0313303011200303-3120311221112230-3103213011123030)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](resources--gcp_vpc_site--reference--group-002.md#canonical-0220310313220112-0023330303232131-2113103022003023-0220221322022133-1102332132322032-2023012010311002-0111003322022130-0232033233001031)
- ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack

<a id="canonical-3020013230200310-3330320213132302-0321100103222221-0000202103110232-1220023310023310-2021233002321310-3332003310101301-3303212101131132"></a>

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

<a id="canonical-2311331333303213-2122300210013003-0020133111001020-1101203022311201-1322211210121212-2012113213133313-3021313332013121-0233303003133031"></a>

## Direct properties — dual_stack / 310001131201 / 3

- [IPv4](resources--gcp_vpc_site--reference--group-002.md#canonical-1133232200000131-3022331213023110-2123300223331211-2212303030331123-2122221202332100-3102310000331320-1210221311322302-0300312013020032): complete subsection reference.

- [IPv6](resources--gcp_vpc_site--reference--group-002.md#canonical-1220313330010131-0122330321101101-1113322113132032-1120110020112120-2130032130030012-0231131220012130-1130212012312232-0312022002231111): complete subsection reference.

<a id="canonical-1002130030330313-0113233130300111-1300100311102111-1121110032022103-1111131322031030-0221312132132132-3313120313330332-0321130211133332"></a>

## Next pages — dual_stack / 310001131201 / 4

- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4](resources--gcp_vpc_site--reference--group-002.md#canonical-1133232200000131-3022331213023110-2123300223331211-2212303030331123-2122221202332100-3102310000331320-1210221311322302-0300312013020032)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6](resources--gcp_vpc_site--reference--group-002.md#canonical-1220313330010131-0122330321101101-1113322113132032-1120110020112120-2130032130030012-0231131220012130-1130212012312232-0312022002231111)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](resources--gcp_vpc_site--reference--group-002.md#canonical-0220310313220112-0023330303232131-2113103022003023-0220221322022133-1102332132322032-2023012010311002-0111003322022130-0232033233001031)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-0131000100012312-2121310130001102-2202230331103203-2230222003223001-1113101010113111-2131313030021223-2030122333203101-2213001121001122)

<a id="canonical-1133232200000131-3022331213023110-2123300223331211-2212303030331123-2122221202332100-3102310000331320-1210221311322302-0300312013020032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0212112333301210-3211120200031313-2331002122301003-3100102231212230-2123112131310122-1202231331201300-3313013031200210-1102020323233201"></a>

## ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.IPv4 — IPv4 / 010232011323 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-0131000100012312-2121310130001102-2202230331103203-2230222003223001-1113101010113111-2131313030021223-2030122333203101-2213001121001122)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-2101111213103332-1031003102012122-3033333002321300-0212010323021100-0301331300303302-3321221003203222-1230133302203013-3232231332123212)
- [ingress_egress_gw](resources--gcp_vpc_site--reference--group-001.md#canonical-3123100130121022-2313112013232222-0213021230302120-3230300230213031-3101322300112331-0011013220032322-1311222312201210-2213020302010100)
- [ingress_egress_gw.inside_static_routes](resources--gcp_vpc_site--reference--group-002.md#canonical-0220321000013112-2030300310000323-2202131321000302-0112111112120211-0223122220232223-0023132200212001-3230213111203011-1330123303013023)
- [ingress_egress_gw.inside_static_routes.static_route_list](resources--gcp_vpc_site--reference--group-002.md#canonical-3112002300031122-1033032203133130-3320033210101302-1013211120120100-3021001213301031-1131300302030232-1123312000113231-0331101202321111)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route](resources--gcp_vpc_site--reference--group-002.md#canonical-3333011102112323-2110222010120203-0032223033000231-0112221131033331-2231222113200233-2301232032001332-0331231203133210-0030103320023020)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop](resources--gcp_vpc_site--reference--group-002.md#canonical-3300033221322203-3221030020310321-0300023333210301-3213020131322111-2123320301321122-0313303011200303-3120311221112230-3103213011123030)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](resources--gcp_vpc_site--reference--group-002.md#canonical-0220310313220112-0023330303232131-2113103022003023-0220221322022133-1102332132322032-2023012010311002-0111003322022130-0232033233001031)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](resources--gcp_vpc_site--reference--group-002.md#canonical-1220333300233211-1300220223202303-2022300303031100-3331112022301020-1110130013231100-2232132020121122-0221101131030113-2213100031012120)
- ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.IPv4

<a id="canonical-1223001332001211-0011011220211030-3000032011232001-2101322031211201-1330233031133312-0201201132031111-0000133301212133-1232113003301232"></a>

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

<a id="canonical-2003132322132332-1030013002321300-3100120210130100-3300223030202212-2322132101201123-3123102222112300-1220030100233201-3123002101123230"></a>

## Direct properties — IPv4 / 010232011323 / 3

<a id="canonical-0111233323303322-3130131323001210-3233203031013220-2021033202201031-0023313130300101-0203021132100113-1102122133223130-3032102221231012"></a>

<a id="canonical-0302011222131001-2031210133002113-1331033020123133-2010101212003221-3032323233212203-1231221333013203-0123313312000202-0101001033102122"></a>

## addr property — IPv4 / 010232011323 / 4

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

<a id="canonical-2120331112011001-1212133302300001-2300020013000321-3212022222322033-0130321232012001-1201230101033222-3312132312012332-2211103001131121"></a>

## Next pages — IPv4 / 010232011323 / 5

- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](resources--gcp_vpc_site--reference--group-002.md#canonical-1220333300233211-1300220223202303-2022300303031100-3331112022301020-1110130013231100-2232132020121122-0221101131030113-2213100031012120)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-0131000100012312-2121310130001102-2202230331103203-2230222003223001-1113101010113111-2131313030021223-2030122333203101-2213001121001122)

<a id="canonical-1220313330010131-0122330321101101-1113322113132032-1120110020112120-2130032130030012-0231131220012130-1130212012312232-0312022002231111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2322022112021211-3030320200231022-1223101201203223-3220023301010300-2033120120303233-3330003221113230-0012330122332001-3002302213113110"></a>

## ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.IPv6 — IPv6 / 033211100001 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-0131000100012312-2121310130001102-2202230331103203-2230222003223001-1113101010113111-2131313030021223-2030122333203101-2213001121001122)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-2101111213103332-1031003102012122-3033333002321300-0212010323021100-0301331300303302-3321221003203222-1230133302203013-3232231332123212)
- [ingress_egress_gw](resources--gcp_vpc_site--reference--group-001.md#canonical-3123100130121022-2313112013232222-0213021230302120-3230300230213031-3101322300112331-0011013220032322-1311222312201210-2213020302010100)
- [ingress_egress_gw.inside_static_routes](resources--gcp_vpc_site--reference--group-002.md#canonical-0220321000013112-2030300310000323-2202131321000302-0112111112120211-0223122220232223-0023132200212001-3230213111203011-1330123303013023)
- [ingress_egress_gw.inside_static_routes.static_route_list](resources--gcp_vpc_site--reference--group-002.md#canonical-3112002300031122-1033032203133130-3320033210101302-1013211120120100-3021001213301031-1131300302030232-1123312000113231-0331101202321111)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route](resources--gcp_vpc_site--reference--group-002.md#canonical-3333011102112323-2110222010120203-0032223033000231-0112221131033331-2231222113200233-2301232032001332-0331231203133210-0030103320023020)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop](resources--gcp_vpc_site--reference--group-002.md#canonical-3300033221322203-3221030020310321-0300023333210301-3213020131322111-2123320301321122-0313303011200303-3120311221112230-3103213011123030)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](resources--gcp_vpc_site--reference--group-002.md#canonical-0220310313220112-0023330303232131-2113103022003023-0220221322022133-1102332132322032-2023012010311002-0111003322022130-0232033233001031)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](resources--gcp_vpc_site--reference--group-002.md#canonical-1220333300233211-1300220223202303-2022300303031100-3331112022301020-1110130013231100-2232132020121122-0221101131030113-2213100031012120)
- ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.IPv6

<a id="canonical-3001120313121113-0130311321323001-3231110210010020-3002203100132233-1221231133120332-1233111320312210-1102132332003301-2323322212330313"></a>

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

<a id="canonical-0332220023300220-1000221223020311-3212202231033122-3220102320122200-2202301021323231-0300311213302030-1013212112131131-0330323131320312"></a>

## Direct properties — IPv6 / 033211100001 / 3

<a id="canonical-1310202121130231-2311023200023233-1203303203013223-2300121233012022-1012202333012120-1120211113203132-2332123122333333-2312203203101321"></a>

<a id="canonical-3231103323112011-3323313013012023-2331231133300022-2003130232021113-3111013312101110-1331020110121201-0131112310203201-0033011313131311"></a>

## addr property — IPv6 / 033211100001 / 4

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

<a id="canonical-1211203330031212-2123213300131000-2210012131223000-1220002333230123-0113312233003223-2203111321223333-3111012223303110-0312213121232111"></a>

## Next pages — IPv6 / 033211100001 / 5

- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](resources--gcp_vpc_site--reference--group-002.md#canonical-1220333300233211-1300220223202303-2022300303031100-3331112022301020-1110130013231100-2232132020121122-0221101131030113-2213100031012120)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-0131000100012312-2121310130001102-2202230331103203-2230222003223001-1113101010113111-2131313030021223-2030122333203101-2213001121001122)

<a id="canonical-3020020211002203-2130030000300210-1001213200123031-1123033133300031-3010030001213211-1300021202212001-0021313231021320-2302321210322022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0213303333210033-2303031103030211-2130002312223201-1322132010311332-3310230213121312-1013122000200223-3023112110332320-3331302230001020"></a>

## ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.IPv4 — IPv4 / 013320233311 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-0131000100012312-2121310130001102-2202230331103203-2230222003223001-1113101010113111-2131313030021223-2030122333203101-2213001121001122)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-2101111213103332-1031003102012122-3033333002321300-0212010323021100-0301331300303302-3321221003203222-1230133302203013-3232231332123212)
- [ingress_egress_gw](resources--gcp_vpc_site--reference--group-001.md#canonical-3123100130121022-2313112013232222-0213021230302120-3230300230213031-3101322300112331-0011013220032322-1311222312201210-2213020302010100)
- [ingress_egress_gw.inside_static_routes](resources--gcp_vpc_site--reference--group-002.md#canonical-0220321000013112-2030300310000323-2202131321000302-0112111112120211-0223122220232223-0023132200212001-3230213111203011-1330123303013023)
- [ingress_egress_gw.inside_static_routes.static_route_list](resources--gcp_vpc_site--reference--group-002.md#canonical-3112002300031122-1033032203133130-3320033210101302-1013211120120100-3021001213301031-1131300302030232-1123312000113231-0331101202321111)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route](resources--gcp_vpc_site--reference--group-002.md#canonical-3333011102112323-2110222010120203-0032223033000231-0112221131033331-2231222113200233-2301232032001332-0331231203133210-0030103320023020)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop](resources--gcp_vpc_site--reference--group-002.md#canonical-3300033221322203-3221030020310321-0300023333210301-3213020131322111-2123320301321122-0313303011200303-3120311221112230-3103213011123030)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](resources--gcp_vpc_site--reference--group-002.md#canonical-0220310313220112-0023330303232131-2113103022003023-0220221322022133-1102332132322032-2023012010311002-0111003322022130-0232033233001031)
- ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.IPv4

<a id="canonical-1020202310020110-1231011321201001-2012022111303213-2031120210120033-3102212022003122-1000133123210300-0321103012220200-1323011102022013"></a>

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

<a id="canonical-1023013321230220-3021231131102302-1100102000312201-3223333232133130-3330321202000330-0230031022023100-1001300012133020-2233012103321222"></a>

## Direct properties — IPv4 / 013320233311 / 3

<a id="canonical-0112030332032011-1212020311011303-0213331221220201-3013110303120322-3230310103101110-0103133102011301-2110333303010313-1002100311310321"></a>

<a id="canonical-1322102031010030-2001303033232121-3322300233121112-2330320333222131-2102122011012201-2303233310221131-1110001033232331-3110020321300112"></a>

## addr property — IPv4 / 013320233311 / 4

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

<a id="canonical-3033133301122002-1220323210030023-3321033123331201-2232321012201133-2222120130010202-1222310131322101-0130112122301233-3103231202331102"></a>

## Next pages — IPv4 / 013320233311 / 5

- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](resources--gcp_vpc_site--reference--group-002.md#canonical-0220310313220112-0023330303232131-2113103022003023-0220221322022133-1102332132322032-2023012010311002-0111003322022130-0232033233001031)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-0131000100012312-2121310130001102-2202230331103203-2230222003223001-1113101010113111-2131313030021223-2030122333203101-2213001121001122)

<a id="canonical-2300321022112010-2020312230101012-3031112123323133-2331111212333221-0002300331222011-0001123221223232-0312020210211103-1333121000330312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0033210122200323-3330311002232331-2210121203322223-1003020201002222-0133000221121132-1020232032331201-0100123020320322-2100201101112311"></a>

## ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.IPv6 — IPv6 / 100203120012 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-0131000100012312-2121310130001102-2202230331103203-2230222003223001-1113101010113111-2131313030021223-2030122333203101-2213001121001122)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-2101111213103332-1031003102012122-3033333002321300-0212010323021100-0301331300303302-3321221003203222-1230133302203013-3232231332123212)
- [ingress_egress_gw](resources--gcp_vpc_site--reference--group-001.md#canonical-3123100130121022-2313112013232222-0213021230302120-3230300230213031-3101322300112331-0011013220032322-1311222312201210-2213020302010100)
- [ingress_egress_gw.inside_static_routes](resources--gcp_vpc_site--reference--group-002.md#canonical-0220321000013112-2030300310000323-2202131321000302-0112111112120211-0223122220232223-0023132200212001-3230213111203011-1330123303013023)
- [ingress_egress_gw.inside_static_routes.static_route_list](resources--gcp_vpc_site--reference--group-002.md#canonical-3112002300031122-1033032203133130-3320033210101302-1013211120120100-3021001213301031-1131300302030232-1123312000113231-0331101202321111)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route](resources--gcp_vpc_site--reference--group-002.md#canonical-3333011102112323-2110222010120203-0032223033000231-0112221131033331-2231222113200233-2301232032001332-0331231203133210-0030103320023020)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop](resources--gcp_vpc_site--reference--group-002.md#canonical-3300033221322203-3221030020310321-0300023333210301-3213020131322111-2123320301321122-0313303011200303-3120311221112230-3103213011123030)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](resources--gcp_vpc_site--reference--group-002.md#canonical-0220310313220112-0023330303232131-2113103022003023-0220221322022133-1102332132322032-2023012010311002-0111003322022130-0232033233001031)
- ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.IPv6

<a id="canonical-1032122123033232-1020000101211023-1222030100201232-1233312002330133-2031230001230232-1010120020013220-3323031201211130-2020322130123021"></a>

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

<a id="canonical-0211222320202320-3221101223202001-2223001130202232-0200220331121133-1332222320000323-0112120110012023-3231310312122123-1133031221103000"></a>

## Direct properties — IPv6 / 100203120012 / 3

<a id="canonical-0212311323003120-2222121111100212-1012321321322111-2131123321321021-1322211012200031-2123033130221211-3000111111200020-1221233123130301"></a>

<a id="canonical-1122003120320332-0002301012000122-0112102230023210-1023333332330132-3222120300013020-0030012033132213-0013220310201031-1300311332021313"></a>

## addr property — IPv6 / 100203120012 / 4

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

<a id="canonical-0033133103332001-2011231313331332-2310320210112003-3323020221333332-1131122321113200-3022031102022002-2130223313330322-3111003003233000"></a>

## Next pages — IPv6 / 100203120012 / 5

- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](resources--gcp_vpc_site--reference--group-002.md#canonical-0220310313220112-0023330303232131-2113103022003023-0220221322022133-1102332132322032-2023012010311002-0111003322022130-0232033233001031)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-0131000100012312-2121310130001102-2202230331103203-2230222003223001-1113101010113111-2131313030021223-2030122333203101-2213001121001122)

<a id="canonical-0233220023132033-0120221223320323-0132122022200220-0021223222210311-3020301231112332-3111211310000102-2300300121313311-0203211200320112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1331130112230032-0213211231112222-2220123120320030-2002302200330011-0331021300022000-3312333212230002-2202123110213301-2213302332312023"></a>

## ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets — subnets / 333223131202 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-0131000100012312-2121310130001102-2202230331103203-2230222003223001-1113101010113111-2131313030021223-2030122333203101-2213001121001122)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-2101111213103332-1031003102012122-3033333002321300-0212010323021100-0301331300303302-3321221003203222-1230133302203013-3232231332123212)
- [ingress_egress_gw](resources--gcp_vpc_site--reference--group-001.md#canonical-3123100130121022-2313112013232222-0213021230302120-3230300230213031-3101322300112331-0011013220032322-1311222312201210-2213020302010100)
- [ingress_egress_gw.inside_static_routes](resources--gcp_vpc_site--reference--group-002.md#canonical-0220321000013112-2030300310000323-2202131321000302-0112111112120211-0223122220232223-0023132200212001-3230213111203011-1330123303013023)
- [ingress_egress_gw.inside_static_routes.static_route_list](resources--gcp_vpc_site--reference--group-002.md#canonical-3112002300031122-1033032203133130-3320033210101302-1013211120120100-3021001213301031-1131300302030232-1123312000113231-0331101202321111)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route](resources--gcp_vpc_site--reference--group-002.md#canonical-3333011102112323-2110222010120203-0032223033000231-0112221131033331-2231222113200233-2301232032001332-0331231203133210-0030103320023020)
- ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets

<a id="canonical-3311220320031023-2130210220123312-0032220220303332-2213033122201123-2123231030013210-2230202002212302-3130320032213301-3331231101000320"></a>

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

<a id="canonical-3032011213122312-2031320033301020-3233321021103012-0110111120022333-2131203323322311-1101110100300213-0122103231321320-0201012001131233"></a>

## Direct properties — subnets / 333223131202 / 3

- [IPv4](resources--gcp_vpc_site--reference--group-002.md#canonical-1122101220000020-0010221111301221-2313213102013113-1000203013113222-2203112310330010-2031201223130030-1222022103300121-1201003232001110): complete subsection reference.

- [IPv6](resources--gcp_vpc_site--reference--group-002.md#canonical-0321213203212110-1123200000030301-2223303011330002-2333211310223331-0123011213131230-1010032313301123-2311010211301302-3220220131003033): complete subsection reference.

<a id="canonical-3200223011121303-2211213222022210-1132001020320233-1231230313022233-1130023331211122-2020100022010303-2031110231122000-2303200121230111"></a>

## Next pages — subnets / 333223131202 / 4

- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets.ipv4](resources--gcp_vpc_site--reference--group-002.md#canonical-1122101220000020-0010221111301221-2313213102013113-1000203013113222-2203112310330010-2031201223130030-1222022103300121-1201003232001110)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets.ipv6](resources--gcp_vpc_site--reference--group-002.md#canonical-0321213203212110-1123200000030301-2223303011330002-2333211310223331-0123011213131230-1010032313301123-2311010211301302-3220220131003033)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route](resources--gcp_vpc_site--reference--group-002.md#canonical-3333011102112323-2110222010120203-0032223033000231-0112221131033331-2231222113200233-2301232032001332-0331231203133210-0030103320023020)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-0131000100012312-2121310130001102-2202230331103203-2230222003223001-1113101010113111-2131313030021223-2030122333203101-2213001121001122)

<a id="canonical-1122101220000020-0010221111301221-2313213102013113-1000203013113222-2203112310330010-2031201223130030-1222022103300121-1201003232001110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2100010123222222-2000330312230132-0202020001332232-0321201321010330-2130310223113321-1110031202311011-0102223002133213-1320220333032100"></a>

## ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets.IPv4 — IPv4 / 332101300233 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-0131000100012312-2121310130001102-2202230331103203-2230222003223001-1113101010113111-2131313030021223-2030122333203101-2213001121001122)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-2101111213103332-1031003102012122-3033333002321300-0212010323021100-0301331300303302-3321221003203222-1230133302203013-3232231332123212)
- [ingress_egress_gw](resources--gcp_vpc_site--reference--group-001.md#canonical-3123100130121022-2313112013232222-0213021230302120-3230300230213031-3101322300112331-0011013220032322-1311222312201210-2213020302010100)
- [ingress_egress_gw.inside_static_routes](resources--gcp_vpc_site--reference--group-002.md#canonical-0220321000013112-2030300310000323-2202131321000302-0112111112120211-0223122220232223-0023132200212001-3230213111203011-1330123303013023)
- [ingress_egress_gw.inside_static_routes.static_route_list](resources--gcp_vpc_site--reference--group-002.md#canonical-3112002300031122-1033032203133130-3320033210101302-1013211120120100-3021001213301031-1131300302030232-1123312000113231-0331101202321111)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route](resources--gcp_vpc_site--reference--group-002.md#canonical-3333011102112323-2110222010120203-0032223033000231-0112221131033331-2231222113200233-2301232032001332-0331231203133210-0030103320023020)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets](resources--gcp_vpc_site--reference--group-002.md#canonical-0233220023132033-0120221223320323-0132122022200220-0021223222210311-3020301231112332-3111211310000102-2300300121313311-0203211200320112)
- ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets.IPv4

<a id="canonical-0211110122301133-2230302231120112-2330212022221020-3123212110311220-3321031113202030-2113202333131213-3310210201000303-3223212320122032"></a>

Type: `"object"`. single nested block, Optional.

IPv4 subnets specified as prefix and prefix-length. Prefix length must be &lt;= 32.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
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

<a id="canonical-1330032003322222-3210322002230130-1002020123210120-2021223212012232-1310222002022010-0121010032220132-1300211200330022-2102302113133121"></a>

## Direct properties — IPv4 / 332101300233 / 3

<a id="canonical-3233013000222222-1202112220322112-2110100120132321-0213301222013233-2121310320332201-1213113023331030-2020222210201132-2222321332012213"></a>

<a id="canonical-3323333033022301-2131213003311330-1030122110203221-0001101300223110-2231000023200102-2232132300332200-1020112103230302-1303231212022121"></a>

## plen property — IPv4 / 332101300233 / 4

Type: `"number"`. Optional.

Prefix-length of the IPv4 subnet. Must be &lt;= 32.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.AtMost(32),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 32,
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
    "ves.io.schema.rules.uint32.lte": "32"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "32"
  }
}
```

<a id="canonical-0230011031302032-2311211013102332-1011230020331010-2011102101230023-2302221330221031-1120300202100133-1210130212103200-3111022330332013"></a>

<a id="canonical-1202222332300120-3130120113121310-1202013333022210-0110000110032132-3302333312122111-3021310110133013-3111321230300102-3122002031000310"></a>

## prefix property — IPv4 / 332101300233 / 5

Type: `"string"`. Optional.

Prefix part of the IPv4 subnet in string form with dot-decimal notation.

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

<a id="canonical-1131031103323220-3232233223030022-2121020112101111-3231303222203121-3202210120303303-1212122210112310-2132332032133330-0022023311302220"></a>

## Next pages — IPv4 / 332101300233 / 6

- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets](resources--gcp_vpc_site--reference--group-002.md#canonical-0233220023132033-0120221223320323-0132122022200220-0021223222210311-3020301231112332-3111211310000102-2300300121313311-0203211200320112)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-0131000100012312-2121310130001102-2202230331103203-2230222003223001-1113101010113111-2131313030021223-2030122333203101-2213001121001122)

<a id="canonical-0321213203212110-1123200000030301-2223303011330002-2333211310223331-0123011213131230-1010032313301123-2311010211301302-3220220131003033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2000301020230112-3221110000130000-1323313210332032-1211233000122233-1031232301023220-3001301201220322-1031021221231212-0130102233200100"></a>

## ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets.IPv6 — IPv6 / 211030302300 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-0131000100012312-2121310130001102-2202230331103203-2230222003223001-1113101010113111-2131313030021223-2030122333203101-2213001121001122)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-2101111213103332-1031003102012122-3033333002321300-0212010323021100-0301331300303302-3321221003203222-1230133302203013-3232231332123212)
- [ingress_egress_gw](resources--gcp_vpc_site--reference--group-001.md#canonical-3123100130121022-2313112013232222-0213021230302120-3230300230213031-3101322300112331-0011013220032322-1311222312201210-2213020302010100)
- [ingress_egress_gw.inside_static_routes](resources--gcp_vpc_site--reference--group-002.md#canonical-0220321000013112-2030300310000323-2202131321000302-0112111112120211-0223122220232223-0023132200212001-3230213111203011-1330123303013023)
- [ingress_egress_gw.inside_static_routes.static_route_list](resources--gcp_vpc_site--reference--group-002.md#canonical-3112002300031122-1033032203133130-3320033210101302-1013211120120100-3021001213301031-1131300302030232-1123312000113231-0331101202321111)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route](resources--gcp_vpc_site--reference--group-002.md#canonical-3333011102112323-2110222010120203-0032223033000231-0112221131033331-2231222113200233-2301232032001332-0331231203133210-0030103320023020)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets](resources--gcp_vpc_site--reference--group-002.md#canonical-0233220023132033-0120221223320323-0132122022200220-0021223222210311-3020301231112332-3111211310000102-2300300121313311-0203211200320112)
- ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets.IPv6

<a id="canonical-0020110003112021-0022202020320313-1202323312300201-1103212222102332-2203012030113100-3113320002113020-0211102012301023-1102220302212130"></a>

Type: `"object"`. single nested block, Optional.

IPv6 subnets specified as prefix and prefix-length. Prefix-legnth must be &lt;= 128.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
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

<a id="canonical-1213222011120223-2322330013332102-1131331322300302-1022311011000322-1120133132032300-0132301021203032-2312122201020110-2233302222131122"></a>

## Direct properties — IPv6 / 211030302300 / 3

<a id="canonical-3010103222202103-0221220033131331-3003302013000021-2200101310130202-2311023100101021-0133023202302122-3022131313220200-0303123201023212"></a>

<a id="canonical-3300121101313132-1001200000033231-1312122131120122-0332232333020013-0210101120223231-3211010332123001-2111311222121011-2332102100233203"></a>

## plen property — IPv6 / 211030302300 / 4

Type: `"number"`. Optional.

Prefix length of the IPv6 subnet. Must be &lt;= 128.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.AtMost(128),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 128,
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
    "ves.io.schema.rules.uint32.lte": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "128"
  }
}
```

<a id="canonical-1100112102210001-0333300200031321-0211301330131232-1130003301021320-2022331020313000-0310201112201012-3112202000223021-2310113330321200"></a>

<a id="canonical-3201231132320112-3103000222030231-1201132030222302-0233301012120212-1233300002122133-2222323113100311-0301131231131013-1330310320101222"></a>

## prefix property — IPv6 / 211030302300 / 5

Type: `"string"`. Optional.

Prefix part of the IPv6 subnet given in form of string. IPv6 address must be specified as
hexadecimal numbers separated by ':' e.g. '2001:db8:0:0:0:2:0:0' The address can be compacted by
suppressing zeros e.g. '2001:db8::2::'.

Upstream description:

Prefix part of the IPv6 subnet given in form of string. IPv6 address must be specified as
hexadecimal numbers separated by ':' e.g. "2001:db8:0:0:0:2:0:0" The address can be compacted by
suppressing zeros e.g. "2001:db8::2::"

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

<a id="canonical-1100101021020032-0221332013002100-1301203211200221-2002121222133021-2311231221123110-3031313121211222-3311111302303322-0311130211011122"></a>

## Next pages — IPv6 / 211030302300 / 6

- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets](resources--gcp_vpc_site--reference--group-002.md#canonical-0233220023132033-0120221223320323-0132122022200220-0021223222210311-3020301231112332-3111211310000102-2300300121313311-0203211200320112)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-0131000100012312-2121310130001102-2202230331103203-2230222003223001-1113101010113111-2131313030021223-2030122333203101-2213001121001122)

<a id="canonical-3222001132100012-1200312003312120-3302312333232311-2232110003311200-1302022010113013-1222012220132232-2322210000001333-2032110122033120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1223020130002033-0321033230232101-1223002011200000-2020111333321121-2123130230310223-1212310230020201-2212030002223222-1221201033001013"></a>

## ingress_egress_gw.inside_subnet — inside_subnet / 301310131133 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-0131000100012312-2121310130001102-2202230331103203-2230222003223001-1113101010113111-2131313030021223-2030122333203101-2213001121001122)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-2101111213103332-1031003102012122-3033333002321300-0212010323021100-0301331300303302-3321221003203222-1230133302203013-3232231332123212)
- [ingress_egress_gw](resources--gcp_vpc_site--reference--group-001.md#canonical-3123100130121022-2313112013232222-0213021230302120-3230300230213031-3101322300112331-0011013220032322-1311222312201210-2213020302010100)
- ingress_egress_gw.inside_subnet

<a id="canonical-1002002222303002-1033332022232012-1132323323301123-3000001323010122-0330000310322033-3013111311001333-1311302322001003-3031333233203121"></a>

Type: `"object"`. single nested block, Optional.

Defines choice about GCP VPC network for a view.

Upstream description:

This defines choice about GCP VPC network for a view.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("existing_subnet",
    "new_subnet")}
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
  "x-ves-oneof-field-choice": "[\"existing_subnet\",\"new_subnet\"]"
}
```

Terraform syntax:

```terraform
inside_subnet {
  # Configure direct properties listed below.
}
```

<a id="canonical-1031202302022132-0101000202133132-2230100013311022-3112200101033110-0012220331223123-2311000202013333-2311203032323203-0101111203003121"></a>

## Direct properties — inside_subnet / 301310131133 / 3

- [existing_subnet](resources--gcp_vpc_site--reference--group-002.md#canonical-1000322103211221-1112011201320010-0333201331033220-2231001323112111-0331321301030211-1012100102312313-1010230101033132-3030020121111113): complete subsection reference.

- [new_subnet](resources--gcp_vpc_site--reference--group-002.md#canonical-2301120001123010-1211300133130112-0301110202030121-0320010130033122-1223103320303031-0220232311030122-0323211110211303-0000002020311211): complete subsection reference.

<a id="canonical-2321333012123132-3201112221302003-1031102121113021-0130033000213020-3102122301220203-3110033222122132-1023020301120310-1021130311230123"></a>

## Next pages — inside_subnet / 301310131133 / 4

- [ingress_egress_gw.inside_subnet.existing_subnet](resources--gcp_vpc_site--reference--group-002.md#canonical-1000322103211221-1112011201320010-0333201331033220-2231001323112111-0331321301030211-1012100102312313-1010230101033132-3030020121111113)
- [ingress_egress_gw.inside_subnet.new_subnet](resources--gcp_vpc_site--reference--group-002.md#canonical-2301120001123010-1211300133130112-0301110202030121-0320010130033122-1223103320303031-0220232311030122-0323211110211303-0000002020311211)
- [ingress_egress_gw](resources--gcp_vpc_site--reference--group-001.md#canonical-3123100130121022-2313112013232222-0213021230302120-3230300230213031-3101322300112331-0011013220032322-1311222312201210-2213020302010100)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-0131000100012312-2121310130001102-2202230331103203-2230222003223001-1113101010113111-2131313030021223-2030122333203101-2213001121001122)

<a id="canonical-1000322103211221-1112011201320010-0333201331033220-2231001323112111-0331321301030211-1012100102312313-1010230101033132-3030020121111113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2132230001123111-1032320220123111-0011131120223131-1112332022010133-1231010103030121-3130311312121331-2230231313021122-2121132303130101"></a>

## ingress_egress_gw.inside_subnet.existing_subnet — existing_subnet / 133032332001 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-0131000100012312-2121310130001102-2202230331103203-2230222003223001-1113101010113111-2131313030021223-2030122333203101-2213001121001122)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-2101111213103332-1031003102012122-3033333002321300-0212010323021100-0301331300303302-3321221003203222-1230133302203013-3232231332123212)
- [ingress_egress_gw](resources--gcp_vpc_site--reference--group-001.md#canonical-3123100130121022-2313112013232222-0213021230302120-3230300230213031-3101322300112331-0011013220032322-1311222312201210-2213020302010100)
- [ingress_egress_gw.inside_subnet](resources--gcp_vpc_site--reference--group-002.md#canonical-3222001132100012-1200312003312120-3302312333232311-2232110003311200-1302022010113013-1222012220132232-2322210000001333-2032110122033120)
- ingress_egress_gw.inside_subnet.existing_subnet

<a id="canonical-2313032121200102-0230111103200121-2312010320123330-1332222323102211-2231302022321331-1032231113211022-0121023311000233-3011030121110012"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for existing subnet.

Upstream description:

Name of existing GCP subnet.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("subnet_name")}
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
existing_subnet {
  # Configure direct properties listed below.
}
```

<a id="canonical-2300221302202312-0033022321123302-0033332200223111-3102112032032311-3020021000103131-0003221312132001-0011203000133030-3323320013002030"></a>

## Direct properties — existing_subnet / 133032332001 / 3

<a id="canonical-0101020313230300-0102321212111323-0111300120202132-2123302021123320-1122212032231130-3311201120121233-1300302211000213-3303222223302123"></a>

<a id="canonical-1101231233213201-3311121112010303-2021113200323130-2102230123321321-1131031200233021-0233233330133113-2032032333201001-3011002001022032"></a>

## subnet_name property — existing_subnet / 133032332001 / 4

Type: `"string"`. Optional.

VPC Subnet Name. Name of your subnet in VPC network.

Upstream description:

Name of your subnet in VPC network.

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

<a id="canonical-1111001111101103-2133323000231210-3012330103211200-3321230100012013-2003312202300231-1112133233220232-3300021013323102-2333202220133330"></a>

## Next pages — existing_subnet / 133032332001 / 5

- [ingress_egress_gw.inside_subnet](resources--gcp_vpc_site--reference--group-002.md#canonical-3222001132100012-1200312003312120-3302312333232311-2232110003311200-1302022010113013-1222012220132232-2322210000001333-2032110122033120)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-0131000100012312-2121310130001102-2202230331103203-2230222003223001-1113101010113111-2131313030021223-2030122333203101-2213001121001122)

<a id="canonical-2301120001123010-1211300133130112-0301110202030121-0320010130033122-1223103320303031-0220232311030122-0323211110211303-0000002020311211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3130222030012302-0111010011311313-2221323201322313-3330013022131131-0210133020030300-0332311231000320-1131211300330032-3112233311100111"></a>

## ingress_egress_gw.inside_subnet.new_subnet — new_subnet / 320021210002 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-0131000100012312-2121310130001102-2202230331103203-2230222003223001-1113101010113111-2131313030021223-2030122333203101-2213001121001122)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-2101111213103332-1031003102012122-3033333002321300-0212010323021100-0301331300303302-3321221003203222-1230133302203013-3232231332123212)
- [ingress_egress_gw](resources--gcp_vpc_site--reference--group-001.md#canonical-3123100130121022-2313112013232222-0213021230302120-3230300230213031-3101322300112331-0011013220032322-1311222312201210-2213020302010100)
- [ingress_egress_gw.inside_subnet](resources--gcp_vpc_site--reference--group-002.md#canonical-3222001132100012-1200312003312120-3302312333232311-2232110003311200-1302022010113013-1222012220132232-2322210000001333-2032110122033120)
- ingress_egress_gw.inside_subnet.new_subnet

<a id="canonical-1230013030223311-3311113110002232-2203032300322310-2220211133121032-0231010212102122-0013131102332302-1000113110222133-3123032012011110"></a>

Type: `"object"`. single nested block, Optional.

GCP subnet parameters Type. Parameters for GCP subnet.

Upstream description:

Parameters for GCP subnet.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("primary_ipv4")}
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
new_subnet {
  # Configure direct properties listed below.
}
```

<a id="canonical-3130121230133023-1222032303332112-0202121030100200-3203133013110221-2200032320321020-2033311120213030-1112223210032301-2013213202330001"></a>

## Direct properties — new_subnet / 320021210002 / 3

<a id="canonical-3333130303110311-3320133302131120-3303322123120300-3101300102311000-3222323320331212-2331313121100333-0121311133211003-1012210013202103"></a>

<a id="canonical-3000111003010330-0121130122311311-3223313020332102-2101313230203133-2013233030322002-3023220133131233-1232032233303302-2001033011230033"></a>

## primary_ipv4 property — new_subnet / 320021210002 / 4

Type: `"string"`. Optional.

IPv4 prefix for this Subnet. It has to be private address space.

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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.ipv4_prefix": "true",
    "ves.io.schema.rules.string.max_ip_prefix_length": "28",
    "ves.io.schema.rules.string.min_ip_prefix_length": "8"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.ipv4_prefix": "true",
    "ves.io.schema.rules.string.max_ip_prefix_length": "28",
    "ves.io.schema.rules.string.min_ip_prefix_length": "8"
  }
}
```

<a id="canonical-3332111002023131-3001200120002310-1333032301100101-2120333020333200-1233011300323221-1320122113123311-2133303021331031-3211332110302322"></a>

<a id="canonical-3233123310003121-1021110110231331-0001222133300213-2130333322313000-3102031213131231-0132010213201301-0302222323222010-3111330120101202"></a>

## subnet_name property — new_subnet / 320021210002 / 5

Type: `"string"`. Optional.

Name of new VPC Subnet, will be autogenerated if empty.

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

<a id="canonical-3010133311301203-2120002322101020-3011122111303013-0212020020212322-3333221131011102-2022000111131312-0113112323200223-1000002101200121"></a>

## Next pages — new_subnet / 320021210002 / 6

- [ingress_egress_gw.inside_subnet](resources--gcp_vpc_site--reference--group-002.md#canonical-3222001132100012-1200312003312120-3302312333232311-2232110003311200-1302022010113013-1222012220132232-2322210000001333-2032110122033120)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-0131000100012312-2121310130001102-2202230331103203-2230222003223001-1113101010113111-2131313030021223-2030122333203101-2213001121001122)

<a id="canonical-0223121212322031-2211320312222312-2300130323120123-2121231012031031-3333321103332333-2002332220233102-1122031312200130-2200321033313231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1003030222302310-2330333031220010-2002123201302111-0130002313002313-1032231023031211-1200233232330013-3133210011203031-0321032023203203"></a>

## ingress_egress_gw.no_dc_cluster_group — no_dc_cluster_group / 010200332001 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-0131000100012312-2121310130001102-2202230331103203-2230222003223001-1113101010113111-2131313030021223-2030122333203101-2213001121001122)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-2101111213103332-1031003102012122-3033333002321300-0212010323021100-0301331300303302-3321221003203222-1230133302203013-3232231332123212)
- [ingress_egress_gw](resources--gcp_vpc_site--reference--group-001.md#canonical-3123100130121022-2313112013232222-0213021230302120-3230300230213031-3101322300112331-0011013220032322-1311222312201210-2213020302010100)
- ingress_egress_gw.no_dc_cluster_group

<a id="canonical-1110303003300303-2300033003302303-2220201231330120-2321123020302000-3333122221010013-1310112220332313-2223112201010233-1220121332330230"></a>

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
no_dc_cluster_group = {}
```

<a id="canonical-1311020322121133-1220330022003120-2112302020312112-3220300030310132-3303201323231331-2132222000103212-3130120011223303-2320332120103331"></a>

## Direct properties — no_dc_cluster_group / 010200332001 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1212003221002201-2003312033021230-0131331111000310-2000232232300010-2123222211201122-3032323202231222-0231032103033001-0110311102303103"></a>

## Next pages — no_dc_cluster_group / 010200332001 / 4

- [ingress_egress_gw](resources--gcp_vpc_site--reference--group-001.md#canonical-3123100130121022-2313112013232222-0213021230302120-3230300230213031-3101322300112331-0011013220032322-1311222312201210-2213020302010100)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-0131000100012312-2121310130001102-2202230331103203-2230222003223001-1113101010113111-2131313030021223-2030122333203101-2213001121001122)

<a id="canonical-1121332000003013-1222000333333223-2121210310100010-2312003033100003-2111220103101112-0133232123131210-0311203213021300-2133223203320312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2011222022012323-0211000320210102-2003132010211012-3111212212032003-2203312001202021-0120303211202112-3132332033311122-1332213231120220"></a>

## ingress_egress_gw.no_forward_proxy — no_forward_proxy / 000202000032 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-0131000100012312-2121310130001102-2202230331103203-2230222003223001-1113101010113111-2131313030021223-2030122333203101-2213001121001122)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-2101111213103332-1031003102012122-3033333002321300-0212010323021100-0301331300303302-3321221003203222-1230133302203013-3232231332123212)
- [ingress_egress_gw](resources--gcp_vpc_site--reference--group-001.md#canonical-3123100130121022-2313112013232222-0213021230302120-3230300230213031-3101322300112331-0011013220032322-1311222312201210-2213020302010100)
- ingress_egress_gw.no_forward_proxy

<a id="canonical-3102003113223311-0133133332230333-0103321130202321-3302301010332120-3110021112012230-2333312213013203-0203000302232033-0012330001200023"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for no forward proxy.

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
no_forward_proxy = {}
```

<a id="canonical-2003103321133113-1101302123100011-0231132031311200-1221122203012211-1123231013022310-1123110210331011-2130123111033132-3030111021312001"></a>

## Direct properties — no_forward_proxy / 000202000032 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2303102220312132-1323302011300102-2132213331200311-1313201013231003-2033332020131323-2111111110010031-2300233113330021-1201213202021010"></a>

## Next pages — no_forward_proxy / 000202000032 / 4

- [ingress_egress_gw](resources--gcp_vpc_site--reference--group-001.md#canonical-3123100130121022-2313112013232222-0213021230302120-3230300230213031-3101322300112331-0011013220032322-1311222312201210-2213020302010100)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-0131000100012312-2121310130001102-2202230331103203-2230222003223001-1113101010113111-2131313030021223-2030122333203101-2213001121001122)

<a id="canonical-3102021312020122-3021300021202033-0210332020111302-3000312313002100-3022110003022221-1123023001332030-0310210102222211-3103101010200133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2312212200331233-2012120323023113-3032112112233001-0111202311300321-3311121112323113-2202320021211002-3201110132222212-0012212232210112"></a>

## ingress_egress_gw.no_global_network — no_global_network / 310032233020 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-0131000100012312-2121310130001102-2202230331103203-2230222003223001-1113101010113111-2131313030021223-2030122333203101-2213001121001122)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-2101111213103332-1031003102012122-3033333002321300-0212010323021100-0301331300303302-3321221003203222-1230133302203013-3232231332123212)
- [ingress_egress_gw](resources--gcp_vpc_site--reference--group-001.md#canonical-3123100130121022-2313112013232222-0213021230302120-3230300230213031-3101322300112331-0011013220032322-1311222312201210-2213020302010100)
- ingress_egress_gw.no_global_network

<a id="canonical-3012032331120131-0123300301311331-3333110130031032-0033010211103203-2203331101221222-0100020013031302-3012323101120232-2220002303311233"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for no global network.

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
no_global_network = {}
```

<a id="canonical-1020201310122231-0220112332110331-2302231321310101-0033201330300320-0320111222010011-2213032313101121-2212022200330300-0213131011211231"></a>

## Direct properties — no_global_network / 310032233020 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3010032121213022-3102320233311101-0102223000111002-2333212201003322-1132322210001221-0022020201322100-3103013103302301-3323123022323011"></a>

## Next pages — no_global_network / 310032233020 / 4

- [ingress_egress_gw](resources--gcp_vpc_site--reference--group-001.md#canonical-3123100130121022-2313112013232222-0213021230302120-3230300230213031-3101322300112331-0011013220032322-1311222312201210-2213020302010100)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-0131000100012312-2121310130001102-2202230331103203-2230222003223001-1113101010113111-2131313030021223-2030122333203101-2213001121001122)

<a id="canonical-0201103101130210-0133000033010210-1231030122220020-3320020013012113-1203001131113000-1103210221011202-0331103332210322-3130012302231213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1121212300011103-0323331012022321-0013121001022103-3311320212313002-2021303113222332-3131131231023201-3021202013220101-1323310323332030"></a>

## ingress_egress_gw.no_inside_static_routes — no_inside_static_routes / 202021123033 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-0131000100012312-2121310130001102-2202230331103203-2230222003223001-1113101010113111-2131313030021223-2030122333203101-2213001121001122)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-2101111213103332-1031003102012122-3033333002321300-0212010323021100-0301331300303302-3321221003203222-1230133302203013-3232231332123212)
- [ingress_egress_gw](resources--gcp_vpc_site--reference--group-001.md#canonical-3123100130121022-2313112013232222-0213021230302120-3230300230213031-3101322300112331-0011013220032322-1311222312201210-2213020302010100)
- ingress_egress_gw.no_inside_static_routes

<a id="canonical-0001210102231111-1222103010132311-2122232301323332-2213132300032113-0303121201210320-2113032110003130-0301033303210030-1102003220032232"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for no inside static routes.

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
no_inside_static_routes = {}
```

<a id="canonical-2002000303321123-1220311001033303-1210230211311012-0113331101133222-3032331311303332-0133213321110122-1323121101213321-0102121133200332"></a>

## Direct properties — no_inside_static_routes / 202021123033 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1332220031022333-1231003230300010-2301300023223332-2322212001323130-3031222101130103-2211310000222003-1103122120013130-3130101020301332"></a>

## Next pages — no_inside_static_routes / 202021123033 / 4

- [ingress_egress_gw](resources--gcp_vpc_site--reference--group-001.md#canonical-3123100130121022-2313112013232222-0213021230302120-3230300230213031-3101322300112331-0011013220032322-1311222312201210-2213020302010100)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-0131000100012312-2121310130001102-2202230331103203-2230222003223001-1113101010113111-2131313030021223-2030122333203101-2213001121001122)

<a id="canonical-1302131011321212-1300221023200122-2300210033331311-0113012022102331-3012030100001333-1111010120111102-1013112203202320-0213131013301310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2221023121211321-1111223330001323-3022233213101311-3132230112120023-1030103321001000-0023023002200123-3001030000013311-0313301332322231"></a>

## ingress_egress_gw.no_network_policy — no_network_policy / 020323020100 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-0131000100012312-2121310130001102-2202230331103203-2230222003223001-1113101010113111-2131313030021223-2030122333203101-2213001121001122)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-2101111213103332-1031003102012122-3033333002321300-0212010323021100-0301331300303302-3321221003203222-1230133302203013-3232231332123212)
- [ingress_egress_gw](resources--gcp_vpc_site--reference--group-001.md#canonical-3123100130121022-2313112013232222-0213021230302120-3230300230213031-3101322300112331-0011013220032322-1311222312201210-2213020302010100)
- ingress_egress_gw.no_network_policy

<a id="canonical-0320101233130300-3320331110011121-3212023231133032-2012311203220301-0010221320330301-2101301010222330-0120320132102021-1321121122012000"></a>

Type: `["object", {}]`. Optional.

Policy configuration for this feature.

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
no_network_policy = {}
```

<a id="canonical-3303020200230012-3013030101033123-1211030203322311-3011002121310320-2310302220331100-1023111223112311-2313131310201133-1331010231021033"></a>

## Direct properties — no_network_policy / 020323020100 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2002000022012011-3020301102312121-0000101223013001-2103030221332112-3333002022212233-3120200230002201-3003110033003001-0303231332212310"></a>

## Next pages — no_network_policy / 020323020100 / 4

- [ingress_egress_gw](resources--gcp_vpc_site--reference--group-001.md#canonical-3123100130121022-2313112013232222-0213021230302120-3230300230213031-3101322300112331-0011013220032322-1311222312201210-2213020302010100)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-0131000100012312-2121310130001102-2202230331103203-2230222003223001-1113101010113111-2131313030021223-2030122333203101-2213001121001122)

<a id="canonical-1122203313010332-0123231301322313-2023101323010232-0202300003230303-1123123013113301-2230221312331320-2113021110010331-3223302020032210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3201031120310231-0321132232232020-0000001203332201-1210303010120230-1313110003212000-3012321121230222-1122313010203000-3331323001223110"></a>

## ingress_egress_gw.no_outside_static_routes — no_outside_static_routes / 230012202321 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-0131000100012312-2121310130001102-2202230331103203-2230222003223001-1113101010113111-2131313030021223-2030122333203101-2213001121001122)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-2101111213103332-1031003102012122-3033333002321300-0212010323021100-0301331300303302-3321221003203222-1230133302203013-3232231332123212)
- [ingress_egress_gw](resources--gcp_vpc_site--reference--group-001.md#canonical-3123100130121022-2313112013232222-0213021230302120-3230300230213031-3101322300112331-0011013220032322-1311222312201210-2213020302010100)
- ingress_egress_gw.no_outside_static_routes

<a id="canonical-2113001200133131-1112322132232023-0120020231113001-2112133003331000-0202021010222221-3011130233330310-3032002132001222-2103121110030230"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for no outside static routes.

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
no_outside_static_routes = {}
```

<a id="canonical-3101221012002201-2133211213233322-1010020320323212-2101220031321100-3031312112020201-3213131122211223-3311233213212323-0211220232023031"></a>

## Direct properties — no_outside_static_routes / 230012202321 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0112112122133301-1120303100112333-1001230120212103-1000120000103010-3310220222301022-3010200133023210-1101131031230120-3333011202133113"></a>

## Next pages — no_outside_static_routes / 230012202321 / 4

- [ingress_egress_gw](resources--gcp_vpc_site--reference--group-001.md#canonical-3123100130121022-2313112013232222-0213021230302120-3230300230213031-3101322300112331-0011013220032322-1311222312201210-2213020302010100)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-0131000100012312-2121310130001102-2202230331103203-2230222003223001-1113101010113111-2131313030021223-2030122333203101-2213001121001122)

<a id="canonical-1211123322330233-3312333012002031-1012211321132000-0012003132301223-2001313132331302-3033013021331012-0200130302231001-1003322233332332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3232321212311002-3003201302223103-1300122020322312-3131210312020031-3101110033032100-2220203312022231-1331203110002121-0211232333303003"></a>

## ingress_egress_gw.outside_network — outside_network / 301230312330 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-0131000100012312-2121310130001102-2202230331103203-2230222003223001-1113101010113111-2131313030021223-2030122333203101-2213001121001122)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-2101111213103332-1031003102012122-3033333002321300-0212010323021100-0301331300303302-3321221003203222-1230133302203013-3232231332123212)
- [ingress_egress_gw](resources--gcp_vpc_site--reference--group-001.md#canonical-3123100130121022-2313112013232222-0213021230302120-3230300230213031-3101322300112331-0011013220032322-1311222312201210-2213020302010100)
- ingress_egress_gw.outside_network

<a id="canonical-3203312133220231-3022221123221333-2023111002132232-2032022120111022-2210223002301200-0023210201022223-3112201223113331-1133020302101221"></a>

Type: `"object"`. single nested block, Optional.

Defines choice about GCP VPC network for a view.

Upstream description:

This defines choice about GCP VPC network for a view.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("existing_network",
    "new_network"),
  validators.ConflictingObjectAttributes("existing_network",
    "new_network_autogenerate"),
  validators.ConflictingObjectAttributes("new_network",
    "new_network_autogenerate")}
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
  "x-ves-oneof-field-choice": "[\"existing_network\",\"new_network\",\"new_network_autogenerate\"]"
}
```

Terraform syntax:

```terraform
outside_network {
  # Configure direct properties listed below.
}
```

<a id="canonical-1113323220110232-3132213210120012-1232212021322220-2221321120221103-2122011331230322-1220112131113010-0223110110232213-3333300322322200"></a>

## Direct properties — outside_network / 301230312330 / 3

- [existing_network](resources--gcp_vpc_site--reference--group-002.md#canonical-3233010001231002-3302002200022312-1213001102003121-2310232300323131-2013122131010231-1122322112303301-2113002133232102-1002112130123003): complete subsection reference.

- [new_network](resources--gcp_vpc_site--reference--group-002.md#canonical-2330002321121130-0323320232112320-0222232300030231-0123012020002130-1330332103212111-2132133133210221-1321022131201233-1321031321320322): complete subsection reference.

- [new_network_autogenerate](resources--gcp_vpc_site--reference--group-002.md#canonical-1310120310131113-0103212131211202-3211213100102032-2222233202232310-0202322223303030-1123032320231002-2123102211010331-3223132202010330): complete subsection reference.

<a id="canonical-1210122002332122-1031231310022201-3320012233200113-2022321311002233-3311211223233300-3102010311033010-2101330221000030-1100010011103230"></a>

## Next pages — outside_network / 301230312330 / 4

- [ingress_egress_gw.outside_network.existing_network](resources--gcp_vpc_site--reference--group-002.md#canonical-3233010001231002-3302002200022312-1213001102003121-2310232300323131-2013122131010231-1122322112303301-2113002133232102-1002112130123003)
- [ingress_egress_gw.outside_network.new_network](resources--gcp_vpc_site--reference--group-002.md#canonical-2330002321121130-0323320232112320-0222232300030231-0123012020002130-1330332103212111-2132133133210221-1321022131201233-1321031321320322)
- [ingress_egress_gw.outside_network.new_network_autogenerate](resources--gcp_vpc_site--reference--group-002.md#canonical-1310120310131113-0103212131211202-3211213100102032-2222233202232310-0202322223303030-1123032320231002-2123102211010331-3223132202010330)
- [ingress_egress_gw](resources--gcp_vpc_site--reference--group-001.md#canonical-3123100130121022-2313112013232222-0213021230302120-3230300230213031-3101322300112331-0011013220032322-1311222312201210-2213020302010100)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-0131000100012312-2121310130001102-2202230331103203-2230222003223001-1113101010113111-2131313030021223-2030122333203101-2213001121001122)

<a id="canonical-3233010001231002-3302002200022312-1213001102003121-2310232300323131-2013122131010231-1122322112303301-2113002133232102-1002112130123003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1020111311031111-0322300322300300-0230230220032232-1103323001122012-1033321133133012-3030330220023111-1330221121200132-2301122211321022"></a>

## ingress_egress_gw.outside_network.existing_network — existing_network / 203303210203 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-0131000100012312-2121310130001102-2202230331103203-2230222003223001-1113101010113111-2131313030021223-2030122333203101-2213001121001122)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-2101111213103332-1031003102012122-3033333002321300-0212010323021100-0301331300303302-3321221003203222-1230133302203013-3232231332123212)
- [ingress_egress_gw](resources--gcp_vpc_site--reference--group-001.md#canonical-3123100130121022-2313112013232222-0213021230302120-3230300230213031-3101322300112331-0011013220032322-1311222312201210-2213020302010100)
- [ingress_egress_gw.outside_network](resources--gcp_vpc_site--reference--group-002.md#canonical-1211123322330233-3312333012002031-1012211321132000-0012003132301223-2001313132331302-3033013021331012-0200130302231001-1003322233332332)
- ingress_egress_gw.outside_network.existing_network

<a id="canonical-0201313200130011-0323213333231121-0231110013103313-2230012232221310-1102021312112320-2010201302032203-2001100122333111-3132121322122202"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for existing network.

Upstream description:

Name of existing VPC network.

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
  },
  "x-ves-oneof-field-routing_type": "[]"
}
```

Terraform syntax:

```terraform
existing_network {
  # Configure direct properties listed below.
}
```

<a id="canonical-3212330003002300-1023030131110003-2223033012231003-1023202323003300-3200033011311221-3210121233213220-0201001101011213-2031000322331033"></a>

## Direct properties — existing_network / 203303210203 / 3

<a id="canonical-0202321222332323-3012233031030221-3121233132223102-3320100001003200-0231302312003132-1110130312121330-1101122211222210-3220310321331200"></a>

<a id="canonical-2322023013220332-0301012222121323-2011030010002333-0302313231132331-1112012213203330-2212330232120223-1323010223313321-2222031311301232"></a>

## name property — existing_network / 203303210203 / 4

Type: `"string"`. Optional.

GCP VPC Network Name. Name for your GCP VPC Network.

Upstream description:

Name for your GCP VPC Network.

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
    "maxLength": 64,
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

<a id="canonical-0010313032301200-0121102123122220-2232121022213020-0103202220231021-3033311001232311-2203133122023213-0011310012012200-2100210121320332"></a>

## Next pages — existing_network / 203303210203 / 5

- [ingress_egress_gw.outside_network](resources--gcp_vpc_site--reference--group-002.md#canonical-1211123322330233-3312333012002031-1012211321132000-0012003132301223-2001313132331302-3033013021331012-0200130302231001-1003322233332332)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-0131000100012312-2121310130001102-2202230331103203-2230222003223001-1113101010113111-2131313030021223-2030122333203101-2213001121001122)

<a id="canonical-2330002321121130-0323320232112320-0222232300030231-0123012020002130-1330332103212111-2132133133210221-1321022131201233-1321031321320322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1122211323002211-3213132113031103-3003301022331103-1202032102123110-3301133000003100-0011120231302202-3331131212103211-0220100102010211"></a>

## ingress_egress_gw.outside_network.new_network — new_network / 121302221131 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-0131000100012312-2121310130001102-2202230331103203-2230222003223001-1113101010113111-2131313030021223-2030122333203101-2213001121001122)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-2101111213103332-1031003102012122-3033333002321300-0212010323021100-0301331300303302-3321221003203222-1230133302203013-3232231332123212)
- [ingress_egress_gw](resources--gcp_vpc_site--reference--group-001.md#canonical-3123100130121022-2313112013232222-0213021230302120-3230300230213031-3101322300112331-0011013220032322-1311222312201210-2213020302010100)
- [ingress_egress_gw.outside_network](resources--gcp_vpc_site--reference--group-002.md#canonical-1211123322330233-3312333012002031-1012211321132000-0012003132301223-2001313132331302-3033013021331012-0200130302231001-1003322233332332)
- ingress_egress_gw.outside_network.new_network

<a id="canonical-2201123223203032-3301012133102120-3322100331002123-2300313110322023-3031001230031103-2311333221012222-2320112000222330-0330032030031120"></a>

Type: `"object"`. single nested block, Optional.

Parameters to create a new GCP VPC Network.

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
new_network {
  # Configure direct properties listed below.
}
```

<a id="canonical-3000111331010230-2313200032201321-1313332201301100-1212230133232133-1313101322212301-2032031310030000-2203221030012323-2122210110110301"></a>

## Direct properties — new_network / 121302221131 / 3

<a id="canonical-3031120022012023-2121133022110023-1031110003320132-3011033202132000-2130000223023310-1203102333321302-1012023321002201-1203312332303020"></a>

<a id="canonical-2022313202030300-3301013301111111-2201103011223200-1111213312110111-3303313130002032-2212230103000103-2123113310122110-1121101332011101"></a>

## name property — new_network / 121302221131 / 4

Type: `"string"`. Optional.

GCP VPC Network Name. Name for your GCP VPC Network.

Upstream description:

Name for your GCP VPC Network.

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
    "maxLength": 64,
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

<a id="canonical-0013111110113122-1000130022323223-1110330131233202-3303013322222300-0130213213010112-0223210203220203-1310023111020023-2012002222130123"></a>

## Next pages — new_network / 121302221131 / 5

- [ingress_egress_gw.outside_network](resources--gcp_vpc_site--reference--group-002.md#canonical-1211123322330233-3312333012002031-1012211321132000-0012003132301223-2001313132331302-3033013021331012-0200130302231001-1003322233332332)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-0131000100012312-2121310130001102-2202230331103203-2230222003223001-1113101010113111-2131313030021223-2030122333203101-2213001121001122)

<a id="canonical-1310120310131113-0103212131211202-3211213100102032-2222233202232310-0202322223303030-1123032320231002-2123102211010331-3223132202010330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->
