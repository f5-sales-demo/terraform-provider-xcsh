---
page_title: "xcsh_bigip_http_proxy reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_bigip_http_proxy reference."
---

# xcsh_bigip_http_proxy reference

<a id="canonical-0101201332222201-3302033021001322-1302101210311331-2210333210320003-0002033231003211-1133020003110203-0200133113131131-0101001111122000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Property reference

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)
- Property reference

<a id="canonical-0320011211132303-1020001003133001-0033303303011230-0022333213012200-1323020030031103-3021221232131220-0130302311103331-1313121031211232"></a>

### Direct properties for `xcsh_bigip_http_proxy`

- [advanced_profile](resources--bigip_http_proxy--reference--group-001.md#canonical-2112133030313203-1202233221221323-0231033212122201-1200202122332313-2022030302010022-1221300001120300-3120211033311213-2323103320212230): complete subsection reference.

<a id="canonical-2313003200321320-1322131322312032-1123031210113332-0230222303330322-3110303033332210-0102330003211132-3213223121223322-3221302223313313"></a>

<a id="canonical-2011132030101211-2230121203223021-0021323010212213-3132330122332230-3002220312003112-0232330221122000-0330000300210331-2013132330010211"></a>

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

- [ddos_profile](resources--bigip_http_proxy--reference--group-001.md#canonical-3223222331220311-2021021213301233-3110010022113211-1113021011323210-1100223221332333-3200220332110002-3301021103103321-3002222130011323): complete subsection reference.

<a id="canonical-3223112011313112-2300120332112033-2323203201323103-0110010303122003-2110223323213310-3023132231200123-3213003330312322-0230032131103003"></a>

<a id="canonical-3022011311033022-0010222020111111-1210330212212130-0300022230321021-1332122202032202-1230222332321313-1212200312110233-2303031011311001"></a>

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-3332330200312223-3023022231023220-0221300011102220-0030301313002013-3022031122012310-0321021013011321-1120223121302333-3201032223210133"></a>

<a id="canonical-1212323300003012-0023303323331123-0200233113033101-3021301332200322-1111133310012232-0232122023231320-0123001122132000-3221331300211001"></a>

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

<a id="canonical-1113110103002001-2200212320222323-2122120023300322-2333320121311313-0323121200121312-0111233201222031-3201301002301133-0200220001100111"></a>

<a id="canonical-1032323302322101-1120113230020033-0122232031211000-0232111211313203-1312333131223231-0220232213111133-3230311023123203-0312200012333202"></a>

#### `id` property

Type: `"string"`. Computed.

Unique identifier for the resource.

- [irules](resources--bigip_http_proxy--reference--group-001.md#canonical-1120320133323032-2322013030113130-0032022100212120-1002211211233300-1000313031332132-3203200330320300-1100023210101120-0131020312332202): complete subsection reference.

<a id="canonical-2221101312132233-1222030332301331-3303031331203320-3020210001233222-1022031333013231-0333031322123322-3001210133130330-3121312232103312"></a>

<a id="canonical-3332023102013000-3313002301101213-3112110121110311-3213321033331132-2202303133201132-1321120113300030-0002232320310030-1123130103233300"></a>

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

- [lb_algorithm](resources--bigip_http_proxy--reference--group-001.md#canonical-0010233021331310-1100023120223111-0232222211022121-1032211023101301-3003123030113013-3230220310130310-0221133011233320-1132331211111231): complete subsection reference.

<a id="canonical-2313110223120123-2330023103302100-0010031021311123-2221031311321002-1221132310231300-3012113332231120-2122203230133202-1130102001002231"></a>

<a id="canonical-0231330120112123-2032330110302221-3211233232230102-1123202103213320-1322320101002103-3330031121031121-1000313003120220-3012023232330101"></a>

#### `name` property

Type: `"string"`. Required.

Name of the BIG-IP HTTP Proxy. Must be unique within the namespace.

Additional upstream details:

This is the name of configuration object. It has to be unique within the namespace. It can only be
specified during create API and cannot be changed during replace API. The value of name has to
follow DNS-1035 format.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.String{
  validators.NameValidator(),
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
    "formatDescription": "DNS-1035 label: must start with a lowercase letter, may contain lowercase alphanumeric and hyphens, must end with alphanumeric",
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

<a id="canonical-3220103110212201-2121223231200113-0202032322320321-2323002000212122-3233303321020013-3111222102102333-1223203001323323-1221101222203301"></a>

<a id="canonical-2113201313030133-3012121032012310-2313300001102230-0101012333023132-2300011212301203-1232331102111222-2132213211230213-2112221103021010"></a>

#### `namespace` property

Type: `"string"`. Required.

Namespace where the BIG-IP HTTP Proxy is created.

Additional upstream details:

This defines the workspace within which each the configuration object is to be created. Must be a
DNS\_LABEL format. For a namespace object itself, namespace value will be ""

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.String{
  validators.NamespaceValidator(),
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

- [origin_pools](resources--bigip_http_proxy--reference--group-001.md#canonical-3130302100213032-1133033213001002-2331132101322013-0203331122233331-3021113033330230-2313122333001332-0303103133232110-2002132020103202): complete subsection reference.

- [proxy_advertisement](resources--bigip_http_proxy--reference--group-002.md#canonical-0222202123223213-1021332332331003-3333013231022023-1200013112100010-3021102122231210-1210031321121132-3303131312332322-2002012022311122): complete subsection reference.

- [proxy_config](resources--bigip_http_proxy--reference--group-003.md#canonical-3003013220031132-0222103102123201-3133022011003320-3300031222211002-0032213213000000-1001011223312300-1133211303323330-1223230120123033): complete subsection reference.

- [timeouts](resources--bigip_http_proxy--reference--group-004.md#canonical-1231333103033012-2112201220312122-2222023320130000-2122221200023233-1122012332200101-1313233133310311-3230012111033123-0101131100232102): complete subsection reference.

<a id="canonical-2233020022001102-3212012112010201-3212100211223201-3002011120313133-1132032213013111-0223022203201223-0301310111132320-3020220223102022"></a>

### All schema paths for `xcsh_bigip_http_proxy`

Each exact path has one authoritative reference destination. Collection element indices are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `advanced_profile` | [advanced_profile](resources--bigip_http_proxy--reference--group-001.md#canonical-2202101110120110-1303332222001222-0101031233131220-1321323001223202-0310111203221022-0121201102003320-3122213002331303-3323322131233320) |
| `advanced_profile.disable_spec` | [advanced_profile.disable_spec](resources--bigip_http_proxy--reference--group-001.md#canonical-3022212011131102-1301213222303303-1233022013233131-1232111020211133-0022133101010300-1003312110023132-2102010321120310-3330031000333121) |
| `advanced_profile.enable_default_profile` | [advanced_profile.enable_default_profile](resources--bigip_http_proxy--reference--group-001.md#canonical-0201313032000320-1323113231232232-2311320230113231-1201223201010012-2330122302202222-1320302221121132-1131200120011200-2100303120132231) |
| `annotations` | [annotations](resources--bigip_http_proxy--reference--group-001.md#canonical-2313003200321320-1322131322312032-1123031210113332-0230222303330322-3110303033332210-0102330003211132-3213223121223322-3221302223313313) |
| `ddos_profile` | [ddos_profile](resources--bigip_http_proxy--reference--group-001.md#canonical-3121300011110230-1223302133011302-0222020300333120-3123321130033230-1112032013102131-3000000213222331-3103220322221300-1131001113201213) |
| `ddos_profile.disable_ddos_mitigation` | [ddos_profile.disable_ddos_mitigation](resources--bigip_http_proxy--reference--group-001.md#canonical-0002230032332111-2200031133321211-3102201011032320-0210302233031102-2003223021300203-0101232321031110-0120123332100222-0301110303322102) |
| `ddos_profile.enable_ddos_mitigation` | [ddos_profile.enable_ddos_mitigation](resources--bigip_http_proxy--reference--group-001.md#canonical-2122033133230010-1301132100003330-2321002022202200-3320221013033020-1001003213001331-1310003033103132-1030301111020113-0033020103032000) |
| `description` | [description](resources--bigip_http_proxy--reference--group-001.md#canonical-3223112011313112-2300120332112033-2323203201323103-0110010303122003-2110223323213310-3023132231200123-3213003330312322-0230032131103003) |
| `disable` | [disable](resources--bigip_http_proxy--reference--group-001.md#canonical-3332330200312223-3023022231023220-0221300011102220-0030301313002013-3022031122012310-0321021013011321-1120223121302333-3201032223210133) |
| `id` | [ID](resources--bigip_http_proxy--reference--group-001.md#canonical-1113110103002001-2200212320222323-2122120023300322-2333320121311313-0323121200121312-0111233201222031-3201301002301133-0200220001100111) |
| `irules` | [irules](resources--bigip_http_proxy--reference--group-001.md#canonical-0300010123030301-1030210311032130-3233322003322132-3033321202133020-2323200020113323-1002201233120210-2022122302231312-2333010232311203) |
| `irules.irules` | [irules.irules](resources--bigip_http_proxy--reference--group-001.md#canonical-2222000011323303-1231101132120330-3032101203102231-2201033030112020-1113322132330120-1012100132313130-0223301321303332-1332333123300032) |
| `irules.irules.name` | [irules.irules.name](resources--bigip_http_proxy--reference--group-001.md#canonical-2031021110012102-1020100013102311-3123211213202333-3133033312013202-2121112013020033-0123020310102313-3211102123301130-2030020013111323) |
| `irules.irules.namespace` | [irules.irules.namespace](resources--bigip_http_proxy--reference--group-001.md#canonical-2313121021230311-0101123000031022-3020323033202133-0010033330111223-2332033203133003-0100231322220232-3031223311131312-1233301302232033) |
| `irules.irules.tenant` | [irules.irules.tenant](resources--bigip_http_proxy--reference--group-001.md#canonical-1312211003113032-2233030002223332-3320331113102300-3320211322131000-0210023313023230-1312033330131100-0123210312230330-3023320000133030) |
| `labels` | [labels](resources--bigip_http_proxy--reference--group-001.md#canonical-2221101312132233-1222030332301331-3303031331203320-3020210001233222-1022031333013231-0333031322123322-3001210133130330-3121312232103312) |
| `lb_algorithm` | [lb_algorithm](resources--bigip_http_proxy--reference--group-001.md#canonical-1203213130303300-2102000300231232-2201011123000301-1210101120003333-3102303320133200-3211331113202003-3113012130100202-3110132101312213) |
| `lb_algorithm.round_robin` | [lb_algorithm.round_robin](resources--bigip_http_proxy--reference--group-001.md#canonical-1020321112021311-3233221200032311-0011011203111303-2223303130211232-2221031111110220-2210300121322212-1200032000022223-3121032112121031) |
| `name` | [name](resources--bigip_http_proxy--reference--group-001.md#canonical-2313110223120123-2330023103302100-0010031021311123-2221031311321002-1221132310231300-3012113332231120-2122203230133202-1130102001002231) |
| `namespace` | [namespace](resources--bigip_http_proxy--reference--group-001.md#canonical-3220103110212201-2121223231200113-0202032322320321-2323002000212122-3233303321020013-3111222102102333-1223203001323323-1221101222203301) |
| `origin_pools` | [origin_pools](resources--bigip_http_proxy--reference--group-001.md#canonical-2333331002003021-3130223021000003-2301230130332012-3101331020300311-2130303003122212-0333033333231300-2133021212030333-0221003231132311) |
| `origin_pools.pools` | [origin_pools.pools](resources--bigip_http_proxy--reference--group-001.md#canonical-3323100232203202-2203032231302200-2233113210023011-1222212333122323-2300100100323012-2013033030123201-3231030123231013-2100302320110102) |
| `origin_pools.pools.name` | [origin_pools.pools.name](resources--bigip_http_proxy--reference--group-001.md#canonical-2311203003130200-1221121201030232-0021202021213322-0223201102200320-3211101212202221-1210000123303323-3202010310133002-2022301231331213) |
| `origin_pools.pools.origin_servers` | [origin_pools.pools.origin_servers](resources--bigip_http_proxy--reference--group-001.md#canonical-3211323131311301-0133100120330120-0131202011131231-1310232023120113-0020232233021020-0320102303231100-3232230122332012-2030010130011220) |
| `origin_pools.pools.origin_servers.automatic_port` | [origin_pools.pools.origin_servers.automatic_port](resources--bigip_http_proxy--reference--group-001.md#canonical-2320233022123330-1122032100312112-1303112231020000-3233012111110303-0020100313232031-3130210321233101-3302122320232132-3111023213102333) |
| `origin_pools.pools.origin_servers.health_checks` | [origin_pools.pools.origin_servers.health_checks](resources--bigip_http_proxy--reference--group-001.md#canonical-1331103210122032-1222112002002132-1010202022302101-0102201221133333-3332310221213320-2032310013300102-0003112303121020-3130211302312231) |
| `origin_pools.pools.origin_servers.health_checks.health_check` | [origin_pools.pools.origin_servers.health_checks.health_check](resources--bigip_http_proxy--reference--group-001.md#canonical-2133223333122032-2011001010233113-3210103130230021-0310202313132323-0013323110212011-1033223333222302-1201000101002132-3102232032310210) |
| `origin_pools.pools.origin_servers.health_checks.health_check.icmp_health_check` | [origin_pools.pools.origin_servers.health_checks.health_check.icmp_health_check](resources--bigip_http_proxy--reference--group-001.md#canonical-2101212012010021-1201233100032322-0000103131120333-2033212232112011-0332121122201032-1300033333312030-2202301331122333-0223322113030102) |
| `origin_pools.pools.origin_servers.health_checks.health_check.tcp_health_check` | [origin_pools.pools.origin_servers.health_checks.health_check.tcp_health_check](resources--bigip_http_proxy--reference--group-001.md#canonical-3101022020020302-0133012122013101-0300331300221300-0131333120113132-2003023310223011-3102032200110003-1222112322032311-3131210333333131) |
| `origin_pools.pools.origin_servers.health_checks.health_check.tcp_health_check.expected_response` | [origin_pools.pools.origin_servers.health_checks.health_check.tcp_health_check.expected_response](resources--bigip_http_proxy--reference--group-001.md#canonical-1122100001101233-1320333212013101-0320120233333010-3321120033021023-3112330322332012-3022330032233120-0211010020223131-2230103033023201) |
| `origin_pools.pools.origin_servers.health_checks.health_check.tcp_health_check.send_payload` | [origin_pools.pools.origin_servers.health_checks.health_check.tcp_health_check.send_payload](resources--bigip_http_proxy--reference--group-001.md#canonical-3003212330320123-0111322003203032-2103103333211012-3300321003003303-2111222120331031-0000303010022232-1332230220110101-3031021232112021) |
| `origin_pools.pools.origin_servers.health_checks.healthy_threshold` | [origin_pools.pools.origin_servers.health_checks.healthy_threshold](resources--bigip_http_proxy--reference--group-001.md#canonical-2030311222030320-2233131102220000-2312113031303002-3231023133311200-0010000202111012-3131231302321333-1313003321030220-3223013312310211) |
| `origin_pools.pools.origin_servers.health_checks.interval` | [origin_pools.pools.origin_servers.health_checks.interval](resources--bigip_http_proxy--reference--group-001.md#canonical-1020230310301001-1232001301033103-2011320133202203-1212212011200310-2200203211020030-3313020122121230-0103120213000230-2130120102211102) |
| `origin_pools.pools.origin_servers.health_checks.timeout` | [origin_pools.pools.origin_servers.health_checks.timeout](resources--bigip_http_proxy--reference--group-001.md#canonical-2033230321000020-3032322202321132-0111000301020320-2123013031331010-1110212311123110-3032122220303103-1320200210312230-2010222011120101) |
| `origin_pools.pools.origin_servers.health_checks.unhealthy_threshold` | [origin_pools.pools.origin_servers.health_checks.unhealthy_threshold](resources--bigip_http_proxy--reference--group-001.md#canonical-0301211311133301-2211031123203031-0233011331213332-2003200221221111-2022312000011011-0013302021213103-0103233302113203-1321300111103122) |
| `origin_pools.pools.origin_servers.lb_port` | [origin_pools.pools.origin_servers.lb_port](resources--bigip_http_proxy--reference--group-001.md#canonical-3202233223233310-1201311120130333-1112100312003213-2102033002011312-0202001200023321-2111323330230120-1022021102001001-0102130001032220) |
| `origin_pools.pools.origin_servers.origin_servers` | [origin_pools.pools.origin_servers.origin_servers](resources--bigip_http_proxy--reference--group-001.md#canonical-0011000100323203-3122231030111012-0303222000311211-3321333133321310-2121010221210121-0033101132030101-0213323113213313-1113220131123200) |
| `origin_pools.pools.origin_servers.origin_servers.k8s_service` | [origin_pools.pools.origin_servers.origin_servers.k8s_service](resources--bigip_http_proxy--reference--group-001.md#canonical-0003000021001331-3303012123010331-3012201100312110-2312003113013033-0231133022022023-2233133132100330-3201322321302220-2033002231323330) |
| `origin_pools.pools.origin_servers.origin_servers.k8s_service.inside_network` | [origin_pools.pools.origin_servers.origin_servers.k8s_service.inside_network](resources--bigip_http_proxy--reference--group-001.md#canonical-2002213121110021-1303111211120211-1223013023211120-2030021223203320-3101103302302131-3122221023311200-3111331202202003-1311203001233000) |
| `origin_pools.pools.origin_servers.origin_servers.k8s_service.outside_network` | [origin_pools.pools.origin_servers.origin_servers.k8s_service.outside_network](resources--bigip_http_proxy--reference--group-001.md#canonical-2113320321333101-1300030113323322-3220012031320010-1213111002021221-0102012203023133-0333012103302102-0310032001311232-3230322021020011) |
| `origin_pools.pools.origin_servers.origin_servers.k8s_service.protocol` | [origin_pools.pools.origin_servers.origin_servers.k8s_service.protocol](resources--bigip_http_proxy--reference--group-001.md#canonical-1311112220202330-0001203101011022-1232022322123302-2330132112023231-2021210011110031-0220103113100210-0022322100030033-0202013213000221) |
| `origin_pools.pools.origin_servers.origin_servers.k8s_service.service_name` | [origin_pools.pools.origin_servers.origin_servers.k8s_service.service_name](resources--bigip_http_proxy--reference--group-001.md#canonical-2111201023220313-2111231110200032-3320333300112123-0203311311301130-3330002311221323-2232313212012301-1322230112203131-3002022322230110) |
| `origin_pools.pools.origin_servers.origin_servers.k8s_service.site_locator` | [origin_pools.pools.origin_servers.origin_servers.k8s_service.site_locator](resources--bigip_http_proxy--reference--group-001.md#canonical-2223011211012000-2322312131103323-1222033230230203-2330032121301233-1311031202121131-0312232021330222-0222012331130132-3323303131200101) |
| `origin_pools.pools.origin_servers.origin_servers.k8s_service.site_locator.site` | [origin_pools.pools.origin_servers.origin_servers.k8s_service.site_locator.site](resources--bigip_http_proxy--reference--group-001.md#canonical-2221311223231120-2102313010233223-0303131021022032-0221030001022012-3013302312213122-1302002131311000-1212310123030322-2203000202021023) |
| `origin_pools.pools.origin_servers.origin_servers.k8s_service.site_locator.site.name` | [origin_pools.pools.origin_servers.origin_servers.k8s_service.site_locator.site.name](resources--bigip_http_proxy--reference--group-001.md#canonical-0023212223030121-2121033320111111-1020232321210030-3302201231221030-0011323010121230-2100112130221130-2113303101300310-3300220302322003) |
| `origin_pools.pools.origin_servers.origin_servers.k8s_service.site_locator.site.namespace` | [origin_pools.pools.origin_servers.origin_servers.k8s_service.site_locator.site.namespace](resources--bigip_http_proxy--reference--group-001.md#canonical-0101113020333213-2223222320111301-2012033213123102-0332201113030223-0003210102103111-0322130211302220-2313311212133331-3133112303211210) |
| `origin_pools.pools.origin_servers.origin_servers.k8s_service.site_locator.site.tenant` | [origin_pools.pools.origin_servers.origin_servers.k8s_service.site_locator.site.tenant](resources--bigip_http_proxy--reference--group-001.md#canonical-3310111322233202-0013031020111122-3011122230303030-2320102023300212-3331221300201101-1031113132122302-1330110122021130-3032321220100321) |
| `origin_pools.pools.origin_servers.origin_servers.k8s_service.site_locator.virtual_site` | [origin_pools.pools.origin_servers.origin_servers.k8s_service.site_locator.virtual_site](resources--bigip_http_proxy--reference--group-001.md#canonical-3331313230220333-1330332111221023-2032211002020313-2213003310003333-0021133331231100-1211121001220210-2222020210232331-1330112331332323) |
| `origin_pools.pools.origin_servers.origin_servers.k8s_service.site_locator.virtual_site.name` | [origin_pools.pools.origin_servers.origin_servers.k8s_service.site_locator.virtual_site.name](resources--bigip_http_proxy--reference--group-001.md#canonical-0223002232003220-2210132010112113-0331021110020110-3002203333312200-3333222311022113-2232320120300303-2233122333223103-1220331323020100) |
| `origin_pools.pools.origin_servers.origin_servers.k8s_service.site_locator.virtual_site.namespace` | [origin_pools.pools.origin_servers.origin_servers.k8s_service.site_locator.virtual_site.namespace](resources--bigip_http_proxy--reference--group-001.md#canonical-3122222102303000-0133132113001303-3000232332020011-3112011012121002-2232003330133030-2321213302201020-1321123321002033-2111330030031312) |
| `origin_pools.pools.origin_servers.origin_servers.k8s_service.site_locator.virtual_site.tenant` | [origin_pools.pools.origin_servers.origin_servers.k8s_service.site_locator.virtual_site.tenant](resources--bigip_http_proxy--reference--group-001.md#canonical-0021231230231113-1300211321002230-0102000110202102-0010311320203013-3111220002022030-2302212302301112-0012223321113332-3112301123120011) |
| `origin_pools.pools.origin_servers.origin_servers.k8s_service.snat_pool` | [origin_pools.pools.origin_servers.origin_servers.k8s_service.snat_pool](resources--bigip_http_proxy--reference--group-001.md#canonical-1210031012130200-3003222231330333-0133200303022131-1223223112102113-3210213223112122-1331002330012323-3032111213321023-3022112132102131) |
| `origin_pools.pools.origin_servers.origin_servers.k8s_service.snat_pool.no_snat_pool` | [origin_pools.pools.origin_servers.origin_servers.k8s_service.snat_pool.no_snat_pool](resources--bigip_http_proxy--reference--group-001.md#canonical-2123321133201330-1101022110010022-2212000130131300-3133310211302131-0030133321202301-2022230213131132-0120023221132230-3102120020000311) |
| `origin_pools.pools.origin_servers.origin_servers.k8s_service.snat_pool.snat_pool` | [origin_pools.pools.origin_servers.origin_servers.k8s_service.snat_pool.snat_pool](resources--bigip_http_proxy--reference--group-002.md#canonical-3211300102000202-3203101313113331-0331003310300202-3303003031202201-2003013322303020-2233320301203211-1122013313021221-2101220030033123) |
| `origin_pools.pools.origin_servers.origin_servers.k8s_service.snat_pool.snat_pool.prefixes` | [origin_pools.pools.origin_servers.origin_servers.k8s_service.snat_pool.snat_pool.prefixes](resources--bigip_http_proxy--reference--group-002.md#canonical-3302231032122000-1312303303312013-0101003333130022-2030011323001031-2022302202032303-2323120111000321-0323332302211130-1223331130300021) |
| `origin_pools.pools.origin_servers.origin_servers.k8s_service.vk8s_networks` | [origin_pools.pools.origin_servers.origin_servers.k8s_service.vk8s_networks](resources--bigip_http_proxy--reference--group-002.md#canonical-0211133031012200-2011221102322000-1112030331321300-3123012032022312-3011231122213110-2012121212030310-1210231103212022-2103221310302031) |
| `origin_pools.pools.origin_servers.origin_servers.private_ip` | [origin_pools.pools.origin_servers.origin_servers.private_ip](resources--bigip_http_proxy--reference--group-002.md#canonical-2321223011000313-2312122221131022-3113333013312112-0230003332100121-2120322130022220-0332332310100321-0133011133031000-2222010013223133) |
| `origin_pools.pools.origin_servers.origin_servers.private_ip.inside_network` | [origin_pools.pools.origin_servers.origin_servers.private_ip.inside_network](resources--bigip_http_proxy--reference--group-002.md#canonical-0233101222100231-3302112213020231-3322003231333122-3130120112112211-2230323102202120-3213030221121330-2131211200131013-0002023100010101) |
| `origin_pools.pools.origin_servers.origin_servers.private_ip.ip` | [origin_pools.pools.origin_servers.origin_servers.private_ip.ip](resources--bigip_http_proxy--reference--group-002.md#canonical-1022233231121011-3310131031220322-1311222313111030-0001211303112323-1101023111231102-0112003203110311-1130131202330210-2022203221012200) |
| `origin_pools.pools.origin_servers.origin_servers.private_ip.outside_network` | [origin_pools.pools.origin_servers.origin_servers.private_ip.outside_network](resources--bigip_http_proxy--reference--group-002.md#canonical-2102223102132210-2203301302312120-2212022310023003-2312013330121102-2312123012322021-3103312132133000-1211223113103113-0030101333222020) |
| `origin_pools.pools.origin_servers.origin_servers.private_ip.segment` | [origin_pools.pools.origin_servers.origin_servers.private_ip.segment](resources--bigip_http_proxy--reference--group-002.md#canonical-2331312210111131-2211020012321111-0003122010132102-3102030130113302-3301000131030110-1021133223220031-1112101220302020-0002021300220101) |
| `origin_pools.pools.origin_servers.origin_servers.private_ip.segment.name` | [origin_pools.pools.origin_servers.origin_servers.private_ip.segment.name](resources--bigip_http_proxy--reference--group-002.md#canonical-0120022000012333-1102013223103310-0233133033123210-2321201113332213-0012320113030213-0221020111122232-0112320323012311-2123032323323113) |
| `origin_pools.pools.origin_servers.origin_servers.private_ip.segment.namespace` | [origin_pools.pools.origin_servers.origin_servers.private_ip.segment.namespace](resources--bigip_http_proxy--reference--group-002.md#canonical-3123003230221133-3313330330000212-3022210002302111-3101100210330222-1323031213202130-3201220102030122-1031021323123122-1303303333031130) |
| `origin_pools.pools.origin_servers.origin_servers.private_ip.segment.tenant` | [origin_pools.pools.origin_servers.origin_servers.private_ip.segment.tenant](resources--bigip_http_proxy--reference--group-002.md#canonical-3000130331013221-0230003332023333-2311121333302030-0010302012200233-3102132302313023-0222303002312201-0132003203212321-0011131231312330) |
| `origin_pools.pools.origin_servers.origin_servers.private_ip.site_locator` | [origin_pools.pools.origin_servers.origin_servers.private_ip.site_locator](resources--bigip_http_proxy--reference--group-002.md#canonical-0233221201203020-0011210022002003-3322230223201030-3133102303102002-2312110201100122-1010223331300130-1131230323330321-1012213232120130) |
| `origin_pools.pools.origin_servers.origin_servers.private_ip.site_locator.site` | [origin_pools.pools.origin_servers.origin_servers.private_ip.site_locator.site](resources--bigip_http_proxy--reference--group-002.md#canonical-1232121222333030-1231220133301011-2002001330312132-3030101003132221-2310333303012230-0322233132212313-0213222210033212-3322311200231121) |
| `origin_pools.pools.origin_servers.origin_servers.private_ip.site_locator.site.name` | [origin_pools.pools.origin_servers.origin_servers.private_ip.site_locator.site.name](resources--bigip_http_proxy--reference--group-002.md#canonical-3122330301332313-3320001113231021-1222231131023330-0201312233003003-3010123320322232-2113303033001022-2000122302003111-0321031133331001) |
| `origin_pools.pools.origin_servers.origin_servers.private_ip.site_locator.site.namespace` | [origin_pools.pools.origin_servers.origin_servers.private_ip.site_locator.site.namespace](resources--bigip_http_proxy--reference--group-002.md#canonical-3212303023220122-1312212030012230-2021301001232021-3020210300201001-0022211333101013-2220113013333212-3112320010101031-3131220330302233) |
| `origin_pools.pools.origin_servers.origin_servers.private_ip.site_locator.site.tenant` | [origin_pools.pools.origin_servers.origin_servers.private_ip.site_locator.site.tenant](resources--bigip_http_proxy--reference--group-002.md#canonical-3201330000011032-2013003301013110-0202221210132320-1301203330202221-1202003030230210-2322323111002311-2003012103303103-1130032130031101) |
| `origin_pools.pools.origin_servers.origin_servers.private_ip.site_locator.virtual_site` | [origin_pools.pools.origin_servers.origin_servers.private_ip.site_locator.virtual_site](resources--bigip_http_proxy--reference--group-002.md#canonical-1330121132112212-1113012313001011-3011201133023223-1130101231133012-0200132203203022-2330010310223301-0022001133112230-2021210201331122) |
| `origin_pools.pools.origin_servers.origin_servers.private_ip.site_locator.virtual_site.name` | [origin_pools.pools.origin_servers.origin_servers.private_ip.site_locator.virtual_site.name](resources--bigip_http_proxy--reference--group-002.md#canonical-1230332201313020-2330100021321133-1331303310100330-3330103231201130-3322233230120032-1330102100303023-1032323031033013-1313213112023122) |
| `origin_pools.pools.origin_servers.origin_servers.private_ip.site_locator.virtual_site.namespace` | [origin_pools.pools.origin_servers.origin_servers.private_ip.site_locator.virtual_site.namespace](resources--bigip_http_proxy--reference--group-002.md#canonical-1112002300210232-1101303000123323-3310110211021222-1323220103103102-0122122012122033-2213312212313011-1211213312333312-2203233301023133) |
| `origin_pools.pools.origin_servers.origin_servers.private_ip.site_locator.virtual_site.tenant` | [origin_pools.pools.origin_servers.origin_servers.private_ip.site_locator.virtual_site.tenant](resources--bigip_http_proxy--reference--group-002.md#canonical-0003021032103112-0222310101301202-1231032020010131-3010002101102100-3333213303333222-2301030011132130-1232002302132211-1122032321002102) |
| `origin_pools.pools.origin_servers.origin_servers.private_ip.snat_pool` | [origin_pools.pools.origin_servers.origin_servers.private_ip.snat_pool](resources--bigip_http_proxy--reference--group-002.md#canonical-2032001111021122-2102131012223222-3102113022022213-1033013131323312-1202313102131023-1232122121101233-2033002200013133-2323203313312002) |
| `origin_pools.pools.origin_servers.origin_servers.private_ip.snat_pool.no_snat_pool` | [origin_pools.pools.origin_servers.origin_servers.private_ip.snat_pool.no_snat_pool](resources--bigip_http_proxy--reference--group-002.md#canonical-2103102333211113-3103223330031332-3123121121221230-2030330200313200-3332330112002033-0231111123111111-2220010232322121-1312101023333033) |
| `origin_pools.pools.origin_servers.origin_servers.private_ip.snat_pool.snat_pool` | [origin_pools.pools.origin_servers.origin_servers.private_ip.snat_pool.snat_pool](resources--bigip_http_proxy--reference--group-002.md#canonical-1032330311300003-3313320310030103-2123221302330333-2031303113322031-0232300010133111-2121223233003002-2132103133022013-3021331100220011) |
| `origin_pools.pools.origin_servers.origin_servers.private_ip.snat_pool.snat_pool.prefixes` | [origin_pools.pools.origin_servers.origin_servers.private_ip.snat_pool.snat_pool.prefixes](resources--bigip_http_proxy--reference--group-002.md#canonical-1123111031032121-3320323233310221-0221311312301131-2122302130030111-1310121313331113-2001103020313120-2322110030023002-1303330231132220) |
| `origin_pools.pools.origin_servers.origin_servers.public_ip` | [origin_pools.pools.origin_servers.origin_servers.public_ip](resources--bigip_http_proxy--reference--group-002.md#canonical-2010223030113203-0233211001220121-0020103030122012-0322021030333003-0222223232011221-1032011103210000-2033333303303230-1032132220010230) |
| `origin_pools.pools.origin_servers.origin_servers.public_ip.ip` | [origin_pools.pools.origin_servers.origin_servers.public_ip.ip](resources--bigip_http_proxy--reference--group-002.md#canonical-1322132030010121-3333233020031300-3012030121233123-2113012311220203-0113131030001030-0212133011301220-2303021123210322-0232133232211212) |
| `origin_pools.pools.origin_servers.origin_servers.public_name` | [origin_pools.pools.origin_servers.origin_servers.public_name](resources--bigip_http_proxy--reference--group-002.md#canonical-3302301122002013-1230313322333200-3202322111200011-1220333212323001-3101001320231000-3121121102003322-1303101233120102-2322002123030033) |
| `origin_pools.pools.origin_servers.origin_servers.public_name.dns_name` | [origin_pools.pools.origin_servers.origin_servers.public_name.dns_name](resources--bigip_http_proxy--reference--group-002.md#canonical-2311332013331031-0313230230033031-1330302213112331-1022123032312131-0230230110331312-0023020021322303-2102130222013013-2213102020213320) |
| `origin_pools.pools.origin_servers.origin_servers.public_name.refresh_interval` | [origin_pools.pools.origin_servers.origin_servers.public_name.refresh_interval](resources--bigip_http_proxy--reference--group-002.md#canonical-0123321023020110-1321322121101332-0223330211031023-2122132021301320-2021211312200021-1032021002012120-3330000212131300-1221331120303333) |
| `origin_pools.pools.origin_servers.port` | [origin_pools.pools.origin_servers.port](resources--bigip_http_proxy--reference--group-001.md#canonical-2202121131021302-3301332301103133-1121102033022033-2123212002323211-1233322000022030-3101300213112020-2332222100311212-3213313100222210) |
| `origin_pools.pools.priority` | [origin_pools.pools.priority](resources--bigip_http_proxy--reference--group-001.md#canonical-2033230213221231-3331312031213203-0110110000100230-1112333311223123-1311000211323100-3333203032112322-3331311121323310-1300330122020302) |
| `origin_pools.pools.weight` | [origin_pools.pools.weight](resources--bigip_http_proxy--reference--group-001.md#canonical-1111202133231300-1221011120122111-1321032030113101-2232101201021311-3000211123312201-0021032203200113-3310003333301132-0113110103130103) |
| `proxy_advertisement` | [proxy_advertisement](resources--bigip_http_proxy--reference--group-002.md#canonical-3132000202003233-2200103133232231-3123032303033200-2022013023131200-1131211311013010-1311220332231213-1130023102200031-2310220333022030) |
| `proxy_advertisement.advertise_custom` | [proxy_advertisement.advertise_custom](resources--bigip_http_proxy--reference--group-002.md#canonical-1311111120023312-0312200122211022-2203203010020131-3030021221200200-1321200313112302-2131221110202100-0000022232000331-3300010301300011) |
| `proxy_advertisement.advertise_custom.advertise_where` | [proxy_advertisement.advertise_custom.advertise_where](resources--bigip_http_proxy--reference--group-002.md#canonical-2111210223102301-3133011231011332-1331231303230113-0322321020123313-2020103311011011-3010200032113023-1323301033000100-1120000033033231) |
| `proxy_advertisement.advertise_custom.advertise_where.advertise_dualstack_on_public` | [proxy_advertisement.advertise_custom.advertise_where.advertise_dualstack_on_public](resources--bigip_http_proxy--reference--group-002.md#canonical-0120021301122100-2030100011101123-0131130333003222-3232101122223312-0100131303233011-3012000102312103-0103001101231010-1132100013233010) |
| `proxy_advertisement.advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip` | [proxy_advertisement.advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip](resources--bigip_http_proxy--reference--group-002.md#canonical-1131123230220011-0333210220032210-3213321011220300-3100303231332312-2100210230201020-1021131232002003-3133131120010320-2132212312210021) |
| `proxy_advertisement.advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip.name` | [proxy_advertisement.advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip.name](resources--bigip_http_proxy--reference--group-002.md#canonical-2113310020200232-0032200112320230-3223332111123311-0333112202030301-2300120312101003-1130310003121322-0123102210102103-1303010303102321) |
| `proxy_advertisement.advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip.namespace` | [proxy_advertisement.advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip.namespace](resources--bigip_http_proxy--reference--group-002.md#canonical-1123121121113111-0103233131303213-0000231220020230-3320123232212101-3200003211330332-3113331102121032-3310130132021333-3100100333233132) |
| `proxy_advertisement.advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip.tenant` | [proxy_advertisement.advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip.tenant](resources--bigip_http_proxy--reference--group-002.md#canonical-0012221033023300-1200302003110303-0000312233211030-1123033332132232-3331200032322120-2001313020012233-0120221001132023-3330112010003200) |
| `proxy_advertisement.advertise_custom.advertise_where.advertise_on_public` | [proxy_advertisement.advertise_custom.advertise_where.advertise_on_public](resources--bigip_http_proxy--reference--group-002.md#canonical-0101302102033232-3222112100301323-3212211020000210-2320032230123112-0101200230332111-2230022301030320-3311101002112211-0302032001313333) |
| `proxy_advertisement.advertise_custom.advertise_where.advertise_on_public.public_ip` | [proxy_advertisement.advertise_custom.advertise_where.advertise_on_public.public_ip](resources--bigip_http_proxy--reference--group-002.md#canonical-3330320333320032-0301001100232020-1023122320222302-2312232013321022-0032122211131102-0222303330101003-1333020302210020-1302211101120101) |
| `proxy_advertisement.advertise_custom.advertise_where.advertise_on_public.public_ip.name` | [proxy_advertisement.advertise_custom.advertise_where.advertise_on_public.public_ip.name](resources--bigip_http_proxy--reference--group-002.md#canonical-3211301212122000-2211303110202200-0122322222103032-0333332200322031-0320223103233003-2313012131222211-1320212131303313-1330110322022032) |
| `proxy_advertisement.advertise_custom.advertise_where.advertise_on_public.public_ip.namespace` | [proxy_advertisement.advertise_custom.advertise_where.advertise_on_public.public_ip.namespace](resources--bigip_http_proxy--reference--group-002.md#canonical-3111103001220101-3310110132323100-3333323301301231-3220001012303132-3133010111220230-3333302010003100-0333201021120111-3100111020330311) |
| `proxy_advertisement.advertise_custom.advertise_where.advertise_on_public.public_ip.tenant` | [proxy_advertisement.advertise_custom.advertise_where.advertise_on_public.public_ip.tenant](resources--bigip_http_proxy--reference--group-002.md#canonical-2300312032231221-2333010112223233-2131132221320322-1012123123121032-3332013031013331-1030313203121021-2133332012233211-2200303312312313) |
| `proxy_advertisement.advertise_custom.advertise_where.advertise_v6_on_public` | [proxy_advertisement.advertise_custom.advertise_where.advertise_v6_on_public](resources--bigip_http_proxy--reference--group-002.md#canonical-0321330023321310-2021100133302101-3031230202132103-1000320203202112-2013020131221203-3003021213321223-0331113302123033-3001212033111001) |
| `proxy_advertisement.advertise_custom.advertise_where.advertise_v6_on_public.public_ip` | [proxy_advertisement.advertise_custom.advertise_where.advertise_v6_on_public.public_ip](resources--bigip_http_proxy--reference--group-002.md#canonical-2232313000130002-3103222300301130-3031101201230022-3210113031221220-0312121000311201-0101320112332001-2223201221230010-3330201120321002) |
| `proxy_advertisement.advertise_custom.advertise_where.advertise_v6_on_public.public_ip.name` | [proxy_advertisement.advertise_custom.advertise_where.advertise_v6_on_public.public_ip.name](resources--bigip_http_proxy--reference--group-002.md#canonical-1030122130221222-1123231212033213-2121222233130301-1011002111310332-2312030212301303-0123221330333123-2111100330100203-3123212120310012) |
| `proxy_advertisement.advertise_custom.advertise_where.advertise_v6_on_public.public_ip.namespace` | [proxy_advertisement.advertise_custom.advertise_where.advertise_v6_on_public.public_ip.namespace](resources--bigip_http_proxy--reference--group-002.md#canonical-2122232012111211-2201002230100210-3113320002001012-0321210102121133-3112120102121300-3310313010111210-2322323002212200-3301201323301023) |
| `proxy_advertisement.advertise_custom.advertise_where.advertise_v6_on_public.public_ip.tenant` | [proxy_advertisement.advertise_custom.advertise_where.advertise_v6_on_public.public_ip.tenant](resources--bigip_http_proxy--reference--group-002.md#canonical-1031320200012330-0213210102302123-0113023211033113-0001330201032332-0222130312113203-0120003300202233-3131201213132321-0103331312223122) |
| `proxy_advertisement.advertise_custom.advertise_where.port` | [proxy_advertisement.advertise_custom.advertise_where.port](resources--bigip_http_proxy--reference--group-002.md#canonical-0301112112213301-1022222002130222-3132203032012111-3323021013100330-3331003120021203-3323231223112033-1332211011201232-0120223313112022) |
| `proxy_advertisement.advertise_custom.advertise_where.port_ranges` | [proxy_advertisement.advertise_custom.advertise_where.port_ranges](resources--bigip_http_proxy--reference--group-002.md#canonical-1333121130033023-0033203310302213-2312201000102001-2123003201020221-3010233323333203-0112012030122312-1132000212122113-2211001211301331) |
| `proxy_advertisement.advertise_custom.advertise_where.site` | [proxy_advertisement.advertise_custom.advertise_where.site](resources--bigip_http_proxy--reference--group-002.md#canonical-0130122121002133-1122221012023322-1233313031202000-3000131030300231-0331002011211222-1203313031022123-2330103220032222-2110122102301030) |
| `proxy_advertisement.advertise_custom.advertise_where.site.ip` | [proxy_advertisement.advertise_custom.advertise_where.site.ip](resources--bigip_http_proxy--reference--group-002.md#canonical-1331232011212220-1113201323130303-0311322010030300-0110303231233323-0032133123301333-2320222023331303-2031311020312121-3130303000021231) |
| `proxy_advertisement.advertise_custom.advertise_where.site.network` | [proxy_advertisement.advertise_custom.advertise_where.site.network](resources--bigip_http_proxy--reference--group-002.md#canonical-3122300101223300-0333311302020123-2030302333102330-1130302103322030-2223332201312031-3123330122102120-2210030131213021-2121132102113022) |
| `proxy_advertisement.advertise_custom.advertise_where.site.site` | [proxy_advertisement.advertise_custom.advertise_where.site.site](resources--bigip_http_proxy--reference--group-002.md#canonical-2013032023310312-0031322113312230-3332200120231012-0001002101013123-0100020003132021-0000122003233322-2200112132303132-0221012231301203) |
| `proxy_advertisement.advertise_custom.advertise_where.site.site.name` | [proxy_advertisement.advertise_custom.advertise_where.site.site.name](resources--bigip_http_proxy--reference--group-002.md#canonical-0201011123020221-0330033131131013-2223231212133022-2132032312023210-3312132331312033-1333133023320321-2110133003323120-1110000111321113) |
| `proxy_advertisement.advertise_custom.advertise_where.site.site.namespace` | [proxy_advertisement.advertise_custom.advertise_where.site.site.namespace](resources--bigip_http_proxy--reference--group-002.md#canonical-0021333213121221-1111132010232103-0003022131033120-2002330113122301-2221302213020010-1212023032023323-0031330123331331-0302232311022310) |
| `proxy_advertisement.advertise_custom.advertise_where.site.site.tenant` | [proxy_advertisement.advertise_custom.advertise_where.site.site.tenant](resources--bigip_http_proxy--reference--group-002.md#canonical-2130110320200002-2110133002112013-2320330233030131-3123212011010022-2302320301231122-2013100300313230-0331332100120331-0301120011332120) |
| `proxy_advertisement.advertise_custom.advertise_where.use_default_port` | [proxy_advertisement.advertise_custom.advertise_where.use_default_port](resources--bigip_http_proxy--reference--group-002.md#canonical-2013220330230101-2003131300000130-1012200023330101-1333300111011111-2013012233321220-0203303100032221-2321301121110103-3000322103322310) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_network` | [proxy_advertisement.advertise_custom.advertise_where.virtual_network](resources--bigip_http_proxy--reference--group-002.md#canonical-2000021012300133-0223211112310033-2331021312120301-2313132120011121-3113333332332113-1322020000320032-0200131310201213-2230332300033033) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_network.default_v6_vip` | [proxy_advertisement.advertise_custom.advertise_where.virtual_network.default_v6_vip](resources--bigip_http_proxy--reference--group-003.md#canonical-3031102321023301-0003233200231212-2100221311200220-1021002103112311-2121132033022033-1101123331113010-3202012121110012-3213323011331103) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_network.default_vip` | [proxy_advertisement.advertise_custom.advertise_where.virtual_network.default_vip](resources--bigip_http_proxy--reference--group-003.md#canonical-2333123311103322-0333002203111322-2010313023322001-0331122223103000-1222300102001212-0033200321233203-3133131011121032-1330311333132333) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_network.specific_v6_vip` | [proxy_advertisement.advertise_custom.advertise_where.virtual_network.specific_v6_vip](resources--bigip_http_proxy--reference--group-002.md#canonical-3032021001323100-1300201313010310-3102031313322320-1113032020320323-1022112212210030-0322230333211300-2233112333333130-3010123031111030) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_network.specific_vip` | [proxy_advertisement.advertise_custom.advertise_where.virtual_network.specific_vip](resources--bigip_http_proxy--reference--group-002.md#canonical-0310203013301232-0022102032113130-3300123120231130-0320021010210123-0231122000222022-2021233021330332-1120211123232010-0020111021223012) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_network.virtual_network` | [proxy_advertisement.advertise_custom.advertise_where.virtual_network.virtual_network](resources--bigip_http_proxy--reference--group-003.md#canonical-1022031103002101-3231301032320231-1213212331322033-1122203102323223-1001122032030333-3110033020100112-1321122200333333-0202022222121102) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_network.virtual_network.name` | [proxy_advertisement.advertise_custom.advertise_where.virtual_network.virtual_network.name](resources--bigip_http_proxy--reference--group-003.md#canonical-0320002002100211-0320000022000020-3223111021211312-1031232232221131-1231232232101020-2130313110100132-1131030330011012-0013020300001302) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_network.virtual_network.namespace` | [proxy_advertisement.advertise_custom.advertise_where.virtual_network.virtual_network.namespace](resources--bigip_http_proxy--reference--group-003.md#canonical-3123332312122310-0331020312300103-0210332230130013-1120310123001102-0202301002200322-1003233131321123-0123222100201111-1123030300112021) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_network.virtual_network.tenant` | [proxy_advertisement.advertise_custom.advertise_where.virtual_network.virtual_network.tenant](resources--bigip_http_proxy--reference--group-003.md#canonical-2321120232220111-0110200332222033-0123103222102023-0131312031323322-0022212331301001-1313333002300202-3002212312210311-2130303331100103) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_site` | [proxy_advertisement.advertise_custom.advertise_where.virtual_site](resources--bigip_http_proxy--reference--group-003.md#canonical-3123331013130022-3233130113001230-2231022020010003-1301212100030103-2112231313201100-2323230330003103-2223313213113322-2200123202013330) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_site.network` | [proxy_advertisement.advertise_custom.advertise_where.virtual_site.network](resources--bigip_http_proxy--reference--group-003.md#canonical-3132133303020301-2123313103311312-1221330302311101-0210211230201022-3210011232011213-1231302112130112-0220330002311000-2021012203023123) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_site.virtual_site` | [proxy_advertisement.advertise_custom.advertise_where.virtual_site.virtual_site](resources--bigip_http_proxy--reference--group-003.md#canonical-3212032320331233-0110203001111003-1213113311213331-0030213020210013-3201211003302002-0133330120022102-3203131031312311-1031003122010022) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_site.virtual_site.name` | [proxy_advertisement.advertise_custom.advertise_where.virtual_site.virtual_site.name](resources--bigip_http_proxy--reference--group-003.md#canonical-2330020322302033-0133011231010320-2101021131302231-0111301120221212-3321012233032031-3132132221203220-0023130202211203-0331133003322230) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_site.virtual_site.namespace` | [proxy_advertisement.advertise_custom.advertise_where.virtual_site.virtual_site.namespace](resources--bigip_http_proxy--reference--group-003.md#canonical-2020122000310211-2023032133333312-1001001102013313-0332101210303201-1010200131233233-3003323232312320-2032311230332221-1102011231223012) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_site.virtual_site.tenant` | [proxy_advertisement.advertise_custom.advertise_where.virtual_site.virtual_site.tenant](resources--bigip_http_proxy--reference--group-003.md#canonical-0323332222331013-0120012112012230-2121323123232122-1221100022330002-2303211203023020-3321133120020230-0210211001311212-2211311303023322) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip` | [proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip](resources--bigip_http_proxy--reference--group-003.md#canonical-0220300231313210-3332311301321130-3030320113320313-3300310102103031-3003312001331100-2123113302222023-2000113330300331-1312213033101121) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip.ip` | [proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip.ip](resources--bigip_http_proxy--reference--group-003.md#canonical-3023233110110223-0112012031033223-3122313322000200-3122102330023132-2330112123322313-0012021000302000-0231330203303003-0310121300110223) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip.network` | [proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip.network](resources--bigip_http_proxy--reference--group-003.md#canonical-2011101021100212-2333221011113000-2233212112220011-0302312103123032-1220002131021311-2003010303300233-0231033033200223-0211103031323101) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip.virtual_site` | [proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip.virtual_site](resources--bigip_http_proxy--reference--group-003.md#canonical-1002001211210232-2213230213230221-2112321102310310-1203230222110311-3133303313213330-2000222130200222-2330233033132232-2002002311111123) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip.virtual_site.name` | [proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip.virtual_site.name](resources--bigip_http_proxy--reference--group-003.md#canonical-3322203021211320-0032023300311101-2101202110112303-1223302233203113-1031331302200130-3311330303220332-1302111321000000-0303100123112311) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip.virtual_site.namespace` | [proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip.virtual_site.namespace](resources--bigip_http_proxy--reference--group-003.md#canonical-2011132310020120-0103301113032110-2233212221323013-1030022200322211-1100120121130111-2100302221021131-3233201230301322-3123332120032111) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip.virtual_site.tenant` | [proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip.virtual_site.tenant](resources--bigip_http_proxy--reference--group-003.md#canonical-2103301231100202-2331001223133322-3110313310200333-2212021103313101-0102333210223031-2210313310233021-1212221121221320-3101100333231030) |
| `proxy_advertisement.advertise_custom.advertise_where.vk8s_service` | [proxy_advertisement.advertise_custom.advertise_where.vk8s_service](resources--bigip_http_proxy--reference--group-003.md#canonical-3301003031203220-1330003030032321-0233023302121333-1101120101330111-3003011300032322-3032233230133012-2022332121321220-0020132330021030) |
| `proxy_advertisement.advertise_custom.advertise_where.vk8s_service.site` | [proxy_advertisement.advertise_custom.advertise_where.vk8s_service.site](resources--bigip_http_proxy--reference--group-003.md#canonical-0310012031112203-2223132312203203-0001211313230131-2222231102010101-2132220320120333-3330331232132100-3220122023022102-1000233332011310) |
| `proxy_advertisement.advertise_custom.advertise_where.vk8s_service.site.name` | [proxy_advertisement.advertise_custom.advertise_where.vk8s_service.site.name](resources--bigip_http_proxy--reference--group-003.md#canonical-3332121322001001-1102133321203012-1332213032333031-2102201023102331-3123101321021122-0211003010021032-1021011023200320-1301112130122321) |
| `proxy_advertisement.advertise_custom.advertise_where.vk8s_service.site.namespace` | [proxy_advertisement.advertise_custom.advertise_where.vk8s_service.site.namespace](resources--bigip_http_proxy--reference--group-003.md#canonical-2202222301231101-2110321011200232-1201322323100130-3231011230322033-0321322320103212-2220201033301013-1301112010311102-3332112121000232) |
| `proxy_advertisement.advertise_custom.advertise_where.vk8s_service.site.tenant` | [proxy_advertisement.advertise_custom.advertise_where.vk8s_service.site.tenant](resources--bigip_http_proxy--reference--group-003.md#canonical-0231132121332322-2101133002010123-2212333310313321-1031203100103133-2020300221011210-2003202232322121-1220211311303132-0131130333021313) |
| `proxy_advertisement.advertise_custom.advertise_where.vk8s_service.virtual_site` | [proxy_advertisement.advertise_custom.advertise_where.vk8s_service.virtual_site](resources--bigip_http_proxy--reference--group-003.md#canonical-2211312030113003-1121203021221302-3130311300032212-0201233220321002-3013003132322011-1300133231011223-0300201001113010-1212112001000130) |
| `proxy_advertisement.advertise_custom.advertise_where.vk8s_service.virtual_site.name` | [proxy_advertisement.advertise_custom.advertise_where.vk8s_service.virtual_site.name](resources--bigip_http_proxy--reference--group-003.md#canonical-2023311033103001-1233202311221312-1310221033223232-0013213213010321-2312321102101120-0121201102100012-0031121210113011-0302200210230322) |
| `proxy_advertisement.advertise_custom.advertise_where.vk8s_service.virtual_site.namespace` | [proxy_advertisement.advertise_custom.advertise_where.vk8s_service.virtual_site.namespace](resources--bigip_http_proxy--reference--group-003.md#canonical-2012122122320120-2312112011002133-3003321120300111-1213230023213320-3200301012202232-0323022100001201-0122230300122200-0232020001312013) |
| `proxy_advertisement.advertise_custom.advertise_where.vk8s_service.virtual_site.tenant` | [proxy_advertisement.advertise_custom.advertise_where.vk8s_service.virtual_site.tenant](resources--bigip_http_proxy--reference--group-003.md#canonical-0320212002300213-2133232013121212-2203101331310233-1203021332113311-2320321100210010-2212332232032331-0232202222103133-1333213203032212) |
| `proxy_advertisement.do_not_advertise` | [proxy_advertisement.do_not_advertise](resources--bigip_http_proxy--reference--group-003.md#canonical-1020302001020012-1111200022023222-0013230011313120-3011031001012211-3220221112212112-3313110121011001-3330011201311013-3230123323013210) |
| `proxy_config` | [proxy_config](resources--bigip_http_proxy--reference--group-003.md#canonical-0321032001132330-3220030230232212-1130321320202323-1011112321033102-0312012203223033-1331130030003002-0021310112303213-3332210012102121) |
| `proxy_config.domains` | [proxy_config.domains](resources--bigip_http_proxy--reference--group-003.md#canonical-0323031323301101-1123202220310110-0213023010203111-1211132003303011-3333123002012113-0013302202301133-1033002132222030-1301222000000021) |
| `proxy_config.http` | [proxy_config.http](resources--bigip_http_proxy--reference--group-003.md#canonical-2033312303200203-3102110023031002-2221300310210132-2211131003033322-0130333102110000-3122333132321131-3203232331122012-3212113300332313) |
| `proxy_config.http.dns_volterra_managed` | [proxy_config.http.dns_volterra_managed](resources--bigip_http_proxy--reference--group-003.md#canonical-1130111333031312-1312121322022202-1303200330122021-0020232133222311-2221010102201131-1232333002230100-1030213112103200-1103232223132113) |
| `proxy_config.http.port` | [proxy_config.http.port](resources--bigip_http_proxy--reference--group-003.md#canonical-0032331110112333-3232212233310031-2332113220323330-1120310101333010-3200332132330002-0203322311233013-0012102322202033-3323223212202131) |
| `proxy_config.http.port_ranges` | [proxy_config.http.port_ranges](resources--bigip_http_proxy--reference--group-003.md#canonical-2112000220313120-2212203110023100-0123222211230130-1301120032301320-2231320311330033-0100212222110111-2020232032111320-1031003200012233) |
| `proxy_config.https` | [proxy_config.https](resources--bigip_http_proxy--reference--group-003.md#canonical-3113232202311223-0130121132011112-2020022212012100-1232012102222301-0330223002132010-2011010113123323-3332121013003100-1231031310202301) |
| `proxy_config.https.add_hsts` | [proxy_config.https.add_hsts](resources--bigip_http_proxy--reference--group-003.md#canonical-0011233003023213-2013302103113212-3200223313232323-1112231002120013-0011123120133300-0230333100230231-1130131101222100-2003103121000020) |
| `proxy_config.https.append_server_name` | [proxy_config.https.append_server_name](resources--bigip_http_proxy--reference--group-003.md#canonical-0330000322012021-1210300011103221-1220023212022000-3322223331213330-0213003023103331-1113030322231233-0120123113111112-2000030130300023) |
| `proxy_config.https.coalescing_options` | [proxy_config.https.coalescing_options](resources--bigip_http_proxy--reference--group-003.md#canonical-1123120103122012-3002213212320332-1213303222100133-2313330212013122-2130201213310200-2023213132130031-3212033231211013-3303232011310133) |
| `proxy_config.https.coalescing_options.default_coalescing` | [proxy_config.https.coalescing_options.default_coalescing](resources--bigip_http_proxy--reference--group-003.md#canonical-1113201100203023-0322202001122031-0002201311000232-1120010130333102-0212120212200132-0102302130320210-3003212313102322-2031131020201233) |
| `proxy_config.https.coalescing_options.strict_coalescing` | [proxy_config.https.coalescing_options.strict_coalescing](resources--bigip_http_proxy--reference--group-003.md#canonical-1303222220230232-1301121333201113-3211221020011002-0022133222100322-3223331321210320-0032332321110223-1010023311230313-1331133032001101) |
| `proxy_config.https.connection_idle_timeout` | [proxy_config.https.connection_idle_timeout](resources--bigip_http_proxy--reference--group-003.md#canonical-3231121300212122-0311230003213132-0312312112230131-1311211332001300-2230120211002231-2310120211102232-2031013311201322-2101002122000030) |
| `proxy_config.https.default_header` | [proxy_config.https.default_header](resources--bigip_http_proxy--reference--group-003.md#canonical-2333033300201021-1330001010030110-3013123320212003-2231113231011030-3020112302212202-1123103312231003-1203013012330222-1320301312330110) |
| `proxy_config.https.default_loadbalancer` | [proxy_config.https.default_loadbalancer](resources--bigip_http_proxy--reference--group-003.md#canonical-0311203212033031-3312023023011121-0320122331211221-1111313321230321-0000211133211113-2102031332230220-1011301333220202-1233110311011322) |
| `proxy_config.https.disable_path_normalize` | [proxy_config.https.disable_path_normalize](resources--bigip_http_proxy--reference--group-003.md#canonical-3223230200123312-3320312110313232-3113113123100232-2330300313111023-0223001001200313-2013300220100200-1330312303221131-0033232210120210) |
| `proxy_config.https.enable_path_normalize` | [proxy_config.https.enable_path_normalize](resources--bigip_http_proxy--reference--group-003.md#canonical-3002223111210111-3303032202311333-2003003230033302-0300003121103233-1321000030203010-1022233010011000-1101211102211102-1223322100200222) |
| `proxy_config.https.http_protocol_options` | [proxy_config.https.http_protocol_options](resources--bigip_http_proxy--reference--group-003.md#canonical-2002102330332111-1023211222312101-2023003021223313-3202222210223131-0212333233320211-2010132210230002-1310020003302130-3112301200103303) |
| `proxy_config.https.http_protocol_options.http_protocol_enable_v1_only` | [proxy_config.https.http_protocol_options.http_protocol_enable_v1_only](resources--bigip_http_proxy--reference--group-003.md#canonical-2130021322001000-3311020211311310-0012122121310011-2230133301010102-2022231302112302-2220300310312013-0031130313222333-3330030031031133) |
| `proxy_config.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation` | [proxy_config.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation](resources--bigip_http_proxy--reference--group-003.md#canonical-0000331132112301-2110210032122110-3003122112112200-0021302031233212-0322232122222122-0231010112223113-1201030012333122-0323030110113022) |
| `proxy_config.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.default_header_transformation` | [proxy_config.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.default_header_transformation](resources--bigip_http_proxy--reference--group-003.md#canonical-3021122201333311-0023112203333120-1300032212002031-0200112321201320-1111330323113331-2012133102322001-1303011200213330-3232030301012322) |
| `proxy_config.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.preserve_case_header_transformation` | [proxy_config.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.preserve_case_header_transformation](resources--bigip_http_proxy--reference--group-003.md#canonical-3333013020131220-1102232030031232-3113130221123203-2223001202022111-1221101202101003-3120302212201231-2303030222021032-3101121211132332) |
| `proxy_config.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.proper_case_header_transformation` | [proxy_config.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.proper_case_header_transformation](resources--bigip_http_proxy--reference--group-003.md#canonical-0020002331120333-3232223011311013-3012321202321013-3123020233233201-0103330222232212-1130002013010311-3112322122012320-2331032231210202) |
| `proxy_config.https.http_protocol_options.http_protocol_enable_v1_v2` | [proxy_config.https.http_protocol_options.http_protocol_enable_v1_v2](resources--bigip_http_proxy--reference--group-003.md#canonical-1322113201320321-0020111120120212-2020032310012201-1111001122020210-0023312133022113-1021222223000331-0113131010233132-1102003322310212) |
| `proxy_config.https.http_protocol_options.http_protocol_enable_v2_only` | [proxy_config.https.http_protocol_options.http_protocol_enable_v2_only](resources--bigip_http_proxy--reference--group-003.md#canonical-0301101221222222-3321312321330131-2012103322002321-0322010332023030-3233030101100230-1020313210212313-1200100231023010-2211212032130031) |
| `proxy_config.https.http_redirect` | [proxy_config.https.http_redirect](resources--bigip_http_proxy--reference--group-003.md#canonical-1011230322303111-2001310102000020-1123033332310231-1030321003330211-1030210002133002-0001102121322320-0333122011231201-2012213213032003) |
| `proxy_config.https.non_default_loadbalancer` | [proxy_config.https.non_default_loadbalancer](resources--bigip_http_proxy--reference--group-003.md#canonical-3121011321013220-0322322003203110-0113320333110021-3212330003310130-0203122002230033-0021210220203030-0121110333330103-3300020120032122) |
| `proxy_config.https.pass_through` | [proxy_config.https.pass_through](resources--bigip_http_proxy--reference--group-003.md#canonical-2100303113221032-3211313031330102-1010121010103101-3302131111021313-1302300103212302-2101103000311001-0223310330002100-0301201121003131) |
| `proxy_config.https.port` | [proxy_config.https.port](resources--bigip_http_proxy--reference--group-003.md#canonical-1313221102122210-2000201233200131-1302220023200122-0200100012021023-0021013233310222-0123230221201223-3131303013212313-3133031223012320) |
| `proxy_config.https.port_ranges` | [proxy_config.https.port_ranges](resources--bigip_http_proxy--reference--group-003.md#canonical-3130020100020311-1330013303003301-3133013022113132-2220101323220121-0330111311122220-2003002013300131-0100322302210330-1201013301313220) |
| `proxy_config.https.server_name` | [proxy_config.https.server_name](resources--bigip_http_proxy--reference--group-003.md#canonical-3031332221013331-1010110220203210-1200013201210120-2101113123120020-0123210311133203-2322230202311301-0223130020200313-1010221032232133) |
| `proxy_config.https.tls_cert_params` | [proxy_config.https.tls_cert_params](resources--bigip_http_proxy--reference--group-003.md#canonical-0233332000132102-1313202110002220-1210132103232310-0011130031230122-2100211121001120-0330212132020020-2331202011301010-0110310320220330) |
| `proxy_config.https.tls_cert_params.certificates` | [proxy_config.https.tls_cert_params.certificates](resources--bigip_http_proxy--reference--group-003.md#canonical-2331102121213232-3113200110102120-2130311031232123-0023131022032013-3310110201230102-3121330232000222-1112023203131030-2233312113200011) |
| `proxy_config.https.tls_cert_params.certificates.name` | [proxy_config.https.tls_cert_params.certificates.name](resources--bigip_http_proxy--reference--group-003.md#canonical-3020130213211231-0110022032110102-0103110201032332-3130011011102112-0000032222200003-0123132112301022-0321313310003111-2101002230132021) |
| `proxy_config.https.tls_cert_params.certificates.namespace` | [proxy_config.https.tls_cert_params.certificates.namespace](resources--bigip_http_proxy--reference--group-003.md#canonical-2221332102220321-1233103023203000-1132122302301322-0032233310020101-1333331321231302-1120223010000332-3232003120111120-1213133331301100) |
| `proxy_config.https.tls_cert_params.certificates.tenant` | [proxy_config.https.tls_cert_params.certificates.tenant](resources--bigip_http_proxy--reference--group-004.md#canonical-1310220210131220-0001132111131111-0320020301231232-3001111231312233-1003032012100032-2020230333123220-1220301033122022-2131101123313200) |
| `proxy_config.https.tls_cert_params.no_mtls` | [proxy_config.https.tls_cert_params.no_mtls](resources--bigip_http_proxy--reference--group-004.md#canonical-3321313102220133-2311213030011233-0012110132133220-3203323230330000-1013313202132332-1232000110102031-0001132120123232-0210022212030103) |
| `proxy_config.https.tls_cert_params.tls_config` | [proxy_config.https.tls_cert_params.tls_config](resources--bigip_http_proxy--reference--group-004.md#canonical-2133230230132223-0013202033222321-1032221331203013-1213030201201130-1012122111333111-0122020311333212-0030313001002123-3011211110212103) |
| `proxy_config.https.tls_cert_params.tls_config.custom_security` | [proxy_config.https.tls_cert_params.tls_config.custom_security](resources--bigip_http_proxy--reference--group-004.md#canonical-0031303213213230-2003331031200220-2032201120012222-1300113230130211-3321122000021131-1133222112012330-2331302010011132-3122231122103321) |
| `proxy_config.https.tls_cert_params.tls_config.custom_security.cipher_suites` | [proxy_config.https.tls_cert_params.tls_config.custom_security.cipher_suites](resources--bigip_http_proxy--reference--group-004.md#canonical-2100012232300311-1022312321330200-0320113031110201-2331220210122300-1332223301233032-3220331312020101-3010122131101022-2001023212020100) |
| `proxy_config.https.tls_cert_params.tls_config.custom_security.max_version` | [proxy_config.https.tls_cert_params.tls_config.custom_security.max_version](resources--bigip_http_proxy--reference--group-004.md#canonical-2131110222323031-1310312332103231-0003203223032310-0211103112012003-0230231030321332-3330232003300312-3200302001020032-0032222000110130) |
| `proxy_config.https.tls_cert_params.tls_config.custom_security.min_version` | [proxy_config.https.tls_cert_params.tls_config.custom_security.min_version](resources--bigip_http_proxy--reference--group-004.md#canonical-2120223232223003-2030201231223003-0131120133022011-3333131302213011-0322333323333033-0320321203031003-1311202231002232-1212222322102011) |
| `proxy_config.https.tls_cert_params.tls_config.default_security` | [proxy_config.https.tls_cert_params.tls_config.default_security](resources--bigip_http_proxy--reference--group-004.md#canonical-3013330031011133-0031101232332000-3320013320321210-2022112121323010-2200333123112331-1312301032101133-0212012003303311-0031300122010323) |
| `proxy_config.https.tls_cert_params.tls_config.low_security` | [proxy_config.https.tls_cert_params.tls_config.low_security](resources--bigip_http_proxy--reference--group-004.md#canonical-2010121312320302-0101102001333221-0103032311000230-0221230332312300-0033111103133232-2123120310213121-2002231312021323-3220230230033103) |
| `proxy_config.https.tls_cert_params.tls_config.medium_security` | [proxy_config.https.tls_cert_params.tls_config.medium_security](resources--bigip_http_proxy--reference--group-004.md#canonical-2322331100001201-2230223332010032-2110202221222210-3220010010231032-0201023110212123-3231021120303112-1132012210021202-1231003202000130) |
| `proxy_config.https.tls_cert_params.use_mtls` | [proxy_config.https.tls_cert_params.use_mtls](resources--bigip_http_proxy--reference--group-004.md#canonical-2002303301031111-2132312201122033-3220010121322013-0333031200101322-0203103201332033-3030000210201212-3320203123132031-1220221212230001) |
| `proxy_config.https.tls_cert_params.use_mtls.client_certificate_optional` | [proxy_config.https.tls_cert_params.use_mtls.client_certificate_optional](resources--bigip_http_proxy--reference--group-004.md#canonical-3013302213222033-3023202323332200-2003131122120220-3031201022201333-1231332233300133-1310320201223312-0312031001032213-1303303333013133) |
| `proxy_config.https.tls_cert_params.use_mtls.crl` | [proxy_config.https.tls_cert_params.use_mtls.crl](resources--bigip_http_proxy--reference--group-004.md#canonical-1100022302103132-2021302201320123-0122000012202212-3221132210021322-1202201031101030-1302010230302232-0323200313222132-3200011223123233) |
| `proxy_config.https.tls_cert_params.use_mtls.crl.name` | [proxy_config.https.tls_cert_params.use_mtls.crl.name](resources--bigip_http_proxy--reference--group-004.md#canonical-3133102110030302-1230311303021030-0313313000311230-3013300130031221-0000223202221333-1310012302032123-3332202133031212-3022312213100312) |
| `proxy_config.https.tls_cert_params.use_mtls.crl.namespace` | [proxy_config.https.tls_cert_params.use_mtls.crl.namespace](resources--bigip_http_proxy--reference--group-004.md#canonical-0321032210002011-1013222001011202-1020121230310020-2311331030000210-3332213103312212-1120303300333012-2213031203030102-2312232330311130) |
| `proxy_config.https.tls_cert_params.use_mtls.crl.tenant` | [proxy_config.https.tls_cert_params.use_mtls.crl.tenant](resources--bigip_http_proxy--reference--group-004.md#canonical-2313032301310211-1030022122233213-3220202310322302-1300012333123332-3321222110230200-2212211102222123-1011013002020001-2333112131210302) |
| `proxy_config.https.tls_cert_params.use_mtls.no_crl` | [proxy_config.https.tls_cert_params.use_mtls.no_crl](resources--bigip_http_proxy--reference--group-004.md#canonical-3020011320103021-1330123313300030-0012020133130110-0332032322023001-0332300202022133-3201101230021303-0211202030211323-1231220220130221) |
| `proxy_config.https.tls_cert_params.use_mtls.trusted_ca` | [proxy_config.https.tls_cert_params.use_mtls.trusted_ca](resources--bigip_http_proxy--reference--group-004.md#canonical-2000232132011120-2011013201322020-1031122031302333-1302021203011212-2020221032312302-1331000203010313-1000222011111201-0103233332312232) |
| `proxy_config.https.tls_cert_params.use_mtls.trusted_ca.name` | [proxy_config.https.tls_cert_params.use_mtls.trusted_ca.name](resources--bigip_http_proxy--reference--group-004.md#canonical-1003202213002313-0031002130223112-1013121001110312-3200311020323310-3102020112003321-2102022112310031-2132332310111321-0201233031302321) |
| `proxy_config.https.tls_cert_params.use_mtls.trusted_ca.namespace` | [proxy_config.https.tls_cert_params.use_mtls.trusted_ca.namespace](resources--bigip_http_proxy--reference--group-004.md#canonical-1320010130103133-0020223220303031-3231323210000230-3023233331131231-0010223232111330-2122110013223310-2223032320210023-2302223020010120) |
| `proxy_config.https.tls_cert_params.use_mtls.trusted_ca.tenant` | [proxy_config.https.tls_cert_params.use_mtls.trusted_ca.tenant](resources--bigip_http_proxy--reference--group-004.md#canonical-2122330320130303-0011132332121021-0203120133010320-1101300202120102-1102330010212331-2012033213121100-2032031012321023-3132131000303301) |
| `proxy_config.https.tls_cert_params.use_mtls.trusted_ca_url` | [proxy_config.https.tls_cert_params.use_mtls.trusted_ca_url](resources--bigip_http_proxy--reference--group-004.md#canonical-0003220231231303-3130033332303332-3322310321001232-3020013321023310-0203000303200202-3302312200201110-2210302301030312-0122303302220230) |
| `proxy_config.https.tls_cert_params.use_mtls.xfcc_disabled` | [proxy_config.https.tls_cert_params.use_mtls.xfcc_disabled](resources--bigip_http_proxy--reference--group-004.md#canonical-0111033302023221-3310003303303131-1310332121313111-1013031100131132-0130010033111122-1113222120323210-3200300013123322-3203323313213131) |
| `proxy_config.https.tls_cert_params.use_mtls.xfcc_options` | [proxy_config.https.tls_cert_params.use_mtls.xfcc_options](resources--bigip_http_proxy--reference--group-004.md#canonical-0332212230300113-2113001211110101-1300133321321132-3203023313200121-3101231200331111-2013121011120303-2311113123200033-1003222322321120) |
| `proxy_config.https.tls_cert_params.use_mtls.xfcc_options.xfcc_header_elements` | [proxy_config.https.tls_cert_params.use_mtls.xfcc_options.xfcc_header_elements](resources--bigip_http_proxy--reference--group-004.md#canonical-2212233001102010-0302012322301022-1331112010300111-0102001013012231-2301313011311012-3012201000313230-0123301213020113-2001130100331302) |
| `proxy_config.https.tls_parameters` | [proxy_config.https.tls_parameters](resources--bigip_http_proxy--reference--group-004.md#canonical-0120102120100020-1100011221033102-0221220220332001-1333311021222120-2023123132110312-0002031111321112-1022231022201102-3131022230333133) |
| `proxy_config.https.tls_parameters.no_mtls` | [proxy_config.https.tls_parameters.no_mtls](resources--bigip_http_proxy--reference--group-004.md#canonical-3310030003230002-1220330232102210-3310322322222322-3120220231131103-2013100022001031-0232220231320012-3222200121113102-2231011021001022) |
| `proxy_config.https.tls_parameters.tls_certificates` | [proxy_config.https.tls_parameters.tls_certificates](resources--bigip_http_proxy--reference--group-004.md#canonical-0220301032330333-1323013100331100-0012033233323112-2320332311233123-3212210220220222-0112021100023210-3201100302313332-0122102311013310) |
| `proxy_config.https.tls_parameters.tls_certificates.certificate_url` | [proxy_config.https.tls_parameters.tls_certificates.certificate_url](resources--bigip_http_proxy--reference--group-004.md#canonical-0032213100120331-0231021203102030-3023322201023331-1112010001322213-3132232010200100-0320222200322132-3211320213133323-0013323130333301) |
| `proxy_config.https.tls_parameters.tls_certificates.custom_hash_algorithms` | [proxy_config.https.tls_parameters.tls_certificates.custom_hash_algorithms](resources--bigip_http_proxy--reference--group-004.md#canonical-2333232103133300-0103332332102100-3233300330020313-0131011021330123-1023112210330310-3321023123023210-2132112123212203-3313221233102030) |
| `proxy_config.https.tls_parameters.tls_certificates.custom_hash_algorithms.hash_algorithms` | [proxy_config.https.tls_parameters.tls_certificates.custom_hash_algorithms.hash_algorithms](resources--bigip_http_proxy--reference--group-004.md#canonical-1100222003230221-2033331123031030-3220021020133123-3012122121031213-3002223122312312-1011012231100202-2131011223020320-2331303222311231) |
| `proxy_config.https.tls_parameters.tls_certificates.description_spec` | [proxy_config.https.tls_parameters.tls_certificates.description_spec](resources--bigip_http_proxy--reference--group-004.md#canonical-3133310331011323-0312013011321130-3133332322303112-2232201032230200-0132020301023012-0013213132203202-2232223120002310-3221330123213032) |
| `proxy_config.https.tls_parameters.tls_certificates.disable_ocsp_stapling` | [proxy_config.https.tls_parameters.tls_certificates.disable_ocsp_stapling](resources--bigip_http_proxy--reference--group-004.md#canonical-3233023212023113-3300203231011311-2232232201223103-2311313202200030-1313313102330213-3112331101121320-0330033210230300-0312131322102203) |
| `proxy_config.https.tls_parameters.tls_certificates.private_key` | [proxy_config.https.tls_parameters.tls_certificates.private_key](resources--bigip_http_proxy--reference--group-004.md#canonical-3013133203032233-3112011013132000-0201232133010111-3321300130121111-0211113323130333-0223111102210233-1321012113010101-2100001101012221) |
| `proxy_config.https.tls_parameters.tls_certificates.private_key.blindfold_secret_info` | [proxy_config.https.tls_parameters.tls_certificates.private_key.blindfold_secret_info](resources--bigip_http_proxy--reference--group-004.md#canonical-3002230331121022-2011220313320330-2221223233211031-3212133202023210-0003333232030203-1221312322100203-2110002312132312-1303232011130120) |
| `proxy_config.https.tls_parameters.tls_certificates.private_key.blindfold_secret_info.decryption_provider` | [proxy_config.https.tls_parameters.tls_certificates.private_key.blindfold_secret_info.decryption_provider](resources--bigip_http_proxy--reference--group-004.md#canonical-0212203223200200-2113132031320133-2132003302301211-2032130301011120-3210003130222203-0312202003201033-3223002230220022-0232200033113102) |
| `proxy_config.https.tls_parameters.tls_certificates.private_key.blindfold_secret_info.location` | [proxy_config.https.tls_parameters.tls_certificates.private_key.blindfold_secret_info.location](resources--bigip_http_proxy--reference--group-004.md#canonical-0121123033113033-1022122331121000-2300103130203112-0213310122131231-2231020003332220-3003231111321113-3232112332320231-0211321310022301) |
| `proxy_config.https.tls_parameters.tls_certificates.private_key.blindfold_secret_info.store_provider` | [proxy_config.https.tls_parameters.tls_certificates.private_key.blindfold_secret_info.store_provider](resources--bigip_http_proxy--reference--group-004.md#canonical-0211002203112010-1101011032112213-0113213212112323-0013311030333312-2233033312212201-0211312122001100-2302320332300003-2100233231323110) |
| `proxy_config.https.tls_parameters.tls_certificates.private_key.clear_secret_info` | [proxy_config.https.tls_parameters.tls_certificates.private_key.clear_secret_info](resources--bigip_http_proxy--reference--group-004.md#canonical-3101321102031322-2010221022033033-3110011233010102-1132312122322000-1121203302023131-3301130212231331-0313000010132021-1200121313012321) |
| `proxy_config.https.tls_parameters.tls_certificates.private_key.clear_secret_info.provider_ref` | [proxy_config.https.tls_parameters.tls_certificates.private_key.clear_secret_info.provider_ref](resources--bigip_http_proxy--reference--group-004.md#canonical-0110203200231022-2023013001221300-2203303231020311-2032020312220302-1203232200100002-2030101223302013-3011021003132223-2201332211032312) |
| `proxy_config.https.tls_parameters.tls_certificates.private_key.clear_secret_info.url` | [proxy_config.https.tls_parameters.tls_certificates.private_key.clear_secret_info.url](resources--bigip_http_proxy--reference--group-004.md#canonical-2023202121220333-3310223010223223-3323320130200213-2022103011011220-3022133301303232-2130320133302202-0031110032320232-3230323032122111) |
| `proxy_config.https.tls_parameters.tls_certificates.use_system_defaults` | [proxy_config.https.tls_parameters.tls_certificates.use_system_defaults](resources--bigip_http_proxy--reference--group-004.md#canonical-3300203102333201-1311112211011302-0031000102110232-2312210220211233-1002211132111022-1232202302023033-2232132100310132-0111012201021031) |
| `proxy_config.https.tls_parameters.tls_config` | [proxy_config.https.tls_parameters.tls_config](resources--bigip_http_proxy--reference--group-004.md#canonical-2103323030031233-2211133023111033-2132023100322132-2331320220012233-0223223032201022-1301023100321302-1102000013211332-3302012003202223) |
| `proxy_config.https.tls_parameters.tls_config.custom_security` | [proxy_config.https.tls_parameters.tls_config.custom_security](resources--bigip_http_proxy--reference--group-004.md#canonical-2210003101010303-3213302030201200-1223323223330133-1101110322332312-1322302203210230-1310233210130310-3032112312000001-1303200130311330) |
| `proxy_config.https.tls_parameters.tls_config.custom_security.cipher_suites` | [proxy_config.https.tls_parameters.tls_config.custom_security.cipher_suites](resources--bigip_http_proxy--reference--group-004.md#canonical-3323033112233302-2102233202022010-2031312223330323-0322112102033321-1130013033033131-0200122333003312-2310031023302223-1101200201111303) |
| `proxy_config.https.tls_parameters.tls_config.custom_security.max_version` | [proxy_config.https.tls_parameters.tls_config.custom_security.max_version](resources--bigip_http_proxy--reference--group-004.md#canonical-2233022230022031-3022210133310300-3311001103120332-0311130301100121-0200020103102300-2210011200302101-1013030320212223-3320200201111120) |
| `proxy_config.https.tls_parameters.tls_config.custom_security.min_version` | [proxy_config.https.tls_parameters.tls_config.custom_security.min_version](resources--bigip_http_proxy--reference--group-004.md#canonical-0101101111032333-0230110103311002-1210133112302333-2000033231011000-0201021000220002-0303311022232002-0330031210020100-1210331230020100) |
| `proxy_config.https.tls_parameters.tls_config.default_security` | [proxy_config.https.tls_parameters.tls_config.default_security](resources--bigip_http_proxy--reference--group-004.md#canonical-1022333100332033-2202313233010321-0033022011010231-0011021211031123-3002112000013110-0333322103010110-2112121010233130-3100123132023301) |
| `proxy_config.https.tls_parameters.tls_config.low_security` | [proxy_config.https.tls_parameters.tls_config.low_security](resources--bigip_http_proxy--reference--group-004.md#canonical-0212100230111030-1120211311122313-0130021100300301-0133232230003123-3130133203032303-2111020323111011-0311231333021201-2012120331102000) |
| `proxy_config.https.tls_parameters.tls_config.medium_security` | [proxy_config.https.tls_parameters.tls_config.medium_security](resources--bigip_http_proxy--reference--group-004.md#canonical-0323133222230223-3132333012313201-0310113302022122-1210022232321131-3031201022012310-3330300002021201-0013300232103323-1331110032131202) |
| `proxy_config.https.tls_parameters.use_mtls` | [proxy_config.https.tls_parameters.use_mtls](resources--bigip_http_proxy--reference--group-004.md#canonical-2202303331113300-3101033121011310-0211110012020211-3202333122112301-0023232301101203-0111110122200310-0230301102120103-3200322322032202) |
| `proxy_config.https.tls_parameters.use_mtls.client_certificate_optional` | [proxy_config.https.tls_parameters.use_mtls.client_certificate_optional](resources--bigip_http_proxy--reference--group-004.md#canonical-1323201330232021-1212100312030110-0020020201022100-0020210302330120-1310203330332002-3122112003211233-2003210131011302-2300010113202331) |
| `proxy_config.https.tls_parameters.use_mtls.crl` | [proxy_config.https.tls_parameters.use_mtls.crl](resources--bigip_http_proxy--reference--group-004.md#canonical-0123233322023122-3310221222012200-2020023031002133-1123201333210201-2322013300203211-3321232331223333-0223313223201132-2332302211210222) |
| `proxy_config.https.tls_parameters.use_mtls.crl.name` | [proxy_config.https.tls_parameters.use_mtls.crl.name](resources--bigip_http_proxy--reference--group-004.md#canonical-0012120003132232-3000121133223031-0022202203213133-2331133320002333-0100222012323013-1220302110223101-2123123130303213-0220030133111102) |
| `proxy_config.https.tls_parameters.use_mtls.crl.namespace` | [proxy_config.https.tls_parameters.use_mtls.crl.namespace](resources--bigip_http_proxy--reference--group-004.md#canonical-3122230101301020-1113132300300211-3302132230010101-3213312010323232-2302300133231223-0330121311303310-0232031121322301-3322330033313312) |
| `proxy_config.https.tls_parameters.use_mtls.crl.tenant` | [proxy_config.https.tls_parameters.use_mtls.crl.tenant](resources--bigip_http_proxy--reference--group-004.md#canonical-1303233313232301-2303300000020320-0133010030102333-3221033112102203-3013323231001321-3111033020213000-0110032201013323-1123020210220221) |
| `proxy_config.https.tls_parameters.use_mtls.no_crl` | [proxy_config.https.tls_parameters.use_mtls.no_crl](resources--bigip_http_proxy--reference--group-004.md#canonical-2101011211001301-0201311300323303-2201111013301002-2221031131030200-1010232101101003-1323003122033203-0300111203311003-3012012030231132) |
| `proxy_config.https.tls_parameters.use_mtls.trusted_ca` | [proxy_config.https.tls_parameters.use_mtls.trusted_ca](resources--bigip_http_proxy--reference--group-004.md#canonical-1133110232210011-2201103002102002-3300133300101313-1300002131011010-1102313112002022-2010003113123022-2332030301132321-2301012231133302) |
| `proxy_config.https.tls_parameters.use_mtls.trusted_ca.name` | [proxy_config.https.tls_parameters.use_mtls.trusted_ca.name](resources--bigip_http_proxy--reference--group-004.md#canonical-0221323322330333-0100032031320121-2203002301212221-1033030131321231-2121120111202113-1232301010000213-0122131201311030-1102311320321331) |
| `proxy_config.https.tls_parameters.use_mtls.trusted_ca.namespace` | [proxy_config.https.tls_parameters.use_mtls.trusted_ca.namespace](resources--bigip_http_proxy--reference--group-004.md#canonical-3103003221321002-1313113211002222-2220311203312212-2302000230201321-3123232102020031-2003300300103223-1121313111101031-1110302211033012) |
| `proxy_config.https.tls_parameters.use_mtls.trusted_ca.tenant` | [proxy_config.https.tls_parameters.use_mtls.trusted_ca.tenant](resources--bigip_http_proxy--reference--group-004.md#canonical-0022203003022302-1131002012121320-3330223230322021-0030111130112002-3210231032020032-0132313112100223-3020132132130123-3232032200333010) |
| `proxy_config.https.tls_parameters.use_mtls.trusted_ca_url` | [proxy_config.https.tls_parameters.use_mtls.trusted_ca_url](resources--bigip_http_proxy--reference--group-004.md#canonical-1321220310111301-1003313120032121-1213103103013301-0223300031223212-1122300033032212-2321330100301100-1131023101100221-2310112221020200) |
| `proxy_config.https.tls_parameters.use_mtls.xfcc_disabled` | [proxy_config.https.tls_parameters.use_mtls.xfcc_disabled](resources--bigip_http_proxy--reference--group-004.md#canonical-2010022220303330-3000033201310110-2101011330223312-2021222033112200-1131330023322120-3123331023020320-0033223001032133-2302012110010220) |
| `proxy_config.https.tls_parameters.use_mtls.xfcc_options` | [proxy_config.https.tls_parameters.use_mtls.xfcc_options](resources--bigip_http_proxy--reference--group-004.md#canonical-2012203301003332-0100201131320312-2113020323131013-3301311113322132-1131210332321203-1232030201233331-3232113310333210-3302121310212100) |
| `proxy_config.https.tls_parameters.use_mtls.xfcc_options.xfcc_header_elements` | [proxy_config.https.tls_parameters.use_mtls.xfcc_options.xfcc_header_elements](resources--bigip_http_proxy--reference--group-004.md#canonical-0231031333310200-2120233223033001-3110213031133110-0303131230132000-0100311020312221-3323122132112010-2023202000232221-3210120322201203) |
| `proxy_config.https_auto_cert` | [proxy_config.https_auto_cert](resources--bigip_http_proxy--reference--group-004.md#canonical-1232130333122220-1122323101230313-3221320113210210-0220120212020321-3021313123110330-3103112331000021-3123222113203022-3022132121330232) |
| `proxy_config.https_auto_cert.add_hsts` | [proxy_config.https_auto_cert.add_hsts](resources--bigip_http_proxy--reference--group-004.md#canonical-1021302330123221-1213000312303202-0300312012313312-3332123011132003-2221001220123113-0301032210010130-3200333222220301-2231021323213320) |
| `proxy_config.https_auto_cert.append_server_name` | [proxy_config.https_auto_cert.append_server_name](resources--bigip_http_proxy--reference--group-004.md#canonical-2013312310300111-0210301002301331-2223112020303113-2033022333121003-0123122221332302-3110133321333031-3100130031313212-0200203022031101) |
| `proxy_config.https_auto_cert.coalescing_options` | [proxy_config.https_auto_cert.coalescing_options](resources--bigip_http_proxy--reference--group-004.md#canonical-2023122032222133-0320312030031012-0233013300033023-1221200320321120-3133332302333212-1123032100210030-0130231233322030-1120012213013033) |
| `proxy_config.https_auto_cert.coalescing_options.default_coalescing` | [proxy_config.https_auto_cert.coalescing_options.default_coalescing](resources--bigip_http_proxy--reference--group-004.md#canonical-1111301103121032-0212011113313202-2110301133301300-1131320230302300-1000211023121031-3213320000010211-1203312213133321-2233331101122311) |
| `proxy_config.https_auto_cert.coalescing_options.strict_coalescing` | [proxy_config.https_auto_cert.coalescing_options.strict_coalescing](resources--bigip_http_proxy--reference--group-004.md#canonical-3010312210331113-1332031313112333-1203120212222333-3131123203123031-0232323022110332-0211303211300311-3112211022131311-0130113320003123) |
| `proxy_config.https_auto_cert.connection_idle_timeout` | [proxy_config.https_auto_cert.connection_idle_timeout](resources--bigip_http_proxy--reference--group-004.md#canonical-0120223312103123-1032030101311003-0000120121302233-1030021320203203-3323202301120112-1300102022312220-0201130132113003-2222120310000023) |
| `proxy_config.https_auto_cert.default_header` | [proxy_config.https_auto_cert.default_header](resources--bigip_http_proxy--reference--group-004.md#canonical-3213133110002310-0013122112302013-3123322222020302-2330120220031131-0203132120210131-2333210023032210-3021313210022233-0320022322031101) |
| `proxy_config.https_auto_cert.default_loadbalancer` | [proxy_config.https_auto_cert.default_loadbalancer](resources--bigip_http_proxy--reference--group-004.md#canonical-1012301213233321-3001220333032221-0220022203120031-0221020003130002-0233211221320112-3021201113020232-3213220112032232-2233231321200133) |
| `proxy_config.https_auto_cert.disable_path_normalize` | [proxy_config.https_auto_cert.disable_path_normalize](resources--bigip_http_proxy--reference--group-004.md#canonical-3220013120110313-3313131231202332-3122133013011123-1203123220120330-2211332312100220-3332123303011222-0112032102312313-0033332013020122) |
| `proxy_config.https_auto_cert.enable_path_normalize` | [proxy_config.https_auto_cert.enable_path_normalize](resources--bigip_http_proxy--reference--group-004.md#canonical-1220202012033120-1312202313313322-2123031100113031-3000313021200012-2202021021030323-0320321331001021-2021031122332232-1031220022013203) |
| `proxy_config.https_auto_cert.http_protocol_options` | [proxy_config.https_auto_cert.http_protocol_options](resources--bigip_http_proxy--reference--group-004.md#canonical-0222300021332000-2201002013200322-3323321302320123-0332321112311323-0021312230312301-2031312023032221-1003001303323233-2322030123112311) |
| `proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only` | [proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only](resources--bigip_http_proxy--reference--group-004.md#canonical-3201111020112210-3030000310010212-0110333212112101-0211101201020130-1211113311123111-1201323311011213-1020020020312010-0313331002002101) |
| `proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation` | [proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation](resources--bigip_http_proxy--reference--group-004.md#canonical-2312031202133103-1112123123021233-3011202231313100-0302223330200300-1030001122213322-0022201000001203-0211121110203221-2120120332110310) |
| `proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.default_header_transformation` | [proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.default_header_transformation](resources--bigip_http_proxy--reference--group-004.md#canonical-0310211332313123-0312102111032001-2303223022211321-1323123311221112-1221220322101130-2211101003130120-1123212301312132-2000001031033123) |
| `proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.preserve_case_header_transformation` | [proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.preserve_case_header_transformation](resources--bigip_http_proxy--reference--group-004.md#canonical-1012232312321101-0203222303321100-3012202112323122-0111022012101111-3031030322000230-2011111111321233-2333331311230112-3230002100031300) |
| `proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.proper_case_header_transformation` | [proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.proper_case_header_transformation](resources--bigip_http_proxy--reference--group-004.md#canonical-2031333333120130-1313333112313030-3110132020032133-2120222100132220-3330211323132222-2220303320133110-1331010103331323-3100230330012010) |
| `proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v1_v2` | [proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v1_v2](resources--bigip_http_proxy--reference--group-004.md#canonical-0131313211230320-2021232221321321-1332111233200011-3023020201330312-1130233230110013-3312320213123131-0032333122232010-3211100302220131) |
| `proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v2_only` | [proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v2_only](resources--bigip_http_proxy--reference--group-004.md#canonical-3212020033102322-1213212323032120-0213200022322013-0301132032210203-2123022130203312-1233303111011011-0031132332020321-0203212120220301) |
| `proxy_config.https_auto_cert.http_redirect` | [proxy_config.https_auto_cert.http_redirect](resources--bigip_http_proxy--reference--group-004.md#canonical-1222321203200333-3031011110100312-0111110013013200-0101110232001131-1011022313132201-2101012233030011-3030123020213222-0022013012320220) |
| `proxy_config.https_auto_cert.no_mtls` | [proxy_config.https_auto_cert.no_mtls](resources--bigip_http_proxy--reference--group-004.md#canonical-1322331222021311-0102013012323033-3011303122032322-0003232210321233-2203203301330223-3013312010111323-0232111113102122-1311032021230010) |
| `proxy_config.https_auto_cert.non_default_loadbalancer` | [proxy_config.https_auto_cert.non_default_loadbalancer](resources--bigip_http_proxy--reference--group-004.md#canonical-2220003321303202-1233223212003031-1222030023201220-1301123311121313-0232031332322311-2313332033001133-0030020313223031-0110022100302222) |
| `proxy_config.https_auto_cert.pass_through` | [proxy_config.https_auto_cert.pass_through](resources--bigip_http_proxy--reference--group-004.md#canonical-1011133132220113-0100022133303333-0221231133011033-1002002123132211-3023320301231111-1031133031030322-3022032122233112-3222010210212021) |
| `proxy_config.https_auto_cert.port` | [proxy_config.https_auto_cert.port](resources--bigip_http_proxy--reference--group-004.md#canonical-2030023110120333-2011333201022013-3323133032330233-2010332011100312-0210033032233131-0203111033011300-0301312303313012-1033012210323011) |
| `proxy_config.https_auto_cert.port_ranges` | [proxy_config.https_auto_cert.port_ranges](resources--bigip_http_proxy--reference--group-004.md#canonical-0131123230120033-1310101200201030-2020332122330312-2313231221221103-2131110020012100-2132220220312212-0300121330110210-0203323011000213) |
| `proxy_config.https_auto_cert.server_name` | [proxy_config.https_auto_cert.server_name](resources--bigip_http_proxy--reference--group-004.md#canonical-3222011031201331-1011131323210012-2121202201131011-2321320301122232-0311200333112331-2031030113203221-2302013023323013-3012330012011031) |
| `proxy_config.https_auto_cert.tls_config` | [proxy_config.https_auto_cert.tls_config](resources--bigip_http_proxy--reference--group-004.md#canonical-1101321313001121-1213000233030003-3313012310020000-2330012200331112-3220032121011033-1011213033101221-2203111301033030-2202212030010211) |
| `proxy_config.https_auto_cert.tls_config.custom_security` | [proxy_config.https_auto_cert.tls_config.custom_security](resources--bigip_http_proxy--reference--group-004.md#canonical-2313203210100331-3033303200001332-0330020331233103-1003001013203200-2212030130112201-2003232021001031-1201133311320120-3103111121112223) |
| `proxy_config.https_auto_cert.tls_config.custom_security.cipher_suites` | [proxy_config.https_auto_cert.tls_config.custom_security.cipher_suites](resources--bigip_http_proxy--reference--group-004.md#canonical-1220222203332022-1213001002122231-0212200110331121-1302000202210220-1131001301023120-1013310330320311-1212011223111221-1320312102230110) |
| `proxy_config.https_auto_cert.tls_config.custom_security.max_version` | [proxy_config.https_auto_cert.tls_config.custom_security.max_version](resources--bigip_http_proxy--reference--group-004.md#canonical-2012113233121211-0203031012013333-2300130132123111-0310032322101213-2131310101231030-2030233303010322-1321212012023123-1330202210210332) |
| `proxy_config.https_auto_cert.tls_config.custom_security.min_version` | [proxy_config.https_auto_cert.tls_config.custom_security.min_version](resources--bigip_http_proxy--reference--group-004.md#canonical-3031111101301111-0001231222303132-3023130202022013-1223313113130120-3230032032201313-1221333103332120-2200223022231013-1200112131203032) |
| `proxy_config.https_auto_cert.tls_config.default_security` | [proxy_config.https_auto_cert.tls_config.default_security](resources--bigip_http_proxy--reference--group-004.md#canonical-2202321002001131-2333010123133101-0011030321323120-1232313310233022-2120120021321132-3313011331301323-2222102103113020-3321002131003302) |
| `proxy_config.https_auto_cert.tls_config.low_security` | [proxy_config.https_auto_cert.tls_config.low_security](resources--bigip_http_proxy--reference--group-004.md#canonical-0131122301123112-2202202013133113-1221311312302230-1300033220330312-3310331212210203-0320001011231020-0202333100023021-3132013101020331) |
| `proxy_config.https_auto_cert.tls_config.medium_security` | [proxy_config.https_auto_cert.tls_config.medium_security](resources--bigip_http_proxy--reference--group-004.md#canonical-1321233002030111-3132110222313321-3111101310030312-2031121301122300-0211120233323123-3031112020210313-2320301013101111-1032300113311130) |
| `proxy_config.https_auto_cert.use_mtls` | [proxy_config.https_auto_cert.use_mtls](resources--bigip_http_proxy--reference--group-004.md#canonical-0322202010222322-3213113223233213-1222022331200032-2110231101321221-1331110231102212-0223001113332212-1021331213120221-1013320011212211) |
| `proxy_config.https_auto_cert.use_mtls.client_certificate_optional` | [proxy_config.https_auto_cert.use_mtls.client_certificate_optional](resources--bigip_http_proxy--reference--group-004.md#canonical-2123130003021302-3310132121121133-2013001301320202-1033210312221022-2323123312300111-3102203322131230-0212300101311131-2010110332031231) |
| `proxy_config.https_auto_cert.use_mtls.crl` | [proxy_config.https_auto_cert.use_mtls.crl](resources--bigip_http_proxy--reference--group-004.md#canonical-3122100112000233-0311120333110101-3212101321123231-2202010132012331-0033101022131030-3300201221020030-1123330330203132-2322022110013331) |
| `proxy_config.https_auto_cert.use_mtls.crl.name` | [proxy_config.https_auto_cert.use_mtls.crl.name](resources--bigip_http_proxy--reference--group-004.md#canonical-2020200021202132-1120122212130130-2012232132011022-0230032223133222-0130100331232202-1022210310210301-1033322010300112-2203131020201312) |
| `proxy_config.https_auto_cert.use_mtls.crl.namespace` | [proxy_config.https_auto_cert.use_mtls.crl.namespace](resources--bigip_http_proxy--reference--group-004.md#canonical-3113223212103113-2323230221002121-1323233001011113-3211313330212202-0133010011102231-1213013021010331-1032031303233310-0110103030221332) |
| `proxy_config.https_auto_cert.use_mtls.crl.tenant` | [proxy_config.https_auto_cert.use_mtls.crl.tenant](resources--bigip_http_proxy--reference--group-004.md#canonical-1112220102103213-1331230331232231-3310220330202123-0322330001033112-3303232111203320-0211320203223331-2013020112232213-0210101232000201) |
| `proxy_config.https_auto_cert.use_mtls.no_crl` | [proxy_config.https_auto_cert.use_mtls.no_crl](resources--bigip_http_proxy--reference--group-004.md#canonical-0331033310013030-0022023111122310-1320200321132021-0002132032312300-0230321310230213-2301131010113132-3021013112023310-1300223223310312) |
| `proxy_config.https_auto_cert.use_mtls.trusted_ca` | [proxy_config.https_auto_cert.use_mtls.trusted_ca](resources--bigip_http_proxy--reference--group-004.md#canonical-1022212220320001-1323001123311200-0203031020331131-3001031320123120-0300323201103002-0313013012200332-1010120203111302-0133232033220103) |
| `proxy_config.https_auto_cert.use_mtls.trusted_ca.name` | [proxy_config.https_auto_cert.use_mtls.trusted_ca.name](resources--bigip_http_proxy--reference--group-004.md#canonical-3230313211001333-0030221213223130-3022110110310110-1221223302200311-1200131131202011-3131321031003321-2212000032123232-3003330200311201) |
| `proxy_config.https_auto_cert.use_mtls.trusted_ca.namespace` | [proxy_config.https_auto_cert.use_mtls.trusted_ca.namespace](resources--bigip_http_proxy--reference--group-004.md#canonical-1133311202323131-1203002110031302-1120023012122321-1023300201222322-3133101132003320-0033322023232100-0301032112323032-1003120013111111) |
| `proxy_config.https_auto_cert.use_mtls.trusted_ca.tenant` | [proxy_config.https_auto_cert.use_mtls.trusted_ca.tenant](resources--bigip_http_proxy--reference--group-004.md#canonical-0123122023022310-2233012221302221-2221032103100031-2103332233020111-1020110332013100-3233312023301213-3211333032133221-0123333033000232) |
| `proxy_config.https_auto_cert.use_mtls.trusted_ca_url` | [proxy_config.https_auto_cert.use_mtls.trusted_ca_url](resources--bigip_http_proxy--reference--group-004.md#canonical-0010310332131023-2011322023211001-1103002232210212-1101102213333321-0122100302133231-0013132220312113-3002130131102033-1102333103012302) |
| `proxy_config.https_auto_cert.use_mtls.xfcc_disabled` | [proxy_config.https_auto_cert.use_mtls.xfcc_disabled](resources--bigip_http_proxy--reference--group-004.md#canonical-0332021101303111-2321212313232113-2311232212001311-2212203011011021-1221333101001300-2013200220200333-1223323310232110-0033131013000023) |
| `proxy_config.https_auto_cert.use_mtls.xfcc_options` | [proxy_config.https_auto_cert.use_mtls.xfcc_options](resources--bigip_http_proxy--reference--group-004.md#canonical-1110022331333333-1303013102221111-1213333230130330-1030133231302313-1112220300122311-1003103212200100-3301013200031011-2122101130311110) |
| `proxy_config.https_auto_cert.use_mtls.xfcc_options.xfcc_header_elements` | [proxy_config.https_auto_cert.use_mtls.xfcc_options.xfcc_header_elements](resources--bigip_http_proxy--reference--group-004.md#canonical-1112003113111331-2123221101023333-3233023323222010-3132322001133120-1111001230300223-2323310313221100-1201321031012002-1310113030111332) |
| `timeouts` | [timeouts](resources--bigip_http_proxy--reference--group-004.md#canonical-1120110110032221-2003012033023222-3122333201020321-3201331110133123-0233100003200311-0001322102030133-3011131312323321-3122331223231202) |
| `timeouts.create` | [timeouts.create](resources--bigip_http_proxy--reference--group-004.md#canonical-3021002223001012-0202020332321100-1312202302020313-3011210111231030-3213131023322321-2103102303123301-0301001032030221-3230333202203011) |
| `timeouts.delete` | [timeouts.delete](resources--bigip_http_proxy--reference--group-004.md#canonical-1221203203020013-1111003330301122-3222020003321200-2021212331111321-2100010321032303-2233033222023113-0222023310011022-0331232212311112) |
| `timeouts.read` | [timeouts.read](resources--bigip_http_proxy--reference--group-004.md#canonical-3111202213100121-2023201200020303-2203123210101010-0311230012132300-2031113223221113-2311030330020320-1102202032333323-3032023221011013) |
| `timeouts.update` | [timeouts.update](resources--bigip_http_proxy--reference--group-004.md#canonical-1102310223003003-3233021331311302-3201131102023300-0121100132013020-1203033102020030-0001010201201022-1132121332021222-3103333330003310) |

<a id="canonical-2112133030313203-1202233221221323-0231033212122201-1200202122332313-2022030302010022-1221300001120300-3120211033311213-2323103320212230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `advanced_profile` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-0101201332222201-3302033021001322-1302101210311331-2210333210320003-0002033231003211-1133020003110203-0200133113131131-0101001111122000)
- advanced_profile

<a id="canonical-2202101110120110-1303332222001222-0101031233131220-1321323001223202-0310111203221022-0121201102003320-3122213002331303-3323322131233320"></a>

Type: `"object"`. single nested block, Optional.

This defines various advanced Profile OPTIONS for a Loadbalancer.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("disable_spec",
    "enable_default_profile")}
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
  "x-ves-oneof-field-choice": "[\"disable\",\"enable_default_profile\"]"
}
```

Terraform syntax:

```terraform
advanced_profile {
  # Configure direct properties listed below.
}
```

<a id="canonical-3001100202212120-2110121302221320-0122211330021333-0332221003131120-0230321211001333-1101302230001033-2320230332001233-3310011232002302"></a>

### Direct properties for `advanced_profile`

- [disable_spec](resources--bigip_http_proxy--reference--group-001.md#canonical-1031030132100222-0002211222122030-3221302130313303-0312000200310233-1220010033231131-3222020210303210-3120021331103212-2132213030020010): complete subsection reference.

- [enable_default_profile](resources--bigip_http_proxy--reference--group-001.md#canonical-3133021223323122-3320012200231021-3301031021103333-1013210113222003-1132233333333102-3313303211213212-1000321311333320-1011120110101032): complete subsection reference.

<a id="canonical-1031030132100222-0002211222122030-3221302130313303-0312000200310233-1220010033231131-3222020210303210-3120021331103212-2132213030020010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `advanced_profile.disable_spec` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-0101201332222201-3302033021001322-1302101210311331-2210333210320003-0002033231003211-1133020003110203-0200133113131131-0101001111122000)
- [advanced_profile](resources--bigip_http_proxy--reference--group-001.md#canonical-2112133030313203-1202233221221323-0231033212122201-1200202122332313-2022030302010022-1221300001120300-3120211033311213-2323103320212230)
- advanced_profile.disable_spec

<a id="canonical-3022212011131102-1301213222303303-1233022013233131-1232111020211133-0022133101010300-1003312110023132-2102010321120310-3330031000333121"></a>

Type: `["object", {}]`. Optional.

Enable this option

Terraform syntax:

```terraform
disable_spec = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3133021223323122-3320012200231021-3301031021103333-1013210113222003-1132233333333102-3313303211213212-1000321311333320-1011120110101032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `advanced_profile.enable_default_profile` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-0101201332222201-3302033021001322-1302101210311331-2210333210320003-0002033231003211-1133020003110203-0200133113131131-0101001111122000)
- [advanced_profile](resources--bigip_http_proxy--reference--group-001.md#canonical-2112133030313203-1202233221221323-0231033212122201-1200202122332313-2022030302010022-1221300001120300-3120211033311213-2323103320212230)
- advanced_profile.enable_default_profile

<a id="canonical-0201313032000320-1323113231232232-2311320230113231-1201223201010012-2330122302202222-1320302221121132-1131200120011200-2100303120132231"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for enable default profile.

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
enable_default_profile = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3223222331220311-2021021213301233-3110010022113211-1113021011323210-1100223221332333-3200220332110002-3301021103103321-3002222130011323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `ddos_profile` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-0101201332222201-3302033021001322-1302101210311331-2210333210320003-0002033231003211-1133020003110203-0200133113131131-0101001111122000)
- ddos_profile

<a id="canonical-3121300011110230-1223302133011302-0222020300333120-3123321130033230-1112032013102131-3000000213222331-3103220322221300-1131001113201213"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for ddos profile.

Additional upstream details:

BIG-IP DDoS Protection Rules.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("disable_ddos_mitigation",
    "enable_ddos_mitigation")}
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
  "x-ves-oneof-field-ddos_mitigation_choice": "[\"disable_ddos_mitigation\",\"enable_ddos_mitigation\"]"
}
```

Terraform syntax:

```terraform
ddos_profile {
  # Configure direct properties listed below.
}
```

<a id="canonical-3100203030011212-3111230021010223-0101011303320332-2210333103331120-0030001332231112-3010003230100330-1231131212232100-0010122313133123"></a>

### Direct properties for `ddos_profile`

- [disable_ddos_mitigation](resources--bigip_http_proxy--reference--group-001.md#canonical-1302201133112223-3100013230232120-1322311122232330-0011201212302201-3313203112300311-0312030100123330-2310313231020013-2322302300131100): complete subsection reference.

- [enable_ddos_mitigation](resources--bigip_http_proxy--reference--group-001.md#canonical-1033233231303030-1112023201021100-0102333202303333-1231203322220103-3121230200123000-2003022222233332-0130133200300322-2330230202020231): complete subsection reference.

<a id="canonical-1302201133112223-3100013230232120-1322311122232330-0011201212302201-3313203112300311-0312030100123330-2310313231020013-2322302300131100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `ddos_profile.disable_ddos_mitigation` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-0101201332222201-3302033021001322-1302101210311331-2210333210320003-0002033231003211-1133020003110203-0200133113131131-0101001111122000)
- [ddos_profile](resources--bigip_http_proxy--reference--group-001.md#canonical-3223222331220311-2021021213301233-3110010022113211-1113021011323210-1100223221332333-3200220332110002-3301021103103321-3002222130011323)
- ddos_profile.disable_ddos_mitigation

<a id="canonical-0002230032332111-2200031133321211-3102201011032320-0210302233031102-2003223021300203-0101232321031110-0120123332100222-0301110303322102"></a>

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
disable_ddos_mitigation = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1033233231303030-1112023201021100-0102333202303333-1231203322220103-3121230200123000-2003022222233332-0130133200300322-2330230202020231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `ddos_profile.enable_ddos_mitigation` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-0101201332222201-3302033021001322-1302101210311331-2210333210320003-0002033231003211-1133020003110203-0200133113131131-0101001111122000)
- [ddos_profile](resources--bigip_http_proxy--reference--group-001.md#canonical-3223222331220311-2021021213301233-3110010022113211-1113021011323210-1100223221332333-3200220332110002-3301021103103321-3002222130011323)
- ddos_profile.enable_ddos_mitigation

<a id="canonical-2122033133230010-1301132100003330-2321002022202200-3320221013033020-1001003213001331-1310003033103132-1030301111020113-0033020103032000"></a>

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
enable_ddos_mitigation = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1120320133323032-2322013030113130-0032022100212120-1002211211233300-1000313031332132-3203200330320300-1100023210101120-0131020312332202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `irules` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-0101201332222201-3302033021001322-1302101210311331-2210333210320003-0002033231003211-1133020003110203-0200133113131131-0101001111122000)
- irules

<a id="canonical-0300010123030301-1030210311032130-3233322003322132-3033321202133020-2323200020113323-1002201233120210-2022122302231312-2333010232311203"></a>

Type: `"object"`. single nested block, Optional.

IRules Configuration for downstream connections.

Receipt-pinned upstream constraints:

```json
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
irules {
  # Configure direct properties listed below.
}
```

<a id="canonical-1022103303333122-2112113311213032-2133313010001200-2032130200201320-2211020122000011-0223111233323112-0110212122022333-2310310332031132"></a>

### Direct properties for `irules`

- [irules](resources--bigip_http_proxy--reference--group-001.md#canonical-1030221032303100-1301132302213012-0012213202330120-3313303203211331-1202022112230230-2223300131130022-0020002023130100-0102102221301112): complete subsection reference.

<a id="canonical-1030221032303100-1301132302213012-0012213202330120-3313303203211331-1202022112230230-2223300131130022-0020002023130100-0102102221301112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `irules.irules` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-0101201332222201-3302033021001322-1302101210311331-2210333210320003-0002033231003211-1133020003110203-0200133113131131-0101001111122000)
- [irules](resources--bigip_http_proxy--reference--group-001.md#canonical-1120320133323032-2322013030113130-0032022100212120-1002211211233300-1000313031332132-3203200330320300-1100023210101120-0131020312332202)
- irules.irules

<a id="canonical-2222000011323303-1231101132120330-3032101203102231-2201033030112020-1113322132330120-1012100132313130-0223301321303332-1332333123300032"></a>

Type: `"object"`. list nested block, Optional.

OPTIONS for attaching iRules to BIG-IP HTTP Proxy.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.RequiredListObjectAttributes("name")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 16,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 16,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
irules {
  # Configure direct properties listed below.
}
```

<a id="canonical-0123223120323212-0212001321213223-2010301003113210-2022210032312031-0320121320233301-3323133011221023-2120130030310312-1032311130331002"></a>

### Direct properties for `irules.irules`

<a id="canonical-2031021110012102-1020100013102311-3123211213202333-3133033312013202-2121112013020033-0123020310102313-3211102123301130-2030020013111323"></a>

#### `irules.irules.name` property

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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

<a id="canonical-2313121021230311-0101123000031022-3020323033202133-0010033330111223-2332033203133003-0100231322220232-3031223311131312-1233301302232033"></a>

<a id="canonical-1122331222031312-2301131103313112-2103322022231212-0000222103031213-0111200102202013-0222110200130330-2310010313333203-3123200213220311"></a>

#### `irules.irules.namespace` property

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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

<a id="canonical-1312211003113032-2233030002223332-3320331113102300-3320211322131000-0210023313023230-1312033330131100-0123210312230330-3023320000133030"></a>

<a id="canonical-1301200221000233-1100313001111011-2020133212121313-3100303022103130-3100113102033100-3101223320323222-0110113300031002-1220002220212031"></a>

#### `irules.irules.tenant` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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

<a id="canonical-0010233021331310-1100023120223111-0232222211022121-1032211023101301-3003123030113013-3230220310130310-0221133011233320-1132331211111231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `lb_algorithm` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-0101201332222201-3302033021001322-1302101210311331-2210333210320003-0002033231003211-1133020003110203-0200133113131131-0101001111122000)
- lb_algorithm

<a id="canonical-1203213130303300-2102000300231232-2201011123000301-1210101120003333-3102303320133200-3211331113202003-3113012130100202-3110132101312213"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for lb algorithm.

Additional upstream details:

Load Balancing Algorithm Type.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-lb_algorithm_choice": "[\"round_robin\"]"
}
```

Terraform syntax:

```terraform
lb_algorithm {
  # Configure direct properties listed below.
}
```

<a id="canonical-3231011102333322-3330302120233013-2121101122223331-0120103211231312-1131323320311312-1203123311310302-1230020333003213-1033131200303010"></a>

### Direct properties for `lb_algorithm`

- [round_robin](resources--bigip_http_proxy--reference--group-001.md#canonical-1032322331030110-3003003122311201-2203010003020020-0033132123302200-2213122211132021-0133023330113310-3102121331113002-0221302122230131): complete subsection reference.

<a id="canonical-1032322331030110-3003003122311201-2203010003020020-0033132123302200-2213122211132021-0133023330113310-3102121331113002-0221302122230131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `lb_algorithm.round_robin` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-0101201332222201-3302033021001322-1302101210311331-2210333210320003-0002033231003211-1133020003110203-0200133113131131-0101001111122000)
- [lb_algorithm](resources--bigip_http_proxy--reference--group-001.md#canonical-0010233021331310-1100023120223111-0232222211022121-1032211023101301-3003123030113013-3230220310130310-0221133011233320-1132331211111231)
- lb_algorithm.round_robin

<a id="canonical-1020321112021311-3233221200032311-0011011203111303-2223303130211232-2221031111110220-2210300121322212-1200032000022223-3121032112121031"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for round robin.

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
round_robin {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3130302100213032-1133033213001002-2331132101322013-0203331122233331-3021113033330230-2313122333001332-0303103133232110-2002132020103202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_pools` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-0101201332222201-3302033021001322-1302101210311331-2210333210320003-0002033231003211-1133020003110203-0200133113131131-0101001111122000)
- origin_pools

<a id="canonical-2333331002003021-3130223021000003-2301230130332012-3101331020300311-2130303003122212-0333033333231300-2133021212030333-0221003231132311"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for origin pools.

Additional upstream details:

List of Origin Pools.

Receipt-pinned upstream constraints:

```json
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
origin_pools {
  # Configure direct properties listed below.
}
```

<a id="canonical-0200101322221231-3313312123100031-3301102301203130-1223313330032002-1230221133300201-3023331000332210-3323120002311333-2033232220102032"></a>

### Direct properties for `origin_pools`

- [pools](resources--bigip_http_proxy--reference--group-001.md#canonical-1023031131101223-0333002000010022-2011021202323002-2121003101012320-0131313002113130-0302023223110202-2203300233110313-3320003203232022): complete subsection reference.

<a id="canonical-1023031131101223-0333002000010022-2011021202323002-2121003101012320-0131313002113130-0302023223110202-2203300233110313-3320003203232022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_pools.pools` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-0101201332222201-3302033021001322-1302101210311331-2210333210320003-0002033231003211-1133020003110203-0200133113131131-0101001111122000)
- [origin_pools](resources--bigip_http_proxy--reference--group-001.md#canonical-3130302100213032-1133033213001002-2331132101322013-0203331122233331-3021113033330230-2313122333001332-0303103133232110-2002132020103202)
- origin_pools.pools

<a id="canonical-3323100232203202-2203032231302200-2233113210023011-1222212333122323-2300100100323012-2013033030123201-3231030123231013-2100302320110102"></a>

Type: `"object"`. list nested block, Optional.

Origin Pools. List of Origin Pools.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.RequiredListObjectAttributes("name")}
```

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minItems": 1,
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
pools {
  # Configure direct properties listed below.
}
```

<a id="canonical-0330231223121222-0230232221331021-3211320122032120-1310311012202023-0203202301212233-1311213301123331-3130233133001121-2212132301133200"></a>

### Direct properties for `origin_pools.pools`

<a id="canonical-2311203003130200-1221121201030232-0021202021213322-0223201102200320-3211101212202221-1210000123303323-3202010310133002-2022301231331213"></a>

#### `origin_pools.pools.name` property

Type: `"string"`. Optional.

Name. Name of the origin pool.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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
    "formatDescription": "DNS-1035 label: must start with a lowercase letter, may contain lowercase alphanumeric and hyphens, must end with alphanumeric",
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

- [origin_servers](resources--bigip_http_proxy--reference--group-001.md#canonical-1312222213103100-0321300110133131-3312031033013332-1010303313020221-3311012303023212-1010020102000132-3130213301103033-0122010331103022): complete subsection reference.

<a id="canonical-2033230213221231-3331312031213203-0110110000100230-1112333311223123-1311000211323100-3333203032112322-3331311121323310-1300330122020302"></a>

<a id="canonical-2111333322311023-3200130202200122-3310031312002031-2331100121303110-3221220100303031-3010230222310121-1310002312031102-1112332001113303"></a>

#### `origin_pools.pools.priority` property

Type: `"number"`. Optional.

Priority of this origin pool, valid only with multiple origin pools. Value of 0 will make the pool
as lowest priority origin pool. When active origin pool is not available, lower priority origin
pools are made active as per the increasing priority.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(0, 32),
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
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minimum": 0,
    "multipleOf": 1
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

<a id="canonical-1111202133231300-1221011120122111-1321032030113101-2232101201021311-3000211123312201-0021032203200113-3310003333301132-0113110103130103"></a>

<a id="canonical-0322100221201010-3102203300123221-2221101232032033-1231112230230020-2021103113123030-0203111321333301-0110321301200102-2202310102021103"></a>

#### `origin_pools.pools.weight` property

Type: `"number"`. Optional.

Weight of this origin pool, valid only with multiple origin pools. Value of 0 will disable the pool.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-1312222213103100-0321300110133131-3312031033013332-1010303313020221-3311012303023212-1010020102000132-3130213301103033-0122010331103022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_pools.pools.origin_servers` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-0101201332222201-3302033021001322-1302101210311331-2210333210320003-0002033231003211-1133020003110203-0200133113131131-0101001111122000)
- [origin_pools](resources--bigip_http_proxy--reference--group-001.md#canonical-3130302100213032-1133033213001002-2331132101322013-0203331122233331-3021113033330230-2313122333001332-0303103133232110-2002132020103202)
- [origin_pools.pools](resources--bigip_http_proxy--reference--group-001.md#canonical-1023031131101223-0333002000010022-2011021202323002-2121003101012320-0131313002113130-0302023223110202-2203300233110313-3320003203232022)
- origin_pools.pools.origin_servers

<a id="canonical-3211323131311301-0133100120330120-0131202011131231-1310232023120113-0020232233021020-0320102303231100-3232230122332012-2030010130011220"></a>

Type: `"object"`. single nested block, Optional.

List of origin Servers for the BIG-IP HTTP Proxy.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("origin_servers"),
  validators.ConflictingObjectAttributes("automatic_port",
    "lb_port"),
  validators.ConflictingObjectAttributes("automatic_port",
    "port"),
  validators.ConflictingObjectAttributes("lb_port",
    "port")}
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
  "x-ves-oneof-field-port_choice": "[\"automatic_port\",\"lb_port\",\"port\"]"
}
```

Terraform syntax:

```terraform
origin_servers {
  # Configure direct properties listed below.
}
```

<a id="canonical-2210120301022120-3023113103011311-2210133133233330-1120323223331301-2311122133120300-0022332333232000-1113233131003212-2323022331021020"></a>

### Direct properties for `origin_pools.pools.origin_servers`

- [automatic_port](resources--bigip_http_proxy--reference--group-001.md#canonical-1333310132310233-1330123223112102-2232131212021210-2003232100032033-1213013313020301-1003021100330001-2121211110220200-3213131231023102): complete subsection reference.

- [health_checks](resources--bigip_http_proxy--reference--group-001.md#canonical-1231223200203212-2223102120201210-0132023001333121-0310321020212013-2303231100301333-3231130313200203-0020112010001202-2200002131110303): complete subsection reference.

- [lb_port](resources--bigip_http_proxy--reference--group-001.md#canonical-0002232232110202-3232011212333210-3202022230001020-0012303133223111-0132230030013013-0022031310113320-1320202122011103-1122210113011222): complete subsection reference.

- [origin_servers](resources--bigip_http_proxy--reference--group-001.md#canonical-3122233331022000-3130002202110102-2203300321323221-1321111230302130-2010113302011332-2020303212023012-2310303023303012-3120020110120320): complete subsection reference.

<a id="canonical-2202121131021302-3301332301103133-1121102033022033-2123212002323211-1233322000022030-3101300213112020-2332222100311212-3213313100222210"></a>

<a id="canonical-3332320321023000-3113113221020133-0131202001212311-1122022103013333-3333012333102202-3321322131202213-3112332203300203-1300133123213222"></a>

#### `origin_pools.pools.origin_servers.port` property

Type: `"number"`. Optional.

Exclusive with \[automatic\_port lb\_port\] Endpoint service is available on this port.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(1, 65535),
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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-1333310132310233-1330123223112102-2232131212021210-2003232100032033-1213013313020301-1003021100330001-2121211110220200-3213131231023102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_pools.pools.origin_servers.automatic_port` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-0101201332222201-3302033021001322-1302101210311331-2210333210320003-0002033231003211-1133020003110203-0200133113131131-0101001111122000)
- [origin_pools](resources--bigip_http_proxy--reference--group-001.md#canonical-3130302100213032-1133033213001002-2331132101322013-0203331122233331-3021113033330230-2313122333001332-0303103133232110-2002132020103202)
- [origin_pools.pools](resources--bigip_http_proxy--reference--group-001.md#canonical-1023031131101223-0333002000010022-2011021202323002-2121003101012320-0131313002113130-0302023223110202-2203300233110313-3320003203232022)
- [origin_pools.pools.origin_servers](resources--bigip_http_proxy--reference--group-001.md#canonical-1312222213103100-0321300110133131-3312031033013332-1010303313020221-3311012303023212-1010020102000132-3130213301103033-0122010331103022)
- origin_pools.pools.origin_servers.automatic_port

<a id="canonical-2320233022123330-1122032100312112-1303112231020000-3233012111110303-0020100313232031-3130210321233101-3302122320232132-3111023213102333"></a>

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
automatic_port = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1231223200203212-2223102120201210-0132023001333121-0310321020212013-2303231100301333-3231130313200203-0020112010001202-2200002131110303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_pools.pools.origin_servers.health_checks` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-0101201332222201-3302033021001322-1302101210311331-2210333210320003-0002033231003211-1133020003110203-0200133113131131-0101001111122000)
- [origin_pools](resources--bigip_http_proxy--reference--group-001.md#canonical-3130302100213032-1133033213001002-2331132101322013-0203331122233331-3021113033330230-2313122333001332-0303103133232110-2002132020103202)
- [origin_pools.pools](resources--bigip_http_proxy--reference--group-001.md#canonical-1023031131101223-0333002000010022-2011021202323002-2121003101012320-0131313002113130-0302023223110202-2203300233110313-3320003203232022)
- [origin_pools.pools.origin_servers](resources--bigip_http_proxy--reference--group-001.md#canonical-1312222213103100-0321300110133131-3312031033013332-1010303313020221-3311012303023212-1010020102000132-3130213301103033-0122010331103022)
- origin_pools.pools.origin_servers.health_checks

<a id="canonical-1331103210122032-1222112002002132-1010202022302101-0102201221133333-3332310221213320-2032310013300102-0003112303121020-3130211302312231"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for health checks.

Additional upstream details:

Origin Server Health Checks.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("health_check",
    "healthy_threshold",
    "interval",
    "timeout",
    "unhealthy_threshold")}
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
health_checks {
  # Configure direct properties listed below.
}
```

<a id="canonical-3213132002031213-0333000231301102-1213101222233102-2321313131212333-1302100012123331-1332223211220111-2213201213221313-0333010031010110"></a>

### Direct properties for `origin_pools.pools.origin_servers.health_checks`

- [health_check](resources--bigip_http_proxy--reference--group-001.md#canonical-2120333231121323-2003031320012101-2302210211013333-0122323100321303-1310022230223303-2202022123002231-0033332311313302-2323130331102303): complete subsection reference.

<a id="canonical-2030311222030320-2233131102220000-2312113031303002-3231023133311200-0010000202111012-3131231302321333-1313003321030220-3223013312310211"></a>

<a id="canonical-0021320100203000-1130010230022202-1130131233121223-2322310032202331-2203002131210013-2111022030130203-3003223111002332-2123222102222002"></a>

#### `origin_pools.pools.origin_servers.health_checks.healthy_threshold` property

Type: `"number"`. Optional.

Number of successful responses before declaring healthy. In other words, this is the number of
healthy health checks required before a host is marked healthy. Note that during startup, only a
single successful health check is required to mark a host healthy.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(1, 16),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 16,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "16"
  }
}
```

<a id="canonical-1020230310301001-1232001301033103-2011320133202203-1212212011200310-2200203211020030-3313020122121230-0103120213000230-2130120102211102"></a>

<a id="canonical-3321021010130100-2300211203012002-3203333112221033-0220130030000010-0013201212231123-1312211131010311-3021102211110321-1233223010210133"></a>

#### `origin_pools.pools.origin_servers.health_checks.interval` property

Type: `"number"`. Optional.

Time interval in seconds between two health check requests.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(1, 600),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 600,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minimum": 1,
    "multipleOf": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "600"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "600"
  }
}
```

<a id="canonical-2033230321000020-3032322202321132-0111000301020320-2123013031331010-1110212311123110-3032122220303103-1320200210312230-2010222011120101"></a>

<a id="canonical-3000123301022003-0310023302131330-0011323212312112-3020220033133332-2220012303303001-3133303310203021-2212223212213333-0212312303022033"></a>

#### `origin_pools.pools.origin_servers.health_checks.timeout` property

Type: `"number"`. Optional.

Timeout in seconds to wait for successful response. In other words, it is the time to wait for a
health check response. If the timeout is reached the health check attempt will be considered a
failure.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(1, 600),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 600,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minimum": 1,
    "multipleOf": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "600"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "600"
  }
}
```

<a id="canonical-0301211311133301-2211031123203031-0233011331213332-2003200221221111-2022312000011011-0013302021213103-0103233302113203-1321300111103122"></a>

<a id="canonical-0120201222011211-0213133131022012-1332133230202303-3023330122021223-3220311113122132-2101332022031013-3133313132332031-0210101000033100"></a>

#### `origin_pools.pools.origin_servers.health_checks.unhealthy_threshold` property

Type: `"number"`. Optional.

Number of failed responses before declaring unhealthy. In other words, this is the number of
unhealthy health checks required before a host is marked unhealthy. Note that for HTTP health check
if a host responds with 503 this threshold is ignored and the host is considered unhealthy
immediately.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(1, 16),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 16,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "16"
  }
}
```

<a id="canonical-2120333231121323-2003031320012101-2302210211013333-0122323100321303-1310022230223303-2202022123002231-0033332311313302-2323130331102303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_pools.pools.origin_servers.health_checks.health_check` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-0101201332222201-3302033021001322-1302101210311331-2210333210320003-0002033231003211-1133020003110203-0200133113131131-0101001111122000)
- [origin_pools](resources--bigip_http_proxy--reference--group-001.md#canonical-3130302100213032-1133033213001002-2331132101322013-0203331122233331-3021113033330230-2313122333001332-0303103133232110-2002132020103202)
- [origin_pools.pools](resources--bigip_http_proxy--reference--group-001.md#canonical-1023031131101223-0333002000010022-2011021202323002-2121003101012320-0131313002113130-0302023223110202-2203300233110313-3320003203232022)
- [origin_pools.pools.origin_servers](resources--bigip_http_proxy--reference--group-001.md#canonical-1312222213103100-0321300110133131-3312031033013332-1010303313020221-3311012303023212-1010020102000132-3130213301103033-0122010331103022)
- [origin_pools.pools.origin_servers.health_checks](resources--bigip_http_proxy--reference--group-001.md#canonical-1231223200203212-2223102120201210-0132023001333121-0310321020212013-2303231100301333-3231130313200203-0020112010001202-2200002131110303)
- origin_pools.pools.origin_servers.health_checks.health_check

<a id="canonical-2133223333122032-2011001010233113-3210103130230021-0310202313132323-0013323110212011-1033223333222302-1201000101002132-3102232032310210"></a>

Type: `"object"`. list nested block, Optional.

List of Health Checks. List of Health Checks.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.ConflictingListObjectAttributes("icmp_health_check",
    "tcp_health_check")}
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
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minItems": 0,
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
    "ves.io.schema.rules.repeated.max_items": "3",
    "ves.io.schema.rules.repeated.min_items": "0",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "3",
    "ves.io.schema.rules.repeated.min_items": "0",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
health_check {
  # Configure direct properties listed below.
}
```

<a id="canonical-1203013122112300-3333201032112200-3133210301003023-2222213012310110-1123312011212032-1020123310233031-2310133302132230-0000003331201111"></a>

### Direct properties for `origin_pools.pools.origin_servers.health_checks.health_check`

- [icmp_health_check](resources--bigip_http_proxy--reference--group-001.md#canonical-3230311221233321-3113131300120012-3312221030103021-1003101031012232-1232320321011133-3123000030020221-3332122231300231-1110331012030203): complete subsection reference.

- [tcp_health_check](resources--bigip_http_proxy--reference--group-001.md#canonical-0112001111130222-0002130202100220-0202302331311333-3332221001111213-1113010103001132-0031221212001123-0221330031031112-3223013011123101): complete subsection reference.

<a id="canonical-3230311221233321-3113131300120012-3312221030103021-1003101031012232-1232320321011133-3123000030020221-3332122231300231-1110331012030203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_pools.pools.origin_servers.health_checks.health_check.icmp_health_check` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-0101201332222201-3302033021001322-1302101210311331-2210333210320003-0002033231003211-1133020003110203-0200133113131131-0101001111122000)
- [origin_pools](resources--bigip_http_proxy--reference--group-001.md#canonical-3130302100213032-1133033213001002-2331132101322013-0203331122233331-3021113033330230-2313122333001332-0303103133232110-2002132020103202)
- [origin_pools.pools](resources--bigip_http_proxy--reference--group-001.md#canonical-1023031131101223-0333002000010022-2011021202323002-2121003101012320-0131313002113130-0302023223110202-2203300233110313-3320003203232022)
- [origin_pools.pools.origin_servers](resources--bigip_http_proxy--reference--group-001.md#canonical-1312222213103100-0321300110133131-3312031033013332-1010303313020221-3311012303023212-1010020102000132-3130213301103033-0122010331103022)
- [origin_pools.pools.origin_servers.health_checks](resources--bigip_http_proxy--reference--group-001.md#canonical-1231223200203212-2223102120201210-0132023001333121-0310321020212013-2303231100301333-3231130313200203-0020112010001202-2200002131110303)
- [origin_pools.pools.origin_servers.health_checks.health_check](resources--bigip_http_proxy--reference--group-001.md#canonical-2120333231121323-2003031320012101-2302210211013333-0122323100321303-1310022230223303-2202022123002231-0033332311313302-2323130331102303)
- origin_pools.pools.origin_servers.health_checks.health_check.icmp_health_check

<a id="canonical-2101212012010021-1201233100032322-0000103131120333-2033212232112011-0332121122201032-1300033333312030-2202301331122333-0223322113030102"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for icmp health check.

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
icmp_health_check = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0112001111130222-0002130202100220-0202302331311333-3332221001111213-1113010103001132-0031221212001123-0221330031031112-3223013011123101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_pools.pools.origin_servers.health_checks.health_check.tcp_health_check` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-0101201332222201-3302033021001322-1302101210311331-2210333210320003-0002033231003211-1133020003110203-0200133113131131-0101001111122000)
- [origin_pools](resources--bigip_http_proxy--reference--group-001.md#canonical-3130302100213032-1133033213001002-2331132101322013-0203331122233331-3021113033330230-2313122333001332-0303103133232110-2002132020103202)
- [origin_pools.pools](resources--bigip_http_proxy--reference--group-001.md#canonical-1023031131101223-0333002000010022-2011021202323002-2121003101012320-0131313002113130-0302023223110202-2203300233110313-3320003203232022)
- [origin_pools.pools.origin_servers](resources--bigip_http_proxy--reference--group-001.md#canonical-1312222213103100-0321300110133131-3312031033013332-1010303313020221-3311012303023212-1010020102000132-3130213301103033-0122010331103022)
- [origin_pools.pools.origin_servers.health_checks](resources--bigip_http_proxy--reference--group-001.md#canonical-1231223200203212-2223102120201210-0132023001333121-0310321020212013-2303231100301333-3231130313200203-0020112010001202-2200002131110303)
- [origin_pools.pools.origin_servers.health_checks.health_check](resources--bigip_http_proxy--reference--group-001.md#canonical-2120333231121323-2003031320012101-2302210211013333-0122323100321303-1310022230223303-2202022123002231-0033332311313302-2323130331102303)
- origin_pools.pools.origin_servers.health_checks.health_check.tcp_health_check

<a id="canonical-3101022020020302-0133012122013101-0300331300221300-0131333120113132-2003023310223011-3102032200110003-1222112322032311-3131210333333131"></a>

Type: `"object"`. single nested block, Optional.

Monitor reports healthy status if UDP connection is successful and response payload matches expected
response pattern.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("expected_response",
    "send_payload")}
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
tcp_health_check {
  # Configure direct properties listed below.
}
```

<a id="canonical-1302303103302212-3202033300301211-1213301202300101-3210213100300011-3101201210013333-3322032311331030-2022203122333320-2103011013002032"></a>

### Direct properties for `origin_pools.pools.origin_servers.health_checks.health_check.tcp_health_check`

<a id="canonical-1122100001101233-1320333212013101-0320120233333010-3321120033021023-3112330322332012-3022330032233120-0211010020223131-2230103033023201"></a>

#### `origin_pools.pools.origin_servers.health_checks.health_check.tcp_health_check.expected_response` property

Type: `"string"`. Optional.

Specifies a regular expression pattern which will be matched against response payload.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(2048),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 2048,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 2048,
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
    "ves.io.schema.rules.string.max_len": "2048"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "2048"
  }
}
```

<a id="canonical-3003212330320123-0111322003203032-2103103333211012-3300321003003303-2111222120331031-0000303010022232-1332230220110101-3031021232112021"></a>

<a id="canonical-3100000121212100-0200231222003320-2310202223301322-3012321022333022-2330301032322331-0222213221313223-3232223202121321-3301113130233022"></a>

#### `origin_pools.pools.origin_servers.health_checks.health_check.tcp_health_check.send_payload` property

Type: `"string"`. Optional.

Send string. Text string sent in the request.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(2048),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 2048,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 2048,
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
    "ves.io.schema.rules.string.max_len": "2048"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "2048"
  }
}
```

<a id="canonical-0002232232110202-3232011212333210-3202022230001020-0012303133223111-0132230030013013-0022031310113320-1320202122011103-1122210113011222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_pools.pools.origin_servers.lb_port` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-0101201332222201-3302033021001322-1302101210311331-2210333210320003-0002033231003211-1133020003110203-0200133113131131-0101001111122000)
- [origin_pools](resources--bigip_http_proxy--reference--group-001.md#canonical-3130302100213032-1133033213001002-2331132101322013-0203331122233331-3021113033330230-2313122333001332-0303103133232110-2002132020103202)
- [origin_pools.pools](resources--bigip_http_proxy--reference--group-001.md#canonical-1023031131101223-0333002000010022-2011021202323002-2121003101012320-0131313002113130-0302023223110202-2203300233110313-3320003203232022)
- [origin_pools.pools.origin_servers](resources--bigip_http_proxy--reference--group-001.md#canonical-1312222213103100-0321300110133131-3312031033013332-1010303313020221-3311012303023212-1010020102000132-3130213301103033-0122010331103022)
- origin_pools.pools.origin_servers.lb_port

<a id="canonical-3202233223233310-1201311120130333-1112100312003213-2102033002011312-0202001200023321-2111323330230120-1022021102001001-0102130001032220"></a>

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
lb_port = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3122233331022000-3130002202110102-2203300321323221-1321111230302130-2010113302011332-2020303212023012-2310303023303012-3120020110120320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_pools.pools.origin_servers.origin_servers` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-0101201332222201-3302033021001322-1302101210311331-2210333210320003-0002033231003211-1133020003110203-0200133113131131-0101001111122000)
- [origin_pools](resources--bigip_http_proxy--reference--group-001.md#canonical-3130302100213032-1133033213001002-2331132101322013-0203331122233331-3021113033330230-2313122333001332-0303103133232110-2002132020103202)
- [origin_pools.pools](resources--bigip_http_proxy--reference--group-001.md#canonical-1023031131101223-0333002000010022-2011021202323002-2121003101012320-0131313002113130-0302023223110202-2203300233110313-3320003203232022)
- [origin_pools.pools.origin_servers](resources--bigip_http_proxy--reference--group-001.md#canonical-1312222213103100-0321300110133131-3312031033013332-1010303313020221-3311012303023212-1010020102000132-3130213301103033-0122010331103022)
- origin_pools.pools.origin_servers.origin_servers

<a id="canonical-0011000100323203-3122231030111012-0303222000311211-3321333133321310-2121010221210121-0033101132030101-0213323113213313-1113220131123200"></a>

Type: `"object"`. list nested block, Optional.

List of Origin Servers. List of origin servers for Proxy.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.ConflictingListObjectAttributes("k8s_service",
    "private_ip"),
  validators.ConflictingListObjectAttributes("k8s_service",
    "public_ip"),
  validators.ConflictingListObjectAttributes("k8s_service",
    "public_name"),
  validators.ConflictingListObjectAttributes("private_ip",
    "public_ip"),
  validators.ConflictingListObjectAttributes("private_ip",
    "public_name"),
  validators.ConflictingListObjectAttributes("public_ip",
    "public_name")}
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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
origin_servers {
  # Configure direct properties listed below.
}
```

<a id="canonical-1122013133210032-3303332303000110-2203133222101113-1202331123112222-3103213101320323-3100003312131220-2320300212112232-2012101033320013"></a>

### Direct properties for `origin_pools.pools.origin_servers.origin_servers`

- [k8s_service](resources--bigip_http_proxy--reference--group-001.md#canonical-0211202201221121-3001000133010112-1131111200013313-3122130332300332-1020212223330013-1311100313303320-2312112110031023-1120330330033233): complete subsection reference.

- [private_ip](resources--bigip_http_proxy--reference--group-002.md#canonical-3232131333111312-0323012101311031-0132212133331330-0022022302222023-3301101031131030-2203330111111210-0202012220323010-3232101203203320): complete subsection reference.

- [public_ip](resources--bigip_http_proxy--reference--group-002.md#canonical-2200332312202210-0032022133011112-0321011230001311-0210031123310011-2312030333320211-1331331202212220-0001233120122303-0011232022323133): complete subsection reference.

- [public_name](resources--bigip_http_proxy--reference--group-002.md#canonical-2002330123011211-1100203020020233-2200222203213100-0323123022123023-2303301212202330-1220032321201102-3002033230123013-0330123233202110): complete subsection reference.

<a id="canonical-0211202201221121-3001000133010112-1131111200013313-3122130332300332-1020212223330013-1311100313303320-2312112110031023-1120330330033233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_pools.pools.origin_servers.origin_servers.k8s_service` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-0101201332222201-3302033021001322-1302101210311331-2210333210320003-0002033231003211-1133020003110203-0200133113131131-0101001111122000)
- [origin_pools](resources--bigip_http_proxy--reference--group-001.md#canonical-3130302100213032-1133033213001002-2331132101322013-0203331122233331-3021113033330230-2313122333001332-0303103133232110-2002132020103202)
- [origin_pools.pools](resources--bigip_http_proxy--reference--group-001.md#canonical-1023031131101223-0333002000010022-2011021202323002-2121003101012320-0131313002113130-0302023223110202-2203300233110313-3320003203232022)
- [origin_pools.pools.origin_servers](resources--bigip_http_proxy--reference--group-001.md#canonical-1312222213103100-0321300110133131-3312031033013332-1010303313020221-3311012303023212-1010020102000132-3130213301103033-0122010331103022)
- [origin_pools.pools.origin_servers.origin_servers](resources--bigip_http_proxy--reference--group-001.md#canonical-3122233331022000-3130002202110102-2203300321323221-1321111230302130-2010113302011332-2020303212023012-2310303023303012-3120020110120320)
- origin_pools.pools.origin_servers.origin_servers.k8s_service

<a id="canonical-0003000021001331-3303012123010331-3012201100312110-2312003113013033-0231133022022023-2233133132100330-3201322321302220-2033002231323330"></a>

Type: `"object"`. single nested block, Optional.

Specify origin server with K8s service name and site information.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("inside_network",
    "outside_network"),
  validators.ConflictingObjectAttributes("inside_network",
    "vk8s_networks"),
  validators.ConflictingObjectAttributes("outside_network",
    "vk8s_networks")}
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
  "x-ves-oneof-field-network_choice": "[\"inside_network\",\"outside_network\",\"vk8s_networks\"]",
  "x-ves-oneof-field-service_info": "[\"service_name\"]"
}
```

Terraform syntax:

```terraform
k8s_service {
  # Configure direct properties listed below.
}
```

<a id="canonical-3300201331203120-3033122101101122-1102112222032033-0232223302211313-0313200221302011-2201313001123031-2212332100223112-1301231203322310"></a>

### Direct properties for `origin_pools.pools.origin_servers.origin_servers.k8s_service`

- [inside_network](resources--bigip_http_proxy--reference--group-001.md#canonical-0001201011312002-0230122321130200-2323211101112333-1002023023021030-1230100032032310-0330303232023002-3102223313011110-1232333132310000): complete subsection reference.

- [outside_network](resources--bigip_http_proxy--reference--group-001.md#canonical-0032102132120233-2002100110212202-0331323313232221-2000013303021212-0332300333030103-0200033233310011-0220033321023302-0132210302231220): complete subsection reference.

<a id="canonical-1311112220202330-0001203101011022-1232022322123302-2330132112023231-2021210011110031-0220103113100210-0022322100030033-0202013213000221"></a>

<a id="canonical-3103103202223030-3122003323010020-3111202020121120-2023133111211200-0301322132321022-2223313113233001-2101123300313230-0012103100120220"></a>

#### `origin_pools.pools.origin_servers.origin_servers.k8s_service.protocol` property

Type: `"string"`. Optional.

\[Enum: PROTOCOL\_TCP|PROTOCOL\_UDP\] Type of protocol - PROTOCOL\_TCP: TCP - PROTOCOL\_UDP: UDP.
Possible values are \`PROTOCOL\_TCP\`, \`PROTOCOL\_UDP\`. Defaults to \`PROTOCOL\_TCP\`.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["PROTOCOL_TCP","PROTOCOL_UDP"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
Validators: []validator.String{
  stringvalidator.OneOf("PROTOCOL_TCP",
    "PROTOCOL_UDP"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "PROTOCOL_TCP",
  "enum": [
    "PROTOCOL_TCP",
    "PROTOCOL_UDP"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2111201023220313-2111231110200032-3320333300112123-0203311311301130-3330002311221323-2232313212012301-1322230112203131-3002022322230110"></a>

<a id="canonical-0021110203301213-0032113212013033-0003223132203131-2132113223102231-3203131222012022-0020122313101112-0021032021113123-3300303213121020"></a>

#### `origin_pools.pools.origin_servers.origin_servers.k8s_service.service_name` property

Type: `"string"`. Optional.

Exclusive with \[\] K8s service name of the origin server will be listed, including the namespace
and cluster-ID. For vK8s services, you need to enter a string with the format
servicename.namespace:example-namespace'frontend', namespace is 'speedtest' and cluster-ID is
'prod', then you will enter..

Additional upstream details:

For vK8s services, you need to enter a string with the format
servicename.namespace:example-namespace"frontend", namespace is "speedtest" and cluster-ID is
"prod", then you will enter "frontend.speedtest:prod". Both namespace and cluster-ID are optional.

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
    "ves.io.schema.rules.string.ves_service_namespace_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ves_service_namespace_name": "true"
  }
}
```

- [site_locator](resources--bigip_http_proxy--reference--group-001.md#canonical-2300311000201002-3301211203010223-3023013122001213-1000230130032212-3031331011220210-0102001321020313-2130001321333332-1320131200011333): complete subsection reference.

- [snat_pool](resources--bigip_http_proxy--reference--group-001.md#canonical-1311211321313201-3321333332130333-1111232100201310-0132003101311310-2303103123122102-2211202330330320-0120121003102110-0330330133201010): complete subsection reference.

- [vk8s_networks](resources--bigip_http_proxy--reference--group-002.md#canonical-0011302332230132-2030011302213232-3301331313023332-1231222203031121-0301012321311130-2330300102211320-1332010123131110-3020003033110321): complete subsection reference.

<a id="canonical-0001201011312002-0230122321130200-2323211101112333-1002023023021030-1230100032032310-0330303232023002-3102223313011110-1232333132310000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_pools.pools.origin_servers.origin_servers.k8s_service.inside_network` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-0101201332222201-3302033021001322-1302101210311331-2210333210320003-0002033231003211-1133020003110203-0200133113131131-0101001111122000)
- [origin_pools](resources--bigip_http_proxy--reference--group-001.md#canonical-3130302100213032-1133033213001002-2331132101322013-0203331122233331-3021113033330230-2313122333001332-0303103133232110-2002132020103202)
- [origin_pools.pools](resources--bigip_http_proxy--reference--group-001.md#canonical-1023031131101223-0333002000010022-2011021202323002-2121003101012320-0131313002113130-0302023223110202-2203300233110313-3320003203232022)
- [origin_pools.pools.origin_servers](resources--bigip_http_proxy--reference--group-001.md#canonical-1312222213103100-0321300110133131-3312031033013332-1010303313020221-3311012303023212-1010020102000132-3130213301103033-0122010331103022)
- [origin_pools.pools.origin_servers.origin_servers](resources--bigip_http_proxy--reference--group-001.md#canonical-3122233331022000-3130002202110102-2203300321323221-1321111230302130-2010113302011332-2020303212023012-2310303023303012-3120020110120320)
- [origin_pools.pools.origin_servers.origin_servers.k8s_service](resources--bigip_http_proxy--reference--group-001.md#canonical-0211202201221121-3001000133010112-1131111200013313-3122130332300332-1020212223330013-1311100313303320-2312112110031023-1120330330033233)
- origin_pools.pools.origin_servers.origin_servers.k8s_service.inside_network

<a id="canonical-2002213121110021-1303111211120211-1223013023211120-2030021223203320-3101103302302131-3122221023311200-3111331202202003-1311203001233000"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for inside network.

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
inside_network = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0032102132120233-2002100110212202-0331323313232221-2000013303021212-0332300333030103-0200033233310011-0220033321023302-0132210302231220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_pools.pools.origin_servers.origin_servers.k8s_service.outside_network` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-0101201332222201-3302033021001322-1302101210311331-2210333210320003-0002033231003211-1133020003110203-0200133113131131-0101001111122000)
- [origin_pools](resources--bigip_http_proxy--reference--group-001.md#canonical-3130302100213032-1133033213001002-2331132101322013-0203331122233331-3021113033330230-2313122333001332-0303103133232110-2002132020103202)
- [origin_pools.pools](resources--bigip_http_proxy--reference--group-001.md#canonical-1023031131101223-0333002000010022-2011021202323002-2121003101012320-0131313002113130-0302023223110202-2203300233110313-3320003203232022)
- [origin_pools.pools.origin_servers](resources--bigip_http_proxy--reference--group-001.md#canonical-1312222213103100-0321300110133131-3312031033013332-1010303313020221-3311012303023212-1010020102000132-3130213301103033-0122010331103022)
- [origin_pools.pools.origin_servers.origin_servers](resources--bigip_http_proxy--reference--group-001.md#canonical-3122233331022000-3130002202110102-2203300321323221-1321111230302130-2010113302011332-2020303212023012-2310303023303012-3120020110120320)
- [origin_pools.pools.origin_servers.origin_servers.k8s_service](resources--bigip_http_proxy--reference--group-001.md#canonical-0211202201221121-3001000133010112-1131111200013313-3122130332300332-1020212223330013-1311100313303320-2312112110031023-1120330330033233)
- origin_pools.pools.origin_servers.origin_servers.k8s_service.outside_network

<a id="canonical-2113320321333101-1300030113323322-3220012031320010-1213111002021221-0102012203023133-0333012103302102-0310032001311232-3230322021020011"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for outside network.

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
outside_network = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2300311000201002-3301211203010223-3023013122001213-1000230130032212-3031331011220210-0102001321020313-2130001321333332-1320131200011333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_pools.pools.origin_servers.origin_servers.k8s_service.site_locator` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-0101201332222201-3302033021001322-1302101210311331-2210333210320003-0002033231003211-1133020003110203-0200133113131131-0101001111122000)
- [origin_pools](resources--bigip_http_proxy--reference--group-001.md#canonical-3130302100213032-1133033213001002-2331132101322013-0203331122233331-3021113033330230-2313122333001332-0303103133232110-2002132020103202)
- [origin_pools.pools](resources--bigip_http_proxy--reference--group-001.md#canonical-1023031131101223-0333002000010022-2011021202323002-2121003101012320-0131313002113130-0302023223110202-2203300233110313-3320003203232022)
- [origin_pools.pools.origin_servers](resources--bigip_http_proxy--reference--group-001.md#canonical-1312222213103100-0321300110133131-3312031033013332-1010303313020221-3311012303023212-1010020102000132-3130213301103033-0122010331103022)
- [origin_pools.pools.origin_servers.origin_servers](resources--bigip_http_proxy--reference--group-001.md#canonical-3122233331022000-3130002202110102-2203300321323221-1321111230302130-2010113302011332-2020303212023012-2310303023303012-3120020110120320)
- [origin_pools.pools.origin_servers.origin_servers.k8s_service](resources--bigip_http_proxy--reference--group-001.md#canonical-0211202201221121-3001000133010112-1131111200013313-3122130332300332-1020212223330013-1311100313303320-2312112110031023-1120330330033233)
- origin_pools.pools.origin_servers.origin_servers.k8s_service.site_locator

<a id="canonical-2223011211012000-2322312131103323-1222033230230203-2330032121301233-1311031202121131-0312232021330222-0222012331130132-3323303131200101"></a>

Type: `"object"`. single nested block, Optional.

This message defines a reference to a site or virtual site object.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("site",
    "virtual_site")}
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
  "x-ves-oneof-field-choice": "[\"site\",\"virtual_site\"]"
}
```

Terraform syntax:

```terraform
site_locator {
  # Configure direct properties listed below.
}
```

<a id="canonical-1032222030113100-0130223003210101-2320000112131020-3021022222231213-1332322210321331-0213103110120030-2031013200100131-2221002002232121"></a>

### Direct properties for `origin_pools.pools.origin_servers.origin_servers.k8s_service.site_locator`

- [site](resources--bigip_http_proxy--reference--group-001.md#canonical-3013112021001103-2320000033322023-3121201132111010-3232310300320100-1223100001232122-3131321002322230-1213011303111222-0101333302023011): complete subsection reference.

- [virtual_site](resources--bigip_http_proxy--reference--group-001.md#canonical-1320333333202130-1021221312321213-1301302001032112-3222031210310132-0101202121033001-3231022133033030-1110200030203123-3321010023123213): complete subsection reference.

<a id="canonical-3013112021001103-2320000033322023-3121201132111010-3232310300320100-1223100001232122-3131321002322230-1213011303111222-0101333302023011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_pools.pools.origin_servers.origin_servers.k8s_service.site_locator.site` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-0101201332222201-3302033021001322-1302101210311331-2210333210320003-0002033231003211-1133020003110203-0200133113131131-0101001111122000)
- [origin_pools](resources--bigip_http_proxy--reference--group-001.md#canonical-3130302100213032-1133033213001002-2331132101322013-0203331122233331-3021113033330230-2313122333001332-0303103133232110-2002132020103202)
- [origin_pools.pools](resources--bigip_http_proxy--reference--group-001.md#canonical-1023031131101223-0333002000010022-2011021202323002-2121003101012320-0131313002113130-0302023223110202-2203300233110313-3320003203232022)
- [origin_pools.pools.origin_servers](resources--bigip_http_proxy--reference--group-001.md#canonical-1312222213103100-0321300110133131-3312031033013332-1010303313020221-3311012303023212-1010020102000132-3130213301103033-0122010331103022)
- [origin_pools.pools.origin_servers.origin_servers](resources--bigip_http_proxy--reference--group-001.md#canonical-3122233331022000-3130002202110102-2203300321323221-1321111230302130-2010113302011332-2020303212023012-2310303023303012-3120020110120320)
- [origin_pools.pools.origin_servers.origin_servers.k8s_service](resources--bigip_http_proxy--reference--group-001.md#canonical-0211202201221121-3001000133010112-1131111200013313-3122130332300332-1020212223330013-1311100313303320-2312112110031023-1120330330033233)
- [origin_pools.pools.origin_servers.origin_servers.k8s_service.site_locator](resources--bigip_http_proxy--reference--group-001.md#canonical-2300311000201002-3301211203010223-3023013122001213-1000230130032212-3031331011220210-0102001321020313-2130001321333332-1320131200011333)
- origin_pools.pools.origin_servers.origin_servers.k8s_service.site_locator.site

<a id="canonical-2221311223231120-2102313010233223-0303131021022032-0221030001022012-3013302312213122-1302002131311000-1212310123030322-2203000202021023"></a>

Type: `"object"`. single nested block, Optional.

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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
site {
  # Configure direct properties listed below.
}
```

<a id="canonical-1113232101113302-3301232300200203-0333120200300210-3333130202013001-0302032220032211-1032013100313003-1231212133131322-3301111110003330"></a>

### Direct properties for `origin_pools.pools.origin_servers.origin_servers.k8s_service.site_locator.site`

<a id="canonical-0023212223030121-2121033320111111-1020232321210030-3302201231221030-0011323010121230-2100112130221130-2113303101300310-3300220302322003"></a>

#### `origin_pools.pools.origin_servers.origin_servers.k8s_service.site_locator.site.name` property

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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

<a id="canonical-0101113020333213-2223222320111301-2012033213123102-0332201113030223-0003210102103111-0322130211302220-2313311212133331-3133112303211210"></a>

<a id="canonical-3030102133110333-2111230232332210-1103330103201031-1001030112300203-2222003330311031-0200320212313000-3233120131313132-0232132313312201"></a>

#### `origin_pools.pools.origin_servers.origin_servers.k8s_service.site_locator.site.namespace` property

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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

<a id="canonical-3310111322233202-0013031020111122-3011122230303030-2320102023300212-3331221300201101-1031113132122302-1330110122021130-3032321220100321"></a>

<a id="canonical-2200033132211201-0012213122021122-2122321302022323-1130110020013232-2111231213331331-2320311130332013-1102111221311203-1332313200100011"></a>

#### `origin_pools.pools.origin_servers.origin_servers.k8s_service.site_locator.site.tenant` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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

<a id="canonical-1320333333202130-1021221312321213-1301302001032112-3222031210310132-0101202121033001-3231022133033030-1110200030203123-3321010023123213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_pools.pools.origin_servers.origin_servers.k8s_service.site_locator.virtual_site` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-0101201332222201-3302033021001322-1302101210311331-2210333210320003-0002033231003211-1133020003110203-0200133113131131-0101001111122000)
- [origin_pools](resources--bigip_http_proxy--reference--group-001.md#canonical-3130302100213032-1133033213001002-2331132101322013-0203331122233331-3021113033330230-2313122333001332-0303103133232110-2002132020103202)
- [origin_pools.pools](resources--bigip_http_proxy--reference--group-001.md#canonical-1023031131101223-0333002000010022-2011021202323002-2121003101012320-0131313002113130-0302023223110202-2203300233110313-3320003203232022)
- [origin_pools.pools.origin_servers](resources--bigip_http_proxy--reference--group-001.md#canonical-1312222213103100-0321300110133131-3312031033013332-1010303313020221-3311012303023212-1010020102000132-3130213301103033-0122010331103022)
- [origin_pools.pools.origin_servers.origin_servers](resources--bigip_http_proxy--reference--group-001.md#canonical-3122233331022000-3130002202110102-2203300321323221-1321111230302130-2010113302011332-2020303212023012-2310303023303012-3120020110120320)
- [origin_pools.pools.origin_servers.origin_servers.k8s_service](resources--bigip_http_proxy--reference--group-001.md#canonical-0211202201221121-3001000133010112-1131111200013313-3122130332300332-1020212223330013-1311100313303320-2312112110031023-1120330330033233)
- [origin_pools.pools.origin_servers.origin_servers.k8s_service.site_locator](resources--bigip_http_proxy--reference--group-001.md#canonical-2300311000201002-3301211203010223-3023013122001213-1000230130032212-3031331011220210-0102001321020313-2130001321333332-1320131200011333)
- origin_pools.pools.origin_servers.origin_servers.k8s_service.site_locator.virtual_site

<a id="canonical-3331313230220333-1330332111221023-2032211002020313-2213003310003333-0021133331231100-1211121001220210-2222020210232331-1330112331332323"></a>

Type: `"object"`. single nested block, Optional.

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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
virtual_site {
  # Configure direct properties listed below.
}
```

<a id="canonical-3331320233210030-3030020312213331-3011133201211222-1232231103323130-3233133120123033-2223210023233111-0110003011023111-2203320122110330"></a>

### Direct properties for `origin_pools.pools.origin_servers.origin_servers.k8s_service.site_locator.virtual_site`

<a id="canonical-0223002232003220-2210132010112113-0331021110020110-3002203333312200-3333222311022113-2232320120300303-2233122333223103-1220331323020100"></a>

#### `origin_pools.pools.origin_servers.origin_servers.k8s_service.site_locator.virtual_site.name` property

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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

<a id="canonical-3122222102303000-0133132113001303-3000232332020011-3112011012121002-2232003330133030-2321213302201020-1321123321002033-2111330030031312"></a>

<a id="canonical-3032301301120303-1310112111103320-2100322122203202-1211001011013213-3133320200032213-3300002333220333-1023332033031100-0123030301120101"></a>

#### `origin_pools.pools.origin_servers.origin_servers.k8s_service.site_locator.virtual_site.namespace` property

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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

<a id="canonical-0021231230231113-1300211321002230-0102000110202102-0010311320203013-3111220002022030-2302212302301112-0012223321113332-3112301123120011"></a>

<a id="canonical-0012223303212012-2221303312023031-3121332133020231-0322231020010130-0112330203030021-1210301122300110-3102311300233333-2013111212032310"></a>

#### `origin_pools.pools.origin_servers.origin_servers.k8s_service.site_locator.virtual_site.tenant` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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

<a id="canonical-1311211321313201-3321333332130333-1111232100201310-0132003101311310-2303103123122102-2211202330330320-0120121003102110-0330330133201010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_pools.pools.origin_servers.origin_servers.k8s_service.snat_pool` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-0101201332222201-3302033021001322-1302101210311331-2210333210320003-0002033231003211-1133020003110203-0200133113131131-0101001111122000)
- [origin_pools](resources--bigip_http_proxy--reference--group-001.md#canonical-3130302100213032-1133033213001002-2331132101322013-0203331122233331-3021113033330230-2313122333001332-0303103133232110-2002132020103202)
- [origin_pools.pools](resources--bigip_http_proxy--reference--group-001.md#canonical-1023031131101223-0333002000010022-2011021202323002-2121003101012320-0131313002113130-0302023223110202-2203300233110313-3320003203232022)
- [origin_pools.pools.origin_servers](resources--bigip_http_proxy--reference--group-001.md#canonical-1312222213103100-0321300110133131-3312031033013332-1010303313020221-3311012303023212-1010020102000132-3130213301103033-0122010331103022)
- [origin_pools.pools.origin_servers.origin_servers](resources--bigip_http_proxy--reference--group-001.md#canonical-3122233331022000-3130002202110102-2203300321323221-1321111230302130-2010113302011332-2020303212023012-2310303023303012-3120020110120320)
- [origin_pools.pools.origin_servers.origin_servers.k8s_service](resources--bigip_http_proxy--reference--group-001.md#canonical-0211202201221121-3001000133010112-1131111200013313-3122130332300332-1020212223330013-1311100313303320-2312112110031023-1120330330033233)
- origin_pools.pools.origin_servers.origin_servers.k8s_service.snat_pool

<a id="canonical-1210031012130200-3003222231330333-0133200303022131-1223223112102113-3210213223112122-1331002330012323-3032111213321023-3022112132102131"></a>

Type: `"object"`. single nested block, Optional.

SNAT Pool. SNAT Pool configuration.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("no_snat_pool",
    "snat_pool")}
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
  "x-ves-oneof-field-snat_pool_choice": "[\"no_snat_pool\",\"snat_pool\"]"
}
```

Terraform syntax:

```terraform
snat_pool {
  # Configure direct properties listed below.
}
```

<a id="canonical-2203113031322030-2330022102103202-0211000330230112-2321023021232000-2132221020212300-2323003202003223-1033331032231323-1223000222120121"></a>

### Direct properties for `origin_pools.pools.origin_servers.origin_servers.k8s_service.snat_pool`

- [no_snat_pool](resources--bigip_http_proxy--reference--group-001.md#canonical-1120300203220233-1323222220312211-0331031300103232-1032020031220021-0122233111303312-2300212310302333-1111201311313111-3122313000013000): complete subsection reference.

- [snat_pool](resources--bigip_http_proxy--reference--group-002.md#canonical-0032212100033122-2311302103202323-1103111310103113-0223312130120302-3012020303232232-2120331002202032-0233210232332320-1011300123211203): complete subsection reference.

<a id="canonical-1120300203220233-1323222220312211-0331031300103232-1032020031220021-0122233111303312-2300212310302333-1111201311313111-3122313000013000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_pools.pools.origin_servers.origin_servers.k8s_service.snat_pool.no_snat_pool` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-0101201332222201-3302033021001322-1302101210311331-2210333210320003-0002033231003211-1133020003110203-0200133113131131-0101001111122000)
- [origin_pools](resources--bigip_http_proxy--reference--group-001.md#canonical-3130302100213032-1133033213001002-2331132101322013-0203331122233331-3021113033330230-2313122333001332-0303103133232110-2002132020103202)
- [origin_pools.pools](resources--bigip_http_proxy--reference--group-001.md#canonical-1023031131101223-0333002000010022-2011021202323002-2121003101012320-0131313002113130-0302023223110202-2203300233110313-3320003203232022)
- [origin_pools.pools.origin_servers](resources--bigip_http_proxy--reference--group-001.md#canonical-1312222213103100-0321300110133131-3312031033013332-1010303313020221-3311012303023212-1010020102000132-3130213301103033-0122010331103022)
- [origin_pools.pools.origin_servers.origin_servers](resources--bigip_http_proxy--reference--group-001.md#canonical-3122233331022000-3130002202110102-2203300321323221-1321111230302130-2010113302011332-2020303212023012-2310303023303012-3120020110120320)
- [origin_pools.pools.origin_servers.origin_servers.k8s_service](resources--bigip_http_proxy--reference--group-001.md#canonical-0211202201221121-3001000133010112-1131111200013313-3122130332300332-1020212223330013-1311100313303320-2312112110031023-1120330330033233)
- [origin_pools.pools.origin_servers.origin_servers.k8s_service.snat_pool](resources--bigip_http_proxy--reference--group-001.md#canonical-1311211321313201-3321333332130333-1111232100201310-0132003101311310-2303103123122102-2211202330330320-0120121003102110-0330330133201010)
- origin_pools.pools.origin_servers.origin_servers.k8s_service.snat_pool.no_snat_pool

<a id="canonical-2123321133201330-1101022110010022-2212000130131300-3133310211302131-0030133321202301-2022230213131132-0120023221132230-3102120020000311"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for no snat pool.

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
no_snat_pool = {}
```

This is an empty object or choice marker. It has no direct properties.
