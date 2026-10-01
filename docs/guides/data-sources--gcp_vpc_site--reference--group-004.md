---
page_title: "xcsh_gcp_vpc_site reference"
subcategory: "Infrastructure"
description: "Complete grouped canonical reference for xcsh_gcp_vpc_site reference."
---

# xcsh_gcp_vpc_site reference

<a id="canonical-0113130013311013-3111132033323120-2313100231130121-1321333020110332-1223002010033030-1022321300112011-2302320321313031-0230110033311012"></a>

## sw.default_sw_version — default_sw_version / 112322102302 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-1132222133321101-1232102200201102-2111112102132122-0010023221133123-2333213300031302-1001011022311003-3121103220313010-2103321200303202)
- [Property reference](data-sources--gcp_vpc_site--reference--group-001.md#canonical-0023231003100101-1010210121112300-0303103132302303-0331000323201313-3003311210331022-3000312312212303-0030202320102022-1232013113103201)
- [sw](data-sources--gcp_vpc_site--reference--group-003.md#canonical-0310330212021223-3303111001132333-2110220022330203-2030023332031212-0003311012011331-0210223302200113-1032100302211312-2001221222220020)
- sw.default_sw_version

<a id="canonical-0132233033323202-1310331310021111-2121232112202121-1202121333223203-2001033123100213-0200202030133023-3101210230330033-0333212311310302"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-2221202103013313-1103100211013221-2121002331201201-3101310002102102-1111113112121132-0032033033010302-1332322220103131-3020110321231122"></a>

## Direct properties — default_sw_version / 112322102302 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0020022223032301-0332311113232100-3303122112100313-3111003121100110-1001201011312033-0102201302233320-3000230011023202-2020032322031201"></a>

## Next pages — default_sw_version / 112322102302 / 4

- [sw](data-sources--gcp_vpc_site--reference--group-003.md#canonical-0310330212021223-3303111001132333-2110220022330203-2030023332031212-0003311012011331-0210223302200113-1032100302211312-2001221222220020)
- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-1132222133321101-1232102200201102-2111112102132122-0010023221133123-2333213300031302-1001011022311003-3121103220313010-2103321200303202)

<a id="canonical-2213313332313212-1230123320203012-3120121021331113-3210101220221000-1023022133330131-3023333003003323-2311003313332333-2203303111102220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1030120132320100-0322221012111320-3322011303323030-3330310012032210-1000231021230121-0122012130311121-3220223130000233-1200212121201122"></a>

## voltstack_cluster — voltstack_cluster / 230131221201 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-1132222133321101-1232102200201102-2111112102132122-0010023221133123-2333213300031302-1001011022311003-3121103220313010-2103321200303202)
- [Property reference](data-sources--gcp_vpc_site--reference--group-001.md#canonical-0023231003100101-1010210121112300-0303103132302303-0331000323201313-3003311210331022-3000312312212303-0030202320102022-1232013113103201)
- voltstack_cluster

<a id="canonical-0203133203303222-3213130300030211-2023033333200210-0202320122220013-3100113312003032-1000031013121311-1032330000300302-1022103102203101"></a>

Type: `"single"`. Computed.

App Stack cluster of single interface GCP site.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-dc_cluster_group_choice": "[\"dc_cluster_group\",\"no_dc_cluster_group\"]",
  "x-ves-oneof-field-forward_proxy_choice": "[\"active_forward_proxy_policies\",\"forward_proxy_allow_all\",\"no_forward_proxy\"]",
  "x-ves-oneof-field-global_network_choice": "[\"global_network_list\",\"no_global_network\"]",
  "x-ves-oneof-field-k8s_cluster_choice": "[\"k8s_cluster\",\"no_k8s_cluster\"]",
  "x-ves-oneof-field-network_policy_choice": "[\"active_enhanced_firewall_policies\",\"active_network_policies\",\"no_network_policy\"]",
  "x-ves-oneof-field-outside_static_route_choice": "[\"no_outside_static_routes\",\"outside_static_routes\"]",
  "x-ves-oneof-field-site_mesh_group_choice": "[\"sm_connection_public_ip\",\"sm_connection_pvt_ip\"]",
  "x-ves-oneof-field-storage_class_choice": "[\"default_storage\",\"storage_class_list\"]"
}
```

<a id="canonical-1021000331302220-0333022201020332-2120020031330303-2000220110322332-2113200023101030-2013202211000212-0002010130230022-3231002303321233"></a>

## Direct properties — voltstack_cluster / 230131221201 / 3

- [active_enhanced_firewall_policies](data-sources--gcp_vpc_site--reference--group-004.md#canonical-0232022202333003-2202033311020122-0013013023122033-0301220230222203-3013000111020120-0001131112012210-2131313320222313-0130122022102322): complete subsection reference.

- [active_forward_proxy_policies](data-sources--gcp_vpc_site--reference--group-004.md#canonical-3000133322203121-1203100320022130-3302113010101322-2022212111130022-0331020321113210-1112313300022331-0123230330213231-3230132100300020): complete subsection reference.

- [active_network_policies](data-sources--gcp_vpc_site--reference--group-004.md#canonical-1003102311022230-1230310103302103-3131100223313320-0300310013321331-2122022003321011-1302022311023210-0333313301103012-2323222210102333): complete subsection reference.

- [dc_cluster_group](data-sources--gcp_vpc_site--reference--group-004.md#canonical-2122123121313303-1002012232131000-1200211200300230-1012313010322103-1232230220220102-2233310202223231-0002222311333333-0020003310133132): complete subsection reference.

- [default_storage](data-sources--gcp_vpc_site--reference--group-004.md#canonical-3120222212131221-0000211132101132-3200332231101123-3312031023220323-2211103012232103-2202133010133203-0332002202213011-1111230220320131): complete subsection reference.

- [forward_proxy_allow_all](data-sources--gcp_vpc_site--reference--group-004.md#canonical-0211331210122330-0313020312231321-0201212032112212-1001003121101001-0013023310013032-1231310020232312-2222211202331132-2321321233200130): complete subsection reference.

<a id="canonical-3330231013222232-3321313230011130-0032321332331231-1213203102220100-0101111323111112-3022123301211323-3002312003220231-0001111121302013"></a>

<a id="canonical-0200311131002103-1002020103103030-2203130203120321-3323302221223021-2212103220103332-2032320100322113-2101111310201002-0323320111330101"></a>

## gcp_certified_hw property — voltstack_cluster / 230131221201 / 4

Type: `"string"`. Computed.

\[Enum: gcp-byol-voltstack-combo\] GCP Certified Hardware. Name for GCP certified hardware. The only
possible value is \`gcp-byol-voltstack-combo\`.

Upstream description:

Name for GCP certified hardware.

Receipt-pinned upstream constraints:

```json
{
  "enum": [
    "gcp-byol-voltstack-combo"
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
    "ves.io.schema.rules.string.in": "[\\\"gcp-byol-voltstack-combo\\\"]",
    "ves.io.schema.rules.string.max_len": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.in": "[\\\"gcp-byol-voltstack-combo\\\"]",
    "ves.io.schema.rules.string.max_len": "64"
  }
}
```

<a id="canonical-0102121311130021-1031012132233013-1231301111110112-3312130101231322-1312320311223301-1222032223203313-0202323112103230-3020203312301021"></a>

<a id="canonical-3101011003133211-1223011210103032-2231131311033113-0001101022003030-2233033100130011-3031001112123033-1313032100030313-0303002131021232"></a>

## gcp_zone_names property — voltstack_cluster / 230131221201 / 5

Type: `["list", "string"]`. Computed.

X-required List of zones when instances will be created, needs to match with region selected.

Upstream description:

X-required List of zones when instances will be created, needs to match with region selected.

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

- [global_network_list](data-sources--gcp_vpc_site--reference--group-004.md#canonical-2312310111103203-2013000010202003-3030233223322200-1312203212233301-1133233103201222-0222203233300312-3210100332133103-2303321111312321): complete subsection reference.

- [k8s_cluster](data-sources--gcp_vpc_site--reference--group-004.md#canonical-3010211030233330-2203020101331110-0232212302122203-0100311223230102-0331113113113020-2101033222222203-1133320113021333-1003131003321311): complete subsection reference.

- [no_dc_cluster_group](data-sources--gcp_vpc_site--reference--group-004.md#canonical-0331331213320330-1030021113201303-3321131320100230-3102322002212313-2331310203330031-1313211211210032-3333001103112233-0023322110030332): complete subsection reference.

- [no_forward_proxy](data-sources--gcp_vpc_site--reference--group-004.md#canonical-1023300033333132-3111232001031012-3110002232131213-1111320011020232-2022032200331210-2013223331130011-0102220132133231-3121320030233121): complete subsection reference.

- [no_global_network](data-sources--gcp_vpc_site--reference--group-004.md#canonical-2121211231121221-0022130123130200-3101033002313113-0111031203013220-1022112230202002-2300333321033013-2001133121303100-3111320300221331): complete subsection reference.

- [no_k8s_cluster](data-sources--gcp_vpc_site--reference--group-004.md#canonical-1300021122101133-3222103223323300-2230003000323301-1332332000001110-2013023301103021-2131301331231302-2001001030302030-0101302213302111): complete subsection reference.

- [no_network_policy](data-sources--gcp_vpc_site--reference--group-004.md#canonical-3300231221023003-3303130230120101-0212230110323330-2012100003321210-0012002101030011-2113122210130011-3022320201002320-1232121200220033): complete subsection reference.

- [no_outside_static_routes](data-sources--gcp_vpc_site--reference--group-004.md#canonical-1300332232132320-0120121023203213-1310021031011123-0130323110333310-0212320211302213-3131302032101122-2020122131013122-0303133222302022): complete subsection reference.

<a id="canonical-3113000302101220-1122020203210231-3103013101232202-2231131200322321-0132322031220132-1012113231130323-0332233120031001-1220311031122013"></a>

<a id="canonical-1110110022011011-3130321101222200-1031330321211201-1022303201222111-1032122030303223-0312331011231232-3211111320322121-2110332031212131"></a>

## node_number property — voltstack_cluster / 230131221201 / 6

Type: `"number"`. Computed.

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

- [outside_static_routes](data-sources--gcp_vpc_site--reference--group-004.md#canonical-3302312100000302-0021313112011031-3212312101300110-3233320120212312-0000122013313331-0312112130203310-2200111320212103-1023131132323330): complete subsection reference.

- [site_local_network](data-sources--gcp_vpc_site--reference--group-004.md#canonical-3111110233303120-3020123200210213-3113011221310021-2032102300121231-0001000301332210-1223222212122033-2103323121303210-2122130003210212): complete subsection reference.

- [site_local_subnet](data-sources--gcp_vpc_site--reference--group-004.md#canonical-0032122331011201-2212323312313011-1322003231020112-2302012332332033-2310121201313332-3003021012113323-3022323300030221-2220332223102330): complete subsection reference.

- [sm_connection_public_ip](data-sources--gcp_vpc_site--reference--group-004.md#canonical-1202113132110030-0331001033001223-1320211110310011-3033232021100220-3110333202300032-3123012232011012-3213130320111011-2032310023331033): complete subsection reference.

- [sm_connection_pvt_ip](data-sources--gcp_vpc_site--reference--group-004.md#canonical-0000101133022021-3020001201030002-1203112303012302-1001300121021122-3203111021332232-1233113103000221-2021020022303003-0332301312102233): complete subsection reference.

- [storage_class_list](data-sources--gcp_vpc_site--reference--group-004.md#canonical-1012101330320212-0113133033002223-0210301022311211-3232000103330311-1233032112203232-2330201133011321-3021003023300101-0302002201010202): complete subsection reference.

<a id="canonical-1230003321330301-3113132210101312-0130113311031310-1322121113022000-1020203122211132-1310022133323320-1101012210231131-1001200123220300"></a>

## Next pages — voltstack_cluster / 230131221201 / 7

- [voltstack_cluster.active_enhanced_firewall_policies](data-sources--gcp_vpc_site--reference--group-004.md#canonical-0232022202333003-2202033311020122-0013013023122033-0301220230222203-3013000111020120-0001131112012210-2131313320222313-0130122022102322)
- [voltstack_cluster.active_forward_proxy_policies](data-sources--gcp_vpc_site--reference--group-004.md#canonical-3000133322203121-1203100320022130-3302113010101322-2022212111130022-0331020321113210-1112313300022331-0123230330213231-3230132100300020)
- [voltstack_cluster.active_network_policies](data-sources--gcp_vpc_site--reference--group-004.md#canonical-1003102311022230-1230310103302103-3131100223313320-0300310013321331-2122022003321011-1302022311023210-0333313301103012-2323222210102333)
- [voltstack_cluster.dc_cluster_group](data-sources--gcp_vpc_site--reference--group-004.md#canonical-2122123121313303-1002012232131000-1200211200300230-1012313010322103-1232230220220102-2233310202223231-0002222311333333-0020003310133132)
- [voltstack_cluster.default_storage](data-sources--gcp_vpc_site--reference--group-004.md#canonical-3120222212131221-0000211132101132-3200332231101123-3312031023220323-2211103012232103-2202133010133203-0332002202213011-1111230220320131)
- [voltstack_cluster.forward_proxy_allow_all](data-sources--gcp_vpc_site--reference--group-004.md#canonical-0211331210122330-0313020312231321-0201212032112212-1001003121101001-0013023310013032-1231310020232312-2222211202331132-2321321233200130)
- [voltstack_cluster.global_network_list](data-sources--gcp_vpc_site--reference--group-004.md#canonical-2312310111103203-2013000010202003-3030233223322200-1312203212233301-1133233103201222-0222203233300312-3210100332133103-2303321111312321)
- [voltstack_cluster.k8s_cluster](data-sources--gcp_vpc_site--reference--group-004.md#canonical-3010211030233330-2203020101331110-0232212302122203-0100311223230102-0331113113113020-2101033222222203-1133320113021333-1003131003321311)
- [voltstack_cluster.no_dc_cluster_group](data-sources--gcp_vpc_site--reference--group-004.md#canonical-0331331213320330-1030021113201303-3321131320100230-3102322002212313-2331310203330031-1313211211210032-3333001103112233-0023322110030332)
- [voltstack_cluster.no_forward_proxy](data-sources--gcp_vpc_site--reference--group-004.md#canonical-1023300033333132-3111232001031012-3110002232131213-1111320011020232-2022032200331210-2013223331130011-0102220132133231-3121320030233121)
- [voltstack_cluster.no_global_network](data-sources--gcp_vpc_site--reference--group-004.md#canonical-2121211231121221-0022130123130200-3101033002313113-0111031203013220-1022112230202002-2300333321033013-2001133121303100-3111320300221331)
- [voltstack_cluster.no_k8s_cluster](data-sources--gcp_vpc_site--reference--group-004.md#canonical-1300021122101133-3222103223323300-2230003000323301-1332332000001110-2013023301103021-2131301331231302-2001001030302030-0101302213302111)
- [voltstack_cluster.no_network_policy](data-sources--gcp_vpc_site--reference--group-004.md#canonical-3300231221023003-3303130230120101-0212230110323330-2012100003321210-0012002101030011-2113122210130011-3022320201002320-1232121200220033)
- [voltstack_cluster.no_outside_static_routes](data-sources--gcp_vpc_site--reference--group-004.md#canonical-1300332232132320-0120121023203213-1310021031011123-0130323110333310-0212320211302213-3131302032101122-2020122131013122-0303133222302022)
- [voltstack_cluster.outside_static_routes](data-sources--gcp_vpc_site--reference--group-004.md#canonical-3302312100000302-0021313112011031-3212312101300110-3233320120212312-0000122013313331-0312112130203310-2200111320212103-1023131132323330)
- [voltstack_cluster.site_local_network](data-sources--gcp_vpc_site--reference--group-004.md#canonical-3111110233303120-3020123200210213-3113011221310021-2032102300121231-0001000301332210-1223222212122033-2103323121303210-2122130003210212)
- [voltstack_cluster.site_local_subnet](data-sources--gcp_vpc_site--reference--group-004.md#canonical-0032122331011201-2212323312313011-1322003231020112-2302012332332033-2310121201313332-3003021012113323-3022323300030221-2220332223102330)
- [voltstack_cluster.sm_connection_public_ip](data-sources--gcp_vpc_site--reference--group-004.md#canonical-1202113132110030-0331001033001223-1320211110310011-3033232021100220-3110333202300032-3123012232011012-3213130320111011-2032310023331033)
- [voltstack_cluster.sm_connection_pvt_ip](data-sources--gcp_vpc_site--reference--group-004.md#canonical-0000101133022021-3020001201030002-1203112303012302-1001300121021122-3203111021332232-1233113103000221-2021020022303003-0332301312102233)
- [voltstack_cluster.storage_class_list](data-sources--gcp_vpc_site--reference--group-004.md#canonical-1012101330320212-0113133033002223-0210301022311211-3232000103330311-1233032112203232-2330201133011321-3021003023300101-0302002201010202)
- [Property reference](data-sources--gcp_vpc_site--reference--group-001.md#canonical-0023231003100101-1010210121112300-0303103132302303-0331000323201313-3003311210331022-3000312312212303-0030202320102022-1232013113103201)
- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-1132222133321101-1232102200201102-2111112102132122-0010023221133123-2333213300031302-1001011022311003-3121103220313010-2103321200303202)

<a id="canonical-0232022202333003-2202033311020122-0013013023122033-0301220230222203-3013000111020120-0001131112012210-2131313320222313-0130122022102322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1003330330111233-2212022211130033-0023220133130013-3223213203120133-1111102303333311-2220022022020011-3003002321223031-3023200003223301"></a>

## voltstack_cluster.active_enhanced_firewall_policies — active_enhanced_firewall_policies / 231233201110 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-1132222133321101-1232102200201102-2111112102132122-0010023221133123-2333213300031302-1001011022311003-3121103220313010-2103321200303202)
- [Property reference](data-sources--gcp_vpc_site--reference--group-001.md#canonical-0023231003100101-1010210121112300-0303103132302303-0331000323201313-3003311210331022-3000312312212303-0030202320102022-1232013113103201)
- [voltstack_cluster](data-sources--gcp_vpc_site--reference--group-004.md#canonical-2213313332313212-1230123320203012-3120121021331113-3210101220221000-1023022133330131-3023333003003323-2311003313332333-2203303111102220)
- voltstack_cluster.active_enhanced_firewall_policies

<a id="canonical-1301310132132100-3120112033223003-2021010133032212-1232003133112021-2302100121300103-2120320013223220-2111310301313110-0130212003130212"></a>

Type: `"single"`. Computed.

List of Enhanced Firewall Policies These policies use session-based rules and provide all OPTIONS
available under firewall policies with an additional option for service insertion.

Upstream description:

List of Enhanced Firewall Policies These policies use session-based rules and provide all OPTIONS
available under firewall policies with an additional option for service insertion.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2302131120222221-2133233023121010-2301120120000312-1330230231122300-2033011310223212-0032203102110130-0023121303032002-3103011333303123"></a>

## Direct properties — active_enhanced_firewall_policies / 231233201110 / 3

- [enhanced_firewall_policies](data-sources--gcp_vpc_site--reference--group-004.md#canonical-3033320230121122-1221231100331222-2201203322023320-3032101322313230-2333233312132222-3121213023030133-0323222201130121-3321321320021131): complete subsection reference.

<a id="canonical-2003232113322300-1013010302322101-1210221222322123-0023133020303031-0012320213200212-1031022123203303-0121311200013232-1213213111213212"></a>

## Next pages — active_enhanced_firewall_policies / 231233201110 / 4

- [voltstack_cluster.active_enhanced_firewall_policies.enhanced_firewall_policies](data-sources--gcp_vpc_site--reference--group-004.md#canonical-3033320230121122-1221231100331222-2201203322023320-3032101322313230-2333233312132222-3121213023030133-0323222201130121-3321321320021131)
- [voltstack_cluster](data-sources--gcp_vpc_site--reference--group-004.md#canonical-2213313332313212-1230123320203012-3120121021331113-3210101220221000-1023022133330131-3023333003003323-2311003313332333-2203303111102220)
- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-1132222133321101-1232102200201102-2111112102132122-0010023221133123-2333213300031302-1001011022311003-3121103220313010-2103321200303202)

<a id="canonical-3033320230121122-1221231100331222-2201203322023320-3032101322313230-2333233312132222-3121213023030133-0323222201130121-3321321320021131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1202211111123300-0111031203201321-1312120130103033-2330031333201201-1211100020023213-2310100310032321-0131311302220132-2222001201031002"></a>

## voltstack_cluster.active_enhanced_firewall_policies.enhanced_firewall_policies — enhanced_firewall_policies / 330200123330 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-1132222133321101-1232102200201102-2111112102132122-0010023221133123-2333213300031302-1001011022311003-3121103220313010-2103321200303202)
- [Property reference](data-sources--gcp_vpc_site--reference--group-001.md#canonical-0023231003100101-1010210121112300-0303103132302303-0331000323201313-3003311210331022-3000312312212303-0030202320102022-1232013113103201)
- [voltstack_cluster](data-sources--gcp_vpc_site--reference--group-004.md#canonical-2213313332313212-1230123320203012-3120121021331113-3210101220221000-1023022133330131-3023333003003323-2311003313332333-2203303111102220)
- [voltstack_cluster.active_enhanced_firewall_policies](data-sources--gcp_vpc_site--reference--group-004.md#canonical-0232022202333003-2202033311020122-0013013023122033-0301220230222203-3013000111020120-0001131112012210-2131313320222313-0130122022102322)
- voltstack_cluster.active_enhanced_firewall_policies.enhanced_firewall_policies

<a id="canonical-2333332221130203-2210321221022230-1312133312201032-0012201031123002-3223313300320233-0122120033331320-1122332113323020-1321222131101011"></a>

Type: `"list"`. Computed.

Ordered List of Enhanced Firewall Policies active.

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

<a id="canonical-1122130323222021-0002221011031032-1020221103323222-1031033021011311-3330303120302221-2211101230021102-1010020201220220-3003111203000020"></a>

## Direct properties — enhanced_firewall_policies / 330200123330 / 3

<a id="canonical-1131133112333322-2331302323202033-3222322111312222-1330222120300023-1310223132103031-0302302330111113-3202210321103123-2330012120113322"></a>

<a id="canonical-3032203101200213-3233121033203023-2333131303013302-2022311103120323-3300200011230210-0012122310121000-1311002031011000-0031121220233020"></a>

## name property — enhanced_firewall_policies / 330200123330 / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

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

<a id="canonical-3023202303130112-2000031300213010-0120330321323132-1100020223032312-3312310022332223-2210202302322323-2100101101001133-1022331320013012"></a>

<a id="canonical-2322030211323110-0023013011032021-3303022130333232-1332100122111221-3310032310103203-3300101001101031-2010332201202030-2123333032333002"></a>

## namespace property — enhanced_firewall_policies / 330200123330 / 5

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

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

<a id="canonical-2331021311101122-0132111120113112-0132223231100123-3230131020130221-2111211113111111-2220120333332120-1230230323022231-3311203112313120"></a>

<a id="canonical-0223102312310031-3020300030121231-1330210011033330-3131003231013223-2111200121230012-3001220200300223-0022123101102003-1222012022013010"></a>

## tenant property — enhanced_firewall_policies / 330200123330 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

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

<a id="canonical-3112130333121032-3120220011221223-3120130122100212-2303323233103032-3120320012131230-2112100223013013-3301202001311021-0321323102113302"></a>

## Next pages — enhanced_firewall_policies / 330200123330 / 7

- [voltstack_cluster.active_enhanced_firewall_policies](data-sources--gcp_vpc_site--reference--group-004.md#canonical-0232022202333003-2202033311020122-0013013023122033-0301220230222203-3013000111020120-0001131112012210-2131313320222313-0130122022102322)
- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-1132222133321101-1232102200201102-2111112102132122-0010023221133123-2333213300031302-1001011022311003-3121103220313010-2103321200303202)

<a id="canonical-3000133322203121-1203100320022130-3302113010101322-2022212111130022-0331020321113210-1112313300022331-0123230330213231-3230132100300020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0330113121001230-3212200032012010-0102001312023323-0300312313211123-3201200311011332-2013022103330103-2111100313003321-3313202320011313"></a>

## voltstack_cluster.active_forward_proxy_policies — active_forward_proxy_policies / 213301010100 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-1132222133321101-1232102200201102-2111112102132122-0010023221133123-2333213300031302-1001011022311003-3121103220313010-2103321200303202)
- [Property reference](data-sources--gcp_vpc_site--reference--group-001.md#canonical-0023231003100101-1010210121112300-0303103132302303-0331000323201313-3003311210331022-3000312312212303-0030202320102022-1232013113103201)
- [voltstack_cluster](data-sources--gcp_vpc_site--reference--group-004.md#canonical-2213313332313212-1230123320203012-3120121021331113-3210101220221000-1023022133330131-3023333003003323-2311003313332333-2203303111102220)
- voltstack_cluster.active_forward_proxy_policies

<a id="canonical-1000212330333001-2333210132332120-0000031021301222-0233333230010220-1100111323210213-2332322013331013-0131223201010213-3303023123220032"></a>

Type: `"single"`. Computed.

Ordered List of Forward Proxy Policies active.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3033303032211220-3211311323332232-3230323320001003-2333032020123111-2230132103231322-2232323102221321-3213200312111211-2033212221203300"></a>

## Direct properties — active_forward_proxy_policies / 213301010100 / 3

- [forward_proxy_policies](data-sources--gcp_vpc_site--reference--group-004.md#canonical-3033020312301221-3132221330132320-3110333000201302-0220130302312101-2032220223100233-0010022231323033-2231220123312001-0211012300202303): complete subsection reference.

<a id="canonical-3213003233120300-0303332210210322-3332312133202101-0121230220222131-1011211033232113-3032301021032030-2323113222200130-3223031231212031"></a>

## Next pages — active_forward_proxy_policies / 213301010100 / 4

- [voltstack_cluster.active_forward_proxy_policies.forward_proxy_policies](data-sources--gcp_vpc_site--reference--group-004.md#canonical-3033020312301221-3132221330132320-3110333000201302-0220130302312101-2032220223100233-0010022231323033-2231220123312001-0211012300202303)
- [voltstack_cluster](data-sources--gcp_vpc_site--reference--group-004.md#canonical-2213313332313212-1230123320203012-3120121021331113-3210101220221000-1023022133330131-3023333003003323-2311003313332333-2203303111102220)
- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-1132222133321101-1232102200201102-2111112102132122-0010023221133123-2333213300031302-1001011022311003-3121103220313010-2103321200303202)

<a id="canonical-3033020312301221-3132221330132320-3110333000201302-0220130302312101-2032220223100233-0010022231323033-2231220123312001-0211012300202303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2022010211033213-1233233002331113-2121230213330212-0132223230303121-1330200021331122-3130113131012301-0113302032103212-1222021210001202"></a>

## voltstack_cluster.active_forward_proxy_policies.forward_proxy_policies — forward_proxy_policies / 300003022330 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-1132222133321101-1232102200201102-2111112102132122-0010023221133123-2333213300031302-1001011022311003-3121103220313010-2103321200303202)
- [Property reference](data-sources--gcp_vpc_site--reference--group-001.md#canonical-0023231003100101-1010210121112300-0303103132302303-0331000323201313-3003311210331022-3000312312212303-0030202320102022-1232013113103201)
- [voltstack_cluster](data-sources--gcp_vpc_site--reference--group-004.md#canonical-2213313332313212-1230123320203012-3120121021331113-3210101220221000-1023022133330131-3023333003003323-2311003313332333-2203303111102220)
- [voltstack_cluster.active_forward_proxy_policies](data-sources--gcp_vpc_site--reference--group-004.md#canonical-3000133322203121-1203100320022130-3302113010101322-2022212111130022-0331020321113210-1112313300022331-0123230330213231-3230132100300020)
- voltstack_cluster.active_forward_proxy_policies.forward_proxy_policies

<a id="canonical-2222200203033223-0121330323113201-2110031330120021-3122313312030303-0301231033011331-1021003010322200-1102013300312210-1121331301210012"></a>

Type: `"list"`. Computed.

Ordered List of Forward Proxy Policies active.

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

<a id="canonical-3130112123113213-1200313102111333-2231230011321101-1233113013311110-2112311100132121-0101211102233123-2203120102213203-2132230323011331"></a>

## Direct properties — forward_proxy_policies / 300003022330 / 3

<a id="canonical-3121333122202223-1001220202013023-2313103301311020-2313323002200120-3321301001100103-3312032331102130-1132130113202133-0130100011033031"></a>

<a id="canonical-1012010032332103-1011301021223200-2200123011320110-3313122113223221-1331031112221112-1130312133102310-0332202013011310-0232211312201033"></a>

## name property — forward_proxy_policies / 300003022330 / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

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

<a id="canonical-3213021113231212-1211031013031311-3022230123003330-2221311323111130-0123213110000322-2121200102202000-1323113331230321-3311320222232130"></a>

<a id="canonical-0212103201102310-3023112023133112-2310111102300123-0121213133312222-0102103003032233-3100131301132320-1213131210033000-0220111201322333"></a>

## namespace property — forward_proxy_policies / 300003022330 / 5

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

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

<a id="canonical-0012310101033132-0021120311023111-2203311310322012-0332221113132301-0003323221101011-1302122123300103-1223322133033302-3212320331102313"></a>

<a id="canonical-2133103300301112-0310302103210000-3220102122111012-2021110201031012-3023110121332121-0123212232311223-1021333100020301-2120120102233203"></a>

## tenant property — forward_proxy_policies / 300003022330 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

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

<a id="canonical-1103012323131302-1301220323323301-2231210033221132-1222031130332201-0221233032321330-0212323333211231-1023312021232302-2333223100121212"></a>

## Next pages — forward_proxy_policies / 300003022330 / 7

- [voltstack_cluster.active_forward_proxy_policies](data-sources--gcp_vpc_site--reference--group-004.md#canonical-3000133322203121-1203100320022130-3302113010101322-2022212111130022-0331020321113210-1112313300022331-0123230330213231-3230132100300020)
- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-1132222133321101-1232102200201102-2111112102132122-0010023221133123-2333213300031302-1001011022311003-3121103220313010-2103321200303202)

<a id="canonical-1003102311022230-1230310103302103-3131100223313320-0300310013321331-2122022003321011-1302022311023210-0333313301103012-2323222210102333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0023020011001121-2321113223201331-0102220103033030-2210101322013010-2302010230001020-0331221321031211-2112003203010313-3212310120211331"></a>

## voltstack_cluster.active_network_policies — active_network_policies / 000232332201 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-1132222133321101-1232102200201102-2111112102132122-0010023221133123-2333213300031302-1001011022311003-3121103220313010-2103321200303202)
- [Property reference](data-sources--gcp_vpc_site--reference--group-001.md#canonical-0023231003100101-1010210121112300-0303103132302303-0331000323201313-3003311210331022-3000312312212303-0030202320102022-1232013113103201)
- [voltstack_cluster](data-sources--gcp_vpc_site--reference--group-004.md#canonical-2213313332313212-1230123320203012-3120121021331113-3210101220221000-1023022133330131-3023333003003323-2311003313332333-2203303111102220)
- voltstack_cluster.active_network_policies

<a id="canonical-1002021111311231-3032230301121310-2010032131113300-3223132213020212-2222331102232232-3331122210212332-1120223203022223-3132312031232300"></a>

Type: `"single"`. Computed.

Configuration parameter for active network policies.

Upstream description:

List of firewall policy views.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2203313221113033-3032020202231012-3022200310132121-2220120221232232-1322202022101021-2300311320200223-2303331120120333-1300020112133112"></a>

## Direct properties — active_network_policies / 000232332201 / 3

- [network_policies](data-sources--gcp_vpc_site--reference--group-004.md#canonical-0311231230100011-2230030332033010-1312202231120022-0222332200030323-3202112033122300-1100031232102033-1321131303202030-1232210131212013): complete subsection reference.

<a id="canonical-2331023311001212-2003222220101231-1101330303202321-2200002321131001-3111120320100130-3003102311200333-0132302232033133-2010211011333021"></a>

## Next pages — active_network_policies / 000232332201 / 4

- [voltstack_cluster.active_network_policies.network_policies](data-sources--gcp_vpc_site--reference--group-004.md#canonical-0311231230100011-2230030332033010-1312202231120022-0222332200030323-3202112033122300-1100031232102033-1321131303202030-1232210131212013)
- [voltstack_cluster](data-sources--gcp_vpc_site--reference--group-004.md#canonical-2213313332313212-1230123320203012-3120121021331113-3210101220221000-1023022133330131-3023333003003323-2311003313332333-2203303111102220)
- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-1132222133321101-1232102200201102-2111112102132122-0010023221133123-2333213300031302-1001011022311003-3121103220313010-2103321200303202)

<a id="canonical-0311231230100011-2230030332033010-1312202231120022-0222332200030323-3202112033122300-1100031232102033-1321131303202030-1232210131212013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3221310110303022-3111101321233330-2033320303212113-2320230300031323-3331333011330312-0320011330130100-2020033321211302-0031302311320011"></a>

## voltstack_cluster.active_network_policies.network_policies — network_policies / 133311011331 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-1132222133321101-1232102200201102-2111112102132122-0010023221133123-2333213300031302-1001011022311003-3121103220313010-2103321200303202)
- [Property reference](data-sources--gcp_vpc_site--reference--group-001.md#canonical-0023231003100101-1010210121112300-0303103132302303-0331000323201313-3003311210331022-3000312312212303-0030202320102022-1232013113103201)
- [voltstack_cluster](data-sources--gcp_vpc_site--reference--group-004.md#canonical-2213313332313212-1230123320203012-3120121021331113-3210101220221000-1023022133330131-3023333003003323-2311003313332333-2203303111102220)
- [voltstack_cluster.active_network_policies](data-sources--gcp_vpc_site--reference--group-004.md#canonical-1003102311022230-1230310103302103-3131100223313320-0300310013321331-2122022003321011-1302022311023210-0333313301103012-2323222210102333)
- voltstack_cluster.active_network_policies.network_policies

<a id="canonical-1002322110130202-3001303120223211-3112321020333103-0100031311010320-0210322031303011-3131133321202022-0210033320202230-0232112031021103"></a>

Type: `"list"`. Computed.

Ordered List of Firewall Policies active for this network firewall.

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

<a id="canonical-0103331331313322-3013200102123203-1000123010000121-3023200323330010-1130002232330010-1130200211200333-0132212313222332-1210132021013113"></a>

## Direct properties — network_policies / 133311011331 / 3

<a id="canonical-1213130331010300-0231312320100112-0031110202212322-2121110120031012-3313101221203310-3332310131313322-3012331221322203-3132233212021011"></a>

<a id="canonical-3301102130031121-3320311120000222-3210002202131022-1110131013213120-2311020221223231-1303213201213222-0313100102001320-2201132132220200"></a>

## name property — network_policies / 133311011331 / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

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

<a id="canonical-1102011223122230-1222011103132012-3022123033220222-1230310230210231-0222202200202311-3201001013132200-3223322113230213-2302203022332310"></a>

<a id="canonical-1303230322132021-1131020000332311-3311320303233331-3233010133032202-2113222112030102-3200120330120311-3032012301100210-0020302123022212"></a>

## namespace property — network_policies / 133311011331 / 5

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

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

<a id="canonical-2320112310120030-0201033222023213-0011221230120110-0130310030121212-2132312013112022-3212301032213000-2222233001331123-1221003011310233"></a>

<a id="canonical-0130233231020333-1022301130333023-0022331332103321-3013202332332330-1302122130202010-2213211110303023-3103102220223110-1212210211330331"></a>

## tenant property — network_policies / 133311011331 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

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

<a id="canonical-0332023323221232-0001330122131101-1132031113201311-3123003322311031-3103310213011330-2213310300331203-3030310132311330-2333002220133003"></a>

## Next pages — network_policies / 133311011331 / 7

- [voltstack_cluster.active_network_policies](data-sources--gcp_vpc_site--reference--group-004.md#canonical-1003102311022230-1230310103302103-3131100223313320-0300310013321331-2122022003321011-1302022311023210-0333313301103012-2323222210102333)
- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-1132222133321101-1232102200201102-2111112102132122-0010023221133123-2333213300031302-1001011022311003-3121103220313010-2103321200303202)

<a id="canonical-2122123121313303-1002012232131000-1200211200300230-1012313010322103-1232230220220102-2233310202223231-0002222311333333-0020003310133132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2000320011031001-0323330322113300-1031133333300323-3032323001023111-2123010213230311-0201133130012312-0222012111033011-1130312211010232"></a>

## voltstack_cluster.dc_cluster_group — dc_cluster_group / 313032112210 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-1132222133321101-1232102200201102-2111112102132122-0010023221133123-2333213300031302-1001011022311003-3121103220313010-2103321200303202)
- [Property reference](data-sources--gcp_vpc_site--reference--group-001.md#canonical-0023231003100101-1010210121112300-0303103132302303-0331000323201313-3003311210331022-3000312312212303-0030202320102022-1232013113103201)
- [voltstack_cluster](data-sources--gcp_vpc_site--reference--group-004.md#canonical-2213313332313212-1230123320203012-3120121021331113-3210101220221000-1023022133330131-3023333003003323-2311003313332333-2203303111102220)
- voltstack_cluster.dc_cluster_group

<a id="canonical-2010012001130330-3232102120221230-0030211202311023-0233202030030312-3023210133022123-0030133300012301-3130331233020110-2211132210031022"></a>

Type: `"single"`. Computed.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

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

<a id="canonical-1331111010012302-2011332231223310-2211120332300123-1320220012003232-2022211332023301-3023312003211133-3302323212010230-0310330003011301"></a>

## Direct properties — dc_cluster_group / 313032112210 / 3

<a id="canonical-1103322322121321-2321013010223003-2013310203213021-3123110210032121-1200130031020212-0300331302300323-0233231033011022-3230323132132033"></a>

<a id="canonical-3031113011020012-0321202301020122-1313320101210212-2033300022032213-3021132100001102-0033331231033203-2101300032230000-3222231100301203"></a>

## name property — dc_cluster_group / 313032112210 / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

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

<a id="canonical-2300201000012210-2031001303011202-2312020110301023-2221102223301202-2321003310301133-3002030331011033-2030100312121122-1200110032221230"></a>

<a id="canonical-2012100202332213-2323032113132202-2132323012211320-2103131312112000-0023300133223211-0001012311021103-1020112232320230-1313300100210200"></a>

## namespace property — dc_cluster_group / 313032112210 / 5

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

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

<a id="canonical-0033133132321230-2203121222212002-2232203223220312-3102022331110303-1113211221003121-0211030102222123-1333310202121233-3101230330301111"></a>

<a id="canonical-0233320232101211-2111002100222311-3332313311002312-2020311001020302-3012331021131111-3033010330302220-0213111231032321-3010113303102000"></a>

## tenant property — dc_cluster_group / 313032112210 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

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

<a id="canonical-2223202120302231-1233023113012221-2122333330112011-1132222113012030-0220203121001002-3322320230200321-2333022110330003-1220322311300023"></a>

## Next pages — dc_cluster_group / 313032112210 / 7

- [voltstack_cluster](data-sources--gcp_vpc_site--reference--group-004.md#canonical-2213313332313212-1230123320203012-3120121021331113-3210101220221000-1023022133330131-3023333003003323-2311003313332333-2203303111102220)
- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-1132222133321101-1232102200201102-2111112102132122-0010023221133123-2333213300031302-1001011022311003-3121103220313010-2103321200303202)

<a id="canonical-3120222212131221-0000211132101132-3200332231101123-3312031023220323-2211103012232103-2202133010133203-0332002202213011-1111230220320131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3020320110031100-1132133200101302-1210230201230030-1332123111000213-1223012220200333-1220003212101313-0310132202332332-0133323321030012"></a>

## voltstack_cluster.default_storage — default_storage / 212310213321 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-1132222133321101-1232102200201102-2111112102132122-0010023221133123-2333213300031302-1001011022311003-3121103220313010-2103321200303202)
- [Property reference](data-sources--gcp_vpc_site--reference--group-001.md#canonical-0023231003100101-1010210121112300-0303103132302303-0331000323201313-3003311210331022-3000312312212303-0030202320102022-1232013113103201)
- [voltstack_cluster](data-sources--gcp_vpc_site--reference--group-004.md#canonical-2213313332313212-1230123320203012-3120121021331113-3210101220221000-1023022133330131-3023333003003323-2311003313332333-2203303111102220)
- voltstack_cluster.default_storage

<a id="canonical-0302310021333323-1321002032210330-1200130211001000-3331230013223013-2310332323102303-0312222201112220-0121100232032101-2313331023033000"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for default storage.

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

<a id="canonical-1120002022330020-0010212130130302-0101030222032220-3313002100011312-3133232013002210-2130310121232233-0231320320311120-2200332322323022"></a>

## Direct properties — default_storage / 212310213321 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1231211303032003-0011230233312333-0122213332222201-2011312203220300-1102220331321002-1221321022322131-2112211213333330-3312100210130130"></a>

## Next pages — default_storage / 212310213321 / 4

- [voltstack_cluster](data-sources--gcp_vpc_site--reference--group-004.md#canonical-2213313332313212-1230123320203012-3120121021331113-3210101220221000-1023022133330131-3023333003003323-2311003313332333-2203303111102220)
- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-1132222133321101-1232102200201102-2111112102132122-0010023221133123-2333213300031302-1001011022311003-3121103220313010-2103321200303202)

<a id="canonical-0211331210122330-0313020312231321-0201212032112212-1001003121101001-0013023310013032-1231310020232312-2222211202331132-2321321233200130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0230320312333233-3022102332211200-3011122003232001-1013011222011020-2310012232021323-2031021021220310-3010200130100003-3023001033312333"></a>

## voltstack_cluster.forward_proxy_allow_all — forward_proxy_allow_all / 202122001022 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-1132222133321101-1232102200201102-2111112102132122-0010023221133123-2333213300031302-1001011022311003-3121103220313010-2103321200303202)
- [Property reference](data-sources--gcp_vpc_site--reference--group-001.md#canonical-0023231003100101-1010210121112300-0303103132302303-0331000323201313-3003311210331022-3000312312212303-0030202320102022-1232013113103201)
- [voltstack_cluster](data-sources--gcp_vpc_site--reference--group-004.md#canonical-2213313332313212-1230123320203012-3120121021331113-3210101220221000-1023022133330131-3023333003003323-2311003313332333-2203303111102220)
- voltstack_cluster.forward_proxy_allow_all

<a id="canonical-3300230133233231-0012011302210331-1112131321010002-3030012112122320-1312121023023311-2221132212111132-1102312312321023-1101302311302100"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-1210000113002010-2033212133331232-2231230022210112-2303210000102230-1303022201102021-3303212033332101-3021103030030200-2333302222001113"></a>

## Direct properties — forward_proxy_allow_all / 202122001022 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3001302012211010-3220001230300023-3010333312210331-1003100102310000-2111130031330023-2312202331132013-1311100332030210-0011122021300213"></a>

## Next pages — forward_proxy_allow_all / 202122001022 / 4

- [voltstack_cluster](data-sources--gcp_vpc_site--reference--group-004.md#canonical-2213313332313212-1230123320203012-3120121021331113-3210101220221000-1023022133330131-3023333003003323-2311003313332333-2203303111102220)
- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-1132222133321101-1232102200201102-2111112102132122-0010023221133123-2333213300031302-1001011022311003-3121103220313010-2103321200303202)

<a id="canonical-2312310111103203-2013000010202003-3030233223322200-1312203212233301-1133233103201222-0222203233300312-3210100332133103-2303321111312321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3111110213312203-2003230123212131-2312323312311001-3213120000132032-2320022122220130-0120212331323031-1303102321133210-3001112202013111"></a>

## voltstack_cluster.global_network_list — global_network_list / 201320000132 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-1132222133321101-1232102200201102-2111112102132122-0010023221133123-2333213300031302-1001011022311003-3121103220313010-2103321200303202)
- [Property reference](data-sources--gcp_vpc_site--reference--group-001.md#canonical-0023231003100101-1010210121112300-0303103132302303-0331000323201313-3003311210331022-3000312312212303-0030202320102022-1232013113103201)
- [voltstack_cluster](data-sources--gcp_vpc_site--reference--group-004.md#canonical-2213313332313212-1230123320203012-3120121021331113-3210101220221000-1023022133330131-3023333003003323-2311003313332333-2203303111102220)
- voltstack_cluster.global_network_list

<a id="canonical-2323211020333101-2331101232201103-3031231202233221-0101202333110323-3311333230212303-3230032001131321-1011021300201231-1220001323300121"></a>

Type: `"single"`. Computed.

Global Network Connection List. List of global network connections.

Upstream description:

List of global network connections.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1100211201211010-2122322323020203-1023311002303320-0322331011103211-1320000210220013-3032322103303333-0200221002313020-1030101011323021"></a>

## Direct properties — global_network_list / 201320000132 / 3

- [global_network_connections](data-sources--gcp_vpc_site--reference--group-004.md#canonical-2112120020213200-3003013232201200-3030221111023103-0322021023311220-2220232100022331-0110011123113200-1312032023131322-1123012033223022): complete subsection reference.

<a id="canonical-1321012231020103-3223330231310222-1132302321100220-3031123003232100-1003113003202320-1311033113001223-3132123313133123-2320233323101011"></a>

## Next pages — global_network_list / 201320000132 / 4

- [voltstack_cluster.global_network_list.global_network_connections](data-sources--gcp_vpc_site--reference--group-004.md#canonical-2112120020213200-3003013232201200-3030221111023103-0322021023311220-2220232100022331-0110011123113200-1312032023131322-1123012033223022)
- [voltstack_cluster](data-sources--gcp_vpc_site--reference--group-004.md#canonical-2213313332313212-1230123320203012-3120121021331113-3210101220221000-1023022133330131-3023333003003323-2311003313332333-2203303111102220)
- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-1132222133321101-1232102200201102-2111112102132122-0010023221133123-2333213300031302-1001011022311003-3121103220313010-2103321200303202)

<a id="canonical-2112120020213200-3003013232201200-3030221111023103-0322021023311220-2220232100022331-0110011123113200-1312032023131322-1123012033223022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0031203033301113-3212322231020011-0320113002100303-0211020320332313-2332003103312210-0211112220230101-3203112021311211-0023232001131332"></a>

## voltstack_cluster.global_network_list.global_network_connections — global_network_connections / 213301131213 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-1132222133321101-1232102200201102-2111112102132122-0010023221133123-2333213300031302-1001011022311003-3121103220313010-2103321200303202)
- [Property reference](data-sources--gcp_vpc_site--reference--group-001.md#canonical-0023231003100101-1010210121112300-0303103132302303-0331000323201313-3003311210331022-3000312312212303-0030202320102022-1232013113103201)
- [voltstack_cluster](data-sources--gcp_vpc_site--reference--group-004.md#canonical-2213313332313212-1230123320203012-3120121021331113-3210101220221000-1023022133330131-3023333003003323-2311003313332333-2203303111102220)
- [voltstack_cluster.global_network_list](data-sources--gcp_vpc_site--reference--group-004.md#canonical-2312310111103203-2013000010202003-3030233223322200-1312203212233301-1133233103201222-0222203233300312-3210100332133103-2303321111312321)
- voltstack_cluster.global_network_list.global_network_connections

<a id="canonical-2323021200122121-1112123121330131-2013232220021033-2313002100221012-1102002000122131-0231303203012030-1021012302212321-1122222002102013"></a>

Type: `"list"`. Computed.

Global Network Connections. Global network connections.

Upstream description:

Global network connections.

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

<a id="canonical-3113332120302231-1121111230122033-1321202102130033-2211010133003133-2020103230321200-2022013333331223-2331230320221321-2000232220201113"></a>

## Direct properties — global_network_connections / 213301131213 / 3

- [sli_to_global_dr](data-sources--gcp_vpc_site--reference--group-004.md#canonical-0320302023022210-1113300310203010-1132100010210222-2300220022131032-3030313101201100-1111030331000222-3330002332231221-3321303001312210): complete subsection reference.

- [slo_to_global_dr](data-sources--gcp_vpc_site--reference--group-004.md#canonical-0220230123010331-1301303122031231-2131030001231213-1310301120231022-0330132202021110-0331113202232100-1223220211021321-0303110101223020): complete subsection reference.

<a id="canonical-3111103301210202-2110322101123102-3231001110303233-3133011011002110-0123310001211130-2030210011301102-2231232021012233-1131331031211000"></a>

## Next pages — global_network_connections / 213301131213 / 4

- [voltstack_cluster.global_network_list.global_network_connections.sli_to_global_dr](data-sources--gcp_vpc_site--reference--group-004.md#canonical-0320302023022210-1113300310203010-1132100010210222-2300220022131032-3030313101201100-1111030331000222-3330002332231221-3321303001312210)
- [voltstack_cluster.global_network_list.global_network_connections.slo_to_global_dr](data-sources--gcp_vpc_site--reference--group-004.md#canonical-0220230123010331-1301303122031231-2131030001231213-1310301120231022-0330132202021110-0331113202232100-1223220211021321-0303110101223020)
- [voltstack_cluster.global_network_list](data-sources--gcp_vpc_site--reference--group-004.md#canonical-2312310111103203-2013000010202003-3030233223322200-1312203212233301-1133233103201222-0222203233300312-3210100332133103-2303321111312321)
- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-1132222133321101-1232102200201102-2111112102132122-0010023221133123-2333213300031302-1001011022311003-3121103220313010-2103321200303202)

<a id="canonical-0320302023022210-1113300310203010-1132100010210222-2300220022131032-3030313101201100-1111030331000222-3330002332231221-3321303001312210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0323320302300131-2201021123110331-0030210321021320-2303030311233033-3220331022220101-0001323031322331-1011230322011110-2212102020121203"></a>

## voltstack_cluster.global_network_list.global_network_connections.sli_to_global_dr — sli_to_global_dr / 121110121303 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-1132222133321101-1232102200201102-2111112102132122-0010023221133123-2333213300031302-1001011022311003-3121103220313010-2103321200303202)
- [Property reference](data-sources--gcp_vpc_site--reference--group-001.md#canonical-0023231003100101-1010210121112300-0303103132302303-0331000323201313-3003311210331022-3000312312212303-0030202320102022-1232013113103201)
- [voltstack_cluster](data-sources--gcp_vpc_site--reference--group-004.md#canonical-2213313332313212-1230123320203012-3120121021331113-3210101220221000-1023022133330131-3023333003003323-2311003313332333-2203303111102220)
- [voltstack_cluster.global_network_list](data-sources--gcp_vpc_site--reference--group-004.md#canonical-2312310111103203-2013000010202003-3030233223322200-1312203212233301-1133233103201222-0222203233300312-3210100332133103-2303321111312321)
- [voltstack_cluster.global_network_list.global_network_connections](data-sources--gcp_vpc_site--reference--group-004.md#canonical-2112120020213200-3003013232201200-3030221111023103-0322021023311220-2220232100022331-0110011123113200-1312032023131322-1123012033223022)
- voltstack_cluster.global_network_list.global_network_connections.sli_to_global_dr

<a id="canonical-3011212102300232-3332222111311222-1331033131013103-3200000321310322-3013111221102012-2321011121210220-1020331211311320-3311201230102100"></a>

Type: `"single"`. Computed.

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

<a id="canonical-0213232303220201-0323131310031101-2210002003011202-1012032302023102-2012103032323221-2021131121003001-3212210030230232-0211013233121201"></a>

## Direct properties — sli_to_global_dr / 121110121303 / 3

- [global_vn](data-sources--gcp_vpc_site--reference--group-004.md#canonical-2202301013021111-0201000202321032-1330130103312012-2222001303010111-1220303113311213-1102220013003230-2212322310223300-2101033212321100): complete subsection reference.

<a id="canonical-1300303132100023-1331103313303011-3300110330312301-1133332220131211-3333132200232213-1011331312323122-0120322201131100-1203022131223132"></a>

## Next pages — sli_to_global_dr / 121110121303 / 4

- [voltstack_cluster.global_network_list.global_network_connections.sli_to_global_dr.global_vn](data-sources--gcp_vpc_site--reference--group-004.md#canonical-2202301013021111-0201000202321032-1330130103312012-2222001303010111-1220303113311213-1102220013003230-2212322310223300-2101033212321100)
- [voltstack_cluster.global_network_list.global_network_connections](data-sources--gcp_vpc_site--reference--group-004.md#canonical-2112120020213200-3003013232201200-3030221111023103-0322021023311220-2220232100022331-0110011123113200-1312032023131322-1123012033223022)
- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-1132222133321101-1232102200201102-2111112102132122-0010023221133123-2333213300031302-1001011022311003-3121103220313010-2103321200303202)

<a id="canonical-2202301013021111-0201000202321032-1330130103312012-2222001303010111-1220303113311213-1102220013003230-2212322310223300-2101033212321100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1103003210313123-1311210220121300-3330003212231110-2111331210130222-1323111322230303-3220120320333021-2223221231021010-3131113332211122"></a>

## voltstack_cluster.global_network_list.global_network_connections.sli_to_global_dr.global_vn — global_vn / 221201112320 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-1132222133321101-1232102200201102-2111112102132122-0010023221133123-2333213300031302-1001011022311003-3121103220313010-2103321200303202)
- [Property reference](data-sources--gcp_vpc_site--reference--group-001.md#canonical-0023231003100101-1010210121112300-0303103132302303-0331000323201313-3003311210331022-3000312312212303-0030202320102022-1232013113103201)
- [voltstack_cluster](data-sources--gcp_vpc_site--reference--group-004.md#canonical-2213313332313212-1230123320203012-3120121021331113-3210101220221000-1023022133330131-3023333003003323-2311003313332333-2203303111102220)
- [voltstack_cluster.global_network_list](data-sources--gcp_vpc_site--reference--group-004.md#canonical-2312310111103203-2013000010202003-3030233223322200-1312203212233301-1133233103201222-0222203233300312-3210100332133103-2303321111312321)
- [voltstack_cluster.global_network_list.global_network_connections](data-sources--gcp_vpc_site--reference--group-004.md#canonical-2112120020213200-3003013232201200-3030221111023103-0322021023311220-2220232100022331-0110011123113200-1312032023131322-1123012033223022)
- [voltstack_cluster.global_network_list.global_network_connections.sli_to_global_dr](data-sources--gcp_vpc_site--reference--group-004.md#canonical-0320302023022210-1113300310203010-1132100010210222-2300220022131032-3030313101201100-1111030331000222-3330002332231221-3321303001312210)
- voltstack_cluster.global_network_list.global_network_connections.sli_to_global_dr.global_vn

<a id="canonical-2123123122031302-0210231213100132-3122300312232303-1111111132220001-2032021112111333-0101000111103233-3331322013001211-0210320312001232"></a>

Type: `"single"`. Computed.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

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

<a id="canonical-1103202111313231-2102203231103000-1120322330220102-2301233020331221-3120003212133100-2211120101111001-0321023001020201-3120310201103200"></a>

## Direct properties — global_vn / 221201112320 / 3

<a id="canonical-0113322312101302-2121111033321030-2300201101220322-2000010310303000-2033132223210322-3133030232012332-1021332110131331-2311030200113002"></a>

<a id="canonical-3030231103102212-0101023010033301-3202330200023112-1220301312121013-2000120311323312-3113121022322101-1231300230330322-2211002321102330"></a>

## name property — global_vn / 221201112320 / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

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

<a id="canonical-2331320022022111-3231200120210022-1200313121021303-3032113013303313-2011311221003120-0203030002011121-1201311131201031-3300310223033033"></a>

<a id="canonical-0333302131233003-0223303001312001-1130332203020003-1121320022111202-1203122003113101-2303320111112221-3301311100110223-3032331112222213"></a>

## namespace property — global_vn / 221201112320 / 5

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

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

<a id="canonical-1200033210122333-1312220100201031-0132132321201131-1200100300231200-0200123110301311-0110323002331322-0231202023123202-2011210230320003"></a>

<a id="canonical-1320230333113333-3233003102002000-3202210302220030-0321231012232201-1020000312211022-2131022030331210-2200000300111121-2031301011300002"></a>

## tenant property — global_vn / 221201112320 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

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

<a id="canonical-0210033220312121-2230230213131022-1020031130102200-2003122132210033-2030030130323130-3311011331322303-0011300113210210-2130133020132303"></a>

## Next pages — global_vn / 221201112320 / 7

- [voltstack_cluster.global_network_list.global_network_connections.sli_to_global_dr](data-sources--gcp_vpc_site--reference--group-004.md#canonical-0320302023022210-1113300310203010-1132100010210222-2300220022131032-3030313101201100-1111030331000222-3330002332231221-3321303001312210)
- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-1132222133321101-1232102200201102-2111112102132122-0010023221133123-2333213300031302-1001011022311003-3121103220313010-2103321200303202)

<a id="canonical-0220230123010331-1301303122031231-2131030001231213-1310301120231022-0330132202021110-0331113202232100-1223220211021321-0303110101223020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2203032232010013-2301132023132032-0020311120221122-2023100032232032-2131211010031213-3112011222220031-3223212213200002-0311223122012322"></a>

## voltstack_cluster.global_network_list.global_network_connections.slo_to_global_dr — slo_to_global_dr / 320133312032 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-1132222133321101-1232102200201102-2111112102132122-0010023221133123-2333213300031302-1001011022311003-3121103220313010-2103321200303202)
- [Property reference](data-sources--gcp_vpc_site--reference--group-001.md#canonical-0023231003100101-1010210121112300-0303103132302303-0331000323201313-3003311210331022-3000312312212303-0030202320102022-1232013113103201)
- [voltstack_cluster](data-sources--gcp_vpc_site--reference--group-004.md#canonical-2213313332313212-1230123320203012-3120121021331113-3210101220221000-1023022133330131-3023333003003323-2311003313332333-2203303111102220)
- [voltstack_cluster.global_network_list](data-sources--gcp_vpc_site--reference--group-004.md#canonical-2312310111103203-2013000010202003-3030233223322200-1312203212233301-1133233103201222-0222203233300312-3210100332133103-2303321111312321)
- [voltstack_cluster.global_network_list.global_network_connections](data-sources--gcp_vpc_site--reference--group-004.md#canonical-2112120020213200-3003013232201200-3030221111023103-0322021023311220-2220232100022331-0110011123113200-1312032023131322-1123012033223022)
- voltstack_cluster.global_network_list.global_network_connections.slo_to_global_dr

<a id="canonical-2333012221331233-1320331201222321-2022130033110123-0213202102333102-2103001033213133-0210222023133031-0132100233300233-0322121003021103"></a>

Type: `"single"`. Computed.

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

<a id="canonical-0301103121111320-1313110101031231-0221301203313332-0232330010311201-1331013120113001-0022213032120011-2132112200101320-0221323223312321"></a>

## Direct properties — slo_to_global_dr / 320133312032 / 3

- [global_vn](data-sources--gcp_vpc_site--reference--group-004.md#canonical-0131033030022133-0330330323102310-3320000212032102-1203011211212020-1112131230122003-2102110020200212-1221023021013233-2221200212022300): complete subsection reference.

<a id="canonical-3321112123222121-1203302020330232-3320111233303323-1301033200111232-1000201230323220-3123320110321211-0120120033010311-2200020013232132"></a>

## Next pages — slo_to_global_dr / 320133312032 / 4

- [voltstack_cluster.global_network_list.global_network_connections.slo_to_global_dr.global_vn](data-sources--gcp_vpc_site--reference--group-004.md#canonical-0131033030022133-0330330323102310-3320000212032102-1203011211212020-1112131230122003-2102110020200212-1221023021013233-2221200212022300)
- [voltstack_cluster.global_network_list.global_network_connections](data-sources--gcp_vpc_site--reference--group-004.md#canonical-2112120020213200-3003013232201200-3030221111023103-0322021023311220-2220232100022331-0110011123113200-1312032023131322-1123012033223022)
- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-1132222133321101-1232102200201102-2111112102132122-0010023221133123-2333213300031302-1001011022311003-3121103220313010-2103321200303202)

<a id="canonical-0131033030022133-0330330323102310-3320000212032102-1203011211212020-1112131230122003-2102110020200212-1221023021013233-2221200212022300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0120020003213132-2332312310213120-3003010010311033-2213003321120110-1213101133213011-2123310001202030-0221321211322001-1132110311012203"></a>

## voltstack_cluster.global_network_list.global_network_connections.slo_to_global_dr.global_vn — global_vn / 033100020110 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-1132222133321101-1232102200201102-2111112102132122-0010023221133123-2333213300031302-1001011022311003-3121103220313010-2103321200303202)
- [Property reference](data-sources--gcp_vpc_site--reference--group-001.md#canonical-0023231003100101-1010210121112300-0303103132302303-0331000323201313-3003311210331022-3000312312212303-0030202320102022-1232013113103201)
- [voltstack_cluster](data-sources--gcp_vpc_site--reference--group-004.md#canonical-2213313332313212-1230123320203012-3120121021331113-3210101220221000-1023022133330131-3023333003003323-2311003313332333-2203303111102220)
- [voltstack_cluster.global_network_list](data-sources--gcp_vpc_site--reference--group-004.md#canonical-2312310111103203-2013000010202003-3030233223322200-1312203212233301-1133233103201222-0222203233300312-3210100332133103-2303321111312321)
- [voltstack_cluster.global_network_list.global_network_connections](data-sources--gcp_vpc_site--reference--group-004.md#canonical-2112120020213200-3003013232201200-3030221111023103-0322021023311220-2220232100022331-0110011123113200-1312032023131322-1123012033223022)
- [voltstack_cluster.global_network_list.global_network_connections.slo_to_global_dr](data-sources--gcp_vpc_site--reference--group-004.md#canonical-0220230123010331-1301303122031231-2131030001231213-1310301120231022-0330132202021110-0331113202232100-1223220211021321-0303110101223020)
- voltstack_cluster.global_network_list.global_network_connections.slo_to_global_dr.global_vn

<a id="canonical-2102210221020223-1230200023033011-2322212130021020-0021311002323302-2111203300012233-2302333210123330-0100201201001132-2302233013230313"></a>

Type: `"single"`. Computed.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

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

<a id="canonical-3303132103222023-1033211033311012-3021223030223320-2101313131330123-3131201020300120-0322320212213121-2230001330323222-0032113013203021"></a>

## Direct properties — global_vn / 033100020110 / 3

<a id="canonical-2220221210120101-3321321212333001-1330232222031231-1213323023130131-3223110032202132-2002001321201122-2010311000330201-3100303012132130"></a>

<a id="canonical-0111330301333311-0220122021322311-0031021110313220-1033313001110000-0303123221211110-1033123031121320-1031232230112022-3213033221331303"></a>

## name property — global_vn / 033100020110 / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

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

<a id="canonical-3101122013110203-2220002102000203-2031331123003300-0220310120003003-1311013303213101-1310013032323232-3110003020223010-1021110203111111"></a>

<a id="canonical-0131330103322302-3030210120022023-3320310022111313-0203322332321113-2010113033103302-3331233312200330-2113320321130120-0012321301101213"></a>

## namespace property — global_vn / 033100020110 / 5

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

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

<a id="canonical-0012123011330122-2103020030100120-3211202012123123-2121313023032121-2130022312120033-3303112110311212-2322230222010311-1330131210011312"></a>

<a id="canonical-0113032233123131-0233330023101012-2210010210221003-2033003210220031-1320001212331311-2333133121221332-1033330131112100-3222231123101122"></a>

## tenant property — global_vn / 033100020110 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

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

<a id="canonical-3203321232012032-0300132232311021-0203102003133310-1132122112133013-3000230021103221-2312022301211220-1123122200130101-2022321011103010"></a>

## Next pages — global_vn / 033100020110 / 7

- [voltstack_cluster.global_network_list.global_network_connections.slo_to_global_dr](data-sources--gcp_vpc_site--reference--group-004.md#canonical-0220230123010331-1301303122031231-2131030001231213-1310301120231022-0330132202021110-0331113202232100-1223220211021321-0303110101223020)
- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-1132222133321101-1232102200201102-2111112102132122-0010023221133123-2333213300031302-1001011022311003-3121103220313010-2103321200303202)

<a id="canonical-3010211030233330-2203020101331110-0232212302122203-0100311223230102-0331113113113020-2101033222222203-1133320113021333-1003131003321311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3000003102200100-1101231222200001-0232201130333303-3332100333202332-2010210330002201-3122221210131011-1213312321031210-0013122120211022"></a>

## voltstack_cluster.k8s_cluster — k8s_cluster / 323103101201 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-1132222133321101-1232102200201102-2111112102132122-0010023221133123-2333213300031302-1001011022311003-3121103220313010-2103321200303202)
- [Property reference](data-sources--gcp_vpc_site--reference--group-001.md#canonical-0023231003100101-1010210121112300-0303103132302303-0331000323201313-3003311210331022-3000312312212303-0030202320102022-1232013113103201)
- [voltstack_cluster](data-sources--gcp_vpc_site--reference--group-004.md#canonical-2213313332313212-1230123320203012-3120121021331113-3210101220221000-1023022133330131-3023333003003323-2311003313332333-2203303111102220)
- voltstack_cluster.k8s_cluster

<a id="canonical-3012333312002112-3332331132132332-1200012230321313-0300211130331212-2232312223111312-0320101302210220-2113233222300301-0222233123133312"></a>

Type: `"single"`. Computed.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

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

<a id="canonical-3331202031310202-2213320031320332-0031032203320132-0021221022313131-0102011200031031-3023213112130031-1103313202223012-3020021222323010"></a>

## Direct properties — k8s_cluster / 323103101201 / 3

<a id="canonical-1231001122333220-2130000001033320-2202100011303030-2032333323130121-0010232322200203-3013233020320013-3303013202222200-3203111210110123"></a>

<a id="canonical-1030312032133111-3130313331230231-3132002332330221-2330212031113113-1322210300232230-2202110212311002-0000232212131100-2300102311332033"></a>

## name property — k8s_cluster / 323103101201 / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

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

<a id="canonical-2322001012002123-0210221031113133-3002222113223311-3323323012100221-2000010113300131-2200210321121122-3330301013000021-0023021210120133"></a>

<a id="canonical-2213022201113101-1232310033233303-3101233000020032-1110300321111231-2220312032000210-2012333123220202-0112213000003100-1310002113113103"></a>

## namespace property — k8s_cluster / 323103101201 / 5

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

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

<a id="canonical-0110323020212110-0011103230330200-0021331211131232-3131120300232313-1301012000122300-1023303111122321-0310123021031230-2332211203001030"></a>

<a id="canonical-1103213211022200-1222011233001222-1000032011331211-0220101223330213-1100231110010012-3031310331133212-0301000123211200-3212221231222213"></a>

## tenant property — k8s_cluster / 323103101201 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

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

<a id="canonical-1033113003033013-0101111222233220-1303230130320301-1333233313212002-1311332131101013-3032023021301230-1030311220213113-0012122030301013"></a>

## Next pages — k8s_cluster / 323103101201 / 7

- [voltstack_cluster](data-sources--gcp_vpc_site--reference--group-004.md#canonical-2213313332313212-1230123320203012-3120121021331113-3210101220221000-1023022133330131-3023333003003323-2311003313332333-2203303111102220)
- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-1132222133321101-1232102200201102-2111112102132122-0010023221133123-2333213300031302-1001011022311003-3121103220313010-2103321200303202)

<a id="canonical-0331331213320330-1030021113201303-3321131320100230-3102322002212313-2331310203330031-1313211211210032-3333001103112233-0023322110030332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1213012321030102-0331032010123233-3221001232302121-3133010322100323-3131332012213213-3132313333331211-3231032310100110-3100002020300012"></a>

## voltstack_cluster.no_dc_cluster_group — no_dc_cluster_group / 000023320313 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-1132222133321101-1232102200201102-2111112102132122-0010023221133123-2333213300031302-1001011022311003-3121103220313010-2103321200303202)
- [Property reference](data-sources--gcp_vpc_site--reference--group-001.md#canonical-0023231003100101-1010210121112300-0303103132302303-0331000323201313-3003311210331022-3000312312212303-0030202320102022-1232013113103201)
- [voltstack_cluster](data-sources--gcp_vpc_site--reference--group-004.md#canonical-2213313332313212-1230123320203012-3120121021331113-3210101220221000-1023022133330131-3023333003003323-2311003313332333-2203303111102220)
- voltstack_cluster.no_dc_cluster_group

<a id="canonical-2331111121102012-1200111113312120-0133223021033211-2121231011121022-3123313001033000-2000300322303101-3312131230231200-0201031103102210"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-1300330203120211-2110123123132313-2313111100000013-1013302202323213-3220133003312223-2020003030023121-2300100020002333-3100110301231232"></a>

## Direct properties — no_dc_cluster_group / 000023320313 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3300322222300320-3233223021303111-1010102303200300-1231020113122113-1300313211303023-3110213122130032-2312202002211202-1202322131033103"></a>

## Next pages — no_dc_cluster_group / 000023320313 / 4

- [voltstack_cluster](data-sources--gcp_vpc_site--reference--group-004.md#canonical-2213313332313212-1230123320203012-3120121021331113-3210101220221000-1023022133330131-3023333003003323-2311003313332333-2203303111102220)
- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-1132222133321101-1232102200201102-2111112102132122-0010023221133123-2333213300031302-1001011022311003-3121103220313010-2103321200303202)

<a id="canonical-1023300033333132-3111232001031012-3110002232131213-1111320011020232-2022032200331210-2013223331130011-0102220132133231-3121320030233121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3031023122232031-0001030333133030-1021001311231312-0213222323223001-0222312131313131-0020133300000133-3110112303211022-2023210323022120"></a>

## voltstack_cluster.no_forward_proxy — no_forward_proxy / 311331112120 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-1132222133321101-1232102200201102-2111112102132122-0010023221133123-2333213300031302-1001011022311003-3121103220313010-2103321200303202)
- [Property reference](data-sources--gcp_vpc_site--reference--group-001.md#canonical-0023231003100101-1010210121112300-0303103132302303-0331000323201313-3003311210331022-3000312312212303-0030202320102022-1232013113103201)
- [voltstack_cluster](data-sources--gcp_vpc_site--reference--group-004.md#canonical-2213313332313212-1230123320203012-3120121021331113-3210101220221000-1023022133330131-3023333003003323-2311003313332333-2203303111102220)
- voltstack_cluster.no_forward_proxy

<a id="canonical-0201303332132022-0202022011323033-0022321111110133-1120332310233103-2210130113113131-1213121212011000-0300130133233101-3220120211132303"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-2002132133032130-3233110031030102-3330310231023310-2131112313311120-2020211032333020-1033022032331300-2103003102000230-0302033223323122"></a>

## Direct properties — no_forward_proxy / 311331112120 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2213131033311231-1003122032202330-0222210231101330-2033303032312102-2313332232320033-2220313011321200-1131300101322130-1313232110110112"></a>

## Next pages — no_forward_proxy / 311331112120 / 4

- [voltstack_cluster](data-sources--gcp_vpc_site--reference--group-004.md#canonical-2213313332313212-1230123320203012-3120121021331113-3210101220221000-1023022133330131-3023333003003323-2311003313332333-2203303111102220)
- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-1132222133321101-1232102200201102-2111112102132122-0010023221133123-2333213300031302-1001011022311003-3121103220313010-2103321200303202)

<a id="canonical-2121211231121221-0022130123130200-3101033002313113-0111031203013220-1022112230202002-2300333321033013-2001133121303100-3111320300221331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1212011102013302-0203010012010000-2322110021013222-3122320312033021-2211001110312103-2100033231013230-2313121323021113-2300001110301232"></a>

## voltstack_cluster.no_global_network — no_global_network / 230310101312 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-1132222133321101-1232102200201102-2111112102132122-0010023221133123-2333213300031302-1001011022311003-3121103220313010-2103321200303202)
- [Property reference](data-sources--gcp_vpc_site--reference--group-001.md#canonical-0023231003100101-1010210121112300-0303103132302303-0331000323201313-3003311210331022-3000312312212303-0030202320102022-1232013113103201)
- [voltstack_cluster](data-sources--gcp_vpc_site--reference--group-004.md#canonical-2213313332313212-1230123320203012-3120121021331113-3210101220221000-1023022133330131-3023333003003323-2311003313332333-2203303111102220)
- voltstack_cluster.no_global_network

<a id="canonical-1013301121033122-0333032031223020-1003021032201301-3101203323033111-2113333301312312-1323330210323203-3031202002100222-3302031122103333"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-3333122213023303-0312133322201212-0023031321321131-1333223330110212-3230121122332232-1333231221130203-3033101213233200-0213201001230011"></a>

## Direct properties — no_global_network / 230310101312 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0211102233122013-2103111200200103-3231222020011312-2231332223200301-0022000122331320-3213231122121130-1232230131331032-0332022120203101"></a>

## Next pages — no_global_network / 230310101312 / 4

- [voltstack_cluster](data-sources--gcp_vpc_site--reference--group-004.md#canonical-2213313332313212-1230123320203012-3120121021331113-3210101220221000-1023022133330131-3023333003003323-2311003313332333-2203303111102220)
- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-1132222133321101-1232102200201102-2111112102132122-0010023221133123-2333213300031302-1001011022311003-3121103220313010-2103321200303202)

<a id="canonical-1300021122101133-3222103223323300-2230003000323301-1332332000001110-2013023301103021-2131301331231302-2001001030302030-0101302213302111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2103112212232313-1010033203230000-1203213231102301-3103310201203333-1332132100313032-2121022221113231-3321312102320120-1200220023103013"></a>

## voltstack_cluster.no_k8s_cluster — no_k8s_cluster / 122000033300 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-1132222133321101-1232102200201102-2111112102132122-0010023221133123-2333213300031302-1001011022311003-3121103220313010-2103321200303202)
- [Property reference](data-sources--gcp_vpc_site--reference--group-001.md#canonical-0023231003100101-1010210121112300-0303103132302303-0331000323201313-3003311210331022-3000312312212303-0030202320102022-1232013113103201)
- [voltstack_cluster](data-sources--gcp_vpc_site--reference--group-004.md#canonical-2213313332313212-1230123320203012-3120121021331113-3210101220221000-1023022133330131-3023333003003323-2311003313332333-2203303111102220)
- voltstack_cluster.no_k8s_cluster

<a id="canonical-3210331011103011-0110012321222232-1130112221121310-3011130231300032-2001010211212300-1320130031211030-3121130212120213-2102021230321321"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-2330233332330302-1222230100331320-1211233332020200-2102331121300310-1102133301113331-3121202000322320-1320120233201132-0201223203120112"></a>

## Direct properties — no_k8s_cluster / 122000033300 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3133111321212313-3030221303210010-0231300333102003-0300022201112322-1110300331221010-0312010322133333-0001312213300123-3020202030030220"></a>

## Next pages — no_k8s_cluster / 122000033300 / 4

- [voltstack_cluster](data-sources--gcp_vpc_site--reference--group-004.md#canonical-2213313332313212-1230123320203012-3120121021331113-3210101220221000-1023022133330131-3023333003003323-2311003313332333-2203303111102220)
- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-1132222133321101-1232102200201102-2111112102132122-0010023221133123-2333213300031302-1001011022311003-3121103220313010-2103321200303202)

<a id="canonical-3300231221023003-3303130230120101-0212230110323330-2012100003321210-0012002101030011-2113122210130011-3022320201002320-1232121200220033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1212133001213123-0211221001211112-0202333122232013-1011323012122213-1212120000320021-1001113120011233-1111131231102221-1231323002012113"></a>

## voltstack_cluster.no_network_policy — no_network_policy / 000113313001 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-1132222133321101-1232102200201102-2111112102132122-0010023221133123-2333213300031302-1001011022311003-3121103220313010-2103321200303202)
- [Property reference](data-sources--gcp_vpc_site--reference--group-001.md#canonical-0023231003100101-1010210121112300-0303103132302303-0331000323201313-3003311210331022-3000312312212303-0030202320102022-1232013113103201)
- [voltstack_cluster](data-sources--gcp_vpc_site--reference--group-004.md#canonical-2213313332313212-1230123320203012-3120121021331113-3210101220221000-1023022133330131-3023333003003323-2311003313332333-2203303111102220)
- voltstack_cluster.no_network_policy

<a id="canonical-2111033332332101-2013233102021121-2213103220013123-0111120113331213-0000210123032012-2302221032112012-3212000303301111-3210230321212113"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-2331013031001310-0013123203300301-1120303321132201-2023120031312123-2130222033200213-2203222122113223-1211231312301022-2210102323212303"></a>

## Direct properties — no_network_policy / 000113313001 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3031220330113311-2130022111213330-1012001210021321-1112310202002300-2332303302223021-0022100110113233-3013311332223213-2013211321323122"></a>

## Next pages — no_network_policy / 000113313001 / 4

- [voltstack_cluster](data-sources--gcp_vpc_site--reference--group-004.md#canonical-2213313332313212-1230123320203012-3120121021331113-3210101220221000-1023022133330131-3023333003003323-2311003313332333-2203303111102220)
- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-1132222133321101-1232102200201102-2111112102132122-0010023221133123-2333213300031302-1001011022311003-3121103220313010-2103321200303202)

<a id="canonical-1300332232132320-0120121023203213-1310021031011123-0130323110333310-0212320211302213-3131302032101122-2020122131013122-0303133222302022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3210322223321313-2221311221322220-1312231011223023-0103022120110133-0322023332301023-1232013320131213-2220021003122130-1110220302212021"></a>

## voltstack_cluster.no_outside_static_routes — no_outside_static_routes / 023232210112 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-1132222133321101-1232102200201102-2111112102132122-0010023221133123-2333213300031302-1001011022311003-3121103220313010-2103321200303202)
- [Property reference](data-sources--gcp_vpc_site--reference--group-001.md#canonical-0023231003100101-1010210121112300-0303103132302303-0331000323201313-3003311210331022-3000312312212303-0030202320102022-1232013113103201)
- [voltstack_cluster](data-sources--gcp_vpc_site--reference--group-004.md#canonical-2213313332313212-1230123320203012-3120121021331113-3210101220221000-1023022133330131-3023333003003323-2311003313332333-2203303111102220)
- voltstack_cluster.no_outside_static_routes

<a id="canonical-2210201330222020-2301110012323202-0230230230200301-2113331313132023-0322023021321120-3031320113032011-1131323011230233-1111212313030111"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-0231032102023323-2021322022221213-1321113001213322-1312332131132131-1223020122111130-0003001101011300-3122033201121102-3303111333210321"></a>

## Direct properties — no_outside_static_routes / 023232210112 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0232300010230113-0220210213032223-0112031103311000-0323220023231031-1321200131323000-2311231023222000-3120132020323011-1331223231100012"></a>

## Next pages — no_outside_static_routes / 023232210112 / 4

- [voltstack_cluster](data-sources--gcp_vpc_site--reference--group-004.md#canonical-2213313332313212-1230123320203012-3120121021331113-3210101220221000-1023022133330131-3023333003003323-2311003313332333-2203303111102220)
- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-1132222133321101-1232102200201102-2111112102132122-0010023221133123-2333213300031302-1001011022311003-3121103220313010-2103321200303202)

<a id="canonical-3302312100000302-0021313112011031-3212312101300110-3233320120212312-0000122013313331-0312112130203310-2200111320212103-1023131132323330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1132330323333211-1013330320002333-1112120120122021-0210010002132123-0223033111321200-3011123031013003-3113202022212123-1023323101312102"></a>

## voltstack_cluster.outside_static_routes — outside_static_routes / 133110310032 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-1132222133321101-1232102200201102-2111112102132122-0010023221133123-2333213300031302-1001011022311003-3121103220313010-2103321200303202)
- [Property reference](data-sources--gcp_vpc_site--reference--group-001.md#canonical-0023231003100101-1010210121112300-0303103132302303-0331000323201313-3003311210331022-3000312312212303-0030202320102022-1232013113103201)
- [voltstack_cluster](data-sources--gcp_vpc_site--reference--group-004.md#canonical-2213313332313212-1230123320203012-3120121021331113-3210101220221000-1023022133330131-3023333003003323-2311003313332333-2203303111102220)
- voltstack_cluster.outside_static_routes

<a id="canonical-3021201222101301-3232131222313211-1022333212333130-2313002013201203-3200211102122100-3311232311232212-1333220230021212-0000300321103203"></a>

Type: `"single"`. Computed.

Configuration parameter for outside static routes.

Upstream description:

List of static routes.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3211330023330312-1211333101120133-1101310123213131-2003310023110102-0113132023322121-1200303221101231-0033221203322310-1033120203003023"></a>

## Direct properties — outside_static_routes / 133110310032 / 3

- [static_route_list](data-sources--gcp_vpc_site--reference--group-004.md#canonical-2132311322103200-3122023210331002-1310201320121212-0233002231320123-3031112223031013-2333033333122020-0311021302322132-2301203303312022): complete subsection reference.

<a id="canonical-0222022120323001-2022213022330001-2200123023321023-3003033003322011-3102032331003022-2113120330213210-0110231022212110-2211312331103011"></a>

## Next pages — outside_static_routes / 133110310032 / 4

- [voltstack_cluster.outside_static_routes.static_route_list](data-sources--gcp_vpc_site--reference--group-004.md#canonical-2132311322103200-3122023210331002-1310201320121212-0233002231320123-3031112223031013-2333033333122020-0311021302322132-2301203303312022)
- [voltstack_cluster](data-sources--gcp_vpc_site--reference--group-004.md#canonical-2213313332313212-1230123320203012-3120121021331113-3210101220221000-1023022133330131-3023333003003323-2311003313332333-2203303111102220)
- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-1132222133321101-1232102200201102-2111112102132122-0010023221133123-2333213300031302-1001011022311003-3121103220313010-2103321200303202)

<a id="canonical-2132311322103200-3122023210331002-1310201320121212-0233002231320123-3031112223031013-2333033333122020-0311021302322132-2301203303312022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0330303032201321-3101023013203031-0020132222201312-2032010220122303-2301001302023220-3312030021223013-3030102021120313-2332322231002131"></a>

## voltstack_cluster.outside_static_routes.static_route_list — static_route_list / 032330110130 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-1132222133321101-1232102200201102-2111112102132122-0010023221133123-2333213300031302-1001011022311003-3121103220313010-2103321200303202)
- [Property reference](data-sources--gcp_vpc_site--reference--group-001.md#canonical-0023231003100101-1010210121112300-0303103132302303-0331000323201313-3003311210331022-3000312312212303-0030202320102022-1232013113103201)
- [voltstack_cluster](data-sources--gcp_vpc_site--reference--group-004.md#canonical-2213313332313212-1230123320203012-3120121021331113-3210101220221000-1023022133330131-3023333003003323-2311003313332333-2203303111102220)
- [voltstack_cluster.outside_static_routes](data-sources--gcp_vpc_site--reference--group-004.md#canonical-3302312100000302-0021313112011031-3212312101300110-3233320120212312-0000122013313331-0312112130203310-2200111320212103-1023131132323330)
- voltstack_cluster.outside_static_routes.static_route_list

<a id="canonical-0132203322010303-3020013120102102-1121323033220132-2332001110202220-1213323102303231-1321131320032130-1023321031001313-0010223033230211"></a>

Type: `"list"`. Computed.

List of Static Routes. List of Static routes.

Upstream description:

List of Static routes.

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

<a id="canonical-2132221032031311-3310222013313222-1003010112313311-3322230011103313-2301201203130021-3031223102122302-1033301032021132-1313220322210211"></a>

## Direct properties — static_route_list / 032330110130 / 3

- [custom_static_route](data-sources--gcp_vpc_site--reference--group-004.md#canonical-2302310321110231-2212332132311032-2113203033022112-2232220102101030-1033100300202321-1301013111211301-1112133213131033-3132202230123212): complete subsection reference.

<a id="canonical-2200213103200322-1020332111133313-2032222032021121-0313303300123012-3322122230330112-3230310323020112-1120133113320032-0011132132120203"></a>

<a id="canonical-0110112130103202-1312132123232132-0220012120233032-2013001333322331-2213331003232230-1333022122301101-1012203301212320-1300012202320133"></a>

## simple_static_route property — static_route_list / 032330110130 / 4

Type: `"string"`. Computed.

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

<a id="canonical-2131331201001003-3103210103202303-2232021112313232-1031321132320211-1302311322213233-1323300102020103-3322022222202312-1102232003103230"></a>

## Next pages — static_route_list / 032330110130 / 5

- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route](data-sources--gcp_vpc_site--reference--group-004.md#canonical-2302310321110231-2212332132311032-2113203033022112-2232220102101030-1033100300202321-1301013111211301-1112133213131033-3132202230123212)
- [voltstack_cluster.outside_static_routes](data-sources--gcp_vpc_site--reference--group-004.md#canonical-3302312100000302-0021313112011031-3212312101300110-3233320120212312-0000122013313331-0312112130203310-2200111320212103-1023131132323330)
- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-1132222133321101-1232102200201102-2111112102132122-0010023221133123-2333213300031302-1001011022311003-3121103220313010-2103321200303202)

<a id="canonical-2302310321110231-2212332132311032-2113203033022112-2232220102101030-1033100300202321-1301013111211301-1112133213131033-3132202230123212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1233221212202030-2123020231223222-2313103200003221-0213032103111021-2012210202231023-3033321100010103-0031303033133313-0313333302301102"></a>

## voltstack_cluster.outside_static_routes.static_route_list.custom_static_route — custom_static_route / 120201231003 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-1132222133321101-1232102200201102-2111112102132122-0010023221133123-2333213300031302-1001011022311003-3121103220313010-2103321200303202)
- [Property reference](data-sources--gcp_vpc_site--reference--group-001.md#canonical-0023231003100101-1010210121112300-0303103132302303-0331000323201313-3003311210331022-3000312312212303-0030202320102022-1232013113103201)
- [voltstack_cluster](data-sources--gcp_vpc_site--reference--group-004.md#canonical-2213313332313212-1230123320203012-3120121021331113-3210101220221000-1023022133330131-3023333003003323-2311003313332333-2203303111102220)
- [voltstack_cluster.outside_static_routes](data-sources--gcp_vpc_site--reference--group-004.md#canonical-3302312100000302-0021313112011031-3212312101300110-3233320120212312-0000122013313331-0312112130203310-2200111320212103-1023131132323330)
- [voltstack_cluster.outside_static_routes.static_route_list](data-sources--gcp_vpc_site--reference--group-004.md#canonical-2132311322103200-3122023210331002-1310201320121212-0233002231320123-3031112223031013-2333033333122020-0311021302322132-2301203303312022)
- voltstack_cluster.outside_static_routes.static_route_list.custom_static_route

<a id="canonical-2322023120212230-3131030211111322-1131300130311332-0031023233220130-3223113022023103-2022203032123023-3123002101200330-0301132330222130"></a>

Type: `"single"`. Computed.

Defines a static route, configuring a list of prefixes and a next-hop to be used for them.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0201321013103101-0031020232320203-1233023020232311-3210011200103130-3230130133100132-1203230101100330-0130333112303021-2330230223332303"></a>

## Direct properties — custom_static_route / 120201231003 / 3

<a id="canonical-1000102113000112-3113022311332201-3000213120031101-1020010130133322-3200333220032022-0322311212121131-3001111010010031-1222121310232212"></a>

<a id="canonical-0323133010102220-3012210311001133-2010013312003120-2122003330111331-0200121011323303-0002310032233230-2100023200112220-0010022303101332"></a>

## attrs property — custom_static_route / 120201231003 / 4

Type: `["list", "string"]`. Computed.

\[Enum:
ROUTE\_ATTR\_NO\_OP|ROUTE\_ATTR\_ADVERTISE|ROUTE\_ATTR\_INSTALL\_HOST|ROUTE\_ATTR\_INSTALL\_FORWARDING|ROUTE\_ATTR\_MERGE\_ONLY\]
List of route attributes associated with the static route. Possible values are
\`ROUTE\_ATTR\_NO\_OP\`, \`ROUTE\_ATTR\_ADVERTISE\`, \`ROUTE\_ATTR\_INSTALL\_HOST\`,
\`ROUTE\_ATTR\_INSTALL\_FORWARDING\`, \`ROUTE\_ATTR\_MERGE\_ONLY\`. Defaults to
\`ROUTE\_ATTR\_NO\_OP\`.

Upstream description:

List of route attributes associated with the static route.

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

- [labels](data-sources--gcp_vpc_site--reference--group-004.md#canonical-3022102313221213-3322020113220010-1212310010111303-3133022002233032-2230130121012020-0032302003101032-2032111020101203-1331311121011231): complete subsection reference.

- [nexthop](data-sources--gcp_vpc_site--reference--group-004.md#canonical-2032001100133322-3313021103221232-1310223201231033-2012223100220002-2010303110300213-1221012220200212-2100100302111000-2103021030212303): complete subsection reference.

- [subnets](data-sources--gcp_vpc_site--reference--group-004.md#canonical-1222320031301012-0201020101100330-0030233313321330-1010321211302301-2321311212300222-0220132031201330-1100132133233122-0011031101010321): complete subsection reference.

<a id="canonical-0220210311210221-2300321223323331-3211010123100111-0120133323103330-0223332021010321-0131333220202330-0110221001020233-1300212321332032"></a>

## Next pages — custom_static_route / 120201231003 / 5

- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.labels](data-sources--gcp_vpc_site--reference--group-004.md#canonical-3022102313221213-3322020113220010-1212310010111303-3133022002233032-2230130121012020-0032302003101032-2032111020101203-1331311121011231)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop](data-sources--gcp_vpc_site--reference--group-004.md#canonical-2032001100133322-3313021103221232-1310223201231033-2012223100220002-2010303110300213-1221012220200212-2100100302111000-2103021030212303)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets](data-sources--gcp_vpc_site--reference--group-004.md#canonical-1222320031301012-0201020101100330-0030233313321330-1010321211302301-2321311212300222-0220132031201330-1100132133233122-0011031101010321)
- [voltstack_cluster.outside_static_routes.static_route_list](data-sources--gcp_vpc_site--reference--group-004.md#canonical-2132311322103200-3122023210331002-1310201320121212-0233002231320123-3031112223031013-2333033333122020-0311021302322132-2301203303312022)
- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-1132222133321101-1232102200201102-2111112102132122-0010023221133123-2333213300031302-1001011022311003-3121103220313010-2103321200303202)

<a id="canonical-3022102313221213-3322020113220010-1212310010111303-3133022002233032-2230130121012020-0032302003101032-2032111020101203-1331311121011231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0122320222133321-1010322031020001-1000110323132133-0222212030110021-1223233022333200-1011332032131303-2203031332000101-1003321201132122"></a>

## voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.labels — labels / 330302101302 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-1132222133321101-1232102200201102-2111112102132122-0010023221133123-2333213300031302-1001011022311003-3121103220313010-2103321200303202)
- [Property reference](data-sources--gcp_vpc_site--reference--group-001.md#canonical-0023231003100101-1010210121112300-0303103132302303-0331000323201313-3003311210331022-3000312312212303-0030202320102022-1232013113103201)
- [voltstack_cluster](data-sources--gcp_vpc_site--reference--group-004.md#canonical-2213313332313212-1230123320203012-3120121021331113-3210101220221000-1023022133330131-3023333003003323-2311003313332333-2203303111102220)
- [voltstack_cluster.outside_static_routes](data-sources--gcp_vpc_site--reference--group-004.md#canonical-3302312100000302-0021313112011031-3212312101300110-3233320120212312-0000122013313331-0312112130203310-2200111320212103-1023131132323330)
- [voltstack_cluster.outside_static_routes.static_route_list](data-sources--gcp_vpc_site--reference--group-004.md#canonical-2132311322103200-3122023210331002-1310201320121212-0233002231320123-3031112223031013-2333033333122020-0311021302322132-2301203303312022)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route](data-sources--gcp_vpc_site--reference--group-004.md#canonical-2302310321110231-2212332132311032-2113203033022112-2232220102101030-1033100300202321-1301013111211301-1112133213131033-3132202230123212)
- voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.labels

<a id="canonical-3011231031233121-0202322231020100-1132022210300020-1031220212213102-2313120031133131-2010300333022113-1331032310132122-3201302303111032"></a>

Type: `"single"`. Computed.

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

<a id="canonical-1302330031203032-1131310313122331-2133101213301202-0331313100321201-1120320302013311-1200031210003321-0111001122213232-2223310303003213"></a>

## Direct properties — labels / 330302101302 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2103200122110123-1203203323212313-2212310220322300-2031022222322010-0332012013022321-2021323100221202-3110332230030232-2130303300000331"></a>

## Next pages — labels / 330302101302 / 4

- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route](data-sources--gcp_vpc_site--reference--group-004.md#canonical-2302310321110231-2212332132311032-2113203033022112-2232220102101030-1033100300202321-1301013111211301-1112133213131033-3132202230123212)
- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-1132222133321101-1232102200201102-2111112102132122-0010023221133123-2333213300031302-1001011022311003-3121103220313010-2103321200303202)

<a id="canonical-2032001100133322-3313021103221232-1310223201231033-2012223100220002-2010303110300213-1221012220200212-2100100302111000-2103021030212303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2333111313131330-1020003012010312-2323132012112101-1132021002322113-3021221233123231-1213202210031223-2020112321310103-1310202330100310"></a>

## voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop — nexthop / 202332121000 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-1132222133321101-1232102200201102-2111112102132122-0010023221133123-2333213300031302-1001011022311003-3121103220313010-2103321200303202)
- [Property reference](data-sources--gcp_vpc_site--reference--group-001.md#canonical-0023231003100101-1010210121112300-0303103132302303-0331000323201313-3003311210331022-3000312312212303-0030202320102022-1232013113103201)
- [voltstack_cluster](data-sources--gcp_vpc_site--reference--group-004.md#canonical-2213313332313212-1230123320203012-3120121021331113-3210101220221000-1023022133330131-3023333003003323-2311003313332333-2203303111102220)
- [voltstack_cluster.outside_static_routes](data-sources--gcp_vpc_site--reference--group-004.md#canonical-3302312100000302-0021313112011031-3212312101300110-3233320120212312-0000122013313331-0312112130203310-2200111320212103-1023131132323330)
- [voltstack_cluster.outside_static_routes.static_route_list](data-sources--gcp_vpc_site--reference--group-004.md#canonical-2132311322103200-3122023210331002-1310201320121212-0233002231320123-3031112223031013-2333033333122020-0311021302322132-2301203303312022)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route](data-sources--gcp_vpc_site--reference--group-004.md#canonical-2302310321110231-2212332132311032-2113203033022112-2232220102101030-1033100300202321-1301013111211301-1112133213131033-3132202230123212)
- voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop

<a id="canonical-1202210222132031-3202202020200022-1131010102230021-0320133101030022-1312202120211120-0230000203331022-3011002032221030-1330020023300003"></a>

Type: `"single"`. Computed.

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

<a id="canonical-0001202032332032-0221323000201320-3032311310000102-1300310001100202-2302100302210233-1031033130222213-2232300101322312-2320301213221230"></a>

## Direct properties — nexthop / 202332121000 / 3

- [interface](data-sources--gcp_vpc_site--reference--group-004.md#canonical-1213012101332232-0133322232231030-1102122121131230-2322121023321230-3223220201223103-0212023121331110-0111313310303313-0230123332312002): complete subsection reference.

- [nexthop_address](data-sources--gcp_vpc_site--reference--group-004.md#canonical-0230132131000303-2200310000221221-3030223133010032-3321313222033320-0211213131333011-3210220000303221-0323311113311031-2120110122130223): complete subsection reference.

<a id="canonical-3212200001332303-1130123202020010-0021212100212030-0010113123031222-0212013332103011-3221223133300331-2220022101330223-0100131103010022"></a>

<a id="canonical-3200330203012221-1130130313033130-1103301123223201-2233033211112321-0320203001313232-0302221100303032-1003313232332101-3112201122002222"></a>

## type property — nexthop / 202332121000 / 4

Type: `"string"`. Computed.

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

<a id="canonical-3100130212023332-3213102112300203-1033221233231002-2220301111221000-1301022230212232-0223320223213110-1203123021013030-3331302303033200"></a>

## Next pages — nexthop / 202332121000 / 5

- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.interface](data-sources--gcp_vpc_site--reference--group-004.md#canonical-1213012101332232-0133322232231030-1102122121131230-2322121023321230-3223220201223103-0212023121331110-0111313310303313-0230123332312002)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](data-sources--gcp_vpc_site--reference--group-004.md#canonical-0230132131000303-2200310000221221-3030223133010032-3321313222033320-0211213131333011-3210220000303221-0323311113311031-2120110122130223)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route](data-sources--gcp_vpc_site--reference--group-004.md#canonical-2302310321110231-2212332132311032-2113203033022112-2232220102101030-1033100300202321-1301013111211301-1112133213131033-3132202230123212)
- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-1132222133321101-1232102200201102-2111112102132122-0010023221133123-2333213300031302-1001011022311003-3121103220313010-2103321200303202)

<a id="canonical-1213012101332232-0133322232231030-1102122121131230-2322121023321230-3223220201223103-0212023121331110-0111313310303313-0230123332312002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1330020021300200-3130130320333201-1213301203123023-2102102133101321-3320330212130230-2003223233232132-3213110320310020-3222313321322002"></a>

## voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.interface — interface / 300002121102 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-1132222133321101-1232102200201102-2111112102132122-0010023221133123-2333213300031302-1001011022311003-3121103220313010-2103321200303202)
- [Property reference](data-sources--gcp_vpc_site--reference--group-001.md#canonical-0023231003100101-1010210121112300-0303103132302303-0331000323201313-3003311210331022-3000312312212303-0030202320102022-1232013113103201)
- [voltstack_cluster](data-sources--gcp_vpc_site--reference--group-004.md#canonical-2213313332313212-1230123320203012-3120121021331113-3210101220221000-1023022133330131-3023333003003323-2311003313332333-2203303111102220)
- [voltstack_cluster.outside_static_routes](data-sources--gcp_vpc_site--reference--group-004.md#canonical-3302312100000302-0021313112011031-3212312101300110-3233320120212312-0000122013313331-0312112130203310-2200111320212103-1023131132323330)
- [voltstack_cluster.outside_static_routes.static_route_list](data-sources--gcp_vpc_site--reference--group-004.md#canonical-2132311322103200-3122023210331002-1310201320121212-0233002231320123-3031112223031013-2333033333122020-0311021302322132-2301203303312022)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route](data-sources--gcp_vpc_site--reference--group-004.md#canonical-2302310321110231-2212332132311032-2113203033022112-2232220102101030-1033100300202321-1301013111211301-1112133213131033-3132202230123212)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop](data-sources--gcp_vpc_site--reference--group-004.md#canonical-2032001100133322-3313021103221232-1310223201231033-2012223100220002-2010303110300213-1221012220200212-2100100302111000-2103021030212303)
- voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.interface

<a id="canonical-3012302233112010-2133011201302032-2330202002002303-3312112320103203-1100103320231031-2212010312132331-2332232231120133-3133122031031301"></a>

Type: `"list"`. Computed.

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

<a id="canonical-0023112310222030-1232033000000202-0332032100212211-3103121220121130-1303201202311122-1211221100311121-1303201330231331-0120032003032203"></a>

## Direct properties — interface / 300002121102 / 3

<a id="canonical-0213312003021032-1300013011312031-2231032032031003-3313013313010023-1223101101320330-1321223003331300-3103010231300111-0002131112111323"></a>

<a id="canonical-1010102003203213-1213303103013302-0331011032321102-2221111230003130-2100331330033103-1000211210222130-1001312003301200-0033200233220203"></a>

## kind property — interface / 300002121102 / 4

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

<a id="canonical-1120332121022223-2301232110331001-0330120021202333-3332212112001010-2020013130101221-2101122231321110-1013212321020032-2100000031322131"></a>

<a id="canonical-2312031111221031-0130021233211323-1200132102112313-1202300300022312-2230320223021220-0210002302133020-3103210231013231-1321113103101110"></a>

## name property — interface / 300002121102 / 5

Type: `"string"`. Computed.

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

<a id="canonical-2112112231330103-2010002200020310-0312121323210300-2003110013200002-2301212210113211-2320002031130031-0120322101132233-0022131201013101"></a>

<a id="canonical-3030323021212101-0000010112323201-2223300030312112-3010301211102321-0231020102021230-0011223022221311-2130302330022313-0210232030022333"></a>

## namespace property — interface / 300002121102 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

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

<a id="canonical-1030311320300122-2330321212331002-1120332220132230-2311022002011012-3022310030030321-1130230122000221-1131013011223010-2032231333213333"></a>

<a id="canonical-2111321333203131-2030021113333031-1220300331030133-2201131011331010-2302213112303000-1013233300010331-2330010113230311-1112121110202302"></a>

## tenant property — interface / 300002121102 / 7

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

<a id="canonical-3020021023110311-3202300103003013-1022131333022230-3223122030033331-1202013202130011-1232130013010313-0002320222331011-0312122322032231"></a>

<a id="canonical-1031200330322331-2321020300122232-0221013111221001-3023111003002302-2132323232121300-0003013211232022-0331223011010300-0211001221030130"></a>

## uid property — interface / 300002121102 / 8

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

<a id="canonical-0332100020312113-1301330323133222-1013112203321001-3032321033003031-1112001323121321-2220302300213300-2110203300112031-2232313113003302"></a>

## Next pages — interface / 300002121102 / 9

- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop](data-sources--gcp_vpc_site--reference--group-004.md#canonical-2032001100133322-3313021103221232-1310223201231033-2012223100220002-2010303110300213-1221012220200212-2100100302111000-2103021030212303)
- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-1132222133321101-1232102200201102-2111112102132122-0010023221133123-2333213300031302-1001011022311003-3121103220313010-2103321200303202)

<a id="canonical-0230132131000303-2200310000221221-3030223133010032-3321313222033320-0211213131333011-3210220000303221-0323311113311031-2120110122130223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3321203323123212-2102230103211302-1322323111322022-2320030123221033-0121302311132333-0033311212220221-3121300303231121-2032002220003012"></a>

## voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address — nexthop_address / 321333203133 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-1132222133321101-1232102200201102-2111112102132122-0010023221133123-2333213300031302-1001011022311003-3121103220313010-2103321200303202)
- [Property reference](data-sources--gcp_vpc_site--reference--group-001.md#canonical-0023231003100101-1010210121112300-0303103132302303-0331000323201313-3003311210331022-3000312312212303-0030202320102022-1232013113103201)
- [voltstack_cluster](data-sources--gcp_vpc_site--reference--group-004.md#canonical-2213313332313212-1230123320203012-3120121021331113-3210101220221000-1023022133330131-3023333003003323-2311003313332333-2203303111102220)
- [voltstack_cluster.outside_static_routes](data-sources--gcp_vpc_site--reference--group-004.md#canonical-3302312100000302-0021313112011031-3212312101300110-3233320120212312-0000122013313331-0312112130203310-2200111320212103-1023131132323330)
- [voltstack_cluster.outside_static_routes.static_route_list](data-sources--gcp_vpc_site--reference--group-004.md#canonical-2132311322103200-3122023210331002-1310201320121212-0233002231320123-3031112223031013-2333033333122020-0311021302322132-2301203303312022)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route](data-sources--gcp_vpc_site--reference--group-004.md#canonical-2302310321110231-2212332132311032-2113203033022112-2232220102101030-1033100300202321-1301013111211301-1112133213131033-3132202230123212)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop](data-sources--gcp_vpc_site--reference--group-004.md#canonical-2032001100133322-3313021103221232-1310223201231033-2012223100220002-2010303110300213-1221012220200212-2100100302111000-2103021030212303)
- voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address

<a id="canonical-1311213101231300-3033023211310010-0033132321301022-0322200313010330-3110100010101331-2313122123113011-3312331320132112-3212300320030301"></a>

Type: `"single"`. Computed.

IP Address used to specify an IPv4 or IPv6 address.

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

<a id="canonical-0202100303133221-1031310213210320-3211103030311022-2002201112221022-0332332231011222-3100312202231233-3102311033312232-2201101022223130"></a>

## Direct properties — nexthop_address / 321333203133 / 3

- [dual_stack](data-sources--gcp_vpc_site--reference--group-004.md#canonical-3211221123012203-3123211320133102-0003002001011031-0231133032200011-1310330232132032-1030303233002221-3113302031110030-2221112033321330): complete subsection reference.

- [ipv4](data-sources--gcp_vpc_site--reference--group-004.md#canonical-3220211222122030-2121321330212013-0232301032320121-0103221031023122-2220232213031202-1201201333321021-3302032300020222-3211212012201212): complete subsection reference.

- [ipv6](data-sources--gcp_vpc_site--reference--group-004.md#canonical-2201312000120001-0013012301221030-0222223002012223-0232300301320011-3201013230313212-2121212022110221-1113013320232331-2312310310003320): complete subsection reference.

<a id="canonical-1031132112200110-2332220321021302-2332010003122110-3302000313112323-3130323200010212-3330113112321102-0232212233233112-3203302000303220"></a>

## Next pages — nexthop_address / 321333203133 / 4

- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](data-sources--gcp_vpc_site--reference--group-004.md#canonical-3211221123012203-3123211320133102-0003002001011031-0231133032200011-1310330232132032-1030303233002221-3113302031110030-2221112033321330)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4](data-sources--gcp_vpc_site--reference--group-004.md#canonical-3220211222122030-2121321330212013-0232301032320121-0103221031023122-2220232213031202-1201201333321021-3302032300020222-3211212012201212)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6](data-sources--gcp_vpc_site--reference--group-004.md#canonical-2201312000120001-0013012301221030-0222223002012223-0232300301320011-3201013230313212-2121212022110221-1113013320232331-2312310310003320)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop](data-sources--gcp_vpc_site--reference--group-004.md#canonical-2032001100133322-3313021103221232-1310223201231033-2012223100220002-2010303110300213-1221012220200212-2100100302111000-2103021030212303)
- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-1132222133321101-1232102200201102-2111112102132122-0010023221133123-2333213300031302-1001011022311003-3121103220313010-2103321200303202)

<a id="canonical-3211221123012203-3123211320133102-0003002001011031-0231133032200011-1310330232132032-1030303233002221-3113302031110030-2221112033321330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1310023100123223-2120103231103321-2301212132113030-0002211300131030-0113021210302111-1312332120210102-2220200221332000-2300223033321012"></a>

## voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack — dual_stack / 021100201211 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-1132222133321101-1232102200201102-2111112102132122-0010023221133123-2333213300031302-1001011022311003-3121103220313010-2103321200303202)
- [Property reference](data-sources--gcp_vpc_site--reference--group-001.md#canonical-0023231003100101-1010210121112300-0303103132302303-0331000323201313-3003311210331022-3000312312212303-0030202320102022-1232013113103201)
- [voltstack_cluster](data-sources--gcp_vpc_site--reference--group-004.md#canonical-2213313332313212-1230123320203012-3120121021331113-3210101220221000-1023022133330131-3023333003003323-2311003313332333-2203303111102220)
- [voltstack_cluster.outside_static_routes](data-sources--gcp_vpc_site--reference--group-004.md#canonical-3302312100000302-0021313112011031-3212312101300110-3233320120212312-0000122013313331-0312112130203310-2200111320212103-1023131132323330)
- [voltstack_cluster.outside_static_routes.static_route_list](data-sources--gcp_vpc_site--reference--group-004.md#canonical-2132311322103200-3122023210331002-1310201320121212-0233002231320123-3031112223031013-2333033333122020-0311021302322132-2301203303312022)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route](data-sources--gcp_vpc_site--reference--group-004.md#canonical-2302310321110231-2212332132311032-2113203033022112-2232220102101030-1033100300202321-1301013111211301-1112133213131033-3132202230123212)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop](data-sources--gcp_vpc_site--reference--group-004.md#canonical-2032001100133322-3313021103221232-1310223201231033-2012223100220002-2010303110300213-1221012220200212-2100100302111000-2103021030212303)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](data-sources--gcp_vpc_site--reference--group-004.md#canonical-0230132131000303-2200310000221221-3030223133010032-3321313222033320-0211213131333011-3210220000303221-0323311113311031-2120110122130223)
- voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack

<a id="canonical-0031320202030113-1113121012311220-1210202222302313-1033203110200110-2111303000101133-2300313122200212-2221230303132223-2012111232220021"></a>

Type: `"single"`. Computed.

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

<a id="canonical-0003011023302131-1122313220022312-2111320123211012-3212030202030131-1101300001212331-1303300311301031-3310322113212200-1030331200220001"></a>

## Direct properties — dual_stack / 021100201211 / 3

- [ipv4](data-sources--gcp_vpc_site--reference--group-004.md#canonical-3101310331031013-3331323310020121-1132200330000333-0101023132001000-3131210333132101-3002000102202303-2311223112100111-3320202320210023): complete subsection reference.

- [ipv6](data-sources--gcp_vpc_site--reference--group-004.md#canonical-1020110233000333-0203021110122213-2331023101321203-3010232010101231-0231301013133303-0023223021200120-1030213100012032-3312203010122200): complete subsection reference.

<a id="canonical-0323113001321200-0230000112233120-2323131312033223-0023122110230302-1121212233211331-3033123120121003-1003310031222220-0330303312122020"></a>

## Next pages — dual_stack / 021100201211 / 4

- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4](data-sources--gcp_vpc_site--reference--group-004.md#canonical-3101310331031013-3331323310020121-1132200330000333-0101023132001000-3131210333132101-3002000102202303-2311223112100111-3320202320210023)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6](data-sources--gcp_vpc_site--reference--group-004.md#canonical-1020110233000333-0203021110122213-2331023101321203-3010232010101231-0231301013133303-0023223021200120-1030213100012032-3312203010122200)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](data-sources--gcp_vpc_site--reference--group-004.md#canonical-0230132131000303-2200310000221221-3030223133010032-3321313222033320-0211213131333011-3210220000303221-0323311113311031-2120110122130223)
- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-1132222133321101-1232102200201102-2111112102132122-0010023221133123-2333213300031302-1001011022311003-3121103220313010-2103321200303202)

<a id="canonical-3101310331031013-3331323310020121-1132200330000333-0101023132001000-3131210333132101-3002000102202303-2311223112100111-3320202320210023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3100221032301230-2023333330121110-2203333010111330-3012002201321011-0301010031020001-0200231033220332-0200002201011213-3312323131320022"></a>

## voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.IPv4 — IPv4 / 132103313032 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-1132222133321101-1232102200201102-2111112102132122-0010023221133123-2333213300031302-1001011022311003-3121103220313010-2103321200303202)
- [Property reference](data-sources--gcp_vpc_site--reference--group-001.md#canonical-0023231003100101-1010210121112300-0303103132302303-0331000323201313-3003311210331022-3000312312212303-0030202320102022-1232013113103201)
- [voltstack_cluster](data-sources--gcp_vpc_site--reference--group-004.md#canonical-2213313332313212-1230123320203012-3120121021331113-3210101220221000-1023022133330131-3023333003003323-2311003313332333-2203303111102220)
- [voltstack_cluster.outside_static_routes](data-sources--gcp_vpc_site--reference--group-004.md#canonical-3302312100000302-0021313112011031-3212312101300110-3233320120212312-0000122013313331-0312112130203310-2200111320212103-1023131132323330)
- [voltstack_cluster.outside_static_routes.static_route_list](data-sources--gcp_vpc_site--reference--group-004.md#canonical-2132311322103200-3122023210331002-1310201320121212-0233002231320123-3031112223031013-2333033333122020-0311021302322132-2301203303312022)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route](data-sources--gcp_vpc_site--reference--group-004.md#canonical-2302310321110231-2212332132311032-2113203033022112-2232220102101030-1033100300202321-1301013111211301-1112133213131033-3132202230123212)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop](data-sources--gcp_vpc_site--reference--group-004.md#canonical-2032001100133322-3313021103221232-1310223201231033-2012223100220002-2010303110300213-1221012220200212-2100100302111000-2103021030212303)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](data-sources--gcp_vpc_site--reference--group-004.md#canonical-0230132131000303-2200310000221221-3030223133010032-3321313222033320-0211213131333011-3210220000303221-0323311113311031-2120110122130223)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](data-sources--gcp_vpc_site--reference--group-004.md#canonical-3211221123012203-3123211320133102-0003002001011031-0231133032200011-1310330232132032-1030303233002221-3113302031110030-2221112033321330)
- voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.IPv4

<a id="canonical-1001200133321233-1100311001112331-2012000233203320-2231110220031210-1132201013003233-1300322333201301-2031211110213103-2303332131230313"></a>

Type: `"single"`. Computed.

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

<a id="canonical-3012310133331201-3032301001330102-2223231000332133-0132232312200030-1120030001210000-1000000100330332-2132121223103011-2211012200110320"></a>

## Direct properties — IPv4 / 132103313032 / 3

<a id="canonical-0130032113112200-2012231230221331-2013203213320130-1223303022021320-0133221200222003-2313203202220210-2211232201302033-1032201203031032"></a>

<a id="canonical-3320311330010133-3213020030012221-3011312011013121-3111323303221012-2001230111200302-1332233301323320-2330312223201210-3301100001010333"></a>

## addr property — IPv4 / 132103313032 / 4

Type: `"string"`. Computed.

IPv4 Address in string form with dot-decimal notation.

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

<a id="canonical-3330011222120210-0110021001132203-0020031233302010-1333122002111010-0310231133232103-2013333032320121-0220302110102110-1020003103212123"></a>

## Next pages — IPv4 / 132103313032 / 5

- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](data-sources--gcp_vpc_site--reference--group-004.md#canonical-3211221123012203-3123211320133102-0003002001011031-0231133032200011-1310330232132032-1030303233002221-3113302031110030-2221112033321330)
- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-1132222133321101-1232102200201102-2111112102132122-0010023221133123-2333213300031302-1001011022311003-3121103220313010-2103321200303202)

<a id="canonical-1020110233000333-0203021110122213-2331023101321203-3010232010101231-0231301013133303-0023223021200120-1030213100012032-3312203010122200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1030202213300223-3211322322100201-0233131012023130-3221233323302311-0233222020300010-1203320112330133-0003030132221131-3331100032201120"></a>

## voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.IPv6 — IPv6 / 210120123011 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-1132222133321101-1232102200201102-2111112102132122-0010023221133123-2333213300031302-1001011022311003-3121103220313010-2103321200303202)
- [Property reference](data-sources--gcp_vpc_site--reference--group-001.md#canonical-0023231003100101-1010210121112300-0303103132302303-0331000323201313-3003311210331022-3000312312212303-0030202320102022-1232013113103201)
- [voltstack_cluster](data-sources--gcp_vpc_site--reference--group-004.md#canonical-2213313332313212-1230123320203012-3120121021331113-3210101220221000-1023022133330131-3023333003003323-2311003313332333-2203303111102220)
- [voltstack_cluster.outside_static_routes](data-sources--gcp_vpc_site--reference--group-004.md#canonical-3302312100000302-0021313112011031-3212312101300110-3233320120212312-0000122013313331-0312112130203310-2200111320212103-1023131132323330)
- [voltstack_cluster.outside_static_routes.static_route_list](data-sources--gcp_vpc_site--reference--group-004.md#canonical-2132311322103200-3122023210331002-1310201320121212-0233002231320123-3031112223031013-2333033333122020-0311021302322132-2301203303312022)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route](data-sources--gcp_vpc_site--reference--group-004.md#canonical-2302310321110231-2212332132311032-2113203033022112-2232220102101030-1033100300202321-1301013111211301-1112133213131033-3132202230123212)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop](data-sources--gcp_vpc_site--reference--group-004.md#canonical-2032001100133322-3313021103221232-1310223201231033-2012223100220002-2010303110300213-1221012220200212-2100100302111000-2103021030212303)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](data-sources--gcp_vpc_site--reference--group-004.md#canonical-0230132131000303-2200310000221221-3030223133010032-3321313222033320-0211213131333011-3210220000303221-0323311113311031-2120110122130223)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](data-sources--gcp_vpc_site--reference--group-004.md#canonical-3211221123012203-3123211320133102-0003002001011031-0231133032200011-1310330232132032-1030303233002221-3113302031110030-2221112033321330)
- voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.IPv6

<a id="canonical-3332220102010130-0311010310033300-0031210321200220-0123120201213012-3131210231210203-1103322020122323-2003010303130213-2031110110012220"></a>

Type: `"single"`. Computed.

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

<a id="canonical-0331232221232012-2000110200311222-1222120101211102-0121302310032331-2003202211033011-3123300002011022-3302101331111101-2202230102130320"></a>

## Direct properties — IPv6 / 210120123011 / 3

<a id="canonical-3232001301020001-0132133332321021-3213111122031303-0020022010321001-2033320030332001-3200000213301320-3031303210220102-2023303202200000"></a>

<a id="canonical-2003010110201131-0013300100110033-2133102312230120-0022320231032111-1221111321020031-2330322313323220-2221031033230311-1030011003003300"></a>

## addr property — IPv6 / 210120123011 / 4

Type: `"string"`. Computed.

IPv6 Address in form of string. IPv6 address must be specified as hexadecimal numbers separated by
':' The address can be compacted by suppressing zeros e.g. '2001:db8:0:0:0:0:2:1' becomes
'2001:db8::2:1' or '2001:db8:0:0:0:2:0:0' becomes '2001:db8::2::'.

Upstream description:

IPv6 Address in form of string. IPv6 address must be specified as hexadecimal numbers separated by
':' The address can be compacted by suppressing zeros e.g. '2001:db8:0:0:0:0:2:1' becomes
'2001:db8::2:1' or '2001:db8:0:0:0:2:0:0' becomes '2001:db8::2::'

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

<a id="canonical-3210331113103123-1132131120033331-2113213030131100-3112120123300130-3113320332031223-3031102122313212-1022101021023111-3132102211033122"></a>

## Next pages — IPv6 / 210120123011 / 5

- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](data-sources--gcp_vpc_site--reference--group-004.md#canonical-3211221123012203-3123211320133102-0003002001011031-0231133032200011-1310330232132032-1030303233002221-3113302031110030-2221112033321330)
- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-1132222133321101-1232102200201102-2111112102132122-0010023221133123-2333213300031302-1001011022311003-3121103220313010-2103321200303202)

<a id="canonical-3220211222122030-2121321330212013-0232301032320121-0103221031023122-2220232213031202-1201201333321021-3302032300020222-3211212012201212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1001221001220030-2312313032113330-1323003313002103-2113231100202133-1030032113130002-0302330112012301-3321020230111002-1230113101002103"></a>

## voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.IPv4 — IPv4 / 000202213203 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-1132222133321101-1232102200201102-2111112102132122-0010023221133123-2333213300031302-1001011022311003-3121103220313010-2103321200303202)
- [Property reference](data-sources--gcp_vpc_site--reference--group-001.md#canonical-0023231003100101-1010210121112300-0303103132302303-0331000323201313-3003311210331022-3000312312212303-0030202320102022-1232013113103201)
- [voltstack_cluster](data-sources--gcp_vpc_site--reference--group-004.md#canonical-2213313332313212-1230123320203012-3120121021331113-3210101220221000-1023022133330131-3023333003003323-2311003313332333-2203303111102220)
- [voltstack_cluster.outside_static_routes](data-sources--gcp_vpc_site--reference--group-004.md#canonical-3302312100000302-0021313112011031-3212312101300110-3233320120212312-0000122013313331-0312112130203310-2200111320212103-1023131132323330)
- [voltstack_cluster.outside_static_routes.static_route_list](data-sources--gcp_vpc_site--reference--group-004.md#canonical-2132311322103200-3122023210331002-1310201320121212-0233002231320123-3031112223031013-2333033333122020-0311021302322132-2301203303312022)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route](data-sources--gcp_vpc_site--reference--group-004.md#canonical-2302310321110231-2212332132311032-2113203033022112-2232220102101030-1033100300202321-1301013111211301-1112133213131033-3132202230123212)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop](data-sources--gcp_vpc_site--reference--group-004.md#canonical-2032001100133322-3313021103221232-1310223201231033-2012223100220002-2010303110300213-1221012220200212-2100100302111000-2103021030212303)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](data-sources--gcp_vpc_site--reference--group-004.md#canonical-0230132131000303-2200310000221221-3030223133010032-3321313222033320-0211213131333011-3210220000303221-0323311113311031-2120110122130223)
- voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.IPv4

<a id="canonical-0030220132002131-1331002300320020-2112131013013131-2003323331212001-1130233333010003-3032220003223212-1302222120332102-1103031003311221"></a>

Type: `"single"`. Computed.

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

<a id="canonical-2333301300133322-0122203212231222-1332320003021121-3020301102001300-1032221200010223-0101210122033121-3312230013021030-3220330010100013"></a>

## Direct properties — IPv4 / 000202213203 / 3

<a id="canonical-0121223312202203-1220232333321110-3331302100230300-2223101301010103-1212012312230012-3321322020010011-2032200112121333-2023233100330013"></a>

<a id="canonical-3133200013232000-2321131200232322-1002221103122303-1031120103310122-2133223002212233-2301111312220322-2303301320121212-1131133013033002"></a>

## addr property — IPv4 / 000202213203 / 4

Type: `"string"`. Computed.

IPv4 Address in string form with dot-decimal notation.

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

<a id="canonical-1221110030101233-2031230333321203-1120110123100033-1101310320212033-1311120211233132-1010212110111232-2200313011223000-1111130230331032"></a>

## Next pages — IPv4 / 000202213203 / 5

- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](data-sources--gcp_vpc_site--reference--group-004.md#canonical-0230132131000303-2200310000221221-3030223133010032-3321313222033320-0211213131333011-3210220000303221-0323311113311031-2120110122130223)
- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-1132222133321101-1232102200201102-2111112102132122-0010023221133123-2333213300031302-1001011022311003-3121103220313010-2103321200303202)

<a id="canonical-2201312000120001-0013012301221030-0222223002012223-0232300301320011-3201013230313212-2121212022110221-1113013320232331-2312310310003320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2300203103230120-1000001222002121-2023231222100112-3010003112233301-0021021201331202-3010130013023222-2102133312011131-0200031221212011"></a>

## voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.IPv6 — IPv6 / 023212313223 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-1132222133321101-1232102200201102-2111112102132122-0010023221133123-2333213300031302-1001011022311003-3121103220313010-2103321200303202)
- [Property reference](data-sources--gcp_vpc_site--reference--group-001.md#canonical-0023231003100101-1010210121112300-0303103132302303-0331000323201313-3003311210331022-3000312312212303-0030202320102022-1232013113103201)
- [voltstack_cluster](data-sources--gcp_vpc_site--reference--group-004.md#canonical-2213313332313212-1230123320203012-3120121021331113-3210101220221000-1023022133330131-3023333003003323-2311003313332333-2203303111102220)
- [voltstack_cluster.outside_static_routes](data-sources--gcp_vpc_site--reference--group-004.md#canonical-3302312100000302-0021313112011031-3212312101300110-3233320120212312-0000122013313331-0312112130203310-2200111320212103-1023131132323330)
- [voltstack_cluster.outside_static_routes.static_route_list](data-sources--gcp_vpc_site--reference--group-004.md#canonical-2132311322103200-3122023210331002-1310201320121212-0233002231320123-3031112223031013-2333033333122020-0311021302322132-2301203303312022)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route](data-sources--gcp_vpc_site--reference--group-004.md#canonical-2302310321110231-2212332132311032-2113203033022112-2232220102101030-1033100300202321-1301013111211301-1112133213131033-3132202230123212)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop](data-sources--gcp_vpc_site--reference--group-004.md#canonical-2032001100133322-3313021103221232-1310223201231033-2012223100220002-2010303110300213-1221012220200212-2100100302111000-2103021030212303)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](data-sources--gcp_vpc_site--reference--group-004.md#canonical-0230132131000303-2200310000221221-3030223133010032-3321313222033320-0211213131333011-3210220000303221-0323311113311031-2120110122130223)
- voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.IPv6

<a id="canonical-0012321212000301-0110110202033002-3100011202320102-0030103031120210-2030210011213331-2232303310211110-0002310231002223-1213012300202111"></a>

Type: `"single"`. Computed.

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

<a id="canonical-3232121001003322-3120311000133111-2221201033320212-3133302010232313-2203023000311231-0322301011212031-0200102320130312-3310132111131031"></a>

## Direct properties — IPv6 / 023212313223 / 3

<a id="canonical-2133003110022331-2330202301333301-3111120203012013-1203120010232231-3020112102123233-0331122111101123-1213101130120010-1112113102221111"></a>

<a id="canonical-0001032010312233-3001333032111213-3303320331321113-0002102212123220-0231331120111100-0031320012013332-1311200001133322-1322002102121210"></a>

## addr property — IPv6 / 023212313223 / 4

Type: `"string"`. Computed.

IPv6 Address in form of string. IPv6 address must be specified as hexadecimal numbers separated by
':' The address can be compacted by suppressing zeros e.g. '2001:db8:0:0:0:0:2:1' becomes
'2001:db8::2:1' or '2001:db8:0:0:0:2:0:0' becomes '2001:db8::2::'.

Upstream description:

IPv6 Address in form of string. IPv6 address must be specified as hexadecimal numbers separated by
':' The address can be compacted by suppressing zeros e.g. '2001:db8:0:0:0:0:2:1' becomes
'2001:db8::2:1' or '2001:db8:0:0:0:2:0:0' becomes '2001:db8::2::'

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

<a id="canonical-0330223130221320-2000022112013122-3232301030233321-3121332232121102-2012100232000233-3302301232201133-2110213321330310-1302010033110033"></a>

## Next pages — IPv6 / 023212313223 / 5

- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](data-sources--gcp_vpc_site--reference--group-004.md#canonical-0230132131000303-2200310000221221-3030223133010032-3321313222033320-0211213131333011-3210220000303221-0323311113311031-2120110122130223)
- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-1132222133321101-1232102200201102-2111112102132122-0010023221133123-2333213300031302-1001011022311003-3121103220313010-2103321200303202)

<a id="canonical-1222320031301012-0201020101100330-0030233313321330-1010321211302301-2321311212300222-0220132031201330-1100132133233122-0011031101010321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3103023321213023-3120021131003120-2011231120212211-2303312323203133-3102330230220223-2313030023302030-3321111132202133-2000311333301000"></a>

## voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets — subnets / 302230313221 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-1132222133321101-1232102200201102-2111112102132122-0010023221133123-2333213300031302-1001011022311003-3121103220313010-2103321200303202)
- [Property reference](data-sources--gcp_vpc_site--reference--group-001.md#canonical-0023231003100101-1010210121112300-0303103132302303-0331000323201313-3003311210331022-3000312312212303-0030202320102022-1232013113103201)
- [voltstack_cluster](data-sources--gcp_vpc_site--reference--group-004.md#canonical-2213313332313212-1230123320203012-3120121021331113-3210101220221000-1023022133330131-3023333003003323-2311003313332333-2203303111102220)
- [voltstack_cluster.outside_static_routes](data-sources--gcp_vpc_site--reference--group-004.md#canonical-3302312100000302-0021313112011031-3212312101300110-3233320120212312-0000122013313331-0312112130203310-2200111320212103-1023131132323330)
- [voltstack_cluster.outside_static_routes.static_route_list](data-sources--gcp_vpc_site--reference--group-004.md#canonical-2132311322103200-3122023210331002-1310201320121212-0233002231320123-3031112223031013-2333033333122020-0311021302322132-2301203303312022)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route](data-sources--gcp_vpc_site--reference--group-004.md#canonical-2302310321110231-2212332132311032-2113203033022112-2232220102101030-1033100300202321-1301013111211301-1112133213131033-3132202230123212)
- voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets

<a id="canonical-3301031311030203-0301302210223303-0123031010113012-2322211301211202-0322200132202301-1323103311230312-0120223020330331-2230201222120333"></a>

Type: `"list"`. Computed.

Subnets. List of route prefixes.

Upstream description:

List of route prefixes.

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

<a id="canonical-3113002311112103-2130231322131331-1303020310332032-0200233301200132-0112321233230310-1023113112330301-2011221212213033-1022312323021231"></a>

## Direct properties — subnets / 302230313221 / 3

- [ipv4](data-sources--gcp_vpc_site--reference--group-004.md#canonical-3010023001210001-2121311132221111-3022220120202211-3101030012033222-1330233232133002-2120210313031210-2012203311331123-3100201220211023): complete subsection reference.

- [ipv6](data-sources--gcp_vpc_site--reference--group-004.md#canonical-2313020111201130-1011222311010033-2031331202103123-2333123232133322-2210001220111222-2122311101002331-2001233331101020-1220103001123030): complete subsection reference.

<a id="canonical-1300110223020102-0120220321320311-1202123321300310-0010133021303233-0311110022012010-1010101202103012-2123000231020113-2333211320213013"></a>

## Next pages — subnets / 302230313221 / 4

- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4](data-sources--gcp_vpc_site--reference--group-004.md#canonical-3010023001210001-2121311132221111-3022220120202211-3101030012033222-1330233232133002-2120210313031210-2012203311331123-3100201220211023)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6](data-sources--gcp_vpc_site--reference--group-004.md#canonical-2313020111201130-1011222311010033-2031331202103123-2333123232133322-2210001220111222-2122311101002331-2001233331101020-1220103001123030)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route](data-sources--gcp_vpc_site--reference--group-004.md#canonical-2302310321110231-2212332132311032-2113203033022112-2232220102101030-1033100300202321-1301013111211301-1112133213131033-3132202230123212)
- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-1132222133321101-1232102200201102-2111112102132122-0010023221133123-2333213300031302-1001011022311003-3121103220313010-2103321200303202)

<a id="canonical-3010023001210001-2121311132221111-3022220120202211-3101030012033222-1330233232133002-2120210313031210-2012203311331123-3100201220211023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1232322311020132-1002233330320112-1212122230120130-0201223322212010-2320123113011313-0212112132303232-1230011212212123-2230013211123321"></a>

## voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets.IPv4 — IPv4 / 232112221201 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-1132222133321101-1232102200201102-2111112102132122-0010023221133123-2333213300031302-1001011022311003-3121103220313010-2103321200303202)
- [Property reference](data-sources--gcp_vpc_site--reference--group-001.md#canonical-0023231003100101-1010210121112300-0303103132302303-0331000323201313-3003311210331022-3000312312212303-0030202320102022-1232013113103201)
- [voltstack_cluster](data-sources--gcp_vpc_site--reference--group-004.md#canonical-2213313332313212-1230123320203012-3120121021331113-3210101220221000-1023022133330131-3023333003003323-2311003313332333-2203303111102220)
- [voltstack_cluster.outside_static_routes](data-sources--gcp_vpc_site--reference--group-004.md#canonical-3302312100000302-0021313112011031-3212312101300110-3233320120212312-0000122013313331-0312112130203310-2200111320212103-1023131132323330)
- [voltstack_cluster.outside_static_routes.static_route_list](data-sources--gcp_vpc_site--reference--group-004.md#canonical-2132311322103200-3122023210331002-1310201320121212-0233002231320123-3031112223031013-2333033333122020-0311021302322132-2301203303312022)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route](data-sources--gcp_vpc_site--reference--group-004.md#canonical-2302310321110231-2212332132311032-2113203033022112-2232220102101030-1033100300202321-1301013111211301-1112133213131033-3132202230123212)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets](data-sources--gcp_vpc_site--reference--group-004.md#canonical-1222320031301012-0201020101100330-0030233313321330-1010321211302301-2321311212300222-0220132031201330-1100132133233122-0011031101010321)
- voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets.IPv4

<a id="canonical-3003203323233113-1222113223233000-0211220202012330-1232332212011102-0310321001131231-1202020001102311-3323330321311232-3103021111200311"></a>

Type: `"single"`. Computed.

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

<a id="canonical-0031203002123133-2300322210123211-2233011100302302-1110000200321310-0011221202112321-0103131000011222-1030230011222133-0303031202120020"></a>

## Direct properties — IPv4 / 232112221201 / 3

<a id="canonical-3311211220132211-1202221310321011-1313022332002211-0332021202132311-1311210303231113-2311120000211220-3111213133131001-0133331213233323"></a>

<a id="canonical-0203100330330100-2003012201303311-2313311113213213-2310313022202130-3230131232311001-3010302120012122-3023323321302133-0130122003003303"></a>

## plen property — IPv4 / 232112221201 / 4

Type: `"number"`. Computed.

Prefix-length of the IPv4 subnet. Must be &lt;= 32.

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

<a id="canonical-1323130310030021-3323001230321031-2013232123211330-2001031211322233-3033121223133003-0112130131012233-1323222021302213-1030302222233320"></a>

<a id="canonical-1211233301001032-1131201210010210-0312213231332310-2223130013130031-2111131111331220-1302211033111212-2021001312231311-3020213103300210"></a>

## prefix property — IPv4 / 232112221201 / 5

Type: `"string"`. Computed.

Prefix part of the IPv4 subnet in string form with dot-decimal notation.

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

<a id="canonical-0213311300201003-3211333231330301-3323102003332031-0112232012020121-3323203330101021-3020121001113313-0332100020003100-3031000223132112"></a>

## Next pages — IPv4 / 232112221201 / 6

- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets](data-sources--gcp_vpc_site--reference--group-004.md#canonical-1222320031301012-0201020101100330-0030233313321330-1010321211302301-2321311212300222-0220132031201330-1100132133233122-0011031101010321)
- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-1132222133321101-1232102200201102-2111112102132122-0010023221133123-2333213300031302-1001011022311003-3121103220313010-2103321200303202)

<a id="canonical-2313020111201130-1011222311010033-2031331202103123-2333123232133322-2210001220111222-2122311101002331-2001233331101020-1220103001123030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3331230230320211-3211031112111230-3330122110210113-0202201021213002-0302022101213213-1103032001333030-3223020122020302-1103012202012010"></a>

## voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets.IPv6 — IPv6 / 322332200322 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-1132222133321101-1232102200201102-2111112102132122-0010023221133123-2333213300031302-1001011022311003-3121103220313010-2103321200303202)
- [Property reference](data-sources--gcp_vpc_site--reference--group-001.md#canonical-0023231003100101-1010210121112300-0303103132302303-0331000323201313-3003311210331022-3000312312212303-0030202320102022-1232013113103201)
- [voltstack_cluster](data-sources--gcp_vpc_site--reference--group-004.md#canonical-2213313332313212-1230123320203012-3120121021331113-3210101220221000-1023022133330131-3023333003003323-2311003313332333-2203303111102220)
- [voltstack_cluster.outside_static_routes](data-sources--gcp_vpc_site--reference--group-004.md#canonical-3302312100000302-0021313112011031-3212312101300110-3233320120212312-0000122013313331-0312112130203310-2200111320212103-1023131132323330)
- [voltstack_cluster.outside_static_routes.static_route_list](data-sources--gcp_vpc_site--reference--group-004.md#canonical-2132311322103200-3122023210331002-1310201320121212-0233002231320123-3031112223031013-2333033333122020-0311021302322132-2301203303312022)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route](data-sources--gcp_vpc_site--reference--group-004.md#canonical-2302310321110231-2212332132311032-2113203033022112-2232220102101030-1033100300202321-1301013111211301-1112133213131033-3132202230123212)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets](data-sources--gcp_vpc_site--reference--group-004.md#canonical-1222320031301012-0201020101100330-0030233313321330-1010321211302301-2321311212300222-0220132031201330-1100132133233122-0011031101010321)
- voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets.IPv6

<a id="canonical-2321031022113230-0031201301200122-2203003130222101-1232323130323233-0203133000020100-3230333122223211-2001131201313211-2112231103111300"></a>

Type: `"single"`. Computed.

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

<a id="canonical-2301223111111223-3023023011210010-0111210023222003-2111010121230321-2222320222311232-2211101102022301-2310311303303033-0232121111230210"></a>

## Direct properties — IPv6 / 322332200322 / 3

<a id="canonical-3303232121301001-3122021130113331-2333101001220313-1111212023001312-1123103013313102-0312100212030012-1000030113112322-3232133212023003"></a>

<a id="canonical-1022121321312133-3112331331323221-2133100031021012-3012003230310012-3311110223003210-3002312133211122-1202312323112321-2000233101302331"></a>

## plen property — IPv6 / 322332200322 / 4

Type: `"number"`. Computed.

Prefix length of the IPv6 subnet. Must be &lt;= 128.

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

<a id="canonical-3210103210212333-2130122210220201-2120121132100231-3312313211213212-0230102031321332-0010023113212202-0330120212001032-2131002320131111"></a>

<a id="canonical-1013203202213022-0200330103013320-1031112012023101-2133003111130132-0230202120211113-0100302332103232-2012301011221222-3131013000320202"></a>

## prefix property — IPv6 / 322332200322 / 5

Type: `"string"`. Computed.

Prefix part of the IPv6 subnet given in form of string. IPv6 address must be specified as
hexadecimal numbers separated by ':' e.g. '2001:db8:0:0:0:2:0:0' The address can be compacted by
suppressing zeros e.g. '2001:db8::2::'.

Upstream description:

Prefix part of the IPv6 subnet given in form of string. IPv6 address must be specified as
hexadecimal numbers separated by ':' e.g. "2001:db8:0:0:0:2:0:0" The address can be compacted by
suppressing zeros e.g. "2001:db8::2::"

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

<a id="canonical-0233200233112133-0233333222121020-0312103102331133-2213022000212121-2130000303210332-2312223003023133-0323322022102013-1222211202133330"></a>

## Next pages — IPv6 / 322332200322 / 6

- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets](data-sources--gcp_vpc_site--reference--group-004.md#canonical-1222320031301012-0201020101100330-0030233313321330-1010321211302301-2321311212300222-0220132031201330-1100132133233122-0011031101010321)
- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-1132222133321101-1232102200201102-2111112102132122-0010023221133123-2333213300031302-1001011022311003-3121103220313010-2103321200303202)

<a id="canonical-3111110233303120-3020123200210213-3113011221310021-2032102300121231-0001000301332210-1223222212122033-2103323121303210-2122130003210212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2313233313010021-3010002100002303-2201302322022320-1013031232132212-0232200122131123-2221331011300322-2012102201313200-3323221311320303"></a>

## voltstack_cluster.site_local_network — site_local_network / 213132022210 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-1132222133321101-1232102200201102-2111112102132122-0010023221133123-2333213300031302-1001011022311003-3121103220313010-2103321200303202)
- [Property reference](data-sources--gcp_vpc_site--reference--group-001.md#canonical-0023231003100101-1010210121112300-0303103132302303-0331000323201313-3003311210331022-3000312312212303-0030202320102022-1232013113103201)
- [voltstack_cluster](data-sources--gcp_vpc_site--reference--group-004.md#canonical-2213313332313212-1230123320203012-3120121021331113-3210101220221000-1023022133330131-3023333003003323-2311003313332333-2203303111102220)
- voltstack_cluster.site_local_network

<a id="canonical-0031131122131032-0222213322201333-1303321220010012-0003220011033221-2301010101203102-3100003023130132-2103302312333210-0303023012023020"></a>

Type: `"single"`. Computed.

Defines choice about GCP VPC network for a view.

Upstream description:

This defines choice about GCP VPC network for a view.

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

<a id="canonical-1303033112233133-1102211033332130-3002122011310200-0121030221122213-2213101313233221-3022023303022322-1121200232231112-2111301303000301"></a>

## Direct properties — site_local_network / 213132022210 / 3

- [existing_network](data-sources--gcp_vpc_site--reference--group-004.md#canonical-1032311213313303-0023202112201133-2322002111200102-1003011002322122-3133020301111311-3230132303233232-3230123101012330-2203311322003301): complete subsection reference.

- [new_network](data-sources--gcp_vpc_site--reference--group-004.md#canonical-1023033311231230-3222301330212210-0113322200221311-2123001030221000-2011133231101231-0310113212203301-1110313300120221-1313112112320132): complete subsection reference.

- [new_network_autogenerate](data-sources--gcp_vpc_site--reference--group-004.md#canonical-3012322021230313-2210131131112301-1230130221210122-1222223012222213-3022231302313322-1233233332220301-2213122013112303-1130223231321321): complete subsection reference.

<a id="canonical-1110112033303012-3323302323211022-0010331111000332-3221231210320333-2230031300332233-1211020300220120-3013321111123032-0203011033122332"></a>

## Next pages — site_local_network / 213132022210 / 4

- [voltstack_cluster.site_local_network.existing_network](data-sources--gcp_vpc_site--reference--group-004.md#canonical-1032311213313303-0023202112201133-2322002111200102-1003011002322122-3133020301111311-3230132303233232-3230123101012330-2203311322003301)
- [voltstack_cluster.site_local_network.new_network](data-sources--gcp_vpc_site--reference--group-004.md#canonical-1023033311231230-3222301330212210-0113322200221311-2123001030221000-2011133231101231-0310113212203301-1110313300120221-1313112112320132)
- [voltstack_cluster.site_local_network.new_network_autogenerate](data-sources--gcp_vpc_site--reference--group-004.md#canonical-3012322021230313-2210131131112301-1230130221210122-1222223012222213-3022231302313322-1233233332220301-2213122013112303-1130223231321321)
- [voltstack_cluster](data-sources--gcp_vpc_site--reference--group-004.md#canonical-2213313332313212-1230123320203012-3120121021331113-3210101220221000-1023022133330131-3023333003003323-2311003313332333-2203303111102220)
- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-1132222133321101-1232102200201102-2111112102132122-0010023221133123-2333213300031302-1001011022311003-3121103220313010-2103321200303202)

<a id="canonical-1032311213313303-0023202112201133-2322002111200102-1003011002322122-3133020301111311-3230132303233232-3230123101012330-2203311322003301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0011003302122230-0222011222232030-1333131221213312-2321321113031321-2310132133102102-1011333020202131-1001011321230031-0101331221212321"></a>

## voltstack_cluster.site_local_network.existing_network — existing_network / 222121312223 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-1132222133321101-1232102200201102-2111112102132122-0010023221133123-2333213300031302-1001011022311003-3121103220313010-2103321200303202)
- [Property reference](data-sources--gcp_vpc_site--reference--group-001.md#canonical-0023231003100101-1010210121112300-0303103132302303-0331000323201313-3003311210331022-3000312312212303-0030202320102022-1232013113103201)
- [voltstack_cluster](data-sources--gcp_vpc_site--reference--group-004.md#canonical-2213313332313212-1230123320203012-3120121021331113-3210101220221000-1023022133330131-3023333003003323-2311003313332333-2203303111102220)
- [voltstack_cluster.site_local_network](data-sources--gcp_vpc_site--reference--group-004.md#canonical-3111110233303120-3020123200210213-3113011221310021-2032102300121231-0001000301332210-1223222212122033-2103323121303210-2122130003210212)
- voltstack_cluster.site_local_network.existing_network

<a id="canonical-0230233221332300-1212003311332221-0320102302000333-2022223103003020-1301101330332332-2202101121020022-2203210100312103-1003210122332020"></a>

Type: `"single"`. Computed.

Configuration parameter for existing network.

Upstream description:

Name of existing VPC network.

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

<a id="canonical-1322303232312303-2020023031231123-2321121132113231-0200200130130013-2331122301212322-1030030333031112-1131200232101302-0113130132230121"></a>

## Direct properties — existing_network / 222121312223 / 3

<a id="canonical-2320321311110020-0231112123331333-1312231033300111-1010331023122022-2013201332000023-1311312133013133-2130022333101113-0002322212212012"></a>

<a id="canonical-3020120011011312-0101000313020220-2133223203022101-0330130201210021-3112330001212222-3110230311113222-3123030022320310-2133221023012021"></a>

## name property — existing_network / 222121312223 / 4

Type: `"string"`. Computed.

GCP VPC Network Name. Name for your GCP VPC Network.

Upstream description:

Name for your GCP VPC Network.

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

<a id="canonical-0103321321322032-1222023321223011-1203101222012311-1022022123312230-2221321233122310-3030003303113311-2322221101131201-2313232313121130"></a>

## Next pages — existing_network / 222121312223 / 5

- [voltstack_cluster.site_local_network](data-sources--gcp_vpc_site--reference--group-004.md#canonical-3111110233303120-3020123200210213-3113011221310021-2032102300121231-0001000301332210-1223222212122033-2103323121303210-2122130003210212)
- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-1132222133321101-1232102200201102-2111112102132122-0010023221133123-2333213300031302-1001011022311003-3121103220313010-2103321200303202)

<a id="canonical-1023033311231230-3222301330212210-0113322200221311-2123001030221000-2011133231101231-0310113212203301-1110313300120221-1313112112320132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0130120313123331-3220023230312010-3131132231313012-2123101221230002-3230200311222011-1200202331220213-1033231023032323-1232223003302310"></a>

## voltstack_cluster.site_local_network.new_network — new_network / 210133203103 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-1132222133321101-1232102200201102-2111112102132122-0010023221133123-2333213300031302-1001011022311003-3121103220313010-2103321200303202)
- [Property reference](data-sources--gcp_vpc_site--reference--group-001.md#canonical-0023231003100101-1010210121112300-0303103132302303-0331000323201313-3003311210331022-3000312312212303-0030202320102022-1232013113103201)
- [voltstack_cluster](data-sources--gcp_vpc_site--reference--group-004.md#canonical-2213313332313212-1230123320203012-3120121021331113-3210101220221000-1023022133330131-3023333003003323-2311003313332333-2203303111102220)
- [voltstack_cluster.site_local_network](data-sources--gcp_vpc_site--reference--group-004.md#canonical-3111110233303120-3020123200210213-3113011221310021-2032102300121231-0001000301332210-1223222212122033-2103323121303210-2122130003210212)
- voltstack_cluster.site_local_network.new_network

<a id="canonical-2201203010022200-2311122333203001-3301213233232200-2202023322122112-3333220302023022-0213230330023220-2000132301313003-0131100301100321"></a>

Type: `"single"`. Computed.

Parameters to create a new GCP VPC Network.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2213320321320121-0132210331120011-1032133112200003-1121132303100311-0020320202213310-2323233220111121-3132121102232310-2312232310013103"></a>

## Direct properties — new_network / 210133203103 / 3

<a id="canonical-0303330030320103-3021002020312130-0210203031022212-1300012231312233-1111322000221112-1311022203020311-1131201103033302-2221213103013221"></a>

<a id="canonical-3032001001203331-0132302220001310-3212130132320132-2322130101102232-0221102113323100-1022313302231331-1011003100101033-0220021203120030"></a>

## name property — new_network / 210133203103 / 4

Type: `"string"`. Computed.

GCP VPC Network Name. Name for your GCP VPC Network.

Upstream description:

Name for your GCP VPC Network.

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

<a id="canonical-2201312233223203-0221002232120113-0020033213213311-2001132302221211-1302101112101230-2200123111301300-1033101102203003-2113201231300102"></a>

## Next pages — new_network / 210133203103 / 5

- [voltstack_cluster.site_local_network](data-sources--gcp_vpc_site--reference--group-004.md#canonical-3111110233303120-3020123200210213-3113011221310021-2032102300121231-0001000301332210-1223222212122033-2103323121303210-2122130003210212)
- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-1132222133321101-1232102200201102-2111112102132122-0010023221133123-2333213300031302-1001011022311003-3121103220313010-2103321200303202)

<a id="canonical-3012322021230313-2210131131112301-1230130221210122-1222223012222213-3022231302313322-1233233332220301-2213122013112303-1130223231321321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2320013013021322-0231233131313322-1100012003212312-3232320320300313-3113313000110313-1323111312232221-3310312011112001-2300012120231202"></a>

## voltstack_cluster.site_local_network.new_network_autogenerate — new_network_autogenerate / 131323123111 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-1132222133321101-1232102200201102-2111112102132122-0010023221133123-2333213300031302-1001011022311003-3121103220313010-2103321200303202)
- [Property reference](data-sources--gcp_vpc_site--reference--group-001.md#canonical-0023231003100101-1010210121112300-0303103132302303-0331000323201313-3003311210331022-3000312312212303-0030202320102022-1232013113103201)
- [voltstack_cluster](data-sources--gcp_vpc_site--reference--group-004.md#canonical-2213313332313212-1230123320203012-3120121021331113-3210101220221000-1023022133330131-3023333003003323-2311003313332333-2203303111102220)
- [voltstack_cluster.site_local_network](data-sources--gcp_vpc_site--reference--group-004.md#canonical-3111110233303120-3020123200210213-3113011221310021-2032102300121231-0001000301332210-1223222212122033-2103323121303210-2122130003210212)
- voltstack_cluster.site_local_network.new_network_autogenerate

<a id="canonical-2022331232201000-3020001030220202-2303033011221130-3001103130222330-0322021012200201-0223101311310302-0120200332103301-1122120231211002"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-0023013222222100-1221011211010220-2310112233233233-3113002100011333-1330221000310203-3132301302000330-1032312232100221-3223221310033123"></a>

## Direct properties — new_network_autogenerate / 131323123111 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2301123133000222-3313023201321111-2331323010130000-0023301212310122-3222230112111113-0333113102323132-3321321033313311-3110222031200101"></a>

## Next pages — new_network_autogenerate / 131323123111 / 4

- [voltstack_cluster.site_local_network](data-sources--gcp_vpc_site--reference--group-004.md#canonical-3111110233303120-3020123200210213-3113011221310021-2032102300121231-0001000301332210-1223222212122033-2103323121303210-2122130003210212)
- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-1132222133321101-1232102200201102-2111112102132122-0010023221133123-2333213300031302-1001011022311003-3121103220313010-2103321200303202)

<a id="canonical-0032122331011201-2212323312313011-1322003231020112-2302012332332033-2310121201313332-3003021012113323-3022323300030221-2220332223102330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1300022023013301-0203121300311030-1230113111112211-1221001223310330-3100300231333003-2131222132010200-3122123210312131-3320030103112302"></a>

## voltstack_cluster.site_local_subnet — site_local_subnet / 020321110332 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-1132222133321101-1232102200201102-2111112102132122-0010023221133123-2333213300031302-1001011022311003-3121103220313010-2103321200303202)
- [Property reference](data-sources--gcp_vpc_site--reference--group-001.md#canonical-0023231003100101-1010210121112300-0303103132302303-0331000323201313-3003311210331022-3000312312212303-0030202320102022-1232013113103201)
- [voltstack_cluster](data-sources--gcp_vpc_site--reference--group-004.md#canonical-2213313332313212-1230123320203012-3120121021331113-3210101220221000-1023022133330131-3023333003003323-2311003313332333-2203303111102220)
- voltstack_cluster.site_local_subnet

<a id="canonical-2111002221121010-1221120331301001-3312232203003331-0110113130200222-2212001320120031-3231320000000000-3311223333033313-0232232132003223"></a>

Type: `"single"`. Computed.

Defines choice about GCP VPC network for a view.

Upstream description:

This defines choice about GCP VPC network for a view.

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

<a id="canonical-0013331032102221-3000230312130311-1132212312120123-0122322120010022-2303013232332113-0100120130003330-3033311301130323-0313322112221303"></a>

## Direct properties — site_local_subnet / 020321110332 / 3

- [existing_subnet](data-sources--gcp_vpc_site--reference--group-004.md#canonical-3232023011332203-3320333301223100-1310313101233011-1231312322330121-0300303210122111-1220121102113133-0301231130003200-0013201322013220): complete subsection reference.

- [new_subnet](data-sources--gcp_vpc_site--reference--group-004.md#canonical-2203131332221101-3012122200001332-1103110132223311-3332121212000213-2203211222112203-0332000233011223-3220222113311213-0323111012212012): complete subsection reference.

<a id="canonical-1332313231333320-2211332210331210-1223200032212303-0111223322002123-2132220111330121-2300231031201022-3131330023011132-3232201102002320"></a>

## Next pages — site_local_subnet / 020321110332 / 4

- [voltstack_cluster.site_local_subnet.existing_subnet](data-sources--gcp_vpc_site--reference--group-004.md#canonical-3232023011332203-3320333301223100-1310313101233011-1231312322330121-0300303210122111-1220121102113133-0301231130003200-0013201322013220)
- [voltstack_cluster.site_local_subnet.new_subnet](data-sources--gcp_vpc_site--reference--group-004.md#canonical-2203131332221101-3012122200001332-1103110132223311-3332121212000213-2203211222112203-0332000233011223-3220222113311213-0323111012212012)
- [voltstack_cluster](data-sources--gcp_vpc_site--reference--group-004.md#canonical-2213313332313212-1230123320203012-3120121021331113-3210101220221000-1023022133330131-3023333003003323-2311003313332333-2203303111102220)
- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-1132222133321101-1232102200201102-2111112102132122-0010023221133123-2333213300031302-1001011022311003-3121103220313010-2103321200303202)

<a id="canonical-3232023011332203-3320333301223100-1310313101233011-1231312322330121-0300303210122111-1220121102113133-0301231130003200-0013201322013220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3202103210231301-2300230001102203-1001200220121013-3220332212302032-1120002032321302-2333321321022032-3113032233103002-3030202303122301"></a>

## voltstack_cluster.site_local_subnet.existing_subnet — existing_subnet / 032202201212 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-1132222133321101-1232102200201102-2111112102132122-0010023221133123-2333213300031302-1001011022311003-3121103220313010-2103321200303202)
- [Property reference](data-sources--gcp_vpc_site--reference--group-001.md#canonical-0023231003100101-1010210121112300-0303103132302303-0331000323201313-3003311210331022-3000312312212303-0030202320102022-1232013113103201)
- [voltstack_cluster](data-sources--gcp_vpc_site--reference--group-004.md#canonical-2213313332313212-1230123320203012-3120121021331113-3210101220221000-1023022133330131-3023333003003323-2311003313332333-2203303111102220)
- [voltstack_cluster.site_local_subnet](data-sources--gcp_vpc_site--reference--group-004.md#canonical-0032122331011201-2212323312313011-1322003231020112-2302012332332033-2310121201313332-3003021012113323-3022323300030221-2220332223102330)
- voltstack_cluster.site_local_subnet.existing_subnet

<a id="canonical-1000332030302133-0320033302130330-0323032113230211-0202332321003031-0113323103031301-1230331131330331-0203133301030001-1002303222122211"></a>

Type: `"single"`. Computed.

Configuration parameter for existing subnet.

Upstream description:

Name of existing GCP subnet.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0110011332320232-2213122330032101-2102011332020100-0330013301301012-2023201131212312-1130201222103103-3001123211222013-1130300023331322"></a>

## Direct properties — existing_subnet / 032202201212 / 3

<a id="canonical-1302022331123000-3103223333202121-0133223021121010-1303131222201303-3101332002320213-2100302201132031-1213031323030203-2222130023230111"></a>

<a id="canonical-0313211222113113-0222231013033123-3133030033023003-1132030302002213-0313212011030302-3013210003302110-3323320303222031-3130101312311130"></a>

## subnet_name property — existing_subnet / 032202201212 / 4

Type: `"string"`. Computed.

VPC Subnet Name. Name of your subnet in VPC network.

Upstream description:

Name of your subnet in VPC network.

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

<a id="canonical-0221203033112102-1000222122321213-0133202130111113-3123122133002113-0100203213003102-2120013203133020-3220033303222133-1011233202103011"></a>

## Next pages — existing_subnet / 032202201212 / 5

- [voltstack_cluster.site_local_subnet](data-sources--gcp_vpc_site--reference--group-004.md#canonical-0032122331011201-2212323312313011-1322003231020112-2302012332332033-2310121201313332-3003021012113323-3022323300030221-2220332223102330)
- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-1132222133321101-1232102200201102-2111112102132122-0010023221133123-2333213300031302-1001011022311003-3121103220313010-2103321200303202)

<a id="canonical-2203131332221101-3012122200001332-1103110132223311-3332121212000213-2203211222112203-0332000233011223-3220222113311213-0323111012212012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0321310203320132-3200230303330220-2203303013001110-0131323221231002-0011301300100013-0032321310331332-2011113023200310-3322300032101223"></a>

## voltstack_cluster.site_local_subnet.new_subnet — new_subnet / 331100013230 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-1132222133321101-1232102200201102-2111112102132122-0010023221133123-2333213300031302-1001011022311003-3121103220313010-2103321200303202)
- [Property reference](data-sources--gcp_vpc_site--reference--group-001.md#canonical-0023231003100101-1010210121112300-0303103132302303-0331000323201313-3003311210331022-3000312312212303-0030202320102022-1232013113103201)
- [voltstack_cluster](data-sources--gcp_vpc_site--reference--group-004.md#canonical-2213313332313212-1230123320203012-3120121021331113-3210101220221000-1023022133330131-3023333003003323-2311003313332333-2203303111102220)
- [voltstack_cluster.site_local_subnet](data-sources--gcp_vpc_site--reference--group-004.md#canonical-0032122331011201-2212323312313011-1322003231020112-2302012332332033-2310121201313332-3003021012113323-3022323300030221-2220332223102330)
- voltstack_cluster.site_local_subnet.new_subnet

<a id="canonical-2113130123002210-0322320302231021-1333220320313321-3101332311201222-0003103031013331-2201002332121201-3022123130300000-2220131301231333"></a>

Type: `"single"`. Computed.

GCP subnet parameters Type. Parameters for GCP subnet.

Upstream description:

Parameters for GCP subnet.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0302021300212033-2112102123202210-2313010103100230-1031313320131023-3221122121031012-2133321312220133-3320310001103220-1131203110123133"></a>

## Direct properties — new_subnet / 331100013230 / 3

<a id="canonical-3220301011121031-0312113333312201-1122202100320002-3331202320033303-1112102231101030-1012213223331300-2130210021202333-3110031313101001"></a>

<a id="canonical-3011320331322222-2121012131202333-2333311020311122-3321302320222023-1113011033322303-1101322201100233-2111331010333210-0232113230332210"></a>

## primary_ipv4 property — new_subnet / 331100013230 / 4

Type: `"string"`. Computed.

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

<a id="canonical-1002021223200332-3011123032023223-0020031111012320-1013212313201132-0033220132003102-0222032322012002-2232112313111221-0123323331223311"></a>

<a id="canonical-3221001131203101-0223122010100220-0013132110223330-0023312313130031-3201312222003131-3002120131321030-2010220122020013-3011200003102103"></a>

## subnet_name property — new_subnet / 331100013230 / 5

Type: `"string"`. Computed.

Name of new VPC Subnet, will be autogenerated if empty.

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

<a id="canonical-3001020120333003-0321200030231322-0131303012110230-2210133312200110-3010201021132220-2313023001022200-0311310200201320-2301221303133131"></a>

## Next pages — new_subnet / 331100013230 / 6

- [voltstack_cluster.site_local_subnet](data-sources--gcp_vpc_site--reference--group-004.md#canonical-0032122331011201-2212323312313011-1322003231020112-2302012332332033-2310121201313332-3003021012113323-3022323300030221-2220332223102330)
- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-1132222133321101-1232102200201102-2111112102132122-0010023221133123-2333213300031302-1001011022311003-3121103220313010-2103321200303202)

<a id="canonical-1202113132110030-0331001033001223-1320211110310011-3033232021100220-3110333202300032-3123012232011012-3213130320111011-2032310023331033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3113103110300122-2013203213312132-1222223320213011-3121133101030213-1231102101121012-1113301232000310-1332122112113230-1132103223100112"></a>

## voltstack_cluster.sm_connection_public_ip — sm_connection_public_ip / 113210230122 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-1132222133321101-1232102200201102-2111112102132122-0010023221133123-2333213300031302-1001011022311003-3121103220313010-2103321200303202)
- [Property reference](data-sources--gcp_vpc_site--reference--group-001.md#canonical-0023231003100101-1010210121112300-0303103132302303-0331000323201313-3003311210331022-3000312312212303-0030202320102022-1232013113103201)
- [voltstack_cluster](data-sources--gcp_vpc_site--reference--group-004.md#canonical-2213313332313212-1230123320203012-3120121021331113-3210101220221000-1023022133330131-3023333003003323-2311003313332333-2203303111102220)
- voltstack_cluster.sm_connection_public_ip

<a id="canonical-0133330010111302-3010032013233030-2000121201122022-3113021012113111-1212232311002200-2231323212021002-2120213221311131-2101331320300210"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-0322330222333110-0233110032132332-1032212122131230-2202223131112203-3202232213010303-1000330303302101-3122113031123111-0323131101312123"></a>

## Direct properties — sm_connection_public_ip / 113210230122 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1303213320131200-3112103212101323-0233033113203303-0013333210110120-1022203120111320-2213112002022313-0000002023220310-1010102210002303"></a>

## Next pages — sm_connection_public_ip / 113210230122 / 4

- [voltstack_cluster](data-sources--gcp_vpc_site--reference--group-004.md#canonical-2213313332313212-1230123320203012-3120121021331113-3210101220221000-1023022133330131-3023333003003323-2311003313332333-2203303111102220)
- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-1132222133321101-1232102200201102-2111112102132122-0010023221133123-2333213300031302-1001011022311003-3121103220313010-2103321200303202)

<a id="canonical-0000101133022021-3020001201030002-1203112303012302-1001300121021122-3203111021332232-1233113103000221-2021020022303003-0332301312102233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0321230032301113-0231223201313201-3232022233100100-0020102303133202-0330210201031202-2131201321130110-3030330313321030-2202331221331030"></a>

## voltstack_cluster.sm_connection_pvt_ip — sm_connection_pvt_ip / 133233131311 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-1132222133321101-1232102200201102-2111112102132122-0010023221133123-2333213300031302-1001011022311003-3121103220313010-2103321200303202)
- [Property reference](data-sources--gcp_vpc_site--reference--group-001.md#canonical-0023231003100101-1010210121112300-0303103132302303-0331000323201313-3003311210331022-3000312312212303-0030202320102022-1232013113103201)
- [voltstack_cluster](data-sources--gcp_vpc_site--reference--group-004.md#canonical-2213313332313212-1230123320203012-3120121021331113-3210101220221000-1023022133330131-3023333003003323-2311003313332333-2203303111102220)
- voltstack_cluster.sm_connection_pvt_ip

<a id="canonical-2011231121031122-2312100022223201-3012211312232211-0031320001201210-3013010102221203-2220132210011011-1030010322132032-0223120313102333"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-0020130030311002-3020102022121301-3320211111002301-1333000323121303-2113221302013013-0122020333000120-2311212202320002-1211112222201313"></a>

## Direct properties — sm_connection_pvt_ip / 133233131311 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2020111023301202-0332311132011023-1230230231201031-2102131202120312-1303132302010232-2112130230131132-1231012002310121-2203121203111121"></a>

## Next pages — sm_connection_pvt_ip / 133233131311 / 4

- [voltstack_cluster](data-sources--gcp_vpc_site--reference--group-004.md#canonical-2213313332313212-1230123320203012-3120121021331113-3210101220221000-1023022133330131-3023333003003323-2311003313332333-2203303111102220)
- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-1132222133321101-1232102200201102-2111112102132122-0010023221133123-2333213300031302-1001011022311003-3121103220313010-2103321200303202)

<a id="canonical-1012101330320212-0113133033002223-0210301022311211-3232000103330311-1233032112203232-2330201133011321-3021003023300101-0302002201010202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3311010201231311-1102110211122222-0233110112310320-3303322011210023-1020011321112113-0303200202211022-2000313311011102-1213230102032132"></a>

## voltstack_cluster.storage_class_list — storage_class_list / 230303330120 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-1132222133321101-1232102200201102-2111112102132122-0010023221133123-2333213300031302-1001011022311003-3121103220313010-2103321200303202)
- [Property reference](data-sources--gcp_vpc_site--reference--group-001.md#canonical-0023231003100101-1010210121112300-0303103132302303-0331000323201313-3003311210331022-3000312312212303-0030202320102022-1232013113103201)
- [voltstack_cluster](data-sources--gcp_vpc_site--reference--group-004.md#canonical-2213313332313212-1230123320203012-3120121021331113-3210101220221000-1023022133330131-3023333003003323-2311003313332333-2203303111102220)
- voltstack_cluster.storage_class_list

<a id="canonical-2101000023011303-1112110022001333-2222022122330310-3130002222022303-1333300120100112-3310021222132202-2031122113112202-0330200101301201"></a>

Type: `"single"`. Computed.

Add additional custom storage classes in Kubernetes for this site.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1333133030020113-0022100230020333-2311211203223302-3231332213220333-1313221210212203-0100231102311310-1210112023331113-3321102231133310"></a>

## Direct properties — storage_class_list / 230303330120 / 3

- [storage_classes](data-sources--gcp_vpc_site--reference--group-004.md#canonical-3001123113010200-1223302333001002-3032231012120301-2221303330130122-1021003003102221-2133012031113311-0110301203203330-1200222103200100): complete subsection reference.

<a id="canonical-0022133200232000-3212310331021202-1100033322313013-2133103321302303-3001113330223000-1213300200203012-0020110033303201-0132321201102133"></a>

## Next pages — storage_class_list / 230303330120 / 4

- [voltstack_cluster.storage_class_list.storage_classes](data-sources--gcp_vpc_site--reference--group-004.md#canonical-3001123113010200-1223302333001002-3032231012120301-2221303330130122-1021003003102221-2133012031113311-0110301203203330-1200222103200100)
- [voltstack_cluster](data-sources--gcp_vpc_site--reference--group-004.md#canonical-2213313332313212-1230123320203012-3120121021331113-3210101220221000-1023022133330131-3023333003003323-2311003313332333-2203303111102220)
- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-1132222133321101-1232102200201102-2111112102132122-0010023221133123-2333213300031302-1001011022311003-3121103220313010-2103321200303202)

<a id="canonical-3001123113010200-1223302333001002-3032231012120301-2221303330130122-1021003003102221-2133012031113311-0110301203203330-1200222103200100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1213113102010132-0020321121111030-0120021312322233-3333020321331002-3223033032211032-1200233030213233-2221313232112221-2302002222233231"></a>

## voltstack_cluster.storage_class_list.storage_classes — storage_classes / 302311330130 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-1132222133321101-1232102200201102-2111112102132122-0010023221133123-2333213300031302-1001011022311003-3121103220313010-2103321200303202)
- [Property reference](data-sources--gcp_vpc_site--reference--group-001.md#canonical-0023231003100101-1010210121112300-0303103132302303-0331000323201313-3003311210331022-3000312312212303-0030202320102022-1232013113103201)
- [voltstack_cluster](data-sources--gcp_vpc_site--reference--group-004.md#canonical-2213313332313212-1230123320203012-3120121021331113-3210101220221000-1023022133330131-3023333003003323-2311003313332333-2203303111102220)
- [voltstack_cluster.storage_class_list](data-sources--gcp_vpc_site--reference--group-004.md#canonical-1012101330320212-0113133033002223-0210301022311211-3232000103330311-1233032112203232-2330201133011321-3021003023300101-0302002201010202)
- voltstack_cluster.storage_class_list.storage_classes

<a id="canonical-0311303011230133-3312312201011010-2223020012222300-0313221020120323-1311312330312113-3230100312232030-3313031020312123-2313032132202133"></a>

Type: `"list"`. Computed.

List of Storage Classes. List of custom storage classes.

Upstream description:

List of custom storage classes.

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
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2113320210103122-3232031003112101-0032013121312330-3301022033222100-1122133102223321-2322330311111232-1223211023111013-1110132023201021"></a>

## Direct properties — storage_classes / 302311330130 / 3

<a id="canonical-0222101230201113-0121011300203232-1330310301112230-2122001123011111-0323231212020321-1311203132330003-3321310320232331-1012220222312330"></a>

<a id="canonical-0112013012201120-2031300110110330-3320333013201302-1131302031233321-3201023012231031-1111112311210001-2001230313120032-3312001030030102"></a>

## default_storage_class property — storage_classes / 302311330130 / 4

Type: `"bool"`. Computed.

Make this storage class default storage class for the K8s cluster.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0302111001000032-1022313333131312-3321201121133102-2223321021202222-1202211110032020-0002330101323232-3212103031111000-2030030113230133"></a>

<a id="canonical-2211313321221021-3022313132330001-2132212221030311-2221212100212021-2222223112123320-0032133320322033-0212323231113030-0232222101310201"></a>

## storage_class_name property — storage_classes / 302311330130 / 5

Type: `"string"`. Computed.

Name of the storage class as it will appear in K8s.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "dns-label",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "pattern": "^[a-z]([-a-z0-9]*[a-z0-9])?$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.ves_object_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.ves_object_name": "true"
  }
}
```

<a id="canonical-0132121221112111-0123230323210203-1313211123300131-2020101203111010-3333200333022010-1200223010130002-1121121200332311-0302223100130200"></a>

## Next pages — storage_classes / 302311330130 / 6

- [voltstack_cluster.storage_class_list](data-sources--gcp_vpc_site--reference--group-004.md#canonical-1012101330320212-0113133033002223-0210301022311211-3232000103330311-1233032112203232-2330201133011321-3021003023300101-0302002201010202)
- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-1132222133321101-1232102200201102-2111112102132122-0010023221133123-2333213300031302-1001011022311003-3121103220313010-2103321200303202)

<a id="canonical-0012012033031300-2110002220223133-3303202210233302-1331310110333111-3031100320201031-2130003013122113-0123330322110322-1011031331030122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3211133312322012-3030032121222311-0211002320312030-1323133230030233-2113301103101212-1313221022230321-1230313020113202-0221213202101113"></a>

## waf_signatures — waf_signatures / 300231301202 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-1132222133321101-1232102200201102-2111112102132122-0010023221133123-2333213300031302-1001011022311003-3121103220313010-2103321200303202)
- [Property reference](data-sources--gcp_vpc_site--reference--group-001.md#canonical-0023231003100101-1010210121112300-0303103132302303-0331000323201313-3003311210331022-3000312312212303-0030202320102022-1232013113103201)
- waf_signatures

<a id="canonical-1332200232203221-0021133013311332-0302302031310200-2231200203011332-2120020021102201-1320223013002030-1302022313101313-3023333000322303"></a>

Type: `"single"`. Computed.

Select F5XC WAF Signatures update mode for the site. By default, new signatures will be applied
manually. Refer to release notes for details about available Signatures update modes.

Upstream description:

Select F5XC WAF Signatures update mode for the site. By default, new signatures will be applied
manually. Refer to release notes for details about available Signatures update modes.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-signatures_update_mode_choice": "[\"automatic\",\"manual\"]"
}
```

<a id="canonical-2120100002313133-1110323023322312-1313231121333200-1312031130121111-2322102112010020-3230111001210030-3003011212301102-2000003202030113"></a>

## Direct properties — waf_signatures / 300231301202 / 3

- [automatic](data-sources--gcp_vpc_site--reference--group-005.md#canonical-3012131100333201-3301022220101113-3012012101203300-0222233012233232-1202011133021130-0221332122021222-1020333131132330-3001022302300313): complete subsection reference.

- [manual](data-sources--gcp_vpc_site--reference--group-005.md#canonical-2311002210233302-2000230013003002-0011322313223130-3000201010223131-2131202222330002-2203313301312212-1331301201211030-3032002311031301): complete subsection reference.
