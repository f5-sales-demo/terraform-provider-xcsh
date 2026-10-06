---
page_title: "xcsh_service_policy reference"
subcategory: "Security"
description: "Complete grouped canonical reference for xcsh_service_policy reference."
---

# xcsh_service_policy reference

<a id="canonical-3312102302111103-0221020013211300-2120323123300113-1101123102331230-1032322121222101-1332231300102132-0233300331323303-0031323122032013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Property reference

Breadcrumbs:

- [xcsh_service_policy](../data-sources/service_policy.md#canonical-1112022302333333-1130100232232301-3023212310221022-3103031003221200-2030020130230132-1201210321332200-2212311003030121-1321310310032303)
- Property reference

<a id="canonical-0030210210313010-1332032311311111-1302002232020010-3033012033321223-1010310030010310-2221302233322132-1221033133101210-3312203222113132"></a>

### Direct properties for `xcsh_service_policy`

- [allow_all_requests](data-sources--service_policy--reference--group-001.md#canonical-1310133010132113-2321211101112030-2032301102311223-3121100330132020-0312130331122112-2311320303111131-3110000220310101-2300003320213301): complete subsection reference.

- [allow_list](data-sources--service_policy--reference--group-001.md#canonical-3210232132133332-2232321312121110-0111111310110022-2030312322330300-2002320322232323-1133002023213102-2130223332331111-3332033011101311): complete subsection reference.

<a id="canonical-2021002001313310-2331021110333232-2222330130301303-1332102030102203-2101203131212113-1210131103122232-1200132131221020-0201033002220313"></a>

<a id="canonical-0200231112013012-2123102323021110-2133301322133110-0302321023303032-1010023311313311-2301131001233330-1101131220110113-2220110202312012"></a>

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

- [any_server](data-sources--service_policy--reference--group-001.md#canonical-3101033101321213-3011013020303022-2020233002120302-0120121332310033-2321321210322320-3100211331032333-1210331221301022-3010231111133032): complete subsection reference.

- [deny_all_requests](data-sources--service_policy--reference--group-001.md#canonical-1321120122221121-2211201320102032-3231122122100213-0112323220103312-2132331120321323-3313311212102223-3013200033323302-3032330332020102): complete subsection reference.

- [deny_list](data-sources--service_policy--reference--group-001.md#canonical-1133210330221300-2302313012231010-0300310022202220-2111123320332210-2101000320030120-3220020323311021-3132220133231132-0021330023303013): complete subsection reference.

<a id="canonical-0232110201311312-0230113113123310-1223322303321330-3003232220213331-3300013301003002-0003333231011012-2023320321101130-1222133000200112"></a>

<a id="canonical-2311223302301222-0333031130231200-0020300031132320-1010313003212031-1111320012103003-1010231010110023-0310032130320312-2002233232211322"></a>

#### `description` property

Type: `"string"`. Computed.

Description of the ServicePolicy.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-0121013310032101-0231302102220221-3223132121331113-3013020002132000-1023031330021300-1121122011200123-2113333321333123-3003012211332203"></a>

<a id="canonical-3121121031022212-2123002201123213-3101300233230002-3122112221333022-2110113303320201-0221333301211232-1003222012232031-0101011122212121"></a>

#### `id` property

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-0022232000222122-2033333001312202-0113233021331200-2010023013110333-0103012302300210-1021303203201002-3330000030220322-1110121110220301"></a>

<a id="canonical-2223210031311200-0333201001301320-3000032101122032-0313033031011223-1001311331132100-1222201103221210-3103002030220333-0232002320031013"></a>

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

<a id="canonical-2022032222200210-0022330332232003-3313001201113221-1110300302113303-3310201201001221-3012033230131210-2021320002112300-0123220313210231"></a>

<a id="canonical-0222312301301230-2030300300013302-0112331300030012-0303013113100230-0230332013031100-1320130101003111-1123103030121332-1212020031333210"></a>

#### `name` property

Type: `"string"`. Required.

Name of the ServicePolicy.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-3020023302113102-3302301232331100-0001100002030002-2233221110023212-3202030102233323-0033220331013203-3332102103132213-3220031313102112"></a>

<a id="canonical-0322303300301302-0113200332301332-1222233130113201-3203022200311002-0300223132023331-0212301333332201-0132032132002222-1333002031010230"></a>

#### `namespace` property

Type: `"string"`. Required.

Namespace where the ServicePolicy exists.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

- [rule_list](data-sources--service_policy--reference--group-001.md#canonical-0200021302120322-3330103111223310-0110033233312323-2211001021302330-2000313233321013-1330302203133111-3101023222321313-3302013113202033): complete subsection reference.

<a id="canonical-0321011233203030-0331130203211121-1031132232322121-0110122031210203-2023032012313300-0320030311200013-1001030211002303-2010110202110133"></a>

<a id="canonical-3312002002323311-1210110023321122-2320222021033022-1331201111001003-1022220302110101-3001010231000022-2001322120021001-0213033312000331"></a>

#### `server_name` property

Type: `"string"`. Computed.

Exclusive with \[any\_server server\_name\_matcher server\_selector\] The expected name of the
server to which the request API is directed. The actual names for the server are extracted from the
HTTP Host header and the name of the virtual\_host to which the request is directed. If the request
is directed to a virtual K8s service, the actual names also contain the name of that service. The
predicate evaluates to true if any of the actual names is the same as the expected server name.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 256
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-06T12:36:10+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256"
  }
}
```

- [server_name_matcher](data-sources--service_policy--reference--group-003.md#canonical-3021113112310011-2312023300100332-3211231333010200-0121033221312122-0113321130233033-3313001233322031-0000023202030312-0011301313100330): complete subsection reference.

- [server_selector](data-sources--service_policy--reference--group-003.md#canonical-3122313323222133-3022130222310211-0113113011303232-0203212121203123-0123123113030131-0111331230021010-3121011333120221-2330322213331033): complete subsection reference.

<a id="canonical-0231132133121201-3311110321100132-3002322012023301-2231221020321200-1210323310311112-0311200132110312-2231110300323322-1222110213332220"></a>

### All schema paths for `xcsh_service_policy`

Each exact path has one authoritative reference destination. Collection element indices are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `allow_all_requests` | [allow_all_requests](data-sources--service_policy--reference--group-001.md#canonical-3130031112221203-3121213003013120-2300321000003133-3013333202132322-3110122103130021-2102312123313230-2212103110313030-3021332211013333) |
| `allow_list` | [allow_list](data-sources--service_policy--reference--group-001.md#canonical-3310213202312001-1301321031302013-3201211132300133-2332211120332103-2310000311022100-2111212212131121-0103323001220022-0232310310100313) |
| `allow_list.asn_list` | [allow_list.asn_list](data-sources--service_policy--reference--group-001.md#canonical-1312111133013001-0010333311120122-2202201333312132-2110322133223101-0202130220103231-1122012112132031-2122112232221223-1130221213013113) |
| `allow_list.asn_list.as_numbers` | [allow_list.asn_list.as_numbers](data-sources--service_policy--reference--group-001.md#canonical-2132211032001022-3322120022132121-2223203200103133-3333123030310232-2330103300033022-2122210021302323-0110230302332032-3033123210210331) |
| `allow_list.asn_set` | [allow_list.asn_set](data-sources--service_policy--reference--group-001.md#canonical-1333322332223310-0113001220232202-1030100121101210-1310033023330202-0213121232200211-1311003320013113-1211220210131201-3010220113203230) |
| `allow_list.asn_set.name` | [allow_list.asn_set.name](data-sources--service_policy--reference--group-001.md#canonical-2221233200011120-0303202320311030-2010201213030012-3131030100213023-0101201103313300-1033323231231210-1233021333033231-3213110030013332) |
| `allow_list.asn_set.namespace` | [allow_list.asn_set.namespace](data-sources--service_policy--reference--group-001.md#canonical-0220223321300300-2111330112030220-0033233310031130-0333311011302202-0102112222320131-3201101201011310-0301002012111222-2231102212310203) |
| `allow_list.asn_set.tenant` | [allow_list.asn_set.tenant](data-sources--service_policy--reference--group-001.md#canonical-3101003101330333-2031030003200212-2300322231023111-2112133320223232-1123030223320230-1330322102132031-0120023220220123-2020200002200212) |
| `allow_list.country_list` | [allow_list.country_list](data-sources--service_policy--reference--group-001.md#canonical-3201223210332200-0210030310212133-1013320230232111-3133223331303232-0320233110212313-1101232230010231-3301233222201211-1310222323131220) |
| `allow_list.default_action_allow` | [allow_list.default_action_allow](data-sources--service_policy--reference--group-001.md#canonical-2322220103311102-1112100020121201-3113332022132223-1220112123203211-2313331003031201-0032310312003122-2330011333301020-0333210132201312) |
| `allow_list.default_action_deny` | [allow_list.default_action_deny](data-sources--service_policy--reference--group-001.md#canonical-1303000123333311-1031110003101232-1221003200013133-2201200210111220-3030330232131002-2213321312330133-3122013323131313-2020111311131332) |
| `allow_list.default_action_next_policy` | [allow_list.default_action_next_policy](data-sources--service_policy--reference--group-001.md#canonical-0012323302131020-3003331031113212-2321012000012330-3321333330122233-3130132221312232-2322011113002330-0131120312133222-2211032312030002) |
| `allow_list.ip_prefix_set` | [allow_list.ip_prefix_set](data-sources--service_policy--reference--group-001.md#canonical-0320313123101223-1031332321320123-2203001322002023-1122002013132033-1213300131001210-2003303030321321-1320101123222000-0102100130012130) |
| `allow_list.ip_prefix_set.name` | [allow_list.ip_prefix_set.name](data-sources--service_policy--reference--group-001.md#canonical-3313001222300132-0112013233111320-3123031221030121-0022331031223121-3033123032302333-1331222302310233-0002232221013032-0220132213323321) |
| `allow_list.ip_prefix_set.namespace` | [allow_list.ip_prefix_set.namespace](data-sources--service_policy--reference--group-001.md#canonical-3323002203313220-2021332210200302-1310110331330112-2322203033010112-2033232330322213-0301233211322321-3203203010023103-1313212020321223) |
| `allow_list.ip_prefix_set.tenant` | [allow_list.ip_prefix_set.tenant](data-sources--service_policy--reference--group-001.md#canonical-0102323230313202-3330101233033303-0102021200202000-1023013233232320-3130010303332323-1200000320323223-3223203013103213-2120123321311311) |
| `allow_list.prefix_list` | [allow_list.prefix_list](data-sources--service_policy--reference--group-001.md#canonical-2011232121313112-2310112313330123-0331312012133013-2112232220323120-0102312002131322-3132131003330101-3222231022030223-0221012130233322) |
| `allow_list.prefix_list.prefixes` | [allow_list.prefix_list.prefixes](data-sources--service_policy--reference--group-001.md#canonical-0233220002202113-0012322213010201-2300000130020120-0123220121012031-1322300320222220-2301121202201031-3312122230120123-2031332203122221) |
| `allow_list.tls_fingerprint_classes` | [allow_list.tls_fingerprint_classes](data-sources--service_policy--reference--group-001.md#canonical-2200032110111002-1232302122323330-0102211101331001-2111120100003320-0222132102301333-2320033332131303-0111121102231032-1111131000112002) |
| `allow_list.tls_fingerprint_values` | [allow_list.tls_fingerprint_values](data-sources--service_policy--reference--group-001.md#canonical-2103232203012222-3010303310201122-2213221301222313-1300301201221030-2030213031211021-3123201120130323-0002110030232031-2113123103033301) |
| `annotations` | [annotations](data-sources--service_policy--reference--group-001.md#canonical-2021002001313310-2331021110333232-2222330130301303-1332102030102203-2101203131212113-1210131103122232-1200132131221020-0201033002220313) |
| `any_server` | [any_server](data-sources--service_policy--reference--group-001.md#canonical-2213123300100210-0321113302220212-0023213232002331-3312020330323132-3010230320130302-3300323232110113-0333123031020303-0132100313012303) |
| `deny_all_requests` | [deny_all_requests](data-sources--service_policy--reference--group-001.md#canonical-1233003033211320-0320102302012323-0122002032101001-0303221201213302-3202101031101002-1123032012131000-3030133100001030-3230132122233001) |
| `deny_list` | [deny_list](data-sources--service_policy--reference--group-001.md#canonical-1022223021201211-2230013123132233-1220120331103012-2232310222203111-3331022203003133-1331133111223111-3330021021323023-2312130331321003) |
| `deny_list.asn_list` | [deny_list.asn_list](data-sources--service_policy--reference--group-001.md#canonical-3332222101231202-1210322010202011-0013010022332202-2202132202003130-1121231101200103-1113013103100003-3220323330201113-2223221100012100) |
| `deny_list.asn_list.as_numbers` | [deny_list.asn_list.as_numbers](data-sources--service_policy--reference--group-001.md#canonical-2101133123131110-1302322012322202-2012210132012210-2300001321102332-0032120222022221-3113212101021210-3313330223232121-3102232211332200) |
| `deny_list.asn_set` | [deny_list.asn_set](data-sources--service_policy--reference--group-001.md#canonical-1223201123120023-1121020133310021-3022130023232101-1232230023010223-2223201020231013-2323332121000001-0032310302313233-3323010302220232) |
| `deny_list.asn_set.name` | [deny_list.asn_set.name](data-sources--service_policy--reference--group-001.md#canonical-3133013100213233-1123003200103120-0132030122013220-0231202323013311-2322330313313103-2120100231012310-0112312102212311-2201222103200013) |
| `deny_list.asn_set.namespace` | [deny_list.asn_set.namespace](data-sources--service_policy--reference--group-001.md#canonical-0221330232123000-3213121022302233-3010022230301220-2123231222031102-0201300132013111-1310110322111231-0312201303222011-3100133111233313) |
| `deny_list.asn_set.tenant` | [deny_list.asn_set.tenant](data-sources--service_policy--reference--group-001.md#canonical-3113232323331121-3002210020002123-2203021210132330-1010102323131133-0320033013222223-3033030203133010-1322010302233233-2332311220021011) |
| `deny_list.country_list` | [deny_list.country_list](data-sources--service_policy--reference--group-001.md#canonical-2311232000332011-2201122232020320-1013331123021023-3310120330200230-1022330200320213-0303123130122130-3221302201321333-1202203312110320) |
| `deny_list.default_action_allow` | [deny_list.default_action_allow](data-sources--service_policy--reference--group-001.md#canonical-0022323322233322-2300133310312232-0030310001120013-3212132010331201-1223030111202311-3133201312001312-1113212221202113-0300221331131001) |
| `deny_list.default_action_deny` | [deny_list.default_action_deny](data-sources--service_policy--reference--group-001.md#canonical-1301120211113332-2311313101103202-1001112231022311-2103330230130230-2210313210133221-0232132212300120-3101021232210033-0233331212311231) |
| `deny_list.default_action_next_policy` | [deny_list.default_action_next_policy](data-sources--service_policy--reference--group-001.md#canonical-2321132010020011-2331132200310002-2231133322202210-2012232120223120-1310221231120300-2321110000333202-3032023320312033-3210120233302010) |
| `deny_list.ip_prefix_set` | [deny_list.ip_prefix_set](data-sources--service_policy--reference--group-001.md#canonical-3102223301301201-2113120232113013-0333021111133010-0210331220212212-2312012220323302-2321230312302303-1001130321111230-2301320331213212) |
| `deny_list.ip_prefix_set.name` | [deny_list.ip_prefix_set.name](data-sources--service_policy--reference--group-001.md#canonical-2210122210300313-2002032131301223-3002133333020310-0322022312111321-3233200013031303-0202131212312020-0213203200110131-1210222212003001) |
| `deny_list.ip_prefix_set.namespace` | [deny_list.ip_prefix_set.namespace](data-sources--service_policy--reference--group-001.md#canonical-3020322220330301-1012122310200221-3001133123110323-1013022100002322-0013000123130021-1101001320302102-1323033102200201-3331302231202311) |
| `deny_list.ip_prefix_set.tenant` | [deny_list.ip_prefix_set.tenant](data-sources--service_policy--reference--group-001.md#canonical-2002313233122232-0221101110210101-0321301020311333-0031332013032230-1031301011201121-3232313331100132-2321000330101113-1331010213131333) |
| `deny_list.prefix_list` | [deny_list.prefix_list](data-sources--service_policy--reference--group-001.md#canonical-0110002022022111-2211102313122002-1132000302031101-2223133030301111-3002101300032113-3010033000121030-1331110001232021-2213100200200012) |
| `deny_list.prefix_list.prefixes` | [deny_list.prefix_list.prefixes](data-sources--service_policy--reference--group-001.md#canonical-3013121103313210-0202220302200012-3031023202212211-0010011131201322-2101113030100213-3000131331113010-3033210223202333-3112030011001321) |
| `deny_list.tls_fingerprint_classes` | [deny_list.tls_fingerprint_classes](data-sources--service_policy--reference--group-001.md#canonical-2011001300111011-3222122130222330-2110323331320203-0200310020130120-2031011133012120-0030102213221301-2230010111321111-0330112130013301) |
| `deny_list.tls_fingerprint_values` | [deny_list.tls_fingerprint_values](data-sources--service_policy--reference--group-001.md#canonical-2302323023100211-1322210121030130-2013000310003202-3013222010301123-1220331331121130-2120332030312110-1111323102030231-2021012013301120) |
| `description` | [description](data-sources--service_policy--reference--group-001.md#canonical-0232110201311312-0230113113123310-1223322303321330-3003232220213331-3300013301003002-0003333231011012-2023320321101130-1222133000200112) |
| `id` | [ID](data-sources--service_policy--reference--group-001.md#canonical-0121013310032101-0231302102220221-3223132121331113-3013020002132000-1023031330021300-1121122011200123-2113333321333123-3003012211332203) |
| `labels` | [labels](data-sources--service_policy--reference--group-001.md#canonical-0022232000222122-2033333001312202-0113233021331200-2010023013110333-0103012302300210-1021303203201002-3330000030220322-1110121110220301) |
| `name` | [name](data-sources--service_policy--reference--group-001.md#canonical-2022032222200210-0022330332232003-3313001201113221-1110300302113303-3310201201001221-3012033230131210-2021320002112300-0123220313210231) |
| `namespace` | [namespace](data-sources--service_policy--reference--group-001.md#canonical-3020023302113102-3302301232331100-0001100002030002-2233221110023212-3202030102233323-0033220331013203-3332102103132213-3220031313102112) |
| `rule_list` | [rule_list](data-sources--service_policy--reference--group-001.md#canonical-1323033222302302-3111221111132320-3301101131323321-3012033101003230-1210201102233222-0102310331101011-1003003201123022-2202030210300232) |
| `rule_list.rules` | [rule_list.rules](data-sources--service_policy--reference--group-001.md#canonical-2101123033112122-2302210303331323-0113210121020123-3130021131003022-3210120301201231-1310022213100132-3100013031031120-0123232230222033) |
| `rule_list.rules.metadata` | [rule_list.rules.metadata](data-sources--service_policy--reference--group-001.md#canonical-2132122221011103-1233111033030100-3231103313200310-0200202333111121-1011321030232023-3132303212002110-0111033113000111-2013212221213320) |
| `rule_list.rules.metadata.description_spec` | [rule_list.rules.metadata.description_spec](data-sources--service_policy--reference--group-001.md#canonical-0212221232311013-3022222221012223-1023030033300300-3020230111333121-1312100333310020-3100233212123211-3121332300122110-1203203222301111) |
| `rule_list.rules.metadata.name` | [rule_list.rules.metadata.name](data-sources--service_policy--reference--group-001.md#canonical-1000132132131211-3110100001311320-2302200310113010-2112212130123330-3230221122002300-2131322303012032-0103232131233232-0233231121323210) |
| `rule_list.rules.spec` | [rule_list.rules.spec](data-sources--service_policy--reference--group-001.md#canonical-1133013120211033-0320113201200113-2030303010023202-2220311210313200-0022221110321003-0302002312313013-0323310011323211-0101213132323023) |
| `rule_list.rules.spec.action` | [rule_list.rules.spec.action](data-sources--service_policy--reference--group-001.md#canonical-0030310321031303-1210300001000323-0330321121202332-0313300002301030-0102113303323012-1221120000322010-0021322223222103-1201002013030122) |
| `rule_list.rules.spec.any_asn` | [rule_list.rules.spec.any_asn](data-sources--service_policy--reference--group-001.md#canonical-1020211012322112-2130223103313230-1221102230302302-3121003131023133-2231320201222110-1300320231013220-3313322103033120-2132132220211101) |
| `rule_list.rules.spec.any_client` | [rule_list.rules.spec.any_client](data-sources--service_policy--reference--group-001.md#canonical-1300330003023333-1333111332220312-2310222203120032-1300110232122131-3203121320022231-3332023213231211-0220220220232233-3210030133023220) |
| `rule_list.rules.spec.any_ip` | [rule_list.rules.spec.any_ip](data-sources--service_policy--reference--group-001.md#canonical-1003211211011222-3123201310100102-0030313120303132-1101032013110102-1011030231231001-0020010301132211-2101301031032131-2211303133020200) |
| `rule_list.rules.spec.api_group_matcher` | [rule_list.rules.spec.api_group_matcher](data-sources--service_policy--reference--group-001.md#canonical-2111211320123130-0231032322312103-3203100222330213-3320320203102003-0313131031333320-3100030220111230-3213103121231300-3330103201102121) |
| `rule_list.rules.spec.api_group_matcher.invert_matcher` | [rule_list.rules.spec.api_group_matcher.invert_matcher](data-sources--service_policy--reference--group-001.md#canonical-3112233332331211-0022302002321222-3203322302302200-2022221123130113-2200320121023302-0223332100303022-3321201300000000-3033331132301132) |
| `rule_list.rules.spec.api_group_matcher.match` | [rule_list.rules.spec.api_group_matcher.match](data-sources--service_policy--reference--group-001.md#canonical-2332001011133303-3010112113111113-0102001020312120-1000113333310200-2112213020130102-1200330031310033-3121120312111100-2021011002103010) |
| `rule_list.rules.spec.arg_matchers` | [rule_list.rules.spec.arg_matchers](data-sources--service_policy--reference--group-001.md#canonical-0231002111103112-1231131210320221-1131032201121113-0310202302313003-2203202001203113-0110112201103112-1210001320120210-2212021302223002) |
| `rule_list.rules.spec.arg_matchers.check_not_present` | [rule_list.rules.spec.arg_matchers.check_not_present](data-sources--service_policy--reference--group-001.md#canonical-3121300003200003-2222002023101200-3000213023322300-1303323332113130-3232031132033301-0013131223012303-0121101101210311-2320311110330232) |
| `rule_list.rules.spec.arg_matchers.check_present` | [rule_list.rules.spec.arg_matchers.check_present](data-sources--service_policy--reference--group-001.md#canonical-0202121332112331-1303132123020311-0302232021111320-2332233023302320-3101031213202022-1021232001032313-2002132122321201-1300212201021323) |
| `rule_list.rules.spec.arg_matchers.invert_matcher` | [rule_list.rules.spec.arg_matchers.invert_matcher](data-sources--service_policy--reference--group-001.md#canonical-0303220323213301-3123011333233123-2030210313111003-1213122331302020-3303010332200110-0000211303203020-1211102200013313-0223220211233121) |
| `rule_list.rules.spec.arg_matchers.item` | [rule_list.rules.spec.arg_matchers.item](data-sources--service_policy--reference--group-001.md#canonical-2332023130032022-1132302310330132-2131001312321022-1113103313112333-3022313223011030-0221021122223210-2112321210013001-0313120133231001) |
| `rule_list.rules.spec.arg_matchers.item.exact_values` | [rule_list.rules.spec.arg_matchers.item.exact_values](data-sources--service_policy--reference--group-001.md#canonical-3213320220301231-2113020331222313-2010113133122330-2000320110223221-3013200303220123-0101100201110201-1233020101313330-3101013100121303) |
| `rule_list.rules.spec.arg_matchers.item.regex_values` | [rule_list.rules.spec.arg_matchers.item.regex_values](data-sources--service_policy--reference--group-001.md#canonical-1301323232211312-3312120320132331-2223023212110302-0300033232021011-2122312312132332-1120111323023123-0100121220311013-1321122113130100) |
| `rule_list.rules.spec.arg_matchers.item.transformers` | [rule_list.rules.spec.arg_matchers.item.transformers](data-sources--service_policy--reference--group-001.md#canonical-2303131320011000-0000330310010011-3121000200001321-0120033112301030-2111000021302120-2220000002100000-0333201011022013-0001133330130001) |
| `rule_list.rules.spec.arg_matchers.name` | [rule_list.rules.spec.arg_matchers.name](data-sources--service_policy--reference--group-001.md#canonical-2322232213331000-2122332021220233-1220132231233002-0212320233033130-0031023103030010-0233202121223223-0111133213013111-2301133322102121) |
| `rule_list.rules.spec.asn_list` | [rule_list.rules.spec.asn_list](data-sources--service_policy--reference--group-001.md#canonical-1030233020210002-0201201310003213-2003323131101020-2023121200110312-0202012201302331-2013100012231021-0132003213302302-0031123023213131) |
| `rule_list.rules.spec.asn_list.as_numbers` | [rule_list.rules.spec.asn_list.as_numbers](data-sources--service_policy--reference--group-001.md#canonical-0230111022020120-2312230000130022-2222323112130003-3202012213102102-3130023312132111-3230231131211301-0231212213230320-0310101302133001) |
| `rule_list.rules.spec.asn_matcher` | [rule_list.rules.spec.asn_matcher](data-sources--service_policy--reference--group-001.md#canonical-3310301311023211-3212333103231332-1301223331002130-1323130023300133-1200010303133320-1231002331210313-3003231230221310-1103000103033223) |
| `rule_list.rules.spec.asn_matcher.asn_sets` | [rule_list.rules.spec.asn_matcher.asn_sets](data-sources--service_policy--reference--group-001.md#canonical-1200202203313103-0013212021220122-2311212212023102-1101120113310332-3303303331023210-0033221122022102-2103312010010023-3020300102222003) |
| `rule_list.rules.spec.asn_matcher.asn_sets.kind` | [rule_list.rules.spec.asn_matcher.asn_sets.kind](data-sources--service_policy--reference--group-001.md#canonical-3213033302123332-1030130013021222-1100202011200203-3232232312322231-3332120020003022-2111203320311011-0213130321310023-1331211301220211) |
| `rule_list.rules.spec.asn_matcher.asn_sets.name` | [rule_list.rules.spec.asn_matcher.asn_sets.name](data-sources--service_policy--reference--group-001.md#canonical-0022113003211033-1230333333111332-0103301202123231-1212133320020313-0323010110222021-1120102322021200-3123001011132311-0233120231022310) |
| `rule_list.rules.spec.asn_matcher.asn_sets.namespace` | [rule_list.rules.spec.asn_matcher.asn_sets.namespace](data-sources--service_policy--reference--group-001.md#canonical-0012112203232311-0022022213112112-0310331232100202-2312202201030312-3312213301020310-0030313001200330-3022131320102210-3032030103212003) |
| `rule_list.rules.spec.asn_matcher.asn_sets.tenant` | [rule_list.rules.spec.asn_matcher.asn_sets.tenant](data-sources--service_policy--reference--group-001.md#canonical-2230200321222023-0122123120032212-3210101013113231-0001200330101313-2030201022302232-2321112011300233-1211230210212321-1203113212322220) |
| `rule_list.rules.spec.asn_matcher.asn_sets.uid` | [rule_list.rules.spec.asn_matcher.asn_sets.uid](data-sources--service_policy--reference--group-001.md#canonical-1033130001323331-3022202233230230-1211200300222233-3321232100201133-2303321102230031-3232032111123102-2102222320112332-1330332112133133) |
| `rule_list.rules.spec.body_matcher` | [rule_list.rules.spec.body_matcher](data-sources--service_policy--reference--group-001.md#canonical-1100103221313312-1312213330003231-2032111201030011-3133232310132122-0013222333002231-3223303322013223-0230113211022321-1120220132031332) |
| `rule_list.rules.spec.body_matcher.exact_values` | [rule_list.rules.spec.body_matcher.exact_values](data-sources--service_policy--reference--group-001.md#canonical-0022310003233020-1313003133110303-1010122030223103-0130120033122203-0212313310003123-1102202301223313-3203213130223100-0021131013200300) |
| `rule_list.rules.spec.body_matcher.regex_values` | [rule_list.rules.spec.body_matcher.regex_values](data-sources--service_policy--reference--group-001.md#canonical-0020121230121202-3110113013122330-0233120202231010-1213321133313201-2302100030201011-2201123312003000-2031111110300123-2111131013013021) |
| `rule_list.rules.spec.body_matcher.transformers` | [rule_list.rules.spec.body_matcher.transformers](data-sources--service_policy--reference--group-001.md#canonical-1311120131232022-0012013333132311-3312020230302200-2330022303322300-3310022211101002-3130222112032300-1202032303113200-1331221103232111) |
| `rule_list.rules.spec.bot_action` | [rule_list.rules.spec.bot_action](data-sources--service_policy--reference--group-002.md#canonical-0133132310332033-1121101121102222-0003231011211121-1121212002010233-2100112013133303-1010033012022130-0202303032132322-1233303100300200) |
| `rule_list.rules.spec.bot_action.bot_skip_processing` | [rule_list.rules.spec.bot_action.bot_skip_processing](data-sources--service_policy--reference--group-002.md#canonical-2130033111300203-1221231230001123-3000103031321132-3101102230212103-0201220011111323-0213130321010030-0231030100131201-1000112123012113) |
| `rule_list.rules.spec.bot_action.none` | [rule_list.rules.spec.bot_action.none](data-sources--service_policy--reference--group-002.md#canonical-2312322002032001-1100123303302213-1212233230320012-2303113313033131-0000233201010130-0133000200102313-0200022313231202-0323012002200221) |
| `rule_list.rules.spec.client_name` | [rule_list.rules.spec.client_name](data-sources--service_policy--reference--group-001.md#canonical-0122122233031311-1200303031303233-2001113120113211-2021121031101002-0003132111113300-0011211221021110-0310231003032211-3202122120001313) |
| `rule_list.rules.spec.client_name_matcher` | [rule_list.rules.spec.client_name_matcher](data-sources--service_policy--reference--group-002.md#canonical-3230203233231333-0311300302111332-3232031330201101-3202221313210010-3002012001210022-1010332200312300-2111010311003001-0211022333113300) |
| `rule_list.rules.spec.client_name_matcher.exact_values` | [rule_list.rules.spec.client_name_matcher.exact_values](data-sources--service_policy--reference--group-002.md#canonical-0110221212023020-2002330213032203-1122033020223320-0222133233310211-1302200233230200-3012131310101303-0323332133032303-0210002203010120) |
| `rule_list.rules.spec.client_name_matcher.regex_values` | [rule_list.rules.spec.client_name_matcher.regex_values](data-sources--service_policy--reference--group-002.md#canonical-3333103100022013-0203233120032003-0111203000021333-3033032013120030-3201130033121101-3121210303011031-1030300222130213-0032330102223102) |
| `rule_list.rules.spec.client_name_matcher.transformers` | [rule_list.rules.spec.client_name_matcher.transformers](data-sources--service_policy--reference--group-002.md#canonical-0100232010003320-1023310102010202-0133323030020233-2013031202310302-3201031130110312-1310332130200103-2212030330200122-0303001130201331) |
| `rule_list.rules.spec.client_selector` | [rule_list.rules.spec.client_selector](data-sources--service_policy--reference--group-002.md#canonical-0021003033023030-1032122012333232-3012330332331020-3331110221001300-3301031012220111-0101233002221021-3122102211210330-3220021210122323) |
| `rule_list.rules.spec.client_selector.expressions` | [rule_list.rules.spec.client_selector.expressions](data-sources--service_policy--reference--group-002.md#canonical-1122131031331101-1132203301003312-1302221230232301-1310200030032231-2230122013120202-0313323033120221-0303123120111012-0131110103212103) |
| `rule_list.rules.spec.cookie_matchers` | [rule_list.rules.spec.cookie_matchers](data-sources--service_policy--reference--group-002.md#canonical-3022303210322011-1331331100001332-1300031010023032-0312201111112020-0230322303020113-0122212030021002-3222211120321133-1021000200001303) |
| `rule_list.rules.spec.cookie_matchers.check_not_present` | [rule_list.rules.spec.cookie_matchers.check_not_present](data-sources--service_policy--reference--group-002.md#canonical-2001300121223031-0020010012022030-3130113233001003-1230221303013313-0333133000323012-3332133023212011-2111011103222023-0313001212330233) |
| `rule_list.rules.spec.cookie_matchers.check_present` | [rule_list.rules.spec.cookie_matchers.check_present](data-sources--service_policy--reference--group-002.md#canonical-2213201131210103-0002320003313033-1022202333132200-0322112311013111-0022021211210301-3232213123101102-3010001301001012-3011130310220011) |
| `rule_list.rules.spec.cookie_matchers.invert_matcher` | [rule_list.rules.spec.cookie_matchers.invert_matcher](data-sources--service_policy--reference--group-002.md#canonical-1202323210300133-0102023012133310-0031212001003222-3233201003001222-2032232001302230-2321102213120033-0021021123101301-3031020301200113) |
| `rule_list.rules.spec.cookie_matchers.item` | [rule_list.rules.spec.cookie_matchers.item](data-sources--service_policy--reference--group-002.md#canonical-0210020002333210-2312000302103031-3121223003210113-2112323112000000-2022110130132001-2332133033313001-2300021121312033-3213220313031033) |
| `rule_list.rules.spec.cookie_matchers.item.exact_values` | [rule_list.rules.spec.cookie_matchers.item.exact_values](data-sources--service_policy--reference--group-002.md#canonical-1303200000020232-3031323201322203-1013212113130000-1031313303232131-2323333311023002-0333011310322133-3110300332330213-1312132022030011) |
| `rule_list.rules.spec.cookie_matchers.item.regex_values` | [rule_list.rules.spec.cookie_matchers.item.regex_values](data-sources--service_policy--reference--group-002.md#canonical-1032032103031000-0030202201111313-3222200213121032-3323103030023212-1123213110311033-0120232012131213-3120031001103201-2033333123033022) |
| `rule_list.rules.spec.cookie_matchers.item.transformers` | [rule_list.rules.spec.cookie_matchers.item.transformers](data-sources--service_policy--reference--group-002.md#canonical-1313200201122022-3031222210230132-1031302211212112-1013320120222001-2101231332103003-0300231011132200-1213030200003311-0303232223221132) |
| `rule_list.rules.spec.cookie_matchers.name` | [rule_list.rules.spec.cookie_matchers.name](data-sources--service_policy--reference--group-002.md#canonical-2032322003122032-2331110210312111-0221031112221332-2010101031300110-2110122330130200-2001102313133300-1000211121210001-2303122201320102) |
| `rule_list.rules.spec.domain_matcher` | [rule_list.rules.spec.domain_matcher](data-sources--service_policy--reference--group-002.md#canonical-3121212333330100-1103320130023311-3111311030320303-1001113330321331-2221021021212000-2300312331100012-1212202110302300-0023223223031331) |
| `rule_list.rules.spec.domain_matcher.exact_values` | [rule_list.rules.spec.domain_matcher.exact_values](data-sources--service_policy--reference--group-002.md#canonical-0023101030131212-3002101010110131-0132020223110202-3210132033203131-2001032203110222-3203311031222231-1202312022113010-2100121331131331) |
| `rule_list.rules.spec.domain_matcher.regex_values` | [rule_list.rules.spec.domain_matcher.regex_values](data-sources--service_policy--reference--group-002.md#canonical-1031132202230330-3321001313001231-1210311301130200-1303022132032102-1322222300103330-2300113210232122-0330202311231131-3033200212020030) |
| `rule_list.rules.spec.domain_matcher.transformers` | [rule_list.rules.spec.domain_matcher.transformers](data-sources--service_policy--reference--group-002.md#canonical-1333211021320022-2003033321121121-1301012022313223-1210023023001022-3301120312130221-2021122032200103-3213102031103112-3202000000022013) |
| `rule_list.rules.spec.expiration_timestamp` | [rule_list.rules.spec.expiration_timestamp](data-sources--service_policy--reference--group-001.md#canonical-2221011121222100-1023211233032112-1321023312012332-0320001222332330-3020230011132223-1202111202123103-3121210033103020-3031010220121000) |
| `rule_list.rules.spec.headers` | [rule_list.rules.spec.headers](data-sources--service_policy--reference--group-002.md#canonical-1001323031223130-3200130200303132-0011101103022322-3221232223200332-2231010031002222-2303021010213331-0222313303011331-3002333222231021) |
| `rule_list.rules.spec.headers.check_not_present` | [rule_list.rules.spec.headers.check_not_present](data-sources--service_policy--reference--group-002.md#canonical-0113220031331322-3121130212333302-2002012121001223-1233211111021323-3012133323312321-3323111301030203-1233332121120031-1301001123313102) |
| `rule_list.rules.spec.headers.check_present` | [rule_list.rules.spec.headers.check_present](data-sources--service_policy--reference--group-002.md#canonical-1323023110013302-3123031303031301-0231221112121332-2000120330213130-3021011110313020-0320012032132231-0010120020330220-1331113111201131) |
| `rule_list.rules.spec.headers.invert_matcher` | [rule_list.rules.spec.headers.invert_matcher](data-sources--service_policy--reference--group-002.md#canonical-2200020300203210-2102123211333203-1000003020231003-0032032011302302-0020100330011102-0103132101213121-0033032033222102-0323132120000132) |
| `rule_list.rules.spec.headers.item` | [rule_list.rules.spec.headers.item](data-sources--service_policy--reference--group-002.md#canonical-0131300331331012-3203011121023013-3230331023200313-0323032320020111-1023002001332211-2210321021023010-3212032223222121-2010132201032032) |
| `rule_list.rules.spec.headers.item.exact_values` | [rule_list.rules.spec.headers.item.exact_values](data-sources--service_policy--reference--group-002.md#canonical-1201123121330233-0320331203133120-0221032002020023-1131330133213111-3101311031300320-1321112001122010-3112021012223303-0100000010302330) |
| `rule_list.rules.spec.headers.item.regex_values` | [rule_list.rules.spec.headers.item.regex_values](data-sources--service_policy--reference--group-002.md#canonical-3212322023122313-3203321001100221-1103013223133210-0312123311031011-2011121001220312-1110330302022100-3230003122332311-3330011112011120) |
| `rule_list.rules.spec.headers.item.transformers` | [rule_list.rules.spec.headers.item.transformers](data-sources--service_policy--reference--group-002.md#canonical-3222213231201121-0022131033222302-2110320132000233-3222020123002023-1313022332112212-0323230221020202-2012001211001001-2012221130200021) |
| `rule_list.rules.spec.headers.name` | [rule_list.rules.spec.headers.name](data-sources--service_policy--reference--group-002.md#canonical-0031122131232323-0311201220123003-3112232130032000-1001113030223200-0222110102200130-3230320200223321-1012032030013303-0020033020131231) |
| `rule_list.rules.spec.http_method` | [rule_list.rules.spec.http_method](data-sources--service_policy--reference--group-002.md#canonical-2313321331200130-0221313021311130-2202301020320121-3033333012111120-3100213131030221-3011122303123000-1123002030313001-3011130013233322) |
| `rule_list.rules.spec.http_method.invert_matcher` | [rule_list.rules.spec.http_method.invert_matcher](data-sources--service_policy--reference--group-002.md#canonical-3002231303110022-0312113102132001-1321121302002030-2132132300112132-2013301103223020-1020300333332330-2310000321321100-3003211030320101) |
| `rule_list.rules.spec.http_method.methods` | [rule_list.rules.spec.http_method.methods](data-sources--service_policy--reference--group-002.md#canonical-1112232112123232-1033332002021200-0313002233110322-1123010013230230-3203010331222302-1020001133220012-3013102202320232-1012312320233133) |
| `rule_list.rules.spec.ip_matcher` | [rule_list.rules.spec.ip_matcher](data-sources--service_policy--reference--group-002.md#canonical-1331103032030102-2202002232213222-1121220301103201-1312101033120123-2222122002112213-3011010203232302-2120300023201203-1221230211311001) |
| `rule_list.rules.spec.ip_matcher.invert_matcher` | [rule_list.rules.spec.ip_matcher.invert_matcher](data-sources--service_policy--reference--group-002.md#canonical-3310210102200031-3113013002201213-2030322331233002-2033200311013101-0000103132112123-0002301231302331-3331202331311030-0301223131302313) |
| `rule_list.rules.spec.ip_matcher.prefix_sets` | [rule_list.rules.spec.ip_matcher.prefix_sets](data-sources--service_policy--reference--group-002.md#canonical-0330011200233103-1103220323320011-1311331221001120-0211020322101330-3032112322011203-2003120203331101-3020030130023113-3230221113002313) |
| `rule_list.rules.spec.ip_matcher.prefix_sets.kind` | [rule_list.rules.spec.ip_matcher.prefix_sets.kind](data-sources--service_policy--reference--group-002.md#canonical-2322013233111133-1011230113131320-2111220110130101-1002020122033333-3122113211323310-2203021121123311-0233021020313210-3313103223230132) |
| `rule_list.rules.spec.ip_matcher.prefix_sets.name` | [rule_list.rules.spec.ip_matcher.prefix_sets.name](data-sources--service_policy--reference--group-002.md#canonical-0330301031313300-3211323220200013-0331221123012112-3001311010313330-0001010300100130-1322232313103100-3331321000223200-1023221322132102) |
| `rule_list.rules.spec.ip_matcher.prefix_sets.namespace` | [rule_list.rules.spec.ip_matcher.prefix_sets.namespace](data-sources--service_policy--reference--group-002.md#canonical-3001031222221102-2221000130030010-3320022311220221-2212010013103330-1103103331101111-1322020012303333-1230210213020203-2131210213330011) |
| `rule_list.rules.spec.ip_matcher.prefix_sets.tenant` | [rule_list.rules.spec.ip_matcher.prefix_sets.tenant](data-sources--service_policy--reference--group-002.md#canonical-2121201132222230-1333213120120023-1233233012012311-3013220323021210-1310301301333321-2312012331003210-2123022211332323-0123210312032101) |
| `rule_list.rules.spec.ip_matcher.prefix_sets.uid` | [rule_list.rules.spec.ip_matcher.prefix_sets.uid](data-sources--service_policy--reference--group-002.md#canonical-0122311311233032-1022010012320301-2033020203222131-3122321101330203-3301300201003220-0210103122311300-2331133000333332-1233220301211333) |
| `rule_list.rules.spec.ip_prefix_list` | [rule_list.rules.spec.ip_prefix_list](data-sources--service_policy--reference--group-002.md#canonical-0332320110133201-3103032311222213-2312021222023232-0330203300133120-1232221103000032-0021300000103213-2220330322332103-1101231122101103) |
| `rule_list.rules.spec.ip_prefix_list.invert_match` | [rule_list.rules.spec.ip_prefix_list.invert_match](data-sources--service_policy--reference--group-002.md#canonical-3020222320221332-3211131100011233-0233021301301223-3120032112132122-2322120322302233-2032302202303103-2332021132210330-0103123213331131) |
| `rule_list.rules.spec.ip_prefix_list.ip_prefixes` | [rule_list.rules.spec.ip_prefix_list.ip_prefixes](data-sources--service_policy--reference--group-002.md#canonical-0333212333332103-3030131233230322-3332112111013110-0013032331201233-2132232001000230-2332021221012113-0313130323113332-0122133333332230) |
| `rule_list.rules.spec.ip_threat_category_list` | [rule_list.rules.spec.ip_threat_category_list](data-sources--service_policy--reference--group-002.md#canonical-0332320111030232-3310113121200022-1031330203210102-2212203332210312-3201012132323000-0200201001103023-0210310132323111-3010331010001100) |
| `rule_list.rules.spec.ip_threat_category_list.ip_threat_categories` | [rule_list.rules.spec.ip_threat_category_list.ip_threat_categories](data-sources--service_policy--reference--group-002.md#canonical-2102010000011313-1231113311132130-1000300113122132-2132213113111123-1310321100110133-1220232223110023-1310101201212000-0201222211103130) |
| `rule_list.rules.spec.ja4_tls_fingerprint` | [rule_list.rules.spec.ja4_tls_fingerprint](data-sources--service_policy--reference--group-002.md#canonical-2133321213201332-0103033131311023-0133012001231300-2210133121211321-1011003323303322-2313311011311000-3310323131231202-0121121103231010) |
| `rule_list.rules.spec.ja4_tls_fingerprint.exact_values` | [rule_list.rules.spec.ja4_tls_fingerprint.exact_values](data-sources--service_policy--reference--group-002.md#canonical-0101100031300100-1122131102312303-0302101013011302-3133300312100221-0330321102320221-0321120232331131-2220323311331113-0113313030320320) |
| `rule_list.rules.spec.jwt_claims` | [rule_list.rules.spec.jwt_claims](data-sources--service_policy--reference--group-002.md#canonical-1333311322310330-3102031002331332-3211032023103233-1223113012113130-3112020132003021-0221131203221330-1123130331202202-3210211021330122) |
| `rule_list.rules.spec.jwt_claims.check_not_present` | [rule_list.rules.spec.jwt_claims.check_not_present](data-sources--service_policy--reference--group-002.md#canonical-1132010000203301-0012121211121100-2112003032022101-3022332231302313-2130000000002101-2231001202321032-0313330101211332-0302231121023303) |
| `rule_list.rules.spec.jwt_claims.check_present` | [rule_list.rules.spec.jwt_claims.check_present](data-sources--service_policy--reference--group-002.md#canonical-1021011111003020-0111321101203321-1322100003311000-2230201220001321-1220231301200001-3321332130230331-3032222232331321-2223302012223311) |
| `rule_list.rules.spec.jwt_claims.invert_matcher` | [rule_list.rules.spec.jwt_claims.invert_matcher](data-sources--service_policy--reference--group-002.md#canonical-1110220012201301-1232232002032110-2011231102301332-0332002321321300-1313230010031102-3212031322203320-1121222302320213-1031033001030322) |
| `rule_list.rules.spec.jwt_claims.item` | [rule_list.rules.spec.jwt_claims.item](data-sources--service_policy--reference--group-002.md#canonical-3303300231310012-2210221200011322-3033332203300223-1033233132303122-1132123103000012-1300123112233213-0301111020322310-2223031312010032) |
| `rule_list.rules.spec.jwt_claims.item.exact_values` | [rule_list.rules.spec.jwt_claims.item.exact_values](data-sources--service_policy--reference--group-002.md#canonical-2131000112113003-3102230203212011-2313111102011211-0320331130011022-1011023030313311-2112120022300022-1102233112231130-2321301031220020) |
| `rule_list.rules.spec.jwt_claims.item.regex_values` | [rule_list.rules.spec.jwt_claims.item.regex_values](data-sources--service_policy--reference--group-002.md#canonical-3331310001031232-0232122121213310-0321013012102312-1122023212311313-0102332100121203-0133130200333312-3332212022020330-0022002031023010) |
| `rule_list.rules.spec.jwt_claims.item.transformers` | [rule_list.rules.spec.jwt_claims.item.transformers](data-sources--service_policy--reference--group-002.md#canonical-0313111331103320-2121332301210233-0223223210320232-0200220232033030-2231330133232122-3103033212010201-2021313300111303-2113203321030200) |
| `rule_list.rules.spec.jwt_claims.name` | [rule_list.rules.spec.jwt_claims.name](data-sources--service_policy--reference--group-002.md#canonical-2133232001211003-0333000330222003-2203013111010320-2233123120202301-3032211102121320-3132232001102102-1102313112110000-2223003010031333) |
| `rule_list.rules.spec.label_matcher` | [rule_list.rules.spec.label_matcher](data-sources--service_policy--reference--group-002.md#canonical-0320223011033230-1210020210111321-2201022210233031-3323031132200230-0303021300212020-2023330312000311-0033033201332310-3310210302123032) |
| `rule_list.rules.spec.label_matcher.keys` | [rule_list.rules.spec.label_matcher.keys](data-sources--service_policy--reference--group-002.md#canonical-1030102212313222-3121203023100211-2231011320110222-2320320110030322-0020232223211230-1231312100213111-1111222000300322-2111331013331012) |
| `rule_list.rules.spec.log_rule_evaluation` | [rule_list.rules.spec.log_rule_evaluation](data-sources--service_policy--reference--group-001.md#canonical-3100030233213202-2131211321003130-0313102000123103-3310210200223132-3313233322332131-2201131301113002-2300200102330321-1132220122121223) |
| `rule_list.rules.spec.mum_action` | [rule_list.rules.spec.mum_action](data-sources--service_policy--reference--group-002.md#canonical-3011112132131102-2033202212222122-3103323112320223-3231220213112223-2002132103222211-2031301130001220-3031203330010021-1003020332333233) |
| `rule_list.rules.spec.mum_action.default` | [rule_list.rules.spec.mum_action.default](data-sources--service_policy--reference--group-002.md#canonical-1120210102212202-0010103020100022-2221301130030031-0111213201001111-0200203303032211-3203122111120132-2210130310110222-0300313032333022) |
| `rule_list.rules.spec.mum_action.skip_processing` | [rule_list.rules.spec.mum_action.skip_processing](data-sources--service_policy--reference--group-002.md#canonical-0330211221313200-0320010013212020-3300331232303321-2310123123312332-2231303100302130-2231120122313032-1232102102120200-1022213221322320) |
| `rule_list.rules.spec.path` | [rule_list.rules.spec.path](data-sources--service_policy--reference--group-002.md#canonical-3112202323333330-1002011132230113-0301131232010302-3223310322220113-2211000203231201-3331203031030111-2211031100031002-2232301030003023) |
| `rule_list.rules.spec.path.encoded_path_matcher` | [rule_list.rules.spec.path.encoded_path_matcher](data-sources--service_policy--reference--group-002.md#canonical-1330100122300331-1010313230320222-2032001201021311-1210133203332112-0000133202122221-2023110121133022-0230312023000201-0210233201223311) |
| `rule_list.rules.spec.path.exact_values` | [rule_list.rules.spec.path.exact_values](data-sources--service_policy--reference--group-002.md#canonical-0120312020113032-1030120030023321-3210021211312110-1120320012301200-2203130312012220-3011232321202021-0110101100132130-2003310023000023) |
| `rule_list.rules.spec.path.invert_matcher` | [rule_list.rules.spec.path.invert_matcher](data-sources--service_policy--reference--group-002.md#canonical-0001322103131130-1110233012103113-1113020300221322-2001213320020330-0201322323121321-0010112211101020-0020022103310202-0233020120132330) |
| `rule_list.rules.spec.path.prefix_values` | [rule_list.rules.spec.path.prefix_values](data-sources--service_policy--reference--group-002.md#canonical-1010312003323223-3201132023023210-3212301221100123-1230300002332211-2301312202110203-1223312121230331-3230222233300220-2301311203112011) |
| `rule_list.rules.spec.path.regex_values` | [rule_list.rules.spec.path.regex_values](data-sources--service_policy--reference--group-002.md#canonical-1320103323133211-0212121220121210-1212100113200310-3113320110302231-0112032211022330-3033331003200220-0120113131332010-2331101223110123) |
| `rule_list.rules.spec.path.suffix_values` | [rule_list.rules.spec.path.suffix_values](data-sources--service_policy--reference--group-002.md#canonical-1011132201120030-2303313012003131-0203001003333022-3103122030222022-3120002133012021-0021000302033131-0302020230322221-0302113102010323) |
| `rule_list.rules.spec.path.transformers` | [rule_list.rules.spec.path.transformers](data-sources--service_policy--reference--group-002.md#canonical-1223332033210322-1002000230132301-1321120211330120-1120110302101211-2001020312003303-0002122010121332-0232222213030002-2012103003111103) |
| `rule_list.rules.spec.port_matcher` | [rule_list.rules.spec.port_matcher](data-sources--service_policy--reference--group-002.md#canonical-0102020310300131-3033033100311323-3102032333033332-0203020321021112-0113020213001120-3112302213320120-3120120130023133-0323220002322302) |
| `rule_list.rules.spec.port_matcher.invert_matcher` | [rule_list.rules.spec.port_matcher.invert_matcher](data-sources--service_policy--reference--group-002.md#canonical-1122010102022223-1313322021133300-3201123231331023-0222100120310123-3233012330020101-3123212231013000-1032300211133103-2202030233330312) |
| `rule_list.rules.spec.port_matcher.ports` | [rule_list.rules.spec.port_matcher.ports](data-sources--service_policy--reference--group-002.md#canonical-0132232331202033-0111310020202112-1310331330133210-0223232110220211-2032012001123101-1210332232122010-3212113213223233-2022211232120231) |
| `rule_list.rules.spec.query_params` | [rule_list.rules.spec.query_params](data-sources--service_policy--reference--group-002.md#canonical-0021032210113322-0333013200203201-0130322031100030-0211101301022012-1220031203102000-0311031333001100-1031333021331331-3030232201313010) |
| `rule_list.rules.spec.query_params.check_not_present` | [rule_list.rules.spec.query_params.check_not_present](data-sources--service_policy--reference--group-002.md#canonical-3031301021203101-3023103333000213-1030313120311310-0102322303232101-1032000211322023-2031321302210032-0232310321133300-1300222013113020) |
| `rule_list.rules.spec.query_params.check_present` | [rule_list.rules.spec.query_params.check_present](data-sources--service_policy--reference--group-002.md#canonical-1320133322002031-0220121202223031-1010313232312202-3021131301000330-3203310133301313-0003232022010121-2120002031030321-0022110001103210) |
| `rule_list.rules.spec.query_params.invert_matcher` | [rule_list.rules.spec.query_params.invert_matcher](data-sources--service_policy--reference--group-002.md#canonical-0220111311323221-1001132313020012-1033021211020310-0300203230123122-0131220111200032-2212102210012111-3210312220230120-3023022012110233) |
| `rule_list.rules.spec.query_params.item` | [rule_list.rules.spec.query_params.item](data-sources--service_policy--reference--group-002.md#canonical-3032123310110312-1100332321020022-2031230101313200-3122001101120303-0222221202231203-3312312201103021-3010330132100200-0331022113103311) |
| `rule_list.rules.spec.query_params.item.exact_values` | [rule_list.rules.spec.query_params.item.exact_values](data-sources--service_policy--reference--group-002.md#canonical-3031110320310201-3213010203020133-1310031020002210-0122202023110333-2211020032131211-0033133300022200-2302013123303130-3332300303002220) |
| `rule_list.rules.spec.query_params.item.regex_values` | [rule_list.rules.spec.query_params.item.regex_values](data-sources--service_policy--reference--group-002.md#canonical-2033303300020230-2200230100301121-0321021232131000-2213031322111010-3000112101020022-0021303000110010-0032001311313020-0301012201202311) |
| `rule_list.rules.spec.query_params.item.transformers` | [rule_list.rules.spec.query_params.item.transformers](data-sources--service_policy--reference--group-002.md#canonical-2012231103001332-3203110113113021-2312122100101111-3101210123200311-2010223300311002-0213030222130002-1122211202301003-1233033133111212) |
| `rule_list.rules.spec.query_params.key` | [rule_list.rules.spec.query_params.key](data-sources--service_policy--reference--group-002.md#canonical-2103011311212312-2032012112102200-0311323221101311-0010312312220323-0300032233313000-3122132111021110-1322111100110320-0323113310320331) |
| `rule_list.rules.spec.request_constraints` | [rule_list.rules.spec.request_constraints](data-sources--service_policy--reference--group-002.md#canonical-2112123331021132-0331130211222303-1220320200120130-2123332130032032-1232330002003032-1003011220032133-0200332210010013-3303233302321003) |
| `rule_list.rules.spec.request_constraints.max_cookie_count_exceeds` | [rule_list.rules.spec.request_constraints.max_cookie_count_exceeds](data-sources--service_policy--reference--group-002.md#canonical-1131210031223213-1132313020110003-1011223120333302-2100330202123121-3223011130121221-3201303130330123-0203131231131233-1331023213220011) |
| `rule_list.rules.spec.request_constraints.max_cookie_count_none` | [rule_list.rules.spec.request_constraints.max_cookie_count_none](data-sources--service_policy--reference--group-002.md#canonical-1111313123200300-3212220133232203-1300023320321201-0122130333002003-3001130311302312-1000301222001313-1323233131120130-1203222310311223) |
| `rule_list.rules.spec.request_constraints.max_cookie_key_size_exceeds` | [rule_list.rules.spec.request_constraints.max_cookie_key_size_exceeds](data-sources--service_policy--reference--group-002.md#canonical-3031223001320213-2021311110120300-3032323002000221-0031033113031331-2101111233203233-0012311311100212-0111322132031123-0220330111110331) |
| `rule_list.rules.spec.request_constraints.max_cookie_key_size_none` | [rule_list.rules.spec.request_constraints.max_cookie_key_size_none](data-sources--service_policy--reference--group-002.md#canonical-2021221011301022-3011212022323322-3001223110300212-0311230002132131-3310122203230113-3332032132012031-2332021012123322-3201221303223213) |
| `rule_list.rules.spec.request_constraints.max_cookie_value_size_exceeds` | [rule_list.rules.spec.request_constraints.max_cookie_value_size_exceeds](data-sources--service_policy--reference--group-002.md#canonical-2121130112321031-1313331103233102-2020121322012102-0301033212030200-3301123312331201-3131123103231230-0001330132231032-3120133300102203) |
| `rule_list.rules.spec.request_constraints.max_cookie_value_size_none` | [rule_list.rules.spec.request_constraints.max_cookie_value_size_none](data-sources--service_policy--reference--group-002.md#canonical-0121330310211302-2003222211210310-0212230230301210-3330023103312300-1220132113112132-3313133311130222-0333022021331003-1113123123000333) |
| `rule_list.rules.spec.request_constraints.max_header_count_exceeds` | [rule_list.rules.spec.request_constraints.max_header_count_exceeds](data-sources--service_policy--reference--group-002.md#canonical-2230132211021102-3303102212231301-2122030321302023-3021323033330122-0131231202013311-2302202323232223-2021201232131232-1132223133210203) |
| `rule_list.rules.spec.request_constraints.max_header_count_none` | [rule_list.rules.spec.request_constraints.max_header_count_none](data-sources--service_policy--reference--group-002.md#canonical-1311313201013113-1212123033301013-0201322220033001-2011103102121023-3232133201022011-1303310131012321-1200120011010321-2033301222023130) |
| `rule_list.rules.spec.request_constraints.max_header_key_size_exceeds` | [rule_list.rules.spec.request_constraints.max_header_key_size_exceeds](data-sources--service_policy--reference--group-002.md#canonical-3302231322211230-0022112201303001-3103031030301013-2032031303130020-1331313020302123-1321222133121203-0013010232032212-1020013221012012) |
| `rule_list.rules.spec.request_constraints.max_header_key_size_none` | [rule_list.rules.spec.request_constraints.max_header_key_size_none](data-sources--service_policy--reference--group-002.md#canonical-2221002200001123-3212303123011102-1222130123011013-2010123222030023-2330230120132310-1102132303311221-0333210211333013-2220232220220200) |
| `rule_list.rules.spec.request_constraints.max_header_value_size_exceeds` | [rule_list.rules.spec.request_constraints.max_header_value_size_exceeds](data-sources--service_policy--reference--group-002.md#canonical-3331202323031333-1320202023223301-3001120321133101-3100313222110002-2233021033000331-0101331233102332-2001120112020010-1002022013112131) |
| `rule_list.rules.spec.request_constraints.max_header_value_size_none` | [rule_list.rules.spec.request_constraints.max_header_value_size_none](data-sources--service_policy--reference--group-002.md#canonical-2033001132200123-2321030230232231-2212130213013221-2121321112321123-3012121310323302-3023133310122302-3100010120101123-0003230301103103) |
| `rule_list.rules.spec.request_constraints.max_parameter_count_exceeds` | [rule_list.rules.spec.request_constraints.max_parameter_count_exceeds](data-sources--service_policy--reference--group-002.md#canonical-0101110133003132-2032222120310332-3002311033110331-1333330113322011-1300032121032031-0101322011033000-3320233021023122-1201022030031223) |
| `rule_list.rules.spec.request_constraints.max_parameter_count_none` | [rule_list.rules.spec.request_constraints.max_parameter_count_none](data-sources--service_policy--reference--group-002.md#canonical-0313322233111221-1210021213200123-3231221212033211-3312310330100321-0201032030100102-0022133123230211-3212130210111222-2333100110021230) |
| `rule_list.rules.spec.request_constraints.max_parameter_name_size_exceeds` | [rule_list.rules.spec.request_constraints.max_parameter_name_size_exceeds](data-sources--service_policy--reference--group-002.md#canonical-2032232301003310-0000322220332033-1020311210122010-1001023223331112-3121120321302001-2022200311210120-3133313102013032-3123033213120001) |
| `rule_list.rules.spec.request_constraints.max_parameter_name_size_none` | [rule_list.rules.spec.request_constraints.max_parameter_name_size_none](data-sources--service_policy--reference--group-002.md#canonical-3101312003323232-0032113110210331-2321113220033213-0023130003312201-2223212030323220-1301323121000213-3320321233033012-3212333220323302) |
| `rule_list.rules.spec.request_constraints.max_parameter_value_size_exceeds` | [rule_list.rules.spec.request_constraints.max_parameter_value_size_exceeds](data-sources--service_policy--reference--group-002.md#canonical-2033122330120011-2122021120213110-2100110310323022-2231120103003212-1310000210331302-2223101111230132-0021100223110221-3303113300121132) |
| `rule_list.rules.spec.request_constraints.max_parameter_value_size_none` | [rule_list.rules.spec.request_constraints.max_parameter_value_size_none](data-sources--service_policy--reference--group-002.md#canonical-0002210320222020-1300110232313333-3012300310132032-1011303323333311-0100331302223311-2003333101032303-2301210011100320-1231223103102123) |
| `rule_list.rules.spec.request_constraints.max_query_size_exceeds` | [rule_list.rules.spec.request_constraints.max_query_size_exceeds](data-sources--service_policy--reference--group-002.md#canonical-0032010330202320-3002001301210231-3323121111223121-3033200010213011-2231113121002333-0210211102212200-2022212130022111-3312221233131001) |
| `rule_list.rules.spec.request_constraints.max_query_size_none` | [rule_list.rules.spec.request_constraints.max_query_size_none](data-sources--service_policy--reference--group-002.md#canonical-2302321312102200-0233313322120033-2310132332032032-1001111220102121-3030020302032323-3321112312323322-2331303333100020-2003311223202201) |
| `rule_list.rules.spec.request_constraints.max_request_line_size_exceeds` | [rule_list.rules.spec.request_constraints.max_request_line_size_exceeds](data-sources--service_policy--reference--group-002.md#canonical-1233202013023031-0202311120202333-3012313030031120-2102101110133111-3210123211101232-0231032013100200-2032313323221111-0012303302203311) |
| `rule_list.rules.spec.request_constraints.max_request_line_size_none` | [rule_list.rules.spec.request_constraints.max_request_line_size_none](data-sources--service_policy--reference--group-002.md#canonical-3131101223332303-3021032331023101-2232323000321112-0230333123331211-0111103111303200-1323112132023102-0111110012123303-2120320320121113) |
| `rule_list.rules.spec.request_constraints.max_request_size_exceeds` | [rule_list.rules.spec.request_constraints.max_request_size_exceeds](data-sources--service_policy--reference--group-002.md#canonical-0011232231110030-3323303223203001-2113221232332112-3121131310333022-0303130033201111-0021300320323302-3322322332212231-2220102220200113) |
| `rule_list.rules.spec.request_constraints.max_request_size_none` | [rule_list.rules.spec.request_constraints.max_request_size_none](data-sources--service_policy--reference--group-002.md#canonical-2032300221313213-2323021323221310-2230003003033312-2101131033103331-0210020313003103-3320100033301110-1102301113012302-2211132101112103) |
| `rule_list.rules.spec.request_constraints.max_url_size_exceeds` | [rule_list.rules.spec.request_constraints.max_url_size_exceeds](data-sources--service_policy--reference--group-002.md#canonical-3102321322200020-0201221200020012-2332200321023110-3021313221230132-2200132210221103-3103222112233222-0212001222221231-3120102233223013) |
| `rule_list.rules.spec.request_constraints.max_url_size_none` | [rule_list.rules.spec.request_constraints.max_url_size_none](data-sources--service_policy--reference--group-002.md#canonical-1313023023203300-0032003021013100-1302123103200301-2121223010312110-2013102002212233-1002000013001322-0121231202230231-0213321321331003) |
| `rule_list.rules.spec.segment_policy` | [rule_list.rules.spec.segment_policy](data-sources--service_policy--reference--group-002.md#canonical-3323232232032330-1122211331303130-1023302210003331-0013122122013120-1002031300033001-2331202230203020-1031113221232130-1002212101031213) |
| `rule_list.rules.spec.segment_policy.dst_any` | [rule_list.rules.spec.segment_policy.dst_any](data-sources--service_policy--reference--group-002.md#canonical-1321312333231211-2012322212003330-1002310223333030-1202100132013202-2030132101123031-1130131223203111-1023332120210333-2122311100310133) |
| `rule_list.rules.spec.segment_policy.dst_segments` | [rule_list.rules.spec.segment_policy.dst_segments](data-sources--service_policy--reference--group-002.md#canonical-0031213203022332-1022112230002233-2203321002120022-1133112331111202-3300311331010213-0022102322032123-1023020002302211-3001021021023002) |
| `rule_list.rules.spec.segment_policy.dst_segments.segments` | [rule_list.rules.spec.segment_policy.dst_segments.segments](data-sources--service_policy--reference--group-002.md#canonical-1002310132133310-0201232300331013-0301123313122333-1131113103202010-3220122322323201-2201223110003021-2303112112230112-1303201201010000) |
| `rule_list.rules.spec.segment_policy.dst_segments.segments.name` | [rule_list.rules.spec.segment_policy.dst_segments.segments.name](data-sources--service_policy--reference--group-002.md#canonical-1221110230301310-0131201212321122-0113003003312103-0232010031031121-1110302312211233-2030202321211333-0203333032201300-0223232330003130) |
| `rule_list.rules.spec.segment_policy.dst_segments.segments.namespace` | [rule_list.rules.spec.segment_policy.dst_segments.segments.namespace](data-sources--service_policy--reference--group-002.md#canonical-2120202321223210-2122112212030221-2221021330201121-1013023030301202-0300121210330333-0100311130313211-3013202223322221-1012320300222201) |
| `rule_list.rules.spec.segment_policy.dst_segments.segments.tenant` | [rule_list.rules.spec.segment_policy.dst_segments.segments.tenant](data-sources--service_policy--reference--group-002.md#canonical-3011301032103032-1233300121303113-0031122021013133-2220333002203010-2233131313023311-1001333033100330-3032030220010002-3233032100103301) |
| `rule_list.rules.spec.segment_policy.intra_segment` | [rule_list.rules.spec.segment_policy.intra_segment](data-sources--service_policy--reference--group-002.md#canonical-1230110130222112-0033112022211020-0200023213002330-1223122322211222-2123221023313332-3013022001202203-0001312021201121-0033213231022131) |
| `rule_list.rules.spec.segment_policy.src_any` | [rule_list.rules.spec.segment_policy.src_any](data-sources--service_policy--reference--group-002.md#canonical-1101103031230102-0230020121302021-2001333033001031-3023002312220232-1302000131223222-3313132130030110-0023031003332310-3132002333101203) |
| `rule_list.rules.spec.segment_policy.src_segments` | [rule_list.rules.spec.segment_policy.src_segments](data-sources--service_policy--reference--group-002.md#canonical-2301133103201303-2001023333321101-0132010032100323-1030203123002021-1100021200021233-1320101333332010-2230331010212112-2210232020122310) |
| `rule_list.rules.spec.segment_policy.src_segments.segments` | [rule_list.rules.spec.segment_policy.src_segments.segments](data-sources--service_policy--reference--group-002.md#canonical-3010233120011013-0220003003232100-1220130101303213-0232203100213013-1122133000103112-1033220213020101-0210122121030133-1030030213021023) |
| `rule_list.rules.spec.segment_policy.src_segments.segments.name` | [rule_list.rules.spec.segment_policy.src_segments.segments.name](data-sources--service_policy--reference--group-002.md#canonical-2320133321111300-3211202300330000-3123020210222102-0112120102130231-2330211230222122-1100212120320010-2221233321323330-1101320131012010) |
| `rule_list.rules.spec.segment_policy.src_segments.segments.namespace` | [rule_list.rules.spec.segment_policy.src_segments.segments.namespace](data-sources--service_policy--reference--group-002.md#canonical-1230301122130022-0203020120330023-2330000013103220-1310100333320033-2213100010322130-2002032220100312-0211030220313133-2322310232202113) |
| `rule_list.rules.spec.segment_policy.src_segments.segments.tenant` | [rule_list.rules.spec.segment_policy.src_segments.segments.tenant](data-sources--service_policy--reference--group-002.md#canonical-1321133222213232-3313200100232002-3231022123330231-0130302331000310-2133101112123201-1030323223131133-1331313020103033-2312303012113023) |
| `rule_list.rules.spec.tls_fingerprint_matcher` | [rule_list.rules.spec.tls_fingerprint_matcher](data-sources--service_policy--reference--group-002.md#canonical-3311100032302102-2323202000133230-2023032001333021-1201210301202211-1223001232030332-3002102203101103-0312130010013122-1133130200001323) |
| `rule_list.rules.spec.tls_fingerprint_matcher.classes` | [rule_list.rules.spec.tls_fingerprint_matcher.classes](data-sources--service_policy--reference--group-002.md#canonical-0333121223330310-3322313032201230-0013120222103300-3003203120231123-1332230003322001-0112022203101232-0013231101021231-2321333333332220) |
| `rule_list.rules.spec.tls_fingerprint_matcher.exact_values` | [rule_list.rules.spec.tls_fingerprint_matcher.exact_values](data-sources--service_policy--reference--group-002.md#canonical-2011231012212202-2021321320303201-0312332010222313-2222120231002131-2330133220203113-0021122300313021-3012121121312031-0223303013010030) |
| `rule_list.rules.spec.tls_fingerprint_matcher.excluded_values` | [rule_list.rules.spec.tls_fingerprint_matcher.excluded_values](data-sources--service_policy--reference--group-002.md#canonical-1200131312313331-0021301110001231-2020033301023203-2020100202230221-0300203113201130-1333021011032201-1312231123332333-0000113031133221) |
| `rule_list.rules.spec.user_identity_matcher` | [rule_list.rules.spec.user_identity_matcher](data-sources--service_policy--reference--group-002.md#canonical-0100212130233102-1310321312010333-2030333021002020-0223013210212033-3201132300013202-3220332031312301-2333203202012002-0211003212223012) |
| `rule_list.rules.spec.user_identity_matcher.exact_values` | [rule_list.rules.spec.user_identity_matcher.exact_values](data-sources--service_policy--reference--group-002.md#canonical-2233132333222131-0231231101030331-1223112100320122-1323202021302012-0122002132001213-2302112323131103-0232310133211212-3132022030202221) |
| `rule_list.rules.spec.user_identity_matcher.regex_values` | [rule_list.rules.spec.user_identity_matcher.regex_values](data-sources--service_policy--reference--group-002.md#canonical-1131312023312021-1230101210102111-2311331301310200-3201232002222003-2202210120300210-2231031302101120-3221231123211010-3233111032100223) |
| `rule_list.rules.spec.waf_action` | [rule_list.rules.spec.waf_action](data-sources--service_policy--reference--group-002.md#canonical-2201213013321111-0021023302222322-3012030010120000-0313121021213130-0213311111011011-3132001123022010-0013200303120001-0010312221111010) |
| `rule_list.rules.spec.waf_action.app_firewall_detection_control` | [rule_list.rules.spec.waf_action.app_firewall_detection_control](data-sources--service_policy--reference--group-002.md#canonical-0020213223313321-2120133310320131-2011210000132312-3113301020021323-0302030302133130-0203011310312310-0112222313103012-1001301021303002) |
| `rule_list.rules.spec.waf_action.app_firewall_detection_control.exclude_attack_type_contexts` | [rule_list.rules.spec.waf_action.app_firewall_detection_control.exclude_attack_type_contexts](data-sources--service_policy--reference--group-002.md#canonical-0313223222013300-2223010213233203-3000112212111221-1300011232121130-1012212112013302-3001330322220022-1201310312333133-3330230221001133) |
| `rule_list.rules.spec.waf_action.app_firewall_detection_control.exclude_attack_type_contexts.context` | [rule_list.rules.spec.waf_action.app_firewall_detection_control.exclude_attack_type_contexts.context](data-sources--service_policy--reference--group-002.md#canonical-3121103133132030-0130203123101231-3223300333120333-0011100233023020-2101231323232111-1112123213130202-3020010302003300-0321200300301321) |
| `rule_list.rules.spec.waf_action.app_firewall_detection_control.exclude_attack_type_contexts.context_name` | [rule_list.rules.spec.waf_action.app_firewall_detection_control.exclude_attack_type_contexts.context_name](data-sources--service_policy--reference--group-002.md#canonical-3110231203333213-2100132103102023-2002232113203223-3011311300330011-2220312103212200-1031213011003321-2002323130010303-1322133302203303) |
| `rule_list.rules.spec.waf_action.app_firewall_detection_control.exclude_attack_type_contexts.exclude_attack_type` | [rule_list.rules.spec.waf_action.app_firewall_detection_control.exclude_attack_type_contexts.exclude_attack_type](data-sources--service_policy--reference--group-002.md#canonical-2012101323113130-2322301023023333-1001022231231003-1131310120203112-2123010102302122-1332001302033030-2332303331031232-0120332211102301) |
| `rule_list.rules.spec.waf_action.app_firewall_detection_control.exclude_bot_name_contexts` | [rule_list.rules.spec.waf_action.app_firewall_detection_control.exclude_bot_name_contexts](data-sources--service_policy--reference--group-002.md#canonical-0100022320130212-0121002013023332-2121022130012331-1013120021320110-1103233103001031-2020212020100033-3013123002120033-3213112223212333) |
| `rule_list.rules.spec.waf_action.app_firewall_detection_control.exclude_bot_name_contexts.bot_name` | [rule_list.rules.spec.waf_action.app_firewall_detection_control.exclude_bot_name_contexts.bot_name](data-sources--service_policy--reference--group-002.md#canonical-1223031212301030-3302120030211111-2311012131320230-3220013222002312-1013232001233222-1110102202001032-1012112103321022-3322203213131131) |
| `rule_list.rules.spec.waf_action.app_firewall_detection_control.exclude_signature_contexts` | [rule_list.rules.spec.waf_action.app_firewall_detection_control.exclude_signature_contexts](data-sources--service_policy--reference--group-002.md#canonical-3232303210031312-3033010222203121-0320100002303131-2030300312300110-2123132031302323-3331330031132312-1222222311001223-1320310302201320) |
| `rule_list.rules.spec.waf_action.app_firewall_detection_control.exclude_signature_contexts.context` | [rule_list.rules.spec.waf_action.app_firewall_detection_control.exclude_signature_contexts.context](data-sources--service_policy--reference--group-002.md#canonical-1332010312113312-1222033200113132-3130310022333322-2022301122032021-3102112122101313-3310200131020332-1012201313030131-1202232330323011) |
| `rule_list.rules.spec.waf_action.app_firewall_detection_control.exclude_signature_contexts.context_name` | [rule_list.rules.spec.waf_action.app_firewall_detection_control.exclude_signature_contexts.context_name](data-sources--service_policy--reference--group-002.md#canonical-0020111011133001-3233222011223222-1131010313311201-3230000233201331-1123103231210002-1232303332322002-2301321003132221-3032102300223220) |
| `rule_list.rules.spec.waf_action.app_firewall_detection_control.exclude_signature_contexts.signature_id` | [rule_list.rules.spec.waf_action.app_firewall_detection_control.exclude_signature_contexts.signature_id](data-sources--service_policy--reference--group-003.md#canonical-3133010203031111-3030010100300220-1213331210330100-3232011213113231-1303230323033130-1022103132032131-0113032213230231-1212102000021003) |
| `rule_list.rules.spec.waf_action.app_firewall_detection_control.exclude_violation_contexts` | [rule_list.rules.spec.waf_action.app_firewall_detection_control.exclude_violation_contexts](data-sources--service_policy--reference--group-003.md#canonical-1221312223200101-1021231322200023-0320023100020323-3313123222113123-3333120013100121-1323222002111003-1103303311030123-1302011122232211) |
| `rule_list.rules.spec.waf_action.app_firewall_detection_control.exclude_violation_contexts.context` | [rule_list.rules.spec.waf_action.app_firewall_detection_control.exclude_violation_contexts.context](data-sources--service_policy--reference--group-003.md#canonical-1220123322101023-0120030220103020-1221310033032110-0100132033311030-3133332212032200-2003113302323023-2131110233001010-2002332333311212) |
| `rule_list.rules.spec.waf_action.app_firewall_detection_control.exclude_violation_contexts.context_name` | [rule_list.rules.spec.waf_action.app_firewall_detection_control.exclude_violation_contexts.context_name](data-sources--service_policy--reference--group-003.md#canonical-2331132200110012-2323133030002030-1103203213222131-3010030102110001-1020320021031023-1223303310332231-1130221321311100-3001002230313130) |
| `rule_list.rules.spec.waf_action.app_firewall_detection_control.exclude_violation_contexts.exclude_violation` | [rule_list.rules.spec.waf_action.app_firewall_detection_control.exclude_violation_contexts.exclude_violation](data-sources--service_policy--reference--group-003.md#canonical-3321133301001022-3230213101020230-0103202203100010-3231032103230102-1320221302200110-0220231231013003-1000002012123322-0223032212312123) |
| `rule_list.rules.spec.waf_action.none` | [rule_list.rules.spec.waf_action.none](data-sources--service_policy--reference--group-003.md#canonical-2203311033100230-3111003011320013-1332130103303122-1310101021212132-1320110233120033-3232112031330130-3310330322330033-1002322232233121) |
| `rule_list.rules.spec.waf_action.waf_skip_processing` | [rule_list.rules.spec.waf_action.waf_skip_processing](data-sources--service_policy--reference--group-003.md#canonical-0132120133010122-1021113000121033-3232231102111113-0123102230223213-2130123312110203-1132011330211331-2121311032221321-0313103021211210) |
| `server_name` | [server_name](data-sources--service_policy--reference--group-001.md#canonical-0321011233203030-0331130203211121-1031132232322121-0110122031210203-2023032012313300-0320030311200013-1001030211002303-2010110202110133) |
| `server_name_matcher` | [server_name_matcher](data-sources--service_policy--reference--group-003.md#canonical-1031033321123202-3020320322203132-2113112002313103-3203303203110001-3300220113221302-3312210031110132-0122213123201000-2230003120010133) |
| `server_name_matcher.exact_values` | [server_name_matcher.exact_values](data-sources--service_policy--reference--group-003.md#canonical-3102111120301111-0023003010303110-1121230010300330-0330133120331220-3102033131230022-0201123130301130-0033023200330101-0202030222301322) |
| `server_name_matcher.regex_values` | [server_name_matcher.regex_values](data-sources--service_policy--reference--group-003.md#canonical-0033203303333232-2330202232201321-0213131123033320-3010001200021203-3332220323302031-0320023322232000-3231002200233122-3223330123330222) |
| `server_selector` | [server_selector](data-sources--service_policy--reference--group-003.md#canonical-3031212021313110-0302012231302222-1103201333123221-1112231301032212-3312311023100330-0321020012301232-0331330301220320-3112122133101112) |
| `server_selector.expressions` | [server_selector.expressions](data-sources--service_policy--reference--group-003.md#canonical-0220131332331303-2310002010012121-2302003200221023-1132030230311202-2010221001010303-1122120300033102-3322323013300002-0101103301110200) |

<a id="canonical-1310133010132113-2321211101112030-2032301102311223-3121100330132020-0312130331122112-2311320303111131-3110000220310101-2300003320213301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `allow_all_requests` properties

Breadcrumbs:

- [xcsh_service_policy](../data-sources/service_policy.md#canonical-1112022302333333-1130100232232301-3023212310221022-3103031003221200-2030020130230132-1201210321332200-2212311003030121-1321310310032303)
- [Property reference](data-sources--service_policy--reference--group-001.md#canonical-3312102302111103-0221020013211300-2120323123300113-1101123102331230-1032322121222101-1332231300102132-0233300331323303-0031323122032013)
- allow_all_requests

<a id="canonical-3130031112221203-3121213003013120-2300321000003133-3013333202132322-3110122103130021-2102312123313230-2212103110313030-3021332211013333"></a>

Type: `["object", {}]`. Computed.

\[OneOf: allow\_all\_requests, allow\_list, deny\_all\_requests, deny\_list, rule\_list\]
Configuration parameter for allow all requests.

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

- [allow_all_requests](data-sources--service_policy--reference--group-001.md#canonical-3130031112221203-3121213003013120-2300321000003133-3013333202132322-3110122103130021-2102312123313230-2212103110313030-3021332211013333)
- [allow_list](data-sources--service_policy--reference--group-001.md#canonical-3310213202312001-1301321031302013-3201211132300133-2332211120332103-2310000311022100-2111212212131121-0103323001220022-0232310310100313)
- [deny_all_requests](data-sources--service_policy--reference--group-001.md#canonical-1233003033211320-0320102302012323-0122002032101001-0303221201213302-3202101031101002-1123032012131000-3030133100001030-3230132122233001)
- [deny_list](data-sources--service_policy--reference--group-001.md#canonical-1022223021201211-2230013123132233-1220120331103012-2232310222203111-3331022203003133-1331133111223111-3330021021323023-2312130331321003)
- [rule_list](data-sources--service_policy--reference--group-001.md#canonical-1323033222302302-3111221111132320-3301101131323321-3012033101003230-1210201102233222-0102310331101011-1003003201123022-2202030210300232)

Select alternatives according to the provider validators above.

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3210232132133332-2232321312121110-0111111310110022-2030312322330300-2002320322232323-1133002023213102-2130223332331111-3332033011101311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `allow_list` properties

Breadcrumbs:

- [xcsh_service_policy](../data-sources/service_policy.md#canonical-1112022302333333-1130100232232301-3023212310221022-3103031003221200-2030020130230132-1201210321332200-2212311003030121-1321310310032303)
- [Property reference](data-sources--service_policy--reference--group-001.md#canonical-3312102302111103-0221020013211300-2120323123300113-1101123102331230-1032322121222101-1332231300102132-0233300331323303-0031323122032013)
- allow_list

<a id="canonical-3310213202312001-1301321031302013-3201211132300133-2332211120332103-2310000311022100-2111212212131121-0103323001220022-0232310310100313"></a>

Type: `"single"`. Computed.

List of sources. A request belongs to this list if it satisfies any of the match criteria.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-default_action_choice": "[\"default_action_allow\",\"default_action_deny\",\"default_action_next_policy\"]"
}
```

<a id="canonical-3302310102112202-2233122010132013-1112213332230000-2313013133313230-2211012322311021-3131333112332212-0312332302000103-2200112222030222"></a>

### Direct properties for `allow_list`

- [asn_list](data-sources--service_policy--reference--group-001.md#canonical-1030023303231232-1203220100313322-1201010102120033-2212032011303213-1002012301022210-2000311031200033-1013133002020321-3332330213230000): complete subsection reference.

- [asn_set](data-sources--service_policy--reference--group-001.md#canonical-2300030202120333-3232001100201110-3312133322102112-1321122132012201-3321011130110232-0322102203130031-3132022230313230-2230311020310000): complete subsection reference.

<a id="canonical-3201223210332200-0210030310212133-1013320230232111-3133223331303232-0320233110212313-1101232230010231-3301233222201211-1310222323131220"></a>

<a id="canonical-3121032103121001-3011220133303303-3203311233320021-0312131110232331-3003131021202012-3300000011231113-0323232130321030-1321231022302332"></a>

#### `allow_list.country_list` property

Type: `["list", "string"]`. Computed.

\[Enum:
COUNTRY\_NONE|COUNTRY\_AD|COUNTRY\_AE|COUNTRY\_AF|COUNTRY\_AG|COUNTRY\_AI|COUNTRY\_AL|COUNTRY\_AM|COUNTRY\_AN|COUNTRY\_AO|COUNTRY\_AQ|COUNTRY\_AR|COUNTRY\_AS|COUNTRY\_AT|COUNTRY\_AU|COUNTRY\_AW|COUNTRY\_AX|COUNTRY\_AZ|COUNTRY\_BA|COUNTRY\_BB|COUNTRY\_BD|COUNTRY\_BE|COUNTRY\_BF|COUNTRY\_BG|COUNTRY\_BH|COUNTRY\_BI|COUNTRY\_BJ|COUNTRY\_BL|COUNTRY\_BM|COUNTRY\_BN|COUNTRY\_BO|COUNTRY\_BQ|COUNTRY\_BR|COUNTRY\_BS|COUNTRY\_BT|COUNTRY\_BV|COUNTRY\_BW|COUNTRY\_BY|COUNTRY\_BZ|COUNTRY\_CA|COUNTRY\_CC|COUNTRY\_CD|COUNTRY\_CF|COUNTRY\_CG|COUNTRY\_CH|COUNTRY\_CI|COUNTRY\_CK|COUNTRY\_CL|COUNTRY\_CM|COUNTRY\_CN|COUNTRY\_CO|COUNTRY\_CR|COUNTRY\_CS|COUNTRY\_CU|COUNTRY\_CV|COUNTRY\_CW|COUNTRY\_CX|COUNTRY\_CY|COUNTRY\_CZ|COUNTRY\_DE|COUNTRY\_DJ|COUNTRY\_DK|COUNTRY\_DM|COUNTRY\_DO|COUNTRY\_DZ|COUNTRY\_EC|COUNTRY\_EE|COUNTRY\_EG|COUNTRY\_EH|COUNTRY\_ER|COUNTRY\_ES|COUNTRY\_ET|COUNTRY\_FI|COUNTRY\_FJ|COUNTRY\_FK|COUNTRY\_FM|COUNTRY\_FO|COUNTRY\_FR|COUNTRY\_GA|COUNTRY\_GB|COUNTRY\_GD|COUNTRY\_GE|COUNTRY\_GF|COUNTRY\_GG|COUNTRY\_GH|COUNTRY\_GI|COUNTRY\_GL|COUNTRY\_GM|COUNTRY\_GN|COUNTRY\_GP|COUNTRY\_GQ|COUNTRY\_GR|COUNTRY\_GS|COUNTRY\_GT|COUNTRY\_GU|COUNTRY\_GW|COUNTRY\_GY|COUNTRY\_HK|COUNTRY\_HM|COUNTRY\_HN|COUNTRY\_HR|COUNTRY\_HT|COUNTRY\_HU|COUNTRY\_ID|COUNTRY\_IE|COUNTRY\_IL|COUNTRY\_IM|COUNTRY\_IN|COUNTRY\_IO|COUNTRY\_IQ|COUNTRY\_IR|COUNTRY\_IS|COUNTRY\_IT|COUNTRY\_JE|COUNTRY\_JM|COUNTRY\_JO|COUNTRY\_JP|COUNTRY\_KE|COUNTRY\_KG|COUNTRY\_KH|COUNTRY\_KI|COUNTRY\_KM|COUNTRY\_KN|COUNTRY\_KP|COUNTRY\_KR|COUNTRY\_KW|COUNTRY\_KY|COUNTRY\_KZ|COUNTRY\_LA|COUNTRY\_LB|COUNTRY\_LC|COUNTRY\_LI|COUNTRY\_LK|COUNTRY\_LR|COUNTRY\_LS|COUNTRY\_LT|COUNTRY\_LU|COUNTRY\_LV|COUNTRY\_LY|COUNTRY\_MA|COUNTRY\_MC|COUNTRY\_MD|COUNTRY\_ME|COUNTRY\_MF|COUNTRY\_MG|COUNTRY\_MH|COUNTRY\_MK|COUNTRY\_ML|COUNTRY\_MM|COUNTRY\_MN|COUNTRY\_MO|COUNTRY\_MP|COUNTRY\_MQ|COUNTRY\_MR|COUNTRY\_MS|COUNTRY\_MT|COUNTRY\_MU|COUNTRY\_MV|COUNTRY\_MW|COUNTRY\_MX|COUNTRY\_MY|COUNTRY\_MZ|COUNTRY\_NA|COUNTRY\_NC|COUNTRY\_NE|COUNTRY\_NF|COUNTRY\_NG|COUNTRY\_NI|COUNTRY\_NL|COUNTRY\_NO|COUNTRY\_NP|COUNTRY\_NR|COUNTRY\_NU|COUNTRY\_NZ|COUNTRY\_OM|COUNTRY\_PA|COUNTRY\_PE|COUNTRY\_PF|COUNTRY\_PG|COUNTRY\_PH|COUNTRY\_PK|COUNTRY\_PL|COUNTRY\_PM|COUNTRY\_PN|COUNTRY\_PR|COUNTRY\_PS|COUNTRY\_PT|COUNTRY\_PW|COUNTRY\_PY|COUNTRY\_QA|COUNTRY\_RE|COUNTRY\_RO|COUNTRY\_RS|COUNTRY\_RU|COUNTRY\_RW|COUNTRY\_SA|COUNTRY\_SB|COUNTRY\_SC|COUNTRY\_SD|COUNTRY\_SE|COUNTRY\_SG|COUNTRY\_SH|COUNTRY\_SI|COUNTRY\_SJ|COUNTRY\_SK|COUNTRY\_SL|COUNTRY\_SM|COUNTRY\_SN|COUNTRY\_SO|COUNTRY\_SR|COUNTRY\_SS|COUNTRY\_ST|COUNTRY\_SV|COUNTRY\_SX|COUNTRY\_SY|COUNTRY\_SZ|COUNTRY\_TC|COUNTRY\_TD|COUNTRY\_TF|COUNTRY\_TG|COUNTRY\_TH|COUNTRY\_TJ|COUNTRY\_TK|COUNTRY\_TL|COUNTRY\_TM|COUNTRY\_TN|COUNTRY\_TO|COUNTRY\_TR|COUNTRY\_TT|COUNTRY\_TV|COUNTRY\_TW|COUNTRY\_TZ|COUNTRY\_UA|COUNTRY\_UG|COUNTRY\_UM|COUNTRY\_US|COUNTRY\_UY|COUNTRY\_UZ|COUNTRY\_VA|COUNTRY\_VC|COUNTRY\_VE|COUNTRY\_VG|COUNTRY\_VI|COUNTRY\_VN|COUNTRY\_VU|COUNTRY\_WF|COUNTRY\_WS|COUNTRY\_XK|COUNTRY\_XT|COUNTRY\_YE|COUNTRY\_YT|COUNTRY\_ZA|COUNTRY\_ZM|COUNTRY\_ZW\]
Addresses that belong to one of the countries in the given list The country is obtained by
performing a lookup for the source IPv4 Address in a GeoIP DB. Possible values are
\`COUNTRY\_NONE\`, \`COUNTRY\_AD\`, \`COUNTRY\_AE\`, \`COUNTRY\_AF\`, \`COUNTRY\_AG\`,
\`COUNTRY\_AI\`, \`COUNTRY\_AL\`, \`COUNTRY\_AM\`, \`COUNTRY\_AN\`, \`COUNTRY\_AO\`,
\`COUNTRY\_AQ\`, \`COUNTRY\_AR\`, \`COUNTRY\_AS\`, \`COUNTRY\_AT\`, \`COUNTRY\_AU\`,
\`COUNTRY\_AW\`, \`COUNTRY\_AX\`, \`COUNTRY\_AZ\`, \`COUNTRY\_BA\`, \`COUNTRY\_BB\`,
\`COUNTRY\_BD\`, \`COUNTRY\_BE\`, \`COUNTRY\_BF\`, \`COUNTRY\_BG\`, \`COUNTRY\_BH\`,
\`COUNTRY\_BI\`, \`COUNTRY\_BJ\`, \`COUNTRY\_BL\`, \`COUNTRY\_BM\`, \`COUNTRY\_BN\`,
\`COUNTRY\_BO\`, \`COUNTRY\_BQ\`, \`COUNTRY\_BR\`, \`COUNTRY\_BS\`, \`COUNTRY\_BT\`,
\`COUNTRY\_BV\`, \`COUNTRY\_BW\`, \`COUNTRY\_BY\`, \`COUNTRY\_BZ\`, \`COUNTRY\_CA\`,
\`COUNTRY\_CC\`, \`COUNTRY\_CD\`, \`COUNTRY\_CF\`, \`COUNTRY\_CG\`, \`COUNTRY\_CH\`,
\`COUNTRY\_CI\`, \`COUNTRY\_CK\`, \`COUNTRY\_CL\`, \`COUNTRY\_CM\`, \`COUNTRY\_CN\`,
\`COUNTRY\_CO\`, \`COUNTRY\_CR\`, \`COUNTRY\_CS\`, \`COUNTRY\_CU\`, \`COUNTRY\_CV\`,
\`COUNTRY\_CW\`, \`COUNTRY\_CX\`, \`COUNTRY\_CY\`, \`COUNTRY\_CZ\`, \`COUNTRY\_DE\`,
\`COUNTRY\_DJ\`, \`COUNTRY\_DK\`, \`COUNTRY\_DM\`, \`COUNTRY\_DO\`, \`COUNTRY\_DZ\`,
\`COUNTRY\_EC\`, \`COUNTRY\_EE\`, \`COUNTRY\_EG\`, \`COUNTRY\_EH\`, \`COUNTRY\_ER\`,
\`COUNTRY\_ES\`, \`COUNTRY\_ET\`, \`COUNTRY\_FI\`, \`COUNTRY\_FJ\`, \`COUNTRY\_FK\`,
\`COUNTRY\_FM\`, \`COUNTRY\_FO\`, \`COUNTRY\_FR\`, \`COUNTRY\_GA\`, \`COUNTRY\_GB\`,
\`COUNTRY\_GD\`, \`COUNTRY\_GE\`, \`COUNTRY\_GF\`, \`COUNTRY\_GG\`, \`COUNTRY\_GH\`,
\`COUNTRY\_GI\`, \`COUNTRY\_GL\`, \`COUNTRY\_GM\`, \`COUNTRY\_GN\`, \`COUNTRY\_GP\`,
\`COUNTRY\_GQ\`, \`COUNTRY\_GR\`, \`COUNTRY\_GS\`, \`COUNTRY\_GT\`, \`COUNTRY\_GU\`,
\`COUNTRY\_GW\`, \`COUNTRY\_GY\`, \`COUNTRY\_HK\`, \`COUNTRY\_HM\`, \`COUNTRY\_HN\`,
\`COUNTRY\_HR\`, \`COUNTRY\_HT\`, \`COUNTRY\_HU\`, \`COUNTRY\_ID\`, \`COUNTRY\_IE\`,
\`COUNTRY\_IL\`, \`COUNTRY\_IM\`, \`COUNTRY\_IN\`, \`COUNTRY\_IO\`, \`COUNTRY\_IQ\`,
\`COUNTRY\_IR\`, \`COUNTRY\_IS\`, \`COUNTRY\_IT\`, \`COUNTRY\_JE\`, \`COUNTRY\_JM\`,
\`COUNTRY\_JO\`, \`COUNTRY\_JP\`, \`COUNTRY\_KE\`, \`COUNTRY\_KG\`, \`COUNTRY\_KH\`,
\`COUNTRY\_KI\`, \`COUNTRY\_KM\`, \`COUNTRY\_KN\`, \`COUNTRY\_KP\`, \`COUNTRY\_KR\`,
\`COUNTRY\_KW\`, \`COUNTRY\_KY\`, \`COUNTRY\_KZ\`, \`COUNTRY\_LA\`, \`COUNTRY\_LB\`,
\`COUNTRY\_LC\`, \`COUNTRY\_LI\`, \`COUNTRY\_LK\`, \`COUNTRY\_LR\`, \`COUNTRY\_LS\`,
\`COUNTRY\_LT\`, \`COUNTRY\_LU\`, \`COUNTRY\_LV\`, \`COUNTRY\_LY\`, \`COUNTRY\_MA\`,
\`COUNTRY\_MC\`, \`COUNTRY\_MD\`, \`COUNTRY\_ME\`, \`COUNTRY\_MF\`, \`COUNTRY\_MG\`,
\`COUNTRY\_MH\`, \`COUNTRY\_MK\`, \`COUNTRY\_ML\`, \`COUNTRY\_MM\`, \`COUNTRY\_MN\`,
\`COUNTRY\_MO\`, \`COUNTRY\_MP\`, \`COUNTRY\_MQ\`, \`COUNTRY\_MR\`, \`COUNTRY\_MS\`,
\`COUNTRY\_MT\`, \`COUNTRY\_MU\`, \`COUNTRY\_MV\`, \`COUNTRY\_MW\`, \`COUNTRY\_MX\`,
\`COUNTRY\_MY\`, \`COUNTRY\_MZ\`, \`COUNTRY\_NA\`, \`COUNTRY\_NC\`, \`COUNTRY\_NE\`,
\`COUNTRY\_NF\`, \`COUNTRY\_NG\`, \`COUNTRY\_NI\`, \`COUNTRY\_NL\`, \`COUNTRY\_NO\`,
\`COUNTRY\_NP\`, \`COUNTRY\_NR\`, \`COUNTRY\_NU\`, \`COUNTRY\_NZ\`, \`COUNTRY\_OM\`,
\`COUNTRY\_PA\`, \`COUNTRY\_PE\`, \`COUNTRY\_PF\`, \`COUNTRY\_PG\`, \`COUNTRY\_PH\`,
\`COUNTRY\_PK\`, \`COUNTRY\_PL\`, \`COUNTRY\_PM\`, \`COUNTRY\_PN\`, \`COUNTRY\_PR\`,
\`COUNTRY\_PS\`, \`COUNTRY\_PT\`, \`COUNTRY\_PW\`, \`COUNTRY\_PY\`, \`COUNTRY\_QA\`,
\`COUNTRY\_RE\`, \`COUNTRY\_RO\`, \`COUNTRY\_RS\`, \`COUNTRY\_RU\`, \`COUNTRY\_RW\`,
\`COUNTRY\_SA\`, \`COUNTRY\_SB\`, \`COUNTRY\_SC\`, \`COUNTRY\_SD\`, \`COUNTRY\_SE\`,
\`COUNTRY\_SG\`, \`COUNTRY\_SH\`, \`COUNTRY\_SI\`, \`COUNTRY\_SJ\`, \`COUNTRY\_SK\`,
\`COUNTRY\_SL\`, \`COUNTRY\_SM\`, \`COUNTRY\_SN\`, \`COUNTRY\_SO\`, \`COUNTRY\_SR\`,
\`COUNTRY\_SS\`, \`COUNTRY\_ST\`, \`COUNTRY\_SV\`, \`COUNTRY\_SX\`, \`COUNTRY\_SY\`,
\`COUNTRY\_SZ\`, \`COUNTRY\_TC\`, \`COUNTRY\_TD\`, \`COUNTRY\_TF\`, \`COUNTRY\_TG\`,
\`COUNTRY\_TH\`, \`COUNTRY\_TJ\`, \`COUNTRY\_TK\`, \`COUNTRY\_TL\`, \`COUNTRY\_TM\`,
\`COUNTRY\_TN\`, \`COUNTRY\_TO\`, \`COUNTRY\_TR\`, \`COUNTRY\_TT\`, \`COUNTRY\_TV\`,
\`COUNTRY\_TW\`, \`COUNTRY\_TZ\`, \`COUNTRY\_UA\`, \`COUNTRY\_UG\`, \`COUNTRY\_UM\`,
\`COUNTRY\_US\`, \`COUNTRY\_UY\`, \`COUNTRY\_UZ\`, \`COUNTRY\_VA\`, \`COUNTRY\_VC\`,
\`COUNTRY\_VE\`, \`COUNTRY\_VG\`, \`COUNTRY\_VI\`, \`COUNTRY\_VN\`, \`COUNTRY\_VU\`,
\`COUNTRY\_WF\`, \`COUNTRY\_WS\`, \`COUNTRY\_XK\`, \`COUNTRY\_XT\`, \`COUNTRY\_YE\`,
\`COUNTRY\_YT\`, \`COUNTRY\_ZA\`, \`COUNTRY\_ZM\`, \`COUNTRY\_ZW\`. Defaults to \`COUNTRY\_NONE\`.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.items.enum.not_in": "[0]",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.items.enum.not_in": "[0]",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [default_action_allow](data-sources--service_policy--reference--group-001.md#canonical-1301032303223002-2333121011032232-3221312123103230-0002332203102010-1213011203333311-3100302221210211-2230101103022102-0102030110323210): complete subsection reference.

- [default_action_deny](data-sources--service_policy--reference--group-001.md#canonical-2222313323121033-0013032032201330-0223330301212312-1200331323300022-2232311323031023-3311101132202333-0200033320310320-1122033110021000): complete subsection reference.

- [default_action_next_policy](data-sources--service_policy--reference--group-001.md#canonical-0123210213101221-1310012000232320-0110030303011022-1133113300201030-0002301300113223-0121203230123330-1202001302332131-0331211230120212): complete subsection reference.

- [ip_prefix_set](data-sources--service_policy--reference--group-001.md#canonical-3331231301111102-2131121010003012-1202332302000012-0200103120202213-1003003323311301-2301002131322213-1111121210233122-2112223211322013): complete subsection reference.

- [prefix_list](data-sources--service_policy--reference--group-001.md#canonical-0323021122221023-1303210122113001-3201333131021030-2021211223130210-3002312031023211-0113002000303112-3101313003210323-0333030021002103): complete subsection reference.

<a id="canonical-2200032110111002-1232302122323330-0102211101331001-2111120100003320-0222132102301333-2320033332131303-0111121102231032-1111131000112002"></a>

<a id="canonical-0202033201000020-2002122112120330-2130130313233310-3003122211313121-0311002013213013-0022100230322321-3202023301123123-3223010113100313"></a>

#### `allow_list.tls_fingerprint_classes` property

Type: `["list", "string"]`. Computed.

\[Enum:
TLS\_FINGERPRINT\_NONE|ANY\_MALICIOUS\_FINGERPRINT|ADWARE|ADWIND|DRIDEX|GOOTKIT|GOZI|JBIFROST|QUAKBOT|RANSOMWARE|TROLDESH|TOFSEE|TORRENTLOCKER|TRICKBOT\]
List of known classes of TLS fingerprints to match the input TLS JA3 fingerprint against. Possible
values are \`TLS\_FINGERPRINT\_NONE\`, \`ANY\_MALICIOUS\_FINGERPRINT\`, \`ADWARE\`, \`ADWIND\`,
\`DRIDEX\`, \`GOOTKIT\`, \`GOZI\`, \`JBIFROST\`, \`QUAKBOT\`, \`RANSOMWARE\`, \`TROLDESH\`,
\`TOFSEE\`, \`TORRENTLOCKER\`, \`TRICKBOT\`. Defaults to \`TLS\_FINGERPRINT\_NONE\`.

Additional upstream details:

A list of known classes of TLS fingerprints to match the input TLS JA3 fingerprint against.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-2103232203012222-3010303310201122-2213221301222313-1300301201221030-2030213031211021-3123201120130323-0002110030232031-2113123103033301"></a>

<a id="canonical-0102123200303210-3220130230003011-2200131032213310-2220010110322111-2033221331323111-1332221012232210-1131000000113233-1323303101100330"></a>

#### `allow_list.tls_fingerprint_values` property

Type: `["list", "string"]`. Computed.

A list of exact TLS JA3 fingerprints to match the input TLS JA3 fingerprint against.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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
    "ves.io.schema.rules.repeated.items.string.len": "32",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.len": "32",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1030023303231232-1203220100313322-1201010102120033-2212032011303213-1002012301022210-2000311031200033-1013133002020321-3332330213230000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `allow_list.asn_list` properties

Breadcrumbs:

- [xcsh_service_policy](../data-sources/service_policy.md#canonical-1112022302333333-1130100232232301-3023212310221022-3103031003221200-2030020130230132-1201210321332200-2212311003030121-1321310310032303)
- [Property reference](data-sources--service_policy--reference--group-001.md#canonical-3312102302111103-0221020013211300-2120323123300113-1101123102331230-1032322121222101-1332231300102132-0233300331323303-0031323122032013)
- [allow_list](data-sources--service_policy--reference--group-001.md#canonical-3210232132133332-2232321312121110-0111111310110022-2030312322330300-2002320322232323-1133002023213102-2130223332331111-3332033011101311)
- allow_list.asn_list

<a id="canonical-1312111133013001-0010333311120122-2202201333312132-2110322133223101-0202130220103231-1122012112132031-2122112232221223-1130221213013113"></a>

Type: `"single"`. Computed.

An unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny
lists for use in network policy or service policy. It can be used to create the allow list only for
DNS Load Balancer.

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

<a id="canonical-0013132233000301-1311102020001322-1201220300313230-2131230311000203-0010121202321023-3122123111102000-3110322321303223-2013302211202010"></a>

### Direct properties for `allow_list.asn_list`

<a id="canonical-2132211032001022-3322120022132121-2223203200103133-3333123030310232-2330103300033022-2122210021302323-0110230302332032-3033123210210331"></a>

#### `allow_list.asn_list.as_numbers` property

Type: `["list", "number"]`. Computed.

An unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny
lists for use in network policy or service policy. It can be used to create the allow list only for
DNS Load Balancer.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 16,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 16,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-06T12:36:10+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2300030202120333-3232001100201110-3312133322102112-1321122132012201-3321011130110232-0322102203130031-3132022230313230-2230311020310000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `allow_list.asn_set` properties

Breadcrumbs:

- [xcsh_service_policy](../data-sources/service_policy.md#canonical-1112022302333333-1130100232232301-3023212310221022-3103031003221200-2030020130230132-1201210321332200-2212311003030121-1321310310032303)
- [Property reference](data-sources--service_policy--reference--group-001.md#canonical-3312102302111103-0221020013211300-2120323123300113-1101123102331230-1032322121222101-1332231300102132-0233300331323303-0031323122032013)
- [allow_list](data-sources--service_policy--reference--group-001.md#canonical-3210232132133332-2232321312121110-0111111310110022-2030312322330300-2002320322232323-1133002023213102-2130223332331111-3332033011101311)
- allow_list.asn_set

<a id="canonical-1333322332223310-0113001220232202-1030100121101210-1310033023330202-0213121232200211-1311003320013113-1211220210131201-3010220113203230"></a>

Type: `"list"`. Computed.

Addresses that belong to the ASNs in the given bgp\_asn\_set The ASN is obtained by performing a
lookup for the source IPv4 Address in a GeoIP DB.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-2211011001301211-0321323112232011-2110120212333302-0322200321101011-2311033023302103-2121032301302222-0130102302213321-2301120030330103"></a>

### Direct properties for `allow_list.asn_set`

<a id="canonical-2221233200011120-0303202320311030-2010201213030012-3131030100213023-0101201103313300-1033323231231210-1233021333033231-3213110030013332"></a>

#### `allow_list.asn_set.name` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-0220223321300300-2111330112030220-0033233310031130-0333311011302202-0102112222320131-3201101201011310-0301002012111222-2231102212310203"></a>

<a id="canonical-2332131202313200-2033313033100103-0110232332320331-3210022012310220-3332321123301332-2133022321210130-0220100020201200-1321320010023121"></a>

#### `allow_list.asn_set.namespace` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-3101003101330333-2031030003200212-2300322231023111-2112133320223232-1123030223320230-1330322102132031-0120023220220123-2020200002200212"></a>

<a id="canonical-1220101023320121-1302113130300022-0320001220211331-3313003112132121-0110303201232013-3110132010021311-3111230330131012-2113223312002011"></a>

#### `allow_list.asn_set.tenant` property

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-1301032303223002-2333121011032232-3221312123103230-0002332203102010-1213011203333311-3100302221210211-2230101103022102-0102030110323210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `allow_list.default_action_allow` properties

Breadcrumbs:

- [xcsh_service_policy](../data-sources/service_policy.md#canonical-1112022302333333-1130100232232301-3023212310221022-3103031003221200-2030020130230132-1201210321332200-2212311003030121-1321310310032303)
- [Property reference](data-sources--service_policy--reference--group-001.md#canonical-3312102302111103-0221020013211300-2120323123300113-1101123102331230-1032322121222101-1332231300102132-0233300331323303-0031323122032013)
- [allow_list](data-sources--service_policy--reference--group-001.md#canonical-3210232132133332-2232321312121110-0111111310110022-2030312322330300-2002320322232323-1133002023213102-2130223332331111-3332033011101311)
- allow_list.default_action_allow

<a id="canonical-2322220103311102-1112100020121201-3113332022132223-1220112123203211-2313331003031201-0032310312003122-2330011333301020-0333210132201312"></a>

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

<a id="canonical-2222313323121033-0013032032201330-0223330301212312-1200331323300022-2232311323031023-3311101132202333-0200033320310320-1122033110021000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `allow_list.default_action_deny` properties

Breadcrumbs:

- [xcsh_service_policy](../data-sources/service_policy.md#canonical-1112022302333333-1130100232232301-3023212310221022-3103031003221200-2030020130230132-1201210321332200-2212311003030121-1321310310032303)
- [Property reference](data-sources--service_policy--reference--group-001.md#canonical-3312102302111103-0221020013211300-2120323123300113-1101123102331230-1032322121222101-1332231300102132-0233300331323303-0031323122032013)
- [allow_list](data-sources--service_policy--reference--group-001.md#canonical-3210232132133332-2232321312121110-0111111310110022-2030312322330300-2002320322232323-1133002023213102-2130223332331111-3332033011101311)
- allow_list.default_action_deny

<a id="canonical-1303000123333311-1031110003101232-1221003200013133-2201200210111220-3030330232131002-2213321312330133-3122013323131313-2020111311131332"></a>

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

<a id="canonical-0123210213101221-1310012000232320-0110030303011022-1133113300201030-0002301300113223-0121203230123330-1202001302332131-0331211230120212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `allow_list.default_action_next_policy` properties

Breadcrumbs:

- [xcsh_service_policy](../data-sources/service_policy.md#canonical-1112022302333333-1130100232232301-3023212310221022-3103031003221200-2030020130230132-1201210321332200-2212311003030121-1321310310032303)
- [Property reference](data-sources--service_policy--reference--group-001.md#canonical-3312102302111103-0221020013211300-2120323123300113-1101123102331230-1032322121222101-1332231300102132-0233300331323303-0031323122032013)
- [allow_list](data-sources--service_policy--reference--group-001.md#canonical-3210232132133332-2232321312121110-0111111310110022-2030312322330300-2002320322232323-1133002023213102-2130223332331111-3332033011101311)
- allow_list.default_action_next_policy

<a id="canonical-0012323302131020-3003331031113212-2321012000012330-3321333330122233-3130132221312232-2322011113002330-0131120312133222-2211032312030002"></a>

Type: `["object", {}]`. Computed.

Policy configuration for this feature.

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

<a id="canonical-3331231301111102-2131121010003012-1202332302000012-0200103120202213-1003003323311301-2301002131322213-1111121210233122-2112223211322013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `allow_list.ip_prefix_set` properties

Breadcrumbs:

- [xcsh_service_policy](../data-sources/service_policy.md#canonical-1112022302333333-1130100232232301-3023212310221022-3103031003221200-2030020130230132-1201210321332200-2212311003030121-1321310310032303)
- [Property reference](data-sources--service_policy--reference--group-001.md#canonical-3312102302111103-0221020013211300-2120323123300113-1101123102331230-1032322121222101-1332231300102132-0233300331323303-0031323122032013)
- [allow_list](data-sources--service_policy--reference--group-001.md#canonical-3210232132133332-2232321312121110-0111111310110022-2030312322330300-2002320322232323-1133002023213102-2130223332331111-3332033011101311)
- allow_list.ip_prefix_set

<a id="canonical-0320313123101223-1031332321320123-2203001322002023-1122002013132033-1213300131001210-2003303030321321-1320101123222000-0102100130012130"></a>

Type: `"list"`. Computed.

Addresses that are covered by the prefixes in the given ip\_prefix\_set.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-2311332120030333-1323232120230301-3202113232132230-2321300030021011-0022210233232221-3021123122312230-0111302101331232-3033210320030233"></a>

### Direct properties for `allow_list.ip_prefix_set`

<a id="canonical-3313001222300132-0112013233111320-3123031221030121-0022331031223121-3033123032302333-1331222302310233-0002232221013032-0220132213323321"></a>

#### `allow_list.ip_prefix_set.name` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-3323002203313220-2021332210200302-1310110331330112-2322203033010112-2033232330322213-0301233211322321-3203203010023103-1313212020321223"></a>

<a id="canonical-3100202220203100-1333303101030013-0230233313333031-0311011212110311-0122323122313021-0321111110323101-2133110020322213-3120332310221211"></a>

#### `allow_list.ip_prefix_set.namespace` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-0102323230313202-3330101233033303-0102021200202000-1023013233232320-3130010303332323-1200000320323223-3223203013103213-2120123321311311"></a>

<a id="canonical-1212113123122020-3133212312230301-2032302200223310-1211311131000103-3203000131312022-1331110110023030-1011131132222003-1011011132232303"></a>

#### `allow_list.ip_prefix_set.tenant` property

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-0323021122221023-1303210122113001-3201333131021030-2021211223130210-3002312031023211-0113002000303112-3101313003210323-0333030021002103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `allow_list.prefix_list` properties

Breadcrumbs:

- [xcsh_service_policy](../data-sources/service_policy.md#canonical-1112022302333333-1130100232232301-3023212310221022-3103031003221200-2030020130230132-1201210321332200-2212311003030121-1321310310032303)
- [Property reference](data-sources--service_policy--reference--group-001.md#canonical-3312102302111103-0221020013211300-2120323123300113-1101123102331230-1032322121222101-1332231300102132-0233300331323303-0031323122032013)
- [allow_list](data-sources--service_policy--reference--group-001.md#canonical-3210232132133332-2232321312121110-0111111310110022-2030312322330300-2002320322232323-1133002023213102-2130223332331111-3332033011101311)
- allow_list.prefix_list

<a id="canonical-2011232121313112-2310112313330123-0331312012133013-2112232220323120-0102312002131322-3132131003330101-3222231022030223-0221012130233322"></a>

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

<a id="canonical-2031221221022011-2033110230320320-3132232023220100-3223100300323132-1011300122000211-0131132122012301-0022320320211230-2333010300111322"></a>

### Direct properties for `allow_list.prefix_list`

<a id="canonical-0233220002202113-0012322213010201-2300000130020120-0123220121012031-1322300320222220-2301121202201031-3312122230120123-2031332203122221"></a>

#### `allow_list.prefix_list.prefixes` property

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-3101033101321213-3011013020303022-2020233002120302-0120121332310033-2321321210322320-3100211331032333-1210331221301022-3010231111133032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `any_server` properties

Breadcrumbs:

- [xcsh_service_policy](../data-sources/service_policy.md#canonical-1112022302333333-1130100232232301-3023212310221022-3103031003221200-2030020130230132-1201210321332200-2212311003030121-1321310310032303)
- [Property reference](data-sources--service_policy--reference--group-001.md#canonical-3312102302111103-0221020013211300-2120323123300113-1101123102331230-1032322121222101-1332231300102132-0233300331323303-0031323122032013)
- any_server

<a id="canonical-2213123300100210-0321113302220212-0023213232002331-3312020330323132-3010230320130302-3300323232110113-0333123031020303-0132100313012303"></a>

Type: `["object", {}]`. Computed.

\[OneOf: any\_server, server\_name, server\_name\_matcher, server\_selector\] Enable this option.
Defaults to \`map\[\]\`. Server applies default when omitted.

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

- [any_server](data-sources--service_policy--reference--group-001.md#canonical-2213123300100210-0321113302220212-0023213232002331-3312020330323132-3010230320130302-3300323232110113-0333123031020303-0132100313012303)
- [server_name](data-sources--service_policy--reference--group-001.md#canonical-0321011233203030-0331130203211121-1031132232322121-0110122031210203-2023032012313300-0320030311200013-1001030211002303-2010110202110133)
- [server_name_matcher](data-sources--service_policy--reference--group-003.md#canonical-1031033321123202-3020320322203132-2113112002313103-3203303203110001-3300220113221302-3312210031110132-0122213123201000-2230003120010133)
- [server_selector](data-sources--service_policy--reference--group-003.md#canonical-3031212021313110-0302012231302222-1103201333123221-1112231301032212-3312311023100330-0321020012301232-0331330301220320-3112122133101112)

Select alternatives according to the provider validators above.

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1321120122221121-2211201320102032-3231122122100213-0112323220103312-2132331120321323-3313311212102223-3013200033323302-3032330332020102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `deny_all_requests` properties

Breadcrumbs:

- [xcsh_service_policy](../data-sources/service_policy.md#canonical-1112022302333333-1130100232232301-3023212310221022-3103031003221200-2030020130230132-1201210321332200-2212311003030121-1321310310032303)
- [Property reference](data-sources--service_policy--reference--group-001.md#canonical-3312102302111103-0221020013211300-2120323123300113-1101123102331230-1032322121222101-1332231300102132-0233300331323303-0031323122032013)
- deny_all_requests

<a id="canonical-1233003033211320-0320102302012323-0122002032101001-0303221201213302-3202101031101002-1123032012131000-3030133100001030-3230132122233001"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for deny all requests.

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

<a id="canonical-1133210330221300-2302313012231010-0300310022202220-2111123320332210-2101000320030120-3220020323311021-3132220133231132-0021330023303013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `deny_list` properties

Breadcrumbs:

- [xcsh_service_policy](../data-sources/service_policy.md#canonical-1112022302333333-1130100232232301-3023212310221022-3103031003221200-2030020130230132-1201210321332200-2212311003030121-1321310310032303)
- [Property reference](data-sources--service_policy--reference--group-001.md#canonical-3312102302111103-0221020013211300-2120323123300113-1101123102331230-1032322121222101-1332231300102132-0233300331323303-0031323122032013)
- deny_list

<a id="canonical-1022223021201211-2230013123132233-1220120331103012-2232310222203111-3331022203003133-1331133111223111-3330021021323023-2312130331321003"></a>

Type: `"single"`. Computed.

List of sources. A request belongs to this list if it satisfies any of the match criteria.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-default_action_choice": "[\"default_action_allow\",\"default_action_deny\",\"default_action_next_policy\"]"
}
```

<a id="canonical-2013211320032321-1201332113211102-0321011231302312-2110312313321211-1200023032111201-1112313003203132-1111021202301320-2002132233120312"></a>

### Direct properties for `deny_list`

- [asn_list](data-sources--service_policy--reference--group-001.md#canonical-0011132331310313-1312032302323313-3130122322230323-0312311112332011-0301221013302132-2232030313211110-3320321122333033-2011332130311213): complete subsection reference.

- [asn_set](data-sources--service_policy--reference--group-001.md#canonical-3330311320223113-1300223221102113-2100032022322001-3132023001102020-2231000000200300-1032221303212201-3031130220103020-2332321310301200): complete subsection reference.

<a id="canonical-2311232000332011-2201122232020320-1013331123021023-3310120330200230-1022330200320213-0303123130122130-3221302201321333-1202203312110320"></a>

<a id="canonical-3333011101122011-3102211231211033-0320031110102312-3113322311032222-0323130232033031-0213313020313332-3333200003033230-0112333300002000"></a>

#### `deny_list.country_list` property

Type: `["list", "string"]`. Computed.

\[Enum:
COUNTRY\_NONE|COUNTRY\_AD|COUNTRY\_AE|COUNTRY\_AF|COUNTRY\_AG|COUNTRY\_AI|COUNTRY\_AL|COUNTRY\_AM|COUNTRY\_AN|COUNTRY\_AO|COUNTRY\_AQ|COUNTRY\_AR|COUNTRY\_AS|COUNTRY\_AT|COUNTRY\_AU|COUNTRY\_AW|COUNTRY\_AX|COUNTRY\_AZ|COUNTRY\_BA|COUNTRY\_BB|COUNTRY\_BD|COUNTRY\_BE|COUNTRY\_BF|COUNTRY\_BG|COUNTRY\_BH|COUNTRY\_BI|COUNTRY\_BJ|COUNTRY\_BL|COUNTRY\_BM|COUNTRY\_BN|COUNTRY\_BO|COUNTRY\_BQ|COUNTRY\_BR|COUNTRY\_BS|COUNTRY\_BT|COUNTRY\_BV|COUNTRY\_BW|COUNTRY\_BY|COUNTRY\_BZ|COUNTRY\_CA|COUNTRY\_CC|COUNTRY\_CD|COUNTRY\_CF|COUNTRY\_CG|COUNTRY\_CH|COUNTRY\_CI|COUNTRY\_CK|COUNTRY\_CL|COUNTRY\_CM|COUNTRY\_CN|COUNTRY\_CO|COUNTRY\_CR|COUNTRY\_CS|COUNTRY\_CU|COUNTRY\_CV|COUNTRY\_CW|COUNTRY\_CX|COUNTRY\_CY|COUNTRY\_CZ|COUNTRY\_DE|COUNTRY\_DJ|COUNTRY\_DK|COUNTRY\_DM|COUNTRY\_DO|COUNTRY\_DZ|COUNTRY\_EC|COUNTRY\_EE|COUNTRY\_EG|COUNTRY\_EH|COUNTRY\_ER|COUNTRY\_ES|COUNTRY\_ET|COUNTRY\_FI|COUNTRY\_FJ|COUNTRY\_FK|COUNTRY\_FM|COUNTRY\_FO|COUNTRY\_FR|COUNTRY\_GA|COUNTRY\_GB|COUNTRY\_GD|COUNTRY\_GE|COUNTRY\_GF|COUNTRY\_GG|COUNTRY\_GH|COUNTRY\_GI|COUNTRY\_GL|COUNTRY\_GM|COUNTRY\_GN|COUNTRY\_GP|COUNTRY\_GQ|COUNTRY\_GR|COUNTRY\_GS|COUNTRY\_GT|COUNTRY\_GU|COUNTRY\_GW|COUNTRY\_GY|COUNTRY\_HK|COUNTRY\_HM|COUNTRY\_HN|COUNTRY\_HR|COUNTRY\_HT|COUNTRY\_HU|COUNTRY\_ID|COUNTRY\_IE|COUNTRY\_IL|COUNTRY\_IM|COUNTRY\_IN|COUNTRY\_IO|COUNTRY\_IQ|COUNTRY\_IR|COUNTRY\_IS|COUNTRY\_IT|COUNTRY\_JE|COUNTRY\_JM|COUNTRY\_JO|COUNTRY\_JP|COUNTRY\_KE|COUNTRY\_KG|COUNTRY\_KH|COUNTRY\_KI|COUNTRY\_KM|COUNTRY\_KN|COUNTRY\_KP|COUNTRY\_KR|COUNTRY\_KW|COUNTRY\_KY|COUNTRY\_KZ|COUNTRY\_LA|COUNTRY\_LB|COUNTRY\_LC|COUNTRY\_LI|COUNTRY\_LK|COUNTRY\_LR|COUNTRY\_LS|COUNTRY\_LT|COUNTRY\_LU|COUNTRY\_LV|COUNTRY\_LY|COUNTRY\_MA|COUNTRY\_MC|COUNTRY\_MD|COUNTRY\_ME|COUNTRY\_MF|COUNTRY\_MG|COUNTRY\_MH|COUNTRY\_MK|COUNTRY\_ML|COUNTRY\_MM|COUNTRY\_MN|COUNTRY\_MO|COUNTRY\_MP|COUNTRY\_MQ|COUNTRY\_MR|COUNTRY\_MS|COUNTRY\_MT|COUNTRY\_MU|COUNTRY\_MV|COUNTRY\_MW|COUNTRY\_MX|COUNTRY\_MY|COUNTRY\_MZ|COUNTRY\_NA|COUNTRY\_NC|COUNTRY\_NE|COUNTRY\_NF|COUNTRY\_NG|COUNTRY\_NI|COUNTRY\_NL|COUNTRY\_NO|COUNTRY\_NP|COUNTRY\_NR|COUNTRY\_NU|COUNTRY\_NZ|COUNTRY\_OM|COUNTRY\_PA|COUNTRY\_PE|COUNTRY\_PF|COUNTRY\_PG|COUNTRY\_PH|COUNTRY\_PK|COUNTRY\_PL|COUNTRY\_PM|COUNTRY\_PN|COUNTRY\_PR|COUNTRY\_PS|COUNTRY\_PT|COUNTRY\_PW|COUNTRY\_PY|COUNTRY\_QA|COUNTRY\_RE|COUNTRY\_RO|COUNTRY\_RS|COUNTRY\_RU|COUNTRY\_RW|COUNTRY\_SA|COUNTRY\_SB|COUNTRY\_SC|COUNTRY\_SD|COUNTRY\_SE|COUNTRY\_SG|COUNTRY\_SH|COUNTRY\_SI|COUNTRY\_SJ|COUNTRY\_SK|COUNTRY\_SL|COUNTRY\_SM|COUNTRY\_SN|COUNTRY\_SO|COUNTRY\_SR|COUNTRY\_SS|COUNTRY\_ST|COUNTRY\_SV|COUNTRY\_SX|COUNTRY\_SY|COUNTRY\_SZ|COUNTRY\_TC|COUNTRY\_TD|COUNTRY\_TF|COUNTRY\_TG|COUNTRY\_TH|COUNTRY\_TJ|COUNTRY\_TK|COUNTRY\_TL|COUNTRY\_TM|COUNTRY\_TN|COUNTRY\_TO|COUNTRY\_TR|COUNTRY\_TT|COUNTRY\_TV|COUNTRY\_TW|COUNTRY\_TZ|COUNTRY\_UA|COUNTRY\_UG|COUNTRY\_UM|COUNTRY\_US|COUNTRY\_UY|COUNTRY\_UZ|COUNTRY\_VA|COUNTRY\_VC|COUNTRY\_VE|COUNTRY\_VG|COUNTRY\_VI|COUNTRY\_VN|COUNTRY\_VU|COUNTRY\_WF|COUNTRY\_WS|COUNTRY\_XK|COUNTRY\_XT|COUNTRY\_YE|COUNTRY\_YT|COUNTRY\_ZA|COUNTRY\_ZM|COUNTRY\_ZW\]
Addresses that belong to one of the countries in the given list The country is obtained by
performing a lookup for the source IPv4 Address in a GeoIP DB. Possible values are
\`COUNTRY\_NONE\`, \`COUNTRY\_AD\`, \`COUNTRY\_AE\`, \`COUNTRY\_AF\`, \`COUNTRY\_AG\`,
\`COUNTRY\_AI\`, \`COUNTRY\_AL\`, \`COUNTRY\_AM\`, \`COUNTRY\_AN\`, \`COUNTRY\_AO\`,
\`COUNTRY\_AQ\`, \`COUNTRY\_AR\`, \`COUNTRY\_AS\`, \`COUNTRY\_AT\`, \`COUNTRY\_AU\`,
\`COUNTRY\_AW\`, \`COUNTRY\_AX\`, \`COUNTRY\_AZ\`, \`COUNTRY\_BA\`, \`COUNTRY\_BB\`,
\`COUNTRY\_BD\`, \`COUNTRY\_BE\`, \`COUNTRY\_BF\`, \`COUNTRY\_BG\`, \`COUNTRY\_BH\`,
\`COUNTRY\_BI\`, \`COUNTRY\_BJ\`, \`COUNTRY\_BL\`, \`COUNTRY\_BM\`, \`COUNTRY\_BN\`,
\`COUNTRY\_BO\`, \`COUNTRY\_BQ\`, \`COUNTRY\_BR\`, \`COUNTRY\_BS\`, \`COUNTRY\_BT\`,
\`COUNTRY\_BV\`, \`COUNTRY\_BW\`, \`COUNTRY\_BY\`, \`COUNTRY\_BZ\`, \`COUNTRY\_CA\`,
\`COUNTRY\_CC\`, \`COUNTRY\_CD\`, \`COUNTRY\_CF\`, \`COUNTRY\_CG\`, \`COUNTRY\_CH\`,
\`COUNTRY\_CI\`, \`COUNTRY\_CK\`, \`COUNTRY\_CL\`, \`COUNTRY\_CM\`, \`COUNTRY\_CN\`,
\`COUNTRY\_CO\`, \`COUNTRY\_CR\`, \`COUNTRY\_CS\`, \`COUNTRY\_CU\`, \`COUNTRY\_CV\`,
\`COUNTRY\_CW\`, \`COUNTRY\_CX\`, \`COUNTRY\_CY\`, \`COUNTRY\_CZ\`, \`COUNTRY\_DE\`,
\`COUNTRY\_DJ\`, \`COUNTRY\_DK\`, \`COUNTRY\_DM\`, \`COUNTRY\_DO\`, \`COUNTRY\_DZ\`,
\`COUNTRY\_EC\`, \`COUNTRY\_EE\`, \`COUNTRY\_EG\`, \`COUNTRY\_EH\`, \`COUNTRY\_ER\`,
\`COUNTRY\_ES\`, \`COUNTRY\_ET\`, \`COUNTRY\_FI\`, \`COUNTRY\_FJ\`, \`COUNTRY\_FK\`,
\`COUNTRY\_FM\`, \`COUNTRY\_FO\`, \`COUNTRY\_FR\`, \`COUNTRY\_GA\`, \`COUNTRY\_GB\`,
\`COUNTRY\_GD\`, \`COUNTRY\_GE\`, \`COUNTRY\_GF\`, \`COUNTRY\_GG\`, \`COUNTRY\_GH\`,
\`COUNTRY\_GI\`, \`COUNTRY\_GL\`, \`COUNTRY\_GM\`, \`COUNTRY\_GN\`, \`COUNTRY\_GP\`,
\`COUNTRY\_GQ\`, \`COUNTRY\_GR\`, \`COUNTRY\_GS\`, \`COUNTRY\_GT\`, \`COUNTRY\_GU\`,
\`COUNTRY\_GW\`, \`COUNTRY\_GY\`, \`COUNTRY\_HK\`, \`COUNTRY\_HM\`, \`COUNTRY\_HN\`,
\`COUNTRY\_HR\`, \`COUNTRY\_HT\`, \`COUNTRY\_HU\`, \`COUNTRY\_ID\`, \`COUNTRY\_IE\`,
\`COUNTRY\_IL\`, \`COUNTRY\_IM\`, \`COUNTRY\_IN\`, \`COUNTRY\_IO\`, \`COUNTRY\_IQ\`,
\`COUNTRY\_IR\`, \`COUNTRY\_IS\`, \`COUNTRY\_IT\`, \`COUNTRY\_JE\`, \`COUNTRY\_JM\`,
\`COUNTRY\_JO\`, \`COUNTRY\_JP\`, \`COUNTRY\_KE\`, \`COUNTRY\_KG\`, \`COUNTRY\_KH\`,
\`COUNTRY\_KI\`, \`COUNTRY\_KM\`, \`COUNTRY\_KN\`, \`COUNTRY\_KP\`, \`COUNTRY\_KR\`,
\`COUNTRY\_KW\`, \`COUNTRY\_KY\`, \`COUNTRY\_KZ\`, \`COUNTRY\_LA\`, \`COUNTRY\_LB\`,
\`COUNTRY\_LC\`, \`COUNTRY\_LI\`, \`COUNTRY\_LK\`, \`COUNTRY\_LR\`, \`COUNTRY\_LS\`,
\`COUNTRY\_LT\`, \`COUNTRY\_LU\`, \`COUNTRY\_LV\`, \`COUNTRY\_LY\`, \`COUNTRY\_MA\`,
\`COUNTRY\_MC\`, \`COUNTRY\_MD\`, \`COUNTRY\_ME\`, \`COUNTRY\_MF\`, \`COUNTRY\_MG\`,
\`COUNTRY\_MH\`, \`COUNTRY\_MK\`, \`COUNTRY\_ML\`, \`COUNTRY\_MM\`, \`COUNTRY\_MN\`,
\`COUNTRY\_MO\`, \`COUNTRY\_MP\`, \`COUNTRY\_MQ\`, \`COUNTRY\_MR\`, \`COUNTRY\_MS\`,
\`COUNTRY\_MT\`, \`COUNTRY\_MU\`, \`COUNTRY\_MV\`, \`COUNTRY\_MW\`, \`COUNTRY\_MX\`,
\`COUNTRY\_MY\`, \`COUNTRY\_MZ\`, \`COUNTRY\_NA\`, \`COUNTRY\_NC\`, \`COUNTRY\_NE\`,
\`COUNTRY\_NF\`, \`COUNTRY\_NG\`, \`COUNTRY\_NI\`, \`COUNTRY\_NL\`, \`COUNTRY\_NO\`,
\`COUNTRY\_NP\`, \`COUNTRY\_NR\`, \`COUNTRY\_NU\`, \`COUNTRY\_NZ\`, \`COUNTRY\_OM\`,
\`COUNTRY\_PA\`, \`COUNTRY\_PE\`, \`COUNTRY\_PF\`, \`COUNTRY\_PG\`, \`COUNTRY\_PH\`,
\`COUNTRY\_PK\`, \`COUNTRY\_PL\`, \`COUNTRY\_PM\`, \`COUNTRY\_PN\`, \`COUNTRY\_PR\`,
\`COUNTRY\_PS\`, \`COUNTRY\_PT\`, \`COUNTRY\_PW\`, \`COUNTRY\_PY\`, \`COUNTRY\_QA\`,
\`COUNTRY\_RE\`, \`COUNTRY\_RO\`, \`COUNTRY\_RS\`, \`COUNTRY\_RU\`, \`COUNTRY\_RW\`,
\`COUNTRY\_SA\`, \`COUNTRY\_SB\`, \`COUNTRY\_SC\`, \`COUNTRY\_SD\`, \`COUNTRY\_SE\`,
\`COUNTRY\_SG\`, \`COUNTRY\_SH\`, \`COUNTRY\_SI\`, \`COUNTRY\_SJ\`, \`COUNTRY\_SK\`,
\`COUNTRY\_SL\`, \`COUNTRY\_SM\`, \`COUNTRY\_SN\`, \`COUNTRY\_SO\`, \`COUNTRY\_SR\`,
\`COUNTRY\_SS\`, \`COUNTRY\_ST\`, \`COUNTRY\_SV\`, \`COUNTRY\_SX\`, \`COUNTRY\_SY\`,
\`COUNTRY\_SZ\`, \`COUNTRY\_TC\`, \`COUNTRY\_TD\`, \`COUNTRY\_TF\`, \`COUNTRY\_TG\`,
\`COUNTRY\_TH\`, \`COUNTRY\_TJ\`, \`COUNTRY\_TK\`, \`COUNTRY\_TL\`, \`COUNTRY\_TM\`,
\`COUNTRY\_TN\`, \`COUNTRY\_TO\`, \`COUNTRY\_TR\`, \`COUNTRY\_TT\`, \`COUNTRY\_TV\`,
\`COUNTRY\_TW\`, \`COUNTRY\_TZ\`, \`COUNTRY\_UA\`, \`COUNTRY\_UG\`, \`COUNTRY\_UM\`,
\`COUNTRY\_US\`, \`COUNTRY\_UY\`, \`COUNTRY\_UZ\`, \`COUNTRY\_VA\`, \`COUNTRY\_VC\`,
\`COUNTRY\_VE\`, \`COUNTRY\_VG\`, \`COUNTRY\_VI\`, \`COUNTRY\_VN\`, \`COUNTRY\_VU\`,
\`COUNTRY\_WF\`, \`COUNTRY\_WS\`, \`COUNTRY\_XK\`, \`COUNTRY\_XT\`, \`COUNTRY\_YE\`,
\`COUNTRY\_YT\`, \`COUNTRY\_ZA\`, \`COUNTRY\_ZM\`, \`COUNTRY\_ZW\`. Defaults to \`COUNTRY\_NONE\`.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.items.enum.not_in": "[0]",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.items.enum.not_in": "[0]",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [default_action_allow](data-sources--service_policy--reference--group-001.md#canonical-3112033300312200-1133221301333110-1031111213320001-1231202110320203-0230001331012002-0000121113020102-2323123112003130-3222311321233002): complete subsection reference.

- [default_action_deny](data-sources--service_policy--reference--group-001.md#canonical-3100231301121321-2302333102303320-1123012022001032-0300112200033312-0012111311103332-2102210033110333-0303232102013103-2132013220231121): complete subsection reference.

- [default_action_next_policy](data-sources--service_policy--reference--group-001.md#canonical-3321002310033200-2102112201330133-0323133010230133-2330223033100332-2131121122102132-2130113211111112-1302100121323202-2332312133133111): complete subsection reference.

- [ip_prefix_set](data-sources--service_policy--reference--group-001.md#canonical-0033133330311211-2231012230313202-0100221313200200-0023202013303210-3312222002033023-0313321120323122-3011233010030100-3233033301200001): complete subsection reference.

- [prefix_list](data-sources--service_policy--reference--group-001.md#canonical-2222000312112021-0020321133320222-2020121021111001-1330320020131232-2032000031320323-3012330110202112-2000030102313310-3220312230110121): complete subsection reference.

<a id="canonical-2011001300111011-3222122130222330-2110323331320203-0200310020130120-2031011133012120-0030102213221301-2230010111321111-0330112130013301"></a>

<a id="canonical-1113221113132311-0013001322230011-2311200303310130-0233100013103103-1322022330022120-1222233121310200-3011132033222121-2210032330030111"></a>

#### `deny_list.tls_fingerprint_classes` property

Type: `["list", "string"]`. Computed.

\[Enum:
TLS\_FINGERPRINT\_NONE|ANY\_MALICIOUS\_FINGERPRINT|ADWARE|ADWIND|DRIDEX|GOOTKIT|GOZI|JBIFROST|QUAKBOT|RANSOMWARE|TROLDESH|TOFSEE|TORRENTLOCKER|TRICKBOT\]
List of known classes of TLS fingerprints to match the input TLS JA3 fingerprint against. Possible
values are \`TLS\_FINGERPRINT\_NONE\`, \`ANY\_MALICIOUS\_FINGERPRINT\`, \`ADWARE\`, \`ADWIND\`,
\`DRIDEX\`, \`GOOTKIT\`, \`GOZI\`, \`JBIFROST\`, \`QUAKBOT\`, \`RANSOMWARE\`, \`TROLDESH\`,
\`TOFSEE\`, \`TORRENTLOCKER\`, \`TRICKBOT\`. Defaults to \`TLS\_FINGERPRINT\_NONE\`.

Additional upstream details:

A list of known classes of TLS fingerprints to match the input TLS JA3 fingerprint against.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-2302323023100211-1322210121030130-2013000310003202-3013222010301123-1220331331121130-2120332030312110-1111323102030231-2021012013301120"></a>

<a id="canonical-2201001312001231-1123221030003312-1111033312210320-1111232112312110-2310120220310121-0111103002231130-0121100000102120-2002331232033213"></a>

#### `deny_list.tls_fingerprint_values` property

Type: `["list", "string"]`. Computed.

A list of exact TLS JA3 fingerprints to match the input TLS JA3 fingerprint against.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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
    "ves.io.schema.rules.repeated.items.string.len": "32",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.len": "32",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0011132331310313-1312032302323313-3130122322230323-0312311112332011-0301221013302132-2232030313211110-3320321122333033-2011332130311213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `deny_list.asn_list` properties

Breadcrumbs:

- [xcsh_service_policy](../data-sources/service_policy.md#canonical-1112022302333333-1130100232232301-3023212310221022-3103031003221200-2030020130230132-1201210321332200-2212311003030121-1321310310032303)
- [Property reference](data-sources--service_policy--reference--group-001.md#canonical-3312102302111103-0221020013211300-2120323123300113-1101123102331230-1032322121222101-1332231300102132-0233300331323303-0031323122032013)
- [deny_list](data-sources--service_policy--reference--group-001.md#canonical-1133210330221300-2302313012231010-0300310022202220-2111123320332210-2101000320030120-3220020323311021-3132220133231132-0021330023303013)
- deny_list.asn_list

<a id="canonical-3332222101231202-1210322010202011-0013010022332202-2202132202003130-1121231101200103-1113013103100003-3220323330201113-2223221100012100"></a>

Type: `"single"`. Computed.

An unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny
lists for use in network policy or service policy. It can be used to create the allow list only for
DNS Load Balancer.

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

<a id="canonical-2012231300321313-1210323330102031-1131313113303013-0030102301211330-3021333021032211-2221101011223323-3023321130032123-0023322331123323"></a>

### Direct properties for `deny_list.asn_list`

<a id="canonical-2101133123131110-1302322012322202-2012210132012210-2300001321102332-0032120222022221-3113212101021210-3313330223232121-3102232211332200"></a>

#### `deny_list.asn_list.as_numbers` property

Type: `["list", "number"]`. Computed.

An unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny
lists for use in network policy or service policy. It can be used to create the allow list only for
DNS Load Balancer.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 16,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 16,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-06T12:36:10+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3330311320223113-1300223221102113-2100032022322001-3132023001102020-2231000000200300-1032221303212201-3031130220103020-2332321310301200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `deny_list.asn_set` properties

Breadcrumbs:

- [xcsh_service_policy](../data-sources/service_policy.md#canonical-1112022302333333-1130100232232301-3023212310221022-3103031003221200-2030020130230132-1201210321332200-2212311003030121-1321310310032303)
- [Property reference](data-sources--service_policy--reference--group-001.md#canonical-3312102302111103-0221020013211300-2120323123300113-1101123102331230-1032322121222101-1332231300102132-0233300331323303-0031323122032013)
- [deny_list](data-sources--service_policy--reference--group-001.md#canonical-1133210330221300-2302313012231010-0300310022202220-2111123320332210-2101000320030120-3220020323311021-3132220133231132-0021330023303013)
- deny_list.asn_set

<a id="canonical-1223201123120023-1121020133310021-3022130023232101-1232230023010223-2223201020231013-2323332121000001-0032310302313233-3323010302220232"></a>

Type: `"list"`. Computed.

Addresses that belong to the ASNs in the given bgp\_asn\_set The ASN is obtained by performing a
lookup for the source IPv4 Address in a GeoIP DB.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-1321221032221030-2202033122110001-0323132333011023-3333212231213100-3333130100300302-3321321200001013-1011202323111120-3231322200211313"></a>

### Direct properties for `deny_list.asn_set`

<a id="canonical-3133013100213233-1123003200103120-0132030122013220-0231202323013311-2322330313313103-2120100231012310-0112312102212311-2201222103200013"></a>

#### `deny_list.asn_set.name` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-0221330232123000-3213121022302233-3010022230301220-2123231222031102-0201300132013111-1310110322111231-0312201303222011-3100133111233313"></a>

<a id="canonical-0100011131011103-3231320211303001-0101110313023322-0130022003102222-2300223131322203-2211013200010011-1312021110330200-2021033313112103"></a>

#### `deny_list.asn_set.namespace` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-3113232323331121-3002210020002123-2203021210132330-1010102323131133-0320033013222223-3033030203133010-1322010302233233-2332311220021011"></a>

<a id="canonical-2001112320122211-3320030300003200-3133011111313230-1003302313022331-0123321310100320-3211222131101033-3333123012220132-3331220012302330"></a>

#### `deny_list.asn_set.tenant` property

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-3112033300312200-1133221301333110-1031111213320001-1231202110320203-0230001331012002-0000121113020102-2323123112003130-3222311321233002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `deny_list.default_action_allow` properties

Breadcrumbs:

- [xcsh_service_policy](../data-sources/service_policy.md#canonical-1112022302333333-1130100232232301-3023212310221022-3103031003221200-2030020130230132-1201210321332200-2212311003030121-1321310310032303)
- [Property reference](data-sources--service_policy--reference--group-001.md#canonical-3312102302111103-0221020013211300-2120323123300113-1101123102331230-1032322121222101-1332231300102132-0233300331323303-0031323122032013)
- [deny_list](data-sources--service_policy--reference--group-001.md#canonical-1133210330221300-2302313012231010-0300310022202220-2111123320332210-2101000320030120-3220020323311021-3132220133231132-0021330023303013)
- deny_list.default_action_allow

<a id="canonical-0022323322233322-2300133310312232-0030310001120013-3212132010331201-1223030111202311-3133201312001312-1113212221202113-0300221331131001"></a>

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

<a id="canonical-3100231301121321-2302333102303320-1123012022001032-0300112200033312-0012111311103332-2102210033110333-0303232102013103-2132013220231121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `deny_list.default_action_deny` properties

Breadcrumbs:

- [xcsh_service_policy](../data-sources/service_policy.md#canonical-1112022302333333-1130100232232301-3023212310221022-3103031003221200-2030020130230132-1201210321332200-2212311003030121-1321310310032303)
- [Property reference](data-sources--service_policy--reference--group-001.md#canonical-3312102302111103-0221020013211300-2120323123300113-1101123102331230-1032322121222101-1332231300102132-0233300331323303-0031323122032013)
- [deny_list](data-sources--service_policy--reference--group-001.md#canonical-1133210330221300-2302313012231010-0300310022202220-2111123320332210-2101000320030120-3220020323311021-3132220133231132-0021330023303013)
- deny_list.default_action_deny

<a id="canonical-1301120211113332-2311313101103202-1001112231022311-2103330230130230-2210313210133221-0232132212300120-3101021232210033-0233331212311231"></a>

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

<a id="canonical-3321002310033200-2102112201330133-0323133010230133-2330223033100332-2131121122102132-2130113211111112-1302100121323202-2332312133133111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `deny_list.default_action_next_policy` properties

Breadcrumbs:

- [xcsh_service_policy](../data-sources/service_policy.md#canonical-1112022302333333-1130100232232301-3023212310221022-3103031003221200-2030020130230132-1201210321332200-2212311003030121-1321310310032303)
- [Property reference](data-sources--service_policy--reference--group-001.md#canonical-3312102302111103-0221020013211300-2120323123300113-1101123102331230-1032322121222101-1332231300102132-0233300331323303-0031323122032013)
- [deny_list](data-sources--service_policy--reference--group-001.md#canonical-1133210330221300-2302313012231010-0300310022202220-2111123320332210-2101000320030120-3220020323311021-3132220133231132-0021330023303013)
- deny_list.default_action_next_policy

<a id="canonical-2321132010020011-2331132200310002-2231133322202210-2012232120223120-1310221231120300-2321110000333202-3032023320312033-3210120233302010"></a>

Type: `["object", {}]`. Computed.

Policy configuration for this feature.

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

<a id="canonical-0033133330311211-2231012230313202-0100221313200200-0023202013303210-3312222002033023-0313321120323122-3011233010030100-3233033301200001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `deny_list.ip_prefix_set` properties

Breadcrumbs:

- [xcsh_service_policy](../data-sources/service_policy.md#canonical-1112022302333333-1130100232232301-3023212310221022-3103031003221200-2030020130230132-1201210321332200-2212311003030121-1321310310032303)
- [Property reference](data-sources--service_policy--reference--group-001.md#canonical-3312102302111103-0221020013211300-2120323123300113-1101123102331230-1032322121222101-1332231300102132-0233300331323303-0031323122032013)
- [deny_list](data-sources--service_policy--reference--group-001.md#canonical-1133210330221300-2302313012231010-0300310022202220-2111123320332210-2101000320030120-3220020323311021-3132220133231132-0021330023303013)
- deny_list.ip_prefix_set

<a id="canonical-3102223301301201-2113120232113013-0333021111133010-0210331220212212-2312012220323302-2321230312302303-1001130321111230-2301320331213212"></a>

Type: `"list"`. Computed.

Addresses that are covered by the prefixes in the given ip\_prefix\_set.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-2311112321112120-2111002022222020-0323121201030320-0322102030113023-2112231012120013-2200032031132230-2222222130112220-1020201021110222"></a>

### Direct properties for `deny_list.ip_prefix_set`

<a id="canonical-2210122210300313-2002032131301223-3002133333020310-0322022312111321-3233200013031303-0202131212312020-0213203200110131-1210222212003001"></a>

#### `deny_list.ip_prefix_set.name` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-3020322220330301-1012122310200221-3001133123110323-1013022100002322-0013000123130021-1101001320302102-1323033102200201-3331302231202311"></a>

<a id="canonical-1311103333223301-0220103122100333-2213123031322110-2230011101303323-2321130112212323-2113110103023112-1230102212130013-1302203011110200"></a>

#### `deny_list.ip_prefix_set.namespace` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-2002313233122232-0221101110210101-0321301020311333-0031332013032230-1031301011201121-3232313331100132-2321000330101113-1331010213131333"></a>

<a id="canonical-2130100111330033-0331002030010233-0100332001321303-2111133231120212-2213333201113213-1300121022030312-1123320123320223-1333001330001010"></a>

#### `deny_list.ip_prefix_set.tenant` property

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-2222000312112021-0020321133320222-2020121021111001-1330320020131232-2032000031320323-3012330110202112-2000030102313310-3220312230110121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `deny_list.prefix_list` properties

Breadcrumbs:

- [xcsh_service_policy](../data-sources/service_policy.md#canonical-1112022302333333-1130100232232301-3023212310221022-3103031003221200-2030020130230132-1201210321332200-2212311003030121-1321310310032303)
- [Property reference](data-sources--service_policy--reference--group-001.md#canonical-3312102302111103-0221020013211300-2120323123300113-1101123102331230-1032322121222101-1332231300102132-0233300331323303-0031323122032013)
- [deny_list](data-sources--service_policy--reference--group-001.md#canonical-1133210330221300-2302313012231010-0300310022202220-2111123320332210-2101000320030120-3220020323311021-3132220133231132-0021330023303013)
- deny_list.prefix_list

<a id="canonical-0110002022022111-2211102313122002-1132000302031101-2223133030301111-3002101300032113-3010033000121030-1331110001232021-2213100200200012"></a>

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

<a id="canonical-1011310110102300-2021213021000332-1233220110323222-3123103023202302-0200303102301230-2203120020031022-3131022320020133-1233302223023330"></a>

### Direct properties for `deny_list.prefix_list`

<a id="canonical-3013121103313210-0202220302200012-3031023202212211-0010011131201322-2101113030100213-3000131331113010-3033210223202333-3112030011001321"></a>

#### `deny_list.prefix_list.prefixes` property

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-0200021302120322-3330103111223310-0110033233312323-2211001021302330-2000313233321013-1330302203133111-3101023222321313-3302013113202033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list` properties

Breadcrumbs:

- [xcsh_service_policy](../data-sources/service_policy.md#canonical-1112022302333333-1130100232232301-3023212310221022-3103031003221200-2030020130230132-1201210321332200-2212311003030121-1321310310032303)
- [Property reference](data-sources--service_policy--reference--group-001.md#canonical-3312102302111103-0221020013211300-2120323123300113-1101123102331230-1032322121222101-1332231300102132-0233300331323303-0031323122032013)
- rule_list

<a id="canonical-1323033222302302-3111221111132320-3301101131323321-3012033101003230-1210201102233222-0102310331101011-1003003201123022-2202030210300232"></a>

Type: `"single"`. Computed.

Ordered service-policy rules for non-geographic predicates and actions. Do not use country\_list for
a geo-only rule here: the platform adds match-all any\_ip and any\_asn selectors on readback, so the
rule can match all traffic. Use deny\_list or allow\_list with country\_list for geographic source
matching.

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

<a id="canonical-2203330202021131-0332010231031330-0103023230022133-2133211323122203-0232312130101333-0033301302113312-3023210113000103-1132313020223020"></a>

### Direct properties for `rule_list`

- [rules](data-sources--service_policy--reference--group-001.md#canonical-2332210012031020-3130121121203233-2133312023012212-1300213020002321-0210122201111201-2313301011130132-1102100221012102-3320333010302121): complete subsection reference.

<a id="canonical-2332210012031020-3130121121203233-2133312023012212-1300213020002321-0210122201111201-2313301011130132-1102100221012102-3320333010302121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules` properties

Breadcrumbs:

- [xcsh_service_policy](../data-sources/service_policy.md#canonical-1112022302333333-1130100232232301-3023212310221022-3103031003221200-2030020130230132-1201210321332200-2212311003030121-1321310310032303)
- [Property reference](data-sources--service_policy--reference--group-001.md#canonical-3312102302111103-0221020013211300-2120323123300113-1101123102331230-1032322121222101-1332231300102132-0233300331323303-0031323122032013)
- [rule_list](data-sources--service_policy--reference--group-001.md#canonical-0200021302120322-3330103111223310-0110033233312323-2211001021302330-2000313233321013-1330302203133111-3101023222321313-3302013113202033)
- rule_list.rules

<a id="canonical-2101123033112122-2302210303331323-0113210121020123-3130021131003022-3210120301201231-1310022213100132-3100013031031120-0123232230222033"></a>

Type: `"list"`. Computed.

Define the list of rules (with an order) that should be evaluated by this service policy. Rules are
evaluated from top to bottom in the list.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
    },
    "minItems": 0,
    "uniqueItems": false
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "256",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "256",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  }
}
```

<a id="canonical-2023030022232023-3033100233003232-3032303211332322-3231010311003101-3000223323232300-3221100111030000-0111012131301112-0132020301133312"></a>

### Direct properties for `rule_list.rules`

- [metadata](data-sources--service_policy--reference--group-001.md#canonical-0312100030223123-0101030013231223-2112033202221103-0033223300021311-0112310011000120-2200101332312223-0221102111131111-0211010031222231): complete subsection reference.

- [spec](data-sources--service_policy--reference--group-001.md#canonical-1313133011110120-1202202112012022-0312300012021010-2230300132201133-1123333233212022-2100210133123323-2030333300021222-2110310211021210): complete subsection reference.

<a id="canonical-0312100030223123-0101030013231223-2112033202221103-0033223300021311-0112310011000120-2200101332312223-0221102111131111-0211010031222231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules.metadata` properties

Breadcrumbs:

- [xcsh_service_policy](../data-sources/service_policy.md#canonical-1112022302333333-1130100232232301-3023212310221022-3103031003221200-2030020130230132-1201210321332200-2212311003030121-1321310310032303)
- [Property reference](data-sources--service_policy--reference--group-001.md#canonical-3312102302111103-0221020013211300-2120323123300113-1101123102331230-1032322121222101-1332231300102132-0233300331323303-0031323122032013)
- [rule_list](data-sources--service_policy--reference--group-001.md#canonical-0200021302120322-3330103111223310-0110033233312323-2211001021302330-2000313233321013-1330302203133111-3101023222321313-3302013113202033)
- [rule_list.rules](data-sources--service_policy--reference--group-001.md#canonical-2332210012031020-3130121121203233-2133312023012212-1300213020002321-0210122201111201-2313301011130132-1102100221012102-3320333010302121)
- rule_list.rules.metadata

<a id="canonical-2132122221011103-1233111033030100-3231103313200310-0200202333111121-1011321030232023-3132303212002110-0111033113000111-2013212221213320"></a>

Type: `"single"`. Computed.

MessageMetaType is metadata (common attributes) of a message that only certain messages have. This
information is propagated to the metadata of a child object that gets created from the containing
message during view processing. The information in this type can be specified by user during create
and replace APIs.

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

<a id="canonical-0121321102101002-2312031302222311-2312312020212300-0003230202221320-3010101211111033-2001332312001120-2102223323220311-2101131201003331"></a>

### Direct properties for `rule_list.rules.metadata`

<a id="canonical-0212221232311013-3022222221012223-1023030033300300-3020230111333121-1312100333310020-3100233212123211-3121332300122110-1203203222301111"></a>

#### `rule_list.rules.metadata.description_spec` property

Type: `"string"`. Computed.

Description. Human readable description.

<a id="canonical-1000132132131211-3110100001311320-2302200310113010-2112212130123330-3230221122002300-2131322303012032-0103232131233232-0233231121323210"></a>

<a id="canonical-2030132103203330-0320311211332002-3020322322302232-3133102002121003-0332002002102000-1301131320313302-1021333133030322-3330011202013103"></a>

#### `rule_list.rules.metadata.name` property

Type: `"string"`. Computed.

This is the name of the message. The value of name has to follow DNS-1035 format.

Receipt-pinned upstream constraints:

```json
{
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
    "maxLength": 63,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-06T12:36:10+00:00"
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
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.ves_object_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.ves_object_name": "true"
  }
}
```

<a id="canonical-1313133011110120-1202202112012022-0312300012021010-2230300132201133-1123333233212022-2100210133123323-2030333300021222-2110310211021210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules.spec` properties

Breadcrumbs:

- [xcsh_service_policy](../data-sources/service_policy.md#canonical-1112022302333333-1130100232232301-3023212310221022-3103031003221200-2030020130230132-1201210321332200-2212311003030121-1321310310032303)
- [Property reference](data-sources--service_policy--reference--group-001.md#canonical-3312102302111103-0221020013211300-2120323123300113-1101123102331230-1032322121222101-1332231300102132-0233300331323303-0031323122032013)
- [rule_list](data-sources--service_policy--reference--group-001.md#canonical-0200021302120322-3330103111223310-0110033233312323-2211001021302330-2000313233321013-1330302203133111-3101023222321313-3302013113202033)
- [rule_list.rules](data-sources--service_policy--reference--group-001.md#canonical-2332210012031020-3130121121203233-2133312023012212-1300213020002321-0210122201111201-2313301011130132-1102100221012102-3320333010302121)
- rule_list.rules.spec

<a id="canonical-1133013120211033-0320113201200113-2030303010023202-2220311210313200-0022221110321003-0302002312313013-0323310011323211-0101213132323023"></a>

Type: `"single"`. Computed.

Shape of service\_policy\_rule in the storage backend.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-asn_choice": "[\"any_asn\",\"asn_list\",\"asn_matcher\"]",
  "x-ves-oneof-field-client_choice": "[\"any_client\",\"client_name\",\"client_name_matcher\",\"client_selector\",\"ip_threat_category_list\"]",
  "x-ves-oneof-field-dst_asn_choice": "[]",
  "x-ves-oneof-field-dst_ip_choice": "[]",
  "x-ves-oneof-field-ip_choice": "[\"any_ip\",\"ip_matcher\",\"ip_prefix_list\"]",
  "x-ves-oneof-field-tls_fingerprint_choice": "[\"ja4_tls_fingerprint\",\"tls_fingerprint_matcher\"]"
}
```

<a id="canonical-2120211230020211-2312002011010111-0122132133001233-1213332102100120-3213201110003320-2211202013322121-1030231330213331-0032030332110012"></a>

### Direct properties for `rule_list.rules.spec`

<a id="canonical-0030310321031303-1210300001000323-0330321121202332-0313300002301030-0102113303323012-1221120000322010-0021322223222103-1201002013030122"></a>

#### `rule_list.rules.spec.action` property

Type: `"string"`. Computed.

\[Enum: DENY|ALLOW|NEXT\_POLICY\] The rule action determines the disposition of the input request
API. If a policy matches a rule with an ALLOW action, the processing of the request proceeds
forward. If it matches a rule with a DENY action, the processing of the request is terminated and an
appropriate message/code returned to.. Possible values are \`DENY\`, \`ALLOW\`, \`NEXT\_POLICY\`.
Defaults to \`DENY\`.

Additional upstream details:

The rule action determines the disposition of the input request API. If it matches a rule with a
DENY action, the processing of the request is terminated and an appropriate message/code returned to
the originator. If it matches a rule with a NEXT\_POLICY\_SET action, evaluation of the current
policy set terminates and evaluation of the next policy set in the chain begins.

&#8203;- DENY: DENY

Deny the request. &#8203;- ALLOW: ALLOW

Allow the request to proceed. &#8203;- NEXT\_POLICY\_SET: NEXT\_POLICY\_SET

Terminate evaluation of the current policy set and begin evaluating the next policy set in the
chain. Note that the evaluation of any remaining policies in the current policy set is skipped.
&#8203;- NEXT\_POLICY: NEXT\_POLICY

Terminate evaluation of the current policy and begin evaluating the next policy in the policy set.
Note that the evaluation of any remaining rules in the current policy is skipped. &#8203;-
LAST\_POLICY: LAST\_POLICY

Terminate evaluation of the current policy and begin evaluating the last policy in the policy set.
Note that the evaluation of any remaining rules in the current policy is skipped. &#8203;-
GOTO\_POLICY: GOTO\_POLICY

Terminate evaluation of the current policy and begin evaluating a specific policy in the policy set.
The policy is specified using the goto\_policy field in the rule and must be after the current
policy in the policy set.

Receipt-pinned upstream constraints:

```json
{
  "default": "DENY",
  "enum": [
    "DENY",
    "ALLOW",
    "NEXT_POLICY"
  ],
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  }
}
```

- [any_asn](data-sources--service_policy--reference--group-001.md#canonical-3111231132021020-0103003300320300-1233210021330021-0321230021103103-0011233333002210-2213133330230222-1121132210022102-3110030011232022): complete subsection reference.

- [any_client](data-sources--service_policy--reference--group-001.md#canonical-1121101031021212-2012012123031320-0013123212100031-2101010322122020-1001232020303133-0230331001201001-2133130000213113-3131210103230101): complete subsection reference.

- [any_ip](data-sources--service_policy--reference--group-001.md#canonical-3322011213231210-0211030030100003-2120231130000333-3110210120020010-2330202211213313-2323310212211220-2020333221321321-0122230020030221): complete subsection reference.

- [api_group_matcher](data-sources--service_policy--reference--group-001.md#canonical-0000330011302112-1212312313102203-0221331010012030-1233110131113232-0132032001100133-1233020003212002-1232220310301031-3113012300230333): complete subsection reference.

- [arg_matchers](data-sources--service_policy--reference--group-001.md#canonical-0322132120012130-1302303202022130-3333202213132012-3332222012222013-1122230102030120-3303120100313000-3131231103000201-3020233310033201): complete subsection reference.

- [asn_list](data-sources--service_policy--reference--group-001.md#canonical-0003110021000230-3123203102130033-2313031311212300-3212013310032302-1030231200103320-2032330222011323-3320101221001100-1101302312333033): complete subsection reference.

- [asn_matcher](data-sources--service_policy--reference--group-001.md#canonical-3003212230320021-1130112110220111-3212200301310132-0123111321310113-3213001310131320-2031031330231203-1132023221020001-3002123211023021): complete subsection reference.

- [body_matcher](data-sources--service_policy--reference--group-001.md#canonical-3120003330322333-0013033122113213-1031303002011013-0121111101233222-2001211210033333-2313220132320101-0011222020013012-0022323101113322): complete subsection reference.

- [bot_action](data-sources--service_policy--reference--group-002.md#canonical-1312131031200012-0323222232212311-0110213200221122-3131110311121211-1002113122132211-1032112220111030-3100011123222312-2232213111221012): complete subsection reference.

<a id="canonical-0122122233031311-1200303031303233-2001113120113211-2021121031101002-0003132111113300-0011211221021110-0310231003032211-3202122120001313"></a>

<a id="canonical-3222230031330032-2200333302122011-1111333022100331-2102203320000300-0020202202020230-2323023303210020-1033121111302030-1221212110323200"></a>

#### `rule_list.rules.spec.client_name` property

Type: `"string"`. Computed.

Exclusive with \[any\_client client\_name\_matcher client\_selector ip\_threat\_category\_list\] The
expected name of the client invoking the request API. The predicate evaluates to true if any of the
actual names is the same as the expected client name.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 256
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-06T12:36:10+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256"
  }
}
```

- [client_name_matcher](data-sources--service_policy--reference--group-002.md#canonical-0121131233020103-2231131133023313-0021123010201112-3322100301021233-0200303032300002-1100203321101230-3201303322100031-3221133200303303): complete subsection reference.

- [client_selector](data-sources--service_policy--reference--group-002.md#canonical-1230321030333321-2133202001231022-0210120112103312-3130210331103021-0123133331212113-1310313101233211-2300330310202031-0032002312032330): complete subsection reference.

- [cookie_matchers](data-sources--service_policy--reference--group-002.md#canonical-2111300231202303-2030301101223223-3032333120002222-2000302311201303-3012020113230021-1223300021220201-3302002113123122-3213110013333130): complete subsection reference.

- [domain_matcher](data-sources--service_policy--reference--group-002.md#canonical-3322203033331023-2233220302031220-2301231103021132-2020113132210223-2001020313023320-0220311122312313-0121320222003303-2131021003133213): complete subsection reference.

<a id="canonical-2221011121222100-1023211233032112-1321023312012332-0320001222332330-3020230011132223-1202111202123103-3121210033103020-3031010220121000"></a>

<a id="canonical-1232022321320030-1333220031120112-1230000200211200-0302332220023013-3300300332321121-0203000021012111-1033302113330020-3021211100211310"></a>

#### `rule_list.rules.spec.expiration_timestamp` property

Type: `"string"`. Computed.

Specifies expiration\_timestamp the RFC 3339 format timestamp at which the containing rule is
considered to be logically expired. The rule continues to exist in the configuration but is not
applied anymore.

Additional upstream details:

The expiration\_timestamp is the RFC 3339 format timestamp at which the containing rule is
considered to be logically expired.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "format": "date-time",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

- [headers](data-sources--service_policy--reference--group-002.md#canonical-2022122312102320-3021230232003303-1301003301023011-2212121031302321-3123311303200100-0020331201223301-3103232233311212-3033201132111122): complete subsection reference.

- [http_method](data-sources--service_policy--reference--group-002.md#canonical-0223130311323131-0303100033332230-3203111320203213-3102331312020011-3013112003331321-2122323011003121-2313130233330233-0103130010303202): complete subsection reference.

- [ip_matcher](data-sources--service_policy--reference--group-002.md#canonical-3001311313101200-0111313102213231-0330101313233022-0210001223230123-2232331303200003-2213132112112331-2233022233020222-3112013102100001): complete subsection reference.

- [ip_prefix_list](data-sources--service_policy--reference--group-002.md#canonical-1133021322231212-2103223100112132-2212113321333203-0113233322020010-0003213320223311-3112100133213021-3323011012332312-1123001232301321): complete subsection reference.

- [ip_threat_category_list](data-sources--service_policy--reference--group-002.md#canonical-0213001221332012-0333230020030103-1200211310131211-0231323331103203-2310310332330330-3313312022312012-1302033120012331-0021212310222310): complete subsection reference.

- [ja4_tls_fingerprint](data-sources--service_policy--reference--group-002.md#canonical-0020023013210332-1201130101003103-1003112321233212-2120210311320132-0321330033001332-2103123112111112-1013321012220033-1032311101122303): complete subsection reference.

- [jwt_claims](data-sources--service_policy--reference--group-002.md#canonical-0302332200230030-0110111001203102-3332033203213321-1013111323001230-3121103113222113-0233230330232331-3022210003302100-3123110320330031): complete subsection reference.

- [label_matcher](data-sources--service_policy--reference--group-002.md#canonical-1221302121333112-2003223202021213-1330010203000001-3101131322012130-0100030033121220-2120330230202021-3310121000131102-1311103021210331): complete subsection reference.

<a id="canonical-3100030233213202-2131211321003130-0313102000123103-3310210200223132-3313233322332131-2201131301113002-2300200102330321-1132220122121223"></a>

<a id="canonical-2030323111132000-2110212011332312-2221012031113013-1113133320203212-1303222110233223-2312200013230031-1101300123222130-0123330103212120"></a>

#### `rule_list.rules.spec.log_rule_evaluation` property

Type: `"bool"`. Computed.

Log the rule match details along with the request and continue to evaluate rules in the sequence.

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

- [mum_action](data-sources--service_policy--reference--group-002.md#canonical-2122103222123111-0310123221120033-3122230001202021-3231333033332103-3101331023112023-0203202032222223-1123223012003012-1103223132011322): complete subsection reference.

- [path](data-sources--service_policy--reference--group-002.md#canonical-0320210310200110-3200233233321020-3202312003012032-3221022312030021-3213113211200131-0321203201123201-1301333233103210-0313130122100133): complete subsection reference.

- [port_matcher](data-sources--service_policy--reference--group-002.md#canonical-0002112321213331-2301331003031102-2101130012331010-0010130231301212-0111003003223033-3131022232020221-3200101100331330-2033223020101132): complete subsection reference.

- [query_params](data-sources--service_policy--reference--group-002.md#canonical-2203212033320122-1121331010023001-1301233321020313-1100002302220320-3313310022002320-0220102300031101-0221111311221203-2002111321113003): complete subsection reference.

- [request_constraints](data-sources--service_policy--reference--group-002.md#canonical-2110021331103130-3023312121033312-2231302230120211-1103201133030222-0221022130001011-3002013103120132-1321211031031000-0320022212322013): complete subsection reference.

- [segment_policy](data-sources--service_policy--reference--group-002.md#canonical-0121322232330011-2001200302133213-0020123132212012-0113020311112312-0201321011303010-1302103121002333-1323212012233310-1210231031031002): complete subsection reference.

- [tls_fingerprint_matcher](data-sources--service_policy--reference--group-002.md#canonical-1333003020033103-2303031332133313-1101333013023130-2313101003001221-0111232133300132-2011202312030000-1012231232313131-2333220102021302): complete subsection reference.

- [user_identity_matcher](data-sources--service_policy--reference--group-002.md#canonical-3302232320022303-1301323100000232-2222223202031122-2132301121130100-3212113133030300-0223231120230100-3101210030303312-3003333002303311): complete subsection reference.

- [waf_action](data-sources--service_policy--reference--group-002.md#canonical-1321030020301232-0323133310020002-1031312201202122-2313331303112210-3101113223232332-0322221031012323-3201210323110133-2121002301220100): complete subsection reference.

<a id="canonical-3111231132021020-0103003300320300-1233210021330021-0321230021103103-0011233333002210-2213133330230222-1121132210022102-3110030011232022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules.spec.any_asn` properties

Breadcrumbs:

- [xcsh_service_policy](../data-sources/service_policy.md#canonical-1112022302333333-1130100232232301-3023212310221022-3103031003221200-2030020130230132-1201210321332200-2212311003030121-1321310310032303)
- [Property reference](data-sources--service_policy--reference--group-001.md#canonical-3312102302111103-0221020013211300-2120323123300113-1101123102331230-1032322121222101-1332231300102132-0233300331323303-0031323122032013)
- [rule_list](data-sources--service_policy--reference--group-001.md#canonical-0200021302120322-3330103111223310-0110033233312323-2211001021302330-2000313233321013-1330302203133111-3101023222321313-3302013113202033)
- [rule_list.rules](data-sources--service_policy--reference--group-001.md#canonical-2332210012031020-3130121121203233-2133312023012212-1300213020002321-0210122201111201-2313301011130132-1102100221012102-3320333010302121)
- [rule_list.rules.spec](data-sources--service_policy--reference--group-001.md#canonical-1313133011110120-1202202112012022-0312300012021010-2230300132201133-1123333233212022-2100210133123323-2030333300021222-2110310211021210)
- rule_list.rules.spec.any_asn

<a id="canonical-1020211012322112-2130223103313230-1221102230302302-3121003131023133-2231320201222110-1300320231013220-3313322103033120-2132132220211101"></a>

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

<a id="canonical-1121101031021212-2012012123031320-0013123212100031-2101010322122020-1001232020303133-0230331001201001-2133130000213113-3131210103230101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules.spec.any_client` properties

Breadcrumbs:

- [xcsh_service_policy](../data-sources/service_policy.md#canonical-1112022302333333-1130100232232301-3023212310221022-3103031003221200-2030020130230132-1201210321332200-2212311003030121-1321310310032303)
- [Property reference](data-sources--service_policy--reference--group-001.md#canonical-3312102302111103-0221020013211300-2120323123300113-1101123102331230-1032322121222101-1332231300102132-0233300331323303-0031323122032013)
- [rule_list](data-sources--service_policy--reference--group-001.md#canonical-0200021302120322-3330103111223310-0110033233312323-2211001021302330-2000313233321013-1330302203133111-3101023222321313-3302013113202033)
- [rule_list.rules](data-sources--service_policy--reference--group-001.md#canonical-2332210012031020-3130121121203233-2133312023012212-1300213020002321-0210122201111201-2313301011130132-1102100221012102-3320333010302121)
- [rule_list.rules.spec](data-sources--service_policy--reference--group-001.md#canonical-1313133011110120-1202202112012022-0312300012021010-2230300132201133-1123333233212022-2100210133123323-2030333300021222-2110310211021210)
- rule_list.rules.spec.any_client

<a id="canonical-1300330003023333-1333111332220312-2310222203120032-1300110232122131-3203121320022231-3332023213231211-0220220220232233-3210030133023220"></a>

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

<a id="canonical-3322011213231210-0211030030100003-2120231130000333-3110210120020010-2330202211213313-2323310212211220-2020333221321321-0122230020030221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules.spec.any_ip` properties

Breadcrumbs:

- [xcsh_service_policy](../data-sources/service_policy.md#canonical-1112022302333333-1130100232232301-3023212310221022-3103031003221200-2030020130230132-1201210321332200-2212311003030121-1321310310032303)
- [Property reference](data-sources--service_policy--reference--group-001.md#canonical-3312102302111103-0221020013211300-2120323123300113-1101123102331230-1032322121222101-1332231300102132-0233300331323303-0031323122032013)
- [rule_list](data-sources--service_policy--reference--group-001.md#canonical-0200021302120322-3330103111223310-0110033233312323-2211001021302330-2000313233321013-1330302203133111-3101023222321313-3302013113202033)
- [rule_list.rules](data-sources--service_policy--reference--group-001.md#canonical-2332210012031020-3130121121203233-2133312023012212-1300213020002321-0210122201111201-2313301011130132-1102100221012102-3320333010302121)
- [rule_list.rules.spec](data-sources--service_policy--reference--group-001.md#canonical-1313133011110120-1202202112012022-0312300012021010-2230300132201133-1123333233212022-2100210133123323-2030333300021222-2110310211021210)
- rule_list.rules.spec.any_ip

<a id="canonical-1003211211011222-3123201310100102-0030313120303132-1101032013110102-1011030231231001-0020010301132211-2101301031032131-2211303133020200"></a>

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

<a id="canonical-0000330011302112-1212312313102203-0221331010012030-1233110131113232-0132032001100133-1233020003212002-1232220310301031-3113012300230333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules.spec.api_group_matcher` properties

Breadcrumbs:

- [xcsh_service_policy](../data-sources/service_policy.md#canonical-1112022302333333-1130100232232301-3023212310221022-3103031003221200-2030020130230132-1201210321332200-2212311003030121-1321310310032303)
- [Property reference](data-sources--service_policy--reference--group-001.md#canonical-3312102302111103-0221020013211300-2120323123300113-1101123102331230-1032322121222101-1332231300102132-0233300331323303-0031323122032013)
- [rule_list](data-sources--service_policy--reference--group-001.md#canonical-0200021302120322-3330103111223310-0110033233312323-2211001021302330-2000313233321013-1330302203133111-3101023222321313-3302013113202033)
- [rule_list.rules](data-sources--service_policy--reference--group-001.md#canonical-2332210012031020-3130121121203233-2133312023012212-1300213020002321-0210122201111201-2313301011130132-1102100221012102-3320333010302121)
- [rule_list.rules.spec](data-sources--service_policy--reference--group-001.md#canonical-1313133011110120-1202202112012022-0312300012021010-2230300132201133-1123333233212022-2100210133123323-2030333300021222-2110310211021210)
- rule_list.rules.spec.api_group_matcher

<a id="canonical-2111211320123130-0231032322312103-3203100222330213-3320320203102003-0313131031333320-3100030220111230-3213103121231300-3330103201102121"></a>

Type: `"single"`. Computed.

A matcher specifies a list of values for matching an input string. The match is considered
successful if the input value is present in the list. The result of the match is inverted if
invert\_matcher is true.

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

<a id="canonical-0132102133000101-1313021220022202-0223210220120200-1113010212210230-0031223120001112-3220333221121022-0102303023212331-0122303202230001"></a>

### Direct properties for `rule_list.rules.spec.api_group_matcher`

<a id="canonical-3112233332331211-0022302002321222-3203322302302200-2022221123130113-2200320121023302-0223332100303022-3321201300000000-3033331132301132"></a>

#### `rule_list.rules.spec.api_group_matcher.invert_matcher` property

Type: `"bool"`. Computed.

Invert String Matcher. Invert the match result.

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

<a id="canonical-2332001011133303-3010112113111113-0102001020312120-1000113333310200-2112213020130102-1200330031310033-3121120312111100-2021011002103010"></a>

<a id="canonical-0333200312000301-2020100113002221-2121110100030321-0023223101311231-0122012100100323-1003100223213213-1120220133321200-3302223103113223"></a>

#### `rule_list.rules.spec.api_group_matcher.match` property

Type: `["list", "string"]`. Computed.

A list of exact values to match the input against.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
    },
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
    "ves.io.schema.rules.repeated.items.string.max_bytes": "63",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.max_bytes": "63",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0322132120012130-1302303202022130-3333202213132012-3332222012222013-1122230102030120-3303120100313000-3131231103000201-3020233310033201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules.spec.arg_matchers` properties

Breadcrumbs:

- [xcsh_service_policy](../data-sources/service_policy.md#canonical-1112022302333333-1130100232232301-3023212310221022-3103031003221200-2030020130230132-1201210321332200-2212311003030121-1321310310032303)
- [Property reference](data-sources--service_policy--reference--group-001.md#canonical-3312102302111103-0221020013211300-2120323123300113-1101123102331230-1032322121222101-1332231300102132-0233300331323303-0031323122032013)
- [rule_list](data-sources--service_policy--reference--group-001.md#canonical-0200021302120322-3330103111223310-0110033233312323-2211001021302330-2000313233321013-1330302203133111-3101023222321313-3302013113202033)
- [rule_list.rules](data-sources--service_policy--reference--group-001.md#canonical-2332210012031020-3130121121203233-2133312023012212-1300213020002321-0210122201111201-2313301011130132-1102100221012102-3320333010302121)
- [rule_list.rules.spec](data-sources--service_policy--reference--group-001.md#canonical-1313133011110120-1202202112012022-0312300012021010-2230300132201133-1123333233212022-2100210133123323-2030333300021222-2110310211021210)
- rule_list.rules.spec.arg_matchers

<a id="canonical-0231002111103112-1231131210320221-1131032201121113-0310202302313003-2203202001203113-0110112201103112-1210001320120210-2212021302223002"></a>

Type: `"list"`. Computed.

A list of predicates for all POST args that need to be matched. The criteria for matching each arg
are described in individual instances of ArgMatcherType. The actual arg values are extracted from
the request API as a list of strings for each arg selector name. Note that all specified arg matcher
predicates must evaluate to true. A request body greater than 64KB will not be evaluated.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16"
  }
}
```

<a id="canonical-1000121220013100-1133210000023102-0300010201223211-0000120021313313-1203302231032102-3211121200300103-2013102132023231-2000022202113331"></a>

### Direct properties for `rule_list.rules.spec.arg_matchers`

- [check_not_present](data-sources--service_policy--reference--group-001.md#canonical-2003200000213131-0030010000212011-2201300000210130-1012010310100130-3330123123110012-1232130120332230-3032300310322030-2310030200022333): complete subsection reference.

- [check_present](data-sources--service_policy--reference--group-001.md#canonical-0312132223002221-3310011220203010-3030230313011212-2320203220300331-0023302210233011-0022122111313031-3003130203001002-0013230313301322): complete subsection reference.

<a id="canonical-0303220323213301-3123011333233123-2030210313111003-1213122331302020-3303010332200110-0000211303203020-1211102200013313-0223220211233121"></a>

<a id="canonical-1203030300011121-2222101002110212-2303233013222110-1020331111313010-1301203220200000-3333212333021012-2202212031013010-1201303020133200"></a>

#### `rule_list.rules.spec.arg_matchers.invert_matcher` property

Type: `"bool"`. Computed.

Invert Matcher. Invert Match of the expression defined.

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

- [item](data-sources--service_policy--reference--group-001.md#canonical-2033010332011023-3320020022223200-3031320122220001-3001131221121032-3012133133113200-1133101211332332-2211213221032213-0333320022301100): complete subsection reference.

<a id="canonical-2322232213331000-2122332021220233-1220132231233002-0212320233033130-0031023103030010-0233202121223223-0111133213013111-2301133322102121"></a>

<a id="canonical-1112031123301020-1211332023201033-2201332033232121-2103000322301021-3023311201021323-0023120330210000-0023213131031312-2103322120233103"></a>

#### `rule_list.rules.spec.arg_matchers.name` property

Type: `"string"`. Computed.

A case-sensitive JSON path in the HTTP request body.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 256
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
    "formatDescription": "DNS-1035 label: must start with a lowercase letter, may contain lowercase alphanumeric and hyphens, must end with alphanumeric",
    "maxLength": 63,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-06T12:36:10+00:00"
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
    "ves.io.schema.rules.string.json_path": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.json_path": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  }
}
```

<a id="canonical-2003200000213131-0030010000212011-2201300000210130-1012010310100130-3330123123110012-1232130120332230-3032300310322030-2310030200022333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules.spec.arg_matchers.check_not_present` properties

Breadcrumbs:

- [xcsh_service_policy](../data-sources/service_policy.md#canonical-1112022302333333-1130100232232301-3023212310221022-3103031003221200-2030020130230132-1201210321332200-2212311003030121-1321310310032303)
- [Property reference](data-sources--service_policy--reference--group-001.md#canonical-3312102302111103-0221020013211300-2120323123300113-1101123102331230-1032322121222101-1332231300102132-0233300331323303-0031323122032013)
- [rule_list](data-sources--service_policy--reference--group-001.md#canonical-0200021302120322-3330103111223310-0110033233312323-2211001021302330-2000313233321013-1330302203133111-3101023222321313-3302013113202033)
- [rule_list.rules](data-sources--service_policy--reference--group-001.md#canonical-2332210012031020-3130121121203233-2133312023012212-1300213020002321-0210122201111201-2313301011130132-1102100221012102-3320333010302121)
- [rule_list.rules.spec](data-sources--service_policy--reference--group-001.md#canonical-1313133011110120-1202202112012022-0312300012021010-2230300132201133-1123333233212022-2100210133123323-2030333300021222-2110310211021210)
- [rule_list.rules.spec.arg_matchers](data-sources--service_policy--reference--group-001.md#canonical-0322132120012130-1302303202022130-3333202213132012-3332222012222013-1122230102030120-3303120100313000-3131231103000201-3020233310033201)
- rule_list.rules.spec.arg_matchers.check_not_present

<a id="canonical-3121300003200003-2222002023101200-3000213023322300-1303323332113130-3232031132033301-0013131223012303-0121101101210311-2320311110330232"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for check not present.

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

<a id="canonical-0312132223002221-3310011220203010-3030230313011212-2320203220300331-0023302210233011-0022122111313031-3003130203001002-0013230313301322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules.spec.arg_matchers.check_present` properties

Breadcrumbs:

- [xcsh_service_policy](../data-sources/service_policy.md#canonical-1112022302333333-1130100232232301-3023212310221022-3103031003221200-2030020130230132-1201210321332200-2212311003030121-1321310310032303)
- [Property reference](data-sources--service_policy--reference--group-001.md#canonical-3312102302111103-0221020013211300-2120323123300113-1101123102331230-1032322121222101-1332231300102132-0233300331323303-0031323122032013)
- [rule_list](data-sources--service_policy--reference--group-001.md#canonical-0200021302120322-3330103111223310-0110033233312323-2211001021302330-2000313233321013-1330302203133111-3101023222321313-3302013113202033)
- [rule_list.rules](data-sources--service_policy--reference--group-001.md#canonical-2332210012031020-3130121121203233-2133312023012212-1300213020002321-0210122201111201-2313301011130132-1102100221012102-3320333010302121)
- [rule_list.rules.spec](data-sources--service_policy--reference--group-001.md#canonical-1313133011110120-1202202112012022-0312300012021010-2230300132201133-1123333233212022-2100210133123323-2030333300021222-2110310211021210)
- [rule_list.rules.spec.arg_matchers](data-sources--service_policy--reference--group-001.md#canonical-0322132120012130-1302303202022130-3333202213132012-3332222012222013-1122230102030120-3303120100313000-3131231103000201-3020233310033201)
- rule_list.rules.spec.arg_matchers.check_present

<a id="canonical-0202121332112331-1303132123020311-0302232021111320-2332233023302320-3101031213202022-1021232001032313-2002132122321201-1300212201021323"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for check present.

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

<a id="canonical-2033010332011023-3320020022223200-3031320122220001-3001131221121032-3012133133113200-1133101211332332-2211213221032213-0333320022301100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules.spec.arg_matchers.item` properties

Breadcrumbs:

- [xcsh_service_policy](../data-sources/service_policy.md#canonical-1112022302333333-1130100232232301-3023212310221022-3103031003221200-2030020130230132-1201210321332200-2212311003030121-1321310310032303)
- [Property reference](data-sources--service_policy--reference--group-001.md#canonical-3312102302111103-0221020013211300-2120323123300113-1101123102331230-1032322121222101-1332231300102132-0233300331323303-0031323122032013)
- [rule_list](data-sources--service_policy--reference--group-001.md#canonical-0200021302120322-3330103111223310-0110033233312323-2211001021302330-2000313233321013-1330302203133111-3101023222321313-3302013113202033)
- [rule_list.rules](data-sources--service_policy--reference--group-001.md#canonical-2332210012031020-3130121121203233-2133312023012212-1300213020002321-0210122201111201-2313301011130132-1102100221012102-3320333010302121)
- [rule_list.rules.spec](data-sources--service_policy--reference--group-001.md#canonical-1313133011110120-1202202112012022-0312300012021010-2230300132201133-1123333233212022-2100210133123323-2030333300021222-2110310211021210)
- [rule_list.rules.spec.arg_matchers](data-sources--service_policy--reference--group-001.md#canonical-0322132120012130-1302303202022130-3333202213132012-3332222012222013-1122230102030120-3303120100313000-3131231103000201-3020233310033201)
- rule_list.rules.spec.arg_matchers.item

<a id="canonical-2332023130032022-1132302310330132-2131001312321022-1113103313112333-3022313223011030-0221021122223210-2112321210013001-0313120133231001"></a>

Type: `"single"`. Computed.

A matcher specifies multiple criteria for matching an input string. The match is considered
successful if any of the criteria are satisfied. The set of supported match criteria includes a list
of exact values and a list of regular expressions.

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

<a id="canonical-0321320131031132-2203013030202001-2222011112331322-0010023303132222-0322230313201300-1313303201012330-2010112103311130-1123320330030120"></a>

### Direct properties for `rule_list.rules.spec.arg_matchers.item`

<a id="canonical-3213320220301231-2113020331222313-2010113133122330-2000320110223221-3013200303220123-0101100201110201-1233020101313330-3101013100121303"></a>

#### `rule_list.rules.spec.arg_matchers.item.exact_values` property

Type: `["list", "string"]`. Computed.

A list of exact values to match the input against.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1301323232211312-3312120320132331-2223023212110302-0300033232021011-2122312312132332-1120111323023123-0100121220311013-1321122113130100"></a>

<a id="canonical-1301133120310333-0103301320110332-2121321131232213-2320321312001312-3221213231212312-3231321111011112-3132023103233232-0211003002300333"></a>

#### `rule_list.rules.spec.arg_matchers.item.regex_values` property

Type: `["list", "string"]`. Computed.

A list of regular expressions to match the input against.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2303131320011000-0000330310010011-3121000200001321-0120033112301030-2111000021302120-2220000002100000-0333201011022013-0001133330130001"></a>

<a id="canonical-1212131220100300-0331012103310230-1323102321232231-3133133121002102-3002131332130023-1201111332230303-1020131022100233-2131012121202130"></a>

#### `rule_list.rules.spec.arg_matchers.item.transformers` property

Type: `["list", "string"]`. Computed.

\[Enum:
LOWER\_CASE|UPPER\_CASE|BASE64\_DECODE|NORMALIZE\_PATH|REMOVE\_WHITESPACE|URL\_DECODE|TRIM\_LEFT|TRIM\_RIGHT|TRIM\]
Ordered list of transformers (starting from index 0) to be applied to the path before matching.
Possible values are \`LOWER\_CASE\`, \`UPPER\_CASE\`, \`BASE64\_DECODE\`, \`NORMALIZE\_PATH\`,
\`REMOVE\_WHITESPACE\`, \`URL\_DECODE\`, \`TRIM\_LEFT\`, \`TRIM\_RIGHT\`, \`TRIM\`.

Additional upstream details:

An ordered list of transformers (starting from index 0) to be applied to the path before matching.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 9,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 9,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-06T12:36:10+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0003110021000230-3123203102130033-2313031311212300-3212013310032302-1030231200103320-2032330222011323-3320101221001100-1101302312333033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules.spec.asn_list` properties

Breadcrumbs:

- [xcsh_service_policy](../data-sources/service_policy.md#canonical-1112022302333333-1130100232232301-3023212310221022-3103031003221200-2030020130230132-1201210321332200-2212311003030121-1321310310032303)
- [Property reference](data-sources--service_policy--reference--group-001.md#canonical-3312102302111103-0221020013211300-2120323123300113-1101123102331230-1032322121222101-1332231300102132-0233300331323303-0031323122032013)
- [rule_list](data-sources--service_policy--reference--group-001.md#canonical-0200021302120322-3330103111223310-0110033233312323-2211001021302330-2000313233321013-1330302203133111-3101023222321313-3302013113202033)
- [rule_list.rules](data-sources--service_policy--reference--group-001.md#canonical-2332210012031020-3130121121203233-2133312023012212-1300213020002321-0210122201111201-2313301011130132-1102100221012102-3320333010302121)
- [rule_list.rules.spec](data-sources--service_policy--reference--group-001.md#canonical-1313133011110120-1202202112012022-0312300012021010-2230300132201133-1123333233212022-2100210133123323-2030333300021222-2110310211021210)
- rule_list.rules.spec.asn_list

<a id="canonical-1030233020210002-0201201310003213-2003323131101020-2023121200110312-0202012201302331-2013100012231021-0132003213302302-0031123023213131"></a>

Type: `"single"`. Computed.

An unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny
lists for use in network policy or service policy. It can be used to create the allow list only for
DNS Load Balancer.

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

<a id="canonical-1021211122231313-3100020230330022-0013321202033012-1200103123300133-1333120031102020-0130202102113133-3121102123132311-0033102000201032"></a>

### Direct properties for `rule_list.rules.spec.asn_list`

<a id="canonical-0230111022020120-2312230000130022-2222323112130003-3202012213102102-3130023312132111-3230231131211301-0231212213230320-0310101302133001"></a>

#### `rule_list.rules.spec.asn_list.as_numbers` property

Type: `["list", "number"]`. Computed.

An unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny
lists for use in network policy or service policy. It can be used to create the allow list only for
DNS Load Balancer.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 16,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 16,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-06T12:36:10+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3003212230320021-1130112110220111-3212200301310132-0123111321310113-3213001310131320-2031031330231203-1132023221020001-3002123211023021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules.spec.asn_matcher` properties

Breadcrumbs:

- [xcsh_service_policy](../data-sources/service_policy.md#canonical-1112022302333333-1130100232232301-3023212310221022-3103031003221200-2030020130230132-1201210321332200-2212311003030121-1321310310032303)
- [Property reference](data-sources--service_policy--reference--group-001.md#canonical-3312102302111103-0221020013211300-2120323123300113-1101123102331230-1032322121222101-1332231300102132-0233300331323303-0031323122032013)
- [rule_list](data-sources--service_policy--reference--group-001.md#canonical-0200021302120322-3330103111223310-0110033233312323-2211001021302330-2000313233321013-1330302203133111-3101023222321313-3302013113202033)
- [rule_list.rules](data-sources--service_policy--reference--group-001.md#canonical-2332210012031020-3130121121203233-2133312023012212-1300213020002321-0210122201111201-2313301011130132-1102100221012102-3320333010302121)
- [rule_list.rules.spec](data-sources--service_policy--reference--group-001.md#canonical-1313133011110120-1202202112012022-0312300012021010-2230300132201133-1123333233212022-2100210133123323-2030333300021222-2110310211021210)
- rule_list.rules.spec.asn_matcher

<a id="canonical-3310301311023211-3212333103231332-1301223331002130-1323130023300133-1200010303133320-1231002331210313-3003231230221310-1103000103033223"></a>

Type: `"single"`. Computed.

Match any AS number contained in the list of bgp\_asn\_sets.

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

<a id="canonical-2330113223221111-2011303210333212-3321013313020131-0310223023123222-2020202211111130-1133130223331033-2103010013200311-0231033333020101"></a>

### Direct properties for `rule_list.rules.spec.asn_matcher`

- [asn_sets](data-sources--service_policy--reference--group-001.md#canonical-2210012001102210-3202102103331221-3030211102203232-2201311003101123-2200102101211312-2121312011323013-3030323231111130-3032033123303113): complete subsection reference.

<a id="canonical-2210012001102210-3202102103331221-3030211102203232-2201311003101123-2200102101211312-2121312011323013-3030323231111130-3032033123303113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules.spec.asn_matcher.asn_sets` properties

Breadcrumbs:

- [xcsh_service_policy](../data-sources/service_policy.md#canonical-1112022302333333-1130100232232301-3023212310221022-3103031003221200-2030020130230132-1201210321332200-2212311003030121-1321310310032303)
- [Property reference](data-sources--service_policy--reference--group-001.md#canonical-3312102302111103-0221020013211300-2120323123300113-1101123102331230-1032322121222101-1332231300102132-0233300331323303-0031323122032013)
- [rule_list](data-sources--service_policy--reference--group-001.md#canonical-0200021302120322-3330103111223310-0110033233312323-2211001021302330-2000313233321013-1330302203133111-3101023222321313-3302013113202033)
- [rule_list.rules](data-sources--service_policy--reference--group-001.md#canonical-2332210012031020-3130121121203233-2133312023012212-1300213020002321-0210122201111201-2313301011130132-1102100221012102-3320333010302121)
- [rule_list.rules.spec](data-sources--service_policy--reference--group-001.md#canonical-1313133011110120-1202202112012022-0312300012021010-2230300132201133-1123333233212022-2100210133123323-2030333300021222-2110310211021210)
- [rule_list.rules.spec.asn_matcher](data-sources--service_policy--reference--group-001.md#canonical-3003212230320021-1130112110220111-3212200301310132-0123111321310113-3213001310131320-2031031330231203-1132023221020001-3002123211023021)
- rule_list.rules.spec.asn_matcher.asn_sets

<a id="canonical-1200202203313103-0013212021220122-2311212212023102-1101120113310332-3303303331023210-0033221122022102-2103312010010023-3020300102222003"></a>

Type: `"list"`. Computed.

A list of references to bgp\_asn\_set objects.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-1321133113320210-1330333221031310-3313013222331213-0203231311032030-0223012123033203-0100013313200120-0200231030032230-0210130222333001"></a>

### Direct properties for `rule_list.rules.spec.asn_matcher.asn_sets`

<a id="canonical-3213033302123332-1030130013021222-1100202011200203-3232232312322231-3332120020003022-2111203320311011-0213130321310023-1331211301220211"></a>

#### `rule_list.rules.spec.asn_matcher.asn_sets.kind` property

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-0022113003211033-1230333333111332-0103301202123231-1212133320020313-0323010110222021-1120102322021200-3123001011132311-0233120231022310"></a>

<a id="canonical-1322223120333020-3023303200113020-3123202122300023-3121312320333213-0233223333220003-1200231330000322-0030030212033021-1113123333230311"></a>

#### `rule_list.rules.spec.asn_matcher.asn_sets.name` property

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-0012112203232311-0022022213112112-0310331232100202-2312202201030312-3312213301020310-0030313001200330-3022131320102210-3032030103212003"></a>

<a id="canonical-0001303001303323-1013133230020101-3301111210012122-2100303130021233-0012131312323220-2023201223113313-1133011320201013-2003113323302310"></a>

#### `rule_list.rules.spec.asn_matcher.asn_sets.namespace` property

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-2230200321222023-0122123120032212-3210101013113231-0001200330101313-2030201022302232-2321112011300233-1211230210212321-1203113212322220"></a>

<a id="canonical-2023230330311201-3221100021301010-2201200233021221-0030012121131103-3122211233122212-2000012310320320-3101120201313232-1103232132120033"></a>

#### `rule_list.rules.spec.asn_matcher.asn_sets.tenant` property

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-1033130001323331-3022202233230230-1211200300222233-3321232100201133-2303321102230031-3232032111123102-2102222320112332-1330332112133133"></a>

<a id="canonical-1231223320130232-2111230231310000-1210213103211303-1000130310211132-1232120312202333-0301132321323112-0011112133323310-3313203010313022"></a>

#### `rule_list.rules.spec.asn_matcher.asn_sets.uid` property

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-3120003330322333-0013033122113213-1031303002011013-0121111101233222-2001211210033333-2313220132320101-0011222020013012-0022323101113322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules.spec.body_matcher` properties

Breadcrumbs:

- [xcsh_service_policy](../data-sources/service_policy.md#canonical-1112022302333333-1130100232232301-3023212310221022-3103031003221200-2030020130230132-1201210321332200-2212311003030121-1321310310032303)
- [Property reference](data-sources--service_policy--reference--group-001.md#canonical-3312102302111103-0221020013211300-2120323123300113-1101123102331230-1032322121222101-1332231300102132-0233300331323303-0031323122032013)
- [rule_list](data-sources--service_policy--reference--group-001.md#canonical-0200021302120322-3330103111223310-0110033233312323-2211001021302330-2000313233321013-1330302203133111-3101023222321313-3302013113202033)
- [rule_list.rules](data-sources--service_policy--reference--group-001.md#canonical-2332210012031020-3130121121203233-2133312023012212-1300213020002321-0210122201111201-2313301011130132-1102100221012102-3320333010302121)
- [rule_list.rules.spec](data-sources--service_policy--reference--group-001.md#canonical-1313133011110120-1202202112012022-0312300012021010-2230300132201133-1123333233212022-2100210133123323-2030333300021222-2110310211021210)
- rule_list.rules.spec.body_matcher

<a id="canonical-1100103221313312-1312213330003231-2032111201030011-3133232310132122-0013222333002231-3223303322013223-0230113211022321-1120220132031332"></a>

Type: `"single"`. Computed.

A matcher specifies multiple criteria for matching an input string. The match is considered
successful if any of the criteria are satisfied. The set of supported match criteria includes a list
of exact values and a list of regular expressions.

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

<a id="canonical-0333013003333221-3110223222303330-3201032310102000-2303102132221312-0131010203112130-1103203211222221-3203310230111322-0112230102010332"></a>

### Direct properties for `rule_list.rules.spec.body_matcher`

<a id="canonical-0022310003233020-1313003133110303-1010122030223103-0130120033122203-0212313310003123-1102202301223313-3203213130223100-0021131013200300"></a>

#### `rule_list.rules.spec.body_matcher.exact_values` property

Type: `["list", "string"]`. Computed.

A list of exact values to match the input against.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0020121230121202-3110113013122330-0233120202231010-1213321133313201-2302100030201011-2201123312003000-2031111110300123-2111131013013021"></a>

<a id="canonical-2310323122130110-2311203332230231-1301121303302010-1302321312102032-2300030013321332-3223001302230222-2003023303311212-1001212210110012"></a>

#### `rule_list.rules.spec.body_matcher.regex_values` property

Type: `["list", "string"]`. Computed.

A list of regular expressions to match the input against.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1311120131232022-0012013333132311-3312020230302200-2330022303322300-3310022211101002-3130222112032300-1202032303113200-1331221103232111"></a>

<a id="canonical-0211311101021333-2323123321031002-2123212301003212-3302011321003322-1113313123102010-3213330322010313-3300230312123002-3212133101310323"></a>

#### `rule_list.rules.spec.body_matcher.transformers` property

Type: `["list", "string"]`. Computed.

\[Enum:
LOWER\_CASE|UPPER\_CASE|BASE64\_DECODE|NORMALIZE\_PATH|REMOVE\_WHITESPACE|URL\_DECODE|TRIM\_LEFT|TRIM\_RIGHT|TRIM\]
Ordered list of transformers (starting from index 0) to be applied to the path before matching.
Possible values are \`LOWER\_CASE\`, \`UPPER\_CASE\`, \`BASE64\_DECODE\`, \`NORMALIZE\_PATH\`,
\`REMOVE\_WHITESPACE\`, \`URL\_DECODE\`, \`TRIM\_LEFT\`, \`TRIM\_RIGHT\`, \`TRIM\`.

Additional upstream details:

An ordered list of transformers (starting from index 0) to be applied to the path before matching.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 9,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 9,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-06T12:36:10+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```
