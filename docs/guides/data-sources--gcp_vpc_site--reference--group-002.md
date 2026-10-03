---
page_title: "xcsh_gcp_vpc_site reference"
subcategory: "Infrastructure"
description: "Complete grouped canonical reference for xcsh_gcp_vpc_site reference."
---

# xcsh_gcp_vpc_site reference

<a id="canonical-2131230022001223-3133111212321212-0231230333203223-1333203132003120-0030021031303301-1212010223300001-3230023211110000-1113301023112322"></a>

## node_number property — ingress_egress_gw / 211132000220 / 6

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

- [outside_network](data-sources--gcp_vpc_site--reference--group-002.md#canonical-0110320013312231-3122213130230201-1122220010231112-0112002300113320-1001032001133133-0203110110301223-0312233112010132-1311303220313332): complete subsection reference.

- [outside_static_routes](data-sources--gcp_vpc_site--reference--group-002.md#canonical-2222212203011002-0122100002033220-1003302221310103-3022130333223033-1321222030130012-0211202033021113-0330312012113010-2122030230011200): complete subsection reference.

- [outside_subnet](data-sources--gcp_vpc_site--reference--group-003.md#canonical-1103131103232101-1232010333133003-2331133213033330-2211321130121123-3021333213002012-1020210032022030-2132233233310110-2212132100003010): complete subsection reference.

- [performance_enhancement_mode](data-sources--gcp_vpc_site--reference--group-003.md#canonical-2230020103223132-0122231020000111-0032312233000321-0112002233310200-3033031210100113-1113131330110003-3303202222310312-1321310002122020): complete subsection reference.

- [sm_connection_public_ip](data-sources--gcp_vpc_site--reference--group-003.md#canonical-2213132301100300-3113003012120111-0030302031310011-0332032320103120-2012121231332023-3230330332132030-0220022002001300-3101210011222233): complete subsection reference.

- [sm_connection_pvt_ip](data-sources--gcp_vpc_site--reference--group-003.md#canonical-1230023111002003-0201102303131321-1123233222131100-0101102023001300-3322113011221122-1300131203311020-1321232210303310-1330312231130201): complete subsection reference.

<a id="canonical-2130203220221232-3133323312022313-3021002211213230-1330313011232231-2121112212303122-1130201011222231-2233020212032300-3013323120110032"></a>

## Next pages — ingress_egress_gw / 211132000220 / 7

- [ingress_egress_gw.active_enhanced_firewall_policies](data-sources--gcp_vpc_site--reference--group-002.md#canonical-3233300222201200-1110331210220311-1213131211123213-3330101101013122-0203013123032112-3220323012000312-3023030231002231-3030233101230200)
- [ingress_egress_gw.active_forward_proxy_policies](data-sources--gcp_vpc_site--reference--group-002.md#canonical-3212012103120301-3323313231002021-0233131311130103-1120023320133210-3310332102220022-3330211100123220-3312233112213331-3030020303332303)
- [ingress_egress_gw.active_network_policies](data-sources--gcp_vpc_site--reference--group-002.md#canonical-3222311311231013-0311301030103101-3030230333023110-0213023110303123-1010031333220100-3130011010011023-1103312132000003-0220313311222311)
- [ingress_egress_gw.dc_cluster_group_inside_vn](data-sources--gcp_vpc_site--reference--group-002.md#canonical-2031022112303103-2210230022033131-0001223112320121-0330132121020222-2211002033113200-0020112222110023-2032232310213122-1123323301301001)
- [ingress_egress_gw.dc_cluster_group_outside_vn](data-sources--gcp_vpc_site--reference--group-002.md#canonical-2303202101201031-1220332112033302-2023202030113102-0030131100101331-2311223011220100-3120013003100133-2300110232001312-1031222330302022)
- [ingress_egress_gw.forward_proxy_allow_all](data-sources--gcp_vpc_site--reference--group-002.md#canonical-0032132213332210-2032133310221201-3121310321330300-0322330120331213-3122212000320102-0310002110102331-0211331112222332-0003201333311232)
- [ingress_egress_gw.global_network_list](data-sources--gcp_vpc_site--reference--group-002.md#canonical-1130312120203102-2122110323320300-1102233103300112-2022130100300120-2110022023011312-0123323310213232-0112231123131301-3101313020332221)
- [ingress_egress_gw.inside_network](data-sources--gcp_vpc_site--reference--group-002.md#canonical-1103010332011311-1122120123211002-0313120122102101-1212112111232323-2002213301120131-2031023110130320-2332020000113011-2101203131100023)
- [ingress_egress_gw.inside_static_routes](data-sources--gcp_vpc_site--reference--group-002.md#canonical-3200313310131232-0210110202101312-2021312210102120-1112102012331303-3023321133320033-3321213133200210-3120212332313312-1233223233330302)
- [ingress_egress_gw.inside_subnet](data-sources--gcp_vpc_site--reference--group-002.md#canonical-3203003222300230-2112031310101033-3303130321221222-3231311000312033-1330123121120320-1222230233122203-2132222101023230-1023322133302111)
- [ingress_egress_gw.no_dc_cluster_group](data-sources--gcp_vpc_site--reference--group-002.md#canonical-0202032310301232-3202321000020000-0210202121112010-2120030213110333-3332322130201230-2131302013000120-3331220230110321-1232320231201231)
- [ingress_egress_gw.no_forward_proxy](data-sources--gcp_vpc_site--reference--group-002.md#canonical-0032301122112030-0230102021011322-1331221201233202-2310133101313001-3300103223100011-2101033003202031-3113223310023220-0023332112202012)
- [ingress_egress_gw.no_global_network](data-sources--gcp_vpc_site--reference--group-002.md#canonical-1220311131221300-0120302031132331-1130330201130100-0013112232110113-3210123101120022-2312223312233301-1213032212023333-1001202002130112)
- [ingress_egress_gw.no_inside_static_routes](data-sources--gcp_vpc_site--reference--group-002.md#canonical-3323101233232200-1121011001111111-2122331313313332-1103001110103122-3203022213321323-3303022232101131-3320313203232111-0001201212301111)
- [ingress_egress_gw.no_network_policy](data-sources--gcp_vpc_site--reference--group-002.md#canonical-2321230320203023-2012002031123220-0211203302300320-0331010030233012-0113111022300021-1131233223230022-0203113021211023-0321203131330223)
- [ingress_egress_gw.no_outside_static_routes](data-sources--gcp_vpc_site--reference--group-002.md#canonical-2121122223003320-1111002330210202-1113133312303121-3331201332000322-1210100303320031-3012330310320210-2111211301122021-3130310200112133)
- [ingress_egress_gw.outside_network](data-sources--gcp_vpc_site--reference--group-002.md#canonical-0110320013312231-3122213130230201-1122220010231112-0112002300113320-1001032001133133-0203110110301223-0312233112010132-1311303220313332)
- [ingress_egress_gw.outside_static_routes](data-sources--gcp_vpc_site--reference--group-002.md#canonical-2222212203011002-0122100002033220-1003302221310103-3022130333223033-1321222030130012-0211202033021113-0330312012113010-2122030230011200)
- [ingress_egress_gw.outside_subnet](data-sources--gcp_vpc_site--reference--group-003.md#canonical-1103131103232101-1232010333133003-2331133213033330-2211321130121123-3021333213002012-1020210032022030-2132233233310110-2212132100003010)
- [ingress_egress_gw.performance_enhancement_mode](data-sources--gcp_vpc_site--reference--group-003.md#canonical-2230020103223132-0122231020000111-0032312233000321-0112002233310200-3033031210100113-1113131330110003-3303202222310312-1321310002122020)
- [ingress_egress_gw.sm_connection_public_ip](data-sources--gcp_vpc_site--reference--group-003.md#canonical-2213132301100300-3113003012120111-0030302031310011-0332032320103120-2012121231332023-3230330332132030-0220022002001300-3101210011222233)
- [ingress_egress_gw.sm_connection_pvt_ip](data-sources--gcp_vpc_site--reference--group-003.md#canonical-1230023111002003-0201102303131321-1123233222131100-0101102023001300-3322113011221122-1300131203311020-1321232210303310-1330312231130201)
- [Property reference](data-sources--gcp_vpc_site--reference--group-001.md#canonical-0023231003100101-1010210121112300-0303103132302303-0331000323201313-3003311210331022-3000312312212303-0030202320102022-1232013113103201)
- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-1132222133321101-1232102200201102-2111112102132122-0010023221133123-2333213300031302-1001011022311003-3121103220313010-2103321200303202)

<a id="canonical-3233300222201200-1110331210220311-1213131211123213-3330101101013122-0203013123032112-3220323012000312-3023030231002231-3030233101230200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0320202100303110-0200211123122012-3321010210102210-0010101213001303-3231113331302021-0303102212020032-1021031010311033-1223032110100332"></a>

## ingress_egress_gw.active_enhanced_firewall_policies — active_enhanced_firewall_policies / 013123112231 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-1132222133321101-1232102200201102-2111112102132122-0010023221133123-2333213300031302-1001011022311003-3121103220313010-2103321200303202)
- [Property reference](data-sources--gcp_vpc_site--reference--group-001.md#canonical-0023231003100101-1010210121112300-0303103132302303-0331000323201313-3003311210331022-3000312312212303-0030202320102022-1232013113103201)
- [ingress_egress_gw](data-sources--gcp_vpc_site--reference--group-001.md#canonical-1123313330020103-2220202203231111-2113330331022333-0133213021020321-0303223331023010-1320133103131013-2322232313123200-2330220010011031)
- ingress_egress_gw.active_enhanced_firewall_policies

<a id="canonical-0011220003212110-2112122203023232-2022222130132320-0133233300323002-0133221111331002-3102203113032303-1212301213110033-2020200302131313"></a>

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

<a id="canonical-2332320322112330-1232101010213101-3303131223133232-3110311112021320-2012331020020302-3201000010322110-0223113102330110-3312302332011133"></a>

## Direct properties — active_enhanced_firewall_policies / 013123112231 / 3

- [enhanced_firewall_policies](data-sources--gcp_vpc_site--reference--group-002.md#canonical-1230003100221012-0323023201302011-2002302110110303-2021021233111103-1321230100031021-1231031030022113-0011010213330102-3000212102031332): complete subsection reference.

<a id="canonical-1130303033213010-3230001102323312-2213200103233321-0310222323002211-2122120230332032-3101003230201311-2021221201030012-0113120033130123"></a>

## Next pages — active_enhanced_firewall_policies / 013123112231 / 4

- [ingress_egress_gw.active_enhanced_firewall_policies.enhanced_firewall_policies](data-sources--gcp_vpc_site--reference--group-002.md#canonical-1230003100221012-0323023201302011-2002302110110303-2021021233111103-1321230100031021-1231031030022113-0011010213330102-3000212102031332)
- [ingress_egress_gw](data-sources--gcp_vpc_site--reference--group-001.md#canonical-1123313330020103-2220202203231111-2113330331022333-0133213021020321-0303223331023010-1320133103131013-2322232313123200-2330220010011031)
- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-1132222133321101-1232102200201102-2111112102132122-0010023221133123-2333213300031302-1001011022311003-3121103220313010-2103321200303202)

<a id="canonical-1230003100221012-0323023201302011-2002302110110303-2021021233111103-1321230100031021-1231031030022113-0011010213330102-3000212102031332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1011011021223210-2310011320220313-0220212133211100-1302232212300132-1201030122312223-0112222301210100-2213130001322321-2211100000020030"></a>

## ingress_egress_gw.active_enhanced_firewall_policies.enhanced_firewall_policies — enhanced_firewall_policies / 322132123023 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-1132222133321101-1232102200201102-2111112102132122-0010023221133123-2333213300031302-1001011022311003-3121103220313010-2103321200303202)
- [Property reference](data-sources--gcp_vpc_site--reference--group-001.md#canonical-0023231003100101-1010210121112300-0303103132302303-0331000323201313-3003311210331022-3000312312212303-0030202320102022-1232013113103201)
- [ingress_egress_gw](data-sources--gcp_vpc_site--reference--group-001.md#canonical-1123313330020103-2220202203231111-2113330331022333-0133213021020321-0303223331023010-1320133103131013-2322232313123200-2330220010011031)
- [ingress_egress_gw.active_enhanced_firewall_policies](data-sources--gcp_vpc_site--reference--group-002.md#canonical-3233300222201200-1110331210220311-1213131211123213-3330101101013122-0203013123032112-3220323012000312-3023030231002231-3030233101230200)
- ingress_egress_gw.active_enhanced_firewall_policies.enhanced_firewall_policies

<a id="canonical-3202013331021013-2210210233100202-2120221222333113-2031031131203300-0112222212312130-3133013103130010-2310132100123120-0202302112013323"></a>

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-0313011201100322-1322233122310231-1221103220320000-2201131321020001-3003132210131023-1233012303313131-2200123130130222-1111310123333131"></a>

## Direct properties — enhanced_firewall_policies / 322132123023 / 3

<a id="canonical-2012031002300210-2031303300132110-2023301013323112-2001322212010320-0131123203022131-3103012011101233-2212132323201102-3222212331330032"></a>

<a id="canonical-0021101333022002-2111313202321000-3310230200131221-2221323103122232-0002311101132230-0212210111000000-0111320010023211-1310021013011312"></a>

## name property — enhanced_firewall_policies / 322132123023 / 4

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

<a id="canonical-2100112101233103-1323232110213301-3120332031113101-2112112123100222-2330220213032133-2003211233022220-3232101213330003-0113103330220020"></a>

<a id="canonical-0022130222213033-0003222212130310-2300200200011231-3011020320000100-2102222032030223-0312223322333123-0031321213022201-0220202000202303"></a>

## namespace property — enhanced_firewall_policies / 322132123023 / 5

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

<a id="canonical-2220031031033123-2122213333313302-2133212211000201-1020233020321123-2300332201300100-3110113122303100-0320133032310110-2023033223112202"></a>

<a id="canonical-2031033113100123-3230030023211210-3031230100220201-0110332112331031-1303310220011233-0110332123023101-2122231020120131-3230032232131233"></a>

## tenant property — enhanced_firewall_policies / 322132123023 / 6

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

<a id="canonical-1102303110023111-1302230012231311-0332203310200120-1121202012030113-1020223110330230-1031113032300322-2113000110013113-0311233212001310"></a>

## Next pages — enhanced_firewall_policies / 322132123023 / 7

- [ingress_egress_gw.active_enhanced_firewall_policies](data-sources--gcp_vpc_site--reference--group-002.md#canonical-3233300222201200-1110331210220311-1213131211123213-3330101101013122-0203013123032112-3220323012000312-3023030231002231-3030233101230200)
- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-1132222133321101-1232102200201102-2111112102132122-0010023221133123-2333213300031302-1001011022311003-3121103220313010-2103321200303202)

<a id="canonical-3212012103120301-3323313231002021-0233131311130103-1120023320133210-3310332102220022-3330211100123220-3312233112213331-3030020303332303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1031002302112232-2222121311010221-0322001321301031-1100222233022013-3112111201112100-0100220213320210-0202232021012302-2302311232020003"></a>

## ingress_egress_gw.active_forward_proxy_policies — active_forward_proxy_policies / 230213313221 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-1132222133321101-1232102200201102-2111112102132122-0010023221133123-2333213300031302-1001011022311003-3121103220313010-2103321200303202)
- [Property reference](data-sources--gcp_vpc_site--reference--group-001.md#canonical-0023231003100101-1010210121112300-0303103132302303-0331000323201313-3003311210331022-3000312312212303-0030202320102022-1232013113103201)
- [ingress_egress_gw](data-sources--gcp_vpc_site--reference--group-001.md#canonical-1123313330020103-2220202203231111-2113330331022333-0133213021020321-0303223331023010-1320133103131013-2322232313123200-2330220010011031)
- ingress_egress_gw.active_forward_proxy_policies

<a id="canonical-1221110013213020-2110232331123000-2112302122103212-0312200021222233-3322022200111321-0220112103210200-0000332211101133-0030012100120023"></a>

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

<a id="canonical-0000311103311003-3033013212210000-1130122103202030-1031212103212023-3323330101131331-1221322013001213-2123331301003013-2013312311320330"></a>

## Direct properties — active_forward_proxy_policies / 230213313221 / 3

- [forward_proxy_policies](data-sources--gcp_vpc_site--reference--group-002.md#canonical-1033312320221222-0333331020123130-2112002311301020-2032002301313220-2301110023302132-1001002200213233-3211230012013122-3101110213313032): complete subsection reference.

<a id="canonical-0003102003332003-3001101020231310-3001301231221012-3031320003122020-2311003011132203-0001220202313313-1122030303301232-0112131301032223"></a>

## Next pages — active_forward_proxy_policies / 230213313221 / 4

- [ingress_egress_gw.active_forward_proxy_policies.forward_proxy_policies](data-sources--gcp_vpc_site--reference--group-002.md#canonical-1033312320221222-0333331020123130-2112002311301020-2032002301313220-2301110023302132-1001002200213233-3211230012013122-3101110213313032)
- [ingress_egress_gw](data-sources--gcp_vpc_site--reference--group-001.md#canonical-1123313330020103-2220202203231111-2113330331022333-0133213021020321-0303223331023010-1320133103131013-2322232313123200-2330220010011031)
- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-1132222133321101-1232102200201102-2111112102132122-0010023221133123-2333213300031302-1001011022311003-3121103220313010-2103321200303202)

<a id="canonical-1033312320221222-0333331020123130-2112002311301020-2032002301313220-2301110023302132-1001002200213233-3211230012013122-3101110213313032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2213010201131100-3132213131022130-0222333112323002-1022113300332012-0102133003003202-3322112223220120-2133023221222300-0213111222113012"></a>

## ingress_egress_gw.active_forward_proxy_policies.forward_proxy_policies — forward_proxy_policies / 210312320001 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-1132222133321101-1232102200201102-2111112102132122-0010023221133123-2333213300031302-1001011022311003-3121103220313010-2103321200303202)
- [Property reference](data-sources--gcp_vpc_site--reference--group-001.md#canonical-0023231003100101-1010210121112300-0303103132302303-0331000323201313-3003311210331022-3000312312212303-0030202320102022-1232013113103201)
- [ingress_egress_gw](data-sources--gcp_vpc_site--reference--group-001.md#canonical-1123313330020103-2220202203231111-2113330331022333-0133213021020321-0303223331023010-1320133103131013-2322232313123200-2330220010011031)
- [ingress_egress_gw.active_forward_proxy_policies](data-sources--gcp_vpc_site--reference--group-002.md#canonical-3212012103120301-3323313231002021-0233131311130103-1120023320133210-3310332102220022-3330211100123220-3312233112213331-3030020303332303)
- ingress_egress_gw.active_forward_proxy_policies.forward_proxy_policies

<a id="canonical-2310112330031102-3120302213200130-3101310311011333-3301213232212003-3013223201110223-0113102100030301-3000112102110313-2303332002322022"></a>

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-3132333311130211-3013230012303232-1213010010331010-0123001010221003-3101201313032013-1233301312220000-0313132230013100-3210131110030220"></a>

## Direct properties — forward_proxy_policies / 210312320001 / 3

<a id="canonical-3132102033133330-2002233111300020-0221112103311230-2003223321113300-0331213322022011-0331030012133123-3032032012120000-3230312312322211"></a>

<a id="canonical-2303111323300301-2101210320222033-0300001131120233-3112120031301310-1132231011012312-1333322133131002-0312222110210112-2111322130112023"></a>

## name property — forward_proxy_policies / 210312320001 / 4

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

<a id="canonical-2110100003323020-3212120122002121-0330102020030021-2303211303201310-1031220233233331-0031133313021302-0012323033231111-3120130220011000"></a>

<a id="canonical-2231021202330002-1333030001110321-2123002133033032-0033122033121313-0121112100302331-2221130201113210-2303122331132002-1321112300223031"></a>

## namespace property — forward_proxy_policies / 210312320001 / 5

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

<a id="canonical-2203131330033320-0300213023321223-0012312110331301-1121231031011003-3031121011323203-1321133230303100-2331001110021331-0320103121033330"></a>

<a id="canonical-3013120123311330-3302211230100300-2102010210033302-1203330320200003-3200020023300300-1110033002101020-0222113122233231-2000213120310322"></a>

## tenant property — forward_proxy_policies / 210312320001 / 6

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

<a id="canonical-3231111212121132-3101100230311201-3320232301322031-3022021122212211-3023312111311203-2033113331100223-0100130123212033-1010133202031120"></a>

## Next pages — forward_proxy_policies / 210312320001 / 7

- [ingress_egress_gw.active_forward_proxy_policies](data-sources--gcp_vpc_site--reference--group-002.md#canonical-3212012103120301-3323313231002021-0233131311130103-1120023320133210-3310332102220022-3330211100123220-3312233112213331-3030020303332303)
- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-1132222133321101-1232102200201102-2111112102132122-0010023221133123-2333213300031302-1001011022311003-3121103220313010-2103321200303202)

<a id="canonical-3222311311231013-0311301030103101-3030230333023110-0213023110303123-1010031333220100-3130011010011023-1103312132000003-0220313311222311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1121123303022320-0013203001001012-2231000233200100-0320322121011101-0230330112010230-1303221010221203-3301033031233302-2010022021121111"></a>

## ingress_egress_gw.active_network_policies — active_network_policies / 333003230033 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-1132222133321101-1232102200201102-2111112102132122-0010023221133123-2333213300031302-1001011022311003-3121103220313010-2103321200303202)
- [Property reference](data-sources--gcp_vpc_site--reference--group-001.md#canonical-0023231003100101-1010210121112300-0303103132302303-0331000323201313-3003311210331022-3000312312212303-0030202320102022-1232013113103201)
- [ingress_egress_gw](data-sources--gcp_vpc_site--reference--group-001.md#canonical-1123313330020103-2220202203231111-2113330331022333-0133213021020321-0303223331023010-1320133103131013-2322232313123200-2330220010011031)
- ingress_egress_gw.active_network_policies

<a id="canonical-2231111100313323-2313020120032212-0033131221311330-0123102201101211-2313023033032222-2003333111332211-1223232103100011-3110022122121322"></a>

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

<a id="canonical-2231302311313121-2221022120213111-0200211200110302-1322002032300302-3102310311303330-2223012331121030-1230131201033213-3123201122030200"></a>

## Direct properties — active_network_policies / 333003230033 / 3

- [network_policies](data-sources--gcp_vpc_site--reference--group-002.md#canonical-0031022201122032-2121320301130301-2020012120122023-3212230032021201-1100322123020233-2120221222231101-0323220313131020-0020303020012031): complete subsection reference.

<a id="canonical-0032310330312331-1123233221131030-3002123220311232-1013000332002121-1103223121320020-2131133211000232-3211101003031330-3333103130122112"></a>

## Next pages — active_network_policies / 333003230033 / 4

- [ingress_egress_gw.active_network_policies.network_policies](data-sources--gcp_vpc_site--reference--group-002.md#canonical-0031022201122032-2121320301130301-2020012120122023-3212230032021201-1100322123020233-2120221222231101-0323220313131020-0020303020012031)
- [ingress_egress_gw](data-sources--gcp_vpc_site--reference--group-001.md#canonical-1123313330020103-2220202203231111-2113330331022333-0133213021020321-0303223331023010-1320133103131013-2322232313123200-2330220010011031)
- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-1132222133321101-1232102200201102-2111112102132122-0010023221133123-2333213300031302-1001011022311003-3121103220313010-2103321200303202)

<a id="canonical-0031022201122032-2121320301130301-2020012120122023-3212230032021201-1100322123020233-2120221222231101-0323220313131020-0020303020012031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2013120020300222-1023212330301133-3113203022032020-1333112231133020-0033333101123300-3023112031000202-2021132031232011-3133112312031012"></a>

## ingress_egress_gw.active_network_policies.network_policies — network_policies / 320012111121 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-1132222133321101-1232102200201102-2111112102132122-0010023221133123-2333213300031302-1001011022311003-3121103220313010-2103321200303202)
- [Property reference](data-sources--gcp_vpc_site--reference--group-001.md#canonical-0023231003100101-1010210121112300-0303103132302303-0331000323201313-3003311210331022-3000312312212303-0030202320102022-1232013113103201)
- [ingress_egress_gw](data-sources--gcp_vpc_site--reference--group-001.md#canonical-1123313330020103-2220202203231111-2113330331022333-0133213021020321-0303223331023010-1320133103131013-2322232313123200-2330220010011031)
- [ingress_egress_gw.active_network_policies](data-sources--gcp_vpc_site--reference--group-002.md#canonical-3222311311231013-0311301030103101-3030230333023110-0213023110303123-1010031333220100-3130011010011023-1103312132000003-0220313311222311)
- ingress_egress_gw.active_network_policies.network_policies

<a id="canonical-1213002002220332-0131122033012131-1011123032302203-2110003020233330-1032200323313011-3320113013002023-1123030010300222-0210312032132033"></a>

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-3011320212030113-2130323223031133-0000212202120302-3311333022301230-0231022213233013-3001301100311321-0212133030311033-1331031011130201"></a>

## Direct properties — network_policies / 320012111121 / 3

<a id="canonical-3223231131301220-2022302332131102-2322231132232110-0233030331123020-0022333033133031-2103322002001102-3103330122230223-2003010000031132"></a>

<a id="canonical-3311032300311210-1022231121121230-2211122120302311-0101323123220200-0002313223012233-3310202310332121-0000123333123131-1330121200033132"></a>

## name property — network_policies / 320012111121 / 4

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

<a id="canonical-2321022013003021-2101323321132213-3001010011302012-2211312330120020-0001010002210310-0203322321211020-3333210130021030-0213310311312132"></a>

<a id="canonical-3312022113102333-0030300233201221-3113331200010022-0303013200210010-1310113230301020-2112012030013301-1213222232203103-1111131313222222"></a>

## namespace property — network_policies / 320012111121 / 5

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

<a id="canonical-1022220131330112-3203132210311030-0112212100220233-1131111200301313-1013230103133302-2011000312232312-0133002330130133-1102201323221011"></a>

<a id="canonical-3323020300132111-3103230020220320-0300331331112201-2110031021022113-2133133203232003-2031113220330100-0321220322232001-0110212121021120"></a>

## tenant property — network_policies / 320012111121 / 6

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

<a id="canonical-1302102310100122-1200220110312133-3310303201311321-0112222203032113-0331212021003001-2223310333113311-2111131200021312-1333103323102002"></a>

## Next pages — network_policies / 320012111121 / 7

- [ingress_egress_gw.active_network_policies](data-sources--gcp_vpc_site--reference--group-002.md#canonical-3222311311231013-0311301030103101-3030230333023110-0213023110303123-1010031333220100-3130011010011023-1103312132000003-0220313311222311)
- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-1132222133321101-1232102200201102-2111112102132122-0010023221133123-2333213300031302-1001011022311003-3121103220313010-2103321200303202)

<a id="canonical-2031022112303103-2210230022033131-0001223112320121-0330132121020222-2211002033113200-0020112222110023-2032232310213122-1123323301301001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3120213101330023-2011302221233311-3203032300231201-2203022333231302-2233202323120110-1320320221033221-2201231230021022-0003323200301031"></a>

## ingress_egress_gw.dc_cluster_group_inside_vn — dc_cluster_group_inside_vn / 330203230322 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-1132222133321101-1232102200201102-2111112102132122-0010023221133123-2333213300031302-1001011022311003-3121103220313010-2103321200303202)
- [Property reference](data-sources--gcp_vpc_site--reference--group-001.md#canonical-0023231003100101-1010210121112300-0303103132302303-0331000323201313-3003311210331022-3000312312212303-0030202320102022-1232013113103201)
- [ingress_egress_gw](data-sources--gcp_vpc_site--reference--group-001.md#canonical-1123313330020103-2220202203231111-2113330331022333-0133213021020321-0303223331023010-1320133103131013-2322232313123200-2330220010011031)
- ingress_egress_gw.dc_cluster_group_inside_vn

<a id="canonical-0122210210301102-0310313022012330-1313320110222103-3221031213200221-0320012323023032-0332202210303322-0210202213011002-3010232230210200"></a>

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

<a id="canonical-3303202313313031-2003320323322302-2113122100323231-0000223310300031-3330333131130303-3000131301021212-3223303332011033-3213111302320101"></a>

## Direct properties — dc_cluster_group_inside_vn / 330203230322 / 3

<a id="canonical-3011023202323032-0313220122123310-2123302012222100-2113103002231032-3221301322212223-0213330102123032-1313031232100203-2022023102122100"></a>

<a id="canonical-1000203301302022-3230101112003010-0232113223020211-0320033211012021-2122200022320012-0303000103212330-2103233231221033-2031320021103021"></a>

## name property — dc_cluster_group_inside_vn / 330203230322 / 4

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

<a id="canonical-1102201102033212-3223223012120332-1313122001222220-2103322100103202-0213123003213333-0331011211332230-3300133000002020-1102332131323333"></a>

<a id="canonical-0202310322002332-3123003202021020-1232013303231022-3200121211130201-3022323231001203-3202232030213202-1223020111013000-3013031333223123"></a>

## namespace property — dc_cluster_group_inside_vn / 330203230322 / 5

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

<a id="canonical-2333220111320323-1333101320023110-0000003300310213-0102223321320123-2132020131203013-2132022122301123-2333300003110332-1322130301022100"></a>

<a id="canonical-1120131302301303-3011201021022010-2202003001011031-1123103202132110-3110333303332021-2111023113331201-3031113022023100-0200303203200300"></a>

## tenant property — dc_cluster_group_inside_vn / 330203230322 / 6

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

<a id="canonical-2033001133021321-1131031301120202-2121111210312133-0311312321303321-2213230221001122-1201013332020001-3233100123130102-0112310022303010"></a>

## Next pages — dc_cluster_group_inside_vn / 330203230322 / 7

- [ingress_egress_gw](data-sources--gcp_vpc_site--reference--group-001.md#canonical-1123313330020103-2220202203231111-2113330331022333-0133213021020321-0303223331023010-1320133103131013-2322232313123200-2330220010011031)
- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-1132222133321101-1232102200201102-2111112102132122-0010023221133123-2333213300031302-1001011022311003-3121103220313010-2103321200303202)

<a id="canonical-2303202101201031-1220332112033302-2023202030113102-0030131100101331-2311223011220100-3120013003100133-2300110232001312-1031222330302022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2020123011223102-3032230203022203-2221303232231310-0202332032103221-0301211120310122-3112222203103103-1032302232002121-0013233210001121"></a>

## ingress_egress_gw.dc_cluster_group_outside_vn — dc_cluster_group_outside_vn / 211002102000 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-1132222133321101-1232102200201102-2111112102132122-0010023221133123-2333213300031302-1001011022311003-3121103220313010-2103321200303202)
- [Property reference](data-sources--gcp_vpc_site--reference--group-001.md#canonical-0023231003100101-1010210121112300-0303103132302303-0331000323201313-3003311210331022-3000312312212303-0030202320102022-1232013113103201)
- [ingress_egress_gw](data-sources--gcp_vpc_site--reference--group-001.md#canonical-1123313330020103-2220202203231111-2113330331022333-0133213021020321-0303223331023010-1320133103131013-2322232313123200-2330220010011031)
- ingress_egress_gw.dc_cluster_group_outside_vn

<a id="canonical-2032030201233212-1133310201222023-1100310323232103-1133131022322332-1133001212121333-2212221220321111-2202310300220313-0022203101332111"></a>

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

<a id="canonical-0332023203230122-2022331212002000-2320023000100200-2101132103111321-2101331210302031-1222100112130333-3133303330220111-1312133311032322"></a>

## Direct properties — dc_cluster_group_outside_vn / 211002102000 / 3

<a id="canonical-0231132111013301-2102211323133112-2132102203030021-1021130110312100-1121023203023333-3021230223311233-1012013221022310-3030210210032210"></a>

<a id="canonical-0233300012020312-0120111201201123-3133010220112200-0321133310221133-2021123110122100-3320303021210003-1210212002113102-1300123013333031"></a>

## name property — dc_cluster_group_outside_vn / 211002102000 / 4

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

<a id="canonical-2223212001031020-3323313102012312-0020032332331332-2132131202122102-3232311112001110-1011232312213202-3130021213221112-1210000132203202"></a>

<a id="canonical-2001022313311021-2233101003113331-0101000203302332-0130133020123001-1301210320030321-3130320030032310-1003032011233012-0013232221120312"></a>

## namespace property — dc_cluster_group_outside_vn / 211002102000 / 5

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

<a id="canonical-1332130011201200-1322110302003323-0101231311323010-3331111102103013-3302310132301112-1331322310332022-1130313002210201-0023001112032000"></a>

<a id="canonical-2321213020000232-3113023120303032-1113002103211023-2323312121123312-3032220101011110-1030131021102122-3300031113003121-1032033213013323"></a>

## tenant property — dc_cluster_group_outside_vn / 211002102000 / 6

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

<a id="canonical-3330000303220210-2010313020110300-1231310011121312-1001132120320133-3210033221131312-0100032023211033-0112201202321023-1233132013122201"></a>

## Next pages — dc_cluster_group_outside_vn / 211002102000 / 7

- [ingress_egress_gw](data-sources--gcp_vpc_site--reference--group-001.md#canonical-1123313330020103-2220202203231111-2113330331022333-0133213021020321-0303223331023010-1320133103131013-2322232313123200-2330220010011031)
- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-1132222133321101-1232102200201102-2111112102132122-0010023221133123-2333213300031302-1001011022311003-3121103220313010-2103321200303202)

<a id="canonical-0032132213332210-2032133310221201-3121310321330300-0322330120331213-3122212000320102-0310002110102331-0211331112222332-0003201333311232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2313102121322313-0311301321020311-3222233003212210-2023311033020021-2323311301120113-1221033320012303-1202203011120112-2030123123011010"></a>

## ingress_egress_gw.forward_proxy_allow_all — forward_proxy_allow_all / 223022023213 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-1132222133321101-1232102200201102-2111112102132122-0010023221133123-2333213300031302-1001011022311003-3121103220313010-2103321200303202)
- [Property reference](data-sources--gcp_vpc_site--reference--group-001.md#canonical-0023231003100101-1010210121112300-0303103132302303-0331000323201313-3003311210331022-3000312312212303-0030202320102022-1232013113103201)
- [ingress_egress_gw](data-sources--gcp_vpc_site--reference--group-001.md#canonical-1123313330020103-2220202203231111-2113330331022333-0133213021020321-0303223331023010-1320133103131013-2322232313123200-2330220010011031)
- ingress_egress_gw.forward_proxy_allow_all

<a id="canonical-1012223021333013-3033013201110301-3233122303221303-3110201233221101-1103310013333123-3323000200013031-2222311220331211-1223210300320132"></a>

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

<a id="canonical-0111011132300200-1201322202101313-3030303323332323-1320110321023110-0121211230123331-0211230301122322-2230310301031302-2222211301010001"></a>

## Direct properties — forward_proxy_allow_all / 223022023213 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2321002121320330-2030001103110303-3133331101130131-0132112102003322-3312101123223323-3022033330320022-2302322321202032-0003002201120303"></a>

## Next pages — forward_proxy_allow_all / 223022023213 / 4

- [ingress_egress_gw](data-sources--gcp_vpc_site--reference--group-001.md#canonical-1123313330020103-2220202203231111-2113330331022333-0133213021020321-0303223331023010-1320133103131013-2322232313123200-2330220010011031)
- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-1132222133321101-1232102200201102-2111112102132122-0010023221133123-2333213300031302-1001011022311003-3121103220313010-2103321200303202)

<a id="canonical-1130312120203102-2122110323320300-1102233103300112-2022130100300120-2110022023011312-0123323310213232-0112231123131301-3101313020332221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1223032023103102-0000100300213230-1023113133332133-1321322221112332-1013211201233212-3103212202020021-0221103203322321-1020013123333303"></a>

## ingress_egress_gw.global_network_list — global_network_list / 110003023120 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-1132222133321101-1232102200201102-2111112102132122-0010023221133123-2333213300031302-1001011022311003-3121103220313010-2103321200303202)
- [Property reference](data-sources--gcp_vpc_site--reference--group-001.md#canonical-0023231003100101-1010210121112300-0303103132302303-0331000323201313-3003311210331022-3000312312212303-0030202320102022-1232013113103201)
- [ingress_egress_gw](data-sources--gcp_vpc_site--reference--group-001.md#canonical-1123313330020103-2220202203231111-2113330331022333-0133213021020321-0303223331023010-1320133103131013-2322232313123200-2330220010011031)
- ingress_egress_gw.global_network_list

<a id="canonical-0132031033010220-0022031212001320-1132002213322013-2133020323001310-3213012203222220-3133032312103301-1223321221112202-1310123112020132"></a>

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

<a id="canonical-3202231303213203-1313120333210101-0323303003130120-2122223013133213-3223210302020121-2103000312231203-2302121102220103-3011103331002220"></a>

## Direct properties — global_network_list / 110003023120 / 3

- [global_network_connections](data-sources--gcp_vpc_site--reference--group-002.md#canonical-1002112130221021-3330011001330322-1320022131222221-3003201102130100-3020213123210002-0101032321103312-2112023202211223-1310132021211212): complete subsection reference.

<a id="canonical-0223220123031100-0222122223111301-0300312302231333-3212013020221231-1200130001212201-1013010203031313-0302213213001211-1330101000312123"></a>

## Next pages — global_network_list / 110003023120 / 4

- [ingress_egress_gw.global_network_list.global_network_connections](data-sources--gcp_vpc_site--reference--group-002.md#canonical-1002112130221021-3330011001330322-1320022131222221-3003201102130100-3020213123210002-0101032321103312-2112023202211223-1310132021211212)
- [ingress_egress_gw](data-sources--gcp_vpc_site--reference--group-001.md#canonical-1123313330020103-2220202203231111-2113330331022333-0133213021020321-0303223331023010-1320133103131013-2322232313123200-2330220010011031)
- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-1132222133321101-1232102200201102-2111112102132122-0010023221133123-2333213300031302-1001011022311003-3121103220313010-2103321200303202)

<a id="canonical-1002112130221021-3330011001330322-1320022131222221-3003201102130100-3020213123210002-0101032321103312-2112023202211223-1310132021211212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2313012033303111-0030122003203123-1002303131020130-3331111102212023-1032303200112303-2313010210213031-1113031202003022-3131100113223131"></a>

## ingress_egress_gw.global_network_list.global_network_connections — global_network_connections / 300312022113 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-1132222133321101-1232102200201102-2111112102132122-0010023221133123-2333213300031302-1001011022311003-3121103220313010-2103321200303202)
- [Property reference](data-sources--gcp_vpc_site--reference--group-001.md#canonical-0023231003100101-1010210121112300-0303103132302303-0331000323201313-3003311210331022-3000312312212303-0030202320102022-1232013113103201)
- [ingress_egress_gw](data-sources--gcp_vpc_site--reference--group-001.md#canonical-1123313330020103-2220202203231111-2113330331022333-0133213021020321-0303223331023010-1320133103131013-2322232313123200-2330220010011031)
- [ingress_egress_gw.global_network_list](data-sources--gcp_vpc_site--reference--group-002.md#canonical-1130312120203102-2122110323320300-1102233103300112-2022130100300120-2110022023011312-0123323310213232-0112231123131301-3101313020332221)
- ingress_egress_gw.global_network_list.global_network_connections

<a id="canonical-3222122122102121-2301132320202131-1011122131303123-0011210202321203-1100003110013211-1110233111320101-1100231220210311-1031102111010003"></a>

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-1122012202211121-1110212020300303-0332300002033212-2031031231033312-3111303210131210-0302120330012110-2111311103000201-0000120012113011"></a>

## Direct properties — global_network_connections / 300312022113 / 3

- [sli_to_global_dr](data-sources--gcp_vpc_site--reference--group-002.md#canonical-3330032003012321-1120210313032212-2312133112233103-3323330212212123-0323133022212321-0132112322301201-0023213202312310-3222103003101221): complete subsection reference.

- [slo_to_global_dr](data-sources--gcp_vpc_site--reference--group-002.md#canonical-3132320012033303-1003211031032332-1312303212112333-2023222030303131-3200013303311031-3300132001132032-0203020233003102-1001212030001002): complete subsection reference.

<a id="canonical-0112123230303103-1211101323333233-3232201000031121-1322010302301030-2321130211220100-1213011212130221-2122301113211213-3211232323102330"></a>

## Next pages — global_network_connections / 300312022113 / 4

- [ingress_egress_gw.global_network_list.global_network_connections.sli_to_global_dr](data-sources--gcp_vpc_site--reference--group-002.md#canonical-3330032003012321-1120210313032212-2312133112233103-3323330212212123-0323133022212321-0132112322301201-0023213202312310-3222103003101221)
- [ingress_egress_gw.global_network_list.global_network_connections.slo_to_global_dr](data-sources--gcp_vpc_site--reference--group-002.md#canonical-3132320012033303-1003211031032332-1312303212112333-2023222030303131-3200013303311031-3300132001132032-0203020233003102-1001212030001002)
- [ingress_egress_gw.global_network_list](data-sources--gcp_vpc_site--reference--group-002.md#canonical-1130312120203102-2122110323320300-1102233103300112-2022130100300120-2110022023011312-0123323310213232-0112231123131301-3101313020332221)
- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-1132222133321101-1232102200201102-2111112102132122-0010023221133123-2333213300031302-1001011022311003-3121103220313010-2103321200303202)

<a id="canonical-3330032003012321-1120210313032212-2312133112233103-3323330212212123-0323133022212321-0132112322301201-0023213202312310-3222103003101221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1123332133232002-3120031100000301-3301312013201330-2032030131011310-3332103213121313-1002023233322213-3000113121313110-0312222100132020"></a>

## ingress_egress_gw.global_network_list.global_network_connections.sli_to_global_dr — sli_to_global_dr / 101322031102 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-1132222133321101-1232102200201102-2111112102132122-0010023221133123-2333213300031302-1001011022311003-3121103220313010-2103321200303202)
- [Property reference](data-sources--gcp_vpc_site--reference--group-001.md#canonical-0023231003100101-1010210121112300-0303103132302303-0331000323201313-3003311210331022-3000312312212303-0030202320102022-1232013113103201)
- [ingress_egress_gw](data-sources--gcp_vpc_site--reference--group-001.md#canonical-1123313330020103-2220202203231111-2113330331022333-0133213021020321-0303223331023010-1320133103131013-2322232313123200-2330220010011031)
- [ingress_egress_gw.global_network_list](data-sources--gcp_vpc_site--reference--group-002.md#canonical-1130312120203102-2122110323320300-1102233103300112-2022130100300120-2110022023011312-0123323310213232-0112231123131301-3101313020332221)
- [ingress_egress_gw.global_network_list.global_network_connections](data-sources--gcp_vpc_site--reference--group-002.md#canonical-1002112130221021-3330011001330322-1320022131222221-3003201102130100-3020213123210002-0101032321103312-2112023202211223-1310132021211212)
- ingress_egress_gw.global_network_list.global_network_connections.sli_to_global_dr

<a id="canonical-2331003130102333-2301322331003201-0112213010332233-3203131030003111-3210232210301311-0110112112013031-0201022022112200-0021222300301221"></a>

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

<a id="canonical-2303322320100302-1131210033011321-0030210213101320-3101133211302323-2122020222100112-0233312100101202-3312013213011001-1321203313130211"></a>

## Direct properties — sli_to_global_dr / 101322031102 / 3

- [global_vn](data-sources--gcp_vpc_site--reference--group-002.md#canonical-1032131322013321-2313331131010322-1030123031020233-3112221000202120-3001213320012301-2331220002033000-0202200033111310-3021013121220331): complete subsection reference.

<a id="canonical-0022302323122033-0233002302202002-1001321333032021-0221232322203222-2003010122320010-1132321012021310-2110021100223322-3313331011132132"></a>

## Next pages — sli_to_global_dr / 101322031102 / 4

- [ingress_egress_gw.global_network_list.global_network_connections.sli_to_global_dr.global_vn](data-sources--gcp_vpc_site--reference--group-002.md#canonical-1032131322013321-2313331131010322-1030123031020233-3112221000202120-3001213320012301-2331220002033000-0202200033111310-3021013121220331)
- [ingress_egress_gw.global_network_list.global_network_connections](data-sources--gcp_vpc_site--reference--group-002.md#canonical-1002112130221021-3330011001330322-1320022131222221-3003201102130100-3020213123210002-0101032321103312-2112023202211223-1310132021211212)
- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-1132222133321101-1232102200201102-2111112102132122-0010023221133123-2333213300031302-1001011022311003-3121103220313010-2103321200303202)

<a id="canonical-1032131322013321-2313331131010322-1030123031020233-3112221000202120-3001213320012301-2331220002033000-0202200033111310-3021013121220331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0131233302133212-2010121320300302-3122003222322233-0330313103023031-0311301333132212-0002013232012330-2302003012212001-2010103123302303"></a>

## ingress_egress_gw.global_network_list.global_network_connections.sli_to_global_dr.global_vn — global_vn / 021332333003 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-1132222133321101-1232102200201102-2111112102132122-0010023221133123-2333213300031302-1001011022311003-3121103220313010-2103321200303202)
- [Property reference](data-sources--gcp_vpc_site--reference--group-001.md#canonical-0023231003100101-1010210121112300-0303103132302303-0331000323201313-3003311210331022-3000312312212303-0030202320102022-1232013113103201)
- [ingress_egress_gw](data-sources--gcp_vpc_site--reference--group-001.md#canonical-1123313330020103-2220202203231111-2113330331022333-0133213021020321-0303223331023010-1320133103131013-2322232313123200-2330220010011031)
- [ingress_egress_gw.global_network_list](data-sources--gcp_vpc_site--reference--group-002.md#canonical-1130312120203102-2122110323320300-1102233103300112-2022130100300120-2110022023011312-0123323310213232-0112231123131301-3101313020332221)
- [ingress_egress_gw.global_network_list.global_network_connections](data-sources--gcp_vpc_site--reference--group-002.md#canonical-1002112130221021-3330011001330322-1320022131222221-3003201102130100-3020213123210002-0101032321103312-2112023202211223-1310132021211212)
- [ingress_egress_gw.global_network_list.global_network_connections.sli_to_global_dr](data-sources--gcp_vpc_site--reference--group-002.md#canonical-3330032003012321-1120210313032212-2312133112233103-3323330212212123-0323133022212321-0132112322301201-0023213202312310-3222103003101221)
- ingress_egress_gw.global_network_list.global_network_connections.sli_to_global_dr.global_vn

<a id="canonical-3302103113202030-2033003301000202-3302211232123023-2320030321330333-0311010211322322-0031111132311322-2213330100310210-2000022022212111"></a>

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

<a id="canonical-1000230012233003-0100233123311300-2301021001123221-3120201110101233-3323013033012221-0202332333300133-1223022112012211-0321200222112300"></a>

## Direct properties — global_vn / 021332333003 / 3

<a id="canonical-0100302310013220-1011013130203202-3320300230310222-2002131121023330-2030100101131121-2310023233301112-3030000011200120-0100200112020333"></a>

<a id="canonical-3211021030322232-0123321102000013-1102020022031023-0302123210301000-1032301011002101-2001001222133000-2003010102010032-3123201001203212"></a>

## name property — global_vn / 021332333003 / 4

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

<a id="canonical-2202311201300030-0120212220201021-0313200330232003-1213332212013103-2130211303122322-1330212300101202-3223023132311010-0220020201233210"></a>

<a id="canonical-0002122210131303-0102321311120010-3011013213121331-1300312302333300-0213231031000012-2113001011132222-3133213031230133-1320021113123323"></a>

## namespace property — global_vn / 021332333003 / 5

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

<a id="canonical-0000030113321100-3210132101003301-2303131230303122-0203112133022032-3231111001313210-3122222132111000-3111001220122002-3110132201133031"></a>

<a id="canonical-3133130011210300-3011333333110103-1230102322223321-0002130220001211-3213303201200013-1101010101110000-3233022311212030-2213322021230200"></a>

## tenant property — global_vn / 021332333003 / 6

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

<a id="canonical-0213230021030113-1021020202001121-1211023030110011-1311030210301003-0320000030002211-3012100210221211-1011200200333030-3330332011122020"></a>

## Next pages — global_vn / 021332333003 / 7

- [ingress_egress_gw.global_network_list.global_network_connections.sli_to_global_dr](data-sources--gcp_vpc_site--reference--group-002.md#canonical-3330032003012321-1120210313032212-2312133112233103-3323330212212123-0323133022212321-0132112322301201-0023213202312310-3222103003101221)
- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-1132222133321101-1232102200201102-2111112102132122-0010023221133123-2333213300031302-1001011022311003-3121103220313010-2103321200303202)

<a id="canonical-3132320012033303-1003211031032332-1312303212112333-2023222030303131-3200013303311031-3300132001132032-0203020233003102-1001212030001002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3220310301322131-3113312321023030-0103213121310000-3310222131312333-1223321011020310-0111013121212013-2011210130330301-3131321212233123"></a>

## ingress_egress_gw.global_network_list.global_network_connections.slo_to_global_dr — slo_to_global_dr / 223203331323 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-1132222133321101-1232102200201102-2111112102132122-0010023221133123-2333213300031302-1001011022311003-3121103220313010-2103321200303202)
- [Property reference](data-sources--gcp_vpc_site--reference--group-001.md#canonical-0023231003100101-1010210121112300-0303103132302303-0331000323201313-3003311210331022-3000312312212303-0030202320102022-1232013113103201)
- [ingress_egress_gw](data-sources--gcp_vpc_site--reference--group-001.md#canonical-1123313330020103-2220202203231111-2113330331022333-0133213021020321-0303223331023010-1320133103131013-2322232313123200-2330220010011031)
- [ingress_egress_gw.global_network_list](data-sources--gcp_vpc_site--reference--group-002.md#canonical-1130312120203102-2122110323320300-1102233103300112-2022130100300120-2110022023011312-0123323310213232-0112231123131301-3101313020332221)
- [ingress_egress_gw.global_network_list.global_network_connections](data-sources--gcp_vpc_site--reference--group-002.md#canonical-1002112130221021-3330011001330322-1320022131222221-3003201102130100-3020213123210002-0101032321103312-2112023202211223-1310132021211212)
- ingress_egress_gw.global_network_list.global_network_connections.slo_to_global_dr

<a id="canonical-1321300321231221-3201210130020211-2130200002123303-1312123311123302-2001130031121322-0013212030331233-3311103331222112-0032231310100123"></a>

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

<a id="canonical-0320022333133220-3130303002012112-2033110120112203-1021003131121112-3233230130010010-1232102301002321-0021203301020212-1303303323212201"></a>

## Direct properties — slo_to_global_dr / 223203331323 / 3

- [global_vn](data-sources--gcp_vpc_site--reference--group-002.md#canonical-3001221001312312-3021221302231213-3323302010121311-2302201010331311-2123033033010102-0031202300011111-2300213330011122-3230203110220313): complete subsection reference.

<a id="canonical-2023132200020130-2002020302031120-1223101133013330-2320101233121312-0332000211101230-2301313023221121-2303210110232223-3123320130330302"></a>

## Next pages — slo_to_global_dr / 223203331323 / 4

- [ingress_egress_gw.global_network_list.global_network_connections.slo_to_global_dr.global_vn](data-sources--gcp_vpc_site--reference--group-002.md#canonical-3001221001312312-3021221302231213-3323302010121311-2302201010331311-2123033033010102-0031202300011111-2300213330011122-3230203110220313)
- [ingress_egress_gw.global_network_list.global_network_connections](data-sources--gcp_vpc_site--reference--group-002.md#canonical-1002112130221021-3330011001330322-1320022131222221-3003201102130100-3020213123210002-0101032321103312-2112023202211223-1310132021211212)
- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-1132222133321101-1232102200201102-2111112102132122-0010023221133123-2333213300031302-1001011022311003-3121103220313010-2103321200303202)

<a id="canonical-3001221001312312-3021221302231213-3323302010121311-2302201010331311-2123033033010102-0031202300011111-2300213330011122-3230203110220313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0323230330310223-1300002203100032-0230202231223102-1101323210012210-3000221021132222-0312321311103013-1202211022033223-0110211231300001"></a>

## ingress_egress_gw.global_network_list.global_network_connections.slo_to_global_dr.global_vn — global_vn / 032101123301 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-1132222133321101-1232102200201102-2111112102132122-0010023221133123-2333213300031302-1001011022311003-3121103220313010-2103321200303202)
- [Property reference](data-sources--gcp_vpc_site--reference--group-001.md#canonical-0023231003100101-1010210121112300-0303103132302303-0331000323201313-3003311210331022-3000312312212303-0030202320102022-1232013113103201)
- [ingress_egress_gw](data-sources--gcp_vpc_site--reference--group-001.md#canonical-1123313330020103-2220202203231111-2113330331022333-0133213021020321-0303223331023010-1320133103131013-2322232313123200-2330220010011031)
- [ingress_egress_gw.global_network_list](data-sources--gcp_vpc_site--reference--group-002.md#canonical-1130312120203102-2122110323320300-1102233103300112-2022130100300120-2110022023011312-0123323310213232-0112231123131301-3101313020332221)
- [ingress_egress_gw.global_network_list.global_network_connections](data-sources--gcp_vpc_site--reference--group-002.md#canonical-1002112130221021-3330011001330322-1320022131222221-3003201102130100-3020213123210002-0101032321103312-2112023202211223-1310132021211212)
- [ingress_egress_gw.global_network_list.global_network_connections.slo_to_global_dr](data-sources--gcp_vpc_site--reference--group-002.md#canonical-3132320012033303-1003211031032332-1312303212112333-2023222030303131-3200013303311031-3300132001132032-0203020233003102-1001212030001002)
- ingress_egress_gw.global_network_list.global_network_connections.slo_to_global_dr.global_vn

<a id="canonical-1313023012101113-0011103032321200-3223022222102332-0001003100122123-1033012331130232-0120332221331123-0133323031003123-0102013320221311"></a>

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

<a id="canonical-0110320120223030-0202123033232301-3112330302302202-2331000023113231-3323013221100132-2013301101113301-3231013200212110-1321221210031233"></a>

## Direct properties — global_vn / 032101123301 / 3

<a id="canonical-2132321002220031-0112023011110111-0133012313200333-0011003222000310-3023100222012330-3001301332023020-1013132202102121-1102021011101330"></a>

<a id="canonical-1120100100301323-1203112111331132-1122013131332313-2320000212132003-0031220320222203-3301020300011003-3233023320213221-3033222300023110"></a>

## name property — global_vn / 032101123301 / 4

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

<a id="canonical-3313230330203201-1013223120012102-2002102110010300-3101233123102331-1200032202022321-1230311203321013-0221223110301301-1210022020322110"></a>

<a id="canonical-3102030222132101-2212100301013331-2112122122213302-2020212330300123-0123100021202301-3322002021131100-1213301011303111-2003013210202311"></a>

## namespace property — global_vn / 032101123301 / 5

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

<a id="canonical-1133033110203021-0220123033000002-1220120222001123-3031100011302302-2302231111203112-1300121120110223-3330230003100110-2321220130321222"></a>

<a id="canonical-3221332331332231-1131300003031010-0033031121003331-2033231120011300-2320332223111131-2000003210222200-3303310003201132-0003233011000212"></a>

## tenant property — global_vn / 032101123301 / 6

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

<a id="canonical-0302213200020202-1223312023101211-2100233133330023-3110301332132112-0220002222300223-3122302111333323-2212030100301100-0310121301231200"></a>

## Next pages — global_vn / 032101123301 / 7

- [ingress_egress_gw.global_network_list.global_network_connections.slo_to_global_dr](data-sources--gcp_vpc_site--reference--group-002.md#canonical-3132320012033303-1003211031032332-1312303212112333-2023222030303131-3200013303311031-3300132001132032-0203020233003102-1001212030001002)
- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-1132222133321101-1232102200201102-2111112102132122-0010023221133123-2333213300031302-1001011022311003-3121103220313010-2103321200303202)

<a id="canonical-1103010332011311-1122120123211002-0313120122102101-1212112111232323-2002213301120131-2031023110130320-2332020000113011-2101203131100023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1002303000220303-0333113002321221-1330312001031121-2323300110201101-1312123321122303-2122221123202311-2012101200130231-2223223033203203"></a>

## ingress_egress_gw.inside_network — inside_network / 120130202011 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-1132222133321101-1232102200201102-2111112102132122-0010023221133123-2333213300031302-1001011022311003-3121103220313010-2103321200303202)
- [Property reference](data-sources--gcp_vpc_site--reference--group-001.md#canonical-0023231003100101-1010210121112300-0303103132302303-0331000323201313-3003311210331022-3000312312212303-0030202320102022-1232013113103201)
- [ingress_egress_gw](data-sources--gcp_vpc_site--reference--group-001.md#canonical-1123313330020103-2220202203231111-2113330331022333-0133213021020321-0303223331023010-1320133103131013-2322232313123200-2330220010011031)
- ingress_egress_gw.inside_network

<a id="canonical-0100211132231310-3310032001321123-1013102103221111-3221220201002330-3031200123011313-3222232120302230-1303003031000232-0221020313323031"></a>

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

<a id="canonical-1022221130221110-1313103313010212-2120121230321130-2320232100300122-2010131010100003-3312133113112333-0011233030300131-3232011313333132"></a>

## Direct properties — inside_network / 120130202011 / 3

- [existing_network](data-sources--gcp_vpc_site--reference--group-002.md#canonical-2003222222331121-3313330221302020-3032020103012030-3203112331110311-3123322101101011-1121113020211032-2120320222230321-2030303330201120): complete subsection reference.

- [new_network](data-sources--gcp_vpc_site--reference--group-002.md#canonical-2003122002300310-1301323321302301-2302130322121032-2112303301233132-0232021023131312-1121330231202132-0201030202001322-0012233333012003): complete subsection reference.

- [new_network_autogenerate](data-sources--gcp_vpc_site--reference--group-002.md#canonical-1210302030101032-0210033310223013-3011033212323100-2220302001201210-3121232032302320-2333103232023202-2302310001131010-1210331322320210): complete subsection reference.

<a id="canonical-1003111132300003-1300022021033003-3020121321110023-3132103220301011-0222022112310222-2212221133022211-2032113123201312-3132322202021011"></a>

## Next pages — inside_network / 120130202011 / 4

- [ingress_egress_gw.inside_network.existing_network](data-sources--gcp_vpc_site--reference--group-002.md#canonical-2003222222331121-3313330221302020-3032020103012030-3203112331110311-3123322101101011-1121113020211032-2120320222230321-2030303330201120)
- [ingress_egress_gw.inside_network.new_network](data-sources--gcp_vpc_site--reference--group-002.md#canonical-2003122002300310-1301323321302301-2302130322121032-2112303301233132-0232021023131312-1121330231202132-0201030202001322-0012233333012003)
- [ingress_egress_gw.inside_network.new_network_autogenerate](data-sources--gcp_vpc_site--reference--group-002.md#canonical-1210302030101032-0210033310223013-3011033212323100-2220302001201210-3121232032302320-2333103232023202-2302310001131010-1210331322320210)
- [ingress_egress_gw](data-sources--gcp_vpc_site--reference--group-001.md#canonical-1123313330020103-2220202203231111-2113330331022333-0133213021020321-0303223331023010-1320133103131013-2322232313123200-2330220010011031)
- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-1132222133321101-1232102200201102-2111112102132122-0010023221133123-2333213300031302-1001011022311003-3121103220313010-2103321200303202)

<a id="canonical-2003222222331121-3313330221302020-3032020103012030-3203112331110311-3123322101101011-1121113020211032-2120320222230321-2030303330201120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0232020132331302-0112012013122312-3103123331310010-1031113100100032-2203211220122232-3020201331320303-1111102331213210-0321133003010212"></a>

## ingress_egress_gw.inside_network.existing_network — existing_network / 103322023112 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-1132222133321101-1232102200201102-2111112102132122-0010023221133123-2333213300031302-1001011022311003-3121103220313010-2103321200303202)
- [Property reference](data-sources--gcp_vpc_site--reference--group-001.md#canonical-0023231003100101-1010210121112300-0303103132302303-0331000323201313-3003311210331022-3000312312212303-0030202320102022-1232013113103201)
- [ingress_egress_gw](data-sources--gcp_vpc_site--reference--group-001.md#canonical-1123313330020103-2220202203231111-2113330331022333-0133213021020321-0303223331023010-1320133103131013-2322232313123200-2330220010011031)
- [ingress_egress_gw.inside_network](data-sources--gcp_vpc_site--reference--group-002.md#canonical-1103010332011311-1122120123211002-0313120122102101-1212112111232323-2002213301120131-2031023110130320-2332020000113011-2101203131100023)
- ingress_egress_gw.inside_network.existing_network

<a id="canonical-1220121103122102-3103121301131010-3231023011230101-3303131303211232-2113213300302312-1030100112201232-2333212110212110-3130210223210203"></a>

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

<a id="canonical-3013030300011123-1213330203010211-1133333013303320-1001311000332002-0323102203303323-0202210330322300-0032032003231123-1121032001310320"></a>

## Direct properties — existing_network / 103322023112 / 3

<a id="canonical-1100203330003302-1203213110222102-1103000211330132-3032021132033230-1312002212323231-2313222101023303-2121102233223221-0331230030213130"></a>

<a id="canonical-3222311122011002-0222121320033331-1332033223123101-0130133133302130-2302103220331123-3020010323022300-0033330011003322-3113122321301312"></a>

## name property — existing_network / 103322023112 / 4

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

<a id="canonical-3321212311012112-0213231201223032-1311321103332311-3210202002302302-0012021120313322-3000332331020132-3222121323200211-1111002311200301"></a>

## Next pages — existing_network / 103322023112 / 5

- [ingress_egress_gw.inside_network](data-sources--gcp_vpc_site--reference--group-002.md#canonical-1103010332011311-1122120123211002-0313120122102101-1212112111232323-2002213301120131-2031023110130320-2332020000113011-2101203131100023)
- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-1132222133321101-1232102200201102-2111112102132122-0010023221133123-2333213300031302-1001011022311003-3121103220313010-2103321200303202)

<a id="canonical-2003122002300310-1301323321302301-2302130322121032-2112303301233132-0232021023131312-1121330231202132-0201030202001322-0012233333012003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2323002223333001-2111001112100223-0321332203132211-1203212231103103-2021101211031122-2222121332133323-0123031000122312-2213031320212030"></a>

## ingress_egress_gw.inside_network.new_network — new_network / 300102332003 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-1132222133321101-1232102200201102-2111112102132122-0010023221133123-2333213300031302-1001011022311003-3121103220313010-2103321200303202)
- [Property reference](data-sources--gcp_vpc_site--reference--group-001.md#canonical-0023231003100101-1010210121112300-0303103132302303-0331000323201313-3003311210331022-3000312312212303-0030202320102022-1232013113103201)
- [ingress_egress_gw](data-sources--gcp_vpc_site--reference--group-001.md#canonical-1123313330020103-2220202203231111-2113330331022333-0133213021020321-0303223331023010-1320133103131013-2322232313123200-2330220010011031)
- [ingress_egress_gw.inside_network](data-sources--gcp_vpc_site--reference--group-002.md#canonical-1103010332011311-1122120123211002-0313120122102101-1212112111232323-2002213301120131-2031023110130320-2332020000113011-2101203131100023)
- ingress_egress_gw.inside_network.new_network

<a id="canonical-2210300111302201-0201113302221132-0132000002112331-1311313110111102-2323132213031100-3310231002121232-1033031203310321-2022102203121210"></a>

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

<a id="canonical-0003030313032030-1133020310213201-0010223332330000-0122323320201101-2212111130231201-2033332312022233-3020233220321321-1202120112131203"></a>

## Direct properties — new_network / 300102332003 / 3

<a id="canonical-1022002303201230-0102322100031231-3011323110211212-2302133030123203-3312013333033332-3122213220313230-2101332321100020-1030222202102202"></a>

<a id="canonical-3033202311110120-1201233010330112-3203100211123112-1211200021111302-2212223330303133-1032111111222331-0020302230100003-1110020012021330"></a>

## name property — new_network / 300102332003 / 4

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

<a id="canonical-0322231313333301-0001313123023100-2202201333100020-2103100210201133-3202323321310110-2110310103113310-1021311230221333-3100333220033110"></a>

## Next pages — new_network / 300102332003 / 5

- [ingress_egress_gw.inside_network](data-sources--gcp_vpc_site--reference--group-002.md#canonical-1103010332011311-1122120123211002-0313120122102101-1212112111232323-2002213301120131-2031023110130320-2332020000113011-2101203131100023)
- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-1132222133321101-1232102200201102-2111112102132122-0010023221133123-2333213300031302-1001011022311003-3121103220313010-2103321200303202)

<a id="canonical-1210302030101032-0210033310223013-3011033212323100-2220302001201210-3121232032302320-2333103232023202-2302310001131010-1210331322320210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1233010332121220-1310130323202131-0330201210023001-2110310221203031-3020222112323231-1321321130310023-2033031200312021-0322100331030120"></a>

## ingress_egress_gw.inside_network.new_network_autogenerate — new_network_autogenerate / 022102323121 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-1132222133321101-1232102200201102-2111112102132122-0010023221133123-2333213300031302-1001011022311003-3121103220313010-2103321200303202)
- [Property reference](data-sources--gcp_vpc_site--reference--group-001.md#canonical-0023231003100101-1010210121112300-0303103132302303-0331000323201313-3003311210331022-3000312312212303-0030202320102022-1232013113103201)
- [ingress_egress_gw](data-sources--gcp_vpc_site--reference--group-001.md#canonical-1123313330020103-2220202203231111-2113330331022333-0133213021020321-0303223331023010-1320133103131013-2322232313123200-2330220010011031)
- [ingress_egress_gw.inside_network](data-sources--gcp_vpc_site--reference--group-002.md#canonical-1103010332011311-1122120123211002-0313120122102101-1212112111232323-2002213301120131-2031023110130320-2332020000113011-2101203131100023)
- ingress_egress_gw.inside_network.new_network_autogenerate

<a id="canonical-2221101310120210-3103010101232031-1303203121200312-1131112222033310-0201330122222120-3120230200032331-3133002122132231-3000331213330220"></a>

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

<a id="canonical-1232103302022133-3033101301010120-2232211231002120-2112221321020020-2213010012123331-0310230212223112-0113023122303221-1311021230023001"></a>

## Direct properties — new_network_autogenerate / 022102323121 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3303033101030110-1330110202103122-2203302221320202-0212313033220210-0102211323232301-0112013020202031-3122203332300123-3012123203031212"></a>

## Next pages — new_network_autogenerate / 022102323121 / 4

- [ingress_egress_gw.inside_network](data-sources--gcp_vpc_site--reference--group-002.md#canonical-1103010332011311-1122120123211002-0313120122102101-1212112111232323-2002213301120131-2031023110130320-2332020000113011-2101203131100023)
- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-1132222133321101-1232102200201102-2111112102132122-0010023221133123-2333213300031302-1001011022311003-3121103220313010-2103321200303202)

<a id="canonical-3200313310131232-0210110202101312-2021312210102120-1112102012331303-3023321133320033-3321213133200210-3120212332313312-1233223233330302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2010231300221121-2303102231213023-1232013020003220-2301102110322010-1023112313002132-2131030300201300-3212231002303112-2102003022313301"></a>

## ingress_egress_gw.inside_static_routes — inside_static_routes / 021120103020 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-1132222133321101-1232102200201102-2111112102132122-0010023221133123-2333213300031302-1001011022311003-3121103220313010-2103321200303202)
- [Property reference](data-sources--gcp_vpc_site--reference--group-001.md#canonical-0023231003100101-1010210121112300-0303103132302303-0331000323201313-3003311210331022-3000312312212303-0030202320102022-1232013113103201)
- [ingress_egress_gw](data-sources--gcp_vpc_site--reference--group-001.md#canonical-1123313330020103-2220202203231111-2113330331022333-0133213021020321-0303223331023010-1320133103131013-2322232313123200-2330220010011031)
- ingress_egress_gw.inside_static_routes

<a id="canonical-0330302001121010-0000023223330023-1311322132211002-3110330223133003-3210302123201210-0131230101121200-1033000331322131-2033330033032113"></a>

Type: `"single"`. Computed.

Configuration parameter for inside static routes.

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

<a id="canonical-3032113012201221-2133013132100232-1302133232003131-3311322232221030-1331303002221002-2311120203303213-3100220230032010-0132212010011130"></a>

## Direct properties — inside_static_routes / 021120103020 / 3

- [static_route_list](data-sources--gcp_vpc_site--reference--group-002.md#canonical-0022032111233021-3032231100003203-2232232213033220-2112100203123200-2101002223201133-1132032110122011-3133130103121231-0111311331112021): complete subsection reference.

<a id="canonical-0031132133223021-0330132021121201-3013301101321101-3133032321330010-2122033301030300-1033213031213123-1230301032123112-0101211032010213"></a>

## Next pages — inside_static_routes / 021120103020 / 4

- [ingress_egress_gw.inside_static_routes.static_route_list](data-sources--gcp_vpc_site--reference--group-002.md#canonical-0022032111233021-3032231100003203-2232232213033220-2112100203123200-2101002223201133-1132032110122011-3133130103121231-0111311331112021)
- [ingress_egress_gw](data-sources--gcp_vpc_site--reference--group-001.md#canonical-1123313330020103-2220202203231111-2113330331022333-0133213021020321-0303223331023010-1320133103131013-2322232313123200-2330220010011031)
- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-1132222133321101-1232102200201102-2111112102132122-0010023221133123-2333213300031302-1001011022311003-3121103220313010-2103321200303202)

<a id="canonical-0022032111233021-3032231100003203-2232232213033220-2112100203123200-2101002223201133-1132032110122011-3133130103121231-0111311331112021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2210320102230101-1101010232222011-3001120100211210-3301330133100233-2332201311033303-1020202311031230-3231113002112331-0300003321021313"></a>

## ingress_egress_gw.inside_static_routes.static_route_list — static_route_list / 202023101322 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-1132222133321101-1232102200201102-2111112102132122-0010023221133123-2333213300031302-1001011022311003-3121103220313010-2103321200303202)
- [Property reference](data-sources--gcp_vpc_site--reference--group-001.md#canonical-0023231003100101-1010210121112300-0303103132302303-0331000323201313-3003311210331022-3000312312212303-0030202320102022-1232013113103201)
- [ingress_egress_gw](data-sources--gcp_vpc_site--reference--group-001.md#canonical-1123313330020103-2220202203231111-2113330331022333-0133213021020321-0303223331023010-1320133103131013-2322232313123200-2330220010011031)
- [ingress_egress_gw.inside_static_routes](data-sources--gcp_vpc_site--reference--group-002.md#canonical-3200313310131232-0210110202101312-2021312210102120-1112102012331303-3023321133320033-3321213133200210-3120212332313312-1233223233330302)
- ingress_egress_gw.inside_static_routes.static_route_list

<a id="canonical-1133110212121311-2212112003103102-0000312111112121-0211310132101302-3203232220132101-3030203111200122-0011110120121120-2121210301211210"></a>

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-0110223210302230-3302323303110232-3313200110001332-3211201003123322-0312102222200030-1213111021100313-0121330221032001-3313000331133103"></a>

## Direct properties — static_route_list / 202023101322 / 3

- [custom_static_route](data-sources--gcp_vpc_site--reference--group-002.md#canonical-0202301330013012-0202311012013203-0330311200232332-1121021133231003-1023203132102220-0022302011110020-0313200112032320-1320010200100322): complete subsection reference.

<a id="canonical-1001130210203312-0101101212231011-0311111010013220-2220101302203232-2233102110230220-2232021133203332-2333231221001310-2033221111022211"></a>

<a id="canonical-0003220203220132-0320303003011001-0001011213321300-3101022022111032-1102201023223111-2101131203210103-3333213121302313-0223020101110031"></a>

## simple_static_route property — static_route_list / 202023101322 / 4

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
    "ves.io.schema.rules.string.ipv4_prefix": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv4_prefix": "true"
  }
}
```

<a id="canonical-0021122021030300-2102120310122202-0003323333031000-3322111013110001-2132233111002002-2022031110133011-3132313301103333-0001312032123130"></a>

## Next pages — static_route_list / 202023101322 / 5

- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route](data-sources--gcp_vpc_site--reference--group-002.md#canonical-0202301330013012-0202311012013203-0330311200232332-1121021133231003-1023203132102220-0022302011110020-0313200112032320-1320010200100322)
- [ingress_egress_gw.inside_static_routes](data-sources--gcp_vpc_site--reference--group-002.md#canonical-3200313310131232-0210110202101312-2021312210102120-1112102012331303-3023321133320033-3321213133200210-3120212332313312-1233223233330302)
- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-1132222133321101-1232102200201102-2111112102132122-0010023221133123-2333213300031302-1001011022311003-3121103220313010-2103321200303202)

<a id="canonical-0202301330013012-0202311012013203-0330311200232332-1121021133231003-1023203132102220-0022302011110020-0313200112032320-1320010200100322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1130031210111000-2130332221202310-3323301022120332-0130121230331313-2201310032010312-2213311211120003-0302121322231130-3031111001221232"></a>

## ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route — custom_static_route / 232223031321 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-1132222133321101-1232102200201102-2111112102132122-0010023221133123-2333213300031302-1001011022311003-3121103220313010-2103321200303202)
- [Property reference](data-sources--gcp_vpc_site--reference--group-001.md#canonical-0023231003100101-1010210121112300-0303103132302303-0331000323201313-3003311210331022-3000312312212303-0030202320102022-1232013113103201)
- [ingress_egress_gw](data-sources--gcp_vpc_site--reference--group-001.md#canonical-1123313330020103-2220202203231111-2113330331022333-0133213021020321-0303223331023010-1320133103131013-2322232313123200-2330220010011031)
- [ingress_egress_gw.inside_static_routes](data-sources--gcp_vpc_site--reference--group-002.md#canonical-3200313310131232-0210110202101312-2021312210102120-1112102012331303-3023321133320033-3321213133200210-3120212332313312-1233223233330302)
- [ingress_egress_gw.inside_static_routes.static_route_list](data-sources--gcp_vpc_site--reference--group-002.md#canonical-0022032111233021-3032231100003203-2232232213033220-2112100203123200-2101002223201133-1132032110122011-3133130103121231-0111311331112021)
- ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route

<a id="canonical-0011003110330210-0021131012120112-1221102110302032-3333000133113130-3023033303230023-2033312202122002-2221023221112123-1122130211012331"></a>

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

<a id="canonical-0012211310123203-0101312200311131-3310231123201133-0120300320321323-3333122100033111-3302301213330330-2331010131320201-0303103130132212"></a>

## Direct properties — custom_static_route / 232223031321 / 3

<a id="canonical-1121211012131102-3032012301112330-0202321233002230-2232303200023021-2311020212100033-1111210032023000-3302131211312032-2211310023300300"></a>

<a id="canonical-0000110112302000-2112001113300121-1223130000310230-1202013213203233-2132122121002023-3000132120103201-3110223030133202-0131212212311123"></a>

## attrs property — custom_static_route / 232223031321 / 4

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
    "ves.io.schema.rules.repeated.max_items": "4"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "4"
  }
}
```

- [labels](data-sources--gcp_vpc_site--reference--group-002.md#canonical-0331233223322020-2313012112130111-3302210023301120-1303320113313301-1130200122102323-0331313320233231-0132003303203330-2123121031100312): complete subsection reference.

- [nexthop](data-sources--gcp_vpc_site--reference--group-002.md#canonical-3321032023000021-0000021032023103-2003122132323032-1121100213300200-2112303323013100-2020032023013330-2311103200013321-3231023002332203): complete subsection reference.

- [subnets](data-sources--gcp_vpc_site--reference--group-002.md#canonical-2122221323130301-1313122121231331-0300323002000300-1301230032030102-2023320221320212-3221013211122031-2220012333303310-2011001021111030): complete subsection reference.

<a id="canonical-0233132200203323-3323030232332123-0133313010011212-0200031321302313-1213301230033120-2230223110012303-1331233000110100-1202130032331222"></a>

## Next pages — custom_static_route / 232223031321 / 5

- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.labels](data-sources--gcp_vpc_site--reference--group-002.md#canonical-0331233223322020-2313012112130111-3302210023301120-1303320113313301-1130200122102323-0331313320233231-0132003303203330-2123121031100312)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop](data-sources--gcp_vpc_site--reference--group-002.md#canonical-3321032023000021-0000021032023103-2003122132323032-1121100213300200-2112303323013100-2020032023013330-2311103200013321-3231023002332203)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets](data-sources--gcp_vpc_site--reference--group-002.md#canonical-2122221323130301-1313122121231331-0300323002000300-1301230032030102-2023320221320212-3221013211122031-2220012333303310-2011001021111030)
- [ingress_egress_gw.inside_static_routes.static_route_list](data-sources--gcp_vpc_site--reference--group-002.md#canonical-0022032111233021-3032231100003203-2232232213033220-2112100203123200-2101002223201133-1132032110122011-3133130103121231-0111311331112021)
- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-1132222133321101-1232102200201102-2111112102132122-0010023221133123-2333213300031302-1001011022311003-3121103220313010-2103321200303202)

<a id="canonical-0331233223322020-2313012112130111-3302210023301120-1303320113313301-1130200122102323-0331313320233231-0132003303203330-2123121031100312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3133130312131003-3303320031030220-0221031032302333-3031100210211103-0002013102103130-1210022020000312-1223010220000332-2333011012221030"></a>

## ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.labels — labels / 013000022000 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-1132222133321101-1232102200201102-2111112102132122-0010023221133123-2333213300031302-1001011022311003-3121103220313010-2103321200303202)
- [Property reference](data-sources--gcp_vpc_site--reference--group-001.md#canonical-0023231003100101-1010210121112300-0303103132302303-0331000323201313-3003311210331022-3000312312212303-0030202320102022-1232013113103201)
- [ingress_egress_gw](data-sources--gcp_vpc_site--reference--group-001.md#canonical-1123313330020103-2220202203231111-2113330331022333-0133213021020321-0303223331023010-1320133103131013-2322232313123200-2330220010011031)
- [ingress_egress_gw.inside_static_routes](data-sources--gcp_vpc_site--reference--group-002.md#canonical-3200313310131232-0210110202101312-2021312210102120-1112102012331303-3023321133320033-3321213133200210-3120212332313312-1233223233330302)
- [ingress_egress_gw.inside_static_routes.static_route_list](data-sources--gcp_vpc_site--reference--group-002.md#canonical-0022032111233021-3032231100003203-2232232213033220-2112100203123200-2101002223201133-1132032110122011-3133130103121231-0111311331112021)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route](data-sources--gcp_vpc_site--reference--group-002.md#canonical-0202301330013012-0202311012013203-0330311200232332-1121021133231003-1023203132102220-0022302011110020-0313200112032320-1320010200100322)
- ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.labels

<a id="canonical-0111323000131300-1132031131213102-2303330201102110-1002310223103331-2120112000011031-0023031132220102-0130132120323321-2201333003010203"></a>

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

<a id="canonical-3020321113220223-1210332030301203-3230231021130303-1232023102330033-2330033212332033-2000102231312033-0033212000102102-0312013121031301"></a>

## Direct properties — labels / 013000022000 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0101320223201112-0030111133202202-3132333122123213-2233311220332112-0122230113230021-2123302003112231-0302231210233110-0222233320012321"></a>

## Next pages — labels / 013000022000 / 4

- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route](data-sources--gcp_vpc_site--reference--group-002.md#canonical-0202301330013012-0202311012013203-0330311200232332-1121021133231003-1023203132102220-0022302011110020-0313200112032320-1320010200100322)
- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-1132222133321101-1232102200201102-2111112102132122-0010023221133123-2333213300031302-1001011022311003-3121103220313010-2103321200303202)

<a id="canonical-3321032023000021-0000021032023103-2003122132323032-1121100213300200-2112303323013100-2020032023013330-2311103200013321-3231023002332203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1121120030120221-0313120000121211-1012111013023010-3231001131100323-2303203202032110-2231033100221012-0001012012103301-2010020132031213"></a>

## ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop — nexthop / 232003123010 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-1132222133321101-1232102200201102-2111112102132122-0010023221133123-2333213300031302-1001011022311003-3121103220313010-2103321200303202)
- [Property reference](data-sources--gcp_vpc_site--reference--group-001.md#canonical-0023231003100101-1010210121112300-0303103132302303-0331000323201313-3003311210331022-3000312312212303-0030202320102022-1232013113103201)
- [ingress_egress_gw](data-sources--gcp_vpc_site--reference--group-001.md#canonical-1123313330020103-2220202203231111-2113330331022333-0133213021020321-0303223331023010-1320133103131013-2322232313123200-2330220010011031)
- [ingress_egress_gw.inside_static_routes](data-sources--gcp_vpc_site--reference--group-002.md#canonical-3200313310131232-0210110202101312-2021312210102120-1112102012331303-3023321133320033-3321213133200210-3120212332313312-1233223233330302)
- [ingress_egress_gw.inside_static_routes.static_route_list](data-sources--gcp_vpc_site--reference--group-002.md#canonical-0022032111233021-3032231100003203-2232232213033220-2112100203123200-2101002223201133-1132032110122011-3133130103121231-0111311331112021)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route](data-sources--gcp_vpc_site--reference--group-002.md#canonical-0202301330013012-0202311012013203-0330311200232332-1121021133231003-1023203132102220-0022302011110020-0313200112032320-1320010200100322)
- ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop

<a id="canonical-2210033313010300-2111131021203311-3133311323312203-0011230211132031-3120332213003022-0302122333003031-2231203323300121-0011003132200323"></a>

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

<a id="canonical-2001032130022023-1000210122310021-3120200130301203-0110323222321330-0020313110321223-3202302100211221-3213230033113301-1112232212020032"></a>

## Direct properties — nexthop / 232003123010 / 3

- [interface](data-sources--gcp_vpc_site--reference--group-002.md#canonical-2013111313022331-0103300023301031-1021131211320133-1102313211303030-3222013212121330-3020231230121123-1203021212021313-0023333110010303): complete subsection reference.

- [nexthop_address](data-sources--gcp_vpc_site--reference--group-002.md#canonical-1231320312112220-0303322321010321-1320233131012130-0021010330203222-3022101101211110-0231313332332320-1320333301122323-0221133132101110): complete subsection reference.

<a id="canonical-0333023003123210-2133331231010203-1200111222022022-0022020033312331-1103033232100102-1302230311030130-1110201013333013-3230312102233111"></a>

<a id="canonical-1010213123103101-0001003330332320-2101112332211133-3120301110312221-2101133111103113-1102301301212220-0333023012313211-0223303210010020"></a>

## type property — nexthop / 232003123010 / 4

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

<a id="canonical-2212323232223320-2231300300133301-3103332122020011-0330122301322031-0122232233013203-2210313020221031-1310232123022013-2321300111321232"></a>

## Next pages — nexthop / 232003123010 / 5

- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.interface](data-sources--gcp_vpc_site--reference--group-002.md#canonical-2013111313022331-0103300023301031-1021131211320133-1102313211303030-3222013212121330-3020231230121123-1203021212021313-0023333110010303)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](data-sources--gcp_vpc_site--reference--group-002.md#canonical-1231320312112220-0303322321010321-1320233131012130-0021010330203222-3022101101211110-0231313332332320-1320333301122323-0221133132101110)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route](data-sources--gcp_vpc_site--reference--group-002.md#canonical-0202301330013012-0202311012013203-0330311200232332-1121021133231003-1023203132102220-0022302011110020-0313200112032320-1320010200100322)
- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-1132222133321101-1232102200201102-2111112102132122-0010023221133123-2333213300031302-1001011022311003-3121103220313010-2103321200303202)

<a id="canonical-2013111313022331-0103300023301031-1021131211320133-1102313211303030-3222013212121330-3020231230121123-1203021212021313-0023333110010303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0133011210021032-0313210330101312-0223001210103200-0122100203220330-1201330311102120-0130100222131310-3312302333201311-2031221302130103"></a>

## ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.interface — interface / 232032300130 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-1132222133321101-1232102200201102-2111112102132122-0010023221133123-2333213300031302-1001011022311003-3121103220313010-2103321200303202)
- [Property reference](data-sources--gcp_vpc_site--reference--group-001.md#canonical-0023231003100101-1010210121112300-0303103132302303-0331000323201313-3003311210331022-3000312312212303-0030202320102022-1232013113103201)
- [ingress_egress_gw](data-sources--gcp_vpc_site--reference--group-001.md#canonical-1123313330020103-2220202203231111-2113330331022333-0133213021020321-0303223331023010-1320133103131013-2322232313123200-2330220010011031)
- [ingress_egress_gw.inside_static_routes](data-sources--gcp_vpc_site--reference--group-002.md#canonical-3200313310131232-0210110202101312-2021312210102120-1112102012331303-3023321133320033-3321213133200210-3120212332313312-1233223233330302)
- [ingress_egress_gw.inside_static_routes.static_route_list](data-sources--gcp_vpc_site--reference--group-002.md#canonical-0022032111233021-3032231100003203-2232232213033220-2112100203123200-2101002223201133-1132032110122011-3133130103121231-0111311331112021)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route](data-sources--gcp_vpc_site--reference--group-002.md#canonical-0202301330013012-0202311012013203-0330311200232332-1121021133231003-1023203132102220-0022302011110020-0313200112032320-1320010200100322)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop](data-sources--gcp_vpc_site--reference--group-002.md#canonical-3321032023000021-0000021032023103-2003122132323032-1121100213300200-2112303323013100-2020032023013330-2311103200013321-3231023002332203)
- ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.interface

<a id="canonical-2101002203120322-0110323120300202-0103013222203002-0131301322222303-1232113130311123-2232003202021321-3020223133133311-2100231113321020"></a>

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
    "ves.io.schema.rules.repeated.max_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1"
  }
}
```

<a id="canonical-2202020213302111-0103202132323020-2131032130222313-0030030312320210-1312230220321331-1213302103030123-1332211122301030-3013223222022303"></a>

## Direct properties — interface / 232032300130 / 3

<a id="canonical-1023313320222021-3103031321023330-0111230331223132-3132300030010002-3001100023130033-1321000321011230-3321023302300132-3120223003213100"></a>

<a id="canonical-3201132222123220-1312311123022220-0211033302033320-3333232220211330-0303310122300322-0021201131001221-1031030223100202-0230102023030320"></a>

## kind property — interface / 232032300130 / 4

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

<a id="canonical-2020122031212133-1021010302012101-2210121120011302-2230001313031030-2320322121010233-1213310221220030-1321103300313111-0120130113311110"></a>

<a id="canonical-2130123020210121-2300330100210110-2111233233222002-0101301013300113-0012233131322100-0330300332021310-3211320202120201-1212221023322002"></a>

## name property — interface / 232032300130 / 5

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

<a id="canonical-0131222332131201-1003012220322001-3020110000022001-3233332212123233-1101133130011013-0133233330213102-1003121230231230-1020132023331031"></a>

<a id="canonical-2301303213110312-0210323012222120-1202223031330010-3310003013231020-3121132211110122-1303120332322002-2002323000031020-0321223210201333"></a>

## namespace property — interface / 232032300130 / 6

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
  }
}
```

<a id="canonical-1132321012203333-0323321301313012-0020012301112222-0131123102322323-0012312322221230-1230000311233012-1021103030020131-3033311003132311"></a>

<a id="canonical-0231010003001201-3033202012013110-1203203312122230-2303333012001013-2001231311133201-2212200100032323-3123332130130100-3001303210321101"></a>

## tenant property — interface / 232032300130 / 7

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

<a id="canonical-1110121210030010-3301012032213333-2323323003211010-3222102012120100-1120111203113230-1131312110032131-1131022322120031-0312312101330220"></a>

<a id="canonical-3332212112302322-2230310130203323-1112123220033300-0333020212002022-2130303311120122-0300301011010222-3032023302103111-3231200132202023"></a>

## uid property — interface / 232032300130 / 8

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

<a id="canonical-2112101232312101-2303012101021023-0010223221333111-2032112203323101-0311301033333130-1113322213022133-3233320111200020-3300301030231220"></a>

## Next pages — interface / 232032300130 / 9

- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop](data-sources--gcp_vpc_site--reference--group-002.md#canonical-3321032023000021-0000021032023103-2003122132323032-1121100213300200-2112303323013100-2020032023013330-2311103200013321-3231023002332203)
- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-1132222133321101-1232102200201102-2111112102132122-0010023221133123-2333213300031302-1001011022311003-3121103220313010-2103321200303202)

<a id="canonical-1231320312112220-0303322321010321-1320233131012130-0021010330203222-3022101101211110-0231313332332320-1320333301122323-0221133132101110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3020321302000301-2331011321100011-0221332313100230-0202301202211303-2320213002121123-2000333122131310-2330223010222323-0303233321133000"></a>

## ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address — nexthop_address / 120220033323 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-1132222133321101-1232102200201102-2111112102132122-0010023221133123-2333213300031302-1001011022311003-3121103220313010-2103321200303202)
- [Property reference](data-sources--gcp_vpc_site--reference--group-001.md#canonical-0023231003100101-1010210121112300-0303103132302303-0331000323201313-3003311210331022-3000312312212303-0030202320102022-1232013113103201)
- [ingress_egress_gw](data-sources--gcp_vpc_site--reference--group-001.md#canonical-1123313330020103-2220202203231111-2113330331022333-0133213021020321-0303223331023010-1320133103131013-2322232313123200-2330220010011031)
- [ingress_egress_gw.inside_static_routes](data-sources--gcp_vpc_site--reference--group-002.md#canonical-3200313310131232-0210110202101312-2021312210102120-1112102012331303-3023321133320033-3321213133200210-3120212332313312-1233223233330302)
- [ingress_egress_gw.inside_static_routes.static_route_list](data-sources--gcp_vpc_site--reference--group-002.md#canonical-0022032111233021-3032231100003203-2232232213033220-2112100203123200-2101002223201133-1132032110122011-3133130103121231-0111311331112021)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route](data-sources--gcp_vpc_site--reference--group-002.md#canonical-0202301330013012-0202311012013203-0330311200232332-1121021133231003-1023203132102220-0022302011110020-0313200112032320-1320010200100322)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop](data-sources--gcp_vpc_site--reference--group-002.md#canonical-3321032023000021-0000021032023103-2003122132323032-1121100213300200-2112303323013100-2020032023013330-2311103200013321-3231023002332203)
- ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address

<a id="canonical-2101302302032302-0023333230100010-1101023031330232-3121120313210231-3123103132003103-2323310020000122-1131221220032222-1200032122110122"></a>

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

<a id="canonical-0311112332212210-3022123203233230-1013130120111222-2111303300130303-1020322331321223-2133110002213033-0232323322112311-3013322132333013"></a>

## Direct properties — nexthop_address / 120220033323 / 3

- [dual_stack](data-sources--gcp_vpc_site--reference--group-002.md#canonical-3031221200010120-0222222313123331-3230111011310333-1310100030122303-3232012313123330-2313311202102303-0133312301030220-3132112120302221): complete subsection reference.

- [IPv4](data-sources--gcp_vpc_site--reference--group-002.md#canonical-0010133121311210-0122000110002013-1333112202201100-1101233231111011-3023123123322022-1120222313102020-3211200212210101-1022123323213302): complete subsection reference.

- [IPv6](data-sources--gcp_vpc_site--reference--group-002.md#canonical-1133122012002123-0103303012001111-1031221110203121-1222300130313122-3001020313232220-3222202301222001-0200332032311021-3202313233201100): complete subsection reference.

<a id="canonical-3030020213122102-1022303010221031-2232122311303313-2003330120021200-3330231312213103-1111133212220021-2221101132321031-3010002001303211"></a>

## Next pages — nexthop_address / 120220033323 / 4

- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](data-sources--gcp_vpc_site--reference--group-002.md#canonical-3031221200010120-0222222313123331-3230111011310333-1310100030122303-3232012313123330-2313311202102303-0133312301030220-3132112120302221)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4](data-sources--gcp_vpc_site--reference--group-002.md#canonical-0010133121311210-0122000110002013-1333112202201100-1101233231111011-3023123123322022-1120222313102020-3211200212210101-1022123323213302)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6](data-sources--gcp_vpc_site--reference--group-002.md#canonical-1133122012002123-0103303012001111-1031221110203121-1222300130313122-3001020313232220-3222202301222001-0200332032311021-3202313233201100)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop](data-sources--gcp_vpc_site--reference--group-002.md#canonical-3321032023000021-0000021032023103-2003122132323032-1121100213300200-2112303323013100-2020032023013330-2311103200013321-3231023002332203)
- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-1132222133321101-1232102200201102-2111112102132122-0010023221133123-2333213300031302-1001011022311003-3121103220313010-2103321200303202)

<a id="canonical-3031221200010120-0222222313123331-3230111011310333-1310100030122303-3232012313123330-2313311202102303-0133312301030220-3132112120302221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3130120201302131-1010222000010213-1232202301032132-0021212030332001-0110002023020031-1302030103031313-2302030311013203-0222012201313111"></a>

## ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack — dual_stack / 311220121303 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-1132222133321101-1232102200201102-2111112102132122-0010023221133123-2333213300031302-1001011022311003-3121103220313010-2103321200303202)
- [Property reference](data-sources--gcp_vpc_site--reference--group-001.md#canonical-0023231003100101-1010210121112300-0303103132302303-0331000323201313-3003311210331022-3000312312212303-0030202320102022-1232013113103201)
- [ingress_egress_gw](data-sources--gcp_vpc_site--reference--group-001.md#canonical-1123313330020103-2220202203231111-2113330331022333-0133213021020321-0303223331023010-1320133103131013-2322232313123200-2330220010011031)
- [ingress_egress_gw.inside_static_routes](data-sources--gcp_vpc_site--reference--group-002.md#canonical-3200313310131232-0210110202101312-2021312210102120-1112102012331303-3023321133320033-3321213133200210-3120212332313312-1233223233330302)
- [ingress_egress_gw.inside_static_routes.static_route_list](data-sources--gcp_vpc_site--reference--group-002.md#canonical-0022032111233021-3032231100003203-2232232213033220-2112100203123200-2101002223201133-1132032110122011-3133130103121231-0111311331112021)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route](data-sources--gcp_vpc_site--reference--group-002.md#canonical-0202301330013012-0202311012013203-0330311200232332-1121021133231003-1023203132102220-0022302011110020-0313200112032320-1320010200100322)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop](data-sources--gcp_vpc_site--reference--group-002.md#canonical-3321032023000021-0000021032023103-2003122132323032-1121100213300200-2112303323013100-2020032023013330-2311103200013321-3231023002332203)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](data-sources--gcp_vpc_site--reference--group-002.md#canonical-1231320312112220-0303322321010321-1320233131012130-0021010330203222-3022101101211110-0231313332332320-1320333301122323-0221133132101110)
- ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack

<a id="canonical-1033011200001310-3333012230203311-3012010132310121-3303222131230011-3300220121032210-1033131130130103-3130001221212003-2222002032221301"></a>

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

<a id="canonical-0231122233030013-1300032312332232-0322001032323021-0202110201331022-0232113013010111-3120011312100002-1223120220130022-0110212330203123"></a>

## Direct properties — dual_stack / 311220121303 / 3

- [IPv4](data-sources--gcp_vpc_site--reference--group-002.md#canonical-0221013311023100-3213131301320303-3322122223111222-1210220100303132-2033110012110003-2322012322000022-2302031011012013-2302123132322200): complete subsection reference.

- [IPv6](data-sources--gcp_vpc_site--reference--group-002.md#canonical-3011122022122123-2321131031031233-1311023303123123-1313300123322313-3312121010011001-1022110022133300-2303023322001103-2010020032012203): complete subsection reference.

<a id="canonical-0200020100212011-3233032103101130-2120202332133121-1310330103313110-2013213111020300-3112032223233331-3222213321313100-0030232000322133"></a>

## Next pages — dual_stack / 311220121303 / 4

- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4](data-sources--gcp_vpc_site--reference--group-002.md#canonical-0221013311023100-3213131301320303-3322122223111222-1210220100303132-2033110012110003-2322012322000022-2302031011012013-2302123132322200)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6](data-sources--gcp_vpc_site--reference--group-002.md#canonical-3011122022122123-2321131031031233-1311023303123123-1313300123322313-3312121010011001-1022110022133300-2303023322001103-2010020032012203)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](data-sources--gcp_vpc_site--reference--group-002.md#canonical-1231320312112220-0303322321010321-1320233131012130-0021010330203222-3022101101211110-0231313332332320-1320333301122323-0221133132101110)
- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-1132222133321101-1232102200201102-2111112102132122-0010023221133123-2333213300031302-1001011022311003-3121103220313010-2103321200303202)

<a id="canonical-0221013311023100-3213131301320303-3322122223111222-1210220100303132-2033110012110003-2322012322000022-2302031011012013-2302123132322200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3210210223130322-2100330310110102-3121021120200020-2021122102021331-1002001210331020-3211212300102002-3033223123101003-3311303300212002"></a>

## ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.IPv4 — IPv4 / 032003303212 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-1132222133321101-1232102200201102-2111112102132122-0010023221133123-2333213300031302-1001011022311003-3121103220313010-2103321200303202)
- [Property reference](data-sources--gcp_vpc_site--reference--group-001.md#canonical-0023231003100101-1010210121112300-0303103132302303-0331000323201313-3003311210331022-3000312312212303-0030202320102022-1232013113103201)
- [ingress_egress_gw](data-sources--gcp_vpc_site--reference--group-001.md#canonical-1123313330020103-2220202203231111-2113330331022333-0133213021020321-0303223331023010-1320133103131013-2322232313123200-2330220010011031)
- [ingress_egress_gw.inside_static_routes](data-sources--gcp_vpc_site--reference--group-002.md#canonical-3200313310131232-0210110202101312-2021312210102120-1112102012331303-3023321133320033-3321213133200210-3120212332313312-1233223233330302)
- [ingress_egress_gw.inside_static_routes.static_route_list](data-sources--gcp_vpc_site--reference--group-002.md#canonical-0022032111233021-3032231100003203-2232232213033220-2112100203123200-2101002223201133-1132032110122011-3133130103121231-0111311331112021)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route](data-sources--gcp_vpc_site--reference--group-002.md#canonical-0202301330013012-0202311012013203-0330311200232332-1121021133231003-1023203132102220-0022302011110020-0313200112032320-1320010200100322)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop](data-sources--gcp_vpc_site--reference--group-002.md#canonical-3321032023000021-0000021032023103-2003122132323032-1121100213300200-2112303323013100-2020032023013330-2311103200013321-3231023002332203)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](data-sources--gcp_vpc_site--reference--group-002.md#canonical-1231320312112220-0303322321010321-1320233131012130-0021010330203222-3022101101211110-0231313332332320-1320333301122323-0221133132101110)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](data-sources--gcp_vpc_site--reference--group-002.md#canonical-3031221200010120-0222222313123331-3230111011310333-1310100030122303-3232012313123330-2313311202102303-0133312301030220-3132112120302221)
- ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.IPv4

<a id="canonical-3303122120233032-3322013203330123-0323301312030022-3122001022033232-3031320121030102-0323220212030222-3301123220023020-1032312322111333"></a>

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

<a id="canonical-0022130112320012-2001320310101012-1311013112222022-2310121030132210-2232123023030212-0123121110323300-0000223003032013-2133013311031201"></a>

## Direct properties — IPv4 / 032003303212 / 3

<a id="canonical-1110022101220222-3033312102022023-2232132120332302-3323330313311111-0322102233121301-1032211213031213-3321120120001100-1313303210001003"></a>

<a id="canonical-3322000000132201-0123122022101110-3020231230233133-2303320121103210-2222002332322302-0312221230323032-1020223102220202-1302123021233320"></a>

## addr property — IPv4 / 032003303212 / 4

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
    "ves.io.schema.rules.string.ipv4": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv4": "true"
  }
}
```

<a id="canonical-1120300321132312-2020323001203313-1033003231303030-3030320300111310-2130033212122201-2313120333202210-3130303223001220-0101012021332320"></a>

## Next pages — IPv4 / 032003303212 / 5

- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](data-sources--gcp_vpc_site--reference--group-002.md#canonical-3031221200010120-0222222313123331-3230111011310333-1310100030122303-3232012313123330-2313311202102303-0133312301030220-3132112120302221)
- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-1132222133321101-1232102200201102-2111112102132122-0010023221133123-2333213300031302-1001011022311003-3121103220313010-2103321200303202)

<a id="canonical-3011122022122123-2321131031031233-1311023303123123-1313300123322313-3312121010011001-1022110022133300-2303023322001103-2010020032012203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2310000332301112-1101002330121013-0220301130011000-1233302311102031-1231312212011112-2312203232132032-0123133311300320-0320010301013021"></a>

## ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.IPv6 — IPv6 / 300322112000 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-1132222133321101-1232102200201102-2111112102132122-0010023221133123-2333213300031302-1001011022311003-3121103220313010-2103321200303202)
- [Property reference](data-sources--gcp_vpc_site--reference--group-001.md#canonical-0023231003100101-1010210121112300-0303103132302303-0331000323201313-3003311210331022-3000312312212303-0030202320102022-1232013113103201)
- [ingress_egress_gw](data-sources--gcp_vpc_site--reference--group-001.md#canonical-1123313330020103-2220202203231111-2113330331022333-0133213021020321-0303223331023010-1320133103131013-2322232313123200-2330220010011031)
- [ingress_egress_gw.inside_static_routes](data-sources--gcp_vpc_site--reference--group-002.md#canonical-3200313310131232-0210110202101312-2021312210102120-1112102012331303-3023321133320033-3321213133200210-3120212332313312-1233223233330302)
- [ingress_egress_gw.inside_static_routes.static_route_list](data-sources--gcp_vpc_site--reference--group-002.md#canonical-0022032111233021-3032231100003203-2232232213033220-2112100203123200-2101002223201133-1132032110122011-3133130103121231-0111311331112021)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route](data-sources--gcp_vpc_site--reference--group-002.md#canonical-0202301330013012-0202311012013203-0330311200232332-1121021133231003-1023203132102220-0022302011110020-0313200112032320-1320010200100322)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop](data-sources--gcp_vpc_site--reference--group-002.md#canonical-3321032023000021-0000021032023103-2003122132323032-1121100213300200-2112303323013100-2020032023013330-2311103200013321-3231023002332203)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](data-sources--gcp_vpc_site--reference--group-002.md#canonical-1231320312112220-0303322321010321-1320233131012130-0021010330203222-3022101101211110-0231313332332320-1320333301122323-0221133132101110)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](data-sources--gcp_vpc_site--reference--group-002.md#canonical-3031221200010120-0222222313123331-3230111011310333-1310100030122303-3232012313123330-2313311202102303-0133312301030220-3132112120302221)
- ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.IPv6

<a id="canonical-3321030110132221-2323320002131110-2300300103222330-0223302200200111-3120232020021020-1310231301330003-0023221312302002-3203331132012322"></a>

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

<a id="canonical-0312321131000302-1231110203302030-3230133220022303-1013030233102132-2333122313333032-0023031010222100-1003110332101120-1120303003201002"></a>

## Direct properties — IPv6 / 300322112000 / 3

<a id="canonical-0123132033211033-2102213201101022-1033101021130302-0023001313132102-0212301123023323-2111100031323113-0013001001002302-1132003201331013"></a>

<a id="canonical-2023100111300301-0231301003110211-2031000123130230-0313332210013232-0221020021131021-1200023313221133-2203333223020213-1222203302212121"></a>

## addr property — IPv6 / 300322112000 / 4

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
    "ves.io.schema.rules.string.ipv6": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv6": "true"
  }
}
```

<a id="canonical-0313013333220302-3330103223201131-3000211233331311-3110201300302232-1321103203213033-1013311020300231-0302221012131303-0001012332313202"></a>

## Next pages — IPv6 / 300322112000 / 5

- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](data-sources--gcp_vpc_site--reference--group-002.md#canonical-3031221200010120-0222222313123331-3230111011310333-1310100030122303-3232012313123330-2313311202102303-0133312301030220-3132112120302221)
- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-1132222133321101-1232102200201102-2111112102132122-0010023221133123-2333213300031302-1001011022311003-3121103220313010-2103321200303202)

<a id="canonical-0010133121311210-0122000110002013-1333112202201100-1101233231111011-3023123123322022-1120222313102020-3211200212210101-1022123323213302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3323302022220021-0331200330310031-2203311122111030-1102211003110320-2002113010033110-2303120332103230-1331232230131133-0200000311300221"></a>

## ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.IPv4 — IPv4 / 100023121201 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-1132222133321101-1232102200201102-2111112102132122-0010023221133123-2333213300031302-1001011022311003-3121103220313010-2103321200303202)
- [Property reference](data-sources--gcp_vpc_site--reference--group-001.md#canonical-0023231003100101-1010210121112300-0303103132302303-0331000323201313-3003311210331022-3000312312212303-0030202320102022-1232013113103201)
- [ingress_egress_gw](data-sources--gcp_vpc_site--reference--group-001.md#canonical-1123313330020103-2220202203231111-2113330331022333-0133213021020321-0303223331023010-1320133103131013-2322232313123200-2330220010011031)
- [ingress_egress_gw.inside_static_routes](data-sources--gcp_vpc_site--reference--group-002.md#canonical-3200313310131232-0210110202101312-2021312210102120-1112102012331303-3023321133320033-3321213133200210-3120212332313312-1233223233330302)
- [ingress_egress_gw.inside_static_routes.static_route_list](data-sources--gcp_vpc_site--reference--group-002.md#canonical-0022032111233021-3032231100003203-2232232213033220-2112100203123200-2101002223201133-1132032110122011-3133130103121231-0111311331112021)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route](data-sources--gcp_vpc_site--reference--group-002.md#canonical-0202301330013012-0202311012013203-0330311200232332-1121021133231003-1023203132102220-0022302011110020-0313200112032320-1320010200100322)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop](data-sources--gcp_vpc_site--reference--group-002.md#canonical-3321032023000021-0000021032023103-2003122132323032-1121100213300200-2112303323013100-2020032023013330-2311103200013321-3231023002332203)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](data-sources--gcp_vpc_site--reference--group-002.md#canonical-1231320312112220-0303322321010321-1320233131012130-0021010330203222-3022101101211110-0231313332332320-1320333301122323-0221133132101110)
- ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.IPv4

<a id="canonical-2122331232333120-1301020213313032-3000312000232000-1101132132200302-3220012333223301-1122121003113303-2122122010022111-3113122203023113"></a>

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

<a id="canonical-2003221010012011-0030213303023031-3012221212211211-3232012121301102-2033221223011103-2331302232012001-1321222131300232-3233221122201203"></a>

## Direct properties — IPv4 / 100023121201 / 3

<a id="canonical-2321010222123213-3133001222120220-0113100120133012-2202010200213130-1300010031312122-1213002002233230-3032113112333300-0232130312131320"></a>

<a id="canonical-3333301112133132-1220221012010230-2330100232300313-3331121033120231-3311310322300012-3133331132302203-1030102100102223-1210110120130333"></a>

## addr property — IPv4 / 100023121201 / 4

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
    "ves.io.schema.rules.string.ipv4": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv4": "true"
  }
}
```

<a id="canonical-2132220210121112-1313213121230332-3010303011331102-1323032112321133-3323232222133030-0322132030101333-3303221103322031-3203111300123011"></a>

## Next pages — IPv4 / 100023121201 / 5

- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](data-sources--gcp_vpc_site--reference--group-002.md#canonical-1231320312112220-0303322321010321-1320233131012130-0021010330203222-3022101101211110-0231313332332320-1320333301122323-0221133132101110)
- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-1132222133321101-1232102200201102-2111112102132122-0010023221133123-2333213300031302-1001011022311003-3121103220313010-2103321200303202)

<a id="canonical-1133122012002123-0103303012001111-1031221110203121-1222300130313122-3001020313232220-3222202301222001-0200332032311021-3202313233201100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2121023232132221-3301313301230121-0103011123200000-0301212330331002-2313022123302000-0232223330013032-1122002222331212-0103011302302033"></a>

## ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.IPv6 — IPv6 / 202302331301 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-1132222133321101-1232102200201102-2111112102132122-0010023221133123-2333213300031302-1001011022311003-3121103220313010-2103321200303202)
- [Property reference](data-sources--gcp_vpc_site--reference--group-001.md#canonical-0023231003100101-1010210121112300-0303103132302303-0331000323201313-3003311210331022-3000312312212303-0030202320102022-1232013113103201)
- [ingress_egress_gw](data-sources--gcp_vpc_site--reference--group-001.md#canonical-1123313330020103-2220202203231111-2113330331022333-0133213021020321-0303223331023010-1320133103131013-2322232313123200-2330220010011031)
- [ingress_egress_gw.inside_static_routes](data-sources--gcp_vpc_site--reference--group-002.md#canonical-3200313310131232-0210110202101312-2021312210102120-1112102012331303-3023321133320033-3321213133200210-3120212332313312-1233223233330302)
- [ingress_egress_gw.inside_static_routes.static_route_list](data-sources--gcp_vpc_site--reference--group-002.md#canonical-0022032111233021-3032231100003203-2232232213033220-2112100203123200-2101002223201133-1132032110122011-3133130103121231-0111311331112021)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route](data-sources--gcp_vpc_site--reference--group-002.md#canonical-0202301330013012-0202311012013203-0330311200232332-1121021133231003-1023203132102220-0022302011110020-0313200112032320-1320010200100322)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop](data-sources--gcp_vpc_site--reference--group-002.md#canonical-3321032023000021-0000021032023103-2003122132323032-1121100213300200-2112303323013100-2020032023013330-2311103200013321-3231023002332203)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](data-sources--gcp_vpc_site--reference--group-002.md#canonical-1231320312112220-0303322321010321-1320233131012130-0021010330203222-3022101101211110-0231313332332320-1320333301122323-0221133132101110)
- ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.IPv6

<a id="canonical-0303023121120203-3133300221131122-1212203103003312-1231122032021012-1210311111110020-1132112313130123-3200220313120220-3300030112220132"></a>

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

<a id="canonical-0020122321320111-1213230330011013-3020021023310002-0233211133223032-0011322020300220-3232110110232203-0313030012110122-1120123011231330"></a>

## Direct properties — IPv6 / 202302331301 / 3

<a id="canonical-3320020103101312-0111112122320112-1310311213310220-1213112022132020-2232012320233101-3030210230101201-0220102103231210-3200101101211003"></a>

<a id="canonical-0232133310021122-0300112333321032-3233320102013331-3031033232203000-1212220313022130-0203121331330002-1100112101011000-0111103001210113"></a>

## addr property — IPv6 / 202302331301 / 4

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
    "ves.io.schema.rules.string.ipv6": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv6": "true"
  }
}
```

<a id="canonical-2130103120201333-2322132302330230-3121232310000111-1332110302322203-2031010033031323-1200321230210113-2131311132313221-0233303323102323"></a>

## Next pages — IPv6 / 202302331301 / 5

- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](data-sources--gcp_vpc_site--reference--group-002.md#canonical-1231320312112220-0303322321010321-1320233131012130-0021010330203222-3022101101211110-0231313332332320-1320333301122323-0221133132101110)
- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-1132222133321101-1232102200201102-2111112102132122-0010023221133123-2333213300031302-1001011022311003-3121103220313010-2103321200303202)

<a id="canonical-2122221323130301-1313122121231331-0300323002000300-1301230032030102-2023320221320212-3221013211122031-2220012333303310-2011001021111030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3213331220010111-0320213321230201-0301130000003023-0200003220313321-2020000123220021-0130300101000000-1001301300202311-2300103032122333"></a>

## ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets — subnets / 310313222313 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-1132222133321101-1232102200201102-2111112102132122-0010023221133123-2333213300031302-1001011022311003-3121103220313010-2103321200303202)
- [Property reference](data-sources--gcp_vpc_site--reference--group-001.md#canonical-0023231003100101-1010210121112300-0303103132302303-0331000323201313-3003311210331022-3000312312212303-0030202320102022-1232013113103201)
- [ingress_egress_gw](data-sources--gcp_vpc_site--reference--group-001.md#canonical-1123313330020103-2220202203231111-2113330331022333-0133213021020321-0303223331023010-1320133103131013-2322232313123200-2330220010011031)
- [ingress_egress_gw.inside_static_routes](data-sources--gcp_vpc_site--reference--group-002.md#canonical-3200313310131232-0210110202101312-2021312210102120-1112102012331303-3023321133320033-3321213133200210-3120212332313312-1233223233330302)
- [ingress_egress_gw.inside_static_routes.static_route_list](data-sources--gcp_vpc_site--reference--group-002.md#canonical-0022032111233021-3032231100003203-2232232213033220-2112100203123200-2101002223201133-1132032110122011-3133130103121231-0111311331112021)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route](data-sources--gcp_vpc_site--reference--group-002.md#canonical-0202301330013012-0202311012013203-0330311200232332-1121021133231003-1023203132102220-0022302011110020-0313200112032320-1320010200100322)
- ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets

<a id="canonical-2130122032312300-2232300321103002-1021101122010022-3021330113131100-0330011210312122-1300001020312011-0313132000032310-3013302001010330"></a>

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
    "ves.io.schema.rules.repeated.max_items": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "256"
  }
}
```

<a id="canonical-3231232110100010-0221101320012033-3331232331122111-2200103301230330-0302112320211023-2101223011213032-2022103003313013-0013200013230102"></a>

## Direct properties — subnets / 310313222313 / 3

- [IPv4](data-sources--gcp_vpc_site--reference--group-002.md#canonical-0311320101010001-1301332130030223-0200020203232222-1202333330033222-1310212301332023-2320310312031331-2221102100323313-2301112120313123): complete subsection reference.

- [IPv6](data-sources--gcp_vpc_site--reference--group-002.md#canonical-1313002130212221-3132103001230130-1332230320021300-1303132001032302-0100122220111233-2033331321302112-0320200331111031-1020330101200323): complete subsection reference.

<a id="canonical-3002100001211302-0022012133331123-0133102322022323-0312201011223121-0213332003132003-0300310222000022-0103100213132113-2221232211102233"></a>

## Next pages — subnets / 310313222313 / 4

- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets.ipv4](data-sources--gcp_vpc_site--reference--group-002.md#canonical-0311320101010001-1301332130030223-0200020203232222-1202333330033222-1310212301332023-2320310312031331-2221102100323313-2301112120313123)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets.ipv6](data-sources--gcp_vpc_site--reference--group-002.md#canonical-1313002130212221-3132103001230130-1332230320021300-1303132001032302-0100122220111233-2033331321302112-0320200331111031-1020330101200323)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route](data-sources--gcp_vpc_site--reference--group-002.md#canonical-0202301330013012-0202311012013203-0330311200232332-1121021133231003-1023203132102220-0022302011110020-0313200112032320-1320010200100322)
- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-1132222133321101-1232102200201102-2111112102132122-0010023221133123-2333213300031302-1001011022311003-3121103220313010-2103321200303202)

<a id="canonical-0311320101010001-1301332130030223-0200020203232222-1202333330033222-1310212301332023-2320310312031331-2221102100323313-2301112120313123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2300130323331111-0203232203122301-3113122112322303-0200221331323103-0311321023211020-1132211311322003-0123230120302120-1100303022310031"></a>

## ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets.IPv4 — IPv4 / 313103200331 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-1132222133321101-1232102200201102-2111112102132122-0010023221133123-2333213300031302-1001011022311003-3121103220313010-2103321200303202)
- [Property reference](data-sources--gcp_vpc_site--reference--group-001.md#canonical-0023231003100101-1010210121112300-0303103132302303-0331000323201313-3003311210331022-3000312312212303-0030202320102022-1232013113103201)
- [ingress_egress_gw](data-sources--gcp_vpc_site--reference--group-001.md#canonical-1123313330020103-2220202203231111-2113330331022333-0133213021020321-0303223331023010-1320133103131013-2322232313123200-2330220010011031)
- [ingress_egress_gw.inside_static_routes](data-sources--gcp_vpc_site--reference--group-002.md#canonical-3200313310131232-0210110202101312-2021312210102120-1112102012331303-3023321133320033-3321213133200210-3120212332313312-1233223233330302)
- [ingress_egress_gw.inside_static_routes.static_route_list](data-sources--gcp_vpc_site--reference--group-002.md#canonical-0022032111233021-3032231100003203-2232232213033220-2112100203123200-2101002223201133-1132032110122011-3133130103121231-0111311331112021)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route](data-sources--gcp_vpc_site--reference--group-002.md#canonical-0202301330013012-0202311012013203-0330311200232332-1121021133231003-1023203132102220-0022302011110020-0313200112032320-1320010200100322)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets](data-sources--gcp_vpc_site--reference--group-002.md#canonical-2122221323130301-1313122121231331-0300323002000300-1301230032030102-2023320221320212-3221013211122031-2220012333303310-2011001021111030)
- ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets.IPv4

<a id="canonical-0330322031201032-1203302302103133-3121121012303102-0022230212030033-2100333030110321-2131301123013323-2012022312202212-2202113321112022"></a>

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

<a id="canonical-3321302220210202-0221310000023102-0033303300233023-2002112303013312-0122203332331133-0102122133031211-0100020233101323-1333303301023300"></a>

## Direct properties — IPv4 / 313103200331 / 3

<a id="canonical-0320203021023111-1330322232221311-1010312023031122-0212233301321300-3323002021210030-3011330222122133-0020000320300033-3123230001323033"></a>

<a id="canonical-1310201212312100-0300023102021130-2212312213100131-0010010210133012-3330020023212331-1003000111120303-3333021113131323-0033322111313111"></a>

## plen property — IPv4 / 313103200331 / 4

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
    "ves.io.schema.rules.uint32.lte": "32"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "32"
  }
}
```

<a id="canonical-2201232223233230-2022021332230231-3201123323101101-1112202003122030-1023210113120023-0132033001112320-0102120233311100-1113301131100323"></a>

<a id="canonical-2322031210213330-0320030301011011-2202022220213211-1310303330112310-0211132110311213-3230010022213231-2333110311232220-2021310232013013"></a>

## prefix property — IPv4 / 313103200331 / 5

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
    "ves.io.schema.rules.string.ipv4": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv4": "true"
  }
}
```

<a id="canonical-1230211312022201-1022013222311220-1331113120120303-1223221113011002-2101303020022122-3011323230030113-2110222101112030-1001112133231211"></a>

## Next pages — IPv4 / 313103200331 / 6

- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets](data-sources--gcp_vpc_site--reference--group-002.md#canonical-2122221323130301-1313122121231331-0300323002000300-1301230032030102-2023320221320212-3221013211122031-2220012333303310-2011001021111030)
- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-1132222133321101-1232102200201102-2111112102132122-0010023221133123-2333213300031302-1001011022311003-3121103220313010-2103321200303202)

<a id="canonical-1313002130212221-3132103001230130-1332230320021300-1303132001032302-0100122220111233-2033331321302112-0320200331111031-1020330101200323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0033031020033321-2223033312332021-3121012010210222-1130132313021311-2330330010310103-0021223112313022-0011113203010033-2131102113120323"></a>

## ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets.IPv6 — IPv6 / 030000220031 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-1132222133321101-1232102200201102-2111112102132122-0010023221133123-2333213300031302-1001011022311003-3121103220313010-2103321200303202)
- [Property reference](data-sources--gcp_vpc_site--reference--group-001.md#canonical-0023231003100101-1010210121112300-0303103132302303-0331000323201313-3003311210331022-3000312312212303-0030202320102022-1232013113103201)
- [ingress_egress_gw](data-sources--gcp_vpc_site--reference--group-001.md#canonical-1123313330020103-2220202203231111-2113330331022333-0133213021020321-0303223331023010-1320133103131013-2322232313123200-2330220010011031)
- [ingress_egress_gw.inside_static_routes](data-sources--gcp_vpc_site--reference--group-002.md#canonical-3200313310131232-0210110202101312-2021312210102120-1112102012331303-3023321133320033-3321213133200210-3120212332313312-1233223233330302)
- [ingress_egress_gw.inside_static_routes.static_route_list](data-sources--gcp_vpc_site--reference--group-002.md#canonical-0022032111233021-3032231100003203-2232232213033220-2112100203123200-2101002223201133-1132032110122011-3133130103121231-0111311331112021)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route](data-sources--gcp_vpc_site--reference--group-002.md#canonical-0202301330013012-0202311012013203-0330311200232332-1121021133231003-1023203132102220-0022302011110020-0313200112032320-1320010200100322)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets](data-sources--gcp_vpc_site--reference--group-002.md#canonical-2122221323130301-1313122121231331-0300323002000300-1301230032030102-2023320221320212-3221013211122031-2220012333303310-2011001021111030)
- ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets.IPv6

<a id="canonical-2213232323302103-0301332323231211-0301230201030020-1010000023011231-3301120231010332-3022011032020031-3032122130302021-3302313122002332"></a>

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

<a id="canonical-3101132130112210-3031133131322231-2232311120023311-3220030030110020-3032213100320103-2011211001233212-1311110333021001-1201213333221110"></a>

## Direct properties — IPv6 / 030000220031 / 3

<a id="canonical-2012322032030221-3023110311121130-2103332020202021-0102121320122110-1000112303010222-3133332113233111-2133111031123110-3103302101002210"></a>

<a id="canonical-1131331322330220-1301033213232300-2120001002031201-1210330120121202-3321010000230010-2022002133133012-0110201313300110-1020322221332033"></a>

## plen property — IPv6 / 030000220031 / 4

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
    "ves.io.schema.rules.uint32.lte": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "128"
  }
}
```

<a id="canonical-1133321320302013-0303231112130101-0023111002133310-3010233330130220-2130221130002322-1013110101210212-1230100322203030-2221110112332011"></a>

<a id="canonical-0011213331111102-1032001332312131-3233110211111220-3233033220202012-2321123001033000-0123333023012312-1120030332220332-3330031221220002"></a>

## prefix property — IPv6 / 030000220031 / 5

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
    "ves.io.schema.rules.string.ipv6": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv6": "true"
  }
}
```

<a id="canonical-2031022013011331-3003133300300321-1311312120211323-1312221011213222-0102020102101210-1303131103120121-0000231022222321-3121310001012211"></a>

## Next pages — IPv6 / 030000220031 / 6

- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets](data-sources--gcp_vpc_site--reference--group-002.md#canonical-2122221323130301-1313122121231331-0300323002000300-1301230032030102-2023320221320212-3221013211122031-2220012333303310-2011001021111030)
- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-1132222133321101-1232102200201102-2111112102132122-0010023221133123-2333213300031302-1001011022311003-3121103220313010-2103321200303202)

<a id="canonical-3203003222300230-2112031310101033-3303130321221222-3231311000312033-1330123121120320-1222230233122203-2132222101023230-1023322133302111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2021011031333112-1101301330030030-3112000330102113-2001002230000031-1011132022122021-0023210130031123-1110230210021013-2230203120221323"></a>

## ingress_egress_gw.inside_subnet — inside_subnet / 113323333212 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-1132222133321101-1232102200201102-2111112102132122-0010023221133123-2333213300031302-1001011022311003-3121103220313010-2103321200303202)
- [Property reference](data-sources--gcp_vpc_site--reference--group-001.md#canonical-0023231003100101-1010210121112300-0303103132302303-0331000323201313-3003311210331022-3000312312212303-0030202320102022-1232013113103201)
- [ingress_egress_gw](data-sources--gcp_vpc_site--reference--group-001.md#canonical-1123313330020103-2220202203231111-2113330331022333-0133213021020321-0303223331023010-1320133103131013-2322232313123200-2330220010011031)
- ingress_egress_gw.inside_subnet

<a id="canonical-3231330330333310-0233032200002020-0320132032223232-2212122102321221-0221013303310012-0220130000213120-0002033021023220-0320002213320033"></a>

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

<a id="canonical-2120023110312122-1023331032122121-2211312100332000-1023213110110223-0111303210222001-0010202101213023-1210132310220033-1102100230313130"></a>

## Direct properties — inside_subnet / 113323333212 / 3

- [existing_subnet](data-sources--gcp_vpc_site--reference--group-002.md#canonical-2330230233200303-0231231323111122-3002213011113101-0002232333230230-2322331121112021-3031313101211211-1003320010102100-1130323313111221): complete subsection reference.

- [new_subnet](data-sources--gcp_vpc_site--reference--group-002.md#canonical-0303200223110012-1013232231211022-2031330330120312-0013112211313230-2113300233210320-3302103030332301-3330120131321121-0113122002133110): complete subsection reference.

<a id="canonical-0212012102033312-0332320100212331-3220101312322320-1223213321001313-2312101312222202-2133311032110133-1010110232233201-3231211023310032"></a>

## Next pages — inside_subnet / 113323333212 / 4

- [ingress_egress_gw.inside_subnet.existing_subnet](data-sources--gcp_vpc_site--reference--group-002.md#canonical-2330230233200303-0231231323111122-3002213011113101-0002232333230230-2322331121112021-3031313101211211-1003320010102100-1130323313111221)
- [ingress_egress_gw.inside_subnet.new_subnet](data-sources--gcp_vpc_site--reference--group-002.md#canonical-0303200223110012-1013232231211022-2031330330120312-0013112211313230-2113300233210320-3302103030332301-3330120131321121-0113122002133110)
- [ingress_egress_gw](data-sources--gcp_vpc_site--reference--group-001.md#canonical-1123313330020103-2220202203231111-2113330331022333-0133213021020321-0303223331023010-1320133103131013-2322232313123200-2330220010011031)
- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-1132222133321101-1232102200201102-2111112102132122-0010023221133123-2333213300031302-1001011022311003-3121103220313010-2103321200303202)

<a id="canonical-2330230233200303-0231231323111122-3002213011113101-0002232333230230-2322331121112021-3031313101211211-1003320010102100-1130323313111221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3310101011321001-2203333310030222-2320110132022101-3310120323320332-2200223112231122-3333102332001313-2121023212313003-1222311323133333"></a>

## ingress_egress_gw.inside_subnet.existing_subnet — existing_subnet / 310202301023 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-1132222133321101-1232102200201102-2111112102132122-0010023221133123-2333213300031302-1001011022311003-3121103220313010-2103321200303202)
- [Property reference](data-sources--gcp_vpc_site--reference--group-001.md#canonical-0023231003100101-1010210121112300-0303103132302303-0331000323201313-3003311210331022-3000312312212303-0030202320102022-1232013113103201)
- [ingress_egress_gw](data-sources--gcp_vpc_site--reference--group-001.md#canonical-1123313330020103-2220202203231111-2113330331022333-0133213021020321-0303223331023010-1320133103131013-2322232313123200-2330220010011031)
- [ingress_egress_gw.inside_subnet](data-sources--gcp_vpc_site--reference--group-002.md#canonical-3203003222300230-2112031310101033-3303130321221222-3231311000312033-1330123121120320-1222230233122203-2132222101023230-1023322133302111)
- ingress_egress_gw.inside_subnet.existing_subnet

<a id="canonical-2301010221110120-0200203121013312-2221131212313031-3012103021213312-1110123012220110-0011320331230232-0123110132033331-3313130101022001"></a>

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

<a id="canonical-1032210231110110-2311113201133123-3020331312011111-2023203202011310-1213201111003103-1101330202312310-3010312313210011-1130333011000002"></a>

## Direct properties — existing_subnet / 310202301023 / 3

<a id="canonical-1022023302023012-2032230123103010-2012123203020331-1123022132030233-0313101232201030-1201023201320102-1131031333100321-0330322033323202"></a>

<a id="canonical-3003103231022122-0313211130233030-3021212133023310-1001200221000003-1102100133212020-0013131033030330-3303302302222332-1011221100033023"></a>

## subnet_name property — existing_subnet / 310202301023 / 4

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

<a id="canonical-3013331133020311-2312211310210031-1322103111210210-1322110110122010-1302333011021121-1121021302101232-0031001231212221-3021232010023312"></a>

## Next pages — existing_subnet / 310202301023 / 5

- [ingress_egress_gw.inside_subnet](data-sources--gcp_vpc_site--reference--group-002.md#canonical-3203003222300230-2112031310101033-3303130321221222-3231311000312033-1330123121120320-1222230233122203-2132222101023230-1023322133302111)
- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-1132222133321101-1232102200201102-2111112102132122-0010023221133123-2333213300031302-1001011022311003-3121103220313010-2103321200303202)

<a id="canonical-0303200223110012-1013232231211022-2031330330120312-0013112211313230-2113300233210320-3302103030332301-3330120131321121-0113122002133110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1010000212303211-3232123103121300-2101113303302020-2013012300313122-2022230330131033-2200131030130020-0032301031333012-0010001321012030"></a>

## ingress_egress_gw.inside_subnet.new_subnet — new_subnet / 303123200203 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-1132222133321101-1232102200201102-2111112102132122-0010023221133123-2333213300031302-1001011022311003-3121103220313010-2103321200303202)
- [Property reference](data-sources--gcp_vpc_site--reference--group-001.md#canonical-0023231003100101-1010210121112300-0303103132302303-0331000323201313-3003311210331022-3000312312212303-0030202320102022-1232013113103201)
- [ingress_egress_gw](data-sources--gcp_vpc_site--reference--group-001.md#canonical-1123313330020103-2220202203231111-2113330331022333-0133213021020321-0303223331023010-1320133103131013-2322232313123200-2330220010011031)
- [ingress_egress_gw.inside_subnet](data-sources--gcp_vpc_site--reference--group-002.md#canonical-3203003222300230-2112031310101033-3303130321221222-3231311000312033-1330123121120320-1222230233122203-2132222101023230-1023322133302111)
- ingress_egress_gw.inside_subnet.new_subnet

<a id="canonical-0132332133032322-0321320111010312-2322110013321011-2321313320321300-2323113101133030-2230123030320302-3111321110321330-1010031031122332"></a>

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

<a id="canonical-2032002312022233-2011011023030123-3032112020000100-2132313023001330-0021132320322121-3030221200231330-3220103203233313-3000122133302302"></a>

## Direct properties — new_subnet / 303123200203 / 3

<a id="canonical-0133312313130312-0323101122201131-2103112122303220-1022002030201201-2002323230013320-3112323232102122-2213100113313033-1200013322331010"></a>

<a id="canonical-2310330330222322-0103103133033033-1033111032210103-0321101300202221-3200333232030022-3032121032131100-3011123312200233-3322310120320123"></a>

## primary_ipv4 property — new_subnet / 303123200203 / 4

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

<a id="canonical-2220123323130313-2333212230331320-2021103203222331-1103033012112311-0112220123100202-2113310312013321-1012232200002013-3222030321212200"></a>

<a id="canonical-2300011003300223-0203203201002333-0010201120232201-3132332121022020-2011212001030033-2011030000023303-0302003023123103-2021122203110023"></a>

## subnet_name property — new_subnet / 303123200203 / 5

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
    "ves.io.schema.rules.string.max_len": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64"
  }
}
```

<a id="canonical-0212233022020231-2121010232033133-1121200222312012-3223330000311231-2231000031121102-0003123120300221-3312230030023013-0110232331321103"></a>

## Next pages — new_subnet / 303123200203 / 6

- [ingress_egress_gw.inside_subnet](data-sources--gcp_vpc_site--reference--group-002.md#canonical-3203003222300230-2112031310101033-3303130321221222-3231311000312033-1330123121120320-1222230233122203-2132222101023230-1023322133302111)
- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-1132222133321101-1232102200201102-2111112102132122-0010023221133123-2333213300031302-1001011022311003-3121103220313010-2103321200303202)

<a id="canonical-0202032310301232-3202321000020000-0210202121112010-2120030213110333-3332322130201230-2131302013000120-3331220230110321-1232320231201231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1220310010032021-1311101132320112-2211123003031100-0120131322010021-2023302303212233-1321202101223103-3121123112300221-3010001013313300"></a>

## ingress_egress_gw.no_dc_cluster_group — no_dc_cluster_group / 222021230210 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-1132222133321101-1232102200201102-2111112102132122-0010023221133123-2333213300031302-1001011022311003-3121103220313010-2103321200303202)
- [Property reference](data-sources--gcp_vpc_site--reference--group-001.md#canonical-0023231003100101-1010210121112300-0303103132302303-0331000323201313-3003311210331022-3000312312212303-0030202320102022-1232013113103201)
- [ingress_egress_gw](data-sources--gcp_vpc_site--reference--group-001.md#canonical-1123313330020103-2220202203231111-2113330331022333-0133213021020321-0303223331023010-1320133103131013-2322232313123200-2330220010011031)
- ingress_egress_gw.no_dc_cluster_group

<a id="canonical-3213110011303132-3311232203201033-2303223001021321-0101223113331223-2001002021102011-1102300113101320-0303102000033120-2311322132130203"></a>

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

<a id="canonical-2111113212011101-2131313131211332-1333333212231131-3322221033000331-2013320113103103-1212030013031103-3221112110121012-0003210311222031"></a>

## Direct properties — no_dc_cluster_group / 222021230210 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0213000032230130-3311103133112311-3303200223030232-1023300223003133-3232203122213002-1223313023330020-0113110233230300-0031022322231320"></a>

## Next pages — no_dc_cluster_group / 222021230210 / 4

- [ingress_egress_gw](data-sources--gcp_vpc_site--reference--group-001.md#canonical-1123313330020103-2220202203231111-2113330331022333-0133213021020321-0303223331023010-1320133103131013-2322232313123200-2330220010011031)
- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-1132222133321101-1232102200201102-2111112102132122-0010023221133123-2333213300031302-1001011022311003-3121103220313010-2103321200303202)

<a id="canonical-0032301122112030-0230102021011322-1331221201233202-2310133101313001-3300103223100011-2101033003202031-3113223310023220-0023332112202012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3302112021312010-2321330231331220-1011030203020313-3131322111033032-0133231312313322-3231210012210200-3020122331022220-1002202321033213"></a>

## ingress_egress_gw.no_forward_proxy — no_forward_proxy / 003203331333 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-1132222133321101-1232102200201102-2111112102132122-0010023221133123-2333213300031302-1001011022311003-3121103220313010-2103321200303202)
- [Property reference](data-sources--gcp_vpc_site--reference--group-001.md#canonical-0023231003100101-1010210121112300-0303103132302303-0331000323201313-3003311210331022-3000312312212303-0030202320102022-1232013113103201)
- [ingress_egress_gw](data-sources--gcp_vpc_site--reference--group-001.md#canonical-1123313330020103-2220202203231111-2113330331022333-0133213021020321-0303223331023010-1320133103131013-2322232313123200-2330220010011031)
- ingress_egress_gw.no_forward_proxy

<a id="canonical-1030111330000031-1013102311022312-1203213012221121-3010200202220312-0003120110223320-2123000223213321-2010211330110032-1010021212130310"></a>

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

<a id="canonical-2200001223122002-0301101032301003-1211100012321113-3331223031313031-0111022113333010-0201110202122303-0113322003001003-0322202103130303"></a>

## Direct properties — no_forward_proxy / 003203331333 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1203132221300222-2201001010313203-0133322131330102-0012022111203132-3201323001301133-2001113320012303-2211031020023303-3031230002332010"></a>

## Next pages — no_forward_proxy / 003203331333 / 4

- [ingress_egress_gw](data-sources--gcp_vpc_site--reference--group-001.md#canonical-1123313330020103-2220202203231111-2113330331022333-0133213021020321-0303223331023010-1320133103131013-2322232313123200-2330220010011031)
- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-1132222133321101-1232102200201102-2111112102132122-0010023221133123-2333213300031302-1001011022311003-3121103220313010-2103321200303202)

<a id="canonical-1220311131221300-0120302031132331-1130330201130100-0013112232110113-3210123101120022-2312223312233301-1213032212023333-1001202002130112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2231322210000303-2302121213320210-0230133330230301-0300032022100100-3301010231103302-1223331212231211-0001220111220131-2231301222201223"></a>

## ingress_egress_gw.no_global_network — no_global_network / 322200003102 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-1132222133321101-1232102200201102-2111112102132122-0010023221133123-2333213300031302-1001011022311003-3121103220313010-2103321200303202)
- [Property reference](data-sources--gcp_vpc_site--reference--group-001.md#canonical-0023231003100101-1010210121112300-0303103132302303-0331000323201313-3003311210331022-3000312312212303-0030202320102022-1232013113103201)
- [ingress_egress_gw](data-sources--gcp_vpc_site--reference--group-001.md#canonical-1123313330020103-2220202203231111-2113330331022333-0133213021020321-0303223331023010-1320133103131013-2322232313123200-2330220010011031)
- ingress_egress_gw.no_global_network

<a id="canonical-3203122321121030-0331210203131120-2310201231021033-0113022301301030-3010231123122331-3323121031332222-0330301302031332-2320212003133213"></a>

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

<a id="canonical-2330332030001111-3010231222021102-0102130100213201-2311120212003022-3220201122020312-0012201202331232-2100322133333131-3233031333021202"></a>

## Direct properties — no_global_network / 322200003102 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0212000121222200-1210303110331132-0030013010213033-3111013013202120-1302010313102032-3202103322302103-3202123202233330-2301230331023322"></a>

## Next pages — no_global_network / 322200003102 / 4

- [ingress_egress_gw](data-sources--gcp_vpc_site--reference--group-001.md#canonical-1123313330020103-2220202203231111-2113330331022333-0133213021020321-0303223331023010-1320133103131013-2322232313123200-2330220010011031)
- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-1132222133321101-1232102200201102-2111112102132122-0010023221133123-2333213300031302-1001011022311003-3121103220313010-2103321200303202)

<a id="canonical-3323101233232200-1121011001111111-2122331313313332-1103001110103122-3203022213321323-3303022232101131-3320313203232111-0001201212301111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2003120201230320-0323011011000120-1000020211113112-3012132000302001-1310312321133020-0233022001130330-3220013302133212-2232131233110312"></a>

## ingress_egress_gw.no_inside_static_routes — no_inside_static_routes / 231003232333 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-1132222133321101-1232102200201102-2111112102132122-0010023221133123-2333213300031302-1001011022311003-3121103220313010-2103321200303202)
- [Property reference](data-sources--gcp_vpc_site--reference--group-001.md#canonical-0023231003100101-1010210121112300-0303103132302303-0331000323201313-3003311210331022-3000312312212303-0030202320102022-1232013113103201)
- [ingress_egress_gw](data-sources--gcp_vpc_site--reference--group-001.md#canonical-1123313330020103-2220202203231111-2113330331022333-0133213021020321-0303223331023010-1320133103131013-2322232313123200-2330220010011031)
- ingress_egress_gw.no_inside_static_routes

<a id="canonical-3221332002031313-2010131221011230-1230231202011312-0202013120103102-1333022113120120-0013320312002211-3102000313211302-2000331100123023"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-0000112221120323-3301231212002233-1201321023002301-0311222333023232-1103201231132003-0032111102010123-1120233233030003-3201320033032200"></a>

## Direct properties — no_inside_static_routes / 231003232333 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0213100203211231-3121030013110331-0103021030303133-0211100331323202-0101002121022221-0132123003100121-1330221001020030-3011132221013123"></a>

## Next pages — no_inside_static_routes / 231003232333 / 4

- [ingress_egress_gw](data-sources--gcp_vpc_site--reference--group-001.md#canonical-1123313330020103-2220202203231111-2113330331022333-0133213021020321-0303223331023010-1320133103131013-2322232313123200-2330220010011031)
- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-1132222133321101-1232102200201102-2111112102132122-0010023221133123-2333213300031302-1001011022311003-3121103220313010-2103321200303202)

<a id="canonical-2321230320203023-2012002031123220-0211203302300320-0331010030233012-0113111022300021-1131233223230022-0203113021211023-0321203131330223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3320223220321213-0123010130210320-3121023113300313-2220330133200103-3202302011333210-2022101311021201-3110222033131302-1132212001022020"></a>

## ingress_egress_gw.no_network_policy — no_network_policy / 020323333323 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-1132222133321101-1232102200201102-2111112102132122-0010023221133123-2333213300031302-1001011022311003-3121103220313010-2103321200303202)
- [Property reference](data-sources--gcp_vpc_site--reference--group-001.md#canonical-0023231003100101-1010210121112300-0303103132302303-0331000323201313-3003311210331022-3000312312212303-0030202320102022-1232013113103201)
- [ingress_egress_gw](data-sources--gcp_vpc_site--reference--group-001.md#canonical-1123313330020103-2220202203231111-2113330331022333-0133213021020321-0303223331023010-1320133103131013-2322232313123200-2330220010011031)
- ingress_egress_gw.no_network_policy

<a id="canonical-0132122202101131-2020132123110110-2113003000123102-2223220313032212-0310133313030033-0100311313301110-0003133333001101-3133301121303302"></a>

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

<a id="canonical-3303013321220312-3023330232031202-1213002010321113-3200231103020320-2101001021132311-0233123120333332-1321310123120012-2130300203132313"></a>

## Direct properties — no_network_policy / 020323333323 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0022100030303112-1110031203103001-2122002030011112-0222233113020112-2202303030222232-3212203323002133-0103031122200130-0022101330112331"></a>

## Next pages — no_network_policy / 020323333323 / 4

- [ingress_egress_gw](data-sources--gcp_vpc_site--reference--group-001.md#canonical-1123313330020103-2220202203231111-2113330331022333-0133213021020321-0303223331023010-1320133103131013-2322232313123200-2330220010011031)
- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-1132222133321101-1232102200201102-2111112102132122-0010023221133123-2333213300031302-1001011022311003-3121103220313010-2103321200303202)

<a id="canonical-2121122223003320-1111002330210202-1113133312303121-3331201332000322-1210100303320031-3012330310320210-2111211301122021-3130310200112133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0021203013123111-2022103332221330-3231111010120102-2113203011313202-1002310330101002-3312021032122332-0032013131013133-0123220033012210"></a>

## ingress_egress_gw.no_outside_static_routes — no_outside_static_routes / 221122133330 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-1132222133321101-1232102200201102-2111112102132122-0010023221133123-2333213300031302-1001011022311003-3121103220313010-2103321200303202)
- [Property reference](data-sources--gcp_vpc_site--reference--group-001.md#canonical-0023231003100101-1010210121112300-0303103132302303-0331000323201313-3003311210331022-3000312312212303-0030202320102022-1232013113103201)
- [ingress_egress_gw](data-sources--gcp_vpc_site--reference--group-001.md#canonical-1123313330020103-2220202203231111-2113330331022333-0133213021020321-0303223331023010-1320133103131013-2322232313123200-2330220010011031)
- ingress_egress_gw.no_outside_static_routes

<a id="canonical-0333203110320330-1132120310113110-3332012300033200-1213033032112101-2211320123122222-1230301130333211-2230110310223013-1213101023032300"></a>

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

<a id="canonical-0013003213120223-2133013002322123-2300120211232023-3312100033233321-3220110210321020-0222130031023022-3033120330333322-3023111220210201"></a>

## Direct properties — no_outside_static_routes / 221122133330 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2130320133222123-2233221313212200-1002103030012200-2230000001121121-0030311321231013-3022001122131211-2000201123302120-3002311100112011"></a>

## Next pages — no_outside_static_routes / 221122133330 / 4

- [ingress_egress_gw](data-sources--gcp_vpc_site--reference--group-001.md#canonical-1123313330020103-2220202203231111-2113330331022333-0133213021020321-0303223331023010-1320133103131013-2322232313123200-2330220010011031)
- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-1132222133321101-1232102200201102-2111112102132122-0010023221133123-2333213300031302-1001011022311003-3121103220313010-2103321200303202)

<a id="canonical-0110320013312231-3122213130230201-1122220010231112-0112002300113320-1001032001133133-0203110110301223-0312233112010132-1311303220313332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2101202022213012-2320320223320310-1321331130232313-0321320211133113-0212210132102121-0011332312231102-3211103030111210-1211020013230010"></a>

## ingress_egress_gw.outside_network — outside_network / 030311333031 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-1132222133321101-1232102200201102-2111112102132122-0010023221133123-2333213300031302-1001011022311003-3121103220313010-2103321200303202)
- [Property reference](data-sources--gcp_vpc_site--reference--group-001.md#canonical-0023231003100101-1010210121112300-0303103132302303-0331000323201313-3003311210331022-3000312312212303-0030202320102022-1232013113103201)
- [ingress_egress_gw](data-sources--gcp_vpc_site--reference--group-001.md#canonical-1123313330020103-2220202203231111-2113330331022333-0133213021020321-0303223331023010-1320133103131013-2322232313123200-2330220010011031)
- ingress_egress_gw.outside_network

<a id="canonical-0313202301100112-1001031023230231-3333101222012102-3202133211302210-3303133320333332-0112122323102331-1102101233320033-0121013330311312"></a>

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

<a id="canonical-0013123111213000-0232112310333303-0132312221232133-3121222300213131-0230001310201123-2023320212022333-3030303212102101-3220203202130130"></a>

## Direct properties — outside_network / 030311333031 / 3

- [existing_network](data-sources--gcp_vpc_site--reference--group-002.md#canonical-2112333113123121-1101312320022003-2301330000311000-2022323031320103-2332003201301030-1000323100131021-1232021122103130-0220123202000112): complete subsection reference.

- [new_network](data-sources--gcp_vpc_site--reference--group-002.md#canonical-3003031120212300-3203113033003113-2131102210302233-0311203030003333-2033222222112012-3103031101213323-3132213300022302-1132302013200212): complete subsection reference.

- [new_network_autogenerate](data-sources--gcp_vpc_site--reference--group-002.md#canonical-0203232211201110-3210200212230212-1330221112013230-2200320230201133-0023311013023032-0110310220103202-3313300221121233-3032200202033103): complete subsection reference.

<a id="canonical-2212023033230001-0112210102011302-0133002320330313-3003221220230230-0333332232011013-2011230112102320-2110023312121233-0030121201230301"></a>

## Next pages — outside_network / 030311333031 / 4

- [ingress_egress_gw.outside_network.existing_network](data-sources--gcp_vpc_site--reference--group-002.md#canonical-2112333113123121-1101312320022003-2301330000311000-2022323031320103-2332003201301030-1000323100131021-1232021122103130-0220123202000112)
- [ingress_egress_gw.outside_network.new_network](data-sources--gcp_vpc_site--reference--group-002.md#canonical-3003031120212300-3203113033003113-2131102210302233-0311203030003333-2033222222112012-3103031101213323-3132213300022302-1132302013200212)
- [ingress_egress_gw.outside_network.new_network_autogenerate](data-sources--gcp_vpc_site--reference--group-002.md#canonical-0203232211201110-3210200212230212-1330221112013230-2200320230201133-0023311013023032-0110310220103202-3313300221121233-3032200202033103)
- [ingress_egress_gw](data-sources--gcp_vpc_site--reference--group-001.md#canonical-1123313330020103-2220202203231111-2113330331022333-0133213021020321-0303223331023010-1320133103131013-2322232313123200-2330220010011031)
- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-1132222133321101-1232102200201102-2111112102132122-0010023221133123-2333213300031302-1001011022311003-3121103220313010-2103321200303202)

<a id="canonical-2112333113123121-1101312320022003-2301330000311000-2022323031320103-2332003201301030-1000323100131021-1232021122103130-0220123202000112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3331132230031223-1302020131321200-0223201120213102-1010012001031231-1322220033010021-2211100212121230-2131112300311012-0022212302313220"></a>

## ingress_egress_gw.outside_network.existing_network — existing_network / 223210113021 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-1132222133321101-1232102200201102-2111112102132122-0010023221133123-2333213300031302-1001011022311003-3121103220313010-2103321200303202)
- [Property reference](data-sources--gcp_vpc_site--reference--group-001.md#canonical-0023231003100101-1010210121112300-0303103132302303-0331000323201313-3003311210331022-3000312312212303-0030202320102022-1232013113103201)
- [ingress_egress_gw](data-sources--gcp_vpc_site--reference--group-001.md#canonical-1123313330020103-2220202203231111-2113330331022333-0133213021020321-0303223331023010-1320133103131013-2322232313123200-2330220010011031)
- [ingress_egress_gw.outside_network](data-sources--gcp_vpc_site--reference--group-002.md#canonical-0110320013312231-3122213130230201-1122220010231112-0112002300113320-1001032001133133-0203110110301223-0312233112010132-1311303220313332)
- ingress_egress_gw.outside_network.existing_network

<a id="canonical-2110022110221201-2320022330033321-3110300110031332-0120013110220221-1200332320202100-0101031321000120-1211011132222203-0310120022223202"></a>

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

<a id="canonical-1200121120301033-1000132232200213-3200010313000221-0113211230033122-2213102021023230-0202001021032303-3101101112211310-2011300321002233"></a>

## Direct properties — existing_network / 223210113021 / 3

<a id="canonical-1123202222333000-0021203322311032-3102302020102323-1030111000220003-3010302231213311-0032221010120321-2230131332010130-0201122221120112"></a>

<a id="canonical-3031230302213303-2033120200122120-3311112333203220-0130010232312231-0222113102121222-1211222113321001-1120330322022200-1200212103033113"></a>

## name property — existing_network / 223210113021 / 4

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

<a id="canonical-1231020132003331-2213310222012102-2213121211332233-3311321130033300-0033113102212100-3123321001002133-2221100213123333-0313112102003131"></a>

## Next pages — existing_network / 223210113021 / 5

- [ingress_egress_gw.outside_network](data-sources--gcp_vpc_site--reference--group-002.md#canonical-0110320013312231-3122213130230201-1122220010231112-0112002300113320-1001032001133133-0203110110301223-0312233112010132-1311303220313332)
- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-1132222133321101-1232102200201102-2111112102132122-0010023221133123-2333213300031302-1001011022311003-3121103220313010-2103321200303202)

<a id="canonical-3003031120212300-3203113033003113-2131102210302233-0311203030003333-2033222222112012-3103031101213323-3132213300022302-1132302013200212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0211300112230111-1223130333330200-1302233312023310-0210230231022331-0001033133233203-1333320012212110-3002210011101213-3223310332232013"></a>

## ingress_egress_gw.outside_network.new_network — new_network / 023031202313 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-1132222133321101-1232102200201102-2111112102132122-0010023221133123-2333213300031302-1001011022311003-3121103220313010-2103321200303202)
- [Property reference](data-sources--gcp_vpc_site--reference--group-001.md#canonical-0023231003100101-1010210121112300-0303103132302303-0331000323201313-3003311210331022-3000312312212303-0030202320102022-1232013113103201)
- [ingress_egress_gw](data-sources--gcp_vpc_site--reference--group-001.md#canonical-1123313330020103-2220202203231111-2113330331022333-0133213021020321-0303223331023010-1320133103131013-2322232313123200-2330220010011031)
- [ingress_egress_gw.outside_network](data-sources--gcp_vpc_site--reference--group-002.md#canonical-0110320013312231-3122213130230201-1122220010231112-0112002300113320-1001032001133133-0203110110301223-0312233112010132-1311303220313332)
- ingress_egress_gw.outside_network.new_network

<a id="canonical-3110302032210223-3133330210133031-2202112321030301-2001010202333221-3210002033222310-3331302330303001-1103003213010012-3312223120333112"></a>

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

<a id="canonical-0113200232132023-1201021201003122-1120121011212101-0121220210113210-0133121322032020-0201112331120113-1310110210031312-0002213120010301"></a>

## Direct properties — new_network / 023031202313 / 3

<a id="canonical-1311211120021301-3013220230111123-0203232332213210-0012311001130001-3013032100332023-2030020310133131-2223011321203113-2332223222031120"></a>

<a id="canonical-2123213330303122-2023023202223020-3112121000033101-0303000032023003-2320103312003030-3132122200221130-1323111023202022-0030202133310311"></a>

## name property — new_network / 023031202313 / 4

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

<a id="canonical-1000113022130000-0323320310103033-1133001102232311-2233221302023103-3133302121331131-2333303303032201-3321223111321300-1020122123331203"></a>

## Next pages — new_network / 023031202313 / 5

- [ingress_egress_gw.outside_network](data-sources--gcp_vpc_site--reference--group-002.md#canonical-0110320013312231-3122213130230201-1122220010231112-0112002300113320-1001032001133133-0203110110301223-0312233112010132-1311303220313332)
- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-1132222133321101-1232102200201102-2111112102132122-0010023221133123-2333213300031302-1001011022311003-3121103220313010-2103321200303202)

<a id="canonical-0203232211201110-3210200212230212-1330221112013230-2200320230201133-0023311013023032-0110310220103202-3313300221121233-3032200202033103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3223301020110201-2210221021013123-0013302120223313-2121030103022213-2313320221231032-2001220120121013-2132300232323033-1120122010113122"></a>

## ingress_egress_gw.outside_network.new_network_autogenerate — new_network_autogenerate / 102302100010 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-1132222133321101-1232102200201102-2111112102132122-0010023221133123-2333213300031302-1001011022311003-3121103220313010-2103321200303202)
- [Property reference](data-sources--gcp_vpc_site--reference--group-001.md#canonical-0023231003100101-1010210121112300-0303103132302303-0331000323201313-3003311210331022-3000312312212303-0030202320102022-1232013113103201)
- [ingress_egress_gw](data-sources--gcp_vpc_site--reference--group-001.md#canonical-1123313330020103-2220202203231111-2113330331022333-0133213021020321-0303223331023010-1320133103131013-2322232313123200-2330220010011031)
- [ingress_egress_gw.outside_network](data-sources--gcp_vpc_site--reference--group-002.md#canonical-0110320013312231-3122213130230201-1122220010231112-0112002300113320-1001032001133133-0203110110301223-0312233112010132-1311303220313332)
- ingress_egress_gw.outside_network.new_network_autogenerate

<a id="canonical-0031111213113013-0103010030232011-0003003202321103-3132010030011132-0231012221320112-0331300220032221-2120210112102030-0333030113032331"></a>

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

<a id="canonical-3022211210203302-3012031202302212-3213023331033221-2000112211030102-3113131122233223-3213210111113033-1103311230121311-2333203113102320"></a>

## Direct properties — new_network_autogenerate / 102302100010 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0223031312111000-0233121012122103-1301312013002310-0120310220020232-0031233100310102-0113030331202233-2313030122201221-2133031022220121"></a>

## Next pages — new_network_autogenerate / 102302100010 / 4

- [ingress_egress_gw.outside_network](data-sources--gcp_vpc_site--reference--group-002.md#canonical-0110320013312231-3122213130230201-1122220010231112-0112002300113320-1001032001133133-0203110110301223-0312233112010132-1311303220313332)
- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-1132222133321101-1232102200201102-2111112102132122-0010023221133123-2333213300031302-1001011022311003-3121103220313010-2103321200303202)

<a id="canonical-2222212203011002-0122100002033220-1003302221310103-3022130333223033-1321222030130012-0211202033021113-0330312012113010-2122030230011200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2210232202110212-2221302021232130-0032100332113032-3001011123322231-1023113302330123-1322320323100123-0133111321211123-0031110010102212"></a>

## ingress_egress_gw.outside_static_routes — outside_static_routes / 023113013011 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-1132222133321101-1232102200201102-2111112102132122-0010023221133123-2333213300031302-1001011022311003-3121103220313010-2103321200303202)
- [Property reference](data-sources--gcp_vpc_site--reference--group-001.md#canonical-0023231003100101-1010210121112300-0303103132302303-0331000323201313-3003311210331022-3000312312212303-0030202320102022-1232013113103201)
- [ingress_egress_gw](data-sources--gcp_vpc_site--reference--group-001.md#canonical-1123313330020103-2220202203231111-2113330331022333-0133213021020321-0303223331023010-1320133103131013-2322232313123200-2330220010011031)
- ingress_egress_gw.outside_static_routes

<a id="canonical-0323023232303223-2121211210022131-3300033022311100-1032131332303230-3001120032220202-1321312121020112-1202102232031111-0312000212011101"></a>

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

<a id="canonical-2123200023021032-2311322102120020-3013332212112212-3011031132213012-1022303311133202-3003013231032133-3330221023130011-3000031003333330"></a>

## Direct properties — outside_static_routes / 023113013011 / 3

- [static_route_list](data-sources--gcp_vpc_site--reference--group-002.md#canonical-1303330302312132-1121332321030032-3232300201320322-2013232231301011-1000303123022011-0120133312103013-3212023032032211-3012112131110121): complete subsection reference.

<a id="canonical-1122113212100011-3100103120111211-2113001101112222-1323123130330313-3123301300320300-0103200311203333-0200022030210210-1213130002303212"></a>

## Next pages — outside_static_routes / 023113013011 / 4

- [ingress_egress_gw.outside_static_routes.static_route_list](data-sources--gcp_vpc_site--reference--group-002.md#canonical-1303330302312132-1121332321030032-3232300201320322-2013232231301011-1000303123022011-0120133312103013-3212023032032211-3012112131110121)
- [ingress_egress_gw](data-sources--gcp_vpc_site--reference--group-001.md#canonical-1123313330020103-2220202203231111-2113330331022333-0133213021020321-0303223331023010-1320133103131013-2322232313123200-2330220010011031)
- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-1132222133321101-1232102200201102-2111112102132122-0010023221133123-2333213300031302-1001011022311003-3121103220313010-2103321200303202)

<a id="canonical-1303330302312132-1121332321030032-3232300201320322-2013232231301011-1000303123022011-0120133312103013-3212023032032211-3012112131110121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3131301020012130-0123310113103231-3023022223211110-0321331210323223-3113000031312121-0303022032323023-0031101111331220-2223300300211323"></a>

## ingress_egress_gw.outside_static_routes.static_route_list — static_route_list / 220320022300 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-1132222133321101-1232102200201102-2111112102132122-0010023221133123-2333213300031302-1001011022311003-3121103220313010-2103321200303202)
- [Property reference](data-sources--gcp_vpc_site--reference--group-001.md#canonical-0023231003100101-1010210121112300-0303103132302303-0331000323201313-3003311210331022-3000312312212303-0030202320102022-1232013113103201)
- [ingress_egress_gw](data-sources--gcp_vpc_site--reference--group-001.md#canonical-1123313330020103-2220202203231111-2113330331022333-0133213021020321-0303223331023010-1320133103131013-2322232313123200-2330220010011031)
- [ingress_egress_gw.outside_static_routes](data-sources--gcp_vpc_site--reference--group-002.md#canonical-2222212203011002-0122100002033220-1003302221310103-3022130333223033-1321222030130012-0211202033021113-0330312012113010-2122030230011200)
- ingress_egress_gw.outside_static_routes.static_route_list

<a id="canonical-1032020001221203-3211103022330321-1111130120110231-2211301200131021-1000121011213311-3222012231020003-3103131101000310-1322333223202230"></a>

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-3230121122211122-1223103010123130-1212220020213022-1101200323030112-3023033012210101-3010113012310012-3002221111323003-0112123131203302"></a>

## Direct properties — static_route_list / 220320022300 / 3

- [custom_static_route](data-sources--gcp_vpc_site--reference--group-002.md#canonical-3330202303321021-0200002222221111-3320121133011222-2300222202210231-0132003120323301-1111330302033032-0033131231022133-1211232101033013): complete subsection reference.

<a id="canonical-2332030301211220-1332310011310010-0320032223000203-2121312203320303-1013300120022111-0102302210223320-1102200301222021-2100311302300133"></a>

<a id="canonical-2203223010230110-2033333121133130-3232132020012302-3002120322023001-0020031301310221-3123030101210131-2323231233020132-1020322233201113"></a>

## simple_static_route property — static_route_list / 220320022300 / 4

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
    "ves.io.schema.rules.string.ipv4_prefix": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv4_prefix": "true"
  }
}
```

<a id="canonical-1000210013110010-2110102330131023-1231132121033120-0221322110100222-3220032133030220-2231120332202212-3223302113330031-3013232300230200"></a>

## Next pages — static_route_list / 220320022300 / 5

- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route](data-sources--gcp_vpc_site--reference--group-002.md#canonical-3330202303321021-0200002222221111-3320121133011222-2300222202210231-0132003120323301-1111330302033032-0033131231022133-1211232101033013)
- [ingress_egress_gw.outside_static_routes](data-sources--gcp_vpc_site--reference--group-002.md#canonical-2222212203011002-0122100002033220-1003302221310103-3022130333223033-1321222030130012-0211202033021113-0330312012113010-2122030230011200)
- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-1132222133321101-1232102200201102-2111112102132122-0010023221133123-2333213300031302-1001011022311003-3121103220313010-2103321200303202)

<a id="canonical-3330202303321021-0200002222221111-3320121133011222-2300222202210231-0132003120323301-1111330302033032-0033131231022133-1211232101033013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3333101103130300-0211322132101130-2012002030012120-2222323102212232-1123232321122222-0113002003211132-1022201321202231-1023122113131121"></a>

## ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route — custom_static_route / 213103212102 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md#canonical-1132222133321101-1232102200201102-2111112102132122-0010023221133123-2333213300031302-1001011022311003-3121103220313010-2103321200303202)
- [Property reference](data-sources--gcp_vpc_site--reference--group-001.md#canonical-0023231003100101-1010210121112300-0303103132302303-0331000323201313-3003311210331022-3000312312212303-0030202320102022-1232013113103201)
- [ingress_egress_gw](data-sources--gcp_vpc_site--reference--group-001.md#canonical-1123313330020103-2220202203231111-2113330331022333-0133213021020321-0303223331023010-1320133103131013-2322232313123200-2330220010011031)
- [ingress_egress_gw.outside_static_routes](data-sources--gcp_vpc_site--reference--group-002.md#canonical-2222212203011002-0122100002033220-1003302221310103-3022130333223033-1321222030130012-0211202033021113-0330312012113010-2122030230011200)
- [ingress_egress_gw.outside_static_routes.static_route_list](data-sources--gcp_vpc_site--reference--group-002.md#canonical-1303330302312132-1121332321030032-3232300201320322-2013232231301011-1000303123022011-0120133312103013-3212023032032211-3012112131110121)
- ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route

<a id="canonical-2103200313332222-1131023303323121-2222313212320320-2303113331130031-1300031032303223-2112212011131213-3031301123323102-2200101331133010"></a>

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

<a id="canonical-1012222100322102-1111121032312303-0223202133100131-1013323113201330-2001300201312122-1133321032320130-1113133201122022-0013220131001322"></a>

## Direct properties — custom_static_route / 213103212102 / 3

<a id="canonical-1113021133222132-0233222330121200-0222023222131011-0303101213230232-2112320320012000-0032320310322333-2202322002203032-2000002032301032"></a>

<a id="canonical-3212000300021020-3112220310021131-1010022123012212-1213210212320022-3023310020222311-2030131033110211-1303131330110011-1100330011323032"></a>

## attrs property — custom_static_route / 213103212102 / 4

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
    "ves.io.schema.rules.repeated.max_items": "4"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "4"
  }
}
```

- [labels](data-sources--gcp_vpc_site--reference--group-003.md#canonical-0322223121033130-3201331110102202-3323013213111212-1223222300100021-0130300132011311-3322320210310233-3331103011001222-2210023213113300): complete subsection reference.

- [nexthop](data-sources--gcp_vpc_site--reference--group-003.md#canonical-0321020222123230-1311211201211220-3333320103332321-3332210131100313-2002203032333220-1002123221331321-1130203022113333-2323212121233330): complete subsection reference.

- [subnets](data-sources--gcp_vpc_site--reference--group-003.md#canonical-2223102112003230-2120223010131011-3311220202110311-0321102303122130-2120231212223300-0313302303012100-0233120210103013-1013201130003200): complete subsection reference.
