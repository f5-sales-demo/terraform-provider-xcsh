---
page_title: "xcsh_nat_policy reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_nat_policy reference."
---

# xcsh_nat_policy reference

<a id="canonical-0313000023303013-1312311300100203-2312121103033010-0011202012203121-2320110020132021-0013320012122120-1332300232020201-2300013013202121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Property reference

Breadcrumbs:

- [xcsh_nat_policy](../data-sources/nat_policy.md#canonical-2211333230110221-1023302013201300-2020112103300231-1033030223223010-0020003220130132-3323021301233201-1030131302223322-2302033131310121)
- Property reference

<a id="canonical-2311022131301122-0331110301022202-3102202111021201-3010101033002002-0312230032001233-1212223221133100-2020213321103213-0110220231333110"></a>

### Direct properties for `xcsh_nat_policy`

<a id="canonical-2322210113220111-2330122330121131-0201111213111132-1223332201321102-0333012303201001-3123033223003231-1101220210001112-2303331133003130"></a>

#### `annotations` property

Type: `["map", "string"]`. Computed.

Annotations applied to this resource.

Additional upstream details:

Annotations is an unstructured key-value map stored with a resource that may be set by external
tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when
modifying objects.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "map",
    "deterministic": true,
    "keys": {
      "maxLength": 64,
      "minLength": 1,
      "type": "string"
    },
    "originalRules": {
      "ves.io.schema.rules.map.keys.string.max_len": "64",
      "ves.io.schema.rules.map.keys.string.min_len": "1",
      "ves.io.schema.rules.map.values.string.max_len": "1024",
      "ves.io.schema.rules.map.values.string.min_len": "1"
    },
    "values": {
      "maxLength": 1024,
      "minLength": 1,
      "type": "string"
    }
  },
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

<a id="canonical-3001021101030032-2033110333000120-0303230021301122-3302031132223100-0320310102321003-0130010133312222-2323322121302103-3301032223212102"></a>

<a id="canonical-1301130120211103-3031131323100221-1012212331002210-1033122033322321-1120331322332112-0203123113030323-0033320131330321-2122312100031021"></a>

#### `description` property

Type: `"string"`. Computed.

Description of the NATPolicy.

Additional upstream details:

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-1200303123303110-2323230001012100-3231130123310103-2213123100220323-1121131330113133-0311002001212232-0203300003001331-2022210221203121"></a>

<a id="canonical-3000100121012302-2321313111010300-2213113212202322-1021310301312230-2220213002232333-3231113013020120-2112210131322130-0123012313013301"></a>

#### `id` property

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-2311203102200033-3330032230102203-0332031033030202-1122313330211021-0020320022230333-1211102302030212-3231310122101333-1032103000031212"></a>

<a id="canonical-2022001020013013-2230030313001011-0313200330011303-2111001223120223-0103013333222021-1211223013133303-1120013001333131-3130321321122220"></a>

#### `labels` property

Type: `["map", "string"]`. Computed.

Labels applied to this resource.

Additional upstream details:

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

<a id="canonical-3033332001313123-1120123222001031-2313010320100023-1030002303133101-0110013221130022-0002120010113002-1323300001320001-3101312203130221"></a>

<a id="canonical-3003103322022212-0333021121213022-1103033133311103-0102312200033310-3232333311010012-2112203101133110-1031323111332123-2122210000300023"></a>

#### `name` property

Type: `"string"`. Required.

Name of the NATPolicy.

Additional upstream details:

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-1312210222202110-3222110211313200-1313101023133031-1300021300121003-3110232031012131-3221222311201330-2021320221032221-2213003010011001"></a>

<a id="canonical-2323222111232121-0300102311110210-1012330022333131-3022310130113331-1132103132333321-1011002322303311-2331121333313032-3312133113001202"></a>

#### `namespace` property

Type: `"string"`. Required.

Namespace where the NATPolicy exists.

Additional upstream details:

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

- [rules](data-sources--nat_policy--reference--group-001.md#canonical-0211110210131101-2303300303122112-0031321323203023-2003313333331231-3032223122221033-1121322010202212-2233233201200212-2000311112203230): complete subsection reference.

- [site](data-sources--nat_policy--reference--group-001.md#canonical-1011301221111022-3101333210103002-3321113303030012-1132311110323211-0123203232212231-0220003002111330-0230131113310100-0023212320213110): complete subsection reference.

<a id="canonical-3202012332302201-2211301201020202-0203110000233200-2100130203111132-1030203203213033-1201333333213012-0131101222112230-0000003010323311"></a>

### All schema paths for `xcsh_nat_policy`

Each exact path has one authoritative reference destination. Collection element indices are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](data-sources--nat_policy--reference--group-001.md#canonical-2322210113220111-2330122330121131-0201111213111132-1223332201321102-0333012303201001-3123033223003231-1101220210001112-2303331133003130) |
| `description` | [description](data-sources--nat_policy--reference--group-001.md#canonical-3001021101030032-2033110333000120-0303230021301122-3302031132223100-0320310102321003-0130010133312222-2323322121302103-3301032223212102) |
| `id` | [ID](data-sources--nat_policy--reference--group-001.md#canonical-1200303123303110-2323230001012100-3231130123310103-2213123100220323-1121131330113133-0311002001212232-0203300003001331-2022210221203121) |
| `labels` | [labels](data-sources--nat_policy--reference--group-001.md#canonical-2311203102200033-3330032230102203-0332031033030202-1122313330211021-0020320022230333-1211102302030212-3231310122101333-1032103000031212) |
| `name` | [name](data-sources--nat_policy--reference--group-001.md#canonical-3033332001313123-1120123222001031-2313010320100023-1030002303133101-0110013221130022-0002120010113002-1323300001320001-3101312203130221) |
| `namespace` | [namespace](data-sources--nat_policy--reference--group-001.md#canonical-1312210222202110-3222110211313200-1313101023133031-1300021300121003-3110232031012131-3221222311201330-2021320221032221-2213003010011001) |
| `rules` | [rules](data-sources--nat_policy--reference--group-001.md#canonical-3122332210112211-2003013132131222-1030202323331303-0032210223011333-1012312202301012-3013130121121002-2001032112100132-2121111303320121) |
| `rules.action` | [rules.action](data-sources--nat_policy--reference--group-001.md#canonical-3130131002322203-1211222203233032-1311300110230122-3012121030120222-3113231211122321-1131020320012101-0120103311111133-0002212303220220) |
| `rules.action.dynamic` | [rules.action.dynamic](data-sources--nat_policy--reference--group-001.md#canonical-1213113100011313-0132133001132220-0002101200320033-3300232131323301-3332021133232020-2323302301131332-1101022103021220-3112233331212100) |
| `rules.action.dynamic.elastic_ips` | [rules.action.dynamic.elastic_ips](data-sources--nat_policy--reference--group-001.md#canonical-1303212022310122-2132110230030121-3203222110200103-0220230312310001-3330031311023220-2222101132232220-3122331330120020-1032231103200132) |
| `rules.action.dynamic.elastic_ips.refs` | [rules.action.dynamic.elastic_ips.refs](data-sources--nat_policy--reference--group-001.md#canonical-0321102202111022-2313022011111211-2232130220031023-2330033230120113-2303030012312202-2311033020112001-2231233323123212-1023232133101130) |
| `rules.action.dynamic.elastic_ips.refs.kind` | [rules.action.dynamic.elastic_ips.refs.kind](data-sources--nat_policy--reference--group-001.md#canonical-2201220011301211-0313213210202201-2022211012211133-2122013100120231-0311300112300000-2122103000311221-1003131320011030-0322302032031323) |
| `rules.action.dynamic.elastic_ips.refs.name` | [rules.action.dynamic.elastic_ips.refs.name](data-sources--nat_policy--reference--group-001.md#canonical-0120022133113213-1320030302333110-3131320121000130-3101113300113023-1121130223113301-0101133103100231-2031221102032020-1022310131222123) |
| `rules.action.dynamic.elastic_ips.refs.namespace` | [rules.action.dynamic.elastic_ips.refs.namespace](data-sources--nat_policy--reference--group-001.md#canonical-0012120123222202-2221332102330311-1123320010220302-0033020123023200-2103211130332321-3131022320331312-1023030320031232-3203231302232032) |
| `rules.action.dynamic.elastic_ips.refs.tenant` | [rules.action.dynamic.elastic_ips.refs.tenant](data-sources--nat_policy--reference--group-001.md#canonical-2020230201301221-0220300233120203-0010111130033031-1202201133033113-2230031330122221-0100033012030031-3220221020131332-0211123033301223) |
| `rules.action.dynamic.elastic_ips.refs.uid` | [rules.action.dynamic.elastic_ips.refs.uid](data-sources--nat_policy--reference--group-001.md#canonical-2010231330101021-1010300031332222-0102213330032123-2323313110133002-2030103331230331-0031320020301131-0213330022320300-3302120202331211) |
| `rules.action.dynamic.pools` | [rules.action.dynamic.pools](data-sources--nat_policy--reference--group-001.md#canonical-3031213031022223-1231011030220200-1331012132312213-3122310012331223-0120301313031111-2023112303331232-3032303011112302-0101122112233133) |
| `rules.action.dynamic.pools.prefixes` | [rules.action.dynamic.pools.prefixes](data-sources--nat_policy--reference--group-001.md#canonical-1021122331203000-3323323223130113-1011323222110212-0332212022130212-2310231013111120-3002230000130121-0030203133223301-1002202210112003) |
| `rules.action.virtual_cidr` | [rules.action.virtual_cidr](data-sources--nat_policy--reference--group-001.md#canonical-0102231110000302-0232113333223222-2112213202000012-1020210323312223-3231300130110010-2103312032320301-1213001001232020-1033233110000331) |
| `rules.cloud_connect` | [rules.cloud_connect](data-sources--nat_policy--reference--group-001.md#canonical-3322312213223330-3332113132330201-2220033320302013-3113220001302332-3010130303130132-3200022302100220-1313310101302111-2301223233223102) |
| `rules.cloud_connect.refs` | [rules.cloud_connect.refs](data-sources--nat_policy--reference--group-001.md#canonical-1332133011313013-2320220332233232-1032312101321121-1322122323323023-1132300020021311-0202110203123202-3021111000003301-1323322223103232) |
| `rules.cloud_connect.refs.kind` | [rules.cloud_connect.refs.kind](data-sources--nat_policy--reference--group-001.md#canonical-0132021312232312-1323130331300112-1223100320211233-1010320003011312-3111120101313102-2233110121221311-3133232231030300-0120121121322121) |
| `rules.cloud_connect.refs.name` | [rules.cloud_connect.refs.name](data-sources--nat_policy--reference--group-001.md#canonical-1110333310011323-0111211113210211-2133013000131133-0322230332112123-0013310321232212-0112212231312111-1331030300130223-1212113020230222) |
| `rules.cloud_connect.refs.namespace` | [rules.cloud_connect.refs.namespace](data-sources--nat_policy--reference--group-001.md#canonical-2133330303203233-0003300120213102-3322220303201333-0013323031123213-0111103231212023-2212100111120221-3221020231103122-2000121200000203) |
| `rules.cloud_connect.refs.tenant` | [rules.cloud_connect.refs.tenant](data-sources--nat_policy--reference--group-001.md#canonical-0110313123311131-2200121330312320-3213032012312110-1220121003133103-2022332203221330-0131211202002202-1033231123102023-2101300210202000) |
| `rules.cloud_connect.refs.uid` | [rules.cloud_connect.refs.uid](data-sources--nat_policy--reference--group-001.md#canonical-2331100023221230-3001213332222331-3321000330333121-3022101131233102-3310301322032030-0003211021012320-3321221013030233-2102000211202111) |
| `rules.criteria` | [rules.criteria](data-sources--nat_policy--reference--group-001.md#canonical-3122212330112311-1203320332110102-2131010302322321-0231013232213211-3311220012233021-3132313202233021-2222220010330031-1203030001113101) |
| `rules.criteria.any` | [rules.criteria.any](data-sources--nat_policy--reference--group-001.md#canonical-0211321130000301-2322030221120333-0203111323101231-1131031002203012-3013323112020311-2330322121333001-0312030123333213-3220321302333311) |
| `rules.criteria.destination_cidr` | [rules.criteria.destination_cidr](data-sources--nat_policy--reference--group-001.md#canonical-2031202320131232-0020330221022011-2021333301331103-1012233311123021-1131120123320231-1132133022110311-0111202201030210-3103012121000012) |
| `rules.criteria.icmp` | [rules.criteria.icmp](data-sources--nat_policy--reference--group-001.md#canonical-0223131210001220-3030121333232313-2110233031020130-0001200211323111-0101320201020210-0223320301013132-1003330021103130-3031003302211310) |
| `rules.criteria.site_local_inside_network` | [rules.criteria.site_local_inside_network](data-sources--nat_policy--reference--group-001.md#canonical-2130133030203031-1232030331232022-1232020233131210-2321013231233330-3332101220132231-1333323103312301-2210032303201211-1023221003221322) |
| `rules.criteria.site_local_network` | [rules.criteria.site_local_network](data-sources--nat_policy--reference--group-001.md#canonical-3331121222223203-2101030333110211-0111033020123321-1303030013330211-0112312323230300-1222220000203232-2303230033110200-0330210223211013) |
| `rules.criteria.source_cidr` | [rules.criteria.source_cidr](data-sources--nat_policy--reference--group-001.md#canonical-2311033122210330-0333113201100003-0213221120232332-3123302303112331-1201322133100002-0131010211223132-3301330233221300-3112110211033122) |
| `rules.criteria.tcp` | [rules.criteria.tcp](data-sources--nat_policy--reference--group-001.md#canonical-2000131203120303-2010121001103002-2221112012330033-1111011011223023-3003113323122133-3031300310330223-0102131331023213-3330223023222200) |
| `rules.criteria.tcp.destination_port` | [rules.criteria.tcp.destination_port](data-sources--nat_policy--reference--group-001.md#canonical-0300122223001102-2113200202111333-3121333213200231-1310302100011202-1111133021103201-1322232023103002-3313332030230000-2313113303033011) |
| `rules.criteria.tcp.destination_port.no_port_match` | [rules.criteria.tcp.destination_port.no_port_match](data-sources--nat_policy--reference--group-001.md#canonical-0022321223023231-2321303123103231-1002223123313032-2301330003310002-2330233311220002-0230230111221310-0122333012213020-3122231213333302) |
| `rules.criteria.tcp.destination_port.port` | [rules.criteria.tcp.destination_port.port](data-sources--nat_policy--reference--group-001.md#canonical-3313320000121021-2032312222103132-1121112230133312-2212010201123121-1310230301003131-0030112003201231-3113321232311322-1221202331032231) |
| `rules.criteria.tcp.destination_port.port_ranges` | [rules.criteria.tcp.destination_port.port_ranges](data-sources--nat_policy--reference--group-001.md#canonical-2231210210220302-2012103203231101-1033323023211220-1110201303011330-1122332312322111-2233213000313110-3011031211023101-1331012003213301) |
| `rules.criteria.tcp.source_port` | [rules.criteria.tcp.source_port](data-sources--nat_policy--reference--group-001.md#canonical-1002221120103133-1301221101122122-0330030102331031-3103313103012103-3120202232133211-2232010230331032-1333032202210330-0123303002011002) |
| `rules.criteria.tcp.source_port.no_port_match` | [rules.criteria.tcp.source_port.no_port_match](data-sources--nat_policy--reference--group-001.md#canonical-1302122230010303-3221013001013100-2231001000002201-1133003022111221-2022033020321003-1012132110101113-0321211220300202-2131231202210100) |
| `rules.criteria.tcp.source_port.port` | [rules.criteria.tcp.source_port.port](data-sources--nat_policy--reference--group-001.md#canonical-2300312310002231-0323222221320013-0332122332003331-3111210331103000-2302033301112311-0220322320033113-3022101200331233-2121312131310213) |
| `rules.criteria.tcp.source_port.port_ranges` | [rules.criteria.tcp.source_port.port_ranges](data-sources--nat_policy--reference--group-001.md#canonical-0002313202330012-0301031023213200-2100333300101021-3120333002023111-2331233333231011-0223111100322322-2000201011102222-3001232312311311) |
| `rules.criteria.udp` | [rules.criteria.udp](data-sources--nat_policy--reference--group-001.md#canonical-2233031112333113-1021120320232133-3011323130001111-0032302210111121-3331033333130002-3002021331102013-1121103310131233-1013020110030233) |
| `rules.criteria.udp.destination_port` | [rules.criteria.udp.destination_port](data-sources--nat_policy--reference--group-001.md#canonical-2331232212120022-3232210133013331-2033333003003011-2201103333033031-1202010131113322-0023022030222323-1233331002213112-2121310102001003) |
| `rules.criteria.udp.destination_port.no_port_match` | [rules.criteria.udp.destination_port.no_port_match](data-sources--nat_policy--reference--group-001.md#canonical-3221112310032120-0113133010200132-0103322203113302-0332303223300233-2131032101322021-3023230232123303-0312010120030110-1031300100133213) |
| `rules.criteria.udp.destination_port.port` | [rules.criteria.udp.destination_port.port](data-sources--nat_policy--reference--group-001.md#canonical-2211122101213110-2111102132130301-1032233303033020-0002133332130300-0000211223302011-0101001033120320-3223022032221321-1332000023223300) |
| `rules.criteria.udp.destination_port.port_ranges` | [rules.criteria.udp.destination_port.port_ranges](data-sources--nat_policy--reference--group-001.md#canonical-3011210010333020-2331332033112233-0200233322331320-1100211302122202-0111121003222122-0220103313021232-2131003321233233-2212012321111112) |
| `rules.criteria.udp.source_port` | [rules.criteria.udp.source_port](data-sources--nat_policy--reference--group-001.md#canonical-2123013333203131-3022300223101203-3000023122100031-1013202112031030-0320012300322122-3220200312212202-1020300223222130-2031223011132310) |
| `rules.criteria.udp.source_port.no_port_match` | [rules.criteria.udp.source_port.no_port_match](data-sources--nat_policy--reference--group-001.md#canonical-0303200101230130-1122010020102233-1011102110003200-1131023313002333-3020230133333330-0321212231103210-1013233203223203-1002231012301112) |
| `rules.criteria.udp.source_port.port` | [rules.criteria.udp.source_port.port](data-sources--nat_policy--reference--group-001.md#canonical-3322233002131211-1323331310001132-1320312232311200-3203110103031223-2013330102233201-1223002231033033-3033033133201212-0300311133312113) |
| `rules.criteria.udp.source_port.port_ranges` | [rules.criteria.udp.source_port.port_ranges](data-sources--nat_policy--reference--group-001.md#canonical-3003322233012202-2003201121330012-3111120030311310-1121213000203332-1001200210012120-3203310213212111-2122000001323212-3210122131331033) |
| `rules.disable_spec` | [rules.disable_spec](data-sources--nat_policy--reference--group-001.md#canonical-2020302132213323-3132131012202311-3330030330033321-0301033101022130-2011103211122011-1020220310001323-3000232030203112-1231212012322122) |
| `rules.enable` | [rules.enable](data-sources--nat_policy--reference--group-001.md#canonical-0122031000202303-0022302233103203-2012232021231003-1200121120101212-1233220221023031-1002021201222310-3031312302231131-1131033112221132) |
| `rules.name` | [rules.name](data-sources--nat_policy--reference--group-001.md#canonical-1113112223023220-2003010322331030-1130220122032212-0323010330313310-0313032031323003-3311201231330112-0112333023032033-0233203213031323) |
| `rules.node_interface` | [rules.node_interface](data-sources--nat_policy--reference--group-001.md#canonical-3111201303131012-2302331132103100-2301220323122021-3301020331003120-0222112230102112-0213212302212023-0223310103010022-3102230211231312) |
| `rules.node_interface.list` | [rules.node_interface.list](data-sources--nat_policy--reference--group-001.md#canonical-3031323123332100-3220331011200233-2021300233233132-1223101120211013-2201310203212113-0232133203011313-1331131031320331-2300113131111113) |
| `rules.node_interface.list.interface` | [rules.node_interface.list.interface](data-sources--nat_policy--reference--group-001.md#canonical-0123330011112010-0120021213211000-3332232121120013-1112333031223301-1313110310133132-1233020020231100-1330032320002332-0202313202032222) |
| `rules.node_interface.list.interface.kind` | [rules.node_interface.list.interface.kind](data-sources--nat_policy--reference--group-001.md#canonical-3000031233200122-2110120321000232-0102121102113111-0321320221031331-2130330120112230-3132001330031102-1132000030012031-2003230033212311) |
| `rules.node_interface.list.interface.name` | [rules.node_interface.list.interface.name](data-sources--nat_policy--reference--group-001.md#canonical-1331112333111202-3231010212032331-0131213302032000-3223201210132333-1333231131112313-2020001320030131-0111002112103231-2102302310303303) |
| `rules.node_interface.list.interface.namespace` | [rules.node_interface.list.interface.namespace](data-sources--nat_policy--reference--group-001.md#canonical-3130122222312033-0020013112302333-1230103131113312-0301232103303121-3010210113101333-3330230123032101-2033113201223003-2313130223222301) |
| `rules.node_interface.list.interface.tenant` | [rules.node_interface.list.interface.tenant](data-sources--nat_policy--reference--group-001.md#canonical-1220123311202130-2010123203013200-0113002332211111-3311002031033010-3231132332021211-2230321132010103-1202322010100122-0222230303122303) |
| `rules.node_interface.list.interface.uid` | [rules.node_interface.list.interface.uid](data-sources--nat_policy--reference--group-001.md#canonical-2131001200213222-0033013112003321-0121030113323102-0002110012121331-2203012223303201-0312102303320000-3110222123123330-3012000322103302) |
| `rules.node_interface.list.node` | [rules.node_interface.list.node](data-sources--nat_policy--reference--group-001.md#canonical-3103003332232231-2331113133313333-1220122202211033-3230100330010002-1011313231023233-3223223002211333-2222232021102320-2100201033003022) |
| `rules.segment` | [rules.segment](data-sources--nat_policy--reference--group-001.md#canonical-0321013101232012-0102301113103020-2033100230301111-1013032320002221-2011331213302331-1030210121103022-2313231111310012-1111010301110013) |
| `rules.segment.refs` | [rules.segment.refs](data-sources--nat_policy--reference--group-001.md#canonical-2303123111311231-1111130223110222-1110130030302000-3331200322002222-1032123312231022-0032102030113312-1222211031132330-0000332313312123) |
| `rules.segment.refs.kind` | [rules.segment.refs.kind](data-sources--nat_policy--reference--group-001.md#canonical-3221023033023031-1331220230221333-0312303101311210-0013122101201331-3230020203210232-1223222320303001-2212101010011202-3333023320321030) |
| `rules.segment.refs.name` | [rules.segment.refs.name](data-sources--nat_policy--reference--group-001.md#canonical-2001023131002230-2032201023011022-2102022303131233-2000323111103322-0032133130210033-2221131222331200-3230302311120132-1311113030121011) |
| `rules.segment.refs.namespace` | [rules.segment.refs.namespace](data-sources--nat_policy--reference--group-001.md#canonical-1013302222021110-0032031302313201-1202020113030011-0102331330200030-2221002312023310-0101121121222322-1101121133201302-3210213223210232) |
| `rules.segment.refs.tenant` | [rules.segment.refs.tenant](data-sources--nat_policy--reference--group-001.md#canonical-3123220331311322-0230222020113112-3300322130302120-1133213021233021-0133002211030023-1223103122003113-3322330311102333-0201121212133012) |
| `rules.segment.refs.uid` | [rules.segment.refs.uid](data-sources--nat_policy--reference--group-001.md#canonical-2110023122102213-3213333202132200-2013012013212110-3122201122312122-1201100011331120-0032202102203032-3223103112121332-2103011000112123) |
| `rules.virtual_network` | [rules.virtual_network](data-sources--nat_policy--reference--group-001.md#canonical-3023313122233333-0323320112021112-3303211323203032-1210010312103310-3213120123002022-1211122110200311-0320302003010100-1032122133112022) |
| `rules.virtual_network.refs` | [rules.virtual_network.refs](data-sources--nat_policy--reference--group-001.md#canonical-3130010321313133-2203113322301312-3231332223320120-0012102321210222-0311023210011303-1013102102032000-2230211210111333-1321331330200110) |
| `rules.virtual_network.refs.kind` | [rules.virtual_network.refs.kind](data-sources--nat_policy--reference--group-001.md#canonical-3133222010032103-2312222322033100-1131312213032331-1122312212313201-2212010310003333-2111130201011013-2030113311301130-0230020223120221) |
| `rules.virtual_network.refs.name` | [rules.virtual_network.refs.name](data-sources--nat_policy--reference--group-001.md#canonical-2212231113312032-3111321103010102-2332320221212220-2202020223313210-0021013032331233-3312002020320321-2221303230100010-3010302022100020) |
| `rules.virtual_network.refs.namespace` | [rules.virtual_network.refs.namespace](data-sources--nat_policy--reference--group-001.md#canonical-3121322132012012-3000130030113003-0210021111210200-0310112021211121-1131010012120130-2012301121220032-1010030322100033-0120121221000221) |
| `rules.virtual_network.refs.tenant` | [rules.virtual_network.refs.tenant](data-sources--nat_policy--reference--group-001.md#canonical-0013220320003321-0211020203313333-1022010023311001-1233002113213020-3131120213231220-3130223112300023-1033030030120112-1322001113000201) |
| `rules.virtual_network.refs.uid` | [rules.virtual_network.refs.uid](data-sources--nat_policy--reference--group-001.md#canonical-1311200221332212-2130312231121313-2201013331233021-1130122102111310-3021133131310203-0103331221230033-1112122013322132-3201110021032001) |
| `site` | [site](data-sources--nat_policy--reference--group-001.md#canonical-1313333331003210-2113112121111033-1002232130112003-0201101313311023-3310102223312210-0330010133312033-1302220120232132-2313200231213101) |
| `site.refs` | [site.refs](data-sources--nat_policy--reference--group-001.md#canonical-3113301200212233-0030300123200000-1320000230030330-1133110033330003-3230212210100303-2013021230111130-2130100133313320-1320021210233303) |
| `site.refs.kind` | [site.refs.kind](data-sources--nat_policy--reference--group-001.md#canonical-2301232201212122-0122002112103002-0100220120103303-1330211123111122-1123330332011112-2200211332011002-3222112113101123-3132032303002000) |
| `site.refs.name` | [site.refs.name](data-sources--nat_policy--reference--group-001.md#canonical-3233331202221322-3311313301002101-3011023101203001-3303332011311221-2031320110103012-2101110020320332-1123301130202202-1012011002313210) |
| `site.refs.namespace` | [site.refs.namespace](data-sources--nat_policy--reference--group-001.md#canonical-1101120132322203-1122311102033011-3320010323323112-0023112021031120-2022220213022002-0221112003321230-3120030203103013-1210002310103131) |
| `site.refs.tenant` | [site.refs.tenant](data-sources--nat_policy--reference--group-001.md#canonical-1031021001202012-0032100120233101-0001003100223302-0022331000322303-2003321221233301-0220020222002033-2130312023100312-1221311002211122) |
| `site.refs.uid` | [site.refs.uid](data-sources--nat_policy--reference--group-001.md#canonical-0101103113033031-1210210103321320-1013212233201211-2302332232223003-2332300223203223-0111113321301002-1103202223323101-0031230111022013) |

<a id="canonical-0211110210131101-2303300303122112-0031321323203023-2003313333331231-3032223122221033-1121322010202212-2233233201200212-2000311112203230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rules` properties

Breadcrumbs:

- [xcsh_nat_policy](../data-sources/nat_policy.md#canonical-2211333230110221-1023302013201300-2020112103300231-1033030223223010-0020003220130132-3323021301233201-1030131302223322-2302033131310121)
- [Property reference](data-sources--nat_policy--reference--group-001.md#canonical-0313000023303013-1312311300100203-2312121103033010-0011202012203121-2320110020132021-0013320012122120-1332300232020201-2300013013202121)
- rules

<a id="canonical-3122332210112211-2003013132131222-1030202323331303-0032210223011333-1012312202301012-3013130121121002-2001032112100132-2121111303320121"></a>

Type: `"list"`. Computed.

List of rules to apply under the NAT Policy. Rule that matches first would be applied.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 64,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-09T12:34:59+00:00"
    },
    "minItems": 0,
    "uniqueItems": false
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "64"
  }
}
```

<a id="canonical-0301311023113001-3033333033220120-3303132202010000-3103212202112013-2131113211210320-2312311110330121-1130213112310330-0003332132112313"></a>

### Direct properties for `rules`

- [action](data-sources--nat_policy--reference--group-001.md#canonical-1113032111000113-3122312303102131-3301332200331203-1032122102301330-3023222303210123-3113201001000301-2112101101312111-3211320212303220): complete subsection reference.

- [cloud_connect](data-sources--nat_policy--reference--group-001.md#canonical-2002332211313233-0323320313100303-1101121333232222-0221010310231023-0110021131222030-1000023213330201-0023023212120113-3230132223300030): complete subsection reference.

- [criteria](data-sources--nat_policy--reference--group-001.md#canonical-3232031321200303-0000110111113213-2023132130333021-2321302103203031-1121302211233200-1333323321303202-3330233302031121-1013033113213023): complete subsection reference.

- [disable_spec](data-sources--nat_policy--reference--group-001.md#canonical-2210320323002232-0032213123210011-2131213200110203-1323222331000202-1100113212133120-2211000002003222-0122301232013100-1201321200321312): complete subsection reference.

- [enable](data-sources--nat_policy--reference--group-001.md#canonical-1332102110100203-3003203212301202-2010122122322333-2201132312131032-3310001132101102-2203022223233333-2323222211330023-1021103211223303): complete subsection reference.

<a id="canonical-1113112223023220-2003010322331030-1130220122032212-0323010330313310-0313032031323003-3311201231330112-0112333023032033-0233203213031323"></a>

<a id="canonical-1231230131223033-1213113023301333-3200132213012212-2101022233111032-1233133101112300-0331131222123312-1110332003030211-0003220203132030"></a>

#### `rules.name` property

Type: `"string"`. Computed.

Name. Name of the Rule.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

- [node_interface](data-sources--nat_policy--reference--group-001.md#canonical-3102133303312131-0221130013231131-3133131012123133-1033013320311022-2201330322003310-3232111120011023-3002022132220023-0032233132110111): complete subsection reference.

- [segment](data-sources--nat_policy--reference--group-001.md#canonical-3122233211112302-2002223232110001-0231200113302123-0202221331112123-2023313220021211-0232101321133010-3000032212300301-3222313102331311): complete subsection reference.

- [virtual_network](data-sources--nat_policy--reference--group-001.md#canonical-1222133003113213-2211003133231313-3030010301222022-1032232221210133-2133032033131032-2133122120012312-3112332233032010-3032123201113221): complete subsection reference.

<a id="canonical-1113032111000113-3122312303102131-3301332200331203-1032122102301330-3023222303210123-3113201001000301-2112101101312111-3211320212303220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rules.action` properties

Breadcrumbs:

- [xcsh_nat_policy](../data-sources/nat_policy.md#canonical-2211333230110221-1023302013201300-2020112103300231-1033030223223010-0020003220130132-3323021301233201-1030131302223322-2302033131310121)
- [Property reference](data-sources--nat_policy--reference--group-001.md#canonical-0313000023303013-1312311300100203-2312121103033010-0011202012203121-2320110020132021-0013320012122120-1332300232020201-2300013013202121)
- [rules](data-sources--nat_policy--reference--group-001.md#canonical-0211110210131101-2303300303122112-0031321323203023-2003313333331231-3032223122221033-1121322010202212-2233233201200212-2000311112203230)
- rules.action

<a id="canonical-3130131002322203-1211222203233032-1311300110230122-3012121030120222-3113231211122321-1131020320012101-0120103311111133-0002212303220220"></a>

Type: `"single"`. Computed.

Action to apply on the packet if the NAT rule is applied.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-source_nat_choice": "[\"dynamic\",\"virtual_cidr\"]"
}
```

<a id="canonical-2000030121212203-1302120132210311-3333331322332321-3113211301110300-2111321331202333-2303232031222110-1031012122312313-3020310013012033"></a>

### Direct properties for `rules.action`

- [dynamic](data-sources--nat_policy--reference--group-001.md#canonical-2111313012302100-0003302302202212-2213222023032033-1320032310010322-0211313001201110-3000032221032332-1303130110101213-2333210330211202): complete subsection reference.

<a id="canonical-0102231110000302-0232113333223222-2112213202000012-1020210323312223-3231300130110010-2103312032320301-1213001001232020-1033233110000331"></a>

<a id="canonical-3032223231021233-0233312330102231-3022212103213020-3023320121033123-0303322103003032-3132233312001032-1323220310122210-3300120300112122"></a>

#### `rules.action.virtual_cidr` property

Type: `"string"`. Computed.

Exclusive with \[dynamic\] Virtual Subnet NAT is static NAT that does a one-to-one translation
between the real source IP CIDR in the policy and the virtual CIDR in a bidirectional fashion. The
range of the real CIDR and virtual CIDRs should be the same (e.g. If the real CIDR has the CIDR
192.0.2.0/24, the virtual CIDR has 100.100.100.0/24.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "cidr",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-09T12:34:59+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.ip_prefix": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ip_prefix": "true"
  }
}
```

<a id="canonical-2111313012302100-0003302302202212-2213222023032033-1320032310010322-0211313001201110-3000032221032332-1303130110101213-2333210330211202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rules.action.dynamic` properties

Breadcrumbs:

- [xcsh_nat_policy](../data-sources/nat_policy.md#canonical-2211333230110221-1023302013201300-2020112103300231-1033030223223010-0020003220130132-3323021301233201-1030131302223322-2302033131310121)
- [Property reference](data-sources--nat_policy--reference--group-001.md#canonical-0313000023303013-1312311300100203-2312121103033010-0011202012203121-2320110020132021-0013320012122120-1332300232020201-2300013013202121)
- [rules](data-sources--nat_policy--reference--group-001.md#canonical-0211110210131101-2303300303122112-0031321323203023-2003313333331231-3032223122221033-1121322010202212-2233233201200212-2000311112203230)
- [rules.action](data-sources--nat_policy--reference--group-001.md#canonical-1113032111000113-3122312303102131-3301332200331203-1032122102301330-3023222303210123-3113201001000301-2112101101312111-3211320212303220)
- rules.action.dynamic

<a id="canonical-1213113100011313-0132133001132220-0002101200320033-3300232131323301-3332021133232020-2323302301131332-1101022103021220-3112233331212100"></a>

Type: `"single"`. Computed.

Dynamic Pool. Dynamic Pool Configuration.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-pool_choice": "[\"elastic_ips\",\"pools\"]"
}
```

<a id="canonical-2301202300030103-3313201122232312-1013222201012133-1210033310232212-2220230311011132-1310013300300230-3100233203223013-0001133111131322"></a>

### Direct properties for `rules.action.dynamic`

- [elastic_ips](data-sources--nat_policy--reference--group-001.md#canonical-0230003220303133-1230022313322321-2000012100313200-3120031001030001-2222023000333022-2331113020212101-0323310011033322-2013300011312320): complete subsection reference.

- [pools](data-sources--nat_policy--reference--group-001.md#canonical-2301012323233132-3133031320102112-3301202112222101-0011103313110310-3103103331213001-3011310023320221-3332212313323202-0123321101112001): complete subsection reference.

<a id="canonical-0230003220303133-1230022313322321-2000012100313200-3120031001030001-2222023000333022-2331113020212101-0323310011033322-2013300011312320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rules.action.dynamic.elastic_ips` properties

Breadcrumbs:

- [xcsh_nat_policy](../data-sources/nat_policy.md#canonical-2211333230110221-1023302013201300-2020112103300231-1033030223223010-0020003220130132-3323021301233201-1030131302223322-2302033131310121)
- [Property reference](data-sources--nat_policy--reference--group-001.md#canonical-0313000023303013-1312311300100203-2312121103033010-0011202012203121-2320110020132021-0013320012122120-1332300232020201-2300013013202121)
- [rules](data-sources--nat_policy--reference--group-001.md#canonical-0211110210131101-2303300303122112-0031321323203023-2003313333331231-3032223122221033-1121322010202212-2233233201200212-2000311112203230)
- [rules.action](data-sources--nat_policy--reference--group-001.md#canonical-1113032111000113-3122312303102131-3301332200331203-1032122102301330-3023222303210123-3113201001000301-2112101101312111-3211320212303220)
- [rules.action.dynamic](data-sources--nat_policy--reference--group-001.md#canonical-2111313012302100-0003302302202212-2213222023032033-1320032310010322-0211313001201110-3000032221032332-1303130110101213-2333210330211202)
- rules.action.dynamic.elastic_ips

<a id="canonical-1303212022310122-2132110230030121-3203222110200103-0220230312310001-3330031311023220-2222101132232220-3122331330120020-1032231103200132"></a>

Type: `"single"`. Computed.

List of references to Cloud Elastic IP Object.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3223233001111130-2100211203112202-2023023102313203-1211033232032322-0321213331300222-2233100102030003-3000310313131333-1330323103113100"></a>

### Direct properties for `rules.action.dynamic.elastic_ips`

- [refs](data-sources--nat_policy--reference--group-001.md#canonical-3313022102230011-3132233300332121-3222330001332203-3003311230003200-2111031221023232-3102330221000131-1031031112303231-0132123332023222): complete subsection reference.

<a id="canonical-3313022102230011-3132233300332121-3222330001332203-3003311230003200-2111031221023232-3102330221000131-1031031112303231-0132123332023222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rules.action.dynamic.elastic_ips.refs` properties

Breadcrumbs:

- [xcsh_nat_policy](../data-sources/nat_policy.md#canonical-2211333230110221-1023302013201300-2020112103300231-1033030223223010-0020003220130132-3323021301233201-1030131302223322-2302033131310121)
- [Property reference](data-sources--nat_policy--reference--group-001.md#canonical-0313000023303013-1312311300100203-2312121103033010-0011202012203121-2320110020132021-0013320012122120-1332300232020201-2300013013202121)
- [rules](data-sources--nat_policy--reference--group-001.md#canonical-0211110210131101-2303300303122112-0031321323203023-2003313333331231-3032223122221033-1121322010202212-2233233201200212-2000311112203230)
- [rules.action](data-sources--nat_policy--reference--group-001.md#canonical-1113032111000113-3122312303102131-3301332200331203-1032122102301330-3023222303210123-3113201001000301-2112101101312111-3211320212303220)
- [rules.action.dynamic](data-sources--nat_policy--reference--group-001.md#canonical-2111313012302100-0003302302202212-2213222023032033-1320032310010322-0211313001201110-3000032221032332-1303130110101213-2333210330211202)
- [rules.action.dynamic.elastic_ips](data-sources--nat_policy--reference--group-001.md#canonical-0230003220303133-1230022313322321-2000012100313200-3120031001030001-2222023000333022-2331113020212101-0323310011033322-2013300011312320)
- rules.action.dynamic.elastic_ips.refs

<a id="canonical-0321102202111022-2313022011111211-2232130220031023-2330033230120113-2303030012312202-2311033020112001-2231233323123212-1023232133101130"></a>

Type: `"list"`. Computed.

Reference to one or more cloud elastic IP objects.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "4"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "4"
  }
}
```

<a id="canonical-1232310101230330-2323301311122332-2331333313201001-3012311000201210-2331031101303001-3111201323022022-1022231132332120-0220220232020320"></a>

### Direct properties for `rules.action.dynamic.elastic_ips.refs`

<a id="canonical-2201220011301211-0313213210202201-2022211012211133-2122013100120231-0311300112300000-2122103000311221-1003131320011030-0322302032031323"></a>

#### `rules.action.dynamic.elastic_ips.refs.kind` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. 'route').

Additional upstream details:

Virtual\_host) refers to another(e.g route) then kind will hold the referred object's kind (e.g.
"route")

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-0120022133113213-1320030302333110-3131320121000130-3101113300113023-1121130223113301-0101133103100231-2031221102032020-1022310131222123"></a>

<a id="canonical-1123203303201213-2211323000313013-1310211111013313-3132123000212013-2122122031023212-1121011220003003-0102311231323231-3320311020022200"></a>

#### `rules.action.dynamic.elastic_ips.refs.name` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-0012120123222202-2221332102330311-1123320010220302-0033020123023200-2103211130332321-3131022320331312-1023030320031232-3203231302232032"></a>

<a id="canonical-2030111022002102-1333202021201333-0230313123233231-2021021303031301-3020010032030123-2101200112012022-2022113203102110-0233101000033221"></a>

#### `rules.action.dynamic.elastic_ips.refs.namespace` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-2020230201301221-0220300233120203-0010111130033031-1202201133033113-2230031330122221-0100033012030031-3220221020131332-0211123033301223"></a>

<a id="canonical-3100133210112213-3022031030021302-3222022220123302-3012110121100302-2303031100202011-0020201030001221-1130201311102033-2020011202221323"></a>

#### `rules.action.dynamic.elastic_ips.refs.tenant` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-2010231330101021-1010300031332222-0102213330032123-2323313110133002-2030103331230331-0031320020301131-0213330022320300-3302120202331211"></a>

<a id="canonical-2112201003202330-0332313321032011-0332012103230003-1331012013031013-2201311232000110-1212230233112233-0321020022111231-2033223212010121"></a>

#### `rules.action.dynamic.elastic_ips.refs.uid` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-2301012323233132-3133031320102112-3301202112222101-0011103313110310-3103103331213001-3011310023320221-3332212313323202-0123321101112001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rules.action.dynamic.pools` properties

Breadcrumbs:

- [xcsh_nat_policy](../data-sources/nat_policy.md#canonical-2211333230110221-1023302013201300-2020112103300231-1033030223223010-0020003220130132-3323021301233201-1030131302223322-2302033131310121)
- [Property reference](data-sources--nat_policy--reference--group-001.md#canonical-0313000023303013-1312311300100203-2312121103033010-0011202012203121-2320110020132021-0013320012122120-1332300232020201-2300013013202121)
- [rules](data-sources--nat_policy--reference--group-001.md#canonical-0211110210131101-2303300303122112-0031321323203023-2003313333331231-3032223122221033-1121322010202212-2233233201200212-2000311112203230)
- [rules.action](data-sources--nat_policy--reference--group-001.md#canonical-1113032111000113-3122312303102131-3301332200331203-1032122102301330-3023222303210123-3113201001000301-2112101101312111-3211320212303220)
- [rules.action.dynamic](data-sources--nat_policy--reference--group-001.md#canonical-2111313012302100-0003302302202212-2213222023032033-1320032310010322-0211313001201110-3000032221032332-1303130110101213-2333210330211202)
- rules.action.dynamic.pools

<a id="canonical-3031213031022223-1231011030220200-1331012132312213-3122310012331223-0120301313031111-2023112303331232-3032303011112302-0101122112233133"></a>

Type: `"single"`. Computed.

List of IPv4 prefixes that represent an endpoint.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3012231111001100-1312110113320002-3000333322002312-2110032313320322-3120133202121013-0330223222032332-2111021031302010-1030120312122300"></a>

### Direct properties for `rules.action.dynamic.pools`

<a id="canonical-1021122331203000-3323323223130113-1011323222110212-0332212022130212-2310231013111120-3002230000130121-0030203133223301-1002202210112003"></a>

#### `rules.action.dynamic.pools.prefixes` property

Type: `["list", "string"]`. Computed.

List of IPv4 prefixes that represent an endpoint.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2002332211313233-0323320313100303-1101121333232222-0221010310231023-0110021131222030-1000023213330201-0023023212120113-3230132223300030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rules.cloud_connect` properties

Breadcrumbs:

- [xcsh_nat_policy](../data-sources/nat_policy.md#canonical-2211333230110221-1023302013201300-2020112103300231-1033030223223010-0020003220130132-3323021301233201-1030131302223322-2302033131310121)
- [Property reference](data-sources--nat_policy--reference--group-001.md#canonical-0313000023303013-1312311300100203-2312121103033010-0011202012203121-2320110020132021-0013320012122120-1332300232020201-2300013013202121)
- [rules](data-sources--nat_policy--reference--group-001.md#canonical-0211110210131101-2303300303122112-0031321323203023-2003313333331231-3032223122221033-1121322010202212-2233233201200212-2000311112203230)
- rules.cloud_connect

<a id="canonical-3322312213223330-3332113132330201-2220033320302013-3113220001302332-3010130303130132-3200022302100220-1313310101302111-2301223233223102"></a>

Type: `"single"`. Computed.

Configuration parameter for cloud connect.

Additional upstream details:

Reference to Cloud connect Object.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2201122112320031-1010110202002330-3232100000302322-0123331321121003-2232023102233212-0323002223232112-2202132130001330-2211122303001031"></a>

### Direct properties for `rules.cloud_connect`

- [refs](data-sources--nat_policy--reference--group-001.md#canonical-2220002330201112-0323112221233231-3311202221211121-1311001221013133-2212100233211222-2310231010213222-2220300021021123-2221120122321302): complete subsection reference.

<a id="canonical-2220002330201112-0323112221233231-3311202221211121-1311001221013133-2212100233211222-2310231010213222-2220300021021123-2221120122321302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rules.cloud_connect.refs` properties

Breadcrumbs:

- [xcsh_nat_policy](../data-sources/nat_policy.md#canonical-2211333230110221-1023302013201300-2020112103300231-1033030223223010-0020003220130132-3323021301233201-1030131302223322-2302033131310121)
- [Property reference](data-sources--nat_policy--reference--group-001.md#canonical-0313000023303013-1312311300100203-2312121103033010-0011202012203121-2320110020132021-0013320012122120-1332300232020201-2300013013202121)
- [rules](data-sources--nat_policy--reference--group-001.md#canonical-0211110210131101-2303300303122112-0031321323203023-2003313333331231-3032223122221033-1121322010202212-2233233201200212-2000311112203230)
- [rules.cloud_connect](data-sources--nat_policy--reference--group-001.md#canonical-2002332211313233-0323320313100303-1101121333232222-0221010310231023-0110021131222030-1000023213330201-0023023212120113-3230132223300030)
- rules.cloud_connect.refs

<a id="canonical-1332133011313013-2320220332233232-1032312101321121-1322122323323023-1132300020021311-0202110203123202-3021111000003301-1323322223103232"></a>

Type: `"list"`. Computed.

Cloud Connect. Reference to Cloud Connect Object.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "1"
  }
}
```

<a id="canonical-1222032020013221-1110313120132221-3000222111000033-3313021122132101-3323101320130312-2123203330323002-2130231022220111-2201012031131310"></a>

### Direct properties for `rules.cloud_connect.refs`

<a id="canonical-0132021312232312-1323130331300112-1223100320211233-1010320003011312-3111120101313102-2233110121221311-3133232231030300-0120121121322121"></a>

#### `rules.cloud_connect.refs.kind` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. 'route').

Additional upstream details:

Virtual\_host) refers to another(e.g route) then kind will hold the referred object's kind (e.g.
"route")

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-1110333310011323-0111211113210211-2133013000131133-0322230332112123-0013310321232212-0112212231312111-1331030300130223-1212113020230222"></a>

<a id="canonical-0132120223112001-0232120023322303-1003211213133300-1003203203120302-2331300121231122-0232111112103200-1121321303320200-2233321133130132"></a>

#### `rules.cloud_connect.refs.name` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-2133330303203233-0003300120213102-3322220303201333-0013323031123213-0111103231212023-2212100111120221-3221020231103122-2000121200000203"></a>

<a id="canonical-2033331032002311-3303100111123322-2111223021211003-0301120110211202-3123330013102233-1202331032322323-1112010001322202-3331333321130020"></a>

#### `rules.cloud_connect.refs.namespace` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-0110313123311131-2200121330312320-3213032012312110-1220121003133103-2022332203221330-0131211202002202-1033231123102023-2101300210202000"></a>

<a id="canonical-2311321301100012-3013213232100233-2333102011201132-2030002233030310-1211230103001222-3001312223021230-1011331202223022-2300111023230233"></a>

#### `rules.cloud_connect.refs.tenant` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-2331100023221230-3001213332222331-3321000330333121-3022101131233102-3310301322032030-0003211021012320-3321221013030233-2102000211202111"></a>

<a id="canonical-2130312323011301-2230001021233202-0103323200122333-3230301303001303-1203232002000203-3220230110230113-1221022012001020-2313120222020300"></a>

#### `rules.cloud_connect.refs.uid` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-3232031321200303-0000110111113213-2023132130333021-2321302103203031-1121302211233200-1333323321303202-3330233302031121-1013033113213023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rules.criteria` properties

Breadcrumbs:

- [xcsh_nat_policy](../data-sources/nat_policy.md#canonical-2211333230110221-1023302013201300-2020112103300231-1033030223223010-0020003220130132-3323021301233201-1030131302223322-2302033131310121)
- [Property reference](data-sources--nat_policy--reference--group-001.md#canonical-0313000023303013-1312311300100203-2312121103033010-0011202012203121-2320110020132021-0013320012122120-1332300232020201-2300013013202121)
- [rules](data-sources--nat_policy--reference--group-001.md#canonical-0211110210131101-2303300303122112-0031321323203023-2003313333331231-3032223122221033-1121322010202212-2233233201200212-2000311112203230)
- rules.criteria

<a id="canonical-3122212330112311-1203320332110102-2131010302322321-0231013232213211-3311220012233021-3132313202233021-2222220010330031-1203030001113101"></a>

Type: `"single"`. Computed.

Match criteria of the packet to apply the NAT Rule.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-network_choice": "[\"site_local_inside_network\",\"site_local_network\"]",
  "x-ves-oneof-field-protocol_choice": "[\"any\",\"icmp\",\"tcp\",\"udp\"]"
}
```

<a id="canonical-0010233311203013-3132212001213321-3222113230330003-2001031301333113-1230120322223023-0220321303002021-2333023013233202-1010121302301030"></a>

### Direct properties for `rules.criteria`

- [any](data-sources--nat_policy--reference--group-001.md#canonical-3111303223302100-0033311113123313-1033330330310033-1332101020223232-3230201111321323-3102001021320013-0000211130123302-2100302020331223): complete subsection reference.

<a id="canonical-2031202320131232-0020330221022011-2021333301331103-1012233311123021-1131120123320231-1132133022110311-0111202201030210-3103012121000012"></a>

<a id="canonical-2230032121202132-2222010311211330-1220221332311312-1010211010330313-0221023311002301-0301001101100101-2002231022232222-1312220232231220"></a>

#### `rules.criteria.destination_cidr` property

Type: `["list", "string"]`. Computed.

Destination IP. Destination IP of the packet to match.

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
    "ves.io.schema.rules.string.ip_prefix": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ip_prefix": "true"
  }
}
```

- [icmp](data-sources--nat_policy--reference--group-001.md#canonical-2211013012030303-1002111222003003-2022203210101302-1113111233221011-3203033321132312-1031021202333121-0012132312111130-3231013230003313): complete subsection reference.

- [site_local_inside_network](data-sources--nat_policy--reference--group-001.md#canonical-2233123033031222-3300211330022223-3000233220003313-3113023202222103-0233102323202113-2303013033021330-2133123202132313-3103321130221003): complete subsection reference.

- [site_local_network](data-sources--nat_policy--reference--group-001.md#canonical-0113313110123111-3131320022230030-1333223303321132-1132232020032203-0132200231232031-3011231201023232-2030213302033221-2033213023010013): complete subsection reference.

<a id="canonical-2311033122210330-0333113201100003-0213221120232332-3123302303112331-1201322133100002-0131010211223132-3301330233221300-3112110211033122"></a>

<a id="canonical-3311032120120200-3000112323011010-1313202222020020-1120302033110330-0101130011121210-0003103122003132-0103000112032000-0030033200112303"></a>

#### `rules.criteria.source_cidr` property

Type: `["list", "string"]`. Computed.

Source IP. Source IP of the packet to match.

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
    "ves.io.schema.rules.string.ip_prefix": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ip_prefix": "true"
  }
}
```

- [tcp](data-sources--nat_policy--reference--group-001.md#canonical-0302030212022102-0022203320220102-2301303202022232-1233212230311300-3123132101033331-0113312012012101-1200202130101220-3132111313230213): complete subsection reference.

- [udp](data-sources--nat_policy--reference--group-001.md#canonical-3030211000021231-0010012323012221-2213120030321012-0333212021213101-0011113222330322-2100201312100232-3322220023030011-0232310121212331): complete subsection reference.

<a id="canonical-3111303223302100-0033311113123313-1033330330310033-1332101020223232-3230201111321323-3102001021320013-0000211130123302-2100302020331223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rules.criteria.any` properties

Breadcrumbs:

- [xcsh_nat_policy](../data-sources/nat_policy.md#canonical-2211333230110221-1023302013201300-2020112103300231-1033030223223010-0020003220130132-3323021301233201-1030131302223322-2302033131310121)
- [Property reference](data-sources--nat_policy--reference--group-001.md#canonical-0313000023303013-1312311300100203-2312121103033010-0011202012203121-2320110020132021-0013320012122120-1332300232020201-2300013013202121)
- [rules](data-sources--nat_policy--reference--group-001.md#canonical-0211110210131101-2303300303122112-0031321323203023-2003313333331231-3032223122221033-1121322010202212-2233233201200212-2000311112203230)
- [rules.criteria](data-sources--nat_policy--reference--group-001.md#canonical-3232031321200303-0000110111113213-2023132130333021-2321302103203031-1121302211233200-1333323321303202-3330233302031121-1013033113213023)
- rules.criteria.any

<a id="canonical-0211321130000301-2322030221120333-0203111323101231-1131031002203012-3013323112020311-2330322121333001-0312030123333213-3220321302333311"></a>

Type: `["object", {}]`. Computed.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2211013012030303-1002111222003003-2022203210101302-1113111233221011-3203033321132312-1031021202333121-0012132312111130-3231013230003313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rules.criteria.icmp` properties

Breadcrumbs:

- [xcsh_nat_policy](../data-sources/nat_policy.md#canonical-2211333230110221-1023302013201300-2020112103300231-1033030223223010-0020003220130132-3323021301233201-1030131302223322-2302033131310121)
- [Property reference](data-sources--nat_policy--reference--group-001.md#canonical-0313000023303013-1312311300100203-2312121103033010-0011202012203121-2320110020132021-0013320012122120-1332300232020201-2300013013202121)
- [rules](data-sources--nat_policy--reference--group-001.md#canonical-0211110210131101-2303300303122112-0031321323203023-2003313333331231-3032223122221033-1121322010202212-2233233201200212-2000311112203230)
- [rules.criteria](data-sources--nat_policy--reference--group-001.md#canonical-3232031321200303-0000110111113213-2023132130333021-2321302103203031-1121302211233200-1333323321303202-3330233302031121-1013033113213023)
- rules.criteria.icmp

<a id="canonical-0223131210001220-3030121333232313-2110233031020130-0001200211323111-0101320201020210-0223320301013132-1003330021103130-3031003302211310"></a>

Type: `["object", {}]`. Computed.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2233123033031222-3300211330022223-3000233220003313-3113023202222103-0233102323202113-2303013033021330-2133123202132313-3103321130221003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rules.criteria.site_local_inside_network` properties

Breadcrumbs:

- [xcsh_nat_policy](../data-sources/nat_policy.md#canonical-2211333230110221-1023302013201300-2020112103300231-1033030223223010-0020003220130132-3323021301233201-1030131302223322-2302033131310121)
- [Property reference](data-sources--nat_policy--reference--group-001.md#canonical-0313000023303013-1312311300100203-2312121103033010-0011202012203121-2320110020132021-0013320012122120-1332300232020201-2300013013202121)
- [rules](data-sources--nat_policy--reference--group-001.md#canonical-0211110210131101-2303300303122112-0031321323203023-2003313333331231-3032223122221033-1121322010202212-2233233201200212-2000311112203230)
- [rules.criteria](data-sources--nat_policy--reference--group-001.md#canonical-3232031321200303-0000110111113213-2023132130333021-2321302103203031-1121302211233200-1333323321303202-3330233302031121-1013033113213023)
- rules.criteria.site_local_inside_network

<a id="canonical-2130133030203031-1232030331232022-1232020233131210-2321013231233330-3332101220132231-1333323103312301-2210032303201211-1023221003221322"></a>

Type: `["object", {}]`. Computed.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0113313110123111-3131320022230030-1333223303321132-1132232020032203-0132200231232031-3011231201023232-2030213302033221-2033213023010013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rules.criteria.site_local_network` properties

Breadcrumbs:

- [xcsh_nat_policy](../data-sources/nat_policy.md#canonical-2211333230110221-1023302013201300-2020112103300231-1033030223223010-0020003220130132-3323021301233201-1030131302223322-2302033131310121)
- [Property reference](data-sources--nat_policy--reference--group-001.md#canonical-0313000023303013-1312311300100203-2312121103033010-0011202012203121-2320110020132021-0013320012122120-1332300232020201-2300013013202121)
- [rules](data-sources--nat_policy--reference--group-001.md#canonical-0211110210131101-2303300303122112-0031321323203023-2003313333331231-3032223122221033-1121322010202212-2233233201200212-2000311112203230)
- [rules.criteria](data-sources--nat_policy--reference--group-001.md#canonical-3232031321200303-0000110111113213-2023132130333021-2321302103203031-1121302211233200-1333323321303202-3330233302031121-1013033113213023)
- rules.criteria.site_local_network

<a id="canonical-3331121222223203-2101030333110211-0111033020123321-1303030013330211-0112312323230300-1222220000203232-2303230033110200-0330210223211013"></a>

Type: `["object", {}]`. Computed.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0302030212022102-0022203320220102-2301303202022232-1233212230311300-3123132101033331-0113312012012101-1200202130101220-3132111313230213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rules.criteria.tcp` properties

Breadcrumbs:

- [xcsh_nat_policy](../data-sources/nat_policy.md#canonical-2211333230110221-1023302013201300-2020112103300231-1033030223223010-0020003220130132-3323021301233201-1030131302223322-2302033131310121)
- [Property reference](data-sources--nat_policy--reference--group-001.md#canonical-0313000023303013-1312311300100203-2312121103033010-0011202012203121-2320110020132021-0013320012122120-1332300232020201-2300013013202121)
- [rules](data-sources--nat_policy--reference--group-001.md#canonical-0211110210131101-2303300303122112-0031321323203023-2003313333331231-3032223122221033-1121322010202212-2233233201200212-2000311112203230)
- [rules.criteria](data-sources--nat_policy--reference--group-001.md#canonical-3232031321200303-0000110111113213-2023132130333021-2321302103203031-1121302211233200-1333323321303202-3330233302031121-1013033113213023)
- rules.criteria.tcp

<a id="canonical-2000131203120303-2010121001103002-2221112012330033-1111011011223023-3003113323122133-3031300310330223-0102131331023213-3330223023222200"></a>

Type: `"single"`. Computed.

Action to apply on the packet if the NAT rule is applied.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3122100021201331-1010000310032210-3310221022321003-3000023210002023-1012213111102000-2033012012011233-2311030323012332-2232210322322020"></a>

### Direct properties for `rules.criteria.tcp`

- [destination_port](data-sources--nat_policy--reference--group-001.md#canonical-1011032220011012-2120123322130012-1201122202113003-2312133030022132-1212301223110310-2000130002100033-0233100101312132-1013213021023312): complete subsection reference.

- [source_port](data-sources--nat_policy--reference--group-001.md#canonical-0120022133030120-1101233122023312-3011033131103313-3302323212102310-0231000100120012-2002033230300210-1023333131311030-0010112120321101): complete subsection reference.

<a id="canonical-1011032220011012-2120123322130012-1201122202113003-2312133030022132-1212301223110310-2000130002100033-0233100101312132-1013213021023312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rules.criteria.tcp.destination_port` properties

Breadcrumbs:

- [xcsh_nat_policy](../data-sources/nat_policy.md#canonical-2211333230110221-1023302013201300-2020112103300231-1033030223223010-0020003220130132-3323021301233201-1030131302223322-2302033131310121)
- [Property reference](data-sources--nat_policy--reference--group-001.md#canonical-0313000023303013-1312311300100203-2312121103033010-0011202012203121-2320110020132021-0013320012122120-1332300232020201-2300013013202121)
- [rules](data-sources--nat_policy--reference--group-001.md#canonical-0211110210131101-2303300303122112-0031321323203023-2003313333331231-3032223122221033-1121322010202212-2233233201200212-2000311112203230)
- [rules.criteria](data-sources--nat_policy--reference--group-001.md#canonical-3232031321200303-0000110111113213-2023132130333021-2321302103203031-1121302211233200-1333323321303202-3330233302031121-1013033113213023)
- [rules.criteria.tcp](data-sources--nat_policy--reference--group-001.md#canonical-0302030212022102-0022203320220102-2301303202022232-1233212230311300-3123132101033331-0113312012012101-1200202130101220-3132111313230213)
- rules.criteria.tcp.destination_port

<a id="canonical-0300122223001102-2113200202111333-3121333213200231-1310302100011202-1111133021103201-1322232023103002-3313332030230000-2313113303033011"></a>

Type: `"single"`. Computed.

Port match of the request can be a range or a specific port.

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

<a id="canonical-0012120302230022-2022013313211212-2223233331010210-2233101112110031-0332321330310330-1201210202002323-1201120000220133-1112201100331222"></a>

### Direct properties for `rules.criteria.tcp.destination_port`

- [no_port_match](data-sources--nat_policy--reference--group-001.md#canonical-2123111233220213-1302011003311213-3113001000031323-0113313030333100-1023103030131203-1121330100022202-1300230313231310-0000000211102022): complete subsection reference.

<a id="canonical-3313320000121021-2032312222103132-1121112230133312-2212010201123121-1310230301003131-0030112003201231-3113321232311322-1221202331032231"></a>

<a id="canonical-3200100302320103-1201020321202203-1131332323313232-0112013213030321-3111313003032231-2223101302303123-2222230230122330-3313013011200011"></a>

#### `rules.criteria.tcp.destination_port.port` property

Type: `"number"`. Computed.

Exclusive with \[no\_port\_match port\_ranges\] Exact Port to match.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-2231210210220302-2012103203231101-1033323023211220-1110201303011330-1122332312322111-2233213000313110-3011031211023101-1331012003213301"></a>

<a id="canonical-2210133121002120-3202220310330212-1331123003300222-0133321321103300-3012332100111123-3210123100110102-1203122202122220-1001222302010103"></a>

#### `rules.criteria.tcp.destination_port.port_ranges` property

Type: `"string"`. Computed.

Exclusive with \[no\_port\_match port\] Port range to match.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-2123111233220213-1302011003311213-3113001000031323-0113313030333100-1023103030131203-1121330100022202-1300230313231310-0000000211102022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rules.criteria.tcp.destination_port.no_port_match` properties

Breadcrumbs:

- [xcsh_nat_policy](../data-sources/nat_policy.md#canonical-2211333230110221-1023302013201300-2020112103300231-1033030223223010-0020003220130132-3323021301233201-1030131302223322-2302033131310121)
- [Property reference](data-sources--nat_policy--reference--group-001.md#canonical-0313000023303013-1312311300100203-2312121103033010-0011202012203121-2320110020132021-0013320012122120-1332300232020201-2300013013202121)
- [rules](data-sources--nat_policy--reference--group-001.md#canonical-0211110210131101-2303300303122112-0031321323203023-2003313333331231-3032223122221033-1121322010202212-2233233201200212-2000311112203230)
- [rules.criteria](data-sources--nat_policy--reference--group-001.md#canonical-3232031321200303-0000110111113213-2023132130333021-2321302103203031-1121302211233200-1333323321303202-3330233302031121-1013033113213023)
- [rules.criteria.tcp](data-sources--nat_policy--reference--group-001.md#canonical-0302030212022102-0022203320220102-2301303202022232-1233212230311300-3123132101033331-0113312012012101-1200202130101220-3132111313230213)
- [rules.criteria.tcp.destination_port](data-sources--nat_policy--reference--group-001.md#canonical-1011032220011012-2120123322130012-1201122202113003-2312133030022132-1212301223110310-2000130002100033-0233100101312132-1013213021023312)
- rules.criteria.tcp.destination_port.no_port_match

<a id="canonical-0022321223023231-2321303123103231-1002223123313032-2301330003310002-2330233311220002-0230230111221310-0122333012213020-3122231213333302"></a>

Type: `["object", {}]`. Computed.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0120022133030120-1101233122023312-3011033131103313-3302323212102310-0231000100120012-2002033230300210-1023333131311030-0010112120321101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rules.criteria.tcp.source_port` properties

Breadcrumbs:

- [xcsh_nat_policy](../data-sources/nat_policy.md#canonical-2211333230110221-1023302013201300-2020112103300231-1033030223223010-0020003220130132-3323021301233201-1030131302223322-2302033131310121)
- [Property reference](data-sources--nat_policy--reference--group-001.md#canonical-0313000023303013-1312311300100203-2312121103033010-0011202012203121-2320110020132021-0013320012122120-1332300232020201-2300013013202121)
- [rules](data-sources--nat_policy--reference--group-001.md#canonical-0211110210131101-2303300303122112-0031321323203023-2003313333331231-3032223122221033-1121322010202212-2233233201200212-2000311112203230)
- [rules.criteria](data-sources--nat_policy--reference--group-001.md#canonical-3232031321200303-0000110111113213-2023132130333021-2321302103203031-1121302211233200-1333323321303202-3330233302031121-1013033113213023)
- [rules.criteria.tcp](data-sources--nat_policy--reference--group-001.md#canonical-0302030212022102-0022203320220102-2301303202022232-1233212230311300-3123132101033331-0113312012012101-1200202130101220-3132111313230213)
- rules.criteria.tcp.source_port

<a id="canonical-1002221120103133-1301221101122122-0330030102331031-3103313103012103-3120202232133211-2232010230331032-1333032202210330-0123303002011002"></a>

Type: `"single"`. Computed.

Port match of the request can be a range or a specific port.

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

<a id="canonical-3322110023023332-2301312123023300-1120113202012032-2123221331022022-1211220222110201-3220022310011003-3322323023233012-1001102301113101"></a>

### Direct properties for `rules.criteria.tcp.source_port`

- [no_port_match](data-sources--nat_policy--reference--group-001.md#canonical-1333011033332133-0011230310033120-3130120310011320-3020320312231102-1322213231312000-3332203200200020-0102003033133310-3110101302001013): complete subsection reference.

<a id="canonical-2300312310002231-0323222221320013-0332122332003331-3111210331103000-2302033301112311-0220322320033113-3022101200331233-2121312131310213"></a>

<a id="canonical-0310321221010133-1203330303311303-3101111202002320-0013110122023113-2231100212213331-3030131333221012-1100220301313011-2210221213333011"></a>

#### `rules.criteria.tcp.source_port.port` property

Type: `"number"`. Computed.

Exclusive with \[no\_port\_match port\_ranges\] Exact Port to match.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-0002313202330012-0301031023213200-2100333300101021-3120333002023111-2331233333231011-0223111100322322-2000201011102222-3001232312311311"></a>

<a id="canonical-0311123130323020-2323322303323113-1123221030322012-0023330101101121-0130001110001223-0111122330113332-2002000111131033-2321230233230312"></a>

#### `rules.criteria.tcp.source_port.port_ranges` property

Type: `"string"`. Computed.

Exclusive with \[no\_port\_match port\] Port range to match.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-1333011033332133-0011230310033120-3130120310011320-3020320312231102-1322213231312000-3332203200200020-0102003033133310-3110101302001013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rules.criteria.tcp.source_port.no_port_match` properties

Breadcrumbs:

- [xcsh_nat_policy](../data-sources/nat_policy.md#canonical-2211333230110221-1023302013201300-2020112103300231-1033030223223010-0020003220130132-3323021301233201-1030131302223322-2302033131310121)
- [Property reference](data-sources--nat_policy--reference--group-001.md#canonical-0313000023303013-1312311300100203-2312121103033010-0011202012203121-2320110020132021-0013320012122120-1332300232020201-2300013013202121)
- [rules](data-sources--nat_policy--reference--group-001.md#canonical-0211110210131101-2303300303122112-0031321323203023-2003313333331231-3032223122221033-1121322010202212-2233233201200212-2000311112203230)
- [rules.criteria](data-sources--nat_policy--reference--group-001.md#canonical-3232031321200303-0000110111113213-2023132130333021-2321302103203031-1121302211233200-1333323321303202-3330233302031121-1013033113213023)
- [rules.criteria.tcp](data-sources--nat_policy--reference--group-001.md#canonical-0302030212022102-0022203320220102-2301303202022232-1233212230311300-3123132101033331-0113312012012101-1200202130101220-3132111313230213)
- [rules.criteria.tcp.source_port](data-sources--nat_policy--reference--group-001.md#canonical-0120022133030120-1101233122023312-3011033131103313-3302323212102310-0231000100120012-2002033230300210-1023333131311030-0010112120321101)
- rules.criteria.tcp.source_port.no_port_match

<a id="canonical-1302122230010303-3221013001013100-2231001000002201-1133003022111221-2022033020321003-1012132110101113-0321211220300202-2131231202210100"></a>

Type: `["object", {}]`. Computed.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3030211000021231-0010012323012221-2213120030321012-0333212021213101-0011113222330322-2100201312100232-3322220023030011-0232310121212331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rules.criteria.udp` properties

Breadcrumbs:

- [xcsh_nat_policy](../data-sources/nat_policy.md#canonical-2211333230110221-1023302013201300-2020112103300231-1033030223223010-0020003220130132-3323021301233201-1030131302223322-2302033131310121)
- [Property reference](data-sources--nat_policy--reference--group-001.md#canonical-0313000023303013-1312311300100203-2312121103033010-0011202012203121-2320110020132021-0013320012122120-1332300232020201-2300013013202121)
- [rules](data-sources--nat_policy--reference--group-001.md#canonical-0211110210131101-2303300303122112-0031321323203023-2003313333331231-3032223122221033-1121322010202212-2233233201200212-2000311112203230)
- [rules.criteria](data-sources--nat_policy--reference--group-001.md#canonical-3232031321200303-0000110111113213-2023132130333021-2321302103203031-1121302211233200-1333323321303202-3330233302031121-1013033113213023)
- rules.criteria.udp

<a id="canonical-2233031112333113-1021120320232133-3011323130001111-0032302210111121-3331033333130002-3002021331102013-1121103310131233-1013020110030233"></a>

Type: `"single"`. Computed.

Action to apply on the packet if the NAT rule is applied.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1203311320021132-0100113221010213-3013311303200030-0010223210130132-3021223111310230-2331210313313132-1333122200113321-2302331233133212"></a>

### Direct properties for `rules.criteria.udp`

- [destination_port](data-sources--nat_policy--reference--group-001.md#canonical-2322103312012322-1302132112232110-1110123020031021-0220032131210330-3111220200222033-0101310201132011-0032023233130101-2202030332003020): complete subsection reference.

- [source_port](data-sources--nat_policy--reference--group-001.md#canonical-3122223323123122-0133220331201330-0012330103223230-1321100120211211-0203021310002301-3231111020112232-2233322002223213-1110112321202131): complete subsection reference.

<a id="canonical-2322103312012322-1302132112232110-1110123020031021-0220032131210330-3111220200222033-0101310201132011-0032023233130101-2202030332003020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rules.criteria.udp.destination_port` properties

Breadcrumbs:

- [xcsh_nat_policy](../data-sources/nat_policy.md#canonical-2211333230110221-1023302013201300-2020112103300231-1033030223223010-0020003220130132-3323021301233201-1030131302223322-2302033131310121)
- [Property reference](data-sources--nat_policy--reference--group-001.md#canonical-0313000023303013-1312311300100203-2312121103033010-0011202012203121-2320110020132021-0013320012122120-1332300232020201-2300013013202121)
- [rules](data-sources--nat_policy--reference--group-001.md#canonical-0211110210131101-2303300303122112-0031321323203023-2003313333331231-3032223122221033-1121322010202212-2233233201200212-2000311112203230)
- [rules.criteria](data-sources--nat_policy--reference--group-001.md#canonical-3232031321200303-0000110111113213-2023132130333021-2321302103203031-1121302211233200-1333323321303202-3330233302031121-1013033113213023)
- [rules.criteria.udp](data-sources--nat_policy--reference--group-001.md#canonical-3030211000021231-0010012323012221-2213120030321012-0333212021213101-0011113222330322-2100201312100232-3322220023030011-0232310121212331)
- rules.criteria.udp.destination_port

<a id="canonical-2331232212120022-3232210133013331-2033333003003011-2201103333033031-1202010131113322-0023022030222323-1233331002213112-2121310102001003"></a>

Type: `"single"`. Computed.

Port match of the request can be a range or a specific port.

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

<a id="canonical-0211133323220013-3310122300120221-2330311000333302-1103230111332222-0101300201002102-2300311302113303-1003122212031221-1210113110332021"></a>

### Direct properties for `rules.criteria.udp.destination_port`

- [no_port_match](data-sources--nat_policy--reference--group-001.md#canonical-0120300021311213-0302123011211331-0022330200101220-0310131333111120-2031011221123223-3210322203322231-1123133133112201-3110111210301230): complete subsection reference.

<a id="canonical-2211122101213110-2111102132130301-1032233303033020-0002133332130300-0000211223302011-0101001033120320-3223022032221321-1332000023223300"></a>

<a id="canonical-1331210323212011-2301001201000120-3233333212133020-0121012231302210-1122210120131103-1102012031003023-0212120112110013-0320003130022330"></a>

#### `rules.criteria.udp.destination_port.port` property

Type: `"number"`. Computed.

Exclusive with \[no\_port\_match port\_ranges\] Exact Port to match.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-3011210010333020-2331332033112233-0200233322331320-1100211302122202-0111121003222122-0220103313021232-2131003321233233-2212012321111112"></a>

<a id="canonical-1321101133303032-3311323130221001-3130232132301223-2102030231231333-1213331211030212-0102203310122020-3102020020321321-3020220101022122"></a>

#### `rules.criteria.udp.destination_port.port_ranges` property

Type: `"string"`. Computed.

Exclusive with \[no\_port\_match port\] Port range to match.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-0120300021311213-0302123011211331-0022330200101220-0310131333111120-2031011221123223-3210322203322231-1123133133112201-3110111210301230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rules.criteria.udp.destination_port.no_port_match` properties

Breadcrumbs:

- [xcsh_nat_policy](../data-sources/nat_policy.md#canonical-2211333230110221-1023302013201300-2020112103300231-1033030223223010-0020003220130132-3323021301233201-1030131302223322-2302033131310121)
- [Property reference](data-sources--nat_policy--reference--group-001.md#canonical-0313000023303013-1312311300100203-2312121103033010-0011202012203121-2320110020132021-0013320012122120-1332300232020201-2300013013202121)
- [rules](data-sources--nat_policy--reference--group-001.md#canonical-0211110210131101-2303300303122112-0031321323203023-2003313333331231-3032223122221033-1121322010202212-2233233201200212-2000311112203230)
- [rules.criteria](data-sources--nat_policy--reference--group-001.md#canonical-3232031321200303-0000110111113213-2023132130333021-2321302103203031-1121302211233200-1333323321303202-3330233302031121-1013033113213023)
- [rules.criteria.udp](data-sources--nat_policy--reference--group-001.md#canonical-3030211000021231-0010012323012221-2213120030321012-0333212021213101-0011113222330322-2100201312100232-3322220023030011-0232310121212331)
- [rules.criteria.udp.destination_port](data-sources--nat_policy--reference--group-001.md#canonical-2322103312012322-1302132112232110-1110123020031021-0220032131210330-3111220200222033-0101310201132011-0032023233130101-2202030332003020)
- rules.criteria.udp.destination_port.no_port_match

<a id="canonical-3221112310032120-0113133010200132-0103322203113302-0332303223300233-2131032101322021-3023230232123303-0312010120030110-1031300100133213"></a>

Type: `["object", {}]`. Computed.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3122223323123122-0133220331201330-0012330103223230-1321100120211211-0203021310002301-3231111020112232-2233322002223213-1110112321202131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rules.criteria.udp.source_port` properties

Breadcrumbs:

- [xcsh_nat_policy](../data-sources/nat_policy.md#canonical-2211333230110221-1023302013201300-2020112103300231-1033030223223010-0020003220130132-3323021301233201-1030131302223322-2302033131310121)
- [Property reference](data-sources--nat_policy--reference--group-001.md#canonical-0313000023303013-1312311300100203-2312121103033010-0011202012203121-2320110020132021-0013320012122120-1332300232020201-2300013013202121)
- [rules](data-sources--nat_policy--reference--group-001.md#canonical-0211110210131101-2303300303122112-0031321323203023-2003313333331231-3032223122221033-1121322010202212-2233233201200212-2000311112203230)
- [rules.criteria](data-sources--nat_policy--reference--group-001.md#canonical-3232031321200303-0000110111113213-2023132130333021-2321302103203031-1121302211233200-1333323321303202-3330233302031121-1013033113213023)
- [rules.criteria.udp](data-sources--nat_policy--reference--group-001.md#canonical-3030211000021231-0010012323012221-2213120030321012-0333212021213101-0011113222330322-2100201312100232-3322220023030011-0232310121212331)
- rules.criteria.udp.source_port

<a id="canonical-2123013333203131-3022300223101203-3000023122100031-1013202112031030-0320012300322122-3220200312212202-1020300223222130-2031223011132310"></a>

Type: `"single"`. Computed.

Port match of the request can be a range or a specific port.

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

<a id="canonical-0211231120233032-0101011133023130-1210303003113020-0232223110120222-1230230133030212-1300230011220122-1000203333230221-0133203113301132"></a>

### Direct properties for `rules.criteria.udp.source_port`

- [no_port_match](data-sources--nat_policy--reference--group-001.md#canonical-1021000011300311-1012310012021100-3130333011231131-1022311113213202-3230113220320203-3311032320110200-0311301022111303-0211331330332221): complete subsection reference.

<a id="canonical-3322233002131211-1323331310001132-1320312232311200-3203110103031223-2013330102233201-1223002231033033-3033033133201212-0300311133312113"></a>

<a id="canonical-1202221332021223-1333320100233122-2102211202110002-2011133321021101-3311330230320022-2302332223331031-2033120212233000-3330112001230031"></a>

#### `rules.criteria.udp.source_port.port` property

Type: `"number"`. Computed.

Exclusive with \[no\_port\_match port\_ranges\] Exact Port to match.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-3003322233012202-2003201121330012-3111120030311310-1121213000203332-1001200210012120-3203310213212111-2122000001323212-3210122131331033"></a>

<a id="canonical-1122002220001131-3231123221210031-0133032010201232-2322333313332020-0323310012210233-1211120013303121-3113010201003111-1120300212210022"></a>

#### `rules.criteria.udp.source_port.port_ranges` property

Type: `"string"`. Computed.

Exclusive with \[no\_port\_match port\] Port range to match.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-1021000011300311-1012310012021100-3130333011231131-1022311113213202-3230113220320203-3311032320110200-0311301022111303-0211331330332221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rules.criteria.udp.source_port.no_port_match` properties

Breadcrumbs:

- [xcsh_nat_policy](../data-sources/nat_policy.md#canonical-2211333230110221-1023302013201300-2020112103300231-1033030223223010-0020003220130132-3323021301233201-1030131302223322-2302033131310121)
- [Property reference](data-sources--nat_policy--reference--group-001.md#canonical-0313000023303013-1312311300100203-2312121103033010-0011202012203121-2320110020132021-0013320012122120-1332300232020201-2300013013202121)
- [rules](data-sources--nat_policy--reference--group-001.md#canonical-0211110210131101-2303300303122112-0031321323203023-2003313333331231-3032223122221033-1121322010202212-2233233201200212-2000311112203230)
- [rules.criteria](data-sources--nat_policy--reference--group-001.md#canonical-3232031321200303-0000110111113213-2023132130333021-2321302103203031-1121302211233200-1333323321303202-3330233302031121-1013033113213023)
- [rules.criteria.udp](data-sources--nat_policy--reference--group-001.md#canonical-3030211000021231-0010012323012221-2213120030321012-0333212021213101-0011113222330322-2100201312100232-3322220023030011-0232310121212331)
- [rules.criteria.udp.source_port](data-sources--nat_policy--reference--group-001.md#canonical-3122223323123122-0133220331201330-0012330103223230-1321100120211211-0203021310002301-3231111020112232-2233322002223213-1110112321202131)
- rules.criteria.udp.source_port.no_port_match

<a id="canonical-0303200101230130-1122010020102233-1011102110003200-1131023313002333-3020230133333330-0321212231103210-1013233203223203-1002231012301112"></a>

Type: `["object", {}]`. Computed.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2210320323002232-0032213123210011-2131213200110203-1323222331000202-1100113212133120-2211000002003222-0122301232013100-1201321200321312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rules.disable_spec` properties

Breadcrumbs:

- [xcsh_nat_policy](../data-sources/nat_policy.md#canonical-2211333230110221-1023302013201300-2020112103300231-1033030223223010-0020003220130132-3323021301233201-1030131302223322-2302033131310121)
- [Property reference](data-sources--nat_policy--reference--group-001.md#canonical-0313000023303013-1312311300100203-2312121103033010-0011202012203121-2320110020132021-0013320012122120-1332300232020201-2300013013202121)
- [rules](data-sources--nat_policy--reference--group-001.md#canonical-0211110210131101-2303300303122112-0031321323203023-2003313333331231-3032223122221033-1121322010202212-2233233201200212-2000311112203230)
- rules.disable_spec

<a id="canonical-2020302132213323-3132131012202311-3330030330033321-0301033101022130-2011103211122011-1020220310001323-3000232030203112-1231212012322122"></a>

Type: `["object", {}]`. Computed.

Enable this option

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1332102110100203-3003203212301202-2010122122322333-2201132312131032-3310001132101102-2203022223233333-2323222211330023-1021103211223303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rules.enable` properties

Breadcrumbs:

- [xcsh_nat_policy](../data-sources/nat_policy.md#canonical-2211333230110221-1023302013201300-2020112103300231-1033030223223010-0020003220130132-3323021301233201-1030131302223322-2302033131310121)
- [Property reference](data-sources--nat_policy--reference--group-001.md#canonical-0313000023303013-1312311300100203-2312121103033010-0011202012203121-2320110020132021-0013320012122120-1332300232020201-2300013013202121)
- [rules](data-sources--nat_policy--reference--group-001.md#canonical-0211110210131101-2303300303122112-0031321323203023-2003313333331231-3032223122221033-1121322010202212-2233233201200212-2000311112203230)
- rules.enable

<a id="canonical-0122031000202303-0022302233103203-2012232021231003-1200121120101212-1233220221023031-1002021201222310-3031312302231131-1131033112221132"></a>

Type: `["object", {}]`. Computed.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3102133303312131-0221130013231131-3133131012123133-1033013320311022-2201330322003310-3232111120011023-3002022132220023-0032233132110111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rules.node_interface` properties

Breadcrumbs:

- [xcsh_nat_policy](../data-sources/nat_policy.md#canonical-2211333230110221-1023302013201300-2020112103300231-1033030223223010-0020003220130132-3323021301233201-1030131302223322-2302033131310121)
- [Property reference](data-sources--nat_policy--reference--group-001.md#canonical-0313000023303013-1312311300100203-2312121103033010-0011202012203121-2320110020132021-0013320012122120-1332300232020201-2300013013202121)
- [rules](data-sources--nat_policy--reference--group-001.md#canonical-0211110210131101-2303300303122112-0031321323203023-2003313333331231-3032223122221033-1121322010202212-2233233201200212-2000311112203230)
- rules.node_interface

<a id="canonical-3111201303131012-2302331132103100-2301220323122021-3301020331003120-0222112230102112-0213212302212023-0223310103010022-3102230211231312"></a>

Type: `"single"`. Computed.

On multinode site, this type holds the information about per node interfaces.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0210321132122112-0032131121122132-2012123302132112-1201030000232330-3301310210323301-3111322212131123-0203232020001202-3332011011331213"></a>

### Direct properties for `rules.node_interface`

- [list](data-sources--nat_policy--reference--group-001.md#canonical-0001022203200300-2112332330120330-3313122022011032-0022320332133210-2030022013033123-2011320113032231-2230002113203220-0112130310111002): complete subsection reference.

<a id="canonical-0001022203200300-2112332330120330-3313122022011032-0022320332133210-2030022013033123-2011320113032231-2230002113203220-0112130310111002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rules.node_interface.list` properties

Breadcrumbs:

- [xcsh_nat_policy](../data-sources/nat_policy.md#canonical-2211333230110221-1023302013201300-2020112103300231-1033030223223010-0020003220130132-3323021301233201-1030131302223322-2302033131310121)
- [Property reference](data-sources--nat_policy--reference--group-001.md#canonical-0313000023303013-1312311300100203-2312121103033010-0011202012203121-2320110020132021-0013320012122120-1332300232020201-2300013013202121)
- [rules](data-sources--nat_policy--reference--group-001.md#canonical-0211110210131101-2303300303122112-0031321323203023-2003313333331231-3032223122221033-1121322010202212-2233233201200212-2000311112203230)
- [rules.node_interface](data-sources--nat_policy--reference--group-001.md#canonical-3102133303312131-0221130013231131-3133131012123133-1033013320311022-2201330322003310-3232111120011023-3002022132220023-0032233132110111)
- rules.node_interface.list

<a id="canonical-3031323123332100-3220331011200233-2021300233233132-1223101120211013-2201310203212113-0232133203011313-1331131031320331-2300113131111113"></a>

Type: `"list"`. Computed.

On a multinode site, this list holds the nodes and corresponding networking\_interface.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 8,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 8,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-09T12:34:59+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "8"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "8"
  }
}
```

<a id="canonical-0111101311223310-0312102030300312-1231001230133011-2131210332113030-1011221010322231-3322031032313301-2320012122112313-1100221203322222"></a>

### Direct properties for `rules.node_interface.list`

- [interface](data-sources--nat_policy--reference--group-001.md#canonical-0000020012110211-2223000331102210-3332100313320202-1320213301133012-3202331303012132-3223002112200113-2111321100122012-0322332330031001): complete subsection reference.

<a id="canonical-3103003332232231-2331113133313333-1220122202211033-3230100330010002-1011313231023233-3223223002211333-2222232021102320-2100201033003022"></a>

<a id="canonical-0312201230311323-1003123002101311-0312231223012002-1032210101121112-2003002200001030-1113100030202131-0200021023132331-1021203122133002"></a>

#### `rules.node_interface.list.node` property

Type: `"string"`. Computed.

Node. Node name on this site.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-0000020012110211-2223000331102210-3332100313320202-1320213301133012-3202331303012132-3223002112200113-2111321100122012-0322332330031001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rules.node_interface.list.interface` properties

Breadcrumbs:

- [xcsh_nat_policy](../data-sources/nat_policy.md#canonical-2211333230110221-1023302013201300-2020112103300231-1033030223223010-0020003220130132-3323021301233201-1030131302223322-2302033131310121)
- [Property reference](data-sources--nat_policy--reference--group-001.md#canonical-0313000023303013-1312311300100203-2312121103033010-0011202012203121-2320110020132021-0013320012122120-1332300232020201-2300013013202121)
- [rules](data-sources--nat_policy--reference--group-001.md#canonical-0211110210131101-2303300303122112-0031321323203023-2003313333331231-3032223122221033-1121322010202212-2233233201200212-2000311112203230)
- [rules.node_interface](data-sources--nat_policy--reference--group-001.md#canonical-3102133303312131-0221130013231131-3133131012123133-1033013320311022-2201330322003310-3232111120011023-3002022132220023-0032233132110111)
- [rules.node_interface.list](data-sources--nat_policy--reference--group-001.md#canonical-0001022203200300-2112332330120330-3313122022011032-0022320332133210-2030022013033123-2011320113032231-2230002113203220-0112130310111002)
- rules.node_interface.list.interface

<a id="canonical-0123330011112010-0120021213211000-3332232121120013-1112333031223301-1313110310133132-1233020020231100-1330032320002332-0202313202032222"></a>

Type: `"list"`. Computed.

Interface. Interface reference on this node.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-2131313221200312-1222312211300222-2200110100010221-1312122022330000-1022130120030130-1212112231310030-0031023222013322-1330311110332033"></a>

### Direct properties for `rules.node_interface.list.interface`

<a id="canonical-3000031233200122-2110120321000232-0102121102113111-0321320221031331-2130330120112230-3132001330031102-1132000030012031-2003230033212311"></a>

#### `rules.node_interface.list.interface.kind` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. 'route').

Additional upstream details:

Virtual\_host) refers to another(e.g route) then kind will hold the referred object's kind (e.g.
"route")

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-1331112333111202-3231010212032331-0131213302032000-3223201210132333-1333231131112313-2020001320030131-0111002112103231-2102302310303303"></a>

<a id="canonical-3321222121232133-1010330330220133-2321312120031132-0333023133021000-2030310201323320-2313121301100132-1310303210132033-0211211011002333"></a>

#### `rules.node_interface.list.interface.name` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-3130122222312033-0020013112302333-1230103131113312-0301232103303121-3010210113101333-3330230123032101-2033113201223003-2313130223222301"></a>

<a id="canonical-2333203120130222-0333100213100120-0022202222033111-2131132333111302-1110110010032300-2311030230311223-2031010220120113-3232100301131231"></a>

#### `rules.node_interface.list.interface.namespace` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-1220123311202130-2010123203013200-0113002332211111-3311002031033010-3231132332021211-2230321132010103-1202322010100122-0222230303122303"></a>

<a id="canonical-3311022221212020-1121021021033210-3222301301211301-1133101103031220-3300130000101303-0231103201303032-2020310030001313-0233221312230011"></a>

#### `rules.node_interface.list.interface.tenant` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-2131001200213222-0033013112003321-0121030113323102-0002110012121331-2203012223303201-0312102303320000-3110222123123330-3012000322103302"></a>

<a id="canonical-2021303231232303-3302033022301110-0032023010020010-1320220200123313-2213330202113032-3332321023102313-0033113230320000-1002130003000330"></a>

#### `rules.node_interface.list.interface.uid` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-3122233211112302-2002223232110001-0231200113302123-0202221331112123-2023313220021211-0232101321133010-3000032212300301-3222313102331311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rules.segment` properties

Breadcrumbs:

- [xcsh_nat_policy](../data-sources/nat_policy.md#canonical-2211333230110221-1023302013201300-2020112103300231-1033030223223010-0020003220130132-3323021301233201-1030131302223322-2302033131310121)
- [Property reference](data-sources--nat_policy--reference--group-001.md#canonical-0313000023303013-1312311300100203-2312121103033010-0011202012203121-2320110020132021-0013320012122120-1332300232020201-2300013013202121)
- [rules](data-sources--nat_policy--reference--group-001.md#canonical-0211110210131101-2303300303122112-0031321323203023-2003313333331231-3032223122221033-1121322010202212-2233233201200212-2000311112203230)
- rules.segment

<a id="canonical-0321013101232012-0102301113103020-2033100230301111-1013032320002221-2011331213302331-1030210121103022-2313231111310012-1111010301110013"></a>

Type: `"single"`. Computed.

Segment Reference Type. Reference to Segment Object.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1120111130013130-2112130210123203-3333101300320130-2210122233212320-1100210202011122-1333302022310121-2023131131001223-2113013003301301"></a>

### Direct properties for `rules.segment`

- [refs](data-sources--nat_policy--reference--group-001.md#canonical-3113010201301011-1030303133221331-2000221301112300-3300300310322211-1021121203302032-2202110322110031-1100123020221002-3211101221121213): complete subsection reference.

<a id="canonical-3113010201301011-1030303133221331-2000221301112300-3300300310322211-1021121203302032-2202110322110031-1100123020221002-3211101221121213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rules.segment.refs` properties

Breadcrumbs:

- [xcsh_nat_policy](../data-sources/nat_policy.md#canonical-2211333230110221-1023302013201300-2020112103300231-1033030223223010-0020003220130132-3323021301233201-1030131302223322-2302033131310121)
- [Property reference](data-sources--nat_policy--reference--group-001.md#canonical-0313000023303013-1312311300100203-2312121103033010-0011202012203121-2320110020132021-0013320012122120-1332300232020201-2300013013202121)
- [rules](data-sources--nat_policy--reference--group-001.md#canonical-0211110210131101-2303300303122112-0031321323203023-2003313333331231-3032223122221033-1121322010202212-2233233201200212-2000311112203230)
- [rules.segment](data-sources--nat_policy--reference--group-001.md#canonical-3122233211112302-2002223232110001-0231200113302123-0202221331112123-2023313220021211-0232101321133010-3000032212300301-3222313102331311)
- rules.segment.refs

<a id="canonical-2303123111311231-1111130223110222-1110130030302000-3331200322002222-1032123312231022-0032102030113312-1222211031132330-0000332313312123"></a>

Type: `"list"`. Computed.

Segment. Reference to Segment Object.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "1"
  }
}
```

<a id="canonical-2013331120111301-1201021110001021-3111011020012313-3221010000310312-0332212202102313-2100010303012221-2210330000120101-0332211202010333"></a>

### Direct properties for `rules.segment.refs`

<a id="canonical-3221023033023031-1331220230221333-0312303101311210-0013122101201331-3230020203210232-1223222320303001-2212101010011202-3333023320321030"></a>

#### `rules.segment.refs.kind` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. 'route').

Additional upstream details:

Virtual\_host) refers to another(e.g route) then kind will hold the referred object's kind (e.g.
"route")

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-2001023131002230-2032201023011022-2102022303131233-2000323111103322-0032133130210033-2221131222331200-3230302311120132-1311113030121011"></a>

<a id="canonical-3331021312223030-3231202031103122-1103322120301110-2023001002021302-0010110021022323-3313121120201311-2200200301223133-0033210330021001"></a>

#### `rules.segment.refs.name` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-1013302222021110-0032031302313201-1202020113030011-0102331330200030-2221002312023310-0101121121222322-1101121133201302-3210213223210232"></a>

<a id="canonical-0230330122130131-0033331000213003-0301303001220013-0303200200202320-2003013033131311-3233031201333231-0012022312022321-3032020011321011"></a>

#### `rules.segment.refs.namespace` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-3123220331311322-0230222020113112-3300322130302120-1133213021233021-0133002211030023-1223103122003113-3322330311102333-0201121212133012"></a>

<a id="canonical-2321003311012303-3300330133022110-0230112202133332-2231132122103222-3230302120002002-1120010030303102-2023101101120330-0033002332312133"></a>

#### `rules.segment.refs.tenant` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-2110023122102213-3213333202132200-2013012013212110-3122201122312122-1201100011331120-0032202102203032-3223103112121332-2103011000112123"></a>

<a id="canonical-3020223302211101-0202120002213320-3023211002133020-2301103203022220-2131031110321023-3232322232321032-1202030001121322-2123201302121320"></a>

#### `rules.segment.refs.uid` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-1222133003113213-2211003133231313-3030010301222022-1032232221210133-2133032033131032-2133122120012312-3112332233032010-3032123201113221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rules.virtual_network` properties

Breadcrumbs:

- [xcsh_nat_policy](../data-sources/nat_policy.md#canonical-2211333230110221-1023302013201300-2020112103300231-1033030223223010-0020003220130132-3323021301233201-1030131302223322-2302033131310121)
- [Property reference](data-sources--nat_policy--reference--group-001.md#canonical-0313000023303013-1312311300100203-2312121103033010-0011202012203121-2320110020132021-0013320012122120-1332300232020201-2300013013202121)
- [rules](data-sources--nat_policy--reference--group-001.md#canonical-0211110210131101-2303300303122112-0031321323203023-2003313333331231-3032223122221033-1121322010202212-2233233201200212-2000311112203230)
- rules.virtual_network

<a id="canonical-3023313122233333-0323320112021112-3303211323203032-1210010312103310-3213120123002022-1211122110200311-0320302003010100-1032122133112022"></a>

Type: `"single"`. Computed.

Carries the reference to virtual network.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1121221331011010-0003013020300132-0300001231113003-3010113132313321-2230023122031021-3000231103001023-0332322030212203-3032133211100323"></a>

### Direct properties for `rules.virtual_network`

- [refs](data-sources--nat_policy--reference--group-001.md#canonical-3211302212221232-1233211323331210-2223231020011010-0212211002202020-3201323002030323-1320121210232201-2012210201312112-3130230212300132): complete subsection reference.

<a id="canonical-3211302212221232-1233211323331210-2223231020011010-0212211002202020-3201323002030323-1320121210232201-2012210201312112-3130230212300132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rules.virtual_network.refs` properties

Breadcrumbs:

- [xcsh_nat_policy](../data-sources/nat_policy.md#canonical-2211333230110221-1023302013201300-2020112103300231-1033030223223010-0020003220130132-3323021301233201-1030131302223322-2302033131310121)
- [Property reference](data-sources--nat_policy--reference--group-001.md#canonical-0313000023303013-1312311300100203-2312121103033010-0011202012203121-2320110020132021-0013320012122120-1332300232020201-2300013013202121)
- [rules](data-sources--nat_policy--reference--group-001.md#canonical-0211110210131101-2303300303122112-0031321323203023-2003313333331231-3032223122221033-1121322010202212-2233233201200212-2000311112203230)
- [rules.virtual_network](data-sources--nat_policy--reference--group-001.md#canonical-1222133003113213-2211003133231313-3030010301222022-1032232221210133-2133032033131032-2133122120012312-3112332233032010-3032123201113221)
- rules.virtual_network.refs

<a id="canonical-3130010321313133-2203113322301312-3231332223320120-0012102321210222-0311023210011303-1013102102032000-2230211210111333-1321331330200110"></a>

Type: `"list"`. Computed.

Virtual Network Reference. Reference to virtual network.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-3223032033013001-2102210120303021-2130013313213301-2131102001031130-3003213211212323-2210200012020321-3303213011230032-0232201000220010"></a>

### Direct properties for `rules.virtual_network.refs`

<a id="canonical-3133222010032103-2312222322033100-1131312213032331-1122312212313201-2212010310003333-2111130201011013-2030113311301130-0230020223120221"></a>

#### `rules.virtual_network.refs.kind` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. 'route').

Additional upstream details:

Virtual\_host) refers to another(e.g route) then kind will hold the referred object's kind (e.g.
"route")

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-2212231113312032-3111321103010102-2332320221212220-2202020223313210-0021013032331233-3312002020320321-2221303230100010-3010302022100020"></a>

<a id="canonical-3030103300303102-1231021222030212-1322030332211233-1030101310111122-1002000300311221-2003021200000320-0332020301130021-2203020032003200"></a>

#### `rules.virtual_network.refs.name` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-3121322132012012-3000130030113003-0210021111210200-0310112021211121-1131010012120130-2012301121220032-1010030322100033-0120121221000221"></a>

<a id="canonical-3210310100112331-3300212003011213-0101303200200210-1103300332233201-3310031212222120-3111020133012213-3222330030000301-0122220121032001"></a>

#### `rules.virtual_network.refs.namespace` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-0013220320003321-0211020203313333-1022010023311001-1233002113213020-3131120213231220-3130223112300023-1033030030120112-1322001113000201"></a>

<a id="canonical-0121210213020223-1103002210210113-1003120231210311-2133330021231223-2003301032131301-0010303132321011-0132203323212133-0310001130213302"></a>

#### `rules.virtual_network.refs.tenant` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-1311200221332212-2130312231121313-2201013331233021-1130122102111310-3021133131310203-0103331221230033-1112122013322132-3201110021032001"></a>

<a id="canonical-0300220021030013-0220101112312333-1301132333200313-0131131022331001-3132331313022312-1113023322311222-3022323332312310-1313211212000310"></a>

#### `rules.virtual_network.refs.uid` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-1011301221111022-3101333210103002-3321113303030012-1132311110323211-0123203232212231-0220003002111330-0230131113310100-0023212320213110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `site` properties

Breadcrumbs:

- [xcsh_nat_policy](../data-sources/nat_policy.md#canonical-2211333230110221-1023302013201300-2020112103300231-1033030223223010-0020003220130132-3323021301233201-1030131302223322-2302033131310121)
- [Property reference](data-sources--nat_policy--reference--group-001.md#canonical-0313000023303013-1312311300100203-2312121103033010-0011202012203121-2320110020132021-0013320012122120-1332300232020201-2300013013202121)
- site

<a id="canonical-1313333331003210-2113112121111033-1002232130112003-0201101313311023-3310102223312210-0330010133312033-1302220120232132-2313200231213101"></a>

Type: `"single"`. Computed.

Site Reference Type. Reference to Site Object.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3220033223103133-1130233213231022-0032010303003112-0023033230233332-2331331022100000-1013101132102002-1102310202213002-1010211123202330"></a>

### Direct properties for `site`

- [refs](data-sources--nat_policy--reference--group-001.md#canonical-2032212332302131-0023213300301211-2232303211012000-0312232213330203-2230211022223200-2221033001131333-3021021223030212-3312303201020013): complete subsection reference.

<a id="canonical-2032212332302131-0023213300301211-2232303211012000-0312232213330203-2230211022223200-2221033001131333-3021021223030212-3312303201020013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `site.refs` properties

Breadcrumbs:

- [xcsh_nat_policy](../data-sources/nat_policy.md#canonical-2211333230110221-1023302013201300-2020112103300231-1033030223223010-0020003220130132-3323021301233201-1030131302223322-2302033131310121)
- [Property reference](data-sources--nat_policy--reference--group-001.md#canonical-0313000023303013-1312311300100203-2312121103033010-0011202012203121-2320110020132021-0013320012122120-1332300232020201-2300013013202121)
- [site](data-sources--nat_policy--reference--group-001.md#canonical-1011301221111022-3101333210103002-3321113303030012-1132311110323211-0123203232212231-0220003002111330-0230131113310100-0023212320213110)
- site.refs

<a id="canonical-3113301200212233-0030300123200000-1320000230030330-1133110033330003-3230212210100303-2013021230111130-2130100133313320-1320021210233303"></a>

Type: `"list"`. Computed.

Site. Reference to Site Object.

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
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  }
}
```

<a id="canonical-3122033231130212-1113032101101313-2320110232023213-0012213300022332-0002321332213321-0012111200333113-3100113333312231-0212002001030311"></a>

### Direct properties for `site.refs`

<a id="canonical-2301232201212122-0122002112103002-0100220120103303-1330211123111122-1123330332011112-2200211332011002-3222112113101123-3132032303002000"></a>

#### `site.refs.kind` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. 'route').

Additional upstream details:

Virtual\_host) refers to another(e.g route) then kind will hold the referred object's kind (e.g.
"route")

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-3233331202221322-3311313301002101-3011023101203001-3303332011311221-2031320110103012-2101110020320332-1123301130202202-1012011002313210"></a>

<a id="canonical-0233300120302110-1212302033013111-3111020010132021-1232001203113312-2131320111030302-3031020332021330-3211300322301303-3033201321132132"></a>

#### `site.refs.name` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-1101120132322203-1122311102033011-3320010323323112-0023112021031120-2022220213022002-0221112003321230-3120030203103013-1210002310103131"></a>

<a id="canonical-2220013322202320-3200312232310102-0102211332210121-0101101323033012-2030120320332002-3023222311123020-1022003103022303-0001003100023131"></a>

#### `site.refs.namespace` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-1031021001202012-0032100120233101-0001003100223302-0022331000322303-2003321221233301-0220020222002033-2130312023100312-1221311002211122"></a>

<a id="canonical-1313303122333112-1103010112221101-2003211030033101-1330131310002232-3330111202233200-1001002121033032-1300222200022133-2201201132110313"></a>

#### `site.refs.tenant` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-0101103113033031-1210210103321320-1013212233201211-2302332232223003-2332300223203223-0111113321301002-1103202223323101-0031230111022013"></a>

<a id="canonical-2322322132210333-2131123300302012-2001021332133100-2012221110232122-3321100302123230-2123112321113022-0232031330030301-3222012132331002"></a>

#### `site.refs.uid` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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
