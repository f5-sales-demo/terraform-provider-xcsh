---
page_title: "xcsh_azure_vnet_site reference"
subcategory: "Infrastructure"
description: "Complete grouped canonical reference for xcsh_azure_vnet_site reference."
---

# xcsh_azure_vnet_site reference

<a id="canonical-3133221213233032-2012301113202223-0102012121222323-1131121110223313-0203120101131111-3003132030001011-0303103103320320-3132303101201031"></a>

## namespace property — k8s_cluster / 212303131311 / 5

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

<a id="canonical-3331203110012321-2222122320120023-1221231331132310-0223330331220131-1103012210322302-1201001201201311-1202030333303211-2331333311030133"></a>

<a id="canonical-0330200111123323-2330310003202201-0201302003020131-1100102031310010-3011300220101332-0333330201002331-2233313332122012-3302310333321102"></a>

## tenant property — k8s_cluster / 212303131311 / 6

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

<a id="canonical-1000220001300002-0133122213120123-0232030113333123-2321103011102212-1221003323233331-2213131020031133-0213020323231022-3320331013200031"></a>

## Next pages — k8s_cluster / 212303131311 / 7

- [voltstack_cluster](data-sources--azure_vnet_site--reference--group-008.md#canonical-0023321213310231-1323112323303023-2220001003301330-3020021300200322-1322303031300210-2010300303132331-3120231121123332-0101313132130000)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)

<a id="canonical-2210301311301020-2323110310321011-1022330333133133-2132303010330033-3311030120302000-1322312221030330-1022201203120312-0312120302231103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1023332321101002-3221220231001102-3332303213320331-1110110331113223-1213311312310221-0130313313332022-2100133103201332-3223102010203102"></a>

## voltstack_cluster.no_dc_cluster_group — no_dc_cluster_group / 230303212122 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-2210100210332112-1233221121133301-0120203222303210-2221022223110220-0313230312203322-2232200300001103-2301220003133113-1333232230300113)
- [voltstack_cluster](data-sources--azure_vnet_site--reference--group-008.md#canonical-0023321213310231-1323112323303023-2220001003301330-3020021300200322-1322303031300210-2010300303132331-3120231121123332-0101313132130000)
- voltstack_cluster.no_dc_cluster_group

<a id="canonical-1313223010311010-3231111203211101-0021321302310020-2000230203212212-1130321313200213-2122103223310022-0303130110131023-3300323020132323"></a>

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

<a id="canonical-1203220311311120-0011233320013132-0022322011010031-1220030300212312-2301232123011023-1333301130211021-0321210111231101-3322012120121203"></a>

## Direct properties — no_dc_cluster_group / 230303212122 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1310102221312233-1112032333023133-3230303133323103-0120333220212233-3332320133203332-0030200021113101-1301130310123212-2201133013010210"></a>

## Next pages — no_dc_cluster_group / 230303212122 / 4

- [voltstack_cluster](data-sources--azure_vnet_site--reference--group-008.md#canonical-0023321213310231-1323112323303023-2220001003301330-3020021300200322-1322303031300210-2010300303132331-3120231121123332-0101313132130000)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)

<a id="canonical-1110113032033220-0120210123012323-2010102301210130-0002330101222001-2100022000011031-0232213310123131-2002100311211200-0011013333113223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2302203203102200-0133012021323132-3213232301133303-3232223120330111-2120223123230330-0220332222122121-3112312022332200-0003220230212232"></a>

## voltstack_cluster.no_forward_proxy — no_forward_proxy / 021020221333 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-2210100210332112-1233221121133301-0120203222303210-2221022223110220-0313230312203322-2232200300001103-2301220003133113-1333232230300113)
- [voltstack_cluster](data-sources--azure_vnet_site--reference--group-008.md#canonical-0023321213310231-1323112323303023-2220001003301330-3020021300200322-1322303031300210-2010300303132331-3120231121123332-0101313132130000)
- voltstack_cluster.no_forward_proxy

<a id="canonical-1102032013103323-2023233031033331-0121330000303013-1320000311120101-3020111102130022-2220313200203123-2101303213131001-2331332000122222"></a>

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

<a id="canonical-0331312121010210-2130231002023220-2220022020213010-2230201211031113-0302110111210120-2112312001232210-2130110121103303-3300220020321033"></a>

## Direct properties — no_forward_proxy / 021020221333 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3133032321201103-2222303323233033-1312132000120022-3130202121311131-0110222223203031-3130303301020022-3012322220012123-3033310121322121"></a>

## Next pages — no_forward_proxy / 021020221333 / 4

- [voltstack_cluster](data-sources--azure_vnet_site--reference--group-008.md#canonical-0023321213310231-1323112323303023-2220001003301330-3020021300200322-1322303031300210-2010300303132331-3120231121123332-0101313132130000)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)

<a id="canonical-1220030220233121-1022201322121111-1320331102132003-2301031222030321-1310201203332031-2133302320010303-1030000210102322-1130231312313320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0030320202222013-2130310201300301-2223033213211122-2032302013212012-2320210102021322-0330200002032321-3221233222113031-1133303120120310"></a>

## voltstack_cluster.no_global_network — no_global_network / 121300223200 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-2210100210332112-1233221121133301-0120203222303210-2221022223110220-0313230312203322-2232200300001103-2301220003133113-1333232230300113)
- [voltstack_cluster](data-sources--azure_vnet_site--reference--group-008.md#canonical-0023321213310231-1323112323303023-2220001003301330-3020021300200322-1322303031300210-2010300303132331-3120231121123332-0101313132130000)
- voltstack_cluster.no_global_network

<a id="canonical-2021232011110223-2321332130332100-3133213030100103-3302101000102111-0300101311000213-2010112232021022-0103102202312011-0302001132033003"></a>

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

<a id="canonical-2121212321133203-1123311203023033-2001313112103012-1130222312210133-0112122210333203-1021021313032030-3013213322223221-2013201202323133"></a>

## Direct properties — no_global_network / 121300223200 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0011200300210011-0310222220210022-0313100300022031-3021230322301212-2232022201002301-1011320330012222-1110323313203123-3122202313333013"></a>

## Next pages — no_global_network / 121300223200 / 4

- [voltstack_cluster](data-sources--azure_vnet_site--reference--group-008.md#canonical-0023321213310231-1323112323303023-2220001003301330-3020021300200322-1322303031300210-2010300303132331-3120231121123332-0101313132130000)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)

<a id="canonical-3023131200121120-0210031311220220-2313211113313123-0123010022333012-2000202121130320-0310021030203121-3111301302311202-0121311111332300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3300233301330030-1102222111111001-2301233311330323-2030033203032300-2322312022330221-1022110022112301-3333102310131130-0132331133110330"></a>

## voltstack_cluster.no_k8s_cluster — no_k8s_cluster / 303121123111 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-2210100210332112-1233221121133301-0120203222303210-2221022223110220-0313230312203322-2232200300001103-2301220003133113-1333232230300113)
- [voltstack_cluster](data-sources--azure_vnet_site--reference--group-008.md#canonical-0023321213310231-1323112323303023-2220001003301330-3020021300200322-1322303031300210-2010300303132331-3120231121123332-0101313132130000)
- voltstack_cluster.no_k8s_cluster

<a id="canonical-1133211102200032-3003223330230330-2113023020223012-3332203210130132-1012313321121021-3130323223220321-2033121111132220-0313133121133321"></a>

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

<a id="canonical-2231231021330301-1003303230332232-0102231121130211-3130320100221032-2130020212232222-1310212100322331-2311202112003200-3101231102300033"></a>

## Direct properties — no_k8s_cluster / 303121123111 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3030330330301003-3302011221013210-2020020113221221-2111031000330111-0221110032231231-2101200222321300-0202011221032333-0100013303022032"></a>

## Next pages — no_k8s_cluster / 303121123111 / 4

- [voltstack_cluster](data-sources--azure_vnet_site--reference--group-008.md#canonical-0023321213310231-1323112323303023-2220001003301330-3020021300200322-1322303031300210-2010300303132331-3120231121123332-0101313132130000)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)

<a id="canonical-2320012013200103-3201003213010030-3011220332330223-2301230333300203-3311323330200302-1003230133130210-3300023131211230-0120133020130112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2201021302110210-0000011120123012-2232211213130000-0123302201212301-3222121103113031-2312012122321213-1330122202211313-3131323322312310"></a>

## voltstack_cluster.no_network_policy — no_network_policy / 300321322202 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-2210100210332112-1233221121133301-0120203222303210-2221022223110220-0313230312203322-2232200300001103-2301220003133113-1333232230300113)
- [voltstack_cluster](data-sources--azure_vnet_site--reference--group-008.md#canonical-0023321213310231-1323112323303023-2220001003301330-3020021300200322-1322303031300210-2010300303132331-3120231121123332-0101313132130000)
- voltstack_cluster.no_network_policy

<a id="canonical-1011232330003120-1222111301022111-0221333211300132-1002010130011221-3222233232021131-2303122332033203-2032131130220003-1113221031011310"></a>

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

<a id="canonical-0212111003112211-1301310222220031-3312301131033210-0311202320112100-1201001133303232-2221211202133011-3110012022002021-1202130331022212"></a>

## Direct properties — no_network_policy / 300321322202 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2321012221333003-0102002000310210-1100320200010132-0201032002221131-2213001101202312-2322012113303003-1221231022102121-1323230113300330"></a>

## Next pages — no_network_policy / 300321322202 / 4

- [voltstack_cluster](data-sources--azure_vnet_site--reference--group-008.md#canonical-0023321213310231-1323112323303023-2220001003301330-3020021300200322-1322303031300210-2010300303132331-3120231121123332-0101313132130000)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)

<a id="canonical-2023320332323311-3020122322021012-1212311120123321-0231331111231212-1121212033301111-0123331233031321-3222033003310101-0123332220333123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1023302233302000-0203313113201130-0213001001212333-0030232033102220-2323303130022103-0010102031013213-3202111320132303-0310212000110301"></a>

## voltstack_cluster.no_outside_static_routes — no_outside_static_routes / 312030123031 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-2210100210332112-1233221121133301-0120203222303210-2221022223110220-0313230312203322-2232200300001103-2301220003133113-1333232230300113)
- [voltstack_cluster](data-sources--azure_vnet_site--reference--group-008.md#canonical-0023321213310231-1323112323303023-2220001003301330-3020021300200322-1322303031300210-2010300303132331-3120231121123332-0101313132130000)
- voltstack_cluster.no_outside_static_routes

<a id="canonical-3310022131023000-1103231001221220-3031213303003033-1231033221230312-0310222313102132-0113111021011122-0302322320121321-0030210012120133"></a>

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

<a id="canonical-3100213120000212-1320233121301003-0303333021011330-2112113312132223-2312033312100100-3221213200101223-1021132112032333-3301131112320303"></a>

## Direct properties — no_outside_static_routes / 312030123031 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1231100210313023-0333120301121301-0101020000323331-3302002000101300-2101112133033100-2030033221120312-2121113331112000-0120102102030223"></a>

## Next pages — no_outside_static_routes / 312030123031 / 4

- [voltstack_cluster](data-sources--azure_vnet_site--reference--group-008.md#canonical-0023321213310231-1323112323303023-2220001003301330-3020021300200322-1322303031300210-2010300303132331-3120231121123332-0101313132130000)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)

<a id="canonical-2122010301323013-0003013121031331-1113323123213010-2220330111003201-1220212213113310-0010110023331303-2022322233101133-1321233223121330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3220232110123020-1223223000132000-2131020200302020-3031023313322310-2031001110231032-1310023011002301-1030303330012013-3230111312122003"></a>

## voltstack_cluster.outside_static_routes — outside_static_routes / 320113023323 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-2210100210332112-1233221121133301-0120203222303210-2221022223110220-0313230312203322-2232200300001103-2301220003133113-1333232230300113)
- [voltstack_cluster](data-sources--azure_vnet_site--reference--group-008.md#canonical-0023321213310231-1323112323303023-2220001003301330-3020021300200322-1322303031300210-2010300303132331-3120231121123332-0101313132130000)
- voltstack_cluster.outside_static_routes

<a id="canonical-2020122222130112-3013112033100331-3012031201303310-0232011022011211-0301010230113223-2030132211113200-0122301020131323-0012010220120200"></a>

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

<a id="canonical-1030121033003023-3031231200231133-0301211003210103-2311221221301002-1002030300133101-1002111011112331-1321100123123132-1003130030300120"></a>

## Direct properties — outside_static_routes / 320113023323 / 3

- [static_route_list](data-sources--azure_vnet_site--reference--group-009.md#canonical-3003111000100013-3323000112311231-0023003330110223-2100010211100300-0031031230032100-0312001210330231-3122230101110102-3002010210221130): complete subsection reference.

<a id="canonical-0001122332302123-0022103133001101-2310032120201332-3202330311211311-2030020132210332-1000131221223313-3212031003120213-3031331112132302"></a>

## Next pages — outside_static_routes / 320113023323 / 4

- [voltstack_cluster.outside_static_routes.static_route_list](data-sources--azure_vnet_site--reference--group-009.md#canonical-3003111000100013-3323000112311231-0023003330110223-2100010211100300-0031031230032100-0312001210330231-3122230101110102-3002010210221130)
- [voltstack_cluster](data-sources--azure_vnet_site--reference--group-008.md#canonical-0023321213310231-1323112323303023-2220001003301330-3020021300200322-1322303031300210-2010300303132331-3120231121123332-0101313132130000)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)

<a id="canonical-3003111000100013-3323000112311231-0023003330110223-2100010211100300-0031031230032100-0312001210330231-3122230101110102-3002010210221130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1010223213221320-3202113011131212-1201011211032022-2312121201133200-0320033001013221-1212231011031220-3031123301131210-2030112120231300"></a>

## voltstack_cluster.outside_static_routes.static_route_list — static_route_list / 131112000233 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-2210100210332112-1233221121133301-0120203222303210-2221022223110220-0313230312203322-2232200300001103-2301220003133113-1333232230300113)
- [voltstack_cluster](data-sources--azure_vnet_site--reference--group-008.md#canonical-0023321213310231-1323112323303023-2220001003301330-3020021300200322-1322303031300210-2010300303132331-3120231121123332-0101313132130000)
- [voltstack_cluster.outside_static_routes](data-sources--azure_vnet_site--reference--group-009.md#canonical-2122010301323013-0003013121031331-1113323123213010-2220330111003201-1220212213113310-0010110023331303-2022322233101133-1321233223121330)
- voltstack_cluster.outside_static_routes.static_route_list

<a id="canonical-1133003131110120-3032121112300001-0000212103230013-3033130303113133-0023132330212213-0301231133013023-1231232310013131-1022101121033133"></a>

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

<a id="canonical-1233232201002210-0023310010332221-3231322312113021-0101123301331301-0102123120223231-3030311020211312-2101101132330313-2331101121022031"></a>

## Direct properties — static_route_list / 131112000233 / 3

- [custom_static_route](data-sources--azure_vnet_site--reference--group-009.md#canonical-1333011212022301-2002223323000032-1012000310011211-2133211221002303-2303303323202232-2111003121220301-1330020031131032-2022300000032110): complete subsection reference.

<a id="canonical-3103122210111313-2330323103221030-0223032322013203-1111000011311322-2211113120033321-3120303200300200-2121033002020313-2212302301133101"></a>

<a id="canonical-2303001222303102-0323202032122100-2120232111321033-2301221003323221-3000100221123011-3022331020022133-3232321311330311-0001230002131022"></a>

## simple_static_route property — static_route_list / 131112000233 / 4

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

<a id="canonical-3313020010223202-3122033233203031-1031320202221112-3113100113200320-1321302231022302-1310310320313303-3300002313323120-1112221011112301"></a>

## Next pages — static_route_list / 131112000233 / 5

- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route](data-sources--azure_vnet_site--reference--group-009.md#canonical-1333011212022301-2002223323000032-1012000310011211-2133211221002303-2303303323202232-2111003121220301-1330020031131032-2022300000032110)
- [voltstack_cluster.outside_static_routes](data-sources--azure_vnet_site--reference--group-009.md#canonical-2122010301323013-0003013121031331-1113323123213010-2220330111003201-1220212213113310-0010110023331303-2022322233101133-1321233223121330)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)

<a id="canonical-1333011212022301-2002223323000032-1012000310011211-2133211221002303-2303303323202232-2111003121220301-1330020031131032-2022300000032110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1231020230103320-1012123022102300-1230300202322023-0300201021312223-3201122010100333-2003100102211121-1111030223202112-2230333301231000"></a>

## voltstack_cluster.outside_static_routes.static_route_list.custom_static_route — custom_static_route / 123200323113 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-2210100210332112-1233221121133301-0120203222303210-2221022223110220-0313230312203322-2232200300001103-2301220003133113-1333232230300113)
- [voltstack_cluster](data-sources--azure_vnet_site--reference--group-008.md#canonical-0023321213310231-1323112323303023-2220001003301330-3020021300200322-1322303031300210-2010300303132331-3120231121123332-0101313132130000)
- [voltstack_cluster.outside_static_routes](data-sources--azure_vnet_site--reference--group-009.md#canonical-2122010301323013-0003013121031331-1113323123213010-2220330111003201-1220212213113310-0010110023331303-2022322233101133-1321233223121330)
- [voltstack_cluster.outside_static_routes.static_route_list](data-sources--azure_vnet_site--reference--group-009.md#canonical-3003111000100013-3323000112311231-0023003330110223-2100010211100300-0031031230032100-0312001210330231-3122230101110102-3002010210221130)
- voltstack_cluster.outside_static_routes.static_route_list.custom_static_route

<a id="canonical-2223133112110021-1330100220023003-1231233101132313-0031301123230313-3020202203211120-3012010110201110-3102020303022320-2130003300131322"></a>

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

<a id="canonical-2321220000333133-3230310120230212-1022201233322310-0202111202010202-0103333002322101-0123110110010320-0313310001320031-0122322022211132"></a>

## Direct properties — custom_static_route / 123200323113 / 3

<a id="canonical-3001033322131303-3131210111011310-1130031010301021-3302121033021300-3123023232030231-0333321132020122-1033110131203103-3123220032023002"></a>

<a id="canonical-3030202102230300-3332122113030303-2103003200210222-1231121021301232-0032333030132333-1002202111301202-1203010313300211-0132202121210310"></a>

## attrs property — custom_static_route / 123200323113 / 4

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

- [labels](data-sources--azure_vnet_site--reference--group-009.md#canonical-1003222333212113-0123211233033013-1301021210113120-3321113222032312-2031112110101331-0210002120112013-2013103233111330-1210332323203031): complete subsection reference.

- [nexthop](data-sources--azure_vnet_site--reference--group-009.md#canonical-3311223101100123-0112100121200001-1122322212113233-3201331232031021-3303311022201202-0103303121332211-1330001200300203-1220221232033320): complete subsection reference.

- [subnets](data-sources--azure_vnet_site--reference--group-009.md#canonical-2113023020113202-1220110303201212-2300303332110112-2301000001202231-3323122003201022-0221020000030212-1223101333201302-2002021122212131): complete subsection reference.

<a id="canonical-1203100133333322-3033300202233312-0110011003100222-1113002323233102-1230102103112032-1322320030223133-0233211010223111-0311121330202212"></a>

## Next pages — custom_static_route / 123200323113 / 5

- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.labels](data-sources--azure_vnet_site--reference--group-009.md#canonical-1003222333212113-0123211233033013-1301021210113120-3321113222032312-2031112110101331-0210002120112013-2013103233111330-1210332323203031)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop](data-sources--azure_vnet_site--reference--group-009.md#canonical-3311223101100123-0112100121200001-1122322212113233-3201331232031021-3303311022201202-0103303121332211-1330001200300203-1220221232033320)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets](data-sources--azure_vnet_site--reference--group-009.md#canonical-2113023020113202-1220110303201212-2300303332110112-2301000001202231-3323122003201022-0221020000030212-1223101333201302-2002021122212131)
- [voltstack_cluster.outside_static_routes.static_route_list](data-sources--azure_vnet_site--reference--group-009.md#canonical-3003111000100013-3323000112311231-0023003330110223-2100010211100300-0031031230032100-0312001210330231-3122230101110102-3002010210221130)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)

<a id="canonical-1003222333212113-0123211233033013-1301021210113120-3321113222032312-2031112110101331-0210002120112013-2013103233111330-1210332323203031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1020001230203200-1221113221132020-0101131120310313-3112202023321003-2110210130210031-1332111233032200-3202331211100032-3013303022031030"></a>

## voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.labels — labels / 213001103310 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-2210100210332112-1233221121133301-0120203222303210-2221022223110220-0313230312203322-2232200300001103-2301220003133113-1333232230300113)
- [voltstack_cluster](data-sources--azure_vnet_site--reference--group-008.md#canonical-0023321213310231-1323112323303023-2220001003301330-3020021300200322-1322303031300210-2010300303132331-3120231121123332-0101313132130000)
- [voltstack_cluster.outside_static_routes](data-sources--azure_vnet_site--reference--group-009.md#canonical-2122010301323013-0003013121031331-1113323123213010-2220330111003201-1220212213113310-0010110023331303-2022322233101133-1321233223121330)
- [voltstack_cluster.outside_static_routes.static_route_list](data-sources--azure_vnet_site--reference--group-009.md#canonical-3003111000100013-3323000112311231-0023003330110223-2100010211100300-0031031230032100-0312001210330231-3122230101110102-3002010210221130)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route](data-sources--azure_vnet_site--reference--group-009.md#canonical-1333011212022301-2002223323000032-1012000310011211-2133211221002303-2303303323202232-2111003121220301-1330020031131032-2022300000032110)
- voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.labels

<a id="canonical-3021031320020103-0133313113203303-2313330103332121-2012312132110111-1113120001000301-1313322120303030-1100111123320221-3013010020030011"></a>

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

<a id="canonical-1212110313033331-1033111113230123-1321203111031130-3302121013201110-1320311223303230-3010310232301132-0022303230222010-1000213132030322"></a>

## Direct properties — labels / 213001103310 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0131213011100032-0033331022311223-2311022022110331-3220222130110210-1111232131330220-1310331212311222-0022102320320303-3100100101200031"></a>

## Next pages — labels / 213001103310 / 4

- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route](data-sources--azure_vnet_site--reference--group-009.md#canonical-1333011212022301-2002223323000032-1012000310011211-2133211221002303-2303303323202232-2111003121220301-1330020031131032-2022300000032110)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)

<a id="canonical-3311223101100123-0112100121200001-1122322212113233-3201331232031021-3303311022201202-0103303121332211-1330001200300203-1220221232033320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2022021212130331-3210231121022002-2103203221010220-2133011013122300-2303002230320111-0101322013123323-2201031321013333-3113133110000330"></a>

## voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop — nexthop / 300330130313 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-2210100210332112-1233221121133301-0120203222303210-2221022223110220-0313230312203322-2232200300001103-2301220003133113-1333232230300113)
- [voltstack_cluster](data-sources--azure_vnet_site--reference--group-008.md#canonical-0023321213310231-1323112323303023-2220001003301330-3020021300200322-1322303031300210-2010300303132331-3120231121123332-0101313132130000)
- [voltstack_cluster.outside_static_routes](data-sources--azure_vnet_site--reference--group-009.md#canonical-2122010301323013-0003013121031331-1113323123213010-2220330111003201-1220212213113310-0010110023331303-2022322233101133-1321233223121330)
- [voltstack_cluster.outside_static_routes.static_route_list](data-sources--azure_vnet_site--reference--group-009.md#canonical-3003111000100013-3323000112311231-0023003330110223-2100010211100300-0031031230032100-0312001210330231-3122230101110102-3002010210221130)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route](data-sources--azure_vnet_site--reference--group-009.md#canonical-1333011212022301-2002223323000032-1012000310011211-2133211221002303-2303303323202232-2111003121220301-1330020031131032-2022300000032110)
- voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop

<a id="canonical-3112220223020212-3332203222003220-0220202103322021-2202133133020111-0300201221232002-3210112233321130-2032021102023013-3023331102200113"></a>

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

<a id="canonical-3033301303220310-1120000132220003-1032003321310022-0311100301103032-0212230201221000-2031333330123111-1013320012022100-1213132222001301"></a>

## Direct properties — nexthop / 300330130313 / 3

- [interface](data-sources--azure_vnet_site--reference--group-009.md#canonical-3002012011002220-0002330011012332-3102232003130133-1023110300023020-1210322021233202-0332122113221232-0120212300231332-3123331103112331): complete subsection reference.

- [nexthop_address](data-sources--azure_vnet_site--reference--group-009.md#canonical-2322030123023233-2001131000132330-0132220122210312-3222311222202003-2202011333132010-1121312311001210-2111310032222233-0211213133000202): complete subsection reference.

<a id="canonical-3203103331331320-3022111312202023-0223012101301303-3303132011123131-0022220231022110-0322211331302103-3010201010331011-3131131330001032"></a>

<a id="canonical-2000130213213203-0310203233212201-0020303323212113-1212121232033221-1200000322122321-2301102210320033-2322220000232313-2033130220232220"></a>

## type property — nexthop / 300330130313 / 4

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

<a id="canonical-1312003101333301-1002032002031312-3002301132133221-0213222232333122-1012223101102101-3222010301030130-1321211210101313-1133123023021131"></a>

## Next pages — nexthop / 300330130313 / 5

- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.interface](data-sources--azure_vnet_site--reference--group-009.md#canonical-3002012011002220-0002330011012332-3102232003130133-1023110300023020-1210322021233202-0332122113221232-0120212300231332-3123331103112331)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](data-sources--azure_vnet_site--reference--group-009.md#canonical-2322030123023233-2001131000132330-0132220122210312-3222311222202003-2202011333132010-1121312311001210-2111310032222233-0211213133000202)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route](data-sources--azure_vnet_site--reference--group-009.md#canonical-1333011212022301-2002223323000032-1012000310011211-2133211221002303-2303303323202232-2111003121220301-1330020031131032-2022300000032110)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)

<a id="canonical-3002012011002220-0002330011012332-3102232003130133-1023110300023020-1210322021233202-0332122113221232-0120212300231332-3123331103112331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3230231112303102-0301120202332110-2101230013321302-3101111211331000-2022222123022021-3103122111331323-3033311311130221-3131323210113122"></a>

## voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.interface — interface / 230013131213 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-2210100210332112-1233221121133301-0120203222303210-2221022223110220-0313230312203322-2232200300001103-2301220003133113-1333232230300113)
- [voltstack_cluster](data-sources--azure_vnet_site--reference--group-008.md#canonical-0023321213310231-1323112323303023-2220001003301330-3020021300200322-1322303031300210-2010300303132331-3120231121123332-0101313132130000)
- [voltstack_cluster.outside_static_routes](data-sources--azure_vnet_site--reference--group-009.md#canonical-2122010301323013-0003013121031331-1113323123213010-2220330111003201-1220212213113310-0010110023331303-2022322233101133-1321233223121330)
- [voltstack_cluster.outside_static_routes.static_route_list](data-sources--azure_vnet_site--reference--group-009.md#canonical-3003111000100013-3323000112311231-0023003330110223-2100010211100300-0031031230032100-0312001210330231-3122230101110102-3002010210221130)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route](data-sources--azure_vnet_site--reference--group-009.md#canonical-1333011212022301-2002223323000032-1012000310011211-2133211221002303-2303303323202232-2111003121220301-1330020031131032-2022300000032110)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop](data-sources--azure_vnet_site--reference--group-009.md#canonical-3311223101100123-0112100121200001-1122322212113233-3201331232031021-3303311022201202-0103303121332211-1330001200300203-1220221232033320)
- voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.interface

<a id="canonical-1002021003231303-2120001102223132-0021220321310020-2000021020212031-1110213321200102-3031303121220113-3320230220201011-2032333101031031"></a>

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

<a id="canonical-1020200101110321-1201023310013101-0030200100313301-2031310120113000-1120230103012201-3013310022130323-0200023003111023-2111222221012103"></a>

## Direct properties — interface / 230013131213 / 3

<a id="canonical-0030220101122310-3130132000000101-2220031011122010-1103301321110103-0201303212211131-2002331213121031-3312023031231231-2010111212110000"></a>

<a id="canonical-1223313220102200-1031011003012132-0022312201033320-3131212123000313-1323033233111232-3303021222011331-1300230310311302-2112020021032331"></a>

## kind property — interface / 230013131213 / 4

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

<a id="canonical-2130020210101122-0003200322220131-2220203130323100-0010201101221320-2002101322012223-0312022123121011-0000201210203301-0003021020111333"></a>

<a id="canonical-3213312230120110-2203323103302212-0300203312211033-0221203302030001-1101211022232231-1201010310331331-2303212231110120-1110212102001300"></a>

## name property — interface / 230013131213 / 5

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

<a id="canonical-0022103313210202-2113110002130111-1121111221103003-2030202121001133-2210022311301002-0110000203312130-0120122003221313-3320303012000202"></a>

<a id="canonical-0302130223032233-1133323323030330-2220131123133131-1202130031011012-3311021003023122-3033021213201201-2022010211323200-1011102320202001"></a>

## namespace property — interface / 230013131213 / 6

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

<a id="canonical-1100011023232232-2131023103031322-0001023033000201-0322203123022122-3213122000033102-1010323220010213-1132130121312203-2032111321232200"></a>

<a id="canonical-0320103023021312-0313210122133133-2033220002030233-3002311120110333-3012222002032231-1003030033110022-0121023230323003-0002322200333111"></a>

## tenant property — interface / 230013131213 / 7

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

<a id="canonical-0333330212101323-2301130331002133-2110210231323330-3221233212210012-0113022011023302-0311221112221200-3000303202003021-0210112310112131"></a>

<a id="canonical-0113031003311111-0333120212133203-1131231132132232-1332122112310200-3321310112212333-0233223212302203-1300111212030021-1112132322011201"></a>

## uid property — interface / 230013131213 / 8

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

<a id="canonical-0322311000213032-2301321133103300-0231321223032002-1222313020121100-2313033310313122-0130312221111210-0032333120022332-3323210030112210"></a>

## Next pages — interface / 230013131213 / 9

- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop](data-sources--azure_vnet_site--reference--group-009.md#canonical-3311223101100123-0112100121200001-1122322212113233-3201331232031021-3303311022201202-0103303121332211-1330001200300203-1220221232033320)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)

<a id="canonical-2322030123023233-2001131000132330-0132220122210312-3222311222202003-2202011333132010-1121312311001210-2111310032222233-0211213133000202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2111130201132320-0122120230322303-3112302332100223-0320213220002012-0332231333033133-2130110021003221-0013102100200202-3323123113320132"></a>

## voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address — nexthop_address / 003321211322 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-2210100210332112-1233221121133301-0120203222303210-2221022223110220-0313230312203322-2232200300001103-2301220003133113-1333232230300113)
- [voltstack_cluster](data-sources--azure_vnet_site--reference--group-008.md#canonical-0023321213310231-1323112323303023-2220001003301330-3020021300200322-1322303031300210-2010300303132331-3120231121123332-0101313132130000)
- [voltstack_cluster.outside_static_routes](data-sources--azure_vnet_site--reference--group-009.md#canonical-2122010301323013-0003013121031331-1113323123213010-2220330111003201-1220212213113310-0010110023331303-2022322233101133-1321233223121330)
- [voltstack_cluster.outside_static_routes.static_route_list](data-sources--azure_vnet_site--reference--group-009.md#canonical-3003111000100013-3323000112311231-0023003330110223-2100010211100300-0031031230032100-0312001210330231-3122230101110102-3002010210221130)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route](data-sources--azure_vnet_site--reference--group-009.md#canonical-1333011212022301-2002223323000032-1012000310011211-2133211221002303-2303303323202232-2111003121220301-1330020031131032-2022300000032110)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop](data-sources--azure_vnet_site--reference--group-009.md#canonical-3311223101100123-0112100121200001-1122322212113233-3201331232031021-3303311022201202-0103303121332211-1330001200300203-1220221232033320)
- voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address

<a id="canonical-1030122123002230-0020301300330231-2000233221222311-1033100330110102-0003321132300020-1311112201223003-0232023121022130-0302030323110331"></a>

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

<a id="canonical-2030012301220131-1300030112022201-3332331312111230-1310113002212302-3230032300232313-0223323321300310-0303220321323232-1101010113221001"></a>

## Direct properties — nexthop_address / 003321211322 / 3

- [dual_stack](data-sources--azure_vnet_site--reference--group-009.md#canonical-1313331112213121-0303132101303233-0200103221220203-0203122323131123-0312211110230112-2202320110102122-0332122003232333-3121311210332310): complete subsection reference.

- [ipv4](data-sources--azure_vnet_site--reference--group-009.md#canonical-3123111203020033-0230300100113302-3112111300123013-1202200301202300-0020201033311310-2032230103122303-0033313320320110-2332003303313213): complete subsection reference.

- [ipv6](data-sources--azure_vnet_site--reference--group-009.md#canonical-2000201333231230-0223330132223032-2122032221311010-3200211303121021-0133121212303321-0332303002200200-1103330101122002-1120121212312000): complete subsection reference.

<a id="canonical-3032323313132222-0020323233131212-2202013033133013-1310201231120210-2301232131222303-1322132100310212-1002202211030222-2003231033012112"></a>

## Next pages — nexthop_address / 003321211322 / 4

- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](data-sources--azure_vnet_site--reference--group-009.md#canonical-1313331112213121-0303132101303233-0200103221220203-0203122323131123-0312211110230112-2202320110102122-0332122003232333-3121311210332310)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4](data-sources--azure_vnet_site--reference--group-009.md#canonical-3123111203020033-0230300100113302-3112111300123013-1202200301202300-0020201033311310-2032230103122303-0033313320320110-2332003303313213)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6](data-sources--azure_vnet_site--reference--group-009.md#canonical-2000201333231230-0223330132223032-2122032221311010-3200211303121021-0133121212303321-0332303002200200-1103330101122002-1120121212312000)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop](data-sources--azure_vnet_site--reference--group-009.md#canonical-3311223101100123-0112100121200001-1122322212113233-3201331232031021-3303311022201202-0103303121332211-1330001200300203-1220221232033320)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)

<a id="canonical-1313331112213121-0303132101303233-0200103221220203-0203122323131123-0312211110230112-2202320110102122-0332122003232333-3121311210332310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2012101031221013-3331212012333023-1232121300112120-1023210032330101-0200021130320100-0322023231123212-0112000010031023-1220101230300101"></a>

## voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack — dual_stack / 232223002200 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-2210100210332112-1233221121133301-0120203222303210-2221022223110220-0313230312203322-2232200300001103-2301220003133113-1333232230300113)
- [voltstack_cluster](data-sources--azure_vnet_site--reference--group-008.md#canonical-0023321213310231-1323112323303023-2220001003301330-3020021300200322-1322303031300210-2010300303132331-3120231121123332-0101313132130000)
- [voltstack_cluster.outside_static_routes](data-sources--azure_vnet_site--reference--group-009.md#canonical-2122010301323013-0003013121031331-1113323123213010-2220330111003201-1220212213113310-0010110023331303-2022322233101133-1321233223121330)
- [voltstack_cluster.outside_static_routes.static_route_list](data-sources--azure_vnet_site--reference--group-009.md#canonical-3003111000100013-3323000112311231-0023003330110223-2100010211100300-0031031230032100-0312001210330231-3122230101110102-3002010210221130)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route](data-sources--azure_vnet_site--reference--group-009.md#canonical-1333011212022301-2002223323000032-1012000310011211-2133211221002303-2303303323202232-2111003121220301-1330020031131032-2022300000032110)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop](data-sources--azure_vnet_site--reference--group-009.md#canonical-3311223101100123-0112100121200001-1122322212113233-3201331232031021-3303311022201202-0103303121332211-1330001200300203-1220221232033320)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](data-sources--azure_vnet_site--reference--group-009.md#canonical-2322030123023233-2001131000132330-0132220122210312-3222311222202003-2202011333132010-1121312311001210-2111310032222233-0211213133000202)
- voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack

<a id="canonical-3321311213210001-0103201330000123-3211100312303321-2232001222132333-0132202232222211-2221303313202112-3121212330202202-0132320000232120"></a>

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

<a id="canonical-0010221003331210-2201131200021301-3030110212303302-3232121303111310-3131131302132211-3020023100300321-1021103030220113-0123132102201103"></a>

## Direct properties — dual_stack / 232223002200 / 3

- [ipv4](data-sources--azure_vnet_site--reference--group-009.md#canonical-1310213012221010-2022102001322101-2013231322113133-0233133211333230-0311210121111132-3320110322003101-3020112111110332-2110023221322321): complete subsection reference.

- [ipv6](data-sources--azure_vnet_site--reference--group-009.md#canonical-2221300211210133-2021023313002110-1201232311332123-0330331011202201-1102321232321333-1022112220230031-0120130232330033-2230222133212111): complete subsection reference.

<a id="canonical-1302120023101120-3211303031310023-0330001021101010-0221231023010120-3110323211211003-1313103030132323-2321101210213213-0032001133020203"></a>

## Next pages — dual_stack / 232223002200 / 4

- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4](data-sources--azure_vnet_site--reference--group-009.md#canonical-1310213012221010-2022102001322101-2013231322113133-0233133211333230-0311210121111132-3320110322003101-3020112111110332-2110023221322321)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6](data-sources--azure_vnet_site--reference--group-009.md#canonical-2221300211210133-2021023313002110-1201232311332123-0330331011202201-1102321232321333-1022112220230031-0120130232330033-2230222133212111)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](data-sources--azure_vnet_site--reference--group-009.md#canonical-2322030123023233-2001131000132330-0132220122210312-3222311222202003-2202011333132010-1121312311001210-2111310032222233-0211213133000202)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)

<a id="canonical-1310213012221010-2022102001322101-2013231322113133-0233133211333230-0311210121111132-3320110322003101-3020112111110332-2110023221322321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3232221002100210-0312100323121120-0233020121200022-0111121112023032-0212002103312210-2123023122012102-2331213331201132-3220230333300301"></a>

## voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.IPv4 — IPv4 / 113103002212 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-2210100210332112-1233221121133301-0120203222303210-2221022223110220-0313230312203322-2232200300001103-2301220003133113-1333232230300113)
- [voltstack_cluster](data-sources--azure_vnet_site--reference--group-008.md#canonical-0023321213310231-1323112323303023-2220001003301330-3020021300200322-1322303031300210-2010300303132331-3120231121123332-0101313132130000)
- [voltstack_cluster.outside_static_routes](data-sources--azure_vnet_site--reference--group-009.md#canonical-2122010301323013-0003013121031331-1113323123213010-2220330111003201-1220212213113310-0010110023331303-2022322233101133-1321233223121330)
- [voltstack_cluster.outside_static_routes.static_route_list](data-sources--azure_vnet_site--reference--group-009.md#canonical-3003111000100013-3323000112311231-0023003330110223-2100010211100300-0031031230032100-0312001210330231-3122230101110102-3002010210221130)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route](data-sources--azure_vnet_site--reference--group-009.md#canonical-1333011212022301-2002223323000032-1012000310011211-2133211221002303-2303303323202232-2111003121220301-1330020031131032-2022300000032110)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop](data-sources--azure_vnet_site--reference--group-009.md#canonical-3311223101100123-0112100121200001-1122322212113233-3201331232031021-3303311022201202-0103303121332211-1330001200300203-1220221232033320)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](data-sources--azure_vnet_site--reference--group-009.md#canonical-2322030123023233-2001131000132330-0132220122210312-3222311222202003-2202011333132010-1121312311001210-2111310032222233-0211213133000202)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](data-sources--azure_vnet_site--reference--group-009.md#canonical-1313331112213121-0303132101303233-0200103221220203-0203122323131123-0312211110230112-2202320110102122-0332122003232333-3121311210332310)
- voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.IPv4

<a id="canonical-1103211302323003-3331130123232203-0011121022230021-0022231002130022-3132132101110231-0111331000011122-0032102011010202-2210200101302020"></a>

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

<a id="canonical-2222221010011202-1322111013311120-1221322332201023-3320231320011021-2301222311010111-2123131121130202-3332121013133302-1320112322020330"></a>

## Direct properties — IPv4 / 113103002212 / 3

<a id="canonical-3302222003122300-2013223301312023-1333231133021120-3102111312302301-2332310203003221-0230012322000102-2021122102033223-1111023000320010"></a>

<a id="canonical-1102132331002330-1312123210302030-0100003011202132-3103232130300231-1130310122302310-3200121112330002-1333311113332301-3203330220001212"></a>

## addr property — IPv4 / 113103002212 / 4

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

<a id="canonical-2202003020221122-0121020321120112-2302033032032022-0322031213220333-2100102030000311-3302013102032022-0002332321322033-2110221211302211"></a>

## Next pages — IPv4 / 113103002212 / 5

- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](data-sources--azure_vnet_site--reference--group-009.md#canonical-1313331112213121-0303132101303233-0200103221220203-0203122323131123-0312211110230112-2202320110102122-0332122003232333-3121311210332310)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)

<a id="canonical-2221300211210133-2021023313002110-1201232311332123-0330331011202201-1102321232321333-1022112220230031-0120130232330033-2230222133212111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3101133320031201-3110220202001313-2302111332132223-3102321133211300-1023002112231333-0123320231003311-3332203210123012-2322002002300103"></a>

## voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.IPv6 — IPv6 / 112111231101 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-2210100210332112-1233221121133301-0120203222303210-2221022223110220-0313230312203322-2232200300001103-2301220003133113-1333232230300113)
- [voltstack_cluster](data-sources--azure_vnet_site--reference--group-008.md#canonical-0023321213310231-1323112323303023-2220001003301330-3020021300200322-1322303031300210-2010300303132331-3120231121123332-0101313132130000)
- [voltstack_cluster.outside_static_routes](data-sources--azure_vnet_site--reference--group-009.md#canonical-2122010301323013-0003013121031331-1113323123213010-2220330111003201-1220212213113310-0010110023331303-2022322233101133-1321233223121330)
- [voltstack_cluster.outside_static_routes.static_route_list](data-sources--azure_vnet_site--reference--group-009.md#canonical-3003111000100013-3323000112311231-0023003330110223-2100010211100300-0031031230032100-0312001210330231-3122230101110102-3002010210221130)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route](data-sources--azure_vnet_site--reference--group-009.md#canonical-1333011212022301-2002223323000032-1012000310011211-2133211221002303-2303303323202232-2111003121220301-1330020031131032-2022300000032110)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop](data-sources--azure_vnet_site--reference--group-009.md#canonical-3311223101100123-0112100121200001-1122322212113233-3201331232031021-3303311022201202-0103303121332211-1330001200300203-1220221232033320)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](data-sources--azure_vnet_site--reference--group-009.md#canonical-2322030123023233-2001131000132330-0132220122210312-3222311222202003-2202011333132010-1121312311001210-2111310032222233-0211213133000202)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](data-sources--azure_vnet_site--reference--group-009.md#canonical-1313331112213121-0303132101303233-0200103221220203-0203122323131123-0312211110230112-2202320110102122-0332122003232333-3121311210332310)
- voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.IPv6

<a id="canonical-3020102303103203-1301110021222030-3133102221332201-3022113132030130-1213211123302303-3123332303133020-2212211003021021-2330110323313312"></a>

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

<a id="canonical-2320220023032130-1022231113333012-0323100101030231-3020012132022213-2122030002011031-3232303113112023-2011323331200132-1231213231221000"></a>

## Direct properties — IPv6 / 112111231101 / 3

<a id="canonical-2111222300311331-1331222302333012-1312103102201322-0323312021300031-1122023103302313-1202133122220103-3131100103233100-3211111130111002"></a>

<a id="canonical-0103320332213022-0300233122110303-3331122120231330-2301232011031113-0311023310113112-0010010303321310-1122011222110100-3020111003222030"></a>

## addr property — IPv6 / 112111231101 / 4

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

<a id="canonical-2223211302230033-1322223011221302-2112020001020333-2010122231230123-1012201311230002-1232012310322321-1211322121032200-3012132123200222"></a>

## Next pages — IPv6 / 112111231101 / 5

- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](data-sources--azure_vnet_site--reference--group-009.md#canonical-1313331112213121-0303132101303233-0200103221220203-0203122323131123-0312211110230112-2202320110102122-0332122003232333-3121311210332310)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)

<a id="canonical-3123111203020033-0230300100113302-3112111300123013-1202200301202300-0020201033311310-2032230103122303-0033313320320110-2332003303313213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1010032200303031-0022103332232212-1031223302033211-2311000111310332-0030122102313300-0202230330212111-0202123300100133-0011033030013201"></a>

## voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.IPv4 — IPv4 / 112200311013 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-2210100210332112-1233221121133301-0120203222303210-2221022223110220-0313230312203322-2232200300001103-2301220003133113-1333232230300113)
- [voltstack_cluster](data-sources--azure_vnet_site--reference--group-008.md#canonical-0023321213310231-1323112323303023-2220001003301330-3020021300200322-1322303031300210-2010300303132331-3120231121123332-0101313132130000)
- [voltstack_cluster.outside_static_routes](data-sources--azure_vnet_site--reference--group-009.md#canonical-2122010301323013-0003013121031331-1113323123213010-2220330111003201-1220212213113310-0010110023331303-2022322233101133-1321233223121330)
- [voltstack_cluster.outside_static_routes.static_route_list](data-sources--azure_vnet_site--reference--group-009.md#canonical-3003111000100013-3323000112311231-0023003330110223-2100010211100300-0031031230032100-0312001210330231-3122230101110102-3002010210221130)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route](data-sources--azure_vnet_site--reference--group-009.md#canonical-1333011212022301-2002223323000032-1012000310011211-2133211221002303-2303303323202232-2111003121220301-1330020031131032-2022300000032110)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop](data-sources--azure_vnet_site--reference--group-009.md#canonical-3311223101100123-0112100121200001-1122322212113233-3201331232031021-3303311022201202-0103303121332211-1330001200300203-1220221232033320)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](data-sources--azure_vnet_site--reference--group-009.md#canonical-2322030123023233-2001131000132330-0132220122210312-3222311222202003-2202011333132010-1121312311001210-2111310032222233-0211213133000202)
- voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.IPv4

<a id="canonical-0000220030110311-1201033103200031-3120123103203030-1230000032103032-1210330202001013-1230130320003121-1200203322010001-1101210201002113"></a>

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

<a id="canonical-1212011221130121-2211200201233012-2123323033320332-1323311111200313-0311203320211112-0220200322221321-0330310213023331-1100201231031313"></a>

## Direct properties — IPv4 / 112200311013 / 3

<a id="canonical-1302232131122212-0011102222101302-0000020231302032-3231003000022010-3213220302010230-3310103031313011-2223221211333110-2333303232032213"></a>

<a id="canonical-1110210212132300-1133332023110201-2232332212012231-2333130201023312-1102223303033313-2121201303003332-0312010001112332-1113030102222223"></a>

## addr property — IPv4 / 112200311013 / 4

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

<a id="canonical-3012132132313011-3323013312121210-3303021120221011-2220321032112021-3212300002313103-3030310022022213-3200221222303213-2312301003100203"></a>

## Next pages — IPv4 / 112200311013 / 5

- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](data-sources--azure_vnet_site--reference--group-009.md#canonical-2322030123023233-2001131000132330-0132220122210312-3222311222202003-2202011333132010-1121312311001210-2111310032222233-0211213133000202)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)

<a id="canonical-2000201333231230-0223330132223032-2122032221311010-3200211303121021-0133121212303321-0332303002200200-1103330101122002-1120121212312000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1131233313223303-2212131313333020-1223301113321300-2313233011212113-0303110202030200-2022231232021033-1210220122222121-0332123031330312"></a>

## voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.IPv6 — IPv6 / 323212023100 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-2210100210332112-1233221121133301-0120203222303210-2221022223110220-0313230312203322-2232200300001103-2301220003133113-1333232230300113)
- [voltstack_cluster](data-sources--azure_vnet_site--reference--group-008.md#canonical-0023321213310231-1323112323303023-2220001003301330-3020021300200322-1322303031300210-2010300303132331-3120231121123332-0101313132130000)
- [voltstack_cluster.outside_static_routes](data-sources--azure_vnet_site--reference--group-009.md#canonical-2122010301323013-0003013121031331-1113323123213010-2220330111003201-1220212213113310-0010110023331303-2022322233101133-1321233223121330)
- [voltstack_cluster.outside_static_routes.static_route_list](data-sources--azure_vnet_site--reference--group-009.md#canonical-3003111000100013-3323000112311231-0023003330110223-2100010211100300-0031031230032100-0312001210330231-3122230101110102-3002010210221130)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route](data-sources--azure_vnet_site--reference--group-009.md#canonical-1333011212022301-2002223323000032-1012000310011211-2133211221002303-2303303323202232-2111003121220301-1330020031131032-2022300000032110)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop](data-sources--azure_vnet_site--reference--group-009.md#canonical-3311223101100123-0112100121200001-1122322212113233-3201331232031021-3303311022201202-0103303121332211-1330001200300203-1220221232033320)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](data-sources--azure_vnet_site--reference--group-009.md#canonical-2322030123023233-2001131000132330-0132220122210312-3222311222202003-2202011333132010-1121312311001210-2111310032222233-0211213133000202)
- voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.IPv6

<a id="canonical-3330202013321222-2323333111212233-1220103331212231-1033120213120202-0132313011320103-1230231300333002-1021111030013013-2121123110033033"></a>

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

<a id="canonical-1120203331331131-1233131200223010-1122030121110313-3132333131310013-0002113200010233-0210302120001331-3232330030320022-2100122010330100"></a>

## Direct properties — IPv6 / 323212023100 / 3

<a id="canonical-3110033002102110-3210233012100001-3033200232123212-1213100113212033-2220221012211220-2202300321200310-0100101213320112-3020110131220233"></a>

<a id="canonical-0322003311231021-2033233202311101-3031211113201120-0323202213332230-3213130022010213-1220302011323302-3032121000320013-2210002132020233"></a>

## addr property — IPv6 / 323212023100 / 4

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

<a id="canonical-3030210211121312-1103210022100212-2020021311211321-2210203031330330-3300103313122020-2313300031211112-1012333231311003-3212301221120120"></a>

## Next pages — IPv6 / 323212023100 / 5

- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](data-sources--azure_vnet_site--reference--group-009.md#canonical-2322030123023233-2001131000132330-0132220122210312-3222311222202003-2202011333132010-1121312311001210-2111310032222233-0211213133000202)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)

<a id="canonical-2113023020113202-1220110303201212-2300303332110112-2301000001202231-3323122003201022-0221020000030212-1223101333201302-2002021122212131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0332123112322213-3321330000101312-0203223211011011-3300313320031302-0200113210310222-1131320222010301-3020023111001101-1113110000000001"></a>

## voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets — subnets / 203033230201 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-2210100210332112-1233221121133301-0120203222303210-2221022223110220-0313230312203322-2232200300001103-2301220003133113-1333232230300113)
- [voltstack_cluster](data-sources--azure_vnet_site--reference--group-008.md#canonical-0023321213310231-1323112323303023-2220001003301330-3020021300200322-1322303031300210-2010300303132331-3120231121123332-0101313132130000)
- [voltstack_cluster.outside_static_routes](data-sources--azure_vnet_site--reference--group-009.md#canonical-2122010301323013-0003013121031331-1113323123213010-2220330111003201-1220212213113310-0010110023331303-2022322233101133-1321233223121330)
- [voltstack_cluster.outside_static_routes.static_route_list](data-sources--azure_vnet_site--reference--group-009.md#canonical-3003111000100013-3323000112311231-0023003330110223-2100010211100300-0031031230032100-0312001210330231-3122230101110102-3002010210221130)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route](data-sources--azure_vnet_site--reference--group-009.md#canonical-1333011212022301-2002223323000032-1012000310011211-2133211221002303-2303303323202232-2111003121220301-1330020031131032-2022300000032110)
- voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets

<a id="canonical-3203300133112230-0100002302303331-0200330130321331-1221101203032220-1031311010210101-0311233232023003-0312113132120012-2332030313012312"></a>

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

<a id="canonical-0301131230101023-2020123030033210-0022333333213031-1032313200003323-0120300303033211-3002032331323001-2022032223321013-0303120302310233"></a>

## Direct properties — subnets / 203033230201 / 3

- [ipv4](data-sources--azure_vnet_site--reference--group-009.md#canonical-2202123011131330-2222232212320320-1003123332010200-0333000331030130-0113131311123230-1302033001330113-1333321232002112-1231130002310021): complete subsection reference.

- [ipv6](data-sources--azure_vnet_site--reference--group-009.md#canonical-2302122331330333-2301113121110032-0132002323130303-1203012122223321-0310231122332200-0002311203133220-3313232112210212-0121002013032023): complete subsection reference.

<a id="canonical-3333011333222301-3020313132323123-0211330022132200-3002302313133330-2013013000023311-3013133323113333-0030312000220230-3022120203333311"></a>

## Next pages — subnets / 203033230201 / 4

- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4](data-sources--azure_vnet_site--reference--group-009.md#canonical-2202123011131330-2222232212320320-1003123332010200-0333000331030130-0113131311123230-1302033001330113-1333321232002112-1231130002310021)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6](data-sources--azure_vnet_site--reference--group-009.md#canonical-2302122331330333-2301113121110032-0132002323130303-1203012122223321-0310231122332200-0002311203133220-3313232112210212-0121002013032023)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route](data-sources--azure_vnet_site--reference--group-009.md#canonical-1333011212022301-2002223323000032-1012000310011211-2133211221002303-2303303323202232-2111003121220301-1330020031131032-2022300000032110)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)

<a id="canonical-2202123011131330-2222232212320320-1003123332010200-0333000331030130-0113131311123230-1302033001330113-1333321232002112-1231130002310021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3133312302023330-3221312321200003-2201030303303003-0132110030212323-3323121223101101-3022330331223233-2233222302301010-2233303213112200"></a>

## voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets.IPv4 — IPv4 / 203312320232 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-2210100210332112-1233221121133301-0120203222303210-2221022223110220-0313230312203322-2232200300001103-2301220003133113-1333232230300113)
- [voltstack_cluster](data-sources--azure_vnet_site--reference--group-008.md#canonical-0023321213310231-1323112323303023-2220001003301330-3020021300200322-1322303031300210-2010300303132331-3120231121123332-0101313132130000)
- [voltstack_cluster.outside_static_routes](data-sources--azure_vnet_site--reference--group-009.md#canonical-2122010301323013-0003013121031331-1113323123213010-2220330111003201-1220212213113310-0010110023331303-2022322233101133-1321233223121330)
- [voltstack_cluster.outside_static_routes.static_route_list](data-sources--azure_vnet_site--reference--group-009.md#canonical-3003111000100013-3323000112311231-0023003330110223-2100010211100300-0031031230032100-0312001210330231-3122230101110102-3002010210221130)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route](data-sources--azure_vnet_site--reference--group-009.md#canonical-1333011212022301-2002223323000032-1012000310011211-2133211221002303-2303303323202232-2111003121220301-1330020031131032-2022300000032110)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets](data-sources--azure_vnet_site--reference--group-009.md#canonical-2113023020113202-1220110303201212-2300303332110112-2301000001202231-3323122003201022-0221020000030212-1223101333201302-2002021122212131)
- voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets.IPv4

<a id="canonical-1111230022020232-3311132132123222-2322303133221202-0021303320323302-2323033232320010-1030323210303033-3300122313333323-2030112201011132"></a>

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

<a id="canonical-1012012320210220-2212101203321132-2230113100301320-3120321010321332-0113130002200223-3121022110313123-0130011030303020-0113020211122002"></a>

## Direct properties — IPv4 / 203312320232 / 3

<a id="canonical-0200332300110112-1211222203211302-0321021232232021-2232220132101210-1110103230130031-0120031110301232-3320300210112113-2031332110323022"></a>

<a id="canonical-1030300032103102-2030113132223131-1102022032020121-0121222302000001-3023303033211101-3131013203110132-0210102101330210-0212330132001102"></a>

## plen property — IPv4 / 203312320232 / 4

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

<a id="canonical-2003221131311232-3330112320233311-0310211223032320-3001122032211233-2032100331020333-3233332210330032-0311210122220213-2300120310320201"></a>

<a id="canonical-3210223131300330-2332211211213001-2213113010300321-3102001110021112-2322011013021320-2300313331130031-3211311331302030-0110320131033120"></a>

## prefix property — IPv4 / 203312320232 / 5

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

<a id="canonical-1322112321001311-3230000221102322-0000101303030012-3221213132333323-2003302021332232-2033021101222331-0110202102111033-1201031301221012"></a>

## Next pages — IPv4 / 203312320232 / 6

- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets](data-sources--azure_vnet_site--reference--group-009.md#canonical-2113023020113202-1220110303201212-2300303332110112-2301000001202231-3323122003201022-0221020000030212-1223101333201302-2002021122212131)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)

<a id="canonical-2302122331330333-2301113121110032-0132002323130303-1203012122223321-0310231122332200-0002311203133220-3313232112210212-0121002013032023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0220002323113030-2203120323220030-1131102222122133-1120233213302333-0013031032220121-3312121120303311-0103313332320100-3323103121310310"></a>

## voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets.IPv6 — IPv6 / 012201333310 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-2210100210332112-1233221121133301-0120203222303210-2221022223110220-0313230312203322-2232200300001103-2301220003133113-1333232230300113)
- [voltstack_cluster](data-sources--azure_vnet_site--reference--group-008.md#canonical-0023321213310231-1323112323303023-2220001003301330-3020021300200322-1322303031300210-2010300303132331-3120231121123332-0101313132130000)
- [voltstack_cluster.outside_static_routes](data-sources--azure_vnet_site--reference--group-009.md#canonical-2122010301323013-0003013121031331-1113323123213010-2220330111003201-1220212213113310-0010110023331303-2022322233101133-1321233223121330)
- [voltstack_cluster.outside_static_routes.static_route_list](data-sources--azure_vnet_site--reference--group-009.md#canonical-3003111000100013-3323000112311231-0023003330110223-2100010211100300-0031031230032100-0312001210330231-3122230101110102-3002010210221130)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route](data-sources--azure_vnet_site--reference--group-009.md#canonical-1333011212022301-2002223323000032-1012000310011211-2133211221002303-2303303323202232-2111003121220301-1330020031131032-2022300000032110)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets](data-sources--azure_vnet_site--reference--group-009.md#canonical-2113023020113202-1220110303201212-2300303332110112-2301000001202231-3323122003201022-0221020000030212-1223101333201302-2002021122212131)
- voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets.IPv6

<a id="canonical-1322130000333310-3212203311211312-1123130011101221-3011103201020220-2230130103213330-1010111031332233-3123002110123112-2320002001212200"></a>

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

<a id="canonical-0112333211123312-3133333201320031-0310123230100122-0231022003310120-3133223032233102-2102320331113102-0322101000132101-0213122100020331"></a>

## Direct properties — IPv6 / 012201333310 / 3

<a id="canonical-3202002203330022-1331132300003101-0013212010013203-1303030113212230-0233210032010222-0121233203201202-2120203320213022-1022333312110012"></a>

<a id="canonical-3201221330021202-0200122202132311-3320310131230310-0302033112003323-3120010221122123-2310200223101103-2333221103002300-2121203210230323"></a>

## plen property — IPv6 / 012201333310 / 4

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

<a id="canonical-2202012223130231-3203310213001130-2110302221200321-2030322223230323-3133011323331223-2201022021201131-3210010301023233-3021212111021112"></a>

<a id="canonical-3332230120010233-3223223213203301-0131312233333313-3130212222213130-3111303031102013-2212211320303002-1321221122330313-3000212103320323"></a>

## prefix property — IPv6 / 012201333310 / 5

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

<a id="canonical-0221001103100030-1233030010022021-2120303111021222-2231100330330131-3322103012022020-3302031332222320-1220203100303013-3320111033212300"></a>

## Next pages — IPv6 / 012201333310 / 6

- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets](data-sources--azure_vnet_site--reference--group-009.md#canonical-2113023020113202-1220110303201212-2300303332110112-2301000001202231-3323122003201022-0221020000030212-1223101333201302-2002021122212131)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)

<a id="canonical-2103312120131130-0311221103003301-2032012022313031-3022112013222312-0332012203320100-1201301121112002-3022201012300030-2230100233031130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1122231231022011-2233011122123002-2030200233323020-0123232031220220-1200320130032013-0031211103323310-1232021033103202-0213300320132212"></a>

## voltstack_cluster.sm_connection_public_ip — sm_connection_public_ip / 203030021123 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-2210100210332112-1233221121133301-0120203222303210-2221022223110220-0313230312203322-2232200300001103-2301220003133113-1333232230300113)
- [voltstack_cluster](data-sources--azure_vnet_site--reference--group-008.md#canonical-0023321213310231-1323112323303023-2220001003301330-3020021300200322-1322303031300210-2010300303132331-3120231121123332-0101313132130000)
- voltstack_cluster.sm_connection_public_ip

<a id="canonical-0023203130100313-0130102020130320-2320030223110220-0023112312110311-0311213323120120-0211112132300332-2110022011300112-1012002223001123"></a>

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

<a id="canonical-0121312220122331-0213331132132200-1312211301022013-1213132310121303-3320330203302000-0203122321223113-1131333133202030-3221321203223102"></a>

## Direct properties — sm_connection_public_ip / 203030021123 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1302333112312322-0233130231231302-2222022230011112-2011223132211201-3021320202120213-2330023002011101-1223110112300010-0121321303010021"></a>

## Next pages — sm_connection_public_ip / 203030021123 / 4

- [voltstack_cluster](data-sources--azure_vnet_site--reference--group-008.md#canonical-0023321213310231-1323112323303023-2220001003301330-3020021300200322-1322303031300210-2010300303132331-3120231121123332-0101313132130000)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)

<a id="canonical-0211033101110033-2320122112231132-2110002313222122-0101211122122231-0321030002111220-0333312133121102-0003110103201133-1132032322113331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1033220023011131-2233310311002302-3210001013010121-0321001003321200-0310212323310333-2321033213210023-3001333303000330-0321013201233113"></a>

## voltstack_cluster.sm_connection_pvt_ip — sm_connection_pvt_ip / 311023122301 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-2210100210332112-1233221121133301-0120203222303210-2221022223110220-0313230312203322-2232200300001103-2301220003133113-1333232230300113)
- [voltstack_cluster](data-sources--azure_vnet_site--reference--group-008.md#canonical-0023321213310231-1323112323303023-2220001003301330-3020021300200322-1322303031300210-2010300303132331-3120231121123332-0101313132130000)
- voltstack_cluster.sm_connection_pvt_ip

<a id="canonical-3130302201030300-3001331313223213-3300220332123333-3033331001021132-1321103021303312-3313233202231211-1313103230102131-2230023320021213"></a>

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

<a id="canonical-2321233001301332-2302302131223121-3321310230300121-3201100033023101-2312303212303333-0003031100200132-1323112221002202-1101131332012110"></a>

## Direct properties — sm_connection_pvt_ip / 311023122301 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2211131021023032-1010303000120310-2010110312321111-2033203110203332-0123321012313303-1311030223213130-0212022332111120-0031030232001023"></a>

## Next pages — sm_connection_pvt_ip / 311023122301 / 4

- [voltstack_cluster](data-sources--azure_vnet_site--reference--group-008.md#canonical-0023321213310231-1323112323303023-2220001003301330-3020021300200322-1322303031300210-2010300303132331-3120231121123332-0101313132130000)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)

<a id="canonical-2211230110200311-3232130012213121-1311020220010320-1010202121311103-1101232001230330-0100002100113133-1013322233220300-3012112231002223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1101332002121220-0333123013300202-1210330330021221-2012233200021002-1023223020001003-3203113112011103-3302203332332322-2320023310031213"></a>

## voltstack_cluster.storage_class_list — storage_class_list / 203312203033 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-2210100210332112-1233221121133301-0120203222303210-2221022223110220-0313230312203322-2232200300001103-2301220003133113-1333232230300113)
- [voltstack_cluster](data-sources--azure_vnet_site--reference--group-008.md#canonical-0023321213310231-1323112323303023-2220001003301330-3020021300200322-1322303031300210-2010300303132331-3120231121123332-0101313132130000)
- voltstack_cluster.storage_class_list

<a id="canonical-3101333311032033-2100111002213303-2113222232030311-1032300300101222-1012032100112102-2330012120121013-3213202232312221-0013120201311201"></a>

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

<a id="canonical-0320213012123302-0010200033233221-2203023000123110-0311113102210213-3301301101200001-3303002101021212-1032310300112021-1003222002020222"></a>

## Direct properties — storage_class_list / 203312203033 / 3

- [storage_classes](data-sources--azure_vnet_site--reference--group-009.md#canonical-2033210312113201-3103122020303021-1302100210012321-1213303310323001-2222211203232131-0131122211022010-0303200000010020-0012032011032130): complete subsection reference.

<a id="canonical-1322130000323022-3233123322101203-3302301012001201-0211102210331100-0110211311020302-1220330120323101-3323331211303221-0132231030203123"></a>

## Next pages — storage_class_list / 203312203033 / 4

- [voltstack_cluster.storage_class_list.storage_classes](data-sources--azure_vnet_site--reference--group-009.md#canonical-2033210312113201-3103122020303021-1302100210012321-1213303310323001-2222211203232131-0131122211022010-0303200000010020-0012032011032130)
- [voltstack_cluster](data-sources--azure_vnet_site--reference--group-008.md#canonical-0023321213310231-1323112323303023-2220001003301330-3020021300200322-1322303031300210-2010300303132331-3120231121123332-0101313132130000)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)

<a id="canonical-2033210312113201-3103122020303021-1302100210012321-1213303310323001-2222211203232131-0131122211022010-0303200000010020-0012032011032130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0320002113130122-1123020232212030-1113220010222323-0303313310001122-1311132221111020-3112002322321031-0323031000230222-1031011233310121"></a>

## voltstack_cluster.storage_class_list.storage_classes — storage_classes / 230131211321 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-2210100210332112-1233221121133301-0120203222303210-2221022223110220-0313230312203322-2232200300001103-2301220003133113-1333232230300113)
- [voltstack_cluster](data-sources--azure_vnet_site--reference--group-008.md#canonical-0023321213310231-1323112323303023-2220001003301330-3020021300200322-1322303031300210-2010300303132331-3120231121123332-0101313132130000)
- [voltstack_cluster.storage_class_list](data-sources--azure_vnet_site--reference--group-009.md#canonical-2211230110200311-3232130012213121-1311020220010320-1010202121311103-1101232001230330-0100002100113133-1013322233220300-3012112231002223)
- voltstack_cluster.storage_class_list.storage_classes

<a id="canonical-2030102303231311-1123031011310302-0132022233103023-1323232211122211-2121023111120121-2022200103233201-2313313311322223-2133212010213310"></a>

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

<a id="canonical-0331022113233313-3200001102121032-1030211212331233-2311221333312130-1022122130331320-0100331010300313-2030320010200232-3200033310133330"></a>

## Direct properties — storage_classes / 230131211321 / 3

<a id="canonical-2031333133001131-3200222221020111-1312220032112013-2111012111300120-3231102202211332-0022023133022032-1023022222210203-3003002133223301"></a>

<a id="canonical-3222201300030100-1203212010332202-2020332333201031-2130323300201333-1020112233232321-3130013320320131-0001230220102213-3201131132230123"></a>

## default_storage_class property — storage_classes / 230131211321 / 4

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

<a id="canonical-2322123233213311-1212022201230133-3332030110313012-2313200000332130-0231131133233103-1022021011022211-0311003232202002-1131200322131203"></a>

<a id="canonical-0223120303303200-2021322030301101-0011012011212020-0321221022010231-0120003130310333-3233033333202232-0312131100210123-2330132302113012"></a>

## storage_class_name property — storage_classes / 230131211321 / 5

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

<a id="canonical-2330121132101230-2023221030301202-3112201232113303-1122013201001123-0331211132211313-2123213103033012-1333303210102023-1300133120303033"></a>

## Next pages — storage_classes / 230131211321 / 6

- [voltstack_cluster.storage_class_list](data-sources--azure_vnet_site--reference--group-009.md#canonical-2211230110200311-3232130012213121-1311020220010320-1010202121311103-1101232001230330-0100002100113133-1013322233220300-3012112231002223)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)

<a id="canonical-0013331330233123-3030030001210332-2301323022202203-2230301230330203-2310333012300123-0110311130202131-0121011120233100-0122113300012113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1133020030230330-0013031021233013-3130103012212112-0103130132221213-0311011132102123-2221232133310010-0031202021313120-0120300300101013"></a>

## voltstack_cluster_ar — voltstack_cluster_ar / 331332202112 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-2210100210332112-1233221121133301-0120203222303210-2221022223110220-0313230312203322-2232200300001103-2301220003133113-1333232230300113)
- voltstack_cluster_ar

<a id="canonical-1132033131013300-3130323102021333-0310010111313220-1101103121102232-0213210000200331-0032222230332302-1320322010211031-2321103223120102"></a>

Type: `"single"`. Computed.

App Stack Cluster of single interface Azure nodes.

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

<a id="canonical-1022312301213100-0122123122112200-0010122312322130-1021110100222313-2132200023302021-1012033100120130-3210112210311131-1112012133303322"></a>

## Direct properties — voltstack_cluster_ar / 331332202112 / 3

- [accelerated_networking](data-sources--azure_vnet_site--reference--group-009.md#canonical-3111320123132202-0111023220022221-2133113121010030-2330203201232211-1210211131300010-3032000021100011-0310321112312100-3232023011332012): complete subsection reference.

- [active_enhanced_firewall_policies](data-sources--azure_vnet_site--reference--group-009.md#canonical-3130311313313102-1223301303022202-0311100110000220-0310130012121123-2131330210121230-2023322103123321-1331013223013013-3231121003310032): complete subsection reference.

- [active_forward_proxy_policies](data-sources--azure_vnet_site--reference--group-009.md#canonical-2211230103012213-0332121122313000-3220010131233032-3312210331230021-3000311103133131-0132232121020210-2302112112112311-1320121110212001): complete subsection reference.

- [active_network_policies](data-sources--azure_vnet_site--reference--group-009.md#canonical-2111312233203311-0010131101022301-0232021221220131-0023122023021222-3110301021302201-0002313230230303-2032003103232221-1230301111133132): complete subsection reference.

<a id="canonical-2110333321230011-0223031220123200-2032032303130022-0300013133000031-2022001232002012-3022032033330302-2012310201320323-0113021310332113"></a>

<a id="canonical-2313203121202301-2221010013221122-3010221030330302-2111332113112210-3300201031131322-2223130300001010-3212210022311222-3013023112333222"></a>

## azure_certified_hw property — voltstack_cluster_ar / 331332202112 / 4

Type: `"string"`. Computed.

\[Enum: Azure-byol-voltstack-combo\] Azure Certified Hardware. Name for Azure certified hardware.
The only possible value is \`azure-byol-voltstack-combo\`.

Upstream description:

Name for Azure certified hardware.

Receipt-pinned upstream constraints:

```json
{
  "enum": [
    "azure-byol-voltstack-combo"
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
    "ves.io.schema.rules.string.in": "[\\\"azure-byol-voltstack-combo\\\"]",
    "ves.io.schema.rules.string.max_len": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.in": "[\\\"azure-byol-voltstack-combo\\\"]",
    "ves.io.schema.rules.string.max_len": "64"
  }
}
```

- [dc_cluster_group](data-sources--azure_vnet_site--reference--group-009.md#canonical-2300233031212012-1232133331121210-0231101211002203-0312000030101310-3220033331333210-1003301023311120-1203212030031110-2000301112300331): complete subsection reference.

- [default_storage](data-sources--azure_vnet_site--reference--group-009.md#canonical-3021110101311100-1223232022130112-1111320311302020-0202121003310301-1000221312200212-1020311233003003-1002010330131110-3322033203002112): complete subsection reference.

- [forward_proxy_allow_all](data-sources--azure_vnet_site--reference--group-009.md#canonical-1003032211122222-0321233202101030-2132030131310302-2232003312202112-1121300211312230-2101233232311113-2302311132300131-1333233122200221): complete subsection reference.

- [global_network_list](data-sources--azure_vnet_site--reference--group-009.md#canonical-3201222313330302-0111330002310231-0113202002012313-3311012321312000-0133101333213210-2221121023210211-1211323221202222-3230133013021203): complete subsection reference.

- [k8s_cluster](data-sources--azure_vnet_site--reference--group-009.md#canonical-2210130302333102-1100012131310231-2210333031222022-1233210202101120-1333223123311122-3033333032332232-3220023213223232-0033332300032221): complete subsection reference.

- [no_dc_cluster_group](data-sources--azure_vnet_site--reference--group-009.md#canonical-2232123200303133-2110101003120023-3302131323101001-2013020002133320-3010303112002323-1100012011010232-2032102310220301-0000313000211123): complete subsection reference.

- [no_forward_proxy](data-sources--azure_vnet_site--reference--group-009.md#canonical-1220213212302011-1111032113103010-3303011132100112-2003203300112313-3313001323110130-1232112200112330-2002133212232031-2110330120213200): complete subsection reference.

- [no_global_network](data-sources--azure_vnet_site--reference--group-009.md#canonical-3103233133102122-3202232301232023-3111301313211310-0001012121230013-3131122323222331-2212212130231230-2323021213233202-2113022210021320): complete subsection reference.

- [no_k8s_cluster](data-sources--azure_vnet_site--reference--group-009.md#canonical-0300213333023101-1320213301033231-2123212001333312-3313222011323111-1312102012202032-1233012021032330-2320122031112023-0113213113303300): complete subsection reference.

- [no_network_policy](data-sources--azure_vnet_site--reference--group-009.md#canonical-3131333313020120-3030311131123210-1130312011123300-0022213231123232-0010320322000333-2313000010320333-3302131010332301-2301310021313300): complete subsection reference.

- [no_outside_static_routes](data-sources--azure_vnet_site--reference--group-009.md#canonical-0101302123000232-0121202321203200-2032301212222100-1101320230302321-3311111221102030-2100313030212333-1202023333300000-2200220000013211): complete subsection reference.

- [node](data-sources--azure_vnet_site--reference--group-009.md#canonical-2001203102000121-3022320233001111-0223132200203002-3110311030033021-1333212300230100-1303202000011321-1201101031310132-3003123330033331): complete subsection reference.

- [outside_static_routes](data-sources--azure_vnet_site--reference--group-010.md#canonical-2000233322000222-3220013132332003-1130221330212230-0332013011222213-1102302302030220-2321302122222103-1113111312230003-1100111110130022): complete subsection reference.

- [sm_connection_public_ip](data-sources--azure_vnet_site--reference--group-010.md#canonical-1110131212223003-1133321131023213-0003333011201323-1023111020320312-3030331031101200-2333320122033131-2021333333032321-3120312230113301): complete subsection reference.

- [sm_connection_pvt_ip](data-sources--azure_vnet_site--reference--group-010.md#canonical-0101231230133230-0333012001301211-2021201233132202-2031222110101312-3222321023121312-2302313333201023-0102233313023002-2013212030332112): complete subsection reference.

- [storage_class_list](data-sources--azure_vnet_site--reference--group-010.md#canonical-1210200133003231-2120022200111332-3101211201010121-0230323020031300-1331003211200131-1321121013311301-0130020230201101-3301222203001333): complete subsection reference.

<a id="canonical-2032130211102030-2032333000210311-0001032022011310-2323303233001021-0212211333102132-1010010330223110-3122021011120123-2201333100231001"></a>

## Next pages — voltstack_cluster_ar / 331332202112 / 5

- [voltstack_cluster_ar.accelerated_networking](data-sources--azure_vnet_site--reference--group-009.md#canonical-3111320123132202-0111023220022221-2133113121010030-2330203201232211-1210211131300010-3032000021100011-0310321112312100-3232023011332012)
- [voltstack_cluster_ar.active_enhanced_firewall_policies](data-sources--azure_vnet_site--reference--group-009.md#canonical-3130311313313102-1223301303022202-0311100110000220-0310130012121123-2131330210121230-2023322103123321-1331013223013013-3231121003310032)
- [voltstack_cluster_ar.active_forward_proxy_policies](data-sources--azure_vnet_site--reference--group-009.md#canonical-2211230103012213-0332121122313000-3220010131233032-3312210331230021-3000311103133131-0132232121020210-2302112112112311-1320121110212001)
- [voltstack_cluster_ar.active_network_policies](data-sources--azure_vnet_site--reference--group-009.md#canonical-2111312233203311-0010131101022301-0232021221220131-0023122023021222-3110301021302201-0002313230230303-2032003103232221-1230301111133132)
- [voltstack_cluster_ar.dc_cluster_group](data-sources--azure_vnet_site--reference--group-009.md#canonical-2300233031212012-1232133331121210-0231101211002203-0312000030101310-3220033331333210-1003301023311120-1203212030031110-2000301112300331)
- [voltstack_cluster_ar.default_storage](data-sources--azure_vnet_site--reference--group-009.md#canonical-3021110101311100-1223232022130112-1111320311302020-0202121003310301-1000221312200212-1020311233003003-1002010330131110-3322033203002112)
- [voltstack_cluster_ar.forward_proxy_allow_all](data-sources--azure_vnet_site--reference--group-009.md#canonical-1003032211122222-0321233202101030-2132030131310302-2232003312202112-1121300211312230-2101233232311113-2302311132300131-1333233122200221)
- [voltstack_cluster_ar.global_network_list](data-sources--azure_vnet_site--reference--group-009.md#canonical-3201222313330302-0111330002310231-0113202002012313-3311012321312000-0133101333213210-2221121023210211-1211323221202222-3230133013021203)
- [voltstack_cluster_ar.k8s_cluster](data-sources--azure_vnet_site--reference--group-009.md#canonical-2210130302333102-1100012131310231-2210333031222022-1233210202101120-1333223123311122-3033333032332232-3220023213223232-0033332300032221)
- [voltstack_cluster_ar.no_dc_cluster_group](data-sources--azure_vnet_site--reference--group-009.md#canonical-2232123200303133-2110101003120023-3302131323101001-2013020002133320-3010303112002323-1100012011010232-2032102310220301-0000313000211123)
- [voltstack_cluster_ar.no_forward_proxy](data-sources--azure_vnet_site--reference--group-009.md#canonical-1220213212302011-1111032113103010-3303011132100112-2003203300112313-3313001323110130-1232112200112330-2002133212232031-2110330120213200)
- [voltstack_cluster_ar.no_global_network](data-sources--azure_vnet_site--reference--group-009.md#canonical-3103233133102122-3202232301232023-3111301313211310-0001012121230013-3131122323222331-2212212130231230-2323021213233202-2113022210021320)
- [voltstack_cluster_ar.no_k8s_cluster](data-sources--azure_vnet_site--reference--group-009.md#canonical-0300213333023101-1320213301033231-2123212001333312-3313222011323111-1312102012202032-1233012021032330-2320122031112023-0113213113303300)
- [voltstack_cluster_ar.no_network_policy](data-sources--azure_vnet_site--reference--group-009.md#canonical-3131333313020120-3030311131123210-1130312011123300-0022213231123232-0010320322000333-2313000010320333-3302131010332301-2301310021313300)
- [voltstack_cluster_ar.no_outside_static_routes](data-sources--azure_vnet_site--reference--group-009.md#canonical-0101302123000232-0121202321203200-2032301212222100-1101320230302321-3311111221102030-2100313030212333-1202023333300000-2200220000013211)
- [voltstack_cluster_ar.node](data-sources--azure_vnet_site--reference--group-009.md#canonical-2001203102000121-3022320233001111-0223132200203002-3110311030033021-1333212300230100-1303202000011321-1201101031310132-3003123330033331)
- [voltstack_cluster_ar.outside_static_routes](data-sources--azure_vnet_site--reference--group-010.md#canonical-2000233322000222-3220013132332003-1130221330212230-0332013011222213-1102302302030220-2321302122222103-1113111312230003-1100111110130022)
- [voltstack_cluster_ar.sm_connection_public_ip](data-sources--azure_vnet_site--reference--group-010.md#canonical-1110131212223003-1133321131023213-0003333011201323-1023111020320312-3030331031101200-2333320122033131-2021333333032321-3120312230113301)
- [voltstack_cluster_ar.sm_connection_pvt_ip](data-sources--azure_vnet_site--reference--group-010.md#canonical-0101231230133230-0333012001301211-2021201233132202-2031222110101312-3222321023121312-2302313333201023-0102233313023002-2013212030332112)
- [voltstack_cluster_ar.storage_class_list](data-sources--azure_vnet_site--reference--group-010.md#canonical-1210200133003231-2120022200111332-3101211201010121-0230323020031300-1331003211200131-1321121013311301-0130020230201101-3301222203001333)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-2210100210332112-1233221121133301-0120203222303210-2221022223110220-0313230312203322-2232200300001103-2301220003133113-1333232230300113)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)

<a id="canonical-3111320123132202-0111023220022221-2133113121010030-2330203201232211-1210211131300010-3032000021100011-0310321112312100-3232023011332012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1310003301313030-1102020021122332-1310333212011322-0220320022211033-0103102202333123-3023001210202130-1121031312031210-2211010323302231"></a>

## voltstack_cluster_ar.accelerated_networking — accelerated_networking / 222211320221 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-2210100210332112-1233221121133301-0120203222303210-2221022223110220-0313230312203322-2232200300001103-2301220003133113-1333232230300113)
- [voltstack_cluster_ar](data-sources--azure_vnet_site--reference--group-009.md#canonical-0013331330233123-3030030001210332-2301323022202203-2230301230330203-2310333012300123-0110311130202131-0121011120233100-0122113300012113)
- voltstack_cluster_ar.accelerated_networking

<a id="canonical-2311330210313021-1213331333300212-2012302131221021-0213103113200302-1333002102001021-2220021031122300-3131312223131133-0330203122133323"></a>

Type: `"single"`. Computed.

Accelerated Networking to reduce Latency, When Mode is toggled, traffic disruption will be seen.

Upstream description:

Accelerated Networking to reduce Latency, When Mode is toggled, traffic disruption will be seen.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-accelerated_networking": "[\"disable\",\"enable\"]"
}
```

<a id="canonical-2112123223312022-0300232213330103-1222131320122230-3212230011000212-0302203110112110-2220122312303222-1013000221013001-1111121313003202"></a>

## Direct properties — accelerated_networking / 222211320221 / 3

- [disable_spec](data-sources--azure_vnet_site--reference--group-009.md#canonical-0210131321111210-3311320300230211-0331301203202122-3320023223011313-2210231130123031-1012300011231023-2030203130103121-0121012113310032): complete subsection reference.

- [enable](data-sources--azure_vnet_site--reference--group-009.md#canonical-1030213322013130-0320311022332002-3312323001321131-1133333232000132-1222313331203311-1310021101222120-2012302232131313-0113210010113011): complete subsection reference.

<a id="canonical-0030220103220033-2331200101033230-2113000121312103-1322030303300303-2321031010022012-3011100003300030-2021323300300310-2202230230112201"></a>

## Next pages — accelerated_networking / 222211320221 / 4

- [voltstack_cluster_ar.accelerated_networking.disable_spec](data-sources--azure_vnet_site--reference--group-009.md#canonical-0210131321111210-3311320300230211-0331301203202122-3320023223011313-2210231130123031-1012300011231023-2030203130103121-0121012113310032)
- [voltstack_cluster_ar.accelerated_networking.enable](data-sources--azure_vnet_site--reference--group-009.md#canonical-1030213322013130-0320311022332002-3312323001321131-1133333232000132-1222313331203311-1310021101222120-2012302232131313-0113210010113011)
- [voltstack_cluster_ar](data-sources--azure_vnet_site--reference--group-009.md#canonical-0013331330233123-3030030001210332-2301323022202203-2230301230330203-2310333012300123-0110311130202131-0121011120233100-0122113300012113)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)

<a id="canonical-0210131321111210-3311320300230211-0331301203202122-3320023223011313-2210231130123031-1012300011231023-2030203130103121-0121012113310032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0230201210300211-2031000212232203-3232021311021233-0030132131111202-3021132321123302-2232323223310201-2210100201122212-2232320323120312"></a>

## voltstack_cluster_ar.accelerated_networking.disable_spec — disable_spec / 013033131003 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-2210100210332112-1233221121133301-0120203222303210-2221022223110220-0313230312203322-2232200300001103-2301220003133113-1333232230300113)
- [voltstack_cluster_ar](data-sources--azure_vnet_site--reference--group-009.md#canonical-0013331330233123-3030030001210332-2301323022202203-2230301230330203-2310333012300123-0110311130202131-0121011120233100-0122113300012113)
- [voltstack_cluster_ar.accelerated_networking](data-sources--azure_vnet_site--reference--group-009.md#canonical-3111320123132202-0111023220022221-2133113121010030-2330203201232211-1210211131300010-3032000021100011-0310321112312100-3232023011332012)
- voltstack_cluster_ar.accelerated_networking.disable_spec

<a id="canonical-3123231312133213-3002201123221130-0000033303312221-0033331113231000-0020130123032211-1102200223321333-2100023221123330-3131300320202231"></a>

Type: `["object", {}]`. Computed.

Enable this option

<a id="canonical-0210233202212201-3131322311112221-1211233122012321-2020213233233213-2313232231133131-2321332233130120-2002302330223102-1033120311011000"></a>

## Direct properties — disable_spec / 013033131003 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2331100033123033-1300002320001033-0022002221122001-2201221033113302-0231011333221200-1302211023132311-3330322202212021-3101331022131331"></a>

## Next pages — disable_spec / 013033131003 / 4

- [voltstack_cluster_ar.accelerated_networking](data-sources--azure_vnet_site--reference--group-009.md#canonical-3111320123132202-0111023220022221-2133113121010030-2330203201232211-1210211131300010-3032000021100011-0310321112312100-3232023011332012)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)

<a id="canonical-1030213322013130-0320311022332002-3312323001321131-1133333232000132-1222313331203311-1310021101222120-2012302232131313-0113210010113011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3200000211302233-2110222303311113-0031302000013333-2210212131013230-1312332312213023-1000313210010103-3201223110132121-3202230102221221"></a>

## voltstack_cluster_ar.accelerated_networking.enable — enable / 323310230221 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-2210100210332112-1233221121133301-0120203222303210-2221022223110220-0313230312203322-2232200300001103-2301220003133113-1333232230300113)
- [voltstack_cluster_ar](data-sources--azure_vnet_site--reference--group-009.md#canonical-0013331330233123-3030030001210332-2301323022202203-2230301230330203-2310333012300123-0110311130202131-0121011120233100-0122113300012113)
- [voltstack_cluster_ar.accelerated_networking](data-sources--azure_vnet_site--reference--group-009.md#canonical-3111320123132202-0111023220022221-2133113121010030-2330203201232211-1210211131300010-3032000021100011-0310321112312100-3232023011332012)
- voltstack_cluster_ar.accelerated_networking.enable

<a id="canonical-0100332331213323-0112102333213021-0031100333312103-2112321222212330-3123210003200122-0121033032303232-3100033331011011-0320131102220111"></a>

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

<a id="canonical-2203012233233011-0332333022033313-1330012202011123-2322020330220010-1113202132322221-3332121320033130-2023033221312203-2301221131112120"></a>

## Direct properties — enable / 323310230221 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3333200312331313-0302201133021112-3033001003030020-1003223101033101-1112302101322132-3021101122010023-0303113113132221-0210110300313332"></a>

## Next pages — enable / 323310230221 / 4

- [voltstack_cluster_ar.accelerated_networking](data-sources--azure_vnet_site--reference--group-009.md#canonical-3111320123132202-0111023220022221-2133113121010030-2330203201232211-1210211131300010-3032000021100011-0310321112312100-3232023011332012)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)

<a id="canonical-3130311313313102-1223301303022202-0311100110000220-0310130012121123-2131330210121230-2023322103123321-1331013223013013-3231121003310032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0002313200313310-1221310103312313-1101312102021120-3001011232022230-2331321112332022-1320230202130221-1022300230213130-3230301031313201"></a>

## voltstack_cluster_ar.active_enhanced_firewall_policies — active_enhanced_firewall_policies / 032003112331 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-2210100210332112-1233221121133301-0120203222303210-2221022223110220-0313230312203322-2232200300001103-2301220003133113-1333232230300113)
- [voltstack_cluster_ar](data-sources--azure_vnet_site--reference--group-009.md#canonical-0013331330233123-3030030001210332-2301323022202203-2230301230330203-2310333012300123-0110311130202131-0121011120233100-0122113300012113)
- voltstack_cluster_ar.active_enhanced_firewall_policies

<a id="canonical-0211311113210310-0003233031020331-2133300232321133-0113001001110331-3100031030300231-1300113021231213-2321202100233010-0222030103311112"></a>

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

<a id="canonical-0131302231301130-3310201120300211-0011332120113121-1232211230122210-1022002103200301-3220301331000322-1012230330033332-1231122323313030"></a>

## Direct properties — active_enhanced_firewall_policies / 032003112331 / 3

- [enhanced_firewall_policies](data-sources--azure_vnet_site--reference--group-009.md#canonical-3333300310233313-3111111223001301-1030111203120131-1211322130123322-1120000112203020-0303200020230301-2113200303011101-2220200202022321): complete subsection reference.

<a id="canonical-3003110000203102-2202012310133023-1330031202321222-1012012200210213-1100112012332200-0211221230111302-2131310101223331-2010223212301003"></a>

## Next pages — active_enhanced_firewall_policies / 032003112331 / 4

- [voltstack_cluster_ar.active_enhanced_firewall_policies.enhanced_firewall_policies](data-sources--azure_vnet_site--reference--group-009.md#canonical-3333300310233313-3111111223001301-1030111203120131-1211322130123322-1120000112203020-0303200020230301-2113200303011101-2220200202022321)
- [voltstack_cluster_ar](data-sources--azure_vnet_site--reference--group-009.md#canonical-0013331330233123-3030030001210332-2301323022202203-2230301230330203-2310333012300123-0110311130202131-0121011120233100-0122113300012113)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)

<a id="canonical-3333300310233313-3111111223001301-1030111203120131-1211322130123322-1120000112203020-0303200020230301-2113200303011101-2220200202022321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1102213101020133-2020232222113103-2132320022010323-3132200030211000-3321122013123020-3113012011212100-0010202003030321-1231311211231122"></a>

## voltstack_cluster_ar.active_enhanced_firewall_policies.enhanced_firewall_policies — enhanced_firewall_policies / 131123300330 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-2210100210332112-1233221121133301-0120203222303210-2221022223110220-0313230312203322-2232200300001103-2301220003133113-1333232230300113)
- [voltstack_cluster_ar](data-sources--azure_vnet_site--reference--group-009.md#canonical-0013331330233123-3030030001210332-2301323022202203-2230301230330203-2310333012300123-0110311130202131-0121011120233100-0122113300012113)
- [voltstack_cluster_ar.active_enhanced_firewall_policies](data-sources--azure_vnet_site--reference--group-009.md#canonical-3130311313313102-1223301303022202-0311100110000220-0310130012121123-2131330210121230-2023322103123321-1331013223013013-3231121003310032)
- voltstack_cluster_ar.active_enhanced_firewall_policies.enhanced_firewall_policies

<a id="canonical-3031002213313321-3302330323322030-0010230031121220-1021320211312121-1002020021232133-3212202010312300-3023111133002111-3232003330133013"></a>

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

<a id="canonical-2321023220033202-3110230010223302-3310103032102032-1231223002112131-3210110023120000-3302011333012013-3031011120212030-3322010321020132"></a>

## Direct properties — enhanced_firewall_policies / 131123300330 / 3

<a id="canonical-1210130103300311-3032210030020331-3313123122330323-3311011131033210-1032101322022200-2232303210320330-3103210232332203-1131122032332112"></a>

<a id="canonical-3333212112313332-3023320300110132-3203013112321102-0333313023102323-0312313033012210-2231022322233221-2310321123220323-3303113211012212"></a>

## name property — enhanced_firewall_policies / 131123300330 / 4

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

<a id="canonical-1133312301322312-3103320322201033-3010202111202111-3012112211022223-2031221320301000-2021232101333301-0020222320220031-1001200231000202"></a>

<a id="canonical-2221003033023032-1301303213010111-2300230223130302-1031013111113333-3031312030120130-0300021032110212-2132220021210310-0130132322330000"></a>

## namespace property — enhanced_firewall_policies / 131123300330 / 5

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

<a id="canonical-1032130103231122-2200000022123113-2121220023120203-3320122012230001-0303120203300222-3230230020110313-0132222133021212-3321232013210131"></a>

<a id="canonical-1123021110103012-2212320123001021-1302031102110203-2132030201103100-3200301220032330-3303032331030120-0010300222210322-3332121211131231"></a>

## tenant property — enhanced_firewall_policies / 131123300330 / 6

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

<a id="canonical-1232003330023102-1002331111020303-3112331211201212-3300233221321310-3103311313323000-1331333200031233-0120300220231332-1030231121321121"></a>

## Next pages — enhanced_firewall_policies / 131123300330 / 7

- [voltstack_cluster_ar.active_enhanced_firewall_policies](data-sources--azure_vnet_site--reference--group-009.md#canonical-3130311313313102-1223301303022202-0311100110000220-0310130012121123-2131330210121230-2023322103123321-1331013223013013-3231121003310032)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)

<a id="canonical-2211230103012213-0332121122313000-3220010131233032-3312210331230021-3000311103133131-0132232121020210-2302112112112311-1320121110212001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1113220020301302-2323331220000230-2021032123302132-3021333322023123-0303332232301302-2221100213012331-2113320223212111-3031222313113103"></a>

## voltstack_cluster_ar.active_forward_proxy_policies — active_forward_proxy_policies / 032123333300 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-2210100210332112-1233221121133301-0120203222303210-2221022223110220-0313230312203322-2232200300001103-2301220003133113-1333232230300113)
- [voltstack_cluster_ar](data-sources--azure_vnet_site--reference--group-009.md#canonical-0013331330233123-3030030001210332-2301323022202203-2230301230330203-2310333012300123-0110311130202131-0121011120233100-0122113300012113)
- voltstack_cluster_ar.active_forward_proxy_policies

<a id="canonical-2122223310132231-3112031330002130-1021022232323030-1222222333133222-1103132112210212-3201201312002302-2301010111130100-3033311001100132"></a>

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

<a id="canonical-3010012320300330-2123213003331033-0332223330030300-2001121233330033-2302310303102202-3013220232221003-3320310321103231-2321313123312332"></a>

## Direct properties — active_forward_proxy_policies / 032123333300 / 3

- [forward_proxy_policies](data-sources--azure_vnet_site--reference--group-009.md#canonical-0202103102032220-0032230031223311-0100230323201210-1012121323132010-2202112030302323-0300020222231023-3332231130323030-2332123133133323): complete subsection reference.

<a id="canonical-3300023331303121-1122313302002100-0312333100331232-2011233110232111-3122301021002100-2322103222103233-3023020111101012-1011001210013213"></a>

## Next pages — active_forward_proxy_policies / 032123333300 / 4

- [voltstack_cluster_ar.active_forward_proxy_policies.forward_proxy_policies](data-sources--azure_vnet_site--reference--group-009.md#canonical-0202103102032220-0032230031223311-0100230323201210-1012121323132010-2202112030302323-0300020222231023-3332231130323030-2332123133133323)
- [voltstack_cluster_ar](data-sources--azure_vnet_site--reference--group-009.md#canonical-0013331330233123-3030030001210332-2301323022202203-2230301230330203-2310333012300123-0110311130202131-0121011120233100-0122113300012113)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)

<a id="canonical-0202103102032220-0032230031223311-0100230323201210-1012121323132010-2202112030302323-0300020222231023-3332231130323030-2332123133133323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0103113320330102-1122113333001021-2202311010323332-2010331221000312-0302010323002201-0123113023030022-2212003022212012-1133021030032300"></a>

## voltstack_cluster_ar.active_forward_proxy_policies.forward_proxy_policies — forward_proxy_policies / 133203320230 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-2210100210332112-1233221121133301-0120203222303210-2221022223110220-0313230312203322-2232200300001103-2301220003133113-1333232230300113)
- [voltstack_cluster_ar](data-sources--azure_vnet_site--reference--group-009.md#canonical-0013331330233123-3030030001210332-2301323022202203-2230301230330203-2310333012300123-0110311130202131-0121011120233100-0122113300012113)
- [voltstack_cluster_ar.active_forward_proxy_policies](data-sources--azure_vnet_site--reference--group-009.md#canonical-2211230103012213-0332121122313000-3220010131233032-3312210331230021-3000311103133131-0132232121020210-2302112112112311-1320121110212001)
- voltstack_cluster_ar.active_forward_proxy_policies.forward_proxy_policies

<a id="canonical-3112030113200331-3212121313320223-3231100002020312-0133211300003330-1202332213132022-1223002113231311-0223300221111230-3022223302313002"></a>

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

<a id="canonical-0332302201110331-3020200222222121-0311030303333011-2321013320300222-2321320222213310-1212210032023322-3130101210103033-1031333002231002"></a>

## Direct properties — forward_proxy_policies / 133203320230 / 3

<a id="canonical-3011112031111112-1103300121320020-1230010010301202-3111221322332313-3212012101032001-0331302321202130-0133221033322121-0331013022022332"></a>

<a id="canonical-3110122031113323-0331332132021112-2212203133012303-1313022100122030-2210121301120003-3322313321311300-3000301302002011-3310131333332332"></a>

## name property — forward_proxy_policies / 133203320230 / 4

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

<a id="canonical-0333323032323033-0220020213201103-2312112033113223-1202010032300113-2331030120223020-3231210221103132-1122322032010323-0322001012322012"></a>

<a id="canonical-3231210011112022-0102221032231123-0330122330010303-0111002111113223-2121122222221121-0110020200102112-0323213230320222-0103333203030303"></a>

## namespace property — forward_proxy_policies / 133203320230 / 5

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

<a id="canonical-2111020013011231-0311113032321230-2033303101102031-0313111323323302-0010130032333302-2302301200032301-0000331211233311-1200132331230220"></a>

<a id="canonical-2110011012302330-0210133112232002-1221220123223230-2230003313011032-3223201113212223-0101320230000101-0132103310333210-2301123233010333"></a>

## tenant property — forward_proxy_policies / 133203320230 / 6

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

<a id="canonical-0213013221212201-0122233102001322-1103330300223231-1112333031013020-1322201023213233-1003332322033100-3320232330332111-3013200131111020"></a>

## Next pages — forward_proxy_policies / 133203320230 / 7

- [voltstack_cluster_ar.active_forward_proxy_policies](data-sources--azure_vnet_site--reference--group-009.md#canonical-2211230103012213-0332121122313000-3220010131233032-3312210331230021-3000311103133131-0132232121020210-2302112112112311-1320121110212001)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)

<a id="canonical-2111312233203311-0010131101022301-0232021221220131-0023122023021222-3110301021302201-0002313230230303-2032003103232221-1230301111133132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3311112323120302-1133210313332030-1221223231232323-1131310000113112-3322033223233321-3031012210102302-3330230312322131-3030132303232130"></a>

## voltstack_cluster_ar.active_network_policies — active_network_policies / 022231323000 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-2210100210332112-1233221121133301-0120203222303210-2221022223110220-0313230312203322-2232200300001103-2301220003133113-1333232230300113)
- [voltstack_cluster_ar](data-sources--azure_vnet_site--reference--group-009.md#canonical-0013331330233123-3030030001210332-2301323022202203-2230301230330203-2310333012300123-0110311130202131-0121011120233100-0122113300012113)
- voltstack_cluster_ar.active_network_policies

<a id="canonical-3333333022032030-2103203332200020-1000302213111313-0203130321311222-2033221010132322-0320312301123000-2201023331102322-2301100301120013"></a>

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

<a id="canonical-3333003223222311-1312001021231021-3203203113211323-3201300303011123-1101323233330331-1321010130331312-2311310232002122-2121210213132033"></a>

## Direct properties — active_network_policies / 022231323000 / 3

- [network_policies](data-sources--azure_vnet_site--reference--group-009.md#canonical-3131300002131011-2113103130302212-1023211222012323-1010000310032202-0131233102101303-1203333121032113-3221320013201113-2001113101112321): complete subsection reference.

<a id="canonical-1323101202132132-3311213221133233-0211032022222112-0310210121213031-3301111213011221-2223220302330330-1211213032322120-0110023120133311"></a>

## Next pages — active_network_policies / 022231323000 / 4

- [voltstack_cluster_ar.active_network_policies.network_policies](data-sources--azure_vnet_site--reference--group-009.md#canonical-3131300002131011-2113103130302212-1023211222012323-1010000310032202-0131233102101303-1203333121032113-3221320013201113-2001113101112321)
- [voltstack_cluster_ar](data-sources--azure_vnet_site--reference--group-009.md#canonical-0013331330233123-3030030001210332-2301323022202203-2230301230330203-2310333012300123-0110311130202131-0121011120233100-0122113300012113)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)

<a id="canonical-3131300002131011-2113103130302212-1023211222012323-1010000310032202-0131233102101303-1203333121032113-3221320013201113-2001113101112321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0002000030213021-2210212011121322-2001303312302120-3110010010331112-1131133212310130-1233003313112111-1301312332223103-0231021301122231"></a>

## voltstack_cluster_ar.active_network_policies.network_policies — network_policies / 110023022331 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-2210100210332112-1233221121133301-0120203222303210-2221022223110220-0313230312203322-2232200300001103-2301220003133113-1333232230300113)
- [voltstack_cluster_ar](data-sources--azure_vnet_site--reference--group-009.md#canonical-0013331330233123-3030030001210332-2301323022202203-2230301230330203-2310333012300123-0110311130202131-0121011120233100-0122113300012113)
- [voltstack_cluster_ar.active_network_policies](data-sources--azure_vnet_site--reference--group-009.md#canonical-2111312233203311-0010131101022301-0232021221220131-0023122023021222-3110301021302201-0002313230230303-2032003103232221-1230301111133132)
- voltstack_cluster_ar.active_network_policies.network_policies

<a id="canonical-2320321302333233-2313313321213230-1123300200132101-2020013130202223-0333003022210312-0312211130110211-0002030021330211-3122023012000223"></a>

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

<a id="canonical-2021333002300302-0121301113133033-1332012212013322-1232120132020132-0330121111020201-2301102201110233-3311022103030011-1032001020103121"></a>

## Direct properties — network_policies / 110023022331 / 3

<a id="canonical-1010211130011132-1002310101320013-2031322222012301-2020132200302032-3330112110021132-1000133030211213-1120122010111102-3213321302002212"></a>

<a id="canonical-0100212111333321-1232321223122201-3221112033010130-1120312203122202-1202002301202110-1220233010031222-2203031232011131-3323322012312320"></a>

## name property — network_policies / 110023022331 / 4

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

<a id="canonical-2232203201332230-3312002231312213-3332212110013110-1021113222011303-2233231311301021-1100303122222131-3023111111321321-3301323110202111"></a>

<a id="canonical-1211333301122221-0122313112212202-0212302023021202-1313200210221201-1112110330032301-0001202022111203-3230201030322023-3133212000121113"></a>

## namespace property — network_policies / 110023022331 / 5

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

<a id="canonical-2223001321022233-3111310112331020-2123123102333132-0133233010303320-0320131112332210-0100320102102020-0030113312003130-2133120301300113"></a>

<a id="canonical-3233223133110320-2313221320312322-1130011111120213-2302103313320221-3312001013101121-2311010000030303-0121312310100311-0122020033122101"></a>

## tenant property — network_policies / 110023022331 / 6

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

<a id="canonical-2230130332111021-2031200003212232-3020113022322031-0300320100013103-2030113320012002-2012320111112213-0301023021032222-2001233020303230"></a>

## Next pages — network_policies / 110023022331 / 7

- [voltstack_cluster_ar.active_network_policies](data-sources--azure_vnet_site--reference--group-009.md#canonical-2111312233203311-0010131101022301-0232021221220131-0023122023021222-3110301021302201-0002313230230303-2032003103232221-1230301111133132)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)

<a id="canonical-2300233031212012-1232133331121210-0231101211002203-0312000030101310-3220033331333210-1003301023311120-1203212030031110-2000301112300331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0132121131213010-1131011200021121-2122112222020333-0300020330210200-0112132332333221-1330332023320031-2303032332023300-1102330001010033"></a>

## voltstack_cluster_ar.dc_cluster_group — dc_cluster_group / 223103233113 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-2210100210332112-1233221121133301-0120203222303210-2221022223110220-0313230312203322-2232200300001103-2301220003133113-1333232230300113)
- [voltstack_cluster_ar](data-sources--azure_vnet_site--reference--group-009.md#canonical-0013331330233123-3030030001210332-2301323022202203-2230301230330203-2310333012300123-0110311130202131-0121011120233100-0122113300012113)
- voltstack_cluster_ar.dc_cluster_group

<a id="canonical-2113320313300213-1023001103022232-1220233302033003-0322321100111200-0213311122320111-2101001130132221-3331112213321322-1023003310201132"></a>

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

<a id="canonical-1322222222111100-0201331320132233-1200012021302001-2132033200100202-0022022331321032-2011131320100233-2013312302100103-3230322320213231"></a>

## Direct properties — dc_cluster_group / 223103233113 / 3

<a id="canonical-0132203300323102-3030232030310322-3222213202013102-2022210212003222-3132103221120311-2202202330320001-3012220011211322-3322013112010113"></a>

<a id="canonical-1002131301222100-3120100301101032-2223113112313330-3101000110111201-3323023032232011-0033130201031013-2131131020022332-1131132010020321"></a>

## name property — dc_cluster_group / 223103233113 / 4

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

<a id="canonical-2311303230300003-0231220323030221-0000200303210301-1220112211222120-0021300231010201-0310300313331332-2311002212011202-0000310312020231"></a>

<a id="canonical-2201212013222100-1113230201300122-0100030103002010-3102332213333322-3132212211033213-1133331320303323-0322103033000311-0230111230032013"></a>

## namespace property — dc_cluster_group / 223103233113 / 5

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

<a id="canonical-3213231133030223-3230110322003323-3232023311133302-3302320121322223-0122102111102211-3133220023211131-1122031202300032-3022023203133033"></a>

<a id="canonical-3221122301033303-0112030230230112-0321021300022333-1113213011003202-3120020123023032-2131101333221121-0013130330310102-1102311211012331"></a>

## tenant property — dc_cluster_group / 223103233113 / 6

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

<a id="canonical-1201301111203002-1012011132302310-1001120223011010-0332323010333202-2123331023323121-3330210132200323-0210002111102220-1303031100003333"></a>

## Next pages — dc_cluster_group / 223103233113 / 7

- [voltstack_cluster_ar](data-sources--azure_vnet_site--reference--group-009.md#canonical-0013331330233123-3030030001210332-2301323022202203-2230301230330203-2310333012300123-0110311130202131-0121011120233100-0122113300012113)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)

<a id="canonical-3021110101311100-1223232022130112-1111320311302020-0202121003310301-1000221312200212-1020311233003003-1002010330131110-3322033203002112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1132032230301121-1123102311300321-1110023312331023-1221310133232002-2012230102231201-2230112102330000-2311030310332211-1230110131231322"></a>

## voltstack_cluster_ar.default_storage — default_storage / 221210013230 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-2210100210332112-1233221121133301-0120203222303210-2221022223110220-0313230312203322-2232200300001103-2301220003133113-1333232230300113)
- [voltstack_cluster_ar](data-sources--azure_vnet_site--reference--group-009.md#canonical-0013331330233123-3030030001210332-2301323022202203-2230301230330203-2310333012300123-0110311130202131-0121011120233100-0122113300012113)
- voltstack_cluster_ar.default_storage

<a id="canonical-1031330020212113-1132023113101111-0112320202032020-2012301301111131-0300223230122003-1013203130203023-0321032012333202-1223122300230300"></a>

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

<a id="canonical-0132200330332212-3120033201233011-3230123001130011-0031120301201330-3021203313202001-3200011202200030-1131100103313101-2120202312301213"></a>

## Direct properties — default_storage / 221210013230 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1011002032022300-0232100022003301-3331210111222231-3021232032323002-2223212031220033-0121323221133120-0112121310023000-3130311332120133"></a>

## Next pages — default_storage / 221210013230 / 4

- [voltstack_cluster_ar](data-sources--azure_vnet_site--reference--group-009.md#canonical-0013331330233123-3030030001210332-2301323022202203-2230301230330203-2310333012300123-0110311130202131-0121011120233100-0122113300012113)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)

<a id="canonical-1003032211122222-0321233202101030-2132030131310302-2232003312202112-1121300211312230-2101233232311113-2302311132300131-1333233122200221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3020033200322302-3201022320303123-1303333012212330-0001313202203002-2022222212332032-3232121312333111-2100011023320021-1113103102303312"></a>

## voltstack_cluster_ar.forward_proxy_allow_all — forward_proxy_allow_all / 203122110213 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-2210100210332112-1233221121133301-0120203222303210-2221022223110220-0313230312203322-2232200300001103-2301220003133113-1333232230300113)
- [voltstack_cluster_ar](data-sources--azure_vnet_site--reference--group-009.md#canonical-0013331330233123-3030030001210332-2301323022202203-2230301230330203-2310333012300123-0110311130202131-0121011120233100-0122113300012113)
- voltstack_cluster_ar.forward_proxy_allow_all

<a id="canonical-3212302220112022-2020312013112111-2333222011021121-1300133300102302-2313000000111102-3302320313233003-2123112132220120-0123303112302311"></a>

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

<a id="canonical-3231033210221201-2323100011302302-1022331333010233-1013310331113110-1112211131103131-3112313112122121-3010310123132222-3310121000302221"></a>

## Direct properties — forward_proxy_allow_all / 203122110213 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2220213313002031-1000223032120201-1032211220212033-0031010101312201-2232123111323103-1231231210103232-2030210033300311-1301101320320031"></a>

## Next pages — forward_proxy_allow_all / 203122110213 / 4

- [voltstack_cluster_ar](data-sources--azure_vnet_site--reference--group-009.md#canonical-0013331330233123-3030030001210332-2301323022202203-2230301230330203-2310333012300123-0110311130202131-0121011120233100-0122113300012113)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)

<a id="canonical-3201222313330302-0111330002310231-0113202002012313-3311012321312000-0133101333213210-2221121023210211-1211323221202222-3230133013021203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1201000333122320-2100020133221032-2321232330133212-1010202332221300-1130201123331023-3322212101310121-2333302300103003-1200302120102133"></a>

## voltstack_cluster_ar.global_network_list — global_network_list / 211013221203 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-2210100210332112-1233221121133301-0120203222303210-2221022223110220-0313230312203322-2232200300001103-2301220003133113-1333232230300113)
- [voltstack_cluster_ar](data-sources--azure_vnet_site--reference--group-009.md#canonical-0013331330233123-3030030001210332-2301323022202203-2230301230330203-2310333012300123-0110311130202131-0121011120233100-0122113300012113)
- voltstack_cluster_ar.global_network_list

<a id="canonical-0210211222102312-2113033213121021-2131100332001003-0333203330131030-1311132011310223-3211000330321120-0013333302003000-3102321101212310"></a>

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

<a id="canonical-0310122200333101-1311123323301111-1000223110132020-0231313233332333-2230330312122002-1301211201322121-0010013110330121-2300311312201322"></a>

## Direct properties — global_network_list / 211013221203 / 3

- [global_network_connections](data-sources--azure_vnet_site--reference--group-009.md#canonical-2221303310131320-3131111130223222-0031032202300003-1322100231033102-0301002110310132-3000122213311021-3123001310013101-0120310233222122): complete subsection reference.

<a id="canonical-2122200300122031-0030303001310232-2302200023011232-1231131301322233-2222310120231212-2031301021231223-0023003032212001-2031210311131331"></a>

## Next pages — global_network_list / 211013221203 / 4

- [voltstack_cluster_ar.global_network_list.global_network_connections](data-sources--azure_vnet_site--reference--group-009.md#canonical-2221303310131320-3131111130223222-0031032202300003-1322100231033102-0301002110310132-3000122213311021-3123001310013101-0120310233222122)
- [voltstack_cluster_ar](data-sources--azure_vnet_site--reference--group-009.md#canonical-0013331330233123-3030030001210332-2301323022202203-2230301230330203-2310333012300123-0110311130202131-0121011120233100-0122113300012113)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)

<a id="canonical-2221303310131320-3131111130223222-0031032202300003-1322100231033102-0301002110310132-3000122213311021-3123001310013101-0120310233222122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1210101211003310-3133301023220313-1310221310000231-3231023310231101-1200331031012201-2323303022200303-1323312301312032-2321331033003310"></a>

## voltstack_cluster_ar.global_network_list.global_network_connections — global_network_connections / 211031223200 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-2210100210332112-1233221121133301-0120203222303210-2221022223110220-0313230312203322-2232200300001103-2301220003133113-1333232230300113)
- [voltstack_cluster_ar](data-sources--azure_vnet_site--reference--group-009.md#canonical-0013331330233123-3030030001210332-2301323022202203-2230301230330203-2310333012300123-0110311130202131-0121011120233100-0122113300012113)
- [voltstack_cluster_ar.global_network_list](data-sources--azure_vnet_site--reference--group-009.md#canonical-3201222313330302-0111330002310231-0113202002012313-3311012321312000-0133101333213210-2221121023210211-1211323221202222-3230133013021203)
- voltstack_cluster_ar.global_network_list.global_network_connections

<a id="canonical-1221232031231031-0030100213210233-0322013200313022-2312010012323320-0110212230330233-1221100303312201-2010311002001202-1302320113303212"></a>

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

<a id="canonical-2210212133010131-0103131302211103-2200233212021231-1213132101021311-2322322323010201-1023233130110112-2103331121202112-1301022302021110"></a>

## Direct properties — global_network_connections / 211031223200 / 3

- [sli_to_global_dr](data-sources--azure_vnet_site--reference--group-009.md#canonical-1303122100023020-1030011112222031-1210313002101032-1110223110003122-3020033302231103-2100300100212331-1231022023020310-3232202123202111): complete subsection reference.

- [slo_to_global_dr](data-sources--azure_vnet_site--reference--group-009.md#canonical-0022100020331303-3310133102213120-1103220100103120-3203220312011213-2110020221232301-2011010012132323-3323223333322300-1310202231101212): complete subsection reference.

<a id="canonical-3101331130030212-2220302222122120-3233223333220300-0102223322012032-3020331221102111-0212303000200133-1333130101321110-2032222000100003"></a>

## Next pages — global_network_connections / 211031223200 / 4

- [voltstack_cluster_ar.global_network_list.global_network_connections.sli_to_global_dr](data-sources--azure_vnet_site--reference--group-009.md#canonical-1303122100023020-1030011112222031-1210313002101032-1110223110003122-3020033302231103-2100300100212331-1231022023020310-3232202123202111)
- [voltstack_cluster_ar.global_network_list.global_network_connections.slo_to_global_dr](data-sources--azure_vnet_site--reference--group-009.md#canonical-0022100020331303-3310133102213120-1103220100103120-3203220312011213-2110020221232301-2011010012132323-3323223333322300-1310202231101212)
- [voltstack_cluster_ar.global_network_list](data-sources--azure_vnet_site--reference--group-009.md#canonical-3201222313330302-0111330002310231-0113202002012313-3311012321312000-0133101333213210-2221121023210211-1211323221202222-3230133013021203)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)

<a id="canonical-1303122100023020-1030011112222031-1210313002101032-1110223110003122-3020033302231103-2100300100212331-1231022023020310-3232202123202111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2100013020102030-2110030012313222-1201132002221211-2311000020031223-3312020102021222-1222012211111021-1121200312110312-1032323210030221"></a>

## voltstack_cluster_ar.global_network_list.global_network_connections.sli_to_global_dr — sli_to_global_dr / 021211122303 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-2210100210332112-1233221121133301-0120203222303210-2221022223110220-0313230312203322-2232200300001103-2301220003133113-1333232230300113)
- [voltstack_cluster_ar](data-sources--azure_vnet_site--reference--group-009.md#canonical-0013331330233123-3030030001210332-2301323022202203-2230301230330203-2310333012300123-0110311130202131-0121011120233100-0122113300012113)
- [voltstack_cluster_ar.global_network_list](data-sources--azure_vnet_site--reference--group-009.md#canonical-3201222313330302-0111330002310231-0113202002012313-3311012321312000-0133101333213210-2221121023210211-1211323221202222-3230133013021203)
- [voltstack_cluster_ar.global_network_list.global_network_connections](data-sources--azure_vnet_site--reference--group-009.md#canonical-2221303310131320-3131111130223222-0031032202300003-1322100231033102-0301002110310132-3000122213311021-3123001310013101-0120310233222122)
- voltstack_cluster_ar.global_network_list.global_network_connections.sli_to_global_dr

<a id="canonical-3223001101231112-2300010022000010-1201223212021201-2221212203010101-3123012022213013-1101103203011113-3220033113002212-1111222310002121"></a>

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

<a id="canonical-0101301001310102-3310323203001230-0312122021230302-1201311221210312-3201213123111322-3113312021221022-2120322122312003-0311313031001100"></a>

## Direct properties — sli_to_global_dr / 021211122303 / 3

- [global_vn](data-sources--azure_vnet_site--reference--group-009.md#canonical-2022320031233311-2313320211002301-1322000301130012-0123302220210332-3213233302101330-0210310010003221-2003210233203200-2101112011201321): complete subsection reference.

<a id="canonical-3330012000230310-2202131331222021-1100131312220222-0333321323320120-2000100201133013-3312032021313323-3111301113002303-0110122221222130"></a>

## Next pages — sli_to_global_dr / 021211122303 / 4

- [voltstack_cluster_ar.global_network_list.global_network_connections.sli_to_global_dr.global_vn](data-sources--azure_vnet_site--reference--group-009.md#canonical-2022320031233311-2313320211002301-1322000301130012-0123302220210332-3213233302101330-0210310010003221-2003210233203200-2101112011201321)
- [voltstack_cluster_ar.global_network_list.global_network_connections](data-sources--azure_vnet_site--reference--group-009.md#canonical-2221303310131320-3131111130223222-0031032202300003-1322100231033102-0301002110310132-3000122213311021-3123001310013101-0120310233222122)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)

<a id="canonical-2022320031233311-2313320211002301-1322000301130012-0123302220210332-3213233302101330-0210310010003221-2003210233203200-2101112011201321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2000303301231200-1030313000111212-0313233033020233-3132120113213101-1030333130233301-3131122112223301-0103021203233332-1310200103221333"></a>

## voltstack_cluster_ar.global_network_list.global_network_connections.sli_to_global_dr.global_vn — global_vn / 221102311311 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-2210100210332112-1233221121133301-0120203222303210-2221022223110220-0313230312203322-2232200300001103-2301220003133113-1333232230300113)
- [voltstack_cluster_ar](data-sources--azure_vnet_site--reference--group-009.md#canonical-0013331330233123-3030030001210332-2301323022202203-2230301230330203-2310333012300123-0110311130202131-0121011120233100-0122113300012113)
- [voltstack_cluster_ar.global_network_list](data-sources--azure_vnet_site--reference--group-009.md#canonical-3201222313330302-0111330002310231-0113202002012313-3311012321312000-0133101333213210-2221121023210211-1211323221202222-3230133013021203)
- [voltstack_cluster_ar.global_network_list.global_network_connections](data-sources--azure_vnet_site--reference--group-009.md#canonical-2221303310131320-3131111130223222-0031032202300003-1322100231033102-0301002110310132-3000122213311021-3123001310013101-0120310233222122)
- [voltstack_cluster_ar.global_network_list.global_network_connections.sli_to_global_dr](data-sources--azure_vnet_site--reference--group-009.md#canonical-1303122100023020-1030011112222031-1210313002101032-1110223110003122-3020033302231103-2100300100212331-1231022023020310-3232202123202111)
- voltstack_cluster_ar.global_network_list.global_network_connections.sli_to_global_dr.global_vn

<a id="canonical-2033101221212222-0202220001032112-2013132122013231-1033121121130012-1230033322310132-0123332321332123-1030222232300212-1233230112312202"></a>

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

<a id="canonical-2103013212010312-0322303232201103-1202102302123012-3130333001301313-3003230121201130-0203020111133320-2210003230113311-3111203232330122"></a>

## Direct properties — global_vn / 221102311311 / 3

<a id="canonical-1310120221313101-1033100113230113-3232313033103302-2203131202033201-1013132212113011-1100310313311232-3320013211111213-0101022112112201"></a>

<a id="canonical-1013101013032123-0022321111331011-0131102001303212-0020133101330222-1232312221212122-3022222022002012-3113021013023201-2312303231010310"></a>

## name property — global_vn / 221102311311 / 4

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

<a id="canonical-2311310230332300-2230121300320120-2001031013232001-2311232123131220-1312023012130221-0301211121010221-0331112303022130-1302303323211311"></a>

<a id="canonical-2001023020101003-2100330333301330-0213021111203302-2021212011113113-0330103333101300-1132332013120131-1030011303002333-2311221123232131"></a>

## namespace property — global_vn / 221102311311 / 5

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

<a id="canonical-2000302112122022-0231322222212023-1112011020210202-2200012023033123-0230010003010212-0112310212112031-3333112221221131-1131200333101103"></a>

<a id="canonical-3200023022201201-1211220100101302-2321331312330320-0013110221222333-3212210203213231-3331302332231330-0300313002221111-0023222103103123"></a>

## tenant property — global_vn / 221102311311 / 6

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

<a id="canonical-2330321231323010-3212113300201010-1230200001103002-3102210332202333-3102101021100011-3113022123121223-1233333202301312-2203301000000322"></a>

## Next pages — global_vn / 221102311311 / 7

- [voltstack_cluster_ar.global_network_list.global_network_connections.sli_to_global_dr](data-sources--azure_vnet_site--reference--group-009.md#canonical-1303122100023020-1030011112222031-1210313002101032-1110223110003122-3020033302231103-2100300100212331-1231022023020310-3232202123202111)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)

<a id="canonical-0022100020331303-3310133102213120-1103220100103120-3203220312011213-2110020221232301-2011010012132323-3323223333322300-1310202231101212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1312120022012300-0310100003032013-3231201003130331-0310012310020201-3213003031032030-0001110303101011-2311330221032120-2301022122333010"></a>

## voltstack_cluster_ar.global_network_list.global_network_connections.slo_to_global_dr — slo_to_global_dr / 102011003002 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-2210100210332112-1233221121133301-0120203222303210-2221022223110220-0313230312203322-2232200300001103-2301220003133113-1333232230300113)
- [voltstack_cluster_ar](data-sources--azure_vnet_site--reference--group-009.md#canonical-0013331330233123-3030030001210332-2301323022202203-2230301230330203-2310333012300123-0110311130202131-0121011120233100-0122113300012113)
- [voltstack_cluster_ar.global_network_list](data-sources--azure_vnet_site--reference--group-009.md#canonical-3201222313330302-0111330002310231-0113202002012313-3311012321312000-0133101333213210-2221121023210211-1211323221202222-3230133013021203)
- [voltstack_cluster_ar.global_network_list.global_network_connections](data-sources--azure_vnet_site--reference--group-009.md#canonical-2221303310131320-3131111130223222-0031032202300003-1322100231033102-0301002110310132-3000122213311021-3123001310013101-0120310233222122)
- voltstack_cluster_ar.global_network_list.global_network_connections.slo_to_global_dr

<a id="canonical-1303100001012020-2333013102033100-0012011233303112-1301013121233131-3002320333110013-0210321031002201-2003213113210111-1000101122122032"></a>

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

<a id="canonical-1110303030333013-3033123301221302-3222100031202223-0102331112102323-1030031202100231-1213212131113131-2003300230000100-0030320133330002"></a>

## Direct properties — slo_to_global_dr / 102011003002 / 3

- [global_vn](data-sources--azure_vnet_site--reference--group-009.md#canonical-2220112021032022-1201003233323203-1112303030313110-3233032131122111-3012113322133011-3300233003121330-1333331030222133-0220130210110003): complete subsection reference.

<a id="canonical-3122020122333331-2202103121100010-0011102212321003-1121322013031110-2031301032102102-1030022222022011-1000200201011001-1320320130332032"></a>

## Next pages — slo_to_global_dr / 102011003002 / 4

- [voltstack_cluster_ar.global_network_list.global_network_connections.slo_to_global_dr.global_vn](data-sources--azure_vnet_site--reference--group-009.md#canonical-2220112021032022-1201003233323203-1112303030313110-3233032131122111-3012113322133011-3300233003121330-1333331030222133-0220130210110003)
- [voltstack_cluster_ar.global_network_list.global_network_connections](data-sources--azure_vnet_site--reference--group-009.md#canonical-2221303310131320-3131111130223222-0031032202300003-1322100231033102-0301002110310132-3000122213311021-3123001310013101-0120310233222122)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)

<a id="canonical-2220112021032022-1201003233323203-1112303030313110-3233032131122111-3012113322133011-3300233003121330-1333331030222133-0220130210110003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0121331222200211-0301233320320331-3310000001311233-2000013211321333-2303122121232211-1013033103201223-1300131130130111-2032102021101101"></a>

## voltstack_cluster_ar.global_network_list.global_network_connections.slo_to_global_dr.global_vn — global_vn / 300133221010 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-2210100210332112-1233221121133301-0120203222303210-2221022223110220-0313230312203322-2232200300001103-2301220003133113-1333232230300113)
- [voltstack_cluster_ar](data-sources--azure_vnet_site--reference--group-009.md#canonical-0013331330233123-3030030001210332-2301323022202203-2230301230330203-2310333012300123-0110311130202131-0121011120233100-0122113300012113)
- [voltstack_cluster_ar.global_network_list](data-sources--azure_vnet_site--reference--group-009.md#canonical-3201222313330302-0111330002310231-0113202002012313-3311012321312000-0133101333213210-2221121023210211-1211323221202222-3230133013021203)
- [voltstack_cluster_ar.global_network_list.global_network_connections](data-sources--azure_vnet_site--reference--group-009.md#canonical-2221303310131320-3131111130223222-0031032202300003-1322100231033102-0301002110310132-3000122213311021-3123001310013101-0120310233222122)
- [voltstack_cluster_ar.global_network_list.global_network_connections.slo_to_global_dr](data-sources--azure_vnet_site--reference--group-009.md#canonical-0022100020331303-3310133102213120-1103220100103120-3203220312011213-2110020221232301-2011010012132323-3323223333322300-1310202231101212)
- voltstack_cluster_ar.global_network_list.global_network_connections.slo_to_global_dr.global_vn

<a id="canonical-2000223010022030-3110321201312110-0202223121112112-0010203112020022-0032030103311100-1320310032110112-3010000202113031-3220023020031203"></a>

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

<a id="canonical-2323333130000332-1301023022013213-3102123021022301-2010222300222203-0012210111033011-2211011100221010-2013200212103110-0321301303233001"></a>

## Direct properties — global_vn / 300133221010 / 3

<a id="canonical-2231302122131131-2022211332202211-2030332022000113-2303333213021203-2012112312010320-1232221102120103-3210230031023331-0102133031100233"></a>

<a id="canonical-1222330212313221-3210013331102333-0032233003332120-0202322200320112-0013212002221023-0003133011330122-1100112232101222-0302222011100121"></a>

## name property — global_vn / 300133221010 / 4

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

<a id="canonical-1213232322000211-3011123323120330-1321131201213012-2030010212110100-3211132210230011-1012231223011333-1100011201233010-1123311120201321"></a>

<a id="canonical-3133010031310230-2020303013130010-0033220231200010-0220221232222121-1033031211313300-0213031111013311-0110311302102310-2301131033011323"></a>

## namespace property — global_vn / 300133221010 / 5

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

<a id="canonical-2103132112122222-0232223221320001-2210230322120302-0022001221132103-1033132313210313-1331011333012212-0031103123000031-2202121010122130"></a>

<a id="canonical-1312233110323010-2202221112222110-3031220302113232-3000222012102313-1210322030302222-2133112231201100-0003021003132002-3132100210002111"></a>

## tenant property — global_vn / 300133221010 / 6

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

<a id="canonical-1220311332330013-0102320012032010-2222001022003121-1203231103331111-0301130000331100-3323211321113202-0333023331221230-1220331201030003"></a>

## Next pages — global_vn / 300133221010 / 7

- [voltstack_cluster_ar.global_network_list.global_network_connections.slo_to_global_dr](data-sources--azure_vnet_site--reference--group-009.md#canonical-0022100020331303-3310133102213120-1103220100103120-3203220312011213-2110020221232301-2011010012132323-3323223333322300-1310202231101212)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)

<a id="canonical-2210130302333102-1100012131310231-2210333031222022-1233210202101120-1333223123311122-3033333032332232-3220023213223232-0033332300032221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3033010021000330-1010312121022211-1020011032231112-3132020032101011-2221223210123103-0310033002133021-3130221102332001-2121020310033023"></a>

## voltstack_cluster_ar.k8s_cluster — k8s_cluster / 200303202023 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-2210100210332112-1233221121133301-0120203222303210-2221022223110220-0313230312203322-2232200300001103-2301220003133113-1333232230300113)
- [voltstack_cluster_ar](data-sources--azure_vnet_site--reference--group-009.md#canonical-0013331330233123-3030030001210332-2301323022202203-2230301230330203-2310333012300123-0110311130202131-0121011120233100-0122113300012113)
- voltstack_cluster_ar.k8s_cluster

<a id="canonical-3331333023300122-0213121210111312-1113200330320023-1222123322200111-2000210011011231-3123232302322230-0221023123322123-1230222323102003"></a>

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

<a id="canonical-1312331001313323-3032130203031020-3031111013332131-2033230221100220-1303113221102032-2103000120202303-2312020101231201-0013012220023332"></a>

## Direct properties — k8s_cluster / 200303202023 / 3

<a id="canonical-2221001122102103-1013020230102033-1332202211213123-1010131033233002-3120133110101211-1132222313023221-1130321313310223-3113031221112123"></a>

<a id="canonical-0030003012222001-3020323123122100-1113332301222002-3221312212210301-0302031210320110-0310333313031310-1333123123102201-0000212303021121"></a>

## name property — k8s_cluster / 200303202023 / 4

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

<a id="canonical-2011212013020330-0330223333023212-3210133020222113-1232013022321130-0213212102013323-1201101313320220-3033211023321233-3231131332203113"></a>

<a id="canonical-2111022331003220-2103330310233022-2222212002113121-0011022002112033-3001220112032132-3101112331320313-0002310310003232-3030231011031211"></a>

## namespace property — k8s_cluster / 200303202023 / 5

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

<a id="canonical-0310003213002011-1133112033003212-1333132010032330-2010233022020202-1023131001133201-3120132312211032-3101010233003101-3113333102001102"></a>

<a id="canonical-2131231210020133-3202312313121332-1220033310303011-1003011013322030-2231133123231303-1231210031130010-2022323312301101-1301000132131202"></a>

## tenant property — k8s_cluster / 200303202023 / 6

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

<a id="canonical-2001012210211323-0001333311133321-2231130321120133-3113323121322023-3212203323301113-0220312110231301-3232312220310013-0310130102120313"></a>

## Next pages — k8s_cluster / 200303202023 / 7

- [voltstack_cluster_ar](data-sources--azure_vnet_site--reference--group-009.md#canonical-0013331330233123-3030030001210332-2301323022202203-2230301230330203-2310333012300123-0110311130202131-0121011120233100-0122113300012113)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)

<a id="canonical-2232123200303133-2110101003120023-3302131323101001-2013020002133320-3010303112002323-1100012011010232-2032102310220301-0000313000211123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2031121110330221-1020112032130103-3223330313020223-2022120233011311-2013100230303121-3320131323232120-2111021332133112-3203331300010122"></a>

## voltstack_cluster_ar.no_dc_cluster_group — no_dc_cluster_group / 022020333203 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-2210100210332112-1233221121133301-0120203222303210-2221022223110220-0313230312203322-2232200300001103-2301220003133113-1333232230300113)
- [voltstack_cluster_ar](data-sources--azure_vnet_site--reference--group-009.md#canonical-0013331330233123-3030030001210332-2301323022202203-2230301230330203-2310333012300123-0110311130202131-0121011120233100-0122113300012113)
- voltstack_cluster_ar.no_dc_cluster_group

<a id="canonical-0331220201110113-1210001232203023-3210211321212231-3211233000233102-0020232221220032-3011333232002213-2001323120321230-2023311121103213"></a>

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

<a id="canonical-1203210112300211-2331331313100002-3002132231112002-0200131023331320-2021013230130030-2123123200032330-2310332102013131-0133233333203201"></a>

## Direct properties — no_dc_cluster_group / 022020333203 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3131113203133022-1331203322012102-1202322112202320-3311033030013112-0031222220320011-2330203332100221-1130130212100311-0020303120203233"></a>

## Next pages — no_dc_cluster_group / 022020333203 / 4

- [voltstack_cluster_ar](data-sources--azure_vnet_site--reference--group-009.md#canonical-0013331330233123-3030030001210332-2301323022202203-2230301230330203-2310333012300123-0110311130202131-0121011120233100-0122113300012113)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)

<a id="canonical-1220213212302011-1111032113103010-3303011132100112-2003203300112313-3313001323110130-1232112200112330-2002133212232031-2110330120213200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1203301213233110-2012021232200323-0001301232023200-0232100301111330-2233021221303310-0131210200002113-2300321023033212-3113332033123113"></a>

## voltstack_cluster_ar.no_forward_proxy — no_forward_proxy / 331100312300 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-2210100210332112-1233221121133301-0120203222303210-2221022223110220-0313230312203322-2232200300001103-2301220003133113-1333232230300113)
- [voltstack_cluster_ar](data-sources--azure_vnet_site--reference--group-009.md#canonical-0013331330233123-3030030001210332-2301323022202203-2230301230330203-2310333012300123-0110311130202131-0121011120233100-0122113300012113)
- voltstack_cluster_ar.no_forward_proxy

<a id="canonical-3230021333031210-3032213100010021-3023223210322031-1323202131111110-3022230022320101-1133203010231122-3032122222303123-0332300223313120"></a>

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

<a id="canonical-0131001310300301-3312320223323032-1123223130000111-2002012321300132-1121001032101203-0100212203023021-2221030111110301-2030002201311013"></a>

## Direct properties — no_forward_proxy / 331100312300 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2110022312301300-1001233203302123-3111020031231311-0212002302110321-0012032130123133-2302330020303222-0201001200330210-0131022030300232"></a>

## Next pages — no_forward_proxy / 331100312300 / 4

- [voltstack_cluster_ar](data-sources--azure_vnet_site--reference--group-009.md#canonical-0013331330233123-3030030001210332-2301323022202203-2230301230330203-2310333012300123-0110311130202131-0121011120233100-0122113300012113)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)

<a id="canonical-3103233133102122-3202232301232023-3111301313211310-0001012121230013-3131122323222331-2212212130231230-2323021213233202-2113022210021320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1331300321132130-1120322103021130-3132310102200230-1213122302130122-0123131013000133-2132021301211210-1132002120231300-0233013111013322"></a>

## voltstack_cluster_ar.no_global_network — no_global_network / 123201102120 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-2210100210332112-1233221121133301-0120203222303210-2221022223110220-0313230312203322-2232200300001103-2301220003133113-1333232230300113)
- [voltstack_cluster_ar](data-sources--azure_vnet_site--reference--group-009.md#canonical-0013331330233123-3030030001210332-2301323022202203-2230301230330203-2310333012300123-0110311130202131-0121011120233100-0122113300012113)
- voltstack_cluster_ar.no_global_network

<a id="canonical-3301201200123210-1213132323211001-0122313333012300-1221320122313310-3203100212310011-1132110203312003-0333130112012011-0023001033301132"></a>

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

<a id="canonical-2303300310213132-3322301111133030-2110023002213322-1122323030310111-3312110213131231-1301300223103122-2130101211223031-2220100012013132"></a>

## Direct properties — no_global_network / 123201102120 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2301331113313310-1220130121232322-0133230012112332-1302221022103331-1012003213000200-3232313000310112-3102322333332003-3112000310311100"></a>

## Next pages — no_global_network / 123201102120 / 4

- [voltstack_cluster_ar](data-sources--azure_vnet_site--reference--group-009.md#canonical-0013331330233123-3030030001210332-2301323022202203-2230301230330203-2310333012300123-0110311130202131-0121011120233100-0122113300012113)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)

<a id="canonical-0300213333023101-1320213301033231-2123212001333312-3313222011323111-1312102012202032-1233012021032330-2320122031112023-0113213113303300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1212120322002102-2120200322213032-2003230210323012-2023000201333333-3230311303310321-1210322101111023-2301330010220213-0013313313102310"></a>

## voltstack_cluster_ar.no_k8s_cluster — no_k8s_cluster / 101233111021 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-2210100210332112-1233221121133301-0120203222303210-2221022223110220-0313230312203322-2232200300001103-2301220003133113-1333232230300113)
- [voltstack_cluster_ar](data-sources--azure_vnet_site--reference--group-009.md#canonical-0013331330233123-3030030001210332-2301323022202203-2230301230330203-2310333012300123-0110311130202131-0121011120233100-0122113300012113)
- voltstack_cluster_ar.no_k8s_cluster

<a id="canonical-3030132210121000-2103033000311023-0113022220310310-2321223333220332-2300202110031030-2013231020122032-0001202030300201-0011223133303231"></a>

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

<a id="canonical-0022321001130022-3213210120011231-1321111323220202-2021202131320110-0033303201200013-1213300103002033-0030323133102310-3131201230013301"></a>

## Direct properties — no_k8s_cluster / 101233111021 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1333003201332102-1210032033123233-2200201221130132-2312031100110022-2011003312231311-1121132212001301-3233031120103103-1031221201313223"></a>

## Next pages — no_k8s_cluster / 101233111021 / 4

- [voltstack_cluster_ar](data-sources--azure_vnet_site--reference--group-009.md#canonical-0013331330233123-3030030001210332-2301323022202203-2230301230330203-2310333012300123-0110311130202131-0121011120233100-0122113300012113)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)

<a id="canonical-3131333313020120-3030311131123210-1130312011123300-0022213231123232-0010320322000333-2313000010320333-3302131010332301-2301310021313300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1320302321030222-0112322132201303-2000220233322023-2331101132102000-0112111233121120-2311312010202031-0101300131302223-3132030201220110"></a>

## voltstack_cluster_ar.no_network_policy — no_network_policy / 203022121031 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-2210100210332112-1233221121133301-0120203222303210-2221022223110220-0313230312203322-2232200300001103-2301220003133113-1333232230300113)
- [voltstack_cluster_ar](data-sources--azure_vnet_site--reference--group-009.md#canonical-0013331330233123-3030030001210332-2301323022202203-2230301230330203-2310333012300123-0110311130202131-0121011120233100-0122113300012113)
- voltstack_cluster_ar.no_network_policy

<a id="canonical-3112001212101323-0032213013022221-2330133120022232-1313220312023111-2330213021233213-2313103201003212-0132000200132022-0001030031132010"></a>

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

<a id="canonical-0233021223030330-0112231111130021-2101233231332110-2303112210103320-0110202222202321-0230310331230110-3213311322123212-3031322030120111"></a>

## Direct properties — no_network_policy / 203022121031 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2230132132120302-2311202200031331-1332102000323010-3111110033131022-1311311202323132-3230101230002310-0021032102221211-0330203000133111"></a>

## Next pages — no_network_policy / 203022121031 / 4

- [voltstack_cluster_ar](data-sources--azure_vnet_site--reference--group-009.md#canonical-0013331330233123-3030030001210332-2301323022202203-2230301230330203-2310333012300123-0110311130202131-0121011120233100-0122113300012113)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)

<a id="canonical-0101302123000232-0121202321203200-2032301212222100-1101320230302321-3311111221102030-2100313030212333-1202023333300000-2200220000013211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3211121110131122-3313010300303222-0131313110010003-1313203310223210-0133302203000003-3111112331123202-2300312100022311-3120232313223233"></a>

## voltstack_cluster_ar.no_outside_static_routes — no_outside_static_routes / 322033322110 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-2210100210332112-1233221121133301-0120203222303210-2221022223110220-0313230312203322-2232200300001103-2301220003133113-1333232230300113)
- [voltstack_cluster_ar](data-sources--azure_vnet_site--reference--group-009.md#canonical-0013331330233123-3030030001210332-2301323022202203-2230301230330203-2310333012300123-0110311130202131-0121011120233100-0122113300012113)
- voltstack_cluster_ar.no_outside_static_routes

<a id="canonical-3200300212303232-3033211232121033-3231030101021202-2011101133302311-0310221013200210-2113003311011312-3033030202113202-1310110033030101"></a>

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

<a id="canonical-2130222013022223-1113001332001222-1322020210023022-0103301111330233-0321213123100311-1303102002121202-1303310031133103-0323132212010212"></a>

## Direct properties — no_outside_static_routes / 322033322110 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2022112102310132-1312023032333021-2130301001031010-2113002001132320-2132222100232023-1020010232211031-1123233113302103-2313210322121321"></a>

## Next pages — no_outside_static_routes / 322033322110 / 4

- [voltstack_cluster_ar](data-sources--azure_vnet_site--reference--group-009.md#canonical-0013331330233123-3030030001210332-2301323022202203-2230301230330203-2310333012300123-0110311130202131-0121011120233100-0122113300012113)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)

<a id="canonical-2001203102000121-3022320233001111-0223132200203002-3110311030033021-1333212300230100-1303202000011321-1201101031310132-3003123330033331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3312130233102203-2222311231220110-2222222323202013-2322332333332133-1231130331220102-3122110321210022-0200033202311231-1322111232320112"></a>

## voltstack_cluster_ar.node — node / 332112001032 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-2210100210332112-1233221121133301-0120203222303210-2221022223110220-0313230312203322-2232200300001103-2301220003133113-1333232230300113)
- [voltstack_cluster_ar](data-sources--azure_vnet_site--reference--group-009.md#canonical-0013331330233123-3030030001210332-2301323022202203-2230301230330203-2310333012300123-0110311130202131-0121011120233100-0122113300012113)
- voltstack_cluster_ar.node

<a id="canonical-1101120320233222-2211312302013203-3210221112221202-0130131110001210-1200231320212223-3133323200111130-2021213203023033-1103231013233001"></a>

Type: `"single"`. Computed.

Parameters for creating Single interface Node for Alternate Region.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1331112213013003-3112030011123111-2313313300023011-0012232222020011-0312331223301222-3133230132333230-3132300121033123-1023233021103103"></a>

## Direct properties — node / 332112001032 / 3

<a id="canonical-3122003022220301-3313320113103012-0312223100213000-2333210133320300-3122112221321200-1030220321232320-2210110111003221-2133310220202013"></a>

<a id="canonical-3210303122300202-3013312032021322-2333203323131231-3210133010213232-1013102100131033-3233100310022020-1000122010123330-1132032332112130"></a>

## fault_domain property — node / 332112001032 / 4

Type: `"number"`. Computed.

Namuber of fault domains to be used while creating the availability set.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 3,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "3"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "3"
  }
}
```

- [local_subnet](data-sources--azure_vnet_site--reference--group-009.md#canonical-3310310310032121-3120002112221320-1200132101132122-2313002323032211-1213111023230133-1232200321322233-1323111103311303-1312321021123022): complete subsection reference.

<a id="canonical-2230031312023221-0312023221301101-2030113001013122-1030202023221100-1000313322231023-2121030001110132-1030120333322102-2223103312131331"></a>

<a id="canonical-3213113223202003-3220102230320001-1011201001333221-3220002113320201-3312101032202210-1231320120113101-2233201032131323-2211022013131332"></a>

## node_number property — node / 332112001032 / 5

Type: `"number"`. Computed.

Number of main nodes to create, either 1 or 3.

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
    "ves.io.schema.rules.uint32.in": "[1,3]"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.in": "[1,3]"
  }
}
```

<a id="canonical-0000200312200013-0302233231013103-2233003033022330-3302033013321232-0231113032212312-0202133130313200-3311113032201033-1202233212110230"></a>

<a id="canonical-3002320021022101-3021201132320032-1031012202120110-0230130212113220-1013020331100322-3310012222033211-3130032323323002-2231232031110201"></a>

## update_domain property — node / 332112001032 / 6

Type: `"number"`. Computed.

Namuber of update domains to be used while creating the availability set.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 20,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "20"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "20"
  }
}
```

<a id="canonical-0312021112332122-1100130333032000-2100213033023332-1230003000133333-3133201011121223-3103220233122202-3302121321021220-3200322130122221"></a>

## Next pages — node / 332112001032 / 7

- [voltstack_cluster_ar.node.local_subnet](data-sources--azure_vnet_site--reference--group-009.md#canonical-3310310310032121-3120002112221320-1200132101132122-2313002323032211-1213111023230133-1232200321322233-1323111103311303-1312321021123022)
- [voltstack_cluster_ar](data-sources--azure_vnet_site--reference--group-009.md#canonical-0013331330233123-3030030001210332-2301323022202203-2230301230330203-2310333012300123-0110311130202131-0121011120233100-0122113300012113)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)

<a id="canonical-3310310310032121-3120002112221320-1200132101132122-2313002323032211-1213111023230133-1232200321322233-1323111103311303-1312321021123022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->
