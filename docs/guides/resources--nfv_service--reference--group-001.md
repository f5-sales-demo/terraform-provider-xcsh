---
page_title: "xcsh_nfv_service reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_nfv_service reference."
---

# xcsh_nfv_service reference

<a id="canonical-1033200313012130-2303320020000220-3032203302003212-3030323031320111-0322032022323123-0203102131300133-1232321110231101-1022222010201320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Property reference

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-0223321113210321-0013210302112221-0313201132021001-1102330023301310-1131322132000230-3303202123330303-3122032322302202-2202303201230222)
- Property reference

<a id="canonical-1122223321213020-1322011321022330-0120001222130201-3212110301312321-0102301113222223-2100332200120010-3031111233112210-1303323320330112"></a>

### Direct properties for `xcsh_nfv_service`

<a id="canonical-2302211320202112-2313200123201322-0333333303023201-0332201111311323-0013223331220222-1001210110123132-1203302311302003-2311031222123020"></a>

#### `annotations` property

Type: `["map", "string"]`. Optional.

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

<a id="canonical-0232102132103133-1132310132021202-3011032120101230-3233031303103311-0100111301323301-1211202021102011-0313122010001123-3012233031112003"></a>

<a id="canonical-0210002133221010-0330333113322010-2033021301230313-2232032003131211-1301022003231310-3200020003233133-2333030011100012-0101311130232131"></a>

#### `description` property

Type: `"string"`. Optional.

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

<a id="canonical-0220012000221310-2132221022010120-0303223003121311-2320313320121122-1301211123023001-3232332100032331-0302200121123130-1001331122201202"></a>

<a id="canonical-3223031011312030-0010230012133313-2130102333120113-0331302130220023-0032010231002231-1033113220113023-2310011221003000-2232311110233220"></a>

#### `disable` property

Type: `"bool"`. Optional.

A value of true administratively disables the object.

Additional upstream details:

A value of true will administratively disable the object.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [disable_https_management](resources--nfv_service--reference--group-001.md#canonical-1213021122302111-2303132302331132-2033120022132312-0232210233303323-2001032101312213-3131012003332023-0101021121323031-2100213002233100): complete subsection reference.

- [disable_ssh_access](resources--nfv_service--reference--group-001.md#canonical-1121230113200220-3301032102311323-1313313310300012-0131301302121233-3323233302031130-1303122223201233-3332210000021300-3200231130103032): complete subsection reference.

- [enabled_ssh_access](resources--nfv_service--reference--group-001.md#canonical-1233323310221201-2110101010301322-3301231231110303-1003102112121320-1012103233232011-0000232331110323-2202130110221200-1133311203200010): complete subsection reference.

- [f5_big_ip_aws_service](resources--nfv_service--reference--group-001.md#canonical-2203212022212003-2212301132213201-2010231030110211-3120210322102100-3130230103331230-0332101002301003-2302320031110011-3121121001302210): complete subsection reference.

- [https_management](resources--nfv_service--reference--group-002.md#canonical-2012312013030011-0102312120121131-2122112110311110-1102330303132211-1031011331133302-3311032130133131-3022001211230302-3310030220221101): complete subsection reference.

<a id="canonical-2202210003230021-2123331330221213-2220013231110122-0233313101232312-2132301010233111-3331221220311002-3100330200012101-3230310003211130"></a>

<a id="canonical-2332000131213002-1201012330032012-1011233020323130-2013320002032310-3002002230230223-3301300310230223-1203313310322203-3210301133011323"></a>

#### `id` property

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-3123203312302020-1021113033213221-3110100013022111-1202032011011331-1202010201312110-3011233013030200-1102102033300321-0130311301000210"></a>

<a id="canonical-3100102312113213-3001212113202100-2010320120110101-3213022010203323-3332011332322303-1032311232212132-0133100030121002-2032002320110210"></a>

#### `labels` property

Type: `["map", "string"]`. Optional.

Labels is a user defined key-value map that can be attached to resources for organization and
filtering.

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

<a id="canonical-0113312013100323-3001203330313203-2130221120123200-2133232231013333-2131013222231223-0001230333223222-1113001322233220-3311213231123230"></a>

<a id="canonical-2302110101023213-0012102102232322-3202110121122001-1220200032003322-0212230203221132-2022333133111323-1013230311320023-3303213122110000"></a>

#### `name` property

Type: `"string"`. Required.

Name of the Nfv Service. Must be unique within the namespace.

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

<a id="canonical-2132211332310221-0321032100122201-0031303002231002-1333110022202310-2030213001222303-0001030222111003-3232120313210031-2300300123223022"></a>

<a id="canonical-3320130022200023-2113301213023323-2132130023113120-1122003232221031-0312011131023112-0131230311313200-3213330323313022-3221201210023202"></a>

#### `namespace` property

Type: `"string"`. Required.

Namespace where the Nfv Service is created.

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

- [palo_alto_fw_service](resources--nfv_service--reference--group-004.md#canonical-1220032300320010-0322310310320112-3102020032000031-0230002231133113-1321103031132221-2112211003001211-2001103020013300-3102213200033033): complete subsection reference.

- [timeouts](resources--nfv_service--reference--group-004.md#canonical-0202232021221122-0000133111310212-2020203112302000-1210111320033230-0001031130310130-3231013003311031-0221323010300330-3130003223303201): complete subsection reference.

<a id="canonical-1301132232301133-3303213000202200-2222211103122001-0322332233313333-3023011113110303-1122323333101230-0303123213203020-3313310033010122"></a>

### All schema paths for `xcsh_nfv_service`

Each exact path has one authoritative reference destination. Collection element indices are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](resources--nfv_service--reference--group-001.md#canonical-2302211320202112-2313200123201322-0333333303023201-0332201111311323-0013223331220222-1001210110123132-1203302311302003-2311031222123020) |
| `description` | [description](resources--nfv_service--reference--group-001.md#canonical-0232102132103133-1132310132021202-3011032120101230-3233031303103311-0100111301323301-1211202021102011-0313122010001123-3012233031112003) |
| `disable` | [disable](resources--nfv_service--reference--group-001.md#canonical-0220012000221310-2132221022010120-0303223003121311-2320313320121122-1301211123023001-3232332100032331-0302200121123130-1001331122201202) |
| `disable_https_management` | [disable_https_management](resources--nfv_service--reference--group-001.md#canonical-3112312130013010-0213122301312322-1233230202121013-2232313130212030-0022130101301033-0101103110230133-2130301211211102-1301123223121110) |
| `disable_ssh_access` | [disable_ssh_access](resources--nfv_service--reference--group-001.md#canonical-0111131333123130-2302200131001101-0111200312022021-3302110232120331-0010113121223001-2211222211321130-2020322122013323-2213211231011123) |
| `enabled_ssh_access` | [enabled_ssh_access](resources--nfv_service--reference--group-001.md#canonical-0201302210200021-1010011310120203-1323221101222200-0312020201221132-0012233110231012-0231010102020210-3032213233301121-3323230130222103) |
| `enabled_ssh_access.advertise_on_sli` | [enabled_ssh_access.advertise_on_sli](resources--nfv_service--reference--group-001.md#canonical-0003013100101131-1122201201230222-1001220013030203-3103121133102032-2013021223011301-0010320102213130-1230031332301211-0111323031201032) |
| `enabled_ssh_access.advertise_on_slo` | [enabled_ssh_access.advertise_on_slo](resources--nfv_service--reference--group-001.md#canonical-3113030011011303-1203312332110201-1123123333313002-2220022010321112-3132022210232301-2322311232111112-0120331201000303-0302121303030110) |
| `enabled_ssh_access.advertise_on_slo_sli` | [enabled_ssh_access.advertise_on_slo_sli](resources--nfv_service--reference--group-001.md#canonical-2301123112320330-1333002212323201-0030013311233203-1100132101100002-2003321311300011-1110200320302232-0231130202120003-1231233022300232) |
| `enabled_ssh_access.domain_suffix` | [enabled_ssh_access.domain_suffix](resources--nfv_service--reference--group-001.md#canonical-2020201313303300-1010202033210223-2101120030023332-1023230302232323-3300231213122233-1000113021111020-0111233220213221-2210102333012311) |
| `enabled_ssh_access.node_ssh_ports` | [enabled_ssh_access.node_ssh_ports](resources--nfv_service--reference--group-001.md#canonical-3131212313033221-1031232020033032-2212111001221102-1233323330113003-3201322203113012-2002120102101320-1230123231121003-0222202133311103) |
| `enabled_ssh_access.node_ssh_ports.node_name` | [enabled_ssh_access.node_ssh_ports.node_name](resources--nfv_service--reference--group-001.md#canonical-0311023330222223-2231312201023331-1310102133100122-3323001122303210-2221201003011301-0013011013323001-0331130103231201-1302311300110220) |
| `enabled_ssh_access.node_ssh_ports.ssh_port` | [enabled_ssh_access.node_ssh_ports.ssh_port](resources--nfv_service--reference--group-001.md#canonical-3213021100333102-0320320221303020-2011100001021130-3322000022023233-0103300220122031-1221211101011311-0311030210021112-1130322030113320) |
| `f5_big_ip_aws_service` | [f5_big_ip_aws_service](resources--nfv_service--reference--group-001.md#canonical-0201100230223303-2001123313212330-0023112220313003-3331311211332102-0202122111230001-2232111322331013-1010031220331333-3331233032322321) |
| `f5_big_ip_aws_service.admin_password` | [f5_big_ip_aws_service.admin_password](resources--nfv_service--reference--group-001.md#canonical-0020332213331212-2101220311133313-2122212010123301-0233103012221122-3212102010332130-3113012033002232-2111100123120310-1023112203303112) |
| `f5_big_ip_aws_service.admin_password.blindfold_secret_info` | [f5_big_ip_aws_service.admin_password.blindfold_secret_info](resources--nfv_service--reference--group-001.md#canonical-3301022230121212-1231300113010201-0231103320022010-2213102330001131-3120203022202301-0012023232321301-3331313021332130-3220130023020300) |
| `f5_big_ip_aws_service.admin_password.blindfold_secret_info.decryption_provider` | [f5_big_ip_aws_service.admin_password.blindfold_secret_info.decryption_provider](resources--nfv_service--reference--group-001.md#canonical-2321300332103030-2323111232301113-2202003220003121-0302121112122102-0113131322320101-0022103312220333-1230113231302133-0220203300110330) |
| `f5_big_ip_aws_service.admin_password.blindfold_secret_info.location` | [f5_big_ip_aws_service.admin_password.blindfold_secret_info.location](resources--nfv_service--reference--group-001.md#canonical-3311303232101323-3130320010021320-3332011001012212-0131000323200013-1220100031022101-2230010020130113-3003203203202220-2320230320323001) |
| `f5_big_ip_aws_service.admin_password.blindfold_secret_info.store_provider` | [f5_big_ip_aws_service.admin_password.blindfold_secret_info.store_provider](resources--nfv_service--reference--group-001.md#canonical-2310103312222302-3132133201331103-1321123300002201-3212313320121213-3023023111211220-1201220230110221-3303233122333130-3000100322001131) |
| `f5_big_ip_aws_service.admin_password.clear_secret_info` | [f5_big_ip_aws_service.admin_password.clear_secret_info](resources--nfv_service--reference--group-001.md#canonical-3313101212220110-0010011312222201-2232103102110313-2330321322322120-0123333232120003-0102122233312221-1133203201120210-3203132203121302) |
| `f5_big_ip_aws_service.admin_password.clear_secret_info.provider_ref` | [f5_big_ip_aws_service.admin_password.clear_secret_info.provider_ref](resources--nfv_service--reference--group-001.md#canonical-2323202200030121-1000333130332100-0122030331233120-1131323222303220-3322321130010322-3102332223302030-3132011203313312-1303202220231201) |
| `f5_big_ip_aws_service.admin_password.clear_secret_info.url` | [f5_big_ip_aws_service.admin_password.clear_secret_info.url](resources--nfv_service--reference--group-001.md#canonical-1033213333031230-3023320130220111-0210002233210013-0132023210312332-2131031103222220-0133022322130021-0013212302020120-0011333001110300) |
| `f5_big_ip_aws_service.admin_username` | [f5_big_ip_aws_service.admin_username](resources--nfv_service--reference--group-001.md#canonical-0303303210230013-0010231111233233-3000013023023101-1112213102230002-2303033320012311-3320331223211221-3021300202321023-3012121221222220) |
| `f5_big_ip_aws_service.aws_tgw_site_params` | [f5_big_ip_aws_service.aws_tgw_site_params](resources--nfv_service--reference--group-001.md#canonical-3320313103200132-0103010113210101-1101031002101012-2010331023022021-2333210311111022-0030031113000212-3001123013330231-2331031203101333) |
| `f5_big_ip_aws_service.aws_tgw_site_params.aws_tgw_site` | [f5_big_ip_aws_service.aws_tgw_site_params.aws_tgw_site](resources--nfv_service--reference--group-001.md#canonical-0003032121222113-0011112101002302-2301303213332311-0110221102323033-1033022110331321-1232332133211020-1113312321331210-2112101100020220) |
| `f5_big_ip_aws_service.aws_tgw_site_params.aws_tgw_site.name` | [f5_big_ip_aws_service.aws_tgw_site_params.aws_tgw_site.name](resources--nfv_service--reference--group-001.md#canonical-1010331213000212-2221032202201100-1020122332020122-1321312121310021-3223131031131020-2023120022023130-3223301320323103-0030030210112312) |
| `f5_big_ip_aws_service.aws_tgw_site_params.aws_tgw_site.namespace` | [f5_big_ip_aws_service.aws_tgw_site_params.aws_tgw_site.namespace](resources--nfv_service--reference--group-001.md#canonical-0311223030320111-2331221013210111-0222200302220102-1212313101302221-0010212231303100-0122210313220220-2022301220113120-1021032320133212) |
| `f5_big_ip_aws_service.aws_tgw_site_params.aws_tgw_site.tenant` | [f5_big_ip_aws_service.aws_tgw_site_params.aws_tgw_site.tenant](resources--nfv_service--reference--group-001.md#canonical-2130333010312310-2122231331023013-1203010000232020-3020113331310320-3302001002000130-1331132211111103-3033321223133312-2111131332222222) |
| `f5_big_ip_aws_service.endpoint_service` | [f5_big_ip_aws_service.endpoint_service](resources--nfv_service--reference--group-001.md#canonical-1032321201122030-3112023110013020-2120212102023120-1301023013310313-1323131301203220-1231311320221303-3203101100210331-2002323023100113) |
| `f5_big_ip_aws_service.endpoint_service.advertise_on_slo_ip` | [f5_big_ip_aws_service.endpoint_service.advertise_on_slo_ip](resources--nfv_service--reference--group-001.md#canonical-2021020020222320-3012232220211220-3202333010211121-3031111112212132-2012111213221301-2112320332101011-2231213112111021-0020221033003003) |
| `f5_big_ip_aws_service.endpoint_service.advertise_on_slo_ip_external` | [f5_big_ip_aws_service.endpoint_service.advertise_on_slo_ip_external](resources--nfv_service--reference--group-001.md#canonical-3310321211232133-1223311010310302-1003310022203022-1332002300203302-3330232100131223-2130132201220301-3203202323230203-2123102210133232) |
| `f5_big_ip_aws_service.endpoint_service.automatic_vip` | [f5_big_ip_aws_service.endpoint_service.automatic_vip](resources--nfv_service--reference--group-001.md#canonical-0312211323110033-3303332203130031-0233200011011313-0311113300222102-1132010013122020-3231022320200232-0020003021321312-3301231312132111) |
| `f5_big_ip_aws_service.endpoint_service.configured_vip` | [f5_big_ip_aws_service.endpoint_service.configured_vip](resources--nfv_service--reference--group-001.md#canonical-2002310120303123-0122223120010230-1331303022213130-1111000013013311-1231112322001123-2300220000333111-1011331130031102-1323301311233223) |
| `f5_big_ip_aws_service.endpoint_service.custom_tcp_ports` | [f5_big_ip_aws_service.endpoint_service.custom_tcp_ports](resources--nfv_service--reference--group-001.md#canonical-1220013311203222-3322132211133220-3222233221333103-3122132300301213-0102113223033230-2303330131202101-2222312332222110-3312331211221202) |
| `f5_big_ip_aws_service.endpoint_service.custom_tcp_ports.ports` | [f5_big_ip_aws_service.endpoint_service.custom_tcp_ports.ports](resources--nfv_service--reference--group-001.md#canonical-0000120102002302-2201133012233123-1230220122301302-1230003031013320-0111023310200223-1101200332310333-2031202100131312-3000113233001212) |
| `f5_big_ip_aws_service.endpoint_service.custom_udp_ports` | [f5_big_ip_aws_service.endpoint_service.custom_udp_ports](resources--nfv_service--reference--group-001.md#canonical-2300333333323321-1030021333102011-1003031021303020-1200323131003122-3011323220310103-2031223201111232-1330110013332133-3120331222011303) |
| `f5_big_ip_aws_service.endpoint_service.custom_udp_ports.ports` | [f5_big_ip_aws_service.endpoint_service.custom_udp_ports.ports](resources--nfv_service--reference--group-001.md#canonical-1012021101223321-3223310022000022-3032110012001230-2333231232021122-3331213302133102-2302311202031300-0213103303313231-1330023100222032) |
| `f5_big_ip_aws_service.endpoint_service.default_tcp_ports` | [f5_big_ip_aws_service.endpoint_service.default_tcp_ports](resources--nfv_service--reference--group-001.md#canonical-2030023210312113-2001323312121333-0012220033200223-2333130023110002-3232001113122221-1121001322301002-2030311213033313-0221213301302220) |
| `f5_big_ip_aws_service.endpoint_service.disable_advertise_on_slo_ip` | [f5_big_ip_aws_service.endpoint_service.disable_advertise_on_slo_ip](resources--nfv_service--reference--group-001.md#canonical-3102223111322131-1130121300023132-2000322232100321-1131123113213121-0302333021013002-0012023130133023-3321221211133301-3101120031020230) |
| `f5_big_ip_aws_service.endpoint_service.http_port` | [f5_big_ip_aws_service.endpoint_service.http_port](resources--nfv_service--reference--group-001.md#canonical-0122010213002333-2131130211001233-2002311002230223-3322212201210133-2002012110231301-1023101000231223-2331131302102202-0230022101212003) |
| `f5_big_ip_aws_service.endpoint_service.https_port` | [f5_big_ip_aws_service.endpoint_service.https_port](resources--nfv_service--reference--group-002.md#canonical-3213200002230332-1112000033222310-2121220020110031-0313333130213222-1320131011223201-1113022013000112-2032221032220100-2322323031100132) |
| `f5_big_ip_aws_service.endpoint_service.no_tcp_ports` | [f5_big_ip_aws_service.endpoint_service.no_tcp_ports](resources--nfv_service--reference--group-002.md#canonical-0031203113313320-2000021332233113-0133122032103002-0021122020111311-3132031120321221-2031132223103012-0313332012111212-3032122211300012) |
| `f5_big_ip_aws_service.endpoint_service.no_udp_ports` | [f5_big_ip_aws_service.endpoint_service.no_udp_ports](resources--nfv_service--reference--group-002.md#canonical-2302111331330122-3120022132102213-2012020331132031-3101313032303321-3300210000112122-1320123212313210-0132221032312033-3012022030210131) |
| `f5_big_ip_aws_service.market_place_image` | [f5_big_ip_aws_service.market_place_image](resources--nfv_service--reference--group-002.md#canonical-0130301203100302-3320011222113103-2233202312103233-1121212120332021-0100201102103322-1132210222233310-0300113210112121-3300112100033101) |
| `f5_big_ip_aws_service.market_place_image.awafpay_g200_mbps` | [f5_big_ip_aws_service.market_place_image.awafpay_g200_mbps](resources--nfv_service--reference--group-002.md#canonical-3013002010302031-0223121012302313-0203311033210231-0012031111210212-1332120130131022-3313200110102232-1112000311101220-1320102333212013) |
| `f5_big_ip_aws_service.market_place_image.awafpay_g3_gbps` | [f5_big_ip_aws_service.market_place_image.awafpay_g3_gbps](resources--nfv_service--reference--group-002.md#canonical-0313103233111113-1301101000002302-0010330233101313-3303031321330022-0101012032023131-0103211222203211-1110103303020001-3311121303001011) |
| `f5_big_ip_aws_service.nodes` | [f5_big_ip_aws_service.nodes](resources--nfv_service--reference--group-002.md#canonical-3322312331030022-2323331220311230-3322211311333031-1010031133322100-2101112112010002-0032031213032211-3121323132132102-3000001312210133) |
| `f5_big_ip_aws_service.nodes.automatic_prefix` | [f5_big_ip_aws_service.nodes.automatic_prefix](resources--nfv_service--reference--group-002.md#canonical-0321222210300212-1313321302213310-3212131023020201-1313312203001232-0200011011002330-1023311023023001-0020023023022302-3030031332331323) |
| `f5_big_ip_aws_service.nodes.aws_az_name` | [f5_big_ip_aws_service.nodes.aws_az_name](resources--nfv_service--reference--group-002.md#canonical-2130312031100011-1223332112201300-2232111031010121-0102002030201230-3022331220301002-3230220001312022-0333100301020311-0321233203103222) |
| `f5_big_ip_aws_service.nodes.mgmt_subnet` | [f5_big_ip_aws_service.nodes.mgmt_subnet](resources--nfv_service--reference--group-002.md#canonical-0122002030110001-1122012022001111-3131220223022030-3231330121100020-3323322303220102-3000332012003230-1021003002122212-2221213011232103) |
| `f5_big_ip_aws_service.nodes.mgmt_subnet.existing_subnet_id` | [f5_big_ip_aws_service.nodes.mgmt_subnet.existing_subnet_id](resources--nfv_service--reference--group-002.md#canonical-3211101130211131-0011331031322321-1102023110023331-2320001321303010-1322301000133102-3222220110020013-2202030032131113-2320311021011230) |
| `f5_big_ip_aws_service.nodes.mgmt_subnet.subnet_param` | [f5_big_ip_aws_service.nodes.mgmt_subnet.subnet_param](resources--nfv_service--reference--group-002.md#canonical-0113312113022013-3003211103033221-1310021323023311-2200310023003223-2012100212032013-2112223103022220-0232031203002222-1203331022030101) |
| `f5_big_ip_aws_service.nodes.mgmt_subnet.subnet_param.ipv4` | [f5_big_ip_aws_service.nodes.mgmt_subnet.subnet_param.ipv4](resources--nfv_service--reference--group-002.md#canonical-0331210223320322-1033003000220102-1110333302210302-1230232212022133-1302201122121312-0112330333310301-2323120322002131-1103222113112102) |
| `f5_big_ip_aws_service.nodes.node_name` | [f5_big_ip_aws_service.nodes.node_name](resources--nfv_service--reference--group-002.md#canonical-3023030020220322-0212220021110223-3203030132120230-0101321102012132-1013312033222321-3010320102301120-0233313212112212-1323233123113323) |
| `f5_big_ip_aws_service.nodes.reserved_mgmt_subnet` | [f5_big_ip_aws_service.nodes.reserved_mgmt_subnet](resources--nfv_service--reference--group-002.md#canonical-0311323132002111-2221023212113103-2131311303120213-3121302331113023-3333100003322230-2013133111321133-0132131002331333-1010312031103321) |
| `f5_big_ip_aws_service.nodes.tunnel_prefix` | [f5_big_ip_aws_service.nodes.tunnel_prefix](resources--nfv_service--reference--group-002.md#canonical-1003211130213231-3220103110102213-1102323030303311-2220213012301221-2230111303001031-3030120113021321-2130203330023333-1333212311302120) |
| `f5_big_ip_aws_service.ssh_key` | [f5_big_ip_aws_service.ssh_key](resources--nfv_service--reference--group-001.md#canonical-2122322100020010-0233231112003322-1100330013212030-1031310200022000-0023300112302121-3230133231122000-3021020020303103-3002313323100011) |
| `f5_big_ip_aws_service.tags` | [f5_big_ip_aws_service.tags](resources--nfv_service--reference--group-001.md#canonical-2300122213223130-2322121232230202-1121102110030323-0233022121232321-3120330012222113-2201303223211330-2200201322003010-1322232110123021) |
| `https_management` | [https_management](resources--nfv_service--reference--group-002.md#canonical-2310110120210321-0011122021130100-1020101231210333-1001101201002320-3213323313112230-0312000102310210-1113012113122222-1223202310021130) |
| `https_management.advertise_on_internet` | [https_management.advertise_on_internet](resources--nfv_service--reference--group-002.md#canonical-0323010000210110-1011321220101213-3302313013003100-3101320312130203-3020110132002010-2121301212013103-3112021220132330-0311003300231200) |
| `https_management.advertise_on_internet.public_ip` | [https_management.advertise_on_internet.public_ip](resources--nfv_service--reference--group-002.md#canonical-1320022302320301-0301323112111020-0012221032212033-1121313221310031-3111210222023102-1211122200200303-0212322333301221-2200310112030020) |
| `https_management.advertise_on_internet.public_ip.name` | [https_management.advertise_on_internet.public_ip.name](resources--nfv_service--reference--group-002.md#canonical-2102211110131331-0031231300331000-3223132133302211-0121311233011100-2333023102113303-0023202021210101-0203210322121200-3020031011101323) |
| `https_management.advertise_on_internet.public_ip.namespace` | [https_management.advertise_on_internet.public_ip.namespace](resources--nfv_service--reference--group-002.md#canonical-3103012222010303-3300122000000203-0133133311121300-1133132131331110-3310310002212023-2330130111330012-2221002133313223-2001011220220002) |
| `https_management.advertise_on_internet.public_ip.tenant` | [https_management.advertise_on_internet.public_ip.tenant](resources--nfv_service--reference--group-002.md#canonical-3011122031333123-0010210111031230-1200100331102003-3323102101001103-1202110111301112-0220022012121232-1031103102221320-3332230022330221) |
| `https_management.advertise_on_internet_default_vip` | [https_management.advertise_on_internet_default_vip](resources--nfv_service--reference--group-002.md#canonical-3032302201112032-1310300120032303-3102310202001023-3132002203013332-0313211203133322-2230000031320222-0030000330311033-3330122112223211) |
| `https_management.advertise_on_sli_vip` | [https_management.advertise_on_sli_vip](resources--nfv_service--reference--group-002.md#canonical-0333202002211320-0113013211010133-1020230123220300-1113103313110221-3232301101001323-3011223320210101-2332032320313031-2300121003201333) |
| `https_management.advertise_on_sli_vip.no_mtls` | [https_management.advertise_on_sli_vip.no_mtls](resources--nfv_service--reference--group-002.md#canonical-1230321233200133-1223000200333331-0303021322210113-3221332003310323-3002313011001201-2232201012122103-0023230322332200-1133302133100111) |
| `https_management.advertise_on_sli_vip.tls_certificates` | [https_management.advertise_on_sli_vip.tls_certificates](resources--nfv_service--reference--group-002.md#canonical-0331213303111210-0202201333130133-3210033310113312-1210101222131032-1311310100111103-0033303013311223-3222200011211033-1110223202033333) |
| `https_management.advertise_on_sli_vip.tls_certificates.blindfold` | [https_management.advertise_on_sli_vip.tls_certificates.blindfold](resources--nfv_service--reference--group-002.md#canonical-1303300220233002-1131012132323222-3222301132332323-1011222111211021-1101311312302300-1211303110223131-1121123110230203-2320300212120102) |
| `https_management.advertise_on_sli_vip.tls_certificates.blindfold.algorithm` | [https_management.advertise_on_sli_vip.tls_certificates.blindfold.algorithm](resources--nfv_service--reference--group-002.md#canonical-1123031101302320-1221333022220333-3000322031113101-0320321323220001-0130101121210232-1311133100001021-0322200112221012-1101310303010100) |
| `https_management.advertise_on_sli_vip.tls_certificates.blindfold.certificate_file` | [https_management.advertise_on_sli_vip.tls_certificates.blindfold.certificate_file](resources--nfv_service--reference--group-002.md#canonical-1212203010203300-2133200123122103-1031020211102303-3202010010113111-0032302012033012-3013300100101210-3303230122203131-3203120303012231) |
| `https_management.advertise_on_sli_vip.tls_certificates.blindfold.certificate_pem` | [https_management.advertise_on_sli_vip.tls_certificates.blindfold.certificate_pem](resources--nfv_service--reference--group-002.md#canonical-3112010001002020-3331131123031302-0132100233020230-3210023103320112-0003223210021210-0111202320111223-1012331130120200-0013022300101321) |
| `https_management.advertise_on_sli_vip.tls_certificates.blindfold.chain_identity` | [https_management.advertise_on_sli_vip.tls_certificates.blindfold.chain_identity](resources--nfv_service--reference--group-002.md#canonical-2312233030201011-2121310220333231-3223030030320311-3321220100121313-2312320332121231-0123123303221121-0300302000120121-0102213120132110) |
| `https_management.advertise_on_sli_vip.tls_certificates.blindfold.context_digest` | [https_management.advertise_on_sli_vip.tls_certificates.blindfold.context_digest](resources--nfv_service--reference--group-002.md#canonical-0201300010123120-2323002011133311-0311220002101003-1211100110130111-0233120033231023-1101213330232100-2031113323121022-0310121200012133) |
| `https_management.advertise_on_sli_vip.tls_certificates.blindfold.encrypted_location` | [https_management.advertise_on_sli_vip.tls_certificates.blindfold.encrypted_location](resources--nfv_service--reference--group-002.md#canonical-1331330213303310-2020023023012032-1012132231003333-3020220230122002-1132231121212223-0332232131011321-3132111011031221-2030210330330311) |
| `https_management.advertise_on_sli_vip.tls_certificates.blindfold.expires_at` | [https_management.advertise_on_sli_vip.tls_certificates.blindfold.expires_at](resources--nfv_service--reference--group-002.md#canonical-2211101020303230-2310123211300310-2020030001321120-3013321011211002-3200000332213312-3202221232220200-3011332002111230-1121300223321111) |
| `https_management.advertise_on_sli_vip.tls_certificates.blindfold.fingerprint` | [https_management.advertise_on_sli_vip.tls_certificates.blindfold.fingerprint](resources--nfv_service--reference--group-002.md#canonical-2020301202301220-2232032021002312-1021010301113323-2020312001230131-3330111333231323-3230013323132111-1122222231122333-2010010022001212) |
| `https_management.advertise_on_sli_vip.tls_certificates.blindfold.id` | [https_management.advertise_on_sli_vip.tls_certificates.blindfold.id](resources--nfv_service--reference--group-002.md#canonical-3013300013130212-3222111011100303-0100033003112300-0110120011000130-2122333023311023-2030022013321213-0302023212232202-3223011031322231) |
| `https_management.advertise_on_sli_vip.tls_certificates.blindfold.material_version` | [https_management.advertise_on_sli_vip.tls_certificates.blindfold.material_version](resources--nfv_service--reference--group-002.md#canonical-3300012003101122-3131202021123310-3210330001103132-3023021201213312-1110002330120302-2122303100100311-0113130321010033-0312020100121011) |
| `https_management.advertise_on_sli_vip.tls_certificates.blindfold.passphrase_env` | [https_management.advertise_on_sli_vip.tls_certificates.blindfold.passphrase_env](resources--nfv_service--reference--group-002.md#canonical-3333033000303201-3212113233233301-3031323021112031-3231230300312331-0320323133333120-3103301333110233-2223002031303223-2300111301323323) |
| `https_management.advertise_on_sli_vip.tls_certificates.blindfold.passphrase_wo` | [https_management.advertise_on_sli_vip.tls_certificates.blindfold.passphrase_wo](resources--nfv_service--reference--group-002.md#canonical-3100001302020122-3201202012222122-0310232031002232-1322011010121223-0311310233120303-3111333311013201-2130331011012311-3333322300232121) |
| `https_management.advertise_on_sli_vip.tls_certificates.blindfold.pkcs12_file` | [https_management.advertise_on_sli_vip.tls_certificates.blindfold.pkcs12_file](resources--nfv_service--reference--group-002.md#canonical-3312233133030201-0121303331003213-1212013133302112-0030123213201010-0123010110131112-3232122030102003-3012023011001020-0131101121331000) |
| `https_management.advertise_on_sli_vip.tls_certificates.blindfold.pkcs12_wo` | [https_management.advertise_on_sli_vip.tls_certificates.blindfold.pkcs12_wo](resources--nfv_service--reference--group-002.md#canonical-3323113231233001-3120111020232322-2210121333123321-3112200211210102-3312330210302201-0122200213333330-2311111132121113-3300010210121123) |
| `https_management.advertise_on_sli_vip.tls_certificates.blindfold.policy` | [https_management.advertise_on_sli_vip.tls_certificates.blindfold.policy](resources--nfv_service--reference--group-002.md#canonical-1302122312221200-1030112220032002-3001203210303030-0113100312203112-0300303333230310-0131112001300021-1210230301231131-0022102130301312) |
| `https_management.advertise_on_sli_vip.tls_certificates.blindfold.prepared_identity` | [https_management.advertise_on_sli_vip.tls_certificates.blindfold.prepared_identity](resources--nfv_service--reference--group-002.md#canonical-3003203231200120-1102013011310230-0232221132333203-1221033233230021-1000322213033223-1232021002222211-0201200120230131-3300122303020022) |
| `https_management.advertise_on_sli_vip.tls_certificates.blindfold.private_key_file` | [https_management.advertise_on_sli_vip.tls_certificates.blindfold.private_key_file](resources--nfv_service--reference--group-002.md#canonical-0120320112003021-2312220113332022-1221112322231223-2203123021122330-2310312021221012-0231200331233011-0110311031211220-0113220211011233) |
| `https_management.advertise_on_sli_vip.tls_certificates.blindfold.private_key_wo` | [https_management.advertise_on_sli_vip.tls_certificates.blindfold.private_key_wo](resources--nfv_service--reference--group-002.md#canonical-0223313202133300-1130002210212312-3110231103133333-0331201113103223-3302320322311201-3010211332000003-2321023333021133-2112121322222023) |
| `https_management.advertise_on_sli_vip.tls_certificates.blindfold.spki_identity` | [https_management.advertise_on_sli_vip.tls_certificates.blindfold.spki_identity](resources--nfv_service--reference--group-002.md#canonical-1032120120003223-2021023232123333-1113302123321230-3102302323101010-2121133121003312-0211221132310322-1230101110321213-0302031233312332) |
| `https_management.advertise_on_sli_vip.tls_certificates.certificate_url` | [https_management.advertise_on_sli_vip.tls_certificates.certificate_url](resources--nfv_service--reference--group-002.md#canonical-2301110001001230-3133131001202021-2012133312201011-3202313010100031-2321133201333031-1032300322103120-2231320233131020-1103033230311302) |
| `https_management.advertise_on_sli_vip.tls_certificates.custom_hash_algorithms` | [https_management.advertise_on_sli_vip.tls_certificates.custom_hash_algorithms](resources--nfv_service--reference--group-002.md#canonical-0232220321113212-1020220210130310-2101311322011200-1200222003133021-0323332113113120-2101113331022333-1233333012312012-1131131310000331) |
| `https_management.advertise_on_sli_vip.tls_certificates.custom_hash_algorithms.hash_algorithms` | [https_management.advertise_on_sli_vip.tls_certificates.custom_hash_algorithms.hash_algorithms](resources--nfv_service--reference--group-002.md#canonical-0102031113100312-1201113023103131-1021123223001021-1022111213330321-2132023302032030-3110311100132010-3122001011002320-0101010303013200) |
| `https_management.advertise_on_sli_vip.tls_certificates.description_spec` | [https_management.advertise_on_sli_vip.tls_certificates.description_spec](resources--nfv_service--reference--group-002.md#canonical-1303031230302323-1022011313213011-2202221313113320-1131013120030112-1121002320213221-0203212213131213-1000231113000221-0110322012013330) |
| `https_management.advertise_on_sli_vip.tls_certificates.disable_ocsp_stapling` | [https_management.advertise_on_sli_vip.tls_certificates.disable_ocsp_stapling](resources--nfv_service--reference--group-002.md#canonical-1223030033333023-2212300113033211-2012300202013332-0000023021211322-1000223311233201-2202311002301120-3312110323120203-0113220323012103) |
| `https_management.advertise_on_sli_vip.tls_certificates.private_key` | [https_management.advertise_on_sli_vip.tls_certificates.private_key](resources--nfv_service--reference--group-002.md#canonical-3302320220131210-1102233221122301-3323021113033023-2313133210200301-2230213022012212-1211032023021122-0202230200132102-1133030113012220) |
| `https_management.advertise_on_sli_vip.tls_certificates.private_key.blindfold_secret_info` | [https_management.advertise_on_sli_vip.tls_certificates.private_key.blindfold_secret_info](resources--nfv_service--reference--group-002.md#canonical-1332203002310312-2112111220132213-3113322303220211-1211011131010323-2331322232132012-1020031202231103-1333331221023222-1233212301001212) |
| `https_management.advertise_on_sli_vip.tls_certificates.private_key.blindfold_secret_info.decryption_provider` | [https_management.advertise_on_sli_vip.tls_certificates.private_key.blindfold_secret_info.decryption_provider](resources--nfv_service--reference--group-002.md#canonical-3303210012011132-1012133030203330-3201232230112320-1033013101132320-2312130012133121-1303012201232310-2332213330012203-2312001312330112) |
| `https_management.advertise_on_sli_vip.tls_certificates.private_key.blindfold_secret_info.location` | [https_management.advertise_on_sli_vip.tls_certificates.private_key.blindfold_secret_info.location](resources--nfv_service--reference--group-002.md#canonical-3001023112033133-0320001100221300-0102131012010031-0000011300311103-0312002201113223-0313301113300100-1313202232213002-3103202203223322) |
| `https_management.advertise_on_sli_vip.tls_certificates.private_key.blindfold_secret_info.store_provider` | [https_management.advertise_on_sli_vip.tls_certificates.private_key.blindfold_secret_info.store_provider](resources--nfv_service--reference--group-002.md#canonical-2133211103020102-2113033001211121-0000012303323311-2221000022131031-3122111311021330-1013112020321133-0322200223310102-0013032033131110) |
| `https_management.advertise_on_sli_vip.tls_certificates.private_key.clear_secret_info` | [https_management.advertise_on_sli_vip.tls_certificates.private_key.clear_secret_info](resources--nfv_service--reference--group-002.md#canonical-1131000321321120-0221313001131013-3313300132223330-2222232210323200-0332202113102203-2200231220233332-1233323310011200-0201332002303323) |
| `https_management.advertise_on_sli_vip.tls_certificates.private_key.clear_secret_info.provider_ref` | [https_management.advertise_on_sli_vip.tls_certificates.private_key.clear_secret_info.provider_ref](resources--nfv_service--reference--group-002.md#canonical-1123121113332012-3130121333322311-1001312002023222-1020120202312022-2303021003313330-3030103000100132-2133320131231030-1310122320112212) |
| `https_management.advertise_on_sli_vip.tls_certificates.private_key.clear_secret_info.url` | [https_management.advertise_on_sli_vip.tls_certificates.private_key.clear_secret_info.url](resources--nfv_service--reference--group-002.md#canonical-3302320120113221-3103312211130001-2302330022300031-1122311103011221-1322100100210001-2232300002112100-1201102211103033-0230100330211022) |
| `https_management.advertise_on_sli_vip.tls_certificates.use_system_defaults` | [https_management.advertise_on_sli_vip.tls_certificates.use_system_defaults](resources--nfv_service--reference--group-002.md#canonical-0201311301000121-3312110021222313-1301033213133020-3010301201332132-0033221323010003-1213113100121213-3001333030330003-1111221031030202) |
| `https_management.advertise_on_sli_vip.tls_config` | [https_management.advertise_on_sli_vip.tls_config](resources--nfv_service--reference--group-002.md#canonical-3101000013132111-3202323103212010-1001233331013303-3230002322121133-3301310201123211-2102000030001330-3100001313231200-0003323130322332) |
| `https_management.advertise_on_sli_vip.tls_config.custom_security` | [https_management.advertise_on_sli_vip.tls_config.custom_security](resources--nfv_service--reference--group-002.md#canonical-0321123112033012-2121202310132012-2011122333133321-2210212033120201-2002112020211321-2111312111101033-1120323130022112-0233200301202213) |
| `https_management.advertise_on_sli_vip.tls_config.custom_security.cipher_suites` | [https_management.advertise_on_sli_vip.tls_config.custom_security.cipher_suites](resources--nfv_service--reference--group-002.md#canonical-3022310133313110-1113102320331031-2202003331102222-3322301100021323-3300300321221222-1003310122032203-3022133011222023-3021002110111010) |
| `https_management.advertise_on_sli_vip.tls_config.custom_security.max_version` | [https_management.advertise_on_sli_vip.tls_config.custom_security.max_version](resources--nfv_service--reference--group-002.md#canonical-1031232122113113-0010221220022222-2222133121300203-2320332202221003-0330113200121011-0131211131012000-1223001032103312-1311111112101310) |
| `https_management.advertise_on_sli_vip.tls_config.custom_security.min_version` | [https_management.advertise_on_sli_vip.tls_config.custom_security.min_version](resources--nfv_service--reference--group-002.md#canonical-3331000101231010-2100023302312123-1232222112211113-0220313020301313-3231111221101133-2322201232211302-0332223231202132-3333332231000232) |
| `https_management.advertise_on_sli_vip.tls_config.default_security` | [https_management.advertise_on_sli_vip.tls_config.default_security](resources--nfv_service--reference--group-002.md#canonical-3002102120320012-2130113211330123-1132122201320022-2121031223310330-1312301322010303-0220133001332221-1032300100302202-3132123001202102) |
| `https_management.advertise_on_sli_vip.tls_config.low_security` | [https_management.advertise_on_sli_vip.tls_config.low_security](resources--nfv_service--reference--group-002.md#canonical-2300130103020101-0200032311313302-2303201020000031-1232220323121123-1030233112213023-1203210323131232-0210100112330022-2032323021300330) |
| `https_management.advertise_on_sli_vip.tls_config.medium_security` | [https_management.advertise_on_sli_vip.tls_config.medium_security](resources--nfv_service--reference--group-002.md#canonical-1012212230321201-0303230111112322-3310112130310203-2212331312130223-3030301113102011-0230021000211310-3203022221320232-1002130202211231) |
| `https_management.advertise_on_sli_vip.use_mtls` | [https_management.advertise_on_sli_vip.use_mtls](resources--nfv_service--reference--group-002.md#canonical-2300110110310032-3001312301300010-0220223223002120-2220311003211122-0201213123321010-2012013311220221-3130210022232112-3220022022002030) |
| `https_management.advertise_on_sli_vip.use_mtls.client_certificate_optional` | [https_management.advertise_on_sli_vip.use_mtls.client_certificate_optional](resources--nfv_service--reference--group-002.md#canonical-1322023323211101-1100203012133221-0323331031021301-1330331231011000-1103213120011303-3302113130213332-0133302332333000-2210220033311003) |
| `https_management.advertise_on_sli_vip.use_mtls.crl` | [https_management.advertise_on_sli_vip.use_mtls.crl](resources--nfv_service--reference--group-002.md#canonical-1020331103213331-0233022022203030-0202022312321030-1133313320111002-1200013103320121-0130231303033122-1030121323110203-2202321202313300) |
| `https_management.advertise_on_sli_vip.use_mtls.crl.name` | [https_management.advertise_on_sli_vip.use_mtls.crl.name](resources--nfv_service--reference--group-002.md#canonical-3213010103333231-3232320133311310-0332230101021222-1102010101232320-1023013110001102-0301012302320330-3230030211002332-0023230312232230) |
| `https_management.advertise_on_sli_vip.use_mtls.crl.namespace` | [https_management.advertise_on_sli_vip.use_mtls.crl.namespace](resources--nfv_service--reference--group-002.md#canonical-2213302030021311-1112110132123223-0020202012000013-1212030030013001-1100201030123303-3013030213233323-1232331312233223-1212033213133001) |
| `https_management.advertise_on_sli_vip.use_mtls.crl.tenant` | [https_management.advertise_on_sli_vip.use_mtls.crl.tenant](resources--nfv_service--reference--group-002.md#canonical-1320110020112330-0313222123311012-2332130013320230-3332222133103132-0000231213121120-3003032112201211-1120110100212231-2303121201030222) |
| `https_management.advertise_on_sli_vip.use_mtls.no_crl` | [https_management.advertise_on_sli_vip.use_mtls.no_crl](resources--nfv_service--reference--group-002.md#canonical-2113233101222310-0101313022111223-3230332230121203-0122131320033132-0002021212122031-2230231302302000-1131220203130012-3013022133320030) |
| `https_management.advertise_on_sli_vip.use_mtls.trusted_ca` | [https_management.advertise_on_sli_vip.use_mtls.trusted_ca](resources--nfv_service--reference--group-002.md#canonical-1111211213203211-1302301322220020-2200133021020022-1210133322213212-0011111112110122-0013312203211022-1030212232312303-2233202121121013) |
| `https_management.advertise_on_sli_vip.use_mtls.trusted_ca.name` | [https_management.advertise_on_sli_vip.use_mtls.trusted_ca.name](resources--nfv_service--reference--group-002.md#canonical-3220330222000010-2322231000130012-2120210203132323-3002311221103312-0303233302011123-1011302130200101-3322220311133111-3221222313203033) |
| `https_management.advertise_on_sli_vip.use_mtls.trusted_ca.namespace` | [https_management.advertise_on_sli_vip.use_mtls.trusted_ca.namespace](resources--nfv_service--reference--group-002.md#canonical-0231220213312112-1000033100000330-3301120321210020-2111213031221313-3221222003302212-0012003010200020-2313000212221120-0113313201100110) |
| `https_management.advertise_on_sli_vip.use_mtls.trusted_ca.tenant` | [https_management.advertise_on_sli_vip.use_mtls.trusted_ca.tenant](resources--nfv_service--reference--group-002.md#canonical-1021220333032132-0022020013030201-1113330122312210-2010201120011321-1112010032331213-3213130121101331-2212301101232030-3003211100330032) |
| `https_management.advertise_on_sli_vip.use_mtls.trusted_ca_url` | [https_management.advertise_on_sli_vip.use_mtls.trusted_ca_url](resources--nfv_service--reference--group-002.md#canonical-0210022220123102-0321023003033212-0210023121031321-2030130302120310-2003303032230020-2223220310020112-2211221101132020-0021213001232002) |
| `https_management.advertise_on_sli_vip.use_mtls.xfcc_disabled` | [https_management.advertise_on_sli_vip.use_mtls.xfcc_disabled](resources--nfv_service--reference--group-002.md#canonical-0330112001111121-1131333312000321-0312020111233011-3320030120332320-1002330110202201-1333133221211210-2032220223133333-3030220122311023) |
| `https_management.advertise_on_sli_vip.use_mtls.xfcc_options` | [https_management.advertise_on_sli_vip.use_mtls.xfcc_options](resources--nfv_service--reference--group-002.md#canonical-3323110102212011-0212001321113331-0130003030101321-2113010100030032-0202102210021130-1000113333321302-1300010131123033-0012331111100223) |
| `https_management.advertise_on_sli_vip.use_mtls.xfcc_options.xfcc_header_elements` | [https_management.advertise_on_sli_vip.use_mtls.xfcc_options.xfcc_header_elements](resources--nfv_service--reference--group-002.md#canonical-3312211122310110-2202300001230110-2020133131220230-1111322322011320-1203030130332031-1012110202311110-1132202133031233-0213001220123000) |
| `https_management.advertise_on_slo_internet_vip` | [https_management.advertise_on_slo_internet_vip](resources--nfv_service--reference--group-002.md#canonical-0000003311331222-2311113333311103-2302132333030302-2232121322202323-0000303312213300-1103320133323031-0001022212013022-3312312223201311) |
| `https_management.advertise_on_slo_internet_vip.no_mtls` | [https_management.advertise_on_slo_internet_vip.no_mtls](resources--nfv_service--reference--group-002.md#canonical-3211330323003300-0001210121130010-2300102110102023-0232132200231300-1213310012120311-2032311322110110-3221130220001220-1303022023222221) |
| `https_management.advertise_on_slo_internet_vip.tls_certificates` | [https_management.advertise_on_slo_internet_vip.tls_certificates](resources--nfv_service--reference--group-002.md#canonical-0333211320302231-1312120110133003-2220203102020210-0120232121231203-2233023003030033-3211033022122212-0310100101221103-3123121320130033) |
| `https_management.advertise_on_slo_internet_vip.tls_certificates.blindfold` | [https_management.advertise_on_slo_internet_vip.tls_certificates.blindfold](resources--nfv_service--reference--group-002.md#canonical-3012113311001130-0301123020321322-0001223331230302-1021121213311020-1320323132331331-2030020213313321-1323320030322321-3011101213130313) |
| `https_management.advertise_on_slo_internet_vip.tls_certificates.blindfold.algorithm` | [https_management.advertise_on_slo_internet_vip.tls_certificates.blindfold.algorithm](resources--nfv_service--reference--group-003.md#canonical-1202110222032323-0301002212312331-0230200112023003-2211303003000001-0223202022003111-1001112330201031-1203322213301313-1321222003302031) |
| `https_management.advertise_on_slo_internet_vip.tls_certificates.blindfold.certificate_file` | [https_management.advertise_on_slo_internet_vip.tls_certificates.blindfold.certificate_file](resources--nfv_service--reference--group-003.md#canonical-0321020010033233-0033030103300121-2230112331332121-3010322331010101-0222312232213231-3120223120222032-0312321001003113-3213221113311311) |
| `https_management.advertise_on_slo_internet_vip.tls_certificates.blindfold.certificate_pem` | [https_management.advertise_on_slo_internet_vip.tls_certificates.blindfold.certificate_pem](resources--nfv_service--reference--group-003.md#canonical-3221021212110211-3101131333000313-2231322101020110-3101233230301031-2303021021131300-0021310121132232-3132000021013203-2330133121211111) |
| `https_management.advertise_on_slo_internet_vip.tls_certificates.blindfold.chain_identity` | [https_management.advertise_on_slo_internet_vip.tls_certificates.blindfold.chain_identity](resources--nfv_service--reference--group-003.md#canonical-2313302203212111-2033220201302130-2302113203231301-0202331311102310-2123312010200331-0222131223213312-0100103303023003-3003330022112113) |
| `https_management.advertise_on_slo_internet_vip.tls_certificates.blindfold.context_digest` | [https_management.advertise_on_slo_internet_vip.tls_certificates.blindfold.context_digest](resources--nfv_service--reference--group-003.md#canonical-3322211121220302-0101233212230202-1122103030210310-3111300133010330-3213222012300303-3200122002233332-0233122332110032-2222220201011123) |
| `https_management.advertise_on_slo_internet_vip.tls_certificates.blindfold.encrypted_location` | [https_management.advertise_on_slo_internet_vip.tls_certificates.blindfold.encrypted_location](resources--nfv_service--reference--group-003.md#canonical-3323223132113111-1303003121020332-1303011010121013-3022231122332013-3022201102012123-3131300123133312-1020121013103123-2132123000113330) |
| `https_management.advertise_on_slo_internet_vip.tls_certificates.blindfold.expires_at` | [https_management.advertise_on_slo_internet_vip.tls_certificates.blindfold.expires_at](resources--nfv_service--reference--group-003.md#canonical-1010111202003330-2333301332020033-3312313030003322-0201113100121220-2122033120110020-2320133112130231-3010332232221222-3120223321122222) |
| `https_management.advertise_on_slo_internet_vip.tls_certificates.blindfold.fingerprint` | [https_management.advertise_on_slo_internet_vip.tls_certificates.blindfold.fingerprint](resources--nfv_service--reference--group-003.md#canonical-3003002200133231-1321212020213313-2032333020300000-3001021030212020-2022031011320130-0303032203023103-1113211332031000-1100202312002201) |
| `https_management.advertise_on_slo_internet_vip.tls_certificates.blindfold.id` | [https_management.advertise_on_slo_internet_vip.tls_certificates.blindfold.id](resources--nfv_service--reference--group-003.md#canonical-1323331022202323-3210123231301022-1100120101102311-3310233101322001-3230233333112022-3203132033100222-2021230212212022-3100301111230321) |
| `https_management.advertise_on_slo_internet_vip.tls_certificates.blindfold.material_version` | [https_management.advertise_on_slo_internet_vip.tls_certificates.blindfold.material_version](resources--nfv_service--reference--group-003.md#canonical-3321220030032103-0313101020213102-1213132121321323-0312303011112132-1221302212000132-0220110023203332-0021200100030202-2310013132210320) |
| `https_management.advertise_on_slo_internet_vip.tls_certificates.blindfold.passphrase_env` | [https_management.advertise_on_slo_internet_vip.tls_certificates.blindfold.passphrase_env](resources--nfv_service--reference--group-003.md#canonical-1223301203112010-2022033000101303-0122200102321302-3131122232223322-0002003323211212-2232133233001333-1312133111212111-2311211022001032) |
| `https_management.advertise_on_slo_internet_vip.tls_certificates.blindfold.passphrase_wo` | [https_management.advertise_on_slo_internet_vip.tls_certificates.blindfold.passphrase_wo](resources--nfv_service--reference--group-003.md#canonical-1101102331021131-3021001323231023-3202232303120303-3222301111323100-3333133023023223-3312102221113313-1023332233330031-2122203322111013) |
| `https_management.advertise_on_slo_internet_vip.tls_certificates.blindfold.pkcs12_file` | [https_management.advertise_on_slo_internet_vip.tls_certificates.blindfold.pkcs12_file](resources--nfv_service--reference--group-003.md#canonical-2321123331131203-2003033011320310-3321232000330230-0021311112231013-3202303032122201-0330000222010303-0230221333113212-0102311333202032) |
| `https_management.advertise_on_slo_internet_vip.tls_certificates.blindfold.pkcs12_wo` | [https_management.advertise_on_slo_internet_vip.tls_certificates.blindfold.pkcs12_wo](resources--nfv_service--reference--group-003.md#canonical-1313220033332300-3031131311211111-3323131002022131-2231321003301032-2123123323223322-3122202020010322-0312030213300032-2001021303331333) |
| `https_management.advertise_on_slo_internet_vip.tls_certificates.blindfold.policy` | [https_management.advertise_on_slo_internet_vip.tls_certificates.blindfold.policy](resources--nfv_service--reference--group-003.md#canonical-3030332330202020-0032332011322012-1323330321213133-2310322102210232-2113210130330003-3322113023200120-0201300312302211-1233032301221023) |
| `https_management.advertise_on_slo_internet_vip.tls_certificates.blindfold.prepared_identity` | [https_management.advertise_on_slo_internet_vip.tls_certificates.blindfold.prepared_identity](resources--nfv_service--reference--group-003.md#canonical-3033233102312122-2203113111121311-2022011000231230-2133130122300113-1323221021201020-0322111013323312-3030133001233131-3123010003301012) |
| `https_management.advertise_on_slo_internet_vip.tls_certificates.blindfold.private_key_file` | [https_management.advertise_on_slo_internet_vip.tls_certificates.blindfold.private_key_file](resources--nfv_service--reference--group-003.md#canonical-2323110233011012-2301211010121111-0122203002010000-2101120011122122-0303123033012031-2300010303110003-0200330123002210-1123222100300312) |
| `https_management.advertise_on_slo_internet_vip.tls_certificates.blindfold.private_key_wo` | [https_management.advertise_on_slo_internet_vip.tls_certificates.blindfold.private_key_wo](resources--nfv_service--reference--group-003.md#canonical-0010032310131212-2232210320213231-3032302013030132-2200001200323013-3222020200121211-2001130031130120-1331213002120010-0233303120121013) |
| `https_management.advertise_on_slo_internet_vip.tls_certificates.blindfold.spki_identity` | [https_management.advertise_on_slo_internet_vip.tls_certificates.blindfold.spki_identity](resources--nfv_service--reference--group-003.md#canonical-3102002100021031-2211321231232123-2110120211130120-1210103011322033-3001130202101221-0220331033002213-0313120310321322-0212321230022323) |
| `https_management.advertise_on_slo_internet_vip.tls_certificates.certificate_url` | [https_management.advertise_on_slo_internet_vip.tls_certificates.certificate_url](resources--nfv_service--reference--group-002.md#canonical-0233103230320023-3013012130111110-0122001211003201-0212103023231110-2332230010120200-0320303300202230-0221332000013012-2333231313230310) |
| `https_management.advertise_on_slo_internet_vip.tls_certificates.custom_hash_algorithms` | [https_management.advertise_on_slo_internet_vip.tls_certificates.custom_hash_algorithms](resources--nfv_service--reference--group-003.md#canonical-1301221302130022-3100011012233003-1220310300032131-1131311020200322-1033101233020303-2133333122320312-3312133031100300-2131003031032121) |
| `https_management.advertise_on_slo_internet_vip.tls_certificates.custom_hash_algorithms.hash_algorithms` | [https_management.advertise_on_slo_internet_vip.tls_certificates.custom_hash_algorithms.hash_algorithms](resources--nfv_service--reference--group-003.md#canonical-0202120302110320-0010312022111110-3332133320230302-1033221223123011-2032311300011030-0323211332013030-2312322301221122-3121301130003300) |
| `https_management.advertise_on_slo_internet_vip.tls_certificates.description_spec` | [https_management.advertise_on_slo_internet_vip.tls_certificates.description_spec](resources--nfv_service--reference--group-002.md#canonical-0322331323100131-2031010003002123-0213202102023122-0303102333133232-0032013001000333-2122001023113002-0201320002312232-0130031013103033) |
| `https_management.advertise_on_slo_internet_vip.tls_certificates.disable_ocsp_stapling` | [https_management.advertise_on_slo_internet_vip.tls_certificates.disable_ocsp_stapling](resources--nfv_service--reference--group-003.md#canonical-0003003331023023-0133102022030100-2320312311111100-1310021121102013-3133032212102321-3102102302031132-3301330333032221-2103102231230031) |
| `https_management.advertise_on_slo_internet_vip.tls_certificates.private_key` | [https_management.advertise_on_slo_internet_vip.tls_certificates.private_key](resources--nfv_service--reference--group-003.md#canonical-3031210022033230-0321121201200100-3300323103310220-0202123202113101-2333210112103113-2203301202312220-0220001102123121-1313213131030031) |
| `https_management.advertise_on_slo_internet_vip.tls_certificates.private_key.blindfold_secret_info` | [https_management.advertise_on_slo_internet_vip.tls_certificates.private_key.blindfold_secret_info](resources--nfv_service--reference--group-003.md#canonical-2323130220122302-0232133013002203-1132201012120133-3310301102012210-1210301031311222-3211022110131203-0203130020001130-2333321320111003) |
| `https_management.advertise_on_slo_internet_vip.tls_certificates.private_key.blindfold_secret_info.decryption_provider` | [https_management.advertise_on_slo_internet_vip.tls_certificates.private_key.blindfold_secret_info.decryption_provider](resources--nfv_service--reference--group-003.md#canonical-2333201221322033-1010132232002223-3221103023330130-2332132120102311-3220023202020021-1302010303022003-2203310030112331-2031213003110012) |
| `https_management.advertise_on_slo_internet_vip.tls_certificates.private_key.blindfold_secret_info.location` | [https_management.advertise_on_slo_internet_vip.tls_certificates.private_key.blindfold_secret_info.location](resources--nfv_service--reference--group-003.md#canonical-2312202201302310-1120101031023022-0332012122103321-2110221332310300-3311203110302331-3321001001012030-2301332000112331-0333311132031321) |
| `https_management.advertise_on_slo_internet_vip.tls_certificates.private_key.blindfold_secret_info.store_provider` | [https_management.advertise_on_slo_internet_vip.tls_certificates.private_key.blindfold_secret_info.store_provider](resources--nfv_service--reference--group-003.md#canonical-2123330132300230-1202230302201303-2012203023201221-0221002002022330-1320123210130001-3211230233122202-0330013332032300-0011313033202331) |
| `https_management.advertise_on_slo_internet_vip.tls_certificates.private_key.clear_secret_info` | [https_management.advertise_on_slo_internet_vip.tls_certificates.private_key.clear_secret_info](resources--nfv_service--reference--group-003.md#canonical-3020222030301022-1301031200022110-2023131313131010-1313030000200200-2111220233122110-2033232230312121-1030123003313101-1202330012010121) |
| `https_management.advertise_on_slo_internet_vip.tls_certificates.private_key.clear_secret_info.provider_ref` | [https_management.advertise_on_slo_internet_vip.tls_certificates.private_key.clear_secret_info.provider_ref](resources--nfv_service--reference--group-003.md#canonical-0321332312131111-1121220202230233-0222212113113201-2121211322321030-2232232011332311-3132212221111030-2200123231330201-2230032113101100) |
| `https_management.advertise_on_slo_internet_vip.tls_certificates.private_key.clear_secret_info.url` | [https_management.advertise_on_slo_internet_vip.tls_certificates.private_key.clear_secret_info.url](resources--nfv_service--reference--group-003.md#canonical-2100131110232100-2221301100222003-3111202002031020-2013031210220012-2102011303101230-3210230313222103-1130132231022202-3223232321022111) |
| `https_management.advertise_on_slo_internet_vip.tls_certificates.use_system_defaults` | [https_management.advertise_on_slo_internet_vip.tls_certificates.use_system_defaults](resources--nfv_service--reference--group-003.md#canonical-0220112311302233-2122212113132013-0233231201302101-1112223231301202-2032011311001030-2231333111130113-2110312212213220-2312301001110300) |
| `https_management.advertise_on_slo_internet_vip.tls_config` | [https_management.advertise_on_slo_internet_vip.tls_config](resources--nfv_service--reference--group-003.md#canonical-1322100212101002-2101201122301130-2332320101220302-2230332100223301-0223311132320010-0212111220032220-1311113301100321-1331322031222320) |
| `https_management.advertise_on_slo_internet_vip.tls_config.custom_security` | [https_management.advertise_on_slo_internet_vip.tls_config.custom_security](resources--nfv_service--reference--group-003.md#canonical-1110212103322000-1013330322012313-3323120112303113-3222010211200331-2120231300311131-1311311201123032-1122221212131220-1121013123212233) |
| `https_management.advertise_on_slo_internet_vip.tls_config.custom_security.cipher_suites` | [https_management.advertise_on_slo_internet_vip.tls_config.custom_security.cipher_suites](resources--nfv_service--reference--group-003.md#canonical-2201132011011232-3111211002233201-3223300323000121-2131002003333323-3102222211110231-0021023210023032-1330032300233203-2131220120110021) |
| `https_management.advertise_on_slo_internet_vip.tls_config.custom_security.max_version` | [https_management.advertise_on_slo_internet_vip.tls_config.custom_security.max_version](resources--nfv_service--reference--group-003.md#canonical-0101023300312311-1303202012010120-2101231121302302-0323233113221003-2121003121210320-0003012333002322-2302232210003302-0321212222021103) |
| `https_management.advertise_on_slo_internet_vip.tls_config.custom_security.min_version` | [https_management.advertise_on_slo_internet_vip.tls_config.custom_security.min_version](resources--nfv_service--reference--group-003.md#canonical-1030303032300000-2113132130001330-2110210002203021-0102132031212110-2222130101203120-3032320030232212-3301331330010330-2200332131330003) |
| `https_management.advertise_on_slo_internet_vip.tls_config.default_security` | [https_management.advertise_on_slo_internet_vip.tls_config.default_security](resources--nfv_service--reference--group-003.md#canonical-2111021302020010-1212233132030233-3033202030220223-1123310233232202-2133023300010023-0233111032020020-3212130022102102-0220103103013121) |
| `https_management.advertise_on_slo_internet_vip.tls_config.low_security` | [https_management.advertise_on_slo_internet_vip.tls_config.low_security](resources--nfv_service--reference--group-003.md#canonical-1233103000320302-0301003332100021-2000232131123320-3333132020020002-3221202311310333-3002301221010220-0303103131123132-0311321102000223) |
| `https_management.advertise_on_slo_internet_vip.tls_config.medium_security` | [https_management.advertise_on_slo_internet_vip.tls_config.medium_security](resources--nfv_service--reference--group-003.md#canonical-0232022330032012-1320012332332331-3202210121333312-3230011333201001-1303210011231011-3331333033113001-3021033312312003-2031223103221202) |
| `https_management.advertise_on_slo_internet_vip.use_mtls` | [https_management.advertise_on_slo_internet_vip.use_mtls](resources--nfv_service--reference--group-003.md#canonical-0203200203212321-2132002301201203-0313000213332021-2310211333332113-3201321010032332-2111233132100032-0030012210220233-2313210003013012) |
| `https_management.advertise_on_slo_internet_vip.use_mtls.client_certificate_optional` | [https_management.advertise_on_slo_internet_vip.use_mtls.client_certificate_optional](resources--nfv_service--reference--group-003.md#canonical-2211002013323210-0013323113233123-1300332312231102-2132323331213003-2303223101100311-0103321113231302-1221123101111300-3112331321012032) |
| `https_management.advertise_on_slo_internet_vip.use_mtls.crl` | [https_management.advertise_on_slo_internet_vip.use_mtls.crl](resources--nfv_service--reference--group-003.md#canonical-3030101113312111-2003312332110222-0313222332010033-3202333010003123-2003132310023121-1012121032302320-0101001103221112-3322013030332022) |
| `https_management.advertise_on_slo_internet_vip.use_mtls.crl.name` | [https_management.advertise_on_slo_internet_vip.use_mtls.crl.name](resources--nfv_service--reference--group-003.md#canonical-0030133212103232-2211333103032311-1321102002312000-3313312033130220-0301202023223203-1020010233111130-0220032201111011-0131101211222331) |
| `https_management.advertise_on_slo_internet_vip.use_mtls.crl.namespace` | [https_management.advertise_on_slo_internet_vip.use_mtls.crl.namespace](resources--nfv_service--reference--group-003.md#canonical-1033110012030301-1330212313012300-0202123011030022-2123212320200231-1011303322121003-1331323233301303-1012000320103321-0233031212331202) |
| `https_management.advertise_on_slo_internet_vip.use_mtls.crl.tenant` | [https_management.advertise_on_slo_internet_vip.use_mtls.crl.tenant](resources--nfv_service--reference--group-003.md#canonical-1102002020022313-3002213222313030-2122200132033223-2020322321000131-2302231201313313-1221233032013320-3103322231322132-2110312132131301) |
| `https_management.advertise_on_slo_internet_vip.use_mtls.no_crl` | [https_management.advertise_on_slo_internet_vip.use_mtls.no_crl](resources--nfv_service--reference--group-003.md#canonical-0131033311220010-0111221133010001-3013200000202122-2103123033323133-3200032303132331-3102220030203012-3020231213211231-2111122120210021) |
| `https_management.advertise_on_slo_internet_vip.use_mtls.trusted_ca` | [https_management.advertise_on_slo_internet_vip.use_mtls.trusted_ca](resources--nfv_service--reference--group-003.md#canonical-3212200110222312-2321131210222313-0321011000130032-0132302100212101-2332033032103311-1311011031113112-1220032102033323-2130201311123212) |
| `https_management.advertise_on_slo_internet_vip.use_mtls.trusted_ca.name` | [https_management.advertise_on_slo_internet_vip.use_mtls.trusted_ca.name](resources--nfv_service--reference--group-003.md#canonical-3212120022223222-3323031213022133-2330220203100002-0332121310012230-3122133121100003-3003222331313311-2331321300310332-1132031122320200) |
| `https_management.advertise_on_slo_internet_vip.use_mtls.trusted_ca.namespace` | [https_management.advertise_on_slo_internet_vip.use_mtls.trusted_ca.namespace](resources--nfv_service--reference--group-003.md#canonical-2030202333213210-1003313111010211-3322033331322103-1330321133302020-1101112213111001-0113001121213312-1211100212211121-2131030200121321) |
| `https_management.advertise_on_slo_internet_vip.use_mtls.trusted_ca.tenant` | [https_management.advertise_on_slo_internet_vip.use_mtls.trusted_ca.tenant](resources--nfv_service--reference--group-003.md#canonical-3030112230010031-2102221330322332-1331103202310132-1202210200111200-3112222202203120-3310100031220003-1331221212102212-1200331033122301) |
| `https_management.advertise_on_slo_internet_vip.use_mtls.trusted_ca_url` | [https_management.advertise_on_slo_internet_vip.use_mtls.trusted_ca_url](resources--nfv_service--reference--group-003.md#canonical-1210221023303021-1333310000102202-0012003221023203-0302201332013210-1021223300210313-1100311023103210-0020112302210330-1121122220201302) |
| `https_management.advertise_on_slo_internet_vip.use_mtls.xfcc_disabled` | [https_management.advertise_on_slo_internet_vip.use_mtls.xfcc_disabled](resources--nfv_service--reference--group-003.md#canonical-3132303122120001-2322032211030112-2133320303120223-1113203331002120-2033212003130223-2210110300121130-0330022321112210-2122012223331113) |
| `https_management.advertise_on_slo_internet_vip.use_mtls.xfcc_options` | [https_management.advertise_on_slo_internet_vip.use_mtls.xfcc_options](resources--nfv_service--reference--group-003.md#canonical-3212211231012230-0230223010133130-2030210000303021-2300303312011203-0132123010231213-0031113032111122-1013031112222301-0210023022120133) |
| `https_management.advertise_on_slo_internet_vip.use_mtls.xfcc_options.xfcc_header_elements` | [https_management.advertise_on_slo_internet_vip.use_mtls.xfcc_options.xfcc_header_elements](resources--nfv_service--reference--group-003.md#canonical-2110300210103203-1200312202131301-1030000130232011-1130002233201310-2210033223021203-1313032132011213-3003010333022333-1032102101221011) |
| `https_management.advertise_on_slo_sli` | [https_management.advertise_on_slo_sli](resources--nfv_service--reference--group-003.md#canonical-1222310230103003-2030221101312331-3000330233111031-2200303020312322-0110001110321021-2333201123200123-1332223311201110-2330030101210230) |
| `https_management.advertise_on_slo_sli.no_mtls` | [https_management.advertise_on_slo_sli.no_mtls](resources--nfv_service--reference--group-003.md#canonical-0322033010311130-0022320003121201-3113011020223313-3232231003321320-0300230030123202-2203122112030232-0323133000022110-0022203221323023) |
| `https_management.advertise_on_slo_sli.tls_certificates` | [https_management.advertise_on_slo_sli.tls_certificates](resources--nfv_service--reference--group-003.md#canonical-2222132112223302-0332312122221233-0320110121020203-1031001022030002-3123021201112213-3020300223003103-0003103121332201-1223202131101201) |
| `https_management.advertise_on_slo_sli.tls_certificates.blindfold` | [https_management.advertise_on_slo_sli.tls_certificates.blindfold](resources--nfv_service--reference--group-003.md#canonical-1033112023221022-3210300330010000-0301302200012331-3120331130002310-1311212112101212-2013213212010301-2203213011113200-0133323010321203) |
| `https_management.advertise_on_slo_sli.tls_certificates.blindfold.algorithm` | [https_management.advertise_on_slo_sli.tls_certificates.blindfold.algorithm](resources--nfv_service--reference--group-003.md#canonical-2323213030201302-2213301110312313-2100210323100102-2113302201010330-3030322100112221-1203022103301220-1003323021011232-3232023302200111) |
| `https_management.advertise_on_slo_sli.tls_certificates.blindfold.certificate_file` | [https_management.advertise_on_slo_sli.tls_certificates.blindfold.certificate_file](resources--nfv_service--reference--group-003.md#canonical-3132022013010100-1120102120101321-0131311111311301-0211011201333120-3100320111121333-3300130320223321-3130123303120022-3320031303032110) |
| `https_management.advertise_on_slo_sli.tls_certificates.blindfold.certificate_pem` | [https_management.advertise_on_slo_sli.tls_certificates.blindfold.certificate_pem](resources--nfv_service--reference--group-003.md#canonical-1310201101010212-3231330133101333-1211020311210023-2330202232121100-0102310332022103-0212111302102001-3220122312023102-1211330123030130) |
| `https_management.advertise_on_slo_sli.tls_certificates.blindfold.chain_identity` | [https_management.advertise_on_slo_sli.tls_certificates.blindfold.chain_identity](resources--nfv_service--reference--group-003.md#canonical-1010331003303320-0101112301220202-0003012101213211-2002333111002030-2031311132123101-1020311032133012-1221313101030312-3111221222031133) |
| `https_management.advertise_on_slo_sli.tls_certificates.blindfold.context_digest` | [https_management.advertise_on_slo_sli.tls_certificates.blindfold.context_digest](resources--nfv_service--reference--group-003.md#canonical-0011033110303102-2120200112031102-1000031322132122-1020033300200112-0221231121121330-0203102201213102-2101300100121031-3323100312320221) |
| `https_management.advertise_on_slo_sli.tls_certificates.blindfold.encrypted_location` | [https_management.advertise_on_slo_sli.tls_certificates.blindfold.encrypted_location](resources--nfv_service--reference--group-003.md#canonical-3000331222310230-1220131020013000-3331012332100102-2331212020310111-2121223301320031-1133002132333013-1312312221023200-1222200321222320) |
| `https_management.advertise_on_slo_sli.tls_certificates.blindfold.expires_at` | [https_management.advertise_on_slo_sli.tls_certificates.blindfold.expires_at](resources--nfv_service--reference--group-003.md#canonical-1013023111300230-0323102210322310-0230312113300202-3331022213310013-0020322022212303-2111110333012212-1001203202103020-1102110033220230) |
| `https_management.advertise_on_slo_sli.tls_certificates.blindfold.fingerprint` | [https_management.advertise_on_slo_sli.tls_certificates.blindfold.fingerprint](resources--nfv_service--reference--group-003.md#canonical-0203223022003222-1103233121122210-3133121210020220-3100101333023302-2001201321020320-1211323231211300-3200111200103313-2132323213330201) |
| `https_management.advertise_on_slo_sli.tls_certificates.blindfold.id` | [https_management.advertise_on_slo_sli.tls_certificates.blindfold.id](resources--nfv_service--reference--group-003.md#canonical-2112002200302221-0313130002201033-1023132001332301-0300333113330330-3010303010221323-3000231202013103-2220303000201331-3010311031121311) |
| `https_management.advertise_on_slo_sli.tls_certificates.blindfold.material_version` | [https_management.advertise_on_slo_sli.tls_certificates.blindfold.material_version](resources--nfv_service--reference--group-003.md#canonical-3221121120131301-0101000111323002-2101210033100310-2031003301121221-2221211210230230-1010230333222330-3201300100213000-3322033002012331) |
| `https_management.advertise_on_slo_sli.tls_certificates.blindfold.passphrase_env` | [https_management.advertise_on_slo_sli.tls_certificates.blindfold.passphrase_env](resources--nfv_service--reference--group-003.md#canonical-0221010111332231-0000301211330130-1113323322210232-3202311203202001-2300302222112001-1103113203210212-1130330122111001-3332221231330232) |
| `https_management.advertise_on_slo_sli.tls_certificates.blindfold.passphrase_wo` | [https_management.advertise_on_slo_sli.tls_certificates.blindfold.passphrase_wo](resources--nfv_service--reference--group-003.md#canonical-3103100333101321-3122103120032111-3100202231212330-0002221212001300-1100102122020002-2030212011011230-3322021112321112-0113232230223033) |
| `https_management.advertise_on_slo_sli.tls_certificates.blindfold.pkcs12_file` | [https_management.advertise_on_slo_sli.tls_certificates.blindfold.pkcs12_file](resources--nfv_service--reference--group-003.md#canonical-1223333230232033-1002121302131310-2011110330332232-2212101003123210-3311310000103313-1321022121232233-3110222322221002-2011203303330322) |
| `https_management.advertise_on_slo_sli.tls_certificates.blindfold.pkcs12_wo` | [https_management.advertise_on_slo_sli.tls_certificates.blindfold.pkcs12_wo](resources--nfv_service--reference--group-003.md#canonical-1111212231022031-1312003330033312-0012200102010212-3131320213302223-2112120033310123-1121211020322212-0001313133330132-1203032113231002) |
| `https_management.advertise_on_slo_sli.tls_certificates.blindfold.policy` | [https_management.advertise_on_slo_sli.tls_certificates.blindfold.policy](resources--nfv_service--reference--group-003.md#canonical-1221022101330200-1310012210323031-0321232332110031-1200201203302110-0131113211002331-1212212103003120-3133231020321132-3311112211313112) |
| `https_management.advertise_on_slo_sli.tls_certificates.blindfold.prepared_identity` | [https_management.advertise_on_slo_sli.tls_certificates.blindfold.prepared_identity](resources--nfv_service--reference--group-003.md#canonical-1231120311000101-1123201321113110-3100301330220323-3101231031023011-2032232132032101-0333212201100120-3133333323233111-2220302121323302) |
| `https_management.advertise_on_slo_sli.tls_certificates.blindfold.private_key_file` | [https_management.advertise_on_slo_sli.tls_certificates.blindfold.private_key_file](resources--nfv_service--reference--group-003.md#canonical-2332333133330202-0131310113302333-2012130013000102-1022220211120230-1130232120211223-1111110311132333-2012210230330213-2333012122000103) |
| `https_management.advertise_on_slo_sli.tls_certificates.blindfold.private_key_wo` | [https_management.advertise_on_slo_sli.tls_certificates.blindfold.private_key_wo](resources--nfv_service--reference--group-003.md#canonical-1010210203200330-3302102332222030-1001231311100130-0302102130030012-3212132022213221-3332222213301330-2313301323212013-2010103010211313) |
| `https_management.advertise_on_slo_sli.tls_certificates.blindfold.spki_identity` | [https_management.advertise_on_slo_sli.tls_certificates.blindfold.spki_identity](resources--nfv_service--reference--group-003.md#canonical-1031312031233300-1030103003110133-2213212301232001-2232110020130332-3100133330120310-0200222213320121-1230000332000303-1131000122121013) |
| `https_management.advertise_on_slo_sli.tls_certificates.certificate_url` | [https_management.advertise_on_slo_sli.tls_certificates.certificate_url](resources--nfv_service--reference--group-003.md#canonical-0110213100123121-1000322120212120-0321033020211011-2203113110023200-3003321002200333-3012330102212333-1000121101130321-0330221303001100) |
| `https_management.advertise_on_slo_sli.tls_certificates.custom_hash_algorithms` | [https_management.advertise_on_slo_sli.tls_certificates.custom_hash_algorithms](resources--nfv_service--reference--group-003.md#canonical-0023013102032200-0233131301222211-1200232321210020-1320032131133020-2133332103012210-1200220200200212-0122333030002123-2030200312110010) |
| `https_management.advertise_on_slo_sli.tls_certificates.custom_hash_algorithms.hash_algorithms` | [https_management.advertise_on_slo_sli.tls_certificates.custom_hash_algorithms.hash_algorithms](resources--nfv_service--reference--group-003.md#canonical-1003032001010322-2110330013122202-2113312222103000-1321220133331130-3013311000212331-1020122231221201-0113000212023313-3131301132321233) |
| `https_management.advertise_on_slo_sli.tls_certificates.description_spec` | [https_management.advertise_on_slo_sli.tls_certificates.description_spec](resources--nfv_service--reference--group-003.md#canonical-1010202211022311-0211131011331123-1002311102113312-2001110233223000-0201022233222112-3310013103000230-1333211321312202-0110111232332013) |
| `https_management.advertise_on_slo_sli.tls_certificates.disable_ocsp_stapling` | [https_management.advertise_on_slo_sli.tls_certificates.disable_ocsp_stapling](resources--nfv_service--reference--group-003.md#canonical-3011233100130321-3013222003232231-0230100121203322-1110102211221320-0033322213102203-2333210010333202-1012231213200221-3121103121133003) |
| `https_management.advertise_on_slo_sli.tls_certificates.private_key` | [https_management.advertise_on_slo_sli.tls_certificates.private_key](resources--nfv_service--reference--group-003.md#canonical-0211221232320211-2013103032113323-3103010030102020-2220102022233032-3200110030000211-3113311111231130-1203213300323010-1202312333113123) |
| `https_management.advertise_on_slo_sli.tls_certificates.private_key.blindfold_secret_info` | [https_management.advertise_on_slo_sli.tls_certificates.private_key.blindfold_secret_info](resources--nfv_service--reference--group-003.md#canonical-0122301330013003-3232321333221120-1303310131300323-0233113032301133-0013112010302101-3222110012001311-0333020021210111-3220003330122300) |
| `https_management.advertise_on_slo_sli.tls_certificates.private_key.blindfold_secret_info.decryption_provider` | [https_management.advertise_on_slo_sli.tls_certificates.private_key.blindfold_secret_info.decryption_provider](resources--nfv_service--reference--group-003.md#canonical-3222003112330000-1031331012330322-3122120131022013-0323322032120322-1102222200231101-0033123130322010-1100332322322301-1202200320002020) |
| `https_management.advertise_on_slo_sli.tls_certificates.private_key.blindfold_secret_info.location` | [https_management.advertise_on_slo_sli.tls_certificates.private_key.blindfold_secret_info.location](resources--nfv_service--reference--group-003.md#canonical-2232312211233331-3100112233111212-0111103012123122-3332012230001002-2020001032202102-0010122320002030-1001301223102123-1322233211313210) |
| `https_management.advertise_on_slo_sli.tls_certificates.private_key.blindfold_secret_info.store_provider` | [https_management.advertise_on_slo_sli.tls_certificates.private_key.blindfold_secret_info.store_provider](resources--nfv_service--reference--group-003.md#canonical-1023100103112010-0003032120211000-0113100122223320-1222121132113200-0103221232233022-0110322100211102-0113022000131302-2023212010112312) |
| `https_management.advertise_on_slo_sli.tls_certificates.private_key.clear_secret_info` | [https_management.advertise_on_slo_sli.tls_certificates.private_key.clear_secret_info](resources--nfv_service--reference--group-003.md#canonical-0223110131031313-3121233223320320-1310233230301303-2212301110220013-2120201122333033-2322323210022020-2220100202302223-1122200130231300) |
| `https_management.advertise_on_slo_sli.tls_certificates.private_key.clear_secret_info.provider_ref` | [https_management.advertise_on_slo_sli.tls_certificates.private_key.clear_secret_info.provider_ref](resources--nfv_service--reference--group-003.md#canonical-3132220103031220-3020022000132030-2212330212120103-1321113013023323-1332002320122100-0201302300201110-1310130301323231-0123101300033121) |
| `https_management.advertise_on_slo_sli.tls_certificates.private_key.clear_secret_info.url` | [https_management.advertise_on_slo_sli.tls_certificates.private_key.clear_secret_info.url](resources--nfv_service--reference--group-003.md#canonical-0212022103102203-1321011132010213-3201211200330313-3223303100001031-1303001322303232-1012303303002020-3321321202031131-2111113002033333) |
| `https_management.advertise_on_slo_sli.tls_certificates.use_system_defaults` | [https_management.advertise_on_slo_sli.tls_certificates.use_system_defaults](resources--nfv_service--reference--group-003.md#canonical-2200301201111133-2313022022013102-3113110020000000-2131032112121023-2222223210213302-1311210202322013-1130322332032001-0033303321223130) |
| `https_management.advertise_on_slo_sli.tls_config` | [https_management.advertise_on_slo_sli.tls_config](resources--nfv_service--reference--group-003.md#canonical-3031201110103220-2332033131020233-3100133000113300-0220032103110233-3333333321111132-0201113321230023-0232030123131331-3113333120110222) |
| `https_management.advertise_on_slo_sli.tls_config.custom_security` | [https_management.advertise_on_slo_sli.tls_config.custom_security](resources--nfv_service--reference--group-003.md#canonical-3000120202122123-3002200103132220-3221200331221031-3213122123223111-0201231213103211-2223130310123132-2002210210302001-2300033213311313) |
| `https_management.advertise_on_slo_sli.tls_config.custom_security.cipher_suites` | [https_management.advertise_on_slo_sli.tls_config.custom_security.cipher_suites](resources--nfv_service--reference--group-003.md#canonical-1302021011020110-0121132011323321-0203323033133333-0102003022301201-2231021221111012-1030210210300110-3121213131013310-0320332220322223) |
| `https_management.advertise_on_slo_sli.tls_config.custom_security.max_version` | [https_management.advertise_on_slo_sli.tls_config.custom_security.max_version](resources--nfv_service--reference--group-003.md#canonical-2023000203003103-0230132300331311-3001013112300212-3110031132030132-2223123101301031-3221220223300210-1210013303311131-3131121331103132) |
| `https_management.advertise_on_slo_sli.tls_config.custom_security.min_version` | [https_management.advertise_on_slo_sli.tls_config.custom_security.min_version](resources--nfv_service--reference--group-003.md#canonical-2331130001230231-0001212313112222-0302001110322100-1200121003210133-1313133012113233-0303313010133222-3122221313111200-3000320023131321) |
| `https_management.advertise_on_slo_sli.tls_config.default_security` | [https_management.advertise_on_slo_sli.tls_config.default_security](resources--nfv_service--reference--group-004.md#canonical-2211111111233321-0331013132321302-3320211030003020-1223022020000203-1100120100000120-0123112203222110-0332211033110130-2103220123030321) |
| `https_management.advertise_on_slo_sli.tls_config.low_security` | [https_management.advertise_on_slo_sli.tls_config.low_security](resources--nfv_service--reference--group-004.md#canonical-3011310322301103-3000300013103011-3101221220033103-3223112202023023-3322113113133100-0211312112220302-0131210110130212-1033311012332032) |
| `https_management.advertise_on_slo_sli.tls_config.medium_security` | [https_management.advertise_on_slo_sli.tls_config.medium_security](resources--nfv_service--reference--group-004.md#canonical-2022100312330102-2103131100103111-2202030202121132-2303013301020101-1311101102203012-2122011310111232-1013212131110203-3332011332302023) |
| `https_management.advertise_on_slo_sli.use_mtls` | [https_management.advertise_on_slo_sli.use_mtls](resources--nfv_service--reference--group-004.md#canonical-0101002012210331-3130133103321132-0122211223123002-3131001203213102-0123232111233223-0032233021132033-1101330321232312-3321311203121303) |
| `https_management.advertise_on_slo_sli.use_mtls.client_certificate_optional` | [https_management.advertise_on_slo_sli.use_mtls.client_certificate_optional](resources--nfv_service--reference--group-004.md#canonical-1202210002310331-1110132200031013-2021320321222332-1300301232123000-0022022022230212-0311032010312010-0003130130122203-3231223002012032) |
| `https_management.advertise_on_slo_sli.use_mtls.crl` | [https_management.advertise_on_slo_sli.use_mtls.crl](resources--nfv_service--reference--group-004.md#canonical-2102223303121033-3003323311233023-1310300103003310-2311001233230132-0333231323232111-1102111212221010-1130033012313322-2222002311022231) |
| `https_management.advertise_on_slo_sli.use_mtls.crl.name` | [https_management.advertise_on_slo_sli.use_mtls.crl.name](resources--nfv_service--reference--group-004.md#canonical-3011022203222132-1132111132033011-1130022201303223-2331312302203300-2313023311231233-2320012223203120-1331203130111020-2300131302121003) |
| `https_management.advertise_on_slo_sli.use_mtls.crl.namespace` | [https_management.advertise_on_slo_sli.use_mtls.crl.namespace](resources--nfv_service--reference--group-004.md#canonical-2230033001232033-2120033231223231-3332201020213230-3332311002030110-3002310202032301-3332100332320202-0133302222320203-3330031322301103) |
| `https_management.advertise_on_slo_sli.use_mtls.crl.tenant` | [https_management.advertise_on_slo_sli.use_mtls.crl.tenant](resources--nfv_service--reference--group-004.md#canonical-0332303302220003-0332222112201002-2310131133010221-2302103302302223-2301223300013003-3233013122002010-2221310212032123-2103302300030311) |
| `https_management.advertise_on_slo_sli.use_mtls.no_crl` | [https_management.advertise_on_slo_sli.use_mtls.no_crl](resources--nfv_service--reference--group-004.md#canonical-0030012230222331-2002311130032222-0011012220130033-1030323331011220-0223000113223232-2231303130322210-3201213101333312-2220132130033011) |
| `https_management.advertise_on_slo_sli.use_mtls.trusted_ca` | [https_management.advertise_on_slo_sli.use_mtls.trusted_ca](resources--nfv_service--reference--group-004.md#canonical-3131023010230002-0231002203322301-1132012311333303-1310010312023212-0320131120220230-2221212300110100-3011231323110122-1312223033212110) |
| `https_management.advertise_on_slo_sli.use_mtls.trusted_ca.name` | [https_management.advertise_on_slo_sli.use_mtls.trusted_ca.name](resources--nfv_service--reference--group-004.md#canonical-1301301030001303-0020323202023312-1313333113033231-2110323101021310-3022100012321201-1322201303111133-1231030012311131-0123301111030233) |
| `https_management.advertise_on_slo_sli.use_mtls.trusted_ca.namespace` | [https_management.advertise_on_slo_sli.use_mtls.trusted_ca.namespace](resources--nfv_service--reference--group-004.md#canonical-1302130320313333-2012313032001211-1200221321210311-2102031331103321-0213303110013013-1203310110112031-3220120122320332-3312200113132232) |
| `https_management.advertise_on_slo_sli.use_mtls.trusted_ca.tenant` | [https_management.advertise_on_slo_sli.use_mtls.trusted_ca.tenant](resources--nfv_service--reference--group-004.md#canonical-0221210212212310-2103332213201110-3032201023222313-3030120103322020-2110210300201123-0013123203310002-0003223203322030-0111101023331002) |
| `https_management.advertise_on_slo_sli.use_mtls.trusted_ca_url` | [https_management.advertise_on_slo_sli.use_mtls.trusted_ca_url](resources--nfv_service--reference--group-004.md#canonical-0300102112310320-0012011321030133-2103130122102211-1002202332202301-2121321330312133-3323213030132013-0221200312122120-3010221100103013) |
| `https_management.advertise_on_slo_sli.use_mtls.xfcc_disabled` | [https_management.advertise_on_slo_sli.use_mtls.xfcc_disabled](resources--nfv_service--reference--group-004.md#canonical-0230021031001003-3311031210100201-2100011213310322-2032231013301012-3033303313211333-0113110111312300-1030312301333101-1132202020211222) |
| `https_management.advertise_on_slo_sli.use_mtls.xfcc_options` | [https_management.advertise_on_slo_sli.use_mtls.xfcc_options](resources--nfv_service--reference--group-004.md#canonical-0320223201130322-1301030333320222-2120313130300112-2112210121030100-0321020002131301-1311112010311130-3323330122211021-2111211212202230) |
| `https_management.advertise_on_slo_sli.use_mtls.xfcc_options.xfcc_header_elements` | [https_management.advertise_on_slo_sli.use_mtls.xfcc_options.xfcc_header_elements](resources--nfv_service--reference--group-004.md#canonical-1232002303002201-3220113203023223-3122213121331031-3021320312302003-3121131223011212-1222310221312021-1103200021011313-1201031011203021) |
| `https_management.advertise_on_slo_vip` | [https_management.advertise_on_slo_vip](resources--nfv_service--reference--group-004.md#canonical-0311030123113020-0310323023030223-2132032033103233-3000210331321312-1131200021222000-1232022213100001-1221331332332222-3313121002001232) |
| `https_management.advertise_on_slo_vip.no_mtls` | [https_management.advertise_on_slo_vip.no_mtls](resources--nfv_service--reference--group-004.md#canonical-0223102102213323-3120131112012331-2011112320231311-1021011120131321-0223330102011211-1303132001130100-2120020021323301-0033120103310330) |
| `https_management.advertise_on_slo_vip.tls_certificates` | [https_management.advertise_on_slo_vip.tls_certificates](resources--nfv_service--reference--group-004.md#canonical-2033311030122032-3110102122003230-0011002203212233-1032211121000322-2122131133333230-0200313213020211-1202331022100100-3301023220233000) |
| `https_management.advertise_on_slo_vip.tls_certificates.blindfold` | [https_management.advertise_on_slo_vip.tls_certificates.blindfold](resources--nfv_service--reference--group-004.md#canonical-1302112220121322-3001022300213131-2323031101332113-3100232002201121-2200211323122202-2300302220111222-2113122223113130-0301311301030102) |
| `https_management.advertise_on_slo_vip.tls_certificates.blindfold.algorithm` | [https_management.advertise_on_slo_vip.tls_certificates.blindfold.algorithm](resources--nfv_service--reference--group-004.md#canonical-2022000111000223-1000103202213233-0303201101122100-1001302223332032-0002302222013012-0321321230331221-1322120102201013-1033032331233220) |
| `https_management.advertise_on_slo_vip.tls_certificates.blindfold.certificate_file` | [https_management.advertise_on_slo_vip.tls_certificates.blindfold.certificate_file](resources--nfv_service--reference--group-004.md#canonical-0131233111303100-1221230310130231-1000122122102013-2311323302121112-0132003111112003-0303202110020101-2110000113332320-3212311121020300) |
| `https_management.advertise_on_slo_vip.tls_certificates.blindfold.certificate_pem` | [https_management.advertise_on_slo_vip.tls_certificates.blindfold.certificate_pem](resources--nfv_service--reference--group-004.md#canonical-0220323011111012-1030123102003100-0002203121100012-3312010230000111-1100230111001213-0223023110330311-3011332011312321-1030020300030011) |
| `https_management.advertise_on_slo_vip.tls_certificates.blindfold.chain_identity` | [https_management.advertise_on_slo_vip.tls_certificates.blindfold.chain_identity](resources--nfv_service--reference--group-004.md#canonical-0110011102203323-2002301022302013-2213223112032300-3023233023113023-3132032001232000-1131201023122100-0123311010333332-0212011212233132) |
| `https_management.advertise_on_slo_vip.tls_certificates.blindfold.context_digest` | [https_management.advertise_on_slo_vip.tls_certificates.blindfold.context_digest](resources--nfv_service--reference--group-004.md#canonical-2120012212231101-0131000031333122-1303201230003023-3120332000320123-2013133332320012-2203201301022211-2210231303221100-3213132313102123) |
| `https_management.advertise_on_slo_vip.tls_certificates.blindfold.encrypted_location` | [https_management.advertise_on_slo_vip.tls_certificates.blindfold.encrypted_location](resources--nfv_service--reference--group-004.md#canonical-2331130312032211-0301121323010130-0110202122212023-0023020321112102-2231130123300121-3010022330202021-1221210022222333-2233112211201023) |
| `https_management.advertise_on_slo_vip.tls_certificates.blindfold.expires_at` | [https_management.advertise_on_slo_vip.tls_certificates.blindfold.expires_at](resources--nfv_service--reference--group-004.md#canonical-3130311011302221-3100220300222020-0331003110203210-3221302301231212-3233031201333230-2302212101022033-1313203223111312-0332011002330002) |
| `https_management.advertise_on_slo_vip.tls_certificates.blindfold.fingerprint` | [https_management.advertise_on_slo_vip.tls_certificates.blindfold.fingerprint](resources--nfv_service--reference--group-004.md#canonical-2113311010213032-1103133100121003-1120303232221223-2110202300101212-3222001311022030-2300101201120023-1311120313323200-3003120230030221) |
| `https_management.advertise_on_slo_vip.tls_certificates.blindfold.id` | [https_management.advertise_on_slo_vip.tls_certificates.blindfold.id](resources--nfv_service--reference--group-004.md#canonical-1002331220321031-2033213202033210-3212101311203031-1100003230022300-0020333211303310-3020303010011003-0233233131030103-0103201132013310) |
| `https_management.advertise_on_slo_vip.tls_certificates.blindfold.material_version` | [https_management.advertise_on_slo_vip.tls_certificates.blindfold.material_version](resources--nfv_service--reference--group-004.md#canonical-1112231311121130-0220300203113030-0122012121033003-0013033302112201-3003122123212002-1313222033332300-3113031133133123-3110010300010101) |
| `https_management.advertise_on_slo_vip.tls_certificates.blindfold.passphrase_env` | [https_management.advertise_on_slo_vip.tls_certificates.blindfold.passphrase_env](resources--nfv_service--reference--group-004.md#canonical-1332113131333313-3002122010313210-1302331311303220-2231303320123333-0221330001121330-3211032312031212-0230111120113121-0330320321131031) |
| `https_management.advertise_on_slo_vip.tls_certificates.blindfold.passphrase_wo` | [https_management.advertise_on_slo_vip.tls_certificates.blindfold.passphrase_wo](resources--nfv_service--reference--group-004.md#canonical-1112032212120220-3330003220200030-1100302333313020-0303120133303301-0032313233232231-1310330002123000-3100020131323100-2323333102131113) |
| `https_management.advertise_on_slo_vip.tls_certificates.blindfold.pkcs12_file` | [https_management.advertise_on_slo_vip.tls_certificates.blindfold.pkcs12_file](resources--nfv_service--reference--group-004.md#canonical-0133320133123222-2331202100103320-1111022202313031-1210213332323111-2122010332100321-1310222003120221-2031303100032321-0330313012133211) |
| `https_management.advertise_on_slo_vip.tls_certificates.blindfold.pkcs12_wo` | [https_management.advertise_on_slo_vip.tls_certificates.blindfold.pkcs12_wo](resources--nfv_service--reference--group-004.md#canonical-3332120102130120-2313313112203031-3122321313322111-2211303302313000-0031210101132212-2112131231333011-0122123230111210-3023320022301020) |
| `https_management.advertise_on_slo_vip.tls_certificates.blindfold.policy` | [https_management.advertise_on_slo_vip.tls_certificates.blindfold.policy](resources--nfv_service--reference--group-004.md#canonical-1131223112120032-0201203303010231-1321212330212111-3120221020301010-0023210012113002-1110003301223202-3212200123100023-2203031100102111) |
| `https_management.advertise_on_slo_vip.tls_certificates.blindfold.prepared_identity` | [https_management.advertise_on_slo_vip.tls_certificates.blindfold.prepared_identity](resources--nfv_service--reference--group-004.md#canonical-1111000203331321-3023133002233230-0311030112003113-3111100102321021-2312000223222230-0103030120323210-3300013102102023-3003200230302002) |
| `https_management.advertise_on_slo_vip.tls_certificates.blindfold.private_key_file` | [https_management.advertise_on_slo_vip.tls_certificates.blindfold.private_key_file](resources--nfv_service--reference--group-004.md#canonical-2020130202023132-3112123303110211-0332332133210132-2310333222013102-2312013013000021-0330301302232231-1133203010100331-3320132232132221) |
| `https_management.advertise_on_slo_vip.tls_certificates.blindfold.private_key_wo` | [https_management.advertise_on_slo_vip.tls_certificates.blindfold.private_key_wo](resources--nfv_service--reference--group-004.md#canonical-3330022333310123-0000110331320333-2112012231321003-2221113313221020-0131323301230203-3203030223211230-3121101231221033-0212302322221322) |
| `https_management.advertise_on_slo_vip.tls_certificates.blindfold.spki_identity` | [https_management.advertise_on_slo_vip.tls_certificates.blindfold.spki_identity](resources--nfv_service--reference--group-004.md#canonical-3032201320223031-0321112222303011-2100123321132110-2211111210222133-1122313312312221-3100220222030100-2012201321313203-0302312103322133) |
| `https_management.advertise_on_slo_vip.tls_certificates.certificate_url` | [https_management.advertise_on_slo_vip.tls_certificates.certificate_url](resources--nfv_service--reference--group-004.md#canonical-1120032033333030-3203101020110111-2300121111222022-0213213231333133-3222331220330300-2303112003300321-1310021131201312-1200123211122223) |
| `https_management.advertise_on_slo_vip.tls_certificates.custom_hash_algorithms` | [https_management.advertise_on_slo_vip.tls_certificates.custom_hash_algorithms](resources--nfv_service--reference--group-004.md#canonical-1123112301132312-1301222223010313-1013321112310331-2130133001312332-2332222000032303-1301231010301310-3120203211332011-2231222100321310) |
| `https_management.advertise_on_slo_vip.tls_certificates.custom_hash_algorithms.hash_algorithms` | [https_management.advertise_on_slo_vip.tls_certificates.custom_hash_algorithms.hash_algorithms](resources--nfv_service--reference--group-004.md#canonical-3213230122223201-0300123032111130-1022332131222302-0101131123311331-1302131210333213-2031331003312132-1220010222210302-1032110020322333) |
| `https_management.advertise_on_slo_vip.tls_certificates.description_spec` | [https_management.advertise_on_slo_vip.tls_certificates.description_spec](resources--nfv_service--reference--group-004.md#canonical-1301110230122310-3000311233031332-3121111031002000-3010230001301101-1101203131321121-0213110012203131-2323323311220321-3023322223020220) |
| `https_management.advertise_on_slo_vip.tls_certificates.disable_ocsp_stapling` | [https_management.advertise_on_slo_vip.tls_certificates.disable_ocsp_stapling](resources--nfv_service--reference--group-004.md#canonical-0221121130213003-0221302100231011-1332330232032013-0000321233021313-2212103103133210-1013212211333213-1330112012111233-0323300330223112) |
| `https_management.advertise_on_slo_vip.tls_certificates.private_key` | [https_management.advertise_on_slo_vip.tls_certificates.private_key](resources--nfv_service--reference--group-004.md#canonical-2123010121130221-0131130220001200-3011012221222210-1103000020112311-0111030031330223-0332201202131330-1311331301232212-1201221111331100) |
| `https_management.advertise_on_slo_vip.tls_certificates.private_key.blindfold_secret_info` | [https_management.advertise_on_slo_vip.tls_certificates.private_key.blindfold_secret_info](resources--nfv_service--reference--group-004.md#canonical-1331111321322033-3212131212031310-2231023021332322-3333003101202020-1330132302130321-1120321003203012-2230023111111213-3203033111022323) |
| `https_management.advertise_on_slo_vip.tls_certificates.private_key.blindfold_secret_info.decryption_provider` | [https_management.advertise_on_slo_vip.tls_certificates.private_key.blindfold_secret_info.decryption_provider](resources--nfv_service--reference--group-004.md#canonical-3211033203210233-3021100103321211-0113213303233220-2003022202221211-0131122331222331-0001221303201332-1211021233013131-0331103213300132) |
| `https_management.advertise_on_slo_vip.tls_certificates.private_key.blindfold_secret_info.location` | [https_management.advertise_on_slo_vip.tls_certificates.private_key.blindfold_secret_info.location](resources--nfv_service--reference--group-004.md#canonical-1132213331011100-0301000132302203-0023301313312020-1213300103201230-0330321011302011-0132002122101302-2301212231130310-0123001202102303) |
| `https_management.advertise_on_slo_vip.tls_certificates.private_key.blindfold_secret_info.store_provider` | [https_management.advertise_on_slo_vip.tls_certificates.private_key.blindfold_secret_info.store_provider](resources--nfv_service--reference--group-004.md#canonical-0330221101110201-1233332133032021-1133101120321202-2210102203220330-2123132220100123-1213332132202002-2013002223331020-3232100021301030) |
| `https_management.advertise_on_slo_vip.tls_certificates.private_key.clear_secret_info` | [https_management.advertise_on_slo_vip.tls_certificates.private_key.clear_secret_info](resources--nfv_service--reference--group-004.md#canonical-3112202333130320-2010030322332211-3130233201012322-1221302032101201-2202113231331232-1101111313303330-3321210023100121-3320231101000232) |
| `https_management.advertise_on_slo_vip.tls_certificates.private_key.clear_secret_info.provider_ref` | [https_management.advertise_on_slo_vip.tls_certificates.private_key.clear_secret_info.provider_ref](resources--nfv_service--reference--group-004.md#canonical-2330310221230030-2033003003301231-3220231311230301-1221311302303332-3010131220222321-3123321302330211-2023320220201312-3330012011300301) |
| `https_management.advertise_on_slo_vip.tls_certificates.private_key.clear_secret_info.url` | [https_management.advertise_on_slo_vip.tls_certificates.private_key.clear_secret_info.url](resources--nfv_service--reference--group-004.md#canonical-1330000113003310-1221211232022200-3332313231310121-2133222211303223-1220312220022113-1201230200203233-2022301023033230-1311032121212121) |
| `https_management.advertise_on_slo_vip.tls_certificates.use_system_defaults` | [https_management.advertise_on_slo_vip.tls_certificates.use_system_defaults](resources--nfv_service--reference--group-004.md#canonical-2031012322002012-2010113222110322-2201220223000202-2103121223311322-0010210210332020-0210210031002211-1021230010011103-0021120103201102) |
| `https_management.advertise_on_slo_vip.tls_config` | [https_management.advertise_on_slo_vip.tls_config](resources--nfv_service--reference--group-004.md#canonical-1313201200001300-0122302203303021-1011222313013121-3222212101023200-2202310100223220-2320332211322110-1030330332233120-0000220030000031) |
| `https_management.advertise_on_slo_vip.tls_config.custom_security` | [https_management.advertise_on_slo_vip.tls_config.custom_security](resources--nfv_service--reference--group-004.md#canonical-1101323212030332-1010223232300102-2033332330332011-2033111302131100-3233012310013022-3031312022211032-0021311300033323-1212103303312230) |
| `https_management.advertise_on_slo_vip.tls_config.custom_security.cipher_suites` | [https_management.advertise_on_slo_vip.tls_config.custom_security.cipher_suites](resources--nfv_service--reference--group-004.md#canonical-0331230020232022-1200003320012120-0033031001113120-1202021230311210-1312302013230323-2320030211223320-0312110303110230-0010022033212131) |
| `https_management.advertise_on_slo_vip.tls_config.custom_security.max_version` | [https_management.advertise_on_slo_vip.tls_config.custom_security.max_version](resources--nfv_service--reference--group-004.md#canonical-3222200121010012-0330230200120233-1230200100000201-3130001333130330-3320212330103012-0330033220330233-2113110113232331-2312100012301030) |
| `https_management.advertise_on_slo_vip.tls_config.custom_security.min_version` | [https_management.advertise_on_slo_vip.tls_config.custom_security.min_version](resources--nfv_service--reference--group-004.md#canonical-2321001301201123-2233213212103110-2222201000022123-1202123011102330-2330201003211002-0232030323211122-2122212133110311-1330102023331302) |
| `https_management.advertise_on_slo_vip.tls_config.default_security` | [https_management.advertise_on_slo_vip.tls_config.default_security](resources--nfv_service--reference--group-004.md#canonical-3332030021111030-1113101321222033-0201230203232021-2213021300313011-2211121121031031-2223003300211230-1012300212211330-2122330213231131) |
| `https_management.advertise_on_slo_vip.tls_config.low_security` | [https_management.advertise_on_slo_vip.tls_config.low_security](resources--nfv_service--reference--group-004.md#canonical-1000233322232110-0231303001213012-3003210203323132-0103012001221131-1023320201112120-0112310313321313-0033200012313231-2132031312203211) |
| `https_management.advertise_on_slo_vip.tls_config.medium_security` | [https_management.advertise_on_slo_vip.tls_config.medium_security](resources--nfv_service--reference--group-004.md#canonical-0023303013333212-1123100113313230-2200121022202303-2320303332321100-2100003333121222-3231121301313320-0130303121200011-3300111302033322) |
| `https_management.advertise_on_slo_vip.use_mtls` | [https_management.advertise_on_slo_vip.use_mtls](resources--nfv_service--reference--group-004.md#canonical-2200112003032112-1010021111303003-1023010233101132-0112120202101103-3020221312010011-0211211303333131-0011021310102020-1321233302210010) |
| `https_management.advertise_on_slo_vip.use_mtls.client_certificate_optional` | [https_management.advertise_on_slo_vip.use_mtls.client_certificate_optional](resources--nfv_service--reference--group-004.md#canonical-2313330321013002-3202013002330222-3233113013221101-2033331002030032-3230000132200021-2200023131323111-2221032030201302-2223312202000202) |
| `https_management.advertise_on_slo_vip.use_mtls.crl` | [https_management.advertise_on_slo_vip.use_mtls.crl](resources--nfv_service--reference--group-004.md#canonical-2130002231333210-1313222133110302-1202303320213031-1120222210010333-2113311301133022-0223110211302332-2001012300332310-1033230133111202) |
| `https_management.advertise_on_slo_vip.use_mtls.crl.name` | [https_management.advertise_on_slo_vip.use_mtls.crl.name](resources--nfv_service--reference--group-004.md#canonical-0002321323323132-1010301303030231-2320113002130013-1020301220023133-2300320200323312-1103232022122330-1120113221331012-0210120301132230) |
| `https_management.advertise_on_slo_vip.use_mtls.crl.namespace` | [https_management.advertise_on_slo_vip.use_mtls.crl.namespace](resources--nfv_service--reference--group-004.md#canonical-3312212220211032-3201123022211330-2033303213322312-3012001011302233-3003213023103321-3032313312200223-2302301132032330-2122321210231301) |
| `https_management.advertise_on_slo_vip.use_mtls.crl.tenant` | [https_management.advertise_on_slo_vip.use_mtls.crl.tenant](resources--nfv_service--reference--group-004.md#canonical-3103200223220101-1112312121012233-0212111223230112-2201232200331020-3211311233032203-0100303010200010-3313313320032223-0311302320002003) |
| `https_management.advertise_on_slo_vip.use_mtls.no_crl` | [https_management.advertise_on_slo_vip.use_mtls.no_crl](resources--nfv_service--reference--group-004.md#canonical-2132332003300103-3113011002312213-2031132200002102-0302130103313111-3131111110203310-2011221310312122-2021223021121012-0233230210033201) |
| `https_management.advertise_on_slo_vip.use_mtls.trusted_ca` | [https_management.advertise_on_slo_vip.use_mtls.trusted_ca](resources--nfv_service--reference--group-004.md#canonical-3031030100112003-0123312130332120-0323100223200332-3030021200122113-0301100330022102-2000230112110033-2212000300220000-3103211000022320) |
| `https_management.advertise_on_slo_vip.use_mtls.trusted_ca.name` | [https_management.advertise_on_slo_vip.use_mtls.trusted_ca.name](resources--nfv_service--reference--group-004.md#canonical-1303103011321322-0323212111022331-0022330123211002-2313223001032311-3011001302302133-2331333300310021-2221313131312021-3030300103002333) |
| `https_management.advertise_on_slo_vip.use_mtls.trusted_ca.namespace` | [https_management.advertise_on_slo_vip.use_mtls.trusted_ca.namespace](resources--nfv_service--reference--group-004.md#canonical-3211013020303230-3011013233210033-2320303320023312-1213303231321100-0131013113021300-2130002332112020-3011333130202031-1003321110012002) |
| `https_management.advertise_on_slo_vip.use_mtls.trusted_ca.tenant` | [https_management.advertise_on_slo_vip.use_mtls.trusted_ca.tenant](resources--nfv_service--reference--group-004.md#canonical-1222311001323310-0121022013112311-0013233000120121-2100230230001020-0333223301123233-0221313310332013-2023012013013330-1012222331123021) |
| `https_management.advertise_on_slo_vip.use_mtls.trusted_ca_url` | [https_management.advertise_on_slo_vip.use_mtls.trusted_ca_url](resources--nfv_service--reference--group-004.md#canonical-0133012010033123-2120131301102210-3100311011020003-1313322222022022-2122121310010311-1112131121303110-0200320302330311-1203223223113322) |
| `https_management.advertise_on_slo_vip.use_mtls.xfcc_disabled` | [https_management.advertise_on_slo_vip.use_mtls.xfcc_disabled](resources--nfv_service--reference--group-004.md#canonical-2032322113310033-0223322130231333-2230312000123322-0031311111333312-2212331330110232-0010301130023102-2331310031032011-0103221310211302) |
| `https_management.advertise_on_slo_vip.use_mtls.xfcc_options` | [https_management.advertise_on_slo_vip.use_mtls.xfcc_options](resources--nfv_service--reference--group-004.md#canonical-2333231231323323-2133022210030203-3003200312100312-3300210313131010-2220111223223331-3213223302002320-1123000021321320-1010110330311303) |
| `https_management.advertise_on_slo_vip.use_mtls.xfcc_options.xfcc_header_elements` | [https_management.advertise_on_slo_vip.use_mtls.xfcc_options.xfcc_header_elements](resources--nfv_service--reference--group-004.md#canonical-2233231322021022-3003020210323122-3331331121100030-0211132331301002-0203020201333001-0201010101210222-1110113202032020-2000010000132101) |
| `https_management.default_https_port` | [https_management.default_https_port](resources--nfv_service--reference--group-004.md#canonical-0320211030330020-2221000001121023-2100203330232133-3311230331201001-0322003221233200-2022031033323333-0112023223112111-3022322032020002) |
| `https_management.domain_suffix` | [https_management.domain_suffix](resources--nfv_service--reference--group-002.md#canonical-3131321201111130-3033222200321231-1201202313023231-1333112310130303-0322120131011100-3120231222222200-3012213001122230-2023330133221023) |
| `https_management.https_port` | [https_management.https_port](resources--nfv_service--reference--group-002.md#canonical-0123023032232303-0323211001202133-2013003322001322-1032312221231232-2112222313233021-0230202320231133-3133101222220000-2000332000110001) |
| `id` | [ID](resources--nfv_service--reference--group-001.md#canonical-2202210003230021-2123331330221213-2220013231110122-0233313101232312-2132301010233111-3331221220311002-3100330200012101-3230310003211130) |
| `labels` | [labels](resources--nfv_service--reference--group-001.md#canonical-3123203312302020-1021113033213221-3110100013022111-1202032011011331-1202010201312110-3011233013030200-1102102033300321-0130311301000210) |
| `name` | [name](resources--nfv_service--reference--group-001.md#canonical-0113312013100323-3001203330313203-2130221120123200-2133232231013333-2131013222231223-0001230333223222-1113001322233220-3311213231123230) |
| `namespace` | [namespace](resources--nfv_service--reference--group-001.md#canonical-2132211332310221-0321032100122201-0031303002231002-1333110022202310-2030213001222303-0001030222111003-3232120313210031-2300300123223022) |
| `palo_alto_fw_service` | [palo_alto_fw_service](resources--nfv_service--reference--group-004.md#canonical-1000211303110321-2132033102233220-3202330310311301-1011321010001303-2002102233013002-0110203131021011-1312303320302102-2012032301010010) |
| `palo_alto_fw_service.auto_setup` | [palo_alto_fw_service.auto_setup](resources--nfv_service--reference--group-004.md#canonical-3310302012302032-0132011203023101-1302100100013130-1010111331312131-2232100232203002-1131312233001023-3132321131020132-0022132333300322) |
| `palo_alto_fw_service.auto_setup.admin_password` | [palo_alto_fw_service.auto_setup.admin_password](resources--nfv_service--reference--group-004.md#canonical-2133330101012003-3220110310231233-0121210023200111-1002312202212001-3001330023110033-2330311203202123-3221100300133330-0231000331331011) |
| `palo_alto_fw_service.auto_setup.admin_password.blindfold_secret_info` | [palo_alto_fw_service.auto_setup.admin_password.blindfold_secret_info](resources--nfv_service--reference--group-004.md#canonical-3312321002331313-0321112300220131-1222100300200220-0131201330212003-2312113203330022-3220333111110133-3232102011322122-0001001301230310) |
| `palo_alto_fw_service.auto_setup.admin_password.blindfold_secret_info.decryption_provider` | [palo_alto_fw_service.auto_setup.admin_password.blindfold_secret_info.decryption_provider](resources--nfv_service--reference--group-004.md#canonical-3322323302230030-3023223200122013-1110200033331311-2030321202301232-3300213330133200-1302232202313101-1021112213231303-0200310310311332) |
| `palo_alto_fw_service.auto_setup.admin_password.blindfold_secret_info.location` | [palo_alto_fw_service.auto_setup.admin_password.blindfold_secret_info.location](resources--nfv_service--reference--group-004.md#canonical-3330312333010233-0120002102020002-0212330222003303-1013221130232232-1003222033102032-0132110022030013-1312303102022103-3322310000133003) |
| `palo_alto_fw_service.auto_setup.admin_password.blindfold_secret_info.store_provider` | [palo_alto_fw_service.auto_setup.admin_password.blindfold_secret_info.store_provider](resources--nfv_service--reference--group-004.md#canonical-2023110012031213-1010322130130203-3021220321011331-2120222032202211-3111132321212100-2233132230231333-2211231321211100-2311313302122002) |
| `palo_alto_fw_service.auto_setup.admin_password.clear_secret_info` | [palo_alto_fw_service.auto_setup.admin_password.clear_secret_info](resources--nfv_service--reference--group-004.md#canonical-3230020201102003-2021032330122033-2311223332112133-2200023321023203-3201300223320020-2201113113022210-0313223121122011-2030300031323033) |
| `palo_alto_fw_service.auto_setup.admin_password.clear_secret_info.provider_ref` | [palo_alto_fw_service.auto_setup.admin_password.clear_secret_info.provider_ref](resources--nfv_service--reference--group-004.md#canonical-3211013132311310-1000323123220322-3313231010331110-3330310111113311-2223310303022001-2213102321110222-2012101231313301-1222230211131132) |
| `palo_alto_fw_service.auto_setup.admin_password.clear_secret_info.url` | [palo_alto_fw_service.auto_setup.admin_password.clear_secret_info.url](resources--nfv_service--reference--group-004.md#canonical-0132311130023313-0333011032233031-2212112101110311-1221320102203033-2130300103212203-2003033101103231-3103331122101111-0233121120222032) |
| `palo_alto_fw_service.auto_setup.admin_username` | [palo_alto_fw_service.auto_setup.admin_username](resources--nfv_service--reference--group-004.md#canonical-2000103103321101-3101320102231022-0310121132320223-2021323230002232-1301322031210000-3013100302223002-2032310312033030-2010023002232300) |
| `palo_alto_fw_service.auto_setup.manual_ssh_keys` | [palo_alto_fw_service.auto_setup.manual_ssh_keys](resources--nfv_service--reference--group-004.md#canonical-0233122221112122-2121000312332323-0210133332122101-3131322212331211-3322102301130031-1333313202212332-1001210012012320-1022303302302113) |
| `palo_alto_fw_service.auto_setup.manual_ssh_keys.private_key` | [palo_alto_fw_service.auto_setup.manual_ssh_keys.private_key](resources--nfv_service--reference--group-004.md#canonical-2332213020122031-2210303020333031-0231323012200101-1013100333121130-2123022022301132-3312311231223300-0122213032331233-0221101031102201) |
| `palo_alto_fw_service.auto_setup.manual_ssh_keys.private_key.blindfold_secret_info` | [palo_alto_fw_service.auto_setup.manual_ssh_keys.private_key.blindfold_secret_info](resources--nfv_service--reference--group-004.md#canonical-1223221033213103-0320031310222133-3233330101313111-1023120120332102-1320303212210223-3100031013232012-1332230300100003-0021230222311211) |
| `palo_alto_fw_service.auto_setup.manual_ssh_keys.private_key.blindfold_secret_info.decryption_provider` | [palo_alto_fw_service.auto_setup.manual_ssh_keys.private_key.blindfold_secret_info.decryption_provider](resources--nfv_service--reference--group-004.md#canonical-1013230101201333-0331303122032312-3322011202230123-1222222313122133-3133222303231021-3322033013320033-3313233110100202-0120010203033133) |
| `palo_alto_fw_service.auto_setup.manual_ssh_keys.private_key.blindfold_secret_info.location` | [palo_alto_fw_service.auto_setup.manual_ssh_keys.private_key.blindfold_secret_info.location](resources--nfv_service--reference--group-004.md#canonical-1310011200102123-0222123131313103-1131321311322302-3330213023031130-2123112233311001-3320002231031221-2310312113000101-1110330201100330) |
| `palo_alto_fw_service.auto_setup.manual_ssh_keys.private_key.blindfold_secret_info.store_provider` | [palo_alto_fw_service.auto_setup.manual_ssh_keys.private_key.blindfold_secret_info.store_provider](resources--nfv_service--reference--group-004.md#canonical-3003230130321023-2030030102121000-2230200133001202-1330131211233213-3022300310130120-2221012131233232-1312200310030010-3102023013100330) |
| `palo_alto_fw_service.auto_setup.manual_ssh_keys.private_key.clear_secret_info` | [palo_alto_fw_service.auto_setup.manual_ssh_keys.private_key.clear_secret_info](resources--nfv_service--reference--group-004.md#canonical-3112020301013232-0033122013032003-2002011010110333-2010120211101130-0132023022030032-2310322323012211-1022031130333130-1221222333230033) |
| `palo_alto_fw_service.auto_setup.manual_ssh_keys.private_key.clear_secret_info.provider_ref` | [palo_alto_fw_service.auto_setup.manual_ssh_keys.private_key.clear_secret_info.provider_ref](resources--nfv_service--reference--group-004.md#canonical-1022223202331020-3101321032221302-1000221220113003-2110122311233131-2103212010001222-1231303322332112-3313032330001100-3333113222203102) |
| `palo_alto_fw_service.auto_setup.manual_ssh_keys.private_key.clear_secret_info.url` | [palo_alto_fw_service.auto_setup.manual_ssh_keys.private_key.clear_secret_info.url](resources--nfv_service--reference--group-004.md#canonical-3032102313230201-2212220322013032-0312113122003212-1213302120031203-2032320110312132-1002121301220101-1002002020333123-3120303100133223) |
| `palo_alto_fw_service.auto_setup.manual_ssh_keys.public_key` | [palo_alto_fw_service.auto_setup.manual_ssh_keys.public_key](resources--nfv_service--reference--group-004.md#canonical-3131111213120100-3221233311012020-3203302111201311-0030020012213132-2323311130232003-0231031010212130-3312322203310221-1322032301133321) |
| `palo_alto_fw_service.aws_tgw_site` | [palo_alto_fw_service.aws_tgw_site](resources--nfv_service--reference--group-004.md#canonical-0012332321113213-2222123132313211-3032030210131202-2223301310111330-2321130132211231-2323012213123330-2132001100232011-1310003010212222) |
| `palo_alto_fw_service.aws_tgw_site.name` | [palo_alto_fw_service.aws_tgw_site.name](resources--nfv_service--reference--group-004.md#canonical-1230133202013231-1200313212023030-2012011113113222-0231101010023130-0121113223230231-1322301130323101-3020002200231331-1132302332033311) |
| `palo_alto_fw_service.aws_tgw_site.namespace` | [palo_alto_fw_service.aws_tgw_site.namespace](resources--nfv_service--reference--group-004.md#canonical-3122310033303300-2130021021220211-3331120031233000-3231023212221200-2210213110100121-1032121012110103-2221100123231321-0101002330032112) |
| `palo_alto_fw_service.aws_tgw_site.tenant` | [palo_alto_fw_service.aws_tgw_site.tenant](resources--nfv_service--reference--group-004.md#canonical-2122313220212011-1122203033130303-3001030133311020-3032212213333113-2323011312302010-2310122003122122-3022313123300033-3101112313021121) |
| `palo_alto_fw_service.disable_panaroma` | [palo_alto_fw_service.disable_panaroma](resources--nfv_service--reference--group-004.md#canonical-3002013203212303-1231321022023203-1113203102303113-2220333121002030-3102123003003201-3102033131320002-0302310223121031-3000223212121300) |
| `palo_alto_fw_service.instance_type` | [palo_alto_fw_service.instance_type](resources--nfv_service--reference--group-004.md#canonical-2213122313101232-3203212022110121-3113222323130102-1332110123031000-1321203232212032-3010311012231201-3130032232010011-1003302030223131) |
| `palo_alto_fw_service.pan_ami_bundle1` | [palo_alto_fw_service.pan_ami_bundle1](resources--nfv_service--reference--group-004.md#canonical-1322111223223220-1200112101001320-3213321113212023-2111020123103132-0110211002031300-0321112200201223-0003223230011200-3113033102023213) |
| `palo_alto_fw_service.pan_ami_bundle2` | [palo_alto_fw_service.pan_ami_bundle2](resources--nfv_service--reference--group-004.md#canonical-1122011022100102-3223030111002132-0211100321303211-0212120111031321-2331113023322213-0002303202323120-1012022233331310-3123312112302312) |
| `palo_alto_fw_service.panorama_server` | [palo_alto_fw_service.panorama_server](resources--nfv_service--reference--group-004.md#canonical-3103121332133303-2123033121010032-2301130103230033-3210111033023111-1130030220000033-2321220010112202-1313313032031300-1012210333022122) |
| `palo_alto_fw_service.panorama_server.authorization_key` | [palo_alto_fw_service.panorama_server.authorization_key](resources--nfv_service--reference--group-004.md#canonical-1030003121212220-1203003232200103-1332300033221223-2312320130233133-3010010111213111-2100321003110122-0121111202200313-2303032003221202) |
| `palo_alto_fw_service.panorama_server.authorization_key.blindfold_secret_info` | [palo_alto_fw_service.panorama_server.authorization_key.blindfold_secret_info](resources--nfv_service--reference--group-004.md#canonical-1210022023113033-0330333212220010-1303232121201320-3313102123300333-3300311333100330-0233201110011101-0203323210123121-0320110202313021) |
| `palo_alto_fw_service.panorama_server.authorization_key.blindfold_secret_info.decryption_provider` | [palo_alto_fw_service.panorama_server.authorization_key.blindfold_secret_info.decryption_provider](resources--nfv_service--reference--group-004.md#canonical-2113011001033121-2132111311022122-3322102330100110-3312033220120032-1231331123323100-3103000223122120-0000320323010233-0101012331321221) |
| `palo_alto_fw_service.panorama_server.authorization_key.blindfold_secret_info.location` | [palo_alto_fw_service.panorama_server.authorization_key.blindfold_secret_info.location](resources--nfv_service--reference--group-004.md#canonical-2001101231131023-3031022011203101-3122231120100302-3003030110100133-1320023000103030-3303021003220032-1203123021122320-1320030230200310) |
| `palo_alto_fw_service.panorama_server.authorization_key.blindfold_secret_info.store_provider` | [palo_alto_fw_service.panorama_server.authorization_key.blindfold_secret_info.store_provider](resources--nfv_service--reference--group-004.md#canonical-2210103213020213-1221213112110312-1101131223123323-0022012330012033-2203212100021112-0210103123322123-2120012231230321-0211232202223210) |
| `palo_alto_fw_service.panorama_server.authorization_key.clear_secret_info` | [palo_alto_fw_service.panorama_server.authorization_key.clear_secret_info](resources--nfv_service--reference--group-004.md#canonical-0232201002030230-0011311321110321-0030011010121223-3311000203000223-1303233320122233-3131022210012030-1023010221321311-0330002122031300) |
| `palo_alto_fw_service.panorama_server.authorization_key.clear_secret_info.provider_ref` | [palo_alto_fw_service.panorama_server.authorization_key.clear_secret_info.provider_ref](resources--nfv_service--reference--group-004.md#canonical-2013233103320313-2231202111101332-3310023303011232-0102033032121222-2013031102202030-0310201133120110-0322111302021120-1011330120223233) |
| `palo_alto_fw_service.panorama_server.authorization_key.clear_secret_info.url` | [palo_alto_fw_service.panorama_server.authorization_key.clear_secret_info.url](resources--nfv_service--reference--group-004.md#canonical-0010212223120021-0301133330003002-0202300011200312-2121301213323210-1233102121213031-2002203123210033-3023000213330302-2213030233013110) |
| `palo_alto_fw_service.panorama_server.device_group_name` | [palo_alto_fw_service.panorama_server.device_group_name](resources--nfv_service--reference--group-004.md#canonical-3103003113222211-3320101331021320-1201131103103013-0103330000212130-2321233113202033-2130221312310210-2002233130210120-0312032133001302) |
| `palo_alto_fw_service.panorama_server.server` | [palo_alto_fw_service.panorama_server.server](resources--nfv_service--reference--group-004.md#canonical-1232210103223223-3200213021033113-0012022103122311-3322032322212330-0311132201233101-3322203230212210-0212320120313231-3131321211020320) |
| `palo_alto_fw_service.panorama_server.template_stack_name` | [palo_alto_fw_service.panorama_server.template_stack_name](resources--nfv_service--reference--group-004.md#canonical-0023310330302002-0230322131303210-3002331023223211-3122231321010002-1021321000121001-3233031321110112-1102131221101031-2022133030233130) |
| `palo_alto_fw_service.service_nodes` | [palo_alto_fw_service.service_nodes](resources--nfv_service--reference--group-004.md#canonical-1302321210221003-3323232100312333-0233111302033230-0212301310130300-1122133130031320-3202102310221101-2302120203200030-3220011203130022) |
| `palo_alto_fw_service.service_nodes.nodes` | [palo_alto_fw_service.service_nodes.nodes](resources--nfv_service--reference--group-004.md#canonical-2210110321103121-1011222232120021-3130222320021302-3222211233202210-3133330213101103-1032200120002023-0000001131032233-2120301200330233) |
| `palo_alto_fw_service.service_nodes.nodes.aws_az_name` | [palo_alto_fw_service.service_nodes.nodes.aws_az_name](resources--nfv_service--reference--group-004.md#canonical-0310331101023321-1020203311312220-2302333233311323-1200310302321103-3033111000201032-1300123111222023-1121102231211320-2023032232113021) |
| `palo_alto_fw_service.service_nodes.nodes.mgmt_subnet` | [palo_alto_fw_service.service_nodes.nodes.mgmt_subnet](resources--nfv_service--reference--group-004.md#canonical-3300303212211100-3223222323322012-2221330302221020-1010123222300000-2103132230233001-2303221320100120-3223300103122320-2032122021331300) |
| `palo_alto_fw_service.service_nodes.nodes.mgmt_subnet.existing_subnet_id` | [palo_alto_fw_service.service_nodes.nodes.mgmt_subnet.existing_subnet_id](resources--nfv_service--reference--group-004.md#canonical-0330302021233022-2203233100301221-3021221103123010-0200022311221300-1121001221013203-0231130111021222-1020020312022033-0313203302233131) |
| `palo_alto_fw_service.service_nodes.nodes.mgmt_subnet.subnet_param` | [palo_alto_fw_service.service_nodes.nodes.mgmt_subnet.subnet_param](resources--nfv_service--reference--group-004.md#canonical-1231310021213332-0211120101313002-2012001021112220-1320120103310100-3030003233311020-3122313321221222-0113210223313010-2220122033200233) |
| `palo_alto_fw_service.service_nodes.nodes.mgmt_subnet.subnet_param.ipv4` | [palo_alto_fw_service.service_nodes.nodes.mgmt_subnet.subnet_param.ipv4](resources--nfv_service--reference--group-004.md#canonical-2231301202103133-2302103323312212-2230003210333200-3310001200132120-1122102101133210-1203120333111320-0213111002011020-2033333003311322) |
| `palo_alto_fw_service.service_nodes.nodes.node_name` | [palo_alto_fw_service.service_nodes.nodes.node_name](resources--nfv_service--reference--group-004.md#canonical-2321023313313102-3323210311131300-2212100330310232-0011111301010312-2112222101213213-1122130030200212-2311110011022222-2131130010323213) |
| `palo_alto_fw_service.service_nodes.nodes.reserved_mgmt_subnet` | [palo_alto_fw_service.service_nodes.nodes.reserved_mgmt_subnet](resources--nfv_service--reference--group-004.md#canonical-3213003231222112-2333030130012231-0233132333121220-1202133030113230-1102002001023102-2120121322111033-2032313103103121-3310203032113132) |
| `palo_alto_fw_service.ssh_key` | [palo_alto_fw_service.ssh_key](resources--nfv_service--reference--group-004.md#canonical-0012322131130123-0120023232110303-0000302311232212-2101211230330313-0021311112213233-0111102100133301-3320122002203211-0103111331231221) |
| `palo_alto_fw_service.tags` | [palo_alto_fw_service.tags](resources--nfv_service--reference--group-004.md#canonical-2222200222102312-2213202200002330-2322120223121012-0211010210303330-1331120313221203-3213032211211020-3002123130003131-1110320202331232) |
| `palo_alto_fw_service.version` | [palo_alto_fw_service.version](resources--nfv_service--reference--group-004.md#canonical-3133201303322112-3200103333121012-2123130323122010-0300222021211033-1022111311000133-1210121113303002-1212230303100230-2231000130111310) |
| `timeouts` | [timeouts](resources--nfv_service--reference--group-004.md#canonical-2232213200300031-0231011230212012-2030223133101200-3212212323011330-2132013202023300-3200300330113012-3333321303101212-3332213233102122) |
| `timeouts.create` | [timeouts.create](resources--nfv_service--reference--group-004.md#canonical-3222232213020000-2113233003111233-0132003233332011-1003013312011030-2000121310103022-1331033313331031-2330020320223021-2100322221223012) |
| `timeouts.delete` | [timeouts.delete](resources--nfv_service--reference--group-004.md#canonical-2000013102223211-2110323122200231-3332122233130011-1113222222023000-0120010230202233-2331332033022313-2230303023200321-3133313220110321) |
| `timeouts.read` | [timeouts.read](resources--nfv_service--reference--group-004.md#canonical-1303103300312222-2332033231033211-0011323303023333-3230002311302303-3001110303202133-1201113231302103-0022123020301201-2122022232232323) |
| `timeouts.update` | [timeouts.update](resources--nfv_service--reference--group-004.md#canonical-2232323323112011-2100000222330312-2220130221023302-0101231132103333-3022311232331110-2230112110031201-3023120213302111-3023232122133100) |

<a id="canonical-1213021122302111-2303132302331132-2033120022132312-0232210233303323-2001032101312213-3131012003332023-0101021121323031-2100213002233100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `disable_https_management` properties

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-0223321113210321-0013210302112221-0313201132021001-1102330023301310-1131322132000230-3303202123330303-3122032322302202-2202303201230222)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-1033200313012130-2303320020000220-3032203302003212-3030323031320111-0322032022323123-0203102131300133-1232321110231101-1022222010201320)
- disable_https_management

<a id="canonical-3112312130013010-0213122301312322-1233230202121013-2232313130212030-0022130101301033-0101103110230133-2130301211211102-1301123223121110"></a>

Type: `["object", {}]`. Optional.

\[OneOf: disable\_https\_management, https\_management; Default: disable\_https\_management\]
Configuration parameter for disable https management.

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

OneOf alternatives in this subsection:

- [disable_https_management](resources--nfv_service--reference--group-001.md#canonical-3112312130013010-0213122301312322-1233230202121013-2232313130212030-0022130101301033-0101103110230133-2130301211211102-1301123223121110)
- [https_management](resources--nfv_service--reference--group-002.md#canonical-2310110120210321-0011122021130100-1020101231210333-1001101201002320-3213323313112230-0312000102310210-1113012113122222-1223202310021130)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
disable_https_management = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1121230113200220-3301032102311323-1313313310300012-0131301302121233-3323233302031130-1303122223201233-3332210000021300-3200231130103032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `disable_ssh_access` properties

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-0223321113210321-0013210302112221-0313201132021001-1102330023301310-1131322132000230-3303202123330303-3122032322302202-2202303201230222)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-1033200313012130-2303320020000220-3032203302003212-3030323031320111-0322032022323123-0203102131300133-1232321110231101-1022222010201320)
- disable_ssh_access

<a id="canonical-0111131333123130-2302200131001101-0111200312022021-3302110232120331-0010113121223001-2211222211321130-2020322122013323-2213211231011123"></a>

Type: `["object", {}]`. Optional.

\[OneOf: disable\_ssh\_access, enabled\_ssh\_access; Default: disable\_ssh\_access\] Configuration
parameter for disable SSH access.

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

OneOf alternatives in this subsection:

- [disable_ssh_access](resources--nfv_service--reference--group-001.md#canonical-0111131333123130-2302200131001101-0111200312022021-3302110232120331-0010113121223001-2211222211321130-2020322122013323-2213211231011123)
- [enabled_ssh_access](resources--nfv_service--reference--group-001.md#canonical-0201302210200021-1010011310120203-1323221101222200-0312020201221132-0012233110231012-0231010102020210-3032213233301121-3323230130222103)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
disable_ssh_access = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1233323310221201-2110101010301322-3301231231110303-1003102112121320-1012103233232011-0000232331110323-2202130110221200-1133311203200010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `enabled_ssh_access` properties

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-0223321113210321-0013210302112221-0313201132021001-1102330023301310-1131322132000230-3303202123330303-3122032322302202-2202303201230222)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-1033200313012130-2303320020000220-3032203302003212-3030323031320111-0322032022323123-0203102131300133-1232321110231101-1022222010201320)
- enabled_ssh_access

<a id="canonical-0201302210200021-1010011310120203-1323221101222200-0312020201221132-0012233110231012-0231010102020210-3032213233301121-3323230130222103"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for enabled SSH access.

Additional upstream details:

SSH based configuration.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-advertise_choice": "[\"advertise_on_sli\",\"advertise_on_slo\",\"advertise_on_slo_sli\"]"
}
```

Terraform syntax:

```terraform
enabled_ssh_access {
  # Configure direct properties listed below.
}
```

<a id="canonical-1313013133021212-3330102103202133-2001002122101002-1321103300112233-2001123332210310-1301013011332232-3131312121011200-1200300130320021"></a>

### Direct properties for `enabled_ssh_access`

- [advertise_on_sli](resources--nfv_service--reference--group-001.md#canonical-0101102232222113-3100331132213111-3122201012102230-2202321122003010-0213100122100312-3013032200113102-3203020232232331-1112200321321030): complete subsection reference.

- [advertise_on_slo](resources--nfv_service--reference--group-001.md#canonical-2313122330330301-1103203312101213-2021113003033113-3132003130032233-2232202322123113-1100211322010032-0033123213103232-1320011203300200): complete subsection reference.

- [advertise_on_slo_sli](resources--nfv_service--reference--group-001.md#canonical-1330031012210003-2220302330022203-3123300332223001-0102312202232013-1111022003030130-0013021223322033-3013230111113303-0003330222120032): complete subsection reference.

<a id="canonical-2020201313303300-1010202033210223-2101120030023332-1023230302232323-3300231213122233-1000113021111020-0111233220213221-2210102333012311"></a>

<a id="canonical-1013010032003313-3233100200301033-2112010122230320-1131010002102102-0111311110131201-1233303322303320-2103221010121020-2032001200132003"></a>

#### `enabled_ssh_access.domain_suffix` property

Type: `"string"`. Optional.

Domain suffix will be used along with node name to form the hostname for SSH node management.

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
    "ves.io.schema.rules.string.hostname": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.hostname": "true"
  }
}
```

- [node_ssh_ports](resources--nfv_service--reference--group-001.md#canonical-2320112001132030-0303302223332210-2023201233202012-2031021302200221-0121132101211303-2310102103021322-3032220021132123-3120320302212102): complete subsection reference.

<a id="canonical-0101102232222113-3100331132213111-3122201012102230-2202321122003010-0213100122100312-3013032200113102-3203020232232331-1112200321321030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `enabled_ssh_access.advertise_on_sli` properties

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-0223321113210321-0013210302112221-0313201132021001-1102330023301310-1131322132000230-3303202123330303-3122032322302202-2202303201230222)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-1033200313012130-2303320020000220-3032203302003212-3030323031320111-0322032022323123-0203102131300133-1232321110231101-1022222010201320)
- [enabled_ssh_access](resources--nfv_service--reference--group-001.md#canonical-1233323310221201-2110101010301322-3301231231110303-1003102112121320-1012103233232011-0000232331110323-2202130110221200-1133311203200010)
- enabled_ssh_access.advertise_on_sli

<a id="canonical-0003013100101131-1122201201230222-1001220013030203-3103121133102032-2013021223011301-0010320102213130-1230031332301211-0111323031201032"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for advertise on sli.

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
advertise_on_sli = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2313122330330301-1103203312101213-2021113003033113-3132003130032233-2232202322123113-1100211322010032-0033123213103232-1320011203300200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `enabled_ssh_access.advertise_on_slo` properties

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-0223321113210321-0013210302112221-0313201132021001-1102330023301310-1131322132000230-3303202123330303-3122032322302202-2202303201230222)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-1033200313012130-2303320020000220-3032203302003212-3030323031320111-0322032022323123-0203102131300133-1232321110231101-1022222010201320)
- [enabled_ssh_access](resources--nfv_service--reference--group-001.md#canonical-1233323310221201-2110101010301322-3301231231110303-1003102112121320-1012103233232011-0000232331110323-2202130110221200-1133311203200010)
- enabled_ssh_access.advertise_on_slo

<a id="canonical-3113030011011303-1203312332110201-1123123333313002-2220022010321112-3132022210232301-2322311232111112-0120331201000303-0302121303030110"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for advertise on slo.

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
advertise_on_slo = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1330031012210003-2220302330022203-3123300332223001-0102312202232013-1111022003030130-0013021223322033-3013230111113303-0003330222120032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `enabled_ssh_access.advertise_on_slo_sli` properties

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-0223321113210321-0013210302112221-0313201132021001-1102330023301310-1131322132000230-3303202123330303-3122032322302202-2202303201230222)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-1033200313012130-2303320020000220-3032203302003212-3030323031320111-0322032022323123-0203102131300133-1232321110231101-1022222010201320)
- [enabled_ssh_access](resources--nfv_service--reference--group-001.md#canonical-1233323310221201-2110101010301322-3301231231110303-1003102112121320-1012103233232011-0000232331110323-2202130110221200-1133311203200010)
- enabled_ssh_access.advertise_on_slo_sli

<a id="canonical-2301123112320330-1333002212323201-0030013311233203-1100132101100002-2003321311300011-1110200320302232-0231130202120003-1231233022300232"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for advertise on slo sli.

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
advertise_on_slo_sli = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2320112001132030-0303302223332210-2023201233202012-2031021302200221-0121132101211303-2310102103021322-3032220021132123-3120320302212102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `enabled_ssh_access.node_ssh_ports` properties

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-0223321113210321-0013210302112221-0313201132021001-1102330023301310-1131322132000230-3303202123330303-3122032322302202-2202303201230222)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-1033200313012130-2303320020000220-3032203302003212-3030323031320111-0322032022323123-0203102131300133-1232321110231101-1022222010201320)
- [enabled_ssh_access](resources--nfv_service--reference--group-001.md#canonical-1233323310221201-2110101010301322-3301231231110303-1003102112121320-1012103233232011-0000232331110323-2202130110221200-1133311203200010)
- enabled_ssh_access.node_ssh_ports

<a id="canonical-3131212313033221-1031232020033032-2212111001221102-1233323330113003-3201322203113012-2002120102101320-1230123231121003-0222202133311103"></a>

Type: `"object"`. list nested block, Optional.

Management Node SSH Port. Enter TCP port and node name per node.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 2,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 2,
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
    "ves.io.schema.rules.repeated.max_items": "2"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "2"
  }
}
```

Terraform syntax:

```terraform
node_ssh_ports {
  # Configure direct properties listed below.
}
```

<a id="canonical-2033331121231211-0102020132132231-0000021001330233-1302222231122111-3030303102130303-2201101222111220-2210221021320213-2130311121011213"></a>

### Direct properties for `enabled_ssh_access.node_ssh_ports`

<a id="canonical-0311023330222223-2231312201023331-1310102133100122-3323001122303210-2221201003011301-0013011013323001-0331130103231201-1302311300110220"></a>

#### `enabled_ssh_access.node_ssh_ports.node_name` property

Type: `"string"`. Optional.

Node name will be used to match a particular node with the desired TCP port.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-09T12:34:59+00:00"
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
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-3213021100333102-0320320221303020-2011100001021130-3322000022023233-0103300220122031-1221211101011311-0311030210021112-1130322030113320"></a>

<a id="canonical-3303222312332101-1202220320220201-1231013130002003-0233231012201220-2012301020220220-2212230212231002-1122223131312130-3223310011121000"></a>

#### `enabled_ssh_access.node_ssh_ports.ssh_port` property

Type: `"number"`. Optional.

SSH Port. Enter TCP port per node.

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
    "minimum": 1024
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1024",
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1024",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-2203212022212003-2212301132213201-2010231030110211-3120210322102100-3130230103331230-0332101002301003-2302320031110011-3121121001302210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `f5_big_ip_aws_service` properties

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-0223321113210321-0013210302112221-0313201132021001-1102330023301310-1131322132000230-3303202123330303-3122032322302202-2202303201230222)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-1033200313012130-2303320020000220-3032203302003212-3030323031320111-0322032022323123-0203102131300133-1232321110231101-1022222010201320)
- f5_big_ip_aws_service

<a id="canonical-0201100230223303-2001123313212330-0023112220313003-3331311211332102-0202122111230001-2232111322331013-1010031220331333-3331233032322321"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: f5\_big\_ip\_aws\_service, palo\_alto\_fw\_service\] Virtual BIG-IP AWS. Virtual BIG-IP
specification for AWS.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-image_choice": "[\"market_place_image\"]",
  "x-ves-oneof-field-site_type_choice": "[\"aws_tgw_site_params\"]"
}
```

OneOf alternatives in this subsection:

- [f5_big_ip_aws_service](resources--nfv_service--reference--group-001.md#canonical-0201100230223303-2001123313212330-0023112220313003-3331311211332102-0202122111230001-2232111322331013-1010031220331333-3331233032322321)
- [palo_alto_fw_service](resources--nfv_service--reference--group-004.md#canonical-1000211303110321-2132033102233220-3202330310311301-1011321010001303-2002102233013002-0110203131021011-1312303320302102-2012032301010010)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
f5_big_ip_aws_service {
  # Configure direct properties listed below.
}
```

<a id="canonical-1130200220213113-1110101312012330-3213323122310332-3011233211031221-1112031230320230-3203310233111131-3322321333022202-1303122230233330"></a>

### Direct properties for `f5_big_ip_aws_service`

- [admin_password](resources--nfv_service--reference--group-001.md#canonical-0010020132110202-2302301000230121-0033122013301323-3131210223203001-1001122123130133-0302231311112103-2111303003131303-1322011213213130): complete subsection reference.

<a id="canonical-0303303210230013-0010231111233233-3000013023023101-1112213102230002-2303033320012311-3320331223211221-3021300202321023-3012121221222220"></a>

<a id="canonical-0103310013010323-0300022100231233-0321222222220023-1013000212202023-3122210101103202-3210133130112330-2030112222122321-1303313020111032"></a>

#### `f5_big_ip_aws_service.admin_username` property

Type: `"string"`. Optional.

Admin Username. Admin Username for BIG-IP.

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
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

- [aws_tgw_site_params](resources--nfv_service--reference--group-001.md#canonical-2233311222002000-2100133010320202-2122321003331003-0221303030313102-0120212003100213-2113310222000310-3130123003223212-1021001222122210): complete subsection reference.

- [endpoint_service](resources--nfv_service--reference--group-001.md#canonical-1010032200210212-2313121332112310-0110323222301213-2211012302000323-2311222030300112-0122023013313002-3223113221233302-2201323331331100): complete subsection reference.

- [market_place_image](resources--nfv_service--reference--group-002.md#canonical-3121221100302221-2312112212213100-2300030203321033-0023133210013231-0322231123223011-1213333100111231-1001202332212022-3001113302213202): complete subsection reference.

- [nodes](resources--nfv_service--reference--group-002.md#canonical-0303213133011232-0201322132022131-1321023330330310-0032213003120332-3011101310033312-2030333102020323-2111032131310203-3332013322121201): complete subsection reference.

<a id="canonical-2122322100020010-0233231112003322-1100330013212030-1031310200022000-0023300112302121-3230133231122000-3021020020303103-3002313323100011"></a>

<a id="canonical-3103121233211222-2101132230330130-1011001001232302-2003200321231220-2100121032033030-0231011103231232-2230000002202132-3302210002100331"></a>

#### `f5_big_ip_aws_service.ssh_key` property

Type: `"string"`. Optional.

Public SSH key for accessing the Big IP nodes.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-2300122213223130-2322121232230202-1121102110030323-0233022121232321-3120330012222113-2201303223211330-2200201322003010-1322232110123021"></a>

<a id="canonical-1121202133330123-3001033001122211-0132100112102132-2331001022103201-0132021002001200-0111302332201002-1202220332003321-0201233001003032"></a>

#### `f5_big_ip_aws_service.tags` property

Type: `["map", "string"]`. Optional.

AWS Tags is a label consisting of a user-defined key and value. It helps to manage, identify,
organize, search for, and filter resources in AWS console.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "cardinality": {
      "maxProperties": 40
    },
    "category": "discovery",
    "constraintType": "map",
    "deterministic": true,
    "keys": {
      "maxLength": 127,
      "type": "string"
    },
    "originalRules": {
      "ves.io.schema.rules.map.keys.string.max_len": "127",
      "ves.io.schema.rules.map.max_pairs": "40",
      "ves.io.schema.rules.map.values.string.max_len": "255"
    },
    "values": {
      "maxLength": 255,
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

<a id="canonical-0010020132110202-2302301000230121-0033122013301323-3131210223203001-1001122123130133-0302231311112103-2111303003131303-1322011213213130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `f5_big_ip_aws_service.admin_password` properties

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-0223321113210321-0013210302112221-0313201132021001-1102330023301310-1131322132000230-3303202123330303-3122032322302202-2202303201230222)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-1033200313012130-2303320020000220-3032203302003212-3030323031320111-0322032022323123-0203102131300133-1232321110231101-1022222010201320)
- [f5_big_ip_aws_service](resources--nfv_service--reference--group-001.md#canonical-2203212022212003-2212301132213201-2010231030110211-3120210322102100-3130230103331230-0332101002301003-2302320031110011-3121121001302210)
- f5_big_ip_aws_service.admin_password

<a id="canonical-0020332213331212-2101220311133313-2122212010123301-0233103012221122-3212102010332130-3113012033002232-2111100123120310-1023112203303112"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
admin_password {
  # Configure direct properties listed below.
}
```

<a id="canonical-0130322223020100-0221323203000211-2021322100031210-1301303120211132-1210121223333313-1022202001001002-0202002120031123-2121023201033022"></a>

### Direct properties for `f5_big_ip_aws_service.admin_password`

- [blindfold_secret_info](resources--nfv_service--reference--group-001.md#canonical-2331000332201331-3301321303213111-3201321310302320-0131002130303003-3101210121103302-3213322231203321-3311212012200011-0112013303012213): complete subsection reference.

- [clear_secret_info](resources--nfv_service--reference--group-001.md#canonical-2220212232310031-3023101212201003-3322202012303003-1302113031202022-0332031222230332-1302010002013330-0233332012031020-1302011101010002): complete subsection reference.

<a id="canonical-2331000332201331-3301321303213111-3201321310302320-0131002130303003-3101210121103302-3213322231203321-3311212012200011-0112013303012213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `f5_big_ip_aws_service.admin_password.blindfold_secret_info` properties

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-0223321113210321-0013210302112221-0313201132021001-1102330023301310-1131322132000230-3303202123330303-3122032322302202-2202303201230222)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-1033200313012130-2303320020000220-3032203302003212-3030323031320111-0322032022323123-0203102131300133-1232321110231101-1022222010201320)
- [f5_big_ip_aws_service](resources--nfv_service--reference--group-001.md#canonical-2203212022212003-2212301132213201-2010231030110211-3120210322102100-3130230103331230-0332101002301003-2302320031110011-3121121001302210)
- [f5_big_ip_aws_service.admin_password](resources--nfv_service--reference--group-001.md#canonical-0010020132110202-2302301000230121-0033122013301323-3131210223203001-1001122123130133-0302231311112103-2111303003131303-1322011213213130)
- f5_big_ip_aws_service.admin_password.blindfold_secret_info

<a id="canonical-3301022230121212-1231300113010201-0231103320022010-2213102330001131-3120203022202301-0012023232321301-3331313021332130-3220130023020300"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
blindfold_secret_info {
  # Configure direct properties listed below.
}
```

<a id="canonical-3102321331123020-2031030311230020-1133032101100222-0033201103222032-1232131232130100-3232123023123223-2323321212321132-0230103232132310"></a>

### Direct properties for `f5_big_ip_aws_service.admin_password.blindfold_secret_info`

<a id="canonical-2321300332103030-2323111232301113-2202003220003121-0302121112122102-0113131322320101-0022103312220333-1230113231302133-0220203300110330"></a>

#### `f5_big_ip_aws_service.admin_password.blindfold_secret_info.decryption_provider` property

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

<a id="canonical-3311303232101323-3130320010021320-3332011001012212-0131000323200013-1220100031022101-2230010020130113-3003203203202220-2320230320323001"></a>

<a id="canonical-2031201003300212-2230330122332022-3231011021310101-1012103010222303-3220302233122301-3223110023332201-0021313203000333-3300000300210233"></a>

#### `f5_big_ip_aws_service.admin_password.blindfold_secret_info.location` property

Type: `"string"`. Optional, Sensitive.

Location is the URI\_ref. It could be in URL format for string:/// Or it could be a path if the
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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-2310103312222302-3132133201331103-1321123300002201-3212313320121213-3023023111211220-1201220230110221-3303233122333130-3000100322001131"></a>

<a id="canonical-2020311101032030-3130021013221311-3311130221013122-1112102333230031-2000212230211133-3330101300022223-0020101330320231-2211023222331332"></a>

#### `f5_big_ip_aws_service.admin_password.blindfold_secret_info.store_provider` property

Type: `"string"`. Optional.

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

<a id="canonical-2220212232310031-3023101212201003-3322202012303003-1302113031202022-0332031222230332-1302010002013330-0233332012031020-1302011101010002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `f5_big_ip_aws_service.admin_password.clear_secret_info` properties

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-0223321113210321-0013210302112221-0313201132021001-1102330023301310-1131322132000230-3303202123330303-3122032322302202-2202303201230222)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-1033200313012130-2303320020000220-3032203302003212-3030323031320111-0322032022323123-0203102131300133-1232321110231101-1022222010201320)
- [f5_big_ip_aws_service](resources--nfv_service--reference--group-001.md#canonical-2203212022212003-2212301132213201-2010231030110211-3120210322102100-3130230103331230-0332101002301003-2302320031110011-3121121001302210)
- [f5_big_ip_aws_service.admin_password](resources--nfv_service--reference--group-001.md#canonical-0010020132110202-2302301000230121-0033122013301323-3131210223203001-1001122123130133-0302231311112103-2111303003131303-1322011213213130)
- f5_big_ip_aws_service.admin_password.clear_secret_info

<a id="canonical-3313101212220110-0010011312222201-2232103102110313-2330321322322120-0123333232120003-0102122233312221-1133203201120210-3203132203121302"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
clear_secret_info {
  # Configure direct properties listed below.
}
```

<a id="canonical-1332230232320321-2302301022003320-0230001310031133-1303210300100210-1101200001212223-1200110020230112-0033122230033201-3122301230000200"></a>

### Direct properties for `f5_big_ip_aws_service.admin_password.clear_secret_info`

<a id="canonical-2323202200030121-1000333130332100-0122030331233120-1131323222303220-3322321130010322-3102332223302030-3132011203313312-1303202220231201"></a>

#### `f5_big_ip_aws_service.admin_password.clear_secret_info.provider_ref` property

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-1033213333031230-3023320130220111-0210002233210013-0132023210312332-2131031103222220-0133022322130021-0013212302020120-0011333001110300"></a>

<a id="canonical-0021201111330012-2311202033132131-1131103000132311-3231023320301113-3020232012320301-2012021033320100-3023021202101303-2300103110320211"></a>

#### `f5_big_ip_aws_service.admin_password.clear_secret_info.url` property

Type: `"string"`. Optional, Sensitive.

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded base64 format. When asked for this secret, caller will GET Secret bytes after
base64 decoding.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-2233311222002000-2100133010320202-2122321003331003-0221303030313102-0120212003100213-2113310222000310-3130123003223212-1021001222122210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `f5_big_ip_aws_service.aws_tgw_site_params` properties

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-0223321113210321-0013210302112221-0313201132021001-1102330023301310-1131322132000230-3303202123330303-3122032322302202-2202303201230222)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-1033200313012130-2303320020000220-3032203302003212-3030323031320111-0322032022323123-0203102131300133-1232321110231101-1022222010201320)
- [f5_big_ip_aws_service](resources--nfv_service--reference--group-001.md#canonical-2203212022212003-2212301132213201-2010231030110211-3120210322102100-3130230103331230-0332101002301003-2302320031110011-3121121001302210)
- f5_big_ip_aws_service.aws_tgw_site_params

<a id="canonical-3320313103200132-0103010113210101-1101031002101012-2010331023022021-2333210311111022-0030031113000212-3001123013330231-2331031203101333"></a>

Type: `"object"`. single nested block, Optional.

BIG-IP AWS TGW Site. BIG-IP AWS TGW site specification.

Receipt-pinned upstream constraints:

```json
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
aws_tgw_site_params {
  # Configure direct properties listed below.
}
```

<a id="canonical-3232112001333120-3333033230130001-1300220333300103-1102022332233201-1010221330133012-2012331123321013-2333112101113203-0000001113211020"></a>

### Direct properties for `f5_big_ip_aws_service.aws_tgw_site_params`

- [aws_tgw_site](resources--nfv_service--reference--group-001.md#canonical-1300102203120232-0232232122011313-1301020010010310-2332200230120321-1331000100232212-0031313231033012-2221132102223233-3313110210032323): complete subsection reference.

<a id="canonical-1300102203120232-0232232122011313-1301020010010310-2332200230120321-1331000100232212-0031313231033012-2221132102223233-3313110210032323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `f5_big_ip_aws_service.aws_tgw_site_params.aws_tgw_site` properties

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-0223321113210321-0013210302112221-0313201132021001-1102330023301310-1131322132000230-3303202123330303-3122032322302202-2202303201230222)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-1033200313012130-2303320020000220-3032203302003212-3030323031320111-0322032022323123-0203102131300133-1232321110231101-1022222010201320)
- [f5_big_ip_aws_service](resources--nfv_service--reference--group-001.md#canonical-2203212022212003-2212301132213201-2010231030110211-3120210322102100-3130230103331230-0332101002301003-2302320031110011-3121121001302210)
- [f5_big_ip_aws_service.aws_tgw_site_params](resources--nfv_service--reference--group-001.md#canonical-2233311222002000-2100133010320202-2122321003331003-0221303030313102-0120212003100213-2113310222000310-3130123003223212-1021001222122210)
- f5_big_ip_aws_service.aws_tgw_site_params.aws_tgw_site

<a id="canonical-0003032121222113-0011112101002302-2301303213332311-0110221102323033-1033022110331321-1232332133211020-1113312321331210-2112101100020220"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
aws_tgw_site {
  # Configure direct properties listed below.
}
```

<a id="canonical-1111211003321333-1310102200012223-2113111100130222-1323210332010203-1201112323202101-2133010221333220-0101321313131323-3132021221203112"></a>

### Direct properties for `f5_big_ip_aws_service.aws_tgw_site_params.aws_tgw_site`

<a id="canonical-1010331213000212-2221032202201100-1020122332020122-1321312121310021-3223131031131020-2023120022023130-3223301320323103-0030030210112312"></a>

#### `f5_big_ip_aws_service.aws_tgw_site_params.aws_tgw_site.name` property

Type: `"string"`. Optional.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-0311223030320111-2331221013210111-0222200302220102-1212313101302221-0010212231303100-0122210313220220-2022301220113120-1021032320133212"></a>

<a id="canonical-0303302210320030-0001220121013233-0303301030032323-3021212100102033-0200210023121121-3233221123102133-0112100033323212-1312103013032213"></a>

#### `f5_big_ip_aws_service.aws_tgw_site_params.aws_tgw_site.namespace` property

Type: `"string"`. Optional, Computed.

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
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-2130333010312310-2122231331023013-1203010000232020-3020113331310320-3302001002000130-1331132211111103-3033321223133312-2111131332222222"></a>

<a id="canonical-2323232303100321-3223112312013231-3100023132212012-2010230110102233-2001332203103203-2201132201001120-2223022302231110-3112032033200132"></a>

#### `f5_big_ip_aws_service.aws_tgw_site_params.aws_tgw_site.tenant` property

Type: `"string"`. Computed.

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
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-1010032200210212-2313121332112310-0110323222301213-2211012302000323-2311222030300112-0122023013313002-3223113221233302-2201323331331100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `f5_big_ip_aws_service.endpoint_service` properties

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-0223321113210321-0013210302112221-0313201132021001-1102330023301310-1131322132000230-3303202123330303-3122032322302202-2202303201230222)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-1033200313012130-2303320020000220-3032203302003212-3030323031320111-0322032022323123-0203102131300133-1232321110231101-1022222010201320)
- [f5_big_ip_aws_service](resources--nfv_service--reference--group-001.md#canonical-2203212022212003-2212301132213201-2010231030110211-3120210322102100-3130230103331230-0332101002301003-2302320031110011-3121121001302210)
- f5_big_ip_aws_service.endpoint_service

<a id="canonical-1032321201122030-3112023110013020-2120212102023120-1301023013310313-1323131301203220-1231311320221303-3203101100210331-2002323023100113"></a>

Type: `"object"`. single nested block, Optional.

Endpoint Service is a type of NFV service where the packets are destined to NFV and service modifies
the destination with a new destination address.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-external_vip_choice": "[\"advertise_on_slo_ip\",\"advertise_on_slo_ip_external\",\"disable_advertise_on_slo_ip\"]",
  "x-ves-oneof-field-inside_vip_choice": "[\"automatic_vip\",\"configured_vip\"]",
  "x-ves-oneof-field-tcp_port_choice": "[\"custom_tcp_ports\",\"default_tcp_ports\",\"http_port\",\"https_port\",\"no_tcp_ports\"]",
  "x-ves-oneof-field-udp_port_choice": "[\"custom_udp_ports\",\"no_udp_ports\"]"
}
```

Terraform syntax:

```terraform
endpoint_service {
  # Configure direct properties listed below.
}
```

<a id="canonical-0230122032113023-2221313121001212-1200231233231233-2031210332203322-3331200122122210-3202110202113313-3031310201320333-1023223022211220"></a>

### Direct properties for `f5_big_ip_aws_service.endpoint_service`

- [advertise_on_slo_ip](resources--nfv_service--reference--group-001.md#canonical-2002023200130011-3221031100321323-1023033331122313-3230213010102032-2121211120330320-2103303103132200-3323232000311320-1022321130223012): complete subsection reference.

- [advertise_on_slo_ip_external](resources--nfv_service--reference--group-001.md#canonical-1330300133332010-2213201212020232-2201233300010002-2200203110002302-1122333112303322-1110111013131122-1030231200003221-2022230211320202): complete subsection reference.

- [automatic_vip](resources--nfv_service--reference--group-001.md#canonical-0213220001313030-3333031231203323-2320020210123302-1123231021030102-1223213312002023-2100333303212210-3222112130330302-0023202210300010): complete subsection reference.

<a id="canonical-2002310120303123-0122223120010230-1331303022213130-1111000013013311-1231112322001123-2300220000333111-1011331130031102-1323301311233223"></a>

<a id="canonical-3113301103232203-3033303123102321-0030222220012011-0233100022313301-2113132000120033-2120212232011112-0102312322031230-0212321000102332"></a>

#### `f5_big_ip_aws_service.endpoint_service.configured_vip` property

Type: `"string"`. Optional.

Exclusive with \[automatic\_vip\] Enter IP address for the default VIP.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ip",
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
    "ves.io.schema.rules.string.ip": "true",
    "ves.io.schema.rules.string.not_in": "192.0.2.26"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ip": "true",
    "ves.io.schema.rules.string.not_in": "192.0.2.26"
  }
}
```

- [custom_tcp_ports](resources--nfv_service--reference--group-001.md#canonical-2020332122210210-0210202110120033-2323010311113002-3020202302003220-3131110333231313-2030202323331202-2201320212203332-0021203122323213): complete subsection reference.

- [custom_udp_ports](resources--nfv_service--reference--group-001.md#canonical-3121301132011211-1100200303020202-2312221303312103-3211230311330023-3210031301110131-0023030101213312-2102223221231013-1101322112212331): complete subsection reference.

- [default_tcp_ports](resources--nfv_service--reference--group-001.md#canonical-0220032001130202-2121033210102332-0021210100021000-3321012202131002-0333333323320303-3311031213001020-3033110131121012-2331200332230113): complete subsection reference.

- [disable_advertise_on_slo_ip](resources--nfv_service--reference--group-001.md#canonical-1133012012021202-1130332102023111-0002320121202131-3023123223110123-2221322121201022-0003100103303331-0030211232030110-2313312110230331): complete subsection reference.

- [http_port](resources--nfv_service--reference--group-001.md#canonical-3320223321311023-3330032333213303-2231201320223001-2111101011102332-0131203102203200-0221001032103310-0112313113120020-2120000231003131): complete subsection reference.

- [https_port](resources--nfv_service--reference--group-002.md#canonical-1113221322000011-1313031012002033-2123303222033131-1023101101330233-1222020101023123-2130012102302313-0302322023032013-3110122331011130): complete subsection reference.

- [no_tcp_ports](resources--nfv_service--reference--group-002.md#canonical-2323230322303213-2133022201311333-0013103330033033-3122301030103131-2303322233110330-3002113230113033-0230331312202212-1232301203313133): complete subsection reference.

- [no_udp_ports](resources--nfv_service--reference--group-002.md#canonical-0032211122030020-3003323212002133-2120003300011011-1012223010030013-1022302100113132-3330231031213131-3320022302013010-0110001302110001): complete subsection reference.

<a id="canonical-2002023200130011-3221031100321323-1023033331122313-3230213010102032-2121211120330320-2103303103132200-3323232000311320-1022321130223012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `f5_big_ip_aws_service.endpoint_service.advertise_on_slo_ip` properties

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-0223321113210321-0013210302112221-0313201132021001-1102330023301310-1131322132000230-3303202123330303-3122032322302202-2202303201230222)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-1033200313012130-2303320020000220-3032203302003212-3030323031320111-0322032022323123-0203102131300133-1232321110231101-1022222010201320)
- [f5_big_ip_aws_service](resources--nfv_service--reference--group-001.md#canonical-2203212022212003-2212301132213201-2010231030110211-3120210322102100-3130230103331230-0332101002301003-2302320031110011-3121121001302210)
- [f5_big_ip_aws_service.endpoint_service](resources--nfv_service--reference--group-001.md#canonical-1010032200210212-2313121332112310-0110323222301213-2211012302000323-2311222030300112-0122023013313002-3223113221233302-2201323331331100)
- f5_big_ip_aws_service.endpoint_service.advertise_on_slo_ip

<a id="canonical-2021020020222320-3012232220211220-3202333010211121-3031111112212132-2012111213221301-2112320332101011-2231213112111021-0020221033003003"></a>

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
advertise_on_slo_ip = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1330300133332010-2213201212020232-2201233300010002-2200203110002302-1122333112303322-1110111013131122-1030231200003221-2022230211320202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `f5_big_ip_aws_service.endpoint_service.advertise_on_slo_ip_external` properties

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-0223321113210321-0013210302112221-0313201132021001-1102330023301310-1131322132000230-3303202123330303-3122032322302202-2202303201230222)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-1033200313012130-2303320020000220-3032203302003212-3030323031320111-0322032022323123-0203102131300133-1232321110231101-1022222010201320)
- [f5_big_ip_aws_service](resources--nfv_service--reference--group-001.md#canonical-2203212022212003-2212301132213201-2010231030110211-3120210322102100-3130230103331230-0332101002301003-2302320031110011-3121121001302210)
- [f5_big_ip_aws_service.endpoint_service](resources--nfv_service--reference--group-001.md#canonical-1010032200210212-2313121332112310-0110323222301213-2211012302000323-2311222030300112-0122023013313002-3223113221233302-2201323331331100)
- f5_big_ip_aws_service.endpoint_service.advertise_on_slo_ip_external

<a id="canonical-3310321211232133-1223311010310302-1003310022203022-1332002300203302-3330232100131223-2130132201220301-3203202323230203-2123102210133232"></a>

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
advertise_on_slo_ip_external = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0213220001313030-3333031231203323-2320020210123302-1123231021030102-1223213312002023-2100333303212210-3222112130330302-0023202210300010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `f5_big_ip_aws_service.endpoint_service.automatic_vip` properties

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-0223321113210321-0013210302112221-0313201132021001-1102330023301310-1131322132000230-3303202123330303-3122032322302202-2202303201230222)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-1033200313012130-2303320020000220-3032203302003212-3030323031320111-0322032022323123-0203102131300133-1232321110231101-1022222010201320)
- [f5_big_ip_aws_service](resources--nfv_service--reference--group-001.md#canonical-2203212022212003-2212301132213201-2010231030110211-3120210322102100-3130230103331230-0332101002301003-2302320031110011-3121121001302210)
- [f5_big_ip_aws_service.endpoint_service](resources--nfv_service--reference--group-001.md#canonical-1010032200210212-2313121332112310-0110323222301213-2211012302000323-2311222030300112-0122023013313002-3223113221233302-2201323331331100)
- f5_big_ip_aws_service.endpoint_service.automatic_vip

<a id="canonical-0312211323110033-3303332203130031-0233200011011313-0311113300222102-1132010013122020-3231022320200232-0020003021321312-3301231312132111"></a>

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
automatic_vip = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2020332122210210-0210202110120033-2323010311113002-3020202302003220-3131110333231313-2030202323331202-2201320212203332-0021203122323213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `f5_big_ip_aws_service.endpoint_service.custom_tcp_ports` properties

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-0223321113210321-0013210302112221-0313201132021001-1102330023301310-1131322132000230-3303202123330303-3122032322302202-2202303201230222)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-1033200313012130-2303320020000220-3032203302003212-3030323031320111-0322032022323123-0203102131300133-1232321110231101-1022222010201320)
- [f5_big_ip_aws_service](resources--nfv_service--reference--group-001.md#canonical-2203212022212003-2212301132213201-2010231030110211-3120210322102100-3130230103331230-0332101002301003-2302320031110011-3121121001302210)
- [f5_big_ip_aws_service.endpoint_service](resources--nfv_service--reference--group-001.md#canonical-1010032200210212-2313121332112310-0110323222301213-2211012302000323-2311222030300112-0122023013313002-3223113221233302-2201323331331100)
- f5_big_ip_aws_service.endpoint_service.custom_tcp_ports

<a id="canonical-1220013311203222-3322132211133220-3222233221333103-3122132300301213-0102113223033230-2303330131202101-2222312332222110-3312331211221202"></a>

Type: `"object"`. single nested block, Optional.

Port Range List. List of port ranges.

Receipt-pinned upstream constraints:

```json
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
custom_tcp_ports {
  # Configure direct properties listed below.
}
```

<a id="canonical-0103322111031113-2111113120122123-0133301101132220-1213001032231003-1001212310233021-1100200112233320-0130220132113310-2231101321101022"></a>

### Direct properties for `f5_big_ip_aws_service.endpoint_service.custom_tcp_ports`

<a id="canonical-0000120102002302-2201133012233123-1230220122301302-1230003031013320-0111023310200223-1101200332310333-2031202100131312-3000113233001212"></a>

#### `f5_big_ip_aws_service.endpoint_service.custom_tcp_ports.ports` property

Type: `["list", "string"]`. Optional.

List of port ranges. Each range is a single port or a pair of start and end ports e.g. 8080-8192.

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
    "ves.io.schema.rules.repeated.items.string.port_range": "true",
    "ves.io.schema.rules.repeated.max_items": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.port_range": "true",
    "ves.io.schema.rules.repeated.max_items": "128"
  }
}
```

<a id="canonical-3121301132011211-1100200303020202-2312221303312103-3211230311330023-3210031301110131-0023030101213312-2102223221231013-1101322112212331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `f5_big_ip_aws_service.endpoint_service.custom_udp_ports` properties

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-0223321113210321-0013210302112221-0313201132021001-1102330023301310-1131322132000230-3303202123330303-3122032322302202-2202303201230222)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-1033200313012130-2303320020000220-3032203302003212-3030323031320111-0322032022323123-0203102131300133-1232321110231101-1022222010201320)
- [f5_big_ip_aws_service](resources--nfv_service--reference--group-001.md#canonical-2203212022212003-2212301132213201-2010231030110211-3120210322102100-3130230103331230-0332101002301003-2302320031110011-3121121001302210)
- [f5_big_ip_aws_service.endpoint_service](resources--nfv_service--reference--group-001.md#canonical-1010032200210212-2313121332112310-0110323222301213-2211012302000323-2311222030300112-0122023013313002-3223113221233302-2201323331331100)
- f5_big_ip_aws_service.endpoint_service.custom_udp_ports

<a id="canonical-2300333333323321-1030021333102011-1003031021303020-1200323131003122-3011323220310103-2031223201111232-1330110013332133-3120331222011303"></a>

Type: `"object"`. single nested block, Optional.

Port Range List. List of port ranges.

Receipt-pinned upstream constraints:

```json
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
custom_udp_ports {
  # Configure direct properties listed below.
}
```

<a id="canonical-3302311320100022-1310301331033311-2202331232201200-2333210201003122-1201221133302112-0320220322312210-3102320120011112-2301321203330311"></a>

### Direct properties for `f5_big_ip_aws_service.endpoint_service.custom_udp_ports`

<a id="canonical-1012021101223321-3223310022000022-3032110012001230-2333231232021122-3331213302133102-2302311202031300-0213103303313231-1330023100222032"></a>

#### `f5_big_ip_aws_service.endpoint_service.custom_udp_ports.ports` property

Type: `["list", "string"]`. Optional.

List of port ranges. Each range is a single port or a pair of start and end ports e.g. 8080-8192.

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
    "ves.io.schema.rules.repeated.items.string.port_range": "true",
    "ves.io.schema.rules.repeated.max_items": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.port_range": "true",
    "ves.io.schema.rules.repeated.max_items": "128"
  }
}
```

<a id="canonical-0220032001130202-2121033210102332-0021210100021000-3321012202131002-0333333323320303-3311031213001020-3033110131121012-2331200332230113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `f5_big_ip_aws_service.endpoint_service.default_tcp_ports` properties

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-0223321113210321-0013210302112221-0313201132021001-1102330023301310-1131322132000230-3303202123330303-3122032322302202-2202303201230222)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-1033200313012130-2303320020000220-3032203302003212-3030323031320111-0322032022323123-0203102131300133-1232321110231101-1022222010201320)
- [f5_big_ip_aws_service](resources--nfv_service--reference--group-001.md#canonical-2203212022212003-2212301132213201-2010231030110211-3120210322102100-3130230103331230-0332101002301003-2302320031110011-3121121001302210)
- [f5_big_ip_aws_service.endpoint_service](resources--nfv_service--reference--group-001.md#canonical-1010032200210212-2313121332112310-0110323222301213-2211012302000323-2311222030300112-0122023013313002-3223113221233302-2201323331331100)
- f5_big_ip_aws_service.endpoint_service.default_tcp_ports

<a id="canonical-2030023210312113-2001323312121333-0012220033200223-2333130023110002-3232001113122221-1121001322301002-2030311213033313-0221213301302220"></a>

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
default_tcp_ports = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1133012012021202-1130332102023111-0002320121202131-3023123223110123-2221322121201022-0003100103303331-0030211232030110-2313312110230331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `f5_big_ip_aws_service.endpoint_service.disable_advertise_on_slo_ip` properties

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-0223321113210321-0013210302112221-0313201132021001-1102330023301310-1131322132000230-3303202123330303-3122032322302202-2202303201230222)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-1033200313012130-2303320020000220-3032203302003212-3030323031320111-0322032022323123-0203102131300133-1232321110231101-1022222010201320)
- [f5_big_ip_aws_service](resources--nfv_service--reference--group-001.md#canonical-2203212022212003-2212301132213201-2010231030110211-3120210322102100-3130230103331230-0332101002301003-2302320031110011-3121121001302210)
- [f5_big_ip_aws_service.endpoint_service](resources--nfv_service--reference--group-001.md#canonical-1010032200210212-2313121332112310-0110323222301213-2211012302000323-2311222030300112-0122023013313002-3223113221233302-2201323331331100)
- f5_big_ip_aws_service.endpoint_service.disable_advertise_on_slo_ip

<a id="canonical-3102223111322131-1130121300023132-2000322232100321-1131123113213121-0302333021013002-0012023130133023-3321221211133301-3101120031020230"></a>

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
disable_advertise_on_slo_ip = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3320223321311023-3330032333213303-2231201320223001-2111101011102332-0131203102203200-0221001032103310-0112313113120020-2120000231003131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `f5_big_ip_aws_service.endpoint_service.http_port` properties

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-0223321113210321-0013210302112221-0313201132021001-1102330023301310-1131322132000230-3303202123330303-3122032322302202-2202303201230222)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-1033200313012130-2303320020000220-3032203302003212-3030323031320111-0322032022323123-0203102131300133-1232321110231101-1022222010201320)
- [f5_big_ip_aws_service](resources--nfv_service--reference--group-001.md#canonical-2203212022212003-2212301132213201-2010231030110211-3120210322102100-3130230103331230-0332101002301003-2302320031110011-3121121001302210)
- [f5_big_ip_aws_service.endpoint_service](resources--nfv_service--reference--group-001.md#canonical-1010032200210212-2313121332112310-0110323222301213-2211012302000323-2311222030300112-0122023013313002-3223113221233302-2201323331331100)
- f5_big_ip_aws_service.endpoint_service.http_port

<a id="canonical-0122010213002333-2131130211001233-2002311002230223-3322212201210133-2002012110231301-1023101000231223-2331131302102202-0230022101212003"></a>

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
http_port = {}
```

This is an empty object or choice marker. It has no direct properties.
