---
page_title: "xcsh_aws_tgw_site reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_aws_tgw_site reference."
---

# xcsh_aws_tgw_site reference

<a id="canonical-1212321311102123-0132200012312320-2312231211033230-2200200312230233-3132320011023221-0111320212132020-3002000002020330-2223302203200011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0221212221131233-1012201132320032-0322332023112013-0123013232112223-3010030010010211-1021320222230112-2230013032103201-0300032321312313"></a>

## Property reference — Property reference / 032233221002 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)
- Property reference

<a id="canonical-1103022002133320-1100303232222303-2102222231303210-1113121300101011-0201111111111103-2322022201221030-0033103302002303-3001210302123120"></a>

## Direct properties — Property reference / 032233221002 / 3

<a id="canonical-2330103032003210-2020031123312023-2321010021232122-3200000300302111-0302023202102232-3223212211001321-3301022230221301-1212020302130031"></a>

<a id="canonical-2222001322121122-0022301210312032-1130333332200131-3031131133031012-2321310222103101-3030110000210322-1132322223102211-0231232131203222"></a>

## annotations property — Property reference / 032233221002 / 4

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

- [aws_parameters](data-sources--aws_tgw_site--reference--group-001.md#canonical-3311120101031201-0112330033331111-0330012133130331-2122223112223133-1300102323001230-2222231003110310-3112001230233232-2000103311201302): complete subsection reference.

- [block_all_services](data-sources--aws_tgw_site--reference--group-002.md#canonical-3121133133112002-0033332221002003-2013023112031223-1233010230303323-2212120132203032-2310303231211001-0211233302000202-0000211330233312): complete subsection reference.

- [blocked_services](data-sources--aws_tgw_site--reference--group-002.md#canonical-2100121132300313-3112323000231331-3322303011221212-3103232201110033-2023130132111022-0020113011113210-0001230331021232-0022021332322021): complete subsection reference.

- [coordinates](data-sources--aws_tgw_site--reference--group-002.md#canonical-3320203122001123-1321021122332021-0300001203322000-1032311002033301-2121100322032201-1222120121102201-2210100212312030-1023311021213310): complete subsection reference.

- [custom_dns](data-sources--aws_tgw_site--reference--group-002.md#canonical-1011223223020320-0302302203312331-1311132331323301-3220231231123022-3310000100111101-2322211321123200-3220013010110222-2213033012101301): complete subsection reference.

- [default_blocked_services](data-sources--aws_tgw_site--reference--group-002.md#canonical-1212131321112002-1213201223220020-0101122010302212-1111013310121030-0020210013331103-1030212031001122-2310211012032310-0030321010032123): complete subsection reference.

<a id="canonical-3021031000112000-1310033321013201-3313031320312022-2310101311021121-0300132102311213-0001112202320132-0122022202020003-2321302333332201"></a>

<a id="canonical-1002002312032100-1023212233130211-1023030300310330-1202302013122303-3220101232131013-3012032232130020-1313131120020102-1213113120212222"></a>

## description property — Property reference / 032233221002 / 5

Type: `"string"`. Computed.

Description of the AWSTGWSite.

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

- [direct_connect_disabled](data-sources--aws_tgw_site--reference--group-002.md#canonical-0122003113200213-2133101112002122-3103022230022203-0333103101220331-3021302000323233-3110133023222210-2033301133212010-3100030200223020): complete subsection reference.

- [direct_connect_enabled](data-sources--aws_tgw_site--reference--group-002.md#canonical-0312112211011132-0333103021031333-1131302121312230-2113231022311313-1112020011110330-3313230111300202-1031033033320221-1032311103133300): complete subsection reference.

<a id="canonical-3210310030001001-1313202330022302-0302033301001012-1013112212022310-1123121233012213-1120121232031223-1032031032130303-3102130100100230"></a>

<a id="canonical-1231330213010113-1202133202210103-3132111031322211-1103001010213231-3323323232033032-2120032223232000-3223132020232030-3300131000113112"></a>

## ID property — Property reference / 032233221002 / 6

Type: `"string"`. Computed.

Unique identifier for the resource.

- [kubernetes_upgrade_drain](data-sources--aws_tgw_site--reference--group-002.md#canonical-0322103200110120-1212132221101122-3031220203231011-0212323213322303-1120210312111203-2300021113210022-2033113122011000-2322120103122230): complete subsection reference.

<a id="canonical-1030020030303212-1331023123103020-2311013113003030-1012203323303200-3010211312121303-3210301010111133-2030311023000320-3100320021100222"></a>

<a id="canonical-1313312131101300-2201333111002321-1031122333332123-0230331033132322-2223212110221122-0030133302202133-2012021200210002-1102203103300220"></a>

## labels property — Property reference / 032233221002 / 7

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

- [log_receiver](data-sources--aws_tgw_site--reference--group-002.md#canonical-2020021301323331-3001033223203323-2332033222202220-3311033003321302-1031013201223320-2333333000113313-1000012203232223-3002210130211211): complete subsection reference.

- [logs_streaming_disabled](data-sources--aws_tgw_site--reference--group-002.md#canonical-1113021113220301-2010123223000213-3012113321101301-3220123032012310-0000133322303120-1033130320232331-1300213221112233-1232321020023000): complete subsection reference.

<a id="canonical-0011202103312201-2033001323313331-0221002211201320-3230233231202111-0130032322210133-3200313130020001-2303302030112133-0333211012201102"></a>

<a id="canonical-2301313121030021-0202020303200202-0023300233323210-1112323100012331-3120001101011332-2003202032331221-1123211013322023-3233311311111110"></a>

## name property — Property reference / 032233221002 / 8

Type: `"string"`. Required.

Name of the AWSTGWSite.

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

<a id="canonical-0230302131303020-2103320111231203-0003002032102103-0220223213120213-2201002031301112-2202301331232212-1213120222100010-1323231332023031"></a>

<a id="canonical-1233213222012000-3112330103132023-2202021010113230-3112322202133111-0013132202202032-1031321112101202-1031203333212323-1330310030023033"></a>

## namespace property — Property reference / 032233221002 / 9

Type: `"string"`. Required.

Namespace where the AWSTGWSite exists.

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

- [offline_survivability_mode](data-sources--aws_tgw_site--reference--group-002.md#canonical-1002101032011221-3313311332231033-1320120021332312-1030110212030332-2020111312332302-2203213132102110-0110133312331031-0330022132023210): complete subsection reference.

- [os](data-sources--aws_tgw_site--reference--group-002.md#canonical-1131211032312230-3213210020110320-2302211020330112-2302302311111200-3123111312222202-0112321312131023-1133232031010020-1131203321212130): complete subsection reference.

- [performance_enhancement_mode](data-sources--aws_tgw_site--reference--group-002.md#canonical-0330322032001013-0000120311203032-3130033011212311-1333301023200011-1202223013022233-0332323331112312-2233100022230001-1003200031300332): complete subsection reference.

- [private_connectivity](data-sources--aws_tgw_site--reference--group-002.md#canonical-2301310322233311-3020201320022232-0210231210312320-0210001101002120-0123230230101130-1201233323300232-2231101213013103-0112210022220330): complete subsection reference.

- [sw](data-sources--aws_tgw_site--reference--group-002.md#canonical-0202131032330112-1333222111331323-2111010000110002-0310200320103110-0231230201000123-3131020001033213-1033001213101002-0122303031233103): complete subsection reference.

<a id="canonical-0302023103101222-2131111213000203-0201003300032103-0330232332321203-1300002031110323-2310211303321131-0232201112300030-1310123100021211"></a>

<a id="canonical-2220113221320220-1023232133232012-2310233323002103-1320010130333110-1323000303103112-2310030033330010-1233330213101213-0003011333300101"></a>

## tags property — Property reference / 032233221002 / 10

Type: `["map", "string"]`. Computed.

AWS Tags is a label consisting of a user-defined key and value. It helps to manage, identify,
organize, search for, and filter resources in AWS console.

Upstream description:

AWS Tags is a label consisting of a user-defined key and value. It helps to manage, identify,
organize, search for, and filter resources in AWS console.

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
    "ves.io.schema.rules.map.keys.string.max_len": "127",
    "ves.io.schema.rules.map.max_pairs": "40",
    "ves.io.schema.rules.map.values.string.max_len": "255"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "127",
    "ves.io.schema.rules.map.max_pairs": "40",
    "ves.io.schema.rules.map.values.string.max_len": "255"
  }
}
```

- [tgw_security](data-sources--aws_tgw_site--reference--group-002.md#canonical-0232013222110121-3223210311032231-0332221012122031-2132020312222132-2303023010020032-2101212313223322-1203003033031332-1111320102110212): complete subsection reference.

- [vn_config](data-sources--aws_tgw_site--reference--group-002.md#canonical-0302230100311030-1210212010000200-3030120211220013-1202212101231321-1322201012322230-0300232300110210-1103133222223202-1011002112321330): complete subsection reference.

- [vpc_attachments](data-sources--aws_tgw_site--reference--group-004.md#canonical-1001011112212113-0131301132102013-2201231332231131-3111000212301222-2111013110131231-3011310233312203-2332313100213022-0323211103120303): complete subsection reference.

- [waf_signatures](data-sources--aws_tgw_site--reference--group-004.md#canonical-0033210300111000-2300033032033230-0020231020220111-0223320021223331-0012031132001010-1110012102032212-1122313103211132-2001123200330212): complete subsection reference.

<a id="canonical-0132133110300230-1011113300121012-1200002310023211-2321131100211110-1121202133032320-0311110331010213-3330101112332003-1102303012031301"></a>

## All schema paths — Property reference / 032233221002 / 11

Each exact path has one authoritative reference destination. Collection element indices are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](data-sources--aws_tgw_site--reference--group-001.md#canonical-2330103032003210-2020031123312023-2321010021232122-3200000300302111-0302023202102232-3223212211001321-3301022230221301-1212020302130031) |
| `aws_parameters` | [aws_parameters](data-sources--aws_tgw_site--reference--group-001.md#canonical-3332111212201100-1330102312211212-0010300231131013-2133303233030103-2222020321321001-2233111333023310-3211211211211201-1310230113020213) |
| `aws_parameters.admin_password` | [aws_parameters.admin_password](data-sources--aws_tgw_site--reference--group-001.md#canonical-1330231030122222-0001220132323320-2032200212031122-3310311110301121-2000111103003023-0213013212032102-1223231011220021-2100211101210010) |
| `aws_parameters.admin_password.blindfold_secret_info` | [aws_parameters.admin_password.blindfold_secret_info](data-sources--aws_tgw_site--reference--group-001.md#canonical-2032220023132303-3221330220321201-2102111303303123-2100303000100211-3030310103033121-3011333210120222-1021012001030002-3230121213031010) |
| `aws_parameters.admin_password.blindfold_secret_info.decryption_provider` | [aws_parameters.admin_password.blindfold_secret_info.decryption_provider](data-sources--aws_tgw_site--reference--group-001.md#canonical-1123330331313200-2313223100133123-0203331033320310-2131232021002122-0330333112113230-1212130023022013-2033113111311012-0322301013331012) |
| `aws_parameters.admin_password.blindfold_secret_info.location` | [aws_parameters.admin_password.blindfold_secret_info.location](data-sources--aws_tgw_site--reference--group-001.md#canonical-1313300330220211-0011203213002320-1110001001110120-2230102122231033-1133301300013013-3130221303130231-0112313023112010-3203220003001123) |
| `aws_parameters.admin_password.blindfold_secret_info.store_provider` | [aws_parameters.admin_password.blindfold_secret_info.store_provider](data-sources--aws_tgw_site--reference--group-001.md#canonical-3030020302220130-1313121013231203-0311320321032000-3221333001333202-0030103232031313-2210333022330233-2130321300213323-0332112210311103) |
| `aws_parameters.admin_password.clear_secret_info` | [aws_parameters.admin_password.clear_secret_info](data-sources--aws_tgw_site--reference--group-001.md#canonical-0102331322020110-1031312000021102-2033221020030011-3320020212012232-3230132122211220-0002001231000113-1110112011311303-0031121322112031) |
| `aws_parameters.admin_password.clear_secret_info.provider_ref` | [aws_parameters.admin_password.clear_secret_info.provider_ref](data-sources--aws_tgw_site--reference--group-001.md#canonical-1121002121123001-3321321320231110-0221300210020221-2122132022222331-1103322200312101-1110002222130113-3302232133211312-3032332233333122) |
| `aws_parameters.admin_password.clear_secret_info.url` | [aws_parameters.admin_password.clear_secret_info.url](data-sources--aws_tgw_site--reference--group-001.md#canonical-3312022033313320-0322200202213210-3110013010013203-1120303222233011-0231333313220111-0120330230212203-2223212113300020-2222033030003010) |
| `aws_parameters.aws_cred` | [aws_parameters.aws_cred](data-sources--aws_tgw_site--reference--group-001.md#canonical-3220133102232332-1123332230300110-0221131132232103-2220031310013221-3331211321303333-1131113100032100-1111121021302330-2003321100301320) |
| `aws_parameters.aws_cred.name` | [aws_parameters.aws_cred.name](data-sources--aws_tgw_site--reference--group-001.md#canonical-0110323103202032-1111121221201001-0203110000322321-0301110211001002-1132031333103300-1132232021031301-1023323202021022-0103132300301331) |
| `aws_parameters.aws_cred.namespace` | [aws_parameters.aws_cred.namespace](data-sources--aws_tgw_site--reference--group-001.md#canonical-2200333020211210-1002023123013310-3333322312012001-1210012103001011-3312111112130211-2221230101232002-1333110313030323-2310231312230322) |
| `aws_parameters.aws_cred.tenant` | [aws_parameters.aws_cred.tenant](data-sources--aws_tgw_site--reference--group-001.md#canonical-3302132302030133-1320322210011121-1021303110001111-1103231012013111-1230103133120100-0320113011230210-1320120001020131-0012022033330220) |
| `aws_parameters.aws_region` | [aws_parameters.aws_region](data-sources--aws_tgw_site--reference--group-001.md#canonical-0321103100202213-1001213001220312-2033202220232322-0333232020120021-0210123102010022-3200303221223033-1123111231221031-1132102111132012) |
| `aws_parameters.az_nodes` | [aws_parameters.az_nodes](data-sources--aws_tgw_site--reference--group-001.md#canonical-1101111332223202-3000223312031132-0111121313002110-2103022120313333-3232331200202213-2223103330203013-3032131333321031-2110311302133101) |
| `aws_parameters.az_nodes.aws_az_name` | [aws_parameters.az_nodes.aws_az_name](data-sources--aws_tgw_site--reference--group-001.md#canonical-2200210322300230-0121200110020122-0311233132011033-0330033302022113-2231302200203323-3000310303322221-3121320330322002-3323000101321001) |
| `aws_parameters.az_nodes.inside_subnet` | [aws_parameters.az_nodes.inside_subnet](data-sources--aws_tgw_site--reference--group-001.md#canonical-0312131210003312-1120101211221121-3100221230332100-1021030010002221-0002220223111313-0213201120021221-1103111011021210-0133202302203003) |
| `aws_parameters.az_nodes.inside_subnet.existing_subnet_id` | [aws_parameters.az_nodes.inside_subnet.existing_subnet_id](data-sources--aws_tgw_site--reference--group-001.md#canonical-1131333233111123-1201333100030321-1333310020102011-3303333030221223-0011022202100331-0300213201100231-2123131230002233-2021210030210320) |
| `aws_parameters.az_nodes.inside_subnet.subnet_param` | [aws_parameters.az_nodes.inside_subnet.subnet_param](data-sources--aws_tgw_site--reference--group-001.md#canonical-3322013001112011-3021300120321031-0133012301321133-0312300211013221-0303222022011202-2331121320033120-2011012023311100-0303211330100210) |
| `aws_parameters.az_nodes.inside_subnet.subnet_param.ipv4` | [aws_parameters.az_nodes.inside_subnet.subnet_param.ipv4](data-sources--aws_tgw_site--reference--group-001.md#canonical-2333231132130323-2303121233101133-2002302302003232-3003023123002030-1112001003021220-2012310311033310-2201320321132312-0023123110123012) |
| `aws_parameters.az_nodes.outside_subnet` | [aws_parameters.az_nodes.outside_subnet](data-sources--aws_tgw_site--reference--group-001.md#canonical-3022300030033130-1121203222003300-2230102101322323-3121002213122013-1023001113110310-3311322031213202-0001230012332211-1002112013300033) |
| `aws_parameters.az_nodes.outside_subnet.existing_subnet_id` | [aws_parameters.az_nodes.outside_subnet.existing_subnet_id](data-sources--aws_tgw_site--reference--group-001.md#canonical-3211223201312103-3313100320121120-0122323331012030-0022130210121010-3133121101100212-3002302011202132-2100002121222231-2013303331010030) |
| `aws_parameters.az_nodes.outside_subnet.subnet_param` | [aws_parameters.az_nodes.outside_subnet.subnet_param](data-sources--aws_tgw_site--reference--group-001.md#canonical-1130322131311001-1002113223313233-0220201012032321-3303102132113202-0033303312220321-1102110201111201-1030121330333030-2011102231001130) |
| `aws_parameters.az_nodes.outside_subnet.subnet_param.ipv4` | [aws_parameters.az_nodes.outside_subnet.subnet_param.ipv4](data-sources--aws_tgw_site--reference--group-001.md#canonical-1002121220112201-2103130310301311-1201201301102123-2013100132003333-1200021232022102-3222232312212131-2320231223221223-0120321212132113) |
| `aws_parameters.az_nodes.reserved_inside_subnet` | [aws_parameters.az_nodes.reserved_inside_subnet](data-sources--aws_tgw_site--reference--group-001.md#canonical-2202321103103310-2121123330102022-2333002300103221-1003312211231033-2320333022223002-3110210111332020-2311203003021123-2020233313222033) |
| `aws_parameters.az_nodes.workload_subnet` | [aws_parameters.az_nodes.workload_subnet](data-sources--aws_tgw_site--reference--group-001.md#canonical-2111032322322011-2321033001021233-2202000210232213-2233200212302300-3230132232222323-0232211031131130-2023311122002120-1212332002310002) |
| `aws_parameters.az_nodes.workload_subnet.existing_subnet_id` | [aws_parameters.az_nodes.workload_subnet.existing_subnet_id](data-sources--aws_tgw_site--reference--group-001.md#canonical-2323120131233121-0213313303020121-1012030103310113-1313023112000312-3012121103131130-2232312032220322-2031103010321121-1011021132001010) |
| `aws_parameters.az_nodes.workload_subnet.subnet_param` | [aws_parameters.az_nodes.workload_subnet.subnet_param](data-sources--aws_tgw_site--reference--group-001.md#canonical-0132330200131113-2010122310001312-0330301001322332-1232021111220022-3300121121013123-0131223233313300-0212220210010021-1133030213033133) |
| `aws_parameters.az_nodes.workload_subnet.subnet_param.ipv4` | [aws_parameters.az_nodes.workload_subnet.subnet_param.ipv4](data-sources--aws_tgw_site--reference--group-001.md#canonical-3012111010100301-3012032211222203-2300223222022030-3210103323323213-2333231001332313-0020313000223202-2120001102100332-0031223132332021) |
| `aws_parameters.custom_security_group` | [aws_parameters.custom_security_group](data-sources--aws_tgw_site--reference--group-001.md#canonical-0211103031320333-2012211311010323-3312130011312130-0012022230012112-0230122320303133-2233311233203023-2131000233321011-1021131200201313) |
| `aws_parameters.custom_security_group.inside_security_group_id` | [aws_parameters.custom_security_group.inside_security_group_id](data-sources--aws_tgw_site--reference--group-001.md#canonical-1100202332213102-1313031231213130-1131120130303111-0112100301112101-2001323123232300-3313002023011222-2033301331310012-0331020202122212) |
| `aws_parameters.custom_security_group.outside_security_group_id` | [aws_parameters.custom_security_group.outside_security_group_id](data-sources--aws_tgw_site--reference--group-001.md#canonical-2003120203131112-2311301101331331-3033332332232010-0011113212013030-1120012230332313-0322111133031021-1303210223323030-1321013112322322) |
| `aws_parameters.disable_encryption` | [aws_parameters.disable_encryption](data-sources--aws_tgw_site--reference--group-001.md#canonical-1100201031001211-3102020200012001-3311002022202321-1320333221221202-1020012021111330-2312030010003231-2232030000111331-3213031001023322) |
| `aws_parameters.disable_internet_vip` | [aws_parameters.disable_internet_vip](data-sources--aws_tgw_site--reference--group-001.md#canonical-2301213032023113-3223313013232230-2122000200113210-0333103311302230-0130100132300313-1302010012331121-2331030130132030-3011310322020320) |
| `aws_parameters.disk_size` | [aws_parameters.disk_size](data-sources--aws_tgw_site--reference--group-001.md#canonical-3012132010013313-2101102221002122-3013133013300233-0313032333103202-3100333001021010-2022013232003010-0301021303222300-3111322231033211) |
| `aws_parameters.enable_encryption` | [aws_parameters.enable_encryption](data-sources--aws_tgw_site--reference--group-001.md#canonical-1023000101100030-0132023223122001-0100332011031111-3313212000120332-0002200120211121-3232032200032330-0012322012133100-3311210022223113) |
| `aws_parameters.enable_encryption.kms_key_id` | [aws_parameters.enable_encryption.kms_key_id](data-sources--aws_tgw_site--reference--group-001.md#canonical-1223211020123121-1332121101302312-1130303221133321-3023113021000332-2022230312023021-1202231212230013-0321010330210033-1310233031032012) |
| `aws_parameters.enable_internet_vip` | [aws_parameters.enable_internet_vip](data-sources--aws_tgw_site--reference--group-001.md#canonical-0312201001302320-3302001312323332-1111132330010322-2222222013202010-1103310312233320-2312322131311102-2332230211223031-3310103120013131) |
| `aws_parameters.existing_tgw` | [aws_parameters.existing_tgw](data-sources--aws_tgw_site--reference--group-001.md#canonical-3301211022320302-1023200231031032-0000201231223332-0121301310113002-3120001111123132-2111333001133210-1313003230330221-0303232032122000) |
| `aws_parameters.existing_tgw.tgw_asn` | [aws_parameters.existing_tgw.tgw_asn](data-sources--aws_tgw_site--reference--group-001.md#canonical-2022111301330033-2311032322203330-2210203013232000-3320333333310310-0100213220002031-2012101311011100-3322333300333322-3121222021303333) |
| `aws_parameters.existing_tgw.tgw_id` | [aws_parameters.existing_tgw.tgw_id](data-sources--aws_tgw_site--reference--group-001.md#canonical-0232031220112233-1221130131000323-1301230330212321-2330131310031113-2303213203302230-0131310101012321-2133323110332311-0102320132001230) |
| `aws_parameters.existing_tgw.volterra_site_asn` | [aws_parameters.existing_tgw.volterra_site_asn](data-sources--aws_tgw_site--reference--group-001.md#canonical-3032333013310113-3101020221121001-3220021001201321-2222132000223203-0313312322320331-0003111001202021-3213312122202230-1020331002211122) |
| `aws_parameters.f5xc_security_group` | [aws_parameters.f5xc_security_group](data-sources--aws_tgw_site--reference--group-001.md#canonical-0333212322020010-1231220133203132-3110002213030232-0021301011321031-3301221333232012-2011220123022311-2000030220322120-1330320102012120) |
| `aws_parameters.instance_type` | [aws_parameters.instance_type](data-sources--aws_tgw_site--reference--group-001.md#canonical-1030012310020131-2311011230112033-3200323020102113-1323330023111023-0133122010203323-3201031131311223-1310313213202223-0333221110213110) |
| `aws_parameters.new_tgw` | [aws_parameters.new_tgw](data-sources--aws_tgw_site--reference--group-001.md#canonical-3030002000022203-2001200200323322-3111132201320132-3313120231103022-1123312203032011-3023033321111002-3300023020023313-2122213313213200) |
| `aws_parameters.new_tgw.system_generated` | [aws_parameters.new_tgw.system_generated](data-sources--aws_tgw_site--reference--group-001.md#canonical-0000201331313213-3020311311101323-3210231302033130-1232112103012020-2001113031101212-2203023012032132-0222021230211330-2002212013023203) |
| `aws_parameters.new_tgw.user_assigned` | [aws_parameters.new_tgw.user_assigned](data-sources--aws_tgw_site--reference--group-001.md#canonical-1102002221000233-2203002321323021-2133130010122230-2023131320300121-2312023033110312-1230031210101331-2011011130223111-1313113002030033) |
| `aws_parameters.new_tgw.user_assigned.tgw_asn` | [aws_parameters.new_tgw.user_assigned.tgw_asn](data-sources--aws_tgw_site--reference--group-001.md#canonical-1112221010133212-1021130310202000-1211203230032311-1101313330323120-1010131320110303-3121013110113303-3011331321020023-0000220332120013) |
| `aws_parameters.new_tgw.user_assigned.volterra_site_asn` | [aws_parameters.new_tgw.user_assigned.volterra_site_asn](data-sources--aws_tgw_site--reference--group-001.md#canonical-1321112303323020-2330012301302102-1212221223310002-0221201231203031-0130023223013202-1230301122311032-0331321131020212-3302312012120230) |
| `aws_parameters.new_vpc` | [aws_parameters.new_vpc](data-sources--aws_tgw_site--reference--group-002.md#canonical-0112231112233111-0033010130321023-2233122310001130-1033203001001231-3331233003331200-0002231332322330-1232333100323122-0230100232233102) |
| `aws_parameters.new_vpc.autogenerate` | [aws_parameters.new_vpc.autogenerate](data-sources--aws_tgw_site--reference--group-002.md#canonical-2332113211022230-3021032320121113-2023111302300120-0101212303023001-1002012100210020-3132130221131223-1110001201232212-0222232323001203) |
| `aws_parameters.new_vpc.name_tag` | [aws_parameters.new_vpc.name_tag](data-sources--aws_tgw_site--reference--group-002.md#canonical-2011222113213212-2330310332201313-2113230133003330-2213200021332221-2203111100211203-2302031012110332-1130123002320003-1121133233313131) |
| `aws_parameters.new_vpc.primary_ipv4` | [aws_parameters.new_vpc.primary_ipv4](data-sources--aws_tgw_site--reference--group-002.md#canonical-0001210313003100-0232001313303120-2333231110002223-0121203201203122-1310113210221201-2313231011332102-0123003103203113-2003200302022221) |
| `aws_parameters.no_worker_nodes` | [aws_parameters.no_worker_nodes](data-sources--aws_tgw_site--reference--group-002.md#canonical-3203003320212301-1211312102132333-0020320012103133-3132112000331121-2132222112010220-2001033122103123-0010230113101110-1001231321221122) |
| `aws_parameters.nodes_per_az` | [aws_parameters.nodes_per_az](data-sources--aws_tgw_site--reference--group-001.md#canonical-1031003233122100-0002211012301200-3203133333033310-0001103123113211-1020010032032311-0232021022010020-1331113213102220-3320223313332310) |
| `aws_parameters.reserved_tgw_cidr` | [aws_parameters.reserved_tgw_cidr](data-sources--aws_tgw_site--reference--group-002.md#canonical-3003230120221223-2022101210131100-0222321322002132-3202120030331132-2200303120212312-2322211020220110-1203211221020302-1130323112221111) |
| `aws_parameters.ssh_key` | [aws_parameters.ssh_key](data-sources--aws_tgw_site--reference--group-001.md#canonical-2212322331222003-2032213220002321-2331130032122222-0120121102222012-1221023112213101-3330010313003132-3100031321230012-2233212010310101) |
| `aws_parameters.tgw_cidr` | [aws_parameters.tgw_cidr](data-sources--aws_tgw_site--reference--group-002.md#canonical-1313132033010001-2200013112203110-3211021301020102-3210132213112210-0123111101232111-1020003223020131-3012021031301030-0203232121220221) |
| `aws_parameters.tgw_cidr.ipv4` | [aws_parameters.tgw_cidr.ipv4](data-sources--aws_tgw_site--reference--group-002.md#canonical-2233130003332021-2110222203320321-2122001222220000-2031202322030213-0133011133030203-1331003313303003-2103330301130230-3122011001213231) |
| `aws_parameters.total_nodes` | [aws_parameters.total_nodes](data-sources--aws_tgw_site--reference--group-001.md#canonical-0003221222331201-2201033333213321-2230011311023003-2332000103221112-2130103333302203-0100123030003301-1230311300131102-2033112322133100) |
| `aws_parameters.vpc_id` | [aws_parameters.vpc_id](data-sources--aws_tgw_site--reference--group-001.md#canonical-2110330210333202-2203213120033331-1021020121111212-3122212121110200-3131113020300102-1120300002301101-3200130233033013-2200020322032010) |
| `block_all_services` | [block_all_services](data-sources--aws_tgw_site--reference--group-002.md#canonical-0313121231222023-3332011302231030-3001320123301303-3330232232020332-3011112101131021-2020100020112103-1300121100013133-2121033212021222) |
| `blocked_services` | [blocked_services](data-sources--aws_tgw_site--reference--group-002.md#canonical-2312023131001001-1001333003132333-2021131112102222-1203313220103300-1203132211301111-1131012000233030-1212011211200331-2002132113103320) |
| `blocked_services.blocked_service` | [blocked_services.blocked_service](data-sources--aws_tgw_site--reference--group-002.md#canonical-1300022131331211-2220302320131023-3032201012222211-2301330200201311-2010123233323313-1003023132003032-0322313122012100-2320001203332310) |
| `blocked_services.blocked_service.dns` | [blocked_services.blocked_service.dns](data-sources--aws_tgw_site--reference--group-002.md#canonical-3323323203231012-0003030211312012-2031323020211133-0230333023000100-0123001312030211-3232103121203222-1220010121100133-0333120000213312) |
| `blocked_services.blocked_service.network_type` | [blocked_services.blocked_service.network_type](data-sources--aws_tgw_site--reference--group-002.md#canonical-1103303203131011-1212122300101020-3102232033110001-2003121321123013-0102121302020101-2132310110001121-1201212132323213-0000223113022322) |
| `blocked_services.blocked_service.ssh` | [blocked_services.blocked_service.ssh](data-sources--aws_tgw_site--reference--group-002.md#canonical-1120221221301121-2001123210002011-2223032131113030-2121310010133212-2022330233012133-2220210231313120-0112113102333322-0233222222000321) |
| `blocked_services.blocked_service.web_user_interface` | [blocked_services.blocked_service.web_user_interface](data-sources--aws_tgw_site--reference--group-002.md#canonical-0120202011113322-0211010203132022-3031213112210133-2002303022203020-1212000212013103-2023221032302232-0311111013301110-1300320111131221) |
| `coordinates` | [coordinates](data-sources--aws_tgw_site--reference--group-002.md#canonical-0101123002002132-3221330333332100-1210130221302010-1300220122000203-0300330031113102-1201320123323231-0023212000321013-3220010023312202) |
| `coordinates.latitude` | [coordinates.latitude](data-sources--aws_tgw_site--reference--group-002.md#canonical-0130030110301303-2110302010301110-1233130020131323-2300233320133222-0220203023302002-0300313113211002-1321010031120122-0121100022013002) |
| `coordinates.longitude` | [coordinates.longitude](data-sources--aws_tgw_site--reference--group-002.md#canonical-2120301010302112-2011211300033332-0200233123121233-0200331303011230-1033312002012122-3332202023130103-3032203312331020-1211220011013111) |
| `custom_dns` | [custom_dns](data-sources--aws_tgw_site--reference--group-002.md#canonical-2221302212322221-1223320221022110-0220332200201012-3312332033230330-3033230112212123-1213011303033320-3203211301032013-0231010203030203) |
| `custom_dns.inside_nameserver` | [custom_dns.inside_nameserver](data-sources--aws_tgw_site--reference--group-002.md#canonical-3300133222021223-3221311321202313-0303231233112200-0222200022022301-2132120102232103-1100331202003113-2133011212230202-0013210131012302) |
| `custom_dns.outside_nameserver` | [custom_dns.outside_nameserver](data-sources--aws_tgw_site--reference--group-002.md#canonical-0302133300333302-3103011221213310-2323220131111231-2220332221001122-1020210303023202-1301031211012131-3211323030132013-3102001233110132) |
| `default_blocked_services` | [default_blocked_services](data-sources--aws_tgw_site--reference--group-002.md#canonical-2312030110333232-3310122233221321-3022221200223202-3223330011103002-2222130122132232-1003220323112212-3013201033332133-3330032030023200) |
| `description` | [description](data-sources--aws_tgw_site--reference--group-001.md#canonical-3021031000112000-1310033321013201-3313031320312022-2310101311021121-0300132102311213-0001112202320132-0122022202020003-2321302333332201) |
| `direct_connect_disabled` | [direct_connect_disabled](data-sources--aws_tgw_site--reference--group-002.md#canonical-0213011121022131-0312130230032300-2210110201103011-3233202121120321-3021130001221322-2220003033230012-1230003130100132-2023103022312300) |
| `direct_connect_enabled` | [direct_connect_enabled](data-sources--aws_tgw_site--reference--group-002.md#canonical-3111131133223020-0022120001021123-3320122232120203-2200102202121330-2100211321231332-3020012303233220-2123122101010000-0003301310111133) |
| `direct_connect_enabled.auto_asn` | [direct_connect_enabled.auto_asn](data-sources--aws_tgw_site--reference--group-002.md#canonical-3023031033312110-1223100000313302-3230110321012023-0221313032300101-0212313301323023-0000212300303101-1202301222323320-3010201300021231) |
| `direct_connect_enabled.custom_asn` | [direct_connect_enabled.custom_asn](data-sources--aws_tgw_site--reference--group-002.md#canonical-1320021202311333-0210201310233010-2132232220231003-3333131201223321-2121321202110002-1213032030120102-3031231122210031-2030113013020010) |
| `direct_connect_enabled.hosted_vifs` | [direct_connect_enabled.hosted_vifs](data-sources--aws_tgw_site--reference--group-002.md#canonical-3132033233301130-3322100201222101-3201320113032030-0132112022220232-3303113200230030-3133300123322201-2031212231330133-3013032031022032) |
| `direct_connect_enabled.hosted_vifs.site_registration_over_direct_connect` | [direct_connect_enabled.hosted_vifs.site_registration_over_direct_connect](data-sources--aws_tgw_site--reference--group-002.md#canonical-2321101003301320-0132020133333132-0201310101022302-1303010132031330-3332032102203103-0323220200012211-2312021210220302-0020311330210000) |
| `direct_connect_enabled.hosted_vifs.site_registration_over_direct_connect.cloudlink_network_name` | [direct_connect_enabled.hosted_vifs.site_registration_over_direct_connect.cloudlink_network_name](data-sources--aws_tgw_site--reference--group-002.md#canonical-2323203213210011-3013013213032101-2202113110121330-2312030232233132-3330211231013102-1303223100100012-3233311003322121-3003200011102110) |
| `direct_connect_enabled.hosted_vifs.site_registration_over_internet` | [direct_connect_enabled.hosted_vifs.site_registration_over_internet](data-sources--aws_tgw_site--reference--group-002.md#canonical-0100313110132030-3220332211321231-0222031210222120-0103203001201320-0020122311031132-2110002113011101-3330201203203203-3030312030000013) |
| `direct_connect_enabled.hosted_vifs.vif_list` | [direct_connect_enabled.hosted_vifs.vif_list](data-sources--aws_tgw_site--reference--group-002.md#canonical-0333121031232201-1332232200302310-1001210212202202-1123012102310122-1013223000010213-0030021221122333-2233001331333303-2131202220112232) |
| `direct_connect_enabled.hosted_vifs.vif_list.other_region` | [direct_connect_enabled.hosted_vifs.vif_list.other_region](data-sources--aws_tgw_site--reference--group-002.md#canonical-2231210011203332-1313212310323222-1230100030000122-2231100312221312-2100221131033012-3300232121302330-1322103101023110-2202203222323023) |
| `direct_connect_enabled.hosted_vifs.vif_list.same_as_site_region` | [direct_connect_enabled.hosted_vifs.vif_list.same_as_site_region](data-sources--aws_tgw_site--reference--group-002.md#canonical-1011200330323330-2323320003233213-0131313033300302-2101223133130123-3113300002101103-1130132111212131-2022101233332032-3011110300312020) |
| `direct_connect_enabled.hosted_vifs.vif_list.vif_id` | [direct_connect_enabled.hosted_vifs.vif_list.vif_id](data-sources--aws_tgw_site--reference--group-002.md#canonical-2002200330330311-2230213211011122-3112203120022233-1102022022011231-2202333301120221-0133133011000032-1230100222110302-3020323211313320) |
| `direct_connect_enabled.standard_vifs` | [direct_connect_enabled.standard_vifs](data-sources--aws_tgw_site--reference--group-002.md#canonical-0021132032220222-3313030201213213-2113013132001223-1322002331101133-0010201200123023-0210113210223122-0122310002230013-0201233233011023) |
| `id` | [ID](data-sources--aws_tgw_site--reference--group-001.md#canonical-3210310030001001-1313202330022302-0302033301001012-1013112212022310-1123121233012213-1120121232031223-1032031032130303-3102130100100230) |
| `kubernetes_upgrade_drain` | [kubernetes_upgrade_drain](data-sources--aws_tgw_site--reference--group-002.md#canonical-1122122331212220-3032221130130121-0313003211221012-2230033000102231-2010301113220211-2131223323213320-1032133012300310-1211323311001222) |
| `kubernetes_upgrade_drain.disable_upgrade_drain` | [kubernetes_upgrade_drain.disable_upgrade_drain](data-sources--aws_tgw_site--reference--group-002.md#canonical-1021212301023113-1001322031023123-2331000332002122-2210103203322303-0320221312311011-3121211231300203-2313212333102212-2013311031300321) |
| `kubernetes_upgrade_drain.enable_upgrade_drain` | [kubernetes_upgrade_drain.enable_upgrade_drain](data-sources--aws_tgw_site--reference--group-002.md#canonical-3303003312323121-3230301232021123-0011200130203332-2313202303331010-1221111311231320-0232233001003320-3123030210020223-2112111333111202) |
| `kubernetes_upgrade_drain.enable_upgrade_drain.disable_vega_upgrade_mode` | [kubernetes_upgrade_drain.enable_upgrade_drain.disable_vega_upgrade_mode](data-sources--aws_tgw_site--reference--group-002.md#canonical-2012012130000103-1102333031131303-2001003112030121-2233320123033111-1202331330123223-2222101111321011-1321233100230301-1223003101022221) |
| `kubernetes_upgrade_drain.enable_upgrade_drain.drain_max_unavailable_node_count` | [kubernetes_upgrade_drain.enable_upgrade_drain.drain_max_unavailable_node_count](data-sources--aws_tgw_site--reference--group-002.md#canonical-1013312200111212-1302123211112313-0320023312020132-3023231031201202-3222120323002232-1202320102023033-1001211210203322-0111321032033113) |
| `kubernetes_upgrade_drain.enable_upgrade_drain.drain_max_unavailable_node_percentage` | [kubernetes_upgrade_drain.enable_upgrade_drain.drain_max_unavailable_node_percentage](data-sources--aws_tgw_site--reference--group-002.md#canonical-3230133031301232-1222330323310331-1012011001303101-2030010222000001-2100301020321122-1221011021200113-0232202121233220-2220332200220200) |
| `kubernetes_upgrade_drain.enable_upgrade_drain.drain_node_timeout` | [kubernetes_upgrade_drain.enable_upgrade_drain.drain_node_timeout](data-sources--aws_tgw_site--reference--group-002.md#canonical-0223320123313300-1203200203010303-3033231212310313-0302113102111120-2201020113220113-1211132011032313-1001010011231223-3112011211001032) |
| `kubernetes_upgrade_drain.enable_upgrade_drain.enable_vega_upgrade_mode` | [kubernetes_upgrade_drain.enable_upgrade_drain.enable_vega_upgrade_mode](data-sources--aws_tgw_site--reference--group-002.md#canonical-1012331302012031-2212033001331030-2002301103211302-0031112311231022-2211122111112203-3321210333112302-3023013103021201-3111322130313001) |
| `labels` | [labels](data-sources--aws_tgw_site--reference--group-001.md#canonical-1030020030303212-1331023123103020-2311013113003030-1012203323303200-3010211312121303-3210301010111133-2030311023000320-3100320021100222) |
| `log_receiver` | [log_receiver](data-sources--aws_tgw_site--reference--group-002.md#canonical-3030230100032313-2200010231230310-1212303202312110-0123013031313320-2201210013021220-0221212203303210-3103200302312121-0213011232213332) |
| `log_receiver.name` | [log_receiver.name](data-sources--aws_tgw_site--reference--group-002.md#canonical-1102033123210212-1211003203000332-0031232321020330-3023210202122331-0122200320332220-2210100123030002-2001202231100220-0000111022122211) |
| `log_receiver.namespace` | [log_receiver.namespace](data-sources--aws_tgw_site--reference--group-002.md#canonical-0232000112231101-2210221322332000-3132210322200313-1233310013122120-2130311213331213-3221031302311310-0231223123302333-1322202323211011) |
| `log_receiver.tenant` | [log_receiver.tenant](data-sources--aws_tgw_site--reference--group-002.md#canonical-0020300220133220-0212313132101101-3100200031123220-1101332023221222-2231233311301010-1321012230212012-1332203133031323-1131112032232111) |
| `logs_streaming_disabled` | [logs_streaming_disabled](data-sources--aws_tgw_site--reference--group-002.md#canonical-1311031232301110-0101211333020012-2300103001323231-0300221120113121-0111103220222123-0222120000232331-1032131220021020-1031000220022031) |
| `name` | [name](data-sources--aws_tgw_site--reference--group-001.md#canonical-0011202103312201-2033001323313331-0221002211201320-3230233231202111-0130032322210133-3200313130020001-2303302030112133-0333211012201102) |
| `namespace` | [namespace](data-sources--aws_tgw_site--reference--group-001.md#canonical-0230302131303020-2103320111231203-0003002032102103-0220223213120213-2201002031301112-2202301331232212-1213120222100010-1323231332023031) |
| `offline_survivability_mode` | [offline_survivability_mode](data-sources--aws_tgw_site--reference--group-002.md#canonical-0121322102202200-1331102101131233-1223212212232333-0003220302220012-3222203311222213-0012212203333031-3200021112131013-3230321023213130) |
| `offline_survivability_mode.enable_offline_survivability_mode` | [offline_survivability_mode.enable_offline_survivability_mode](data-sources--aws_tgw_site--reference--group-002.md#canonical-0102012010001030-3220201221032100-2120133313313132-0123130011110111-3101000221211001-0010312231021102-3100100130032023-0132030020030311) |
| `offline_survivability_mode.no_offline_survivability_mode` | [offline_survivability_mode.no_offline_survivability_mode](data-sources--aws_tgw_site--reference--group-002.md#canonical-2231222223113003-3000113203203111-1012301003122221-0000113331323212-2203010210322130-2021323312022301-2103300333001023-0102222231221303) |
| `os` | [os](data-sources--aws_tgw_site--reference--group-002.md#canonical-1232012323003122-0120303332133123-2220313303021232-3032131333312031-2332331321012213-0130213020300022-2332023032222111-0213321231100301) |
| `os.default_os_version` | [os.default_os_version](data-sources--aws_tgw_site--reference--group-002.md#canonical-2023023303222013-3211003212023212-0030332022103301-1322023030320033-0123212232130033-2331311012021033-1132330010310023-3103310322022200) |
| `os.operating_system_version` | [os.operating_system_version](data-sources--aws_tgw_site--reference--group-002.md#canonical-3000032112303030-1002121020003022-0033031230011321-3111330301032002-2132332130202133-3301201213132032-1012012233313102-3322322232331010) |
| `performance_enhancement_mode` | [performance_enhancement_mode](data-sources--aws_tgw_site--reference--group-002.md#canonical-1101020121232223-3210011003013000-0223332002021320-2201131210333333-2013100221112302-2031013202200013-1321130210022200-3220232202332300) |
| `performance_enhancement_mode.perf_mode_l3_enhanced` | [performance_enhancement_mode.perf_mode_l3_enhanced](data-sources--aws_tgw_site--reference--group-002.md#canonical-3030233020320202-2123120321021312-3301010310332203-0312012120000012-2020033002111011-3130201222210002-3122331110102120-2332221031222311) |
| `performance_enhancement_mode.perf_mode_l3_enhanced.jumbo` | [performance_enhancement_mode.perf_mode_l3_enhanced.jumbo](data-sources--aws_tgw_site--reference--group-002.md#canonical-1120100000332311-0213302031131332-2130020033033212-1321032031113103-2312332111020132-1201312132100100-2131100022230023-3111333101220000) |
| `performance_enhancement_mode.perf_mode_l3_enhanced.no_jumbo` | [performance_enhancement_mode.perf_mode_l3_enhanced.no_jumbo](data-sources--aws_tgw_site--reference--group-002.md#canonical-2303321201113123-0310112320213100-0022301322230013-2212031111120202-0222112200233211-0323210013003203-3213000123031322-1012033332102033) |
| `performance_enhancement_mode.perf_mode_l7_enhanced` | [performance_enhancement_mode.perf_mode_l7_enhanced](data-sources--aws_tgw_site--reference--group-002.md#canonical-0313233110130320-1012221112013120-1320201110033212-0310132123111323-2120230311233202-2201201200311223-1310202120033202-3323211112010332) |
| `performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_disabled` | [performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_disabled](data-sources--aws_tgw_site--reference--group-002.md#canonical-3203213023101120-2213121223303013-0021110100010223-0312131212003131-2101030021100213-2331301030302033-1210122111011301-0011201003333111) |
| `performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_enabled` | [performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_enabled](data-sources--aws_tgw_site--reference--group-002.md#canonical-0021300211302203-1220210311311133-0213221020033130-0102020023033133-2131012231303302-2120021000032232-3300233020321001-1112002320220121) |
| `private_connectivity` | [private_connectivity](data-sources--aws_tgw_site--reference--group-002.md#canonical-2330311303310000-3201101003320222-1332021212103110-0001203211113110-3100222000231102-3210211331232231-0231310132333032-2201013101212203) |
| `private_connectivity.cloud_link` | [private_connectivity.cloud_link](data-sources--aws_tgw_site--reference--group-002.md#canonical-0030010303001102-3211101233101010-1031003113032013-0003010312130110-1313122011002123-0021012031012203-0322032211223232-3000130131121322) |
| `private_connectivity.cloud_link.name` | [private_connectivity.cloud_link.name](data-sources--aws_tgw_site--reference--group-002.md#canonical-1110230310311330-3030213103330201-1132113100030031-3313120002221230-0313103030122001-3100020001330322-1101313133301031-3300231230030200) |
| `private_connectivity.cloud_link.namespace` | [private_connectivity.cloud_link.namespace](data-sources--aws_tgw_site--reference--group-002.md#canonical-1130022210102202-3121323120301202-3201331203010103-1321223103002313-3110033132022301-1102300212301323-3130230302010323-2311113332301010) |
| `private_connectivity.cloud_link.tenant` | [private_connectivity.cloud_link.tenant](data-sources--aws_tgw_site--reference--group-002.md#canonical-0303001311113201-2100003213330112-2111230133002301-1011212132331303-0131311223322113-3233310020100203-3200033221011232-3211202121021303) |
| `private_connectivity.inside` | [private_connectivity.inside](data-sources--aws_tgw_site--reference--group-002.md#canonical-2113000200131023-3030031001000010-2031030333032233-1130320311133321-1001132220112220-3112020210200002-2223233211113211-0131103322131020) |
| `private_connectivity.outside` | [private_connectivity.outside](data-sources--aws_tgw_site--reference--group-002.md#canonical-1310323331313330-2313213231101102-1002212000230223-0302233232112023-2001011101231112-3011331220000223-1302130030333120-3200112231331203) |
| `sw` | [sw](data-sources--aws_tgw_site--reference--group-002.md#canonical-3121000022210322-1121332102321311-2330332310212001-3020230120010333-3233120322221131-3221030110012133-1312011032321302-1100323303132121) |
| `sw.default_sw_version` | [sw.default_sw_version](data-sources--aws_tgw_site--reference--group-002.md#canonical-0213122220113013-2113220022300013-1200232313120203-2121301030033221-3020132013131221-0310011121112011-3023131312311131-0120230221100220) |
| `sw.volterra_software_version` | [sw.volterra_software_version](data-sources--aws_tgw_site--reference--group-002.md#canonical-0330021302200232-2310211202123212-1331202201113120-2131102000102223-1310033221203321-1303033021202233-2230300133302213-0021203312101100) |
| `tags` | [tags](data-sources--aws_tgw_site--reference--group-001.md#canonical-0302023103101222-2131111213000203-0201003300032103-0330232332321203-1300002031110323-2310211303321131-0232201112300030-1310123100021211) |
| `tgw_security` | [tgw_security](data-sources--aws_tgw_site--reference--group-002.md#canonical-0023100022300302-1300301103130303-3021020032020322-1333132001230213-2103322223003320-2033221013203121-1133002230031233-3102111313311311) |
| `tgw_security.active_east_west_service_policies` | [tgw_security.active_east_west_service_policies](data-sources--aws_tgw_site--reference--group-002.md#canonical-1330121210213301-1310222323032202-3323310311102020-2200111100020032-0130333132111101-2231032131121030-1310201332001110-2121012121032111) |
| `tgw_security.active_east_west_service_policies.service_policies` | [tgw_security.active_east_west_service_policies.service_policies](data-sources--aws_tgw_site--reference--group-002.md#canonical-0023300202230132-1213113002112211-1202311310233132-0312223330202013-3230303223120001-3000130201211111-1130012002321020-3102311111032133) |
| `tgw_security.active_east_west_service_policies.service_policies.name` | [tgw_security.active_east_west_service_policies.service_policies.name](data-sources--aws_tgw_site--reference--group-002.md#canonical-2300210223323132-3313002130320201-1030201002213231-0303301201001001-3111300003120320-3020131111300101-0221220122100103-0232320210200113) |
| `tgw_security.active_east_west_service_policies.service_policies.namespace` | [tgw_security.active_east_west_service_policies.service_policies.namespace](data-sources--aws_tgw_site--reference--group-002.md#canonical-0130222230131121-1203201023221202-0010222302100210-2131231312003300-2332303113222313-1312020213310011-1031110113211013-1133032222311030) |
| `tgw_security.active_east_west_service_policies.service_policies.tenant` | [tgw_security.active_east_west_service_policies.service_policies.tenant](data-sources--aws_tgw_site--reference--group-002.md#canonical-0022003203333222-2210130213112320-1300320303103010-2021203111322003-0300211231203121-0011012331130010-2122003210222003-1113310111023330) |
| `tgw_security.active_enhanced_firewall_policies` | [tgw_security.active_enhanced_firewall_policies](data-sources--aws_tgw_site--reference--group-002.md#canonical-3100122300123201-0233210101101332-2032133310332201-0300320223332330-3200233303333021-1030110010222113-0321103313321202-3330202320330210) |
| `tgw_security.active_enhanced_firewall_policies.enhanced_firewall_policies` | [tgw_security.active_enhanced_firewall_policies.enhanced_firewall_policies](data-sources--aws_tgw_site--reference--group-002.md#canonical-3311213201332131-3213123121330331-1300121311023210-1332223320132301-0222323200222112-0111220013320122-3023300131232031-1000110333232202) |
| `tgw_security.active_enhanced_firewall_policies.enhanced_firewall_policies.name` | [tgw_security.active_enhanced_firewall_policies.enhanced_firewall_policies.name](data-sources--aws_tgw_site--reference--group-002.md#canonical-3010323312003203-1230332212330132-3203200313101303-1013321303231302-1303122010103021-1131201010202322-0330311022312110-0010113221001020) |
| `tgw_security.active_enhanced_firewall_policies.enhanced_firewall_policies.namespace` | [tgw_security.active_enhanced_firewall_policies.enhanced_firewall_policies.namespace](data-sources--aws_tgw_site--reference--group-002.md#canonical-0021201300132102-1033001333301030-2100222303012100-2012300120100133-3000300111331202-0212112023323221-2000001110213232-2130121112323313) |
| `tgw_security.active_enhanced_firewall_policies.enhanced_firewall_policies.tenant` | [tgw_security.active_enhanced_firewall_policies.enhanced_firewall_policies.tenant](data-sources--aws_tgw_site--reference--group-002.md#canonical-2331103130010323-3130111132200330-1220112112021230-1313101012330230-0322123320300202-3322200302213133-0212023012032010-2110201031210310) |
| `tgw_security.active_forward_proxy_policies` | [tgw_security.active_forward_proxy_policies](data-sources--aws_tgw_site--reference--group-002.md#canonical-0033102321332031-2332102301121223-0313232211023310-1333321222213222-2323210303021100-1232122101112302-2000133233122212-0213120201022001) |
| `tgw_security.active_forward_proxy_policies.forward_proxy_policies` | [tgw_security.active_forward_proxy_policies.forward_proxy_policies](data-sources--aws_tgw_site--reference--group-002.md#canonical-2112321233233312-1331021033300013-0121102131110002-0002213231223000-2102111010210002-2210100012200132-0000132223211122-3311302100310103) |
| `tgw_security.active_forward_proxy_policies.forward_proxy_policies.name` | [tgw_security.active_forward_proxy_policies.forward_proxy_policies.name](data-sources--aws_tgw_site--reference--group-002.md#canonical-0101232113031113-3303103321202312-3321311030230001-0210302133011113-3320003223332311-2021222320333133-0201200132023312-2013212133012211) |
| `tgw_security.active_forward_proxy_policies.forward_proxy_policies.namespace` | [tgw_security.active_forward_proxy_policies.forward_proxy_policies.namespace](data-sources--aws_tgw_site--reference--group-002.md#canonical-2320121310321232-1111130131203130-1122301313230332-0000310233130122-0230110132321011-0132003031322313-3012313021023200-1130001130313331) |
| `tgw_security.active_forward_proxy_policies.forward_proxy_policies.tenant` | [tgw_security.active_forward_proxy_policies.forward_proxy_policies.tenant](data-sources--aws_tgw_site--reference--group-002.md#canonical-2022002112133322-0001323210012012-2122131022023033-2030323322033323-3230220030202231-3131031133111101-2131031212213311-1231303211030322) |
| `tgw_security.active_network_policies` | [tgw_security.active_network_policies](data-sources--aws_tgw_site--reference--group-002.md#canonical-1222303002231310-2201021313212000-2230020330321333-0103103301323010-1311200333300131-2003333010222302-1233201021213103-0013202220231320) |
| `tgw_security.active_network_policies.network_policies` | [tgw_security.active_network_policies.network_policies](data-sources--aws_tgw_site--reference--group-002.md#canonical-1121302101300001-0202001231021112-1133232032322033-2212222312131313-0013221133130332-1202112133200120-0110000130321123-2213121112123313) |
| `tgw_security.active_network_policies.network_policies.name` | [tgw_security.active_network_policies.network_policies.name](data-sources--aws_tgw_site--reference--group-002.md#canonical-2222103201031010-2021112100232323-2312023031220012-0012202330003132-3011313120333133-1311231010020103-3023303100100332-2033232032313113) |
| `tgw_security.active_network_policies.network_policies.namespace` | [tgw_security.active_network_policies.network_policies.namespace](data-sources--aws_tgw_site--reference--group-002.md#canonical-2130022010310232-0013213301223112-1333023200203321-3300030300102313-3331312033233303-1310202032231221-2311113013201133-0021310010322203) |
| `tgw_security.active_network_policies.network_policies.tenant` | [tgw_security.active_network_policies.network_policies.tenant](data-sources--aws_tgw_site--reference--group-002.md#canonical-0000101223311230-3112013221302121-3003231312211322-2010132220131122-2103100220123033-1130020230102113-2120333013213320-2120232003320200) |
| `tgw_security.east_west_service_policy_allow_all` | [tgw_security.east_west_service_policy_allow_all](data-sources--aws_tgw_site--reference--group-002.md#canonical-3331033230110031-0230010010022130-3101101003311012-0111122010311331-2232323013303200-3321223132332001-1302302302000103-2002310323013031) |
| `tgw_security.forward_proxy_allow_all` | [tgw_security.forward_proxy_allow_all](data-sources--aws_tgw_site--reference--group-002.md#canonical-0112211023013312-3330223220001323-0232131133032102-2222132133003102-2103033103232202-3003213132210012-0120320112213113-3033320133131220) |
| `tgw_security.no_east_west_policy` | [tgw_security.no_east_west_policy](data-sources--aws_tgw_site--reference--group-002.md#canonical-2331002212110332-2000100311121000-3311310302011001-1221132123002103-1301013220332310-2100320330222211-3132010312103123-1200002303212020) |
| `tgw_security.no_forward_proxy` | [tgw_security.no_forward_proxy](data-sources--aws_tgw_site--reference--group-002.md#canonical-1023311231031313-0033101023212023-3332201011001102-1101033121233122-3233303123232221-0311331213021211-0210132002102121-2022313001210110) |
| `tgw_security.no_network_policy` | [tgw_security.no_network_policy](data-sources--aws_tgw_site--reference--group-002.md#canonical-1033230301303120-1031212130330003-1132331221223312-0331311231330122-0322302032233132-1111331031123131-2122130311310110-0232313101113211) |
| `vn_config` | [vn_config](data-sources--aws_tgw_site--reference--group-002.md#canonical-1132130212133133-1300231313303102-3012311220221030-3002121133320010-3113121310300211-1221101102322223-2121132222023331-2213031113012122) |
| `vn_config.allowed_vip_port` | [vn_config.allowed_vip_port](data-sources--aws_tgw_site--reference--group-003.md#canonical-2021313021303233-3323323013313223-0333210033112010-1232123310002001-2222321001133230-1020020221222202-0000100100203030-1002102111231323) |
| `vn_config.allowed_vip_port.custom_ports` | [vn_config.allowed_vip_port.custom_ports](data-sources--aws_tgw_site--reference--group-003.md#canonical-3313033231023322-0101121202310330-3030300332321333-2202123130101331-0222330212301312-1020022312032212-1223212310201132-3201300012133002) |
| `vn_config.allowed_vip_port.custom_ports.port_ranges` | [vn_config.allowed_vip_port.custom_ports.port_ranges](data-sources--aws_tgw_site--reference--group-003.md#canonical-0312113303023131-0302230332300302-2131130001212303-3030103022122010-3333102103131232-2001202003001203-1200012313223323-1003032020232130) |
| `vn_config.allowed_vip_port.disable_allowed_vip_port` | [vn_config.allowed_vip_port.disable_allowed_vip_port](data-sources--aws_tgw_site--reference--group-003.md#canonical-0303321021030230-3022302111323231-3333211023022112-3021032032021311-3303332103003000-1000132323112233-2302013200113123-0110103011233222) |
| `vn_config.allowed_vip_port.use_http_https_port` | [vn_config.allowed_vip_port.use_http_https_port](data-sources--aws_tgw_site--reference--group-003.md#canonical-3133013112010310-3310120103132101-3121201100233122-0232322211001033-1120012011021023-0201010133233032-2222031012320302-3230033231310010) |
| `vn_config.allowed_vip_port.use_http_port` | [vn_config.allowed_vip_port.use_http_port](data-sources--aws_tgw_site--reference--group-003.md#canonical-2333123322021212-1012023313221123-1303112311003132-2232321331112220-3200201022002112-0113013200223201-2233132203032031-0100010102023021) |
| `vn_config.allowed_vip_port.use_https_port` | [vn_config.allowed_vip_port.use_https_port](data-sources--aws_tgw_site--reference--group-003.md#canonical-1200321100121303-2301132301320123-2030011030302332-1202332213131301-2221101132330331-2112302322202111-0122101220301032-3132320113111332) |
| `vn_config.allowed_vip_port_sli` | [vn_config.allowed_vip_port_sli](data-sources--aws_tgw_site--reference--group-003.md#canonical-2012010110313302-0031030303221221-0212130003032210-0103022103101330-1121011332203210-3322231133001220-0210003233312011-2201212132202112) |
| `vn_config.allowed_vip_port_sli.custom_ports` | [vn_config.allowed_vip_port_sli.custom_ports](data-sources--aws_tgw_site--reference--group-003.md#canonical-2132203123002302-3022122030102110-2021323121303210-3003001133331223-2003231112010010-1130201331230213-1130102231222311-2220230213211223) |
| `vn_config.allowed_vip_port_sli.custom_ports.port_ranges` | [vn_config.allowed_vip_port_sli.custom_ports.port_ranges](data-sources--aws_tgw_site--reference--group-003.md#canonical-2200002113221332-3122230301311312-1333302323312111-2113110113231231-3110212101120322-1130131200133333-0121103210001313-2000022203321122) |
| `vn_config.allowed_vip_port_sli.disable_allowed_vip_port` | [vn_config.allowed_vip_port_sli.disable_allowed_vip_port](data-sources--aws_tgw_site--reference--group-003.md#canonical-3133023221300322-0022231131012001-3020331130323331-2203301301301110-0013021232212311-2011222112133221-1101201001112201-3112101023131003) |
| `vn_config.allowed_vip_port_sli.use_http_https_port` | [vn_config.allowed_vip_port_sli.use_http_https_port](data-sources--aws_tgw_site--reference--group-003.md#canonical-0120321211232223-2202313121132120-1023101232323213-3320013301010030-2103312322221212-1221302203300223-3332102232303210-0201032200100121) |
| `vn_config.allowed_vip_port_sli.use_http_port` | [vn_config.allowed_vip_port_sli.use_http_port](data-sources--aws_tgw_site--reference--group-003.md#canonical-0211200231120212-3310113121000230-0020031120211333-1221212332022333-2132221131331120-3103111210110300-2223120331003003-2100103320332000) |
| `vn_config.allowed_vip_port_sli.use_https_port` | [vn_config.allowed_vip_port_sli.use_https_port](data-sources--aws_tgw_site--reference--group-003.md#canonical-1301331312330112-1113302222101133-3311030300002003-0022112200130103-3321202103020121-2031023010110231-2323021230321222-1202322110102131) |
| `vn_config.dc_cluster_group_inside_vn` | [vn_config.dc_cluster_group_inside_vn](data-sources--aws_tgw_site--reference--group-003.md#canonical-0013221021112231-1220211022001000-2002131201333221-2213322013121012-2002133003222322-2100333301322012-1030233103201023-0201213121203021) |
| `vn_config.dc_cluster_group_inside_vn.name` | [vn_config.dc_cluster_group_inside_vn.name](data-sources--aws_tgw_site--reference--group-003.md#canonical-0130130012013132-2300233232112103-0230223031010323-0113212232330233-1313311102210112-3101133202322032-2220330303131202-3121132030312313) |
| `vn_config.dc_cluster_group_inside_vn.namespace` | [vn_config.dc_cluster_group_inside_vn.namespace](data-sources--aws_tgw_site--reference--group-003.md#canonical-2303032022333220-3320331302213202-0333111003001020-0032211013322201-0011023003331111-1233323233133021-2002123231203333-3220330233231301) |
| `vn_config.dc_cluster_group_inside_vn.tenant` | [vn_config.dc_cluster_group_inside_vn.tenant](data-sources--aws_tgw_site--reference--group-003.md#canonical-2033211012311130-0003032022110310-1132022330031332-0311321010332032-1303230331212230-0200011030201210-1331113001300102-3023301011333100) |
| `vn_config.dc_cluster_group_outside_vn` | [vn_config.dc_cluster_group_outside_vn](data-sources--aws_tgw_site--reference--group-003.md#canonical-0233133011211132-2131203201020322-3230113132223030-2002012011220003-0123130322020233-2120111202223113-1120010301320030-2020010100231012) |
| `vn_config.dc_cluster_group_outside_vn.name` | [vn_config.dc_cluster_group_outside_vn.name](data-sources--aws_tgw_site--reference--group-003.md#canonical-2133100021103131-1330030113013000-3320033222110232-0232311320301031-1313222010322120-3131212202333203-1121030122001022-3302012233011232) |
| `vn_config.dc_cluster_group_outside_vn.namespace` | [vn_config.dc_cluster_group_outside_vn.namespace](data-sources--aws_tgw_site--reference--group-003.md#canonical-2031201202102112-0122223001113232-0332332211213133-1120311111131232-1022331102113132-0201112313322223-2120133230223100-1303120301020123) |
| `vn_config.dc_cluster_group_outside_vn.tenant` | [vn_config.dc_cluster_group_outside_vn.tenant](data-sources--aws_tgw_site--reference--group-003.md#canonical-2030233232330121-3021000323212233-3130203033023301-2002123331133020-3231012331322223-2232220231123030-0003012110232021-3313320111230212) |
| `vn_config.global_network_list` | [vn_config.global_network_list](data-sources--aws_tgw_site--reference--group-003.md#canonical-0121312111011012-3323133333133203-1010313033013332-2332300333213313-2110102210211122-3310120233201320-1122212200301130-3312101302212101) |
| `vn_config.global_network_list.global_network_connections` | [vn_config.global_network_list.global_network_connections](data-sources--aws_tgw_site--reference--group-003.md#canonical-2322021111233212-0103333122321321-2110103100030123-2202020033001310-1112003010320323-1132131322232010-3020310032021301-1121132212032210) |
| `vn_config.global_network_list.global_network_connections.sli_to_global_dr` | [vn_config.global_network_list.global_network_connections.sli_to_global_dr](data-sources--aws_tgw_site--reference--group-003.md#canonical-1101003112123200-3223021222301310-1022211100330111-2301311020332022-1203331121111333-1100122031031310-0001311321331112-1200222131020302) |
| `vn_config.global_network_list.global_network_connections.sli_to_global_dr.global_vn` | [vn_config.global_network_list.global_network_connections.sli_to_global_dr.global_vn](data-sources--aws_tgw_site--reference--group-003.md#canonical-0201011223203322-1333212230332110-1100022033123123-2023310321030311-3111020120022132-1102313123011130-3203022033021010-3110201233122122) |
| `vn_config.global_network_list.global_network_connections.sli_to_global_dr.global_vn.name` | [vn_config.global_network_list.global_network_connections.sli_to_global_dr.global_vn.name](data-sources--aws_tgw_site--reference--group-003.md#canonical-0221321122023031-3213202301300223-0013212131313212-0101222133212203-2103110312222000-2031230132212200-1013130222322031-2101230020220000) |
| `vn_config.global_network_list.global_network_connections.sli_to_global_dr.global_vn.namespace` | [vn_config.global_network_list.global_network_connections.sli_to_global_dr.global_vn.namespace](data-sources--aws_tgw_site--reference--group-003.md#canonical-0113020102130003-3212212111323213-3301233112102122-2020113020112031-3202202023231213-3121031223112301-0311303311133332-0021131310033310) |
| `vn_config.global_network_list.global_network_connections.sli_to_global_dr.global_vn.tenant` | [vn_config.global_network_list.global_network_connections.sli_to_global_dr.global_vn.tenant](data-sources--aws_tgw_site--reference--group-003.md#canonical-3222221221220023-0103312111013002-3301213202312222-0330322202103012-3312311122310023-0100031212110232-2103333211123030-0233121100021111) |
| `vn_config.global_network_list.global_network_connections.slo_to_global_dr` | [vn_config.global_network_list.global_network_connections.slo_to_global_dr](data-sources--aws_tgw_site--reference--group-003.md#canonical-2230110122313132-3311303310031123-2320112113023110-2320023100231120-1000323021131230-3033212213033033-0103021030103032-3120330230222321) |
| `vn_config.global_network_list.global_network_connections.slo_to_global_dr.global_vn` | [vn_config.global_network_list.global_network_connections.slo_to_global_dr.global_vn](data-sources--aws_tgw_site--reference--group-003.md#canonical-0113203233312220-0100030312221122-3001023301121130-1033200222312220-3011111021030330-2301210302323211-0330012003201011-2033230222202331) |
| `vn_config.global_network_list.global_network_connections.slo_to_global_dr.global_vn.name` | [vn_config.global_network_list.global_network_connections.slo_to_global_dr.global_vn.name](data-sources--aws_tgw_site--reference--group-003.md#canonical-0113021023301300-1030223201011201-1232100320312021-2023121230122031-0300220002113113-2211223012110232-2102100013323221-2210211322003111) |
| `vn_config.global_network_list.global_network_connections.slo_to_global_dr.global_vn.namespace` | [vn_config.global_network_list.global_network_connections.slo_to_global_dr.global_vn.namespace](data-sources--aws_tgw_site--reference--group-003.md#canonical-2112032332220023-1133031330230200-2131131332031222-3321133233033313-1233312303200001-2020302202123300-2331030020312021-0220311203031123) |
| `vn_config.global_network_list.global_network_connections.slo_to_global_dr.global_vn.tenant` | [vn_config.global_network_list.global_network_connections.slo_to_global_dr.global_vn.tenant](data-sources--aws_tgw_site--reference--group-003.md#canonical-3233131200300320-1112232031133023-0123022322000131-1101133331000112-2201120100330211-0110301223021213-1012033302032312-2133031021122133) |
| `vn_config.inside_static_routes` | [vn_config.inside_static_routes](data-sources--aws_tgw_site--reference--group-003.md#canonical-1313032111231222-1231012302001313-3202220000221030-1033322000331202-0030110100210113-1122200020232300-3332323330233113-1122230121113222) |
| `vn_config.inside_static_routes.static_route_list` | [vn_config.inside_static_routes.static_route_list](data-sources--aws_tgw_site--reference--group-003.md#canonical-0320002011302122-2122010312320313-0213313220322210-3123112210233132-3032302102133103-3002001323232200-0203023131333333-2302210102122012) |
| `vn_config.inside_static_routes.static_route_list.custom_static_route` | [vn_config.inside_static_routes.static_route_list.custom_static_route](data-sources--aws_tgw_site--reference--group-003.md#canonical-2310021111233213-1012310201322132-3112230020320023-0311220010023201-3133100023302210-2012120011101133-3213300102332120-3022221302111111) |
| `vn_config.inside_static_routes.static_route_list.custom_static_route.attrs` | [vn_config.inside_static_routes.static_route_list.custom_static_route.attrs](data-sources--aws_tgw_site--reference--group-003.md#canonical-1332000033103023-0230330202313023-0022230300202122-1233322022202330-2300211220020133-0310223320100130-1231013301031101-1112213221232021) |
| `vn_config.inside_static_routes.static_route_list.custom_static_route.labels` | [vn_config.inside_static_routes.static_route_list.custom_static_route.labels](data-sources--aws_tgw_site--reference--group-003.md#canonical-2100122130313122-2012212301333201-1301201011012123-2010221001030233-0110020031123212-1101002212232030-1013231123033032-0022332311321213) |
| `vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop` | [vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop](data-sources--aws_tgw_site--reference--group-003.md#canonical-3120101033331320-1312230322103201-1232131023010323-1033111101322031-2001111302222330-2022311122330222-2300120120323232-1332003321333321) |
| `vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.interface` | [vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.interface](data-sources--aws_tgw_site--reference--group-003.md#canonical-2330222223012030-2332010113310222-3223122232213332-3112200212330220-0022123030011321-0130012103230330-2221121132002033-0302200211100330) |
| `vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.interface.kind` | [vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.interface.kind](data-sources--aws_tgw_site--reference--group-003.md#canonical-1222102310100101-2032210110121111-2331331223332223-2033131032112231-2110300231132011-2031032103213222-3230010322020003-0313200000331111) |
| `vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.interface.name` | [vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.interface.name](data-sources--aws_tgw_site--reference--group-003.md#canonical-3012113222122100-1002121132322213-2100123120121210-2331121121030101-0133121303023011-1202100122230032-1110112132100221-1120121133332113) |
| `vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.interface.namespace` | [vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.interface.namespace](data-sources--aws_tgw_site--reference--group-003.md#canonical-3101301103033322-0300201303200202-0303221210100322-0210121012123133-0111233130230332-3010102203312130-2112102002130103-2102032023022211) |
| `vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.interface.tenant` | [vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.interface.tenant](data-sources--aws_tgw_site--reference--group-003.md#canonical-1210303100333230-0132230002233103-2031322232320222-1111231113332012-1322121320302202-2331233301123011-1031113310012032-2010300230232110) |
| `vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.interface.uid` | [vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.interface.uid](data-sources--aws_tgw_site--reference--group-003.md#canonical-1310331122002021-0210331320003020-0120203013130322-2103223301012232-0320103021200320-0111123321301112-3011220003210203-3111032033311322) |
| `vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address` | [vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](data-sources--aws_tgw_site--reference--group-003.md#canonical-1013222313003032-0211213232202032-2133210222311210-0101222320223321-0200103130311203-1321203221111023-1102002221031300-2303302121200111) |
| `vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack` | [vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](data-sources--aws_tgw_site--reference--group-003.md#canonical-3212110100222132-0321030100231331-2032121020130133-3103233323312021-2100102030201222-2000020223303032-2132322110221231-1013303111333221) |
| `vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4` | [vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4](data-sources--aws_tgw_site--reference--group-003.md#canonical-3201323101310213-0201300303321103-1103100222221302-3121010213301012-0211121123001003-1223121103103112-2313331223011110-3022010212223203) |
| `vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4.addr` | [vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4.addr](data-sources--aws_tgw_site--reference--group-003.md#canonical-1120211020320310-2110311303323101-0110231123310302-0220213231001233-1002233322103222-1303100313231021-0302013122101030-0001111012222300) |
| `vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6` | [vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6](data-sources--aws_tgw_site--reference--group-003.md#canonical-2222101131023131-0103023101032330-0212010301002232-2332300101232131-0202111210120122-1222103020122111-0110210023222103-2231222300111011) |
| `vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6.addr` | [vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6.addr](data-sources--aws_tgw_site--reference--group-003.md#canonical-2011221310330123-1210030312201102-3312101301013302-1002232100201311-2011103131003023-2302031303021301-0310002100131221-3120121210111032) |
| `vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4` | [vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4](data-sources--aws_tgw_site--reference--group-003.md#canonical-3130333220113003-2202210202210003-0222123203223311-2033131211200003-1303203131220303-1213220002200232-3121101323110033-3031002120012333) |
| `vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4.addr` | [vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4.addr](data-sources--aws_tgw_site--reference--group-003.md#canonical-2021120302302210-3132013102020002-1301011202020203-1102231331101200-0110233300113000-3123132300300310-1101202221101023-3011031103120211) |
| `vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6` | [vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6](data-sources--aws_tgw_site--reference--group-003.md#canonical-0021203312233331-1033331332223313-2021310311210032-1110232020110231-0032020110222021-1200212320001022-2301020323230133-2330323103312211) |
| `vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6.addr` | [vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6.addr](data-sources--aws_tgw_site--reference--group-003.md#canonical-3132333302213220-2330322322111130-3132133333101111-2330320330130021-2001100331221231-1112210302020233-3331221313230032-3103020231013300) |
| `vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.type` | [vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.type](data-sources--aws_tgw_site--reference--group-003.md#canonical-0310302220123031-0202020003223230-0310021011311023-1330300102112121-0313232113220331-3230222321121310-0100231023002132-2033010033203310) |
| `vn_config.inside_static_routes.static_route_list.custom_static_route.subnets` | [vn_config.inside_static_routes.static_route_list.custom_static_route.subnets](data-sources--aws_tgw_site--reference--group-003.md#canonical-3200013132220230-3021210131110211-1022321232003030-2220221123232320-1032113230311233-3213220101322311-2101111013111131-0322202022033222) |
| `vn_config.inside_static_routes.static_route_list.custom_static_route.subnets.ipv4` | [vn_config.inside_static_routes.static_route_list.custom_static_route.subnets.ipv4](data-sources--aws_tgw_site--reference--group-003.md#canonical-1302220333000102-1013212303331210-3312322222110110-1121313100321112-2120200232022013-0012202023103103-2302332302111001-0011212121321110) |
| `vn_config.inside_static_routes.static_route_list.custom_static_route.subnets.ipv4.plen` | [vn_config.inside_static_routes.static_route_list.custom_static_route.subnets.ipv4.plen](data-sources--aws_tgw_site--reference--group-003.md#canonical-1023220021310133-3233310121023203-1123313132021232-0102132310332101-1332220000231332-1221011221112103-2101210120012331-2013003002211301) |
| `vn_config.inside_static_routes.static_route_list.custom_static_route.subnets.ipv4.prefix` | [vn_config.inside_static_routes.static_route_list.custom_static_route.subnets.ipv4.prefix](data-sources--aws_tgw_site--reference--group-003.md#canonical-1230311132220223-0302201033103122-1022300231232002-0221332131302202-3211030130211002-2121000322323221-0033020022032313-2010121110230322) |
| `vn_config.inside_static_routes.static_route_list.custom_static_route.subnets.ipv6` | [vn_config.inside_static_routes.static_route_list.custom_static_route.subnets.ipv6](data-sources--aws_tgw_site--reference--group-003.md#canonical-1320103021133030-2201231130312121-0032202113321330-0202330322202001-3210310213031302-2012222312130101-2321101320330211-0313220100010030) |
| `vn_config.inside_static_routes.static_route_list.custom_static_route.subnets.ipv6.plen` | [vn_config.inside_static_routes.static_route_list.custom_static_route.subnets.ipv6.plen](data-sources--aws_tgw_site--reference--group-003.md#canonical-2133213011200203-1011120033311032-3122211031300032-1212002331300023-3101102133233322-0331302000212311-0212230102122102-2220022003300131) |
| `vn_config.inside_static_routes.static_route_list.custom_static_route.subnets.ipv6.prefix` | [vn_config.inside_static_routes.static_route_list.custom_static_route.subnets.ipv6.prefix](data-sources--aws_tgw_site--reference--group-003.md#canonical-0123002033213010-3232030101212011-0202001333301122-2300300202321230-0310222030121302-2313213230130133-3101222312320032-3103000220121200) |
| `vn_config.inside_static_routes.static_route_list.simple_static_route` | [vn_config.inside_static_routes.static_route_list.simple_static_route](data-sources--aws_tgw_site--reference--group-003.md#canonical-0000022032200223-2111110233201011-2121311132230210-1220210010033012-0202102223121130-3002010010313330-3021123333002101-2020120022310312) |
| `vn_config.no_dc_cluster_group` | [vn_config.no_dc_cluster_group](data-sources--aws_tgw_site--reference--group-003.md#canonical-1221132013211032-2133130332031123-0020022331000110-1020112200130302-2221330030123310-2200000332132201-3323133213320200-2002033322023101) |
| `vn_config.no_global_network` | [vn_config.no_global_network](data-sources--aws_tgw_site--reference--group-003.md#canonical-0232130030022032-1310310331201023-3330012201032223-2211221012303032-1302021301000203-0031201201302110-2112021232201230-0112130210311333) |
| `vn_config.no_inside_static_routes` | [vn_config.no_inside_static_routes](data-sources--aws_tgw_site--reference--group-003.md#canonical-3211213202330233-3012313121131230-2013211112102123-2031021122221202-1011013022110021-0322312213020332-1313233123312013-0320110211001120) |
| `vn_config.no_outside_static_routes` | [vn_config.no_outside_static_routes](data-sources--aws_tgw_site--reference--group-003.md#canonical-0103130211231220-3101133330300313-1003120032333320-0221210331301022-0033011320330133-0201022133331023-3012023002131111-1333013320013122) |
| `vn_config.outside_static_routes` | [vn_config.outside_static_routes](data-sources--aws_tgw_site--reference--group-003.md#canonical-3312331323320232-1123101132211022-0310221210303331-0212202212132310-2310113100003003-2223113312321203-0233211130330023-0202100123021131) |
| `vn_config.outside_static_routes.static_route_list` | [vn_config.outside_static_routes.static_route_list](data-sources--aws_tgw_site--reference--group-003.md#canonical-0113233320300210-2031112020313320-2102000211221331-0202033322100012-2002012213332231-2220020103030331-1130002210003113-3132221101321313) |
| `vn_config.outside_static_routes.static_route_list.custom_static_route` | [vn_config.outside_static_routes.static_route_list.custom_static_route](data-sources--aws_tgw_site--reference--group-003.md#canonical-2213212111110031-3020033110300133-3222312223101012-0331213222000212-0300223130123303-0201322001212013-3310103201202133-0131103112300101) |
| `vn_config.outside_static_routes.static_route_list.custom_static_route.attrs` | [vn_config.outside_static_routes.static_route_list.custom_static_route.attrs](data-sources--aws_tgw_site--reference--group-003.md#canonical-0001303303211212-3001111213303202-3222021300102010-1103231103221211-1200001132301233-2113311333220121-3113202133223102-1201220321022233) |
| `vn_config.outside_static_routes.static_route_list.custom_static_route.labels` | [vn_config.outside_static_routes.static_route_list.custom_static_route.labels](data-sources--aws_tgw_site--reference--group-003.md#canonical-1200203223130301-3233201020010130-3130331231202133-1313120013020000-0111010303222113-0002313002132032-0120332310110033-0031212223302232) |
| `vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop` | [vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop](data-sources--aws_tgw_site--reference--group-003.md#canonical-0303303111000031-3110221030332012-3211302230221131-2023221302013320-3330221133131302-3133122201021323-2221213033133133-1133332323203332) |
| `vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.interface` | [vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.interface](data-sources--aws_tgw_site--reference--group-003.md#canonical-1202330203013031-1112313331003211-0111303132223300-2011103033201100-1321230321102111-3001021132320211-0310321003311133-1121012212230303) |
| `vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.kind` | [vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.kind](data-sources--aws_tgw_site--reference--group-003.md#canonical-1122300000320232-2111323223330210-1120011012210213-0020230302231310-3010030331012013-3122330020121132-2013312000223123-1301022232000131) |
| `vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.name` | [vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.name](data-sources--aws_tgw_site--reference--group-003.md#canonical-1003010012113301-2211123031012101-3231202212302110-3021121230122211-2123021002011130-3213302303201322-2330031230313100-2001120221101311) |
| `vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.namespace` | [vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.namespace](data-sources--aws_tgw_site--reference--group-003.md#canonical-1203111333202212-2232210323031203-1320313330032123-3323010300222323-0123233321110132-0311031022201001-0131121113301001-0332122032310220) |
| `vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.tenant` | [vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.tenant](data-sources--aws_tgw_site--reference--group-003.md#canonical-1311223023122001-3103311303102301-3302212201130312-0200013322330010-2302131321202033-2311011010110121-1123130330312110-3201032022010110) |
| `vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.uid` | [vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.uid](data-sources--aws_tgw_site--reference--group-003.md#canonical-0101101020123331-0130010332033310-1220021033101033-3332212310020202-0031020323200332-3200033331203023-1213130103030301-2131300100302221) |
| `vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address` | [vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](data-sources--aws_tgw_site--reference--group-003.md#canonical-2233233232001103-0002203203030122-2123212203102233-3221131213133022-2112300231102300-3012111103001110-2232123230331020-0210031130200321) |
| `vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack` | [vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](data-sources--aws_tgw_site--reference--group-003.md#canonical-0203002321101210-2231123310022232-2110012102133201-3323130013003103-3011010032210321-2021023330233101-2032001112302120-0000001202021130) |
| `vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4` | [vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4](data-sources--aws_tgw_site--reference--group-003.md#canonical-0310230212201012-3103130231100011-2211110123223211-2102332201211301-3212232200112221-1001223001222330-2233032210030202-2113222230021202) |
| `vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4.addr` | [vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4.addr](data-sources--aws_tgw_site--reference--group-003.md#canonical-2223103022020330-1302110210312023-2323321020023021-2311232000002013-2013011220111333-0301223212330323-3100201201130001-1113212330121333) |
| `vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6` | [vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6](data-sources--aws_tgw_site--reference--group-003.md#canonical-1221100310122323-1210112302002031-1123333322000202-1112220132100203-2131202031302033-3122310112023222-1120222031021100-1013333002231122) |
| `vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6.addr` | [vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6.addr](data-sources--aws_tgw_site--reference--group-003.md#canonical-2320321323203023-2312201211021003-2323013110122011-1312321210320121-0002033212011202-2220223201130122-3322033131332023-2021001000133222) |
| `vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4` | [vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4](data-sources--aws_tgw_site--reference--group-003.md#canonical-0010211210103132-1323322113103230-0233202212330001-0223003022112012-3011003032000221-2101201101011112-2323013201133033-0301312200322320) |
| `vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4.addr` | [vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4.addr](data-sources--aws_tgw_site--reference--group-003.md#canonical-1201323120011333-1020302002302301-3100031202203313-1120233012302232-2100010331000323-1102201313010010-2201313130002211-1010001011011123) |
| `vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6` | [vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6](data-sources--aws_tgw_site--reference--group-004.md#canonical-3220110210311121-3022012202112021-3122120111220131-2023200001102101-2311022032322213-3001131023213100-3023213301223013-3101333330023101) |
| `vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6.addr` | [vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6.addr](data-sources--aws_tgw_site--reference--group-004.md#canonical-3201112131210203-2131102211032122-2102133331002220-0003133032133033-0123303112210130-0003002232333332-3101012033332322-3012212320021323) |
| `vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.type` | [vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.type](data-sources--aws_tgw_site--reference--group-003.md#canonical-0332120320021321-0120002212111200-0321020213012333-3302102302002202-3003221012323210-2201031020223132-0220021113312022-0212100030313221) |
| `vn_config.outside_static_routes.static_route_list.custom_static_route.subnets` | [vn_config.outside_static_routes.static_route_list.custom_static_route.subnets](data-sources--aws_tgw_site--reference--group-004.md#canonical-0211132213222123-3313121302300003-1033302311201033-2101302103331323-1230011112201332-0322112133300301-0331211333132110-2020230101322322) |
| `vn_config.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4` | [vn_config.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4](data-sources--aws_tgw_site--reference--group-004.md#canonical-2110132003201010-0123021110032021-0131103030113213-1211030313121031-1210311021323310-0100020101300023-2030200301222012-3021212230223221) |
| `vn_config.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4.plen` | [vn_config.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4.plen](data-sources--aws_tgw_site--reference--group-004.md#canonical-3101030202221301-0310210022010030-2212003233000112-2322313321230121-1030132212033302-2200331333111021-1230333232231230-1110032021103213) |
| `vn_config.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4.prefix` | [vn_config.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4.prefix](data-sources--aws_tgw_site--reference--group-004.md#canonical-1002010002023332-2203202332021021-3220032113232112-2020113002223010-1321131201011022-2133021230122232-1311203333301203-0233332313323213) |
| `vn_config.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6` | [vn_config.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6](data-sources--aws_tgw_site--reference--group-004.md#canonical-1013100231330102-3203031332213110-0103010113231000-2003233002031230-3333110300223113-1210330013211000-0131222103013000-2323113031132130) |
| `vn_config.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6.plen` | [vn_config.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6.plen](data-sources--aws_tgw_site--reference--group-004.md#canonical-1332131310302023-2101212321221302-2311322023101202-1023020000013000-2310303010222232-2320330200002012-1303333201210112-2031022312310303) |
| `vn_config.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6.prefix` | [vn_config.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6.prefix](data-sources--aws_tgw_site--reference--group-004.md#canonical-1000212111231002-3320232222202000-1100300120121132-0332301013231133-0233322002232023-2100002313300123-0202003130023223-1312212030313033) |
| `vn_config.outside_static_routes.static_route_list.simple_static_route` | [vn_config.outside_static_routes.static_route_list.simple_static_route](data-sources--aws_tgw_site--reference--group-003.md#canonical-0302310312132022-1303002322032333-0013112310120111-2101132212221122-1023123220210213-1023230013002132-1301233203113031-0320233230331120) |
| `vn_config.sm_connection_public_ip` | [vn_config.sm_connection_public_ip](data-sources--aws_tgw_site--reference--group-004.md#canonical-1100310330332101-0230332232221303-2122130112022222-1310322111320331-2211022002330202-0322231113130100-0022011100120321-1020200030213102) |
| `vn_config.sm_connection_pvt_ip` | [vn_config.sm_connection_pvt_ip](data-sources--aws_tgw_site--reference--group-004.md#canonical-0003013021032322-2210300002230132-1231002011213030-3122001130312031-3131201223003311-1302201313021331-1223322112022211-1320320301021323) |
| `vpc_attachments` | [vpc_attachments](data-sources--aws_tgw_site--reference--group-004.md#canonical-2032121101122220-3001123200110112-1100220002312213-0333230202320131-0130022322132321-3303232323223223-1023210032011101-2130332210311003) |
| `vpc_attachments.vpc_list` | [vpc_attachments.vpc_list](data-sources--aws_tgw_site--reference--group-004.md#canonical-2200213132321220-0322000303311301-1201222033223222-1311001311321011-2310321200330030-0203210200102312-0202112230320030-3112323011313012) |
| `vpc_attachments.vpc_list.labels` | [vpc_attachments.vpc_list.labels](data-sources--aws_tgw_site--reference--group-004.md#canonical-3111123302102312-1011330022033313-0123230031120033-3200002230203112-1020332211232000-2021203001003301-1321003320311301-2333022111203013) |
| `vpc_attachments.vpc_list.vpc_id` | [vpc_attachments.vpc_list.vpc_id](data-sources--aws_tgw_site--reference--group-004.md#canonical-0003020000210021-2302311213221331-1222202130232200-2313023213123201-2233301033102213-2210001200131021-2033003301033323-1213122200323301) |
| `waf_signatures` | [waf_signatures](data-sources--aws_tgw_site--reference--group-004.md#canonical-2002121323312202-1100332113202321-0200320322132330-1320113233311320-1103132121103200-3321312301302332-3101302103212110-2021233200122010) |
| `waf_signatures.automatic` | [waf_signatures.automatic](data-sources--aws_tgw_site--reference--group-004.md#canonical-1303233131211010-2113312203131111-1201110023200102-2101320311203333-3311033011212212-2001012312212311-0322322101200002-0303323033101201) |
| `waf_signatures.manual` | [waf_signatures.manual](data-sources--aws_tgw_site--reference--group-004.md#canonical-1202003002301220-3032303103333101-1212030300222032-2330000212022123-0130012232212100-3331210133113330-0202132311220223-0103123211201313) |

<a id="canonical-3312230002003010-0203000030323311-0012133331232333-2130031113032031-2121211233010103-0321001311112101-1323310000010020-0010331313100312"></a>

## Next pages — Property reference / 032233221002 / 12

- [aws_parameters](data-sources--aws_tgw_site--reference--group-001.md#canonical-3311120101031201-0112330033331111-0330012133130331-2122223112223133-1300102323001230-2222231003110310-3112001230233232-2000103311201302)
- [block_all_services](data-sources--aws_tgw_site--reference--group-002.md#canonical-3121133133112002-0033332221002003-2013023112031223-1233010230303323-2212120132203032-2310303231211001-0211233302000202-0000211330233312)
- [blocked_services](data-sources--aws_tgw_site--reference--group-002.md#canonical-2100121132300313-3112323000231331-3322303011221212-3103232201110033-2023130132111022-0020113011113210-0001230331021232-0022021332322021)
- [coordinates](data-sources--aws_tgw_site--reference--group-002.md#canonical-3320203122001123-1321021122332021-0300001203322000-1032311002033301-2121100322032201-1222120121102201-2210100212312030-1023311021213310)
- [custom_dns](data-sources--aws_tgw_site--reference--group-002.md#canonical-1011223223020320-0302302203312331-1311132331323301-3220231231123022-3310000100111101-2322211321123200-3220013010110222-2213033012101301)
- [default_blocked_services](data-sources--aws_tgw_site--reference--group-002.md#canonical-1212131321112002-1213201223220020-0101122010302212-1111013310121030-0020210013331103-1030212031001122-2310211012032310-0030321010032123)
- [direct_connect_disabled](data-sources--aws_tgw_site--reference--group-002.md#canonical-0122003113200213-2133101112002122-3103022230022203-0333103101220331-3021302000323233-3110133023222210-2033301133212010-3100030200223020)
- [direct_connect_enabled](data-sources--aws_tgw_site--reference--group-002.md#canonical-0312112211011132-0333103021031333-1131302121312230-2113231022311313-1112020011110330-3313230111300202-1031033033320221-1032311103133300)
- [kubernetes_upgrade_drain](data-sources--aws_tgw_site--reference--group-002.md#canonical-0322103200110120-1212132221101122-3031220203231011-0212323213322303-1120210312111203-2300021113210022-2033113122011000-2322120103122230)
- [log_receiver](data-sources--aws_tgw_site--reference--group-002.md#canonical-2020021301323331-3001033223203323-2332033222202220-3311033003321302-1031013201223320-2333333000113313-1000012203232223-3002210130211211)
- [logs_streaming_disabled](data-sources--aws_tgw_site--reference--group-002.md#canonical-1113021113220301-2010123223000213-3012113321101301-3220123032012310-0000133322303120-1033130320232331-1300213221112233-1232321020023000)
- [offline_survivability_mode](data-sources--aws_tgw_site--reference--group-002.md#canonical-1002101032011221-3313311332231033-1320120021332312-1030110212030332-2020111312332302-2203213132102110-0110133312331031-0330022132023210)
- [os](data-sources--aws_tgw_site--reference--group-002.md#canonical-1131211032312230-3213210020110320-2302211020330112-2302302311111200-3123111312222202-0112321312131023-1133232031010020-1131203321212130)
- [performance_enhancement_mode](data-sources--aws_tgw_site--reference--group-002.md#canonical-0330322032001013-0000120311203032-3130033011212311-1333301023200011-1202223013022233-0332323331112312-2233100022230001-1003200031300332)
- [private_connectivity](data-sources--aws_tgw_site--reference--group-002.md#canonical-2301310322233311-3020201320022232-0210231210312320-0210001101002120-0123230230101130-1201233323300232-2231101213013103-0112210022220330)
- [sw](data-sources--aws_tgw_site--reference--group-002.md#canonical-0202131032330112-1333222111331323-2111010000110002-0310200320103110-0231230201000123-3131020001033213-1033001213101002-0122303031233103)
- [tgw_security](data-sources--aws_tgw_site--reference--group-002.md#canonical-0232013222110121-3223210311032231-0332221012122031-2132020312222132-2303023010020032-2101212313223322-1203003033031332-1111320102110212)
- [vn_config](data-sources--aws_tgw_site--reference--group-002.md#canonical-0302230100311030-1210212010000200-3030120211220013-1202212101231321-1322201012322230-0300232300110210-1103133222223202-1011002112321330)
- [vpc_attachments](data-sources--aws_tgw_site--reference--group-004.md#canonical-1001011112212113-0131301132102013-2201231332231131-3111000212301222-2111013110131231-3011310233312203-2332313100213022-0323211103120303)
- [waf_signatures](data-sources--aws_tgw_site--reference--group-004.md#canonical-0033210300111000-2300033032033230-0020231020220111-0223320021223331-0012031132001010-1110012102032212-1122313103211132-2001123200330212)
- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)

<a id="canonical-3311120101031201-0112330033331111-0330012133130331-2122223112223133-1300102323001230-2222231003110310-3112001230233232-2000103311201302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2232112320021303-1312020323331310-1322321101102310-2320312232013121-3203003123202010-3233310123312133-2111311132100311-1322320313200200"></a>

## aws_parameters — aws_parameters / 112032012001 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-1212321311102123-0132200012312320-2312231211033230-2200200312230233-3132320011023221-0111320212132020-3002000002020330-2223302203200011)
- aws_parameters

<a id="canonical-3332111212201100-1330102312211212-0010300231131013-2133303233030103-2222020321321001-2233111333023310-3211211211211201-1310230113020213"></a>

Type: `"single"`. Computed.

Setup AWS services VPC, transit gateway and site.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-deployment": "[\"aws_cred\"]",
  "x-ves-oneof-field-encryption_choice": "[\"disable_encryption\",\"enable_encryption\"]",
  "x-ves-oneof-field-internet_vip_choice": "[\"disable_internet_vip\",\"enable_internet_vip\"]",
  "x-ves-oneof-field-security_group_choice": "[\"custom_security_group\",\"f5xc_security_group\"]",
  "x-ves-oneof-field-service_vpc_choice": "[\"new_vpc\",\"vpc_id\"]",
  "x-ves-oneof-field-tgw_choice": "[\"existing_tgw\",\"new_tgw\"]",
  "x-ves-oneof-field-tgw_cidr_choice": "[\"reserved_tgw_cidr\",\"tgw_cidr\"]",
  "x-ves-oneof-field-worker_nodes": "[\"no_worker_nodes\",\"nodes_per_az\",\"total_nodes\"]"
}
```

<a id="canonical-0231312223033300-3022323331113233-3030330210023300-1120320112011102-0310021133300122-1330302011222021-1001133031023311-3200100021222021"></a>

## Direct properties — aws_parameters / 112032012001 / 3

- [admin_password](data-sources--aws_tgw_site--reference--group-001.md#canonical-1002133331211101-0233223330233302-3212000222100122-1010022102103121-2301122302222011-2303211211331121-1332000310212023-2210331000312001): complete subsection reference.

- [aws_cred](data-sources--aws_tgw_site--reference--group-001.md#canonical-2201133331133130-1312301001201200-3223131300300232-1213322111112011-2311320112113302-0203120123101333-1331011130203122-0213231113212031): complete subsection reference.

<a id="canonical-0321103100202213-1001213001220312-2033202220232322-0333232020120021-0210123102010022-3200303221223033-1123111231221031-1132102111132012"></a>

<a id="canonical-3113332333302231-3203110221232131-3031220233133200-3031320210112221-2303002111210303-1331032222112212-2330113013310212-3011332232130033"></a>

## aws_region property — aws_parameters / 112032012001 / 4

Type: `"string"`. Computed.

AWS Region of your services VPC, where F5XC site will be deployed.

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
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  }
}
```

- [az_nodes](data-sources--aws_tgw_site--reference--group-001.md#canonical-0300223333010011-0113000202023103-2021003320322111-1023131301113200-1102213013332011-1213031130233303-2311321321203130-2101213331332202): complete subsection reference.

- [custom_security_group](data-sources--aws_tgw_site--reference--group-001.md#canonical-1200330320023230-0322020333133231-1333202000222201-1300210113301131-3022222013031320-2323003101133112-0221332102112021-1220220320212322): complete subsection reference.

- [disable_encryption](data-sources--aws_tgw_site--reference--group-001.md#canonical-2332232011200302-1331212102222001-3302032333323332-3333033110331323-1123030122233131-1023032022330333-0330100101333210-2033132031213230): complete subsection reference.

- [disable_internet_vip](data-sources--aws_tgw_site--reference--group-001.md#canonical-1023220201100022-2321100222123310-3300233121120123-0322013200031311-1122223320323313-3130123020102012-1203103303102113-0202301101203333): complete subsection reference.

<a id="canonical-3012132010013313-2101102221002122-3013133013300233-0313032333103202-3100333001021010-2022013232003010-0301021303222300-3111322231033211"></a>

<a id="canonical-2130001131230311-2303301212221300-2003320330121132-1212131001000333-0322202112200010-3223113332303120-2212211223102201-3330323232203032"></a>

## disk_size property — aws_parameters / 112032012001 / 5

Type: `"number"`. Computed.

Node disk size for all node in the F5XC site. Unit is GiB.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 64000,
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
    "ves.io.schema.rules.uint32.lte": "64000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "64000"
  }
}
```

- [enable_encryption](data-sources--aws_tgw_site--reference--group-001.md#canonical-3010302002120100-2103002021320330-3110033102312003-2010320201333323-3230020212131301-2330231333302020-1201123130210012-0300222330022213): complete subsection reference.

- [enable_internet_vip](data-sources--aws_tgw_site--reference--group-001.md#canonical-0122200323013210-2301312003231211-3101311301201122-2120230202020100-0323212320030013-0332302101213300-2231312100322202-2213132110133310): complete subsection reference.

- [existing_tgw](data-sources--aws_tgw_site--reference--group-001.md#canonical-2022022330331132-1202301000122022-0131230132121103-2300033201303301-1103312012020121-1311322301301210-0321211103202233-0213003312233030): complete subsection reference.

- [f5xc_security_group](data-sources--aws_tgw_site--reference--group-001.md#canonical-2111203120120332-0311322131220100-1310033300211130-0302112330331312-0000101223200133-3210013202322002-0101000113321133-3320021300310020): complete subsection reference.

<a id="canonical-1030012310020131-2311011230112033-3200323020102113-1323330023111023-0133122010203323-3201031131311223-1310313213202223-0333221110213110"></a>

<a id="canonical-1131330123322210-0201132131030312-2133103332303113-3111120312302021-0113003322111312-3232110133332223-3011233301001100-1011302300130000"></a>

## instance_type property — aws_parameters / 112032012001 / 6

Type: `"string"`. Computed.

AWS Instance Type for Node. Instance size based on the performance.

Upstream description:

Instance size based on the performance.

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

- [new_tgw](data-sources--aws_tgw_site--reference--group-001.md#canonical-1232303130313003-0000010222131313-2230032301201222-3310233212211201-1303130113131201-2022121312300000-0120130132002311-3330030310301103): complete subsection reference.

- [new_vpc](data-sources--aws_tgw_site--reference--group-001.md#canonical-2203311030300201-1320113131100101-3012322103333133-3003203313213332-2310302023313002-3013210130001312-2033332213210030-0030330322300233): complete subsection reference.

- [no_worker_nodes](data-sources--aws_tgw_site--reference--group-002.md#canonical-0231100231122132-2021112330123120-0002300113022030-2301010222102233-1213302232110033-3022113202020003-0322300221102323-2233201022000200): complete subsection reference.

<a id="canonical-1031003233122100-0002211012301200-3203133333033310-0001103123113211-1020010032032311-0232021022010020-1331113213102220-3320223313332310"></a>

<a id="canonical-1100230322100232-2012011030230323-0201013310011120-0330302333302131-3012302101133220-1100212321111213-2201201232223302-0213313332320032"></a>

## nodes_per_az property — aws_parameters / 112032012001 / 7

Type: `"number"`. Computed.

Exclusive with \[no\_worker\_nodes total\_nodes\] Desired Worker Nodes Per AZ. Max limit is up to
21.

Upstream description:

Exclusive with \[no\_worker\_nodes total\_nodes\] Desired Worker Nodes Per AZ. Max limit is up to
21.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 21,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 0
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "21"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "21"
  }
}
```

- [reserved_tgw_cidr](data-sources--aws_tgw_site--reference--group-002.md#canonical-3033320011130322-2011303033033001-0310202021301030-3231002233002112-2303103000310110-0332011032300123-0320013032020212-2331110113213020): complete subsection reference.

<a id="canonical-2212322331222003-2032213220002321-2331130032122222-0120121102222012-1221023112213101-3330010313003132-3100031321230012-2233212010310101"></a>

<a id="canonical-1210013021330122-1223000033212333-1230211130122121-0010331302321012-2333021333320121-2200032101101002-2111123210322201-0222001212333202"></a>

## ssh_key property — aws_parameters / 112032012001 / 8

Type: `"string"`. Computed.

Public SSH key for accessing nodes of the site.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 8192,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 8192,
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
    "ves.io.schema.rules.string.max_len": "8192",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "8192",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

- [tgw_cidr](data-sources--aws_tgw_site--reference--group-002.md#canonical-3202021121111032-1032113112203313-1011031301321120-1031130323101023-1123113100211230-0030303011131111-1220032320323021-3021110302323002): complete subsection reference.

<a id="canonical-0003221222331201-2201033333213321-2230011311023003-2332000103221112-2130103333302203-0100123030003301-1230311300131102-2033112322133100"></a>

<a id="canonical-2311010202111133-0002302023223113-0030300002330313-2223103223103331-3022320312031223-3221002011120003-2100321013211023-2222000222120133"></a>

## total_nodes property — aws_parameters / 112032012001 / 9

Type: `"number"`. Computed.

Exclusive with \[no\_worker\_nodes nodes\_per\_az\] Total number of worker nodes to be deployed
across all AZ's used in the Site.

Upstream description:

Exclusive with \[no\_worker\_nodes nodes\_per\_az\] Total number of worker nodes to be deployed
across all AZ's used in the Site.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 61,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 0
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "61"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "61"
  }
}
```

<a id="canonical-2110330210333202-2203213120033331-1021020121111212-3122212121110200-3131113020300102-1120300002301101-3200130233033013-2200020322032010"></a>

<a id="canonical-0332000311332233-3313222010332110-2112012030101020-0221222002233223-2310011303320213-0030212311313131-1333323002210121-1231233312000323"></a>

## vpc_id property — aws_parameters / 112032012001 / 10

Type: `"string"`. Computed.

Exclusive with \[new\_vpc\] Existing VPC ID.

Upstream description:

Exclusive with \[new\_vpc\] Existing VPC ID.

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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.pattern": "^(vpc-)([a-z0-9]{8}|[a-z0-9]{17})$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.pattern": "^(vpc-)([a-z0-9]{8}|[a-z0-9]{17})$"
  }
}
```

<a id="canonical-0031010100012232-1031221121310220-3203013122231001-0322101011303320-0111320300102322-2301233232201311-1203322230100323-0101200101321202"></a>

## Next pages — aws_parameters / 112032012001 / 11

- [aws_parameters.admin_password](data-sources--aws_tgw_site--reference--group-001.md#canonical-1002133331211101-0233223330233302-3212000222100122-1010022102103121-2301122302222011-2303211211331121-1332000310212023-2210331000312001)
- [aws_parameters.aws_cred](data-sources--aws_tgw_site--reference--group-001.md#canonical-2201133331133130-1312301001201200-3223131300300232-1213322111112011-2311320112113302-0203120123101333-1331011130203122-0213231113212031)
- [aws_parameters.az_nodes](data-sources--aws_tgw_site--reference--group-001.md#canonical-0300223333010011-0113000202023103-2021003320322111-1023131301113200-1102213013332011-1213031130233303-2311321321203130-2101213331332202)
- [aws_parameters.custom_security_group](data-sources--aws_tgw_site--reference--group-001.md#canonical-1200330320023230-0322020333133231-1333202000222201-1300210113301131-3022222013031320-2323003101133112-0221332102112021-1220220320212322)
- [aws_parameters.disable_encryption](data-sources--aws_tgw_site--reference--group-001.md#canonical-2332232011200302-1331212102222001-3302032333323332-3333033110331323-1123030122233131-1023032022330333-0330100101333210-2033132031213230)
- [aws_parameters.disable_internet_vip](data-sources--aws_tgw_site--reference--group-001.md#canonical-1023220201100022-2321100222123310-3300233121120123-0322013200031311-1122223320323313-3130123020102012-1203103303102113-0202301101203333)
- [aws_parameters.enable_encryption](data-sources--aws_tgw_site--reference--group-001.md#canonical-3010302002120100-2103002021320330-3110033102312003-2010320201333323-3230020212131301-2330231333302020-1201123130210012-0300222330022213)
- [aws_parameters.enable_internet_vip](data-sources--aws_tgw_site--reference--group-001.md#canonical-0122200323013210-2301312003231211-3101311301201122-2120230202020100-0323212320030013-0332302101213300-2231312100322202-2213132110133310)
- [aws_parameters.existing_tgw](data-sources--aws_tgw_site--reference--group-001.md#canonical-2022022330331132-1202301000122022-0131230132121103-2300033201303301-1103312012020121-1311322301301210-0321211103202233-0213003312233030)
- [aws_parameters.f5xc_security_group](data-sources--aws_tgw_site--reference--group-001.md#canonical-2111203120120332-0311322131220100-1310033300211130-0302112330331312-0000101223200133-3210013202322002-0101000113321133-3320021300310020)
- [aws_parameters.new_tgw](data-sources--aws_tgw_site--reference--group-001.md#canonical-1232303130313003-0000010222131313-2230032301201222-3310233212211201-1303130113131201-2022121312300000-0120130132002311-3330030310301103)
- [aws_parameters.new_vpc](data-sources--aws_tgw_site--reference--group-001.md#canonical-2203311030300201-1320113131100101-3012322103333133-3003203313213332-2310302023313002-3013210130001312-2033332213210030-0030330322300233)
- [aws_parameters.no_worker_nodes](data-sources--aws_tgw_site--reference--group-002.md#canonical-0231100231122132-2021112330123120-0002300113022030-2301010222102233-1213302232110033-3022113202020003-0322300221102323-2233201022000200)
- [aws_parameters.reserved_tgw_cidr](data-sources--aws_tgw_site--reference--group-002.md#canonical-3033320011130322-2011303033033001-0310202021301030-3231002233002112-2303103000310110-0332011032300123-0320013032020212-2331110113213020)
- [aws_parameters.tgw_cidr](data-sources--aws_tgw_site--reference--group-002.md#canonical-3202021121111032-1032113112203313-1011031301321120-1031130323101023-1123113100211230-0030303011131111-1220032320323021-3021110302323002)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-1212321311102123-0132200012312320-2312231211033230-2200200312230233-3132320011023221-0111320212132020-3002000002020330-2223302203200011)
- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)

<a id="canonical-1002133331211101-0233223330233302-3212000222100122-1010022102103121-2301122302222011-2303211211331121-1332000310212023-2210331000312001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0010130203030213-3120103001330311-2032132131133122-2320232001003032-1221000323203222-2022100131031023-1330112030332231-1001301231232320"></a>

## aws_parameters.admin_password — admin_password / 213333122232 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-1212321311102123-0132200012312320-2312231211033230-2200200312230233-3132320011023221-0111320212132020-3002000002020330-2223302203200011)
- [aws_parameters](data-sources--aws_tgw_site--reference--group-001.md#canonical-3311120101031201-0112330033331111-0330012133130331-2122223112223133-1300102323001230-2222231003110310-3112001230233232-2000103311201302)
- aws_parameters.admin_password

<a id="canonical-1330231030122222-0001220132323320-2032200212031122-3310311110301121-2000111103003023-0213013212032102-1223231011220021-2100211101210010"></a>

Type: `"single"`. Computed.

SecretType is used in an object to indicate a sensitive/confidential field.

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

<a id="canonical-2101303031202131-1220132200333212-0313322100333112-2132130230201022-0101020212230313-0313031302332023-0230112231002302-3030021133131323"></a>

## Direct properties — admin_password / 213333122232 / 3

- [blindfold_secret_info](data-sources--aws_tgw_site--reference--group-001.md#canonical-1010123322011111-3133110332311101-0330133021020322-0301301300321213-2103123012111011-3312101222310223-2031110323003100-1300210102111133): complete subsection reference.

- [clear_secret_info](data-sources--aws_tgw_site--reference--group-001.md#canonical-3131102232320013-1033303200032300-0233132213002023-2200310123202313-0223202013123202-1301310331232132-0131002323003333-1330210003311121): complete subsection reference.

<a id="canonical-2233312100312012-0110321212103002-1221013130023322-2302132032000122-3001101030313213-0012031332222311-0321320323333302-3233323211131313"></a>

## Next pages — admin_password / 213333122232 / 4

- [aws_parameters.admin_password.blindfold_secret_info](data-sources--aws_tgw_site--reference--group-001.md#canonical-1010123322011111-3133110332311101-0330133021020322-0301301300321213-2103123012111011-3312101222310223-2031110323003100-1300210102111133)
- [aws_parameters.admin_password.clear_secret_info](data-sources--aws_tgw_site--reference--group-001.md#canonical-3131102232320013-1033303200032300-0233132213002023-2200310123202313-0223202013123202-1301310331232132-0131002323003333-1330210003311121)
- [aws_parameters](data-sources--aws_tgw_site--reference--group-001.md#canonical-3311120101031201-0112330033331111-0330012133130331-2122223112223133-1300102323001230-2222231003110310-3112001230233232-2000103311201302)
- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)

<a id="canonical-1010123322011111-3133110332311101-0330133021020322-0301301300321213-2103123012111011-3312101222310223-2031110323003100-1300210102111133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2303011210320212-1023112031333212-0300331022022301-0323231123020030-2203303101313320-2303230010232133-2232001111112132-2030231112323323"></a>

## aws_parameters.admin_password.blindfold_secret_info — blindfold_secret_info / 123210033212 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-1212321311102123-0132200012312320-2312231211033230-2200200312230233-3132320011023221-0111320212132020-3002000002020330-2223302203200011)
- [aws_parameters](data-sources--aws_tgw_site--reference--group-001.md#canonical-3311120101031201-0112330033331111-0330012133130331-2122223112223133-1300102323001230-2222231003110310-3112001230233232-2000103311201302)
- [aws_parameters.admin_password](data-sources--aws_tgw_site--reference--group-001.md#canonical-1002133331211101-0233223330233302-3212000222100122-1010022102103121-2301122302222011-2303211211331121-1332000310212023-2210331000312001)
- aws_parameters.admin_password.blindfold_secret_info

<a id="canonical-2032220023132303-3221330220321201-2102111303303123-2100303000100211-3030310103033121-3011333210120222-1021012001030002-3230121213031010"></a>

Type: `"single"`. Computed.

BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0131030313233022-2110213332202011-1200311311231322-1130000220120013-3020333230221121-1203233030321320-2031101111103121-0132200311330333"></a>

## Direct properties — blindfold_secret_info / 123210033212 / 3

<a id="canonical-1123330331313200-2313223100133123-0203331033320310-2131232021002122-0330333112113230-1212130023022013-2033113111311012-0322301013331012"></a>

<a id="canonical-3021023223021120-2330020311322201-3331202311110320-0311013133233302-3123213032010121-1213212320211200-2133212011311321-1333301230020000"></a>

## decryption_provider property — blindfold_secret_info / 123210033212 / 4

Type: `"string"`. Computed.

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

<a id="canonical-1313300330220211-0011203213002320-1110001001110120-2230102122231033-1133301300013013-3130221303130231-0112313023112010-3203220003001123"></a>

<a id="canonical-0120302200213130-1001211022202001-0103230230103203-3003003321113200-3221121301302132-3302332120323303-1031212301213332-3031101200032303"></a>

## location property — blindfold_secret_info / 123210033212 / 5

Type: `"string"`. Computed, Sensitive.

Location is the uri\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Upstream description:

Location is the uri\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

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

<a id="canonical-3030020302220130-1313121013231203-0311320321032000-3221333001333202-0030103232031313-2210333022330233-2130321300213323-0332112210311103"></a>

<a id="canonical-2101312212013133-3123313210023212-2222212112331121-2022131121322112-1233231222111200-0032113220010203-3222230302312213-2203111113323331"></a>

## store_provider property — blindfold_secret_info / 123210033212 / 6

Type: `"string"`. Computed.

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

<a id="canonical-1110132133133231-0111021132101211-3223021101221202-2011212100000023-0321330131000022-3313023122010120-1332211331320210-1312013330230021"></a>

## Next pages — blindfold_secret_info / 123210033212 / 7

- [aws_parameters.admin_password](data-sources--aws_tgw_site--reference--group-001.md#canonical-1002133331211101-0233223330233302-3212000222100122-1010022102103121-2301122302222011-2303211211331121-1332000310212023-2210331000312001)
- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)

<a id="canonical-3131102232320013-1033303200032300-0233132213002023-2200310123202313-0223202013123202-1301310331232132-0131002323003333-1330210003311121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3032230020110000-1131013100010233-3010011220103203-3200031323100303-2220101011300022-3100210012321122-0102323033003003-0302130201121130"></a>

## aws_parameters.admin_password.clear_secret_info — clear_secret_info / 123211200223 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-1212321311102123-0132200012312320-2312231211033230-2200200312230233-3132320011023221-0111320212132020-3002000002020330-2223302203200011)
- [aws_parameters](data-sources--aws_tgw_site--reference--group-001.md#canonical-3311120101031201-0112330033331111-0330012133130331-2122223112223133-1300102323001230-2222231003110310-3112001230233232-2000103311201302)
- [aws_parameters.admin_password](data-sources--aws_tgw_site--reference--group-001.md#canonical-1002133331211101-0233223330233302-3212000222100122-1010022102103121-2301122302222011-2303211211331121-1332000310212023-2210331000312001)
- aws_parameters.admin_password.clear_secret_info

<a id="canonical-0102331322020110-1031312000021102-2033221020030011-3320020212012232-3230132122211220-0002001231000113-1110112011311303-0031121322112031"></a>

Type: `"single"`. Computed.

ClearSecretInfoType specifies information about the Secret that is not encrypted.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0111212330300032-1323303003011112-0330131033232312-1023002302233022-2001332313321301-0132001133010203-0301033112120301-3023003330232300"></a>

## Direct properties — clear_secret_info / 123211200223 / 3

<a id="canonical-1121002121123001-3321321320231110-0221300210020221-2122132022222331-1103322200312101-1110002222130113-3302232133211312-3032332233333122"></a>

<a id="canonical-1123232132200230-0333131210123103-0011133100232212-3011132022320231-0121131220233012-2303300203001130-1233212110330223-2010302021133233"></a>

## provider_ref property — clear_secret_info / 123211200223 / 4

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-3312022033313320-0322200202213210-3110013010013203-1120303222233011-0231333313220111-0120330230212203-2223212113300020-2222033030003010"></a>

<a id="canonical-2201313002011232-0311230223220220-2132023201310211-2002121113103312-0202013200302231-2212312221020321-2202313310020322-3023223020022222"></a>

## URL property — clear_secret_info / 123211200223 / 5

Type: `"string"`. Computed, Sensitive.

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded Base64 format. When asked for this secret, caller will GET Secret bytes after
Base64 decoding.

Upstream description:

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded Base64 format. When asked for this secret, caller will GET Secret bytes after
Base64 decoding.

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

<a id="canonical-0300121220211333-0100030231211323-1213130003130120-0121112220211003-3323122133112311-3002032231312202-0231221130300033-3200233010002332"></a>

## Next pages — clear_secret_info / 123211200223 / 6

- [aws_parameters.admin_password](data-sources--aws_tgw_site--reference--group-001.md#canonical-1002133331211101-0233223330233302-3212000222100122-1010022102103121-2301122302222011-2303211211331121-1332000310212023-2210331000312001)
- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)

<a id="canonical-2201133331133130-1312301001201200-3223131300300232-1213322111112011-2311320112113302-0203120123101333-1331011130203122-0213231113212031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2201302023323303-3230331001213011-3322233121223122-1232031120132302-1201001110323032-3103311132033033-0112220113100112-1120100012300221"></a>

## aws_parameters.aws_cred — aws_cred / 210320310223 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-1212321311102123-0132200012312320-2312231211033230-2200200312230233-3132320011023221-0111320212132020-3002000002020330-2223302203200011)
- [aws_parameters](data-sources--aws_tgw_site--reference--group-001.md#canonical-3311120101031201-0112330033331111-0330012133130331-2122223112223133-1300102323001230-2222231003110310-3112001230233232-2000103311201302)
- aws_parameters.aws_cred

<a id="canonical-3220133102232332-1123332230300110-0221131132232103-2220031310013221-3331211321303333-1131113100032100-1111121021302330-2003321100301320"></a>

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

<a id="canonical-3031020230020031-0213322211200231-1011010121301222-0322021323233232-0101122331202101-0002111030332032-1212003021311131-0002133231113310"></a>

## Direct properties — aws_cred / 210320310223 / 3

<a id="canonical-0110323103202032-1111121221201001-0203110000322321-0301110211001002-1132031333103300-1132232021031301-1023323202021022-0103132300301331"></a>

<a id="canonical-3120301021131001-0000323133230331-1222021003231012-1100311033302020-1133020102030132-1012233030003011-2032322331113231-2032122220032322"></a>

## name property — aws_cred / 210320310223 / 4

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

<a id="canonical-2200333020211210-1002023123013310-3333322312012001-1210012103001011-3312111112130211-2221230101232002-1333110313030323-2310231312230322"></a>

<a id="canonical-3323100200221202-0002110130000011-2131103211112132-3002000212330031-2100020321033313-2102313031021030-1331223000112103-2002100301001203"></a>

## namespace property — aws_cred / 210320310223 / 5

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

<a id="canonical-3302132302030133-1320322210011121-1021303110001111-1103231012013111-1230103133120100-0320113011230210-1320120001020131-0012022033330220"></a>

<a id="canonical-0133130222010312-3122322313012132-2031332212221303-0122230103310320-0331012332100332-1022121100222012-1222010310021123-0230203013132323"></a>

## tenant property — aws_cred / 210320310223 / 6

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

<a id="canonical-3320211311233131-0101010030222131-3330011323200310-2303121020031233-2322000131213011-1200313010121130-3221000301103200-2322223020021332"></a>

## Next pages — aws_cred / 210320310223 / 7

- [aws_parameters](data-sources--aws_tgw_site--reference--group-001.md#canonical-3311120101031201-0112330033331111-0330012133130331-2122223112223133-1300102323001230-2222231003110310-3112001230233232-2000103311201302)
- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)

<a id="canonical-0300223333010011-0113000202023103-2021003320322111-1023131301113200-1102213013332011-1213031130233303-2311321321203130-2101213331332202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1002323200013230-1332310121323030-3120301000121031-3301120300320322-1033010132133100-2332112120132311-0133200113122020-2021313010121010"></a>

## aws_parameters.az_nodes — az_nodes / 010103300213 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-1212321311102123-0132200012312320-2312231211033230-2200200312230233-3132320011023221-0111320212132020-3002000002020330-2223302203200011)
- [aws_parameters](data-sources--aws_tgw_site--reference--group-001.md#canonical-3311120101031201-0112330033331111-0330012133130331-2122223112223133-1300102323001230-2222231003110310-3112001230233232-2000103311201302)
- aws_parameters.az_nodes

<a id="canonical-1101111332223202-3000223312031132-0111121313002110-2103022120313333-3232331200202213-2223103330203013-3032131333321031-2110311302133101"></a>

Type: `"list"`. Computed.

Only Single AZ or Three AZ(s) nodes are supported currently.

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
    "ves.io.schema.rules.repeated.num_items": "1,3"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.num_items": "1,3"
  }
}
```

<a id="canonical-3230300300333013-3223100213022021-2303001210021201-2233303211103220-2110123131032120-1332211221231103-0211003030202111-1121203102133320"></a>

## Direct properties — az_nodes / 010103300213 / 3

<a id="canonical-2200210322300230-0121200110020122-0311233132011033-0330033302022113-2231302200203323-3000310303322221-3121320330322002-3323000101321001"></a>

<a id="canonical-0331201101213103-0223011220022213-3101132003320301-1220302323130122-1332330110322212-1222313220201102-3300000221213233-1330103221023223"></a>

## aws_az_name property — az_nodes / 010103300213 / 4

Type: `"string"`. Computed.

AWS availability zone, must be consistent with the selected AWS region.

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
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  }
}
```

- [inside_subnet](data-sources--aws_tgw_site--reference--group-001.md#canonical-1032211122202210-1202000032012123-2001320322331313-3230111133303222-2113110333110212-0113031331331211-1113233121202133-2000021132210131): complete subsection reference.

- [outside_subnet](data-sources--aws_tgw_site--reference--group-001.md#canonical-2300323120032113-1330312103113310-0130330012012230-0302321002222031-2202010231030133-3210110230001301-2010001203220310-0123221213330122): complete subsection reference.

- [reserved_inside_subnet](data-sources--aws_tgw_site--reference--group-001.md#canonical-1123033132220002-3032212132320031-2121230131101203-2102220211022103-0220312011230111-3111100232123222-3130311202310230-0021311232300120): complete subsection reference.

- [workload_subnet](data-sources--aws_tgw_site--reference--group-001.md#canonical-2001023333111231-1300333312131200-2123102003330300-1321222313100103-0300131320232133-0333313233211023-0323320201232213-3213103210002203): complete subsection reference.

<a id="canonical-3322322310010302-1000202133011320-3013133303201112-2210213302022212-2313012001222221-0122320111211123-0211322313303221-0323322003213202"></a>

## Next pages — az_nodes / 010103300213 / 5

- [aws_parameters.az_nodes.inside_subnet](data-sources--aws_tgw_site--reference--group-001.md#canonical-1032211122202210-1202000032012123-2001320322331313-3230111133303222-2113110333110212-0113031331331211-1113233121202133-2000021132210131)
- [aws_parameters.az_nodes.outside_subnet](data-sources--aws_tgw_site--reference--group-001.md#canonical-2300323120032113-1330312103113310-0130330012012230-0302321002222031-2202010231030133-3210110230001301-2010001203220310-0123221213330122)
- [aws_parameters.az_nodes.reserved_inside_subnet](data-sources--aws_tgw_site--reference--group-001.md#canonical-1123033132220002-3032212132320031-2121230131101203-2102220211022103-0220312011230111-3111100232123222-3130311202310230-0021311232300120)
- [aws_parameters.az_nodes.workload_subnet](data-sources--aws_tgw_site--reference--group-001.md#canonical-2001023333111231-1300333312131200-2123102003330300-1321222313100103-0300131320232133-0333313233211023-0323320201232213-3213103210002203)
- [aws_parameters](data-sources--aws_tgw_site--reference--group-001.md#canonical-3311120101031201-0112330033331111-0330012133130331-2122223112223133-1300102323001230-2222231003110310-3112001230233232-2000103311201302)
- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)

<a id="canonical-1032211122202210-1202000032012123-2001320322331313-3230111133303222-2113110333110212-0113031331331211-1113233121202133-2000021132210131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2102332121001122-0301201012221232-2300111021221023-3301103001331202-3212130230212131-2112323213312222-1001223012001311-3123121320210112"></a>

## aws_parameters.az_nodes.inside_subnet — inside_subnet / 010132111200 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-1212321311102123-0132200012312320-2312231211033230-2200200312230233-3132320011023221-0111320212132020-3002000002020330-2223302203200011)
- [aws_parameters](data-sources--aws_tgw_site--reference--group-001.md#canonical-3311120101031201-0112330033331111-0330012133130331-2122223112223133-1300102323001230-2222231003110310-3112001230233232-2000103311201302)
- [aws_parameters.az_nodes](data-sources--aws_tgw_site--reference--group-001.md#canonical-0300223333010011-0113000202023103-2021003320322111-1023131301113200-1102213013332011-1213031130233303-2311321321203130-2101213331332202)
- aws_parameters.az_nodes.inside_subnet

<a id="canonical-0312131210003312-1120101211221121-3100221230332100-1021030010002221-0002220223111313-0213201120021221-1103111011021210-0133202302203003"></a>

Type: `"single"`. Computed.

Configuration parameter for inside subnet.

Upstream description:

Parameters for AWS subnet.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-choice": "[\"existing_subnet_id\",\"subnet_param\"]"
}
```

<a id="canonical-1111132030131122-1023003120111120-1300213031332311-1322223201310101-3131333203320000-0121133330313333-1203231020122233-2212002100331013"></a>

## Direct properties — inside_subnet / 010132111200 / 3

<a id="canonical-1131333233111123-1201333100030321-1333310020102011-3303333030221223-0011022202100331-0300213201100231-2123131230002233-2021210030210320"></a>

<a id="canonical-2210111022333233-1020320301100033-0221001000330110-3302220030020110-3221021011000002-2122200131021100-1303300020303303-3002232133010131"></a>

## existing_subnet_id property — inside_subnet / 010132111200 / 4

Type: `"string"`. Computed.

Exclusive with \[subnet\_param\] Information about existing subnet ID.

Upstream description:

Exclusive with \[subnet\_param\] Information about existing subnet ID.

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
    "pattern": "^(subnet-)([a-z0-9]{8}|[a-z0-9]{17})$"
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.pattern": "^(subnet-)([a-z0-9]{8}|[a-z0-9]{17})$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.pattern": "^(subnet-)([a-z0-9]{8}|[a-z0-9]{17})$"
  }
}
```

- [subnet_param](data-sources--aws_tgw_site--reference--group-001.md#canonical-3300021312023210-2030102111000020-3011130030231300-1130220233022223-1103303211110002-2120322230330122-0131030132112232-3110101000230312): complete subsection reference.

<a id="canonical-3320332021120030-2302213111102130-0213312113120302-1133003121303103-2223203032133310-3130220020311322-3120033021220112-3012002132001213"></a>

## Next pages — inside_subnet / 010132111200 / 5

- [aws_parameters.az_nodes.inside_subnet.subnet_param](data-sources--aws_tgw_site--reference--group-001.md#canonical-3300021312023210-2030102111000020-3011130030231300-1130220233022223-1103303211110002-2120322230330122-0131030132112232-3110101000230312)
- [aws_parameters.az_nodes](data-sources--aws_tgw_site--reference--group-001.md#canonical-0300223333010011-0113000202023103-2021003320322111-1023131301113200-1102213013332011-1213031130233303-2311321321203130-2101213331332202)
- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)

<a id="canonical-3300021312023210-2030102111000020-3011130030231300-1130220233022223-1103303211110002-2120322230330122-0131030132112232-3110101000230312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0332323320133020-2323221201311100-2011323312330331-1112113303200020-1103032233032000-2301123213331022-1231011011112113-2131032131020302"></a>

## aws_parameters.az_nodes.inside_subnet.subnet_param — subnet_param / 200010111321 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-1212321311102123-0132200012312320-2312231211033230-2200200312230233-3132320011023221-0111320212132020-3002000002020330-2223302203200011)
- [aws_parameters](data-sources--aws_tgw_site--reference--group-001.md#canonical-3311120101031201-0112330033331111-0330012133130331-2122223112223133-1300102323001230-2222231003110310-3112001230233232-2000103311201302)
- [aws_parameters.az_nodes](data-sources--aws_tgw_site--reference--group-001.md#canonical-0300223333010011-0113000202023103-2021003320322111-1023131301113200-1102213013332011-1213031130233303-2311321321203130-2101213331332202)
- [aws_parameters.az_nodes.inside_subnet](data-sources--aws_tgw_site--reference--group-001.md#canonical-1032211122202210-1202000032012123-2001320322331313-3230111133303222-2113110333110212-0113031331331211-1113233121202133-2000021132210131)
- aws_parameters.az_nodes.inside_subnet.subnet_param

<a id="canonical-3322013001112011-3021300120321031-0133012301321133-0312300211013221-0303222022011202-2331121320033120-2011012023311100-0303211330100210"></a>

Type: `"single"`. Computed.

Parameters for creating a new cloud subnet.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2013120232213003-0213111010021022-0130130222201110-3221213100221113-0100210001101102-3030132121132331-1330131012232111-2031300302321113"></a>

## Direct properties — subnet_param / 200010111321 / 3

<a id="canonical-2333231132130323-2303121233101133-2002302302003232-3003023123002030-1112001003021220-2012310311033310-2201320321132312-0023123110123012"></a>

<a id="canonical-1000222230222011-3113230301103031-2222202332302102-2320310011223200-3231233122022022-0312110010200102-1113202012100033-1023001133013122"></a>

## IPv4 property — subnet_param / 200010111321 / 4

Type: `"string"`. Computed.

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

<a id="canonical-3121203100332102-1200133320001012-3020122112203003-3023103121022112-0100200331310321-3231301003132002-3321312223002312-1311030302032012"></a>

## Next pages — subnet_param / 200010111321 / 5

- [aws_parameters.az_nodes.inside_subnet](data-sources--aws_tgw_site--reference--group-001.md#canonical-1032211122202210-1202000032012123-2001320322331313-3230111133303222-2113110333110212-0113031331331211-1113233121202133-2000021132210131)
- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)

<a id="canonical-2300323120032113-1330312103113310-0130330012012230-0302321002222031-2202010231030133-3210110230001301-2010001203220310-0123221213330122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1122321201320203-3000230231313021-3323101203123003-1130121021211213-0022331331211232-3332230010122300-3300020023231011-1130123230202213"></a>

## aws_parameters.az_nodes.outside_subnet — outside_subnet / 112123300001 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-1212321311102123-0132200012312320-2312231211033230-2200200312230233-3132320011023221-0111320212132020-3002000002020330-2223302203200011)
- [aws_parameters](data-sources--aws_tgw_site--reference--group-001.md#canonical-3311120101031201-0112330033331111-0330012133130331-2122223112223133-1300102323001230-2222231003110310-3112001230233232-2000103311201302)
- [aws_parameters.az_nodes](data-sources--aws_tgw_site--reference--group-001.md#canonical-0300223333010011-0113000202023103-2021003320322111-1023131301113200-1102213013332011-1213031130233303-2311321321203130-2101213331332202)
- aws_parameters.az_nodes.outside_subnet

<a id="canonical-3022300030033130-1121203222003300-2230102101322323-3121002213122013-1023001113110310-3311322031213202-0001230012332211-1002112013300033"></a>

Type: `"single"`. Computed.

Configuration parameter for outside subnet.

Upstream description:

Parameters for AWS subnet.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-choice": "[\"existing_subnet_id\",\"subnet_param\"]"
}
```

<a id="canonical-1000210233000333-2133200213120302-0022021023111222-2033131220121301-0130030320312000-3330030233313132-1302132233102030-1001330232012213"></a>

## Direct properties — outside_subnet / 112123300001 / 3

<a id="canonical-3211223201312103-3313100320121120-0122323331012030-0022130210121010-3133121101100212-3002302011202132-2100002121222231-2013303331010030"></a>

<a id="canonical-0200122103221302-0222123201112313-1303322332301301-0130321322012201-2200300331103301-2211111103020002-3301111200100301-1031320320012331"></a>

## existing_subnet_id property — outside_subnet / 112123300001 / 4

Type: `"string"`. Computed.

Exclusive with \[subnet\_param\] Information about existing subnet ID.

Upstream description:

Exclusive with \[subnet\_param\] Information about existing subnet ID.

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
    "pattern": "^(subnet-)([a-z0-9]{8}|[a-z0-9]{17})$"
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.pattern": "^(subnet-)([a-z0-9]{8}|[a-z0-9]{17})$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.pattern": "^(subnet-)([a-z0-9]{8}|[a-z0-9]{17})$"
  }
}
```

- [subnet_param](data-sources--aws_tgw_site--reference--group-001.md#canonical-1222233101021211-2100032213203333-3312200232010002-3111213220123023-3002020211133022-3121202102302130-0020301032233011-1303133121312313): complete subsection reference.

<a id="canonical-3010331321131211-1211002101321223-0120100022223132-2022002213130023-3311202101303302-1033122000210301-1313303102002012-1133030011133302"></a>

## Next pages — outside_subnet / 112123300001 / 5

- [aws_parameters.az_nodes.outside_subnet.subnet_param](data-sources--aws_tgw_site--reference--group-001.md#canonical-1222233101021211-2100032213203333-3312200232010002-3111213220123023-3002020211133022-3121202102302130-0020301032233011-1303133121312313)
- [aws_parameters.az_nodes](data-sources--aws_tgw_site--reference--group-001.md#canonical-0300223333010011-0113000202023103-2021003320322111-1023131301113200-1102213013332011-1213031130233303-2311321321203130-2101213331332202)
- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)

<a id="canonical-1222233101021211-2100032213203333-3312200232010002-3111213220123023-3002020211133022-3121202102302130-0020301032233011-1303133121312313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1003002100211011-3030323323123131-1021032233111003-0302032020101100-2110112323231203-1121130103122302-2330232331300101-2211222003122202"></a>

## aws_parameters.az_nodes.outside_subnet.subnet_param — subnet_param / 022221232303 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-1212321311102123-0132200012312320-2312231211033230-2200200312230233-3132320011023221-0111320212132020-3002000002020330-2223302203200011)
- [aws_parameters](data-sources--aws_tgw_site--reference--group-001.md#canonical-3311120101031201-0112330033331111-0330012133130331-2122223112223133-1300102323001230-2222231003110310-3112001230233232-2000103311201302)
- [aws_parameters.az_nodes](data-sources--aws_tgw_site--reference--group-001.md#canonical-0300223333010011-0113000202023103-2021003320322111-1023131301113200-1102213013332011-1213031130233303-2311321321203130-2101213331332202)
- [aws_parameters.az_nodes.outside_subnet](data-sources--aws_tgw_site--reference--group-001.md#canonical-2300323120032113-1330312103113310-0130330012012230-0302321002222031-2202010231030133-3210110230001301-2010001203220310-0123221213330122)
- aws_parameters.az_nodes.outside_subnet.subnet_param

<a id="canonical-1130322131311001-1002113223313233-0220201012032321-3303102132113202-0033303312220321-1102110201111201-1030121330333030-2011102231001130"></a>

Type: `"single"`. Computed.

Parameters for creating a new cloud subnet.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1231100132101003-0031323212013212-0132032102233030-2220220002333230-2120200221030321-2112131222221230-1330103111113231-3013300111301120"></a>

## Direct properties — subnet_param / 022221232303 / 3

<a id="canonical-1002121220112201-2103130310301311-1201201301102123-2013100132003333-1200021232022102-3222232312212131-2320231223221223-0120321212132113"></a>

<a id="canonical-2132330010323031-1111130001011322-1300022310113032-1011323003130313-1313023000310021-3000223220320232-1332211121301112-3332132211012321"></a>

## IPv4 property — subnet_param / 022221232303 / 4

Type: `"string"`. Computed.

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

<a id="canonical-1032233222201003-3302303231130113-0101113231211200-2210203030332023-2130001310020000-2120233302121121-3222223232222032-2000221333030122"></a>

## Next pages — subnet_param / 022221232303 / 5

- [aws_parameters.az_nodes.outside_subnet](data-sources--aws_tgw_site--reference--group-001.md#canonical-2300323120032113-1330312103113310-0130330012012230-0302321002222031-2202010231030133-3210110230001301-2010001203220310-0123221213330122)
- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)

<a id="canonical-1123033132220002-3032212132320031-2121230131101203-2102220211022103-0220312011230111-3111100232123222-3130311202310230-0021311232300120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1312332210222213-1330312203013313-0110231003123203-0110220113330230-0010321033231210-3112200310122031-0011231313310030-2123332323112111"></a>

## aws_parameters.az_nodes.reserved_inside_subnet — reserved_inside_subnet / 121110311013 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-1212321311102123-0132200012312320-2312231211033230-2200200312230233-3132320011023221-0111320212132020-3002000002020330-2223302203200011)
- [aws_parameters](data-sources--aws_tgw_site--reference--group-001.md#canonical-3311120101031201-0112330033331111-0330012133130331-2122223112223133-1300102323001230-2222231003110310-3112001230233232-2000103311201302)
- [aws_parameters.az_nodes](data-sources--aws_tgw_site--reference--group-001.md#canonical-0300223333010011-0113000202023103-2021003320322111-1023131301113200-1102213013332011-1213031130233303-2311321321203130-2101213331332202)
- aws_parameters.az_nodes.reserved_inside_subnet

<a id="canonical-2202321103103310-2121123330102022-2333002300103221-1003312211231033-2320333022223002-3110210111332020-2311203003021123-2020233313222033"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for reserved inside subnet.

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

<a id="canonical-0201322230030201-2122002121320221-2002122003212300-2203003103300300-2313221233131012-2131130330012202-0202103210310000-1100321221212011"></a>

## Direct properties — reserved_inside_subnet / 121110311013 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1113113113210101-3010310212300203-0202001222231313-3002132133102002-3201232030113330-1333321000032302-2322201223021330-3021232103010303"></a>

## Next pages — reserved_inside_subnet / 121110311013 / 4

- [aws_parameters.az_nodes](data-sources--aws_tgw_site--reference--group-001.md#canonical-0300223333010011-0113000202023103-2021003320322111-1023131301113200-1102213013332011-1213031130233303-2311321321203130-2101213331332202)
- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)

<a id="canonical-2001023333111231-1300333312131200-2123102003330300-1321222313100103-0300131320232133-0333313233211023-0323320201232213-3213103210002203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1311211320230210-1101321123003122-1130123011203320-3010302221120302-1121220100230110-0102003130201323-0201012211121211-2232013211023210"></a>

## aws_parameters.az_nodes.workload_subnet — workload_subnet / 002200330121 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-1212321311102123-0132200012312320-2312231211033230-2200200312230233-3132320011023221-0111320212132020-3002000002020330-2223302203200011)
- [aws_parameters](data-sources--aws_tgw_site--reference--group-001.md#canonical-3311120101031201-0112330033331111-0330012133130331-2122223112223133-1300102323001230-2222231003110310-3112001230233232-2000103311201302)
- [aws_parameters.az_nodes](data-sources--aws_tgw_site--reference--group-001.md#canonical-0300223333010011-0113000202023103-2021003320322111-1023131301113200-1102213013332011-1213031130233303-2311321321203130-2101213331332202)
- aws_parameters.az_nodes.workload_subnet

<a id="canonical-2111032322322011-2321033001021233-2202000210232213-2233200212302300-3230132232222323-0232211031131130-2023311122002120-1212332002310002"></a>

Type: `"single"`. Computed.

Configuration parameter for workload subnet.

Upstream description:

Parameters for AWS subnet.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-choice": "[\"existing_subnet_id\",\"subnet_param\"]"
}
```

<a id="canonical-3132002200320030-3210221302231313-3212200203011011-0103311230301201-2113210322023220-2313223212220221-0321121222002021-0101022011123302"></a>

## Direct properties — workload_subnet / 002200330121 / 3

<a id="canonical-2323120131233121-0213313303020121-1012030103310113-1313023112000312-3012121103131130-2232312032220322-2031103010321121-1011021132001010"></a>

<a id="canonical-2130310213103333-3333203210303313-3220113332331033-2103100011310323-0322200200220333-3001000303321300-3020100110302330-1213322203132012"></a>

## existing_subnet_id property — workload_subnet / 002200330121 / 4

Type: `"string"`. Computed.

Exclusive with \[subnet\_param\] Information about existing subnet ID.

Upstream description:

Exclusive with \[subnet\_param\] Information about existing subnet ID.

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
    "pattern": "^(subnet-)([a-z0-9]{8}|[a-z0-9]{17})$"
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.pattern": "^(subnet-)([a-z0-9]{8}|[a-z0-9]{17})$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.pattern": "^(subnet-)([a-z0-9]{8}|[a-z0-9]{17})$"
  }
}
```

- [subnet_param](data-sources--aws_tgw_site--reference--group-001.md#canonical-2300310230300231-0030210001211030-2300200223210223-2211100301322113-0030101102322213-1310311232131111-0013230020123223-0122021220313130): complete subsection reference.

<a id="canonical-0201122022230231-2020222112223330-2222121003310312-2113011000311200-1002121223031321-2231003200121232-1302132013122201-1232032313122123"></a>

## Next pages — workload_subnet / 002200330121 / 5

- [aws_parameters.az_nodes.workload_subnet.subnet_param](data-sources--aws_tgw_site--reference--group-001.md#canonical-2300310230300231-0030210001211030-2300200223210223-2211100301322113-0030101102322213-1310311232131111-0013230020123223-0122021220313130)
- [aws_parameters.az_nodes](data-sources--aws_tgw_site--reference--group-001.md#canonical-0300223333010011-0113000202023103-2021003320322111-1023131301113200-1102213013332011-1213031130233303-2311321321203130-2101213331332202)
- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)

<a id="canonical-2300310230300231-0030210001211030-2300200223210223-2211100301322113-0030101102322213-1310311232131111-0013230020123223-0122021220313130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1020222130201002-2200103332032010-2121023103233103-1212320202031022-0301300222000123-2233213113321320-2302231023023321-3313002230030200"></a>

## aws_parameters.az_nodes.workload_subnet.subnet_param — subnet_param / 031212003222 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-1212321311102123-0132200012312320-2312231211033230-2200200312230233-3132320011023221-0111320212132020-3002000002020330-2223302203200011)
- [aws_parameters](data-sources--aws_tgw_site--reference--group-001.md#canonical-3311120101031201-0112330033331111-0330012133130331-2122223112223133-1300102323001230-2222231003110310-3112001230233232-2000103311201302)
- [aws_parameters.az_nodes](data-sources--aws_tgw_site--reference--group-001.md#canonical-0300223333010011-0113000202023103-2021003320322111-1023131301113200-1102213013332011-1213031130233303-2311321321203130-2101213331332202)
- [aws_parameters.az_nodes.workload_subnet](data-sources--aws_tgw_site--reference--group-001.md#canonical-2001023333111231-1300333312131200-2123102003330300-1321222313100103-0300131320232133-0333313233211023-0323320201232213-3213103210002203)
- aws_parameters.az_nodes.workload_subnet.subnet_param

<a id="canonical-0132330200131113-2010122310001312-0330301001322332-1232021111220022-3300121121013123-0131223233313300-0212220210010021-1133030213033133"></a>

Type: `"single"`. Computed.

Parameters for creating a new cloud subnet.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2132232122123300-1002231003222210-3120121232001201-2302101012130322-3013332121113310-1333013012210333-1200001233331232-2223233131031000"></a>

## Direct properties — subnet_param / 031212003222 / 3

<a id="canonical-3012111010100301-3012032211222203-2300223222022030-3210103323323213-2333231001332313-0020313000223202-2120001102100332-0031223132332021"></a>

<a id="canonical-3132310311113310-3331031000001030-0233102013111020-3101323230120310-3020230011311103-0132001210131130-3213312110123000-0113221310312021"></a>

## IPv4 property — subnet_param / 031212003222 / 4

Type: `"string"`. Computed.

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

<a id="canonical-3230122020311002-2020111113302332-2120033301302012-1331101003120100-0212230313322100-1111021221211320-0301011310232001-0203230102100012"></a>

## Next pages — subnet_param / 031212003222 / 5

- [aws_parameters.az_nodes.workload_subnet](data-sources--aws_tgw_site--reference--group-001.md#canonical-2001023333111231-1300333312131200-2123102003330300-1321222313100103-0300131320232133-0333313233211023-0323320201232213-3213103210002203)
- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)

<a id="canonical-1200330320023230-0322020333133231-1333202000222201-1300210113301131-3022222013031320-2323003101133112-0221332102112021-1220220320212322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1001122002332123-2202103132303332-0103311121132213-3012323322102010-0113200011313131-2211033213010120-1212231210003332-2133313100202311"></a>

## aws_parameters.custom_security_group — custom_security_group / 101232021231 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-1212321311102123-0132200012312320-2312231211033230-2200200312230233-3132320011023221-0111320212132020-3002000002020330-2223302203200011)
- [aws_parameters](data-sources--aws_tgw_site--reference--group-001.md#canonical-3311120101031201-0112330033331111-0330012133130331-2122223112223133-1300102323001230-2222231003110310-3112001230233232-2000103311201302)
- aws_parameters.custom_security_group

<a id="canonical-0211103031320333-2012211311010323-3312130011312130-0012022230012112-0230122320303133-2233311233203023-2131000233321011-1021131200201313"></a>

Type: `"single"`. Computed.

Enter pre created security groups for slo(Site Local Outside) and sli(Site Local Inside) interface.
Supported only for sites deployed on existing VPC.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1001101003303313-1220030323132012-3003230100110110-0232002313120002-1101302000222330-0130312000302021-1030120311202020-0103112230302211"></a>

## Direct properties — custom_security_group / 101232021231 / 3

<a id="canonical-1100202332213102-1313031231213130-1131120130303111-0112100301112101-2001323123232300-3313002023011222-2033301331310012-0331020202122212"></a>

<a id="canonical-0102002131213310-1133331130211333-0113031001122021-0332301100100021-1230130233303001-2301232300302033-0030133223301203-3021102331100333"></a>

## inside_security_group_id property — custom_security_group / 101232021231 / 4

Type: `"string"`. Computed.

Security Group ID to be attached to SLI(Site Local Inside) Interface.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 20,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 20,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "pattern": "^(sg-)([a-z0-9]{8}|[a-z0-9]{17})$|^$"
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "20",
    "ves.io.schema.rules.string.pattern": "^(sg-)([a-z0-9]{8}|[a-z0-9]{17})$|^$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "20",
    "ves.io.schema.rules.string.pattern": "^(sg-)([a-z0-9]{8}|[a-z0-9]{17})$|^$"
  }
}
```

<a id="canonical-2003120203131112-2311301101331331-3033332332232010-0011113212013030-1120012230332313-0322111133031021-1303210223323030-1321013112322322"></a>

<a id="canonical-2223312332001333-0313312303023133-3233102210010031-2300321331232221-3123213333312111-0122220013222223-0131310131223023-3011211122132001"></a>

## outside_security_group_id property — custom_security_group / 101232021231 / 5

Type: `"string"`. Computed.

Security Group ID to be attached to SLO(Site Local Outside) Interface.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 20,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 20,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "pattern": "^(sg-)([a-z0-9]{8}|[a-z0-9]{17})$|^$"
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "20",
    "ves.io.schema.rules.string.pattern": "^(sg-)([a-z0-9]{8}|[a-z0-9]{17})$|^$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "20",
    "ves.io.schema.rules.string.pattern": "^(sg-)([a-z0-9]{8}|[a-z0-9]{17})$|^$"
  }
}
```

<a id="canonical-2111000110012311-2313122232003010-3011300303201003-1012333231331030-1000332222310011-0033200021210002-1112332000212033-3100022012011010"></a>

## Next pages — custom_security_group / 101232021231 / 6

- [aws_parameters](data-sources--aws_tgw_site--reference--group-001.md#canonical-3311120101031201-0112330033331111-0330012133130331-2122223112223133-1300102323001230-2222231003110310-3112001230233232-2000103311201302)
- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)

<a id="canonical-2332232011200302-1331212102222001-3302032333323332-3333033110331323-1123030122233131-1023032022330333-0330100101333210-2033132031213230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1301300331120132-1013202203123030-3301311323132030-1023031030301212-3023110031001020-1201330203223101-3313321020212122-1032001330203103"></a>

## aws_parameters.disable_encryption — disable_encryption / 121222213333 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-1212321311102123-0132200012312320-2312231211033230-2200200312230233-3132320011023221-0111320212132020-3002000002020330-2223302203200011)
- [aws_parameters](data-sources--aws_tgw_site--reference--group-001.md#canonical-3311120101031201-0112330033331111-0330012133130331-2122223112223133-1300102323001230-2222231003110310-3112001230233232-2000103311201302)
- aws_parameters.disable_encryption

<a id="canonical-1100201031001211-3102020200012001-3311002022202321-1320333221221202-1020012021111330-2312030010003231-2232030000111331-3213031001023322"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for disable encryption.

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

<a id="canonical-1013322311333300-0002120232111101-2233000121333000-0101201133002332-3113122212112232-2310200033300003-2233133220020230-3110022031230222"></a>

## Direct properties — disable_encryption / 121222213333 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0033212232000200-2223311210303001-3332003221003210-0121202301313003-3030210012311033-1202123121220311-1222130313121211-1213112320200322"></a>

## Next pages — disable_encryption / 121222213333 / 4

- [aws_parameters](data-sources--aws_tgw_site--reference--group-001.md#canonical-3311120101031201-0112330033331111-0330012133130331-2122223112223133-1300102323001230-2222231003110310-3112001230233232-2000103311201302)
- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)

<a id="canonical-1023220201100022-2321100222123310-3300233121120123-0322013200031311-1122223320323313-3130123020102012-1203103303102113-0202301101203333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0102133223213330-2310322323302013-3130302313231320-0303323101031021-2102020321111022-3231103122133222-2211131011303113-3001212220311223"></a>

## aws_parameters.disable_internet_vip — disable_internet_vip / 320333213130 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-1212321311102123-0132200012312320-2312231211033230-2200200312230233-3132320011023221-0111320212132020-3002000002020330-2223302203200011)
- [aws_parameters](data-sources--aws_tgw_site--reference--group-001.md#canonical-3311120101031201-0112330033331111-0330012133130331-2122223112223133-1300102323001230-2222231003110310-3112001230233232-2000103311201302)
- aws_parameters.disable_internet_vip

<a id="canonical-2301213032023113-3223313013232230-2122000200113210-0333103311302230-0130100132300313-1302010012331121-2331030130132030-3011310322020320"></a>

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

<a id="canonical-0321211211103333-2033110311300133-3032023321113232-3003123300002213-1102011233323020-1122312313200130-1301302211232300-3001121211212323"></a>

## Direct properties — disable_internet_vip / 320333213130 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1220231010322001-1203023332003220-3122232332031333-2033033221021112-3103031221013111-3020302313310302-2131222323102031-2211130003111220"></a>

## Next pages — disable_internet_vip / 320333213130 / 4

- [aws_parameters](data-sources--aws_tgw_site--reference--group-001.md#canonical-3311120101031201-0112330033331111-0330012133130331-2122223112223133-1300102323001230-2222231003110310-3112001230233232-2000103311201302)
- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)

<a id="canonical-3010302002120100-2103002021320330-3110033102312003-2010320201333323-3230020212131301-2330231333302020-1201123130210012-0300222330022213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2033130100222313-2233232222202220-2302013313001132-2130020322032210-3301123320323003-3330210122103302-0020023133000220-3201223332123230"></a>

## aws_parameters.enable_encryption — enable_encryption / 121133300303 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-1212321311102123-0132200012312320-2312231211033230-2200200312230233-3132320011023221-0111320212132020-3002000002020330-2223302203200011)
- [aws_parameters](data-sources--aws_tgw_site--reference--group-001.md#canonical-3311120101031201-0112330033331111-0330012133130331-2122223112223133-1300102323001230-2222231003110310-3112001230233232-2000103311201302)
- aws_parameters.enable_encryption

<a id="canonical-1023000101100030-0132023223122001-0100332011031111-3313212000120332-0002200120211121-3232032200032330-0012322012133100-3311210022223113"></a>

Type: `"single"`. Computed.

Configuration parameter for enable encryption.

Upstream description:

Information related to disk encryption.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2311303210023003-3231122230000331-3202112033020131-1231211300103031-0001202333333222-0331200112211321-2120020033201120-0330310112020200"></a>

## Direct properties — enable_encryption / 121133300303 / 3

<a id="canonical-1223211020123121-1332121101302312-1130303221133321-3023113021000332-2022230312023021-1202231212230013-0321010330210033-1310233031032012"></a>

<a id="canonical-1000113033113233-2230220312332211-0132333300123023-3302333322332121-1101232100300212-1122322212210020-2001201321000302-0102001030123313"></a>

## kms_key_id property — enable_encryption / 121133300303 / 4

Type: `"string"`. Computed.

AWS KMS Key to be used to encrypt the disk attached to the VM.

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
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  }
}
```

<a id="canonical-2200133131011233-0122010000102211-2122220133201110-1010021301221022-3302231232301212-3011221021102031-1100220302033131-0012210200302300"></a>

## Next pages — enable_encryption / 121133300303 / 5

- [aws_parameters](data-sources--aws_tgw_site--reference--group-001.md#canonical-3311120101031201-0112330033331111-0330012133130331-2122223112223133-1300102323001230-2222231003110310-3112001230233232-2000103311201302)
- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)

<a id="canonical-0122200323013210-2301312003231211-3101311301201122-2120230202020100-0323212320030013-0332302101213300-2231312100322202-2213132110133310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0303311321113333-0120110332321021-3211321121200102-3302231133201131-1113031331210302-1100213100313322-2130123102203002-0020222323201111"></a>

## aws_parameters.enable_internet_vip — enable_internet_vip / 303223230000 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-1212321311102123-0132200012312320-2312231211033230-2200200312230233-3132320011023221-0111320212132020-3002000002020330-2223302203200011)
- [aws_parameters](data-sources--aws_tgw_site--reference--group-001.md#canonical-3311120101031201-0112330033331111-0330012133130331-2122223112223133-1300102323001230-2222231003110310-3112001230233232-2000103311201302)
- aws_parameters.enable_internet_vip

<a id="canonical-0312201001302320-3302001312323332-1111132330010322-2222222013202010-1103310312233320-2312322131311102-2332230211223031-3310103120013131"></a>

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

<a id="canonical-1212102103102233-1331101130032011-3113200233022230-0012321010232003-0103321232320223-2203003131312110-1210330332020112-3302211022131003"></a>

## Direct properties — enable_internet_vip / 303223230000 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2112021323010323-0111120233103300-2202031333000102-2100331103223331-3311121223123232-2000011301222113-3112001132100023-2231332021331220"></a>

## Next pages — enable_internet_vip / 303223230000 / 4

- [aws_parameters](data-sources--aws_tgw_site--reference--group-001.md#canonical-3311120101031201-0112330033331111-0330012133130331-2122223112223133-1300102323001230-2222231003110310-3112001230233232-2000103311201302)
- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)

<a id="canonical-2022022330331132-1202301000122022-0131230132121103-2300033201303301-1103312012020121-1311322301301210-0321211103202233-0213003312233030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2330202130001210-0312100010211202-2131120201223113-2012023200022210-0212133332133020-2013113301030121-2303222321113320-0323101223233020"></a>

## aws_parameters.existing_tgw — existing_tgw / 232112200112 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-1212321311102123-0132200012312320-2312231211033230-2200200312230233-3132320011023221-0111320212132020-3002000002020330-2223302203200011)
- [aws_parameters](data-sources--aws_tgw_site--reference--group-001.md#canonical-3311120101031201-0112330033331111-0330012133130331-2122223112223133-1300102323001230-2222231003110310-3112001230233232-2000103311201302)
- aws_parameters.existing_tgw

<a id="canonical-3301211022320302-1023200231031032-0000201231223332-0121301310113002-3120001111123132-2111333001133210-1313003230330221-0303232032122000"></a>

Type: `"single"`. Computed.

Configuration parameter for existing tgw.

Upstream description:

Information needed for existing TGW.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0013212300032331-1200201320033310-3121102131322230-2131231111002330-2032230012330221-1021021231212310-2110331211111031-0113313210121021"></a>

## Direct properties — existing_tgw / 232112200112 / 3

<a id="canonical-2022111301330033-2311032322203330-2210203013232000-3320333333310310-0100213220002031-2012101311011100-3322333300333322-3121222021303333"></a>

<a id="canonical-2301132211110233-1101110012112333-0302010311312233-0322232313012102-0322122200103123-1112110213120021-1301110320221301-1010033111213312"></a>

## tgw_asn property — existing_tgw / 232112200112 / 4

Type: `"number"`. Computed.

Enter TGW ASN. TGW ASN.

Upstream description:

TGW ASN.

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

<a id="canonical-0232031220112233-1221130131000323-1301230330212321-2330131310031113-2303213203302230-0131310101012321-2133323110332311-0102320132001230"></a>

<a id="canonical-0323003330121321-3302203311222330-0110013013302321-1310331200300101-0003011202103330-1121310112122112-3223000203133222-1120103120302233"></a>

## tgw_id property — existing_tgw / 232112200112 / 5

Type: `"string"`. Computed.

Existing TGW ID. Existing TGW ID.

Upstream description:

Existing TGW ID.

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
    "pattern": "^(tgw-)([a-z0-9]{8}|[a-z0-9]{17})$"
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.pattern": "^(tgw-)([a-z0-9]{8}|[a-z0-9]{17})$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.pattern": "^(tgw-)([a-z0-9]{8}|[a-z0-9]{17})$"
  }
}
```

<a id="canonical-3032333013310113-3101020221121001-3220021001201321-2222132000223203-0313312322320331-0003111001202021-3213312122202230-1020331002211122"></a>

<a id="canonical-2201220303112212-1223132031032323-0032000032232222-3120010312333102-1111202313131231-0131103133300000-2003312211203011-1011202211201020"></a>

## volterra_site_asn property — existing_tgw / 232112200112 / 6

Type: `"number"`. Computed.

Enter F5XC Site ASN. F5XC Site ASN.

Upstream description:

F5XC Site ASN.

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

<a id="canonical-2310300003321302-0123132003331233-2121212101302110-2123303322100330-0221133310211012-1133310201210013-2123122110031303-3111020333313310"></a>

## Next pages — existing_tgw / 232112200112 / 7

- [aws_parameters](data-sources--aws_tgw_site--reference--group-001.md#canonical-3311120101031201-0112330033331111-0330012133130331-2122223112223133-1300102323001230-2222231003110310-3112001230233232-2000103311201302)
- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)

<a id="canonical-2111203120120332-0311322131220100-1310033300211130-0302112330331312-0000101223200133-3210013202322002-0101000113321133-3320021300310020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2333213231133320-3223123122130323-2312322003020102-0132033033010001-1003323220203231-0330120122023103-0313120122002203-2200322002222022"></a>

## aws_parameters.f5xc_security_group — f5xc_security_group / 302101011021 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-1212321311102123-0132200012312320-2312231211033230-2200200312230233-3132320011023221-0111320212132020-3002000002020330-2223302203200011)
- [aws_parameters](data-sources--aws_tgw_site--reference--group-001.md#canonical-3311120101031201-0112330033331111-0330012133130331-2122223112223133-1300102323001230-2222231003110310-3112001230233232-2000103311201302)
- aws_parameters.f5xc_security_group

<a id="canonical-0333212322020010-1231220133203132-3110002213030232-0021301011321031-3301221333232012-2011220123022311-2000030220322120-1330320102012120"></a>

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

<a id="canonical-3002103320112321-0030230202232012-0231112123313001-2310230202130220-2130302220023212-3012321132121032-1111322332023333-1012100100213300"></a>

## Direct properties — f5xc_security_group / 302101011021 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3230230213012011-2202103013113232-0320000201311030-0223213031331233-2130121131210010-3132333311220333-3120123331210330-0312223033001132"></a>

## Next pages — f5xc_security_group / 302101011021 / 4

- [aws_parameters](data-sources--aws_tgw_site--reference--group-001.md#canonical-3311120101031201-0112330033331111-0330012133130331-2122223112223133-1300102323001230-2222231003110310-3112001230233232-2000103311201302)
- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)

<a id="canonical-1232303130313003-0000010222131313-2230032301201222-3310233212211201-1303130113131201-2022121312300000-0120130132002311-3330030310301103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1132132330111020-1223113133012231-0230331101320332-0202110223333000-1031133013013013-3222212303232030-1133331303233222-3312201110000300"></a>

## aws_parameters.new_tgw — new_tgw / 101313310330 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-1212321311102123-0132200012312320-2312231211033230-2200200312230233-3132320011023221-0111320212132020-3002000002020330-2223302203200011)
- [aws_parameters](data-sources--aws_tgw_site--reference--group-001.md#canonical-3311120101031201-0112330033331111-0330012133130331-2122223112223133-1300102323001230-2222231003110310-3112001230233232-2000103311201302)
- aws_parameters.new_tgw

<a id="canonical-3030002000022203-2001200200323322-3111132201320132-3313120231103022-1123312203032011-3023033321111002-3300023020023313-2122213313213200"></a>

Type: `"single"`. Computed.

TGWParamsType.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-asn_choice": "[\"system_generated\",\"user_assigned\"]"
}
```

<a id="canonical-0233302122300331-0131313101301000-1131300310302301-3021200222112332-3300212013101302-2331321223213303-3311323200231201-3020302000232332"></a>

## Direct properties — new_tgw / 101313310330 / 3

- [system_generated](data-sources--aws_tgw_site--reference--group-001.md#canonical-0223113321211100-3331033131103000-2022021223322202-2332222301220320-0131021012311303-0223011311231301-2122331311211112-1321203302100002): complete subsection reference.

- [user_assigned](data-sources--aws_tgw_site--reference--group-001.md#canonical-3033000002110200-3033022233303022-3300212013011111-2102201030002331-0311023013100201-0221032010201101-1211022300233331-3331121021211111): complete subsection reference.

<a id="canonical-2332223132320031-1302232300013130-3333231333302313-0302330021002113-1202120210202121-2022010231330110-0221001103231303-2303202102132212"></a>

## Next pages — new_tgw / 101313310330 / 4

- [aws_parameters.new_tgw.system_generated](data-sources--aws_tgw_site--reference--group-001.md#canonical-0223113321211100-3331033131103000-2022021223322202-2332222301220320-0131021012311303-0223011311231301-2122331311211112-1321203302100002)
- [aws_parameters.new_tgw.user_assigned](data-sources--aws_tgw_site--reference--group-001.md#canonical-3033000002110200-3033022233303022-3300212013011111-2102201030002331-0311023013100201-0221032010201101-1211022300233331-3331121021211111)
- [aws_parameters](data-sources--aws_tgw_site--reference--group-001.md#canonical-3311120101031201-0112330033331111-0330012133130331-2122223112223133-1300102323001230-2222231003110310-3112001230233232-2000103311201302)
- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)

<a id="canonical-0223113321211100-3331033131103000-2022021223322202-2332222301220320-0131021012311303-0223011311231301-2122331311211112-1321203302100002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2220113001120131-2232222000233121-0213131332333212-0331031222032323-0131300002111011-3300122220133000-0011100000331122-3320310010200323"></a>

## aws_parameters.new_tgw.system_generated — system_generated / 030032013310 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-1212321311102123-0132200012312320-2312231211033230-2200200312230233-3132320011023221-0111320212132020-3002000002020330-2223302203200011)
- [aws_parameters](data-sources--aws_tgw_site--reference--group-001.md#canonical-3311120101031201-0112330033331111-0330012133130331-2122223112223133-1300102323001230-2222231003110310-3112001230233232-2000103311201302)
- [aws_parameters.new_tgw](data-sources--aws_tgw_site--reference--group-001.md#canonical-1232303130313003-0000010222131313-2230032301201222-3310233212211201-1303130113131201-2022121312300000-0120130132002311-3330030310301103)
- aws_parameters.new_tgw.system_generated

<a id="canonical-0000201331313213-3020311311101323-3210231302033130-1232112103012020-2001113031101212-2203023012032132-0222021230211330-2002212013023203"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for system generated.

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

<a id="canonical-2231210010122130-3031202020010333-2323303013112322-2022203322011220-0102112230103221-1101021301331101-1331200113031133-1313323221032011"></a>

## Direct properties — system_generated / 030032013310 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2130023033102200-1321301112221100-1033203132000320-3011032321110221-2033032302112121-0120122320322303-0221113213212030-2313023333100202"></a>

## Next pages — system_generated / 030032013310 / 4

- [aws_parameters.new_tgw](data-sources--aws_tgw_site--reference--group-001.md#canonical-1232303130313003-0000010222131313-2230032301201222-3310233212211201-1303130113131201-2022121312300000-0120130132002311-3330030310301103)
- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)

<a id="canonical-3033000002110200-3033022233303022-3300212013011111-2102201030002331-0311023013100201-0221032010201101-1211022300233331-3331121021211111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3210303211210111-1322010020031300-3102020033033210-2122210322013011-3132210023201310-2120113102301011-2122001220012031-3013311313303312"></a>

## aws_parameters.new_tgw.user_assigned — user_assigned / 133030103231 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-1212321311102123-0132200012312320-2312231211033230-2200200312230233-3132320011023221-0111320212132020-3002000002020330-2223302203200011)
- [aws_parameters](data-sources--aws_tgw_site--reference--group-001.md#canonical-3311120101031201-0112330033331111-0330012133130331-2122223112223133-1300102323001230-2222231003110310-3112001230233232-2000103311201302)
- [aws_parameters.new_tgw](data-sources--aws_tgw_site--reference--group-001.md#canonical-1232303130313003-0000010222131313-2230032301201222-3310233212211201-1303130113131201-2022121312300000-0120130132002311-3330030310301103)
- aws_parameters.new_tgw.user_assigned

<a id="canonical-1102002221000233-2203002321323021-2133130010122230-2023131320300121-2312023033110312-1230031210101331-2011011130223111-1313113002030033"></a>

Type: `"single"`. Computed.

Information needed when ASNs are assigned by the user.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3102330330011101-2003231330303111-1223012220300323-2033313031032231-3002230102003132-0332032321030033-0111223110113213-0212112203103123"></a>

## Direct properties — user_assigned / 133030103231 / 3

<a id="canonical-1112221010133212-1021130310202000-1211203230032311-1101313330323120-1010131320110303-3121013110113303-3011331321020023-0000220332120013"></a>

<a id="canonical-1130101122223321-3210111323231231-2312311231121212-2130132100310203-1230220202232112-2220022113130213-1321311222112210-2010211201312123"></a>

## tgw_asn property — user_assigned / 133030103231 / 4

Type: `"number"`. Computed.

TGW ASN. Allowed range for 16-bit private ASNs include 64512 to 65534.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 65534,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 64513
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gt": "64512",
    "ves.io.schema.rules.uint32.lte": "65534"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gt": "64512",
    "ves.io.schema.rules.uint32.lte": "65534"
  }
}
```

<a id="canonical-1321112303323020-2330012301302102-1212221223310002-0221201231203031-0130023223013202-1230301122311032-0331321131020212-3302312012120230"></a>

<a id="canonical-1133000121021200-1011030303101021-2121210110021133-0212112321333023-0100220123200333-0222212202200332-0111031132130210-3003323232002200"></a>

## volterra_site_asn property — user_assigned / 133030103231 / 5

Type: `"number"`. Computed.

Enter F5XC Site ASN. F5XC Site ASN.

Upstream description:

F5XC Site ASN.

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

<a id="canonical-0310100202312130-1233223301300312-2133010003333313-1122002022301221-3321110313021212-1231202032322322-0212032230033032-1210133110012133"></a>

## Next pages — user_assigned / 133030103231 / 6

- [aws_parameters.new_tgw](data-sources--aws_tgw_site--reference--group-001.md#canonical-1232303130313003-0000010222131313-2230032301201222-3310233212211201-1303130113131201-2022121312300000-0120130132002311-3330030310301103)
- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)

<a id="canonical-2203311030300201-1320113131100101-3012322103333133-3003203313213332-2310302023313002-3013210130001312-2033332213210030-0030330322300233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->
