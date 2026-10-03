---
page_title: "xcsh_dns_lb_health_check reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_dns_lb_health_check reference."
---

# xcsh_dns_lb_health_check reference

<a id="canonical-3120123222202123-1012301220221110-0112203121332121-2102322302132322-2211003220321202-1031302101200212-1132031013203122-0302130033110030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1011201122001211-3332033221332221-0122112322320320-3223100001200002-2032102303012323-2113002221110232-2200121033112012-3210123133321330"></a>

## Property reference — Property reference / 301021012110 / 2

Breadcrumbs:

- [xcsh_dns_lb_health_check](../data-sources/dns_lb_health_check.md#canonical-3100122033112233-3012231223301113-3200002222023330-1201232003223131-2023031121303302-1301302210202321-2031302202212103-1331213200311221)
- Property reference

<a id="canonical-1321323312320312-1230200121002221-2113030133110111-3212213313231122-2001010322110210-0301312103123010-1123223232332001-2211202210203220"></a>

## Direct properties — Property reference / 301021012110 / 3

<a id="canonical-2131031031300123-1330320201331322-2302201211222222-3012221032330201-2130102011231032-3222133223320223-0232011323210023-1120031101222003"></a>

<a id="canonical-2022002200320012-2310123221003202-0200022231332201-3000112320100322-0021301322311320-0303130000220100-3122003020113032-1110213012222333"></a>

## annotations property — Property reference / 301021012110 / 4

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

<a id="canonical-1200311313232332-2213232223021100-3303021230122301-3123010112221331-3033103021121011-0301111100321210-3230033332333112-0010133031202120"></a>

<a id="canonical-1310230210120032-0100110121231123-0103012133101031-1103001033331313-1110321010033332-0212021123320203-0121223020020013-2003131030300302"></a>

## description property — Property reference / 301021012110 / 5

Type: `"string"`. Computed.

Description of the DNSLBHealthCheck.

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

- [http_health_check](data-sources--dns_lb_health_check--reference--group-001.md#canonical-3120202302210131-2013203312000310-1203131203130120-0120100221013232-1023312201211131-3223130332023123-1212032120133301-3322130133202012): complete subsection reference.

- [https_health_check](data-sources--dns_lb_health_check--reference--group-001.md#canonical-3131313330131332-3032110303033222-3102220010232323-3233202032230113-2132122013000013-0120022003012100-3300210122101101-3321103031102112): complete subsection reference.

- [icmp_health_check](data-sources--dns_lb_health_check--reference--group-001.md#canonical-2031001321110111-3201322332320002-2032300111031233-1331230303302132-2302302011221202-1003200033112211-3020333321013130-1221321211221033): complete subsection reference.

<a id="canonical-1233113022200123-2202300020111200-0210321100330022-3313100112030300-3330111033120030-0013113132212310-3200133100332211-3310132233001202"></a>

<a id="canonical-0102313022002200-2310201023121200-2202003020322113-2012203223322110-1332130332021120-3321121020323003-0210202210113013-1332032332220220"></a>

## ID property — Property reference / 301021012110 / 6

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-0111231022230002-1122000223310212-0010333023210131-1110003023210112-1231321230200210-2202001022321211-1032332120311033-3322320030130122"></a>

<a id="canonical-1213110302030132-3031233010211200-3303200313200012-2122000020332231-0310320030111002-3221333021023103-0202120302113302-3322301202023113"></a>

## labels property — Property reference / 301021012110 / 7

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

<a id="canonical-2211210221111200-0031321331232132-1122012230133311-1222201233131012-1033113222213012-1102221030131112-1032323110111121-3030301001222312"></a>

<a id="canonical-1302210122123111-3300111233331333-3222103312231312-0020330320331332-3200023023022132-0220211030001230-2122000131323221-0012300121232023"></a>

## name property — Property reference / 301021012110 / 8

Type: `"string"`. Required.

Name of the DNSLBHealthCheck.

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

<a id="canonical-0333020331203212-3102303132122211-0312221003330221-0221003233110131-2111103311113130-2210000221120022-0323032101030230-2310223112223130"></a>

<a id="canonical-2121010010030110-3320312001100023-2112002333313002-3301313130021030-2212002000120023-1121200230300103-1112101322120312-1100031322333111"></a>

## namespace property — Property reference / 301021012110 / 9

Type: `"string"`. Optional, Computed.

Namespace where the DNSLBHealthCheck exists.

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

- [tcp_health_check](data-sources--dns_lb_health_check--reference--group-001.md#canonical-3312322330200030-1232321021303111-3101103331113211-2113200131132022-1122310212123331-1332031232103330-0012301023331021-3103320231031323): complete subsection reference.

- [tcp_hex_health_check](data-sources--dns_lb_health_check--reference--group-001.md#canonical-2303133011132012-0032300330121222-0130321113203111-3133103302322001-3331022011323230-2300302212313100-3303032312332312-1223301313120331): complete subsection reference.

- [udp_health_check](data-sources--dns_lb_health_check--reference--group-001.md#canonical-1003110030133312-2333313002120220-2022222203021020-3311213231330220-2111301321213201-0000331121303100-0212210000033032-1003333133210312): complete subsection reference.

<a id="canonical-2103200311011012-3311320100101010-2002001010301121-2322323310020101-0320031331323222-3212220301331130-2312300130120122-1101223223030102"></a>

## All schema paths — Property reference / 301021012110 / 10

Each exact path has one authoritative reference destination. Collection element indices are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](data-sources--dns_lb_health_check--reference--group-001.md#canonical-2131031031300123-1330320201331322-2302201211222222-3012221032330201-2130102011231032-3222133223320223-0232011323210023-1120031101222003) |
| `description` | [description](data-sources--dns_lb_health_check--reference--group-001.md#canonical-1200311313232332-2213232223021100-3303021230122301-3123010112221331-3033103021121011-0301111100321210-3230033332333112-0010133031202120) |
| `http_health_check` | [http_health_check](data-sources--dns_lb_health_check--reference--group-001.md#canonical-3330000311113230-3031002222120111-2312101202220232-0333131212232013-3003232113110223-1032210113330322-1033133020300231-2210031232222131) |
| `http_health_check.disable_virtual_host` | [http_health_check.disable_virtual_host](data-sources--dns_lb_health_check--reference--group-001.md#canonical-0132000123022112-3120322321203200-0111110301202010-3322032031210020-1032112213222021-2132131200303223-0203020003331022-1330022200031011) |
| `http_health_check.health_check_port` | [http_health_check.health_check_port](data-sources--dns_lb_health_check--reference--group-001.md#canonical-3223101032130112-0111212322113001-1333020331120101-2100020033102322-3032322202000002-1203201103013031-1113123203220000-1101200213312010) |
| `http_health_check.health_check_secondary_port` | [http_health_check.health_check_secondary_port](data-sources--dns_lb_health_check--reference--group-001.md#canonical-3122131023020222-2031212320201011-1010032202220330-3303223012301010-3331131213113001-0131323123222223-2200201002231331-3223101110010323) |
| `http_health_check.inherit_load_balancer_fqdn` | [http_health_check.inherit_load_balancer_fqdn](data-sources--dns_lb_health_check--reference--group-001.md#canonical-0302023111020203-3313230033120313-2031221102321203-1302213232133313-3003023022330012-0023221000311233-0013300123313302-1021301132333223) |
| `http_health_check.receive` | [http_health_check.receive](data-sources--dns_lb_health_check--reference--group-001.md#canonical-0110130210002330-3130133132030231-0001002312323300-0320312023201221-1030020301022130-2200322233220210-3300222202120002-2133322133021010) |
| `http_health_check.send` | [http_health_check.send](data-sources--dns_lb_health_check--reference--group-001.md#canonical-3110112203222030-0120011001121332-3200123121033222-1311222011021012-0033212021110322-3211003203001111-1133313133032222-3132201010022033) |
| `http_health_check.virtual_host` | [http_health_check.virtual_host](data-sources--dns_lb_health_check--reference--group-001.md#canonical-2123312012123210-0213332112200110-1111023300211312-1211013232022011-2302012020221332-2021100022321122-3000012303032111-0111013320032301) |
| `https_health_check` | [https_health_check](data-sources--dns_lb_health_check--reference--group-001.md#canonical-1202002013023111-3001300030220321-3121221121203333-0312202310022100-3110222322303120-0303013130221021-1310032302101021-2330303031033012) |
| `https_health_check.disable_virtual_host` | [https_health_check.disable_virtual_host](data-sources--dns_lb_health_check--reference--group-001.md#canonical-1003221002321133-0321031303022323-3020313021130320-3120122300232012-2020121221330313-0113202213032230-2332201222230301-2103121302013232) |
| `https_health_check.health_check_port` | [https_health_check.health_check_port](data-sources--dns_lb_health_check--reference--group-001.md#canonical-3012120200320200-1032122030313021-3212031021323110-2022122103030332-0221133022002113-0102030001032210-1020300212320302-0030320113203210) |
| `https_health_check.health_check_secondary_port` | [https_health_check.health_check_secondary_port](data-sources--dns_lb_health_check--reference--group-001.md#canonical-0130221222200221-2113300220002023-1311021200012102-3300220200231223-0333031000223322-3300231003330332-1132132330010132-2103010010332021) |
| `https_health_check.inherit_load_balancer_fqdn` | [https_health_check.inherit_load_balancer_fqdn](data-sources--dns_lb_health_check--reference--group-001.md#canonical-2310321131210323-1121223220121222-2121310010222331-2111200213023011-2322301332223333-1012012302113222-0332321131010221-3122332032132333) |
| `https_health_check.receive` | [https_health_check.receive](data-sources--dns_lb_health_check--reference--group-001.md#canonical-2333201120013111-2301000221312013-2110022222312201-0032233230311000-1313323313103112-1222120100013312-2303330022202210-2201100012301232) |
| `https_health_check.send` | [https_health_check.send](data-sources--dns_lb_health_check--reference--group-001.md#canonical-3001001003123213-1120222313020322-2332220001320010-3131312123303212-1133211031023123-2331331110230132-0303213312022332-3200330300233023) |
| `https_health_check.virtual_host` | [https_health_check.virtual_host](data-sources--dns_lb_health_check--reference--group-001.md#canonical-3000031110211120-2320113000312322-2302112031213200-2130033001303010-3230213000221212-2131132022322132-0030322222031100-1020303113120321) |
| `icmp_health_check` | [icmp_health_check](data-sources--dns_lb_health_check--reference--group-001.md#canonical-1321231321300110-2233110121112311-3321213113333300-3201113022102101-1331022133200131-2001030322021222-0330100122322121-1123103132320103) |
| `id` | [ID](data-sources--dns_lb_health_check--reference--group-001.md#canonical-1233113022200123-2202300020111200-0210321100330022-3313100112030300-3330111033120030-0013113132212310-3200133100332211-3310132233001202) |
| `labels` | [labels](data-sources--dns_lb_health_check--reference--group-001.md#canonical-0111231022230002-1122000223310212-0010333023210131-1110003023210112-1231321230200210-2202001022321211-1032332120311033-3322320030130122) |
| `name` | [name](data-sources--dns_lb_health_check--reference--group-001.md#canonical-2211210221111200-0031321331232132-1122012230133311-1222201233131012-1033113222213012-1102221030131112-1032323110111121-3030301001222312) |
| `namespace` | [namespace](data-sources--dns_lb_health_check--reference--group-001.md#canonical-0333020331203212-3102303132122211-0312221003330221-0221003233110131-2111103311113130-2210000221120022-0323032101030230-2310223112223130) |
| `tcp_health_check` | [tcp_health_check](data-sources--dns_lb_health_check--reference--group-001.md#canonical-0220021320200322-2101330233302312-3231330303221312-0310022032033012-2210203320220201-2330212033301022-0033112123000121-3201200003202203) |
| `tcp_health_check.health_check_port` | [tcp_health_check.health_check_port](data-sources--dns_lb_health_check--reference--group-001.md#canonical-1120310012121030-0301300012210331-1230101010322101-2220333123323133-2331023300111031-3130121312200321-2202110132010311-2330320120002130) |
| `tcp_health_check.health_check_secondary_port` | [tcp_health_check.health_check_secondary_port](data-sources--dns_lb_health_check--reference--group-001.md#canonical-0002102013231001-2312231032232011-3202121313321031-1023231302100103-2311332330213022-1122313032213310-3322232110121122-2100303130321103) |
| `tcp_health_check.receive` | [tcp_health_check.receive](data-sources--dns_lb_health_check--reference--group-001.md#canonical-0320022033030000-0000101013303303-2131313311000002-0103021113122200-2101112223213010-1323312210120323-1302111132223313-1032220030311303) |
| `tcp_health_check.send` | [tcp_health_check.send](data-sources--dns_lb_health_check--reference--group-001.md#canonical-2202001033000311-0010321033320033-0002220301003113-2031230021012211-0201303230130233-1312102002233333-3330301203203130-3002222332201010) |
| `tcp_hex_health_check` | [tcp_hex_health_check](data-sources--dns_lb_health_check--reference--group-001.md#canonical-0121210010121103-1223321111223031-1023300133012310-1133122223220133-3310102210222002-0112100230310130-0223203310223333-3221000113220121) |
| `tcp_hex_health_check.health_check_port` | [tcp_hex_health_check.health_check_port](data-sources--dns_lb_health_check--reference--group-001.md#canonical-0201013213212312-1123023310223123-1013123231232232-2310010122101012-3130233013001123-3212210121321221-2031320333000322-0211031233133030) |
| `tcp_hex_health_check.health_check_secondary_port` | [tcp_hex_health_check.health_check_secondary_port](data-sources--dns_lb_health_check--reference--group-001.md#canonical-3232212222030201-3202103211232210-1202123322200223-2022102300113000-0133003013312002-3022320113210033-2031031210023223-1301303302003021) |
| `tcp_hex_health_check.receive` | [tcp_hex_health_check.receive](data-sources--dns_lb_health_check--reference--group-001.md#canonical-0300022111112323-3312002021122313-2002213001021021-3310310033222210-2101000300212123-2112210230200013-2101322232101203-3311012100010101) |
| `tcp_hex_health_check.send` | [tcp_hex_health_check.send](data-sources--dns_lb_health_check--reference--group-001.md#canonical-3002232302221113-3322013231101020-3333131133200230-0213330120312223-3322132203332222-0012300330310130-1112131112102123-1312201013231112) |
| `udp_health_check` | [udp_health_check](data-sources--dns_lb_health_check--reference--group-001.md#canonical-0332303000012330-0113320233211013-2113221201312020-0130010210030223-1120022123020203-0210002131333002-2313031011222300-1131130002110020) |
| `udp_health_check.health_check_port` | [udp_health_check.health_check_port](data-sources--dns_lb_health_check--reference--group-001.md#canonical-1001212233231002-1002313130130100-0200211222130210-1331000200130133-2120030221133310-2103120200310302-3300310300012332-3030033303121231) |
| `udp_health_check.health_check_secondary_port` | [udp_health_check.health_check_secondary_port](data-sources--dns_lb_health_check--reference--group-001.md#canonical-3333110010321330-0231033230231200-0211023003030133-3112313111313313-3000000100122101-1110200302001023-0311233022022021-2221031202301313) |
| `udp_health_check.receive` | [udp_health_check.receive](data-sources--dns_lb_health_check--reference--group-001.md#canonical-0230231131333201-0033022311023232-1222211133030311-0031123112100211-3101133010212311-1110032003020223-3333031311203100-2202322100130133) |
| `udp_health_check.send` | [udp_health_check.send](data-sources--dns_lb_health_check--reference--group-001.md#canonical-3002123021130131-2220133311320122-0032021323221332-2213001032122031-2112033323320013-3031021300120031-1201302032121000-1021121221333331) |

<a id="canonical-0113223331003312-0313303221002102-0112212301120131-2330321330012302-0212123033211310-3212320231201013-2211021033031120-1112222301330012"></a>

## Next pages — Property reference / 301021012110 / 11

- [http_health_check](data-sources--dns_lb_health_check--reference--group-001.md#canonical-3120202302210131-2013203312000310-1203131203130120-0120100221013232-1023312201211131-3223130332023123-1212032120133301-3322130133202012)
- [https_health_check](data-sources--dns_lb_health_check--reference--group-001.md#canonical-3131313330131332-3032110303033222-3102220010232323-3233202032230113-2132122013000013-0120022003012100-3300210122101101-3321103031102112)
- [icmp_health_check](data-sources--dns_lb_health_check--reference--group-001.md#canonical-2031001321110111-3201322332320002-2032300111031233-1331230303302132-2302302011221202-1003200033112211-3020333321013130-1221321211221033)
- [tcp_health_check](data-sources--dns_lb_health_check--reference--group-001.md#canonical-3312322330200030-1232321021303111-3101103331113211-2113200131132022-1122310212123331-1332031232103330-0012301023331021-3103320231031323)
- [tcp_hex_health_check](data-sources--dns_lb_health_check--reference--group-001.md#canonical-2303133011132012-0032300330121222-0130321113203111-3133103302322001-3331022011323230-2300302212313100-3303032312332312-1223301313120331)
- [udp_health_check](data-sources--dns_lb_health_check--reference--group-001.md#canonical-1003110030133312-2333313002120220-2022222203021020-3311213231330220-2111301321213201-0000331121303100-0212210000033032-1003333133210312)
- [xcsh_dns_lb_health_check](../data-sources/dns_lb_health_check.md#canonical-3100122033112233-3012231223301113-3200002222023330-1201232003223131-2023031121303302-1301302210202321-2031302202212103-1331213200311221)

<a id="canonical-3120202302210131-2013203312000310-1203131203130120-0120100221013232-1023312201211131-3223130332023123-1212032120133301-3322130133202012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0203013031131302-2133332111000222-0313200012003102-1310320120111302-3022130232330103-3222212302120303-1311212031100131-3323332033322121"></a>

## http_health_check — http_health_check / 302221023203 / 2

Breadcrumbs:

- [xcsh_dns_lb_health_check](../data-sources/dns_lb_health_check.md#canonical-3100122033112233-3012231223301113-3200002222023330-1201232003223131-2023031121303302-1301302210202321-2031302202212103-1331213200311221)
- [Property reference](data-sources--dns_lb_health_check--reference--group-001.md#canonical-3120123222202123-1012301220221110-0112203121332121-2102322302132322-2211003220321202-1031302101200212-1132031013203122-0302130033110030)
- http_health_check

<a id="canonical-3330000311113230-3031002222120111-2312101202220232-0333131212232013-3003232113110223-1032210113330322-1033133020300231-2210031232222131"></a>

Type: `"single"`. Computed.

\[OneOf: http\_health\_check, https\_health\_check, icmp\_health\_check, tcp\_health\_check,
tcp\_hex\_health\_check, udp\_health\_check\] Configuration parameter for http health check.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-virtual_host_choice": "[\"disable_virtual_host\",\"inherit_load_balancer_fqdn\",\"virtual_host\"]"
}
```

OneOf alternatives in this subsection:

- [http_health_check](data-sources--dns_lb_health_check--reference--group-001.md#canonical-3330000311113230-3031002222120111-2312101202220232-0333131212232013-3003232113110223-1032210113330322-1033133020300231-2210031232222131)
- [https_health_check](data-sources--dns_lb_health_check--reference--group-001.md#canonical-1202002013023111-3001300030220321-3121221121203333-0312202310022100-3110222322303120-0303013130221021-1310032302101021-2330303031033012)
- [icmp_health_check](data-sources--dns_lb_health_check--reference--group-001.md#canonical-1321231321300110-2233110121112311-3321213113333300-3201113022102101-1331022133200131-2001030322021222-0330100122322121-1123103132320103)
- [tcp_health_check](data-sources--dns_lb_health_check--reference--group-001.md#canonical-0220021320200322-2101330233302312-3231330303221312-0310022032033012-2210203320220201-2330212033301022-0033112123000121-3201200003202203)
- [tcp_hex_health_check](data-sources--dns_lb_health_check--reference--group-001.md#canonical-0121210010121103-1223321111223031-1023300133012310-1133122223220133-3310102210222002-0112100230310130-0223203310223333-3221000113220121)
- [udp_health_check](data-sources--dns_lb_health_check--reference--group-001.md#canonical-0332303000012330-0113320233211013-2113221201312020-0130010210030223-1120022123020203-0210002131333002-2313031011222300-1131130002110020)

Select alternatives according to the provider validators above.

<a id="canonical-0112200322332212-3122220313123101-3030333023010213-0230332111132011-0013031333302330-2012110123113232-1101320002310001-3030322113201210"></a>

## Direct properties — http_health_check / 302221023203 / 3

- [disable_virtual_host](data-sources--dns_lb_health_check--reference--group-001.md#canonical-2220331332012210-0111333213302200-0110121003302022-1123310200133231-0200221202022320-3333013211130303-3101121011312033-2221223113312112): complete subsection reference.

<a id="canonical-3223101032130112-0111212322113001-1333020331120101-2100020033102322-3032322202000002-1203201103013031-1113123203220000-1101200213312010"></a>

<a id="canonical-3120332022202211-0111213212210000-0312131212011122-0202232321233222-2010113300222210-3020012223302313-3330213323021132-2013112123331012"></a>

## health_check_port property — http_health_check / 302221023203 / 4

Type: `"number"`. Computed.

Health Check Port. Port used for performing health check.

Upstream description:

Port used for performing health check.

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
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-3122131023020222-2031212320201011-1010032202220330-3303223012301010-3331131213113001-0131323123222223-2200201002231331-3223101110010323"></a>

<a id="canonical-2113322001013312-3302101001033200-1122222103320200-2022013221311331-1121103120302212-3303131112310133-3303132200012232-2302003102003321"></a>

## health_check_secondary_port property — http_health_check / 302221023203 / 5

Type: `"number"`. Computed.

Secondary port used for performing health check. If included, both ports must be healthy for the
health check to pass.

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
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

- [inherit_load_balancer_fqdn](data-sources--dns_lb_health_check--reference--group-001.md#canonical-0330332110210230-2100032203010033-0123232220110122-0102333322301132-1210323100210233-2002210103320320-0001331031312211-1010102203123131): complete subsection reference.

<a id="canonical-0110130210002330-3130133132030231-0001002312323300-0320312023201221-1030020301022130-2200322233220210-3300222202120002-2133322133021010"></a>

<a id="canonical-1232320221023110-1322001100313311-0312021003202212-2010003023223100-1210220210030120-1103233230113310-0333320031213111-2132000102103322"></a>

## receive property — http_health_check / 302221023203 / 6

Type: `"string"`. Computed.

Regular expression used to match against the response to the health check's request. Mark node up
upon receipt of a successful regular expression match. Uses re2 regular expression syntax.

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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "2048",
    "ves.io.schema.rules.string.regex": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "2048",
    "ves.io.schema.rules.string.regex": "true"
  }
}
```

<a id="canonical-3110112203222030-0120011001121332-3200123121033222-1311222011021012-0033212021110322-3211003203001111-1133313133032222-3132201010022033"></a>

<a id="canonical-1133200310200100-3203122100002011-1100030012321220-0333010213232023-1230112300202221-0313301300111213-1211131111132101-3303322011133302"></a>

## send property — http_health_check / 302221023203 / 7

Type: `"string"`. Computed.

Send String. HTTP payload to send to the target.

Upstream description:

HTTP payload to send to the target.

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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "2048"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "2048"
  }
}
```

<a id="canonical-2123312012123210-0213332112200110-1111023300211312-1211013232022011-2302012020221332-2021100022321122-3000012303032111-0111013320032301"></a>

<a id="canonical-1312012022330113-3101101110001323-2330133213033032-1123110312003113-2311301211233212-2223212131312012-1100011233303322-3123131233213023"></a>

## virtual_host property — http_health_check / 302221023203 / 8

Type: `"string"`. Computed.

Exclusive with \[disable\_virtual\_host inherit\_load\_balancer\_fqdn\] Name of the virtual host to
use for SNI.

Upstream description:

Exclusive with \[disable\_virtual\_host inherit\_load\_balancer\_fqdn\] Name of the virtual host to
use for SNI.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 2048,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "maxLength": 2048,
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
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "2048"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "2048"
  }
}
```

<a id="canonical-1111213011110021-0230023210122100-3021311200003212-0120311230023302-3003102223110223-3301223203000100-1002333300332300-3200213102113000"></a>

## Next pages — http_health_check / 302221023203 / 9

- [http_health_check.disable_virtual_host](data-sources--dns_lb_health_check--reference--group-001.md#canonical-2220331332012210-0111333213302200-0110121003302022-1123310200133231-0200221202022320-3333013211130303-3101121011312033-2221223113312112)
- [http_health_check.inherit_load_balancer_fqdn](data-sources--dns_lb_health_check--reference--group-001.md#canonical-0330332110210230-2100032203010033-0123232220110122-0102333322301132-1210323100210233-2002210103320320-0001331031312211-1010102203123131)
- [Property reference](data-sources--dns_lb_health_check--reference--group-001.md#canonical-3120123222202123-1012301220221110-0112203121332121-2102322302132322-2211003220321202-1031302101200212-1132031013203122-0302130033110030)
- [xcsh_dns_lb_health_check](../data-sources/dns_lb_health_check.md#canonical-3100122033112233-3012231223301113-3200002222023330-1201232003223131-2023031121303302-1301302210202321-2031302202212103-1331213200311221)

<a id="canonical-2220331332012210-0111333213302200-0110121003302022-1123310200133231-0200221202022320-3333013211130303-3101121011312033-2221223113312112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1331020112001201-0131022300132302-0330000211332321-1113221321331213-1233221311100132-0211330321202022-2100020303130010-2301023301223130"></a>

## http_health_check.disable_virtual_host — disable_virtual_host / 331123023022 / 2

Breadcrumbs:

- [xcsh_dns_lb_health_check](../data-sources/dns_lb_health_check.md#canonical-3100122033112233-3012231223301113-3200002222023330-1201232003223131-2023031121303302-1301302210202321-2031302202212103-1331213200311221)
- [Property reference](data-sources--dns_lb_health_check--reference--group-001.md#canonical-3120123222202123-1012301220221110-0112203121332121-2102322302132322-2211003220321202-1031302101200212-1132031013203122-0302130033110030)
- [http_health_check](data-sources--dns_lb_health_check--reference--group-001.md#canonical-3120202302210131-2013203312000310-1203131203130120-0120100221013232-1023312201211131-3223130332023123-1212032120133301-3322130133202012)
- http_health_check.disable_virtual_host

<a id="canonical-0132000123022112-3120322321203200-0111110301202010-3322032031210020-1032112213222021-2132131200303223-0203020003331022-1330022200031011"></a>

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

<a id="canonical-0333100010112311-3211302323113002-1132330233332313-0223130322111301-3010131202200202-2203231013331010-2233333321313321-1022223202210332"></a>

## Direct properties — disable_virtual_host / 331123023022 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1021333201101101-2021303232020210-1233202011212301-3323021112002303-3231211212220301-2022003322132110-3000322123023232-0103112330233302"></a>

## Next pages — disable_virtual_host / 331123023022 / 4

- [http_health_check](data-sources--dns_lb_health_check--reference--group-001.md#canonical-3120202302210131-2013203312000310-1203131203130120-0120100221013232-1023312201211131-3223130332023123-1212032120133301-3322130133202012)
- [xcsh_dns_lb_health_check](../data-sources/dns_lb_health_check.md#canonical-3100122033112233-3012231223301113-3200002222023330-1201232003223131-2023031121303302-1301302210202321-2031302202212103-1331213200311221)

<a id="canonical-0330332110210230-2100032203010033-0123232220110122-0102333322301132-1210323100210233-2002210103320320-0001331031312211-1010102203123131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2131211321123232-1203103203333211-1332023113120101-2332313130001120-2023120222131320-2131301201311011-1102313201033210-1323203130303100"></a>

## http_health_check.inherit_load_balancer_fqdn — inherit_load_balancer_fqdn / 100023123131 / 2

Breadcrumbs:

- [xcsh_dns_lb_health_check](../data-sources/dns_lb_health_check.md#canonical-3100122033112233-3012231223301113-3200002222023330-1201232003223131-2023031121303302-1301302210202321-2031302202212103-1331213200311221)
- [Property reference](data-sources--dns_lb_health_check--reference--group-001.md#canonical-3120123222202123-1012301220221110-0112203121332121-2102322302132322-2211003220321202-1031302101200212-1132031013203122-0302130033110030)
- [http_health_check](data-sources--dns_lb_health_check--reference--group-001.md#canonical-3120202302210131-2013203312000310-1203131203130120-0120100221013232-1023312201211131-3223130332023123-1212032120133301-3322130133202012)
- http_health_check.inherit_load_balancer_fqdn

<a id="canonical-0302023111020203-3313230033120313-2031221102321203-1302213232133313-3003023022330012-0023221000311233-0013300123313302-1021301132333223"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for inherit load balancer fqdn.

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

<a id="canonical-3012110220033103-2103121102133120-3213201202321312-3131130333110213-0100223303021333-3212212021231233-1333002202020223-0120321320102131"></a>

## Direct properties — inherit_load_balancer_fqdn / 100023123131 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1313131200213113-3103231020211202-1200220223112103-2123332111230031-2012121212100201-1313210101102202-2010000002100101-0200303000202132"></a>

## Next pages — inherit_load_balancer_fqdn / 100023123131 / 4

- [http_health_check](data-sources--dns_lb_health_check--reference--group-001.md#canonical-3120202302210131-2013203312000310-1203131203130120-0120100221013232-1023312201211131-3223130332023123-1212032120133301-3322130133202012)
- [xcsh_dns_lb_health_check](../data-sources/dns_lb_health_check.md#canonical-3100122033112233-3012231223301113-3200002222023330-1201232003223131-2023031121303302-1301302210202321-2031302202212103-1331213200311221)

<a id="canonical-3131313330131332-3032110303033222-3102220010232323-3233202032230113-2132122013000013-0120022003012100-3300210122101101-3321103031102112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3123133201323030-1122303002133103-0312121012002100-0323031331221032-3320220013233031-0021320203033221-3131032122330111-0021032132323012"></a>

## https_health_check — https_health_check / 121012231023 / 2

Breadcrumbs:

- [xcsh_dns_lb_health_check](../data-sources/dns_lb_health_check.md#canonical-3100122033112233-3012231223301113-3200002222023330-1201232003223131-2023031121303302-1301302210202321-2031302202212103-1331213200311221)
- [Property reference](data-sources--dns_lb_health_check--reference--group-001.md#canonical-3120123222202123-1012301220221110-0112203121332121-2102322302132322-2211003220321202-1031302101200212-1132031013203122-0302130033110030)
- https_health_check

<a id="canonical-1202002013023111-3001300030220321-3121221121203333-0312202310022100-3110222322303120-0303013130221021-1310032302101021-2330303031033012"></a>

Type: `"single"`. Computed.

Configuration parameter for https health check.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-virtual_host_choice": "[\"disable_virtual_host\",\"inherit_load_balancer_fqdn\",\"virtual_host\"]"
}
```

<a id="canonical-2300013230120313-0330023003003033-2012221323112113-3032331232212330-3310303003232113-0321000112102110-0001002301013022-3113202003110310"></a>

## Direct properties — https_health_check / 121012231023 / 3

- [disable_virtual_host](data-sources--dns_lb_health_check--reference--group-001.md#canonical-3032023230023312-3113300231210103-0230333301321322-3321313200200133-3001233021210131-3221302102201133-3313331012310302-3232232023102330): complete subsection reference.

<a id="canonical-3012120200320200-1032122030313021-3212031021323110-2022122103030332-0221133022002113-0102030001032210-1020300212320302-0030320113203210"></a>

<a id="canonical-3132300232310123-1030101101113203-2013110131322230-0322212332303312-2032322100133131-3100002110213031-3322313210322022-1133303222023130"></a>

## health_check_port property — https_health_check / 121012231023 / 4

Type: `"number"`. Computed.

Health Check Port. Port used for performing health check.

Upstream description:

Port used for performing health check.

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
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-0130221222200221-2113300220002023-1311021200012102-3300220200231223-0333031000223322-3300231003330332-1132132330010132-2103010010332021"></a>

<a id="canonical-0003330210331111-2202030202313330-1130130220012033-0300123123003102-3010303123313221-2002113110202231-3003232133210010-0310133121330213"></a>

## health_check_secondary_port property — https_health_check / 121012231023 / 5

Type: `"number"`. Computed.

Secondary port used for performing health check. If included, both ports must be healthy for the
health check to pass.

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
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

- [inherit_load_balancer_fqdn](data-sources--dns_lb_health_check--reference--group-001.md#canonical-2131010110313022-3330112232023122-0021221130110333-0020122020231031-3231110023332300-3222220201000013-3121222203233021-1322320113302221): complete subsection reference.

<a id="canonical-2333201120013111-2301000221312013-2110022222312201-0032233230311000-1313323313103112-1222120100013312-2303330022202210-2201100012301232"></a>

<a id="canonical-2011211121021020-0211221002002231-1001303000132301-3212332210133233-2131331202033133-0120332320121102-2011333033011102-1232123231101032"></a>

## receive property — https_health_check / 121012231023 / 6

Type: `"string"`. Computed.

Regular expression used to match against the response to the health check's request. Mark node up
upon receipt of a successful regular expression match. Uses re2 regular expression syntax.

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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "2048",
    "ves.io.schema.rules.string.regex": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "2048",
    "ves.io.schema.rules.string.regex": "true"
  }
}
```

<a id="canonical-3001001003123213-1120222313020322-2332220001320010-3131312123303212-1133211031023123-2331331110230132-0303213312022332-3200330300233023"></a>

<a id="canonical-2002002303031101-1202011333000323-0002022101321223-2023132230000102-2103231232220121-0112130313310022-0302023232103311-2002121002011013"></a>

## send property — https_health_check / 121012231023 / 7

Type: `"string"`. Computed.

Send String. HTTP payload to send to the target.

Upstream description:

HTTP payload to send to the target.

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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "2048"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "2048"
  }
}
```

<a id="canonical-3000031110211120-2320113000312322-2302112031213200-2130033001303010-3230213000221212-2131132022322132-0030322222031100-1020303113120321"></a>

<a id="canonical-2333333010020002-0303231232333013-3303101321220013-1201323110000001-3222213013020320-0110230233031211-3302113012122132-3202302020200122"></a>

## virtual_host property — https_health_check / 121012231023 / 8

Type: `"string"`. Computed.

Exclusive with \[disable\_virtual\_host inherit\_load\_balancer\_fqdn\] Name of the virtual host to
use for SNI.

Upstream description:

Exclusive with \[disable\_virtual\_host inherit\_load\_balancer\_fqdn\] Name of the virtual host to
use for SNI.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 2048,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "maxLength": 2048,
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
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "2048"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "2048"
  }
}
```

<a id="canonical-0230320323101123-1222002233212122-1031030002100000-2110130320002201-3301123110302111-2030013122130233-1213100000201310-3120232210130330"></a>

## Next pages — https_health_check / 121012231023 / 9

- [https_health_check.disable_virtual_host](data-sources--dns_lb_health_check--reference--group-001.md#canonical-3032023230023312-3113300231210103-0230333301321322-3321313200200133-3001233021210131-3221302102201133-3313331012310302-3232232023102330)
- [https_health_check.inherit_load_balancer_fqdn](data-sources--dns_lb_health_check--reference--group-001.md#canonical-2131010110313022-3330112232023122-0021221130110333-0020122020231031-3231110023332300-3222220201000013-3121222203233021-1322320113302221)
- [Property reference](data-sources--dns_lb_health_check--reference--group-001.md#canonical-3120123222202123-1012301220221110-0112203121332121-2102322302132322-2211003220321202-1031302101200212-1132031013203122-0302130033110030)
- [xcsh_dns_lb_health_check](../data-sources/dns_lb_health_check.md#canonical-3100122033112233-3012231223301113-3200002222023330-1201232003223131-2023031121303302-1301302210202321-2031302202212103-1331213200311221)

<a id="canonical-3032023230023312-3113300231210103-0230333301321322-3321313200200133-3001233021210131-3221302102201133-3313331012310302-3232232023102330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1233303001200211-0311100032121303-3001103230010203-0013012200120312-3200020001013322-0122220213121002-2200213321033102-1132020323300231"></a>

## https_health_check.disable_virtual_host — disable_virtual_host / 103223221330 / 2

Breadcrumbs:

- [xcsh_dns_lb_health_check](../data-sources/dns_lb_health_check.md#canonical-3100122033112233-3012231223301113-3200002222023330-1201232003223131-2023031121303302-1301302210202321-2031302202212103-1331213200311221)
- [Property reference](data-sources--dns_lb_health_check--reference--group-001.md#canonical-3120123222202123-1012301220221110-0112203121332121-2102322302132322-2211003220321202-1031302101200212-1132031013203122-0302130033110030)
- [https_health_check](data-sources--dns_lb_health_check--reference--group-001.md#canonical-3131313330131332-3032110303033222-3102220010232323-3233202032230113-2132122013000013-0120022003012100-3300210122101101-3321103031102112)
- https_health_check.disable_virtual_host

<a id="canonical-1003221002321133-0321031303022323-3020313021130320-3120122300232012-2020121221330313-0113202213032230-2332201222230301-2103121302013232"></a>

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

<a id="canonical-1120202032233103-1301331132023031-3210310100310220-3120000321312103-0011311223033122-2002213331321321-3332032302311330-1310030331003312"></a>

## Direct properties — disable_virtual_host / 103223221330 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0221332112312022-0323333010331103-3030212222131100-1030312032222012-3030210121120331-2331331202202112-0301102200222100-1121000002300222"></a>

## Next pages — disable_virtual_host / 103223221330 / 4

- [https_health_check](data-sources--dns_lb_health_check--reference--group-001.md#canonical-3131313330131332-3032110303033222-3102220010232323-3233202032230113-2132122013000013-0120022003012100-3300210122101101-3321103031102112)
- [xcsh_dns_lb_health_check](../data-sources/dns_lb_health_check.md#canonical-3100122033112233-3012231223301113-3200002222023330-1201232003223131-2023031121303302-1301302210202321-2031302202212103-1331213200311221)

<a id="canonical-2131010110313022-3330112232023122-0021221130110333-0020122020231031-3231110023332300-3222220201000013-3121222203233021-1322320113302221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0223221232010201-0122203201200202-2331032221202022-1211102232331321-1110321023003131-0301103222023000-2003300021130100-3130220003333200"></a>

## https_health_check.inherit_load_balancer_fqdn — inherit_load_balancer_fqdn / 100002033030 / 2

Breadcrumbs:

- [xcsh_dns_lb_health_check](../data-sources/dns_lb_health_check.md#canonical-3100122033112233-3012231223301113-3200002222023330-1201232003223131-2023031121303302-1301302210202321-2031302202212103-1331213200311221)
- [Property reference](data-sources--dns_lb_health_check--reference--group-001.md#canonical-3120123222202123-1012301220221110-0112203121332121-2102322302132322-2211003220321202-1031302101200212-1132031013203122-0302130033110030)
- [https_health_check](data-sources--dns_lb_health_check--reference--group-001.md#canonical-3131313330131332-3032110303033222-3102220010232323-3233202032230113-2132122013000013-0120022003012100-3300210122101101-3321103031102112)
- https_health_check.inherit_load_balancer_fqdn

<a id="canonical-2310321131210323-1121223220121222-2121310010222331-2111200213023011-2322301332223333-1012012302113222-0332321131010221-3122332032132333"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for inherit load balancer fqdn.

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

<a id="canonical-3213311322310131-2202230132020130-3210032111310301-3311313332002113-0031021233132231-2101300231102022-0131310302332221-2003233012303313"></a>

## Direct properties — inherit_load_balancer_fqdn / 100002033030 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3000330213111212-2012132002302302-1010322312023032-3221132023120210-2001230003311100-0300100203122323-3111320331333100-0102011200120233"></a>

## Next pages — inherit_load_balancer_fqdn / 100002033030 / 4

- [https_health_check](data-sources--dns_lb_health_check--reference--group-001.md#canonical-3131313330131332-3032110303033222-3102220010232323-3233202032230113-2132122013000013-0120022003012100-3300210122101101-3321103031102112)
- [xcsh_dns_lb_health_check](../data-sources/dns_lb_health_check.md#canonical-3100122033112233-3012231223301113-3200002222023330-1201232003223131-2023031121303302-1301302210202321-2031302202212103-1331213200311221)

<a id="canonical-2031001321110111-3201322332320002-2032300111031233-1331230303302132-2302302011221202-1003200033112211-3020333321013130-1221321211221033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2133330030303113-3122330223212113-1321120122133110-2321313201112231-2213221132030232-3130011202110013-2302111112013203-2020202130121331"></a>

## icmp_health_check — icmp_health_check / 212002112121 / 2

Breadcrumbs:

- [xcsh_dns_lb_health_check](../data-sources/dns_lb_health_check.md#canonical-3100122033112233-3012231223301113-3200002222023330-1201232003223131-2023031121303302-1301302210202321-2031302202212103-1331213200311221)
- [Property reference](data-sources--dns_lb_health_check--reference--group-001.md#canonical-3120123222202123-1012301220221110-0112203121332121-2102322302132322-2211003220321202-1031302101200212-1132031013203122-0302130033110030)
- icmp_health_check

<a id="canonical-1321231321300110-2233110121112311-3321213113333300-3201113022102101-1331022133200131-2001030322021222-0330100122322121-1123103132320103"></a>

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

<a id="canonical-1331020132111303-2200020003013033-1022223330231120-1032310021122220-1131203210033220-3200103003022112-0113233032230321-2232222103232122"></a>

## Direct properties — icmp_health_check / 212002112121 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0103300102131320-1211120322330202-3001100222331110-0033312122023103-3113031033000303-0122133101333201-1213323110001203-0312010332001123"></a>

## Next pages — icmp_health_check / 212002112121 / 4

- [Property reference](data-sources--dns_lb_health_check--reference--group-001.md#canonical-3120123222202123-1012301220221110-0112203121332121-2102322302132322-2211003220321202-1031302101200212-1132031013203122-0302130033110030)
- [xcsh_dns_lb_health_check](../data-sources/dns_lb_health_check.md#canonical-3100122033112233-3012231223301113-3200002222023330-1201232003223131-2023031121303302-1301302210202321-2031302202212103-1331213200311221)

<a id="canonical-3312322330200030-1232321021303111-3101103331113211-2113200131132022-1122310212123331-1332031232103330-0012301023331021-3103320231031323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2302033321201122-0103233023323112-2030301303322120-2122003110300011-1311101302310321-0233301303332310-1122201111110033-0321002000211023"></a>

## tcp_health_check — tcp_health_check / 333122000230 / 2

Breadcrumbs:

- [xcsh_dns_lb_health_check](../data-sources/dns_lb_health_check.md#canonical-3100122033112233-3012231223301113-3200002222023330-1201232003223131-2023031121303302-1301302210202321-2031302202212103-1331213200311221)
- [Property reference](data-sources--dns_lb_health_check--reference--group-001.md#canonical-3120123222202123-1012301220221110-0112203121332121-2102322302132322-2211003220321202-1031302101200212-1132031013203122-0302130033110030)
- tcp_health_check

<a id="canonical-0220021320200322-2101330233302312-3231330303221312-0310022032033012-2210203320220201-2330212033301022-0033112123000121-3201200003202203"></a>

Type: `"single"`. Computed.

Configuration parameter for tcp health check.

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

<a id="canonical-2213313213103210-0223002310300113-2102232231311331-0122333022221320-1021310330221010-0333120331013032-3233222130303123-0321002020102121"></a>

## Direct properties — tcp_health_check / 333122000230 / 3

<a id="canonical-1120310012121030-0301300012210331-1230101010322101-2220333123323133-2331023300111031-3130121312200321-2202110132010311-2330320120002130"></a>

<a id="canonical-3121003031132102-0030203220133200-0131320021002133-2333022112213011-0213313102320312-2023013222013230-1133313111132233-2202332101313121"></a>

## health_check_port property — tcp_health_check / 333122000230 / 4

Type: `"number"`. Computed.

Health Check Port. Port used for performing health check.

Upstream description:

Port used for performing health check.

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
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-0002102013231001-2312231032232011-3202121313321031-1023231302100103-2311332330213022-1122313032213310-3322232110121122-2100303130321103"></a>

<a id="canonical-3331012200111300-1110330110132111-1321130120003231-1122000001102311-1131031222303113-2300113303212000-2211321301102302-0313011323320101"></a>

## health_check_secondary_port property — tcp_health_check / 333122000230 / 5

Type: `"number"`. Computed.

Secondary port used for performing health check. If included, both ports must be healthy for the
health check to pass.

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
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-0320022033030000-0000101013303303-2131313311000002-0103021113122200-2101112223213010-1323312210120323-1302111132223313-1032220030311303"></a>

<a id="canonical-3333223211200233-1103001332233300-2200211211113301-1022303301320001-3003201201123320-3320030310310232-1022003020110313-0013311211300303"></a>

## receive property — tcp_health_check / 333122000230 / 6

Type: `"string"`. Computed.

Regular expression used to match against the response to the monitor's request. Mark node up upon
receipt of a successful regular expression match. Uses re2 regular expression syntax.

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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "2048",
    "ves.io.schema.rules.string.regex": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "2048",
    "ves.io.schema.rules.string.regex": "true"
  }
}
```

<a id="canonical-2202001033000311-0010321033320033-0002220301003113-2031230021012211-0201303230130233-1312102002233333-3330301203203130-3002222332201010"></a>

<a id="canonical-1003320031233323-1021122100121301-3101131122010133-2123222100220033-1102231322101011-2201100002032111-1002310312133010-2233201001331113"></a>

## send property — tcp_health_check / 333122000230 / 7

Type: `"string"`. Computed.

Send this string to target (default empty. When send and receive are both empty, monitor just tests
3WHS).

Upstream description:

Send this string to target (default empty. When send and receive are both empty, monitor just tests
3WHS)

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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "2048"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "2048"
  }
}
```

<a id="canonical-2201320021222311-0203121301121103-3002203221112202-1132002000300133-0100332223233103-1000130100312332-3320002003010113-2102233332222103"></a>

## Next pages — tcp_health_check / 333122000230 / 8

- [Property reference](data-sources--dns_lb_health_check--reference--group-001.md#canonical-3120123222202123-1012301220221110-0112203121332121-2102322302132322-2211003220321202-1031302101200212-1132031013203122-0302130033110030)
- [xcsh_dns_lb_health_check](../data-sources/dns_lb_health_check.md#canonical-3100122033112233-3012231223301113-3200002222023330-1201232003223131-2023031121303302-1301302210202321-2031302202212103-1331213200311221)

<a id="canonical-2303133011132012-0032300330121222-0130321113203111-3133103302322001-3331022011323230-2300302212313100-3303032312332312-1223301313120331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3013333201230320-2210111203200300-3233002222011132-2323030310303123-1210103032100011-0212121333332330-0123201311030130-2130233310112013"></a>

## tcp_hex_health_check — tcp_hex_health_check / 211321021301 / 2

Breadcrumbs:

- [xcsh_dns_lb_health_check](../data-sources/dns_lb_health_check.md#canonical-3100122033112233-3012231223301113-3200002222023330-1201232003223131-2023031121303302-1301302210202321-2031302202212103-1331213200311221)
- [Property reference](data-sources--dns_lb_health_check--reference--group-001.md#canonical-3120123222202123-1012301220221110-0112203121332121-2102322302132322-2211003220321202-1031302101200212-1132031013203122-0302130033110030)
- tcp_hex_health_check

<a id="canonical-0121210010121103-1223321111223031-1023300133012310-1133122223220133-3310102210222002-0112100230310130-0223203310223333-3221000113220121"></a>

Type: `"single"`. Computed.

Configuration parameter for tcp hex health check.

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

<a id="canonical-0313333000132333-0032231231012211-1103330110300030-0201011121102011-3010313112211131-0332002203222200-0222022221200221-1010332220012210"></a>

## Direct properties — tcp_hex_health_check / 211321021301 / 3

<a id="canonical-0201013213212312-1123023310223123-1013123231232232-2310010122101012-3130233013001123-3212210121321221-2031320333000322-0211031233133030"></a>

<a id="canonical-1011231310302302-2233133330312311-2302111331322331-0203031313011303-1213120233111030-1321233023301133-3013013112120013-0113231301131031"></a>

## health_check_port property — tcp_hex_health_check / 211321021301 / 4

Type: `"number"`. Computed.

Health Check Port. Port used for performing health check.

Upstream description:

Port used for performing health check.

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
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-3232212222030201-3202103211232210-1202123322200223-2022102300113000-0133003013312002-3022320113210033-2031031210023223-1301303302003021"></a>

<a id="canonical-2332231131111130-0220123233132123-0331123120310113-2301220321101200-2202210210002020-3300203013022000-1330101300102203-2313033013021322"></a>

## health_check_secondary_port property — tcp_hex_health_check / 211321021301 / 5

Type: `"number"`. Computed.

Secondary port used for performing health check. If included, both ports must be healthy for the
health check to pass.

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
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-0300022111112323-3312002021122313-2002213001021021-3310310033222210-2101000300212123-2112210230200013-2101322232101203-3311012100010101"></a>

<a id="canonical-2310130312122210-3220210122203210-3313311312222023-1003010120232203-0013202100333010-3320312301212320-1033132213103022-2033101011330000"></a>

## receive property — tcp_hex_health_check / 211321021301 / 6

Type: `"string"`. Computed.

Hex encoded raw bytes expected in the response.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 2048,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hex",
    "maxLength": 2048,
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
    "ves.io.schema.rules.string.hex": "true",
    "ves.io.schema.rules.string.max_len": "2048"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hex": "true",
    "ves.io.schema.rules.string.max_len": "2048"
  }
}
```

<a id="canonical-3002232302221113-3322013231101020-3333131133200230-0213330120312223-3322132203332222-0012300330310130-1112131112102123-1312201013231112"></a>

<a id="canonical-2032332201023121-1130201102010001-0210323330100012-2311033122113323-1110322322221122-0001213303103130-1130330132220211-1020303212033233"></a>

## send property — tcp_hex_health_check / 211321021301 / 7

Type: `"string"`. Computed.

Hex encoded raw bytes sent in the request. Empty payloads imply a connect-only health check.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 2048,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hex",
    "maxLength": 2048,
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
    "ves.io.schema.rules.string.hex": "true",
    "ves.io.schema.rules.string.max_len": "2048"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hex": "true",
    "ves.io.schema.rules.string.max_len": "2048"
  }
}
```

<a id="canonical-3203220030310013-3233201200131301-0233001230300313-2233002000300112-1303313120023232-0132132131020333-1212203121323312-2030111121102330"></a>

## Next pages — tcp_hex_health_check / 211321021301 / 8

- [Property reference](data-sources--dns_lb_health_check--reference--group-001.md#canonical-3120123222202123-1012301220221110-0112203121332121-2102322302132322-2211003220321202-1031302101200212-1132031013203122-0302130033110030)
- [xcsh_dns_lb_health_check](../data-sources/dns_lb_health_check.md#canonical-3100122033112233-3012231223301113-3200002222023330-1201232003223131-2023031121303302-1301302210202321-2031302202212103-1331213200311221)

<a id="canonical-1003110030133312-2333313002120220-2022222203021020-3311213231330220-2111301321213201-0000331121303100-0212210000033032-1003333133210312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2012011203302033-1112233103023101-3333210122032101-1332001000103000-2100312312332203-2002311133013231-3223133223131201-2102303022132112"></a>

## udp_health_check — udp_health_check / 100122001320 / 2

Breadcrumbs:

- [xcsh_dns_lb_health_check](../data-sources/dns_lb_health_check.md#canonical-3100122033112233-3012231223301113-3200002222023330-1201232003223131-2023031121303302-1301302210202321-2031302202212103-1331213200311221)
- [Property reference](data-sources--dns_lb_health_check--reference--group-001.md#canonical-3120123222202123-1012301220221110-0112203121332121-2102322302132322-2211003220321202-1031302101200212-1132031013203122-0302130033110030)
- udp_health_check

<a id="canonical-0332303000012330-0113320233211013-2113221201312020-0130010210030223-1120022123020203-0210002131333002-2313031011222300-1131130002110020"></a>

Type: `"single"`. Computed.

Configuration parameter for udp health check.

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

<a id="canonical-0310130031331331-3233122303303232-0010231331103132-1033012130313321-2110323132123112-0213030032223010-0210022012130030-2101331121221323"></a>

## Direct properties — udp_health_check / 100122001320 / 3

<a id="canonical-1001212233231002-1002313130130100-0200211222130210-1331000200130133-2120030221133310-2103120200310302-3300310300012332-3030033303121231"></a>

<a id="canonical-1133232101230111-0210120211211131-1021032330321031-3102220303111213-0020221202112123-3331122133130020-3030203011110013-2000131320020010"></a>

## health_check_port property — udp_health_check / 100122001320 / 4

Type: `"number"`. Computed.

Health Check Port. Port used for performing health check.

Upstream description:

Port used for performing health check.

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
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-3333110010321330-0231033230231200-0211023003030133-3112313111313313-3000000100122101-1110200302001023-0311233022022021-2221031202301313"></a>

<a id="canonical-3011120123120212-3133132032120012-0300010232312130-1333312201303033-1220212122010033-0201201311133001-2232221121113203-1003003203210111"></a>

## health_check_secondary_port property — udp_health_check / 100122001320 / 5

Type: `"number"`. Computed.

Secondary port used for performing health check. If included, both ports must be healthy for the
health check to pass.

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
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-0230231131333201-0033022311023232-1222211133030311-0031123112100211-3101133010212311-1110032003020223-3333031311203100-2202322100130133"></a>

<a id="canonical-3230201003213110-0311201102112111-0323221003103313-0312113311030010-1132011121120300-2002303201330323-1202002213232000-1230312003133130"></a>

## receive property — udp_health_check / 100122001320 / 6

Type: `"string"`. Computed.

UDP response to be matched. It can be a regular expression.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 2048,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 2048,
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
    "ves.io.schema.rules.string.max_len": "2048",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.regex": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "2048",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.regex": "true"
  }
}
```

<a id="canonical-3002123021130131-2220133311320122-0032021323221332-2213001032122031-2112033323320013-3031021300120031-1201302032121000-1021121221333331"></a>

<a id="canonical-0023321123212000-1112320122011232-2312022101303021-0311120233132220-2200302331223301-3000313021013210-1322332013103113-1332333000333032"></a>

## send property — udp_health_check / 100122001320 / 7

Type: `"string"`. Computed.

Send String. UDP payload.

Upstream description:

UDP payload.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 2048,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 2048,
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
    "ves.io.schema.rules.string.max_len": "2048",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "2048",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-1021321030030212-0133101021312010-0302331030231002-3012223313121120-0031213211113300-1013300023131030-3332032120201303-1213121300232303"></a>

## Next pages — udp_health_check / 100122001320 / 8

- [Property reference](data-sources--dns_lb_health_check--reference--group-001.md#canonical-3120123222202123-1012301220221110-0112203121332121-2102322302132322-2211003220321202-1031302101200212-1132031013203122-0302130033110030)
- [xcsh_dns_lb_health_check](../data-sources/dns_lb_health_check.md#canonical-3100122033112233-3012231223301113-3200002222023330-1201232003223131-2023031121303302-1301302210202321-2031302202212103-1331213200311221)
