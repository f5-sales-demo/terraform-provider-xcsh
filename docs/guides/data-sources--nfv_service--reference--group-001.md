---
page_title: "xcsh_nfv_service reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_nfv_service reference."
---

# xcsh_nfv_service reference

<a id="canonical-2313010333323131-3010030231311121-2301323120300012-0122321020201331-3020223211220022-0100121100023011-2323321311300323-0122100132003020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1223310212302332-1201003320332030-0200020022311113-2032020310021000-1000133112023033-1000003010120303-2021233203122112-0301333020102313"></a>

## Property reference — Property reference / 011321300323 / 2

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-2212320103310132-2221300222221032-1233031123200120-3022110322311231-3301003023302233-3133232120201022-0322021211220213-1332311000003300)
- Property reference

<a id="canonical-2230112311100111-1313113310123330-3213320000203112-3313312011323100-0301013131020023-3030103220330223-0233030313312103-2301020323002110"></a>

## Direct properties — Property reference / 011321300323 / 3

<a id="canonical-1201233331031300-3200303302121020-1012331201121322-2221320031323312-1022012003010133-3120010200212033-2001003221020132-1311013332010301"></a>

<a id="canonical-0122131123230113-0233231301223113-2101331333313220-1310222000010211-1220022023103013-3300230301032230-0033213032302312-1311322033111330"></a>

## annotations property — Property reference / 011321300323 / 4

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

<a id="canonical-3302132100111220-2203003231221300-3323021130011031-3032000300222331-0311233201021120-3232033020022101-1100133223002013-3213302210001103"></a>

<a id="canonical-0212033112132120-2322021120132012-2023213103102111-3000301220322133-2111323232013232-0021213023033323-0003201102330010-0103123110322201"></a>

## description property — Property reference / 011321300323 / 5

Type: `"string"`. Computed.

Description of the NfvService.

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

- [disable_https_management](data-sources--nfv_service--reference--group-001.md#canonical-3233032123222311-3001001210121231-0312011132321313-1232230220130301-3231301131202311-0233202113011321-1032221330102003-2312002332132102): complete subsection reference.

- [disable_ssh_access](data-sources--nfv_service--reference--group-001.md#canonical-1221033000332001-2223320310121121-0001202121031020-2311110013220222-3223210220123033-3303322313032323-2103300120311112-3302300312112212): complete subsection reference.

- [enabled_ssh_access](data-sources--nfv_service--reference--group-001.md#canonical-2323120212210131-3332303320001030-2120313330111100-0101001020220303-0323332221022001-1220132333131230-1200303110333213-2001222021201023): complete subsection reference.

- [f5_big_ip_aws_service](data-sources--nfv_service--reference--group-001.md#canonical-3121203202232102-3301112102333133-2103303100231002-1321223111220201-3103311113020000-0303332002322201-3021331113310133-2201221310222221): complete subsection reference.

- [https_management](data-sources--nfv_service--reference--group-002.md#canonical-0322223112333211-3323030031122111-0131002222131221-2002123332012320-0013112013120020-2030102011102030-0020230010011012-0231301033231112): complete subsection reference.

<a id="canonical-3010132210033022-2121332002321211-0011101120213002-1122103003320001-0100212223322322-0203302332101331-2123211012210311-3332022223100230"></a>

<a id="canonical-3322213100112033-3301231001200231-2301220131101313-3122033022000002-2213232110032131-0211312213132330-2031010002122103-1010023202333032"></a>

## ID property — Property reference / 011321300323 / 6

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-0011002330323300-2232322212023333-2203202300321130-3111000101202103-2323320322031122-0033122000210232-2203102320001121-1033312130312122"></a>

<a id="canonical-1010131302133100-3231210231003320-0321300221002201-2321331230033103-1202222033322230-1132230102112010-2222331030232021-1330332103000100"></a>

## labels property — Property reference / 011321300323 / 7

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

<a id="canonical-0323220030233032-3210121100100012-3321333223121023-1213311013233123-0332000032210012-2312133323332212-2012321010011122-0301313203120111"></a>

<a id="canonical-1112020102122211-1003333000312022-2202312333102321-3003222032031112-0113301030332211-2023213232212100-1312233021123210-3011121022112112"></a>

## name property — Property reference / 011321300323 / 8

Type: `"string"`. Required.

Name of the NfvService.

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

<a id="canonical-2310122223303213-0030303221000210-2220113231313200-0230103312203301-1022221200303123-0331221220210132-1032110132003001-1330231330302100"></a>

<a id="canonical-3023321103012230-3120132331100223-0303213321232030-1011202332012222-1311111200323023-1231313311132313-0311022323200330-1203030320300312"></a>

## namespace property — Property reference / 011321300323 / 9

Type: `"string"`. Required.

Namespace where the NfvService exists.

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

- [palo_alto_fw_service](data-sources--nfv_service--reference--group-003.md#canonical-0130222001222131-3123210310203012-0200310203211313-0131331300023100-2310131123330002-1011033223002202-2221031230120202-1331030023120030): complete subsection reference.

<a id="canonical-1223011220320210-1311311033202121-1120320020101220-3023322212123123-3022330003132210-1133023001030303-0313311331202320-1230312130022210"></a>

## All schema paths — Property reference / 011321300323 / 10

Each exact path has one authoritative reference destination. Collection element indices are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](data-sources--nfv_service--reference--group-001.md#canonical-1201233331031300-3200303302121020-1012331201121322-2221320031323312-1022012003010133-3120010200212033-2001003221020132-1311013332010301) |
| `description` | [description](data-sources--nfv_service--reference--group-001.md#canonical-3302132100111220-2203003231221300-3323021130011031-3032000300222331-0311233201021120-3232033020022101-1100133223002013-3213302210001103) |
| `disable_https_management` | [disable_https_management](data-sources--nfv_service--reference--group-001.md#canonical-1123230132031022-2230003222203310-3220003002031030-0130133021301033-1110001030131201-2010020310131122-3212333113023313-1013010302300230) |
| `disable_ssh_access` | [disable_ssh_access](data-sources--nfv_service--reference--group-001.md#canonical-3322300221020310-1022320213222010-1003030332021023-1003202032123101-2331012113322001-3022212232102110-0102011121223220-2033000110230323) |
| `enabled_ssh_access` | [enabled_ssh_access](data-sources--nfv_service--reference--group-001.md#canonical-3231302020013323-0132223211210223-2112032102231030-2123122230000331-1031332220111111-2232322012102201-0222213113233310-1330222012001120) |
| `enabled_ssh_access.advertise_on_sli` | [enabled_ssh_access.advertise_on_sli](data-sources--nfv_service--reference--group-001.md#canonical-3021212023001322-1213203013021332-3102001100030201-0103222211030022-3202220221332023-3131300030321000-3122101301112223-2001033303230210) |
| `enabled_ssh_access.advertise_on_slo` | [enabled_ssh_access.advertise_on_slo](data-sources--nfv_service--reference--group-001.md#canonical-0302030102011032-3020122123312323-0333020011112021-0223012110013000-2212100200102231-2003200120101003-2012021310120131-3101113030133130) |
| `enabled_ssh_access.advertise_on_slo_sli` | [enabled_ssh_access.advertise_on_slo_sli](data-sources--nfv_service--reference--group-001.md#canonical-3123230300313323-3022122023222303-3221203330310313-0120003300231313-0101033323132021-2202333103331202-3113112222013323-1211002331322133) |
| `enabled_ssh_access.domain_suffix` | [enabled_ssh_access.domain_suffix](data-sources--nfv_service--reference--group-001.md#canonical-2213222112030101-3333013110223330-2222121133111131-0111221330032032-2332201233311323-0111133303131220-3133111320332322-2331112320010201) |
| `enabled_ssh_access.node_ssh_ports` | [enabled_ssh_access.node_ssh_ports](data-sources--nfv_service--reference--group-001.md#canonical-3302313201101220-0011123021313023-3220210111020130-1000331220221102-3312202323230323-1211223301110032-2020012303311032-1030332102230220) |
| `enabled_ssh_access.node_ssh_ports.node_name` | [enabled_ssh_access.node_ssh_ports.node_name](data-sources--nfv_service--reference--group-001.md#canonical-1211232230300010-2003330312233223-0203213131011110-0332132213221020-0323010210000323-2022333130123311-3103222332022023-1111101321331303) |
| `enabled_ssh_access.node_ssh_ports.ssh_port` | [enabled_ssh_access.node_ssh_ports.ssh_port](data-sources--nfv_service--reference--group-001.md#canonical-2212222320011231-2301211123023133-1111210130033300-1230021321211212-3321032202212031-1103312333333310-2332030332333233-2003010303110202) |
| `f5_big_ip_aws_service` | [f5_big_ip_aws_service](data-sources--nfv_service--reference--group-001.md#canonical-1000212232223202-3102233333221300-1111113133333303-1002121012300313-0201010302013330-2021303201311211-0323112122022121-0330010210332033) |
| `f5_big_ip_aws_service.admin_password` | [f5_big_ip_aws_service.admin_password](data-sources--nfv_service--reference--group-001.md#canonical-0203130131012333-1200022021221312-3313033221032333-3331123311232131-1331130312030123-3023012121113133-0302213012201312-3323111130003003) |
| `f5_big_ip_aws_service.admin_password.blindfold_secret_info` | [f5_big_ip_aws_service.admin_password.blindfold_secret_info](data-sources--nfv_service--reference--group-001.md#canonical-1220030023200320-2012032331323102-2111333022321113-2110023100032203-2003221223102111-2211121301022112-2232113013003332-0000300012203320) |
| `f5_big_ip_aws_service.admin_password.blindfold_secret_info.decryption_provider` | [f5_big_ip_aws_service.admin_password.blindfold_secret_info.decryption_provider](data-sources--nfv_service--reference--group-001.md#canonical-3100233013221321-2222001321002200-0223111203310311-1101031222101103-1010030112021230-1313103013300031-0122203030221011-2230023211000230) |
| `f5_big_ip_aws_service.admin_password.blindfold_secret_info.location` | [f5_big_ip_aws_service.admin_password.blindfold_secret_info.location](data-sources--nfv_service--reference--group-001.md#canonical-0130113201112121-2103101102302030-2230121310031123-1030321002010121-2303100113213331-1311033211000232-3103123113213023-1120231220220022) |
| `f5_big_ip_aws_service.admin_password.blindfold_secret_info.store_provider` | [f5_big_ip_aws_service.admin_password.blindfold_secret_info.store_provider](data-sources--nfv_service--reference--group-001.md#canonical-3003310333102211-0301120311203122-0321033133002021-3313122331223223-2303013002030023-0033122300322210-0322120002332301-2203022231300111) |
| `f5_big_ip_aws_service.admin_password.clear_secret_info` | [f5_big_ip_aws_service.admin_password.clear_secret_info](data-sources--nfv_service--reference--group-001.md#canonical-1202031110132212-3212011233300213-2131223321233010-1000032230101221-3211201332333313-0011213130133333-0020333111112003-3301020121020010) |
| `f5_big_ip_aws_service.admin_password.clear_secret_info.provider_ref` | [f5_big_ip_aws_service.admin_password.clear_secret_info.provider_ref](data-sources--nfv_service--reference--group-001.md#canonical-1213030032230021-3011131310320110-1221032002021113-1032222320311101-3113301111020320-0110010313000332-0032111013112222-3031030303303331) |
| `f5_big_ip_aws_service.admin_password.clear_secret_info.url` | [f5_big_ip_aws_service.admin_password.clear_secret_info.url](data-sources--nfv_service--reference--group-001.md#canonical-1213013020300133-2130222000312102-0231000123103132-2112211300301310-1322033003033002-3320031102110003-2300130232030313-1023210121223210) |
| `f5_big_ip_aws_service.admin_username` | [f5_big_ip_aws_service.admin_username](data-sources--nfv_service--reference--group-001.md#canonical-0211302020113331-0302232110022122-1020122001132313-2022332110130120-2312013200102330-3120323133102300-3100221000210303-3100120312311121) |
| `f5_big_ip_aws_service.aws_tgw_site_params` | [f5_big_ip_aws_service.aws_tgw_site_params](data-sources--nfv_service--reference--group-001.md#canonical-1023032330102020-1233233323001332-0332223010131311-0301003330030102-3120221211003023-1001212032320202-0332230021030201-1320220120030321) |
| `f5_big_ip_aws_service.aws_tgw_site_params.aws_tgw_site` | [f5_big_ip_aws_service.aws_tgw_site_params.aws_tgw_site](data-sources--nfv_service--reference--group-001.md#canonical-1010020320213320-3103132210221212-0302122111111330-0333131100203133-3102220211332332-3130331302020130-2211320202313113-2330101212133200) |
| `f5_big_ip_aws_service.aws_tgw_site_params.aws_tgw_site.name` | [f5_big_ip_aws_service.aws_tgw_site_params.aws_tgw_site.name](data-sources--nfv_service--reference--group-001.md#canonical-3203121021223331-1210221010331110-3323011032103102-1011230201113303-3031131301111322-2100031331321130-3202031130333112-0230101033203330) |
| `f5_big_ip_aws_service.aws_tgw_site_params.aws_tgw_site.namespace` | [f5_big_ip_aws_service.aws_tgw_site_params.aws_tgw_site.namespace](data-sources--nfv_service--reference--group-001.md#canonical-2121231111121232-3123320130230313-1200320302011211-0231321032031202-0010120101001321-1203311012130022-0303000222321010-1322301111332331) |
| `f5_big_ip_aws_service.aws_tgw_site_params.aws_tgw_site.tenant` | [f5_big_ip_aws_service.aws_tgw_site_params.aws_tgw_site.tenant](data-sources--nfv_service--reference--group-001.md#canonical-2120003030330223-1032220221021203-2023012113032300-1122031212233210-1021022111030021-2210233301111023-3211221003121321-1103123013013100) |
| `f5_big_ip_aws_service.endpoint_service` | [f5_big_ip_aws_service.endpoint_service](data-sources--nfv_service--reference--group-001.md#canonical-2102032020301333-2031221001211012-0233322230323103-3321200112100000-0120101110020230-2100102323111221-1122312221333020-2322022001320220) |
| `f5_big_ip_aws_service.endpoint_service.advertise_on_slo_ip` | [f5_big_ip_aws_service.endpoint_service.advertise_on_slo_ip](data-sources--nfv_service--reference--group-001.md#canonical-3222013222330332-0211222310100332-0002230130002330-1010130113302110-2323103102020223-1023330032200123-0103122001232021-2223330130333001) |
| `f5_big_ip_aws_service.endpoint_service.advertise_on_slo_ip_external` | [f5_big_ip_aws_service.endpoint_service.advertise_on_slo_ip_external](data-sources--nfv_service--reference--group-001.md#canonical-1210312003330221-1031021202310023-0001001333111211-0201210010212230-3211123013111203-1330300331313303-0112001011230213-2110030312101120) |
| `f5_big_ip_aws_service.endpoint_service.automatic_vip` | [f5_big_ip_aws_service.endpoint_service.automatic_vip](data-sources--nfv_service--reference--group-001.md#canonical-2211133003033002-3323220103223323-0101023323110303-0202123111110000-3103020022103321-3020323221123120-3030020132101303-3230323202002203) |
| `f5_big_ip_aws_service.endpoint_service.configured_vip` | [f5_big_ip_aws_service.endpoint_service.configured_vip](data-sources--nfv_service--reference--group-001.md#canonical-1321211322210330-3103122021213020-1000013022312021-2230023210122030-2111010320203111-1100021211320001-0122032302202122-2311003330131200) |
| `f5_big_ip_aws_service.endpoint_service.custom_tcp_ports` | [f5_big_ip_aws_service.endpoint_service.custom_tcp_ports](data-sources--nfv_service--reference--group-001.md#canonical-3303211321230323-2103233110022011-2333311220020202-2201311302332100-1120200002101022-1032213330332001-3231010121310323-2111110102031123) |
| `f5_big_ip_aws_service.endpoint_service.custom_tcp_ports.ports` | [f5_big_ip_aws_service.endpoint_service.custom_tcp_ports.ports](data-sources--nfv_service--reference--group-001.md#canonical-3003103212223022-2132210022233231-0323210112300222-2221300310303320-3333103002013110-3211313312110011-1323333330300311-3031001013202230) |
| `f5_big_ip_aws_service.endpoint_service.custom_udp_ports` | [f5_big_ip_aws_service.endpoint_service.custom_udp_ports](data-sources--nfv_service--reference--group-001.md#canonical-0123132012113300-2332213331331000-2031221221000011-0023323201312303-3003310332011202-3112132110223330-0101230332201201-0323012212332023) |
| `f5_big_ip_aws_service.endpoint_service.custom_udp_ports.ports` | [f5_big_ip_aws_service.endpoint_service.custom_udp_ports.ports](data-sources--nfv_service--reference--group-001.md#canonical-3102120312302230-0101321303222301-1302130013332031-0233210212220111-0101203203323202-3100131322312001-2212002301311113-1333331111210032) |
| `f5_big_ip_aws_service.endpoint_service.default_tcp_ports` | [f5_big_ip_aws_service.endpoint_service.default_tcp_ports](data-sources--nfv_service--reference--group-001.md#canonical-2003222221331333-3011032112013113-0333210312032000-1313331311213020-3113333111130010-2010112032201212-1223333322211300-0030101221300311) |
| `f5_big_ip_aws_service.endpoint_service.disable_advertise_on_slo_ip` | [f5_big_ip_aws_service.endpoint_service.disable_advertise_on_slo_ip](data-sources--nfv_service--reference--group-001.md#canonical-0111122032013011-2311120202300221-1213330000033031-2012311233300122-0322310020132321-3223211223220303-0021213223032003-1211122022033313) |
| `f5_big_ip_aws_service.endpoint_service.http_port` | [f5_big_ip_aws_service.endpoint_service.http_port](data-sources--nfv_service--reference--group-001.md#canonical-0210321120300122-2222022122301331-0112303233133203-1333113201302113-2031002112221123-1203003000120230-0302302031112233-1103010110120232) |
| `f5_big_ip_aws_service.endpoint_service.https_port` | [f5_big_ip_aws_service.endpoint_service.https_port](data-sources--nfv_service--reference--group-001.md#canonical-2031102231123010-0030111011112002-3103023001012110-1131303301121101-3232103213202001-3321132132233202-2300200033011323-1223212302232233) |
| `f5_big_ip_aws_service.endpoint_service.no_tcp_ports` | [f5_big_ip_aws_service.endpoint_service.no_tcp_ports](data-sources--nfv_service--reference--group-001.md#canonical-3301111110211233-2032030200120103-0001202120131113-0333212211023003-1111003200131230-2021121030111331-1311302132210012-1002333331132130) |
| `f5_big_ip_aws_service.endpoint_service.no_udp_ports` | [f5_big_ip_aws_service.endpoint_service.no_udp_ports](data-sources--nfv_service--reference--group-001.md#canonical-3132223311022000-1302212113332033-2312101113201202-3012012302022033-3303122020331102-1330101113122120-0033123111123323-0330302021032232) |
| `f5_big_ip_aws_service.market_place_image` | [f5_big_ip_aws_service.market_place_image](data-sources--nfv_service--reference--group-001.md#canonical-1230320202121112-2210332203231211-1311220012032233-3011113003020230-0131303213211021-2302212122000003-1322131212121302-3003022001123313) |
| `f5_big_ip_aws_service.market_place_image.awafpay_g200_mbps` | [f5_big_ip_aws_service.market_place_image.awafpay_g200_mbps](data-sources--nfv_service--reference--group-002.md#canonical-2023103103032113-2233100103000332-2100113030023103-2310202201033013-1133032201021000-3203232200332312-2002311001001000-1103212031110000) |
| `f5_big_ip_aws_service.market_place_image.awafpay_g3_gbps` | [f5_big_ip_aws_service.market_place_image.awafpay_g3_gbps](data-sources--nfv_service--reference--group-002.md#canonical-3020211211011321-1202101033121111-0132202313332202-3300332111101302-1222302323320211-3233131313323111-0101323220232120-0103311300030323) |
| `f5_big_ip_aws_service.nodes` | [f5_big_ip_aws_service.nodes](data-sources--nfv_service--reference--group-002.md#canonical-3033221013330113-0111102030012021-1203332011032233-0013223120022233-1110100331123321-3011032322222123-0002130033213331-3221302233322132) |
| `f5_big_ip_aws_service.nodes.automatic_prefix` | [f5_big_ip_aws_service.nodes.automatic_prefix](data-sources--nfv_service--reference--group-002.md#canonical-3301330100331131-1203312023002220-1330022330332232-2313231133331212-0233220100033233-0130233021012132-0323323132123300-2300300313311232) |
| `f5_big_ip_aws_service.nodes.aws_az_name` | [f5_big_ip_aws_service.nodes.aws_az_name](data-sources--nfv_service--reference--group-002.md#canonical-2322223222320122-1001111010231232-0212113231102101-1232210032323110-3320011112300330-2001301112303310-1102003212321233-3301122232121231) |
| `f5_big_ip_aws_service.nodes.mgmt_subnet` | [f5_big_ip_aws_service.nodes.mgmt_subnet](data-sources--nfv_service--reference--group-002.md#canonical-3232011020301101-0232133132033112-0010001211011103-2220323232220333-0112032031132230-2322132310301120-0032330030120132-1133330022023310) |
| `f5_big_ip_aws_service.nodes.mgmt_subnet.existing_subnet_id` | [f5_big_ip_aws_service.nodes.mgmt_subnet.existing_subnet_id](data-sources--nfv_service--reference--group-002.md#canonical-2012303310230221-1023320122032033-0120231002102323-1330301120012312-2223031331231131-2120233011220131-0021110032302201-2103122332200020) |
| `f5_big_ip_aws_service.nodes.mgmt_subnet.subnet_param` | [f5_big_ip_aws_service.nodes.mgmt_subnet.subnet_param](data-sources--nfv_service--reference--group-002.md#canonical-0132000211013003-2132221200311331-1331032322313213-3131310110131322-2010022200213312-1312303111222003-1010332330120023-3232232101300303) |
| `f5_big_ip_aws_service.nodes.mgmt_subnet.subnet_param.ipv4` | [f5_big_ip_aws_service.nodes.mgmt_subnet.subnet_param.ipv4](data-sources--nfv_service--reference--group-002.md#canonical-3012230301310303-0232203012123001-0312203113132122-0001000313112331-2030011221000300-1112301000113202-1330311031012311-0222033321123010) |
| `f5_big_ip_aws_service.nodes.node_name` | [f5_big_ip_aws_service.nodes.node_name](data-sources--nfv_service--reference--group-002.md#canonical-1211313212222122-1120301000231300-2032013322030030-1013000301302010-3022130310112012-0211122222132202-1123022010302102-3100310120232221) |
| `f5_big_ip_aws_service.nodes.reserved_mgmt_subnet` | [f5_big_ip_aws_service.nodes.reserved_mgmt_subnet](data-sources--nfv_service--reference--group-002.md#canonical-3120333313030110-0102200123002322-0003111222130230-1103133132000102-3001223223101001-2203010231233000-2201312023330033-2302200003202233) |
| `f5_big_ip_aws_service.nodes.tunnel_prefix` | [f5_big_ip_aws_service.nodes.tunnel_prefix](data-sources--nfv_service--reference--group-002.md#canonical-2220110021213212-1133300002201302-2112101232122010-2303100030032012-0112130023333301-1220200302330231-2030122111113232-2011303110333301) |
| `f5_big_ip_aws_service.ssh_key` | [f5_big_ip_aws_service.ssh_key](data-sources--nfv_service--reference--group-001.md#canonical-3223311100111002-2023301003333220-0220310100330022-2131030032203122-1022123103030002-3123111220122031-3120303211113233-2232201220312321) |
| `f5_big_ip_aws_service.tags` | [f5_big_ip_aws_service.tags](data-sources--nfv_service--reference--group-001.md#canonical-1132203121331230-1220113023121303-3300332302230221-2013221001023101-3230133132130132-1321301112033301-1003113232132101-0000202033121102) |
| `https_management` | [https_management](data-sources--nfv_service--reference--group-002.md#canonical-3331102201020133-2333120011110303-1201032002311221-0030302301022231-0303203311202320-0211011020011033-2330232112032332-3132123012211010) |
| `https_management.advertise_on_internet` | [https_management.advertise_on_internet](data-sources--nfv_service--reference--group-002.md#canonical-0322012231122202-3131030023120301-3032203011320203-0210322001323011-3212012102003303-3021001002231001-2232232330113012-1203313030313113) |
| `https_management.advertise_on_internet.public_ip` | [https_management.advertise_on_internet.public_ip](data-sources--nfv_service--reference--group-002.md#canonical-0213331133222302-1331232222130003-3120033101313110-3002212000013223-1102130033200212-1100203121213302-1031022221132203-1101203310123333) |
| `https_management.advertise_on_internet.public_ip.name` | [https_management.advertise_on_internet.public_ip.name](data-sources--nfv_service--reference--group-002.md#canonical-3122000221211031-0032032220203201-0201231130103023-3031121200302023-2322033211223002-3303203220311201-1012120211300213-3002123100311111) |
| `https_management.advertise_on_internet.public_ip.namespace` | [https_management.advertise_on_internet.public_ip.namespace](data-sources--nfv_service--reference--group-002.md#canonical-0201010102311021-1021220203200101-0310233212132133-0110002330003322-1032223312031212-2010000232231001-2023030110310103-3200333331031232) |
| `https_management.advertise_on_internet.public_ip.tenant` | [https_management.advertise_on_internet.public_ip.tenant](data-sources--nfv_service--reference--group-002.md#canonical-1011123322213122-2003021131022310-1133323323010203-2133320322133132-3221020212301220-2220200331133321-1232003100311100-1032120023221202) |
| `https_management.advertise_on_internet_default_vip` | [https_management.advertise_on_internet_default_vip](data-sources--nfv_service--reference--group-002.md#canonical-1202031020130101-0130232303012032-3223123013110133-3313200232301311-2321312331302332-2032022300000203-0121231002231223-2220233313332130) |
| `https_management.advertise_on_sli_vip` | [https_management.advertise_on_sli_vip](data-sources--nfv_service--reference--group-002.md#canonical-1201133103322011-2203213002022303-0233023300323012-0323323030331323-3102012332003210-1020121110232132-0312301203333020-0221022121313231) |
| `https_management.advertise_on_sli_vip.no_mtls` | [https_management.advertise_on_sli_vip.no_mtls](data-sources--nfv_service--reference--group-002.md#canonical-3032111232130310-0101311100010111-0203033311133313-3323030110112221-0112012301303303-1102032322232302-2222030221313220-3311011203100312) |
| `https_management.advertise_on_sli_vip.tls_certificates` | [https_management.advertise_on_sli_vip.tls_certificates](data-sources--nfv_service--reference--group-002.md#canonical-2002011332302033-2103010220231001-2131132223300011-1100010011003023-0012012023230311-3000323303302221-1000303100301321-0133233221301133) |
| `https_management.advertise_on_sli_vip.tls_certificates.certificate_url` | [https_management.advertise_on_sli_vip.tls_certificates.certificate_url](data-sources--nfv_service--reference--group-002.md#canonical-2312312122020313-2200133300310220-0231213303203133-2110303222222003-1032300220221333-0333122130001031-3002102121223310-0131023010100202) |
| `https_management.advertise_on_sli_vip.tls_certificates.custom_hash_algorithms` | [https_management.advertise_on_sli_vip.tls_certificates.custom_hash_algorithms](data-sources--nfv_service--reference--group-002.md#canonical-2133333032023132-2001313021223012-1301232203301223-3030303331231333-3333332000223112-3311232121231200-3112333013211313-1113133210330330) |
| `https_management.advertise_on_sli_vip.tls_certificates.custom_hash_algorithms.hash_algorithms` | [https_management.advertise_on_sli_vip.tls_certificates.custom_hash_algorithms.hash_algorithms](data-sources--nfv_service--reference--group-002.md#canonical-2110220311230131-1023202021222203-0013333322022220-1032133002132112-2112322023200131-2031112313333233-0201231201231113-0332221133223231) |
| `https_management.advertise_on_sli_vip.tls_certificates.description_spec` | [https_management.advertise_on_sli_vip.tls_certificates.description_spec](data-sources--nfv_service--reference--group-002.md#canonical-1103133323103302-1120223211130233-2102120032312201-0200212300233100-2330302333101011-1210312133022013-1033031222232132-2220003313022311) |
| `https_management.advertise_on_sli_vip.tls_certificates.disable_ocsp_stapling` | [https_management.advertise_on_sli_vip.tls_certificates.disable_ocsp_stapling](data-sources--nfv_service--reference--group-002.md#canonical-0323202120103231-2101300122033322-2310221022102003-1013302220223313-1010213003023320-3123313001223132-2133202033032120-0113201010112122) |
| `https_management.advertise_on_sli_vip.tls_certificates.private_key` | [https_management.advertise_on_sli_vip.tls_certificates.private_key](data-sources--nfv_service--reference--group-002.md#canonical-2003023012103221-2303020132131020-0210210010030010-3033031330033322-1102233211203113-3020202112213031-3123112003323313-3101020032233313) |
| `https_management.advertise_on_sli_vip.tls_certificates.private_key.blindfold_secret_info` | [https_management.advertise_on_sli_vip.tls_certificates.private_key.blindfold_secret_info](data-sources--nfv_service--reference--group-002.md#canonical-3123121232110011-3313311313212012-1300330330333101-2101323003032312-1321001223210113-2112032011321230-2012101123033331-1222211302322332) |
| `https_management.advertise_on_sli_vip.tls_certificates.private_key.blindfold_secret_info.decryption_provider` | [https_management.advertise_on_sli_vip.tls_certificates.private_key.blindfold_secret_info.decryption_provider](data-sources--nfv_service--reference--group-002.md#canonical-3203302012303220-3033030311100321-2230320330213323-0200031213332212-2211130033210112-1020003132330322-1100233323230003-0210013303221011) |
| `https_management.advertise_on_sli_vip.tls_certificates.private_key.blindfold_secret_info.location` | [https_management.advertise_on_sli_vip.tls_certificates.private_key.blindfold_secret_info.location](data-sources--nfv_service--reference--group-002.md#canonical-2102230020303312-1220011101001200-3301033332012313-3001120200230313-3233122011111103-1111122132213033-2031332231130303-3332303033210201) |
| `https_management.advertise_on_sli_vip.tls_certificates.private_key.blindfold_secret_info.store_provider` | [https_management.advertise_on_sli_vip.tls_certificates.private_key.blindfold_secret_info.store_provider](data-sources--nfv_service--reference--group-002.md#canonical-1323121223331321-2020200230211112-3003301123100233-0130303003231020-2110101013333120-3111312231213220-0023001230001010-2101222331310012) |
| `https_management.advertise_on_sli_vip.tls_certificates.private_key.clear_secret_info` | [https_management.advertise_on_sli_vip.tls_certificates.private_key.clear_secret_info](data-sources--nfv_service--reference--group-002.md#canonical-3003133322202002-1213102113032213-2303212210103202-0010032222011122-1300221311222031-1133330111111210-2101230310032121-1200300120211103) |
| `https_management.advertise_on_sli_vip.tls_certificates.private_key.clear_secret_info.provider_ref` | [https_management.advertise_on_sli_vip.tls_certificates.private_key.clear_secret_info.provider_ref](data-sources--nfv_service--reference--group-002.md#canonical-2220300013322301-1031303333033030-3203012010023322-3320333300012310-1023211202132320-1031310202302212-1311030013021000-0113300033102301) |
| `https_management.advertise_on_sli_vip.tls_certificates.private_key.clear_secret_info.url` | [https_management.advertise_on_sli_vip.tls_certificates.private_key.clear_secret_info.url](data-sources--nfv_service--reference--group-002.md#canonical-3022100031020120-1211233320231230-2100323200110313-3033210313001013-1303213102322312-3102002020003301-0203102321001021-0131203123203001) |
| `https_management.advertise_on_sli_vip.tls_certificates.use_system_defaults` | [https_management.advertise_on_sli_vip.tls_certificates.use_system_defaults](data-sources--nfv_service--reference--group-002.md#canonical-0020213310000231-2233132011003122-2001133111200230-3033111323112101-0023321032022012-0220302303020021-0020010210310001-3233100221222123) |
| `https_management.advertise_on_sli_vip.tls_config` | [https_management.advertise_on_sli_vip.tls_config](data-sources--nfv_service--reference--group-002.md#canonical-0111120012132101-3232121210300120-3212320033032323-2212131200132202-2110121322023331-2030121310131123-1313310030210020-2320303320122310) |
| `https_management.advertise_on_sli_vip.tls_config.custom_security` | [https_management.advertise_on_sli_vip.tls_config.custom_security](data-sources--nfv_service--reference--group-002.md#canonical-1323203100003013-2120031013223203-2121212313211112-0101132213112302-0200302233311101-2130111231033012-1033310010023110-1031303113012210) |
| `https_management.advertise_on_sli_vip.tls_config.custom_security.cipher_suites` | [https_management.advertise_on_sli_vip.tls_config.custom_security.cipher_suites](data-sources--nfv_service--reference--group-002.md#canonical-3013102311030301-3020301121310222-3200112312300101-0021221311200022-3213112203031322-3011001220302232-1100013231120310-0332332302131331) |
| `https_management.advertise_on_sli_vip.tls_config.custom_security.max_version` | [https_management.advertise_on_sli_vip.tls_config.custom_security.max_version](data-sources--nfv_service--reference--group-002.md#canonical-1201100111002231-3121312012123012-2102230231333013-2102320210002330-1301210201332211-0031310013101311-1111111301001112-2211030232233230) |
| `https_management.advertise_on_sli_vip.tls_config.custom_security.min_version` | [https_management.advertise_on_sli_vip.tls_config.custom_security.min_version](data-sources--nfv_service--reference--group-002.md#canonical-2323123121233223-1121100202301211-2223103223301111-0111010102010001-2021212211020302-1201232002323032-2000033110221100-1103001030100312) |
| `https_management.advertise_on_sli_vip.tls_config.default_security` | [https_management.advertise_on_sli_vip.tls_config.default_security](data-sources--nfv_service--reference--group-002.md#canonical-0011111012213111-0231313131222030-0200233203223202-0131212311233033-1032121021312232-1203103102201110-0130131123103211-1123133323330221) |
| `https_management.advertise_on_sli_vip.tls_config.low_security` | [https_management.advertise_on_sli_vip.tls_config.low_security](data-sources--nfv_service--reference--group-002.md#canonical-0021201100033021-2022323301302233-2001011031322022-3000300013110112-3303310322313230-0332313022123133-0203013230111103-0031223331233113) |
| `https_management.advertise_on_sli_vip.tls_config.medium_security` | [https_management.advertise_on_sli_vip.tls_config.medium_security](data-sources--nfv_service--reference--group-002.md#canonical-1002321133101300-1322312322332303-3032232333223223-2010220113200231-2220012022220132-3113101121012023-2002003123201231-1123100032212010) |
| `https_management.advertise_on_sli_vip.use_mtls` | [https_management.advertise_on_sli_vip.use_mtls](data-sources--nfv_service--reference--group-002.md#canonical-1001022002323202-1003220103203221-3102321122222223-1220123331133033-0000000203210302-2311232223020231-3111330110201032-3021021110011332) |
| `https_management.advertise_on_sli_vip.use_mtls.client_certificate_optional` | [https_management.advertise_on_sli_vip.use_mtls.client_certificate_optional](data-sources--nfv_service--reference--group-002.md#canonical-1321000201030302-0302101210302121-3202123002320020-0031030032231130-3322011302221230-3010321322213001-0003331310310023-0303030303032112) |
| `https_management.advertise_on_sli_vip.use_mtls.crl` | [https_management.advertise_on_sli_vip.use_mtls.crl](data-sources--nfv_service--reference--group-002.md#canonical-2111311003313112-0212233103031023-3031003311030302-2101000110320333-3220033330211321-1303013332213120-1211021222001331-0210032203232210) |
| `https_management.advertise_on_sli_vip.use_mtls.crl.name` | [https_management.advertise_on_sli_vip.use_mtls.crl.name](data-sources--nfv_service--reference--group-002.md#canonical-0311310310131200-3003323300011133-3211003311302221-3000103220220102-1021130120012011-3300311102201331-0303122321213101-3123123022103121) |
| `https_management.advertise_on_sli_vip.use_mtls.crl.namespace` | [https_management.advertise_on_sli_vip.use_mtls.crl.namespace](data-sources--nfv_service--reference--group-002.md#canonical-3223312123303120-2102223013111003-3002311210103002-1000331300002211-0300022121110232-3022032131121120-1231320013000330-0300303030322101) |
| `https_management.advertise_on_sli_vip.use_mtls.crl.tenant` | [https_management.advertise_on_sli_vip.use_mtls.crl.tenant](data-sources--nfv_service--reference--group-002.md#canonical-1121131201333310-3021103220131320-2220213000111313-2332230122303000-1320310123302003-2103021001120101-1110123300011230-3003021330310132) |
| `https_management.advertise_on_sli_vip.use_mtls.no_crl` | [https_management.advertise_on_sli_vip.use_mtls.no_crl](data-sources--nfv_service--reference--group-002.md#canonical-1203102112003003-2222321112021301-2231202123113330-2201101232323102-0212220013233320-1121021123203311-3220332111033222-3302110132312200) |
| `https_management.advertise_on_sli_vip.use_mtls.trusted_ca` | [https_management.advertise_on_sli_vip.use_mtls.trusted_ca](data-sources--nfv_service--reference--group-002.md#canonical-1113102033001013-1230300323112202-2303102330011011-2130120110011203-0300033333023231-3201031330100333-3023202031232110-3322121020033103) |
| `https_management.advertise_on_sli_vip.use_mtls.trusted_ca.name` | [https_management.advertise_on_sli_vip.use_mtls.trusted_ca.name](data-sources--nfv_service--reference--group-002.md#canonical-1122032020201031-3031203120130222-3310201212101200-0302003332203231-0021001110101021-1223233210022331-1303002132011333-0122120223133213) |
| `https_management.advertise_on_sli_vip.use_mtls.trusted_ca.namespace` | [https_management.advertise_on_sli_vip.use_mtls.trusted_ca.namespace](data-sources--nfv_service--reference--group-002.md#canonical-0120223310210023-2023133302323012-1302221123102120-3131120031222320-0322322110130130-0133223000002303-2011200102301022-2033020210013211) |
| `https_management.advertise_on_sli_vip.use_mtls.trusted_ca.tenant` | [https_management.advertise_on_sli_vip.use_mtls.trusted_ca.tenant](data-sources--nfv_service--reference--group-002.md#canonical-2213122103030213-0013221003001121-1023032033231202-3233020111321300-3101332112323322-0331123011202210-2323012210322022-0230320020212330) |
| `https_management.advertise_on_sli_vip.use_mtls.trusted_ca_url` | [https_management.advertise_on_sli_vip.use_mtls.trusted_ca_url](data-sources--nfv_service--reference--group-002.md#canonical-1113312021332032-0322131313313120-1022030122113200-2311130121020202-0012301332211230-1120203221010310-0222203130131131-2122113131232031) |
| `https_management.advertise_on_sli_vip.use_mtls.xfcc_disabled` | [https_management.advertise_on_sli_vip.use_mtls.xfcc_disabled](data-sources--nfv_service--reference--group-002.md#canonical-0231200133003302-2200110003112212-0232101020203002-3333030301310333-1032130213023313-1310132100232200-2010131100031221-3020201033101333) |
| `https_management.advertise_on_sli_vip.use_mtls.xfcc_options` | [https_management.advertise_on_sli_vip.use_mtls.xfcc_options](data-sources--nfv_service--reference--group-002.md#canonical-1203011120223321-1122002020121213-0232211113032302-3002320021302233-2010033032012130-0002101221211220-0003003032332212-2123112233021230) |
| `https_management.advertise_on_sli_vip.use_mtls.xfcc_options.xfcc_header_elements` | [https_management.advertise_on_sli_vip.use_mtls.xfcc_options.xfcc_header_elements](data-sources--nfv_service--reference--group-002.md#canonical-1200202322022330-3131313001210130-0131110313131220-3121213132113322-1201132121032011-0300031110022303-3133120210103231-1011010112102021) |
| `https_management.advertise_on_slo_internet_vip` | [https_management.advertise_on_slo_internet_vip](data-sources--nfv_service--reference--group-002.md#canonical-1203311020012322-1033102332022113-3010133021301203-3032233002303212-1000302333022202-2130111233333110-1111002132103113-0233112103101201) |
| `https_management.advertise_on_slo_internet_vip.no_mtls` | [https_management.advertise_on_slo_internet_vip.no_mtls](data-sources--nfv_service--reference--group-002.md#canonical-3232111202132332-1232222320111223-1022120320323000-0232303031023312-1122022222320111-1012200212003222-3111023101120302-3210313323101303) |
| `https_management.advertise_on_slo_internet_vip.tls_certificates` | [https_management.advertise_on_slo_internet_vip.tls_certificates](data-sources--nfv_service--reference--group-002.md#canonical-1130033232300013-0312233123122301-3300210001110330-3322203312320000-0013310112310303-2132031322320023-2331102130030313-3133201223130301) |
| `https_management.advertise_on_slo_internet_vip.tls_certificates.certificate_url` | [https_management.advertise_on_slo_internet_vip.tls_certificates.certificate_url](data-sources--nfv_service--reference--group-002.md#canonical-1010202301331020-3122320111311201-2122232120002320-2213011123111012-3110201232021231-1030321021111022-2301302001101021-2122121232312122) |
| `https_management.advertise_on_slo_internet_vip.tls_certificates.custom_hash_algorithms` | [https_management.advertise_on_slo_internet_vip.tls_certificates.custom_hash_algorithms](data-sources--nfv_service--reference--group-002.md#canonical-3111312223003312-2332031022132100-0121333010211321-1300133203003213-3321221123030123-1003230312301233-1300123331023112-2231112232011300) |
| `https_management.advertise_on_slo_internet_vip.tls_certificates.custom_hash_algorithms.hash_algorithms` | [https_management.advertise_on_slo_internet_vip.tls_certificates.custom_hash_algorithms.hash_algorithms](data-sources--nfv_service--reference--group-002.md#canonical-3031202111312232-3233323230022211-2031321222202333-3231200310203321-3110321002230332-2132000213312020-0201103301012013-1230230001333202) |
| `https_management.advertise_on_slo_internet_vip.tls_certificates.description_spec` | [https_management.advertise_on_slo_internet_vip.tls_certificates.description_spec](data-sources--nfv_service--reference--group-002.md#canonical-3003233333100213-3013231033112131-1032121222102330-1213122122232301-3201330122030113-0123312021333302-3320020010002232-2100122222133231) |
| `https_management.advertise_on_slo_internet_vip.tls_certificates.disable_ocsp_stapling` | [https_management.advertise_on_slo_internet_vip.tls_certificates.disable_ocsp_stapling](data-sources--nfv_service--reference--group-002.md#canonical-2113310221322311-3000133132033220-2000130022203102-3121301220223331-0000301101033122-0121323121201131-2222222332022311-2320311110202332) |
| `https_management.advertise_on_slo_internet_vip.tls_certificates.private_key` | [https_management.advertise_on_slo_internet_vip.tls_certificates.private_key](data-sources--nfv_service--reference--group-002.md#canonical-3011002031002112-3303303102112202-3332200211030323-3100231213033022-2211321321302123-0020232011313020-1311113110331110-0213202101330031) |
| `https_management.advertise_on_slo_internet_vip.tls_certificates.private_key.blindfold_secret_info` | [https_management.advertise_on_slo_internet_vip.tls_certificates.private_key.blindfold_secret_info](data-sources--nfv_service--reference--group-002.md#canonical-1213212033222310-2323301103130310-2331013331321013-2100030220102033-0332320231201200-0220120131031000-0201332120110320-3330222303333111) |
| `https_management.advertise_on_slo_internet_vip.tls_certificates.private_key.blindfold_secret_info.decryption_provider` | [https_management.advertise_on_slo_internet_vip.tls_certificates.private_key.blindfold_secret_info.decryption_provider](data-sources--nfv_service--reference--group-002.md#canonical-0023023033103323-0011210333113002-3313110002022112-0313210010203221-0222201333100030-3312023130131003-3100212213030030-2131200223312111) |
| `https_management.advertise_on_slo_internet_vip.tls_certificates.private_key.blindfold_secret_info.location` | [https_management.advertise_on_slo_internet_vip.tls_certificates.private_key.blindfold_secret_info.location](data-sources--nfv_service--reference--group-002.md#canonical-3011121102330011-0132232233033110-1003332122000231-1023301012020020-2033301213201320-2231211010013200-1001321310310121-2303333232313020) |
| `https_management.advertise_on_slo_internet_vip.tls_certificates.private_key.blindfold_secret_info.store_provider` | [https_management.advertise_on_slo_internet_vip.tls_certificates.private_key.blindfold_secret_info.store_provider](data-sources--nfv_service--reference--group-002.md#canonical-2013221312203301-2111202220223331-3133201103210313-1333100102103331-0122000120010030-3120113121112333-2311123110222001-3102320311300303) |
| `https_management.advertise_on_slo_internet_vip.tls_certificates.private_key.clear_secret_info` | [https_management.advertise_on_slo_internet_vip.tls_certificates.private_key.clear_secret_info](data-sources--nfv_service--reference--group-002.md#canonical-1303210011000222-1110120000133001-3322101000312030-1100032221002100-3133022222210332-2023231201312131-2201331012013102-3332113321121310) |
| `https_management.advertise_on_slo_internet_vip.tls_certificates.private_key.clear_secret_info.provider_ref` | [https_management.advertise_on_slo_internet_vip.tls_certificates.private_key.clear_secret_info.provider_ref](data-sources--nfv_service--reference--group-002.md#canonical-1000002003101031-3200003330120130-0300311311301331-3020210232111202-3120221003101313-3020323013130031-0201203332331002-2131013011113301) |
| `https_management.advertise_on_slo_internet_vip.tls_certificates.private_key.clear_secret_info.url` | [https_management.advertise_on_slo_internet_vip.tls_certificates.private_key.clear_secret_info.url](data-sources--nfv_service--reference--group-002.md#canonical-0021321111030313-1321320233123130-3203012230300301-1303113301122101-3331300312030130-0323231031131022-0233002213201112-0031321203302230) |
| `https_management.advertise_on_slo_internet_vip.tls_certificates.use_system_defaults` | [https_management.advertise_on_slo_internet_vip.tls_certificates.use_system_defaults](data-sources--nfv_service--reference--group-002.md#canonical-2013231231231111-2002232233030233-3011211000112013-3130221030120032-1233013220130230-2311132113232223-1330132111111130-1233030210311301) |
| `https_management.advertise_on_slo_internet_vip.tls_config` | [https_management.advertise_on_slo_internet_vip.tls_config](data-sources--nfv_service--reference--group-002.md#canonical-2133101321102213-3301133322312003-0321023330322310-1231221102032100-0111222103330301-0121021033013033-3310020021031010-0311012233012233) |
| `https_management.advertise_on_slo_internet_vip.tls_config.custom_security` | [https_management.advertise_on_slo_internet_vip.tls_config.custom_security](data-sources--nfv_service--reference--group-002.md#canonical-1100210220010022-1333331213031102-2111010103133000-3331302233000033-3211032012102222-1022111322032302-0332303310331322-3032323202231223) |
| `https_management.advertise_on_slo_internet_vip.tls_config.custom_security.cipher_suites` | [https_management.advertise_on_slo_internet_vip.tls_config.custom_security.cipher_suites](data-sources--nfv_service--reference--group-002.md#canonical-2012313333233111-0301320331031001-0101120230312102-0223022022222111-2223001010301300-1133232313201313-0101011000212103-3301102330212110) |
| `https_management.advertise_on_slo_internet_vip.tls_config.custom_security.max_version` | [https_management.advertise_on_slo_internet_vip.tls_config.custom_security.max_version](data-sources--nfv_service--reference--group-002.md#canonical-0203313312110220-2030130022202300-0013111002113122-0301330101110311-0220112123012100-3313211221233120-3312121001320012-0220302222001202) |
| `https_management.advertise_on_slo_internet_vip.tls_config.custom_security.min_version` | [https_management.advertise_on_slo_internet_vip.tls_config.custom_security.min_version](data-sources--nfv_service--reference--group-002.md#canonical-0230122022302300-0032323100102022-3210231213213201-2112323132111131-1000011133300221-2000100322200230-0221313223022012-0213022122303013) |
| `https_management.advertise_on_slo_internet_vip.tls_config.default_security` | [https_management.advertise_on_slo_internet_vip.tls_config.default_security](data-sources--nfv_service--reference--group-002.md#canonical-3010213023313123-1232320013323033-0233033000211101-3302201302130303-1022103032320213-3123020133110310-2103112020120203-1331332312210200) |
| `https_management.advertise_on_slo_internet_vip.tls_config.low_security` | [https_management.advertise_on_slo_internet_vip.tls_config.low_security](data-sources--nfv_service--reference--group-002.md#canonical-3212111233100100-2312313110310021-0121302321322213-2123011031322322-1303222303123211-1001211032130330-3101101121230131-3022323133113323) |
| `https_management.advertise_on_slo_internet_vip.tls_config.medium_security` | [https_management.advertise_on_slo_internet_vip.tls_config.medium_security](data-sources--nfv_service--reference--group-002.md#canonical-3300030310231313-3300001210102222-3103200200003330-1332330123310210-0331021121303210-3231133003211312-2020300323113210-0001110133323331) |
| `https_management.advertise_on_slo_internet_vip.use_mtls` | [https_management.advertise_on_slo_internet_vip.use_mtls](data-sources--nfv_service--reference--group-002.md#canonical-3313100011123200-0111211110233233-3222320301030333-2130023333331312-1110122032320312-0023202002320032-1202031120002202-0100313303112131) |
| `https_management.advertise_on_slo_internet_vip.use_mtls.client_certificate_optional` | [https_management.advertise_on_slo_internet_vip.use_mtls.client_certificate_optional](data-sources--nfv_service--reference--group-002.md#canonical-1223302010310002-3032230321001023-0301213100122303-1033023312330103-0331220330130123-0303130003300022-2321120013331032-0231130202002012) |
| `https_management.advertise_on_slo_internet_vip.use_mtls.crl` | [https_management.advertise_on_slo_internet_vip.use_mtls.crl](data-sources--nfv_service--reference--group-002.md#canonical-1100102110000330-1313323030210113-0230232331201102-0220120030221021-2110121323303300-1213002011030110-2322030122222033-1331013123101020) |
| `https_management.advertise_on_slo_internet_vip.use_mtls.crl.name` | [https_management.advertise_on_slo_internet_vip.use_mtls.crl.name](data-sources--nfv_service--reference--group-002.md#canonical-1212132101303310-3213021103310002-2020030023020311-0222112213310303-1212033103033013-1202311022220121-1320000330103230-2000331011310212) |
| `https_management.advertise_on_slo_internet_vip.use_mtls.crl.namespace` | [https_management.advertise_on_slo_internet_vip.use_mtls.crl.namespace](data-sources--nfv_service--reference--group-002.md#canonical-2110121001033031-0113223121121100-2331130010202132-2313233202030100-0003222221331320-1201331211222101-0032330003121230-2311033330132203) |
| `https_management.advertise_on_slo_internet_vip.use_mtls.crl.tenant` | [https_management.advertise_on_slo_internet_vip.use_mtls.crl.tenant](data-sources--nfv_service--reference--group-002.md#canonical-2302001202321232-1332032300010312-3222221303311323-3233210013102213-1103332023133312-0203230101012331-0031320131331202-1111202030103113) |
| `https_management.advertise_on_slo_internet_vip.use_mtls.no_crl` | [https_management.advertise_on_slo_internet_vip.use_mtls.no_crl](data-sources--nfv_service--reference--group-002.md#canonical-3333013112033210-3002032201322122-2322121003103110-1032212133202000-0133113221031330-0203031313121321-2001220300333333-3123221001322323) |
| `https_management.advertise_on_slo_internet_vip.use_mtls.trusted_ca` | [https_management.advertise_on_slo_internet_vip.use_mtls.trusted_ca](data-sources--nfv_service--reference--group-002.md#canonical-3310200000001312-0022330101103213-0303021002213203-2121211321011120-3203010132122123-0103001113132022-3010212210212113-3101113111232010) |
| `https_management.advertise_on_slo_internet_vip.use_mtls.trusted_ca.name` | [https_management.advertise_on_slo_internet_vip.use_mtls.trusted_ca.name](data-sources--nfv_service--reference--group-002.md#canonical-3020203000230302-0223300123232002-2203133300000323-1200313032102111-3331212230022332-1101221033113220-1010310212032200-2331311302302322) |
| `https_management.advertise_on_slo_internet_vip.use_mtls.trusted_ca.namespace` | [https_management.advertise_on_slo_internet_vip.use_mtls.trusted_ca.namespace](data-sources--nfv_service--reference--group-002.md#canonical-1133300222301111-1010023010333102-0333201333311301-0021221102202321-3013001102011312-0331332033103030-2113321100103112-2303120320112223) |
| `https_management.advertise_on_slo_internet_vip.use_mtls.trusted_ca.tenant` | [https_management.advertise_on_slo_internet_vip.use_mtls.trusted_ca.tenant](data-sources--nfv_service--reference--group-002.md#canonical-0030302221012003-2033232303003110-2210312220121313-3113132022331022-2333121200122321-1131201021232012-1011210323103103-2111303221023111) |
| `https_management.advertise_on_slo_internet_vip.use_mtls.trusted_ca_url` | [https_management.advertise_on_slo_internet_vip.use_mtls.trusted_ca_url](data-sources--nfv_service--reference--group-002.md#canonical-0203300203210101-3012300320213023-1310103332331102-3022011013230211-1112320211121220-0303121021322102-0220012112221322-0021000211131202) |
| `https_management.advertise_on_slo_internet_vip.use_mtls.xfcc_disabled` | [https_management.advertise_on_slo_internet_vip.use_mtls.xfcc_disabled](data-sources--nfv_service--reference--group-002.md#canonical-3223123030001322-1031301010313102-0001110332031221-3113330211233311-2032011301320130-0202200213120013-0000322022221233-1222331202301322) |
| `https_management.advertise_on_slo_internet_vip.use_mtls.xfcc_options` | [https_management.advertise_on_slo_internet_vip.use_mtls.xfcc_options](data-sources--nfv_service--reference--group-003.md#canonical-0212332210313031-0032001310031302-1113322132113321-0121200130233013-0333220210322331-2020333113033223-0320213233100031-1332233013120303) |
| `https_management.advertise_on_slo_internet_vip.use_mtls.xfcc_options.xfcc_header_elements` | [https_management.advertise_on_slo_internet_vip.use_mtls.xfcc_options.xfcc_header_elements](data-sources--nfv_service--reference--group-003.md#canonical-2011010030320310-0101110311022121-1203120331220103-0313110322223331-3021110313131103-0333231301121101-2203330103303000-3001113320313312) |
| `https_management.advertise_on_slo_sli` | [https_management.advertise_on_slo_sli](data-sources--nfv_service--reference--group-003.md#canonical-3131211132010013-0223123101133212-0231212131233312-1000123022221333-1001103213230001-2301212201113200-2033231301231311-2103100233111011) |
| `https_management.advertise_on_slo_sli.no_mtls` | [https_management.advertise_on_slo_sli.no_mtls](data-sources--nfv_service--reference--group-003.md#canonical-1321122231003130-1300322123110300-1020220201321121-2333132000103302-2013131323023332-2231220132132223-2320100200322300-0331220031323033) |
| `https_management.advertise_on_slo_sli.tls_certificates` | [https_management.advertise_on_slo_sli.tls_certificates](data-sources--nfv_service--reference--group-003.md#canonical-3222023331233302-0313302322333022-0221212310233102-3200222230312110-2123022333032330-3113301231331322-1332033212110313-3232022021012231) |
| `https_management.advertise_on_slo_sli.tls_certificates.certificate_url` | [https_management.advertise_on_slo_sli.tls_certificates.certificate_url](data-sources--nfv_service--reference--group-003.md#canonical-0010332003012210-0331200330210321-2133011201123012-1330310203330220-1303223200031231-1101013022001201-2120122232212023-2103320003312233) |
| `https_management.advertise_on_slo_sli.tls_certificates.custom_hash_algorithms` | [https_management.advertise_on_slo_sli.tls_certificates.custom_hash_algorithms](data-sources--nfv_service--reference--group-003.md#canonical-1321301323022332-0030100002131133-0223033322200000-0122311100311112-3023232321332103-3001020333313011-2211002022332312-3202213321030212) |
| `https_management.advertise_on_slo_sli.tls_certificates.custom_hash_algorithms.hash_algorithms` | [https_management.advertise_on_slo_sli.tls_certificates.custom_hash_algorithms.hash_algorithms](data-sources--nfv_service--reference--group-003.md#canonical-1121330301113311-2130213323201030-0002200303001022-3222110230301233-0132221022001023-2023301220303132-0110002303232101-2333210313200310) |
| `https_management.advertise_on_slo_sli.tls_certificates.description_spec` | [https_management.advertise_on_slo_sli.tls_certificates.description_spec](data-sources--nfv_service--reference--group-003.md#canonical-3233331210221210-1201302113031023-2011220312131102-0011113211323020-0012002222111310-3021113311101022-1300300231120113-3112021202213123) |
| `https_management.advertise_on_slo_sli.tls_certificates.disable_ocsp_stapling` | [https_management.advertise_on_slo_sli.tls_certificates.disable_ocsp_stapling](data-sources--nfv_service--reference--group-003.md#canonical-0021311210231103-2310002322101333-1130131301313130-1313223220220312-1001320000200121-1332133220111122-1210221332301210-0002003223230022) |
| `https_management.advertise_on_slo_sli.tls_certificates.private_key` | [https_management.advertise_on_slo_sli.tls_certificates.private_key](data-sources--nfv_service--reference--group-003.md#canonical-2010232010033330-2133223331322131-1002033103313321-1123230100033321-0122223001030033-2200202013012020-3323321202130202-1212030312122302) |
| `https_management.advertise_on_slo_sli.tls_certificates.private_key.blindfold_secret_info` | [https_management.advertise_on_slo_sli.tls_certificates.private_key.blindfold_secret_info](data-sources--nfv_service--reference--group-003.md#canonical-3013111330211010-1310323113231010-0020212013333233-0332213130210110-1223122312212233-2102300100220321-0022003310213020-2113022201003013) |
| `https_management.advertise_on_slo_sli.tls_certificates.private_key.blindfold_secret_info.decryption_provider` | [https_management.advertise_on_slo_sli.tls_certificates.private_key.blindfold_secret_info.decryption_provider](data-sources--nfv_service--reference--group-003.md#canonical-3112321331322210-0013020320113211-2033103220133202-1313110010121310-1213223023112010-2210202010233110-3301023021211121-3210200330003103) |
| `https_management.advertise_on_slo_sli.tls_certificates.private_key.blindfold_secret_info.location` | [https_management.advertise_on_slo_sli.tls_certificates.private_key.blindfold_secret_info.location](data-sources--nfv_service--reference--group-003.md#canonical-1122212111011122-0213031131013010-2201000331031110-1102010000012323-3222222100030012-0200132132102021-2331123203020011-3233131333033131) |
| `https_management.advertise_on_slo_sli.tls_certificates.private_key.blindfold_secret_info.store_provider` | [https_management.advertise_on_slo_sli.tls_certificates.private_key.blindfold_secret_info.store_provider](data-sources--nfv_service--reference--group-003.md#canonical-0212023201302100-3310001221033300-0002330220110011-2300232031213101-0230110331230130-3202131313013112-1102303100303210-1000331211332323) |
| `https_management.advertise_on_slo_sli.tls_certificates.private_key.clear_secret_info` | [https_management.advertise_on_slo_sli.tls_certificates.private_key.clear_secret_info](data-sources--nfv_service--reference--group-003.md#canonical-3102200103003321-1303323013231331-1122013110333002-1210301123122030-3033313101202132-1311301021302112-1301022033310023-1103302103201223) |
| `https_management.advertise_on_slo_sli.tls_certificates.private_key.clear_secret_info.provider_ref` | [https_management.advertise_on_slo_sli.tls_certificates.private_key.clear_secret_info.provider_ref](data-sources--nfv_service--reference--group-003.md#canonical-2022120033333001-2123301212310101-3002111223030100-1322023022301333-2200123231203100-2320233223002123-1111110002313230-0332102320130020) |
| `https_management.advertise_on_slo_sli.tls_certificates.private_key.clear_secret_info.url` | [https_management.advertise_on_slo_sli.tls_certificates.private_key.clear_secret_info.url](data-sources--nfv_service--reference--group-003.md#canonical-3221220200233112-0200333223112331-2211201121120200-1120321332203232-2131233202221200-3031212102000303-3023331030323300-3232112200333201) |
| `https_management.advertise_on_slo_sli.tls_certificates.use_system_defaults` | [https_management.advertise_on_slo_sli.tls_certificates.use_system_defaults](data-sources--nfv_service--reference--group-003.md#canonical-1011300131103331-1101222232230232-1112012020002203-2121322200203100-1331102312310223-0320112330212201-2223113220331301-1232301233020200) |
| `https_management.advertise_on_slo_sli.tls_config` | [https_management.advertise_on_slo_sli.tls_config](data-sources--nfv_service--reference--group-003.md#canonical-0203331200121210-1010331113302133-1223110132233332-3101312231211102-3203121012112221-2220302010011030-0222212022112201-2332333300213223) |
| `https_management.advertise_on_slo_sli.tls_config.custom_security` | [https_management.advertise_on_slo_sli.tls_config.custom_security](data-sources--nfv_service--reference--group-003.md#canonical-2003310123302022-3031022321010231-1333130130333103-2232003202202030-1310013131300333-1200303020223010-0013332312011010-1220033210010011) |
| `https_management.advertise_on_slo_sli.tls_config.custom_security.cipher_suites` | [https_management.advertise_on_slo_sli.tls_config.custom_security.cipher_suites](data-sources--nfv_service--reference--group-003.md#canonical-3131101130000202-1333202011001023-2213000332123230-3321210310321211-0120112321312202-1031011320133203-3210023220320021-3013131113001013) |
| `https_management.advertise_on_slo_sli.tls_config.custom_security.max_version` | [https_management.advertise_on_slo_sli.tls_config.custom_security.max_version](data-sources--nfv_service--reference--group-003.md#canonical-2011213120131012-3003130120232020-0301101201220211-3320111022223203-1231220321103033-0001210213001233-2202331012332113-2033122020311232) |
| `https_management.advertise_on_slo_sli.tls_config.custom_security.min_version` | [https_management.advertise_on_slo_sli.tls_config.custom_security.min_version](data-sources--nfv_service--reference--group-003.md#canonical-3212120022230003-2200332320323032-0201133302320001-0101233322032113-2101231010130023-3212001032220013-2131311022131322-0213131101133122) |
| `https_management.advertise_on_slo_sli.tls_config.default_security` | [https_management.advertise_on_slo_sli.tls_config.default_security](data-sources--nfv_service--reference--group-003.md#canonical-0121021131322000-2220310131010230-0212113211102310-3322020111010211-3103232102321111-1101211203003332-2000103012223301-0223133231001233) |
| `https_management.advertise_on_slo_sli.tls_config.low_security` | [https_management.advertise_on_slo_sli.tls_config.low_security](data-sources--nfv_service--reference--group-003.md#canonical-3221103332010201-3130003221220300-0020020213203002-0303311312032012-2033100312303223-2220333303122320-3321012200331030-0330211212100212) |
| `https_management.advertise_on_slo_sli.tls_config.medium_security` | [https_management.advertise_on_slo_sli.tls_config.medium_security](data-sources--nfv_service--reference--group-003.md#canonical-3300332300100133-3110222201010132-0030300131320200-0213033221101010-1133210000223322-2303213302330013-0021322331230331-2201021323230302) |
| `https_management.advertise_on_slo_sli.use_mtls` | [https_management.advertise_on_slo_sli.use_mtls](data-sources--nfv_service--reference--group-003.md#canonical-0220301210013231-2031133121322131-2321211303313123-2321200230333121-2131103031021331-3332000122011320-2322321100210003-3102211223333212) |
| `https_management.advertise_on_slo_sli.use_mtls.client_certificate_optional` | [https_management.advertise_on_slo_sli.use_mtls.client_certificate_optional](data-sources--nfv_service--reference--group-003.md#canonical-1123210220103213-0002301223201301-1303320332033302-0021312232213330-0233022103102312-1011211113003232-0302303323010010-3231032001331011) |
| `https_management.advertise_on_slo_sli.use_mtls.crl` | [https_management.advertise_on_slo_sli.use_mtls.crl](data-sources--nfv_service--reference--group-003.md#canonical-1030211031122330-3300021113200032-1201113121300013-0320000010120221-0011312021003132-0321030112211320-2221313231201121-1022103031202123) |
| `https_management.advertise_on_slo_sli.use_mtls.crl.name` | [https_management.advertise_on_slo_sli.use_mtls.crl.name](data-sources--nfv_service--reference--group-003.md#canonical-3312103323122202-3101213120330233-1010033020031001-2331030331112122-2333320112020201-3112131333200211-0002132031223322-0110203220202131) |
| `https_management.advertise_on_slo_sli.use_mtls.crl.namespace` | [https_management.advertise_on_slo_sli.use_mtls.crl.namespace](data-sources--nfv_service--reference--group-003.md#canonical-2113331133210320-1212322022132230-1310002023132313-1133202322220032-1103311031131212-3332031101013323-2030003000300101-1230022011122011) |
| `https_management.advertise_on_slo_sli.use_mtls.crl.tenant` | [https_management.advertise_on_slo_sli.use_mtls.crl.tenant](data-sources--nfv_service--reference--group-003.md#canonical-2010303300211231-3130302310322031-0310302211202010-0300023323222232-2032232221323020-1121320202132131-1203330203132303-2011333123332000) |
| `https_management.advertise_on_slo_sli.use_mtls.no_crl` | [https_management.advertise_on_slo_sli.use_mtls.no_crl](data-sources--nfv_service--reference--group-003.md#canonical-1103022002012221-1011003023310022-3101111122130111-0223010003023020-1131320030220203-2213002111033012-3022033003111112-3210211330010030) |
| `https_management.advertise_on_slo_sli.use_mtls.trusted_ca` | [https_management.advertise_on_slo_sli.use_mtls.trusted_ca](data-sources--nfv_service--reference--group-003.md#canonical-2323320112132010-2030320110022023-1023110313032123-1031010112113332-0030011211211310-0202033011120201-3203030032303120-0011231232111112) |
| `https_management.advertise_on_slo_sli.use_mtls.trusted_ca.name` | [https_management.advertise_on_slo_sli.use_mtls.trusted_ca.name](data-sources--nfv_service--reference--group-003.md#canonical-1032102323222210-1021132002133311-2332331221320313-3112230213220211-0231211103021231-0221022210230000-0201223010231330-3023132223313022) |
| `https_management.advertise_on_slo_sli.use_mtls.trusted_ca.namespace` | [https_management.advertise_on_slo_sli.use_mtls.trusted_ca.namespace](data-sources--nfv_service--reference--group-003.md#canonical-0331311121231131-3031113302131210-0002203003122023-3010010311021013-3232000030100113-2133120100001003-0021023230013000-2203213121133011) |
| `https_management.advertise_on_slo_sli.use_mtls.trusted_ca.tenant` | [https_management.advertise_on_slo_sli.use_mtls.trusted_ca.tenant](data-sources--nfv_service--reference--group-003.md#canonical-3232133223312131-3330132031300022-1022212102011202-0112131330013300-2013020010023302-3112012212231022-1002213020032122-1033311111233100) |
| `https_management.advertise_on_slo_sli.use_mtls.trusted_ca_url` | [https_management.advertise_on_slo_sli.use_mtls.trusted_ca_url](data-sources--nfv_service--reference--group-003.md#canonical-3332101211031221-3302201201320010-2300211121103020-3311210223133013-3013231103121313-3321233213233231-0003001010132331-1122330310010301) |
| `https_management.advertise_on_slo_sli.use_mtls.xfcc_disabled` | [https_management.advertise_on_slo_sli.use_mtls.xfcc_disabled](data-sources--nfv_service--reference--group-003.md#canonical-3103311232022002-2312013011333210-3003200010031311-0203032221101331-1233201002213213-0300130200123013-2132030112233220-0013212223310122) |
| `https_management.advertise_on_slo_sli.use_mtls.xfcc_options` | [https_management.advertise_on_slo_sli.use_mtls.xfcc_options](data-sources--nfv_service--reference--group-003.md#canonical-1302302203011030-0000223002133131-3223123331103311-3202223313033231-1130130330313121-2311313031100212-2111122111101000-1001121202230323) |
| `https_management.advertise_on_slo_sli.use_mtls.xfcc_options.xfcc_header_elements` | [https_management.advertise_on_slo_sli.use_mtls.xfcc_options.xfcc_header_elements](data-sources--nfv_service--reference--group-003.md#canonical-3123223231030223-0000100000012111-0132202320002301-1322331223213123-3110221033012013-2122101312300132-2013202111000111-3012123101202332) |
| `https_management.advertise_on_slo_vip` | [https_management.advertise_on_slo_vip](data-sources--nfv_service--reference--group-003.md#canonical-1201300311232210-3120333123332212-2021000011210110-3212320212312103-0200203032021102-1102331330302300-0230103102031032-3001223111313130) |
| `https_management.advertise_on_slo_vip.no_mtls` | [https_management.advertise_on_slo_vip.no_mtls](data-sources--nfv_service--reference--group-003.md#canonical-1220200323331332-3032011231113212-1312000202230132-1011301100021333-0033311332023202-0111021031311213-1130311120121212-0300120223131223) |
| `https_management.advertise_on_slo_vip.tls_certificates` | [https_management.advertise_on_slo_vip.tls_certificates](data-sources--nfv_service--reference--group-003.md#canonical-0121122103200222-1012011112210020-3110030201112113-1012023330221331-3211121123032211-3233211021301202-2123020032213103-2022312131000303) |
| `https_management.advertise_on_slo_vip.tls_certificates.certificate_url` | [https_management.advertise_on_slo_vip.tls_certificates.certificate_url](data-sources--nfv_service--reference--group-003.md#canonical-0322323131220131-0002320112021220-3232331022302301-3103332210210312-2301321223012031-0010200132003211-0123213200211031-2003221121202331) |
| `https_management.advertise_on_slo_vip.tls_certificates.custom_hash_algorithms` | [https_management.advertise_on_slo_vip.tls_certificates.custom_hash_algorithms](data-sources--nfv_service--reference--group-003.md#canonical-1221210211321123-0213132013332310-0012011303330123-2232330212002132-3212320012133301-3022110111320211-1131031333300023-0210010213323010) |
| `https_management.advertise_on_slo_vip.tls_certificates.custom_hash_algorithms.hash_algorithms` | [https_management.advertise_on_slo_vip.tls_certificates.custom_hash_algorithms.hash_algorithms](data-sources--nfv_service--reference--group-003.md#canonical-2302211003133230-3221322033323131-3003230303233231-2003232320002003-2011002333101230-1212200322011312-0210003220303101-2321132113312333) |
| `https_management.advertise_on_slo_vip.tls_certificates.description_spec` | [https_management.advertise_on_slo_vip.tls_certificates.description_spec](data-sources--nfv_service--reference--group-003.md#canonical-3123200033312303-2110122033222221-2331313212332013-3221231103322023-0022203331020131-0132232310003130-0232203022021221-3312230311310233) |
| `https_management.advertise_on_slo_vip.tls_certificates.disable_ocsp_stapling` | [https_management.advertise_on_slo_vip.tls_certificates.disable_ocsp_stapling](data-sources--nfv_service--reference--group-003.md#canonical-2023321032211332-0103323222331103-2030223023200111-0231130223102001-3230000021032000-3122211131230232-0111220220031111-2122010233012110) |
| `https_management.advertise_on_slo_vip.tls_certificates.private_key` | [https_management.advertise_on_slo_vip.tls_certificates.private_key](data-sources--nfv_service--reference--group-003.md#canonical-3123231300111133-3110002301033002-3123233302201310-0310111320110032-2002320013002103-1032022011231322-2010301212332331-3211212321122010) |
| `https_management.advertise_on_slo_vip.tls_certificates.private_key.blindfold_secret_info` | [https_management.advertise_on_slo_vip.tls_certificates.private_key.blindfold_secret_info](data-sources--nfv_service--reference--group-003.md#canonical-3011320311322121-1221001133201211-0330102213303330-1230130220032130-2221130002331221-1111201303222201-3302310010021133-0211023002331111) |
| `https_management.advertise_on_slo_vip.tls_certificates.private_key.blindfold_secret_info.decryption_provider` | [https_management.advertise_on_slo_vip.tls_certificates.private_key.blindfold_secret_info.decryption_provider](data-sources--nfv_service--reference--group-003.md#canonical-3122211320123100-0333111203023112-2022331012133010-1030202301011023-2230330131321032-3211010302120120-3113213011231130-0212110001211212) |
| `https_management.advertise_on_slo_vip.tls_certificates.private_key.blindfold_secret_info.location` | [https_management.advertise_on_slo_vip.tls_certificates.private_key.blindfold_secret_info.location](data-sources--nfv_service--reference--group-003.md#canonical-2201302300211210-1001313133331333-0222223200112033-2030133101323223-0021101303313032-3211132022032203-2331022010123131-1030111310232011) |
| `https_management.advertise_on_slo_vip.tls_certificates.private_key.blindfold_secret_info.store_provider` | [https_management.advertise_on_slo_vip.tls_certificates.private_key.blindfold_secret_info.store_provider](data-sources--nfv_service--reference--group-003.md#canonical-1332200202330211-2033111200022200-2121333221212013-3212321212221231-1202130312133021-2210032201203233-0032221212011220-1200133112331210) |
| `https_management.advertise_on_slo_vip.tls_certificates.private_key.clear_secret_info` | [https_management.advertise_on_slo_vip.tls_certificates.private_key.clear_secret_info](data-sources--nfv_service--reference--group-003.md#canonical-2321112223130110-1300120031131131-1330232121012122-3330303311000200-0131011223223121-2130122313300320-2232112032103203-1322233313303312) |
| `https_management.advertise_on_slo_vip.tls_certificates.private_key.clear_secret_info.provider_ref` | [https_management.advertise_on_slo_vip.tls_certificates.private_key.clear_secret_info.provider_ref](data-sources--nfv_service--reference--group-003.md#canonical-3202321232020211-1123033003113021-1333000330131303-3322112112132202-3303130122320112-0030300132300122-2112222303210023-0111021001211013) |
| `https_management.advertise_on_slo_vip.tls_certificates.private_key.clear_secret_info.url` | [https_management.advertise_on_slo_vip.tls_certificates.private_key.clear_secret_info.url](data-sources--nfv_service--reference--group-003.md#canonical-0121211120313132-1121222130130003-3223311113021333-2221102312001112-0033300113003032-0210022111223333-1212111120011322-1331021011132120) |
| `https_management.advertise_on_slo_vip.tls_certificates.use_system_defaults` | [https_management.advertise_on_slo_vip.tls_certificates.use_system_defaults](data-sources--nfv_service--reference--group-003.md#canonical-2113020101012132-1303202132222230-1102010311001030-2033300013033310-2101111202131100-1232223233133230-2320013211022331-0123322321000103) |
| `https_management.advertise_on_slo_vip.tls_config` | [https_management.advertise_on_slo_vip.tls_config](data-sources--nfv_service--reference--group-003.md#canonical-0331311322030123-3123133111123223-1321020012313323-3133130002311013-1233121302202011-0213001010011132-2011231203312111-1113211312013310) |
| `https_management.advertise_on_slo_vip.tls_config.custom_security` | [https_management.advertise_on_slo_vip.tls_config.custom_security](data-sources--nfv_service--reference--group-003.md#canonical-3302111021202022-3232013023000120-1321033123102103-3002320312322001-3003112130302021-3202310211000131-2011022332212301-0211232211010110) |
| `https_management.advertise_on_slo_vip.tls_config.custom_security.cipher_suites` | [https_management.advertise_on_slo_vip.tls_config.custom_security.cipher_suites](data-sources--nfv_service--reference--group-003.md#canonical-0120031320113330-1320031212132112-1303331130121003-3300022211100111-2000131300102320-3112301022100211-3111010033201332-2010123310330023) |
| `https_management.advertise_on_slo_vip.tls_config.custom_security.max_version` | [https_management.advertise_on_slo_vip.tls_config.custom_security.max_version](data-sources--nfv_service--reference--group-003.md#canonical-2311111223013322-2322021130101320-0210000303302121-0233011213331320-2020221233213003-3131111001022032-3210231113331221-0130312220310333) |
| `https_management.advertise_on_slo_vip.tls_config.custom_security.min_version` | [https_management.advertise_on_slo_vip.tls_config.custom_security.min_version](data-sources--nfv_service--reference--group-003.md#canonical-2102113320230302-2212022233111301-1220300111302322-0221023133011023-2012033122312100-1330313131210110-2332323222011130-3111000130030010) |
| `https_management.advertise_on_slo_vip.tls_config.default_security` | [https_management.advertise_on_slo_vip.tls_config.default_security](data-sources--nfv_service--reference--group-003.md#canonical-2200010202331020-1200013310313322-2313030002130232-0121322330210222-2312122231202323-0233333320220120-3120103220121320-3200122331130231) |
| `https_management.advertise_on_slo_vip.tls_config.low_security` | [https_management.advertise_on_slo_vip.tls_config.low_security](data-sources--nfv_service--reference--group-003.md#canonical-3133102031001330-1133232102133002-0012111012020211-0311202012223202-3000223300202331-3322112211133232-1311122002202211-3131322113203211) |
| `https_management.advertise_on_slo_vip.tls_config.medium_security` | [https_management.advertise_on_slo_vip.tls_config.medium_security](data-sources--nfv_service--reference--group-003.md#canonical-3113230120313131-0002323223202320-1213312022212221-1013101021213222-3230022302002201-0230310131221212-0133031022000120-2120230100313321) |
| `https_management.advertise_on_slo_vip.use_mtls` | [https_management.advertise_on_slo_vip.use_mtls](data-sources--nfv_service--reference--group-003.md#canonical-0000010113012233-0231013031313230-1002310210333220-2221312211012332-1102321002200313-1222223113011310-1202220320011330-1012321200301023) |
| `https_management.advertise_on_slo_vip.use_mtls.client_certificate_optional` | [https_management.advertise_on_slo_vip.use_mtls.client_certificate_optional](data-sources--nfv_service--reference--group-003.md#canonical-1213101133011003-2002203131213231-2311122102010132-1203231100131202-3310233000023331-1201103020233031-1231332201313122-2332023302003331) |
| `https_management.advertise_on_slo_vip.use_mtls.crl` | [https_management.advertise_on_slo_vip.use_mtls.crl](data-sources--nfv_service--reference--group-003.md#canonical-2121101103120003-2101323332022131-0213220111201201-2013003331111312-3230203201101023-1231330220300233-3231122302111103-3213000123010131) |
| `https_management.advertise_on_slo_vip.use_mtls.crl.name` | [https_management.advertise_on_slo_vip.use_mtls.crl.name](data-sources--nfv_service--reference--group-003.md#canonical-0320312210333311-1031213201120212-0220222031023031-3130311013332213-1320331230310131-1130301111013212-2332202020221213-2222030322202103) |
| `https_management.advertise_on_slo_vip.use_mtls.crl.namespace` | [https_management.advertise_on_slo_vip.use_mtls.crl.namespace](data-sources--nfv_service--reference--group-003.md#canonical-3123021301121120-1133310312200012-2211000301302111-3131031121121123-2302022203202322-3210130100011230-3213312212032101-2113332001201010) |
| `https_management.advertise_on_slo_vip.use_mtls.crl.tenant` | [https_management.advertise_on_slo_vip.use_mtls.crl.tenant](data-sources--nfv_service--reference--group-003.md#canonical-3020013210220320-3001003320133310-3133021231110233-0112300121232311-2313332301113311-0203330102001120-3012222110102230-1222332102231302) |
| `https_management.advertise_on_slo_vip.use_mtls.no_crl` | [https_management.advertise_on_slo_vip.use_mtls.no_crl](data-sources--nfv_service--reference--group-003.md#canonical-2011333123320013-1301332010323222-3103233230311201-0231020002332200-1111021120023211-0212121131213112-3131122111310211-3003203112131111) |
| `https_management.advertise_on_slo_vip.use_mtls.trusted_ca` | [https_management.advertise_on_slo_vip.use_mtls.trusted_ca](data-sources--nfv_service--reference--group-003.md#canonical-3122023333112123-2322332100022023-2000333101313100-2023211212023330-0230333331231000-0213323220300001-2202303131202211-3201231330033322) |
| `https_management.advertise_on_slo_vip.use_mtls.trusted_ca.name` | [https_management.advertise_on_slo_vip.use_mtls.trusted_ca.name](data-sources--nfv_service--reference--group-003.md#canonical-0221201022212102-3001301113332211-0202200213223121-2233300031221033-1333010302032022-3113000100333201-2120200131321332-2301023010122311) |
| `https_management.advertise_on_slo_vip.use_mtls.trusted_ca.namespace` | [https_management.advertise_on_slo_vip.use_mtls.trusted_ca.namespace](data-sources--nfv_service--reference--group-003.md#canonical-2022032030323020-2130120113122112-2211203120220021-2323321120211313-2010111312321102-1130000120103000-1111110033021002-3010333333323222) |
| `https_management.advertise_on_slo_vip.use_mtls.trusted_ca.tenant` | [https_management.advertise_on_slo_vip.use_mtls.trusted_ca.tenant](data-sources--nfv_service--reference--group-003.md#canonical-2133012010032133-0310333111101020-1120111123013212-1203210023103203-0110312332231102-1133023233311000-0313202223100221-0303323302313313) |
| `https_management.advertise_on_slo_vip.use_mtls.trusted_ca_url` | [https_management.advertise_on_slo_vip.use_mtls.trusted_ca_url](data-sources--nfv_service--reference--group-003.md#canonical-1132330210130302-1233232301102303-1302303010333330-3223321332130332-3032023003302020-3330011023210031-3222310033302102-3003133031123101) |
| `https_management.advertise_on_slo_vip.use_mtls.xfcc_disabled` | [https_management.advertise_on_slo_vip.use_mtls.xfcc_disabled](data-sources--nfv_service--reference--group-003.md#canonical-2300103102230012-0003231020030332-0303122002220110-3123230023101003-2211231120002021-1233321033332121-3302232022232311-1100032232321302) |
| `https_management.advertise_on_slo_vip.use_mtls.xfcc_options` | [https_management.advertise_on_slo_vip.use_mtls.xfcc_options](data-sources--nfv_service--reference--group-003.md#canonical-2230321013321121-2230030032332200-0020111102220121-3230221201200123-2031323031111213-0313220102332003-2300330110320103-0310032301321230) |
| `https_management.advertise_on_slo_vip.use_mtls.xfcc_options.xfcc_header_elements` | [https_management.advertise_on_slo_vip.use_mtls.xfcc_options.xfcc_header_elements](data-sources--nfv_service--reference--group-003.md#canonical-3301020001203301-1120002212030212-2212021121102322-2223011023112321-0120333223221333-0022123233322032-2101223223231220-2233230321002032) |
| `https_management.default_https_port` | [https_management.default_https_port](data-sources--nfv_service--reference--group-003.md#canonical-1031231313033223-2323221332031020-3012130221213012-2133111033212130-2221030321332220-1012301313121210-2203113230230000-3112120101323131) |
| `https_management.domain_suffix` | [https_management.domain_suffix](data-sources--nfv_service--reference--group-002.md#canonical-0131010220012321-3102000010123201-2220322013303211-1011331201113303-3103120203202031-0123113331000302-2101303030002032-0203102112331330) |
| `https_management.https_port` | [https_management.https_port](data-sources--nfv_service--reference--group-002.md#canonical-1313122300232232-1322213101201203-3100210023023211-2111223110011030-1211203302100102-2211012120221221-1103223030322320-1213013020321223) |
| `id` | [id](data-sources--nfv_service--reference--group-001.md#canonical-3010132210033022-2121332002321211-0011101120213002-1122103003320001-0100212223322322-0203302332101331-2123211012210311-3332022223100230) |
| `labels` | [labels](data-sources--nfv_service--reference--group-001.md#canonical-0011002330323300-2232322212023333-2203202300321130-3111000101202103-2323320322031122-0033122000210232-2203102320001121-1033312130312122) |
| `name` | [name](data-sources--nfv_service--reference--group-001.md#canonical-0323220030233032-3210121100100012-3321333223121023-1213311013233123-0332000032210012-2312133323332212-2012321010011122-0301313203120111) |
| `namespace` | [namespace](data-sources--nfv_service--reference--group-001.md#canonical-2310122223303213-0030303221000210-2220113231313200-0230103312203301-1022221200303123-0331221220210132-1032110132003001-1330231330302100) |
| `palo_alto_fw_service` | [palo_alto_fw_service](data-sources--nfv_service--reference--group-003.md#canonical-2212323110330013-3101222023103001-2303100101011111-2301211332231032-2320133322031322-3100212302103322-2302012130013323-3313030213120310) |
| `palo_alto_fw_service.auto_setup` | [palo_alto_fw_service.auto_setup](data-sources--nfv_service--reference--group-003.md#canonical-2011011022311021-3012212223133002-0201233101202032-0133213022203213-3200021112003023-1223223003210011-0133222313233113-3212333102110333) |
| `palo_alto_fw_service.auto_setup.admin_password` | [palo_alto_fw_service.auto_setup.admin_password](data-sources--nfv_service--reference--group-003.md#canonical-0020233003201323-0221001312103003-1232123332133003-1202032301221022-1302202211220011-1310310203301000-3333323133313032-2122012111230330) |
| `palo_alto_fw_service.auto_setup.admin_password.blindfold_secret_info` | [palo_alto_fw_service.auto_setup.admin_password.blindfold_secret_info](data-sources--nfv_service--reference--group-003.md#canonical-1023123103310310-3310102330023203-2331103030230201-2003313021311303-0101313032120220-2100011011322303-2111231332332333-1102113111103020) |
| `palo_alto_fw_service.auto_setup.admin_password.blindfold_secret_info.decryption_provider` | [palo_alto_fw_service.auto_setup.admin_password.blindfold_secret_info.decryption_provider](data-sources--nfv_service--reference--group-003.md#canonical-1220031212200222-2321111001232312-1103003232023320-0133131202011220-2321022210022010-3221110302030301-2130012332220112-2321122303001332) |
| `palo_alto_fw_service.auto_setup.admin_password.blindfold_secret_info.location` | [palo_alto_fw_service.auto_setup.admin_password.blindfold_secret_info.location](data-sources--nfv_service--reference--group-003.md#canonical-2210220122110110-0123032113200212-0132131003312020-2333333022032220-3112102230001302-2332121302130301-0103113130003203-3311300202210130) |
| `palo_alto_fw_service.auto_setup.admin_password.blindfold_secret_info.store_provider` | [palo_alto_fw_service.auto_setup.admin_password.blindfold_secret_info.store_provider](data-sources--nfv_service--reference--group-003.md#canonical-2111033033022230-2113120103311230-0121100111012211-0001123023210031-0322112221220221-0131312032303112-3233100220310111-1121330320323133) |
| `palo_alto_fw_service.auto_setup.admin_password.clear_secret_info` | [palo_alto_fw_service.auto_setup.admin_password.clear_secret_info](data-sources--nfv_service--reference--group-003.md#canonical-3131022021011233-1103132133133231-0233232102303312-3213223223010013-1102133330023213-3221132332133020-0323122330222331-2212031133131011) |
| `palo_alto_fw_service.auto_setup.admin_password.clear_secret_info.provider_ref` | [palo_alto_fw_service.auto_setup.admin_password.clear_secret_info.provider_ref](data-sources--nfv_service--reference--group-003.md#canonical-1200201212101203-1100333121212101-3022212012032230-2012131211212231-2311210331001030-2132001001003130-0132301003212112-1232123202321211) |
| `palo_alto_fw_service.auto_setup.admin_password.clear_secret_info.url` | [palo_alto_fw_service.auto_setup.admin_password.clear_secret_info.url](data-sources--nfv_service--reference--group-003.md#canonical-3232013331012300-0232221033113323-0332232322300302-1032232122320113-0210323221022311-3322000002332333-2320011103123010-1022310311112010) |
| `palo_alto_fw_service.auto_setup.admin_username` | [palo_alto_fw_service.auto_setup.admin_username](data-sources--nfv_service--reference--group-003.md#canonical-0321301213031111-2122012313210232-1031100320200110-3023223100301103-0010332230312203-1002023101211303-0132003221312221-1022232022032132) |
| `palo_alto_fw_service.auto_setup.manual_ssh_keys` | [palo_alto_fw_service.auto_setup.manual_ssh_keys](data-sources--nfv_service--reference--group-003.md#canonical-2031010030111102-2301013001331211-0320133301320000-3222112031320120-0013332010023210-0121210123321000-3103002222232231-0000130112133013) |
| `palo_alto_fw_service.auto_setup.manual_ssh_keys.private_key` | [palo_alto_fw_service.auto_setup.manual_ssh_keys.private_key](data-sources--nfv_service--reference--group-003.md#canonical-2130112210210233-0210012001001312-2121021123301120-2031031113112221-3003031101212022-0220302223230133-3223332310233132-1303231021210000) |
| `palo_alto_fw_service.auto_setup.manual_ssh_keys.private_key.blindfold_secret_info` | [palo_alto_fw_service.auto_setup.manual_ssh_keys.private_key.blindfold_secret_info](data-sources--nfv_service--reference--group-004.md#canonical-0131101220320321-0313303200021110-0313202000312312-2213031323010011-3111202320133003-3100022210003103-3110230231122210-1012203221031221) |
| `palo_alto_fw_service.auto_setup.manual_ssh_keys.private_key.blindfold_secret_info.decryption_provider` | [palo_alto_fw_service.auto_setup.manual_ssh_keys.private_key.blindfold_secret_info.decryption_provider](data-sources--nfv_service--reference--group-004.md#canonical-2333321302102031-3012222023322113-2130333010020013-2313320020102010-3221033020102012-1131301303013121-0231331120012110-3331200012112102) |
| `palo_alto_fw_service.auto_setup.manual_ssh_keys.private_key.blindfold_secret_info.location` | [palo_alto_fw_service.auto_setup.manual_ssh_keys.private_key.blindfold_secret_info.location](data-sources--nfv_service--reference--group-004.md#canonical-1313233132303031-2332331330330013-3332020103112003-2033223300122100-1103212221321321-3310303201031110-1201012133231122-2030113200321220) |
| `palo_alto_fw_service.auto_setup.manual_ssh_keys.private_key.blindfold_secret_info.store_provider` | [palo_alto_fw_service.auto_setup.manual_ssh_keys.private_key.blindfold_secret_info.store_provider](data-sources--nfv_service--reference--group-004.md#canonical-2233130003103312-2300113023130212-1023220323031121-1211003201132320-3320112313222013-2311013001103123-1302133302203310-1321301223230233) |
| `palo_alto_fw_service.auto_setup.manual_ssh_keys.private_key.clear_secret_info` | [palo_alto_fw_service.auto_setup.manual_ssh_keys.private_key.clear_secret_info](data-sources--nfv_service--reference--group-004.md#canonical-1303113033003331-1300002003301322-2300033220333211-3012213221302010-3202321231000330-3221223012312031-2033331000212301-2300113321320110) |
| `palo_alto_fw_service.auto_setup.manual_ssh_keys.private_key.clear_secret_info.provider_ref` | [palo_alto_fw_service.auto_setup.manual_ssh_keys.private_key.clear_secret_info.provider_ref](data-sources--nfv_service--reference--group-004.md#canonical-0321131021223332-3113212120310012-2300210232031202-3133001331232303-0212122122200101-0032322122000233-3002132333321023-2103130130032321) |
| `palo_alto_fw_service.auto_setup.manual_ssh_keys.private_key.clear_secret_info.url` | [palo_alto_fw_service.auto_setup.manual_ssh_keys.private_key.clear_secret_info.url](data-sources--nfv_service--reference--group-004.md#canonical-2033003333221223-2023033030102030-2313030011320331-3123112010221232-3010113333131013-3120221031322122-0330030110112112-3120211211312233) |
| `palo_alto_fw_service.auto_setup.manual_ssh_keys.public_key` | [palo_alto_fw_service.auto_setup.manual_ssh_keys.public_key](data-sources--nfv_service--reference--group-003.md#canonical-3313303030201130-3002023313231213-1103122331213121-3300133203130131-1301112320031300-2303221220012213-2003203232010121-0300302103203201) |
| `palo_alto_fw_service.aws_tgw_site` | [palo_alto_fw_service.aws_tgw_site](data-sources--nfv_service--reference--group-004.md#canonical-1002223122113121-3102330210031330-2023303032233111-3110002013231221-0222001203101033-1103102012010210-1033311010133002-2320320332000021) |
| `palo_alto_fw_service.aws_tgw_site.name` | [palo_alto_fw_service.aws_tgw_site.name](data-sources--nfv_service--reference--group-004.md#canonical-2011200112100303-2113221130301032-1020031310220113-0312012303011002-2313023131233031-2203312220133303-2101132020311130-3310033321003230) |
| `palo_alto_fw_service.aws_tgw_site.namespace` | [palo_alto_fw_service.aws_tgw_site.namespace](data-sources--nfv_service--reference--group-004.md#canonical-3233201302011022-1103301222013221-0212130131033132-2230321233132313-3112033131102103-3021233222313212-3322200111113130-0330202012232210) |
| `palo_alto_fw_service.aws_tgw_site.tenant` | [palo_alto_fw_service.aws_tgw_site.tenant](data-sources--nfv_service--reference--group-004.md#canonical-1323010012311002-0321020310020020-2033310013012130-1302121232102033-1201031122210122-2022330033121232-0332330333032000-2030033200301222) |
| `palo_alto_fw_service.disable_panaroma` | [palo_alto_fw_service.disable_panaroma](data-sources--nfv_service--reference--group-004.md#canonical-3221023100000033-1202313210103312-3322133323120231-3000201322320020-3021210032113302-1210030220232012-0013231312030123-1130332311001022) |
| `palo_alto_fw_service.instance_type` | [palo_alto_fw_service.instance_type](data-sources--nfv_service--reference--group-003.md#canonical-3130120003123322-3320211312112110-1130232133011123-2031302212033032-0001210331320200-0110120132020313-2300231321333023-1013120201330312) |
| `palo_alto_fw_service.pan_ami_bundle1` | [palo_alto_fw_service.pan_ami_bundle1](data-sources--nfv_service--reference--group-004.md#canonical-0330002021030120-2230211133322211-0322022233201113-0121123002222100-2103200023222111-0023201000002102-1333333220100033-3012321122230133) |
| `palo_alto_fw_service.pan_ami_bundle2` | [palo_alto_fw_service.pan_ami_bundle2](data-sources--nfv_service--reference--group-004.md#canonical-2223112200023231-1001103221120030-2303002012232110-3011010221231033-3213030210211200-1320232030102222-0001210032111120-0120232002201202) |
| `palo_alto_fw_service.panorama_server` | [palo_alto_fw_service.panorama_server](data-sources--nfv_service--reference--group-004.md#canonical-3312312313220311-0222321000231020-1113120020312122-0023033300023311-3103313121032222-3120303003203103-2133330320302300-1300212000230212) |
| `palo_alto_fw_service.panorama_server.authorization_key` | [palo_alto_fw_service.panorama_server.authorization_key](data-sources--nfv_service--reference--group-004.md#canonical-1300201323132212-0130320020131010-0012100122031033-2021312203100320-2011210223030021-2111020223323311-1100100120202203-1013321001200303) |
| `palo_alto_fw_service.panorama_server.authorization_key.blindfold_secret_info` | [palo_alto_fw_service.panorama_server.authorization_key.blindfold_secret_info](data-sources--nfv_service--reference--group-004.md#canonical-3032113131223123-2112301032200002-3321232112022210-1212230031332322-2023203332302323-3323121303203323-2222303223130133-2323112211012310) |
| `palo_alto_fw_service.panorama_server.authorization_key.blindfold_secret_info.decryption_provider` | [palo_alto_fw_service.panorama_server.authorization_key.blindfold_secret_info.decryption_provider](data-sources--nfv_service--reference--group-004.md#canonical-3211002012201312-2102133313332013-0222032101022011-0103230132020210-2233122331330103-1223320022210030-3032212332212131-0303122321120231) |
| `palo_alto_fw_service.panorama_server.authorization_key.blindfold_secret_info.location` | [palo_alto_fw_service.panorama_server.authorization_key.blindfold_secret_info.location](data-sources--nfv_service--reference--group-004.md#canonical-1312020013310023-2033013132221001-1202201120000031-1132131100330111-0133300321000311-3310203231220333-3320031322121232-1000013323303012) |
| `palo_alto_fw_service.panorama_server.authorization_key.blindfold_secret_info.store_provider` | [palo_alto_fw_service.panorama_server.authorization_key.blindfold_secret_info.store_provider](data-sources--nfv_service--reference--group-004.md#canonical-3330132002022100-0232331211122021-0030110101011202-2102031202010103-0023013221122303-3213322121110030-1000212210103221-1002203210202002) |
| `palo_alto_fw_service.panorama_server.authorization_key.clear_secret_info` | [palo_alto_fw_service.panorama_server.authorization_key.clear_secret_info](data-sources--nfv_service--reference--group-004.md#canonical-3331302012321122-2021321300322212-0320221013133233-1112323231110223-2113032232003300-3231031320202230-0313330031130123-2111021313003230) |
| `palo_alto_fw_service.panorama_server.authorization_key.clear_secret_info.provider_ref` | [palo_alto_fw_service.panorama_server.authorization_key.clear_secret_info.provider_ref](data-sources--nfv_service--reference--group-004.md#canonical-2320201313233220-3130320020233012-2133121123222132-3132032000123233-0002023102103002-3010120110123013-3030203112033332-2202001132123231) |
| `palo_alto_fw_service.panorama_server.authorization_key.clear_secret_info.url` | [palo_alto_fw_service.panorama_server.authorization_key.clear_secret_info.url](data-sources--nfv_service--reference--group-004.md#canonical-3033102111332302-0322302000133202-3230232331101200-1011213021300030-2130231330010010-3201023132001221-2033313002011120-2003201033023213) |
| `palo_alto_fw_service.panorama_server.device_group_name` | [palo_alto_fw_service.panorama_server.device_group_name](data-sources--nfv_service--reference--group-004.md#canonical-3200133030102030-0211013233110300-1122121200100202-0222210221223101-1232310232212320-0023200231323303-0322332012323231-1313312001212231) |
| `palo_alto_fw_service.panorama_server.server` | [palo_alto_fw_service.panorama_server.server](data-sources--nfv_service--reference--group-004.md#canonical-3223212320221030-1333020123302023-0101101001003301-3211021212320331-0033113223322211-0310130210013322-3302220101220310-0031120212023231) |
| `palo_alto_fw_service.panorama_server.template_stack_name` | [palo_alto_fw_service.panorama_server.template_stack_name](data-sources--nfv_service--reference--group-004.md#canonical-3323031213320330-1120211030313102-0001202233313203-0202020310312220-2030003003132032-3033200212320223-0220332201231320-0101122231113330) |
| `palo_alto_fw_service.service_nodes` | [palo_alto_fw_service.service_nodes](data-sources--nfv_service--reference--group-004.md#canonical-3111023333233030-2330222103003112-0333322321203330-2233021223312021-1311003210112212-0201320111330121-3113223003301321-2332313012211021) |
| `palo_alto_fw_service.service_nodes.nodes` | [palo_alto_fw_service.service_nodes.nodes](data-sources--nfv_service--reference--group-004.md#canonical-2021220321123132-3121211221022002-2333302232122320-1322013320133220-0220222131220211-0123121131133031-2232100210132222-0123101033023021) |
| `palo_alto_fw_service.service_nodes.nodes.aws_az_name` | [palo_alto_fw_service.service_nodes.nodes.aws_az_name](data-sources--nfv_service--reference--group-004.md#canonical-1212203330302023-0131200333033302-0213212023322000-0132220302323231-0203230222323020-3311231002213323-0002311031211103-3033221232211313) |
| `palo_alto_fw_service.service_nodes.nodes.mgmt_subnet` | [palo_alto_fw_service.service_nodes.nodes.mgmt_subnet](data-sources--nfv_service--reference--group-004.md#canonical-2311321201203132-1310331112122323-2301020112322233-3333013021230213-2322031010121303-1331330201210223-2223210123211211-0123030122001331) |
| `palo_alto_fw_service.service_nodes.nodes.mgmt_subnet.existing_subnet_id` | [palo_alto_fw_service.service_nodes.nodes.mgmt_subnet.existing_subnet_id](data-sources--nfv_service--reference--group-004.md#canonical-3232211322303120-0100222302302010-1030331011330122-2011131013200322-3313310112201001-1021103012230001-3012131113100320-3023012232113201) |
| `palo_alto_fw_service.service_nodes.nodes.mgmt_subnet.subnet_param` | [palo_alto_fw_service.service_nodes.nodes.mgmt_subnet.subnet_param](data-sources--nfv_service--reference--group-004.md#canonical-2001031322313022-2203210011032310-1001000230012200-2121011033212133-2202312233130121-1331310233120212-3323110113333130-1300102320322333) |
| `palo_alto_fw_service.service_nodes.nodes.mgmt_subnet.subnet_param.ipv4` | [palo_alto_fw_service.service_nodes.nodes.mgmt_subnet.subnet_param.ipv4](data-sources--nfv_service--reference--group-004.md#canonical-1300000212301130-1002333033031311-1222113002221212-0211220010323011-3321313331333222-3201122133231203-2223201333103202-2000110311201121) |
| `palo_alto_fw_service.service_nodes.nodes.node_name` | [palo_alto_fw_service.service_nodes.nodes.node_name](data-sources--nfv_service--reference--group-004.md#canonical-0113120213023203-2000323212100300-2010013111002031-3010330120110202-3103302100122102-0222121002102230-2002211002221013-3211202030012003) |
| `palo_alto_fw_service.service_nodes.nodes.reserved_mgmt_subnet` | [palo_alto_fw_service.service_nodes.nodes.reserved_mgmt_subnet](data-sources--nfv_service--reference--group-004.md#canonical-3222001131002220-2123300131030320-3233023211132200-1031230303203110-2102302012331332-1010303213302032-0113112311121212-3213000222020120) |
| `palo_alto_fw_service.ssh_key` | [palo_alto_fw_service.ssh_key](data-sources--nfv_service--reference--group-003.md#canonical-1213322231133131-0331010312322200-0210312021033021-0231030210232123-3233001223133231-1103321221211111-0031203121211000-3233120123132012) |
| `palo_alto_fw_service.tags` | [palo_alto_fw_service.tags](data-sources--nfv_service--reference--group-003.md#canonical-2231312321311130-0301220130301031-3110131112331020-0310122220330231-0003132130320123-0232021132311022-1003101231203231-2133331332011100) |
| `palo_alto_fw_service.version` | [palo_alto_fw_service.version](data-sources--nfv_service--reference--group-003.md#canonical-2231020202202023-2220020331032032-1331002231230321-3222023221302020-3130033201111120-2001012300233302-2310113313122230-1201000110101032) |

<a id="canonical-3213103303301133-3333303210230310-3323002302303301-2320020003110332-0230322231122232-2012333331221302-3333330213301323-2233113002011321"></a>

## Next pages — Property reference / 011321300323 / 11

- [disable_https_management](data-sources--nfv_service--reference--group-001.md#canonical-3233032123222311-3001001210121231-0312011132321313-1232230220130301-3231301131202311-0233202113011321-1032221330102003-2312002332132102)
- [disable_ssh_access](data-sources--nfv_service--reference--group-001.md#canonical-1221033000332001-2223320310121121-0001202121031020-2311110013220222-3223210220123033-3303322313032323-2103300120311112-3302300312112212)
- [enabled_ssh_access](data-sources--nfv_service--reference--group-001.md#canonical-2323120212210131-3332303320001030-2120313330111100-0101001020220303-0323332221022001-1220132333131230-1200303110333213-2001222021201023)
- [f5_big_ip_aws_service](data-sources--nfv_service--reference--group-001.md#canonical-3121203202232102-3301112102333133-2103303100231002-1321223111220201-3103311113020000-0303332002322201-3021331113310133-2201221310222221)
- [https_management](data-sources--nfv_service--reference--group-002.md#canonical-0322223112333211-3323030031122111-0131002222131221-2002123332012320-0013112013120020-2030102011102030-0020230010011012-0231301033231112)
- [palo_alto_fw_service](data-sources--nfv_service--reference--group-003.md#canonical-0130222001222131-3123210310203012-0200310203211313-0131331300023100-2310131123330002-1011033223002202-2221031230120202-1331030023120030)
- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-2212320103310132-2221300222221032-1233031123200120-3022110322311231-3301003023302233-3133232120201022-0322021211220213-1332311000003300)

<a id="canonical-3233032123222311-3001001210121231-0312011132321313-1232230220130301-3231301131202311-0233202113011321-1032221330102003-2312002332132102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3332330212303222-2013123031013231-2110312212020003-0113230113130222-0113002210223320-3112213310120111-2103112000101211-0123213330131031"></a>

## disable_https_management — disable_https_management / 020230313333 / 2

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-2212320103310132-2221300222221032-1233031123200120-3022110322311231-3301003023302233-3133232120201022-0322021211220213-1332311000003300)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-2313010333323131-3010030231311121-2301323120300012-0122321020201331-3020223211220022-0100121100023011-2323321311300323-0122100132003020)
- disable_https_management

<a id="canonical-1123230132031022-2230003222203310-3220003002031030-0130133021301033-1110001030131201-2010020310131122-3212333113023313-1013010302300230"></a>

Type: `["object", {}]`. Computed.

\[OneOf: disable\_https\_management, https\_management; Default: disable\_https\_management\]
Configuration parameter for disable https management.

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

OneOf alternatives in this subsection:

- [disable_https_management](data-sources--nfv_service--reference--group-001.md#canonical-1123230132031022-2230003222203310-3220003002031030-0130133021301033-1110001030131201-2010020310131122-3212333113023313-1013010302300230)
- [https_management](data-sources--nfv_service--reference--group-002.md#canonical-3331102201020133-2333120011110303-1201032002311221-0030302301022231-0303203311202320-0211011020011033-2330232112032332-3132123012211010)

Select alternatives according to the provider validators above.

<a id="canonical-3220310110111220-2002003211303323-0030323222333112-1313333311000210-2032101101002202-0300110300030312-0023233032011201-3332330112013232"></a>

## Direct properties — disable_https_management / 020230313333 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1021103131113110-3010202311003111-0102232033103013-0102110320222012-2102120133132232-1222323321002332-0330021233021220-2123312301010303"></a>

## Next pages — disable_https_management / 020230313333 / 4

- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-2313010333323131-3010030231311121-2301323120300012-0122321020201331-3020223211220022-0100121100023011-2323321311300323-0122100132003020)
- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-2212320103310132-2221300222221032-1233031123200120-3022110322311231-3301003023302233-3133232120201022-0322021211220213-1332311000003300)

<a id="canonical-1221033000332001-2223320310121121-0001202121031020-2311110013220222-3223210220123033-3303322313032323-2103300120311112-3302300312112212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3002110000113332-2311031031110030-2120310101100100-3322230302010103-0213231102313212-3100213121312122-0211132103320132-0310000101013303"></a>

## disable_ssh_access — disable_ssh_access / 003120300010 / 2

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-2212320103310132-2221300222221032-1233031123200120-3022110322311231-3301003023302233-3133232120201022-0322021211220213-1332311000003300)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-2313010333323131-3010030231311121-2301323120300012-0122321020201331-3020223211220022-0100121100023011-2323321311300323-0122100132003020)
- disable_ssh_access

<a id="canonical-3322300221020310-1022320213222010-1003030332021023-1003202032123101-2331012113322001-3022212232102110-0102011121223220-2033000110230323"></a>

Type: `["object", {}]`. Computed.

\[OneOf: disable\_ssh\_access, enabled\_ssh\_access; Default: disable\_ssh\_access\] Configuration
parameter for disable ssh access.

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

OneOf alternatives in this subsection:

- [disable_ssh_access](data-sources--nfv_service--reference--group-001.md#canonical-3322300221020310-1022320213222010-1003030332021023-1003202032123101-2331012113322001-3022212232102110-0102011121223220-2033000110230323)
- [enabled_ssh_access](data-sources--nfv_service--reference--group-001.md#canonical-3231302020013323-0132223211210223-2112032102231030-2123122230000331-1031332220111111-2232322012102201-0222213113233310-1330222012001120)

Select alternatives according to the provider validators above.

<a id="canonical-3300220211133110-2220331232130133-0203320200313133-1101101030021320-1012333023331120-3303023230223022-0113001201020310-0032021301203132"></a>

## Direct properties — disable_ssh_access / 003120300010 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1311103022120000-3132333211331031-0320301231310232-1020211200121321-2222102012123230-2320211222022112-1130022302132133-1221111323112311"></a>

## Next pages — disable_ssh_access / 003120300010 / 4

- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-2313010333323131-3010030231311121-2301323120300012-0122321020201331-3020223211220022-0100121100023011-2323321311300323-0122100132003020)
- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-2212320103310132-2221300222221032-1233031123200120-3022110322311231-3301003023302233-3133232120201022-0322021211220213-1332311000003300)

<a id="canonical-2323120212210131-3332303320001030-2120313330111100-0101001020220303-0323332221022001-1220132333131230-1200303110333213-2001222021201023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0302331201330132-3322013210133122-0222211132233302-1022321330230203-3303033032113310-2311001321031321-1332133330332022-3122120212211021"></a>

## enabled_ssh_access — enabled_ssh_access / 311130011031 / 2

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-2212320103310132-2221300222221032-1233031123200120-3022110322311231-3301003023302233-3133232120201022-0322021211220213-1332311000003300)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-2313010333323131-3010030231311121-2301323120300012-0122321020201331-3020223211220022-0100121100023011-2323321311300323-0122100132003020)
- enabled_ssh_access

<a id="canonical-3231302020013323-0132223211210223-2112032102231030-2123122230000331-1031332220111111-2232322012102201-0222213113233310-1330222012001120"></a>

Type: `"single"`. Computed.

Configuration parameter for enabled ssh access.

Upstream description:

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

<a id="canonical-1202131100111200-2000333033201213-1313110301333210-1311233032210202-2221122032110302-1313220132322301-2110333222023200-0003311302330030"></a>

## Direct properties — enabled_ssh_access / 311130011031 / 3

- [advertise_on_sli](data-sources--nfv_service--reference--group-001.md#canonical-0032103101313111-2331311201113010-1311130022121321-2233231132233312-2311112221312303-1131013203031032-3312320312131033-0332213100033303): complete subsection reference.

- [advertise_on_slo](data-sources--nfv_service--reference--group-001.md#canonical-1301300322221023-1333023120330231-0010020202302131-2212310300333031-3301001122210220-1302203021023111-0231020330113123-0120010100312120): complete subsection reference.

- [advertise_on_slo_sli](data-sources--nfv_service--reference--group-001.md#canonical-3101031022003013-2320313333011120-1133330122231111-0122320330123233-0121311010310331-2013231103120212-1212233312310322-1033301103122102): complete subsection reference.

<a id="canonical-2213222112030101-3333013110223330-2222121133111131-0111221330032032-2332201233311323-0111133303131220-3133111320332322-2331112320010201"></a>

<a id="canonical-0231322030302202-2300111103322123-1033131013223301-2003330302122011-2303310013210002-1110313202302020-0210121130321232-1312331102231333"></a>

## domain_suffix property — enabled_ssh_access / 311130011031 / 4

Type: `"string"`. Computed.

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
    "ves.io.schema.rules.string.hostname": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.hostname": "true"
  }
}
```

- [node_ssh_ports](data-sources--nfv_service--reference--group-001.md#canonical-0022331032213222-2321032331132022-3101102101312200-2333131201211231-1003000002013302-0032130202121131-1232322121311231-1032230112233023): complete subsection reference.

<a id="canonical-0320000220221313-2302322301110312-2301022332111211-0200022313300133-1112012321223101-2132030020231131-3122213230223033-0003230200100001"></a>

## Next pages — enabled_ssh_access / 311130011031 / 5

- [enabled_ssh_access.advertise_on_sli](data-sources--nfv_service--reference--group-001.md#canonical-0032103101313111-2331311201113010-1311130022121321-2233231132233312-2311112221312303-1131013203031032-3312320312131033-0332213100033303)
- [enabled_ssh_access.advertise_on_slo](data-sources--nfv_service--reference--group-001.md#canonical-1301300322221023-1333023120330231-0010020202302131-2212310300333031-3301001122210220-1302203021023111-0231020330113123-0120010100312120)
- [enabled_ssh_access.advertise_on_slo_sli](data-sources--nfv_service--reference--group-001.md#canonical-3101031022003013-2320313333011120-1133330122231111-0122320330123233-0121311010310331-2013231103120212-1212233312310322-1033301103122102)
- [enabled_ssh_access.node_ssh_ports](data-sources--nfv_service--reference--group-001.md#canonical-0022331032213222-2321032331132022-3101102101312200-2333131201211231-1003000002013302-0032130202121131-1232322121311231-1032230112233023)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-2313010333323131-3010030231311121-2301323120300012-0122321020201331-3020223211220022-0100121100023011-2323321311300323-0122100132003020)
- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-2212320103310132-2221300222221032-1233031123200120-3022110322311231-3301003023302233-3133232120201022-0322021211220213-1332311000003300)

<a id="canonical-0032103101313111-2331311201113010-1311130022121321-2233231132233312-2311112221312303-1131013203031032-3312320312131033-0332213100033303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0230000233321102-0323210301020230-1102103333013333-2303201210202123-3023013101332201-0130222301233121-0123033023302001-0010221122232202"></a>

## enabled_ssh_access.advertise_on_sli — advertise_on_sli / 003300012321 / 2

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-2212320103310132-2221300222221032-1233031123200120-3022110322311231-3301003023302233-3133232120201022-0322021211220213-1332311000003300)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-2313010333323131-3010030231311121-2301323120300012-0122321020201331-3020223211220022-0100121100023011-2323321311300323-0122100132003020)
- [enabled_ssh_access](data-sources--nfv_service--reference--group-001.md#canonical-2323120212210131-3332303320001030-2120313330111100-0101001020220303-0323332221022001-1220132333131230-1200303110333213-2001222021201023)
- enabled_ssh_access.advertise_on_sli

<a id="canonical-3021212023001322-1213203013021332-3102001100030201-0103222211030022-3202220221332023-3131300030321000-3122101301112223-2001033303230210"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for advertise on sli.

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

<a id="canonical-3232201221001102-2213101022002201-0113211131101232-2332030103020322-1202021233213020-0301231202102330-2123230210020011-3233002300323131"></a>

## Direct properties — advertise_on_sli / 003300012321 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0322012221333133-3223302131033331-3312113220331103-0310122323330023-0220003123021200-2103200032010313-0021232211233023-0101330303021232"></a>

## Next pages — advertise_on_sli / 003300012321 / 4

- [enabled_ssh_access](data-sources--nfv_service--reference--group-001.md#canonical-2323120212210131-3332303320001030-2120313330111100-0101001020220303-0323332221022001-1220132333131230-1200303110333213-2001222021201023)
- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-2212320103310132-2221300222221032-1233031123200120-3022110322311231-3301003023302233-3133232120201022-0322021211220213-1332311000003300)

<a id="canonical-1301300322221023-1333023120330231-0010020202302131-2212310300333031-3301001122210220-1302203021023111-0231020330113123-0120010100312120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0010202002313210-2101303323132133-3031022013332312-1221211032123030-2223310020003010-1011020001103001-0301320132103132-3231321233201001"></a>

## enabled_ssh_access.advertise_on_slo — advertise_on_slo / 203103323210 / 2

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-2212320103310132-2221300222221032-1233031123200120-3022110322311231-3301003023302233-3133232120201022-0322021211220213-1332311000003300)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-2313010333323131-3010030231311121-2301323120300012-0122321020201331-3020223211220022-0100121100023011-2323321311300323-0122100132003020)
- [enabled_ssh_access](data-sources--nfv_service--reference--group-001.md#canonical-2323120212210131-3332303320001030-2120313330111100-0101001020220303-0323332221022001-1220132333131230-1200303110333213-2001222021201023)
- enabled_ssh_access.advertise_on_slo

<a id="canonical-0302030102011032-3020122123312323-0333020011112021-0223012110013000-2212100200102231-2003200120101003-2012021310120131-3101113030133130"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for advertise on slo.

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

<a id="canonical-2132213111100312-1023022200221023-2221310003330301-3231303331331001-2020123213320313-1103113332230331-0202330313222310-2202120102310121"></a>

## Direct properties — advertise_on_slo / 203103323210 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1003101013103203-3102113213212310-0020000322300133-2011331121300103-1103300011300122-3230210003330212-0000233200113133-1130320023210120"></a>

## Next pages — advertise_on_slo / 203103323210 / 4

- [enabled_ssh_access](data-sources--nfv_service--reference--group-001.md#canonical-2323120212210131-3332303320001030-2120313330111100-0101001020220303-0323332221022001-1220132333131230-1200303110333213-2001222021201023)
- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-2212320103310132-2221300222221032-1233031123200120-3022110322311231-3301003023302233-3133232120201022-0322021211220213-1332311000003300)

<a id="canonical-3101031022003013-2320313333011120-1133330122231111-0122320330123233-0121311010310331-2013231103120212-1212233312310322-1033301103122102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2323031231000011-0311220210332222-3331202212321010-2200033200323231-3000300313032231-3320032223002011-2012211130121230-1212113321122130"></a>

## enabled_ssh_access.advertise_on_slo_sli — advertise_on_slo_sli / 000133311000 / 2

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-2212320103310132-2221300222221032-1233031123200120-3022110322311231-3301003023302233-3133232120201022-0322021211220213-1332311000003300)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-2313010333323131-3010030231311121-2301323120300012-0122321020201331-3020223211220022-0100121100023011-2323321311300323-0122100132003020)
- [enabled_ssh_access](data-sources--nfv_service--reference--group-001.md#canonical-2323120212210131-3332303320001030-2120313330111100-0101001020220303-0323332221022001-1220132333131230-1200303110333213-2001222021201023)
- enabled_ssh_access.advertise_on_slo_sli

<a id="canonical-3123230300313323-3022122023222303-3221203330310313-0120003300231313-0101033323132021-2202333103331202-3113112222013323-1211002331322133"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for advertise on slo sli.

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

<a id="canonical-3131031212200111-1230333221000133-1113113233322322-2223102323213310-0110331002100130-1013112220220332-3323221120202012-1213002103020122"></a>

## Direct properties — advertise_on_slo_sli / 000133311000 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1212001232003003-1020300213011032-1012002223002301-0233110232333002-0222312130100203-0031331333231300-1203010001321103-3001113103201102"></a>

## Next pages — advertise_on_slo_sli / 000133311000 / 4

- [enabled_ssh_access](data-sources--nfv_service--reference--group-001.md#canonical-2323120212210131-3332303320001030-2120313330111100-0101001020220303-0323332221022001-1220132333131230-1200303110333213-2001222021201023)
- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-2212320103310132-2221300222221032-1233031123200120-3022110322311231-3301003023302233-3133232120201022-0322021211220213-1332311000003300)

<a id="canonical-0022331032213222-2321032331132022-3101102101312200-2333131201211231-1003000002013302-0032130202121131-1232322121311231-1032230112233023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2133210010000101-2201223003303011-0320223333310121-1103221122332331-0130000010120030-0301201300201320-0002221223130203-3022322210300020"></a>

## enabled_ssh_access.node_ssh_ports — node_ssh_ports / 101221223221 / 2

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-2212320103310132-2221300222221032-1233031123200120-3022110322311231-3301003023302233-3133232120201022-0322021211220213-1332311000003300)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-2313010333323131-3010030231311121-2301323120300012-0122321020201331-3020223211220022-0100121100023011-2323321311300323-0122100132003020)
- [enabled_ssh_access](data-sources--nfv_service--reference--group-001.md#canonical-2323120212210131-3332303320001030-2120313330111100-0101001020220303-0323332221022001-1220132333131230-1200303110333213-2001222021201023)
- enabled_ssh_access.node_ssh_ports

<a id="canonical-3302313201101220-0011123021313023-3220210111020130-1000331220221102-3312202323230323-1211223301110032-2020012303311032-1030332102230220"></a>

Type: `"list"`. Computed.

Management Node SSH Port. Enter TCP port and node name per node.

Upstream description:

Enter TCP port and node name per node.

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
    "ves.io.schema.rules.repeated.max_items": "2"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "2"
  }
}
```

<a id="canonical-1023102213313331-1231300213330033-2211122112113120-3110211033112321-3020203302333103-0202201311031310-2231012303110121-3311023101323101"></a>

## Direct properties — node_ssh_ports / 101221223221 / 3

<a id="canonical-1211232230300010-2003330312233223-0203213131011110-0332132213221020-0323010210000323-2022333130123311-3103222332022023-1111101321331303"></a>

<a id="canonical-0130313302030311-3300031332301000-1302203110213210-2333102030111300-0120122320131312-1120221212323132-3212101223101010-1311202101221030"></a>

## node_name property — node_ssh_ports / 101221223221 / 4

Type: `"string"`. Computed.

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

<a id="canonical-2212222320011231-2301211123023133-1111210130033300-1230021321211212-3321032202212031-1103312333333310-2332030332333233-2003010303110202"></a>

<a id="canonical-1222102101301200-1220110031311130-3301123010231332-2222312222110200-0211303332131210-0222131203001130-1231133233123213-0011313230203212"></a>

## ssh_port property — node_ssh_ports / 101221223221 / 5

Type: `"number"`. Computed.

SSH Port. Enter TCP port per node.

Upstream description:

Enter TCP port per node.

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

<a id="canonical-0030301211313000-3130202332002203-3110021011212222-3310100221232021-2211021322132310-3033133322100210-3202331012020021-3303233002131130"></a>

## Next pages — node_ssh_ports / 101221223221 / 6

- [enabled_ssh_access](data-sources--nfv_service--reference--group-001.md#canonical-2323120212210131-3332303320001030-2120313330111100-0101001020220303-0323332221022001-1220132333131230-1200303110333213-2001222021201023)
- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-2212320103310132-2221300222221032-1233031123200120-3022110322311231-3301003023302233-3133232120201022-0322021211220213-1332311000003300)

<a id="canonical-3121203202232102-3301112102333133-2103303100231002-1321223111220201-3103311113020000-0303332002322201-3021331113310133-2201221310222221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3113312333313132-3333211120302323-1100112212001011-0230130003311322-1130202100302201-0021333111223220-0010120322011130-3003321223130322"></a>

## f5_big_ip_aws_service — f5_big_ip_aws_service / 221030201132 / 2

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-2212320103310132-2221300222221032-1233031123200120-3022110322311231-3301003023302233-3133232120201022-0322021211220213-1332311000003300)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-2313010333323131-3010030231311121-2301323120300012-0122321020201331-3020223211220022-0100121100023011-2323321311300323-0122100132003020)
- f5_big_ip_aws_service

<a id="canonical-1000212232223202-3102233333221300-1111113133333303-1002121012300313-0201010302013330-2021303201311211-0323112122022121-0330010210332033"></a>

Type: `"single"`. Computed.

\[OneOf: f5\_big\_ip\_aws\_service, palo\_alto\_fw\_service\] Virtual BIG-IP AWS. Virtual BIG-IP
specification for AWS.

Upstream description:

Virtual BIG-IP specification for AWS.

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

- [f5_big_ip_aws_service](data-sources--nfv_service--reference--group-001.md#canonical-1000212232223202-3102233333221300-1111113133333303-1002121012300313-0201010302013330-2021303201311211-0323112122022121-0330010210332033)
- [palo_alto_fw_service](data-sources--nfv_service--reference--group-003.md#canonical-2212323110330013-3101222023103001-2303100101011111-2301211332231032-2320133322031322-3100212302103322-2302012130013323-3313030213120310)

Select alternatives according to the provider validators above.

<a id="canonical-3223313030013313-0300012101031301-0101111211322213-3332312311301120-2020023230012122-2030122221331322-2013013131313111-0011101331102021"></a>

## Direct properties — f5_big_ip_aws_service / 221030201132 / 3

- [admin_password](data-sources--nfv_service--reference--group-001.md#canonical-1321011123103003-0223330110312101-0111101021303212-3301023331031210-0021230321333022-0133210003220220-3321222301023200-0111121311331331): complete subsection reference.

<a id="canonical-0211302020113331-0302232110022122-1020122001132313-2022332110130120-2312013200102330-3120323133102300-3100221000210303-3100120312311121"></a>

<a id="canonical-0332113033021230-0033302303330223-3203112211211320-1011123130330230-2211101001001323-3000222200233221-1033321132232030-0020231301102201"></a>

## admin_username property — f5_big_ip_aws_service / 221030201132 / 4

Type: `"string"`. Computed.

Admin Username. Admin Username for BIG-IP.

Upstream description:

Admin Username for BIG-IP.

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

- [aws_tgw_site_params](data-sources--nfv_service--reference--group-001.md#canonical-1130322021203231-2300212333020310-3210110221221221-1003120300332113-1323310310320132-3100020313313322-3031020210022022-2110121011300330): complete subsection reference.

- [endpoint_service](data-sources--nfv_service--reference--group-001.md#canonical-1100221120003321-3013113033000220-0121013203232101-0223100022123001-0102122201021202-3023132102201031-0030303212302022-0220303202322013): complete subsection reference.

- [market_place_image](data-sources--nfv_service--reference--group-001.md#canonical-2210031222201112-3331110222012210-2132033313310012-3320121221230020-0131021130301020-0223033002000032-3021303122131320-2001020121212112): complete subsection reference.

- [nodes](data-sources--nfv_service--reference--group-002.md#canonical-3131321100033332-2120130100102110-2101000032232132-2120201123123102-3322330003003013-3102100333000113-3113220113131113-1101122011233323): complete subsection reference.

<a id="canonical-3223311100111002-2023301003333220-0220310100330022-2131030032203122-1022123103030002-3123111220122031-3120303211113233-2232201220312321"></a>

<a id="canonical-2101130203322310-3132031300132032-3213313000131030-1221110333300010-2301321222230202-1211120101300200-2302200031330302-0010003310010100"></a>

## ssh_key property — f5_big_ip_aws_service / 221030201132 / 5

Type: `"string"`. Computed.

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

<a id="canonical-1132203121331230-1220113023121303-3300332302230221-2013221001023101-3230133132130132-1321301112033301-1003113232132101-0000202033121102"></a>

<a id="canonical-3203110213003321-0021103331212332-2113313003113331-0103111301013011-0201213310312312-2110220330131032-1220301230232000-3011302233201321"></a>

## tags property — f5_big_ip_aws_service / 221030201132 / 6

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

<a id="canonical-2103333232100020-1203030102112302-0321122113022300-1223032222230133-1023301111300213-1322223030302022-1012231032033311-1330321022222001"></a>

## Next pages — f5_big_ip_aws_service / 221030201132 / 7

- [f5_big_ip_aws_service.admin_password](data-sources--nfv_service--reference--group-001.md#canonical-1321011123103003-0223330110312101-0111101021303212-3301023331031210-0021230321333022-0133210003220220-3321222301023200-0111121311331331)
- [f5_big_ip_aws_service.aws_tgw_site_params](data-sources--nfv_service--reference--group-001.md#canonical-1130322021203231-2300212333020310-3210110221221221-1003120300332113-1323310310320132-3100020313313322-3031020210022022-2110121011300330)
- [f5_big_ip_aws_service.endpoint_service](data-sources--nfv_service--reference--group-001.md#canonical-1100221120003321-3013113033000220-0121013203232101-0223100022123001-0102122201021202-3023132102201031-0030303212302022-0220303202322013)
- [f5_big_ip_aws_service.market_place_image](data-sources--nfv_service--reference--group-001.md#canonical-2210031222201112-3331110222012210-2132033313310012-3320121221230020-0131021130301020-0223033002000032-3021303122131320-2001020121212112)
- [f5_big_ip_aws_service.nodes](data-sources--nfv_service--reference--group-002.md#canonical-3131321100033332-2120130100102110-2101000032232132-2120201123123102-3322330003003013-3102100333000113-3113220113131113-1101122011233323)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-2313010333323131-3010030231311121-2301323120300012-0122321020201331-3020223211220022-0100121100023011-2323321311300323-0122100132003020)
- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-2212320103310132-2221300222221032-1233031123200120-3022110322311231-3301003023302233-3133232120201022-0322021211220213-1332311000003300)

<a id="canonical-1321011123103003-0223330110312101-0111101021303212-3301023331031210-0021230321333022-0133210003220220-3321222301023200-0111121311331331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1213101312321103-0322030201013232-1233312112332312-3223012333311320-1122323121022220-1121120022133123-1202203322330000-1300330033111321"></a>

## f5_big_ip_aws_service.admin_password — admin_password / 023101122113 / 2

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-2212320103310132-2221300222221032-1233031123200120-3022110322311231-3301003023302233-3133232120201022-0322021211220213-1332311000003300)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-2313010333323131-3010030231311121-2301323120300012-0122321020201331-3020223211220022-0100121100023011-2323321311300323-0122100132003020)
- [f5_big_ip_aws_service](data-sources--nfv_service--reference--group-001.md#canonical-3121203202232102-3301112102333133-2103303100231002-1321223111220201-3103311113020000-0303332002322201-3021331113310133-2201221310222221)
- f5_big_ip_aws_service.admin_password

<a id="canonical-0203130131012333-1200022021221312-3313033221032333-3331123311232131-1331130312030123-3023012121113133-0302213012201312-3323111130003003"></a>

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

<a id="canonical-1211011113113130-2200221320033232-2120102320231303-0130122013100222-1302100103011003-1222202231321232-0001020012302322-0322332313323212"></a>

## Direct properties — admin_password / 023101122113 / 3

- [blindfold_secret_info](data-sources--nfv_service--reference--group-001.md#canonical-2211303232122322-1210002130131111-2301020102011120-0202001113330020-2013333333113200-1103021333313302-2030321001212101-3100220231210122): complete subsection reference.

- [clear_secret_info](data-sources--nfv_service--reference--group-001.md#canonical-0011321212002310-3013230101301223-2330001032101120-0321312211320102-3120320103221010-3103322221303132-1213230200011001-2211121120130201): complete subsection reference.

<a id="canonical-3131013113233201-2031312021021101-2321203330011323-1213211232131030-1201213320112231-1102100103301202-1213103331233130-1323313312021003"></a>

## Next pages — admin_password / 023101122113 / 4

- [f5_big_ip_aws_service.admin_password.blindfold_secret_info](data-sources--nfv_service--reference--group-001.md#canonical-2211303232122322-1210002130131111-2301020102011120-0202001113330020-2013333333113200-1103021333313302-2030321001212101-3100220231210122)
- [f5_big_ip_aws_service.admin_password.clear_secret_info](data-sources--nfv_service--reference--group-001.md#canonical-0011321212002310-3013230101301223-2330001032101120-0321312211320102-3120320103221010-3103322221303132-1213230200011001-2211121120130201)
- [f5_big_ip_aws_service](data-sources--nfv_service--reference--group-001.md#canonical-3121203202232102-3301112102333133-2103303100231002-1321223111220201-3103311113020000-0303332002322201-3021331113310133-2201221310222221)
- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-2212320103310132-2221300222221032-1233031123200120-3022110322311231-3301003023302233-3133232120201022-0322021211220213-1332311000003300)

<a id="canonical-2211303232122322-1210002130131111-2301020102011120-0202001113330020-2013333333113200-1103021333313302-2030321001212101-3100220231210122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2213120303022021-0031021200311310-0131101001102113-0230210230033103-1131330210303302-1132111211222302-2212032101110130-0111312313010022"></a>

## f5_big_ip_aws_service.admin_password.blindfold_secret_info — blindfold_secret_info / 012021013011 / 2

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-2212320103310132-2221300222221032-1233031123200120-3022110322311231-3301003023302233-3133232120201022-0322021211220213-1332311000003300)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-2313010333323131-3010030231311121-2301323120300012-0122321020201331-3020223211220022-0100121100023011-2323321311300323-0122100132003020)
- [f5_big_ip_aws_service](data-sources--nfv_service--reference--group-001.md#canonical-3121203202232102-3301112102333133-2103303100231002-1321223111220201-3103311113020000-0303332002322201-3021331113310133-2201221310222221)
- [f5_big_ip_aws_service.admin_password](data-sources--nfv_service--reference--group-001.md#canonical-1321011123103003-0223330110312101-0111101021303212-3301023331031210-0021230321333022-0133210003220220-3321222301023200-0111121311331331)
- f5_big_ip_aws_service.admin_password.blindfold_secret_info

<a id="canonical-1220030023200320-2012032331323102-2111333022321113-2110023100032203-2003221223102111-2211121301022112-2232113013003332-0000300012203320"></a>

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

<a id="canonical-3220231003122000-3102201330222211-1012131022320111-2233021133303123-2003222032211131-0022102121222020-0210011223220011-3201012211020111"></a>

## Direct properties — blindfold_secret_info / 012021013011 / 3

<a id="canonical-3100233013221321-2222001321002200-0223111203310311-1101031222101103-1010030112021230-1313103013300031-0122203030221011-2230023211000230"></a>

<a id="canonical-2200313103320202-0120212221110020-0010300103231331-2211131003220203-2222111222303130-3200031323003303-3103323103011203-1002011121103010"></a>

## decryption_provider property — blindfold_secret_info / 012021013011 / 4

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

<a id="canonical-0130113201112121-2103101102302030-2230121310031123-1030321002010121-2303100113213331-1311033211000232-3103123113213023-1120231220220022"></a>

<a id="canonical-1310332101333302-2210110123332020-0212103022210031-3331100110122203-3111120302120231-3121130202012002-0133022330300103-3023220030013221"></a>

## location property — blindfold_secret_info / 012021013011 / 5

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

<a id="canonical-3003310333102211-0301120311203122-0321033133002021-3313122331223223-2303013002030023-0033122300322210-0322120002332301-2203022231300111"></a>

<a id="canonical-2132212301211330-3112022333301322-3013132232033013-2302101310331230-2123201202113203-3231131300110123-0033013003331213-0022323033313321"></a>

## store_provider property — blindfold_secret_info / 012021013011 / 6

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

<a id="canonical-1221012203202220-2021022332221233-0302013031331202-1120030301111323-2311312110210333-2120123220033001-1011022123210232-2032203213003231"></a>

## Next pages — blindfold_secret_info / 012021013011 / 7

- [f5_big_ip_aws_service.admin_password](data-sources--nfv_service--reference--group-001.md#canonical-1321011123103003-0223330110312101-0111101021303212-3301023331031210-0021230321333022-0133210003220220-3321222301023200-0111121311331331)
- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-2212320103310132-2221300222221032-1233031123200120-3022110322311231-3301003023302233-3133232120201022-0322021211220213-1332311000003300)

<a id="canonical-0011321212002310-3013230101301223-2330001032101120-0321312211320102-3120320103221010-3103322221303132-1213230200011001-2211121120130201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1111212031232333-0011323211333233-3132020331230001-3000303132002103-0120000312022120-1110301200203031-2132311031203230-3310322103103303"></a>

## f5_big_ip_aws_service.admin_password.clear_secret_info — clear_secret_info / 202202130312 / 2

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-2212320103310132-2221300222221032-1233031123200120-3022110322311231-3301003023302233-3133232120201022-0322021211220213-1332311000003300)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-2313010333323131-3010030231311121-2301323120300012-0122321020201331-3020223211220022-0100121100023011-2323321311300323-0122100132003020)
- [f5_big_ip_aws_service](data-sources--nfv_service--reference--group-001.md#canonical-3121203202232102-3301112102333133-2103303100231002-1321223111220201-3103311113020000-0303332002322201-3021331113310133-2201221310222221)
- [f5_big_ip_aws_service.admin_password](data-sources--nfv_service--reference--group-001.md#canonical-1321011123103003-0223330110312101-0111101021303212-3301023331031210-0021230321333022-0133210003220220-3321222301023200-0111121311331331)
- f5_big_ip_aws_service.admin_password.clear_secret_info

<a id="canonical-1202031110132212-3212011233300213-2131223321233010-1000032230101221-3211201332333313-0011213130133333-0020333111112003-3301020121020010"></a>

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

<a id="canonical-0303121333321122-2333222301112131-2022101010230333-1002121011130131-0000331211212010-0001220212202233-1220220020121100-0230110233232020"></a>

## Direct properties — clear_secret_info / 202202130312 / 3

<a id="canonical-1213030032230021-3011131310320110-1221032002021113-1032222320311101-3113301111020320-0110010313000332-0032111013112222-3031030303303331"></a>

<a id="canonical-0223302330012130-3221301201100122-3320330123132101-2122333112112320-2022300302103012-1322312233032103-1301333333323220-2030323231000012"></a>

## provider_ref property — clear_secret_info / 202202130312 / 4

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-1213013020300133-2130222000312102-0231000123103132-2112211300301310-1322033003033002-3320031102110003-2300130232030313-1023210121223210"></a>

<a id="canonical-0221302021221210-2101013020320322-0110132110301203-3300001331321120-0033000333331023-0212133032010332-1131213111303210-2233302313111211"></a>

## URL property — clear_secret_info / 202202130312 / 5

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

<a id="canonical-2020302013301212-3203032200230311-0213131113000230-3232332332020111-3033301222012200-2123230130020022-3231322321231202-3222113312203210"></a>

## Next pages — clear_secret_info / 202202130312 / 6

- [f5_big_ip_aws_service.admin_password](data-sources--nfv_service--reference--group-001.md#canonical-1321011123103003-0223330110312101-0111101021303212-3301023331031210-0021230321333022-0133210003220220-3321222301023200-0111121311331331)
- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-2212320103310132-2221300222221032-1233031123200120-3022110322311231-3301003023302233-3133232120201022-0322021211220213-1332311000003300)

<a id="canonical-1130322021203231-2300212333020310-3210110221221221-1003120300332113-1323310310320132-3100020313313322-3031020210022022-2110121011300330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0202201232321132-3232330213120311-2111023001103201-3200231223022200-1202331333032032-2330202002310132-3033023010230230-1232012222020232"></a>

## f5_big_ip_aws_service.aws_tgw_site_params — aws_tgw_site_params / 210201122131 / 2

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-2212320103310132-2221300222221032-1233031123200120-3022110322311231-3301003023302233-3133232120201022-0322021211220213-1332311000003300)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-2313010333323131-3010030231311121-2301323120300012-0122321020201331-3020223211220022-0100121100023011-2323321311300323-0122100132003020)
- [f5_big_ip_aws_service](data-sources--nfv_service--reference--group-001.md#canonical-3121203202232102-3301112102333133-2103303100231002-1321223111220201-3103311113020000-0303332002322201-3021331113310133-2201221310222221)
- f5_big_ip_aws_service.aws_tgw_site_params

<a id="canonical-1023032330102020-1233233323001332-0332223010131311-0301003330030102-3120221211003023-1001212032320202-0332230021030201-1320220120030321"></a>

Type: `"single"`. Computed.

BIG-IP AWS TGW Site. BIG-IP AWS TGW site specification.

Upstream description:

BIG-IP AWS TGW site specification.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3303221121122101-3032113211103213-2332233013313130-0203302031012211-2321121211030331-0320023012030220-3222132133012132-2333000321012322"></a>

## Direct properties — aws_tgw_site_params / 210201122131 / 3

- [aws_tgw_site](data-sources--nfv_service--reference--group-001.md#canonical-2020003301300312-3101112321320332-2330012023223333-3212012331110301-0023332310200331-2132121123021011-2132123301130122-0230302003313213): complete subsection reference.

<a id="canonical-0123111103322231-2030032112303011-1203230132301112-0033211130310131-0300130232102321-1332021210021032-3300311333313300-0032222001312222"></a>

## Next pages — aws_tgw_site_params / 210201122131 / 4

- [f5_big_ip_aws_service.aws_tgw_site_params.aws_tgw_site](data-sources--nfv_service--reference--group-001.md#canonical-2020003301300312-3101112321320332-2330012023223333-3212012331110301-0023332310200331-2132121123021011-2132123301130122-0230302003313213)
- [f5_big_ip_aws_service](data-sources--nfv_service--reference--group-001.md#canonical-3121203202232102-3301112102333133-2103303100231002-1321223111220201-3103311113020000-0303332002322201-3021331113310133-2201221310222221)
- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-2212320103310132-2221300222221032-1233031123200120-3022110322311231-3301003023302233-3133232120201022-0322021211220213-1332311000003300)

<a id="canonical-2020003301300312-3101112321320332-2330012023223333-3212012331110301-0023332310200331-2132121123021011-2132123301130122-0230302003313213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1012231212233312-2212100130122130-3202030133233102-2332321322301031-1301010312331131-1300310132230000-3302200230222000-0033321130031302"></a>

## f5_big_ip_aws_service.aws_tgw_site_params.aws_tgw_site — aws_tgw_site / 320223210121 / 2

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-2212320103310132-2221300222221032-1233031123200120-3022110322311231-3301003023302233-3133232120201022-0322021211220213-1332311000003300)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-2313010333323131-3010030231311121-2301323120300012-0122321020201331-3020223211220022-0100121100023011-2323321311300323-0122100132003020)
- [f5_big_ip_aws_service](data-sources--nfv_service--reference--group-001.md#canonical-3121203202232102-3301112102333133-2103303100231002-1321223111220201-3103311113020000-0303332002322201-3021331113310133-2201221310222221)
- [f5_big_ip_aws_service.aws_tgw_site_params](data-sources--nfv_service--reference--group-001.md#canonical-1130322021203231-2300212333020310-3210110221221221-1003120300332113-1323310310320132-3100020313313322-3031020210022022-2110121011300330)
- f5_big_ip_aws_service.aws_tgw_site_params.aws_tgw_site

<a id="canonical-1010020320213320-3103132210221212-0302122111111330-0333131100203133-3102220211332332-3130331302020130-2211320202313113-2330101212133200"></a>

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

<a id="canonical-3020312202331213-3011230333332102-0010332213303220-2323232210320223-2320001113033101-3013320101221212-1312021222013333-3330213311022212"></a>

## Direct properties — aws_tgw_site / 320223210121 / 3

<a id="canonical-3203121021223331-1210221010331110-3323011032103102-1011230201113303-3031131301111322-2100031331321130-3202031130333112-0230101033203330"></a>

<a id="canonical-2213000012220222-0231223203220023-2132301122210003-0122231013133231-1301121303122023-1333220120103100-1330300131010132-2233303121020223"></a>

## name property — aws_tgw_site / 320223210121 / 4

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

<a id="canonical-2121231111121232-3123320130230313-1200320302011211-0231321032031202-0010120101001321-1203311012130022-0303000222321010-1322301111332331"></a>

<a id="canonical-2122103133230211-1331102313231321-1103003001331320-0331311211203021-0232233130032000-0130322033211121-3112211131222120-1332321202130201"></a>

## namespace property — aws_tgw_site / 320223210121 / 5

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

<a id="canonical-2120003030330223-1032220221021203-2023012113032300-1122031212233210-1021022111030021-2210233301111023-3211221003121321-1103123013013100"></a>

<a id="canonical-0112312112012313-0333333000311320-0331113200030332-2311131103132232-2032202011030110-0320211122311230-0110221221232123-0332122110023320"></a>

## tenant property — aws_tgw_site / 320223210121 / 6

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

<a id="canonical-2302030102320002-2211222131322132-3211213232322232-0210230013013022-1320312033013010-0233020231311020-0123322333303000-3301133032103312"></a>

## Next pages — aws_tgw_site / 320223210121 / 7

- [f5_big_ip_aws_service.aws_tgw_site_params](data-sources--nfv_service--reference--group-001.md#canonical-1130322021203231-2300212333020310-3210110221221221-1003120300332113-1323310310320132-3100020313313322-3031020210022022-2110121011300330)
- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-2212320103310132-2221300222221032-1233031123200120-3022110322311231-3301003023302233-3133232120201022-0322021211220213-1332311000003300)

<a id="canonical-1100221120003321-3013113033000220-0121013203232101-0223100022123001-0102122201021202-3023132102201031-0030303212302022-0220303202322013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3312303112023022-0100132023313103-3002123031012220-2223121013100020-0303100001302103-3331133300030001-2202330012123321-3013210303221012"></a>

## f5_big_ip_aws_service.endpoint_service — endpoint_service / 021301022231 / 2

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-2212320103310132-2221300222221032-1233031123200120-3022110322311231-3301003023302233-3133232120201022-0322021211220213-1332311000003300)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-2313010333323131-3010030231311121-2301323120300012-0122321020201331-3020223211220022-0100121100023011-2323321311300323-0122100132003020)
- [f5_big_ip_aws_service](data-sources--nfv_service--reference--group-001.md#canonical-3121203202232102-3301112102333133-2103303100231002-1321223111220201-3103311113020000-0303332002322201-3021331113310133-2201221310222221)
- f5_big_ip_aws_service.endpoint_service

<a id="canonical-2102032020301333-2031221001211012-0233322230323103-3321200112100000-0120101110020230-2100102323111221-1122312221333020-2322022001320220"></a>

Type: `"single"`. Computed.

Endpoint Service is a type of NFV service where the packets are destined to NFV and service modifies
the destination with a new destination address.

Upstream description:

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

<a id="canonical-1302031011213101-1223202122223201-3131131012211331-1011320230111112-0031210013223303-3110102201220332-3033221121021313-2103322313232201"></a>

## Direct properties — endpoint_service / 021301022231 / 3

- [advertise_on_slo_ip](data-sources--nfv_service--reference--group-001.md#canonical-3133220003102223-0111331313121130-3221131310211011-3301311203303332-1201131110233320-1212213333322231-0132232110221230-1230200302331102): complete subsection reference.

- [advertise_on_slo_ip_external](data-sources--nfv_service--reference--group-001.md#canonical-3231102101102033-2333001220131233-2120213210302103-3002103222310310-0231223023302211-0130331230023312-2001011100021111-3131220110111010): complete subsection reference.

- [automatic_vip](data-sources--nfv_service--reference--group-001.md#canonical-0333230310002130-2321211131011030-3113123131303102-0120320110302311-0323002123332313-0022221013003200-3001222032223002-3100113323311033): complete subsection reference.

<a id="canonical-1321211322210330-3103122021213020-1000013022312021-2230023210122030-2111010320203111-1100021211320001-0122032302202122-2311003330131200"></a>

<a id="canonical-2321301220001122-3311213123200213-2030022013310203-3203012303111111-2112220022203112-0110011310132222-0021121301003120-3100201130203032"></a>

## configured_vip property — endpoint_service / 021301022231 / 4

Type: `"string"`. Computed.

Exclusive with \[automatic\_vip\] Enter IP address for the default VIP.

Upstream description:

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
    "ves.io.schema.rules.string.ip": "true",
    "ves.io.schema.rules.string.not_in": "192.0.2.26"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ip": "true",
    "ves.io.schema.rules.string.not_in": "192.0.2.26"
  }
}
```

- [custom_tcp_ports](data-sources--nfv_service--reference--group-001.md#canonical-0213322323131212-1001122322320131-2123022303303303-3323230213213130-3012321230000310-3010320103232323-1203032232302030-2233101033210301): complete subsection reference.

- [custom_udp_ports](data-sources--nfv_service--reference--group-001.md#canonical-2222122210103230-1131323112201210-3301032223301002-3002210110332123-0120110321323130-2132321002301000-1313332131011110-3211011222211220): complete subsection reference.

- [default_tcp_ports](data-sources--nfv_service--reference--group-001.md#canonical-0111313031020233-3111102301113132-3102322010322221-2113320130033322-0131133030003030-3220211303230302-3003313313302002-1103010311323030): complete subsection reference.

- [disable_advertise_on_slo_ip](data-sources--nfv_service--reference--group-001.md#canonical-0111010031211300-2221210230303232-1213331122313130-3311113220333323-1111113120313122-1102202220230210-0002002122331210-0321301332033033): complete subsection reference.

- [http_port](data-sources--nfv_service--reference--group-001.md#canonical-3120330013002133-2303230122133023-0112123320322113-2302002021001231-3122033221301330-0103011033100011-2231031101222321-3003322003211000): complete subsection reference.

- [https_port](data-sources--nfv_service--reference--group-001.md#canonical-2132212120230331-0030101332230130-2211231130201303-1133310210021132-0221332020013000-3213323030231123-0330331011131012-2311110220130022): complete subsection reference.

- [no_tcp_ports](data-sources--nfv_service--reference--group-001.md#canonical-2031223303233203-0333000312131101-1020313012131303-2010023311132121-3022101111302003-2102231231011103-3023013230122120-3212301313330230): complete subsection reference.

- [no_udp_ports](data-sources--nfv_service--reference--group-001.md#canonical-1211332033202133-0033000313310303-3212021231010121-2320200313300213-1222203000210103-2322212212121121-0030203223133230-3210220110101211): complete subsection reference.

<a id="canonical-3210123013310112-1313301220231332-0030032310221200-1231211312332030-0030031300323023-0102110130202300-2233230012331031-1132330300230103"></a>

## Next pages — endpoint_service / 021301022231 / 5

- [f5_big_ip_aws_service.endpoint_service.advertise_on_slo_ip](data-sources--nfv_service--reference--group-001.md#canonical-3133220003102223-0111331313121130-3221131310211011-3301311203303332-1201131110233320-1212213333322231-0132232110221230-1230200302331102)
- [f5_big_ip_aws_service.endpoint_service.advertise_on_slo_ip_external](data-sources--nfv_service--reference--group-001.md#canonical-3231102101102033-2333001220131233-2120213210302103-3002103222310310-0231223023302211-0130331230023312-2001011100021111-3131220110111010)
- [f5_big_ip_aws_service.endpoint_service.automatic_vip](data-sources--nfv_service--reference--group-001.md#canonical-0333230310002130-2321211131011030-3113123131303102-0120320110302311-0323002123332313-0022221013003200-3001222032223002-3100113323311033)
- [f5_big_ip_aws_service.endpoint_service.custom_tcp_ports](data-sources--nfv_service--reference--group-001.md#canonical-0213322323131212-1001122322320131-2123022303303303-3323230213213130-3012321230000310-3010320103232323-1203032232302030-2233101033210301)
- [f5_big_ip_aws_service.endpoint_service.custom_udp_ports](data-sources--nfv_service--reference--group-001.md#canonical-2222122210103230-1131323112201210-3301032223301002-3002210110332123-0120110321323130-2132321002301000-1313332131011110-3211011222211220)
- [f5_big_ip_aws_service.endpoint_service.default_tcp_ports](data-sources--nfv_service--reference--group-001.md#canonical-0111313031020233-3111102301113132-3102322010322221-2113320130033322-0131133030003030-3220211303230302-3003313313302002-1103010311323030)
- [f5_big_ip_aws_service.endpoint_service.disable_advertise_on_slo_ip](data-sources--nfv_service--reference--group-001.md#canonical-0111010031211300-2221210230303232-1213331122313130-3311113220333323-1111113120313122-1102202220230210-0002002122331210-0321301332033033)
- [f5_big_ip_aws_service.endpoint_service.http_port](data-sources--nfv_service--reference--group-001.md#canonical-3120330013002133-2303230122133023-0112123320322113-2302002021001231-3122033221301330-0103011033100011-2231031101222321-3003322003211000)
- [f5_big_ip_aws_service.endpoint_service.https_port](data-sources--nfv_service--reference--group-001.md#canonical-2132212120230331-0030101332230130-2211231130201303-1133310210021132-0221332020013000-3213323030231123-0330331011131012-2311110220130022)
- [f5_big_ip_aws_service.endpoint_service.no_tcp_ports](data-sources--nfv_service--reference--group-001.md#canonical-2031223303233203-0333000312131101-1020313012131303-2010023311132121-3022101111302003-2102231231011103-3023013230122120-3212301313330230)
- [f5_big_ip_aws_service.endpoint_service.no_udp_ports](data-sources--nfv_service--reference--group-001.md#canonical-1211332033202133-0033000313310303-3212021231010121-2320200313300213-1222203000210103-2322212212121121-0030203223133230-3210220110101211)
- [f5_big_ip_aws_service](data-sources--nfv_service--reference--group-001.md#canonical-3121203202232102-3301112102333133-2103303100231002-1321223111220201-3103311113020000-0303332002322201-3021331113310133-2201221310222221)
- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-2212320103310132-2221300222221032-1233031123200120-3022110322311231-3301003023302233-3133232120201022-0322021211220213-1332311000003300)

<a id="canonical-3133220003102223-0111331313121130-3221131310211011-3301311203303332-1201131110233320-1212213333322231-0132232110221230-1230200302331102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0022203032100110-0222010123210213-2301222012023200-1301033301102202-2200132332010023-0201230202311221-3013012311303132-1213032023011310"></a>

## f5_big_ip_aws_service.endpoint_service.advertise_on_slo_ip — advertise_on_slo_ip / 211103311011 / 2

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-2212320103310132-2221300222221032-1233031123200120-3022110322311231-3301003023302233-3133232120201022-0322021211220213-1332311000003300)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-2313010333323131-3010030231311121-2301323120300012-0122321020201331-3020223211220022-0100121100023011-2323321311300323-0122100132003020)
- [f5_big_ip_aws_service](data-sources--nfv_service--reference--group-001.md#canonical-3121203202232102-3301112102333133-2103303100231002-1321223111220201-3103311113020000-0303332002322201-3021331113310133-2201221310222221)
- [f5_big_ip_aws_service.endpoint_service](data-sources--nfv_service--reference--group-001.md#canonical-1100221120003321-3013113033000220-0121013203232101-0223100022123001-0102122201021202-3023132102201031-0030303212302022-0220303202322013)
- f5_big_ip_aws_service.endpoint_service.advertise_on_slo_ip

<a id="canonical-3222013222330332-0211222310100332-0002230130002330-1010130113302110-2323103102020223-1023330032200123-0103122001232021-2223330130333001"></a>

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

<a id="canonical-1131303330112232-1001130021233221-2221211213121133-2301002111000013-0033020313211212-0131013331232003-3303220302330230-1300200112323100"></a>

## Direct properties — advertise_on_slo_ip / 211103311011 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3103003131200332-0102313002012323-2100123133311310-3231311003222323-2311220212322201-0331121032011333-1233010123300011-2122203332311110"></a>

## Next pages — advertise_on_slo_ip / 211103311011 / 4

- [f5_big_ip_aws_service.endpoint_service](data-sources--nfv_service--reference--group-001.md#canonical-1100221120003321-3013113033000220-0121013203232101-0223100022123001-0102122201021202-3023132102201031-0030303212302022-0220303202322013)
- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-2212320103310132-2221300222221032-1233031123200120-3022110322311231-3301003023302233-3133232120201022-0322021211220213-1332311000003300)

<a id="canonical-3231102101102033-2333001220131233-2120213210302103-3002103222310310-0231223023302211-0130331230023312-2001011100021111-3131220110111010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3100322030101231-1003113102223211-0222101133220010-2023202322200332-1312220102101130-1200320131131321-2121110133020201-3131311233332221"></a>

## f5_big_ip_aws_service.endpoint_service.advertise_on_slo_ip_external — advertise_on_slo_ip_external / 111001101113 / 2

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-2212320103310132-2221300222221032-1233031123200120-3022110322311231-3301003023302233-3133232120201022-0322021211220213-1332311000003300)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-2313010333323131-3010030231311121-2301323120300012-0122321020201331-3020223211220022-0100121100023011-2323321311300323-0122100132003020)
- [f5_big_ip_aws_service](data-sources--nfv_service--reference--group-001.md#canonical-3121203202232102-3301112102333133-2103303100231002-1321223111220201-3103311113020000-0303332002322201-3021331113310133-2201221310222221)
- [f5_big_ip_aws_service.endpoint_service](data-sources--nfv_service--reference--group-001.md#canonical-1100221120003321-3013113033000220-0121013203232101-0223100022123001-0102122201021202-3023132102201031-0030303212302022-0220303202322013)
- f5_big_ip_aws_service.endpoint_service.advertise_on_slo_ip_external

<a id="canonical-1210312003330221-1031021202310023-0001001333111211-0201210010212230-3211123013111203-1330300331313303-0112001011230213-2110030312101120"></a>

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

<a id="canonical-0320201233110112-1203232301212032-3000103011102002-2003223101311122-2133330022111032-0330123021122322-2001323303033322-2212112122220222"></a>

## Direct properties — advertise_on_slo_ip_external / 111001101113 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2313030212010030-0122131331020103-3213112122303133-1231303132332112-3130130233031322-2210002302111332-0012022301123022-3122203100013133"></a>

## Next pages — advertise_on_slo_ip_external / 111001101113 / 4

- [f5_big_ip_aws_service.endpoint_service](data-sources--nfv_service--reference--group-001.md#canonical-1100221120003321-3013113033000220-0121013203232101-0223100022123001-0102122201021202-3023132102201031-0030303212302022-0220303202322013)
- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-2212320103310132-2221300222221032-1233031123200120-3022110322311231-3301003023302233-3133232120201022-0322021211220213-1332311000003300)

<a id="canonical-0333230310002130-2321211131011030-3113123131303102-0120320110302311-0323002123332313-0022221013003200-3001222032223002-3100113323311033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1312123323011322-3021111212233232-1210021203000033-2213220023031310-2232111232121232-2022122112210131-3101111123300333-0303213202001330"></a>

## f5_big_ip_aws_service.endpoint_service.automatic_vip — automatic_vip / 020123102200 / 2

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-2212320103310132-2221300222221032-1233031123200120-3022110322311231-3301003023302233-3133232120201022-0322021211220213-1332311000003300)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-2313010333323131-3010030231311121-2301323120300012-0122321020201331-3020223211220022-0100121100023011-2323321311300323-0122100132003020)
- [f5_big_ip_aws_service](data-sources--nfv_service--reference--group-001.md#canonical-3121203202232102-3301112102333133-2103303100231002-1321223111220201-3103311113020000-0303332002322201-3021331113310133-2201221310222221)
- [f5_big_ip_aws_service.endpoint_service](data-sources--nfv_service--reference--group-001.md#canonical-1100221120003321-3013113033000220-0121013203232101-0223100022123001-0102122201021202-3023132102201031-0030303212302022-0220303202322013)
- f5_big_ip_aws_service.endpoint_service.automatic_vip

<a id="canonical-2211133003033002-3323220103223323-0101023323110303-0202123111110000-3103020022103321-3020323221123120-3030020132101303-3230323202002203"></a>

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

<a id="canonical-0131133010121103-1310103112110033-2303221033320003-3232101210021213-0030112322121131-1113033003011033-3003103102113112-0003323301223130"></a>

## Direct properties — automatic_vip / 020123102200 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1210332332210300-0032113031233302-2312222223231132-0333302223010312-1133202303131110-0310223103013210-2030032300112220-0321311113221210"></a>

## Next pages — automatic_vip / 020123102200 / 4

- [f5_big_ip_aws_service.endpoint_service](data-sources--nfv_service--reference--group-001.md#canonical-1100221120003321-3013113033000220-0121013203232101-0223100022123001-0102122201021202-3023132102201031-0030303212302022-0220303202322013)
- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-2212320103310132-2221300222221032-1233031123200120-3022110322311231-3301003023302233-3133232120201022-0322021211220213-1332311000003300)

<a id="canonical-0213322323131212-1001122322320131-2123022303303303-3323230213213130-3012321230000310-3010320103232323-1203032232302030-2233101033210301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1201120002203200-3013011211020323-0203320023001023-1230001203300223-2010003211221012-1031231120101120-3200101222111312-0313001322213012"></a>

## f5_big_ip_aws_service.endpoint_service.custom_tcp_ports — custom_tcp_ports / 330223102210 / 2

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-2212320103310132-2221300222221032-1233031123200120-3022110322311231-3301003023302233-3133232120201022-0322021211220213-1332311000003300)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-2313010333323131-3010030231311121-2301323120300012-0122321020201331-3020223211220022-0100121100023011-2323321311300323-0122100132003020)
- [f5_big_ip_aws_service](data-sources--nfv_service--reference--group-001.md#canonical-3121203202232102-3301112102333133-2103303100231002-1321223111220201-3103311113020000-0303332002322201-3021331113310133-2201221310222221)
- [f5_big_ip_aws_service.endpoint_service](data-sources--nfv_service--reference--group-001.md#canonical-1100221120003321-3013113033000220-0121013203232101-0223100022123001-0102122201021202-3023132102201031-0030303212302022-0220303202322013)
- f5_big_ip_aws_service.endpoint_service.custom_tcp_ports

<a id="canonical-3303211321230323-2103233110022011-2333311220020202-2201311302332100-1120200002101022-1032213330332001-3231010121310323-2111110102031123"></a>

Type: `"single"`. Computed.

Port Range List. List of port ranges.

Upstream description:

List of port ranges.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1013310232303032-0111023100113012-0020102232233201-0021110131323121-2011211120120001-0213211321300330-2230012133323002-3013213102313001"></a>

## Direct properties — custom_tcp_ports / 330223102210 / 3

<a id="canonical-3003103212223022-2132210022233231-0323210112300222-2221300310303320-3333103002013110-3211313312110011-1323333330300311-3031001013202230"></a>

<a id="canonical-0011201122122321-3120010021122122-0020110130232003-0122133322222322-1323220223031221-2200233001203001-1332022111010123-2233100010102003"></a>

## ports property — custom_tcp_ports / 330223102210 / 4

Type: `["list", "string"]`. Computed.

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

<a id="canonical-1130320100333003-2310303030333003-2121020202032311-1220210021002011-0233322100132313-1133212132112032-3133003331203322-3103111301331013"></a>

## Next pages — custom_tcp_ports / 330223102210 / 5

- [f5_big_ip_aws_service.endpoint_service](data-sources--nfv_service--reference--group-001.md#canonical-1100221120003321-3013113033000220-0121013203232101-0223100022123001-0102122201021202-3023132102201031-0030303212302022-0220303202322013)
- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-2212320103310132-2221300222221032-1233031123200120-3022110322311231-3301003023302233-3133232120201022-0322021211220213-1332311000003300)

<a id="canonical-2222122210103230-1131323112201210-3301032223301002-3002210110332123-0120110321323130-2132321002301000-1313332131011110-3211011222211220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2221310130100213-3333322203103130-2030230300223211-3121123101313323-1020123331212010-1211031211200022-0003213131322322-2320102213201131"></a>

## f5_big_ip_aws_service.endpoint_service.custom_udp_ports — custom_udp_ports / 003111112111 / 2

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-2212320103310132-2221300222221032-1233031123200120-3022110322311231-3301003023302233-3133232120201022-0322021211220213-1332311000003300)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-2313010333323131-3010030231311121-2301323120300012-0122321020201331-3020223211220022-0100121100023011-2323321311300323-0122100132003020)
- [f5_big_ip_aws_service](data-sources--nfv_service--reference--group-001.md#canonical-3121203202232102-3301112102333133-2103303100231002-1321223111220201-3103311113020000-0303332002322201-3021331113310133-2201221310222221)
- [f5_big_ip_aws_service.endpoint_service](data-sources--nfv_service--reference--group-001.md#canonical-1100221120003321-3013113033000220-0121013203232101-0223100022123001-0102122201021202-3023132102201031-0030303212302022-0220303202322013)
- f5_big_ip_aws_service.endpoint_service.custom_udp_ports

<a id="canonical-0123132012113300-2332213331331000-2031221221000011-0023323201312303-3003310332011202-3112132110223330-0101230332201201-0323012212332023"></a>

Type: `"single"`. Computed.

Port Range List. List of port ranges.

Upstream description:

List of port ranges.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0300023321300203-0333030233333020-2111301333212122-2011033113333022-2133122302113203-0031233110203213-1112120023313000-2332333021000032"></a>

## Direct properties — custom_udp_ports / 003111112111 / 3

<a id="canonical-3102120312302230-0101321303222301-1302130013332031-0233210212220111-0101203203323202-3100131322312001-2212002301311113-1333331111210032"></a>

<a id="canonical-3231313030213212-1113331231111213-1221021301332133-0011320200312012-3020213323102213-2210032310030230-3121220233111120-1320011013002012"></a>

## ports property — custom_udp_ports / 003111112111 / 4

Type: `["list", "string"]`. Computed.

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

<a id="canonical-2023132330112223-1221113033212112-2020322023111210-0312101322023301-0302333313330310-3202333220113300-1321332132301210-1013301322220330"></a>

## Next pages — custom_udp_ports / 003111112111 / 5

- [f5_big_ip_aws_service.endpoint_service](data-sources--nfv_service--reference--group-001.md#canonical-1100221120003321-3013113033000220-0121013203232101-0223100022123001-0102122201021202-3023132102201031-0030303212302022-0220303202322013)
- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-2212320103310132-2221300222221032-1233031123200120-3022110322311231-3301003023302233-3133232120201022-0322021211220213-1332311000003300)

<a id="canonical-0111313031020233-3111102301113132-3102322010322221-2113320130033322-0131133030003030-3220211303230302-3003313313302002-1103010311323030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3310301222222133-0323211133300332-0133330202103310-1000103133133033-3230323031000332-1230313330032221-3000231000311113-1130130303301002"></a>

## f5_big_ip_aws_service.endpoint_service.default_tcp_ports — default_tcp_ports / 102203321310 / 2

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-2212320103310132-2221300222221032-1233031123200120-3022110322311231-3301003023302233-3133232120201022-0322021211220213-1332311000003300)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-2313010333323131-3010030231311121-2301323120300012-0122321020201331-3020223211220022-0100121100023011-2323321311300323-0122100132003020)
- [f5_big_ip_aws_service](data-sources--nfv_service--reference--group-001.md#canonical-3121203202232102-3301112102333133-2103303100231002-1321223111220201-3103311113020000-0303332002322201-3021331113310133-2201221310222221)
- [f5_big_ip_aws_service.endpoint_service](data-sources--nfv_service--reference--group-001.md#canonical-1100221120003321-3013113033000220-0121013203232101-0223100022123001-0102122201021202-3023132102201031-0030303212302022-0220303202322013)
- f5_big_ip_aws_service.endpoint_service.default_tcp_ports

<a id="canonical-2003222221331333-3011032112013113-0333210312032000-1313331311213020-3113333111130010-2010112032201212-1223333322211300-0030101221300311"></a>

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

<a id="canonical-1321212231120312-1323312103230130-3121002333212300-1233122212131130-3013013230022030-3223003202311223-1301213233022330-2200300022130313"></a>

## Direct properties — default_tcp_ports / 102203321310 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3111222000320201-0322111010121010-3302222202313101-2101001211021310-1121210032313202-0320131222021230-2113101200030220-1332002003333311"></a>

## Next pages — default_tcp_ports / 102203321310 / 4

- [f5_big_ip_aws_service.endpoint_service](data-sources--nfv_service--reference--group-001.md#canonical-1100221120003321-3013113033000220-0121013203232101-0223100022123001-0102122201021202-3023132102201031-0030303212302022-0220303202322013)
- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-2212320103310132-2221300222221032-1233031123200120-3022110322311231-3301003023302233-3133232120201022-0322021211220213-1332311000003300)

<a id="canonical-0111010031211300-2221210230303232-1213331122313130-3311113220333323-1111113120313122-1102202220230210-0002002122331210-0321301332033033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2221230301312202-2112230112013222-3312032021021113-3233212012003210-1231111211320303-1131210030122002-1212000131112203-2003222131023231"></a>

## f5_big_ip_aws_service.endpoint_service.disable_advertise_on_slo_ip — disable_advertise_on_slo_ip / 131113321300 / 2

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-2212320103310132-2221300222221032-1233031123200120-3022110322311231-3301003023302233-3133232120201022-0322021211220213-1332311000003300)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-2313010333323131-3010030231311121-2301323120300012-0122321020201331-3020223211220022-0100121100023011-2323321311300323-0122100132003020)
- [f5_big_ip_aws_service](data-sources--nfv_service--reference--group-001.md#canonical-3121203202232102-3301112102333133-2103303100231002-1321223111220201-3103311113020000-0303332002322201-3021331113310133-2201221310222221)
- [f5_big_ip_aws_service.endpoint_service](data-sources--nfv_service--reference--group-001.md#canonical-1100221120003321-3013113033000220-0121013203232101-0223100022123001-0102122201021202-3023132102201031-0030303212302022-0220303202322013)
- f5_big_ip_aws_service.endpoint_service.disable_advertise_on_slo_ip

<a id="canonical-0111122032013011-2311120202300221-1213330000033031-2012311233300122-0322310020132321-3223211223220303-0021213223032003-1211122022033313"></a>

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

<a id="canonical-3221330133200130-3212131213300031-3331000123010011-0113210332112032-0332113331223001-3231010312120023-2332323223321103-3310022220021110"></a>

## Direct properties — disable_advertise_on_slo_ip / 131113321300 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0302110230222330-3111100110333233-2130130202023111-2102112202102213-2123103221002031-3023003201320203-3123331302320332-0002132101132113"></a>

## Next pages — disable_advertise_on_slo_ip / 131113321300 / 4

- [f5_big_ip_aws_service.endpoint_service](data-sources--nfv_service--reference--group-001.md#canonical-1100221120003321-3013113033000220-0121013203232101-0223100022123001-0102122201021202-3023132102201031-0030303212302022-0220303202322013)
- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-2212320103310132-2221300222221032-1233031123200120-3022110322311231-3301003023302233-3133232120201022-0322021211220213-1332311000003300)

<a id="canonical-3120330013002133-2303230122133023-0112123320322113-2302002021001231-3122033221301330-0103011033100011-2231031101222321-3003322003211000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2303221132300323-0020311123012211-3011203320102012-2132002022022133-3002022300222222-3120131200221033-2003232012111310-1232231203333022"></a>

## f5_big_ip_aws_service.endpoint_service.http_port — http_port / 210332311100 / 2

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-2212320103310132-2221300222221032-1233031123200120-3022110322311231-3301003023302233-3133232120201022-0322021211220213-1332311000003300)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-2313010333323131-3010030231311121-2301323120300012-0122321020201331-3020223211220022-0100121100023011-2323321311300323-0122100132003020)
- [f5_big_ip_aws_service](data-sources--nfv_service--reference--group-001.md#canonical-3121203202232102-3301112102333133-2103303100231002-1321223111220201-3103311113020000-0303332002322201-3021331113310133-2201221310222221)
- [f5_big_ip_aws_service.endpoint_service](data-sources--nfv_service--reference--group-001.md#canonical-1100221120003321-3013113033000220-0121013203232101-0223100022123001-0102122201021202-3023132102201031-0030303212302022-0220303202322013)
- f5_big_ip_aws_service.endpoint_service.http_port

<a id="canonical-0210321120300122-2222022122301331-0112303233133203-1333113201302113-2031002112221123-1203003000120230-0302302031112233-1103010110120232"></a>

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

<a id="canonical-1021012013312332-0112100030021131-3222232023312310-3330122203322211-3313111023323330-0310000321110303-0222033032313032-0321200322003030"></a>

## Direct properties — http_port / 210332311100 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0113300131103230-1031103313303020-2103230222132202-2211210213313312-1131113303122313-1311232333112322-0320300020120123-3230221132113012"></a>

## Next pages — http_port / 210332311100 / 4

- [f5_big_ip_aws_service.endpoint_service](data-sources--nfv_service--reference--group-001.md#canonical-1100221120003321-3013113033000220-0121013203232101-0223100022123001-0102122201021202-3023132102201031-0030303212302022-0220303202322013)
- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-2212320103310132-2221300222221032-1233031123200120-3022110322311231-3301003023302233-3133232120201022-0322021211220213-1332311000003300)

<a id="canonical-2132212120230331-0030101332230130-2211231130201303-1133310210021132-0221332020013000-3213323030231123-0330331011131012-2311110220130022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1002200022313220-1331122132310202-1321030130303321-2121331310210030-0133103000033231-0131301232212011-1312333013010012-2331302102202300"></a>

## f5_big_ip_aws_service.endpoint_service.https_port — https_port / 220310010212 / 2

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-2212320103310132-2221300222221032-1233031123200120-3022110322311231-3301003023302233-3133232120201022-0322021211220213-1332311000003300)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-2313010333323131-3010030231311121-2301323120300012-0122321020201331-3020223211220022-0100121100023011-2323321311300323-0122100132003020)
- [f5_big_ip_aws_service](data-sources--nfv_service--reference--group-001.md#canonical-3121203202232102-3301112102333133-2103303100231002-1321223111220201-3103311113020000-0303332002322201-3021331113310133-2201221310222221)
- [f5_big_ip_aws_service.endpoint_service](data-sources--nfv_service--reference--group-001.md#canonical-1100221120003321-3013113033000220-0121013203232101-0223100022123001-0102122201021202-3023132102201031-0030303212302022-0220303202322013)
- f5_big_ip_aws_service.endpoint_service.https_port

<a id="canonical-2031102231123010-0030111011112002-3103023001012110-1131303301121101-3232103213202001-3321132132233202-2300200033011323-1223212302232233"></a>

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

<a id="canonical-3201012211021002-3211232332010113-3211023312022102-1013233021200201-2113212213111002-1022022131002101-0312013330110332-3100332020233121"></a>

## Direct properties — https_port / 220310010212 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1012330111312023-1313200210311311-0313203030312323-2101320320020013-1111113313210323-0002013312022130-1120210123301213-3222221203001312"></a>

## Next pages — https_port / 220310010212 / 4

- [f5_big_ip_aws_service.endpoint_service](data-sources--nfv_service--reference--group-001.md#canonical-1100221120003321-3013113033000220-0121013203232101-0223100022123001-0102122201021202-3023132102201031-0030303212302022-0220303202322013)
- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-2212320103310132-2221300222221032-1233031123200120-3022110322311231-3301003023302233-3133232120201022-0322021211220213-1332311000003300)

<a id="canonical-2031223303233203-0333000312131101-1020313012131303-2010023311132121-3022101111302003-2102231231011103-3023013230122120-3212301313330230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3002222300221110-0010033332010322-3200223013321131-3010311023212312-1020303121121123-1012022233013002-2110122122323201-0333010000310032"></a>

## f5_big_ip_aws_service.endpoint_service.no_tcp_ports — no_tcp_ports / 201011231133 / 2

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-2212320103310132-2221300222221032-1233031123200120-3022110322311231-3301003023302233-3133232120201022-0322021211220213-1332311000003300)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-2313010333323131-3010030231311121-2301323120300012-0122321020201331-3020223211220022-0100121100023011-2323321311300323-0122100132003020)
- [f5_big_ip_aws_service](data-sources--nfv_service--reference--group-001.md#canonical-3121203202232102-3301112102333133-2103303100231002-1321223111220201-3103311113020000-0303332002322201-3021331113310133-2201221310222221)
- [f5_big_ip_aws_service.endpoint_service](data-sources--nfv_service--reference--group-001.md#canonical-1100221120003321-3013113033000220-0121013203232101-0223100022123001-0102122201021202-3023132102201031-0030303212302022-0220303202322013)
- f5_big_ip_aws_service.endpoint_service.no_tcp_ports

<a id="canonical-3301111110211233-2032030200120103-0001202120131113-0333212211023003-1111003200131230-2021121030111331-1311302132210012-1002333331132130"></a>

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

<a id="canonical-2001102033202212-2131013132131133-3111122201103032-3212231310001301-1131031211330213-3021311031333211-0002213221130031-0002001010002200"></a>

## Direct properties — no_tcp_ports / 201011231133 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2233213322221010-3212312012323022-1002301132302222-0012120002122131-0312120033231121-3203200112112130-2320023300313213-0011210110020300"></a>

## Next pages — no_tcp_ports / 201011231133 / 4

- [f5_big_ip_aws_service.endpoint_service](data-sources--nfv_service--reference--group-001.md#canonical-1100221120003321-3013113033000220-0121013203232101-0223100022123001-0102122201021202-3023132102201031-0030303212302022-0220303202322013)
- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-2212320103310132-2221300222221032-1233031123200120-3022110322311231-3301003023302233-3133232120201022-0322021211220213-1332311000003300)

<a id="canonical-1211332033202133-0033000313310303-3212021231010121-2320200313300213-1222203000210103-2322212212121121-0030203223133230-3210220110101211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2111130120311211-0303212101003220-3131031132230313-1031321123110123-2231110230012000-0000312332210132-3032222332103301-0130311321133200"></a>

## f5_big_ip_aws_service.endpoint_service.no_udp_ports — no_udp_ports / 133020302013 / 2

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-2212320103310132-2221300222221032-1233031123200120-3022110322311231-3301003023302233-3133232120201022-0322021211220213-1332311000003300)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-2313010333323131-3010030231311121-2301323120300012-0122321020201331-3020223211220022-0100121100023011-2323321311300323-0122100132003020)
- [f5_big_ip_aws_service](data-sources--nfv_service--reference--group-001.md#canonical-3121203202232102-3301112102333133-2103303100231002-1321223111220201-3103311113020000-0303332002322201-3021331113310133-2201221310222221)
- [f5_big_ip_aws_service.endpoint_service](data-sources--nfv_service--reference--group-001.md#canonical-1100221120003321-3013113033000220-0121013203232101-0223100022123001-0102122201021202-3023132102201031-0030303212302022-0220303202322013)
- f5_big_ip_aws_service.endpoint_service.no_udp_ports

<a id="canonical-3132223311022000-1302212113332033-2312101113201202-3012012302022033-3303122020331102-1330101113122120-0033123111123323-0330302021032232"></a>

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

<a id="canonical-3002303113321303-3230231021213222-1211001331033100-3121320022332130-0203203210223130-3103200121210222-1111130323213121-1021112131303213"></a>

## Direct properties — no_udp_ports / 133020302013 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2011213310030322-3112223212003012-3133220302302200-0223230212212311-3032032100031212-3110202311220211-3000121133031333-3213010313313031"></a>

## Next pages — no_udp_ports / 133020302013 / 4

- [f5_big_ip_aws_service.endpoint_service](data-sources--nfv_service--reference--group-001.md#canonical-1100221120003321-3013113033000220-0121013203232101-0223100022123001-0102122201021202-3023132102201031-0030303212302022-0220303202322013)
- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-2212320103310132-2221300222221032-1233031123200120-3022110322311231-3301003023302233-3133232120201022-0322021211220213-1332311000003300)

<a id="canonical-2210031222201112-3331110222012210-2132033313310012-3320121221230020-0131021130301020-0223033002000032-3021303122131320-2001020121212112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2131300201131332-2332201223113002-3001330321113023-3123022213131200-1222322323100321-2312103202020012-2303033120111201-1330032010311000"></a>

## f5_big_ip_aws_service.market_place_image — market_place_image / 133201330123 / 2

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-2212320103310132-2221300222221032-1233031123200120-3022110322311231-3301003023302233-3133232120201022-0322021211220213-1332311000003300)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-2313010333323131-3010030231311121-2301323120300012-0122321020201331-3020223211220022-0100121100023011-2323321311300323-0122100132003020)
- [f5_big_ip_aws_service](data-sources--nfv_service--reference--group-001.md#canonical-3121203202232102-3301112102333133-2103303100231002-1321223111220201-3103311113020000-0303332002322201-3021331113310133-2201221310222221)
- f5_big_ip_aws_service.market_place_image

<a id="canonical-1230320202121112-2210332203231211-1311220012032233-3011113003020230-0131303213211021-2302212122000003-1322131212121302-3003022001123313"></a>

Type: `"single"`. Computed.

BIG-IP AWS Pay as You Go Image Selection.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-ami_choice": "[\"AWAFPayG200Mbps\",\"AWAFPayG3Gbps\",\"BestPlusPayG200Mbps\",\"best_plus_payg_1gbps\"]"
}
```
