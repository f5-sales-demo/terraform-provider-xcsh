---
page_title: "xcsh_cloud_connect reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_cloud_connect reference."
---

# xcsh_cloud_connect reference

<a id="canonical-0222032211111313-0003010022101312-0310023330122333-0201122213032311-0122213101110201-2002020111110210-1221320313132123-3300313332231200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0313331031110100-0300032103310002-2201101132212223-2230130133330022-2221022021200221-3033132301110310-0210111120202101-1303213312021022"></a>

## Property reference — Property reference / 031022023211 / 2

Breadcrumbs:

- [xcsh_cloud_connect](../data-sources/cloud_connect.md#canonical-2011032310310033-3311323122002031-2213302012003111-3033233110012013-2021203231210000-3221202000202111-0001331133100330-3301210132312000)
- Property reference

<a id="canonical-0232002111120012-0230030113030320-0110132220230231-1220331213131131-3213023312101112-1220301013323203-0203113203103203-0000211130030303"></a>

## Direct properties — Property reference / 031022023211 / 3

<a id="canonical-2131301123220101-1023320102123110-0201332110133111-3030003302033333-2202120113101013-0323213021000330-2023133110023323-3233110202013033"></a>

<a id="canonical-2212010033221113-0202321303311332-0210201213031221-2021101310310202-0203033223120120-1303020003103212-3120000020300222-0231332221000020"></a>

## annotations property — Property reference / 031022023211 / 4

Type: `["map", "string"]`. Computed.

Annotations applied to this resource.

Upstream description:

Annotations is an unstructured key-value map stored with a resource that may be set by external
tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when
modifying objects.

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
    "ves.io.schema.rules.map.keys.string.max_len": "64",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.values.string.max_len": "1024",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "64",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.values.string.max_len": "1024",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  }
}
```

- [aws_provider](data-sources--cloud_connect--reference--group-001.md#canonical-3322210313330233-3203313001001101-1033130101220000-1313031021321122-3011231212132102-2121333313020110-1123112301023233-1202303010121200): complete subsection reference.

- [azure_vnet_site](data-sources--cloud_connect--reference--group-001.md#canonical-1310213230010202-0010122300232220-3030003101222331-3033212232021303-0201232322203122-3303121123112230-3233222203332331-2222210203131132): complete subsection reference.

<a id="canonical-0322301302023230-1000010020002332-2332311311222230-0313130102233333-0320103301300100-2131212102231320-0210121023031303-0002310332000211"></a>

<a id="canonical-3322220210231203-0202233112113323-0321023010112001-1300331122312133-3223032223133303-3132311232131123-0120010212203203-0313110333233003"></a>

## description property — Property reference / 031022023211 / 5

Type: `"string"`. Computed.

Description of the CloudConnect.

Upstream description:

Human readable description for the object.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 1200,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 1200
    },
    "category": "discovery",
    "characterSet": {
      "description": "Free text with UTF-8 support"
    },
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 1200,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 0
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "1200"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "1200"
  }
}
```

<a id="canonical-2022011201212303-2312301100103312-3223313130221101-2012303113223112-2030102222002130-3220303030010110-2123020100021330-3102033101123013"></a>

<a id="canonical-2313211213223131-0301031012120203-1012330132002020-1222023003312120-3000133003001103-3232321303332021-3221222232020021-1021022221132002"></a>

## ID property — Property reference / 031022023211 / 6

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-0033320133021013-0322003311213233-3001022231022123-2020320201310020-2123133303111123-2122312220102333-0230213230013310-3032201012001132"></a>

<a id="canonical-0311330023020222-0323123222032323-0030320301103323-1012323313101300-2313011113102112-1113210021320012-1212233231212320-1112130331302131"></a>

## labels property — Property reference / 031022023211 / 7

Type: `["map", "string"]`. Computed.

Labels applied to this resource.

Upstream description:

Map of string keys and values that can be used to organize and categorize (scope and select) objects
as chosen by the user. Values specified here will be used by selector expression.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1333313323220130-2232002220300220-2200122030303222-2321130331320010-2013210033003221-2201113110203323-0223330100333022-2130320012013033"></a>

<a id="canonical-1220223233332023-1120233203321320-3120133301003031-2001320121122130-1312101211100203-3212201023113122-0200112310013211-3131312000022221"></a>

## name property — Property reference / 031022023211 / 8

Type: `"string"`. Required.

Name of the CloudConnect.

Upstream description:

This is the name of configuration object. It has to be unique within the namespace. It can only be
specified during create API and cannot be changed during replace API. The value of name has to
follow DNS-1035 format.

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
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  }
}
```

<a id="canonical-2230020311230332-0133222222111101-2332110002313012-2110113133222103-0310312313232002-2201132103331100-1332210133100123-3213221320232223"></a>

<a id="canonical-1220200200302332-0101310000302103-0000010221013312-1123112113110022-0233132313300230-1113101210222123-3220202132030300-3110011321121303"></a>

## namespace property — Property reference / 031022023211 / 9

Type: `"string"`. Required.

Namespace where the CloudConnect exists.

Upstream description:

This defines the workspace within which each the configuration object is to be created. Must be a
DNS\_LABEL format. For a namespace object itself, namespace value will be ""

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

- [segment](data-sources--cloud_connect--reference--group-001.md#canonical-0231311031231003-1321013301003130-3313203023221333-0013202013320101-2320200121101002-2203011312223323-1203000323320122-1200110112102001): complete subsection reference.

<a id="canonical-0232103022330330-3021112332023331-1020221112331200-3123101033101200-2311002303323203-3230302033202122-3112113132203210-0120231122220110"></a>

## All schema paths — Property reference / 031022023211 / 10

Each exact path has one authoritative reference destination. Collection element indices are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](data-sources--cloud_connect--reference--group-001.md#canonical-2131301123220101-1023320102123110-0201332110133111-3030003302033333-2202120113101013-0323213021000330-2023133110023323-3233110202013033) |
| `aws_provider` | [aws_provider](data-sources--cloud_connect--reference--group-001.md#canonical-0323202020332112-0003310021012122-0333333120010311-0010312113123221-2210103131221221-1102330110210113-2101030230221010-1133032313332022) |
| `aws_provider.aws_tgw_site` | [aws_provider.aws_tgw_site](data-sources--cloud_connect--reference--group-001.md#canonical-0323101110301210-2203102021231111-0132323321120133-0010211122013322-2002211110011322-2213313201010232-3023023313322231-3113320131200221) |
| `aws_provider.aws_tgw_site.cred` | [aws_provider.aws_tgw_site.cred](data-sources--cloud_connect--reference--group-001.md#canonical-3220202231021112-2232322331331033-0313231233130001-3023102010030032-1331033133122231-1313220330223112-1310100020020013-2123333312111123) |
| `aws_provider.aws_tgw_site.cred.name` | [aws_provider.aws_tgw_site.cred.name](data-sources--cloud_connect--reference--group-001.md#canonical-1220122320022201-0102121123003021-1313312212333220-0032322022223312-1011123221322200-1221033230322311-3101132001203111-2223123123023003) |
| `aws_provider.aws_tgw_site.cred.namespace` | [aws_provider.aws_tgw_site.cred.namespace](data-sources--cloud_connect--reference--group-001.md#canonical-2022011120013023-0333100101231203-3300330333320103-0100213101210311-1021331011220020-2321230123131210-1031233321202000-3323000032330000) |
| `aws_provider.aws_tgw_site.cred.tenant` | [aws_provider.aws_tgw_site.cred.tenant](data-sources--cloud_connect--reference--group-001.md#canonical-3023221212110203-0330210221321101-2213100310112030-2100123133321120-0212230022321312-0011312302322220-2013103120321202-3120312221222210) |
| `aws_provider.aws_tgw_site.site` | [aws_provider.aws_tgw_site.site](data-sources--cloud_connect--reference--group-001.md#canonical-2031121000320021-1003023101300120-3203010221230023-2021213133000030-0212031331330330-3302233112123130-3121110010232020-3323020022311132) |
| `aws_provider.aws_tgw_site.site.name` | [aws_provider.aws_tgw_site.site.name](data-sources--cloud_connect--reference--group-001.md#canonical-0201223212302023-0311203230321020-1003203120221011-2202121213011102-3310333321131030-2310012022013103-3322010321130330-1020022212132321) |
| `aws_provider.aws_tgw_site.site.namespace` | [aws_provider.aws_tgw_site.site.namespace](data-sources--cloud_connect--reference--group-001.md#canonical-1313310112133300-2032311231000203-1102221330031221-2332330301033213-1131101302210113-2310131230131010-0320331330201133-0023220030132230) |
| `aws_provider.aws_tgw_site.site.tenant` | [aws_provider.aws_tgw_site.site.tenant](data-sources--cloud_connect--reference--group-001.md#canonical-0020202012230021-0322102200322103-3031320102231003-2003231301200003-1323300113111012-3211233300331030-1030011213312321-2211111020103312) |
| `aws_provider.aws_tgw_site.vpc_attachments` | [aws_provider.aws_tgw_site.vpc_attachments](data-sources--cloud_connect--reference--group-001.md#canonical-1310322300123113-1010220023211111-1111023033321032-3202323020313312-2121003311010202-2220130203213130-2321100213002313-0122333010010301) |
| `aws_provider.aws_tgw_site.vpc_attachments.vpc_list` | [aws_provider.aws_tgw_site.vpc_attachments.vpc_list](data-sources--cloud_connect--reference--group-001.md#canonical-3003132211300032-3100010103222301-3330301023321103-1333103221331222-3100313331211232-0003003003112120-3122000003121322-2030311201201002) |
| `aws_provider.aws_tgw_site.vpc_attachments.vpc_list.custom_routing` | [aws_provider.aws_tgw_site.vpc_attachments.vpc_list.custom_routing](data-sources--cloud_connect--reference--group-001.md#canonical-0100223003113001-1323032321321022-1222002133203021-3203331203320023-2121133303110333-0002312002210221-0111113313023030-2223133103032322) |
| `aws_provider.aws_tgw_site.vpc_attachments.vpc_list.custom_routing.route_tables` | [aws_provider.aws_tgw_site.vpc_attachments.vpc_list.custom_routing.route_tables](data-sources--cloud_connect--reference--group-001.md#canonical-3010211203312021-1302213213133231-0011022122211320-3321231223222101-1332202221223232-2232222102210133-2300123302012213-2022031000013111) |
| `aws_provider.aws_tgw_site.vpc_attachments.vpc_list.custom_routing.route_tables.route_table_id` | [aws_provider.aws_tgw_site.vpc_attachments.vpc_list.custom_routing.route_tables.route_table_id](data-sources--cloud_connect--reference--group-001.md#canonical-3102330220020133-3310013221330302-0113231231011300-0213013032012000-3300201321130232-1133030112020121-2203030001313322-1122230013003232) |
| `aws_provider.aws_tgw_site.vpc_attachments.vpc_list.custom_routing.route_tables.static_routes` | [aws_provider.aws_tgw_site.vpc_attachments.vpc_list.custom_routing.route_tables.static_routes](data-sources--cloud_connect--reference--group-001.md#canonical-1031211001233031-1330330331211230-3003320212010220-2302322323130103-0100203222023230-1302220122132130-3000101303302032-2130112210300210) |
| `aws_provider.aws_tgw_site.vpc_attachments.vpc_list.default_route` | [aws_provider.aws_tgw_site.vpc_attachments.vpc_list.default_route](data-sources--cloud_connect--reference--group-001.md#canonical-3212213132131302-2112023130010330-2030203133211133-2012100213322031-1000303202032333-1201011223203121-2133301002010231-0130133102103020) |
| `aws_provider.aws_tgw_site.vpc_attachments.vpc_list.default_route.all_route_tables` | [aws_provider.aws_tgw_site.vpc_attachments.vpc_list.default_route.all_route_tables](data-sources--cloud_connect--reference--group-001.md#canonical-3131010131022100-0233201030003013-1200231122103303-0301322200022312-2023112310022132-2132001002312221-1223112333212031-1311203110123200) |
| `aws_provider.aws_tgw_site.vpc_attachments.vpc_list.default_route.selective_route_tables` | [aws_provider.aws_tgw_site.vpc_attachments.vpc_list.default_route.selective_route_tables](data-sources--cloud_connect--reference--group-001.md#canonical-2330320210133210-2322002012313033-1000333010011302-3102320333120102-0231002310122120-3330312213300012-1312221111232231-0321211331120332) |
| `aws_provider.aws_tgw_site.vpc_attachments.vpc_list.default_route.selective_route_tables.route_table_id` | [aws_provider.aws_tgw_site.vpc_attachments.vpc_list.default_route.selective_route_tables.route_table_id](data-sources--cloud_connect--reference--group-001.md#canonical-2122110230303202-2233320102123223-2303130111221002-0121232312221302-2331123322213111-0133212202300333-3100331331222203-0202130300231321) |
| `aws_provider.aws_tgw_site.vpc_attachments.vpc_list.labels` | [aws_provider.aws_tgw_site.vpc_attachments.vpc_list.labels](data-sources--cloud_connect--reference--group-001.md#canonical-1301323010121110-1131312032002230-3012021122010113-1232320303122212-0312233023020230-0013310302011112-2131111230222221-1211232220113012) |
| `aws_provider.aws_tgw_site.vpc_attachments.vpc_list.manual_routing` | [aws_provider.aws_tgw_site.vpc_attachments.vpc_list.manual_routing](data-sources--cloud_connect--reference--group-001.md#canonical-2211322220313012-2023130030001222-3230220303023301-0020223230120100-1212010330010012-0311131310333212-2210123000121200-3303003330300323) |
| `aws_provider.aws_tgw_site.vpc_attachments.vpc_list.vpc_id` | [aws_provider.aws_tgw_site.vpc_attachments.vpc_list.vpc_id](data-sources--cloud_connect--reference--group-001.md#canonical-0132121310112132-0023131020323103-0133321003102330-2001111300003332-3010032022132221-0233232202033113-0202131123031032-2011321332132020) |
| `azure_vnet_site` | [azure_vnet_site](data-sources--cloud_connect--reference--group-001.md#canonical-2303331131320220-1121203202332230-2002103033303200-1330122233203313-2233002333201203-2210330300013330-3112012031201022-0212312313033300) |
| `azure_vnet_site.site` | [azure_vnet_site.site](data-sources--cloud_connect--reference--group-001.md#canonical-2030211203020302-2303220330131231-2132310220022100-0110121302033223-0012301010131310-3300300132231012-1000212321020331-3302233200123022) |
| `azure_vnet_site.site.name` | [azure_vnet_site.site.name](data-sources--cloud_connect--reference--group-001.md#canonical-1102020303211010-0011102030131312-3313213020002113-3201331333223201-2233201212100101-0230023220232002-2222020103111320-3312330220011012) |
| `azure_vnet_site.site.namespace` | [azure_vnet_site.site.namespace](data-sources--cloud_connect--reference--group-001.md#canonical-1123101002330133-0100211213003113-3100331210303222-3300101131211310-0132122212123130-3133202321121230-1012030033333132-1033203010033331) |
| `azure_vnet_site.site.tenant` | [azure_vnet_site.site.tenant](data-sources--cloud_connect--reference--group-001.md#canonical-0230203120301011-1310331130030032-1311123321321232-0033113331231222-1021023112100110-2300102011323003-3320333031232321-3232020302000021) |
| `azure_vnet_site.vnet_attachments` | [azure_vnet_site.vnet_attachments](data-sources--cloud_connect--reference--group-001.md#canonical-2200223322013231-3133121010213111-1011022302132032-2113023221120221-2132002221103103-2223222302332033-0221332000013030-3101203113330001) |
| `azure_vnet_site.vnet_attachments.vnet_list` | [azure_vnet_site.vnet_attachments.vnet_list](data-sources--cloud_connect--reference--group-001.md#canonical-3003310002001023-3031001112033032-3023222232312120-2010022113001233-0021331032103003-1002330002103220-3311131233333331-3101112031331012) |
| `azure_vnet_site.vnet_attachments.vnet_list.custom_routing` | [azure_vnet_site.vnet_attachments.vnet_list.custom_routing](data-sources--cloud_connect--reference--group-001.md#canonical-1202221332023212-0021111010003202-0332121131112101-2133121213122232-2020012110303132-3303200101100002-2300330101113311-0211300303212322) |
| `azure_vnet_site.vnet_attachments.vnet_list.custom_routing.route_tables` | [azure_vnet_site.vnet_attachments.vnet_list.custom_routing.route_tables](data-sources--cloud_connect--reference--group-001.md#canonical-0322202123202203-3320322122311122-3011131023001003-2011122010211101-1301001323000311-3300320003312003-0122122302223331-0132320320200013) |
| `azure_vnet_site.vnet_attachments.vnet_list.custom_routing.route_tables.route_table_id` | [azure_vnet_site.vnet_attachments.vnet_list.custom_routing.route_tables.route_table_id](data-sources--cloud_connect--reference--group-001.md#canonical-3223020110011033-0133232003231230-0031311132200013-0323000031022320-2332012230133221-2112203031301101-3201013120200201-2010301211032020) |
| `azure_vnet_site.vnet_attachments.vnet_list.custom_routing.route_tables.static_routes` | [azure_vnet_site.vnet_attachments.vnet_list.custom_routing.route_tables.static_routes](data-sources--cloud_connect--reference--group-001.md#canonical-2031120031312132-0033232203303033-1303013331322211-1302302001213013-1210130110212201-3011232233112323-2201010021332101-0212231120332221) |
| `azure_vnet_site.vnet_attachments.vnet_list.default_route` | [azure_vnet_site.vnet_attachments.vnet_list.default_route](data-sources--cloud_connect--reference--group-001.md#canonical-2230222032302001-3202100312302000-3030232332233133-0302122230331133-2311023110302213-2302031112200203-1122020301131101-2320100031310112) |
| `azure_vnet_site.vnet_attachments.vnet_list.default_route.all_route_tables` | [azure_vnet_site.vnet_attachments.vnet_list.default_route.all_route_tables](data-sources--cloud_connect--reference--group-001.md#canonical-0231123112110200-3020001130310300-0110120210011323-3000312013320332-0212223012023303-3033220331321010-1001002231230332-0310122230300221) |
| `azure_vnet_site.vnet_attachments.vnet_list.default_route.selective_route_tables` | [azure_vnet_site.vnet_attachments.vnet_list.default_route.selective_route_tables](data-sources--cloud_connect--reference--group-001.md#canonical-1213112003003022-0113113333201230-3313032110303120-2013110223231233-3210020321220013-0003101222313030-3312122022322302-1322012310221220) |
| `azure_vnet_site.vnet_attachments.vnet_list.default_route.selective_route_tables.route_table_id` | [azure_vnet_site.vnet_attachments.vnet_list.default_route.selective_route_tables.route_table_id](data-sources--cloud_connect--reference--group-001.md#canonical-1333111113123110-0013232011000233-2211001221022032-2031103132010131-3112200310121113-2201321201122231-3020102213311311-3031332222312213) |
| `azure_vnet_site.vnet_attachments.vnet_list.labels` | [azure_vnet_site.vnet_attachments.vnet_list.labels](data-sources--cloud_connect--reference--group-001.md#canonical-1023212103223232-0131121001222212-0312023302021112-0301132001221032-0210310233210332-3202210301133332-0022021110011303-1100133130311011) |
| `azure_vnet_site.vnet_attachments.vnet_list.manual_routing` | [azure_vnet_site.vnet_attachments.vnet_list.manual_routing](data-sources--cloud_connect--reference--group-001.md#canonical-1310113022310312-1130103101120312-0303132100020203-0213110130311020-0123122320100230-3030331101022321-3020031113100220-1112233320020031) |
| `azure_vnet_site.vnet_attachments.vnet_list.subscription_id` | [azure_vnet_site.vnet_attachments.vnet_list.subscription_id](data-sources--cloud_connect--reference--group-001.md#canonical-0220321320030211-3023202301032231-1102022223123030-3310132100132110-0021011330100021-3013131020111220-3223102133323132-2113031221220310) |
| `azure_vnet_site.vnet_attachments.vnet_list.vnet_id` | [azure_vnet_site.vnet_attachments.vnet_list.vnet_id](data-sources--cloud_connect--reference--group-001.md#canonical-0323302110032333-3001322223023311-0332333021212230-1212021230000211-1222310000333223-0302320103223332-0220302321101012-3220312311132311) |
| `description` | [description](data-sources--cloud_connect--reference--group-001.md#canonical-0322301302023230-1000010020002332-2332311311222230-0313130102233333-0320103301300100-2131212102231320-0210121023031303-0002310332000211) |
| `id` | [id](data-sources--cloud_connect--reference--group-001.md#canonical-2022011201212303-2312301100103312-3223313130221101-2012303113223112-2030102222002130-3220303030010110-2123020100021330-3102033101123013) |
| `labels` | [labels](data-sources--cloud_connect--reference--group-001.md#canonical-0033320133021013-0322003311213233-3001022231022123-2020320201310020-2123133303111123-2122312220102333-0230213230013310-3032201012001132) |
| `name` | [name](data-sources--cloud_connect--reference--group-001.md#canonical-1333313323220130-2232002220300220-2200122030303222-2321130331320010-2013210033003221-2201113110203323-0223330100333022-2130320012013033) |
| `namespace` | [namespace](data-sources--cloud_connect--reference--group-001.md#canonical-2230020311230332-0133222222111101-2332110002313012-2110113133222103-0310312313232002-2201132103331100-1332210133100123-3213221320232223) |
| `segment` | [segment](data-sources--cloud_connect--reference--group-001.md#canonical-1321211200101101-3231233333312303-1023333013331310-2312022030301213-2122012103101310-0112330331030330-2210330113331103-0220230112321012) |
| `segment.name` | [segment.name](data-sources--cloud_connect--reference--group-001.md#canonical-1120103033202211-1312123030321303-2121112230310001-0313102102323020-0231032203133220-1223211211033102-1102210222303132-1223222011101103) |
| `segment.namespace` | [segment.namespace](data-sources--cloud_connect--reference--group-001.md#canonical-2131121001302133-2210232310322021-2212130023331232-2010213202000112-1023020323103002-3200313013202220-3202212133222302-2200320031010213) |
| `segment.tenant` | [segment.tenant](data-sources--cloud_connect--reference--group-001.md#canonical-3012212122202301-3333000220102202-3201010032223100-0131231003220021-3231300302102303-1133233003213230-0303100232233313-3332110102203222) |

<a id="canonical-0223203323331101-2013101330302110-0001032332112101-1230310312223332-3220110131203102-1213112203012003-2102220112311133-0310312002103132"></a>

## Next pages — Property reference / 031022023211 / 11

- [aws_provider](data-sources--cloud_connect--reference--group-001.md#canonical-3322210313330233-3203313001001101-1033130101220000-1313031021321122-3011231212132102-2121333313020110-1123112301023233-1202303010121200)
- [azure_vnet_site](data-sources--cloud_connect--reference--group-001.md#canonical-1310213230010202-0010122300232220-3030003101222331-3033212232021303-0201232322203122-3303121123112230-3233222203332331-2222210203131132)
- [segment](data-sources--cloud_connect--reference--group-001.md#canonical-0231311031231003-1321013301003130-3313203023221333-0013202013320101-2320200121101002-2203011312223323-1203000323320122-1200110112102001)
- [xcsh_cloud_connect](../data-sources/cloud_connect.md#canonical-2011032310310033-3311323122002031-2213302012003111-3033233110012013-2021203231210000-3221202000202111-0001331133100330-3301210132312000)

<a id="canonical-3322210313330233-3203313001001101-1033130101220000-1313031021321122-3011231212132102-2121333313020110-1123112301023233-1202303010121200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1003310200130211-0120302100020022-1113103113203021-1210233101332230-0111312132102331-1033033320323301-2002203303031313-1130132001112100"></a>

## aws_provider — aws_provider / 310122011033 / 2

Breadcrumbs:

- [xcsh_cloud_connect](../data-sources/cloud_connect.md#canonical-2011032310310033-3311323122002031-2213302012003111-3033233110012013-2021203231210000-3221202000202111-0001331133100330-3301210132312000)
- [Property reference](data-sources--cloud_connect--reference--group-001.md#canonical-0222032211111313-0003010022101312-0310023330122333-0201122213032311-0122213101110201-2002020111110210-1221320313132123-3300313332231200)
- aws_provider

<a id="canonical-0323202020332112-0003310021012122-0333333120010311-0010312113123221-2210103131221221-1102330110210113-2101030230221010-1133032313332022"></a>

Type: `"single"`. Computed.

\[OneOf: aws\_provider, Azure\_vnet\_site\] Configuration parameter for aws provider.

Upstream description:

Cloud Connect with AWS.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-site_type": "[\"aws_tgw_site\"]"
}
```

OneOf alternatives in this subsection:

- [aws_provider](data-sources--cloud_connect--reference--group-001.md#canonical-0323202020332112-0003310021012122-0333333120010311-0010312113123221-2210103131221221-1102330110210113-2101030230221010-1133032313332022)
- [azure_vnet_site](data-sources--cloud_connect--reference--group-001.md#canonical-2303331131320220-1121203202332230-2002103033303200-1330122233203313-2233002333201203-2210330300013330-3112012031201022-0212312313033300)

Select alternatives according to the provider validators above.

<a id="canonical-2033110032101000-1021121030102003-3322221232331200-0110030303030231-3013300210021310-3010323320212333-0121200032301221-2120022230212123"></a>

## Direct properties — aws_provider / 310122011033 / 3

- [aws_tgw_site](data-sources--cloud_connect--reference--group-001.md#canonical-1210201012310011-1202111032113023-2120321323312011-1101121121031011-0113230133300122-3231021212301121-1230133321330321-3230000102301202): complete subsection reference.

<a id="canonical-2032103333230020-1212313312202212-3222112030133103-0231131303113033-3301200103123330-1020103011313330-3032130233131320-0111011313301012"></a>

## Next pages — aws_provider / 310122011033 / 4

- [aws_provider.aws_tgw_site](data-sources--cloud_connect--reference--group-001.md#canonical-1210201012310011-1202111032113023-2120321323312011-1101121121031011-0113230133300122-3231021212301121-1230133321330321-3230000102301202)
- [Property reference](data-sources--cloud_connect--reference--group-001.md#canonical-0222032211111313-0003010022101312-0310023330122333-0201122213032311-0122213101110201-2002020111110210-1221320313132123-3300313332231200)
- [xcsh_cloud_connect](../data-sources/cloud_connect.md#canonical-2011032310310033-3311323122002031-2213302012003111-3033233110012013-2021203231210000-3221202000202111-0001331133100330-3301210132312000)

<a id="canonical-1210201012310011-1202111032113023-2120321323312011-1101121121031011-0113230133300122-3231021212301121-1230133321330321-3230000102301202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2132321033221003-3323113302212012-0213032000023322-3103221300221130-0312223011211132-2033033213100221-0131032021103032-3103001301221030"></a>

## aws_provider.aws_tgw_site — aws_tgw_site / 110000311001 / 2

Breadcrumbs:

- [xcsh_cloud_connect](../data-sources/cloud_connect.md#canonical-2011032310310033-3311323122002031-2213302012003111-3033233110012013-2021203231210000-3221202000202111-0001331133100330-3301210132312000)
- [Property reference](data-sources--cloud_connect--reference--group-001.md#canonical-0222032211111313-0003010022101312-0310023330122333-0201122213032311-0122213101110201-2002020111110210-1221320313132123-3300313332231200)
- [aws_provider](data-sources--cloud_connect--reference--group-001.md#canonical-3322210313330233-3203313001001101-1033130101220000-1313031021321122-3011231212132102-2121333313020110-1123112301023233-1202303010121200)
- aws_provider.aws_tgw_site

<a id="canonical-0323101110301210-2203102021231111-0132323321120133-0010211122013322-2002211110011322-2213313201010232-3023023313322231-3113320131200221"></a>

Type: `"single"`. Computed.

AWS TGW Site Type. Cloud Connect AWS TGW Site Type.

Upstream description:

Cloud Connect AWS TGW Site Type.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0103002121110320-3123133323210221-2010121300330123-0331031033221021-2103002133220001-2012033300133221-0113012103103030-0100111133121230"></a>

## Direct properties — aws_tgw_site / 110000311001 / 3

- [cred](data-sources--cloud_connect--reference--group-001.md#canonical-3030032220101032-1330222322300222-1133100011310011-3230102311112133-1231033210023330-0201311101322003-3001110010120312-2121113132123332): complete subsection reference.

- [site](data-sources--cloud_connect--reference--group-001.md#canonical-0331121033111213-0100013022110333-0100002030121222-0212130132233330-0223131031010322-0030330010012112-3111032322130210-1122111132220301): complete subsection reference.

- [vpc_attachments](data-sources--cloud_connect--reference--group-001.md#canonical-3200311203003333-3212000013100213-1102221330110123-2132202003301122-3202310210102320-3120022021203120-3110111203122330-1233233203001313): complete subsection reference.

<a id="canonical-3232033330321320-2031003000230212-2331111211231212-1313212103302031-2223121131010122-0311113220113310-2200231133232222-3300102311332122"></a>

## Next pages — aws_tgw_site / 110000311001 / 4

- [aws_provider.aws_tgw_site.cred](data-sources--cloud_connect--reference--group-001.md#canonical-3030032220101032-1330222322300222-1133100011310011-3230102311112133-1231033210023330-0201311101322003-3001110010120312-2121113132123332)
- [aws_provider.aws_tgw_site.site](data-sources--cloud_connect--reference--group-001.md#canonical-0331121033111213-0100013022110333-0100002030121222-0212130132233330-0223131031010322-0030330010012112-3111032322130210-1122111132220301)
- [aws_provider.aws_tgw_site.vpc_attachments](data-sources--cloud_connect--reference--group-001.md#canonical-3200311203003333-3212000013100213-1102221330110123-2132202003301122-3202310210102320-3120022021203120-3110111203122330-1233233203001313)
- [aws_provider](data-sources--cloud_connect--reference--group-001.md#canonical-3322210313330233-3203313001001101-1033130101220000-1313031021321122-3011231212132102-2121333313020110-1123112301023233-1202303010121200)
- [xcsh_cloud_connect](../data-sources/cloud_connect.md#canonical-2011032310310033-3311323122002031-2213302012003111-3033233110012013-2021203231210000-3221202000202111-0001331133100330-3301210132312000)

<a id="canonical-3030032220101032-1330222322300222-1133100011310011-3230102311112133-1231033210023330-0201311101322003-3001110010120312-2121113132123332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0201231121110131-0131312221312202-0332201320221102-0130223322222203-3201023200233320-2200302100222122-2003200321133333-0030321231220313"></a>

## aws_provider.aws_tgw_site.cred — cred / 113223210112 / 2

Breadcrumbs:

- [xcsh_cloud_connect](../data-sources/cloud_connect.md#canonical-2011032310310033-3311323122002031-2213302012003111-3033233110012013-2021203231210000-3221202000202111-0001331133100330-3301210132312000)
- [Property reference](data-sources--cloud_connect--reference--group-001.md#canonical-0222032211111313-0003010022101312-0310023330122333-0201122213032311-0122213101110201-2002020111110210-1221320313132123-3300313332231200)
- [aws_provider](data-sources--cloud_connect--reference--group-001.md#canonical-3322210313330233-3203313001001101-1033130101220000-1313031021321122-3011231212132102-2121333313020110-1123112301023233-1202303010121200)
- [aws_provider.aws_tgw_site](data-sources--cloud_connect--reference--group-001.md#canonical-1210201012310011-1202111032113023-2120321323312011-1101121121031011-0113230133300122-3231021212301121-1230133321330321-3230000102301202)
- aws_provider.aws_tgw_site.cred

<a id="canonical-3220202231021112-2232322331331033-0313231233130001-3023102010030032-1331033133122231-1313220330223112-1310100020020013-2123333312111123"></a>

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

<a id="canonical-3103201311303113-2012201222002321-3032321303212303-1123100031211322-3310200333001030-1013032302310113-3033312002300032-0031020110231001"></a>

## Direct properties — cred / 113223210112 / 3

<a id="canonical-1220122320022201-0102121123003021-1313312212333220-0032322022223312-1011123221322200-1221033230322311-3101132001203111-2223123123023003"></a>

<a id="canonical-3001123003323110-0010323221132301-1331011011010130-3110032012322313-2320313011223211-3201223020002320-3033220103103201-1222223101023300"></a>

## name property — cred / 113223210112 / 4

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

<a id="canonical-2022011120013023-0333100101231203-3300330333320103-0100213101210311-1021331011220020-2321230123131210-1031233321202000-3323000032330000"></a>

<a id="canonical-2330012323100230-2012010321232112-1322233003223122-0300002223222013-3131100121023233-0022321322321332-0001222011013222-0202110303033022"></a>

## namespace property — cred / 113223210112 / 5

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

<a id="canonical-3023221212110203-0330210221321101-2213100310112030-2100123133321120-0212230022321312-0011312302322220-2013103120321202-3120312221222210"></a>

<a id="canonical-3112331112331232-0232111302333023-0301113332312202-0103130212333210-0131032320210121-1130133210100031-0323210332103032-0132012202302131"></a>

## tenant property — cred / 113223210112 / 6

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

<a id="canonical-0001011323202130-0310303333210300-0333002133221230-2312103330113022-3233311311021322-2030203131103320-2210020031303002-3010330233030333"></a>

## Next pages — cred / 113223210112 / 7

- [aws_provider.aws_tgw_site](data-sources--cloud_connect--reference--group-001.md#canonical-1210201012310011-1202111032113023-2120321323312011-1101121121031011-0113230133300122-3231021212301121-1230133321330321-3230000102301202)
- [xcsh_cloud_connect](../data-sources/cloud_connect.md#canonical-2011032310310033-3311323122002031-2213302012003111-3033233110012013-2021203231210000-3221202000202111-0001331133100330-3301210132312000)

<a id="canonical-0331121033111213-0100013022110333-0100002030121222-0212130132233330-0223131031010322-0030330010012112-3111032322130210-1122111132220301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3101100003301301-0123320122002103-0220233323100102-3031003233312200-0233223213122111-2201122321030223-3032302322202133-2133202032211132"></a>

## aws_provider.aws_tgw_site.site — site / 322233233200 / 2

Breadcrumbs:

- [xcsh_cloud_connect](../data-sources/cloud_connect.md#canonical-2011032310310033-3311323122002031-2213302012003111-3033233110012013-2021203231210000-3221202000202111-0001331133100330-3301210132312000)
- [Property reference](data-sources--cloud_connect--reference--group-001.md#canonical-0222032211111313-0003010022101312-0310023330122333-0201122213032311-0122213101110201-2002020111110210-1221320313132123-3300313332231200)
- [aws_provider](data-sources--cloud_connect--reference--group-001.md#canonical-3322210313330233-3203313001001101-1033130101220000-1313031021321122-3011231212132102-2121333313020110-1123112301023233-1202303010121200)
- [aws_provider.aws_tgw_site](data-sources--cloud_connect--reference--group-001.md#canonical-1210201012310011-1202111032113023-2120321323312011-1101121121031011-0113230133300122-3231021212301121-1230133321330321-3230000102301202)
- aws_provider.aws_tgw_site.site

<a id="canonical-2031121000320021-1003023101300120-3203010221230023-2021213133000030-0212031331330330-3302233112123130-3121110010232020-3323020022311132"></a>

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

<a id="canonical-0120321032123310-1223122023023323-3133030023332223-2023133123021101-2332012323333222-2111200211310311-2220321211323333-1031110113130122"></a>

## Direct properties — site / 322233233200 / 3

<a id="canonical-0201223212302023-0311203230321020-1003203120221011-2202121213011102-3310333321131030-2310012022013103-3322010321130330-1020022212132321"></a>

<a id="canonical-1210033002210110-3312312310101000-1220130123002200-2313011020131032-3103322300122000-2102013321030322-3112203311121213-0000023110011331"></a>

## name property — site / 322233233200 / 4

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

<a id="canonical-1313310112133300-2032311231000203-1102221330031221-2332330301033213-1131101302210113-2310131230131010-0320331330201133-0023220030132230"></a>

<a id="canonical-0233223300213020-2012203302220311-0023032113130211-2102003022323023-1001332130030311-2222002203001113-0301332330001302-2323033303232131"></a>

## namespace property — site / 322233233200 / 5

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

<a id="canonical-0020202012230021-0322102200322103-3031320102231003-2003231301200003-1323300113111012-3211233300331030-1030011213312321-2211111020103312"></a>

<a id="canonical-3111303231211232-2002312022203332-2331331202201310-2023233023303002-1111201000210211-0223003230302202-1021311301030301-1310331031313200"></a>

## tenant property — site / 322233233200 / 6

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

<a id="canonical-2131103210020321-3311023000222323-0331202212122202-3033332321220231-0202312122231310-3002211233331302-1110012323310100-0322031231022111"></a>

## Next pages — site / 322233233200 / 7

- [aws_provider.aws_tgw_site](data-sources--cloud_connect--reference--group-001.md#canonical-1210201012310011-1202111032113023-2120321323312011-1101121121031011-0113230133300122-3231021212301121-1230133321330321-3230000102301202)
- [xcsh_cloud_connect](../data-sources/cloud_connect.md#canonical-2011032310310033-3311323122002031-2213302012003111-3033233110012013-2021203231210000-3221202000202111-0001331133100330-3301210132312000)

<a id="canonical-3200311203003333-3212000013100213-1102221330110123-2132202003301122-3202310210102320-3120022021203120-3110111203122330-1233233203001313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1002132232103222-1013331302133003-3110110233122122-3210331011123102-1322101010023303-2131112033331203-0002323213100201-2221102232013012"></a>

## aws_provider.aws_tgw_site.vpc_attachments — vpc_attachments / 013132133023 / 2

Breadcrumbs:

- [xcsh_cloud_connect](../data-sources/cloud_connect.md#canonical-2011032310310033-3311323122002031-2213302012003111-3033233110012013-2021203231210000-3221202000202111-0001331133100330-3301210132312000)
- [Property reference](data-sources--cloud_connect--reference--group-001.md#canonical-0222032211111313-0003010022101312-0310023330122333-0201122213032311-0122213101110201-2002020111110210-1221320313132123-3300313332231200)
- [aws_provider](data-sources--cloud_connect--reference--group-001.md#canonical-3322210313330233-3203313001001101-1033130101220000-1313031021321122-3011231212132102-2121333313020110-1123112301023233-1202303010121200)
- [aws_provider.aws_tgw_site](data-sources--cloud_connect--reference--group-001.md#canonical-1210201012310011-1202111032113023-2120321323312011-1101121121031011-0113230133300122-3231021212301121-1230133321330321-3230000102301202)
- aws_provider.aws_tgw_site.vpc_attachments

<a id="canonical-1310322300123113-1010220023211111-1111023033321032-3202323020313312-2121003311010202-2220130203213130-2321100213002313-0122333010010301"></a>

Type: `"single"`. Computed.

Configuration parameter for vpc attachments.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1232123322333120-1000332010321032-0200321302302203-3131301003322331-3220301220300211-2323023333222212-2132132332011322-1230022333123113"></a>

## Direct properties — vpc_attachments / 013132133023 / 3

- [vpc_list](data-sources--cloud_connect--reference--group-001.md#canonical-3130311012000102-3303222030000301-2101203322100102-3320221211100031-0320011012231032-2322023202333123-1013132123020102-3330213131230103): complete subsection reference.

<a id="canonical-2100233010133020-3122230221233313-1300200310031000-2332112020303030-0102020023110030-2222233013010303-2031321003011211-3133232001033303"></a>

## Next pages — vpc_attachments / 013132133023 / 4

- [aws_provider.aws_tgw_site.vpc_attachments.vpc_list](data-sources--cloud_connect--reference--group-001.md#canonical-3130311012000102-3303222030000301-2101203322100102-3320221211100031-0320011012231032-2322023202333123-1013132123020102-3330213131230103)
- [aws_provider.aws_tgw_site](data-sources--cloud_connect--reference--group-001.md#canonical-1210201012310011-1202111032113023-2120321323312011-1101121121031011-0113230133300122-3231021212301121-1230133321330321-3230000102301202)
- [xcsh_cloud_connect](../data-sources/cloud_connect.md#canonical-2011032310310033-3311323122002031-2213302012003111-3033233110012013-2021203231210000-3221202000202111-0001331133100330-3301210132312000)

<a id="canonical-3130311012000102-3303222030000301-2101203322100102-3320221211100031-0320011012231032-2322023202333123-1013132123020102-3330213131230103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3010212021003032-1203013123032022-1302122322213313-0003300112311233-1012301303112103-3300002032000013-2222010012010003-3121123110010323"></a>

## aws_provider.aws_tgw_site.vpc_attachments.vpc_list — vpc_list / 200220222031 / 2

Breadcrumbs:

- [xcsh_cloud_connect](../data-sources/cloud_connect.md#canonical-2011032310310033-3311323122002031-2213302012003111-3033233110012013-2021203231210000-3221202000202111-0001331133100330-3301210132312000)
- [Property reference](data-sources--cloud_connect--reference--group-001.md#canonical-0222032211111313-0003010022101312-0310023330122333-0201122213032311-0122213101110201-2002020111110210-1221320313132123-3300313332231200)
- [aws_provider](data-sources--cloud_connect--reference--group-001.md#canonical-3322210313330233-3203313001001101-1033130101220000-1313031021321122-3011231212132102-2121333313020110-1123112301023233-1202303010121200)
- [aws_provider.aws_tgw_site](data-sources--cloud_connect--reference--group-001.md#canonical-1210201012310011-1202111032113023-2120321323312011-1101121121031011-0113230133300122-3231021212301121-1230133321330321-3230000102301202)
- [aws_provider.aws_tgw_site.vpc_attachments](data-sources--cloud_connect--reference--group-001.md#canonical-3200311203003333-3212000013100213-1102221330110123-2132202003301122-3202310210102320-3120022021203120-3110111203122330-1233233203001313)
- aws_provider.aws_tgw_site.vpc_attachments.vpc_list

<a id="canonical-3003132211300032-3100010103222301-3330301023321103-1333103221331222-3100313331211232-0003003003112120-3122000003121322-2030311201201002"></a>

Type: `"list"`. Computed.

VPC List. Collection of items or values

Upstream description:

Collection of items or values

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 128,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 128,
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
    "ves.io.schema.rules.repeated.max_items": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "128"
  }
}
```

<a id="canonical-3233320012020312-0110212022232302-0001222120300233-2322320131123221-3133233020230122-3101212002102132-3001121302302003-1101020300100120"></a>

## Direct properties — vpc_list / 200220222031 / 3

- [custom_routing](data-sources--cloud_connect--reference--group-001.md#canonical-3131100103223010-2311321131333301-3213320232002220-1202330330221100-0320030133210202-3102321111313012-0202120103120113-2032331112101222): complete subsection reference.

- [default_route](data-sources--cloud_connect--reference--group-001.md#canonical-0333233202010131-1000222030131223-0202030023000321-3132123300323333-2010021110110333-0120331213112021-0323310011103333-2021311200221130): complete subsection reference.

- [labels](data-sources--cloud_connect--reference--group-001.md#canonical-2331221100110300-1113031003120322-3010333302212321-1033202110023100-3311223310010132-0120300201330202-3323031111310023-1002112333320021): complete subsection reference.

- [manual_routing](data-sources--cloud_connect--reference--group-001.md#canonical-2000103210333331-1332323303303110-0301003132012213-2120303213221223-2133323002201033-2233212101203122-0101202023103300-0301321331310131): complete subsection reference.

<a id="canonical-0132121310112132-0023131020323103-0133321003102330-2001111300003332-3010032022132221-0233232202033113-0202131123031032-2011321332132020"></a>

<a id="canonical-3000333311111222-0203103232211310-2032332112331320-2020203022212113-3231011122333320-2313200220003213-2123130001102120-2231123210303130"></a>

## vpc_id property — vpc_list / 200220222031 / 4

Type: `"string"`. Computed.

Enter the VPC ID of the VPC to be attached.

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
    },
    "pattern": "^(vpc-)([a-z0-9]{8}|[a-z0-9]{17})$"
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
    "ves.io.schema.rules.string.pattern": "^(vpc-)([a-z0-9]{8}|[a-z0-9]{17})$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.pattern": "^(vpc-)([a-z0-9]{8}|[a-z0-9]{17})$"
  }
}
```

<a id="canonical-3313200032220133-1221321113020310-3112310321320000-1111212121001101-0120132312200330-3333200222300131-0300210232203322-1201331113102232"></a>

## Next pages — vpc_list / 200220222031 / 5

- [aws_provider.aws_tgw_site.vpc_attachments.vpc_list.custom_routing](data-sources--cloud_connect--reference--group-001.md#canonical-3131100103223010-2311321131333301-3213320232002220-1202330330221100-0320030133210202-3102321111313012-0202120103120113-2032331112101222)
- [aws_provider.aws_tgw_site.vpc_attachments.vpc_list.default_route](data-sources--cloud_connect--reference--group-001.md#canonical-0333233202010131-1000222030131223-0202030023000321-3132123300323333-2010021110110333-0120331213112021-0323310011103333-2021311200221130)
- [aws_provider.aws_tgw_site.vpc_attachments.vpc_list.labels](data-sources--cloud_connect--reference--group-001.md#canonical-2331221100110300-1113031003120322-3010333302212321-1033202110023100-3311223310010132-0120300201330202-3323031111310023-1002112333320021)
- [aws_provider.aws_tgw_site.vpc_attachments.vpc_list.manual_routing](data-sources--cloud_connect--reference--group-001.md#canonical-2000103210333331-1332323303303110-0301003132012213-2120303213221223-2133323002201033-2233212101203122-0101202023103300-0301321331310131)
- [aws_provider.aws_tgw_site.vpc_attachments](data-sources--cloud_connect--reference--group-001.md#canonical-3200311203003333-3212000013100213-1102221330110123-2132202003301122-3202310210102320-3120022021203120-3110111203122330-1233233203001313)
- [xcsh_cloud_connect](../data-sources/cloud_connect.md#canonical-2011032310310033-3311323122002031-2213302012003111-3033233110012013-2021203231210000-3221202000202111-0001331133100330-3301210132312000)

<a id="canonical-3131100103223010-2311321131333301-3213320232002220-1202330330221100-0320030133210202-3102321111313012-0202120103120113-2032331112101222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3200020003100231-1031132320220032-3300213002133001-1203033213302022-1331303230330133-3112011101110003-0322013302332001-1010232020000003"></a>

## aws_provider.aws_tgw_site.vpc_attachments.vpc_list.custom_routing — custom_routing / 320300230002 / 2

Breadcrumbs:

- [xcsh_cloud_connect](../data-sources/cloud_connect.md#canonical-2011032310310033-3311323122002031-2213302012003111-3033233110012013-2021203231210000-3221202000202111-0001331133100330-3301210132312000)
- [Property reference](data-sources--cloud_connect--reference--group-001.md#canonical-0222032211111313-0003010022101312-0310023330122333-0201122213032311-0122213101110201-2002020111110210-1221320313132123-3300313332231200)
- [aws_provider](data-sources--cloud_connect--reference--group-001.md#canonical-3322210313330233-3203313001001101-1033130101220000-1313031021321122-3011231212132102-2121333313020110-1123112301023233-1202303010121200)
- [aws_provider.aws_tgw_site](data-sources--cloud_connect--reference--group-001.md#canonical-1210201012310011-1202111032113023-2120321323312011-1101121121031011-0113230133300122-3231021212301121-1230133321330321-3230000102301202)
- [aws_provider.aws_tgw_site.vpc_attachments](data-sources--cloud_connect--reference--group-001.md#canonical-3200311203003333-3212000013100213-1102221330110123-2132202003301122-3202310210102320-3120022021203120-3110111203122330-1233233203001313)
- [aws_provider.aws_tgw_site.vpc_attachments.vpc_list](data-sources--cloud_connect--reference--group-001.md#canonical-3130311012000102-3303222030000301-2101203322100102-3320221211100031-0320011012231032-2322023202333123-1013132123020102-3330213131230103)
- aws_provider.aws_tgw_site.vpc_attachments.vpc_list.custom_routing

<a id="canonical-0100223003113001-1323032321321022-1222002133203021-3203331203320023-2121133303110333-0002312002210221-0111113313023030-2223133103032322"></a>

Type: `"single"`. Computed.

AWS Route Table List. AWS Route Table List.

Upstream description:

AWS Route Table List.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3102322221030003-0201111133330132-1130200113203330-0103301302112311-0102011301323021-2232003223302221-0320022020002112-2100300120211023"></a>

## Direct properties — custom_routing / 320300230002 / 3

- [route_tables](data-sources--cloud_connect--reference--group-001.md#canonical-0311100001212222-1230333310312121-3021122331021313-1021001230001000-2232211320020210-2112020110301122-2210102002220203-1131110313113001): complete subsection reference.

<a id="canonical-1131210132023311-0122310010303003-0102313201201131-1030012302130201-3110001231313301-2332212131212310-2000232312320101-2032110323120332"></a>

## Next pages — custom_routing / 320300230002 / 4

- [aws_provider.aws_tgw_site.vpc_attachments.vpc_list.custom_routing.route_tables](data-sources--cloud_connect--reference--group-001.md#canonical-0311100001212222-1230333310312121-3021122331021313-1021001230001000-2232211320020210-2112020110301122-2210102002220203-1131110313113001)
- [aws_provider.aws_tgw_site.vpc_attachments.vpc_list](data-sources--cloud_connect--reference--group-001.md#canonical-3130311012000102-3303222030000301-2101203322100102-3320221211100031-0320011012231032-2322023202333123-1013132123020102-3330213131230103)
- [xcsh_cloud_connect](../data-sources/cloud_connect.md#canonical-2011032310310033-3311323122002031-2213302012003111-3033233110012013-2021203231210000-3221202000202111-0001331133100330-3301210132312000)

<a id="canonical-0311100001212222-1230333310312121-3021122331021313-1021001230001000-2232211320020210-2112020110301122-2210102002220203-1131110313113001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2210320120220320-1320213322122001-0032230221032233-2223121013002231-1130231122303321-1113220021021310-2323313203121021-0133001213222102"></a>

## aws_provider.aws_tgw_site.vpc_attachments.vpc_list.custom_routing.route_tables — route_tables / 012001021030 / 2

Breadcrumbs:

- [xcsh_cloud_connect](../data-sources/cloud_connect.md#canonical-2011032310310033-3311323122002031-2213302012003111-3033233110012013-2021203231210000-3221202000202111-0001331133100330-3301210132312000)
- [Property reference](data-sources--cloud_connect--reference--group-001.md#canonical-0222032211111313-0003010022101312-0310023330122333-0201122213032311-0122213101110201-2002020111110210-1221320313132123-3300313332231200)
- [aws_provider](data-sources--cloud_connect--reference--group-001.md#canonical-3322210313330233-3203313001001101-1033130101220000-1313031021321122-3011231212132102-2121333313020110-1123112301023233-1202303010121200)
- [aws_provider.aws_tgw_site](data-sources--cloud_connect--reference--group-001.md#canonical-1210201012310011-1202111032113023-2120321323312011-1101121121031011-0113230133300122-3231021212301121-1230133321330321-3230000102301202)
- [aws_provider.aws_tgw_site.vpc_attachments](data-sources--cloud_connect--reference--group-001.md#canonical-3200311203003333-3212000013100213-1102221330110123-2132202003301122-3202310210102320-3120022021203120-3110111203122330-1233233203001313)
- [aws_provider.aws_tgw_site.vpc_attachments.vpc_list](data-sources--cloud_connect--reference--group-001.md#canonical-3130311012000102-3303222030000301-2101203322100102-3320221211100031-0320011012231032-2322023202333123-1013132123020102-3330213131230103)
- [aws_provider.aws_tgw_site.vpc_attachments.vpc_list.custom_routing](data-sources--cloud_connect--reference--group-001.md#canonical-3131100103223010-2311321131333301-3213320232002220-1202330330221100-0320030133210202-3102321111313012-0202120103120113-2032331112101222)
- aws_provider.aws_tgw_site.vpc_attachments.vpc_list.custom_routing.route_tables

<a id="canonical-3010211203312021-1302213213133231-0011022122211320-3321231223222101-1332202221223232-2232222102210133-2300123302012213-2022031000013111"></a>

Type: `"list"`. Computed.

List of route tables. Route Tables.

Upstream description:

Route Tables.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 200,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 200,
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
    "ves.io.schema.rules.repeated.max_items": "200",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "200",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2210221221301203-1112310032132210-3301201321001133-1013111311222222-0011332113333321-0310123332111303-3013213311110031-0211333311331231"></a>

## Direct properties — route_tables / 012001021030 / 3

<a id="canonical-3102330220020133-3310013221330302-0113231231011300-0213013032012000-3300201321130232-1133030112020121-2203030001313322-1122230013003232"></a>

<a id="canonical-3330312000202112-0330010013000333-3231001103112332-2001201320031022-1321323022221133-2113100320230320-1123212321210202-2300100212323222"></a>

## route_table_id property — route_tables / 012001021030 / 4

Type: `"string"`. Computed.

Route table ID. Route table ID.

Upstream description:

Route table ID.

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
    },
    "pattern": "^(rtb-)([a-z0-9]{8}|[a-z0-9]{17})$"
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.pattern": "^(rtb-)([a-z0-9]{8}|[a-z0-9]{17})$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.pattern": "^(rtb-)([a-z0-9]{8}|[a-z0-9]{17})$"
  }
}
```

<a id="canonical-1031211001233031-1330330331211230-3003320212010220-2302322323130103-0100203222023230-1302220122132130-3000101303302032-2130112210300210"></a>

<a id="canonical-2102303212321101-3300131132011003-1132033003311010-0000231200210131-0303010112101102-2000020210230221-0113103232320103-3221333131223103"></a>

## static_routes property — route_tables / 012001021030 / 5

Type: `["list", "string"]`. Computed.

Static Routes. List of Static Routes.

Upstream description:

List of Static Routes.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 50,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 50,
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
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "50",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "50",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0202221332131031-3233032002103011-3231330113300320-0202313313231020-1330030022010121-2221112022102032-2111100300211110-0110133323331301"></a>

## Next pages — route_tables / 012001021030 / 6

- [aws_provider.aws_tgw_site.vpc_attachments.vpc_list.custom_routing](data-sources--cloud_connect--reference--group-001.md#canonical-3131100103223010-2311321131333301-3213320232002220-1202330330221100-0320030133210202-3102321111313012-0202120103120113-2032331112101222)
- [xcsh_cloud_connect](../data-sources/cloud_connect.md#canonical-2011032310310033-3311323122002031-2213302012003111-3033233110012013-2021203231210000-3221202000202111-0001331133100330-3301210132312000)

<a id="canonical-0333233202010131-1000222030131223-0202030023000321-3132123300323333-2010021110110333-0120331213112021-0323310011103333-2021311200221130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0011300203122322-0302212312030313-3230223002001300-1123221111030333-2213200102231330-0100320300201011-3222102202302000-1133230220201031"></a>

## aws_provider.aws_tgw_site.vpc_attachments.vpc_list.default_route — default_route / 333000030332 / 2

Breadcrumbs:

- [xcsh_cloud_connect](../data-sources/cloud_connect.md#canonical-2011032310310033-3311323122002031-2213302012003111-3033233110012013-2021203231210000-3221202000202111-0001331133100330-3301210132312000)
- [Property reference](data-sources--cloud_connect--reference--group-001.md#canonical-0222032211111313-0003010022101312-0310023330122333-0201122213032311-0122213101110201-2002020111110210-1221320313132123-3300313332231200)
- [aws_provider](data-sources--cloud_connect--reference--group-001.md#canonical-3322210313330233-3203313001001101-1033130101220000-1313031021321122-3011231212132102-2121333313020110-1123112301023233-1202303010121200)
- [aws_provider.aws_tgw_site](data-sources--cloud_connect--reference--group-001.md#canonical-1210201012310011-1202111032113023-2120321323312011-1101121121031011-0113230133300122-3231021212301121-1230133321330321-3230000102301202)
- [aws_provider.aws_tgw_site.vpc_attachments](data-sources--cloud_connect--reference--group-001.md#canonical-3200311203003333-3212000013100213-1102221330110123-2132202003301122-3202310210102320-3120022021203120-3110111203122330-1233233203001313)
- [aws_provider.aws_tgw_site.vpc_attachments.vpc_list](data-sources--cloud_connect--reference--group-001.md#canonical-3130311012000102-3303222030000301-2101203322100102-3320221211100031-0320011012231032-2322023202333123-1013132123020102-3330213131230103)
- aws_provider.aws_tgw_site.vpc_attachments.vpc_list.default_route

<a id="canonical-3212213132131302-2112023130010330-2030203133211133-2012100213322031-1000303202032333-1201011223203121-2133301002010231-0130133102103020"></a>

Type: `"single"`. Computed.

Configuration parameter for default route.

Upstream description:

Select Override Default Route Choice.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-default_route_choice": "[\"all_route_tables\",\"selective_route_tables\"]"
}
```

<a id="canonical-2303101211121230-2331333301232123-3002223203301013-1011032021330211-3121232003230000-3333031122100301-2022112212223120-1333303131222323"></a>

## Direct properties — default_route / 333000030332 / 3

- [all_route_tables](data-sources--cloud_connect--reference--group-001.md#canonical-0121011012231301-1220030300020032-3031213002133100-0330230111013030-3101323323230222-3211313203023002-1123022003011030-2331202210333210): complete subsection reference.

- [selective_route_tables](data-sources--cloud_connect--reference--group-001.md#canonical-0333032303100212-0202213010201203-0022200021121321-2233231220313020-0201131133202111-2103012232030003-0121333331200112-3302003203021122): complete subsection reference.

<a id="canonical-1231121100230121-1032303312223301-3020311101330210-0222011012010311-3123332002311312-2020010201220011-2010001213123020-0121330012111231"></a>

## Next pages — default_route / 333000030332 / 4

- [aws_provider.aws_tgw_site.vpc_attachments.vpc_list.default_route.all_route_tables](data-sources--cloud_connect--reference--group-001.md#canonical-0121011012231301-1220030300020032-3031213002133100-0330230111013030-3101323323230222-3211313203023002-1123022003011030-2331202210333210)
- [aws_provider.aws_tgw_site.vpc_attachments.vpc_list.default_route.selective_route_tables](data-sources--cloud_connect--reference--group-001.md#canonical-0333032303100212-0202213010201203-0022200021121321-2233231220313020-0201131133202111-2103012232030003-0121333331200112-3302003203021122)
- [aws_provider.aws_tgw_site.vpc_attachments.vpc_list](data-sources--cloud_connect--reference--group-001.md#canonical-3130311012000102-3303222030000301-2101203322100102-3320221211100031-0320011012231032-2322023202333123-1013132123020102-3330213131230103)
- [xcsh_cloud_connect](../data-sources/cloud_connect.md#canonical-2011032310310033-3311323122002031-2213302012003111-3033233110012013-2021203231210000-3221202000202111-0001331133100330-3301210132312000)

<a id="canonical-0121011012231301-1220030300020032-3031213002133100-0330230111013030-3101323323230222-3211313203023002-1123022003011030-2331202210333210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0320030123120322-0312322210331333-3201303301023111-0122032101300233-1113101103032032-0213222032321123-0211001222300123-0033212202233113"></a>

## aws_provider.aws_tgw_site.vpc_attachments.vpc_list.default_route.all_route_tables — all_route_tables / 320233113110 / 2

Breadcrumbs:

- [xcsh_cloud_connect](../data-sources/cloud_connect.md#canonical-2011032310310033-3311323122002031-2213302012003111-3033233110012013-2021203231210000-3221202000202111-0001331133100330-3301210132312000)
- [Property reference](data-sources--cloud_connect--reference--group-001.md#canonical-0222032211111313-0003010022101312-0310023330122333-0201122213032311-0122213101110201-2002020111110210-1221320313132123-3300313332231200)
- [aws_provider](data-sources--cloud_connect--reference--group-001.md#canonical-3322210313330233-3203313001001101-1033130101220000-1313031021321122-3011231212132102-2121333313020110-1123112301023233-1202303010121200)
- [aws_provider.aws_tgw_site](data-sources--cloud_connect--reference--group-001.md#canonical-1210201012310011-1202111032113023-2120321323312011-1101121121031011-0113230133300122-3231021212301121-1230133321330321-3230000102301202)
- [aws_provider.aws_tgw_site.vpc_attachments](data-sources--cloud_connect--reference--group-001.md#canonical-3200311203003333-3212000013100213-1102221330110123-2132202003301122-3202310210102320-3120022021203120-3110111203122330-1233233203001313)
- [aws_provider.aws_tgw_site.vpc_attachments.vpc_list](data-sources--cloud_connect--reference--group-001.md#canonical-3130311012000102-3303222030000301-2101203322100102-3320221211100031-0320011012231032-2322023202333123-1013132123020102-3330213131230103)
- [aws_provider.aws_tgw_site.vpc_attachments.vpc_list.default_route](data-sources--cloud_connect--reference--group-001.md#canonical-0333233202010131-1000222030131223-0202030023000321-3132123300323333-2010021110110333-0120331213112021-0323310011103333-2021311200221130)
- aws_provider.aws_tgw_site.vpc_attachments.vpc_list.default_route.all_route_tables

<a id="canonical-3131010131022100-0233201030003013-1200231122103303-0301322200022312-2023112310022132-2132001002312221-1223112333212031-1311203110123200"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for all route tables.

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

<a id="canonical-0101112013021332-2012012102101210-0003020020210110-2330030001311023-1113310212112003-0002020301102031-3221113230011100-0113103212022021"></a>

## Direct properties — all_route_tables / 320233113110 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0110032313030021-0301022231311201-3333013110113221-0332232323001222-0133012202323222-0300133102111000-0330100131100131-3120333132300021"></a>

## Next pages — all_route_tables / 320233113110 / 4

- [aws_provider.aws_tgw_site.vpc_attachments.vpc_list.default_route](data-sources--cloud_connect--reference--group-001.md#canonical-0333233202010131-1000222030131223-0202030023000321-3132123300323333-2010021110110333-0120331213112021-0323310011103333-2021311200221130)
- [xcsh_cloud_connect](../data-sources/cloud_connect.md#canonical-2011032310310033-3311323122002031-2213302012003111-3033233110012013-2021203231210000-3221202000202111-0001331133100330-3301210132312000)

<a id="canonical-0333032303100212-0202213010201203-0022200021121321-2233231220313020-0201131133202111-2103012232030003-0121333331200112-3302003203021122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0213220110012012-2333221100203220-2210021222130023-0012131120222010-0203122013030200-3223112122201322-3002030200011103-0001020232121022"></a>

## aws_provider.aws_tgw_site.vpc_attachments.vpc_list.default_route.selective_route_tables — selective_route_tables / 002322120233 / 2

Breadcrumbs:

- [xcsh_cloud_connect](../data-sources/cloud_connect.md#canonical-2011032310310033-3311323122002031-2213302012003111-3033233110012013-2021203231210000-3221202000202111-0001331133100330-3301210132312000)
- [Property reference](data-sources--cloud_connect--reference--group-001.md#canonical-0222032211111313-0003010022101312-0310023330122333-0201122213032311-0122213101110201-2002020111110210-1221320313132123-3300313332231200)
- [aws_provider](data-sources--cloud_connect--reference--group-001.md#canonical-3322210313330233-3203313001001101-1033130101220000-1313031021321122-3011231212132102-2121333313020110-1123112301023233-1202303010121200)
- [aws_provider.aws_tgw_site](data-sources--cloud_connect--reference--group-001.md#canonical-1210201012310011-1202111032113023-2120321323312011-1101121121031011-0113230133300122-3231021212301121-1230133321330321-3230000102301202)
- [aws_provider.aws_tgw_site.vpc_attachments](data-sources--cloud_connect--reference--group-001.md#canonical-3200311203003333-3212000013100213-1102221330110123-2132202003301122-3202310210102320-3120022021203120-3110111203122330-1233233203001313)
- [aws_provider.aws_tgw_site.vpc_attachments.vpc_list](data-sources--cloud_connect--reference--group-001.md#canonical-3130311012000102-3303222030000301-2101203322100102-3320221211100031-0320011012231032-2322023202333123-1013132123020102-3330213131230103)
- [aws_provider.aws_tgw_site.vpc_attachments.vpc_list.default_route](data-sources--cloud_connect--reference--group-001.md#canonical-0333233202010131-1000222030131223-0202030023000321-3132123300323333-2010021110110333-0120331213112021-0323310011103333-2021311200221130)
- aws_provider.aws_tgw_site.vpc_attachments.vpc_list.default_route.selective_route_tables

<a id="canonical-2330320210133210-2322002012313033-1000333010011302-3102320333120102-0231002310122120-3330312213300012-1312221111232231-0321211331120332"></a>

Type: `"single"`. Computed.

Configuration parameter for selective route tables.

Upstream description:

AWS Route Table.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0121233232102130-2221132123002310-1210330010003013-2112112111320122-3131310310133201-2320233012210012-2111121132301130-3303020023022210"></a>

## Direct properties — selective_route_tables / 002322120233 / 3

<a id="canonical-2122110230303202-2233320102123223-2303130111221002-0121232312221302-2331123322213111-0133212202300333-3100331331222203-0202130300231321"></a>

<a id="canonical-3213300322320312-2211230210120033-3201302131313322-3032302210111131-0111333023311322-2232303203313132-3133012233213130-2313130131223302"></a>

## route_table_id property — selective_route_tables / 002322120233 / 4

Type: `["list", "string"]`. Computed.

Route table ID. Route table ID.

Upstream description:

Route table ID.

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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_len": "64",
    "ves.io.schema.rules.repeated.items.string.pattern": "^(rtb-)([a-z0-9]{8}|[a-z0-9]{17})$",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_len": "64",
    "ves.io.schema.rules.repeated.items.string.pattern": "^(rtb-)([a-z0-9]{8}|[a-z0-9]{17})$",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1033230133020310-3320113302320121-0201023312333200-1102103310332300-2220302211110002-3022231230032200-1000022213102320-2012331131323021"></a>

## Next pages — selective_route_tables / 002322120233 / 5

- [aws_provider.aws_tgw_site.vpc_attachments.vpc_list.default_route](data-sources--cloud_connect--reference--group-001.md#canonical-0333233202010131-1000222030131223-0202030023000321-3132123300323333-2010021110110333-0120331213112021-0323310011103333-2021311200221130)
- [xcsh_cloud_connect](../data-sources/cloud_connect.md#canonical-2011032310310033-3311323122002031-2213302012003111-3033233110012013-2021203231210000-3221202000202111-0001331133100330-3301210132312000)

<a id="canonical-2331221100110300-1113031003120322-3010333302212321-1033202110023100-3311223310010132-0120300201330202-3323031111310023-1002112333320021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1203323013330021-1011020301102313-0110030332323130-0310213330210333-1000333103311111-2010111223203313-2220123123213102-1231201133313333"></a>

## aws_provider.aws_tgw_site.vpc_attachments.vpc_list.labels — labels / 323222120030 / 2

Breadcrumbs:

- [xcsh_cloud_connect](../data-sources/cloud_connect.md#canonical-2011032310310033-3311323122002031-2213302012003111-3033233110012013-2021203231210000-3221202000202111-0001331133100330-3301210132312000)
- [Property reference](data-sources--cloud_connect--reference--group-001.md#canonical-0222032211111313-0003010022101312-0310023330122333-0201122213032311-0122213101110201-2002020111110210-1221320313132123-3300313332231200)
- [aws_provider](data-sources--cloud_connect--reference--group-001.md#canonical-3322210313330233-3203313001001101-1033130101220000-1313031021321122-3011231212132102-2121333313020110-1123112301023233-1202303010121200)
- [aws_provider.aws_tgw_site](data-sources--cloud_connect--reference--group-001.md#canonical-1210201012310011-1202111032113023-2120321323312011-1101121121031011-0113230133300122-3231021212301121-1230133321330321-3230000102301202)
- [aws_provider.aws_tgw_site.vpc_attachments](data-sources--cloud_connect--reference--group-001.md#canonical-3200311203003333-3212000013100213-1102221330110123-2132202003301122-3202310210102320-3120022021203120-3110111203122330-1233233203001313)
- [aws_provider.aws_tgw_site.vpc_attachments.vpc_list](data-sources--cloud_connect--reference--group-001.md#canonical-3130311012000102-3303222030000301-2101203322100102-3320221211100031-0320011012231032-2322023202333123-1013132123020102-3330213131230103)
- aws_provider.aws_tgw_site.vpc_attachments.vpc_list.labels

<a id="canonical-1301323010121110-1131312032002230-3012021122010113-1232320303122212-0312233023020230-0013310302011112-2131111230222221-1211232220113012"></a>

Type: `"single"`. Computed.

Add labels for the VPC attachment. These labels can then be used in policies such as enhanced
firewall.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2022010022303221-1230100210232312-0032002002203313-0320122102302102-1202010230003201-0102010001111100-1121011203121111-1230030120111311"></a>

## Direct properties — labels / 323222120030 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3101202121130310-1000231232020212-1010312000323022-0121221101332232-0202313213021211-2130220220120132-3031020303112301-0023222210103122"></a>

## Next pages — labels / 323222120030 / 4

- [aws_provider.aws_tgw_site.vpc_attachments.vpc_list](data-sources--cloud_connect--reference--group-001.md#canonical-3130311012000102-3303222030000301-2101203322100102-3320221211100031-0320011012231032-2322023202333123-1013132123020102-3330213131230103)
- [xcsh_cloud_connect](../data-sources/cloud_connect.md#canonical-2011032310310033-3311323122002031-2213302012003111-3033233110012013-2021203231210000-3221202000202111-0001331133100330-3301210132312000)

<a id="canonical-2000103210333331-1332323303303110-0301003132012213-2120303213221223-2133323002201033-2233212101203122-0101202023103300-0301321331310131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2022003022002001-2211330300110220-0212303213120212-2022321032101121-0322210313030003-2022212102200003-2320201221322022-2033232101110233"></a>

## aws_provider.aws_tgw_site.vpc_attachments.vpc_list.manual_routing — manual_routing / 311321311133 / 2

Breadcrumbs:

- [xcsh_cloud_connect](../data-sources/cloud_connect.md#canonical-2011032310310033-3311323122002031-2213302012003111-3033233110012013-2021203231210000-3221202000202111-0001331133100330-3301210132312000)
- [Property reference](data-sources--cloud_connect--reference--group-001.md#canonical-0222032211111313-0003010022101312-0310023330122333-0201122213032311-0122213101110201-2002020111110210-1221320313132123-3300313332231200)
- [aws_provider](data-sources--cloud_connect--reference--group-001.md#canonical-3322210313330233-3203313001001101-1033130101220000-1313031021321122-3011231212132102-2121333313020110-1123112301023233-1202303010121200)
- [aws_provider.aws_tgw_site](data-sources--cloud_connect--reference--group-001.md#canonical-1210201012310011-1202111032113023-2120321323312011-1101121121031011-0113230133300122-3231021212301121-1230133321330321-3230000102301202)
- [aws_provider.aws_tgw_site.vpc_attachments](data-sources--cloud_connect--reference--group-001.md#canonical-3200311203003333-3212000013100213-1102221330110123-2132202003301122-3202310210102320-3120022021203120-3110111203122330-1233233203001313)
- [aws_provider.aws_tgw_site.vpc_attachments.vpc_list](data-sources--cloud_connect--reference--group-001.md#canonical-3130311012000102-3303222030000301-2101203322100102-3320221211100031-0320011012231032-2322023202333123-1013132123020102-3330213131230103)
- aws_provider.aws_tgw_site.vpc_attachments.vpc_list.manual_routing

<a id="canonical-2211322220313012-2023130030001222-3230220303023301-0020223230120100-1212010330010012-0311131310333212-2210123000121200-3303003330300323"></a>

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

<a id="canonical-3110000123000100-3100332030323102-1232212012322020-1213312001010031-3212033010111332-3331320331110233-3022211212211332-3000112011130301"></a>

## Direct properties — manual_routing / 311321311133 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1302001312332113-2133122310323120-0133220113122030-3003202103100121-1230233003103310-1111020221022201-1221133331100123-0133112112030201"></a>

## Next pages — manual_routing / 311321311133 / 4

- [aws_provider.aws_tgw_site.vpc_attachments.vpc_list](data-sources--cloud_connect--reference--group-001.md#canonical-3130311012000102-3303222030000301-2101203322100102-3320221211100031-0320011012231032-2322023202333123-1013132123020102-3330213131230103)
- [xcsh_cloud_connect](../data-sources/cloud_connect.md#canonical-2011032310310033-3311323122002031-2213302012003111-3033233110012013-2021203231210000-3221202000202111-0001331133100330-3301210132312000)

<a id="canonical-1310213230010202-0010122300232220-3030003101222331-3033212232021303-0201232322203122-3303121123112230-3233222203332331-2222210203131132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1321201102031302-3100003021133201-1132011011121022-2313221112120231-3110310311230332-3203301012320132-3200200020232002-1310031130112303"></a>

## azure_vnet_site — azure_vnet_site / 303300102133 / 2

Breadcrumbs:

- [xcsh_cloud_connect](../data-sources/cloud_connect.md#canonical-2011032310310033-3311323122002031-2213302012003111-3033233110012013-2021203231210000-3221202000202111-0001331133100330-3301210132312000)
- [Property reference](data-sources--cloud_connect--reference--group-001.md#canonical-0222032211111313-0003010022101312-0310023330122333-0201122213032311-0122213101110201-2002020111110210-1221320313132123-3300313332231200)
- azure_vnet_site

<a id="canonical-2303331131320220-1121203202332230-2002103033303200-1330122233203313-2233002333201203-2210330300013330-3112012031201022-0212312313033300"></a>

Type: `"single"`. Computed.

Azure VNet Site Type. Cloud Connect Azure VNet Site Type.

Upstream description:

Cloud Connect Azure VNet Site Type.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2333203323303232-3203302331130213-1113032232221233-2103133013031310-0211032302332111-0000031301321001-3001131101202330-2323111320121212"></a>

## Direct properties — azure_vnet_site / 303300102133 / 3

- [site](data-sources--cloud_connect--reference--group-001.md#canonical-3102021212202223-3220103301032310-2200030332333111-0033031323121002-3121001133300202-2031320131310303-0302133200031210-3102323202103310): complete subsection reference.

- [vnet_attachments](data-sources--cloud_connect--reference--group-001.md#canonical-0230320312132323-2200131002110130-0233330032303133-0331032222032113-2033113223103310-2310201323303211-0102210022233122-2210201210032231): complete subsection reference.

<a id="canonical-3023131313200332-2110111031213232-3212022330113310-2113130303033201-0133013131221021-3122211021230231-3100020001232013-2201221123023011"></a>

## Next pages — azure_vnet_site / 303300102133 / 4

- [azure_vnet_site.site](data-sources--cloud_connect--reference--group-001.md#canonical-3102021212202223-3220103301032310-2200030332333111-0033031323121002-3121001133300202-2031320131310303-0302133200031210-3102323202103310)
- [azure_vnet_site.vnet_attachments](data-sources--cloud_connect--reference--group-001.md#canonical-0230320312132323-2200131002110130-0233330032303133-0331032222032113-2033113223103310-2310201323303211-0102210022233122-2210201210032231)
- [Property reference](data-sources--cloud_connect--reference--group-001.md#canonical-0222032211111313-0003010022101312-0310023330122333-0201122213032311-0122213101110201-2002020111110210-1221320313132123-3300313332231200)
- [xcsh_cloud_connect](../data-sources/cloud_connect.md#canonical-2011032310310033-3311323122002031-2213302012003111-3033233110012013-2021203231210000-3221202000202111-0001331133100330-3301210132312000)

<a id="canonical-3102021212202223-3220103301032310-2200030332333111-0033031323121002-3121001133300202-2031320131310303-0302133200031210-3102323202103310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3010031232122111-3133030211021022-1033231230110232-2333131010110232-0022131020110021-0030203111010110-0321110002011003-1230100110203013"></a>

## azure_vnet_site.site — site / 030100033311 / 2

Breadcrumbs:

- [xcsh_cloud_connect](../data-sources/cloud_connect.md#canonical-2011032310310033-3311323122002031-2213302012003111-3033233110012013-2021203231210000-3221202000202111-0001331133100330-3301210132312000)
- [Property reference](data-sources--cloud_connect--reference--group-001.md#canonical-0222032211111313-0003010022101312-0310023330122333-0201122213032311-0122213101110201-2002020111110210-1221320313132123-3300313332231200)
- [azure_vnet_site](data-sources--cloud_connect--reference--group-001.md#canonical-1310213230010202-0010122300232220-3030003101222331-3033212232021303-0201232322203122-3303121123112230-3233222203332331-2222210203131132)
- azure_vnet_site.site

<a id="canonical-2030211203020302-2303220330131231-2132310220022100-0110121302033223-0012301010131310-3300300132231012-1000212321020331-3302233200123022"></a>

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

<a id="canonical-1211222113112003-2200213230213003-1132023221122221-3332213110031333-1220120113233211-1232203120220021-1113022302212333-1023231212011313"></a>

## Direct properties — site / 030100033311 / 3

<a id="canonical-1102020303211010-0011102030131312-3313213020002113-3201331333223201-2233201212100101-0230023220232002-2222020103111320-3312330220011012"></a>

<a id="canonical-1103110201102203-1000031233101200-1302221131132302-3012233010203023-2221113330320010-3220230332201323-2120310203330221-1131031110303101"></a>

## name property — site / 030100033311 / 4

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

<a id="canonical-1123101002330133-0100211213003113-3100331210303222-3300101131211310-0132122212123130-3133202321121230-1012030033333132-1033203010033331"></a>

<a id="canonical-0112120310122310-0321311002230230-1313032221211201-3303301032312213-3322201212212003-2113303132222002-0221301233002303-3031021003303223"></a>

## namespace property — site / 030100033311 / 5

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

<a id="canonical-0230203120301011-1310331130030032-1311123321321232-0033113331231222-1021023112100110-2300102011323003-3320333031232321-3232020302000021"></a>

<a id="canonical-2210122213213222-2131312021111022-0001233130312222-0100123023123320-0012312323310010-3223133113012222-1232033031302103-3132201102302011"></a>

## tenant property — site / 030100033311 / 6

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

<a id="canonical-1011001003012312-3212331311110321-2213131001032033-1022200301111312-2133222311311031-0100131223020233-0022032233021123-1312332011302103"></a>

## Next pages — site / 030100033311 / 7

- [azure_vnet_site](data-sources--cloud_connect--reference--group-001.md#canonical-1310213230010202-0010122300232220-3030003101222331-3033212232021303-0201232322203122-3303121123112230-3233222203332331-2222210203131132)
- [xcsh_cloud_connect](../data-sources/cloud_connect.md#canonical-2011032310310033-3311323122002031-2213302012003111-3033233110012013-2021203231210000-3221202000202111-0001331133100330-3301210132312000)

<a id="canonical-0230320312132323-2200131002110130-0233330032303133-0331032222032113-2033113223103310-2310201323303211-0102210022233122-2210201210032231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0100112112121322-2112002012123230-3213231031013022-1130113120021312-3133123330000031-3132202200131032-0312022213121003-3212311101310013"></a>

## azure_vnet_site.vnet_attachments — vnet_attachments / 302221133202 / 2

Breadcrumbs:

- [xcsh_cloud_connect](../data-sources/cloud_connect.md#canonical-2011032310310033-3311323122002031-2213302012003111-3033233110012013-2021203231210000-3221202000202111-0001331133100330-3301210132312000)
- [Property reference](data-sources--cloud_connect--reference--group-001.md#canonical-0222032211111313-0003010022101312-0310023330122333-0201122213032311-0122213101110201-2002020111110210-1221320313132123-3300313332231200)
- [azure_vnet_site](data-sources--cloud_connect--reference--group-001.md#canonical-1310213230010202-0010122300232220-3030003101222331-3033212232021303-0201232322203122-3303121123112230-3233222203332331-2222210203131132)
- azure_vnet_site.vnet_attachments

<a id="canonical-2200223322013231-3133121010213111-1011022302132032-2113023221120221-2132002221103103-2223222302332033-0221332000013030-3101203113330001"></a>

Type: `"single"`. Computed.

Configuration parameter for vnet attachments.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3313100011233211-3020121300213100-0303203222012210-0331033230233133-2112320020310101-0201033100102022-0303032020022232-3031112110231002"></a>

## Direct properties — vnet_attachments / 302221133202 / 3

- [vnet_list](data-sources--cloud_connect--reference--group-001.md#canonical-2031031013220233-3021013330231020-1303212211322333-3220331102233013-3000313130132021-1130102221322202-2020300300220203-2010133303333300): complete subsection reference.

<a id="canonical-2331332313300102-0302032123331220-0321211313213303-1112301002321102-0022011112313202-3133332033321033-2102220031201102-1031102133123022"></a>

## Next pages — vnet_attachments / 302221133202 / 4

- [azure_vnet_site.vnet_attachments.vnet_list](data-sources--cloud_connect--reference--group-001.md#canonical-2031031013220233-3021013330231020-1303212211322333-3220331102233013-3000313130132021-1130102221322202-2020300300220203-2010133303333300)
- [azure_vnet_site](data-sources--cloud_connect--reference--group-001.md#canonical-1310213230010202-0010122300232220-3030003101222331-3033212232021303-0201232322203122-3303121123112230-3233222203332331-2222210203131132)
- [xcsh_cloud_connect](../data-sources/cloud_connect.md#canonical-2011032310310033-3311323122002031-2213302012003111-3033233110012013-2021203231210000-3221202000202111-0001331133100330-3301210132312000)

<a id="canonical-2031031013220233-3021013330231020-1303212211322333-3220331102233013-3000313130132021-1130102221322202-2020300300220203-2010133303333300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3011322001133130-1123333101202120-3301211110100223-0203013130130333-3022012002221301-3033232002111020-0232021022320201-3113303133332312"></a>

## azure_vnet_site.vnet_attachments.vnet_list — vnet_list / 033010113021 / 2

Breadcrumbs:

- [xcsh_cloud_connect](../data-sources/cloud_connect.md#canonical-2011032310310033-3311323122002031-2213302012003111-3033233110012013-2021203231210000-3221202000202111-0001331133100330-3301210132312000)
- [Property reference](data-sources--cloud_connect--reference--group-001.md#canonical-0222032211111313-0003010022101312-0310023330122333-0201122213032311-0122213101110201-2002020111110210-1221320313132123-3300313332231200)
- [azure_vnet_site](data-sources--cloud_connect--reference--group-001.md#canonical-1310213230010202-0010122300232220-3030003101222331-3033212232021303-0201232322203122-3303121123112230-3233222203332331-2222210203131132)
- [azure_vnet_site.vnet_attachments](data-sources--cloud_connect--reference--group-001.md#canonical-0230320312132323-2200131002110130-0233330032303133-0331032222032113-2033113223103310-2310201323303211-0102210022233122-2210201210032231)
- azure_vnet_site.vnet_attachments.vnet_list

<a id="canonical-3003310002001023-3031001112033032-3023222232312120-2010022113001233-0021331032103003-1002330002103220-3311131233333331-3101112031331012"></a>

Type: `"list"`. Computed.

VNet List. Collection of items or values

Upstream description:

Collection of items or values

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 128,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 128,
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
    "ves.io.schema.rules.repeated.max_items": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "128"
  }
}
```

<a id="canonical-0131103133133002-0313131211030112-3001030300221131-1011301120210112-3300232030332020-2111213213120230-0012301222030203-0020221322113102"></a>

## Direct properties — vnet_list / 033010113021 / 3

- [custom_routing](data-sources--cloud_connect--reference--group-001.md#canonical-3023000011113031-3220300103221132-2031322323031310-2213300011110321-2130130030002322-1333121213113220-3323100123212200-1200200013310120): complete subsection reference.

- [default_route](data-sources--cloud_connect--reference--group-001.md#canonical-2102123111210300-0012120231222223-3002320210023113-3021222322000312-0001200233013131-0013131303130220-3031110011220123-1123232001321010): complete subsection reference.

- [labels](data-sources--cloud_connect--reference--group-001.md#canonical-0230311100303121-2102113133213300-2312332110032230-2310202220321320-1233123301020033-2101313013320320-0332231230332301-3133302223312101): complete subsection reference.

- [manual_routing](data-sources--cloud_connect--reference--group-001.md#canonical-3203001103101101-2100030111030200-3101020211011020-0021302321012010-2122212011333020-3010031231201131-0233301332312201-0333130210323012): complete subsection reference.

<a id="canonical-0220321320030211-3023202301032231-1102022223123030-3310132100132110-0021011330100021-3013131020111220-3223102133323132-2113031221220310"></a>

<a id="canonical-3331111221232300-1132020313233313-0130312120223122-1122202232202102-3200001323013003-2310221111122211-3021013132023003-1033231130111133"></a>

## subscription_id property — vnet_list / 033010113021 / 4

Type: `"string"`. Computed.

Enter the Subscription ID of the VNet to be attached.

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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64"
  }
}
```

<a id="canonical-0323302110032333-3001322223023311-0332333021212230-1212021230000211-1222310000333223-0302320103223332-0220302321101012-3220312311132311"></a>

<a id="canonical-0200111013311202-0200322013332300-1223311003022211-0030123333222333-1011320233320302-0320130112000311-1031131202032333-3320203130210101"></a>

## vnet_id property — vnet_list / 033010113021 / 5

Type: `"string"`. Computed.

Enter the VNet ID of the VNet to be attached in format
/&lt;resource-group-name&gt;/&lt;VNet-name&gt;.

Upstream description:

Enter the VNet ID of the VNet to be attached in format
/&lt;resource-group-name&gt;/&lt;VNet-name&gt;

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

<a id="canonical-3312320320122103-1021200030321003-0110301103333000-0110311332310010-2120120322011001-2121000332302023-1323013310123023-2230310002210031"></a>

## Next pages — vnet_list / 033010113021 / 6

- [azure_vnet_site.vnet_attachments.vnet_list.custom_routing](data-sources--cloud_connect--reference--group-001.md#canonical-3023000011113031-3220300103221132-2031322323031310-2213300011110321-2130130030002322-1333121213113220-3323100123212200-1200200013310120)
- [azure_vnet_site.vnet_attachments.vnet_list.default_route](data-sources--cloud_connect--reference--group-001.md#canonical-2102123111210300-0012120231222223-3002320210023113-3021222322000312-0001200233013131-0013131303130220-3031110011220123-1123232001321010)
- [azure_vnet_site.vnet_attachments.vnet_list.labels](data-sources--cloud_connect--reference--group-001.md#canonical-0230311100303121-2102113133213300-2312332110032230-2310202220321320-1233123301020033-2101313013320320-0332231230332301-3133302223312101)
- [azure_vnet_site.vnet_attachments.vnet_list.manual_routing](data-sources--cloud_connect--reference--group-001.md#canonical-3203001103101101-2100030111030200-3101020211011020-0021302321012010-2122212011333020-3010031231201131-0233301332312201-0333130210323012)
- [azure_vnet_site.vnet_attachments](data-sources--cloud_connect--reference--group-001.md#canonical-0230320312132323-2200131002110130-0233330032303133-0331032222032113-2033113223103310-2310201323303211-0102210022233122-2210201210032231)
- [xcsh_cloud_connect](../data-sources/cloud_connect.md#canonical-2011032310310033-3311323122002031-2213302012003111-3033233110012013-2021203231210000-3221202000202111-0001331133100330-3301210132312000)

<a id="canonical-3023000011113031-3220300103221132-2031322323031310-2213300011110321-2130130030002322-1333121213113220-3323100123212200-1200200013310120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3323223210313322-0303110230021103-3202102000030210-2102121022110031-1033201111020101-2023333212231301-0233320213332121-0203023310030313"></a>

## azure_vnet_site.vnet_attachments.vnet_list.custom_routing — custom_routing / 212231133223 / 2

Breadcrumbs:

- [xcsh_cloud_connect](../data-sources/cloud_connect.md#canonical-2011032310310033-3311323122002031-2213302012003111-3033233110012013-2021203231210000-3221202000202111-0001331133100330-3301210132312000)
- [Property reference](data-sources--cloud_connect--reference--group-001.md#canonical-0222032211111313-0003010022101312-0310023330122333-0201122213032311-0122213101110201-2002020111110210-1221320313132123-3300313332231200)
- [azure_vnet_site](data-sources--cloud_connect--reference--group-001.md#canonical-1310213230010202-0010122300232220-3030003101222331-3033212232021303-0201232322203122-3303121123112230-3233222203332331-2222210203131132)
- [azure_vnet_site.vnet_attachments](data-sources--cloud_connect--reference--group-001.md#canonical-0230320312132323-2200131002110130-0233330032303133-0331032222032113-2033113223103310-2310201323303211-0102210022233122-2210201210032231)
- [azure_vnet_site.vnet_attachments.vnet_list](data-sources--cloud_connect--reference--group-001.md#canonical-2031031013220233-3021013330231020-1303212211322333-3220331102233013-3000313130132021-1130102221322202-2020300300220203-2010133303333300)
- azure_vnet_site.vnet_attachments.vnet_list.custom_routing

<a id="canonical-1202221332023212-0021111010003202-0332121131112101-2133121213122232-2020012110303132-3303200101100002-2300330101113311-0211300303212322"></a>

Type: `"single"`. Computed.

List Azure Route Table with Static Route.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3302032031331323-3210312331032222-3102320032201202-2300022133113021-1022321210022231-3123130112300121-3022023020033022-3223200023313003"></a>

## Direct properties — custom_routing / 212231133223 / 3

- [route_tables](data-sources--cloud_connect--reference--group-001.md#canonical-2232200020030213-2233122101320313-0020210103330312-0131131212000202-2313322032131132-0101021311112010-0331323202330030-3023202020332330): complete subsection reference.

<a id="canonical-2333023301200312-2012122203111133-3003133000130121-2313200300320023-0021332311202322-2212131201331201-1031022322033232-2111230221022301"></a>

## Next pages — custom_routing / 212231133223 / 4

- [azure_vnet_site.vnet_attachments.vnet_list.custom_routing.route_tables](data-sources--cloud_connect--reference--group-001.md#canonical-2232200020030213-2233122101320313-0020210103330312-0131131212000202-2313322032131132-0101021311112010-0331323202330030-3023202020332330)
- [azure_vnet_site.vnet_attachments.vnet_list](data-sources--cloud_connect--reference--group-001.md#canonical-2031031013220233-3021013330231020-1303212211322333-3220331102233013-3000313130132021-1130102221322202-2020300300220203-2010133303333300)
- [xcsh_cloud_connect](../data-sources/cloud_connect.md#canonical-2011032310310033-3311323122002031-2213302012003111-3033233110012013-2021203231210000-3221202000202111-0001331133100330-3301210132312000)

<a id="canonical-2232200020030213-2233122101320313-0020210103330312-0131131212000202-2313322032131132-0101021311112010-0331323202330030-3023202020332330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3022322302033200-2222102131230021-1320301032332302-1011030302000210-1130111320101203-2323322020023323-3001030210333013-2222211322212001"></a>

## azure_vnet_site.vnet_attachments.vnet_list.custom_routing.route_tables — route_tables / 210120121133 / 2

Breadcrumbs:

- [xcsh_cloud_connect](../data-sources/cloud_connect.md#canonical-2011032310310033-3311323122002031-2213302012003111-3033233110012013-2021203231210000-3221202000202111-0001331133100330-3301210132312000)
- [Property reference](data-sources--cloud_connect--reference--group-001.md#canonical-0222032211111313-0003010022101312-0310023330122333-0201122213032311-0122213101110201-2002020111110210-1221320313132123-3300313332231200)
- [azure_vnet_site](data-sources--cloud_connect--reference--group-001.md#canonical-1310213230010202-0010122300232220-3030003101222331-3033212232021303-0201232322203122-3303121123112230-3233222203332331-2222210203131132)
- [azure_vnet_site.vnet_attachments](data-sources--cloud_connect--reference--group-001.md#canonical-0230320312132323-2200131002110130-0233330032303133-0331032222032113-2033113223103310-2310201323303211-0102210022233122-2210201210032231)
- [azure_vnet_site.vnet_attachments.vnet_list](data-sources--cloud_connect--reference--group-001.md#canonical-2031031013220233-3021013330231020-1303212211322333-3220331102233013-3000313130132021-1130102221322202-2020300300220203-2010133303333300)
- [azure_vnet_site.vnet_attachments.vnet_list.custom_routing](data-sources--cloud_connect--reference--group-001.md#canonical-3023000011113031-3220300103221132-2031322323031310-2213300011110321-2130130030002322-1333121213113220-3323100123212200-1200200013310120)
- azure_vnet_site.vnet_attachments.vnet_list.custom_routing.route_tables

<a id="canonical-0322202123202203-3320322122311122-3011131023001003-2011122010211101-1301001323000311-3300320003312003-0122122302223331-0132320320200013"></a>

Type: `"list"`. Computed.

List of route tables with static routes. Route Tables with static routes.

Upstream description:

Route Tables with static routes.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 200,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 200,
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
    "ves.io.schema.rules.repeated.max_items": "200",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "200",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2121203130101002-3123211131233102-2022323030021322-3310113311300101-3310010310220303-3032112232020020-0001113022031020-0213233020032103"></a>

## Direct properties — route_tables / 210120121133 / 3

<a id="canonical-3223020110011033-0133232003231230-0031311132200013-0323000031022320-2332012230133221-2112203031301101-3201013120200201-2010301211032020"></a>

<a id="canonical-3010010310222300-3120331120123101-1011322333102002-3100322112130233-3221030032230301-3302202100233012-2301022103103331-1102311033312203"></a>

## route_table_id property — route_tables / 210120121133 / 4

Type: `"string"`. Computed.

Route table ID in the format /&lt;resource-group-name&gt;/&lt;route-table-name&gt;.

Upstream description:

Route table ID in the format /&lt;resource-group-name&gt;/&lt;route-table-name&gt;

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
    "pattern": "^\\\\/[-\\\\w\\\\._\\\\(\\\\)]+\\\\/[a-zA-Z0-9][a-zA-Z0-9-._]+[a-zA-Z0-9_]$"
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.pattern": "^\\\\/[-\\\\w\\\\._\\\\(\\\\)]+\\\\/[a-zA-Z0-9][a-zA-Z0-9-._]+[a-zA-Z0-9_]$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.pattern": "^\\\\/[-\\\\w\\\\._\\\\(\\\\)]+\\\\/[a-zA-Z0-9][a-zA-Z0-9-._]+[a-zA-Z0-9_]$"
  }
}
```

<a id="canonical-2031120031312132-0033232203303033-1303013331322211-1302302001213013-1210130110212201-3011232233112323-2201010021332101-0212231120332221"></a>

<a id="canonical-2001221331223301-0012022320000201-3310203333220133-2220031011030201-1312132203011233-3112323201210103-3210313313331012-1122112033210002"></a>

## static_routes property — route_tables / 210120121133 / 5

Type: `["list", "string"]`. Computed.

Static Routes. List of Static Routes.

Upstream description:

List of Static Routes.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 50,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 50,
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
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "50",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "50",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0203002013100210-0212303321222202-0123100322130201-0320233210112201-1113000111031230-0010002220210201-3213330021221000-3010133010323322"></a>

## Next pages — route_tables / 210120121133 / 6

- [azure_vnet_site.vnet_attachments.vnet_list.custom_routing](data-sources--cloud_connect--reference--group-001.md#canonical-3023000011113031-3220300103221132-2031322323031310-2213300011110321-2130130030002322-1333121213113220-3323100123212200-1200200013310120)
- [xcsh_cloud_connect](../data-sources/cloud_connect.md#canonical-2011032310310033-3311323122002031-2213302012003111-3033233110012013-2021203231210000-3221202000202111-0001331133100330-3301210132312000)

<a id="canonical-2102123111210300-0012120231222223-3002320210023113-3021222322000312-0001200233013131-0013131303130220-3031110011220123-1123232001321010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3023012010030332-3023101322023323-1312311120031122-0123123132201022-1230112001002220-2013320302332331-2020113010001102-0302313312021131"></a>

## azure_vnet_site.vnet_attachments.vnet_list.default_route — default_route / 032003120321 / 2

Breadcrumbs:

- [xcsh_cloud_connect](../data-sources/cloud_connect.md#canonical-2011032310310033-3311323122002031-2213302012003111-3033233110012013-2021203231210000-3221202000202111-0001331133100330-3301210132312000)
- [Property reference](data-sources--cloud_connect--reference--group-001.md#canonical-0222032211111313-0003010022101312-0310023330122333-0201122213032311-0122213101110201-2002020111110210-1221320313132123-3300313332231200)
- [azure_vnet_site](data-sources--cloud_connect--reference--group-001.md#canonical-1310213230010202-0010122300232220-3030003101222331-3033212232021303-0201232322203122-3303121123112230-3233222203332331-2222210203131132)
- [azure_vnet_site.vnet_attachments](data-sources--cloud_connect--reference--group-001.md#canonical-0230320312132323-2200131002110130-0233330032303133-0331032222032113-2033113223103310-2310201323303211-0102210022233122-2210201210032231)
- [azure_vnet_site.vnet_attachments.vnet_list](data-sources--cloud_connect--reference--group-001.md#canonical-2031031013220233-3021013330231020-1303212211322333-3220331102233013-3000313130132021-1130102221322202-2020300300220203-2010133303333300)
- azure_vnet_site.vnet_attachments.vnet_list.default_route

<a id="canonical-2230222032302001-3202100312302000-3030232332233133-0302122230331133-2311023110302213-2302031112200203-1122020301131101-2320100031310112"></a>

Type: `"single"`. Computed.

Configuration parameter for default route.

Upstream description:

Select Override Default Route Choice.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-default_route_choice": "[\"all_route_tables\",\"selective_route_tables\"]"
}
```

<a id="canonical-2130311102323001-2133120010302203-0032302030322022-1132033330223011-3121321122013121-2102103333120322-0332011220011231-3212332310231233"></a>

## Direct properties — default_route / 032003120321 / 3

- [all_route_tables](data-sources--cloud_connect--reference--group-001.md#canonical-1213133112330221-0301001002031300-1011122113010112-1203221001320223-2202001013102230-0113101102102301-3023313121010021-3102030210122211): complete subsection reference.

- [selective_route_tables](data-sources--cloud_connect--reference--group-001.md#canonical-0020001020221203-3323031110111102-3223132120232113-3031230333001003-2022221320110213-0220331100120123-3332110302231002-1322333311212331): complete subsection reference.

<a id="canonical-0100330332030332-3303022223012201-2211130002321113-0133120002100132-3020323310320101-1310032322123100-1332022000113012-3120133312223011"></a>

## Next pages — default_route / 032003120321 / 4

- [azure_vnet_site.vnet_attachments.vnet_list.default_route.all_route_tables](data-sources--cloud_connect--reference--group-001.md#canonical-1213133112330221-0301001002031300-1011122113010112-1203221001320223-2202001013102230-0113101102102301-3023313121010021-3102030210122211)
- [azure_vnet_site.vnet_attachments.vnet_list.default_route.selective_route_tables](data-sources--cloud_connect--reference--group-001.md#canonical-0020001020221203-3323031110111102-3223132120232113-3031230333001003-2022221320110213-0220331100120123-3332110302231002-1322333311212331)
- [azure_vnet_site.vnet_attachments.vnet_list](data-sources--cloud_connect--reference--group-001.md#canonical-2031031013220233-3021013330231020-1303212211322333-3220331102233013-3000313130132021-1130102221322202-2020300300220203-2010133303333300)
- [xcsh_cloud_connect](../data-sources/cloud_connect.md#canonical-2011032310310033-3311323122002031-2213302012003111-3033233110012013-2021203231210000-3221202000202111-0001331133100330-3301210132312000)

<a id="canonical-1213133112330221-0301001002031300-1011122113010112-1203221001320223-2202001013102230-0113101102102301-3023313121010021-3102030210122211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3221212213200221-3210210123021020-2223000123120033-1001313323322222-2320232321213123-1230212231021131-1113301100130102-3312001232331023"></a>

## azure_vnet_site.vnet_attachments.vnet_list.default_route.all_route_tables — all_route_tables / 321102301130 / 2

Breadcrumbs:

- [xcsh_cloud_connect](../data-sources/cloud_connect.md#canonical-2011032310310033-3311323122002031-2213302012003111-3033233110012013-2021203231210000-3221202000202111-0001331133100330-3301210132312000)
- [Property reference](data-sources--cloud_connect--reference--group-001.md#canonical-0222032211111313-0003010022101312-0310023330122333-0201122213032311-0122213101110201-2002020111110210-1221320313132123-3300313332231200)
- [azure_vnet_site](data-sources--cloud_connect--reference--group-001.md#canonical-1310213230010202-0010122300232220-3030003101222331-3033212232021303-0201232322203122-3303121123112230-3233222203332331-2222210203131132)
- [azure_vnet_site.vnet_attachments](data-sources--cloud_connect--reference--group-001.md#canonical-0230320312132323-2200131002110130-0233330032303133-0331032222032113-2033113223103310-2310201323303211-0102210022233122-2210201210032231)
- [azure_vnet_site.vnet_attachments.vnet_list](data-sources--cloud_connect--reference--group-001.md#canonical-2031031013220233-3021013330231020-1303212211322333-3220331102233013-3000313130132021-1130102221322202-2020300300220203-2010133303333300)
- [azure_vnet_site.vnet_attachments.vnet_list.default_route](data-sources--cloud_connect--reference--group-001.md#canonical-2102123111210300-0012120231222223-3002320210023113-3021222322000312-0001200233013131-0013131303130220-3031110011220123-1123232001321010)
- azure_vnet_site.vnet_attachments.vnet_list.default_route.all_route_tables

<a id="canonical-0231123112110200-3020001130310300-0110120210011323-3000312013320332-0212223012023303-3033220331321010-1001002231230332-0310122230300221"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for all route tables.

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

<a id="canonical-1322111233110100-2301011311002203-1021323003201200-2030132102000123-0130022013323233-0310133122201202-0120111310203012-0300123330010302"></a>

## Direct properties — all_route_tables / 321102301130 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2322010122232032-2330133030310002-0220332223200122-1102211033022001-3303233231102102-2030033222213222-1031221230311110-0313333221103112"></a>

## Next pages — all_route_tables / 321102301130 / 4

- [azure_vnet_site.vnet_attachments.vnet_list.default_route](data-sources--cloud_connect--reference--group-001.md#canonical-2102123111210300-0012120231222223-3002320210023113-3021222322000312-0001200233013131-0013131303130220-3031110011220123-1123232001321010)
- [xcsh_cloud_connect](../data-sources/cloud_connect.md#canonical-2011032310310033-3311323122002031-2213302012003111-3033233110012013-2021203231210000-3221202000202111-0001331133100330-3301210132312000)

<a id="canonical-0020001020221203-3323031110111102-3223132120232113-3031230333001003-2022221320110213-0220331100120123-3332110302231002-1322333311212331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1001220032331103-1021211331213200-2302301211310130-2021223122213331-3002133213101221-3223031120321002-2013112132222122-1032202002010200"></a>

## azure_vnet_site.vnet_attachments.vnet_list.default_route.selective_route_tables — selective_route_tables / 122002122030 / 2

Breadcrumbs:

- [xcsh_cloud_connect](../data-sources/cloud_connect.md#canonical-2011032310310033-3311323122002031-2213302012003111-3033233110012013-2021203231210000-3221202000202111-0001331133100330-3301210132312000)
- [Property reference](data-sources--cloud_connect--reference--group-001.md#canonical-0222032211111313-0003010022101312-0310023330122333-0201122213032311-0122213101110201-2002020111110210-1221320313132123-3300313332231200)
- [azure_vnet_site](data-sources--cloud_connect--reference--group-001.md#canonical-1310213230010202-0010122300232220-3030003101222331-3033212232021303-0201232322203122-3303121123112230-3233222203332331-2222210203131132)
- [azure_vnet_site.vnet_attachments](data-sources--cloud_connect--reference--group-001.md#canonical-0230320312132323-2200131002110130-0233330032303133-0331032222032113-2033113223103310-2310201323303211-0102210022233122-2210201210032231)
- [azure_vnet_site.vnet_attachments.vnet_list](data-sources--cloud_connect--reference--group-001.md#canonical-2031031013220233-3021013330231020-1303212211322333-3220331102233013-3000313130132021-1130102221322202-2020300300220203-2010133303333300)
- [azure_vnet_site.vnet_attachments.vnet_list.default_route](data-sources--cloud_connect--reference--group-001.md#canonical-2102123111210300-0012120231222223-3002320210023113-3021222322000312-0001200233013131-0013131303130220-3031110011220123-1123232001321010)
- azure_vnet_site.vnet_attachments.vnet_list.default_route.selective_route_tables

<a id="canonical-1213112003003022-0113113333201230-3313032110303120-2013110223231233-3210020321220013-0003101222313030-3312122022322302-1322012310221220"></a>

Type: `"single"`. Computed.

Configuration parameter for selective route tables.

Upstream description:

Azure Route Table.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2020003130013322-2113300322131303-1333310201110310-3213121321132012-2322310320333031-3322322321113112-3131323332200010-0130113203003123"></a>

## Direct properties — selective_route_tables / 122002122030 / 3

<a id="canonical-1333111113123110-0013232011000233-2211001221022032-2031103132010131-3112200310121113-2201321201122231-3020102213311311-3031332222312213"></a>

<a id="canonical-3002221133333111-2110311021002330-2032302231310110-3231103120213210-0323120223312000-1322221310302032-2302232002232311-1330232103330303"></a>

## route_table_id property — selective_route_tables / 122002122030 / 4

Type: `["list", "string"]`. Computed.

Route table ID in the format /&lt;resource-group-name&gt;/&lt;route-table-name&gt;.

Upstream description:

Route table ID in the format /&lt;resource-group-name&gt;/&lt;route-table-name&gt;

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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_len": "256",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_len": "256",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2110001201330110-0102030321023331-3013003011011121-2302202001132320-3013321330222200-0102322212311211-0230122330012103-3110200311132333"></a>

## Next pages — selective_route_tables / 122002122030 / 5

- [azure_vnet_site.vnet_attachments.vnet_list.default_route](data-sources--cloud_connect--reference--group-001.md#canonical-2102123111210300-0012120231222223-3002320210023113-3021222322000312-0001200233013131-0013131303130220-3031110011220123-1123232001321010)
- [xcsh_cloud_connect](../data-sources/cloud_connect.md#canonical-2011032310310033-3311323122002031-2213302012003111-3033233110012013-2021203231210000-3221202000202111-0001331133100330-3301210132312000)

<a id="canonical-0230311100303121-2102113133213300-2312332110032230-2310202220321320-1233123301020033-2101313013320320-0332231230332301-3133302223312101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1021130301312303-0310110203331232-2011220212102332-1301001110030012-1323201320001311-1032332201131200-0232232023031220-3233011303102301"></a>

## azure_vnet_site.vnet_attachments.vnet_list.labels — labels / 213003313120 / 2

Breadcrumbs:

- [xcsh_cloud_connect](../data-sources/cloud_connect.md#canonical-2011032310310033-3311323122002031-2213302012003111-3033233110012013-2021203231210000-3221202000202111-0001331133100330-3301210132312000)
- [Property reference](data-sources--cloud_connect--reference--group-001.md#canonical-0222032211111313-0003010022101312-0310023330122333-0201122213032311-0122213101110201-2002020111110210-1221320313132123-3300313332231200)
- [azure_vnet_site](data-sources--cloud_connect--reference--group-001.md#canonical-1310213230010202-0010122300232220-3030003101222331-3033212232021303-0201232322203122-3303121123112230-3233222203332331-2222210203131132)
- [azure_vnet_site.vnet_attachments](data-sources--cloud_connect--reference--group-001.md#canonical-0230320312132323-2200131002110130-0233330032303133-0331032222032113-2033113223103310-2310201323303211-0102210022233122-2210201210032231)
- [azure_vnet_site.vnet_attachments.vnet_list](data-sources--cloud_connect--reference--group-001.md#canonical-2031031013220233-3021013330231020-1303212211322333-3220331102233013-3000313130132021-1130102221322202-2020300300220203-2010133303333300)
- azure_vnet_site.vnet_attachments.vnet_list.labels

<a id="canonical-1023212103223232-0131121001222212-0312023302021112-0301132001221032-0210310233210332-3202210301133332-0022021110011303-1100133130311011"></a>

Type: `"single"`. Computed.

Add labels for the VNet attachments. These labels can then be used in policies such as enhanced
firewall policies.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2202002323020303-2310322221023320-0023002002211301-1102012022233112-3100221223200300-0210020330102113-3301331211233101-3323000231101022"></a>

## Direct properties — labels / 213003313120 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3200301322012222-1013232320213112-2222132211020300-2122332020213113-3001220210213222-3211101122231313-2030223003302333-2123312310113203"></a>

## Next pages — labels / 213003313120 / 4

- [azure_vnet_site.vnet_attachments.vnet_list](data-sources--cloud_connect--reference--group-001.md#canonical-2031031013220233-3021013330231020-1303212211322333-3220331102233013-3000313130132021-1130102221322202-2020300300220203-2010133303333300)
- [xcsh_cloud_connect](../data-sources/cloud_connect.md#canonical-2011032310310033-3311323122002031-2213302012003111-3033233110012013-2021203231210000-3221202000202111-0001331133100330-3301210132312000)

<a id="canonical-3203001103101101-2100030111030200-3101020211011020-0021302321012010-2122212011333020-3010031231201131-0233301332312201-0333130210323012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3223231133121232-3103123023133122-3200132211330221-2201332013133000-2000331223222001-1223122133123121-1120310233320133-0023121020322232"></a>

## azure_vnet_site.vnet_attachments.vnet_list.manual_routing — manual_routing / 220220113100 / 2

Breadcrumbs:

- [xcsh_cloud_connect](../data-sources/cloud_connect.md#canonical-2011032310310033-3311323122002031-2213302012003111-3033233110012013-2021203231210000-3221202000202111-0001331133100330-3301210132312000)
- [Property reference](data-sources--cloud_connect--reference--group-001.md#canonical-0222032211111313-0003010022101312-0310023330122333-0201122213032311-0122213101110201-2002020111110210-1221320313132123-3300313332231200)
- [azure_vnet_site](data-sources--cloud_connect--reference--group-001.md#canonical-1310213230010202-0010122300232220-3030003101222331-3033212232021303-0201232322203122-3303121123112230-3233222203332331-2222210203131132)
- [azure_vnet_site.vnet_attachments](data-sources--cloud_connect--reference--group-001.md#canonical-0230320312132323-2200131002110130-0233330032303133-0331032222032113-2033113223103310-2310201323303211-0102210022233122-2210201210032231)
- [azure_vnet_site.vnet_attachments.vnet_list](data-sources--cloud_connect--reference--group-001.md#canonical-2031031013220233-3021013330231020-1303212211322333-3220331102233013-3000313130132021-1130102221322202-2020300300220203-2010133303333300)
- azure_vnet_site.vnet_attachments.vnet_list.manual_routing

<a id="canonical-1310113022310312-1130103101120312-0303132100020203-0213110130311020-0123122320100230-3030331101022321-3020031113100220-1112233320020031"></a>

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

<a id="canonical-0020122111331102-0002000003321030-3331003121222110-3100001210022213-1021211121010302-3300330203313000-1231233011122232-2222113223031323"></a>

## Direct properties — manual_routing / 220220113100 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0330300232213233-1331331110200102-3003123022312011-1101022123332201-0211103010133132-3010210013031233-2021131111231102-1333200030210202"></a>

## Next pages — manual_routing / 220220113100 / 4

- [azure_vnet_site.vnet_attachments.vnet_list](data-sources--cloud_connect--reference--group-001.md#canonical-2031031013220233-3021013330231020-1303212211322333-3220331102233013-3000313130132021-1130102221322202-2020300300220203-2010133303333300)
- [xcsh_cloud_connect](../data-sources/cloud_connect.md#canonical-2011032310310033-3311323122002031-2213302012003111-3033233110012013-2021203231210000-3221202000202111-0001331133100330-3301210132312000)

<a id="canonical-0231311031231003-1321013301003130-3313203023221333-0013202013320101-2320200121101002-2203011312223323-1203000323320122-1200110112102001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0211021312203321-3312132133123013-2112121132200120-0023300000333010-3302223131123331-1313203021002103-0212031230133231-2302200022201222"></a>

## segment — segment / 001220102221 / 2

Breadcrumbs:

- [xcsh_cloud_connect](../data-sources/cloud_connect.md#canonical-2011032310310033-3311323122002031-2213302012003111-3033233110012013-2021203231210000-3221202000202111-0001331133100330-3301210132312000)
- [Property reference](data-sources--cloud_connect--reference--group-001.md#canonical-0222032211111313-0003010022101312-0310023330122333-0201122213032311-0122213101110201-2002020111110210-1221320313132123-3300313332231200)
- segment

<a id="canonical-1321211200101101-3231233333312303-1023333013331310-2312022030301213-2122012103101310-0112330331030330-2210330113331103-0220230112321012"></a>

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

<a id="canonical-3002322021231123-2212113312022110-2031031130010213-1230021121032022-2233012131211032-0001310230113102-0320221000202223-2022003233202133"></a>

## Direct properties — segment / 001220102221 / 3

<a id="canonical-1120103033202211-1312123030321303-2121112230310001-0313102102323020-0231032203133220-1223211211033102-1102210222303132-1223222011101103"></a>

<a id="canonical-0113312310100221-0100112023132223-1220213223320022-1112123121130123-3330313112321210-3310003210303010-2322032103122102-1010112330310130"></a>

## name property — segment / 001220102221 / 4

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

<a id="canonical-2131121001302133-2210232310322021-2212130023331232-2010213202000112-1023020323103002-3200313013202220-3202212133222302-2200320031010213"></a>

<a id="canonical-3222110213321133-0032300101002333-3132300121312212-2301120113003101-0022202312023212-2012030221012321-3100101230022011-0311320002112330"></a>

## namespace property — segment / 001220102221 / 5

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

<a id="canonical-3012212122202301-3333000220102202-3201010032223100-0131231003220021-3231300302102303-1133233003213230-0303100232233313-3332110102203222"></a>

<a id="canonical-0102123120110120-2021122300110203-3100331102312212-1132123233030013-1030333301320201-0210322201331020-0232323212301110-0010222333130223"></a>

## tenant property — segment / 001220102221 / 6

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

<a id="canonical-2210013102110130-1301311311111323-3110330322200330-0000223223221300-2213003111122000-0120322310132002-2103232002203301-0200002022323331"></a>

## Next pages — segment / 001220102221 / 7

- [Property reference](data-sources--cloud_connect--reference--group-001.md#canonical-0222032211111313-0003010022101312-0310023330122333-0201122213032311-0122213101110201-2002020111110210-1221320313132123-3300313332231200)
- [xcsh_cloud_connect](../data-sources/cloud_connect.md#canonical-2011032310310033-3311323122002031-2213302012003111-3033233110012013-2021203231210000-3221202000202111-0001331133100330-3301210132312000)
