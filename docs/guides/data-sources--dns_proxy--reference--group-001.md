---
page_title: "xcsh_dns_proxy reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_dns_proxy reference."
---

# xcsh_dns_proxy reference

<a id="canonical-1133330113022100-3110333211021121-3111232100100000-3311232321023320-3202221331130133-0013122201302330-1103000312200021-1203111112100020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0212301112313200-3222110320321000-1110201202131331-0031113221323210-0332333203003022-3013112033201022-0232333021221112-0210322231303203"></a>

## Property reference — Property reference / 223330232111 / 2

Breadcrumbs:

- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-1122013202133033-3201201122310010-3303231012021212-2123120003213300-3130001133310203-2011013122222333-0130313333122120-0323311233013112)
- Property reference

<a id="canonical-1332302113212130-3232031200312200-2113112330011231-2330130101102301-1200101121201013-1303112233302030-1321113022223130-2302220100220020"></a>

## Direct properties — Property reference / 223330232111 / 3

<a id="canonical-0021233123110303-1100220000120033-1213023003120311-2201011302130022-3220320230202030-0331213230110333-3022303301010000-0113032331212032"></a>

<a id="canonical-3112132300203003-3101201230033030-0312133301300110-3223321212101331-2220203120131323-2221231133103313-3303020111113321-3323130131132003"></a>

## annotations property — Property reference / 223330232111 / 4

Type: `["map", "string"]`. Computed.

Annotations applied to this resource.

Upstream description:

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

- [cache_profile](data-sources--dns_proxy--reference--group-001.md#canonical-2123200303011300-0332210102202000-2233020323203102-0102011230031003-3000331233003002-2033021112321102-1333122002121100-0202123300222123): complete subsection reference.

- [ddos_profile](data-sources--dns_proxy--reference--group-001.md#canonical-0232013023122233-1012213323101023-2001112012330030-3013030103202023-2011131011012030-2021102202320120-0100103221220130-0320000200033100): complete subsection reference.

<a id="canonical-2030300032212210-3012122021002300-0232211130222320-1322301023120310-1331201311011300-3233220133022001-1013300112322300-1202302212122133"></a>

<a id="canonical-0131123321213022-3322111110033201-2113210131121230-3122302103132323-2002122300222202-1203321130223210-0233231011022220-3223020021302211"></a>

## description property — Property reference / 223330232111 / 5

Type: `"string"`. Computed.

Description of the DNSProxy.

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

<a id="canonical-2203303302232133-1321201322111312-2120233310331030-1221231320222231-2120221032123330-1310211301031032-0132120111301302-0001123212302130"></a>

<a id="canonical-2323320120102023-3011203103022333-2213031012113301-1002323033312211-3320222112330001-2231323003333321-2203333320101022-3310032321000020"></a>

## ID property — Property reference / 223330232111 / 6

Type: `"string"`. Computed.

Unique identifier for the resource.

- [irules](data-sources--dns_proxy--reference--group-001.md#canonical-1323010111003201-0303002312013010-3103100312030112-0020032232131232-1312132221212131-1312102001102023-3223132230233333-1312111033010102): complete subsection reference.

<a id="canonical-1033202112132120-1000201200221111-2002121313311212-2110002011120232-0102201031322022-3123012033101020-0203203201302121-1321033030123110"></a>

<a id="canonical-3203100133311023-2031212233332220-1210020122311320-0213201201021131-1001023220123003-2133131322030212-3033030321300213-0011100121001120"></a>

## labels property — Property reference / 223330232111 / 7

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

- [lb_algorithm](data-sources--dns_proxy--reference--group-001.md#canonical-2013211100010010-0023312102121310-3323002033212331-3103023022222212-0210130112023200-0302223111210332-0203211111130203-1233001021233013): complete subsection reference.

<a id="canonical-0331030223323232-0131121113130002-3011010202233303-0211222100332101-1010311220010101-0112232303322302-3002312212222211-0300121000213330"></a>

<a id="canonical-2331123002233131-2202310332010131-0222032022210033-1013030221310220-2000102311000100-3312311110313223-0201211022000013-1200312202021221"></a>

## name property — Property reference / 223330232111 / 8

Type: `"string"`. Required.

Name of the DNSProxy.

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

<a id="canonical-1220331021231011-1023322332200031-0320023103322213-0111120132201313-2210133013310020-2113232032220100-1323301313012233-2230021323030230"></a>

<a id="canonical-1211001111200022-2131211023023322-1022002022333013-1302030322013303-0331303111111212-2011320331300200-0033133311033322-2302011111112323"></a>

## namespace property — Property reference / 223330232111 / 9

Type: `"string"`. Optional, Computed.

Namespace where the DNSProxy exists.

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

- [origin_servers](data-sources--dns_proxy--reference--group-001.md#canonical-3102221313302120-2300133213200331-0320001332113013-1100133203331131-3231110233212130-1312212033030000-3302023133321213-2023021333201010): complete subsection reference.

- [protocol_inspection](data-sources--dns_proxy--reference--group-001.md#canonical-1222310103202331-3022213131203002-0003133222013103-3321032133113232-1300333320113032-0110233333210122-2321331132232223-2213213103001030): complete subsection reference.

- [proxy_advertisement](data-sources--dns_proxy--reference--group-001.md#canonical-1233321320320303-1032110303233312-2131113231101333-3032323203221001-0110213212120103-3322121223310111-0201121112313233-2310233121220133): complete subsection reference.

<a id="canonical-1321320200212302-1122131020300010-0003132323300200-2230221011020133-0113211213013320-0312212000110013-0000213213122003-1103230312232020"></a>

<a id="canonical-3321111102030100-1330223120303230-1212301322223002-0333123332210313-3001303201113132-2320000201032312-1333031300100033-0212232231111332"></a>

## transport_type property — Property reference / 223330232111 / 10

Type: `"string"`. Computed.

\[Enum: UDP|TCP|BothTCPAndUDP\] Transport Type - UDP: UDP - TCP: TCP - BothTCPAndUDP: Both TCP and
UDP. Possible values are \`UDP\`, \`TCP\`, \`BothTCPAndUDP\`. Defaults to \`UDP\`.

Upstream description:

Transport Type

&#8203;- UDP: UDP

&#8203;- TCP: TCP

&#8203;- BothTCPAndUDP: Both TCP and UDP.

Receipt-pinned upstream constraints:

```json
{
  "default": "UDP",
  "enum": [
    "UDP",
    "TCP",
    "BothTCPAndUDP"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2323000222121212-0113200022223203-0322031313331220-3002310010230103-2101111331132231-3330021132023001-3103122203010323-0123220002303122"></a>

## All schema paths — Property reference / 223330232111 / 11

Each exact path has one authoritative reference destination. Collection element indices are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](data-sources--dns_proxy--reference--group-001.md#canonical-0021233123110303-1100220000120033-1213023003120311-2201011302130022-3220320230202030-0331213230110333-3022303301010000-0113032331212032) |
| `cache_profile` | [cache_profile](data-sources--dns_proxy--reference--group-001.md#canonical-3321112031013310-2302300123101133-1301002112323322-2110120312023332-2330210320100011-2232131120222333-2022230120011331-1031232000230020) |
| `cache_profile.cache_size` | [cache_profile.cache_size](data-sources--dns_proxy--reference--group-001.md#canonical-1003330202301011-1111132331000213-1330120332111201-2023130232031023-1312233321001131-3112203203213101-2100313211302120-2333230200121303) |
| `cache_profile.disable_cache_profile` | [cache_profile.disable_cache_profile](data-sources--dns_proxy--reference--group-001.md#canonical-2233313110000322-3200300023100220-0113120122323022-0032113311003331-2123123301211032-1102123033133130-3122223030100030-2031003303132212) |
| `ddos_profile` | [ddos_profile](data-sources--dns_proxy--reference--group-001.md#canonical-1022223003131002-1132102030033213-1323221023322321-2122033121232110-2022100132101232-2123330301323211-1202121032232222-0312332210210001) |
| `ddos_profile.disable_ddos_mitigation` | [ddos_profile.disable_ddos_mitigation](data-sources--dns_proxy--reference--group-001.md#canonical-1321022133100300-2033012213201211-0132031113312130-0102130000000033-1030002213232013-2012311313322111-0300032101231022-1321211133131202) |
| `ddos_profile.enable_ddos_mitigation` | [ddos_profile.enable_ddos_mitigation](data-sources--dns_proxy--reference--group-001.md#canonical-0102200000332022-1222110321303101-0311222301222330-3021331032033122-0212203131033013-3323130000110332-3200120223010203-0032322103010320) |
| `description` | [description](data-sources--dns_proxy--reference--group-001.md#canonical-2030300032212210-3012122021002300-0232211130222320-1322301023120310-1331201311011300-3233220133022001-1013300112322300-1202302212122133) |
| `id` | [ID](data-sources--dns_proxy--reference--group-001.md#canonical-2203303302232133-1321201322111312-2120233310331030-1221231320222231-2120221032123330-1310211301031032-0132120111301302-0001123212302130) |
| `irules` | [irules](data-sources--dns_proxy--reference--group-001.md#canonical-2002111320201302-1212001233333213-0221101100323211-1003033322020300-3033313321222010-0111222002222112-1313103231103203-1331333230331133) |
| `irules.name` | [irules.name](data-sources--dns_proxy--reference--group-001.md#canonical-1012132332122103-2223200313032320-0031200121023030-3030111213200133-1012132120001020-0233100021202133-1311312031313033-3222203030302203) |
| `irules.namespace` | [irules.namespace](data-sources--dns_proxy--reference--group-001.md#canonical-1011311111011113-3310003203120102-3010231111020231-2123032220222302-0030232312003233-1300130111310221-1113033302021010-1021123233122021) |
| `irules.tenant` | [irules.tenant](data-sources--dns_proxy--reference--group-001.md#canonical-1001102032312020-2031012001011101-0233211312102320-3000000011011231-0013123022120132-2110211020320202-2121102122200102-2121031211320222) |
| `labels` | [labels](data-sources--dns_proxy--reference--group-001.md#canonical-1033202112132120-1000201200221111-2002121313311212-2110002011120232-0102201031322022-3123012033101020-0203203201302121-1321033030123110) |
| `lb_algorithm` | [lb_algorithm](data-sources--dns_proxy--reference--group-001.md#canonical-1032100130110113-1223331012221000-2322201311222131-0303102221030023-3022121323130110-3020211210120002-1022102120123021-2220110310322033) |
| `lb_algorithm.round_robin` | [lb_algorithm.round_robin](data-sources--dns_proxy--reference--group-001.md#canonical-2300031331322231-2112331103131303-2301320023332231-1203322121033220-3000323330012000-2111331010302010-2122113303313213-3213312331232303) |
| `name` | [name](data-sources--dns_proxy--reference--group-001.md#canonical-0331030223323232-0131121113130002-3011010202233303-0211222100332101-1010311220010101-0112232303322302-3002312212222211-0300121000213330) |
| `namespace` | [namespace](data-sources--dns_proxy--reference--group-001.md#canonical-1220331021231011-1023322332200031-0320023103322213-0111120132201313-2210133013310020-2113232032220100-1323301313012233-2230021323030230) |
| `origin_servers` | [origin_servers](data-sources--dns_proxy--reference--group-001.md#canonical-2330321113221132-2103301201322102-1112102121233110-3220023132102032-0221311220130111-1110210010230022-0131030013111333-0212031031203123) |
| `origin_servers.health_checks` | [origin_servers.health_checks](data-sources--dns_proxy--reference--group-001.md#canonical-2000022203130112-2201210303311033-3232101013110101-2130001033002122-1023213101122131-1310132333302031-0000200311323010-2322003201302032) |
| `origin_servers.health_checks.health_check` | [origin_servers.health_checks.health_check](data-sources--dns_proxy--reference--group-001.md#canonical-1010223322301010-0133331031200222-1310111020203213-2022021333023212-0013010221030230-1011231113033121-3012130230022033-2221123013301103) |
| `origin_servers.health_checks.health_check.dns_health_check` | [origin_servers.health_checks.health_check.dns_health_check](data-sources--dns_proxy--reference--group-001.md#canonical-1210333032102033-2101213202313333-2000200012021000-1312312320133002-3010332213132220-3013211100013303-1320230331323333-2330111332202011) |
| `origin_servers.health_checks.health_check.dns_health_check.expected_rcode` | [origin_servers.health_checks.health_check.dns_health_check.expected_rcode](data-sources--dns_proxy--reference--group-001.md#canonical-1312300323231221-0322213122300311-1320013030122201-0133302110022031-0030331121023233-0020210021232303-0322303222020021-3212030130303023) |
| `origin_servers.health_checks.health_check.dns_health_check.expected_record_type` | [origin_servers.health_checks.health_check.dns_health_check.expected_record_type](data-sources--dns_proxy--reference--group-001.md#canonical-1330131113132210-3131232113103133-2010030332331223-1220231013321222-3323101232221312-0100320100010233-3113030312010031-3222012033231200) |
| `origin_servers.health_checks.health_check.dns_health_check.expected_response` | [origin_servers.health_checks.health_check.dns_health_check.expected_response](data-sources--dns_proxy--reference--group-001.md#canonical-3321312103320331-0313331201202312-2200331133200000-1131123213110132-0233202310213023-1223312333132102-0312212220031201-2121100103101110) |
| `origin_servers.health_checks.health_check.dns_health_check.query_name` | [origin_servers.health_checks.health_check.dns_health_check.query_name](data-sources--dns_proxy--reference--group-001.md#canonical-0222031101213312-0300103222032303-2100330123330313-0323220121323221-2231222321012232-2310032113302130-1121212312023123-2111011331221302) |
| `origin_servers.health_checks.health_check.dns_health_check.query_type` | [origin_servers.health_checks.health_check.dns_health_check.query_type](data-sources--dns_proxy--reference--group-001.md#canonical-3012231333323031-2303020101202202-2130131103102332-3121000101232131-0230121203333231-0032130121222033-2300333330121232-3321003013011210) |
| `origin_servers.health_checks.health_check.dns_health_check.reverse` | [origin_servers.health_checks.health_check.dns_health_check.reverse](data-sources--dns_proxy--reference--group-001.md#canonical-2321223202022331-2132203003331121-0312013211203231-0000302301230002-2200302233232001-0213020031303220-3212311222013210-3323021013212231) |
| `origin_servers.health_checks.health_check.icmp_health_check` | [origin_servers.health_checks.health_check.icmp_health_check](data-sources--dns_proxy--reference--group-001.md#canonical-1233301013101320-0132231330112002-2303123213201031-0230033033020323-2303111311112032-3313121233033321-1011131300003330-3110031213302111) |
| `origin_servers.health_checks.health_check.tcp_health_check` | [origin_servers.health_checks.health_check.tcp_health_check](data-sources--dns_proxy--reference--group-001.md#canonical-2220222330033021-0101100322232023-2120333033002331-0202023311220021-3320301303330302-0301100302002313-1211003321122211-0212100010310003) |
| `origin_servers.health_checks.health_check.tcp_health_check.expected_response` | [origin_servers.health_checks.health_check.tcp_health_check.expected_response](data-sources--dns_proxy--reference--group-001.md#canonical-3232012320211020-3033130113121123-3112333103212132-3301022120032230-0030213211302000-0230000001302322-3022022033023231-1313131212030131) |
| `origin_servers.health_checks.health_check.tcp_health_check.send_payload` | [origin_servers.health_checks.health_check.tcp_health_check.send_payload](data-sources--dns_proxy--reference--group-001.md#canonical-1232333323131311-1310130003121213-2232121221201132-3013230323100220-1222012000333311-3020133212230230-1313310133121332-0112232112232002) |
| `origin_servers.health_checks.healthy_threshold` | [origin_servers.health_checks.healthy_threshold](data-sources--dns_proxy--reference--group-001.md#canonical-2012132130031101-1200121131120100-1102330200130202-0001220033013303-3221223301310030-3231310021301101-3030020020300020-2123130222301311) |
| `origin_servers.health_checks.interval` | [origin_servers.health_checks.interval](data-sources--dns_proxy--reference--group-001.md#canonical-2012220001232131-0203100020212132-1133303111223202-2220033320330333-3221302022213100-2221103200031131-3300122220211202-2230330212022211) |
| `origin_servers.health_checks.timeout` | [origin_servers.health_checks.timeout](data-sources--dns_proxy--reference--group-001.md#canonical-3313312120002122-0201023203012301-1132130330223331-0301210322323213-0012012332323022-3322212001313113-1233312132123122-3301323222313313) |
| `origin_servers.health_checks.unhealthy_threshold` | [origin_servers.health_checks.unhealthy_threshold](data-sources--dns_proxy--reference--group-001.md#canonical-2232003033232100-1320311101131233-2103003020212121-0113013213031302-1321113320001230-0211220111201210-2330021120122211-0013202222201320) |
| `origin_servers.origin_servers` | [origin_servers.origin_servers](data-sources--dns_proxy--reference--group-001.md#canonical-1100231120111211-0333332200130232-3222310002132022-2201111203131122-0122330332013221-0130311023023003-0202312200100112-3213311103322112) |
| `origin_servers.origin_servers.k8s_service` | [origin_servers.origin_servers.k8s_service](data-sources--dns_proxy--reference--group-001.md#canonical-0133021023131133-1330130202321201-2210230130122320-2113232101031102-1202133220121022-1110312032300213-2033122121322110-1032323022121221) |
| `origin_servers.origin_servers.k8s_service.inside_network` | [origin_servers.origin_servers.k8s_service.inside_network](data-sources--dns_proxy--reference--group-001.md#canonical-2100211203310230-3331112213210123-0003001310023223-1313333102100132-3122322232101233-2101202220100103-3011000210202133-2201331210110023) |
| `origin_servers.origin_servers.k8s_service.outside_network` | [origin_servers.origin_servers.k8s_service.outside_network](data-sources--dns_proxy--reference--group-001.md#canonical-2223212311231312-2301120003020212-2031331311203101-2230200011220020-3320011130120011-0331001200131020-3200301112100002-1330312212020023) |
| `origin_servers.origin_servers.k8s_service.protocol` | [origin_servers.origin_servers.k8s_service.protocol](data-sources--dns_proxy--reference--group-001.md#canonical-1312200010103203-1313001330002320-2200203222121221-1222320310002130-2003203333321012-1030232131220202-2033020122232231-2321021220211200) |
| `origin_servers.origin_servers.k8s_service.service_name` | [origin_servers.origin_servers.k8s_service.service_name](data-sources--dns_proxy--reference--group-001.md#canonical-1021030233122010-2230121232100301-3212022000332232-1232111132213212-0232203111303033-1012133102212021-3013311232202213-3222333023122031) |
| `origin_servers.origin_servers.k8s_service.site_locator` | [origin_servers.origin_servers.k8s_service.site_locator](data-sources--dns_proxy--reference--group-001.md#canonical-2010112330321003-3311033302201023-2111101233233101-0001331223312123-3232302213111001-2212302220301032-3013003221300131-1121201102212012) |
| `origin_servers.origin_servers.k8s_service.site_locator.site` | [origin_servers.origin_servers.k8s_service.site_locator.site](data-sources--dns_proxy--reference--group-001.md#canonical-2310310021013020-0221321023223201-1201201103310302-2311011202203021-1321031223131331-0202000220201130-3001110301023331-0100011022332121) |
| `origin_servers.origin_servers.k8s_service.site_locator.site.name` | [origin_servers.origin_servers.k8s_service.site_locator.site.name](data-sources--dns_proxy--reference--group-001.md#canonical-1120130322132032-2002103311200231-2032110021222101-0032301102222000-3010323032000132-1222112210210211-3110330201000231-3020023032302323) |
| `origin_servers.origin_servers.k8s_service.site_locator.site.namespace` | [origin_servers.origin_servers.k8s_service.site_locator.site.namespace](data-sources--dns_proxy--reference--group-001.md#canonical-1012000211200102-3102223312121122-3230120030202231-3030302200021002-2130113132213123-1230130030220001-2223313023100032-3231002101002020) |
| `origin_servers.origin_servers.k8s_service.site_locator.site.tenant` | [origin_servers.origin_servers.k8s_service.site_locator.site.tenant](data-sources--dns_proxy--reference--group-001.md#canonical-2320230122010330-3011103012012232-2311202233120330-1302103022031332-1103210032013133-3232032130331011-1122201111203012-2123210100311001) |
| `origin_servers.origin_servers.k8s_service.site_locator.virtual_site` | [origin_servers.origin_servers.k8s_service.site_locator.virtual_site](data-sources--dns_proxy--reference--group-001.md#canonical-1112122201202133-3220123300200322-0203011023212130-2332303121201022-1132032312112323-3302101322330023-3022322330222000-2011023023001213) |
| `origin_servers.origin_servers.k8s_service.site_locator.virtual_site.name` | [origin_servers.origin_servers.k8s_service.site_locator.virtual_site.name](data-sources--dns_proxy--reference--group-001.md#canonical-1122101020120033-1303201331013123-0112233200101213-1000122313231002-3001032101023122-3333110201211330-1021122032320030-3310123333030322) |
| `origin_servers.origin_servers.k8s_service.site_locator.virtual_site.namespace` | [origin_servers.origin_servers.k8s_service.site_locator.virtual_site.namespace](data-sources--dns_proxy--reference--group-001.md#canonical-1022101100002032-3120130312102111-0232213331121300-3032033011001121-1120321023100113-2211310021201231-0110111220000100-0101312302311030) |
| `origin_servers.origin_servers.k8s_service.site_locator.virtual_site.tenant` | [origin_servers.origin_servers.k8s_service.site_locator.virtual_site.tenant](data-sources--dns_proxy--reference--group-001.md#canonical-0310213102000020-1200022320300101-3223001311213331-0312021312333212-1201011111012122-2000022201202221-3121013230320211-0202010112131322) |
| `origin_servers.origin_servers.k8s_service.snat_pool` | [origin_servers.origin_servers.k8s_service.snat_pool](data-sources--dns_proxy--reference--group-001.md#canonical-0213031102313032-0210132230313201-1131101102112201-2001120200002223-1212233031310302-2010210013003113-3120121130110012-0013130232003111) |
| `origin_servers.origin_servers.k8s_service.snat_pool.no_snat_pool` | [origin_servers.origin_servers.k8s_service.snat_pool.no_snat_pool](data-sources--dns_proxy--reference--group-001.md#canonical-0012103130012331-1001032031002321-0200230312303000-0003012320221221-1030321132331121-3233103220122113-3110013032321013-0010013112123003) |
| `origin_servers.origin_servers.k8s_service.snat_pool.snat_pool` | [origin_servers.origin_servers.k8s_service.snat_pool.snat_pool](data-sources--dns_proxy--reference--group-001.md#canonical-2231331122122323-0032131131322122-0322022032132030-3010133011101000-1300323302031223-3231031321310202-1111200033312032-2301102313222220) |
| `origin_servers.origin_servers.k8s_service.snat_pool.snat_pool.prefixes` | [origin_servers.origin_servers.k8s_service.snat_pool.snat_pool.prefixes](data-sources--dns_proxy--reference--group-001.md#canonical-3021031213333123-2130132032102121-0013313111333033-1210011330100213-2120131211011321-1000331323021222-2121222311022121-3222202002331100) |
| `origin_servers.origin_servers.k8s_service.vk8s_networks` | [origin_servers.origin_servers.k8s_service.vk8s_networks](data-sources--dns_proxy--reference--group-001.md#canonical-1120011231010002-3131210232000201-0011313230030020-3032202223020231-2310132311213132-1001302322202110-0311201013001101-0320320322213132) |
| `origin_servers.origin_servers.no_preference` | [origin_servers.origin_servers.no_preference](data-sources--dns_proxy--reference--group-001.md#canonical-1030232110030201-0332310310231100-0120123002133000-2211122020112200-3132120022222002-0102333330133323-2223113300330032-2033033123102221) |
| `origin_servers.origin_servers.public_ip` | [origin_servers.origin_servers.public_ip](data-sources--dns_proxy--reference--group-001.md#canonical-0321001032211202-1233331120121310-0223033013230012-3212010132301122-2232303212011113-1303301103033203-3220211121330302-0210021011221220) |
| `origin_servers.origin_servers.public_ip.ip` | [origin_servers.origin_servers.public_ip.ip](data-sources--dns_proxy--reference--group-001.md#canonical-1000113130113222-2130123230010002-0132330210111021-1020113212020211-0000220222012202-0031100322231223-3322302220201031-1330030120011210) |
| `origin_servers.origin_servers.public_name` | [origin_servers.origin_servers.public_name](data-sources--dns_proxy--reference--group-001.md#canonical-0231003220033013-2212103223332030-0202001221333110-0020001310113101-3122003022121313-1011013001221031-3112222033122223-1101210102202013) |
| `origin_servers.origin_servers.public_name.dns_name` | [origin_servers.origin_servers.public_name.dns_name](data-sources--dns_proxy--reference--group-001.md#canonical-1202113103101030-3223203110210203-1223323030130320-1013113210321230-2331032013130202-0100100312213313-1232211303012122-1302320222300231) |
| `origin_servers.origin_servers.public_name.refresh_interval` | [origin_servers.origin_servers.public_name.refresh_interval](data-sources--dns_proxy--reference--group-001.md#canonical-0322012333020212-2201213102303100-0110303000131003-3110133030232133-3130323213132302-0203221102032132-3310003102322203-0021311323323030) |
| `origin_servers.origin_servers.site_preferences` | [origin_servers.origin_servers.site_preferences](data-sources--dns_proxy--reference--group-001.md#canonical-1311300121121300-3301310121221020-3210222301121133-2222020202302323-2133012221132103-3123003202312213-2203030022311023-1000133221022333) |
| `origin_servers.origin_servers.site_preferences.refs` | [origin_servers.origin_servers.site_preferences.refs](data-sources--dns_proxy--reference--group-001.md#canonical-0130033232121321-1210302322100122-2210101203021123-0000021003200311-0102222203110210-2113101021313331-0021320123311022-1320332010013322) |
| `origin_servers.origin_servers.site_preferences.refs.name` | [origin_servers.origin_servers.site_preferences.refs.name](data-sources--dns_proxy--reference--group-001.md#canonical-1112310223310102-2333030012300102-1222000330203232-0133301330110303-2101113102102003-2310113333022120-1323213012123330-3222031001030320) |
| `origin_servers.origin_servers.site_preferences.refs.namespace` | [origin_servers.origin_servers.site_preferences.refs.namespace](data-sources--dns_proxy--reference--group-001.md#canonical-2120130133212133-2321133020011032-2213003130302313-0200101120010013-1100210023301131-0231212311021223-1233013213130021-0121333121010321) |
| `origin_servers.origin_servers.site_preferences.refs.tenant` | [origin_servers.origin_servers.site_preferences.refs.tenant](data-sources--dns_proxy--reference--group-001.md#canonical-2102310230213331-1121231203003132-2333321112110221-3132202130002020-1230313232010302-1012320231113033-0333002122011322-1330031012021233) |
| `protocol_inspection` | [protocol_inspection](data-sources--dns_proxy--reference--group-001.md#canonical-0110322022112223-0222213023211202-1210123301332202-1131323120012022-1102200130013220-1010212300233210-3212330310023220-2231113331112133) |
| `protocol_inspection.name` | [protocol_inspection.name](data-sources--dns_proxy--reference--group-001.md#canonical-3221100310203000-1003122011103113-2221111123013120-3003002101012312-1110003301222321-1131023330200303-3023010020103111-1031331000122303) |
| `protocol_inspection.namespace` | [protocol_inspection.namespace](data-sources--dns_proxy--reference--group-001.md#canonical-2222021032131300-3011220032311303-2300131032013033-3103202313110001-0322310302210120-3130233202012112-1000012213322222-2112220130021013) |
| `protocol_inspection.tenant` | [protocol_inspection.tenant](data-sources--dns_proxy--reference--group-001.md#canonical-0123102233211003-3230300321332003-1332031211013233-0113213033310000-3023012302033213-0213022320231113-2120102110303311-2013011230302322) |
| `proxy_advertisement` | [proxy_advertisement](data-sources--dns_proxy--reference--group-001.md#canonical-1031132200030301-2123232013300210-2020123320233102-0130132220102021-3002201030310313-3323113323131301-1112231331321231-3321032123323210) |
| `proxy_advertisement.advertise_custom` | [proxy_advertisement.advertise_custom](data-sources--dns_proxy--reference--group-001.md#canonical-3330330022111103-2111320100301223-0213202312220003-3030303021230223-2131101010300223-1223112213001030-2110223233220201-1210020323032303) |
| `proxy_advertisement.advertise_custom.advertise_where` | [proxy_advertisement.advertise_custom.advertise_where](data-sources--dns_proxy--reference--group-001.md#canonical-0331120301030313-3011021013112113-0213322212210312-3021133012333232-3011100330302310-1021013321000311-2201103213133012-2110011202202300) |
| `proxy_advertisement.advertise_custom.advertise_where.advertise_dualstack_on_public` | [proxy_advertisement.advertise_custom.advertise_where.advertise_dualstack_on_public](data-sources--dns_proxy--reference--group-002.md#canonical-3333011233220233-0132031303033103-1200130311003002-1220223133333222-3210331022303333-1313100030202130-0022003323210300-2312312312013212) |
| `proxy_advertisement.advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip` | [proxy_advertisement.advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip](data-sources--dns_proxy--reference--group-002.md#canonical-3221332131020000-2223023303302233-3122330110113213-1300333121320221-1113203323233300-2012331110002132-0323300112000101-2122301211032031) |
| `proxy_advertisement.advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip.name` | [proxy_advertisement.advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip.name](data-sources--dns_proxy--reference--group-002.md#canonical-3113202300333111-2031000321111301-3312303333333301-3003230321223202-0213333202203211-1231130123032130-2320233233211132-2020113231123030) |
| `proxy_advertisement.advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip.namespace` | [proxy_advertisement.advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip.namespace](data-sources--dns_proxy--reference--group-002.md#canonical-0121011210201103-1332310323122222-2002020310113220-1210122023132033-2021212303232202-2223212333202002-3003232123032303-3302303223000023) |
| `proxy_advertisement.advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip.tenant` | [proxy_advertisement.advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip.tenant](data-sources--dns_proxy--reference--group-002.md#canonical-2222312310301200-0031310330232132-3132332100202313-2023213103030200-2022233003320012-0330232100133130-2013221112112212-2030331100122132) |
| `proxy_advertisement.advertise_custom.advertise_where.advertise_on_public` | [proxy_advertisement.advertise_custom.advertise_where.advertise_on_public](data-sources--dns_proxy--reference--group-002.md#canonical-0102320013321202-2003132020213231-1031032130321121-3320101220023021-1010301112022212-1002333223131300-0133013302020110-1330220221310202) |
| `proxy_advertisement.advertise_custom.advertise_where.advertise_on_public.public_ip` | [proxy_advertisement.advertise_custom.advertise_where.advertise_on_public.public_ip](data-sources--dns_proxy--reference--group-002.md#canonical-0331033300133222-3020331123101331-1103000323331302-3222021321231122-0123110301202120-1310113330311111-2300030212122203-3231023323323312) |
| `proxy_advertisement.advertise_custom.advertise_where.advertise_on_public.public_ip.name` | [proxy_advertisement.advertise_custom.advertise_where.advertise_on_public.public_ip.name](data-sources--dns_proxy--reference--group-002.md#canonical-0120230132131222-2111030003302231-1330300131330110-1130033330321100-1010131021012030-2322222013312210-0121331011203003-1112220110320001) |
| `proxy_advertisement.advertise_custom.advertise_where.advertise_on_public.public_ip.namespace` | [proxy_advertisement.advertise_custom.advertise_where.advertise_on_public.public_ip.namespace](data-sources--dns_proxy--reference--group-002.md#canonical-2020110113012221-1200033231000212-0020322112333100-2221123221311100-2130330002220312-3020333203210032-2333210102222032-1300121220200332) |
| `proxy_advertisement.advertise_custom.advertise_where.advertise_on_public.public_ip.tenant` | [proxy_advertisement.advertise_custom.advertise_where.advertise_on_public.public_ip.tenant](data-sources--dns_proxy--reference--group-002.md#canonical-2033100113103210-2022101030230022-0002202212131000-2203013131002330-0233101230332301-2122013330202222-2333201030002300-1332312212322303) |
| `proxy_advertisement.advertise_custom.advertise_where.advertise_v6_on_public` | [proxy_advertisement.advertise_custom.advertise_where.advertise_v6_on_public](data-sources--dns_proxy--reference--group-002.md#canonical-2311012002011031-1223120322223032-2113131002211203-3012101321121220-2120213001132301-1213102103133223-2122320020322001-0233333321001020) |
| `proxy_advertisement.advertise_custom.advertise_where.advertise_v6_on_public.public_ip` | [proxy_advertisement.advertise_custom.advertise_where.advertise_v6_on_public.public_ip](data-sources--dns_proxy--reference--group-002.md#canonical-3303133212201001-2323201302023212-3031232322031010-2203013022102122-1203213201030223-0022312212200001-1113301031320020-3023003112001122) |
| `proxy_advertisement.advertise_custom.advertise_where.advertise_v6_on_public.public_ip.name` | [proxy_advertisement.advertise_custom.advertise_where.advertise_v6_on_public.public_ip.name](data-sources--dns_proxy--reference--group-002.md#canonical-2110331301031230-1022303220022312-3233202030311310-0221223322001203-1003112203332300-2333032030212132-0330221231031202-3112122102102200) |
| `proxy_advertisement.advertise_custom.advertise_where.advertise_v6_on_public.public_ip.namespace` | [proxy_advertisement.advertise_custom.advertise_where.advertise_v6_on_public.public_ip.namespace](data-sources--dns_proxy--reference--group-002.md#canonical-1231103032103101-2210013100311311-2130302200033133-0112200032321013-2231213003210303-3231332122002102-3330211112320021-2221010310003322) |
| `proxy_advertisement.advertise_custom.advertise_where.advertise_v6_on_public.public_ip.tenant` | [proxy_advertisement.advertise_custom.advertise_where.advertise_v6_on_public.public_ip.tenant](data-sources--dns_proxy--reference--group-002.md#canonical-0233332010022201-2303332230210232-0322022122211313-3131033213111333-0313212032220201-1031033203311331-2102311331332302-0220333210031321) |
| `proxy_advertisement.advertise_custom.advertise_where.port` | [proxy_advertisement.advertise_custom.advertise_where.port](data-sources--dns_proxy--reference--group-001.md#canonical-1202132000203113-1033303203013103-0001213223220220-1231123023012210-2220010121321103-0200111001330310-3333220111103311-0103101333123330) |
| `proxy_advertisement.advertise_custom.advertise_where.port_ranges` | [proxy_advertisement.advertise_custom.advertise_where.port_ranges](data-sources--dns_proxy--reference--group-001.md#canonical-1121333312202023-2203311203212300-2302333220221022-1131302110200310-2130220312320130-2320000333320213-3201032122023202-2320221221123110) |
| `proxy_advertisement.advertise_custom.advertise_where.site` | [proxy_advertisement.advertise_custom.advertise_where.site](data-sources--dns_proxy--reference--group-002.md#canonical-3002230103303213-1130233102321131-2032122030302012-2331322320012303-1022213021202010-2032312130311310-2203301213133312-2032013302233323) |
| `proxy_advertisement.advertise_custom.advertise_where.site.ip` | [proxy_advertisement.advertise_custom.advertise_where.site.ip](data-sources--dns_proxy--reference--group-002.md#canonical-1132221200331030-2312132310313022-0002223320201022-0222331302200133-1023031332021211-0123232113001021-3003220212100320-1231111013023301) |
| `proxy_advertisement.advertise_custom.advertise_where.site.network` | [proxy_advertisement.advertise_custom.advertise_where.site.network](data-sources--dns_proxy--reference--group-002.md#canonical-0032001201113330-2220220203311302-0033010021002310-2002231302221221-0130320031332300-0320113300000102-1000300310200012-3220212022220313) |
| `proxy_advertisement.advertise_custom.advertise_where.site.site` | [proxy_advertisement.advertise_custom.advertise_where.site.site](data-sources--dns_proxy--reference--group-002.md#canonical-3330302210303232-3122012001122102-0123311121100203-2332212012030202-2312333303223021-3021222200321200-1131233203320131-2120032102303301) |
| `proxy_advertisement.advertise_custom.advertise_where.site.site.name` | [proxy_advertisement.advertise_custom.advertise_where.site.site.name](data-sources--dns_proxy--reference--group-002.md#canonical-1201113331101030-2320323211103110-0231103302300221-0133303112002000-2302021133303031-2223320201000213-1212200200203311-1130122013231321) |
| `proxy_advertisement.advertise_custom.advertise_where.site.site.namespace` | [proxy_advertisement.advertise_custom.advertise_where.site.site.namespace](data-sources--dns_proxy--reference--group-002.md#canonical-0301323222203222-1103223311303200-3312030032012003-3223303113030110-0032022003132210-2120200033010031-1000021120322323-0131121012113103) |
| `proxy_advertisement.advertise_custom.advertise_where.site.site.tenant` | [proxy_advertisement.advertise_custom.advertise_where.site.site.tenant](data-sources--dns_proxy--reference--group-002.md#canonical-0002010213113313-0201133320121003-1023202001331320-2011300332211331-3002331023210213-2203031103221102-2223201120113102-0223132313332003) |
| `proxy_advertisement.advertise_custom.advertise_where.use_default_port` | [proxy_advertisement.advertise_custom.advertise_where.use_default_port](data-sources--dns_proxy--reference--group-002.md#canonical-0030223121030210-1312320332032211-2132220111210333-3120012113201313-1201011211211312-0230323203332331-2100131112231211-2213300021132021) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_network` | [proxy_advertisement.advertise_custom.advertise_where.virtual_network](data-sources--dns_proxy--reference--group-002.md#canonical-2331322221302011-3011100222100032-2112223213132132-0130233310221230-1333323303333101-3320202323320212-3232021230101321-0333331200220302) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_network.default_v6_vip` | [proxy_advertisement.advertise_custom.advertise_where.virtual_network.default_v6_vip](data-sources--dns_proxy--reference--group-002.md#canonical-2120300033223332-0212003101300300-2101201012221121-0113103100130113-1211203203003233-2203010222330231-1332032200130301-1332101203131313) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_network.default_vip` | [proxy_advertisement.advertise_custom.advertise_where.virtual_network.default_vip](data-sources--dns_proxy--reference--group-002.md#canonical-3331302000200313-1002221312100232-0120212033130322-1031103310303210-1320222122131212-2321123302123101-3330102010230003-0032130313133111) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_network.specific_v6_vip` | [proxy_advertisement.advertise_custom.advertise_where.virtual_network.specific_v6_vip](data-sources--dns_proxy--reference--group-002.md#canonical-1203001012231023-2200120113320101-0022301103303333-2130322003311132-1022001233211220-2310333023012330-2221331123001310-3121003323033121) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_network.specific_vip` | [proxy_advertisement.advertise_custom.advertise_where.virtual_network.specific_vip](data-sources--dns_proxy--reference--group-002.md#canonical-0010203030001200-2130022020310002-3022300332011011-0202232131123321-2203321211320321-0012002102131312-2320030323231322-0033103130222112) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_network.virtual_network` | [proxy_advertisement.advertise_custom.advertise_where.virtual_network.virtual_network](data-sources--dns_proxy--reference--group-002.md#canonical-0203010212131123-2330300232303102-1103313022101132-0211010211130121-2133200333110230-3130203022023103-1021321322223332-1200333210210223) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_network.virtual_network.name` | [proxy_advertisement.advertise_custom.advertise_where.virtual_network.virtual_network.name](data-sources--dns_proxy--reference--group-002.md#canonical-2311133313101213-2021133103021123-2323002033030120-3011003123030130-3313111232103231-0321312010000003-2213313223213021-3100300010323302) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_network.virtual_network.namespace` | [proxy_advertisement.advertise_custom.advertise_where.virtual_network.virtual_network.namespace](data-sources--dns_proxy--reference--group-002.md#canonical-3003031010233103-2013001132101032-2222222133113222-0230211101222311-0112213112000232-3330112110102200-0220313012221330-1020011323123023) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_network.virtual_network.tenant` | [proxy_advertisement.advertise_custom.advertise_where.virtual_network.virtual_network.tenant](data-sources--dns_proxy--reference--group-002.md#canonical-0210211020220132-3200312320131231-1032011233300200-2022322110033032-1103100001303112-2210132020222110-1333000221101310-3233022122212110) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_site` | [proxy_advertisement.advertise_custom.advertise_where.virtual_site](data-sources--dns_proxy--reference--group-002.md#canonical-2030102301102310-3322020220231032-3033002123231113-1030022231331202-3231312301032213-2301023233022303-3031131201132213-2220032200033123) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_site.network` | [proxy_advertisement.advertise_custom.advertise_where.virtual_site.network](data-sources--dns_proxy--reference--group-002.md#canonical-0332203300131313-0233111331002231-3030103103333230-3301031331232012-0110031302332022-1122001011312131-0222331321203132-3213132101002323) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_site.virtual_site` | [proxy_advertisement.advertise_custom.advertise_where.virtual_site.virtual_site](data-sources--dns_proxy--reference--group-002.md#canonical-1230302220330002-1201213102013312-2333020322031000-2303231223301033-0013101310221231-1101133203012323-3230020110101223-0100203131023232) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_site.virtual_site.name` | [proxy_advertisement.advertise_custom.advertise_where.virtual_site.virtual_site.name](data-sources--dns_proxy--reference--group-002.md#canonical-2032030110020013-0012111102233320-0011233121202322-2122220311313123-3130121333000231-0312013310231300-0322312213221201-3330311113212023) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_site.virtual_site.namespace` | [proxy_advertisement.advertise_custom.advertise_where.virtual_site.virtual_site.namespace](data-sources--dns_proxy--reference--group-002.md#canonical-0222020302101131-2211101131110232-2011221000010132-2010313321230320-3002133310331212-1231320312202221-2322200203033110-3002310210312122) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_site.virtual_site.tenant` | [proxy_advertisement.advertise_custom.advertise_where.virtual_site.virtual_site.tenant](data-sources--dns_proxy--reference--group-002.md#canonical-2032333332030030-3202030323230123-3321221002101010-0222003313202013-0333031223032002-3123332021313110-1032111230203200-3203303323322201) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip` | [proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip](data-sources--dns_proxy--reference--group-002.md#canonical-1332100020122100-2213112023010000-2322331010010130-1113001101313131-1132321130101011-2103010110333013-2011103123233303-1312022022102233) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip.ip` | [proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip.ip](data-sources--dns_proxy--reference--group-002.md#canonical-0003331001023312-1213212033132123-0232221022122222-0030231331013012-3011001033100033-0203223231300133-3000130212310332-1001002213132201) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip.network` | [proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip.network](data-sources--dns_proxy--reference--group-002.md#canonical-3021000131331310-2113113231220321-2210201320223000-1232321320122223-1233200320332133-0021021221323331-1033302111310302-1003013012333332) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip.virtual_site` | [proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip.virtual_site](data-sources--dns_proxy--reference--group-002.md#canonical-0321113323330103-1222322133010120-0233132221011111-1321122110102210-1112231321310303-2203102111213301-1113310000132302-0330231123301021) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip.virtual_site.name` | [proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip.virtual_site.name](data-sources--dns_proxy--reference--group-002.md#canonical-2212102111221101-3100300312221123-1100020321303032-1200231221013030-1321303222003303-2333111201022010-0333320001000010-0032101301212023) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip.virtual_site.namespace` | [proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip.virtual_site.namespace](data-sources--dns_proxy--reference--group-002.md#canonical-2030310102002022-3102331323011030-2132221222201003-1131330011201313-2132011213023202-2021120300220101-1220022133311221-0101033321231013) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip.virtual_site.tenant` | [proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip.virtual_site.tenant](data-sources--dns_proxy--reference--group-002.md#canonical-1101023220003033-2313312100323330-3210321223312300-2220203003302132-1322130313323033-2310101211311121-0033313022201100-3201003212012233) |
| `proxy_advertisement.advertise_custom.advertise_where.vk8s_service` | [proxy_advertisement.advertise_custom.advertise_where.vk8s_service](data-sources--dns_proxy--reference--group-002.md#canonical-1033112100123332-2013010212202232-1333211020033213-0120032111021023-0233033021033101-1203002012112211-2110220212023321-1230300000122231) |
| `proxy_advertisement.advertise_custom.advertise_where.vk8s_service.site` | [proxy_advertisement.advertise_custom.advertise_where.vk8s_service.site](data-sources--dns_proxy--reference--group-002.md#canonical-3102211002111302-2222211312331122-0031322132221331-3012231000123103-1223231203330221-1102003221333021-3333310232323222-2000302312212130) |
| `proxy_advertisement.advertise_custom.advertise_where.vk8s_service.site.name` | [proxy_advertisement.advertise_custom.advertise_where.vk8s_service.site.name](data-sources--dns_proxy--reference--group-002.md#canonical-3200323101031203-3211102110021310-3323303132003123-3330210011000220-0033302310103312-1303323001113212-3003132232323003-0202010103203123) |
| `proxy_advertisement.advertise_custom.advertise_where.vk8s_service.site.namespace` | [proxy_advertisement.advertise_custom.advertise_where.vk8s_service.site.namespace](data-sources--dns_proxy--reference--group-002.md#canonical-3100230232202330-1213233121303210-2223212101222022-1121321000111203-3331303222203323-0213332021210010-1011010212003012-1011031203100111) |
| `proxy_advertisement.advertise_custom.advertise_where.vk8s_service.site.tenant` | [proxy_advertisement.advertise_custom.advertise_where.vk8s_service.site.tenant](data-sources--dns_proxy--reference--group-002.md#canonical-3211311003012203-3300013122331133-3101231222011223-3011321212301330-3232313200113220-2313223323230203-3133130301001221-1201320200223222) |
| `proxy_advertisement.advertise_custom.advertise_where.vk8s_service.virtual_site` | [proxy_advertisement.advertise_custom.advertise_where.vk8s_service.virtual_site](data-sources--dns_proxy--reference--group-002.md#canonical-1020330201012223-1013003233110010-2120313231101200-1033232133013211-2312320323231303-2213320102131223-2111313033313313-1310100101202113) |
| `proxy_advertisement.advertise_custom.advertise_where.vk8s_service.virtual_site.name` | [proxy_advertisement.advertise_custom.advertise_where.vk8s_service.virtual_site.name](data-sources--dns_proxy--reference--group-002.md#canonical-1020210112011102-0032103032012201-2110310110332023-0130203302120322-2032231222110020-2231132203132311-2303301320302331-2132130011300122) |
| `proxy_advertisement.advertise_custom.advertise_where.vk8s_service.virtual_site.namespace` | [proxy_advertisement.advertise_custom.advertise_where.vk8s_service.virtual_site.namespace](data-sources--dns_proxy--reference--group-002.md#canonical-2022123211100103-2301211113323202-1321001223021030-0201333100112322-1313311001231231-2030202002000312-0213023132231232-2031212113302010) |
| `proxy_advertisement.advertise_custom.advertise_where.vk8s_service.virtual_site.tenant` | [proxy_advertisement.advertise_custom.advertise_where.vk8s_service.virtual_site.tenant](data-sources--dns_proxy--reference--group-002.md#canonical-1333222213132301-0301212022320032-3101201223022103-3122312203300303-3323130120131210-3233001122120021-3221202230311121-3222313030303110) |
| `proxy_advertisement.advertise_dualstack_on_public` | [proxy_advertisement.advertise_dualstack_on_public](data-sources--dns_proxy--reference--group-002.md#canonical-0321120322221213-3230322211221020-3033020113111230-3330130321232023-3212302102201312-3202323102023302-1111220132323122-2323200203010022) |
| `proxy_advertisement.advertise_dualstack_on_public.public_ip` | [proxy_advertisement.advertise_dualstack_on_public.public_ip](data-sources--dns_proxy--reference--group-002.md#canonical-3132031011321222-1222130130003221-1011323212132333-1200311333232201-3033211212101232-2321002132302213-3200121312022003-0132113311131320) |
| `proxy_advertisement.advertise_dualstack_on_public.public_ip.name` | [proxy_advertisement.advertise_dualstack_on_public.public_ip.name](data-sources--dns_proxy--reference--group-002.md#canonical-0330211001301212-3011302010302213-0211330220211020-2232022321302301-0022201023032033-0011301201223232-2102133012102013-2331002302100123) |
| `proxy_advertisement.advertise_dualstack_on_public.public_ip.namespace` | [proxy_advertisement.advertise_dualstack_on_public.public_ip.namespace](data-sources--dns_proxy--reference--group-002.md#canonical-0101313111332023-1323012133113303-2222132230032020-3233000220233120-3200232111321012-0301111221323132-2211112331313203-3203303110132323) |
| `proxy_advertisement.advertise_dualstack_on_public.public_ip.tenant` | [proxy_advertisement.advertise_dualstack_on_public.public_ip.tenant](data-sources--dns_proxy--reference--group-002.md#canonical-2300201110000333-1212300211222212-0112132022012221-0131003301103013-0002300202130132-0323311133210122-0200221131201102-2220111021322302) |
| `proxy_advertisement.advertise_on_public` | [proxy_advertisement.advertise_on_public](data-sources--dns_proxy--reference--group-002.md#canonical-2201111330223031-3100333012303300-1322103101012023-3113003022102223-2113331212030300-2101002211303002-2312100311212232-3220203201203203) |
| `proxy_advertisement.advertise_on_public.public_ip` | [proxy_advertisement.advertise_on_public.public_ip](data-sources--dns_proxy--reference--group-002.md#canonical-0311110300000112-2122110131010333-2302313301012321-1233101230312222-2131300313320323-3223230201323122-0132320200003303-1112203021210312) |
| `proxy_advertisement.advertise_on_public.public_ip.name` | [proxy_advertisement.advertise_on_public.public_ip.name](data-sources--dns_proxy--reference--group-002.md#canonical-0313301030330320-3212022003311130-0111130200010033-1332003012230031-1020333003200220-1300010113332202-0300200322201132-2002103000132133) |
| `proxy_advertisement.advertise_on_public.public_ip.namespace` | [proxy_advertisement.advertise_on_public.public_ip.namespace](data-sources--dns_proxy--reference--group-002.md#canonical-2100010122300000-1011033203333301-1010010023022330-2031030321120320-3120031123031132-1221130123230321-3110112301221131-2301322220130300) |
| `proxy_advertisement.advertise_on_public.public_ip.tenant` | [proxy_advertisement.advertise_on_public.public_ip.tenant](data-sources--dns_proxy--reference--group-002.md#canonical-3010310230103100-1331012023321322-2300303011213312-2113121322022121-3103013300132103-3033310200021301-3230022310013123-1301311321113030) |
| `proxy_advertisement.advertise_on_public_default_dualstack_vip` | [proxy_advertisement.advertise_on_public_default_dualstack_vip](data-sources--dns_proxy--reference--group-002.md#canonical-0002131110333021-0323213313130330-0303222033120113-3121123012310001-0310022123203211-3132333211233222-0021222122221202-0311333213020313) |
| `proxy_advertisement.advertise_on_public_default_ipv6_vip` | [proxy_advertisement.advertise_on_public_default_ipv6_vip](data-sources--dns_proxy--reference--group-002.md#canonical-2100231102122202-1003100230223330-0201110101011120-0310333333020133-0210311033230211-0033222310022021-0001122310313100-0202211021231331) |
| `proxy_advertisement.advertise_on_public_default_vip` | [proxy_advertisement.advertise_on_public_default_vip](data-sources--dns_proxy--reference--group-002.md#canonical-0302120000223013-1311100103032323-1130112221333032-2022213100113122-3233203233211300-2000133011000311-1310023331233323-2121130232011303) |
| `proxy_advertisement.advertise_v6_on_public` | [proxy_advertisement.advertise_v6_on_public](data-sources--dns_proxy--reference--group-002.md#canonical-1101300330101121-0301312132001130-1100120010001300-3010321212103232-1010000210322330-2323121332332221-0032001133121212-2233330130210210) |
| `proxy_advertisement.advertise_v6_on_public.public_ip` | [proxy_advertisement.advertise_v6_on_public.public_ip](data-sources--dns_proxy--reference--group-002.md#canonical-3100220323023013-2123101123301222-3300213013002132-2211313010132301-0100201000131223-0332010232210323-3031111311321013-0110011322011300) |
| `proxy_advertisement.advertise_v6_on_public.public_ip.name` | [proxy_advertisement.advertise_v6_on_public.public_ip.name](data-sources--dns_proxy--reference--group-002.md#canonical-0031200223103212-2302232323030101-2332323010002310-0323212023001321-1321231202110001-1133313121312031-3331022002321302-2001033210120233) |
| `proxy_advertisement.advertise_v6_on_public.public_ip.namespace` | [proxy_advertisement.advertise_v6_on_public.public_ip.namespace](data-sources--dns_proxy--reference--group-002.md#canonical-3323103212220131-1332012031332222-2001313132200200-0300213113223102-3323201103321333-1320112020033002-1020212203333020-2013133310021331) |
| `proxy_advertisement.advertise_v6_on_public.public_ip.tenant` | [proxy_advertisement.advertise_v6_on_public.public_ip.tenant](data-sources--dns_proxy--reference--group-002.md#canonical-2201022200321133-0211213110211113-1230210323023132-2302223133020011-0233112121221232-3220120032022220-0000211013202012-2313232110222222) |
| `proxy_advertisement.do_not_advertise` | [proxy_advertisement.do_not_advertise](data-sources--dns_proxy--reference--group-002.md#canonical-1200102101022133-0203120300303311-2321030000103133-0230002201020300-1210232213020300-2023300023002312-3221333213021123-1111131110311202) |
| `transport_type` | [transport_type](data-sources--dns_proxy--reference--group-001.md#canonical-1321320200212302-1122131020300010-0003132323300200-2230221011020133-0113211213013320-0312212000110013-0000213213122003-1103230312232020) |

<a id="canonical-0322322223123302-2010132100330321-3030313020333012-0333301232303002-2212222210331021-0211123003031332-3000121303000130-3313132200022032"></a>

## Next pages — Property reference / 223330232111 / 12

- [cache_profile](data-sources--dns_proxy--reference--group-001.md#canonical-2123200303011300-0332210102202000-2233020323203102-0102011230031003-3000331233003002-2033021112321102-1333122002121100-0202123300222123)
- [ddos_profile](data-sources--dns_proxy--reference--group-001.md#canonical-0232013023122233-1012213323101023-2001112012330030-3013030103202023-2011131011012030-2021102202320120-0100103221220130-0320000200033100)
- [irules](data-sources--dns_proxy--reference--group-001.md#canonical-1323010111003201-0303002312013010-3103100312030112-0020032232131232-1312132221212131-1312102001102023-3223132230233333-1312111033010102)
- [lb_algorithm](data-sources--dns_proxy--reference--group-001.md#canonical-2013211100010010-0023312102121310-3323002033212331-3103023022222212-0210130112023200-0302223111210332-0203211111130203-1233001021233013)
- [origin_servers](data-sources--dns_proxy--reference--group-001.md#canonical-3102221313302120-2300133213200331-0320001332113013-1100133203331131-3231110233212130-1312212033030000-3302023133321213-2023021333201010)
- [protocol_inspection](data-sources--dns_proxy--reference--group-001.md#canonical-1222310103202331-3022213131203002-0003133222013103-3321032133113232-1300333320113032-0110233333210122-2321331132232223-2213213103001030)
- [proxy_advertisement](data-sources--dns_proxy--reference--group-001.md#canonical-1233321320320303-1032110303233312-2131113231101333-3032323203221001-0110213212120103-3322121223310111-0201121112313233-2310233121220133)
- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-1122013202133033-3201201122310010-3303231012021212-2123120003213300-3130001133310203-2011013122222333-0130313333122120-0323311233013112)

<a id="canonical-2123200303011300-0332210102202000-2233020323203102-0102011230031003-3000331233003002-2033021112321102-1333122002121100-0202123300222123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1210100132010311-0003011011133323-0120022330130203-3100310000112202-1223202310023102-3021232022010102-1323130030231132-1103321132312312"></a>

## cache_profile — cache_profile / 202303000320 / 2

Breadcrumbs:

- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-1122013202133033-3201201122310010-3303231012021212-2123120003213300-3130001133310203-2011013122222333-0130313333122120-0323311233013112)
- [Property reference](data-sources--dns_proxy--reference--group-001.md#canonical-1133330113022100-3110333211021121-3111232100100000-3311232321023320-3202221331130133-0013122201302330-1103000312200021-1203111112100020)
- cache_profile

<a id="canonical-3321112031013310-2302300123101133-1301002112323322-2110120312023332-2330210320100011-2232131120222333-2022230120011331-1031232000230020"></a>

Type: `"single"`. Computed.

DNS Cache specifies cache configuration.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-cache_profile_choice": "[\"cache_size\",\"disable_cache_profile\"]"
}
```

<a id="canonical-3213030201023121-2210313322303230-0200031310330012-3312312311131022-2030021020122300-3010033303223223-0001233302130310-3000200303123313"></a>

## Direct properties — cache_profile / 202303000320 / 3

<a id="canonical-1003330202301011-1111132331000213-1330120332111201-2023130232031023-1312233321001131-3112203203213101-2100313211302120-2333230200121303"></a>

<a id="canonical-0211120132003020-3303111210002220-2313110202120233-1123312331132233-0003210331200130-0102213012002133-2233313212122120-1323333322031330"></a>

## cache_size property — cache_profile / 202303000320 / 4

Type: `"number"`. Computed.

Exclusive with \[disable\_cache\_profile\] cache size.

Upstream description:

Exclusive with \[disable\_cache\_profile\] cache size.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 10240,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "10240"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "10240"
  }
}
```

- [disable_cache_profile](data-sources--dns_proxy--reference--group-001.md#canonical-1233112130130202-3000310023210033-3013013212033332-1221220030120202-2102131031221320-0000000010323112-2222321323302012-3101233302200110): complete subsection reference.

<a id="canonical-1022232330303230-2131133113312203-0100311000313213-1311131213210011-3102020020131102-1130131202010020-2013200210311133-0333300212111310"></a>

## Next pages — cache_profile / 202303000320 / 5

- [cache_profile.disable_cache_profile](data-sources--dns_proxy--reference--group-001.md#canonical-1233112130130202-3000310023210033-3013013212033332-1221220030120202-2102131031221320-0000000010323112-2222321323302012-3101233302200110)
- [Property reference](data-sources--dns_proxy--reference--group-001.md#canonical-1133330113022100-3110333211021121-3111232100100000-3311232321023320-3202221331130133-0013122201302330-1103000312200021-1203111112100020)
- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-1122013202133033-3201201122310010-3303231012021212-2123120003213300-3130001133310203-2011013122222333-0130313333122120-0323311233013112)

<a id="canonical-1233112130130202-3000310023210033-3013013212033332-1221220030120202-2102131031221320-0000000010323112-2222321323302012-3101233302200110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3103001223212311-0020120331131013-3221011130011312-3023120231203220-3000023133330320-2310120121100002-2201002302223031-0021330033211322"></a>

## cache_profile.disable_cache_profile — disable_cache_profile / 322121133122 / 2

Breadcrumbs:

- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-1122013202133033-3201201122310010-3303231012021212-2123120003213300-3130001133310203-2011013122222333-0130313333122120-0323311233013112)
- [Property reference](data-sources--dns_proxy--reference--group-001.md#canonical-1133330113022100-3110333211021121-3111232100100000-3311232321023320-3202221331130133-0013122201302330-1103000312200021-1203111112100020)
- [cache_profile](data-sources--dns_proxy--reference--group-001.md#canonical-2123200303011300-0332210102202000-2233020323203102-0102011230031003-3000331233003002-2033021112321102-1333122002121100-0202123300222123)
- cache_profile.disable_cache_profile

<a id="canonical-2233313110000322-3200300023100220-0113120122323022-0032113311003331-2123123301211032-1102123033133130-3122223030100030-2031003303132212"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for disable cache profile.

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

<a id="canonical-0001303303113132-0122022002223131-1313320223310313-1332201303332223-3023120023102023-1321222013213313-1223232220233111-3313332121301010"></a>

## Direct properties — disable_cache_profile / 322121133122 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0022012311132011-3310300203211303-2130212100233030-1322132210220132-1030011031123000-1230030033331232-1111211023331030-2013011310323003"></a>

## Next pages — disable_cache_profile / 322121133122 / 4

- [cache_profile](data-sources--dns_proxy--reference--group-001.md#canonical-2123200303011300-0332210102202000-2233020323203102-0102011230031003-3000331233003002-2033021112321102-1333122002121100-0202123300222123)
- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-1122013202133033-3201201122310010-3303231012021212-2123120003213300-3130001133310203-2011013122222333-0130313333122120-0323311233013112)

<a id="canonical-0232013023122233-1012213323101023-2001112012330030-3013030103202023-2011131011012030-2021102202320120-0100103221220130-0320000200033100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2133333213011333-1322321231212332-2211231023232112-0223032221010303-1320001002321130-0111010033020103-1211121222032321-3132221123113321"></a>

## ddos_profile — ddos_profile / 330223222220 / 2

Breadcrumbs:

- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-1122013202133033-3201201122310010-3303231012021212-2123120003213300-3130001133310203-2011013122222333-0130313333122120-0323311233013112)
- [Property reference](data-sources--dns_proxy--reference--group-001.md#canonical-1133330113022100-3110333211021121-3111232100100000-3311232321023320-3202221331130133-0013122201302330-1103000312200021-1203111112100020)
- ddos_profile

<a id="canonical-1022223003131002-1132102030033213-1323221023322321-2122033121232110-2022100132101232-2123330301323211-1202121032232222-0312332210210001"></a>

Type: `"single"`. Computed.

Configuration parameter for ddos profile.

Upstream description:

DDoS Protection Rule for DNS.

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

<a id="canonical-2123120022031322-0210010223221203-3133001203012230-2003123103103201-1311103131110212-0122133222220303-2232122213233010-0113312003011020"></a>

## Direct properties — ddos_profile / 330223222220 / 3

- [disable_ddos_mitigation](data-sources--dns_proxy--reference--group-001.md#canonical-3101333200001011-3113102300203110-3322212002113223-2003210311103322-1223122110101110-2301202102322301-1012112211113321-0022331112213033): complete subsection reference.

- [enable_ddos_mitigation](data-sources--dns_proxy--reference--group-001.md#canonical-2202321321110202-1121121331123102-1012113221303013-3211220023032223-3232012220133213-0220311321030322-1130112303232212-3320033022310000): complete subsection reference.

<a id="canonical-0123331102120002-2011201102323233-0331233232113331-0230022322010333-0003223201023113-1003322131011033-1020100211211320-1111320001212232"></a>

## Next pages — ddos_profile / 330223222220 / 4

- [ddos_profile.disable_ddos_mitigation](data-sources--dns_proxy--reference--group-001.md#canonical-3101333200001011-3113102300203110-3322212002113223-2003210311103322-1223122110101110-2301202102322301-1012112211113321-0022331112213033)
- [ddos_profile.enable_ddos_mitigation](data-sources--dns_proxy--reference--group-001.md#canonical-2202321321110202-1121121331123102-1012113221303013-3211220023032223-3232012220133213-0220311321030322-1130112303232212-3320033022310000)
- [Property reference](data-sources--dns_proxy--reference--group-001.md#canonical-1133330113022100-3110333211021121-3111232100100000-3311232321023320-3202221331130133-0013122201302330-1103000312200021-1203111112100020)
- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-1122013202133033-3201201122310010-3303231012021212-2123120003213300-3130001133310203-2011013122222333-0130313333122120-0323311233013112)

<a id="canonical-3101333200001011-3113102300203110-3322212002113223-2003210311103322-1223122110101110-2301202102322301-1012112211113321-0022331112213033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1213201221111001-0111201221022030-0300223020321121-0210112232121303-1323201331230231-1312332120010330-1313233032220013-0311112312013133"></a>

## ddos_profile.disable_ddos_mitigation — disable_ddos_mitigation / 210311311010 / 2

Breadcrumbs:

- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-1122013202133033-3201201122310010-3303231012021212-2123120003213300-3130001133310203-2011013122222333-0130313333122120-0323311233013112)
- [Property reference](data-sources--dns_proxy--reference--group-001.md#canonical-1133330113022100-3110333211021121-3111232100100000-3311232321023320-3202221331130133-0013122201302330-1103000312200021-1203111112100020)
- [ddos_profile](data-sources--dns_proxy--reference--group-001.md#canonical-0232013023122233-1012213323101023-2001112012330030-3013030103202023-2011131011012030-2021102202320120-0100103221220130-0320000200033100)
- ddos_profile.disable_ddos_mitigation

<a id="canonical-1321022133100300-2033012213201211-0132031113312130-0102130000000033-1030002213232013-2012311313322111-0300032101231022-1321211133131202"></a>

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

<a id="canonical-1013000103321222-2022012112210310-1130313003120330-2133101202231201-2213313112220130-2302131103232221-3320001113302202-1121031201332232"></a>

## Direct properties — disable_ddos_mitigation / 210311311010 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0110020312202012-2000132323101020-1013330011312200-0133101231332212-3202111030231202-1012032333221330-2321033101333123-2131331232012132"></a>

## Next pages — disable_ddos_mitigation / 210311311010 / 4

- [ddos_profile](data-sources--dns_proxy--reference--group-001.md#canonical-0232013023122233-1012213323101023-2001112012330030-3013030103202023-2011131011012030-2021102202320120-0100103221220130-0320000200033100)
- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-1122013202133033-3201201122310010-3303231012021212-2123120003213300-3130001133310203-2011013122222333-0130313333122120-0323311233013112)

<a id="canonical-2202321321110202-1121121331123102-1012113221303013-3211220023032223-3232012220133213-0220311321030322-1130112303232212-3320033022310000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3321003013330121-2300231221332121-1332000112202030-3110010103113030-0212023113232330-2023113003333222-3232012013103100-1333311120123133"></a>

## ddos_profile.enable_ddos_mitigation — enable_ddos_mitigation / 101331202221 / 2

Breadcrumbs:

- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-1122013202133033-3201201122310010-3303231012021212-2123120003213300-3130001133310203-2011013122222333-0130313333122120-0323311233013112)
- [Property reference](data-sources--dns_proxy--reference--group-001.md#canonical-1133330113022100-3110333211021121-3111232100100000-3311232321023320-3202221331130133-0013122201302330-1103000312200021-1203111112100020)
- [ddos_profile](data-sources--dns_proxy--reference--group-001.md#canonical-0232013023122233-1012213323101023-2001112012330030-3013030103202023-2011131011012030-2021102202320120-0100103221220130-0320000200033100)
- ddos_profile.enable_ddos_mitigation

<a id="canonical-0102200000332022-1222110321303101-0311222301222330-3021331032033122-0212203131033013-3323130000110332-3200120223010203-0032322103010320"></a>

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

<a id="canonical-1312320322020321-3303330030201221-3012313311102112-1001012321101100-3330103312232213-3002000100132101-2023211023330131-3320133123313100"></a>

## Direct properties — enable_ddos_mitigation / 101331202221 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1320003013222002-1201211320000003-0121023012310302-1211023031001121-1302113320233103-2222330331233132-2002032323133010-2320311021003333"></a>

## Next pages — enable_ddos_mitigation / 101331202221 / 4

- [ddos_profile](data-sources--dns_proxy--reference--group-001.md#canonical-0232013023122233-1012213323101023-2001112012330030-3013030103202023-2011131011012030-2021102202320120-0100103221220130-0320000200033100)
- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-1122013202133033-3201201122310010-3303231012021212-2123120003213300-3130001133310203-2011013122222333-0130313333122120-0323311233013112)

<a id="canonical-1323010111003201-0303002312013010-3103100312030112-0020032232131232-1312132221212131-1312102001102023-3223132230233333-1312111033010102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2031322131031120-0331332033032320-1301303223221301-3300000022203031-2031022113221001-2111313021023333-2331303110320302-3030202000301123"></a>

## irules — irules / 111203210302 / 2

Breadcrumbs:

- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-1122013202133033-3201201122310010-3303231012021212-2123120003213300-3130001133310203-2011013122222333-0130313333122120-0323311233013112)
- [Property reference](data-sources--dns_proxy--reference--group-001.md#canonical-1133330113022100-3110333211021121-3111232100100000-3311232321023320-3202221331130133-0013122201302330-1103000312200021-1203111112100020)
- irules

<a id="canonical-2002111320201302-1212001233333213-0221101100323211-1003033322020300-3033313321222010-0111222002222112-1313103231103203-1331333230331133"></a>

Type: `"list"`. Computed.

OPTIONS for attaching iRules to DNS proxy.

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

<a id="canonical-0022120130301321-0231321100300200-0033323023101323-0212323131110210-1323022321311103-1222233232013131-1112231303031321-2013233313013211"></a>

## Direct properties — irules / 111203210302 / 3

<a id="canonical-1012132332122103-2223200313032320-0031200121023030-3030111213200133-1012132120001020-0233100021202133-1311312031313033-3222203030302203"></a>

<a id="canonical-1023321033311202-3300301021202303-0033010000112111-3002031322212001-3232030010223301-0123230000021221-2300323201211133-2300232331030031"></a>

## name property — irules / 111203210302 / 4

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

<a id="canonical-1011311111011113-3310003203120102-3010231111020231-2123032220222302-0030232312003233-1300130111310221-1113033302021010-1021123233122021"></a>

<a id="canonical-3301032110110132-0311112320020203-0211303320233203-1020111331223330-0112033201013120-1021021312202113-0132003330213222-3202100221031211"></a>

## namespace property — irules / 111203210302 / 5

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

<a id="canonical-1001102032312020-2031012001011101-0233211312102320-3000000011011231-0013123022120132-2110211020320202-2121102122200102-2121031211320222"></a>

<a id="canonical-3122320320210012-2130302221320231-0323203222301012-2302120100231133-1230003213222232-0303311132101212-3011000011013202-3332122321012332"></a>

## tenant property — irules / 111203210302 / 6

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

<a id="canonical-0120020023220133-2103303020322331-1301312121121023-1331003102313003-1011310003010223-1222101203110300-3312112122031212-1132231111111003"></a>

## Next pages — irules / 111203210302 / 7

- [Property reference](data-sources--dns_proxy--reference--group-001.md#canonical-1133330113022100-3110333211021121-3111232100100000-3311232321023320-3202221331130133-0013122201302330-1103000312200021-1203111112100020)
- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-1122013202133033-3201201122310010-3303231012021212-2123120003213300-3130001133310203-2011013122222333-0130313333122120-0323311233013112)

<a id="canonical-2013211100010010-0023312102121310-3323002033212331-3103023022222212-0210130112023200-0302223111210332-0203211111130203-1233001021233013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0202132010011020-0302220130321121-2122003203333223-0330112110313132-2233310210003010-0331013322323122-2231310020011333-1310223212113102"></a>

## lb_algorithm — lb_algorithm / 300011332102 / 2

Breadcrumbs:

- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-1122013202133033-3201201122310010-3303231012021212-2123120003213300-3130001133310203-2011013122222333-0130313333122120-0323311233013112)
- [Property reference](data-sources--dns_proxy--reference--group-001.md#canonical-1133330113022100-3110333211021121-3111232100100000-3311232321023320-3202221331130133-0013122201302330-1103000312200021-1203111112100020)
- lb_algorithm

<a id="canonical-1032100130110113-1223331012221000-2322201311222131-0303102221030023-3022121323130110-3020211210120002-1022102120123021-2220110310322033"></a>

Type: `"single"`. Computed.

Configuration parameter for lb algorithm.

Upstream description:

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

<a id="canonical-0201033230012031-3330322131022021-1110311002220113-0231322231123111-1130112330030203-1313222222322132-1222233202313333-3131131000031202"></a>

## Direct properties — lb_algorithm / 300011332102 / 3

- [round_robin](data-sources--dns_proxy--reference--group-001.md#canonical-1000333222000010-2332222030023020-0103031132200112-1200233201021122-1222101311030121-2112301212131002-3203310020231232-0331221201033200): complete subsection reference.

<a id="canonical-0300230200211131-1222201020320113-0133330310112222-0233131310010120-3023032031222032-0221013121003023-3202303002130321-1212133011221313"></a>

## Next pages — lb_algorithm / 300011332102 / 4

- [lb_algorithm.round_robin](data-sources--dns_proxy--reference--group-001.md#canonical-1000333222000010-2332222030023020-0103031132200112-1200233201021122-1222101311030121-2112301212131002-3203310020231232-0331221201033200)
- [Property reference](data-sources--dns_proxy--reference--group-001.md#canonical-1133330113022100-3110333211021121-3111232100100000-3311232321023320-3202221331130133-0013122201302330-1103000312200021-1203111112100020)
- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-1122013202133033-3201201122310010-3303231012021212-2123120003213300-3130001133310203-2011013122222333-0130313333122120-0323311233013112)

<a id="canonical-1000333222000010-2332222030023020-0103031132200112-1200233201021122-1222101311030121-2112301212131002-3203310020231232-0331221201033200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3203002100313211-0012032011010032-1200031010003003-0233223201100021-1211312022313130-2033020210032222-3200010113310233-1300001001212213"></a>

## lb_algorithm.round_robin — round_robin / 220010131022 / 2

Breadcrumbs:

- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-1122013202133033-3201201122310010-3303231012021212-2123120003213300-3130001133310203-2011013122222333-0130313333122120-0323311233013112)
- [Property reference](data-sources--dns_proxy--reference--group-001.md#canonical-1133330113022100-3110333211021121-3111232100100000-3311232321023320-3202221331130133-0013122201302330-1103000312200021-1203111112100020)
- [lb_algorithm](data-sources--dns_proxy--reference--group-001.md#canonical-2013211100010010-0023312102121310-3323002033212331-3103023022222212-0210130112023200-0302223111210332-0203211111130203-1233001021233013)
- lb_algorithm.round_robin

<a id="canonical-2300031331322231-2112331103131303-2301320023332231-1203322121033220-3000323330012000-2111331010302010-2122113303313213-3213312331232303"></a>

Type: `"single"`. Computed.

Configuration parameter for round robin.

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

<a id="canonical-0333320220233111-0032212203321202-1320131310001221-3110123103231012-2301312102332021-0222223321121121-0113002203001013-0311030213032323"></a>

## Direct properties — round_robin / 220010131022 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1213310213210322-3120330211100002-2013222322220322-1210312111122003-3301111332312212-1012320113200130-0213123200220120-2121003000102022"></a>

## Next pages — round_robin / 220010131022 / 4

- [lb_algorithm](data-sources--dns_proxy--reference--group-001.md#canonical-2013211100010010-0023312102121310-3323002033212331-3103023022222212-0210130112023200-0302223111210332-0203211111130203-1233001021233013)
- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-1122013202133033-3201201122310010-3303231012021212-2123120003213300-3130001133310203-2011013122222333-0130313333122120-0323311233013112)

<a id="canonical-3102221313302120-2300133213200331-0320001332113013-1100133203331131-3231110233212130-1312212033030000-3302023133321213-2023021333201010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2322112312233320-3321301110002230-3120102312131002-2221103330100030-1022333023311322-2333331112132002-3211103210131310-2100313020302022"></a>

## origin_servers — origin_servers / 231320212202 / 2

Breadcrumbs:

- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-1122013202133033-3201201122310010-3303231012021212-2123120003213300-3130001133310203-2011013122222333-0130313333122120-0323311233013112)
- [Property reference](data-sources--dns_proxy--reference--group-001.md#canonical-1133330113022100-3110333211021121-3111232100100000-3311232321023320-3202221331130133-0013122201302330-1103000312200021-1203111112100020)
- origin_servers

<a id="canonical-2330321113221132-2103301201322102-1112102121233110-3220023132102032-0221311220130111-1110210010230022-0131030013111333-0212031031203123"></a>

Type: `"single"`. Computed.

List of origin Servers for the DNS proxy.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3300200311312221-1211302110311310-2033132222213311-2221002323100002-2010333000021210-1323221331101120-1111310030233221-3103102202323232"></a>

## Direct properties — origin_servers / 231320212202 / 3

- [health_checks](data-sources--dns_proxy--reference--group-001.md#canonical-1332222220302031-2101033010320001-3031202103311021-0220323130130120-2010133232320203-1030011000121012-2122110231200210-1132323122311231): complete subsection reference.

- [origin_servers](data-sources--dns_proxy--reference--group-001.md#canonical-1333131311220333-1231120200230111-1310003132330122-1302102303223232-1331131102023222-3003000012212311-2300033331021112-0011131102103213): complete subsection reference.

<a id="canonical-0301333003133013-1330322130101223-3312223010030222-2212320232200111-1030121320000222-0312121021023313-3230031231130331-0220313220103031"></a>

## Next pages — origin_servers / 231320212202 / 4

- [origin_servers.health_checks](data-sources--dns_proxy--reference--group-001.md#canonical-1332222220302031-2101033010320001-3031202103311021-0220323130130120-2010133232320203-1030011000121012-2122110231200210-1132323122311231)
- [origin_servers.origin_servers](data-sources--dns_proxy--reference--group-001.md#canonical-1333131311220333-1231120200230111-1310003132330122-1302102303223232-1331131102023222-3003000012212311-2300033331021112-0011131102103213)
- [Property reference](data-sources--dns_proxy--reference--group-001.md#canonical-1133330113022100-3110333211021121-3111232100100000-3311232321023320-3202221331130133-0013122201302330-1103000312200021-1203111112100020)
- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-1122013202133033-3201201122310010-3303231012021212-2123120003213300-3130001133310203-2011013122222333-0130313333122120-0323311233013112)

<a id="canonical-1332222220302031-2101033010320001-3031202103311021-0220323130130120-2010133232320203-1030011000121012-2122110231200210-1132323122311231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0112312021121101-0311000121021221-0133111102232301-2021322012020331-2132302101310300-2232220020332202-3210330211233220-1322331233213222"></a>

## origin_servers.health_checks — health_checks / 211020331333 / 2

Breadcrumbs:

- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-1122013202133033-3201201122310010-3303231012021212-2123120003213300-3130001133310203-2011013122222333-0130313333122120-0323311233013112)
- [Property reference](data-sources--dns_proxy--reference--group-001.md#canonical-1133330113022100-3110333211021121-3111232100100000-3311232321023320-3202221331130133-0013122201302330-1103000312200021-1203111112100020)
- [origin_servers](data-sources--dns_proxy--reference--group-001.md#canonical-3102221313302120-2300133213200331-0320001332113013-1100133203331131-3231110233212130-1312212033030000-3302023133321213-2023021333201010)
- origin_servers.health_checks

<a id="canonical-2000022203130112-2201210303311033-3232101013110101-2130001033002122-1023213101122131-1310132333302031-0000200311323010-2322003201302032"></a>

Type: `"single"`. Computed.

Configuration parameter for health checks.

Upstream description:

Origin Server Health Checks.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2332222303132211-2312012213123003-1133111132023101-0132022012230120-2332100212230021-1130121002320101-1123121033022100-3330222313031110"></a>

## Direct properties — health_checks / 211020331333 / 3

- [health_check](data-sources--dns_proxy--reference--group-001.md#canonical-0132233331302130-0031011000210212-0333020033133010-3331030123330032-1000313120301310-3130300220132220-2010000011303122-2321022100023132): complete subsection reference.

<a id="canonical-2012132130031101-1200121131120100-1102330200130202-0001220033013303-3221223301310030-3231310021301101-3030020020300020-2123130222301311"></a>

<a id="canonical-0113333330103111-0213333301202023-3203123223331031-3021131202231031-2110220303103301-0310000312102011-2012220311230311-0330231122322022"></a>

## healthy_threshold property — health_checks / 211020331333 / 4

Type: `"number"`. Computed.

Number of successful responses before declaring healthy. In other words, this is the number of
healthy health checks required before a host is marked healthy. Note that during startup, only a
single successful health check is required to mark a host healthy.

Upstream description:

Number of successful responses before declaring healthy. In other words, this is the number of
healthy health checks required before a host is marked healthy. Note that during startup, only a
single successful health check is required to mark a host healthy.

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

<a id="canonical-2012220001232131-0203100020212132-1133303111223202-2220033320330333-3221302022213100-2221103200031131-3300122220211202-2230330212022211"></a>

<a id="canonical-0111033312333302-1023120233210133-2303021300102302-3300323223021111-3310321210010133-0210012330030210-3133200302330131-2230231220233231"></a>

## interval property — health_checks / 211020331333 / 5

Type: `"number"`. Computed.

Time interval in seconds between two healthcheck requests.

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

<a id="canonical-3313312120002122-0201023203012301-1132130330223331-0301210322323213-0012012332323022-3322212001313113-1233312132123122-3301323222313313"></a>

<a id="canonical-0030013300313100-3323330023311013-2311130132301203-3311111131110313-2221111221023223-2123210320213222-1030132222232333-3020001130110210"></a>

## timeout property — health_checks / 211020331333 / 6

Type: `"number"`. Computed.

Timeout in seconds to wait for successful response. In other words, it is the time to wait for a
health check response. If the timeout is reached the health check attempt will be considered a
failure.

Upstream description:

Timeout in seconds to wait for successful response. In other words, it is the time to wait for a
health check response. If the timeout is reached the health check attempt will be considered a
failure.

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

<a id="canonical-2232003033232100-1320311101131233-2103003020212121-0113013213031302-1321113320001230-0211220111201210-2330021120122211-0013202222201320"></a>

<a id="canonical-0133020130121232-3001102133131323-2212203233233110-2231332302213203-1231212133333120-0113132002201321-0121102210020032-2210312003323021"></a>

## unhealthy_threshold property — health_checks / 211020331333 / 7

Type: `"number"`. Computed.

Number of failed responses before declaring unhealthy. In other words, this is the number of
unhealthy health checks required before a host is marked unhealthy. Note that for HTTP health
checking if a host responds with 503 this threshold is ignored and the host is considered unhealthy
immediately.

Upstream description:

Number of failed responses before declaring unhealthy. In other words, this is the number of
unhealthy health checks required before a host is marked unhealthy. Note that for HTTP health
checking if a host responds with 503 this threshold is ignored and the host is considered unhealthy
immediately.

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

<a id="canonical-3022233032203330-0023012330113201-3221103111311310-0110221002303001-0213310111203311-1210202213112021-2030020202322131-0322302101122230"></a>

## Next pages — health_checks / 211020331333 / 8

- [origin_servers.health_checks.health_check](data-sources--dns_proxy--reference--group-001.md#canonical-0132233331302130-0031011000210212-0333020033133010-3331030123330032-1000313120301310-3130300220132220-2010000011303122-2321022100023132)
- [origin_servers](data-sources--dns_proxy--reference--group-001.md#canonical-3102221313302120-2300133213200331-0320001332113013-1100133203331131-3231110233212130-1312212033030000-3302023133321213-2023021333201010)
- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-1122013202133033-3201201122310010-3303231012021212-2123120003213300-3130001133310203-2011013122222333-0130313333122120-0323311233013112)

<a id="canonical-0132233331302130-0031011000210212-0333020033133010-3331030123330032-1000313120301310-3130300220132220-2010000011303122-2321022100023132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3100223032022123-3110113131323212-2131221100320032-3220332200121303-3111112213303020-1221021313211210-3033130222013021-3320131232111322"></a>

## origin_servers.health_checks.health_check — health_check / 323121101103 / 2

Breadcrumbs:

- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-1122013202133033-3201201122310010-3303231012021212-2123120003213300-3130001133310203-2011013122222333-0130313333122120-0323311233013112)
- [Property reference](data-sources--dns_proxy--reference--group-001.md#canonical-1133330113022100-3110333211021121-3111232100100000-3311232321023320-3202221331130133-0013122201302330-1103000312200021-1203111112100020)
- [origin_servers](data-sources--dns_proxy--reference--group-001.md#canonical-3102221313302120-2300133213200331-0320001332113013-1100133203331131-3231110233212130-1312212033030000-3302023133321213-2023021333201010)
- [origin_servers.health_checks](data-sources--dns_proxy--reference--group-001.md#canonical-1332222220302031-2101033010320001-3031202103311021-0220323130130120-2010133232320203-1030011000121012-2122110231200210-1132323122311231)
- origin_servers.health_checks.health_check

<a id="canonical-1010223322301010-0133331031200222-1310111020203213-2022021333023212-0013010221030230-1011231113033121-3012130230022033-2221123013301103"></a>

Type: `"list"`. Computed.

List of Health Checks. List of Health Checks.

Upstream description:

List of Health Checks.

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
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.min_items": "0",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.min_items": "0",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3320103021330122-2123322313100111-2232131031013310-3211213233311110-2013323031113313-1302201013222022-1123231312103110-3010100303312220"></a>

## Direct properties — health_check / 323121101103 / 3

- [dns_health_check](data-sources--dns_proxy--reference--group-001.md#canonical-0312002212212310-1033020011102231-3132201333011310-0231131210110000-1233123020130232-2103231311123121-2200303122030210-1110233021320201): complete subsection reference.

- [icmp_health_check](data-sources--dns_proxy--reference--group-001.md#canonical-1332332322032123-1332320201200211-0331221121100120-1223323202123020-0011113212320303-2321130232201013-1330013323033022-3102003003300222): complete subsection reference.

- [tcp_health_check](data-sources--dns_proxy--reference--group-001.md#canonical-0210222031302311-3203021130323311-2101200201210030-0301320231303232-1113221323222013-1220230322321332-0301321130133311-2222031032113100): complete subsection reference.

<a id="canonical-2122123012310110-2323210023211200-1112211021202113-3022112200123222-2232102103223023-2003233013323030-3312030212201131-1003013133033130"></a>

## Next pages — health_check / 323121101103 / 4

- [origin_servers.health_checks.health_check.dns_health_check](data-sources--dns_proxy--reference--group-001.md#canonical-0312002212212310-1033020011102231-3132201333011310-0231131210110000-1233123020130232-2103231311123121-2200303122030210-1110233021320201)
- [origin_servers.health_checks.health_check.icmp_health_check](data-sources--dns_proxy--reference--group-001.md#canonical-1332332322032123-1332320201200211-0331221121100120-1223323202123020-0011113212320303-2321130232201013-1330013323033022-3102003003300222)
- [origin_servers.health_checks.health_check.tcp_health_check](data-sources--dns_proxy--reference--group-001.md#canonical-0210222031302311-3203021130323311-2101200201210030-0301320231303232-1113221323222013-1220230322321332-0301321130133311-2222031032113100)
- [origin_servers.health_checks](data-sources--dns_proxy--reference--group-001.md#canonical-1332222220302031-2101033010320001-3031202103311021-0220323130130120-2010133232320203-1030011000121012-2122110231200210-1132323122311231)
- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-1122013202133033-3201201122310010-3303231012021212-2123120003213300-3130001133310203-2011013122222333-0130313333122120-0323311233013112)

<a id="canonical-0312002212212310-1033020011102231-3132201333011310-0231131210110000-1233123020130232-2103231311123121-2200303122030210-1110233021320201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0302121220113233-3313030131302110-2321030210001210-0122002322121322-0122032023102000-2133331311212020-3211010032133202-1111030101230112"></a>

## origin_servers.health_checks.health_check.dns_health_check — dns_health_check / 222003312323 / 2

Breadcrumbs:

- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-1122013202133033-3201201122310010-3303231012021212-2123120003213300-3130001133310203-2011013122222333-0130313333122120-0323311233013112)
- [Property reference](data-sources--dns_proxy--reference--group-001.md#canonical-1133330113022100-3110333211021121-3111232100100000-3311232321023320-3202221331130133-0013122201302330-1103000312200021-1203111112100020)
- [origin_servers](data-sources--dns_proxy--reference--group-001.md#canonical-3102221313302120-2300133213200331-0320001332113013-1100133203331131-3231110233212130-1312212033030000-3302023133321213-2023021333201010)
- [origin_servers.health_checks](data-sources--dns_proxy--reference--group-001.md#canonical-1332222220302031-2101033010320001-3031202103311021-0220323130130120-2010133232320203-1030011000121012-2122110231200210-1132323122311231)
- [origin_servers.health_checks.health_check](data-sources--dns_proxy--reference--group-001.md#canonical-0132233331302130-0031011000210212-0333020033133010-3331030123330032-1000313120301310-3130300220132220-2010000011303122-2321022100023132)
- origin_servers.health_checks.health_check.dns_health_check

<a id="canonical-1210333032102033-2101213202313333-2000200012021000-1312312320133002-3010332213132220-3013211100013303-1320230331323333-2330111332202011"></a>

Type: `"single"`. Computed.

DNS health check reports healthy if DNS query is successful and response header and answer matches
the given value.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0033333122300320-3102112100213232-2201033220220033-1123302212032300-3213231202022123-0201032120211032-3112301033221113-0030233003330231"></a>

## Direct properties — dns_health_check / 222003312323 / 3

<a id="canonical-1312300323231221-0322213122300311-1320013030122201-0133302110022031-0030331121023233-0020210021232303-0322303222020021-3212030130303023"></a>

<a id="canonical-3133302301001313-3013032013312301-0331102223023313-1223302001231123-1331102023030133-2023101202203320-0302303111201111-0022230220011000"></a>

## expected_rcode property — dns_health_check / 222003312323 / 4

Type: `"string"`. Computed.

\[Enum: DNS\_RES\_RCODE\_NOERROR|DNS\_RES\_RCODE\_ANY\] Expected DNS Response Rcode Type -
DNS\_RES\_RCODE\_NOERROR: Rcode NOERROR - DNS\_RES\_RCODE\_ANY: RCODE ANY. Possible values are
\`DNS\_RES\_RCODE\_NOERROR\`, \`DNS\_RES\_RCODE\_ANY\`. Defaults to \`DNS\_RES\_RCODE\_NOERROR\`.

Upstream description:

Expected DNS Response Rcode Type

&#8203;- DNS\_RES\_RCODE\_NOERROR: Rcode NOERROR

&#8203;- DNS\_RES\_RCODE\_ANY: RCODE ANY.

Receipt-pinned upstream constraints:

```json
{
  "default": "DNS_RES_RCODE_NOERROR",
  "enum": [
    "DNS_RES_RCODE_NOERROR",
    "DNS_RES_RCODE_ANY"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1330131113132210-3131232113103133-2010030332331223-1220231013321222-3323101232221312-0100320100010233-3113030312010031-3222012033231200"></a>

<a id="canonical-2222000321320213-3320123203320221-3303223131011303-3332331132213330-1311300020210111-3101203220211210-1101132001100312-1313133313213102"></a>

## expected_record_type property — dns_health_check / 222003312323 / 5

Type: `"string"`. Computed.

\[Enum: DNS\_REQUESTED\_QUERY\_TYPE|DNS\_RES\_RECORD\_TYPE\_ANY\] DNS Response Record Type -
DNS\_REQUESTED\_QUERY\_TYPE: Requested Query Type - DNS\_RES\_RECORD\_TYPE\_ANY: Any Record Type.
Possible values are \`DNS\_REQUESTED\_QUERY\_TYPE\`, \`DNS\_RES\_RECORD\_TYPE\_ANY\`. Defaults to
\`DNS\_REQUESTED\_QUERY\_TYPE\`.

Upstream description:

DNS Response Record Type

&#8203;- DNS\_REQUESTED\_QUERY\_TYPE: Requested Query Type

&#8203;- DNS\_RES\_RECORD\_TYPE\_ANY: Any Record Type.

Receipt-pinned upstream constraints:

```json
{
  "default": "DNS_REQUESTED_QUERY_TYPE",
  "enum": [
    "DNS_REQUESTED_QUERY_TYPE",
    "DNS_RES_RECORD_TYPE_ANY"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3321312103320331-0313331201202312-2200331133200000-1131123213110132-0233202310213023-1223312333132102-0312212220031201-2121100103101110"></a>

<a id="canonical-1300303233302333-3001101111211010-0203210010322121-2023100320033110-1231110300321121-2233131012203000-3322303222022003-3301123213101220"></a>

## expected_response property — dns_health_check / 222003312323 / 6

Type: `"string"`. Computed.

Specifies an IPv4 or IPv6 address in the answer section of DNS Response.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 1024,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 1024,
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
    "ves.io.schema.rules.string.max_len": "1024"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "1024"
  }
}
```

<a id="canonical-0222031101213312-0300103222032303-2100330123330313-0323220121323221-2231222321012232-2310032113302130-1121212312023123-2111011331221302"></a>

<a id="canonical-1021120202011020-0122110202223310-0322001102020220-3013212202013310-2002013101113032-1312020011111113-0120100121113111-2021121231213232"></a>

## query_name property — dns_health_check / 222003312323 / 7

Type: `"string"`. Computed.

The query name that the monitor sends a DNS query for.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 255,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 255,
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
    "ves.io.schema.rules.string.max_len": "255",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "255",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-3012231333323031-2303020101202202-2130131103102332-3121000101232131-0230121203333231-0032130121222033-2300333330121232-3321003013011210"></a>

<a id="canonical-1030223032113230-0132121302230032-3100121220301221-3300123220131220-1320202020222030-1031030322013132-1300003132212210-2311332202100023"></a>

## query_type property — dns_health_check / 222003312323 / 8

Type: `"string"`. Computed.

\[Enum: DNS\_QTYPE\_A|DNS\_QTYPE\_AAAA\] DNS Query Type - DNS\_QTYPE\_A: Query Type A -
DNS\_QTYPE\_AAAA: Query Type AAAA. Possible values are \`DNS\_QTYPE\_A\`, \`DNS\_QTYPE\_AAAA\`.
Defaults to \`DNS\_QTYPE\_A\`.

Upstream description:

DNS Query Type

&#8203;- DNS\_QTYPE\_A: Query Type A

&#8203;- DNS\_QTYPE\_AAAA: Query Type AAAA.

Receipt-pinned upstream constraints:

```json
{
  "default": "DNS_QTYPE_A",
  "enum": [
    "DNS_QTYPE_A",
    "DNS_QTYPE_AAAA"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2321223202022331-2132203003331121-0312013211203231-0000302301230002-2200302233232001-0213020031303220-3212311222013210-3323021013212231"></a>

<a id="canonical-3333131332303021-2320103010101113-1122010012002122-2331323133011022-1001311113310233-1322103222001330-0023103223233231-2322232223221203"></a>

## reverse property — dns_health_check / 222003312323 / 9

Type: `"bool"`. Computed.

Enable the monitor operation in reverse mode. When the monitor is in reverse mode, a successful
receive string match marks the monitored object down instead of up.

Upstream description:

Enable the monitor operation in reverse mode. When the monitor is in reverse mode, a successful
receive string match marks the monitored object down instead of up.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0203121021001333-0131030200223132-1031010321320310-0112230130220200-3111123112010133-1020330201013312-3013023320032321-3032133113032102"></a>

## Next pages — dns_health_check / 222003312323 / 10

- [origin_servers.health_checks.health_check](data-sources--dns_proxy--reference--group-001.md#canonical-0132233331302130-0031011000210212-0333020033133010-3331030123330032-1000313120301310-3130300220132220-2010000011303122-2321022100023132)
- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-1122013202133033-3201201122310010-3303231012021212-2123120003213300-3130001133310203-2011013122222333-0130313333122120-0323311233013112)

<a id="canonical-1332332322032123-1332320201200211-0331221121100120-1223323202123020-0011113212320303-2321130232201013-1330013323033022-3102003003300222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1121311221333223-2000211310310300-3102321310330012-1122213323011111-3030002320112010-2211200232312032-0203221023310002-2100322230231203"></a>

## origin_servers.health_checks.health_check.icmp_health_check — icmp_health_check / 212232211230 / 2

Breadcrumbs:

- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-1122013202133033-3201201122310010-3303231012021212-2123120003213300-3130001133310203-2011013122222333-0130313333122120-0323311233013112)
- [Property reference](data-sources--dns_proxy--reference--group-001.md#canonical-1133330113022100-3110333211021121-3111232100100000-3311232321023320-3202221331130133-0013122201302330-1103000312200021-1203111112100020)
- [origin_servers](data-sources--dns_proxy--reference--group-001.md#canonical-3102221313302120-2300133213200331-0320001332113013-1100133203331131-3231110233212130-1312212033030000-3302023133321213-2023021333201010)
- [origin_servers.health_checks](data-sources--dns_proxy--reference--group-001.md#canonical-1332222220302031-2101033010320001-3031202103311021-0220323130130120-2010133232320203-1030011000121012-2122110231200210-1132323122311231)
- [origin_servers.health_checks.health_check](data-sources--dns_proxy--reference--group-001.md#canonical-0132233331302130-0031011000210212-0333020033133010-3331030123330032-1000313120301310-3130300220132220-2010000011303122-2321022100023132)
- origin_servers.health_checks.health_check.icmp_health_check

<a id="canonical-1233301013101320-0132231330112002-2303123213201031-0230033033020323-2303111311112032-3313121233033321-1011131300003330-3110031213302111"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for icmp health check.

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

<a id="canonical-2002030302223220-1132311003102133-0323002101003323-0133122010322223-2232031112213110-1022112223030223-2111312102123121-0220310033003013"></a>

## Direct properties — icmp_health_check / 212232211230 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2110230322200213-1121322132013031-0110212320300003-0302022101011130-1322332030303113-0110121200101122-0212011122201021-1333210110231022"></a>

## Next pages — icmp_health_check / 212232211230 / 4

- [origin_servers.health_checks.health_check](data-sources--dns_proxy--reference--group-001.md#canonical-0132233331302130-0031011000210212-0333020033133010-3331030123330032-1000313120301310-3130300220132220-2010000011303122-2321022100023132)
- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-1122013202133033-3201201122310010-3303231012021212-2123120003213300-3130001133310203-2011013122222333-0130313333122120-0323311233013112)

<a id="canonical-0210222031302311-3203021130323311-2101200201210030-0301320231303232-1113221323222013-1220230322321332-0301321130133311-2222031032113100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2231230132300332-2210011110030031-2322203132033000-1111123133302030-3122120132103323-0120301303031332-0301011013331303-0023230030003301"></a>

## origin_servers.health_checks.health_check.tcp_health_check — tcp_health_check / 201322000101 / 2

Breadcrumbs:

- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-1122013202133033-3201201122310010-3303231012021212-2123120003213300-3130001133310203-2011013122222333-0130313333122120-0323311233013112)
- [Property reference](data-sources--dns_proxy--reference--group-001.md#canonical-1133330113022100-3110333211021121-3111232100100000-3311232321023320-3202221331130133-0013122201302330-1103000312200021-1203111112100020)
- [origin_servers](data-sources--dns_proxy--reference--group-001.md#canonical-3102221313302120-2300133213200331-0320001332113013-1100133203331131-3231110233212130-1312212033030000-3302023133321213-2023021333201010)
- [origin_servers.health_checks](data-sources--dns_proxy--reference--group-001.md#canonical-1332222220302031-2101033010320001-3031202103311021-0220323130130120-2010133232320203-1030011000121012-2122110231200210-1132323122311231)
- [origin_servers.health_checks.health_check](data-sources--dns_proxy--reference--group-001.md#canonical-0132233331302130-0031011000210212-0333020033133010-3331030123330032-1000313120301310-3130300220132220-2010000011303122-2321022100023132)
- origin_servers.health_checks.health_check.tcp_health_check

<a id="canonical-2220222330033021-0101100322232023-2120333033002331-0202023311220021-3320301303330302-0301100302002313-1211003321122211-0212100010310003"></a>

Type: `"single"`. Computed.

Monitor reports healthy status if UDP connection is successful and response payload matches expected
response pattern.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2000322133300300-3102333323303032-0113313012322033-0320133022323023-3011021030102330-3123022333100221-0211202111223010-3220320233132300"></a>

## Direct properties — tcp_health_check / 201322000101 / 3

<a id="canonical-3232012320211020-3033130113121123-3112333103212132-3301022120032230-0030213211302000-0230000001302322-3022022033023231-1313131212030131"></a>

<a id="canonical-2110333220130200-1111011301132320-1033303201232300-0201030201003131-1001013310123130-0330120031113323-2223031313003233-0313313303011002"></a>

## expected_response property — tcp_health_check / 201322000101 / 4

Type: `"string"`. Computed.

Specifies a regular expression pattern which will be matched against response payload.

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

<a id="canonical-1232333323131311-1310130003121213-2232121221201132-3013230323100220-1222012000333311-3020133212230230-1313310133121332-0112232112232002"></a>

<a id="canonical-1220303102223022-2302033102121121-1232032230331321-3003311102212201-2231221321212120-1330330112331010-2303231030012310-1210223122123102"></a>

## send_payload property — tcp_health_check / 201322000101 / 5

Type: `"string"`. Computed.

Send string. Text string sent in the request.

Upstream description:

Text string sent in the request.

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

<a id="canonical-3312032222302322-0112330222031133-0332330033301102-1023130312033223-3012222132003211-1320220020133011-1233332031101010-3330120221012231"></a>

## Next pages — tcp_health_check / 201322000101 / 6

- [origin_servers.health_checks.health_check](data-sources--dns_proxy--reference--group-001.md#canonical-0132233331302130-0031011000210212-0333020033133010-3331030123330032-1000313120301310-3130300220132220-2010000011303122-2321022100023132)
- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-1122013202133033-3201201122310010-3303231012021212-2123120003213300-3130001133310203-2011013122222333-0130313333122120-0323311233013112)

<a id="canonical-1333131311220333-1231120200230111-1310003132330122-1302102303223232-1331131102023222-3003000012212311-2300033331021112-0011131102103213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0013012030203232-3103211210330110-3301000213231012-0302131110012213-3031112020031111-2102033233101132-1002213123022012-3131023112310132"></a>

## origin_servers.origin_servers — origin_servers / 100113122333 / 2

Breadcrumbs:

- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-1122013202133033-3201201122310010-3303231012021212-2123120003213300-3130001133310203-2011013122222333-0130313333122120-0323311233013112)
- [Property reference](data-sources--dns_proxy--reference--group-001.md#canonical-1133330113022100-3110333211021121-3111232100100000-3311232321023320-3202221331130133-0013122201302330-1103000312200021-1203111112100020)
- [origin_servers](data-sources--dns_proxy--reference--group-001.md#canonical-3102221313302120-2300133213200331-0320001332113013-1100133203331131-3231110233212130-1312212033030000-3302023133321213-2023021333201010)
- origin_servers.origin_servers

<a id="canonical-1100231120111211-0333332200130232-3222310002132022-2201111203131122-0122330332013221-0130311023023003-0202312200100112-3213311103322112"></a>

Type: `"list"`. Computed.

List Of Origin Servers. List of origin servers for Proxy.

Upstream description:

List of origin servers for Proxy.

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

<a id="canonical-1122222323122313-0331311023101310-0303103011330323-3021122012111010-3333333332311112-0033101012030113-1203212011113102-3020131100313021"></a>

## Direct properties — origin_servers / 100113122333 / 3

- [k8s_service](data-sources--dns_proxy--reference--group-001.md#canonical-3013321012300130-2001201333112023-2331221100330200-3212311300200201-0022113122003202-1313310102332230-0201213331323330-2230122103332312): complete subsection reference.

- [no_preference](data-sources--dns_proxy--reference--group-001.md#canonical-3222002000231230-0322310333333030-1212100212023011-3133123030311110-3111101221013101-3103003132020103-3203131120331332-2222122021111030): complete subsection reference.

- [public_ip](data-sources--dns_proxy--reference--group-001.md#canonical-1130222201211312-0231111130203012-0031022133111331-0230322101311202-3230020113113213-0202213031223101-2223023003321102-2203102122311123): complete subsection reference.

- [public_name](data-sources--dns_proxy--reference--group-001.md#canonical-2201000022322101-2100232132233100-1010112323033302-0011012222002121-3123112101120233-2131123013132001-3123013223323331-0223223201110012): complete subsection reference.

- [site_preferences](data-sources--dns_proxy--reference--group-001.md#canonical-2030003201020323-1323020011011112-2000103202210133-0303032330110302-0332002331302301-3130130232111302-3222033101030313-1331120330031102): complete subsection reference.

<a id="canonical-1301232010211023-1033321230102201-1222030231011110-1111212002102020-0113231101131012-1311300103332210-2102313110201011-2311301222113133"></a>

## Next pages — origin_servers / 100113122333 / 4

- [origin_servers.origin_servers.k8s_service](data-sources--dns_proxy--reference--group-001.md#canonical-3013321012300130-2001201333112023-2331221100330200-3212311300200201-0022113122003202-1313310102332230-0201213331323330-2230122103332312)
- [origin_servers.origin_servers.no_preference](data-sources--dns_proxy--reference--group-001.md#canonical-3222002000231230-0322310333333030-1212100212023011-3133123030311110-3111101221013101-3103003132020103-3203131120331332-2222122021111030)
- [origin_servers.origin_servers.public_ip](data-sources--dns_proxy--reference--group-001.md#canonical-1130222201211312-0231111130203012-0031022133111331-0230322101311202-3230020113113213-0202213031223101-2223023003321102-2203102122311123)
- [origin_servers.origin_servers.public_name](data-sources--dns_proxy--reference--group-001.md#canonical-2201000022322101-2100232132233100-1010112323033302-0011012222002121-3123112101120233-2131123013132001-3123013223323331-0223223201110012)
- [origin_servers.origin_servers.site_preferences](data-sources--dns_proxy--reference--group-001.md#canonical-2030003201020323-1323020011011112-2000103202210133-0303032330110302-0332002331302301-3130130232111302-3222033101030313-1331120330031102)
- [origin_servers](data-sources--dns_proxy--reference--group-001.md#canonical-3102221313302120-2300133213200331-0320001332113013-1100133203331131-3231110233212130-1312212033030000-3302023133321213-2023021333201010)
- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-1122013202133033-3201201122310010-3303231012021212-2123120003213300-3130001133310203-2011013122222333-0130313333122120-0323311233013112)

<a id="canonical-3013321012300130-2001201333112023-2331221100330200-3212311300200201-0022113122003202-1313310102332230-0201213331323330-2230122103332312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2333133022321231-2312132200223222-2132002020033020-1100032120100023-2311000100133132-1332211221011332-1233023012220200-3123122302002211"></a>

## origin_servers.origin_servers.k8s_service — k8s_service / 030102022303 / 2

Breadcrumbs:

- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-1122013202133033-3201201122310010-3303231012021212-2123120003213300-3130001133310203-2011013122222333-0130313333122120-0323311233013112)
- [Property reference](data-sources--dns_proxy--reference--group-001.md#canonical-1133330113022100-3110333211021121-3111232100100000-3311232321023320-3202221331130133-0013122201302330-1103000312200021-1203111112100020)
- [origin_servers](data-sources--dns_proxy--reference--group-001.md#canonical-3102221313302120-2300133213200331-0320001332113013-1100133203331131-3231110233212130-1312212033030000-3302023133321213-2023021333201010)
- [origin_servers.origin_servers](data-sources--dns_proxy--reference--group-001.md#canonical-1333131311220333-1231120200230111-1310003132330122-1302102303223232-1331131102023222-3003000012212311-2300033331021112-0011131102103213)
- origin_servers.origin_servers.k8s_service

<a id="canonical-0133021023131133-1330130202321201-2210230130122320-2113232101031102-1202133220121022-1110312032300213-2033122121322110-1032323022121221"></a>

Type: `"single"`. Computed.

Specify origin server with K8s service name and site information.

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

<a id="canonical-0212101100131312-2321002321103220-2003231212013300-2101222103000323-3222301220111210-1113311213123312-0003001133323023-2003311032132132"></a>

## Direct properties — k8s_service / 030102022303 / 3

- [inside_network](data-sources--dns_proxy--reference--group-001.md#canonical-0130321123330002-2303201132100133-1010332333022131-2330332331120112-1001320031101020-1130300203311213-1021133101020211-3023032212231312): complete subsection reference.

- [outside_network](data-sources--dns_proxy--reference--group-001.md#canonical-0103022302130302-1022321001203201-1132123231310223-3322322110301130-0122021113200300-3332233101301001-1312123201012210-2201033130013203): complete subsection reference.

<a id="canonical-1312200010103203-1313001330002320-2200203222121221-1222320310002130-2003203333321012-1030232131220202-2033020122232231-2321021220211200"></a>

<a id="canonical-2022002130002212-3002230013310003-0001301121023102-3211321301233200-1031311320131021-2102020210111311-0222221113302311-0021002310221311"></a>

## protocol property — k8s_service / 030102022303 / 4

Type: `"string"`. Computed.

\[Enum: PROTOCOL\_TCP|PROTOCOL\_UDP\] Type of protocol - PROTOCOL\_TCP: TCP - PROTOCOL\_UDP: UDP.
Possible values are \`PROTOCOL\_TCP\`, \`PROTOCOL\_UDP\`. Defaults to \`PROTOCOL\_TCP\`.

Upstream description:

Type of protocol

&#8203;- PROTOCOL\_TCP: TCP

&#8203;- PROTOCOL\_UDP: UDP.

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

<a id="canonical-1021030233122010-2230121232100301-3212022000332232-1232111132213212-0232203111303033-1012133102212021-3013311232202213-3222333023122031"></a>

<a id="canonical-0123213012322031-0022313002301120-1332023331001000-0120033313120231-1331212032032002-3011120030020001-1013230111100122-0311232313301121"></a>

## service_name property — k8s_service / 030102022303 / 5

Type: `"string"`. Computed.

Exclusive with \[\] K8s service name of the origin server will be listed, including the namespace
and cluster-ID. For vK8s services, you need to enter a string with the format
servicename.namespace:example-namespace'frontend', namespace is 'speedtest' and cluster-ID is
'prod', then you will enter..

Upstream description:

Exclusive with \[\] K8s service name of the origin server will be listed, including the namespace
and cluster-ID. For vK8s services, you need to enter a string with the format
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

- [site_locator](data-sources--dns_proxy--reference--group-001.md#canonical-1333000312220030-0121223332311112-2333321231000010-2121133023122032-0321220131313212-3020003021311223-3030303313103321-1000111130123030): complete subsection reference.

- [snat_pool](data-sources--dns_proxy--reference--group-001.md#canonical-3333103011112101-1212323122212031-3321322300123033-1122012010120322-3221322213131001-0023023210020130-3110311321211112-0333030011001132): complete subsection reference.

- [vk8s_networks](data-sources--dns_proxy--reference--group-001.md#canonical-0001132323000120-0331100211112033-2313220311231020-2211230031310220-1232233013130320-1022300101222031-1302121003303302-1132211123313120): complete subsection reference.

<a id="canonical-1332222133112220-1120220323310010-3131011322002121-1123022231333132-2103311130133131-0220120210212133-2103133023002302-2323213332322201"></a>

## Next pages — k8s_service / 030102022303 / 6

- [origin_servers.origin_servers.k8s_service.inside_network](data-sources--dns_proxy--reference--group-001.md#canonical-0130321123330002-2303201132100133-1010332333022131-2330332331120112-1001320031101020-1130300203311213-1021133101020211-3023032212231312)
- [origin_servers.origin_servers.k8s_service.outside_network](data-sources--dns_proxy--reference--group-001.md#canonical-0103022302130302-1022321001203201-1132123231310223-3322322110301130-0122021113200300-3332233101301001-1312123201012210-2201033130013203)
- [origin_servers.origin_servers.k8s_service.site_locator](data-sources--dns_proxy--reference--group-001.md#canonical-1333000312220030-0121223332311112-2333321231000010-2121133023122032-0321220131313212-3020003021311223-3030303313103321-1000111130123030)
- [origin_servers.origin_servers.k8s_service.snat_pool](data-sources--dns_proxy--reference--group-001.md#canonical-3333103011112101-1212323122212031-3321322300123033-1122012010120322-3221322213131001-0023023210020130-3110311321211112-0333030011001132)
- [origin_servers.origin_servers.k8s_service.vk8s_networks](data-sources--dns_proxy--reference--group-001.md#canonical-0001132323000120-0331100211112033-2313220311231020-2211230031310220-1232233013130320-1022300101222031-1302121003303302-1132211123313120)
- [origin_servers.origin_servers](data-sources--dns_proxy--reference--group-001.md#canonical-1333131311220333-1231120200230111-1310003132330122-1302102303223232-1331131102023222-3003000012212311-2300033331021112-0011131102103213)
- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-1122013202133033-3201201122310010-3303231012021212-2123120003213300-3130001133310203-2011013122222333-0130313333122120-0323311233013112)

<a id="canonical-0130321123330002-2303201132100133-1010332333022131-2330332331120112-1001320031101020-1130300203311213-1021133101020211-3023032212231312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0131132302003211-1231223300031300-0300032100111230-2131112210312310-3020233321103201-3101103120021220-3020301033021302-2032001012033222"></a>

## origin_servers.origin_servers.k8s_service.inside_network — inside_network / 231020130302 / 2

Breadcrumbs:

- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-1122013202133033-3201201122310010-3303231012021212-2123120003213300-3130001133310203-2011013122222333-0130313333122120-0323311233013112)
- [Property reference](data-sources--dns_proxy--reference--group-001.md#canonical-1133330113022100-3110333211021121-3111232100100000-3311232321023320-3202221331130133-0013122201302330-1103000312200021-1203111112100020)
- [origin_servers](data-sources--dns_proxy--reference--group-001.md#canonical-3102221313302120-2300133213200331-0320001332113013-1100133203331131-3231110233212130-1312212033030000-3302023133321213-2023021333201010)
- [origin_servers.origin_servers](data-sources--dns_proxy--reference--group-001.md#canonical-1333131311220333-1231120200230111-1310003132330122-1302102303223232-1331131102023222-3003000012212311-2300033331021112-0011131102103213)
- [origin_servers.origin_servers.k8s_service](data-sources--dns_proxy--reference--group-001.md#canonical-3013321012300130-2001201333112023-2331221100330200-3212311300200201-0022113122003202-1313310102332230-0201213331323330-2230122103332312)
- origin_servers.origin_servers.k8s_service.inside_network

<a id="canonical-2100211203310230-3331112213210123-0003001310023223-1313333102100132-3122322232101233-2101202220100103-3011000210202133-2201331210110023"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for inside network.

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

<a id="canonical-2103211231132211-0113021122222213-3213122331222332-3122331212213130-2123232331012110-2230333202102131-0122211022313111-1203031112212200"></a>

## Direct properties — inside_network / 231020130302 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1312310211313131-2130111123301213-3223113301003110-0013313330312311-0233220022332333-3330020030032002-3231203323122133-2101311003233301"></a>

## Next pages — inside_network / 231020130302 / 4

- [origin_servers.origin_servers.k8s_service](data-sources--dns_proxy--reference--group-001.md#canonical-3013321012300130-2001201333112023-2331221100330200-3212311300200201-0022113122003202-1313310102332230-0201213331323330-2230122103332312)
- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-1122013202133033-3201201122310010-3303231012021212-2123120003213300-3130001133310203-2011013122222333-0130313333122120-0323311233013112)

<a id="canonical-0103022302130302-1022321001203201-1132123231310223-3322322110301130-0122021113200300-3332233101301001-1312123201012210-2201033130013203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2222211030203121-3310100120230111-1110023223121331-1013301202103020-0022213302202322-2201213021232312-0203233032030300-0100013023212111"></a>

## origin_servers.origin_servers.k8s_service.outside_network — outside_network / 200222312111 / 2

Breadcrumbs:

- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-1122013202133033-3201201122310010-3303231012021212-2123120003213300-3130001133310203-2011013122222333-0130313333122120-0323311233013112)
- [Property reference](data-sources--dns_proxy--reference--group-001.md#canonical-1133330113022100-3110333211021121-3111232100100000-3311232321023320-3202221331130133-0013122201302330-1103000312200021-1203111112100020)
- [origin_servers](data-sources--dns_proxy--reference--group-001.md#canonical-3102221313302120-2300133213200331-0320001332113013-1100133203331131-3231110233212130-1312212033030000-3302023133321213-2023021333201010)
- [origin_servers.origin_servers](data-sources--dns_proxy--reference--group-001.md#canonical-1333131311220333-1231120200230111-1310003132330122-1302102303223232-1331131102023222-3003000012212311-2300033331021112-0011131102103213)
- [origin_servers.origin_servers.k8s_service](data-sources--dns_proxy--reference--group-001.md#canonical-3013321012300130-2001201333112023-2331221100330200-3212311300200201-0022113122003202-1313310102332230-0201213331323330-2230122103332312)
- origin_servers.origin_servers.k8s_service.outside_network

<a id="canonical-2223212311231312-2301120003020212-2031331311203101-2230200011220020-3320011130120011-0331001200131020-3200301112100002-1330312212020023"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for outside network.

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

<a id="canonical-2202103031131031-1033022023021121-2310133032010103-3121213013220202-3311233103211322-0102233202323133-3220030231330010-3231212102000110"></a>

## Direct properties — outside_network / 200222312111 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2100130011130321-2113202033130232-3201231003221222-1133020303102112-3333211232322311-1222231302220101-1010300032010122-3322010310032121"></a>

## Next pages — outside_network / 200222312111 / 4

- [origin_servers.origin_servers.k8s_service](data-sources--dns_proxy--reference--group-001.md#canonical-3013321012300130-2001201333112023-2331221100330200-3212311300200201-0022113122003202-1313310102332230-0201213331323330-2230122103332312)
- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-1122013202133033-3201201122310010-3303231012021212-2123120003213300-3130001133310203-2011013122222333-0130313333122120-0323311233013112)

<a id="canonical-1333000312220030-0121223332311112-2333321231000010-2121133023122032-0321220131313212-3020003021311223-3030303313103321-1000111130123030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3113130121011333-1032110000203123-0102311220301031-0010322112111022-2112331213122011-2321022331012123-3303222313102110-0323302111003111"></a>

## origin_servers.origin_servers.k8s_service.site_locator — site_locator / 330013131112 / 2

Breadcrumbs:

- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-1122013202133033-3201201122310010-3303231012021212-2123120003213300-3130001133310203-2011013122222333-0130313333122120-0323311233013112)
- [Property reference](data-sources--dns_proxy--reference--group-001.md#canonical-1133330113022100-3110333211021121-3111232100100000-3311232321023320-3202221331130133-0013122201302330-1103000312200021-1203111112100020)
- [origin_servers](data-sources--dns_proxy--reference--group-001.md#canonical-3102221313302120-2300133213200331-0320001332113013-1100133203331131-3231110233212130-1312212033030000-3302023133321213-2023021333201010)
- [origin_servers.origin_servers](data-sources--dns_proxy--reference--group-001.md#canonical-1333131311220333-1231120200230111-1310003132330122-1302102303223232-1331131102023222-3003000012212311-2300033331021112-0011131102103213)
- [origin_servers.origin_servers.k8s_service](data-sources--dns_proxy--reference--group-001.md#canonical-3013321012300130-2001201333112023-2331221100330200-3212311300200201-0022113122003202-1313310102332230-0201213331323330-2230122103332312)
- origin_servers.origin_servers.k8s_service.site_locator

<a id="canonical-2010112330321003-3311033302201023-2111101233233101-0001331223312123-3232302213111001-2212302220301032-3013003221300131-1121201102212012"></a>

Type: `"single"`. Computed.

Message defines a reference to a site or virtual site object.

Upstream description:

This message defines a reference to a site or virtual site object.

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

<a id="canonical-1222100201110031-1202112130203331-3202031301120200-1123230211302123-3302230323100021-3101302120121322-2031302003313001-1020120210001221"></a>

## Direct properties — site_locator / 330013131112 / 3

- [site](data-sources--dns_proxy--reference--group-001.md#canonical-2033212103133131-1313112111300210-0100110222222011-3102032010011010-2212211331221301-1112213002302320-2220323213310331-0323132002212101): complete subsection reference.

- [virtual_site](data-sources--dns_proxy--reference--group-001.md#canonical-3311231132113330-1111221303112132-3313112100220013-2020033000203020-1132130020123223-0322022303131112-2331301011203020-1100012233000012): complete subsection reference.

<a id="canonical-1231312102320212-0212013103322123-0000203122320103-3132012210210111-0120022312220312-1230211333320302-0013022033021220-1300331132130111"></a>

## Next pages — site_locator / 330013131112 / 4

- [origin_servers.origin_servers.k8s_service.site_locator.site](data-sources--dns_proxy--reference--group-001.md#canonical-2033212103133131-1313112111300210-0100110222222011-3102032010011010-2212211331221301-1112213002302320-2220323213310331-0323132002212101)
- [origin_servers.origin_servers.k8s_service.site_locator.virtual_site](data-sources--dns_proxy--reference--group-001.md#canonical-3311231132113330-1111221303112132-3313112100220013-2020033000203020-1132130020123223-0322022303131112-2331301011203020-1100012233000012)
- [origin_servers.origin_servers.k8s_service](data-sources--dns_proxy--reference--group-001.md#canonical-3013321012300130-2001201333112023-2331221100330200-3212311300200201-0022113122003202-1313310102332230-0201213331323330-2230122103332312)
- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-1122013202133033-3201201122310010-3303231012021212-2123120003213300-3130001133310203-2011013122222333-0130313333122120-0323311233013112)

<a id="canonical-2033212103133131-1313112111300210-0100110222222011-3102032010011010-2212211331221301-1112213002302320-2220323213310331-0323132002212101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2313212010132122-3303202010120232-2203330012110123-1311013021031123-3010300100203321-0103020132000231-0211320101312121-2210113033322222"></a>

## origin_servers.origin_servers.k8s_service.site_locator.site — site / 131000332033 / 2

Breadcrumbs:

- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-1122013202133033-3201201122310010-3303231012021212-2123120003213300-3130001133310203-2011013122222333-0130313333122120-0323311233013112)
- [Property reference](data-sources--dns_proxy--reference--group-001.md#canonical-1133330113022100-3110333211021121-3111232100100000-3311232321023320-3202221331130133-0013122201302330-1103000312200021-1203111112100020)
- [origin_servers](data-sources--dns_proxy--reference--group-001.md#canonical-3102221313302120-2300133213200331-0320001332113013-1100133203331131-3231110233212130-1312212033030000-3302023133321213-2023021333201010)
- [origin_servers.origin_servers](data-sources--dns_proxy--reference--group-001.md#canonical-1333131311220333-1231120200230111-1310003132330122-1302102303223232-1331131102023222-3003000012212311-2300033331021112-0011131102103213)
- [origin_servers.origin_servers.k8s_service](data-sources--dns_proxy--reference--group-001.md#canonical-3013321012300130-2001201333112023-2331221100330200-3212311300200201-0022113122003202-1313310102332230-0201213331323330-2230122103332312)
- [origin_servers.origin_servers.k8s_service.site_locator](data-sources--dns_proxy--reference--group-001.md#canonical-1333000312220030-0121223332311112-2333321231000010-2121133023122032-0321220131313212-3020003021311223-3030303313103321-1000111130123030)
- origin_servers.origin_servers.k8s_service.site_locator.site

<a id="canonical-2310310021013020-0221321023223201-1201201103310302-2311011202203021-1321031223131331-0202000220201130-3001110301023331-0100011022332121"></a>

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

<a id="canonical-2120313201323331-0321000121313000-1112133012232130-2013312002020110-0101012302121222-3133230003210000-0031033313112201-0000123020333200"></a>

## Direct properties — site / 131000332033 / 3

<a id="canonical-1120130322132032-2002103311200231-2032110021222101-0032301102222000-3010323032000132-1222112210210211-3110330201000231-3020023032302323"></a>

<a id="canonical-2223230302030222-2130000100222201-0100320311003102-0021212011230033-0001210201120333-2003122122332032-1033132233111201-1101133033302001"></a>

## name property — site / 131000332033 / 4

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

<a id="canonical-1012000211200102-3102223312121122-3230120030202231-3030302200021002-2130113132213123-1230130030220001-2223313023100032-3231002101002020"></a>

<a id="canonical-3302122133102011-0221021330120201-2102010000021111-0202013222113112-2122131321133222-0232220332210121-3032030133032302-3010222103011112"></a>

## namespace property — site / 131000332033 / 5

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

<a id="canonical-2320230122010330-3011103012012232-2311202233120330-1302103022031332-1103210032013133-3232032130331011-1122201111203012-2123210100311001"></a>

<a id="canonical-0120201110211322-1112133032222133-0313002313311021-3132111330111122-0233121321312112-3020233313122031-1210123011212132-3112321223002211"></a>

## tenant property — site / 131000332033 / 6

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

<a id="canonical-2313122113312021-1012131131320212-1010212333132123-2211230320300233-1212322320111010-3331221321330100-3012232213033200-3021133033230310"></a>

## Next pages — site / 131000332033 / 7

- [origin_servers.origin_servers.k8s_service.site_locator](data-sources--dns_proxy--reference--group-001.md#canonical-1333000312220030-0121223332311112-2333321231000010-2121133023122032-0321220131313212-3020003021311223-3030303313103321-1000111130123030)
- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-1122013202133033-3201201122310010-3303231012021212-2123120003213300-3130001133310203-2011013122222333-0130313333122120-0323311233013112)

<a id="canonical-3311231132113330-1111221303112132-3313112100220013-2020033000203020-1132130020123223-0322022303131112-2331301011203020-1100012233000012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2010001233322321-1110101231300323-1222222031031322-1010213101123232-2232100133001023-2001013231112002-2131230132203132-3330311212031120"></a>

## origin_servers.origin_servers.k8s_service.site_locator.virtual_site — virtual_site / 133220121002 / 2

Breadcrumbs:

- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-1122013202133033-3201201122310010-3303231012021212-2123120003213300-3130001133310203-2011013122222333-0130313333122120-0323311233013112)
- [Property reference](data-sources--dns_proxy--reference--group-001.md#canonical-1133330113022100-3110333211021121-3111232100100000-3311232321023320-3202221331130133-0013122201302330-1103000312200021-1203111112100020)
- [origin_servers](data-sources--dns_proxy--reference--group-001.md#canonical-3102221313302120-2300133213200331-0320001332113013-1100133203331131-3231110233212130-1312212033030000-3302023133321213-2023021333201010)
- [origin_servers.origin_servers](data-sources--dns_proxy--reference--group-001.md#canonical-1333131311220333-1231120200230111-1310003132330122-1302102303223232-1331131102023222-3003000012212311-2300033331021112-0011131102103213)
- [origin_servers.origin_servers.k8s_service](data-sources--dns_proxy--reference--group-001.md#canonical-3013321012300130-2001201333112023-2331221100330200-3212311300200201-0022113122003202-1313310102332230-0201213331323330-2230122103332312)
- [origin_servers.origin_servers.k8s_service.site_locator](data-sources--dns_proxy--reference--group-001.md#canonical-1333000312220030-0121223332311112-2333321231000010-2121133023122032-0321220131313212-3020003021311223-3030303313103321-1000111130123030)
- origin_servers.origin_servers.k8s_service.site_locator.virtual_site

<a id="canonical-1112122201202133-3220123300200322-0203011023212130-2332303121201022-1132032312112323-3302101322330023-3022322330222000-2011023023001213"></a>

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

<a id="canonical-1321231310112003-1222032001121311-3122203300032311-3011212122323233-1111132132202222-0012001302332322-3223302102132012-2021303312321101"></a>

## Direct properties — virtual_site / 133220121002 / 3

<a id="canonical-1122101020120033-1303201331013123-0112233200101213-1000122313231002-3001032101023122-3333110201211330-1021122032320030-3310123333030322"></a>

<a id="canonical-2020313230112233-1230100230302112-1311213132213111-3021001300202031-2211032100102300-1300320300303310-2102223033332033-3200111111012232"></a>

## name property — virtual_site / 133220121002 / 4

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

<a id="canonical-1022101100002032-3120130312102111-0232213331121300-3032033011001121-1120321023100113-2211310021201231-0110111220000100-0101312302311030"></a>

<a id="canonical-1330301121210110-2100231131322230-2101100203120013-0000232311220023-1012023112232131-3001023103320332-3103311000312022-2321002031020123"></a>

## namespace property — virtual_site / 133220121002 / 5

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

<a id="canonical-0310213102000020-1200022320300101-3223001311213331-0312021312333212-1201011111012122-2000022201202221-3121013230320211-0202010112131322"></a>

<a id="canonical-0313332110130200-2333330223231302-0230033032011210-2311011321011322-0012310133033022-2030023230313310-1231303300303023-0203201311203012"></a>

## tenant property — virtual_site / 133220121002 / 6

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

<a id="canonical-3321000132331023-3313223310313123-2013303311121302-1220320132002320-0131300322302320-0003221322203132-3302101213220021-2031211023222331"></a>

## Next pages — virtual_site / 133220121002 / 7

- [origin_servers.origin_servers.k8s_service.site_locator](data-sources--dns_proxy--reference--group-001.md#canonical-1333000312220030-0121223332311112-2333321231000010-2121133023122032-0321220131313212-3020003021311223-3030303313103321-1000111130123030)
- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-1122013202133033-3201201122310010-3303231012021212-2123120003213300-3130001133310203-2011013122222333-0130313333122120-0323311233013112)

<a id="canonical-3333103011112101-1212323122212031-3321322300123033-1122012010120322-3221322213131001-0023023210020130-3110311321211112-0333030011001132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1100030301330000-2333010220102100-0322020221232031-1121303203100320-3230132202300023-1301101012120110-0000333231330302-1313321300301011"></a>

## origin_servers.origin_servers.k8s_service.snat_pool — snat_pool / 221102112012 / 2

Breadcrumbs:

- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-1122013202133033-3201201122310010-3303231012021212-2123120003213300-3130001133310203-2011013122222333-0130313333122120-0323311233013112)
- [Property reference](data-sources--dns_proxy--reference--group-001.md#canonical-1133330113022100-3110333211021121-3111232100100000-3311232321023320-3202221331130133-0013122201302330-1103000312200021-1203111112100020)
- [origin_servers](data-sources--dns_proxy--reference--group-001.md#canonical-3102221313302120-2300133213200331-0320001332113013-1100133203331131-3231110233212130-1312212033030000-3302023133321213-2023021333201010)
- [origin_servers.origin_servers](data-sources--dns_proxy--reference--group-001.md#canonical-1333131311220333-1231120200230111-1310003132330122-1302102303223232-1331131102023222-3003000012212311-2300033331021112-0011131102103213)
- [origin_servers.origin_servers.k8s_service](data-sources--dns_proxy--reference--group-001.md#canonical-3013321012300130-2001201333112023-2331221100330200-3212311300200201-0022113122003202-1313310102332230-0201213331323330-2230122103332312)
- origin_servers.origin_servers.k8s_service.snat_pool

<a id="canonical-0213031102313032-0210132230313201-1131101102112201-2001120200002223-1212233031310302-2010210013003113-3120121130110012-0013130232003111"></a>

Type: `"single"`. Computed.

SNAT Pool. SNAT Pool configuration.

Upstream description:

SNAT Pool configuration.

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

<a id="canonical-0000001030203232-0312302322001123-1332320333310020-2120221102220112-2313011301313313-0001110121002323-3230123222013201-0101320300233310"></a>

## Direct properties — snat_pool / 221102112012 / 3

- [no_snat_pool](data-sources--dns_proxy--reference--group-001.md#canonical-2103310230333320-1233103000210031-2102031033320222-3120231020022020-0033012322332112-2110233123333121-1120000103112332-3321122332201120): complete subsection reference.

- [snat_pool](data-sources--dns_proxy--reference--group-001.md#canonical-2113213331213213-2011010023133102-1012200233231212-3010202012002201-1202030003210112-1223332003302322-3320032312021201-2201120302321231): complete subsection reference.

<a id="canonical-1111123012032110-3230002331303311-2322103001220030-0303010230302232-3330320320320121-1331320123020212-2002020210100333-1201333300111212"></a>

## Next pages — snat_pool / 221102112012 / 4

- [origin_servers.origin_servers.k8s_service.snat_pool.no_snat_pool](data-sources--dns_proxy--reference--group-001.md#canonical-2103310230333320-1233103000210031-2102031033320222-3120231020022020-0033012322332112-2110233123333121-1120000103112332-3321122332201120)
- [origin_servers.origin_servers.k8s_service.snat_pool.snat_pool](data-sources--dns_proxy--reference--group-001.md#canonical-2113213331213213-2011010023133102-1012200233231212-3010202012002201-1202030003210112-1223332003302322-3320032312021201-2201120302321231)
- [origin_servers.origin_servers.k8s_service](data-sources--dns_proxy--reference--group-001.md#canonical-3013321012300130-2001201333112023-2331221100330200-3212311300200201-0022113122003202-1313310102332230-0201213331323330-2230122103332312)
- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-1122013202133033-3201201122310010-3303231012021212-2123120003213300-3130001133310203-2011013122222333-0130313333122120-0323311233013112)

<a id="canonical-2103310230333320-1233103000210031-2102031033320222-3120231020022020-0033012322332112-2110233123333121-1120000103112332-3321122332201120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3310221213001231-1223212221101011-1223201230102301-1203113123022012-3010013110231332-2103221311132333-1113333321032012-1230313320111031"></a>

## origin_servers.origin_servers.k8s_service.snat_pool.no_snat_pool — no_snat_pool / 211212231023 / 2

Breadcrumbs:

- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-1122013202133033-3201201122310010-3303231012021212-2123120003213300-3130001133310203-2011013122222333-0130313333122120-0323311233013112)
- [Property reference](data-sources--dns_proxy--reference--group-001.md#canonical-1133330113022100-3110333211021121-3111232100100000-3311232321023320-3202221331130133-0013122201302330-1103000312200021-1203111112100020)
- [origin_servers](data-sources--dns_proxy--reference--group-001.md#canonical-3102221313302120-2300133213200331-0320001332113013-1100133203331131-3231110233212130-1312212033030000-3302023133321213-2023021333201010)
- [origin_servers.origin_servers](data-sources--dns_proxy--reference--group-001.md#canonical-1333131311220333-1231120200230111-1310003132330122-1302102303223232-1331131102023222-3003000012212311-2300033331021112-0011131102103213)
- [origin_servers.origin_servers.k8s_service](data-sources--dns_proxy--reference--group-001.md#canonical-3013321012300130-2001201333112023-2331221100330200-3212311300200201-0022113122003202-1313310102332230-0201213331323330-2230122103332312)
- [origin_servers.origin_servers.k8s_service.snat_pool](data-sources--dns_proxy--reference--group-001.md#canonical-3333103011112101-1212323122212031-3321322300123033-1122012010120322-3221322213131001-0023023210020130-3110311321211112-0333030011001132)
- origin_servers.origin_servers.k8s_service.snat_pool.no_snat_pool

<a id="canonical-0012103130012331-1001032031002321-0200230312303000-0003012320221221-1030321132331121-3233103220122113-3110013032321013-0010013112123003"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for no snat pool.

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

<a id="canonical-3122323103322133-1013332212323332-2322233230310112-1230123002212300-1113021213231101-0301222123310232-3101233121211202-0322111132111013"></a>

## Direct properties — no_snat_pool / 211212231023 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0213333301021312-1001200020111303-3203123100012030-3120303212003130-1113222211221313-2230232120311131-2020202321311230-2003201213222002"></a>

## Next pages — no_snat_pool / 211212231023 / 4

- [origin_servers.origin_servers.k8s_service.snat_pool](data-sources--dns_proxy--reference--group-001.md#canonical-3333103011112101-1212323122212031-3321322300123033-1122012010120322-3221322213131001-0023023210020130-3110311321211112-0333030011001132)
- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-1122013202133033-3201201122310010-3303231012021212-2123120003213300-3130001133310203-2011013122222333-0130313333122120-0323311233013112)

<a id="canonical-2113213331213213-2011010023133102-1012200233231212-3010202012002201-1202030003210112-1223332003302322-3320032312021201-2201120302321231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2111232233003323-1330302310311212-0003331111203220-2012210330321112-2011331233323130-3232212111301112-2102323332311333-3113011130033103"></a>

## origin_servers.origin_servers.k8s_service.snat_pool.snat_pool — snat_pool / 113203222333 / 2

Breadcrumbs:

- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-1122013202133033-3201201122310010-3303231012021212-2123120003213300-3130001133310203-2011013122222333-0130313333122120-0323311233013112)
- [Property reference](data-sources--dns_proxy--reference--group-001.md#canonical-1133330113022100-3110333211021121-3111232100100000-3311232321023320-3202221331130133-0013122201302330-1103000312200021-1203111112100020)
- [origin_servers](data-sources--dns_proxy--reference--group-001.md#canonical-3102221313302120-2300133213200331-0320001332113013-1100133203331131-3231110233212130-1312212033030000-3302023133321213-2023021333201010)
- [origin_servers.origin_servers](data-sources--dns_proxy--reference--group-001.md#canonical-1333131311220333-1231120200230111-1310003132330122-1302102303223232-1331131102023222-3003000012212311-2300033331021112-0011131102103213)
- [origin_servers.origin_servers.k8s_service](data-sources--dns_proxy--reference--group-001.md#canonical-3013321012300130-2001201333112023-2331221100330200-3212311300200201-0022113122003202-1313310102332230-0201213331323330-2230122103332312)
- [origin_servers.origin_servers.k8s_service.snat_pool](data-sources--dns_proxy--reference--group-001.md#canonical-3333103011112101-1212323122212031-3321322300123033-1122012010120322-3221322213131001-0023023210020130-3110311321211112-0333030011001132)
- origin_servers.origin_servers.k8s_service.snat_pool.snat_pool

<a id="canonical-2231331122122323-0032131131322122-0322022032132030-3010133011101000-1300323302031223-3231031321310202-1111200033312032-2301102313222220"></a>

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

<a id="canonical-2222022221202133-3100133123302110-0021032121000210-1202100331323303-0220000233103233-2313121103231320-0201212211302300-3030311112121010"></a>

## Direct properties — snat_pool / 113203222333 / 3

<a id="canonical-3021031213333123-2130132032102121-0013313111333033-1210011330100213-2120131211011321-1000331323021222-2121222311022121-3222202002331100"></a>

<a id="canonical-3203101221012121-2002230033013213-2210011310002013-3212312002002303-2210231220002122-1011212333201003-0112223221033300-2011311233002032"></a>

## prefixes property — snat_pool / 113203222333 / 4

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

<a id="canonical-1332310232131001-3222023122032103-3111131133011301-0312332120301011-3222003110332022-1230110303213203-1222111331213213-2102110332022000"></a>

## Next pages — snat_pool / 113203222333 / 5

- [origin_servers.origin_servers.k8s_service.snat_pool](data-sources--dns_proxy--reference--group-001.md#canonical-3333103011112101-1212323122212031-3321322300123033-1122012010120322-3221322213131001-0023023210020130-3110311321211112-0333030011001132)
- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-1122013202133033-3201201122310010-3303231012021212-2123120003213300-3130001133310203-2011013122222333-0130313333122120-0323311233013112)

<a id="canonical-0001132323000120-0331100211112033-2313220311231020-2211230031310220-1232233013130320-1022300101222031-1302121003303302-1132211123313120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3012001031220002-0013110130112111-3003220111002311-2001130202022303-1221013023303030-0333003303200023-1003302001313230-1331131310323232"></a>

## origin_servers.origin_servers.k8s_service.vk8s_networks — vk8s_networks / 220212011011 / 2

Breadcrumbs:

- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-1122013202133033-3201201122310010-3303231012021212-2123120003213300-3130001133310203-2011013122222333-0130313333122120-0323311233013112)
- [Property reference](data-sources--dns_proxy--reference--group-001.md#canonical-1133330113022100-3110333211021121-3111232100100000-3311232321023320-3202221331130133-0013122201302330-1103000312200021-1203111112100020)
- [origin_servers](data-sources--dns_proxy--reference--group-001.md#canonical-3102221313302120-2300133213200331-0320001332113013-1100133203331131-3231110233212130-1312212033030000-3302023133321213-2023021333201010)
- [origin_servers.origin_servers](data-sources--dns_proxy--reference--group-001.md#canonical-1333131311220333-1231120200230111-1310003132330122-1302102303223232-1331131102023222-3003000012212311-2300033331021112-0011131102103213)
- [origin_servers.origin_servers.k8s_service](data-sources--dns_proxy--reference--group-001.md#canonical-3013321012300130-2001201333112023-2331221100330200-3212311300200201-0022113122003202-1313310102332230-0201213331323330-2230122103332312)
- origin_servers.origin_servers.k8s_service.vk8s_networks

<a id="canonical-1120011231010002-3131210232000201-0011313230030020-3032202223020231-2310132311213132-1001302322202110-0311201013001101-0320320322213132"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for vk8s networks.

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

<a id="canonical-1231023220113330-1301223311033233-3330310021212000-3112231310103320-3012132330021231-3333232131001210-0032030222321113-3011113222332212"></a>

## Direct properties — vk8s_networks / 220212011011 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1101032322331131-2300012123030103-2333300033111231-1131330122030222-1223111100302132-2021122112231310-0231023011030222-1321000303121332"></a>

## Next pages — vk8s_networks / 220212011011 / 4

- [origin_servers.origin_servers.k8s_service](data-sources--dns_proxy--reference--group-001.md#canonical-3013321012300130-2001201333112023-2331221100330200-3212311300200201-0022113122003202-1313310102332230-0201213331323330-2230122103332312)
- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-1122013202133033-3201201122310010-3303231012021212-2123120003213300-3130001133310203-2011013122222333-0130313333122120-0323311233013112)

<a id="canonical-3222002000231230-0322310333333030-1212100212023011-3133123030311110-3111101221013101-3103003132020103-3203131120331332-2222122021111030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1130010323310102-2031233331330110-1213312113130203-1001332210000031-3131000021121222-0210232222133220-2111303033030131-1332013211102023"></a>

## origin_servers.origin_servers.no_preference — no_preference / 020333312120 / 2

Breadcrumbs:

- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-1122013202133033-3201201122310010-3303231012021212-2123120003213300-3130001133310203-2011013122222333-0130313333122120-0323311233013112)
- [Property reference](data-sources--dns_proxy--reference--group-001.md#canonical-1133330113022100-3110333211021121-3111232100100000-3311232321023320-3202221331130133-0013122201302330-1103000312200021-1203111112100020)
- [origin_servers](data-sources--dns_proxy--reference--group-001.md#canonical-3102221313302120-2300133213200331-0320001332113013-1100133203331131-3231110233212130-1312212033030000-3302023133321213-2023021333201010)
- [origin_servers.origin_servers](data-sources--dns_proxy--reference--group-001.md#canonical-1333131311220333-1231120200230111-1310003132330122-1302102303223232-1331131102023222-3003000012212311-2300033331021112-0011131102103213)
- origin_servers.origin_servers.no_preference

<a id="canonical-1030232110030201-0332310310231100-0120123002133000-2211122020112200-3132120022222002-0102333330133323-2223113300330032-2033033123102221"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for no preference.

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

<a id="canonical-1230120102313122-3100320101221131-0012213331203001-1021113222003121-2133033303112031-1331012313120120-1200112232233113-0322310130003220"></a>

## Direct properties — no_preference / 020333312120 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1302130102313100-1123231102113031-0333010101130203-1030203030210002-2101031331031031-0113313210010031-3300030011302333-0210201302330220"></a>

## Next pages — no_preference / 020333312120 / 4

- [origin_servers.origin_servers](data-sources--dns_proxy--reference--group-001.md#canonical-1333131311220333-1231120200230111-1310003132330122-1302102303223232-1331131102023222-3003000012212311-2300033331021112-0011131102103213)
- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-1122013202133033-3201201122310010-3303231012021212-2123120003213300-3130001133310203-2011013122222333-0130313333122120-0323311233013112)

<a id="canonical-1130222201211312-0231111130203012-0031022133111331-0230322101311202-3230020113113213-0202213031223101-2223023003321102-2203102122311123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0132200222331130-3130202211211220-2222131013003302-3132230002332110-3110111010030310-2311330111321323-3111123230230222-2230030333003033"></a>

## origin_servers.origin_servers.public_ip — public_ip / 111301322023 / 2

Breadcrumbs:

- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-1122013202133033-3201201122310010-3303231012021212-2123120003213300-3130001133310203-2011013122222333-0130313333122120-0323311233013112)
- [Property reference](data-sources--dns_proxy--reference--group-001.md#canonical-1133330113022100-3110333211021121-3111232100100000-3311232321023320-3202221331130133-0013122201302330-1103000312200021-1203111112100020)
- [origin_servers](data-sources--dns_proxy--reference--group-001.md#canonical-3102221313302120-2300133213200331-0320001332113013-1100133203331131-3231110233212130-1312212033030000-3302023133321213-2023021333201010)
- [origin_servers.origin_servers](data-sources--dns_proxy--reference--group-001.md#canonical-1333131311220333-1231120200230111-1310003132330122-1302102303223232-1331131102023222-3003000012212311-2300033331021112-0011131102103213)
- origin_servers.origin_servers.public_ip

<a id="canonical-0321001032211202-1233331120121310-0223033013230012-3212010132301122-2232303212011113-1303301103033203-3220211121330302-0210021011221220"></a>

Type: `"single"`. Computed.

Specify origin server with public IP address.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-public_ip_choice": "[\"ip\"]"
}
```

<a id="canonical-0200203213032233-1101311123131021-2030130300201101-3201301220122001-2033201110121021-1110112012102113-2012230123031311-1112123313123123"></a>

## Direct properties — public_ip / 111301322023 / 3

<a id="canonical-1000113130113222-2130123230010002-0132330210111021-1020113212020211-0000220222012202-0031100322231223-3322302220201031-1330030120011210"></a>

<a id="canonical-3100212230323131-3022311013310102-3303223101122212-2023233003012323-1030101200303313-3221311010113331-2230313313203310-3111331110203313"></a>

## ip property — public_ip / 111301322023 / 4

Type: `"string"`. Computed.

Public IPv4. Exclusive with \[\] Public IPv4 address.

Upstream description:

Exclusive with \[\] Public IPv4 address.

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

<a id="canonical-2302311011332232-3010030303032120-3211213112130121-1223302001120002-0103220001233233-1332001313121310-2310231031321331-0030220123213310"></a>

## Next pages — public_ip / 111301322023 / 5

- [origin_servers.origin_servers](data-sources--dns_proxy--reference--group-001.md#canonical-1333131311220333-1231120200230111-1310003132330122-1302102303223232-1331131102023222-3003000012212311-2300033331021112-0011131102103213)
- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-1122013202133033-3201201122310010-3303231012021212-2123120003213300-3130001133310203-2011013122222333-0130313333122120-0323311233013112)

<a id="canonical-2201000022322101-2100232132233100-1010112323033302-0011012222002121-3123112101120233-2131123013132001-3123013223323331-0223223201110012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3122200131222230-3020222233222230-1210231020131212-2131031122200331-2323302000232023-3311013221322122-3320321002201333-1301313113311211"></a>

## origin_servers.origin_servers.public_name — public_name / 011100231130 / 2

Breadcrumbs:

- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-1122013202133033-3201201122310010-3303231012021212-2123120003213300-3130001133310203-2011013122222333-0130313333122120-0323311233013112)
- [Property reference](data-sources--dns_proxy--reference--group-001.md#canonical-1133330113022100-3110333211021121-3111232100100000-3311232321023320-3202221331130133-0013122201302330-1103000312200021-1203111112100020)
- [origin_servers](data-sources--dns_proxy--reference--group-001.md#canonical-3102221313302120-2300133213200331-0320001332113013-1100133203331131-3231110233212130-1312212033030000-3302023133321213-2023021333201010)
- [origin_servers.origin_servers](data-sources--dns_proxy--reference--group-001.md#canonical-1333131311220333-1231120200230111-1310003132330122-1302102303223232-1331131102023222-3003000012212311-2300033331021112-0011131102103213)
- origin_servers.origin_servers.public_name

<a id="canonical-0231003220033013-2212103223332030-0202001221333110-0020001310113101-3122003022121313-1011013001221031-3112222033122223-1101210102202013"></a>

Type: `"single"`. Computed.

Specify origin server with public DNS name.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1312111101233002-3330122213221033-1320203332200032-3033013103020032-1020121200232000-3201032230032333-0310323330221310-2213310022111320"></a>

## Direct properties — public_name / 011100231130 / 3

<a id="canonical-1202113103101030-3223203110210203-1223323030130320-1013113210321230-2331032013130202-0100100312213313-1232211303012122-1302320222300231"></a>

<a id="canonical-0203301310002132-3330120120030032-1101102111201120-3320311101213310-0112202330022320-1132223120011133-0112022222300021-2322201331232120"></a>

## dns_name property — public_name / 011100231130 / 4

Type: `"string"`. Computed.

DNS Name. DNS Name

Upstream description:

DNS Name

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "characterSet": {
      "allowed": "[a-z0-9.-]",
      "description": "Dot-separated DNS labels"
    },
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "formatDescription": "RFC 1123 FQDN: lowercase, dot-separated labels, max 253 chars total",
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1,
    "pattern": "^([a-z0-9]([a-z0-9-]{0,61}[a-z0-9])?\\.)*[a-z0-9]([a-z0-9-]{0,61}[a-z0-9])?$",
    "validation": {
      "rfc": "RFC 1123"
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

<a id="canonical-0322012333020212-2201213102303100-0110303000131003-3110133030232133-3130323213132302-0203221102032132-3310003102322203-0021311323323030"></a>

<a id="canonical-0120100022031322-3211121221021130-1233113311120012-1302021202303000-3323231003001303-1133122132113132-0300201003100222-1031021212302022"></a>

## refresh_interval property — public_name / 011100231130 / 5

Type: `"number"`. Computed.

Interval for DNS refresh in seconds. Max value is 7 days as per
https&#58;//datatracker.ietf.org/doc/HTML/rfc8767.

Upstream description:

Interval for DNS refresh in seconds. Max value is 7 days as per
https&#58;//datatracker.ietf.org/doc/HTML/rfc8767.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 604800,
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
    "ves.io.schema.rules.uint32.ranges": "0,10-604800"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.ranges": "0,10-604800"
  }
}
```

<a id="canonical-3221030120212133-2121032100001330-1332223013133230-2331213112100311-1011323303110330-0322222310333103-0223121201030120-2230101330133110"></a>

## Next pages — public_name / 011100231130 / 6

- [origin_servers.origin_servers](data-sources--dns_proxy--reference--group-001.md#canonical-1333131311220333-1231120200230111-1310003132330122-1302102303223232-1331131102023222-3003000012212311-2300033331021112-0011131102103213)
- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-1122013202133033-3201201122310010-3303231012021212-2123120003213300-3130001133310203-2011013122222333-0130313333122120-0323311233013112)

<a id="canonical-2030003201020323-1323020011011112-2000103202210133-0303032330110302-0332002331302301-3130130232111302-3222033101030313-1331120330031102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1110122201100321-1232331301300012-3301212220322102-2112322013221220-3231032102331032-2001002212021232-0103323022111200-2003230011001310"></a>

## origin_servers.origin_servers.site_preferences — site_preferences / 133300223200 / 2

Breadcrumbs:

- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-1122013202133033-3201201122310010-3303231012021212-2123120003213300-3130001133310203-2011013122222333-0130313333122120-0323311233013112)
- [Property reference](data-sources--dns_proxy--reference--group-001.md#canonical-1133330113022100-3110333211021121-3111232100100000-3311232321023320-3202221331130133-0013122201302330-1103000312200021-1203111112100020)
- [origin_servers](data-sources--dns_proxy--reference--group-001.md#canonical-3102221313302120-2300133213200331-0320001332113013-1100133203331131-3231110233212130-1312212033030000-3302023133321213-2023021333201010)
- [origin_servers.origin_servers](data-sources--dns_proxy--reference--group-001.md#canonical-1333131311220333-1231120200230111-1310003132330122-1302102303223232-1331131102023222-3003000012212311-2300033331021112-0011131102103213)
- origin_servers.origin_servers.site_preferences

<a id="canonical-1311300121121300-3301310121221020-3210222301121133-2222020202302323-2133012221132103-3123003202312213-2203030022311023-1000133221022333"></a>

Type: `"single"`. Computed.

Carries the references to one or more sites.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2001311110120121-1012233200323331-1011010010200022-0320322000130132-1203103222212132-3232101022203312-1300232220230011-0303021101322112"></a>

## Direct properties — site_preferences / 133300223200 / 3

- [refs](data-sources--dns_proxy--reference--group-001.md#canonical-1012330211313113-0023200231222312-0002033131021220-1122201013200212-3011033222113113-0110310330102131-2100223210132131-3122312202000023): complete subsection reference.

<a id="canonical-2200021032210200-3113330103022111-3222020103203332-2230222212010103-3102000010331210-2200231332101113-3311031031131100-1223311113102100"></a>

## Next pages — site_preferences / 133300223200 / 4

- [origin_servers.origin_servers.site_preferences.refs](data-sources--dns_proxy--reference--group-001.md#canonical-1012330211313113-0023200231222312-0002033131021220-1122201013200212-3011033222113113-0110310330102131-2100223210132131-3122312202000023)
- [origin_servers.origin_servers](data-sources--dns_proxy--reference--group-001.md#canonical-1333131311220333-1231120200230111-1310003132330122-1302102303223232-1331131102023222-3003000012212311-2300033331021112-0011131102103213)
- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-1122013202133033-3201201122310010-3303231012021212-2123120003213300-3130001133310203-2011013122222333-0130313333122120-0323311233013112)

<a id="canonical-1012330211313113-0023200231222312-0002033131021220-1122201013200212-3011033222113113-0110310330102131-2100223210132131-3122312202000023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3221203133330320-3102012232203132-2213313123333002-3322223031113132-0031002313203132-2202030331002202-2332122322220320-3323302111221312"></a>

## origin_servers.origin_servers.site_preferences.refs — refs / 013230212312 / 2

Breadcrumbs:

- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-1122013202133033-3201201122310010-3303231012021212-2123120003213300-3130001133310203-2011013122222333-0130313333122120-0323311233013112)
- [Property reference](data-sources--dns_proxy--reference--group-001.md#canonical-1133330113022100-3110333211021121-3111232100100000-3311232321023320-3202221331130133-0013122201302330-1103000312200021-1203111112100020)
- [origin_servers](data-sources--dns_proxy--reference--group-001.md#canonical-3102221313302120-2300133213200331-0320001332113013-1100133203331131-3231110233212130-1312212033030000-3302023133321213-2023021333201010)
- [origin_servers.origin_servers](data-sources--dns_proxy--reference--group-001.md#canonical-1333131311220333-1231120200230111-1310003132330122-1302102303223232-1331131102023222-3003000012212311-2300033331021112-0011131102103213)
- [origin_servers.origin_servers.site_preferences](data-sources--dns_proxy--reference--group-001.md#canonical-2030003201020323-1323020011011112-2000103202210133-0303032330110302-0332002331302301-3130130232111302-3222033101030313-1331120330031102)
- origin_servers.origin_servers.site_preferences.refs

<a id="canonical-0130033232121321-1210302322100122-2210101203021123-0000021003200311-0102222203110210-2113101021313331-0021320123311022-1320332010013322"></a>

Type: `"list"`. Computed.

Site References. Reference to one or more sites.

Upstream description:

Reference to one or more sites.

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
    "ves.io.schema.rules.repeated.max_items": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "64"
  }
}
```

<a id="canonical-0321012000202320-0203111331102032-3202332332300331-2001300212320012-1113301010223330-0032010231302031-2221113321221233-3313012033133132"></a>

## Direct properties — refs / 013230212312 / 3

<a id="canonical-1112310223310102-2333030012300102-1222000330203232-0133301330110303-2101113102102003-2310113333022120-1323213012123330-3222031001030320"></a>

<a id="canonical-3021113130201231-3001033321012313-0120102111302032-1021230200022033-2321200330112132-3031112232203111-2022032103323310-1000300113000001"></a>

## name property — refs / 013230212312 / 4

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

<a id="canonical-2120130133212133-2321133020011032-2213003130302313-0200101120010013-1100210023301131-0231212311021223-1233013213130021-0121333121010321"></a>

<a id="canonical-1132100201231212-3230003303020120-1030102131331102-0322322230131103-1022233201302232-1123332320032003-3321311321320301-2330021022010313"></a>

## namespace property — refs / 013230212312 / 5

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

<a id="canonical-2102310230213331-1121231203003132-2333321112110221-3132202130002020-1230313232010302-1012320231113033-0333002122011322-1330031012021233"></a>

<a id="canonical-1323322302102002-3233033122210011-3021210332021220-3333122102230221-3102000003233302-3330123221202023-0000031220301032-3131331222321003"></a>

## tenant property — refs / 013230212312 / 6

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

<a id="canonical-1021123101311122-0022000210313331-0003321133321120-3323303020320223-2322100232201333-0113133223233021-2110021133110113-2010223332203223"></a>

## Next pages — refs / 013230212312 / 7

- [origin_servers.origin_servers.site_preferences](data-sources--dns_proxy--reference--group-001.md#canonical-2030003201020323-1323020011011112-2000103202210133-0303032330110302-0332002331302301-3130130232111302-3222033101030313-1331120330031102)
- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-1122013202133033-3201201122310010-3303231012021212-2123120003213300-3130001133310203-2011013122222333-0130313333122120-0323311233013112)

<a id="canonical-1222310103202331-3022213131203002-0003133222013103-3321032133113232-1300333320113032-0110233333210122-2321331132232223-2213213103001030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2301213101010130-3302230030330231-2321333002002113-2213010312031222-1023130002010332-1023113233302212-2121003312233233-0322313032303311"></a>

## protocol_inspection — protocol_inspection / 201123102213 / 2

Breadcrumbs:

- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-1122013202133033-3201201122310010-3303231012021212-2123120003213300-3130001133310203-2011013122222333-0130313333122120-0323311233013112)
- [Property reference](data-sources--dns_proxy--reference--group-001.md#canonical-1133330113022100-3110333211021121-3111232100100000-3311232321023320-3202221331130133-0013122201302330-1103000312200021-1203111112100020)
- protocol_inspection

<a id="canonical-0110322022112223-0222213023211202-1210123301332202-1131323120012022-1102200130013220-1010212300233210-3212330310023220-2231113331112133"></a>

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

<a id="canonical-1101022231122123-2322232020011333-0132100330120331-0110300033322310-1301011033313201-3211100003232301-1002203202132210-2121130133212112"></a>

## Direct properties — protocol_inspection / 201123102213 / 3

<a id="canonical-3221100310203000-1003122011103113-2221111123013120-3003002101012312-1110003301222321-1131023330200303-3023010020103111-1031331000122303"></a>

<a id="canonical-0112323211111123-1211202333232203-0000030000132303-0022002031232131-2131221203222113-2321320201221113-2013310322111022-0032013323303313"></a>

## name property — protocol_inspection / 201123102213 / 4

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

<a id="canonical-2222021032131300-3011220032311303-2300131032013033-3103202313110001-0322310302210120-3130233202012112-1000012213322222-2112220130021013"></a>

<a id="canonical-3101131003223012-1003313111211310-1313123332330033-1202012302312000-2222331002033323-2312201302131112-1032310223031320-3033202201133313"></a>

## namespace property — protocol_inspection / 201123102213 / 5

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

<a id="canonical-0123102233211003-3230300321332003-1332031211013233-0113213033310000-3023012302033213-0213022320231113-2120102110303311-2013011230302322"></a>

<a id="canonical-2321201011121300-3223030310103311-2031132332322223-1230203100002113-3222113103113103-3101201223011010-0303003030220201-2311210332101321"></a>

## tenant property — protocol_inspection / 201123102213 / 6

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

<a id="canonical-2312300321321003-3032222031103311-2302310232121132-2110220133300131-1203022103200132-3311032201203120-0332203030333223-3212111011033331"></a>

## Next pages — protocol_inspection / 201123102213 / 7

- [Property reference](data-sources--dns_proxy--reference--group-001.md#canonical-1133330113022100-3110333211021121-3111232100100000-3311232321023320-3202221331130133-0013122201302330-1103000312200021-1203111112100020)
- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-1122013202133033-3201201122310010-3303231012021212-2123120003213300-3130001133310203-2011013122222333-0130313333122120-0323311233013112)

<a id="canonical-1233321320320303-1032110303233312-2131113231101333-3032323203221001-0110213212120103-3322121223310111-0201121112313233-2310233121220133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1102103330200032-3121000302331303-2002103122210110-0010110330020121-3202002212203310-0030300132102132-2203211310122300-1121121102210210"></a>

## proxy_advertisement — proxy_advertisement / 100133110112 / 2

Breadcrumbs:

- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-1122013202133033-3201201122310010-3303231012021212-2123120003213300-3130001133310203-2011013122222333-0130313333122120-0323311233013112)
- [Property reference](data-sources--dns_proxy--reference--group-001.md#canonical-1133330113022100-3110333211021121-3111232100100000-3311232321023320-3202221331130133-0013122201302330-1103000312200021-1203111112100020)
- proxy_advertisement

<a id="canonical-1031132200030301-2123232013300210-2020123320233102-0130132220102021-3002201030310313-3323113323131301-1112231331321231-3321032123323210"></a>

Type: `"single"`. Computed.

Configuration parameter for proxy advertisement.

Upstream description:

Proxy Advertisement Type.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-advertise_choice": "[\"advertise_custom\",\"advertise_dualstack_on_public\",\"advertise_on_public\",\"advertise_on_public_default_dualstack_vip\",\"advertise_on_public_default_ipv6_vip\",\"advertise_on_public_default_vip\",\"advertise_v6_on_public\",\"do_not_advertise\"]"
}
```

<a id="canonical-1130101120022001-0302221223030210-2132322000222001-0121322030330331-2011233013120102-1112310230233113-0113221103121100-3023232310021130"></a>

## Direct properties — proxy_advertisement / 100133110112 / 3

- [advertise_custom](data-sources--dns_proxy--reference--group-001.md#canonical-0121320103123000-0212221121020201-1332101102231101-0230012010231220-1120001321320132-3232322222212322-3211300112301312-1201311231022101): complete subsection reference.

- [advertise_dualstack_on_public](data-sources--dns_proxy--reference--group-002.md#canonical-0201333130332013-0112020301131201-2021022313321201-3013233130311223-3022300202201021-2131301231003023-3121323331203223-2002102000332121): complete subsection reference.

- [advertise_on_public](data-sources--dns_proxy--reference--group-002.md#canonical-3310110220212000-2133311323002011-3330302201310131-3311331023300100-3031000200023101-0321033223132322-1202213132212101-2303002101010130): complete subsection reference.

- [advertise_on_public_default_dualstack_vip](data-sources--dns_proxy--reference--group-002.md#canonical-1131222002311300-2222212030001300-0012300103101232-0031023232101211-3302302310331123-2301120201231323-1120203233013132-2310220201330210): complete subsection reference.

- [advertise_on_public_default_ipv6_vip](data-sources--dns_proxy--reference--group-002.md#canonical-0010111020311300-1202110002311312-3111100031020231-3233312101103231-2321101300233120-2111220013030021-0203011312322031-2032301213311120): complete subsection reference.

- [advertise_on_public_default_vip](data-sources--dns_proxy--reference--group-002.md#canonical-1022322020332313-1201203221223320-0201230131212313-1100002123112013-2313010311202123-1200121102113312-1212301113321232-0123230022132320): complete subsection reference.

- [advertise_v6_on_public](data-sources--dns_proxy--reference--group-002.md#canonical-3023010033002312-1211102200100233-0013330212203322-0033130332302222-1020020000213233-0021211213221212-1221003001332210-3122311210303213): complete subsection reference.

- [do_not_advertise](data-sources--dns_proxy--reference--group-002.md#canonical-1012133330232223-3320102120220023-3002110212112031-0330112322100132-3202033312322102-3301100130033303-3200110031011233-3223231020123322): complete subsection reference.

<a id="canonical-0023011023103113-1330130033230012-1021320320132011-3100112112123000-1002010223123323-2323002330103001-1320301130102122-3133110301110113"></a>

## Next pages — proxy_advertisement / 100133110112 / 4

- [proxy_advertisement.advertise_custom](data-sources--dns_proxy--reference--group-001.md#canonical-0121320103123000-0212221121020201-1332101102231101-0230012010231220-1120001321320132-3232322222212322-3211300112301312-1201311231022101)
- [proxy_advertisement.advertise_dualstack_on_public](data-sources--dns_proxy--reference--group-002.md#canonical-0201333130332013-0112020301131201-2021022313321201-3013233130311223-3022300202201021-2131301231003023-3121323331203223-2002102000332121)
- [proxy_advertisement.advertise_on_public](data-sources--dns_proxy--reference--group-002.md#canonical-3310110220212000-2133311323002011-3330302201310131-3311331023300100-3031000200023101-0321033223132322-1202213132212101-2303002101010130)
- [proxy_advertisement.advertise_on_public_default_dualstack_vip](data-sources--dns_proxy--reference--group-002.md#canonical-1131222002311300-2222212030001300-0012300103101232-0031023232101211-3302302310331123-2301120201231323-1120203233013132-2310220201330210)
- [proxy_advertisement.advertise_on_public_default_ipv6_vip](data-sources--dns_proxy--reference--group-002.md#canonical-0010111020311300-1202110002311312-3111100031020231-3233312101103231-2321101300233120-2111220013030021-0203011312322031-2032301213311120)
- [proxy_advertisement.advertise_on_public_default_vip](data-sources--dns_proxy--reference--group-002.md#canonical-1022322020332313-1201203221223320-0201230131212313-1100002123112013-2313010311202123-1200121102113312-1212301113321232-0123230022132320)
- [proxy_advertisement.advertise_v6_on_public](data-sources--dns_proxy--reference--group-002.md#canonical-3023010033002312-1211102200100233-0013330212203322-0033130332302222-1020020000213233-0021211213221212-1221003001332210-3122311210303213)
- [proxy_advertisement.do_not_advertise](data-sources--dns_proxy--reference--group-002.md#canonical-1012133330232223-3320102120220023-3002110212112031-0330112322100132-3202033312322102-3301100130033303-3200110031011233-3223231020123322)
- [Property reference](data-sources--dns_proxy--reference--group-001.md#canonical-1133330113022100-3110333211021121-3111232100100000-3311232321023320-3202221331130133-0013122201302330-1103000312200021-1203111112100020)
- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-1122013202133033-3201201122310010-3303231012021212-2123120003213300-3130001133310203-2011013122222333-0130313333122120-0323311233013112)

<a id="canonical-0121320103123000-0212221121020201-1332101102231101-0230012010231220-1120001321320132-3232322222212322-3211300112301312-1201311231022101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1302033233123233-0003222303211332-2000020132121121-2233310131301311-3133133302202110-2231312323212000-1202311302130102-2301011313221320"></a>

## proxy_advertisement.advertise_custom — advertise_custom / 030101222202 / 2

Breadcrumbs:

- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-1122013202133033-3201201122310010-3303231012021212-2123120003213300-3130001133310203-2011013122222333-0130313333122120-0323311233013112)
- [Property reference](data-sources--dns_proxy--reference--group-001.md#canonical-1133330113022100-3110333211021121-3111232100100000-3311232321023320-3202221331130133-0013122201302330-1103000312200021-1203111112100020)
- [proxy_advertisement](data-sources--dns_proxy--reference--group-001.md#canonical-1233321320320303-1032110303233312-2131113231101333-3032323203221001-0110213212120103-3322121223310111-0201121112313233-2310233121220133)
- proxy_advertisement.advertise_custom

<a id="canonical-3330330022111103-2111320100301223-0213202312220003-3030303021230223-2131101010300223-1223112213001030-2110223233220201-1210020323032303"></a>

Type: `"single"`. Computed.

Defines a way to advertise a VIP on specific sites.

Upstream description:

This defines a way to advertise a VIP on specific sites.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0232310331100230-0311303033023030-2111111132111310-1322013333022220-1121130311120303-1200223023021333-2312030322001201-2002022212111301"></a>

## Direct properties — advertise_custom / 030101222202 / 3

- [advertise_where](data-sources--dns_proxy--reference--group-001.md#canonical-0310210230001331-3033021211312312-1002120201130001-1301322301121003-1122112322223303-1331210131110332-2313010230213100-2302012312313220): complete subsection reference.

<a id="canonical-0110031233110202-0000333303132003-3230320120023220-1331322333232130-1020010113130212-1003300333300132-1232032300200331-2323130331201220"></a>

## Next pages — advertise_custom / 030101222202 / 4

- [proxy_advertisement.advertise_custom.advertise_where](data-sources--dns_proxy--reference--group-001.md#canonical-0310210230001331-3033021211312312-1002120201130001-1301322301121003-1122112322223303-1331210131110332-2313010230213100-2302012312313220)
- [proxy_advertisement](data-sources--dns_proxy--reference--group-001.md#canonical-1233321320320303-1032110303233312-2131113231101333-3032323203221001-0110213212120103-3322121223310111-0201121112313233-2310233121220133)
- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-1122013202133033-3201201122310010-3303231012021212-2123120003213300-3130001133310203-2011013122222333-0130313333122120-0323311233013112)

<a id="canonical-0310210230001331-3033021211312312-1002120201130001-1301322301121003-1122112322223303-1331210131110332-2313010230213100-2302012312313220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2323310122011231-2221012233121310-1133200010103203-1322112123303312-1232220023011200-3000322231211031-1232222112020030-0122323322310132"></a>

## proxy_advertisement.advertise_custom.advertise_where — advertise_where / 311210300032 / 2

Breadcrumbs:

- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-1122013202133033-3201201122310010-3303231012021212-2123120003213300-3130001133310203-2011013122222333-0130313333122120-0323311233013112)
- [Property reference](data-sources--dns_proxy--reference--group-001.md#canonical-1133330113022100-3110333211021121-3111232100100000-3311232321023320-3202221331130133-0013122201302330-1103000312200021-1203111112100020)
- [proxy_advertisement](data-sources--dns_proxy--reference--group-001.md#canonical-1233321320320303-1032110303233312-2131113231101333-3032323203221001-0110213212120103-3322121223310111-0201121112313233-2310233121220133)
- [proxy_advertisement.advertise_custom](data-sources--dns_proxy--reference--group-001.md#canonical-0121320103123000-0212221121020201-1332101102231101-0230012010231220-1120001321320132-3232322222212322-3211300112301312-1201311231022101)
- proxy_advertisement.advertise_custom.advertise_where

<a id="canonical-0331120301030313-3011021013112113-0213322212210312-3021133012333232-3011100330302310-1021013321000311-2201103213133012-2110011202202300"></a>

Type: `"list"`. Computed.

Where should this load balancer be available.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 32,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 32,
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
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2312323121001210-0333100023211323-3100133210112102-2203233331100010-0313333011112030-3130110031100333-3011013132130302-3113131121333232"></a>

## Direct properties — advertise_where / 311210300032 / 3

- [advertise_dualstack_on_public](data-sources--dns_proxy--reference--group-001.md#canonical-0132123303300300-0302300023320120-1302232100300133-2220222110112113-0223322131221330-1300221223321322-0321220131002313-3321231233002031): complete subsection reference.

- [advertise_on_public](data-sources--dns_proxy--reference--group-002.md#canonical-3333022132212222-3111201032102213-1013002313023331-0213231031023210-2222300310332023-1013233020000121-0221011301231012-2133231303223103): complete subsection reference.

- [advertise_v6_on_public](data-sources--dns_proxy--reference--group-002.md#canonical-3011320001333033-3323010222033222-2312310300222031-3102010203100032-0322103312102132-1310220022003211-2201223332202331-3012320333320133): complete subsection reference.

<a id="canonical-1202132000203113-1033303203013103-0001213223220220-1231123023012210-2220010121321103-0200111001330310-3333220111103311-0103101333123330"></a>

<a id="canonical-0030311100310313-3000222310002003-1113133031301122-1021122331033001-3233310032201222-0031300223021100-2330303222113333-2023122112100201"></a>

## port property — advertise_where / 311210300032 / 4

Type: `"number"`. Computed.

Exclusive with \[port\_ranges use\_default\_port\] Port to Listen.

Upstream description:

Exclusive with \[port\_ranges use\_default\_port\] Port to Listen.

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

<a id="canonical-1121333312202023-2203311203212300-2302333220221022-1131302110200310-2130220312320130-2320000333320213-3201032122023202-2320221221123110"></a>

<a id="canonical-0132013003120202-3032220133231322-3021113312220300-3122112130130220-3120000030321022-3032313110132331-0031101310023311-2311003003221212"></a>

## port_ranges property — advertise_where / 311210300032 / 5

Type: `"string"`. Computed.

Exclusive with \[port use\_default\_port\] A string containing a comma separated list of port
ranges. Each port range consists of a single port or two ports separated by '-'.

Upstream description:

Exclusive with \[port use\_default\_port\] A string containing a comma separated list of port
ranges. Each port range consists of a single port or two ports separated by "-".

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 512,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 512,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.string.max_len": "512",
    "ves.io.schema.rules.string.max_ports": "64",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.unique_port_range_list": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "512",
    "ves.io.schema.rules.string.max_ports": "64",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.unique_port_range_list": "true"
  }
}
```

- [site](data-sources--dns_proxy--reference--group-002.md#canonical-1202013332121321-0333321230221233-0300031321311331-0202212330103032-2001113111023230-2302012002012120-1001301130310311-1312022200123212): complete subsection reference.

- [use_default_port](data-sources--dns_proxy--reference--group-002.md#canonical-1322010222011223-3333222022032231-0303010321312313-1213220133011033-0310121331102331-1013011301322221-0033302111332320-1211310333310013): complete subsection reference.

- [virtual_network](data-sources--dns_proxy--reference--group-002.md#canonical-1210321133310320-2021112303102212-2320133003123233-1211100103020321-2023323111103133-3210003112302123-0003332023102101-3033302113213130): complete subsection reference.

- [virtual_site](data-sources--dns_proxy--reference--group-002.md#canonical-0032203302302013-3111220213021010-0303003102222033-2220212203222302-1212122000221101-3333223131032300-0332333012122110-1013233212322320): complete subsection reference.

- [virtual_site_with_vip](data-sources--dns_proxy--reference--group-002.md#canonical-0033131110232301-2002233322320212-3111233001121322-2023001311103132-3232133223101330-2202030133010212-2111220111013323-0233220313313133): complete subsection reference.

- [vk8s_service](data-sources--dns_proxy--reference--group-002.md#canonical-2033011233111310-3303121203100330-1200131323233001-0011132112303132-2021223312002302-0232310330222110-2120123123021030-3030010023303130): complete subsection reference.

<a id="canonical-0101322030132013-0222300002221131-2031221322223013-0312021011123000-3122202121101003-1112203222133303-1000331202212303-2022302312223120"></a>

## Next pages — advertise_where / 311210300032 / 6

- [proxy_advertisement.advertise_custom.advertise_where.advertise_dualstack_on_public](data-sources--dns_proxy--reference--group-001.md#canonical-0132123303300300-0302300023320120-1302232100300133-2220222110112113-0223322131221330-1300221223321322-0321220131002313-3321231233002031)
- [proxy_advertisement.advertise_custom.advertise_where.advertise_on_public](data-sources--dns_proxy--reference--group-002.md#canonical-3333022132212222-3111201032102213-1013002313023331-0213231031023210-2222300310332023-1013233020000121-0221011301231012-2133231303223103)
- [proxy_advertisement.advertise_custom.advertise_where.advertise_v6_on_public](data-sources--dns_proxy--reference--group-002.md#canonical-3011320001333033-3323010222033222-2312310300222031-3102010203100032-0322103312102132-1310220022003211-2201223332202331-3012320333320133)
- [proxy_advertisement.advertise_custom.advertise_where.site](data-sources--dns_proxy--reference--group-002.md#canonical-1202013332121321-0333321230221233-0300031321311331-0202212330103032-2001113111023230-2302012002012120-1001301130310311-1312022200123212)
- [proxy_advertisement.advertise_custom.advertise_where.use_default_port](data-sources--dns_proxy--reference--group-002.md#canonical-1322010222011223-3333222022032231-0303010321312313-1213220133011033-0310121331102331-1013011301322221-0033302111332320-1211310333310013)
- [proxy_advertisement.advertise_custom.advertise_where.virtual_network](data-sources--dns_proxy--reference--group-002.md#canonical-1210321133310320-2021112303102212-2320133003123233-1211100103020321-2023323111103133-3210003112302123-0003332023102101-3033302113213130)
- [proxy_advertisement.advertise_custom.advertise_where.virtual_site](data-sources--dns_proxy--reference--group-002.md#canonical-0032203302302013-3111220213021010-0303003102222033-2220212203222302-1212122000221101-3333223131032300-0332333012122110-1013233212322320)
- [proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip](data-sources--dns_proxy--reference--group-002.md#canonical-0033131110232301-2002233322320212-3111233001121322-2023001311103132-3232133223101330-2202030133010212-2111220111013323-0233220313313133)
- [proxy_advertisement.advertise_custom.advertise_where.vk8s_service](data-sources--dns_proxy--reference--group-002.md#canonical-2033011233111310-3303121203100330-1200131323233001-0011132112303132-2021223312002302-0232310330222110-2120123123021030-3030010023303130)
- [proxy_advertisement.advertise_custom](data-sources--dns_proxy--reference--group-001.md#canonical-0121320103123000-0212221121020201-1332101102231101-0230012010231220-1120001321320132-3232322222212322-3211300112301312-1201311231022101)
- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-1122013202133033-3201201122310010-3303231012021212-2123120003213300-3130001133310203-2011013122222333-0130313333122120-0323311233013112)

<a id="canonical-0132123303300300-0302300023320120-1302232100300133-2220222110112113-0223322131221330-1300221223321322-0321220131002313-3321231233002031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->
